/*
Provides a resource to manage TKE cluster user permissions via RBAC.

# Example Usage

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_user_permissions" "example" {
	  cluster_id = "cls-xxxxxxxx"
	  target_uin = "110000000000"

	  permissions {
	    role_name = "tke:admin"
	    role_type = "cluster"
	  }

	  permissions {
	    role_name = "tke:dev"
	    role_type = "namespace"
	    namespace = "default"
	  }
	}

```

# Import

terraform import tencentcloudenterprise_tke_kubernetes_user_permissions.example cls-xxxxxxxx:110000000000
*/
package tencentcloud

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// tkeRoleNameMap maps friendly role names from console to K8s ClusterRole names.
var tkeRoleNameMap = map[string]string{
	"admin":           "tke:admin",
	"tke:admin":       "tke:admin",
	"ops":             "tke:ops",
	"tke:ops":         "tke:ops",
	"ops team":        "tke:ops",
	"developer":       "tke:dev",
	"tke:dev":         "tke:dev",
	"read-only":       "tke:ro",
	"tke:ro":          "tke:ro",
	"read-only users": "tke:ro",
	"namespace dev":   "tke:ns:dev",
	"tke:ns:dev":      "tke:ns:dev",
	"namespace ro":    "tke:ns:ro",
	"tke:ns:ro":       "tke:ns:ro",
}

// normalizeRoleName converts user input to the K8s ClusterRole name.
func normalizeRoleName(input string) string {
	lower := strings.ToLower(strings.TrimSpace(input))
	if mapped, ok := tkeRoleNameMap[lower]; ok {
		return mapped
	}
	return input
}

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_user_permissions", CNDescription{
		TerraformTypeCN: "TKE 集群用户权限",
		DescriptionCN:   "用于管理 TKE 集群用户 RBAC 权限。通过声明式语义管理子账号在指定集群中的 ClusterRoleBinding 和 RoleBinding。",
		AttributesCN: map[string]string{
			"cluster_id":  "集群 ID。",
			"target_uin":  "被授权的用户 UIN（支持子账号 UIN 和角色 UIN）。",
			"permissions": "用户最终应拥有的完整权限列表。使用声明式语义，传入的列表代表用户最终应有的所有权限，系统会自动计算差异并执行增删操作。",
			"role_name":   "角色名称。预置角色：tke:admin（集群管理员）、tke:ops（运维人员）、tke:dev（开发者）、tke:ro（只读用户）、tke:ns:dev（命名空间开发者）、tke:ns:ro（命名空间只读用户），其他为自定义角色。",
			"role_type":   "授权类型。cluster：集群级权限（ClusterRoleBinding），namespace：命名空间级权限（RoleBinding）。",
			"is_custom":   "是否为自定义角色，默认 false。",
			"namespace":   "命名空间。role_type 为 namespace 时必填。",
		},
	})
}

func resourceTencentCloudTkeKubernetesUserPermissions() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to manage TKE cluster user permissions via RBAC.",
		Create:      resourceTencentCloudTkeKubernetesUserPermissionsCreate,
		Read:        resourceTencentCloudTkeKubernetesUserPermissionsRead,
		Update:      resourceTencentCloudTkeKubernetesUserPermissionsUpdate,
		Delete:      resourceTencentCloudTkeKubernetesUserPermissionsDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		CustomizeDiff: func(_ context.Context, d *schema.ResourceDiff, _ interface{}) error {
			return validateTkeUserPermissions(d.Get("permissions"))
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Cluster ID.",
			},
			"target_uin": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique identifier of the user to be authorized (supports sub-account UIN and role UIN).",
			},
			"permissions": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Complete list of permissions that the user should ultimately have. Uses declarative semantics, the passed list represents all permissions the user should ultimately have, the system will automatically calculate differences and perform necessary create/delete operations. When empty or not provided, all permissions for this user will be cleared.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Role name. Predefined roles include: tke:admin (cluster administrator), tke:ops (operations personnel), tke:dev (developer), tke:ro (read-only user), tke:ns:dev (namespace developer), tke:ns:ro (namespace read-only user), others are user-defined roles.",
						},
						"role_type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"cluster", "namespace"}, false),
							Description:  "Authorization type. Enum values: cluster (cluster-level permissions, corresponding to ClusterRoleBinding), namespace (namespace-level permissions, corresponding to RoleBinding).",
						},
						"is_custom": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether it is a custom role, default false.",
						},
						"namespace": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Namespace. Required when role_type is namespace.",
						},
					},
				},
			},
		},
	}
}

func validateTkeUserPermissions(value interface{}) error {
	permissions, ok := value.(*schema.Set)
	if !ok {
		return nil
	}

	for i, item := range permissions.List() {
		permission, ok := item.(map[string]interface{})
		if !ok {
			return fmt.Errorf("permissions[%d] has an invalid value", i)
		}

		roleType, _ := permission["role_type"].(string)
		roleType = strings.TrimSpace(roleType)
		namespace := ""
		if value, ok := permission["namespace"].(string); ok {
			namespace = strings.TrimSpace(value)
		}

		switch roleType {
		case "namespace":
			if namespace == "" {
				return fmt.Errorf("permissions[%d].namespace must be set when role_type is \"namespace\"", i)
			}
		case "cluster":
			if namespace != "" {
				return fmt.Errorf("permissions[%d].namespace must not be set when role_type is \"cluster\"", i)
			}
		default:
			return fmt.Errorf("permissions[%d].role_type must be either \"cluster\" or \"namespace\"", i)
		}
	}

	return nil
}

// buildClusterRoleBindingYAML builds a ClusterRoleBinding YAML for apply.
func buildClusterRoleBindingYAML(uin, nickname, subjectName, roleName string) string {
	tmpl := `{"apiVersion":"rbac.authorization.k8s.io/v1","kind":"ClusterRoleBinding","metadata":{"name":"%s-ClusterRole","labels":{"cloud.tencent.com/tke-account":"%s"},"annotations":{"cloud.tencent.com/tke-account-nickname":"%s"}},"roleRef":{"apiGroup":"rbac.authorization.k8s.io","kind":"ClusterRole","name":"%s"},"subjects":[{"kind":"User","name":"%s"}]}`
	return fmt.Sprintf(tmpl, uin, uin, nickname, roleName, subjectName)
}

// buildRoleBindingYAML builds a RoleBinding YAML for apply.
func buildRoleBindingYAML(uin, nickname, subjectName, roleName, namespace string) string {
	tmpl := `{"apiVersion":"rbac.authorization.k8s.io/v1","kind":"RoleBinding","metadata":{"name":"%s-Role-%s","namespace":"%s","labels":{"cloud.tencent.com/tke-account":"%s"},"annotations":{"cloud.tencent.com/tke-account-nickname":"%s"}},"roleRef":{"apiGroup":"rbac.authorization.k8s.io","kind":"ClusterRole","name":"%s"},"subjects":[{"kind":"User","name":"%s"}]}`
	return fmt.Sprintf(tmpl, uin, namespace, namespace, uin, nickname, roleName, subjectName)
}

// resolveTkeAccountNickname resolves a sub-account UIN for console display.
// Role UINs and lookup failures fall back to the UIN to preserve compatibility.
func resolveTkeAccountNickname(ctx context.Context, client *TencentCloudClient, targetUin string) string {
	uin, err := strconv.ParseInt(targetUin, 10, 64)
	if err != nil {
		return targetUin
	}

	service := CamService{client: client.apiV3Conn}
	userName, err := service.DescribeUserNameByUin(ctx, uin)
	if err != nil {
		log.Printf("[WARN] resolve CAM user name for UIN %s failed: %v", targetUin, err)
		return targetUin
	}
	if userName == "" {
		return targetUin
	}

	return userName
}

func resolveTkeSubjectName(ctx context.Context, client *TencentCloudClient, clusterId, targetUin string) (string, error) {
	service := TkeService{client: client.apiV3Conn}
	commonName, err := service.DescribeClusterCommonName(ctx, clusterId, targetUin)
	if err != nil {
		return "", err
	}
	if commonName == "" {
		return "", fmt.Errorf("no cluster certificate common name found for account %s in cluster %s", targetUin, clusterId)
	}
	return commonName, nil
}

// buildApplyPath returns the apply endpoint path for a given cluster.
func buildApplyPath(clusterId string) string {
	return fmt.Sprintf("/apis/platform.tke/v1/clusters/%s/apply", clusterId)
}

// applyBinding calls the TCE apply endpoint to create/update a K8s RBAC binding.
func applyBinding(me *TkeService, ctx context.Context, clusterId, yamlBody string) error {
	encodedBody := base64.StdEncoding.EncodeToString([]byte(yamlBody))
	path := buildApplyPath(clusterId)
	encoded := true
	_, err := me.ForwardPlatformRequestV3WithOptions(ctx, "POST", path, clusterId, encodedBody, "", "", &encoded)
	return err
}

// deleteBinding deletes a K8s RBAC binding.
func deleteBinding(me *TkeService, ctx context.Context, clusterId, path string) error {
	_, err := me.ForwardPlatformRequestV3(ctx, "DELETE", path, clusterId, "")
	return err
}

// parseBindingList parses a K8s binding list JSON response and returns binding names and their roleRef.
func parseBindingList(responseBody string) ([]map[string]string, error) {
	var result struct {
		Items []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			RoleRef struct {
				Name string `json:"name"`
			} `json:"roleRef"`
		} `json:"items"`
	}
	if err := json.Unmarshal([]byte(responseBody), &result); err != nil {
		return nil, err
	}
	var bindings []map[string]string
	for _, item := range result.Items {
		bindings = append(bindings, map[string]string{
			"name":     item.Metadata.Name,
			"roleName": item.RoleRef.Name,
		})
	}
	return bindings, nil
}

// bindingKey generates a unique key for a permission for diff comparison.
func bindingKey(perm map[string]interface{}) string {
	roleName := normalizeRoleName(perm["role_name"].(string))
	roleType := perm["role_type"].(string)
	namespace := ""
	if v, ok := perm["namespace"].(string); ok {
		namespace = v
	}
	if roleType == "namespace" {
		return fmt.Sprintf("%s|%s|%s", roleName, roleType, namespace)
	}
	return fmt.Sprintf("%s|%s", roleName, roleType)
}

func resourceTencentCloudTkeKubernetesUserPermissionsCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_user_permissions.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var clusterId, targetUin string
	if v, ok := d.GetOk("cluster_id"); ok {
		clusterId = v.(string)
	}
	if v, ok := d.GetOk("target_uin"); ok {
		targetUin = v.(string)
	}
	if err := validateTkeUserPermissions(d.Get("permissions")); err != nil {
		return err
	}

	if v, ok := d.GetOk("permissions"); ok {
		permissions := v.(*schema.Set).List()
		if len(permissions) == 0 {
			d.SetId(fmt.Sprintf("%s:%s", clusterId, targetUin))
			return resourceTencentCloudTkeKubernetesUserPermissionsRead(d, meta)
		}

		nickname := resolveTkeAccountNickname(ctx, meta.(*TencentCloudClient), targetUin)
		subjectName, err := resolveTkeSubjectName(ctx, meta.(*TencentCloudClient), clusterId, targetUin)
		if err != nil {
			return err
		}

		for _, item := range permissions {
			permMap := item.(map[string]interface{})
			roleName := normalizeRoleName(permMap["role_name"].(string))
			roleType := permMap["role_type"].(string)
			namespace := ""
			if v, ok := permMap["namespace"].(string); ok {
				namespace = v
			}

			var yamlBody string
			if roleType == "namespace" {
				yamlBody = buildRoleBindingYAML(targetUin, nickname, subjectName, roleName, namespace)
			} else {
				yamlBody = buildClusterRoleBindingYAML(targetUin, nickname, subjectName, roleName)
			}

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				if e := applyBinding(&service, ctx, clusterId, yamlBody); e != nil {
					return retryError(e)
				}
				return nil
			})
			if err != nil {
				log.Printf("[CRITAL]%s create kubernetes user permission failed, clusterId: %s, role: %s, reason:%+v",
					logId, clusterId, roleName, err)
				return err
			}
		}
	}

	d.SetId(fmt.Sprintf("%s:%s", clusterId, targetUin))
	return resourceTencentCloudTkeKubernetesUserPermissionsRead(d, meta)
}

func resourceTencentCloudTkeKubernetesUserPermissionsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_user_permissions.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	idParts := strings.SplitN(d.Id(), ":", 2)
	if len(idParts) != 2 {
		return fmt.Errorf("invalid ID format, expected cluster_id:target_uin, got: %s", d.Id())
	}
	clusterId := idParts[0]
	targetUin := idParts[1]

	_ = d.Set("cluster_id", clusterId)
	_ = d.Set("target_uin", targetUin)

	bindings, err := readUserPermissionsByCluster(&service, ctx, targetUin, clusterId)
	if err != nil {
		log.Printf("[WARN]%s read permissions for cluster %s failed: %v", logId, clusterId, err)
	}

	if len(bindings) == 0 {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_tke_kubernetes_user_permissions` [%s] has no permissions, please check if it has been deleted.\n", logId, d.Id())
	}

	_ = d.Set("permissions", bindings)
	return nil
}

func readUserPermissionsByCluster(service *TkeService, ctx context.Context, uin, clusterId string) ([]map[string]interface{}, error) {
	var permissions []map[string]interface{}

	// Read ClusterRoleBindings
	crbBody, err := service.DescribeClusterRoleBindingsByUin(ctx, clusterId, uin)
	if err == nil && crbBody != "" {
		bindings, parseErr := parseBindingList(crbBody)
		if parseErr == nil {
			for _, b := range bindings {
				permissions = append(permissions, map[string]interface{}{
					"role_name": b["roleName"],
					"role_type": "cluster",
					"is_custom": !isPredefinedRole(b["roleName"]),
				})
			}
		}
	}

	// Read RoleBindings (namespace level)
	rbBody, err := service.DescribeRoleBindingsByUin(ctx, clusterId, uin)
	if err == nil && rbBody != "" {
		bindings, parseErr := parseBindingList(rbBody)
		if parseErr == nil {
			for _, b := range bindings {
				// Extract namespace from binding name pattern: {uin}-Role-{namespace}
				namespace := extractNamespaceFromBindingName(b["name"], uin)
				permissions = append(permissions, map[string]interface{}{
					"role_name": b["roleName"],
					"role_type": "namespace",
					"namespace": namespace,
					"is_custom": !isPredefinedRole(b["roleName"]),
				})
			}
		}
	}

	return permissions, nil
}

func isPredefinedRole(roleName string) bool {
	predefined := map[string]bool{
		"tke:admin":  true,
		"tke:ops":    true,
		"tke:dev":    true,
		"tke:ro":     true,
		"tke:ns:dev": true,
		"tke:ns:ro":  true,
	}
	return predefined[roleName]
}

func extractNamespaceFromBindingName(name, uin string) string {
	prefix := fmt.Sprintf("%s-Role-", uin)
	if strings.HasPrefix(name, prefix) {
		return strings.TrimPrefix(name, prefix)
	}
	return ""
}

func resourceTencentCloudTkeKubernetesUserPermissionsUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_user_permissions.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	idParts := strings.SplitN(d.Id(), ":", 2)
	if len(idParts) != 2 {
		return fmt.Errorf("invalid ID format, expected cluster_id:target_uin, got: %s", d.Id())
	}
	clusterId := idParts[0]
	targetUin := idParts[1]

	if d.HasChange("permissions") {
		oldInterface, newInterface := d.GetChange("permissions")
		if err := validateTkeUserPermissions(newInterface); err != nil {
			return err
		}
		olds := oldInterface.(*schema.Set)
		news := newInterface.(*schema.Set)
		remove := olds.Difference(news).List()
		add := news.Difference(olds).List()
		nickname := ""
		subjectName := ""
		if len(add) > 0 {
			nickname = resolveTkeAccountNickname(ctx, meta.(*TencentCloudClient), targetUin)
			resolvedSubjectName, resolveErr := resolveTkeSubjectName(ctx, meta.(*TencentCloudClient), clusterId, targetUin)
			if resolveErr != nil {
				return resolveErr
			}
			subjectName = resolvedSubjectName
		}

		// Remove old permissions
		for _, item := range remove {
			permMap := item.(map[string]interface{})
			roleType := permMap["role_type"].(string)
			namespace := ""
			if v, ok := permMap["namespace"].(string); ok {
				namespace = v
			}

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				var deleteErr error
				if roleType == "namespace" {
					bindingName := fmt.Sprintf("%s-Role-%s", targetUin, namespace)
					path := fmt.Sprintf("/apis/rbac.authorization.k8s.io/v1/namespaces/%s/rolebindings/%s", namespace, bindingName)
					deleteErr = deleteBinding(&service, ctx, clusterId, path)
				} else {
					bindingName := fmt.Sprintf("%s-ClusterRole", targetUin)
					path := fmt.Sprintf("/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/%s", bindingName)
					deleteErr = deleteBinding(&service, ctx, clusterId, path)
				}
				if deleteErr != nil {
					log.Printf("[WARN]%s delete binding failed (may already not exist), err: %v", logId, deleteErr)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}

		// Add new permissions
		for _, item := range add {
			permMap := item.(map[string]interface{})
			roleName := normalizeRoleName(permMap["role_name"].(string))
			roleType := permMap["role_type"].(string)
			namespace := ""
			if v, ok := permMap["namespace"].(string); ok {
				namespace = v
			}

			var yamlBody string
			if roleType == "namespace" {
				yamlBody = buildRoleBindingYAML(targetUin, nickname, subjectName, roleName, namespace)
			} else {
				yamlBody = buildClusterRoleBindingYAML(targetUin, nickname, subjectName, roleName)
			}

			err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				if e := applyBinding(&service, ctx, clusterId, yamlBody); e != nil {
					return retryError(e)
				}
				return nil
			})
			if err != nil {
				log.Printf("[CRITAL]%s create kubernetes user permission failed, clusterId: %s, role: %s, reason:%+v",
					logId, clusterId, roleName, err)
				return err
			}
		}
	}

	return resourceTencentCloudTkeKubernetesUserPermissionsRead(d, meta)
}

func resourceTencentCloudTkeKubernetesUserPermissionsDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_user_permissions.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	idParts := strings.SplitN(d.Id(), ":", 2)
	if len(idParts) != 2 {
		return fmt.Errorf("invalid ID format, expected cluster_id:target_uin, got: %s", d.Id())
	}
	clusterId := idParts[0]
	targetUin := idParts[1]

	// Delete ClusterRoleBinding
	crbName := fmt.Sprintf("%s-ClusterRole", targetUin)
	crbPath := fmt.Sprintf("/apis/rbac.authorization.k8s.io/v1/clusterrolebindings/%s", crbName)
	if err := deleteBinding(&service, ctx, clusterId, crbPath); err != nil {
		log.Printf("[WARN]%s delete ClusterRoleBinding %s in cluster %s failed: %v", logId, crbName, clusterId, err)
	}

	// Delete all RoleBindings for this uin
	rbBody, err := service.DescribeRoleBindingsByUin(ctx, clusterId, targetUin)
	if err == nil && rbBody != "" {
		bindings, parseErr := parseBindingList(rbBody)
		if parseErr == nil {
			for _, b := range bindings {
				ns := extractNamespaceFromBindingName(b["name"], targetUin)
				if ns == "" {
					continue
				}
				rbPath := fmt.Sprintf("/apis/rbac.authorization.k8s.io/v1/namespaces/%s/rolebindings/%s", ns, b["name"])
				if err := deleteBinding(&service, ctx, clusterId, rbPath); err != nil {
					log.Printf("[WARN]%s delete RoleBinding %s in cluster %s ns %s failed: %v", logId, b["name"], clusterId, ns, err)
				}
			}
		}
	}

	return nil
}

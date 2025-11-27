/*
Provide a resource to increase instance to cluster

~> **NOTE:** To use the custom Kubernetes component startup parameter function (parameter `extra_args`), you need to submit a ticket for application.

# Example Usage

```hcl

	resource "tencentcloudenterprise_kubernetes_cluster_secret" "app-csp-sm" {
	  cluster_id = cloud_tke_kubernetes_cluster.cluster.id
	  namespace  = "app-csp-sm"
	  path = "/apis/platform.tkestack.io/v1/clusters/cls-x8lxd2jx/apply"
	  request_body = "{\"kind\":\"Secret\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"app-csp-sm\",\"annotations\":{\"description\":\"hkjc1\"}}}{\"kind\":\"Secret\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"qcloudregistrykey\",\"namespace\":\"app-csp-sm\",\"labels\":{\"qcloud-app\":\"qcloudregistrykey\"}},\"type\":\"kubernetes.io/dockercfg\",\"data\":{\".dockercfg\":\"eyJjY3IudGNlMzEwMHBvYy5mc3BoZXJlLmNuIjp7InVzZXJuYW1lIjoiMTAwMDA0NjAzMTU3IiwicGFzc3dvcmQiOiJ7QXBwbGljYXRpb25Ub2tlbjo0OGJlNzY2ZTVkZmRmN2JhZTAwZjdlZTQ3NTQyNDJlMX0iLCJlbWFpbCI6Im5vdEB2YWwuaWQiLCJhdXRoIjoiTVRBd01EQTBOakF6TVRVM09udEJjSEJzYVdOaGRHbHZibFJ2YTJWdU9qUTRZbVUzTmpabE5XUm1aR1kzWW1GbE1EQm1OMlZsTkRjMU5ESTBNbVV4ZlE9PSJ9fQ==\"}}"
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"
	// "encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_cluster_secret", CNDescription{
		TerraformTypeCN: "集群密钥配置",
		DescriptionCN:   "提供集群密钥配置资源，用于向集群增加密钥。",
		AttributesCN: map[string]string{
			"cluster_id":   "集群ID",
			"namespace":    "命名空间名称",
			"path":         "请求路径",
			"request_body": "请求体",
			"secret_name":  "密钥名称",
		},
	})
}

type SecretList struct {
	Kind       string       `json:"kind"`
	ApiVersion string       `json:"apiVersion"`
	Metadata   ObjectMeta   `json:"metadata"`
	Items      []SecretItem `json:"items"`
}

type ObjectMeta struct {
	SelfLink          string            `json:"selfLink"`
	ResourceVersion   string            `json:"resourceVersion"`
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Uid               string            `json:"uid"`
	CreationTimestamp string            `json:"creationTimestamp"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	ManagedFields     []ManagedField    `json:"managedFields"`
}

type ManagedField struct {
	Manager    string                 `json:"manager"`
	Operation  string                 `json:"operation"`
	ApiVersion string                 `json:"apiVersion"`
	Time       string                 `json:"time"`
	FieldsType string                 `json:"fieldsType"`
	FieldsV1   map[string]interface{} `json:"fieldsV1"`
}

type SecretItem struct {
	Metadata      ObjectMeta        `json:"metadata"`
	Data          map[string]string `json:"data"`
	Type          string            `json:"type"`
	ManagedFields []ManagedField    `json:"managedFields"`
}

type TkeForwardRequestSecretReadResponse struct {
	MetaData map[string]interface{} `json:"metadata"`
	Code     int                    `json:"code"`
	Status   string                 `json:"status"`
	Message  string                 `json:"message"`
}

func resourceTencentCloudTkeClusterSecret() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to increase ns to cluster",
		Create:      resourceTencentCloudTkeTkeClusterSecretCreate,
		Read:        resourceTencentCloudTkeTkeClusterSecretRead,
		Delete:      resourceTencentCloudTkeTkeClusterSecretDelete,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "ID of the cluster.",
			},
			"path": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "ns name",
			},
			"request_body": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "request_body",
			},
			"namespace": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "namespace",
			},
			"secret_name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "secret_name",
			},
		},
	}
}

func resourceTencentCloudTkeTkeClusterSecretCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_secret.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	var (
		clusterId   = d.Get("cluster_id").(string)
		path        = d.Get("path").(string)
		requestBody = d.Get("request_body").(string)

		secret = d.Get("secret_name").(string)
	)
	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_POST, path, clusterId, requestBody)
	if err != nil {
		return err
	}
	d.SetId(secret)
	return resourceTencentCloudTkeTkeClusterSecretRead(d, meta)
}

func resourceTencentCloudTkeTkeClusterSecretRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_secret.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var (
		ns        = d.Get("namespace").(string)
		secret    = d.Get("secret_name").(string)
		clusterId = d.Get("cluster_id").(string)
		path      = fmt.Sprintf("/api/v1/namespaces/%s/secrets?fieldSelector=metadata.name=%s", ns, secret)
	)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_GET, path, clusterId, "")
	if err != nil {
		return err
	}
	// var response SecretList

	// err = json.Unmarshal([]byte(body), &response)
	// if err != nil {
	// 	return err
	// }
	// if len(response.Items) == 0 {
	// 	return fmt.Errorf("secret not found")
	// }
	return nil
}
func resourceTencentCloudTkeTkeClusterSecretDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_kubernetes_cluster_secret.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		namespace   = d.Get("namespace").(string)
		clusterId   = d.Get("cluster_id").(string)
		secret      = d.Get("secret_name").(string)
		path        = fmt.Sprintf("/api/v1/namespaces/%s/secrets/%s", namespace, secret)
		requestBody = "{\"propagationPolicy\":\"Background\"}"
	)

	// delete secret
	_, err := service.ForwardRequest(ctx, TKE_FORWARD_METHOD_DELETE, path, clusterId, requestBody)
	if err != nil {
		return err
	}

	return nil
}

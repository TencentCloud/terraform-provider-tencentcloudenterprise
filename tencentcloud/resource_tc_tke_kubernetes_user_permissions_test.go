package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestBuildTkeClusterRoleBindingYAML(t *testing.T) {
	body := buildClusterRoleBindingYAML("100000000001", "test-user", "1234567890", "tke:ro")

	var binding struct {
		Metadata struct {
			Annotations map[string]string `json:"annotations"`
		} `json:"metadata"`
		RoleRef struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
		} `json:"roleRef"`
		Subjects []struct {
			Name string `json:"name"`
		} `json:"subjects"`
	}
	if err := json.Unmarshal([]byte(body), &binding); err != nil {
		t.Fatalf("unmarshal ClusterRoleBinding: %v", err)
	}
	if got := binding.Metadata.Annotations["cloud.tencent.com/tke-account-nickname"]; got != "test-user" {
		t.Fatalf("unexpected account nickname: %q", got)
	}
	if binding.RoleRef.Kind != "ClusterRole" || binding.RoleRef.Name != "tke:ro" {
		t.Fatalf("unexpected roleRef: kind=%q name=%q", binding.RoleRef.Kind, binding.RoleRef.Name)
	}
	if len(binding.Subjects) != 1 || binding.Subjects[0].Name != "1234567890" {
		t.Fatalf("unexpected subject: %+v", binding.Subjects)
	}
}

func TestBuildTkeRoleBindingYAML(t *testing.T) {
	body := buildRoleBindingYAML("100000000001", "test-user", "1234567890", "tke:dev", "kube-system")

	var binding struct {
		Metadata struct {
			Name        string            `json:"name"`
			Namespace   string            `json:"namespace"`
			Annotations map[string]string `json:"annotations"`
		} `json:"metadata"`
		RoleRef struct {
			Kind string `json:"kind"`
			Name string `json:"name"`
		} `json:"roleRef"`
		Subjects []struct {
			Name string `json:"name"`
		} `json:"subjects"`
	}
	if err := json.Unmarshal([]byte(body), &binding); err != nil {
		t.Fatalf("unmarshal RoleBinding: %v", err)
	}
	if binding.Metadata.Name != "100000000001-Role-kube-system" {
		t.Fatalf("unexpected binding name: %q", binding.Metadata.Name)
	}
	if binding.Metadata.Namespace != "kube-system" {
		t.Fatalf("unexpected namespace: %q", binding.Metadata.Namespace)
	}
	if got := binding.Metadata.Annotations["cloud.tencent.com/tke-account-nickname"]; got != "test-user" {
		t.Fatalf("unexpected account nickname: %q", got)
	}
	if binding.RoleRef.Kind != "ClusterRole" || binding.RoleRef.Name != "tke:dev" {
		t.Fatalf("unexpected roleRef: kind=%q name=%q", binding.RoleRef.Kind, binding.RoleRef.Name)
	}
	if len(binding.Subjects) != 1 || binding.Subjects[0].Name != "1234567890" {
		t.Fatalf("unexpected subject: %+v", binding.Subjects)
	}
}

func TestValidateTkeUserPermissions(t *testing.T) {
	permissionResource := resourceTencentCloudTkeKubernetesUserPermissions().Schema["permissions"].Elem.(*schema.Resource)
	hash := schema.HashResource(permissionResource)

	tests := []struct {
		name       string
		permission map[string]interface{}
		wantError  string
	}{
		{
			name: "valid cluster permission",
			permission: map[string]interface{}{
				"role_name": "tke:ro",
				"role_type": "cluster",
			},
		},
		{
			name: "valid namespace permission",
			permission: map[string]interface{}{
				"role_name": "tke:dev",
				"role_type": "namespace",
				"namespace": "default",
			},
		},
		{
			name: "missing namespace",
			permission: map[string]interface{}{
				"role_name": "tke:dev",
				"role_type": "namespace",
			},
			wantError: `namespace must be set when role_type is "namespace"`,
		},
		{
			name: "blank namespace",
			permission: map[string]interface{}{
				"role_name": "tke:dev",
				"role_type": "namespace",
				"namespace": "   ",
			},
			wantError: `namespace must be set when role_type is "namespace"`,
		},
		{
			name: "cluster permission with namespace",
			permission: map[string]interface{}{
				"role_name": "tke:ro",
				"role_type": "cluster",
				"namespace": "default",
			},
			wantError: `namespace must not be set when role_type is "cluster"`,
		},
		{
			name: "invalid role type",
			permission: map[string]interface{}{
				"role_name": "tke:ro",
				"role_type": "project",
			},
			wantError: `role_type must be either "cluster" or "namespace"`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			permissions := schema.NewSet(hash, []interface{}{test.permission})
			err := validateTkeUserPermissions(permissions)
			if test.wantError == "" {
				if err != nil {
					t.Fatalf("unexpected validation error: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), test.wantError) {
				t.Fatalf("validation error = %v, want substring %q", err, test.wantError)
			}
		})
	}
}

func TestAccTencentCloudTkeKubernetesUserPermissions_basic(t *testing.T) {
	t.Parallel()

	clusterId := os.Getenv("TF_ACC_TKE_CLUSTER_ID")
	targetUin := os.Getenv("TF_ACC_TKE_TARGET_UIN")
	if clusterId == "" || targetUin == "" {
		t.Skip("set TF_ACC_TKE_CLUSTER_ID and TF_ACC_TKE_TARGET_UIN to run TKE user permissions acceptance test")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTkeKubernetesUserPermissionsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTkeKubernetesUserPermissionsEmpty(clusterId, targetUin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTkeKubernetesUserPermissionsExists("tencentcloudenterprise_tke_kubernetes_user_permissions.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_user_permissions.test", "permissions.#", "0"),
				),
			},
			{
				Config: testAccTkeKubernetesUserPermissionsBasic(clusterId, targetUin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTkeKubernetesUserPermissionsExists("tencentcloudenterprise_tke_kubernetes_user_permissions.test"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tke_kubernetes_user_permissions.test", "target_uin"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_user_permissions.test", "permissions.#", "1"),
					testAccCheckTkeBindingAPI("tencentcloudenterprise_tke_kubernetes_user_permissions.test", "cluster", "", "tke:ro"),
				),
			},
			{
				Config: testAccTkeKubernetesUserPermissionsUpdate(clusterId, targetUin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTkeKubernetesUserPermissionsExists("tencentcloudenterprise_tke_kubernetes_user_permissions.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_user_permissions.test", "permissions.#", "2"),
					testAccCheckTkeBindingAPI("tencentcloudenterprise_tke_kubernetes_user_permissions.test", "namespace", "default", "tke:dev"),
				),
			},
			{
				Config: testAccTkeKubernetesUserPermissionsNamespaceUpdate(clusterId, targetUin),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTkeKubernetesUserPermissionsExists("tencentcloudenterprise_tke_kubernetes_user_permissions.test"),
					resource.TestCheckTypeSetElemNestedAttrs("tencentcloudenterprise_tke_kubernetes_user_permissions.test", "permissions.*", map[string]string{
						"role_name": "tke:dev",
						"role_type": "namespace",
						"namespace": "kube-system",
					}),
					testAccCheckTkeBindingAPI("tencentcloudenterprise_tke_kubernetes_user_permissions.test", "namespace", "kube-system", "tke:dev"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tke_kubernetes_user_permissions.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateId:     fmt.Sprintf("%s:%s", clusterId, targetUin),
			},
		},
	})
}

func testAccTkeKubernetesUserPermissionsEmpty(clusterId, targetUin string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_tke_kubernetes_user_permissions" "test" {
  cluster_id = "%s"
  target_uin = "%s"
}
`, clusterId, targetUin)
}

func testAccCheckTkeBindingAPI(resourceName, roleType, namespace, roleName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}
		idParts := strings.SplitN(rs.Primary.ID, ":", 2)
		if len(idParts) != 2 {
			return fmt.Errorf("invalid resource ID: %s", rs.Primary.ID)
		}
		clusterId, targetUin := idParts[0], idParts[1]

		uin, err := strconv.ParseInt(targetUin, 10, 64)
		if err != nil {
			return fmt.Errorf("parse target UIN: %w", err)
		}

		client := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn
		camService := CamService{client: client}
		expectedNickname, err := camService.DescribeUserNameByUin(context.Background(), uin)
		if err != nil {
			return fmt.Errorf("query CAM user name: %w", err)
		}
		if expectedNickname == "" || expectedNickname == targetUin {
			return fmt.Errorf("CAM did not resolve a user name for UIN %s", targetUin)
		}

		tkeService := TkeService{client: client}
		expectedSubjectName, err := tkeService.DescribeClusterCommonName(context.Background(), clusterId, targetUin)
		if err != nil {
			return fmt.Errorf("query TKE cluster common name: %w", err)
		}
		if expectedSubjectName == "" || expectedSubjectName == targetUin {
			return fmt.Errorf("TKE did not resolve a cluster common name for UIN %s", targetUin)
		}

		var responseBody, bindingName string
		if roleType == "namespace" {
			responseBody, err = tkeService.DescribeRoleBindingsByUin(context.Background(), clusterId, targetUin)
			bindingName = fmt.Sprintf("%s-Role-%s", targetUin, namespace)
		} else {
			responseBody, err = tkeService.DescribeClusterRoleBindingsByUin(context.Background(), clusterId, targetUin)
			bindingName = fmt.Sprintf("%s-ClusterRole", targetUin)
		}
		if err != nil {
			return fmt.Errorf("query TKE bindings: %w", err)
		}

		var list struct {
			Items []struct {
				Metadata struct {
					Name        string            `json:"name"`
					Namespace   string            `json:"namespace"`
					Annotations map[string]string `json:"annotations"`
				} `json:"metadata"`
				RoleRef struct {
					Kind string `json:"kind"`
					Name string `json:"name"`
				} `json:"roleRef"`
				Subjects []struct {
					Name string `json:"name"`
				} `json:"subjects"`
			} `json:"items"`
		}
		if err := json.Unmarshal([]byte(responseBody), &list); err != nil {
			return fmt.Errorf("parse TKE binding response: %w", err)
		}
		for _, binding := range list.Items {
			if binding.Metadata.Name != bindingName {
				continue
			}
			if got := binding.Metadata.Annotations["cloud.tencent.com/tke-account-nickname"]; got != expectedNickname {
				return fmt.Errorf("binding nickname mismatch: got %q, want %q", got, expectedNickname)
			}
			if binding.Metadata.Namespace != namespace {
				return fmt.Errorf("binding namespace mismatch: got %q, want %q", binding.Metadata.Namespace, namespace)
			}
			if binding.RoleRef.Kind != "ClusterRole" || binding.RoleRef.Name != roleName {
				return fmt.Errorf("binding roleRef mismatch: kind=%q name=%q", binding.RoleRef.Kind, binding.RoleRef.Name)
			}
			if len(binding.Subjects) != 1 || binding.Subjects[0].Name != expectedSubjectName {
				return fmt.Errorf("binding subject mismatch: got %+v, want %q", binding.Subjects, expectedSubjectName)
			}
			return nil
		}

		return fmt.Errorf("binding %s not found", bindingName)
	}
}

func testAccCheckTkeKubernetesUserPermissionsExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resource %s not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("resource ID is not set")
		}
		return nil
	}
}

func testAccCheckTkeKubernetesUserPermissionsDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_tke_kubernetes_user_permissions" {
			continue
		}
		return nil
	}
	return nil
}

func testAccTkeKubernetesUserPermissionsBasic(clusterId, targetUin string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_tke_kubernetes_user_permissions" "test" {
  cluster_id = "%s"
  target_uin = "%s"

  permissions {
    role_name = "tke:ro"
    role_type = "cluster"
  }
}
`, clusterId, targetUin)
}

func testAccTkeKubernetesUserPermissionsUpdate(clusterId, targetUin string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_tke_kubernetes_user_permissions" "test" {
  cluster_id = "%s"
  target_uin = "%s"

  permissions {
    role_name = "tke:ro"
    role_type = "cluster"
  }

  permissions {
    role_name = "tke:dev"
    role_type = "namespace"
    namespace = "default"
  }
}
`, clusterId, targetUin)
}

func testAccTkeKubernetesUserPermissionsNamespaceUpdate(clusterId, targetUin string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_tke_kubernetes_user_permissions" "test" {
  cluster_id = "%s"
  target_uin = "%s"

  permissions {
    role_name = "tke:ro"
    role_type = "cluster"
  }

  permissions {
    role_name = "tke:dev"
    role_type = "namespace"
    namespace = "kube-system"
  }
}
`, clusterId, targetUin)
}

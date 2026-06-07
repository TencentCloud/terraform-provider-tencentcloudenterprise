package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const defaultTkeHealthCheckPolicyClusterId = "cls-pwyy5m0s"

func TestAccTencentCloudTkeKubernetesHealthCheckPolicy_basic(t *testing.T) {
	t.Parallel()

	resourceName := "tencentcloudenterprise_tke_kubernetes_health_check_policy.test"
	policyName := fmt.Sprintf("tf-testacc-health-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTkeKubernetesHealthCheckPolicyBasic(defaultTkeHealthCheckPolicyClusterId, policyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_id", defaultTkeHealthCheckPolicyClusterId),
					resource.TestCheckResourceAttr(resourceName, "name", policyName),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),
				),
			},
			{
				Config: testAccTkeKubernetesHealthCheckPolicyUpdate(defaultTkeHealthCheckPolicyClusterId, policyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "cluster_id", defaultTkeHealthCheckPolicyClusterId),
					resource.TestCheckResourceAttr(resourceName, "name", policyName),
					resource.TestCheckResourceAttr(resourceName, "rules.#", "2"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"rules"},
			},
		},
	})
}

func testAccTkeKubernetesHealthCheckPolicyBasic(clusterId, policyName string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_tke_kubernetes_health_check_policy" "test" {
  cluster_id = %q
  name       = %q

  rules {
    name                = "OOMKilling"
    enabled             = true
    auto_repair_enabled = true
  }

  rules {
    name                = "KubeletUnhealthy"
    enabled             = true
    auto_repair_enabled = false
  }
}
`, clusterId, policyName)
}

func testAccTkeKubernetesHealthCheckPolicyUpdate(clusterId, policyName string) string {
	return fmt.Sprintf(`
resource "tencentcloudenterprise_tke_kubernetes_health_check_policy" "test" {
  cluster_id = %q
  name       = %q

  rules {
    name                = "OOMKilling"
    enabled             = true
    auto_repair_enabled = false
  }

  rules {
    name                = "KubeletUnhealthy"
    enabled             = false
    auto_repair_enabled = false
  }
}
`, clusterId, policyName)
}

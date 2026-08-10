package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudTkeKubernetesAddonConfigResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTkeKubernetesAddonConfigCbs,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tke_kubernetes_addon_config.kubernetes_addon_config", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_addon_config.kubernetes_addon_config", "addon_name", "cbs"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tke_kubernetes_addon_config.kubernetes_addon_config", "phase", "Succeeded"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tke_kubernetes_addon_config.kubernetes_addon_config",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"raw_values",
				},
			},
		},
	})
}

const testAccTkeKubernetesAddonConfigCbs = `
locals {
  cluster_id = "cls-3uj3avg6"
}

resource "tencentcloudenterprise_tke_kubernetes_addon_config" "kubernetes_addon_config" {
	cluster_id = local.cluster_id
	addon_name = "cbs"
	raw_values = "{}"
}
`

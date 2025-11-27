package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudKubernetesCommonNamesDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccKubernetesCommonNamesWithSubAccount,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.cloud_tke_kubernetes_cluster_common_names.foo", "cluster_id"),
					//resource.TestCheckResourceAttr("data.cloud_tke_kubernetes_cluster_common_names.foo", "role_ids.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_tke_kubernetes_cluster_common_names.foo", "list.#"),
				),
			},
		},
	})
}

const KeepTkeCNRoleName = `
variable "keep_tke_cn" {
  default = "keep-for-tke-cn"
}
`

const testAccKubernetesCommonNamesBasic = KeepTkeCNRoleName + `
data "tencentcloudenterprise_user_info" "info" {}

locals {
  app_id = data.cloud_user_info.info.app_id
  uin    = data.cloud_user_info.info.uin
}

data "tencentcloudenterprise_tke_kubernetes_clusters" "cls" {
  cluster_name = "` + defaultTkeClusterName + `"
}

data "tencentcloudenterprise_cam_roles" "role_basic" {
  name          = var.keep_tke_cn
}

data "tencentcloudenterprise_tke_kubernetes_cluster_common_names" "foo" {
  cluster_id = data.cloud_tke_kubernetes_clusters.cls.list.0.cluster_id
  role_ids = [data.cloud_cam_roles.role_basic.role_list.0.role_id]
}
`

const testAccKubernetesCommonNamesWithSubAccount = `
data "tencentcloudenterprise_tke_kubernetes_cluster_common_names" "foo" {
	result_output_file = "333.json"
	cluster_id = "cls-9h465v27"
	subaccount_uins = ["100004614032"]
}
`

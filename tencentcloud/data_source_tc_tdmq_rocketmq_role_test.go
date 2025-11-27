package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudTdmqRocketmqRoleDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTdmqRocketmqRole,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tdmq_rocketmq_role.role"),
					resource.TestCheckResourceAttr("data.cloud_tdmq_rocketmq_role.role", "role_sets.#", "1"),
				),
			},
		},
	})
}

const testAccDataSourceTdmqRocketmqRole = `
resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
	cluster_name = "test_rocketmq_datasource_role"
	remark = "test recket mq"
}

resource "tencentcloudenterprise_tdmq_rocketmq_role" "role" {
  role_name = "test_rocketmq_role"
  remark = "test rocketmq role"
  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
}

data "tencentcloudenterprise_tdmq_rocketmq_role" "role" {
  role_name = cloud_tdmq_rocketmq_role.role.role_name
  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
}
`

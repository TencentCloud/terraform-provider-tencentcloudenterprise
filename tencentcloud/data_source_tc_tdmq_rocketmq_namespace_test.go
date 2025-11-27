package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudTdmqRocketmqNamespaceDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTdmqRocketmqNamespace,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tdmq_rocketmq_namespace.namespace"),
					resource.TestCheckResourceAttr("data.cloud_tdmq_rocketmq_namespace.namespace", "namespaces.#", "1"),
				),
			},
		},
	})
}

const testAccDataSourceTdmqRocketmqNamespace = `
resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
	cluster_name = "test_rocketmq_namespace_sdatasource"
	remark = "test recket mq"
}

resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespacedata" {
	cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	namespace_name = "test_namespace_datasource"
	ttl = 65000
	retention_time = 65000
	remark = "test namespace"
}

data "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
	cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	name_keyword = cloud_tdmq_rocketmq_namespace.namespacedata.namespace_name
}
`

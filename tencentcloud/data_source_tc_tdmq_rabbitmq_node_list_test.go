package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixTdmqRabbitmqNodeListDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTdmqRabbitmqNodeListDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_tdmq_rabbitmq_node_list.rabbitmq_node_list"),
				),
			},
		},
	})
}

const testAccTdmqRabbitmqNodeListDataSource = `
data "tencentcloudenterprise_tdmq_rabbitmq_node_list" "specific_node" {
  instance_id = "amqp-xxxxxxxx"
  node_name   = "rabbit@rabbitmq-broker-1.rabbitmq-broker-internal.amqp-7d39mjvn.svc.cluster.local"
}

output "node_status" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_node_list.specific_node.node_list[0].node_status
}
`

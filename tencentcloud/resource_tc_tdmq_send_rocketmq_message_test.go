package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudNeedFixTdmqSendRocketmqMessageResource_basic -v
func TestAccTencentCloudNeedFixTdmqSendRocketmqMessageResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTdmqSendRocketmqMessage,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_send_rocketmq_message.send_rocketmq_message", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_send_rocketmq_message.send_rocketmq_message", "cluster_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_send_rocketmq_message.send_rocketmq_message", "namespace_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_send_rocketmq_message.send_rocketmq_message", "topic_name"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_send_rocketmq_message.send_rocketmq_message", "msg_body"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_send_rocketmq_message.send_rocketmq_message", "msg_key"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_send_rocketmq_message.send_rocketmq_message", "msg_tag"),
				),
			},
		},
	})
}

const testAccTdmqSendRocketmqMessage = `
resource "tencentcloudenterprise_tdmq_send_rocketmq_message" "send_rocketmq_message" {
cluster_id     = "rocketmq-o4e5rnddo9r4"
  namespace_id = "test_namespace"
  topic_name     = "test_rocketmq_topic"
  msg_body     = "msg key"
  msg_key      = "msg tag"
  msg_tag      = "msg value"
}
`

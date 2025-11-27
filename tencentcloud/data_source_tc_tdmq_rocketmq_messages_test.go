package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudNeedFixTdmqRocketmqMessagesDataSource_basic -v
func TestAccTencentCloudNeedFixTdmqRocketmqMessagesDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTdmqMessageDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_tdmq_rocketmq_messages.message"),
				),
			},
		},
	})
}

const testAccTdmqMessageDataSource = `
data "tencentcloudenterprise_tdmq_rocketmq_messages" "message" {
  cluster_id     = "rocketmq-o4e5rnddo9r4"
  environment_id = "test_namespace"
  topic_name     = "test_rocketmq_topic"
  msg_id         = "7F000001001C21B8D17C0366C21D00C2"
  query_dlq_msg  = false
}
`

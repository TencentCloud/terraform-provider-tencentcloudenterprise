package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudClbLogTopic_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbLogTopic_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbLogTopicExists("tencentcloudenterprise_clb_log_topic.topic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_log_topic.topic", "topic_name", "clb-topic-test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_log_topic.topic", "status", "true"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_clb_log_topic.topic", "create_time"),
				),
			},
			{
				Config: testAccClbLogTopic_update,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckClbLogTopicExists("tencentcloudenterprise_clb_log_topic.topic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_log_topic.topic", "topic_name", "clb-topic-test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_clb_log_topic.topic", "status", "false"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_clb_log_topic.topic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckClbLogTopicExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("[CHECK][CLB log topic][Exists] check: CLB log topic %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[CHECK][CLB log topic][Exists] check: CLB log topic id is not set")
		}
		clsService := ClsService{
			client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn,
		}
		instance, err := clsService.DescribeClsTopicById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if instance == nil {
			return fmt.Errorf("[CHECK][CLB log topic][Exists] id %s is not exist", rs.Primary.ID)
		}
		return nil
	}
}

const testAccClbLogTopic_basic = `
resource "tencentcloudenterprise_clb_log_set" "set" {
  period = 7
}

resource "tencentcloudenterprise_clb_log_topic" "topic" {
  log_set_id = tencentcloudenterprise_clb_log_set.set.id
  topic_name = "clb-topic-test"
}
`

const testAccClbLogTopic_update = `
resource "tencentcloudenterprise_clb_log_set" "set" {
  period = 7
}

resource "tencentcloudenterprise_clb_log_topic" "topic" {
  log_set_id = tencentcloudenterprise_clb_log_set.set.id
  topic_name = "clb-topic-test"
  status     = false
}
`

package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbServerAttachmentsDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckClbServerAttachmentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataSourceClbServerAttachments,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckClbServerAttachmentExists("tencentcloudenterprise_clb_attachment.foo"),
					resource.TestCheckResourceAttr("data.cloud_clb_attachments.foo", "attachment_list.#", "1"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_attachments.foo", "attachment_list.0.clb_id"),
					resource.TestCheckResourceAttrSet("data.cloud_clb_attachments.foo", "attachment_list.0.listener_id"),
					resource.TestCheckResourceAttr("data.cloud_clb_attachments.foo", "attachment_list.0.targets.#", "1"),
				),
			},
		},
	})
}

const testAccTencentCloudDataSourceClbServerAttachments = instanceCommonTestCase + `
resource "tencentcloudenterprise_clb_instance" "foo" {
  network_type = "OPEN"
  clb_name     = var.instance_name
  vpc_id       = var.cvm_vpc_id
}

resource "tencentcloudenterprise_clb_listener" "foo" {
  clb_id                     = cloud_clb_instance.foo.id
  listener_name              = var.instance_name
  port                       = 44
  protocol                   = "TCP"
  health_check_switch        = true
  health_check_time_out      = 30
  health_check_interval_time = 100
  health_check_health_num    = 2
  health_check_unhealth_num  = 2
  session_expire_time        = 30
  scheduler                  = "WRR"
}

resource "tencentcloudenterprise_clb_attachment" "foo" {
  clb_id      = cloud_clb_instance.foo.id
  listener_id = cloud_clb_listener.foo.listener_id

  targets {
    instance_id = cloud_cvm_instance.default.id
    port        = 23
    weight      = 10
  }
}

data "tencentcloudenterprise_clb_attachments" "foo" {
  clb_id      = cloud_clb_instance.foo.id
  listener_id = cloud_clb_attachment.foo.listener_id
}
`

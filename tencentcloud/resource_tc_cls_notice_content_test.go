package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClsNoticeContentResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClsNoticeContent,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cls_notice_content.notice_content", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cls_notice_content.notice_content", "type", "0"),
				),
			},
			{
				Config: testAccClsNoticeContentUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_cls_notice_content.notice_content", "name", "terraform-notice-content-for-test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_cls_notice_content.notice_content", "type", "1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cls_notice_content.notice_content",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccClsNoticeContent = `
resource "tencentcloudenterprise_cls_notice_content" "notice_content" {
  name = "terraform-notice-content-test"
  type = 0

  notice_contents {
    type = "Email"

    trigger_content {
      title   = "Trigger title"
      content = "Trigger content"
    }

    recovery_content {
      title   = "Recovery title"
      content = "Recovery content"
    }
  }
}
`

const testAccClsNoticeContentUpdate = `
resource "tencentcloudenterprise_cls_notice_content" "notice_content" {
  name = "terraform-notice-content-for-test"
  type = 1

  notice_contents {
    type = "Email"

    trigger_content {
      title   = "Trigger title updated"
      content = "Trigger content updated"
    }

    recovery_content {
      title   = "Recovery title updated"
      content = "Recovery content updated"
    }
  }
}
`

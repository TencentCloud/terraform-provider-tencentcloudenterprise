package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClsAlarmNoticeResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClsAlarmNotice,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_cls_alarm_notice.alarm_notice", "id")),
			},
			{
				Config: testAccClsAlarmNoticeUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_cls_alarm_notice.alarm_notice", "name", "terraform-alarm-notice-for-test"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_cls_alarm_notice.alarm_notice",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccClsAlarmNotice = `

resource "tencentcloudenterprise_cls_alarm_notice" "alarm_notice" {
  name = "terraform-alarm-notice-test"
  tags = {
    "createdBy" = "terraform"
  }
  type = "All"

  notice_receivers {
    index             = 0
    receiver_channels = [
      "Sms",
    ]
    receiver_ids = [
      13478043,
      15972111,
    ]
    receiver_type = "Uin"
    start_time    = "00:00:00"
    end_time      = "23:59:59"
  }
}

`

const testAccClsAlarmNoticeUpdate = `

resource "tencentcloudenterprise_cls_alarm_notice" "alarm_notice" {
  name = "terraform-alarm-notice-for-test"
  tags = {
    "createdBy" = "terraform"
  }
  type = "All"

  notice_receivers {
    index             = 0
    receiver_channels = [
      "Sms",
    ]
    receiver_ids = [
      13478043,
      15972111,
    ]
    receiver_type = "Uin"
    start_time    = "00:00:00"
    end_time      = "23:59:59"
  }
}

`

package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudAuditTrackResource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAuditTrack,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_audit_track.track", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_audit_track.track", "name", "tf_track_basic"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_audit_track.track", "action_type", "*"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_audit_track.track", "status", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_audit_track.track", "track_for_all_members", "0"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_audit_track.track",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccAuditTrack = `

resource "tencentcloudenterprise_audit_track" "track" {
  action_type           = "*"
  event_names           = [
    "*",
  ]
  name                  = "tf_track_basic"
  resource_type         = "*"
  status                = 1
  track_for_all_members = 0

  storage {
    storage_name   = "tf-bucket-basic"
    storage_prefix = "tftrackbasic"
    storage_region = "kazakhstan-1"
    storage_type   = "cos"
  }
}
`

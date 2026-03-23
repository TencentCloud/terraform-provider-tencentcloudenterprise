package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamLoginRules_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamLoginRulesBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_login_rules.login_rules", "session_duration", "120"),
				),
			},
			{
				Config: testAccCamLoginRulesUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_login_rules.login_rules", "session_duration", "60"),
				),
			},
		},
	})
}

const testAccCamLoginRulesBasic = `
resource "tencentcloudenterprise_cam_login_rules" "login_rules" {
  session_duration = 120
}
`

const testAccCamLoginRulesUpdate = `
resource "tencentcloudenterprise_cam_login_rules" "login_rules" {
  session_duration = 60
}
`

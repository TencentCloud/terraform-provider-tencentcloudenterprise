package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamPasswordRules_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamPasswordRulesBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_cam_password_rules.password_rules", "id", "cam-password-rules"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_cam_password_rules.password_rules", "minimum_length"),
				),
			},
		},
	})
}

const testAccCamPasswordRulesBasic = `
resource "tencentcloudenterprise_cam_password_rules" "password_rules" {
}
`

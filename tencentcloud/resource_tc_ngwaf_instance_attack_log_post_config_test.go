package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNgwafInstanceAttackLogPostConfig_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNgwafInstanceAttackLogPostConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ngwaf_instance_attack_log_post_config.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_instance_attack_log_post_config.example", "instance_id", defaultNgwafInstanceAttackLogPostInstanceId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_instance_attack_log_post_config.example", "attack_log_post", "1"),
				),
			},
			{
				Config: testAccNgwafInstanceAttackLogPostConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ngwaf_instance_attack_log_post_config.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_instance_attack_log_post_config.example", "instance_id", defaultNgwafInstanceAttackLogPostInstanceId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ngwaf_instance_attack_log_post_config.example", "attack_log_post", "0"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_ngwaf_instance_attack_log_post_config.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccNgwafInstanceAttackLogPostConfigBasic = defaultNgwafVariable + `
resource "tencentcloudenterprise_ngwaf_instance_attack_log_post_config" "example" {
  instance_id     = var.ngwaf_instance_attack_log_post_instance_id
  attack_log_post = 1
}
`

const testAccNgwafInstanceAttackLogPostConfigUpdate = defaultNgwafVariable + `
resource "tencentcloudenterprise_ngwaf_instance_attack_log_post_config" "example" {
  instance_id     = var.ngwaf_instance_attack_log_post_instance_id
  attack_log_post = 0
}
`

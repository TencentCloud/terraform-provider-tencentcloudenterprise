package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTsfApplicationFileConfigReleaseResource_basic -v
func TestAccTencentCloudTsfApplicationFileConfigReleaseResource_basic(t *testing.T) {
	// t.Parallel()

	resource.Test(t, resource.TestCase{
		//PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfApplicationFileConfigRelease2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tsf_application_file_config_release.application_file_config_release", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_application_file_config_release.application_file_config_release", "release_desc", "product release"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tsf_application_file_config_release.application_file_config_release",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccTsfApplicationFileConfigReleaseVar = `
variable "config_id" {
	default = "` + defaultTsfFileConfigId + `"
}
`
const testAccTsfApplicationFileConfigRelease = testAccTsfGroup + testAccTsfApplicationFileConfigReleaseVar + `

resource "tencentcloudenterprise_tsf_application_file_config_release" "application_file_config_release" {
  config_id = var.config_id
  group_id = cloud_tsf_group.group.id
  release_desc = "product release"
}

`

const testAccTsfApplicationFileConfigRelease2 = `
resource "tencentcloudenterprise_tsf_application_file_config_release" "application_file_config_release" {
  config_id = "dcfg-f-qevjmlvb"
  group_id = "group-byxmmovl"
  release_desc = "product release"
}
`

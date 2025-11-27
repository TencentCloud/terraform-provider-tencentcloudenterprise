package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTsfApplicationPublicConfigReleaseResource_basic -v
func TestAccTencentCloudTsfApplicationPublicConfigReleaseResource_basic(t *testing.T) {
	// t.Parallel()

	resource.Test(t, resource.TestCase{
		//PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfApplicationPublicConfigRelease2,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tsf_application_public_config_release.application_public_config_release", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_tsf_application_public_config_release.application_public_config_release", "release_desc", "v1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tsf_application_public_config_release.application_public_config_release",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccTsfApplicationPublicConfigRelease = testAccTsfApplicationPublicConfig + testAccTsfNamespace + `

resource "tencentcloudenterprise_tsf_application_public_config_release" "application_public_config_release" {
  config_id = cloud_tsf_application_public_config.application_public_config.id
  namespace_id = cloud_tsf_namespace.namespace.id
  release_desc = "v1"
}

`

const testAccTsfApplicationPublicConfigRelease2 = `
resource "tencentcloudenterprise_tsf_application_public_config_release" "application_public_config_release" {
  config_id = "dcfg-p-ldap93vo"
  namespace_id = "namespace-jqv3pga7"
  release_desc = "v1"
}	
`

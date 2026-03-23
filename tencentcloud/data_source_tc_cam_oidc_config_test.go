package tencentcloud

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// 测试查询 OIDC 配置
// 需要在环境中预先配置一个 OIDC IdP
func TestAccTencentCloudCamOidcConfigDataSource_basic(t *testing.T) {
	t.Parallel()

	testIdpName := os.Getenv("TENCENTCLOUD_TEST_OIDC_IDP_NAME")
	if testIdpName == "" {
		t.Skip("TENCENTCLOUD_TEST_OIDC_IDP_NAME is not set, skipping test. Please create an OIDC IdP first and set the environment variable.")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamOidcConfigDataSource(testIdpName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_cam_oidc_config.test"),
					resource.TestCheckResourceAttr("data.tencentcloudenterprise_cam_oidc_config.test", "name", testIdpName),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_oidc_config.test", "provider_type"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_oidc_config.test", "identity_url"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_oidc_config.test", "identity_key"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_oidc_config.test", "client_id.#"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_oidc_config.test", "status"),
				),
			},
		},
	})
}

func testAccCamOidcConfigDataSource(idpName string) string {
	return `
data "tencentcloudenterprise_cam_oidc_config" "test" {
  name = "` + idpName + `"
}
`
}

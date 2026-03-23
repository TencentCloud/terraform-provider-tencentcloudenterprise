package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudUserInfoDataSource_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUserInfoDataSource,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.tencentcloudenterprise_user_info.current"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_user_info.current", "app_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_user_info.current", "uin"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_user_info.current", "owner_uin"),
				),
			},
		},
	})
}

const testAccUserInfoDataSource = `
data "tencentcloudenterprise_user_info" "current" {
}
`

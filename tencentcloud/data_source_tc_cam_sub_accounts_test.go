package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamSubAccountsDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamSubAccountsDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_sub_accounts.foo", "sub_accounts.0.uin"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_sub_accounts.foo", "sub_accounts.0.name"),
				),
			},
		},
	})
}

const testAccCamSubAccountsDataSource_basic = `
data "tencentcloudenterprise_cam_users" "users" {}

data "tencentcloudenterprise_cam_sub_accounts" "foo" {
  filter_sub_account_uin = [data.tencentcloudenterprise_cam_users.users.user_list.0.uin]
}
`

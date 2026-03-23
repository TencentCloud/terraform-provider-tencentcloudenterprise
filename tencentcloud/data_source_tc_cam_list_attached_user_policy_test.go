package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamListAttachedUserPolicyDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamListAttachedUserPolicyDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_list_attached_user_policy.foo", "id"),
				),
			},
		},
	})
}

const testAccCamListAttachedUserPolicyDataSource_basic = `
data "tencentcloudenterprise_cam_users" "users" {}

data "tencentcloudenterprise_cam_list_attached_user_policy" "foo" {
  target_uin  = data.tencentcloudenterprise_cam_users.users.user_list.0.uin
  attach_type = 0
}
`

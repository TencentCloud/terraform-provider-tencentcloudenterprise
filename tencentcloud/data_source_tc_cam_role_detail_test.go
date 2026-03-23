package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamRoleDetailDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamRoleDetailDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_role_detail.foo", "role_info.0.role_id"),
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_role_detail.foo", "role_info.0.role_name"),
				),
			},
		},
	})
}

const testAccCamRoleDetailDataSource_basic = `
data "tencentcloudenterprise_cam_roles" "roles" {}

data "tencentcloudenterprise_cam_role_detail" "foo" {
  role_id = data.tencentcloudenterprise_cam_roles.roles.role_list.0.role_id
}
`

package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCamListEntitiesForPolicyDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCamListEntitiesForPolicyDataSource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.tencentcloudenterprise_cam_list_entities_for_policy.foo", "list.0.uin"),
				),
			},
		},
	})
}

const testAccCamListEntitiesForPolicyDataSource_basic = `
data "tencentcloudenterprise_cam_policies" "policies" {
  type = 2 # Preset policies usually have entities
}

data "tencentcloudenterprise_cam_list_entities_for_policy" "foo" {
  policy_id = data.tencentcloudenterprise_cam_policies.policies.policy_list.0.policy_id
}
`

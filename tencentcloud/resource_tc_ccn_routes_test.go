package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCcnRoutesResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcCcnRoutes,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_ccn_routes.ccn_routes", "id")),
			},
			{
				Config: testAccVpcCcnRoutesUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_ccn_routes.ccn_routes", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_ccn_routes.ccn_routes", "switch", "on"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_ccn_routes.ccn_routes",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcCcnRoutes = `

resource "tencentcloudenterprise_ccn_routes" "ccn_routes" {
  ccn_id = "ccn-39lqkygf"
  route_id = "ccnr-3o0dfyuw"
  switch = "off"
}

`

const testAccVpcCcnRoutesUpdate = `

resource "tencentcloudenterprise_ccn_routes" "ccn_routes" {
  ccn_id = "ccn-39lqkygf"
  route_id = "ccnr-3o0dfyuw"
  switch = "on"
}

`

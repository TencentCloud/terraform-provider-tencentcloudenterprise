package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixVpnGatewayCcnRoutesResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcVpnGatewayCcnRoutes,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_vpn_gateway_ccn_routes.vpn_gateway_ccn_routes", "id")),
			},
			{
				ResourceName:      "cloud_vpn_gateway_ccn_routes.vpn_gateway_ccn_routes",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcVpnGatewayCcnRoutes = `

resource "cloud_vpn_gateway_ccn_routes" "vpn_gateway_ccn_routes" {
  destination_cidr_block = "192.168.1.0/24"
  route_id               = "vpnr-akdy0757"
  status                 = "DISABLE"
  vpn_gateway_id         = "vpngw-lie1a4u7"
}

`

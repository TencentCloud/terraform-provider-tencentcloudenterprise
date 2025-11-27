package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcIpv6EniAddressResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcIpv6EniAddress,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_ipv6_eni_address.ipv6_eni_address", "id")),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_ipv6_eni_address.ipv6_eni_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcIpv6EniAddress = `

resource "tencentcloudenterprise_vpc_ipv6_eni_address" "ipv6_eni_address" {
  vpc_id = ""
  network_interface_id = ""
  ipv6_addresses {
		address = ""
		primary = False
		address_id = ""
		description = ""
		is_wan_ip_blocked = 
		state = ""

  }
}

`

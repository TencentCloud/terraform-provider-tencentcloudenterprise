package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixVpcDhcpAssociateAddressResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcDhcpAssociateAddress,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_dhcp_associate_address.dhcp_associate_address", "id")),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_dhcp_associate_address.dhcp_associate_address",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcDhcpAssociateAddress = `

resource "tencentcloudenterprise_vpc_dhcp_associate_address" "dhcp_associate_address" {
  dhcp_ip_id = ""
  address_ip = ""
}

`

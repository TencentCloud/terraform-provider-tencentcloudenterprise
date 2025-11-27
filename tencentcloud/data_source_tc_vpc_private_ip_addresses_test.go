package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcPrivateIpAddressesDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcPrivateIpAddressesDataSource,
				Check:  resource.ComposeTestCheckFunc(testAccCheckTencentCloudDataSourceID("data.cloud_vpc_private_ip_addresses.private_ip_addresses")),
			},
		},
	})
}

const testAccVpcPrivateIpAddressesDataSource = `

data "tencentcloudenterprise_vpc_private_ip_addresses" "private_ip_addresses" {
  vpc_id = "vpc-l0dw94uh"
  private_ip_addresses = ["203.0.113.1"]
}

`

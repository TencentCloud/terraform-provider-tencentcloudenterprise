package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpnGatewaysDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudVpnGatewaysDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_vpn_gateways.cgws"),
					resource.TestCheckResourceAttr("data.cloud_vpn_gateways.cgws", "gateway_list.#", "1"),
					resource.TestCheckResourceAttr("data.cloud_vpn_gateways.cgws", "gateway_list.0.name", "terraform_test"),
					resource.TestCheckResourceAttr("data.cloud_vpn_gateways.cgws", "gateway_list.0.bandwidth", "10"),
					resource.TestCheckResourceAttr("data.cloud_vpn_gateways.cgws", "gateway_list.0.tags.test", "tf"),
				),
			},
		},
	})
}

const testAccTencentCloudVpnGatewaysDataSourceConfig_basic = `
# Create VPC
data "tencentcloudenterprise_vpc_instances" "foo" {
  name = "Default-VPC"
}

resource "tencentcloudenterprise_vpn_gateway" "my_cgw" {
  name      = "terraform_test"
  vpc_id    = data.cloud_vpc_instances.foo.instance_list.0.vpc_id
  bandwidth = 10
  zone      = "ap-guangzhou-3"

  tags = {
    test = "tf"
  }
}

data "tencentcloudenterprise_vpn_gateways" "cgws" {
  id = cloud_vpn_gateway.my_cgw.id
}
`

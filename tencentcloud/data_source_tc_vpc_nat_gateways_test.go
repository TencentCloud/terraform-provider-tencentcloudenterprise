package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNatGatewaysDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudNatGatewaysDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_nat_gateways.multi_nat"),
					resource.TestCheckResourceAttr("data.cloud_vpc_nat_gateways.multi_nat", "nats.#", "2"),
					resource.TestCheckResourceAttr("data.cloud_vpc_nat_gateways.multi_nat", "nats.0.name", "terraform_test_nats"),
					resource.TestCheckResourceAttr("data.cloud_vpc_nat_gateways.multi_nat", "nats.1.bandwidth", "500"),
					//resource.TestCheckResourceAttr("data.cloud_vpc_nat_gateways.multi_nat", "nats.0.tags.tf", "test"),
				),
			},
		},
	})
}

const testAccTencentCloudNatGatewaysDataSourceConfig_basic = `
resource "tencentcloudenterprise_vpc" "main" {
  name       = "terraform_test_nats"
  cidr_block = "10.6.0.0/16"
}
resource "tencentcloudenterprise_eip" "eip_dev_dnat" {
  name = "terraform_test"
}
resource "tencentcloudenterprise_eip" "eip_test_dnat" {
  name = "terraform_test"
}

resource "tencentcloudenterprise_vpc_nat_gateway" "dev_nat" {
  vpc_id           = cloud_vpc.main.id
  name             = "terraform_test_nats"
  max_concurrent   = 3000000
  bandwidth        = 500
  assigned_eip_set = [
    cloud_eip.eip_dev_dnat.public_ip,
  ]
}
resource "tencentcloudenterprise_vpc_nat_gateway" "test_nat" {
  vpc_id           = cloud_vpc.main.id
  name             = "terraform_test_nats"
  max_concurrent   = 3000000
  bandwidth        = 500
  assigned_eip_set = [
    cloud_eip.eip_test_dnat.public_ip,
  ]
}

data "tencentcloudenterprise_vpc_nat_gateways" "multi_nat" {
  name           = cloud_vpc_nat_gateway.dev_nat.name
  vpc_id         = cloud_vpc.main.id
}
`

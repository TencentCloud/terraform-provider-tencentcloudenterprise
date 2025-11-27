package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNatGatewayTransRuleDataSource(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTencentCloudDataSourceDnatsBase,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_dnats.multi_dnats"),
					resource.TestCheckResourceAttr("data.cloud_vpc_dnats.multi_dnats", "dnat_list.#", "1"),
					resource.TestCheckResourceAttr("data.cloud_vpc_dnats.multi_dnats", "dnat_list.0.description", defaultInsName),
				),
			},
		},
	})
}

const testAccTencentCloudDataSourceDnatsBase = instanceCommonTestCase + `
# Create EIP 
resource "tencentcloudenterprise_eip" "eip_dev_dnat" {
  name = var.instance_name
}

resource "tencentcloudenterprise_eip" "eip_test_dnat" {
  name = var.instance_name
}

# Create NAT Gateway
resource "tencentcloudenterprise_vpc_nat_gateway" "my_nat" {
  vpc_id         = var.cvm_vpc_id
  name           = var.instance_name
  max_concurrent = 3000000
  bandwidth      = 500

  assigned_eip_set = [
    cloud_eip.eip_dev_dnat.public_ip,
    cloud_eip.eip_test_dnat.public_ip,
  ]
}

# Add DNAT Entry
resource "tencentcloudenterprise_vpc_dnat" "dev_dnat" {
  vpc_id       = cloud_vpc_nat_gateway.my_nat.vpc_id
  nat_id       = cloud_vpc_nat_gateway.my_nat.id
  protocol     = "TCP"
  elastic_ip   = cloud_eip.eip_dev_dnat.public_ip
  elastic_port = "80"
  private_ip   = cloud_cvm_instance.default.private_ip
  private_port = "9001"
  description  = var.instance_name
}

data "tencentcloudenterprise_vpc_dnats" "multi_dnats" {
  nat_id = cloud_vpc_dnat.dev_dnat.nat_id
  vpc_id = cloud_vpc_dnat.dev_dnat.vpc_id
}
`

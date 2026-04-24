package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTencentCloudDcGatewayInstancesBasic(t *testing.T) {
	t.Parallel()

	var nameKey = "data.tencentcloudenterprise_dc_gateway_instances.name_select"
	var idKey = "data.tencentcloudenterprise_dc_gateway_instances.id_select"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTencentCloudDcGatewayInstancesConfig,
				Check: resource.ComposeTestCheckFunc(
					// name filter
					testAccCheckTencentCloudDataSourceID(nameKey),
					resource.TestCheckResourceAttrSet(nameKey, "instance_list.#"),

					// id filter
					testAccCheckTencentCloudDataSourceID(idKey),
					resource.TestCheckResourceAttr(idKey, "instance_list.#", "1"),

					resource.TestCheckResourceAttrSet(idKey, "instance_list.0.dcg_id"),
					resource.TestCheckResourceAttrSet(idKey, "instance_list.0.name"),
					resource.TestCheckResourceAttrSet(idKey, "instance_list.0.network_type"),
					resource.TestCheckResourceAttrSet(idKey, "instance_list.0.gateway_type"),
					resource.TestCheckResourceAttrSet(idKey, "instance_list.0.create_time"),
				),
			},
		},
	})
}

// VPC-type dc_gateway is simpler and does not require CCN
const testAccDataSourceTencentCloudDcGatewayInstancesConfig = `
resource "tencentcloudenterprise_vpc" "main" {
  name       = "ci-dcg-test-vpc"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_dc_gateway" "main" {
  name                = "ci-dcg-test"
  network_instance_id = tencentcloudenterprise_vpc.main.id
  network_type        = "VPC"
  gateway_type        = "NAT"
}

data "tencentcloudenterprise_dc_gateway_instances" "name_select" {
  name = tencentcloudenterprise_vpc_dc_gateway.main.name
}

data "tencentcloudenterprise_dc_gateway_instances" "id_select" {
  dcg_id = tencentcloudenterprise_vpc_dc_gateway.main.id
}
`

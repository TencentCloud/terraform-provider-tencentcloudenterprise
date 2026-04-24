package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTencentCloudDcGatewayCcnRoutesBasic(t *testing.T) {
	t.Parallel()

	var rKey = "data.tencentcloudenterprise_dc_gateway_ccn_routes.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceTencentCloudDcGatewayCcnRoutesConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID(rKey),
					// instance_list may be empty for a VPC-type gateway without CCN routes,
					// just verify the data source can be read without error
					resource.TestCheckResourceAttrSet(rKey, "instance_list.#"),
				),
			},
		},
	})
}

// Use VPC-type dc_gateway since CCN creation is not available in TCE
const testAccDataSourceTencentCloudDcGatewayCcnRoutesConfig = `
resource "tencentcloudenterprise_vpc" "main" {
  name       = "ci-dcg-route-test-vpc"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_dc_gateway" "main" {
  name                = "ci-dcg-route-test"
  network_instance_id = tencentcloudenterprise_vpc.main.id
  network_type        = "VPC"
  gateway_type        = "NAT"
}

data "tencentcloudenterprise_dc_gateway_ccn_routes" "test" {
  dcg_id = tencentcloudenterprise_vpc_dc_gateway.main.id
}
`

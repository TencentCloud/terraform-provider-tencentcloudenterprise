package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudVpcNatGatewayFlowMonitorResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcNatGatewayFlowMonitor_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_nat_gateway_flow_monitor.example", "id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_nat_gateway_flow_monitor.example", "enable", "true"),
				),
			},
			{
				Config: testAccVpcNatGatewayFlowMonitor_disable,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_nat_gateway_flow_monitor.example", "enable", "false"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_nat_gateway_flow_monitor.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// testAccVpcNatGatewayFlowMonitorBase creates a NAT gateway with required EIPs as a dependency.
const testAccVpcNatGatewayFlowMonitorBase = `
data "tencentcloudenterprise_vpc_instances" "foo" {
  name = "Default-VPC"
}

resource "tencentcloudenterprise_eip" "eip1" {
  name = "tf-test-nat-flow-monitor-eip1"
}

resource "tencentcloudenterprise_vpc_nat_gateway" "nat" {
  vpc_id           = data.tencentcloudenterprise_vpc_instances.foo.instance_list.0.vpc_id
  name             = "tf-test-nat-flow-monitor"
  max_concurrent   = 1000000
  bandwidth        = 100
  assigned_eip_set = [tencentcloudenterprise_eip.eip1.public_ip]
}
`

const testAccVpcNatGatewayFlowMonitor_basic = testAccVpcNatGatewayFlowMonitorBase + `
resource "tencentcloudenterprise_vpc_nat_gateway_flow_monitor" "example" {
  gateway_id = tencentcloudenterprise_vpc_nat_gateway.nat.id
  enable     = true
}
`

const testAccVpcNatGatewayFlowMonitor_disable = testAccVpcNatGatewayFlowMonitorBase + `
resource "tencentcloudenterprise_vpc_nat_gateway_flow_monitor" "example" {
  gateway_id = tencentcloudenterprise_vpc_nat_gateway.nat.id
  enable     = false
}
`

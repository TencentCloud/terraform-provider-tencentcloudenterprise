package tencentcloud

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestBuildVpcNatGatewaySnat_subnet(t *testing.T) {
	d := resourceTencentCloudVpcNatGatewaySnat().TestResourceData()
	_ = d.Set("resource_type", NAT_GATEWAY_TYPE_SUBNET)
	_ = d.Set("subnet_id", "subnet-abc123")
	_ = d.Set("subnet_cidr_block", "172.31.1.0/24")
	_ = d.Set("description", "unit test subnet snat")
	_ = d.Set("public_ip_addr", []interface{}{"1.2.3.4", "5.6.7.8"})

	rule := buildVpcNatGatewaySnat(d)
	if rule == nil {
		t.Fatal("expected SNAT rule")
	}
	if *rule.ResourceType != NAT_GATEWAY_TYPE_SUBNET {
		t.Fatalf("resource type = %s", *rule.ResourceType)
	}
	if *rule.ResourceId != "subnet-abc123" {
		t.Fatalf("resource id = %s", *rule.ResourceId)
	}
	if *rule.PrivateIpAddress != "172.31.1.0/24" {
		t.Fatalf("private ip = %s", *rule.PrivateIpAddress)
	}
	if len(rule.PublicIpAddresses) != 2 {
		t.Fatalf("public ip count = %d", len(rule.PublicIpAddresses))
	}
}

func TestAccTencentCloudVpcNatGatewaySnat_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcNatGatewaySnatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcNatGatewaySnatConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcNatGatewaySnatExists("tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat", "resource_type", "SUBNET"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat", "description", "tf acc snat"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat", "public_ip_addr.#", "2"),
				),
			},
			{
				Config: testAccVpcNatGatewaySnatConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcNatGatewaySnatExists("tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat", "description", "tf acc snat updated"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat", "public_ip_addr.#", "1"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckVpcNatGatewaySnatDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_vpc_nat_gateway_snat" {
			continue
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("nat gateway snat id is not set")
		}
		ids := strings.Split(rs.Primary.ID, FILED_SP)
		if len(ids) != 2 {
			return fmt.Errorf("unexpected snat id %s", rs.Primary.ID)
		}
		err, result := service.DescribeNatGatewaySnats(contextNil, ids[0], nil)
		if err != nil {
			log.Printf("[CRITAL]%s read nat gateway snat failed, reason:%s\n", logId, err.Error())
			return err
		}
		for _, item := range result {
			if item != nil && item.ResourceId != nil && *item.ResourceId == ids[1] {
				return fmt.Errorf("nat gateway snat %s still exists", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckVpcNatGatewaySnatExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("nat gateway snat instance %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("nat gateway snat id is not set")
		}
		ids := strings.Split(rs.Primary.ID, FILED_SP)
		if len(ids) != 2 {
			return fmt.Errorf("unexpected snat id %s", rs.Primary.ID)
		}
		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		err, result := service.DescribeNatGatewaySnats(contextNil, ids[0], nil)
		if err != nil {
			return err
		}
		for _, item := range result {
			if item != nil && item.ResourceId != nil && *item.ResourceId == ids[1] {
				return nil
			}
		}
		return fmt.Errorf("nat gateway snat %s is not found", rs.Primary.ID)
	}
}

const testAccVpcNatGatewaySnatNetwork = `
data "tencentcloudenterprise_vpc_instances" "default" {}

data "tencentcloudenterprise_vpc_subnets" "default" {
  vpc_id = data.tencentcloudenterprise_vpc_instances.default.instance_list.0.vpc_id
}

resource "tencentcloudenterprise_vpc" "vpc" {
  name       = "tf-acc-vpc-snat"
  cidr_block = "172.31.0.0/16"
}

resource "tencentcloudenterprise_vpc_route_table" "snat" {
  vpc_id = tencentcloudenterprise_vpc.vpc.id
  name   = "tf-acc-vpc-snat"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  vpc_id            = tencentcloudenterprise_vpc.vpc.id
  name              = "tf-acc-vpc-snat"
  cidr_block        = "172.31.1.0/24"
  availability_zone = data.tencentcloudenterprise_vpc_subnets.default.instance_list.0.availability_zone
  route_table_id    = tencentcloudenterprise_vpc_route_table.snat.id
}

resource "tencentcloudenterprise_eip" "eip1" {
  name = "tf-acc-vpc-snat-1"
}

resource "tencentcloudenterprise_eip" "eip2" {
  name = "tf-acc-vpc-snat-2"
}

resource "tencentcloudenterprise_vpc_nat_gateway" "nat" {
  vpc_id         = tencentcloudenterprise_vpc.vpc.id
  name           = "tf-acc-vpc-snat"
  max_concurrent = 3000000
  bandwidth      = 500
  assigned_eip_set = [
    tencentcloudenterprise_eip.eip1.public_ip,
    tencentcloudenterprise_eip.eip2.public_ip,
  ]
}

resource "tencentcloudenterprise_vpc_route_table_entry" "nat" {
  route_table_id         = tencentcloudenterprise_vpc_route_table.snat.id
  destination_cidr_block = "0.0.0.0/0"
  next_type              = "NAT"
  next_hop               = tencentcloudenterprise_vpc_nat_gateway.nat.id
  description            = "tf acc snat default route"
}
`

const testAccVpcNatGatewaySnatConfig = testAccVpcNatGatewaySnatNetwork + `
resource "tencentcloudenterprise_vpc_nat_gateway_snat" "subnet_snat" {
  nat_gateway_id    = tencentcloudenterprise_vpc_nat_gateway.nat.id
  resource_type     = "SUBNET"
  subnet_id         = tencentcloudenterprise_vpc_subnet.subnet.id
  subnet_cidr_block = tencentcloudenterprise_vpc_subnet.subnet.cidr_block
  description       = "tf acc snat"
  public_ip_addr = [
    tencentcloudenterprise_eip.eip1.public_ip,
    tencentcloudenterprise_eip.eip2.public_ip,
  ]
  depends_on = [tencentcloudenterprise_vpc_route_table_entry.nat]
}
`

const testAccVpcNatGatewaySnatConfigUpdate = testAccVpcNatGatewaySnatNetwork + `
resource "tencentcloudenterprise_vpc_nat_gateway_snat" "subnet_snat" {
  nat_gateway_id    = tencentcloudenterprise_vpc_nat_gateway.nat.id
  resource_type     = "SUBNET"
  subnet_id         = tencentcloudenterprise_vpc_subnet.subnet.id
  subnet_cidr_block = tencentcloudenterprise_vpc_subnet.subnet.cidr_block
  description       = "tf acc snat updated"
  public_ip_addr = [
    tencentcloudenterprise_eip.eip1.public_ip,
  ]
  depends_on = [tencentcloudenterprise_vpc_route_table_entry.nat]
}
`

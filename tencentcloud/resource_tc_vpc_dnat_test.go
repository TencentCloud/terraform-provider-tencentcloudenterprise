package tencentcloud

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func TestAccTencentCloudNatGatewayTransRule_basic(t *testing.T) {
	t.Parallel()
	var dnatId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDnatDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccDnatConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnatExists("cloud_vpc_dnat.dev_dnat", &dnatId),
				),
			},
			{
				Config: testAccDnatConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDnatExists("cloud_vpc_dnat.dev_dnat", &dnatId),
				),
			},
		},
	})
}

func testAccCheckDnatExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("DNAT instance %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("DNAT id is not set")
		}
		_, params, e := parseDnatId(rs.Primary.ID)
		if e != nil {
			return fmt.Errorf("[CRITAL]parse DNAT id fail, reason[%s]\n", e.Error())
		}
		conn := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn
		request := vpc.NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest()
		request.Filters = make([]*vpc.Filter, 0, len(params))
		for k, v := range params {
			filter := &vpc.Filter{
				Name:   helper.String(k),
				Values: []*string{helper.String(v)},
			}
			request.Filters = append(request.Filters, filter)
		}
		var response *vpc.DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := conn.UseVpcClient().DescribeNatGatewayDestinationIpPortTranslationNatRules(request)
			if e != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, request.GetAction(), request.ToJsonString(), e.Error())
				return retryError(e)
			}
			response = result
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s read DNAT failed, reason:%s\n", logId, err.Error())
			return err
		}
		if len(response.Response.NatGatewayDestinationIpPortTranslationNatRuleSet) != 1 {
			return fmt.Errorf("DNAT is not exists")
		}
		return nil
	}
}

func testAccCheckDnatDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloud_vpc_dnat" {
			continue
		}
		_, params, e := parseDnatId(rs.Primary.ID)
		if e != nil {
			log.Printf("[CRITAL]parse DNAT id fail, reason[%s]\n", e.Error())
		}
		conn := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn
		request := vpc.NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest()
		request.Filters = make([]*vpc.Filter, 0, len(params))
		for k, v := range params {
			filter := &vpc.Filter{
				Name:   helper.String(k),
				Values: []*string{helper.String(v)},
			}
			request.Filters = append(request.Filters, filter)
		}
		var response *vpc.DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := conn.UseVpcClient().DescribeNatGatewayDestinationIpPortTranslationNatRules(request)
			if e != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, request.GetAction(), request.ToJsonString(), e.Error())
				return retryError(e)
			}
			response = result
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s read DNAT failed, reason:%s\n", logId, err.Error())
			return err
		}
		if len(response.Response.NatGatewayDestinationIpPortTranslationNatRuleSet) != 0 {
			return fmt.Errorf("DNAT is still exists")
		}
	}
	return nil

}

const testAccDnatConfig = instanceCommonTestCase + `
# Create EIP 
resource "cloud_eip" "eip_dev_dnat" {
  name = var.instance_name
}

resource "cloud_eip" "eip_test_dnat" {
  name = var.instance_name
}

# Create NAT Gateway
resource "cloud_vpc_nat_gateway" "my_nat" {
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
resource "cloud_vpc_dnat" "dev_dnat" {
  vpc_id       = cloud_vpc_nat_gateway.my_nat.vpc_id
  nat_id       = cloud_vpc_nat_gateway.my_nat.id
  protocol     = "TCP"
  elastic_ip   = cloud_eip.eip_dev_dnat.public_ip
  elastic_port = "80"
  private_ip   = cloud_cvm_instance.default.private_ip
  private_port = "9001"
}
`

const testAccDnatConfigUpdate = instanceCommonTestCase + `
# Create EIP 
resource "cloud_eip" "eip_dev_dnat" {
  name = var.instance_name
}

resource "cloud_eip" "eip_test_dnat" {
  name = var.instance_name
}

# Create NAT Gateway
resource "cloud_vpc_nat_gateway" "my_nat" {
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
resource "cloud_vpc_dnat" "dev_dnat" {
  vpc_id       = cloud_vpc_nat_gateway.my_nat.vpc_id
  nat_id       = cloud_vpc_nat_gateway.my_nat.id
  protocol     = "TCP"
  elastic_ip   = cloud_eip.eip_dev_dnat.public_ip
  elastic_port = "80"
  private_ip   = cloud_cvm_instance.default.private_ip
  private_port = "9001"
  description  = var.instance_name
}
`

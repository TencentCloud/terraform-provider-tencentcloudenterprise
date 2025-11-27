package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudSecurityGroupLiteRule_basic(t *testing.T) {
	t.Parallel()

	var liteRuleId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityGroupLiteRuleDestroy(&liteRuleId),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupLiteRuleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupLiteRuleExists("tencentcloudenterprise_vpc_security_group_lite_rule.foo", &liteRuleId),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.0", "ACCEPT#192.168.1.0/24#80#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.1", "DROP#8.8.8.8#80,90#UDP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.2", "ACCEPT#0.0.0.0/0#80-90#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.0", "ACCEPT#192.168.0.0/16#ALL#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.1", "ACCEPT#10.0.0.0/8#ALL#ICMP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.2", "DROP#0.0.0.0/0#ALL#ALL"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_security_group_lite_rule.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudSecurityGroupLiteRule_update(t *testing.T) {
	t.Parallel()

	var liteRuleId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityGroupLiteRuleDestroy(&liteRuleId),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupLiteRuleBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupLiteRuleExists("tencentcloudenterprise_vpc_security_group_lite_rule.foo", &liteRuleId),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.#", "3"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.0", "ACCEPT#192.168.1.0/24#80#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.1", "DROP#8.8.8.8#80,90#UDP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.2", "ACCEPT#0.0.0.0/0#80-90#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.#", "3"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.0", "ACCEPT#192.168.0.0/16#ALL#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.1", "ACCEPT#10.0.0.0/8#ALL#ICMP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.2", "DROP#0.0.0.0/0#ALL#ALL"),
				),
			},
			{
				Config: testAccSecurityGroupLiteRuleUpdate1,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupLiteRuleExists("tencentcloudenterprise_vpc_security_group_lite_rule.foo", &liteRuleId),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.#", "4"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.0", "ACCEPT#192.168.1.0/24#80#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.1", "ACCEPT#192.168.1.0/26#800#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.2", "DROP#8.8.8.8#80,90#UDP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.3", "ACCEPT#0.0.0.0/0#80-90#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.#", "4"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.0", "ACCEPT#192.168.0.0/16#ALL#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.1", "ACCEPT#192.168.0.0/24#ALL#UDP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.2", "ACCEPT#10.0.0.0/8#ALL#ICMP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.3", "DROP#0.0.0.0/0#ALL#ALL"),
				),
			},
			{
				Config: testAccSecurityGroupLiteRuleUpdate2,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupLiteRuleExists("tencentcloudenterprise_vpc_security_group_lite_rule.foo", &liteRuleId),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.#", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.0", "ACCEPT#192.168.1.0/24#80#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.#", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.0", "ACCEPT#192.168.0.0/16#ALL#TCP"),
				),
			},
			{
				Config: testAccSecurityGroupLiteRuleUpdate3,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupLiteRuleExists("tencentcloudenterprise_vpc_security_group_lite_rule.foo", &liteRuleId),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.#", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.0", "ACCEPT#192.168.1.0/24#80#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.#", "0"),
				),
			},
			{
				Config: testAccSecurityGroupLiteRuleUpdate4,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupLiteRuleExists("tencentcloudenterprise_vpc_security_group_lite_rule.foo", &liteRuleId),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.#", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.#", "1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.0", "ACCEPT#192.168.0.0/16#ALL#TCP"),
				),
			},
			{
				Config: testAccSecurityGroupLiteRuleUpdate5,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupLiteRuleExists("tencentcloudenterprise_vpc_security_group_lite_rule.foo", &liteRuleId),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.#", "0"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.#", "0"),
				),
			},
			{
				Config: testAccSecurityGroupLiteRuleUpdate6,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupLiteRuleExists("tencentcloudenterprise_vpc_security_group_lite_rule.foo", &liteRuleId),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "security_group_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.#", "5"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.0"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.1"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.2", "ACCEPT#0.0.0.0/0#80-90#TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.3", "DROP#8.8.8.8#80,90#UDP"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "ingress.4"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.#", "2"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.0", "ACCEPT#192.168.0.0/16#ALL#TCP"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_lite_rule.foo", "egress.1"),
				),
			},
		},
	})
}

func testAccCheckSecurityGroupLiteRuleExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no security group rule ID is set")
		}

		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

		_, _, exist, err := service.DescribeSecurityGroupPolices(context.TODO(), rs.Primary.ID)
		if err != nil {
			return err
		}

		if !exist {
			return errors.New("security group lite rule not exist")
		}

		*id = rs.Primary.ID

		return nil
	}
}

func testAccCheckSecurityGroupLiteRuleDestroy(id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

		ingress, egress, exist, err := service.DescribeSecurityGroupPolices(context.TODO(), *id)
		if err != nil {
			return err
		}

		if !exist || (len(ingress) == 0 && len(egress) == 0) {
			return nil
		}

		return errors.New("security group lite rule still exist")
	}
}

const testAccSecurityGroupLiteRuleBasic = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}

resource "tencentcloudenterprise_vpc_security_group_lite_rule" "foo" {
  security_group_id = cloud_vpc_security_group.foo.id

  ingress = [
    "ACCEPT#192.168.1.0/24#80#TCP",
    "DROP#8.8.8.8#80,90#UDP",
    "ACCEPT#0.0.0.0/0#80-90#TCP",
  ]

  egress = [
    "ACCEPT#192.168.0.0/16#ALL#TCP",
    "ACCEPT#10.0.0.0/8#ALL#ICMP",
    "DROP#0.0.0.0/0#ALL#ALL",
  ]
}
`

const testAccSecurityGroupLiteRuleUpdate1 = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}

resource "tencentcloudenterprise_vpc_security_group_lite_rule" "foo" {
  security_group_id = cloud_vpc_security_group.foo.id

  ingress = [
    "ACCEPT#192.168.1.0/24#80#TCP",
    "ACCEPT#192.168.1.0/26#800#TCP",
    "DROP#8.8.8.8#80,90#UDP",
    "ACCEPT#0.0.0.0/0#80-90#TCP",
  ]

  egress = [
    "ACCEPT#192.168.0.0/16#ALL#TCP",
    "ACCEPT#192.168.0.0/24#ALL#UDP",
    "ACCEPT#10.0.0.0/8#ALL#ICMP",
    "DROP#0.0.0.0/0#ALL#ALL",
  ]
}
`

const testAccSecurityGroupLiteRuleUpdate2 = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}

resource "tencentcloudenterprise_vpc_security_group_lite_rule" "foo" {
  security_group_id = cloud_vpc_security_group.foo.id

  ingress = [
    "ACCEPT#192.168.1.0/24#80#TCP",
  ]

  egress = [
    "ACCEPT#192.168.0.0/16#ALL#TCP",
  ]
}
`

const testAccSecurityGroupLiteRuleUpdate3 = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}

resource "tencentcloudenterprise_vpc_security_group_lite_rule" "foo" {
  security_group_id = cloud_vpc_security_group.foo.id

  ingress = [
    "ACCEPT#192.168.1.0/24#80#TCP",
  ]
}
`

const testAccSecurityGroupLiteRuleUpdate4 = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}

resource "tencentcloudenterprise_vpc_security_group_lite_rule" "foo" {
  security_group_id = cloud_vpc_security_group.foo.id

  egress = [
    "ACCEPT#192.168.0.0/16#ALL#TCP",
  ]
}
`

const testAccSecurityGroupLiteRuleUpdate5 = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}

resource "tencentcloudenterprise_vpc_security_group_lite_rule" "foo" {
  security_group_id = cloud_vpc_security_group.foo.id
}
`
const testAccSecurityGroupLiteRuleUpdate6 = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}

resource "tencentcloudenterprise_vpc_security_group" "group1" {
  name = "tf-test-sec"
}

resource "tencentcloudenterprise_vpc_address_template" "addr-foo" {
  name      = "tf-test-addr"
  addresses = ["203.0.113.1", "10.0.1.0/24", "203.0.113.1-203.0.113.100"]
}

resource "tencentcloudenterprise_vpc_address_template" "addr-bar" {
  name      = "cam-user-test"
  addresses = ["203.0.113.21", "10.0.3.0/24"]
}

resource "tencentcloudenterprise_vpc_address_template_group" "foo" {
  name      = "group-test"
  template_ids = [cloud_vpc_address_template.addr-foo.id, cloud_vpc_address_template.addr-bar.id]
}

resource "tencentcloudenterprise_vpc_security_group_lite_rule" "foo" {
  security_group_id = cloud_vpc_security_group.foo.id
  ingress = [
    "ACCEPT#${cloud_vpc_address_template_group.foo.id}#8080#TCP",
    "DROP#${cloud_vpc_address_template.addr-foo.id}#8080#TCP",
    "ACCEPT#0.0.0.0/0#80-90#TCP",
    "DROP#8.8.8.8#80,90#UDP",
    "ACCEPT#${cloud_vpc_security_group.group1.id}#80#TCP",
  ]
  egress = [
    "ACCEPT#192.168.0.0/16#ALL#TCP",
    "ACCEPT#${cloud_vpc_security_group.group1.id}#ALL#TCP",
  ]
}
`

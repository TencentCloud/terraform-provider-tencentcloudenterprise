package tencentcloud

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudSecurityGroupRuleSetResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupRuleSetResource_basic,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSecurityGroupRuleSetResource_basicExists("tencentcloudenterprise_vpc_security_group_rule_set.base"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.#", "5"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.0.cidr_block", "10.0.0.0/22"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.0.action", "ACCEPT"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.1.cidr_block", "203.0.113.21"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.1.action", "ACCEPT"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.2.cidr_block", "203.0.113.21"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.2.action", "ACCEPT"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.3.cidr_block", "203.0.113.182"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.3.action", "ACCEPT"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.4.description", "E:Block relative"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.4.action", "DROP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.#", "3"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_rule_set.base", "version"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_security_group_rule_set.base",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccSecurityGroupRuleSetResource_sort,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSecurityGroupRuleSetResource_basicExists("tencentcloudenterprise_vpc_security_group_rule_set.base"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.#", "5"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.#", "3"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.0.action", "DROP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.0.protocol", "TCP"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.0.port", "82"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.0.source_security_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.1.port", "80-90"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.1.description", "A:Allow Ips and 80-90"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.2.cidr_block", "203.0.113.21"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.2.description", "B:Allow UDP 8080"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.3.protocol", "ALL"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.4.protocol", "UDP"),

					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.0.cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.0.protocol", "ICMP"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.1.address_template_id"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.1.description", "B:Allow template"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.2.action", "ACCEPT"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.2.address_template_group"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_rule_set.base", "version"),
				),
			},
			{
				Config: testAccSecurityGroupRuleSetResource_modify,
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckSecurityGroupRuleSetResource_basicExists("tencentcloudenterprise_vpc_security_group_rule_set.base"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "ingress.#", "3"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.#", "4"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group_rule_set.base", "egress.1.description", "A:Block ping4"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_security_group_rule_set.base", "version"),
				),
			},
		},
	})
}

func testAccCheckSecurityGroupRuleSetResource_basicExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		cli, _ := sharedClientForRegion(defaultRegion)
		client := cli.(*TencentCloudClient).apiV3Conn
		service := VpcService{client}

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("resoure %s is not found", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("id is not set")
		}

		id := rs.Primary.ID

		_, err := service.DescribeSecurityGroupPolicies(ctx, id)
		if err != nil {
			return err
		}

		return nil
	}
}

const testAccSecurityGroupRuleSetDeps = `
resource "tencentcloudenterprise_vpc_security_group" "base" {
  name        = "test-set-sg"
  description = "Testing Rule Set Security"
}

resource "tencentcloudenterprise_vpc_security_group" "relative" {
  name        = "for-relative"
  description = "Used for attach security policy"
}

resource "tencentcloudenterprise_vpc_address_template" "foo" {
  name      = "test-set-aTemp"
  addresses = ["203.0.113.1", "10.0.1.0/24", "203.0.113.1-203.0.113.100"]
}

resource "tencentcloudenterprise_vpc_address_template_group" "foo" {
  name         = "test-set-atg"
  template_ids = [cloud_vpc_address_template.foo.id]
}
`

const testAccSecurityGroupRuleSetResource_basic = testAccSecurityGroupRuleSetDeps + `
resource "tencentcloudenterprise_vpc_security_group_rule_set" "base" {
  security_group_id = cloud_vpc_security_group.base.id

  ingress {
    action      = "ACCEPT"
    cidr_block  = "10.0.0.0/22"
    protocol    = "TCP"
    port        = "80-90"
    description = "A:Allow Ips and 80-90"
  }

  ingress {
    action      = "ACCEPT"
    cidr_block  = "203.0.113.21"
    protocol    = "UDP"
    port        = "8080"
    description = "B:Allow UDP 8080"
  }

  ingress {
    action      = "ACCEPT"
    cidr_block  = "203.0.113.21"
    protocol    = "UDP"
    port        = "8080"
    description = "C:Allow UDP 8080"
  }

  ingress {
    action      = "ACCEPT"
    cidr_block  = "203.0.113.182"
    protocol    = "ALL"
    port        = "ALL"
    description = "D:Allow ALL"
  }

  ingress {
    action             = "DROP"
    protocol           = "TCP"
    port               = "80"
    source_security_id = cloud_vpc_security_group.relative.id
    description        = "E:Block relative"
  }

  egress {
    action      = "DROP"
    cidr_block  = "10.0.0.0/16"
    protocol    = "ICMP"
    description = "A:Block ping3"
  }

  egress {
    action              = "DROP"
    address_template_id = cloud_vpc_address_template.foo.id
    description         = "B:Allow template"
  }

  egress {
    action              = "DROP"
    address_template_group = cloud_vpc_address_template_group.foo.id
    description         = "C:DROP template group"
  }
}
`

const testAccSecurityGroupRuleSetResource_sort = testAccSecurityGroupRuleSetDeps + `
resource "tencentcloudenterprise_vpc_security_group_rule_set" "base" {
  security_group_id = cloud_vpc_security_group.base.id

  ingress {
    action             = "DROP"
    protocol           = "TCP"
    port               = "82"
    source_security_id = cloud_vpc_security_group.relative.id
    description        = "E:Block relative and set 82"
  }

  ingress {
    action      = "ACCEPT"
    cidr_block  = "10.0.0.0/22"
    protocol    = "TCP"
    port        = "80-90"
    description = "A:Allow Ips and 80-90"
  }

  ingress {
    action      = "ACCEPT"
    cidr_block  = "203.0.113.21"
    protocol    = "UDP"
    port        = "8080"
    description = "B:Allow UDP 8080"
  }

  ingress {
    action      = "ACCEPT"
    cidr_block  = "203.0.113.182"
    protocol    = "ALL"
    port        = "ALL"
    description = "D:Allow ALL"
  }

  ingress {
    action      = "ACCEPT"
    cidr_block  = "203.0.113.21"
    protocol    = "UDP"
    port        = "8080"
    description = "C:Allow UDP 8080"
  }

  egress {
    action      = "DROP"
    cidr_block  = "10.0.0.0/16"
    protocol    = "ICMP"
    description = "A:Block ping3"
  }

  egress {
    action              = "DROP"
    address_template_id = cloud_vpc_address_template.foo.id
    description         = "B:Allow template"
  }

  egress {
    action              = "ACCEPT"
    address_template_group = cloud_vpc_address_template_group.foo.id
    description         = "C:ACCEPT template group"
  }
}
`

const testAccSecurityGroupRuleSetResource_modify = testAccSecurityGroupRuleSetDeps + `
resource "tencentcloudenterprise_vpc_security_group_rule_set" "base" {
  security_group_id = cloud_vpc_security_group.base.id

  ingress {
    action      = "ACCEPT"
    cidr_block  = "10.0.0.0/22"
    protocol    = "TCP"
    port        = "80-90"
    description = "A:Allow Ips and 80-90"
  }

  ingress {
    action      = "ACCEPT"
    cidr_block  = "203.0.113.21"
    protocol    = "UDP"
    port        = "8080"
    description = "C:Allow UDP 8080"
  }

  ingress {
    action             = "DROP"
    protocol           = "TCP"
    port               = "80"
    source_security_id = cloud_vpc_security_group.relative.id
    description        = "E:Block relative"
  }

  egress {
    action      = "DROP"
    cidr_block  = "10.0.0.0/16"
    protocol    = "ICMP"
    description = "A:Block ping3"
  }

  egress {
    action      = "DROP"
    cidr_block  = "203.0.113.182"
    protocol    = "ICMP"
    description = "A:Block ping4"
  }

  egress {
    action              = "DROP"
    address_template_id = cloud_vpc_address_template.foo.id
    description         = "B:Allow template"
  }

  egress {
    action              = "DROP"
    address_template_group = cloud_vpc_address_template_group.foo.id
    description         = "C:DROP template group"
  }
}
`

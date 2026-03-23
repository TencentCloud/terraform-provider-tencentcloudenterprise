package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccTencentCloudCcnAttachmentV2_basic(t *testing.T) {
	keyName := "tencentcloudenterprise_ccn_attachment_v2.example"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCcnAttachmentV2Destroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCcnAttachmentV2Config("attachment description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCcnAttachmentV2Exists(keyName),
					resource.TestCheckResourceAttrSet(keyName, "ccn_id"),
					resource.TestCheckResourceAttrSet(keyName, "instance_id"),
					resource.TestCheckResourceAttr(keyName, "instance_type", "VPC"),
					resource.TestCheckResourceAttrSet(keyName, "state"),
					resource.TestCheckResourceAttrSet(keyName, "attached_time"),
					resource.TestCheckResourceAttrSet(keyName, "cidr_block.#"),
				),
			},
			{
				Config: testAccCcnAttachmentV2Config("attachment description updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(keyName, "description", "attachment description updated"),
				),
			},
			{
				ResourceName:      keyName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudCcnRouteTable_basic(t *testing.T) {
	keyName := "tencentcloudenterprise_ccn_route_table.example"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCcnRouteTableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCcnRouteTableConfig("ci-temp-test-route-table", "route table description"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCcnRouteTableExists(keyName),
					resource.TestCheckResourceAttr(keyName, "name", "ci-temp-test-route-table"),
					resource.TestCheckResourceAttr(keyName, "description", "route table description"),
					resource.TestCheckResourceAttrSet(keyName, "is_default_table"),
					resource.TestCheckResourceAttrSet(keyName, "create_time"),
				),
			},
			{
				Config: testAccCcnRouteTableConfig("ci-temp-test-route-table-updated", "route table description updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(keyName, "name", "ci-temp-test-route-table-updated"),
					resource.TestCheckResourceAttr(keyName, "description", "route table description updated"),
				),
			},
			{
				ResourceName:      keyName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudCcnRouteTableAssociateInstanceConfig_basic(t *testing.T) {
	keyName := "tencentcloudenterprise_ccn_route_table_associate_instance_config.example"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCcnRouteTableAssociateInstanceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(keyName, "ccn_id"),
					resource.TestCheckResourceAttrSet(keyName, "route_table_id"),
					resource.TestCheckResourceAttr(keyName, "instances.#", "1"),
				),
			},
			{
				ResourceName:      keyName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudCcnRouteTableBroadcastPolicies_basic(t *testing.T) {
	keyName := "tencentcloudenterprise_ccn_route_table_broadcast_policies.example"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCcnRouteTableBroadcastPoliciesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCcnRouteTableBroadcastPoliciesConfig("broadcast policy"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(keyName, "ccn_id"),
					resource.TestCheckResourceAttrSet(keyName, "route_table_id"),
					resource.TestCheckResourceAttr(keyName, "policies.#", "1"),
					resource.TestCheckResourceAttr(keyName, "policies.0.action", "accept"),
					resource.TestCheckResourceAttr(keyName, "policies.0.description", "broadcast policy"),
					resource.TestCheckResourceAttr(keyName, "policies.0.route_conditions.#", "1"),
					resource.TestCheckResourceAttr(keyName, "policies.0.broadcast_conditions.#", "1"),
				),
			},
			{
				Config: testAccCcnRouteTableBroadcastPoliciesConfig("broadcast policy updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(keyName, "policies.0.description", "broadcast policy updated"),
				),
			},
			{
				ResourceName:      keyName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudCcnRouteTableInputPolicies_basic(t *testing.T) {
	keyName := "tencentcloudenterprise_ccn_route_table_input_policies.example"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckCcnRouteTableInputPoliciesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCcnRouteTableInputPoliciesConfig("input policy"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(keyName, "ccn_id"),
					resource.TestCheckResourceAttrSet(keyName, "route_table_id"),
					resource.TestCheckResourceAttr(keyName, "policies.#", "1"),
					resource.TestCheckResourceAttr(keyName, "policies.0.action", "accept"),
					resource.TestCheckResourceAttr(keyName, "policies.0.description", "input policy"),
					resource.TestCheckResourceAttr(keyName, "policies.0.route_conditions.#", "1"),
				),
			},
			{
				Config: testAccCcnRouteTableInputPoliciesConfig("input policy updated"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(keyName, "policies.0.description", "input policy updated"),
				),
			},
			{
				ResourceName:      keyName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCcnAttachmentV2Exists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}

		parts := strings.Split(rs.Primary.ID, FILED_SP)
		if len(parts) != 4 {
			return fmt.Errorf("invalid attachment id %s", rs.Primary.ID)
		}

		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		info, err := service.DescribeCcnAttachedInstanceByFilter(ctx, parts[0], parts[1], parts[2], parts[3])
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("ccn attachment %s not found", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckCcnAttachmentV2Destroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_ccn_attachment_v2" {
			continue
		}

		time.Sleep(5 * time.Second)
		parts := strings.Split(rs.Primary.ID, FILED_SP)
		if len(parts) != 4 {
			return fmt.Errorf("invalid attachment id %s", rs.Primary.ID)
		}

		info, err := service.DescribeCcnAttachedInstanceByFilter(ctx, parts[0], parts[1], parts[2], parts[3])
		if err != nil {
			return err
		}
		if info != nil {
			return fmt.Errorf("ccn attachment %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCcnRouteTableExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("resource %s not found", resourceName)
		}

		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		info, err := service.DescribeVpcCcnRouteTablesById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("ccn route table %s not found", rs.Primary.ID)
		}
		return nil
	}
}

func testAccCheckCcnRouteTableDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_ccn_route_table" {
			continue
		}

		time.Sleep(5 * time.Second)
		info, err := service.DescribeVpcCcnRouteTablesById(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}
		if info != nil {
			return fmt.Errorf("ccn route table %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCcnRouteTableBroadcastPoliciesDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_ccn_route_table_broadcast_policies" {
			continue
		}

		time.Sleep(5 * time.Second)
		parts := strings.Split(rs.Primary.ID, FILED_SP)
		if len(parts) != 2 {
			return fmt.Errorf("invalid broadcast policy id %s", rs.Primary.ID)
		}

		info, err := service.DescribeVpcReplaceCcnRouteTableBroadcastPolicysById(ctx, parts[0], parts[1])
		if err != nil {
			return err
		}
		if info != nil {
			return fmt.Errorf("broadcast policies %s still exist", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckCcnRouteTableInputPoliciesDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_ccn_route_table_input_policies" {
			continue
		}

		time.Sleep(5 * time.Second)
		parts := strings.Split(rs.Primary.ID, FILED_SP)
		if len(parts) != 2 {
			return fmt.Errorf("invalid input policy id %s", rs.Primary.ID)
		}

		info, err := service.DescribeVpcReplaceCcnRouteTableInputPolicysById(ctx, parts[0], parts[1])
		if err != nil {
			return err
		}
		if info != nil {
			return fmt.Errorf("input policies %s still exist", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCcnOnlyBaseConfig() string {
	return `
resource "tencentcloudenterprise_ccn" "example" {
  name                 = "ci-temp-test-ccn-migrate"
  description          = "ci-temp-test-ccn-migrate-des"
  qos                  = "AG"
  charge_type          = "PREPAID"
  bandwidth_limit_type = "INTER_REGION_LIMIT"
}
`
}

func testAccCcnVpcBaseConfig() string {
	return fmt.Sprintf(`
variable "region" {
  default = "%s"
}

resource "tencentcloudenterprise_vpc" "vpc" {
  name         = "ci-temp-test-vpc-ccn-migrate"
  cidr_block   = "%s"
  dns_servers  = ["119.29.29.29", "8.8.8.8"]
  is_multicast = false
}

%s
`, defaultRegion, defaultVpcCidr, testAccCcnOnlyBaseConfig())
}

func testAccCcnAttachmentV2Config(description string) string {
	return testAccCcnVpcBaseConfig() + fmt.Sprintf(`
resource "tencentcloudenterprise_ccn_attachment_v2" "example" {
  ccn_id          = tencentcloudenterprise_ccn.example.id
  instance_id     = tencentcloudenterprise_vpc.vpc.id
  instance_type   = "VPC"
  instance_region = var.region
  description     = %q
}
`, description)
}

func testAccCcnRouteTableConfig(name, description string) string {
	return testAccCcnOnlyBaseConfig() + fmt.Sprintf(`
resource "tencentcloudenterprise_ccn_route_table" "example" {
  ccn_id      = tencentcloudenterprise_ccn.example.id
  name        = %q
  description = %q
}
`, name, description)
}

func testAccCcnRouteTableAssociateInstanceConfig() string {
	return testAccCcnVpcBaseConfig() + `
resource "tencentcloudenterprise_ccn_route_table" "example" {
  ccn_id      = tencentcloudenterprise_ccn.example.id
  name        = "ci-temp-test-route-table-associate"
  description = "route table associate description"
}

resource "tencentcloudenterprise_ccn_attachment_v2" "attachment" {
  ccn_id          = tencentcloudenterprise_ccn.example.id
  instance_id     = tencentcloudenterprise_vpc.vpc.id
  instance_type   = "VPC"
  instance_region = var.region
}

resource "tencentcloudenterprise_ccn_route_table_associate_instance_config" "example" {
  ccn_id         = tencentcloudenterprise_ccn.example.id
  route_table_id = tencentcloudenterprise_ccn_route_table.example.id
  instances {
    instance_id   = tencentcloudenterprise_vpc.vpc.id
    instance_type = "VPC"
  }

  depends_on = [tencentcloudenterprise_ccn_attachment_v2.attachment]
}
`
}

func testAccCcnRouteTableBroadcastPoliciesConfig(description string) string {
	return testAccCcnVpcBaseConfig() + fmt.Sprintf(`
resource "tencentcloudenterprise_ccn_route_table" "example" {
  ccn_id      = tencentcloudenterprise_ccn.example.id
  name        = "ci-temp-test-route-table-broadcast"
  description = "route table broadcast description"
}

resource "tencentcloudenterprise_ccn_attachment_v2" "attachment" {
  ccn_id          = tencentcloudenterprise_ccn.example.id
  instance_id     = tencentcloudenterprise_vpc.vpc.id
  instance_type   = "VPC"
  instance_region = var.region
  route_table_id  = tencentcloudenterprise_ccn_route_table.example.id
}

resource "tencentcloudenterprise_ccn_route_table_broadcast_policies" "example" {
  ccn_id         = tencentcloudenterprise_ccn.example.id
  route_table_id = tencentcloudenterprise_ccn_route_table.example.id
  policies {
    action      = "accept"
    description = %q
    route_conditions {
      name          = "instance-region"
      values        = [var.region]
      match_pattern = 1
    }
    broadcast_conditions {
      name          = "instance-region"
      values        = ["ap-shanghai"]
      match_pattern = 1
    }
  }

  depends_on = [tencentcloudenterprise_ccn_attachment_v2.attachment]
}
`, description)
}

func testAccCcnRouteTableInputPoliciesConfig(description string) string {
	return testAccCcnVpcBaseConfig() + fmt.Sprintf(`
resource "tencentcloudenterprise_ccn_route_table" "example" {
  ccn_id      = tencentcloudenterprise_ccn.example.id
  name        = "ci-temp-test-route-table-input"
  description = "route table input description"
}

resource "tencentcloudenterprise_ccn_attachment_v2" "attachment" {
  ccn_id          = tencentcloudenterprise_ccn.example.id
  instance_id     = tencentcloudenterprise_vpc.vpc.id
  instance_type   = "VPC"
  instance_region = var.region
  route_table_id  = tencentcloudenterprise_ccn_route_table.example.id
}

resource "tencentcloudenterprise_ccn_route_table_input_policies" "example" {
  ccn_id         = tencentcloudenterprise_ccn.example.id
  route_table_id = tencentcloudenterprise_ccn_route_table.example.id
  policies {
    action      = "accept"
    description = %q
    route_conditions {
      name          = "instance-region"
      values        = [var.region]
      match_pattern = 1
    }
  }

  depends_on = [tencentcloudenterprise_ccn_attachment_v2.attachment]
}
`, description)
}

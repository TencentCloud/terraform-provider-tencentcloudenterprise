package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func init() {
	resource.AddTestSweepers("tencentcloudenterprise_vpc_security_group", &resource.Sweeper{
		Name: "tencentcloudenterprise_vpc_security_group",
		F:    testSweepSecurityGroups,
	})
}

func testSweepSecurityGroups(region string) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	sharedClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("getting tencentcloud client error: %s", err.Error())
	}
	client := sharedClient.(*TencentCloudClient)
	service := VpcService{
		client: client.apiV3Conn,
	}

	sgs, err := service.DescribeSecurityGroups(ctx, nil, nil, nil, nil)

	if err != nil {
		return fmt.Errorf("DescribeSecurityGroups error: %s", err.Error())
	}

	for _, v := range sgs {
		name := *v.SecurityGroupName
		id := *v.SecurityGroupId

		now := time.Now()
		createTime := stringTotime(*v.CreatedTime)
		interval := now.Sub(createTime).Minutes()

		if strings.HasPrefix(name, keepResource) || strings.HasPrefix(name, defaultResource) {
			continue
		}

		// less than 30 minute, not delete
		if needProtect == 1 && int64(interval) < 30 {
			continue
		}

		if err := service.DeleteSecurityGroup(ctx, id); err != nil {
			log.Printf("[ERROR] sweep security group %s error: %s", id, err.Error())
		}
	}

	return nil
}

func TestAccTencentCloudSecurityGroup_basic(t *testing.T) {
	t.Parallel()
	var sgId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityGroupDestroy(&sgId),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupExists("tencentcloudenterprise_vpc_security_group.foo", &sgId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "name", "ci-temp-test-sg"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "description", ""),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "project_id", "0"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_security_group.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudSecurityGroup_update(t *testing.T) {
	t.Parallel()
	var sgId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityGroupDestroy(&sgId),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupExists("tencentcloudenterprise_vpc_security_group.foo", &sgId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "name", "ci-temp-test-sg"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "description", "ci-temp-test-sg-desc"),
				),
			},
			{
				Config: testAccSecurityGroupConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupExists("tencentcloudenterprise_vpc_security_group.foo", &sgId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "name", "ci-temp-test-sg-updated"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "description", "ci-temp-test-sg-desc-updated"),
				),
			},
		},
	})
}

func TestAccTencentCloudSecurityGroup_tags(t *testing.T) {
	t.Parallel()
	var sgId string

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckSecurityGroupDestroy(&sgId),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupConfigTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupExists("tencentcloudenterprise_vpc_security_group.foo", &sgId),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "description", ""),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "tags.test", "test"),
				),
			},
			{
				Config: testAccSecurityGroupConfigTagsUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckSecurityGroupExists("tencentcloudenterprise_vpc_security_group.foo", &sgId),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "tags.test"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_security_group.foo", "tags.abc", "abc"),
				),
			},
		},
	})
}

func testAccCheckSecurityGroupDestroy(id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*TencentCloudClient).apiV3Conn
		service := VpcService{client: client}

		sg, err := service.DescribeSecurityGroup(context.TODO(), *id)
		if err != nil {
			return err
		}

		if sg != nil {
			return fmt.Errorf("security group still exists")
		}

		return nil
	}
}

func testAccCheckSecurityGroupExists(n string, id *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("no security group ID is set")
		}

		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}

		sg, err := service.DescribeSecurityGroup(context.TODO(), rs.Primary.ID)
		if err != nil {
			return err
		}

		if sg == nil {
			return fmt.Errorf("security group not found: %s", rs.Primary.ID)
		}

		*id = rs.Primary.ID

		return nil
	}
}

const testAccSecurityGroupConfigBasic = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"
}
`

const testAccSecurityGroupConfig = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name        = "ci-temp-test-sg"
  description = "ci-temp-test-sg-desc"
}
`
const testAccSecurityGroupConfigUpdate = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name        = "ci-temp-test-sg-updated"
  description = "ci-temp-test-sg-desc-updated"
}
`

const testAccSecurityGroupConfigTags = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"

  tags = {
    "test" = "test"
  }
}
`

const testAccSecurityGroupConfigTagsUpdate = `
resource "tencentcloudenterprise_vpc_security_group" "foo" {
  name = "ci-temp-test-sg"

  tags = {
    "abc" = "abc"
  }
}
`

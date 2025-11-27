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
	resource.AddTestSweepers("tencentcloudenterprise_vpc", &resource.Sweeper{
		Name: "tencentcloudenterprise_vpc",
		F:    testSweepVpcInstance,
	})
}

func testSweepVpcInstance(region string) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	sharedClient, err := sharedClientForRegion(region)
	if err != nil {
		return fmt.Errorf("getting tencentcloud client error: %s", err.Error())
	}
	client := sharedClient.(*TencentCloudClient)

	vpcService := VpcService{
		client: client.apiV3Conn,
	}

	instances, err := vpcService.DescribeVpcs(ctx, "", "", nil, nil, "", "")
	if err != nil {
		return fmt.Errorf("get instance list error: %s", err.Error())
	}

	for _, v := range instances {
		instanceId := v.vpcId
		instanceName := v.name

		now := time.Now()

		createTime := stringTotime(v.createTime)
		interval := now.Sub(createTime).Minutes()
		if strings.HasPrefix(instanceName, keepResource) || strings.HasPrefix(instanceName, defaultResource) {
			continue
		}
		// less than 30 minute, not delete
		if needProtect == 1 && int64(interval) < 30 {
			continue
		}

		if err = vpcService.DeleteVpc(ctx, instanceId); err != nil {
			log.Printf("[ERROR] sweep instance %s error: %s", instanceId, err.Error())
		}
	}

	return nil
}

func TestAccTencentCloudVpcV3Basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcExists("tencentcloudenterprise_vpc.foo"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "cidr_block", defaultVpcCidr),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "is_multicast", "true"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "default_route_table_id"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudVpcV3Update(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcExists("tencentcloudenterprise_vpc.foo"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "cidr_block", defaultVpcCidr),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "is_multicast", "true"),

					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "is_default"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "create_time"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "dns_servers.#"),
				),
			},
			{
				Config: testAccVpcConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcExists("tencentcloudenterprise_vpc.foo"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "cidr_block", defaultVpcCidrLess),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "name", defaultInsNameUpdate),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "is_multicast", "false"),

					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "is_default"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "create_time"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "dns_servers.#"),

					resource.TestCheckTypeSetElemAttr("tencentcloudenterprise_vpc.foo", "dns_servers.*", "203.0.113.29"),
					resource.TestCheckTypeSetElemAttr("tencentcloudenterprise_vpc.foo", "dns_servers.*", "203.0.113.116"),
				),
			},
		},
	})
}

func TestAccTencentCloudVpcV3WithTags(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcConfigWithTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcExists("tencentcloudenterprise_vpc.foo"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "cidr_block", defaultVpcCidr),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "is_multicast", "true"),

					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "is_default"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "create_time"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc.foo", "dns_servers.#"),

					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "tags.test", "test"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc.foo", "tags.abc"),
				),
			},
			{
				Config: testAccVpcConfigWithTagsUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcExists("tencentcloudenterprise_vpc.foo"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "cidr_block", defaultVpcCidr),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "is_multicast", "true"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc.foo", "tags.abc", "abc"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc.foo", "tags.test"),
				),
			},
		},
	})
}

func testAccCheckVpcExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		_, has, err := service.DescribeVpc(ctx, rs.Primary.ID, "", "")
		if err != nil {
			return err
		}
		if has > 0 {
			return nil
		}

		return fmt.Errorf("vpc %s not exists", rs.Primary.ID)
	}
}

func testAccCheckVpcDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_vpc" {
			continue
		}
		time.Sleep(5 * time.Second)
		_, has, err := service.DescribeVpc(ctx, rs.Primary.ID, "", "")
		if err != nil {
			return err
		}
		if has == 0 {
			return nil
		}
		return fmt.Errorf("vpc %s still exists", rs.Primary.ID)
	}

	return nil
}

const testAccVpcConfig = defaultVpcVariable + `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = var.instance_name
  cidr_block = var.vpc_cidr
}
`

const testAccVpcConfigUpdate = defaultVpcVariable + `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = var.instance_name_update
  cidr_block = var.vpc_cidr_less
  dns_servers  = ["203.0.113.29", "203.0.113.116"]
  is_multicast = false
}
`

const testAccVpcConfigWithTags = defaultVpcVariable + `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = var.instance_name
  cidr_block = var.vpc_cidr

  tags = {
    "test" = "test"
  }
}
`

const testAccVpcConfigWithTagsUpdate = defaultVpcVariable + `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = var.instance_name
  cidr_block = var.vpc_cidr

  tags = {
    "abc" = "abc"
  }
}
`

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
	resource.AddTestSweepers("tencentcloudenterprise_vpc_subnet", &resource.Sweeper{
		Name: "tencentcloudenterprise_vpc_subnet",
		F:    testSweepSubnet,
	})
}

func testSweepSubnet(region string) error {
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

	instances, err := vpcService.DescribeSubnets(ctx, "", "", "", "",
		nil, nil, nil, "", "")
	if err != nil {
		return fmt.Errorf("get instance list error: %s", err.Error())
	}

	for _, v := range instances {

		instanceId := v.subnetId
		instanceName := v.name
		now := time.Now()
		createTime := stringTotime(v.createTime)
		interval := now.Sub(createTime).Minutes()

		if strings.HasPrefix(instanceName, keepResource) || strings.HasPrefix(instanceName, defaultResource) {
			continue
		}

		if needProtect == 1 && int64(interval) < 30 {
			continue
		}

		if err = vpcService.DeleteSubnet(ctx, instanceId); err != nil {
			log.Printf("[ERROR] sweep instance %s error: %s", instanceId, err.Error())
		}
	}

	return nil
}

func TestAccTencentCloudVpcV3SubnetBasic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcSubnetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcSubnetExists("tencentcloudenterprise_vpc_subnet.subnet"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "cidr_block", defaultSubnetCidr),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "is_multicast", "false"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_vpc_subnet.subnet",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccTencentCloudVpcV3SubnetUpdate(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcSubnetConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcSubnetExists("tencentcloudenterprise_vpc_subnet.subnet"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "cidr_block", defaultSubnetCidr),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "is_multicast", "false"),

					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "availability_zone"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "is_default"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "available_ip_count"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "route_table_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "create_time"),
				),
			},
			{
				Config: testAccVpcSubnetConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcSubnetExists("tencentcloudenterprise_vpc_subnet.subnet"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "cidr_block", defaultSubnetCidrLess),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "name", defaultInsNameUpdate),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "is_multicast", "true"),

					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "availability_zone"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "is_default"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "available_ip_count"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "route_table_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "create_time"),
				),
			},
		},
	})
}

func TestAccTencentCloudVpcV3SubnetWithTags(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVpcSubnetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcSubnetConfigWithTags,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcSubnetExists("tencentcloudenterprise_vpc_subnet.subnet"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "cidr_block", defaultSubnetCidr),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "is_multicast", "false"),

					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "availability_zone"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "is_default"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "available_ip_count"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "route_table_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_vpc_subnet.subnet", "create_time"),

					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "tags.test", "test"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "tags.abc"),
				),
			},
			{
				Config: testAccVpcSubnetConfigWithTagsUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckVpcSubnetExists("tencentcloudenterprise_vpc_subnet.subnet"),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "cidr_block", defaultSubnetCidr),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "name", defaultInsName),
					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "is_multicast", "false"),

					resource.TestCheckResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "tags.abc", "abc"),
					resource.TestCheckNoResourceAttr("tencentcloudenterprise_vpc_subnet.subnet", "tags.test"),
				),
			},
		},
	})
}

func testAccCheckVpcSubnetExists(r string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[r]
		if !ok {
			return fmt.Errorf("resource %s is not found", r)
		}

		service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		_, has, err := service.DescribeSubnet(ctx, rs.Primary.ID, nil, "", "")
		if err != nil {
			return err
		}
		if has > 0 {
			return nil
		}

		return fmt.Errorf("subnet %s not exists", rs.Primary.ID)
	}
}

func testAccCheckVpcSubnetDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_vpc_subnet" {
			continue
		}
		time.Sleep(5 * time.Second)
		_, has, err := service.DescribeSubnet(ctx, rs.Primary.ID, nil, "", "")
		if err != nil {
			return err
		}
		if has == 0 {
			return nil
		}
		return fmt.Errorf("subnet %s still exists", rs.Primary.ID)
	}

	return nil
}

const testAccVpcSubnetConfig = defaultVpcVariable + `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = var.instance_name
  cidr_block = var.vpc_cidr
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  name              = var.instance_name
  vpc_id            = cloud_vpc.foo.id
  availability_zone = var.availability_zone
  cidr_block        = var.subnet_cidr
  is_multicast      = false
}
`

const testAccVpcSubnetConfigUpdate = defaultVpcVariable + `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = var.instance_name
  cidr_block = var.vpc_cidr
}

resource "tencentcloudenterprise_route_table" "route_table" {
  name   = var.instance_name
  vpc_id = cloud_vpc.foo.id
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  name              = var.instance_name_update
  vpc_id            = cloud_vpc.foo.id
  availability_zone = var.availability_zone
  cidr_block        = var.subnet_cidr_less
  is_multicast      = true
  route_table_id    = cloud_route_table.route_table.id
}
`

const testAccVpcSubnetConfigWithTags = defaultVpcVariable + `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = var.instance_name
  cidr_block = var.vpc_cidr
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  name              = var.instance_name
  vpc_id            = cloud_vpc.foo.id
  availability_zone = var.availability_zone
  cidr_block        = var.subnet_cidr
  is_multicast      = false

  tags = {
    "test" = "test"
  }
}
`

const testAccVpcSubnetConfigWithTagsUpdate = defaultVpcVariable + `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = var.instance_name
  cidr_block = var.vpc_cidr
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  name              = var.instance_name
  vpc_id            = cloud_vpc.foo.id
  availability_zone = var.availability_zone
  cidr_block        = var.subnet_cidr
  is_multicast      = false

  tags = {
    "abc" = "abc"
  }
}
`

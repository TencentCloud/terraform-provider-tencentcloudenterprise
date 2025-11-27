package tencentcloud

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
)

// go test -i; go test -test.run TestAccTencentCloudTdmqRabbitmqVipInstanceResource_basic -v
func TestAccTencentCloudTdmqRabbitmqVipInstanceResource_basic(t *testing.T) {
	//t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccCheckTdmqRabbitmqVipInstanceDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTdmqRabbitmqVipInstance,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTdmqRabbitmqVipInstanceExists("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "zone_ids.#"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "subnet_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "cluster_name"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "node_spec"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "node_num"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "storage_size"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "enable_create_default_ha_mirror_queue"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "auto_renew_flag"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "time_span"),
					testAccStepTimeSleepDuration(1*time.Minute),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccTdmqRabbitmqVipInstanceUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTdmqRabbitmqVipInstanceExists("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "zone_ids.#"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "vpc_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "subnet_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "cluster_name"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "node_spec"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "node_num"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "storage_size"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "enable_create_default_ha_mirror_queue"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "auto_renew_flag"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_vip_instance.example", "time_span"),
					testAccStepTimeSleepDuration(1*time.Minute),
				),
			},
		},
	})
}

func testAccCheckTdmqRabbitmqVipInstanceExists(re string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[re]
		if !ok {
			return fmt.Errorf("rabbitmq vip instance  %s is not found", re)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("rabbitmq vip instance id is not set")
		}

		service := TdmqService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		id := rs.Primary.ID

		ret, err := service.DescribeTdmqRabbitmqVipInstanceById(ctx, id)
		if err != nil {
			return err
		}

		if ret == nil {
			return fmt.Errorf("tdmq rabbitmq vip instance not found, id: %v", id)
		}

		return nil
	}
}

func testAccCheckTdmqRabbitmqVipInstanceDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" {
			continue
		}

		id := rs.Primary.ID
		ret, err := service.DescribeTdmqRabbitmqVipInstanceById(ctx, id)
		if err != nil {
			code := err.(*sdkErrors.CloudSDKError).Code
			if code == "InternalError" || code == "FailedOperation" {
				return nil
			}

			return err
		}

		if ret != nil {
			return fmt.Errorf("tdmq rabbitmq vip instance exist, id: %v", id)
		}
	}

	return nil
}

const testAccTdmqRabbitmqVipInstance = `
data "tencentcloudenterprise_availability_zones" "zones" {
  name = "ap-guangzhou-6"
}

# create vpc
resource "tencentcloudenterprise_vpc" "vpc" {
  name       = "vpc"
  cidr_block = "10.0.0.0/16"
}

# create vpc subnet
resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  name              = "subnet"
  vpc_id            = cloud_vpc.vpc.id
  availability_zone = "ap-guangzhou-6"
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

# create rabbitmq instance
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "example" {
  zone_ids                              = [data.cloud_availability_zones.zones.zones.0.id]
  vpc_id                                = cloud_vpc.vpc.id
  subnet_id                             = cloud_vpc_subnet.subnet.id
  cluster_name                          = "tf-example-rabbitmq-vip-instance"
  node_spec                             = "rabbit-vip-basic-1"
  node_num                              = 1
  storage_size                          = 200
  enable_create_default_ha_mirror_queue = false
  auto_renew_flag                       = true
  time_span                             = 1
}
`

const testAccTdmqRabbitmqVipInstanceUpdate = `
data "tencentcloudenterprise_availability_zones" "zones" {
  name = "ap-guangzhou-6"
}

# create vpc
resource "tencentcloudenterprise_vpc" "vpc" {
  name       = "vpc"
  cidr_block = "10.0.0.0/16"
}

# create vpc subnet
resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  name              = "subnet"
  vpc_id            = cloud_vpc.vpc.id
  availability_zone = "ap-guangzhou-6"
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

# create rabbitmq instance
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "example" {
  zone_ids                              = [data.cloud_availability_zones.zones.zones.0.id]
  vpc_id                                = cloud_vpc.vpc.id
  subnet_id                             = cloud_vpc_subnet.subnet.id
  cluster_name                          = "tf-example-rabbitmq-vip-instance-update"
  node_spec                             = "rabbit-vip-basic-1"
  node_num                              = 1
  storage_size                          = 200
  enable_create_default_ha_mirror_queue = false
  auto_renew_flag                       = true
  time_span                             = 1
}
`

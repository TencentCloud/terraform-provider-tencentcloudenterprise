package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// go test -i; go test -test.run TestAccTencentCloudTdmqRabbitmqUserResource_basic -v
func TestAccTencentCloudTdmqRabbitmqUserResource_basic(t *testing.T) {
	//t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTdmqRabbitmqUser,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "instance_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "user"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "password"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "description"),
				),
			},
			{
				ResourceName:      "tencentcloudenterprise_tdmq_rabbitmq_user.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccTdmqRabbitmqUserUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "instance_id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "user"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "password"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "description"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "max_connections"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_tdmq_rabbitmq_user.rabbitmq_user", "max_channels"),
				),
			},
		},
	})
}

const testAccTdmqRabbitmqUser = `
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

# create rabbitmq user
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "example" {
  instance_id     = cloud_tdmq_rabbitmq_vip_instance.example.id
  user            = "tf-example-user"
  password        = "$Password"
  description     = "desc."
  tags            = ["management", "monitoring", "example"]
}
`

const testAccTdmqRabbitmqUserUpdate = `
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

# create rabbitmq user
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "example" {
  instance_id     = cloud_tdmq_rabbitmq_vip_instance.example.id
  user            = "tf-example-user"
  password        = "$Password"
  description     = "desc update."
  tags            = ["management", "monitoring", "example_update"]
  max_connections = 3
  max_channels    = 3
}
`

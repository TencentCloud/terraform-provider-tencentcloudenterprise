package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTencentCloudVpcV3Subnets_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudVpcSubnets,
				Check: resource.ComposeTestCheckFunc(
					// id filter
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_subnets.id_instances"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnets.id_instances", "instance_list.#", "1"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnets.id_instances", "instance_list.0.availability_zone", "ap-guangzhou-3"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnets.id_instances", "instance_list.0.name", "guagua_vpc_subnet_test"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnets.id_instances", "instance_list.0.cidr_block", "10.0.20.0/28"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnets.id_instances", "instance_list.0.is_multicast", "false"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.id_instances", "instance_list.0.vpc_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.id_instances", "instance_list.0.subnet_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.id_instances", "instance_list.0.route_table_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.id_instances", "instance_list.0.is_default"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.id_instances", "instance_list.0.available_ip_count"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.id_instances", "instance_list.0.create_time"),

					// vpc_id filter ,Every subnet with the query vpc_id will be found
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_subnets.vpc_instances"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.availability_zone"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.name"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.cidr_block"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.is_multicast"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.vpc_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.subnet_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.route_table_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.is_default"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.available_ip_count"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.vpc_instances", "instance_list.0.create_time"),

					// name filter ,Every subnet with a "guagua_vpc_subnet_test" name will be found
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_subnets.name_instances"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.availability_zone"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.name"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.cidr_block"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.is_multicast"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.vpc_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.subnet_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.route_table_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.is_default"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.available_ip_count"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.name_instances", "instance_list.0.create_time"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnets.name_instances", "instance_list.0.name", "guagua_vpc_subnet_test"),

					// tags filter ,Every subnet with a tag test:test will be found
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_subnets.tags_instances"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.availability_zone"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.name"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.cidr_block"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.is_multicast"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.vpc_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.subnet_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.route_table_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.is_default"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.available_ip_count"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.tags_instances", "instance_list.0.create_time"),
					resource.TestCheckResourceAttr("data.cloud_vpc_subnets.tags_instances", "instance_list.0.tags.test", "test"),

					// name filter ,Every subnet with cidr_block "10.0.20.0/28" will be found
					//testAccCheckTencentCloudDataSourceID("data.cloud_vpc_subnets.cidr_block_instances"),
					//resource.TestCheckResourceAttrSet("data.cloud_vpc_subnets.cidr_block_instances", "instance_list.#"),
					//resource.TestCheckResourceAttr("data.cloud_vpc_subnets.cidr_block_instances", "instance_list.0.cidr_block", "10.0.20.0/28"),
				),
			},
		},
	})
}

const TestAccDataSourceTencentCloudVpcSubnets = `
variable "availability_zone" {
  default = "ap-guangzhou-3"
}

resource "tencentcloudenterprise_vpc" "foo" {
  name       = "guagua_vpc_instance_test"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  availability_zone = var.availability_zone
  name              = "guagua_vpc_subnet_test"
  vpc_id            = cloud_vpc.foo.id
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_vpc_subnets" "vpc_instances" {
  vpc_id = cloud_vpc_subnet.subnet.vpc_id
}

data "tencentcloudenterprise_vpc_subnets" "id_instances" {
  subnet_id = cloud_vpc_subnet.subnet.id
}

data "tencentcloudenterprise_vpc_subnets" "cidr_block_instances" {
  cidr_block = cloud_vpc_subnet.subnet.cidr_block
}

data "tencentcloudenterprise_vpc_subnets" "name_instances" {
  name = cloud_vpc_subnet.subnet.name
}

data "tencentcloudenterprise_vpc_subnets" "tags_instances" {
  tags = cloud_vpc_subnet.subnet.tags
}
`

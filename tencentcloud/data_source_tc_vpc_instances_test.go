package tencentcloud

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceTencentCloudVpcV3Instances_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: TestAccDataSourceTencentCloudVpcInstances,

				Check: resource.ComposeTestCheckFunc(
					// id filter
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_instances.id_instances"),
					resource.TestCheckResourceAttr("data.cloud_vpc_instances.id_instances", "instance_list.#", "1"),
					resource.TestCheckResourceAttr("data.cloud_vpc_instances.id_instances", "instance_list.0.name", "guagua_vpc_instance_test"),
					resource.TestCheckResourceAttr("data.cloud_vpc_instances.id_instances", "instance_list.0.cidr_block", "10.0.0.0/16"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.id_instances", "instance_list.0.vpc_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.id_instances", "instance_list.0.is_default"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.id_instances", "instance_list.0.is_multicast"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.id_instances", "instance_list.0.dns_servers.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.id_instances", "instance_list.0.subnet_ids.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.id_instances", "instance_list.0.create_time"),
					resource.TestCheckResourceAttr("data.cloud_vpc_instances.id_instances", "instance_list.0.tags.test", "test"),

					// name filter ,Every VPC with a "guagua_vpc_instance_test" name will be found
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_instances.name_instances"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.0.name"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.0.cidr_block"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.0.vpc_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.0.is_default"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.0.is_multicast"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.0.dns_servers.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.0.subnet_ids.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.name_instances", "instance_list.0.create_time"),

					// tag filter ,Every VPC with a tag test:test will be found
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_instances.tags_instances"),
					resource.TestMatchResourceAttr("data.cloud_vpc_instances.tags_instances", "instance_list.#", regexp.MustCompile(`^[1-9]\d*$`)),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.tags_instances", "instance_list.0.name"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.tags_instances", "instance_list.0.cidr_block"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.tags_instances", "instance_list.0.vpc_id"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.tags_instances", "instance_list.0.is_default"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.tags_instances", "instance_list.0.is_multicast"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.tags_instances", "instance_list.0.dns_servers.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.tags_instances", "instance_list.0.subnet_ids.#"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.tags_instances", "instance_list.0.create_time"),
					resource.TestCheckResourceAttr("data.cloud_vpc_instances.tags_instances", "instance_list.0.tags.test", "test"),

					// cidr filter ,Every VPC with  "10.0.0.0/16" cidr will be found
					testAccCheckTencentCloudDataSourceID("data.cloud_vpc_instances.cidr_instances"),
					resource.TestCheckResourceAttrSet("data.cloud_vpc_instances.cidr_instances", "instance_list.#"),
					resource.TestCheckResourceAttr("data.cloud_vpc_instances.cidr_instances", "instance_list.0.cidr_block", "10.0.0.0/16"),
				),
			},
		},
	})
}

const TestAccDataSourceTencentCloudVpcInstances = `
resource "tencentcloudenterprise_vpc" "foo" {
  name       = "guagua_vpc_instance_test"
  cidr_block = "10.0.0.0/16"

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_vpc_instances" "id_instances" {
  vpc_id = cloud_vpc.foo.id
}

data "tencentcloudenterprise_vpc_instances" "cidr_instances" {
  cidr_block = cloud_vpc.foo.cidr_block
}

data "tencentcloudenterprise_vpc_instances" "name_instances" {
  name = cloud_vpc.foo.name
}

data "tencentcloudenterprise_vpc_instances" "tags_instances" {
  tags = cloud_vpc.foo.tags
}
`

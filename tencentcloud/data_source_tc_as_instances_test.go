package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudAsInstancesDataSource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAsInstancesDataSource_basic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTencentCloudDataSourceID("data.cloud_as_instances.instances"),
					resource.TestCheckResourceAttrSet("data.cloud_as_instances.instances", "instance_list.#"),
				),
			},
		},
	})
}

func testAccAsInstancesDataSource_basic() string {
	return defaultAsVariable + `
resource "tencentcloudenterprise_vpc" "vpc" {
  name       = "tf-as-vpc"
  cidr_block = "10.2.0.0/16"
}

resource "tencentcloudenterprise_vpc_subnet" "subnet" {
  vpc_id            = cloud_vpc.vpc.id
  name              = "tf-as-subnet"
  cidr_block        = "10.2.11.0/24"
  availability_zone = var.availability_zone
}

resource "tencentcloudenterprise_as_scaling_config" "launch_configuration" {
  configuration_name = "tf-as-configuration-ds-ins-basic"
  image_id           = "img-2lr9q49h"
  instance_types     = [data.cloud_cvm_instance_types.default.instance_types.0.instance_type]
}

resource "tencentcloudenterprise_as_scaling_group" "scaling_group" {
  scaling_group_name = "tf-as-group-ds-ins-basic"
  configuration_id   = cloud_as_scaling_config.launch_configuration.id
  max_size           = 1
  min_size           = 1
  vpc_id             = cloud_vpc.vpc.id
  subnet_ids         = [cloud_vpc_subnet.subnet.id]

  tags = {
    "test" = "test"
  }
}

data "tencentcloudenterprise_as_instances" "instances" {
	filters {
		name = "auto-scaling-group-id"
		values = [cloud_as_scaling_group.scaling_group.id]
  }
}

`
}

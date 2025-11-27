package tencentcloud

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDcdbActivateHourInstanceOperationResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDcdbIsolateHourInstanceOperation, "oper_isolate"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_isolate_hour_instance_operation.isolate_operation", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_isolate_hour_instance_operation.isolate_operation", "instance_id"),
				),
			},
			{
				Config: fmt.Sprintf(testAccDcdbActivateHourInstanceOperation, "oper_isolate"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_activate_hour_instance_operation.activate_operation", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_activate_hour_instance_operation.activate_operation", "instance_id"),
				),
			},
		},
	})
}

const testAccDcdbHourInsOperation = defaultAzVariable + `
data "tencentcloudenterprise_vpc_security_groups" "internal" {
	name = "default"
}

data "tencentcloudenterprise_vpc_instances" "vpc" {
	name ="Default-VPC"
}
	
data "tencentcloudenterprise_vpc_subnets" "subnet" {
	vpc_id = data.cloud_vpc_instances.vpc.instance_list.0.vpc_id
}

locals {
	vpc_id = data.cloud_vpc_subnets.subnet.instance_list.0.vpc_id
	subnet_id = data.cloud_vpc_subnets.subnet.instance_list.0.subnet_id
	sg_id = data.cloud_vpc_security_groups.internal.security_groups.0.security_group_id
}

resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance" {
	instance_name = "test_dcdb_db_hourdb_instance_%s"
	zones = [var.default_az]
	shard_memory = "2"
	shard_storage = "10"
	shard_node_count = "2"
	shard_count = "2"
	vpc_id = local.vpc_id
	subnet_id = local.subnet_id
	security_group_id = local.sg_id
	db_version_id = "8.0"
	resource_tags {
	  tag_key = "aaa"
	  tag_value = "bbb"
	}
}

  locals {
	dcdb_id = cloud_dcdb_instance.hourdb_instance.id
  }
`

const testAccDcdbActivateHourInstanceOperation = testAccDcdbHourInsOperation + `

resource "tencentcloudenterprise_dcdb_activate_hour_instance_operation" "activate_operation" {
  instance_id = local.dcdb_id
}

`

const testAccDcdbIsolateHourInstanceOperation = testAccDcdbHourInsOperation + `

resource "tencentcloudenterprise_dcdb_isolate_hour_instance_operation" "isolate_operation" {
	instance_id = local.dcdb_id
}

`

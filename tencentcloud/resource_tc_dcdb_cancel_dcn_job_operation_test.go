package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDcdbCancelDcnJobOperationResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDcdbCancelDcnJobOperation,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_cancel_dcn_job_operation.cancel_operation", "id"),
					resource.TestCheckResourceAttrSet("tencentcloudenterprise_dcdb_cancel_dcn_job_operation.cancel_operation", "instance_id"),
				),
			},
		},
	})
}

const testAccDcdbHourInsOperation_dcn = defaultAzVariable + CommonPresetDcdb + `
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

resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance_dcn" {
	instance_name     = "test_dcdb_db_hourdb_instance_dcn"
	zones             = [var.default_az]
	shard_memory      = "2"
	shard_storage     = "10"
	shard_node_count  = "2"
	shard_count       = "2"
	vpc_id            = local.vpc_id
	subnet_id         = local.subnet_id
	security_group_id = local.sg_id
	db_version_id     = "8.0"
	dcn_region        = "ap-guangzhou"
	dcn_instance_id   = local.dcdb_id  //master_instance
	resource_tags {
	  tag_key = "aaa"
	  tag_value = "bbb"
	}
}

// resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance_master" {
// 	instance_name = "test_dcdb_db_hourdb_instance_master"
// 	zones = [var.default_az]
// 	shard_memory = "2"
// 	shard_storage = "10"
// 	shard_node_count = "2"
// 	shard_count = "2"
// 	vpc_id = local.vpc_id
// 	subnet_id = local.subnet_id
// 	security_group_id = local.sg_id
// 	db_version_id = "8.0"
// 	resource_tags {
// 	  tag_key = "aaa"
// 	  tag_value = "bbb"
// 	}
// }

  locals {
	// master_dcdb_id = cloud_dcdb_instance.hourdb_instance_master.id
	dcn_dcdb_id = cloud_dcdb_instance.hourdb_instance_dcn.id
  }
`

const testAccDcdbCancelDcnJobOperation = testAccDcdbHourInsOperation_dcn + `

resource "tencentcloudenterprise_dcdb_cancel_dcn_job_operation" "cancel_operation" {
  instance_id = local.dcn_dcdb_id
}

`

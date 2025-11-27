---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dcdb_instances"
sidebar_current: "docs-tencentcloudenterprise-datasource-dcdb_instances"
description: |-
  Use this data source to query detailed information of dcdb instances
---

# tencentcloudenterprise_dcdb_instances

Use this data source to query detailed information of dcdb instances

## Example Usage

```hcl
data "tencentcloudenterprise_dcdb_instances" "instances1" {
  instance_ids        = ["your_dcdb_instance1_id"]
  search_name         = "instancename"
  search_key          = "search_key"
  project_ids         = [0]
  excluster_type      = 0
  is_filter_excluster = true
  excluster_type      = 0
  is_filter_vpc       = true
  vpc_id              = "your_vpc_id"
  subnet_id           = "your_subnet_id"
}

data "tencentcloudenterprise_dcdb_instances" "instances2" {
  instance_ids = ["your_dcdb_instance2_id"]
}

data "tencentcloudenterprise_dcdb_instances" "instances3" {
  search_name         = "instancename"
  search_key          = "instances3"
  is_filter_excluster = false
  excluster_type      = 2
}
```

## Argument Reference

The following arguments are supported:

* `excluster_type` - (Optional, Int) Cluster excluster type.
* `instance_ids` - (Optional, Set: [`String`]) Instance ids.
* `is_filter_excluster` - (Optional, Bool) Search according to the cluster excluter type.
* `is_filter_vpc` - (Optional, Bool) Search according to the vpc.
* `project_ids` - (Optional, Set: [`Int`]) Project ids.
* `result_output_file` - (Optional, String) Used to save results.
* `search_key` - (Optional, String) Search key, support fuzzy query.
* `search_name` - (Optional, String) Search name, support instancename, vip, all.
* `subnet_id` - (Optional, String) Subnet id, valid when IsFilterVpc is true.
* `vpc_id` - (Optional, String) Vpc id, valid when IsFilterVpc is true.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - Instance list.
  * `app_id` - App id.
  * `auto_renew_flag` - Auto renew flag.
  * `create_time` - Create time.
  * `db_engine` - Db engine.
  * `db_version` - Db engine version.
  * `instance_id` - Instance id.
  * `instance_name` - Instance name.
  * `instance_type` - Instance type.
  * `is_audit_supported` - Aduit support, 0:support, 1:unsupport.
  * `is_tmp` - Tmp instance mark.
  * `isolated_timestamp` - Isolated time.
  * `memory` - Memory, the unit is GB.
  * `node_count` - Node count.
  * `paymode` - Pay mode.
  * `period_end_time` - Expired time.
  * `project_id` - Project id.
  * `region` - Region.
  * `resource_tags` - Resource tags.
    * `tag_key` - Tag key.
    * `tag_value` - Tag value.
  * `shard_count` - Shard count.
  * `shard_detail` - Shard detail.
    * `cpu` - Cpu cores.
    * `createtime` - Shard create time.
    * `memory` - Memory.
    * `node_count` - Node count.
    * `shard_id` - Shard id.
    * `shard_instance_id` - Shard instance id.
    * `shard_serial_id` - Shard serial id.
    * `status` - Shard status.
    * `storage` - Storage.
  * `status_desc` - Status description.
  * `status` - Status.
  * `storage` - Memory, the unit is GB.
  * `subnet_id` - Subnet id.
  * `uin` - Account uin.
  * `update_time` - Update time.
  * `vip` - Vip.
  * `vpc_id` - Vpc id.
  * `vport` - Vport.
  * `wan_domain` - Wan domain.
  * `wan_port` - Wan port.
  * `wan_status` - Wan status, 0:nonactivated, 1:activated, 2:closed, 3:activating.
  * `wan_vip` - Wan vip.



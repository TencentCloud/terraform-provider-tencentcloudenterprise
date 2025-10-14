---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_shards"
sidebar_current: "docs-cloud-datasource-dcdb_shards"
description: |-
  Use this data source to query detailed information of dcdb shards
---

# cloud_dcdb_shards

Use this data source to query detailed information of dcdb shards

## Example Usage

```hcl
data "cloud_dcdb_shards" "shards" {
  instance_id        = "your_instance_id"
  shard_instance_ids = ["shard1_id"]
}

data "cloud_dcdb_shards" "shards" {
  instance_id = "your_instance_id"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) Instance id.
* `result_output_file` - (Optional, String) Used to save results.
* `shard_instance_ids` - (Optional, Set: [`String`]) Shard instance ids.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - Shard list.
  * `cpu` - Cpu cores.
  * `create_time` - Create time.
  * `instance_id` - Instance id.
  * `memory_usage` - Memory usage.
  * `memory` - Memory, the unit is GB.
  * `node_count` - Node count.
  * `paymode` - Pay mode.
  * `period_end_time` - Expired time.
  * `project_id` - Project id.
  * `proxy_version` - Proxy version.
  * `range` - The range of shard key.
  * `region` - Region.
  * `shard_instance_id` - Shard instance id.
  * `shard_master_zone` - Shard master zone.
  * `shard_serial_id` - Shard serial id.
  * `shard_slave_zones` - Shard slave zones.
  * `status_desc` - Status description.
  * `status` - Status.
  * `storage_usage` - Storage usage.
  * `storage` - Memory, the unit is GB.
  * `subnet_id` - Subnet id.
  * `vpc_id` - Vpc id.
  * `zone` - Zone.



---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dcdb_instance"
sidebar_current: "docs-tencentcloudenterprise-resources-dcdb_instance"
description: |-
  Provides a resource to create a dcdb hourdb_instance
---

# tencentcloudenterprise_dcdb_instance

Provides a resource to create a dcdb hourdb_instance

## Example Usage

```hcl
resource "tencentcloudenterprise_dcdb_instance" "hourdb_instance" {
  instance_name    = "111"
  zones            = ["yfm18", "yfm18"]
  cpu_arch         = "X86"
  ipv6_flag        = 0
  shard_memory     = 2
  shard_storage    = 10
  shard_node_count = 2
  shard_cpu        = 1
  shard_count      = 2
  vpc_id           = "vpc-cs6ffr73"
  subnet_id        = "subnet-mfbxe9zk"
  db_version_id    = "5.7"
  project_id       = "pr-bae40f73"

  resource_tags {
    tag_key   = "createdBy"
    tag_value = "terraform3"
  }

  init_params {

  }
}
```

## Argument Reference

The following arguments are supported:

* `init_params` - (Required, List) The optional values for this interface are: character_set_server (character set,required), lower_case_table_names (case sensitivity of table names, required, 0 - sensitive; 1 - insensitive), innodb_page_size (innodb data page, default 16K), sync_mode (synchronization mode: 0 - asynchronous; 1 - strong synchronization; 2 - degradable strong synchronization. Default is strong synchronization).
* `shard_count` - (Required, Int) Instance shard count.
* `shard_memory` - (Required, Int) Memory(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
* `shard_node_count` - (Required, Int) Node count for each shard. It can be obtained by querying api DescribeShardSpec.
* `shard_storage` - (Required, Int) Storage(GB) for each shard. It can be obtained by querying api DescribeShardSpec.
* `subnet_id` - (Required, String) Subnet id, its required when vpcId is set.
* `vpc_id` - (Required, String) Vpc id.
* `zones` - (Required, Set: [`String`]) Available zone.
* `cpu_arch` - (Optional, String) Cpu architecture, default to x86_64.
* `db_version_id` - (Optional, String) Db engine version, default to Percona 5.7.17.
* `dcn_instance_id` - (Optional, String) DCN source instance ID.
* `dcn_region` - (Optional, String) DCN source region.
* `extranet_access` - (Optional, Bool) Whether to open the extranet access.
* `instance_count` - (Optional, Int) Instance count, default to 1.
* `instance_name` - (Optional, String) Name of this instance.
* `ipv6_flag` - (Optional, Int) Whether to support IPv6.
* `project_id` - (Optional, String) Project id.
* `resource_tags` - (Optional, List) Resource tags.
* `security_group_id` - (Optional, String) Security group id.
* `shard_cpu` - (Optional, Int) Cpu core for each shard. It can be obtained by querying api DescribeShardSpec.
* `vip` - (Optional, String) The field is required to specify VIP.
* `vipv6` - (Optional, String) The field is required to specify VIPv6.

The `init_params` object supports the following:

* `character_set_server` - (Optional, String) The available zone ID of the instance.
* `lower_case_table_names` - (Optional, Int) (case sensitivity of table names, required, 0 - sensitive; 1 - insensitive).
* `sync_mode` - (Optional, Int) (synchronization mode: 0 - asynchronous; 1 - strong synchronization; 2 - degradable strong synchronization. Default is strong synchronization).

The `resource_tags` object supports the following:

* `tag_key` - (Required, String) Tag key.
* `tag_value` - (Required, String) Tag value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `vport` - Intranet port.


## Import

dcdb hourdb_instance can be imported using the id, e.g.
```
$ terraform import tencentcloudenterprise_dcdb_instance.hourdb_instance hourdbInstance_id
```


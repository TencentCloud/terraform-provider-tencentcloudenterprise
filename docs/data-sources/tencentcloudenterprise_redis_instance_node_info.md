---
subcategory: "Cloud Redis®(Redis)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_redis_instance_node_info"
sidebar_current: "docs-tencentcloudenterprise-datasource-redis_instance_node_info"
description: |-
  Use this data source to query detailed information of instance_node_info
---

# tencentcloudenterprise_redis_instance_node_info

Use this data source to query detailed information of instance_node_info

## Example Usage

```hcl
data "tencentcloudenterprise_redis_instance_node_info" "instance_node_info" {
  instance_id = "crs-c1nl9rpv"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) The ID of instance.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `proxy_count` - Number of proxy nodes.
* `proxy` - Proxy node information.
  * `node_id` - Node ID.
  * `zone_id` - Zone ID.
* `redis_count` - Number of redis nodes.
* `redis` - Redis node information.
  * `cluster_id` - Shard ID.
  * `node_id` - Node ID.
  * `node_role` - Node role.
  * `zone_id` - Zone ID.



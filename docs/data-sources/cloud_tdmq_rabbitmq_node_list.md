---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_rabbitmq_node_list"
sidebar_current: "docs-cloud-datasource-tdmq_rabbitmq_node_list"
description: |-
  Use this data source to query detailed information of tdmq rabbitmq_node_list
---

# cloud_tdmq_rabbitmq_node_list

Use this data source to query detailed information of tdmq rabbitmq_node_list

## Example Usage

```hcl
data "cloud_tdmq_rabbitmq_node_list" "rabbitmq_node_list" {
  instance_id = "amqp-testtesttest"
  node_name   = "keep-node"
  filters {
    name   = "nodeStatus"
    values = ["running", "down"]
  }
  sort_element = "cpuUsage"
  sort_order   = "descend"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) Rabbitmq cluster ID.
* `filters` - (Optional, List) Filter parameter name and value. Now there is only one nodeStatus: running/down. Array type, compatible with adding filter parameters later.
* `node_name` - (Optional, String) Fuzzy search node name.
* `result_output_file` - (Optional, String) Used to save results.
* `sort_element` - (Optional, String) Sort by the specified element, now there are only 2: cpuUsage/diskUsage.
* `sort_order` - (Optional, String) Sort order: ascend/descend.

The `filters` object supports the following:

* `name` - (Optional, String) The name of the filter parameter.
* `values` - (Optional, Set) Value.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `node_list` - Cluster list. Note: This field may return null, indicating that no valid value can be obtained.
  * `cpu_usage` - CPU usageNote: This field may return null, indicating that no valid value can be obtained.
  * `disk_usage` - Disk usageNote: This field may return null, indicating that no valid value can be obtained.
  * `memory` - Memory usage, in MBNote: This field may return null, indicating that no valid value can be obtained.
  * `node_name` - Node nameNote: This field may return null, indicating that no valid value can be obtained.
  * `node_status` - Node statusNote: This field may return null, indicating that no valid value can be obtained.
  * `process_number` - Number of Erlang processes for RabbitmqNote: This field may return null, indicating that no valid value can be obtained.



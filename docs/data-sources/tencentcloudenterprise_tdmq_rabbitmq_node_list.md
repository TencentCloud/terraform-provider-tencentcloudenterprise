---
subcategory: "TDMQ for RabbitMQ(trabbit)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_rabbitmq_node_list"
sidebar_current: "docs-tencentcloudenterprise-datasource-tdmq_rabbitmq_node_list"
description: |-
  Use this data source to query detailed information of tdmq rabbitmq node list
---

# tencentcloudenterprise_tdmq_rabbitmq_node_list

Use this data source to query detailed information of tdmq rabbitmq node list

## Example Usage

### ### Query all nodes in a RabbitMQ cluster

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_node_list" "all_nodes" {
  instance_id = "amqp-xxxxxxxx"
}

output "node_list" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_node_list.all_nodes.node_list
}
```

### ### Query RabbitMQ nodes by exact node name

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_node_list" "specific_node" {
  instance_id = "amqp-xxxxxxxx"
  node_name   = "rabbit@rabbitmq-broker-1.rabbitmq-broker-internal.amqp-7d39mjvn.svc.cluster.local"
}

output "node_status" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_node_list.specific_node.node_list[0].node_status
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) RabbitMQ cluster instance ID. The ID of the RabbitMQ instance to query node information from.
* `filters` - (Optional, List) Filter conditions for querying nodes. Currently supports filtering by `nodeStatus` with values: `running` (node is running), `down` (node is down). Multiple filters can be specified. Array type design allows for future filter parameter extensions.
* `node_name` - (Optional, String) Node name for fuzzy search filtering. Returns nodes whose names contain this string.
* `result_output_file` - (Optional, String) File path to save the query results. If specified, the node list will be written to this file in JSON format.
* `sort_element` - (Optional, String) Field to sort results by. Currently supported value: `diskUsage` (sort by disk usage).
* `sort_order` - (Optional, String) Sort order. Valid values: `ascend` (ascending order), `descend` (descending order).

The `filters` object supports the following:

* `name` - (Optional, String) Name of the filter parameter. Currently supported: `nodeStatus` (filter by node running status).
* `values` - (Optional, Set) Value(s) for the filter parameter. For `nodeStatus`, valid values are: `running`, `down`. Multiple values can be specified.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `node_list` - List of RabbitMQ cluster nodes matching the filter criteria. Note: This field may return null if no nodes match the query.
  * `disk_usage` - Disk usage percentage of the node. Indicates how much disk storage the node is consuming. Note: This field may return null.
  * `memory` - Memory usage in MB (megabytes). The amount of memory currently being used by the node. Note: This field may return null.
  * `node_name` - Name of the RabbitMQ node. Note: This field may return null.
  * `node_status` - Node running status. Possible values: `running` (node is operational), `down` (node is not responding). Note: This field may return null.
  * `process_number` - Number of Erlang processes running on the RabbitMQ node. RabbitMQ is built on Erlang, and this metric indicates the total number of Erlang processes active on the node. Note: This field may return null.


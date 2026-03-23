---
subcategory: "TDMQ"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_rabbitmq_vip_instance"
sidebar_current: "docs-tencentcloudenterprise-datasource-tdmq_rabbitmq_vip_instance"
description: |-
  Use this data source to query detailed information of tdmq rabbitmq_vip_instance
---

# tencentcloudenterprise_tdmq_rabbitmq_vip_instance

Use this data source to query detailed information of tdmq rabbitmq_vip_instance

## Example Usage

### ### Query all RabbitMQ VIP instances

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "all_instances" {
}

output "instance_list" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.all_instances.instances
}
```

### ### Query RabbitMQ VIP instance by instance ID

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "by_instance_id" {
  filters {
    name   = "instanceIds"
    values = ["amqp-xxxxxxxx"]
  }
}

output "instance_info" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.by_instance_id.instances[0]
}
```

### ### Query RabbitMQ VIP instance by instance name

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "by_name" {
  filters {
    name   = "instanceName"
    values = ["my-rabbitmq-cluster"]
  }
}
```

### ### Query RabbitMQ VIP instances via endpoint

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "by_endpoint" {
  filters {
    name   = "endpoint"
    values = ["10.0.0.17"]
  }
}
```

### ### Query and export to file

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "export_instances" {
  result_output_file = "./rabbitmq_instances.json"
}
```

### ### Access instance details including VPC and zone information

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "detailed_info" {
  filters {
    name   = "instanceIds"
    values = ["rmq-xxxxxxxx"]
  }
}

output "instance_zone_ids" {
  description = "Availability zones where the instance is deployed"
  value       = data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.detailed_info.instances[0].zone_ids
}

output "instance_vpcs" {
  description = "VPC access endpoints"
  value       = data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.detailed_info.instances[0].vpcs
}

output "vpc_endpoint" {
  description = "First VPC endpoint address for connection"
  value       = length(data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.detailed_info.instances[0].vpcs) > 0 ? data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.detailed_info.instances[0].vpcs[0].vpc_endpoint : ""
}
```

## Argument Reference

The following arguments are supported:

* `filters` - (Optional, List) Filter conditions for querying RabbitMQ instances. Multiple filters can be specified to narrow down the query results.
* `result_output_file` - (Optional, String) File path to save the query results. If specified, the instance list will be written to this file in JSON format.

The `filters` object supports the following:

* `name` - (Optional, String) Name of the filter parameter. Common filter names include: `instanceIds` (filter by instance IDs), `instanceName` (filter by instance name), `status` (filter by instance status).
* `values` - (Optional, Set) Value(s) for the filter parameter. Multiple values can be specified, and instances matching any of the values will be returned.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `instances` - List of RabbitMQ VIP instances matching the filter criteria.
  * `auto_renew_flag` - Auto-renewal flag. Valid values: `0` (default/manual renewal - user has not set auto-renewal), `1` (auto-renewal enabled), `2` (auto-renewal explicitly disabled by user).
  * `config_display` - Display name of the instance configuration specification (e.g., "4 cores 8GB").
  * `exception_information` - Exception or error information if the cluster is in abnormal state. Note: This field may return null when the instance is running normally.
  * `expire_time` - Instance expiration timestamp in milliseconds. The time when the prepaid instance will expire.
  * `instance_id` - Instance ID. Unique identifier for the RabbitMQ instance.
  * `instance_name` - Instance name. The user-defined name of the RabbitMQ instance.
  * `instance_version` - RabbitMQ version of the instance (e.g., 3.8.30, 3.11.8). Note: This field may return null.
  * `max_band_width` - Peak bandwidth in Mbps (megabits per second). Maximum network bandwidth capacity of the instance.
  * `max_storage` - Total storage capacity in GB (gigabytes). Maximum disk space available for the instance.
  * `max_tps` - Peak TPS (Transactions Per Second). Maximum throughput capacity of the instance.
  * `node_count` - Number of nodes in the cluster.
  * `pay_mode` - Billing mode. Valid values: `0` (postpaid/pay-as-you-go), `1` (prepaid/monthly subscription).
  * `remark` - Remark or note about the instance. Note: This field may return null.
  * `spec_name` - Instance specification name (e.g., rabbit-vip-basic-1). The internal identifier for the instance configuration.
  * `status` - Instance status. Valid values: `0` (creating), `1` (running/normal), `2` (isolating), `3` (destroyed), `4` (abnormal/failed), `5` (delivery failed).
  * `vpcs` - List of VPC access points. Contains VPC network endpoint information for accessing the RabbitMQ instance.
    * `subnet_id` - Subnet ID where the access endpoint is located.
    * `vpc_data_stream_endpoint_status` - Status of the VPC endpoint. Valid values: `OFF`, `ON`, `CREATING`, `DELETING`.
    * `vpc_endpoint` - VPC private network access endpoint address. Use this address to connect to RabbitMQ from within the VPC.
    * `vpc_id` - VPC ID where the access endpoint is located.
  * `zone_ids` - Availability zone ID list where the instance is deployed. Multiple zone IDs indicate multi-availability zone deployment.


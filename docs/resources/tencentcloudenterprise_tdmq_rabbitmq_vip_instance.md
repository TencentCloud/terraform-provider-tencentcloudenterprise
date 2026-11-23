---
subcategory: "TDMQ"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_rabbitmq_vip_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_rabbitmq_vip_instance"
description: |-
  Manages a dedicated RabbitMQ instance.
---

# tencentcloudenterprise_tdmq_rabbitmq_vip_instance

Manages a dedicated RabbitMQ instance.

## Example Usage

```hcl
variable "zone_ids" { type = set(number) }
variable "subnet_id" { type = string }
variable "vpc_id" { type = string }

resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "example" {
  cluster_name                          = "tf-example"
  zone_ids                              = var.zone_ids
  pay_mode                              = 0
  node_num                              = 3
  subnet_id                             = var.subnet_id
  vpc_id                                = var.vpc_id
  enable_create_default_ha_mirror_queue = false
}
```

## Argument Reference

The following arguments are supported:

* `cluster_name` - (Required, String) RabbitMQ cluster name. Length must be between 3-64 characters. Only letters, numbers, hyphens (-), and underscores (_) are allowed.
* `enable_create_default_ha_mirror_queue` - (Required, Bool) Whether to create a default HA (High Availability) mirrored queue. When enabled, queues will be automatically mirrored across nodes for high availability. Default is true.
* `node_num` - (Required, Int) Number of nodes in the cluster. Must be greater than 0. For single availability zone deployment, typically use 1 node; for multi-availability zone deployment, minimum 3 nodes are required for high availability. Please configure based on your actual environment and requirements.
* `subnet_id` - (Required, String) Subnet ID within the specified VPC where the instance will be deployed. Format: subnet-xxxxxxxx.
* `vpc_id` - (Required, String) VPC (Virtual Private Cloud) ID where the RabbitMQ instance will be deployed. Format: vpc-xxxxxxxx.
* `zone_ids` - (Required, Set: [`Int`]) Availability zone ID list. For single availability zone deployment, provide one zone ID; for multi-availability zone deployment, provide multiple zone IDs. Multi-availability zone instances require at least 3 nodes.
* `cluster_version` - (Optional, String) RabbitMQ cluster version. Valid values: `3.8.30` (default), `3.11.8`. Different versions may have different features and performance characteristics.
* `node_spec` - (Optional, String) Node specification. Valid values: `rabbit-vip-basic-5` (2C4G), `rabbit-vip-profession-2c8g` (2C8G), `rabbit-vip-basic-1` (4C8G, default), `rabbit-vip-profession-4c16g` (4C16G), `rabbit-vip-basic-2` (8C16G), `rabbit-vip-profession-8c32g` (8C32G), `rabbit-vip-basic-4` (16C32G), `rabbit-vip-profession-16c64g` (16C64G). Note: Some specifications may be unavailable due to stock limitations.
* `pay_mode` - (Optional, Int, ForceNew) Billing mode. Use 0 for postpaid or 1 for prepaid (default).
* `storage_size` - (Optional, Int) Storage capacity per node in GB. Default is 200GB.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `public_access_endpoint` - Public network access endpoint address. Used to access the RabbitMQ instance from the internet.
* `vpcs` - List of VPC access points. Contains VPC network endpoint information for accessing the RabbitMQ instance from within VPC.
  * `subnet_id` - Subnet ID where the access endpoint is located.
  * `vpc_data_stream_endpoint_status` - Status of the VPC endpoint. Indicates the availability status of the VPC access point.
  * `vpc_endpoint` - VPC private network access endpoint address. Use this address to connect to RabbitMQ from within the VPC.
  * `vpc_id` - VPC ID where the access endpoint is located.


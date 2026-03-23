---
subcategory: "TDMQ"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_rabbitmq_vip_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_rabbitmq_vip_instance"
description: |-
  Provides a resource to create and manage TDMQ RabbitMQ VIP instance
---

# tencentcloudenterprise_tdmq_rabbitmq_vip_instance

Provides a resource to create and manage TDMQ RabbitMQ VIP instance

## Example Usage

### ### Create a basic RabbitMQ VIP instance with single node

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "example" {
  cluster_name    = "rabbitmq-cluster"
  zone_ids        = ["ap-chongqing-1"]
  vpc_id          = "vpc-xxxxxxxx"
  subnet_id       = "subnet-xxxxxxxx"
  node_spec       = "rabbit-vip-basic-1"
  node_num        = 3
  storage_size    = 200
  cluster_version = "3.8.30"
}
```

### ### Create a high-availability RabbitMQ VIP instance with multi-zone deployment

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "ha_instance" {
  cluster_name                          = "rabbitmq-ha-cluster"
  zone_ids                              = ["ap-chongqing-1", "ap-chongqing-2", "ap-chongqing-3"]
  vpc_id                                = "vpc-xxxxxxxx"
  subnet_id                             = "subnet-xxxxxxxx"
  node_spec                             = "rabbit-vip-basic-2"
  node_num                              = 3
  storage_size                          = 500
  enable_create_default_ha_mirror_queue = true
  cluster_version                       = "3.11.8"
}
```

### ### Create a production RabbitMQ VIP instance with enhanced resources

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "production" {
  cluster_name    = "rabbitmq-prod"
  zone_ids        = ["ap-chongqing-1"]
  vpc_id          = "vpc-xxxxxxxx"
  subnet_id       = "subnet-xxxxxxxx"
  node_spec       = "rabbit-vip-basic-4"
  node_num        = 3
  storage_size    = 1000
  cluster_version = "3.11.8"
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


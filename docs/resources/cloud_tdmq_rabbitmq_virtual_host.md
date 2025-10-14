---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_rabbitmq_virtual_host"
sidebar_current: "docs-cloud-resources-tdmq_rabbitmq_virtual_host"
description: |-
  Provides a resource to create a tdmq rabbitmq_virtual_host
---

# cloud_tdmq_rabbitmq_virtual_host

Provides a resource to create a tdmq rabbitmq_virtual_host

## Example Usage

```hcl
data "cloud_availability_zones" "zones" {
  name = "ap-guangzhou-6"
}

# create vpc
resource "cloud_vpc" "vpc" {
  name       = "vpc"
  cidr_block = "10.0.0.0/16"
}

# create vpc subnet
resource "cloud_vpc_subnet" "subnet" {
  name              = "subnet"
  vpc_id            = cloud_vpc.vpc.id
  availability_zone = "ap-guangzhou-6"
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

# create rabbitmq instance
resource "cloud_tdmq_rabbitmq_vip_instance" "example" {
  zone_ids                              = [data.cloud_availability_zones.zones.zones.0.id]
  vpc_id                                = cloud_vpc.vpc.id
  subnet_id                             = cloud_vpc_subnet.subnet.id
  cluster_name                          = "tf-example-rabbitmq-vip-instance"
  node_spec                             = "rabbit-vip-basic-1"
  node_num                              = 1
  storage_size                          = 200
  enable_create_default_ha_mirror_queue = false
  auto_renew_flag                       = true
  time_span                             = 1
}

# create virtual host
resource "cloud_tdmq_rabbitmq_virtual_host" "example" {
  instance_id  = cloud_tdmq_rabbitmq_vip_instance.example.id
  virtual_host = "tf-example-vhost"
  description  = "desc."
  trace_flag   = true
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) Cluster instance ID.
* `virtual_host` - (Required, String) vhost name.
* `description` - (Optional, String) describe.
* `trace_flag` - (Optional, Bool) Message track switch, true is on, false is off, default is off.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

tdmq rabbitmq_virtual_host can be imported using the id, e.g.

```
terraform import cloud_tdmq_rabbitmq_virtual_host.example amqp-pbavw2wd#tf-example-vhost
```


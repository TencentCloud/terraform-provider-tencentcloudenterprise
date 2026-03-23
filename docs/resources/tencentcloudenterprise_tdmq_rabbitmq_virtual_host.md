---
subcategory: "TDMQ for RabbitMQ(trabbit)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_rabbitmq_virtual_host"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_rabbitmq_virtual_host"
description: |-
  Provides a resource to create and manage TDMQ RabbitMQ virtual host
---

# tencentcloudenterprise_tdmq_rabbitmq_virtual_host

Provides a resource to create and manage TDMQ RabbitMQ virtual host

## Example Usage

### ### Create a basic virtual host

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_virtual_host" "example" {
  instance_id  = "amqp-xxxxxxxx"
  virtual_host = "my_vhost"
  description  = "Virtual host for application 1"
}
```

### ### Create a virtual host with mirror queue policy

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_virtual_host" "with_mirror" {
  instance_id              = "amqp-xxxxxxxx"
  virtual_host             = "mirror_vhost"
  description              = "Virtual host with mirror queue policy enabled"
  mirror_queue_policy_flag = true
}
```

### ### Create a virtual host without mirror queue policy

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_virtual_host" "without_mirror" {
  instance_id              = "amqp-xxxxxxxx"
  virtual_host             = "no_mirror_vhost"
  description              = "Virtual host without mirror queue policy"
  mirror_queue_policy_flag = false
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) RabbitMQ cluster instance ID. The ID of the RabbitMQ instance where the virtual host will be created.
* `virtual_host` - (Required, String) Virtual host (vhost) name. Virtual hosts provide logical grouping and separation of resources (exchanges, queues, bindings) within a RabbitMQ instance, allowing multiple applications to share the same RabbitMQ instance securely.
* `description` - (Optional, String) Description for the virtual host. Provides additional information about the vhost's purpose or usage.
* `mirror_queue_policy_flag` - (Optional, Bool, ForceNew) Whether to create a mirror queue policy. When enabled (`true`), a mirror queue policy will be automatically created to replicate queues across cluster nodes for high availability. When disabled (`false`), no mirror queue policy is created. Default is `true` (enabled). Note: This can only be set during virtual host creation and cannot be modified afterwards.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `trace_flag` - Message tracing switch status (read-only).


---
subcategory: "TDMQ for RabbitMQ(trabbit)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_rabbitmq_user"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_rabbitmq_user"
description: |-
  Provides a resource to create and manage TDMQ RabbitMQ user
---

# tencentcloudenterprise_tdmq_rabbitmq_user

Provides a resource to create and manage TDMQ RabbitMQ user

## Example Usage

### ### Create a RabbitMQ user with administrator role

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "example" {
  instance_id = "amqp-xxxxxxxx"
  user        = "admin_user"
  password    = "AdminPassword123!"
  description = "Administrator user for RabbitMQ"
  tags        = ["administrator"]
}
```

### ### Create a RabbitMQ user with monitoring role

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "monitoring" {
  instance_id = "amqp-xxxxxxxx"
  user        = "monitor_user"
  password    = "MonitorPass123!"
  description = "Monitoring user for RabbitMQ"
  tags        = ["monitoring"]
}
```

### ### Create a RabbitMQ user with management role

```hcl
resource "tencentcloudenterprise_tdmq_rabbitmq_user" "management" {
  instance_id = "amqp-xxxxxxxx"
  user        = "mgmt_user"
  password    = "ManagementPass123!"
  description = "Management user for RabbitMQ"
  tags        = ["management"]
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) RabbitMQ cluster instance ID. The ID of the RabbitMQ instance where the user will be created.
* `password` - (Required, String) Password for RabbitMQ authentication. This will be used to log in to the RabbitMQ server and management console. The password is stored securely and will not be displayed in logs.
* `tags` - (Required, List: [`String`]) User tags. Determines the user's access permissions to RabbitMQ Management console. Valid values: `administrator` (full admin access, default), `monitoring` (read-only monitoring access), `policymaker` (can manage policies), `management` (can manage resources), `none` (no management console access).
* `user` - (Required, String) Username for RabbitMQ authentication. This will be used to log in to the RabbitMQ server and management console.
* `description` - (Optional, String) Description for the user. Provides additional information about the user's purpose or role.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



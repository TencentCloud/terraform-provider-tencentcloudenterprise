---
subcategory: "TDMQ"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_environment_role_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_pulsar_environment_role_attachment"
description: |-
  Provides a resource to attach a TDMQ Pulsar role to a namespace with specific permissions.
---

# tencentcloudenterprise_tdmq_pulsar_environment_role_attachment

Provides a resource to attach a TDMQ Pulsar role to a namespace with specific permissions.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_environment_role_attachment" "example" {
  cluster_id  = "pulsar-xxxxxxxx"
  environ_id  = "my-namespace"
  role_name   = "my-role"
  permissions = ["produce", "consume"]
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) The id of tdmq cluster.
* `environ_id` - (Required, String) The name of tdmq namespace.
* `permissions` - (Required, List: [`String`]) The permissions of tdmq role.
* `role_name` - (Required, String) The name of tdmq role.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of resource.

## Import

tencentcloudenterprise_tdmq_pulsar_environment_role_attachment can be imported using the id, e.g.

```
# TDMQ Pulsar namespace role attachment can be imported using the id, e.g

```
$ terraform import tencentcloudenterprise_tdmq_pulsar_environment_role_attachment.example pulsar-xxxxxxxx#my-namespace#my-role
```
```


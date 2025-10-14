---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_namespace_role_attachment"
sidebar_current: "docs-cloud-resources-tdmq_namespace_role_attachment"
description: |-
  Provide a resource to create a TDMQ role.
---

# cloud_tdmq_namespace_role_attachment

Provide a resource to create a TDMQ role.

## Example Usage

```hcl
resource "cloud_tdmq_instance" "foo" {
  cluster_name = "example"
  remark       = "this is description."
}

resource "cloud_tdmq_namespace" "bar" {
  environ_name = "example"
  msg_ttl      = 300
  cluster_id   = cloud_tdmq_instance.foo.id
  remark       = "this is description."
}

resource "cloud_tdmq_topic" "bar" {
  environ_id = cloud_tdmq_namespace.bar.id
  topic_name = "example"
  partitions = 6
  topic_type = 0
  cluster_id = cloud_tdmq_instance.foo.id
  remark     = "this is description."
}

resource "cloud_tdmq_role" "bar" {
  role_name  = "example"
  cluster_id = cloud_tdmq_instance.foo.id
  remark     = "this is description world"
}

resource "cloud_tdmq_namespace_role_attachment" "bar" {
  environ_id  = cloud_tdmq_namespace.bar.id
  role_name   = cloud_tdmq_role.bar.role_name
  permissions = ["produce", "consume"]
  cluster_id  = cloud_tdmq_instance.foo.id
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



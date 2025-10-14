---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_role"
sidebar_current: "docs-cloud-resources-tdmq_role"
description: |-
  Provide a resource to create a TDMQ role.
---

# cloud_tdmq_role

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
  cluster_id   = "cloud_tdmq_instance.foo.id"
  remark       = "this is description."
}

resource "cloud_tdmq_topic" "bar" {
  environ_id = "cloud_tdmq_namespace.bar.id"
  topic_name = "example"
  partitions = 6
  topic_type = 0
  cluster_id = "cloud_tdmq_instance.foo.id"
  remark     = "this is description."
}

resource "cloud_tdmq_role" "bar" {
  role_name  = "example"
  cluster_id = cloud_tdmq_instance.foo.id
  remark     = "this is description world"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) The id of tdmq cluster.
* `remark` - (Required, String) The description of tdmq role.
* `role_name` - (Required, String) The name of tdmq role.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Tdmq instance can be imported, e.g.

```
$ terraform import cloud_tdmq_instance.test tdmq_id
```


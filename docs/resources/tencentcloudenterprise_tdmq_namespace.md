---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_namespace"
sidebar_current: "docs-tencentcloudenterprise-resources-tdmq_namespace"
description: |-
  Provide a resource to create a tdmq namespace.
---

# tencentcloudenterprise_tdmq_namespace

Provide a resource to create a tdmq namespace.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_instance" "foo" {
  cluster_name = "example111"
  remark       = "this is description111."
  tags = {
    "createdBy" = "terraform"
    "test"      = "111"
  }
  bind_cluster_id   = 0
  bind_cluster_name = "default"
}

resource "tencentcloudenterprise_tdmq_namespace" "bar" {
  environ_name = "example"
  msg_ttl      = 200
  cluster_id   = tencentcloudenterprise_tdmq_instance.foo.id
  remark       = "this is description.22222222"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) The Dedicated Cluster Id.
* `environ_name` - (Required, String) The name of namespace to be created.
* `msg_ttl` - (Required, Int) The expiration time of unconsumed message.
* `remark` - (Optional, String) Description of the namespace.
* `retention_policy` - (Optional, Map) The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `time_in_minutes`: the time of message to retain; `size_in_mb`: the size of message to retain.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Tdmq namespace can be imported, e.g.

```
$ terraform import tencentcloudenterprise_tdmq_instance.test namespace_id
```


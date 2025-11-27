---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_environment"
sidebar_current: "docs-tencentcloudenterprise-resources-tdmq_pulsar_environment"
description: |-
  Provide a resource to create a TDMQ namespace.
---

# tencentcloudenterprise_tdmq_pulsar_environment

Provide a resource to create a TDMQ namespace.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_cluster" "example" {
  cluster_name = "tf_example"
  remark       = "remark."
  tags = {
    "createdBy" = "terraform"
  }
}

resource "tencentcloudenterprise_tdmq_pulsar_environment" "example" {
  environ_name = "tf_example"
  msg_ttl      = 300
  cluster_id   = tencentcloudenterprise_tdmq_pulsar_cluster.example.id
  retention_policy {
    time_in_minutes = 60
    size_in_mb      = 10
  }
  remark = "remark."
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) The Dedicated Cluster Id.
* `environ_name` - (Required, String) The name of namespace to be created.
* `msg_ttl` - (Required, Int) The expiration time of unconsumed message.
* `remark` - (Optional, String) Description of the namespace.
* `retention_policy` - (Optional, List) The Policy of message to retain. Format like: `{time_in_minutes: Int, size_in_mb: Int}`. `time_in_minutes`: the time of message to retain; `size_in_mb`: the size of message to retain.

The `retention_policy` object supports the following:

* `size_in_mb` - (Optional, Int) the size of message to retain.
* `time_in_minutes` - (Optional, Int) the time of message to retain.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Tdmq namespace can be imported, e.g.

```
$ terraform import tencentcloudenterprise_tdmq_pulsar_cluster.example tf_example#pulsar-78bwjaj8epxv
```


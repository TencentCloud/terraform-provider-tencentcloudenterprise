---
subcategory: "TDMQ"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_environment"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_pulsar_environment"
description: |-
  Provides a resource to create a TDMQ Pulsar namespace (environment).
---

# tencentcloudenterprise_tdmq_pulsar_environment

Provides a resource to create a TDMQ Pulsar namespace (environment).

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_environment" "example" {
  environ_name = "my-namespace"
  cluster_id   = "pulsar-xxxxxxxx"
  msg_ttl      = 86400
  remark       = "Example namespace"

  retention_policy {
    time_in_minutes = 1440
    size_in_mb      = 1024
  }
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

tencentcloudenterprise_tdmq_pulsar_environment can be imported using the id, e.g.

```
# TDMQ Pulsar namespace can be imported using the id, e.g

```
$ terraform import tencentcloudenterprise_tdmq_pulsar_environment.example pulsar-xxxxxxxx#my-namespace
```
```


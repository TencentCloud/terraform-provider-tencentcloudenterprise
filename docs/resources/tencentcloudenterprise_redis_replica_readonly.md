---
subcategory: "Cloud Redis®(Redis)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_redis_replica_readonly"
sidebar_current: "docs-tencentcloudenterprise-resources-redis_replica_readonly"
description: |-
  Provides a resource to create a replica_readonly
---

# tencentcloudenterprise_redis_replica_readonly

Provides a resource to create a replica_readonly

## Example Usage

```hcl
resource "tencentcloudenterprise_redis_replica_readonly" "replica_readonly" {
  instance_id = "crs-c1nl9rpv"
  operate     = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) The ID of instance.
* `operate` - (Required, String) The replica is read-only, `enable` - enable read-write splitting, `disable`- disable read-write splitting.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




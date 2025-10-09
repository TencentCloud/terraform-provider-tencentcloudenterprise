---
subcategory: "Cloud Redis®(Redis)"
layout: "cloud"
page_title: "TencentCloud: cloud_redis_clear_instance_operation"
sidebar_current: "docs-cloud-resource-redis_clear_instance_operation"
description: |-
  Provides a resource to create a clear_instance_operation
---

# cloud_redis_clear_instance_operation

Provides a resource to create a clear_instance_operation

## Example Usage

```hcl
resource "cloud_redis_clear_instance_operation" "clear_instance_operation" {
  instance_id = "crs-c1nl9rpv"
  password    = ""
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) The ID of instance.
* `password` - (Optional, String, ForceNew) Redis® instance password (password-free instances do not need to pass passwords, non-password-free instances must be transmitted).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




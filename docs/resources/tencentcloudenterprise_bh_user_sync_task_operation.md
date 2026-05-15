---
subcategory: "BH"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bh_user_sync_task_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-bh_user_sync_task_operation"
description: |-
  Provide a resource to create a BH user sync task operation
---

# tencentcloudenterprise_bh_user_sync_task_operation

Provide a resource to create a BH user sync task operation

## Example Usage

```hcl
resource "tencentcloudenterprise_bh_user_sync_task_operation" "example" {
  user_kind = 1
}
```

## Argument Reference

The following arguments are supported:

* `user_kind` - (Required, Int, ForceNew) Synchronized user type, 1-synchronize IOA users.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



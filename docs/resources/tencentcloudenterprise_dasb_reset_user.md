---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_reset_user"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_reset_user"
description: |-
  Provide a resource to create a DASB reset user operation
---

# tencentcloudenterprise_dasb_reset_user

Provide a resource to create a DASB reset user operation

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_reset_user" "example" {
  user_id = 1
}
```

## Argument Reference

The following arguments are supported:

* `user_id` - (Required, Int, ForceNew) User Id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



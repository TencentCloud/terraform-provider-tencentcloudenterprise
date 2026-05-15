---
subcategory: "BH"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bh_user_group"
sidebar_current: "docs-tencentcloudenterprise-resource-bh_user_group"
description: |-
  Provide a resource to create a BH user group
---

# tencentcloudenterprise_bh_user_group

Provide a resource to create a BH user group

## Example Usage

```hcl
resource "tencentcloudenterprise_bh_user_group" "example" {
  name = "example-user-group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) User group name, maximum length 32 characters.
* `department_id` - (Optional, String) Department ID to which the user group belongs, e.g.: 1.2.3.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `user_group_id` - User group ID.

## Import

tencentcloudenterprise_bh_user_group can be imported using the id, e.g.

```
BH user group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_user_group.example 12345
```
```


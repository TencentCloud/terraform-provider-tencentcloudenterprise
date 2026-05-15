---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_user_group"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_user_group"
description: |-
  Provide a resource to create a DASB user group
---

# tencentcloudenterprise_dasb_user_group

Provide a resource to create a DASB user group

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_user_group" "example" {
  name = "example-user-group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) User group name, maximum length 32 characters.
* `department_id` - (Optional, String) ID of the department to which the user group belongs, such as: 1.2.3.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_dasb_user_group can be imported using the id, e.g.

```
DASB user group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_user_group.example 12345
```
```


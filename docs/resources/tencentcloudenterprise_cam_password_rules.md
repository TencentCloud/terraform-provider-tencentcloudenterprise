---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_password_rules"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_password_rules"
description: |-
  Provides a resource to manage CAM password rules.
---

# tencentcloudenterprise_cam_password_rules

Provides a resource to manage CAM password rules.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_password_rules" "foo" {
  minimum_length                = 12
  must_contain                  = "Aa1!"
  force_password_change         = 90
  reuse_password_limit          = 3
  retry_password_limit          = 5
  only_admin_can_reset_password = false
  must_not_contain_username     = true
}
```

## Argument Reference

The following arguments are supported:

* `black_list` - (Optional, String) JSON array string of disallowed password keywords.
* `force_password_change` - (Optional, Int) Days after which a password must be changed.
* `minimum_length` - (Optional, Int) Minimum password length.
* `must_contain` - (Optional, String) Character classes that must be contained in the password (for example: Aa).
* `must_not_contain_username` - (Optional, Bool) Whether passwords must not contain the username.
* `only_admin_can_reset_password` - (Optional, Bool) Whether only administrators can reset passwords.
* `retry_password_limit` - (Optional, Int) Maximum retry attempts before lockout.
* `reuse_password_limit` - (Optional, Int) Number of previous passwords that cannot be reused.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



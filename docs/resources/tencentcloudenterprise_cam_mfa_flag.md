---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_mfa_flag"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_mfa_flag"
description: |-
  Provides a resource to manage CAM MFA (Multi-Factor Authentication) protection flags.
---

# tencentcloudenterprise_cam_mfa_flag

Provides a resource to manage CAM MFA (Multi-Factor Authentication) protection flags.

This resource allows you to enable or disable MFA protection for login and sensitive operations.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_mfa_flag" "example" {
  op_uin = 100000000001

  login_flag {
    phone  = 1 # Enable phone verification for login
    stoken = 1 # Enable soft token for login
    token  = 0 # Disable hard token for login
    ukey   = 0 # Disable Ukey for login
  }

  action_flag {
    phone  = 1 # Enable phone verification for sensitive operations
    stoken = 0 # Disable soft token for sensitive operations
    token  = 0 # Disable hard token for sensitive operations
    ukey   = 0 # Disable Ukey for sensitive operations
  }
}
```

## Argument Reference

The following arguments are supported:

* `op_uin` - (Required, Int, ForceNew) The UIN of the user to operate on.
* `action_flag` - (Optional, List) MFA protection settings for sensitive operations. Each field accepts 0 (disabled) or 1 (enabled).
* `login_flag` - (Optional, List) MFA protection settings for login. Each field accepts 0 (disabled) or 1 (enabled).

The `action_flag` object supports the following:

* `phone` - (Optional, Int) Phone verification for sensitive operations. 0: disabled, 1: enabled.
* `stoken` - (Optional, Int) Soft token verification for sensitive operations. 0: disabled, 1: enabled.
* `token` - (Optional, Int) Hardware token verification for sensitive operations. 0: disabled, 1: enabled.
* `ukey` - (Optional, Int) Ukey verification for sensitive operations. 0: disabled, 1: enabled.

The `login_flag` object supports the following:

* `phone` - (Optional, Int) Phone verification for login. 0: disabled, 1: enabled.
* `stoken` - (Optional, Int) Soft token verification for login. 0: disabled, 1: enabled.
* `token` - (Optional, Int) Hardware token verification for login. 0: disabled, 1: enabled.
* `ukey` - (Optional, Int) Ukey verification for login. 0: disabled, 1: enabled.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cam_mfa_flag can be imported using the id, e.g.

```
CAM MFA flag can be imported using the uin, e.g.

```
$ terraform import tencentcloudenterprise_cam_mfa_flag.example 100000000001
```
```


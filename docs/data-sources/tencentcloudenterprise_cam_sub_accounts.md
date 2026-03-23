---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_sub_accounts"
sidebar_current: "docs-tencentcloudenterprise-datasource-cam_sub_accounts"
description: |-
  Use this data source to query CAM sub accounts.
---

# tencentcloudenterprise_cam_sub_accounts

Use this data source to query CAM sub accounts.

## Example Usage

```hcl
data "tencentcloudenterprise_cam_sub_accounts" "foo" {
  filter_sub_account_uin = [100000000001]
}
```

## Argument Reference

The following arguments are supported:

* `filter_sub_account_uin` - (Required, Set: [`Int`]) List of sub-user UINs. Up to 50 UINs are supported.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `sub_accounts` - Sub-user list.
  * `create_time` - Creation time.
  * `name` - Sub-user name.
  * `remark` - Sub-user remarks.
  * `uid` - Sub-user UID. UID is the unique identifier of a user who is a message recipient, while UIN is a unique identifier of a user.
  * `uin` - Sub-user ID.
  * `user_type` - User type (1: root account; 2: sub-user; 3: WeCom sub-user; 4: collaborator; 5: message recipient).


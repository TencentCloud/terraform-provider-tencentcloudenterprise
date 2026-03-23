---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_group_user_account"
sidebar_current: "docs-tencentcloudenterprise-datasource-cam_group_user_account"
description: |-
  Use this data source to query the list of user groups that a sub-user has joined.
---

# tencentcloudenterprise_cam_group_user_account

Use this data source to query the list of user groups that a sub-user has joined.

## Example Usage

```hcl
# Query by uid

data "tencentcloudenterprise_cam_group_user_account" "example" {
  uid = 4364
}

# Query by uin

data "tencentcloudenterprise_cam_group_user_account" "example" {
  uin = 110000001037
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.
* `uid` - (Optional, Int) Sub-user uid. Conflicts with uin.
* `uin` - (Optional, Int) Sub-user uin. Conflicts with uid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `group_info` - User group information.
  * `create_time` - Create time.
  * `group_id` - User group ID.
  * `group_name` - User group name.
  * `remark` - Remark.
* `total_num` - The total number of user groups the sub-user has joined.


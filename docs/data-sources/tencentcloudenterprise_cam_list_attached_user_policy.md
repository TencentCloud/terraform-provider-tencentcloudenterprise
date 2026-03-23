---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_list_attached_user_policy"
sidebar_current: "docs-tencentcloudenterprise-datasource-cam_list_attached_user_policy"
description: |-
  Use this data source to query policies attached to a CAM user.
---

# tencentcloudenterprise_cam_list_attached_user_policy

Use this data source to query policies attached to a CAM user.

## Example Usage

```hcl
data "tencentcloudenterprise_cam_list_attached_user_policy" "foo" {
  target_uin  = 100000000001
  attach_type = 0
}
```

## Argument Reference

The following arguments are supported:

* `attach_type` - (Required, Int) 0: Return direct association and group association policies, 1: Only return direct association policies, 2: Only return group association policies.
* `target_uin` - (Required, Int) Target User ID.
* `keyword` - (Optional, String) Search Keywords.
* `result_output_file` - (Optional, String) Used to save results.
* `strategy_type` - (Optional, Int) Policy type.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `policy_list` - Policy List Data.
  * `add_time` - Creation time.
  * `create_mode` - Creation mode (1 represents policies created by product or project permissions, others represent policies created by policy syntax).
  * `description` - Policy Description.
  * `groups` - Associated information with group.
    * `group_id` - Group ID.
    * `group_name` - Group Name.
  * `policy_id` - Policy ID.
  * `policy_name` - Policy Name.
  * `strategy_type` - Policy type (1 represents custom policy, 2 represents preset policy).


---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_group_policy_attachments"
sidebar_current: "docs-tencentcloudenterprise-datasource-cam_group_policy_attachments"
description: |-
  Use this data source to query detailed information of CAM group policy attachments
---

# tencentcloudenterprise_cam_group_policy_attachments

Use this data source to query detailed information of CAM group policy attachments

## Example Usage

```hcl
# query by group_id

data "tencentcloudenterprise_cam_group_policy_attachments" "foo" {
  group_id = tencentcloudenterprise_cam_group.foo.id
}

# query by group_id and policy_id

data "tencentcloudenterprise_cam_group_policy_attachments" "bar" {
  group_id  = tencentcloudenterprise_cam_group.foo.id
  policy_id = tencentcloudenterprise_cam_policy.foo.id
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required, String) ID of the attached CAM group to be queried.
* `create_mode` - (Optional, Int) Mode of creation of the CAM user policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
* `policy_id` - (Optional, String) ID of CAM policy to be queried.
* `policy_type` - (Optional, String) Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `group_policy_attachment_list` - A list of CAM group policy attachments. Each element contains the following attributes:
  * `create_mode` - Mode of Creation of the CAM group policy attachment. 1 means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
  * `create_time` - Create time of the CAM group policy attachment.
  * `group_id` - ID of CAM group.
  * `policy_id` - Name of CAM group.
  * `policy_name` - Name of the policy.
  * `policy_type` - Type of the policy strategy. 'User' means customer strategy and 'QCS' means preset strategy.


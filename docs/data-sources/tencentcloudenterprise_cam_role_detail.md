---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_role_detail"
sidebar_current: "docs-tencentcloudenterprise-datasource-cam_role_detail"
description: |-
  Use this data source to query CAM role details.
---

# tencentcloudenterprise_cam_role_detail

Use this data source to query CAM role details.

## Example Usage

```hcl
data "tencentcloudenterprise_cam_role_detail" "foo" {
  role_name = "test-role"
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.
* `role_id` - (Optional, String) Role ID, used to specify role. Input either `role_id` or `role_name`.
* `role_name` - (Optional, String) Role name, used to specify role. Input either `role_id` or `role_name`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `role_info` - Role details.
  * `add_time` - Time role created.
  * `console_login` - If login is allowed for the role.
  * `deletion_task_id` - Task identifier for deleting a service-linked role.
  * `description` - Role description.
  * `policy_document` - Role policy document.
  * `role_id` - Role ID.
  * `role_name` - Role name.
  * `role_type` - User role. Valid values: `user`, `system`, `service_linked`.
  * `session_duration` - Valid period.
  * `update_time` - Time role last updated.


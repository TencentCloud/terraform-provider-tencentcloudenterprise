---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_list_entities_for_policy"
sidebar_current: "docs-tencentcloudenterprise-datasource-cam_list_entities_for_policy"
description: |-
  Use this data source to query entities associated with a CAM policy.
---

# tencentcloudenterprise_cam_list_entities_for_policy

Use this data source to query entities associated with a CAM policy.

## Example Usage

```hcl
data "tencentcloudenterprise_cam_list_entities_for_policy" "foo" {
  policy_id     = 12345678
  entity_filter = "User"
}
```

## Argument Reference

The following arguments are supported:

* `policy_id` - (Required, Int) Policy Id.
* `entity_filter` - (Optional, String) Can take values of 'All', 'User', 'Group', and 'Role'. 'All' represents obtaining all entity types, 'User' represents only obtaining sub accounts, 'Group' represents only obtaining user groups, and 'Role' represents only obtaining roles. The default value is 'All'.
* `result_output_file` - (Optional, String) Used to save results.
* `rp` - (Optional, Int) Per page size, default value is 20.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - Entity List.
  * `id` - Entity ID.
  * `name` - Entity Name.
  * `related_type` - Association type. 1. User association; 2 User Group Association.
  * `uin` - Entity Uin.


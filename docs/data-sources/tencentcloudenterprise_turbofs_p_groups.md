---
subcategory: "Parallel File Storage(TurboFS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_turbofs_p_groups"
sidebar_current: "docs-tencentcloudenterprise-datasource-turbofs_p_groups"
description: |-
  Use this data source to query the detail information of Turbofs permission group.
---

# tencentcloudenterprise_turbofs_p_groups

Use this data source to query the detail information of Turbofs permission group.

## Example Usage

```hcl
data "tencentcloudenterprise_turbofs_p_groups" "pgroup" {
  p_group_id = "pgroup-7nx89k7l"
}
```

## Argument Reference

The following arguments are supported:

* `p_group_id` - (Optional, String) A specified permission group ID used to query.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `p_group_list` - An information list of TurboFS permission group.
  * `bind_cfs_num` - The number of file systems associated with the permission group.
  * `c_date` - Creation datetime of the permission group.
  * `desc_info` - Description info of the permission group.
  * `name` - Name of the permission group.
  * `p_group_id` - ID of the permission group.



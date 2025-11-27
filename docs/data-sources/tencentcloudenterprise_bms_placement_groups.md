---
subcategory: "Bare Metal Server(BMS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bms_placement_groups"
sidebar_current: "docs-tencentcloudenterprise-datasource-bms_placement_groups"
description: |-
  Use this data source to query bms placement group.
---

# tencentcloudenterprise_bms_placement_groups

Use this data source to query bms placement group.

## Example Usage

```hcl
data "tencentcloudenterprise_bms_placement_groups" "bms" {
  result_output_file = "bms.json"
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Optional, String) The ID of the placement group.
* `result_output_file` - (Optional, String) The file path to output the result.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - The list of placement group.
  * `create_time` - Create time of the instance.
  * `current_num` - The current num of the instance.
  * `group_id` - The group ID of the instance.
  * `name` - The name of the instance.
  * `type` - The type of the instance.
  * `update_time` - Update time of the instance.



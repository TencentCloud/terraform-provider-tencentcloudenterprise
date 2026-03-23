---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_groups"
sidebar_current: "docs-tencentcloudenterprise-datasource-cic_groups"
description: |-
  Use this data source to query detailed information of cic groups
---

# tencentcloudenterprise_cic_groups

Use this data source to query detailed information of cic groups

## Example Usage

```hcl
data "tencentcloudenterprise_cic_groups" "groups" {
  zone_id = "z-xxxxxxxxxx"
}

output "groups_list" {
  value = data.tencentcloudenterprise_cic_groups.groups.groups
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required, String) Space ID.
* `filter_users` - (Optional, Set: [`String`]) Filtered user list. IsSelected=true will be returned for the user group associated with this user.
* `filter` - (Optional, String) Filter criterion. Format: <Attribute> <Operator> <Value>, case-insensitive. Currently, <Attribute> supports only GroupName, and <Operator> supports only eq (Equals) and sw (Start With).
* `group_type` - (Optional, String) User group type. Manual: manually created; Synchronized: externally imported.
* `result_output_file` - (Optional, String) Used to save results.
* `sort_field` - (Optional, String) Sorting field, which currently only supports CreateTime.
* `sort_type` - (Optional, String) Sorting type. Desc: descending order; Asc: ascending order. It should be set along with SortField.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `groups` - User group list.
  * `create_time` - Creation time of the user group.
  * `description` - User group description.
  * `group_id` - User group ID.
  * `group_name` - User group name.
  * `group_type` - User group type. Manual: manually created; Synchronized: externally imported.
  * `is_selected` - If the input parameter FilterUsers is provided, return true when the user is in the user group; otherwise, return false.
  * `member_count` - Number of group members.
  * `update_time` - Modification time of the user group.


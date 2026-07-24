---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_users"
sidebar_current: "docs-tencentcloudenterprise-datasource-cic_users"
description: |-
  Use this data source to query detailed information of cic users
---

# tencentcloudenterprise_cic_users

Use this data source to query detailed information of cic users

## Example Usage

```hcl
data "tencentcloudenterprise_cic_users" "users" {
  zone_id = "z-xxxxxxxxxx"
}

output "users_list" {
  value = data.tencentcloudenterprise_cic_users.users.users
}
```

### Query with filters

```hcl
data "tencentcloudenterprise_cic_users" "users" {
  zone_id     = "z-xxxxxxxxxx"
  user_name   = "test_user"
  user_status = "Enabled"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required, String) Space ID.
* `description` - (Optional, String) User description for exact match filtering.
* `email` - (Optional, String) Email address for exact match filtering.
* `result_output_file` - (Optional, String) Used to save results.
* `user_id` - (Optional, String) User ID for exact match filtering.
* `user_name` - (Optional, String) User name for exact match filtering.
* `user_status` - (Optional, String) User status. Enabled: enabled; Disabled: disabled.
* `user_type` - (Optional, String) User type. Manual: manually created; Synchronized: externally imported.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `users` - User list.
  * `create_time` - Creation time.
  * `description` - Description.
  * `display_name` - Display name.
  * `email` - Email address.
  * `first_name` - First name.
  * `is_selected` - If the input parameter FilterGroups is provided, return true when the user is in the user group; otherwise, return false.
  * `last_name` - Last name.
  * `update_time` - Update time.
  * `user_id` - User ID.
  * `user_name` - User name.
  * `user_status` - User status. Enabled: enabled; Disabled: disabled.
  * `user_type` - User type. Manual: manually created; Synchronized: externally imported.


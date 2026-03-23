---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_role_configurations"
sidebar_current: "docs-tencentcloudenterprise-datasource-cic_role_configurations"
description: |-
  Use this data source to query detailed information of cic role configurations
---

# tencentcloudenterprise_cic_role_configurations

Use this data source to query detailed information of cic role configurations

## Example Usage

```hcl
data "tencentcloudenterprise_cic_role_configurations" "role_configurations" {
  zone_id = "z-xxxxxxxxxx"
}

output "role_configurations_list" {
  value = data.tencentcloudenterprise_cic_role_configurations.role_configurations.role_configurations
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required, String) Space ID.
* `filter_targets` - (Optional, Set: [`Int`]) Check whether the member account has been configured with permissions. If configured, return IsSelected: true; otherwise, return false.
* `filter` - (Optional, String) Filter criteria, case insensitive. Currently supports RoleConfigurationName and Description.
* `principal_id` - (Optional, String) UserId of the authorized user or GroupId of the authorized user group, which must be set together with the input parameter FilterTargets.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `role_configurations` - Permission configuration list.
  * `create_time` - Creation time of the permission configuration.
  * `description` - Permission configuration description.
  * `is_selected` - If the input parameter FilterTargets is provided, check whether the member account has been configured with permissions. If configured, return true; otherwise, return false.
  * `relay_state` - Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access member accounts.
  * `role_configuration_id` - Permission configuration ID.
  * `role_configuration_name` - Permission configuration name.
  * `session_duration` - Session duration. It indicates the maximum session duration when CIC users use the access configuration to access member accounts. Unit: seconds.
  * `update_time` - Update time of the permission configuration.


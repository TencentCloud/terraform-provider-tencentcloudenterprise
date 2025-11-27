---
subcategory: "Backup and Recovery Center(BRC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_brc_autobackup_policies"
sidebar_current: "docs-tencentcloudenterprise-datasource-brc_autobackup_policies"
description: |-
  Use this data source to query detailed information of brc auto backup policies
---

# tencentcloudenterprise_brc_autobackup_policies

Use this data source to query detailed information of brc auto backup policies

## Example Usage

```hcl
data "tencentcloudenterprise_brc_autobackup_policies" "instance_policies" {
  resource_type      = "INSTANCE"
  result_output_file = "cvm_policies.json"
}
```

## Argument Reference

The following arguments are supported:

* `auto_backup_policy_id` - (Optional, String) Auto backup policy ID filter.
* `auto_backup_policy_name` - (Optional, String) Auto backup policy name filter.
* `is_activated` - (Optional, Bool) Filter by activation status.
* `resource_type` - (Optional, String) Resource type filter. Valid values: INSTANCE, DISK, CFS, COS, CSP, MySQL_MariaDB, TDSQL_MySQL.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `auto_backup_policy_list` - Auto backup policy list.
  * `account_name` - Account name.
  * `advanced_retention_policy` - Advanced retention policy configuration.
    * `days` - Retention days.
    * `months` - Retention months.
    * `weeks` - Retention weeks.
    * `years` - Retention years.
  * `auto_backup_policy_id` - Auto backup policy ID.
  * `auto_backup_policy_name` - Auto backup policy name.
  * `auto_backup_policy_state` - Auto backup policy state.
  * `create_time` - Create time.
  * `full_backup_interval` - Full backup interval.
  * `instance_id_set` - Bound instance IDs.
  * `is_activated` - Whether the auto backup policy is activated.
  * `is_permanent` - Whether the backup is permanent.
  * `next_trigger_time` - Next trigger time.
  * `policy` - Backup policy configuration.
    * `hour` - Backup hours (0-23).
    * `interval_days` - Backup interval in days.
  * `resource_type` - Resource type.
  * `retention_amount` - Retention amount.



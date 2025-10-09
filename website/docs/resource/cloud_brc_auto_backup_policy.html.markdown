---
subcategory: "Backup and Recovery Center(BRC)"
layout: "cloud"
page_title: "TencentCloud: cloud_brc_auto_backup_policy"
sidebar_current: "docs-cloud-resource-brc_auto_backup_policy"
description: |-
  Provides a resource to create a brc auto_backup_policy
---

# cloud_brc_auto_backup_policy

Provides a resource to create a brc auto_backup_policy

## Example Usage

```hcl
resource "cloud_brc_auto_backup_policy" "example" {
  resource_type           = "INSTANCE"
  auto_backup_policy_name = "example-policy"

  policy {
    hour          = [0, 11, 12, 13, 14]
    interval_days = 12
  }

  is_permanent         = false
  full_backup_interval = 2
  retention_amount     = 5

  advanced_retention_policy {
    days   = 1
    weeks  = 1
    months = 1
    years  = 1
  }

  is_activated = false
  dry_run      = false
}
```

## Argument Reference

The following arguments are supported:

* `auto_backup_policy_name` - (Required, String) Name of the scheduled backup policy.
* `policy` - (Required, List) Execution policy for scheduled backup.
* `resource_type` - (Required, String, ForceNew) Product type for creating scheduled policy. Valid values: DISK (Cloud Block Storage), CFS (Cloud File Storage), CVM (Cloud Virtual Machine), TDSQL_MySQL (TDSQL MySQL Edition), MySQL_MariaDB (Relational Database MySQL_MariaDB), COS (Cloud Object Storage).
* `advanced_retention_policy` - (Optional, List) Advanced retention policy.
* `create_speed` - (Optional, Int) Bandwidth limit for backup creation, range: [0, 100].
* `dry_run` - (Optional, Bool) Whether to actually create the scheduled backup policy. true means only get the first backup start time without actually creating the scheduled backup policy, false means create, default is false.
* `full_backup_interval` - (Optional, Int) How many backups to do a full backup, 0 means all full backups.
* `is_activated` - (Optional, Bool) Whether to activate the scheduled backup policy.
* `is_permanent` - (Optional, Bool) Whether backups created through this scheduled backup policy are permanently retained. false means not permanently retained, true means permanently retained, default is false.
* `need_archive` - (Optional, Bool) Indicates that backups created by this policy will all be used for archiving.
* `retention_amount` - (Optional, Int) Maximum number of backups retained through this scheduled backup policy. When exceeding this limit, the earliest created backup will be automatically deleted. This parameter cannot conflict with the IsPermanent parameter.

The `advanced_retention_policy` object supports the following:

* `days` - (Optional, Int) Retention days.
* `months` - (Optional, Int) Retention months.
* `weeks` - (Optional, Int) Retention weeks.
* `years` - (Optional, Int) Retention years.

The `policy` object supports the following:

* `hour` - (Required, List) Trigger time for scheduled snapshot policy. Unit: hour, value range: [0, 23]. 00:00 ~ 23:00, 24 time points available, 1 means 01:00, and so on.
* `day_of_month` - (Optional, List) Specify dates from the beginning to the end of each month that need to trigger scheduled backup, value range: [1, 31]. 1-31 represent specific dates of each month, for example, 5 represents the 5th of each month. Note: If dates like 29, 30, 31 that don't exist in some months are set, the corresponding months without these dates will skip scheduled backup.
* `day_of_week` - (Optional, List) Selected dates from Monday to Sunday for creating snapshots, value range: [0, 6]. 0 means Sunday trigger, 1 means Monday trigger, and so on.
* `interval_days` - (Optional, Int) Backup interval in days.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `auto_backup_policy_id` - Auto backup policy ID.
* `next_trigger_time` - First backup start time.


## Import

brc auto_backup_policy can be imported using the id, e.g.

```
terraform import cloud_brc auto_backup_policy.example auto_backup_policy_id
```


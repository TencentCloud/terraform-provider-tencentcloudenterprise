---
subcategory: "Backup and Recovery Center(BRC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_brc_backups"
sidebar_current: "docs-tencentcloudenterprise-datasource-brc_backups"
description: |-
  Use this data source to query detailed information of brc backups
---

# tencentcloudenterprise_brc_backups

Use this data source to query detailed information of brc backups

## Example Usage

```hcl
data "tencentcloudenterprise_brc_backups" "all_backups" {
  result_output_file = "all_backups.json"
}
```

## Argument Reference

The following arguments are supported:

* `backup_group_id` - (Optional, String) Backup group ID filter.
* `backup_id` - (Optional, String) Backup ID filter.
* `backup_name` - (Optional, String) Backup name filter.
* `backup_state` - (Optional, String) Backup state filter.
* `disk_id` - (Optional, String) Disk ID filter for creating backup.
* `disk_usage` - (Optional, String) Disk usage type filter, SYSTEM_DISK (system disk) or DATA_DISK (data disk).
* `platform_project_id` - (Optional, String) Platform project ID filter.
* `result_output_file` - (Optional, String) Used to save results.
* `zone` - (Optional, String) Availability zone filter.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `backup_set` - Backup list.
  * `account_name` - Account name.
  * `account_uin` - Main account UIN.
  * `appid` - User AppId.
  * `archive_status` - Archive status.
  * `auto_backup_policy_id` - ID of the scheduled backup policy that created the current backup, null if manually created.
  * `backup_class` - Full or incremental backup information; FULL indicates full backup, INC indicates incremental backup.
  * `backup_group_id` - Backup group ID associated with the backup.
  * `backup_id` - Backup ID.
  * `backup_name` - Backup name.
  * `backup_size` - Backup size in bytes.
  * `backup_state` - Backup state.
  * `backup_type` - Backup type.
  * `copy_from_remote` - Whether the backup is copied from remote.
  * `copying_to_regions` - List of target regions where the backup is currently being remotely copied.
  * `create_speed` - Backup creation speed.
  * `create_time` - Create time.
  * `deadline_time` - Backup expiration time.
  * `disk_details` - Details of disk attributes at the time of backup creation.
  * `disk_id` - ID of the disk used to create the backup.
  * `disk_name` - Disk name.
  * `disk_size` - Size of the disk used to create the backup, in GB.
  * `disk_usage` - Disk usage type for creating the backup, SYSTEM_DISK (system disk) or DATA_DISK (data disk).
  * `encrypt` - Whether the backup is encrypted.
  * `expire_time` - Expire time.
  * `is_permanent` - Whether the backup is permanently retained.
  * `need_archive` - Whether archiving is needed.
  * `percent` - Backup creation progress percentage.
  * `placement` - Location information of the backup.
    * `cage_id` - Cage ID.
    * `cdc_id` - CDC ID.
    * `cdc_name` - CDC name.
    * `dedicated_cluster_id` - Dedicated cluster ID.
    * `project_id` - Project ID.
    * `project_name` - Project name.
    * `zone` - Availability zone.
  * `platform_project_id` - Project ID to which the backup belongs.
  * `share_reference` - Number of times the backup has been shared.
  * `sub_account_uin` - Sub-account UIN that created the backup.
  * `tags` - List of tags bound to the backup.
    * `key` - Tag key.
    * `value` - Tag value.
  * `zone` - Zone.



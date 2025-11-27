---
subcategory: "Backup and Recovery Center(BRC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_brc_backup_group"
sidebar_current: "docs-tencentcloudenterprise-resources-brc_backup_group"
description: |-
  Provides a resource to create a brc group(cvm instance disks) backup
---

# tencentcloudenterprise_brc_backup_group

Provides a resource to create a brc group(cvm instance disks) backup

## Example Usage

```hcl
resource "tencentcloudenterprise_brc_backup_group" "example" {
  disk_ids          = ["disk-23h487qfj", "disk-da73ha9dj"]
  backup_group_name = "basic-backup-group"
  deadline          = "2024-12-31 23:59:59"
  backup_class      = "FULL"
  create_speed      = 50
  need_archive      = false
}
```

## Argument Reference

The following arguments are supported:

* `disk_ids` - (Required, List: [`String`], ForceNew) CBS disk IDs of the CVM instance.
* `backup_class` - (Optional, String, ForceNew) Backup class. FULL: full backup, INC: incremental backup. Default: FULL.
* `backup_group_name` - (Optional, String, ForceNew) Backup group name.
* `create_speed` - (Optional, Int, ForceNew) Bandwidth limit for creating backup. Range: [0, 100].
* `deadline` - (Optional, String, ForceNew) Deadline of the backup group. If not set, the backup group will be permanent. Standard time format: 2024-04-26 10:00:00.
* `need_archive` - (Optional, Bool, ForceNew) Whether archiving is needed. true: need archive, false: don't need archive. Default: false.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `backup_group_id` - ID of the CVM backup group.


## Import

brc backup_group can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_brc_backup_group.example backup_id
```


---
subcategory: "Backup and Recovery Center(BRC)"
layout: "cloud"
page_title: "TencentCloud: cloud_brc_backup_cfs"
sidebar_current: "docs-cloud-resource-brc_backup_cfs"
description: |-
  Provides a resource to create a brc cfs backup
---

# cloud_brc_backup_cfs

Provides a resource to create a brc cfs backup

## Example Usage

```hcl
resource "cloud_brc_backup_cfs" "example" {
  file_system_id = "cfs-bo74chkp"
  backup_name    = "test"
  deadline       = "2025-07-06 03:06:09"
  backup_class   = "INC"
}
```

## Argument Reference

The following arguments are supported:

* `file_system_id` - (Required, String, ForceNew) The ID of the CFS file system to be backed up.
* `backup_class` - (Optional, String, ForceNew) The class of the backup (e.g., 'FULL' or 'INC'). default: FULL.
* `backup_name` - (Optional, String, ForceNew) The name of the backup.
* `create_speed` - (Optional, Int, ForceNew) The bandwidth limit for creating backup, range: [0, 100].
* `deadline` - (Optional, String, ForceNew) The deadline for the backup. If not set, the backup will be permanent. standard time format: 2024-04-26 10:00:00.
* `need_archive` - (Optional, Bool, ForceNew) Indicates whether this file system backup will be used for archiving.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `backup_id` - The ID of the backup.


## Import

brc backup_cfs can be imported using the id, e.g.

```
terraform import cloud_brc_backup_cfs.example backup_id
```


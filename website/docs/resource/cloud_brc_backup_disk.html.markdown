---
subcategory: "Backup and Recovery Center(BRC)"
layout: "cloud"
page_title: "TencentCloud: cloud_brc_backup_disk"
sidebar_current: "docs-cloud-resource-brc_backup_disk"
description: |-
  Provides a resource to create a brc disk(cbs) backup action
---

# cloud_brc_backup_disk

Provides a resource to create a brc disk(cbs) backup action

## Example Usage

```hcl
resource "cloud_brc_backup_disk" "example" {
  disk_id      = "disk-ewei0a2q"
  backup_name  = "my-backup"
  deadline     = "2025-07-05 19:03:29"
  backup_class = "INC"
}
```

## Argument Reference

The following arguments are supported:

* `disk_id` - (Required, String, ForceNew) The ID of the CBS disk to be backed up.
* `backup_class` - (Optional, String, ForceNew) The class of the backup (e.g., 'FULL' or 'INC'). default: FULL.
* `backup_name` - (Optional, String, ForceNew) The name of the backup.
* `create_speed` - (Optional, Int, ForceNew) The bandwidth limit for creating backup, range: [0, 100].
* `deadline` - (Optional, String, ForceNew) The deadline for the backup. If not set, the backup will be permanent. standard time format: 2024-04-26 10:00:00.
* `need_archive` - (Optional, Bool, ForceNew) Indicates whether this disk backup will be used for archiving.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `backup_id` - The ID of the backup.


## Import

brc backup_disk can be imported using the id, e.g.

```
terraform import cloud_brc_backup_disk.example backup_id
```


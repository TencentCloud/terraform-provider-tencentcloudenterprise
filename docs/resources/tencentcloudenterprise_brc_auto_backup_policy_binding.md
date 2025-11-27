---
subcategory: "Backup and Recovery Center(BRC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_brc_auto_backup_policy_binding"
sidebar_current: "docs-tencentcloudenterprise-resources-brc_auto_backup_policy_binding"
description: |-
  Provides a resource to create a brc auto_snapshot_policy binding
---

# tencentcloudenterprise_brc_auto_backup_policy_binding

Provides a resource to create a brc auto_snapshot_policy binding

## Example Usage

```hcl
resource "tencentcloudenterprise_brc_auto_backup_policy_binding" "example" {
  auto_backup_policy_id = tencentcloudenterprise_brc_autobackup_policy.example.id
  instance_ids          = ["ins-21ahx7qj"]
  resource_type         = "INSTANCE"
}
```

## Argument Reference

The following arguments are supported:

* `auto_backup_policy_id` - (Required, String, ForceNew) Auto backup policy ID.
* `bucket_details` - (Optional, List, ForceNew) The mapping of corresponding relationships between the backup source bucket to be bound, prefix, and the region where the bucket belongs, which is used for regular COS backup.
* `disk_ids` - (Optional, List: [`String`], ForceNew) Disk IDs to bind with the auto backup policy.
* `file_system_ids` - (Optional, List: [`String`], ForceNew) File system IDs to bind with the auto backup policy.
* `instance_ids` - (Optional, List: [`String`], ForceNew) Instance IDs to bind with the auto backup policy.
* `resource_ids` - (Optional, List: [`String`], ForceNew) Resource IDs to bind with the auto backup policy.
* `resource_type` - (Optional, String, ForceNew) Resource type. Valid values: INSTANCE, DISK, CFS, COS, CSP, MySQL_MariaDB, TDSQL_MySQL.

The `bucket_details` object supports the following:

* `cos_region` - (Optional, String, ForceNew) cos bucket backup region, required for cos backup.
* `prefix` - (Optional, String, ForceNew) cos bucket backup prefix, required for cos backup.
* `resource_id` - (Optional, String, ForceNew) cos bucket backup info, required for cos backup.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




---
subcategory: "Backup and Recovery Center(BRC)"
layout: "cloud"
page_title: "TencentCloud: cloud_brc_backup_resource"
sidebar_current: "docs-cloud-resource-brc_backup_resource"
description: |-
  Provides a resource to create a brc resource backup
---

# cloud_brc_backup_resource

Provides a resource to create a brc resource backup

## Example Usage

```hcl
# cos backup

resource "cloud_brc_backup_resource" "cos_backup" {
  resource_type = "COS"
  backup_name   = "my-cos-backup"
  bucket_detail {
    cos_region  = "chongqing"
    prefix      = "/"
    resource_id = "test-456-1255000252"
  }
}

# mysql backup

resource "cloud_brc_backup_resource" "mysql_backup" {
  resource_type = "MySQL_MariaDB"
  backup_name   = "my-mysql-backup"
  resource_id   = "tdsql-baxmqi05"
}
```

## Argument Reference

The following arguments are supported:

* `resource_type` - (Required, String, ForceNew) Resource type for backup creation. Valid values: COS, CSP, MySQL_MariaDB, TDSQL_MySQL. Make sure the resource is deployed in BRC.
* `backup_name` - (Optional, String, ForceNew) Name of the backup.
* `bucket_detail` - (Optional, List, ForceNew) Bucket details mapping for COS backup. Required for COS resource type.
* `create_speed` - (Optional, Int, ForceNew) Bandwidth limit for backup creation, range: [0, 100] Mbps.
* `deadline` - (Optional, String, ForceNew) Backup expiration time. If not specified, the backup will be kept permanently. Example format: '2024-04-26 10:00:00'.
* `resource_id` - (Optional, String, ForceNew) ID of the resource to be backed up.

The `bucket_detail` object supports the following:

* `cos_region` - (Required, String, ForceNew) COS bucket region.
* `resource_id` - (Required, String, ForceNew) COS bucket ID.
* `prefix` - (Optional, String, ForceNew) COS bucket prefix path.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `backup_id` - ID of the created backup.
* `delete_retreated_backup` - Whether to delete the rolled-back backup data, which is used when deleting resources.


## Import

brc backup_resource can be imported using the id, e.g.

```
terraform import cloud_brc_backup_resource.example backup_id
```


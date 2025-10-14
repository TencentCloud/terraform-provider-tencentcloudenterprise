---
subcategory: "Cloud File Storage(CFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cfs_snapshot"
sidebar_current: "docs-cloud-resources-cfs_snapshot"
description: |-
  Provides a resource to create a cfs snapshot
---

# cloud_cfs_snapshot

Provides a resource to create a cfs snapshot

## Example Usage

```hcl
resource "cloud_cfs_snapshot" "snapshot" {
  file_system_id = "cfs-iobiaxtj"
  snapshot_name  = "test"
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `file_system_id` - (Required, String, ForceNew) Id of file system.
* `snapshot_name` - (Required, String) Name of snapshot.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `job_id` - Job id of create snapshot.


## Import

cfs snapshot can be imported using the id, e.g.

```
terraform import cloud_cfs_snapshot.snapshot snapshot_id
```


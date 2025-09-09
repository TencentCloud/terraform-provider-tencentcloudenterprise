---
subcategory: "Cloud File Storage(CFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cfs_auto_snapshot_policy_attachment"
sidebar_current: "docs-cloud-resource-cfs_auto_snapshot_policy_attachment"
description: |-
  Provides a resource to create a cfs auto_snapshot_policy_attachment
---

# cloud_cfs_auto_snapshot_policy_attachment

Provides a resource to create a cfs auto_snapshot_policy_attachment

## Example Usage

```hcl
resource "cloud_cfs_auto_snapshot_policy_attachment" "auto_snapshot_policy_attachment" {
  auto_snapshot_policy_id = "asp-basic"
  file_system_ids         = "cfs-4xzkct19,cfs-iobiaxtj"
}
```

## Argument Reference

The following arguments are supported:

* `auto_snapshot_policy_id` - (Required, String, ForceNew) ID of the snapshot to be unbound.
* `file_system_ids` - (Required, String, ForceNew) List of IDs of the file systems to be unbound, separated by comma.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cfs auto_snapshot_policy_attachment can be imported using the id, e.g.

```
terraform import cloud_cfs_auto_snapshot_policy_attachment.auto_snapshot_policy_attachment auto_snapshot_policy_id#file_system_ids
```


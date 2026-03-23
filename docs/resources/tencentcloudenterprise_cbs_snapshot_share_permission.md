---
subcategory: "Cloud Block Storage(CBS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cbs_snapshot_share_permission"
sidebar_current: "docs-tencentcloudenterprise-resource-cbs_snapshot_share_permission"
description: |-
  Provides a resource to create a cbs snapshot_share_permission
---

# tencentcloudenterprise_cbs_snapshot_share_permission

Provides a resource to create a cbs snapshot_share_permission

## Example Usage

```hcl
resource "tencentcloudenterprise_cbs_snapshot_share_permission" "snapshot_share_permission" {
  account_ids = ["1xxxxxx", "2xxxxxx"]
  snapshot_id = "snap-xxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `account_ids` - (Required, Set: [`String`]) List of account IDs with which a snapshot is shared.
* `snapshot_id` - (Required, String) The ID of the snapshot to be queried.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cbs_snapshot_share_permission can be imported using the id, e.g.

```
cbs snapshot_share_permission can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cbs_snapshot_share_permission.snapshot_share_permission snap-xxxxxx
```
```


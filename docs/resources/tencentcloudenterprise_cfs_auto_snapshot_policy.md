---
subcategory: "Cloud File Storage(CFS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfs_auto_snapshot_policy"
sidebar_current: "docs-tencentcloudenterprise-resources-cfs_auto_snapshot_policy"
description: |-
  Provides a resource to create a cfs auto_snapshot_policy
---

# tencentcloudenterprise_cfs_auto_snapshot_policy

Provides a resource to create a cfs auto_snapshot_policy

## Example Usage

### use day of week

```hcl
resource "tencentcloudenterprise_cfs_auto_snapshot_policy" "auto_snapshot_policy" {
  day_of_week = "1,2"
  hour        = "2,3"
  policy_name = "policy_name"
  alive_days  = 7
}
```

### use day of month

```hcl
resource "tencentcloudenterprise_cfs_auto_snapshot_policy" "auto_snapshot_policy" {
  hour         = "2,3"
  policy_name  = "policy_name"
  alive_days   = 7
  day_of_month = "2,3,4"
}
```

### use interval days

```hcl
resource "tencentcloudenterprise_cfs_auto_snapshot_policy" "auto_snapshot_policy" {
  hour          = "2,3"
  policy_name   = "policy_name"
  alive_days    = 7
  interval_days = 1
}
```

## Argument Reference

The following arguments are supported:

* `day_of_week` - (Required, String) The day of the week on which to repeat the snapshot operation.
* `hour` - (Required, String) The time point when to repeat the snapshot operation.
* `alive_days` - (Optional, Int) Snapshot retention period.
* `policy_name` - (Optional, String) Policy name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `auto_snapshot_policy_id` - The ID of the auto snapshot policy.


## Import

cfs auto_snapshot_policy can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cfs_auto_snapshot_policy.auto_snapshot_policy auto_snapshot_policy_id
```


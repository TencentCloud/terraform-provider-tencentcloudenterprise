---
subcategory: "Parallel File Storage(TurboFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_auto_snapshot_policy"
sidebar_current: "docs-cloud-resources-turbofs_auto_snapshot_policy"
description: |-
  Provides a resource to create a turbofs auto_snapshot_policy
---

# cloud_turbofs_auto_snapshot_policy

Provides a resource to create a turbofs auto_snapshot_policy

## Example Usage

### use day of week

```hcl
resource "cloud_turbofs_auto_snapshot_policy" "auto_snapshot_policy" {
  day_of_week = "1,2"
  hour        = "2,3"
  policy_name = "policy_name"
  alive_days  = 7
}
```

### use day of month

```hcl
resource "cloud_turbofs_auto_snapshot_policy" "auto_snapshot_policy" {
  hour         = "2,3"
  policy_name  = "policy_name"
  alive_days   = 7
  day_of_month = "2,3,4"
}
```

## Argument Reference

The following arguments are supported:

* `hour` - (Required, String) The time point when to repeat the snapshot operation.
* `policy_name` - (Required, String) Policy name.
* `alive_days` - (Optional, Int) Snapshot retention period. default value is 0, which means the snapshot will be retained permanently.
* `day_of_month` - (Optional, String) The day of the month on which to repeat the snapshot operation, conflict with interval_days and day_of_week.
* `day_of_week` - (Optional, String) The day of the week on which to repeat the snapshot operation, conflict with interval_days and day_of_month.
* `interval_days` - (Optional, Int) The interval days which to repeat the snapshot operation, conflict with day_of_week and day_of_month.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `auto_snapshot_policy_id` - The ID of the auto snapshot policy.


## Import

turbofs auto_snapshot_policy can be imported using the id, e.g.

```
terraform import cloud_turbofs_auto_snapshot_policy.auto_snapshot_policy auto_snapshot_policy_id
```


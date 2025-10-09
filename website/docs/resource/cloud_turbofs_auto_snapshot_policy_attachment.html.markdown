---
subcategory: "Parallel File Storage(TurboFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_auto_snapshot_policy_attachment"
sidebar_current: "docs-cloud-resource-turbofs_auto_snapshot_policy_attachment"
description: |-
  Provides a resource to create a turbofs auto_snapshot_policy_attachment
---

# cloud_turbofs_auto_snapshot_policy_attachment

Provides a resource to create a turbofs auto_snapshot_policy_attachment

## Example Usage

```hcl
resource "cloud_turbofs_auto_snapshot_policy_attachment" "auto_snapshot_policy_attachment" {
  auto_snapshot_policy_id = "asp-basic"
  file_system_id          = "turbofs-4xzkct19"
}
```

## Argument Reference

The following arguments are supported:

* `auto_snapshot_policy_id` - (Required, String, ForceNew) ID of the snapshot to bound.
* `file_system_id` - (Required, String, ForceNew) ID of the file systems to bound.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

turbofs auto_snapshot_policy_attachment can be imported using the id, e.g.

```
terraform import cloud_turbofs_auto_snapshot_policy_attachment.auto_snapshot_policy_attachment auto_snapshot_policy_id#file_system_ids
```


---
subcategory: "Parallel File Storage(TurboFS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_turbofs_auto_snapshot_policy_attachment"
sidebar_current: "docs-tencentcloudenterprise-resources-turbofs_auto_snapshot_policy_attachment"
description: |-
  Provides a resource to create a turbofs auto_snapshot_policy_attachment
---

# tencentcloudenterprise_turbofs_auto_snapshot_policy_attachment

Provides a resource to create a turbofs auto_snapshot_policy_attachment

## Example Usage

```hcl
resource "tencentcloudenterprise_turbofs_auto_snapshot_policy_attachment" "auto_snapshot_policy_attachment" {
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
terraform import tencentcloudenterprise_turbofs_auto_snapshot_policy_attachment.auto_snapshot_policy_attachment auto_snapshot_policy_id#file_system_ids
```


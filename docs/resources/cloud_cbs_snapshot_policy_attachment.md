---
subcategory: "Cloud Block Storage(CBS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cbs_snapshot_policy_attachment"
sidebar_current: "docs-cloud-resources-cbs_snapshot_policy_attachment"
description: |-
  Provides a CBS snapshot policy attachment resource.
---

# cloud_cbs_snapshot_policy_attachment

Provides a CBS snapshot policy attachment resource.

## Example Usage

```hcl
resource "cloud_cbs_snapshot_policy_attachment" "foo" {
  storage_id         = cloud_cbs_storage.foo.id
  snapshot_policy_id = cloud_cbs_snapshot_policy.policy.id
}
```

## Argument Reference

The following arguments are supported:

* `snapshot_policy_id` - (Required, String, ForceNew) ID of CBS snapshot policy.
* `storage_id` - (Required, String, ForceNew) ID of CBS.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




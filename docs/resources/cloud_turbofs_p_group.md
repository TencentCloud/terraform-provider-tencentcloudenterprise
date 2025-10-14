---
subcategory: "Parallel File Storage(TurboFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_p_group"
sidebar_current: "docs-cloud-resources-turbofs_p_group"
description: |-
  Provides a resource to create a TurboFS permission group.
---

# cloud_turbofs_p_group

Provides a resource to create a TurboFS permission group.

## Example Usage

```hcl
resource "cloud_turbofs_p_group" "foo" {
  name      = "test_p_group"
  desc_info = "test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the permission group.
* `desc_info` - (Optional, String) Description of the permission group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `bind_cfs_num` - The number of file systems associated with the permission group.
* `c_date` - Create datetime of the permission group.
* `p_group_id` - Id of the permission group.


## Import

TurboFS permission group can be imported using the id, e.g.

```
$ terraform import cloud_turbofs_p_group.foo pgroup-7nx89k7l
```


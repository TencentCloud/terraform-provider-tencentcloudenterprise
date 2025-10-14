---
subcategory: "Bare Metal Server(BMS)"
layout: "cloud"
page_title: "TencentCloud: cloud_bms_placement_group"
sidebar_current: "docs-cloud-resources-bms_placement_group"
description: |-
  Provide a resource to create a placement group.
---

# cloud_bms_placement_group

Provide a resource to create a placement group.

## Example Usage

```hcl
resource "cloud_bms_placement_group" "foo" {
  name = "test"
  type = "RACK"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the placement group, 1-60 characters in length.
* `type` - (Required, String, ForceNew) Type of the placement group. Valid values: `RACK_SAME_SW` and `RACK`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of the placement group.
* `current_num` - Number of hosts in the placement group.


## Import

Placement group can be imported using the id, e.g.

```
$ terraform import cloud_bms_placement_group.foo ps-ilan8vjf
```


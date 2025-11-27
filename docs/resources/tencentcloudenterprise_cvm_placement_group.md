---
subcategory: "Cloud Virtual Machine(CVM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cvm_placement_group"
sidebar_current: "docs-tencentcloudenterprise-resources-cvm_placement_group"
description: |-
  Provide a resource to create a placement group.
---

# tencentcloudenterprise_cvm_placement_group

Provide a resource to create a placement group.

## Example Usage

```hcl
resource "tencentcloudenterprise_cvm_placement_group" "foo" {
  name = "test"
  type = "HOST"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the placement group, 1-60 characters in length.
* `type` - (Required, String, ForceNew) Type of the placement group. Valid values: `HOST`, `SW` and `RACK`.
* `strategy` - (Optional, String) Types are: `SPREAD`,`WEAK_SPREAD`,`AFFINITY` and `WEAK_AFFINITY`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of the placement group.
* `current_num` - Number of hosts in the placement group.
* `cvm_quota_total` - Maximum number of hosts in the placement group.


## Import

Placement group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cvm_placement_group.foo ps-ilan8vjf
```


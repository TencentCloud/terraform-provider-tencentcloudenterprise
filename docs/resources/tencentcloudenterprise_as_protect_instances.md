---
subcategory: "Auto Scaling(AS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_as_protect_instances"
sidebar_current: "docs-tencentcloudenterprise-resources-as_protect_instances"
description: |-
  Provides a resource to create a as protect_instances
---

# tencentcloudenterprise_as_protect_instances

Provides a resource to create a as protect_instances

## Example Usage

```hcl
resource "tencentcloudenterprise_as_protect_instances" "protect_instances" {
  auto_scaling_group_id   = tencentcloudenterprise_as_scaling_group.scaling_group.id
  instance_ids            = ["ins-xxxxx"]
  protected_from_scale_in = true
}
```

## Argument Reference

The following arguments are supported:

* `auto_scaling_group_id` - (Required, String, ForceNew) Launch configuration ID.
* `instance_ids` - (Required, Set: [`String`], ForceNew) List of cvm instances to remove.
* `protected_from_scale_in` - (Required, Bool, ForceNew) If instances need protect.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




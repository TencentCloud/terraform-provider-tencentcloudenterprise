---
subcategory: "Auto Scaling(AS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_as_remove_instances"
sidebar_current: "docs-tencentcloudenterprise-resources-as_remove_instances"
description: |-
  Provides a resource to create a as remove_instances
---

# tencentcloudenterprise_as_remove_instances

Provides a resource to create a as remove_instances

## Example Usage

```hcl
resource "tencentcloudenterprise_as_remove_instances" "remove_instances" {
  auto_scaling_group_id = tencentcloudenterprise_as_scaling_group.scaling_group.id
  instance_ids          = ["ins-xxxxxx"]
}
```

## Argument Reference

The following arguments are supported:

* `auto_scaling_group_id` - (Required, String, ForceNew) Launch configuration ID.
* `instance_ids` - (Required, Set: [`String`], ForceNew) List of cvm instances to remove.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




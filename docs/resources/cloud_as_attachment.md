---
subcategory: "Auto Scaling(AS)"
layout: "cloud"
page_title: "TencentCloud: cloud_as_attachment"
sidebar_current: "docs-cloud-resources-as_attachment"
description: |-
  Provides a resource to attach or detach CVM instances to a specified scaling group.
---

# cloud_as_attachment

Provides a resource to attach or detach CVM instances to a specified scaling group.

## Example Usage

```hcl
resource "cloud_as_attachment" "attachment" {
  scaling_group_id = "sg-afasfa"
  instance_ids     = ["ins-01", "ins-02"]
}
```

## Argument Reference

The following arguments are supported:

* `instance_ids` - (Required, Set: [`String`]) ID list of CVM instances to be attached to the scaling group.
* `scaling_group_id` - (Required, String, ForceNew) ID of a scaling group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




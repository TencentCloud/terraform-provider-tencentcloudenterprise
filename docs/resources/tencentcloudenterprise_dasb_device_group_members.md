---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_device_group_members"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_device_group_members"
description: |-
  Provide a resource to create a DASB device group members
---

# tencentcloudenterprise_dasb_device_group_members

Provide a resource to create a DASB device group members

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_device_group_members" "example" {
  device_group_id = 1
  member_id_set   = [1, 2]
}
```

## Argument Reference

The following arguments are supported:

* `device_group_id` - (Required, Int, ForceNew) Device Group ID.
* `member_id_set` - (Required, Set: [`Int`], ForceNew) A collection of device IDs that need to be added to the device group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_dasb_device_group_members can be imported using the id, e.g.

```
DASB device group members can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_device_group_members.example 1#1,2
```
```


---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_group_membership"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_group_membership"
description: |-
  Provides a resource to create a CAM group membership.
---

# tencentcloudenterprise_cam_group_membership

Provides a resource to create a CAM group membership.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_group_membership" "foo" {
  group_id   = tencentcloudenterprise_cam_group.foo.id
  user_names = [tencentcloudenterprise_cam_user.foo.name, tencentcloudenterprise_cam_user.bar.name]
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required, String, ForceNew) ID of CAM group.
* `user_names` - (Required, Set: [`String`]) User name set as ID of the CAM group members.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cam_group_membership can be imported using the id, e.g.

```
CAM group membership can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_group_membership.foo 12515263
```
```


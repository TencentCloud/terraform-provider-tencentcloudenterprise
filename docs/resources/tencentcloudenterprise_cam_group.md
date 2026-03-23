---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_group"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_group"
description: |-
  Provides a resource to create a CAM group.
---

# tencentcloudenterprise_cam_group

Provides a resource to create a CAM group.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_group" "foo" {
  name   = "cam-group-test"
  remark = "test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of CAM group.
* `remark` - (Optional, String) Description of the CAM group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `channel` - Message receiving channel. 1: SMS, 2: Email, 3: SMS+Email.
* `create_time` - Create time of the CAM group.
* `group_type` - Group type. 0: custom group, 1: preset group.

## Import

tencentcloudenterprise_cam_group can be imported using the id, e.g.

```
CAM group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_group.foo 90496
```
```


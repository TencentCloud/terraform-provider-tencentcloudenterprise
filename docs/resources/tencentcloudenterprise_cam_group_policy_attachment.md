---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_group_policy_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_group_policy_attachment"
description: |-
  Provides a resource to create a CAM group policy attachment.
---

# tencentcloudenterprise_cam_group_policy_attachment

Provides a resource to create a CAM group policy attachment.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_group_policy_attachment" "foo" {
  group_id  = tencentcloudenterprise_cam_group.foo.id
  policy_id = tencentcloudenterprise_cam_policy.foo.id
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required, String, ForceNew) ID of the attached CAM group.
* `policy_id` - (Required, String, ForceNew) ID of the policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_mode` - Mode of Creation of the CAM group policy attachment. `1` means the cam policy attachment is created by production, and the others indicate syntax strategy ways.
* `create_time` - Create time of the CAM group policy attachment.
* `policy_name` - Name of the policy.
* `policy_type` - Type of the policy strategy. 'Group' means customer strategy and 'QCS' means preset strategy.

## Import

tencentcloudenterprise_cam_group_policy_attachment can be imported using the id, e.g.

```
CAM group policy attachment can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cam_group_policy_attachment.foo 12515263#26800353
```
```


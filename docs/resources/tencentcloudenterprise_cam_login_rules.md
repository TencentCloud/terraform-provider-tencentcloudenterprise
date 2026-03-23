---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_login_rules"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_login_rules"
description: |-
  Provides a resource to manage CAM login rules.
---

# tencentcloudenterprise_cam_login_rules

Provides a resource to manage CAM login rules.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_login_rules" "foo" {
  session_duration = 1440
}
```

## Argument Reference

The following arguments are supported:

* `session_duration` - (Optional, Int) Login session duration in minutes.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



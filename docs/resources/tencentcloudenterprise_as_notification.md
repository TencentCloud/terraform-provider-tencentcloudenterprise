---
subcategory: "Auto Scaling(AS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_as_notification"
sidebar_current: "docs-tencentcloudenterprise-resources-as_notification"
description: |-
  Provides a resource for an AS (Auto scaling) notification.
---

# tencentcloudenterprise_as_notification

Provides a resource for an AS (Auto scaling) notification.

## Example Usage

```hcl
resource "tencentcloudenterprise_as_notification" "as_notification" {
  scaling_group_id            = "sg-12af45"
  notification_types          = ["SCALE_OUT_FAILED", "SCALE_IN_SUCCESSFUL", "SCALE_IN_FAILED", "REPLACE_UNHEALTHY_INSTANCE_FAILED"]
  notification_user_group_ids = ["76955"]
}
```

## Argument Reference

The following arguments are supported:

* `notification_types` - (Required, List: [`String`]) A list of Notification Types that trigger notifications. Acceptable values are `SCALE_OUT_FAILED`, `SCALE_IN_SUCCESSFUL`, `SCALE_IN_FAILED`, `REPLACE_UNHEALTHY_INSTANCE_SUCCESSFUL` and `REPLACE_UNHEALTHY_INSTANCE_FAILED`.
* `notification_user_group_ids` - (Required, List: [`String`]) A group of user IDs to be notified.
* `scaling_group_id` - (Required, String, ForceNew) ID of a scaling group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




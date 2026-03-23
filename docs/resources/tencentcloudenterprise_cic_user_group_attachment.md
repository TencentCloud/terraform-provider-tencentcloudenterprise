---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_user_group_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-cic_user_group_attachment"
description: |-
  Provides a resource to create an identity center user group attachment
---

# tencentcloudenterprise_cic_user_group_attachment

Provides a resource to create an identity center user group attachment

## Example Usage

```hcl
resource "tencentcloudenterprise_cic_user_group_attachment" "cic_user_group_attachment" {
  zone_id  = "z-xxxxxx"
  user_id  = "u-xxxxxx"
  group_id = "g-xxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `group_id` - (Required, String, ForceNew) User group ID.
* `user_id` - (Required, String, ForceNew) User ID.
* `zone_id` - (Required, String, ForceNew) Zone id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cic_user_group_attachment can be imported using the id, e.g.

```
organization cic_user_group_attachment can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_user_group_attachment.cic_user_group_attachment ${zoneId}#${groupId}#${userId}
```
```


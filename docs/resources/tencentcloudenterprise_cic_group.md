---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_group"
sidebar_current: "docs-tencentcloudenterprise-resource-cic_group"
description: |-
  Provides a resource to create an identity center group
---

# tencentcloudenterprise_cic_group

Provides a resource to create an identity center group

## Example Usage

```hcl
resource "tencentcloudenterprise_cic_group" "cic_group" {
  zone_id     = "z-xxxxxx"
  group_name  = "test-group"
  description = "test"
}
```

## Argument Reference

The following arguments are supported:

* `group_name` - (Required, String) The name of the user group. Format: Allow English letters, numbers and special characters-. Length: Maximum 128 characters.
* `zone_id` - (Required, String) Zone id.
* `description` - (Optional, String) A description of the user group. Length: Maximum 1024 characters.
* `group_type` - (Optional, String) Type of user group. `Manual`: manual creation, `Synchronized`: external import.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of the user group.
* `group_id` - ID of the user group.
* `member_count` - Number of team members.
* `update_time` - Modification time for the user group.

## Import

tencentcloudenterprise_cic_group can be imported using the id, e.g.

```
tencentcloudenterprise_cic_group can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_group.cic_group ${zoneId}#${groupId}
```
```


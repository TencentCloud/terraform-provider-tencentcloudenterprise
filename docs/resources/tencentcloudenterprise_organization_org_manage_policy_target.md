---
subcategory: "Tencent Cloud Organization"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_organization_org_manage_policy_target"
sidebar_current: "docs-tencentcloudenterprise-resource-organization_org_manage_policy_target"
description: |-
  Provides a resource to create an organization manage policy target
---

# tencentcloudenterprise_organization_org_manage_policy_target

Provides a resource to create an organization manage policy target

## Example Usage

```hcl
resource "tencentcloudenterprise_organization_org_manage_policy_target" "example" {
  target_id   = 86
  target_type = "NODE"
  policy_id   = 44909
  policy_type = "SERVICE_CONTROL_POLICY"
}
```

## Argument Reference

The following arguments are supported:

* `policy_id` - (Required, Int, ForceNew) Policy Id.
* `target_id` - (Required, Int, ForceNew) Binding target ID of the policy. Member Uin or Department ID.
* `target_type` - (Required, String, ForceNew) Target type.
Valid values:
  - `NODE`: Department.
  - `MEMBER`: Check Member.
* `policy_type` - (Optional, String, ForceNew) Policy type. Default value is SERVICE_CONTROL_POLICY.
Valid values:
  - `SERVICE_CONTROL_POLICY`: Service control policy.
  - `TAG_POLICY`: Tag policy.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



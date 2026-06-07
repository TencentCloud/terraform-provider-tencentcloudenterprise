---
subcategory: "Tencent Cloud Organization"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_organization_org_service_assign_member"
sidebar_current: "docs-tencentcloudenterprise-resource-organization_org_service_assign_member"
description: |-
  Provides a resource to manage an organization trusted-service delegated administrator (assign member).
---

# tencentcloudenterprise_organization_org_service_assign_member

Provides a resource to manage an organization trusted-service delegated administrator (assign member).

Each resource instance represents a single delegated administrator binding for one
trusted service (`service_id`) on one member account (`member_uin`).

## Example Usage

```hcl
resource "tencentcloudenterprise_organization_org_service_assign_member" "demo" {
  service_id       = 24
  member_uin       = 110000003055
  management_scope = 1
}
```

## Argument Reference

The following arguments are supported:

* `member_uin` - (Required, Int, ForceNew) The delegated administrator member UIN.
* `service_id` - (Required, Int, ForceNew) Trusted service ID. e.g. 24 stands for cloudaudit.
* `management_scope_node_ids` - (Optional, Set: [`Int`], ForceNew) Node ID list. Effective only when `management_scope` is `2`.
* `management_scope_uins` - (Optional, Set: [`Int`], ForceNew) Member UIN list. Effective only when `management_scope` is `2`.
* `management_scope` - (Optional, Int, ForceNew) Management scope. `1`: all members; `2`: partial members (must specify `management_scope_uins` and/or `management_scope_node_ids`).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Time when the assignment is created.
* `member_name` - Member account name.
* `product_name` - Product name of the trusted service.

## Import

tencentcloudenterprise_organization_org_service_assign_member can be imported using the id, e.g.

```
Organization service assign member can be imported using `service_id#member_uin`, e.g.

```
$ terraform import tencentcloudenterprise_organization_org_service_assign_member.demo 24#110000003055
```
```


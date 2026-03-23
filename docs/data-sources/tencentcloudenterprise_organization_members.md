---
subcategory: "Tencent Cloud Organization"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_organization_members"
sidebar_current: "docs-tencentcloudenterprise-datasource-organization_members"
description: |-
  Use this data source to query detailed information of organization members
---

# tencentcloudenterprise_organization_members

Use this data source to query detailed information of organization members

## Example Usage

```hcl
data "tencentcloudenterprise_organization_members" "members" {
}

output "members_list" {
  value = data.tencentcloudenterprise_organization_members.members.items
}
```

## Argument Reference

The following arguments are supported:

* `auth_name` - (Optional, String) Entity name.
* `lang` - (Optional, String) Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).
* `product` - (Optional, String) Abbreviation of the trusted service, which is required during querying the trusted service admin.
* `result_output_file` - (Optional, String) Used to save results.
* `search_key` - (Optional, String) Search by member name or ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `items` - Member list.
  * `create_time` - Creation time.
  * `is_allow_quit` - Whether the member is allowed to leave. Valid values: `Allow`, `Denied`.
  * `member_type` - Member type. Valid values: `Invite` (invited); `Create` (created).
  * `member_uin` - Member UIN.
  * `name` - Member name.
  * `node_id` - Node ID.
  * `node_name` - Node name.
  * `org_identity` - Management identity.
    * `identity_alias_name` - Identity name.
    * `identity_id` - Identity ID.
  * `org_permission` - Relationship policy permission.
    * `id` - Permission ID.
    * `name` - Permission name.
  * `org_policy_name` - Relationship policy name.
  * `org_policy_type` - Relationship policy type.
  * `pay_name` - Payer name.
  * `pay_uin` - Payer UIN.
  * `permission_status` - Member permission status. Valid values: `Confirmed`, `UnConfirmed`.
  * `remark` - Remarks.
  * `update_time` - Update time.


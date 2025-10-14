---
subcategory: "CFS Turbo(Turbofs)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_access_rules"
sidebar_current: "docs-cloud-datasource-turbofs_access_rules"
description: |-
  Use this data source to query the detail information of Turbofs access rule.
---

# cloud_turbofs_access_rules

Use this data source to query the detail information of Turbofs access rule.

## Example Usage

```hcl
data "cloud_turbofs_access_rules" "access_rules" {
  access_group_id = "pgroup-7nx89k7l"
  access_rule_id  = "rule-qcndbqzj"
}
```

## Argument Reference

The following arguments are supported:

* `access_group_id` - (Required, String) A specified access group ID used to query.
* `access_rule_id` - (Optional, String) A specified access rule ID used to query.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `access_rule_list` - An information list of Turbofs access rule. Each element contains the following attributes:
  * `access_rule_id` - ID of the access rule.
  * `auth_client_ip` - Allowed IP of the access rule.
  * `priority` - The priority level of access rule.
  * `rw_permission` - Read and write permissions.
  * `user_permission` - The permissions of accessing users.



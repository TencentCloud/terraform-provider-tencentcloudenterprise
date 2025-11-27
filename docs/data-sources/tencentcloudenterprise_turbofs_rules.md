---
subcategory: "Parallel File Storage(TurboFS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_turbofs_rules"
sidebar_current: "docs-tencentcloudenterprise-datasource-turbofs_rules"
description: |-
  Use this data source to query the detail information of Turbofs permission rule.
---

# tencentcloudenterprise_turbofs_rules

Use this data source to query the detail information of Turbofs permission rule.

## Example Usage

```hcl
data "tencentcloudenterprise_turbofs_rules" "rules" {
  p_group_id = "pgroup-7nx89k7l"
  rule_id    = "rule-qcndbqzj"
}
```

## Argument Reference

The following arguments are supported:

* `p_group_id` - (Required, String) A specified permission group ID used to query.
* `result_output_file` - (Optional, String) Used to save results.
* `rule_id` - (Optional, String) A specified permission rule ID used to query.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `rule_list` - An information list of Turbofs permission rule. Each element contains the following attributes:
  * `auth_client_ip` - Allowed IP of the permission rule.
  * `priority` - The priority level of rule.
  * `rule_id` - ID of the permission rule.
  * `rw_permission` - Read and write permissions.
  * `user_permission` - The permissions of users.



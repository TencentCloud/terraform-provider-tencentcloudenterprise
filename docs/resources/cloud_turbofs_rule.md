---
subcategory: "Parallel File Storage(TurboFS)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_rule"
sidebar_current: "docs-cloud-resources-turbofs_rule"
description: |-
  Provides a resource to create a TurboFS permission group rule.
---

# cloud_turbofs_rule

Provides a resource to create a TurboFS permission group rule.

## Example Usage

```hcl
resource "cloud_turbofs_rule" "foo" {
  p_group_id      = "pgroup-7nx89k7l"
  auth_client_ip  = "10.10.1.0/24"
  priority        = 1
  rw_permission   = "ro"
  user_permission = "root_squash"
}
```

## Argument Reference

The following arguments are supported:

* `auth_client_ip` - (Required, String) A single IP or a single IP address range such as 10.1.10.11 or 10.10.1.0/24 indicates that all IPs are allowed. Please note that the IP entered should be CVM's private IP.
* `p_group_id` - (Required, String, ForceNew) ID of a permission group.
* `priority` - (Required, Int) The priority level of rule. Valid value ranges: (1~100). `1` indicates the highest priority.
* `rw_permission` - (Optional, String) Read and write permissions. Valid values are `ro` and `rw`, and default is `rw`.
* `user_permission` - (Optional, String) The permissions of users. Valid values are `all_squash`, `no_all_squash`, `root_squash` and `no_root_squash`. and default is `root_squash`. `all_squash` indicates that all users are mapped as anonymous users or user groups; `no_all_squash` indicates that users will match local users first and be mapped to anonymous users or user groups after matching failed; `root_squash` indicates that map root users to anonymous users or user groups; `no_root_squash` indicates that root users keep root account permission.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `rule_id` - The id of rule.



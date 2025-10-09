---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_security_groups"
sidebar_current: "docs-cloud-datasource-dcdb_security_groups"
description: |-
  Use this data source to query detailed information of dcdb securityGroups
---

# cloud_dcdb_security_groups

Use this data source to query detailed information of dcdb securityGroups

## Example Usage

```hcl
data "cloud_dcdb_security_groups" "securityGroups" {
  instance_id = "your_instance_id"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) Instance id.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - Security group list.
  * `create_time` - Create time.
  * `inbound` - Inbound rules.
    * `action` - Policy action.
    * `cidr_ip` - Cidr ip.
    * `ip_protocol` - Internet protocol.
    * `port_range` - Port range.
  * `outbound` - Outbound rules.
    * `action` - Policy action.
    * `cidr_ip` - Cidr ip.
    * `ip_protocol` - Internet protocol.
    * `port_range` - Port range.
  * `project_id` - Project id.
  * `security_group_id` - Security group id.
  * `security_group_name` - Security group name.



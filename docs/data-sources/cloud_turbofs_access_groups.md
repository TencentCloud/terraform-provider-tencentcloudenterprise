---
subcategory: "CFS Turbo(Turbofs)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_access_groups"
sidebar_current: "docs-cloud-datasource-turbofs_access_groups"
description: |-
  Use this data source to query the detail information of Turbofs access group.
---

# cloud_turbofs_access_groups

Use this data source to query the detail information of Turbofs access group.

## Example Usage

```hcl
data "cloud_turbofs_access_groups" "access_groups" {
  access_group_id = "pgroup-7nx89k7l"
  name            = "test"
}
```

## Argument Reference

The following arguments are supported:

* `access_group_id` - (Optional, String) A specified access group ID used to query.
* `name` - (Optional, String) A access group Name used to query.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `access_group_list` - An information list of Turbofs access group. Each element contains the following attributes:
  * `access_group_id` - ID of the access group.
  * `create_time` - Creation time of the access group.
  * `description` - Description of the access group.
  * `name` - Name of the access group.



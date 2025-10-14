---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_address_template_groups"
sidebar_current: "docs-cloud-datasource-vpc_address_template_groups"
description: |-
  Use this data source to query detailed information of address template groups.
---

# cloud_vpc_address_template_groups

Use this data source to query detailed information of address template groups.

## Example Usage

```hcl
data "cloud_vpc_address_template_groups" "name" {
  name = "test"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional, String) Id of the address template group to query.
* `name` - (Optional, String) Name of the address template group to query.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `group_list` - Information list of the dedicated address template groups.
  * `id` - Id of the address template group.
  * `name` - Name of address template group.
  * `template_ids` - ID set of the address template.



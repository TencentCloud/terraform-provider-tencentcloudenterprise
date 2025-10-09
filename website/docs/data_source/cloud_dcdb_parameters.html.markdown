---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_parameters"
sidebar_current: "docs-cloud-datasource-dcdb_parameters"
description: |-
  Use this data source to query detailed information of dcdb parameters
---

# cloud_dcdb_parameters

Use this data source to query detailed information of dcdb parameters

## Example Usage

```hcl
data "cloud_dcdb_parameters" "parameters" {
  instance_id = "your_instance_id"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) Instance id.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - Parameter list.
  * `constraint` - Params constraint.
    * `enum` - A list of optional values of type num.
    * `range` - Range constraint.
      * `max` - Max value.
      * `min` - Min value.
    * `string` - Constraint type is string.
    * `type` - Type.
  * `default` - Default value.
  * `have_set_value` - Have set value.
  * `need_restart` - Need restart.
  * `param` - Parameter name.
  * `value` - Parameter value.



---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_database_objects"
sidebar_current: "docs-cloud-datasource-dcdb_database_objects"
description: |-
  Use this data source to query detailed information of dcdb database_objects
---

# cloud_dcdb_database_objects

Use this data source to query detailed information of dcdb database_objects

## Example Usage

```hcl
data "cloud_dcdb_database_objects" "database_objects" {
  instance_id        = "tdsqlshard-973xatu3"
  db_name            = "SysDB"
  result_output_file = "database_objects.json"
}
```

## Argument Reference

The following arguments are supported:

* `db_name` - (Required, String) Database name, obtained through the DescribeDatabases api.
* `instance_id` - (Required, String) The ID of instance.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `funcs` - Function list.
  * `func` - The name of function.
* `procs` - Procedure list.
  * `proc` - The name of procedure.
* `tables` - Table list.
  * `table` - The name of table.
* `views` - View list.
  * `view` - The name of view.



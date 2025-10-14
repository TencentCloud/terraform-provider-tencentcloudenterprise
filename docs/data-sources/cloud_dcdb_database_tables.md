---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_database_tables"
sidebar_current: "docs-cloud-datasource-dcdb_database_tables"
description: |-
  Use this data source to query detailed information of dcdb database_tables
---

# cloud_dcdb_database_tables

Use this data source to query detailed information of dcdb database_tables

## Example Usage

```hcl
data "cloud_dcdb_database_tables" "database_tables" {
  instance_id        = "tdsqlshard-973xatu3"
  db_name            = "SysDB"
  table              = "StatusTable"
  result_output_file = "database_tables.json"
}
```

## Argument Reference

The following arguments are supported:

* `db_name` - (Required, String) Database name, obtained through the DescribeDatabases api.
* `instance_id` - (Required, String) The ID of instance.
* `table` - (Required, String) Table name, obtained through the DescribeDatabaseObjects api.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cols` - Column information.
  * `col` - The name of column.
  * `type` - Column type.



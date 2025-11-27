---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_mariadb_database_table"
sidebar_current: "docs-tencentcloudenterprise-datasource-mariadb_database_table"
description: |-
  Use this data source to query detailed information of mariadb database_table
---

# tencentcloudenterprise_mariadb_database_table

Use this data source to query detailed information of mariadb database_table

## Example Usage

```hcl
data "tencentcloudenterprise_mariadb_database_table" "database_table" {
  instance_id = "tdsql-e9tklsgz"
  db_name     = "mysql"
  table       = "server_cost"
}
```

## Argument Reference

The following arguments are supported:

* `db_name` - (Required, String) database name.
* `instance_id` - (Required, String) instance id.
* `table` - (Required, String) table name.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cols` - column list.
  * `col` - column name.
  * `type` - column type.



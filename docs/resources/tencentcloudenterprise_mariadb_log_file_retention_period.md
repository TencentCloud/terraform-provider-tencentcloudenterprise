---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_mariadb_log_file_retention_period"
sidebar_current: "docs-tencentcloudenterprise-resources-mariadb_log_file_retention_period"
description: |-
  Provides a resource to create a mariadb log_file_retention_period
---

# tencentcloudenterprise_mariadb_log_file_retention_period

Provides a resource to create a mariadb log_file_retention_period

## Example Usage

```hcl
resource "tencentcloudenterprise_mariadb_log_file_retention_period" "log_file_retention_period" {
  instance_id = "tdsql-4pzs5b67"
  days        = "8"
}
```

## Argument Reference

The following arguments are supported:

* `days` - (Required, Int) The number of days to save, cannot exceed 30.
* `instance_id` - (Required, String) instance id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

mariadb log_file_retention_period can be imported using the id, e.g.
```
$ terraform import tencentcloudenterprise_mariadb_log_file_retention_period.log_file_retention_period tdsql-4pzs5b67
```


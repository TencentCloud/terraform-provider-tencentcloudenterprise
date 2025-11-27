---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_mariadb_flush_binlog"
sidebar_current: "docs-tencentcloudenterprise-resources-mariadb_flush_binlog"
description: |-
  Provides a resource to create a mariadb flush_binlog
---

# tencentcloudenterprise_mariadb_flush_binlog

Provides a resource to create a mariadb flush_binlog

## Example Usage

```hcl
resource "tencentcloudenterprise_mariadb_flush_binlog" "flush_binlog" {
  instance_id = "tdsql-9vqvls95"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Instance ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




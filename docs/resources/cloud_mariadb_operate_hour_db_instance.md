---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_mariadb_operate_hour_db_instance"
sidebar_current: "docs-cloud-resources-mariadb_operate_hour_db_instance"
description: |-
  Provides a resource to create a mariadb activate_hour_db_instance
---

# cloud_mariadb_operate_hour_db_instance

Provides a resource to create a mariadb activate_hour_db_instance

## Example Usage

```hcl
resource "cloud_mariadb_operate_hour_db_instance" "activate_hour_db_instance" {
  instance_id = "tdsql-9vqvls95"
  operate     = "activate"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) Instance ID.
* `operate` - (Required, String) Operation, `activate`- activate the hour db instance, `isolate`- isolate the hour db instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




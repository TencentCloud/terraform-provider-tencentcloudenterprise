---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_mariadb_renew_instance"
sidebar_current: "docs-cloud-resources-mariadb_renew_instance"
description: |-
  Provides a resource to create a mariadb renew_instance
---

# cloud_mariadb_renew_instance

Provides a resource to create a mariadb renew_instance

## Example Usage

```hcl
resource "cloud_mariadb_renew_instance" "renew_instance" {
  instance_id = "tdsql-9vqvls95"
  period      = 1
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Instance ID.
* `period` - (Required, Int, ForceNew) Renewal duration, unit: month.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




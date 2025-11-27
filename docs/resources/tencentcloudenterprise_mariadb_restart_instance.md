---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_mariadb_restart_instance"
sidebar_current: "docs-tencentcloudenterprise-resources-mariadb_restart_instance"
description: |-
  Provides a resource to create a mariadb restart_instance
---

# tencentcloudenterprise_mariadb_restart_instance

Provides a resource to create a mariadb restart_instance

## Example Usage

```hcl
resource "tencentcloudenterprise_mariadb_restart_instance" "restart_instance" {
  instance_id = "tdsql-9vqvls95"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) instance ID.
* `restart_time` - (Optional, String, ForceNew) expected restart time.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_mariadb_switch_ha"
sidebar_current: "docs-cloud-resources-mariadb_switch_ha"
description: |-
  Provides a resource to create a mariadb switch_h_a
---

# cloud_mariadb_switch_ha

Provides a resource to create a mariadb switch_h_a

## Example Usage

```hcl
resource "cloud_mariadb_switch_ha" "switch_ha" {
  instance_id = "tdsql-9vqvls95"
  zone        = "ap-guangzhou-2"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Instance ID in the format of tdsql-ow728lmc.
* `zone` - (Required, String, ForceNew) Target AZ. The node with the lowest delay in the target AZ will be automatically promoted to primary node.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




---
subcategory: "MARIADB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_mariadb_hour_db_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-mariadb_hour_db_instance"
description: |-
  Manages a pay-as-you-go MariaDB instance.
---

# tencentcloudenterprise_mariadb_hour_db_instance

Manages a pay-as-you-go MariaDB instance.

## Example Usage

```hcl
variable "zones" { type = set(string) }

resource "tencentcloudenterprise_mariadb_hour_db_instance" "example" {
  memory     = 2
  node_count = 2
  storage    = 10
  zones      = var.zones
}
```

## Argument Reference

The following arguments are supported:

* `memory` - (Required, Int) instance memory.
* `node_count` - (Required, Int) number of node for instance.
* `storage` - (Required, Int) instance disk storage.
* `zones` - (Required, Set: [`String`]) available zone of instance.
* `cpu_arch` - (Optional, String, ForceNew) CPU architecture of the instance. Valid values: `X86`, `ARM`.
* `db_version_id` - (Optional, String) Database engine version. Use DescribeDBEngines to find MariaDB versions available in the region.
* `instance_name` - (Optional, String) name of this instance.
* `project_id` - (Optional, Int) project id.
* `subnet_id` - (Optional, String) subnet id, it's required when vpc_id is set.
* `tags` - (Optional, Map) Tag description list.
* `vip` - (Optional, String) vip.
* `vpc_id` - (Optional, String) vpc id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



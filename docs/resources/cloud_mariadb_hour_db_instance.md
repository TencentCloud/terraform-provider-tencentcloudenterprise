---
subcategory: "TencentDB for MariaDB(MariaDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_mariadb_hour_db_instance"
sidebar_current: "docs-cloud-resources-mariadb_hour_db_instance"
description: |-
  Provides a resource to create a mariadb hour_db_instance
---

# cloud_mariadb_hour_db_instance

Provides a resource to create a mariadb hour_db_instance

## Example Usage

```hcl
resource "cloud_mariadb_hour_db_instance" "basic" {
  db_version_id = "10.0"
  instance_name = "db-test-del"
  memory        = 2
  node_count    = 2
  storage       = 10
  subnet_id     = "subnet-jdi5xn22"
  vpc_id        = "vpc-k1t8ickr"
  vip           = "10.0.0.197"
  zones = [
    "ap-guangzhou-6",
    "ap-guangzhou-7",
  ]
  tags = {
    createdBy = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `memory` - (Required, Int) instance memory.
* `node_count` - (Required, Int) number of node for instance.
* `storage` - (Required, Int) instance disk storage.
* `zones` - (Required, Set: [`String`]) available zone of instance.
* `db_version_id` - (Optional, String) db engine version, default to 10.1.9.
* `instance_name` - (Optional, String) name of this instance.
* `project_id` - (Optional, Int) project id.
* `subnet_id` - (Optional, String) subnet id, it&amp;#39;s required when vpcId is set.
* `tags` - (Optional, Map) Tag description list.
* `vip` - (Optional, String) vip.
* `vpc_id` - (Optional, String) vpc id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

mariadb hour_db_instance can be imported using the id, e.g.
```
$ terraform import cloud_mariadb_hour_db_instance.hour_db_instance tdsql-kjqih9nn
```


---
subcategory: "TDSQL PostgreSQL (Tbase)"
layout: "cloud"
page_title: "TencentCloud: cloud_tbase_pg_instance"
sidebar_current: "docs-cloud-resource-tbase_pg_instance"
description: |-
  Provides a resource to create a tbase instance.
---

# cloud_tbase_pg_instance

Provides a resource to create a tbase instance.

~> **NOTE:** This resource is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.

## Example Usage

```hcl
resource "cloud_tbase_pg_instance" "example" {
  vpc_id                 = "vpc-cs6ffr73"
  subnet_id              = "subnet-38oi34ta"
  net_type               = "vpc"
  pool_id                = 1
  instance_name          = "tf-test"
  engine_type            = "PostgreSQL"
  engine_version         = "10.5.2"
  memory                 = 1
  cpu                    = 1
  admin_password         = "Xx5kjgw%RXKf"
  charset                = "UTF8"
  instance_count         = 1
  instance_role          = "master"
  master_node_zone       = "yfm18"
  follow_nodes_zones     = ["yfm18"]
  security_group_id_list = ["sg-9s7k6qgw", "sg-rrpst23s"]
  pg_node {
    storage   = 10
    spec_code = "pg.dn.1g"
  }
  project_id = "pr-bae40f73"
}
```

## Argument Reference

The following arguments are supported:

* `charset` - (Required, String) charset.
* `cpu` - (Required, Int) cpu.
* `engine_type` - (Required, String) engine type.
* `engine_version` - (Required, String) engine version.
* `instance_count` - (Required, Int) instance count.
* `instance_role` - (Required, String) instance role, current only support `master` type.
* `master_node_zone` - (Required, String) The zone of the master node.
* `memory` - (Required, Int) memory.
* `net_type` - (Required, String) net type, optional value is VPC or BASIC.
* `pg_node` - (Required, List) Node specification information.
* `pool_id` - (Required, Int) pool id.
* `security_group_id_list` - (Required, Set: [`String`]) The security group id list of the instance.
* `subnet_id` - (Required, String) subnet id.
* `vpc_id` - (Required, String) vpc id.
* `admin_password` - (Optional, String) admin pass.
* `follow_nodes_zones` - (Optional, Set: [`String`]) The zone of the follow nodes.
* `instance_name` - (Optional, String) instance name.
* `project_id` - (Optional, String) project id.

The `pg_node` object supports the following:

* `spec_code` - (Required, String) The spec code of pg nodes to be created.
* `storage` - (Required, Int) The sync type of pg nodes to be created.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

tbase instance can be imported using the id, e.g.
```
$ terraform import cloud_tbase_instance.instance cluster_id#instance_id
```


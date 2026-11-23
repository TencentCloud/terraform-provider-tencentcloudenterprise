---
subcategory: "TDSQL PostgreSQL(Tbase)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tbase_pg_instance"
sidebar_current: "docs-tencentcloudenterprise-resource-tbase_pg_instance"
description: |-
  Manages a TDSQL PostgreSQL instance.
---

# tencentcloudenterprise_tbase_pg_instance

Manages a TDSQL PostgreSQL instance.

## Example Usage

```hcl
variable "engine_type" { type = string }
variable "engine_version" { type = string }
variable "master_node_zone" { type = string }
variable "pool_id" { type = number }
variable "security_group_id_list" { type = set(string) }
variable "subnet_id" { type = string }
variable "vpc_id" { type = string }
variable "pg_spec_code" { type = string }

resource "tencentcloudenterprise_tbase_pg_instance" "example" {
  charset                = "UTF8"
  cpu                    = 1
  engine_type            = var.engine_type
  engine_version         = var.engine_version
  instance_count         = 1
  instance_role          = "master"
  master_node_zone       = var.master_node_zone
  memory                 = 1
  net_type               = "vpc"
  pool_id                = var.pool_id
  security_group_id_list = var.security_group_id_list
  subnet_id              = var.subnet_id
  vpc_id                 = var.vpc_id
  pg_node {
    spec_code = var.pg_spec_code
    storage   = 10
  }
}
```

## Argument Reference

The following arguments are supported:

* `charset` - (Required, String) charset.
* `cpu` - (Required, Int) cpu.
* `engine_type` - (Required, String) PostgreSQL engine type. Specify PostgreSQL or TbaseXC according to the engine type and version offered in the region.
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



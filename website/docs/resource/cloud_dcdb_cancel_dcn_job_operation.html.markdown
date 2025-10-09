---
subcategory: "Distributed Database For MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_cancel_dcn_job_operation"
sidebar_current: "docs-cloud-resource-dcdb_cancel_dcn_job_operation"
description: |-
  Provides a resource to create a dcdb cancel_dcn_job_operation
---

# cloud_dcdb_cancel_dcn_job_operation

Provides a resource to create a dcdb cancel_dcn_job_operation

## Example Usage

```hcl
data "cloud_vpc_security_groups" "internal" {
  name = "default"
}

data "cloud_vpc_instances" "vpc" {
  name = "Default-VPC"
}

data "cloud_vpc_subnets" "subnet" {
  vpc_id = data.cloud_vpc_instances.vpc.instance_list.0.vpc_id
}

locals {
  vpc_id    = data.cloud_vpc_subnets.subnet.instance_list.0.vpc_id
  subnet_id = data.cloud_vpc_subnets.subnet.instance_list.0.subnet_id
  sg_id     = data.cloud_vpc_security_groups.internal.security_groups.0.security_group_id
}

resource "cloud_dcdb_instance" "hourdb_instance_dcn" {
  instance_name     = "test_dcdb_db_hourdb_instance_dcn"
  zones             = [var.default_az]
  shard_memory      = "2"
  shard_storage     = "10"
  shard_node_count  = "2"
  shard_count       = "2"
  vpc_id            = local.vpc_id
  subnet_id         = local.subnet_id
  security_group_id = local.sg_id
  db_version_id     = "8.0"
  dcn_region        = "ap-guangzhou"
  dcn_instance_id   = local.dcdb_id //master_instance
  resource_tags {
    tag_key   = "aaa"
    tag_value = "bbb"
  }
}

locals {
  dcn_dcdb_id = cloud_dcdb_instance.hourdb_instance_dcn.id
}

resource "cloud_dcdb_cancel_dcn_job_operation" "cancel_operation" {
  instance_id = local.dcn_dcdb_id // cancel the dcn/readonly instance
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) Instance ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




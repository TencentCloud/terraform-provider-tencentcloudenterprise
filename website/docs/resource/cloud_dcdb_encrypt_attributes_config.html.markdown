---
subcategory: "TDSQL for MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_encrypt_attributes_config"
sidebar_current: "docs-cloud-resource-dcdb_encrypt_attributes_config"
description: |-
  Provides a resource to create a dcdb encrypt_attributes_config
---

# cloud_dcdb_encrypt_attributes_config

Provides a resource to create a dcdb encrypt_attributes_config

~> **NOTE:**  This resource currently only supports the newly created MySQL 8.0.24 version.

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

resource "cloud_dcdb_db_instance" "prepaid_instance" {
  instance_name    = "test_dcdb_db_post_instance"
  zones            = [var.default_az]
  period           = 1
  shard_memory     = "2"
  shard_storage    = "10"
  shard_node_count = "2"
  shard_count      = "2"
  vpc_id           = local.vpc_id
  subnet_id        = local.subnet_id
  db_version_id    = "8.0"
  resource_tags {
    tag_key   = "aaa"
    tag_value = "bbb"
  }
  security_group_ids = [local.sg_id]
}

resource "cloud_dcdb_instance" "hourdb_instance" {
  instance_name     = "test_dcdb_db_hourdb_instance"
  zones             = [var.default_az]
  shard_memory      = "2"
  shard_storage     = "10"
  shard_node_count  = "2"
  shard_count       = "2"
  vpc_id            = local.vpc_id
  subnet_id         = local.subnet_id
  security_group_id = local.sg_id
  db_version_id     = "8.0"
  resource_tags {
    tag_key   = "aaa"
    tag_value = "bbb"
  }
}

locals {
  prepaid_dcdb_id = cloud_dcdb_db_instance.prepaid_instance.id
  hourdb_dcdb_id  = cloud_dcdb_instance.hourdb_instance.id
}

// for postpaid instance

resource "cloud_dcdb_encrypt_attributes_config" "config_hourdb" {
  instance_id     = local.hourdb_dcdb_id
  encrypt_enabled = 1
}

// for prepaid instance

resource "cloud_dcdb_encrypt_attributes_config" "config_prepaid" {
  instance_id     = local.prepaid_dcdb_id
  encrypt_enabled = 1
}
```

### # Import

dcdb encrypt_attributes_config can be imported using the id, e.g.

```hcl
terraform import cloud_dcdb_encrypt_attributes_config.encrypt_attributes_config encrypt_attributes_config_id
```

## Argument Reference

The following arguments are supported:

* `encrypt_enabled` - (Required, Int) whether to enable data encryption. Notice: it is not supported to turn it off after it is turned on. The optional values: 0-disable, 1-enable.
* `instance_id` - (Required, String) instance id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




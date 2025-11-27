---
subcategory: "Cloud Distributed SQL For PostGreSQL(TBASE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tbase_instance"
sidebar_current: "docs-tencentcloudenterprise-resources-tbase_instance"
description: |-
  Provides a resource to create a tbase instance.
---

# tencentcloudenterprise_tbase_instance

Provides a resource to create a tbase instance.

~> **NOTE:** This resource is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.

## Example Usage

```hcl
resource "tencentcloudenterprise_tbase_instance" "instance1" {
  admin_password         = "Tcdn@2007"
  charset                = "UTF8"
  engine_type            = "TbaseV5"
  engine_version         = "5.06.1.1"
  goods_num              = 1
  net_type               = "vpc"
  security_group_id_list = ["sg-h45zkk4g"]
  subnet_id              = "subnet-ebbhkayi"
  vpc_id                 = "vpc-1dz8zfpr"
  zone                   = "yfdb1"
  dn_node {
    node_count = 4
    set_count  = 2
    spec_code  = "dn.v5.1g"
    storage    = 10
    sync_type  = "sync2async"
  }
  cn_node {
    node_count = 4
    set_count  = 2
    spec_code  = "cn.v5.1g"
    storage    = 10
  }
}
```

## Argument Reference

The following arguments are supported:

* `admin_password` - (Required, String) Admin pass.
* `charset` - (Required, String) Character set, based on DescribeAvailableCharset.
* `dn_node` - (Required, List) Create the specification information related to the dn node of the instance.
* `engine_type` - (Required, String) Engine type, optional value is `TbaseV2`, `TbaseV5`, `TbaseV5_C`, `TbaseH`,	`TbaseHC`.
* `engine_version` - (Required, String) Engine version.
* `goods_num` - (Required, Int) Purchase quantity, (0,10), no default create 1.
* `net_type` - (Required, String) Net type, optional value is VPC or BASIC.
* `security_group_id_list` - (Required, Set: [`String`]) The security group id list of the instance.
* `subnet_id` - (Required, String) Subnet id.
* `vpc_id` - (Required, String) Vpc id.
* `zone` - (Required, String) Zone.
* `cn_node` - (Optional, List) Example Create specifications of the instance cn node.
* `instance_name` - (Optional, String) Instance name.
* `project_id` - (Optional, String) Project id.
* `recovery` - (Optional, List) Mandatory parameter for creating a backfile instance.
* `slave_zones` - (Optional, List: [`String`]) The list of slave zones.
* `tags` - (Optional, Map) Tag list.

The `cn_node` object supports the following:

* `node_count` - (Required, Int) The number of cn nodes to be created.
* `set_count` - (Required, Int) The number of cn nodes to be created.
* `spec_code` - (Required, String) The spec code of cn nodes to be created.
* `storage` - (Required, Int) The stroage of cn nodes to be created.
* `sync_type` - (Optional, String) The sync type of cn nodes to be created.

The `dn_node` object supports the following:

* `node_count` - (Required, Int) The number of nodes to be created.
* `set_count` - (Required, Int) The number of nodes to be created.
* `spec_code` - (Required, String) The spec code of nodes to be created.
* `storage` - (Required, Int) The storage of nodes to be created.
* `sync_type` - (Optional, String) The sync type of nodes to be created.

The `recovery` object supports the following:

* `recovery_instance_id` - (Required, String) The recovery instance id of the instance.
* `recovery_time_point` - (Required, String) The recovery time point of the instance.
* `recovery_time` - (Required, String) The recovery time of the instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

tbase instance can be imported using the id, e.g.
```
$ terraform import tencentcloudenterprise_tbase_instance.instance cluster_id#instance_id
```


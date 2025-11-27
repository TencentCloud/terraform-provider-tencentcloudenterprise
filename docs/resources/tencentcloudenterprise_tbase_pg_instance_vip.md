---
subcategory: "Cloud Distributed SQL For PostGreSQL(TBASE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tbase_pg_instance_vip"
sidebar_current: "docs-tencentcloudenterprise-resources-tbase_pg_instance_vip"
description: |-
  Provides a resource to create a tbase instance.
---

# tencentcloudenterprise_tbase_pg_instance_vip

Provides a resource to create a tbase instance.

~> **NOTE:** This resource is still in internal testing. To experience its functions, you need to apply for a whitelist from Tencent Cloud.

## Example Usage

```hcl
resource "tencentcloudenterprise_tbase_pg_instance_vip" "foo" {
  instance_id    = tencentcloudenterprise_tbase_pg_instance.example.id
  uniq_subnet_id = "subnet-38oi34ta"
  uniq_vpc_id    = "vpc-cs6ffr73"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) Instance id.
* `uniq_subnet_id` - (Required, String) Subnet id.
* `uniq_vpc_id` - (Required, String) Vpc id.
* `vip` - (Optional, String) Vip.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `vport` - Vport.


## Import

tbase instance can be imported using the id, e.g.
```
$ terraform import tencentcloudenterprise_tbase_pg_instance_vip.instance cluster_id#instance_id
```


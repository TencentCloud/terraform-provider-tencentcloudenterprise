---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_route"
sidebar_current: "docs-cloud-resources-tdmq_route"
description: |-
  Provide a resource to create a TDMQ Route.
---

# cloud_tdmq_route

Provide a resource to create a TDMQ Route.

## Example Usage

```hcl
resource "cloud_tdmq_Route" "foo" {
  remark     = "this is description111."
  cluster_id = 0
  net_type   = 2
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) The Dedicated Cluster Id.
* `unique_subnet_id` - (Required, String) uniq subnet id, the value contains a maximum of 128 characters.
* `unique_vpc_id` - (Required, String) Uniq vpc id.
* `encrypt_type` - (Optional, Int) Encrypt type of route. 0: no encryption, 1: TLS encrypt. Default: 0.
* `net_type` - (Optional, Int) Net type.
* `remark` - (Optional, String) Remark of the route.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `router_id` - The id of the route.
* `vip` - The vip of the route.
* `vport` - The vport of the route.


## Import

Tdmq Route can be imported, e.g.

```
$ terraform import cloud_tdmq_Route.test tdmq_id
```


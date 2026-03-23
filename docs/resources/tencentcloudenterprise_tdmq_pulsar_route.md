---
subcategory: "TDMQ for Pulsar(tpulsar)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_route"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_pulsar_route"
description: |-
  Provide a resource to create a TDMQ Pulsar Route.
---

# tencentcloudenterprise_tdmq_pulsar_route

Provide a resource to create a TDMQ Pulsar Route.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_route" "foo" {
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

tencentcloudenterprise_tdmq_pulsar_route can be imported using the id, e.g.

```
Tdmq Route can be imported, e.g.

```
$ terraform import tencentcloudenterprise_tdmq_pulsar_route.test tdmq_id
```
```


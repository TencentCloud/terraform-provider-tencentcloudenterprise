---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_dc_gateway"
sidebar_current: "docs-cloud-resources-vpc_dc_gateway"
description: |-
  Provides a resource to creating direct connect gateway instance.
---

# cloud_vpc_dc_gateway

Provides a resource to creating direct connect gateway instance.

## Example Usage

```hcl
resource "cloud_vpc" "main" {
  name       = "ci-vpc-instance-test"
  cidr_block = "10.0.0.0/16"
}

resource "cloud_vpc_dc_gateway" "vpc_main" {
  name                = "ci-cdg-vpc-test"
  network_instance_id = cloud_vpc.main.id
  network_type        = "VPC"
  gateway_type        = "NAT"
}
```

## Argument Reference

The following arguments are supported:

* `gateway_type` - (Required, String) Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`.
* `name` - (Required, String) Name of the DCG.
* `network_instance_id` - (Required, String, ForceNew) If the `network_type` value is `VPC`, the available value is VPC ID.
* `network_type` - (Required, String, ForceNew) Type of associated network. Valid value: `VPC`.
* `band_with` - (Optional, Int) The bandwith speed limit of the gateway.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `cnn_route_type` - Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP enabled.
* `create_time` - Creation time of resource.
* `enable_bgp` - Indicates whether the BGP is enabled.


## Import

Direct connect gateway instance can be imported, e.g.

```
$ terraform import cloud_vpc_dc_gateway.instance dcg-id
```


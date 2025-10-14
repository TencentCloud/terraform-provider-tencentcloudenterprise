---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_nat_gateway"
sidebar_current: "docs-cloud-resources-vpc_nat_gateway"
description: |-
  Provides a resource to create a NAT gateway.
---

# cloud_vpc_nat_gateway

Provides a resource to create a NAT gateway.

## Example Usage

```hcl
resource "cloud_vpc_nat_gateway" "foo" {
  name             = "test_nat_gateway"
  vpc_id           = "vpc-4xxr2cy7"
  bandwidth        = 100
  max_concurrent   = 1000000
  assigned_eip_set = ["1.1.1.1"]

  tags = {
    test = "tf"
  }
}
```

## Argument Reference

The following arguments are supported:

* `assigned_eip_set` - (Required, Set: [`String`]) EIP IP address set bound to the gateway. The value of at least 1 and at most 10.
* `name` - (Required, String) Name of the NAT gateway.
* `vpc_id` - (Required, String, ForceNew) ID of the vpc.
* `bandwidth` - (Optional, Int) The maximum public network output bandwidth of NAT gateway (unit: Mbps). Valid values: `20`, `50`, `100`, `200`, `500`, `1000`, `2000`, `5000`. Default is 100.
* `internet_service_provider` - (Optional, String) The ISP of the NAT gateway. Valid values: `CTCC` (China Telecom), `CUCC` (China Unicom), `CMCC` (China Mobile). Default is `CTCC`.
* `max_concurrent` - (Optional, Int) The upper limit of concurrent connection of NAT gateway. Valid values: `1000000`, `3000000`, `10000000`. Default is `1000000`.
* `tags` - (Optional, Map) The available tags within this NAT gateway.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `created_time` - Create time of the NAT gateway.


## Import

NAT gateway can be imported using the id, e.g.

```
$ terraform import cloud_vpc_nat_gateway.foo nat-1asg3t63
```


---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_route_entry"
sidebar_current: "docs-cloud-resources-vpc_route_entry"
description: |-
  Provides a resource to create a routing entry in a VPC routing table.
---

# cloud_vpc_route_entry

Provides a resource to create a routing entry in a VPC routing table.

~> **NOTE:** It has been deprecated and replaced by cloud_route_table_entry.

## Example Usage

```hcl
resource "cloud_vpc" "main" {
  name       = "Used to test the routing entry"
  cidr_block = "10.4.0.0/16"
}

resource "cloud_route_table" "r" {
  name   = "Used to test the routing entry"
  vpc_id = cloud_vpc.main.id
}

resource "cloud_route_entry" "rtb_entry_instance" {
  vpc_id         = cloud_route_table.main.vpc_id
  route_table_id = cloud_route_table.r.id
  cidr_block     = "10.4.8.0/24"
  next_type      = "instance"
  next_hub       = "10.16.1.7"
}

resource "cloud_route_entry" "rtb_entry_instance" {
  vpc_id         = cloud_route_table.main.vpc_id
  route_table_id = cloud_route_table.r.id
  cidr_block     = "10.4.5.0/24"
  next_type      = "vpn_gateway"
  next_hub       = "vpngw-db52irtl"
}
```

## Argument Reference

The following arguments are supported:

* `cidr_block` - (Required, String, ForceNew) The RouteEntry's target network segment.
* `next_hub` - (Required, String, ForceNew) The route entry's next hub. CVM instance ID or VPC router interface ID.
* `next_type` - (Required, String, ForceNew) The next hop type. Valid values: `public_gateway`,`vpn_gateway`,`sslvpn_gateway`,`dc_gateway`,`peering_connection`,`nat_gateway`,`havip`,`local_gateway` and `instance`. `instance` points to CVM Instance.
* `route_table_id` - (Required, String, ForceNew) The ID of the route table.
* `vpc_id` - (Required, String, ForceNew) The VPC ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




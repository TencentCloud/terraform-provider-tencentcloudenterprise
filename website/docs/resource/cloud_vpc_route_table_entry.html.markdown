---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_route_table_entry"
sidebar_current: "docs-cloud-resource-vpc_route_table_entry"
description: |-
  Provides a resource to create an entry of a routing table.
---

# cloud_vpc_route_table_entry

Provides a resource to create an entry of a routing table.

## Example Usage

```hcl
variable "availability_zone" {
  default = "na-siliconvalley-1"
}

resource "cloud_vpc" "foo" {
  name       = "ci-temp-test"
  cidr_block = "10.0.0.0/16"
}

resource "cloud_vpc_subnet" "foo" {
  vpc_id            = cloud_vpc.foo.id
  name              = "terraform test subnet"
  cidr_block        = "10.0.12.0/24"
  availability_zone = var.availability_zone
  route_table_id    = cloud_route_table.foo.id
}

resource "cloud_route_table" "foo" {
  vpc_id = cloud_vpc.foo.id
  name   = "ci-temp-test-rt"
}

resource "cloud_vpc_route_table_entry" "instance" {
  route_table_id         = cloud_route_table.foo.id
  destination_cidr_block = "10.4.4.0/24"
  next_type              = "EIP"
  next_hop               = "0"
  description            = "ci-test-route-table-entry"
}
```

## Argument Reference

The following arguments are supported:

* `destination_cidr_block` - (Required, String, ForceNew) Destination address block.
* `next_hop` - (Required, String, ForceNew) ID of next-hop gateway. Note: when `next_type` is EIP, GatewayId should be `0`.
* `next_type` - (Required, String, ForceNew) Type of next-hop. Valid values: `VPN`, `DIRECTCONNECT`, `PEERCONNECTION`, `HAVIP`, `NAT`, `NORMAL_CVM`, `EIP`.
* `route_table_id` - (Required, String, ForceNew) ID of routing table to which this entry belongs.
* `description` - (Optional, String, ForceNew) Description of the routing table entry.
* `disabled` - (Optional, Bool) Whether the entry is disabled, default is `false`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Route table entry can be imported using the id, e.g.

```
$ terraform import cloud_vpc_route_table_entry.foo 83517.rtb-mlhpg09u
```


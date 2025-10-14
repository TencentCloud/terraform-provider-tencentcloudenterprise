---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_dnat"
sidebar_current: "docs-cloud-resources-vpc_dnat"
description: |-
  Provides a resource to create a NAT forwarding.
---

# cloud_vpc_dnat

Provides a resource to create a NAT forwarding.

## Example Usage

```hcl
resource "cloud_vpc_dnat" "foo" {
  vpc_id       = "vpc-asg3sfa3"
  nat_id       = "nat-2515tdg"
  protocol     = "tcp"
  elastic_ip   = "203.0.113.1"
  elastic_port = 80
  private_ip   = "203.0.113.2"
  private_port = 22
  description  = "test"
}
```

## Argument Reference

The following arguments are supported:

* `elastic_ip` - (Required, String, ForceNew) Network address of the EIP.
* `elastic_port` - (Required, String, ForceNew) Port of the EIP.
* `nat_id` - (Required, String, ForceNew) ID of the NAT gateway.
* `private_ip` - (Required, String, ForceNew) Network address of the backend service.
* `private_port` - (Required, String, ForceNew) Port of intranet.
* `protocol` - (Required, String, ForceNew) Type of the network protocol. Valid value: `TCP` and `UDP`.
* `vpc_id` - (Required, String, ForceNew) ID of the VPC.
* `description` - (Optional, String) Description of the NAT forward.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

NAT forwarding can be imported using the id, e.g.

```
$ terraform import cloud_vpc_dnat.foo tcp://vpc-asg3sfa3:nat-1asg3t63@127.15.2.3:8080
```


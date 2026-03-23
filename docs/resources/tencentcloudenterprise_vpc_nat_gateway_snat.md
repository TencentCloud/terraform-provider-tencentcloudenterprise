---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_nat_gateway_snat"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_nat_gateway_snat"
description: |-
  Provides a resource to create a NAT Gateway SNat rule.
---

# tencentcloudenterprise_vpc_nat_gateway_snat

Provides a resource to create a NAT Gateway SNat rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_nat_gateway_snat" "subnet_snat" {
  nat_gateway_id    = tencentcloudenterprise_vpc_nat_gateway.my_nat.id
  resource_type     = "SUBNET"
  subnet_id         = tencentcloudenterprise_vpc_subnet.my_subnet.id
  subnet_cidr_block = tencentcloudenterprise_vpc_subnet.my_subnet.cidr_block
  description       = "terraform test"
  public_ip_addr = [
    tencentcloudenterprise_eip.eip1.public_ip,
    tencentcloudenterprise_eip.eip2.public_ip,
  ]
}

resource "tencentcloudenterprise_vpc_nat_gateway_snat" "instance_snat" {
  nat_gateway_id           = tencentcloudenterprise_vpc_nat_gateway.my_nat.id
  resource_type            = "NETWORKINTERFACE"
  instance_id              = tencentcloudenterprise_cvm_instance.my_instance.id
  instance_private_ip_addr = tencentcloudenterprise_cvm_instance.my_instance.private_ip
  description              = "terraform test"
  public_ip_addr = [
    tencentcloudenterprise_eip.eip1.public_ip,
  ]
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required, String) Description.
* `nat_gateway_id` - (Required, String, ForceNew) NAT gateway ID.
* `public_ip_addr` - (Required, List: [`String`]) Elastic IP address pool.
* `resource_type` - (Required, String, ForceNew) Resource type. Valid values: SUBNET, NETWORKINTERFACE.
* `instance_id` - (Optional, String, ForceNew) Instance ID, required when `resource_type` is NETWORKINTERFACE.
* `instance_private_ip_addr` - (Optional, String, ForceNew) Private IPs of the instance's primary ENI, required when `resource_type` is NETWORKINTERFACE.
* `subnet_cidr_block` - (Optional, String, ForceNew) The IPv4 CIDR of the subnet, required when `resource_type` is SUBNET.
* `subnet_id` - (Optional, String, ForceNew) Subnet instance ID, required when `resource_type` is SUBNET.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time.
* `snat_id` - SNAT rule ID.

## Import

tencentcloudenterprise_vpc_nat_gateway_snat can be imported using the id, e.g.

```
NAT gateway snat rule can be imported using the id, the id format must be '{nat_gateway_id}#{resource_id}',
resource_id range `subnet_id`, `instance_id`, e.g.

SUBNET SNat
```
$ terraform import tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat nat-r4ip1cwt#subnet-2ap74y35
```

NETWORKINTERFACE SNat
```
$ terraform import tencentcloudenterprise_vpc_nat_gateway_snat.instance_snat nat-r4ip1cwt#ins-da412f5a
```
```


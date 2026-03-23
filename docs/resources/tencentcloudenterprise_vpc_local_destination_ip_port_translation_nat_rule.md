---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_local_destination_ip_port_translation_nat_rule"
description: |-
  Provides a resource to creating VPC local destination IP port translation NAT rule.
---

# tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule

Provides a resource to creating VPC local destination IP port translation NAT rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc" "main" {
  name       = "ci-vpc-instance-test"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_dc_gateway" "dcg_main" {
  name                = "ci-dcg-test"
  network_instance_id = tencentcloudenterprise_vpc.main.id
  network_type        = "VPC"
  gateway_type        = "NAT"
}

resource "tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule" "nat_rule" {
  vpc_id                    = tencentcloudenterprise_vpc.main.id
  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
  protocol                  = "tcp"
  original_ip               = "10.0.1.1"
  original_port             = 80
  translation_ip            = "10.0.2.1"
  translation_port          = 8080
  description               = "test destination nat rule"
}
```

## Argument Reference

The following arguments are supported:

* `direct_connect_gateway_id` - (Required, String, ForceNew) Direct connect gateway ID.
* `original_ip` - (Required, String) Original IP address.
* `original_port` - (Required, Int) Original port.
* `protocol` - (Required, String) Protocol type. Valid values: `tcp`, `udp`.
* `translation_ip` - (Required, String) Translation IP address.
* `translation_port` - (Required, Int) Translation port.
* `vpc_id` - (Required, String, ForceNew) VPC instance ID.
* `description` - (Optional, String) Description.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule can be imported using the id, e.g.

```
VPC local destination IP port translation NAT rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule.instance vpc-id#dcg-id#protocol#original-ip#original-port
```
```


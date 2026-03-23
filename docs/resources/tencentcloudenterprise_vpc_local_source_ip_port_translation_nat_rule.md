---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_local_source_ip_port_translation_nat_rule"
description: |-
  Provides a resource to creating VPC local source IP port translation NAT rule.
---

# tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule

Provides a resource to creating VPC local source IP port translation NAT rule.

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

resource "tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule" "nat_rule" {
  vpc_id                    = tencentcloudenterprise_vpc.main.id
  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
  ip_pool                   = "10.40.31.45"
  description               = "test nat rule"
}
```

## Argument Reference

The following arguments are supported:

* `direct_connect_gateway_id` - (Required, String, ForceNew) Direct connect gateway ID.
* `ip_pool` - (Required, String) IP pool.
* `vpc_id` - (Required, String, ForceNew) VPC instance ID.
* `description` - (Optional, String) Description.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule can be imported using the id, e.g.

```
VPC local source IP port translation NAT rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule.instance vpc-id#dcg-id#ip-pool
```
```


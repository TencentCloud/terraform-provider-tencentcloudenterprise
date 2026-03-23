---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_local_ip_translation_acl_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_local_ip_translation_acl_rule"
description: |-
  Provides a resource to creating VPC local IP translation ACL rule.
---

# tencentcloudenterprise_vpc_local_ip_translation_acl_rule

Provides a resource to creating VPC local IP translation ACL rule.

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

resource "tencentcloudenterprise_vpc_local_ip_translation_nat_rule" "nat_rule" {
  vpc_id                    = tencentcloudenterprise_vpc.main.id
  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
  original_ip               = "8.123.40.13"
  translation_ip            = "8.123.45.14"
  description               = "test nat rule"
}

resource "tencentcloudenterprise_vpc_local_ip_translation_acl_rule" "main" {
  vpc_id                    = tencentcloudenterprise_vpc.main.id
  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
  original_ip               = tencentcloudenterprise_vpc_local_ip_translation_nat_rule.nat_rule.original_ip
  translation_ip            = tencentcloudenterprise_vpc_local_ip_translation_nat_rule.nat_rule.translation_ip
  protocol                  = "all"
  source_port               = "0"
  destination_port          = "0"
  destination_cidr          = "10.0.0.0/30"
}
```

## Argument Reference

The following arguments are supported:

* `destination_cidr` - (Required, String) Destination CIDR.
* `destination_port` - (Required, String) Destination port. Use `0` for all ports.
* `direct_connect_gateway_id` - (Required, String, ForceNew) Direct connect gateway ID.
* `original_ip` - (Required, String, ForceNew) Original IP address.
* `protocol` - (Required, String) Protocol type. Valid values: `all`, `tcp`, `udp`.
* `source_port` - (Required, String) Source port. Use `0` for all ports.
* `translation_ip` - (Required, String, ForceNew) Translation IP address.
* `vpc_id` - (Required, String, ForceNew) VPC instance ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `acl_rule_id` - ACL rule ID.

## Import

tencentcloudenterprise_vpc_local_ip_translation_acl_rule can be imported using the id, e.g.

```
VPC local IP translation ACL rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_local_ip_translation_acl_rule.instance vpc-id#dcg-id#original-ip#translation-ip#acl-rule-id
```
```


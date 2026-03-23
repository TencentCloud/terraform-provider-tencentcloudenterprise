---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_local_ip_translation_nat_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_local_ip_translation_nat_rule"
description: |-
  Provides a resource to creating VPC local IP translation NAT rule.
---

# tencentcloudenterprise_vpc_local_ip_translation_nat_rule

Provides a resource to creating VPC local IP translation NAT rule.

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

resource "tencentcloudenterprise_vpc_local_ip_translation_nat_rule" "main" {
  vpc_id                    = tencentcloudenterprise_vpc.main.id
  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
  original_ip               = "8.123.40.13"
  translation_ip            = "8.123.45.14"
  description               = "test local ip translation nat rule"
}
```

## Argument Reference

The following arguments are supported:

* `direct_connect_gateway_id` - (Required, String, ForceNew) Direct connect gateway ID.
* `original_ip` - (Required, String) Original IP address.
* `translation_ip` - (Required, String) Translation IP address.
* `vpc_id` - (Required, String, ForceNew) VPC instance ID.
* `description` - (Optional, String) Description of the rule.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_vpc_local_ip_translation_nat_rule can be imported using the id, e.g.

```
VPC local IP translation NAT rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_local_ip_translation_nat_rule.instance vpc-id#dcg-id#original-ip#translation-ip
```
```


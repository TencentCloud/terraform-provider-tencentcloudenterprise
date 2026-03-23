---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_vpc_firewall_switch"
sidebar_current: "docs-tencentcloudenterprise-resource-cfw_vpc_firewall_switch"
description: |-
  Provides a resource to create a cloud firewall (cfw) VPC firewall switch.
---

# tencentcloudenterprise_cfw_vpc_firewall_switch

Provides a resource to create a cloud firewall (cfw) VPC firewall switch.

## Example Usage

```hcl
resource "tencentcloudenterprise_cfw_vpc_firewall_switch" "example" {
  vpc_ins_id = "cfwg-xxxxxxxx"
  switch_id  = "switch-xxxxxxxx"
  enable     = 1
}
```

## Argument Reference

The following arguments are supported:

* `enable` - (Required, Int) Turn the switch on or off. 0: turn off the switch; 1: Turn on the switch.
* `switch_id` - (Required, String, ForceNew) Firewall switch ID.
* `vpc_ins_id` - (Required, String, ForceNew) Firewall instance id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cfw_vpc_firewall_switch can be imported using the id, e.g.

```
Cloud firewall VPC firewall switch can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_vpc_firewall_switch.example vpc_ins_id#switch_id
```
```


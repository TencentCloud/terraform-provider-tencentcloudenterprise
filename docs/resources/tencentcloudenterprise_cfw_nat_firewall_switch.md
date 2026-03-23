---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_nat_firewall_switch"
sidebar_current: "docs-tencentcloudenterprise-resource-cfw_nat_firewall_switch"
description: |-
  Provides a resource to create a cloud firewall (cfw) NAT firewall switch.
---

# tencentcloudenterprise_cfw_nat_firewall_switch

Provides a resource to create a cloud firewall (cfw) NAT firewall switch.

## Example Usage

```hcl
resource "tencentcloudenterprise_cfw_nat_firewall_switch" "example" {
  nat_ins_id = "cfwnat-xxxxxxxx"
  subnet_id  = "subnet-xxxxxxxx"
  enable     = 1
}
```

## Argument Reference

The following arguments are supported:

* `enable` - (Required, Int) Switch, 0: off, 1: on.
* `nat_ins_id` - (Required, String, ForceNew) Firewall instance id.
* `subnet_id` - (Required, String, ForceNew) subnet id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cfw_nat_firewall_switch can be imported using the id, e.g.

```
Cloud firewall NAT firewall switch can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_nat_firewall_switch.example nat_ins_id#subnet_id
```
```


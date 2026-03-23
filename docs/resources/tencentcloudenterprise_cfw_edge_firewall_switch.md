---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_edge_firewall_switch"
sidebar_current: "docs-tencentcloudenterprise-resource-cfw_edge_firewall_switch"
description: |-
  Provides a resource to create a cloud firewall (cfw) edge firewall switch.
---

# tencentcloudenterprise_cfw_edge_firewall_switch

Provides a resource to create a cloud firewall (cfw) edge firewall switch.

## Example Usage

```hcl
resource "tencentcloudenterprise_cfw_edge_firewall_switch" "example" {
  public_ip   = "1.1.1.1"
  subnet_id   = "subnet-xxxxxxxx"
  switch_mode = 1
  enable      = 1
}
```

## Argument Reference

The following arguments are supported:

* `enable` - (Required, Int) Switch, 0: off, 1: on.
* `public_ip` - (Required, String, ForceNew) Public Ip.
* `switch_mode` - (Required, Int) 0: bypass; 1: serial.
* `subnet_id` - (Optional, String) The first EIP switch in the vpc is turned on, and you need to specify a subnet to create a private connection. If `switch_mode` is 1 and `enable` is 1, this field is required.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cfw_edge_firewall_switch can be imported using the id, e.g.

```
Cloud firewall edge firewall switch can be imported using the public_ip, e.g.

```
$ terraform import tencentcloudenterprise_cfw_edge_firewall_switch.example 1.1.1.1
```
```


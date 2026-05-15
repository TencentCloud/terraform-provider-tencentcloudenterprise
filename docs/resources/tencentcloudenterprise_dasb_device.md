---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_device"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_device"
description: |-
  Provide a resource to create a DASB device
---

# tencentcloudenterprise_dasb_device

Provide a resource to create a DASB device

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_device" "example" {
  os_name = "Linux"
  ip      = "192.168.1.1"
  port    = 22
}
```

## Argument Reference

The following arguments are supported:

* `ip` - (Required, String) IP address.
* `os_name` - (Required, String) Operating system name, only Linux, Windows or MySQL.
* `port` - (Required, Int) Management port.
* `department_id` - (Optional, String) The department ID to which the device belongs.
* `ip_port_set` - (Optional, Set: [`String`]) Asset multi-node: fields ip and port.
* `name` - (Optional, String) Hostname, can be empty.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_dasb_device can be imported using the id, e.g.

```
DASB device can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_device.example 12345
```
```


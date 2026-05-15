---
subcategory: "BH"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bh_device"
sidebar_current: "docs-tencentcloudenterprise-resource-bh_device"
description: |-
  Provide a resource to create a BH device
---

# tencentcloudenterprise_bh_device

Provide a resource to create a BH device

## Example Usage

```hcl
resource "tencentcloudenterprise_bh_device" "example" {
  device_set {
    os_name = "Linux"
    ip      = "10.0.0.1"
    port    = 22
    name    = "example-device"
  }
}
```

## Argument Reference

The following arguments are supported:

* `device_set` - (Required, List) Asset parameter list.

The `device_set` object supports the following:

* `ip` - (Required, String, ForceNew) IP address.
* `os_name` - (Required, String, ForceNew) The operating system name can only be one of the following: Host (Linux, Windows), Database (MySQL, SQL Server, MariaDB, PostgreSQL, MongoDBReplicaSet, MongoDBSharded, Redis), or Container (TKE, EKS).
* `port` - (Required, Int) Management port.
* `department_id` - (Optional, String) Department ID to which the asset belongs.
* `enable_ssl` - (Optional, Int, ForceNew) Whether to enable SSL, 1: enable, 0: disable, only supports Redis assets.
* `ip_port_set` - (Optional, Set, ForceNew) Asset multi-node: IP and port fields.
* `name` - (Optional, String, ForceNew) Host name, can be empty.
* `ssl_cert_name` - (Optional, String, ForceNew) SSL certificate name, required when EnableSSL is enabled.
* `ssl_cert` - (Optional, String, ForceNew) SSL certificate, required when EnableSSL is enabled.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `device_id` - ID of the device.

## Import

tencentcloudenterprise_bh_device can be imported using the id, e.g.

```
BH device can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_device.example 12345
```
```


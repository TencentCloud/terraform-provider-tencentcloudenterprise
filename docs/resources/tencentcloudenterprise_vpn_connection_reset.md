---
subcategory: "VPN Connections(VPN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpn_connection_reset"
sidebar_current: "docs-tencentcloudenterprise-resources-vpn_connection_reset"
description: |-
  Provides a resource to create a vpc vpn_connection_reset
---

# tencentcloudenterprise_vpn_connection_reset

Provides a resource to create a vpc vpn_connection_reset

## Example Usage

```hcl
resource "tencentcloudenterprise_vpn_connection_reset" "vpn_connection_reset" {
  vpn_gateway_id    = "vpngw-gt8bianl"
  vpn_connection_id = "vpnx-kme2tx8m"
}
```

## Argument Reference

The following arguments are supported:

* `vpn_connection_id` - (Required, String, ForceNew) VPN CONNECTION INSTANCE ID.
* `vpn_gateway_id` - (Required, String, ForceNew) VPN GATEWAY INSTANCE ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




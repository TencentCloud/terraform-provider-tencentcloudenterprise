---
subcategory: "Direct Connect(DC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dc_approve_tunnel"
sidebar_current: "docs-tencentcloudenterprise-resource-dc_approve_tunnel"
description: |-
  Provides a resource to approve a dc tunnel
---

# tencentcloudenterprise_dc_approve_tunnel

Provides a resource to approve a dc tunnel

## Example Usage

```hcl
resource "tencentcloudenterprise_dc_approve_tunnel" "example {
	direct_connect_tunnel_name = " test-tunnel "
	direct_connect_tunnel_name = " dcx-ua8ej92a "
	approved				   = true
	comments				   = " ok "
}
```

## Argument Reference

The following arguments are supported:

* `approved` - (Required, Bool, ForceNew) Approval result for the resource: `true` (approved), `false` (disapproved).
* `comments` - (Required, String, ForceNew) Comment of the approval.
* `direct_connect_tunnel_id` - (Required, String, ForceNew) ID of direct connect tunnel.
* `direct_connect_tunnel_name` - (Required, String, ForceNew) Connection name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_dc_approve_tunnel can be imported using the id, e.g.

```
approve_tunnel can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_dc_approve_tunnel.example
```
```


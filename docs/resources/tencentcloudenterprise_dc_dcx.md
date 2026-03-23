---
subcategory: "Direct Connect(DC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dc_dcx"
sidebar_current: "docs-tencentcloudenterprise-resource-dc_dcx"
description: |-
  Provides a resource to creating dedicated tunnels instances.
---

# tencentcloudenterprise_dc_dcx

Provides a resource to creating dedicated tunnels instances.

~> **NOTE:** 1. ID of the DC is queried, can only apply for this resource offline.

## Example Usage

```hcl
variable "dc_id" {
  default = "dc-kax48sg7"
}

variable "dcg_id" {
  default = "dcg-dmbhf7jf"
}

variable "vpc_id" {
  default = "vpc-4h9v4mo3"
}

resource "tencentcloudenterprise_dc_dcx" "bgp_main" {
  bandwidth    = 900
  dc_id        = var.dc_id
  dcg_id       = var.dcg_id
  name         = "bgp_main"
  network_type = "VPC"
  route_type   = "BGP"
  vlan         = 306
  vpc_id       = var.vpc_id
}

resource "tencentcloudenterprise_dc_dcx" "static_main" {
  bandwidth                      = 900
  dc_id                          = var.dc_id
  dcg_id                         = var.dcg_id
  name                           = "static_main"
  network_type                   = "VPC"
  route_type                     = "STATIC"
  vlan                           = 301
  vpc_id                         = var.vpc_id
  tencentcloudenterprise_address = "100.93.46.1/30"
  customer_address               = "100.93.46.2/30"
  idc_routes = [
    "10.0.0.0/16",
    "10.2.0.0/16"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `bandwidth` - (Required, Int) Bandwidth of the DC in Mbps.
* `connect_subnet_mask` - (Required, Int, ForceNew) Mask of the interconnect address.
* `customer_address` - (Required, String, ForceNew) Interconnect IP of the DC within client.
* `direct_connect_gateway_id` - (Required, String, ForceNew) ID of the DC Gateway. Currently only new in the console.
* `direct_connect_id` - (Required, String, ForceNew) ID of the DC to be queried, application deployment offline.
* `direct_connect_tunnel_name` - (Required, String) Name of the dedicated tunnel.
* `enable_bfd` - (Required, Bool) Whether enables BFD.
* `load_mode` - (Required, String) Tunnel Load Balancing Mode: None (Non-redundant Mode), LoadBalance (Load Balancing), MasterSlave (Active-Standby).
* `network_region` - (Required, String, ForceNew) Region of the vpc.
* `route_type` - (Required, String, ForceNew) Type of the route, and available values include `BGP` and `STATIC`.
* `tencentcloudenterprise_address` - (Required, String, ForceNew) Interconnect IP of the DC within cloud.
* `vlan` - (Required, Int, ForceNew) Vlan: Range: 11 ~ 4000; it must be consistent with the VLAN ID on the user side.
* `vpc_id` - (Required, Int, ForceNew) ID of the VPC or BMVPC.
* `vpc_name` - (Required, String, ForceNew) ID of the VPC or BMVPC.
* `bfd_interval` - (Optional, Int) BFD Protocol Interval Configuration.
* `bgp_peer` - (Optional, List) BGP peer information configured by the user, including Asn and AuthKey.
* `direct_connect_owner_account` - (Optional, String) Direct connect owner: The default value is the current customer. When sharing the dedicated line (by the physical dedicated line owner), the developer account ID of the shared dedicated line must be filled in here.
* `enable_multicast` - (Optional, Bool) Whether to enable multicast. Can only be modified after creation via updates.
* `idc_routes` - (Optional, Set: [`String`]) Static route, the network segment address of the user's IDC.
* `ip_type` - (Optional, String, ForceNew) Protocol type of the tunnel IP.
* `multicast_groups` - (Optional, String) Multicast group addresses supported by the tunnel. Can only be modified after creation via updates. Only valid when enable_multicast is true.
* `related_direct_connect_tunnel_id` - (Optional, String) ID of the related redundant DC.

The `bgp_peer` object supports the following:

* `bgp_asn` - (Optional, Int) BGP ASN of the user. A required field within BGP.
* `bgp_auth_key` - (Optional, String) BGP key of the user.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `bfd_state` - The BFD state of the dedicated tunnel. Possible values: DISABLED, ENABLE, UP, DOWN.
* `created_time` - Creation time of the direct connect tunnel.
* `direct_connect_gateway_name` - Direct connect gateway name.
* `nat_type` - Whether it is a NAT tunnel.
* `net_detect_id` - Network detection ID.
* `state` - The state of the dedicated tunnel. Possible values: AVAILABLE, APPLYING, ALLOCATING, ALLOCATED, ALTERING, DELETING, DELETED, PENDING, REJECTED.
* `vpc_region` - VPC region where the tunnel is connected.


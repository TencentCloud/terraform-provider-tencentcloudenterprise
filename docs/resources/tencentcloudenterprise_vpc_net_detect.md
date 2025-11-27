---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_net_detect"
sidebar_current: "docs-tencentcloudenterprise-resources-vpc_net_detect"
description: |-
  Provides a resource to create a vpc net_detect
---

# tencentcloudenterprise_vpc_net_detect

Provides a resource to create a vpc net_detect

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_net_detect" "net_detect" {
  net_detect_name      = "terrform-test"
  vpc_id               = "vpc-4owdpnwr"
  subnet_id            = "subnet-c1l35990"
  next_hop_destination = "203.0.113.57"
  next_hop_type        = "NORMAL_CVM"
  detect_destination_ip = [
    "203.0.113.1",
    "203.0.113.2",
  ]
}
```

## Argument Reference

The following arguments are supported:

* `detect_destination_ip` - (Required, Set: [`String`]) An array of probe destination IPv4 addresses. Up to two.
* `net_detect_name` - (Required, String) Network probe name, the maximum length cannot exceed 60 bytes.
* `next_hop_destination` - (Required, String) The destination gateway of the next hop, the value is related to the next hop type.If the next hop type is VPN, and the value is the VPN gateway ID, such as: vpngw-12345678;If the next hop type is DIRECTCONNECT, and the value is the private line gateway ID, such as: dcg-12345678; If the next hop type is PEERCONNECTION, which takes the value of the peer connection ID, such as: pcx-12345678;If the next hop type is NAT, and the value is Nat gateway, such as: nat-12345678; If the next hop type is NORMAL_CVM, which takes the IPv4 address of the cloud server, such as: 203.0.113.12.
* `next_hop_type` - (Required, String) NextHopType: indicates the next hop type.Currently, the following types are supported: VPN: indicates the VPN gateway. DIRECTCONNECT: private line gateway; PEERCONNECTION: peer connection; NAT: indicates the NAT gateway. NORMAL_CVM: common cloud server.
* `subnet_id` - (Required, String, ForceNew) Subnet instance ID. Such as:subnet-12345678.
* `vpc_id` - (Required, String, ForceNew) `VPC` instance `ID`. Such as:`vpc-12345678`.
* `net_detect_description` - (Optional, String) Network probe description.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

vpc net_detect can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_vpc_net_detect.net_detect net_detect_id
```


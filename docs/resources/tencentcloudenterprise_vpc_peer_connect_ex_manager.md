---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_peer_connect_ex_manager"
sidebar_current: "docs-tencentcloudenterprise-resources-vpc_peer_connect_ex_manager"
description: |-
  Provides a resource to create and manage a VPC peering connection with extended features.
---

# tencentcloudenterprise_vpc_peer_connect_ex_manager

Provides a resource to create and manage a VPC peering connection with extended features.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_peer_connect_ex_manager" "example" {
  vpc_id                  = "vpc-45dvaaw9"
  peering_connection_name = "test-peer-connection-ex"
  peer_vpc_id             = "vpc-44fnavba"
  peer_uin                = "110000053176"
  peer_region             = "ap-shenzhen-region"
  bandwidth               = 500
}
```

## Argument Reference

The following arguments are supported:

* `bandwidth` - (Required, Int) Peering connection bandwidth value.
* `peer_region` - (Required, String) Peer region.
* `peer_uin` - (Required, String) Peer user UIN.
* `peer_vpc_id` - (Required, String) The unique ID of the peer VPC.
* `peering_connection_name` - (Required, String) Peer connection name.
* `vpc_id` - (Required, String) The unique ID of the local VPC.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

VPC peering connection ex can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_vpc_peer_connect_ex_manager.example pcx-1asg3t63
```


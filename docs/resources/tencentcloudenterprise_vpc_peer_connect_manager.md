---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_peer_connect_manager"
sidebar_current: "docs-tencentcloudenterprise-resources-vpc_peer_connect_manager"
description: |-
  Provides a resource to create and manage a VPC peering connection.
---

# tencentcloudenterprise_vpc_peer_connect_manager

Provides a resource to create and manage a VPC peering connection.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_peer_connect_manager" "foo" {
  vpc_id                  = "vpc-4xxr2cy7"
  peering_connection_name = "test_peer_connection"
  peer_vpc_id             = "vpc-5ggr3dx8"
  peer_uin                = "100001234567"
  peer_region             = "ap-beijing"
}
```

## Argument Reference

The following arguments are supported:

* `peer_region` - (Required, String, ForceNew) Peer region.
* `peer_uin` - (Required, String, ForceNew) Peer user UIN.
* `peer_vpc_id` - (Required, String, ForceNew) The unique ID of the peer VPC.
* `peering_connection_name` - (Required, String) Peer connection name.
* `vpc_id` - (Required, String, ForceNew) The unique ID of the local VPC.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

VPC peering connection can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_vpc_peer_connect_manager.foo pcx-1asg3t63
```


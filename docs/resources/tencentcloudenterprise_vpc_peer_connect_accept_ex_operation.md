---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_peer_connect_accept_ex_operation"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_peer_connect_accept_ex_operation"
description: |-
  Provides a resource to accept a cross-region cross-account VPC peering connection.
---

# tencentcloudenterprise_vpc_peer_connect_accept_ex_operation

Provides a resource to accept a cross-region cross-account VPC peering connection.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_peer_connect_accept_ex_operation" "example" {
  peering_connection_id = "pcx-1asg3t63"
}
```

## Argument Reference

The following arguments are supported:

* `peering_connection_id` - (Required, String, ForceNew) The unique ID of the peering connection.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



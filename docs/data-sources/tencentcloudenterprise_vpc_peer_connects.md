---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_peer_connects"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpc_peer_connects"
description: |-
  Use this data source to query detailed information of VPC peering connections.
---

# tencentcloudenterprise_vpc_peer_connects

Use this data source to query detailed information of VPC peering connections.

## Example Usage

```hcl
data "tencentcloudenterprise_vpc_peer_connects" "by_id" {
  peering_connection_id = "pcx-xxxxxxxx"
}

data "tencentcloudenterprise_vpc_peer_connects" "by_vpc" {
  vpc_id = "vpc-xxxxxxxx"
}

data "tencentcloudenterprise_vpc_peer_connects" "by_name" {
  peering_connection_name = "my-peer"
}

data "tencentcloudenterprise_vpc_peer_connects" "by_state" {
  state = "ACTIVE"
}
```

## Argument Reference

The following arguments are supported:

* `peering_connection_id` - (Optional, String) Query by exact peering connection ID. Cannot be used with other filters.
* `peering_connection_name` - (Optional, String) Filter by peering connection name (fuzzy match via Filters).
* `result_output_file` - (Optional, String) Used to save results.
* `state` - (Optional, String) Filter by state: PENDING / ACTIVE / EXPIRED / REJECTED.
* `vpc_id` - (Optional, String) Filter by local VPC ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - A list of peering connections. Each element contains the following attributes:
  * `app_id` - AppId of the local end.
  * `bandwidth` - Bandwidth of the peering connection (Mbps).
  * `charge_type` - Charge type of the peering connection.
  * `create_time` - Creation time of the peering connection.
  * `dst_region` - Region of the peer end.
  * `peer_app_id` - AppId of the peer end.
  * `peer_uin` - UIN of the peer end.
  * `peer_vpc_cidr_block` - IPv4 CIDR of the peer VPC.
  * `peer_vpc_id` - VPC ID of the peer end.
  * `peer_vpc_name` - VPC name of the peer end.
  * `peering_connection_id` - ID of the peering connection.
  * `peering_connection_name` - Name of the peering connection.
  * `src_region` - Region of the local end.
  * `state` - State of the peering connection: PENDING / ACTIVE / REJECTED / DELETED / FAILED.
  * `tags` - Tags of the peering connection.
  * `type` - 0: basic network interconnection; 1: VPC-to-VPC; 2: VPC to BM network.
  * `uin` - UIN of the local end.
  * `vpc_cidr_block` - IPv4 CIDR of the local VPC.
  * `vpc_id` - VPC ID of the local end.
  * `vpc_name` - VPC name of the local end.


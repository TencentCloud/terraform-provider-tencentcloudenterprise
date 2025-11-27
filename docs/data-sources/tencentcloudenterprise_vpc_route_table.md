---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_route_table"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpc_route_table"
description: |-
  Provides details about a specific Route Table.
---

# tencentcloudenterprise_vpc_route_table

Provides details about a specific Route Table.

This resource can prove useful when a module accepts a Subnet id as an input variable and needs to, for example, add a route in the Route Table.

~> **NOTE:** It has been deprecated and replaced by tencentcloudenterprise_vpc_route_tables.

## Example Usage

```hcl
variable "route_table_id" {}

data "tencentcloudenterprise_route_table" "selected" {
  route_table_id = var.route_table_id
}

resource "tencentcloudenterprise_route_entry" "rtb_entry_instance" {
  vpc_id         = "{data.tencentcloudenterprise_route_table.selected.vpc_id}"
  route_table_id = var.route_table_id
  cidr_block     = "10.4.8.0/24"
  next_type      = "instance"
  next_hub       = "10.16.1.7"
}
```

## Argument Reference

The following arguments are supported:

* `route_table_id` - (Required, String) The Route Table ID.
* `name` - (Optional, String) The Route Table name.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `create_time` - Creation time of routing table.
* `routes` - The information list of the VPC route table.
  * `cidr_block` - The RouteEntry's target network segment.
  * `description` - The RouteEntry's description.
  * `next_hub` - The RouteEntry's next hub.
  * `next_type` - The `next_hub` type.
* `subnet_num` - Number of associated subnets.
* `vpc_id` - The VPC ID.



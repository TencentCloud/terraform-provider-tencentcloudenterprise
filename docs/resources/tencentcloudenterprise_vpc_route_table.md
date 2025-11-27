---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_route_table"
sidebar_current: "docs-tencentcloudenterprise-resources-vpc_route_table"
description: |-
  Provides a resource to create a VPC routing table.
---

# tencentcloudenterprise_vpc_route_table

Provides a resource to create a VPC routing table.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc" "foo" {
  name       = "ci-temp-test"
  cidr_block = "10.0.0.0/16"
}

resource "tencentcloudenterprise_vpc_route_table" "foo" {
  vpc_id = tencentcloudenterprise_vpc.foo.id
  name   = "ci-temp-test-rt"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) The name of routing table.
* `vpc_id` - (Required, String, ForceNew) ID of VPC to which the route table should be associated.
* `tags` - (Optional, Map) The tags of routing table.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of the routing table.
* `is_default` - Indicates whether it is the default routing table.
* `route_entry_ids` - ID list of the routing entries.
* `subnet_ids` - ID list of the subnets associated with this route table.


## Import

Vpc routetable instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_route_table.test route_table_id
```


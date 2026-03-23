---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_nat_gateway_flow_monitor"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_nat_gateway_flow_monitor"
description: |-
  Provides a resource to manage the flow monitor of a NAT gateway.
---

# tencentcloudenterprise_vpc_nat_gateway_flow_monitor

Provides a resource to manage the flow monitor of a NAT gateway.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_nat_gateway_flow_monitor" "example" {
  gateway_id = "nat-xxxxxxxx"
  enable     = true
}
```

## Argument Reference

The following arguments are supported:

* `enable` - (Required, Bool) Whether to enable flow monitor.
* `gateway_id` - (Required, String, ForceNew) ID of Gateway.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `bandwidth` - Bandwidth of flow monitor.

## Import

tencentcloudenterprise_vpc_nat_gateway_flow_monitor can be imported using the id, e.g.

```
NAT gateway flow monitor can be imported using the gateway id, e.g.

```
$ terraform import tencentcloudenterprise_vpc_nat_gateway_flow_monitor.example nat-xxxxxxxx
```
```


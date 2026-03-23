---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_end_point"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_end_point"
description: |-
  Provides a resource to create a vpc end_point
---

# tencentcloudenterprise_vpc_end_point

Provides a resource to create a vpc end_point

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_end_point" "example" {
  vpc_id               = "vpc-ffwo6rid"
  subnet_id            = "subnet-o7v0wz10"
  end_point_name       = "123tf"
  end_point_service_id = "vpcsvc-o9u88lu5"
  end_point_vip        = "192.168.32.20"
  ip_address_type      = "IPv4"
  security_group_id    = "sg-iz7ipqme"
}
```

## Argument Reference

The following arguments are supported:

* `end_point_name` - (Required, String) Name of endpoint.
* `end_point_service_id` - (Required, String) ID of endpoint service.
* `subnet_id` - (Required, String) ID of subnet instance.
* `vpc_id` - (Required, String) ID of vpc instance.
* `end_point_vip` - (Optional, String) VIP of endpoint ip.
* `ip_address_type` - (Optional, String) IP address type: IPv4/IPv6.
* `security_group_id` - (Optional, Set: [`String`]) List of security group IDs.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create Time.
* `end_point_owner` - APPID.
* `state` - State of end point.

## Import

tencentcloudenterprise_vpc_end_point can be imported using the id, e.g.

```
vpc end_point can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_vpc_end_point.end_point end_point_id
```
```


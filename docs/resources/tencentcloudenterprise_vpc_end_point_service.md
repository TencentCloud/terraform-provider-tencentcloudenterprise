---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_end_point_service"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_end_point_service"
description: |-
  Provides a resource to create a vpc end_point_service
---

# tencentcloudenterprise_vpc_end_point_service

Provides a resource to create a vpc end_point_service

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_end_point_service" "end_point_service" {
  vpc_id                 = "vpc-391sv4w3"
  end_point_service_name = "terraform-endpoint-service"
  auto_accept_flag       = false
  service_instance_id    = "lb-o5f6x7ke"
}
```

## Argument Reference

The following arguments are supported:

* `auto_accept_flag` - (Required, Bool) Whether to automatically accept.
* `end_point_service_name` - (Required, String) Name of end point service.
* `service_instance_id` - (Required, String) Id of service instance, like lb-xxx.
* `vpc_id` - (Required, String) ID of vpc instance.
* `ip_address_type` - (Optional, String) Type of the IP address: IPv4/IPv6.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `end_point_service` - End point service details.
  * `auto_accept_flag` - Whether to automatically accept.
  * `create_time` - Create time.
  * `end_point_count` - Number of associated end points.
  * `end_point_service_id` - End point service ID.
  * `end_point_service` - End point object array.
    * `create_time` - Create time.
    * `end_point_id` - End point ID.
    * `end_point_name` - End point name.
    * `end_point_owner` - End point owner (APPID).
    * `end_point_service_id` - End point service ID.
    * `end_point_vip` - End point VIP.
    * `group_set` - Security group instance ID list bound to end point.
    * `service_name` - End point service name.
    * `service_vip` - End point service VIP.
    * `service_vpc_id` - End point service VPC ID.
    * `state` - End point state: ACTIVE, PENDING, ACCEPTING, REJECTED, FAILED.
    * `subnet_id` - Subnet ID.
    * `vpc_id` - VPC ID.
  * `ip_address_type` - Type of IP address: IPv4/IPv6.
  * `service_instance_id` - Backend service ID, like lb-xxx.
  * `service_name` - End point service name.
  * `service_owner` - Service owner (APPID).
  * `service_vip` - Backend service VIP.
  * `vpc_id` - VPC ID.

## Import

tencentcloudenterprise_vpc_end_point_service can be imported using the id, e.g.

```
vpc end_point_service can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_vpc_end_point_service.end_point_service end_point_service_id
```
```


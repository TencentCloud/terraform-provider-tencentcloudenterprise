---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_end_point_service_white_list"
sidebar_current: "docs-tencentcloudenterprise-resource-vpc_end_point_service_white_list"
description: |-
  Provides a resource to create a vpc end_point_service_white_list
---

# tencentcloudenterprise_vpc_end_point_service_white_list

Provides a resource to create a vpc end_point_service_white_list

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_end_point_service_white_list" "end_point_service_white_list" {
  user_uin             = "100020512675"
  end_point_service_id = "vpcsvc-69y13tdb"
  description          = "terraform for test"
}
```

## Argument Reference

The following arguments are supported:

* `end_point_service_id` - (Required, String) ID of endpoint service.
* `user_uin` - (Required, String) UIN.
* `description` - (Optional, String) Description of white list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create Time.
* `owner` - APPID.

## Import

tencentcloudenterprise_vpc_end_point_service_white_list can be imported using the id, e.g.

```
vpc end_point_service_white_list can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_vpc_end_point_service_white_list.end_point_service_white_list end_point_service_white_list_id
```
```


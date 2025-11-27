---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_limits"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpc_limits"
description: |-
  Use this data source to query detailed information of vpc limits
---

# tencentcloudenterprise_vpc_limits

Use this data source to query detailed information of vpc limits

## Example Usage

```hcl
data "tencentcloudenterprise_vpc_limits" "limits" {
  limit_types = ["appid-max-vpcs", "vpc-max-subnets"]
}
```

## Argument Reference

The following arguments are supported:

* `limit_types` - (Required, Set: [`String`]) Quota name. A maximum of 100 quota types can be queried each time.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `vpc_limit_set` - Vpc limit.
  * `limit_type` - Type of vpc limit.
  * `limit_value` - Value of vpc limit.



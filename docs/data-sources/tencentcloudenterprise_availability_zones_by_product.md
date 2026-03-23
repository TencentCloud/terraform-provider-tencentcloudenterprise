---
subcategory: "Provider Data Sources"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_availability_zones_by_product"
sidebar_current: "docs-tencentcloudenterprise-datasource-availability_zones_by_product"
description: |-
  Use this data source to query availability zones by product.
---

# tencentcloudenterprise_availability_zones_by_product

Use this data source to query availability zones by product.

## Example Usage

```hcl
data "tencentcloudenterprise_availability_zones_by_product" "example" {
  product = "cvm"
}
```

## Argument Reference

The following arguments are supported:

* `product` - (Required, String) A string variable indicates that the query will use product information.
* `include_unavailable` - (Optional, Bool) A bool variable indicates that the query will include `UNAVAILABLE` zones.
* `name` - (Optional, String) When specified, only the zone with the exactly name match will be returned.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `zones` - A list of zones will be exported and its every element contains the following attributes:
  * `description` - The description of the zone, like `Guangzhou Zone 3`.
  * `id` - An internal id for the zone, like `200003`, usually not so useful.
  * `name` - The name of the zone, like `ap-guangzhou-3`.
  * `state` - The state of the zone, indicate availability using `AVAILABLE` and `UNAVAILABLE` values.


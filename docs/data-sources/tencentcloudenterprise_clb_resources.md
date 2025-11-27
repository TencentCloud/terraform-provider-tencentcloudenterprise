---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_resources"
sidebar_current: "docs-tencentcloudenterprise-datasource-clb_resources"
description: |-
  Use this data source to query detailed information of clb resources
---

# tencentcloudenterprise_clb_resources

Use this data source to query detailed information of clb resources

## Example Usage

```hcl
data "tencentcloudenterprise_clb_resources" "resources" {
  filters {
    name   = "isp"
    values = ["BGP"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `filters` - (Optional, List) Filter to query the list of AZ resources.
* `result_output_file` - (Optional, String) Used to save results.

The `filters` object supports the following:

* `name` - (Required, String) Filter name.
* `values` - (Required, Set) Filter value array.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `resource_set` - List of resources supported by the AZ.
  * `domain` - Domain of clb.
  * `load_balancer_id` - Id of clb.
  * `load_balancer_name` - Name of clb.
  * `load_balancer_type` - Type of clb.
  * `load_balancer_vips` - Vips of clb.
  * `master_zone` - Master zone of clb.
    * `local_zone` - Is the local zone.
    * `zone_id` - Zone id of clb.
    * `zone_name` - Zone name of clb.
    * `zone_region` - Zone region of clb.
    * `zone` - Zone of clb.



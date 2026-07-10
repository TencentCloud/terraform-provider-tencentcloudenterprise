---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn_route_tables"
sidebar_current: "docs-tencentcloudenterprise-datasource-ccn_route_tables"
description: |-
  Use this data source to query CCN (Cloud Connect Network) route tables.
---

# tencentcloudenterprise_ccn_route_tables

Use this data source to query CCN (Cloud Connect Network) route tables.

## Example Usage

```hcl
data "tencentcloudenterprise_ccn_route_tables" "by_ccn" {
  ccn_id = "ccn-xxxxxxxx"
}

data "tencentcloudenterprise_ccn_route_tables" "by_name" {
  ccn_id           = "ccn-xxxxxxxx"
  route_table_name = "my-rtb"
}

data "tencentcloudenterprise_ccn_route_tables" "by_id" {
  route_table_id = "ccnrtb-xxxxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `ccn_id` - (Optional, String) Filter by CCN instance ID.
* `result_output_file` - (Optional, String) Used to save results.
* `route_table_description` - (Optional, String) Filter by route table description.
* `route_table_id` - (Optional, String) Filter by exact route table ID.
* `route_table_name` - (Optional, String) Filter by route table name (fuzzy match).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - A list of CCN route tables. Each element contains the following attributes:
  * `ccn_id` - ID of the CCN instance.
  * `create_time` - Creation time of the route table.
  * `is_default_table` - True: default route table; False: custom route table.
  * `route_table_description` - Description of the CCN route table.
  * `route_table_id` - ID of the CCN route table.
  * `route_table_name` - Name of the CCN route table.


---
subcategory: "Direct Connect(DC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dc_gateway_ccn_routes"
sidebar_current: "docs-tencentcloudenterprise-datasource-dc_gateway_ccn_routes"
description: |-
  Use this data source to query detailed information of direct connect gateway route entries.
---

# tencentcloudenterprise_dc_gateway_ccn_routes

Use this data source to query detailed information of direct connect gateway route entries.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn" "main" {
  name        = "ci-temp-test-ccn"
  description = "ci-temp-test-ccn-des"
  qos         = "AG"
}

resource "tencentcloudenterprise_vpc_dc_gateway" "ccn_main" {
  name                = "ci-cdg-ccn-test"
  network_instance_id = tencentcloudenterprise_ccn.main.id
  network_type        = "CCN"
  gateway_type        = "NORMAL"
}

resource "tencentcloudenterprise_dc_gateway_ccn_route" "route1" {
  dcg_id     = tencentcloudenterprise_vpc_dc_gateway.ccn_main.id
  cidr_block = "10.1.1.0/32"
}

resource "tencentcloudenterprise_dc_gateway_ccn_route" "route2" {
  dcg_id     = tencentcloudenterprise_vpc_dc_gateway.ccn_main.id
  cidr_block = "192.1.1.0/32"
}

#You need to sleep for a few seconds because there is a cache on the server
data "tencentcloudenterprise_dc_gateway_ccn_routes" "test" {
  dcg_id = tencentcloudenterprise_vpc_dc_gateway.ccn_main.id
}
```

## Argument Reference

The following arguments are supported:

* `dcg_id` - (Required, String) ID of the DCG to be queried.
* `ccn_route_type` - (Optional, String) Cloud networking routing learning type, optional values: BGP - Automatic Learning; STATIC - User configured. Default is STATIC.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `instance_list` - Information list of the DCG route entries.
  * `as_path` - As path list of the BGP.
  * `cidr_block` - A network address segment of IDC.
  * `dcg_id` - ID of the DCG.
  * `route_id` - ID of the DCG route.


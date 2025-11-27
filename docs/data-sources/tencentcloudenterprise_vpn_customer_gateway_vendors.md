---
subcategory: "VPN Connections(VPN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpn_customer_gateway_vendors"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpn_customer_gateway_vendors"
description: |-
  Use this data source to query detailed information of vpc vpn_customer_gateway_vendors
---

# tencentcloudenterprise_vpn_customer_gateway_vendors

Use this data source to query detailed information of vpc vpn_customer_gateway_vendors

## Example Usage

```hcl
data "tencentcloudenterprise_vpn_customer_gateway_vendors" "vpn_customer_gateway_vendors" {}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `customer_gateway_vendor_set` - Customer Gateway Vendor Set.
  * `platform` - Platform.
  * `software_version` - SoftwareVersion.
  * `vendor_name` - VendorName.



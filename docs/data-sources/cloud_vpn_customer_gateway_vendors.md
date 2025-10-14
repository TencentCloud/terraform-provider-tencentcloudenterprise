---
subcategory: "VPN Connections(VPN)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpn_customer_gateway_vendors"
sidebar_current: "docs-cloud-datasource-vpn_customer_gateway_vendors"
description: |-
  Use this data source to query detailed information of vpc vpn_customer_gateway_vendors
---

# cloud_vpn_customer_gateway_vendors

Use this data source to query detailed information of vpc vpn_customer_gateway_vendors

## Example Usage

```hcl
data "cloud_vpn_customer_gateway_vendors" "vpn_customer_gateway_vendors" {}
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



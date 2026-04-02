---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_zones"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpcdns_zones"
description: |-
  Provide a data source to query VPCDNS zones.
---

# tencentcloudenterprise_vpcdns_zones

Provide a data source to query VPCDNS zones.

## Example Usage

```hcl
data "tencentcloudenterprise_vpcdns_zones" "foo" {
  result_output_file = "zones.json"
}

data "tencentcloudenterprise_vpcdns_zones" "by_domain" {
  domain = "example.com"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Optional, String) Domain name to filter.
* `result_output_file` - (Optional, String) The file path to output the result.
* `zone_id` - (Optional, String) Private zone ID to filter.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `zone_list` - List of private zones.
  * `account_vpc_set` - Cross-account VPC bindings.
    * `region` - Region.
    * `uin` - Account UIN.
    * `uniq_vpc_id` - VPC ID.
  * `created_on` - Creation time.
  * `dns_forward_status` - DNS forward status: ENABLED, DISABLED.
  * `domain_id` - Numeric domain ID.
  * `domain` - Domain name.
  * `record_count` - Record count.
  * `remark` - Remark.
  * `status` - Status: SUSPEND, ENABLED, FAILED.
  * `updated_on` - Last update time.
  * `vpc_set` - VPC binddings.
    * `region` - Region.
    * `uniq_vpc_id` - VPC ID.
  * `zone_id` - Private zone ID.


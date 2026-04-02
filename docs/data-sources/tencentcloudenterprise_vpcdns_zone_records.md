---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_zone_records"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpcdns_zone_records"
description: |-
  Provide a data source to query VPCDNS zone records.
---

# tencentcloudenterprise_vpcdns_zone_records

Provide a data source to query VPCDNS zone records.

## Example Usage

```hcl
data "tencentcloudenterprise_vpcdns_zone_records" "foo" {
  zone_id            = "zone-xxxxxxxx"
  result_output_file = "zone_records.json"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required, String) Private zone ID.
* `result_output_file` - (Optional, String) The file path to output the result.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `record_list` - List of zone records.
  * `created_on` - Creation time.
  * `enabled` - Enabled: 0 paused, 1 enabled.
  * `mx` - MX priority.
  * `record_id` - Record ID.
  * `record_type` - Record type: A, AAAA, CNAME, MX, TXT, PTR.
  * `record_value` - Record value.
  * `status` - Record status.
  * `sub_domain` - Subdomain.
  * `ttl` - TTL.
  * `updated_on` - Last update time.
  * `weight` - Record weight, 1-100.
  * `zone_id` - Zone ID.


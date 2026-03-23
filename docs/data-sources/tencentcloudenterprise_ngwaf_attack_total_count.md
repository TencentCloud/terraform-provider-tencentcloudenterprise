---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_attack_total_count"
sidebar_current: "docs-tencentcloudenterprise-datasource-ngwaf_attack_total_count"
description: |-
  Use this data source to query the total count of NGWAF attacks within a specified time period.
---

# tencentcloudenterprise_ngwaf_attack_total_count

Use this data source to query the total count of NGWAF attacks within a specified time period.

## Example Usage

```hcl
data "tencentcloudenterprise_ngwaf_attack_total_count" "example" {
  start_time   = "2023-01-01 00:00:00"
  end_time     = "2023-01-31 23:59:59"
  query_string = "severity:high"
}
```

## Argument Reference

The following arguments are supported:

* `end_time` - (Required, String) End time.
* `start_time` - (Required, String) Begin time.
* `query_string` - (Optional, String) Query conditions.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `total_count` - Total number of attacks.


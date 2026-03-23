---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_peak_points"
sidebar_current: "docs-tencentcloudenterprise-datasource-ngwaf_peak_points"
description: |-
  Use this data source to query NGWAF peak statistics for specific metrics within a time period.
---

# tencentcloudenterprise_ngwaf_peak_points

Use this data source to query NGWAF peak statistics for specific metrics within a time period.

## Example Usage

```hcl
data "tencentcloudenterprise_ngwaf_peak_points" "example" {
  from_time   = "2023-01-01 00:00:00"
  to_time     = "2023-01-31 23:59:59"
  metric_name = "access"
  domain      = "example.com"
  edition     = "clb-waf"
  instance_id = "waf-xxxxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `from_time` - (Required, String) Begin time.
* `metric_name` - (Required, String) Three values are available: `access`-Peak qps trend chart; `cc`-Trend chart of total number of CC attacks; `attack`-Trend chart of total number of web attacks.
* `to_time` - (Required, String) End time.
* `domain` - (Optional, String) The domain name to be queried. If all domain name data is queried, this parameter is not filled in.
* `edition` - (Optional, String) Only support sparta-waf and clb-waf. If not passed, there will be no filtering.
* `instance_id` - (Optional, String) WAF instance ID, if not passed, there will be no filtering.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `points` - point list.
  * `access` - qps.
  * `attack` - Number of web attacks.
  * `bot_access` - Bot qps.
  * `cc` - Number of cc attacks.
  * `down` - Peak downlink bandwidth, unit B.
  * `status_client_error` - Trend chart of the number of status codes returned by WAF to the client.
  * `status_ok` - Trend chart of the number of status codes returned by WAF to the client.
  * `status_redirect` - Trend chart of the number of status codes returned by WAF to the client.
  * `status_server_error` - Trend chart of the number of status codes returned by WAF to the server.
  * `time` - Second level timestamp.
  * `up` - Peak uplink bandwidth, unit B.
  * `upstream_client_error` - Trend chart of the number of status codes returned to WAF by the origin site.
  * `upstream_redirect` - Trend chart of the number of status codes returned to WAF by the origin site.
  * `upstream_server_error` - Trend chart of the number of status codes returned to WAF by the origin site.


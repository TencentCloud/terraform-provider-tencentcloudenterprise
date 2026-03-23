---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_attack_overview"
sidebar_current: "docs-tencentcloudenterprise-datasource-ngwaf_attack_overview"
description: |-
  Use this data source to query NGWAF attack overview statistics for a specified time period.
---

# tencentcloudenterprise_ngwaf_attack_overview

Use this data source to query NGWAF attack overview statistics for a specified time period.

## Example Usage

```hcl
data "tencentcloudenterprise_ngwaf_attack_overview" "example" {
  from_time   = "2023-01-01 00:00:00"
  to_time     = "2023-01-31 23:59:59"
  domain      = "example.com"
  instance_id = "waf-xxxxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `from_time` - (Required, String) Begin time.
* `to_time` - (Required, String) End time.
* `domain` - (Optional, String) Domain.
* `instance_id` - (Optional, String) Waf instanceId, otherwise not filter.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `access_circle_count` - Access circle count.
* `access_count` - Access count.
* `acl_circle_count` - Access control circle count.
* `acl_count` - Access control count.
* `api_assets_circle_count` - API assets circle count.
* `api_assets_count` - Api asset count.
* `api_risk_event_circle_count` - API risk event circle count.
* `api_risk_event_count` - Number of API risk events.
* `applet_circle_count` - Applet circle count.
* `applet_count` - Applet attack count.
* `attack_circle_count` - Attack circle count.
* `attack_count` - Attack count.
* `bot_circle_count` - Bot circle count.
* `bot_count` - Bot attack count.
* `cc_circle_count` - CC circle count.
* `cc_count` - CC attack count.
* `ip_black_circle_count` - IP blacklist circle count.
* `ip_black_count` - IP blacklist count.
* `leak_circle_count` - Leak circle count.
* `leak_count` - Information leak count.
* `tamper_circle_count` - Tamper circle count.
* `tamper_count` - Tamper proof count.


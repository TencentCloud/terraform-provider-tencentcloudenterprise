---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_attack_log_list"
sidebar_current: "docs-tencentcloudenterprise-datasource-ngwaf_attack_log_list"
description: |-
  Use this data source to query detailed attack logs from NGWAF (Next-Generation Web Application Firewall).
---

# tencentcloudenterprise_ngwaf_attack_log_list

Use this data source to query detailed attack logs from NGWAF (Next-Generation Web Application Firewall).

## Example Usage

```hcl
data "tencentcloudenterprise_ngwaf_attack_log_list" "example" {
  start_time   = "2023-01-01 00:00:00"
  end_time     = "2023-01-31 23:59:59"
  query_string = "action:block"
  query_count  = 20
  page         = 0
  sort         = "desc"
}
```

## Argument Reference

The following arguments are supported:

* `end_time` - (Required, String) End time.
* `query_string` - (Required, String) Lucene grammar.
* `start_time` - (Required, String) Begin time.
* `page` - (Optional, Int) Number of pages, starting from 0 by default.
* `query_count` - (Optional, Int) Number of queries, default to 10, maximum of 100.
* `result_output_file` - (Optional, String) Used to save results.
* `sort` - (Optional, String) Default desc, support desc, asc.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `data` - Attack log array.
  * `content` - The detail of attack log.
  * `file_name` - Useless.
  * `source` - Useless.
  * `time_stamp` - Time string.


---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_custom_white_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_custom_white_rule"
description: |-
  Provides a resource to create a NGWAF custom white rule.
---

# tencentcloudenterprise_ngwaf_custom_white_rule

Provides a resource to create a NGWAF custom white rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_custom_white_rule" "example" {
  name        = "example-white-rule"
  sort_id     = "10"
  expire_time = "0"
  domain      = "example.com"
  bypass      = "ACL,OWASP,CC"
  logical_op  = "and"

  strategies {
    field              = "IP"
    compare_func       = "ipmatch"
    content            = "192.168.1.100"
    arg                = ""
    case_not_sensitive = 0
  }

  strategies {
    field              = "URL"
    compare_func       = "contains"
    content            = "/api"
    arg                = ""
    case_not_sensitive = 1
  }
}
```

## Argument Reference

The following arguments are supported:

* `bypass` - (Required, String) Comma-separated list of WAF modules to bypass when the rule matches. Supported modules: ACL (Custom Rules), OWASP (Rule Engine), Webshell (Malicious File Detection), GeoIP (Geographic Block), BWIP (IP Black and White List), CC (Rate Limiting), BotRPC (BOT Protection), AntiLeakage (Information Leakage Prevention), API (API Security), AI (AI Engine), ip_outo_deny (IP Block), Applet (Mini Program Traffic Risk Control). Example: 'ACL,OWASP,CC'.
* `domain` - (Required, String) Domain name to which this custom white rule applies. The rule will only be effective for requests to this specific domain.
* `expire_time` - (Required, String) Expiration time for the rule in Unix timestamp format (seconds). Example: 1677254399 represents 2023-02-24 23:59:59. Use 0 for permanent rules that never expire.
* `name` - (Required, String) Name of the custom white rule. Must be unique within the domain. Used for identification and management purposes.
* `sort_id` - (Required, String) Priority level for rule execution. Valid range: 1-100. Lower numbers indicate higher priority (e.g., 1 has the highest priority). Rules are executed in ascending order of sort_id.
* `strategies` - (Required, List) List of matching strategies that define the conditions for the white rule to trigger. Each strategy specifies a field to match, comparison logic, and the content to match against. Multiple strategies can be combined using logical operators.
* `job_date_time` - (Optional, List) Rule execution time configuration. Required when job_type is set. Contains timing parameters for scheduled or periodic execution.
* `job_type` - (Optional, String) Rule execution mode. Determines how the rule timing is configured. Valid values: 'TimedJob' for scheduled execution (specific start/end times), 'CronJob' for periodic execution (recurring pattern). Required when job_date_time is specified.
* `logical_op` - (Optional, String) Logical operator for combining multiple strategy conditions. Determines how multiple matching conditions are evaluated. Valid values: 'and' (all conditions must be true), 'or' (any condition can be true). Default is 'and' if not specified.

The `cron` object supports the following:

* `days` - (Optional, Set) Days in each month for execution (1-31). Used for monthly scheduling. Can specify multiple days. Note: This field may return null, indicating that no valid values can be obtained.
* `end_time` - (Optional, String) End time for execution window in HH:mm format (24-hour clock). Example: "18:00" means execution ends at 6:00 PM. Must be later than start_time. Note: This field may return null, indicating that no valid values can be obtained.
* `start_time` - (Optional, String) Start time for execution window in HH:mm format (24-hour clock). Example: "09:00" means execution starts at 9:00 AM. Note: This field may return null, indicating that no valid values can be obtained.
* `w_days` - (Optional, Set) Days of each week for execution (0-6, where 0=Sunday, 1=Monday, ..., 6=Saturday). Used for weekly scheduling. Can specify multiple days. Note: This field may return null, indicating that no valid values can be obtained.

The `job_date_time` object supports the following:

* `cron` - (Optional, List) Time parameters for periodic execution (CronJob). Defines recurring execution patterns with days of week/month and time ranges. Note: This field may return null, indicating that no valid values can be obtained.
* `time_t_zone` - (Optional, String) Time zone for rule execution. Format: "UTC+X" or "UTC-X", where X is the offset from UTC. Example: "UTC+8" for China Standard Time, "UTC-5" for Eastern Standard Time. Note: This field may return null, indicating that no valid values can be obtained.
* `timed` - (Optional, List) Time parameters for scheduled execution (TimedJob). Defines specific start and end timestamps for one-time execution. Note: This field may return null, indicating that no valid values can be obtained.

The `strategies` object supports the following:

* `compare_func` - (Required, String) Logic symbol
                            Logical symbols are divided into the following types:
								empty (content is empty)
								null (do not exist)
								eq (equal to)
								neq (not equal to)
								contains (contain)
								ncontains (do not contain)
								strprefix (prefix matching)
								strsuffix (suffix matching)
								len_eq (length equals to)
								len_gt (length is greater than)
								len_lt (length is less than)
								ipmatch (belong to)
								ipnmatch (do not belong to)
								numgt (number greater than)
								numlt (number less than)
								geo_in (IP geo belongs to)
								geo_not_in (IP geo not belongs to)
								rematch (regex match)
								numgt (numerically greater than)
								numlt (numerically less than)
								numeq (numerically equal to)
								numneq (numerically not equal to)
								numle (numerically less than or equal to)
								numge (numerically greater than or equal to)
                            Different matching fields correspond to different logical operators. For details, see the matching field table above.
                        Note: This field may return null, indicating that no valid values can be obtained.
* `content` - (Required, String) Matching content
                            Currently, when the matching field is COOKIE (cookie), the matching content is not required. In other scenes, the matching content is required.
                        Note: This field may return null, indicating that no valid values can be obtained.
* `field` - (Required, String) Matching field
                            Different matching fields result in different matching parameters, logical operators, and matching contents. The details are as follows:
                        	<table><thead><tr><th>Matching Field</th><th>Matching Parameter</th><th>Logical Symbol</th><th>Matching Content</th></tr></thead><tbody><tr><td>IP (source IP)</td><td>Parameters are not supported.</td><td>ipmatch (match)<br>ipnmatch (mismatch)</td><td>Multiple IP addresses are separated by commas. A maximum of 20 IP addresses are allowed.</td></tr><tr><td>IPv6 (source IPv6)</td><td>Parameters are not supported.</td><td>ipmatch (match)<br>ipnmatch (mismatch)</td><td>A single IPv6 address is supported.</td></tr><tr><td>Referer (referer)</td><td>Parameters are not supported.</td><td>empty (Content is empty.)<br>null (do not exist)<br>eq (equal to)<br>neq (not equal to)<br>contains (contain)<br>ncontains (do not contain)<br>len_eq (length equals to)<br>len_gt (length is greater than)<br>len_lt (length is less than)<br>strprefix (prefix matching)<br>strsuffix (suffix matching)<br>rematch (regular expression matching)</td><td>Enter the content, with a maximum of 512 characters.</td></tr><tr><td>URL (request path)</td><td>Parameters are not supported.</td><td>eq (equal to)<br>neq (not equal to)<br>contains (contain)<br>ncontains (do not contain)<br>len_eq (length equals to)<br>len_gt (length is greater than)<br>len_lt (length is less than)<br>strprefix (prefix matching)<br>strsuffix (suffix matching)<br>rematch (regular expression matching)</td><td>Enter the content starting with /, with a maximum of 512 characters.</td></tr><tr><td>UserAgent (UserAgent)</td><td>Parameters are not supported.</td><td>Same logical symbols as the matching field <font color="Red">Referer</font></td><td>Enter the content with a maximum of 512 characters.</td></tr><tr><td>HTTP_METHOD (HTTP request method)</td><td>Parameters are not supported.</td><td>eq (equal to)<br>neq (not equal to)</td><td>Enter the method name. The uppercase is recommended.</td></tr><tr><td>QUERY_STRING (request string)</td><td>Parameters are not supported.</td><td>Same logical symbol as the matchin... [truncated]
          	  				Note: This field may return null, indicating that no valid values can be obtained.
* `arg` - (Optional, String) Matching parameter
                            There are two types of configuration parameters: unsupported parameters and supported parameters.
                            The matching parameter can be entered only when the matching field is one of the following four. Otherwise, the parameter is not supported.
                                GET (GET parameter value)		
                                POST (POST parameter value)		
                                ARGS_COOKIE (Cookie parameter value)		
                                ARGS_HEADER (Header parameter value)
                        Note: This field may return null, indicating that no valid values can be obtained.
* `case_not_sensitive` - (Optional, Int) 0: case-sensitive, 1: case-insensitive. Note: This field may return null, indicating that no valid values can be obtained.

The `timed` object supports the following:

* `end_date_time` - (Optional, Int) End timestamp for scheduled execution, in Unix timestamp format (seconds). Must be greater than start_date_time. Example: 1677340799 represents 2023-02-25 23:59:59. Note: This field may return null, indicating that no valid values can be obtained.
* `start_date_time` - (Optional, Int) Start timestamp for scheduled execution, in Unix timestamp format (seconds). Example: 1677254399 represents 2023-02-24 23:59:59. Note: This field may return null, indicating that no valid values can be obtained.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `rule_id` - The unique identifier of the custom white rule, automatically generated upon creation. Used to identify and manage the rule.

## Import

tencentcloudenterprise_ngwaf_custom_white_rule can be imported using the id, e.g.

```
NGWAF custom white rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_custom_white_rule.example example.com#rule-xxxxxxxx
```
```


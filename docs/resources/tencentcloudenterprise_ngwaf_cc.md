---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_cc"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_cc"
description: |-
  Provides a resource to create a NGWAF CC protection rule.
---

# tencentcloudenterprise_ngwaf_cc

Provides a resource to create a NGWAF CC protection rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_cc" "basic_protection" {
  domain      = "example.com"
  name        = "basic-cc-protection"
  status      = 1
  advance     = "0"
  limit       = "100"
  interval    = "60"
  url         = "/*"
  match_func  = 1
  action_type = "22"
  priority    = 10
  valid_time  = 300
  edition     = "clb-waf"
}
```

## Argument Reference

The following arguments are supported:

* `action_type` - (Required, String) Rule Action, 20 means observation, 21 means human-machine identification, 22 means interception, 23 means precise interception, 26 means precise human-machine identification.
* `advance` - (Required, String) Advanced mode (whether to use session detection). 0(disabled) 1(enabled).
* `domain` - (Required, String) Domain.
* `edition` - (Required, String) WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf.
* `interval` - (Required, String) CC detection cycle.
* `limit` - (Required, String) CC detection threshold.
* `match_func` - (Required, Int) Match method, 0(equal), 1(prefix), 2(contains), 3(not equal), 6(suffix), 7(not contains).
* `name` - (Required, String) Rule Name.
* `priority` - (Required, Int) Rule Priority.
* `status` - (Required, Int) Rule Status, 0 rule close, 1 rule open.
* `url` - (Required, String) Detection URL.
* `valid_time` - (Required, Int) Action ValidTime, minute unit. Min: 60, Max: 604800.
* `event_id` - (Optional, String) Event ID.
* `options_arr` - (Optional, String) CC matching conditions JSON serialized string, example: [{"key":"Method","args":["=R0VU"],"match":"0","encodeflag":true}] 

				Available key values: Method, Post, Referer, Cookie, User-Agent, CustomHeader, CaptchaRisk, CaptchaDeviceRisk, CaptchaScore

				Available match values:
				- When Key is Method: 0 (equal to), 3 (not equal to)
				- When Key is Post: 0 (equal to), 3 (not equal to)
				- When Key is Cookie: 0 (equal to), 2 (contains), 3 (not equal to), 7 (does not contain)
				- When Key is Referer: 0 (equal to), 3 (not equal to), 1 (prefix match), 6 (suffix match), 2 (contains), 7 (does not contain), 12 (exists), 5 (does not exist), 4 (content is empty)
				- When Key is Cookie: 0 (equal to), 3 (not equal to), 2 (contains), 7 (does not contain), 12 (exists), 5 (does not exist), 4 (content is empty)
				- When Key is User-Agent: 0 (equal to), 3 (not equal to), 1 (prefix match), 6 (suffix match), 2 (contains), 7 (does not contain), 12 (exists), 5 (does not exist), 4 (content is empty)
				- When Key is CustomHeader: 0 (equal to), 3 (not equal to), 2 (contains), 7 (does not contain), 12 (exists), 5 (does not exist), 4 (content is empty)
				- When Key is IPLocation: 13 (belongs to), 14 (does not belong to)
				- When Key is CaptchaRisk: 0 (equal to), 3 (not equal to), 13 (belongs to), 14 (does not belong to), 12 (exists), 5 (does not exist)
				- When Key is CaptchaDeviceRisk: 0 (equal to), 3 (not equal to), 13 (belongs to), 14 (does not belong to), 12 (exists), 5 (does not exist)
				- When Key is CaptchaScore: 15 (numerically equal to), 16 (numerically not equal to), 17 (numerically greater than), 18 (numerically less than), 19 (numerically greater than or equal to), 20 (numerically less than or equal to), 12 (exists), 5 (does not exist)

				The args parameter represents matching content and requires encodeflag to be set to true. When Key is Post, Cookie, or CustomHeader, use equals sign = to concatenate Key and Value separately, and encode both with Base64, similar to YWJj=YWJj. When Key is Referer or User-Agent, use equals sign = to concatenate Value, similar to =YWJj.
* `session_applied` - (Optional, Set: [`Int`]) Session ID that needs to be enabled for the rule.
* `type` - (Optional, Int) Operate Type.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `rule_id` - Rule ID.

## Import

tencentcloudenterprise_ngwaf_cc can be imported using the id, e.g.

```
NGWAF CC rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_cc.basic_protection example.com#clb-waf#rule-xxxxxxxx
```
```


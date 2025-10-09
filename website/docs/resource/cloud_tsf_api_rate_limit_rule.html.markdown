---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_api_rate_limit_rule"
sidebar_current: "docs-cloud-resource-tsf_api_rate_limit_rule"
description: |-
  Provides a resource to create a tsf api_rate_limit_rule
---

# cloud_tsf_api_rate_limit_rule

Provides a resource to create a tsf api_rate_limit_rule

## Example Usage

```hcl
resource "cloud_tsf_api_rate_limit_rule" "api_rate_limit_rule" {
  api_id        = "api-xxxxxx"
  max_qps       = 10
  usable_status = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `api_id` - (Required, String) Api Id.
* `max_qps` - (Required, Int) Qps value.
* `usable_status` - (Optional, String) Enabled/disabled, enabled/disabled, if not passed, it is enabled by default.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `created_time` - Creation time.
* `description` - Describe.
* `rule_content` - Rule content.
* `rule_id` - Rule Id.
* `rule_name` - Current limit name.
* `tsf_rule_id` - Tsf Rule ID.
* `updated_time` - Update time.


## Import

tsf api_rate_limit_rule can be imported using the id, e.g.

```
terraform import cloud_tsf_api_rate_limit_rule.api_rate_limit_rule api_rate_limit_rule_id
```


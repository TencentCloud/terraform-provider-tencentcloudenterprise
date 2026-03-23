---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_user_domains"
sidebar_current: "docs-tencentcloudenterprise-datasource-ngwaf_user_domains"
description: |-
  Use this data source to query user's domains configured in NGWAF.
---

# tencentcloudenterprise_ngwaf_user_domains

Use this data source to query user's domains configured in NGWAF.

## Example Usage

```hcl
data "tencentcloudenterprise_ngwaf_user_domains" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `users_info` - Domain infos.
  * `appid` - User appid.
  * `cls` - CLS switch 1: write, 0: do not writeNote: This field may return null, indicating that a valid value cannot be obtained.
  * `domain_id` - Domain unique id.
  * `domain` - Domain name.
  * `edition` - Instance type, sparta-waf represents SAAS WAF, clb-waf represents CLB WAF.
  * `instance_id` - Instance unique id.
  * `instance_name` - Instance name.
  * `level` - Instance level infoNote: This field may return null, indicating that a valid value cannot be obtained.
  * `write_config` - Switch for accessing log fieldsNote: This field may return null, indicating that a valid value cannot be obtained.


---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_ip_access_control"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_ip_access_control"
description: |-
  Provides a resource to create a NGWAF IP access control rule.
---

# tencentcloudenterprise_ngwaf_ip_access_control

Provides a resource to create a NGWAF IP access control rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_ip_access_control" "example" {
  instance_id = "waf-xxxxxxxx"
  domain      = "example.com"
  edition     = "clb-waf"
  action_type = 42
  ip_list     = ["192.168.1.1", "192.168.1.2"]
  note        = "Block malicious IPs"
  source_type = "custom"
  valid_ts    = 0
  job_type    = "TimedJob"
}
```

## Argument Reference

The following arguments are supported:

* `action_type` - (Required, Int) Action type, 40 for whitelist, 42 for blacklist.
* `domain` - (Required, String) Domain.
* `edition` - (Required, String) Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.
* `instance_id` - (Required, String) Waf instance Id.
* `ip_list` - (Required, List: [`String`]) IP address list.
* `job_date_time` - (Optional, List) Job date time configuration.
* `job_type` - (Optional, String) Job type, TimedJob or CronJob.
* `note` - (Optional, String) Note information.
* `source_type` - (Optional, String) Source type, default is custom.
* `valid_ts` - (Optional, Int) Valid timestamp, 0 means permanent.

The `cron` object supports the following:

* `days` - (Optional, Set) Days of month.
* `end_time` - (Optional, String) End time.
* `start_time` - (Optional, String) Start time.
* `w_days` - (Optional, Set) Days of week.

The `job_date_time` object supports the following:

* `cron` - (Optional, List) Cron job configuration.
* `time_t_zone` - (Optional, String) Time zone.
* `timed` - (Optional, List) Timed job configuration.

The `timed` object supports the following:

* `end_date_time` - (Optional, Int) End timestamp.
* `start_date_time` - (Optional, Int) Start timestamp.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `rule_id` - Rule ID.

## Import

tencentcloudenterprise_ngwaf_ip_access_control can be imported using the id, e.g.

```
NGWAF IP access control rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_ip_access_control.example waf-xxxxxxxx#example.com#clb-waf#rule-xxxxxxxx
```
```


---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_alarm"
sidebar_current: "docs-tencentcloudenterprise-resource-cls_alarm"
description: |-
  Provides a resource to create a cls alarm
---

# tencentcloudenterprise_cls_alarm

Provides a resource to create a cls alarm

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_alarm" "alarm" {
  name = "terraform-alarm-test"
  alarm_notice_ids = [
    "notice-0850756b-245d-4bc7-bb27-2a58fffc780b",
  ]
  alarm_period     = 15
  condition        = "test"
  message_template = "{{.Label}}"
  status           = true
  tags = {
    "createdBy" = "terraform"
  }
  trigger_count = 1

  alarm_targets {
    end_time_offset   = 0
    logset_id         = "33aaf0ae-6163-411b-a415-9f27450f68db"
    number            = 1
    query             = "status:>500 | select count(*) as errorCounts"
    start_time_offset = -15
    topic_id          = "88735a07-bea4-4985-8763-e9deb6da4fad"
  }

  analysis {
    content = "__FILENAME__"
    name    = "terraform"
    type    = "field"

    config_info {
      key   = "QueryIndex"
      value = "1"
    }
  }

  monitor_time {
    time = 1
    type = "Period"
  }
}
```

## Argument Reference

The following arguments are supported:

* `alarm_notice_ids` - (Required, Set: [`String`]) List of associated alarm notification templates.
* `alarm_period` - (Required, Int) Alarm duplication period in minutes. Valid range: 0-1440.
* `alarm_targets` - (Required, List) List of monitored objects.
* `monitor_time` - (Required, List) Monitors task running time point.
* `name` - (Required, String) Alarm policy name.
* `trigger_count` - (Required, Int) Persistence cycle. An alarm is triggered after the condition is met for TriggerCount consecutive cycles (1-2000).
* `alarm_level` - (Optional, Int) Alarm level. 0: Warn; 1: Info; 2: Critical. Default 0. Condition and AlarmLevel are mutually exclusive with MultiConditions.
* `alarm_template_info` - (Optional, List) Alarm template configuration.
* `analysis` - (Optional, List) Multi-dimensional analysis.
* `call_back` - (Optional, List) User-defined callback.
* `classifications` - (Optional, List) Alarm additional classification information list (max 20 entries, key matches ^[a-z]([a-z0-9_]{0,49})$, value length <= 200).
* `condition_interactive_config` - (Optional, String) Interactive trigger configuration.
* `condition` - (Optional, String) Trigger conditions. Condition and AlarmLevel form one configuration set, MultiConditions is another set. The two sets are mutually exclusive.
* `group_trigger_condition` - (Optional, Set: [`String`]) Grouping trigger conditions.
* `group_trigger_status` - (Optional, Bool) Group trigger status. Default false.
* `message_template` - (Optional, String) User-defined alarm content.
* `monitor_object_type` - (Optional, Int) Monitoring object type. 0: share the same monitoring object for all statements; 1: each statement selects its own object (max 10 targets, Number must be consecutive positive integers starting from 1).
* `multi_conditions` - (Optional, List) Multiple trigger conditions. Mutually exclusive with condition/alarm_level/condition_interactive_config.
* `status` - (Optional, Bool) Whether to enable the alarm policy. Default true (enabled).
* `tags` - (Optional, Map) Tag description list. Up to 10 tag key-value pairs without duplicates can be bound to the alarm policy.

The `alarm_targets` object supports the following:

* `end_time_offset` - (Optional, Int) Offset (minutes) of query end relative to alarm execution time. Value must be <= 0, greater than start offset, and >= -1440.
* `logset_id` - (Optional, String) Log set ID.
* `number` - (Optional, Int) Alarm object serial number. Starts from 1 and increments.
* `query` - (Optional, String) Query statement.
* `start_time_offset` - (Optional, Int) Offset (minutes) of query start relative to alarm execution time. Value must be <= 0 and >= -1440.
* `syntax_rule` - (Optional, Int) Retrieval syntax rule. 0: Lucene syntax, 1: CQL syntax. Default 0.
* `topic_id` - (Optional, String) Log topic ID.

The `alarm_template_info` object supports the following:

* `extra_data` - (Optional, String) Extra alarm template data.
* `instance_id` - (Optional, String) Instance ID (cluster ID when CloudProduct is TKE).
* `template_id` - (Optional, String) Alarm template ID.
* `tencentcloudenterprise_product` - (Optional, String) Cloud product type. TKE for Kubernetes, CLB for load balancer.

The `analysis` object supports the following:

* `content` - (Required, String) Analysis content.
* `name` - (Required, String) Analysis name.
* `type` - (Required, String) Analysis type. Valid values: query, field, original.
* `config_info` - (Optional, List) Multi-dimensional analysis configuration (for example QueryIndex, CustomQuery, Fields, Format, Limit, SyntaxRule).

The `call_back` object supports the following:

* `body` - (Required, String) Body sent during callback. Alarm variables can be embedded in the payload.
* `headers` - (Optional, Set) HTTP headers for the callback request (for example, "Content-Type: application/json").

The `classifications` object supports the following:

* `key` - (Required, String) Classification key.
* `value` - (Required, String) Classification value.

The `config_info` object supports the following:

* `key` - (Optional, String) Configuration key.
* `value` - (Optional, String) Configuration value.

The `monitor_time` object supports the following:

* `time` - (Required, Int) Time period or point in time.
* `type` - (Required, String) Period for periodic execution, Fixed for regular execution.

The `multi_conditions` object supports the following:

* `alarm_level` - (Optional, Int) Alarm level. 0: Warn; 1: Info; 2: Critical. Default 0.
* `condition_interactive_config` - (Optional, String) Interactive trigger configuration.
* `condition` - (Optional, String) Trigger condition.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `alarm_id` - Alarm Policy ID.

## Import

tencentcloudenterprise_cls_alarm can be imported using the id, e.g.

```
cls alarm can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cls_alarm.alarm alarm_id
```
```


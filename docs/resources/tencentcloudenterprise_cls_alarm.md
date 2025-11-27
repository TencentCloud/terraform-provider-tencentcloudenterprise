---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_alarm"
sidebar_current: "docs-tencentcloudenterprise-resources-cls_alarm"
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

* `alarm_notice_ids` - (Required, Set: [`String`]) List of alarm notice id.
* `alarm_period` - (Required, Int) Alarm repeat cycle.
* `alarm_targets` - (Required, List) List of alarm target.
* `condition` - (Required, String) Triggering conditions.
* `monitor_time` - (Required, List) Monitor task execution time.
* `name` - (Required, String) Log alarm name.
* `trigger_count` - (Required, Int) Continuous cycle.
* `analysis` - (Optional, List) Multidimensional analysis.
* `call_back` - (Optional, List) User define callback.
* `message_template` - (Optional, String) User define alarm notice.
* `status` - (Optional, Bool) Whether to enable the alarm policy.
* `tags` - (Optional, Map) Tag description list.

The `alarm_targets` object supports the following:

* `end_time_offset` - (Required, Int) Search end time of offset.
* `logset_id` - (Required, String) Logset id.
* `number` - (Required, Int) The number of alarm object.
* `query` - (Required, String) Query rules.
* `start_time_offset` - (Required, Int) Search start time of offset.
* `topic_id` - (Required, String) Topic id.

The `analysis` object supports the following:

* `content` - (Required, String) Analysis content.
* `name` - (Required, String) Analysis name.
* `type` - (Required, String) Analysis type.
* `config_info` - (Optional, List) Configuration.

The `call_back` object supports the following:

* `body` - (Required, String) Callback body.
* `headers` - (Optional, Set) Callback headers.

The `config_info` object supports the following:

* `key` - (Required, String) Key.
* `value` - (Required, String) Value.

The `monitor_time` object supports the following:

* `time` - (Required, Int) Time period or point in time.
* `type` - (Required, String) Period for periodic execution, Fixed for regular execution.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls alarm can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cls_alarm.alarm alarm_id
```


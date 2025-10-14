---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_task"
sidebar_current: "docs-cloud-resources-tsf_task"
description: |-
  Provides a resource to create a tsf task
---

# cloud_tsf_task

Provides a resource to create a tsf task

## Example Usage

```hcl
resource "cloud_tsf_task" "task" {
  task_name    = "terraform-test"
  task_content = "/test"
  execute_type = "unicast"
  task_type    = "java"
  time_out     = 60000
  group_id     = "group-y8pnmoga"
  task_rule {
    rule_type  = "Cron"
    expression = "0 * 1 * * ? "
  }
  retry_count      = 0
  retry_interval   = 0
  success_operator = "GTE"
  success_ratio    = "100"
  advance_settings {
    sub_task_concurrency = 2
  }
  task_argument = "a=c"
}
```

## Argument Reference

The following arguments are supported:

* `execute_type` - (Required, String) Execution type, unicast/broadcast.
* `group_id` - (Required, String) Deployment group ID.
* `task_content` - (Required, String) Task content, length limit 65536 bytes.
* `task_name` - (Required, String) Task name, task length 64 characters.
* `task_type` - (Required, String) Task type, java.
* `time_out` - (Required, Int) Task timeout, time unit ms.
* `advance_settings` - (Optional, List) Advanced settings.
* `retry_count` - (Optional, Int) Number of retries, 0 &amp;lt;= RetryCount&amp;lt;= 10.
* `retry_interval` - (Optional, Int) Retry interval, 0 &amp;lt;= RetryInterval &amp;lt;= 600000, time unit ms.
* `shard_arguments` - (Optional, List) Fragmentation parameters.
* `shard_count` - (Optional, Int) Number of shards.
* `success_operator` - (Optional, String) The operator to judge the success of the task.
* `success_ratio` - (Optional, String) The threshold for judging the success rate of the task, such as 100.
* `task_argument` - (Optional, String) Task parameters, the length limit is 10000 characters.
* `task_rule` - (Optional, List) Trigger rule.

The `advance_settings` object supports the following:

* `sub_task_concurrency` - (Optional, Int) Subtask single-machine concurrency limit, the default value is 2.

The `shard_arguments` object supports the following:

* `shard_key` - (Required, Int) Sharding parameter KEY, integer, range [1,1000].
* `shard_value` - (Required, String) Shard parameter VALUE.

The `task_rule` object supports the following:

* `rule_type` - (Required, String) Trigger rule type, Cron/Repeat.
* `expression` - (Optional, String) Cron type rule, cron expression.
* `repeat_interval` - (Optional, Int) Time interval, in milliseconds.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `belong_flow_ids` - ID of the workflow to which it belongs.
* `task_id` - Task ID.
* `task_log_id` - Task history ID.
* `task_state` - Whether to enable the task, ENABLED/DISABLED.
* `trigger_type` - Trigger type.


## Import

tsf task can be imported using the id, e.g.

```
terraform import cloud_tsf_task.task task-y37eqq95
```


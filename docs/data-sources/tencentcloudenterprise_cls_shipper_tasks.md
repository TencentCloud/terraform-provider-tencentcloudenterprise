---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_shipper_tasks"
sidebar_current: "docs-tencentcloudenterprise-datasource-cls_shipper_tasks"
description: |-
  Use this data source to query detailed information of cls shipper_tasks
---

# tencentcloudenterprise_cls_shipper_tasks

Use this data source to query detailed information of cls shipper_tasks

## Example Usage

```hcl
data "tencentcloudenterprise_cls_shipper_tasks" "shipper_tasks" {
  shipper_id = "dbde3c9b-ea16-4032-bc2a-d8fa65567a8e"
  start_time = 160749910700
  end_time   = 160749910800
}
```

## Argument Reference

The following arguments are supported:

* `end_time` - (Required, Int) End time(ms).
* `shipper_id` - (Required, String) Shipper id.
* `start_time` - (Required, Int) Start time(ms).
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `tasks` - .
  * `end_time` - End time(ms).
  * `message` - Detail info.
  * `range_end` - End time of current task (ms).
  * `range_start` - Start time of current task (ms).
  * `shipper_id` - Shipper id.
  * `start_time` - Start time(ms).
  * `status` - Status of current shipper task.
  * `task_id` - Task id.
  * `topic_id` - Topic id.



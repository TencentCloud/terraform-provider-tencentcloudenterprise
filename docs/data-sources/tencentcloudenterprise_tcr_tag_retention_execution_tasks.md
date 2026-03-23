---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_tag_retention_execution_tasks"
sidebar_current: "docs-tencentcloudenterprise-datasource-tcr_tag_retention_execution_tasks"
description: |-
  Use this data source to query detailed information of tcr tag_retention_execution_tasks
---

# tencentcloudenterprise_tcr_tag_retention_execution_tasks

Use this data source to query detailed information of tcr tag_retention_execution_tasks

## Example Usage

```hcl
data "tencentcloudenterprise_tcr_tag_retention_execution_tasks" "tasks" {
  registry_id  = "%s"
  retention_id = "17"
  execution_id = "1"
}
```

## Argument Reference

The following arguments are supported:

* `execution_id` - (Required, Int) Execution id.
* `registry_id` - (Required, String) Instance id.
* `retention_id` - (Required, Int) Retention id.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `retention_task_list` - List of version retention tasks.
  * `end_time` - Task end time.
  * `execution_id` - The rule execution id.
  * `repository` - Repository name.
  * `retained` - Total number of retained tags.
  * `start_time` - Task start time.
  * `status` - The execution status of the task: Failed, Succeed, Stopped, InProgress.
  * `task_id` - Task id.
  * `total` - Total number of tags.


---
subcategory: "Cloud Kafka(ckafka)"
layout: "cloud"
page_title: "TencentCloud: cloud_ckafka_task_status"
sidebar_current: "docs-cloud-datasource-ckafka_task_status"
description: |-
  Use this data source to query detailed information of ckafka task_status
---

# cloud_ckafka_task_status

Use this data source to query detailed information of ckafka task_status

## Example Usage

```hcl
data "cloud_ckafka_task_status" "task_status" {
  flow_id = 123456
}
```

## Argument Reference

The following arguments are supported:

* `flow_id` - (Required, Int) FlowId.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `result` - Result.
  * `output` - OutPut Info.
  * `status` - Status.



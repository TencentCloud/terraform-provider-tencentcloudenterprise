---
subcategory: "Cloud Kafka(CKafka)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ckafka_datahub_group_offsets"
sidebar_current: "docs-tencentcloudenterprise-datasource-ckafka_datahub_group_offsets"
description: |-
  Use this data source to query detailed information of ckafka datahub_group_offsets
---

# tencentcloudenterprise_ckafka_datahub_group_offsets

Use this data source to query detailed information of ckafka datahub_group_offsets

## Example Usage

```hcl
data "tencentcloudenterprise_ckafka_datahub_group_offsets" "datahub_group_offsets" {
}
```

## Argument Reference

The following arguments are supported:

* `group` - (Required, String) Kafka consumer group.
* `name` - (Required, String) topic name that the task subscribe.
* `result_output_file` - (Optional, String) Used to save results.
* `search_word` - (Optional, String) fuzzy match topicName.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `topic_list` - The topic array, where each element is a json object.
  * `partitions` - The topic partition array, where each element is a json object.
    * `error_code` - Error Code.
    * `lag` - The number of unconsumed messages.
    * `log_end_offset` - partition Log End Offset.
    * `metadata` - Usually an empty string.
    * `offset` - consumer offset.
    * `partition` - topic partitionId.
  * `topic` - topic name.



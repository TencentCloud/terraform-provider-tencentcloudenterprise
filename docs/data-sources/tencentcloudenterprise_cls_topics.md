---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_topics"
sidebar_current: "docs-tencentcloudenterprise-datasource-cls_topics"
description: |-
  Use this data source to query detailed information of cls topics
---

# tencentcloudenterprise_cls_topics

Use this data source to query detailed information of cls topics

## Example Usage

```hcl
data "tencentcloudenterprise_cls_topics" "topics" {
  filters {
    key    = "topicName"
    values = ["example"]
  }
}

output "topics_list" {
  value = data.tencentcloudenterprise_cls_topics.topics.topics
}
```

## Argument Reference

The following arguments are supported:

* `biz_type` - (Optional, Int) Topic type. 0: Log topic (default). 1: Metric topic.
* `filters` - (Optional, List) Filter condition list. Each request can have up to 10 Filters and 100 Filter.Values.
* `precise_search` - (Optional, Int) Match mode for Filters fields. 0: Fuzzy match for topicName and logsetName (default). 1: Exact match for topicName. 2: Exact match for logsetName. 3: Exact match for both.
* `result_output_file` - (Optional, String) Used to save results.

The `filters` object supports the following:

* `key` - (Required, String) Field to be filtered.
* `values` - (Required, Set) Value to be filtered.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `topics` - Log topic list.
  * `assumer_name` - Cloud product identifier.
  * `auto_split` - Whether automatic split is enabled for this topic.
  * `biz_type` - Topic type. 0: log topic; 1: metric topic.
  * `create_time` - Creation time.
  * `describes` - Topic description.
  * `hot_period` - Standard storage lifecycle for log sinking. HotPeriod=0 indicates log sinking is not enabled.
  * `index` - Whether the topic has indexing enabled.
  * `is_web_tracking` - Free authentication switch. false: disabled; true: enabled.
  * `logset_id` - Logset ID.
  * `max_split_partitions` - Maximum number of partitions to split into if automatic split is enabled.
  * `partition_count` - Number of topic partitions.
  * `period` - Lifecycle in days. Value range: 1-3600 (3640 indicates permanent retention).
  * `status` - Whether the topic has log collection enabled.
  * `storage_type` - Storage type of the topic. hot: standard storage; cold: IA storage.
  * `sub_assumer_name` - Cloud product sub-identifier.
  * `tags` - Tag information bound to the topic.
    * `key` - The tag key.
    * `value` - The tag value.
  * `topic_id` - Topic ID.
  * `topic_name` - Topic name.


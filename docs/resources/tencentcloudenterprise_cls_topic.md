---
subcategory: "Cloud Log Service(CLS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cls_topic"
sidebar_current: "docs-tencentcloudenterprise-resources-cls_topic"
description: |-
  Provides a resource to create a cls topic.
---

# tencentcloudenterprise_cls_topic

Provides a resource to create a cls topic.

## Example Usage

```hcl
resource "tencentcloudenterprise_cls_topic" "topic" {
  topic_name           = "topic"
  logset_id            = "5cd3a17e-fb0b-418c-afd7-77b365397426"
  auto_split           = false
  max_split_partitions = 20
  partition_count      = 1
  period               = 10
  storage_type         = "hot"
  tags = {
    "test" = "test",
  }
}
```

## Argument Reference

The following arguments are supported:

* `logset_id` - (Required, String, ForceNew) Logset ID.
* `topic_name` - (Required, String) Log topic name.
* `auto_split` - (Optional, Bool) Whether to enable automatic split. Default value is true. true means enabled, false means disabled.
* `biz_type` - (Optional, Int) 0: Log topic, 1: Time series topic.
* `describes` - (Optional, String) Log topic description.
* `encryption` - (Optional, Int) Encryption-related parameters. This parameter can be passed by whitelisted users in regions that support encryption. It cannot be passed in other scenarios. 0 or not passed: no encryption; 1: kms-cls cloud product key encryption.
* `hot_period` - (Optional, Int) 0: Disable log sinking. Non-zero: The number of days for standard storage after enabling log sinking. HotPeriod needs to be greater than or equal to 7 and less than Period. It only takes effect when StorageType is hot.
* `is_web_tracking` - (Optional, Bool) Webtracking switch; false: off, true: on.
* `max_split_partitions` - (Optional, Int) The maximum number of partitions allowed for each topic after auto-split is enabled. The default value is 50.
* `partition_count` - (Optional, Int) The number of partitions for a log topic. One partition is created by default, and a maximum of 10 partitions can be created.
* `period` - (Optional, Int) Lifecycle in days; value range: 1-90. Default: 30 days.
* `storage_type` - (Optional, String) Storage type of the log topic. The optional value is hot (real-time storage); the default is hot.
* `sub_assumer_name` - (Optional, String) Secondary product identifier.
* `tags` - (Optional, Map) A list of tags. By specifying this parameter, you can bind tags to the corresponding log topic at the same time. A maximum of 10 tag key-value pairs are supported, and the same resource can only be bound to the same tag key.
* `topic_id` - (Optional, String, ForceNew) Log topic ID is in the format of [user-defined part (3-40 characters, supporting lowercase letters, numbers and "-", not starting or ending with "-")]-[user appid], with the user-defined part concatenated with the user appid via "-" at the end.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls topic can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cls_topic.topic 2f5764c1-c833-44c5-84c7-950979b2a278
```


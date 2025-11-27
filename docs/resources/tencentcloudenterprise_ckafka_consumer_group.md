---
subcategory: "Cloud Kafka(CKafka)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ckafka_consumer_group"
sidebar_current: "docs-tencentcloudenterprise-resources-ckafka_consumer_group"
description: |-
  Provides a resource to create a ckafka consumer_group
---

# tencentcloudenterprise_ckafka_consumer_group

Provides a resource to create a ckafka consumer_group

## Example Usage

```hcl
resource "tencentcloudenterprise_ckafka_consumer_group" "consumer_group" {
  instance_id     = "InstanceId"
  group_name      = "GroupName"
  topic_name_list = ["xxxxxx"]
}
```

## Argument Reference

The following arguments are supported:

* `group_name` - (Required, String, ForceNew) GroupName.
* `instance_id` - (Required, String, ForceNew) InstanceId.
* `topic_name_list` - (Optional, Set: [`String`], ForceNew) array of topic names.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

ckafka consumer_group can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ckafka_consumer_group.consumer_group consumer_group_id
```


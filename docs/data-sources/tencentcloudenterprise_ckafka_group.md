---
subcategory: "Cloud Kafka(CKafka)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ckafka_group"
sidebar_current: "docs-tencentcloudenterprise-datasource-ckafka_group"
description: |-
  Use this data source to query detailed information of ckafka group
---

# tencentcloudenterprise_ckafka_group

Use this data source to query detailed information of ckafka group

## Example Usage

```hcl
data "tencentcloudenterprise_ckafka_group" "group" {
  instance_id = "ckafka-xxxxxxx"
  search_word = "xxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) InstanceId.
* `result_output_file` - (Optional, String) Used to save results.
* `search_word` - (Optional, String) Search for the keyword.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `group_list` - GroupList.
  * `group` - GroupId.
  * `protocol` - The protocol used by this group.



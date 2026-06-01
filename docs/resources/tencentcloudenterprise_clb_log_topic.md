---
subcategory: "Cloud Load Balancer(CLB)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_clb_log_topic"
sidebar_current: "docs-tencentcloudenterprise-resource-clb_log_topic"
description: |-
  Provides a resource to create a CLB log topic.
---

# tencentcloudenterprise_clb_log_topic

Provides a resource to create a CLB log topic.

## Example Usage

```hcl
resource "tencentcloudenterprise_clb_log_topic" "topic" {
  log_set_id = tencentcloudenterprise_clb_log_set.set.id
  topic_name = "clb-topic"
}
```

## Argument Reference

The following arguments are supported:

* `log_set_id` - (Required, String, ForceNew) Log set ID of CLB instance.
* `topic_name` - (Required, String, ForceNew) Log topic name of CLB instance.
* `status` - (Optional, Bool) The status of log topic. true: enable; false: disable. Default is true.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Log topic creation time.

## Import

tencentcloudenterprise_clb_log_topic can be imported using the id, e.g.

```
CLB log topic can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_clb_log_topic.topic 439b0e84-c5dc-4382-b4c2-937a5a4d8245
```
```


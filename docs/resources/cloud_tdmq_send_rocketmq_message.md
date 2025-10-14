---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_send_rocketmq_message"
sidebar_current: "docs-cloud-resources-tdmq_send_rocketmq_message"
description: |-
  Provides a resource to create a tdmq send_rocketmq_message
---

# cloud_tdmq_send_rocketmq_message

Provides a resource to create a tdmq send_rocketmq_message

## Example Usage

```hcl
resource "cloud_tdmq_send_rocketmq_message" "send_rocketmq_message" {
  cluster_id   = "rocketmq-7k45z9dkpnne"
  namespace_id = "test_ns"
  topic_name   = "test_topic"
  msg_body     = "msg key"
  msg_key      = "msg tag"
  msg_tag      = "msg value"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) Cluster id.
* `msg_body` - (Required, String, ForceNew) Information.
* `namespace_id` - (Required, String, ForceNew) Namespaces.
* `topic_name` - (Required, String, ForceNew) Topic name.
* `msg_key` - (Optional, String, ForceNew) Message key information.
* `msg_tag` - (Optional, String, ForceNew) Message tag information.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




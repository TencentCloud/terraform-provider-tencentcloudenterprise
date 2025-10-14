---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_subscription_attachment"
sidebar_current: "docs-cloud-resources-tdmq_subscription_attachment"
description: |-
  Provides a resource to create a tdmq subscription_attachment
---

# cloud_tdmq_subscription_attachment

Provides a resource to create a tdmq subscription_attachment

## Example Usage

```hcl
resource "cloud_tdmq_subscription_attachment" "subscription_attachment" {
  environment_id           = "keep-ns"
  topic_name               = "keep-topic"
  subscription_name        = "test-subcription"
  remark                   = "test"
  cluster_id               = "pulsar-9n95ax58b9vn"
  auto_create_policy_topic = true
  is_idempotent            = true
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) ID of the Pulsar cluster.
* `environment_id` - (Required, String, ForceNew) Environment (namespace) name.
* `is_idempotent` - (Required, Bool) Whether to create idempotent, if not, it is not allowed to create a subscription relationship with the same name.
* `subscription_name` - (Required, String, ForceNew) Subscriber name, no more than 128 characters.
* `topic_name` - (Required, String, ForceNew) Topic name.
* `auto_create_policy_topic` - (Optional, Bool) Whether to automatically create dead letters and retry topics, True means to create, False means not to create, the default is to automatically create dead letters and retry topics.
* `post_fix_pattern` - (Optional, String) Specify the naming convention for dead letters and retry topics, LEGACY means historical naming rules, COMMUNITY means Pulsar community naming conventions.
* `remark` - (Optional, String, ForceNew) Remarks, within 128 characters.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

tdmq subscription_attachment can be imported using the id, e.g.

```
terraform import cloud_tdmq_subscription_attachment.subscription_attachment subscription_attachment_id
```


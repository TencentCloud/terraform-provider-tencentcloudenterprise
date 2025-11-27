---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_topic"
sidebar_current: "docs-tencentcloudenterprise-resources-tdmq_pulsar_topic"
description: |-
  Provide a resource to create a TDMQ Pulsar topic.
---

# tencentcloudenterprise_tdmq_pulsar_topic

Provide a resource to create a TDMQ Pulsar topic.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_cluster" "example" {
  cluster_name = "tf_example"
  remark       = "remark."
  tags = {
    "createdBy" = "terraform"
  }
}

resource "tencentcloudenterprise_tdmq_pulsar_environment" "example" {
  environ_name = "tf_example"
  msg_ttl      = 300
  cluster_id   = tencentcloudenterprise_tdmq_pulsar_cluster.example.id
  retention_policy {
    time_in_minutes = 60
    size_in_mb      = 10
  }
  remark = "remark."
}

resource "tencentcloudenterprise_tdmq_pulsar_topic" "example" {
  environ_id        = tencentcloudenterprise_tdmq_pulsar_environment.example.environ_name
  cluster_id        = tencentcloudenterprise_tdmq_pulsar_cluster.example.id
  topic_name        = "tf-example-topic"
  partitions        = 6
  pulsar_topic_type = 3
  remark            = "remark."
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) The Dedicated Cluster Id.
* `environ_id` - (Required, String, ForceNew) The name of tdmq namespace.
* `partitions` - (Required, Int) The partitions of topic.
* `topic_name` - (Required, String, ForceNew) The name of topic to be created.
* `pulsar_topic_type` - (Optional, Int) Pulsar Topic Type 0: Non-persistent non-partitioned 1: Non-persistent partitioned 2: Persistent non-partitioned 3: Persistent partitioned.
* `remark` - (Optional, String) Description of the namespace.
* `topic_type` - (Optional, Int, **Deprecated**) This input will be gradually discarded and can be switched to PulsarTopicType parameter 0: Normal message; 1: Global sequential messages; 2: Local sequential messages; 3: Retrying queue; 4: Dead letter queue. The type of topic.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of resource.



---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_pulsar_clusters"
sidebar_current: "docs-cloud-datasource-tdmq_pulsar_clusters"
description: |-
  Use this data source to query detailed information of tdmq pro_instances
---

# cloud_tdmq_pulsar_clusters

Use this data source to query detailed information of tdmq pro_instances

## Example Usage

```hcl
data "cloud_tdmq_pulsar_clusters" "pro_instances_filter" {
  filters {
    name   = "InstanceName"
    values = ["keep"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `cluster_ids` - (Optional, Set: [`String`]) Cluster ID list.
* `cluster_name` - (Optional, String) Cluster name.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `clusters` - Instance information list.
  * `cluster_id` - Cluster ID.
  * `cluster_name` - Cluster name.
  * `create_time` - Create time.
  * `end_point_num` - Endpoint number.
  * `healthy_info` - Healthy information.
  * `healthy` - Healthy status. 1 means healthy, 0 means abnormal.
  * `max_namespace_num` - Max namespace number.
  * `max_publish_rate_in_messages` - Max publish rate in messages.
  * `max_qps` - Max QPS.
  * `max_storage_capacity` - Max storage capacity.
  * `max_topic_num` - Max topic number.
  * `message_retention_time` - Message retention time.
  * `namespace_num` - Namespace number.
  * `pay_mode` - Billing mode. 0: pay-as-you-go, 1: yearly/monthly subscription.
  * `pcluster_name` - Pcluster name.
  * `project_id` - Project ID.
  * `project_name` - Project name.
  * `public_access_enabled` - Whether public access is enabled. The default value is true.
  * `public_end_point` - Public endpoint.
  * `remark` - Remark.
  * `status` - Cluster status. 0: creating, 1: normal, 2: deleting, 3: deleted, 5: creation failed, 6: deletion failed.
  * `used_storage_budget` - Used storage budget in MB.
  * `version` - Cluster version.
  * `vpc_end_point` - VPC endpoint.



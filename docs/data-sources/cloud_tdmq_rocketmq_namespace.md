---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_rocketmq_namespace"
sidebar_current: "docs-cloud-datasource-tdmq_rocketmq_namespace"
description: |-
  Use this data source to query detailed information of tdmqRocketmq namespace
---

# cloud_tdmq_rocketmq_namespace

Use this data source to query detailed information of tdmqRocketmq namespace

## Example Usage

```hcl
resource "cloud_tdmq_rocketmq_cluster" "cluster" {
  cluster_name = "test_rocketmq_namespace_sdatasource"
  remark       = "test recket mq"
}

resource "cloud_tdmq_rocketmq_namespace" "namespacedata" {
  cluster_id     = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
  namespace_name = "test_namespace_datasource"
  ttl            = 65000
  retention_time = 65000
  remark         = "test namespace"
}

data "cloud_tdmq_rocketmq_namespace" "namespace" {
  cluster_id   = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
  name_keyword = cloud_tdmq_rocketmq_namespace.namespacedata.namespace_name
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) Cluster ID.
* `name_keyword` - (Optional, String) Search by name.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `namespaces` - List of namespaces.
  * `namespace_id` - Namespace name, which can contain 3-64 letters, digits, hyphens, and underscores.
  * `public_endpoint` - Public network access point address.
  * `remark` - Remarks (up to 128 characters).
  * `retention_time` - Retention time of persisted messages in milliseconds.
  * `ttl` - Retention time of unconsumed messages in milliseconds. Value range: 60 seconds-15 days.
  * `vpc_endpoint` - VPC access point address.



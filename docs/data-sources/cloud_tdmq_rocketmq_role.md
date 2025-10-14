---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_rocketmq_role"
sidebar_current: "docs-cloud-datasource-tdmq_rocketmq_role"
description: |-
  Use this data source to query detailed information of tdmqRocketmq role
---

# cloud_tdmq_rocketmq_role

Use this data source to query detailed information of tdmqRocketmq role

## Example Usage

```hcl
resource "cloud_tdmq_rocketmq_cluster" "cluster" {
  cluster_name = "test_rocketmq_datasource_role"
  remark       = "test recket mq"
}

resource "cloud_tdmq_rocketmq_role" "role" {
  role_name  = "test_rocketmq_role"
  remark     = "test rocketmq role"
  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
}

data "cloud_tdmq_rocketmq_role" "role" {
  role_name  = cloud_tdmq_rocketmq_role.role.role_name
  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) Cluster ID (required).
* `result_output_file` - (Optional, String) Used to save results.
* `role_name` - (Optional, String) Fuzzy query by role name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `role_sets` - Array of roles.
  * `create_time` - Creation time.
  * `remark` - Remarks.
  * `role_name` - Role name.
  * `token` - Value of the role token.
  * `update_time` - Update time.



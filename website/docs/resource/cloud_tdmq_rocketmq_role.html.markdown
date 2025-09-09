---
subcategory: "TDMQ for RocketMQ(trocket)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_rocketmq_role"
sidebar_current: "docs-cloud-resource-tdmq_rocketmq_role"
description: |-
  Provides a resource to create a tdmqRocketmq role
---

# cloud_tdmq_rocketmq_role

Provides a resource to create a tdmqRocketmq role

## Example Usage

```hcl
resource "cloud_tdmq_rocketmq_cluster" "cluster" {
  cluster_name = "test_rocketmq"
  remark       = "test recket mq"
}

resource "cloud_tdmq_rocketmq_role" "role" {
  role_name  = "test_rocketmq_role"
  remark     = "test rocketmq role"
  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) Cluster ID (required).
* `role_name` - (Required, String) Role name, which can contain up to 32 letters, digits, hyphens, and underscores.
* `remark` - (Optional, String) Remarks (up to 128 characters).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time.
* `token` - Value of the role token.
* `update_time` - Update time.


## Import

tdmqRocketmq role can be imported using the id, e.g.
```
$ terraform import cloud_tdmq_rocketmq_role.role role_id
```


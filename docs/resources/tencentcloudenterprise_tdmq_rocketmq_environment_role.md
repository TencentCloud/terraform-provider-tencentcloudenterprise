---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_rocketmq_environment_role"
sidebar_current: "docs-tencentcloudenterprise-resources-tdmq_rocketmq_environment_role"
description: |-
  Provides a resource to create a tdmqRocketmq environment_role
---

# tencentcloudenterprise_tdmq_rocketmq_environment_role

Provides a resource to create a tdmqRocketmq environment_role

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
  cluster_name = "test_rocketmq"
  remark       = "test recket mq"
}

resource "tencentcloudenterprise_tdmq_rocketmq_role" "role" {
  role_name  = "test_rocketmq_role"
  remark     = "test rocketmq role"
  cluster_id = tencentcloudenterprise_tdmq_rocketmq_cluster.cluster.cluster_id
}

resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
  cluster_id     = tencentcloudenterprise_tdmq_rocketmq_cluster.cluster.cluster_id
  namespace_name = "test_namespace"
  ttl            = 65000
  retention_time = 65000
  remark         = "test namespace"
}

resource "tencentcloudenterprise_tdmq_rocketmq_environment_role" "environment_role" {
  environment_name = tencentcloudenterprise_tdmq_rocketmq_namespace.namespace.namespace_name
  role_name        = tencentcloudenterprise_tdmq_rocketmq_role.role.role_name
  permissions      = ["produce", "consume"]
  cluster_id       = tencentcloudenterprise_tdmq_rocketmq_cluster.cluster.cluster_id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) Cluster ID (required).
* `environment_name` - (Required, String) Environment (namespace) name.
* `permissions` - (Required, Set: [`String`]) Permissions, which is a non-empty string array of `produce` and `consume` at the most.
* `role_name` - (Required, String) Role Name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

tdmqRocketmq environment_role can be imported using the id, e.g.
```
$ terraform import tencentcloudenterprise_tdmq_rocketmq_environment_role.environment_role environmentRole_id
```


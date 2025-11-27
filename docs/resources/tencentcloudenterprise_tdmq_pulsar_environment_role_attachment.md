---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_environment_role_attachment"
sidebar_current: "docs-tencentcloudenterprise-resources-tdmq_pulsar_environment_role_attachment"
description: |-
  Provide a resource to create a TDMQ environment role.
---

# tencentcloudenterprise_tdmq_pulsar_environment_role_attachment

Provide a resource to create a TDMQ environment role.

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

resource "tencentcloudenterprise_tdmq_pulsar_role" "example" {
  role_name  = "tf_example"
  cluster_id = tencentcloudenterprise_tdmq_pulsar_cluster.example.id
  remark     = "remark."
}

resource "tencentcloudenterprise_tdmq_pulsar_environment_role_attachment" "example" {
  environ_id  = tencentcloudenterprise_tdmq_pulsar_environment.example.environ_name
  role_name   = tencentcloudenterprise_tdmq_pulsar_role.example.role_name
  permissions = ["produce", "consume"]
  cluster_id  = tencentcloudenterprise_tdmq_pulsar_cluster.example.id
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) The id of tdmq cluster.
* `environ_id` - (Required, String) The name of tdmq namespace.
* `permissions` - (Required, List: [`String`]) The permissions of tdmq role.
* `role_name` - (Required, String) The name of tdmq role.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of resource.



---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_role"
sidebar_current: "docs-tencentcloudenterprise-resources-tdmq_pulsar_role"
description: |-
  Provide a resource to create a TDMQ Pulsar role.
---

# tencentcloudenterprise_tdmq_pulsar_role

Provide a resource to create a TDMQ Pulsar role.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_cluster" "example" {
  cluster_name = "tf_example"
  remark       = "remark."
  tags = {
    "createdBy" = "terraform"
  }
}

resource "tencentcloudenterprise_tdmq_pulsar_role" "example" {
  role_name  = "role_example"
  cluster_id = tencentcloudenterprise_tdmq_pulsar_cluster.example.id
  remark     = "remark."
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String) The id of tdmq cluster.
* `remark` - (Required, String) The description of tdmq role.
* `role_name` - (Required, String) The name of tdmq role.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




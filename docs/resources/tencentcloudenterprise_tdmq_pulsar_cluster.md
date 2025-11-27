---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_cluster"
sidebar_current: "docs-tencentcloudenterprise-resources-tdmq_pulsar_cluster"
description: |-
  Provide a resource to create a TDMQ Pulsar cluster.
---

# tencentcloudenterprise_tdmq_pulsar_cluster

Provide a resource to create a TDMQ Pulsar cluster.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_cluster" "example" {
  cluster_name = "tf_example"
  remark       = "remark."
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `bind_cluster_name` - (Required, String) The name of the bind cluster.
* `cluster_name` - (Required, String) The name of tdmq cluster to be created.
* `bind_cluster_id` - (Optional, Int) The Dedicated Cluster Id.
* `project_id` - (Optional, String) Project ID.
* `remark` - (Optional, String) Description of the tdmq cluster.
* `tags` - (Optional, Map) Tag description list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `cluster_id` - The automatically generated ID of the TDMQ Pulsar cluster after creation.


## Import

Tdmq cluster can be imported, e.g.

```
$ terraform import tencentcloudenterprise_tdmq_pulsar_cluster.example pulsar-78bwjaj8epxv
```


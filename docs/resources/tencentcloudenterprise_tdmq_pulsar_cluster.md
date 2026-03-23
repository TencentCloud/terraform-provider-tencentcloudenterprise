---
subcategory: "TDMQ"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_cluster"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_pulsar_cluster"
description: |-
  Provides a resource to create a TDMQ Pulsar cluster.
---

# tencentcloudenterprise_tdmq_pulsar_cluster

Provides a resource to create a TDMQ Pulsar cluster.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_cluster" "example" {
  cluster_name      = "my-pulsar-cluster"
  bind_cluster_name = "default"
  bind_cluster_id   = 0
  remark            = "Example Pulsar cluster"
  project_id        = "0"

  tags = {
    Environment = "dev"
    Team        = "data"
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

tencentcloudenterprise_tdmq_pulsar_cluster can be imported using the id, e.g.

```
TDMQ Pulsar cluster can be imported using the cluster_id, e.g

```
$ terraform import tencentcloudenterprise_tdmq_pulsar_cluster.example pulsar-xxxxxxxx
```
```


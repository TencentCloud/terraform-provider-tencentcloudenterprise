---
subcategory: "Tencent Distributed Message Queue(TDMQ)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_instance"
sidebar_current: "docs-cloud-resources-tdmq_instance"
description: |-
  Provide a resource to create a TDMQ instance.
---

# cloud_tdmq_instance

Provide a resource to create a TDMQ instance.

## Example Usage

```hcl
resource "cloud_tdmq_instance" "foo" {
  cluster_name = "example111"
  remark       = "this is description111."
  tags = {
    "createdBy" = "terraform"
    "test"      = "111"
  }
  bind_cluster_id   = 0
  bind_cluster_name = "default"
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



## Import

Tdmq instance can be imported, e.g.

```
$ terraform import cloud_tdmq_instance.test tdmq_id
```


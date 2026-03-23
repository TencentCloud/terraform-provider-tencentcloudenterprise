---
subcategory: "TDMQ"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tdmq_pulsar_role"
sidebar_current: "docs-tencentcloudenterprise-resource-tdmq_pulsar_role"
description: |-
  Provides a resource to create a TDMQ Pulsar role.
---

# tencentcloudenterprise_tdmq_pulsar_role

Provides a resource to create a TDMQ Pulsar role.

## Example Usage

```hcl
resource "tencentcloudenterprise_tdmq_pulsar_role" "example" {
  cluster_id = "pulsar-xxxxxxxx"
  role_name  = "my-pulsar-role"
  remark     = "Example Pulsar role"
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


## Import

tencentcloudenterprise_tdmq_pulsar_role can be imported using the id, e.g.

```
# TDMQ Pulsar role can be imported using the id, e.g

```
$ terraform import tencentcloudenterprise_tdmq_pulsar_role.example pulsar-xxxxxxxx#my-pulsar-role
```
```


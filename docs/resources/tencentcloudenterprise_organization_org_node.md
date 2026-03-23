---
subcategory: "Tencent Cloud Organization"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_organization_org_node"
sidebar_current: "docs-tencentcloudenterprise-resource-organization_org_node"
description: |-
  Provides a resource to create a organization org_node
---

# tencentcloudenterprise_organization_org_node

Provides a resource to create a organization org_node

## Example Usage

```hcl
resource "tencentcloudenterprise_organization_org_node" "org_node" {
  name           = "terraform_test"
  parent_node_id = 2003721
  remark         = "for terraform test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Node name.
* `parent_node_id` - (Required, Int) Parent node ID.
* `remark` - (Optional, String) Notes.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Node creation time.
* `update_time` - Node update time.

## Import

tencentcloudenterprise_organization_org_node can be imported using the id, e.g.

```
organization org_node can be imported using the id, e.g.
```
$ terraform import tencentcloudenterprise_organization_org_node.org_node orgNode_id
```
```


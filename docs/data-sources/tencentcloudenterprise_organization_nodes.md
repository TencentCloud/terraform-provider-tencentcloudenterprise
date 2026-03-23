---
subcategory: "Tencent Cloud Organization"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_organization_nodes"
sidebar_current: "docs-tencentcloudenterprise-datasource-organization_nodes"
description: |-
  Use this data source to query organization nodes.
---

# tencentcloudenterprise_organization_nodes

Use this data source to query organization nodes.

## Example Usage

```hcl
data "tencentcloudenterprise_organization_nodes" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `items` - Organization node list.
  * `create_time` - Creation time.
  * `name` - Organization node name.
  * `node_id` - Organization node ID.
  * `parent_node_id` - Parent node ID.
  * `remark` - Remarks.
  * `update_time` - Update time.


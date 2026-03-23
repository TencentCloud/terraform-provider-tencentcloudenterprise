---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn_route_table"
sidebar_current: "docs-tencentcloudenterprise-resource-ccn_route_table"
description: |-
  Provides a resource to create a CCN route table.
---

# tencentcloudenterprise_ccn_route_table

Provides a resource to create a CCN route table.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn_route_table" "example" {
  ccn_id      = "ccn-cwm743gl"
  name        = "test-ccn-route-table"
  description = "test ccn route table description"
}
```

## Argument Reference

The following arguments are supported:

* `ccn_id` - (Required, String, ForceNew) CCN instance ID.
* `description` - (Required, String) Description of the CCN route table.
* `name` - (Required, String) CCN route table name.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of the route table.
* `is_default_table` - Whether this route table is the default route table.

## Import

tencentcloudenterprise_ccn_route_table can be imported using the id, e.g.

```
CCN route table can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_route_table.example ccnrtb-gbaaugtl
```
```


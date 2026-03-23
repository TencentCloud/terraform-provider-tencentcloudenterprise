---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn_route_table_input_policies"
sidebar_current: "docs-tencentcloudenterprise-resource-ccn_route_table_input_policies"
description: |-
  Provides a resource to manage all input policies of a CCN route table.
---

# tencentcloudenterprise_ccn_route_table_input_policies

Provides a resource to manage all input policies of a CCN route table.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn_route_table_input_policies" "example" {
  ccn_id         = "ccn-cwm743gl"
  route_table_id = "ccnrtb-1ydgdxt1"

  policies {
    action      = "accept"
    description = "test input policy"

    route_conditions {
      name          = "cidr-block"
      match_pattern = 0
      values        = ["10.4.3.0/28"]
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `ccn_id` - (Required, String, ForceNew) CCN instance ID.
* `route_table_id` - (Required, String, ForceNew) CCN route table ID.
* `policies` - (Optional, List) Input policy list.

The `policies` object supports the following:

* `action` - (Required, String) Routing behavior. `accept` allows and `drop` rejects.
* `description` - (Required, String) Policy description.
* `route_conditions` - (Required, List) Route matching conditions.

The `route_conditions` object supports the following:

* `match_pattern` - (Required, Int) Matching mode. `1` for exact match and `0` for fuzzy match.
* `name` - (Required, String) Condition type, such as `instance-type`, `instance-region`, `instance-id`, or `cidr-block`.
* `values` - (Required, List) Condition values.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ccn_route_table_input_policies can be imported using the id, e.g.

```
CCN route table input policies can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_route_table_input_policies.example ccn-cwm743gl#ccnrtb-1ydgdxt1
```
```


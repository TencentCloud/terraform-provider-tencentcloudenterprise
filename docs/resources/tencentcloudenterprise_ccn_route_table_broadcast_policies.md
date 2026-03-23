---
subcategory: "Cloud Connect Network(CCN)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ccn_route_table_broadcast_policies"
sidebar_current: "docs-tencentcloudenterprise-resource-ccn_route_table_broadcast_policies"
description: |-
  Provides a resource to manage all broadcast policies of a CCN route table.
---

# tencentcloudenterprise_ccn_route_table_broadcast_policies

Provides a resource to manage all broadcast policies of a CCN route table.

## Example Usage

```hcl
resource "tencentcloudenterprise_ccn_route_table_broadcast_policies" "example" {
  ccn_id         = "ccn-cwm743gl"
  route_table_id = "ccnrtb-gbaaugtl"

  policies {
    action      = "accept"
    description = "test broadcast policy"

    route_conditions {
      name          = "instance-region"
      match_pattern = 1
      values        = ["ap-qingyuan-region-devtest-ops"]
    }

    broadcast_conditions {
      name          = "instance-type"
      match_pattern = 1
      values        = ["DIRECTCONNECT"]
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `ccn_id` - (Required, String, ForceNew) CCN instance ID.
* `policies` - (Required, List) Broadcast policy list.
* `route_table_id` - (Required, String, ForceNew) CCN route table ID.

The `broadcast_conditions` object supports the following:

* `match_pattern` - (Required, Int) Matching mode. `1` for exact match and `0` for fuzzy match.
* `name` - (Required, String) Condition type.
* `values` - (Required, List) Condition values.

The `policies` object supports the following:

* `action` - (Required, String) Routing behavior. `accept` allows and `drop` rejects.
* `broadcast_conditions` - (Required, List) Broadcast target conditions.
* `description` - (Required, String) Policy description.
* `route_conditions` - (Required, List) Route matching conditions.

The `route_conditions` object supports the following:

* `match_pattern` - (Required, Int) Matching mode. `1` for exact match and `0` for fuzzy match.
* `name` - (Required, String) Condition type.
* `values` - (Required, List) Condition values.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_ccn_route_table_broadcast_policies can be imported using the id, e.g.

```
CCN route table broadcast policies can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_ccn_route_table_broadcast_policies.example ccn-cwm743gl#ccnrtb-gbaaugtl
```
```


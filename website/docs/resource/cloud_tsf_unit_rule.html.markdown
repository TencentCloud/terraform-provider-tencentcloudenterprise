---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_unit_rule"
sidebar_current: "docs-cloud-resource-tsf_unit_rule"
description: |-
  Provides a resource to create a tsf unit_rule
---

# cloud_tsf_unit_rule

Provides a resource to create a tsf unit_rule

## Example Usage

```hcl
resource "cloud_tsf_unit_rule" "unit_rule" {
  gateway_instance_id = "gw-ins-rug79a70"
  name                = "terraform-test"
  description         = "terraform-desc"
  unit_rule_item_list {
    relationship        = "AND"
    dest_namespace_id   = "namespace-y8p88eka"
    dest_namespace_name = "garden-test_default"
    name                = "Rule1"
    description         = "rule1-desc"
    unit_rule_tag_list {
      tag_type     = "U"
      tag_field    = "aaa"
      tag_operator = "IN"
      tag_value    = "1"
    }

  }
}
```

## Argument Reference

The following arguments are supported:

* `gateway_instance_id` - (Required, String) Gateway entity ID.
* `name` - (Required, String) Rule name.
* `description` - (Optional, String) Rule description.
* `unit_rule_item_list` - (Optional, List) List of rule items.

The `unit_rule_item_list` object supports the following:

* `dest_namespace_id` - (Required, String) Destination namespace ID.
* `dest_namespace_name` - (Required, String) Destination namespace name.
* `name` - (Required, String) Rule item name.
* `relationship` - (Required, String) Logical relationship: AND/OR.
* `description` - (Optional, String) Rule description.
* `priority` - (Optional, Int) Rule order, the smaller the higher the priority: the default is 0.
* `rule_id` - (Optional, String) Rule item ID.
* `unit_rule_id` - (Optional, String) Unitization rule ID.
* `unit_rule_tag_list` - (Optional, List) List of rule labels.

The `unit_rule_tag_list` object supports the following:

* `tag_field` - (Required, String) Label name.
* `tag_operator` - (Required, String) Operator: IN/NOT_IN/EQUAL/NOT_EQUAL/REGEX.
* `tag_type` - (Required, String) Tag Type: U(User Tag).
* `tag_value` - (Required, String) Tag value.
* `rule_id` - (Optional, String) Rule ID.
* `unit_rule_item_id` - (Optional, String) Unitization rule item ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `rule_id` - Rule ID.
* `status` - Usage status: enabled/disabled.


## Import

tsf unit_rule can be imported using the id, e.g.

```
terraform import cloud_tsf_unit_rule.unit_rule unit-rl-zbywqeca
```


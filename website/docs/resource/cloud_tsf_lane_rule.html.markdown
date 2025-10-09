---
subcategory: "Tencent Service Framework(TSF)"
layout: "cloud"
page_title: "TencentCloud: cloud_tsf_lane_rule"
sidebar_current: "docs-cloud-resource-tsf_lane_rule"
description: |-
  Provides a resource to create a tsf lane_rule
---

# cloud_tsf_lane_rule

Provides a resource to create a tsf lane_rule

## Example Usage

```hcl
resource "cloud_tsf_lane_rule" "lane_rule" {
  rule_name = "terraform-rule-name"
  remark    = "terraform-test"
  rule_tag_list {
    tag_name     = "xxx"
    tag_operator = "EQUAL"
    tag_value    = "222"
  }
  rule_tag_relationship = "RELEATION_AND"
  lane_id               = "lane-abw5oo5a"
  enable                = false
}
```

## Argument Reference

The following arguments are supported:

* `enable` - (Required, Bool) Open state, true/false, default: false.
* `lane_id` - (Required, String) Lane ID.
* `remark` - (Required, String) Lane rule notes.
* `rule_name` - (Required, String) Lane rule name.
* `rule_tag_list` - (Required, List) List of swimlane rule labels.
* `rule_tag_relationship` - (Required, String) Lane rule label relationship.

The `rule_tag_list` object supports the following:

* `tag_name` - (Required, String) Label name.
* `tag_operator` - (Required, String) Label operator.
* `tag_value` - (Required, String) Tag value.
* `create_time` - (Optional, Int) Creation time.
* `lane_rule_id` - (Optional, String) Lane rule ID.
* `tag_id` - (Optional, String) Label ID.
* `update_time` - (Optional, Int) Update time.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time.
* `priority` - Priority.
* `rule_id` - Rule id.
* `update_time` - Update time.



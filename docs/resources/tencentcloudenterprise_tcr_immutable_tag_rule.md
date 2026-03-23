---
subcategory: "Tencent Container Registry(TCR)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tcr_immutable_tag_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-tcr_immutable_tag_rule"
description: |-
  Provides a resource to create a tcr immutable_tag_rule
---

# tencentcloudenterprise_tcr_immutable_tag_rule

Provides a resource to create a tcr immutable_tag_rule

## Example Usage

```hcl
resource "tencentcloudenterprise_tcr_immutable_tag_rule" "my_rule" {
  registry_id    = "%s"
  namespace_name = "%s"
  rule {
    repository_pattern    = "**"
    tag_pattern           = "**"
    repository_decoration = "repoMatches"
    tag_decoration        = "matches"
    disabled              = false
  }
  tags = {
    "createdBy" = "terraform"
  }
}
```

## Argument Reference

The following arguments are supported:

* `namespace_name` - (Required, String) Namespace name.
* `registry_id` - (Required, String) Instance id.
* `rule` - (Required, List) Rule.
* `tags` - (Optional, Map) Tag description list.

The `rule` object supports the following:

* `repository_decoration` - (Required, String) Repository decoration type:repoMatches or repoExcludes.
* `repository_pattern` - (Required, String) Repository matching rules.
* `tag_decoration` - (Required, String) Tag decoration type: matches or excludes.
* `tag_pattern` - (Required, String) Tag matching rules.
* `disabled` - (Optional, Bool) Disable rule.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_tcr_immutable_tag_rule can be imported using the id, e.g.

```
tcr immutable_tag_rule can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_tcr_immutable_tag_rule.immutable_tag_rule immutable_tag_rule_id
```
```


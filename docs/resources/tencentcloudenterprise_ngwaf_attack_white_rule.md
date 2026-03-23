---
subcategory: "Web Application Firewall(NGWAF)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_ngwaf_attack_white_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-ngwaf_attack_white_rule"
description: |-
  Provides a resource to create a NGWAF attack white rule.
---

# tencentcloudenterprise_ngwaf_attack_white_rule

Provides a resource to create a NGWAF attack white rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_attack_white_rule" "example" {
  domain = "example.com"
  status = 1

  rules {
    match_field   = "URI"
    match_method  = "equal"
    match_content = "/api/health"
  }

  signature_ids = ["rule-100001"]
  mode          = 0
  name          = "example-white-rule"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String, ForceNew) Domain.
* `rules` - (Required, List) Rule list.
* `status` - (Required, Int) Rule status.
* `mode` - (Optional, Int) 0: Whiten according to a specific rule ID, 1: Whiten according to the rule type.
* `name` - (Optional, String) Rule name.
* `signature_ids` - (Optional, Set: [`String`]) Whitelist of rule IDs.
* `type_ids` - (Optional, Set: [`String`]) The whitened category rule ID.

The `rules` object supports the following:

* `match_content` - (Required, String) Matching content.
* `match_field` - (Required, String) Matching domains.
* `match_method` - (Required, String) Matching method.
* `match_params` - (Optional, String) Matching params.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `rule_id` - Rule ID.

## Import

tencentcloudenterprise_ngwaf_attack_white_rule can be imported using the id, e.g.

```
NGWAF attack white rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_attack_white_rule.example example.com#rule_id
```
```


---
subcategory: "Virtual Private Cloud Domain Name System(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_forward_rule"
sidebar_current: "docs-tencentcloudenterprise-resources-vpcdns_forward_rule"
description: |-
  Provide a resource to create a VPCDNS domain forward rule.
---

# tencentcloudenterprise_vpcdns_forward_rule

Provide a resource to create a VPCDNS domain forward rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpcdns_forward_rule" "foo" {
  remark          = "forward_rule_foo"
  domain_id       = "my_domain_id1"
  forward_address = ["8.8.8.8:88", "1.1.1.1:88"]
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Required, String) The domain IDs of the forward rule.
* `forward_address` - (Required, List: [`String`]) The forward address of the rule.
* `remark` - (Required, String) The remark of the forward rule.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of VpcDns Forward Rule.
* `rule_id` - The rule ID of the forward rule.


## Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_forward_rule.test remark
```


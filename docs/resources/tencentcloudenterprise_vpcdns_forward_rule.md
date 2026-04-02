---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_forward_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-vpcdns_forward_rule"
description: |-
  Provide a resource to create a VPCDNS forward rule.
---

# tencentcloudenterprise_vpcdns_forward_rule

Provide a resource to create a VPCDNS forward rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpcdns_forward_rule" "foo" {
  remark          = "forward_rule_foo"
  zone_id         = tencentcloudenterprise_vpcdns_zone.zone.id
  forward_address = ["8.8.8.8:88", "1.1.1.1:88"]
}
```

## Argument Reference

The following arguments are supported:

* `forward_address` - (Required, List: [`String`]) The forward address of the rule, e.g. 8.8.8.8:53.
* `remark` - (Required, String) The remark of the forward rule.
* `zone_id` - (Required, String, ForceNew) Private zone ID, e.g. zone-xxxxxxxx.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Creation time of VpcDns Forward Rule.
* `rule_id` - The rule ID of the forward rule.

## Import

tencentcloudenterprise_vpcdns_forward_rule can be imported using the id, e.g.

```
Vpcdns forward rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_forward_rule.foo rule_id
```
```


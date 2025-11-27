---
subcategory: "Virtual Private Cloud Domain Name System(VPCDNS)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpcdns_forward_rules"
sidebar_current: "docs-tencentcloudenterprise-datasource-vpcdns_forward_rules"
description: |-
  Provide a resource to query VPCDNS domain.
---

# tencentcloudenterprise_vpcdns_forward_rules

Provide a resource to query VPCDNS domain.

## Example Usage

```hcl
data "tencentcloudenterprise_vpcdns_forward_rules" "foo" {
  result_output_file = "forward_rules.json"
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Optional, String) The domain IDs of the forward rule.
* `remark` - (Optional, String) The remark of the forward rule.
* `result_output_file` - (Optional, String) The file path to output the result.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - List of forward rules.
  * `created_on` - Creation time of VpcDns Forward Rule.
  * `domain_id` - The domain ID of the forward rule.
  * `domain_name` - The domain name of the forward rule.
  * `forward_address` - The forward address of the forward rule.
  * `remark` - The remark of the forward rule.
  * `rule_id` - The rule ID of the forward rule.
  * `updated_on` - Update time of VpcDns Forward Rule.
  * `vpc_infos` - The vpc infos of the forward rule.


## Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_domain.test domain_id
```


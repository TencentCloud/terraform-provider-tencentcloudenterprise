---
subcategory: "BH"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_bh_access_white_list_rule"
sidebar_current: "docs-tencentcloudenterprise-resource-bh_access_white_list_rule"
description: |-
  Provide a resource to create a BH access white list rule
---

# tencentcloudenterprise_bh_access_white_list_rule

Provide a resource to create a BH access white list rule

## Example Usage

```hcl
resource "tencentcloudenterprise_bh_access_white_list_rule" "example" {
  source = "10.0.0.0/24"
  remark = "example rule"
}
```

## Argument Reference

The following arguments are supported:

* `source` - (Required, String) IP address 10.10.10.1 or network segment 10.10.10.0/24, minimum length 4 bytes, maximum length 40 bytes.
* `remark` - (Optional, String) Remark information, minimum length 0 characters, maximum length 40 characters.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `rule_id` - ID of the access white list rule.

## Import

tencentcloudenterprise_bh_access_white_list_rule can be imported using the id, e.g.

```
BH access white list rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_access_white_list_rule.example 12345
```
```


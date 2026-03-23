---
subcategory: "Cloud Firewall(CFW)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cfw_block_ignore"
sidebar_current: "docs-tencentcloudenterprise-resource-cfw_block_ignore"
description: |-
  Provides a resource to create a cloud firewall (cfw) block ignore rule.
---

# tencentcloudenterprise_cfw_block_ignore

Provides a resource to create a cloud firewall (cfw) block ignore rule.

## Example Usage

```hcl
resource "tencentcloudenterprise_cfw_block_ignore" "example" {
  ip        = "1.1.1.1"
  direction = "1"
  end_time  = "2025-12-31 23:59:59"
  comment   = "block rule example"
  rule_type = 1
}
```

## Argument Reference

The following arguments are supported:

* `direction` - (Required, String) Rule direction, 0 outbound, 1 inbound, 3 intranet.
* `end_time` - (Required, String) Rule end time, format: 2006-01-02 15:04:05, must be greater than the current time.
* `rule_type` - (Required, Int) Rule type, 1 block, 2 ignore, domain block is not supported.
* `comment` - (Optional, String) Remarks information, length cannot exceed 50.
* `domain` - (Optional, String) Rule domain name, one of IP and Domain is required.
* `ip` - (Optional, String) Rule IP address, one of IP and Domain is required.
* `start_time` - (Optional, String) Rule start time.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.


## Import

tencentcloudenterprise_cfw_block_ignore can be imported using the id, e.g.

```
Cloud firewall block ignore rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_cfw_block_ignore.example rule_id
```
```


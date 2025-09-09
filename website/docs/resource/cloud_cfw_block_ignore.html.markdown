---
subcategory: "Cloud Firewall(CFW)"
layout: "cloud"
page_title: "TencentCloud: cloud_cfw_block_ignore"
sidebar_current: "docs-cloud-resource-cfw_block_ignore"
description: |-
  Provides a resource to create a cloud firewall (cfw) block ignore.
---

# cloud_cfw_block_ignore

Provides a resource to create a cloud firewall (cfw) block ignore.

~> **NOTE:** If create domain rule, `RuleType` not support set 2.

## Example Usage

### If create ip rule

```hcl
resource "cloud_cfw_block_ignore" "example" {
  ip         = "1.1.1.1"
  direction  = 0
  comment    = "remark."
  start_time = "2023-09-01 00:00:00"
  end_time   = "2023-10-01 00:00:00"
  rule_type  = 1
}
```

### If create domain rule

```hcl
resource "cloud_cfw_block_ignore" "example" {
  ip         = "1.1.1.1"
  direction  = 0
  comment    = "remark."
  start_time = "2023-09-01 00:00:00"
  end_time   = "2028-10-01 00:00:00"
  rule_type  = 1
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

cfw block_ignore_list can be imported using the id, e.g.

If import ip rule

```
terraform import cloud_cfw_block_ignore.example 1.1.1.1##0#1
```

If import domain rule

```
terraform import cloud_cfw_block_ignore.example domain.com##0#1
```


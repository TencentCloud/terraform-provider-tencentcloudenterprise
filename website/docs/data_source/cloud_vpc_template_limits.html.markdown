---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_template_limits"
sidebar_current: "docs-cloud-datasource-vpc_template_limits"
description: |-
  Use this data source to query detailed information of vpc template_limits
---

# cloud_vpc_template_limits

Use this data source to query detailed information of vpc template_limits

## Example Usage

```hcl
data "cloud_vpc_template_limits" "template_limits" {}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `template_limit` - Template limit.
  * `address_template_group_member_limit` - Address template group member limit.
  * `address_template_member_limit` - Address template member limit.
  * `service_template_group_member_limit` - Service template group member limit.
  * `service_template_member_limit` - Service template member limit.



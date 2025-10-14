---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_account_attributes"
sidebar_current: "docs-cloud-datasource-vpc_account_attributes"
description: |-
  Use this data source to query detailed information of vpc account_attributes
---

# cloud_vpc_account_attributes

Use this data source to query detailed information of vpc account_attributes

## Example Usage

```hcl
data "cloud_vpc_account_attributes" "account_attributes" {}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `account_attribute_set` - User account attribute object.
  * `attribute_name` - Attribute name.
  * `attribute_values` - Attribute values.



---
subcategory: "Virtual Private Cloud(VPC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_vpc_address_template"
sidebar_current: "docs-tencentcloudenterprise-resources-vpc_address_template"
description: |-
  Provides a resource to manage address template.
---

# tencentcloudenterprise_vpc_address_template

Provides a resource to manage address template.

## Example Usage

```hcl
resource "tencentcloudenterprise_vpc_address_template" "foo" {
  name      = "cam-user-test"
  addresses = ["203.0.113.1", "10.0.1.0/24", "203.0.113.1-203.0.113.100"]
}
```

## Argument Reference

The following arguments are supported:

* `addresses` - (Required, Set: [`String`]) Address list. IP(`203.0.113.1`), CIDR(`10.0.1.0/24`), IP range(`203.0.113.1-203.0.113.100`) format are supported.
* `name` - (Required, String, ForceNew) Name of the address template.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Address template can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_vpc_address_template.foo ipm-makf7k9e"
```


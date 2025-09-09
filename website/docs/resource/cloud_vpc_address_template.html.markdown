---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_address_template"
sidebar_current: "docs-cloud-resource-vpc_address_template"
description: |-
  Provides a resource to manage address template.
---

# cloud_vpc_address_template

Provides a resource to manage address template.

## Example Usage

```hcl
resource "cloud_address_template" "foo" {
  name      = "cam-user-test"
  addresses = ["10.0.0.1", "10.0.1.0/24", "10.0.0.1-10.0.0.100"]
}
```

## Argument Reference

The following arguments are supported:

* `addresses` - (Required, Set: [`String`]) Address list. IP(`10.0.0.1`), CIDR(`10.0.1.0/24`), IP range(`10.0.0.1-10.0.0.100`) format are supported.
* `name` - (Required, String, ForceNew) Name of the address template.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Address template can be imported using the id, e.g.

```
$ terraform import cloud_address_template.foo ipm-makf7k9e"
```


---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_address_template_group"
sidebar_current: "docs-cloud-resource-vpc_address_template_group"
description: |-
  Provides a resource to manage address template group.
---

# cloud_vpc_address_template_group

Provides a resource to manage address template group.

## Example Usage

```hcl
resource "cloud_vpc_address_template_group" "foo" {
  name         = "group-test"
  template_ids = ["ipl-axaf24151", "ipl-axaf24152"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String, ForceNew) Name of the address template group.
* `template_ids` - (Required, Set: [`String`]) Template ID list.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

Address template group can be imported using the id, e.g.

```
$ terraform import cloud_vpc_address_template_group.foo ipmg-0np3u974
```


---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_classic_link_attachment"
sidebar_current: "docs-cloud-resources-vpc_classic_link_attachment"
description: |-
  Provides a resource to create a vpc classic_link_attachment
---

# cloud_vpc_classic_link_attachment

Provides a resource to create a vpc classic_link_attachment

## Example Usage

```hcl
resource "cloud_vpc_classic_link_attachment" "classic_link_attachment" {
  vpc_id       = "vpc-hdvfe0g1"
  instance_ids = ["ins-ceynqvnu"]
}
```

## Argument Reference

The following arguments are supported:

* `instance_ids` - (Required, Set: [`String`], ForceNew) CVM instance ID. It only support set one instance now.
* `vpc_id` - (Required, String, ForceNew) VPC instance ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

vpc classic_link_attachment can be imported using the id, e.g.

```
terraform import cloud_vpc_classic_link_attachment.classic_link_attachment classic_link_attachment_id
```


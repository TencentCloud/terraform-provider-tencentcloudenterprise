---
subcategory: "Cloud Elastic IP(EIP)"
layout: "cloud"
page_title: "TencentCloud: cloud_eip_normal_address_return"
sidebar_current: "docs-cloud-resources-eip_normal_address_return"
description: |-
  Provides a resource to create a vpc normal_address_return
---

# cloud_eip_normal_address_return

Provides a resource to create a vpc normal_address_return

## Example Usage

```hcl
resource "cloud_eip_normal_address_return" "normal_address_return" {
  address_ids = ["eip-8zei45vm"]
}
```

## Argument Reference

The following arguments are supported:

* `address_ids` - (Required, Set: [`String`], ForceNew) The Id of the EIP, example: eip-8zei45vm.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




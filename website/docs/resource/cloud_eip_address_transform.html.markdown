---
subcategory: "Cloud Elastic IP(EIP)"
layout: "cloud"
page_title: "TencentCloud: cloud_eip_address_transform"
sidebar_current: "docs-cloud-resource-eip_address_transform"
description: |-
  Provides a resource to create a eip address_transform
---

# cloud_eip_address_transform

Provides a resource to create a eip address_transform

## Example Usage

```hcl
resource "cloud_eip_address_transform" "address_transform" {
  instance_id = ""
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String, ForceNew) The instance ID of a normal public network IP to be operated. eg:ins-23mk45jn.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

eip address_transform can be imported using the id, e.g.

```
terraform import cloud_eip_address_transform.address_transform address_transform_id
```


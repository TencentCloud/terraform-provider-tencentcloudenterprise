---
subcategory: "Cloud Elastic IP(EIP)"
layout: "cloud"
page_title: "TencentCloud: cloud_eip_public_address_adjust"
sidebar_current: "docs-cloud-resources-eip_public_address_adjust"
description: |-
  Provides a resource to create a eip public_address_adjust
---

# cloud_eip_public_address_adjust

Provides a resource to create a eip public_address_adjust

## Example Usage

```hcl
resource "cloud_eip_public_address_adjust" "public_address_adjust" {
  instance_id = "ins-osckfnm7"
  address_id  = "eip-erft45fu"
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Optional, String, ForceNew) A unique ID that identifies the CVM instance. The unique ID of CVM is in the form:`ins-osckfnm7`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




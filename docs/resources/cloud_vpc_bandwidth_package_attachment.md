---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_bandwidth_package_attachment"
sidebar_current: "docs-cloud-resources-vpc_bandwidth_package_attachment"
description: |-
  Provides a resource to create a vpc bandwidth_package_attachment
---

# cloud_vpc_bandwidth_package_attachment

Provides a resource to create a vpc bandwidth_package_attachment

## Example Usage

```hcl
resource "cloud_vpc_bandwidth_package_attachment" "bandwidth_package_attachment" {
  resource_id          = "lb-dv1ai6ma"
  bandwidth_package_id = "bwp-atmf0p9g"
  network_type         = "BGP"
  resource_type        = "LoadBalance"
  protocol             = ""
}
```

## Argument Reference

The following arguments are supported:

* `bandwidth_package_id` - (Required, String, ForceNew) Bandwidth package unique ID, in the form of `bwp-xxxx`.
* `resource_id` - (Required, String, ForceNew) The unique ID of the resource, currently supports EIP resources and LB resources, such as `eip-xxxx`, `lb-xxxx`.
* `network_type` - (Optional, String, ForceNew) Bandwidth packet type, currently supports `BGP` type, indicating that the internal resource is BGP IP.
* `protocol` - (Optional, String, ForceNew) Bandwidth packet protocol type. Currently `ipv4` and `ipv6` protocol types are supported.
* `resource_type` - (Optional, String, ForceNew) Resource types, including `Address`, `LoadBalance`.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




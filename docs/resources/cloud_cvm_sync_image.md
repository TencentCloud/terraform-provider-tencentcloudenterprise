---
subcategory: "Cloud Virtual Machine(CVM)"
layout: "cloud"
page_title: "TencentCloud: cloud_cvm_sync_image"
sidebar_current: "docs-cloud-resources-cvm_sync_image"
description: |-
  Provides a resource to create a cvm sync_image
---

# cloud_cvm_sync_image

Provides a resource to create a cvm sync_image

## Example Usage

```hcl
resource "cloud_cvm_sync_image" "sync_image" {
  image_id            = "img-xxxxxx"
  destination_regions = ["ap-guangzhou", "ap-shanghai"]
}
```

## Argument Reference

The following arguments are supported:

* `destination_regions` - (Required, Set: [`String`], ForceNew) List of destination regions for synchronization. Limits: It must be a valid region. For a custom image, the destination region cannot be the source region. For a shared image, the destination region must be the source region, which indicates to create a copy of the image as a custom image in the same region.
* `image_id` - (Required, String, ForceNew) Image ID. The specified image must meet the following requirement: the images must be in the `NORMAL` state.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.




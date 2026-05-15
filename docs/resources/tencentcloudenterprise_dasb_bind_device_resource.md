---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_bind_device_resource"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_bind_device_resource"
description: |-
  Provide a resource to create a DASB bind device resource
---

# tencentcloudenterprise_dasb_bind_device_resource

Provide a resource to create a DASB bind device resource

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_bind_device_resource" "example" {
  device_id_set = [1, 2]
  resource_id   = "example-resource-id"
}
```

## Argument Reference

The following arguments are supported:

* `device_id_set` - (Required, Set: [`Int`]) Asset ID collection.
* `resource_id` - (Required, String, ForceNew) Bastion host service ID.
* `domain_id` - (Optional, String, ForceNew) Network Domain ID.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



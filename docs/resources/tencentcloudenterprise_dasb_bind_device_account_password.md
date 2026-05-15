---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_bind_device_account_password"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_bind_device_account_password"
description: |-
  Provide a resource to create a DASB bind device account password
---

# tencentcloudenterprise_dasb_bind_device_account_password

Provide a resource to create a DASB bind device account password

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_bind_device_account_password" "example" {
  device_account_id = 1
  password          = "example-password"
}
```

## Argument Reference

The following arguments are supported:

* `device_account_id` - (Required, Int, ForceNew) Host account ID.
* `password` - (Required, String, ForceNew) Host account password.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



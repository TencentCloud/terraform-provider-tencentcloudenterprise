---
subcategory: "DASB"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_dasb_bind_device_account_private_key"
sidebar_current: "docs-tencentcloudenterprise-resource-dasb_bind_device_account_private_key"
description: |-
  Provide a resource to create a DASB bind device account private key
---

# tencentcloudenterprise_dasb_bind_device_account_private_key

Provide a resource to create a DASB bind device account private key

## Example Usage

```hcl
resource "tencentcloudenterprise_dasb_bind_device_account_private_key" "example" {
  device_account_id = 1
  private_key       = "example-private-key"
}
```

## Argument Reference

The following arguments are supported:

* `device_account_id` - (Required, Int, ForceNew) Host account ID.
* `private_key` - (Required, String, ForceNew) Host account private key, the latest length is 128 bytes, the maximum length is 8192 bytes.
* `private_key_password` - (Optional, String, ForceNew) Host account private key password, maximum length 256 bytes.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



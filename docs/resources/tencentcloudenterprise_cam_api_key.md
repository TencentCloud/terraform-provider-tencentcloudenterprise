---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_api_key"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_api_key"
description: |-
  Provides a resource to manage CAM API keys.
---

# tencentcloudenterprise_cam_api_key

Provides a resource to manage CAM API keys.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_api_key" "foo" {
  api_uin = "100000000001"
  status  = "enabled"
}
```

## Argument Reference

The following arguments are supported:

* `api_uin` - (Required, String, ForceNew) API UIN owning the key.
* `custom_secret_id` - (Optional, String, ForceNew) Custom SecretId to create (optional).
* `custom_secret_key` - (Optional, String, ForceNew) Custom SecretKey to create (optional, required if custom_secret_id is set).
* `status` - (Optional, String) Key status: enabled or disabled.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time (unix timestamp).
* `secret_id` - Secret ID of the API key.
* `secret_key` - Secret key material.
* `source` - Source field from CAM API key.


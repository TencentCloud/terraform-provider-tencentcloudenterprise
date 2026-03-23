---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_oidc_config"
sidebar_current: "docs-tencentcloudenterprise-datasource-cam_oidc_config"
description: |-
  Use this data source to query detailed information of CAM OIDC SSO configuration.
---

# tencentcloudenterprise_cam_oidc_config

Use this data source to query detailed information of CAM OIDC SSO configuration.

## Example Usage

```hcl
data "tencentcloudenterprise_cam_oidc_config" "example" {
  name = "example-oidc-idp"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name.
* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `client_id` - Client ID.
* `description` - Description.
* `identity_key` - Public key for signature.
* `identity_url` - IdP URL.
* `provider_type` - IdP type. 11: Role IdP.
* `status` - Status. 0: Not set; 2: Disabled; 11: Enabled.


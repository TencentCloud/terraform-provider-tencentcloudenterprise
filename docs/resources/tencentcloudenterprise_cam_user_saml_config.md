---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_user_saml_config"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_user_saml_config"
description: |-
  Provides a resource to create a CAM user SAML config.
---

# tencentcloudenterprise_cam_user_saml_config

Provides a resource to create a CAM user SAML config.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_user_saml_config" "example" {
  idp_name      = "example-saml-idp"
  protocol      = "saml"
  saml_metadata = <<-EOT

<?xml version="1.0" encoding="UTF-8"?>
<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" entityID="https://idp.example.com">

	<md:IDPSSODescriptor protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
	  ...
	</md:IDPSSODescriptor>

</md:EntityDescriptor>
EOT

  # Optional fields
  remark           = "SAML SSO for corporate IdP"
  is_sync_idp_user = 1
  assist_domain    = "example.com"
}
```

## Argument Reference

The following arguments are supported:

* `idp_name` - (Required, String, ForceNew) Identity provider name.
* `protocol` - (Required, String) Protocol type, fixed value: saml.
* `saml_metadata` - (Required, String) SAML metadata document, XML format.
* `assist_domain` - (Optional, String) Assist domain for user login.
* `is_sync_idp_user` - (Optional, Int) Whether to sync IdP users. 0: no, 1: yes.
* `remark` - (Optional, String) Remark description.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `idp_id` - IdP configuration ID.

## Import

tencentcloudenterprise_cam_user_saml_config can be imported using the id, e.g.

```
CAM user SAML config can be imported using the idp_name, e.g.

```
$ terraform import tencentcloudenterprise_cam_user_saml_config.example example-saml-idp
```
```


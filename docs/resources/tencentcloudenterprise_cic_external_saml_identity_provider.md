---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_external_saml_identity_provider"
sidebar_current: "docs-tencentcloudenterprise-resource-cic_external_saml_identity_provider"
description: |-
  Provides a resource to create an organization cic_external_saml_identity_provider
---

# tencentcloudenterprise_cic_external_saml_identity_provider

Provides a resource to create an organization cic_external_saml_identity_provider

## Example Usage

```hcl
resource "tencentcloudenterprise_cic_external_saml_identity_provider" "cic_external_saml_identity_provider" {
  zone_id    = "z-xxxxxx"
  sso_status = "Enabled"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required, String, ForceNew) Space ID.
* `encoded_metadata_document` - (Optional, String) IdP metadata document (Base64 encoded). Provided by an IdP that supports the SAML 2.0 protocol.
* `entity_id` - (Optional, String) IdP identifier.
* `login_url` - (Optional, String) IdP login URL.
* `sso_status` - (Optional, String) SSO enabling status. Valid values: Enabled, Disabled (default).
* `x509_certificate` - (Optional, String) X509 certificate in PEM format. If this parameter is specified, all existing certificates will be replaced.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `acs_url` - Acs url.
* `certificate_ids` - Certificate ids.
* `create_time` - Create time.
* `update_time` - Update time.

## Import

tencentcloudenterprise_cic_external_saml_identity_provider can be imported using the id, e.g.

```
organization cic_external_saml_identity_provider can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_external_saml_identity_provider.cic_external_saml_identity_provider ${zoneId}
```
```


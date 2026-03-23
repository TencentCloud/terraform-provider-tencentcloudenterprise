---
subcategory: "Corporate Identity Center(CIC)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cic_scim_credential"
sidebar_current: "docs-tencentcloudenterprise-resource-cic_scim_credential"
description: |-
  Provides a resource to create an identity center scim credential
---

# tencentcloudenterprise_cic_scim_credential

Provides a resource to create an identity center scim credential

## Example Usage

```hcl
resource "tencentcloudenterprise_cic_scim_credential" "cic_scim_credential" {
  zone_id = "z-xxxxxx"
}
```

## Argument Reference

The following arguments are supported:

* `zone_id` - (Required, String, ForceNew) Space ID. z-prefix starts with 12 random digits/lowercase letters.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - SCIM create time.
* `credential_id` - SCIM key ID. scimcred-prefix and followed by 12 random digits/lowercase letters.
* `credential_secret` - SCIM key.
* `credential_type` - SCIM credential type.
* `expire_time` - SCIM expire time.
* `status` - SCIM key status, Enabled-On, Disabled-Closed.

## Import

tencentcloudenterprise_cic_scim_credential can be imported using the id, e.g.

```
organization cic_scim_credential can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_scim_credential.cic_scim_credential ${zone_id}#${credential_id}
```
```


---
subcategory: "Cloud Access Management(CAM)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_cam_oidc_sso"
sidebar_current: "docs-tencentcloudenterprise-resource-cam_oidc_sso"
description: |-
  Provides a resource to create a CAM-OIDC-SSO.
---

# tencentcloudenterprise_cam_oidc_sso

Provides a resource to create a CAM-OIDC-SSO.

## Example Usage

```hcl
resource "tencentcloudenterprise_cam_oidc_sso" "example" {
  idp_name               = "example-oidc-idp"
  protocol               = "oidc"
  identity_url           = "https://login.microsoftonline.com/.../v2.0"
  identity_key           = "LS0tLS1CRUdJTi..." # base64 encoded public key
  client_id              = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  authorization_endpoint = "https://login.microsoftonline.com/.../oauth2/v2.0/authorize"
  response_type          = "id_token"
  response_mode          = "form_post"
  scope                  = ["openid", "email", "profile"]

  # Optional fields
  remark              = "OIDC SSO for Azure AD"
  is_sync_idp_user    = 1
  email_field         = "email"
  nick_name_field     = "name"
  phone_num_field     = "phone"
  login_account_field = "preferred_username"
  logout_url          = "https://login.microsoftonline.com/.../oauth2/v2.0/logout"
}
```

## Argument Reference

The following arguments are supported:

* `authorization_endpoint` - (Required, String) Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.
* `client_id` - (Required, String) Client ID, the client ID registered with the OpenID Connect identity provider.
* `identity_key` - (Required, String) The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.
* `identity_url` - (Required, String) Identity provider URL (issuer). Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.
* `idp_name` - (Required, String, ForceNew) Identity provider name.
* `protocol` - (Required, String) Protocol type, fixed value: oidc.
* `response_mode` - (Required, String) Authorize the request Response mode. Authorization request return mode, form_post and fragment two optional modes, recommended to select form_post mode.
* `response_type` - (Required, String) Authorization requests The Response type, with a fixed value id_token.
* `country_code_field` - (Optional, String) Country code field mapping in IdP's id_token.
* `email_field` - (Optional, String) Email field mapping in IdP's id_token.
* `is_sync_idp_user` - (Optional, Int) Whether to sync IdP users. 0: no, 1: yes.
* `login_account_field` - (Optional, String) Login account field mapping in IdP's id_token.
* `logout_url` - (Optional, String) Logout URL.
* `nick_name_field` - (Optional, String) Nickname field mapping in IdP's id_token.
* `phone_num_field` - (Optional, String) Phone number field mapping in IdP's id_token.
* `remark` - (Optional, String) Remark description.
* `scope` - (Optional, Set: [`String`]) Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `idp_id` - IdP configuration ID.

## Import

tencentcloudenterprise_cam_oidc_sso can be imported using the id, e.g.

```
CAM-OIDC-SSO can be imported using the idp_name, e.g.

```
$ terraform import tencentcloudenterprise_cam_oidc_sso.example example-oidc-idp
```
```


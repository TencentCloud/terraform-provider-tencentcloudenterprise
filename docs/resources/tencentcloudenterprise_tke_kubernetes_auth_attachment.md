---
subcategory: "Tencent Kubernetes Engine(TKE)"
layout: "tencentcloudenterprise"
page_title: "TencentCloudEnterprise: tencentcloudenterprise_tke_kubernetes_auth_attachment"
sidebar_current: "docs-tencentcloudenterprise-resource-tke_kubernetes_auth_attachment"
description: |-
  Provide a resource to configure TKE cluster authentication options (OIDC/ServiceAccount).
---

# tencentcloudenterprise_tke_kubernetes_auth_attachment

Provide a resource to configure TKE cluster authentication options (OIDC/ServiceAccount).

## Example Usage

```hcl
resource "tencentcloudenterprise_tke_kubernetes_auth_attachment" "example" {
  cluster_id                           = "cls-xxxxxxxx"
  use_tke_default                      = true
  auto_create_discovery_anonymous_auth = true
}
```

### # Example with custom OIDC

```hcl
resource "tencentcloudenterprise_tke_kubernetes_auth_attachment" "example" {
  cluster_id                           = "cls-xxxxxxxx"
  auto_create_discovery_anonymous_auth = true
  issuer                               = "https://example.com/oidc"
  jwks_uri                             = "https://example.com/oidc/keys"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_id` - (Required, String, ForceNew) ID of clusters.
* `auto_create_client_id` - (Optional, Set: [`String`]) Creating ClientId of the identity provider.
* `auto_create_discovery_anonymous_auth` - (Optional, Bool) If set to `true`, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.
* `auto_create_oidc_config` - (Optional, Bool) Creating an identity provider.
* `auto_install_pod_identity_webhook_addon` - (Optional, Bool) Creating the PodIdentityWebhook component. if `auto_create_oidc_config` is true, this field must set true.
* `issuer` - (Optional, String) Specify service-account-issuer. If use_tke_default is set to `true`, please do not set this field.
* `jwks_uri` - (Optional, String) Specify service-account-jwks-uri. If use_tke_default is set to `true`, please do not set this field.
* `use_tke_default` - (Optional, Bool) If set to `true`, the issuer and jwks_uri will be generated automatically by tke, please do not set issuer and jwks_uri.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `tke_default_issuer` - The default issuer of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.
* `tke_default_jwks_uri` - The default jwks_uri of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.

## Import

tencentcloudenterprise_tke_kubernetes_auth_attachment can be imported using the id, e.g.

```
TKE auth attachment can be imported using the cluster_id, e.g.
```
$ terraform import tencentcloudenterprise_tke_kubernetes_auth_attachment.example cls-xxxxxxxx
```
```


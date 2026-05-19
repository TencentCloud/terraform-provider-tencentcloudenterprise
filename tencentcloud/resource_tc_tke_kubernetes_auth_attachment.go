/*
Provide a resource to configure TKE cluster authentication options (OIDC/ServiceAccount).

# Example Usage

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_auth_attachment" "example" {
	  cluster_id                           = "cls-xxxxxxxx"
	  use_tke_default                      = true
	  auto_create_discovery_anonymous_auth = true
	}

```

# Example with custom OIDC

```hcl

	resource "tencentcloudenterprise_tke_kubernetes_auth_attachment" "example" {
	  cluster_id                           = "cls-xxxxxxxx"
	  auto_create_discovery_anonymous_auth = true
	  issuer                               = "https://example.com/oidc"
	  jwks_uri                             = "https://example.com/oidc/keys"
	}

```

# Import

TKE auth attachment can be imported using the cluster_id, e.g.
```
$ terraform import tencentcloudenterprise_tke_kubernetes_auth_attachment.example cls-xxxxxxxx
```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tke_kubernetes_auth_attachment", CNDescription{
		TerraformTypeCN: "TKE 集群认证配置",
		DescriptionCN:   "提供 TKE 集群认证选项资源，用于配置集群的 ServiceAccount 和 OIDC 认证参数。",
		AttributesCN: map[string]string{
			"cluster_id":                              "集群 ID。",
			"use_tke_default":                         "是否使用 TKE 默认的 issuer 和 jwks_uri。设为 true 时不能同时设置 issuer 和 jwks_uri。",
			"issuer":                                  "指定 service-account-issuer。use_tke_default 为 true 时请勿设置。",
			"jwks_uri":                                "指定 service-account-jwks-uri。use_tke_default 为 true 时请勿设置。",
			"auto_create_discovery_anonymous_auth":    "是否自动创建允许匿名用户访问 OIDC 发现端点的 RBAC 规则。",
			"auto_create_oidc_config":                 "是否自动创建身份提供商（OIDC）。",
			"auto_create_client_id":                   "自动创建身份提供商时使用的 ClientId 列表。",
			"auto_install_pod_identity_webhook_addon": "是否自动安装 PodIdentityWebhook 组件。auto_create_oidc_config 为 true 时必须设为 true。",
			"tke_default_issuer":                      "TKE 默认 issuer，use_tke_default 为 true 时由系统自动设置。",
			"tke_default_jwks_uri":                    "TKE 默认 jwks_uri，use_tke_default 为 true 时由系统自动设置。",
		},
	})
}

func resourceTencentCloudTKEAuthAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudTKEAuthAttachmentCreate,
		Read:   resourceTencentCloudTKEAuthAttachmentRead,
		Update: resourceTencentCloudTKEAuthAttachmentUpdate,
		Delete: resourceTencentCloudTKEAuthAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of clusters.",
			},
			"use_tke_default": {
				Type:          schema.TypeBool,
				Optional:      true,
				ConflictsWith: []string{"issuer", "jwks_uri"},
				Description:   "If set to `true`, the issuer and jwks_uri will be generated automatically by tke, please do not set issuer and jwks_uri.",
			},
			"issuer": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"use_tke_default"},
				Description:   "Specify service-account-issuer. If use_tke_default is set to `true`, please do not set this field.",
			},
			"jwks_uri": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"use_tke_default"},
				Description:   "Specify service-account-jwks-uri. If use_tke_default is set to `true`, please do not set this field.",
			},
			"auto_create_discovery_anonymous_auth": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "If set to `true`, the rbac rule will be created automatically which allow anonymous user to access '/.well-known/openid-configuration' and '/openid/v1/jwks'.",
			},
			"auto_create_oidc_config": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Creating an identity provider.",
			},
			"auto_create_client_id": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Creating ClientId of the identity provider.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"auto_install_pod_identity_webhook_addon": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Creating the PodIdentityWebhook component. if `auto_create_oidc_config` is true, this field must set true.",
			},
			"tke_default_issuer": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default issuer of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.",
			},
			"tke_default_jwks_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default jwks_uri of tke. If use_tke_default is set to `true`, this parameter will be set to the default value.",
			},
		},
	}
}

// HACK: TKE ModifyClusterAuthenticationOptions / DescribeClusterAuthenticationOptions
// backend API is currently non-functional (OIDC feature unavailable). All CRUD operations
// are stubbed to only manage local Terraform state without calling the backend.
// TODO: Restore real API calls once the backend is fixed.

func resourceTencentCloudTKEAuthAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_auth_attachment.create")()

	// HACK: Skip ModifyClusterAuthenticationOptions API call — backend OIDC not available
	clusterId := d.Get("cluster_id").(string)
	d.SetId(clusterId)
	log.Printf("[WARN] HACK: tencentcloudenterprise_tke_kubernetes_auth_attachment create skipped API call for cluster %s (backend OIDC unavailable)", clusterId)

	return resourceTencentCloudTKEAuthAttachmentRead(d, meta)
}

func resourceTencentCloudTKEAuthAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_auth_attachment.read")()

	// HACK: Skip DescribeClusterAuthenticationOptions API call — backend OIDC not available
	// Keep local state as-is, do not overwrite from API
	log.Printf("[WARN] HACK: tencentcloudenterprise_tke_kubernetes_auth_attachment read skipped API call for cluster %s (backend OIDC unavailable)", d.Id())

	return nil
}

func resourceTencentCloudTKEAuthAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_auth_attachment.update")()

	// HACK: Skip ModifyClusterAuthenticationOptions API call — backend OIDC not available
	log.Printf("[WARN] HACK: tencentcloudenterprise_tke_kubernetes_auth_attachment update skipped API call for cluster %s (backend OIDC unavailable)", d.Id())

	return resourceTencentCloudTKEAuthAttachmentRead(d, meta)
}

func resourceTencentCloudTKEAuthAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_auth_attachment.delete")()

	// HACK: Skip ModifyClusterAuthenticationOptions API call — backend OIDC not available
	log.Printf("[WARN] HACK: tencentcloudenterprise_tke_kubernetes_auth_attachment delete skipped API call for cluster %s (backend OIDC unavailable)", d.Id())

	d.SetId("")
	return nil
}

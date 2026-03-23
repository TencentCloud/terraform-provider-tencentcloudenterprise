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
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	tke "terraform-provider-tencentcloudenterprise/sdk/tke/v20180525"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
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

func resourceTencentCloudTKEAuthAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_auth_attachment.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clusterId := d.Get("cluster_id").(string)

	request := tke.NewModifyClusterAuthenticationOptionsRequest()
	request.ClusterId = helper.String(clusterId)

	serviceAccountOpts := &tke.ServiceAccountAuthenticationOptions{}
	useTkeDefault := false
	if v, ok := d.GetOkExists("use_tke_default"); ok {
		serviceAccountOpts.UseTKEDefault = helper.Bool(v.(bool))
		useTkeDefault = v.(bool)
	}
	if !useTkeDefault {
		if v, ok := d.GetOk("issuer"); ok {
			serviceAccountOpts.Issuer = helper.String(v.(string))
		}
		if v, ok := d.GetOk("jwks_uri"); ok {
			serviceAccountOpts.JWKSURI = helper.String(v.(string))
		}
	}
	if v, ok := d.GetOkExists("auto_create_discovery_anonymous_auth"); ok {
		serviceAccountOpts.AutoCreateDiscoveryAnonymousAuth = helper.Bool(v.(bool))
	}
	request.ServiceAccounts = serviceAccountOpts

	oidcOpts := &tke.OIDCConfigAuthenticationOptions{}
	if v, ok := d.GetOkExists("auto_create_oidc_config"); ok {
		oidcOpts.AutoCreateOIDCConfig = helper.Bool(v.(bool))
	}
	if v, ok := d.GetOk("auto_create_client_id"); ok {
		clientIdSet := v.(*schema.Set).List()
		for i := range clientIdSet {
			oidcOpts.AutoCreateClientId = append(oidcOpts.AutoCreateClientId, helper.String(clientIdSet[i].(string)))
		}
	}
	if v, ok := d.GetOkExists("auto_install_pod_identity_webhook_addon"); ok {
		oidcOpts.AutoInstallPodIdentityWebhookAddon = helper.Bool(v.(bool))
	}
	request.OIDCConfig = oidcOpts

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().ModifyClusterAuthenticationOptions(request)
		if e != nil {
			return retryError(e, tke.RESOURCEUNAVAILABLE_CLUSTERSTATE)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tke kubernetes auth attachment failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(clusterId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	if _, err := service.WaitForAuthenticationOptionsUpdateSuccess(ctx, clusterId); err != nil {
		log.Printf("[CRITAL]%s wait for tke auth attachment create failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTKEAuthAttachmentRead(d, meta)
}

func resourceTencentCloudTKEAuthAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_auth_attachment.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clusterId := d.Id()
	_ = d.Set("cluster_id", clusterId)

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		serviceAccountOpts *tke.ServiceAccountAuthenticationOptions
		oidcConfig         *tke.OIDCConfigAuthenticationOptions
	)

	reqErr := resource.Retry(3*readRetryTimeout, func() *resource.RetryError {
		opts, _, oidc, err := service.DescribeClusterAuthenticationOptions(ctx, clusterId)
		if err != nil {
			return retryError(err)
		}
		serviceAccountOpts = opts
		oidcConfig = oidc
		return nil
	})
	if reqErr != nil {
		log.Printf("[CRITAL]%s read tke kubernetes auth attachment failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	if serviceAccountOpts == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `tencentcloudenterprise_tke_kubernetes_auth_attachment` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	useTkeDefault := serviceAccountOpts.UseTKEDefault != nil && *serviceAccountOpts.UseTKEDefault
	if serviceAccountOpts.UseTKEDefault != nil {
		_ = d.Set("use_tke_default", serviceAccountOpts.UseTKEDefault)
	}

	if useTkeDefault {
		// When use_tke_default=true, issuer/jwks_uri are TKE-managed values, expose as computed fields
		_ = d.Set("tke_default_issuer", serviceAccountOpts.Issuer)
		_ = d.Set("tke_default_jwks_uri", serviceAccountOpts.JWKSURI)
	} else {
		if serviceAccountOpts.Issuer != nil {
			_ = d.Set("issuer", serviceAccountOpts.Issuer)
		}
		if serviceAccountOpts.JWKSURI != nil {
			_ = d.Set("jwks_uri", serviceAccountOpts.JWKSURI)
		}
	}

	// NOTE: AutoCreateDiscoveryAnonymousAuth always returns null by design (API behavior)
	// We do not override the local state for this field

	if oidcConfig != nil {
		if oidcConfig.AutoCreateOIDCConfig != nil {
			_ = d.Set("auto_create_oidc_config", oidcConfig.AutoCreateOIDCConfig)
		}
		if len(oidcConfig.AutoCreateClientId) > 0 {
			clientIds := make([]string, 0, len(oidcConfig.AutoCreateClientId))
			for _, v := range oidcConfig.AutoCreateClientId {
				if v != nil {
					clientIds = append(clientIds, *v)
				}
			}
			_ = d.Set("auto_create_client_id", clientIds)
		}
		if oidcConfig.AutoInstallPodIdentityWebhookAddon != nil {
			_ = d.Set("auto_install_pod_identity_webhook_addon", oidcConfig.AutoInstallPodIdentityWebhookAddon)
		}
	}

	return nil
}

func resourceTencentCloudTKEAuthAttachmentUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_auth_attachment.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clusterId := d.Id()

	mutableArgs := []string{
		"use_tke_default", "issuer", "jwks_uri",
		"auto_create_discovery_anonymous_auth",
		"auto_create_oidc_config", "auto_create_client_id",
		"auto_install_pod_identity_webhook_addon",
	}
	needChange := false
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if !needChange {
		return resourceTencentCloudTKEAuthAttachmentRead(d, meta)
	}

	request := tke.NewModifyClusterAuthenticationOptionsRequest()
	request.ClusterId = helper.String(clusterId)

	serviceAccountOpts := &tke.ServiceAccountAuthenticationOptions{}
	useTkeDefault := false
	if v, ok := d.GetOk("use_tke_default"); ok {
		serviceAccountOpts.UseTKEDefault = helper.Bool(v.(bool))
		useTkeDefault = v.(bool)
	} else {
		serviceAccountOpts.UseTKEDefault = helper.Bool(false)
	}

	if !useTkeDefault {
		if d.HasChange("jwks_uri") {
			serviceAccountOpts.JWKSURI = helper.String(d.Get("jwks_uri").(string))
		}
		if d.HasChange("issuer") {
			serviceAccountOpts.Issuer = helper.String(d.Get("issuer").(string))
		}
	}
	if v, ok := d.GetOkExists("auto_create_discovery_anonymous_auth"); ok {
		serviceAccountOpts.AutoCreateDiscoveryAnonymousAuth = helper.Bool(v.(bool))
	}
	request.ServiceAccounts = serviceAccountOpts

	oidcOpts := &tke.OIDCConfigAuthenticationOptions{}
	if v, ok := d.GetOkExists("auto_create_oidc_config"); ok {
		oidcOpts.AutoCreateOIDCConfig = helper.Bool(v.(bool))
	}
	if v, ok := d.GetOk("auto_create_client_id"); ok {
		clientIdSet := v.(*schema.Set).List()
		for i := range clientIdSet {
			oidcOpts.AutoCreateClientId = append(oidcOpts.AutoCreateClientId, helper.String(clientIdSet[i].(string)))
		}
	}
	if v, ok := d.GetOkExists("auto_install_pod_identity_webhook_addon"); ok {
		oidcOpts.AutoInstallPodIdentityWebhookAddon = helper.Bool(v.(bool))
	}
	request.OIDCConfig = oidcOpts

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().ModifyClusterAuthenticationOptions(request)
		if e != nil {
			return retryError(e, tke.RESOURCEUNAVAILABLE_CLUSTERSTATE)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update tke kubernetes auth attachment failed, reason:%+v", logId, err)
		return err
	}

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	if _, err := service.WaitForAuthenticationOptionsUpdateSuccess(ctx, clusterId); err != nil {
		log.Printf("[CRITAL]%s wait for tke auth attachment update failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTKEAuthAttachmentRead(d, meta)
}

func resourceTencentCloudTKEAuthAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_tke_kubernetes_auth_attachment.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	clusterId := d.Id()

	// Reset to default state: clear jwks_uri and reset issuer to default
	request := tke.NewModifyClusterAuthenticationOptionsRequest()
	request.ClusterId = helper.String(clusterId)
	request.ServiceAccounts = &tke.ServiceAccountAuthenticationOptions{
		JWKSURI: helper.String(""),
		Issuer:  helper.String(DefaultAuthenticationOptionsIssuer),
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseTkeClient().ModifyClusterAuthenticationOptions(request)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete tke kubernetes auth attachment failed, reason:%+v", logId, err)
		return err
	}

	service := TkeService{client: meta.(*TencentCloudClient).apiV3Conn}
	if _, err := service.WaitForAuthenticationOptionsUpdateSuccess(ctx, clusterId); err != nil {
		log.Printf("[CRITAL]%s wait for tke auth attachment delete failed, reason:%+v", logId, err)
		return err
	}

	return nil
}

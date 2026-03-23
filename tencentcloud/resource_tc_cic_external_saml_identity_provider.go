/*
Provides a resource to create an organization cic_external_saml_identity_provider

Example Usage

```hcl
resource "tencentcloudenterprise_cic_external_saml_identity_provider" "cic_external_saml_identity_provider" {
    zone_id = "z-xxxxxx"
    sso_status = "Enabled"
}
```

Import

organization cic_external_saml_identity_provider can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_external_saml_identity_provider.cic_external_saml_identity_provider ${zoneId}
```

 */
package tencentcloud

import (
	"context"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_external_saml_identity_provider", CNDescription{
		TerraformTypeCN: "身份中心外部SAML身份提供商",
		DescriptionCN:   "提供身份中心外部SAML身份提供商资源，用于配置和管理外部SAML身份提供商。",
		AttributesCN: map[string]string{
			"zone_id":                    "空间ID",
			"encoded_metadata_document":  "编码的元数据文档",
			"sso_status":                 "SSO状态",
			"certificate_ids":            "证书ID列表",
			"login_url":                  "登录URL",
			"x509_certificate":           "X509证书",
		},
	})
}

func resourceTencentCloudCicExternalSamlIdentityProvider() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to manage cic external saml identity provider",
		Create: resourceTencentCloudCicExternalSamlIdentityProviderCreate,
		Read:   resourceTencentCloudCicExternalSamlIdentityProviderRead,
		Update: resourceTencentCloudCicExternalSamlIdentityProviderUpdate,
		Delete: resourceTencentCloudCicExternalSamlIdentityProviderDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Space ID.",
			},

			"encoded_metadata_document": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "IdP metadata document (Base64 encoded). Provided by an IdP that supports the SAML 2.0 protocol.",
			},

			"sso_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "SSO enabling status. Valid values: Enabled, Disabled (default).",
			},

			"entity_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "IdP identifier.",
			},

			"login_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "IdP login URL.",
			},

			"x509_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "X509 certificate in PEM format. If this parameter is specified, all existing certificates will be replaced.",
			},
			"acs_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Acs url.",
			},
			"certificate_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Certificate ids.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update time.",
			},
		},
	}
}

func resourceTencentCloudCicExternalSamlIdentityProviderCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_external_saml_identity_provider.create")()
	defer inconsistentCheck(d, meta)()

	var (
		zoneId string
	)
	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}

	d.SetId(zoneId)

	return resourceTencentCloudCicExternalSamlIdentityProviderUpdate(d, meta)
}

func resourceTencentCloudCicExternalSamlIdentityProviderRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_external_saml_identity_provider.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	cicService := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	zoneId := d.Id()

	_ = d.Set("zone_id", zoneId)

	respData, err := cicService.DescribeCicExternalSamlIdentityProviderById(ctx, zoneId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_external_saml_identity_provider` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if respData.EntityId != nil {
		_ = d.Set("entity_id", respData.EntityId)
	}

	if respData.ZoneId != nil {
		_ = d.Set("zone_id", respData.ZoneId)
	}

	if respData.EncodedMetadataDocument != nil {
		_ = d.Set("encoded_metadata_document", respData.EncodedMetadataDocument)
	}

	if respData.AcsUrl != nil {
		_ = d.Set("acs_url", respData.AcsUrl)
	}

	respData1, err := cicService.DescribeCicExternalSamlIdentityProviderConfigById(ctx, zoneId)
	if err != nil {
		return err
	}

	if respData1 == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_external_saml_identity_provider` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}
	if respData1.EntityId != nil {
		_ = d.Set("entity_id", respData1.EntityId)
	}

	if respData1.SSOStatus != nil {
		_ = d.Set("sso_status", respData1.SSOStatus)
	}

	if respData1.EncodedMetadataDocument != nil {
		_ = d.Set("encoded_metadata_document", respData1.EncodedMetadataDocument)
	}

	if respData1.CertificateIds != nil {
		_ = d.Set("certificate_ids", respData1.CertificateIds)
	}

	if respData1.LoginUrl != nil {
		_ = d.Set("login_url", respData1.LoginUrl)
	}

	if respData1.CreateTime != nil {
		_ = d.Set("create_time", respData1.CreateTime)
	}

	if respData1.UpdateTime != nil {
		_ = d.Set("update_time", respData1.UpdateTime)
	}

	_ = zoneId
	return nil
}

func resourceTencentCloudCicExternalSamlIdentityProviderUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_external_saml_identity_provider.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	zoneId := d.Id()

	needChange := false
	mutableArgs := []string{"encoded_metadata_document", "sso_status", "entity_id", "login_url", "x509_certificate"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := cic.NewSetExternalSAMLIdentityProviderRequest()

		if v, ok := d.GetOk("zone_id"); ok {
			request.ZoneId = helper.String(v.(string))
		}

		if v, ok := d.GetOk("encoded_metadata_document"); ok {
			request.EncodedMetadataDocument = helper.String(v.(string))
		}

		if v, ok := d.GetOk("sso_status"); ok {
			request.SSOStatus = helper.String(v.(string))
		}

		if v, ok := d.GetOk("entity_id"); ok {
			request.EntityId = helper.String(v.(string))
		}

		if v, ok := d.GetOk("login_url"); ok {
			request.LoginUrl = helper.String(v.(string))
		}

		if v, ok := d.GetOk("x509_certificate"); ok {
			request.X509Certificate = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().SetExternalSAMLIdentityProvider(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update identity center external saml identity provider failed, reason:%+v", logId, err)
			return err
		}
	}

	_ = zoneId
	return resourceTencentCloudCicExternalSamlIdentityProviderRead(d, meta)
}

func resourceTencentCloudCicExternalSamlIdentityProviderDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_external_saml_identity_provider.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)
	cicService := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	zoneId := d.Id()
	respData1, err := cicService.DescribeCicExternalSamlIdentityProviderConfigById(ctx, zoneId)
	if err != nil {
		return err
	}
	if respData1.SSOStatus != nil && *respData1.SSOStatus == "Enabled" {
		request := cic.NewSetExternalSAMLIdentityProviderRequest()
		request.ZoneId = helper.String(zoneId)
		request.SSOStatus = helper.String("Disabled")
		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().SetExternalSAMLIdentityProvider(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update identity center external saml identity provider failed, reason:%+v", logId, err)
			return err
		}
	}
	var (
		request  = cic.NewClearExternalSAMLIdentityProviderRequest()
		response = cic.NewClearExternalSAMLIdentityProviderResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
	}

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().ClearExternalSAMLIdentityProvider(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete identity center external saml identity provider failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	_ = zoneId
	return nil
}

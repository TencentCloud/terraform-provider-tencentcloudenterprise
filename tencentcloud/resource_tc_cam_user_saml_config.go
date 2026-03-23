/*
Provides a resource to create a CAM user SAML config.

# Example Usage

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

# Import

CAM user SAML config can be imported using the idp_name, e.g.

```
$ terraform import tencentcloudenterprise_cam_user_saml_config.example example-saml-idp
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	open "terraform-provider-tencentcloudenterprise/sdk/open/v20201202"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_user_saml_config", CNDescription{
		TerraformTypeCN: "CAM用户SAML配置",
		DescriptionCN:   "提供 CAM 用户 SAML 配置资源，用于管理用户的 SAML SSO 配置。",
		AttributesCN: map[string]string{
			"idp_name":         "身份提供商名称",
			"protocol":         "协议类型（固定为 saml）",
			"saml_metadata":    "SAML 元数据文档，XML 格式",
			"remark":           "备注",
			"is_sync_idp_user": "是否同步 IdP 用户：0-否，1-是",
			"assist_domain":    "辅助域名",
			"idp_id":           "IdP 配置 ID",
		},
	})
}

func resourceTencentCloudCamUserSamlConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamUserSamlConfigCreate,
		Read:   resourceTencentCloudCamUserSamlConfigRead,
		Update: resourceTencentCloudCamUserSamlConfigUpdate,
		Delete: resourceTencentCloudCamUserSamlConfigDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"idp_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Identity provider name.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Protocol type, fixed value: saml.",
			},
			"saml_metadata": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SAML metadata document, XML format.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remark description.",
			},
			"is_sync_idp_user": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Whether to sync IdP users. 0: no, 1: yes.",
			},
			"assist_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Assist domain for user login.",
			},
			"idp_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IdP configuration ID.",
			},
		},
	}
}

func resourceTencentCloudCamUserSamlConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user_saml_config.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := open.NewAddSamlConfigRequest()
	request.IdpName = helper.String(d.Get("idp_name").(string))
	request.Protocol = helper.String(d.Get("protocol").(string))
	request.SamlMetaData = helper.String(d.Get("saml_metadata").(string))

	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}

	if v, ok := d.GetOk("is_sync_idp_user"); ok {
		request.IsSyncIdpUser = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("assist_domain"); ok {
		request.AssistDomain = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOpenClient().AddSamlConfig(request)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create CAM user SAML config failed, reason:%s\n", logId, err.Error())
		return err
	}

	d.SetId(d.Get("idp_name").(string))
	return resourceTencentCloudCamUserSamlConfigRead(d, meta)
}

func resourceTencentCloudCamUserSamlConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user_saml_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	idpName := d.Id()
	openService := OpenService{client: meta.(*TencentCloudClient).apiV3Conn}

	samlConfig, err := openService.DescribeSamlConfigByName(ctx, idpName)
	if err != nil {
		log.Printf("[CRITAL]%s read CAM user SAML config failed, reason:%s\n", logId, err.Error())
		return err
	}

	if samlConfig == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CamUserSamlConfig` [%s] not found, please check if it has been deleted.\n", logId, idpName)
		return nil
	}

	_ = d.Set("idp_name", samlConfig.IdpName)
	_ = d.Set("protocol", samlConfig.Protocol)
	_ = d.Set("saml_metadata", samlConfig.SamlMetaData)

	if samlConfig.Remark != nil {
		_ = d.Set("remark", samlConfig.Remark)
	}

	if samlConfig.IsSyncIdpUser != nil {
		_ = d.Set("is_sync_idp_user", samlConfig.IsSyncIdpUser)
	}

	if samlConfig.AssistDomain != nil {
		_ = d.Set("assist_domain", samlConfig.AssistDomain)
	}

	if samlConfig.Id != nil {
		_ = d.Set("idp_id", samlConfig.Id)
	}

	return nil
}

func resourceTencentCloudCamUserSamlConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user_saml_config.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	request := open.NewUpdateSamlConfigRequest()

	// Get IdP ID for update
	if v, ok := d.GetOk("idp_id"); ok {
		request.IdpId = helper.IntUint64(v.(int))
	} else {
		return fmt.Errorf("idp_id is required for update operation")
	}

	request.IdpName = helper.String(d.Get("idp_name").(string))
	request.Protocol = helper.String(d.Get("protocol").(string))
	request.SamlMetaData = helper.String(d.Get("saml_metadata").(string))

	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}

	if v, ok := d.GetOk("is_sync_idp_user"); ok {
		request.IsSyncIdpUser = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("assist_domain"); ok {
		request.AssistDomain = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOpenClient().UpdateSamlConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update CAM user SAML config failed, reason:%s\n", logId, err.Error())
		return err
	}

	return resourceTencentCloudCamUserSamlConfigRead(d, meta)
}

func resourceTencentCloudCamUserSamlConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user_saml_config.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	request := open.NewDisabledIdpConfigRequest()

	if v, ok := d.GetOk("idp_id"); ok {
		request.IdpId = helper.IntInt64(v.(int))
	} else {
		return fmt.Errorf("idp_id is required for delete operation")
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOpenClient().DisabledIdpConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s disable cam user saml config failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

// SamlConfigDetail represents the SAML configuration details
type SamlConfigDetail struct {
	Id            *int64
	IdpName       *string
	Protocol      *string
	SamlMetaData  *string
	Remark        *string
	IsSyncIdpUser *int64
	AssistDomain  *string
}

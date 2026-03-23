/*
Provides a resource to create a CAM-OIDC-SSO.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_oidc_sso" "example" {
	  idp_name                = "example-oidc-idp"
	  protocol                = "oidc"
	  identity_url            = "https://login.microsoftonline.com/.../v2.0"
	  identity_key            = "LS0tLS1CRUdJTi..."  # base64 encoded public key
	  client_id               = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	  authorization_endpoint  = "https://login.microsoftonline.com/.../oauth2/v2.0/authorize"
	  response_type           = "id_token"
	  response_mode           = "form_post"
	  scope                   = ["openid", "email", "profile"]

	  # Optional fields
	  remark                  = "OIDC SSO for Azure AD"
	  is_sync_idp_user        = 1
	  email_field             = "email"
	  nick_name_field         = "name"
	  phone_num_field         = "phone"
	  login_account_field     = "preferred_username"
	  logout_url              = "https://login.microsoftonline.com/.../oauth2/v2.0/logout"
	}

```

# Import

CAM-OIDC-SSO can be imported using the idp_name, e.g.

```
$ terraform import tencentcloudenterprise_cam_oidc_sso.example example-oidc-idp
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_oidc_sso", CNDescription{
		TerraformTypeCN: "CAM OIDC SSO配置",
		DescriptionCN:   "提供 CAM OIDC SSO 资源，用于配置 OIDC 身份提供商的单点登录。",
		AttributesCN: map[string]string{
			"idp_name":               "身份提供商名称",
			"protocol":               "协议类型（固定为 oidc）",
			"identity_url":           "身份提供商 URL (issuer)",
			"identity_key":           "签名公钥 (Base64 编码)",
			"client_id":              "客户端 ID",
			"authorization_endpoint": "授权请求端点",
			"response_type":          "响应类型（通常为 id_token）",
			"response_mode":          "响应模式（form_post 或 fragment）",
			"scope":                  "授权范围",
			"remark":                 "备注",
			"is_sync_idp_user":       "是否同步 IdP 用户：0-否，1-是",
			"email_field":            "邮箱字段映射",
			"nick_name_field":        "昵称字段映射",
			"phone_num_field":        "电话字段映射",
			"login_account_field":    "登录账号字段映射",
			"country_code_field":     "国家代码字段映射",
			"logout_url":             "登出 URL",
			"idp_id":                 "IdP 配置 ID",
		},
	})
}

func resourceTencentCloudCamOIDCSSO() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamOIDCSSOCreate,
		Read:   resourceTencentCloudCamOIDCSSORead,
		Update: resourceTencentCloudCamOIDCSSOUpdate,
		Delete: resourceTencentCloudCamOIDCSSODelete,
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
				Description: "Protocol type, fixed value: oidc.",
			},
			"identity_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Identity provider URL (issuer). Corresponds to the value of the `issuer` field in the Openid-configuration provided by the Enterprise IdP.",
			},
			"identity_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The signature public key requires base64_encode. Verify the public key signed by the OpenID Connect identity provider ID Token. For the security of your account, we recommend that you rotate the signed public key regularly.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Client ID, the client ID registered with the OpenID Connect identity provider.",
			},
			"authorization_endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Authorization request Endpoint, OpenID Connect identity provider authorization address. Corresponds to the value of the `authorization_endpoint` field in the Openid-configuration provided by the Enterprise IdP.",
			},
			"response_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Authorization requests The Response type, with a fixed value id_token.",
			},
			"response_mode": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Authorize the request Response mode. Authorization request return mode, form_post and fragment two optional modes, recommended to select form_post mode.",
			},
			"scope": {
				Type:        schema.TypeSet,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Authorize the request Scope. openid; email; profile; Authorization request information scope. The default is required openid.",
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
			"email_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email field mapping in IdP's id_token.",
			},
			"nick_name_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Nickname field mapping in IdP's id_token.",
			},
			"phone_num_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Phone number field mapping in IdP's id_token.",
			},
			"login_account_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Login account field mapping in IdP's id_token.",
			},
			"country_code_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Country code field mapping in IdP's id_token.",
			},
			"logout_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Logout URL.",
			},
			"idp_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IdP configuration ID.",
			},
		},
	}
}

func resourceTencentCloudCamOIDCSSOCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_oidc_sso.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := open.NewAddOidcConfigRequest()
	request.IdpName = helper.String(d.Get("idp_name").(string))
	request.Protocol = helper.String(d.Get("protocol").(string))
	request.IdentityUrl = helper.String(d.Get("identity_url").(string))
	request.IdentityKey = helper.String(d.Get("identity_key").(string))
	request.ClientId = helper.String(d.Get("client_id").(string))
	request.AuthorizationEndpoint = helper.String(d.Get("authorization_endpoint").(string))
	request.ResponseType = helper.String(d.Get("response_type").(string))
	request.ResponseMode = helper.String(d.Get("response_mode").(string))

	if v, ok := d.GetOk("scope"); ok {
		request.Scope = helper.InterfacesStringsPoint(v.(*schema.Set).List())
	} else {
		request.Scope = helper.InterfacesStringsPoint([]interface{}{"openid"})
	}

	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}

	if v, ok := d.GetOk("is_sync_idp_user"); ok {
		request.IsSyncIdpUser = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("email_field"); ok {
		request.EmailField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("nick_name_field"); ok {
		request.NickNameField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("phone_num_field"); ok {
		request.PhoneNumField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("login_account_field"); ok {
		request.LoginAccountField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("country_code_field"); ok {
		request.CountryCodeField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("logout_url"); ok {
		request.LogoutUrl = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOpenClient().AddOidcConfig(request)
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
		log.Printf("[CRITAL]%s create CAM OIDC SSO failed, reason:%s\n", logId, err.Error())
		return err
	}

	d.SetId(d.Get("idp_name").(string))
	return resourceTencentCloudCamOIDCSSORead(d, meta)
}

func resourceTencentCloudCamOIDCSSORead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_oidc_sso.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	idpName := d.Id()
	openService := OpenService{client: meta.(*TencentCloudClient).apiV3Conn}

	oidcConfig, err := openService.DescribeOidcConfigByName(ctx, idpName)
	if err != nil {
		log.Printf("[CRITAL]%s read CAM OIDC SSO failed, reason:%s\n", logId, err.Error())
		return err
	}

	if oidcConfig == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `CamOIDCSSO` [%s] not found, please check if it has been deleted.\n", logId, idpName)
		return nil
	}

	_ = d.Set("idp_name", oidcConfig.IdpName)
	_ = d.Set("protocol", oidcConfig.Protocol)
	_ = d.Set("identity_url", oidcConfig.IdentityUrl)
	_ = d.Set("identity_key", oidcConfig.IdentityKey)
	_ = d.Set("client_id", oidcConfig.ClientId)
	_ = d.Set("authorization_endpoint", oidcConfig.AuthorizationEndpoint)
	_ = d.Set("response_type", oidcConfig.ResponseType)
	_ = d.Set("response_mode", oidcConfig.ResponseMode)

	if oidcConfig.Scope != nil && len(oidcConfig.Scope) > 0 {
		_ = d.Set("scope", oidcConfig.Scope)
	}

	if oidcConfig.Remark != nil {
		_ = d.Set("remark", oidcConfig.Remark)
	}

	if oidcConfig.IsSyncIdpUser != nil {
		_ = d.Set("is_sync_idp_user", oidcConfig.IsSyncIdpUser)
	}

	if oidcConfig.EmailField != nil {
		_ = d.Set("email_field", oidcConfig.EmailField)
	}

	if oidcConfig.NickNameField != nil {
		_ = d.Set("nick_name_field", oidcConfig.NickNameField)
	}

	if oidcConfig.PhoneNumField != nil {
		_ = d.Set("phone_num_field", oidcConfig.PhoneNumField)
	}

	if oidcConfig.LoginAccountField != nil {
		_ = d.Set("login_account_field", oidcConfig.LoginAccountField)
	}

	if oidcConfig.CountryCodeField != nil {
		_ = d.Set("country_code_field", oidcConfig.CountryCodeField)
	}

	if oidcConfig.LogoutUrl != nil {
		_ = d.Set("logout_url", oidcConfig.LogoutUrl)
	}

	if oidcConfig.Id != nil {
		_ = d.Set("idp_id", oidcConfig.Id)
	}

	return nil
}

func resourceTencentCloudCamOIDCSSOUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_oidc_sso.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	request := open.NewUpdateOidcConfigRequest()

	// Get IdP ID for update
	if v, ok := d.GetOk("idp_id"); ok {
		request.Id = helper.IntInt64(v.(int))
	} else {
		return fmt.Errorf("idp_id is required for update operation")
	}

	request.IdpName = helper.String(d.Get("idp_name").(string))
	request.Protocol = helper.String(d.Get("protocol").(string))
	request.IdentityUrl = helper.String(d.Get("identity_url").(string))
	request.IdentityKey = helper.String(d.Get("identity_key").(string))
	request.ClientId = helper.String(d.Get("client_id").(string))
	request.AuthorizationEndpoint = helper.String(d.Get("authorization_endpoint").(string))
	request.ResponseType = helper.String(d.Get("response_type").(string))
	request.ResponseMode = helper.String(d.Get("response_mode").(string))

	if v, ok := d.GetOk("scope"); ok {
		request.Scope = helper.InterfacesStringsPoint(v.(*schema.Set).List())
	} else {
		request.Scope = helper.InterfacesStringsPoint([]interface{}{"openid"})
	}

	if v, ok := d.GetOk("remark"); ok {
		request.Remark = helper.String(v.(string))
	}

	if v, ok := d.GetOk("is_sync_idp_user"); ok {
		request.IsSyncIdpUser = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("email_field"); ok {
		request.EmailField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("nick_name_field"); ok {
		request.NickNameField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("phone_num_field"); ok {
		request.PhoneNumField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("login_account_field"); ok {
		request.LoginAccountField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("country_code_field"); ok {
		request.CountryCodeField = helper.String(v.(string))
	}

	if v, ok := d.GetOk("logout_url"); ok {
		request.LogoutUrl = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseOpenClient().UpdateOidcConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update CAM OIDC SSO failed, reason:%s\n", logId, err.Error())
		return err
	}

	return resourceTencentCloudCamOIDCSSORead(d, meta)
}

func resourceTencentCloudCamOIDCSSODelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_oidc_sso.delete")()
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
		log.Printf("[CRITAL]%s disable cam oidc sso failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

// OidcConfigDetail represents the OIDC configuration details
type OidcConfigDetail struct {
	Id                    *int64
	IdpName               *string
	Protocol              *string
	IdentityUrl           *string
	IdentityKey           *string
	ClientId              *string
	AuthorizationEndpoint *string
	ResponseType          *string
	ResponseMode          *string
	Scope                 []*string
	Remark                *string
	IsSyncIdpUser         *int64
	EmailField            *string
	NickNameField         *string
	PhoneNumField         *string
	LoginAccountField     *string
	CountryCodeField      *string
	LogoutUrl             *string
}

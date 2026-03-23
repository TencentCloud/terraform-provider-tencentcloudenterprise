/*
Provides a resource to manage CAM MFA (Multi-Factor Authentication) protection flags.

This resource allows you to enable or disable MFA protection for login and sensitive operations.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_mfa_flag" "example" {
	  op_uin = 100000000001

	  login_flag {
	    phone  = 1  # Enable phone verification for login
	    stoken = 1  # Enable soft token for login
	    token  = 0  # Disable hard token for login
	    ukey   = 0  # Disable Ukey for login
	  }

	  action_flag {
	    phone  = 1  # Enable phone verification for sensitive operations
	    stoken = 0  # Disable soft token for sensitive operations
	    token  = 0  # Disable hard token for sensitive operations
	    ukey   = 0  # Disable Ukey for sensitive operations
	  }
	}

```

# Import

CAM MFA flag can be imported using the uin, e.g.

```
$ terraform import tencentcloudenterprise_cam_mfa_flag.example 100000000001
```
*/
package tencentcloud

import (
	"context"
	"log"
	"strconv"

	account "terraform-provider-tencentcloudenterprise/sdk/account/v20190325"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_mfa_flag", CNDescription{
		TerraformTypeCN: "CAM多因素认证保护",
		DescriptionCN:   "提供 CAM 多因素认证保护资源，用于管理用户的登录保护和操作保护设置。",
		AttributesCN: map[string]string{
			"op_uin":      "操作的用户 UIN",
			"login_flag":  "登录保护设置",
			"action_flag": "操作保护设置",
			"phone":       "手机验证：0-关闭，1-开启",
			"stoken":      "软令牌验证：0-关闭，1-开启",
			"token":       "硬件令牌验证：0-关闭，1-开启",
			"ukey":        "U盾验证：0-关闭，1-开启",
		},
	})
}

func resourceTencentCloudCamMfaFlag() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamMfaFlagCreate,
		Read:   resourceTencentCloudCamMfaFlagRead,
		Update: resourceTencentCloudCamMfaFlagUpdate,
		Delete: resourceTencentCloudCamMfaFlagDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"op_uin": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "The UIN of the user to operate on.",
			},

			"login_flag": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "MFA protection settings for login. Each field accepts 0 (disabled) or 1 (enabled).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phone": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Phone verification for login. 0: disabled, 1: enabled.",
						},
						"stoken": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Soft token verification for login. 0: disabled, 1: enabled.",
						},
						"token": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Hardware token verification for login. 0: disabled, 1: enabled.",
						},
						"ukey": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Ukey verification for login. 0: disabled, 1: enabled.",
						},
					},
				},
			},

			"action_flag": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "MFA protection settings for sensitive operations. Each field accepts 0 (disabled) or 1 (enabled).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"phone": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Phone verification for sensitive operations. 0: disabled, 1: enabled.",
						},
						"stoken": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Soft token verification for sensitive operations. 0: disabled, 1: enabled.",
						},
						"token": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Hardware token verification for sensitive operations. 0: disabled, 1: enabled.",
						},
						"ukey": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Ukey verification for sensitive operations. 0: disabled, 1: enabled.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudCamMfaFlagCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_mfa_flag.create")()
	defer inconsistentCheck(d, meta)()
	var opUin int

	if v, ok := d.GetOk("op_uin"); ok {
		opUin = v.(int)
	}
	d.SetId(strconv.Itoa(opUin))
	return resourceTencentCloudCamMfaFlagUpdate(d, meta)
}

func resourceTencentCloudCamMfaFlagRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_mfa_flag.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	opUin := d.Id()
	uin, err := strconv.Atoi(opUin)
	if err != nil {
		return err
	}
	service := &CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	loginFlag, actionFlag, err := service.DescribeCamMfaFlagById(ctx, uint64(uin))
	if err != nil {
		return err
	}

	if loginFlag == nil && actionFlag == nil {
		log.Printf("[WARN]%s resource `CamMfaFlag` not found, please check if it has been deleted.\n", logId)
		return nil
	}

	_ = d.Set("op_uin", uin)

	if loginFlag != nil {
		loginFlagMap := map[string]interface{}{}

		if loginFlag.Phone != nil {
			loginFlagMap["phone"] = loginFlag.Phone
		}

		if loginFlag.Stoken != nil {
			loginFlagMap["stoken"] = loginFlag.Stoken
		}

		if loginFlag.Token != nil {
			loginFlagMap["token"] = loginFlag.Token
		}

		if loginFlag.Ukey != nil {
			loginFlagMap["ukey"] = loginFlag.Ukey
		}

		_ = d.Set("login_flag", []interface{}{loginFlagMap})
	}

	if actionFlag != nil {
		actionFlagMap := map[string]interface{}{}

		if actionFlag.Phone != nil {
			actionFlagMap["phone"] = actionFlag.Phone
		}

		if actionFlag.Stoken != nil {
			actionFlagMap["stoken"] = actionFlag.Stoken
		}

		if actionFlag.Token != nil {
			actionFlagMap["token"] = actionFlag.Token
		}

		if actionFlag.Ukey != nil {
			actionFlagMap["ukey"] = actionFlag.Ukey
		}

		_ = d.Set("action_flag", []interface{}{actionFlagMap})
	}

	return nil
}

func resourceTencentCloudCamMfaFlagUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_mfa_flag.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	opUin := d.Id()
	request := account.NewSetSafeAuthFlagRequest()
	uin, err := strconv.Atoi(opUin)
	if err != nil {
		return err
	}
	request.UserUin = helper.String(strconv.Itoa(uin))

	if d.HasChange("login_flag") {
		if dMap, ok := helper.InterfacesHeadMap(d, "login_flag"); ok {
			safeAuthFlag := account.SafeAuthFlag{}
			if v, ok := dMap["phone"]; ok {
				safeAuthFlag.Phone = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["stoken"]; ok {
				safeAuthFlag.Stoken = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["token"]; ok {
				safeAuthFlag.Token = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["ukey"]; ok {
				safeAuthFlag.Ukey = helper.IntInt64(v.(int))
			}
			request.LoginFlag = &safeAuthFlag
		}
	}

	if d.HasChange("action_flag") {
		if dMap, ok := helper.InterfacesHeadMap(d, "action_flag"); ok {
			safeAuthFlag := account.SafeAuthFlag{}
			if v, ok := dMap["phone"]; ok {
				safeAuthFlag.Phone = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["stoken"]; ok {
				safeAuthFlag.Stoken = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["token"]; ok {
				safeAuthFlag.Token = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["ukey"]; ok {
				safeAuthFlag.Ukey = helper.IntInt64(v.(int))
			}
			request.ActionFlag = &safeAuthFlag
		}
	}

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseAccountClient().SetSafeAuthFlag(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update cam mfaFlag failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudCamMfaFlagRead(d, meta)
}

func resourceTencentCloudCamMfaFlagDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_mfa_flag.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

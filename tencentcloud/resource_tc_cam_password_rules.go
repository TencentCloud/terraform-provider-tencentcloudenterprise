/*
Provides a resource to manage CAM password rules.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_password_rules" "foo" {
	  minimum_length               = 12
	  must_contain                 = "Aa1!"
	  force_password_change        = 90
	  reuse_password_limit         = 3
	  retry_password_limit         = 5
	  only_admin_can_reset_password = false
	  must_not_contain_username    = true
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_password_rules", CNDescription{
		TerraformTypeCN: "CAM密码规则",
		DescriptionCN:   "提供 CAM 密码规则资源，用于管理子用户的密码强度和有效期等规则。",
		AttributesCN: map[string]string{
			"minimum_length":                "密码最小长度",
			"must_contain":                  "密码必须包含的字符类型 (如: Aa)",
			"force_password_change":         "密码强制更换天数",
			"reuse_password_limit":          "密码历史重用限制次数",
			"retry_password_limit":          "密码错误锁定前的重试次数",
			"only_admin_can_reset_password": "是否仅管理员可重置密码",
			"must_not_contain_username":     "密码是否不能包含用户名",
			"black_list":                    "密码关键字黑名单 (JSON 数组字符串)",
		},
	})
}

func resourceTencentCloudCamPasswordRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamPasswordRulesCreate,
		Read:   resourceTencentCloudCamPasswordRulesRead,
		Update: resourceTencentCloudCamPasswordRulesUpdate,
		Delete: resourceTencentCloudCamPasswordRulesDelete,

		Schema: map[string]*schema.Schema{
			"minimum_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Minimum password length.",
			},
			"must_contain": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Character classes that must be contained in the password (for example: Aa).",
			},
			"force_password_change": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Days after which a password must be changed.",
			},
			"reuse_password_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Number of previous passwords that cannot be reused.",
			},
			"retry_password_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Maximum retry attempts before lockout.",
			},
			"only_admin_can_reset_password": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether only administrators can reset passwords.",
			},
			"must_not_contain_username": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether passwords must not contain the username.",
			},
			"black_list": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "JSON array string of disallowed password keywords.",
			},
		},
	}
}

func resourceTencentCloudCamPasswordRulesCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_password_rules.create")()
	d.SetId("cam-password-rules")

	if err := resourceTencentCloudCamPasswordRulesApply(d, meta); err != nil {
		return err
	}
	return resourceTencentCloudCamPasswordRulesRead(d, meta)
}

func resourceTencentCloudCamPasswordRulesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_password_rules.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{client: meta.(*TencentCloudClient).apiV3Conn}
	rules, blacklist, err := camService.DescribePasswordRules(ctx)
	if err != nil {
		return err
	}

	if rules == nil {
		d.SetId("")
		return nil
	}

	if rules.MinimumLength != nil {
		_ = d.Set("minimum_length", int(*rules.MinimumLength))
	} else {
		_ = d.Set("minimum_length", nil)
	}
	if rules.MustContain != nil {
		_ = d.Set("must_contain", *rules.MustContain)
	} else {
		_ = d.Set("must_contain", nil)
	}
	if rules.ForcePasswordChange != nil {
		_ = d.Set("force_password_change", int(*rules.ForcePasswordChange))
	} else {
		_ = d.Set("force_password_change", nil)
	}
	if rules.ReusePasswordLimit != nil {
		_ = d.Set("reuse_password_limit", int(*rules.ReusePasswordLimit))
	} else {
		_ = d.Set("reuse_password_limit", nil)
	}
	if rules.RetryPasswordLimit != nil {
		_ = d.Set("retry_password_limit", int(*rules.RetryPasswordLimit))
	} else {
		_ = d.Set("retry_password_limit", nil)
	}
	if rules.OnlyAdminCanResetPassword != nil {
		_ = d.Set("only_admin_can_reset_password", *rules.OnlyAdminCanResetPassword == 1)
	}
	if rules.MustNotContainUsername != nil {
		_ = d.Set("must_not_contain_username", *rules.MustNotContainUsername == 1)
	}
	if blacklist != nil {
		_ = d.Set("black_list", *blacklist)
	} else {
		_ = d.Set("black_list", nil)
	}

	return nil
}

func resourceTencentCloudCamPasswordRulesUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_password_rules.update")()
	defer inconsistentCheck(d, meta)()

	if err := resourceTencentCloudCamPasswordRulesApply(d, meta); err != nil {
		return err
	}
	return resourceTencentCloudCamPasswordRulesRead(d, meta)
}

func resourceTencentCloudCamPasswordRulesDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_password_rules.delete")()
	defer inconsistentCheck(d, meta)()
	// No delete API; keep state only.
	return nil
}

func resourceTencentCloudCamPasswordRulesApply(d *schema.ResourceData, meta interface{}) error {
	logId := getLogId(contextNil)

	rules, blacklist := expandCamPasswordRules(d)
	if rules == nil && blacklist == nil {
		return nil
	}

	request := cam.NewUpdatePasswordRulesRequest()
	request.Rules = rules
	request.BlackList = blacklist

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().UpdatePasswordRules(request)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n", logId, request.GetAction(), request.ToJsonString(), e.Error())
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update CAM password rules failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

func expandCamPasswordRules(d *schema.ResourceData) (rules *cam.PasswordRules, blacklist *string) {
	pr := &cam.PasswordRules{}
	hasRules := false

	if v, ok := d.GetOk("minimum_length"); ok {
		pr.MinimumLength = helper.Int64(int64(v.(int)))
		hasRules = true
	}
	if v, ok := d.GetOk("must_contain"); ok {
		pr.MustContain = helper.String(v.(string))
		hasRules = true
	}
	if v, ok := d.GetOk("force_password_change"); ok {
		pr.ForcePasswordChange = helper.Int64(int64(v.(int)))
		hasRules = true
	}
	if v, ok := d.GetOk("reuse_password_limit"); ok {
		pr.ReusePasswordLimit = helper.Int64(int64(v.(int)))
		hasRules = true
	}
	if v, ok := d.GetOk("retry_password_limit"); ok {
		pr.RetryPasswordLimit = helper.Int64(int64(v.(int)))
		hasRules = true
	}
	if v, ok := d.GetOkExists("only_admin_can_reset_password"); ok {
		val := v.(bool)
		if val {
			pr.OnlyAdminCanResetPassword = helper.Int64(1)
		} else {
			pr.OnlyAdminCanResetPassword = helper.Int64(0)
		}
		hasRules = true
	}
	if v, ok := d.GetOkExists("must_not_contain_username"); ok {
		val := v.(bool)
		if val {
			pr.MustNotContainUsername = helper.Int64(1)
		} else {
			pr.MustNotContainUsername = helper.Int64(0)
		}
		hasRules = true
	}

	if hasRules {
		rules = pr
	}

	if v, ok := d.GetOk("black_list"); ok {
		blacklist = helper.String(v.(string))
	}

	return
}

/*
Provides a resource to manage CAM user.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_user" "foo" {
	  name                = "cam-user-test"
	  remark              = "test"
	  console_login       = true
	  use_api             = true
	  need_reset_password = true
	  password            = "Gail@1234"
	  phone_num           = "12345678910"
	  email               = "hello@test.com"
	  country_code        = "86"
	  force_delete        = true
	  tags = {
	    test  = "tf-cam-user",
	  }
	}

```

# Import

CAM user can be imported using the user name, e.g.

```
$ terraform import tencentcloudenterprise_cam_user.foo cam-user-test
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"log"
	"strings"
	"time"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_user", CNDescription{
		TerraformTypeCN: "CAM用户",
		DescriptionCN:   "提供 CAM 用户资源，用于创建和管理子用户。",
		AttributesCN: map[string]string{
			"name":                "子账号名称",
			"remark":              "备注",
			"force_delete":        "是否强制删除",
			"use_api":             "是否生成 API 密钥",
			"console_login":       "是否允许登录控制台",
			"password":            "密码",
			"need_reset_password": "登录是否需要重置密码",
			"phone_num":           "手机号码",
			"country_code":        "国家代码",
			"email":               "邮箱",
			"uin":                 "帐号 UIN",
			"uid":                 "子用户 UID",
			"secret_id":           "密钥 ID",
			"secret_key":          "密钥 Key",
			"tags":                "标签",
		},
	})
}

func resourceTencentCloudCamUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCamUserCreate,
		Read:   resourceTencentCloudCamUserRead,
		Update: resourceTencentCloudCamUserUpdate,
		Delete: resourceTencentCloudCamUserDelete,
		Importer: &schema.ResourceImporter{
			State: helper.ImportWithDefaultValue(map[string]interface{}{
				"remark":              "",
				"force_delete":        false,
				"use_api":             true,
				"console_login":       false,
				"need_reset_password": true,
			}),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "Name of the CAM user.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Remark of the CAM user.",
			},
			"force_delete": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicate whether to force deletes the CAM user. If set false, the API secret key will be checked and failed when exists; otherwise the user will be deleted directly. Default is false.",
			},
			"use_api": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicate whether to generate the API secret key or not.",
			},
			"console_login": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicate whether the CAM user can login to the web console or not.",
			},
			"password": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Sensitive:    true,
				ValidateFunc: validateAsConfigPassword,
				Description:  "The password of the CAM user. Password should be at least 8 characters and no more than 32 characters, includes uppercase letters, lowercase letters, numbers and special characters. Only required when `console_login` is true. If not set, a random password will be automatically generated.",
			},
			"need_reset_password": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Indicate whether the CAM user need to reset the password when first logins.",
			},
			"phone_num": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Phone number of the CAM user.",
			},
			"country_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Country code of the phone number, for example: '86'.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email of the CAM user.",
			},
			"uin": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Uin of the CAM User.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Secret key of the CAM user.",
			},
			"secret_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Secret ID of the CAM user.",
			},
			"uid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ID of the CAM user.",
			},
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "A list of tags used to associate different resources.",
			},
		},
	}
}

func resourceTencentCloudCamUserCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user.create")()

	logId := getLogId(contextNil)

	boolToStr := func(b bool) string {
		if b {
			return "1"
		}
		return "0"
	}

	request := cam.NewAddSubAccountRequest()

	user := &cam.UserInfo{
		Name:       helper.String(d.Get("name").(string)),
		Remark:     helper.String(d.Get("remark").(string)),
		SystemType: helper.String("SubAccount"),
	}
	// Phone & email
	if v, ok := d.GetOk("phone_num"); ok {
		user.PhoneNum = helper.String(v.(string))
	}
	if v, ok := d.GetOk("country_code"); ok {
		user.CountryCode = helper.String(v.(string))
	}
	if v, ok := d.GetOk("email"); ok {
		user.Email = helper.String(v.(string))
	}

	// Detail flags
	detail := &cam.AccountDetail{}
	detail.UseApi = helper.String(boolToStr(d.Get("use_api").(bool)))
	detail.ConsoleLogin = helper.String(boolToStr(d.Get("console_login").(bool)))
	detail.NeedResetPassword = helper.String(boolToStr(d.Get("need_reset_password").(bool)))
	if v, ok := d.GetOk("password"); ok {
		detail.Password = helper.String(v.(string))
	}
	user.Detail = detail
	// CanLogin align with console_login
	user.CanLogin = helper.String(boolToStr(d.Get("console_login").(bool)))

	request.UserInfo = []*cam.UserInfo{user}
	request.FromAPI = helper.Int64(1)
	request.Lang = helper.String("zh-CN")

	var response *cam.AddSubAccountResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().AddSubAccount(request)
		if e != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), e.Error())
			if ee, ok := e.(*sdkErrors.CloudSDKError); ok {
				errCode := ee.GetCode()
				if strings.Contains(errCode, "SubUserNameInUse") {
					return resource.NonRetryableError(e)
				}
			}
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create user failed, reason:%s\n", logId, err.Error())
		return err
	}
	if response.Response == nil || len(response.Response.SubAccounts) == 0 || response.Response.SubAccounts[0].Name == nil {
		return fmt.Errorf("[CHECK][CAM sub account][Create] response missing SubAccounts")
	}

	sub := response.Response.SubAccounts[0]
	d.SetId(*sub.Name)
	if sub.SecretKey != nil {
		_ = d.Set("secret_key", *sub.SecretKey)
	}
	if sub.Password != nil {
		_ = d.Set("password", *sub.Password)
	}
	if sub.SecretId != nil {
		_ = d.Set("secret_id", *sub.SecretId)
	}

	//get real instance then read
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		instance, e := camService.DescribeUserById(ctx, *sub.Name)
		if e != nil {
			return retryError(e)
		}
		if instance == nil {
			return resource.RetryableError(fmt.Errorf("creation not done"))
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s wait for CAM user ready failed, reason:%s\n", logId, err.Error())
		return err
	}

	//modify tags
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		// prefer Uin from response if available
		if sub.Uin != nil {
			resourceName := BuildTagResourceName("cam", "uin", region, helper.UInt64ToStr(*sub.Uin))
			if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
				return err
			}
		}
	}
	time.Sleep(10 * time.Second)
	return resourceTencentCloudCamUserRead(d, meta)
}

func resourceTencentCloudCamUserRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	deleteForce := false
	if v, ok := d.GetOkExists("force_delete"); ok {
		deleteForce = v.(bool)
		_ = d.Set("force_delete", deleteForce)
	}

	userId := d.Id()
	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var instance *cam.SubAccountFilter
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribeUserById(ctx, userId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM user failed, reason:%s\n", logId, err.Error())
		return err
	}

	if instance == nil || instance.Uid == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("name", userId)
	_ = d.Set("uin", int(*instance.Uin))
	_ = d.Set("uid", int(*instance.Uid))
	_ = d.Set("remark", instance.Remark)
	_ = d.Set("phone_num", instance.PhoneNum)
	_ = d.Set("country_code", instance.CountryCode)
	_ = d.Set("email", instance.Email)
	if instance.ConsoleLogin != nil && int(*instance.ConsoleLogin) == 0 {
		_ = d.Set("console_login", false)
	} else if instance.ConsoleLogin != nil && int(*instance.ConsoleLogin) == 1 {
		_ = d.Set("console_login", true)
	}

	//tags
	tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
	region := meta.(*TencentCloudClient).apiV3Conn.Region
	tags, err := tagService.DescribeResourceTags(ctx, "cam", "uin", region, helper.UInt64ToStr(*instance.Uin))
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudCamUserUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	userId := d.Id()

	boolToStr := func(b bool) string {
		if b {
			return "1"
		}
		return "0"
	}

	var updateAttrs []string

	info := &cam.SubAccountInfo{
		Name:       &userId,
		SystemType: helper.String("SubAccount"),
	}

	if d.HasChange("remark") {
		info.Remark = helper.String(d.Get("remark").(string))
		updateAttrs = append(updateAttrs, "remark")
	}

	if d.HasChange("console_login") {
		info.ConsoleLogin = helper.String(boolToStr(d.Get("console_login").(bool)))
		info.CanLogin = helper.String(boolToStr(d.Get("console_login").(bool)))
		updateAttrs = append(updateAttrs, "console_login")
	}

	if d.HasChange("password") {
		password := d.Get("password").(string)
		info.Password = helper.String(password)
		updateAttrs = append(updateAttrs, "password")
	}

	if d.HasChange("need_reset_password") {
		info.NeedResetPassword = helper.String(boolToStr(d.Get("need_reset_password").(bool)))
		updateAttrs = append(updateAttrs, "need_reset_password")
	}

	if d.HasChange("phone_num") || d.HasChange("country_code") {
		info.PhoneNum = helper.String(d.Get("phone_num").(string))
		updateAttrs = append(updateAttrs, "phone_num")
		info.CountryCode = helper.String(d.Get("country_code").(string))
		updateAttrs = append(updateAttrs, "country_code")
	}
	if d.HasChange("email") {
		info.Email = helper.String(d.Get("email").(string))
		updateAttrs = append(updateAttrs, "email")
	}

	if len(updateAttrs) > 0 {
		request := cam.NewUpdateSubAccountRequest()
		// Use UpdateType=4 as per API examples for profile updates
		updateType := int64(4)
		request.UpdateType = &updateType
		// Keep in whitelist to allow updates
		request.InWhiteList = helper.Int64(1)
		request.UserInfo = []*cam.SubAccountInfo{info}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			response, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().UpdateSubAccount(request)

			if e != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, request.GetAction(), request.ToJsonString(), e.Error())
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update CAM sub account failed, reason:%s\n", logId, err.Error())
			return err
		}
	}

	//tag
	if d.HasChange("tags") {
		camService := CamService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}

		var instance *cam.SubAccountFilter
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := camService.DescribeUserById(ctx, userId)
			if e != nil {
				return retryError(e)
			}
			instance = result
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s read CAM user failed, reason:%s\n", logId, err.Error())
			return err
		}

		if instance == nil || instance.Uid == nil {
			d.SetId("")
			return nil
		}

		oldInterface, newInterface := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldInterface.(map[string]interface{}), newInterface.(map[string]interface{}))
		tagService := TagService{
			client: meta.(*TencentCloudClient).apiV3Conn,
		}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := BuildTagResourceName("cam", "uin", region, helper.UInt64ToStr(*instance.Uin))
		err = tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags)
		if err != nil {
			return err
		}

	}

	return nil
}

func resourceTencentCloudCamUserDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_user.delete")()

	logId := getLogId(contextNil)

	userId := d.Id()
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	camService := CamService{client: meta.(*TencentCloudClient).apiV3Conn}
	var instance *cam.SubAccountFilter
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := camService.DescribeUserById(ctx, userId)
		if e != nil {
			return retryError(e)
		}
		instance = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM sub account failed before delete, reason:%s\n", logId, err.Error())
		return err
	}
	if instance == nil || instance.Uid == nil || instance.Uin == nil {
		// treat as already gone
		d.SetId("")
		return nil
	}

	uinfo := &cam.GroupUidUinInfo{
		Uid:     instance.Uid,
		Uin:     instance.Uin,
		GroupId: helper.Int64(-1),
	}

	request := cam.NewDeleteSubAccountRequest()
	request.UserInfo = []*cam.GroupUidUinInfo{uinfo}
	request.FromAPI = helper.Uint64(1)
	request.Lang = helper.String("zh-CN")

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		_, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().DeleteSubAccount(request)
		if e != nil {
			log.Printf("[CRITAL]%s reason[%s]\n", logId, e.Error())
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete CAM user failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

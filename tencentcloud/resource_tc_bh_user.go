/*
Provide a resource to create a BH user

# Example Usage

```hcl

	resource "tencentcloudenterprise_bh_user" "example" {
	  user_name = "example_user"
	  real_name = "Example User"
	  email     = "example@example.com"
	}

```

# Import

BH user can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_user.example 12345
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_bh_user", CNDescription{
		TerraformTypeCN: "BH 用户",
		DescriptionCN:   "提供 BH 用户资源，用于创建和管理堡垒机用户。",
		AttributesCN: map[string]string{
			"user_name":     "用户名",
			"real_name":     "用户姓名",
			"phone":         "手机号",
			"email":         "邮箱",
			"validate_from": "用户生效时间",
			"validate_to":   "用户失效时间",
			"group_id_set":  "用户所属用户组ID集合",
			"auth_type":     "认证方式",
			"validate_time": "访问时间限制",
			"department_id": "用户所属部门ID",
			"user_id":       "用户ID",
		},
	})
}

func ResourceTencentCloudBhUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudBhUserCreate,
		Read:   resourceTencentCloudBhUserRead,
		Update: resourceTencentCloudBhUserUpdate,
		Delete: resourceTencentCloudBhUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Username, 3-20 characters, must start with an English letter and cannot contain characters other than `letters`, `numbers`, `.`, `_`, `-`.",
			},

			"real_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User's real name, maximum length 20 characters, cannot contain whitespace characters.",
			},

			"phone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Input in the format of \"country code|phone number\", e.g.: \"+86|xxxxxxxx\". At least one of phone and email parameters must be provided.",
			},

			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email address. At least one of phone and email parameters must be provided.",
			},

			"validate_from": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "User effective time, e.g.: \"2021-09-22T00:00:00+00:00\". If effective and expiration times are not filled, the user will be valid permanently.",
			},

			"validate_to": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "User expiration time, e.g.: \"2021-09-23T00:00:00+00:00\". If effective and expiration times are not filled, the user will be valid permanently.",
			},

			"group_id_set": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "User group ID set to which the user belongs.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			"auth_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Authentication method, 0 - local, 1 - LDAP, 2 - OAuth. Default is 0 if not provided.",
			},

			"validate_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Access time restriction, a string composed of 0 and 1 with length 168 (7 * 24), representing the time slots allowed for the user in a week. The Nth character in the string represents the Nth hour in the week, 0 - not allowed to access, 1 - allowed to access.",
			},

			"department_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Department ID to which the user belongs, e.g.: \"1.2.3\".",
			},

			// computed
			"user_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "User ID.",
			},
		},
	}
}

func resourceTencentCloudBhUserCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = bhsaas.NewCreateUserRequest()
		response = bhsaas.NewCreateUserResponse()
		userId   string
	)

	if v, ok := d.GetOk("user_name"); ok {
		request.UserName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("real_name"); ok {
		request.RealName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("phone"); ok {
		request.Phone = helper.String(v.(string))
	}

	if v, ok := d.GetOk("email"); ok {
		request.Email = helper.String(v.(string))
	}

	if v, ok := d.GetOk("validate_from"); ok {
		request.ValidateFrom = helper.String(v.(string))
	}

	if v, ok := d.GetOk("validate_to"); ok {
		request.ValidateTo = helper.String(v.(string))
	}

	if v, ok := d.GetOk("group_id_set"); ok {
		groupIdSetSet := v.(*schema.Set).List()
		for i := range groupIdSetSet {
			groupIdSet := groupIdSetSet[i].(int)
			request.GroupIdSet = append(request.GroupIdSet, helper.IntUint64(groupIdSet))
		}
	}

	if v, ok := d.GetOkExists("auth_type"); ok {
		request.AuthType = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("validate_time"); ok {
		request.ValidateTime = helper.String(v.(string))
	}

	if v, ok := d.GetOk("department_id"); ok {
		request.DepartmentId = helper.String(v.(string))
	}

	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateUser(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("Create bh user failed, Response is nil."))
		}

		response = result
		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s create bh user failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	if response.Response.Id == nil {
		return fmt.Errorf("Id is nil.")
	}

	userId = helper.UInt64ToStr(*response.Response.Id)
	d.SetId(userId)
	return resourceTencentCloudBhUserRead(d, meta)
}

func resourceTencentCloudBhUserRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = BhService{client: meta.(*TencentCloudClient).apiV3Conn}
		userId  = d.Id()
	)

	respData, err := service.DescribeBhUserById(ctx, userId)
	if err != nil {
		return err
	}

	if respData == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_bh_user` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	if respData.UserName != nil {
		_ = d.Set("user_name", respData.UserName)
	}

	if respData.RealName != nil {
		_ = d.Set("real_name", respData.RealName)
	}

	if respData.Phone != nil {
		_ = d.Set("phone", respData.Phone)
	}

	if respData.Email != nil {
		_ = d.Set("email", respData.Email)
	}

	if respData.ValidateFrom != nil {
		_ = d.Set("validate_from", respData.ValidateFrom)
	}

	if respData.ValidateTo != nil {
		_ = d.Set("validate_to", respData.ValidateTo)
	}

	if respData.GroupSet != nil {
		groupIdSetList := make([]uint64, 0, len(respData.GroupSet))
		for _, item := range respData.GroupSet {
			if item.Id != nil {
				groupIdSetList = append(groupIdSetList, *item.Id)
			}
		}

		_ = d.Set("group_id_set", groupIdSetList)
	}

	if respData.AuthType != nil {
		_ = d.Set("auth_type", respData.AuthType)
	}

	if respData.ValidateTime != nil {
		_ = d.Set("validate_time", respData.ValidateTime)
	}

	if respData.DepartmentId != nil {
		dResp, err := service.DescribeBhDepartments(ctx)
		if err != nil {
			return err
		}

		if dResp == nil {
			return fmt.Errorf("Departments is nil")
		}

		if dResp.Enabled != nil && *dResp.Enabled {
			_ = d.Set("department_id", respData.DepartmentId)
		}
	}

	if respData.Id != nil {
		_ = d.Set("user_id", respData.Id)
	}

	return nil
}

func resourceTencentCloudBhUserUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId  = getLogId(contextNil)
		userId = d.Id()
	)

	needChange := false
	mutableArgs := []string{"real_name", "phone", "email", "validate_from", "validate_to", "group_id_set", "auth_type", "validate_time", "department_id"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := bhsaas.NewModifyUserRequest()
		if v, ok := d.GetOk("real_name"); ok {
			request.RealName = helper.String(v.(string))
		}

		if v, ok := d.GetOk("phone"); ok {
			request.Phone = helper.String(v.(string))
		}

		if v, ok := d.GetOk("email"); ok {
			request.Email = helper.String(v.(string))
		}

		if v, ok := d.GetOk("validate_from"); ok {
			request.ValidateFrom = helper.String(v.(string))
		}

		if v, ok := d.GetOk("validate_to"); ok {
			request.ValidateTo = helper.String(v.(string))
		}

		if v, ok := d.GetOk("group_id_set"); ok {
			groupIdSetSet := v.(*schema.Set).List()
			for i := range groupIdSetSet {
				groupIdSet := groupIdSetSet[i].(int)
				request.GroupIdSet = append(request.GroupIdSet, helper.IntUint64(groupIdSet))
			}
		}

		if v, ok := d.GetOkExists("auth_type"); ok {
			request.AuthType = helper.IntUint64(v.(int))
		}

		if v, ok := d.GetOk("validate_time"); ok {
			request.ValidateTime = helper.String(v.(string))
		}

		if v, ok := d.GetOk("department_id"); ok {
			request.DepartmentId = helper.String(v.(string))
		}

		request.Id = helper.StrToUint64Point(userId)
		reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyUser(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if reqErr != nil {
			log.Printf("[CRITAL]%s update bh user failed, reason:%+v", logId, reqErr)
			return reqErr
		}
	}

	return resourceTencentCloudBhUserRead(d, meta)
}

func resourceTencentCloudBhUserDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = bhsaas.NewDeleteUsersRequest()
		userId  = d.Id()
	)

	request.IdSet = append(request.IdSet, helper.StrToUint64Point(userId))
	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DeleteUsers(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s delete bh user failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return nil
}

/*
Provides a resource to create an identity center user

Example Usage

```hcl
resource "tencentcloudenterprise_cic_user" "example" {
  zone_id     = "z-1os7c9tyugct"
  user_name   = "tf-example"
  description = "desc."
}
```

Or

```hcl
resource "tencentcloudenterprise_cic_user" "example" {
  zone_id      = "z-1os7c9tyugct"
  user_name    = "tf-example"
  description  = "desc."
  first_name   = "FirstName"
  last_name    = "LastName"
  display_name = "DisplayName"
  email        = "user@example.com"
  user_status  = "Enabled"
}
```

Import

organization identity center user can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_user.example z-1os7c9tyugct#u-rdvm4xdqi8pr
```

 */
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strings"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_user", CNDescription{
		TerraformTypeCN: "身份中心用户",
		DescriptionCN:   "提供身份中心用户资源，用于创建和管理身份中心用户。",
		AttributesCN: map[string]string{
			"zone_id":      "空间ID",
			"user_name":    "用户名",
			"description":  "用户描述",
			"first_name":   "名",
			"last_name":    "姓",
			"display_name": "显示名称",
			"email":        "邮箱地址",
			"user_status":  "用户状态",
			"user_type":    "用户类型",
			"user_id":      "用户ID",
			"create_time":  "创建时间",
			"update_time":  "更新时间",
		},
	})
}

func resourceTencentCloudCicUser() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create an identity center user",
		Create: resourceTencentCloudCicUserCreate,
		Read:   resourceTencentCloudCicUserRead,
		Update: resourceTencentCloudCicUserUpdate,
		Delete: resourceTencentCloudCicUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Zone id.",
			},

			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User name. It must be unique in space. Modifications are not supported. Format: Contains numbers, English letters and special symbols(`+`, `=`, `,`, `.`, `@`, `-`, `_`). Length: Maximum 64 characters.",
			},

			"first_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user's last name. Length: Maximum 64 characters.",
			},

			"last_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user's name. Length: Maximum 64 characters.",
			},

			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The display name of the user. Length: Maximum 256 characters.",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User's description. Length: Maximum 1024 characters.",
			},

			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user's email address. Must be unique within the catalog. Length: Maximum 128 characters.",
			},

			"user_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The status of the user. Value: Enabled (default): Enabled. Disabled: Disabled.",
			},
			"user_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User id.",
			},
			"user_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User type.",
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

func resourceTencentCloudCicUserCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId string
		userId string
	)
	var (
		request  = cic.NewCreateUserRequest()
		response = cic.NewCreateUserResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("user_name"); ok {
		request.UserName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("first_name"); ok {
		request.FirstName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("last_name"); ok {
		request.LastName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("display_name"); ok {
		request.DisplayName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOk("email"); ok {
		request.Email = helper.String(v.(string))
	}

	if v, ok := d.GetOk("user_status"); ok {
		request.UserStatus = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().CreateUserWithContext(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create identity center user failed, reason:%+v", logId, err)
		return err
	}

	userId = *response.Response.UserInfo.UserId

	d.SetId(strings.Join([]string{zoneId, userId}, FILED_SP))

	return resourceTencentCloudCicUserRead(d, meta)
}

func resourceTencentCloudCicUserRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	cicServicex := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	userId := idSplit[1]

	_ = d.Set("zone_id", zoneId)
	_ = d.Set("user_id", userId)

	respData, err := cicServicex.DescribeCicUserById(ctx, zoneId, userId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `cic_user` [%s] not found, please check if it has been deleted.\n",
			logId, d.Id())
		return nil
	}
	if respData.UserName != nil {
		_ = d.Set("user_name", respData.UserName)
	}

	if respData.FirstName != nil {
		_ = d.Set("first_name", respData.FirstName)
	}

	if respData.LastName != nil {
		_ = d.Set("last_name", respData.LastName)
	}

	if respData.DisplayName != nil {
		_ = d.Set("display_name", respData.DisplayName)
	}

	if respData.Description != nil {
		_ = d.Set("description", respData.Description)
	}

	if respData.Email != nil {
		_ = d.Set("email", respData.Email)
	}

	if respData.UserStatus != nil {
		_ = d.Set("user_status", respData.UserStatus)
	}

	if respData.UserType != nil {
		_ = d.Set("user_type", respData.UserType)
	}

	if respData.UserId != nil {
		_ = d.Set("user_id", respData.UserId)
	}

	if respData.CreateTime != nil {
		_ = d.Set("create_time", respData.CreateTime)
	}

	if respData.UpdateTime != nil {
		_ = d.Set("update_time", respData.UpdateTime)
	}

	return nil
}

func resourceTencentCloudCicUserUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)


	immutableArgs := []string{"zone_id", "user_name"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}
	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	userId := idSplit[1]

	needChange := false
	mutableArgs := []string{"user_status"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := cic.NewUpdateUserStatusRequest()

		request.ZoneId = helper.String(zoneId)

		request.UserId = helper.String(userId)

		if v, ok := d.GetOk("user_status"); ok {
			request.NewUserStatus = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().UpdateUserStatus(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update identity center user failed, reason:%+v", logId, err)
			return err
		}
	}

	needChange1 := false
	mutableArgs1 := []string{"first_name", "last_name", "display_name", "description", "email"}
	for _, v := range mutableArgs1 {
		if d.HasChange(v) {
			needChange1 = true
			break
		}
	}

	if needChange1 {
		request1 := cic.NewUpdateUserRequest()

		request1.ZoneId = helper.String(zoneId)

		request1.UserId = helper.String(userId)

		if v, ok := d.GetOk("first_name"); ok {
			request1.NewFirstName = helper.String(v.(string))
		}

		if v, ok := d.GetOk("last_name"); ok {
			request1.NewLastName = helper.String(v.(string))
		}

		if v, ok := d.GetOk("display_name"); ok {
			request1.NewDisplayName = helper.String(v.(string))
		}

		if v, ok := d.GetOk("description"); ok {
			request1.NewDescription = helper.String(v.(string))
		}

		if v, ok := d.GetOk("email"); ok {
			request1.NewEmail = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().UpdateUser(request1)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request1.GetAction(), request1.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update identity center user failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCicUserRead(d, meta)
}

func resourceTencentCloudCicUserDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_user.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	userId := idSplit[1]

	var (
		request  = cic.NewDeleteUserRequest()
		response = cic.NewDeleteUserResponse()
	)

	request.ZoneId = helper.String(zoneId)

	request.UserId = helper.String(userId)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().DeleteUser(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete identity center user failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	return nil
}

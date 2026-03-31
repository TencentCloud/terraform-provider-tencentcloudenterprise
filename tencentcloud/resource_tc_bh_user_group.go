/*
Provide a resource to create a BH user group

# Example Usage

```hcl

	resource "tencentcloudenterprise_bh_user_group" "example" {
	  name = "example-user-group"
	}

```

# Import

BH user group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_bh_user_group.example 12345
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
	registerResourceDescriptionProvider("tencentcloudenterprise_bh_user_group", CNDescription{
		TerraformTypeCN: "BH 用户组",
		DescriptionCN:   "提供 BH 用户组资源，用于创建和管理堡垒机用户组。",
		AttributesCN: map[string]string{
			"name":          "用户组名称",
			"department_id": "用户组所属部门ID",
			"user_group_id": "用户组ID",
		},
	})
}

func ResourceTencentCloudBhUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudBhUserGroupCreate,
		Read:   resourceTencentCloudBhUserGroupRead,
		Update: resourceTencentCloudBhUserGroupUpdate,
		Delete: resourceTencentCloudBhUserGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User group name, maximum length 32 characters.",
			},

			"department_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Department ID to which the user group belongs, e.g.: 1.2.3.",
			},

			// computed
			"user_group_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "User group ID.",
			},
		},
	}
}

func resourceTencentCloudBhUserGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user_group.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		_           = context.WithValue(context.TODO(), logIdKey, logId)
		request     = bhsaas.NewCreateUserGroupRequest()
		response    = bhsaas.NewCreateUserGroupResponse()
		userGroupId string
	)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("department_id"); ok {
		request.DepartmentId = helper.String(v.(string))
	}

	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateUserGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("Create bh user group failed, Response is nil."))
		}

		response = result
		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s create bh user group failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	if response.Response.Id == nil {
		return fmt.Errorf("Id is nil.")
	}

	userGroupId = helper.UInt64ToStr(*response.Response.Id)
	d.SetId(userGroupId)
	return resourceTencentCloudBhUserGroupRead(d, meta)
}

func resourceTencentCloudBhUserGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user_group.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = BhService{client: meta.(*TencentCloudClient).apiV3Conn}
		userGroupId = d.Id()
	)

	respData, err := service.DescribeBhUserGroupById(ctx, userGroupId)
	if err != nil {
		return err
	}

	if respData == nil {
		log.Printf("[WARN]%s resource `tencentcloudenterprise_bh_user_group` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		d.SetId("")
		return nil
	}

	if respData.Name != nil {
		_ = d.Set("name", respData.Name)
	}

	if respData.Department != nil {
		if respData.Department.Id != nil {
			dResp, err := service.DescribeBhDepartments(ctx)
			if err != nil {
				return err
			}

			if dResp == nil {
				return fmt.Errorf("Departments is nil")
			}

			if dResp.Enabled != nil && *dResp.Enabled {
				_ = d.Set("department_id", respData.Department.Id)
			}
		}
	}

	if respData.Id != nil {
		_ = d.Set("user_group_id", respData.Id)
	}

	return nil
}

func resourceTencentCloudBhUserGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user_group.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		_           = context.WithValue(context.TODO(), logIdKey, logId)
		userGroupId = d.Id()
	)

	needChange := false
	mutableArgs := []string{"name", "department_id"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := bhsaas.NewModifyUserGroupRequest()
		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}

		if v, ok := d.GetOk("department_id"); ok {
			request.DepartmentId = helper.String(v.(string))
		}

		request.Id = helper.StrToUint64Point(userGroupId)
		reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyUserGroup(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if reqErr != nil {
			log.Printf("[CRITAL]%s update bh user group failed, reason:%+v", logId, reqErr)
			return reqErr
		}
	}

	return resourceTencentCloudBhUserGroupRead(d, meta)
}

func resourceTencentCloudBhUserGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_bh_user_group.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		_           = context.WithValue(context.TODO(), logIdKey, logId)
		request     = bhsaas.NewDeleteUserGroupsRequest()
		userGroupId = d.Id()
	)

	request.IdSet = append(request.IdSet, helper.StrToUint64Point(userGroupId))
	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().DeleteUserGroups(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s delete bh user group failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return nil
}

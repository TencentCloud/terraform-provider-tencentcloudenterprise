/*
Provide a resource to create a DASB user group

# Example Usage

```hcl

	resource "tencentcloudenterprise_dasb_user_group" "example" {
	  name = "example-user-group"
	}

```

# Import

DASB user group can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_dasb_user_group.example 12345
```
*/
package tencentcloud

import (
	"context"
	"log"
	"strconv"

	bhsaas "terraform-provider-tencentcloudenterprise/sdk/bhsaas/v20191018"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dasb_user_group", CNDescription{
		TerraformTypeCN: "DASB 用户组",
		DescriptionCN:   "提供 DASB 用户组资源，用于创建和管理用户组。",
		AttributesCN: map[string]string{
			"name":          "名称",
			"department_id": "部门ID",
		},
	})
}

func ResourceTencentCloudDasbUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudDasbUserGroupCreate,
		Read:   resourceTencentCloudDasbUserGroupRead,
		Update: resourceTencentCloudDasbUserGroupUpdate,
		Delete: resourceTencentCloudDasbUserGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "User group name, maximum length 32 characters.",
			},
			"department_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "ID of the department to which the user group belongs, such as: 1.2.3.",
			},
		},
	}
}

func resourceTencentCloudDasbUserGroupCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_user_group.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
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

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().CreateUserGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create dasb UserGroup failed, reason:%+v", logId, err)
		return err
	}

	userGroupIdInt := *response.Response.Id
	userGroupId = strconv.FormatUint(userGroupIdInt, 10)
	d.SetId(userGroupId)

	return resourceTencentCloudDasbUserGroupRead(d, meta)
}

func resourceTencentCloudDasbUserGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_user_group.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		userGroupId = d.Id()
	)

	UserGroup, err := service.DescribeDasbUserGroupById(ctx, userGroupId)
	if err != nil {
		return err
	}

	if UserGroup == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `DasbUserGroup` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if UserGroup.Name != nil {
		_ = d.Set("name", UserGroup.Name)
	}

	if UserGroup.Department != nil {
		departmentId := *UserGroup.Department.Id
		if departmentId != "1" {
			_ = d.Set("department_id", departmentId)
		}
	}

	return nil
}

func resourceTencentCloudDasbUserGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_user_group.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		request     = bhsaas.NewModifyUserGroupRequest()
		userGroupId = d.Id()
	)

	userGroupIdInt, _ := strconv.ParseUint(userGroupId, 10, 64)
	request.Id = &userGroupIdInt

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("department_id"); ok {
		request.DepartmentId = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseBhsaasClient().ModifyUserGroup(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update dasb UserGroup failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudDasbUserGroupRead(d, meta)
}

func resourceTencentCloudDasbUserGroupDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_dasb_user_group.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = DasbService{client: meta.(*TencentCloudClient).apiV3Conn}
		userGroupId = d.Id()
	)

	if err := service.DeleteDasbUserGroupById(ctx, userGroupId); err != nil {
		return err
	}

	return nil
}

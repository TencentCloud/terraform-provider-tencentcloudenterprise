/*
Provides a resource to create a cam service_linked_role

# Example Usage

```hcl

	resource "tencentcloudenterprise_cam_service_linked_role" "service_linked_role" {
	  qcs_service_name = ["tke.cloud.com"]
	  custom_suffix    = "test-suffix"
	  description      = "test-description"
	}

```
*/
package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cam_service_linked_role", CNDescription{
		TerraformTypeCN: "CAM服务相关角色",
		DescriptionCN:   "提供 CAM 服务相关角色资源，用于创建和管理服务相关角色。",
		AttributesCN: map[string]string{
			"qcs_service_name": "授权服务，即该角色关联的腾讯云服务主体",
			"custom_suffix":    "自定义后缀，与服务提供的固定前缀结合形成全名",
			"description":      "角色描述",
		},
	})
}

func resourceTencentCloudCamServiceLinkedRole() *schema.Resource {
	return &schema.Resource{
		Read:   resourceTencentCloudCamServiceLinkedRoleRead,
		Create: resourceTencentCloudCamServiceLinkedRoleCreate,
		Update: resourceTencentCloudCamServiceLinkedRoleUpdate,
		Delete: resourceTencentCloudCamServiceLinkedRoleDelete,
		Schema: map[string]*schema.Schema{
			"qcs_service_name": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Required:    true,
				ForceNew:    true,
				Description: "Authorization service, the Tencent Cloud service principal with this role attached.",
			},

			"custom_suffix": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The custom suffix, based on the string you provide, is combined with the prefix provided by the service to form the full role name. This field is not allowed to contain the character `_`.",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Role description.",
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceTencentCloudCamServiceLinkedRoleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_service_linked_role.create")()

	logId := getLogId(contextNil)

	var (
		request  = cam.NewCreateServiceLinkedRoleRequest()
		response *cam.CreateServiceLinkedRoleResponse
		roleId   string
	)

	if v, ok := d.GetOk("qcs_service_name"); ok {
		serviceName := v.(*schema.Set).List()
		serviceNameArr := make([]*string, 0, len(serviceName))
		for _, name := range serviceName {
			serviceNameArr = append(serviceNameArr, helper.String(name.(string)))
		}
		request.QCSServiceName = serviceNameArr
	}

	if v, ok := d.GetOk("custom_suffix"); ok {
		request.CustomSuffix = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().CreateServiceLinkedRole(request)
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
		log.Printf("[CRITAL]%s create cam serviceLinkedRole failed, reason:%+v", logId, err)
		return err
	}

	roleId = *response.Response.RoleId

	d.SetId(roleId)

	return resourceTencentCloudCamServiceLinkedRoleRead(d, meta)
}

func resourceTencentCloudCamServiceLinkedRoleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_service_linked_role.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	roleId := d.Id()

	serviceLinkedRole, err := service.DescribeCamServiceLinkedRole(ctx, roleId)
	if err != nil {
		return err
	}

	if serviceLinkedRole == nil {
		d.SetId("")
		return fmt.Errorf("resource `serviceLinkedRole` %s does not exist", roleId)
	}

	if serviceLinkedRole.PolicyDocument != nil {
		var documentJson CamDocument
		err = json.Unmarshal([]byte(*serviceLinkedRole.PolicyDocument), &documentJson)
		if err != nil {
			return err
		}
		if documentJson.Statement != nil && len(documentJson.Statement) > 0 {
			principal := documentJson.Statement[0].Principal
			if principal.Service != nil && len(principal.Service) > 0 {
				_ = d.Set("qcs_service_name", principal.Service)
			}
		}
	}

	if serviceLinkedRole.RoleName != nil {
		roleName := strings.Split(*serviceLinkedRole.RoleName, "_")
		if len(roleName) > 0 {
			_ = d.Set("custom_suffix", roleName[len(roleName)-1])
		}
	}

	if serviceLinkedRole.Description != nil {
		_ = d.Set("description", serviceLinkedRole.Description)
	}

	return nil
}

func resourceTencentCloudCamServiceLinkedRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_service_linked_role.update")()

	logId := getLogId(contextNil)

	roleId := d.Id()

	if d.HasChange("description") {
		request := cam.NewUpdateRoleDescriptionRequest()
		request.RoleId = &roleId

		if v, ok := d.GetOk("description"); ok {
			request.Description = helper.String(v.(string))
		}

		// ConsoleLogin is required, ServiceLinkedRole default to 0 (no console login)
		consoleLogin := uint64(0)
		request.ConsoleLogin = &consoleLogin

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCamClient().UpdateRoleDescription(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
					logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update cam serviceLinkedRole failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCamServiceLinkedRoleRead(d, meta)
}

func resourceTencentCloudCamServiceLinkedRoleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cam_service_linked_role.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	roleId := d.Id()

	serviceLinkedRole, err := service.DescribeCamServiceLinkedRole(ctx, roleId)
	if err != nil {
		return err
	}
	if serviceLinkedRole == nil || serviceLinkedRole.RoleName == nil {
		return fmt.Errorf("when querying serviceLinkedRole, an error occurs")
	}

	// 本地SDK不支持删除状态轮询，直接调用删除接口
	err = service.DeleteCamServiceLinkedRoleByName(ctx, *serviceLinkedRole.RoleName)
	if err != nil {
		return err
	}

	return nil
}

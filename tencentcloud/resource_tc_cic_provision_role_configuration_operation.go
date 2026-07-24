/*
Provides a resource to create a cic provision_role_configuration_operation

Example Usage

```hcl
data "tencentcloudenterprise_cic_identity_center" "center" {}

resource "tencentcloudenterprise_cic_role_configuration" "example" {
  zone_id                 = data.tencentcloudenterprise_cic_identity_center.center.zone_id
  role_configuration_name = "tf-provision-example"
  description             = "tf example"
}

resource "tencentcloudenterprise_cic_provision_role_configuration_operation" "example" {
  zone_id               = data.tencentcloudenterprise_cic_identity_center.center.zone_id
  role_configuration_id = tencentcloudenterprise_cic_role_configuration.example.role_configuration_id
  target_type           = "MemberUin"
  target_uin            = 100001234567
}
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strings"
	"time"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_provision_role_configuration_operation", CNDescription{
		TerraformTypeCN: "身份中心权限配置部署操作",
		DescriptionCN:   "提供身份中心权限配置部署操作资源，用于将权限配置（重新）部署到成员账号。",
		AttributesCN: map[string]string{
			"zone_id":               "空间ID",
			"role_configuration_id": "权限配置ID",
			"target_type":           "目标账号类型",
			"target_uin":            "目标账号UIN",
		},
	})
}

func resourceTencentCloudCicProvisionRoleConfigurationOperation() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudCicProvisionRoleConfigurationOperationCreate,
		Read:   resourceTencentCloudCicProvisionRoleConfigurationOperationRead,
		Delete: resourceTencentCloudCicProvisionRoleConfigurationOperationDelete,
		Description: "Provide an identity center permission configuration provision operation resource to deploy or redeploy a role configuration to a member account.",
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Space ID.",
			},
			"role_configuration_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Permission configuration ID.",
			},
			"target_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Type of the synchronized target account of the Tencent Cloud Organization. ManagerUin: admin account; MemberUin: member account.",
			},
			"target_uin": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "UIN of the target account of the Tencent Cloud Organization.",
			},
		},
	}
}

func resourceTencentCloudCicProvisionRoleConfigurationOperationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_provision_role_configuration_operation.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId              string
		roleConfigurationId string
		targetType          string
		targetUin           int
		request             = cic.NewProvisionRoleConfigurationRequest()
		response            = cic.NewProvisionRoleConfigurationResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
		request.ZoneId = helper.String(zoneId)
	}
	if v, ok := d.GetOk("role_configuration_id"); ok {
		roleConfigurationId = v.(string)
		request.RoleConfigurationId = helper.String(roleConfigurationId)
	}
	if v, ok := d.GetOk("target_type"); ok {
		targetType = v.(string)
		request.TargetType = helper.String(targetType)
	}
	if v, ok := d.GetOkExists("target_uin"); ok {
		targetUin = v.(int)
		request.TargetUin = helper.IntInt64(targetUin)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().ProvisionRoleConfiguration(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cic provision role configuration operation failed, reason:%+v", logId, err)
		return err
	}

	if response.Response == nil || response.Response.Task == nil || response.Response.Task.TaskId == nil {
		return fmt.Errorf("can not find taskId")
	}

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}
	taskId := *response.Response.Task.TaskId
	if _, err := (&resource.StateChangeConf{
		Delay:      5 * time.Second,
		MinTimeout: 3 * time.Second,
		Pending:    []string{"InProgress", ""},
		Refresh:    service.AssignmentTaskStatusStateRefreshFunc(zoneId, taskId, []string{"Failed"}),
		Target:     []string{"Success"},
		Timeout:    600 * time.Second,
	}).WaitForStateContext(ctx); err != nil {
		return err
	}

	d.SetId(strings.Join([]string{zoneId, roleConfigurationId, targetType, helper.IntToStr(targetUin)}, FILED_SP))
	return resourceTencentCloudCicProvisionRoleConfigurationOperationRead(d, meta)
}

func resourceTencentCloudCicProvisionRoleConfigurationOperationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_provision_role_configuration_operation.read")()
	defer inconsistentCheck(d, meta)()
	return nil
}

func resourceTencentCloudCicProvisionRoleConfigurationOperationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_provision_role_configuration_operation.delete")()
	defer inconsistentCheck(d, meta)()
	return nil
}

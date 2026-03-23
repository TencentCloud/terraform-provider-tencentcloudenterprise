/*
Provides a resource to create an organization cic_role_configuration

Example Usage

```hcl
resource "tencentcloudenterprise_cic_role_configuration" "cic_role_configuration" {
    zone_id = "z-xxxxxx"
    role_configuration_name = "tf-test"
    description = "test"
}
```

Import

organization cic_role_configuration can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cic_role_configuration.cic_role_configuration ${zoneId}#${roleConfigurationId}
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cic_role_configuration", CNDescription{
		TerraformTypeCN: "身份中心角色配置",
		DescriptionCN:   "提供身份中心角色配置资源，用于创建和管理身份中心角色配置。",
		AttributesCN: map[string]string{
			"zone_id":                  "空间ID",
			"role_configuration_name":  "角色配置名称",
			"description":              "角色配置描述",
			"session_duration":         "会话持续时间",
			"relay_state":              "中继状态",
			"role_configuration_id":    "角色配置ID",
			"create_time":              "创建时间",
			"update_time":              "更新时间",
		},
	})
}

func resourceTencentCloudCicRoleConfiguration() *schema.Resource {
	return &schema.Resource{
		Description: "Provide identity center role configuration resources for creating and managing identity center role configurations.",
		Create: resourceTencentCloudCicRoleConfigurationCreate,
		Read:   resourceTencentCloudCicRoleConfigurationRead,
		Update: resourceTencentCloudCicRoleConfigurationUpdate,
		Delete: resourceTencentCloudCicRoleConfigurationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Space ID.",
			},

			"role_configuration_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Access configuration name, which contains up to 128 characters, including English letters, digits, and hyphens (-).",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Access configuration description, which contains up to 1024 characters.",
			},

			"session_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Session duration. It indicates the maximum session duration when CIC users use the access configuration to access the target account of the Tencent Cloud Cic. Unit: seconds. Value range: 900-43,200 (15 minutes to 12 hours). Default value: 3600 (1 hour).",
			},

			"relay_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access the target account of the Tencent Cloud Cic. This page must be the Tencent Cloud console page. The default is null, which indicates navigating to the home page of the Tencent Cloud console.",
			},
			"role_configuration_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Role configuration id.",
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

func resourceTencentCloudCicRoleConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	var (
		zoneId              string
		roleConfigurationId string
	)
	var (
		request  = cic.NewCreateRoleConfigurationRequest()
		response = cic.NewCreateRoleConfigurationResponse()
	)

	if v, ok := d.GetOk("zone_id"); ok {
		zoneId = v.(string)
	}

	if v, ok := d.GetOk("zone_id"); ok {
		request.ZoneId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("role_configuration_name"); ok {
		request.RoleConfigurationName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("description"); ok {
		request.Description = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("session_duration"); ok {
		request.SessionDuration = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("relay_state"); ok {
		request.RelayState = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().CreateRoleConfiguration(request)
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
		log.Printf("[CRITAL]%s create identity center role configuration failed, reason:%+v", logId, err)
		return err
	}

	roleConfigurationId = *response.Response.RoleConfigurationInfo.RoleConfigurationId

	d.SetId(strings.Join([]string{zoneId, roleConfigurationId}, FILED_SP))

	return resourceTencentCloudCicRoleConfigurationRead(d, meta)
}

func resourceTencentCloudCicRoleConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.Background(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	roleConfigurationId := idSplit[1]

	_ = d.Set("zone_id", zoneId)

	respData, err := service.DescribeCicRoleConfigurationById(ctx, zoneId, roleConfigurationId)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf(
			"[WARN]%s resource `cic_role_configuration` [%s] not found, please check if it has been deleted.\n",
			logId, d.Id())
		return nil
	}
	if respData.RoleConfigurationId != nil {
		_ = d.Set("role_configuration_id", respData.RoleConfigurationId)
	}

	if respData.RoleConfigurationName != nil {
		_ = d.Set("role_configuration_name", respData.RoleConfigurationName)
	}

	if respData.Description != nil {
		_ = d.Set("description", respData.Description)
	}

	if respData.SessionDuration != nil {
		_ = d.Set("session_duration", respData.SessionDuration)
	}

	if respData.RelayState != nil {
		_ = d.Set("relay_state", respData.RelayState)
	}

	if respData.CreateTime != nil {
		_ = d.Set("create_time", respData.CreateTime)
	}

	if respData.UpdateTime != nil {
		_ = d.Set("update_time", respData.UpdateTime)
	}

	return nil
}

func resourceTencentCloudCicRoleConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	immutableArgs := []string{"zone_id", "role_configuration_name"}
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
	roleConfigurationId := idSplit[1]

	needChange := false
	mutableArgs := []string{"description", "session_duration", "relay_state"}
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		request := cic.NewUpdateRoleConfigurationRequest()

		request.ZoneId = helper.String(zoneId)

		request.RoleConfigurationId = helper.String(roleConfigurationId)

		if v, ok := d.GetOk("description"); ok {
			request.NewDescription = helper.String(v.(string))
		}

		if v, ok := d.GetOkExists("session_duration"); ok {
			request.NewSessionDuration = helper.IntInt64(v.(int))
		}

		if v, ok := d.GetOk("relay_state"); ok {
			request.NewRelayState = helper.String(v.(string))
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().UpdateRoleConfiguration(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId,
					request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update identity center role configuration failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudCicRoleConfigurationRead(d, meta)
}

func resourceTencentCloudCicRoleConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cic_role_configuration.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	// ctx := context.WithValue(context.Background(), logIdKey, logId)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	zoneId := idSplit[0]
	roleConfigurationId := idSplit[1]

	var (
		request  = cic.NewDeleteRoleConfigurationRequest()
		response = cic.NewDeleteRoleConfigurationResponse()
	)

	request.ZoneId = helper.String(zoneId)

	request.RoleConfigurationId = helper.String(roleConfigurationId)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseCicClient().DeleteRoleConfiguration(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete identity center role configuration failed, reason:%+v", logId, err)
		return err
	}

	_ = response
	return nil
}

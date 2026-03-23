/*
Use this data source to query detailed information of cic role configurations

# Example Usage

```hcl

	data "tencentcloudenterprise_cic_role_configurations" "role_configurations" {
	  zone_id = "z-xxxxxxxxxx"
	}

	output "role_configurations_list" {
	  value = data.tencentcloudenterprise_cic_role_configurations.role_configurations.role_configurations
	}

```
*/
package tencentcloud

import (
	"context"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cic_role_configurations", CNDescription{
		TerraformTypeCN: "CIC 权限配置列表",
		DescriptionCN:   "提供 CIC 权限配置列表数据源，用于查询满足条件的权限配置信息。",
		AttributesCN: map[string]string{
			"zone_id":                 "空间 ID",
			"filter":                  "过滤文本，不区分大小写，支持 RoleConfigurationName 和 Description",
			"filter_targets":          "检索成员账号是否配置过权限的账号 UIN 列表",
			"principal_id":            "授权的用户 UserId 或用户组 GroupId，需与 filter_targets 一起设置",
			"role_configurations":     "权限配置列表",
			"role_configuration_id":   "权限配置 ID",
			"role_configuration_name": "权限配置名称",
			"description":             "权限配置描述",
			"session_duration":        "会话持续时间，单位秒",
			"relay_state":             "初始访问页面 URL",
			"create_time":             "权限配置创建时间",
			"update_time":             "权限配置更新时间",
			"is_selected":             "若 filter_targets 有值，配置了权限则返回 true，否则 false",
			"result_output_file":      "用于保存结果，可视化界面不可用",
		},
	})
}

func dataSourceTencentCloudCicRoleConfigurations() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudCicRoleConfigurationsRead,
		Description: "Use this data source to query detailed information of cic role configurations.",
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Space ID.",
			},

			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter criteria, case insensitive. Currently supports RoleConfigurationName and Description.",
			},

			"filter_targets": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Check whether the member account has been configured with permissions. If configured, return IsSelected: true; otherwise, return false.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			"principal_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "UserId of the authorized user or GroupId of the authorized user group, which must be set together with the input parameter FilterTargets.",
			},

			"role_configurations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Permission configuration list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_configuration_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Permission configuration ID.",
						},
						"role_configuration_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Permission configuration name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Permission configuration description.",
						},
						"session_duration": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Session duration. It indicates the maximum session duration when CIC users use the access configuration to access member accounts. Unit: seconds.",
						},
						"relay_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Initial access page. It indicates the initial access page URL when CIC users use the access configuration to access member accounts.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the permission configuration.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of the permission configuration.",
						},
						"is_selected": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If the input parameter FilterTargets is provided, check whether the member account has been configured with permissions. If configured, return true; otherwise, return false.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudCicRoleConfigurationsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cic_role_configurations.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("zone_id"); ok {
		paramMap["ZoneId"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("filter"); ok {
		paramMap["Filter"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("filter_targets"); ok {
		filterTargetsList := []*int64{}
		filterTargetsSet := v.(*schema.Set).List()
		for i := range filterTargetsSet {
			filterTargets := filterTargetsSet[i].(int)
			filterTargetsList = append(filterTargetsList, helper.IntInt64(filterTargets))
		}
		paramMap["FilterTargets"] = filterTargetsList
	}
	if v, ok := d.GetOk("principal_id"); ok {
		paramMap["PrincipalId"] = helper.String(v.(string))
	}

	var roleConfigurations []*cic.RoleConfiguration
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCicRoleConfigurationsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		roleConfigurations = result
		return nil
	})
	if err != nil {
		return err
	}

	roleConfigurationsList := make([]map[string]interface{}, 0, len(roleConfigurations))
	ids := make([]string, 0, len(roleConfigurations))
	for _, roleConfiguration := range roleConfigurations {
		roleConfigurationsMap := map[string]interface{}{}
		if roleConfiguration.RoleConfigurationId != nil {
			roleConfigurationsMap["role_configuration_id"] = roleConfiguration.RoleConfigurationId
			ids = append(ids, *roleConfiguration.RoleConfigurationId)
		}
		if roleConfiguration.RoleConfigurationName != nil {
			roleConfigurationsMap["role_configuration_name"] = roleConfiguration.RoleConfigurationName
		}
		if roleConfiguration.Description != nil {
			roleConfigurationsMap["description"] = roleConfiguration.Description
		}
		if roleConfiguration.SessionDuration != nil {
			roleConfigurationsMap["session_duration"] = roleConfiguration.SessionDuration
		}
		if roleConfiguration.RelayState != nil {
			roleConfigurationsMap["relay_state"] = roleConfiguration.RelayState
		}
		if roleConfiguration.CreateTime != nil {
			roleConfigurationsMap["create_time"] = roleConfiguration.CreateTime
		}
		if roleConfiguration.UpdateTime != nil {
			roleConfigurationsMap["update_time"] = roleConfiguration.UpdateTime
		}
		if roleConfiguration.IsSelected != nil {
			roleConfigurationsMap["is_selected"] = roleConfiguration.IsSelected
		}
		roleConfigurationsList = append(roleConfigurationsList, roleConfigurationsMap)
	}

	_ = d.Set("role_configurations", roleConfigurationsList)
	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), roleConfigurationsList); e != nil {
			return e
		}
	}

	return nil
}

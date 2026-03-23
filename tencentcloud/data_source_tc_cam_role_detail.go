/*
Use this data source to query CAM role details.

# Example Usage

```hcl

	data "tencentcloudenterprise_cam_role_detail" "foo" {
	  role_name = "test-role"
	}

```
*/
package tencentcloud

import (
	"context"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_role_detail", CNDescription{
		TerraformTypeCN: "CAM角色详情",
		DescriptionCN:   "用于查询 CAM 角色的详细信息。",
		AttributesCN: map[string]string{
			"role_id":          "角色 ID",
			"role_name":        "角色名称",
			"role_info":        "角色详细信息",
			"policy_document":  "策略文档",
			"description":      "角色描述",
			"add_time":         "创建时间",
			"update_time":      "修改时间",
			"console_login":    "是否允许登录控制台",
			"role_type":        "角色类型",
			"session_duration": "有效期",
			"deletion_task_id": "删除任务 ID",
		},
	})
}

func dataSourceTencentCloudCamRoleDetail() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamRoleDetailRead,
		Schema: map[string]*schema.Schema{
			"role_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Role ID, used to specify role. Input either `role_id` or `role_name`.",
			},

			"role_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Role name, used to specify role. Input either `role_id` or `role_name`.",
			},

			"role_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Role details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Role ID.",
						},
						"role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Role name.",
						},
						"policy_document": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Role policy document.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Role description.",
						},
						"add_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time role created.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time role last updated.",
						},
						"console_login": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "If login is allowed for the role.",
						},
						"role_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User role. Valid values: `user`, `system`, `service_linked`.",
						},
						"session_duration": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Valid period.",
						},
						"deletion_task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task identifier for deleting a service-linked role.",
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

func dataSourceTencentCloudCamRoleDetailRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencenttencentcloudenterprise_cam_role_detail.read")()

	ctx := context.Background()
	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("role_id"); ok {
		paramMap["RoleId"] = v.(string)
	}

	if v, ok := d.GetOk("role_name"); ok {
		paramMap["RoleName"] = v.(string)
	}

	var respData *cam.RoleInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCamRoleDetailByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		respData = result
		return nil
	})
	if err != nil {
		return err
	}

	var roleId string
	roleInfoMap := map[string]interface{}{}

	if respData != nil {
		if respData.RoleId != nil {
			roleInfoMap["role_id"] = *respData.RoleId
			roleId = *respData.RoleId
		}

		if respData.RoleName != nil {
			roleInfoMap["role_name"] = *respData.RoleName
		}

		if respData.PolicyDocument != nil {
			roleInfoMap["policy_document"] = *respData.PolicyDocument
		}

		if respData.Description != nil {
			roleInfoMap["description"] = *respData.Description
		}

		if respData.AddTime != nil {
			roleInfoMap["add_time"] = *respData.AddTime
		}

		if respData.UpdateTime != nil {
			roleInfoMap["update_time"] = *respData.UpdateTime
		}

		if respData.ConsoleLogin != nil {
			roleInfoMap["console_login"] = int(*respData.ConsoleLogin)
		}

		if respData.RoleType != nil {
			roleInfoMap["role_type"] = *respData.RoleType
		}

		if respData.SessionDuration != nil {
			roleInfoMap["session_duration"] = int(*respData.SessionDuration)
		}

		if respData.DeletionTaskId != nil {
			roleInfoMap["deletion_task_id"] = *respData.DeletionTaskId
		}

		_ = d.Set("role_info", []interface{}{roleInfoMap})
	}

	d.SetId(roleId)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), roleInfoMap); e != nil {
			return e
		}
	}

	return nil
}

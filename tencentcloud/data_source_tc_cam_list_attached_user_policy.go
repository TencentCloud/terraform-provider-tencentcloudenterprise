/*
Use this data source to query policies attached to a CAM user.

# Example Usage

```hcl

	data "tencentcloudenterprise_cam_list_attached_user_policy" "foo" {
	  target_uin  = 100000000001
	  attach_type = 0
	}

```
*/
package tencentcloud

import (
	"context"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_list_attached_user_policy", CNDescription{
		TerraformTypeCN: "CAM用户关联策略列表",
		DescriptionCN:   "用于查询指定 CAM 用户关联的所有策略（包括直接关联和随组关联）。",
		AttributesCN: map[string]string{
			"target_uin":    "目标用户 UIN",
			"attach_type":   "关联类型 (0: 全部; 1: 直接关联; 2: 随组关联)",
			"keyword":       "搜索关键词",
			"policy_list":   "策略列表",
			"policy_id":     "策略 ID",
			"policy_name":   "策略名称",
			"description":   "描述",
			"add_time":      "添加时间",
			"strategy_type": "策略类型 (1: 自定义策略; 2: 预设策略)",
			"create_mode":   "创建模式",
			"groups":        "所属用户组信息",
			"group_id":      "用户组 ID",
			"group_name":    "用户组名称",
		},
	})
}

func dataSourceTencentCloudCamListAttachedUserPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamListAttachedUserPolicyRead,
		Schema: map[string]*schema.Schema{
			"target_uin": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Target User ID.",
			},

			"attach_type": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "0: Return direct association and group association policies, 1: Only return direct association policies, 2: Only return group association policies.",
			},

			"strategy_type": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Policy type.",
			},

			"keyword": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Search Keywords.",
			},

			"policy_list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Policy List Data.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy ID.",
						},
						"policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy Name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy Description.",
						},
						"add_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"strategy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Policy type (1 represents custom policy, 2 represents preset policy).",
						},
						"create_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation mode (1 represents policies created by product or project permissions, others represent policies created by policy syntax).",
						},
						"groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Associated information with group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Group ID.",
									},
									"group_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Group Name.",
									},
								},
							},
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

func dataSourceTencentCloudCamListAttachedUserPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencenttencentcloudenterprise_cam_list_attached_user_policy.read")()

	ctx := context.Background()
	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("target_uin"); ok {
		paramMap["TargetUin"] = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("attach_type"); ok {
		paramMap["AttachType"] = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("strategy_type"); ok {
		paramMap["StrategyType"] = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("keyword"); ok {
		paramMap["Keyword"] = v.(string)
	}

	var policyList []*cam.AttachedUserPolicy

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCamListAttachedUserPolicyByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		policyList = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(policyList))
	tmpList := make([]map[string]interface{}, 0, len(policyList))

	if policyList != nil {
		for _, attachedUserPolicy := range policyList {
			attachedUserPolicyMap := map[string]interface{}{}

			if attachedUserPolicy.PolicyId != nil {
				attachedUserPolicyMap["policy_id"] = helper.UInt64ToStr(*attachedUserPolicy.PolicyId)
				ids = append(ids, helper.UInt64ToStr(*attachedUserPolicy.PolicyId))
			}

			if attachedUserPolicy.PolicyName != nil {
				attachedUserPolicyMap["policy_name"] = *attachedUserPolicy.PolicyName
			}

			if attachedUserPolicy.Description != nil {
				attachedUserPolicyMap["description"] = *attachedUserPolicy.Description
			}

			if attachedUserPolicy.AddTime != nil {
				attachedUserPolicyMap["add_time"] = *attachedUserPolicy.AddTime
			}

			if attachedUserPolicy.StrategyType != nil {
				attachedUserPolicyMap["strategy_type"] = helper.UInt64ToStr(*attachedUserPolicy.StrategyType)
			}

			if attachedUserPolicy.CreateMode != nil {
				attachedUserPolicyMap["create_mode"] = helper.UInt64ToStr(*attachedUserPolicy.CreateMode)
			}

			if attachedUserPolicy.Groups != nil {
				groupsList := []interface{}{}
				for _, groups := range attachedUserPolicy.Groups {
					groupsMap := map[string]interface{}{}

					if groups.GroupId != nil {
						groupsMap["group_id"] = int(*groups.GroupId)
					}

					if groups.GroupName != nil {
						groupsMap["group_name"] = *groups.GroupName
					}

					groupsList = append(groupsList, groupsMap)
				}

				attachedUserPolicyMap["groups"] = groupsList
			}

			tmpList = append(tmpList, attachedUserPolicyMap)
		}

		_ = d.Set("policy_list", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}

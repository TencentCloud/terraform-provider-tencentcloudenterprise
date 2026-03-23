/*
Use this data source to query entities associated with a CAM policy.

# Example Usage

```hcl

	data "tencentcloudenterprise_cam_list_entities_for_policy" "foo" {
	  policy_id     = 12345678
	  entity_filter = "User"
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
	registerDataDescriptionProvider("tencentcloudenterprise_cam_list_entities_for_policy", CNDescription{
		TerraformTypeCN: "CAM策略关联实体列表",
		DescriptionCN:   "用于查询与 CAM 策略关联的实体（用户、用户组、角色）列表。",
		AttributesCN: map[string]string{
			"policy_id":     "策略 ID",
			"rp":            "每页数量，默认 20",
			"entity_filter": "实体过滤器，可选值: 'All', 'User', 'Group', 'Role'",
			"list":          "关联实体列表",
			"id":            "实体 ID",
			"name":          "实体名称",
			"uin":           "实体 Uin",
			"related_type":  "关联类型 (1: 用户关联; 2: 用户组关联)",
		},
	})
}

func dataSourceTencentCloudCamListEntitiesForPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamListEntitiesForPolicyRead,
		Schema: map[string]*schema.Schema{
			"policy_id": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Policy Id.",
			},

			"rp": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Per page size, default value is 20.",
			},

			"entity_filter": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Can take values of 'All', 'User', 'Group', and 'Role'. 'All' represents obtaining all entity types, 'User' represents only obtaining sub accounts, 'Group' represents only obtaining user groups, and 'Role' represents only obtaining roles. The default value is 'All'.",
			},

			"list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Entity List.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Entity ID.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Entity Name.",
						},
						"uin": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Entity Uin.",
						},
						"related_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Association type. 1. User association; 2 User Group Association.",
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

func dataSourceTencentCloudCamListEntitiesForPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencenttencentcloudenterprise_cam_list_entities_for_policy.read")()

	ctx := context.Background()
	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("policy_id"); ok {
		paramMap["PolicyId"] = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("rp"); ok {
		paramMap["Rp"] = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("entity_filter"); ok {
		paramMap["EntityFilter"] = v.(string)
	}

	var listEntitiesForPolicy []*cam.AttachEntityOfPolicy
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCamListEntitiesForPolicyByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		listEntitiesForPolicy = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(listEntitiesForPolicy))
	tmpList := make([]map[string]interface{}, 0)

	if listEntitiesForPolicy != nil {
		for _, attachEntityOfPolicy := range listEntitiesForPolicy {
			attachEntityOfPolicyMap := map[string]interface{}{}

			if attachEntityOfPolicy.Id != nil {
				attachEntityOfPolicyMap["id"] = *attachEntityOfPolicy.Id
			}

			if attachEntityOfPolicy.Name != nil {
				attachEntityOfPolicyMap["name"] = *attachEntityOfPolicy.Name
			}

			if attachEntityOfPolicy.Uin != nil {
				attachEntityOfPolicyMap["uin"] = int(*attachEntityOfPolicy.Uin)
			}

			if attachEntityOfPolicy.RelatedType != nil {
				attachEntityOfPolicyMap["related_type"] = int(*attachEntityOfPolicy.RelatedType)
			}

			if attachEntityOfPolicy.Uin != nil {
				ids = append(ids, helper.UInt64ToStr(*attachEntityOfPolicy.Uin))
			}
			tmpList = append(tmpList, attachEntityOfPolicyMap)
		}

		_ = d.Set("list", tmpList)
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

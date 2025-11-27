/*
Use this data source to query detailed information of tdmqRocketmq role

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
		cluster_name = "test_rocketmq_datasource_role"
		remark = "test recket mq"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_role" "role" {
	  role_name = "test_rocketmq_role"
	  remark = "test rocketmq role"
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	}

	data "tencentcloudenterprise_tdmq_rocketmq_role" "role" {
	  role_name = cloud_tdmq_rocketmq_role.role.role_name
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmqRocketmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rocketmq_role", CNDescription{
		TerraformTypeCN: "RocketMQ 角色",
		DescriptionCN:   "提供TDMQ RocketMQ角色数据源，用于查询TDMQ RocketMQ角色的详细信息。",
		AttributesCN: map[string]string{
			"role_name":          "按角色名称进行模糊查询",
			"cluster_id":         "群集ID（必需）",
			"role_sets":          "角色数组",
			"token":              "角色令牌的值",
			"remark":             "备注",
			"create_time":        "创建时间",
			"update_time":        "更新时间",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTdmqRocketmqRole() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tdmqRocketmq role",
		Read:        dataSourceTencentCloudTdmqRocketmqRoleRead,
		Schema: map[string]*schema.Schema{
			"role_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Fuzzy query by role name.",
			},

			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster ID (required).",
			},

			"role_sets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of roles.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Role name.",
						},
						"token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Value of the role token.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remarks.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
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

func dataSourceTencentCloudTdmqRocketmqRoleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmqRocketmq_role.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("role_name"); ok {
		paramMap["role_name"] = v.(string)
	}

	if v, ok := d.GetOk("cluster_id"); ok {
		paramMap["cluster_id"] = v.(string)
	}

	tdmqRocketmqService := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var roleSets []*tdmqRocketmq.Role
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := tdmqRocketmqService.DescribeTdmqRocketmqRoleByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		roleSets = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read TdmqRocketmq roleSets failed, reason:%+v", logId, err)
		return err
	}

	roleSetList := []interface{}{}
	ids := make([]string, 0)
	for _, roleSet := range roleSets {
		roleSetMap := map[string]interface{}{}
		ids = append(ids, *roleSet.RoleName)
		roleSetMap["role_name"] = roleSet.RoleName
		if roleSet.Token != nil {
			roleSetMap["token"] = roleSet.Token
		}
		if roleSet.Remark != nil {
			roleSetMap["remark"] = roleSet.Remark
		}
		if roleSet.CreateTime != nil {
			roleSetMap["create_time"] = roleSet.CreateTime
		}
		if roleSet.UpdateTime != nil {
			roleSetMap["update_time"] = roleSet.UpdateTime
		}

		roleSetList = append(roleSetList, roleSetMap)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("role_sets", roleSetList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), roleSetList); e != nil {
			return e
		}
	}

	return nil
}

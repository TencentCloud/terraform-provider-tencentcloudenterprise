/*
Use this data source to query detailed information of param records

# Example Usage

```hcl

	data "tencentcloudenterprise_redis_param_records" "param_records" {
		instance_id = "crs-c1nl9rpv"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	redis "terraform-provider-tencentcloudenterprise/sdk/redis/v20180412"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_redis_param_records", CNDescription{
		TerraformTypeCN: "redis®参数历史",
		DescriptionCN:   "提供Redis参数记录数据源，用于查询Redis参数记录的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":            "实例的ID",
			"instance_param_history": "参数名称",
			"param_name":             "参数名称",
			"pre_value":              "修改前的值",
			"new_value":              "修改后的值",
			"status":                 "参数状态：1:参数配置修改；2:参数配置修改成功；3:参数配置修改失败",
			"modify_time":            "修改时间",
			"result_output_file":     "结果输出文件",
		},
	})
}

func dataSourceTencentCloudRedisRecordsParam() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of param records",
		Read:        dataSourceTencentCloudRedisParamRecordsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The ID of instance.",
			},

			"instance_param_history": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "The parameter name.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"param_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The parameter name.",
						},
						"pre_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Modify the previous value.",
						},
						"new_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modified value.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Parameter status:1: parameter configuration modification.2: The parameter configuration is modified successfully.3: Parameter configuration modification failed.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Modification time.",
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

func dataSourceTencentCloudRedisParamRecordsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_redis_param_records.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	var instanceId string

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		paramMap["InstanceId"] = helper.String(v.(string))
	}

	service := RedisService{client: meta.(*TencentCloudClient).apiV3Conn}

	var instanceParamHistory []*redis.InstanceParamHistory

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeRedisParamRecordsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		instanceParamHistory = result
		return nil
	})
	if err != nil {
		return err
	}

	tmpList := make([]map[string]interface{}, 0, len(instanceParamHistory))

	if instanceParamHistory != nil {
		for _, instanceParamHistory := range instanceParamHistory {
			instanceParamHistoryMap := map[string]interface{}{}

			if instanceParamHistory.ParamName != nil {
				instanceParamHistoryMap["param_name"] = instanceParamHistory.ParamName
			}

			if instanceParamHistory.PreValue != nil {
				instanceParamHistoryMap["pre_value"] = instanceParamHistory.PreValue
			}

			if instanceParamHistory.NewValue != nil {
				instanceParamHistoryMap["new_value"] = instanceParamHistory.NewValue
			}

			if instanceParamHistory.Status != nil {
				instanceParamHistoryMap["status"] = instanceParamHistory.Status
			}

			if instanceParamHistory.ModifyTime != nil {
				instanceParamHistoryMap["modify_time"] = instanceParamHistory.ModifyTime
			}

			tmpList = append(tmpList, instanceParamHistoryMap)
		}

		_ = d.Set("instance_param_history", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash([]string{instanceId}))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}

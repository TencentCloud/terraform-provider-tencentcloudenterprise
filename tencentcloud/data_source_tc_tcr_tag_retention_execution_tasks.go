/*
Use this data source to query detailed information of tcr tag_retention_execution_tasks

Example Usage

```hcl
data "tencentcloudenterprise_tcr_tag_retention_execution_tasks" "tasks" {
  registry_id = "%s"
  retention_id = "17"
  execution_id = "1"
  }
```
*/
package tencentcloud

import (
	"context"

	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tcr_tag_retention_execution_tasks", CNDescription{
		TerraformTypeCN: "TCR标签保留执行任务",
		DescriptionCN:   "用于查询TCR标签保留执行任务的详细信息",
		AttributesCN: map[string]string{
			"registry_id":        "实例ID",
			"retention_id":       "保留策略ID",
			"result_output_file": "用于保存结果",
			"retention_task_list": "保留任务列表",
			"task_id":            "任务ID",
			"execution_id":       "执行ID",
			"start_time":         "开始时间",
			"end_time":           "结束时间",
			"status":             "任务状态",
			"total":              "总数",
			"retained":           "保留数量",
			"repository":         "仓库名称",
		},
	})
}

func dataSourceTencentCloudTcrTagRetentionExecutionTasks() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudTcrTagRetentionExecutionTasksRead,
		Description: "Use this data source to query detailed information of TCR tag retention execution tasks",
		Schema: map[string]*schema.Schema{
			"registry_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Instance id.",
			},

			"retention_id": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Retention id.",
			},

			"execution_id": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Execution id.",
			},

			"retention_task_list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "List of version retention tasks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Task id.",
						},
						"execution_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The rule execution id.",
						},
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task start time.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task end time.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The execution status of the task: Failed, Succeed, Stopped, InProgress.",
						},
						"total": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total number of tags.",
						},
						"retained": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total number of retained tags.",
						},
						"repository": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Repository name.",
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

func dataSourceTencentCloudTcrTagRetentionExecutionTasksRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tcr_tag_retention_execution_tasks.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		registryId  string
		retentionId string
		executionId string
	)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("registry_id"); ok {
		paramMap["registry_id"] = helper.String(v.(string))
		registryId = v.(string)
	}

	if v, _ := d.GetOk("retention_id"); v != nil {
		paramMap["retention_id"] = helper.IntInt64(v.(int))
		retentionId = helper.IntToStr(v.(int))
	}

	if v, _ := d.GetOk("execution_id"); v != nil {
		paramMap["execution_id"] = helper.IntInt64(v.(int))
		executionId = helper.IntToStr(v.(int))
	}

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}

	var retentionTaskList []*tcr.RetentionTask

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTcrTagRetentionExecutionTasksByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		retentionTaskList = result
		return nil
	})
	if err != nil {
		return err
	}

	tmpList := make([]map[string]interface{}, 0, len(retentionTaskList))

	if retentionTaskList != nil {
		for _, retentionTask := range retentionTaskList {
			retentionTaskMap := map[string]interface{}{}

			if retentionTask.TaskId != nil {
				retentionTaskMap["task_id"] = retentionTask.TaskId
			}

			if retentionTask.ExecutionId != nil {
				retentionTaskMap["execution_id"] = retentionTask.ExecutionId
			}

			if retentionTask.StartTime != nil {
				retentionTaskMap["start_time"] = retentionTask.StartTime
			}

			if retentionTask.EndTime != nil {
				retentionTaskMap["end_time"] = retentionTask.EndTime
			}

			if retentionTask.Status != nil {
				retentionTaskMap["status"] = retentionTask.Status
			}

			if retentionTask.Total != nil {
				retentionTaskMap["total"] = retentionTask.Total
			}

			if retentionTask.Retained != nil {
				retentionTaskMap["retained"] = retentionTask.Retained
			}

			if retentionTask.Repository != nil {
				retentionTaskMap["repository"] = retentionTask.Repository
			}

			tmpList = append(tmpList, retentionTaskMap)
		}

		_ = d.Set("retention_task_list", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash([]string{registryId, retentionId, executionId}))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}

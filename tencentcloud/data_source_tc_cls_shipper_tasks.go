/*
Use this data source to query detailed information of cls shipper_tasks

# Example Usage

```hcl

	data "tencentcloudenterprise_cls_shipper_tasks" "shipper_tasks" {
	  shipper_id = "dbde3c9b-ea16-4032-bc2a-d8fa65567a8e"
	  start_time = 160749910700
	  end_time = 160749910800
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cls_shipper_tasks", CNDescription{
		TerraformTypeCN: "CLS投递任务列表",
		DescriptionCN:   "提供CLS投递任务列表数据源，用于查询指定投递器的任务执行状态和详细信息。",
		AttributesCN: map[string]string{
			"shipper_id":         "投递器ID",
			"start_time":         "开始时间（毫秒）",
			"end_time":           "结束时间（毫秒）",
			"result_output_file": "用于保存结果",
			"tasks":              "投递任务列表",
		},
	})
}

func dataSourceTencentCloudClsShipperTasks() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudClsShipperTasksRead,
		Description: "Provides a list of CLS shipper tasks within the specified time range",
		Schema: map[string]*schema.Schema{
			"shipper_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Shipper id.",
			},

			"start_time": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Start time(ms).",
			},

			"end_time": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "End time(ms).",
			},

			"tasks": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "List of shipper tasks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task id.",
						},
						"shipper_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shipper id.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topic id.",
						},
						"range_start": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Start time of current task (ms).",
						},
						"range_end": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "End time of current task (ms).",
						},
						"start_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Start time(ms).",
						},
						"end_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "End time(ms).",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of current shipper task.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Detail info.",
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

func dataSourceTencentCloudClsShipperTasksRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cls_shipper_tasks.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("shipper_id"); ok {
		paramMap["ShipperId"] = helper.String(v.(string))
	}

	if v, _ := d.GetOk("start_time"); v != nil {
		paramMap["StartTime"] = helper.IntInt64(v.(int))
	}

	if v, _ := d.GetOk("end_time"); v != nil {
		paramMap["EndTime"] = helper.IntInt64(v.(int))
	}

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var tasks []*cls.ShipperTaskInfo

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeClsShipperTasksByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		tasks = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(tasks))
	tmpList := make([]map[string]interface{}, 0, len(tasks))

	if tasks != nil {
		for _, shipperTaskInfo := range tasks {
			shipperTaskInfoMap := map[string]interface{}{}

			if shipperTaskInfo.TaskId != nil {
				shipperTaskInfoMap["task_id"] = shipperTaskInfo.TaskId
			}

			if shipperTaskInfo.ShipperId != nil {
				shipperTaskInfoMap["shipper_id"] = shipperTaskInfo.ShipperId
			}

			if shipperTaskInfo.TopicId != nil {
				shipperTaskInfoMap["topic_id"] = shipperTaskInfo.TopicId
			}

			if shipperTaskInfo.RangeStart != nil {
				shipperTaskInfoMap["range_start"] = shipperTaskInfo.RangeStart
			}

			if shipperTaskInfo.RangeEnd != nil {
				shipperTaskInfoMap["range_end"] = shipperTaskInfo.RangeEnd
			}

			if shipperTaskInfo.StartTime != nil {
				shipperTaskInfoMap["start_time"] = shipperTaskInfo.StartTime
			}

			if shipperTaskInfo.EndTime != nil {
				shipperTaskInfoMap["end_time"] = shipperTaskInfo.EndTime
			}

			if shipperTaskInfo.Status != nil {
				shipperTaskInfoMap["status"] = shipperTaskInfo.Status
			}

			if shipperTaskInfo.Message != nil {
				shipperTaskInfoMap["message"] = shipperTaskInfo.Message
			}

			ids = append(ids, *shipperTaskInfo.ShipperId)
			tmpList = append(tmpList, shipperTaskInfoMap)
		}

		_ = d.Set("tasks", tmpList)
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

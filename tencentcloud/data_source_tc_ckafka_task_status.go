/*
Use this data source to query detailed information of ckafka task_status

# Example Usage

```hcl

	data "tencentcloudenterprise_ckafka_task_status" "task_status" {
	  flow_id = 123456
	}

```
*/
package tencentcloud

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_task_status", CNDescription{
		TerraformTypeCN: "ckafka任务状态",
		DescriptionCN:   "提供CKafka任务状态数据源，用于查询CKafka任务状态的详细信息。",
		AttributesCN: map[string]string{
			"flow_id":            "流Id",
			"result":             "任务结果",
			"status":             "状态",
			"output":             "输出信息",
			"result_output_file": "结果输出文件",
		},
	})
}

func dataSourceTencentCloudCkafkaTaskStatus() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of ckafka task_status.",
		Read:        dataSourceTencentCloudCkafkaTaskStatusRead,
		Schema: map[string]*schema.Schema{
			"flow_id": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "FlowId.",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Result.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Status.",
						},
						"output": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "OutPut Info.",
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

func dataSourceTencentCloudCkafkaTaskStatusRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_task_status.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	flowId := d.Get("flow_id").(int)

	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}

	var result *ckafka.TaskStatusResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		taskStatus, e := service.DescribeCkafkaTaskStatusByFilter(ctx, flowId)
		if e != nil {
			return retryError(e)
		}
		result = taskStatus
		return nil
	})
	if err != nil {
		return err
	}
	taskStatusResponseMapList := make([]interface{}, 0)
	if result != nil {
		taskStatusResponseMap := map[string]interface{}{}
		if result.Status != nil {
			taskStatusResponseMap["status"] = result.Status
		}

		if result.Output != nil {
			taskStatusResponseMap["output"] = result.Output
		}
		taskStatusResponseMapList = append(taskStatusResponseMapList, taskStatusResponseMap)
		_ = d.Set("result", taskStatusResponseMapList)
	}

	d.SetId(strconv.Itoa(flowId))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), taskStatusResponseMapList); e != nil {
			return e
		}
	}
	return nil
}

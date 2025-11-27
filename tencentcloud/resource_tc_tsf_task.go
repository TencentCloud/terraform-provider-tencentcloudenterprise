/*
Provides a resource to create a tsf task

# Example Usage

```hcl

	resource "tencentcloudenterprise_tsf_task" "task" {
	  task_name = "terraform-test"
	  task_content = "/test"
	  execute_type = "unicast"
	  task_type = "java"
	  time_out = 60000
	  group_id = "group-y8pnmoga"
	  task_rule {
		rule_type = "Cron"
		expression = "0 * 1 * * ? "
	  }
	  retry_count = 0
	  retry_interval = 0
	  success_operator = "GTE"
	  success_ratio = "100"
	  advance_settings {
		sub_task_concurrency = 2
	  }
	  task_argument = "a=c"
	}

```

# Import

tsf task can be imported using the id, e.g.

```
terraform import cloud_tsf_task.task task-y37eqq95
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tsf "terraform-provider-tencentcloudenterprise/sdk/tsf/v20180326"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_tsf_task", CNDescription{
		TerraformTypeCN: "创建任务",
		DescriptionCN:   "提供TSF任务资源，用于创建和管理任务。",
		AttributesCN: map[string]string{
			"task_id":              "任务ID",
			"task_name":            "任务名称",
			"task_content":         "任务内容",
			"execute_type":         "执行类型",
			"task_type":            "任务类型",
			"time_out":             "任务超时时间",
			"group_id":             "部署组ID",
			"task_rule":            "触发规则",
			"retry_count":          "重试次数",
			"retry_interval":       "重试间隔",
			"shard_count":          "分片数量",
			"shard_arguments":      "分片参数",
			"success_operator":     "成功率计算方式",
			"success_ratio":        "成功率阈值",
			"advance_settings":     "高级设置",
			"task_argument":        "任务参数",
			"task_state":           "任务状态",
			"belong_flow_ids":      "所属工作流ID",
			"task_log_id":          "任务历史ID",
			"trigger_type":         "触发类型",
			"expression":           "Cron类型规则，Cron表达式",
			"repeat_interval":      "时间间隔，单位为毫秒",
			"rule_type":            "触发规则类型，Cron/Repeat",
			"shard_key":            "分片参数KEY，整数，范围[11000]",
			"shard_value":          "碎片参数VALUE",
			"sub_task_concurrency": "子任务单机并发限制，默认值为2",
		},
	})
}

func resourceTencentCloudTsfTask() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a tsf task",
		Create:      resourceTencentCloudTsfTaskCreate,
		Read:        resourceTencentCloudTsfTaskRead,
		Update:      resourceTencentCloudTsfTaskUpdate,
		Delete:      resourceTencentCloudTsfTaskDelete,
		// Importer: &schema.ResourceImporter{
		// 	State: schema.ImportStatePassthrough,
		// },
		Schema: map[string]*schema.Schema{
			"task_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Task ID.",
			},

			"task_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Task name, task length 64 characters.",
			},

			"task_content": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Task content, length limit 65536 bytes.",
			},

			"execute_type": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Execution type, unicast/broadcast.",
			},

			"task_type": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Task type, java.",
			},

			"time_out": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Task timeout, time unit ms.",
			},

			"group_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Deployment group ID.",
			},

			"task_rule": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Trigger rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Trigger rule type, Cron/Repeat.",
						},
						"expression": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cron type rule, cron expression.",
						},
						"repeat_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Time interval, in milliseconds.",
						},
					},
				},
			},

			"retry_count": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Number of retries, 0 &amp;lt;= RetryCount&amp;lt;= 10.",
			},

			"retry_interval": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Retry interval, 0 &amp;lt;= RetryInterval &amp;lt;= 600000, time unit ms.",
			},

			"shard_count": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeInt,
				Description: "Number of shards.",
			},

			"shard_arguments": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Fragmentation parameters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shard_key": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Sharding parameter KEY, integer, range [1,1000].",
						},
						"shard_value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Shard parameter VALUE.",
						},
					},
				},
			},

			"success_operator": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The operator to judge the success of the task.",
			},

			"success_ratio": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The threshold for judging the success rate of the task, such as 100.",
			},

			"advance_settings": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Advanced settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sub_task_concurrency": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Subtask single-machine concurrency limit, the default value is 2.",
						},
					},
				},
			},

			"task_argument": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Task parameters, the length limit is 10000 characters.",
			},

			"task_state": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Whether to enable the task, ENABLED/DISABLED.",
			},

			"belong_flow_ids": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "ID of the workflow to which it belongs.",
			},

			"task_log_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Task history ID.",
			},

			"trigger_type": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Trigger type.",
			},

			//"program_id_list": {
			//	Optional: true,
			//	Type:     schema.TypeSet,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString,
			//	},
			//	Description: "Program id list.",
			//},
		},
	}
}

func resourceTencentCloudTsfTaskCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_task.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = tsf.NewCreateTaskRequest()
		response = tsf.NewCreateTaskResponse()
		taskId   string
	)
	if v, ok := d.GetOk("task_name"); ok {
		request.TaskName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("task_content"); ok {
		request.TaskContent = helper.String(v.(string))
	}

	if v, ok := d.GetOk("execute_type"); ok {
		request.ExecuteType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("task_type"); ok {
		request.TaskType = helper.String(v.(string))
	}

	if v, _ := d.GetOk("time_out"); v != nil {
		request.TimeOut = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("group_id"); ok {
		request.GroupId = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "task_rule"); ok {
		taskRule := tsf.TaskRule{}
		if v, ok := dMap["rule_type"]; ok {
			taskRule.RuleType = helper.String(v.(string))
		}
		if v, ok := dMap["expression"]; ok {
			taskRule.Expression = helper.String(v.(string))
		}
		if v, ok := dMap["repeat_interval"]; ok {
			taskRule.RepeatInterval = helper.IntUint64(v.(int))
		}
		request.TaskRule = &taskRule
	}

	if v, _ := d.GetOk("retry_count"); v != nil {
		request.RetryCount = helper.IntUint64(v.(int))
	}

	if v, _ := d.GetOk("retry_interval"); v != nil {
		request.RetryInterval = helper.IntUint64(v.(int))
	}

	if v, _ := d.GetOk("shard_count"); v != nil {
		request.ShardCount = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("shard_arguments"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			shardArgument := tsf.ShardArgument{}
			if v, ok := dMap["shard_key"]; ok {
				shardArgument.ShardKey = helper.IntUint64(v.(int))
			}
			if v, ok := dMap["shard_value"]; ok {
				shardArgument.ShardValue = helper.String(v.(string))
			}
			request.ShardArguments = append(request.ShardArguments, &shardArgument)
		}
	}

	if v, ok := d.GetOk("success_operator"); ok {
		request.SuccessOperator = helper.String(v.(string))
	}

	if v, ok := d.GetOk("success_ratio"); ok {
		request.SuccessRatio = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "advance_settings"); ok {
		advanceSettings := tsf.AdvanceSettings{}
		if v, ok := dMap["sub_task_concurrency"]; ok {
			advanceSettings.SubTaskConcurrency = helper.IntInt64(v.(int))
		}
		request.AdvanceSettings = &advanceSettings
	}

	if v, ok := d.GetOk("task_argument"); ok {
		request.TaskArgument = helper.String(v.(string))
	}

	//if v, ok := d.GetOk("program_id_list"); ok {
	//	programIdListSet := v.(*schema.Set).List()
	//	for i := range programIdListSet {
	//		programIdList := programIdListSet[i].(string)
	//		request.ProgramIdList = append(request.ProgramIdList, &programIdList)
	//	}
	//}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().CreateTask(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create tsf task failed, reason:%+v", logId, err)
		return err
	}

	taskId = *response.Response.Result
	d.SetId(taskId)

	return resourceTencentCloudTsfTaskRead(d, meta)
}

func resourceTencentCloudTsfTaskRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_task.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}

	taskId := d.Id()

	task, err := service.DescribeTsfTaskById(ctx, taskId)
	if err != nil {
		return err
	}

	if task == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `TsfTask` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if task.TaskId != nil {
		_ = d.Set("task_id", task.TaskId)
	}

	if task.TaskName != nil {
		_ = d.Set("task_name", task.TaskName)
	}

	if task.TaskContent != nil {
		_ = d.Set("task_content", task.TaskContent)
	}

	if task.ExecuteType != nil {
		_ = d.Set("execute_type", task.ExecuteType)
	}

	if task.TaskType != nil {
		_ = d.Set("task_type", task.TaskType)
	}

	if task.TimeOut != nil {
		_ = d.Set("time_out", task.TimeOut)
	}

	if task.GroupId != nil {
		_ = d.Set("group_id", task.GroupId)
	}

	if task.TaskRule != nil {
		taskRuleMap := map[string]interface{}{}

		if task.TaskRule.RuleType != nil {
			taskRuleMap["rule_type"] = task.TaskRule.RuleType
		}

		if task.TaskRule.Expression != nil {
			taskRuleMap["expression"] = task.TaskRule.Expression
		}

		if task.TaskRule.RepeatInterval != nil {
			taskRuleMap["repeat_interval"] = task.TaskRule.RepeatInterval
		}

		_ = d.Set("task_rule", []interface{}{taskRuleMap})
	}

	if task.RetryCount != nil {
		_ = d.Set("retry_count", task.RetryCount)
	}

	if task.RetryInterval != nil {
		_ = d.Set("retry_interval", task.RetryInterval)
	}

	if task.ShardCount != nil {
		_ = d.Set("shard_count", task.ShardCount)
	}

	if task.ShardArguments != nil {
		shardArgumentsList := []interface{}{}
		for _, shardArguments := range task.ShardArguments {
			shardArgumentsMap := map[string]interface{}{}

			if shardArguments.ShardKey != nil {
				shardArgumentsMap["shard_key"] = shardArguments.ShardKey
			}

			if shardArguments.ShardValue != nil {
				shardArgumentsMap["shard_value"] = shardArguments.ShardValue
			}

			shardArgumentsList = append(shardArgumentsList, shardArgumentsMap)
		}

		_ = d.Set("shard_arguments", shardArgumentsList)

	}

	if task.SuccessOperator != nil {
		_ = d.Set("success_operator", task.SuccessOperator)
	}

	if task.SuccessRatio != nil {
		_ = d.Set("success_ratio", strconv.FormatInt(*task.SuccessRatio, 10))
	}

	if task.AdvanceSettings != nil {
		advanceSettingsMap := map[string]interface{}{}

		if task.AdvanceSettings.SubTaskConcurrency != nil {
			advanceSettingsMap["sub_task_concurrency"] = task.AdvanceSettings.SubTaskConcurrency
		}

		_ = d.Set("advance_settings", []interface{}{advanceSettingsMap})
	}

	if task.TaskArgument != nil {
		_ = d.Set("task_argument", task.TaskArgument)
	}

	if task.TaskState != nil {
		_ = d.Set("task_state", task.TaskState)
	}

	var belongFlowIds []string
	if task.BelongFlowIds != nil && len(task.BelongFlowIds) > 0 {
		for _, v := range task.BelongFlowIds {
			belongFlowIds = append(belongFlowIds, *v)
		}
	}
	_ = d.Set("belong_flow_ids", belongFlowIds)

	if task.TaskLogId != nil {
		_ = d.Set("task_log_id", task.TaskLogId)
	}

	if task.TriggerType != nil {
		_ = d.Set("trigger_type", task.TriggerType)
	}

	// if task.ProgramIdList != nil {
	// 	_ = d.Set("program_id_list", task.ProgramIdList)
	// }

	return nil
}

func resourceTencentCloudTsfTaskUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_task.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := tsf.NewModifyTaskRequest()

	taskId := d.Id()

	request.TaskId = &taskId

	immutableArgs := []string{"task_id", "task_state", "belong_flow_ids", "task_log_id", "trigger_type"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	if d.HasChange("task_name") {
		if v, ok := d.GetOk("task_name"); ok {
			request.TaskName = helper.String(v.(string))
		}
	}

	if d.HasChange("task_content") {
		if v, ok := d.GetOk("task_content"); ok {
			request.TaskContent = helper.String(v.(string))
		}
	}

	if d.HasChange("execute_type") {
		if v, ok := d.GetOk("execute_type"); ok {
			request.ExecuteType = helper.String(v.(string))
		}
	}

	if d.HasChange("task_type") {
		if v, ok := d.GetOk("task_type"); ok {
			request.TaskType = helper.String(v.(string))
		}
	}

	if d.HasChange("time_out") {
		if v, _ := d.GetOk("time_out"); v != nil {
			request.TimeOut = helper.IntUint64(v.(int))
		}
	}

	if d.HasChange("group_id") {
		if v, ok := d.GetOk("group_id"); ok {
			request.GroupId = helper.String(v.(string))
		}
	}

	if d.HasChange("task_rule") {
		if dMap, ok := helper.InterfacesHeadMap(d, "task_rule"); ok {
			taskRule := tsf.TaskRule{}
			if v, ok := dMap["rule_type"]; ok {
				taskRule.RuleType = helper.String(v.(string))
			}
			if v, ok := dMap["expression"]; ok {
				taskRule.Expression = helper.String(v.(string))
			}
			if v, ok := dMap["repeat_interval"]; ok {
				taskRule.RepeatInterval = helper.IntUint64(v.(int))
			}
			request.TaskRule = &taskRule
		}
	}

	if d.HasChange("retry_count") {
		if v, _ := d.GetOk("retry_count"); v != nil {
			request.RetryCount = helper.IntUint64(v.(int))
		}
	}

	if d.HasChange("retry_interval") {
		if v, _ := d.GetOk("retry_interval"); v != nil {
			request.RetryInterval = helper.IntUint64(v.(int))
		}
	}

	if d.HasChange("shard_count") {
		if v, _ := d.GetOk("shard_count"); v != nil {
			request.ShardCount = helper.IntInt64(v.(int))
		}
	}

	if d.HasChange("shard_arguments") {
		if v, ok := d.GetOk("shard_arguments"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				shardArgument := tsf.ShardArgument{}
				if v, ok := dMap["shard_key"]; ok {
					shardArgument.ShardKey = helper.IntUint64(v.(int))
				}
				if v, ok := dMap["shard_value"]; ok {
					shardArgument.ShardValue = helper.String(v.(string))
				}
				request.ShardArguments = append(request.ShardArguments, &shardArgument)
			}
		}
	}

	if d.HasChange("success_operator") {
		if v, ok := d.GetOk("success_operator"); ok {
			request.SuccessOperator = helper.String(v.(string))
		}
	}

	if d.HasChange("success_ratio") {
		if v, ok := d.GetOk("success_ratio"); ok {
			request.SuccessRatio = helper.StrToInt64Point(v.(string))
		}
	}

	if d.HasChange("advance_settings") {
		if dMap, ok := helper.InterfacesHeadMap(d, "advance_settings"); ok {
			advanceSettings := tsf.AdvanceSettings{}
			if v, ok := dMap["sub_task_concurrency"]; ok {
				advanceSettings.SubTaskConcurrency = helper.IntInt64(v.(int))
			}
			request.AdvanceSettings = &advanceSettings
		}
	}

	if d.HasChange("task_argument") {
		if v, ok := d.GetOk("task_argument"); ok {
			request.TaskArgument = helper.String(v.(string))
		}
	}

	//if d.HasChange("program_id_list") {
	//	if v, ok := d.GetOk("program_id_list"); ok {
	//		programIdListSet := v.(*schema.Set).List()
	//		for i := range programIdListSet {
	//			programIdList := programIdListSet[i].(string)
	//			request.ProgramIdList = append(request.ProgramIdList, &programIdList)
	//		}
	//	}
	//}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTsfClient().ModifyTask(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update tsf task failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudTsfTaskRead(d, meta)
}

func resourceTencentCloudTsfTaskDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_tsf_task.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TsfService{client: meta.(*TencentCloudClient).apiV3Conn}
	taskId := d.Id()

	if err := service.DeleteTsfTaskById(ctx, taskId); err != nil {
		return err
	}

	return nil
}

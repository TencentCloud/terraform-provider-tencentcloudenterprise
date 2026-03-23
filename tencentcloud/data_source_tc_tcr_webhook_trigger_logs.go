/*
Use this data source to query detailed information of tencentcloudenterprise_tcr_webhook_trigger_logs

Example Usage

```hcl
data "tencentcloudenterprise_tcr_webhook_trigger_logs" "my_logs" {
  registry_id = local.tcr_id
  namespace = var.tcr_namespace
  trigger_id = var.trigger_id
    tags = {
    "createdBy" = "terraform"
  }
}
```
*/
package tencentcloud

import (
	"context"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tcr "terraform-provider-tencentcloudenterprise/sdk/tcr/v20190924"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tcr_webhook_trigger_logs", CNDescription{
		TerraformTypeCN: "TCR Webhook触发日志",
		DescriptionCN:   "用于查询TCR Webhook触发日志的详细信息",
		AttributesCN: map[string]string{
			"registry_id":        "实例ID",
			"namespace_name":     "命名空间名称",
			"trigger_id":         "触发器ID",
			"result_output_file": "用于保存结果",
			"logs":               "日志列表",
			"id":                 "日志ID",
			"trigger_name":       "触发器名称",
			"invoke_source":      "调用来源",
			"invoke_action":      "调用动作",
			"invoke_time":        "调用时间",
			"invoke_condition":   "调用条件",
			"invoke_params":      "调用参数",
			"invoke_result":      "调用结果",
		},
	})
}

func dataSourceTencentCloudTcrWebhookTriggerLogs() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudTcrWebhookTriggerLogsRead,
		Description: "Use this data source to query detailed information of TCR webhook trigger logs",
		Schema: map[string]*schema.Schema{
			"registry_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Instance Id.",
			},

			"namespace": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Namespace.",
			},

			"trigger_id": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Trigger id.",
			},

			"logs": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Log list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Log id.",
						},
						"trigger_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trigger Id.",
						},
						"event_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Event type.",
						},
						"notify_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Notification type.",
						},
						"detail": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Webhook trigger detail.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status.",
						},
					},
				},
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudTcrWebhookTriggerLogsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tcr_webhook_trigger_logs.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("registry_id"); ok {
		paramMap["registry_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("namespace"); ok {
		paramMap["namespace"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("trigger_id"); ok {
		paramMap["trigger_id"] = helper.IntInt64(v.(int))
	}

	service := TCRService{client: meta.(*TencentCloudClient).apiV3Conn}

	var logs []*tcr.WebhookTriggerLog

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTcrWebhookTriggerLogByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		logs = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(logs))
	tmpList := make([]map[string]interface{}, 0, len(logs))

	if logs != nil {
		for _, webhookTriggerLog := range logs {
			webhookTriggerLogMap := map[string]interface{}{}

			if webhookTriggerLog.Id != nil {
				webhookTriggerLogMap["id"] = webhookTriggerLog.Id
			}

			if webhookTriggerLog.TriggerId != nil {
				webhookTriggerLogMap["trigger_id"] = webhookTriggerLog.TriggerId
			}

			if webhookTriggerLog.EventType != nil {
				webhookTriggerLogMap["event_type"] = webhookTriggerLog.EventType
			}

			if webhookTriggerLog.NotifyType != nil {
				webhookTriggerLogMap["notify_type"] = webhookTriggerLog.NotifyType
			}

			if webhookTriggerLog.Detail != nil {
				webhookTriggerLogMap["detail"] = webhookTriggerLog.Detail
			}

			if webhookTriggerLog.CreationTime != nil {
				webhookTriggerLogMap["creation_time"] = webhookTriggerLog.CreationTime
			}

			if webhookTriggerLog.UpdateTime != nil {
				webhookTriggerLogMap["update_time"] = webhookTriggerLog.UpdateTime
			}

			if webhookTriggerLog.Status != nil {
				webhookTriggerLogMap["status"] = webhookTriggerLog.Status
			}

			ids = append(ids, helper.Int64ToStr(*webhookTriggerLog.Id))
			tmpList = append(tmpList, webhookTriggerLogMap)
		}

		_ = d.Set("logs", tmpList)
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

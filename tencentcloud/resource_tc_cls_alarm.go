/*
Provides a resource to create a cls alarm

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_alarm" "alarm" {
	  name             = "terraform-alarm-test"
	  alarm_notice_ids = [
	    "notice-0850756b-245d-4bc7-bb27-2a58fffc780b",
	  ]
	  alarm_period     = 15
	  condition        = "test"
	  message_template = "{{.Label}}"
	  status           = true
	  tags             = {
	    "createdBy" = "terraform"
	  }
	  trigger_count = 1

	  alarm_targets {
	    end_time_offset   = 0
	    logset_id         = "33aaf0ae-6163-411b-a415-9f27450f68db"
	    number            = 1
	    query             = "status:>500 | select count(*) as errorCounts"
	    start_time_offset = -15
	    topic_id          = "88735a07-bea4-4985-8763-e9deb6da4fad"
	  }

	  analysis {
	    content = "__FILENAME__"
	    name    = "terraform"
	    type    = "field"

	    config_info {
	      key   = "QueryIndex"
	      value = "1"
	    }
	  }

	  monitor_time {
	    time = 1
	    type = "Period"
	  }
	}

```

# Import

cls alarm can be imported using the id, e.g.

```
terraform import cloud_cls_alarm.alarm alarm_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_alarm", CNDescription{
		TerraformTypeCN: "CLS告警策略",
		DescriptionCN:   "提供CLS告警策略资源，用于创建和管理日志服务告警策略。",
		AttributesCN: map[string]string{
			"name":             "告警策略名称",
			"alarm_targets":    "告警对象",
			"monitor_time":     "监控任务运行时间点",
			"condition":        "触发条件",
			"trigger_count":    "触发次数",
			"alarm_period":     "告警间隔",
			"alarm_notice_ids": "告警通知模板ID列表",
			"status":           "是否开启告警策略",
			"message_template": "自定义告警内容",
			"call_back":        "用户回调URL",
			"analysis":         "多维分析",
			"tags":             "标签",
		},
	})
}

func resourceTencentCloudClsAlarm() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsAlarmCreate,
		Read:        resourceTencentCloudClsAlarmRead,
		Update:      resourceTencentCloudClsAlarmUpdate,
		Delete:      resourceTencentCloudClsAlarmDelete,
		Description: "Provides a resource to create and manage CLS alarm policy",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Log alarm name.",
			},

			"alarm_targets": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "List of alarm target.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Topic id.",
						},
						"query": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Query rules.",
						},
						"number": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The number of alarm object.",
						},
						"start_time_offset": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Search start time of offset.",
						},
						"end_time_offset": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Search end time of offset.",
						},
						"logset_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Logset id.",
						},
					},
				},
			},

			"monitor_time": {
				Required:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Monitor task execution time.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Period for periodic execution, Fixed for regular execution.",
						},
						"time": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Time period or point in time.",
						},
					},
				},
			},

			"condition": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Triggering conditions.",
			},

			"trigger_count": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Continuous cycle.",
			},

			"alarm_period": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Alarm repeat cycle.",
			},

			"alarm_notice_ids": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of alarm notice id.",
			},

			"status": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to enable the alarm policy.",
			},

			"message_template": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "User define alarm notice.",
			},

			"call_back": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "User define callback.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"body": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Callback body.",
						},
						"headers": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "Callback headers.",
						},
					},
				},
			},

			"analysis": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Multidimensional analysis.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Analysis name.",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Analysis type.",
						},
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Analysis content.",
						},
						"config_info": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Key.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Value.",
									},
								},
							},
						},
					},
				},
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list.",
			},
		},
	}
}

func resourceTencentCloudClsAlarmCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_alarm.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request  = cls.NewCreateAlarmRequest()
		response = cls.NewCreateAlarmResponse()
		alarmId  string
	)
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("alarm_targets"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			alarmTarget := cls.AlarmTarget{}
			if v, ok := dMap["topic_id"]; ok {
				alarmTarget.TopicId = helper.String(v.(string))
			}
			if v, ok := dMap["query"]; ok {
				alarmTarget.Query = helper.String(v.(string))
			}
			if v, ok := dMap["number"]; ok {
				alarmTarget.Number = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["start_time_offset"]; ok {
				alarmTarget.StartTimeOffset = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["end_time_offset"]; ok {
				alarmTarget.EndTimeOffset = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["logset_id"]; ok {
				alarmTarget.LogsetId = helper.String(v.(string))
			}
			request.AlarmTargets = append(request.AlarmTargets, &alarmTarget)
		}
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "monitor_time"); ok {
		monitorTime := cls.MonitorTime{}
		if v, ok := dMap["type"]; ok {
			monitorTime.Type = helper.String(v.(string))
		}
		if v, ok := dMap["time"]; ok {
			monitorTime.Time = helper.IntInt64(v.(int))
		}
		request.MonitorTime = &monitorTime
	}

	if v, ok := d.GetOk("condition"); ok {
		request.Condition = helper.String(v.(string))
	}

	if v, ok := d.GetOkExists("trigger_count"); ok {
		request.TriggerCount = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("alarm_period"); ok {
		request.AlarmPeriod = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("alarm_notice_ids"); ok {
		alarmNoticeIdsSet := v.(*schema.Set).List()
		for i := range alarmNoticeIdsSet {
			alarmNoticeIds := alarmNoticeIdsSet[i].(string)
			request.AlarmNoticeIds = append(request.AlarmNoticeIds, &alarmNoticeIds)
		}
	}

	if v, ok := d.GetOkExists("status"); ok {
		request.Status = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("message_template"); ok {
		request.MessageTemplate = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "call_back"); ok {
		callBackInfo := cls.CallBackInfo{}
		if v, ok := dMap["body"]; ok {
			callBackInfo.Body = helper.String(v.(string))
		}
		if v, ok := dMap["headers"]; ok {
			headersSet := v.(*schema.Set).List()
			for i := range headersSet {
				headers := headersSet[i].(string)
				callBackInfo.Headers = append(callBackInfo.Headers, &headers)
			}
		}
		request.CallBack = &callBackInfo
	}

	if v, ok := d.GetOk("analysis"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			analysisDimensional := cls.AnalysisDimensional{}
			if v, ok := dMap["name"]; ok {
				analysisDimensional.Name = helper.String(v.(string))
			}
			if v, ok := dMap["type"]; ok {
				analysisDimensional.Type = helper.String(v.(string))
			}
			if v, ok := dMap["content"]; ok {
				analysisDimensional.Content = helper.String(v.(string))
			}
			if v, ok := dMap["config_info"]; ok {
				for _, item := range v.([]interface{}) {
					configInfoMap := item.(map[string]interface{})
					alarmAnalysisConfig := cls.AlarmAnalysisConfig{}
					if v, ok := configInfoMap["key"]; ok {
						alarmAnalysisConfig.Key = helper.String(v.(string))
					}
					if v, ok := configInfoMap["value"]; ok {
						alarmAnalysisConfig.Value = helper.String(v.(string))
					}
					analysisDimensional.ConfigInfo = append(analysisDimensional.ConfigInfo, &alarmAnalysisConfig)
				}
			}
			request.Analysis = append(request.Analysis, &analysisDimensional)
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateAlarm(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cls alarm failed, reason:%+v", logId, err)
		return err
	}

	alarmId = *response.Response.AlarmId
	d.SetId(alarmId)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::cls:%s:uin/:alarm/%s", region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudClsAlarmRead(d, meta)
}

func resourceTencentCloudClsAlarmRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_alarm.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	alarmId := d.Id()

	alarm, err := service.DescribeClsAlarmById(ctx, alarmId)
	if err != nil {
		return err
	}

	if alarm == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ClsAlarm` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if alarm.Name != nil {
		_ = d.Set("name", alarm.Name)
	}

	if alarm.AlarmTargets != nil {
		alarmTargetsList := []interface{}{}
		for _, alarmTarget := range alarm.AlarmTargets {
			alarmTargetsMap := map[string]interface{}{}

			if alarmTarget.TopicId != nil {
				alarmTargetsMap["topic_id"] = alarmTarget.TopicId
			}

			if alarmTarget.Query != nil {
				alarmTargetsMap["query"] = alarmTarget.Query
			}

			if alarmTarget.Number != nil {
				alarmTargetsMap["number"] = alarmTarget.Number
			}

			if alarmTarget.StartTimeOffset != nil {
				alarmTargetsMap["start_time_offset"] = alarmTarget.StartTimeOffset
			}

			if alarmTarget.EndTimeOffset != nil {
				alarmTargetsMap["end_time_offset"] = alarmTarget.EndTimeOffset
			}

			if alarmTarget.LogsetId != nil {
				alarmTargetsMap["logset_id"] = alarmTarget.LogsetId
			}

			alarmTargetsList = append(alarmTargetsList, alarmTargetsMap)
		}

		_ = d.Set("alarm_targets", alarmTargetsList)

	}

	if alarm.MonitorTime != nil {
		monitorTimeMap := map[string]interface{}{}

		if alarm.MonitorTime.Type != nil {
			monitorTimeMap["type"] = alarm.MonitorTime.Type
		}

		if alarm.MonitorTime.Time != nil {
			monitorTimeMap["time"] = alarm.MonitorTime.Time
		}

		_ = d.Set("monitor_time", []interface{}{monitorTimeMap})
	}

	if alarm.Condition != nil {
		_ = d.Set("condition", alarm.Condition)
	}

	if alarm.TriggerCount != nil {
		_ = d.Set("trigger_count", alarm.TriggerCount)
	}

	if alarm.AlarmPeriod != nil {
		_ = d.Set("alarm_period", alarm.AlarmPeriod)
	}

	if alarm.AlarmNoticeIds != nil {
		_ = d.Set("alarm_notice_ids", alarm.AlarmNoticeIds)
	}

	if alarm.Status != nil {
		_ = d.Set("status", alarm.Status)
	}

	if alarm.MessageTemplate != nil {
		_ = d.Set("message_template", alarm.MessageTemplate)
	}

	if alarm.CallBack != nil {
		callBackMap := map[string]interface{}{}

		if alarm.CallBack.Body != nil {
			callBackMap["body"] = alarm.CallBack.Body
		}

		if alarm.CallBack.Headers != nil {
			callBackMap["headers"] = alarm.CallBack.Headers
		}

		_ = d.Set("call_back", []interface{}{callBackMap})
	}

	if alarm.Analysis != nil {
		analysisList := []interface{}{}
		for _, analysis := range alarm.Analysis {
			analysisMap := map[string]interface{}{}

			if analysis.Name != nil {
				analysisMap["name"] = analysis.Name
			}

			if analysis.Type != nil {
				analysisMap["type"] = analysis.Type
			}

			if analysis.Content != nil {
				analysisMap["content"] = analysis.Content
			}

			if analysis.ConfigInfo != nil {
				configInfoList := []interface{}{}
				for _, configInfo := range analysis.ConfigInfo {
					configInfoMap := map[string]interface{}{}

					if configInfo.Key != nil {
						configInfoMap["key"] = configInfo.Key
					}

					if configInfo.Value != nil {
						configInfoMap["value"] = configInfo.Value
					}

					configInfoList = append(configInfoList, configInfoMap)
				}

				analysisMap["config_info"] = configInfoList
			}

			analysisList = append(analysisList, analysisMap)
		}

		_ = d.Set("analysis", analysisList)

	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := &TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "cls", "alarm", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudClsAlarmUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_alarm.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	needChange := false

	request := cls.NewModifyAlarmRequest()

	alarmId := d.Id()

	request.AlarmId = &alarmId

	mutableArgs := []string{
		"name", "alarm_targets", "monitor_time", "condition",
		"trigger_count", "alarm_period", "alarm_notice_ids",
		"status", "message_template", "call_back", "analysis",
	}

	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {

		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}

		if v, ok := d.GetOk("alarm_targets"); ok {
			for _, item := range v.([]interface{}) {
				alarmTarget := cls.AlarmTarget{}
				dMap := item.(map[string]interface{})
				if v, ok := dMap["topic_id"]; ok {
					alarmTarget.TopicId = helper.String(v.(string))
				}
				if v, ok := dMap["query"]; ok {
					alarmTarget.Query = helper.String(v.(string))
				}
				if v, ok := dMap["number"]; ok {
					alarmTarget.Number = helper.IntInt64(v.(int))
				}
				if v, ok := dMap["start_time_offset"]; ok {
					alarmTarget.StartTimeOffset = helper.IntInt64(v.(int))
				}
				if v, ok := dMap["end_time_offset"]; ok {
					alarmTarget.EndTimeOffset = helper.IntInt64(v.(int))
				}
				if v, ok := dMap["logset_id"]; ok {
					alarmTarget.LogsetId = helper.String(v.(string))
				}
				request.AlarmTargets = append(request.AlarmTargets, &alarmTarget)
			}
		}

		if dMap, ok := helper.InterfacesHeadMap(d, "monitor_time"); ok {
			monitorTime := cls.MonitorTime{}
			if v, ok := dMap["type"]; ok {
				monitorTime.Type = helper.String(v.(string))
			}
			if v, ok := dMap["time"]; ok {
				monitorTime.Time = helper.IntInt64(v.(int))
			}
			request.MonitorTime = &monitorTime
		}

		if v, ok := d.GetOk("condition"); ok {
			request.Condition = helper.String(v.(string))
		}

		if v, ok := d.GetOkExists("trigger_count"); ok {
			request.TriggerCount = helper.IntInt64(v.(int))
		}

		if v, ok := d.GetOkExists("alarm_period"); ok {
			request.AlarmPeriod = helper.IntInt64(v.(int))
		}

		if v, ok := d.GetOk("alarm_notice_ids"); ok {
			alarmNoticeIdsSet := v.(*schema.Set).List()
			for i := range alarmNoticeIdsSet {
				alarmNoticeIds := alarmNoticeIdsSet[i].(string)
				request.AlarmNoticeIds = append(request.AlarmNoticeIds, &alarmNoticeIds)
			}
		}

		if v, ok := d.GetOkExists("status"); ok {
			request.Status = helper.Bool(v.(bool))
		}

		if v, ok := d.GetOk("message_template"); ok {
			request.MessageTemplate = helper.String(v.(string))
		}

		if dMap, ok := helper.InterfacesHeadMap(d, "call_back"); ok {
			callBackInfo := cls.CallBackInfo{}
			if v, ok := dMap["body"]; ok {
				callBackInfo.Body = helper.String(v.(string))
			}
			if v, ok := dMap["headers"]; ok {
				headersSet := v.(*schema.Set).List()
				for i := range headersSet {
					headers := headersSet[i].(string)
					callBackInfo.Headers = append(callBackInfo.Headers, &headers)
				}
			}
			request.CallBack = &callBackInfo
		}

		if v, ok := d.GetOk("analysis"); ok {
			for _, item := range v.([]interface{}) {
				analysisDimensional := cls.AnalysisDimensional{}
				dMap := item.(map[string]interface{})
				if v, ok := dMap["name"]; ok {
					analysisDimensional.Name = helper.String(v.(string))
				}
				if v, ok := dMap["type"]; ok {
					analysisDimensional.Type = helper.String(v.(string))
				}
				if v, ok := dMap["content"]; ok {
					analysisDimensional.Content = helper.String(v.(string))
				}
				if v, ok := dMap["config_info"]; ok {
					for _, item := range v.([]interface{}) {
						configInfoMap := item.(map[string]interface{})
						alarmAnalysisConfig := cls.AlarmAnalysisConfig{}
						if v, ok := configInfoMap["key"]; ok {
							alarmAnalysisConfig.Key = helper.String(v.(string))
						}
						if v, ok := configInfoMap["value"]; ok {
							alarmAnalysisConfig.Value = helper.String(v.(string))
						}
						analysisDimensional.ConfigInfo = append(analysisDimensional.ConfigInfo, &alarmAnalysisConfig)
					}
				}
				request.Analysis = append(request.Analysis, &analysisDimensional)
			}
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyAlarm(request)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update cls alarm failed, reason:%+v", logId, err)
			return err
		}
	}

	if d.HasChange("tags") {
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := &TagService{client: tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("cls", "alarm", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudClsAlarmRead(d, meta)
}

func resourceTencentCloudClsAlarmDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_alarm.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	alarmId := d.Id()

	if err := service.DeleteClsAlarmById(ctx, alarmId); err != nil {
		return err
	}

	return nil
}

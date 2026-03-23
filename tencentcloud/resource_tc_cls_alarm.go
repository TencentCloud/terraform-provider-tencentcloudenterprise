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
terraform import tencentcloudenterprise_cls_alarm.alarm alarm_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_alarm", CNDescription{
		TerraformTypeCN: "CLS告警策略",
		DescriptionCN:   "提供CLS告警策略资源，用于创建和管理日志服务告警策略。",
		AttributesCN: map[string]string{
			"name":                         "告警策略名称（或分析名称）",
			"alarm_targets":                "监控对象列表。",
			"monitor_time":                 "监控任务运行时间点。",
			"condition":                    "触发条件 注意:  Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。",
			"condition_interactive_config": "触发条件交互模式配置信息。",
			"trigger_count":                "持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为2000。",
			"alarm_period":                 "告警重复的周期。单位是分钟。取值范围是0~1440。",
			"alarm_notice_ids":             "关联的告警通知模板列表。",
			"status":                       "是否开启告警策略。默认值为true，true表示开启告警策略，false表示关闭告警策略。",
			"message_template":             "用户自定义告警内容",
			"call_back":                    "用户自定义回调",
			"analysis":                     "多维分析",
			"multi_conditions":             "多触发条件。注意: Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。",
			"tags":                         "标签描述列表，通过指定该参数可以同时绑定标签到相应的告警策略。  最大支持10个标签键值对，并且不能有重复的键值对。",
			"group_trigger_status":         "分组触发状态。 默认值false\t",
			"group_trigger_condition":      "分组触发条件。",
			"alarm_level":                  "告警级别 0:警告(Warn); 1:提醒(Info); 2:紧急 (Critical)。 注意:  不填则默认为0。 Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。",
			"monitor_object_type":          "监控对象类型。0:执行语句共用监控对象; 1:每个执行语句单独选择监控对象。 不填则默认为0。 当值为1时，AlarmTargets元素个数不能超过10个，AlarmTargets中的Number必须是从1开始的连续正整数，不能重复。",
			"classifications":              "告警附加分类信息列表。 Classifications元素个数不能超过20个。 Classifications元素的Key不能为空，不能重复，长度不能超过50个字符，符合正则 ^[a-z]([a-z0-9_]{0,49})$。 Classifications元素的Value长度不能超过200个字符。",
			"alarm_template_info":          "告警模板相关信息",
			"tencentcloudenterprise_product":                "产品类型。TKE：容器服务 CLB：负载均衡",
			"template_id":                  "告警模板ID",
			"instance_id":                  "实例ID。当CloudProduct为TKE时，InstanceId为集群ID",
			"extra_data":                   "告警模板额外信息",
			"alarm_id":                     "告警策略ID。",
			"topic_id":                     "日志主题ID",
			"query":                        "查询语句",
			"number":                       "告警对象序号；从1开始递增",
			"start_time_offset":            "查询范围起始时间相对于告警执行时间的偏移，单位为分钟，取值为非正，最大值为0，最小值为-1440",
			"end_time_offset":              "查询范围终止时间相对于告警执行时间的偏移，单位为分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440",
			"logset_id":                    "日志集ID",
			"syntax_rule":                  "检索语法规则，默认值为0。0：Lucene语法，1：CQL语法。详细说明参见检索条件语法规则",
			"type":                         "执行周期， 可选值：Period、Fixed、Cron。  Period：固定频率 Fixed：固定时间 Cron：Cron表达式",
			"time":                         "执行的周期，或者定制执行的时间节点。单位为分钟，取值范围为1~1440。 当type为Period,Fixed时，time字段生效。",
			"body":                         "回调时的Body。 可将各类告警变量放在请求内容中，详见帮助文档。 如下示例： {\"data\":\"data\"}",
			"headers":                      "回调时的HTTP请求头部字段。 例如：下面请求头部字段来告知服务器请求主体的内容类型为JSON。  \"Content-Type: application/json\"",
			"content":                      "分析内容",
			"config_info":                  "多维分析配置。  当Analysis的Type字段为query（自定义）时，支持 { \"Key\": \"SyntaxRule\", // 语法规则 \"Value\": \"1\" //0：Lucene语法 ，1： CQL语法 }  当Analysis的Type字段为field（top5）时, 支持 { \"Key\": \"QueryIndex\", \"Value\": \"-1\" // -1：自定义， 1：执行语句1， 2：执行语句2 },{ \"Key\": \"CustomQuery\", //检索语句。 QueryIndex为-1时有效且必填 \"Value\": \"* | select count(*) as count\" },{ \"Key\": \"SyntaxRule\", // 查不到这个字段也是老语法（Lucene） \"Value\": \"0\"//0:Lucene, 1:CQL }  当Analysis的Type字段为original（原始日志）时, 支持 { \"Key\": \"Fields\", \"Value\": \"SOURCE,HOSTNAME,TIMESTAMP,PKG_LOGID,TAG.pod_ip\" }, { \"Key\": \"QueryIndex\", \"Value\": \"-1\" // -1：自定义， 1：执行语句1， 2：执行语句2 },{ \"Key\": \"CustomQuery\", // //检索语句。 QueryIndex为-1时有效且必填 \"Value\": \"* | select count(*) as count\" },{ \"Key\": \"Format\", //显示形式。1：每条日志一行，2：每条日志每个字段一行 \"Value\": \"2\" }, { \"Key\": \"Limit\", //最大日志条数 \"Value\": \"5\" },{ \"Key\": \"SyntaxRule\", // 查不到这个字段也是老语法 \"Value\": \"0\"//0:Lucene, 1:CQL }",
			"key":                          "键。支持以下key： SyntaxRule：语法规则，value支持 0：Lucene语法；1： CQL语法。 QueryIndex：执行语句序号。value支持 -1：自定义； 1：执行语句1； 2：执行语句2。 CustomQuery：检索语句。 QueryIndex为-1时有效且必填，value示例： \"* | select count(*) as count\"。 Fields：字段。value支持 SOURCE；FILENAME；HOSTNAME；TIMESTAMP；INDEX_STATUS；PKG_LOGID；TOPIC。 Format：显示形式。value支持 1：每条日志一行；2：每条日志每个字段一行。 Limit：最大日志条数。 value示例： 5。",
			"value":                        "值。 键对应值如下： SyntaxRule：语法规则，value支持 0：Lucene语法；1： CQL语法。 QueryIndex：执行语句序号。value支持 -1：自定义； 1：执行语句1； 2：执行语句2。 CustomQuery：检索语句。 QueryIndex为-1时有效且必填，value示例： \"* | select count(*) as count\"。 Fields：字段。value支持 SOURCE；FILENAME；HOSTNAME；TIMESTAMP；INDEX_STATUS；PKG_LOGID；TOPIC。 Format：显示形式。value支持 1：每条日志一行；2：每条日志每个字段一行。 Limit：最大日志条数。 value示例： 5。",
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
			"alarm_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Alarm Policy ID.",
			},

			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Alarm policy name.",
			},

			"alarm_targets": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "List of monitored objects.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Log topic ID.",
						},
						"query": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Query statement.",
						},
						"number": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateIntegerMin(1),
							Description:  "Alarm object serial number. Starts from 1 and increments.",
						},
						"start_time_offset": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateIntegerInRange(-1440, 0),
							Description:  "Offset (minutes) of query start relative to alarm execution time. Value must be <= 0 and >= -1440.",
						},
						"end_time_offset": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateIntegerInRange(-1440, 0),
							Description:  "Offset (minutes) of query end relative to alarm execution time. Value must be <= 0, greater than start offset, and >= -1440.",
						},
						"logset_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Log set ID.",
						},
						"syntax_rule": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateAllowedIntValue([]int{0, 1}),
							Description:  "Retrieval syntax rule. 0: Lucene syntax, 1: CQL syntax. Default 0.",
						},
					},
				},
			},

			"monitor_time": {
				Required:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Monitors task running time point.",
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
				Optional:     true,
				Type:         schema.TypeString,
				ExactlyOneOf: []string{"condition", "multi_conditions"},
				Description:  "Trigger conditions. Condition and AlarmLevel form one configuration set, MultiConditions is another set. The two sets are mutually exclusive.",
			},

			"condition_interactive_config": {
				Optional:      true,
				Type:          schema.TypeString,
				ConflictsWith: []string{"multi_conditions"},
				Description:   "Interactive trigger configuration.",
			},

			"trigger_count": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateIntegerInRange(1, 2000),
				Description:  "Persistence cycle. An alarm is triggered after the condition is met for TriggerCount consecutive cycles (1-2000).",
			},

			"alarm_period": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateIntegerInRange(0, 1440),
				Description:  "Alarm duplication period in minutes. Valid range: 0-1440.",
			},

			"alarm_notice_ids": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "List of associated alarm notification templates.",
			},

			"status": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to enable the alarm policy. Default true (enabled).",
			},

			"message_template": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "User-defined alarm content.",
			},

			"call_back": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "User-defined callback.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"body": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Body sent during callback. Alarm variables can be embedded in the payload.",
						},
						"headers": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "HTTP headers for the callback request (for example, \"Content-Type: application/json\").",
						},
					},
				},
			},

			"analysis": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Multi-dimensional analysis.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Analysis name.",
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateAllowedStringValue([]string{"query", "field", "original"}),
							Description:  "Analysis type. Valid values: query, field, original.",
						},
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Analysis content.",
						},
						"config_info": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Multi-dimensional analysis configuration (for example QueryIndex, CustomQuery, Fields, Format, Limit, SyntaxRule).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Configuration key.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Configuration value.",
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
				Description: "Tag description list. Up to 10 tag key-value pairs without duplicates can be bound to the alarm policy.",
			},

			"group_trigger_status": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Group trigger status. Default false.",
			},

			"group_trigger_condition": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Grouping trigger conditions.",
			},

			"alarm_level": {
				Optional:      true,
				Type:          schema.TypeInt,
				ValidateFunc:  validateAllowedIntValue([]int{0, 1, 2}),
				ConflictsWith: []string{"multi_conditions"},
				Description:   "Alarm level. 0: Warn; 1: Info; 2: Critical. Default 0. Condition and AlarmLevel are mutually exclusive with MultiConditions.",
			},

			"multi_conditions": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Multiple trigger conditions. Mutually exclusive with condition/alarm_level/condition_interactive_config.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"condition": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Trigger condition.",
						},
						"condition_interactive_config": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Interactive trigger configuration.",
						},
						"alarm_level": {
							Type:         schema.TypeInt,
							Optional:     true,
							ValidateFunc: validateAllowedIntValue([]int{0, 1, 2}),
							Description:  "Alarm level. 0: Warn; 1: Info; 2: Critical. Default 0.",
						},
					},
				},
			},

			"monitor_object_type": {
				Optional:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{0, 1}),
				Description:  "Monitoring object type. 0: share the same monitoring object for all statements; 1: each statement selects its own object (max 10 targets, Number must be consecutive positive integers starting from 1).",
			},

			"classifications": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Alarm additional classification information list (max 20 entries, key matches ^[a-z]([a-z0-9_]{0,49})$, value length <= 200).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Classification key.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Classification value.",
						},
					},
				},
			},

			"alarm_template_info": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Alarm template configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tencentcloudenterprise_product": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cloud product type. TKE for Kubernetes, CLB for load balancer.",
						},
						"template_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Alarm template ID.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Instance ID (cluster ID when CloudProduct is TKE).",
						},
						"extra_data": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Extra alarm template data.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudClsAlarmCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_alarm.create")()
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
			if v, ok := dMap["syntax_rule"]; ok {
				alarmTarget.SyntaxRule = helper.IntUint64(v.(int))
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

	if v, ok := d.GetOk("condition_interactive_config"); ok {
		request.ConditionInteractiveConfig = helper.String(v.(string))
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

	if v, ok := d.GetOkExists("group_trigger_status"); ok {
		request.GroupTriggerStatus = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("group_trigger_condition"); ok {
		groupTriggerConditionSet := v.(*schema.Set).List()
		for i := range groupTriggerConditionSet {
			groupTriggerCondition := groupTriggerConditionSet[i].(string)
			request.GroupTriggerCondition = append(request.GroupTriggerCondition, &groupTriggerCondition)
		}
	}

	if v, ok := d.GetOkExists("alarm_level"); ok {
		request.AlarmLevel = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("multi_conditions"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			multiCondition := cls.MultiCondition{}
			if v, ok := dMap["condition"]; ok {
				multiCondition.Condition = helper.String(v.(string))
			}
			if v, ok := dMap["condition_interactive_config"]; ok {
				multiCondition.ConditionInteractiveConfig = helper.String(v.(string))
			}
			if v, ok := dMap["alarm_level"]; ok {
				multiCondition.AlarmLevel = helper.IntUint64(v.(int))
			}
			request.MultiConditions = append(request.MultiConditions, &multiCondition)
		}
	}

	if v, ok := d.GetOkExists("monitor_object_type"); ok {
		request.MonitorObjectType = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("classifications"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			classification := cls.AlarmClassification{}
			if v, ok := dMap["key"]; ok {
				classification.Key = helper.String(v.(string))
			}
			if v, ok := dMap["value"]; ok {
				classification.Value = helper.String(v.(string))
			}
			request.Classifications = append(request.Classifications, &classification)
		}
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "alarm_template_info"); ok {
		alarmTemplateInfo := cls.AlarmTemplateConfig{}
		if v, ok := dMap["tencentcloudenterprise_product"]; ok {
			alarmTemplateInfo.CloudProduct = helper.String(v.(string))
		}
		if v, ok := dMap["template_id"]; ok {
			alarmTemplateInfo.TemplateId = helper.String(v.(string))
		}
		if v, ok := dMap["instance_id"]; ok {
			alarmTemplateInfo.InstanceId = helper.String(v.(string))
		}
		if v, ok := dMap["extra_data"]; ok {
			alarmTemplateInfo.ExtraData = helper.String(v.(string))
		}
		request.AlarmTemplateInfo = &alarmTemplateInfo
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
	defer logElapsed("resource.tencentcloudenterprise_cls_alarm.read")()
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

	hasMultiConditions := alarm.MultiConditions != nil && len(alarm.MultiConditions) > 0

	_ = d.Set("alarm_id", alarmId)

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

			if alarmTarget.SyntaxRule != nil {
				alarmTargetsMap["syntax_rule"] = alarmTarget.SyntaxRule
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

	if !hasMultiConditions && alarm.Condition != nil {
		_ = d.Set("condition", alarm.Condition)
	}

	if !hasMultiConditions && alarm.ConditionInteractiveConfig != nil {
		_ = d.Set("condition_interactive_config", alarm.ConditionInteractiveConfig)
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

	if alarm.GroupTriggerStatus != nil {
		_ = d.Set("group_trigger_status", alarm.GroupTriggerStatus)
	}

	if alarm.GroupTriggerCondition != nil {
		_ = d.Set("group_trigger_condition", alarm.GroupTriggerCondition)
	}

	if !hasMultiConditions && alarm.AlarmLevel != nil {
		_ = d.Set("alarm_level", alarm.AlarmLevel)
	}

	if hasMultiConditions {
		multiConditionsList := []interface{}{}
		for _, multiCondition := range alarm.MultiConditions {
			multiConditionMap := map[string]interface{}{}

			if multiCondition.Condition != nil {
				multiConditionMap["condition"] = multiCondition.Condition
			}

			if multiCondition.ConditionInteractiveConfig != nil {
				multiConditionMap["condition_interactive_config"] = multiCondition.ConditionInteractiveConfig
			}

			if multiCondition.AlarmLevel != nil {
				multiConditionMap["alarm_level"] = multiCondition.AlarmLevel
			}

			multiConditionsList = append(multiConditionsList, multiConditionMap)
		}

		_ = d.Set("multi_conditions", multiConditionsList)
	}

	if alarm.MonitorObjectType != nil {
		_ = d.Set("monitor_object_type", alarm.MonitorObjectType)
	}

	if alarm.Classifications != nil {
		classificationsList := []interface{}{}
		for _, classification := range alarm.Classifications {
			classificationMap := map[string]interface{}{}

			if classification.Key != nil {
				classificationMap["key"] = classification.Key
			}

			if classification.Value != nil {
				classificationMap["value"] = classification.Value
			}

			classificationsList = append(classificationsList, classificationMap)
		}

		_ = d.Set("classifications", classificationsList)
	}

	if alarm.AlarmTemplateInfo != nil {
		alarmTemplateInfoMap := map[string]interface{}{}
		if alarm.AlarmTemplateInfo.CloudProduct != nil {
			alarmTemplateInfoMap["tencentcloudenterprise_product"] = alarm.AlarmTemplateInfo.CloudProduct
		}
		if alarm.AlarmTemplateInfo.TemplateId != nil {
			alarmTemplateInfoMap["template_id"] = alarm.AlarmTemplateInfo.TemplateId
		}
		if alarm.AlarmTemplateInfo.InstanceId != nil {
			alarmTemplateInfoMap["instance_id"] = alarm.AlarmTemplateInfo.InstanceId
		}
		if alarm.AlarmTemplateInfo.ExtraData != nil {
			alarmTemplateInfoMap["extra_data"] = alarm.AlarmTemplateInfo.ExtraData
		}
		_ = d.Set("alarm_template_info", []interface{}{alarmTemplateInfoMap})
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
	defer logElapsed("resource.tencentcloudenterprise_cls_alarm.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	needChange := false

	request := cls.NewModifyAlarmRequest()

	alarmId := d.Id()

	request.AlarmId = &alarmId

	mutableArgs := []string{
		"name", "alarm_targets", "monitor_time", "condition", "condition_interactive_config",
		"trigger_count", "alarm_period", "alarm_notice_ids",
		"status", "message_template", "call_back", "analysis",
		"group_trigger_status", "group_trigger_condition", "alarm_level", "multi_conditions",
		"monitor_object_type", "classifications", "alarm_template_info",
	}

	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}

	if needChange {
		hasMultiConditions := false

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
				if v, ok := dMap["syntax_rule"]; ok {
					alarmTarget.SyntaxRule = helper.IntUint64(v.(int))
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

		if v, ok := d.GetOkExists("group_trigger_status"); ok {
			request.GroupTriggerStatus = helper.Bool(v.(bool))
		}

		if v, ok := d.GetOk("group_trigger_condition"); ok {
			groupTriggerConditionSet := v.(*schema.Set).List()
			for i := range groupTriggerConditionSet {
				groupTriggerCondition := groupTriggerConditionSet[i].(string)
				request.GroupTriggerCondition = append(request.GroupTriggerCondition, &groupTriggerCondition)
			}
		}

		if v, ok := d.GetOk("multi_conditions"); ok {
			hasMultiConditions = true
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				multiCondition := cls.MultiCondition{}
				if v, ok := dMap["condition"]; ok {
					multiCondition.Condition = helper.String(v.(string))
				}
				if v, ok := dMap["condition_interactive_config"]; ok {
					multiCondition.ConditionInteractiveConfig = helper.String(v.(string))
				}
				if v, ok := dMap["alarm_level"]; ok {
					multiCondition.AlarmLevel = helper.IntUint64(v.(int))
				}
				request.MultiConditions = append(request.MultiConditions, &multiCondition)
			}
		}

		if !hasMultiConditions {
			if v, ok := d.GetOk("condition"); ok {
				request.Condition = helper.String(v.(string))
			}

			if v, ok := d.GetOk("condition_interactive_config"); ok {
				request.ConditionInteractiveConfig = helper.String(v.(string))
			}

			if v, ok := d.GetOkExists("alarm_level"); ok {
				request.AlarmLevel = helper.IntUint64(v.(int))
			}
		}

		if v, ok := d.GetOkExists("monitor_object_type"); ok {
			request.MonitorObjectType = helper.IntUint64(v.(int))
		}

		if v, ok := d.GetOk("classifications"); ok {
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				classification := cls.AlarmClassification{}
				if v, ok := dMap["key"]; ok {
					classification.Key = helper.String(v.(string))
				}
				if v, ok := dMap["value"]; ok {
					classification.Value = helper.String(v.(string))
				}
				request.Classifications = append(request.Classifications, &classification)
			}
		}

		if dMap, ok := helper.InterfacesHeadMap(d, "alarm_template_info"); ok {
			alarmTemplateInfo := cls.AlarmTemplateConfig{}
			if v, ok := dMap["tencentcloudenterprise_product"]; ok {
				alarmTemplateInfo.CloudProduct = helper.String(v.(string))
			}
			if v, ok := dMap["template_id"]; ok {
				alarmTemplateInfo.TemplateId = helper.String(v.(string))
			}
			if v, ok := dMap["instance_id"]; ok {
				alarmTemplateInfo.InstanceId = helper.String(v.(string))
			}
			if v, ok := dMap["extra_data"]; ok {
				alarmTemplateInfo.ExtraData = helper.String(v.(string))
			}
			request.AlarmTemplateInfo = &alarmTemplateInfo
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
	defer logElapsed("resource.tencentcloudenterprise_cls_alarm.delete")()
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

/*
Provides a resource to create a CLS alarm notice.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_alarm_notice" "alarm_notice" {
	  name = "terraform-alarm-notice"
	  type = "All"

	  notice_receivers {
	    receiver_type     = "Uin"
	    receiver_ids      = [10000001]
	    receiver_channels = ["Email"]
	    notice_content_id = "Default-zh"
	    start_time        = "00:00:00"
	    end_time          = "23:59:59"
	  }

	  web_callbacks {
	    callback_type = "Http"
	    url           = "https://example.com/callback"
	    method        = "POST"
	  }

	  tags = {
	    "createdBy" = "terraform"
	  }
	}

```

# Advanced Mode Example

```hcl

	resource "tencentcloudenterprise_cls_alarm_notice" "alarm_notice_rule" {
	  name = "terraform-alarm-notice-rule"

	  notice_rules {
	    rule = jsonencode({
	      Value = "AND"
	      Type  = "Operation"
	      Children = [
	        {
	          Type  = "Condition"
	          Value = "NotifyType"
	          Children = [
	            { Value = "In",   Type = "Compare" },
	            { Value = "[1,2]", Type = "Value" }
	          ]
	        }
	      ]
	    })

	    notice_receivers {
	      receiver_type     = "Uin"
	      receiver_ids      = [4364, 3563]
	      receiver_channels = ["Sms", "WeChat"]
	      start_time        = "00:00:00"
	      end_time          = "23:59:59"
	      notice_content_id = "Default-zh"
	    }

	    escalate = false
	    interval = 10
	    type     = 1
	  }

	  deliver_status      = 1
	  alarm_shield_status = 1
	}

```

# Import

cls alarm notice can be imported using the id, e.g.

```
terraform import tencentcloudenterprise_cls_alarm_notice.alarm_notice alarm_notice_id
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_alarm_notice", CNDescription{
		TerraformTypeCN: "CLS告警通知模板",
		DescriptionCN:   "提供CLS告警通知模板资源，用于创建和管理告警通知模板。",
		AttributesCN: map[string]string{
			"alarm_notice_id":     "告警通知模板ID",
			"name":                "告警通知模板名称",
			"type":                "通知类型。Trigger：触发；Recovery：恢复；All：触发与恢复",
			"notice_receivers":    "通知接收者信息",
			"receiver_type":       "接收者类型。Uin：用户ID；Group：用户组ID",
			"receiver_ids":        "接收者ID列表",
			"receiver_channels":   "接收渠道。Email/Sms/WeChat/Phone",
			"notice_content_id":   "通知内容模板ID",
			"start_time":          "允许接收消息的起始时间",
			"end_time":            "允许接收消息的结束时间",
			"index":               "序号（仅出参有效）",
			"web_callbacks":       "回调信息",
			"callback_type":       "回调类型。Http/WeCom/DingTalk/Lark",
			"url":                 "回调地址",
			"web_callback_id":     "集成配置ID",
			"method":              "回调方法",
			"remind_type":         "提醒类型。0不提醒；1指定人；2所有人",
			"mobiles":             "电话列表",
			"user_ids":            "用户ID列表",
			"headers":             "请求头（已废弃）",
			"body":                "请求体（已废弃）",
			"tags":                "标签描述列表",
			"jump_domain":         "查询数据跳转域名",
			"notice_rules":        "高级模式通知规则列表",
			"rule":                "通知规则JSON字符串",
			"escalate":            "告警升级开关",
			"interval":            "告警升级间隔，单位分钟",
			"deliver_status":      "投递日志开关。1关闭；2开启",
			"deliver_config":      "投递日志配置",
			"deliver_flag":        "投递日志标识",
			"alarm_shield_status": "免登录操作告警开关。1关闭；2开启",
			"deliver_err_msg":     "投递失败原因",
			"region":              "投递地域",
			"topic_id":            "投递日志主题ID",
			"scope":               "投递范围。0全部日志；1仅告警触发及恢复日志",
		},
	})
}

func resourceTencentCloudClsAlarmNotice() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsAlarmNoticeCreate,
		Read:        resourceTencentCloudClsAlarmNoticeRead,
		Update:      resourceTencentCloudClsAlarmNoticeUpdate,
		Delete:      resourceTencentCloudClsAlarmNoticeDelete,
		Description: "Provides a resource to create and manage CLS alarm notice.",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"alarm_notice_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Alarm notice ID.",
			},

			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Alarm notice name.",
			},

			"type": {
				Optional: true,
				Type:     schema.TypeString,
				ConflictsWith: []string{
					"notice_rules",
				},
				Description: "Notice type. Valid values: Trigger, Recovery, All.",
			},

			"jump_domain": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Jump domain for query links.",
			},

			"notice_receivers": {
				Optional: true,
				Type:     schema.TypeList,
				ConflictsWith: []string{
					"notice_rules",
				},
				Description: "Notice receivers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"receiver_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Receiver type. Valid values: Uin, Group.",
						},
						"receiver_ids": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Required:    true,
							Description: "Receiver ID list.",
						},
						"receiver_channels": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "Receiver channels. Valid values: Email, Sms, WeChat, Phone.",
						},
						"notice_content_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Notice content template ID.",
						},
						"start_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Start time allowed to receive messages.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "End time allowed to receive messages.",
						},
						"index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Index. Input is invalid, output is valid.",
						},
					},
				},
			},

			"web_callbacks": {
				Optional: true,
				Type:     schema.TypeList,
				ConflictsWith: []string{
					"notice_rules",
				},
				Description: "Callback information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"callback_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Callback type. Valid values: Http, WeCom, DingTalk, Lark.",
						},
						"url": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Callback URL.",
						},
						"web_callback_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Integration configuration ID.",
						},
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Method. Valid values: POST, PUT.",
						},
						"notice_content_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Notice content template ID.",
						},
						"remind_type": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Remind type. 0: Do not remind; 1: Specified person; 2: Everyone.",
						},
						"mobiles": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Telephone list.",
						},
						"user_ids": {
							Type:        schema.TypeSet,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "User ID list.",
						},
						"headers": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Deprecated:  "This parameter is deprecated. Please use `notice_content_id`.",
							Description: "Request headers.",
						},
						"body": {
							Type:        schema.TypeString,
							Optional:    true,
							Deprecated:  "This parameter is deprecated. Please use `notice_content_id`.",
							Description: "Request body.",
						},
						"index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Index. Input is invalid, output is valid.",
						},
					},
				},
			},

			"notice_rules": {
				Optional: true,
				Type:     schema.TypeList,
				ConflictsWith: []string{
					"type",
					"notice_receivers",
					"web_callbacks",
				},
				Description: "Advanced mode notice rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Rule JSON string.",
						},
						"notice_receivers": {
							Optional:    true,
							Type:        schema.TypeList,
							Description: "Notice receivers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"receiver_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Receiver type. Valid values: Uin, Group.",
									},
									"receiver_ids": {
										Type: schema.TypeSet,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
										Required:    true,
										Description: "Receiver ID list.",
									},
									"receiver_channels": {
										Type: schema.TypeSet,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Required:    true,
										Description: "Receiver channels. Valid values: Email, Sms, WeChat, Phone.",
									},
									"notice_content_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Notice content template ID.",
									},
									"start_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Start time allowed to receive messages.",
									},
									"end_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "End time allowed to receive messages.",
									},
									"index": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Index. Input is invalid, output is valid.",
									},
								},
							},
						},
						"web_callbacks": {
							Optional:    true,
							Type:        schema.TypeList,
							Description: "Callback information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"callback_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Callback type. Valid values: Http, WeCom, DingTalk, Lark.",
									},
									"url": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Callback URL.",
									},
									"web_callback_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Integration configuration ID.",
									},
									"method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Method. Valid values: POST, PUT.",
									},
									"notice_content_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Notice content template ID.",
									},
									"remind_type": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Remind type. 0: Do not remind; 1: Specified person; 2: Everyone.",
									},
									"mobiles": {
										Type:        schema.TypeSet,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Telephone list.",
									},
									"user_ids": {
										Type:        schema.TypeSet,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "User ID list.",
									},
									"headers": {
										Type: schema.TypeSet,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Optional:    true,
										Deprecated:  "This parameter is deprecated. Please use `notice_content_id`.",
										Description: "Request headers.",
									},
									"body": {
										Type:        schema.TypeString,
										Optional:    true,
										Deprecated:  "This parameter is deprecated. Please use `notice_content_id`.",
										Description: "Request body.",
									},
									"index": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Index. Input is invalid, output is valid.",
									},
								},
							},
						},
						"escalate": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to enable escalation.",
						},
						"interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Escalation interval in minutes.",
						},
						"type": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Escalation condition type.",
						},
					},
				},
			},

			"deliver_status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Log delivery switch. 1: Off; 2: On.",
			},

			"deliver_config": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Log delivery configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Delivery region.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Delivery topic ID.",
						},
						"scope": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delivery scope. 0: All logs; 1: Only alarm trigger and recovery logs.",
						},
					},
				},
			},

			"deliver_flag": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Log delivery status flag.",
			},

			"alarm_shield_status": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Alarm shield status. 1: Off; 2: On.",
			},

			"deliver_err_msg": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Log delivery error message.",
			},

			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Tag description list.",
			},
		},
	}
}

func resourceTencentCloudClsAlarmNoticeCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_alarm_notice.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cls.NewCreateAlarmNoticeRequest()
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}
	if v, ok := d.GetOk("jump_domain"); ok {
		request.JumpDomain = helper.String(v.(string))
	}
	if v, ok := d.GetOkExists("deliver_status"); ok {
		request.DeliverStatus = helper.IntUint64(v.(int))
	}
	if v, ok := d.GetOk("deliver_config"); ok {
		request.DeliverConfig = expandAlarmNoticeDeliverConfig(v)
	}
	if v, ok := d.GetOkExists("alarm_shield_status"); ok {
		request.AlarmShieldStatus = helper.IntUint64(v.(int))
	}

	var noticeRules []*cls.NoticeRule
	if v, ok := d.GetOk("notice_rules"); ok {
		noticeRules = expandAlarmNoticeRules(v.([]interface{}))
	}
	if len(noticeRules) > 0 {
		request.NoticeRules = noticeRules
	} else {
		if v, ok := d.GetOk("type"); ok {
			request.Type = helper.String(v.(string))
		}
		if v, ok := d.GetOk("notice_receivers"); ok {
			request.NoticeReceivers = expandAlarmNoticeReceivers(v.([]interface{}))
		}
		if v, ok := d.GetOk("web_callbacks"); ok {
			request.WebCallbacks = expandAlarmNoticeWebCallbacks(v.([]interface{}))
		}
	}

	if request.DeliverStatus != nil && *request.DeliverStatus == 2 && request.DeliverConfig == nil {
		return fmt.Errorf("deliver_config is required when deliver_status is 2")
	}
	if request.Type == nil && len(request.NoticeRules) == 0 {
		return fmt.Errorf("either notice_rules or type must be set")
	}

	var response *cls.CreateAlarmNoticeResponse
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateAlarmNotice(request)
		if e != nil {
			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cls alarm notice failed, reason:%+v", logId, err)
		return err
	}

	alarmNoticeId := *response.Response.AlarmNoticeId
	d.SetId(alarmNoticeId)
	_ = d.Set("alarm_notice_id", alarmNoticeId)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	if tags := helper.GetTags(d, "tags"); len(tags) > 0 {
		tagService := TagService{client: meta.(*TencentCloudClient).apiV3Conn}
		region := meta.(*TencentCloudClient).apiV3Conn.Region
		resourceName := fmt.Sprintf("qcs::cls:%s:uin/:alarmNotice/%s", region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, tags, nil); err != nil {
			return err
		}
	}

	return resourceTencentCloudClsAlarmNoticeRead(d, meta)
}

func resourceTencentCloudClsAlarmNoticeRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_alarm_notice.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	alarmNoticeId := d.Id()
	alarmNotice, err := service.DescribeClsAlarmNoticeById(ctx, alarmNoticeId)
	if err != nil {
		return err
	}

	if alarmNotice == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ClsAlarmNotice` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("alarm_notice_id", alarmNoticeId)

	if alarmNotice.Name != nil {
		_ = d.Set("name", alarmNotice.Name)
	}
	if alarmNotice.Type != nil {
		_ = d.Set("type", alarmNotice.Type)
	}
	if alarmNotice.JumpDomain != nil {
		_ = d.Set("jump_domain", alarmNotice.JumpDomain)
	}
	if alarmNotice.NoticeReceivers != nil {
		_ = d.Set("notice_receivers", flattenAlarmNoticeReceivers(alarmNotice.NoticeReceivers))
	}
	if alarmNotice.WebCallbacks != nil {
		_ = d.Set("web_callbacks", flattenAlarmNoticeWebCallbacks(alarmNotice.WebCallbacks))
	}
	if alarmNotice.NoticeRules != nil {
		_ = d.Set("notice_rules", flattenAlarmNoticeRules(alarmNotice.NoticeRules))
	}
	if alarmNotice.DeliverStatus != nil {
		_ = d.Set("deliver_status", int(*alarmNotice.DeliverStatus))
	}
	if alarmNotice.DeliverFlag != nil {
		_ = d.Set("deliver_flag", int(*alarmNotice.DeliverFlag))
	}
	if alarmNotice.AlarmShieldStatus != nil {
		_ = d.Set("alarm_shield_status", int(*alarmNotice.AlarmShieldStatus))
	}
	if alarmNotice.AlarmNoticeDeliverConfig != nil {
		if alarmNotice.AlarmNoticeDeliverConfig.DeliverConfig != nil {
			_ = d.Set("deliver_config", flattenAlarmNoticeDeliverConfig(alarmNotice.AlarmNoticeDeliverConfig.DeliverConfig))
		}
		if alarmNotice.AlarmNoticeDeliverConfig.ErrMsg != nil {
			_ = d.Set("deliver_err_msg", alarmNotice.AlarmNoticeDeliverConfig.ErrMsg)
		}
	}

	tcClient := meta.(*TencentCloudClient).apiV3Conn
	tagService := TagService{client: tcClient}
	tags, err := tagService.DescribeResourceTags(ctx, "cls", "alarmNotice", tcClient.Region, d.Id())
	if err != nil {
		return err
	}
	_ = d.Set("tags", tags)

	return nil
}

func resourceTencentCloudClsAlarmNoticeUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_alarm_notice.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cls.NewModifyAlarmNoticeRequest()
	request.AlarmNoticeId = helper.String(d.Id())

	mutableArgs := []string{
		"name",
		"type",
		"notice_receivers",
		"web_callbacks",
		"notice_rules",
		"jump_domain",
		"deliver_status",
		"deliver_config",
		"alarm_shield_status",
	}
	needChange := false
	for _, v := range mutableArgs {
		if d.HasChange(v) {
			needChange = true
			break
		}
	}
	noticeConfigChanged := d.HasChange("notice_rules") || d.HasChange("type") || d.HasChange("notice_receivers") || d.HasChange("web_callbacks")
	deliverConfigChanged := d.HasChange("deliver_status") || d.HasChange("deliver_config")

	if needChange {
		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}
		if v, ok := d.GetOk("jump_domain"); ok {
			request.JumpDomain = helper.String(v.(string))
		}
		if v, ok := d.GetOkExists("deliver_status"); ok {
			request.DeliverStatus = helper.IntUint64(v.(int))
		}
		if v, ok := d.GetOk("deliver_config"); ok {
			request.DeliverConfig = expandAlarmNoticeDeliverConfig(v)
		}
		if v, ok := d.GetOkExists("alarm_shield_status"); ok {
			request.AlarmShieldStatus = helper.IntUint64(v.(int))
		}

		var noticeRules []*cls.NoticeRule
		if v, ok := d.GetOk("notice_rules"); ok {
			noticeRules = expandAlarmNoticeRules(v.([]interface{}))
		}
		if len(noticeRules) > 0 {
			request.NoticeRules = noticeRules
		} else {
			if v, ok := d.GetOk("type"); ok {
				request.Type = helper.String(v.(string))
			}
			if v, ok := d.GetOk("notice_receivers"); ok {
				request.NoticeReceivers = expandAlarmNoticeReceivers(v.([]interface{}))
			}
			if v, ok := d.GetOk("web_callbacks"); ok {
				request.WebCallbacks = expandAlarmNoticeWebCallbacks(v.([]interface{}))
			}
		}

		if deliverConfigChanged && request.DeliverStatus != nil && *request.DeliverStatus == 2 && request.DeliverConfig == nil {
			return fmt.Errorf("deliver_config is required when deliver_status is 2")
		}
		if noticeConfigChanged && request.Type == nil && len(request.NoticeRules) == 0 {
			return fmt.Errorf("either notice_rules or type must be set")
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyAlarmNotice(request)
			if e != nil {
				return retryError(e)
			}
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update cls alarm notice failed, reason:%+v", logId, err)
			return err
		}
	}

	if d.HasChange("tags") {
		ctx := context.WithValue(context.TODO(), logIdKey, logId)
		tcClient := meta.(*TencentCloudClient).apiV3Conn
		tagService := TagService{client: tcClient}
		oldTags, newTags := d.GetChange("tags")
		replaceTags, deleteTags := diffTags(oldTags.(map[string]interface{}), newTags.(map[string]interface{}))
		resourceName := BuildTagResourceName("cls", "alarmNotice", tcClient.Region, d.Id())
		if err := tagService.ModifyTags(ctx, resourceName, replaceTags, deleteTags); err != nil {
			return err
		}
	}

	return resourceTencentCloudClsAlarmNoticeRead(d, meta)
}

func resourceTencentCloudClsAlarmNoticeDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_cls_alarm_notice.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	alarmNoticeId := d.Id()

	if err := service.DeleteClsAlarmNoticeById(ctx, alarmNoticeId); err != nil {
		return err
	}

	return nil
}

func expandAlarmNoticeReceivers(items []interface{}) []*cls.NoticeReceiver {
	if len(items) == 0 {
		return nil
	}
	results := make([]*cls.NoticeReceiver, 0, len(items))
	for _, item := range items {
		itemMap := item.(map[string]interface{})
		receiver := cls.NoticeReceiver{}
		if v, ok := itemMap["receiver_type"]; ok {
			receiver.ReceiverType = helper.String(v.(string))
		}
		if v, ok := itemMap["receiver_ids"]; ok {
			idsSet := v.(*schema.Set).List()
			for _, id := range idsSet {
				receiver.ReceiverIds = append(receiver.ReceiverIds, helper.IntInt64(id.(int)))
			}
		}
		if v, ok := itemMap["receiver_channels"]; ok {
			channelsSet := v.(*schema.Set).List()
			for _, channel := range channelsSet {
				value := channel.(string)
				receiver.ReceiverChannels = append(receiver.ReceiverChannels, &value)
			}
		}
		if v, ok := itemMap["notice_content_id"].(string); ok && v != "" {
			receiver.NoticeContentId = helper.String(v)
		}
		if v, ok := itemMap["start_time"].(string); ok && v != "" {
			receiver.StartTime = helper.String(v)
		}
		if v, ok := itemMap["end_time"].(string); ok && v != "" {
			receiver.EndTime = helper.String(v)
		}
		if v, ok := itemMap["index"].(int); ok && v != 0 {
			receiver.Index = helper.IntInt64(v)
		}
		results = append(results, &receiver)
	}
	return results
}

func expandAlarmNoticeWebCallbacks(items []interface{}) []*cls.WebCallback {
	if len(items) == 0 {
		return nil
	}
	results := make([]*cls.WebCallback, 0, len(items))
	for _, item := range items {
		itemMap := item.(map[string]interface{})
		webCallback := cls.WebCallback{}
		if v, ok := itemMap["callback_type"]; ok {
			webCallback.CallbackType = helper.String(v.(string))
		}
		if v, ok := itemMap["url"]; ok {
			webCallback.Url = helper.String(v.(string))
		}
		if v, ok := itemMap["web_callback_id"].(string); ok && v != "" {
			webCallback.WebCallbackId = helper.String(v)
		}
		if v, ok := itemMap["method"].(string); ok && v != "" {
			webCallback.Method = helper.String(v)
		}
		if v, ok := itemMap["notice_content_id"].(string); ok && v != "" {
			webCallback.NoticeContentId = helper.String(v)
		}
		if v, ok := itemMap["remind_type"]; ok {
			webCallback.RemindType = helper.IntUint64(v.(int))
		}
		if v, ok := itemMap["mobiles"]; ok {
			mobilesSet := v.(*schema.Set).List()
			for _, mobile := range mobilesSet {
				value := mobile.(string)
				webCallback.Mobiles = append(webCallback.Mobiles, &value)
			}
		}
		if v, ok := itemMap["user_ids"]; ok {
			userIdsSet := v.(*schema.Set).List()
			for _, userId := range userIdsSet {
				value := userId.(string)
				webCallback.UserIds = append(webCallback.UserIds, &value)
			}
		}
		if v, ok := itemMap["headers"]; ok {
			headersSet := v.(*schema.Set).List()
			for _, header := range headersSet {
				value := header.(string)
				webCallback.Headers = append(webCallback.Headers, &value)
			}
		}
		if v, ok := itemMap["body"]; ok {
			webCallback.Body = helper.String(v.(string))
		}
		if v, ok := itemMap["index"].(int); ok && v != 0 {
			webCallback.Index = helper.IntInt64(v)
		}
		results = append(results, &webCallback)
	}
	return results
}

func expandAlarmNoticeRules(items []interface{}) []*cls.NoticeRule {
	if len(items) == 0 {
		return nil
	}
	results := make([]*cls.NoticeRule, 0, len(items))
	for _, item := range items {
		itemMap := item.(map[string]interface{})
		rule := cls.NoticeRule{}
		if v, ok := itemMap["rule"].(string); ok && v != "" {
			rule.Rule = helper.String(v)
		}
		if raw, ok := itemMap["notice_receivers"]; ok && raw != nil {
			if list, ok := raw.([]interface{}); ok {
				rule.NoticeReceivers = expandAlarmNoticeReceivers(list)
			}
		}
		if raw, ok := itemMap["web_callbacks"]; ok && raw != nil {
			if list, ok := raw.([]interface{}); ok {
				rule.WebCallbacks = expandAlarmNoticeWebCallbacks(list)
			}
		}
		if v, ok := itemMap["escalate"]; ok {
			rule.Escalate = helper.Bool(v.(bool))
		}
		if v, ok := itemMap["interval"]; ok {
			rule.Interval = helper.IntUint64(v.(int))
		}
		if v, ok := itemMap["type"]; ok {
			rule.Type = helper.IntUint64(v.(int))
		}
		results = append(results, &rule)
	}
	return results
}

func expandAlarmNoticeDeliverConfig(raw interface{}) *cls.DeliverConfig {
	if raw == nil {
		return nil
	}
	list, ok := raw.([]interface{})
	if !ok || len(list) == 0 || list[0] == nil {
		return nil
	}
	itemMap := list[0].(map[string]interface{})
	config := cls.DeliverConfig{}
	if v, ok := itemMap["region"].(string); ok && v != "" {
		config.Region = helper.String(v)
	}
	if v, ok := itemMap["topic_id"].(string); ok && v != "" {
		config.TopicId = helper.String(v)
	}
	if v, ok := itemMap["scope"]; ok {
		config.Scope = helper.IntUint64(v.(int))
	}
	return &config
}

func flattenAlarmNoticeReceivers(receivers []*cls.NoticeReceiver) []interface{} {
	if len(receivers) == 0 {
		return nil
	}
	results := make([]interface{}, 0, len(receivers))
	for _, receiver := range receivers {
		itemMap := map[string]interface{}{}
		if receiver.ReceiverType != nil {
			itemMap["receiver_type"] = *receiver.ReceiverType
		}
		if receiver.ReceiverIds != nil {
			ids := make([]int, 0, len(receiver.ReceiverIds))
			for _, id := range receiver.ReceiverIds {
				if id != nil {
					ids = append(ids, int(*id))
				}
			}
			itemMap["receiver_ids"] = ids
		}
		if receiver.ReceiverChannels != nil {
			channels := make([]string, 0, len(receiver.ReceiverChannels))
			for _, channel := range receiver.ReceiverChannels {
				if channel != nil {
					channels = append(channels, *channel)
				}
			}
			itemMap["receiver_channels"] = channels
		}
		if receiver.NoticeContentId != nil {
			itemMap["notice_content_id"] = *receiver.NoticeContentId
		}
		if receiver.StartTime != nil {
			itemMap["start_time"] = *receiver.StartTime
		}
		if receiver.EndTime != nil {
			itemMap["end_time"] = *receiver.EndTime
		}
		if receiver.Index != nil {
			itemMap["index"] = int(*receiver.Index)
		}
		results = append(results, itemMap)
	}
	return results
}

func flattenAlarmNoticeRules(rules []*cls.NoticeRule) []interface{} {
	if len(rules) == 0 {
		return nil
	}
	results := make([]interface{}, 0, len(rules))
	for _, rule := range rules {
		itemMap := map[string]interface{}{}
		if rule.Rule != nil {
			itemMap["rule"] = *rule.Rule
		}
		if rule.NoticeReceivers != nil {
			itemMap["notice_receivers"] = flattenAlarmNoticeReceivers(rule.NoticeReceivers)
		}
		if rule.WebCallbacks != nil {
			itemMap["web_callbacks"] = flattenAlarmNoticeWebCallbacks(rule.WebCallbacks)
		}
		if rule.Escalate != nil {
			itemMap["escalate"] = *rule.Escalate
		}
		if rule.Interval != nil {
			itemMap["interval"] = int(*rule.Interval)
		}
		if rule.Type != nil {
			itemMap["type"] = int(*rule.Type)
		}
		results = append(results, itemMap)
	}
	return results
}

func flattenAlarmNoticeDeliverConfig(config *cls.DeliverConfig) []interface{} {
	if config == nil {
		return nil
	}
	itemMap := map[string]interface{}{}
	if config.Region != nil {
		itemMap["region"] = *config.Region
	}
	if config.TopicId != nil {
		itemMap["topic_id"] = *config.TopicId
	}
	if config.Scope != nil {
		itemMap["scope"] = int(*config.Scope)
	}
	return []interface{}{itemMap}
}

func flattenAlarmNoticeWebCallbacks(callbacks []*cls.WebCallback) []interface{} {
	if len(callbacks) == 0 {
		return nil
	}
	results := make([]interface{}, 0, len(callbacks))
	for _, callback := range callbacks {
		itemMap := map[string]interface{}{}
		if callback.CallbackType != nil {
			itemMap["callback_type"] = *callback.CallbackType
		}
		if callback.Url != nil {
			itemMap["url"] = *callback.Url
		}
		if callback.WebCallbackId != nil {
			itemMap["web_callback_id"] = *callback.WebCallbackId
		}
		if callback.Method != nil {
			itemMap["method"] = *callback.Method
		}
		if callback.NoticeContentId != nil {
			itemMap["notice_content_id"] = *callback.NoticeContentId
		}
		if callback.RemindType != nil {
			itemMap["remind_type"] = int(*callback.RemindType)
		}
		if callback.Mobiles != nil {
			mobiles := make([]string, 0, len(callback.Mobiles))
			for _, mobile := range callback.Mobiles {
				if mobile != nil {
					mobiles = append(mobiles, *mobile)
				}
			}
			itemMap["mobiles"] = mobiles
		}
		if callback.UserIds != nil {
			userIds := make([]string, 0, len(callback.UserIds))
			for _, userId := range callback.UserIds {
				if userId != nil {
					userIds = append(userIds, *userId)
				}
			}
			itemMap["user_ids"] = userIds
		}
		if callback.Headers != nil {
			headers := make([]string, 0, len(callback.Headers))
			for _, header := range callback.Headers {
				if header != nil {
					headers = append(headers, *header)
				}
			}
			itemMap["headers"] = headers
		}
		if callback.Body != nil {
			itemMap["body"] = *callback.Body
		}
		if callback.Index != nil {
			itemMap["index"] = int(*callback.Index)
		}
		results = append(results, itemMap)
	}
	return results
}

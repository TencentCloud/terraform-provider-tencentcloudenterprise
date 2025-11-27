/*
Provides a resource to create a cls cos_recharge

~> **NOTE:** This resource can not be deleted if you run `terraform destroy`.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_cos_recharge" "cos_recharge" {
	  bucket        = "cos-lock-1308919341"
	  bucket_region = "ap-guangzhou"
	  log_type      = "minimalist_log"
	  logset_id     = "dd426d1a-95bc-4bca-b8c2-baa169261812"
	  name          = "cos_recharge_for_test"
	  prefix        = "test"
	  topic_id      = "7e34a3a7-635e-4da8-9005-88106c1fde69"

	  extract_rule_info {
	    backtracking            = 0
	    is_gbk                  = 0
	    json_standard           = 0
	    keys                    = []
	    metadata_type           = 0
	    un_match_up_load_switch = false

	    filter_key_regex {
	      key   = "__CONTENT__"
	      regex = "dasd"
	    }
	  }
	}

```

# Import

cls cos_recharge can be imported using the id, e.g.

```
terraform import cloud_cls_cos_recharge.cos_recharge topic_id#cos_recharge_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_cos_recharge", CNDescription{
		TerraformTypeCN: "CLS COS导入",
		DescriptionCN:   "提供CLS COS导入资源，用于从COS导入日志数据到日志服务。",
		AttributesCN: map[string]string{
			"id":                      "COS导入任务ID",
			"topic_id":                "日志主题ID",
			"logset_id":               "日志集ID",
			"name":                    "COS导入任务名称",
			"bucket":                  "COS存储桶",
			"bucket_region":           "COS存储桶地域",
			"prefix":                  "COS对象前缀",
			"log_type":                "日志类型",
			"compress":                "压缩格式，支持无压缩，gzip，lzop，snappy。默认值为无压缩。",
			"extract_rule_info":       "日志提取规则",
			"time_key":                "时间字段的key名字",
			"time_format":             "时间字段的格式",
			"delimiter":               "分隔符类型日志的分隔符",
			"log_regex":               "整条日志匹配规则",
			"begin_regex":             "行首匹配规则",
			"keys":                    "提取的每个字段的key名字",
			"filter_key_regex":        "日志过滤规则列表（旧版），需要过滤日志的key，及其对应的regex。 注意：2.9.3及以上版本LogListener ，建议使用AdvanceFilterRules配置日志过滤规则。",
			"un_match_up_load_switch": "是否上传解析失败日志",
			"un_match_log_key":        "解析失败日志的key",
			"backtracking":            "增量采集模式下的回溯数据量",
			"is_gbk":                  "GBK编码",
			"json_standard":           "是否为标准json",
			"protocol":                "syslog传输协议",
			"address":                 "syslog系统日志采集指定采集地址",
			"parse_protocol":          "解析协议",
			"metadata_type":           "元数据类型。0: 不使用元数据信息；1:使用机器组元数据；2:使用用户自定义元数据；3:使用采集配置路径。 注意：  COS导入不支持此字段。",
			"path_regex":              "元数据信息中的采集路径正则表达式",
			"meta_tags":               "用户自定义元数据信息。 注意：  MetadataType为2时必填。 COS导入不支持此字段。",
			"event_log_rules":         "Windows事件日志采集规则，只有在LogType为windows_event_log时生效，其余类型无需填写。",
			"event_channel":           "事件通道，支持Application，Security，Setup，System，ALL",
			"time_type":               "时间类型，1:用户自定义，2:当前时间",
			"time_stamp":              "时间，用户选择自定义时间类型时，需要指定时间",
			"event_ids":               "事件ID过滤列表  选填，为空表示不做过滤 支持正向过滤单个值（例：20）或范围（例：0-20），也支持反向过滤单个值(例：-20) 多个过滤项之间可由逗号隔开，例：1-200,-100表示采集1-200范围内除了100以外的事件日志",
			"key":                     "需要过滤日志的key/元数据key",
			"regex":                   "key对应的过滤规则regex",
			"value":                   "元数据value",
		},
	})
}

func resourceTencentCloudClsCosRecharge() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsCosRechargeCreate,
		Read:        resourceTencentCloudClsCosRechargeRead,
		Update:      resourceTencentCloudClsCosRechargeUpdate,
		Delete:      resourceTencentCloudClsCosRechargeDelete,
		Description: "Provides a resource to create and manage CLS COS recharge task for importing log data from COS",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Topic id.",
			},

			"logset_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Logset id.",
			},

			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Recharge name.",
			},

			"bucket": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Cos bucket.",
			},

			"bucket_region": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Cos bucket region.",
			},

			"prefix": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Cos file prefix.",
			},

			"log_type": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Log type.",
			},

			"compress": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Supported gzip, lzop, snappy.",
			},

			"extract_rule_info": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				ForceNew:    true,
				Description: "Extract rule info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time key.",
						},
						"time_format": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time format.",
						},
						"delimiter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Log delimiter.",
						},
						"log_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Log regex.",
						},
						"begin_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "begin line regex.",
						},
						"keys": {
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Optional:    true,
							Description: "Key list.",
						},
						"filter_key_regex": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Rules that need to filter logs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Need filter log key.",
									},
									"regex": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Need filter log regex.",
									},
								},
							},
						},
						"un_match_up_load_switch": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to upload the parsing failure log.",
						},
						"un_match_log_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Parsing failure log key.",
						},
						"backtracking": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "backtracking data volume in incremental acquisition mode.",
						},
						"is_gbk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Gbk encoding.",
						},
						"json_standard": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Is standard json.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Syslog protocol.",
						},
						"address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Syslog address.",
						},
						"parse_protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Parse protocol.",
						},
						"metadata_type": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Metadata type.",
						},
						"path_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Metadata path regex.",
						},
						"meta_tags": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Metadata tag list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Metadata key.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Metadata value.",
									},
								},
							},
						},
						"event_log_rules": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"event_channel": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "event channel, optional value: Application, Security, Setup, System and ALL",
									},
									"time_type": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "time type, 1: user-defined, 2: current time.",
									},
									"time_stamp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "time, When the user selects the custom time type, a specific time needs to be specified.",
									},
									"event_ids": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "event id filter list, Supports positive filtering for single values (e.g.: 20) or ranges (e.g.: 0-20), as well as negative filtering for single values (e.g.: -20). Multiple filter items can be separated by commas. For example: 1-200,-100 means collecting event logs in the range 1-200 excluding 100.",
									},
								},
							},
							Description: "Only take effect when LogType is set to \"windows_event_log\"; no need to fill in for other types.",
						},
					},
				},
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "COS recharge task ID.",
			},
		},
	}
}

func resourceTencentCloudClsCosRechargeCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_recharge.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request    = cls.NewCreateCosRechargeRequest()
		response   = cls.NewCreateCosRechargeResponse()
		topicId    string
		reChargeId string
	)
	if v, ok := d.GetOk("topic_id"); ok {
		topicId = v.(string)
		request.TopicId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("logset_id"); ok {
		request.LogsetId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bucket"); ok {
		request.Bucket = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bucket_region"); ok {
		request.BucketRegion = helper.String(v.(string))
	}

	if v, ok := d.GetOk("prefix"); ok {
		request.Prefix = helper.String(v.(string))
	}

	if v, ok := d.GetOk("log_type"); ok {
		request.LogType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("compress"); ok {
		request.Compress = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "extract_rule_info"); ok {
		extractRuleInfo := cls.ExtractRuleInfo{}
		if v, ok := dMap["time_key"]; ok {
			extractRuleInfo.TimeKey = helper.String(v.(string))
		}
		if v, ok := dMap["time_format"]; ok {
			extractRuleInfo.TimeFormat = helper.String(v.(string))
		}
		if v, ok := dMap["delimiter"]; ok {
			extractRuleInfo.Delimiter = helper.String(v.(string))
		}
		if v, ok := dMap["log_regex"]; ok {
			extractRuleInfo.LogRegex = helper.String(v.(string))
		}
		if v, ok := dMap["begin_regex"]; ok {
			extractRuleInfo.BeginRegex = helper.String(v.(string))
		}
		if v, ok := dMap["keys"]; ok {
			keysList := v.([]interface{})
			for i := range keysList {
				keys := keysList[i].(string)
				extractRuleInfo.Keys = append(extractRuleInfo.Keys, &keys)
			}
		}
		if v, ok := dMap["filter_key_regex"]; ok {
			for _, item := range v.([]interface{}) {
				filterKeyRegexMap := item.(map[string]interface{})
				keyRegexInfo := cls.KeyRegexInfo{}
				if v, ok := filterKeyRegexMap["key"]; ok {
					keyRegexInfo.Key = helper.String(v.(string))
				}
				if v, ok := filterKeyRegexMap["regex"]; ok {
					keyRegexInfo.Regex = helper.String(v.(string))
				}
				extractRuleInfo.FilterKeyRegex = append(extractRuleInfo.FilterKeyRegex, &keyRegexInfo)
			}
		}
		if v, ok := dMap["un_match_up_load_switch"]; ok {
			extractRuleInfo.UnMatchUpLoadSwitch = helper.Bool(v.(bool))
		}
		if v, ok := dMap["un_match_log_key"]; ok {
			extractRuleInfo.UnMatchLogKey = helper.String(v.(string))
		}
		if v, ok := dMap["backtracking"]; ok {
			extractRuleInfo.Backtracking = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["is_gbk"]; ok {
			extractRuleInfo.IsGBK = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["json_standard"]; ok {
			extractRuleInfo.JsonStandard = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["protocol"]; ok {
			extractRuleInfo.Protocol = helper.String(v.(string))
		}
		if v, ok := dMap["address"]; ok {
			extractRuleInfo.Address = helper.String(v.(string))
		}
		if v, ok := dMap["parse_protocol"]; ok {
			extractRuleInfo.ParseProtocol = helper.String(v.(string))
		}
		if v, ok := dMap["metadata_type"]; ok {
			extractRuleInfo.MetadataType = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["path_regex"]; ok {
			extractRuleInfo.PathRegex = helper.String(v.(string))
		}
		if v, ok := dMap["meta_tags"]; ok {
			for _, item := range v.([]interface{}) {
				metaTagsMap := item.(map[string]interface{})
				metaTagInfo := cls.MetaTagInfo{}
				if v, ok := metaTagsMap["key"]; ok {
					metaTagInfo.Key = helper.String(v.(string))
				}
				if v, ok := metaTagsMap["value"]; ok {
					metaTagInfo.Value = helper.String(v.(string))
				}
				extractRuleInfo.MetaTags = append(extractRuleInfo.MetaTags, &metaTagInfo)
			}
		}
		if v, ok := dMap["event_log_rules"]; ok {
			for _, item := range v.([]interface{}) {
				eventLogRulesMap := item.(map[string]interface{})
				eventLogRule := cls.EventLogRuleInfo{}
				if v, ok := eventLogRulesMap["event_channel"]; ok {
					eventLogRule.EventChannel = helper.String(v.(string))
				}
				if v, ok := eventLogRulesMap["time_type"]; ok {
					eventLogRule.TimeType = helper.Uint64(uint64(v.(int)))
				}
				if v, ok := eventLogRulesMap["time_stamp"]; ok {
					eventLogRule.Timestamp = helper.Uint64(uint64(v.(int)))
				}
				if v, ok := eventLogRulesMap["event_ids"]; ok {
					eventIdsList := v.([]interface{})
					for _, eventId := range eventIdsList {
						eventLogRule.EventIDs = append(eventLogRule.EventIDs, helper.String(eventId.(string)))
					}
				}
				extractRuleInfo.EventLogRules = append(extractRuleInfo.EventLogRules, &eventLogRule)
			}
		}
		request.ExtractRuleInfo = &extractRuleInfo
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateCosRecharge(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create cls cosRecharge failed, reason:%+v", logId, err)
		return err
	}

	reChargeId = *response.Response.Id

	d.SetId(topicId + FILED_SP + reChargeId)

	return resourceTencentCloudClsCosRechargeRead(d, meta)
}

func resourceTencentCloudClsCosRechargeRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_recharge.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	topicId := idSplit[0]
	rechargeId := idSplit[1]

	cosRecharge, err := service.DescribeClsCosRechargeById(ctx, topicId, rechargeId)
	if err != nil {
		return err
	}

	if cosRecharge == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ClsCosRecharge` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if cosRecharge.TopicId != nil {
		_ = d.Set("topic_id", cosRecharge.TopicId)
	}

	if cosRecharge.LogsetId != nil {
		_ = d.Set("logset_id", cosRecharge.LogsetId)
	}

	if cosRecharge.Name != nil {
		_ = d.Set("name", cosRecharge.Name)
	}

	if cosRecharge.Bucket != nil {
		_ = d.Set("bucket", cosRecharge.Bucket)
	}

	if cosRecharge.BucketRegion != nil {
		_ = d.Set("bucket_region", cosRecharge.BucketRegion)
	}

	if cosRecharge.Prefix != nil {
		_ = d.Set("prefix", cosRecharge.Prefix)
	}

	if cosRecharge.LogType != nil {
		_ = d.Set("log_type", cosRecharge.LogType)
	}

	if cosRecharge.Compress != nil {
		_ = d.Set("compress", cosRecharge.Compress)
	}

	// Set the recharge ID
	_ = d.Set("id", rechargeId)

	if cosRecharge.ExtractRuleInfo != nil {
		extractRuleInfoMap := map[string]interface{}{}

		if cosRecharge.ExtractRuleInfo.TimeKey != nil {
			extractRuleInfoMap["time_key"] = cosRecharge.ExtractRuleInfo.TimeKey
		}

		if cosRecharge.ExtractRuleInfo.TimeFormat != nil {
			extractRuleInfoMap["time_format"] = cosRecharge.ExtractRuleInfo.TimeFormat
		}

		if cosRecharge.ExtractRuleInfo.Delimiter != nil {
			extractRuleInfoMap["delimiter"] = cosRecharge.ExtractRuleInfo.Delimiter
		}

		if cosRecharge.ExtractRuleInfo.LogRegex != nil {
			extractRuleInfoMap["log_regex"] = cosRecharge.ExtractRuleInfo.LogRegex
		}

		if cosRecharge.ExtractRuleInfo.BeginRegex != nil {
			extractRuleInfoMap["begin_regex"] = cosRecharge.ExtractRuleInfo.BeginRegex
		}

		if cosRecharge.ExtractRuleInfo.Keys != nil {
			extractRuleInfoMap["keys"] = cosRecharge.ExtractRuleInfo.Keys
		}

		if cosRecharge.ExtractRuleInfo.FilterKeyRegex != nil {
			filterKeyRegexList := []interface{}{}
			for _, filterKeyRegex := range cosRecharge.ExtractRuleInfo.FilterKeyRegex {
				filterKeyRegexMap := map[string]interface{}{}

				if filterKeyRegex.Key != nil {
					filterKeyRegexMap["key"] = filterKeyRegex.Key
				}

				if filterKeyRegex.Regex != nil {
					filterKeyRegexMap["regex"] = filterKeyRegex.Regex
				}

				filterKeyRegexList = append(filterKeyRegexList, filterKeyRegexMap)
			}

			extractRuleInfoMap["filter_key_regex"] = filterKeyRegexList
		}

		if cosRecharge.ExtractRuleInfo.UnMatchUpLoadSwitch != nil {
			extractRuleInfoMap["un_match_up_load_switch"] = cosRecharge.ExtractRuleInfo.UnMatchUpLoadSwitch
		}

		if cosRecharge.ExtractRuleInfo.UnMatchLogKey != nil {
			extractRuleInfoMap["un_match_log_key"] = cosRecharge.ExtractRuleInfo.UnMatchLogKey
		}

		if cosRecharge.ExtractRuleInfo.Backtracking != nil {
			extractRuleInfoMap["backtracking"] = cosRecharge.ExtractRuleInfo.Backtracking
		}

		if cosRecharge.ExtractRuleInfo.IsGBK != nil {
			extractRuleInfoMap["is_gbk"] = cosRecharge.ExtractRuleInfo.IsGBK
		}

		if cosRecharge.ExtractRuleInfo.JsonStandard != nil {
			extractRuleInfoMap["json_standard"] = cosRecharge.ExtractRuleInfo.JsonStandard
		}

		if cosRecharge.ExtractRuleInfo.Protocol != nil {
			extractRuleInfoMap["protocol"] = cosRecharge.ExtractRuleInfo.Protocol
		}

		if cosRecharge.ExtractRuleInfo.Address != nil {
			extractRuleInfoMap["address"] = cosRecharge.ExtractRuleInfo.Address
		}

		if cosRecharge.ExtractRuleInfo.ParseProtocol != nil {
			extractRuleInfoMap["parse_protocol"] = cosRecharge.ExtractRuleInfo.ParseProtocol
		}

		if cosRecharge.ExtractRuleInfo.MetadataType != nil {
			extractRuleInfoMap["metadata_type"] = cosRecharge.ExtractRuleInfo.MetadataType
		}

		if cosRecharge.ExtractRuleInfo.PathRegex != nil {
			extractRuleInfoMap["path_regex"] = cosRecharge.ExtractRuleInfo.PathRegex
		}

		if cosRecharge.ExtractRuleInfo.MetaTags != nil {
			metaTagsList := []interface{}{}
			for _, metaTags := range cosRecharge.ExtractRuleInfo.MetaTags {
				metaTagsMap := map[string]interface{}{}

				if metaTags.Key != nil {
					metaTagsMap["key"] = metaTags.Key
				}

				if metaTags.Value != nil {
					metaTagsMap["value"] = metaTags.Value
				}

				metaTagsList = append(metaTagsList, metaTagsMap)
			}

			extractRuleInfoMap["meta_tags"] = metaTagsList
		}

		if cosRecharge.ExtractRuleInfo.EventLogRules != nil {
			eventLogRulesList := []interface{}{}
			for _, eventLogRule := range cosRecharge.ExtractRuleInfo.EventLogRules {
				eventLogRuleMap := map[string]interface{}{}

				if eventLogRule.EventChannel != nil {
					eventLogRuleMap["event_channel"] = eventLogRule.EventChannel
				}

				if eventLogRule.TimeType != nil {
					eventLogRuleMap["time_type"] = int(*eventLogRule.TimeType)
				}

				if eventLogRule.Timestamp != nil {
					eventLogRuleMap["time_stamp"] = int(*eventLogRule.Timestamp)
				}

				if eventLogRule.EventIDs != nil {
					eventIdsList := []interface{}{}
					for _, eventId := range eventLogRule.EventIDs {
						eventIdsList = append(eventIdsList, *eventId)
					}
					eventLogRuleMap["event_ids"] = eventIdsList
				}

				eventLogRulesList = append(eventLogRulesList, eventLogRuleMap)
			}

			extractRuleInfoMap["event_log_rules"] = eventLogRulesList
		}

		_ = d.Set("extract_rule_info", []interface{}{extractRuleInfoMap})
	}

	return nil
}

func resourceTencentCloudClsCosRechargeUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_recharge.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	request := cls.NewModifyCosRechargeRequest()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}
	topicId := idSplit[0]
	id := idSplit[1]

	immutableArgs := []string{
		"logset_id", "bucket", "bucket_region", "prefix",
		"log_type", "compress", "extract_rule_info",
	}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	request.TopicId = &topicId
	request.Id = &id

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyCosRecharge(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s update cls cosRecharge failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudClsCosRechargeRead(d, meta)
}

func resourceTencentCloudClsCosRechargeDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_recharge.delete")()
	defer inconsistentCheck(d, meta)()

	return fmt.Errorf("due to cls api, resource `cls cos_recharge` can not be deleted")
}

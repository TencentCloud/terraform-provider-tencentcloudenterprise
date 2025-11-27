/*
Provides a resource to create a cls config

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_config" "config" {
	  name             = "config_hello"
	  output           = "4d07fba0-b93e-4e0b-9a7f-d58542560bbb"
	  path             = "/var/log/kubernetes"
	  log_type         = "json_log"
	  extract_rule {
	    filter_key_regex {
	      key   = "key1"
	      regex = "value1"
	    }
	    filter_key_regex {
	      key   = "key2"
	      regex = "value2"
	    }
	    un_match_up_load_switch = true
	    un_match_log_key        = "config"
	    backtracking            = -1
	  }
	  exclude_paths {
	    type  = "Path"
	    value = "/data"
	  }
	  exclude_paths {
	    type  = "File"
	    value = "/file"
	  }

#  user_define_rule = ""
}
```

# Import

cls config can be imported using the id, e.g.

```
terraform import cloud_cls_config.config config_id
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_config", CNDescription{
		TerraformTypeCN: "CLS采集配置",
		DescriptionCN:   "提供CLS采集配置资源，用于创建和管理日志服务采集配置。",
		AttributesCN: map[string]string{
			"name":                    "采集配置名称",
			"output":                  "采集配置所属日志主题ID即TopicId",
			"path":                    "日志采集路径，包含文件名，支持多个路径，多个路径之间英文逗号分隔，文件采集情况下必填",
			"log_type":                "采集的日志类型，默认为minimalist_log。支持以下类型：  json_log代表：JSON-文件日志（详见使用 JSON 提取模式采集日志）； delimiter_log代表：分隔符-文件日志（详见使用分隔符提取模式采集日志）； minimalist_log代表：单行全文-文件日志（详见使用单行全文提取模式采集日志）； fullregex_log代表：单行完全正则-文件日志（详见使用单行-完全正则提取模式采集日志）； multiline_log代表：多行全文-文件日志（详见使用多行全文提取模式采集日志）； multiline_fullregex_log代表：多行完全正则-文件日志（详见使用多行-完全正则提取模式采集日志）； user_define_log代表：组合解析（适用于多格式嵌套的日志，详见使用组合解析提取模式采集日志）； service_syslog代表：syslog 采集（详见采集 Syslog）； windows_event_log代表：Windows事件日志（详见采集 Windows 事件日志）。",
			"extract_rule":            "提取规则，如果设置了ExtractRule，则必须设置LogType",
			"exclude_paths":           "采集黑名单路径列表",
			"user_define_rule":        "用户自定义采集规则，Json格式序列化的字符串。当LogType为user_define_log时，必填。",
			"log_format":              "日志格式化方式",
			"config_extra_id":         "config_extra表主键ID",
			"config_flag":             "采集配置标签",
			"advanced_config":         "高级采集配置。 Json字符串， Key/Value定义为如下：  ClsAgentFileTimeout(超时属性), 取值范围: 大于等于0的整数， 0为不超时 ClsAgentMaxDepth(最大目录深度)，取值范围: 大于等于0的整数 ClsAgentParseFailMerge(合并解析失败日志)，取值范围: true或false 样例： {\\\"ClsAgentFileTimeout\\\":0,\\\"ClsAgentMaxDepth\\\":10,\\\"ClsAgentParseFailMerge\\\":true} 控制台默认占位值：{\\\"ClsAgentDefault\\\":0}",
			"source":                  "采集配置来源，0：默认来源，1: TKE",
			"event_log_rules":         "Windows事件日志采集规则，只有在LogType为windows_event_log时生效，其余类型无需填写。",
			"event_channel":           "事件通道，支持Application，Security，Setup，System，ALL",
			"time_type":               "时间类型，1:用户自定义，2:当前时间",
			"time_stamp":              "时间，用户选择自定义时间类型时，需要指定时间",
			"event_ids":               "事件ID过滤列表  选填，为空表示不做过滤 支持正向过滤单个值（例：20）或范围（例：0-20），也支持反向过滤单个值(例：-20) 多个过滤项之间可由逗号隔开，例：1-200,-100表示采集1-200范围内除了100以外的事件日志",
			"config_id":               "采集配置ID",
			"input_type":              "日志输入类型，支持file、window_event、syslog、k8s_stdout、k8s_file",
			"meta_tags":               "机器组元数据信息列表",
			"key":                     "日志过滤字段/元数据key",
			"value":                   "日志过滤值/元数据value",
			"time_key":                "时间字段的key名字，TimeKey和TimeFormat必须成对出现",
			"time_format":             "时间字段的格式，参考c语言的strftime函数对于时间的格式说明输出参数",
			"delimiter":               "分隔符日志的分隔符，仅当log_type为delimiter_log时有效",
			"log_regex":               "完整日志匹配规则，仅当log_type为fullregex_log时有效",
			"begin_regex":             "首行匹配规则，仅当log_type为multiline_log或fullregex_log时有效",
			"keys":                    "每个提取字段的key名字，如果key为空，则表示丢弃该字段，仅当log_type为delimiter_log时有效，json_log日志使用json本身的key",
			"filter_key_regex":        "需要过滤日志的key及其对应的regex",
			"regex":                   "key对应的过滤规则regex",
			"un_match_up_load_switch": "解析失败日志是否上传，true表示上传，false表示不上传",
			"un_match_log_key":        "失败日志的key，当UnMatchUpLoadSwitch为true时必填",
			"backtracking":            "增量采集模式下的回溯数据量，默认：-1（全量采集）；其他非负数表示增量采集（从最新的位置，往前采集${Backtracking}字节（Byte）的日志）最大支持1073741824（1G）。 注意：  COS导入不支持此字段。",
			"is_gbk":                  "是否为Gbk编码。 0：否；1：是。 注意  目前取0值时，表示UTF-8编码 COS导入不支持此字段。",
			"json_standard":           "是否为标准json。 0：否； 1：是。  标准json指采集器使用业界标准开源解析器进行json解析，非标json指采集器使用CLS自研json解析器进行解析，两种解析器没有本质区别，建议客户使用标准json进行解析。",
			"protocol":                "syslog传输协议，取值为tcp或者udp，只有在LogType为service_syslog时生效，其余类型无需填写。 注意：  该字段适用于：创建采集规则配置、修改采集规则配置。 COS导入不支持此字段。",
			"address":                 "syslog系统日志采集指定采集器监听的地址和端口 ，形式：[ip]:[port]，只有在LogType为service_syslog时生效，其余类型无需填写。 注意：  该字段适用于：创建采集规则配置、修改采集规则配置。 COS导入不支持此字段。",
			"parse_protocol":          "rfc3164：指定系统日志采集使用RFC3164协议解析日志。 rfc5424：指定系统日志采集使用RFC5424协议解析日志。 auto：自动匹配rfc3164或者rfc5424其中一种协议。 只有在LogType为service_syslog时生效，其余类型无需填写。 注意：  该字段适用于：创建采集规则配置、修改采集规则配置 COS导入不支持此字段。",
			"metadata_type":           "元数据类型。0: 不使用元数据信息；1:使用机器组元数据；2:使用用户自定义元数据；3:使用采集配置路径。 注意：  COS导入不支持此字段。",
			"path_regex":              "采集配置路径正则表达式。  请用\"()\"标识路径中目标字段对应的正则表达式，解析时将\"()\"视为捕获组，并以__TAG__.{i}:{目标字段}的形式与日志一起上报，其中i为捕获组的序号。 若不希望以序号为键名，可以通过命名捕获组\"(?<{键名}>{正则})\"自定义键名，并以__TAG__.{键名}:{目标字段}的形式与日志一起上报。最多支持5个捕获组 注意：  MetadataType为3时必填。 COS导入不支持此字段。",
			"type":                    "类型，选填File或Path",
			"rule":                    "过滤规则，0:等于，1:字段存在，2:字段不存在",
		},
	})
}

func resourceTencentCloudClsConfig() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsConfigCreate,
		Read:        resourceTencentCloudClsConfigRead,
		Delete:      resourceTencentCloudClsConfigDelete,
		Update:      resourceTencentCloudClsConfigUpdate,
		Description: "Provides a resource to create and manage CLS config",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Collection configuration name.",
			},
			"output": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Log topic ID (TopicId) of collection configuration.",
			},
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Log collection path, including filename. Supports multiple paths separated by English commas. Required for file collection.",
			},
			"log_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of the log to be collected. Default value: minimalist_log. Supported types: json_log: JSON file log; delimiter_log: delimiter file log; minimalist_log: single-line full text file log; fullregex_log: single-line full regex file log; multiline_log: multi-line full text file log; multiline_fullregex_log: multi-line full regex file log; user_define_log: composite parsing (for multi-format nested logs); service_syslog: syslog collection; windows_event_log: Windows event log.",
			},
			"extract_rule": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Extraction rule. If ExtractRule is set, LogType must be set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"event_log_rules": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"event_channel": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Event channel. Supported values: Application, Security, Setup, System, ALL.",
									},
									"time_type": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Time type. 1: user-defined, 2: current time.",
									},
									"time_stamp": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Time. When the user selects the custom time type, a specific time needs to be specified.",
									},
									"event_ids": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "Event ID filter list. Optional, empty means no filtering. Supports positive filtering for single values (e.g.: 20) or ranges (e.g.: 0-20), also supports negative filtering for single values (e.g.: -20). Multiple filter items can be separated by commas. For example: 1-200,-100 means collecting event logs in the range 1-200 excluding 100.",
									},
								},
							},
							Description: "Windows event log collection rules. Only effective when LogType is windows_event_log, no need to fill for other types.",
						},
						"time_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time field key name. time_key and time_format must appear in pair.",
						},
						"time_format": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time field format. For more information, please see the output parameters of the time format description of the strftime function in C language.",
						},
						"delimiter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Delimiter for delimited log, which is valid only if log_type is delimiter_log.",
						},
						"log_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Full log matching rule, which is valid only if log_type is fullregex_log.",
						},
						"begin_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "First-Line matching rule, which is valid only if log_type is multiline_log or fullregex_log.",
						},
						"keys": {
							Type:        schema.TypeList,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if log_type is delimiter_log. json_log logs use the key of JSON itself.",
						},
						"filter_key_regex": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Log keys to be filtered and the corresponding regex.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Log key to be filtered.",
									},
									"regex": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Filter rule regex corresponding to key.",
									},
								},
							},
						},
						"un_match_up_load_switch": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to upload the logs that failed to be parsed. Valid values: true: yes; false: no.",
						},
						"un_match_log_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Unmatched log key.",
						},
						"backtracking": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection).",
						},
						"is_gbk": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "GBK encoding. Default 0.",
						},
						"json_standard": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Standard json. Default 0.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Syslog protocol, tcp or udp.",
						},
						"address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Syslog system log collection specifies the address and port that the collector listens to.",
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
							Description: "List of metadata tags for the machine group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Key of the metadata tag for the machine group.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Value of the metadata tag for the machine group.",
									},
								},
							},
						},
					},
				},
			},
			"exclude_paths": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Collection blacklist path list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type. Valid values: File, Path.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specific content corresponding to Type.",
						},
					},
				},
			},
			"user_define_rule": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User-defined collection rule, JSON format serialized string. Required when LogType is user_define_log.",
			},
			"config_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Collection configuration ID.",
			},
			"log_format": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Log formatting method.",
			},
			"config_extra_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "config_extra table primary key ID.",
			},
			"advanced_config": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Advanced collection configuration. JSON string, Key/Value definitions: ClsAgentFileTimeout (timeout attribute), value range: integer >= 0, 0 means no timeout; ClsAgentMaxDepth (maximum directory depth), value range: integer >= 0; ClsAgentParseFailMerge (merge parse failed logs), value range: true or false. Example: {\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}. Console default placeholder: {\"ClsAgentDefault\":0}.",
			},
			"source": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Collection configuration source. 0: default source, 1: TKE.",
			},
			//"input_type": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Description: "Log input type. Supported values: file, window_event, syslog, k8s_stdout, k8s_file.",
			//},
			"config_flag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Collection configuration tag.",
			},
		},
	}
}

func resourceTencentCloudClsConfigCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_config.create")()

	logId := getLogId(contextNil)

	var (
		request  = cls.NewCreateConfigRequest()
		response *cls.CreateConfigResponse
	)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}
	if v, ok := d.GetOk("output"); ok {
		request.Output = helper.String(v.(string))
	}
	if v, ok := d.GetOk("path"); ok {
		request.Path = helper.String(v.(string))
	}
	if v, ok := d.GetOk("log_type"); ok {
		request.LogType = helper.String(v.(string))
	}
	if v, ok := d.GetOk("extract_rule"); ok {
		var extractRules []*cls.ExtractRuleInfo
		if len(v.([]interface{})) != 1 {
			return fmt.Errorf("need only one extract rule.")
		}
		extractRule := cls.ExtractRuleInfo{}
		dMap := v.([]interface{})[0].(map[string]interface{})
		if v, ok := dMap["time_key"]; ok {
			extractRule.TimeKey = helper.String(v.(string))
		}
		if v, ok := dMap["time_format"]; ok {
			extractRule.TimeFormat = helper.String(v.(string))
		}
		if v, ok := dMap["delimiter"]; ok {
			extractRule.Delimiter = helper.String(v.(string))
		}
		if v, ok := dMap["log_regex"]; ok {
			extractRule.LogRegex = helper.String(v.(string))
		}
		if v, ok := dMap["begin_regex"]; ok {
			extractRule.BeginRegex = helper.String(v.(string))
		}
		if v, ok := dMap["keys"]; ok {
			keys := v.([]interface{})
			for _, key := range keys {
				extractRule.Keys = append(extractRule.Keys, helper.String(key.(string)))
			}
		}
		if v, ok := dMap["filter_key_regex"]; ok {
			keyRegexs := make([]*cls.KeyRegexInfo, 0, 10)
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				keyRegex := cls.KeyRegexInfo{}
				if v, ok := dMap["key"]; ok {
					keyRegex.Key = helper.String(v.(string))
				}
				if v, ok := dMap["regex"]; ok {
					keyRegex.Regex = helper.String(v.(string))
				}
				keyRegexs = append(keyRegexs, &keyRegex)
			}
			extractRule.FilterKeyRegex = keyRegexs
		}
		if v, ok := dMap["un_match_up_load_switch"]; ok {
			extractRule.UnMatchUpLoadSwitch = helper.Bool(v.(bool))
		}
		if v, ok := dMap["un_match_log_key"]; ok {
			extractRule.UnMatchLogKey = helper.String(v.(string))
		}
		if v, ok := dMap["backtracking"]; ok {
			extractRule.Backtracking = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["is_gbk"]; ok {
			extractRule.IsGBK = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["json_standard"]; ok {
			extractRule.JsonStandard = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["protocol"]; ok {
			extractRule.Protocol = helper.String(v.(string))
		}
		if v, ok := dMap["address"]; ok {
			extractRule.Address = helper.String(v.(string))
		}
		if v, ok := dMap["parse_protocol"]; ok {
			extractRule.ParseProtocol = helper.String(v.(string))
		}
		if v, ok := dMap["metadata_type"]; ok {
			extractRule.MetadataType = helper.IntInt64(v.(int))
		}
		if v, ok := dMap["path_regex"]; ok {
			extractRule.PathRegex = helper.String(v.(string))
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
				extractRule.MetaTags = append(extractRule.MetaTags, &metaTagInfo)
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
				extractRule.EventLogRules = append(extractRule.EventLogRules, &eventLogRule)
			}
		}
		extractRules = append(extractRules, &extractRule)
		request.ExtractRule = extractRules[0]
	}
	if v, ok := d.GetOk("exclude_paths"); ok {
		excludePaths := make([]*cls.ExcludePathInfo, 0, 10)
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			excludePath := cls.ExcludePathInfo{}
			if v, ok := dMap["type"]; ok {
				excludePath.Type = helper.String(v.(string))
			}
			if v, ok := dMap["value"]; ok {
				excludePath.Value = helper.String(v.(string))
			}
			excludePaths = append(excludePaths, &excludePath)
		}
		request.ExcludePaths = excludePaths
	}
	if v, ok := d.GetOk("user_define_rule"); ok {
		request.UserDefineRule = helper.String(v.(string))
	}
	if v, ok := d.GetOk("log_format"); ok {
		request.LogFormat = helper.String(v.(string))
	}
	if v, ok := d.GetOk("config_extra_id"); ok {
		request.ConfigExtraId = helper.String(v.(string))
	}
	if v, ok := d.GetOk("advanced_config"); ok {
		request.AdvancedConfig = helper.String(v.(string))
	}
	if v, ok := d.GetOk("source"); ok {
		request.Source = helper.Uint64(uint64(v.(int)))
	}
	//if v, ok := d.GetOk("input_type"); ok {
	//	request.InputType = helper.String(v.(string))
	//}
	if v, ok := d.GetOk("config_flag"); ok {
		request.ConfigFlag = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cls config extra failed, reason:%+v", logId, err)
		return err
	}

	id := *response.Response.ConfigId
	d.SetId(id)
	return resourceTencentCloudClsConfigRead(d, meta)
}

func resourceTencentCloudClsConfigRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_config.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	configId := d.Id()

	var config *cls.ConfigInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeClsConfigById(ctx, configId)
		if e != nil {
			return resource.NonRetryableError(e)
		}

		if result == nil {
			return resource.RetryableError(fmt.Errorf("config %s not found, retrying", configId))
		}

		config = result
		return nil
	})

	if err != nil {
		return err
	}

	if config == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ClsConfig` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if config.Name != nil {
		_ = d.Set("name", config.Name)
	}

	if config.Output != nil {
		_ = d.Set("output", config.Output)
	}

	if config.Path != nil {
		_ = d.Set("path", config.Path)
	}

	if config.LogType != nil {
		_ = d.Set("log_type", config.LogType)
	}

	if config.ExtractRule != nil {
		extractRuleMap := map[string]interface{}{}

		if config.ExtractRule.TimeKey != nil {
			extractRuleMap["time_key"] = config.ExtractRule.TimeKey
		}

		if config.ExtractRule.TimeFormat != nil {
			extractRuleMap["time_format"] = config.ExtractRule.TimeFormat
		}

		if config.ExtractRule.Delimiter != nil {
			extractRuleMap["delimiter"] = config.ExtractRule.Delimiter
		}

		if config.ExtractRule.LogRegex != nil {
			extractRuleMap["log_regex"] = config.ExtractRule.LogRegex
		}

		if config.ExtractRule.BeginRegex != nil {
			extractRuleMap["begin_regex"] = config.ExtractRule.BeginRegex
		}

		if config.ExtractRule.Keys != nil {
			extractRuleMap["keys"] = config.ExtractRule.Keys
		}

		if config.ExtractRule.FilterKeyRegex != nil {
			filterKeyRegexList := []interface{}{}
			for _, filterKeyRegex := range config.ExtractRule.FilterKeyRegex {
				filterKeyRegexMap := map[string]interface{}{}

				if filterKeyRegex.Key != nil {
					filterKeyRegexMap["key"] = filterKeyRegex.Key
				}

				if filterKeyRegex.Regex != nil {
					filterKeyRegexMap["regex"] = filterKeyRegex.Regex
				}

				filterKeyRegexList = append(filterKeyRegexList, filterKeyRegexMap)
			}

			extractRuleMap["filter_key_regex"] = filterKeyRegexList
		}

		if config.ExtractRule.UnMatchUpLoadSwitch != nil {
			extractRuleMap["un_match_up_load_switch"] = config.ExtractRule.UnMatchUpLoadSwitch
		}

		if config.ExtractRule.UnMatchLogKey != nil {
			extractRuleMap["un_match_log_key"] = config.ExtractRule.UnMatchLogKey
		}

		if config.ExtractRule.Backtracking != nil {
			extractRuleMap["backtracking"] = config.ExtractRule.Backtracking
		}

		if config.ExtractRule.IsGBK != nil {
			extractRuleMap["is_gbk"] = config.ExtractRule.IsGBK
		}

		if config.ExtractRule.JsonStandard != nil {
			extractRuleMap["json_standard"] = config.ExtractRule.JsonStandard
		}

		if config.ExtractRule.Protocol != nil {
			extractRuleMap["protocol"] = config.ExtractRule.Protocol
		}

		if config.ExtractRule.Address != nil {
			extractRuleMap["address"] = config.ExtractRule.Address
		}

		if config.ExtractRule.ParseProtocol != nil {
			extractRuleMap["parse_protocol"] = config.ExtractRule.ParseProtocol
		}

		if config.ExtractRule.MetadataType != nil {
			extractRuleMap["metadata_type"] = config.ExtractRule.MetadataType
		}

		if config.ExtractRule.PathRegex != nil {
			extractRuleMap["path_regex"] = config.ExtractRule.PathRegex
		}

		if config.ExtractRule.MetaTags != nil {
			metaTagsList := []interface{}{}
			for _, metaTags := range config.ExtractRule.MetaTags {
				metaTagsMap := map[string]interface{}{}

				if metaTags.Key != nil {
					metaTagsMap["key"] = *metaTags.Key
				}

				if metaTags.Value != nil {
					metaTagsMap["value"] = *metaTags.Value
				}

				metaTagsList = append(metaTagsList, metaTagsMap)
			}
			if metaTagsList != nil {
				extractRuleMap["meta_tags"] = metaTagsList
			}
		}

		if config.ExtractRule.EventLogRules != nil {
			var eventLogRulesList []interface{}
			for _, eventLogRule := range config.ExtractRule.EventLogRules {
				eventLogRuleMap := map[string]interface{}{}

				if eventLogRule.EventChannel != nil {
					eventLogRuleMap["event_channel"] = eventLogRule.EventChannel
				}

				if eventLogRule.TimeType != nil {
					eventLogRuleMap["time_type"] = eventLogRule.TimeType
				}

				if eventLogRule.Timestamp != nil {
					eventLogRuleMap["time_stamp"] = eventLogRule.Timestamp
				}

				if eventLogRule.EventIDs != nil {
					eventLogRuleMap["event_ids"] = eventLogRule.EventIDs
				}

				eventLogRulesList = append(eventLogRulesList, eventLogRuleMap)
			}

			extractRuleMap["event_log_rules"] = eventLogRulesList
		}

		_ = d.Set("extract_rule", []interface{}{extractRuleMap})
	}

	if config.ExcludePaths != nil {
		excludePathsList := []interface{}{}
		for _, excludePath := range config.ExcludePaths {
			excludePathsMap := map[string]interface{}{}

			if excludePath.Type != nil {
				excludePathsMap["type"] = excludePath.Type
			}

			if excludePath.Value != nil {
				excludePathsMap["value"] = excludePath.Value
			}

			excludePathsList = append(excludePathsList, excludePathsMap)
		}

		_ = d.Set("exclude_paths", excludePathsList)
	}

	if config.UserDefineRule != nil {
		_ = d.Set("user_define_rule", config.UserDefineRule)
	}

	if config.LogFormat != nil {
		_ = d.Set("log_format", config.LogFormat)
	}

	if config.ConfigExtraId != nil {
		_ = d.Set("config_extra_id", config.ConfigExtraId)
	}

	if config.AdvancedConfig != nil {
		_ = d.Set("advanced_config", config.AdvancedConfig)
	}

	if config.Source != nil {
		_ = d.Set("source", config.Source)
	}

	//if config.InputType != nil {
	//	_ = d.Set("input_type", config.InputType)
	//}

	if config.ConfigFlag != nil {
		_ = d.Set("config_flag", config.ConfigFlag)
	}

	if config.ConfigId != nil {
		_ = d.Set("config_id", config.ConfigId)
	}

	return nil
}

func resourceTencentCloudClsConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_config.update")()
	logId := getLogId(contextNil)
	request := cls.NewModifyConfigRequest()

	request.ConfigId = helper.String(d.Id())

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok {
			request.Name = helper.String(v.(string))
		}
	}
	if d.HasChange("output") {
		if v, ok := d.GetOk("output"); ok {
			request.Output = helper.String(v.(string))
		}
	}
	if d.HasChange("path") {
		if v, ok := d.GetOk("path"); ok {
			request.Path = helper.String(v.(string))
		}
	}
	if d.HasChange("log_type") || d.HasChange("extract_rule") {
		if v, ok := d.GetOk("log_type"); ok {
			request.LogType = helper.String(v.(string))
		}
		if v, ok := d.GetOk("extract_rule"); ok {
			extractRules := make([]*cls.ExtractRuleInfo, 0, 10)
			if len(v.([]interface{})) != 1 {
				return fmt.Errorf("need only one extract rule.")
			}
			extractRule := cls.ExtractRuleInfo{}
			dMap := v.([]interface{})[0].(map[string]interface{})
			if v, ok := dMap["time_key"]; ok {
				extractRule.TimeKey = helper.String(v.(string))
			}
			if v, ok := dMap["time_format"]; ok {
				extractRule.TimeFormat = helper.String(v.(string))
			}
			if v, ok := dMap["delimiter"]; ok {
				extractRule.Delimiter = helper.String(v.(string))
			}
			if v, ok := dMap["log_regex"]; ok {
				extractRule.LogRegex = helper.String(v.(string))
			}
			if v, ok := dMap["begin_regex"]; ok {
				extractRule.BeginRegex = helper.String(v.(string))
			}
			if v, ok := dMap["keys"]; ok {
				keys := v.([]interface{})
				for _, key := range keys {
					extractRule.Keys = append(extractRule.Keys, helper.String(key.(string)))
				}
			}
			if v, ok := dMap["filter_key_regex"]; ok {
				keyRegexs := make([]*cls.KeyRegexInfo, 0, 10)
				for _, item := range v.([]interface{}) {
					dMap := item.(map[string]interface{})
					keyRegex := cls.KeyRegexInfo{}
					if v, ok := dMap["key"]; ok {
						keyRegex.Key = helper.String(v.(string))
					}
					if v, ok := dMap["regex"]; ok {
						keyRegex.Regex = helper.String(v.(string))
					}
					keyRegexs = append(keyRegexs, &keyRegex)
				}
				extractRule.FilterKeyRegex = keyRegexs
			}
			if v, ok := dMap["un_match_up_load_switch"]; ok {
				extractRule.UnMatchUpLoadSwitch = helper.Bool(v.(bool))
			}
			if v, ok := dMap["un_match_log_key"]; ok {
				extractRule.UnMatchLogKey = helper.String(v.(string))
			}
			if v, ok := dMap["backtracking"]; ok {
				extractRule.Backtracking = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["is_gbk"]; ok {
				extractRule.IsGBK = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["json_standard"]; ok {
				extractRule.JsonStandard = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["protocol"]; ok {
				extractRule.Protocol = helper.String(v.(string))
			}
			if v, ok := dMap["address"]; ok {
				extractRule.Address = helper.String(v.(string))
			}
			if v, ok := dMap["parse_protocol"]; ok {
				extractRule.ParseProtocol = helper.String(v.(string))
			}
			if v, ok := dMap["metadata_type"]; ok {
				extractRule.MetadataType = helper.IntInt64(v.(int))
			}
			if v, ok := dMap["path_regex"]; ok {
				extractRule.PathRegex = helper.String(v.(string))
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
					extractRule.MetaTags = append(extractRule.MetaTags, &metaTagInfo)
				}
			}
			if v, ok := dMap["event_log_rules"]; ok {
				eventLogRules := make([]*cls.EventLogRuleInfo, 0, 10)
				for _, item := range v.([]interface{}) {
					eventLogRuleMap := item.(map[string]interface{})
					eventLogRule := cls.EventLogRuleInfo{}
					if v, ok := eventLogRuleMap["event_channel"]; ok {
						eventLogRule.EventChannel = helper.String(v.(string))
					}
					if v, ok := eventLogRuleMap["time_type"]; ok {
						eventLogRule.TimeType = helper.Uint64(uint64(v.(int)))
					}
					if v, ok := eventLogRuleMap["time_stamp"]; ok {
						eventLogRule.Timestamp = helper.Uint64(uint64(v.(int)))
					}
					if v, ok := eventLogRuleMap["event_ids"]; ok {
						eventIds := v.([]interface{})
						for _, eventId := range eventIds {
							eventLogRule.EventIDs = append(eventLogRule.EventIDs, helper.String(eventId.(string)))
						}
					}
					eventLogRules = append(eventLogRules, &eventLogRule)
				}
				extractRule.EventLogRules = eventLogRules
			}
			extractRules = append(extractRules, &extractRule)
			request.ExtractRule = extractRules[0]
		}
	}
	if d.HasChange("exclude_paths") {
		if v, ok := d.GetOk("exclude_paths"); ok {
			excludePaths := make([]*cls.ExcludePathInfo, 0, 10)
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				excludePath := cls.ExcludePathInfo{}
				if v, ok := dMap["type"]; ok {
					excludePath.Type = helper.String(v.(string))
				}
				if v, ok := dMap["value"]; ok {
					excludePath.Value = helper.String(v.(string))
				}
				excludePaths = append(excludePaths, &excludePath)
			}
			request.ExcludePaths = excludePaths
		}
	}

	if d.HasChange("user_define_rule") {
		if v, ok := d.GetOk("user_define_rule"); ok {
			request.UserDefineRule = helper.String(v.(string))
		}
	}

	if d.HasChange("log_format") {
		if v, ok := d.GetOk("log_format"); ok {
			request.LogFormat = helper.String(v.(string))
		}
	}

	if d.HasChange("config_extra_id") {
		if v, ok := d.GetOk("config_extra_id"); ok {
			request.ConfigExtraId = helper.String(v.(string))
		}
	}

	if d.HasChange("advanced_config") {
		if v, ok := d.GetOk("advanced_config"); ok {
			request.AdvancedConfig = helper.String(v.(string))
		}
	}

	if d.HasChange("source") {
		if v, ok := d.GetOk("source"); ok {
			request.Source = helper.Uint64(uint64(v.(int)))
		}
	}

	if d.HasChange("config_flag") {
		if v, ok := d.GetOk("config_flag"); ok {
			request.ConfigFlag = helper.String(v.(string))
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyConfig(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		return err
	}

	return resourceTencentCloudClsConfigRead(d, meta)
}

func resourceTencentCloudClsConfigDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_config.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	id := d.Id()

	if err := service.DeleteClsConfig(ctx, id); err != nil {
		return err
	}

	return nil
}

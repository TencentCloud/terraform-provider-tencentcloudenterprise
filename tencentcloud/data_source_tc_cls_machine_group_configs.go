/*
Use this data source to query detailed information of cls machine_group_configs

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_machine_group" "group" {
	  group_name        = "tf-describe-mg-config-test"
	  service_logging   = true
	  auto_update       = true
	  update_end_time   = "19:05:00"
	  update_start_time = "17:05:00"

	  machine_group_type {
	    type   = "ip"
	    values = [
	      "203.0.113.101",
	      "203.0.113.102",
	    ]
	  }
	}

	data "tencentcloudenterprise_cls_machine_group_configs" "machine_group_configs" {
	  group_id = cloud_cls_machine_group.group.id
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
	registerDataDescriptionProvider("tencentcloudenterprise_cls_machine_group_configs", CNDescription{
		TerraformTypeCN: "CLS机器组配置列表",
		DescriptionCN:   "提供CLS机器组配置列表数据源，用于查询指定机器组关联的采集配置信息。",
		AttributesCN: map[string]string{
			"result_output_file":      "用于保存查询结果，在线不可用，仅可在本地Provider使用",
			"group_id":                "机器组ID",
			"configs":                 "采集配置列表",
			"config_id":               "采集配置ID",
			"name":                    "采集配置名称",
			"log_format":              "日志格式化方式",
			"path":                    "日志采集路径",
			"log_type":                "采集的日志类型",
			"extract_rule":            "提取规则，如果设置了ExtractRule，则必须设置LogType",
			"exclude_paths":           "采集黑名单路径列表",
			"output":                  "日志主题ID",
			"update_time":             "更新时间",
			"create_time":             "创建时间",
			"user_define_rule":        "用户自定义解析字符串，详见使用组合解析提取模式采集日志。",
			"config_extra_id":         "config_extra主键ID",
			"config_flag":             "采集配置标签",
			"advanced_config":         "日志过滤规则列表（新版）。 注意：  2.9.3以下版本LogListener不支持， 请使用FilterKeyRegex配置日志过滤规则。 自建k8s采集配置（CreateConfigExtra、ModifyConfigExtra）不支持此字段。",
			"source":                  "采集配置来源，0: 默认来源，1: TKE",
			"event_log_rules":         "Windows事件日志采集规则",
			"time_key":                "时间字段的key名字，TimeKey和TimeFormat必须成对出现",
			"time_format":             "时间字段的格式，参考c语言的strftime函数对于时间的格式说明输出参数",
			"delimiter":               "分隔符类型日志的分隔符，只有LogType为delimiter_log时有效",
			"log_regex":               "整条日志匹配规则，只有LogType为fullregex_log时有效",
			"begin_regex":             "行首匹配规则，只有LogType为multiline_log或fullregex_log时有效",
			"keys":                    "取的每个字段的key名字，为空的key代表丢弃这个字段，只有LogType为delimiter_log时有效，json_log的日志使用json本身的key。限制100个。",
			"filter_key_regex":        "日志过滤规则列表（旧版），需要过滤日志的key，及其对应的regex。 注意：2.9.3及以上版本LogListener ，建议使用AdvanceFilterRules配置日志过滤规则。",
			"meta_tags":               "用户自定义元数据信息。 注意：  MetadataType为2时必填。 COS导入不支持此字段。",
			"un_match_up_load_switch": "解析失败日志是否上传，true表示上传，false表示不上传",
			"un_match_log_key":        "失败日志的key，当UnMatchUpLoadSwitch为true时必填",
			"backtracking":            "增量采集模式下的回溯数据量，默认：-1（全量采集）；其他非负数表示增量采集（从最新的位置，往前采集${Backtracking}字节（Byte）的日志）最大支持1073741824（1G）。 注意：  COS导入不支持此字段。",
			"is_gbk":                  "是否为Gbk编码。 0：否；1：是。 注意  目前取0值时，表示UTF-8编码 COS导入不支持此字段。",
			"json_standard":           "是否为标准json。 0：否； 1：是。  标准json指采集器使用业界标准开源解析器进行json解析，非标json指采集器使用CLS自研json解析器进行解析，两种解析器没有本质区别，建议客户使用标准json进行解析。",
			"protocol":                "syslog传输协议，取值为tcp或者udp，只有在LogType为service_syslog时生效，其余类型无需填写。 注意：  该字段适用于：创建采集规则配置、修改采集规则配置。 COS导入不支持此字段。",
			"address":                 "syslog系统日志采集指定采集器监听的地址和端口 ，形式：[ip]:[port]，只有在LogType为service_syslog时生效，其余类型无需填写。 注意：  该字段适用于：创建采集规则配置、修改采集规则配置。 COS导入不支持此字段。",
			"parse_protocol":          "解析协议",
			"metadata_type":           "元数据类型。0: 不使用元数据信息；1:使用机器组元数据；2:使用用户自定义元数据；3:使用采集配置路径。 注意：  COS导入不支持此字段。",
			"path_regex":              "元数据路径匹配规则",
			"type":                    "类型，选填File或Path",
			"value":                   "Type对应的具体内容/元数据标签的Value",
			"event_channel":           "事件通道",
			"time_type":               "时间类型",
			"time_stamp":              "时间戳",
			"event_ids":               "事件ID列表",
			"key":                     "过滤规则的Key/元数据标签的Key",
			"regex":                   "过滤规则的Regex",
		},
	})
}

func dataSourceTencentCloudClsMachineGroupConfigs() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudClsMachineGroupConfigsRead,
		Description: "Retrieves collection configurations associated with a CLS machine group",
		Schema: map[string]*schema.Schema{
			"group_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Group id.",
			},
			"result_output_file": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "path to save result file.",
			},
			"configs": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Collection configuration list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"config_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection configuration ID.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Collection configuration name.",
						},
						"log_format": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Log formatting method.",
						},
						"path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Log collection path, including filename.",
						},
						"log_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the log to be collected. json_log represents JSON format logs, delimiter_log represents delimiter format logs, minimalist_log represents minimalist logs, multiline_log represents multi-line logs, fullregex_log represents full regex, default is minimalist_log.",
						},
						"extract_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Extraction rule. If ExtractRule is set, LogType must be set.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"event_log_rules": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Windows event log collection rules, which only take effect when LogType is windows_event_log.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"event_channel": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Event channel, supports Application, Security, Setup, System, ALL.",
												},
												"time_type": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Time type, 1: user-defined, 2: current time.",
												},
												"time_stamp": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Time, when the user selects a custom time type, a specific time needs to be specified.",
												},
												"event_ids": {
													Type:        schema.TypeList,
													Computed:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "Event ID filter list. An empty list means no filtering. It supports positive filtering of single values or ranges, and also supports negative filtering of single values.",
												},
											},
										},
									},
									"time_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key name of the time field. time_key and time_format must appear in pairs.",
									},
									"time_format": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Time field format. For more information, please see the output parameters of the time format description of the strftime function in C language.",
									},
									"delimiter": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Delimiter for delimited log, which is valid only if log_type is delimiter_log.",
									},
									"log_regex": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Full log matching rule, which is valid only if log_type is fullregex_log.",
									},
									"begin_regex": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "First-Line matching rule, which is valid only if log_type is multiline_log or fullregex_log.",
									},
									"keys": {
										Type: schema.TypeList,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Computed:    true,
										Description: "Key name of each extracted field. An empty key indicates to discard the field. This parameter is valid only if log_type is delimiter_log. json_log logs use the key of JSON itself.",
									},
									"filter_key_regex": {
										Type:        schema.TypeList,
										Computed:    true,
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
										Computed:    true,
										Description: "Whether to upload the logs that failed to be parsed. Valid values: true: yes; false: no.",
									},
									"un_match_log_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Unmatched log key.",
									},
									"backtracking": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Size of the data to be rewound in incremental collection mode. Default value: -1 (full collection).",
									},
									"is_gbk": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "GBK encoding. Default 0.",
									},
									"json_standard": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Standard json. Default 0.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Syslog protocol, tcp or udp.",
									},
									"address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Syslog system log collection specifies the address and port that the collector listens to.",
									},
									"parse_protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parse protocol.",
									},
									"metadata_type": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Metadata type.",
									},
									"path_regex": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Metadata path regex.",
									},
									"meta_tags": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Metadata tags.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Metadata tag key.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Metadata tag value.",
												},
											},
										},
									},
								},
							},
						},
						"exclude_paths": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Collection path blocklist.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Type of blocklist. Valid values: File, Path.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specific content corresponding to Type.",
									},
								},
							},
						},
						"output": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topicid.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"user_define_rule": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User define rule.",
						},
						"config_extra_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Config extra id.",
						},
						"config_flag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Config flag.",
						},
						"advanced_config": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Advanced config.",
						},
						"source": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Config source.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudClsMachineGroupConfigsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cls_machine_group_configs.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("group_id"); ok {
		paramMap["GroupId"] = helper.String(v.(string))
	}

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var configs []*cls.ConfigInfo

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeClsMachineGroupConfigsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		configs = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(configs))
	tmpList := make([]map[string]interface{}, 0, len(configs))

	if configs != nil {
		for _, configInfo := range configs {
			configInfoMap := map[string]interface{}{}

			if configInfo.ConfigId != nil {
				configInfoMap["config_id"] = configInfo.ConfigId
			}

			if configInfo.Name != nil {
				configInfoMap["name"] = configInfo.Name
			}

			if configInfo.LogFormat != nil {
				configInfoMap["log_format"] = configInfo.LogFormat
			}

			if configInfo.Path != nil {
				configInfoMap["path"] = configInfo.Path
			}

			if configInfo.LogType != nil {
				configInfoMap["log_type"] = configInfo.LogType
			}

			if configInfo.ExtractRule != nil {
				extractRuleMap := map[string]interface{}{}

				// 处理EventLogRules
				if configInfo.ExtractRule.EventLogRules != nil {
					eventLogRulesList := []interface{}{}
					for _, eventLogRule := range configInfo.ExtractRule.EventLogRules {
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

				if configInfo.ExtractRule.TimeKey != nil {
					extractRuleMap["time_key"] = configInfo.ExtractRule.TimeKey
				}

				if configInfo.ExtractRule.TimeFormat != nil {
					extractRuleMap["time_format"] = configInfo.ExtractRule.TimeFormat
				}

				if configInfo.ExtractRule.Delimiter != nil {
					extractRuleMap["delimiter"] = configInfo.ExtractRule.Delimiter
				}

				if configInfo.ExtractRule.LogRegex != nil {
					extractRuleMap["log_regex"] = configInfo.ExtractRule.LogRegex
				}

				if configInfo.ExtractRule.BeginRegex != nil {
					extractRuleMap["begin_regex"] = configInfo.ExtractRule.BeginRegex
				}

				if configInfo.ExtractRule.Keys != nil {
					extractRuleMap["keys"] = configInfo.ExtractRule.Keys
				}

				if configInfo.ExtractRule.FilterKeyRegex != nil {
					filterKeyRegexList := []interface{}{}
					for _, filterKeyRegex := range configInfo.ExtractRule.FilterKeyRegex {
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

				if configInfo.ExtractRule.UnMatchUpLoadSwitch != nil {
					extractRuleMap["un_match_up_load_switch"] = configInfo.ExtractRule.UnMatchUpLoadSwitch
				}

				if configInfo.ExtractRule.UnMatchLogKey != nil {
					extractRuleMap["un_match_log_key"] = configInfo.ExtractRule.UnMatchLogKey
				}

				if configInfo.ExtractRule.Backtracking != nil {
					extractRuleMap["backtracking"] = configInfo.ExtractRule.Backtracking
				}

				if configInfo.ExtractRule.IsGBK != nil {
					extractRuleMap["is_gbk"] = configInfo.ExtractRule.IsGBK
				}

				if configInfo.ExtractRule.JsonStandard != nil {
					extractRuleMap["json_standard"] = configInfo.ExtractRule.JsonStandard
				}

				if configInfo.ExtractRule.Protocol != nil {
					extractRuleMap["protocol"] = configInfo.ExtractRule.Protocol
				}

				if configInfo.ExtractRule.Address != nil {
					extractRuleMap["address"] = configInfo.ExtractRule.Address
				}

				if configInfo.ExtractRule.ParseProtocol != nil {
					extractRuleMap["parse_protocol"] = configInfo.ExtractRule.ParseProtocol
				}

				if configInfo.ExtractRule.MetadataType != nil {
					extractRuleMap["metadata_type"] = configInfo.ExtractRule.MetadataType
				}

				if configInfo.ExtractRule.PathRegex != nil {
					extractRuleMap["path_regex"] = configInfo.ExtractRule.PathRegex
				}

				if configInfo.ExtractRule.MetaTags != nil {
					metaTagsList := []interface{}{}
					for _, metaTags := range configInfo.ExtractRule.MetaTags {
						metaTagsMap := map[string]interface{}{}

						if metaTags.Key != nil {
							metaTagsMap["key"] = metaTags.Key
						}

						if metaTags.Value != nil {
							metaTagsMap["value"] = metaTags.Value
						}

						metaTagsList = append(metaTagsList, metaTagsMap)
					}

					extractRuleMap["meta_tags"] = metaTagsList
				}

				configInfoMap["extract_rule"] = []interface{}{extractRuleMap}
			}

			if configInfo.ExcludePaths != nil {
				excludePathsList := []interface{}{}
				for _, excludePaths := range configInfo.ExcludePaths {
					excludePathsMap := map[string]interface{}{}

					if excludePaths.Type != nil {
						excludePathsMap["type"] = excludePaths.Type
					}

					if excludePaths.Value != nil {
						excludePathsMap["value"] = excludePaths.Value
					}

					excludePathsList = append(excludePathsList, excludePathsMap)
				}

				configInfoMap["exclude_paths"] = excludePathsList
			}

			if configInfo.Output != nil {
				configInfoMap["output"] = configInfo.Output
			}

			if configInfo.UpdateTime != nil {
				configInfoMap["update_time"] = configInfo.UpdateTime
			}

			if configInfo.CreateTime != nil {
				configInfoMap["create_time"] = configInfo.CreateTime
			}

			if configInfo.UserDefineRule != nil {
				configInfoMap["user_define_rule"] = configInfo.UserDefineRule
			}

			// 添加新字段的处理
			if configInfo.ConfigExtraId != nil {
				configInfoMap["config_extra_id"] = configInfo.ConfigExtraId
			}

			if configInfo.ConfigFlag != nil {
				configInfoMap["config_flag"] = configInfo.ConfigFlag
			}

			if configInfo.AdvancedConfig != nil {
				configInfoMap["advanced_config"] = configInfo.AdvancedConfig
			}

			if configInfo.Source != nil {
				configInfoMap["source"] = configInfo.Source
			}

			ids = append(ids, *configInfo.ConfigId)
			tmpList = append(tmpList, configInfoMap)
		}

		_ = d.Set("configs", tmpList)
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

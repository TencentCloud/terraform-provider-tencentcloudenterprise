/*
Provides a resource to create a cls index.

# Example Usage

```hcl
resource "tencentcloudenterprise_cls_index" "complete_index" {

	  topic_id       = "a22f98e2-7331-44af-adc3-e21abcac1ae3"
	  coverage_field = "parse_error_field"

	  rule {
	    # 全文索引配置
	    full_text {
	      case_sensitive = true
	      tokenizer      = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
	      contain_z_h    = true
	    }

	    # 键值索引配置
	    key_value {
	      case_sensitive = true
	      template_type  = ""

	      key_values {
	        key = "level"
	        value {
	          type        = "text"
	          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
	          sql_flag    = true
	          contain_z_h = false
	          alias       = "log_level"
	        }
	      }

	      key_values {
	        key = "timestamp"
	        value {
	          type     = "long"
	          sql_flag = true
	          alias    = "log_time"
	        }
	      }

	      key_values {
	        key = "message"
	        value {
	          type        = "text"
	          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
	          sql_flag    = true
	          contain_z_h = true
	          alias       = "log_message"
	        }
	      }
	    }

	    dynamic_index {
	      status = true
	    }

	    # 元字段索引配置
	    tag {
	      case_sensitive = false

	      key_values {
	        key = "source"
	        value {
	          type        = "text"
	          tokenizer   = "@&?|#()='\",;:<>[]{}/ \n\t\r\\"
	          sql_flag    = true
	          contain_z_h = false
	          alias       = "log_source"
	        }
	      }

	      key_values {
	        key = "environment"
	        value {
	          type     = "text"
	          sql_flag = true
	          alias    = "env"
	        }
	      }
	    }
	  }
	  status = false
	  include_internal_fields = true
	  metadata_flag          = 1
	}

```

# Import

cls cos index can be imported using the id, e.g.

```
$ terraform import cloud_cls_index.index 0937e56f-4008-49d2-ad2d-69c52a9f11cc
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_index", CNDescription{
		TerraformTypeCN: "CLS索引配置",
		DescriptionCN:   "提供CLS索引配置资源，用于创建和管理日志服务索引配置。",
		AttributesCN: map[string]string{
			"topic_id":                "日志主题ID",
			"rule":                    "索引规则，FullText、KeyValue、Tag参数必须输入一个有效参数",
			"status":                  "是否生效，默认为true。true表示索引生效，false表示索引不生效。",
			"include_internal_fields": "内置保留字段（__FILENAME__，__HOSTNAME__及__SOURCE__）是否包含至全文索引，默认为false，推荐设置为true  false:不包含 true:包含",
			"metadata_flag":           "元数据字段（前缀为__TAG__的字段）是否包含至全文索引，默认为0，推荐设置为1  0:仅包含开启键值索引的元数据字段 1:包含所有元数据字段 2:不包含任何元数据字段",
			"coverage_field":          "自定义日志解析异常存储字段。",
			"dynamic_index":           "键值索引自动配置，为空时代表未开启该功能。 启用后自动将日志内的字段添加到键值索引中，包括日志中后续新增的字段。",
			"case_sensitive":          "是否大小写敏感。true表示大小写敏感，false表示大小写不敏感。",
			"key_values":              "需要建立索引的键值对信息；最大只能配置100个键值对",
			"key_value":               "键值索引配置，为空时代表未开启键值索引",
			"template_type":           "索引是否开启动态模板；若开启，则会根据上报的键值对配置索引，但是所有字段类型都是text，大小写敏感，不支持分析，采用默认分词符。已废弃",
			"key":                     "需要配置键值或者元字段索引的字段名称，仅支持字母、数字、下划线和-./@，且不能以下划线开头  注意： 1，元字段（tag）的Key无需额外添加__TAG__.前缀，与上传日志时对应的字段Key一致即可，云控制台展示时将自动添加__TAG__.前缀 2，键值索引（KeyValue）及元字段索引（Tag）中的Key总数不能超过300 3，Key的层级不能超过10层，例如a.b.c.d.e.f.g.h.j.k 4，不允许同时包含json父子级字段，例如a及a.b",
			"value":                   "需要配置键值或者元字段索引的字段名称，仅支持字母、数字、下划线和-./@，且不能以下划线开头  注意： 1，元字段（tag）的Key无需额外添加__TAG__.前缀，与上传日志时对应的字段Key一致即可，云控制台展示时将自动添加__TAG__.前缀 2，键值索引（KeyValue）及元字段索引（Tag）中的Key总数不能超过300 3，Key的层级不能超过10层，例如a.b.c.d.e.f.g.h.j.k 4，不允许同时包含json父子级字段，例如a及a.b",
			"full_text":               "全文索引配置, 为空时代表未开启全文索引",
			"tag":                     "元字段索引配置，为空时代表未开启元字段索引",
			"contain_z_h":             "是否包含中文。true代表包含中文，false代表不包含中文。",
			"alias":                   "索引键值别名",
			"tokenizer":               "分词符，其中的每个字符代表一个分词符； 仅支持英文符号、\\n\\t\\r及转义符\\； 注意：\\n\\t\\r本身已被转义，直接使用双引号包裹即可作为入参，无需再次转义。使用API Explorer进行调试时请使用JSON参数输入方式，以避免\\n\\t\\r被重复转义",
			"type":                    "字段是否开启分析功能。true表示开启分析功能，false表示关闭分析功能。",
			"sql_flag":                "字段是否开启分析功能。true表示开启分析功能，false表示关闭分析功能。",
		},
	})
}

func resourceTencentCloudClsIndex() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsIndexCreate,
		Read:        resourceTencentCloudClsIndexRead,
		Delete:      resourceTencentCloudClsIndexDelete,
		Update:      resourceTencentCloudClsIndexUpdate,
		Description: "Provides a resource to create and manage CLS index",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Log topic ID.",
			},
			"coverage_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Custom log parsing exception storage field.",
			},
			"rule": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "Index rule. FullText, KeyValue, Tag parameters must input one valid parameter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"full_text": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Full-Text index configuration. Empty means full-text index is not enabled.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"case_sensitive": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Case sensitivity. true means case sensitive, false means case insensitive.",
									},
									"tokenizer": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Full-Text index delimiter. Each character in the string represents a delimiter.",
									},
									"contain_z_h": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether Chinese characters are contained. true means contains Chinese, false means does not contain Chinese.",
									},
								},
							},
						},
						"dynamic_index": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Key-value index automatic configuration. Empty means this feature is not enabled. When enabled, fields in logs will be automatically added to key-value index, including new fields added later.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Dynamic index configuration switch.",
									},
								},
							},
						},
						"key_value": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Key-Value index configuration. Empty means key-value index is not enabled.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"case_sensitive": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Case sensitivity. true means case sensitive, false means case insensitive.",
									},
									"template_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Whether index enables dynamic template. If enabled, index will be configured based on reported key-value pairs, but all field types are text, case sensitive, analysis not supported, using default delimiters. Deprecated.",
									},
									"key_values": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Field name for key-value or metafield index configuration. Only supports letters, numbers, underscores and -./@, and cannot start with underscore. Note: 1. Metafield (tag) Key does not need __TAG__. prefix, consistent with field Key when uploading logs, console will automatically add __TAG__. prefix for display. 2. Total number of Keys in key-value index (KeyValue) and metafield index (Tag) cannot exceed 300. 3. Key hierarchy cannot exceed 10 levels, e.g. a.b.c.d.e.f.g.h.j.k. 4. Cannot contain both json parent and child fields, e.g. a and a.b.",
												},
												"value": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Required:    true,
													Description: "Field index description information.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Field type. Valid values: long, text, double.",
															},
															"tokenizer": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.",
															},
															"sql_flag": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether the analysis feature is enabled for the field.",
															},
															"contain_z_h": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether Chinese characters are contained. true means contains Chinese, false means does not contain Chinese.",
															},
															"alias": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Index key-value alias.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"tag": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Metafield index configuration. Empty means metafield index is not enabled.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"case_sensitive": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Case sensitivity. true means case sensitive, false means case insensitive.",
									},
									"key_values": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Key-Value pair information of the index to be created. Up to 100 key-value pairs can be configured.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Field name for key-value or metafield index configuration. Only supports letters, numbers, underscores and -./@, and cannot start with underscore. Note: 1. Metafield (tag) Key does not need __TAG__. prefix, consistent with field Key when uploading logs, console will automatically add __TAG__. prefix for display. 2. Total number of Keys in key-value index (KeyValue) and metafield index (Tag) cannot exceed 300. 3. Key hierarchy cannot exceed 10 levels, e.g. a.b.c.d.e.f.g.h.j.k. 4. Cannot contain both json parent and child fields, e.g. a and a.b.",
												},
												"value": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Required:    true,
													Description: "Field index description information.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"type": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Field type. Valid values: long, text, double.",
															},
															"tokenizer": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Field delimiter, which is meaningful only if the field type is text. Each character in the entered string represents a delimiter.",
															},
															"sql_flag": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether the analysis feature is enabled for the field.",
															},
															"contain_z_h": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Whether Chinese characters are contained. true means contains Chinese, false means does not contain Chinese.",
															},
															"alias": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Index key-value alias.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"include_internal_fields": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Internal field marker of full-text index. Default value: false. Valid value: false: excluding internal fields; true: including internal fields.",
			},
			"metadata_flag": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Metadata flag. Default value: 0. Valid value: 0: full-text index (including the metadata field with key-value index enabled); 1: full-text index (including all metadata fields); 2: full-text index (excluding metadata fields).",
			},
			"status": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to take effect. Default value: true.",
			},
		},
	}
}

func resourceTencentCloudClsIndexCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_index.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = cls.NewCreateIndexRequest()
		indexId string
	)

	if v, ok := d.GetOk("topic_id"); ok {
		indexId = v.(string)
		request.TopicId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("coverage_field"); ok {
		request.CoverageField = helper.String(v.(string))
	}

	if dMap, ok := helper.InterfacesHeadMap(d, "rule"); ok {
		ruleInfo := cls.RuleInfo{}
		if fullTextMap, ok := helper.InterfaceToMap(dMap, "full_text"); ok {
			fullTextInfo := cls.FullTextInfo{}
			if v, ok := fullTextMap["case_sensitive"]; ok {
				fullTextInfo.CaseSensitive = helper.Bool(v.(bool))
			}
			if v, ok := fullTextMap["tokenizer"]; ok {
				fullTextInfo.Tokenizer = helper.String(v.(string))
			}
			if v, ok := fullTextMap["contain_z_h"]; ok {
				fullTextInfo.ContainZH = helper.Bool(v.(bool))
			}
			ruleInfo.FullText = &fullTextInfo
		}

		if dynamicIndexMap, ok := helper.InterfaceToMap(dMap, "dynamic_index"); ok {
			dynamicIndex := cls.DynamicIndex{}
			if v, ok := dynamicIndexMap["status"]; ok {
				dynamicIndex.Status = helper.Bool(v.(bool))
			}
			ruleInfo.DynamicIndex = &dynamicIndex
		}

		if ruleKeyValueMap, ok := helper.InterfaceToMap(dMap, "key_value"); ok {
			ruleKeyValueInfo := cls.RuleKeyValueInfo{}
			if v, ok := ruleKeyValueMap["case_sensitive"]; ok {
				ruleKeyValueInfo.CaseSensitive = helper.Bool(v.(bool))
			}
			if v, ok := ruleKeyValueMap["template_type"]; ok {
				ruleKeyValueInfo.TemplateType = helper.String(v.(string))
			}
			if v, ok := ruleKeyValueMap["key_values"]; ok {
				for _, keyValue := range v.([]interface{}) {
					keyValueInfo := cls.KeyValueInfo{}
					keyValueMap := keyValue.(map[string]interface{})
					if v, ok := keyValueMap["key"]; ok {
						keyValueInfo.Key = helper.String(v.(string))
					}
					if valueMap, ok := helper.InterfaceToMap(keyValueMap, "value"); ok {
						valueInfo := cls.ValueInfo{}
						if v, ok := valueMap["type"]; ok {
							valueInfo.Type = helper.String(v.(string))
						}
						if v, ok := valueMap["tokenizer"]; ok {
							valueInfo.Tokenizer = helper.String(v.(string))
						}
						if v, ok := valueMap["sql_flag"]; ok {
							valueInfo.SqlFlag = helper.Bool(v.(bool))
						}
						if v, ok := valueMap["contain_z_h"]; ok {
							valueInfo.ContainZH = helper.Bool(v.(bool))
						}
						if v, ok := valueMap["alias"]; ok {
							valueInfo.Alias = helper.String(v.(string))
						}
						keyValueInfo.Value = &valueInfo
					}
					ruleKeyValueInfo.KeyValues = append(ruleKeyValueInfo.KeyValues, &keyValueInfo)
				}
			}
			ruleInfo.KeyValue = &ruleKeyValueInfo
		}

		if tagMap, ok := helper.InterfaceToMap(dMap, "tag"); ok {
			ruleTagInfo := cls.RuleTagInfo{}
			if v, ok := tagMap["case_sensitive"]; ok {
				ruleTagInfo.CaseSensitive = helper.Bool(v.(bool))
			}
			if v, ok := tagMap["key_values"]; ok {
				for _, keyValue := range v.([]interface{}) {
					keyValueInfo := cls.KeyValueInfo{}
					keyValueMap := keyValue.(map[string]interface{})
					if v, ok := keyValueMap["key"]; ok {
						keyValueInfo.Key = helper.String(v.(string))
					}
					if valueMap, ok := helper.InterfaceToMap(keyValueMap, "value"); ok {
						valueInfo := cls.ValueInfo{}
						if v, ok := valueMap["type"]; ok {
							valueInfo.Type = helper.String(v.(string))
						}
						if v, ok := valueMap["tokenizer"]; ok {
							valueInfo.Tokenizer = helper.String(v.(string))
						}
						if v, ok := valueMap["sql_flag"]; ok {
							valueInfo.SqlFlag = helper.Bool(v.(bool))
						}
						if v, ok := valueMap["contain_z_h"]; ok {
							valueInfo.ContainZH = helper.Bool(v.(bool))
						}
						if v, ok := valueMap["alias"]; ok {
							valueInfo.Alias = helper.String(v.(string))
						}
						keyValueInfo.Value = &valueInfo
					}
					ruleTagInfo.KeyValues = append(ruleTagInfo.KeyValues, &keyValueInfo)
				}
			}
			ruleInfo.Tag = &ruleTagInfo
		}
		request.Rule = &ruleInfo
	}

	if v, ok := d.GetOk("include_internal_fields"); ok {
		request.IncludeInternalFields = helper.Bool(v.(bool))
	}

	if v, ok := d.GetOk("metadata_flag"); ok {
		request.MetadataFlag = helper.IntUint64(v.(int))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateIndex(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create cls index failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(indexId)

	return resourceTencentCloudClsIndexRead(d, meta)
}

func resourceTencentCloudClsIndexRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_index.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = cls.NewDescribeIndexRequest()
		result  *cls.DescribeIndexResponse
	)
	id := d.Id()

	request.TopicId = &id

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().DescribeIndex(request)
		if e != nil {
			return retryError(e)
		}
		result = response
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read cls index failed, reason:%s\n", logId, err.Error())
		return err
	}

	res := result.Response

	if res.TopicId != nil {
		_ = d.Set("topic_id", res.TopicId)
	}

	if res.CoverageField != nil {
		_ = d.Set("coverage_field", res.CoverageField)
	}

	if res.Rule != nil {
		ruleMap := map[string]interface{}{}

		if res.Rule.FullText != nil {
			FullTextMap := map[string]interface{}{}
			if res.Rule.FullText.CaseSensitive != nil {
				FullTextMap["case_sensitive"] = res.Rule.FullText.CaseSensitive
			}
			if res.Rule.FullText.Tokenizer != nil {
				FullTextMap["tokenizer"] = res.Rule.FullText.Tokenizer
			}
			if res.Rule.FullText.ContainZH != nil {
				FullTextMap["contain_z_h"] = res.Rule.FullText.ContainZH
			}

			ruleMap["full_text"] = []interface{}{FullTextMap}
		}

		if res.Rule.DynamicIndex != nil {
			DynamicIndexMap := map[string]interface{}{}
			if res.Rule.DynamicIndex.Status != nil {
				DynamicIndexMap["status"] = res.Rule.DynamicIndex.Status
			}
			ruleMap["dynamic_index"] = []interface{}{DynamicIndexMap}
		}

		if res.Rule.KeyValue != nil {
			RuleKeyValueMap := map[string]interface{}{}
			if res.Rule.KeyValue.CaseSensitive != nil {
				RuleKeyValueMap["case_sensitive"] = res.Rule.KeyValue.CaseSensitive
			}
			if res.Rule.KeyValue.TemplateType != nil {
				RuleKeyValueMap["template_type"] = res.Rule.KeyValue.TemplateType
			}

			if res.Rule.KeyValue.KeyValues != nil {
				keyValuesList := []interface{}{}
				for _, keyValueInfo := range res.Rule.KeyValue.KeyValues {
					keyValueInfoMap := map[string]interface{}{}
					if keyValueInfo.Key != nil {
						keyValueInfoMap["key"] = keyValueInfo.Key
					}
					if keyValueInfo.Value != nil {
						valueInfoMap := map[string]interface{}{}
						if keyValueInfo.Value.Type != nil {
							valueInfoMap["type"] = keyValueInfo.Value.Type
						}
						if keyValueInfo.Value.Tokenizer != nil {
							valueInfoMap["tokenizer"] = keyValueInfo.Value.Tokenizer
						}
						if keyValueInfo.Value.SqlFlag != nil {
							valueInfoMap["sql_flag"] = keyValueInfo.Value.SqlFlag
						}
						if keyValueInfo.Value.ContainZH != nil {
							valueInfoMap["contain_z_h"] = keyValueInfo.Value.ContainZH
						}
						if keyValueInfo.Value.Alias != nil {
							valueInfoMap["alias"] = keyValueInfo.Value.Alias
						}
						keyValueInfoMap["value"] = []interface{}{valueInfoMap}
					}
					keyValuesList = append(keyValuesList, keyValueInfoMap)
				}
				RuleKeyValueMap["key_values"] = keyValuesList
			}
			ruleMap["key_value"] = []interface{}{RuleKeyValueMap}
		}

		if res.Rule.Tag != nil {
			ruleTagMap := map[string]interface{}{
				"case_sensitive": res.Rule.Tag.CaseSensitive,
			}
			if res.Rule.Tag.KeyValues != nil {
				keyValuesList := []interface{}{}
				for _, keyValueInfo := range res.Rule.Tag.KeyValues {
					keyValueInfoMap := map[string]interface{}{
						"key": keyValueInfo.Key,
					}
					if keyValueInfo.Value != nil {
						valueInfoMap := map[string]interface{}{}
						if keyValueInfo.Value.Type != nil {
							valueInfoMap["type"] = keyValueInfo.Value.Type
						}
						if keyValueInfo.Value.Tokenizer != nil {
							valueInfoMap["tokenizer"] = keyValueInfo.Value.Tokenizer
						}
						if keyValueInfo.Value.SqlFlag != nil {
							valueInfoMap["sql_flag"] = keyValueInfo.Value.SqlFlag
						}
						if keyValueInfo.Value.ContainZH != nil {
							valueInfoMap["contain_z_h"] = keyValueInfo.Value.ContainZH
						}
						if keyValueInfo.Value.Alias != nil {
							valueInfoMap["alias"] = keyValueInfo.Value.Alias
						}
						keyValueInfoMap["value"] = []interface{}{valueInfoMap}
					}
					keyValuesList = append(keyValuesList, keyValueInfoMap)
				}
				ruleTagMap["key_values"] = keyValuesList
			}
			ruleMap["tag"] = []interface{}{ruleTagMap}
		}

		_ = d.Set("rule", []interface{}{ruleMap})
	}

	if res.IncludeInternalFields != nil {
		_ = d.Set("include_internal_fields", res.IncludeInternalFields)
	}

	if res.MetadataFlag != nil {
		_ = d.Set("metadata_flag", res.MetadataFlag)
	}

	return nil
}

func resourceTencentCloudClsIndexUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_index.update")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = cls.NewModifyIndexRequest()
	)
	id := d.Id()

	request.TopicId = &id

	if d.HasChange("coverage_field") {
		if v, ok := d.GetOk("coverage_field"); ok {
			request.CoverageField = helper.String(v.(string))
		}
	}

	if d.HasChange("rule") {
		if dMap, ok := helper.InterfacesHeadMap(d, "rule"); ok {
			ruleInfo := cls.RuleInfo{}
			if fullTextMap, ok := helper.InterfaceToMap(dMap, "full_text"); ok {
				fullTextInfo := cls.FullTextInfo{}
				if v, ok := fullTextMap["case_sensitive"]; ok {
					fullTextInfo.CaseSensitive = helper.Bool(v.(bool))
				}
				if v, ok := fullTextMap["tokenizer"]; ok {
					fullTextInfo.Tokenizer = helper.String(v.(string))
				}
				if v, ok := fullTextMap["contain_z_h"]; ok {
					fullTextInfo.ContainZH = helper.Bool(v.(bool))
				}
				ruleInfo.FullText = &fullTextInfo
			}

			if dynamicIndexMap, ok := helper.InterfaceToMap(dMap, "dynamic_index"); ok {
				dynamicIndex := cls.DynamicIndex{}
				if v, ok := dynamicIndexMap["status"]; ok {
					dynamicIndex.Status = helper.Bool(v.(bool))
				}
				ruleInfo.DynamicIndex = &dynamicIndex
			}

			if ruleKeyValueMap, ok := helper.InterfaceToMap(dMap, "key_value"); ok {
				ruleKeyValueInfo := cls.RuleKeyValueInfo{}
				if v, ok := ruleKeyValueMap["case_sensitive"]; ok {
					ruleKeyValueInfo.CaseSensitive = helper.Bool(v.(bool))
				}
				if v, ok := ruleKeyValueMap["template_type"]; ok {
					ruleKeyValueInfo.TemplateType = helper.String(v.(string))
				}
				if v, ok := ruleKeyValueMap["key_values"]; ok {
					for _, keyValue := range v.([]interface{}) {
						keyValueInfo := cls.KeyValueInfo{}
						keyValueMap := keyValue.(map[string]interface{})
						if v, ok := keyValueMap["key"]; ok {
							keyValueInfo.Key = helper.String(v.(string))
						}
						if v, ok := keyValueMap["value"]; ok {
							valueMap := v.([]interface{})[0].(map[string]interface{})
							valueInfo := cls.ValueInfo{}
							if v, ok := valueMap["type"]; ok {
								valueInfo.Type = helper.String(v.(string))
							}
							if v, ok := valueMap["tokenizer"]; ok {
								valueInfo.Tokenizer = helper.String(v.(string))
							}
							if v, ok := valueMap["sql_flag"]; ok {
								valueInfo.SqlFlag = helper.Bool(v.(bool))
							}
							if v, ok := valueMap["contain_z_h"]; ok {
								valueInfo.ContainZH = helper.Bool(v.(bool))
							}
							if v, ok := valueMap["alias"]; ok {
								valueInfo.Alias = helper.String(v.(string))
							}
							keyValueInfo.Value = &valueInfo
						}
						ruleKeyValueInfo.KeyValues = append(ruleKeyValueInfo.KeyValues, &keyValueInfo)
					}
				}

				ruleInfo.KeyValue = &ruleKeyValueInfo
			}

			if tagMap, ok := helper.InterfaceToMap(dMap, "tag"); ok {
				ruleTagInfo := cls.RuleTagInfo{}
				if v, ok := tagMap["case_sensitive"]; ok {
					ruleTagInfo.CaseSensitive = helper.Bool(v.(bool))
				}
				if v, ok := tagMap["key_values"]; ok {
					for _, keyValue := range v.([]interface{}) {
						keyValueInfo := cls.KeyValueInfo{}
						keyValueMap := keyValue.(map[string]interface{})
						if v, ok := keyValueMap["key"]; ok {
							keyValueInfo.Key = helper.String(v.(string))
						}
						if v, ok := keyValueMap["value"]; ok {
							valueMap := v.([]interface{})[0].(map[string]interface{})
							valueInfo := cls.ValueInfo{}
							if v, ok := valueMap["type"]; ok {
								valueInfo.Type = helper.String(v.(string))
							}
							if v, ok := valueMap["tokenizer"]; ok {
								valueInfo.Tokenizer = helper.String(v.(string))
							}
							if v, ok := valueMap["sql_flag"]; ok {
								valueInfo.SqlFlag = helper.Bool(v.(bool))
							}
							if v, ok := valueMap["contain_z_h"]; ok {
								valueInfo.ContainZH = helper.Bool(v.(bool))
							}
							if v, ok := valueMap["alias"]; ok {
								valueInfo.Alias = helper.String(v.(string))
							}
							keyValueInfo.Value = &valueInfo
						}
						ruleTagInfo.KeyValues = append(ruleTagInfo.KeyValues, &keyValueInfo)
					}
				}
				ruleInfo.Tag = &ruleTagInfo
			}
			request.Rule = &ruleInfo
		}
	}

	if d.HasChange("include_internal_fields") {
		if v, ok := d.GetOk("include_internal_fields"); ok {
			request.IncludeInternalFields = helper.Bool(v.(bool))
		}
	}
	if d.HasChange("metadata_flag") {
		if v, ok := d.GetOk("metadata_flag"); ok {
			request.MetadataFlag = helper.IntUint64(v.(int))
		}
	}
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyIndex(request)
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

	return resourceTencentCloudClsIndexRead(d, meta)
}

func resourceTencentCloudClsIndexDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_shipper.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	id := d.Id()

	if err := service.DeleteClsIndex(ctx, id); err != nil {
		return err
	}
	return nil
}

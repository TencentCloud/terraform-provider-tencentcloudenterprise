/*
Provides a resource to create a cls cos shipper.

# Example Usage

```hcl

	resource "tencentcloudenterprise_cls_cos_shipper" "shipper" {
	  bucket       = "preset-scf-bucket-1308919341"
	  interval     = 300
	  max_size     = 200
	  partition    = "/%Y/%m/%d/%H/"
	  prefix       = "ap-guangzhou-fffsasad-1649734752"
	  shipper_name = "ap-guangzhou-fffsasad-1649734752"
	  topic_id     = "4d07fba0-b93e-4e0b-9a7f-d58542560bbb"

	  compress {
	    format = "lzop"
	  }

	  content {
	    format = "json"

	    json {
	      enable_tag  = true
	      meta_fields = [
	        "__FILENAME__",
	        "__SOURCE__",
	        "__TIMESTAMP__",
	      ]
	    }
	  }
	}

```

# Import

cls cos shipper can be imported using the id, e.g.

```
$ terraform import cloud_cls_cos_shipper.shipper 5d1b7b2a-c163-4c48-bb01-9ee00584d761
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
	registerResourceDescriptionProvider("tencentcloudenterprise_cls_cos_shipper", CNDescription{
		TerraformTypeCN: "CLS COS投递",
		DescriptionCN:   "提供CLS COS投递资源，用于将日志数据投递到COS存储桶。",
		AttributesCN: map[string]string{
			"topic_id":               "创建的投递规则所属的日志主题ID",
			"bucket":                 "创建的投递规则投递的bucket",
			"prefix":                 "创建的投递规则投递目录的前缀",
			"shipper_name":           "投递规则的名字",
			"interval":               "投递的时间间隔，单位 秒，默认300，范围 300-900",
			"max_size":               "投递的文件的最大值，单位 MB，默认100，范围 5-256",
			"filter_rules":           "投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递",
			"partition":              "投递日志的分区规则，支持strftime的时间格式表示",
			"compress":               "投递日志的压缩配置",
			"content":                "投递日志的内容格式配置",
			"filename_mode":          "投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）",
			"custom_uin":             "跨账户uin，用于支持跨账户bucket投递",
			"start_time":             "投递数据范围的开始时间点，不能超出日志主题的生命周期起点。如果用户不填写，默认为用户新建投递任务的时间。",
			"end_time":               "投递数据范围的结束时间点，不能填写未来时间。如果用户不填写，默认为持续投递，即无限。",
			"shipper_id":             "投递规则ID",
			"key":                    "过滤规则Key",
			"regex":                  "过滤规则",
			"value":                  "过滤规则Value",
			"format":                 "压缩格式，支持gzip、lzop、snappy和none不压缩",
			"parquet":                "parquet格式内容描述",
			"parquet_key_info":       "ParquetKeyInfo数组",
			"key_name":               "键值名称",
			"key_type":               "数据类型，目前支持6种类型：string、boolean、int32、int64、float、double",
			"key_non_existing_field": "解析失败赋值信息",
			"csv":                    "csv格式内容描述",
			"print_key":              "csv首行是否打印key。true表示首行打印key，false表示首行不打印key",
			"keys":                   "每列key的名字",
			"delimiter":              "各字段间的分隔符",
			"escape_char":            "各字段间的分隔符",
			"non_existing":           "对于上面指定的不存在字段使用该内容填充",
			"json":                   "json格式内容描述",
			"meta_fields":            "元数据信息列表",
			"json_type":              "投递Json格式，0：字符串方式投递；1:以结构化方式投递",
			"enable_tag":             "启用标志。true表示启用TAG标志，false表示关闭TAG标志。",
			"non_existing_field":     "对于上面指定的不存在字段使用该内容填充",
		},
	})
}

func resourceTencentCloudClsCosShipper() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudClsCosShipperCreate,
		Read:        resourceTencentCloudClsCosShipperRead,
		Delete:      resourceTencentCloudClsCosShipperDelete,
		Update:      resourceTencentCloudClsCosShipperUpdate,
		Description: "Provides a resource to create and manage CLS COS shipper for shipping log data to COS bucket",
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the log topic to which the shipping rule to be created belongs.",
			},
			"bucket": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Destination bucket in the shipping rule to be created.",
			},
			"prefix": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Prefix of the shipping directory in the shipping rule to be created.",
			},
			"shipper_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Shipping rule name.",
			},
			"interval": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Shipping time interval in seconds. Default value: 300. Value range: 300~900.",
			},
			"max_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Maximum size of a file to be shipped, in MB. Default value: 256. Value range: 100~256.",
			},
			"filename_mode": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "File naming configuration for shipped logs. 0: random number naming, 1: shipping time naming. Default value: 0 (random number naming).",
			},
			"custom_uin": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Cross-account UIN for supporting cross-account bucket shipping.",
			},
			"start_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Start time point for the shipping data range, cannot exceed the lifecycle start point of the log topic. If not filled by user, defaults to the time when the user creates the shipping task.",
			},
			"end_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "End time point for the shipping data range, cannot be a future time. If not filled by user, defaults to continuous shipping, i.e., infinite.",
			},
			"filter_rules": {
				Type:     schema.TypeList,
				Optional: true,
				Description: "Filter rules for shipped logs. Only logs matching the rules can be shipped. All rules are in the AND relationship, and up to five rules can be added. " +
					"If the array is empty, no filtering will be performed, and all logs will be shipped.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Filter rule key.",
						},
						"regex": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Filter rule.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Filter rule value.",
						},
					},
				},
			},
			"partition": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Partition rule of shipped log, which can be represented in strftime time format.",
			},
			"compress": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"format": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Compression format. Valid values: gzip, lzop, none (no compression).",
						},
					},
				},
				Description: "Compression configuration of shipped log.",
			},
			"content": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Format configuration of shipped log content.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"format": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Content format. Valid values: json, csv.",
						},
						"parquet": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"parquet_key_info": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key_name": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "key name",
												},
												"key_type": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "key type: string、boolean、int32、int64、float、double",
												},
												"key_non_existing_field": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "default value when parsing failed",
												},
											},
										},
									},
								},
							},
							Description: "",
						},
						"csv": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"print_key": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Whether to print key on the first row of the CSV file.",
									},
									"keys": {
										Type:        schema.TypeList,
										Required:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Names of keys.Note: this field may return null, indicating that no valid values can be obtained.",
									},
									"delimiter": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Field delimiter.",
									},
									"escape_char": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Field delimiter.",
									},
									"non_existing_field": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Content used to populate non-existing fields.",
									},
								},
							},
							Description: "CSV format content description.Note: this field may return null, indicating that no valid values can be obtained.",
						},
						"json": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_tag": {
										Type:        schema.TypeBool,
										Required:    true,
										Description: "Enablement flag.",
									},
									"meta_fields": {
										Type:        schema.TypeList,
										Required:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Metadata information list. Note: this field may return null, indicating that no valid values can be obtained..",
									},
									"json_type": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "JSON delivery format. 0: string delivery; 1: structured delivery.",
									},
								},
							},
							Description: "JSON format content description.Note: this field may return null, indicating that no valid values can be obtained.",
						},
					},
				},
			},
			"shipper_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Shipping rule ID.",
			},
		},
	}
}

func resourceTencentCloudClsCosShipperCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_shipper.create")()

	logId := getLogId(contextNil)

	var (
		request  = cls.NewCreateShipperRequest()
		response *cls.CreateShipperResponse
	)

	if v, ok := d.GetOk("topic_id"); ok {
		request.TopicId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bucket"); ok {
		request.Bucket = helper.String(v.(string))
	}

	if v, ok := d.GetOk("prefix"); ok {
		request.Prefix = helper.String(v.(string))
	}

	if v, ok := d.GetOk("shipper_name"); ok {
		request.ShipperName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("interval"); ok {
		request.Interval = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("max_size"); ok {
		request.MaxSize = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("filename_mode"); ok {
		request.FilenameMode = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("custom_uin"); ok {
		request.CustomUin = helper.Uint64(uint64(v.(int)))
	}

	if v, ok := d.GetOk("filter_rules"); ok {
		filterRules := make([]*cls.FilterRuleInfo, 0, 10)
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			filterRule := cls.FilterRuleInfo{}
			if v, ok := dMap["key"]; ok {
				filterRule.Key = helper.String(v.(string))
			}
			if v, ok := dMap["regex"]; ok {
				filterRule.Regex = helper.String(v.(string))
			}
			if v, ok := dMap["value"]; ok {
				filterRule.Value = helper.String(v.(string))
			}
			filterRules = append(filterRules, &filterRule)
		}
		request.FilterRules = filterRules
	}

	if v, ok := d.GetOk("partition"); ok {
		request.Partition = helper.String(v.(string))
	}

	if v, ok := d.GetOk("compress"); ok {
		compresses := make([]*cls.CompressInfo, 0, 10)
		if len(v.([]interface{})) != 1 {
			return fmt.Errorf("need only one compress.")
		}
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			compress := cls.CompressInfo{}
			if v, ok := dMap["format"]; ok {
				compress.Format = helper.String(v.(string))
			}
			compresses = append(compresses, &compress)
		}
		request.Compress = compresses[0]
	}

	if v, ok := d.GetOk("content"); ok {
		contents := make([]*cls.ContentInfo, 0, 10)
		if len(v.([]interface{})) != 1 {
			return fmt.Errorf("need only one content.")
		}
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			content := cls.ContentInfo{}
			if v, ok := dMap["format"]; ok {
				content.Format = helper.String(v.(string))
			}
			if v, ok := dMap["csv"]; ok {
				if len(v.([]interface{})) == 1 {
					csv := v.([]interface{})[0].(map[string]interface{})
					csvInfo := cls.CsvInfo{}
					csvInfo.PrintKey = helper.Bool(csv["print_key"].(bool))
					keys := csv["keys"].([]interface{})
					for _, key := range keys {
						csvInfo.Keys = append(csvInfo.Keys, helper.String(key.(string)))
					}
					csvInfo.Delimiter = helper.String(csv["delimiter"].(string))
					csvInfo.EscapeChar = helper.String(csv["escape_char"].(string))
					csvInfo.NonExistingField = helper.String(csv["non_existing_field"].(string))
					content.Csv = &csvInfo
				}
			}
			if v, ok := dMap["parquet"]; ok {
				if len(v.([]interface{})) == 1 {
					parquet := v.([]interface{})[0].(map[string]interface{})
					parquetInfo := cls.ParquetInfo{}
					if parquetKeyInfo, ok := parquet["parquet_key_info"].([]interface{}); ok {
						for _, keyInfo := range parquetKeyInfo {
							keyInfoMap := keyInfo.(map[string]interface{})
							parquetKeyInfoItem := cls.ParquetKeyInfo{}
							if keyName, ok := keyInfoMap["key_name"].(string); ok {
								parquetKeyInfoItem.KeyName = helper.String(keyName)
							}
							if keyType, ok := keyInfoMap["key_type"].(string); ok {
								parquetKeyInfoItem.KeyType = helper.String(keyType)
							}
							if keyNonExistingField, ok := keyInfoMap["key_non_existing_field"].(string); ok {
								parquetKeyInfoItem.KeyNonExistingField = helper.String(keyNonExistingField)
							}
							parquetInfo.ParquetKeyInfo = append(parquetInfo.ParquetKeyInfo, &parquetKeyInfoItem)
						}
					}
					content.Parquet = &parquetInfo
				}
			}
			if v, ok := dMap["json"]; ok {
				if len(v.([]interface{})) == 1 {

					json := v.([]interface{})[0].(map[string]interface{})
					jsonInfo := cls.JsonInfo{}
					jsonInfo.EnableTag = helper.Bool(json["enable_tag"].(bool))
					metaFields := json["meta_fields"].([]interface{})
					for _, metaField := range metaFields {
						jsonInfo.MetaFields = append(jsonInfo.MetaFields, helper.String(metaField.(string)))
					}
					if jsonType, ok := json["json_type"].(int); ok {
						jsonInfo.JsonType = helper.Int64(int64(jsonType))
					}
					content.Json = &jsonInfo
				}
			}
			contents = append(contents, &content)
		}
		request.Content = contents[0]
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().CreateShipper(request)
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
		log.Printf("[CRITAL]%s create cls cos shipper failed, reason:%+v", logId, err)
		return err
	}

	id := *response.Response.ShipperId
	d.SetId(id)
	return resourceTencentCloudClsCosShipperRead(d, meta)
}

func resourceTencentCloudClsCosShipperRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_shipper.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	id := d.Id()

	shipper, err := service.DescribeClsCosShipperById(ctx, id)

	if err != nil {
		return err
	}

	if shipper == nil {
		d.SetId("")
		return fmt.Errorf("resource `Shipper` %s does not exist", id)
	}

	_ = d.Set("topic_id", shipper.TopicId)
	_ = d.Set("bucket", shipper.Bucket)
	_ = d.Set("prefix", shipper.Prefix)
	_ = d.Set("shipper_name", shipper.ShipperName)
	if shipper.Interval != nil {
		_ = d.Set("interval", shipper.Interval)
	}
	if shipper.MaxSize != nil {
		_ = d.Set("max_size", shipper.MaxSize)
	}

	if shipper.FilenameMode != nil {
		_ = d.Set("filename_mode", shipper.FilenameMode)
	}

	if shipper.CustomUin != nil {
		_ = d.Set("custom_uin", shipper.CustomUin)
	}

	if shipper.StartTime != nil {
		_ = d.Set("start_time", shipper.StartTime)
	}

	if shipper.EndTime != nil {
		_ = d.Set("end_time", shipper.EndTime)
	}

	_ = d.Set("shipper_id", shipper.ShipperId)

	if shipper.FilterRules != nil {
		filterRules := make([]interface{}, 0, 100)
		for _, v := range shipper.FilterRules {
			filterRule := map[string]interface{}{
				"key":   v.Key,
				"regex": v.Regex,
				"value": v.Value,
			}
			filterRules = append(filterRules, filterRule)
		}
		_ = d.Set("filter_rules", filterRules)
	}

	if shipper.Partition != nil {
		_ = d.Set("partition", shipper.Partition)
	}

	if shipper.Compress != nil {
		compress := map[string]interface{}{
			"format": shipper.Compress.Format,
		}
		_ = d.Set("compress", []interface{}{compress})
	}

	if shipper.Content != nil {
		content := map[string]interface{}{
			"format": shipper.Content.Format,
		}
		if shipper.Content.Csv != nil {
			csv := map[string]interface{}{
				"print_key":          shipper.Content.Csv.PrintKey,
				"keys":               shipper.Content.Csv.Keys,
				"delimiter":          shipper.Content.Csv.Delimiter,
				"escape_char":        shipper.Content.Csv.EscapeChar,
				"non_existing_field": shipper.Content.Csv.NonExistingField,
			}
			content["csv"] = []interface{}{csv}
		}
		if shipper.Content.Parquet != nil {
			parquetKeyInfoList := make([]interface{}, 0)
			for _, keyInfo := range shipper.Content.Parquet.ParquetKeyInfo {
				parquetKeyInfoItem := map[string]interface{}{
					"key_name":               keyInfo.KeyName,
					"key_type":               keyInfo.KeyType,
					"key_non_existing_field": keyInfo.KeyNonExistingField,
				}
				parquetKeyInfoList = append(parquetKeyInfoList, parquetKeyInfoItem)
			}
			parquet := map[string]interface{}{
				"parquet_key_info": parquetKeyInfoList,
			}
			content["parquet"] = []interface{}{parquet}
		}
		if shipper.Content.Json != nil {
			json := map[string]interface{}{
				"enable_tag":  shipper.Content.Json.EnableTag,
				"meta_fields": shipper.Content.Json.MetaFields,
			}
			if shipper.Content.Json.JsonType != nil {
				json["json_type"] = shipper.Content.Json.JsonType
			}
			content["json"] = []interface{}{json}
		}
		_ = d.Set("content", []interface{}{content})
	}
	return nil
}

func resourceTencentCloudClsCosShipperUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_shipper.update")()
	logId := getLogId(contextNil)
	request := cls.NewModifyShipperRequest()

	request.ShipperId = helper.String(d.Id())

	if d.HasChange("bucket") {
		if v, ok := d.GetOk("bucket"); ok {
			request.Bucket = helper.String(v.(string))
		}
	}

	if d.HasChange("prefix") {
		if v, ok := d.GetOk("prefix"); ok {
			request.Prefix = helper.String(v.(string))
		}
	}

	if d.HasChange("shipper_name") {
		if v, ok := d.GetOk("shipper_name"); ok {
			request.ShipperName = helper.String(v.(string))
		}
	}

	if d.HasChange("interval") {
		if v, ok := d.GetOk("interval"); ok {
			request.Interval = helper.IntUint64(v.(int))
		}
	}

	if d.HasChange("max_size") {
		if v, ok := d.GetOk("max_size"); ok {
			request.MaxSize = helper.IntUint64(v.(int))
		}
	}

	if v, ok := d.GetOk("filename_mode"); ok {
		request.FilenameMode = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("custom_uin"); ok {
		request.CustomUin = helper.Uint64(uint64(v.(int)))
	}

	if v, ok := d.GetOk("start_time"); ok {
		request.StartTime = helper.Int64(int64(v.(int)))
	}

	if v, ok := d.GetOk("end_time"); ok {
		request.EndTime = helper.Int64(int64(v.(int)))
	}

	if d.HasChange("filter_rules") {
		if v, ok := d.GetOk("filter_rules"); ok {
			filterRules := make([]*cls.FilterRuleInfo, 0, 10)
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				filterRule := cls.FilterRuleInfo{}
				if v, ok := dMap["key"]; ok {
					filterRule.Key = helper.String(v.(string))
				}
				if v, ok := dMap["regex"]; ok {
					filterRule.Regex = helper.String(v.(string))
				}
				if v, ok := dMap["value"]; ok {
					filterRule.Value = helper.String(v.(string))
				}
				filterRules = append(filterRules, &filterRule)
			}
			request.FilterRules = filterRules
		}
	}

	if d.HasChange("partition") {
		if v, ok := d.GetOk("partition"); ok {
			request.Partition = helper.String(v.(string))
		}
	}

	if d.HasChange("compress") {
		if v, ok := d.GetOk("compress"); ok {
			compresses := make([]*cls.CompressInfo, 0, 10)
			if len(v.([]interface{})) != 1 {
				return fmt.Errorf("need only one compress.")
			}
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				compress := cls.CompressInfo{}
				if v, ok := dMap["format"]; ok {
					compress.Format = helper.String(v.(string))
				}
				compresses = append(compresses, &compress)
			}
			request.Compress = compresses[0]
		}
	}

	if d.HasChange("content") {
		if v, ok := d.GetOk("content"); ok {
			contents := make([]*cls.ContentInfo, 0, 10)
			if len(v.([]interface{})) != 1 {
				return fmt.Errorf("need only one content.")
			}
			for _, item := range v.([]interface{}) {
				dMap := item.(map[string]interface{})
				content := cls.ContentInfo{}
				if v, ok := dMap["format"]; ok {
					content.Format = helper.String(v.(string))
				}
				if v, ok := dMap["csv"]; ok {
					if len(v.([]interface{})) == 1 {
						csv := v.([]interface{})[0].(map[string]interface{})
						csvInfo := cls.CsvInfo{}
						if printKey, ok := csv["print_key"].(bool); ok {
							csvInfo.PrintKey = helper.Bool(printKey)
						}
						if keysInterface, ok := csv["keys"].([]interface{}); ok {
							keys := make([]*string, 0, len(keysInterface))
							for _, key := range keysInterface {
								if keyStr, ok := key.(string); ok {
									keys = append(keys, helper.String(keyStr))
								}
							}
							csvInfo.Keys = keys
						}
						if delimiter, ok := csv["delimiter"].(string); ok {
							csvInfo.Delimiter = helper.String(delimiter)
						}
						if escapeChar, ok := csv["escape_char"].(string); ok {
							csvInfo.EscapeChar = helper.String(escapeChar)
						}
						if nonExistingField, ok := csv["non_existing_field"].(string); ok {
							csvInfo.NonExistingField = helper.String(nonExistingField)
						}
						content.Csv = &csvInfo
					}
				}
				if v, ok := dMap["parquet"]; ok {
					if len(v.([]interface{})) == 1 {
						parquet := v.([]interface{})[0].(map[string]interface{})
						parquetInfo := cls.ParquetInfo{}
						if parquetKeyInfo, ok := parquet["parquet_key_info"].([]interface{}); ok {
							for _, keyInfo := range parquetKeyInfo {
								keyInfoMap := keyInfo.(map[string]interface{})
								parquetKeyInfoItem := cls.ParquetKeyInfo{}
								if keyName, ok := keyInfoMap["key_name"].(string); ok {
									parquetKeyInfoItem.KeyName = helper.String(keyName)
								}
								if keyType, ok := keyInfoMap["key_type"].(string); ok {
									parquetKeyInfoItem.KeyType = helper.String(keyType)
								}
								if keyNonExistingField, ok := keyInfoMap["key_non_existing_field"].(string); ok {
									parquetKeyInfoItem.KeyNonExistingField = helper.String(keyNonExistingField)
								}
								parquetInfo.ParquetKeyInfo = append(parquetInfo.ParquetKeyInfo, &parquetKeyInfoItem)
							}
						}
						content.Parquet = &parquetInfo
					}
				}
				if v, ok := dMap["json"]; ok {
					if len(v.([]interface{})) == 1 {

						json := v.([]interface{})[0].(map[string]interface{})
						jsonInfo := cls.JsonInfo{}
						jsonInfo.EnableTag = helper.Bool(json["enable_tag"].(bool))
						metaFields := json["meta_fields"].([]interface{})
						for _, metaField := range metaFields {
							jsonInfo.MetaFields = append(jsonInfo.MetaFields, helper.String(metaField.(string)))
						}
						if jsonType, ok := json["json_type"].(int); ok {
							jsonInfo.JsonType = helper.Int64(int64(jsonType))
						}
						content.Json = &jsonInfo
					}
				}
				contents = append(contents, &content)
			}
			request.Content = contents[0]
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseClsClient().ModifyShipper(request)
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

	return resourceTencentCloudClsCosShipperRead(d, meta)
}

func resourceTencentCloudClsCosShipperDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_cls_cos_shipper.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}
	id := d.Id()

	if err := service.DeleteClsCosShipper(ctx, id); err != nil {
		return err
	}

	return nil
}

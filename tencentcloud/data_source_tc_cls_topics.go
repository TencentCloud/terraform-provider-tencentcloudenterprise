/*
Use this data source to query detailed information of cls topics

# Example Usage

```hcl

	data "tencentcloudenterprise_cls_topics" "topics" {
	  filters {
	    key    = "topicName"
	    values = ["example"]
	  }
	}

	output "topics_list" {
	  value = data.tencentcloudenterprise_cls_topics.topics.topics
	}

```
*/
package tencentcloud

import (
	"context"

	cls "terraform-provider-tencentcloudenterprise/sdk/cls/v20201016"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cls_topics", CNDescription{
		TerraformTypeCN: "CLS 日志主题列表",
		DescriptionCN:   "提供 CLS 日志主题列表数据源，用于查询满足条件的日志主题信息。",
		AttributesCN: map[string]string{
			"filters":              "过滤条件列表",
			"key":                  "过滤字段名",
			"values":               "过滤字段值列表",
			"precise_search":       "Filters 字段的匹配模式：0 模糊匹配（默认），1 精确匹配 topicName，2 精确匹配 logsetName，3 同时精确匹配",
			"biz_type":             "主题类型：0 日志主题（默认），1 指标主题",
			"topics":               "日志主题列表",
			"logset_id":            "日志集 ID",
			"topic_id":             "日志主题 ID",
			"topic_name":           "日志主题名称",
			"partition_count":      "主题分区个数",
			"index":                "是否开启索引",
			"assumer_name":         "云产品标识",
			"create_time":          "创建时间",
			"status":               "是否开启采集",
			"tags":                 "标签信息",
			"auto_split":           "是否开启自动分裂",
			"max_split_partitions": "自动分裂时允许的最大分区数",
			"storage_type":         "存储类型，hot 标准存储，cold 低频存储",
			"period":               "生命周期，单位天",
			"sub_assumer_name":     "云产品二级标识",
			"describes":            "日志主题描述",
			"hot_period":           "标准存储的生命周期（日志沉降）",
			"biz_type_out":         "主题类型",
			"is_web_tracking":      "免鉴权开关",
			"result_output_file":   "用于保存结果，可视化界面不可用",
		},
	})
}

func dataSourceTencentCloudClsTopics() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudClsTopicsRead,
		Description: "Use this data source to query detailed information of cls topics.",
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter condition list. Each request can have up to 10 Filters and 100 Filter.Values.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Field to be filtered.",
						},
						"values": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "Value to be filtered.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},

			"precise_search": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Match mode for Filters fields. 0: Fuzzy match for topicName and logsetName (default). 1: Exact match for topicName. 2: Exact match for logsetName. 3: Exact match for both.",
			},

			"biz_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Topic type. 0: Log topic (default). 1: Metric topic.",
			},

			"topics": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log topic list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logset_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Logset ID.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topic ID.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topic name.",
						},
						"partition_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of topic partitions.",
						},
						"index": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the topic has indexing enabled.",
						},
						"assumer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud product identifier.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"status": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the topic has log collection enabled.",
						},
						"tags": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Tag information bound to the topic.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The tag key.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The tag value.",
									},
								},
							},
						},
						"auto_split": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether automatic split is enabled for this topic.",
						},
						"max_split_partitions": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of partitions to split into if automatic split is enabled.",
						},
						"storage_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Storage type of the topic. hot: standard storage; cold: IA storage.",
						},
						"period": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Lifecycle in days. Value range: 1-3600 (3640 indicates permanent retention).",
						},
						"sub_assumer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud product sub-identifier.",
						},
						"describes": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topic description.",
						},
						"hot_period": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Standard storage lifecycle for log sinking. HotPeriod=0 indicates log sinking is not enabled.",
						},
						"biz_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Topic type. 0: log topic; 1: metric topic.",
						},
						"is_web_tracking": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Free authentication switch. false: disabled; true: enabled.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudClsTopicsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cls_topics.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := ClsService{client: meta.(*TencentCloudClient).apiV3Conn}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*cls.Filter, 0, len(filtersSet))
		for _, item := range filtersSet {
			filtersMap := item.(map[string]interface{})
			filter := cls.Filter{}
			if v, ok := filtersMap["key"].(string); ok && v != "" {
				filter.Key = helper.String(v)
			}
			if v, ok := filtersMap["values"]; ok {
				valuesSet := v.(*schema.Set).List()
				for i := range valuesSet {
					values := valuesSet[i].(string)
					filter.Values = append(filter.Values, helper.String(values))
				}
			}
			tmpSet = append(tmpSet, &filter)
		}
		paramMap["Filters"] = tmpSet
	}

	if v, ok := d.GetOkExists("precise_search"); ok {
		paramMap["PreciseSearch"] = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOkExists("biz_type"); ok {
		paramMap["BizType"] = helper.IntUint64(v.(int))
	}

	var respData []*cls.TopicInfo
	reqErr := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeClsTopicsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		respData = result
		return nil
	})

	if reqErr != nil {
		return reqErr
	}

	ids := make([]string, 0, len(respData))
	topicsList := make([]map[string]interface{}, 0, len(respData))
	if respData != nil {
		for _, topics := range respData {
			topicsMap := map[string]interface{}{}
			if topics.LogsetId != nil {
				topicsMap["logset_id"] = topics.LogsetId
			}
			if topics.TopicId != nil {
				topicsMap["topic_id"] = topics.TopicId
				ids = append(ids, *topics.TopicId)
			}
			if topics.TopicName != nil {
				topicsMap["topic_name"] = topics.TopicName
			}
			if topics.PartitionCount != nil {
				topicsMap["partition_count"] = topics.PartitionCount
			}
			if topics.Index != nil {
				topicsMap["index"] = topics.Index
			}
			if topics.AssumerName != nil {
				topicsMap["assumer_name"] = topics.AssumerName
			}
			if topics.CreateTime != nil {
				topicsMap["create_time"] = topics.CreateTime
			}
			if topics.Status != nil {
				topicsMap["status"] = topics.Status
			}

			tagsList := make([]map[string]interface{}, 0, len(topics.Tags))
			if topics.Tags != nil {
				for _, tags := range topics.Tags {
					tagsMap := map[string]interface{}{}
					if tags.Key != nil {
						tagsMap["key"] = tags.Key
					}
					if tags.Value != nil {
						tagsMap["value"] = tags.Value
					}
					tagsList = append(tagsList, tagsMap)
				}
				topicsMap["tags"] = tagsList
			}

			if topics.AutoSplit != nil {
				topicsMap["auto_split"] = topics.AutoSplit
			}
			if topics.MaxSplitPartitions != nil {
				topicsMap["max_split_partitions"] = topics.MaxSplitPartitions
			}
			if topics.StorageType != nil {
				topicsMap["storage_type"] = topics.StorageType
			}
			if topics.Period != nil {
				topicsMap["period"] = topics.Period
			}
			if topics.SubAssumerName != nil {
				topicsMap["sub_assumer_name"] = topics.SubAssumerName
			}
			if topics.Describes != nil {
				topicsMap["describes"] = topics.Describes
			}
			if topics.HotPeriod != nil {
				topicsMap["hot_period"] = topics.HotPeriod
			}
			if topics.BizType != nil {
				topicsMap["biz_type"] = topics.BizType
			}
			if topics.IsWebTracking != nil {
				topicsMap["is_web_tracking"] = topics.IsWebTracking
			}

			topicsList = append(topicsList, topicsMap)
		}

		_ = d.Set("topics", topicsList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), topicsList); e != nil {
			return e
		}
	}

	return nil
}

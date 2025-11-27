// /*
// Use this data source to query detailed information of ckafka datahub_group_offsets
//
// # Example Usage
//
// ```hcl
// data "tencentcloudenterprise_ckafka_datahub_group_offsets" "datahub_group_offsets" {
// }
// ```
// */
package tencentcloud

//
//import (
//	"context"
//
//	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	ckafka "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka/v20190819"
//)
//
//func init() {
//	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_datahub_group_offsets", CNDescription{
//		TerraformTypeCN: "ckafka_datahub_group_offsets",
//		AttributesCN: map[string]string{
//			"name":               "任务订阅的主题名称",
//			"group":              "卡夫卡消费群体",
//			"search_word":        "模糊匹配topicName",
//			"topic_list":         "主题数组，其中每个元素都是一个json对象",
//			"topic":              "主题名称",
//			"partitions":         "主题分区数组，其中每个元素都是一个json对象",
//			"partition":          "主题分区Id",
//			"offset":             "消费者补偿",
//			"metadata":           "通常是一个空字符串",
//			"error_code":         "错误代码",
//			"log_end_offset":     "分区日志结束偏移",
//			"lag":                "未使用的邮件数",
//			"result_output_file": "用于保存结果",
//		},
//	})
//}
//
//func dataSourceTencentCloudCkafkaDatahubGroupOffsets() *schema.Resource {
//	return &schema.Resource{
//		Description: "Use this data source to query detailed information of ckafka datahub_group_offsets",
//		Read:        dataSourceTencentCloudCkafkaDatahubGroupOffsetsRead,
//		Schema: map[string]*schema.Schema{
//			"name": {
//				Required:    true,
//				Type:        schema.TypeString,
//				Description: "Topic name that the task subscribe.",
//			},
//
//			"group": {
//				Required:    true,
//				Type:        schema.TypeString,
//				Description: "Kafka consumer group.",
//			},
//
//			"search_word": {
//				Optional:    true,
//				Type:        schema.TypeString,
//				Description: "Fuzzy match topicName.",
//			},
//
//			"topic_list": {
//				Type:        schema.TypeList,
//				Computed:    true,
//				Description: "The topic array, where each element is a json object.",
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"topic": {
//							Type:        schema.TypeString,
//							Computed:    true,
//							Description: "Topic name.",
//						},
//						"partitions": {
//							Type:        schema.TypeList,
//							Computed:    true,
//							Description: "The topic partition array, where each element is a json object.",
//							Elem: &schema.Resource{
//								Schema: map[string]*schema.Schema{
//									"partition": {
//										Type:        schema.TypeInt,
//										Computed:    true,
//										Description: "Topic partitionId.",
//									},
//									"offset": {
//										Type:        schema.TypeInt,
//										Computed:    true,
//										Description: "Consumer offset.",
//									},
//									"metadata": {
//										Type:        schema.TypeString,
//										Computed:    true,
//										Description: "Usually an empty string.",
//									},
//									"error_code": {
//										Type:        schema.TypeInt,
//										Computed:    true,
//										Description: "Error Code.",
//									},
//									"log_end_offset": {
//										Type:        schema.TypeInt,
//										Computed:    true,
//										Description: "Partition Log End Offset.",
//									},
//									"lag": {
//										Type:        schema.TypeInt,
//										Computed:    true,
//										Description: "The number of unconsumed messages.",
//									},
//								},
//							},
//						},
//					},
//				},
//			},
//
//			"result_output_file": {
//				Type:        schema.TypeString,
//				Optional:    true,
//				Description: "Used to save results.",
//			},
//		},
//	}
//}
//
//func dataSourceTencentCloudCkafkaDatahubGroupOffsetsRead(d *schema.ResourceData, meta interface{}) error {
//	defer logElapsed("data_source.cloud_ckafka_datahub_group_offsets.read")()
//	defer inconsistentCheck(d, meta)()
//
//	logId := getLogId(contextNil)
//
//	ctx := context.WithValue(context.TODO(), logIdKey, logId)
//
//	paramMap := make(map[string]interface{})
//	if v, ok := d.GetOk("name"); ok {
//		paramMap["name"] = helper.String(v.(string))
//	}
//
//	if v, ok := d.GetOk("group"); ok {
//		paramMap["group"] = helper.String(v.(string))
//	}
//
//	if v, ok := d.GetOk("search_word"); ok {
//		paramMap["search_word"] = helper.String(v.(string))
//	}
//
//	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}
//
//	var result []*ckafka.GroupOffsetTopic
//
//	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
//		groupOffsetTopics, e := service.DescribeCkafkaDatahubGroupOffsetsByFilter(ctx, paramMap)
//		if e != nil {
//			return retryError(e)
//		}
//		result = groupOffsetTopics
//		return nil
//	})
//	if err != nil {
//		return err
//	}
//
//	ids := make([]string, 0, len(result))
//	topicList := make([]map[string]interface{}, 0, len(result))
//	for _, topic := range result {
//		topicMap := make(map[string]interface{})
//
//		if topic.Topic != nil {
//			topicMap["topic"] = topic.Topic
//			ids = append(ids, *topic.Topic)
//
//		}
//
//		if topic.Partitions != nil {
//			partitionsList := make([]map[string]interface{}, 0)
//			for _, partitions := range topic.Partitions {
//				partitionsMap := map[string]interface{}{}
//
//				if partitions.Partition != nil {
//					partitionsMap["partition"] = partitions.Partition
//				}
//
//				if partitions.Offset != nil {
//					partitionsMap["offset"] = partitions.Offset
//				}
//
//				if partitions.Metadata != nil {
//					partitionsMap["metadata"] = partitions.Metadata
//				}
//
//				if partitions.ErrorCode != nil {
//					partitionsMap["error_code"] = partitions.ErrorCode
//				}
//
//				if partitions.LogEndOffset != nil {
//					partitionsMap["log_end_offset"] = partitions.LogEndOffset
//				}
//
//				if partitions.Lag != nil {
//					partitionsMap["lag"] = partitions.Lag
//				}
//
//				partitionsList = append(partitionsList, partitionsMap)
//			}
//
//			topicMap["partitions"] = partitionsList
//		}
//
//		topicList = append(topicList, topicMap)
//
//	}
//
//	d.SetId(helper.DataResourceIdsHash(ids))
//	_ = d.Set("topic_list", topicList)
//
//	output, ok := d.GetOk("result_output_file")
//	if ok && output.(string) != "" {
//		if e := writeToFile(output.(string), topicList); e != nil {
//			return e
//		}
//	}
//	return nil
//}

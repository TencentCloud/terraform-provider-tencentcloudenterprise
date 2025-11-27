/*
Use this data source to query detailed information of ckafka group_offsets

# Example Usage

```hcl

	data "tencentcloudenterprise_ckafka_group_offsets" "group_offsets" {
	  instance_id = "ckafka-xxxxxx"
	  group = "xxxxxx"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_group_offsets", CNDescription{
		TerraformTypeCN: "消费分组offset",
		DescriptionCN:   "提供Ckafka消费组Offset数据源，用于查询Ckafka消费组Offset的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "（过滤条件）按照实例 ID 过滤",
			"group":              "Kafka 消费分组",
			"topics":             "group 订阅的主题名称数组，如果没有该数组，则表示指定的 group 下所有 topic 信息",
			"search_word":        "模糊匹配 topicName",
			"topic_list":         "主题数组，其中每个元素为一个 json object",
			"topic":              "主题名称",
			"partitions":         "该主题分区数组，其中每个元素为一个 json object",
			"partition":          "topic 的 partitionId",
			"offset":             "consumer 提交的 offset 位置",
			"metadata":           "支持消费者提交消息时，传入 metadata 作为它用，当前一般为空字符串",
			"error_code":         "错误码",
			"log_end_offset":     "当前 partition 最新的 offset",
			"lag":                "未消费的消息个数",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudCkafkaGroupOffsets() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of ckafka group_offsets",
		Read:        dataSourceTencentCloudCkafkaGroupOffsetsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "InstanceId.",
			},

			"group": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Kafka consumer group name.",
			},

			"topics": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "An array of topic names subscribed by the group, if there is no such array, it means all topic information under the specified group.",
			},

			"search_word": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Fuzzy match topicName.",
			},

			"topic_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The topic array, where each element is a json object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "TopicName.",
						},
						"partitions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The topic partition array, where each element is a json object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"partition": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Topic partitionId.",
									},
									"offset": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The offset of the position.",
									},
									"metadata": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "When consumers submit messages, they can pass in metadata for other purposes. Currently, it is usually an empty string.",
									},
									"error_code": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "ErrorCode.",
									},
									"log_end_offset": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The latest offset of the current partition.",
									},
									"lag": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of unconsumed messages.",
									},
								},
							},
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

func dataSourceTencentCloudCkafkaGroupOffsetsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_group_offsets.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("group"); ok {
		paramMap["group"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("topics"); ok {
		topicsSet := v.(*schema.Set).List()
		paramMap["topics"] = helper.InterfacesStringsPoint(topicsSet)
	}

	if v, ok := d.GetOk("search_word"); ok {
		paramMap["search_word"] = helper.String(v.(string))
	}

	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}

	var groupOffsetTopics []*ckafka.GroupOffsetTopic

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCkafkaGroupOffsetsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		groupOffsetTopics = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(groupOffsetTopics))
	groupOffsetResponseMap := map[string]interface{}{}

	topicList := []interface{}{}
	for _, topic := range groupOffsetTopics {
		topicMap := map[string]interface{}{}

		if topic.Topic != nil {
			topicMap["topic"] = topic.Topic
			ids = append(ids, *topic.Topic)

		}

		if topic.Partitions != nil {
			partitionsList := []interface{}{}
			for _, partitions := range topic.Partitions {
				partitionsMap := map[string]interface{}{}

				if partitions.Partition != nil {
					partitionsMap["partition"] = partitions.Partition
				}

				if partitions.Offset != nil {
					partitionsMap["offset"] = partitions.Offset
				}

				if partitions.Metadata != nil {
					partitionsMap["metadata"] = partitions.Metadata
				}

				if partitions.ErrorCode != nil {
					partitionsMap["error_code"] = partitions.ErrorCode
				}

				if partitions.LogEndOffset != nil {
					partitionsMap["log_end_offset"] = partitions.LogEndOffset
				}

				if partitions.Lag != nil {
					partitionsMap["lag"] = partitions.Lag
				}

				partitionsList = append(partitionsList, partitionsMap)
			}

			topicMap["partitions"] = partitionsList
		}

		topicList = append(topicList, topicMap)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("topic_list", topicList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), groupOffsetResponseMap); e != nil {
			return e
		}
	}
	return nil
}

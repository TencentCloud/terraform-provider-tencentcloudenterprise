/*
Use this data source to query detailed information of tdmq message

# Example Usage

```hcl

	data "tencentcloudenterprise_tdmq_rocketmq_messages" "message" {
	  cluster_id     = "rocketmq-rkrbm52djmro"
	  environment_id = "keep_ns"
	  topic_name     = "keep-topic"
	  msg_id         = "A9FE8D0567FE15DB97425FC08EEF0000"
	  query_dlq_msg  = false
	}

```
*/
package tencentcloud

import (
	"context"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rocketmq_messages", CNDescription{
		TerraformTypeCN: "rocketmq消息",
		DescriptionCN:   "提供TDMQ RocketMQ消息数据源，用于查询TDMQ RocketMQ消息的详细信息。",
		AttributesCN: map[string]string{
			"cluster_id":         "群集id",
			"environment_id":     "环境",
			"topic_name":         "Topic，查询死信时传递groupId",
			"msg_id":             "消息ID",
			"query_dlq_msg":      "查询死信时该值为true，仅对Rocketmq有效",
			"body":               "消息正文",
			"properties":         "详细参数",
			"produce_time":       "生产时间",
			"producer_addr":      "生产者地址",
			"message_tracks":     "消费者组消费注意：此字段可能返回null，表示无法获得有效值",
			"group":              "消费者群体",
			"consume_status":     "消费状态",
			"track_type":         "消息跟踪类型",
			"exception_desc":     "异常信息注意：此字段可能返回null，表示无法获得有效值",
			"show_topic_name":    "详细信息页面上显示的主题名称注意：此字段可能返回null，表示无法获得有效值",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTdmqRocketmqMessages() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tdmq message",
		Read:        dataSourceTencentCloudTdmqMessageRead,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Cluster id.",
			},
			"environment_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Environment.",
			},
			"topic_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Topic, groupId is passed when querying dead letters.",
			},
			"msg_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Message ID.",
			},
			"query_dlq_msg": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "The value is true when querying dead letters, only valid for Rocketmq.",
			},
			"body": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Message body.",
			},
			"properties": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Detailed parameters.",
			},
			"produce_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Production time.",
			},
			"producer_addr": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Producer address.",
			},
			"message_tracks": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Consumer Group ConsumptionNote: This field may return null, indicating that no valid value can be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Consumer group.",
						},
						"consume_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Consumption status.",
						},
						"track_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message track type.",
						},
						"exception_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Exception informationNote: This field may return null, indicating that no valid value can be obtained.",
						},
					},
				},
			},
			"show_topic_name": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The topic name displayed on the details pageNote: This field may return null, indicating that no valid value can be obtained.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudTdmqMessageRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmq_rocketmq_messages.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId     = getLogId(contextNil)
		ctx       = context.WithValue(context.TODO(), logIdKey, logId)
		service   = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		msgDetail *tdmq.DescribeRocketMQMsgResponse
		clusterId string
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("cluster_id"); ok {
		paramMap["ClusterId"] = helper.String(v.(string))
		clusterId = v.(string)
	}

	if v, ok := d.GetOk("environment_id"); ok {
		paramMap["EnvironmentId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("topic_name"); ok {
		paramMap["TopicName"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("msg_id"); ok {
		paramMap["MsgId"] = helper.String(v.(string))
	}

	if v, _ := d.GetOk("query_dlq_msg"); v != nil {
		paramMap["QueryDlqMsg"] = helper.Bool(v.(bool))
	}

	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTdmqMessageByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		msgDetail = result
		return nil
	})

	if err != nil {
		return err
	}

	detail := msgDetail.Response

	if detail.Body != nil {
		_ = d.Set("body", detail.Body)
	}

	if detail.Properties != nil {
		_ = d.Set("properties", detail.Properties)
	}

	if detail.ProduceTime != nil {
		_ = d.Set("produce_time", detail.ProduceTime)
	}

	if detail.ProducerAddr != nil {
		_ = d.Set("producer_addr", detail.ProducerAddr)
	}

	if detail.MessageTracks != nil {
		tmpList := []interface{}{}
		for _, rocketMQMessageTrack := range detail.MessageTracks {
			rocketMQMessageTrackMap := map[string]interface{}{}

			if rocketMQMessageTrack.Group != nil {
				rocketMQMessageTrackMap["group"] = rocketMQMessageTrack.Group
			}

			if rocketMQMessageTrack.ConsumeStatus != nil {
				rocketMQMessageTrackMap["consume_status"] = rocketMQMessageTrack.ConsumeStatus
			}

			if rocketMQMessageTrack.TrackType != nil {
				rocketMQMessageTrackMap["track_type"] = rocketMQMessageTrack.TrackType
			}

			if rocketMQMessageTrack.ExceptionDesc != nil {
				rocketMQMessageTrackMap["exception_desc"] = rocketMQMessageTrack.ExceptionDesc
			}

			tmpList = append(tmpList, rocketMQMessageTrackMap)
		}

		_ = d.Set("message_tracks", tmpList)
	}

	if detail.ShowTopicName != nil {
		_ = d.Set("show_topic_name", detail.ShowTopicName)
	}

	d.SetId(clusterId)
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), d); e != nil {
			return e
		}
	}

	return nil
}

/*
Use this data source to query detailed information of ckafka topic.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ckafka_topic" "foo" {
		instance_id                     = "ckafka-f9ife4zz"
		topic_name                      = "example"
		note                            = "topic note"
		replica_num                     = 2
		partition_num                   = 1
		enable_white_list               = true
		ip_white_list                   = ["ip1","ip2"]
		clean_up_policy                 = "delete"
		sync_replica_min_num            = 1
		unclean_leader_election_enable  = false
		segment                         = 3600000
		retention                       = 60000
		max_message_bytes               = 1024
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_topics", CNDescription{
		TerraformTypeCN: "ckafka主题",
		DescriptionCN:   "提供CKafka主题数据源，用于查询CKafka主题的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":                    "Ckafka实例ID",
			"topic_name":                     "CKafka主题的名称它必须以字母开头，其余部分可以包含字母、数字和短划线（-）长度范围从1到64",
			"result_output_file":             "用于存储结果",
			"topic_list":                     "返回的主题详情列表",
			"topic_id":                       "CKafka主题的ID",
			"partition_num":                  "分区的数目",
			"replica_num":                    "复制副本的数量",
			"note":                           "CKafka主题说明",
			"create_time":                    "创建CKafka主题的时间",
			"enable_white_list":              "是否打开IP白名单`true`：打开，false`：关闭",
			"ip_white_list_count":            "IP白名单计数",
			"forward_interval":               "定期将数据备份到cos",
			"forward_cos_bucket":             "数据备份cos bucket：转储到cos的bucket地址",
			"forward_status":                 "数据备份cos状态`1`：不打开数据备份，`0`：打开数据备份",
			"retention":                      "可以选择消息保留时间（单位ms）",
			"sync_replica_min_num":           "同步副本的最小数目",
			"clean_up_policy":                "清除日志策略，日志清除模式`delete：日志根据存储时间删除，compact：日志根据密钥压缩，compact，delete：日志按照密钥压缩，并将根据存储时间进行删除",
			"unclean_leader_election_enable": "是否允许选择未同步的副本作为前导，默认为“false”、“true:”允许，“false”：不允许",
			"max_message_bytes":              "最大消息字节数",
			"segment":                        "段滚动时间，以毫秒为单位",
			"segment_bytes":                  "碎片滚动的字节数",
		},
	})
}

func dataSourceTencentCloudCkafkaTopics() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides the CKafka topic details of the specified instance.",
		Read:        dataSourceTencentCloudCkafkaTopicsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Ckafka instance ID.",
			},
			"topic_name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringLengthInRange(1, 64),
				Description:  "Name of the CKafka topic. It must start with a letter, the rest can contain letters, numbers and dashes(-). The length range is from 1 to 64.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to store results.",
			},
			"topic_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of instances. Each element contains the following attributes.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the CKafka topic.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the CKafka topic.",
						},
						"partition_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of partition.",
						},
						"replica_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of replica.",
						},
						"note": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CKafka topic note description.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of the CKafka topic.",
						},
						"enable_white_list": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to open the IP Whitelist. `true`: open, `false`: close.",
						},
						"ip_white_list_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IP Whitelist count.",
						},
						"forward_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Periodic frequency of data backup to cos.",
						},
						"forward_cos_bucket": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Data backup cos bucket: the bucket address that is dumped to cos.",
						},
						"forward_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Data backup cos status. `1`: do not open data backup, `0`: open data backup.",
						},
						"retention": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Message can be selected. Retention time(unit ms).",
						},
						"sync_replica_min_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Min number of sync replicas.",
						},
						"clean_up_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Clear log policy, log clear mode. `delete`: logs are deleted according to the storage time, `compact`: logs are compressed according to the key, `compact, delete`: logs are compressed according to the key and will be deleted according to the storage time.",
						},
						"unclean_leader_election_enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to allow unsynchronized replicas to be selected as leader, default is `false`, `true: `allowed, `false`: not allowed.",
						},
						"max_message_bytes": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max message bytes.",
						},
						"segment": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Segment scrolling time, in ms.",
						},
						"segment_bytes": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of bytes rolled by shard.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCkafkaTopicsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_topics.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	instanceId := d.Get("instance_id").(string)
	topicName := d.Get("topic_name").(string)
	ckafkcService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	topicDetails, err := ckafkcService.DescribeCkafkaTopics(ctx, instanceId, topicName)
	if err != nil {
		return err
	}

	instanceList := make([]map[string]interface{}, 0, len(topicDetails))
	ids := make([]string, 0, len(topicDetails))

	for _, topic := range topicDetails {
		var uncleanLeaderElectionEnable bool
		if topic.Config.UncleanLeaderElectionEnable != nil {
			uncleanLeaderElectionEnable = *topic.Config.UncleanLeaderElectionEnable != 0
		}
		instance := map[string]interface{}{
			"topic_name":                     topic.TopicName,
			"topic_id":                       topic.TopicId,
			"partition_num":                  topic.PartitionNum,
			"replica_num":                    topic.ReplicaNum,
			"note":                           topic.Note,
			"create_time":                    helper.FormatUnixTime(uint64(*topic.CreateTime)),
			"enable_white_list":              topic.EnableWhiteList,
			"ip_white_list_count":            topic.IpWhiteListCount,
			"forward_interval":               topic.ForwardInterval,
			"forward_cos_bucket":             topic.ForwardCosBucket,
			"forward_status":                 topic.ForwardStatus,
			"retention":                      topic.Config.Retention,
			"sync_replica_min_num":           topic.Config.MinInsyncReplicas,
			"clean_up_policy":                topic.Config.CleanUpPolicy,
			"unclean_leader_election_enable": uncleanLeaderElectionEnable,
			"max_message_bytes":              topic.Config.MaxMessageBytes,
			"segment":                        topic.Config.SegmentMs,
			"segment_bytes":                  topic.Config.SegmentBytes,
		}
		resourceId := instanceId + FILED_SP + *topic.TopicName
		instanceList = append(instanceList, instance)
		ids = append(ids, resourceId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if err = d.Set("topic_list", instanceList); err != nil {
		log.Printf("[CRITAL]%s provider set ckafka topic list fail, reason:%s\n ", logId, err.Error())
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), instanceList); err != nil {
			return err
		}
	}

	return nil
}

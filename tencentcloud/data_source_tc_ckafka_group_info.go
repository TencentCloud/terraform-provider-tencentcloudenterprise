/*
Use this data source to query detailed information of ckafka group_info

# Example Usage

```hcl

	data "tencentcloudenterprise_ckafka_group_info" "group_info" {
	  instance_id = "ckafka-xxxxxx"
	  group_list = ["xxxxxx"]
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
	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_group_info", CNDescription{
		TerraformTypeCN: "ckafka消费分组信息",
		DescriptionCN:   "提供Ckafka消费组信息数据源，用于查询Ckafka消费组的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "（过滤条件）按照实例 ID 过滤。",
			"group_list":         "Kafka 消费分组，Consumer-group，这里是数组形式。如果使用拼接数组元素方式传值（一般 GET 请求用），可以使用格式：GroupList.0=xxx&GroupList.1=yyy。",
			"result":             "返回的消费分组信息",
			"error_code":         "错误代码，通常为0",
			"state":              "组状态描述（通常为Empty、Stable和Dead状态）：Dead:消费组不存在Empty：消费组当前没有任何消费者订阅PreparingRebalance:消费组处于重新平衡状态CompletingRebalance：消费组处于再平衡状态Stable：消费组中的每个消费者都已加入并处于稳定状态",
			"protocol_type":      "消费组选择的协议类型通常是消费者，但有些系统使用自己的协议，例如kafka connect，它使用connect只有标准的消费者协议，这个接口才知道具体分配方法的格式，并且可以分析具体的分区分配",
			"protocol":           "常见的consumer分区分配算法如下（Kafka consumer SDK默认选项为range）range|roundrobin|stick",
			"members":            "仅当state为Stable并且protocol_type为consumer时，此数组才包含信息",
			"member_id":          "coordinator 为消费分组中的消费者生成的唯一 ID",
			"client_id":          "客户消费者 SDK 自己设置的 client.id 信息",
			"client_host":        "一般存储客户的 IP 地址",
			"assignment":         "存储着分配给该消费者的 partition 信息",
			"version":            "assingment版本信息",
			"topics":             "topic信息列表",
			"topic":              "分配的 topic 名称",
			"partitions":         "分配的 partition 信息",
			"group":              "ckafka消费组",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudCkafkaGroupInfo() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of ckafka group_info",
		Read:        dataSourceTencentCloudCkafkaGroupInfoRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "InstanceId.",
			},

			"group_list": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Kafka consumption group, Consumer-group, here is an array format, format GroupList.0=xxx&amp;amp;GroupList.1=yyy.",
			},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Result.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error code, normally 0.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Group state description (commonly Empty, Stable, and Dead states): Dead: The consumption group does not exist Empty: The consumption group does not currently have any consumer subscriptions PreparingRebalance: The consumption group is in the rebalance state CompletingRebalance: The consumption group is in the rebalance state Stable: Each consumer in the consumption group has joined and is in a stable state.",
						},
						"protocol_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol type selected by the consumption group is normally the consumer, but some systems use their own protocol, such as kafka-connect, which uses connect. Only the standard consumer protocol, this interface knows the format of the specific allocation method, and can analyze the specific partition allocation.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Common consumer partition allocation algorithms are as follows (the default option for Kafka consumer SDK is range)  range|roundrobin|sticky.",
						},
						"members": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "This array contains information only if state is Stable and protocol_type is consumer.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"member_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID that the coordinator generated for consumer.",
									},
									"client_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The client.id information set by the client consumer SDK itself.",
									},
									"client_host": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Generally store the customer&#39;s IP address.",
									},
									"assignment": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Stores the partition information assigned to the consumer.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"version": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Assignment version information.",
												},
												"topics": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Topic list.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"topic": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Assigned topic name.",
															},
															"partitions": {
																Type: schema.TypeSet,
																Elem: &schema.Schema{
																	Type: schema.TypeInt,
																},
																Computed:    true,
																Description: "Allocated partition information.",
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
						"group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Kafka consumer group.",
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

func dataSourceTencentCloudCkafkaGroupInfoRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_group_info.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("group_list"); ok {
		groupListSet := v.(*schema.Set).List()
		paramMap["group_list"] = helper.InterfacesStringsPoint(groupListSet)
	}

	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}

	var result []*ckafka.GroupInfoResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		groupInfo, e := service.DescribeCkafkaGroupInfoByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		result = groupInfo
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(result))
	tmpList := make([]map[string]interface{}, 0, len(result))

	if result != nil {
		for _, groupInfoResponse := range result {
			groupInfoResponseMap := map[string]interface{}{}

			if groupInfoResponse.ErrorCode != nil {
				groupInfoResponseMap["error_code"] = groupInfoResponse.ErrorCode
			}

			if groupInfoResponse.State != nil {
				groupInfoResponseMap["state"] = groupInfoResponse.State
			}

			if groupInfoResponse.ProtocolType != nil {
				groupInfoResponseMap["protocol_type"] = groupInfoResponse.ProtocolType
			}

			if groupInfoResponse.Protocol != nil {
				groupInfoResponseMap["protocol"] = groupInfoResponse.Protocol
			}

			if groupInfoResponse.Members != nil {
				membersList := []interface{}{}
				for _, members := range groupInfoResponse.Members {
					membersMap := map[string]interface{}{}

					if members.MemberId != nil {
						membersMap["member_id"] = members.MemberId
					}

					if members.ClientId != nil {
						membersMap["client_id"] = members.ClientId
					}

					if members.ClientHost != nil {
						membersMap["client_host"] = members.ClientHost
					}

					if members.Assignment != nil {
						assignmentMap := map[string]interface{}{}

						if members.Assignment.Version != nil {
							assignmentMap["version"] = members.Assignment.Version
						}

						if members.Assignment.Topics != nil {
							topicsList := []interface{}{}
							for _, topics := range members.Assignment.Topics {
								topicsMap := map[string]interface{}{}

								if topics.Topic != nil {
									topicsMap["topic"] = topics.Topic
								}

								if topics.Partitions != nil {
									topicsMap["partitions"] = topics.Partitions
								}

								topicsList = append(topicsList, topicsMap)
							}

							assignmentMap["topics"] = topicsList
						}

						membersMap["assignment"] = []interface{}{assignmentMap}
					}

					membersList = append(membersList, membersMap)
				}

				groupInfoResponseMap["members"] = membersList
			}

			if groupInfoResponse.Group != nil {
				groupInfoResponseMap["group"] = groupInfoResponse.Group
				ids = append(ids, *groupInfoResponse.Group)
			}

			tmpList = append(tmpList, groupInfoResponseMap)
		}

		_ = d.Set("result", tmpList)
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

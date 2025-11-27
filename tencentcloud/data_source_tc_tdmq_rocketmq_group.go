/*
Use this data source to query detailed information of tdmqRocketmq group

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
	  cluster_name = "test_rocketmq_datasource_group"
	  remark = "test recket mq"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	  namespace_name = "test_namespace_datasource"
	  ttl = 65000
	  retention_time = 65000
	  remark = "test namespace"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_group" "group" {
	  group_name = "test_rocketmq_group"
	  namespace = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
	  read_enable = true
	  broadcast_enable = true
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	  remark = "test rocketmq group"
	}

	data "tencentcloudenterprise_tdmq_rocketmq_group" "group" {
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	  namespace_id = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
	  filter_group = cloud_tdmq_rocketmq_group.group.group_name
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmqRocketmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rocketmq_group", CNDescription{
		TerraformTypeCN: "RocketMQ 消费组",
		DescriptionCN:   "提供TDMQ RocketMQ消费组数据源，用于查询TDMQ RocketMQ消费组的详细信息。",
		AttributesCN: map[string]string{
			"cluster_id":          "群集ID",
			"namespace_id":        "命名空间",
			"filter_topic":        "主题名称，可用于查询该主题下的所有订阅组",
			"filter_group":        "按使用者组名称查询使用者组支持模糊查询",
			"filter_one_group":    "订阅组名称指定后，只返回该订阅组的信息",
			"groups":              "订阅组的列表",
			"name":                "消费者组名称",
			"consumer_num":        "在线消费者的数量",
			"tps":                 "消耗TPS",
			"total_accumulative":  "堆积的邮件总数",
			"consumption_mode":    "`0`：集群消费模式`1`：广播消费模式`-1`：未知",
			"read_enable":         "是否启用消费",
			"retry_partition_num": "重试主题中的分区数",
			"create_time":         "创建时间（以毫秒为单位）",
			"update_time":         "修改时间（以毫秒为单位）",
			"client_protocol":     "客户端协议",
			"remark":              "备注（最多128个字符）",
			"consumer_type":       "消费者类型枚举的值：ACTIVELY或PASSIVELY",
			"broadcast_enable":    "是否启用广播消费",
			"result_output_file":  "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTdmqRocketmqGroup() *schema.Resource {
	return &schema.Resource{
		Description: "Cloud TDMQ RocketMQ Group.",
		Read:        dataSourceTencentCloudTdmqRocketmqGroupRead,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster ID.",
			},

			"namespace_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Namespace.",
			},

			"filter_topic": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Topic name, which can be used to query all subscription groups under the topic.",
			},

			"filter_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Consumer group query by consumer group name. Fuzzy query is supported.",
			},

			"filter_one_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Subscription group name. After it is specified, the information of only this subscription group will be returned.",
			},

			"groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of subscription groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Consumer group name.",
						},
						"consumer_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of online consumers.",
						},
						"tps": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Consumption TPS.",
						},
						"total_accumulative": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of heaped messages.",
						},
						"consumption_mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "`0`: Cluster consumption mode; `1`: Broadcast consumption mode; `-1`: Unknown.",
						},
						"read_enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable consumption.",
						},
						"retry_partition_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of partitions in a retry topic.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Creation time in milliseconds.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Modification time in milliseconds.",
						},
						"client_protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Client protocol.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remarks (up to 128 characters).",
						},
						"consumer_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Consumer type. Enumerated values: ACTIVELY or PASSIVELY.",
						},
						"broadcast_enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable broadcast consumption.",
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

func dataSourceTencentCloudTdmqRocketmqGroupRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmqRocketmq_group.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("cluster_id"); ok {
		paramMap["cluster_id"] = v.(string)
	}

	if v, ok := d.GetOk("namespace_id"); ok {
		paramMap["namespace_id"] = v.(string)
	}

	if v, ok := d.GetOk("filter_topic"); ok {
		paramMap["filter_topic"] = v.(string)
	}

	if v, ok := d.GetOk("filter_group"); ok {
		paramMap["filter_group"] = v.(string)
	}

	if v, ok := d.GetOk("filter_one_group"); ok {
		paramMap["filter_one_group"] = v.(string)
	}

	tdmqRocketmqService := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var groups []*tdmqRocketmq.RocketMQGroup
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := tdmqRocketmqService.DescribeTdmqRocketmqGroupByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		groups = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read TdmqRocketmq groups failed, reason:%+v", logId, err)
		return err
	}

	groupList := []interface{}{}
	ids := make([]string, 0)
	for _, group := range groups {
		groupMap := map[string]interface{}{}
		if group.Name != nil {
			groupMap["name"] = group.Name
		}
		if group.ConsumerNum != nil {
			groupMap["consumer_num"] = group.ConsumerNum
		}
		if group.TPS != nil {
			groupMap["tps"] = group.TPS
		}
		if group.TotalAccumulative != nil {
			groupMap["total_accumulative"] = group.TotalAccumulative
		}
		if group.ConsumptionMode != nil {
			groupMap["consumption_mode"] = group.ConsumptionMode
		}
		if group.ReadEnabled != nil {
			groupMap["read_enable"] = group.ReadEnabled
		}
		if group.RetryPartitionNum != nil {
			groupMap["retry_partition_num"] = group.RetryPartitionNum
		}
		if group.CreateTime != nil {
			groupMap["create_time"] = group.CreateTime
		}
		if group.UpdateTime != nil {
			groupMap["update_time"] = group.UpdateTime
		}
		if group.ClientProtocol != nil {
			groupMap["client_protocol"] = group.ClientProtocol
		}
		if group.Remark != nil {
			groupMap["remark"] = group.Remark
		}
		if group.ConsumerType != nil {
			groupMap["consumer_type"] = group.ConsumerType
		}
		if group.BroadcastEnabled != nil {
			groupMap["broadcast_enable"] = group.BroadcastEnabled
		}

		groupList = append(groupList, groupMap)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("groups", groupList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), groupList); e != nil {
			return e
		}
	}

	return nil
}

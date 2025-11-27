/*
Use this data source to query detailed information of tdmqRocketmq topic

# Example Usage

```hcl

	resource "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
		cluster_name = "test_rocketmq"
		remark = "test recket mq"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_namespace" "namespace" {
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	  namespace_name = "test_namespace"
	  ttl = 65000
	  retention_time = 65000
	  remark = "test namespace"
	}

	resource "tencentcloudenterprise_tdmq_rocketmq_topic" "topic" {
	  topic_name = "test_rocketmq_topic"
	  namespace_name = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
	  type = "Normal"
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	  remark = "test rocketmq topic"
	}

	data "tencentcloudenterprise_tdmq_rocketmq_topic" "topic" {
	  cluster_id = cloud_tdmq_rocketmq_cluster.cluster.cluster_id
	  namespace_id = cloud_tdmq_rocketmq_namespace.namespace.namespace_name
	  filter_name = cloud_tdmq_rocketmq_topic.topic.topic_name
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
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rocketmq_topic", CNDescription{
		TerraformTypeCN: "RocketMQ Topic",
		DescriptionCN:   "提供TDMQ RocketMQ主题数据源，用于查询TDMQ RocketMQ主题的详细信息。",
		AttributesCN: map[string]string{
			"cluster_id":         "群集ID",
			"namespace_id":       "命名空间",
			"filter_type":        "按主题类型筛选有效值：“Normal”、“GlobalOrder”、“PartitionedOrder”和“Transaction”",
			"filter_name":        "按主题名称搜索支持模糊查询",
			"topics":             "主题信息列表",
			"name":               "主题名称",
			"remark":             "主题名称",
			"partition_num":      "读/写分区的数目",
			"create_time":        "创建时间（以毫秒为单位）",
			"update_time":        "更新时间（以毫秒为单位）",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTdmqRocketmqTopic() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tdmqRocketmq topic",
		Read:        dataSourceTencentCloudTdmqRocketmqTopicRead,
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

			"filter_type": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Filter by topic type. Valid values: `Normal`, `GlobalOrder`, `PartitionedOrder`, `Transaction`.",
			},

			"filter_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search by topic name. Fuzzy query is supported.",
			},

			"topics": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of topic information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topic name.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topic name.",
						},
						"partition_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of read/write partitions.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Creation time in milliseconds.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Update time in milliseconds.",
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

func dataSourceTencentCloudTdmqRocketmqTopicRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmqRocketmq_topic.read")()
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

	if v, ok := d.GetOk("filter_type"); ok {
		filterTypes := v.(*schema.Set).List()
		filterTypeList := make([]string, 0)
		for _, filterType := range filterTypes {
			filter_type := filterType.(string)
			filterTypeList = append(filterTypeList, filter_type)
		}
		paramMap["filter_type"] = filterTypeList
	}

	if v, ok := d.GetOk("filter_name"); ok {
		paramMap["filter_name"] = v.(string)
	}

	tdmqRocketmqService := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var topics []*tdmqRocketmq.RocketMQTopic
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := tdmqRocketmqService.DescribeTdmqRocketmqTopicByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		topics = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read TdmqRocketmq topics failed, reason:%+v", logId, err)
		return err
	}

	ids := make([]string, 0)
	topicList := []interface{}{}
	for _, topic := range topics {
		topicMap := map[string]interface{}{}
		ids = append(ids, *topic.Name)
		topicMap["name"] = topic.Name
		if topic.Remark != nil {
			topicMap["remark"] = topic.Remark
		}
		if topic.PartitionNum != nil {
			topicMap["partition_num"] = topic.PartitionNum
		}
		if topic.CreateTime != nil {
			topicMap["create_time"] = topic.CreateTime
		}
		if topic.UpdateTime != nil {
			topicMap["update_time"] = topic.UpdateTime
		}

		topicList = append(topicList, topicMap)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("topics", topicList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), topicList); e != nil {
			return e
		}
	}

	return nil
}

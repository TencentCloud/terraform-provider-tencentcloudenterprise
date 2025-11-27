/*
Use this data source to query detailed information of tdmq vip_instance

# Example Usage

```hcl

	data "tencentcloudenterprise_tdmq_vip_instance" "vip_instance" {
	  cluster_id = "rocketmq-rd3545bkkj49"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

//func init() {
//	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_vip_instance", CNDescription{
//		TerraformTypeCN: "TDMQ专享版实例",
//		DescriptionCN:   "提供TDMQ专享版实例数据源，用于查询专享版实例的详细信息。",
//		AttributesCN: map[string]string{
//			"cluster_id":         "集群ID",
//			"result_output_file": "用于保存结果",
//			"instance_list":      "实例列表",
//		},
//	})
//}

func dataSourceTencentCloudTdmqVipInstance() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudTdmqVipInstanceRead,
		Description: "Use this data source to query detailed information of tdmq vip instance",
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Cluster ID.",
			},
			// computed
			"cluster_info": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Cluster information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster ID.",
						},
						"cluster_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster Name.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region.",
						},
						"create_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Creation time, in milliseconds.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster description information. Note: This field may return null, indicating that no valid value can be obtained.",
						},
						"public_end_point": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Public network access address.",
						},
						"vpc_end_point": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC access address.",
						},
						"support_namespace_endpoint": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether namespace access points are supported. Note: This field may return null, indicating that no valid value can be obtained.",
						},
						"vpcs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "VPC information. Note: This field may return null, indicating that no valid value can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "VPC ID.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Subnet Id.",
									},
								},
							},
						},
						"is_vip": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether it is a dedicated instance. Note: This field may return null, indicating that no valid value can be obtained.",
						},
						"rocket_mq_flag": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Rocketmq cluster identification. Note: This field may return null, indicating that no valid value can be obtained.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Billing status, 1 means normal, 2 means stopped, 3 means destroyed. Note: This field may return null, indicating that no valid value can be obtained.",
						},
						"isolate_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Overdue suspension time, in milliseconds. Note: This field may return null, indicating that no valid value can be obtained.",
						},
						"http_public_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP protocol public network access address. Note: This field may return null, indicating that no valid value can be obtained.",
						},
						"http_vpc_endpoint": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTP protocol VPC access address. Note: This field may return null, indicating that no valid value can be obtained.",
						},
					},
				},
			},
			"instance_config": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Cluster configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"max_tps_per_namespace": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Single namespace TPS upper limit.",
						},
						"max_namespace_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of namespaces.",
						},
						"used_namespace_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of used namespaces.",
						},
						"max_topic_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of topics.",
						},
						"used_topic_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of topics used.",
						},
						"max_group_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of groups.",
						},
						"used_group_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of used groups.",
						},
						"config_display": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster type.",
						},
						"node_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of cluster nodes.",
						},
						"node_distribution": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Node distribution.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Availability zone.",
									},
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Availability zone id.",
									},
									"node_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of nodes.",
									},
								},
							},
						},
						"topic_distribution": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Topic distribution.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"topic_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Topic type.",
									},
									"count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of topics.",
									},
								},
							},
						},
						"max_queues_per_topic": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of queues per topic. Note: This field may return null, indicating that no valid value can be obtained.",
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

func dataSourceTencentCloudTdmqVipInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmq_vip_instance.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId       = getLogId(contextNil)
		ctx         = context.WithValue(context.TODO(), logIdKey, logId)
		service     = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		vipInstance *tdmq.DescribeRocketMQVipInstanceDetailResponse
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("cluster_id"); ok {
		paramMap["ClusterId"] = helper.String(v.(string))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTdmqVipInstanceByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		if result == nil {
			return nil
		}

		vipInstance = result
		return nil
	})

	if err != nil {
		return err
	}

	ids := make([]string, 0)
	temp := vipInstance.Response
	if temp.ClusterInfo != nil {
		rocketMQClusterInfoMap := map[string]interface{}{}

		if temp.ClusterInfo.ClusterId != nil {
			rocketMQClusterInfoMap["cluster_id"] = temp.ClusterInfo.ClusterId
		}

		if temp.ClusterInfo.ClusterName != nil {
			rocketMQClusterInfoMap["cluster_name"] = temp.ClusterInfo.ClusterName
		}

		if temp.ClusterInfo.Region != nil {
			rocketMQClusterInfoMap["region"] = temp.ClusterInfo.Region
		}

		if temp.ClusterInfo.CreateTime != nil {
			rocketMQClusterInfoMap["create_time"] = temp.ClusterInfo.CreateTime
		}

		if temp.ClusterInfo.Remark != nil {
			rocketMQClusterInfoMap["remark"] = temp.ClusterInfo.Remark
		}

		if temp.ClusterInfo.PublicEndPoint != nil {
			rocketMQClusterInfoMap["public_end_point"] = temp.ClusterInfo.PublicEndPoint
		}

		if temp.ClusterInfo.VpcEndPoint != nil {
			rocketMQClusterInfoMap["vpc_end_point"] = temp.ClusterInfo.VpcEndPoint
		}

		if temp.ClusterInfo.SupportNamespaceEndpoint != nil {
			rocketMQClusterInfoMap["support_namespace_endpoint"] = temp.ClusterInfo.SupportNamespaceEndpoint
		}

		if temp.ClusterInfo.Vpcs != nil {
			vpcsList := []interface{}{}
			for _, vpcs := range temp.ClusterInfo.Vpcs {
				vpcsMap := map[string]interface{}{}

				if vpcs.VpcId != nil {
					vpcsMap["vpc_id"] = vpcs.VpcId
				}

				if vpcs.SubnetId != nil {
					vpcsMap["subnet_id"] = vpcs.SubnetId
				}

				vpcsList = append(vpcsList, vpcsMap)
			}

			rocketMQClusterInfoMap["vpcs"] = []interface{}{vpcsList}
		}

		if temp.ClusterInfo.IsVip != nil {
			rocketMQClusterInfoMap["is_vip"] = temp.ClusterInfo.IsVip
		}

		if temp.ClusterInfo.RocketMQFlag != nil {
			rocketMQClusterInfoMap["rocket_mq_flag"] = temp.ClusterInfo.RocketMQFlag
		}

		if temp.ClusterInfo.Status != nil {
			rocketMQClusterInfoMap["status"] = temp.ClusterInfo.Status
		}

		if temp.ClusterInfo.IsolateTime != nil {
			rocketMQClusterInfoMap["isolate_time"] = temp.ClusterInfo.IsolateTime
		}

		if temp.ClusterInfo.HttpPublicEndpoint != nil {
			rocketMQClusterInfoMap["http_public_endpoint"] = temp.ClusterInfo.HttpPublicEndpoint
		}

		if temp.ClusterInfo.HttpVpcEndpoint != nil {
			rocketMQClusterInfoMap["http_vpc_endpoint"] = temp.ClusterInfo.HttpVpcEndpoint
		}

		ids = append(ids, *temp.ClusterInfo.ClusterId)
		_ = d.Set("cluster_info", rocketMQClusterInfoMap)
	}

	if temp.InstanceConfig != nil {
		rocketMQInstanceConfigMap := map[string]interface{}{}

		if temp.InstanceConfig.MaxTpsPerNamespace != nil {
			rocketMQInstanceConfigMap["max_tps_per_namespace"] = temp.InstanceConfig.MaxTpsPerNamespace
		}

		if temp.InstanceConfig.MaxNamespaceNum != nil {
			rocketMQInstanceConfigMap["max_namespace_num"] = temp.InstanceConfig.MaxNamespaceNum
		}

		if temp.InstanceConfig.UsedNamespaceNum != nil {
			rocketMQInstanceConfigMap["used_namespace_num"] = temp.InstanceConfig.UsedNamespaceNum
		}

		if temp.InstanceConfig.MaxTopicNum != nil {
			rocketMQInstanceConfigMap["max_topic_num"] = temp.InstanceConfig.MaxTopicNum
		}

		if temp.InstanceConfig.UsedTopicNum != nil {
			rocketMQInstanceConfigMap["used_topic_num"] = temp.InstanceConfig.UsedTopicNum
		}

		if temp.InstanceConfig.MaxGroupNum != nil {
			rocketMQInstanceConfigMap["max_group_num"] = temp.InstanceConfig.MaxGroupNum
		}

		if temp.InstanceConfig.UsedGroupNum != nil {
			rocketMQInstanceConfigMap["used_group_num"] = temp.InstanceConfig.UsedGroupNum
		}

		if temp.InstanceConfig.ConfigDisplay != nil {
			rocketMQInstanceConfigMap["config_display"] = temp.InstanceConfig.ConfigDisplay
		}

		if temp.InstanceConfig.NodeCount != nil {
			rocketMQInstanceConfigMap["node_count"] = temp.InstanceConfig.NodeCount
		}

		if temp.InstanceConfig.NodeDistribution != nil {
			nodeDistributionList := []interface{}{}
			for _, nodeDistribution := range temp.InstanceConfig.NodeDistribution {
				nodeDistributionMap := map[string]interface{}{}

				if nodeDistribution.ZoneName != nil {
					nodeDistributionMap["zone_name"] = nodeDistribution.ZoneName
				}

				if nodeDistribution.ZoneId != nil {
					nodeDistributionMap["zone_id"] = nodeDistribution.ZoneId
				}

				if nodeDistribution.NodeCount != nil {
					nodeDistributionMap["node_count"] = nodeDistribution.NodeCount
				}

				nodeDistributionList = append(nodeDistributionList, nodeDistributionMap)
			}

			rocketMQInstanceConfigMap["node_distribution"] = []interface{}{nodeDistributionList}
		}

		if temp.InstanceConfig.TopicDistribution != nil {
			topicDistributionList := []interface{}{}
			for _, topicDistribution := range temp.InstanceConfig.TopicDistribution {
				topicDistributionMap := map[string]interface{}{}

				if topicDistribution.TopicType != nil {
					topicDistributionMap["topic_type"] = topicDistribution.TopicType
				}

				if topicDistribution.Count != nil {
					topicDistributionMap["count"] = topicDistribution.Count
				}

				topicDistributionList = append(topicDistributionList, topicDistributionMap)
			}

			rocketMQInstanceConfigMap["topic_distribution"] = []interface{}{topicDistributionList}
		}

		if temp.InstanceConfig.MaxQueuesPerTopic != nil {
			rocketMQInstanceConfigMap["max_queues_per_topic"] = temp.InstanceConfig.MaxQueuesPerTopic
		}

		_ = d.Set("instance_config", rocketMQInstanceConfigMap)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), d); e != nil {
			return e
		}
	}

	return nil
}

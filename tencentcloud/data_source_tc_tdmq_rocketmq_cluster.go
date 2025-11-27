/*
Use this data source to query detailed information of tdmqRocketmq cluster

# Example Usage

```hcl

	data "tencentcloudenterprise_tdmq_rocketmq_cluster" "cluster" {
	  name_keyword = "test_rocketmq"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	rocketmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rocketmq_cluster", CNDescription{
		TerraformTypeCN: "RocketMQ集群",
		DescriptionCN:   "提供TDMQ RocketMQ集群数据源，用于查询TDMQ RocketMQ集群的详细信息。",
		AttributesCN: map[string]string{
			"id_keyword":                 "按群集ID搜索",
			"name_keyword":               "按群集名称搜索",
			"cluster_id_list":            "按群集ID筛选",
			"cluster_list":               "群集信息",
			"info":                       "基本群集信息",
			"remark":                     "群集描述（最多128个字符）",
			"cluster_id":                 "群集ID",
			"cluster_name":               "群集名称",
			"region":                     "区域信息",
			"create_time":                "创建时间（以毫秒为单位）",
			"public_end_point":           "公用网络访问地址",
			"vpc_end_point":              "VPC访问地址",
			"support_namespace_endpoint": "是否支持命名空间访问点",
			"vpcs":                       "Vpc列表",
			"vpc_id":                     "Vpc ID",
			"subnet_id":                  "子网ID",
			"is_vip":                     "是否为独占实例",
			"rocketmq_flag":              "Rocketmq集群标识",
			"config":                     "群集配置信息",
			"max_tps_per_namespace":      "每个命名空间的最大TPS",
			"max_namespace_num":          "命名空间的最大数目",
			"used_namespace_num":         "已使用的命名空间数",
			"max_topic_num":              "最大主题数",
			"used_topic_num":             "已使用的主题数",
			"max_group_num":              "最大组数",
			"used_group_num":             "已使用的组数",
			"max_retention_time":         "最大邮件保留期（以毫秒为单位）",
			"max_latency_time":           "最大消息延迟（毫秒）",
			"status":                     "群集状态`0`：正在创建`1`：正常`2`：终止`3`：已删除`4`：孤立的`5`：创建失败`6`：删除失败",
			"result_output_file":         "用于保存结果",
		},
	})
}

func dataSourceTencentCloudTdmqRocketmqCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tdmqRocketmq cluster",
		Read:        dataSourceTencentCloudRocketmqClusterRead,
		Schema: map[string]*schema.Schema{
			"id_keyword": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search by cluster ID.",
			},

			"name_keyword": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search by cluster name.",
			},

			"cluster_id_list": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional:    true,
				Description: "Filter by cluster ID.",
			},

			"cluster_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cluster information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Basic cluster information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"remark": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cluster description (up to 128 characters).",
									},
									"cluster_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cluster ID.",
									},
									"cluster_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cluster name.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Region information.",
									},
									"create_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Creation time in milliseconds.",
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
										Description: "Whether the namespace access point is supported.",
									},
									"vpcs": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Vpc list.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"vpc_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Vpc ID.",
												},
												"subnet_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Subnet ID.",
												},
											},
										},
									},
									"is_vip": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether it is an exclusive instance.",
									},
									"rocketmq_flag": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Rocketmq cluster identification.",
									},
								},
							},
						},
						"config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Cluster configuration information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"max_tps_per_namespace": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum TPS per namespace.",
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
										Description: "Number of used topics.",
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
									"max_retention_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum message retention period in milliseconds.",
									},
									"max_latency_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum message delay in millisecond.",
									},
								},
							},
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Cluster status. `0`: Creating; `1`: Normal; `2`: Terminating; `3`: Deleted; `4`: Isolated; `5`: Creation failed; `6`: Deletion failed.",
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

func dataSourceTencentCloudRocketmqClusterRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_rocketmq_cluster.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("id_keyword"); ok {
		paramMap["id_keyword"] = v.(string)
	}

	if v, ok := d.GetOk("name_keyword"); ok {
		paramMap["name_keyword"] = v.(string)
	}

	if v, ok := d.GetOk("cluster_id_list"); ok {
		clusterIds := v.(*schema.Set).List()
		clusterIdList := make([]string, 0)
		for _, i := range clusterIds {
			clusterId := i.(string)
			clusterIdList = append(clusterIdList, clusterId)
		}
		paramMap["cluster_id_list"] = clusterIdList
	}

	service := TdmqRocketmqService{client: meta.(*TencentCloudClient).apiV3Conn}

	var clusterList []*rocketmq.RocketMQClusterDetail
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := service.DescribeRocketmqClusterByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		clusterList = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Rocketmq clusterList failed, reason:%+v", logId, err)
		return err
	}
	clusterListList := []interface{}{}
	ids := make([]string, 0)

	for _, clusterList := range clusterList {
		clusterListMap := map[string]interface{}{}
		if clusterList.Info != nil {
			infoMap := map[string]interface{}{}
			infoMap["cluster_id"] = clusterList.Info.ClusterId
			ids = append(ids, *clusterList.Info.ClusterId)
			if clusterList.Info.Remark != nil {
				infoMap["remark"] = clusterList.Info.Remark
			}
			if clusterList.Info.ClusterName != nil {
				infoMap["cluster_name"] = clusterList.Info.ClusterName
			}
			if clusterList.Info.Region != nil {
				infoMap["region"] = clusterList.Info.Region
			}
			if clusterList.Info.CreateTime != nil {
				infoMap["create_time"] = clusterList.Info.CreateTime
			}
			if clusterList.Info.PublicEndPoint != nil {
				infoMap["public_end_point"] = clusterList.Info.PublicEndPoint
			}
			if clusterList.Info.VpcEndPoint != nil {
				infoMap["vpc_end_point"] = clusterList.Info.VpcEndPoint
			}
			if clusterList.Info.SupportNamespaceEndpoint != nil {
				infoMap["support_namespace_endpoint"] = clusterList.Info.SupportNamespaceEndpoint
			}
			if clusterList.Info.Vpcs != nil {
				vpcsList := []interface{}{}
				for _, vpcs := range clusterList.Info.Vpcs {
					vpcsMap := map[string]interface{}{}
					if vpcs.VpcId != nil {
						vpcsMap["vpc_id"] = vpcs.VpcId
					}
					if vpcs.SubnetId != nil {
						vpcsMap["subnet_id"] = vpcs.SubnetId
					}

					vpcsList = append(vpcsList, vpcsMap)
				}
				infoMap["vpcs"] = vpcsList
			}
			if clusterList.Info.IsVip != nil {
				infoMap["is_vip"] = clusterList.Info.IsVip
			}
			if clusterList.Info.RocketMQFlag != nil {
				infoMap["rocketmq_flag"] = clusterList.Info.RocketMQFlag
			}

			clusterListMap["info"] = []interface{}{infoMap}
		}
		if clusterList.Config != nil {
			configMap := map[string]interface{}{}
			if clusterList.Config.MaxTpsPerNamespace != nil {
				configMap["max_tps_per_namespace"] = clusterList.Config.MaxTpsPerNamespace
			}
			if clusterList.Config.MaxNamespaceNum != nil {
				configMap["max_namespace_num"] = clusterList.Config.MaxNamespaceNum
			}
			if clusterList.Config.UsedNamespaceNum != nil {
				configMap["used_namespace_num"] = clusterList.Config.UsedNamespaceNum
			}
			if clusterList.Config.MaxTopicNum != nil {
				configMap["max_topic_num"] = clusterList.Config.MaxTopicNum
			}
			if clusterList.Config.UsedTopicNum != nil {
				configMap["used_topic_num"] = clusterList.Config.UsedTopicNum
			}
			if clusterList.Config.MaxGroupNum != nil {
				configMap["max_group_num"] = clusterList.Config.MaxGroupNum
			}
			if clusterList.Config.UsedGroupNum != nil {
				configMap["used_group_num"] = clusterList.Config.UsedGroupNum
			}
			if clusterList.Config.MaxRetentionTime != nil {
				configMap["max_retention_time"] = clusterList.Config.MaxRetentionTime
			}
			if clusterList.Config.MaxLatencyTime != nil {
				configMap["max_latency_time"] = clusterList.Config.MaxLatencyTime
			}

			clusterListMap["config"] = []interface{}{configMap}
		}
		if clusterList.Status != nil {
			clusterListMap["status"] = clusterList.Status
		}

		clusterListList = append(clusterListList, clusterListMap)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("cluster_list", clusterListList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), clusterListList); e != nil {
			return e
		}
	}

	return nil
}

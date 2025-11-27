package tencentcloud

import (
	"context"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_pulsar_clusters", CNDescription{
		TerraformTypeCN: "pulsar集群",
		DescriptionCN:   "提供TDMQ Pulsar集群数据源，用于查询TDMQ Pulsar集群的详细信息。",
		AttributesCN: map[string]string{
			"cluster_ids":                  "群集ID列表",
			"cluster_name":                 "群集名称",
			"result_output_file":           "用于保存结果",
			"clusters":                     "实例信息列表",
			"cluster_id":                   "群集ID",
			"remark":                       "评论",
			"end_point_num":                "端点编号",
			"create_time":                  "创造时间",
			"healthy":                      "健康状态1表示健康，0表示异常",
			"healthy_info":                 "健康信息",
			"status":                       "群集状态0:创建，1:正常，2:删除，3:删除，5:创建失败，6:删除失败",
			"max_namespace_num":            "最大命名空间编号",
			"max_topic_num":                "最大主题编号",
			"max_qps":                      "最大QPS",
			"message_retention_time":       "邮件保留时间",
			"max_storage_capacity":         "最大存储容量",
			"version":                      "群集版本",
			"public_end_point":             "公共端点",
			"vpc_end_point":                "VPC端点",
			"namespace_num":                "命名空间编号",
			"used_storage_budget":          "已用存储预算（MB）",
			"project_id":                   "项目ID",
			"project_name":                 "项目名称",
			"pay_mode":                     "计费模式0：现收现付，1：年/月订阅",
			"max_publish_rate_in_messages": "消息中的最大发布速率",
			"pcluster_name":                "P集群名称",
			"public_access_enabled":        "是否启用公共访问默认值为true",
		},
	})
}

func dataSourceTencentCloudTdmqPulsarCluster() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of tdmq pro_instances",
		Read:        dataSourceTencentCloudTdmqPulsarClusterRead,
		Schema: map[string]*schema.Schema{
			// 集群ID列表
			"cluster_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Cluster ID list.",
			},
			// 集群名称
			"cluster_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster name.",
			},
			// 输出文件
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			// Computed
			// 集群列表
			"clusters": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Instance information list.",
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
							Description: "Cluster name.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remark.",
						},
						"end_point_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Endpoint number.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"healthy": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Healthy status. 1 means healthy, 0 means abnormal.",
						},
						"healthy_info": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Healthy information.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Cluster status. 0: creating, 1: normal, 2: deleting, 3: deleted, 5: creation failed, 6: deletion failed.",
						},
						"max_namespace_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max namespace number.",
						},
						"max_topic_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max topic number.",
						},
						"max_qps": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max QPS.",
						},
						"message_retention_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Message retention time.",
						},
						"max_storage_capacity": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max storage capacity.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster version.",
						},
						"public_end_point": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Public endpoint.",
						},
						"vpc_end_point": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC endpoint.",
						},
						"namespace_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Namespace number.",
						},
						"used_storage_budget": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Used storage budget in MB.",
						},
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Project ID.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Project name.",
						},
						"pay_mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Billing mode. 0: pay-as-you-go, 1: yearly/monthly subscription.",
						},
						"max_publish_rate_in_messages": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max publish rate in messages.",
						},
						"pcluster_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Pcluster name.",
						},
						"public_access_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether public access is enabled. The default value is true.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudTdmqPulsarClusterRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmq_pulsar_clusters.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		ctx      = context.WithValue(context.TODO(), logIdKey, logId)
		request  = tdmq.NewDescribeClustersRequest()
		clusters []*tdmq.Cluster
	)

	if v, ok := d.GetOk("cluster_ids"); ok {
		clusterIds := v.(*schema.Set).List()
		clusterIdList := make([]*string, 0)
		for _, i := range clusterIds {
			clusterId := i.(string)
			clusterIdList = append(clusterIdList, &clusterId)
		}
		request.ClusterIdList = clusterIdList
	}

	if v, ok := d.GetOk("cluster_name"); ok {
		request.ClusterName = helper.String(v.(string))
	}

	err := resource.RetryContext(ctx, readRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseTdmqClient().DescribeClusters(request)
		if e != nil {
			return retryError(e)
		}

		clusters = result.Response.ClusterSet
		return nil
	})

	if err != nil {
		return err
	}

	ids := make([]string, 0, len(clusters))
	tmpList := make([]map[string]interface{}, 0, len(clusters))

	for _, pulsarProInstance := range clusters {
		pulsarProInstanceMap := map[string]interface{}{}

		if pulsarProInstance.ClusterId != nil {
			pulsarProInstanceMap["cluster_id"] = pulsarProInstance.ClusterId
		}

		if pulsarProInstance.ClusterName != nil {
			pulsarProInstanceMap["cluster_name"] = pulsarProInstance.ClusterName
		}

		if pulsarProInstance.Remark != nil {
			pulsarProInstanceMap["remark"] = pulsarProInstance.Remark
		}

		if pulsarProInstance.EndPointNum != nil {
			pulsarProInstanceMap["end_point_num"] = pulsarProInstance.EndPointNum
		}

		if pulsarProInstance.CreateTime != nil {
			pulsarProInstanceMap["create_time"] = pulsarProInstance.CreateTime
		}

		if pulsarProInstance.Healthy != nil {
			pulsarProInstanceMap["healthy"] = pulsarProInstance.Healthy
		}

		if pulsarProInstance.HealthyInfo != nil {
			pulsarProInstanceMap["healthy_info"] = pulsarProInstance.HealthyInfo
		}

		if pulsarProInstance.Status != nil {
			pulsarProInstanceMap["status"] = pulsarProInstance.Status
		}

		if pulsarProInstance.MaxNamespaceNum != nil {
			pulsarProInstanceMap["max_namespace_num"] = pulsarProInstance.MaxNamespaceNum
		}

		if pulsarProInstance.MaxTopicNum != nil {
			pulsarProInstanceMap["max_topic_num"] = pulsarProInstance.MaxTopicNum
		}

		if pulsarProInstance.MaxQps != nil {
			pulsarProInstanceMap["max_qps"] = pulsarProInstance.MaxQps
		}

		if pulsarProInstance.MessageRetentionTime != nil {
			pulsarProInstanceMap["message_retention_time"] = pulsarProInstance.MessageRetentionTime
		}

		if pulsarProInstance.MaxStorageCapacity != nil {
			pulsarProInstanceMap["max_storage_capacity"] = pulsarProInstance.MaxStorageCapacity
		}

		if pulsarProInstance.Version != nil {
			pulsarProInstanceMap["version"] = pulsarProInstance.Version
		}

		if pulsarProInstance.PublicEndPoint != nil {
			pulsarProInstanceMap["public_end_point"] = pulsarProInstance.PublicEndPoint
		}

		if pulsarProInstance.VpcEndPoint != nil {
			pulsarProInstanceMap["vpc_end_point"] = pulsarProInstance.VpcEndPoint
		}

		if pulsarProInstance.NamespaceNum != nil {
			pulsarProInstanceMap["namespace_num"] = pulsarProInstance.NamespaceNum
		}

		if pulsarProInstance.UsedStorageBudget != nil {
			pulsarProInstanceMap["used_storage_budget"] = pulsarProInstance.UsedStorageBudget
		}

		if pulsarProInstance.ProjectId != nil {
			pulsarProInstanceMap["project_id"] = pulsarProInstance.ProjectId
		}

		if pulsarProInstance.ProjectName != nil {
			pulsarProInstanceMap["project_name"] = pulsarProInstance.ProjectName
		}

		if pulsarProInstance.PayMode != nil {
			pulsarProInstanceMap["pay_mode"] = pulsarProInstance.PayMode
		}

		if pulsarProInstance.MaxPublishRateInMessages != nil {
			pulsarProInstanceMap["max_publish_rate_in_messages"] = pulsarProInstance.MaxPublishRateInMessages
		}

		if pulsarProInstance.PclusterName != nil {
			pulsarProInstanceMap["pcluster_name"] = pulsarProInstance.PclusterName
		}

		if pulsarProInstance.PublicAccessEnabled != nil {
			pulsarProInstanceMap["public_access_enabled"] = pulsarProInstance.PublicAccessEnabled
		}

		//if pulsarProInstance.MaxDispatchRateInBytes != nil {
		//	pulsarProInstanceMap["max_dispatch_rate_in_bytes"] = pulsarProInstance.MaxDispatchRateInBytes
		//}

		// For "tags" field, you need to loop through the tags and create a list of maps
		//if pulsarProInstance.Tags != nil {
		//	tags := make([]map[string]interface{}, 0, len(pulsarProInstance.Tags))
		//	for _, tag := range pulsarProInstance.Tags {
		//		tagMap := map[string]interface{}{
		//			"key":   tag.TagKey,
		//			"value": tag.TagValue,
		//		}
		//		tags = append(tags, tagMap)
		//	}
		//	pulsarProInstanceMap["tags"] = tags
		//}

		//if pulsarProInstance.TopicNum != nil {
		//	pulsarProInstanceMap["topic_num"] = pulsarProInstance.TopicNum
		//}

		//if pulsarProInstance.InternalEndPoint != nil {
		//	pulsarProInstanceMap["internal_end_point"] = pulsarProInstance.InternalEndPoint
		//}

		ids = append(ids, *pulsarProInstance.ClusterId)
		tmpList = append(tmpList, pulsarProInstanceMap)
	}

	_ = d.Set("clusters", tmpList)

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}

	return nil
}

package tencentcloud

import (
	"context"

	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rabbitmq_node_list", CNDescription{
		TerraformTypeCN: "TDMQ RabbitMQ节点列表",
		DescriptionCN:   "提供TDMQ RabbitMQ节点列表数据源，用于查询RabbitMQ集群节点的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "RabbitMQ集群ID",
			"node_name":          "模糊搜索节点名称",
			"filters":            "过滤参数名称和值",
			"sort_element":       "按指定元素排序",
			"sort_order":         "排序顺序",
			"result_output_file": "用于保存结果",
			"node_list":          "节点列表",
		},
	})
}

func dataSourceTencentCloudTdmqRabbitmqNodeList() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudTdmqRabbitmqNodeListRead,
		Description: "Use this data source to query detailed information of tdmq rabbitmq_node_list",
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Rabbitmq cluster ID.",
			},
			"node_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Fuzzy search node name.",
			},
			"filters": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Filter parameter name and value. Now there is only one nodeStatus: running/down. Array type, compatible with adding filter parameters later.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the filter parameter.",
						},
						"values": {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Description: "Value.",
						},
					},
				},
			},
			"sort_element": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Sort by the specified element, now there are only 2: cpuUsage/diskUsage.",
			},
			"sort_order": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Sort order: ascend/descend.",
			},
			// computed
			"node_list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Cluster list. Note: This field may return null, indicating that no valid value can be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node nameNote: This field may return null, indicating that no valid value can be obtained.",
						},
						"node_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node statusNote: This field may return null, indicating that no valid value can be obtained.",
						},
						"cpu_usage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CPU usageNote: This field may return null, indicating that no valid value can be obtained.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory usage, in MBNote: This field may return null, indicating that no valid value can be obtained.",
						},
						"disk_usage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Disk usageNote: This field may return null, indicating that no valid value can be obtained.",
						},
						"process_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of Erlang processes for RabbitmqNote: This field may return null, indicating that no valid value can be obtained.",
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

func dataSourceTencentCloudTdmqRabbitmqNodeListRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmq_rabbitmq_node_list.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		service    = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		nodeList   []*tdmq.RabbitMQPrivateNode
		instanceId string
		nodeName   string
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["InstanceId"] = helper.String(v.(string))
		instanceId = v.(string)
	}

	if v, ok := d.GetOk("node_name"); ok {
		paramMap["NodeName"] = helper.String(v.(string))
		nodeName = v.(string)
	}

	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*tdmq.Filter, 0, len(filtersSet))

		for _, item := range filtersSet {
			filter := tdmq.Filter{}
			filterMap := item.(map[string]interface{})

			if v, ok := filterMap["name"]; ok {
				filter.Name = helper.String(v.(string))
			}
			if v, ok := filterMap["values"]; ok {
				valuesSet := v.(*schema.Set).List()
				filter.Values = helper.InterfacesStringsPoint(valuesSet)
			}
			tmpSet = append(tmpSet, &filter)
		}
		paramMap["filters"] = tmpSet
	}

	if v, ok := d.GetOk("sort_element"); ok {
		paramMap["SortElement"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("sort_order"); ok {
		paramMap["SortOrder"] = helper.String(v.(string))
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTdmqRabbitmqNodeListByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		nodeList = result
		return nil
	})

	if err != nil {
		return err
	}

	ids := make([]string, 0)
	tmpList := make([]map[string]interface{}, 0, len(nodeList))
	if nodeList != nil {
		for _, rabbitMQPrivateNode := range nodeList {
			rabbitMQPrivateNodeMap := map[string]interface{}{}

			if rabbitMQPrivateNode.NodeName != nil {
				rabbitMQPrivateNodeMap["node_name"] = rabbitMQPrivateNode.NodeName
			}

			if rabbitMQPrivateNode.NodeStatus != nil {
				rabbitMQPrivateNodeMap["node_status"] = rabbitMQPrivateNode.NodeStatus
			}

			if rabbitMQPrivateNode.CPUUsage != nil {
				rabbitMQPrivateNodeMap["cpu_usage"] = rabbitMQPrivateNode.CPUUsage
			}

			if rabbitMQPrivateNode.Memory != nil {
				rabbitMQPrivateNodeMap["memory"] = rabbitMQPrivateNode.Memory
			}

			if rabbitMQPrivateNode.DiskUsage != nil {
				rabbitMQPrivateNodeMap["disk_usage"] = rabbitMQPrivateNode.DiskUsage
			}

			if rabbitMQPrivateNode.ProcessNumber != nil {
				rabbitMQPrivateNodeMap["process_number"] = rabbitMQPrivateNode.ProcessNumber
			}

			tmpList = append(tmpList, rabbitMQPrivateNodeMap)
		}

		_ = d.Set("node_list", tmpList)
	}

	ids = append(ids, instanceId)
	ids = append(ids, nodeName)
	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}

	return nil
}

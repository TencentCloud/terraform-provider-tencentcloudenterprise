/*
Use this data source to query detailed information of tdmq rabbitmq node list

Example Usage

### Query all nodes in a RabbitMQ cluster

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_node_list" "all_nodes" {
  instance_id = "amqp-xxxxxxxx"
}

output "node_list" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_node_list.all_nodes.node_list
}
```

### Query RabbitMQ nodes by exact node name

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_node_list" "specific_node" {
  instance_id = "amqp-xxxxxxxx"
  node_name   = "rabbit@rabbitmq-broker-1.rabbitmq-broker-internal.amqp-7d39mjvn.svc.cluster.local"
}

output "node_status" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_node_list.specific_node.node_list[0].node_status
}
```

*/
package tencentcloud

import (
	"context"

	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rabbitmq_node_list", CNDescription{
		TerraformTypeCN: "TDMQ RabbitMQ节点列表",
		DescriptionCN:   "提供TDMQ RabbitMQ节点列表数据源，用于查询RabbitMQ集群节点的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "RabbitMQ集群实例ID",
			"node_name":          "节点名称，需要使用精确名称查询",
			"filters":            "过滤条件,如nodeStatus",
			"name":               "过滤器参数名称",
			"values":             "过滤器参数值",
			"sort_element":       "排序字段,如diskUsage",
			"sort_order":         "排序顺序,升序(ascend)或降序(descend)",
			"node_list":          "RabbitMQ节点信息列表",
			"node_status":        "节点运行状态,running运行中/down已停止",
			//"cpu_usage":          "节点CPU使用率百分比",
			"memory":             "节点内存使用量,单位MB",
			"disk_usage":         "节点磁盘使用率百分比",
			"process_number":     "节点上运行的Erlang进程数量",
			"result_output_file": "结果输出文件路径，前端不支持该参数的使用",
		},
	})
}

func dataSourceTencentCloudTdmqRabbitmqNodeList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudTdmqRabbitmqNodeListRead,
		Description: "Use this data source to query detailed information of tdmq rabbitmq_node_list",
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "RabbitMQ cluster instance ID. The ID of the RabbitMQ instance to query node information from.",
			},
			"node_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Node name for fuzzy search filtering. Returns nodes whose names contain this string.",
			},
			"filters": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Filter conditions for querying nodes. Currently supports filtering by `nodeStatus` with values: `running` (node is running), `down` (node is down). Multiple filters can be specified. Array type design allows for future filter parameter extensions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the filter parameter. Currently supported: `nodeStatus` (filter by node running status).",
						},
						"values": {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Description: "Value(s) for the filter parameter. For `nodeStatus`, valid values are: `running`, `down`. Multiple values can be specified.",
						},
					},
				},
			},
			"sort_element": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Field to sort results by. Currently supported value: `diskUsage` (sort by disk usage).",
			},
			"sort_order": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Sort order. Valid values: `ascend` (ascending order), `descend` (descending order).",
			},
			// computed
			"node_list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "List of RabbitMQ cluster nodes matching the filter criteria. Note: This field may return null if no nodes match the query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"node_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the RabbitMQ node. Note: This field may return null.",
						},
						"node_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node running status. Possible values: `running` (node is operational), `down` (node is not responding). Note: This field may return null.",
						},
					//"cpu_usage": {
					//	Type:        schema.TypeString,
					//	Computed:    true,
					//	Description: "CPU usage percentage of the node. Indicates how much CPU resources the node is consuming. Note: This field may return null.",
					//},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory usage in MB (megabytes). The amount of memory currently being used by the node. Note: This field may return null.",
						},
						"disk_usage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Disk usage percentage of the node. Indicates how much disk storage the node is consuming. Note: This field may return null.",
						},
						"process_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of Erlang processes running on the RabbitMQ node. RabbitMQ is built on Erlang, and this metric indicates the total number of Erlang processes active on the node. Note: This field may return null.",
						},
					},
				},
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File path to save the query results. If specified, the node list will be written to this file in JSON format.",
			},
		},
	}
}

func dataSourceTencentCloudTdmqRabbitmqNodeListRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tdmq_rabbitmq_node_list.read")()
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

			//if rabbitMQPrivateNode.CPUUsage != nil {
			//	rabbitMQPrivateNodeMap["cpu_usage"] = rabbitMQPrivateNode.CPUUsage
			//}

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

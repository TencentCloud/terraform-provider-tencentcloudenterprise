/*
Use this data source to query detailed information of tdmq rabbitmq_vip_instance

Example Usage

### Query all RabbitMQ VIP instances

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "all_instances" {
}

output "instance_list" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.all_instances.instances
}
```

### Query RabbitMQ VIP instance by instance ID

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "by_instance_id" {
  filters {
    name   = "instanceIds"
    values = ["amqp-xxxxxxxx"]
  }
}

output "instance_info" {
  value = data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.by_instance_id.instances[0]
}
```

### Query RabbitMQ VIP instance by instance name

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "by_name" {
  filters {
    name   = "instanceName"
    values = ["my-rabbitmq-cluster"]
  }
}
```

### Query RabbitMQ VIP instances via endpoint

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "by_endpoint" {
  filters {
    name   = "endpoint"
    values = ["10.0.0.17"]
  }
}
```

### Query and export to file

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "export_instances" {
  result_output_file = "./rabbitmq_instances.json"
}
```

### Access instance details including VPC and zone information

```hcl
data "tencentcloudenterprise_tdmq_rabbitmq_vip_instance" "detailed_info" {
  filters {
    name   = "instanceIds"
    values = ["rmq-xxxxxxxx"]
  }
}

output "instance_zone_ids" {
  description = "Availability zones where the instance is deployed"
  value       = data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.detailed_info.instances[0].zone_ids
}

output "instance_vpcs" {
  description = "VPC access endpoints"
  value       = data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.detailed_info.instances[0].vpcs
}

output "vpc_endpoint" {
  description = "First VPC endpoint address for connection"
  value       = length(data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.detailed_info.instances[0].vpcs) > 0 ? data.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.detailed_info.instances[0].vpcs[0].vpc_endpoint : ""
}
```

 */
package tencentcloud

import (
	"context"
	"fmt"
	"log"

	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rabbitmq_vip_instance", CNDescription{
		TerraformTypeCN: "TDMQ RabbitMQ专享版实例列表",
		DescriptionCN:   "提供TDMQ RabbitMQ专享版实例列表数据源，用于查询RabbitMQ专享版实例的详细信息。",
		AttributesCN: map[string]string{
			"filters":                         "查询条件过滤器,用于筛选实例",
			"name":                            "过滤器参数名称",
			"values":                          "过滤器参数值",
			"instances":                       "RabbitMQ实例信息列表",
			"instance_id":                     "实例ID,实例的唯一标识符",
			"instance_name":                   "实例名称,用户定义的实例名称",
			"instance_version":                "实例版本号,如3.8.30或3.11.8",
			"status":                          "实例状态,0创建中/1正常/2隔离中/3已销毁/4异常/5发货失败",
			"node_count":                      "集群节点数量",
			"config_display":                  "实例配置规格显示名称,如4核8GB",
			"max_tps":                         "峰值TPS(每秒事务数),实例最大吞吐能力",
			"max_band_width":                  "峰值带宽,单位Mbps",
			"max_storage":                     "总存储容量,单位GB",
			"expire_time":                     "实例到期时间,毫秒时间戳",
			"auto_renew_flag":                 "自动续费标志,0默认手动续费/1自动续费/2明确不自动续费",
			"pay_mode":                        "付费模式,0后付费/1预付费",
			"remark":                          "实例备注信息",
			"spec_name":                       "实例规格名称,如rabbit-vip-basic-1",
			"exception_information":           "实例异常信息,正常时为空",
			"zone_ids":                        "实例部署的可用区ID列表",
			"vpcs":                            "VPC接入点列表",
			"vpc_id":                          "VPC ID",
			"subnet_id":                       "子网ID",
			"vpc_endpoint":                    "VPC访问端点地址",
			"vpc_data_stream_endpoint_status": "VPC端点状态,OFF/ON/CREATING/DELETING",
			"result_output_file":              "结果输出文件路径，前端不支持该参数的使用",
		},
	})
}

func dataSourceTencentCloudTdmqRabbitmqVipInstance() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudTdmqRabbitmqVipInstanceRead,
		Description: "Use this data source to query detailed information of tdmq rabbitmq vip instance",
		Schema: map[string]*schema.Schema{
			"filters": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Filter conditions for querying RabbitMQ instances. Multiple filters can be specified to narrow down the query results.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the filter parameter. Common filter names include: `instanceIds` (filter by instance IDs), `instanceName` (filter by instance name), `status` (filter by instance status).",
						},
						"values": {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Optional:    true,
							Description: "Value(s) for the filter parameter. Multiple values can be specified, and instances matching any of the values will be returned.",
						},
					},
				},
			},
			// computed
			"instances": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "List of RabbitMQ VIP instances matching the filter criteria.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance ID. Unique identifier for the RabbitMQ instance.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance name. The user-defined name of the RabbitMQ instance.",
						},
						"instance_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "RabbitMQ version of the instance (e.g., 3.8.30, 3.11.8). Note: This field may return null.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance status. Valid values: `0` (creating), `1` (running/normal), `2` (isolating), `3` (destroyed), `4` (abnormal/failed), `5` (delivery failed).",
						},
						"node_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of nodes in the cluster.",
						},
						"config_display": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Display name of the instance configuration specification (e.g., \"4 cores 8GB\").",
						},
						"max_tps": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Peak TPS (Transactions Per Second). Maximum throughput capacity of the instance.",
						},
						"max_band_width": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Peak bandwidth in Mbps (megabits per second). Maximum network bandwidth capacity of the instance.",
						},
						"max_storage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total storage capacity in GB (gigabytes). Maximum disk space available for the instance.",
						},
						"expire_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance expiration timestamp in milliseconds. The time when the prepaid instance will expire.",
						},
						"auto_renew_flag": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Auto-renewal flag. Valid values: `0` (default/manual renewal - user has not set auto-renewal), `1` (auto-renewal enabled), `2` (auto-renewal explicitly disabled by user).",
						},
						"pay_mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Billing mode. Valid values: `0` (postpaid/pay-as-you-go), `1` (prepaid/monthly subscription).",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remark or note about the instance. Note: This field may return null.",
						},
						"spec_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance specification name (e.g., rabbit-vip-basic-1). The internal identifier for the instance configuration.",
						},
						"exception_information": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Exception or error information if the cluster is in abnormal state. Note: This field may return null when the instance is running normally.",
						},
						"zone_ids": {
							Type:        schema.TypeSet,
							Elem:        &schema.Schema{Type: schema.TypeInt},
							Computed:    true,
							Description: "Availability zone ID list where the instance is deployed. Multiple zone IDs indicate multi-availability zone deployment.",
						},
						"vpcs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of VPC access points. Contains VPC network endpoint information for accessing the RabbitMQ instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "VPC ID where the access endpoint is located.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Subnet ID where the access endpoint is located.",
									},
									"vpc_endpoint": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "VPC private network access endpoint address. Use this address to connect to RabbitMQ from within the VPC.",
									},
									"vpc_data_stream_endpoint_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of the VPC endpoint. Valid values: `OFF`, `ON`, `CREATING`, `DELETING`.",
									},
								},
							},
						},
					},
				},
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File path to save the query results. If specified, the instance list will be written to this file in JSON format.",
			},
		},
	}
}

func dataSourceTencentCloudTdmqRabbitmqVipInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_tdmq_rabbitmq_vip_instance.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId     = getLogId(contextNil)
		ctx       = context.WithValue(context.TODO(), logIdKey, logId)
		service   = TdmqService{client: meta.(*TencentCloudClient).apiV3Conn}
		instances []*tdmq.RabbitMQVipInstance
	)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*tdmq.Filter, 0, len(filtersSet))
		log.Printf(fmt.Sprintf("[DEBUG] %s: %#v", "filters1", filtersSet))
		for _, item := range filtersSet {
			log.Printf(fmt.Sprintf("[DEBUG] %s: %#v", "filters", item))
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

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeTdmqRabbitmqVipInstanceByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}

		instances = result
		return nil
	})

	if err != nil {
		return err
	}

	ids := make([]string, 0, len(instances))
	tmpList := make([]map[string]interface{}, 0, len(instances))

	if instances != nil {
		for _, rabbitMQVipInstance := range instances {
			rabbitMQVipInstanceMap := map[string]interface{}{}

			if rabbitMQVipInstance.InstanceId != nil {
				rabbitMQVipInstanceMap["instance_id"] = rabbitMQVipInstance.InstanceId
			}

			if rabbitMQVipInstance.InstanceName != nil {
				rabbitMQVipInstanceMap["instance_name"] = rabbitMQVipInstance.InstanceName
			}

			if rabbitMQVipInstance.InstanceVersion != nil {
				rabbitMQVipInstanceMap["instance_version"] = rabbitMQVipInstance.InstanceVersion
			}

			if rabbitMQVipInstance.Status != nil {
				rabbitMQVipInstanceMap["status"] = rabbitMQVipInstance.Status
			}

			if rabbitMQVipInstance.NodeCount != nil {
				rabbitMQVipInstanceMap["node_count"] = rabbitMQVipInstance.NodeCount
			}

			if rabbitMQVipInstance.ConfigDisplay != nil {
				rabbitMQVipInstanceMap["config_display"] = rabbitMQVipInstance.ConfigDisplay
			}

			if rabbitMQVipInstance.MaxTps != nil {
				rabbitMQVipInstanceMap["max_tps"] = rabbitMQVipInstance.MaxTps
			}

			if rabbitMQVipInstance.MaxBandWidth != nil {
				rabbitMQVipInstanceMap["max_band_width"] = rabbitMQVipInstance.MaxBandWidth
			}

			if rabbitMQVipInstance.MaxStorage != nil {
				rabbitMQVipInstanceMap["max_storage"] = rabbitMQVipInstance.MaxStorage
			}

			if rabbitMQVipInstance.ExpireTime != nil {
				rabbitMQVipInstanceMap["expire_time"] = rabbitMQVipInstance.ExpireTime
			}

			if rabbitMQVipInstance.AutoRenewFlag != nil {
				rabbitMQVipInstanceMap["auto_renew_flag"] = rabbitMQVipInstance.AutoRenewFlag
			}

			if rabbitMQVipInstance.PayMode != nil {
				rabbitMQVipInstanceMap["pay_mode"] = rabbitMQVipInstance.PayMode
			}

			if rabbitMQVipInstance.Remark != nil {
				rabbitMQVipInstanceMap["remark"] = rabbitMQVipInstance.Remark
			}

			if rabbitMQVipInstance.SpecName != nil {
				rabbitMQVipInstanceMap["spec_name"] = rabbitMQVipInstance.SpecName
			}

			//if rabbitMQVipInstance.ExceptionInformation != nil {
			//	rabbitMQVipInstanceMap["exception_information"] = rabbitMQVipInstance.ExceptionInformation
			//}

			if rabbitMQVipInstance.ZoneIds != nil {
				rabbitMQVipInstanceMap["zone_ids"] = rabbitMQVipInstance.ZoneIds
			}

			if rabbitMQVipInstance.Vpcs != nil {
				vpcsList := make([]map[string]interface{}, 0, len(rabbitMQVipInstance.Vpcs))
				for _, vpc := range rabbitMQVipInstance.Vpcs {
					vpcMap := map[string]interface{}{}
					if vpc.VpcId != nil {
						vpcMap["vpc_id"] = vpc.VpcId
					}
					if vpc.SubnetId != nil {
						vpcMap["subnet_id"] = vpc.SubnetId
					}
					if vpc.VpcEndpoint != nil {
						vpcMap["vpc_endpoint"] = vpc.VpcEndpoint
					}
					if vpc.VpcDataStreamEndpointStatus != nil {
						vpcMap["vpc_data_stream_endpoint_status"] = vpc.VpcDataStreamEndpointStatus
					}
					vpcsList = append(vpcsList, vpcMap)
				}
				rabbitMQVipInstanceMap["vpcs"] = vpcsList
			}

			ids = append(ids, *rabbitMQVipInstance.InstanceId)
			tmpList = append(tmpList, rabbitMQVipInstanceMap)
		}

		_ = d.Set("instances", tmpList)
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

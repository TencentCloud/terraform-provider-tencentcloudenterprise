package tencentcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tdmq "terraform-provider-tencentcloudenterprise/sdk/tdmq/v20200217"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_tdmq_rabbitmq_vip_instance", CNDescription{
		TerraformTypeCN: "TDMQ RabbitMQ专享版实例列表",
		DescriptionCN:   "提供TDMQ RabbitMQ专享版实例列表数据源，用于查询RabbitMQ专享版实例的详细信息。",
		AttributesCN: map[string]string{
			"filters":   "查询条件过滤器",
			"instances": "实例信息列表",
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
				Description: "Query condition filter.",
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
			// computed
			"instances": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Instance information list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance id.",
						},
						"instance_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance name.",
						},
						"instance_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance versionNote: This field may return null, indicating that no valid value can be obtained.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance status, 0 means creating, 1 means normal, 2 means isolating, 3 means destroyed, 4 - abnormal, 5 - delivery failed.",
						},
						"node_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of nodes.",
						},
						"config_display": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance configuration specification name.",
						},
						"max_tps": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Peak TPS.",
						},
						"max_band_width": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Peak bandwidth, in Mbps.",
						},
						"max_storage": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Storage capacity, in GB.",
						},
						"expire_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Instance expiration time, in milliseconds.",
						},
						"auto_renew_flag": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Automatic renewal mark, 0 indicates the default state (the user has not set it, that is, the initial state is manual renewal), 1 indicates automatic renewal, 2 indicates that the automatic renewal is not specified (user setting).",
						},
						"pay_mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "0-postpaid, 1-prepaid.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "RemarksNote: This field may return null, indicating that no valid value can be obtained.",
						},
						"spec_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance Configuration ID.",
						},
						"exception_information": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cluster is abnormal.Note: This field may return null, indicating that no valid value can be obtained.",
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

func dataSourceTencentCloudTdmqRabbitmqVipInstanceRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_tdmq_rabbitmq_vip_instance.read")()
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

/*
Use this data source to query cvm instances modification.

# Example Usage

```hcl

	data "tencentcloudenterprise_cvm_instances_modification" "foo" {
	  instance_ids = ["ins-xxxxxxx"]
	}

```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cvm_instances_modification", CNDescription{
		TerraformTypeCN: "可调整配置机型的信息",
		DescriptionCN:   "提供CVM实例可调整配置机型信息数据源，用于查询实例可调整的配置信息。",
		AttributesCN: map[string]string{
			"instance_ids":                     "实例ID列表",
			"filters":                          "过滤器",
			"name":                             "过滤字段",
			"values":                           "过滤值",
			"result_output_file":               "用于保存结果的文件",
			"status":                           "状态描述",
			"message":                          "状态描述信息",
			"zone":                             "可用区",
			"instance_type":                    "机型",
			"instance_family":                  "机型族",
			"gpu":                              "GPU核数",
			"cpu":                              "CPU核数",
			"memory":                           "内存容量（GB）",
			"fpga":                             "FPGA核数",
			"instance_type_config_status_list": "可调整配置机型的信息列表",
		},
	})
}

func dataSourceTencentCloudCvmInstancesModification() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query cvm instances modification.",
		Read:        dataSourceTencentCloudCvmInstancesModificationRead,

		Schema: map[string]*schema.Schema{
			"instance_ids": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "One or more instance ID to be queried. It can be obtained from the InstanceId in the returned value of API DescribeInstances. The maximum number of instances in batch for each request is 20.",
			},
			"filters": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "The upper limit of Filters for each request is 10 and the upper limit for Filter.Values is 2.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Fields to be filtered.",
						},
						"values": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "Value of the field.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			"instance_type_config_status_list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "The list of model configurations that can be adjusted by the instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State description.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status description information.",
						},
						"instance_type_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Availability zone.",
									},
									"instance_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Instance type.",
									},
									"instance_family": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Instance family.",
									},
									"gpu": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of GPU kernels, in cores.",
									},
									"cpu": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of CPU kernels, in cores.",
									},
									"memory": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Memory capacity (in GB).",
									},
									"fpga": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of FPGA kernels, in cores.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCvmInstancesModificationRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cvm_instances_modification.read")()
	logId := getLogId(contextNil)

	var (
		request  = cvm.NewDescribeInstancesModificationRequest()
		response = cvm.NewDescribeInstancesModificationResponse()
	)
	if v, ok := d.GetOk("instance_ids"); ok {
		request.InstanceIds = helper.InterfacesStringsPoint(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("filters"); ok {
		filters := make([]*cvm.Filter, 0)
		for _, item := range v.(*schema.Set).List() {
			filter := item.(map[string]interface{})
			name := filter["name"].(string)
			filters = append(filters, &cvm.Filter{
				Name:   &name,
				Values: helper.StringsStringsPoint(filter["values"].([]string)),
			})
		}
		request.Filters = filters
	}

	instanceTypeConfigStatusList := make([]map[string]interface{}, 0)

	var innerErr error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, innerErr = meta.(*TencentCloudClient).apiV3Conn.UseCvmClient().DescribeInstancesModification(request)
		if innerErr != nil {
			return retryError(innerErr)
		}
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0)
	for _, instanceTypeConfigStatusSetItem := range response.Response.InstanceTypeConfigStatusSet {
		instanceTypeConfigStatus := make(map[string]interface{})
		instanceTypeConfigStatus["status"] = instanceTypeConfigStatusSetItem.Status
		instanceTypeConfigStatus["message"] = instanceTypeConfigStatusSetItem.Message

		instanceTypeConfigMaps := make([]map[string]interface{}, 0)
		instanceTypeConfigMap := make(map[string]interface{})
		instanceTypeConfig := instanceTypeConfigStatusSetItem.InstanceTypeConfig
		instanceTypeConfigMap["zone"] = instanceTypeConfig.Zone
		ids = append(ids, *instanceTypeConfig.InstanceType)
		instanceTypeConfigMap["instance_type"] = instanceTypeConfig.InstanceType
		instanceTypeConfigMap["instance_family"] = instanceTypeConfig.InstanceFamily
		instanceTypeConfigMap["gpu"] = instanceTypeConfig.GPU
		instanceTypeConfigMap["cpu"] = instanceTypeConfig.CPU
		instanceTypeConfigMap["memory"] = instanceTypeConfig.Memory
		instanceTypeConfigMap["fpga"] = instanceTypeConfig.FPGA
		instanceTypeConfigMaps = append(instanceTypeConfigMaps, instanceTypeConfigMap)
		instanceTypeConfigStatus["instance_type_config"] = instanceTypeConfigMaps

		instanceTypeConfigStatusList = append(instanceTypeConfigStatusList, instanceTypeConfigStatus)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("instance_type_config_status_list", instanceTypeConfigStatusList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set instance list fail, reason:%s\n ", logId, err.Error())
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), instanceTypeConfigStatusList); err != nil {
			return err
		}
	}
	return nil

}

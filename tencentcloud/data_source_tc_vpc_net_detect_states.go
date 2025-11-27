/*
Use this data source to query detailed information of vpc net_detect_states

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_net_detect_states" "net_detect_states" {
	  net_detect_ids = ["netd-12345678"]
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_net_detect_states", CNDescription{
		TerraformTypeCN: "vpc网络检测实例",
		DescriptionCN:   "提供VPC网络探测状态数据源，用于查询VPC网络探测状态的详细信息。",
		AttributesCN: map[string]string{
			"net_detect_ids": "网络检测实例“ID”的数组，例如“netd-12345678”",
			"filters":        "过滤条件`不能同时指定NetDetectIds`和`Filters'net-detect-id-String-（筛选器条件）网络检测实例id，如netd-12345678",
			"name":           "属性名称如果存在多个筛选器，则这些筛选器之间的逻辑关系为“与”",
			"属性值如果一个筛选器中存在多个值，则这些值之间的逻辑关系为“OR”对于“bool”参数，有效值包括“TRUE”和“FALSE”s": "属性值如果一个筛选器中存在多个值，则这些值之间的逻辑关系为“OR”对于“bool”参数，有效值包括“TRUE”和“FALSE”",
			"net_detect_state_set":    "阵列网络检测验证结果符合要求注意：此字段可能返回null，表示无法获得有效值",
			"net_detect_id":           "网络检测实例的ID，例如netd-12345678",
			"net_detect_ip_state_set": "网络检测目标IP验证结果的数组",
			"detect_destination_ip":   "网络检测的目标IPv4地址",
			"state":                   "检测结果0:成功-1：路由过程中未发生丢包-2：出站流量被ACL阻塞时发生丢包-3：入站流量被ACL阻塞时发生丢包-4：其他错误",
			"delay":                   "延迟单位：ms",
			"packet_loss_rate":        "数据包丢失率",
			"result_output_file":      "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcNetDetectStates() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc net_detect_states",
		Read:        dataSourceTencentCloudVpcNetDetectStatesRead,
		Schema: map[string]*schema.Schema{
			"net_detect_ids": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The array of network detection instance `IDs`, such as [`netd-12345678`].",
			},

			"filters": {
				Optional:    true,
				Type:        schema.TypeList,
				Description: "Filter conditions. `NetDetectIds` and `Filters` cannot be specified at the same time.net-detect-id - String - (Filter condition) The network detection instance ID, such as netd-12345678.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The attribute name. If more than one Filter exists, the logical relation between these Filters is `AND`.",
						},
						"values": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Required:    true,
							Description: "Attribute value. If multiple values exist in one filter, the logical relationship between these values is `OR`. For a `bool` parameter, the valid values include `TRUE` and `FALSE`.",
						},
					},
				},
			},

			"net_detect_state_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "The array of network detection verification results that meet requirements.Note: This field may return null, indicating that no valid values can be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"net_detect_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of a network detection instance, such as netd-12345678.",
						},
						"net_detect_ip_state_set": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The array of network detection destination IP verification results.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"detect_destination_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The destination IPv4 address of network detection.",
									},
									"state": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The detection result.0: successful;-1: no packet loss occurred during routing;-2: packet loss occurred when outbound traffic is blocked by the ACL;-3: packet loss occurred when inbound traffic is blocked by the ACL;-4: other errors.",
									},
									"delay": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The latency. Unit: ms.",
									},
									"packet_loss_rate": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The packet loss rate.",
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
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudVpcNetDetectStatesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_net_detect_states.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("net_detect_ids"); ok {
		netDetectIdsSet := v.(*schema.Set).List()
		paramMap["NetDetectIds"] = helper.InterfacesStringsPoint(netDetectIdsSet)
	}

	if v, ok := d.GetOk("filters"); ok {
		filtersSet := v.([]interface{})
		tmpSet := make([]*vpc.Filter, 0, len(filtersSet))

		for _, item := range filtersSet {
			filter := vpc.Filter{}
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
		paramMap["Filters"] = tmpSet
	}

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var netDetectStateSet []*vpc.NetDetectState

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcNetDetectStatesByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		netDetectStateSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(netDetectStateSet))
	tmpList := make([]map[string]interface{}, 0, len(netDetectStateSet))

	if netDetectStateSet != nil {
		for _, netDetectState := range netDetectStateSet {
			netDetectStateMap := map[string]interface{}{}

			if netDetectState.NetDetectId != nil {
				netDetectStateMap["net_detect_id"] = netDetectState.NetDetectId
			}

			if netDetectState.NetDetectIpStateSet != nil {
				netDetectIpStateSetList := []interface{}{}
				for _, netDetectIpStateSet := range netDetectState.NetDetectIpStateSet {
					netDetectIpStateSetMap := map[string]interface{}{}

					if netDetectIpStateSet.DetectDestinationIp != nil {
						netDetectIpStateSetMap["detect_destination_ip"] = netDetectIpStateSet.DetectDestinationIp
					}

					if netDetectIpStateSet.State != nil {
						netDetectIpStateSetMap["state"] = netDetectIpStateSet.State
					}

					if netDetectIpStateSet.Delay != nil {
						netDetectIpStateSetMap["delay"] = netDetectIpStateSet.Delay
					}

					if netDetectIpStateSet.PacketLossRate != nil {
						netDetectIpStateSetMap["packet_loss_rate"] = netDetectIpStateSet.PacketLossRate
					}

					netDetectIpStateSetList = append(netDetectIpStateSetList, netDetectIpStateSetMap)
				}

				netDetectStateMap["net_detect_ip_state_set"] = netDetectIpStateSetList
			}

			ids = append(ids, *netDetectState.NetDetectId)
			tmpList = append(tmpList, netDetectStateMap)
		}

		_ = d.Set("net_detect_state_set", tmpList)
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

/*
Use this data source to query detailed information of vpc gateway_flow_qos

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_gateway_flow_qos" "gateway_flow_qos" {
	  gateway_id = "vpngw-gt8bianl"
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_gateway_flow_qos", CNDescription{
		TerraformTypeCN: "vpc网关流量限制",
		DescriptionCN:   "提供VPC网关流量QoS数据源，用于查询VPC网关流量QoS的详细信息。",
		AttributesCN: map[string]string{
			"gateway_id":         "网络实例ID，我们目前支持的网络实例类型有：专线网关实例ID，格式为`dcg-ltjahce6`；Nat网关实例ID，格式为“Nat-ltjahce6”；VPN网关实例ID，格式为“VPN-ltjahce6”",
			"ip_addresses":       "具有流量限制的云服务器的Intranet IP",
			"gateway_qos_set":    "实例详细信息列表",
			"vpc_id":             "vpc id",
			"ip_address":         "cvm ip地址",
			"bandwidth":          "带宽值",
			"create_time":        "创造时间",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcGatewayFlowQos() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc gateway_flow_qos",
		Read:        dataSourceTencentCloudVpcGatewayFlowQosRead,
		Schema: map[string]*schema.Schema{
			"gateway_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Network instance ID, the network instance types we currently support are:Private line gateway instance ID, in the form of `dcg-ltjahce6`;Nat gateway instance ID, in the form of `nat-ltjahce6`;VPN gateway instance ID, in the form of `vpn-ltjahce6`.",
			},

			"ip_addresses": {
				Optional: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Intranet IP of the cloud server with traffic limitation.",
			},

			"gateway_qos_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Instance detail list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Vpc id.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cvm ip address.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "bandwidth value.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
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

func dataSourceTencentCloudVpcGatewayFlowQosRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_gateway_flow_qos.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("gateway_id"); ok {
		paramMap["GatewayId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("ip_addresses"); ok {
		ipAddressesSet := v.(*schema.Set).List()
		paramMap["IpAddresses"] = helper.InterfacesStringsPoint(ipAddressesSet)
	}

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var gatewayQosSet []*vpc.GatewayQos

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcGatewayFlowQosByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		gatewayQosSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(gatewayQosSet))
	tmpList := make([]map[string]interface{}, 0, len(gatewayQosSet))

	if gatewayQosSet != nil {
		for _, gatewayQos := range gatewayQosSet {
			gatewayQosMap := map[string]interface{}{}

			if gatewayQos.VpcId != nil {
				gatewayQosMap["vpc_id"] = gatewayQos.VpcId
			}

			if gatewayQos.IpAddress != nil {
				gatewayQosMap["ip_address"] = gatewayQos.IpAddress
			}

			if gatewayQos.Bandwidth != nil {
				gatewayQosMap["bandwidth"] = gatewayQos.Bandwidth
			}

			if gatewayQos.CreateTime != nil {
				gatewayQosMap["create_time"] = gatewayQos.CreateTime
			}

			ids = append(ids, *gatewayQos.IpAddress)
			tmpList = append(tmpList, gatewayQosMap)
		}

		_ = d.Set("gateway_qos_set", tmpList)
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

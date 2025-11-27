/*
Use this data source to query detailed information of vpc net_detect_state_check

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_net_detect_state_check" "net_detect_state_check" {
	  net_detect_id         = "netd-12345678"
	  detect_destination_ip = [
	    "203.0.113.3",
	    "203.0.113.2"
	  ]
	  next_hop_type        = "NORMAL_CVM"
	  next_hop_destination = "203.0.113.4"
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_net_detect_state_check", CNDescription{
		TerraformTypeCN: "VPC网络探测状态检查",
		DescriptionCN:   "提供VPC网络探测状态检查数据源，用于查询VPC网络探测状态检查的详细信息。",
		AttributesCN: map[string]string{
			"detect_destination_ip":   "检测目标IPv4地址的阵列，最多包含两个IP地址",
			"next_hop_type":           "下一个跃点的类型目前支持的类型有：VPN:VPN网关；DIRECTCONNECT：专线网关；PEERCONNECTION：对等连接；NAT：NAT网关；NORMAL_CVM：正常CVM",
			"next_hop_destination":    "下一个跃点目标网关该值与NextHopType有关如果NextHopType设置为VPN，则该参数的值为VPN网关ID，如vpngw-1235678如果NextThopType设置为DIRECTCONNECT，则该值为直接连接网关ID，例如dcg-12345678如果NextMopType设置为PEERCONNECTION，则该参数值为对等连接ID，例如pcx-12345678如果NextHopType设置为NAT，则此参数的值为NAT网关ID，例如NAT-12345678",
			"net_detect_id":           "网络检查器实例的ID，例如netd-12345678请至少输入以下参数之一：VpcId、SubnetId和NetDetectName如果存在NetDetectId，请使用它",
			"vpc_id":                  "“VPC”实例的ID，例如“VPC-12345678”，它与SubnetId和NetDetectName一起使用您应该输入此参数或NetDetectId，或者同时输入两者如果存在NetDetectId，请使用它",
			"subnet_id":               "子网实例的ID，例如“子网-12345678”，它与VpcId和NetDetectName一起使用您应该输入此参数或NetDetectId，或者同时输入两者如果存在NetDetectId，请使用它",
			"net_detect_name":         "网络检查器的名称，长度最多为60字节它与VpcId和NetDetectName一起使用您应该输入此参数或NetDetectId，或者同时输入两者如果存在NetDetectId，请使用它",
			"net_detect_ip_state_set": "阵列网络检测验证结果",
			"state":                   "检测结果0:成功-1：路由过程中未发生丢包-2：出站流量被ACL阻塞时发生丢包-3：入站流量被ACL阻塞时发生丢包-4：其他错误",
			"delay":                   "延迟单位：ms",
			"packet_loss_rate":        "数据包丢失率",
			"result_output_file":      "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcNetDetectStateCheck() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc net_detect_state_check",
		Read:        dataSourceTencentCloudVpcNetDetectStateCheckRead,
		Schema: map[string]*schema.Schema{
			"detect_destination_ip": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The array of detection destination IPv4 addresses, which contains at most two IP addresses.",
			},

			"next_hop_type": {
				// Required:    true,
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The type of the next hop. Currently supported types are:VPN: VPN gateway;DIRECTCONNECT: direct connect gateway;PEERCONNECTION: peering connection;NAT: NAT gateway;NORMAL_CVM: normal CVM.",
			},

			"next_hop_destination": {
				// Required:    true,
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The next-hop destination gateway. The value is related to NextHopType.If NextHopType is set to VPN, the value of this parameter is the VPN gateway ID, such as vpngw-12345678.If NextHopType is set to DIRECTCONNECT, the value of this parameter is the direct connect gateway ID, such as dcg-12345678.If NextHopType is set to PEERCONNECTION, the value of this parameter is the peering connection ID, such as pcx-12345678.If NextHopType is set to NAT, the value of this parameter is the NAT gateway ID, such as nat-12345678.If NextHopType is set to NORMAL_CVM, the value of this parameter is the IPv4 address of the CVM, such as 203.0.113.12.",
			},

			"net_detect_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "ID of a network inspector instance, e.g. netd-12345678. Enter at least one of this parameter, VpcId, SubnetId, and NetDetectName. Use NetDetectId if it is present.",
			},

			"vpc_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "ID of a `VPC` instance, e.g. `vpc-12345678`, which is used together with SubnetId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.",
			},

			"subnet_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "ID of a subnet instance, e.g. `subnet-12345678`, which is used together with VpcId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.",
			},

			"net_detect_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of a network inspector, up to 60 bytes in length. It is used together with VpcId and NetDetectName. You should enter either this parameter or NetDetectId, or both. Use NetDetectId if it is present.",
			},

			"net_detect_ip_state_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "The array of network detection verification results.",
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

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudVpcNetDetectStateCheckRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_net_detect_state_check.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("detect_destination_ip"); ok {
		detectDestinationIpSet := v.(*schema.Set).List()
		paramMap["DetectDestinationIp"] = helper.InterfacesStringsPoint(detectDestinationIpSet)
	}

	if v, ok := d.GetOk("next_hop_type"); ok {
		paramMap["NextHopType"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("next_hop_destination"); ok {
		paramMap["NextHopDestination"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("net_detect_id"); ok {
		paramMap["NetDetectId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		paramMap["VpcId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("subnet_id"); ok {
		paramMap["SubnetId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("net_detect_name"); ok {
		paramMap["NetDetectName"] = helper.String(v.(string))
	}

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var netDetectIpStateSet []*vpc.NetDetectIpState

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcNetDetectStateCheck(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		netDetectIpStateSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(netDetectIpStateSet))
	tmpList := make([]map[string]interface{}, 0, len(netDetectIpStateSet))

	if netDetectIpStateSet != nil {
		for _, netDetectIpState := range netDetectIpStateSet {
			netDetectIpStateMap := map[string]interface{}{}

			if netDetectIpState.DetectDestinationIp != nil {
				netDetectIpStateMap["detect_destination_ip"] = netDetectIpState.DetectDestinationIp
			}

			if netDetectIpState.State != nil {
				netDetectIpStateMap["state"] = netDetectIpState.State
			}

			if netDetectIpState.Delay != nil {
				netDetectIpStateMap["delay"] = netDetectIpState.Delay
			}

			if netDetectIpState.PacketLossRate != nil {
				netDetectIpStateMap["packet_loss_rate"] = netDetectIpState.PacketLossRate
			}

			ids = append(ids, *netDetectIpState.DetectDestinationIp)
			tmpList = append(tmpList, netDetectIpStateMap)
		}

		_ = d.Set("net_detect_ip_state_set", tmpList)
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

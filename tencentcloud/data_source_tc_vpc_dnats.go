/*
Use this data source to query detailed information of DNATs.

# Example Usage

```hcl
# query by nat gateway id

	data "tencentcloudenterprise_vpc_dnats" "foo" {
	  nat_id = "nat-xfaq1"
	}

# query by vpc id

	data "tencentcloudenterprise_vpc_dnats" "foo" {
	  vpc_id = "vpc-xfqag"
	}

# query by elastic ip

	data "tencentcloudenterprise_vpc_dnats" "foo" {
	  elastic_ip = "203.0.113.1"
	}

```
*/
package tencentcloud

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_dnats", CNDescription{
		TerraformTypeCN: "NAT网关端口转发规则对象数组",
		DescriptionCN:   "提供VPC DNAT数据源，用于查询VPC DNAT的详细信息。",
		AttributesCN: map[string]string{
			"vpc_id":             "专有网络的ID",
			"nat_id":             "NAT网关的ID",
			"elastic_ip":         "EIP的网络地址",
			"elastic_port":       "EIP的端口",
			"private_ip":         "后端服务的网络地址",
			"description":        "NAT转发的描述",
			"private_port":       "intranet的端口",
			"result_output_file": "用于保存结果",
			"dnat_list":          "DNAT的信息列表",
			"protocol":           "网络协议的类型有效值：“TCP”和“UDP”",
		},
	})
}

func dataSourceTencentCloudDnats() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides details about a specific DNAT instance.",
		Read:        dataSourceTencentCloudDnatsRead,

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the VPC.",
			},
			"nat_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the NAT gateway.",
			},
			"elastic_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIp,
				Description:  "Network address of the EIP.",
			},
			"elastic_port": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validatePort,
				Description:  "Port of the EIP.",
			},
			"private_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIp,
				Description:  "Network address of the backend service.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the NAT forward.",
			},
			"private_port": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validatePort,
				Description:  "Port of intranet.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// Computed values
			"dnat_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information list of the DNATs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the VPC.",
						},
						"nat_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the NAT.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the network protocol. Valid values: `TCP` and `UDP`.",
						},
						"elastic_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network address of the EIP.",
						},
						"elastic_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port of the EIP.",
						},
						"private_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network address of the backend service.",
						},
						"private_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Port of intranet.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of the NAT forward.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudDnatsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_dnats.read")()

	logId := getLogId(contextNil)
	request := vpc.NewDescribeNatGatewayDestinationIpPortTranslationNatRulesRequest()
	params := make(map[string]string)
	if v, ok := d.GetOk("nat_id"); ok {
		params["nat-gateway-id"] = v.(string)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		params["vpc-id"] = v.(string)
	}
	if v, ok := d.GetOk("elastic_ip"); ok {
		params["public-ip-address"] = v.(string)
	}
	if v, ok := d.GetOk("elastic_port"); ok {
		params["public-port"] = v.(string)
	}
	if v, ok := d.GetOk("private_ip"); ok {
		params["private-ip-address"] = v.(string)
	}
	if v, ok := d.GetOk("private_port"); ok {
		params["private-port"] = v.(string)
	}
	if v, ok := d.GetOk("description"); ok {
		params["description"] = v.(string)
	}
	request.Filters = make([]*vpc.Filter, 0, len(params))
	for k, v := range params {
		filter := &vpc.Filter{
			Name:   helper.String(k),
			Values: []*string{helper.String(v)},
		}
		request.Filters = append(request.Filters, filter)
	}
	var response *vpc.DescribeNatGatewayDestinationIpPortTranslationNatRulesResponse

	offset := uint64(0)
	request.Offset = &offset
	result := make([]*vpc.NatGatewayDestinationIpPortTranslationNatRule, 0)
	limit := uint64(NAT_DESCRIBE_LIMIT)
	request.Limit = &limit
	for {
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().DescribeNatGatewayDestinationIpPortTranslationNatRules(request)
			if e != nil {
				log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
					logId, request.GetAction(), request.ToJsonString(), e.Error())
				return retryError(e)
			}
			response = result
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s read DNAT failed, reason:%s\n", logId, err.Error())
			return err
		} else {
			result = append(result, response.Response.NatGatewayDestinationIpPortTranslationNatRuleSet...)
			if len(response.Response.NatGatewayDestinationIpPortTranslationNatRuleSet) < NAT_DESCRIBE_LIMIT {
				break
			} else {
				offset = offset + limit
			}
		}
	}
	ids := make([]string, 0, len(result))
	dnatList := make([]map[string]interface{}, 0, len(result))
	for _, dnat := range result {
		mapping := map[string]interface{}{
			"nat_id":       *dnat.NatGatewayId,
			"vpc_id":       *dnat.VpcId,
			"elastic_ip":   *dnat.PublicIpAddress,
			"elastic_port": strconv.Itoa(int(*dnat.PublicPort)),
			"private_ip":   *dnat.PrivateIpAddress,
			"private_port": strconv.Itoa(int(*dnat.PrivatePort)),
			"protocol":     *dnat.IpProtocol,
			"description":  *dnat.Description,
		}
		dnatList = append(dnatList, mapping)
		var entry = &vpc.DestinationIpPortTranslationNatRule{}
		entry.IpProtocol = dnat.IpProtocol
		entry.PublicIpAddress = dnat.PublicIpAddress
		entry.PublicPort = dnat.PublicPort
		ids = append(ids, buildDnatId(entry, *dnat.VpcId, *dnat.NatGatewayId))
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("dnat_list", dnatList); e != nil {
		log.Printf("[CRITAL]%s provider set DNAT list fail, reason:%s\n", logId, e.Error())
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), dnatList); e != nil {
			return e
		}
	}

	return nil

}

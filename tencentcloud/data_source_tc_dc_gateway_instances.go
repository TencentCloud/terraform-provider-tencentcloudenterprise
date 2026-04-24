/*
Use this data source to query detailed information of direct connect gateway instances.

Example Usage

```hcl
resource "tencentcloudenterprise_ccn" "main" {
  name        = "ci-temp-test-ccn"
  description = "ci-temp-test-ccn-des"
  qos         = "AG"
}

resource "tencentcloudenterprise_vpc_dc_gateway" "ccn_main" {
  name                = "ci-cdg-ccn-test"
  network_instance_id = tencentcloudenterprise_ccn.main.id
  network_type        = "CCN"
  gateway_type        = "NORMAL"
}

#You need to sleep for a few seconds because there is a cache on the server
data "tencentcloudenterprise_dc_gateway_instances" "name_select" {
  name = tencentcloudenterprise_vpc_dc_gateway.ccn_main.name
}

data "tencentcloudenterprise_dc_gateway_instances" "id_select" {
  dcg_id = tencentcloudenterprise_vpc_dc_gateway.ccn_main.id
}
```
*/
package tencentcloud

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dc_gateway_instances", CNDescription{
		TerraformTypeCN: "专线网关实例",
		DescriptionCN:   "提供专线网关实例数据源查询，用于获取专线网关实例列表信息。",
		AttributesCN: map[string]string{
			"dcg_id":               "查询的专线网关ID",
			"name":                 "查询的专线网关名称",
			"result_output_file":   "结果输出文件路径",
			"instance_list":        "专线网关实例列表",
			"instance_list.dcg_id": "专线网关ID",
			"instance_list.name":   "专线网关名称",
			"instance_list.dcg_ip": "专线网关IP",
			"instance_list.network_type":        "关联网络类型，可选值：VPC、CCN",
			"instance_list.network_instance_id": "关联网络实例ID",
			"instance_list.gateway_type":        "网关类型，可选值：NORMAL、NAT",
			"instance_list.cnn_route_type":      "云联网路由学习类型，可选值：BGP、STATIC",
			"instance_list.enable_bgp":          "是否启用BGP",
			"instance_list.create_time":         "创建时间",
			"instance_list.mode_type":           "是否发布云联网VPC的Cidr，standard:发布，exquisite:不发布",
		},
	})
}

func dataSourceTencentCloudDcGatewayInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudDcGatewayInstancesRead,
		Schema: map[string]*schema.Schema{
			"dcg_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the DCG to be queried.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of the DCG to be queried.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			// Computed values
			"instance_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information list of the DCG.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dcg_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the DCG.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the DCG.",
						},
						"dcg_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP of the DCG.",
						},
						"network_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of associated network. Valid values: `VPC` and `CCN`.",
						},
						"network_instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the associated network instance.",
						},
						"gateway_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of the gateway. Valid values: `NORMAL` and `NAT`.",
						},
						"cnn_route_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of CCN route. Valid values: `BGP` and `STATIC`.",
						},
						"enable_bgp": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the BGP is enabled.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of resource.",
						},
						"mode_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to publish VPC CIDR to CCN. Valid values: `standard` and `exquisite`.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudDcGatewayInstancesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_dc_gateway_instances.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		id   = ""
		name = ""
	)

	if temp, ok := d.GetOk("dcg_id"); ok {
		if tempStr := temp.(string); tempStr != "" {
			id = tempStr
		}
	}

	if temp, ok := d.GetOk("name"); ok {
		if tempStr := temp.(string); tempStr != "" {
			name = tempStr
		}
	}

	var infos, err = service.DescribeDirectConnectGatewaysByFilter(ctx, id, name)
	if err != nil {
		return err
	}

	var infoList = make([]map[string]interface{}, 0, len(infos))

	for _, item := range infos {
		var infoMap = make(map[string]interface{})
		infoMap["dcg_id"] = item.dcgId
		infoMap["name"] = item.name
		infoMap["dcg_ip"] = item.dcgIp
		infoMap["network_type"] = item.networkType
		infoMap["network_instance_id"] = item.networkInstanceId
		infoMap["gateway_type"] = item.gatewayType
		infoMap["cnn_route_type"] = item.cnnRouteType
		infoMap["create_time"] = item.createTime
		infoMap["enable_bgp"] = item.enableBGP
		infoMap["mode_type"] = item.modeType

		infoList = append(infoList, infoMap)
	}
	if err := d.Set("instance_list", infoList); err != nil {
		log.Printf("[CRITAL]%s provider set  dcg instances fail, reason:%s\n ",
			logId,
			err.Error())
		return err
	}

	m := md5.New()
	_, err = m.Write([]byte("dcg_instances" + id + "_" + name))
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%x", m.Sum(nil)))

	if output, ok := d.GetOk("result_output_file"); ok && output.(string) != "" {
		if err := writeToFile(output.(string), infoList); err != nil {
			log.Printf("[CRITAL]%s output file[%s] fail, reason[%s]\n",
				logId,
				output.(string),
				err.Error())
			return err
		}
	}
	return nil

}

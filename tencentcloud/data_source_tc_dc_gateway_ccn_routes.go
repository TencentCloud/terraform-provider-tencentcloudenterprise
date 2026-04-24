/*
Use this data source to query detailed information of direct connect gateway route entries.

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

resource "tencentcloudenterprise_dc_gateway_ccn_route" "route1" {
  dcg_id     = tencentcloudenterprise_vpc_dc_gateway.ccn_main.id
  cidr_block = "10.1.1.0/32"
}

resource "tencentcloudenterprise_dc_gateway_ccn_route" "route2" {
  dcg_id     = tencentcloudenterprise_vpc_dc_gateway.ccn_main.id
  cidr_block = "192.1.1.0/32"
}

#You need to sleep for a few seconds because there is a cache on the server
data "tencentcloudenterprise_dc_gateway_ccn_routes" "test" {
  dcg_id = tencentcloudenterprise_vpc_dc_gateway.ccn_main.id
}
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_dc_gateway_ccn_routes", CNDescription{
		TerraformTypeCN: "专线网关云联网路由",
		DescriptionCN:   "提供专线网关云联网路由数据源查询，用于获取专线网关的CCN路由条目列表。",
		AttributesCN: map[string]string{
			"dcg_id":               "查询的专线网关ID",
			"ccn_route_type":       "云联网路由学习类型，可选值：BGP（自动学习）、STATIC（用户配置），默认STATIC",
			"result_output_file":   "结果输出文件路径",
			"instance_list":        "路由条目列表",
			"instance_list.dcg_id": "专线网关ID",
			"instance_list.route_id":    "路由条目ID",
			"instance_list.cidr_block":  "目标网段",
			"instance_list.as_path":     "BGP AS路径列表",
		},
	})
}

func dataSourceTencentCloudDcGatewayCCNRoutes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudDcGatewayCCNRoutesRead,
		Schema: map[string]*schema.Schema{
			"dcg_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the DCG to be queried.",
			},
			"ccn_route_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cloud networking routing learning type, optional values: BGP - Automatic Learning; STATIC - User configured. Default is STATIC.",
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
				Description: "Information list of the DCG route entries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dcg_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the DCG.",
						},
						"route_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the DCG route.",
						},
						"cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A network address segment of IDC.",
						},
						"as_path": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "As path list of the BGP.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudDcGatewayCCNRoutesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_dc_gateway_ccn_routes.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		id           string
		ccnRouteType string
	)

	if v, ok := d.GetOk("dcg_id"); ok {
		id = v.(string)
	}

	if v, ok := d.GetOk("ccn_route_type"); ok {
		ccnRouteType = v.(string)
	}

	var infos, err = service.DescribeDirectConnectGatewayCcnRoutes(ctx, id, ccnRouteType)
	if err != nil {
		return err
	}

	var infoList = make([]map[string]interface{}, 0, len(infos))

	for _, item := range infos {
		var infoMap = make(map[string]interface{})
		infoMap["dcg_id"] = item.dcgId
		infoMap["route_id"] = item.routeId
		infoMap["cidr_block"] = item.cidrBlock
		infoMap["as_path"] = item.asPaths
		infoList = append(infoList, infoMap)
	}
	if err := d.Set("instance_list", infoList); err != nil {
		log.Printf("[CRITAL]%s provider set  dcg  ccn routes fail, reason:%s\n ",
			logId,
			err.Error())
		return err
	}

	d.SetId(id)

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

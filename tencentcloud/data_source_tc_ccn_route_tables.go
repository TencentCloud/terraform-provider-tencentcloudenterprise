/*
Use this data source to query CCN (Cloud Connect Network) route tables.

# Example Usage

```hcl

data "tencentcloudenterprise_ccn_route_tables" "by_ccn" {
  ccn_id = "ccn-xxxxxxxx"
}

data "tencentcloudenterprise_ccn_route_tables" "by_name" {
  ccn_id           = "ccn-xxxxxxxx"
  route_table_name = "my-rtb"
}

data "tencentcloudenterprise_ccn_route_tables" "by_id" {
  route_table_id = "ccnrtb-xxxxxxxx"
}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ccnRouteTableInfo() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ccn_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the CCN instance.",
		},
		"route_table_id": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "ID of the CCN route table.",
		},
		"route_table_name": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Name of the CCN route table.",
		},
		"route_table_description": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Description of the CCN route table.",
		},
		"is_default_table": {
			Type:        schema.TypeBool,
			Computed:    true,
			Description: "True: default route table; False: custom route table.",
		},
		"create_time": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Creation time of the route table.",
		},
	}
}

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ccn_route_tables", CNDescription{
		TerraformTypeCN: "ccn路由表",
		DescriptionCN:   "提供CCN（云联网）路由表数据源，用于查询云联网路由表的详细信息。支持按CCN ID/路由表ID/路由表名称/路由表描述过滤。",
		AttributesCN: map[string]string{
			"ccn_id":                  "云联网实例ID（过滤条件）",
			"route_table_id":          "路由表ID（过滤条件，精确查询）",
			"route_table_name":        "路由表名称（过滤条件，模糊查询）",
			"route_table_description": "路由表描述（过滤条件）",
			"is_default_table":        "是否默认路由表",
			"create_time":             "创建时间",
			"result_output_file":      "用于保存结果",
			"list":                    "路由表信息列表",
		},
	})
}

func dataSourceTencentCloudCcnRouteTables() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a data source to query CCN route tables.",
		Read:        dataSourceTencentCloudCcnRouteTablesRead,
		Schema: map[string]*schema.Schema{
			"ccn_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by CCN instance ID.",
			},
			"route_table_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by exact route table ID.",
			},
			"route_table_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by route table name (fuzzy match).",
			},
			"route_table_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by route table description.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of CCN route tables. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: ccnRouteTableInfo(),
				},
			},
		},
	}
}

func dataSourceTencentCloudCcnRouteTablesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ccn_route_tables.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var (
		ccnId              string
		routeTableId       string
		routeTableName     string
		routeTableDesc     string
	)

	if v, ok := d.GetOk("ccn_id"); ok {
		ccnId = v.(string)
	}
	if v, ok := d.GetOk("route_table_id"); ok {
		routeTableId = v.(string)
	}
	if v, ok := d.GetOk("route_table_name"); ok {
		routeTableName = v.(string)
	}
	if v, ok := d.GetOk("route_table_description"); ok {
		routeTableDesc = v.(string)
	}

	infos, err := service.DescribeCcnRouteTables(ctx, ccnId, routeTableId, routeTableName, routeTableDesc)
	if err != nil && routeTableId == "" {
		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			infos, err = service.DescribeCcnRouteTables(ctx, ccnId, routeTableId, routeTableName, routeTableDesc)
			if err != nil {
				return retryError(err)
			}
			return nil
		})
	}
	if err != nil {
		return err
	}

	list := make([]map[string]interface{}, 0, len(infos))
	for _, info := range infos {
		infoMap := map[string]interface{}{}
		infoMap["ccn_id"] = info.CcnId
		infoMap["route_table_id"] = info.CcnRouteTableId
		infoMap["route_table_name"] = info.RouteTableName
		infoMap["route_table_description"] = info.RouteTableDescription
		infoMap["is_default_table"] = info.IsDefaultTable
		infoMap["create_time"] = info.CreateTime
		list = append(list, infoMap)
	}

	d.SetId("CcnRouteTables" + ccnId + routeTableId + routeTableName + routeTableDesc)
	err = d.Set("list", list)
	if err != nil {
		log.Printf("[CRITAL]%s provider set tencentcloudenterprise_ccn_route_tables list fail, reason:%s\n ", logId, err.Error())
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err = writeToFile(output.(string), list); err != nil {
			return err
		}
	}
	return nil
}

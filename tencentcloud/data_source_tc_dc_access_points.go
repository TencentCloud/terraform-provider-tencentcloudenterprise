/*
Use this data source to query detailed information of dc access_points

# Example Usage

```hcl

	data "tencentcloudenterprise_dc_access_points" "access_points" {
	  region_id = "ap-guangzhou"
	}

```
*/
package tencentcloud

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dc "terraform-provider-tencentcloudenterprise/sdk/dc/v20180410"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_dc_access_points", CNDescription{
		TerraformTypeCN: "云物理服务器接入点",
		AttributesCN: map[string]string{
			"region_id":           "接入点区域，可以通过`DescribeRegions`查询，可以调用`DescribeRegions`获取区域ID",
			"access_point_set":    "接入点信息",
			"access_point_name":   "接入点名称",
			"access_point_id":     "唯一的访问点ID",
			"state":               "接入点状态有效值：可用、不可用",
			"location":            "接入点位置",
			"line_operator":       "接入点支持的ISP列表",
			"available_port_type": "访问点上的可用端口类型有效值：1000BASE-T：千兆电口；1000BASE-LX:10km千兆单模光端口；1000BASE-ZX:80km千兆单模光端口；10GBASE-LR：10km万兆单模光端口；10GBASE-ZR:80km万兆单模光端口；10GBASE-LH:40km万兆单模光端口；100GBASE-LR4:10 km 100千兆单模光端口光纤端口注意：此字段可能返回“null”，表示没有得到有效值",
			"city":                "访问点所在的城市注意：此字段可能返回“null”，表示无法获得有效值",
			"area":                "接入点区域注意：此字段可能返回“null”，表示无法获得有效值",
			"access_point_type":   "接入点类型有效值：“VXLAN”、“QCPL”和“QCAR”注意：此字段可能返回“null”，表示无法获得有效值",
		},
	})
}

func dataSourceTencentCloudDcAccessPoints() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of dc access_points",
		Read:        dataSourceTencentCloudDcAccessPointsRead,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Access point region, which can be queried through `DescribeRegions`.You can call `DescribeRegions` to get the region ID.",
			},

			"access_point_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Access point information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_point_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access point name.",
						},
						"access_point_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique access point ID.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access point status. Valid values: available, unavailable.",
						},
						"location": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access point location.",
						},
						"line_operator": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "List of ISPs supported by access point.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the region that manages the access point.",
						},
						"available_port_type": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "Available port type at the access point. Valid values: 1000BASE-T: gigabit electrical port; 1000BASE-LX: 10 km gigabit single-mode optical port; 1000BASE-ZX: 80 km gigabit single-mode optical port; 10GBASE-LR: 10 km 10-gigabit single-mode optical port; 10GBASE-ZR: 80 km 10-gigabit single-mode optical port; 10GBASE-LH: 40 km 10-gigabit single-mode optical port; 100GBASE-LR4: 10 km 100-gigabit single-mode optical portfiber optic port.Note: this field may return `null`, indicating that no valid value is obtained.",
						},
						//"coordinate": {
						//	Type:        schema.TypeList,
						//	Computed:    true,
						//	Description: "Latitude and longitude of the access pointNote: this field may return `null`, indicating that no valid values can be obtained.",
						//	Elem: &schema.Resource{
						//		Schema: map[string]*schema.Schema{
						//			"lat": {
						//				Type:        schema.TypeFloat,
						//				Computed:    true,
						//				Description: "Latitude.",
						//			},
						//			"lng": {
						//				Type:        schema.TypeFloat,
						//				Computed:    true,
						//				Description: "Longitude.",
						//			},
						//		},
						//	},
						//},
						"city": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "City where the access point is locatedNote: this field may return `null`, indicating that no valid values can be obtained.",
						},
						"area": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access point regionNote: this field may return `null`, indicating that no valid values can be obtained.",
						},
						"access_point_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access point type. Valid values: `VXLAN`, `QCPL`, and `QCAR`.Note: this field may return `null`, indicating that no valid values can be obtained.",
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

func dataSourceTencentCloudDcAccessPointsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_dc_access_points.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("region_id"); ok {
		paramMap["RegionId"] = helper.String(v.(string))
	}

	service := DcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var accessPointSet []*dc.AccessPoint

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeDcAccessPointsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		accessPointSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(accessPointSet))
	tmpList := make([]map[string]interface{}, 0, len(accessPointSet))

	if accessPointSet != nil {
		for _, accessPoint := range accessPointSet {
			accessPointMap := map[string]interface{}{}

			if accessPoint.AccessPointName != nil {
				accessPointMap["access_point_name"] = accessPoint.AccessPointName
			}

			if accessPoint.AccessPointId != nil {
				accessPointMap["access_point_id"] = accessPoint.AccessPointId
			}

			if accessPoint.State != nil {
				accessPointMap["state"] = accessPoint.State
			}

			if accessPoint.Location != nil {
				accessPointMap["location"] = accessPoint.Location
			}

			if accessPoint.LineOperator != nil {
				accessPointMap["line_operator"] = accessPoint.LineOperator
			}

			if accessPoint.Region != nil {
				accessPointMap["region"] = accessPoint.Region
			}

			if accessPoint.CloudPortType != nil {
				accessPointMap["tencentcloudenterprise_port_type"] = accessPoint.CloudPortType
			}

			if accessPoint.City != nil {
				accessPointMap["city"] = accessPoint.City
			}

			if accessPoint.IdcPortType != nil {
				accessPointMap["idc_port_type"] = accessPoint.IdcPortType
			}

			ids = append(ids, strconv.FormatInt(*accessPoint.AccessPointId, 10))
			tmpList = append(tmpList, accessPointMap)
		}

		_ = d.Set("access_point_set", tmpList)
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

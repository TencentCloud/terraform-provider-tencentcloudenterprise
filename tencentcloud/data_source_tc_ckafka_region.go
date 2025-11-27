// /*
// Use this data source to query detailed information of ckafka region
//
// # Example Usage
//
// ```hcl
// data "tencentcloudenterprise_ckafka_region" "region" {
// }
// ```
// */
package tencentcloud

//
//import (
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//)
//
//func init() {
//	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_region", CNDescription{
//		TerraformTypeCN: "ckafka 区域配置",
//		AttributesCN: map[string]string{
//			"result":             "返回区域枚举结果的列表",
//			"region_id":          "区域ID",
//			"region_name":        "地理名称",
//			"area_name":          "区域名称",
//			"region_code":        "地区代码",
//			"region_code_v3":     "地区代码（V3版本）",
//			"support":            "NONE：默认值不支持任何特殊型号CVM：支持CVM类型",
//			"ipv6":               "是否支持ipv6，0:表示不支持，1:表示支持",
//			"multi_zone":         "是否支持跨可用区，0:表示不支持，1:表示支持",
//			"result_output_file": "用于保存结果",
//		},
//	})
//}
//
//func dataSourceTencentCloudCkafkaRegion() *schema.Resource {
//	return &schema.Resource{
//		Description: "Use this data source to query detailed information of ckafka region",
//		Read:        dataSourceTencentCloudCkafkaRegionRead,
//		Schema: map[string]*schema.Schema{
//			"result": {
//				Computed:    true,
//				Type:        schema.TypeList,
//				Description: "Return a list of region enumeration results.",
//				Elem: &schema.Resource{
//					Schema: map[string]*schema.Schema{
//						"region_id": {
//							Type:        schema.TypeInt,
//							Computed:    true,
//							Description: "Region ID.",
//						},
//						"region_name": {
//							Type:        schema.TypeString,
//							Computed:    true,
//							Description: "Geographical name.",
//						},
//						"area_name": {
//							Type:        schema.TypeString,
//							Computed:    true,
//							Description: "Area name.",
//						},
//						"region_code": {
//							Type:        schema.TypeString,
//							Computed:    true,
//							Description: "Region Code.",
//						},
//						"region_code_v3": {
//							Type:        schema.TypeString,
//							Computed:    true,
//							Description: "Region Code(V3 version).",
//						},
//						"support": {
//							Type:        schema.TypeString,
//							Computed:    true,
//							Description: "NONE: The default value does not support any special models CVM: Supports CVM types.",
//						},
//						"ipv6": {
//							Type:        schema.TypeInt,
//							Computed:    true,
//							Description: "Whether to support ipv6, 0: means not supported, 1: means supported.",
//						},
//						"multi_zone": {
//							Type:        schema.TypeInt,
//							Computed:    true,
//							Description: "Whether to support cross-availability zones, 0: means not supported, 1: means supported.",
//						},
//					},
//				},
//			},
//
//			"result_output_file": {
//				Type:        schema.TypeString,
//				Optional:    true,
//				Description: "Used to save results.",
//			},
//		},
//	}
//}
//
////func dataSourceTencentCloudCkafkaRegionRead(d *schema.ResourceData, meta interface{}) error {
////	defer logElapsed("data_source.cloud_ckafka_region.read")()
////	defer inconsistentCheck(d, meta)()
////
////	var result []*ckafka.Region
////	request := ckafka.NewDescribeRegionRequest()
////	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
////		response, e := meta.(*TencentCloudClient).apiV3Conn.UseCkafkaClient().DescribeRegion(request)
////		if e != nil {
////			return retryError(e)
////		}
////		if response == nil || response.Response == nil {
////			return retryError(errors.New("Response is null"))
////		}
////		result = response.Response.Result
////		return nil
////	})
////	if err != nil {
////		return err
////	}
////
////	ids := make([]string, 0, len(result))
////	tmpList := make([]map[string]interface{}, 0, len(result))
////
////	if result != nil {
////		for _, region := range result {
////			regionMap := map[string]interface{}{}
////
////			if region.RegionId != nil {
////				regionMap["region_id"] = region.RegionId
////			}
////
////			if region.RegionName != nil {
////				regionMap["region_name"] = region.RegionName
////			}
////
////			if region.AreaName != nil {
////				regionMap["area_name"] = region.AreaName
////			}
////
////			if region.RegionCode != nil {
////				regionMap["region_code"] = region.RegionCode
////			}
////
////			if region.RegionCodeV3 != nil {
////				regionMap["region_code_v3"] = region.RegionCodeV3
////			}
////
////			if region.Support != nil {
////				regionMap["support"] = region.Support
////			}
////
////			if region.Ipv6 != nil {
////				regionMap["ipv6"] = region.Ipv6
////			}
////
////			if region.MultiZone != nil {
////				regionMap["multi_zone"] = region.MultiZone
////			}
////
////			ids = append(ids, strconv.FormatInt(*region.RegionId, 10))
////			tmpList = append(tmpList, regionMap)
////		}
////
////		_ = d.Set("result", tmpList)
////	}
////
////	d.SetId(helper.DataResourceIdsHash(ids))
////	output, ok := d.GetOk("result_output_file")
////	if ok && output.(string) != "" {
////		if e := writeToFile(output.(string), tmpList); e != nil {
////			return e
////		}
////	}
////	return nil
////}

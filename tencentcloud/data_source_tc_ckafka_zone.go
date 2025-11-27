/*
Use this data source to query detailed information of ckafka zone

# Example Usage

```hcl
data "tencentcloudenterprise_ckafka_zone" "ckafka_zone" {
}
```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ckafka "terraform-provider-tencentcloudenterprise/sdk/ckafka/v20190819"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_ckafka_zone", CNDescription{
		TerraformTypeCN: "ckafka区域信息",
		DescriptionCN:   "提供CKafka区域信息数据源，用于查询CKafka可用区域的详细信息。",
		AttributesCN: map[string]string{
			"cdc_id":          "cdc专业集群业务参数",
			"result":          "查询结果复杂对象实体",
			"zone_list":       "区域列表",
			"zone_id":         "区域id",
			"is_internal_app": "内部APP",
			"app_id":          "应用程序id",
			"flag":            "旗帜",
			"zone_name":       "区域名称",
			"zone_status":     "区域状态",
			"exflag":          "额外的标志",
			"sold_out":        "json对象，key是模型，value true是卖光的，false不是卖光的",
			//"sales_info":           "标准版售罄信息",
			"version":              "ckakfa版本（1.1.1/2.4.2/0.10.2）",
			"platform":             "专业版、标准版标志",
			"max_buy_instance_num": "已购买实例的最大数量",
			"max_bandwidth":        "最大购买带宽（MB）",
			"unit_price":           "后付费单价",
			"real_total_cost":      "折扣价",
			"total_cost":           "原价",
			"message_price":        "后付费消息单价",
			"cluster_info":         "用户专用群集信息",
			"cluster_id":           "ClusterId",
			"cluster_name":         "ClusterName",
			"max_disk_size":        "群集中最大的磁盘，以GB为单位",
			"max_band_width":       "最大群集带宽（MB）",
			"available_disk_size":  "群集的当前可用磁盘，以GB为单位",
			"available_band_width": "群集的当前可用带宽（MB）",
			"zone_ids":             "群集节点所在的可用性区域如果群集是跨可用性区域群集，则它包括群集节点所在的多个可用性区域",
			//"standard":             "购买标准版配置",
			//"standard_s2":          "标准版S2配置",
			//"profession":           "专业版配置",
			//"physical":             "物理独占版配置",
			//"public_network":       "公共网络带宽",
			//"public_network_limit": "公用网络带宽配置",
			//"result_output_file":   "用于保存结果",
		},
	})
}

func dataSourceTencentCloudCkafkaZone() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of ckafka zone",
		Read:        dataSourceTencentCloudCkafkaZoneRead,
		Schema: map[string]*schema.Schema{
			//"cdc_id": {
			//	Optional:    true,
			//	Type:        schema.TypeString,
			//	Description: "Cdc professional cluster business parameters.",
			//},

			"result": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Query result complex object entity.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_list": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Zone list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Zone id.",
									},
									"is_internal_app": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Internal APP.",
									},
									"app_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "App id.",
									},
									"flag": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Flag.",
									},
									"zone_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Zone name.",
									},
									"zone_status": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Zone status.",
									},
									"exflag": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Extra flag.",
									},
									"sold_out": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Json object, key is model, value true is sold out, false is not sold out.",
									},
									//"sales_info": {
									//	Type:        schema.TypeList,
									//	Computed:    true,
									//	Description: "Standard Edition Sold Out Information.",
									//	Elem: &schema.Resource{
									//		Schema: map[string]*schema.Schema{
									//			"flag": {
									//				Type:        schema.TypeBool,
									//				Computed:    true,
									//				Description: "Manually set flags.",
									//			},
									//			"version": {
									//				Type:        schema.TypeString,
									//				Computed:    true,
									//				Description: "Ckakfa version(1.1.1/2.4.2/0.10.2).",
									//			},
									//			"platform": {
									//				Type:        schema.TypeString,
									//				Computed:    true,
									//				Description: "Professional Edition, Standard Edition flag.",
									//			},
									//			"sold_out": {
									//				Type:        schema.TypeBool,
									//				Computed:    true,
									//				Description: "Sold out flag: true sold out.",
									//			},
									//		},
									//	},
									//},
								},
							},
						},
						"max_buy_instance_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The maximum number of purchased instances.",
						},
						"max_bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum purchased bandwidth in Mbs.",
						},
						"unit_price": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Postpaid unit price.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"real_total_cost": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Discount price.",
									},
									"total_cost": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Original price.",
									},
								},
							},
						},
						"message_price": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Postpaid message unit price.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"real_total_cost": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Discount price.",
									},
									"total_cost": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "Original price.",
									},
								},
							},
						},
						"cluster_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "User exclusive cluster information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cluster_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "ClusterId.",
									},
									"cluster_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ClusterName.",
									},
									"max_disk_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The largest disk in the cluster, in GB.",
									},
									"max_band_width": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum cluster bandwidth in MBs.",
									},
									"available_disk_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The current available disk of the cluster, in GB.",
									},
									"available_band_width": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The current available bandwidth of the cluster in MBs.",
									},
									"zone_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Availability zone to which the cluster belongs, indicating the availability zone to which the cluster belongs.",
									},
									"zone_ids": {
										Type: schema.TypeSet,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
										Computed:    true,
										Description: "The availability zone where the cluster node is located. If the cluster is a cross-availability zone cluster, it includes multiple availability zones where the cluster node is located.",
									},
								},
							},
						},
						//"standard": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "Purchase Standard Edition Configuration.",
						//},
						//"standard_s2": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "Standard Edition S2 configuration.",
						//},
						//"profession": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "Professional Edition configuration.",
						//},
						//"physical": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "Physical Exclusive Edition Configuration.",
						//},
						//"public_network": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "Public network bandwidth.",
						//},
						//"public_network_limit": {
						//	Type:        schema.TypeString,
						//	Computed:    true,
						//	Description: "Public network bandwidth configuration.",
						//},
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

func dataSourceTencentCloudCkafkaZoneRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_zone.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("cdc_id"); ok {
		paramMap["CdcId"] = helper.String(v.(string))
	}

	service := CkafkaService{client: meta.(*TencentCloudClient).apiV3Conn}

	var result *ckafka.ZoneResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, e := service.DescribeCkafkaCkafkaZoneByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		result = response
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0)
	zoneResponseMapList := make([]interface{}, 0)
	if result != nil {
		zoneResponseMap := map[string]interface{}{}

		if result.ZoneList != nil {
			zoneListList := []interface{}{}
			for _, zoneList := range result.ZoneList {
				zoneListMap := map[string]interface{}{}

				if zoneList.ZoneId != nil {
					ids = append(ids, *zoneList.ZoneId)
					zoneListMap["zone_id"] = zoneList.ZoneId
				}

				if zoneList.IsInternalApp != nil {
					zoneListMap["is_internal_app"] = zoneList.IsInternalApp
				}

				if zoneList.AppId != nil {
					zoneListMap["app_id"] = zoneList.AppId
				}

				if zoneList.Flag != nil {
					zoneListMap["flag"] = zoneList.Flag
				}

				if zoneList.ZoneName != nil {
					zoneListMap["zone_name"] = zoneList.ZoneName
				}

				if zoneList.ZoneStatus != nil {
					zoneListMap["zone_status"] = zoneList.ZoneStatus
				}

				if zoneList.Exflag != nil {
					zoneListMap["exflag"] = zoneList.Exflag
				}

				if zoneList.SoldOut != nil {
					zoneListMap["sold_out"] = zoneList.SoldOut
				}

				//if zoneList.SalesInfo != nil {
				//	salesInfoList := []interface{}{}
				//	for _, salesInfo := range zoneList.SalesInfo {
				//		salesInfoMap := map[string]interface{}{}
				//
				//		if salesInfo.Flag != nil {
				//			salesInfoMap["flag"] = salesInfo.Flag
				//		}
				//
				//		if salesInfo.Version != nil {
				//			salesInfoMap["version"] = salesInfo.Version
				//		}
				//
				//		if salesInfo.Platform != nil {
				//			salesInfoMap["platform"] = salesInfo.Platform
				//		}
				//
				//		if salesInfo.SoldOut != nil {
				//			salesInfoMap["sold_out"] = salesInfo.SoldOut
				//		}
				//
				//		salesInfoList = append(salesInfoList, salesInfoMap)
				//	}
				//
				//	zoneListMap["sales_info"] = salesInfoList
				//}

				zoneListList = append(zoneListList, zoneListMap)
			}

			zoneResponseMap["zone_list"] = zoneListList
		}

		if result.MaxBuyInstanceNum != nil {
			zoneResponseMap["max_buy_instance_num"] = result.MaxBuyInstanceNum
		}

		if result.MaxBandwidth != nil {
			zoneResponseMap["max_bandwidth"] = result.MaxBandwidth
		}

		if result.UnitPrice != nil {
			unitPriceMap := map[string]interface{}{}

			if result.UnitPrice.RealTotalCost != nil {
				unitPriceMap["real_total_cost"] = result.UnitPrice.RealTotalCost
			}

			if result.UnitPrice.TotalCost != nil {
				unitPriceMap["total_cost"] = result.UnitPrice.TotalCost
			}

			zoneResponseMap["unit_price"] = []interface{}{unitPriceMap}
		}

		if result.MessagePrice != nil {
			messagePriceMap := map[string]interface{}{}

			if result.MessagePrice.RealTotalCost != nil {
				messagePriceMap["real_total_cost"] = result.MessagePrice.RealTotalCost
			}

			if result.MessagePrice.TotalCost != nil {
				messagePriceMap["total_cost"] = result.MessagePrice.TotalCost
			}

			zoneResponseMap["message_price"] = []interface{}{messagePriceMap}
		}

		if result.ClusterInfo != nil {
			clusterInfoList := []interface{}{}
			for _, clusterInfo := range result.ClusterInfo {
				clusterInfoMap := map[string]interface{}{}

				if clusterInfo.ClusterId != nil {
					clusterInfoMap["cluster_id"] = clusterInfo.ClusterId
				}

				if clusterInfo.ClusterName != nil {
					clusterInfoMap["cluster_name"] = clusterInfo.ClusterName
				}

				if clusterInfo.MaxDiskSize != nil {
					clusterInfoMap["max_disk_size"] = clusterInfo.MaxDiskSize
				}

				if clusterInfo.MaxBandWidth != nil {
					clusterInfoMap["max_band_width"] = clusterInfo.MaxBandWidth
				}

				if clusterInfo.AvailableDiskSize != nil {
					clusterInfoMap["available_disk_size"] = clusterInfo.AvailableDiskSize
				}

				if clusterInfo.AvailableBandWidth != nil {
					clusterInfoMap["available_band_width"] = clusterInfo.AvailableBandWidth
				}

				if clusterInfo.ZoneId != nil {
					clusterInfoMap["zone_id"] = clusterInfo.ZoneId
				}

				if clusterInfo.ZoneIds != nil {
					clusterInfoMap["zone_ids"] = clusterInfo.ZoneIds
				}

				clusterInfoList = append(clusterInfoList, clusterInfoMap)
			}

			zoneResponseMap["cluster_info"] = clusterInfoList
		}

		//if result.Standard != nil {
		//	zoneResponseMap["standard"] = result.Standard
		//}
		//
		//if result.StandardS2 != nil {
		//	zoneResponseMap["standard_s2"] = result.StandardS2
		//}
		//
		//if result.Profession != nil {
		//	zoneResponseMap["profession"] = result.Profession
		//}
		//
		//if result.Physical != nil {
		//	zoneResponseMap["physical"] = result.Physical
		//}
		//
		//if result.PublicNetwork != nil {
		//	zoneResponseMap["public_network"] = result.PublicNetwork
		//}
		//
		//if result.PublicNetworkLimit != nil {
		//	zoneResponseMap["public_network_limit"] = result.PublicNetworkLimit
		//}
		zoneResponseMapList = append(zoneResponseMapList, zoneResponseMap)
		_ = d.Set("result", zoneResponseMapList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), zoneResponseMapList); e != nil {
			return e
		}
	}
	return nil
}

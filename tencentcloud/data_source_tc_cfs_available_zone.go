/*
Use this data source to query detailed information of cfs available_zone

# Example Usage

```hcl
data "tencentcloudenterprise_cfs_available_zone" "available_zone" {}
```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cfs "terraform-provider-tencentcloudenterprise/sdk/cfs/v20190719"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cfs_available_zone", CNDescription{
		TerraformTypeCN: "云文件存储可用区",
		DescriptionCN:   "提供CFS文件存储可用区数据源，用于查询文件存储可用区的详细信息。",
		AttributesCN: map[string]string{
			"region_zones":  "信息，如每个AZ中的资源可用性以及支持的存储类和协议",
			"region":        "地区名称，例如“北京”",
			"region_name":   "区域名称，例如“bj”",
			"region_status": "区域可用性如果一个地区至少有一个可购买资源的AZ，则该值将为AVAILABLE；否则，它将不可用",
			"zones":         "AZ阵列",
			"zone":          "AZ名称",
			"zone_id":       "AZ ID",
			"zone_cn_name":  "AZ的中文名称",
			"types":         "类的数组",
			"protocols":     "协议和销售细节",
			"sale_status":   "销售状态有效值：sale_out（已售完）、saling（可购买）、no_saling（不可购买）",
			"protocol":      "协议类型有效值：NFS、CIFS",
			"type":          "存储类别有效值：SD（标准存储）和HP（高性能存储）",
		},
	})
}

func dataSourceTencentCloudCfsAvailableZone() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of cfs available_zone",
		Read:        dataSourceTencentCloudCfsAvailableZoneRead,
		Schema: map[string]*schema.Schema{
			"region_zones": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Information such as resource availability in each AZ and the supported storage classes and protocols.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region name, such as `ap-beijing`.",
						},
						"region_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region name, such as `bj`.",
						},
						"region_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region availability. If a region has at least one AZ where resources are purchasable, this value will be AVAILABLE; otherwise, it will be UNAVAILABLE.",
						},
						"zones": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Array of AZs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "AZ name.",
									},
									"zone_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "AZ ID.",
									},
									"zone_cn_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Chinese name of an AZ.",
									},
									"types": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Array of classes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"protocols": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Protocol and sale details.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"sale_status": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "	Sale status. Valid values: sale_out (sold out), saling (purchasable), no_saling (non-purchasable).",
															},
															"protocol": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Protocol type. Valid values: NFS, CIFS.",
															},
														},
													},
												},
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Storage class. Valid values: SD (standard storage) and HP (high-performance storage).",
												},
												/*
													"prepayment": {
														Type:        schema.TypeBool,
														Computed:    true,
														Description: "Indicates whether prepaid is supported. true: yes; false: no.",
													},

												*/
											},
										},
									},
									"zone_en_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Chinese and English names of an AZ.",
									},
								},
							},
						},
						"region_cn_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region chinese name, such as `Guangzhou`.",
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

func dataSourceTencentCloudCfsAvailableZoneRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_cfs_available_zone.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CfsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var regionZones []*cfs.AvailableRegion

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCfsAvailableZoneByFilter(ctx)
		if e != nil {
			return retryError(e)
		}
		regionZones = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(regionZones))
	tmpList := make([]map[string]interface{}, 0, len(regionZones))

	for _, availableRegion := range regionZones {
		availableRegionMap := map[string]interface{}{}

		if availableRegion.Region != nil {
			availableRegionMap["region"] = availableRegion.Region
		}

		if availableRegion.RegionName != nil {
			availableRegionMap["region_name"] = availableRegion.RegionName
		}

		if availableRegion.RegionStatus != nil {
			availableRegionMap["region_status"] = availableRegion.RegionStatus
		}

		if availableRegion.Zones != nil {
			zonesList := []interface{}{}
			for _, zones := range availableRegion.Zones {
				zonesMap := map[string]interface{}{}

				if zones.Zone != nil {
					zonesMap["zone"] = zones.Zone
				}

				if zones.ZoneId != nil {
					zonesMap["zone_id"] = zones.ZoneId
				}

				if zones.ZoneCnName != nil {
					zonesMap["zone_cn_name"] = zones.ZoneCnName
				}

				if zones.Types != nil {
					typesList := []interface{}{}
					for _, types := range zones.Types {
						typesMap := map[string]interface{}{}

						if types.Protocols != nil {
							protocolsList := []interface{}{}
							for _, protocols := range types.Protocols {
								protocolsMap := map[string]interface{}{}

								if protocols.SaleStatus != nil {
									protocolsMap["sale_status"] = protocols.SaleStatus
								}

								if protocols.Protocol != nil {
									protocolsMap["protocol"] = protocols.Protocol
								}

								protocolsList = append(protocolsList, protocolsMap)
							}

							typesMap["protocols"] = protocolsList
						}

						if types.Type != nil {
							typesMap["type"] = types.Type
						}
						/*
							if types.Prepayment != nil {
								typesMap["prepayment"] = types.Prepayment
							}

						*/

						typesList = append(typesList, typesMap)
					}

					zonesMap["types"] = typesList
				}

				if zones.ZoneEnName != nil {
					zonesMap["zone_en_name"] = zones.ZoneEnName
				}

				zonesList = append(zonesList, zonesMap)
			}

			availableRegionMap["zones"] = zonesList
		}

		if availableRegion.RegionCnName != nil {
			availableRegionMap["region_cn_name"] = availableRegion.RegionCnName
		}
		ids = append(ids, *availableRegion.Region)
		tmpList = append(tmpList, availableRegionMap)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("region_zones", tmpList)
	if err != nil {
		return err
	}
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}

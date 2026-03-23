/*
Use this data source to query availability zones by product.

# Example Usage

```hcl

	data "tencentcloudenterprise_availability_zones_by_product" "example" {
	  product = "cvm"
	}

```
*/
package tencentcloud

import (
	"log"

	location "terraform-provider-tencentcloudenterprise/sdk/location/v20191128"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTencentCloudAvailabilityZonesByProduct() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudAvailabilityZonesByProductRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "When specified, only the zone with the exactly name match will be returned.",
			},
			"product": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A string variable indicates that the query will use product information.",
			},
			"include_unavailable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A bool variable indicates that the query will include `UNAVAILABLE` zones.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// Computed values.
			"zones": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of zones will be exported and its every element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An internal id for the zone, like `200003`, usually not so useful.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the zone, like `ap-guangzhou-3`.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the zone, like `Guangzhou Zone 3`.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the zone, indicate availability using `AVAILABLE` and `UNAVAILABLE` values.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudAvailabilityZonesByProductRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_availability_zones_by_product.read")()

	logId := getLogId(contextNil)

	var name string
	var product string
	var includeUnavailable = false
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("product"); ok {
		product = v.(string)
	}
	if v, ok := d.GetOkExists("include_unavailable"); ok {
		includeUnavailable = v.(bool)
	}

	request := location.NewDescribeRegionZoneRequest()
	request.ProductId = helper.String(product)

	var regions []*location.RegionEx
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseLocationClient().DescribeRegionZone(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())

		if response == nil || response.Response == nil {
			return nil
		}
		regions = response.Response.RegionSet
		return nil
	})

	if err != nil {
		return err
	}

	// 获取当前 region
	currentRegion := meta.(*TencentCloudClient).apiV3Conn.Region

	zoneList := make([]map[string]interface{}, 0)
	ids := make([]string, 0)

	for _, region := range regions {
		// 只返回当前 region 的可用区
		if region.Region == nil || *region.Region != currentRegion {
			continue
		}

		for _, zone := range region.ZoneSet {
			if zone.Zone == nil {
				continue
			}
			if name != "" && name != *zone.Zone {
				continue
			}
			if !includeUnavailable && zone.ZoneState != nil && *zone.ZoneState == ZONE_STATE_UNAVAILABLE {
				continue
			}

			zoneId := ""
			if zone.ZoneID != nil {
				zoneId = *zone.ZoneID
			}

			mapping := map[string]interface{}{
				"id":          zoneId,
				"name":        zone.Zone,
				"description": zone.ZoneName,
				"state":       zone.ZoneState,
			}
			zoneList = append(zoneList, mapping)
			ids = append(ids, zoneId)
		}
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("zones", zoneList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set zones list fail, reason:%s\n ", logId, err.Error())
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), zoneList); err != nil {
			return err
		}
	}

	return nil
}

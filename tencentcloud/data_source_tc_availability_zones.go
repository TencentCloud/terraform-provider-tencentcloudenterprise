/*
Use this data source to get the available zones in current region. By default only `AVAILABLE` zones will be returned, but `UNAVAILABLE` zones can also be fetched when `include_unavailable` is specified.

Example Usage

```hcl
data "tencentcloudenterprise_availability_zones" "my_favourite_zone" {
  name = "ap-guangzhou-3"
}
```
*/
package tencentcloud

import (
	"context"
	"log"

	cvm "terraform-provider-tencentcloudenterprise/sdk/cvm/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// 注册资源到资源中文描述提供者映射里
	registerDataDescriptionProvider("tencentcloudenterprise_cvm_instances", CNDescription{
		TerraformTypeCN: "查询可用区信息",
		AttributesCN: map[string]string{
			"id":          "可用区的内部ID，如 ‘200003’，通常不太有用",
			"name":        "可用区的名称，如 `ap-guangzhou-3`",
			"description": "可用区的描述，如 `广州三区`",
			"state":       "可用区的状态，使用 ‘AVAILABLE’ 和 ‘UNAVAILABLE’ 的值来表示可用性",
		}},
	)
}

func dataSourceTencentCloudAvailabilityZones() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to get the available zones in current region." +
			" By default only `AVAILABLE` zones will be returned, " +
			"but `UNAVAILABLE` zones can also be fetched when `include_unavailable` is specified.",
		Read: dataSourceTencentCloudAvailabilityZonesRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "When specified, only the zone with the exactly name match will be returned.",
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

func dataSourceTencentCloudAvailabilityZonesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_availability_zones.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	cvmService := CvmService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var name string
	var includeUnavailable = false
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOkExists("include_unavailable"); ok {
		includeUnavailable = v.(bool)
	}

	var zones []*cvm.ZoneInfo
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		zones, errRet = cvmService.DescribeZones(ctx)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}

	zoneList := make([]map[string]interface{}, 0, len(zones))
	ids := make([]string, 0, len(zones))
	for _, zone := range zones {
		if name != "" && name != *zone.Zone {
			continue
		}
		if !includeUnavailable && *zone.ZoneState == ZONE_STATE_UNAVAILABLE {
			continue
		}
		mapping := map[string]interface{}{
			"id":          zone.ZoneId,
			"name":        zone.Zone,
			"description": zone.ZoneName,
			"state":       zone.ZoneState,
		}
		zoneList = append(zoneList, mapping)
		ids = append(ids, *zone.ZoneId)
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

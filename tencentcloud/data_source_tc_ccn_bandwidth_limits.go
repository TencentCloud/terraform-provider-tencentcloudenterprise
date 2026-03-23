/*
Use this data source to query detailed information of CCN bandwidth limits.

Example Usage

```hcl

	variable "other_region1" {
	  default = "ap-shanghai"
	}

	resource "tencentcloudenterprise_ccn" "main" {
	  name        = "ci-temp-test-ccn"
	  description = "ci-temp-test-ccn-des"
	  qos         = "AG"
	}

	data "tencentcloudenterprise_ccn_bandwidth_limits" "limit" {
	  ccn_id = tencentcloudenterprise_ccn.main.id
	}

	resource "tencentcloudenterprise_ccn_bandwidth_limit" "limit1" {
	  ccn_id          = tencentcloudenterprise_ccn.main.id
	  region          = var.other_region1
	  bandwidth_limit = 500
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
	registerDataDescriptionProvider("tencentcloudenterprise_ccn_bandwidth_limits", CNDescription{
		TerraformTypeCN: "CCN带宽限制",
		AttributesCN: map[string]string{
			"ccn_id":             "要查询的CCN的ID",
			"result_output_file": "用于保存结果",
			"limits":             "区域的带宽限制：",
			"region":             "区域限制",
			"bandwidth_limit":    "带宽限制",
			"dst_region":         "目的地区域限制",
		},
	})
}

func dataSourceTencentCloudCcnBandwidthLimits() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of CCN bandwidth limits.",
		Read:        dataSourceTencentCloudCcnBandwidthLimitsRead,

		Schema: map[string]*schema.Schema{
			"ccn_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the CCN to be queried.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			// Computed values
			"limits": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The bandwidth limits of regions:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Limitation of region.",
						},
						"bandwidth_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Limitation of bandwidth.",
						},
						"dst_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination area restriction.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCcnBandwidthLimitsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_ccn_bandwidth_limit.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		ccnId = d.Get("ccn_id").(string)
	)

	var infos, err = service.GetCcnRegionBandwidthLimits(ctx, ccnId)
	if err != nil {
		return err
	}

	var infoList = make([]map[string]interface{}, 0, len(infos))

	for _, item := range infos {
		var infoMap = make(map[string]interface{})
		infoMap["region"] = item.Region
		infoMap["bandwidth_limit"] = item.BandwidthLimit
		infoMap["dst_region"] = item.DstRegion
		infoList = append(infoList, infoMap)
	}
	if err := d.Set("limits", infoList); err != nil {
		log.Printf("[CRITAL]%s provider set  ccn  bandwidth limits fail, reason:%s\n ", logId, err.Error())
		return err
	}

	d.SetId(ccnId)

	if output, ok := d.GetOk("result_output_file"); ok && output.(string) != "" {
		if err := writeToFile(output.(string), infoList); err != nil {
			log.Printf("[CRITAL]%s output file[%s] fail, reason[%s]\n",
				logId, output.(string), err.Error())
			return err
		}
	}
	return nil
}

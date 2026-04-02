/*
Provide a data source to query VPCDNS zone records.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpcdns_zone_records" "foo" {
	  zone_id            = "zone-xxxxxxxx"
	  result_output_file = "zone_records.json"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpcdns_zone_records", CNDescription{
		TerraformTypeCN: "VPCDNS私有域解析记录列表",
		DescriptionCN:   "提供VPCDNS私有域解析记录数据源，用于批量查询私有域的解析记录。",
		AttributesCN: map[string]string{
			"zone_id":            "私有域ID",
			"result_output_file": "数据源查询结果文件",
		},
	})
}

func dataSourceTencentCloudVpcDnsZoneRecords() *schema.Resource {
	return &schema.Resource{
		Description: "Query VPCDNS private zone record list.",
		Read:        dataSourceTencentCloudVpcDnsZoneRecordsRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Private zone ID.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The file path to output the result.",
			},
			// Computed Values
			"record_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of zone records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"record_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Record ID.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Zone ID.",
						},
						"sub_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subdomain.",
						},
						"record_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Record type: A, AAAA, CNAME, MX, TXT, PTR.",
						},
						"record_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Record value.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TTL.",
						},
						"mx": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "MX priority.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Record weight, 1-100.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Record status.",
						},
						"enabled": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Enabled: 0 paused, 1 enabled.",
						},
						"created_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"updated_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last update time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudVpcDnsZoneRecordsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_vpcdns_zone_records.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	zoneId := d.Get("zone_id").(string)

	var records []*vpcdns.PrivateZoneRecord
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcDnsZoneRecordByFilter(ctx, zoneId, "")
		if e != nil {
			return retryError(e, InternalError)
		}
		records = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(records))
	recordList := make([]map[string]interface{}, 0, len(records))

	for _, record := range records {
		recordId := ""
		if record.RecordId != nil {
			recordId = *record.RecordId
		}
		ids = append(ids, recordId)

		recordList = append(recordList, map[string]interface{}{
			"record_id":    record.RecordId,
			"zone_id":      record.ZoneId,
			"sub_domain":   record.SubDomain,
			"record_type":  record.RecordType,
			"record_value": record.RecordValue,
			"ttl":          record.TTL,
			"mx":           record.MX,
			"weight":       record.Weight,
			"status":       record.Status,
			"enabled":      record.Enabled,
			"created_on":   record.CreatedOn,
			"updated_on":   record.UpdatedOn,
		})
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("record_list", recordList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), recordList); e != nil {
			return e
		}
	}

	log.Printf("[DEBUG]%s data source tencentcloudenterprise_vpcdns_zone_records read, total %d records", logId, len(records))
	return nil
}

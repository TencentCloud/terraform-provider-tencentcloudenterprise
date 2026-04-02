/*
Provide a data source to query VPCDNS zones.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpcdns_zones" "foo" {
	  result_output_file = "zones.json"
	}

	data "tencentcloudenterprise_vpcdns_zones" "by_domain" {
	  domain = "example.com"
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpcdns_zones", CNDescription{
		TerraformTypeCN: "VPCDNS私有域列表",
		DescriptionCN:   "提供VPCDNS私有域数据源，用于批量查询私有域信息。",
		AttributesCN: map[string]string{
			"zone_id":            "私有域ID",
			"domain":             "域名",
			"result_output_file": "数据源查询结果文件",
		},
	})
}

func dataSourceTencentCloudVpcDnsZones() *schema.Resource {
	return &schema.Resource{
		Description: "Query VPCDNS private zone list.",
		Read:        dataSourceTencentCloudVpcDnsZonesRead,

		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Private zone ID to filter.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Domain name to filter.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The file path to output the result.",
			},
			// Computed Values
			"zone_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of private zones.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Private zone ID.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name.",
						},
						"domain_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Numeric domain ID.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remark.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status: SUSPEND, ENABLED, FAILED.",
						},
						"dns_forward_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DNS forward status: ENABLED, DISABLED.",
						},
						"record_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Record count.",
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
						"vpc_set": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "VPC binddings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uniq_vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "VPC ID.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Region.",
									},
								},
							},
						},
						"account_vpc_set": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Cross-account VPC bindings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uin": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Account UIN.",
									},
									"uniq_vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "VPC ID.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Region.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudVpcDnsZonesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_vpcdns_zones.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	filters := make([]*vpcdns.Filter, 0)
	if v, ok := d.GetOk("zone_id"); ok {
		filters = append(filters, &vpcdns.Filter{
			Name:   helper.String("ZoneId"),
			Values: []*string{helper.String(v.(string))},
		})
	}
	if v, ok := d.GetOk("domain"); ok {
		filters = append(filters, &vpcdns.Filter{
			Name:   helper.String("Domain"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	var zones []*vpcdns.PrivateZone
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcDnsZoneList(ctx, filters)
		if e != nil {
			return retryError(e, InternalError)
		}
		zones = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(zones))
	zoneList := make([]map[string]interface{}, 0, len(zones))

	for _, zone := range zones {
		ids = append(ids, *zone.ZoneId)

		vpcSet := make([]map[string]interface{}, 0)
		if zone.VpcSet != nil {
			for _, vpc := range zone.VpcSet {
				vpcSet = append(vpcSet, map[string]interface{}{
					"uniq_vpc_id": vpc.UniqVpcId,
					"region":      vpc.Region,
				})
			}
		}

		accountVpcSet := make([]map[string]interface{}, 0)
		if zone.AccountVpcSet != nil {
			for _, acc := range zone.AccountVpcSet {
				accountVpcSet = append(accountVpcSet, map[string]interface{}{
					"uin":         acc.Uin,
					"uniq_vpc_id": acc.UniqVpcId,
					"region":      acc.Region,
				})
			}
		}

		item := map[string]interface{}{
			"zone_id":            zone.ZoneId,
			"domain":             zone.Domain,
			"remark":             zone.Remark,
			"status":             zone.Status,
			"dns_forward_status": zone.DnsForwardStatus,
			"record_count":       zone.RecordCount,
			"created_on":         zone.CreatedOn,
			"updated_on":         zone.UpdatedOn,
			"vpc_set":            vpcSet,
			"account_vpc_set":    accountVpcSet,
		}
		if zone.DomainId != nil {
			item["domain_id"] = *zone.DomainId
		}

		zoneList = append(zoneList, item)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("zone_list", zoneList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), zoneList); e != nil {
			return e
		}
	}

	log.Printf("[DEBUG]%s data source tencentcloudenterprise_vpcdns_zones read, total %d zones", logId, len(zones))
	return nil
}

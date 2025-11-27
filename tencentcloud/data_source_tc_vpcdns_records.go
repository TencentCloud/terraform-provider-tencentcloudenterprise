/*
Provide a resource to query VPCDNS records.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpcdns_records" "foo" {
	  domain_id      = "xxx"
	}

```

# Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import cloud_vpcdns_domain.test domain_id
```
*/
package tencentcloud

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpcdns_records", CNDescription{
		TerraformTypeCN: "VPCDNS记录列表",
		DescriptionCN:   "提供VPCDNS记录数据源，用于查询VPCDNS记录的详细信息。",
		AttributesCN: map[string]string{
			"domain_id":          "域名ID",
			"record_id":          "记录ID",
			"sub_domain":         "子域名",
			"record_type":        "记录类型",
			"enabled":            "是否启用, 0: 禁用, 1: 启用",
			"status":             "状态",
			"ttl":                "TTL",
			"mx":                 "MX",
			"create_time":        "创建时间",
			"update_time":        "更新时间",
			"result_output_file": "数据源查询结果文件，白屏化界面不可用",
		},
	})
}
func dataSourceTencentCloudVpcDnsRecord() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to query VPCDNS records.",
		Read:        dataTencentCloudVpcDnsRecordRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The domain id of the VpcDns.",
			},
			"sub_domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain of the VpcDns.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The file path to output the result.",
			},
			// Computed values
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of records.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the domain.",
						},
						"record_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the record.",
						},
						"sub_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sub domain of the record.",
						},
						"record_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of record.",
						},
						"enabled": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Is enabled, 0: disabled, 1: enabled.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of record.",
						},
						"ttl": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "TTL.",
						},
						"mx": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "MX.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of record.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of record.",
						},
					},
				},
			},
		},
	}
}

func dataTencentCloudVpcDnsRecordRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentclout_vpcdns_domain.read")()
	defer inconsistentCheck(d, meta)()

	var (
		request = vpcdns.NewDescribeVpcDnsRecordListRequest()
		filters []*vpcdns.RecordListFilters
		records []*vpcdns.RecordDetail
	)
	domainId := d.Get("domain_id").(int)
	request.DomainId = helper.IntUint64(domainId)

	if v, ok := d.GetOk("sub_domain"); ok && v.(string) != "" {
		filters = append(filters, &vpcdns.RecordListFilters{
			Name: helper.String("SubDomain"),
			Values: []*string{
				helper.String(v.(string)),
			},
		})
		request.Filters = filters
	}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().DescribeVpcDnsRecordList(request)
		if e != nil {
			return retryError(e)
		}

		if response.Response == nil || len(response.Response.Records) < 1 {
			return retryError(errors.New("vpc dns records not found"))
		}
		records = response.Response.Records

		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(records))
	list := make([]map[string]interface{}, 0, len(records))

	for _, v := range records {
		ids = append(ids, fmt.Sprintf("%d", *v.DomainId))
		list = append(list, map[string]interface{}{
			"domain_id":   v.DomainId,
			"record_id":   v.RecordId,
			"sub_domain":  v.SubDomain,
			"record_type": v.RecordType,
			"enabled":     v.Enabled,
			"status":      v.Status,
			"ttl":         v.Ttl,
			"mx":          v.Mx,
			"create_time": v.CreatedOn,
			"update_time": v.UpdatedOn,
		})
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("list", list)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), list); e != nil {
			return e
		}
	}

	return nil
}

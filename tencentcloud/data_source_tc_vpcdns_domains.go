/*
Provide a resource to query VPCDNS domain.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpcdns_domains" "foo" {
	  domain      = "brucezylin.cc"
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpcdns_domains", CNDescription{
		TerraformTypeCN: "VPCDNS域名列表",
		DescriptionCN:   "提供VPCDNS域名数据源，用于查询VPCDNS域名的详细信息。",
		AttributesCN: map[string]string{
			"domain":             "域名",
			"create_time":        "创建时间",
			"update_time":        "更新时间",
			"record_count":       "记录数量",
			"remark":             "备注",
			"dns_forward_status": "DNS转发状态",
			"list":               "	域名列表",
			"result_output_file": "输出结果到文件",
			"domain_id":          "域名ID",
		},
	})
}
func dataSourceTencentCloudVpcDnsDomain() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to query VPCDNS domain.",
		Read:        dataTencentCloudVpcDnsDomainRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": {
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
				Description: "List of domain.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "ID of the domain.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Content of domain.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of domain.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of domain.",
						},
						"record_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Record count of domain.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remark of domain.",
						},
						"dns_forward_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Dns forward status of domain.",
						},
					},
				},
			},
		},
	}
}

func dataTencentCloudVpcDnsDomainRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentclout_vpcdns_domain.read")()
	defer inconsistentCheck(d, meta)()

	var (
		request = vpcdns.NewDescribeVpcDnsDomainListRequest()
		filters []*vpcdns.DomainListFilters
		domains []*vpcdns.DomainDetail
	)

	if v, ok := d.GetOk("domain"); ok && v.(string) != "" {
		filters = append(filters, &vpcdns.DomainListFilters{
			Name: helper.String("Domain"),
			Values: []*string{
				helper.String(v.(string)),
			},
		})
		request.Filters = filters
	}
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		response, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().DescribeVpcDnsDomainList(request)
		if e != nil {
			return retryError(e)
		}

		if response.Response == nil || len(response.Response.Domains) < 1 {
			return retryError(errors.New("vpc dns domain not found"))
		}
		domains = response.Response.Domains

		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(domains))
	list := make([]map[string]interface{}, 0, len(domains))

	for _, v := range domains {
		ids = append(ids, fmt.Sprintf("%d", *v.DomainId))
		list = append(list, map[string]interface{}{
			"domain_id":          v.DomainId,
			"domain":             v.Domain,
			"create_time":        v.CreatedOn,
			"update_time":        v.UpdatedOn,
			"record_count":       v.RecordCount,
			"remark":             v.Remark,
			"dns_forward_status": v.DnsForwardStatus,
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

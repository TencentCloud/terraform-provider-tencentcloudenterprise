/*
Provide a datasource to query Domains.

Example Usage

```hcl
data "tencentcloudenterprise_vpc_domains" "foo" {
}
```
*/
package tencentcloud

import (
	"context"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	domain "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/domain/v20180808"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_domains", CNDescription{
		TerraformTypeCN: "域名列表",
		AttributesCN: map[string]string{
			"offset":             "指定数据偏移量默认值：0",
			"limit":              "指定范围为[1100]的数据限制默认值：20",
			"result_output_file": "用于将响应保存为本地文件",
			"list":               "域结果列表",
			"auto_renew":         "域是否自动续订，0-手动续订，1-自动续订",
			"is_premium":         "域名是否为高级域名",
			"domain_id":          "域ID",
			"expiration_date":    "域过期日期",
			"domain_name":        "域名",
			"code_tld":           "域名有限公司",
			"creation_date":      "域创建时间",
			"tld":                "域有限公司",
			"buy_status":         "域购买状态",
		},
	})
}

func dataSourceTencentCloudDomains() *schema.Resource {
	return &schema.Resource{
		Description: "A data source to query Domains.",
		Read:        datasourceTencentCloudDomainsRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "Specify data offset. Default: 0.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     20,
				Description: "Specify data limit in range [1, 100]. Default: 20.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used for save response as file locally.",
			},
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Domain result list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_renew": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether the domain auto renew, 0 - manual renew, 1 - auto renew.",
						},
						"is_premium": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the domain is premium.",
						},
						"domain_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain ID.",
						},
						"expiration_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain expiration date.",
						},
						"domain_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name.",
						},
						"code_tld": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain code ltd.",
						},
						"creation_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain create time.",
						},
						"tld": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain ltd.",
						},
						"buy_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain buy status.",
						},
					},
				},
			},
		},
	}
}

func datasourceTencentCloudDomainsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("datasource.tencentcloudenterprise_vpc_domains.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	client := meta.(*TencentCloudClient).apiV3Conn
	service := DomainService{client}
	request := domain.NewDescribeDomainNameListRequest()

	if v, ok := d.GetOk("limit"); ok {
		request.Limit = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("offset"); ok {
		request.Offset = helper.IntUint64(v.(int))
	}

	result, err := service.DescribeDomainNameList(ctx, request)

	if err != nil {
		d.SetId("")
		return err
	}

	list := make([]interface{}, 0, len(result))
	ids := make([]string, 0)

	for i := range result {
		item := result[i]
		ids = append(ids, *item.DomainId)
		list = append(list, map[string]interface{}{
			"auto_renew":      item.AutoRenew,
			"is_premium":      item.IsPremium,
			"domain_id":       item.DomainId,
			"expiration_date": item.ExpirationDate,
			"domain_name":     item.DomainName,
			"code_tld":        item.CodeTld,
			"creation_date":   item.CreationDate,
			"tld":             item.Tld,
			"buy_status":      item.BuyStatus,
		})
	}

	d.SetId("domains-" + helper.DataResourceIdsHash(ids))
	if err := d.Set("list", list); err != nil {
		return err
	}

	if output, ok := d.GetOk("result_output_file"); ok {
		return writeToFile(output.(string), result)
	}

	return nil
}

/*
Use this data source to query detailed information of vpc address_quota

# Example Usage

```hcl
data "tencentcloudenterprise_eip_address_quota" "address_quota" {}
```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_eip_address_quota", CNDescription{
		TerraformTypeCN: "EIP地址配额",
		DescriptionCN:   "提供EIP地址配额数据源，用于查询EIP地址配额的详细信息。",
		AttributesCN: map[string]string{
			"quota_set":     "指定的帐户EIP配额信息",
			"quota_id":      "配额名称：TOTAL_EIP_Quota、DAILY_IP_APPLY、DAILY_PUBLIC_IP_ASSIGN",
			"quota_current": "当前计数",
			"quota_limit":   "配额计数",
		},
	})
}

func dataSourceTencentCloudEipAddressQuota() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc address_quota",
		Read:        dataSourceTencentCloudEipAddressQuotaRead,
		Schema: map[string]*schema.Schema{
			"quota_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "The specified account EIP quota information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quota_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Quota name: TOTAL_EIP_QUOTA,DAILY_EIP_APPLY,DAILY_PUBLIC_IP_ASSIGN.",
						},
						"quota_current": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Current count.",
						},
						"quota_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Quota count.",
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

func dataSourceTencentCloudEipAddressQuotaRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_address_quota.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var quotaSet []*vpc.Quota

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeEipAddressQuota(ctx)
		if e != nil {
			return retryError(e)
		}
		quotaSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(quotaSet))
	tmpList := make([]map[string]interface{}, 0, len(quotaSet))

	if quotaSet != nil {
		for _, quota := range quotaSet {
			quotaMap := map[string]interface{}{}

			if quota.QuotaId != nil {
				quotaMap["quota_id"] = quota.QuotaId
			}

			if quota.QuotaCurrent != nil {
				quotaMap["quota_current"] = quota.QuotaCurrent
			}

			if quota.QuotaLimit != nil {
				quotaMap["quota_limit"] = quota.QuotaLimit
			}

			ids = append(ids, *quota.QuotaId)
			tmpList = append(tmpList, quotaMap)
		}

		_ = d.Set("quota_set", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}

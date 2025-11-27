/*
Use this data source to query detailed information of vpc bandwidth_package_quota

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_bandwidth_package_quota" "bandwidth_package_quota" {
	  }

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
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_bandwidth_package_quota", CNDescription{
		TerraformTypeCN: "带宽包配额详细信息",
		AttributesCN: map[string]string{
			"quota_set":          "带宽包配额详细信息",
			"quota_id":           "配额类型",
			"quota_current":      "当前金额",
			"quota_limit":        "配额金额",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcBandwidthPackageQuota() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc bandwidth_package_quota",
		Read:        dataSourceTencentCloudVpcBandwidthPackageQuotaRead,
		Schema: map[string]*schema.Schema{
			"quota_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Bandwidth Package Quota Details.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"quota_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Quota type.",
						},
						"quota_current": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Current amount.",
						},
						"quota_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Quota amount.",
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

func dataSourceTencentCloudVpcBandwidthPackageQuotaRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_bandwidth_package_quota.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var quotaSet []*vpc.Quota

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcBandwidthPackageQuota(ctx)
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

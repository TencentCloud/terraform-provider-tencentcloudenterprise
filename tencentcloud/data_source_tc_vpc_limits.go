/*
Use this data source to query detailed information of vpc limits

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_limits" "limits" {
	  limit_types = ["appid-max-vpcs", "vpc-max-subnets"]
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_limits", CNDescription{
		TerraformTypeCN: "获取私有网络配额",
		DescriptionCN:   "提供VPC限额数据源，用于查询VPC限额的详细信息。",
		AttributesCN: map[string]string{
			"limit_types":        "配额名称每次最多可以查询100种配额类型",
			"vpc_limit_set":      "私有网络配额",
			"limit_type":         "配额类型",
			"limit_value":        "配额值",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcLimits() *schema.Resource {
	return &schema.Resource{
		Description: "Query detailed information of vpc limits",
		Read:        dataSourceTencentCloudVpcLimitsRead,
		Schema: map[string]*schema.Schema{
			"limit_types": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Quota name. A maximum of 100 quota types can be queried each time.",
			},

			"vpc_limit_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Vpc limit.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of vpc limit.",
						},
						"limit_value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Value of vpc limit.",
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

func dataSourceTencentCloudVpcLimitsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_limits.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("limit_types"); ok {
		limitTypesSet := v.(*schema.Set).List()
		paramMap["LimitTypes"] = helper.InterfacesStringsPoint(limitTypesSet)
	}

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var vpcLimitSet []*vpc.VpcLimit

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcLimitsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		vpcLimitSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(vpcLimitSet))
	tmpList := make([]map[string]interface{}, 0, len(vpcLimitSet))

	if vpcLimitSet != nil {
		for _, vpcLimit := range vpcLimitSet {
			vpcLimitMap := map[string]interface{}{}

			if vpcLimit.LimitType != nil {
				vpcLimitMap["limit_type"] = vpcLimit.LimitType
			}

			if vpcLimit.LimitValue != nil {
				vpcLimitMap["limit_value"] = vpcLimit.LimitValue
			}

			ids = append(ids, *vpcLimit.LimitType)
			tmpList = append(tmpList, vpcLimitMap)
		}

		_ = d.Set("vpc_limit_set", tmpList)
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

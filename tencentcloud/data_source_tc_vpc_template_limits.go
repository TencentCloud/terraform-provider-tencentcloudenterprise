/*
Use this data source to query detailed information of vpc template_limits

# Example Usage

```hcl
data "tencentcloudenterprise_vpc_template_limits" "template_limits" {}
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_template_limits", CNDescription{
		TerraformTypeCN: "查询参数模板配额列表",
		DescriptionCN:   "提供VPC模板限额数据源，用于查询VPC模板限额的详细信息。",
		AttributesCN: map[string]string{
			"template_limit":                      "参数模板配额列表",
			"address_template_member_limit":       "地址模板成员限制",
			"address_template_group_member_limit": "地址模板组成员限制",
			"service_template_member_limit":       "服务模板成员限制",
			"service_template_group_member_limit": "服务模板组成员限制",
			"result_output_file":                  "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcTemplateLimits() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc template_limits",
		Read:        dataSourceTencentCloudVpcTemplateLimitsRead,
		Schema: map[string]*schema.Schema{
			"template_limit": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Template limit.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address_template_member_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Address template member limit.",
						},
						"address_template_group_member_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Address template group member limit.",
						},
						"service_template_member_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Service template member limit.",
						},
						"service_template_group_member_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Service template group member limit.",
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

func dataSourceTencentCloudVpcTemplateLimitsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_template_limits.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var templateLimit *vpc.TemplateLimit

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcTemplateLimits(ctx)
		if e != nil {
			return retryError(e)
		}
		templateLimit = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0)
	templateLimitMap := map[string]interface{}{}

	if templateLimit != nil {

		if templateLimit.AddressTemplateMemberLimit != nil {
			templateLimitMap["address_template_member_limit"] = templateLimit.AddressTemplateMemberLimit
		}

		if templateLimit.AddressTemplateGroupMemberLimit != nil {
			templateLimitMap["address_template_group_member_limit"] = templateLimit.AddressTemplateGroupMemberLimit
		}

		if templateLimit.ServiceTemplateMemberLimit != nil {
			templateLimitMap["service_template_member_limit"] = templateLimit.ServiceTemplateMemberLimit
		}

		if templateLimit.ServiceTemplateGroupMemberLimit != nil {
			templateLimitMap["service_template_group_member_limit"] = templateLimit.ServiceTemplateGroupMemberLimit
		}

		ids = append(ids, helper.UInt64ToStr(*templateLimit.AddressTemplateMemberLimit))
		_ = d.Set("template_limit", []interface{}{templateLimitMap})
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), templateLimitMap); e != nil {
			return e
		}
	}
	return nil
}

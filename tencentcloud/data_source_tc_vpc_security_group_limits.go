/*
Use this data source to query detailed information of vpc security_group_limits

# Example Usage

```hcl
data "tencentcloudenterprise_vpc_security_group_limits" "security_group_limits" {}
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_security_group_limits", CNDescription{
		TerraformTypeCN: "用户安全组配额",
		DescriptionCN:   "提供VPC安全组限额数据源，用于查询VPC安全组限额的详细信息。",
		AttributesCN: map[string]string{
			"security_group_limit_set":                  "sg限制设置",
			"security_group_limit":                      "可以创建多个sg",
			"security_group_policy_limit":               "可以创建多个sg-polciy",
			"referred_security_group_limit":             "可以参考sg的数量",
			"security_group_instance_limit":             "sg关联实例的数量",
			"instance_security_group_limit":             "关联sg的实例数",
			"security_group_extended_policy_limit":      "sg扩展策略的数量",
			"security_group_referred_cvm_and_eni_limit": "可以参考eni和cvm的数量",
			"security_group_referred_svc_limit":         "可以参考svc的数量",
			"result_output_file":                        "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcSecurityGroupLimits() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc security_group_limits",
		Read:        dataSourceTencentCloudVpcSecurityGroupLimitsRead,
		Schema: map[string]*schema.Schema{
			"security_group_limit_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Sg limit set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_group_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of sg can be created.",
						},
						"security_group_policy_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of sg polciy can be created.",
						},
						"referred_security_group_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of sg can be referred.",
						},
						"security_group_instance_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of sg associated instances.",
						},
						"instance_security_group_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of instances associated sg.",
						},
						"security_group_extended_policy_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of sg extended policy.",
						},
						"security_group_referred_cvm_and_eni_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of eni and cvm can be referred.",
						},
						"security_group_referred_svc_limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of svc can be referred.",
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

func dataSourceTencentCloudVpcSecurityGroupLimitsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_security_group_limits.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var securityGroupLimitSet *vpc.SecurityGroupLimitSet

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcSecurityGroupLimits(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		securityGroupLimitSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0)
	securityGroupLimitSetMap := map[string]interface{}{}

	if securityGroupLimitSet != nil {
		if securityGroupLimitSet.SecurityGroupLimit != nil {
			securityGroupLimitSetMap["security_group_limit"] = securityGroupLimitSet.SecurityGroupLimit
		}

		if securityGroupLimitSet.SecurityGroupPolicyLimit != nil {
			securityGroupLimitSetMap["security_group_policy_limit"] = securityGroupLimitSet.SecurityGroupPolicyLimit
		}

		if securityGroupLimitSet.ReferedSecurityGroupLimit != nil {
			securityGroupLimitSetMap["referred_security_group_limit"] = securityGroupLimitSet.ReferedSecurityGroupLimit
		}

		if securityGroupLimitSet.SecurityGroupInstanceLimit != nil {
			securityGroupLimitSetMap["security_group_instance_limit"] = securityGroupLimitSet.SecurityGroupInstanceLimit
		}

		if securityGroupLimitSet.InstanceSecurityGroupLimit != nil {
			securityGroupLimitSetMap["instance_security_group_limit"] = securityGroupLimitSet.InstanceSecurityGroupLimit
		}

		if securityGroupLimitSet.SecurityGroupExtendedPolicyLimit != nil {
			securityGroupLimitSetMap["security_group_extended_policy_limit"] = securityGroupLimitSet.SecurityGroupExtendedPolicyLimit
		}

		if securityGroupLimitSet.SecurityGroupReferedCvmAndEniLimit != nil {
			securityGroupLimitSetMap["security_group_referred_cvm_and_eni_limit"] = securityGroupLimitSet.SecurityGroupReferedCvmAndEniLimit
		}

		if securityGroupLimitSet.SecurityGroupReferedSvcLimit != nil {
			securityGroupLimitSetMap["security_group_referred_svc_limit"] = securityGroupLimitSet.SecurityGroupReferedSvcLimit
		}

		ids = append(ids, helper.UInt64ToStr(*securityGroupLimitSet.SecurityGroupLimit))
		_ = d.Set("security_group_limit_set", []interface{}{securityGroupLimitSetMap})
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), securityGroupLimitSetMap); e != nil {
			return e
		}
	}
	return nil
}

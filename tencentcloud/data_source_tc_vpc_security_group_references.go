/*
Use this data source to query detailed information of vpc security_group_references

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_security_group_references" "security_group_references" {
	  security_group_ids = ["sg-edmur627"]
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_security_group_references", CNDescription{
		TerraformTypeCN: "安全组关联",
		DescriptionCN:   "提供VPC安全组引用数据源，用于查询VPC安全组引用的详细信息。",
		AttributesCN: map[string]string{
			"security_group_ids":          "一组安全组实例ID，例如[sg-12345678]",
			"referred_security_group_set": "推荐的安全组",
			"security_group_id":           "安全组实例ID",
			"referred_security_group_ids": "所有引用的安全组实例的ID",
			"result_output_file":          "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcSecurityGroupReferences() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc security_group_references",
		Read:        dataSourceTencentCloudVpcSecurityGroupReferencesRead,
		Schema: map[string]*schema.Schema{
			"security_group_ids": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A set of security group instance IDs, e.g. [sg-12345678].",
			},

			"referred_security_group_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Referred security groups.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Security group instance ID.",
						},
						"referred_security_group_ids": {
							Type: schema.TypeSet,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Computed:    true,
							Description: "IDs of all referred security group instances.",
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

func dataSourceTencentCloudVpcSecurityGroupReferencesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_security_group_references.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("security_group_ids"); ok {
		securityGroupIdsSet := v.(*schema.Set).List()
		paramMap["SecurityGroupIds"] = helper.InterfacesStringsPoint(securityGroupIdsSet)
	}

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var referredSecurityGroupSet []*vpc.ReferredSecurityGroup

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcSecurityGroupReferences(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		referredSecurityGroupSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(referredSecurityGroupSet))
	tmpList := make([]map[string]interface{}, 0, len(referredSecurityGroupSet))

	if referredSecurityGroupSet != nil {
		for _, referredSecurityGroup := range referredSecurityGroupSet {
			referredSecurityGroupMap := map[string]interface{}{}

			if referredSecurityGroup.SecurityGroupId != nil {
				referredSecurityGroupMap["security_group_id"] = referredSecurityGroup.SecurityGroupId
			}

			if referredSecurityGroup.ReferredSecurityGroupIds != nil {
				referredSecurityGroupMap["referred_security_group_ids"] = referredSecurityGroup.ReferredSecurityGroupIds
			}

			ids = append(ids, *referredSecurityGroup.SecurityGroupId)
			tmpList = append(tmpList, referredSecurityGroupMap)
		}

		_ = d.Set("referred_security_group_set", tmpList)
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

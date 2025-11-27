/*
Provide a resource to query VPCDNS domain.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpcdns_forward_rules" "foo" {
	  result_output_file = "forward_rules.json"
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
	registerDataDescriptionProvider("tencentcloudenterprise_vpcdns_forward_rules", CNDescription{
		TerraformTypeCN: "VPCDNS转发规则",
		DescriptionCN:   "提供VPCDNS转发规则数据源，用于查询VPCDNS转发规则的详细信息。",
		AttributesCN: map[string]string{
			"rule_id":            "转发规则id",
			"remark":             "转发规则名称",
			"domain_id":          "转发域名id",
			"domain_name":        "转发域名",
			"forward_address":    "dns地址",
			"vpc_infos":          "vpc关联信息",
			"created_on":         "创建时间",
			"updated_on":         "更新时间",
			"result_output_file": "数据源查询结果文件，白屏化界面不可用",
		},
	})
}

func dataSourceTencentCloudVpcDnsForwardRules() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to query VPCDNS forward rules.",
		Read:        dataTencentCloudVpcDnsForwardRuleRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The remark of the forward rule.",
			},
			"domain_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain IDs of the forward rule.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The file path to output the result.",
			},
			// Computed Values
			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of forward rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rule ID of the forward rule.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The remark of the forward rule.",
						},
						"domain_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain ID of the forward rule.",
						},
						"domain_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain name of the forward rule.",
						},
						"forward_address": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The forward address of the forward rule.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vpc_infos": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The vpc infos of the forward rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"region_id": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"un_vpc_id": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"created_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of VpcDns Forward Rule.",
						},
						"updated_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of VpcDns Forward Rule.",
						},
					},
				},
			},
		},
	}
}

func dataTencentCloudVpcDnsForwardRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentclout_vpcdns_domain.read")()
	defer inconsistentCheck(d, meta)()

	var (
		request  = vpcdns.NewDescribeVpcDnsForwardRuleRequest()
		ruleList []*vpcdns.VpcDnsForwardRuleDetail
		offset   = 0
		limit    = 20
	)

	for {
		request.Offset = helper.IntUint64(offset)
		request.Limit = helper.IntUint64(limit)

		var pageRules []*vpcdns.VpcDnsForwardRuleDetail
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			response, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().
				DescribeVpcDnsForwardRule(request)
			if e != nil {
				return retryError(e)
			}
			if response.Response == nil || len(response.Response.ForwardRuleList) < 1 {
				return retryError(errors.New("vpc dns domain not found"))
			}
			pageRules = response.Response.ForwardRuleList
			return nil
		})
		if err != nil {
			return err
		}
		ruleList = append(ruleList, pageRules...)
		if (len(pageRules)) < limit {
			break
		}
		offset += limit
	}

	ids := make([]string, 0, len(ruleList))
	list := make([]map[string]interface{}, 0, len(ruleList))

	for _, v := range ruleList {
		ids = append(ids, fmt.Sprintf("%s", *v.RuleId))

		vpcInfos := make([]map[string]interface{}, 0, len(v.VpcInfos))
		for _, vpc := range v.VpcInfos {
			vpcInfos = append(vpcInfos, map[string]interface{}{
				"vpc_id":    vpc.VpcId,
				"region_id": vpc.RegionId,
				"un_vpc_id": vpc.UnVpcId,
			})
		}

		list = append(list, map[string]interface{}{
			"rule_id":         v.RuleId,
			"remark":          v.Remark,
			"domain_id":       v.DomainId,
			"domain_name":     v.DomainName,
			"forward_address": v.ForwardAddress,
			"vpc_infos":       vpcInfos,
			"created_on":      v.CreatedOn,
			"updated_on":      v.UpdatedOn,
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

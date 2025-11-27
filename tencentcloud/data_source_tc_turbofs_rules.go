/*
Use this data source to query the detail information of Turbofs permission rule.

# Example Usage

```hcl

	data "tencentcloudenterprise_turbofs_rules" "rules" {
	  p_group_id = "pgroup-7nx89k7l"
	  rule_id  = "rule-qcndbqzj"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"
	turbofs "terraform-provider-tencentcloudenterprise/sdk/turbofs/v20190719"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_turbofs_rules", CNDescription{
		TerraformTypeCN: "权限组规则",
		DescriptionCN:   "提供TurboFS权限组规则数据源，用于查询TurboFS权限组规则的详细信息。",
		AttributesCN: map[string]string{
			"p_group_id":         "用于查询的指定权限组ID",
			"rule_id":            "用于查询的指定权限组规则ID",
			"result_output_file": "用于保存结果，可视化界面不可用",
			"rule_list":          "TurboFs权限组规则的信息列表",
			"auth_client_ip":     "允许访问的客户端IP地址或地址段",
			"rw_permission":      "读取和写入权限",
			"user_permission":    "权限组用户的权限",
			"priority":           "权限组规则的优先级",
		},
	})
}

func dataSourceTencentCloudTurbofsRules() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides the detail information of Turbofs permission rule.",
		Read:        dataSourceTencentCloudTurbofsAccessRulesRead,

		Schema: map[string]*schema.Schema{
			"p_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A specified permission group ID used to query.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A specified permission rule ID used to query.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"rule_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "An information list of Turbofs permission rule. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the permission rule.",
						},
						"auth_client_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Allowed IP of the permission rule.",
						},
						"rw_permission": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Read and write permissions.",
						},
						"user_permission": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The permissions of users.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority level of rule.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudTurbofsAccessRulesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_turbofs_rules.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	turbofsService := TurbofsService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	var ruleId string
	pGroupId := d.Get("p_group_id").(string)
	if v, ok := d.GetOk("rule_id"); ok {
		ruleId = v.(string)
	}

	var rules []*turbofs.PGroupRuleInfo
	var errRet error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		rules, errRet = turbofsService.DescribeCfsRule(ctx, pGroupId, ruleId)
		if errRet != nil {
			return retryError(errRet, InternalError)
		}
		return nil
	})
	if err != nil {
		return err
	}

	ruleList := make([]map[string]interface{}, 0, len(rules))
	ids := make([]string, 0, len(rules))
	for _, rule := range rules {
		mapping := map[string]interface{}{
			"rule_id":         rule.RuleId,
			"auth_client_ip":  rule.AuthClientIp,
			"rw_permission":   rule.RWPermission,
			"user_permission": rule.UserPermission,
			"priority":        rule.Priority,
		}
		ruleList = append(ruleList, mapping)
		ids = append(ids, *rule.RuleId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("rule_list", ruleList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set turbofs rule list fail, reason:%s\n ", logId, err.Error())
		return err
	}
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), ruleList); err != nil {
			return err
		}
	}
	return nil
}

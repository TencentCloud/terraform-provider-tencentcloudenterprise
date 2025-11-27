/*
Use this data source to query detailed information of brc auto backup policies

# Example Usage

```hcl

	data "tencentcloudenterprise_brc_autobackup_policies" "instance_policies" {
	  resource_type = "INSTANCE"
	  result_output_file = "cvm_policies.json"
	}

```
*/
package tencentcloud

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_brc_autobackup_policies", CNDescription{
		TerraformTypeCN: "BRC自动备份策略列表",
		DescriptionCN:   "提供BRC自动备份策略列表数据源，用于查询自动备份策略的详细信息。",
		AttributesCN: map[string]string{
			"resource_type":             "备份资源类型，有效值：INSTANCE(CVM实例)、DISK(CBS云硬盘)、CFS、COS、CSP、MySQL_MariaDB和TDSQL_MySQL",
			"auto_backup_policy_name":   "自动备份策略名称过滤器",
			"auto_backup_policy_id":     "自动备份策略ID过滤器",
			"is_activated":              "按激活状态过滤",
			"result_output_file":        "用于保存结果",
			"auto_backup_policy_list":   "自动备份策略列表",
			"create_time":               "创建时间",
			"policy":                    "定期备份的执行策略。",
			"instance_id_set":           "绑定的实例ID列表",
			"next_trigger_time":         "首次开始备份的时间。",
			"retention_amount":          "通过该定期备份策略最多保留的备份个数，超过该个数限制后自动删除最先创建的备份，该参数不可与IsPermanent参数冲突。",
			"is_permanent":              "通过该定期备份策略创建的备份是否永久保留。false表示非永久保留，true表示永久保留，默认为false。",
			"account_name":              "账户名称",
			"advanced_retention_policy": "高级保留策略",
			"auto_backup_policy_state":  "自动备份策略状态",
			"full_backup_interval":      "每隔几个备份做一个全量备份，0表示全部做全量备份。",
		},
	})
}

func dataSourceTencentCloudBrcAutoBackupPolicies() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudBrcAutoBackupPoliciesRead,
		Description: "Use this data source to query detailed information of brc auto backup policies",
		Schema: map[string]*schema.Schema{
			"resource_type": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Resource type filter. Valid values: INSTANCE, DISK, CFS, COS, CSP, MySQL_MariaDB, TDSQL_MySQL.",
				//ValidateFunc: validateAllowedStringValue(BackupResouceTypes),
			},
			"auto_backup_policy_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Auto backup policy name filter.",
			},
			"auto_backup_policy_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Auto backup policy ID filter.",
			},
			"is_activated": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Filter by activation status.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"auto_backup_policy_list": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Auto backup policy list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_backup_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auto backup policy ID.",
						},
						"auto_backup_policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auto backup policy name.",
						},
						"resource_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Resource type.",
						},
						"auto_backup_policy_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Auto backup policy state.",
						},
						"is_activated": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the auto backup policy is activated.",
						},
						"is_permanent": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the backup is permanent.",
						},
						"full_backup_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Full backup interval.",
						},
						"retention_amount": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Retention amount.",
						},
						"next_trigger_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Next trigger time.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account name.",
						},
						"policy": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Backup policy configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hour": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Backup hours (0-23).",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"interval_days": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Backup interval in days.",
									},
								},
							},
						},
						"advanced_retention_policy": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Advanced retention policy configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Retention days.",
									},
									"weeks": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Retention weeks.",
									},
									"months": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Retention months.",
									},
									"years": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Retention years.",
									},
								},
							},
						},
						"instance_id_set": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Bound instance IDs.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudBrcAutoBackupPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_brc_autobackup_policies.read")()

	var (
		logId      = getLogId(contextNil)
		brcService = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
	)

	request := brc.NewDescribeAutoBackupPoliciesRequest()
	request.Limit = helper.Uint64(100)
	request.Offset = helper.Uint64(0)

	// 添加过滤条件
	filters := make([]*brc.Filter, 0)

	if v, ok := d.GetOk("resource_type"); ok {
		request.ResourceType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("auto_backup_policy_name"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("auto-backup-policy-name"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if v, ok := d.GetOk("auto_backup_policy_id"); ok {
		filters = append(filters, &brc.Filter{
			Name:   helper.String("auto-backup-policy-id"),
			Values: []*string{helper.String(v.(string))},
		})
	}

	if len(filters) > 0 {
		request.Filters = filters
	}

	ratelimit.Check(request.GetAction())
	response, err := brcService.client.UseBrcClient().DescribeAutoBackupPolicies(request)
	if err != nil {
		log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), err.Error())
		return err
	}

	policyList := make([]map[string]interface{}, 0)
	if response.Response != nil && response.Response.AutoBackupPolicySet != nil {
		for _, policy := range response.Response.AutoBackupPolicySet {
			// 过滤激活状态
			if v, ok := d.GetOk("is_activated"); ok {
				if policy.IsActivated == nil || *policy.IsActivated != v.(bool) {
					continue
				}
			}

			policyMap := map[string]interface{}{
				"auto_backup_policy_id":    policy.AutoBackupPolicyId,
				"auto_backup_policy_name":  policy.AutoBackupPolicyName,
				"resource_type":            policy.ResourceType,
				"auto_backup_policy_state": policy.AutoBackupPolicyState,
				"is_activated":             policy.IsActivated,
				"is_permanent":             policy.IsPermanent,
				"full_backup_interval":     policy.FullBackupInterval,
				"retention_amount":         policy.RetentionAmount,
				"next_trigger_time":        policy.NextTriggerTime,
				"create_time":              policy.CreateTime,
				"account_name":             policy.AccountName,
			}

			// 设置policy
			if policy.Policy != nil {
				policyItems := make([]map[string]interface{}, 0, len(policy.Policy))
				for _, item := range policy.Policy {
					policyItem := map[string]interface{}{
						"interval_days": item.IntervalDays,
					}
					if item.Hour != nil {
						hours := make([]int, 0, len(item.Hour))
						for _, hour := range item.Hour {
							hours = append(hours, int(*hour))
						}
						policyItem["hour"] = hours
					}
					policyItems = append(policyItems, policyItem)
				}
				policyMap["policy"] = policyItems
			}

			// 设置advanced_retention_policy
			if policy.AdvancedRetentionPolicy != nil {
				advancedPolicy := map[string]interface{}{
					"days":   policy.AdvancedRetentionPolicy.Days,
					"weeks":  policy.AdvancedRetentionPolicy.Weeks,
					"months": policy.AdvancedRetentionPolicy.Months,
					"years":  policy.AdvancedRetentionPolicy.Years,
				}
				policyMap["advanced_retention_policy"] = []map[string]interface{}{advancedPolicy}
			}

			// 设置instance_id_set
			if policy.InstanceIdSet != nil {
				instanceIds := make([]string, 0, len(policy.InstanceIdSet))
				for _, instanceId := range policy.InstanceIdSet {
					if instanceId != nil {
						instanceIds = append(instanceIds, *instanceId)
					}
				}
				policyMap["instance_id_set"] = instanceIds
			}

			policyList = append(policyList, policyMap)
		}
	}

	// 创建ID列表用于hash
	policyIds := make([]string, 0, len(policyList))
	for _, policy := range policyList {
		if id, ok := policy["auto_backup_policy_id"].(*string); ok && id != nil {
			policyIds = append(policyIds, *id)
		}
	}

	d.SetId(fmt.Sprintf("brc_autobackup_policies_%s", helper.DataResourceIdsHash(policyIds)))
	_ = d.Set("auto_backup_policy_list", policyList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), policyList); err != nil {
			return err
		}
	}

	return nil
}

/*
Provides a resource to create a brc auto_backup_policy

# Example Usage

```hcl

	resource "tencentcloudenterprise_brc_auto_backup_policy" "example" {
	  resource_type = "INSTANCE"
	  auto_backup_policy_name = "example-policy"

	  policy {
	    hour = [0, 11, 12,13,14]
	    interval_days = 12
	  }

	  is_permanent = false
	  full_backup_interval = 2
	  retention_amount = 5

	  advanced_retention_policy {
	    days = 1
	    weeks = 1
	    months = 1
	    years = 1
	  }

	  is_activated = false
	  dry_run = false
	}

```

# Import

brc auto_backup_policy can be imported using the id, e.g.

```
terraform import cloud_brc auto_backup_policy.example auto_backup_policy_id
```
*/
package tencentcloud

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	brc "terraform-provider-tencentcloudenterprise/sdk/brc/v20220516"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_brc_auto_backup_policy", CNDescription{
		TerraformTypeCN: "自动备份策略",
		DescriptionCN:   "提供BRC自动备份策略资源，用于创建和管理自动备份策略。",
		AttributesCN: map[string]string{
			"resource_type":             "要创建定期策略的产品类型。取值范围：DISK：云硬盘,CFS：文件存储,CVM：云服务器,TDSQL_MySQL：TDSQL_MySQL版,MySQL_MariaDB：关系型数据库(MySQL_MariaDB),COS：对象存储。",
			"auto_backup_policy_name":   "定期备份策略的名称。",
			"policy":                    "定期备份的执行策略。",
			"is_permanent":              "通过该定期备份策略创建的备份是否永久保留。false表示非永久保留，true表示永久保留，默认为false。",
			"full_backup_interval":      "每隔几个备份做一个全量备份，0表示全部做全量备份。",
			"retention_amount":          "通过该定期备份策略最多保留的备份个数，超过该个数限制后自动删除最先创建的备份，该参数不可与IsPermanent参数冲突。",
			"advanced_retention_policy": "高级保留策略",
			"is_activated":              "是否激活定期备份策略。",
			"dry_run":                   "是否实际创建定期备份策略。true表示只需获取首次开始备份的时间，不实际创建定期备份策略，false表示创建，默认为false。",
			"auto_backup_policy_id":     "自动备份策略ID",
			"create_speed":              "创建备份的带宽上限，范围:[0, 100]",
			"retention_days":            "通过定期备份策略创建出的备份保留时间。",
			"retention_months":          "该定期备份策略创建的备份可以保留的月数，该参数不可与IsPermanent/RetentionDays参数冲突。",
			"need_archive":              "标识该策略创建的备份均将用于归档。",
			"day_of_month":              "指定每月从月初到月底需要触发定期备份的日期,取值范围：[1, 31]，1-31分别表示每月的具体日期，比如5表示每月的5号。注：若设置29、30、31等部分月份不存在的日期，则对应不存在日期的月份会跳过不打定期备份。",
			"day_of_week":               "选定周一到周日中需要创建快照的日期，取值范围：[0, 6]。0表示周日触发，1表示周一触发，依次类推。",
			"hour":                      "指定定期快照策略的触发时间。单位为小时，取值范围：[0, 23]。00:00 ~ 23:00 共 24 个时间点可选，1表示 01:00，依此类推。",
			"next_trigger":              "首次开始备份的时间。",
			"interval_days":             "指定创建定期快照的间隔天数，默认为1，取值范围：[1, 365]，例如设置为5，则间隔5天即触发定期快照创建。注：当选择按天备份时，理论上第一次备份的时间为备份策略创建当天。如果当天备份策略创建的时间已经晚于设置的备份时间，那么将会等到第二个备份周期再进行第一次备份。",
			"next_trigger_time":         "首次开始备份的时间。",
		},
	})
}

func resourceTencentCloudBrcAutoBackupPolicy() *schema.Resource {
	return &schema.Resource{
		Create:      resourceTencentCloudBrcAutoBackupPolicyCreate,
		Read:        resourceTencentCloudBrcAutoBackupPolicyRead,
		Update:      resourceTencentCloudBrcAutoBackupPolicyUpdate,
		Delete:      resourceTencentCloudBrcAutoBackupPolicyDelete,
		Description: "Provides a resource to create and manage BRC auto backup policy",

		Schema: map[string]*schema.Schema{
			"resource_type": {
				Required:     true,
				ForceNew:     true,
				Type:         schema.TypeString,
				Description:  "Product type for creating scheduled policy. Valid values: DISK (Cloud Block Storage), CFS (Cloud File Storage), CVM (Cloud Virtual Machine), TDSQL_MySQL (TDSQL MySQL Edition), MySQL_MariaDB (Relational Database MySQL_MariaDB), COS (Cloud Object Storage).",
				ValidateFunc: validateAllowedStringValue(BackupResouceTypes),
			},
			"auto_backup_policy_name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Name of the scheduled backup policy.",
			},
			"policy": {
				Required:    true,
				MaxItems:    1,
				Type:        schema.TypeList,
				Description: "Execution policy for scheduled backup.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hour": {
							Required:    true,
							Type:        schema.TypeList,
							Description: "Trigger time for scheduled snapshot policy. Unit: hour, value range: [0, 23]. 00:00 ~ 23:00, 24 time points available, 1 means 01:00, and so on.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"interval_days": {
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Backup interval in days.",
						},
						"day_of_week": {
							Optional:    true,
							Type:        schema.TypeList,
							Description: "Selected dates from Monday to Sunday for creating snapshots, value range: [0, 6]. 0 means Sunday trigger, 1 means Monday trigger, and so on.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"day_of_month": {
							Optional:    true,
							Type:        schema.TypeList,
							Description: "Specify dates from the beginning to the end of each month that need to trigger scheduled backup, value range: [1, 31]. 1-31 represent specific dates of each month, for example, 5 represents the 5th of each month. Note: If dates like 29, 30, 31 that don't exist in some months are set, the corresponding months without these dates will skip scheduled backup.",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
			"is_permanent": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether backups created through this scheduled backup policy are permanently retained. false means not permanently retained, true means permanently retained, default is false.",
			},
			"full_backup_interval": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "How many backups to do a full backup, 0 means all full backups.",
			},
			"retention_amount": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Maximum number of backups retained through this scheduled backup policy. When exceeding this limit, the earliest created backup will be automatically deleted. This parameter cannot conflict with the IsPermanent parameter.",
			},
			"advanced_retention_policy": {
				Optional:    true,
				Type:        schema.TypeList,
				MaxItems:    1,
				Description: "Advanced retention policy.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"days": {
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Retention days.",
						},
						"weeks": {
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Retention weeks.",
						},
						"months": {
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Retention months.",
						},
						"years": {
							Optional:    true,
							Type:        schema.TypeInt,
							Description: "Retention years.",
						},
					},
				},
			},
			"is_activated": {
				Optional:    true,
				Default:     true,
				Type:        schema.TypeBool,
				Description: "Whether to activate the scheduled backup policy.",
			},
			"dry_run": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Whether to actually create the scheduled backup policy. true means only get the first backup start time without actually creating the scheduled backup policy, false means create, default is false.",
			},
			"need_archive": {
				Optional:    true,
				Type:        schema.TypeBool,
				Description: "Indicates that backups created by this policy will all be used for archiving.",
			},
			"create_speed": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Bandwidth limit for backup creation, range: [0, 100].",
			},
			"auto_backup_policy_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Auto backup policy ID.",
			},
			"next_trigger_time": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "First backup start time.",
			},
		},
	}
}

func resourceTencentCloudBrcAutoBackupPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_auto_backup_policy.create")()

	var (
		logId        = getLogId(contextNil)
		ctx          = context.WithValue(context.TODO(), logIdKey, logId)
		brcService   = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
		resourceType = d.Get("resource_type").(string)
		policyName   = d.Get("auto_backup_policy_name").(string)
		policyItem   *brc.PolicyItem
	)

	if v, ok := d.GetOk("policy"); ok {
		policyList := make([]*brc.PolicyItem, 0, 10)
		for _, item := range v.([]interface{}) {
			policyMap := item.(map[string]interface{})
			policy := brc.PolicyItem{}

			if hours, ok := policyMap["hour"]; ok {
				hourList := hours.([]interface{})
				hourInts := make([]*uint64, 0, len(hourList))
				for _, h := range hourList {
					hourInts = append(hourInts, helper.Uint64(uint64(h.(int))))
				}
				policy.Hour = hourInts
			}

			if intervalDays, ok := policyMap["interval_days"]; ok {
				policy.IntervalDays = helper.Uint64(uint64(intervalDays.(int)))
			}

			if dayOfWeek, ok := policyMap["day_of_week"]; ok {
				dayOfWeekList := dayOfWeek.([]interface{})
				dayOfWeekInts := make([]*uint64, 0, len(dayOfWeekList))
				for _, d := range dayOfWeekList {
					dayOfWeekInts = append(dayOfWeekInts, helper.Uint64(uint64(d.(int))))
				}
				policy.DayOfWeek = dayOfWeekInts
			}

			if dayOfMonth, ok := policyMap["day_of_month"]; ok {
				dayOfMonthList := dayOfMonth.([]interface{})
				dayOfMonthInts := make([]*uint64, 0, len(dayOfMonthList))
				for _, d := range dayOfMonthList {
					dayOfMonthInts = append(dayOfMonthInts, helper.Uint64(uint64(d.(int))))
				}
				policy.DayOfMonth = dayOfMonthInts
			}
			policyList = append(policyList, &policy)
		}
		policyItem = policyList[0]
	}

	// 处理 advanced_retention_policy 字段
	var advancedRetentionPolicy *brc.AdvancedRetentionPolicy
	if v, ok := d.GetOk("advanced_retention_policy"); ok {
		advancedPolicyList := v.([]interface{})
		if len(advancedPolicyList) > 0 {
			advancedPolicyMap := advancedPolicyList[0].(map[string]interface{})
			advancedRetentionPolicy = &brc.AdvancedRetentionPolicy{}

			if days, ok := advancedPolicyMap["days"]; ok {
				advancedRetentionPolicy.Days = helper.Uint64(uint64(days.(int)))
			}
			if weeks, ok := advancedPolicyMap["weeks"]; ok {
				advancedRetentionPolicy.Weeks = helper.Uint64(uint64(weeks.(int)))
			}
			if months, ok := advancedPolicyMap["months"]; ok {
				advancedRetentionPolicy.Months = helper.Uint64(uint64(months.(int)))
			}
			if years, ok := advancedPolicyMap["years"]; ok {
				advancedRetentionPolicy.Years = helper.Uint64(uint64(years.(int)))
			}
		}
	}

	// 处理其他参数
	isPermanent := d.Get("is_permanent").(bool)
	fullBackupInterval := int64(d.Get("full_backup_interval").(int))

	var retentionAmount uint64
	if !isPermanent {
		retentionAmount = uint64(d.Get("retention_amount").(int))
	}

	dryRun := d.Get("dry_run").(bool)

	var autoBackupPolicyId string
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := brcService.CreateAutoBackupPolicy(ctx, policyItem, isPermanent, policyName,
			fullBackupInterval, retentionAmount, advancedRetentionPolicy, dryRun, resourceType)
		if e != nil {
			return retryError(e)
		}
		autoBackupPolicyId = result
		return nil
	})
	if err != nil {
		return err
	}

	d.SetId(autoBackupPolicyId)

	return resourceTencentCloudBrcAutoBackupPolicyRead(d, meta)
}

func resourceTencentCloudBrcAutoBackupPolicyRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_auto_backup_policy.read")()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		brcService = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
		policyId   = helper.String(d.Id())
	)

	var policy *brc.AutoBackupPolicy
	var e error
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		policy, e = brcService.DescribeAutoBackupPolicyById(ctx, policyId)
		if e != nil {
			return retryError(e)
		}
		return nil
	})
	if err != nil {
		return err
	}

	if policy == nil {
		d.SetId("")
		return nil
	}

	if policy.AutoBackupPolicyId != nil {
		_ = d.Set("auto_backup_policy_id", *policy.AutoBackupPolicyId)
	}
	if policy.AutoBackupPolicyName != nil {
		_ = d.Set("auto_backup_policy_name", *policy.AutoBackupPolicyName)
	}
	if policy.ResourceType != nil {
		_ = d.Set("resource_type", *policy.ResourceType)
	}
	if policy.IsPermanent != nil {
		_ = d.Set("is_permanent", *policy.IsPermanent)
	}
	if policy.FullBackupInterval != nil {
		_ = d.Set("full_backup_interval", int(*policy.FullBackupInterval))
	}
	if policy.RetentionAmount != nil {
		_ = d.Set("retention_amount", int(*policy.RetentionAmount))
	}
	if policy.IsActivated != nil {
		_ = d.Set("is_activated", *policy.IsActivated)
	}
	if policy.NextTriggerTime != nil {
		_ = d.Set("next_trigger_time", *policy.NextTriggerTime)
	}
	// 添加 Terraform schema 中定义但之前缺失的字段
	if policy.CreateSpeed != nil {
		_ = d.Set("create_speed", int(*policy.CreateSpeed))
	}
	if policy.IsArchive != nil {
		_ = d.Set("need_archive", *policy.IsArchive)
	}

	if policy.Policy != nil {
		policyList := make([]map[string]interface{}, 0, 1)

		var itemsToProcess []*brc.PolicyItem
		if len(policy.Policy) > 0 {
			itemsToProcess = policy.Policy
		}

		for _, item := range itemsToProcess {
			policyMap := make(map[string]interface{})

			if item.Hour != nil && len(item.Hour) > 0 {
				hours := make([]int, 0, len(item.Hour))
				for _, hour := range item.Hour {
					if hour != nil {
						hours = append(hours, int(*hour))
					}
				}
				policyMap["hour"] = hours
			}

			if item.IntervalDays != nil {
				policyMap["interval_days"] = int(*item.IntervalDays)
			}

			if item.DayOfWeek != nil && len(item.DayOfWeek) > 0 {
				dayOfWeeks := make([]int, 0, len(item.DayOfWeek))
				for _, dayOfWeek := range item.DayOfWeek {
					if dayOfWeek != nil {
						dayOfWeeks = append(dayOfWeeks, int(*dayOfWeek))
					}
				}
				policyMap["day_of_week"] = dayOfWeeks
			}

			if item.DayOfMonth != nil && len(item.DayOfMonth) > 0 {
				dayOfMonths := make([]int, 0, len(item.DayOfMonth))
				for _, dayOfMonth := range item.DayOfMonth {
					if dayOfMonth != nil {
						dayOfMonths = append(dayOfMonths, int(*dayOfMonth))
					}
				}
				policyMap["day_of_month"] = dayOfMonths
			}

			policyList = append(policyList, policyMap)
		}
		_ = d.Set("policy", policyList)
	}

	// 处理 advanced_retention_policy 复杂字段
	if policy.AdvancedRetentionPolicy != nil {
		advancedPolicyMap := make(map[string]interface{})

		if policy.AdvancedRetentionPolicy.Days != nil {
			advancedPolicyMap["days"] = int(*policy.AdvancedRetentionPolicy.Days)
		}
		if policy.AdvancedRetentionPolicy.Weeks != nil {
			advancedPolicyMap["weeks"] = int(*policy.AdvancedRetentionPolicy.Weeks)
		}
		if policy.AdvancedRetentionPolicy.Months != nil {
			advancedPolicyMap["months"] = int(*policy.AdvancedRetentionPolicy.Months)
		}
		if policy.AdvancedRetentionPolicy.Years != nil {
			advancedPolicyMap["years"] = int(*policy.AdvancedRetentionPolicy.Years)
		}

		_ = d.Set("advanced_retention_policy", []map[string]interface{}{advancedPolicyMap})
	}

	return nil
}

func resourceTencentCloudBrcAutoBackupPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_auto_backup_policy.update")()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		brcService = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
		policyId   = d.Id()
	)

	hasChanges := false
	updateRequest := &brc.ModifyAutoBackupPolicyAttributeRequest{
		AutoBackupPolicyId: helper.String(policyId),
	}

	if d.HasChange("policy") {
		hasChanges = true
		if v, ok := d.GetOk("policy"); ok {
			policyList := v.([]interface{})
			if len(policyList) > 0 && policyList[0] != nil {
				policyMap := policyList[0].(map[string]interface{})
				policyItem := &brc.PolicyItem{}

				if hours, ok := policyMap["hour"]; ok {
					hourList := hours.([]interface{})
					hourInts := make([]*uint64, 0, len(hourList))
					for _, h := range hourList {
						hourInts = append(hourInts, helper.Uint64(uint64(h.(int))))
					}
					policyItem.Hour = hourInts
				}

				if intervalDays, ok := policyMap["interval_days"]; ok {
					policyItem.IntervalDays = helper.Uint64(uint64(intervalDays.(int)))
				}

				if dayOfWeek, ok := policyMap["day_of_week"]; ok {
					dayOfWeekList := dayOfWeek.([]interface{})
					dayOfWeekInts := make([]*uint64, 0, len(dayOfWeekList))
					for _, d := range dayOfWeekList {
						dayOfWeekInts = append(dayOfWeekInts, helper.Uint64(uint64(d.(int))))
					}
					policyItem.DayOfWeek = dayOfWeekInts
				}

				if dayOfMonth, ok := policyMap["day_of_month"]; ok {
					dayOfMonthList := dayOfMonth.([]interface{})
					dayOfMonthInts := make([]*uint64, 0, len(dayOfMonthList))
					for _, d := range dayOfMonthList {
						dayOfMonthInts = append(dayOfMonthInts, helper.Uint64(uint64(d.(int))))
					}
					policyItem.DayOfMonth = dayOfMonthInts
				}
				updateRequest.Policy = append(updateRequest.Policy, policyItem)
			}
		}
	}

	// 检查 advanced_retention_policy 字段变化
	if d.HasChange("advanced_retention_policy") {
		hasChanges = true
		if v, ok := d.GetOk("advanced_retention_policy"); ok {
			advancedPolicyList := v.([]interface{})
			if len(advancedPolicyList) > 0 {
				advancedPolicyMap := advancedPolicyList[0].(map[string]interface{})
				advancedRetentionPolicy := &brc.AdvancedRetentionPolicy{}

				if days, ok := advancedPolicyMap["days"]; ok {
					advancedRetentionPolicy.Days = helper.Uint64(uint64(days.(int)))
				}
				if weeks, ok := advancedPolicyMap["weeks"]; ok {
					advancedRetentionPolicy.Weeks = helper.Uint64(uint64(weeks.(int)))
				}
				if months, ok := advancedPolicyMap["months"]; ok {
					advancedRetentionPolicy.Months = helper.Uint64(uint64(months.(int)))
				}
				if years, ok := advancedPolicyMap["years"]; ok {
					advancedRetentionPolicy.Years = helper.Uint64(uint64(years.(int)))
				}

				updateRequest.AdvancedRetentionPolicy = advancedRetentionPolicy
			}
		}
	}

	// 检查其他字段变化
	if d.HasChange("is_permanent") {
		hasChanges = true
		updateRequest.IsPermanent = helper.Bool(d.Get("is_permanent").(bool))
	}

	if d.HasChange("auto_backup_policy_name") {
		hasChanges = true
		updateRequest.AutoBackupPolicyName = helper.String(d.Get("auto_backup_policy_name").(string))
	}

	if d.HasChange("full_backup_interval") {
		hasChanges = true
		updateRequest.FullBackupInterval = helper.Int64(int64(d.Get("full_backup_interval").(int)))
	}

	if d.HasChange("retention_amount") {
		hasChanges = true
		updateRequest.RetentionAmount = helper.Uint64(uint64(d.Get("retention_amount").(int)))
	}

	if d.HasChange("is_activated") {
		hasChanges = true
		updateRequest.IsActivated = helper.Bool(d.Get("is_activated").(bool))
	}

	if hasChanges {
		err := brcService.ModifyAutoBackupPolicyAttribute(ctx, updateRequest)
		if err != nil {
			return err
		}
	}

	return resourceTencentCloudBrcAutoBackupPolicyRead(d, meta)
}

func resourceTencentCloudBrcAutoBackupPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_brc_auto_backup_policy.delete")()

	var (
		logId      = getLogId(contextNil)
		ctx        = context.WithValue(context.TODO(), logIdKey, logId)
		brcService = BRCService{client: meta.(*TencentCloudClient).apiV3Conn}
		policyId   = d.Id()
	)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		err := brcService.DeleteAutoBackupPolicies(ctx, []string{policyId})
		if err != nil {
			return retryError(err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

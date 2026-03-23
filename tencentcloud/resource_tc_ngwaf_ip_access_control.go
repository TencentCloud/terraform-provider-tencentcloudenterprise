/*
Provides a resource to create a NGWAF IP access control rule.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_ip_access_control" "example" {
  instance_id = "waf-xxxxxxxx"
  domain      = "example.com"
  edition     = "clb-waf"
  action_type = 42
  ip_list     = ["192.168.1.1", "192.168.1.2"]
  note        = "Block malicious IPs"
  source_type = "custom"
  valid_ts    = 0
  job_type    = "TimedJob"
}
```

Import

NGWAF IP access control rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_ip_access_control.example waf-xxxxxxxx#example.com#clb-waf#rule-xxxxxxxx
```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudNgwafIpAccessControl() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafIpAccessControlCreate,
		Read:   resourceTencentCloudNgwafIpAccessControlRead,
		Update: resourceTencentCloudNgwafIpAccessControlUpdate,
		Delete: resourceTencentCloudNgwafIpAccessControlDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Waf instance Id.",
			},
			"domain": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"edition": {
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(EDITION_TYPE),
				Description:  "Waf edition. clb-waf means clb-waf, sparta-waf means saas-waf.",
			},
			"action_type": {
				Required:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateAllowedIntValue([]int{40, 42}),
				Description:  "Action type, 40 for whitelist, 42 for blacklist.",
			},
			"ip_list": {
				Required:    true,
				Type:        schema.TypeList,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "IP address list.",
			},
			"note": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Note information.",
			},
			"source_type": {
				Optional:    true,
				Default:     "custom",
				Type:        schema.TypeString,
				Description: "Source type, default is custom.",
			},
			"valid_ts": {
				Optional:    true,
				Default:     0,
				Type:        schema.TypeInt,
				Description: "Valid timestamp, 0 means permanent.",
			},
			"job_type": {
				Optional:    true,
				Default:     "TimedJob",
				Type:        schema.TypeString,
				Description: "Job type, TimedJob or CronJob.",
			},
			"job_date_time": {
				Optional: true,
				Type:     schema.TypeList,
				MaxItems: 1,
				Description: "Job date time configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timed": {
							Optional: true,
							Type:     schema.TypeList,
							Description: "Timed job configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_date_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Start timestamp.",
									},
									"end_date_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "End timestamp.",
									},
								},
							},
						},
						"cron": {
							Optional: true,
							Type:     schema.TypeList,
							Description: "Cron job configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeInt},
										Description: "Days of month.",
									},
									"w_days": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem:     &schema.Schema{Type: schema.TypeInt},
										Description: "Days of week.",
									},
									"start_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Start time.",
									},
									"end_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "End time.",
									},
								},
							},
						},
						"time_t_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "UTC+8",
							Description: "Time zone.",
						},
					},
				},
			},
			"rule_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Rule ID.",
			},
		},
	}
}

func resourceTencentCloudNgwafIpAccessControlCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_ip_access_control.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = ngwaf.NewCreateIpAccessControlRequest()
		response = ngwaf.NewCreateIpAccessControlResponse()
		domain   string
		ruleId   int64
	)

	// Set basic fields
	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		domain = v.(string)
	}

	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("edition"); ok {
		request.Edition = helper.String(v.(string))
	}

	if v, ok := d.GetOk("action_type"); ok {
		actionType := int64(v.(int))
		request.ActionType = &actionType
	}

	if v, ok := d.GetOk("ip_list"); ok {
		ipList := v.([]interface{})
		for _, ip := range ipList {
			request.IpList = append(request.IpList, helper.String(ip.(string)))
		}
	}

	if v, ok := d.GetOk("note"); ok {
		request.Note = helper.String(v.(string))
	}

	if v, ok := d.GetOk("source_type"); ok {
		request.SourceType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("valid_ts"); ok {
		validTs := int64(v.(int))
		request.ValidTS = &validTs
	}

	if v, ok := d.GetOk("job_type"); ok {
		request.JobType = helper.String(v.(string))
	}

	// Set job date time
	if v, ok := d.GetOk("job_date_time"); ok {
		jobDateTimeList := v.([]interface{})
		if len(jobDateTimeList) > 0 {
			jobDateTimeMap := jobDateTimeList[0].(map[string]interface{})
			jobDateTime := &ngwaf.JobDateTime{}

			if timedList, ok := jobDateTimeMap["timed"].([]interface{}); ok && len(timedList) > 0 {
				for _, timedItem := range timedList {
					timedMap := timedItem.(map[string]interface{})
					timedJob := &ngwaf.TimedJob{}

					if startTime, ok := timedMap["start_date_time"].(int); ok {
						timedJob.StartDateTime = helper.IntUint64(startTime)
					}
					if endTime, ok := timedMap["end_date_time"].(int); ok {
						timedJob.EndDateTime = helper.IntUint64(endTime)
					}

					jobDateTime.Timed = append(jobDateTime.Timed, timedJob)
				}
			}

			if cronList, ok := jobDateTimeMap["cron"].([]interface{}); ok && len(cronList) > 0 {
				for _, cronItem := range cronList {
					cronMap := cronItem.(map[string]interface{})
					cronJob := &ngwaf.CronJob{}

					if daysSet, ok := cronMap["days"].(*schema.Set); ok {
						for _, day := range daysSet.List() {
							cronJob.Days = append(cronJob.Days, helper.IntUint64(day.(int)))
						}
					}

					if wDaysSet, ok := cronMap["w_days"].(*schema.Set); ok {
						for _, wDay := range wDaysSet.List() {
							cronJob.WDays = append(cronJob.WDays, helper.IntUint64(wDay.(int)))
						}
					}

					if startTime, ok := cronMap["start_time"].(string); ok {
						cronJob.StartTime = helper.String(startTime)
					}

					if endTime, ok := cronMap["end_time"].(string); ok {
						cronJob.EndTime = helper.String(endTime)
					}

					jobDateTime.Cron = append(jobDateTime.Cron, cronJob)
				}
			}

			if timeZone, ok := jobDateTimeMap["time_t_zone"].(string); ok {
				jobDateTime.TimeTZone = helper.String(timeZone)
			}

			request.JobDateTime = jobDateTime
		}
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().CreateIpAccessControl(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.RuleId == nil {
			return resource.NonRetryableError(fmt.Errorf("create ip access control failed, Response is nil"))
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create ip access control failed, reason:%+v", logId, err)
		return err
	}

	ruleId = int64(*response.Response.RuleId)
	d.SetId(strings.Join([]string{domain, strconv.FormatInt(ruleId, 10)}, FILED_SP))

	return resourceTencentCloudNgwafIpAccessControlRead(d, meta)
}

func resourceTencentCloudNgwafIpAccessControlRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_ip_access_control.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId = getLogId(contextNil)
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	ruleId := idSplit[1]

	// Use DescribeIpAccessControl to get the rule details
	request := ngwaf.NewDescribeIpAccessControlRequest()
	request.Domain = &domain
	
	// Get action_type from state or use default
	actionType := uint64(40) // default to whitelist
	if v, ok := d.GetOk("action_type"); ok {
		actionType = uint64(v.(int))
	}
	request.ActionType = &actionType
	request.OffSet = helper.IntUint64(0)
	request.Limit = helper.IntUint64(100)
	count := uint64(1)
	request.Count = &count

	response, err := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().DescribeIpAccessControl(request)
	if err != nil {
		return err
	}

	if response == nil || response.Response == nil || response.Response.Data == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource not found, please check if it has been deleted.\n", logId)
		return nil
	}

	// Find the specific rule by ruleId
	var rule *ngwaf.IpAccessControlItem
	for _, item := range response.Response.Data.Res {
		if item.RuleId != nil && strconv.FormatUint(*item.RuleId, 10) == ruleId {
			rule = item
			break
		}
	}

	if rule == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource not found, please check if it has been deleted.\n", logId)
		return nil
	}

	if rule.RuleId == nil {
		return fmt.Errorf("rule RuleId is nil")
	}

	_ = d.Set("domain", domain)
	_ = d.Set("rule_id", strconv.FormatUint(*rule.RuleId, 10))

	if rule.IpList != nil {
		_ = d.Set("ip_list", rule.IpList)
	}

	if rule.ActionType != nil {
		_ = d.Set("action_type", int(*rule.ActionType))
	}

	if rule.Note != nil {
		_ = d.Set("note", rule.Note)
	}

	if rule.Source != nil {
		_ = d.Set("source_type", rule.Source)
	}

	if rule.ValidTs != nil {
		_ = d.Set("valid_ts", int(*rule.ValidTs))
	}

	if rule.JobType != nil {
		_ = d.Set("job_type", rule.JobType)
	}

	if rule.JobDateTime != nil {
		jobDateTimeList := make([]map[string]interface{}, 0)
		jobDateTimeMap := make(map[string]interface{})

		if rule.JobDateTime.Timed != nil {
			timedList := make([]map[string]interface{}, 0)
			for _, timed := range rule.JobDateTime.Timed {
				timedMap := make(map[string]interface{})
				if timed.StartDateTime != nil {
					timedMap["start_date_time"] = int(*timed.StartDateTime)
				}
				if timed.EndDateTime != nil {
					timedMap["end_date_time"] = int(*timed.EndDateTime)
				}
				timedList = append(timedList, timedMap)
			}
			jobDateTimeMap["timed"] = timedList
		}

		if rule.JobDateTime.Cron != nil {
			cronList := make([]map[string]interface{}, 0)
			for _, cron := range rule.JobDateTime.Cron {
				cronMap := make(map[string]interface{})
				if cron.Days != nil {
					cronMap["days"] = cron.Days
				}
				if cron.WDays != nil {
					cronMap["w_days"] = cron.WDays
				}
				if cron.StartTime != nil {
					cronMap["start_time"] = cron.StartTime
				}
				if cron.EndTime != nil {
					cronMap["end_time"] = cron.EndTime
				}
				cronList = append(cronList, cronMap)
			}
			jobDateTimeMap["cron"] = cronList
		}

		if rule.JobDateTime.TimeTZone != nil {
			jobDateTimeMap["time_t_zone"] = rule.JobDateTime.TimeTZone
		}

		jobDateTimeList = append(jobDateTimeList, jobDateTimeMap)
		_ = d.Set("job_date_time", jobDateTimeList)
	}

	return nil
}

func resourceTencentCloudNgwafIpAccessControlUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_ip_access_control.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyIpAccessControlRequest()
	)

	immutableArgs := []string{"instance_id", "domain", "edition"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	ruleIdStr := idSplit[1]

	ruleId, err := strconv.ParseUint(ruleIdStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid rule id: %s", ruleIdStr)
	}

	request.RuleId = &ruleId
	request.Domain = &domain

	if v, ok := d.GetOk("instance_id"); ok {
		request.InstanceId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("edition"); ok {
		request.Edition = helper.String(v.(string))
	}

	if v, ok := d.GetOk("action_type"); ok {
		actionType := int64(v.(int))
		request.ActionType = &actionType
	}

	if v, ok := d.GetOk("ip_list"); ok {
		ipList := v.([]interface{})
		for _, ip := range ipList {
			request.IpList = append(request.IpList, helper.String(ip.(string)))
		}
	}

	if v, ok := d.GetOk("note"); ok {
		request.Note = helper.String(v.(string))
	}

	if v, ok := d.GetOk("source_type"); ok {
		request.SourceType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("valid_ts"); ok {
		validTs := int64(v.(int))
		request.ValidTS = &validTs
	}

	if v, ok := d.GetOk("job_type"); ok {
		request.JobType = helper.String(v.(string))
	}

	// Set job date time
	if v, ok := d.GetOk("job_date_time"); ok {
		jobDateTimeList := v.([]interface{})
		if len(jobDateTimeList) > 0 {
			jobDateTimeMap := jobDateTimeList[0].(map[string]interface{})
			jobDateTime := &ngwaf.JobDateTime{}

			if timedList, ok := jobDateTimeMap["timed"].([]interface{}); ok && len(timedList) > 0 {
				for _, timedItem := range timedList {
					timedMap := timedItem.(map[string]interface{})
					timedJob := &ngwaf.TimedJob{}

					if startTime, ok := timedMap["start_date_time"].(int); ok {
						timedJob.StartDateTime = helper.IntUint64(startTime)
					}
					if endTime, ok := timedMap["end_date_time"].(int); ok {
						timedJob.EndDateTime = helper.IntUint64(endTime)
					}

					jobDateTime.Timed = append(jobDateTime.Timed, timedJob)
				}
			}

			if cronList, ok := jobDateTimeMap["cron"].([]interface{}); ok && len(cronList) > 0 {
				for _, cronItem := range cronList {
					cronMap := cronItem.(map[string]interface{})
					cronJob := &ngwaf.CronJob{}

					if daysSet, ok := cronMap["days"].(*schema.Set); ok {
						for _, day := range daysSet.List() {
							cronJob.Days = append(cronJob.Days, helper.IntUint64(day.(int)))
						}
					}

					if wDaysSet, ok := cronMap["w_days"].(*schema.Set); ok {
						for _, wDay := range wDaysSet.List() {
							cronJob.WDays = append(cronJob.WDays, helper.IntUint64(wDay.(int)))
						}
					}

					if startTime, ok := cronMap["start_time"].(string); ok {
						cronJob.StartTime = helper.String(startTime)
					}

					if endTime, ok := cronMap["end_time"].(string); ok {
						cronJob.EndTime = helper.String(endTime)
					}

					jobDateTime.Cron = append(jobDateTime.Cron, cronJob)
				}
			}

			if timeZone, ok := jobDateTimeMap["time_t_zone"].(string); ok {
				jobDateTime.TimeTZone = helper.String(timeZone)
			}

			request.JobDateTime = jobDateTime
		}
	}

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyIpAccessControl(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf ipAccessControl failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudNgwafIpAccessControlRead(d, meta)
}

func resourceTencentCloudNgwafIpAccessControlDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_ip_access_control.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewDeleteIpAccessControlRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	ruleIdStr := idSplit[1]

	// Set required parameters
	request.Domain = &domain
	
	// Set action type
	if v, ok := d.GetOk("action_type"); ok {
		actionType := uint64(v.(int))
		request.ActionType = &actionType
	}

	// Set IsId to true since we're deleting by rule ID
	isId := true
	request.IsId = &isId

	// Add rule ID to items list
	request.Items = append(request.Items, &ruleIdStr)

	// Set delete all to false since we're deleting specific rule
	deleteAll := false
	request.DeleteAll = &deleteAll

	// Set source type if available
	if v, ok := d.GetOk("source_type"); ok {
		request.SourceType = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().DeleteIpAccessControl(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s delete ip access control failed, reason:%+v", logId, err)
		return err
	}

	return nil
}

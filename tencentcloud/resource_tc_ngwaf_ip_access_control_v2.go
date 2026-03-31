/*
Provides a resource to create a NGWAF IP access control rule v2.

# Example Usage

```hcl

	resource "tencentcloudenterprise_ngwaf_ip_access_control_v2" "example" {
	  instance_id = "waf_2kxtlbky11bbcr4b"
	  domain      = "example.com"
	  action_type = 40
	  note        = "note."

	  ip_list = [
	    "10.0.0.10",
	    "172.0.0.16",
	    "192.168.0.30",
	  ]

	  job_type = "TimedJob"

	  job_date_time {
	    time_t_zone = "UTC+8"

	    timed {
	      end_date_time   = 0
	      start_date_time = 0
	    }
	  }
	}

```

# Import

NGWAF IP access control rule v2 can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_ip_access_control_v2.example waf_2kxtlbky11bbcr4b#example.com#55000001
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_ngwaf_ip_access_control_v2", CNDescription{
		TerraformTypeCN: "NGWAF IP 黑白名单规则 V2",
		DescriptionCN:   "提供 NGWAF IP 黑白名单规则 V2 资源，用于创建和管理按规则 ID 维度维护的 IP 黑白名单。",
		AttributesCN: map[string]string{
			"domain":          "域名，global 表示全局域名",
			"ip_list":         "IP 地址列表",
			"action_type":     "动作类型，40 表示白名单，42 表示黑名单",
			"instance_id":     "WAF 实例 ID",
			"note":            "备注",
			"job_type":        "定时任务类型",
			"job_date_time":   "定时任务详情",
			"timed":           "定时执行配置",
			"cron":            "周期执行配置",
			"time_t_zone":     "时区",
			"start_date_time": "定时开始时间戳",
			"end_date_time":   "定时结束时间戳",
			"days":            "每月执行日期列表",
			"w_days":          "每周执行星期列表",
			"start_time":      "周期执行开始时间",
			"end_time":        "周期执行结束时间",
		},
	})
}

func resourceTencentCloudNgwafIpAccessControlV2() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafIpAccessControlV2Create,
		Read:   resourceTencentCloudNgwafIpAccessControlV2Read,
		Update: resourceTencentCloudNgwafIpAccessControlV2Update,
		Delete: resourceTencentCloudNgwafIpAccessControlV2Delete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specific domain name, for example test.qcloudwaf.com. Global domain name can be set to global.",
			},
			"ip_list": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "IP parameter list.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"action_type": {
				Type:         schema.TypeInt,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAllowedIntValue([]int{40, 42}),
				Description:  "42 means blocklist and 40 means allowlist.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Instance ID.",
			},
			"note": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remarks.",
			},
			"job_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Scheduled configuration type.",
			},
			"job_date_time": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Details of scheduled configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timed": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Time parameters for scheduled execution.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_date_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Start timestamp in seconds.",
									},
									"end_date_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "End timestamp in seconds.",
									},
								},
							},
						},
						"cron": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Time parameters for periodic execution.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Days in each month for execution.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"w_days": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Days of each week for execution.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
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
							Description: "Time zone.",
						},
					},
				},
			},
		},
	}
}

func resourceTencentCloudNgwafIpAccessControlV2Create(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_ip_access_control_v2.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = ngwaf.NewCreateIpAccessControlRequest()
		response   = ngwaf.NewCreateIpAccessControlResponse()
		instanceId string
		domain     string
	)

	if v, ok := d.GetOk("instance_id"); ok {
		instanceId = v.(string)
		request.InstanceId = helper.String(instanceId)
	}

	if v, ok := d.GetOk("domain"); ok {
		domain = v.(string)
		request.Domain = helper.String(domain)
	}

	if v, ok := d.GetOkExists("action_type"); ok {
		request.ActionType = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("ip_list"); ok {
		for _, item := range v.(*schema.Set).List() {
			request.IpList = append(request.IpList, helper.String(item.(string)))
		}
	}

	sourceType := "custom"
	request.SourceType = &sourceType

	if v, ok := d.GetOk("note"); ok {
		request.Note = helper.String(v.(string))
	}

	if v, ok := d.GetOk("job_type"); ok {
		request.JobType = helper.String(v.(string))
	}

	if jobDateTime := expandNgwafIpAccessControlV2JobDateTime(d); jobDateTime != nil {
		request.JobDateTime = jobDateTime
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().CreateIpAccessControl(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		response = result
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create ngwaf ip access control v2 failed, reason:%+v", logId, err)
		return err
	}

	if response == nil || response.Response == nil || response.Response.RuleId == nil {
		return fmt.Errorf("create ngwaf ip access control v2 failed: invalid response")
	}

	d.SetId(strings.Join([]string{instanceId, domain, helper.UInt64ToStr(*response.Response.RuleId)}, FILED_SP))

	return resourceTencentCloudNgwafIpAccessControlV2Read(d, meta)
}

func resourceTencentCloudNgwafIpAccessControlV2Read(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_ip_access_control_v2.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.Background(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	instanceId := idSplit[0]
	domain := idSplit[1]
	ruleID := idSplit[2]

	respData, err := service.DescribeWafIpAccessControlV2ById(ctx, domain, ruleID)
	if err != nil {
		return err
	}

	rule := findNgwafIpAccessControlV2Rule(respData, ruleID)
	if rule == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `ngwaf_ip_access_control_v2` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	return setNgwafIpAccessControlV2State(d, instanceId, domain, rule)
}

func resourceTencentCloudNgwafIpAccessControlV2Update(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_ip_access_control_v2.update")()
	defer inconsistentCheck(d, meta)()

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	instanceId := idSplit[0]
	domain := idSplit[1]
	ruleID, err := strconv.ParseUint(idSplit[2], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid rule id %q: %w", idSplit[2], err)
	}

	needChange := false
	for _, field := range []string{"ip_list", "note", "job_type", "job_date_time"} {
		if d.HasChange(field) {
			needChange = true
			break
		}
	}

	if needChange {
		var (
			logId   = getLogId(contextNil)
			request = ngwaf.NewModifyIpAccessControlRequest()
		)

		request.InstanceId = helper.String(instanceId)
		request.Domain = helper.String(domain)
		request.RuleId = &ruleID

		if v, ok := d.GetOkExists("action_type"); ok {
			request.ActionType = helper.IntInt64(v.(int))
		}

		sourceType := "custom"
		request.SourceType = &sourceType

		if v, ok := d.GetOk("ip_list"); ok {
			for _, item := range v.(*schema.Set).List() {
				request.IpList = append(request.IpList, helper.String(item.(string)))
			}
		}

		if v, ok := d.GetOk("note"); ok {
			request.Note = helper.String(v.(string))
		}

		if v, ok := d.GetOk("job_type"); ok {
			request.JobType = helper.String(v.(string))
		}

		if jobDateTime := expandNgwafIpAccessControlV2JobDateTime(d); jobDateTime != nil {
			request.JobDateTime = jobDateTime
		}

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyIpAccessControl(request)
			if e != nil {
				return retryError(e)
			}

			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s update ngwaf ip access control v2 failed, reason:%+v", logId, err)
			return err
		}
	}

	return resourceTencentCloudNgwafIpAccessControlV2Read(d, meta)
}

func resourceTencentCloudNgwafIpAccessControlV2Delete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_ip_access_control_v2.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewDeleteIpAccessControlV2Request()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	domain := idSplit[1]
	ruleID, err := strconv.ParseUint(idSplit[2], 10, 64)
	if err != nil {
		return fmt.Errorf("invalid rule id %q: %w", idSplit[2], err)
	}

	request.Domain = helper.String(domain)
	request.RuleIds = []*uint64{&ruleID}

	err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().DeleteIpAccessControlV2(request)
		if e != nil {
			return retryError(e)
		}

		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s delete ngwaf ip access control v2 failed, reason:%+v", logId, err)
		return err
	}

	return nil
}

func expandNgwafIpAccessControlV2JobDateTime(d *schema.ResourceData) *ngwaf.JobDateTime {
	jobDateTimeMap, ok := helper.InterfacesHeadMap(d, "job_date_time")
	if !ok {
		return nil
	}

	jobDateTime := &ngwaf.JobDateTime{}

	if v, ok := jobDateTimeMap["timed"]; ok {
		for _, item := range v.([]interface{}) {
			timedMap := item.(map[string]interface{})
			timedJob := &ngwaf.TimedJob{}

			if startDateTime, ok := timedMap["start_date_time"]; ok {
				timedJob.StartDateTime = helper.IntUint64(startDateTime.(int))
			}

			if endDateTime, ok := timedMap["end_date_time"]; ok {
				timedJob.EndDateTime = helper.IntUint64(endDateTime.(int))
			}

			jobDateTime.Timed = append(jobDateTime.Timed, timedJob)
		}
	}

	if v, ok := jobDateTimeMap["cron"]; ok {
		for _, item := range v.([]interface{}) {
			cronMap := item.(map[string]interface{})
			cronJob := &ngwaf.CronJob{}

			if days, ok := cronMap["days"]; ok {
				for _, day := range days.(*schema.Set).List() {
					cronJob.Days = append(cronJob.Days, helper.IntUint64(day.(int)))
				}
			}

			if wDays, ok := cronMap["w_days"]; ok {
				for _, wDay := range wDays.(*schema.Set).List() {
					cronJob.WDays = append(cronJob.WDays, helper.IntUint64(wDay.(int)))
				}
			}

			if startTime, ok := cronMap["start_time"]; ok {
				cronJob.StartTime = helper.String(startTime.(string))
			}

			if endTime, ok := cronMap["end_time"]; ok {
				cronJob.EndTime = helper.String(endTime.(string))
			}

			jobDateTime.Cron = append(jobDateTime.Cron, cronJob)
		}
	}

	if v, ok := jobDateTimeMap["time_t_zone"]; ok {
		jobDateTime.TimeTZone = helper.String(v.(string))
	}

	return jobDateTime
}

func setNgwafIpAccessControlV2State(d *schema.ResourceData, instanceId, domain string, rule *ngwaf.IpAccessControlItem) error {
	_ = d.Set("instance_id", instanceId)
	_ = d.Set("domain", domain)

	ipList := make([]string, 0)
	if rule.IpList != nil {
		for _, item := range rule.IpList {
			if item != nil {
				ipList = append(ipList, *item)
			}
		}
	}
	_ = d.Set("ip_list", ipList)

	if rule.ActionType != nil {
		_ = d.Set("action_type", int(*rule.ActionType))
	}

	if rule.Note != nil {
		_ = d.Set("note", *rule.Note)
	} else {
		_ = d.Set("note", "")
	}

	if rule.JobType != nil {
		_ = d.Set("job_type", *rule.JobType)
	} else {
		_ = d.Set("job_type", "")
	}

	_ = d.Set("job_date_time", flattenNgwafIpAccessControlV2JobDateTime(rule.JobDateTime))

	return nil
}

func flattenNgwafIpAccessControlV2JobDateTime(jobDateTime *ngwaf.JobDateTime) []interface{} {
	if jobDateTime == nil {
		return []interface{}{}
	}

	jobDateTimeMap := make(map[string]interface{})

	if jobDateTime.Timed != nil {
		timedList := make([]map[string]interface{}, 0, len(jobDateTime.Timed))
		for _, timed := range jobDateTime.Timed {
			if timed == nil {
				continue
			}

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

	if jobDateTime.Cron != nil {
		cronList := make([]map[string]interface{}, 0, len(jobDateTime.Cron))
		for _, cron := range jobDateTime.Cron {
			if cron == nil {
				continue
			}

			cronMap := make(map[string]interface{})
			if cron.Days != nil {
				days := make([]int, 0, len(cron.Days))
				for _, day := range cron.Days {
					if day != nil {
						days = append(days, int(*day))
					}
				}
				cronMap["days"] = days
			}

			if cron.WDays != nil {
				wDays := make([]int, 0, len(cron.WDays))
				for _, wDay := range cron.WDays {
					if wDay != nil {
						wDays = append(wDays, int(*wDay))
					}
				}
				cronMap["w_days"] = wDays
			}

			if cron.StartTime != nil {
				cronMap["start_time"] = *cron.StartTime
			}
			if cron.EndTime != nil {
				cronMap["end_time"] = *cron.EndTime
			}

			cronList = append(cronList, cronMap)
		}
		jobDateTimeMap["cron"] = cronList
	}

	if jobDateTime.TimeTZone != nil {
		jobDateTimeMap["time_t_zone"] = *jobDateTime.TimeTZone
	}

	return []interface{}{jobDateTimeMap}
}

func findNgwafIpAccessControlV2Rule(respData *ngwaf.DescribeIpAccessControlResponse, ruleID string) *ngwaf.IpAccessControlItem {
	if respData == nil || respData.Response == nil || respData.Response.Data == nil {
		return nil
	}

	for _, item := range respData.Response.Data.Res {
		if item == nil || item.RuleId == nil {
			continue
		}

		if helper.UInt64ToStr(*item.RuleId) == ruleID {
			return item
		}
	}

	return nil
}

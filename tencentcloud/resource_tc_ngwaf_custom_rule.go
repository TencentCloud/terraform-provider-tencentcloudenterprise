/*
Provides a resource to create a NGWAF custom rule.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_custom_rule" "example" {
  name        = "example-rule"
  sort_id     = "10"
  expire_time = "0"
  domain      = "example.com"
  edition     = "clb-waf"
  action_type = "1"
  status      = "1"
  logical_op  = "and"
  
  strategies {
    field              = "IP"
    compare_func       = "ipmatch"
    content            = "192.168.1.1,192.168.1.2"
    arg                = ""
    case_not_sensitive = 0
  }
  
  strategies {
    field              = "URL"
    compare_func       = "contains"
    content            = "/admin"
    arg                = ""
    case_not_sensitive = 1
  }
}
```

Import

NGWAF custom rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_custom_rule.example example.com#clb-waf#rule-xxxxxxxx
```
*/
package tencentcloud

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"strconv"
	"strings"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudNgwafCustomRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafCustomRuleCreate,
		Read:   resourceTencentCloudNgwafCustomRuleRead,
		Update: resourceTencentCloudNgwafCustomRuleUpdate,
		Delete: resourceTencentCloudNgwafCustomRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Rule Name.",
			},
			"sort_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Priority, value range 0-100.",
			},
			"redirect": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "If the action is a Redirect, it represents the redirect address; Other situations can be left blank.",
			},
			"expire_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Expiration time in second-level timestamp, for example, 1677254399 indicates the expiration time is 2023-02-24 23:59:59; 0 indicates it will never expire.",
			},
			"strategies": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "Strategies detail.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Matching field
              Different matching fields result in different matching parameters, logical operators, and matching contents. The details are as follows:
			  <table><thead><tr><th>Matching Field</th><th>Matching Parameter</th><th>Logical Symbol</th><th>Matching Content</th></tr></thead><tbody><tr><td>IP (source IP)</td><td>Parameters are not supported.</td><td>ipmatch (match)<br>ipnmatch (mismatch)</td><td>Multiple IP addresses are separated by commas. A maximum of 20 IP addresses are allowed.</td></tr><tr><td>IPv6 (source IPv6)</td><td>Parameters are not supported.</td><td>ipmatch (match)<br>ipnmatch (mismatch)</td><td>A single IPv6 address is supported.</td></tr><tr><td>Referer (referer)</td><td>Parameters are not supported.</td><td>empty (Content is empty.)<br>null (do not exist)<br>eq (equal to)<br>neq (not equal to)<br>contains (contain)<br>ncontains (do not contain)<br>len_eq (length equals to)<br>len_gt (length is greater than)<br>len_lt (length is less than)<br>strprefix (prefix matching)<br>strsuffix (suffix matching)<br>rematch (regular expression matching)</td><td>Enter the content, with a maximum of 512 characters.</td></tr><tr><td>URL (request path)</td><td>Parameters are not supported.</td><td>eq (equal to)<br>neq (not equal to)<br>contains (contain)<br>ncontains (do not contain)<br>len_eq (length equals to)<br>len_gt (length is greater than)<br>len_lt (length is less than)<br>strprefix (prefix matching)<br>strsuffix (suffix matching)<br>rematch (regular expression matching)</td><td>Enter the content starting with /, with a maximum of 512 characters.</td></tr><tr><td>UserAgent (UserAgent)</td><td>Parameters are not supported.</td><td>Same logical symbols as the matching field <font color="Red">Referer</font></td><td>Enter the content with a maximum of 512 characters.</td></tr><tr><td>HTTP_METHOD (HTTP request method)</td><td>Parameters are not supported.</td><td>eq (equal to)<br>neq (not equal to)</td><td>Enter the method name. The uppercase is recommended.</td></tr><tr><td>QUERY_STRING (request string)</td><td>Parameters are not supported.</td><td>Same logical symbol as the matching field <font color=... [truncated]
          	  Note: This field may return null, indicating that no valid values can be obtained.`,
						},
						"compare_func": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Logic symbol
              Logical symbols are divided into the following types:
                  empty (content is empty)
                  null (do not exist)
                  eq (equal to)
                  neq (not equal to)
                  contains (contain)
                  ncontains (do not contain)
                  strprefix (prefix matching)
                  strsuffix (suffix matching)
                  len_eq (length equals to)
                  len_gt (length is greater than)
                  len_lt (length is less than)
                  ipmatch (belong to)
                  ipnmatch (do not belong to)
                  numgt (number greater than)
                  numlt (number less than)
                  geo_in (IP geo belongs to)
                  geo_not_in (IP geo not belongs to)
                  rematch (regex match)
				  numgt (numerically greater than)
				  numlt (numerically less than)
				  numeq (numerically equal to)
				  numneq (numerically not equal to)
				  numle (numerically less than or equal to)
				  numge (numerically greater than or equal to)
              Different matching fields correspond to different logical operators. For details, see the matching field table above.
          Note: This field may return null, indicating that no valid values can be obtained.`,
						},
						"content": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Matching content
              Currently, when the matching field is COOKIE (cookie), the matching content is not required. In other scenes, the matching content is required.
          Note: This field may return null, indicating that no valid values can be obtained.`,
						},
						"arg": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Matching parameter
              There are two types of configuration parameters: unsupported parameters and supported parameters.
              The matching parameter can be entered only when the matching field is one of the following four. Otherwise, the parameter is not supported.
                  GET (GET parameter value)		
                  POST (POST parameter value)		
                  ARGS_COOKIE (Cookie parameter value)		
                  ARGS_HEADER (Header parameter value)
          Note: This field may return null, indicating that no valid values can be obtained.`,
						},
						"case_not_sensitive": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "0: case-sensitive, 1: case-insensitive. Note: This field may return null, indicating that no valid values can be obtained.",
						},
					},
				},
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
			Type:         schema.TypeString,
			ValidateFunc: validateAllowedStringValue(CUSTOM_RULE_ACTION_TYPE),
			Description:  "Action type, 1(Block), 2(Captcha), 3(log), 4(Redirect).",
		},
			"status": {
				Optional:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(CUSTOM_RULE_STATUS),
				Default:      CUSTOM_RULE_STATUS_1,
				Description:  "The status of the rule, 1(open), 0(close).",
			},
			"job_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Rule execution mode: TimedJob indicates scheduled execution. CronJob indicates periodic execution.",
			},
			"job_date_time": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Rule execution time.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timed": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Time parameters for scheduled execution. Note: This field may return null, indicating that no valid values can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_date_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Start timestamp, in seconds. Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"end_date_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "End timestamp, in seconds. Note: This field may return null, indicating that no valid values can be obtained.",
									},
								},
							},
						},
						"cron": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Time parameters for periodic execution. Note: This field may return null, indicating that no valid values can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Days in each month for execution. Note: This field may return null, indicating that no valid values can be obtained.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"w_days": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Days of each week for execution. Note: This field may return null, indicating that no valid values can be obtained.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"start_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Start time. Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"end_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "End time. Note: This field may return null, indicating that no valid values can be obtained.",
									},
								},
							},
						},
						"time_t_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time zone. Note: This field may return null, indicating that no valid values can be obtained.",
						},
					},
				},
			},
			"rule_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "rule ID.",
			},
			"logical_op": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Logical relation between strategies, can be 'and' or 'or'.",
			},
		},
	}
}

func resourceTencentCloudNgwafCustomRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_custom_rule.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		request       = ngwaf.NewAddCustomRuleRequest()
		response      = ngwaf.NewAddCustomRuleResponse()
		statusRequest = ngwaf.NewModifyCustomRuleStatusRequest()
		domain        string
		ruleIdStr     string
		status        string
	)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("sort_id"); ok {
		request.SortId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("redirect"); ok {
		request.Redirect = helper.String(v.(string))
	}

	if v, ok := d.GetOk("expire_time"); ok {
		request.ExpireTime = helper.String(v.(string))
	}

	if v, ok := d.GetOk("strategies"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			strategy := ngwaf.Strategy{}
			var base46Flag bool
			if v, ok := dMap["field"]; ok {
				strategy.Field = helper.String(v.(string))
			}

			if v, ok := dMap["compare_func"]; ok {
				strategy.CompareFunc = helper.String(v.(string))
				if v.(string) == "rematch" {
					base46Flag = true
				}
			}

			if v, ok := dMap["content"]; ok {
				if base46Flag {
					strategy.Content = helper.String(base64.URLEncoding.EncodeToString([]byte(v.(string))))
				} else {
					strategy.Content = helper.String(v.(string))
				}
			}

			if v, ok := dMap["arg"]; ok {
				strategy.Arg = helper.String(v.(string))
			}

			if v, ok := dMap["case_not_sensitive"]; ok {
				strategy.CaseNotSensitive = helper.IntUint64(v.(int))
			}

			request.Strategies = append(request.Strategies, &strategy)
		}
	}

if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		domain = v.(string)
	}

	if v, ok := d.GetOk("edition"); ok {
		request.Edition = helper.String(v.(string))
	}

	if v, ok := d.GetOk("action_type"); ok {
		request.ActionType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("job_type"); ok {
		request.JobType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("logical_op"); ok {
		request.LogicalOp = helper.String(v.(string))
	}

	if jobDateTimeMap, ok := helper.InterfacesHeadMap(d, "job_date_time"); ok {
		jobDateTime := ngwaf.JobDateTime{}
		if v, ok := jobDateTimeMap["timed"]; ok {
			for _, item := range v.([]interface{}) {
				timedMap := item.(map[string]interface{})
				timedJob := ngwaf.TimedJob{}
				if v, ok := timedMap["start_date_time"].(int); ok {
					timedJob.StartDateTime = helper.IntUint64(v)
				}

				if v, ok := timedMap["end_date_time"].(int); ok {
					timedJob.EndDateTime = helper.IntUint64(v)
				}

				jobDateTime.Timed = append(jobDateTime.Timed, &timedJob)
			}
		}

		if v, ok := jobDateTimeMap["cron"]; ok {
			for _, item := range v.([]interface{}) {
				cronMap := item.(map[string]interface{})
				cronJob := ngwaf.CronJob{}
				if v, ok := cronMap["days"]; ok {
					daysSet := v.(*schema.Set).List()
					for i := range daysSet {
						days := daysSet[i].(int)
						cronJob.Days = append(cronJob.Days, helper.IntUint64(days))
					}
				}

				if v, ok := cronMap["w_days"]; ok {
					wDaysSet := v.(*schema.Set).List()
					for i := range wDaysSet {
						wDays := wDaysSet[i].(int)
						cronJob.WDays = append(cronJob.WDays, helper.IntUint64(wDays))
					}
				}

				if v, ok := cronMap["start_time"].(string); ok && v != "" {
					cronJob.StartTime = helper.String(v)
				}

				if v, ok := cronMap["end_time"].(string); ok && v != "" {
					cronJob.EndTime = helper.String(v)
				}

				jobDateTime.Cron = append(jobDateTime.Cron, &cronJob)
			}
		}

		if v, ok := jobDateTimeMap["time_t_zone"].(string); ok && v != "" {
			jobDateTime.TimeTZone = helper.String(v)
		}

		request.JobDateTime = &jobDateTime
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().AddCustomRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create waf CustomRule failed, reason:%+v", logId, err)
		return err
	}

	ruleId := *response.Response.RuleId
	ruleIdStr = strconv.FormatInt(ruleId, 10)

	if v, ok := d.GetOk("status"); ok {
		status = v.(string)
	}

	if status == CUSTOM_RULE_STATUS_0 {
		statusRequest.Domain = &domain
		tmpRuleId := uint64(ruleId)
		statusRequest.RuleId = &tmpRuleId
		statusRequest.Status = helper.IntUint64(CUSTOM_RULE_STATUS_0_INT)
		err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyCustomRuleStatus(statusRequest)
			if e != nil {
				return retryError(e)
			} else {
				log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
			}

			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s modify waf CustomRule status failed, reason:%+v", logId, err)
			return err
		}
	}

	d.SetId(strings.Join([]string{domain, ruleIdStr}, FILED_SP))
	return resourceTencentCloudNgwafCustomRuleRead(d, meta)
}

func resourceTencentCloudNgwafCustomRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_custom_rule.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken, %s", idSplit)
	}

	domain := idSplit[0]
	ruleId := idSplit[1]

	customRule, err := service.DescribeWafCustomRuleById(ctx, domain, ruleId)
	if err != nil {
		return err
	}

	if customRule == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafCustomRule` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if customRule.Name != nil {
		_ = d.Set("name", customRule.Name)
	}

	if customRule.SortId != nil {
		_ = d.Set("sort_id", customRule.SortId)
	}

	if customRule.Redirect != nil {
		_ = d.Set("redirect", customRule.Redirect)
	}

	if customRule.ExpireTime != nil {
		_ = d.Set("expire_time", customRule.ExpireTime)
	}

	if customRule.Strategies != nil {
		strategiesList := []interface{}{}
		for _, strategies := range customRule.Strategies {
			strategiesMap := map[string]interface{}{}
			var base46Flag bool
			if strategies.Field != nil {
				strategiesMap["field"] = strategies.Field
			}

			if strategies.CompareFunc != nil {
				strategiesMap["compare_func"] = strategies.CompareFunc
				if *strategies.CompareFunc == "rematch" {
					base46Flag = true
				}
			}

			if strategies.Content != nil {
				if base46Flag {
					decoded, e := base64.StdEncoding.DecodeString(*strategies.Content)
					if e != nil {
						return fmt.Errorf("[%s] base64 decode error: %s", *strategies.Content, e.Error())
					}

					strategiesMap["content"] = string(decoded)
				} else {
					strategiesMap["content"] = strategies.Content
				}
			}

			if strategies.Arg != nil {
				strategiesMap["arg"] = strategies.Arg
			}

			if strategies.CaseNotSensitive != nil {
				strategiesMap["case_not_sensitive"] = strategies.CaseNotSensitive
			}

			strategiesList = append(strategiesList, strategiesMap)
		}

		_ = d.Set("strategies", strategiesList)

	}

	_ = d.Set("domain", domain)

	if customRule.ActionType != nil {
		_ = d.Set("action_type", customRule.ActionType)
	}

	if customRule.JobType != nil {
		_ = d.Set("job_type", customRule.JobType)
	}

	if customRule.JobDateTime != nil {
		tmpList := make([]map[string]interface{}, 0)
		dMap := map[string]interface{}{}
		if customRule.JobDateTime.Timed != nil {
			timedList := []interface{}{}
			for _, v := range customRule.JobDateTime.Timed {
				timedMap := map[string]interface{}{}
				if v.StartDateTime != nil {
					timedMap["start_date_time"] = v.StartDateTime
				}

				if v.EndDateTime != nil {
					timedMap["end_date_time"] = v.EndDateTime
				}

				timedList = append(timedList, timedMap)
			}

			dMap["timed"] = timedList
		}

		if customRule.JobDateTime.Cron != nil {
			cronList := []interface{}{}
			for _, v := range customRule.JobDateTime.Cron {
				cronMap := map[string]interface{}{}
				if v.Days != nil {
					cronMap["days"] = v.Days
				}

				if v.WDays != nil {
					cronMap["w_days"] = v.WDays
				}

				if v.StartTime != nil {
					cronMap["start_time"] = v.StartTime
				}

				if v.EndTime != nil {
					cronMap["end_time"] = v.EndTime
				}

				cronList = append(cronList, cronMap)
			}

			dMap["cron"] = cronList
		}

		if customRule.JobDateTime.TimeTZone != nil {
			dMap["time_t_zone"] = customRule.JobDateTime.TimeTZone
		}

		tmpList = append(tmpList, dMap)

		_ = d.Set("job_date_time", tmpList)
	}

	if customRule.Status != nil {
		_ = d.Set("status", customRule.Status)
	}

	if customRule.RuleId != nil {
		_ = d.Set("rule_id", customRule.RuleId)
	}

	if customRule.LogicalOp != nil {
		_ = d.Set("logical_op", customRule.LogicalOp)
	}

	return nil
}

func resourceTencentCloudNgwafCustomRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_custom_rule.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId         = getLogId(contextNil)
		request       = ngwaf.NewModifyCustomRuleRequest()
		statusRequest = ngwaf.NewModifyCustomRuleStatusRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	domain := idSplit[0]
	ruleId := idSplit[1]

immutableArgs := []string{"domain"}

	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	request.Domain = &domain
	ruleIdInt, _ := strconv.ParseInt(ruleId, 10, 64)
	ruleIdUInt := uint64(ruleIdInt)
	request.RuleId = &ruleIdUInt

	if v, ok := d.GetOk("edition"); ok {
		request.Edition = helper.String(v.(string))
	}

	if v, ok := d.GetOk("name"); ok {
		request.RuleName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("action_type"); ok {
		request.RuleAction = helper.String(v.(string))
	}

	if v, ok := d.GetOk("redirect"); ok {
		request.Redirect = helper.String(v.(string))
	}

	if v, ok := d.GetOk("sort_id"); ok {
		sortIdInt, _ := strconv.ParseInt(v.(string), 10, 64)
		sortIdUInt := uint64(sortIdInt)
		request.SortId = &sortIdUInt
	}

	if v, ok := d.GetOk("expire_time"); ok {
		expireTimeInt, _ := strconv.ParseInt(v.(string), 10, 64)
		expireTimeUInt := uint64(expireTimeInt)
		request.ExpireTime = &expireTimeUInt
	}

	if v, ok := d.GetOk("strategies"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			strategy := ngwaf.Strategy{}
			var base46Flag bool
			if v, ok := dMap["field"]; ok {
				strategy.Field = helper.String(v.(string))
			}

			if v, ok := dMap["compare_func"]; ok {
				strategy.CompareFunc = helper.String(v.(string))
				if v.(string) == "rematch" {
					base46Flag = true
				}
			}

			if v, ok := dMap["content"]; ok {
				if base46Flag {
					strategy.Content = helper.String(base64.URLEncoding.EncodeToString([]byte(v.(string))))
				} else {
					strategy.Content = helper.String(v.(string))
				}
			}

			if v, ok := dMap["arg"]; ok {
				strategy.Arg = helper.String(v.(string))
			}

			if v, ok := dMap["case_not_sensitive"]; ok {
				strategy.CaseNotSensitive = helper.IntUint64(v.(int))
			}

			request.Strategies = append(request.Strategies, &strategy)
		}
	}

	if v, ok := d.GetOk("job_type"); ok {
		request.JobType = helper.String(v.(string))
	}

	if jobDateTimeMap, ok := helper.InterfacesHeadMap(d, "job_date_time"); ok {
		jobDateTime := ngwaf.JobDateTime{}
		if v, ok := jobDateTimeMap["timed"]; ok {
			for _, item := range v.([]interface{}) {
				timedMap := item.(map[string]interface{})
				timedJob := ngwaf.TimedJob{}
				if v, ok := timedMap["start_date_time"].(int); ok {
					timedJob.StartDateTime = helper.IntUint64(v)
				}

				if v, ok := timedMap["end_date_time"].(int); ok {
					timedJob.EndDateTime = helper.IntUint64(v)
				}

				jobDateTime.Timed = append(jobDateTime.Timed, &timedJob)
			}
		}

		if v, ok := jobDateTimeMap["cron"]; ok {
			for _, item := range v.([]interface{}) {
				cronMap := item.(map[string]interface{})
				cronJob := ngwaf.CronJob{}
				if v, ok := cronMap["days"]; ok {
					daysSet := v.(*schema.Set).List()
					for i := range daysSet {
						days := daysSet[i].(int)
						cronJob.Days = append(cronJob.Days, helper.IntUint64(days))
					}
				}

				if v, ok := cronMap["w_days"]; ok {
					wDaysSet := v.(*schema.Set).List()
					for i := range wDaysSet {
						wDays := wDaysSet[i].(int)
						cronJob.WDays = append(cronJob.WDays, helper.IntUint64(wDays))
					}
				}

				if v, ok := cronMap["start_time"].(string); ok && v != "" {
					cronJob.StartTime = helper.String(v)
				}

				if v, ok := cronMap["end_time"].(string); ok && v != "" {
					cronJob.EndTime = helper.String(v)
				}

				jobDateTime.Cron = append(jobDateTime.Cron, &cronJob)
			}
		}

		if v, ok := jobDateTimeMap["time_t_zone"].(string); ok && v != "" {
			jobDateTime.TimeTZone = helper.String(v)
		}

		request.JobDateTime = &jobDateTime
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyCustomRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf CustomRule failed, reason:%+v", logId, err)
		return err
	}

	if d.HasChange("status") {
		if v, ok := d.GetOk("status"); ok {
			status := v.(string)
			statusRequest.Domain = &domain
			statusRequest.RuleId = &ruleIdUInt
			if status == CUSTOM_RULE_STATUS_0 {
				statusRequest.Status = helper.IntUint64(CUSTOM_RULE_STATUS_0_INT)
			} else {
				statusRequest.Status = helper.IntUint64(CUSTOM_RULE_STATUS_1_INT)
			}

			err = resource.Retry(writeRetryTimeout, func() *resource.RetryError {
				result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyCustomRuleStatus(statusRequest)
				if e != nil {
					return retryError(e)
				} else {
					log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
				}

				return nil
			})

			if err != nil {
				log.Printf("[CRITAL]%s modify waf CustomRule status failed, reason:%+v", logId, err)
				return err
			}
		}
	}

	return resourceTencentCloudNgwafCustomRuleRead(d, meta)
}

func resourceTencentCloudNgwafCustomRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_custom_rule.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}

	domain := idSplit[0]
	ruleId := idSplit[1]

	if err := service.DeleteWafCustomRuleById(ctx, domain, ruleId); err != nil {
		return err
	}

	return nil
}

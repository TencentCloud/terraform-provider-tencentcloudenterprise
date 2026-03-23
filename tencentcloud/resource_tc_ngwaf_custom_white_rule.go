/*
Provides a resource to create a NGWAF custom white rule.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_custom_white_rule" "example" {
  name        = "example-white-rule"
  sort_id     = "10"
  expire_time = "0"
  domain      = "example.com"
  bypass      = "ACL,OWASP,CC"
  logical_op  = "and"
  
  strategies {
    field              = "IP"
    compare_func       = "ipmatch"
    content            = "192.168.1.100"
    arg                = ""
    case_not_sensitive = 0
  }
  
  strategies {
    field              = "URL"
    compare_func       = "contains"
    content            = "/api"
    arg                = ""
    case_not_sensitive = 1
  }
}
```

Import

NGWAF custom white rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_custom_white_rule.example example.com#rule-xxxxxxxx
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

func resourceTencentCloudNgwafCustomWhiteRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafCustomWhiteRuleCreate,
		Read:   resourceTencentCloudNgwafCustomWhiteRuleRead,
		Update: resourceTencentCloudNgwafCustomWhiteRuleUpdate,
		Delete: resourceTencentCloudNgwafCustomWhiteRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Name of the custom white rule. Must be unique within the domain. Used for identification and management purposes.",
			},
			"sort_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Priority level for rule execution. Valid range: 1-100. Lower numbers indicate higher priority (e.g., 1 has the highest priority). Rules are executed in ascending order of sort_id.",
			},
			"expire_time": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Expiration time for the rule in Unix timestamp format (seconds). Example: 1677254399 represents 2023-02-24 23:59:59. Use 0 for permanent rules that never expire.",
			},
			"strategies": {
				Required:    true,
				Type:        schema.TypeList,
				Description: "List of matching strategies that define the conditions for the white rule to trigger. Each strategy specifies a field to match, comparison logic, and the content to match against. Multiple strategies can be combined using logical operators.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field": {
							Type:     schema.TypeString,
							Required: true,
							Description: `Matching field
                            Different matching fields result in different matching parameters, logical operators, and matching contents. The details are as follows:
                        	<table><thead><tr><th>Matching Field</th><th>Matching Parameter</th><th>Logical Symbol</th><th>Matching Content</th></tr></thead><tbody><tr><td>IP (source IP)</td><td>Parameters are not supported.</td><td>ipmatch (match)<br>ipnmatch (mismatch)</td><td>Multiple IP addresses are separated by commas. A maximum of 20 IP addresses are allowed.</td></tr><tr><td>IPv6 (source IPv6)</td><td>Parameters are not supported.</td><td>ipmatch (match)<br>ipnmatch (mismatch)</td><td>A single IPv6 address is supported.</td></tr><tr><td>Referer (referer)</td><td>Parameters are not supported.</td><td>empty (Content is empty.)<br>null (do not exist)<br>eq (equal to)<br>neq (not equal to)<br>contains (contain)<br>ncontains (do not contain)<br>len_eq (length equals to)<br>len_gt (length is greater than)<br>len_lt (length is less than)<br>strprefix (prefix matching)<br>strsuffix (suffix matching)<br>rematch (regular expression matching)</td><td>Enter the content, with a maximum of 512 characters.</td></tr><tr><td>URL (request path)</td><td>Parameters are not supported.</td><td>eq (equal to)<br>neq (not equal to)<br>contains (contain)<br>ncontains (do not contain)<br>len_eq (length equals to)<br>len_gt (length is greater than)<br>len_lt (length is less than)<br>strprefix (prefix matching)<br>strsuffix (suffix matching)<br>rematch (regular expression matching)</td><td>Enter the content starting with /, with a maximum of 512 characters.</td></tr><tr><td>UserAgent (UserAgent)</td><td>Parameters are not supported.</td><td>Same logical symbols as the matching field <font color="Red">Referer</font></td><td>Enter the content with a maximum of 512 characters.</td></tr><tr><td>HTTP_METHOD (HTTP request method)</td><td>Parameters are not supported.</td><td>eq (equal to)<br>neq (not equal to)</td><td>Enter the method name. The uppercase is recommended.</td></tr><tr><td>QUERY_STRING (request string)</td><td>Parameters are not supported.</td><td>Same logical symbol as the matchin... [truncated]
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
							Optional: true,
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
				Description: "Domain name to which this custom white rule applies. The rule will only be effective for requests to this specific domain.",
			},
			"bypass": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Comma-separated list of WAF modules to bypass when the rule matches. Supported modules: ACL (Custom Rules), OWASP (Rule Engine), Webshell (Malicious File Detection), GeoIP (Geographic Block), BWIP (IP Black and White List), CC (Rate Limiting), BotRPC (BOT Protection), AntiLeakage (Information Leakage Prevention), API (API Security), AI (AI Engine), ip_outo_deny (IP Block), Applet (Mini Program Traffic Risk Control). Example: 'ACL,OWASP,CC'.",
			},
			"job_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Rule execution mode. Determines how the rule timing is configured. Valid values: 'TimedJob' for scheduled execution (specific start/end times), 'CronJob' for periodic execution (recurring pattern). Required when job_date_time is specified.",
			},
			"job_date_time": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Rule execution time configuration. Required when job_type is set. Contains timing parameters for scheduled or periodic execution.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"timed": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Time parameters for scheduled execution (TimedJob). Defines specific start and end timestamps for one-time execution. Note: This field may return null, indicating that no valid values can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"start_date_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Start timestamp for scheduled execution, in Unix timestamp format (seconds). Example: 1677254399 represents 2023-02-24 23:59:59. Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"end_date_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "End timestamp for scheduled execution, in Unix timestamp format (seconds). Must be greater than start_date_time. Example: 1677340799 represents 2023-02-25 23:59:59. Note: This field may return null, indicating that no valid values can be obtained.",
									},
								},
							},
						},
						"cron": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Time parameters for periodic execution (CronJob). Defines recurring execution patterns with days of week/month and time ranges. Note: This field may return null, indicating that no valid values can be obtained.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"days": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Days in each month for execution (1-31). Used for monthly scheduling. Can specify multiple days. Note: This field may return null, indicating that no valid values can be obtained.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"w_days": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Days of each week for execution (0-6, where 0=Sunday, 1=Monday, ..., 6=Saturday). Used for weekly scheduling. Can specify multiple days. Note: This field may return null, indicating that no valid values can be obtained.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"start_time": {
										Type:        schema.TypeString,
										Optional:    true,
									Description: "Start time for execution window in HH:mm format (24-hour clock). Example: \"09:00\" means execution starts at 9:00 AM. Note: This field may return null, indicating that no valid values can be obtained.",
									},
									"end_time": {
										Type:        schema.TypeString,
										Optional:    true,
									Description: "End time for execution window in HH:mm format (24-hour clock). Example: \"18:00\" means execution ends at 6:00 PM. Must be later than start_time. Note: This field may return null, indicating that no valid values can be obtained.",
									},
								},
							},
						},
						"time_t_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time zone for rule execution. Format: \"UTC+X\" or \"UTC-X\", where X is the offset from UTC. Example: \"UTC+8\" for China Standard Time, \"UTC-5\" for Eastern Standard Time. Note: This field may return null, indicating that no valid values can be obtained.",
						},
					},
				},
			},
			"logical_op": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Logical operator for combining multiple strategy conditions. Determines how multiple matching conditions are evaluated. Valid values: 'and' (all conditions must be true), 'or' (any condition can be true). Default is 'and' if not specified.",
			},
			"rule_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "The unique identifier of the custom white rule, automatically generated upon creation. Used to identify and manage the rule.",
			},
		},
	}
}

func resourceTencentCloudNgwafCustomWhiteRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_custom_white_rule.create")()
	defer inconsistentCheck(d, meta)()

var (
		logId    = getLogId(contextNil)
		request  = ngwaf.NewAddCustomWhiteRuleRequest()
		response = ngwaf.NewAddCustomWhiteRuleResponse()
		domain   string
		ruleIdStr string
	)

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	if v, ok := d.GetOk("sort_id"); ok {
		request.SortId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("expire_time"); ok {
		request.ExpireTime = helper.String(v.(string))
	}

	if v, ok := d.GetOk("strategies"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			strategy := ngwaf.Strategy{}
			if v, ok := dMap["field"]; ok {
				strategy.Field = helper.String(v.(string))
			}

			if v, ok := dMap["compare_func"]; ok {
				strategy.CompareFunc = helper.String(v.(string))
			}

			if v, ok := dMap["content"]; ok {
				strategy.Content = helper.String(v.(string))
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

	if v, ok := d.GetOk("bypass"); ok {
		request.Bypass = helper.String(v.(string))
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

	if v, ok := d.GetOk("logical_op"); ok {
		request.LogicalOp = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().AddCustomWhiteRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create waf CustomWhiteRule failed, reason:%+v", logId, err)
		return err
	}

	ruleId := *response.Response.RuleId
	ruleIdStr = strconv.FormatUint(ruleId, 10)



	d.SetId(strings.Join([]string{domain, ruleIdStr}, FILED_SP))
	return resourceTencentCloudNgwafCustomWhiteRuleRead(d, meta)
}

func resourceTencentCloudNgwafCustomWhiteRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_custom_white_rule.read")()
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

	customWhiteRule, err := service.DescribeWafCustomWhiteRuleById(ctx, domain, ruleId)
	if err != nil {
		return err
	}

	if customWhiteRule == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafCustomWhiteRule` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if customWhiteRule.Name != nil {
		_ = d.Set("name", customWhiteRule.Name)
	}

	if customWhiteRule.SortId != nil {
		_ = d.Set("sort_id", customWhiteRule.SortId)
	}

	if customWhiteRule.ExpireTime != nil {
		_ = d.Set("expire_time", customWhiteRule.ExpireTime)
	}

	if customWhiteRule.Strategies != nil {
		strategiesList := []interface{}{}
		for _, strategies := range customWhiteRule.Strategies {
			strategiesMap := map[string]interface{}{}

			if strategies.Field != nil {
				strategiesMap["field"] = strategies.Field
			}

			if strategies.CompareFunc != nil {
				strategiesMap["compare_func"] = strategies.CompareFunc
			}

			if strategies.Content != nil {
				strategiesMap["content"] = strategies.Content
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

	if customWhiteRule.JobType != nil {
		_ = d.Set("job_type", customWhiteRule.JobType)
	}

	if customWhiteRule.JobDateTime != nil {
		tmpList := make([]map[string]interface{}, 0)
		dMap := map[string]interface{}{}
		if customWhiteRule.JobDateTime.Timed != nil {
			timedList := []interface{}{}
			for _, v := range customWhiteRule.JobDateTime.Timed {
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

		if customWhiteRule.JobDateTime.Cron != nil {
			cronList := []interface{}{}
			for _, v := range customWhiteRule.JobDateTime.Cron {
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

		if customWhiteRule.JobDateTime.TimeTZone != nil {
			dMap["time_t_zone"] = customWhiteRule.JobDateTime.TimeTZone
		}

		tmpList = append(tmpList, dMap)

		_ = d.Set("job_date_time", tmpList)
	}

	if customWhiteRule.Bypass != nil {
		_ = d.Set("bypass", customWhiteRule.Bypass)
	}



	if customWhiteRule.LogicalOp != nil {
		_ = d.Set("logical_op", customWhiteRule.LogicalOp)
	}

	if customWhiteRule.RuleId != nil {
		_ = d.Set("rule_id", customWhiteRule.RuleId)
	}

	return nil
}

func resourceTencentCloudNgwafCustomWhiteRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_custom_white_rule.update")()
	defer inconsistentCheck(d, meta)()

var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyCustomWhiteRuleRequest()
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

	if v, ok := d.GetOk("name"); ok {
		request.RuleName = helper.String(v.(string))
	}

	if v, ok := d.GetOk("bypass"); ok {
		request.Bypass = helper.String(v.(string))
	}

	if v, ok := d.GetOk("sort_id"); ok {
		tmpSortId, _ := strconv.ParseInt(v.(string), 10, 64)
		request.SortId = helper.Int64Uint64(tmpSortId)
	}

	if v, ok := d.GetOk("expire_time"); ok {
		tmpExpireTime, _ := strconv.ParseInt(v.(string), 10, 64)
		request.ExpireTime = helper.Int64Uint64(tmpExpireTime)
	}

	if v, ok := d.GetOk("strategies"); ok {
		for _, item := range v.([]interface{}) {
			dMap := item.(map[string]interface{})
			strategy := ngwaf.Strategy{}
			if v, ok := dMap["field"]; ok {
				strategy.Field = helper.String(v.(string))
			}

			if v, ok := dMap["compare_func"]; ok {
				strategy.CompareFunc = helper.String(v.(string))
			}

			if v, ok := dMap["content"]; ok {
				strategy.Content = helper.String(v.(string))
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

	if v, ok := d.GetOk("logical_op"); ok {
		request.LogicalOp = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyCustomWhiteRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf CustomWhiteRule failed, reason:%+v", logId, err)
		return err
	}



	return resourceTencentCloudNgwafCustomWhiteRuleRead(d, meta)
}

func resourceTencentCloudNgwafCustomWhiteRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_custom_white_rule.delete")()
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

	if err := service.DeleteWafCustomWhiteRuleById(ctx, domain, ruleId); err != nil {
		return err
	}

	return nil
}

/*
Provides a resource to create a NGWAF CC protection rule.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_cc" "basic_protection" {
  domain      = "example.com"
  name        = "basic-cc-protection"
  status      = 1
  advance     = "0"
  limit       = "100"
  interval    = "60"
  url         = "/*"
  match_func  = 1
  action_type = "22"
  priority    = 10
  valid_time  = 300
  edition     = "clb-waf"
}
```

Import

NGWAF CC rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_cc.basic_protection example.com#clb-waf#rule-xxxxxxxx
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

func resourceTencentCloudNgwafCc() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafCcCreate,
		Read:   resourceTencentCloudNgwafCcRead,
		Update: resourceTencentCloudNgwafCcUpdate,
		Delete: resourceTencentCloudNgwafCcDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Rule Name.",
			},
			"status": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Rule Status, 0 rule close, 1 rule open.",
			},
			"advance": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Advanced mode (whether to use session detection). 0(disabled) 1(enabled).",
			},
			"limit": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "CC detection threshold.",
			},
			"interval": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "CC detection cycle.",
			},
			"url": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Detection URL.",
			},
			"match_func": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Match method, 0(equal), 1(prefix), 2(contains), 3(not equal), 6(suffix), 7(not contains).",
			},
			"action_type": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Rule Action, 20 means observation, 21 means human-machine identification, 22 means interception, 23 means precise interception, 26 means precise human-machine identification.",
			},
			"priority": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Rule Priority.",
			},
			"valid_time": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "Action ValidTime, minute unit. Min: 60, Max: 604800.",
			},
			"options_arr": {
				Optional: true,
				Type:     schema.TypeString,
				Description: `CC matching conditions JSON serialized string, example: [{"key":"Method","args":["=R0VU"],"match":"0","encodeflag":true}] 

				Available key values: Method, Post, Referer, Cookie, User-Agent, CustomHeader, CaptchaRisk, CaptchaDeviceRisk, CaptchaScore

				Available match values:
				- When Key is Method: 0 (equal to), 3 (not equal to)
				- When Key is Post: 0 (equal to), 3 (not equal to)
				- When Key is Cookie: 0 (equal to), 2 (contains), 3 (not equal to), 7 (does not contain)
				- When Key is Referer: 0 (equal to), 3 (not equal to), 1 (prefix match), 6 (suffix match), 2 (contains), 7 (does not contain), 12 (exists), 5 (does not exist), 4 (content is empty)
				- When Key is Cookie: 0 (equal to), 3 (not equal to), 2 (contains), 7 (does not contain), 12 (exists), 5 (does not exist), 4 (content is empty)
				- When Key is User-Agent: 0 (equal to), 3 (not equal to), 1 (prefix match), 6 (suffix match), 2 (contains), 7 (does not contain), 12 (exists), 5 (does not exist), 4 (content is empty)
				- When Key is CustomHeader: 0 (equal to), 3 (not equal to), 2 (contains), 7 (does not contain), 12 (exists), 5 (does not exist), 4 (content is empty)
				- When Key is IPLocation: 13 (belongs to), 14 (does not belong to)
				- When Key is CaptchaRisk: 0 (equal to), 3 (not equal to), 13 (belongs to), 14 (does not belong to), 12 (exists), 5 (does not exist)
				- When Key is CaptchaDeviceRisk: 0 (equal to), 3 (not equal to), 13 (belongs to), 14 (does not belong to), 12 (exists), 5 (does not exist)
				- When Key is CaptchaScore: 15 (numerically equal to), 16 (numerically not equal to), 17 (numerically greater than), 18 (numerically less than), 19 (numerically greater than or equal to), 20 (numerically less than or equal to), 12 (exists), 5 (does not exist)

				The args parameter represents matching content and requires encodeflag to be set to true. When Key is Post, Cookie, or CustomHeader, use equals sign = to concatenate Key and Value separately, and encode both with Base64, similar to YWJj=YWJj. When Key is Referer or User-Agent, use equals sign = to concatenate Value, similar to =YWJj.`,
			},
			"edition": {
				Required:     true,
				Type:         schema.TypeString,
				ValidateFunc: validateAllowedStringValue(EDITION_TYPE),
				Description:  "WAF edition. clb-waf means clb-waf, sparta-waf means saas-waf.",
			},
			"type": {
				Optional:    true,
				Type:        schema.TypeInt,
				Description: "Operate Type.",
			},
			"event_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Event ID.",
			},
			"session_applied": {
				Optional:    true,
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeInt},
				Description: "Session ID that needs to be enabled for the rule.",
			},
			"rule_id": {
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Rule ID.",
			},
		},
	}
}

func resourceTencentCloudNgwafCcCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = ngwaf.NewUpsertCCRuleRequest()
		response = ngwaf.NewUpsertCCRuleResponse()
		domain   string
		ruleId   string
		name     string
	)

	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		domain = v.(string)
	}

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
		name = v.(string)
	}

	if v, ok := d.GetOk("status"); ok {
		request.Status = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("advance"); ok {
		request.Advance = helper.String(v.(string))
	}

	if v, ok := d.GetOk("limit"); ok {
		request.Limit = helper.String(v.(string))
	}

	if v, ok := d.GetOk("interval"); ok {
		request.Interval = helper.String(v.(string))
	}

	if v, ok := d.GetOk("url"); ok {
		request.Url = helper.String(v.(string))
	}

	if v, ok := d.GetOk("match_func"); ok {
		request.MatchFunc = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("action_type"); ok {
		request.ActionType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("priority"); ok {
		request.Priority = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("valid_time"); ok {
		request.ValidTime = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("options_arr"); ok {
		request.OptionsArr = helper.String(v.(string))
	}

	if v, ok := d.GetOk("edition"); ok {
		request.Edition = helper.String(v.(string))
	}

	if v, ok := d.GetOk("type"); ok {
		request.Type = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("event_id"); ok {
		request.EventId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("session_applied"); ok {
		sessionAppliedSet := v.(*schema.Set).List()
		for i := range sessionAppliedSet {
			if sessionAppliedSet[i] != nil {
				sessionApplied := sessionAppliedSet[i].(int)
				request.SessionApplied = append(request.SessionApplied, helper.IntInt64(sessionApplied))
			}
		}
	}

	request.RuleId = helper.IntInt64(0)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().UpsertCCRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil || result.Response.RuleId == nil {
			return resource.NonRetryableError(fmt.Errorf("Create waf cc failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create waf cc failed, reason:%+v", logId, err)
		return err
	}

	ruleIdInt := *response.Response.RuleId
	ruleId = strconv.FormatInt(ruleIdInt, 10)
	d.SetId(strings.Join([]string{domain, ruleId, name}, FILED_SP))

	return resourceTencentCloudNgwafCcRead(d, meta)
}

func resourceTencentCloudNgwafCcRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	ruleId := idSplit[1]

	cc, err := service.DescribeWafCcById(ctx, domain, ruleId)
	if err != nil {
		return err
	}

	if cc == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafCc` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)

	if cc.Name != nil {
		_ = d.Set("name", cc.Name)
	}

	if cc.Status != nil {
		_ = d.Set("status", cc.Status)
	}

	if cc.Advance != nil {
		advanceStr := strconv.FormatUint(*cc.Advance, 10)
		_ = d.Set("advance", advanceStr)
	}

	if cc.Limit != nil {
		limitStr := strconv.FormatUint(*cc.Limit, 10)
		_ = d.Set("limit", limitStr)
	}

	if cc.Interval != nil {
		intervalStr := strconv.FormatUint(*cc.Interval, 10)
		_ = d.Set("interval", intervalStr)
	}

	if cc.Url != nil {
		_ = d.Set("url", cc.Url)
	}

	if cc.MatchFunc != nil {
		_ = d.Set("match_func", cc.MatchFunc)
	}

	if cc.ActionType != nil {
		actionTypeStr := strconv.FormatUint(*cc.ActionType, 10)
		_ = d.Set("action_type", actionTypeStr)
	}

	if cc.Priority != nil {
		_ = d.Set("priority", cc.Priority)
	}

	if cc.ValidTime != nil {
		_ = d.Set("valid_time", cc.ValidTime)
	}

	if cc.Options != nil {
		_ = d.Set("options_arr", cc.Options)
	}

	if cc.EventId != nil {
		_ = d.Set("event_id", cc.EventId)
	}

	if cc.SessionApplied != nil {
		_ = d.Set("session_applied", cc.SessionApplied)
	}

	if cc.RuleId != nil {
		ruleIdStr := strconv.FormatUint(*cc.RuleId, 10)
		_ = d.Set("rule_id", ruleIdStr)
	}

	// Update ID if name has changed to keep it consistent
	if cc.Name != nil {
		currentName := *cc.Name
		newId := strings.Join([]string{domain, ruleId, currentName}, FILED_SP)
		d.SetId(newId)
	}

	return nil
}

func resourceTencentCloudNgwafCcUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewUpsertCCRuleRequest()
	)

	immutableArgs := []string{"domain"}
	for _, v := range immutableArgs {
		if d.HasChange(v) {
			return fmt.Errorf("argument `%s` cannot be changed", v)
		}
	}

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	ruleId := idSplit[1]
	// name from ID is the old name, use current value instead
	oldName := idSplit[2]

	request.Domain = &domain
	ruleIdInt, _ := strconv.ParseInt(ruleId, 10, 64)
	request.RuleId = &ruleIdInt
	
	// Use current name value from configuration
	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	} else {
		request.Name = &oldName
	}

	if v, ok := d.GetOk("status"); ok {
		request.Status = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("advance"); ok {
		request.Advance = helper.String(v.(string))
	}

	if v, ok := d.GetOk("limit"); ok {
		request.Limit = helper.String(v.(string))
	}

	if v, ok := d.GetOk("interval"); ok {
		request.Interval = helper.String(v.(string))
	}

	if v, ok := d.GetOk("url"); ok {
		request.Url = helper.String(v.(string))
	}

	if v, ok := d.GetOk("match_func"); ok {
		request.MatchFunc = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("action_type"); ok {
		request.ActionType = helper.String(v.(string))
	}

	if v, ok := d.GetOk("priority"); ok {
		request.Priority = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("valid_time"); ok {
		request.ValidTime = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("options_arr"); ok {
		request.OptionsArr = helper.String(v.(string))
	}

	if v, ok := d.GetOk("edition"); ok {
		request.Edition = helper.String(v.(string))
	}

	if v, ok := d.GetOk("type"); ok {
		request.Type = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("event_id"); ok {
		request.EventId = helper.String(v.(string))
	}

	if v, ok := d.GetOk("session_applied"); ok {
		sessionAppliedSet := v.(*schema.Set).List()
		for i := range sessionAppliedSet {
			sessionApplied := sessionAppliedSet[i].(int)
			request.SessionApplied = append(request.SessionApplied, helper.IntInt64(sessionApplied))
		}
	}

	if v, ok := d.GetOk("limit_method"); ok {
		request.LimitMethod = helper.String(v.(string))
	}

	if v, ok := d.GetOk("cel_rule"); ok {
		request.CelRule = helper.String(v.(string))
	}

	if v, ok := d.GetOk("logical_op"); ok {
		request.LogicalOp = helper.String(v.(string))
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().UpsertCCRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s update waf cc failed, reason:%+v", logId, err)
		return err
	}

	return resourceTencentCloudNgwafCcRead(d, meta)
}

func resourceTencentCloudNgwafCcDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_cc.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken,%s", idSplit)
	}
	domain := idSplit[0]
	ruleId := idSplit[1]
	name := idSplit[2]

	if err := service.DeleteWafCcById(ctx, domain, ruleId, name); err != nil {
		return err
	}

	return nil
}

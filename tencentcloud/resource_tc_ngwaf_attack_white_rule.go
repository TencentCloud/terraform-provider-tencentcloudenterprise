/*
Provides a resource to create a NGWAF attack white rule.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_attack_white_rule" "example" {
  domain = "example.com"
  status = 1

  rules {
    match_field   = "URI"
    match_method  = "equal"
    match_content = "/api/health"
  }

  signature_ids = ["rule-100001"]
  mode          = 0
  name          = "example-white-rule"
}
```

Import

NGWAF attack white rule can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_attack_white_rule.example example.com#rule_id
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

func resourceTencentCloudNgwafAttackWhiteRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafAttackWhiteRuleCreate,
		Read:   resourceTencentCloudNgwafAttackWhiteRuleRead,
		Update: resourceTencentCloudNgwafAttackWhiteRuleUpdate,
		Delete: resourceTencentCloudNgwafAttackWhiteRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain.",
			},

			"status": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Rule status.",
			},

			"rules": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Rule list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"match_field": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Matching domains.",
						},
						"match_method": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Matching method.",
						},
						"match_content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Matching content.",
						},
						"match_params": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Matching params.",
						},
					},
				},
			},

			"signature_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Whitelist of rule IDs.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"type_ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "The whitened category rule ID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"mode": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "0: Whiten according to a specific rule ID, 1: Whiten according to the rule type.",
			},

			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Rule name.",
			},

			// computed
			"rule_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Rule ID.",
			},
		},
	}
}

func resourceTencentCloudNgwafAttackWhiteRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_attack_white_rule.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId    = getLogId(contextNil)
		request  = ngwaf.NewAddAttackWhiteRuleRequest()
		response = ngwaf.NewAddAttackWhiteRuleResponse()
		domain   string
		ruleId   string
	)

	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		domain = v.(string)
	}

	if v, ok := d.GetOkExists("status"); ok {
		request.Status = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("rules"); ok {
		for _, item := range v.([]interface{}) {
			rulesMap := item.(map[string]interface{})
			userWhiteRuleItem := ngwaf.UserWhiteRuleItem{}
			if v, ok := rulesMap["match_field"].(string); ok && v != "" {
				userWhiteRuleItem.MatchField = helper.String(v)
			}

			if v, ok := rulesMap["match_method"].(string); ok && v != "" {
				userWhiteRuleItem.MatchMethod = helper.String(v)
			}

			if v, ok := rulesMap["match_content"].(string); ok && v != "" {
				userWhiteRuleItem.MatchContent = helper.String(v)
			}

			if v, ok := rulesMap["match_params"].(string); ok && v != "" {
				userWhiteRuleItem.MatchParams = helper.String(v)
			}

			request.Rules = append(request.Rules, &userWhiteRuleItem)
		}
	}

	if v, ok := d.GetOk("signature_ids"); ok {
		signatureIdsSet := v.(*schema.Set).List()
		for i := range signatureIdsSet {
			signatureIds := signatureIdsSet[i].(string)
			request.SignatureIds = append(request.SignatureIds, helper.String(signatureIds))
		}
	}

	if v, ok := d.GetOk("type_ids"); ok {
		typeIdsSet := v.(*schema.Set).List()
		for i := range typeIdsSet {
			typeIds := typeIdsSet[i].(string)
			request.TypeIds = append(request.TypeIds, helper.String(typeIds))
		}
	}

	if v, ok := d.GetOkExists("mode"); ok {
		request.Mode = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().AddAttackWhiteRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		if result == nil || result.Response == nil {
			return resource.NonRetryableError(fmt.Errorf("Create waf attack white rule failed, Response is nil."))
		}

		response = result
		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s create waf attack white rule failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	if response.Response.RuleId == nil {
		return fmt.Errorf("RuleId is nil.")
	}

	ruleId = strconv.FormatUint(*response.Response.RuleId, 10)
	d.SetId(strings.Join([]string{domain, ruleId}, FILED_SP))
	return resourceTencentCloudNgwafAttackWhiteRuleRead(d, meta)
}

func resourceTencentCloudNgwafAttackWhiteRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_attack_white_rule.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	domain := idSplit[0]
	ruleId := idSplit[1]

	ruleIdInt := helper.StrToUInt64(ruleId)
	respData, err := service.DescribeWafAttackWhiteRuleById(ctx, domain, ruleIdInt)
	if err != nil {
		return err
	}

	if respData == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `waf_attack_white_rule` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)
	_ = d.Set("rule_id", ruleIdInt)

	if respData.Status != nil {
		_ = d.Set("status", respData.Status)
	}

	if respData.MatchInfo != nil {
		tmpList := make([]map[string]interface{}, 0, len(respData.MatchInfo))
		for _, item := range respData.MatchInfo {
			dMap := make(map[string]interface{}, 0)
			if item.MatchField != nil {
				dMap["match_field"] = item.MatchField
			}

			if item.MatchMethod != nil {
				dMap["match_method"] = item.MatchMethod
			}

			if item.MatchContent != nil {
				dMap["match_content"] = item.MatchContent
			}

			if item.MatchParams != nil {
				dMap["match_params"] = item.MatchParams
			}

			tmpList = append(tmpList, dMap)
		}

		_ = d.Set("rules", tmpList)
	}

	if respData.SignatureIds != nil {
		_ = d.Set("signature_ids", respData.SignatureIds)
	}

	if respData.TypeIds != nil {
		_ = d.Set("type_ids", respData.TypeIds)
	}

	if respData.Mode != nil {
		_ = d.Set("mode", respData.Mode)
	}

	if respData.Name != nil {
		_ = d.Set("name", respData.Name)
	}

	return nil
}

func resourceTencentCloudNgwafAttackWhiteRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_attack_white_rule.update")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyAttackWhiteRuleRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	domain := idSplit[0]
	ruleId := idSplit[1]

	if v, ok := d.GetOkExists("status"); ok {
		request.Status = helper.IntUint64(v.(int))
	}

	if v, ok := d.GetOk("rules"); ok {
		for _, item := range v.([]interface{}) {
			rulesMap := item.(map[string]interface{})
			userWhiteRuleItem := ngwaf.UserWhiteRuleItem{}
			if v, ok := rulesMap["match_field"].(string); ok && v != "" {
				userWhiteRuleItem.MatchField = helper.String(v)
			}

			if v, ok := rulesMap["match_method"].(string); ok && v != "" {
				userWhiteRuleItem.MatchMethod = helper.String(v)
			}

			if v, ok := rulesMap["match_content"].(string); ok && v != "" {
				userWhiteRuleItem.MatchContent = helper.String(v)
			}

			if v, ok := rulesMap["match_params"].(string); ok && v != "" {
				userWhiteRuleItem.MatchParams = helper.String(v)
			}

			request.Rules = append(request.Rules, &userWhiteRuleItem)
		}
	}

	if v, ok := d.GetOk("signature_ids"); ok {
		signatureIdsSet := v.(*schema.Set).List()
		for i := range signatureIdsSet {
			signatureIds := signatureIdsSet[i].(string)
			request.SignatureIds = append(request.SignatureIds, helper.String(signatureIds))
		}
	}

	if v, ok := d.GetOk("type_ids"); ok {
		typeIdsSet := v.(*schema.Set).List()
		for i := range typeIdsSet {
			typeIds := typeIdsSet[i].(string)
			request.TypeIds = append(request.TypeIds, helper.String(typeIds))
		}
	}

	if v, ok := d.GetOkExists("mode"); ok {
		request.Mode = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOk("name"); ok {
		request.Name = helper.String(v.(string))
	}

	request.Domain = &domain
	request.RuleId = helper.StrToUint64Point(ruleId)
	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyAttackWhiteRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s update waf attack white rule failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return resourceTencentCloudNgwafAttackWhiteRuleRead(d, meta)
}

func resourceTencentCloudNgwafAttackWhiteRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_attack_white_rule.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewDeleteAttackWhiteRuleRequest()
	)

	idSplit := strings.Split(d.Id(), FILED_SP)
	if len(idSplit) != 2 {
		return fmt.Errorf("id is broken,%s", d.Id())
	}

	domain := idSplit[0]
	ruleId := idSplit[1]

	request.Domain = &domain
	request.Ids = []*uint64{helper.StrToUint64Point(ruleId)}
	reqErr := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().DeleteAttackWhiteRule(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}

		return nil
	})

	if reqErr != nil {
		log.Printf("[CRITAL]%s delete waf attack white rule failed, reason:%+v", logId, reqErr)
		return reqErr
	}

	return nil
}

/*
Provides a resource to create NGWAF auto deny rules.

Example Usage

```hcl
resource "tencentcloudenterprise_ngwaf_auto_deny_rules" "example" {
  domain               = "example.com"
  attack_threshold     = 10
  time_threshold       = 10
  deny_time_threshold  = 30
}
```

Import

NGWAF auto deny rules can be imported using the id, e.g.

```
$ terraform import tencentcloudenterprise_ngwaf_auto_deny_rules.example example.com
```
*/
package tencentcloud

import (
	"context"
	"log"

	ngwaf "terraform-provider-tencentcloudenterprise/sdk/ngwaf/v20180125"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceTencentCloudNgwafAutoDenyRules() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudNgwafAutoDenyRulesCreate,
		Read:   resourceTencentCloudNgwafAutoDenyRulesRead,
		Delete: resourceTencentCloudNgwafAutoDenyRulesDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "Domain.",
			},
			"attack_threshold": {
				Required:     true,
				ForceNew:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateIntegerInRange(2, 100),
				Description:  "The threshold number of attacks that triggers IP autodeny, ranging from 2 to 100 times.",
			},
			"time_threshold": {
				Required:     true,
				ForceNew:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateIntegerInRange(1, 60),
				Description:  "IP autodeny statistical time, ranging from 1-60 minutes.",
			},
			"deny_time_threshold": {
				Required:     true,
				ForceNew:     true,
				Type:         schema.TypeInt,
				ValidateFunc: validateIntegerInRange(5, 360),
				Description:  "The IP autodeny time after triggering the IP autodeny, ranging from 5 to 360 minutes.",
			},
		},
	}
}

func resourceTencentCloudNgwafAutoDenyRulesCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_auto_deny_rules.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyWafAutoDenyRulesRequest()
		domain  string
	)

	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
		domain = v.(string)
	}

	if v, ok := d.GetOkExists("attack_threshold"); ok {
		request.AttackThreshold = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("time_threshold"); ok {
		request.TimeThreshold = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("deny_time_threshold"); ok {
		request.DenyTimeThreshold = helper.IntInt64(v.(int))
	}

	request.DefenseStatus = helper.IntInt64(1)

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyWafAutoDenyRules(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create waf autoDenyRules failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(domain)

	return resourceTencentCloudNgwafAutoDenyRulesRead(d, meta)
}

func resourceTencentCloudNgwafAutoDenyRulesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_auto_deny_rules.read")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = &NgwafService{client: meta.(*TencentCloudClient)}
		domain  = d.Id()
	)

	autoDenyRules, err := service.DescribeWafAutoDenyRulesById(ctx, domain)
	if err != nil {
		return err
	}

	if autoDenyRules == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `WafAutoDenyRules` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	_ = d.Set("domain", domain)

	if autoDenyRules.Response != nil && autoDenyRules.Response.AttackThreshold != nil {
		_ = d.Set("attack_threshold", autoDenyRules.Response.AttackThreshold)
	}

	if autoDenyRules.Response != nil && autoDenyRules.Response.TimeThreshold != nil {
		_ = d.Set("time_threshold", autoDenyRules.Response.TimeThreshold)
	}

	if autoDenyRules.Response != nil && autoDenyRules.Response.DenyTimeThreshold != nil {
		_ = d.Set("deny_time_threshold", autoDenyRules.Response.DenyTimeThreshold)
	}

	return nil
}

func resourceTencentCloudNgwafAutoDenyRulesDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_ngwaf_auto_deny_rules.delete")()
	defer inconsistentCheck(d, meta)()

	var (
		logId   = getLogId(contextNil)
		request = ngwaf.NewModifyWafAutoDenyRulesRequest()
		domain  = d.Id()
	)

	if v, ok := d.GetOkExists("attack_threshold"); ok {
		request.AttackThreshold = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("time_threshold"); ok {
		request.TimeThreshold = helper.IntInt64(v.(int))
	}

	if v, ok := d.GetOkExists("deny_time_threshold"); ok {
		request.DenyTimeThreshold = helper.IntInt64(v.(int))
	}

	request.Domain = &domain
	request.DefenseStatus = helper.IntInt64(0)
	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseNgwafClient().ModifyWafAutoDenyRules(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s delete waf autoDenyRules failed, reason:%+v", logId, err)
		return err
	}

	return nil
}

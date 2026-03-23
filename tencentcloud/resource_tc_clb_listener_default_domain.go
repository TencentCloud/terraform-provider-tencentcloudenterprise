/*
Provides a resource to set CLB listener default domain.

Example Usage

```hcl
resource "tencentcloudenterprise_clb_listener_default_domain" "example" {
  clb_id      = "lb-7a0t6zqb"
  listener_id = "lbl-hh141sn9"
  domain      = "www.example.com"
}
```

Import

CLB listener default domain can be imported using the id (clb_id#listener_id), e.g.

```
$ terraform import tencentcloudenterprise_clb_listener_default_domain.example lb-7a0t6zqb#lbl-hh141sn9
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	clb "terraform-provider-tencentcloudenterprise/sdk/clb/v20180317"
	sdkErrors "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_clb_listener_default_domain", CNDescription{
		TerraformTypeCN: "CLB 默认域名",
		DescriptionCN:   "提供 CLB 监听器默认域名资源，用于设置 CLB 七层监听器的默认域名。",
		AttributesCN: map[string]string{
			"clb_id":      "CLB 实例 ID",
			"listener_id": "CLB 监听器 ID",
			"domain":      "监听器规则的域名，单域名规则传入 domain，多域名规则传入 domains",
			"rule_id":     "CLB 监听器规则 ID",
		},
	})
}

func resourceTencentCloudClbListenerDefaultDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceTencentCloudClbListenerDefaultDomainCreate,
		Read:   resourceTencentCloudClbListenerDefaultDomainRead,
		Update: resourceTencentCloudClbListenerDefaultDomainUpdate,
		Delete: resourceTencentCloudClbListenerDefaultDomainDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"clb_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of CLB instance.",
			},
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "ID of CLB listener.",
			},
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name of the listener rule. Single domain rules are passed to `domain`, and multi domain rules are passed to `domains`.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "ID of this CLB listener rule.",
			},
		},
	}
}

func resourceTencentCloudClbListenerDefaultDomainCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_listener_default_domain.create")()
	defer inconsistentCheck(d, meta)()

	var (
		logId      = getLogId(contextNil)
		request    = clb.NewModifyDomainAttributesRequest()
		response   *clb.ModifyDomainAttributesResponse
		clbId      string
		listenerId string
	)

	if v, ok := d.GetOk("clb_id"); ok {
		clbId = v.(string)
		request.LoadBalancerId = helper.String(clbId)
	}

	if v, ok := d.GetOk("listener_id"); ok {
		listenerId = v.(string)
		request.ListenerId = helper.String(listenerId)
	}

	if v, ok := d.GetOk("domain"); ok {
		request.Domain = helper.String(v.(string))
	}

	request.DefaultServer = helper.Bool(true)

	client := meta.(*TencentCloudClient).apiV3Conn.UseClbClient()

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := client.ModifyDomainAttributes(request)
		if e != nil {
			if sdkError, ok := e.(*sdkErrors.CloudSDKError); ok {
				if sdkError.Code == "FailedOperation.ResourceInOperating" {
					return resource.RetryableError(e)
				}
			}

			return retryError(e)
		}
		log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
			logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())

		if result == nil || result.Response == nil || result.Response.RequestId == nil {
			return resource.NonRetryableError(fmt.Errorf("Modify domain attributes failed, Response is nil."))
		}

		response = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s create clb listener default domain failed, reason:%+v", logId, err)
		return err
	}

	taskId := *response.Response.RequestId
	if retryErr := waitForTaskFinish(taskId, client); retryErr != nil {
		return retryErr
	}

	d.SetId(clbId + FILED_SP + listenerId)
	return resourceTencentCloudClbListenerDefaultDomainRead(d, meta)
}

func resourceTencentCloudClbListenerDefaultDomainRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_listener_default_domain.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	resourceId := d.Id()

	items := strings.Split(resourceId, FILED_SP)
	if len(items) != 2 {
		return fmt.Errorf("id is broken,%s", resourceId)
	}
	clbId := items[0]
	listenerId := items[1]

	clbService := ClbService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}

	filter := map[string]string{"listener_id": listenerId, "clb_id": clbId}
	var instances []*clb.RuleOutput
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := clbService.DescribeRulesByFilter(ctx, filter)
		if e != nil {
			return retryError(e)
		}
		instances = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CLB listener rule failed, reason:%+v", logId, err)
		return err
	}

	if len(instances) == 0 {
		d.SetId("")
		return nil
	}

	var (
		domain string
		ruleId string
	)

	for _, rule := range instances {
		if rule.DefaultServer != nil && *rule.DefaultServer {
			domain = *rule.Domain
			ruleId = *rule.LocationId
			break
		}
	}

	_ = d.Set("clb_id", clbId)
	_ = d.Set("listener_id", listenerId)
	_ = d.Set("domain", domain)
	_ = d.Set("rule_id", ruleId)

	return nil
}

func resourceTencentCloudClbListenerDefaultDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_listener_default_domain.update")()
	defer inconsistentCheck(d, meta)()

	if d.HasChange("domain") {
		var (
			logId    = getLogId(contextNil)
			request  = clb.NewModifyDomainAttributesRequest()
			response *clb.ModifyDomainAttributesResponse
		)

		if v, ok := d.GetOk("clb_id"); ok {
			request.LoadBalancerId = helper.String(v.(string))
		}

		if v, ok := d.GetOk("listener_id"); ok {
			request.ListenerId = helper.String(v.(string))
		}

		if v, ok := d.GetOk("domain"); ok {
			request.Domain = helper.String(v.(string))
		}

		request.DefaultServer = helper.Bool(true)

		client := meta.(*TencentCloudClient).apiV3Conn.UseClbClient()

		err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
			result, e := client.ModifyDomainAttributes(request)
			if e != nil {
				if sdkError, ok := e.(*sdkErrors.CloudSDKError); ok {
					if sdkError.Code == "FailedOperation.ResourceInOperating" {
						return resource.RetryableError(e)
					}
				}

				return retryError(e)
			}
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n",
				logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())

			if result == nil || result.Response == nil || result.Response.RequestId == nil {
				return resource.NonRetryableError(fmt.Errorf("Modify domain attributes failed, Response is nil."))
			}

			response = result
			return nil
		})

		if err != nil {
			log.Printf("[CRITAL]%s update clb listener default domain failed, reason:%+v", logId, err)
			return err
		}

		taskId := *response.Response.RequestId
		if retryErr := waitForTaskFinish(taskId, client); retryErr != nil {
			return retryErr
		}
	}

	return resourceTencentCloudClbListenerDefaultDomainRead(d, meta)
}

func resourceTencentCloudClbListenerDefaultDomainDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_clb_listener_default_domain.delete")()
	defer inconsistentCheck(d, meta)()

	return nil
}

/*
Provide a resource to create a VPCDNS domain forward rule.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpcdns_forward_rule" "foo" {
	  remark     = "forward_rule_foo"
	  domain_id = "my_domain_id1"
	  forward_address = ["8.8.8.8:88", "1.1.1.1:88"]
	}

```

# Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import cloud_vpcdns_forward_rule.test remark
```
*/
package tencentcloud

import (
	"context"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkError "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	"time"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpcdns_forward_rule", CNDescription{
		TerraformTypeCN: "VPCDNS转发规则",
		DescriptionCN:   "提供VPCDNS转发规则资源，用于创建和管理DNS转发规则。",
		AttributesCN: map[string]string{
			"remark":          "转发规则名称",
			"domain_id":       "转发域名id",
			"forward_address": "dns地址",
			"create_time":     "创建时间",
			"rule_id":         "转发规则id",
		},
	})
}

func resourceTencentCloudVpcDnsForwardRule() *schema.Resource {
	return &schema.Resource{
		Description: "Provide a resource to create a VPCDNS forward rule.",
		Create:      resourceTencentCloudVpcDnsForwardRuleCreate,
		Read:        resourceTencentCloudVpcDnsForwardRuleRead,
		Update:      resourceTencentCloudVpcDnsForwardRuleUpdate,
		Delete:      resourceTencentCloudVpcDnsForwardRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"remark": {
				Type:        schema.TypeString,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The remark of the forward rule.",
			},
			"domain_id": {
				Type:        schema.TypeString,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The domain IDs of the forward rule.",
			},
			"forward_address": {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The forward address of the rule.",
			},

			// Computed values
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Creation time of VpcDns Forward Rule.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rule ID of the forward rule.",
			},
		},
	}
}

func resourceTencentCloudVpcDnsForwardRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpcdns_forward_rule.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	vpcDnsService := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		remark         string
		domainId       string
		forwardAddress []string
	)
	if temp, ok := d.GetOk("remark"); ok {
		remark = temp.(string)
	}
	if temp, ok := d.GetOk("domain_id"); ok {
		domainId = temp.(string)
	}

	if temp, ok := d.GetOk("forward_address"); ok {
		// 将[]interface{}转换为[]string
		if list, ok := temp.([]interface{}); ok {
			forwardAddress = make([]string, len(list))
			for i, v := range list {
				forwardAddress[i] = v.(string)
			}
		}
	}
	ruleId, err := vpcDnsService.CreateVpcDnsForwardRule(ctx, remark, domainId, forwardAddress)
	if err != nil {
		return err
	}

	time.Sleep(waitReadTimeout)
	d.SetId(ruleId)

	return resourceTencentCloudVpcDnsForwardRuleRead(d, meta)
}

func resourceTencentCloudVpcDnsForwardRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpcdns_forward_rule.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	// only filter by DomainId and DomainName
	domainId := d.Get("domain_id").(string)
	filterMap := map[string][]string{
		"DomainId": {domainId},
	}

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		forwardRule, e := service.DescribeVpcDnsForwardRuleList(ctx, filterMap)
		if e != nil {
			return retryError(e)
		}

		if forwardRule == nil {
			return retryError(errors.New("vpc dns forward rule not found"))
		}
		_ = d.Set("rule_id", forwardRule.RuleId)
		_ = d.Set("remark", forwardRule.Remark)
		_ = d.Set("domain_id", forwardRule.DomainId)
		_ = d.Set("forward_address", forwardRule.ForwardAddress)
		return nil
	})

	return err
}

func resourceTencentCloudVpcDnsForwardRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpcdns_forward_rule.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	id := d.Id()

	vpcDnsService := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	d.Partial(true)

	var (
		remark         string
		forwardAddress []string
	)

	old, now := d.GetChange("remark")
	if d.HasChange("remark") {
		remark = now.(string)
	} else {
		remark = old.(string)
	}

	old, now = d.GetChange("forward_address")
	if d.HasChange("forward_address") {
		if list, ok := now.([]interface{}); ok {
			forwardAddress = make([]string, len(list))
			for i, v := range list {
				forwardAddress[i] = v.(string)
			}
		}
	} else {
		if list, ok := old.([]interface{}); ok {
			forwardAddress = make([]string, len(list))
			for i, v := range list {
				forwardAddress[i] = v.(string)
			}
		}
	}

	if err := vpcDnsService.ModifyVpcDnsForwardRule(ctx, id, remark, forwardAddress); err != nil {
		return err
	}

	d.Partial(false)

	return resourceTencentCloudVpcDnsForwardRuleRead(d, meta)
}

func resourceTencentCloudVpcDnsForwardRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpcdns_forward_rule.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		if err := service.DeleteVpcDnsForwardRule(ctx, d.Id()); err != nil {
			if sdkErr, ok := err.(*sdkError.CloudSDKError); ok {
				if sdkErr.Code == VPCNotFound {
					return nil
				}
			}
			return resource.RetryableError(err)
		}
		return nil
	})

	return err
}

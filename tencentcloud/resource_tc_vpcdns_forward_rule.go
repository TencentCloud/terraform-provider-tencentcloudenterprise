/*
Provide a resource to create a VPCDNS forward rule.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpcdns_forward_rule" "foo" {
	  remark          = "forward_rule_foo"
	  zone_id         = tencentcloudenterprise_vpcdns_zone.zone.id
	  forward_address = ["8.8.8.8:88", "1.1.1.1:88"]
	}

```

# Import

Vpcdns forward rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpcdns_forward_rule.foo rule_id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"

	sdkError "terraform-provider-tencentcloudenterprise/sdk/common/errors"
	vpcdns "terraform-provider-tencentcloudenterprise/sdk/vpcdns/v20191025"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpcdns_forward_rule", CNDescription{
		TerraformTypeCN: "VPCDNS转发规则",
		DescriptionCN:   "提供VPCDNS转发规则资源，用于创建和管理DNS转发规则。",
		AttributesCN: map[string]string{
			"remark":          "转发规则名称",
			"zone_id":         "私有域ID",
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
				Description: "The remark of the forward rule.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Private zone ID, e.g. zone-xxxxxxxx.",
			},
			"forward_address": {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "The forward address of the rule, e.g. 8.8.8.8:53.",
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
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_forward_rule.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	vpcDnsService := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		remark         string
		zoneId         string
		forwardAddress []string
	)
	if temp, ok := d.GetOk("remark"); ok {
		remark = temp.(string)
	}
	if temp, ok := d.GetOk("zone_id"); ok {
		zoneId = temp.(string)
	}

	if temp, ok := d.GetOk("forward_address"); ok {
		if list, ok := temp.([]interface{}); ok {
			forwardAddress = make([]string, len(list))
			for i, v := range list {
				forwardAddress[i] = v.(string)
			}
		}
	}

	// Resolve zone_id to domain_id via DescribePrivateZone
	request := vpcdns.NewDescribePrivateZoneRequest()
	request.ZoneId = &zoneId
	response, err := meta.(*TencentCloudClient).apiV3Conn.UseVpcDnsClient().DescribePrivateZone(request)
	if err != nil {
		return fmt.Errorf("failed to describe private zone %s: %v", zoneId, err)
	}
	if response.Response == nil || response.Response.PrivateZone == nil || response.Response.PrivateZone.DomainId == nil {
		return fmt.Errorf("private zone %s has no DomainId", zoneId)
	}
	domainId := strconv.FormatInt(*response.Response.PrivateZone.DomainId, 10)
	log.Printf("[DEBUG]%s resolved zone_id %s to domain_id %s", logId, zoneId, domainId)

	ruleId, err := vpcDnsService.CreateVpcDnsForwardRule(ctx, remark, domainId, forwardAddress)
	if err != nil {
		return err
	}

	d.SetId(ruleId)

	return resourceTencentCloudVpcDnsForwardRuleRead(d, meta)
}

func resourceTencentCloudVpcDnsForwardRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_forward_rule.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcDnsService{client: meta.(*TencentCloudClient).apiV3Conn}

	ruleId := d.Id()

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		forwardRule, e := service.DescribeVpcDnsForwardRuleById(ctx, ruleId)
		if e != nil {
			return retryError(e)
		}

		if forwardRule == nil {
			d.SetId("")
			return nil
		}
		_ = d.Set("rule_id", forwardRule.RuleId)
		_ = d.Set("remark", forwardRule.Remark)
		_ = d.Set("forward_address", forwardRule.ForwardAddress)
		return nil
	})

	return err
}

func resourceTencentCloudVpcDnsForwardRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_forward_rule.update")()

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
	defer logElapsed("resource.tencentcloudenterprise_vpcdns_forward_rule.delete")()

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

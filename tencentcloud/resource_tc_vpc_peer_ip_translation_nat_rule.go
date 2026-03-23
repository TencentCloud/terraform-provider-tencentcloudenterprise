/*
Provides a resource to creating VPC peer IP translation NAT rule.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc" "main" {
	  name       = "ci-vpc-instance-test"
	  cidr_block = "10.0.0.0/16"
	}

	resource "tencentcloudenterprise_vpc_dc_gateway" "dcg_main" {
	  name                = "ci-dcg-test"
	  network_instance_id = tencentcloudenterprise_vpc.main.id
	  network_type        = "VPC"
	  gateway_type        = "NAT"
	}

	resource "tencentcloudenterprise_vpc_peer_ip_translation_nat_rule" "main" {
	  vpc_id                    = tencentcloudenterprise_vpc.main.id
	  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
	  original_ip               = "10.0.1.100"
	  translation_ip            = "192.168.1.100"
	  description               = "IDC侧IP转换NAT规则"
	}

```

# Import

VPC peer IP translation NAT rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_peer_ip_translation_nat_rule.instance vpc-id#dcg-id#original-ip#translation-ip
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
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_peer_ip_translation_nat_rule", CNDescription{
		TerraformTypeCN: "IDC侧IP转换NAT规则",
		DescriptionCN:   "提供IDC侧IP转换NAT规则资源，用于创建和管理IDC侧IP转换NAT规则。",
		AttributesCN: map[string]string{
			"vpc_id":                    "VPC实例ID",
			"direct_connect_gateway_id": "专线网关ID",
			"original_ip":               "原始IP地址",
			"translation_ip":            "映射IP地址",
			"description":               "规则描述",
		},
	})
}

func resourceTencentCloudVpcPeerIpTranslationNatRule() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to creating VPC peer IP translation NAT rule.",
		Create:      resourceTencentCloudVpcPeerIpTranslationNatRuleCreate,
		Read:        resourceTencentCloudVpcPeerIpTranslationNatRuleRead,
		Update:      resourceTencentCloudVpcPeerIpTranslationNatRuleUpdate,
		Delete:      resourceTencentCloudVpcPeerIpTranslationNatRuleDelete,
		Importer: &schema.ResourceImporter{
			State: resourceTencentCloudVpcPeerIpTranslationNatRuleImport,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPC ID.",
			},
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Direct connect gateway ID.",
			},
			"original_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Original IP address.",
			},
			"translation_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Translation IP address.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the peer IP translation NAT rule.",
			},
		},
	}
}

func resourceTencentCloudVpcPeerIpTranslationNatRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_peer_ip_translation_nat_rule.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		vpcId                  = d.Get("vpc_id").(string)
		directConnectGatewayId = d.Get("direct_connect_gateway_id").(string)
		originalIp             = d.Get("original_ip").(string)
		translationIp          = d.Get("translation_ip").(string)
		description            = d.Get("description").(string)
	)

	err := service.CreatePeerIpTranslationNatRule(ctx, vpcId, directConnectGatewayId, originalIp, translationIp, description)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s#%s#%s#%s", vpcId, directConnectGatewayId, originalIp, translationIp))

	return resourceTencentCloudVpcPeerIpTranslationNatRuleRead(d, meta)
}

func resourceTencentCloudVpcPeerIpTranslationNatRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_peer_ip_translation_nat_rule.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	originalIp := idSplit[2]
	translationIp := idSplit[3]

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		rule, e := service.DescribePeerIpTranslationNatRule(ctx, vpcId, directConnectGatewayId, originalIp, translationIp)
		if e != nil {
			return retryError(e)
		}

		if rule == nil {
			log.Printf("[WARN] VPC Peer IP Translation NAT Rule %s not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		_ = d.Set("vpc_id", vpcId)
		_ = d.Set("direct_connect_gateway_id", directConnectGatewayId)
		if rule.OriginalIp != nil {
			_ = d.Set("original_ip", *rule.OriginalIp)
		}
		if rule.TranslationIp != nil {
			_ = d.Set("translation_ip", *rule.TranslationIp)
		}
		if rule.Description != nil {
			_ = d.Set("description", *rule.Description)
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudVpcPeerIpTranslationNatRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_peer_ip_translation_nat_rule.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	oldOriginalIp := idSplit[2]
	oldTranslationIp := idSplit[3]

	var (
		originalIp    = d.Get("original_ip").(string)
		translationIp = d.Get("translation_ip").(string)
		description   = d.Get("description").(string)
	)

	err := service.ModifyPeerIpTranslationNatRule(ctx, vpcId, directConnectGatewayId, oldOriginalIp, oldTranslationIp, originalIp, translationIp, description)
	if err != nil {
		return err
	}

	// If IP addresses changed, update resource ID
	if originalIp != oldOriginalIp || translationIp != oldTranslationIp {
		d.SetId(fmt.Sprintf("%s#%s#%s#%s", vpcId, directConnectGatewayId, originalIp, translationIp))
	}

	return resourceTencentCloudVpcPeerIpTranslationNatRuleRead(d, meta)
}

func resourceTencentCloudVpcPeerIpTranslationNatRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_peer_ip_translation_nat_rule.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	originalIp := idSplit[2]
	translationIp := idSplit[3]

	err := service.DeletePeerIpTranslationNatRule(ctx, vpcId, directConnectGatewayId, originalIp, translationIp)
	if err != nil {
		return err
	}

	return nil
}

func resourceTencentCloudVpcPeerIpTranslationNatRuleImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 4 {
		return nil, fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	originalIp := idSplit[2]
	translationIp := idSplit[3]

	_ = d.Set("vpc_id", vpcId)
	_ = d.Set("direct_connect_gateway_id", directConnectGatewayId)
	_ = d.Set("original_ip", originalIp)
	_ = d.Set("translation_ip", translationIp)

	return []*schema.ResourceData{d}, nil
}

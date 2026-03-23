/*
Provides a resource to creating VPC local source IP port translation ACL rule.

Example Usage

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

	resource "tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule" "nat_rule" {
	  vpc_id                    = tencentcloudenterprise_vpc.main.id
	  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
	  ip_pool                   = "10.40.31.45"
	  description               = "test nat rule"
	}

	resource "tencentcloudenterprise_vpc_local_source_ip_port_translation_acl_rule" "main" {
	  vpc_id                    = tencentcloudenterprise_vpc.main.id
	  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
	  translation_ip_pool       = tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule.nat_rule.ip_pool
	  protocol                  = "all"
	  source_port               = "0"
	  source_cidr               = "10.33.44.55"
	  destination_port          = "0"
	  destination_cidr          = "10.44.33.11"
	}

```

Import

VPC local source IP port translation ACL rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_local_source_ip_port_translation_acl_rule.instance vpc-id#dcg-id#translation-ip-pool#acl-rule-id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_local_source_ip_port_translation_acl_rule", CNDescription{
		TerraformTypeCN: "VPC本地源IP端口转换ACL规则",
		DescriptionCN:   "提供创建VPC本地源IP端口转换ACL规则的资源。",
		AttributesCN: map[string]string{
			"vpc_id":                    "VPC实例ID",
			"direct_connect_gateway_id": "专线网关ID",
			"translation_ip_pool":       "转换后IP池",
			"protocol":                  "协议类型",
			"source_port":               "源端口",
			"source_cidr":               "源CIDR",
			"destination_port":          "目标端口",
			"destination_cidr":          "目标CIDR",
			"acl_rule_id":               "ACL规则ID",
		},
	})
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationAclRule() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to creating VPC local source IP port translation ACL rule.",
		Create:      resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleCreate,
		Read:        resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleRead,
		Update:      resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleUpdate,
		Delete:      resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		CustomizeDiff: customdiff.All(
			validateVpcDcGatewayProtocol),
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPC instance ID.",
			},
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Direct connect gateway ID.",
			},
			"translation_ip_pool": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Translation IP pool.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Protocol type. Valid values: `all`, `tcp`, `udp`.",
			},
			"source_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Source port. Use `0` for all ports.",
			},
		"source_cidr": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "Source CIDR.",
			DiffSuppressFunc: diffSuppressSingleIpCidr,
		},
			"destination_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Destination port. Use `0` for all ports.",
			},
		"destination_cidr": {
			Type:             schema.TypeString,
			Required:         true,
			Description:      "Destination CIDR.",
			DiffSuppressFunc: diffSuppressSingleIpCidr,
		},
			"acl_rule_id": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "ACL rule ID.",
			},
		},
	}
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_source_ip_port_translation_acl_rule.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		vpcId                  = d.Get("vpc_id").(string)
		directConnectGatewayId = d.Get("direct_connect_gateway_id").(string)
		translationIpPool      = d.Get("translation_ip_pool").(string)
		protocol               = d.Get("protocol").(string)
		sourcePort             = d.Get("source_port").(string)
		sourceCidr             = d.Get("source_cidr").(string)
		destinationPort        = d.Get("destination_port").(string)
		destinationCidr        = d.Get("destination_cidr").(string)
	)

	aclRuleId, err := service.CreateLocalSourceIpPortTranslationAclRule(ctx, vpcId, directConnectGatewayId, translationIpPool, protocol, sourcePort, sourceCidr, destinationPort, destinationCidr)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s#%s#%s#%d", vpcId, directConnectGatewayId, translationIpPool, aclRuleId))

	return resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleRead(d, meta)
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_source_ip_port_translation_acl_rule.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	translationIpPool := idSplit[2]
	aclRuleIdStr := idSplit[3]

	aclRuleId, err := strconv.Atoi(aclRuleIdStr)
	if err != nil {
		return fmt.Errorf("invalid acl rule id: %s", aclRuleIdStr)
	}

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		rule, e := service.DescribeLocalSourceIpPortTranslationAclRule(ctx, vpcId, directConnectGatewayId, aclRuleId)
		if e != nil {
			return retryError(e)
		}

		if rule == nil {
			log.Printf("[WARN] VPC Local Source IP Port Translation ACL Rule %s not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		_ = d.Set("vpc_id", vpcId)
		_ = d.Set("direct_connect_gateway_id", directConnectGatewayId)
		_ = d.Set("translation_ip_pool", translationIpPool)
		if rule.Protocol != nil {
			_ = d.Set("protocol", *rule.Protocol)
		}
		if rule.SourcePort != nil {
			_ = d.Set("source_port", *rule.SourcePort)
		}
	if rule.SourceCidr != nil {
		_ = d.Set("source_cidr", normalizeSingleIpCidr(*rule.SourceCidr))
	}
		if rule.DestinationPort != nil {
			_ = d.Set("destination_port", *rule.DestinationPort)
		}
	if rule.DestinationCidr != nil {
		_ = d.Set("destination_cidr", normalizeSingleIpCidr(*rule.DestinationCidr))
	}
		if rule.AclRuleId != nil {
			_ = d.Set("acl_rule_id", *rule.AclRuleId)
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_source_ip_port_translation_acl_rule.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	translationIpPool := idSplit[2]
	aclRuleIdStr := idSplit[3]

	aclRuleId, err := strconv.Atoi(aclRuleIdStr)
	if err != nil {
		return fmt.Errorf("invalid acl rule id: %s", aclRuleIdStr)
	}

	if d.HasChange("protocol") || d.HasChange("source_port") || d.HasChange("source_cidr") ||
		d.HasChange("destination_port") || d.HasChange("destination_cidr") {

		var (
			protocol        = d.Get("protocol").(string)
			sourcePort      = d.Get("source_port").(string)
			sourceCidr      = d.Get("source_cidr").(string)
			destinationPort = d.Get("destination_port").(string)
			destinationCidr = d.Get("destination_cidr").(string)
		)

		err = service.ModifyLocalSourceIpPortTranslationAclRule(ctx, vpcId, directConnectGatewayId, translationIpPool, aclRuleId, protocol, sourcePort, sourceCidr, destinationPort, destinationCidr)
		if err != nil {
			return err
		}
	}

	return resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleRead(d, meta)
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationAclRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_source_ip_port_translation_acl_rule.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 4 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	translationIpPool := idSplit[2]
	aclRuleIdStr := idSplit[3]

	aclRuleId, err := strconv.Atoi(aclRuleIdStr)
	if err != nil {
		return fmt.Errorf("invalid acl rule id: %s", aclRuleIdStr)
	}

	return service.DeleteLocalSourceIpPortTranslationAclRule(ctx, vpcId, directConnectGatewayId, translationIpPool, aclRuleId)
}

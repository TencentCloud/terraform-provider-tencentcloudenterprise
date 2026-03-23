/*
Provides a resource to creating VPC local IP translation ACL rule.

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

resource "tencentcloudenterprise_vpc_local_ip_translation_nat_rule" "nat_rule" {
  vpc_id                    = tencentcloudenterprise_vpc.main.id
  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
  original_ip               = "8.123.40.13"
  translation_ip            = "8.123.45.14"
  description               = "test nat rule"
}

resource "tencentcloudenterprise_vpc_local_ip_translation_acl_rule" "main" {
  vpc_id                    = tencentcloudenterprise_vpc.main.id
  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
  original_ip               = tencentcloudenterprise_vpc_local_ip_translation_nat_rule.nat_rule.original_ip
  translation_ip            = tencentcloudenterprise_vpc_local_ip_translation_nat_rule.nat_rule.translation_ip
  protocol                  = "all"
  source_port               = "0"
  destination_port          = "0"
  destination_cidr          = "10.0.0.0/30"
}
```

Import

VPC local IP translation ACL rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_local_ip_translation_acl_rule.instance vpc-id#dcg-id#original-ip#translation-ip#acl-rule-id
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_local_ip_translation_acl_rule", CNDescription{
		TerraformTypeCN: "本地IP转换ACL规则",
		DescriptionCN:   "提供VPC本地IP转换ACL规则资源，用于创建本地IP转换ACL规则。",
		AttributesCN: map[string]string{
			"vpc_id":                    "VPC实例ID",
			"direct_connect_gateway_id": "专线网关ID",
			"original_ip":               "原始IP地址",
			"translation_ip":            "转换后IP地址",
			"protocol":                  "协议类型",
			"source_port":               "源端口",
			"destination_port":          "目标端口",
			"destination_cidr":          "目标CIDR",
			"acl_rule_id":               "ACL规则ID",
		},
	})
}

func resourceTencentCloudVpcLocalIpTranslationAclRule() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to creating VPC local IP translation ACL rule.",
		Create:      resourceTencentCloudVpcLocalIpTranslationAclRuleCreate,
		Read:        resourceTencentCloudVpcLocalIpTranslationAclRuleRead,
		Update:      resourceTencentCloudVpcLocalIpTranslationAclRuleUpdate,
		Delete:      resourceTencentCloudVpcLocalIpTranslationAclRuleDelete,
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
			"original_ip": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Original IP address.",
			},
			"translation_ip": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Translation IP address.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				//ValidateFunc: validateVpcDcGatewayProtocol(),
				Description: "Protocol type. Valid values: `all`, `tcp`, `udp`.",
			},
			"source_port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Source port. Use `0` for all ports.",
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

func resourceTencentCloudVpcLocalIpTranslationAclRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_ip_translation_acl_rule.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		vpcId                   = d.Get("vpc_id").(string)
		directConnectGatewayId  = d.Get("direct_connect_gateway_id").(string)
		originalIp              = d.Get("original_ip").(string)
		translationIp           = d.Get("translation_ip").(string)
		protocol                = d.Get("protocol").(string)
		sourcePort              = d.Get("source_port").(string)
		destinationPort         = d.Get("destination_port").(string)
		destinationCidr         = d.Get("destination_cidr").(string)
	)

	aclRuleId, err := service.CreateLocalIpTranslationAclRule(ctx, vpcId, directConnectGatewayId, 
		originalIp, translationIp, protocol, sourcePort, destinationPort, destinationCidr)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s#%s#%s#%s#%d", vpcId, directConnectGatewayId, originalIp, translationIp, aclRuleId))

	return resourceTencentCloudVpcLocalIpTranslationAclRuleRead(d, meta)
}

func resourceTencentCloudVpcLocalIpTranslationAclRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_ip_translation_acl_rule.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 5 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	originalIp := idSplit[2]
	translationIp := idSplit[3]
	aclRuleIdStr := idSplit[4]
	
	aclRuleId, err := strconv.Atoi(aclRuleIdStr)
	if err != nil {
		return fmt.Errorf("invalid acl rule id: %s", aclRuleIdStr)
	}

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		rule, e := service.DescribeLocalIpTranslationAclRule(ctx, vpcId, directConnectGatewayId, 
			originalIp, translationIp, aclRuleId)
		if e != nil {
			return retryError(e)
		}

		if rule == nil {
			log.Printf("[WARN] VPC Local IP Translation ACL Rule %s not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		_ = d.Set("vpc_id", vpcId)
		_ = d.Set("direct_connect_gateway_id", directConnectGatewayId)
		_ = d.Set("original_ip", originalIp)
		_ = d.Set("translation_ip", translationIp)
		if rule.Protocol != nil {
			_ = d.Set("protocol", *rule.Protocol)
		}
		if rule.SourcePort != nil {
			_ = d.Set("source_port", *rule.SourcePort)
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

func resourceTencentCloudVpcLocalIpTranslationAclRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_ip_translation_acl_rule.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 5 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	originalIp := idSplit[2]
	translationIp := idSplit[3]
	aclRuleIdStr := idSplit[4]
	
	aclRuleId, err := strconv.Atoi(aclRuleIdStr)
	if err != nil {
		return fmt.Errorf("invalid acl rule id: %s", aclRuleIdStr)
	}

	if d.HasChange("protocol") || d.HasChange("source_port") || 
		d.HasChange("destination_port") || d.HasChange("destination_cidr") {
		
		var (
			protocol        = d.Get("protocol").(string)
			sourcePort      = d.Get("source_port").(string)
			destinationPort = d.Get("destination_port").(string)
			destinationCidr = d.Get("destination_cidr").(string)
		)

		err = service.ModifyLocalIpTranslationAclRule(ctx, vpcId, directConnectGatewayId,
			originalIp, translationIp, aclRuleId, protocol, sourcePort, destinationPort, destinationCidr)
		if err != nil {
			return err
		}
	}


	return resourceTencentCloudVpcLocalIpTranslationAclRuleRead(d, meta)
}

func resourceTencentCloudVpcLocalIpTranslationAclRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_ip_translation_acl_rule.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 5 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	originalIp := idSplit[2]
	translationIp := idSplit[3]
	aclRuleIdStr := idSplit[4]
	
	aclRuleId, err := strconv.Atoi(aclRuleIdStr)
	if err != nil {
		return fmt.Errorf("invalid acl rule id: %s", aclRuleIdStr)
	}

	return service.DeleteLocalIpTranslationAclRule(ctx, vpcId, directConnectGatewayId, 
		originalIp, translationIp, aclRuleId)
}
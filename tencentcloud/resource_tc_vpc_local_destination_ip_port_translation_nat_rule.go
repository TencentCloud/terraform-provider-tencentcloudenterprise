/*
Provides a resource to creating VPC local destination IP port translation NAT rule.

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

resource "tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule" "nat_rule" {
  vpc_id                    = tencentcloudenterprise_vpc.main.id
  direct_connect_gateway_id = tencentcloudenterprise_vpc_dc_gateway.dcg_main.id
  protocol                  = "tcp"
  original_ip               = "10.0.1.1"
  original_port             = 80
  translation_ip            = "10.0.2.1"
  translation_port          = 8080
  description               = "test destination nat rule"
}
```

Import

VPC local destination IP port translation NAT rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule.instance vpc-id#dcg-id#protocol#original-ip#original-port
```
*/
package tencentcloud

import (
	"context"
	"fmt"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule", CNDescription{
		TerraformTypeCN: "VPC本地目的IP端口转换NAT规则",
		DescriptionCN:   "提供创建VPC本地目的IP端口转换NAT规则的资源。",
		AttributesCN: map[string]string{
			"vpc_id":                    "VPC实例ID",
			"direct_connect_gateway_id": "专线网关ID",
			"protocol":                  "协议类型",
			"original_ip":               "源IP",
			"original_port":             "源端口",
			"translation_ip":            "转换后IP",
			"translation_port":          "转换后端口",
			"description":               "描述",
		},
	})
}

func resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRule() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to creating VPC local destination IP port translation NAT rule.",
		Create:      resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleCreate,
		Read:        resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleRead,
		Update:      resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleUpdate,
		Delete:      resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Protocol type. Valid values: `tcp`, `udp`.",
			},
			"original_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Original IP address.",
			},
			"original_port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Original port.",
			},
			"translation_ip": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Translation IP address.",
			},
			"translation_port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Translation port.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description.",
			},
		},
	}
}

func resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		vpcId                   = d.Get("vpc_id").(string)
		directConnectGatewayId  = d.Get("direct_connect_gateway_id").(string)
		protocol                = d.Get("protocol").(string)
		originalIp              = d.Get("original_ip").(string)
		originalPort            = int64(d.Get("original_port").(int))
		translationIp           = d.Get("translation_ip").(string)
		translationPort         = int64(d.Get("translation_port").(int))
		description             = d.Get("description").(string)
	)

	err := service.CreateLocalDestinationIpPortTranslationNatRule(ctx, vpcId, directConnectGatewayId, protocol, originalIp, translationIp, description, originalPort, translationPort)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s#%s#%s#%s#%d", vpcId, directConnectGatewayId, protocol, originalIp, originalPort))

	return resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleRead(d, meta)
}

func resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 5 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	protocol := idSplit[2]
	originalIp := idSplit[3]
	originalPortStr := idSplit[4]

	originalPort, err := strconv.ParseInt(originalPortStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid original port: %s", originalPortStr)
	}

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		rule, e := service.DescribeLocalDestinationIpPortTranslationNatRule(ctx, vpcId, directConnectGatewayId, protocol, originalIp, originalPort)
		if e != nil {
			return retryError(e)
		}

		if rule == nil {
			log.Printf("[WARN] VPC Local Destination IP Port Translation NAT Rule %s not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		_ = d.Set("vpc_id", vpcId)
		_ = d.Set("direct_connect_gateway_id", directConnectGatewayId)
		if rule.Protocol != nil {
			_ = d.Set("protocol", *rule.Protocol)
		}
		if rule.OriginalIp != nil {
			_ = d.Set("original_ip", *rule.OriginalIp)
		}
		if rule.OriginalPort != nil {
			_ = d.Set("original_port", int(*rule.OriginalPort))
		}
		if rule.TranslationIp != nil {
			_ = d.Set("translation_ip", *rule.TranslationIp)
		}
		if rule.TranslationPort != nil {
			_ = d.Set("translation_port", int(*rule.TranslationPort))
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

func resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 5 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	oldProtocol := idSplit[2]
	oldOriginalIp := idSplit[3]
	oldOriginalPortStr := idSplit[4]
	var oldTranslationPort int64
	var oldTranslationIp string

	if d.HasChange("translation_port") {
		oldTranslationPortValue, _ := d.GetChange("translation_port")
		if v, ok := oldTranslationPortValue.(int); ok {
			oldTranslationPort = int64(v)
		} else {
			return fmt.Errorf("translation_port type assertion failed")
		}
	} else {
		oldTranslationPortValue := d.Get("translation_port")
		if v, ok := oldTranslationPortValue.(int); ok {
			oldTranslationPort = int64(v)
		} else {
			return fmt.Errorf("translation_port type assertion failed")
		}
	}

	if d.HasChange("translation_ip") {
		oldTranslationIpValue, _ := d.GetChange("translation_ip")
		if v, ok := oldTranslationIpValue.(string); ok {
			oldTranslationIp = v
		} else {
			return fmt.Errorf("translation_ip type assertion failed")
		}
	} else {
		oldTranslationIpValue := d.Get("translation_ip")
		if v, ok := oldTranslationIpValue.(string); ok {
			oldTranslationIp = v
		} else {
			return fmt.Errorf("translation_ip type assertion failed")
		}
	}



	oldOriginalPort, err := strconv.ParseInt(oldOriginalPortStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid old original port: %s", oldOriginalPortStr)
	}

	var (
		newProtocol         = d.Get("protocol").(string)
		newOriginalIp       = d.Get("original_ip").(string)
		newOriginalPort     = int64(d.Get("original_port").(int))
		newTranslationIp    = d.Get("translation_ip").(string)
		newTranslationPort  = int64(d.Get("translation_port").(int))
		description         = d.Get("description").(string)
	)

	err = service.ModifyLocalDestinationIpPortTranslationNatRule(ctx, vpcId, directConnectGatewayId, oldProtocol, oldOriginalIp,oldTranslationIp,
		newProtocol, newOriginalIp, newTranslationIp, description, oldOriginalPort, newOriginalPort, newTranslationPort, oldTranslationPort)
	if err != nil {
		return err
	}

	// 如果关键字段发生变化，需要更新资源ID
	if newProtocol != oldProtocol || newOriginalIp != oldOriginalIp || newOriginalPort != oldOriginalPort {
		d.SetId(fmt.Sprintf("%s#%s#%s#%s#%d", vpcId, directConnectGatewayId, newProtocol, newOriginalIp, newOriginalPort))
	}

	return resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleRead(d, meta)
}

func resourceTencentCloudVpcLocalDestinationIpPortTranslationNatRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_destination_ip_port_translation_nat_rule.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 5 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	protocol := idSplit[2]
	originalIp := idSplit[3]
	originalPortStr := idSplit[4]

	originalPort, err := strconv.ParseInt(originalPortStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid original port: %s", originalPortStr)
	}

	translationIp := d.Get("translation_ip").(string)
	translationPort := int64(d.Get("translation_port").(int))
	var description *string
	if v, ok := d.GetOk("description"); ok{
		description = helper.String(v.(string))

	}


	return service.DeleteLocalDestinationIpPortTranslationNatRule(ctx, vpcId, directConnectGatewayId, protocol,
		originalIp, originalPort, translationIp, translationPort, description)
}
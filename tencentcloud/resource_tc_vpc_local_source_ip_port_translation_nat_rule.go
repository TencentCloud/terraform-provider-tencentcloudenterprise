
/*
Provides a resource to creating VPC local source IP port translation NAT rule.

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
```

Import

VPC local source IP port translation NAT rule can be imported, e.g.

```
$ terraform import tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule.instance vpc-id#dcg-id#ip-pool
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
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule", CNDescription{
		TerraformTypeCN: "VPC本地源IP端口转换NAT规则",
		DescriptionCN:   "提供创建VPC本地源IP端口转换NAT规则的资源。",
		AttributesCN: map[string]string{
			"vpc_id":                    "VPC实例ID",
			"direct_connect_gateway_id": "专线网关ID",
			"ip_pool":                   "IP池",
			"description":               "描述",
		},
	})
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationNatRule() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to creating VPC local source IP port translation NAT rule.",
		Create:      resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleCreate,
		Read:        resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleRead,
		Update:      resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleUpdate,
		Delete:      resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleDelete,
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
			"ip_pool": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP pool.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description.",
			},
		},
	}
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		vpcId                   = d.Get("vpc_id").(string)
		directConnectGatewayId  = d.Get("direct_connect_gateway_id").(string)
		ipPool                  = d.Get("ip_pool").(string)
		description             = d.Get("description").(string)
	)

	err := service.CreateLocalSourceIpPortTranslationNatRule(ctx, vpcId, directConnectGatewayId, ipPool, description)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s#%s#%s", vpcId, directConnectGatewayId, ipPool))

	return resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleRead(d, meta)
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	ipPool := idSplit[2]

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		rule, e := service.DescribeLocalSourceIpPortTranslationNatRule(ctx, vpcId, directConnectGatewayId, ipPool)
		if e != nil {
			return retryError(e)
		}

		if rule == nil {
			log.Printf("[WARN] VPC Local Source IP Port Translation NAT Rule %s not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}

		_ = d.Set("vpc_id", vpcId)
		_ = d.Set("direct_connect_gateway_id", directConnectGatewayId)
		if rule.IpPool != nil {
			_ = d.Set("ip_pool", *rule.IpPool)
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

func resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	oldIpPool := idSplit[2]

	var (
		newIpPool    = d.Get("ip_pool").(string)
		description   = d.Get("description").(string)
	)

	err := service.ModifyLocalSourceIpPortTranslationNatRule(ctx, vpcId, directConnectGatewayId, oldIpPool, newIpPool, description)
	if err != nil {
		return err
	}

	if oldIpPool != newIpPool {
		d.SetId(fmt.Sprintf("%s#%s#%s", vpcId, directConnectGatewayId, newIpPool))
	}

	return resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleRead(d, meta)
}

func resourceTencentCloudVpcLocalSourceIpPortTranslationNatRuleDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_local_source_ip_port_translation_nat_rule.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	idSplit := strings.Split(d.Id(), "#")
	if len(idSplit) != 3 {
		return fmt.Errorf("id is broken, id is %s", d.Id())
	}
	vpcId := idSplit[0]
	directConnectGatewayId := idSplit[1]
	ipPool := idSplit[2]

	return service.DeleteLocalSourceIpPortTranslationNatRule(ctx, vpcId, directConnectGatewayId, ipPool)
}

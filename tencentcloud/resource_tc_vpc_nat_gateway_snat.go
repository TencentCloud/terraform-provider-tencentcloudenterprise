/*
Provides a resource to create a NAT Gateway SNat rule.

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc_nat_gateway_snat" "subnet_snat" {
	  nat_gateway_id    = tencentcloudenterprise_vpc_nat_gateway.my_nat.id
	  resource_type     = "SUBNET"
	  subnet_id         = tencentcloudenterprise_vpc_subnet.my_subnet.id
	  subnet_cidr_block = tencentcloudenterprise_vpc_subnet.my_subnet.cidr_block
	  description       = "terraform test"
	  public_ip_addr    = [
	    tencentcloudenterprise_eip.eip1.public_ip,
	    tencentcloudenterprise_eip.eip2.public_ip,
	  ]
	}

	resource "tencentcloudenterprise_vpc_nat_gateway_snat" "instance_snat" {
	  nat_gateway_id           = tencentcloudenterprise_vpc_nat_gateway.my_nat.id
	  resource_type            = "NETWORKINTERFACE"
	  instance_id              = tencentcloudenterprise_cvm_instance.my_instance.id
	  instance_private_ip_addr = tencentcloudenterprise_cvm_instance.my_instance.private_ip
	  description              = "terraform test"
	  public_ip_addr           = [
	    tencentcloudenterprise_eip.eip1.public_ip,
	  ]
	}

```

# Import

NAT gateway snat rule can be imported using the id, the id format must be '{nat_gateway_id}#{resource_id}',
resource_id range `subnet_id`, `instance_id`, e.g.

SUBNET SNat
```
$ terraform import tencentcloudenterprise_vpc_nat_gateway_snat.subnet_snat nat-r4ip1cwt#subnet-2ap74y35
```

NETWORKINTERFACE SNat
```
$ terraform import tencentcloudenterprise_vpc_nat_gateway_snat.instance_snat nat-r4ip1cwt#ins-da412f5a
```
*/
package tencentcloud

import (
	"context"
	"errors"
	"log"

	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_nat_gateway_snat", CNDescription{
		TerraformTypeCN: "NAT网关SNAT规则",
		DescriptionCN:   "提供NAT网关SNAT规则资源，用于创建NAT网关的源IP转换规则。",
		AttributesCN: map[string]string{
			"nat_gateway_id":           "NAT网关ID",
			"resource_type":            "资源类型",
			"subnet_id":                "子网实例ID",
			"subnet_cidr_block":        "子网IPv4 CIDR",
			"instance_id":              "实例ID",
			"instance_private_ip_addr": "实例主ENI的内网IP",
			"public_ip_addr":           "弹性IP地址池",
			"description":              "描述",
			"snat_id":                  "SNAT规则ID",
			"create_time":              "创建时间",
		},
	})
}

func resourceTencentCloudVpcNatGatewaySnat() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a NAT Gateway SNat rule.",
		Create:      resourceTencentCloudVpcNatGatewaySnatCreate,
		Read:        resourceTencentCloudVpcNatGatewaySnatRead,
		Update:      resourceTencentCloudVpcNatGatewaySnatUpdate,
		Delete:      resourceTencentCloudVpcNatGatewaySnatDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"nat_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "NAT gateway ID.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Resource type. Valid values: " + NAT_GATEWAY_TYPE_SUBNET + ", " + NAT_GATEWAY_TYPE_NETWORK_INTERFACE + ".",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Subnet instance ID, required when `resource_type` is " + NAT_GATEWAY_TYPE_SUBNET + ".",
			},
			"subnet_cidr_block": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The IPv4 CIDR of the subnet, required when `resource_type` is " + NAT_GATEWAY_TYPE_SUBNET + ".",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Instance ID, required when `resource_type` is " + NAT_GATEWAY_TYPE_NETWORK_INTERFACE + ".",
			},
			"instance_private_ip_addr": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Private IPs of the instance's primary ENI, required when `resource_type` is " + NAT_GATEWAY_TYPE_NETWORK_INTERFACE + ".",
			},
			"public_ip_addr": {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Elastic IP address pool.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description.",
			},
			"snat_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SNAT rule ID.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Create time.",
			},
		},
	}
}

func resourceTencentCloudVpcNatGatewaySnatCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_nat_gateway_snat.create")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	vpcService := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	natGatewayId := d.Get("nat_gateway_id").(string)

	if err := vpcNatGatewaySnatParamValid(d); err != nil {
		return err
	}

	snat := buildVpcNatGatewaySnat(d)
	if err := vpcService.CreateNatGatewaySnat(ctx, natGatewayId, snat); err != nil {
		log.Printf("[CRITAL]%s create nat gateway snat failed, reason:%s\n", logId, err.Error())
		return err
	}

	err, result := vpcService.DescribeNatGatewaySnats(ctx, natGatewayId, []*vpc.Filter{
		{Name: helper.String("resource-id"), Values: []*string{snat.ResourceId}},
		{Name: helper.String("public-ip-address"), Values: snat.PublicIpAddresses},
		{Name: helper.String("description"), Values: []*string{snat.Description}},
	})
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("[CRITAL] create nat gateway snat failed: read result is empty")
	}
	rule := result[len(result)-1]
	d.SetId(helper.IdFormat(*rule.NatGatewayId, *rule.ResourceId))

	return resourceTencentCloudVpcNatGatewaySnatRead(d, meta)
}

func resourceTencentCloudVpcNatGatewaySnatRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_nat_gateway_snat.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	compositeId := helper.IdParse(d.Id())
	if len(compositeId) != 2 {
		return errors.New("the id format must be '{nat_gateway_id}#{resource_id}'")
	}

	err, snatList := service.DescribeNatGatewaySnats(ctx, compositeId[0], nil)
	if err != nil {
		log.Printf("[CRITAL]%s read nat gateway snat failed, reason:%s\n", logId, err.Error())
		return err
	}

	var snat *vpc.SourceIpTranslationNatRule
	for _, s := range snatList {
		if compositeId[1] == *s.ResourceId {
			snat = s
			break
		}
	}
	if snat == nil {
		d.SetId("")
		return nil
	}

	_ = d.Set("nat_gateway_id", snat.NatGatewayId)
	_ = d.Set("resource_type", snat.ResourceType)
	_ = d.Set("public_ip_addr", vpcSnatSortPublicIpAddr(d, snat.PublicIpAddresses))
	_ = d.Set("description", snat.Description)
	_ = d.Set("snat_id", snat.NatGatewaySnatId)
	_ = d.Set("create_time", snat.CreatedTime)

	if *snat.ResourceType == NAT_GATEWAY_TYPE_SUBNET {
		_ = d.Set("subnet_id", snat.ResourceId)
		_ = d.Set("subnet_cidr_block", snat.PrivateIpAddress)
	} else if *snat.ResourceType == NAT_GATEWAY_TYPE_NETWORK_INTERFACE {
		_ = d.Set("instance_id", snat.ResourceId)
		_ = d.Set("instance_private_ip_addr", snat.PrivateIpAddress)
	}

	return nil
}

func resourceTencentCloudVpcNatGatewaySnatUpdate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_nat_gateway_snat.update")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	compositeId := helper.IdParse(d.Id())
	if len(compositeId) != 2 {
		return errors.New("the id format must be '{nat_gateway_id}#{snat_id}'")
	}

	if err := vpcNatGatewaySnatParamValid(d); err != nil {
		return err
	}

	snat := buildVpcNatGatewaySnat(d)
	if err := service.ModifyNatGatewaySnat(ctx, compositeId[0], snat); err != nil {
		log.Printf("[CRITAL]%s modify nat gateway snat failed, reason:%s\n", logId, err.Error())
		return err
	}

	return resourceTencentCloudVpcNatGatewaySnatRead(d, meta)
}

func resourceTencentCloudVpcNatGatewaySnatDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.tencentcloudenterprise_vpc_nat_gateway_snat.delete")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	natGatewayId := d.Get("nat_gateway_id").(string)
	snatId := d.Get("snat_id").(string)
	if err := service.DeleteNatGatewaySnat(ctx, natGatewayId, snatId); err != nil {
		log.Printf("[CRITAL]%s delete nat gateway snat failed, reason:%s\n", logId, err.Error())
		return err
	}

	return nil
}

func vpcNatGatewaySnatParamValid(d *schema.ResourceData) error {
	logId := getLogId(contextNil)
	resourceType := d.Get("resource_type")
	_, hasSubnetId := d.GetOk("subnet_id")
	_, hasCidrBlock := d.GetOk("subnet_cidr_block")
	_, hasInstanceId := d.GetOk("instance_id")
	_, hasPrivateIpAddr := d.GetOk("instance_private_ip_addr")

	if resourceType == NAT_GATEWAY_TYPE_SUBNET && !(hasSubnetId && hasCidrBlock) {
		log.Printf("[CRITAL]%s `resource_type` is %s, but hasSubnetId=%v, hasCidrBlock=%v",
			logId, NAT_GATEWAY_TYPE_SUBNET, hasSubnetId, hasCidrBlock)
		return errors.New("`subnet_id` and `subnet_cidr_block` required when `resource_type` is " + NAT_GATEWAY_TYPE_SUBNET)
	} else if resourceType == NAT_GATEWAY_TYPE_NETWORK_INTERFACE && !(hasInstanceId && hasPrivateIpAddr) {
		log.Printf("[CRITAL]%s `resource_type` is %s, but hasInstanceId=%v, hasPrivateIpAddr=%v",
			logId, NAT_GATEWAY_TYPE_NETWORK_INTERFACE, hasInstanceId, hasPrivateIpAddr)
		return errors.New("`instance_id` and `instance_private_ip_addr` required when `resource_type` is " + NAT_GATEWAY_TYPE_NETWORK_INTERFACE)
	}
	return nil
}

func buildVpcNatGatewaySnat(d *schema.ResourceData) *vpc.SourceIpTranslationNatRule {
	resourceType := d.Get("resource_type").(string)
	publicIpAddrs := helper.InterfacesStringsPoint(d.Get("public_ip_addr").([]interface{}))
	description := helper.String(d.Get("description").(string))

	var resourceId, privateIpAddr string
	if resourceType == NAT_GATEWAY_TYPE_SUBNET {
		resourceId = d.Get("subnet_id").(string)
		privateIpAddr = d.Get("subnet_cidr_block").(string)
	} else if resourceType == NAT_GATEWAY_TYPE_NETWORK_INTERFACE {
		resourceId = d.Get("instance_id").(string)
		privateIpAddr = d.Get("instance_private_ip_addr").(string)
	}

	snat := &vpc.SourceIpTranslationNatRule{
		ResourceId:        &resourceId,
		ResourceType:      &resourceType,
		PrivateIpAddress:  &privateIpAddr,
		PublicIpAddresses: publicIpAddrs,
		Description:       description,
	}
	if v, ok := d.GetOk("snat_id"); ok {
		snat.NatGatewaySnatId = helper.String(v.(string))
	}
	return snat
}

func vpcSnatSortPublicIpAddr(d *schema.ResourceData, publicIpAddresses []*string) []*string {
	v, ok := d.GetOk("public_ip_addr")
	if !ok {
		return publicIpAddresses
	}
	result := make([]*string, 0, len(publicIpAddresses))
	for _, paramObj := range v.([]interface{}) {
		for _, obj := range publicIpAddresses {
			if paramObj.(string) == *obj {
				result = append(result, obj)
				break
			}
		}
	}
	for _, obj := range publicIpAddresses {
		found := false
		for _, exist := range result {
			if exist == obj {
				found = true
				break
			}
		}
		if !found {
			result = append(result, obj)
		}
	}
	return result
}

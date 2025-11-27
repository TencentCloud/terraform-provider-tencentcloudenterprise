/*
Provides a resource to create a vpc ipv6_cidr_block

# Example Usage

```hcl

	resource "tencentcloudenterprise_vpc" "cidr-block" {
	  name         = "ipv6-cidr-block-for-test"
	  cidr_block   = "10.0.0.0/16"
	  is_multicast = false
	}

	resource "tencentcloudenterprise_vpc_ipv6_cidr_block" "ipv6_cidr_block" {
	  vpc_id = cloud_vpc.cidr-block.id
	}

```

# Import

vpc ipv6_cidr_block can be imported using the id, e.g.

```
terraform import cloud_vpc_ipv6_cidr_block.ipv6_cidr_block vpc_id
```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerResourceDescriptionProvider("tencentcloudenterprise_vpc_ipv6_cidr_block", CNDescription{
		TerraformTypeCN: "创建VPC IPv6网段",
		DescriptionCN:   "提供VPC IPv6网段资源，用于创建VPC IPv6网段。",
		AttributesCN: map[string]string{
			"vpc_id":          "VPC实例ID",
			"ipv6_cidr_block": "IPv6网段",
		},
	})
}
func resourceTencentCloudVpcIpv6CidrBlock() *schema.Resource {
	return &schema.Resource{
		Description: "Provides a resource to create a vpc ipv6_cidr_block",
		Create:      resourceTencentCloudVpcIpv6CidrBlockCreate,
		Read:        resourceTencentCloudVpcIpv6CidrBlockRead,
		Delete:      resourceTencentCloudVpcIpv6CidrBlockDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "`VPC` instance `ID`, in the form of `vpc-f49l6u0z`.",
			},
			"ipv6_cidr_block": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Ipv6 cidr block.",
			},
		},
	}
}

func resourceTencentCloudVpcIpv6CidrBlockCreate(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_ipv6_cidr_block.create")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	var (
		request = vpc.NewAssignIpv6CidrBlockRequest()
		vpcId   string
	)
	if v, ok := d.GetOk("vpc_id"); ok {
		vpcId = v.(string)
		request.VpcId = helper.String(vpcId)
	}

	err := resource.Retry(writeRetryTimeout, func() *resource.RetryError {
		result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().AssignIpv6CidrBlock(request)
		if e != nil {
			return retryError(e)
		} else {
			log.Printf("[DEBUG]%s api[%s] success, request body [%s], response body [%s]\n", logId, request.GetAction(), request.ToJsonString(), result.ToJsonString())
		}
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s create vpc ipv6CidrBlock failed, reason:%+v", logId, err)
		return err
	}

	d.SetId(vpcId)

	return resourceTencentCloudVpcIpv6CidrBlockRead(d, meta)
}

func resourceTencentCloudVpcIpv6CidrBlockRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_ipv6_cidr_block.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	vpcId := d.Id()

	instance, err := service.DescribeVpcById(ctx, vpcId)
	if err != nil {
		return err
	}

	if instance == nil {
		d.SetId("")
		log.Printf("[WARN]%s resource `VpcIpv6CidrBlock` [%s] not found, please check if it has been deleted.\n", logId, d.Id())
		return nil
	}

	if instance.VpcId != nil {
		_ = d.Set("vpc_id", instance.VpcId)
	}

	if instance.Ipv6CidrBlock != nil {
		_ = d.Set("ipv6_cidr_block", instance.Ipv6CidrBlock)
	}

	return nil
}

func resourceTencentCloudVpcIpv6CidrBlockDelete(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("resource.cloud_vpc_ipv6_cidr_block.delete")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}
	vpcId := d.Id()

	if err := service.DeleteVpcIpv6CidrBlockById(ctx, vpcId); err != nil {
		return err
	}

	return nil
}

/*
Provides details about a specific VPC.

This resource can prove useful when a module accepts a vpc id as an input variable and needs to, for example, determine the CIDR block of that VPC.

~> **NOTE:** It has been deprecated and replaced by cloud_vpc_instances.

# Example Usage

```hcl
variable "vpc_id" {}

	data "tencentcloudenterprise_vpc" "selected" {
	  id = var.vpc_id
	}

	resource "tencentcloudenterprise_vpc_subnet" "main" {
	  name              = "my test subnet"
	  cidr_block        = cidrsubnet(data.cloud_vpc.selected.cidr_block, 4, 1)
	  availability_zone = "eu-frankfurt-1"
	  vpc_id            = data.cloud_vpc.selected.id
	}

```
*/
package tencentcloud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc", CNDescription{
		TerraformTypeCN: "VPC",
		AttributesCN: map[string]string{
			"id":           "要检索的特定VPC的ID",
			"name":         "要检索的特定VPC的名称",
			"cidr_block":   "VPC的CIDR块",
			"is_default":   "是否为默认VPC",
			"is_multicast": "VPC是否支持组播",
		},
	})
}

func dataSourceTencentCloudVpc() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This data source has been deprecated in Terraform TencentCloud provider version 1.10.0. Please use 'cloud_vpc_instances' instead.",
		Read:               dataSourceTencentCloudVpcRead,

		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ID of the specific VPC to retrieve.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the specific VPC to retrieve.",
			},
			"cidr_block": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CIDR block of the VPC.",
			},
			"is_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether or not the default VPC.",
			},
			"is_multicast": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether or not the VPC has Multicast support.",
			},
		},
	}
}

func dataSourceTencentCloudVpcRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)
	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var (
		vpcId string
		name  string
	)
	if temp, ok := d.GetOk("id"); ok {
		tempStr := temp.(string)
		if tempStr != "" {
			vpcId = tempStr
		}
	}
	if temp, ok := d.GetOk("name"); ok {
		tempStr := temp.(string)
		if tempStr != "" {
			name = tempStr
		}
	}

	var vpcInfos, err = service.DescribeVpcs(ctx, vpcId, name, map[string]string{}, nil, "", "")
	if err != nil {
		return err
	}

	if len(vpcInfos) == 0 {
		return fmt.Errorf("vpc id=%s, name=%s not found", vpcId, name)
	}

	vpc := vpcInfos[0]
	d.SetId(vpc.vpcId)
	_ = d.Set("name", vpc.name)
	_ = d.Set("cidr_block", vpc.cidr)
	_ = d.Set("is_default", vpc.isDefault)
	_ = d.Set("is_multicast", vpc.isMulticast)

	return nil
}

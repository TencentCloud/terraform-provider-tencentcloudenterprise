/*
Use this data source to query detailed information of vpc private_ip_addresses

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_private_ip_addresses" "private_ip_addresses" {
	  vpc_id = "vpc-l0dw94uh"
	  private_ip_addresses = ["203.0.113.1"]
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_private_ip_addresses", CNDescription{
		TerraformTypeCN: "私有网络专用IP查询",
		DescriptionCN:   "提供VPC私有IP地址数据源，用于查询VPC私有IP地址的详细信息。",
		AttributesCN: map[string]string{
			"vpc_id":                     "“VPC”的“ID”，例如“VPC-f49l6u0z”",
			"private_ip_addresses":       "专用“IP”地址列表每个请求最多支持“10”个批处理查询",
			"vpc_private_ip_address_set": "专用“IP”地址信息的列表",
			"private_ip_address":         "`专有网络`private`IP`",
			"cidr_block":                 "属于子网的“CIDR”",
			"private_ip_address_type":    "专用“IP”类型",
			"created_time":               "`IP`应用程序时间",
			"result_output_file":         "用于保存结果",
		},
	})
}

func dataSourceTencentCloudVpcPrivateIpAddresses() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed information of vpc private_ip_addresses",
		Read:        dataSourceTencentCloudVpcPrivateIpAddressesRead,
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The `ID` of the `VPC`, such as `vpc-f49l6u0z`.",
			},

			"private_ip_addresses": {
				Required: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The private `IP` address list. Each request supports a maximum of `10` batch querying.",
			},

			"vpc_private_ip_address_set": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "The list of private `IP` address information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"private_ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "`VPC` private `IP`.",
						},
						"cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The `CIDR` belonging to the subnet.",
						},
						"private_ip_address_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Private `IP` type.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "`IP` application time.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudVpcPrivateIpAddressesRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_private_ip_addresses.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)

	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("vpc_id"); ok {
		paramMap["VpcId"] = helper.String(v.(string))
	}

	if v, ok := d.GetOk("private_ip_addresses"); ok {
		privateIpAddressesSet := v.(*schema.Set).List()
		paramMap["PrivateIpAddresses"] = helper.InterfacesStringsPoint(privateIpAddressesSet)
	}

	service := VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

	var vpcPrivateIpAddressSet []*vpc.VpcPrivateIpAddress

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeVpcPrivateIpAddresses(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		vpcPrivateIpAddressSet = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(vpcPrivateIpAddressSet))
	tmpList := make([]map[string]interface{}, 0, len(vpcPrivateIpAddressSet))

	if vpcPrivateIpAddressSet != nil {
		for _, vpcPrivateIpAddress := range vpcPrivateIpAddressSet {
			vpcPrivateIpAddressMap := map[string]interface{}{}

			if vpcPrivateIpAddress.PrivateIpAddress != nil {
				vpcPrivateIpAddressMap["private_ip_address"] = vpcPrivateIpAddress.PrivateIpAddress
			}

			if vpcPrivateIpAddress.CidrBlock != nil {
				vpcPrivateIpAddressMap["cidr_block"] = vpcPrivateIpAddress.CidrBlock
			}

			if vpcPrivateIpAddress.PrivateIpAddressType != nil {
				vpcPrivateIpAddressMap["private_ip_address_type"] = vpcPrivateIpAddress.PrivateIpAddressType
			}

			if vpcPrivateIpAddress.CreatedTime != nil {
				vpcPrivateIpAddressMap["created_time"] = vpcPrivateIpAddress.CreatedTime
			}

			ids = append(ids, *vpcPrivateIpAddress.PrivateIpAddress)
			tmpList = append(tmpList, vpcPrivateIpAddressMap)
		}

		_ = d.Set("vpc_private_ip_address_set", tmpList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), tmpList); e != nil {
			return e
		}
	}
	return nil
}

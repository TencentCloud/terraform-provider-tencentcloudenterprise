/*
Use this data source to query detailed information of HA VIPs.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_ha_vips" "havips" {
	  id         = "havip-kjqwe4ba"
	  name       = "test"
	  vpc_id     = "vpc-gzea3dd7"
	  subnet_id  = "subnet-4d4m4cd4"
	  address_ip = "10.0.4.16"
	}

```
*/
package tencentcloud

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pkg/errors"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_ha_vips", CNDescription{
		TerraformTypeCN: "高可用虚拟IP（HAVIP）列表",
		DescriptionCN:   "提供HA VIP数据源，用于查询HA VIP的详细信息。",
		AttributesCN: map[string]string{
			"name":                 "HA VIP的名称字符长度限制为1-60",
			"id":                   "要查询的HA VIP的ID",
			"vpc_id":               "要查询的HA VIP的VPC id",
			"subnet_id":            "要查询的HA VIP的子网id",
			"address_ip":           "要查询的HA VIP的EIP",
			"result_output_file":   "用于保存结果",
			"ha_vip_list":          "医管局专用贵宾名单",
			"vip":                  "虚拟IP地址，不能被占用，并且在这个VPC网段中如果未设置，它将在自动创建资源后分配",
			"state":                "HA VIP的状态有效值：“AVAILABLE”、“UNBIND”",
			"instance_id":          "关联的实例id",
			"network_interface_id": "关联的网络接口id",
			"create_time":          "创建HA VIP的时间",
		},
	})
}

func dataSourceTencentCloudHaVips() *schema.Resource {
	return &schema.Resource{
		Description: "This data source provides the details of a HA VIP.",
		Read:        dataSourceTencentCloudHaVipsRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringLengthInRange(1, 60),
				Description:  "Name of the HA VIP. The length of character is limited to 1-60.",
			},
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the HA VIP to be queried.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VPC id of the HA VIP to be queried.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Subnet id of the HA VIP to be queried.",
			},
			"address_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateIp,
				Description:  "EIP of the HA VIP to be queried.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// Computed values
			"ha_vip_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information list of the dedicated HA VIPs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the HA VIP.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the HA VIP.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPC id.",
						},
						"subnet_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subnet id.",
						},
						"vip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Virtual IP address, it must not be occupied and in this VPC network segment. If not set, it will be assigned after resource created automatically.",
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "State of the HA VIP. Valid values: `AVAILABLE`, `UNBIND`.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance id that is associated.",
						},
						"network_interface_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network interface id that is associated.",
						},
						"address_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "EIP that is associated.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time of the HA VIP.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudHaVipsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_ha_vips.read")()

	logId := getLogId(contextNil)

	request := vpc.NewDescribeHaVipsRequest()

	params := make(map[string]string)
	if v, ok := d.GetOk("id"); ok {
		params["havip-id"] = v.(string)
	}
	if v, ok := d.GetOk("name"); ok {
		params["havip-name"] = v.(string)
	}
	if v, ok := d.GetOk("address_ip"); ok {
		params["address-ip"] = v.(string)
	}
	if v, ok := d.GetOk("subnet_id"); ok {
		params["subnet-id"] = v.(string)
	}
	if v, ok := d.GetOk("vpc_id"); ok {
		params["vpc-ip"] = v.(string)
	}
	request.Filters = make([]*vpc.Filter, 0, len(params))
	for k, v := range params {
		filter := &vpc.Filter{
			Name:   helper.String(k),
			Values: []*string{helper.String(v)},
		}
		request.Filters = append(request.Filters, filter)
	}
	offset := uint64(0)
	request.Offset = &offset
	result := make([]*vpc.HaVip, 0)
	limit := uint64(HAVIP_DESCRIBE_LIMIT)
	request.Limit = &limit
	for {
		var response *vpc.DescribeHaVipsResponse
		err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
			result, e := meta.(*TencentCloudClient).apiV3Conn.UseVpcClient().DescribeHaVips(request)
			if e != nil {
				return retryError(errors.WithStack(e))
			}
			response = result
			return nil
		})
		if err != nil {
			log.Printf("[CRITAL]%s read HA VIP failed, reason:%+v", logId, err)
			return err
		} else {
			result = append(result, response.Response.HaVipSet...)
			if len(response.Response.HaVipSet) < HAVIP_DESCRIBE_LIMIT {
				break
			} else {
				offset = offset + limit
			}
		}
	}
	ids := make([]string, 0, len(result))
	haVipList := make([]map[string]interface{}, 0, len(result))
	for _, haVip := range result {
		mapping := map[string]interface{}{
			"id":          *haVip.HaVipId,
			"vip":         *haVip.Vip,
			"name":        *haVip.HaVipName,
			"state":       *haVip.State,
			"vpc_id":      *haVip.VpcId,
			"subnet_id":   *haVip.SubnetId,
			"create_time": *haVip.CreatedTime,
		}
		if haVip.NetworkInterfaceId != nil {
			mapping["network_interface_id"] = *haVip.NetworkInterfaceId
		}
		if haVip.AddressIp != nil {
			mapping["address_ip"] = *haVip.AddressIp
		}
		if haVip.InstanceId != nil {
			mapping["instance_id"] = *haVip.InstanceId
		}
		haVipList = append(haVipList, mapping)
		ids = append(ids, *haVip.HaVipId)
	}
	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("ha_vip_list", haVipList); e != nil {
		log.Printf("[CRITAL]%s provider set haVip list fail, reason:%s\n", logId, e)
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), haVipList); e != nil {
			return e
		}
	}

	return nil

}

/*
Use this data source to query VPC Network ACL information.

# Example Usage

```hcl
data "tencentcloudenterprise_vpc_instances" "foo" {
}

	data "tencentcloudenterprise_vpc_acls" "foo" {
	  vpc_id            = data.cloud_vpc_instances.foo.instance_list.0.vpc_id
	}

	data "tencentcloudenterprise_vpc_acls" "foo" {
	  name            	= "test_acl"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_acls", CNDescription{
		TerraformTypeCN: "私有网络ACL",
		DescriptionCN:   "提供VPC网络ACL数据源，用于查询VPC网络ACL的详细信息。",
		AttributesCN: map[string]string{
			"vpc_id":             "VPC实例的ID",
			"name":               "网络ACL的名称",
			"id":                 "网络ACL实例的ID",
			"result_output_file": "用于保存结果",
			"acl_list":           "专有网络的信息列表每个元素都包含以下属性：",
			"create_time":        "创建时间",
			"subnets":            "与网络ACL关联的子网",
			"subnet_id":          "子网实例ID",
			"subnet_name":        "子网名称",
			"cidr_block":         "子网的IPv4 CIDR",
			"tags":               "子网的标记",
			"ingress":            "网络ACL的入站规则",
			"protocol":           "IP协议的类型",
			"port":               "端口的范围",
			"policy":             "网络ACL的规则策略",
			"description":        "规则说明",
			"egress":             "网络ACL的出站规则",
		},
	})
}

func dataSourceTencentCloudVpcAcls() *schema.Resource {
	return &schema.Resource{
		Description:        "Use this data source to query VPC Network ACL information.",
		DeprecationMessage: "This data source is deprecated. Please use cloud_vpc_acl_list instead.",
		Read:               dataSourceTencentCloudVpcACLRead,

		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateNotEmpty,
				Description:  "ID of the VPC instance.",
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateStringLengthInRange(0, 60),
				Description:  "Name of the network ACL.",
			},
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateNotEmpty,
				Description:  "ID of the network ACL instance.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"acl_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The information list of the VPC. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the VPC instance.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the network ACL instance.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the network ACL.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"subnets": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Subnets associated with the network ACL.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the VPC instance.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Subnet instance ID.",
									},
									"subnet_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Subnet name.",
									},
									"cidr_block": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IPv4 CIDR of the subnet.",
									},
									"tags": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Tags of the subnet.",
									},
								},
							},
						},
						"ingress": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Inbound rules of the network ACL.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of IP protocol.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Range of the port.",
									},
									"policy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rule policy of Network ACL.",
									},
									"cidr_block": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An IP address network or segment.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rule description.",
									},
								},
							},
						},
						"egress": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Outbound rules of the network ACL.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of IP protocol.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Range of the port.",
									},
									"policy": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rule policy of Network ACL.",
									},
									"cidr_block": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "An IP address network or segment.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rule description.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudVpcACLRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_vpc_acls.read")()
	var (
		logId   = getLogId(contextNil)
		ctx     = context.WithValue(context.TODO(), logIdKey, logId)
		service = VpcService{client: meta.(*TencentCloudClient).apiV3Conn}

		vpcID = d.Get("vpc_id").(string)
		name  = d.Get("name").(string)
		id    = d.Get("id").(string)
	)

	networkAcls, err := service.DescribeNetWorkAcls(ctx, id, vpcID, name)
	if err != nil {
		return err
	}

	aclList := make([]map[string]interface{}, 0, len(networkAcls))
	ids := make([]string, 0, len(networkAcls))

	for _, info := range networkAcls {
		subnetInfo := info.SubnetSet
		subnets := make([]map[string]interface{}, 0, len(subnetInfo))
		for i := range subnetInfo {
			v := subnetInfo[i]
			subnet := make(map[string]interface{}, 5)
			subnet["vpc_id"] = v.VpcId
			subnet["subnet_id"] = v.SubnetId
			subnet["subnet_name"] = v.SubnetName
			subnet["cidr_block"] = v.CidrBlock

			tag := make(map[string]interface{}, len(v.TagSet))
			for t := range v.TagSet {
				tagValue := v.TagSet[t]
				tag[*tagValue.Key] = tagValue.Value
			}
			subnet["tags"] = tag

			subnets = append(subnets, subnet)
		}

		ingressInfo := info.IngressEntries
		ingress := make([]map[string]interface{}, 0, len(ingressInfo))
		for i := range ingressInfo {
			v := ingressInfo[i]
			egressMap := map[string]interface{}{
				"protocol":    v.Protocol,
				"port":        v.Port,
				"cidr_block":  v.CidrBlock,
				"policy":      v.Action,
				"description": v.Description,
			}
			ingress = append(ingress, egressMap)
		}

		egressInfo := info.EgressEntries
		egress := make([]map[string]interface{}, 0, len(egressInfo))
		for i := range egressInfo {
			v := egressInfo[i]
			egressMap := map[string]interface{}{
				"protocol":    v.Protocol,
				"port":        v.Port,
				"cidr_block":  v.CidrBlock,
				"policy":      v.Action,
				"description": v.Description,
			}
			egress = append(egress, egressMap)
		}

		aclResult := map[string]interface{}{
			"vpc_id":      info.VpcId,
			"id":          info.NetworkAclId,
			"name":        info.NetworkAclName,
			"create_time": info.CreatedTime,
			"subnets":     subnets,
			"ingress":     ingress,
			"egress":      egress,
		}
		aclList = append(aclList, aclResult)
		ids = append(ids, *info.NetworkAclId)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	err = d.Set("acl_list", aclList)
	if err != nil {
		log.Printf("[CRITAL]%s provider set ACL list fail, reason:%v \n ", logId, err)
		return err
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if err := writeToFile(output.(string), aclList); err != nil {
			return err
		}
	}
	return nil
}

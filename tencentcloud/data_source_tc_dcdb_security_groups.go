/*
Use this data source to query detailed information of dcdb securityGroups

# Example Usage

```hcl

	data "tencentcloudenterprise_dcdb_security_groups" "securityGroups" {
	  instance_id = "your_instance_id"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	dcdb "terraform-provider-tencentcloudenterprise/sdk/dcdb/v20180411"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_dcdb_security_groups", CNDescription{
		TerraformTypeCN: "数据库安全组",
		DescriptionCN:   "提供DCDB安全组数据源，用于查询DCDB安全组的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":         "实例id",
			"list":                "安全组列表",
			"project_id":          "项目id",
			"create_time":         "创造时间",
			"security_group_id":   "安全组id",
			"security_group_name": "安全组名称",
			"inbound":             "入站规则",
			"cidr_ip":             "cidr ip",
			"action":              "政策行动",
			"port_range":          "端口范围",
			"ip_protocol":         "互联网协议",
			"outbound":            "出站规则",
			"result_output_file":  "用于保存结果",
		},
	})
}

func dataSourceTencentCloudDcdbSecurityGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Cloud Database Security Groups",
		Read:        dataSourceTencentCloudDcdbSecurityGroupsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance id.",
			},

			"list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Security group list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Project id.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Create time.",
						},
						"security_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Security group id.",
						},
						"security_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Security group name.",
						},
						"inbound": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Inbound rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cidr_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cidr ip.",
									},
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Policy action.",
									},
									"port_range": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Port range.",
									},
									"ip_protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Internet protocol.",
									},
								},
							},
						},
						"outbound": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Outbound rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cidr_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cidr ip.",
									},
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Policy action.",
									},
									"port_range": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Port range.",
									},
									"ip_protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Internet protocol.",
									},
								},
							},
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

func dataSourceTencentCloudDcdbSecurityGroupsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_dcdb_security_groups.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("instance_id"); ok {
		paramMap["instance_id"] = helper.String(v.(string))
	}

	dcdbService := DcdbService{client: meta.(*TencentCloudClient).apiV3Conn}

	var groups []*dcdb.SecurityGroup
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := dcdbService.DescribeDcdbSecurityGroupsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		groups = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read Dcdb groups failed, reason:%+v", logId, err)
		return err
	}

	groupList := []interface{}{}
	ids := make([]string, 0, len(groups))
	if groups != nil {
		for _, group := range groups {
			groupMap := map[string]interface{}{}
			if group.ProjectId != nil {
				groupMap["project_id"] = group.ProjectId
			}
			if group.CreateTime != nil {
				groupMap["create_time"] = group.CreateTime
			}
			if group.SecurityGroupId != nil {
				groupMap["security_group_id"] = group.SecurityGroupId
			}
			if group.SecurityGroupName != nil {
				groupMap["security_group_name"] = group.SecurityGroupName
			}
			if group.Inbound != nil {
				inboundList := []interface{}{}
				for _, inbound := range group.Inbound {
					inboundMap := map[string]interface{}{}
					if inbound.CidrIp != nil {
						inboundMap["cidr_ip"] = inbound.CidrIp
					}
					if inbound.Action != nil {
						inboundMap["action"] = inbound.Action
					}
					if inbound.PortRange != nil {
						inboundMap["port_range"] = inbound.PortRange
					}
					if inbound.IpProtocol != nil {
						inboundMap["ip_protocol"] = inbound.IpProtocol
					}

					inboundList = append(inboundList, inboundMap)
				}
				groupMap["inbound"] = inboundList
			}
			if group.Outbound != nil {
				outboundList := []interface{}{}
				for _, outbound := range group.Outbound {
					outboundMap := map[string]interface{}{}
					if outbound.CidrIp != nil {
						outboundMap["cidr_ip"] = outbound.CidrIp
					}
					if outbound.Action != nil {
						outboundMap["action"] = outbound.Action
					}
					if outbound.PortRange != nil {
						outboundMap["port_range"] = outbound.PortRange
					}
					if outbound.IpProtocol != nil {
						outboundMap["ip_protocol"] = outbound.IpProtocol
					}

					outboundList = append(outboundList, outboundMap)
				}
				groupMap["outbound"] = outboundList
			}
			ids = append(ids, *group.SecurityGroupId)
			groupList = append(groupList, groupMap)
		}
		d.SetId(helper.DataResourceIdsHash(ids))
		_ = d.Set("list", groupList)
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), groupList); e != nil {
			return e
		}
	}

	return nil
}

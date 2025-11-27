/*
Use this data source to query detailed information of security groups.

# Example Usage

```hcl

	data "tencentcloudenterprise_vpc_security_groups" "sglab" {
	  security_group_id = cloud_vpc_security_group.sglab.id
	}

```
*/
package tencentcloud

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	vpc "terraform-provider-tencentcloudenterprise/sdk/vpc/v20170312"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_vpc_security_groups", CNDescription{
		TerraformTypeCN: "安全组",
		DescriptionCN:   "提供VPC安全组数据源，用于查询VPC安全组的详细信息。",
		AttributesCN: map[string]string{
			"security_group_id":  "要查询的安全组的ID与“name”和“project_id”冲突",
			"name":               "要查询的安全组的名称与“security_group_id”冲突",
			"project_id":         "要查询的安全组的项目ID与“security_group_id”冲突",
			"tags":               "要查询的安全组的标记与“security_group_id”冲突",
			"result_output_file": "用于保存结果",
			"security_groups":    "安全组的信息列表",
			"description":        "安全组的描述",
			"create_time":        "安全组的创建时间",
			"be_associate_count": "安全组绑定资源的数目",
			"ingress":            "入口规则集对于`[action]#[cidr_ip]#[port]#[protocol]`之类的项，它表示一个规则；对于像“sg XXXX”这样的项目，它意味着嵌套的安全组",
			"egress":             "设置出口规则对于`[action]#[cidr_ip]#[port]#[protocol]`之类的项，它表示一个规则；对于像“sg XXXX”这样的项目，它意味着嵌套的安全组",
		},
	})
}

func dataSourceTencentCloudSecurityGroups() *schema.Resource {
	return &schema.Resource{
		Description: "Query Security Groups",
		Read:        dataSourceTencentCloudSecurityGroupsRead,
		Schema: map[string]*schema.Schema{
			"security_group_id": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"name", "project_id"},
				Description:   "ID of the security group to be queried. Conflict with `name` and `project_id`.",
			},
			"name": {
				Type:          schema.TypeString,
				Optional:      true,
				ValidateFunc:  validateStringLengthInRange(1, 60),
				ConflictsWith: []string{"security_group_id"},
				Description:   "Name of the security group to be queried. Conflict with `security_group_id`.",
			},
			"project_id": {
				Type:          schema.TypeInt,
				Optional:      true,
				ConflictsWith: []string{"security_group_id"},
				Description:   "Project ID of the security group to be queried. Conflict with `security_group_id`.",
			},
			"tags": {
				Type:          schema.TypeMap,
				Optional:      true,
				ConflictsWith: []string{"security_group_id"},
				Description:   "Tags of the security group to be queried. Conflict with `security_group_id`.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},

			// computed
			"security_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information list of security group.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"security_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the security group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the security group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the security group.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of security group.",
						},
						"be_associate_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of security group binding resources.",
						},
						"project_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Project ID of the security group.",
						},
						"ingress": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Ingress rules set. For items like `[action]#[cidr_ip]#[port]#[protocol]`, it means a regular rule; for items like `sg-XXXX`, it means a nested security group.",
						},
						"egress": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Egress rules set. For items like `[action]#[cidr_ip]#[port]#[protocol]`, it means a regular rule; for items like `sg-XXXX`, it means a nested security group.",
						},
						"tags": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Tags of the security group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudSecurityGroupsRead(d *schema.ResourceData, m interface{}) error {
	defer logElapsed("data_source.cloud_vpc_security_groups.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	client := m.(*TencentCloudClient).apiV3Conn
	vpcService := VpcService{client: client}
	tagService := TagService{client: client}
	region := client.Region

	var (
		sgId      *string
		sgName    *string
		projectId *int
	)

	idBuilder := strings.Builder{}
	idBuilder.WriteString("securityGroups-")

	if raw, ok := d.GetOk("security_group_id"); ok {
		sgId = helper.String(raw.(string))
		idBuilder.WriteString(*sgId)
		idBuilder.WriteRune('-')
	}

	if raw, ok := d.GetOk("name"); ok {
		sgName = common.StringPtr(raw.(string))
		idBuilder.WriteString(*sgName)
		idBuilder.WriteRune('-')
	}

	if raw, ok := d.GetOkExists("project_id"); ok {
		projectId = helper.Int(raw.(int))
		idBuilder.WriteString(strconv.Itoa(*projectId))
	}

	tags := helper.GetTags(d, "tags")

	sgs, err := vpcService.DescribeSecurityGroups(ctx, sgId, sgName, projectId, tags)
	if err != nil {
		return err
	}

	if len(sgs) == 0 {
		_ = d.Set("security_groups", []map[string]interface{}{})
		d.SetId(idBuilder.String())
		return nil
	}

	sgMap := make(map[string]*vpc.SecurityGroup, len(sgs))
	for _, sg := range sgs {
		if sg.SecurityGroupId == nil {
			return errors.New("security group id is nil")
		}
		sgMap[*sg.SecurityGroupId] = sg
	}

	sgIds := make([]string, 0, len(sgs))
	for _, sg := range sgs {
		sgIds = append(sgIds, *sg.SecurityGroupId)
	}

	associateSet, err := vpcService.DescribeSecurityGroupsAssociate(ctx, sgIds)
	if err != nil {
		return err
	}

	sgInstances := make([]map[string]interface{}, 0, len(sgs))
	for _, associate := range associateSet {
		if associate.SecurityGroupId == nil {
			return errors.New("associate statistics security group id is nil")
		}

		if sg, ok := sgMap[*associate.SecurityGroupId]; ok {
			count := int(*associate.CVM + *associate.ENI + *associate.CDB + *associate.CLB)

			// normally projectId default value is 0
			if sg.ProjectId == nil {
				return errors.New("associate statistics project id is nil")
			}

			projectId, err := strconv.Atoi(*sg.ProjectId)
			if err != nil {
				return fmt.Errorf("securtiy group %s project id invalid: %v", *sg.SecurityGroupId, err)
			}

			respIngress, respEgress, exist, err := vpcService.DescribeSecurityGroupPolices(ctx, *sg.SecurityGroupId)
			if err != nil {
				return err
			}

			if !exist {
				// when read security group all rules, it doesn't exist, maybe delete on other places, ignore it
				continue
			}

			respTags, err := tagService.DescribeResourceTags(ctx, "cvm", "sg", region, *sg.SecurityGroupId)
			if err != nil {
				return err
			}

			ingress := make([]string, 0, len(respIngress))
			for _, in := range respIngress {
				ingress = append(ingress, in.String())
			}

			egress := make([]string, 0, len(respEgress))
			for _, eg := range respEgress {
				egress = append(egress, eg.String())
			}

			sgInstances = append(sgInstances, map[string]interface{}{
				"security_group_id":  *sg.SecurityGroupId,
				"name":               *sg.SecurityGroupName,
				"description":        *sg.SecurityGroupDesc,
				"create_time":        *sg.CreatedTime,
				"be_associate_count": count,
				"project_id":         projectId,
				"ingress":            ingress,
				"egress":             egress,
				"tags":               respTags,
			})
		}
	}

	if len(sgInstances) != len(sgs) {
		return errors.New("security group associate statistics is not enough")
	}

	_ = d.Set("security_groups", sgInstances)
	d.SetId(idBuilder.String())

	if output, ok := d.GetOk("result_output_file"); ok && output.(string) != "" {
		if err := writeToFile(output.(string), sgInstances); err != nil {
			log.Printf("[CRITAL]%s output file[%s] fail, reason[%v]", logId, output.(string), err)
			return err
		}
	}

	return nil
}

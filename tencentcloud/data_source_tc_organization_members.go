/*
Use this data source to query detailed information of organization members

# Example Usage

```hcl
data "tencentcloudenterprise_organization_members" "members" {
}

	output "members_list" {
	  value = data.tencentcloudenterprise_organization_members.members.items
	}

```
*/
package tencentcloud

import (
	"context"

	organization "terraform-provider-tencentcloudenterprise/sdk/organization/v20220508"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_organization_members", CNDescription{
		TerraformTypeCN: "组织成员列表",
		DescriptionCN:   "提供组织成员列表数据源，用于查询满足条件的组织成员信息。",
		AttributesCN: map[string]string{
			"lang":                "语言版本：en 英文版，zh 中文版",
			"search_key":          "搜索关键字，支持账号名称和 UIN",
			"auth_name":           "主体名称",
			"product":             "可信服务简称",
			"items":               "成员列表",
			"member_uin":          "成员账号 UIN",
			"name":                "成员名称",
			"member_type":         "成员类型：Invite 邀请加入，Create 创建",
			"org_policy_type":     "关系策略类型",
			"org_policy_name":     "关系策略名称",
			"org_permission":      "关系策略权限列表",
			"id":                  "权限 ID",
			"node_id":             "节点 ID",
			"node_name":           "节点名称",
			"remark":              "备注",
			"create_time":         "创建时间",
			"update_time":         "更新时间",
			"is_allow_quit":       "是否允许退出：Allow，Denied",
			"pay_uin":             "代付账号 UIN",
			"pay_name":            "代付账号名称",
			"org_identity":        "管理身份列表",
			"identity_id":         "身份 ID",
			"identity_alias_name": "身份别名",
			"permission_status":   "成员权限状态：Confirmed，UnConfirmed",
			"result_output_file":  "用于保存结果，可视化界面不可用",
		},
	})
}

func dataSourceTencentCloudOrganizationMembers() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudOrganizationMembersRead,
		Description: "Use this data source to query detailed information of organization members.",
		Schema: map[string]*schema.Schema{
			"lang": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Valid values: `en` (Tencent Cloud International); `zh` (Tencent Cloud).",
			},

			"search_key": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Search by member name or ID.",
			},

			"auth_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Entity name.",
			},

			"product": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Abbreviation of the trusted service, which is required during querying the trusted service admin.",
			},

			"items": {
				Computed:    true,
				Type:        schema.TypeList,
				Description: "Member list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_uin": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Member UIN.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Member name.",
						},
						"member_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Member type. Valid values: `Invite` (invited); `Create` (created).",
						},
						"org_policy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Relationship policy type.",
						},
						"org_policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Relationship policy name.",
						},
						"org_permission": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Relationship policy permission.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Permission ID.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Permission name.",
									},
								},
							},
						},
						"node_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Node ID.",
						},
						"node_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Node name.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remarks.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
						},
						"is_allow_quit": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether the member is allowed to leave. Valid values: `Allow`, `Denied`.",
						},
						"pay_uin": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Payer UIN.",
						},
						"pay_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Payer name.",
						},
						"org_identity": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Management identity.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"identity_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Identity ID.",
									},
									"identity_alias_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Identity name.",
									},
								},
							},
						},
						"permission_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Member permission status. Valid values: `Confirmed`, `UnConfirmed`.",
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

func dataSourceTencentCloudOrganizationMembersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_organization_members.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("lang"); ok {
		paramMap["Lang"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("search_key"); ok {
		paramMap["SearchKey"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("auth_name"); ok {
		paramMap["AuthName"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("product"); ok {
		paramMap["Product"] = helper.String(v.(string))
	}

	service := OrganizationService{client: meta.(*TencentCloudClient).apiV3Conn}

	var items []*organization.OrgMember
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeOrganizationMembersByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		items = result
		return nil
	})
	if err != nil {
		return err
	}

	ids := make([]string, 0, len(items))
	tmpList := make([]map[string]interface{}, 0, len(items))

	if items != nil {
		for _, orgMember := range items {
			orgMemberMap := map[string]interface{}{}

			if orgMember.MemberUin != nil {
				orgMemberMap["member_uin"] = orgMember.MemberUin
			}
			if orgMember.Name != nil {
				orgMemberMap["name"] = orgMember.Name
			}
			if orgMember.MemberType != nil {
				orgMemberMap["member_type"] = orgMember.MemberType
			}
			if orgMember.OrgPolicyType != nil {
				orgMemberMap["org_policy_type"] = orgMember.OrgPolicyType
			}
			if orgMember.OrgPolicyName != nil {
				orgMemberMap["org_policy_name"] = orgMember.OrgPolicyName
			}

			if orgMember.OrgPermission != nil {
				orgPermissionList := []interface{}{}
				for _, orgPermission := range orgMember.OrgPermission {
					orgPermissionMap := map[string]interface{}{}
					if orgPermission.Id != nil {
						orgPermissionMap["id"] = orgPermission.Id
					}
					if orgPermission.Name != nil {
						orgPermissionMap["name"] = orgPermission.Name
					}
					orgPermissionList = append(orgPermissionList, orgPermissionMap)
				}
				orgMemberMap["org_permission"] = orgPermissionList
			}

			if orgMember.NodeId != nil {
				orgMemberMap["node_id"] = orgMember.NodeId
			}
			if orgMember.NodeName != nil {
				orgMemberMap["node_name"] = orgMember.NodeName
			}
			if orgMember.Remark != nil {
				orgMemberMap["remark"] = orgMember.Remark
			}
			if orgMember.CreateTime != nil {
				orgMemberMap["create_time"] = orgMember.CreateTime
			}
			if orgMember.UpdateTime != nil {
				orgMemberMap["update_time"] = orgMember.UpdateTime
			}
			if orgMember.IsAllowQuit != nil {
				orgMemberMap["is_allow_quit"] = orgMember.IsAllowQuit
			}
			if orgMember.PayUin != nil {
				orgMemberMap["pay_uin"] = orgMember.PayUin
			}
			if orgMember.PayName != nil {
				orgMemberMap["pay_name"] = orgMember.PayName
			}

			if orgMember.OrgIdentity != nil {
				orgIdentityList := []interface{}{}
				for _, orgIdentity := range orgMember.OrgIdentity {
					orgIdentityMap := map[string]interface{}{}
					if orgIdentity.IdentityId != nil {
						orgIdentityMap["identity_id"] = orgIdentity.IdentityId
					}
					if orgIdentity.IdentityAliasName != nil {
						orgIdentityMap["identity_alias_name"] = orgIdentity.IdentityAliasName
					}
					orgIdentityList = append(orgIdentityList, orgIdentityMap)
				}
				orgMemberMap["org_identity"] = orgIdentityList
			}

			if orgMember.PermissionStatus != nil {
				orgMemberMap["permission_status"] = orgMember.PermissionStatus
			}

			if orgMember.Name != nil {
				ids = append(ids, *orgMember.Name)
			}
			tmpList = append(tmpList, orgMemberMap)
		}

		_ = d.Set("items", tmpList)
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

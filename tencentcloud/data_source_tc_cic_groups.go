/*
Use this data source to query detailed information of cic groups

# Example Usage

```hcl

	data "tencentcloudenterprise_cic_groups" "groups" {
	  zone_id = "z-xxxxxxxxxx"
	}

	output "groups_list" {
	  value = data.tencentcloudenterprise_cic_groups.groups.groups
	}

```
*/
package tencentcloud

import (
	"context"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cic_groups", CNDescription{
		TerraformTypeCN: "CIC 用户组列表",
		DescriptionCN:   "提供 CIC 用户组列表数据源，用于查询满足条件的用户组信息。",
		AttributesCN: map[string]string{
			"zone_id":            "空间 ID",
			"filter":             "过滤条件，格式为 <Attribute> <Operator> <Value>",
			"group_type":         "用户组类型：Manual 手动创建，Synchronized 外部导入",
			"filter_users":       "筛选的用户列表，关联该用户的用户组会返回 is_selected=true",
			"sort_field":         "排序字段，目前仅支持 CreateTime",
			"sort_type":          "排序类型：Desc 倒序，Asc 正序",
			"groups":             "用户组列表",
			"group_name":         "用户组名称",
			"description":        "用户组描述",
			"create_time":        "用户组创建时间",
			"group_type_out":     "用户组类型",
			"update_time":        "用户组修改时间",
			"group_id":           "用户组 ID",
			"member_count":       "组员数量",
			"is_selected":        "若 filter_users 有值，用户在该用户组中则返回 true，否则 false",
			"result_output_file": "用于保存结果，可视化界面不可用",
		},
	})
}

func dataSourceTencentCloudCicGroups() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudCicGroupsRead,
		Description: "Use this data source to query detailed information of cic groups.",
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Space ID.",
			},

			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter criterion. Format: <Attribute> <Operator> <Value>, case-insensitive. Currently, <Attribute> supports only GroupName, and <Operator> supports only eq (Equals) and sw (Start With).",
			},

			"group_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User group type. Manual: manually created; Synchronized: externally imported.",
			},

			"filter_users": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Filtered user list. IsSelected=true will be returned for the user group associated with this user.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"sort_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sorting field, which currently only supports CreateTime.",
			},

			"sort_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sorting type. Desc: descending order; Asc: ascending order. It should be set along with SortField.",
			},

			"groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User group list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User group name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User group description.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the user group.",
						},
						"group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User group type. Manual: manually created; Synchronized: externally imported.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Modification time of the user group.",
						},
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User group ID.",
						},
						"member_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of group members.",
						},
						"is_selected": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If the input parameter FilterUsers is provided, return true when the user is in the user group; otherwise, return false.",
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

func dataSourceTencentCloudCicGroupsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cic_groups.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("zone_id"); ok {
		paramMap["ZoneId"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("filter"); ok {
		paramMap["Filter"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("group_type"); ok {
		paramMap["GroupType"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("filter_users"); ok {
		filterUsersList := []*string{}
		filterUsersSet := v.(*schema.Set).List()
		for i := range filterUsersSet {
			filterUsers := filterUsersSet[i].(string)
			filterUsersList = append(filterUsersList, helper.String(filterUsers))
		}
		paramMap["FilterUsers"] = filterUsersList
	}
	if v, ok := d.GetOk("sort_field"); ok {
		paramMap["SortField"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("sort_type"); ok {
		paramMap["SortType"] = helper.String(v.(string))
	}

	var groups []*cic.GroupInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCicGroupsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		groups = result
		return nil
	})
	if err != nil {
		return err
	}

	groupsList := make([]map[string]interface{}, 0, len(groups))
	ids := make([]string, 0, len(groups))
	for _, group := range groups {
		groupsMap := map[string]interface{}{}
		if group.GroupName != nil {
			groupsMap["group_name"] = group.GroupName
		}
		if group.Description != nil {
			groupsMap["description"] = group.Description
		}
		if group.CreateTime != nil {
			groupsMap["create_time"] = group.CreateTime
		}
		if group.GroupType != nil {
			groupsMap["group_type"] = group.GroupType
		}
		if group.UpdateTime != nil {
			groupsMap["update_time"] = group.UpdateTime
		}
		if group.GroupId != nil {
			groupsMap["group_id"] = group.GroupId
			ids = append(ids, *group.GroupId)
		}
		if group.MemberCount != nil {
			groupsMap["member_count"] = group.MemberCount
		}
		if group.IsSelected != nil {
			groupsMap["is_selected"] = group.IsSelected
		}
		groupsList = append(groupsList, groupsMap)
	}

	_ = d.Set("groups", groupsList)
	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), groupsList); e != nil {
			return e
		}
	}

	return nil
}

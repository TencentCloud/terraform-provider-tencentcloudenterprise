/*
Use this data source to query detailed information of cic users

# Example Usage

```hcl

	data "tencentcloudenterprise_cic_users" "users" {
	  zone_id = "z-xxxxxxxxxx"
	  user_name = "admin"
	  user_status = "Enabled"
	}

	output "users_list" {
	  value = data.tencentcloudenterprise_cic_users.users.users
	}

*/
package tencentcloud

import (
	"context"
	"strings"

	cic "terraform-provider-tencentcloudenterprise/sdk/cic/v20210331"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cic_users", CNDescription{
		TerraformTypeCN: "CIC 用户列表",
		DescriptionCN:   "提供 CIC 用户列表数据源，用于查询满足条件的用户信息。",
		AttributesCN: map[string]string{
			"zone_id":            "空间 ID",
			"user_name":          "用户名，精确匹配",
			"user_id":            "用户 ID，精确匹配",
			"email":              "邮箱地址，精确匹配",
			"description":        "用户描述，精确匹配",
			"user_status":        "用户状态：Enabled 启用，Disabled 禁用",
			"user_type":          "用户类型：Manual 手动创建，Synchronized 外部导入",
			"users":              "用户列表",
			"display_name":       "显示名称",
			"first_name":         "名",
			"last_name":          "姓",
			"create_time":        "创建时间",
			"update_time":        "更新时间",
			"is_selected":        "若 filter_groups 有值，用户在该用户组中则返回 true，否则 false",
			"result_output_file": "用于保存结果，可视化界面不可用",
		},
	})
}

func dataSourceTencentCloudCicUsers() *schema.Resource {
	return &schema.Resource{
		Read:        dataSourceTencentCloudCicUsersRead,
		Description: "Use this data source to query detailed information of cic users.",
		Schema: map[string]*schema.Schema{
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Space ID.",
			},

			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User name for exact match filtering.",
			},

			"user_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User ID for exact match filtering.",
			},

			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email address for exact match filtering.",
			},

			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User description for exact match filtering.",
			},

			"user_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User status. Enabled: enabled; Disabled: disabled.",
			},

			"user_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User type. Manual: manually created; Synchronized: externally imported.",
			},

			"users": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User ID.",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User name.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Display name.",
						},
						"first_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "First name.",
						},
						"last_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last name.",
						},
						"email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Email address.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description.",
						},
						"user_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User status. Enabled: enabled; Disabled: disabled.",
						},
						"user_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User type. Manual: manually created; Synchronized: externally imported.",
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
						"is_selected": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If the input parameter FilterGroups is provided, return true when the user is in the user group; otherwise, return false.",
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

func dataSourceTencentCloudCicUsersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cic_users.read")()
	defer inconsistentCheck(d, meta)()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := CicService{client: meta.(*TencentCloudClient).apiV3Conn}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("zone_id"); ok {
		paramMap["ZoneId"] = helper.String(v.(string))
	}

	// Build Filter string from individual search fields for API-side fuzzy matching
	filterParts := []string{}
	if v, ok := d.GetOk("user_name"); ok {
		filterParts = append(filterParts, v.(string))
	}
	if v, ok := d.GetOk("user_id"); ok {
		filterParts = append(filterParts, v.(string))
	}
	if v, ok := d.GetOk("email"); ok {
		filterParts = append(filterParts, v.(string))
	}
	if v, ok := d.GetOk("description"); ok {
		filterParts = append(filterParts, v.(string))
	}
	if len(filterParts) > 0 {
		paramMap["Filter"] = helper.String(strings.Join(filterParts, " "))
	}

	if v, ok := d.GetOk("user_status"); ok {
		paramMap["UserStatus"] = helper.String(v.(string))
	}
	if v, ok := d.GetOk("user_type"); ok {
		paramMap["UserType"] = helper.String(v.(string))
	}

	var users []*cic.UserInfo
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCicUsersByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		users = result
		return nil
	})
	if err != nil {
		return err
	}

	// Client-side exact match filtering on individual fields
	filterUserName, hasUserName := d.GetOk("user_name")
	filterUserId, hasUserId := d.GetOk("user_id")
	filterEmail, hasEmail := d.GetOk("email")
	filterDesc, hasDesc := d.GetOk("description")

	filtered := make([]*cic.UserInfo, 0, len(users))
	for _, user := range users {
		if hasUserName && (user.UserName == nil || *user.UserName != filterUserName.(string)) {
			continue
		}
		if hasUserId && (user.UserId == nil || *user.UserId != filterUserId.(string)) {
			continue
		}
		if hasEmail && (user.Email == nil || *user.Email != filterEmail.(string)) {
			continue
		}
		if hasDesc && (user.Description == nil || *user.Description != filterDesc.(string)) {
			continue
		}
		filtered = append(filtered, user)
	}
	users = filtered

	usersList := make([]map[string]interface{}, 0, len(users))
	ids := make([]string, 0, len(users))
	for _, user := range users {
		usersMap := map[string]interface{}{}
		if user.UserId != nil {
			usersMap["user_id"] = user.UserId
			ids = append(ids, *user.UserId)
		}
		if user.UserName != nil {
			usersMap["user_name"] = user.UserName
		}
		if user.DisplayName != nil {
			usersMap["display_name"] = user.DisplayName
		}
		if user.FirstName != nil {
			usersMap["first_name"] = user.FirstName
		}
		if user.LastName != nil {
			usersMap["last_name"] = user.LastName
		}
		if user.Email != nil {
			usersMap["email"] = user.Email
		}
		if user.Description != nil {
			usersMap["description"] = user.Description
		}
		if user.UserStatus != nil {
			usersMap["user_status"] = user.UserStatus
		}
		if user.UserType != nil {
			usersMap["user_type"] = user.UserType
		}
		if user.CreateTime != nil {
			usersMap["create_time"] = user.CreateTime
		}
		if user.UpdateTime != nil {
			usersMap["update_time"] = user.UpdateTime
		}
		if user.IsSelected != nil {
			usersMap["is_selected"] = user.IsSelected
		}
		usersList = append(usersList, usersMap)
	}

	_ = d.Set("users", usersList)
	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), usersList); e != nil {
			return e
		}
	}

	return nil
}

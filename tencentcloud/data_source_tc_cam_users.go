/*
Use this data source to query detailed information of CAM users

# Example Usage

```hcl
# query by name

	data "tencentcloudenterprise_cam_users" "foo" {
	  name = "cam-user-test"
	}

# query by email

	data "tencentcloudenterprise_cam_users" "bar" {
	  email = "hello@test.com"
	}

# query by phone

	data "tencentcloudenterprise_cam_users" "far" {
	  phone_num = "12345678910"
	}

```
*/
package tencentcloud

import (
	"context"
	"log"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_users", CNDescription{
		TerraformTypeCN: "CAM用户列表",
		DescriptionCN:   "用于查询 CAM 用户信息的列表。",
		AttributesCN: map[string]string{
			"name":          "子账号名称",
			"remark":        "备注",
			"phone_num":     "手机号码",
			"country_code":  "国家代码",
			"email":         "邮箱",
			"uin":           "帐号 UIN",
			"uid":           "子用户 UID",
			"console_login": "是否允许登录控制台",
			"user_list":     "用户列表",
			"user_id":       "子用户 UID",
			"can_login":     "是否允许登录控制台",
		},
	})
}

func dataSourceTencentCloudCamUsers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamUsersRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name of CAM user to be queried.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Remark of the CAM user to be queried.",
			},
			"phone_num": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Phone num of the CAM user to be queried.",
			},
			"country_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Country code of the CAM user to be queried.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email of the CAM user to be queried.",
			},
			"uin": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Uin of the CAM user to be queried.",
			},
			"uid": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Uid of the CAM user to be queried.",
			},
			"console_login": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicate whether the user can login in.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"user_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of CAM users. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of CAM user. Its value equals to `name` argument.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of CAM user.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Remark of the CAM user.",
						},
						"phone_num": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Phone num of the CAM user.",
						},
						"country_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Country code of the CAM user.",
						},
						"email": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Email of the CAM user.",
						},
						"uin": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Uin of the CAM user.",
						},
						"uid": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Uid of the CAM user.",
						},
						"console_login": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicate whether the user can login in.",
						},
						"can_login": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicate whether the user can login in.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCamUsersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencentcloudenterprise_cam_users.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		params["name"] = v.(string)
	}
	if v, ok := d.GetOk("uin"); ok {
		params["uin"] = v.(int)
	}
	if v, ok := d.GetOk("remark"); ok {
		params["remark"] = v.(string)
	}
	if v, ok := d.GetOk("uid"); ok {
		params["uid"] = v.(int)
	}
	if v, ok := d.GetOk("phone_num"); ok {
		params["phone_num"] = v.(string)
	}
	if v, ok := d.GetOk("country_code"); ok {
		params["country_code"] = v.(string)
	}
	if v, ok := d.GetOk("email"); ok {
		params["email"] = v.(string)
	}
	if v, ok := d.GetOkExists("console_login"); ok {
		consoleLogin := v.(bool)
		if consoleLogin {
			params["console_login"] = 1
		} else {
			params["console_login"] = 0
		}
	}

	camService := CamService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	var users []*cam.SubAccountFilter
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		results, e := camService.DescribeUsersByFilter(ctx, params)
		if e != nil {
			return retryError(e)
		}
		users = results
		return nil
	})
	if err != nil {
		log.Printf("[CRITAL]%s read CAM users failed, reason:%s\n", logId, err.Error())
		return err
	}
	userList := make([]map[string]interface{}, 0, len(users))
	ids := make([]string, 0, len(users))
	for _, user := range users {
		mapping := map[string]interface{}{}
		if user.Uin != nil {
			mapping["uin"] = int(*user.Uin)
		}
		if user.Uid != nil {
			mapping["uid"] = int(*user.Uid)
		}
		if user.Name != nil {
			mapping["name"] = *user.Name
			mapping["user_id"] = *user.Name
			ids = append(ids, *user.Name)
		}
		if user.Remark != nil {
			mapping["remark"] = *user.Remark
		}
		if user.PhoneNum != nil {
			mapping["phone_num"] = *user.PhoneNum
		}
		if user.CountryCode != nil {
			mapping["country_code"] = *user.CountryCode
		}
		if user.Email != nil {
			mapping["email"] = *user.Email
		}
		if user.ConsoleLogin != nil {
			mapping["console_login"] = (*user.ConsoleLogin == 1)
		}
		if user.CanLogin != nil {
			mapping["can_login"] = (*user.CanLogin == 1)
		}

		userList = append(userList, mapping)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	if e := d.Set("user_list", userList); e != nil {
		log.Printf("[CRITAL]%s provider set CAM user list fail, reason:%s\n", logId, e.Error())
		return e
	}

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), userList); e != nil {
			return e
		}
	}

	return nil
}

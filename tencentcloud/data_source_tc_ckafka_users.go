/*
Use this data source to query detailed user information of Ckafka

# Example Usage

```hcl

	data "cloud_ckafka_users" "foo" {
	  instance_id  = "ckafka-f9ife4zz"
	  account_name = "test"
	}

```
*/
package tencentcloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("cloud_ckafka_users", CNDescription{
		TerraformTypeCN: "ckafka用户实例",
		DescriptionCN:   "提供CKafka用户实例数据源，用于查询CKafka用户的详细信息。",
		AttributesCN: map[string]string{
			"instance_id":        "ckafka实例的Id",
			"account_name":       "查询ckafka用户信息时使用的帐户名称可能是用户名的子字符串",
			"result_output_file": "用于保存结果",
			"user_list":          "ckafka用户列表每个元素都包含以下属性：",
			"create_time":        "帐户的创建时间",
			"update_time":        "帐户的上次更新时间",
		},
	})
}

func dataSourceTencentCloudCkafkaUsers() *schema.Resource {
	return &schema.Resource{
		Description: "Use this data source to query detailed user information of Ckafka",
		Read:        dataSourceTencentCloudCkafkaUsersRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the ckafka instance.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account name used when query ckafka users' infos. Could be a substr of user name.",
			},
			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
			"user_list": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of ckafka users. Each element contains the following attributes:",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account name of user.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the account.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of the account.",
						},
					},
				},
			},
		},
	}
}

func dataSourceTencentCloudCkafkaUsersRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.cloud_ckafka_users.read")()

	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	params := make(map[string]interface{})
	params["instance_id"] = d.Get("instance_id").(string)
	if v, ok := d.GetOk("account_name"); ok {
		params["account_name"] = v.(string)
	}

	ckafkaService := CkafkaService{
		client: meta.(*TencentCloudClient).apiV3Conn,
	}
	userInfos, err := ckafkaService.DescribeUserByFilter(ctx, params)
	if err != nil {
		return err
	}
	userList := make([]map[string]interface{}, 0, len(userInfos))
	ids := make([]string, 0, len(userInfos))
	for _, user := range userInfos {
		userList = append(userList, map[string]interface{}{
			"account_name": *user.Name,
			"create_time":  *user.CreateTime,
			"update_time":  *user.UpdateTime,
		})

		ids = append(ids, params["instance_id"].(string)+FILED_SP+*user.Name)
	}

	d.SetId(helper.DataResourceIdsHash(ids))
	_ = d.Set("user_list", userList)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), userList); e != nil {
			return e
		}
	}

	return nil
}

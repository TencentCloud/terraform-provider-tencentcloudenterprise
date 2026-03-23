/*
Use this data source to query CAM sub accounts.

# Example Usage

```hcl

	data "tencentcloudenterprise_cam_sub_accounts" "foo" {
	  filter_sub_account_uin = [100000000001]
	}

```
*/
package tencentcloud

import (
	"context"

	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_cam_sub_accounts", CNDescription{
		TerraformTypeCN: "CAM子账号",
		DescriptionCN:   "用于查询 CAM 子账号信息的列表。",
		AttributesCN: map[string]string{
			"filter_sub_account_uin": "子账号 UIN 列表，最大支持 50 个",
			"sub_accounts":           "子账号列表",
			"uin":                    "帐号 UIN",
			"name":                   "子账号名称",
			"uid":                    "子用户 UID",
			"remark":                 "备注",
			"create_time":            "创建时间",
			"user_type":              "用户类型 (1: 根账号; 2: 子账号; 3: 企微子账号; 4: 协作者; 5: 消息接收者)",
		},
	})
}

func dataSourceTencentCloudCamSubAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudCamSubAccountsRead,
		Schema: map[string]*schema.Schema{
			"filter_sub_account_uin": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of sub-user UINs. Up to 50 UINs are supported.",
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			"sub_accounts": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Sub-user list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"uin": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Sub-user ID.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sub-user name.",
						},
						"uid": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Sub-user UID. UID is the unique identifier of a user who is a message recipient, while UIN is a unique identifier of a user.",
						},
						"remark": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sub-user remarks.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"user_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "User type (1: root account; 2: sub-user; 3: WeCom sub-user; 4: collaborator; 5: message recipient).",
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

func dataSourceTencentCloudCamSubAccountsRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("data_source.tencenttencentcloudenterprise_cam_sub_accounts.read")()

	ctx := context.Background()
	service := CamService{client: meta.(*TencentCloudClient).apiV3Conn}

	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("filter_sub_account_uin"); ok {
		filterSubAccountUinList := []*uint64{}
		filterSubAccountUinSet := v.(*schema.Set).List()
		for i := range filterSubAccountUinSet {
			filterSubAccountUin := filterSubAccountUinSet[i].(int)
			filterSubAccountUinList = append(filterSubAccountUinList, helper.IntUint64(filterSubAccountUin))
		}
		paramMap["FilterSubAccountUin"] = filterSubAccountUinList
	}

	var respData []*cam.SubAccountFilter
	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeCamSubAccountsByFilter(ctx, paramMap)
		if e != nil {
			return retryError(e)
		}
		respData = result
		return nil
	})
	if err != nil {
		return err
	}

	var ids []string
	subAccountsList := make([]map[string]interface{}, 0, len(respData))
	if respData != nil {
		for _, subAccounts := range respData {
			subAccountsMap := map[string]interface{}{}

			var uin uint64
			if subAccounts.Uin != nil {
				subAccountsMap["uin"] = int(*subAccounts.Uin)
				uin = *subAccounts.Uin
			}

			if subAccounts.Name != nil {
				subAccountsMap["name"] = *subAccounts.Name
			}

			if subAccounts.Uid != nil {
				subAccountsMap["uid"] = int(*subAccounts.Uid)
			}

			if subAccounts.Remark != nil {
				subAccountsMap["remark"] = *subAccounts.Remark
			}

			if subAccounts.CreateTime != nil {
				subAccountsMap["create_time"] = *subAccounts.CreateTime
			}

			if subAccounts.UserType != nil {
				subAccountsMap["user_type"] = int(*subAccounts.UserType)
			}

			ids = append(ids, helper.UInt64ToStr(uin))
			subAccountsList = append(subAccountsList, subAccountsMap)
		}

		_ = d.Set("sub_accounts", subAccountsList)
	}

	d.SetId(helper.DataResourceIdsHash(ids))

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), subAccountsList); e != nil {
			return e
		}
	}

	return nil
}

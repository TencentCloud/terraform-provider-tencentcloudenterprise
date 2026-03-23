/*
使用Account SDK接口查询当前账号信息。

Use Account SDK API to query current account information.

# Example Usage

```hcl
data "tencentcloudenterprise_user_info" "current" {
}

	output "app_id" {
	  value = data.tencentcloudenterprise_user_info.current.app_id
	}

	output "uin" {
	  value = data.tencentcloudenterprise_user_info.current.uin
	}

```
*/
package tencentcloud

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	account20181225 "terraform-provider-tencentcloudenterprise/sdk/account/v20181225"
	account "terraform-provider-tencentcloudenterprise/sdk/account/v20190325"
	"terraform-provider-tencentcloudenterprise/tencentcloud/internal/helper"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	registerDataDescriptionProvider("tencentcloudenterprise_user_info", CNDescription{
		TerraformTypeCN: "当前账号信息",
		DescriptionCN:   "用于查询当前账号的基本信息。",
		AttributesCN: map[string]string{
			"app_id":             "当前账号AppID",
			"uin":                "当前账号UIN",
			"owner_uin":          "当前账号OwnerUIN",
			"name":               "当前账号名称（仅子账号支持）",
			"result_output_file": "用于保存结果",
		},
	})
}

func dataSourceTencentCloudUserInfo() *schema.Resource {
	return &schema.Resource{
		Read: datasourceTencentCloudUserInfoRead,
		Schema: map[string]*schema.Schema{
			"app_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current account App ID.",
			},

			"uin": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current account UIN.",
			},

			"owner_uin": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current account OwnerUIN.",
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current account Name. NOTE: only support subaccount.",
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used for save results.",
			},
		},
	}
}

func datasourceTencentCloudUserInfoRead(d *schema.ResourceData, meta interface{}) error {
	defer logElapsed("datasource.tencentcloudenterprise_user_info.read")()

	logId := getLogId(contextNil)
	client := meta.(*TencentCloudClient).apiV3Conn

	var appId, uin, ownerUin, name string

	// 1. Get user info (Uin, OwnerUin, UserName)
	userInfoRequest := account.NewGetUserInfoRequest()
	var userInfoResponse *account.GetUserInfoResponse

	err := resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := client.UseAccountClient().GetUserInfo(userInfoRequest)
		if e != nil {
			return retryError(e)
		}
		userInfoResponse = result
		return nil
	})

	if err != nil {
		log.Printf("[CRITAL]%s api[GetUserInfo] fail, reason[%s]\n", logId, err.Error())
		return err
	}

	if userInfoResponse == nil || userInfoResponse.Response == nil {
		return fmt.Errorf("get user info error: empty response")
	}

	// Extract Uin and OwnerUin
	if userInfoResponse.Response.Uin != nil {
		uin = strconv.FormatUint(*userInfoResponse.Response.Uin, 10)
	}
	if userInfoResponse.Response.OwnerUin != nil {
		ownerUin = strconv.FormatUint(*userInfoResponse.Response.OwnerUin, 10)
	}
	if userInfoResponse.Response.UserName != nil {
		name = *userInfoResponse.Response.UserName
	}

	log.Printf("[DEBUG]%s api[GetUserInfo] success, uin[%s], ownerUin[%s]\n", logId, uin, ownerUin)

	// 2. Get AppId using v20181225 GetAppIdByUin API
	appIdRequest := account20181225.NewGetAppIdByUinRequest()
	// OpUin is optional, if not set, uses current login uin
	if userInfoResponse.Response.Uin != nil {
		appIdRequest.OpUin = userInfoResponse.Response.Uin
	}

	var appIdResponse *account20181225.GetAppIdByUinResponse

	err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
		result, e := client.UseAccount20181225Client().GetAppIdByUin(appIdRequest)
		if e != nil {
			return retryError(e)
		}
		appIdResponse = result
		return nil
	})

	if err != nil {
		log.Printf("[WARN]%s api[GetAppIdByUin] fail, reason[%s]\n", logId, err.Error())
	} else if appIdResponse != nil && appIdResponse.Response != nil && appIdResponse.Response.AppId != nil {
		appId = strconv.FormatUint(*appIdResponse.Response.AppId, 10)
		log.Printf("[DEBUG]%s api[GetAppIdByUin] success, appId[%s]\n", logId, appId)
	}

	// 3. If it's a sub-account, try to get the name from DescribeSubAccounts
	if uin != ownerUin && name == "" {
		accountInfoRequest := account.NewDescribeSubAccountsRequest()
		uinUint64 := helper.StrToUInt64(uin)
		accountInfoRequest.FilterSubAccountUin = []*uint64{&uinUint64}

		err = resource.Retry(readRetryTimeout, func() *resource.RetryError {
			accountInfoResult, e := client.UseAccountClient().DescribeSubAccounts(accountInfoRequest)
			if e != nil {
				return retryError(e)
			}
			if accountInfoResult != nil && accountInfoResult.Response != nil {
				subAccounts := accountInfoResult.Response.SubAccounts
				if len(subAccounts) > 0 && subAccounts[0].Name != nil {
					name = *subAccounts[0].Name
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("[WARN]%s read sub-account name failed, reason:%s\n", logId, err.Error())
			// Don't return error, just continue without the name
		}
	}

	d.SetId(fmt.Sprintf("user-%s-%s-%d", uin, appId, rand.Intn(10000)))

	_ = d.Set("app_id", appId)
	_ = d.Set("uin", uin)
	_ = d.Set("owner_uin", ownerUin)
	_ = d.Set("name", name)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := writeToFile(output.(string), map[string]interface{}{
			"app_id":    appId,
			"uin":       uin,
			"owner_uin": ownerUin,
			"name":      name,
		}); e != nil {
			return e
		}
	}

	return nil
}

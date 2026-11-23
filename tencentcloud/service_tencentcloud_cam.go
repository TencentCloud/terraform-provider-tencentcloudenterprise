package tencentcloud

import (
	"context"
	"log"
	cam "terraform-provider-tencentcloudenterprise/sdk/cam/v20190116"
	"terraform-provider-tencentcloudenterprise/tencentcloud/connectivity"
	"terraform-provider-tencentcloudenterprise/tencentcloud/ratelimit"

	"fmt"
)

type CamService struct {
	client *connectivity.TencentCloudClient
}

func (me *CamService) DescribeUserById(ctx context.Context, name string) (*cam.SubAccountFilter, error) {
	users, err := me.DescribeUsersByFilter(ctx, map[string]interface{}{"name": name})
	if err != nil || len(users) == 0 {
		return nil, err
	}
	return users[0], nil
}

func (me *CamService) QueryApiKey(ctx context.Context, targetUin uint64) (keys []*cam.ApiKey, errRet error) {
	logId := getLogId(ctx)
	request := cam.NewQueryApiKeyRequest()
	request.TargetUin = &targetUin

	ratelimit.Check(request.GetAction())
	response, err := me.client.UseCamClient().QueryApiKey(request)
	if err != nil {
		log.Printf("[ERROR]%s api[%s] failed: %s", logId, request.GetAction(), err.Error())
		errRet = err
		return
	}
	log.Printf("[DEBUG]%s api[%s] success", logId, request.GetAction())

	if response != nil && response.Response != nil {
		keys = response.Response.IdKeys
	}
	return
}

func (me *CamService) DisableApiKey(ctx context.Context, apiUin uint64, secretId string) (errRet error) {
	logId := getLogId(ctx)
	request := cam.NewDisableApiKeyRequest()
	request.ApiUin = &apiUin
	request.ApiSecretId = &secretId

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseCamClient().DisableApiKey(request)
	if err != nil {
		log.Printf("[ERROR]%s api[%s] failed: %s", logId, request.GetAction(), err.Error())
		errRet = err
		return
	}
	return
}

func (me *CamService) DeleteApiKey(ctx context.Context, apiUin uint64, secretId string) (errRet error) {
	logId := getLogId(ctx)
	request := cam.NewDeleteApiKeyRequest()
	request.ApiUin = &apiUin
	request.ApiSecretId = &secretId

	ratelimit.Check(request.GetAction())
	_, err := me.client.UseCamClient().DeleteApiKey(request)
	if err != nil {
		log.Printf("[ERROR]%s api[%s] failed: %s", logId, request.GetAction(), err.Error())
		errRet = err
		return
	}
	return
}

func (me *CamService) DescribeUsersByFilter(ctx context.Context, params map[string]interface{}) ([]*cam.SubAccountFilter, error) {
	request := cam.NewListSubAccountsRequest()
	var offset uint64
	limit := uint64(50)
	request.Offset, request.Limit = &offset, &limit
	if name, ok := params["name"].(string); ok && name != "" {
		request.Keyword = &name
	}
	var users []*cam.SubAccountFilter
	for {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		ratelimit.Check(request.GetAction())
		response, err := me.client.UseCamClient().ListSubAccounts(request)
		if err != nil {
			return nil, err
		}
		if response == nil || response.Response == nil {
			return nil, fmt.Errorf("ListSubAccounts returned no response")
		}
		page := response.Response.UserInfo
		// Filter results based on params
		for _, user := range page {
			if user == nil {
				continue
			}
			// Filter by name
			if params["name"] != nil {
				if user.Name == nil || *user.Name != params["name"].(string) {
					continue
				}
			}

			// Filter by remark
			if params["remark"] != nil {
				if user.Remark == nil || *user.Remark != params["remark"].(string) {
					continue
				}
			}

			// Filter by phone_num
			if params["phone_num"] != nil {
				if user.PhoneNum == nil || *user.PhoneNum != params["phone_num"].(string) {
					continue
				}
			}

			// Filter by country_code
			if params["country_code"] != nil {
				if user.CountryCode == nil || *user.CountryCode != params["country_code"].(string) {
					continue
				}
			}

			// Filter by email
			if params["email"] != nil {
				if user.Email == nil || *user.Email != params["email"].(string) {
					continue
				}
			}

			// Filter by console_login
			if params["console_login"] != nil {
				if user.ConsoleLogin == nil || int(*user.ConsoleLogin) != params["console_login"].(int) {
					continue
				}
			}

			// Filter by uin
			if params["uin"] != nil {
				if user.Uin == nil || int(*user.Uin) != params["uin"].(int) {
					continue
				}
			}

			// Filter by uid
			if params["uid"] != nil {
				if user.Uid == nil || int(*user.Uid) != params["uid"].(int) {
					continue
				}
			}

			users = append(users, user)
		}
		offset += uint64(len(page))
		if len(page) == 0 || (response.Response.TotalNum != nil && offset >= *response.Response.TotalNum) || (response.Response.TotalNum == nil && uint64(len(page)) < limit) {
			return users, nil
		}
	}
}

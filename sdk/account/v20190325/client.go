// All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190325

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-03-25"

var _ = tchttp.POST

type Client struct {
	common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	cpf := profile.NewClientProfile()
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewCheckVerifyCodeRequest() (request *CheckVerifyCodeRequest) {
	request = &CheckVerifyCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CheckVerifyCode")
	return
}

func NewCheckVerifyCodeResponse() (response *CheckVerifyCodeResponse) {
	response = &CheckVerifyCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证码校验
func (c *Client) CheckVerifyCode(request *CheckVerifyCodeRequest) (response *CheckVerifyCodeResponse, err error) {
	if request == nil {
		request = NewCheckVerifyCodeRequest()
	}
	response = NewCheckVerifyCodeResponse()
	err = c.Send(request, response)
	return
}

func NewListRetentionSubAccountsRequest() (request *ListRetentionSubAccountsRequest) {
	request = &ListRetentionSubAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ListRetentionSubAccounts")
	return
}

func NewListRetentionSubAccountsResponse() (response *ListRetentionSubAccountsResponse) {
	response = &ListRetentionSubAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看回收站中的账号
func (c *Client) ListRetentionSubAccounts(request *ListRetentionSubAccountsRequest) (response *ListRetentionSubAccountsResponse, err error) {
	if request == nil {
		request = NewListRetentionSubAccountsRequest()
	}
	response = NewListRetentionSubAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewModifySensitiveActionRequest() (request *ModifySensitiveActionRequest) {
	request = &ModifySensitiveActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ModifySensitiveAction")
	return
}

func NewModifySensitiveActionResponse() (response *ModifySensitiveActionResponse) {
	response = &ModifySensitiveActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改敏感接口
func (c *Client) ModifySensitiveAction(request *ModifySensitiveActionRequest) (response *ModifySensitiveActionResponse, err error) {
	if request == nil {
		request = NewModifySensitiveActionRequest()
	}
	response = NewModifySensitiveActionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSelfApiKeyRequest() (request *CreateSelfApiKeyRequest) {
	request = &CreateSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CreateSelfApiKey")
	return
}

func NewCreateSelfApiKeyResponse() (response *CreateSelfApiKeyResponse) {
	response = &CreateSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建自己的密钥
func (c *Client) CreateSelfApiKey(request *CreateSelfApiKeyRequest) (response *CreateSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewCreateSelfApiKeyRequest()
	}
	response = NewCreateSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetLoginInfoRequest() (request *GetLoginInfoRequest) {
	request = &GetLoginInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetLoginInfo")
	return
}

func NewGetLoginInfoResponse() (response *GetLoginInfoResponse) {
	response = &GetLoginInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取登录历史信息
func (c *Client) GetLoginInfo(request *GetLoginInfoRequest) (response *GetLoginInfoResponse, err error) {
	if request == nil {
		request = NewGetLoginInfoRequest()
	}
	response = NewGetLoginInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetMasterListRequest() (request *GetMasterListRequest) {
	request = &GetMasterListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetMasterList")
	return
}

func NewGetMasterListResponse() (response *GetMasterListResponse) {
	response = &GetMasterListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账户所关联的主账户列表
func (c *Client) GetMasterList(request *GetMasterListRequest) (response *GetMasterListResponse, err error) {
	if request == nil {
		request = NewGetMasterListRequest()
	}
	response = NewGetMasterListResponse()
	err = c.Send(request, response)
	return
}

func NewGetTokenRequest() (request *GetTokenRequest) {
	request = &GetTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetToken")
	return
}

func NewGetTokenResponse() (response *GetTokenResponse) {
	response = &GetTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建并获取一个token
func (c *Client) GetToken(request *GetTokenRequest) (response *GetTokenResponse, err error) {
	if request == nil {
		request = NewGetTokenRequest()
	}
	response = NewGetTokenResponse()
	err = c.Send(request, response)
	return
}

func NewEnableSelfApiKeyRequest() (request *EnableSelfApiKeyRequest) {
	request = &EnableSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "EnableSelfApiKey")
	return
}

func NewEnableSelfApiKeyResponse() (response *EnableSelfApiKeyResponse) {
	response = &EnableSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用自己的密钥
func (c *Client) EnableSelfApiKey(request *EnableSelfApiKeyRequest) (response *EnableSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewEnableSelfApiKeyRequest()
	}
	response = NewEnableSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetOwnerUinByAppidRequest() (request *GetOwnerUinByAppidRequest) {
	request = &GetOwnerUinByAppidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetOwnerUinByAppid")
	return
}

func NewGetOwnerUinByAppidResponse() (response *GetOwnerUinByAppidResponse) {
	response = &GetOwnerUinByAppidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据appid获取主账户uin
func (c *Client) GetOwnerUinByAppid(request *GetOwnerUinByAppidRequest) (response *GetOwnerUinByAppidResponse, err error) {
	if request == nil {
		request = NewGetOwnerUinByAppidRequest()
	}
	response = NewGetOwnerUinByAppidResponse()
	err = c.Send(request, response)
	return
}

func NewGetInfoByFieldsRequest() (request *GetInfoByFieldsRequest) {
	request = &GetInfoByFieldsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetInfoByFields")
	return
}

func NewGetInfoByFieldsResponse() (response *GetInfoByFieldsResponse) {
	response = &GetInfoByFieldsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据账户查询指定属性信息
func (c *Client) GetInfoByFields(request *GetInfoByFieldsRequest) (response *GetInfoByFieldsResponse, err error) {
	if request == nil {
		request = NewGetInfoByFieldsRequest()
	}
	response = NewGetInfoByFieldsResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserAreaByLoginUinRequest() (request *GetUserAreaByLoginUinRequest) {
	request = &GetUserAreaByLoginUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetUserAreaByLoginUin")
	return
}

func NewGetUserAreaByLoginUinResponse() (response *GetUserAreaByLoginUinResponse) {
	response = &GetUserAreaByLoginUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户注册地区域信息
func (c *Client) GetUserAreaByLoginUin(request *GetUserAreaByLoginUinRequest) (response *GetUserAreaByLoginUinResponse, err error) {
	if request == nil {
		request = NewGetUserAreaByLoginUinRequest()
	}
	response = NewGetUserAreaByLoginUinResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubAccountsRequest() (request *DescribeSubAccountsRequest) {
	request = &DescribeSubAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DescribeSubAccounts")
	return
}

func NewDescribeSubAccountsResponse() (response *DescribeSubAccountsResponse) {
	response = &DescribeSubAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主账号下的子账号信息
func (c *Client) DescribeSubAccounts(request *DescribeSubAccountsRequest) (response *DescribeSubAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeSubAccountsRequest()
	}
	response = NewDescribeSubAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserInfoByLoginUinRequest() (request *GetUserInfoByLoginUinRequest) {
	request = &GetUserInfoByLoginUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetUserInfoByLoginUin")
	return
}

func NewGetUserInfoByLoginUinResponse() (response *GetUserInfoByLoginUinResponse) {
	response = &GetUserInfoByLoginUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin查询用户详细信息
func (c *Client) GetUserInfoByLoginUin(request *GetUserInfoByLoginUinRequest) (response *GetUserInfoByLoginUinResponse, err error) {
	if request == nil {
		request = NewGetUserInfoByLoginUinRequest()
	}
	response = NewGetUserInfoByLoginUinResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRelatedUinSessionKeyRequest() (request *DescribeRelatedUinSessionKeyRequest) {
	request = &DescribeRelatedUinSessionKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DescribeRelatedUinSessionKey")
	return
}

func NewDescribeRelatedUinSessionKeyResponse() (response *DescribeRelatedUinSessionKeyResponse) {
	response = &DescribeRelatedUinSessionKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取关联账号的登录态
func (c *Client) DescribeRelatedUinSessionKey(request *DescribeRelatedUinSessionKeyRequest) (response *DescribeRelatedUinSessionKeyResponse, err error) {
	if request == nil {
		request = NewDescribeRelatedUinSessionKeyRequest()
	}
	response = NewDescribeRelatedUinSessionKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppIdByOwnerUinRequest() (request *GetAppIdByOwnerUinRequest) {
	request = &GetAppIdByOwnerUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAppIdByOwnerUin")
	return
}

func NewGetAppIdByOwnerUinResponse() (response *GetAppIdByOwnerUinResponse) {
	response = &GetAppIdByOwnerUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据主账户id获取appid
func (c *Client) GetAppIdByOwnerUin(request *GetAppIdByOwnerUinRequest) (response *GetAppIdByOwnerUinResponse, err error) {
	if request == nil {
		request = NewGetAppIdByOwnerUinRequest()
	}
	response = NewGetAppIdByOwnerUinResponse()
	err = c.Send(request, response)
	return
}

func NewSetSafeAuthFlagRequest() (request *SetSafeAuthFlagRequest) {
	request = &SetSafeAuthFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SetSafeAuthFlag")
	return
}

func NewSetSafeAuthFlagResponse() (response *SetSafeAuthFlagResponse) {
	response = &SetSafeAuthFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置安全认证配置
func (c *Client) SetSafeAuthFlag(request *SetSafeAuthFlagRequest) (response *SetSafeAuthFlagResponse, err error) {
	if request == nil {
		request = NewSetSafeAuthFlagRequest()
	}
	response = NewSetSafeAuthFlagResponse()
	err = c.Send(request, response)
	return
}

func NewTokenUnBindRequest() (request *TokenUnBindRequest) {
	request = &TokenUnBindRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "TokenUnBind")
	return
}

func NewTokenUnBindResponse() (response *TokenUnBindResponse) {
	response = &TokenUnBindResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解绑账号mfa认证token设备
func (c *Client) TokenUnBind(request *TokenUnBindRequest) (response *TokenUnBindResponse, err error) {
	if request == nil {
		request = NewTokenUnBindRequest()
	}
	response = NewTokenUnBindResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSubAccountNameRequest() (request *CheckSubAccountNameRequest) {
	request = &CheckSubAccountNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CheckSubAccountName")
	return
}

func NewCheckSubAccountNameResponse() (response *CheckSubAccountNameResponse) {
	response = &CheckSubAccountNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验子账户名称是和否存在
func (c *Client) CheckSubAccountName(request *CheckSubAccountNameRequest) (response *CheckSubAccountNameResponse, err error) {
	if request == nil {
		request = NewCheckSubAccountNameRequest()
	}
	response = NewCheckSubAccountNameResponse()
	err = c.Send(request, response)
	return
}

func NewGetSelfApiKeyRequest() (request *GetSelfApiKeyRequest) {
	request = &GetSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetSelfApiKey")
	return
}

func NewGetSelfApiKeyResponse() (response *GetSelfApiKeyResponse) {
	response = &GetSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取自己的密钥
func (c *Client) GetSelfApiKey(request *GetSelfApiKeyRequest) (response *GetSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewGetSelfApiKeyRequest()
	}
	response = NewGetSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserByAttributeValueRequest() (request *GetUserByAttributeValueRequest) {
	request = &GetUserByAttributeValueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetUserByAttributeValue")
	return
}

func NewGetUserByAttributeValueResponse() (response *GetUserByAttributeValueResponse) {
	response = &GetUserByAttributeValueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据扩展属性查询用户
func (c *Client) GetUserByAttributeValue(request *GetUserByAttributeValueRequest) (response *GetUserByAttributeValueResponse, err error) {
	if request == nil {
		request = NewGetUserByAttributeValueRequest()
	}
	response = NewGetUserByAttributeValueResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAccountExistRequest() (request *CheckAccountExistRequest) {
	request = &CheckAccountExistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CheckAccountExist")
	return
}

func NewCheckAccountExistResponse() (response *CheckAccountExistResponse) {
	response = &CheckAccountExistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验账户是否存在
func (c *Client) CheckAccountExist(request *CheckAccountExistRequest) (response *CheckAccountExistResponse, err error) {
	if request == nil {
		request = NewCheckAccountExistRequest()
	}
	response = NewCheckAccountExistResponse()
	err = c.Send(request, response)
	return
}

func NewCheckCaptchaRequest() (request *CheckCaptchaRequest) {
	request = &CheckCaptchaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CheckCaptcha")
	return
}

func NewCheckCaptchaResponse() (response *CheckCaptchaResponse) {
	response = &CheckCaptchaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验图形验证码
func (c *Client) CheckCaptcha(request *CheckCaptchaRequest) (response *CheckCaptchaResponse, err error) {
	if request == nil {
		request = NewCheckCaptchaRequest()
	}
	response = NewCheckCaptchaResponse()
	err = c.Send(request, response)
	return
}

func NewGetFrequentLoginConfigRequest() (request *GetFrequentLoginConfigRequest) {
	request = &GetFrequentLoginConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetFrequentLoginConfig")
	return
}

func NewGetFrequentLoginConfigResponse() (response *GetFrequentLoginConfigResponse) {
	response = &GetFrequentLoginConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 常用登录信息配置
func (c *Client) GetFrequentLoginConfig(request *GetFrequentLoginConfigRequest) (response *GetFrequentLoginConfigResponse, err error) {
	if request == nil {
		request = NewGetFrequentLoginConfigRequest()
	}
	response = NewGetFrequentLoginConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRetentionPolicyRequest() (request *ModifyRetentionPolicyRequest) {
	request = &ModifyRetentionPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ModifyRetentionPolicy")
	return
}

func NewModifyRetentionPolicyResponse() (response *ModifyRetentionPolicyResponse) {
	response = &ModifyRetentionPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改账号删除保留策略
func (c *Client) ModifyRetentionPolicy(request *ModifyRetentionPolicyRequest) (response *ModifyRetentionPolicyResponse, err error) {
	if request == nil {
		request = NewModifyRetentionPolicyRequest()
	}
	response = NewModifyRetentionPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetAutoLockPeriodRequest() (request *GetAutoLockPeriodRequest) {
	request = &GetAutoLockPeriodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAutoLockPeriod")
	return
}

func NewGetAutoLockPeriodResponse() (response *GetAutoLockPeriodResponse) {
	response = &GetAutoLockPeriodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取自动锁定时间
func (c *Client) GetAutoLockPeriod(request *GetAutoLockPeriodRequest) (response *GetAutoLockPeriodResponse, err error) {
	if request == nil {
		request = NewGetAutoLockPeriodRequest()
	}
	response = NewGetAutoLockPeriodResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubAccountInfoRequest() (request *GetSubAccountInfoRequest) {
	request = &GetSubAccountInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetSubAccountInfo")
	return
}

func NewGetSubAccountInfoResponse() (response *GetSubAccountInfoResponse) {
	response = &GetSubAccountInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据用户名查询子账户信息
func (c *Client) GetSubAccountInfo(request *GetSubAccountInfoRequest) (response *GetSubAccountInfoResponse, err error) {
	if request == nil {
		request = NewGetSubAccountInfoRequest()
	}
	response = NewGetSubAccountInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTokenRequest() (request *DeleteTokenRequest) {
	request = &DeleteTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DeleteToken")
	return
}

func NewDeleteTokenResponse() (response *DeleteTokenResponse) {
	response = &DeleteTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 将token置为无效
func (c *Client) DeleteToken(request *DeleteTokenRequest) (response *DeleteTokenResponse, err error) {
	if request == nil {
		request = NewDeleteTokenRequest()
	}
	response = NewDeleteTokenResponse()
	err = c.Send(request, response)
	return
}

func NewGetNicknameRequest() (request *GetNicknameRequest) {
	request = &GetNicknameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetNickname")
	return
}

func NewGetNicknameResponse() (response *GetNicknameResponse) {
	response = &GetNicknameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户昵称
func (c *Client) GetNickname(request *GetNicknameRequest) (response *GetNicknameResponse, err error) {
	if request == nil {
		request = NewGetNicknameRequest()
	}
	response = NewGetNicknameResponse()
	err = c.Send(request, response)
	return
}

func NewQuerySelfApiKeyRequest() (request *QuerySelfApiKeyRequest) {
	request = &QuerySelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "QuerySelfApiKey")
	return
}

func NewQuerySelfApiKeyResponse() (response *QuerySelfApiKeyResponse) {
	response = &QuerySelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取自己的API密钥列表
func (c *Client) QuerySelfApiKey(request *QuerySelfApiKeyRequest) (response *QuerySelfApiKeyResponse, err error) {
	if request == nil {
		request = NewQuerySelfApiKeyRequest()
	}
	response = NewQuerySelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewSetSingleLoginFlagRequest() (request *SetSingleLoginFlagRequest) {
	request = &SetSingleLoginFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SetSingleLoginFlag")
	return
}

func NewSetSingleLoginFlagResponse() (response *SetSingleLoginFlagResponse) {
	response = &SetSingleLoginFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置单点登陆标识
func (c *Client) SetSingleLoginFlag(request *SetSingleLoginFlagRequest) (response *SetSingleLoginFlagResponse, err error) {
	if request == nil {
		request = NewSetSingleLoginFlagRequest()
	}
	response = NewSetSingleLoginFlagResponse()
	err = c.Send(request, response)
	return
}

func NewGetCustomSessionSettingRequest() (request *GetCustomSessionSettingRequest) {
	request = &GetCustomSessionSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetCustomSessionSetting")
	return
}

func NewGetCustomSessionSettingResponse() (response *GetCustomSessionSettingResponse) {
	response = &GetCustomSessionSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户端会话时间设置
func (c *Client) GetCustomSessionSetting(request *GetCustomSessionSettingRequest) (response *GetCustomSessionSettingResponse, err error) {
	if request == nil {
		request = NewGetCustomSessionSettingRequest()
	}
	response = NewGetCustomSessionSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserCategoryRequest() (request *DescribeUserCategoryRequest) {
	request = &DescribeUserCategoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DescribeUserCategory")
	return
}

func NewDescribeUserCategoryResponse() (response *DescribeUserCategoryResponse) {
	response = &DescribeUserCategoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeUserCategory
func (c *Client) DescribeUserCategory(request *DescribeUserCategoryRequest) (response *DescribeUserCategoryResponse, err error) {
	if request == nil {
		request = NewDescribeUserCategoryRequest()
	}
	response = NewDescribeUserCategoryResponse()
	err = c.Send(request, response)
	return
}

func NewGetCountryCodeRequest() (request *GetCountryCodeRequest) {
	request = &GetCountryCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetCountryCode")
	return
}

func NewGetCountryCodeResponse() (response *GetCountryCodeResponse) {
	response = &GetCountryCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全部国家编码列表
func (c *Client) GetCountryCode(request *GetCountryCodeRequest) (response *GetCountryCodeResponse, err error) {
	if request == nil {
		request = NewGetCountryCodeRequest()
	}
	response = NewGetCountryCodeResponse()
	err = c.Send(request, response)
	return
}

func NewSetLoginFlagRequest() (request *SetLoginFlagRequest) {
	request = &SetLoginFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SetLoginFlag")
	return
}

func NewSetLoginFlagResponse() (response *SetLoginFlagResponse) {
	response = &SetLoginFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置登陆标识
func (c *Client) SetLoginFlag(request *SetLoginFlagRequest) (response *SetLoginFlagResponse, err error) {
	if request == nil {
		request = NewSetLoginFlagRequest()
	}
	response = NewSetLoginFlagResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSoftTokenRequest() (request *CreateSoftTokenRequest) {
	request = &CreateSoftTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CreateSoftToken")
	return
}

func NewCreateSoftTokenResponse() (response *CreateSoftTokenResponse) {
	response = &CreateSoftTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成虚拟设备token
func (c *Client) CreateSoftToken(request *CreateSoftTokenRequest) (response *CreateSoftTokenResponse, err error) {
	if request == nil {
		request = NewCreateSoftTokenRequest()
	}
	response = NewCreateSoftTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSelfApiKeyRequest() (request *DeleteSelfApiKeyRequest) {
	request = &DeleteSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DeleteSelfApiKey")
	return
}

func NewDeleteSelfApiKeyResponse() (response *DeleteSelfApiKeyResponse) {
	response = &DeleteSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除自己的密钥
func (c *Client) DeleteSelfApiKey(request *DeleteSelfApiKeyRequest) (response *DeleteSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewDeleteSelfApiKeyRequest()
	}
	response = NewDeleteSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetUinOwnerInOpenRequest() (request *GetUinOwnerInOpenRequest) {
	request = &GetUinOwnerInOpenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetUinOwnerInOpen")
	return
}

func NewGetUinOwnerInOpenResponse() (response *GetUinOwnerInOpenResponse) {
	response = &GetUinOwnerInOpenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取开放平台主账号
func (c *Client) GetUinOwnerInOpen(request *GetUinOwnerInOpenRequest) (response *GetUinOwnerInOpenResponse, err error) {
	if request == nil {
		request = NewGetUinOwnerInOpenRequest()
	}
	response = NewGetUinOwnerInOpenResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaxSubAccountAndGroupNumRequest() (request *GetMaxSubAccountAndGroupNumRequest) {
	request = &GetMaxSubAccountAndGroupNumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetMaxSubAccountAndGroupNum")
	return
}

func NewGetMaxSubAccountAndGroupNumResponse() (response *GetMaxSubAccountAndGroupNumResponse) {
	response = &GetMaxSubAccountAndGroupNumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetMaxSubAccountAndGroupNum
func (c *Client) GetMaxSubAccountAndGroupNum(request *GetMaxSubAccountAndGroupNumRequest) (response *GetMaxSubAccountAndGroupNumResponse, err error) {
	if request == nil {
		request = NewGetMaxSubAccountAndGroupNumRequest()
	}
	response = NewGetMaxSubAccountAndGroupNumResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubLoginUinListRequest() (request *GetSubLoginUinListRequest) {
	request = &GetSubLoginUinListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetSubLoginUinList")
	return
}

func NewGetSubLoginUinListResponse() (response *GetSubLoginUinListResponse) {
	response = &GetSubLoginUinListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主账户下的子账户列表
func (c *Client) GetSubLoginUinList(request *GetSubLoginUinListRequest) (response *GetSubLoginUinListResponse, err error) {
	if request == nil {
		request = NewGetSubLoginUinListRequest()
	}
	response = NewGetSubLoginUinListResponse()
	err = c.Send(request, response)
	return
}

func NewMFAStatusRequest() (request *MFAStatusRequest) {
	request = &MFAStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "MFAStatus")
	return
}

func NewMFAStatusResponse() (response *MFAStatusResponse) {
	response = &MFAStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 多因子认证状态信息
func (c *Client) MFAStatus(request *MFAStatusRequest) (response *MFAStatusResponse, err error) {
	if request == nil {
		request = NewMFAStatusRequest()
	}
	response = NewMFAStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserInfoRequest() (request *GetUserInfoRequest) {
	request = &GetUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetUserInfo")
	return
}

func NewGetUserInfoResponse() (response *GetUserInfoResponse) {
	response = &GetUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户信息
func (c *Client) GetUserInfo(request *GetUserInfoRequest) (response *GetUserInfoResponse, err error) {
	if request == nil {
		request = NewGetUserInfoRequest()
	}
	response = NewGetUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAutoLockPeriodRequest() (request *ModifyAutoLockPeriodRequest) {
	request = &ModifyAutoLockPeriodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ModifyAutoLockPeriod")
	return
}

func NewModifyAutoLockPeriodResponse() (response *ModifyAutoLockPeriodResponse) {
	response = &ModifyAutoLockPeriodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自动锁定时间
func (c *Client) ModifyAutoLockPeriod(request *ModifyAutoLockPeriodRequest) (response *ModifyAutoLockPeriodResponse, err error) {
	if request == nil {
		request = NewModifyAutoLockPeriodRequest()
	}
	response = NewModifyAutoLockPeriodResponse()
	err = c.Send(request, response)
	return
}

func NewCheckSubAccountUinRequest() (request *CheckSubAccountUinRequest) {
	request = &CheckSubAccountUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CheckSubAccountUin")
	return
}

func NewCheckSubAccountUinResponse() (response *CheckSubAccountUinResponse) {
	response = &CheckSubAccountUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验子账户uin
func (c *Client) CheckSubAccountUin(request *CheckSubAccountUinRequest) (response *CheckSubAccountUinResponse, err error) {
	if request == nil {
		request = NewCheckSubAccountUinRequest()
	}
	response = NewCheckSubAccountUinResponse()
	err = c.Send(request, response)
	return
}

func NewCheckTokenRequest() (request *CheckTokenRequest) {
	request = &CheckTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CheckToken")
	return
}

func NewCheckTokenResponse() (response *CheckTokenResponse) {
	response = &CheckTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验token是否有效
func (c *Client) CheckToken(request *CheckTokenRequest) (response *CheckTokenResponse, err error) {
	if request == nil {
		request = NewCheckTokenRequest()
	}
	response = NewCheckTokenResponse()
	err = c.Send(request, response)
	return
}

func NewLoginVerifyRequest() (request *LoginVerifyRequest) {
	request = &LoginVerifyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "LoginVerify")
	return
}

func NewLoginVerifyResponse() (response *LoginVerifyResponse) {
	response = &LoginVerifyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 登陆账户校验
func (c *Client) LoginVerify(request *LoginVerifyRequest) (response *LoginVerifyResponse, err error) {
	if request == nil {
		request = NewLoginVerifyRequest()
	}
	response = NewLoginVerifyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNicknameRequest() (request *ModifyNicknameRequest) {
	request = &ModifyNicknameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ModifyNickname")
	return
}

func NewModifyNicknameResponse() (response *ModifyNicknameResponse) {
	response = &ModifyNicknameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户昵称
func (c *Client) ModifyNickname(request *ModifyNicknameRequest) (response *ModifyNicknameResponse, err error) {
	if request == nil {
		request = NewModifyNicknameRequest()
	}
	response = NewModifyNicknameResponse()
	err = c.Send(request, response)
	return
}

func NewSeedLoginTokenRequest() (request *SeedLoginTokenRequest) {
	request = &SeedLoginTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SeedLoginToken")
	return
}

func NewSeedLoginTokenResponse() (response *SeedLoginTokenResponse) {
	response = &SeedLoginTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成登陆用的token
func (c *Client) SeedLoginToken(request *SeedLoginTokenRequest) (response *SeedLoginTokenResponse, err error) {
	if request == nil {
		request = NewSeedLoginTokenRequest()
	}
	response = NewSeedLoginTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAttributeValuesRequest() (request *DeleteAttributeValuesRequest) {
	request = &DeleteAttributeValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DeleteAttributeValues")
	return
}

func NewDeleteAttributeValuesResponse() (response *DeleteAttributeValuesResponse) {
	response = &DeleteAttributeValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除扩展属性值
func (c *Client) DeleteAttributeValues(request *DeleteAttributeValuesRequest) (response *DeleteAttributeValuesResponse, err error) {
	if request == nil {
		request = NewDeleteAttributeValuesRequest()
	}
	response = NewDeleteAttributeValuesResponse()
	err = c.Send(request, response)
	return
}

func NewSendVerifyCodeRequest() (request *SendVerifyCodeRequest) {
	request = &SendVerifyCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SendVerifyCode")
	return
}

func NewSendVerifyCodeResponse() (response *SendVerifyCodeResponse) {
	response = &SendVerifyCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送验证码
func (c *Client) SendVerifyCode(request *SendVerifyCodeRequest) (response *SendVerifyCodeResponse, err error) {
	if request == nil {
		request = NewSendVerifyCodeRequest()
	}
	response = NewSendVerifyCodeResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppIdRequest() (request *GetAppIdRequest) {
	request = &GetAppIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAppId")
	return
}

func NewGetAppIdResponse() (response *GetAppIdResponse) {
	response = &GetAppIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取appid
func (c *Client) GetAppId(request *GetAppIdRequest) (response *GetAppIdResponse, err error) {
	if request == nil {
		request = NewGetAppIdRequest()
	}
	response = NewGetAppIdResponse()
	err = c.Send(request, response)
	return
}

func NewGetAppIdByLoginUinRequest() (request *GetAppIdByLoginUinRequest) {
	request = &GetAppIdByLoginUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAppIdByLoginUin")
	return
}

func NewGetAppIdByLoginUinResponse() (response *GetAppIdByLoginUinResponse) {
	response = &GetAppIdByLoginUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 使用uin获取账户所对应的appid
func (c *Client) GetAppIdByLoginUin(request *GetAppIdByLoginUinRequest) (response *GetAppIdByLoginUinResponse, err error) {
	if request == nil {
		request = NewGetAppIdByLoginUinRequest()
	}
	response = NewGetAppIdByLoginUinResponse()
	err = c.Send(request, response)
	return
}

func NewSetCaptchaRequest() (request *SetCaptchaRequest) {
	request = &SetCaptchaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SetCaptcha")
	return
}

func NewSetCaptchaResponse() (response *SetCaptchaResponse) {
	response = &SetCaptchaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置图形验证码
func (c *Client) SetCaptcha(request *SetCaptchaRequest) (response *SetCaptchaResponse, err error) {
	if request == nil {
		request = NewSetCaptchaRequest()
	}
	response = NewSetCaptchaResponse()
	err = c.Send(request, response)
	return
}

func NewGetRetentionPolicyRequest() (request *GetRetentionPolicyRequest) {
	request = &GetRetentionPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetRetentionPolicy")
	return
}

func NewGetRetentionPolicyResponse() (response *GetRetentionPolicyResponse) {
	response = &GetRetentionPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账号删除保留策略
func (c *Client) GetRetentionPolicy(request *GetRetentionPolicyRequest) (response *GetRetentionPolicyResponse, err error) {
	if request == nil {
		request = NewGetRetentionPolicyRequest()
	}
	response = NewGetRetentionPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifySelfApiKeyRequest() (request *ModifySelfApiKeyRequest) {
	request = &ModifySelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ModifySelfApiKey")
	return
}

func NewModifySelfApiKeyResponse() (response *ModifySelfApiKeyResponse) {
	response = &ModifySelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改自己的密钥
func (c *Client) ModifySelfApiKey(request *ModifySelfApiKeyRequest) (response *ModifySelfApiKeyResponse, err error) {
	if request == nil {
		request = NewModifySelfApiKeyRequest()
	}
	response = NewModifySelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewAddSubAccountRequest() (request *AddSubAccountRequest) {
	request = &AddSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "AddSubAccount")
	return
}

func NewAddSubAccountResponse() (response *AddSubAccountResponse) {
	response = &AddSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加子账号
func (c *Client) AddSubAccount(request *AddSubAccountRequest) (response *AddSubAccountResponse, err error) {
	if request == nil {
		request = NewAddSubAccountRequest()
	}
	response = NewAddSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewChangeMailPasswordRequest() (request *ChangeMailPasswordRequest) {
	request = &ChangeMailPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ChangeMailPassword")
	return
}

func NewChangeMailPasswordResponse() (response *ChangeMailPasswordResponse) {
	response = &ChangeMailPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改邮箱账号密码
func (c *Client) ChangeMailPassword(request *ChangeMailPasswordRequest) (response *ChangeMailPasswordResponse, err error) {
	if request == nil {
		request = NewChangeMailPasswordRequest()
	}
	response = NewChangeMailPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewGetSafeAuthConfigRequest() (request *GetSafeAuthConfigRequest) {
	request = &GetSafeAuthConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetSafeAuthConfig")
	return
}

func NewGetSafeAuthConfigResponse() (response *GetSafeAuthConfigResponse) {
	response = &GetSafeAuthConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 读取账户关联的安全认证相关配置
func (c *Client) GetSafeAuthConfig(request *GetSafeAuthConfigRequest) (response *GetSafeAuthConfigResponse, err error) {
	if request == nil {
		request = NewGetSafeAuthConfigRequest()
	}
	response = NewGetSafeAuthConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttributeValuesRequest() (request *GetAttributeValuesRequest) {
	request = &GetAttributeValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAttributeValues")
	return
}

func NewGetAttributeValuesResponse() (response *GetAttributeValuesResponse) {
	response = &GetAttributeValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取扩展字段（企业员信息）
func (c *Client) GetAttributeValues(request *GetAttributeValuesRequest) (response *GetAttributeValuesResponse, err error) {
	if request == nil {
		request = NewGetAttributeValuesRequest()
	}
	response = NewGetAttributeValuesResponse()
	err = c.Send(request, response)
	return
}

func NewChangeSubAccountPasswordRequest() (request *ChangeSubAccountPasswordRequest) {
	request = &ChangeSubAccountPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ChangeSubAccountPassword")
	return
}

func NewChangeSubAccountPasswordResponse() (response *ChangeSubAccountPasswordResponse) {
	response = &ChangeSubAccountPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改子账号密码
func (c *Client) ChangeSubAccountPassword(request *ChangeSubAccountPasswordRequest) (response *ChangeSubAccountPasswordResponse, err error) {
	if request == nil {
		request = NewChangeSubAccountPasswordRequest()
	}
	response = NewChangeSubAccountPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewGetMultiFactorParasRequest() (request *GetMultiFactorParasRequest) {
	request = &GetMultiFactorParasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetMultiFactorParas")
	return
}

func NewGetMultiFactorParasResponse() (response *GetMultiFactorParasResponse) {
	response = &GetMultiFactorParasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取多因子登录参数
func (c *Client) GetMultiFactorParas(request *GetMultiFactorParasRequest) (response *GetMultiFactorParasResponse, err error) {
	if request == nil {
		request = NewGetMultiFactorParasRequest()
	}
	response = NewGetMultiFactorParasResponse()
	err = c.Send(request, response)
	return
}

func NewSetMfaDeviceRequest() (request *SetMfaDeviceRequest) {
	request = &SetMfaDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SetMfaDevice")
	return
}

func NewSetMfaDeviceResponse() (response *SetMfaDeviceResponse) {
	response = &SetMfaDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对子用户分配mfa设备类型，1-三方设备，2-虚拟mfa设备
func (c *Client) SetMfaDevice(request *SetMfaDeviceRequest) (response *SetMfaDeviceResponse, err error) {
	if request == nil {
		request = NewSetMfaDeviceRequest()
	}
	response = NewSetMfaDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewSendLoginVerifyCodeRequest() (request *SendLoginVerifyCodeRequest) {
	request = &SendLoginVerifyCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SendLoginVerifyCode")
	return
}

func NewSendLoginVerifyCodeResponse() (response *SendLoginVerifyCodeResponse) {
	response = &SendLoginVerifyCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送登陆验证码
func (c *Client) SendLoginVerifyCode(request *SendLoginVerifyCodeRequest) (response *SendLoginVerifyCodeResponse, err error) {
	if request == nil {
		request = NewSendLoginVerifyCodeRequest()
	}
	response = NewSendLoginVerifyCodeResponse()
	err = c.Send(request, response)
	return
}

func NewGetAttributeNameRequest() (request *GetAttributeNameRequest) {
	request = &GetAttributeNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAttributeName")
	return
}

func NewGetAttributeNameResponse() (response *GetAttributeNameResponse) {
	response = &GetAttributeNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取扩展字段名称
func (c *Client) GetAttributeName(request *GetAttributeNameRequest) (response *GetAttributeNameResponse, err error) {
	if request == nil {
		request = NewGetAttributeNameRequest()
	}
	response = NewGetAttributeNameResponse()
	err = c.Send(request, response)
	return
}

func NewOmitOffsiteDeviceRequest() (request *OmitOffsiteDeviceRequest) {
	request = &OmitOffsiteDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "OmitOffsiteDevice")
	return
}

func NewOmitOffsiteDeviceResponse() (response *OmitOffsiteDeviceResponse) {
	response = &OmitOffsiteDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 忽略非常用设备登录记录
func (c *Client) OmitOffsiteDevice(request *OmitOffsiteDeviceRequest) (response *OmitOffsiteDeviceResponse, err error) {
	if request == nil {
		request = NewOmitOffsiteDeviceRequest()
	}
	response = NewOmitOffsiteDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewTokenBindRequest() (request *TokenBindRequest) {
	request = &TokenBindRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "TokenBind")
	return
}

func NewTokenBindResponse() (response *TokenBindResponse) {
	response = &TokenBindResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定账号mfa认证token
func (c *Client) TokenBind(request *TokenBindRequest) (response *TokenBindResponse, err error) {
	if request == nil {
		request = NewTokenBindRequest()
	}
	response = NewTokenBindResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaskedUserInfoRequest() (request *GetMaskedUserInfoRequest) {
	request = &GetMaskedUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetMaskedUserInfo")
	return
}

func NewGetMaskedUserInfoResponse() (response *GetMaskedUserInfoResponse) {
	response = &GetMaskedUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取脱敏用户信息
func (c *Client) GetMaskedUserInfo(request *GetMaskedUserInfoRequest) (response *GetMaskedUserInfoResponse, err error) {
	if request == nil {
		request = NewGetMaskedUserInfoRequest()
	}
	response = NewGetMaskedUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetMasterListWithStatusRequest() (request *GetMasterListWithStatusRequest) {
	request = &GetMasterListWithStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetMasterListWithStatus")
	return
}

func NewGetMasterListWithStatusResponse() (response *GetMasterListWithStatusResponse) {
	response = &GetMasterListWithStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主账户列表，提供状态信息
func (c *Client) GetMasterListWithStatus(request *GetMasterListWithStatusRequest) (response *GetMasterListWithStatusResponse, err error) {
	if request == nil {
		request = NewGetMasterListWithStatusRequest()
	}
	response = NewGetMasterListWithStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetAccountLoginStatusRequest() (request *GetAccountLoginStatusRequest) {
	request = &GetAccountLoginStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetAccountLoginStatus")
	return
}

func NewGetAccountLoginStatusResponse() (response *GetAccountLoginStatusResponse) {
	response = &GetAccountLoginStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账号的登陆状态。 包括正常，锁定等。
func (c *Client) GetAccountLoginStatus(request *GetAccountLoginStatusRequest) (response *GetAccountLoginStatusResponse, err error) {
	if request == nil {
		request = NewGetAccountLoginStatusRequest()
	}
	response = NewGetAccountLoginStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetTradeRequest() (request *GetTradeRequest) {
	request = &GetTradeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetTrade")
	return
}

func NewGetTradeResponse() (response *GetTradeResponse) {
	response = &GetTradeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取行业列表
func (c *Client) GetTrade(request *GetTradeRequest) (response *GetTradeResponse, err error) {
	if request == nil {
		request = NewGetTradeRequest()
	}
	response = NewGetTradeResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserIdAttrRequest() (request *GetUserIdAttrRequest) {
	request = &GetUserIdAttrRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetUserIdAttr")
	return
}

func NewGetUserIdAttrResponse() (response *GetUserIdAttrResponse) {
	response = &GetUserIdAttrResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量交换uin/appid
func (c *Client) GetUserIdAttr(request *GetUserIdAttrRequest) (response *GetUserIdAttrResponse, err error) {
	if request == nil {
		request = NewGetUserIdAttrRequest()
	}
	response = NewGetUserIdAttrResponse()
	err = c.Send(request, response)
	return
}

func NewClearLoginFlagRequest() (request *ClearLoginFlagRequest) {
	request = &ClearLoginFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "ClearLoginFlag")
	return
}

func NewClearLoginFlagResponse() (response *ClearLoginFlagResponse) {
	response = &ClearLoginFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清理登陆态
func (c *Client) ClearLoginFlag(request *ClearLoginFlagRequest) (response *ClearLoginFlagResponse, err error) {
	if request == nil {
		request = NewClearLoginFlagRequest()
	}
	response = NewClearLoginFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDisableSelfApiKeyRequest() (request *DisableSelfApiKeyRequest) {
	request = &DisableSelfApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DisableSelfApiKey")
	return
}

func NewDisableSelfApiKeyResponse() (response *DisableSelfApiKeyResponse) {
	response = &DisableSelfApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用自己的密钥
func (c *Client) DisableSelfApiKey(request *DisableSelfApiKeyRequest) (response *DisableSelfApiKeyResponse, err error) {
	if request == nil {
		request = NewDisableSelfApiKeyRequest()
	}
	response = NewDisableSelfApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetSecurityLastUsedRequest() (request *GetSecurityLastUsedRequest) {
	request = &GetSecurityLastUsedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetSecurityLastUsed")
	return
}

func NewGetSecurityLastUsedResponse() (response *GetSecurityLastUsedResponse) {
	response = &GetSecurityLastUsedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetSecurityLastUsed
func (c *Client) GetSecurityLastUsed(request *GetSecurityLastUsedRequest) (response *GetSecurityLastUsedResponse, err error) {
	if request == nil {
		request = NewGetSecurityLastUsedRequest()
	}
	response = NewGetSecurityLastUsedResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserAreaRequest() (request *GetUserAreaRequest) {
	request = &GetUserAreaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetUserArea")
	return
}

func NewGetUserAreaResponse() (response *GetUserAreaResponse) {
	response = &GetUserAreaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户地域信息
func (c *Client) GetUserArea(request *GetUserAreaRequest) (response *GetUserAreaResponse, err error) {
	if request == nil {
		request = NewGetUserAreaRequest()
	}
	response = NewGetUserAreaResponse()
	err = c.Send(request, response)
	return
}

func NewCheckMailPasswordRequest() (request *CheckMailPasswordRequest) {
	request = &CheckMailPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "CheckMailPassword")
	return
}

func NewCheckMailPasswordResponse() (response *CheckMailPasswordResponse) {
	response = &CheckMailPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验输入账户密码是否正确
func (c *Client) CheckMailPassword(request *CheckMailPasswordRequest) (response *CheckMailPasswordResponse, err error) {
	if request == nil {
		request = NewCheckMailPasswordRequest()
	}
	response = NewCheckMailPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewQuserGetUserInfoRequest() (request *QuserGetUserInfoRequest) {
	request = &QuserGetUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "QuserGetUserInfo")
	return
}

func NewQuserGetUserInfoResponse() (response *QuserGetUserInfoResponse) {
	response = &QuserGetUserInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin查询用户详细信息，可指定查询字段
func (c *Client) QuserGetUserInfo(request *QuserGetUserInfoRequest) (response *QuserGetUserInfoResponse, err error) {
	if request == nil {
		request = NewQuserGetUserInfoRequest()
	}
	response = NewQuserGetUserInfoResponse()
	err = c.Send(request, response)
	return
}

func NewSetAttributeValuesRequest() (request *SetAttributeValuesRequest) {
	request = &SetAttributeValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "SetAttributeValues")
	return
}

func NewSetAttributeValuesResponse() (response *SetAttributeValuesResponse) {
	response = &SetAttributeValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置扩展属性值
func (c *Client) SetAttributeValues(request *SetAttributeValuesRequest) (response *SetAttributeValuesResponse, err error) {
	if request == nil {
		request = NewSetAttributeValuesRequest()
	}
	response = NewSetAttributeValuesResponse()
	err = c.Send(request, response)
	return
}

func NewGetMasterListV2Request() (request *GetMasterListV2Request) {
	request = &GetMasterListV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetMasterListV2")
	return
}

func NewGetMasterListV2Response() (response *GetMasterListV2Response) {
	response = &GetMasterListV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin查询主账户列表
func (c *Client) GetMasterListV2(request *GetMasterListV2Request) (response *GetMasterListV2Response, err error) {
	if request == nil {
		request = NewGetMasterListV2Request()
	}
	response = NewGetMasterListV2Response()
	err = c.Send(request, response)
	return
}

func NewRecoverRetentionSubaccountRequest() (request *RecoverRetentionSubaccountRequest) {
	request = &RecoverRetentionSubaccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "RecoverRetentionSubaccount")
	return
}

func NewRecoverRetentionSubaccountResponse() (response *RecoverRetentionSubaccountResponse) {
	response = &RecoverRetentionSubaccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 恢复回收站账号
func (c *Client) RecoverRetentionSubaccount(request *RecoverRetentionSubaccountRequest) (response *RecoverRetentionSubaccountResponse, err error) {
	if request == nil {
		request = NewRecoverRetentionSubaccountRequest()
	}
	response = NewRecoverRetentionSubaccountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRetentionSubaccountRequest() (request *DeleteRetentionSubaccountRequest) {
	request = &DeleteRetentionSubaccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "DeleteRetentionSubaccount")
	return
}

func NewDeleteRetentionSubaccountResponse() (response *DeleteRetentionSubaccountResponse) {
	response = &DeleteRetentionSubaccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除回收站账号
func (c *Client) DeleteRetentionSubaccount(request *DeleteRetentionSubaccountRequest) (response *DeleteRetentionSubaccountResponse, err error) {
	if request == nil {
		request = NewDeleteRetentionSubaccountRequest()
	}
	response = NewDeleteRetentionSubaccountResponse()
	err = c.Send(request, response)
	return
}

func NewGetLastLoginInfoRequest() (request *GetLastLoginInfoRequest) {
	request = &GetLastLoginInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetLastLoginInfo")
	return
}

func NewGetLastLoginInfoResponse() (response *GetLastLoginInfoResponse) {
	response = &GetLastLoginInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取上一次的登录历史记录（ip，时间，登录方式）
func (c *Client) GetLastLoginInfo(request *GetLastLoginInfoRequest) (response *GetLastLoginInfoResponse, err error) {
	if request == nil {
		request = NewGetLastLoginInfoRequest()
	}
	response = NewGetLastLoginInfoResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBindAccountByUinRequest() (request *QueryBindAccountByUinRequest) {
	request = &QueryBindAccountByUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "QueryBindAccountByUin")
	return
}

func NewQueryBindAccountByUinResponse() (response *QueryBindAccountByUinResponse) {
	response = &QueryBindAccountByUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据uin查询账户信息
func (c *Client) QueryBindAccountByUin(request *QueryBindAccountByUinRequest) (response *QueryBindAccountByUinResponse, err error) {
	if request == nil {
		request = NewQueryBindAccountByUinRequest()
	}
	response = NewQueryBindAccountByUinResponse()
	err = c.Send(request, response)
	return
}

func NewGetMaskedUserInfoByLoginUinRequest() (request *GetMaskedUserInfoByLoginUinRequest) {
	request = &GetMaskedUserInfoByLoginUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("account", APIVersion, "GetMaskedUserInfoByLoginUin")
	return
}

func NewGetMaskedUserInfoByLoginUinResponse() (response *GetMaskedUserInfoByLoginUinResponse) {
	response = &GetMaskedUserInfoByLoginUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetMaskedUserInfoByLoginUin
func (c *Client) GetMaskedUserInfoByLoginUin(request *GetMaskedUserInfoByLoginUinRequest) (response *GetMaskedUserInfoByLoginUinResponse, err error) {
	if request == nil {
		request = NewGetMaskedUserInfoByLoginUinRequest()
	}
	response = NewGetMaskedUserInfoByLoginUinResponse()
	err = c.Send(request, response)
	return
}

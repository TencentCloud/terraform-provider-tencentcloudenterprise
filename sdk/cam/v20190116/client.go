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

package v20190116

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-01-16"

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

func NewGetUserListByUinListRequest() (request *GetUserListByUinListRequest) {
	request = &GetUserListByUinListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetUserListByUinList")
	return
}

func NewGetUserListByUinListResponse() (response *GetUserListByUinListResponse) {
	response = &GetUserListByUinListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetUserListByUinList
func (c *Client) GetUserListByUinList(request *GetUserListByUinListRequest) (response *GetUserListByUinListResponse, err error) {
	if request == nil {
		request = NewGetUserListByUinListRequest()
	}
	response = NewGetUserListByUinListResponse()
	err = c.Send(request, response)
	return
}

func NewGetRoleRequest() (request *GetRoleRequest) {
	request = &GetRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetRole")
	return
}

func NewGetRoleResponse() (response *GetRoleResponse) {
	response = &GetRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GetRole）用于获取指定角色的详细信息。
func (c *Client) GetRole(request *GetRoleRequest) (response *GetRoleResponse, err error) {
	if request == nil {
		request = NewGetRoleRequest()
	}
	response = NewGetRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
	request = &DeleteUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteUser")
	return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
	response = &DeleteUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除子帐号
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
	if request == nil {
		request = NewDeleteUserRequest()
	}
	response = NewDeleteUserResponse()
	err = c.Send(request, response)
	return
}

func NewDisableCollApiKeyRequest() (request *DisableCollApiKeyRequest) {
	request = &DisableCollApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DisableCollApiKey")
	return
}

func NewDisableCollApiKeyResponse() (response *DisableCollApiKeyResponse) {
	response = &DisableCollApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用密钥
func (c *Client) DisableCollApiKey(request *DisableCollApiKeyRequest) (response *DisableCollApiKeyResponse, err error) {
	if request == nil {
		request = NewDisableCollApiKeyRequest()
	}
	response = NewDisableCollApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewBindMailToOrganizationMemberRequest() (request *BindMailToOrganizationMemberRequest) {
	request = &BindMailToOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "BindMailToOrganizationMember")
	return
}

func NewBindMailToOrganizationMemberResponse() (response *BindMailToOrganizationMemberResponse) {
	response = &BindMailToOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// BindMailToOrganizationMember
func (c *Client) BindMailToOrganizationMember(request *BindMailToOrganizationMemberRequest) (response *BindMailToOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewBindMailToOrganizationMemberRequest()
	}
	response = NewBindMailToOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewListPoliciesRequest() (request *ListPoliciesRequest) {
	request = &ListPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListPolicies")
	return
}

func NewListPoliciesResponse() (response *ListPoliciesResponse) {
	response = &ListPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ListPolicies）可用于查询策略列表。
func (c *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
	if request == nil {
		request = NewListPoliciesRequest()
	}
	response = NewListPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCollApiKeyRequest() (request *EnableCollApiKeyRequest) {
	request = &EnableCollApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "EnableCollApiKey")
	return
}

func NewEnableCollApiKeyResponse() (response *EnableCollApiKeyResponse) {
	response = &EnableCollApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用密钥
func (c *Client) EnableCollApiKey(request *EnableCollApiKeyRequest) (response *EnableCollApiKeyResponse, err error) {
	if request == nil {
		request = NewEnableCollApiKeyRequest()
	}
	response = NewEnableCollApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubReceiverRequest() (request *DescribeSubReceiverRequest) {
	request = &DescribeSubReceiverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DescribeSubReceiver")
	return
}

func NewDescribeSubReceiverResponse() (response *DescribeSubReceiverResponse) {
	response = &DescribeSubReceiverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 子账号查询联系方式
func (c *Client) DescribeSubReceiver(request *DescribeSubReceiverRequest) (response *DescribeSubReceiverResponse, err error) {
	if request == nil {
		request = NewDescribeSubReceiverRequest()
	}
	response = NewDescribeSubReceiverResponse()
	err = c.Send(request, response)
	return
}

func NewCancelModifyContactRequest() (request *CancelModifyContactRequest) {
	request = &CancelModifyContactRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CancelModifyContact")
	return
}

func NewCancelModifyContactResponse() (response *CancelModifyContactResponse) {
	response = &CancelModifyContactResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消更换联系信息
func (c *Client) CancelModifyContact(request *CancelModifyContactRequest) (response *CancelModifyContactResponse, err error) {
	if request == nil {
		request = NewCancelModifyContactRequest()
	}
	response = NewCancelModifyContactResponse()
	err = c.Send(request, response)
	return
}

func NewEnableApiKeyRequest() (request *EnableApiKeyRequest) {
	request = &EnableApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "EnableApiKey")
	return
}

func NewEnableApiKeyResponse() (response *EnableApiKeyResponse) {
	response = &EnableApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用密钥
func (c *Client) EnableApiKey(request *EnableApiKeyRequest) (response *EnableApiKeyResponse, err error) {
	if request == nil {
		request = NewEnableApiKeyRequest()
	}
	response = NewEnableApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetServiceRoleInfoRequest() (request *GetServiceRoleInfoRequest) {
	request = &GetServiceRoleInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetServiceRoleInfo")
	return
}

func NewGetServiceRoleInfoResponse() (response *GetServiceRoleInfoResponse) {
	response = &GetServiceRoleInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务角色信息
func (c *Client) GetServiceRoleInfo(request *GetServiceRoleInfoRequest) (response *GetServiceRoleInfoResponse, err error) {
	if request == nil {
		request = NewGetServiceRoleInfoRequest()
	}
	response = NewGetServiceRoleInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEmailRequest() (request *ModifyEmailRequest) {
	request = &ModifyEmailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ModifyEmail")
	return
}

func NewModifyEmailResponse() (response *ModifyEmailResponse) {
	response = &ModifyEmailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改邮箱
func (c *Client) ModifyEmail(request *ModifyEmailRequest) (response *ModifyEmailResponse, err error) {
	if request == nil {
		request = NewModifyEmailRequest()
	}
	response = NewModifyEmailResponse()
	err = c.Send(request, response)
	return
}

func NewListAttachedUserAllPoliciesRequest() (request *ListAttachedUserAllPoliciesRequest) {
	request = &ListAttachedUserAllPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListAttachedUserAllPolicies")
	return
}

func NewListAttachedUserAllPoliciesResponse() (response *ListAttachedUserAllPoliciesResponse) {
	response = &ListAttachedUserAllPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 列出用户关联的策略（包括随组关联）
func (c *Client) ListAttachedUserAllPolicies(request *ListAttachedUserAllPoliciesRequest) (response *ListAttachedUserAllPoliciesResponse, err error) {
	if request == nil {
		request = NewListAttachedUserAllPoliciesRequest()
	}
	response = NewListAttachedUserAllPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserAccessTokenRequest() (request *GetUserAccessTokenRequest) {
	request = &GetUserAccessTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetUserAccessToken")
	return
}

func NewGetUserAccessTokenResponse() (response *GetUserAccessTokenResponse) {
	response = &GetUserAccessTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户第三方开放平台的access token
func (c *Client) GetUserAccessToken(request *GetUserAccessTokenRequest) (response *GetUserAccessTokenResponse, err error) {
	if request == nil {
		request = NewGetUserAccessTokenRequest()
	}
	response = NewGetUserAccessTokenResponse()
	err = c.Send(request, response)
	return
}

func NewGetPolicyRequest() (request *GetPolicyRequest) {
	request = &GetPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetPolicy")
	return
}

func NewGetPolicyResponse() (response *GetPolicyResponse) {
	response = &GetPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（GetPolicy）可用于查询查看策略详情。
func (c *Client) GetPolicy(request *GetPolicyRequest) (response *GetPolicyResponse, err error) {
	if request == nil {
		request = NewGetPolicyRequest()
	}
	response = NewGetPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewBatchOperateCamStrategyRequest() (request *BatchOperateCamStrategyRequest) {
	request = &BatchOperateCamStrategyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "BatchOperateCamStrategy")
	return
}

func NewBatchOperateCamStrategyResponse() (response *BatchOperateCamStrategyResponse) {
	response = &BatchOperateCamStrategyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量关联-移除策略
func (c *Client) BatchOperateCamStrategy(request *BatchOperateCamStrategyRequest) (response *BatchOperateCamStrategyResponse, err error) {
	if request == nil {
		request = NewBatchOperateCamStrategyRequest()
	}
	response = NewBatchOperateCamStrategyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApiKeyRequest() (request *CreateApiKeyRequest) {
	request = &CreateApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CreateApiKey")
	return
}

func NewCreateApiKeyResponse() (response *CreateApiKeyResponse) {
	response = &CreateApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建密钥
func (c *Client) CreateApiKey(request *CreateApiKeyRequest) (response *CreateApiKeyResponse, err error) {
	if request == nil {
		request = NewCreateApiKeyRequest()
	}
	response = NewCreateApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDetachUserPoliciesRequest() (request *DetachUserPoliciesRequest) {
	request = &DetachUserPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DetachUserPolicies")
	return
}

func NewDetachUserPoliciesResponse() (response *DetachUserPoliciesResponse) {
	response = &DetachUserPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除绑定多个策略到用户
func (c *Client) DetachUserPolicies(request *DetachUserPoliciesRequest) (response *DetachUserPoliciesResponse, err error) {
	if request == nil {
		request = NewDetachUserPoliciesRequest()
	}
	response = NewDetachUserPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAccountRequest() (request *CheckAccountRequest) {
	request = &CheckAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CheckAccount")
	return
}

func NewCheckAccountResponse() (response *CheckAccountResponse) {
	response = &CheckAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检验账号
func (c *Client) CheckAccount(request *CheckAccountRequest) (response *CheckAccountResponse, err error) {
	if request == nil {
		request = NewCheckAccountRequest()
	}
	response = NewCheckAccountResponse()
	err = c.Send(request, response)
	return
}

func NewAttachRolesPolicyRequest() (request *AttachRolesPolicyRequest) {
	request = &AttachRolesPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AttachRolesPolicy")
	return
}

func NewAttachRolesPolicyResponse() (response *AttachRolesPolicyResponse) {
	response = &AttachRolesPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定多个角色到策略
func (c *Client) AttachRolesPolicy(request *AttachRolesPolicyRequest) (response *AttachRolesPolicyResponse, err error) {
	if request == nil {
		request = NewAttachRolesPolicyRequest()
	}
	response = NewAttachRolesPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewRefreshUserTokenRequest() (request *RefreshUserTokenRequest) {
	request = &RefreshUserTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "RefreshUserToken")
	return
}

func NewRefreshUserTokenResponse() (response *RefreshUserTokenResponse) {
	response = &RefreshUserTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 刷新用户第三方access_token
func (c *Client) RefreshUserToken(request *RefreshUserTokenRequest) (response *RefreshUserTokenResponse, err error) {
	if request == nil {
		request = NewRefreshUserTokenRequest()
	}
	response = NewRefreshUserTokenResponse()
	err = c.Send(request, response)
	return
}

func NewAttachRolePoliciesRequest() (request *AttachRolePoliciesRequest) {
	request = &AttachRolePoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AttachRolePolicies")
	return
}

func NewAttachRolePoliciesResponse() (response *AttachRolePoliciesResponse) {
	response = &AttachRolePoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定多个策略到角色
func (c *Client) AttachRolePolicies(request *AttachRolePoliciesRequest) (response *AttachRolePoliciesResponse, err error) {
	if request == nil {
		request = NewAttachRolePoliciesRequest()
	}
	response = NewAttachRolePoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckUserPolicyAttachmentRequest() (request *CheckUserPolicyAttachmentRequest) {
	request = &CheckUserPolicyAttachmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CheckUserPolicyAttachment")
	return
}

func NewCheckUserPolicyAttachmentResponse() (response *CheckUserPolicyAttachmentResponse) {
	response = &CheckUserPolicyAttachmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CheckUserPolicyAttachment）用于查询用户是否关联策略。
func (c *Client) CheckUserPolicyAttachment(request *CheckUserPolicyAttachmentRequest) (response *CheckUserPolicyAttachmentResponse, err error) {
	if request == nil {
		request = NewCheckUserPolicyAttachmentRequest()
	}
	response = NewCheckUserPolicyAttachmentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceLinkedRoleRequest() (request *DescribeServiceLinkedRoleRequest) {
	request = &DescribeServiceLinkedRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DescribeServiceLinkedRole")
	return
}

func NewDescribeServiceLinkedRoleResponse() (response *DescribeServiceLinkedRoleResponse) {
	response = &DescribeServiceLinkedRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定服务相关角色的详细信息。
func (c *Client) DescribeServiceLinkedRole(request *DescribeServiceLinkedRoleRequest) (response *DescribeServiceLinkedRoleResponse, err error) {
	if request == nil {
		request = NewDescribeServiceLinkedRoleRequest()
	}
	response = NewDescribeServiceLinkedRoleResponse()
	err = c.Send(request, response)
	return
}

func NewGetGroupListRequest() (request *GetGroupListRequest) {
	request = &GetGroupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetGroupList")
	return
}

func NewGetGroupListResponse() (response *GetGroupListResponse) {
	response = &GetGroupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户所属于的用户组的信息
func (c *Client) GetGroupList(request *GetGroupListRequest) (response *GetGroupListResponse, err error) {
	if request == nil {
		request = NewGetGroupListRequest()
	}
	response = NewGetGroupListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPhoneNumRequest() (request *ModifyPhoneNumRequest) {
	request = &ModifyPhoneNumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ModifyPhoneNum")
	return
}

func NewModifyPhoneNumResponse() (response *ModifyPhoneNumResponse) {
	response = &ModifyPhoneNumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改手机号
func (c *Client) ModifyPhoneNum(request *ModifyPhoneNumRequest) (response *ModifyPhoneNumResponse, err error) {
	if request == nil {
		request = NewModifyPhoneNumRequest()
	}
	response = NewModifyPhoneNumResponse()
	err = c.Send(request, response)
	return
}

func NewGetGroupsSubAccountRequest() (request *GetGroupsSubAccountRequest) {
	request = &GetGroupsSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetGroupsSubAccount")
	return
}

func NewGetGroupsSubAccountResponse() (response *GetGroupsSubAccountResponse) {
	response = &GetGroupsSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户组下账户列表
func (c *Client) GetGroupsSubAccount(request *GetGroupsSubAccountRequest) (response *GetGroupsSubAccountResponse, err error) {
	if request == nil {
		request = NewGetGroupsSubAccountRequest()
	}
	response = NewGetGroupsSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewQueryCollApiKeyRequest() (request *QueryCollApiKeyRequest) {
	request = &QueryCollApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "QueryCollApiKey")
	return
}

func NewQueryCollApiKeyResponse() (response *QueryCollApiKeyResponse) {
	response = &QueryCollApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取API密钥列表
func (c *Client) QueryCollApiKey(request *QueryCollApiKeyRequest) (response *QueryCollApiKeyResponse, err error) {
	if request == nil {
		request = NewQueryCollApiKeyRequest()
	}
	response = NewQueryCollApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewAddAttributeValuesRequest() (request *AddAttributeValuesRequest) {
	request = &AddAttributeValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AddAttributeValues")
	return
}

func NewAddAttributeValuesResponse() (response *AddAttributeValuesResponse) {
	response = &AddAttributeValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加属性值
func (c *Client) AddAttributeValues(request *AddAttributeValuesRequest) (response *AddAttributeValuesResponse, err error) {
	if request == nil {
		request = NewAddAttributeValuesRequest()
	}
	response = NewAddAttributeValuesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOauthProviderRequest() (request *CreateOauthProviderRequest) {
	request = &CreateOauthProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CreateOauthProvider")
	return
}

func NewCreateOauthProviderResponse() (response *CreateOauthProviderResponse) {
	response = &CreateOauthProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增oauth配置
func (c *Client) CreateOauthProvider(request *CreateOauthProviderRequest) (response *CreateOauthProviderResponse, err error) {
	if request == nil {
		request = NewCreateOauthProviderRequest()
	}
	response = NewCreateOauthProviderResponse()
	err = c.Send(request, response)
	return
}

func NewSeedOIDCLoginTokenRequest() (request *SeedOIDCLoginTokenRequest) {
	request = &SeedOIDCLoginTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "SeedOIDCLoginToken")
	return
}

func NewSeedOIDCLoginTokenResponse() (response *SeedOIDCLoginTokenResponse) {
	response = &SeedOIDCLoginTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取oidc登录的认证token
func (c *Client) SeedOIDCLoginToken(request *SeedOIDCLoginTokenRequest) (response *SeedOIDCLoginTokenResponse, err error) {
	if request == nil {
		request = NewSeedOIDCLoginTokenRequest()
	}
	response = NewSeedOIDCLoginTokenResponse()
	err = c.Send(request, response)
	return
}

func NewQueryKeyBySecretIdRequest() (request *QueryKeyBySecretIdRequest) {
	request = &QueryKeyBySecretIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "QueryKeyBySecretId")
	return
}

func NewQueryKeyBySecretIdResponse() (response *QueryKeyBySecretIdResponse) {
	response = &QueryKeyBySecretIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据SecretId查询个人 API 密钥
func (c *Client) QueryKeyBySecretId(request *QueryKeyBySecretIdRequest) (response *QueryKeyBySecretIdResponse, err error) {
	if request == nil {
		request = NewQueryKeyBySecretIdRequest()
	}
	response = NewQueryKeyBySecretIdResponse()
	err = c.Send(request, response)
	return
}

func NewConfirmCASProviderRequest() (request *ConfirmCASProviderRequest) {
	request = &ConfirmCASProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ConfirmCASProvider")
	return
}

func NewConfirmCASProviderResponse() (response *ConfirmCASProviderResponse) {
	response = &ConfirmCASProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启企业认证登录入口
func (c *Client) ConfirmCASProvider(request *ConfirmCASProviderRequest) (response *ConfirmCASProviderResponse, err error) {
	if request == nil {
		request = NewConfirmCASProviderRequest()
	}
	response = NewConfirmCASProviderResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupBatchRequest() (request *DeleteGroupBatchRequest) {
	request = &DeleteGroupBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteGroupBatch")
	return
}

func NewDeleteGroupBatchResponse() (response *DeleteGroupBatchResponse) {
	response = &DeleteGroupBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除用户组
func (c *Client) DeleteGroupBatch(request *DeleteGroupBatchRequest) (response *DeleteGroupBatchResponse, err error) {
	if request == nil {
		request = NewDeleteGroupBatchRequest()
	}
	response = NewDeleteGroupBatchResponse()
	err = c.Send(request, response)
	return
}

func NewGetSubsGroupRequest() (request *GetSubsGroupRequest) {
	request = &GetSubsGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetSubsGroup")
	return
}

func NewGetSubsGroupResponse() (response *GetSubsGroupResponse) {
	response = &GetSubsGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子账户所属用户组列表
func (c *Client) GetSubsGroup(request *GetSubsGroupRequest) (response *GetSubsGroupResponse, err error) {
	if request == nil {
		request = NewGetSubsGroupRequest()
	}
	response = NewGetSubsGroupResponse()
	err = c.Send(request, response)
	return
}

func NewListEntitiesForPolicyRequest() (request *ListEntitiesForPolicyRequest) {
	request = &ListEntitiesForPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListEntitiesForPolicy")
	return
}

func NewListEntitiesForPolicyResponse() (response *ListEntitiesForPolicyResponse) {
	response = &ListEntitiesForPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ListEntitiesForPolicy）可用于查询策略关联的实体列表。
func (c *Client) ListEntitiesForPolicy(request *ListEntitiesForPolicyRequest) (response *ListEntitiesForPolicyResponse, err error) {
	if request == nil {
		request = NewListEntitiesForPolicyRequest()
	}
	response = NewListEntitiesForPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewQueryApiKeyRequest() (request *QueryApiKeyRequest) {
	request = &QueryApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "QueryApiKey")
	return
}

func NewQueryApiKeyResponse() (response *QueryApiKeyResponse) {
	response = &QueryApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取API密钥列表
func (c *Client) QueryApiKey(request *QueryApiKeyRequest) (response *QueryApiKeyResponse, err error) {
	if request == nil {
		request = NewQueryApiKeyRequest()
	}
	response = NewQueryApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDetachGroupsPolicyRequest() (request *DetachGroupsPolicyRequest) {
	request = &DetachGroupsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DetachGroupsPolicy")
	return
}

func NewDetachGroupsPolicyResponse() (response *DetachGroupsPolicyResponse) {
	response = &DetachGroupsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除绑定策略到多个用户组
func (c *Client) DetachGroupsPolicy(request *DetachGroupsPolicyRequest) (response *DetachGroupsPolicyResponse, err error) {
	if request == nil {
		request = NewDetachGroupsPolicyRequest()
	}
	response = NewDetachGroupsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOauthProviderRequest() (request *UpdateOauthProviderRequest) {
	request = &UpdateOauthProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateOauthProvider")
	return
}

func NewUpdateOauthProviderResponse() (response *UpdateOauthProviderResponse) {
	response = &UpdateOauthProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Oauth配置信息
func (c *Client) UpdateOauthProvider(request *UpdateOauthProviderRequest) (response *UpdateOauthProviderResponse, err error) {
	if request == nil {
		request = NewUpdateOauthProviderRequest()
	}
	response = NewUpdateOauthProviderResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserToGroupRequest() (request *AddUserToGroupRequest) {
	request = &AddUserToGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AddUserToGroup")
	return
}

func NewAddUserToGroupResponse() (response *AddUserToGroupResponse) {
	response = &AddUserToGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户加入到用户组
func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
	if request == nil {
		request = NewAddUserToGroupRequest()
	}
	response = NewAddUserToGroupResponse()
	err = c.Send(request, response)
	return
}

func NewAttachGroupsPolicyRequest() (request *AttachGroupsPolicyRequest) {
	request = &AttachGroupsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AttachGroupsPolicy")
	return
}

func NewAttachGroupsPolicyResponse() (response *AttachGroupsPolicyResponse) {
	response = &AttachGroupsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定策略到多个用户组
func (c *Client) AttachGroupsPolicy(request *AttachGroupsPolicyRequest) (response *AttachGroupsPolicyResponse, err error) {
	if request == nil {
		request = NewAttachGroupsPolicyRequest()
	}
	response = NewAttachGroupsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetAllSubUserRequest() (request *GetAllSubUserRequest) {
	request = &GetAllSubUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetAllSubUser")
	return
}

func NewGetAllSubUserResponse() (response *GetAllSubUserResponse) {
	response = &GetAllSubUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所属主账号下的子账号和用户组信息
func (c *Client) GetAllSubUser(request *GetAllSubUserRequest) (response *GetAllSubUserResponse, err error) {
	if request == nil {
		request = NewGetAllSubUserRequest()
	}
	response = NewGetAllSubUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAttachedEntityPoliciesRequest() (request *DescribeAttachedEntityPoliciesRequest) {
	request = &DescribeAttachedEntityPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DescribeAttachedEntityPolicies")
	return
}

func NewDescribeAttachedEntityPoliciesResponse() (response *DescribeAttachedEntityPoliciesResponse) {
	response = &DescribeAttachedEntityPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查多个用户或用户组关联的策略详情
func (c *Client) DescribeAttachedEntityPolicies(request *DescribeAttachedEntityPoliciesRequest) (response *DescribeAttachedEntityPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeAttachedEntityPoliciesRequest()
	}
	response = NewDescribeAttachedEntityPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewAddSubAccountRequest() (request *AddSubAccountRequest) {
	request = &AddSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AddSubAccount")
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

func NewGetServiceApiListRequest() (request *GetServiceApiListRequest) {
	request = &GetServiceApiListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetServiceApiList")
	return
}

func NewGetServiceApiListResponse() (response *GetServiceApiListResponse) {
	response = &GetServiceApiListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取服务及其API列表
func (c *Client) GetServiceApiList(request *GetServiceApiListRequest) (response *GetServiceApiListResponse, err error) {
	if request == nil {
		request = NewGetServiceApiListRequest()
	}
	response = NewGetServiceApiListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
	request = &DeleteGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteGroup")
	return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
	response = &DeleteGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据用户组id删除用户组
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
	if request == nil {
		request = NewDeleteGroupRequest()
	}
	response = NewDeleteGroupResponse()
	err = c.Send(request, response)
	return
}

func NewListMaskedSubAccountsRequest() (request *ListMaskedSubAccountsRequest) {
	request = &ListMaskedSubAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListMaskedSubAccounts")
	return
}

func NewListMaskedSubAccountsResponse() (response *ListMaskedSubAccountsResponse) {
	response = &ListMaskedSubAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取加密子帐号列表
func (c *Client) ListMaskedSubAccounts(request *ListMaskedSubAccountsRequest) (response *ListMaskedSubAccountsResponse, err error) {
	if request == nil {
		request = NewListMaskedSubAccountsRequest()
	}
	response = NewListMaskedSubAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewListOpenPlatformsRequest() (request *ListOpenPlatformsRequest) {
	request = &ListOpenPlatformsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListOpenPlatforms")
	return
}

func NewListOpenPlatformsResponse() (response *ListOpenPlatformsResponse) {
	response = &ListOpenPlatformsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开放平台列表
func (c *Client) ListOpenPlatforms(request *ListOpenPlatformsRequest) (response *ListOpenPlatformsResponse, err error) {
	if request == nil {
		request = NewListOpenPlatformsRequest()
	}
	response = NewListOpenPlatformsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
	request = &DeleteRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteRole")
	return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
	response = &DeleteRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DeleteRole）用于删除指定角色。
func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
	if request == nil {
		request = NewDeleteRoleRequest()
	}
	response = NewDeleteRoleResponse()
	err = c.Send(request, response)
	return
}

func NewAttachUserPoliciesRequest() (request *AttachUserPoliciesRequest) {
	request = &AttachUserPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AttachUserPolicies")
	return
}

func NewAttachUserPoliciesResponse() (response *AttachUserPoliciesResponse) {
	response = &AttachUserPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定多个策略到用户
func (c *Client) AttachUserPolicies(request *AttachUserPoliciesRequest) (response *AttachUserPoliciesResponse, err error) {
	if request == nil {
		request = NewAttachUserPoliciesRequest()
	}
	response = NewAttachUserPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewListGroupsRequest() (request *ListGroupsRequest) {
	request = &ListGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListGroups")
	return
}

func NewListGroupsResponse() (response *ListGroupsResponse) {
	response = &ListGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户组列表
func (c *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
	if request == nil {
		request = NewListGroupsRequest()
	}
	response = NewListGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCASProviderRequest() (request *CreateCASProviderRequest) {
	request = &CreateCASProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CreateCASProvider")
	return
}

func NewCreateCASProviderResponse() (response *CreateCASProviderResponse) {
	response = &CreateCASProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置身份提供商信息
func (c *Client) CreateCASProvider(request *CreateCASProviderRequest) (response *CreateCASProviderResponse, err error) {
	if request == nil {
		request = NewCreateCASProviderRequest()
	}
	response = NewCreateCASProviderResponse()
	err = c.Send(request, response)
	return
}

func NewAttachGroupPoliciesRequest() (request *AttachGroupPoliciesRequest) {
	request = &AttachGroupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AttachGroupPolicies")
	return
}

func NewAttachGroupPoliciesResponse() (response *AttachGroupPoliciesResponse) {
	response = &AttachGroupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定多个策略到用户组
func (c *Client) AttachGroupPolicies(request *AttachGroupPoliciesRequest) (response *AttachGroupPoliciesResponse, err error) {
	if request == nil {
		request = NewAttachGroupPoliciesRequest()
	}
	response = NewAttachGroupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewCheckStrategyNameRequest() (request *CheckStrategyNameRequest) {
	request = &CheckStrategyNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CheckStrategyName")
	return
}

func NewCheckStrategyNameResponse() (response *CheckStrategyNameResponse) {
	response = &CheckStrategyNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检测策略名是否重复
func (c *Client) CheckStrategyName(request *CheckStrategyNameRequest) (response *CheckStrategyNameResponse, err error) {
	if request == nil {
		request = NewCheckStrategyNameRequest()
	}
	response = NewCheckStrategyNameResponse()
	err = c.Send(request, response)
	return
}

func NewListAttachedRolePoliciesRequest() (request *ListAttachedRolePoliciesRequest) {
	request = &ListAttachedRolePoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListAttachedRolePolicies")
	return
}

func NewListAttachedRolePoliciesResponse() (response *ListAttachedRolePoliciesResponse) {
	response = &ListAttachedRolePoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ListAttachedRolePolicies）用于获取角色绑定的策略列表。
func (c *Client) ListAttachedRolePolicies(request *ListAttachedRolePoliciesRequest) (response *ListAttachedRolePoliciesResponse, err error) {
	if request == nil {
		request = NewListAttachedRolePoliciesRequest()
	}
	response = NewListAttachedRolePoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewSendPhoneVerifyLinkRequest() (request *SendPhoneVerifyLinkRequest) {
	request = &SendPhoneVerifyLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "SendPhoneVerifyLink")
	return
}

func NewSendPhoneVerifyLinkResponse() (response *SendPhoneVerifyLinkResponse) {
	response = &SendPhoneVerifyLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送手机号验证链接
func (c *Client) SendPhoneVerifyLink(request *SendPhoneVerifyLinkRequest) (response *SendPhoneVerifyLinkResponse, err error) {
	if request == nil {
		request = NewSendPhoneVerifyLinkRequest()
	}
	response = NewSendPhoneVerifyLinkResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
	request = &CreateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CreateGroup")
	return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
	response = &CreateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用户组
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
	if request == nil {
		request = NewCreateGroupRequest()
	}
	response = NewCreateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAssumeRolePolicyRequest() (request *UpdateAssumeRolePolicyRequest) {
	request = &UpdateAssumeRolePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateAssumeRolePolicy")
	return
}

func NewUpdateAssumeRolePolicyResponse() (response *UpdateAssumeRolePolicyResponse) {
	response = &UpdateAssumeRolePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateAssumeRolePolicy）用于修改角色信任策略的策略文档。
func (c *Client) UpdateAssumeRolePolicy(request *UpdateAssumeRolePolicyRequest) (response *UpdateAssumeRolePolicyResponse, err error) {
	if request == nil {
		request = NewUpdateAssumeRolePolicyRequest()
	}
	response = NewUpdateAssumeRolePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCopyUserPolicyRequest() (request *CopyUserPolicyRequest) {
	request = &CopyUserPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CopyUserPolicy")
	return
}

func NewCopyUserPolicyResponse() (response *CopyUserPolicyResponse) {
	response = &CopyUserPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 多个用户拷贝某一个用户的权限
func (c *Client) CopyUserPolicy(request *CopyUserPolicyRequest) (response *CopyUserPolicyResponse, err error) {
	if request == nil {
		request = NewCopyUserPolicyRequest()
	}
	response = NewCopyUserPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGroupMemberRequest() (request *UpdateGroupMemberRequest) {
	request = &UpdateGroupMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateGroupMember")
	return
}

func NewUpdateGroupMemberResponse() (response *UpdateGroupMemberResponse) {
	response = &UpdateGroupMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量更新用户组成员，支持增加和删除
func (c *Client) UpdateGroupMember(request *UpdateGroupMemberRequest) (response *UpdateGroupMemberResponse, err error) {
	if request == nil {
		request = NewUpdateGroupMemberRequest()
	}
	response = NewUpdateGroupMemberResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOwnerAccountRequest() (request *UpdateOwnerAccountRequest) {
	request = &UpdateOwnerAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateOwnerAccount")
	return
}

func NewUpdateOwnerAccountResponse() (response *UpdateOwnerAccountResponse) {
	response = &UpdateOwnerAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新主账户属性
func (c *Client) UpdateOwnerAccount(request *UpdateOwnerAccountRequest) (response *UpdateOwnerAccountResponse, err error) {
	if request == nil {
		request = NewUpdateOwnerAccountRequest()
	}
	response = NewUpdateOwnerAccountResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePasswordRulesRequest() (request *UpdatePasswordRulesRequest) {
	request = &UpdatePasswordRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdatePasswordRules")
	return
}

func NewUpdatePasswordRulesResponse() (response *UpdatePasswordRulesResponse) {
	response = &UpdatePasswordRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新用户密码设置规则
func (c *Client) UpdatePasswordRules(request *UpdatePasswordRulesRequest) (response *UpdatePasswordRulesResponse, err error) {
	if request == nil {
		request = NewUpdatePasswordRulesRequest()
	}
	response = NewUpdatePasswordRulesResponse()
	err = c.Send(request, response)
	return
}

func NewGetServicePermListRequest() (request *GetServicePermListRequest) {
	request = &GetServicePermListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetServicePermList")
	return
}

func NewGetServicePermListResponse() (response *GetServicePermListResponse) {
	response = &GetServicePermListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取接口列表
func (c *Client) GetServicePermList(request *GetServicePermListRequest) (response *GetServicePermListResponse, err error) {
	if request == nil {
		request = NewGetServicePermListRequest()
	}
	response = NewGetServicePermListResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiKeyRequest() (request *DeleteApiKeyRequest) {
	request = &DeleteApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteApiKey")
	return
}

func NewDeleteApiKeyResponse() (response *DeleteApiKeyResponse) {
	response = &DeleteApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除密钥
func (c *Client) DeleteApiKey(request *DeleteApiKeyRequest) (response *DeleteApiKeyResponse, err error) {
	if request == nil {
		request = NewDeleteApiKeyRequest()
	}
	response = NewDeleteApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePolicyRequest() (request *DeletePolicyRequest) {
	request = &DeletePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeletePolicy")
	return
}

func NewDeletePolicyResponse() (response *DeletePolicyResponse) {
	response = &DeletePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除策略
func (c *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
	if request == nil {
		request = NewDeletePolicyRequest()
	}
	response = NewDeletePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDisableCASProviderRequest() (request *DisableCASProviderRequest) {
	request = &DisableCASProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DisableCASProvider")
	return
}

func NewDisableCASProviderResponse() (response *DisableCASProviderResponse) {
	response = &DisableCASProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭企业认证登录入口
func (c *Client) DisableCASProvider(request *DisableCASProviderRequest) (response *DisableCASProviderResponse, err error) {
	if request == nil {
		request = NewDisableCASProviderRequest()
	}
	response = NewDisableCASProviderResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserInfoRequest() (request *GetUserInfoRequest) {
	request = &GetUserInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetUserInfo")
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

func NewSetLoginRulesRequest() (request *SetLoginRulesRequest) {
	request = &SetLoginRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "SetLoginRules")
	return
}

func NewSetLoginRulesResponse() (response *SetLoginRulesResponse) {
	response = &SetLoginRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置登录策略规则
func (c *Client) SetLoginRules(request *SetLoginRulesRequest) (response *SetLoginRulesResponse, err error) {
	if request == nil {
		request = NewSetLoginRulesRequest()
	}
	response = NewSetLoginRulesResponse()
	err = c.Send(request, response)
	return
}

func NewAttachUsersPolicyRequest() (request *AttachUsersPolicyRequest) {
	request = &AttachUsersPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AttachUsersPolicy")
	return
}

func NewAttachUsersPolicyResponse() (response *AttachUsersPolicyResponse) {
	response = &AttachUsersPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定策略到多个用户
func (c *Client) AttachUsersPolicy(request *AttachUsersPolicyRequest) (response *AttachUsersPolicyResponse, err error) {
	if request == nil {
		request = NewAttachUsersPolicyRequest()
	}
	response = NewAttachUsersPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCasProviderRequest() (request *DescribeCasProviderRequest) {
	request = &DescribeCasProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DescribeCasProvider")
	return
}

func NewDescribeCasProviderResponse() (response *DescribeCasProviderResponse) {
	response = &DescribeCasProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 企业认证登录管理
func (c *Client) DescribeCasProvider(request *DescribeCasProviderRequest) (response *DescribeCasProviderResponse, err error) {
	if request == nil {
		request = NewDescribeCasProviderRequest()
	}
	response = NewDescribeCasProviderResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSubAccountRequest() (request *UpdateSubAccountRequest) {
	request = &UpdateSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateSubAccount")
	return
}

func NewUpdateSubAccountResponse() (response *UpdateSubAccountResponse) {
	response = &UpdateSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新子账户信息
func (c *Client) UpdateSubAccount(request *UpdateSubAccountRequest) (response *UpdateSubAccountResponse, err error) {
	if request == nil {
		request = NewUpdateSubAccountRequest()
	}
	response = NewUpdateSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewListUsersForGroupRequest() (request *ListUsersForGroupRequest) {
	request = &ListUsersForGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListUsersForGroup")
	return
}

func NewListUsersForGroupResponse() (response *ListUsersForGroupResponse) {
	response = &ListUsersForGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户组关联的用户列表
func (c *Client) ListUsersForGroup(request *ListUsersForGroupRequest) (response *ListUsersForGroupResponse, err error) {
	if request == nil {
		request = NewListUsersForGroupRequest()
	}
	response = NewListUsersForGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRoleStatusRequest() (request *UpdateRoleStatusRequest) {
	request = &UpdateRoleStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateRoleStatus")
	return
}

func NewUpdateRoleStatusResponse() (response *UpdateRoleStatusResponse) {
	response = &UpdateRoleStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新角色状态
func (c *Client) UpdateRoleStatus(request *UpdateRoleStatusRequest) (response *UpdateRoleStatusResponse, err error) {
	if request == nil {
		request = NewUpdateRoleStatusRequest()
	}
	response = NewUpdateRoleStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCheckPolicyRequest() (request *CheckPolicyRequest) {
	request = &CheckPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CheckPolicy")
	return
}

func NewCheckPolicyResponse() (response *CheckPolicyResponse) {
	response = &CheckPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CheckPolicy）可用于检查策略语法。
func (c *Client) CheckPolicy(request *CheckPolicyRequest) (response *CheckPolicyResponse, err error) {
	if request == nil {
		request = NewCheckPolicyRequest()
	}
	response = NewCheckPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserRequest() (request *AddUserRequest) {
	request = &AddUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AddUser")
	return
}

func NewAddUserResponse() (response *AddUserResponse) {
	response = &AddUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建子用户
func (c *Client) AddUser(request *AddUserRequest) (response *AddUserResponse, err error) {
	if request == nil {
		request = NewAddUserRequest()
	}
	response = NewAddUserResponse()
	err = c.Send(request, response)
	return
}

func NewDetachUsersPolicyRequest() (request *DetachUsersPolicyRequest) {
	request = &DetachUsersPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DetachUsersPolicy")
	return
}

func NewDetachUsersPolicyResponse() (response *DetachUsersPolicyResponse) {
	response = &DetachUsersPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除绑定策略到多个用户
func (c *Client) DetachUsersPolicy(request *DetachUsersPolicyRequest) (response *DetachUsersPolicyResponse, err error) {
	if request == nil {
		request = NewDetachUsersPolicyRequest()
	}
	response = NewDetachUsersPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewListAttachedGroupPoliciesRequest() (request *ListAttachedGroupPoliciesRequest) {
	request = &ListAttachedGroupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListAttachedGroupPolicies")
	return
}

func NewListAttachedGroupPoliciesResponse() (response *ListAttachedGroupPoliciesResponse) {
	response = &ListAttachedGroupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户组关联的策略列表
func (c *Client) ListAttachedGroupPolicies(request *ListAttachedGroupPoliciesRequest) (response *ListAttachedGroupPoliciesResponse, err error) {
	if request == nil {
		request = NewListAttachedGroupPoliciesRequest()
	}
	response = NewListAttachedGroupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubAccountRequest() (request *DeleteSubAccountRequest) {
	request = &DeleteSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteSubAccount")
	return
}

func NewDeleteSubAccountResponse() (response *DeleteSubAccountResponse) {
	response = &DeleteSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除子帐号
func (c *Client) DeleteSubAccount(request *DeleteSubAccountRequest) (response *DeleteSubAccountResponse, err error) {
	if request == nil {
		request = NewDeleteSubAccountRequest()
	}
	response = NewDeleteSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDetachGroupPoliciesRequest() (request *DetachGroupPoliciesRequest) {
	request = &DetachGroupPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DetachGroupPolicies")
	return
}

func NewDetachGroupPoliciesResponse() (response *DetachGroupPoliciesResponse) {
	response = &DetachGroupPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除绑定多个策略到用户组
func (c *Client) DetachGroupPolicies(request *DetachGroupPoliciesRequest) (response *DetachGroupPoliciesResponse, err error) {
	if request == nil {
		request = NewDetachGroupPoliciesRequest()
	}
	response = NewDetachGroupPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCollApiKeyRequest() (request *DeleteCollApiKeyRequest) {
	request = &DeleteCollApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteCollApiKey")
	return
}

func NewDeleteCollApiKeyResponse() (response *DeleteCollApiKeyResponse) {
	response = &DeleteCollApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除密钥
func (c *Client) DeleteCollApiKey(request *DeleteCollApiKeyRequest) (response *DeleteCollApiKeyResponse, err error) {
	if request == nil {
		request = NewDeleteCollApiKeyRequest()
	}
	response = NewDeleteCollApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetUinBySecretIdRequest() (request *GetUinBySecretIdRequest) {
	request = &GetUinBySecretIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetUinBySecretId")
	return
}

func NewGetUinBySecretIdResponse() (response *GetUinBySecretIdResponse) {
	response = &GetUinBySecretIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据SecretId查询Uin
func (c *Client) GetUinBySecretId(request *GetUinBySecretIdRequest) (response *GetUinBySecretIdResponse, err error) {
	if request == nil {
		request = NewGetUinBySecretIdRequest()
	}
	response = NewGetUinBySecretIdResponse()
	err = c.Send(request, response)
	return
}

func NewListUsersForPolicyRequest() (request *ListUsersForPolicyRequest) {
	request = &ListUsersForPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListUsersForPolicy")
	return
}

func NewListUsersForPolicyResponse() (response *ListUsersForPolicyResponse) {
	response = &ListUsersForPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ListUsersForPolicy）用于列出策略关联的用户列表（包括随组关联）。
func (c *Client) ListUsersForPolicy(request *ListUsersForPolicyRequest) (response *ListUsersForPolicyResponse, err error) {
	if request == nil {
		request = NewListUsersForPolicyRequest()
	}
	response = NewListUsersForPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePolicyRequest() (request *CreatePolicyRequest) {
	request = &CreatePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CreatePolicy")
	return
}

func NewCreatePolicyResponse() (response *CreatePolicyResponse) {
	response = &CreatePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreatePolicy）可用于创建策略。
func (c *Client) CreatePolicy(request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
	if request == nil {
		request = NewCreatePolicyRequest()
	}
	response = NewCreatePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewGetApiKeyRequest() (request *GetApiKeyRequest) {
	request = &GetApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetApiKey")
	return
}

func NewGetApiKeyResponse() (response *GetApiKeyResponse) {
	response = &GetApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取密钥
func (c *Client) GetApiKey(request *GetApiKeyRequest) (response *GetApiKeyResponse, err error) {
	if request == nil {
		request = NewGetApiKeyRequest()
	}
	response = NewGetApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCASProviderRequest() (request *UpdateCASProviderRequest) {
	request = &UpdateCASProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateCASProvider")
	return
}

func NewUpdateCASProviderResponse() (response *UpdateCASProviderResponse) {
	response = &UpdateCASProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新身份提供商信息
func (c *Client) UpdateCASProvider(request *UpdateCASProviderRequest) (response *UpdateCASProviderResponse, err error) {
	if request == nil {
		request = NewUpdateCASProviderRequest()
	}
	response = NewUpdateCASProviderResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupListRequest() (request *DeleteGroupListRequest) {
	request = &DeleteGroupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteGroupList")
	return
}

func NewDeleteGroupListResponse() (response *DeleteGroupListResponse) {
	response = &DeleteGroupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据用户组id删除用户组
func (c *Client) DeleteGroupList(request *DeleteGroupListRequest) (response *DeleteGroupListResponse, err error) {
	if request == nil {
		request = NewDeleteGroupListRequest()
	}
	response = NewDeleteGroupListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubAccountContactsRequest() (request *DescribeSubAccountContactsRequest) {
	request = &DescribeSubAccountContactsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DescribeSubAccountContacts")
	return
}

func NewDescribeSubAccountContactsResponse() (response *DescribeSubAccountContactsResponse) {
	response = &DescribeSubAccountContactsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子账号联系方式列表
func (c *Client) DescribeSubAccountContacts(request *DescribeSubAccountContactsRequest) (response *DescribeSubAccountContactsResponse, err error) {
	if request == nil {
		request = NewDescribeSubAccountContactsRequest()
	}
	response = NewDescribeSubAccountContactsResponse()
	err = c.Send(request, response)
	return
}

func NewDetachRolesPolicyRequest() (request *DetachRolesPolicyRequest) {
	request = &DetachRolesPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DetachRolesPolicy")
	return
}

func NewDetachRolesPolicyResponse() (response *DetachRolesPolicyResponse) {
	response = &DetachRolesPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除绑定多个角色到策略
func (c *Client) DetachRolesPolicy(request *DetachRolesPolicyRequest) (response *DetachRolesPolicyResponse, err error) {
	if request == nil {
		request = NewDetachRolesPolicyRequest()
	}
	response = NewDetachRolesPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewSetAccountLoginStatusRequest() (request *SetAccountLoginStatusRequest) {
	request = &SetAccountLoginStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "SetAccountLoginStatus")
	return
}

func NewSetAccountLoginStatusResponse() (response *SetAccountLoginStatusResponse) {
	response = &SetAccountLoginStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置登陆状态，正常/锁定/临时锁定
func (c *Client) SetAccountLoginStatus(request *SetAccountLoginStatusRequest) (response *SetAccountLoginStatusResponse, err error) {
	if request == nil {
		request = NewSetAccountLoginStatusRequest()
	}
	response = NewSetAccountLoginStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCollApiKeyRequest() (request *CreateCollApiKeyRequest) {
	request = &CreateCollApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CreateCollApiKey")
	return
}

func NewCreateCollApiKeyResponse() (response *CreateCollApiKeyResponse) {
	response = &CreateCollApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建密钥
func (c *Client) CreateCollApiKey(request *CreateCollApiKeyRequest) (response *CreateCollApiKeyResponse, err error) {
	if request == nil {
		request = NewCreateCollApiKeyRequest()
	}
	response = NewCreateCollApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewGetPasswordRulesRequest() (request *GetPasswordRulesRequest) {
	request = &GetPasswordRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetPasswordRules")
	return
}

func NewGetPasswordRulesResponse() (response *GetPasswordRulesResponse) {
	response = &GetPasswordRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口{GetPasswordRules}用于获取用户的密码设置规则
func (c *Client) GetPasswordRules(request *GetPasswordRulesRequest) (response *GetPasswordRulesResponse, err error) {
	if request == nil {
		request = NewGetPasswordRulesRequest()
	}
	response = NewGetPasswordRulesResponse()
	err = c.Send(request, response)
	return
}

func NewListGroupsForConsoleRequest() (request *ListGroupsForConsoleRequest) {
	request = &ListGroupsForConsoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListGroupsForConsole")
	return
}

func NewListGroupsForConsoleResponse() (response *ListGroupsForConsoleResponse) {
	response = &ListGroupsForConsoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ListGroupsForConsole
func (c *Client) ListGroupsForConsole(request *ListGroupsForConsoleRequest) (response *ListGroupsForConsoleResponse, err error) {
	if request == nil {
		request = NewListGroupsForConsoleRequest()
	}
	response = NewListGroupsForConsoleResponse()
	err = c.Send(request, response)
	return
}

func NewAttachRolePolicyRequest() (request *AttachRolePolicyRequest) {
	request = &AttachRolePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AttachRolePolicy")
	return
}

func NewAttachRolePolicyResponse() (response *AttachRolePolicyResponse) {
	response = &AttachRolePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（AttachRolePolicy）用于绑定策略到角色。
func (c *Client) AttachRolePolicy(request *AttachRolePolicyRequest) (response *AttachRolePolicyResponse, err error) {
	if request == nil {
		request = NewAttachRolePolicyRequest()
	}
	response = NewAttachRolePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceLinkedRoleRequest() (request *DeleteServiceLinkedRoleRequest) {
	request = &DeleteServiceLinkedRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteServiceLinkedRole")
	return
}

func NewDeleteServiceLinkedRoleResponse() (response *DeleteServiceLinkedRoleResponse) {
	response = &DeleteServiceLinkedRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除服务相关角色
func (c *Client) DeleteServiceLinkedRole(request *DeleteServiceLinkedRoleRequest) (response *DeleteServiceLinkedRoleResponse, err error) {
	if request == nil {
		request = NewDeleteServiceLinkedRoleRequest()
	}
	response = NewDeleteServiceLinkedRoleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
	request = &UpdateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateGroup")
	return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
	response = &UpdateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新用户组信息
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
	if request == nil {
		request = NewUpdateGroupRequest()
	}
	response = NewUpdateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewListAllUserGroupRequest() (request *ListAllUserGroupRequest) {
	request = &ListAllUserGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListAllUserGroup")
	return
}

func NewListAllUserGroupResponse() (response *ListAllUserGroupResponse) {
	response = &ListAllUserGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户及用户组列表
func (c *Client) ListAllUserGroup(request *ListAllUserGroupRequest) (response *ListAllUserGroupResponse, err error) {
	if request == nil {
		request = NewListAllUserGroupRequest()
	}
	response = NewListAllUserGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRoleDescriptionRequest() (request *UpdateRoleDescriptionRequest) {
	request = &UpdateRoleDescriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateRoleDescription")
	return
}

func NewUpdateRoleDescriptionResponse() (response *UpdateRoleDescriptionResponse) {
	response = &UpdateRoleDescriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdateRoleDescription）用于修改角色的描述信息。
func (c *Client) UpdateRoleDescription(request *UpdateRoleDescriptionRequest) (response *UpdateRoleDescriptionResponse, err error) {
	if request == nil {
		request = NewUpdateRoleDescriptionRequest()
	}
	response = NewUpdateRoleDescriptionResponse()
	err = c.Send(request, response)
	return
}

func NewListPoliciesByActionRequest() (request *ListPoliciesByActionRequest) {
	request = &ListPoliciesByActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListPoliciesByAction")
	return
}

func NewListPoliciesByActionResponse() (response *ListPoliciesByActionResponse) {
	response = &ListPoliciesByActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据接口名称查询策略列表
func (c *Client) ListPoliciesByAction(request *ListPoliciesByActionRequest) (response *ListPoliciesByActionResponse, err error) {
	if request == nil {
		request = NewListPoliciesByActionRequest()
	}
	response = NewListPoliciesByActionResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyUserAccessTokenRequest() (request *VerifyUserAccessTokenRequest) {
	request = &VerifyUserAccessTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "VerifyUserAccessToken")
	return
}

func NewVerifyUserAccessTokenResponse() (response *VerifyUserAccessTokenResponse) {
	response = &VerifyUserAccessTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证用户第三方开放平台access_token
func (c *Client) VerifyUserAccessToken(request *VerifyUserAccessTokenRequest) (response *VerifyUserAccessTokenResponse, err error) {
	if request == nil {
		request = NewVerifyUserAccessTokenRequest()
	}
	response = NewVerifyUserAccessTokenResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceLinkedRoleRequest() (request *CreateServiceLinkedRoleRequest) {
	request = &CreateServiceLinkedRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CreateServiceLinkedRole")
	return
}

func NewCreateServiceLinkedRoleResponse() (response *CreateServiceLinkedRoleResponse) {
	response = &CreateServiceLinkedRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建服务相关角色
func (c *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (response *CreateServiceLinkedRoleResponse, err error) {
	if request == nil {
		request = NewCreateServiceLinkedRoleRequest()
	}
	response = NewCreateServiceLinkedRoleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
	request = &UpdatePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdatePolicy")
	return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
	response = &UpdatePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（UpdatePolicy ）可用于更新策略。
// 如果已存在策略版本，本接口会直接更新策略的默认版本，不会创建新版本，如果不存在任何策略版本，则直接创建一个默认版本。
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyRequest()
	}
	response = NewUpdatePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewListSubAccountsRequest() (request *ListSubAccountsRequest) {
	request = &ListSubAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "ListSubAccounts")
	return
}

func NewListSubAccountsResponse() (response *ListSubAccountsResponse) {
	response = &ListSubAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取子帐号列表
func (c *Client) ListSubAccounts(request *ListSubAccountsRequest) (response *ListSubAccountsResponse, err error) {
	if request == nil {
		request = NewListSubAccountsRequest()
	}
	response = NewListSubAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
	request = &DescribeGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DescribeGroups")
	return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
	response = &DescribeGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量查询用户组详情
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeGroupsRequest()
	}
	response = NewDescribeGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDetachRolePoliciesRequest() (request *DetachRolePoliciesRequest) {
	request = &DetachRolePoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DetachRolePolicies")
	return
}

func NewDetachRolePoliciesResponse() (response *DetachRolePoliciesResponse) {
	response = &DetachRolePoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解除绑定多个策略到角色
func (c *Client) DetachRolePolicies(request *DetachRolePoliciesRequest) (response *DetachRolePoliciesResponse, err error) {
	if request == nil {
		request = NewDetachRolePoliciesRequest()
	}
	response = NewDetachRolePoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRolePolicyRequest() (request *DeleteRolePolicyRequest) {
	request = &DeleteRolePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DeleteRolePolicy")
	return
}

func NewDeleteRolePolicyResponse() (response *DeleteRolePolicyResponse) {
	response = &DeleteRolePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除角色策略
func (c *Client) DeleteRolePolicy(request *DeleteRolePolicyRequest) (response *DeleteRolePolicyResponse, err error) {
	if request == nil {
		request = NewDeleteRolePolicyRequest()
	}
	response = NewDeleteRolePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewSendEmailVerifyLinkRequest() (request *SendEmailVerifyLinkRequest) {
	request = &SendEmailVerifyLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "SendEmailVerifyLink")
	return
}

func NewSendEmailVerifyLinkResponse() (response *SendEmailVerifyLinkResponse) {
	response = &SendEmailVerifyLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送邮件修改链接
func (c *Client) SendEmailVerifyLink(request *SendEmailVerifyLinkRequest) (response *SendEmailVerifyLinkResponse, err error) {
	if request == nil {
		request = NewSendEmailVerifyLinkRequest()
	}
	response = NewSendEmailVerifyLinkResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRoleListRequest() (request *DescribeRoleListRequest) {
	request = &DescribeRoleListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DescribeRoleList")
	return
}

func NewDescribeRoleListResponse() (response *DescribeRoleListResponse) {
	response = &DescribeRoleListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeRoleList）用于获取账号下的角色列表。
func (c *Client) DescribeRoleList(request *DescribeRoleListRequest) (response *DescribeRoleListResponse, err error) {
	if request == nil {
		request = NewDescribeRoleListRequest()
	}
	response = NewDescribeRoleListResponse()
	err = c.Send(request, response)
	return
}

func NewAddReceiverRequest() (request *AddReceiverRequest) {
	request = &AddReceiverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "AddReceiver")
	return
}

func NewAddReceiverResponse() (response *AddReceiverResponse) {
	response = &AddReceiverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加消息接收者
func (c *Client) AddReceiver(request *AddReceiverRequest) (response *AddReceiverResponse, err error) {
	if request == nil {
		request = NewAddReceiverRequest()
	}
	response = NewAddReceiverResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
	request = &CreateRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "CreateRole")
	return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
	response = &CreateRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreateRole）用于创建角色。
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
	if request == nil {
		request = NewCreateRoleRequest()
	}
	response = NewCreateRoleResponse()
	err = c.Send(request, response)
	return
}

func NewGetLoginRulesRequest() (request *GetLoginRulesRequest) {
	request = &GetLoginRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetLoginRules")
	return
}

func NewGetLoginRulesResponse() (response *GetLoginRulesResponse) {
	response = &GetLoginRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取登录策略配置
func (c *Client) GetLoginRules(request *GetLoginRulesRequest) (response *GetLoginRulesResponse, err error) {
	if request == nil {
		request = NewGetLoginRulesRequest()
	}
	response = NewGetLoginRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDisableApiKeyRequest() (request *DisableApiKeyRequest) {
	request = &DisableApiKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "DisableApiKey")
	return
}

func NewDisableApiKeyResponse() (response *DisableApiKeyResponse) {
	response = &DisableApiKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用密钥
func (c *Client) DisableApiKey(request *DisableApiKeyRequest) (response *DisableApiKeyResponse, err error) {
	if request == nil {
		request = NewDisableApiKeyRequest()
	}
	response = NewDisableApiKeyResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
	request = &RemoveUserFromGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "RemoveUserFromGroup")
	return
}

func NewRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
	response = &RemoveUserFromGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从用户组删除用户
func (c *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
	if request == nil {
		request = NewRemoveUserFromGroupRequest()
	}
	response = NewRemoveUserFromGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetServiceListRequest() (request *GetServiceListRequest) {
	request = &GetServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetServiceList")
	return
}

func NewGetServiceListResponse() (response *GetServiceListResponse) {
	response = &GetServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetServiceList
func (c *Client) GetServiceList(request *GetServiceListRequest) (response *GetServiceListResponse, err error) {
	if request == nil {
		request = NewGetServiceListRequest()
	}
	response = NewGetServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserGroupListRequest() (request *GetUserGroupListRequest) {
	request = &GetUserGroupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "GetUserGroupList")
	return
}

func NewGetUserGroupListResponse() (response *GetUserGroupListResponse) {
	response = &GetUserGroupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户及用户组列表
func (c *Client) GetUserGroupList(request *GetUserGroupListRequest) (response *GetUserGroupListResponse, err error) {
	if request == nil {
		request = NewGetUserGroupListRequest()
	}
	response = NewGetUserGroupListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSubAccountTypeRequest() (request *UpdateSubAccountTypeRequest) {
	request = &UpdateSubAccountTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cam", APIVersion, "UpdateSubAccountType")
	return
}

func NewUpdateSubAccountTypeResponse() (response *UpdateSubAccountTypeResponse) {
	response = &UpdateSubAccountTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新子账号类型
func (c *Client) UpdateSubAccountType(request *UpdateSubAccountTypeRequest) (response *UpdateSubAccountTypeResponse, err error) {
	if request == nil {
		request = NewUpdateSubAccountTypeRequest()
	}
	response = NewUpdateSubAccountTypeResponse()
	err = c.Send(request, response)
	return
}

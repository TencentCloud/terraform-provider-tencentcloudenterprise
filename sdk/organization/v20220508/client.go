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

package v20220508

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2022-05-08"

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

func NewCheckMemberExistRequest() (request *CheckMemberExistRequest) {
	request = &CheckMemberExistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckMemberExist")
	return
}

func NewCheckMemberExistResponse() (response *CheckMemberExistResponse) {
	response = &CheckMemberExistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量检查成员账号是否存在，返回存在的账号列表
func (c *Client) CheckMemberExist(request *CheckMemberExistRequest) (response *CheckMemberExistResponse, err error) {
	if request == nil {
		request = NewCheckMemberExistRequest()
	}
	response = NewCheckMemberExistResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationRequest() (request *DescribeOrganizationRequest) {
	request = &DescribeOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganization")
	return
}

func NewDescribeOrganizationResponse() (response *DescribeOrganizationResponse) {
	response = &DescribeOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所属集团组织信息
func (c *Client) DescribeOrganization(request *DescribeOrganizationRequest) (response *DescribeOrganizationResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationRequest()
	}
	response = NewDescribeOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMemberRequest() (request *DescribeOrganizationMemberRequest) {
	request = &DescribeOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMember")
	return
}

func NewDescribeOrganizationMemberResponse() (response *DescribeOrganizationMemberResponse) {
	response = &DescribeOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员详情
func (c *Client) DescribeOrganizationMember(request *DescribeOrganizationMemberRequest) (response *DescribeOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMemberRequest()
	}
	response = NewDescribeOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewMoveOrganizationNodeRequest() (request *MoveOrganizationNodeRequest) {
	request = &MoveOrganizationNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "MoveOrganizationNode")
	return
}

func NewMoveOrganizationNodeResponse() (response *MoveOrganizationNodeResponse) {
	response = &MoveOrganizationNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 移动组织节点
func (c *Client) MoveOrganizationNode(request *MoveOrganizationNodeRequest) (response *MoveOrganizationNodeResponse, err error) {
	if request == nil {
		request = NewMoveOrganizationNodeRequest()
	}
	response = NewMoveOrganizationNodeResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationNodeRequest() (request *UpdateOrganizationNodeRequest) {
	request = &UpdateOrganizationNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationNode")
	return
}

func NewUpdateOrganizationNodeResponse() (response *UpdateOrganizationNodeResponse) {
	response = &UpdateOrganizationNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新组织节点
func (c *Client) UpdateOrganizationNode(request *UpdateOrganizationNodeRequest) (response *UpdateOrganizationNodeResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationNodeRequest()
	}
	response = NewUpdateOrganizationNodeResponse()
	err = c.Send(request, response)
	return
}

func NewBindOrganizationMemberAuthAccountRequest() (request *BindOrganizationMemberAuthAccountRequest) {
	request = &BindOrganizationMemberAuthAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationMemberAuthAccount")
	return
}

func NewBindOrganizationMemberAuthAccountResponse() (response *BindOrganizationMemberAuthAccountResponse) {
	response = &BindOrganizationMemberAuthAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定成员和子账号授权关系
func (c *Client) BindOrganizationMemberAuthAccount(request *BindOrganizationMemberAuthAccountRequest) (response *BindOrganizationMemberAuthAccountResponse, err error) {
	if request == nil {
		request = NewBindOrganizationMemberAuthAccountRequest()
	}
	response = NewBindOrganizationMemberAuthAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCancelOrganizationMemberAuthAccountRequest() (request *CancelOrganizationMemberAuthAccountRequest) {
	request = &CancelOrganizationMemberAuthAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationMemberAuthAccount")
	return
}

func NewCancelOrganizationMemberAuthAccountResponse() (response *CancelOrganizationMemberAuthAccountResponse) {
	response = &CancelOrganizationMemberAuthAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消成员和子账号的授权关系
func (c *Client) CancelOrganizationMemberAuthAccount(request *CancelOrganizationMemberAuthAccountRequest) (response *CancelOrganizationMemberAuthAccountResponse, err error) {
	if request == nil {
		request = NewCancelOrganizationMemberAuthAccountRequest()
	}
	response = NewCancelOrganizationMemberAuthAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationNodeMembersRequest() (request *DeleteOrganizationNodeMembersRequest) {
	request = &DeleteOrganizationNodeMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationNodeMembers")
	return
}

func NewDeleteOrganizationNodeMembersResponse() (response *DeleteOrganizationNodeMembersResponse) {
	response = &DeleteOrganizationNodeMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除节点成员
func (c *Client) DeleteOrganizationNodeMembers(request *DeleteOrganizationNodeMembersRequest) (response *DeleteOrganizationNodeMembersResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationNodeMembersRequest()
	}
	response = NewDeleteOrganizationNodeMembersResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrganizationNodeRequest() (request *AddOrganizationNodeRequest) {
	request = &AddOrganizationNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationNode")
	return
}

func NewAddOrganizationNodeResponse() (response *AddOrganizationNodeResponse) {
	response = &AddOrganizationNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加集团组织节点
func (c *Client) AddOrganizationNode(request *AddOrganizationNodeRequest) (response *AddOrganizationNodeResponse, err error) {
	if request == nil {
		request = NewAddOrganizationNodeRequest()
	}
	response = NewAddOrganizationNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationNodesRequest() (request *DeleteOrganizationNodesRequest) {
	request = &DeleteOrganizationNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationNodes")
	return
}

func NewDeleteOrganizationNodesResponse() (response *DeleteOrganizationNodesResponse) {
	response = &DeleteOrganizationNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除组织节点
func (c *Client) DeleteOrganizationNodes(request *DeleteOrganizationNodesRequest) (response *DeleteOrganizationNodesResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationNodesRequest()
	}
	response = NewDeleteOrganizationNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMemberAuthAccountsRequest() (request *DescribeOrganizationMemberAuthAccountsRequest) {
	request = &DescribeOrganizationMemberAuthAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthAccounts")
	return
}

func NewDescribeOrganizationMemberAuthAccountsResponse() (response *DescribeOrganizationMemberAuthAccountsResponse) {
	response = &DescribeOrganizationMemberAuthAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询成员授权子账号列表
func (c *Client) DescribeOrganizationMemberAuthAccounts(request *DescribeOrganizationMemberAuthAccountsRequest) (response *DescribeOrganizationMemberAuthAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMemberAuthAccountsRequest()
	}
	response = NewDescribeOrganizationMemberAuthAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMembersRequest() (request *DescribeOrganizationMembersRequest) {
	request = &DescribeOrganizationMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembers")
	return
}

func NewDescribeOrganizationMembersResponse() (response *DescribeOrganizationMembersResponse) {
	response = &DescribeOrganizationMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 成员账号列表
func (c *Client) DescribeOrganizationMembers(request *DescribeOrganizationMembersRequest) (response *DescribeOrganizationMembersResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMembersRequest()
	}
	response = NewDescribeOrganizationMembersResponse()
	err = c.Send(request, response)
	return
}

func NewMoveOrganizationNodeMembersRequest() (request *MoveOrganizationNodeMembersRequest) {
	request = &MoveOrganizationNodeMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "MoveOrganizationNodeMembers")
	return
}

func NewMoveOrganizationNodeMembersResponse() (response *MoveOrganizationNodeMembersResponse) {
	response = &MoveOrganizationNodeMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 移动节点成员
func (c *Client) MoveOrganizationNodeMembers(request *MoveOrganizationNodeMembersRequest) (response *MoveOrganizationNodeMembersResponse, err error) {
	if request == nil {
		request = NewMoveOrganizationNodeMembersRequest()
	}
	response = NewMoveOrganizationNodeMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationMemberPolicyRequest() (request *CreateOrganizationMemberPolicyRequest) {
	request = &CreateOrganizationMemberPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMemberPolicy")
	return
}

func NewCreateOrganizationMemberPolicyResponse() (response *CreateOrganizationMemberPolicyResponse) {
	response = &CreateOrganizationMemberPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加成员授权策略
func (c *Client) CreateOrganizationMemberPolicy(request *CreateOrganizationMemberPolicyRequest) (response *CreateOrganizationMemberPolicyResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationMemberPolicyRequest()
	}
	response = NewCreateOrganizationMemberPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodesByParentRequest() (request *DescribeOrganizationNodesByParentRequest) {
	request = &DescribeOrganizationNodesByParentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodesByParent")
	return
}

func NewDescribeOrganizationNodesByParentResponse() (response *DescribeOrganizationNodesByParentResponse) {
	response = &DescribeOrganizationNodesByParentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据父节点查询子节点
func (c *Client) DescribeOrganizationNodesByParent(request *DescribeOrganizationNodesByParentRequest) (response *DescribeOrganizationNodesByParentResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodesByParentRequest()
	}
	response = NewDescribeOrganizationNodesByParentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodeRequest() (request *DescribeOrganizationNodeRequest) {
	request = &DescribeOrganizationNodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNode")
	return
}

func NewDescribeOrganizationNodeResponse() (response *DescribeOrganizationNodeResponse) {
	response = &DescribeOrganizationNodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单节点信息
func (c *Client) DescribeOrganizationNode(request *DescribeOrganizationNodeRequest) (response *DescribeOrganizationNodeResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodeRequest()
	}
	response = NewDescribeOrganizationNodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodesRequest() (request *DescribeOrganizationNodesRequest) {
	request = &DescribeOrganizationNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodes")
	return
}

func NewDescribeOrganizationNodesResponse() (response *DescribeOrganizationNodesResponse) {
	response = &DescribeOrganizationNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织节点列表
func (c *Client) DescribeOrganizationNodes(request *DescribeOrganizationNodesRequest) (response *DescribeOrganizationNodesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodesRequest()
	}
	response = NewDescribeOrganizationNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodeMembersRequest() (request *DescribeOrganizationNodeMembersRequest) {
	request = &DescribeOrganizationNodeMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodeMembers")
	return
}

func NewDescribeOrganizationNodeMembersResponse() (response *DescribeOrganizationNodeMembersResponse) {
	response = &DescribeOrganizationNodeMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取节点成员列表
func (c *Client) DescribeOrganizationNodeMembers(request *DescribeOrganizationNodeMembersRequest) (response *DescribeOrganizationNodeMembersResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodeMembersRequest()
	}
	response = NewDescribeOrganizationNodeMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCancelOrganizationMemberAuthAccountForDeletionSubAccountRequest() (request *CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest) {
	request = &CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationMemberAuthAccountForDeletionSubAccount")
	return
}

func NewCancelOrganizationMemberAuthAccountForDeletionSubAccountResponse() (response *CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse) {
	response = &CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消成员和子账号的授权关系(删除子账号消息调用)
func (c *Client) CancelOrganizationMemberAuthAccountForDeletionSubAccount(request *CancelOrganizationMemberAuthAccountForDeletionSubAccountRequest) (response *CancelOrganizationMemberAuthAccountForDeletionSubAccountResponse, err error) {
	if request == nil {
		request = NewCancelOrganizationMemberAuthAccountForDeletionSubAccountRequest()
	}
	response = NewCancelOrganizationMemberAuthAccountForDeletionSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationRequest() (request *CreateOrganizationRequest) {
	request = &CreateOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganization")

	return
}

func NewCreateOrganizationResponse() (response *CreateOrganizationResponse) {
	response = &CreateOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateOrganization
// 创建企业组织
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//	FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//	FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	UNSUPPORTEDOPERATION = "UnsupportedOperation"
//	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWCREATEORGANIZATION = "UnsupportedOperation.CreateMemberNotAllowCreateOrganization"
func (c *Client) CreateOrganization(request *CreateOrganizationRequest) (response *CreateOrganizationResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationRequest()
	}

	response = NewCreateOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationRequest() (request *DeleteOrganizationRequest) {
	request = &DeleteOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganization")

	return
}

func NewDeleteOrganizationResponse() (response *DeleteOrganizationResponse) {
	response = &DeleteOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteOrganization
// 删除企业组织
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//	FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//	FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"
//	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//	FAILEDOPERATION_QUITSHAREUINT = "FailedOperation.QuitShareUint"
//	FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//	FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganization(request *DeleteOrganizationRequest) (response *DeleteOrganizationResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationRequest()
	}

	response = NewDeleteOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationMembersRequest() (request *DeleteOrganizationMembersRequest) {
	request = &DeleteOrganizationMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMembers")

	return
}

func NewDeleteOrganizationMembersResponse() (response *DeleteOrganizationMembersResponse) {
	response = &DeleteOrganizationMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteOrganizationMembers
// 从组织中移除成员账号，不会删除账号。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//	FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//	FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//	FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//	FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//	FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//	FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//	FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"
//	UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//	UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//	UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
func (c *Client) DeleteOrganizationMembers(request *DeleteOrganizationMembersRequest) (response *DeleteOrganizationMembersResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationMembersRequest()
	}

	response = NewDeleteOrganizationMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationIdentityRequest() (request *CreateOrganizationIdentityRequest) {
	request = &CreateOrganizationIdentityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationIdentity")

	return
}

func NewCreateOrganizationIdentityResponse() (response *CreateOrganizationIdentityResponse) {
	response = &CreateOrganizationIdentityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateOrganizationIdentity
// 添加组织身份
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//	FAILEDOPERATION_ORGANIZATIONIDENTITYNAMEUSED = "FailedOperation.OrganizationIdentityNameUsed"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	LIMITEXCEEDED_IDENTITYEXCEEDLIMIT = "LimitExceeded.IdentityExceedLimit"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) CreateOrganizationIdentity(request *CreateOrganizationIdentityRequest) (response *CreateOrganizationIdentityResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationIdentityRequest()
	}

	response = NewCreateOrganizationIdentityResponse()
	err = c.Send(request, response)
	return
}

func NewListOrganizationIdentityRequest() (request *ListOrganizationIdentityRequest) {
	request = &ListOrganizationIdentityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationIdentity")

	return
}

func NewListOrganizationIdentityResponse() (response *ListOrganizationIdentityResponse) {
	response = &ListOrganizationIdentityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// ListOrganizationIdentity
// 获取组织成员访问身份列表
//
// 可能返回的错误码:
//
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) ListOrganizationIdentity(request *ListOrganizationIdentityRequest) (response *ListOrganizationIdentityResponse, err error) {
	if request == nil {
		request = NewListOrganizationIdentityRequest()
	}

	response = NewListOrganizationIdentityResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationIdentityRequest() (request *UpdateOrganizationIdentityRequest) {
	request = &UpdateOrganizationIdentityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationIdentity")

	return
}

func NewUpdateOrganizationIdentityResponse() (response *UpdateOrganizationIdentityResponse) {
	response = &UpdateOrganizationIdentityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdateOrganizationIdentity
// 更新组织身份
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) UpdateOrganizationIdentity(request *UpdateOrganizationIdentityRequest) (response *UpdateOrganizationIdentityResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationIdentityRequest()
	}

	response = NewUpdateOrganizationIdentityResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationIdentityRequest() (request *DeleteOrganizationIdentityRequest) {
	request = &DeleteOrganizationIdentityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationIdentity")

	return
}

func NewDeleteOrganizationIdentityResponse() (response *DeleteOrganizationIdentityResponse) {
	response = &DeleteOrganizationIdentityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteOrganizationIdentity
// 删除组织身份
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ORGANIZATIONIDENTITYINUSED = "FailedOperation.OrganizationIdentityInUsed"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationIdentity(request *DeleteOrganizationIdentityRequest) (response *DeleteOrganizationIdentityResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationIdentityRequest()
	}

	response = NewDeleteOrganizationIdentityResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePolicyRequest() (request *CreatePolicyRequest) {
	request = &CreatePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "CreatePolicy")

	return
}

func NewCreatePolicyResponse() (response *CreatePolicyResponse) {
	response = &CreatePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreatePolicy
// 创建一个特殊类型的策略，您可以关联到企业组织Root节点、企业部门节点或者企业的成员账号。
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//	FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//	FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//	INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//	INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//	INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//	INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//	INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//	INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//	INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//	INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//	INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//	INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"
//	INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//	INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"
//	INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//	INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//	INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//	INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//	INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//	INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//	INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//	INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//	INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//	INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//	INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//	INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//	INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//	INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//	INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//	INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"
//	LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"
//	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) CreatePolicy(request *CreatePolicyRequest) (response *CreatePolicyResponse, err error) {
	if request == nil {
		request = NewCreatePolicyRequest()
	}

	response = NewCreatePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewListPoliciesRequest() (request *ListPoliciesRequest) {
	request = &ListPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "ListPolicies")

	return
}

func NewListPoliciesResponse() (response *ListPoliciesResponse) {
	response = &ListPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// ListPolicies
// 本接口（ListPolicies）可用于查询查看策略列表数据
//
// 可能返回的错误码:
//
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//	INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//	MISSINGPARAMETER = "MissingParameter"
//	RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//	RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//	RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//	RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//	RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//	RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//	RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//	RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//	RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//	RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//	RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//	RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//	RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//	RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//	RESOURCEUNAVAILABLE = "ResourceUnavailable"
func (c *Client) ListPolicies(request *ListPoliciesRequest) (response *ListPoliciesResponse, err error) {
	if request == nil {
		request = NewListPoliciesRequest()
	}

	response = NewListPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePolicyRequest() (request *DescribePolicyRequest) {
	request = &DescribePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DescribePolicy")

	return
}

func NewDescribePolicyResponse() (response *DescribePolicyResponse) {
	response = &DescribePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DescribePolicy
// 本接口（DescribePolicy）可用于查询查看策略详情。
//
// 可能返回的错误码:
//
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//	INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//	MISSINGPARAMETER = "MissingParameter"
//	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
func (c *Client) DescribePolicy(request *DescribePolicyRequest) (response *DescribePolicyResponse, err error) {
	if request == nil {
		request = NewDescribePolicyRequest()
	}

	response = NewDescribePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyRequest() (request *UpdatePolicyRequest) {
	request = &UpdatePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "UpdatePolicy")

	return
}

func NewUpdatePolicyResponse() (response *UpdatePolicyResponse) {
	response = &UpdatePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdatePolicy
// 编辑策略
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//	FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//	INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//	INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//	INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//	INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//	INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//	INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//	INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//	INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//	INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//	INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//	INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//	INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"
//	INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//	INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"
//	INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//	INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//	INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//	INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//	INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//	INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//	INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//	INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//	INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//	INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//	INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//	INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//	INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//	INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//	INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//	INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"
//	LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"
//	RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) UpdatePolicy(request *UpdatePolicyRequest) (response *UpdatePolicyResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyRequest()
	}

	response = NewUpdatePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePolicyRequest() (request *DeletePolicyRequest) {
	request = &DeletePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DeletePolicy")

	return
}

func NewDeletePolicyResponse() (response *DeletePolicyResponse) {
	response = &DeletePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeletePolicy
// 删除策略
//
// 可能返回的错误码:
//
//	FAILEDOPERATION = "FailedOperation"
//	FAILEDOPERATION_ORGANIZATIONPOLICYINUSED = "FailedOperation.OrganizationPolicyInUsed"
//	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//	RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DeletePolicy(request *DeletePolicyRequest) (response *DeletePolicyResponse, err error) {
	if request == nil {
		request = NewDeletePolicyRequest()
	}

	response = NewDeletePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewEnablePolicyTypeRequest() (request *EnablePolicyTypeRequest) {
	request = &EnablePolicyTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "EnablePolicyType")

	return
}

func NewEnablePolicyTypeResponse() (response *EnablePolicyTypeResponse) {
	response = &EnablePolicyTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// EnablePolicyType
// 启用策略类型
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//	FAILEDOPERATION_POLICYENABLEINVALID = "FailedOperation.PolicyEnableInvalid"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) EnablePolicyType(request *EnablePolicyTypeRequest) (response *EnablePolicyTypeResponse, err error) {
	if request == nil {
		request = NewEnablePolicyTypeRequest()
	}

	response = NewEnablePolicyTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePolicyConfigRequest() (request *DescribePolicyConfigRequest) {
	request = &DescribePolicyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DescribePolicyConfig")

	return
}

func NewDescribePolicyConfigResponse() (response *DescribePolicyConfigResponse) {
	response = &DescribePolicyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DescribePolicyConfig
// 本接口（DescribePolicyConfig）可用于查询企业组织策略配置
//
// 可能返回的错误码:
//
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//	INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//	MISSINGPARAMETER = "MissingParameter"
//	RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//	RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//	RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//	RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//	RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//	RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//	RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//	RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//	RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//	RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//	RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//	RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//	RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//	RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribePolicyConfig(request *DescribePolicyConfigRequest) (response *DescribePolicyConfigResponse, err error) {
	if request == nil {
		request = NewDescribePolicyConfigRequest()
	}

	response = NewDescribePolicyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDisablePolicyTypeRequest() (request *DisablePolicyTypeRequest) {
	request = &DisablePolicyTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DisablePolicyType")

	return
}

func NewDisablePolicyTypeResponse() (response *DisablePolicyTypeResponse) {
	response = &DisablePolicyTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DisablePolicyType
// 禁用策略类型
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DisablePolicyType(request *DisablePolicyTypeRequest) (response *DisablePolicyTypeResponse, err error) {
	if request == nil {
		request = NewDisablePolicyTypeRequest()
	}

	response = NewDisablePolicyTypeResponse()
	err = c.Send(request, response)
	return
}

func NewAttachPolicyRequest() (request *AttachPolicyRequest) {
	request = &AttachPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "AttachPolicy")

	return
}

func NewAttachPolicyResponse() (response *AttachPolicyResponse) {
	response = &AttachPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// AttachPolicy
// 绑定策略
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) AttachPolicy(request *AttachPolicyRequest) (response *AttachPolicyResponse, err error) {
	if request == nil {
		request = NewAttachPolicyRequest()
	}

	response = NewAttachPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewListTargetsForPolicyRequest() (request *ListTargetsForPolicyRequest) {
	request = &ListTargetsForPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "ListTargetsForPolicy")

	return
}

func NewListTargetsForPolicyResponse() (response *ListTargetsForPolicyResponse) {
	response = &ListTargetsForPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// ListTargetsForPolicy
// 本接口（ListTargetsForPolicy）查询某个指定策略关联的目标列表
//
// 可能返回的错误码:
//
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//	INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//	MISSINGPARAMETER = "MissingParameter"
//	RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//	RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//	RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//	RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//	RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//	RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//	RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//	RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//	RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//	RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//	RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//	RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//	RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//	RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//	RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
func (c *Client) ListTargetsForPolicy(request *ListTargetsForPolicyRequest) (response *ListTargetsForPolicyResponse, err error) {
	if request == nil {
		request = NewListTargetsForPolicyRequest()
	}

	response = NewListTargetsForPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDetachPolicyRequest() (request *DetachPolicyRequest) {
	request = &DetachPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DetachPolicy")

	return
}

func NewDetachPolicyResponse() (response *DetachPolicyResponse) {
	response = &DetachPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DetachPolicy
// 解绑策略
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ORGANIZATIONDETACHLASTPOLICYERROR = "FailedOperation.OrganizationDetachLastPolicyError"
//	FAILEDOPERATION_ORGANIZATIONDETACHPOLICYERROR = "FailedOperation.OrganizationDetachPolicyError"
//	FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//	RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
func (c *Client) DetachPolicy(request *DetachPolicyRequest) (response *DetachPolicyResponse, err error) {
	if request == nil {
		request = NewDetachPolicyRequest()
	}

	response = NewDetachPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationMemberRequest() (request *CreateOrganizationMemberRequest) {
	request = &CreateOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMember")

	return
}

func NewCreateOrganizationMemberResponse() (response *CreateOrganizationMemberResponse) {
	response = &CreateOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateOrganizationMember
// 创建组织成员
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//	FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//	FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//	FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//	FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//	FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//	FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//	FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//	FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"
//	FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//	FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//	FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//	FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//	FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//	FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//	FAILEDOPERATION_PARTNERMANAGEMENTERR = "FailedOperation.PartnerManagementErr"
//	FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//	FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//	INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
//	LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//	LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//	RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//	UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//	UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//	UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//	UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//	UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//	UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//	UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//	UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//	UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//	UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//	UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//	UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//	UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//	UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//	UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
func (c *Client) CreateOrganizationMember(request *CreateOrganizationMemberRequest) (response *CreateOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationMemberRequest()
	}

	response = NewCreateOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationMemberAuthIdentityRequest() (request *CreateOrganizationMemberAuthIdentityRequest) {
	request = &CreateOrganizationMemberAuthIdentityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMemberAuthIdentity")

	return
}

func NewCreateOrganizationMemberAuthIdentityResponse() (response *CreateOrganizationMemberAuthIdentityResponse) {
	response = &CreateOrganizationMemberAuthIdentityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateOrganizationMemberAuthIdentity
// 添加组织成员访问授权
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//	FAILEDOPERATION_ORGANIZATIONIDENTITYPOLICYERROR = "FailedOperation.OrganizationIdentityPolicyError"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMemberAuthIdentity(request *CreateOrganizationMemberAuthIdentityRequest) (response *CreateOrganizationMemberAuthIdentityResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationMemberAuthIdentityRequest()
	}

	response = NewCreateOrganizationMemberAuthIdentityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMemberAuthIdentitiesRequest() (request *DescribeOrganizationMemberAuthIdentitiesRequest) {
	request = &DescribeOrganizationMemberAuthIdentitiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberAuthIdentities")

	return
}

func NewDescribeOrganizationMemberAuthIdentitiesResponse() (response *DescribeOrganizationMemberAuthIdentitiesResponse) {
	response = &DescribeOrganizationMemberAuthIdentitiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DescribeOrganizationMemberAuthIdentities
// 获取组织成员访问授权列表
//
// 可能返回的错误码:
//
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DescribeOrganizationMemberAuthIdentities(request *DescribeOrganizationMemberAuthIdentitiesRequest) (response *DescribeOrganizationMemberAuthIdentitiesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMemberAuthIdentitiesRequest()
	}

	response = NewDescribeOrganizationMemberAuthIdentitiesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationMemberAuthIdentityRequest() (request *DeleteOrganizationMemberAuthIdentityRequest) {
	request = &DeleteOrganizationMemberAuthIdentityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMemberAuthIdentity")

	return
}

func NewDeleteOrganizationMemberAuthIdentityResponse() (response *DeleteOrganizationMemberAuthIdentityResponse) {
	response = &DeleteOrganizationMemberAuthIdentityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteOrganizationMemberAuthIdentity
// 删除组织成员访问授权
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_MEMBERIDENTITYAUTHUSED = "FailedOperation.MemberIdentityAuthUsed"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMemberAuthIdentity(request *DeleteOrganizationMemberAuthIdentityRequest) (response *DeleteOrganizationMemberAuthIdentityResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationMemberAuthIdentityRequest()
	}

	response = NewDeleteOrganizationMemberAuthIdentityResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrganizationMemberEmailRequest() (request *AddOrganizationMemberEmailRequest) {
	request = &AddOrganizationMemberEmailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationMemberEmail")

	return
}

func NewAddOrganizationMemberEmailResponse() (response *AddOrganizationMemberEmailResponse) {
	response = &AddOrganizationMemberEmailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// AddOrganizationMemberEmail
// 添加组织成员邮箱
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_CHECKACCOUNTPHONEBINDLIMIT = "FailedOperation.CheckAccountPhoneBindLimit"
//	FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//	FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//	FAILEDOPERATION_MEMBEREMAILEXIST = "FailedOperation.MemberEmailExist"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	LIMITEXCEEDED_EMAILBINDOVERLIMIT = "LimitExceeded.EmailBindOverLimit"
//	LIMITEXCEEDED_PHONENUMBOUND = "LimitExceeded.PhoneNumBound"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) AddOrganizationMemberEmail(request *AddOrganizationMemberEmailRequest) (response *AddOrganizationMemberEmailResponse, err error) {
	if request == nil {
		request = NewAddOrganizationMemberEmailRequest()
	}

	response = NewAddOrganizationMemberEmailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMemberEmailBindRequest() (request *DescribeOrganizationMemberEmailBindRequest) {
	request = &DescribeOrganizationMemberEmailBindRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberEmailBind")

	return
}

func NewDescribeOrganizationMemberEmailBindResponse() (response *DescribeOrganizationMemberEmailBindResponse) {
	response = &DescribeOrganizationMemberEmailBindResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DescribeOrganizationMemberEmailBind
// 查询成员邮箱绑定详细信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ACCOUNTALREADYREGISTER = "FailedOperation.AccountAlreadyRegister"
//	FAILEDOPERATION_BINDEMAILLINKEXPIRED = "FailedOperation.BindEmailLinkExpired"
//	FAILEDOPERATION_BINDEMAILLINKINVALID = "FailedOperation.BindEmailLinkInvalid"
//	FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//	FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//	FAILEDOPERATION_MEMBERBINDEMAILERROR = "FailedOperation.MemberBindEmailError"
//	FAILEDOPERATION_MEMBERBINDPHONEERROR = "FailedOperation.MemberBindPhoneError"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	INVALIDPARAMETER_CODEERROR = "InvalidParameter.CodeError"
//	INVALIDPARAMETER_CODEEXPIRED = "InvalidParameter.CodeExpired"
//	INVALIDPARAMETER_INVALIDEMAIL = "InvalidParameter.InvalidEmail"
//	INVALIDPARAMETER_PASSWORDILLEGAL = "InvalidParameter.PasswordIllegal"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DescribeOrganizationMemberEmailBind(request *DescribeOrganizationMemberEmailBindRequest) (response *DescribeOrganizationMemberEmailBindResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMemberEmailBindRequest()
	}

	response = NewDescribeOrganizationMemberEmailBindResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationMemberEmailBindRequest() (request *UpdateOrganizationMemberEmailBindRequest) {
	request = &UpdateOrganizationMemberEmailBindRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMemberEmailBind")

	return
}

func NewUpdateOrganizationMemberEmailBindResponse() (response *UpdateOrganizationMemberEmailBindResponse) {
	response = &UpdateOrganizationMemberEmailBindResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdateOrganizationMemberEmailBind
// 修改绑定成员邮箱
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//	FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//	FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	LIMITEXCEEDED_UPDATEEMAILBINDOVERLIMIT = "LimitExceeded.UpdateEmailBindOverLimit"
//	RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) UpdateOrganizationMemberEmailBind(request *UpdateOrganizationMemberEmailBindRequest) (response *UpdateOrganizationMemberEmailBindResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationMemberEmailBindRequest()
	}

	response = NewUpdateOrganizationMemberEmailBindResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationMembersPolicyRequest() (request *CreateOrganizationMembersPolicyRequest) {
	request = &CreateOrganizationMembersPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMembersPolicy")

	return
}

func NewCreateOrganizationMembersPolicyResponse() (response *CreateOrganizationMembersPolicyResponse) {
	response = &CreateOrganizationMembersPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateOrganizationMembersPolicy
// 创建组织成员访问策略
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//	FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	UNSUPPORTEDOPERATION = "UnsupportedOperation"
func (c *Client) CreateOrganizationMembersPolicy(request *CreateOrganizationMembersPolicyRequest) (response *CreateOrganizationMembersPolicyResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationMembersPolicyRequest()
	}

	response = NewCreateOrganizationMembersPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationMembersPolicyRequest() (request *DeleteOrganizationMembersPolicyRequest) {
	request = &DeleteOrganizationMembersPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationMembersPolicy")

	return
}

func NewDeleteOrganizationMembersPolicyResponse() (response *DeleteOrganizationMembersPolicyResponse) {
	response = &DeleteOrganizationMembersPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteOrganizationMembersPolicy
// 删除组织成员访问策略
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DELETEPOLICY = "FailedOperation.DeletePolicy"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
func (c *Client) DeleteOrganizationMembersPolicy(request *DeleteOrganizationMembersPolicyRequest) (response *DeleteOrganizationMembersPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationMembersPolicyRequest()
	}

	response = NewDeleteOrganizationMembersPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewQuitOrganizationRequest() (request *QuitOrganizationRequest) {
	request = &QuitOrganizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "QuitOrganization")

	return
}

func NewQuitOrganizationResponse() (response *QuitOrganizationResponse) {
	response = &QuitOrganizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// QuitOrganization
// 退出企业组织
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//	FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//	FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//	FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//	FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//	FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//	FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//	FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETER = "InvalidParameter"
//	RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//	RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//	UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWQUIT = "UnsupportedOperation.CreateMemberNotAllowQuit"
//	UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//	UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//	UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
//	UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
func (c *Client) QuitOrganization(request *QuitOrganizationRequest) (response *QuitOrganizationResponse, err error) {
	if request == nil {
		request = NewQuitOrganizationRequest()
	}

	response = NewQuitOrganizationResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationMemberRequest() (request *UpdateOrganizationMemberRequest) {
	request = &UpdateOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMember")
	return
}

func NewUpdateOrganizationMemberResponse() (response *UpdateOrganizationMemberResponse) {
	response = &UpdateOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新组织成员信息
func (c *Client) UpdateOrganizationMember(request *UpdateOrganizationMemberRequest) (response *UpdateOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationMemberRequest()
	}
	response = NewUpdateOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewAcceptMemberChangePermissionRequest() (request *AcceptMemberChangePermissionRequest) {
	request = &AcceptMemberChangePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "AcceptMemberChangePermission")
	return
}

func NewAcceptMemberChangePermissionResponse() (response *AcceptMemberChangePermissionResponse) {
	response = &AcceptMemberChangePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 确认成员权限变更
func (c *Client) AcceptMemberChangePermission(request *AcceptMemberChangePermissionRequest) (response *AcceptMemberChangePermissionResponse, err error) {
	if request == nil {
		request = NewAcceptMemberChangePermissionRequest()
	}
	response = NewAcceptMemberChangePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewAcceptOrganizationInvitationRequest() (request *AcceptOrganizationInvitationRequest) {
	request = &AcceptOrganizationInvitationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "AcceptOrganizationInvitation")
	return
}

func NewAcceptOrganizationInvitationResponse() (response *AcceptOrganizationInvitationResponse) {
	response = &AcceptOrganizationInvitationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接受加入企业组织邀请
func (c *Client) AcceptOrganizationInvitation(request *AcceptOrganizationInvitationRequest) (response *AcceptOrganizationInvitationResponse, err error) {
	if request == nil {
		request = NewAcceptOrganizationInvitationRequest()
	}
	response = NewAcceptOrganizationInvitationResponse()
	err = c.Send(request, response)
	return
}

func NewAddOrganizationNodeTagsRequest() (request *AddOrganizationNodeTagsRequest) {
	request = &AddOrganizationNodeTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "AddOrganizationNodeTags")
	return
}

func NewAddOrganizationNodeTagsResponse() (response *AddOrganizationNodeTagsResponse) {
	response = &AddOrganizationNodeTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 给组织节点打标签
func (c *Client) AddOrganizationNodeTags(request *AddOrganizationNodeTagsRequest) (response *AddOrganizationNodeTagsResponse, err error) {
	if request == nil {
		request = NewAddOrganizationNodeTagsRequest()
	}
	response = NewAddOrganizationNodeTagsResponse()
	err = c.Send(request, response)
	return
}

func NewBindOrganizationMemberEmailRequest() (request *BindOrganizationMemberEmailRequest) {
	request = &BindOrganizationMemberEmailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationMemberEmail")
	return
}

func NewBindOrganizationMemberEmailResponse() (response *BindOrganizationMemberEmailResponse) {
	response = &BindOrganizationMemberEmailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 成员绑定邮箱和安全手机
func (c *Client) BindOrganizationMemberEmail(request *BindOrganizationMemberEmailRequest) (response *BindOrganizationMemberEmailResponse, err error) {
	if request == nil {
		request = NewBindOrganizationMemberEmailRequest()
	}
	response = NewBindOrganizationMemberEmailResponse()
	err = c.Send(request, response)
	return
}

func NewBindOrganizationPolicyGroupRequest() (request *BindOrganizationPolicyGroupRequest) {
	request = &BindOrganizationPolicyGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationPolicyGroup")
	return
}

func NewBindOrganizationPolicyGroupResponse() (response *BindOrganizationPolicyGroupResponse) {
	response = &BindOrganizationPolicyGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定成员访问授权策略和用户组
func (c *Client) BindOrganizationPolicyGroup(request *BindOrganizationPolicyGroupRequest) (response *BindOrganizationPolicyGroupResponse, err error) {
	if request == nil {
		request = NewBindOrganizationPolicyGroupRequest()
	}
	response = NewBindOrganizationPolicyGroupResponse()
	err = c.Send(request, response)
	return
}

func NewBindOrganizationPolicySubAccountRequest() (request *BindOrganizationPolicySubAccountRequest) {
	request = &BindOrganizationPolicySubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "BindOrganizationPolicySubAccount")
	return
}

func NewBindOrganizationPolicySubAccountResponse() (response *BindOrganizationPolicySubAccountResponse) {
	response = &BindOrganizationPolicySubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定成员访问授权策略和组织管理员子账号
func (c *Client) BindOrganizationPolicySubAccount(request *BindOrganizationPolicySubAccountRequest) (response *BindOrganizationPolicySubAccountResponse, err error) {
	if request == nil {
		request = NewBindOrganizationPolicySubAccountRequest()
	}
	response = NewBindOrganizationPolicySubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCancelMemberChangePermissionRequest() (request *CancelMemberChangePermissionRequest) {
	request = &CancelMemberChangePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CancelMemberChangePermission")
	return
}

func NewCancelMemberChangePermissionResponse() (response *CancelMemberChangePermissionResponse) {
	response = &CancelMemberChangePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消成员权限变更
func (c *Client) CancelMemberChangePermission(request *CancelMemberChangePermissionRequest) (response *CancelMemberChangePermissionResponse, err error) {
	if request == nil {
		request = NewCancelMemberChangePermissionRequest()
	}
	response = NewCancelMemberChangePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewCancelOrganizationInvitationRequest() (request *CancelOrganizationInvitationRequest) {
	request = &CancelOrganizationInvitationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationInvitation")
	return
}

func NewCancelOrganizationInvitationResponse() (response *CancelOrganizationInvitationResponse) {
	response = &CancelOrganizationInvitationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消企业组织邀请
func (c *Client) CancelOrganizationInvitation(request *CancelOrganizationInvitationRequest) (response *CancelOrganizationInvitationResponse, err error) {
	if request == nil {
		request = NewCancelOrganizationInvitationRequest()
	}
	response = NewCancelOrganizationInvitationResponse()
	err = c.Send(request, response)
	return
}

func NewCancelOrganizationPolicyGroupRequest() (request *CancelOrganizationPolicyGroupRequest) {
	request = &CancelOrganizationPolicyGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationPolicyGroup")
	return
}

func NewCancelOrganizationPolicyGroupResponse() (response *CancelOrganizationPolicyGroupResponse) {
	response = &CancelOrganizationPolicyGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解绑成员访问授权策略和用户组
func (c *Client) CancelOrganizationPolicyGroup(request *CancelOrganizationPolicyGroupRequest) (response *CancelOrganizationPolicyGroupResponse, err error) {
	if request == nil {
		request = NewCancelOrganizationPolicyGroupRequest()
	}
	response = NewCancelOrganizationPolicyGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCancelOrganizationPolicySubAccountRequest() (request *CancelOrganizationPolicySubAccountRequest) {
	request = &CancelOrganizationPolicySubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CancelOrganizationPolicySubAccount")
	return
}

func NewCancelOrganizationPolicySubAccountResponse() (response *CancelOrganizationPolicySubAccountResponse) {
	response = &CancelOrganizationPolicySubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解绑成员访问授权策略和组织管理员子账号
func (c *Client) CancelOrganizationPolicySubAccount(request *CancelOrganizationPolicySubAccountRequest) (response *CancelOrganizationPolicySubAccountResponse, err error) {
	if request == nil {
		request = NewCancelOrganizationPolicySubAccountRequest()
	}
	response = NewCancelOrganizationPolicySubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAccountIsSubClientRequest() (request *CheckAccountIsSubClientRequest) {
	request = &CheckAccountIsSubClientRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckAccountIsSubClient")
	return
}

func NewCheckAccountIsSubClientResponse() (response *CheckAccountIsSubClientResponse) {
	response = &CheckAccountIsSubClientResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查账号是否经销子客
func (c *Client) CheckAccountIsSubClient(request *CheckAccountIsSubClientRequest) (response *CheckAccountIsSubClientResponse, err error) {
	if request == nil {
		request = NewCheckAccountIsSubClientRequest()
	}
	response = NewCheckAccountIsSubClientResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAccountStatusRequest() (request *CheckAccountStatusRequest) {
	request = &CheckAccountStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckAccountStatus")
	return
}

func NewCheckAccountStatusResponse() (response *CheckAccountStatusResponse) {
	response = &CheckAccountStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查账号状态
func (c *Client) CheckAccountStatus(request *CheckAccountStatusRequest) (response *CheckAccountStatusResponse, err error) {
	if request == nil {
		request = NewCheckAccountStatusRequest()
	}
	response = NewCheckAccountStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAddOrganizationMemberRequest() (request *CheckAddOrganizationMemberRequest) {
	request = &CheckAddOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckAddOrganizationMember")
	return
}

func NewCheckAddOrganizationMemberResponse() (response *CheckAddOrganizationMemberResponse) {
	response = &CheckAddOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查添加组织成员
func (c *Client) CheckAddOrganizationMember(request *CheckAddOrganizationMemberRequest) (response *CheckAddOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewCheckAddOrganizationMemberRequest()
	}
	response = NewCheckAddOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewCheckBindOrganizationMemberEmailRequest() (request *CheckBindOrganizationMemberEmailRequest) {
	request = &CheckBindOrganizationMemberEmailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckBindOrganizationMemberEmail")
	return
}

func NewCheckBindOrganizationMemberEmailResponse() (response *CheckBindOrganizationMemberEmailResponse) {
	response = &CheckBindOrganizationMemberEmailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 验证成员绑定邮箱
func (c *Client) CheckBindOrganizationMemberEmail(request *CheckBindOrganizationMemberEmailRequest) (response *CheckBindOrganizationMemberEmailResponse, err error) {
	if request == nil {
		request = NewCheckBindOrganizationMemberEmailRequest()
	}
	response = NewCheckBindOrganizationMemberEmailResponse()
	err = c.Send(request, response)
	return
}

func NewCheckOrganizationAuthManageUinRequest() (request *CheckOrganizationAuthManageUinRequest) {
	request = &CheckOrganizationAuthManageUinRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckOrganizationAuthManageUin")
	return
}

func NewCheckOrganizationAuthManageUinResponse() (response *CheckOrganizationAuthManageUinResponse) {
	response = &CheckOrganizationAuthManageUinResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查是否企业组织主体管理账号
func (c *Client) CheckOrganizationAuthManageUin(request *CheckOrganizationAuthManageUinRequest) (response *CheckOrganizationAuthManageUinResponse, err error) {
	if request == nil {
		request = NewCheckOrganizationAuthManageUinRequest()
	}
	response = NewCheckOrganizationAuthManageUinResponse()
	err = c.Send(request, response)
	return
}

func NewCheckOrganizationMemberAuthRequest() (request *CheckOrganizationMemberAuthRequest) {
	request = &CheckOrganizationMemberAuthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckOrganizationMemberAuth")
	return
}

func NewCheckOrganizationMemberAuthResponse() (response *CheckOrganizationMemberAuthResponse) {
	response = &CheckOrganizationMemberAuthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查成员企业实名是否和管理员一样
func (c *Client) CheckOrganizationMemberAuth(request *CheckOrganizationMemberAuthRequest) (response *CheckOrganizationMemberAuthResponse, err error) {
	if request == nil {
		request = NewCheckOrganizationMemberAuthRequest()
	}
	response = NewCheckOrganizationMemberAuthResponse()
	err = c.Send(request, response)
	return
}

func NewCheckOrganizationMemberAuthRelationRequest() (request *CheckOrganizationMemberAuthRelationRequest) {
	request = &CheckOrganizationMemberAuthRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckOrganizationMemberAuthRelation")
	return
}

func NewCheckOrganizationMemberAuthRelationResponse() (response *CheckOrganizationMemberAuthRelationResponse) {
	response = &CheckOrganizationMemberAuthRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查管理员和成员主体是否互信
func (c *Client) CheckOrganizationMemberAuthRelation(request *CheckOrganizationMemberAuthRelationRequest) (response *CheckOrganizationMemberAuthRelationResponse, err error) {
	if request == nil {
		request = NewCheckOrganizationMemberAuthRelationRequest()
	}
	response = NewCheckOrganizationMemberAuthRelationResponse()
	err = c.Send(request, response)
	return
}

func NewCheckOrganizationMemberPermissionRequest() (request *CheckOrganizationMemberPermissionRequest) {
	request = &CheckOrganizationMemberPermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CheckOrganizationMemberPermission")
	return
}

func NewCheckOrganizationMemberPermissionResponse() (response *CheckOrganizationMemberPermissionResponse) {
	response = &CheckOrganizationMemberPermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验成员权限设置
func (c *Client) CheckOrganizationMemberPermission(request *CheckOrganizationMemberPermissionRequest) (response *CheckOrganizationMemberPermissionResponse, err error) {
	if request == nil {
		request = NewCheckOrganizationMemberPermissionRequest()
	}
	response = NewCheckOrganizationMemberPermissionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrgMemberProductServiceRoleRequest() (request *CreateOrgMemberProductServiceRoleRequest) {
	request = &CreateOrgMemberProductServiceRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CreateOrgMemberProductServiceRole")
	return
}

func NewCreateOrgMemberProductServiceRoleResponse() (response *CreateOrgMemberProductServiceRoleResponse) {
	response = &CreateOrgMemberProductServiceRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建成员产品服务角色
func (c *Client) CreateOrgMemberProductServiceRole(request *CreateOrgMemberProductServiceRoleRequest) (response *CreateOrgMemberProductServiceRoleResponse, err error) {
	if request == nil {
		request = NewCreateOrgMemberProductServiceRoleRequest()
	}
	response = NewCreateOrgMemberProductServiceRoleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrgServiceAssignRequest() (request *CreateOrgServiceAssignRequest) {
	request = &CreateOrgServiceAssignRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CreateOrgServiceAssign")
	return
}

func NewCreateOrgServiceAssignResponse() (response *CreateOrgServiceAssignResponse) {
	response = &CreateOrgServiceAssignResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加集团服务委派管理员
func (c *Client) CreateOrgServiceAssign(request *CreateOrgServiceAssignRequest) (response *CreateOrgServiceAssignResponse, err error) {
	if request == nil {
		request = NewCreateOrgServiceAssignRequest()
	}
	response = NewCreateOrgServiceAssignResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrganizationMembersRequest() (request *CreateOrganizationMembersRequest) {
	request = &CreateOrganizationMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "CreateOrganizationMembers")
	return
}

func NewCreateOrganizationMembersResponse() (response *CreateOrganizationMembersResponse) {
	response = &CreateOrganizationMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量添加组织成员
func (c *Client) CreateOrganizationMembers(request *CreateOrganizationMembersRequest) (response *CreateOrganizationMembersResponse, err error) {
	if request == nil {
		request = NewCreateOrganizationMembersRequest()
	}
	response = NewCreateOrganizationMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrgServiceAssignRequest() (request *DeleteOrgServiceAssignRequest) {
	request = &DeleteOrgServiceAssignRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrgServiceAssign")
	return
}

func NewDeleteOrgServiceAssignResponse() (response *DeleteOrgServiceAssignResponse) {
	response = &DeleteOrgServiceAssignResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集团服务委派管理员
func (c *Client) DeleteOrgServiceAssign(request *DeleteOrgServiceAssignRequest) (response *DeleteOrgServiceAssignResponse, err error) {
	if request == nil {
		request = NewDeleteOrgServiceAssignRequest()
	}
	response = NewDeleteOrgServiceAssignResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteOrganizationNodeTagsRequest() (request *DeleteOrganizationNodeTagsRequest) {
	request = &DeleteOrganizationNodeTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DeleteOrganizationNodeTags")
	return
}

func NewDeleteOrganizationNodeTagsResponse() (response *DeleteOrganizationNodeTagsResponse) {
	response = &DeleteOrganizationNodeTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除组织节点标签
func (c *Client) DeleteOrganizationNodeTags(request *DeleteOrganizationNodeTagsRequest) (response *DeleteOrganizationNodeTagsResponse, err error) {
	if request == nil {
		request = NewDeleteOrganizationNodeTagsRequest()
	}
	response = NewDeleteOrganizationNodeTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDenyMemberChangePermissionRequest() (request *DenyMemberChangePermissionRequest) {
	request = &DenyMemberChangePermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DenyMemberChangePermission")
	return
}

func NewDenyMemberChangePermissionResponse() (response *DenyMemberChangePermissionResponse) {
	response = &DenyMemberChangePermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拒绝成员权限变更
func (c *Client) DenyMemberChangePermission(request *DenyMemberChangePermissionRequest) (response *DenyMemberChangePermissionResponse, err error) {
	if request == nil {
		request = NewDenyMemberChangePermissionRequest()
	}
	response = NewDenyMemberChangePermissionResponse()
	err = c.Send(request, response)
	return
}

func NewDenyOrganizationCreateRecordRequest() (request *DenyOrganizationCreateRecordRequest) {
	request = &DenyOrganizationCreateRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DenyOrganizationCreateRecord")
	return
}

func NewDenyOrganizationCreateRecordResponse() (response *DenyOrganizationCreateRecordResponse) {
	response = &DenyOrganizationCreateRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拒绝企业组织创建成员
func (c *Client) DenyOrganizationCreateRecord(request *DenyOrganizationCreateRecordRequest) (response *DenyOrganizationCreateRecordResponse, err error) {
	if request == nil {
		request = NewDenyOrganizationCreateRecordRequest()
	}
	response = NewDenyOrganizationCreateRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDenyOrganizationInvitationRequest() (request *DenyOrganizationInvitationRequest) {
	request = &DenyOrganizationInvitationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DenyOrganizationInvitation")
	return
}

func NewDenyOrganizationInvitationResponse() (response *DenyOrganizationInvitationResponse) {
	response = &DenyOrganizationInvitationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拒绝企业组织邀请
func (c *Client) DenyOrganizationInvitation(request *DenyOrganizationInvitationRequest) (response *DenyOrganizationInvitationResponse, err error) {
	if request == nil {
		request = NewDenyOrganizationInvitationRequest()
	}
	response = NewDenyOrganizationInvitationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllParentNodesRequest() (request *DescribeAllParentNodesRequest) {
	request = &DescribeAllParentNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeAllParentNodes")
	return
}

func NewDescribeAllParentNodesResponse() (response *DescribeAllParentNodesResponse) {
	response = &DescribeAllParentNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据节点获取所有父节点,直到根节点为止。
func (c *Client) DescribeAllParentNodes(request *DescribeAllParentNodesRequest) (response *DescribeAllParentNodesResponse, err error) {
	if request == nil {
		request = NewDescribeAllParentNodesRequest()
	}
	response = NewDescribeAllParentNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEffectivePolicyRequest() (request *DescribeEffectivePolicyRequest) {
	request = &DescribeEffectivePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeEffectivePolicy")
	return
}

func NewDescribeEffectivePolicyResponse() (response *DescribeEffectivePolicyResponse) {
	response = &DescribeEffectivePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询目标关联的有效策略
func (c *Client) DescribeEffectivePolicy(request *DescribeEffectivePolicyRequest) (response *DescribeEffectivePolicyResponse, err error) {
	if request == nil {
		request = NewDescribeEffectivePolicyRequest()
	}
	response = NewDescribeEffectivePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMemberChangePermissionRecordsRequest() (request *DescribeMemberChangePermissionRecordsRequest) {
	request = &DescribeMemberChangePermissionRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeMemberChangePermissionRecords")
	return
}

func NewDescribeMemberChangePermissionRecordsResponse() (response *DescribeMemberChangePermissionRecordsResponse) {
	response = &DescribeMemberChangePermissionRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员权限变更记录列表
func (c *Client) DescribeMemberChangePermissionRecords(request *DescribeMemberChangePermissionRecordsRequest) (response *DescribeMemberChangePermissionRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeMemberChangePermissionRecordsRequest()
	}
	response = NewDescribeMemberChangePermissionRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMemberDeletionPermissionRequest() (request *DescribeMemberDeletionPermissionRequest) {
	request = &DescribeMemberDeletionPermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeMemberDeletionPermission")
	return
}

func NewDescribeMemberDeletionPermissionResponse() (response *DescribeMemberDeletionPermissionResponse) {
	response = &DescribeMemberDeletionPermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询成员删除许可
func (c *Client) DescribeMemberDeletionPermission(request *DescribeMemberDeletionPermissionRequest) (response *DescribeMemberDeletionPermissionResponse, err error) {
	if request == nil {
		request = NewDescribeMemberDeletionPermissionRequest()
	}
	response = NewDescribeMemberDeletionPermissionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationAuthPoliciesRequest() (request *DescribeOrganizationAuthPoliciesRequest) {
	request = &DescribeOrganizationAuthPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationAuthPolicies")
	return
}

func NewDescribeOrganizationAuthPoliciesResponse() (response *DescribeOrganizationAuthPoliciesResponse) {
	response = &DescribeOrganizationAuthPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织访问授权策略列表
func (c *Client) DescribeOrganizationAuthPolicies(request *DescribeOrganizationAuthPoliciesRequest) (response *DescribeOrganizationAuthPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationAuthPoliciesRequest()
	}
	response = NewDescribeOrganizationAuthPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationBeInviteRecordRequest() (request *DescribeOrganizationBeInviteRecordRequest) {
	request = &DescribeOrganizationBeInviteRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationBeInviteRecord")
	return
}

func NewDescribeOrganizationBeInviteRecordResponse() (response *DescribeOrganizationBeInviteRecordResponse) {
	response = &DescribeOrganizationBeInviteRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取我被组织邀请信息列表
func (c *Client) DescribeOrganizationBeInviteRecord(request *DescribeOrganizationBeInviteRecordRequest) (response *DescribeOrganizationBeInviteRecordResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationBeInviteRecordRequest()
	}
	response = NewDescribeOrganizationBeInviteRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationCollPoliciesRequest() (request *DescribeOrganizationCollPoliciesRequest) {
	request = &DescribeOrganizationCollPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationCollPolicies")
	return
}

func NewDescribeOrganizationCollPoliciesResponse() (response *DescribeOrganizationCollPoliciesResponse) {
	response = &DescribeOrganizationCollPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 子账号登录控制台，获取被绑定授权策略列表
func (c *Client) DescribeOrganizationCollPolicies(request *DescribeOrganizationCollPoliciesRequest) (response *DescribeOrganizationCollPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationCollPoliciesRequest()
	}
	response = NewDescribeOrganizationCollPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationCreateRecordRequest() (request *DescribeOrganizationCreateRecordRequest) {
	request = &DescribeOrganizationCreateRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationCreateRecord")
	return
}

func NewDescribeOrganizationCreateRecordResponse() (response *DescribeOrganizationCreateRecordResponse) {
	response = &DescribeOrganizationCreateRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员创建列表
func (c *Client) DescribeOrganizationCreateRecord(request *DescribeOrganizationCreateRecordRequest) (response *DescribeOrganizationCreateRecordResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationCreateRecordRequest()
	}
	response = NewDescribeOrganizationCreateRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationFinancialByMemberRequest() (request *DescribeOrganizationFinancialByMemberRequest) {
	request = &DescribeOrganizationFinancialByMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationFinancialByMember")
	return
}

func NewDescribeOrganizationFinancialByMemberResponse() (response *DescribeOrganizationFinancialByMemberResponse) {
	response = &DescribeOrganizationFinancialByMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 以成员维度获取组织财务信息
func (c *Client) DescribeOrganizationFinancialByMember(request *DescribeOrganizationFinancialByMemberRequest) (response *DescribeOrganizationFinancialByMemberResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationFinancialByMemberRequest()
	}
	response = NewDescribeOrganizationFinancialByMemberResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationFinancialByMonthRequest() (request *DescribeOrganizationFinancialByMonthRequest) {
	request = &DescribeOrganizationFinancialByMonthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationFinancialByMonth")
	return
}

func NewDescribeOrganizationFinancialByMonthResponse() (response *DescribeOrganizationFinancialByMonthResponse) {
	response = &DescribeOrganizationFinancialByMonthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 以月维度获取组织财务信息趋势
func (c *Client) DescribeOrganizationFinancialByMonth(request *DescribeOrganizationFinancialByMonthRequest) (response *DescribeOrganizationFinancialByMonthResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationFinancialByMonthRequest()
	}
	response = NewDescribeOrganizationFinancialByMonthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationFinancialByProductRequest() (request *DescribeOrganizationFinancialByProductRequest) {
	request = &DescribeOrganizationFinancialByProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationFinancialByProduct")
	return
}

func NewDescribeOrganizationFinancialByProductResponse() (response *DescribeOrganizationFinancialByProductResponse) {
	response = &DescribeOrganizationFinancialByProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 以产品维度获取组织财务信息
func (c *Client) DescribeOrganizationFinancialByProduct(request *DescribeOrganizationFinancialByProductRequest) (response *DescribeOrganizationFinancialByProductResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationFinancialByProductRequest()
	}
	response = NewDescribeOrganizationFinancialByProductResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationFinancialMemberNumRequest() (request *DescribeOrganizationFinancialMemberNumRequest) {
	request = &DescribeOrganizationFinancialMemberNumRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationFinancialMemberNum")
	return
}

func NewDescribeOrganizationFinancialMemberNumResponse() (response *DescribeOrganizationFinancialMemberNumResponse) {
	response = &DescribeOrganizationFinancialMemberNumResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织财务架构统计数据
func (c *Client) DescribeOrganizationFinancialMemberNum(request *DescribeOrganizationFinancialMemberNumRequest) (response *DescribeOrganizationFinancialMemberNumResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationFinancialMemberNumRequest()
	}
	response = NewDescribeOrganizationFinancialMemberNumResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationIdentityRequest() (request *DescribeOrganizationIdentityRequest) {
	request = &DescribeOrganizationIdentityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationIdentity")
	return
}

func NewDescribeOrganizationIdentityResponse() (response *DescribeOrganizationIdentityResponse) {
	response = &DescribeOrganizationIdentityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织身份详情
func (c *Client) DescribeOrganizationIdentity(request *DescribeOrganizationIdentityRequest) (response *DescribeOrganizationIdentityResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationIdentityRequest()
	}
	response = NewDescribeOrganizationIdentityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationInviteRecordRequest() (request *DescribeOrganizationInviteRecordRequest) {
	request = &DescribeOrganizationInviteRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationInviteRecord")
	return
}

func NewDescribeOrganizationInviteRecordResponse() (response *DescribeOrganizationInviteRecordResponse) {
	response = &DescribeOrganizationInviteRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员邀请信息列表
func (c *Client) DescribeOrganizationInviteRecord(request *DescribeOrganizationInviteRecordRequest) (response *DescribeOrganizationInviteRecordResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationInviteRecordRequest()
	}
	response = NewDescribeOrganizationInviteRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMemberNodesRequest() (request *DescribeOrganizationMemberNodesRequest) {
	request = &DescribeOrganizationMemberNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberNodes")
	return
}

func NewDescribeOrganizationMemberNodesResponse() (response *DescribeOrganizationMemberNodesResponse) {
	response = &DescribeOrganizationMemberNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询成员节点列表。
func (c *Client) DescribeOrganizationMemberNodes(request *DescribeOrganizationMemberNodesRequest) (response *DescribeOrganizationMemberNodesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMemberNodesRequest()
	}
	response = NewDescribeOrganizationMemberNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMemberPoliciesRequest() (request *DescribeOrganizationMemberPoliciesRequest) {
	request = &DescribeOrganizationMemberPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMemberPolicies")
	return
}

func NewDescribeOrganizationMemberPoliciesResponse() (response *DescribeOrganizationMemberPoliciesResponse) {
	response = &DescribeOrganizationMemberPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员的授权策略列表
func (c *Client) DescribeOrganizationMemberPolicies(request *DescribeOrganizationMemberPoliciesRequest) (response *DescribeOrganizationMemberPoliciesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMemberPoliciesRequest()
	}
	response = NewDescribeOrganizationMemberPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMembersAuthAccountRequest() (request *DescribeOrganizationMembersAuthAccountRequest) {
	request = &DescribeOrganizationMembersAuthAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembersAuthAccount")
	return
}

func NewDescribeOrganizationMembersAuthAccountResponse() (response *DescribeOrganizationMembersAuthAccountResponse) {
	response = &DescribeOrganizationMembersAuthAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员访问授权子账号关系列表
func (c *Client) DescribeOrganizationMembersAuthAccount(request *DescribeOrganizationMembersAuthAccountRequest) (response *DescribeOrganizationMembersAuthAccountResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMembersAuthAccountRequest()
	}
	response = NewDescribeOrganizationMembersAuthAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMembersAuthPolicyRequest() (request *DescribeOrganizationMembersAuthPolicyRequest) {
	request = &DescribeOrganizationMembersAuthPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembersAuthPolicy")
	return
}

func NewDescribeOrganizationMembersAuthPolicyResponse() (response *DescribeOrganizationMembersAuthPolicyResponse) {
	response = &DescribeOrganizationMembersAuthPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织成员访问策略列表
func (c *Client) DescribeOrganizationMembersAuthPolicy(request *DescribeOrganizationMembersAuthPolicyRequest) (response *DescribeOrganizationMembersAuthPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMembersAuthPolicyRequest()
	}
	response = NewDescribeOrganizationMembersAuthPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationMembersCanAuthIdentitiesRequest() (request *DescribeOrganizationMembersCanAuthIdentitiesRequest) {
	request = &DescribeOrganizationMembersCanAuthIdentitiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationMembersCanAuthIdentities")
	return
}

func NewDescribeOrganizationMembersCanAuthIdentitiesResponse() (response *DescribeOrganizationMembersCanAuthIdentitiesResponse) {
	response = &DescribeOrganizationMembersCanAuthIdentitiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员都被授权的访问权限列表
func (c *Client) DescribeOrganizationMembersCanAuthIdentities(request *DescribeOrganizationMembersCanAuthIdentitiesRequest) (response *DescribeOrganizationMembersCanAuthIdentitiesResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationMembersCanAuthIdentitiesRequest()
	}
	response = NewDescribeOrganizationMembersCanAuthIdentitiesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodeByNameRequest() (request *DescribeOrganizationNodeByNameRequest) {
	request = &DescribeOrganizationNodeByNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodeByName")
	return
}

func NewDescribeOrganizationNodeByNameResponse() (response *DescribeOrganizationNodeByNameResponse) {
	response = &DescribeOrganizationNodeByNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据节点名获取节点
func (c *Client) DescribeOrganizationNodeByName(request *DescribeOrganizationNodeByNameRequest) (response *DescribeOrganizationNodeByNameResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodeByNameRequest()
	}
	response = NewDescribeOrganizationNodeByNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodeMemberRecordsRequest() (request *DescribeOrganizationNodeMemberRecordsRequest) {
	request = &DescribeOrganizationNodeMemberRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodeMemberRecords")
	return
}

func NewDescribeOrganizationNodeMemberRecordsResponse() (response *DescribeOrganizationNodeMemberRecordsResponse) {
	response = &DescribeOrganizationNodeMemberRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织节点成员变更记录列表
func (c *Client) DescribeOrganizationNodeMemberRecords(request *DescribeOrganizationNodeMemberRecordsRequest) (response *DescribeOrganizationNodeMemberRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodeMemberRecordsRequest()
	}
	response = NewDescribeOrganizationNodeMemberRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodeRecordsRequest() (request *DescribeOrganizationNodeRecordsRequest) {
	request = &DescribeOrganizationNodeRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodeRecords")
	return
}

func NewDescribeOrganizationNodeRecordsResponse() (response *DescribeOrganizationNodeRecordsResponse) {
	response = &DescribeOrganizationNodeRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织节点变更记录列表
func (c *Client) DescribeOrganizationNodeRecords(request *DescribeOrganizationNodeRecordsRequest) (response *DescribeOrganizationNodeRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodeRecordsRequest()
	}
	response = NewDescribeOrganizationNodeRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationNodeTagsRequest() (request *DescribeOrganizationNodeTagsRequest) {
	request = &DescribeOrganizationNodeTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationNodeTags")
	return
}

func NewDescribeOrganizationNodeTagsResponse() (response *DescribeOrganizationNodeTagsResponse) {
	response = &DescribeOrganizationNodeTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织节点标签
func (c *Client) DescribeOrganizationNodeTags(request *DescribeOrganizationNodeTagsRequest) (response *DescribeOrganizationNodeTagsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationNodeTagsRequest()
	}
	response = NewDescribeOrganizationNodeTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationOverViewRequest() (request *DescribeOrganizationOverViewRequest) {
	request = &DescribeOrganizationOverViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationOverView")
	return
}

func NewDescribeOrganizationOverViewResponse() (response *DescribeOrganizationOverViewResponse) {
	response = &DescribeOrganizationOverViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按天获取组织概览信息，包含成员，节点，子账号，角色数信息。
func (c *Client) DescribeOrganizationOverView(request *DescribeOrganizationOverViewRequest) (response *DescribeOrganizationOverViewResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationOverViewRequest()
	}
	response = NewDescribeOrganizationOverViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationPendingCreateRecordRequest() (request *DescribeOrganizationPendingCreateRecordRequest) {
	request = &DescribeOrganizationPendingCreateRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationPendingCreateRecord")
	return
}

func NewDescribeOrganizationPendingCreateRecordResponse() (response *DescribeOrganizationPendingCreateRecordResponse) {
	response = &DescribeOrganizationPendingCreateRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员创建待审核列表
func (c *Client) DescribeOrganizationPendingCreateRecord(request *DescribeOrganizationPendingCreateRecordRequest) (response *DescribeOrganizationPendingCreateRecordResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationPendingCreateRecordRequest()
	}
	response = NewDescribeOrganizationPendingCreateRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationRecordsRequest() (request *DescribeOrganizationRecordsRequest) {
	request = &DescribeOrganizationRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationRecords")
	return
}

func NewDescribeOrganizationRecordsResponse() (response *DescribeOrganizationRecordsResponse) {
	response = &DescribeOrganizationRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取组织成员动态列表
func (c *Client) DescribeOrganizationRecords(request *DescribeOrganizationRecordsRequest) (response *DescribeOrganizationRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationRecordsRequest()
	}
	response = NewDescribeOrganizationRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationServiceRoleRequest() (request *DescribeOrganizationServiceRoleRequest) {
	request = &DescribeOrganizationServiceRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationServiceRole")
	return
}

func NewDescribeOrganizationServiceRoleResponse() (response *DescribeOrganizationServiceRoleResponse) {
	response = &DescribeOrganizationServiceRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取企业组织服务角色
func (c *Client) DescribeOrganizationServiceRole(request *DescribeOrganizationServiceRoleRequest) (response *DescribeOrganizationServiceRoleResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationServiceRoleRequest()
	}
	response = NewDescribeOrganizationServiceRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationSubAccountByDayRequest() (request *DescribeOrganizationSubAccountByDayRequest) {
	request = &DescribeOrganizationSubAccountByDayRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationSubAccountByDay")
	return
}

func NewDescribeOrganizationSubAccountByDayResponse() (response *DescribeOrganizationSubAccountByDayResponse) {
	response = &DescribeOrganizationSubAccountByDayResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 以天维度获取组织成员和子账号数量趋势
func (c *Client) DescribeOrganizationSubAccountByDay(request *DescribeOrganizationSubAccountByDayRequest) (response *DescribeOrganizationSubAccountByDayResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationSubAccountByDayRequest()
	}
	response = NewDescribeOrganizationSubAccountByDayResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrganizationSubAccountByMonthRequest() (request *DescribeOrganizationSubAccountByMonthRequest) {
	request = &DescribeOrganizationSubAccountByMonthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeOrganizationSubAccountByMonth")
	return
}

func NewDescribeOrganizationSubAccountByMonthResponse() (response *DescribeOrganizationSubAccountByMonthResponse) {
	response = &DescribeOrganizationSubAccountByMonthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 以月维度获取组织成员和子账号数量趋势
func (c *Client) DescribeOrganizationSubAccountByMonth(request *DescribeOrganizationSubAccountByMonthRequest) (response *DescribeOrganizationSubAccountByMonthResponse, err error) {
	if request == nil {
		request = NewDescribeOrganizationSubAccountByMonthRequest()
	}
	response = NewDescribeOrganizationSubAccountByMonthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReportCreationRequest() (request *DescribeReportCreationRequest) {
	request = &DescribeReportCreationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "DescribeReportCreation")
	return
}

func NewDescribeReportCreationResponse() (response *DescribeReportCreationResponse) {
	response = &DescribeReportCreationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询成员标签检测合规报告
func (c *Client) DescribeReportCreation(request *DescribeReportCreationRequest) (response *DescribeReportCreationResponse, err error) {
	if request == nil {
		request = NewDescribeReportCreationRequest()
	}
	response = NewDescribeReportCreationResponse()
	err = c.Send(request, response)
	return
}

func NewInviteOrganizationMemberRequest() (request *InviteOrganizationMemberRequest) {
	request = &InviteOrganizationMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "InviteOrganizationMember")
	return
}

func NewInviteOrganizationMemberResponse() (response *InviteOrganizationMemberResponse) {
	response = &InviteOrganizationMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 邀请组织成员
func (c *Client) InviteOrganizationMember(request *InviteOrganizationMemberRequest) (response *InviteOrganizationMemberResponse, err error) {
	if request == nil {
		request = NewInviteOrganizationMemberRequest()
	}
	response = NewInviteOrganizationMemberResponse()
	err = c.Send(request, response)
	return
}

func NewListComplianceSummaryRequest() (request *ListComplianceSummaryRequest) {
	request = &ListComplianceSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "ListComplianceSummary")
	return
}

func NewListComplianceSummaryResponse() (response *ListComplianceSummaryResponse) {
	response = &ListComplianceSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取成员标签检测结果列表
func (c *Client) ListComplianceSummary(request *ListComplianceSummaryRequest) (response *ListComplianceSummaryResponse, err error) {
	if request == nil {
		request = NewListComplianceSummaryRequest()
	}
	response = NewListComplianceSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewListFinancialProductRequest() (request *ListFinancialProductRequest) {
	request = &ListFinancialProductRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "ListFinancialProduct")
	return
}

func NewListFinancialProductResponse() (response *ListFinancialProductResponse) {
	response = &ListFinancialProductResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询组织成员账单的产品列表
func (c *Client) ListFinancialProduct(request *ListFinancialProductRequest) (response *ListFinancialProductResponse, err error) {
	if request == nil {
		request = NewListFinancialProductRequest()
	}
	response = NewListFinancialProductResponse()
	err = c.Send(request, response)
	return
}

func NewListNonCompliantResourceRequest() (request *ListNonCompliantResourceRequest) {
	request = &ListNonCompliantResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "ListNonCompliantResource")
	return
}

func NewListNonCompliantResourceResponse() (response *ListNonCompliantResourceResponse) {
	response = &ListNonCompliantResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取成员标签检测不合规资源列表
func (c *Client) ListNonCompliantResource(request *ListNonCompliantResourceRequest) (response *ListNonCompliantResourceResponse, err error) {
	if request == nil {
		request = NewListNonCompliantResourceRequest()
	}
	response = NewListNonCompliantResourceResponse()
	err = c.Send(request, response)
	return
}

func NewListOrgMemberSubAccountRequest() (request *ListOrgMemberSubAccountRequest) {
	request = &ListOrgMemberSubAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "ListOrgMemberSubAccount")
	return
}

func NewListOrgMemberSubAccountResponse() (response *ListOrgMemberSubAccountResponse) {
	response = &ListOrgMemberSubAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询同一负责人子账号的列表
func (c *Client) ListOrgMemberSubAccount(request *ListOrgMemberSubAccountRequest) (response *ListOrgMemberSubAccountResponse, err error) {
	if request == nil {
		request = NewListOrgMemberSubAccountRequest()
	}
	response = NewListOrgMemberSubAccountResponse()
	err = c.Send(request, response)
	return
}

func NewListOrgServiceAssignMemberRequest() (request *ListOrgServiceAssignMemberRequest) {
	request = &ListOrgServiceAssignMemberRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "ListOrgServiceAssignMember")
	return
}

func NewListOrgServiceAssignMemberResponse() (response *ListOrgServiceAssignMemberResponse) {
	response = &ListOrgServiceAssignMemberResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集团服务委派管理员列表
func (c *Client) ListOrgServiceAssignMember(request *ListOrgServiceAssignMemberRequest) (response *ListOrgServiceAssignMemberResponse, err error) {
	if request == nil {
		request = NewListOrgServiceAssignMemberRequest()
	}
	response = NewListOrgServiceAssignMemberResponse()
	err = c.Send(request, response)
	return
}

func NewListOrganizationServiceRequest() (request *ListOrganizationServiceRequest) {
	request = &ListOrganizationServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "ListOrganizationService")
	return
}

func NewListOrganizationServiceResponse() (response *ListOrganizationServiceResponse) {
	response = &ListOrganizationServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集团服务设置列表
func (c *Client) ListOrganizationService(request *ListOrganizationServiceRequest) (response *ListOrganizationServiceResponse, err error) {
	if request == nil {
		request = NewListOrganizationServiceRequest()
	}
	response = NewListOrganizationServiceResponse()
	err = c.Send(request, response)
	return
}

func NewListPoliciesForTargetRequest() (request *ListPoliciesForTargetRequest) {
	request = &ListPoliciesForTargetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "ListPoliciesForTarget")
	return
}

func NewListPoliciesForTargetResponse() (response *ListPoliciesForTargetResponse) {
	response = &ListPoliciesForTargetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ListPoliciesForTarget）查询目标关联的策略列表
func (c *Client) ListPoliciesForTarget(request *ListPoliciesForTargetRequest) (response *ListPoliciesForTargetResponse, err error) {
	if request == nil {
		request = NewListPoliciesForTargetRequest()
	}
	response = NewListPoliciesForTargetResponse()
	err = c.Send(request, response)
	return
}

func NewSendOrgMemberAccountBindEmailRequest() (request *SendOrgMemberAccountBindEmailRequest) {
	request = &SendOrgMemberAccountBindEmailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "SendOrgMemberAccountBindEmail")
	return
}

func NewSendOrgMemberAccountBindEmailResponse() (response *SendOrgMemberAccountBindEmailResponse) {
	response = &SendOrgMemberAccountBindEmailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新发送成员绑定邮箱激活邮件
func (c *Client) SendOrgMemberAccountBindEmail(request *SendOrgMemberAccountBindEmailRequest) (response *SendOrgMemberAccountBindEmailResponse, err error) {
	if request == nil {
		request = NewSendOrgMemberAccountBindEmailRequest()
	}
	response = NewSendOrgMemberAccountBindEmailResponse()
	err = c.Send(request, response)
	return
}

func NewSendSmsVerifyCodeForBindPhoneRequest() (request *SendSmsVerifyCodeForBindPhoneRequest) {
	request = &SendSmsVerifyCodeForBindPhoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "SendSmsVerifyCodeForBindPhone")
	return
}

func NewSendSmsVerifyCodeForBindPhoneResponse() (response *SendSmsVerifyCodeForBindPhoneResponse) {
	response = &SendSmsVerifyCodeForBindPhoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定安全手机发送验证码
func (c *Client) SendSmsVerifyCodeForBindPhone(request *SendSmsVerifyCodeForBindPhoneRequest) (response *SendSmsVerifyCodeForBindPhoneResponse, err error) {
	if request == nil {
		request = NewSendSmsVerifyCodeForBindPhoneRequest()
	}
	response = NewSendSmsVerifyCodeForBindPhoneResponse()
	err = c.Send(request, response)
	return
}

func NewSetMemberDeletionPermissionRequest() (request *SetMemberDeletionPermissionRequest) {
	request = &SetMemberDeletionPermissionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "SetMemberDeletionPermission")
	return
}

func NewSetMemberDeletionPermissionResponse() (response *SetMemberDeletionPermissionResponse) {
	response = &SetMemberDeletionPermissionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启或关闭成员删除许可
func (c *Client) SetMemberDeletionPermission(request *SetMemberDeletionPermissionRequest) (response *SetMemberDeletionPermissionResponse, err error) {
	if request == nil {
		request = NewSetMemberDeletionPermissionRequest()
	}
	response = NewSetMemberDeletionPermissionResponse()
	err = c.Send(request, response)
	return
}

func NewStartReportCreationRequest() (request *StartReportCreationRequest) {
	request = &StartReportCreationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "StartReportCreation")
	return
}

func NewStartReportCreationResponse() (response *StartReportCreationResponse) {
	response = &StartReportCreationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成成员标签检测合规报告
func (c *Client) StartReportCreation(request *StartReportCreationRequest) (response *StartReportCreationResponse, err error) {
	if request == nil {
		request = NewStartReportCreationRequest()
	}
	response = NewStartReportCreationResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationMembersPolicyRequest() (request *UpdateOrganizationMembersPolicyRequest) {
	request = &UpdateOrganizationMembersPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationMembersPolicy")
	return
}

func NewUpdateOrganizationMembersPolicyResponse() (response *UpdateOrganizationMembersPolicyResponse) {
	response = &UpdateOrganizationMembersPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改组织成员访问策略
func (c *Client) UpdateOrganizationMembersPolicy(request *UpdateOrganizationMembersPolicyRequest) (response *UpdateOrganizationMembersPolicyResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationMembersPolicyRequest()
	}
	response = NewUpdateOrganizationMembersPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateOrganizationNodeTagRequest() (request *UpdateOrganizationNodeTagRequest) {
	request = &UpdateOrganizationNodeTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("organization", APIVersion, "UpdateOrganizationNodeTag")
	return
}

func NewUpdateOrganizationNodeTagResponse() (response *UpdateOrganizationNodeTagResponse) {
	response = &UpdateOrganizationNodeTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新组织节点标签
func (c *Client) UpdateOrganizationNodeTag(request *UpdateOrganizationNodeTagRequest) (response *UpdateOrganizationNodeTagResponse, err error) {
	if request == nil {
		request = NewUpdateOrganizationNodeTagRequest()
	}
	response = NewUpdateOrganizationNodeTagResponse()
	err = c.Send(request, response)
	return
}

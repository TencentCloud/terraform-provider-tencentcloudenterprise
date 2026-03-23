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
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_ORGANIZATIONEXISTALREADY = "FailedOperation.OrganizationExistAlready"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWCREATEORGANIZATION = "UnsupportedOperation.CreateMemberNotAllowCreateOrganization"
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
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_ORGANIZATIONNOTEMPTY = "FailedOperation.OrganizationNotEmpty"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//  FAILEDOPERATION_QUITSHAREUINT = "FailedOperation.QuitShareUint"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHAREUNITNOTEMPTY = "FailedOperation.ShareUnitNotEmpty"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
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
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITSHAREUINTERROR = "FailedOperation.QuitShareUintError"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWDELETE = "UnsupportedOperation.CreateMemberNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
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
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYNAMEUSED = "FailedOperation.OrganizationIdentityNameUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_IDENTITYEXCEEDLIMIT = "LimitExceeded.IdentityExceedLimit"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
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
//  FAILEDOPERATION_GETPOLICYDETAIL = "FailedOperation.GetPolicyDetail"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
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
//  FAILEDOPERATION_ORGANIZATIONIDENTITYINUSED = "FailedOperation.OrganizationIdentityInUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
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
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  FAILEDOPERATION_POLICYFULL = "FailedOperation.PolicyFull"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"
//  LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
//  RESOURCEUNAVAILABLE = "ResourceUnavailable"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
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
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  FAILEDOPERATION_POLICYNAMEINUSE = "FailedOperation.PolicyNameInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ACTIONERROR = "InvalidParameter.ActionError"
//  INVALIDPARAMETER_ACTIONMISS = "InvalidParameter.ActionMiss"
//  INVALIDPARAMETER_ACTIONNOTEXIST = "InvalidParameter.ActionNotExist"
//  INVALIDPARAMETER_ACTIONSERVICENOTEXIST = "InvalidParameter.ActionServiceNotExist"
//  INVALIDPARAMETER_CONDITIONCONTENTERROR = "InvalidParameter.ConditionContentError"
//  INVALIDPARAMETER_CONDITIONERROR = "InvalidParameter.ConditionError"
//  INVALIDPARAMETER_CONDITIONTYPEERROR = "InvalidParameter.ConditionTypeError"
//  INVALIDPARAMETER_DESCRIPTIONLENGTHOVERLIMIT = "InvalidParameter.DescriptionLengthOverlimit"
//  INVALIDPARAMETER_EFFECTERROR = "InvalidParameter.EffectError"
//  INVALIDPARAMETER_NOTSUPPORTPRODUCT = "InvalidParameter.NotSupportProduct"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYDOCUMENTERROR = "InvalidParameter.PolicyDocumentError"
//  INVALIDPARAMETER_POLICYDOCUMENTLENGTHOVERLIMIT = "InvalidParameter.PolicyDocumentLengthOverLimit"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  INVALIDPARAMETER_POLICYKEYDUPLICATED = "InvalidParameter.PolicyKeyDuplicated"
//  INVALIDPARAMETER_POLICYNAMEERROR = "InvalidParameter.PolicyNameError"
//  INVALIDPARAMETER_POLICYNAMEEXISTED = "InvalidParameter.PolicyNameExisted"
//  INVALIDPARAMETER_PRINCIPALERROR = "InvalidParameter.PrincipalError"
//  INVALIDPARAMETER_PRINCIPALQCSERROR = "InvalidParameter.PrincipalQcsError"
//  INVALIDPARAMETER_PRINCIPALQCSNOTEXIST = "InvalidParameter.PrincipalQcsNotExist"
//  INVALIDPARAMETER_PRINCIPALSERVICENOTEXIST = "InvalidParameter.PrincipalServiceNotExist"
//  INVALIDPARAMETER_RESERVEDTAGKEY = "InvalidParameter.ReservedTagKey"
//  INVALIDPARAMETER_RESOURCECONTENTERROR = "InvalidParameter.ResourceContentError"
//  INVALIDPARAMETER_RESOURCEERROR = "InvalidParameter.ResourceError"
//  INVALIDPARAMETER_RESOURCEPROJECTERROR = "InvalidParameter.ResourceProjectError"
//  INVALIDPARAMETER_RESOURCEQCSERROR = "InvalidParameter.ResourceQcsError"
//  INVALIDPARAMETER_RESOURCEREGIONERROR = "InvalidParameter.ResourceRegionError"
//  INVALIDPARAMETER_RESOURCESERVICENOTEXIST = "InvalidParameter.ResourceServiceNotExist"
//  INVALIDPARAMETER_RESOURCEUINERROR = "InvalidParameter.ResourceUinError"
//  INVALIDPARAMETER_STATEMENTERROR = "InvalidParameter.StatementError"
//  INVALIDPARAMETER_UNSUPPORTEDSERVICE = "InvalidParameter.UnsupportedService"
//  INVALIDPARAMETER_VERSIONERROR = "InvalidParameter.VersionError"
//  INVALIDPARAMETERVALUE_POLICYCONTENTINVALID = "InvalidParameterValue.PolicyContentInvalid"
//  LIMITEXCEEDED_TAGPOLICY = "LimitExceeded.TagPolicy"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  FAILEDOPERATION = "FailedOperation"
//  FAILEDOPERATION_ORGANIZATIONPOLICYINUSED = "FailedOperation.OrganizationPolicyInUsed"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTDISABLED = "FailedOperation.OrganizationPolicyIsNotDisabled"
//  FAILEDOPERATION_POLICYENABLEINVALID = "FailedOperation.PolicyEnableInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_APPLYNOTEXIST = "ResourceNotFound.ApplyNotExist"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ATTACHMENTFULL = "InvalidParameter.AttachmentFull"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_INTERFACENOTEXIST = "InvalidParameter.InterfaceNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_ORGANIZATIONNODENOTEXIST = "InvalidParameter.OrganizationNodeNotExist"
//  INVALIDPARAMETER_ORGANIZATIONNOTEXIST = "InvalidParameter.OrganizationNotExist"
//  MISSINGPARAMETER = "MissingParameter"
//  RESOURCENOTFOUND_CHANGEPERMISSIONNOTEXIST = "ResourceNotFound.ChangePermissionNotExist"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_INVITATIONNOTEXIST = "ResourceNotFound.InvitationNotExist"
//  RESOURCENOTFOUND_MEMBEREVENTNOTEXIST = "ResourceNotFound.MemberEventNotExist"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_MEMBERNOTEXIST = "ResourceNotFound.MemberNotExist"
//  RESOURCENOTFOUND_MEMBEROPERATEPROCESSNOTEXIST = "ResourceNotFound.MemberOperateProcessNotExist"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_NODENOTEXIST = "ResourceNotFound.NodeNotExist"
//  RESOURCENOTFOUND_NOTFOUND = "ResourceNotFound.NotFound"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICEASSIGNNOTEXIST = "ResourceNotFound.OrganizationServiceAssignNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONSERVICENOTEXIST = "ResourceNotFound.OrganizationServiceNotExist"
//  RESOURCENOTFOUND_POLICYIDNOTFOUND = "ResourceNotFound.PolicyIdNotFound"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  RESOURCENOTFOUND_RESOURCETYPENOTEXIST = "ResourceNotFound.ResourceTypeNotExist"
//  RESOURCENOTFOUND_SERVICEROLENOTEXIST = "ResourceNotFound.ServiceRoleNotExist"
//  RESOURCENOTFOUND_SHARERESOURCEMEMBERNOTEXIST = "ResourceNotFound.ShareResourceMemberNotExist"
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
//  FAILEDOPERATION_ORGANIZATIONDETACHLASTPOLICYERROR = "FailedOperation.OrganizationDetachLastPolicyError"
//  FAILEDOPERATION_ORGANIZATIONDETACHPOLICYERROR = "FailedOperation.OrganizationDetachPolicyError"
//  FAILEDOPERATION_ORGANIZATIONPOLICYISNOTENABLED = "FailedOperation.OrganizationPolicyIsNotEnabled"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTMANAGER = "InvalidParameter.OrganizationMemberNotManager"
//  INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//  INVALIDPARAMETER_POLICYIDNOTEXIST = "InvalidParameter.PolicyIdNotExist"
//  RESOURCENOTFOUND_POLICYNOTEXIST = "ResourceNotFound.PolicyNotExist"
//  UNAUTHORIZEDOPERATION = "UnauthorizedOperation"
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
//  FAILEDOPERATION_AUTHINFOEMPTY = "FailedOperation.AuthInfoEmpty"
//  FAILEDOPERATION_AUTHNOTENTERPRISE = "FailedOperation.AuthNotEnterprise"
//  FAILEDOPERATION_CREATEACCOUNT = "FailedOperation.CreateAccount"
//  FAILEDOPERATION_CREATEBILLINGPERMISSIONERR = "FailedOperation.CreateBillingPermissionErr"
//  FAILEDOPERATION_CREATEMEMBERAUTHOVERLIMIT = "FailedOperation.CreateMemberAuthOverLimit"
//  FAILEDOPERATION_CREATERECORDALREADYSUCCESS = "FailedOperation.CreateRecordAlreadySuccess"
//  FAILEDOPERATION_CREATERECORDNOTEXIST = "FailedOperation.CreateRecordNotExist"
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_GETAUTHINFO = "FailedOperation.GetAuthInfo"
//  FAILEDOPERATION_MEMBERNAMEUSED = "FailedOperation.MemberNameUsed"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONMEMBERNAMEUSED = "FailedOperation.OrganizationMemberNameUsed"
//  FAILEDOPERATION_ORGANIZATIONNODENOTEXIST = "FailedOperation.OrganizationNodeNotExist"
//  FAILEDOPERATION_ORGANIZATIONPERMISSIONILLEGAL = "FailedOperation.OrganizationPermissionIllegal"
//  FAILEDOPERATION_ORGANIZATIONPOLICYILLEGAL = "FailedOperation.OrganizationPolicyIllegal"
//  FAILEDOPERATION_PARTNERMANAGEMENTERR = "FailedOperation.PartnerManagementErr"
//  FAILEDOPERATION_PAYUINILLEGAL = "FailedOperation.PayUinIllegal"
//  FAILEDOPERATION_TAGRESOURCESERROR = "FailedOperation.TagResourcesError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_ORGANIZATIONMEMBERNOTEXIST = "InvalidParameter.OrganizationMemberNotExist"
//  INVALIDPARAMETER_TAGERROR = "InvalidParameter.TagError"
//  LIMITEXCEEDED_CREATEMEMBEROVERLIMIT = "LimitExceeded.CreateMemberOverLimit"
//  LIMITEXCEEDED_ORGANIZATIONMEMBEROVERLIMIT = "LimitExceeded.OrganizationMemberOverLimit"
//  RESOURCENOTFOUND_ORGANIZATIONAUTHRELATIONNOTEXIST = "ResourceNotFound.OrganizationAuthRelationNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNODENOTEXIST = "ResourceNotFound.OrganizationNodeNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_ABNORMALFINANCIALSTATUSOFADMIN = "UnsupportedOperation.AbnormalFinancialStatusOfAdmin"
//  UNSUPPORTEDOPERATION_ADDDELEGATEPAYERNOTALLOW = "UnsupportedOperation.AddDelegatePayerNotAllow"
//  UNSUPPORTEDOPERATION_ADDDISCOUNTINHERITNOTALLOW = "UnsupportedOperation.AddDiscountInheritNotAllow"
//  UNSUPPORTEDOPERATION_EXISTEDAGENT = "UnsupportedOperation.ExistedAgent"
//  UNSUPPORTEDOPERATION_EXISTEDCLIENT = "UnsupportedOperation.ExistedClient"
//  UNSUPPORTEDOPERATION_INCONSISTENTUSERTYPES = "UnsupportedOperation.InconsistentUserTypes"
//  UNSUPPORTEDOPERATION_MANAGEMENTSYSTEMERROR = "UnsupportedOperation.ManagementSystemError"
//  UNSUPPORTEDOPERATION_MEMBERACCOUNTARREARS = "UnsupportedOperation.MemberAccountArrears"
//  UNSUPPORTEDOPERATION_MEMBERDISCOUNTINHERITEXISTED = "UnsupportedOperation.MemberDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_MEMBEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.MemberExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_MEMBERISAGENT = "UnsupportedOperation.MemberIsAgent"
//  UNSUPPORTEDOPERATION_ORDERINPROGRESSEXISTED = "UnsupportedOperation.OrderInProgressExisted"
//  UNSUPPORTEDOPERATION_OWNERDISCOUNTINHERITEXISTED = "UnsupportedOperation.OwnerDiscountInheritExisted"
//  UNSUPPORTEDOPERATION_PAYERARREARSANDNOCREDITACCOUNT = "UnsupportedOperation.PayerArrearsAndNoCreditAccount"
//  UNSUPPORTEDOPERATION_PAYEREXISTACCOUNTLEVELDISCOUNTINHERIT = "UnsupportedOperation.PayerExistAccountLevelDiscountInherit"
//  UNSUPPORTEDOPERATION_SECONDARYDISTRIBUTORSUBCLIENTEXISTED = "UnsupportedOperation.SecondaryDistributorSubClientExisted"
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
//  FAILEDOPERATION_CREATEROLE = "FailedOperation.CreateRole"
//  FAILEDOPERATION_ORGANIZATIONIDENTITYPOLICYERROR = "FailedOperation.OrganizationIdentityPolicyError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
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
//  FAILEDOPERATION_MEMBERIDENTITYAUTHUSED = "FailedOperation.MemberIdentityAuthUsed"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONIDENTITYNOTEXIST = "ResourceNotFound.OrganizationIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
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
//  FAILEDOPERATION_CHECKACCOUNTPHONEBINDLIMIT = "FailedOperation.CheckAccountPhoneBindLimit"
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_MEMBEREMAILEXIST = "FailedOperation.MemberEmailExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_EMAILBINDOVERLIMIT = "LimitExceeded.EmailBindOverLimit"
//  LIMITEXCEEDED_PHONENUMBOUND = "LimitExceeded.PhoneNumBound"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION_ACCOUNTALREADYREGISTER = "FailedOperation.AccountAlreadyRegister"
//  FAILEDOPERATION_BINDEMAILLINKEXPIRED = "FailedOperation.BindEmailLinkExpired"
//  FAILEDOPERATION_BINDEMAILLINKINVALID = "FailedOperation.BindEmailLinkInvalid"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  FAILEDOPERATION_MEMBERBINDEMAILERROR = "FailedOperation.MemberBindEmailError"
//  FAILEDOPERATION_MEMBERBINDPHONEERROR = "FailedOperation.MemberBindPhoneError"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETER_CODEERROR = "InvalidParameter.CodeError"
//  INVALIDPARAMETER_CODEEXPIRED = "InvalidParameter.CodeExpired"
//  INVALIDPARAMETER_INVALIDEMAIL = "InvalidParameter.InvalidEmail"
//  INVALIDPARAMETER_PASSWORDILLEGAL = "InvalidParameter.PasswordIllegal"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
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
//  FAILEDOPERATION_CHECKMAILACCOUNT = "FailedOperation.CheckMailAccount"
//  FAILEDOPERATION_EMAILALREADYUSED = "FailedOperation.EmailAlreadyUsed"
//  FAILEDOPERATION_EMAILBINDRECORDINVALID = "FailedOperation.EmailBindRecordInvalid"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  LIMITEXCEEDED_UPDATEEMAILBINDOVERLIMIT = "LimitExceeded.UpdateEmailBindOverLimit"
//  RESOURCENOTFOUND_EMAILBINDRECORDNOTEXIST = "ResourceNotFound.EmailBindRecordNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION_CREATEPOLICY = "FailedOperation.CreatePolicy"
//  FAILEDOPERATION_MEMBERPOLICYNAMEEXIST = "FailedOperation.MemberPolicyNameExist"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERIDENTITYNOTEXIST = "ResourceNotFound.MemberIdentityNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION = "UnsupportedOperation"
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
//  FAILEDOPERATION_DELETEPOLICY = "FailedOperation.DeletePolicy"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_MEMBERPOLICYNOTEXIST = "ResourceNotFound.MemberPolicyNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
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
//  FAILEDOPERATION_DISABLEQUITSELFCREATEDORGANIZATION = "FailedOperation.DisableQuitSelfCreatedOrganization"
//  FAILEDOPERATION_MEMBEREXISTDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberExistDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERISDELEGATEPAYERNOTALLOWDELETE = "FailedOperation.MemberIsDelegatePayerNotAllowDelete"
//  FAILEDOPERATION_MEMBERSHARERESOURCE = "FailedOperation.MemberShareResource"
//  FAILEDOPERATION_OPERATEBILLINGPERMISSIONERR = "FailedOperation.OperateBillingPermissionErr"
//  FAILEDOPERATION_ORGANIZATIONAUTHMANAGENOTALLOWDELETE = "FailedOperation.OrganizationAuthManageNotAllowDelete"
//  FAILEDOPERATION_QUITESHAREUNIT = "FailedOperation.QuiteShareUnit"
//  FAILEDOPERATION_SHARERESOURCEMEMBERINUSE = "FailedOperation.ShareResourceMemberInUse"
//  INTERNALERROR = "InternalError"
//  INVALIDPARAMETER = "InvalidParameter"
//  RESOURCENOTFOUND_ORGANIZATIONMEMBERNOTEXIST = "ResourceNotFound.OrganizationMemberNotExist"
//  RESOURCENOTFOUND_ORGANIZATIONNOTEXIST = "ResourceNotFound.OrganizationNotExist"
//  UNSUPPORTEDOPERATION_CREATEMEMBERNOTALLOWQUIT = "UnsupportedOperation.CreateMemberNotAllowQuit"
//  UNSUPPORTEDOPERATION_MEMBEREXISTOPERATEPROCESSNOTALLOWDELETE = "UnsupportedOperation.MemberExistOperateProcessNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBEREXISTSERVICENOTALLOWDELETE = "UnsupportedOperation.MemberExistServiceNotAllowDelete"
//  UNSUPPORTEDOPERATION_MEMBERNOPAYMENT = "UnsupportedOperation.MemberNoPayment"
//  UNSUPPORTEDOPERATION_MEMBERNOTALLOWQUIT = "UnsupportedOperation.MemberNotAllowQuit"
func (c *Client) QuitOrganization(request *QuitOrganizationRequest) (response *QuitOrganizationResponse, err error) {
	if request == nil {
		request = NewQuitOrganizationRequest()
	}


	response = NewQuitOrganizationResponse()
	err = c.Send(request, response)
	return
}
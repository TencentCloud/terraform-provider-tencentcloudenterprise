package v20210331

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2021-03-31"

type Client struct {
	common.Client
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithCredential(credential).
		WithProfile(clientProfile)
	return
}

func NewGetZoneSAMLServiceProviderInfoRequest() (request *GetZoneSAMLServiceProviderInfoRequest) {
	request = &GetZoneSAMLServiceProviderInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetZoneSAMLServiceProviderInfo")

	return
}

func NewGetZoneSAMLServiceProviderInfoResponse() (response *GetZoneSAMLServiceProviderInfoResponse) {
	response = &GetZoneSAMLServiceProviderInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetZoneSAMLServiceProviderInfo
// 查询SAML服务提供商配置信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_SAMLMETADATADOCUMENTGETFAILED = "FailedOperation.SAMLMetadataDocumentGetFailed"
//	INTERNALERROR = "InternalError"
//	RESOURCENOTFOUND_SAMLSERVICEPROVIDERNOTFOUND = "ResourceNotFound.SAMLServiceProviderNotFound"
func (c *Client) GetZoneSAMLServiceProviderInfo(request *GetZoneSAMLServiceProviderInfoRequest) (
	response *GetZoneSAMLServiceProviderInfoResponse, err error) {
	if request == nil {
		request = NewGetZoneSAMLServiceProviderInfoRequest()
	}

	response = NewGetZoneSAMLServiceProviderInfoResponse()
	err = c.Send(request, response)
	return
}

// NewGetExternalSAMLIdentityProviderRequest ...
func NewGetExternalSAMLIdentityProviderRequest() (request *GetExternalSAMLIdentityProviderRequest) {
	request = &GetExternalSAMLIdentityProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetExternalSAMLIdentityProvider")

	return
}

// NewGetExternalSAMLIdentityProviderResponse ...
func NewGetExternalSAMLIdentityProviderResponse() (response *GetExternalSAMLIdentityProviderResponse) {
	response = &GetExternalSAMLIdentityProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetExternalSAMLIdentityProvider ...
func (c *Client) GetExternalSAMLIdentityProvider(request *GetExternalSAMLIdentityProviderRequest) (response *GetExternalSAMLIdentityProviderResponse, err error) {
	if request == nil {
		request = NewGetExternalSAMLIdentityProviderRequest()
	}

	response = NewGetExternalSAMLIdentityProviderResponse()
	err = c.Send(request, response)
	return
}

// SetExternalSAMLIdentityProvider
// 配置SAML身份提供商信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_DECODEMETADATADOCUMENTFAILED = "FailedOperation.DecodeMetadataDocumentFailed"
//	FAILEDOPERATION_METADATACOSPARSINGFAILED = "FailedOperation.MetadataCosParsingFailed"
//	FAILEDOPERATION_SAMLSERVICEPROVIDERCREATEFAILED = "FailedOperation.SAMLServiceProviderCreateFailed"
//	FAILEDOPERATION_UPLOADMETADATAFAILED = "FailedOperation.UploadMetadataFailed"
//	FAILEDOPERATION_X509CERTIFICATEPARSINGFAILED = "FailedOperation.X509CertificateParsingFailed"
//	FAILEDOPERATION_XMLDATAUNMARSHALFAILED = "FailedOperation.XMLDataUnmarshalFailed"
//	INTERNALERROR = "InternalError"
//	INVALIDPARAMETERVALUE_SSOSTATUSINVALID = "InvalidParameterValue.SSoStatusInvalid"
//	INVALIDPARAMETERVALUE_X509CERTIFICATEFORMATERROR = "InvalidParameterValue.X509CertificateFormatError"
func (c *Client) SetExternalSAMLIdentityProvider(request *SetExternalSAMLIdentityProviderRequest) (response *SetExternalSAMLIdentityProviderResponse, err error) {
	if request == nil {
		request = NewSetExternalSAMLIdentityProviderRequest()
	}

	response = NewSetExternalSAMLIdentityProviderResponse()
	err = c.Send(request, response)
	return
}

func NewSetExternalSAMLIdentityProviderRequest() (request *SetExternalSAMLIdentityProviderRequest) {
	request = &SetExternalSAMLIdentityProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "SetExternalSAMLIdentityProvider")

	return
}

func NewSetExternalSAMLIdentityProviderResponse() (response *SetExternalSAMLIdentityProviderResponse) {
	response = &SetExternalSAMLIdentityProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

func NewClearExternalSAMLIdentityProviderRequest() (request *ClearExternalSAMLIdentityProviderRequest) {
	request = &ClearExternalSAMLIdentityProviderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "ClearExternalSAMLIdentityProvider")

	return
}

func NewClearExternalSAMLIdentityProviderResponse() (response *ClearExternalSAMLIdentityProviderResponse) {
	response = &ClearExternalSAMLIdentityProviderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

func (c *Client) ClearExternalSAMLIdentityProvider(request *ClearExternalSAMLIdentityProviderRequest) (response *ClearExternalSAMLIdentityProviderResponse, err error) {
	if request == nil {
		request = NewClearExternalSAMLIdentityProviderRequest()
	}

	response = NewClearExternalSAMLIdentityProviderResponse()
	err = c.Send(request, response)
	return
}

// UpdateGroup
// 修改用户组信息
//
// 可能返回的错误码:
//  FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//  FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//  FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//  FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//  INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//  INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
//	if request == nil {
//		request = NewUpdateGroupRequest()
//	}
//
//	response = NewUpdateGroupResponse()
//	err = c.Send(request, response)
//	return
//}

func NewCreateUserRequest() (request *CreateUserRequest) {
	request = &CreateUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "CreateUser")

	return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
	response = &CreateUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateUser
// 创建用户
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_USEROVERUPPERLIMIT = "FailedOperation.UserOverUpperLimit"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//	INVALIDPARAMETER_USERNAMEALREADYEXISTS = "InvalidParameter.UsernameAlreadyExists"
//	INVALIDPARAMETER_USERNAMEFORMATERROR = "InvalidParameter.UsernameFormatError"
//	LIMITEXCEEDED_CREATEUSERLIMITEXCEEDED = "LimitExceeded.CreateUserLimitExceeded"
func (c *Client) CreateUserWithContext(request *CreateUserRequest) (response *CreateUserResponse, err error) {
	if request == nil {
		request = NewCreateUserRequest()
	}

	response = NewCreateUserResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserSyncProvisioningRequest() (request *CreateUserSyncProvisioningRequest) {
	request = &CreateUserSyncProvisioningRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "CreateUserSyncProvisioning")

	return
}

func NewCreateUserSyncProvisioningResponse() (response *CreateUserSyncProvisioningResponse) {
	response = &CreateUserSyncProvisioningResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateUserSyncProvisioning
// 创建子用户同步任务
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//	FAILEDOPERATION_USERPROVISIONINGALREADYEXISTS = "FailedOperation.UserProvisioningAlreadyExists"
//	FAILEDOPERATION_USERPROVISIONINGOVERLIMIT = "FailedOperation.UserProvisioningOverLimit"
//	INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//	LIMITEXCEEDED_CREATEUSERSYNCPROVISIONINGLIMITEXCEEDED = "LimitExceeded.CreateUserSyncProvisioningLimitExceeded"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateUserSyncProvisioning(request *CreateUserSyncProvisioningRequest) (response *CreateUserSyncProvisioningResponse, err error) {
	if request == nil {
		request = NewCreateUserSyncProvisioningRequest()
	}

	response = NewCreateUserSyncProvisioningResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserRequest() (request *GetUserRequest) {
	request = &GetUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetUser")

	return
}

func NewGetUserResponse() (response *GetUserResponse) {
	response = &GetUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetUser
// 查询用户信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) GetUser(request *GetUserRequest) (response *GetUserResponse, err error) {
	if request == nil {
		request = NewGetUserRequest()
	}

	response = NewGetUserResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUserStatusRequest() (request *UpdateUserStatusRequest) {
	request = &UpdateUserStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "UpdateUserStatus")

	return
}

func NewUpdateUserStatusResponse() (response *UpdateUserStatusResponse) {
	response = &UpdateUserStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdateUserStatus
// 修改用户状态
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//	FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUserStatus(request *UpdateUserStatusRequest) (response *UpdateUserStatusResponse, err error) {
	if request == nil {
		request = NewUpdateUserStatusRequest()
	}

	response = NewUpdateUserStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRoleAssignmentRequest() (request *CreateRoleAssignmentRequest) {
	request = &CreateRoleAssignmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "CreateRoleAssignment")

	return
}

func NewCreateRoleAssignmentResponse() (response *CreateRoleAssignmentResponse) {
	response = &CreateRoleAssignmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateRoleAssignment
// 在成员账号上授权
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//	FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONALREADYEXIST = "FailedOperation.RoleConfigurationAuthorizationAlreadyExist"
//	FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONOVERLIMIT = "FailedOperation.RoleConfigurationAuthorizationOverLimit"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//	INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//	LIMITEXCEEDED_CREATEROLEASSIGNMENTLIMITEXCEEDED = "LimitExceeded.CreateRoleAssignmentLimitExceeded"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) CreateRoleAssignment(request *CreateRoleAssignmentRequest) (response *CreateRoleAssignmentResponse, err error) {
	if request == nil {
		request = NewCreateRoleAssignmentRequest()
	}

	response = NewCreateRoleAssignmentResponse()
	err = c.Send(request, response)
	return
}

func NewListRoleAssignmentsRequest() (request *ListRoleAssignmentsRequest) {
	request = &ListRoleAssignmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "ListRoleAssignments")

	return
}

func NewListRoleAssignmentsResponse() (response *ListRoleAssignmentsResponse) {
	response = &ListRoleAssignmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// ListRoleAssignments
// 查询授权列表
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_NEXTTOKENINVALID = "InvalidParameter.NextTokenInvalid"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
func (c *Client) ListRoleAssignments(request *ListRoleAssignmentsRequest) (response *ListRoleAssignmentsResponse, err error) {
	if request == nil {
		request = NewListRoleAssignmentsRequest()
	}

	response = NewListRoleAssignmentsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRoleAssignmentRequest() (request *DeleteRoleAssignmentRequest) {
	request = &DeleteRoleAssignmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "DeleteRoleAssignment")

	return
}

func NewDeleteRoleAssignmentResponse() (response *DeleteRoleAssignmentResponse) {
	response = &DeleteRoleAssignmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteRoleAssignment
// 移除成员账号上的授权
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_ROLECONFIGURATIONAUTHORIZATIONNOTFOUND = "ResourceNotFound.RoleConfigurationAuthorizationNotFound"
func (c *Client) DeleteRoleAssignment(request *DeleteRoleAssignmentRequest) (response *DeleteRoleAssignmentResponse, err error) {
	if request == nil {
		request = NewDeleteRoleAssignmentRequest()
	}
	response = NewDeleteRoleAssignmentResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRoleConfigurationRequest() (request *DeleteRoleConfigurationRequest) {
	request = &DeleteRoleConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "DeleteRoleConfiguration")

	return
}

func NewDeleteRoleConfigurationResponse() (response *DeleteRoleConfigurationResponse) {
	response = &DeleteRoleConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteRoleConfiguration
// 删除权限配置信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGALREADYDEPLOYED = "FailedOperation.RoleConfigurationProvisioningAlreadyDeployed"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//	INVALIDPARAMETER_ROLEPOLICYALREADYEXIST = "InvalidParameter.RolePolicyAlreadyExist"
func (c *Client) DeleteRoleConfiguration(request *DeleteRoleConfigurationRequest) (response *DeleteRoleConfigurationResponse, err error) {
	if request == nil {
		request = NewDeleteRoleConfigurationRequest()
	}

	response = NewDeleteRoleConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewDismantleRoleConfigurationRequest() (request *DismantleRoleConfigurationRequest) {
	request = &DismantleRoleConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "DismantleRoleConfiguration")

	return
}

func NewDismantleRoleConfigurationResponse() (response *DismantleRoleConfigurationResponse) {
	response = &DismantleRoleConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DismantleRoleConfiguration
// 解除权限配置在成员账号上的部署
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ORGANIZATIONMEMBERNOTEXIST = "FailedOperation.OrganizationMemberNotExist"
//	FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//	FAILEDOPERATION_ROLECONFIGURATIONPROVISIONINGSTATUSERROR = "FailedOperation.RoleConfigurationProvisioningStatusError"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_ROLECONFIGURATIONPROVISIONINGNOTFOUND = "ResourceNotFound.RoleConfigurationProvisioningNotFound"
func (c *Client) DismantleRoleConfiguration(request *DismantleRoleConfigurationRequest) (response *DismantleRoleConfigurationResponse, err error) {
	if request == nil {
		request = NewDismantleRoleConfigurationRequest()
	}

	response = NewDismantleRoleConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
	request = &CreateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "CreateGroup")

	return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
	response = &CreateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateGroup
// 创建用户组
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_GROUPOVERUPPERLIMIT = "FailedOperation.GroupOverUpperLimit"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//	INVALIDPARAMETER_GROUPNAMEFORMATERROR = "InvalidParameter.GroupNameFormatError"
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
	if request == nil {
		request = NewCreateGroupRequest()
	}

	response = NewCreateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetGroupRequest() (request *GetGroupRequest) {
	request = &GetGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetGroup")

	return
}

func NewGetGroupResponse() (response *GetGroupResponse) {
	response = &GetGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetGroup
// 查询用户组信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) GetGroup(request *GetGroupRequest) (response *GetGroupResponse, err error) {
	if request == nil {
		request = NewGetGroupRequest()
	}

	response = NewGetGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGroupRequest() (request *UpdateGroupRequest) {
	request = &UpdateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "UpdateGroup")

	return
}

func NewUpdateGroupResponse() (response *UpdateGroupResponse) {
	response = &UpdateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdateGroup
// 修改用户组信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_GROUPNAMEALREADYEXISTS = "InvalidParameter.GroupNameAlreadyExists"
//	INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) UpdateGroup(request *UpdateGroupRequest) (response *UpdateGroupResponse, err error) {
	if request == nil {
		request = NewUpdateGroupRequest()
	}

	response = NewUpdateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
	request = &DeleteGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "DeleteGroup")

	return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
	response = &DeleteGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteGroup
// 删除用户组
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DELETEGROUPNOTALLOWEXISTUSER = "FailedOperation.DeleteGroupNotAllowExistUser"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_MANUALGROUPNOTDELETE = "FailedOperation.ManualGroupNotDelete"
//	FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTDELETE = "FailedOperation.SynchronizedGroupNotDelete"
//	FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
	if request == nil {
		request = NewDeleteGroupRequest()
	}

	response = NewDeleteGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetTaskStatusRequest() (request *GetTaskStatusRequest) {
	request = &GetTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetTaskStatus")

	return
}

func NewGetTaskStatusResponse() (response *GetTaskStatusResponse) {
	response = &GetTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetTaskStatus
// 查询异步任务的状态
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_ROLECONFIGURATIONTASKNOTFOUND = "ResourceNotFound.RoleConfigurationTaskNotFound"
func (c *Client) GetTaskStatus(request *GetTaskStatusRequest) (response *GetTaskStatusResponse, err error) {
	if request == nil {
		request = NewGetTaskStatusRequest()
	}

	response = NewGetTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRoleConfigurationRequest() (request *CreateRoleConfigurationRequest) {
	request = &CreateRoleConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "CreateRoleConfiguration")

	return
}

func NewCreateRoleConfigurationResponse() (response *CreateRoleConfigurationResponse) {
	response = &CreateRoleConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateRoleConfiguration
// 创建权限配置
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_CONFIGURATIONOVERUPPERLIMIT = "FailedOperation.ConfigurationOverUpperLimit"
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_CONFIGURATIONNAMEALREADYEXISTS = "InvalidParameter.ConfigurationNameAlreadyExists"
//	INVALIDPARAMETER_CONFIGURATIONNAMEFORMATERROR = "InvalidParameter.ConfigurationNameFormatError"
func (c *Client) CreateRoleConfiguration(request *CreateRoleConfigurationRequest) (response *CreateRoleConfigurationResponse, err error) {
	if request == nil {
		request = NewCreateRoleConfigurationRequest()
	}

	response = NewCreateRoleConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewGetRoleConfigurationRequest() (request *GetRoleConfigurationRequest) {
	request = &GetRoleConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetRoleConfiguration")

	return
}

func NewGetRoleConfigurationResponse() (response *GetRoleConfigurationResponse) {
	response = &GetRoleConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetRoleConfiguration
// 查询权限配置信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) GetRoleConfiguration(request *GetRoleConfigurationRequest) (response *GetRoleConfigurationResponse, err error) {
	if request == nil {
		request = NewGetRoleConfigurationRequest()
	}

	response = NewGetRoleConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRoleConfigurationRequest() (request *UpdateRoleConfigurationRequest) {
	request = &UpdateRoleConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "UpdateRoleConfiguration")

	return
}

func NewUpdateRoleConfigurationResponse() (response *UpdateRoleConfigurationResponse) {
	response = &UpdateRoleConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdateRoleConfiguration
// 修改权限配置信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) UpdateRoleConfiguration(request *UpdateRoleConfigurationRequest) (response *UpdateRoleConfigurationResponse, err error) {
	if request == nil {
		request = NewUpdateRoleConfigurationRequest()
	}

	response = NewUpdateRoleConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewAddPermissionPolicyToRoleConfigurationRequest() (request *AddPermissionPolicyToRoleConfigurationRequest) {
	request = &AddPermissionPolicyToRoleConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "AddPermissionPolicyToRoleConfiguration")

	return
}

func NewAddPermissionPolicyToRoleConfigurationResponse() (response *AddPermissionPolicyToRoleConfigurationResponse) {
	response = &AddPermissionPolicyToRoleConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// AddPermissionPolicyToRoleConfiguration
// 为权限配置添加策略
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_CUSTOMPOLICYOVERUPPERLIMIT = "FailedOperation.CustomPolicyOverUpperLimit"
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_SYSTEMPOLICYOVERUPPERLIMIT = "FailedOperation.SystemPolicyOverUpperLimit"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_BINDPOLICYNAMENOTALLOWED = "InvalidParameter.BindPolicyNameNotAllowed"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	INVALIDPARAMETER_POLICYDOCUMENTEMPTY = "InvalidParameter.PolicyDocumentEmpty"
//	INVALIDPARAMETER_POLICYNAMEALREADYEXISTS = "InvalidParameter.PolicyNameAlreadyExists"
//	INVALIDPARAMETER_POLICYNAMESIZEOVERUPPERLIMIT = "InvalidParameter.PolicyNameSizeOverUpperLimit"
//	INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) AddPermissionPolicyToRoleConfiguration(request *AddPermissionPolicyToRoleConfigurationRequest) (response *AddPermissionPolicyToRoleConfigurationResponse, err error) {
	if request == nil {
		request = NewAddPermissionPolicyToRoleConfigurationRequest()
	}

	response = NewAddPermissionPolicyToRoleConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewListPermissionPoliciesInRoleConfigurationRequest() (request *ListPermissionPoliciesInRoleConfigurationRequest) {
	request = &ListPermissionPoliciesInRoleConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "ListPermissionPoliciesInRoleConfiguration")

	return
}

func NewListPermissionPoliciesInRoleConfigurationResponse() (response *ListPermissionPoliciesInRoleConfigurationResponse) {
	response = &ListPermissionPoliciesInRoleConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// ListPermissionPoliciesInRoleConfiguration
// 获取权限配置中的策略列表
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
func (c *Client) ListPermissionPoliciesInRoleConfiguration(request *ListPermissionPoliciesInRoleConfigurationRequest) (response *ListPermissionPoliciesInRoleConfigurationResponse, err error) {
	if request == nil {
		request = NewListPermissionPoliciesInRoleConfigurationRequest()
	}

	response = NewListPermissionPoliciesInRoleConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewRemovePermissionPolicyFromRoleConfigurationRequest() (request *RemovePermissionPolicyFromRoleConfigurationRequest) {
	request = &RemovePermissionPolicyFromRoleConfigurationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "RemovePermissionPolicyFromRoleConfiguration")

	return
}

func NewRemovePermissionPolicyFromRoleConfigurationResponse() (response *RemovePermissionPolicyFromRoleConfigurationResponse) {
	response = &RemovePermissionPolicyFromRoleConfigurationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// RemovePermissionPolicyFromRoleConfiguration
// 为权限配置移除策略
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_PARAMERROR = "InvalidParameter.ParamError"
//	INVALIDPARAMETER_ROLECONFIGURATIONNOTEXIST = "InvalidParameter.RoleConfigurationNotExist"
//	INVALIDPARAMETER_ROLEPOLICYNOTEXIST = "InvalidParameter.RolePolicyNotExist"
func (c *Client) RemovePermissionPolicyFromRoleConfiguration(request *RemovePermissionPolicyFromRoleConfigurationRequest) (response *RemovePermissionPolicyFromRoleConfigurationResponse, err error) {
	if request == nil {
		request = NewRemovePermissionPolicyFromRoleConfigurationRequest()
	}

	response = NewRemovePermissionPolicyFromRoleConfigurationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSCIMCredentialRequest() (request *CreateSCIMCredentialRequest) {
	request = &CreateSCIMCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "CreateSCIMCredential")

	return
}

func NewCreateSCIMCredentialResponse() (response *CreateSCIMCredentialResponse) {
	response = &CreateSCIMCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// CreateSCIMCredential
// 创建SCIM密钥
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_SCIMCREDENTIALGENERATEERROR = "FailedOperation.ScimCredentialGenerateError"
//	LIMITEXCEEDED_SCIMCREDENTIALLIMITEXCEEDED = "LimitExceeded.ScimCredentialLimitExceeded"
func (c *Client) CreateSCIMCredential(request *CreateSCIMCredentialRequest) (response *CreateSCIMCredentialResponse, err error) {
	if request == nil {
		request = NewCreateSCIMCredentialRequest()
	}

	response = NewCreateSCIMCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewListSCIMCredentialsRequest() (request *ListSCIMCredentialsRequest) {
	request = &ListSCIMCredentialsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "ListSCIMCredentials")

	return
}

func NewListSCIMCredentialsResponse() (response *ListSCIMCredentialsResponse) {
	response = &ListSCIMCredentialsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// ListSCIMCredentials
// 查询用户SCIM密钥列表
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) ListSCIMCredentials(request *ListSCIMCredentialsRequest) (response *ListSCIMCredentialsResponse, err error) {
	if request == nil {
		request = NewListSCIMCredentialsRequest()
	}

	response = NewListSCIMCredentialsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSCIMCredentialRequest() (request *DeleteSCIMCredentialRequest) {
	request = &DeleteSCIMCredentialRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("organization", APIVersion, "DeleteSCIMCredential")

	return
}

func NewDeleteSCIMCredentialResponse() (response *DeleteSCIMCredentialResponse) {
	response = &DeleteSCIMCredentialResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteSCIMCredential
// 删除SCIM密钥
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_SCIMCREDENTIALNOTFOUND = "InvalidParameter.ScimCredentialNotFound"
func (c *Client) DeleteSCIMCredential(request *DeleteSCIMCredentialRequest) (response *DeleteSCIMCredentialResponse, err error) {
	if request == nil {
		request = NewDeleteSCIMCredentialRequest()
	}

	response = NewDeleteSCIMCredentialResponse()
	err = c.Send(request, response)
	return
}

func NewGetSCIMSynchronizationStatusRequest() (request *GetSCIMSynchronizationStatusRequest) {
	request = &GetSCIMSynchronizationStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetSCIMSynchronizationStatus")

	return
}

func NewGetSCIMSynchronizationStatusResponse() (response *GetSCIMSynchronizationStatusResponse) {
	response = &GetSCIMSynchronizationStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetSCIMSynchronizationStatus
// 获取SCIM同步状态
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
func (c *Client) GetSCIMSynchronizationStatus(request *GetSCIMSynchronizationStatusRequest) (response *GetSCIMSynchronizationStatusResponse, err error) {
	if request == nil {
		request = NewGetSCIMSynchronizationStatusRequest()
	}

	response = NewGetSCIMSynchronizationStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSCIMSynchronizationStatusRequest() (request *UpdateSCIMSynchronizationStatusRequest) {
	request = &UpdateSCIMSynchronizationStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "UpdateSCIMSynchronizationStatus")

	return
}

func NewUpdateSCIMSynchronizationStatusResponse() (response *UpdateSCIMSynchronizationStatusResponse) {
	response = &UpdateSCIMSynchronizationStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdateSCIMSynchronizationStatus
// 启用/禁用用户SCIM同步
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_DBOPERATIONERROR = "FailedOperation.DBOperationError"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_SCIMSYNCSTATUSERROR = "InvalidParameter.ScimSyncStatusError"
func (c *Client) UpdateSCIMSynchronizationStatus(request *UpdateSCIMSynchronizationStatusRequest) (response *UpdateSCIMSynchronizationStatusResponse, err error) {
	if request == nil {
		request = NewUpdateSCIMSynchronizationStatusRequest()
	}

	response = NewUpdateSCIMSynchronizationStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUserRequest() (request *UpdateUserRequest) {
	request = &UpdateUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "UpdateUser")

	return
}

func NewUpdateUserResponse() (response *UpdateUserResponse) {
	response = &UpdateUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdateUser
// 修改用户信息
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_MANUALUSERNOTUPDATE = "FailedOperation.ManualUserNotUpdate"
//	FAILEDOPERATION_SYNCHRONIZEDUSERNOTUPDATE = "FailedOperation.SynchronizedUserNotUpdate"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_EMAILALREADYEXISTS = "InvalidParameter.EmailAlreadyExists"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) UpdateUser(request *UpdateUserRequest) (response *UpdateUserResponse, err error) {
	if request == nil {
		request = NewUpdateUserRequest()
	}

	response = NewUpdateUserResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
	request = &DeleteUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "DeleteUser")

	return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
	response = &DeleteUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteUser
// 删除用户
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_MANUALUSERNOTDELETE = "FailedOperation.ManualUserNotDelete"
//	FAILEDOPERATION_ROLECONFIGURATIONAUTHORIZATIONEXIST = "FailedOperation.RoleConfigurationAuthorizationExist"
//	FAILEDOPERATION_SYNCHRONIZEDUSERNOTDELETE = "FailedOperation.SynchronizedUserNotDelete"
//	FAILEDOPERATION_USERPROVISIONINGEXISTS = "FailedOperation.UserProvisioningExists"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_USERALREADYEXISTSGROUP = "InvalidParameter.UserAlreadyExistsGroup"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
	if request == nil {
		request = NewDeleteUserRequest()
	}

	response = NewDeleteUserResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserToGroupRequest() (request *AddUserToGroupRequest) {
	request = &AddUserToGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "AddUserToGroup")

	return
}

func NewAddUserToGroupResponse() (response *AddUserToGroupResponse) {
	response = &AddUserToGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// AddUserToGroup
// 为用户组添加用户
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_GROUPTYPEUSERTYPENOTMATCH = "FailedOperation.GroupTypeUserTypeNotMatch"
//	FAILEDOPERATION_GROUPUSERCOUNTOVERUPPERLIMIT = "FailedOperation.GroupUserCountOverUpperLimit"
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTADDUSER = "FailedOperation.SynchronizedGroupNotAddUser"
//	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//	FAILEDOPERATION_USERADDGROUPCOUNTOVERUPPERLIMIT = "FailedOperation.UserAddGroupCountOverUpperLimit"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//	INVALIDPARAMETER_GROUPUSERALREADYEXISTS = "InvalidParameter.GroupUserAlreadyExists"
//	LIMITEXCEEDED_ADDUSERTOGROUPLIMITEXCEEDED = "LimitExceeded.AddUserToGroupLimitExceeded"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) AddUserToGroup(request *AddUserToGroupRequest) (response *AddUserToGroupResponse, err error) {
	if request == nil {
		request = NewAddUserToGroupRequest()
	}

	response = NewAddUserToGroupResponse()
	err = c.Send(request, response)
	return
}

func NewListJoinedGroupsForUserRequest() (request *ListJoinedGroupsForUserRequest) {
	request = &ListJoinedGroupsForUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "ListJoinedGroupsForUser")

	return
}

func NewListJoinedGroupsForUserResponse() (response *ListJoinedGroupsForUserResponse) {
	response = &ListJoinedGroupsForUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// ListJoinedGroupsForUser
// 查询用户加入的用户组
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_USERNOTEXIST = "ResourceNotFound.UserNotExist"
func (c *Client) ListJoinedGroupsForUser(request *ListJoinedGroupsForUserRequest) (response *ListJoinedGroupsForUserResponse, err error) {
	if request == nil {
		request = NewListJoinedGroupsForUserRequest()
	}

	response = NewListJoinedGroupsForUserResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveUserFromGroupRequest() (request *RemoveUserFromGroupRequest) {
	request = &RemoveUserFromGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "RemoveUserFromGroup")

	return
}

func NewRemoveUserFromGroupResponse() (response *RemoveUserFromGroupResponse) {
	response = &RemoveUserFromGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// RemoveUserFromGroup
// 从用户组中移除用户
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_MANUALGROUPNOTUPDATE = "FailedOperation.ManualGroupNotUpdate"
//	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTREMOVEUSER = "FailedOperation.SynchronizedGroupNotRemoveUser"
//	FAILEDOPERATION_SYNCHRONIZEDGROUPNOTUPDATE = "FailedOperation.SynchronizedGroupNotUpdate"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	INVALIDPARAMETER_GROUPNOTEXIST = "InvalidParameter.GroupNotExist"
//	INVALIDPARAMETER_GROUPUSERNOTEXIST = "InvalidParameter.GroupUserNotExist"
//	LIMITEXCEEDED_REMOVEUSERFROMGROUPLIMITEXCEEDED = "LimitExceeded.RemoveUserFromGroupLimitExceeded"
func (c *Client) RemoveUserFromGroup(request *RemoveUserFromGroupRequest) (response *RemoveUserFromGroupResponse, err error) {
	if request == nil {
		request = NewRemoveUserFromGroupRequest()
	}

	response = NewRemoveUserFromGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetProvisioningTaskStatusRequest() (request *GetProvisioningTaskStatusRequest) {
	request = &GetProvisioningTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetProvisioningTaskStatus")

	return
}

func NewGetProvisioningTaskStatusResponse() (response *GetProvisioningTaskStatusResponse) {
	response = &GetProvisioningTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetProvisioningTaskStatus
// 查询用户同步异步任务的状态
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_USERPROVISIONINGTASKNOTFOUND = "ResourceNotFound.UserProvisioningTaskNotFound"
func (c *Client) GetProvisioningTaskStatus(request *GetProvisioningTaskStatusRequest) (response *GetProvisioningTaskStatusResponse, err error) {
	if request == nil {
		request = NewGetProvisioningTaskStatusRequest()
	}

	response = NewGetProvisioningTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetUserSyncProvisioningRequest() (request *GetUserSyncProvisioningRequest) {
	request = &GetUserSyncProvisioningRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "GetUserSyncProvisioning")

	return
}

func NewGetUserSyncProvisioningResponse() (response *GetUserSyncProvisioningResponse) {
	response = &GetUserSyncProvisioningResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// GetUserSyncProvisioning
// 查询CAM用户同步
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) GetUserSyncProvisioning(request *GetUserSyncProvisioningRequest) (response *GetUserSyncProvisioningResponse, err error) {
	if request == nil {
		request = NewGetUserSyncProvisioningRequest()
	}

	response = NewGetUserSyncProvisioningResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUserSyncProvisioningRequest() (request *UpdateUserSyncProvisioningRequest) {
	request = &UpdateUserSyncProvisioningRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "UpdateUserSyncProvisioning")

	return
}

func NewUpdateUserSyncProvisioningResponse() (response *UpdateUserSyncProvisioningResponse) {
	response = &UpdateUserSyncProvisioningResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// UpdateUserSyncProvisioning
// 创建子用户同步任务
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) UpdateUserSyncProvisioning(request *UpdateUserSyncProvisioningRequest) (response *UpdateUserSyncProvisioningResponse, err error) {
	if request == nil {
		request = NewUpdateUserSyncProvisioningRequest()
	}

	response = NewUpdateUserSyncProvisioningResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserSyncProvisioningRequest() (request *DeleteUserSyncProvisioningRequest) {
	request = &DeleteUserSyncProvisioningRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}

	request.Init().WithApiInfo("cic", APIVersion, "DeleteUserSyncProvisioning")

	return
}

func NewDeleteUserSyncProvisioningResponse() (response *DeleteUserSyncProvisioningResponse) {
	response = &DeleteUserSyncProvisioningResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return

}

// DeleteUserSyncProvisioning
// 删除子用户同步任务
//
// 可能返回的错误码:
//
//	FAILEDOPERATION_IDENTITYCENTERNOTOPEN = "FailedOperation.IdentityCenterNotOpen"
//	FAILEDOPERATION_USERPROVISIONINGFAILED = "FailedOperation.UserProvisioningFailed"
//	FAILEDOPERATION_ZONEIDNOTEXIST = "FailedOperation.ZoneIdNotExist"
//	RESOURCENOTFOUND_USERPROVISIONINGNOTFOUND = "ResourceNotFound.UserProvisioningNotFound"
func (c *Client) DeleteUserSyncProvisioning(request *DeleteUserSyncProvisioningRequest) (response *DeleteUserSyncProvisioningResponse, err error) {
	if request == nil {
		request = NewDeleteUserSyncProvisioningRequest()
	}

	response = NewDeleteUserSyncProvisioningResponse()
	err = c.Send(request, response)
	return
}

func NewListGroupsRequest() (request *ListGroupsRequest) {
	request = &ListGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cic", APIVersion, "ListGroups")
	return
}

func NewListGroupsResponse() (response *ListGroupsResponse) {
	response = &ListGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) ListGroups(request *ListGroupsRequest) (response *ListGroupsResponse, err error) {
	if request == nil {
		request = NewListGroupsRequest()
	}

	response = NewListGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewListRoleConfigurationsRequest() (request *ListRoleConfigurationsRequest) {
	request = &ListRoleConfigurationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cic", APIVersion, "ListRoleConfigurations")
	return
}

func NewListRoleConfigurationsResponse() (response *ListRoleConfigurationsResponse) {
	response = &ListRoleConfigurationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

func (c *Client) ListRoleConfigurations(request *ListRoleConfigurationsRequest) (response *ListRoleConfigurationsResponse, err error) {
	if request == nil {
		request = NewListRoleConfigurationsRequest()
	}

	response = NewListRoleConfigurationsResponse()
	err = c.Send(request, response)
	return
}

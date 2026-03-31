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

package v20191018

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-10-18"

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

func NewReplaySessionRequest() (request *ReplaySessionRequest) {
	request = &ReplaySessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ReplaySession")
	return
}

func NewReplaySessionResponse() (response *ReplaySessionResponse) {
	response = &ReplaySessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 会话回放
func (c *Client) ReplaySession(request *ReplaySessionRequest) (response *ReplaySessionResponse, err error) {
	if request == nil {
		request = NewReplaySessionRequest()
	}
	response = NewReplaySessionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessControlTemplateRequest() (request *CreateAccessControlTemplateRequest) {
	request = &CreateAccessControlTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAccessControlTemplate")
	return
}

func NewCreateAccessControlTemplateResponse() (response *CreateAccessControlTemplateResponse) {
	response = &CreateAccessControlTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建访问控制模板
func (c *Client) CreateAccessControlTemplate(request *CreateAccessControlTemplateRequest) (response *CreateAccessControlTemplateResponse, err error) {
	if request == nil {
		request = NewCreateAccessControlTemplateRequest()
	}
	response = NewCreateAccessControlTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
	request = &DescribeResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeResources")
	return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
	response = &DescribeResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户购买的堡垒机服务信息，包括资源ID、授权点数、VPC、过期时间等。
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeResourcesRequest()
	}
	response = NewDescribeResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessWhiteListRulesRequest() (request *DescribeAccessWhiteListRulesRequest) {
	request = &DescribeAccessWhiteListRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccessWhiteListRules")
	return
}

func NewDescribeAccessWhiteListRulesResponse() (response *DescribeAccessWhiteListRulesResponse) {
	response = &DescribeAccessWhiteListRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问白名单规则列表
func (c *Client) DescribeAccessWhiteListRules(request *DescribeAccessWhiteListRulesRequest) (response *DescribeAccessWhiteListRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAccessWhiteListRulesRequest()
	}
	response = NewDescribeAccessWhiteListRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDisableIntranetAccessRequest() (request *DisableIntranetAccessRequest) {
	request = &DisableIntranetAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DisableIntranetAccess")
	return
}

func NewDisableIntranetAccessResponse() (response *DisableIntranetAccessResponse) {
	response = &DisableIntranetAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭内网访问
func (c *Client) DisableIntranetAccess(request *DisableIntranetAccessRequest) (response *DisableIntranetAccessResponse, err error) {
	if request == nil {
		request = NewDisableIntranetAccessRequest()
	}
	response = NewDisableIntranetAccessResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExportUserTaskRequest() (request *DeleteExportUserTaskRequest) {
	request = &DeleteExportUserTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteExportUserTask")
	return
}

func NewDeleteExportUserTaskResponse() (response *DeleteExportUserTaskResponse) {
	response = &DeleteExportUserTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除导出用户任务
func (c *Client) DeleteExportUserTask(request *DeleteExportUserTaskRequest) (response *DeleteExportUserTaskResponse, err error) {
	if request == nil {
		request = NewDeleteExportUserTaskRequest()
	}
	response = NewDeleteExportUserTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMFAPreCheckRequest() (request *DescribeMFAPreCheckRequest) {
	request = &DescribeMFAPreCheckRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeMFAPreCheck")
	return
}

func NewDescribeMFAPreCheckResponse() (response *DescribeMFAPreCheckResponse) {
	response = &DescribeMFAPreCheckResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询多因子认证预校验
func (c *Client) DescribeMFAPreCheck(request *DescribeMFAPreCheckRequest) (response *DescribeMFAPreCheckResponse, err error) {
	if request == nil {
		request = NewDescribeMFAPreCheckRequest()
	}
	response = NewDescribeMFAPreCheckResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePushAccountTaskDetailRequest() (request *DescribePushAccountTaskDetailRequest) {
	request = &DescribePushAccountTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribePushAccountTaskDetail")
	return
}

func NewDescribePushAccountTaskDetailResponse() (response *DescribePushAccountTaskDetailResponse) {
	response = &DescribePushAccountTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账号推送任务详情
func (c *Client) DescribePushAccountTaskDetail(request *DescribePushAccountTaskDetailRequest) (response *DescribePushAccountTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribePushAccountTaskDetailRequest()
	}
	response = NewDescribePushAccountTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewKillSessionRequest() (request *KillSessionRequest) {
	request = &KillSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "KillSession")
	return
}

func NewKillSessionResponse() (response *KillSessionResponse) {
	response = &KillSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 会话切断
func (c *Client) KillSession(request *KillSessionRequest) (response *KillSessionResponse, err error) {
	if request == nil {
		request = NewKillSessionRequest()
	}
	response = NewKillSessionResponse()
	err = c.Send(request, response)
	return
}

func NewSearchPushAccountTaskInfoRequest() (request *SearchPushAccountTaskInfoRequest) {
	request = &SearchPushAccountTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchPushAccountTaskInfo")
	return
}

func NewSearchPushAccountTaskInfoResponse() (response *SearchPushAccountTaskInfoResponse) {
	response = &SearchPushAccountTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索账号推送任务详情
func (c *Client) SearchPushAccountTaskInfo(request *SearchPushAccountTaskInfoRequest) (response *SearchPushAccountTaskInfoResponse, err error) {
	if request == nil {
		request = NewSearchPushAccountTaskInfoRequest()
	}
	response = NewSearchPushAccountTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCdcSettingRequest() (request *DescribeCdcSettingRequest) {
	request = &DescribeCdcSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeCdcSetting")
	return
}

func NewDescribeCdcSettingResponse() (response *DescribeCdcSettingResponse) {
	response = &DescribeCdcSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于描述用户的CDC环境参数设置信息，比如用户能否在CDC环境下开通堡垒机资源、用户的CDC集群ID等。
func (c *Client) DescribeCdcSetting(request *DescribeCdcSettingRequest) (response *DescribeCdcSettingResponse, err error) {
	if request == nil {
		request = NewDescribeCdcSettingRequest()
	}
	response = NewDescribeCdcSettingResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppAssetRequest() (request *ModifyAppAssetRequest) {
	request = &ModifyAppAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAppAsset")
	return
}

func NewModifyAppAssetResponse() (response *ModifyAppAssetResponse) {
	response = &ModifyAppAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改应用资产
func (c *Client) ModifyAppAsset(request *ModifyAppAssetRequest) (response *ModifyAppAssetResponse, err error) {
	if request == nil {
		request = NewModifyAppAssetRequest()
	}
	response = NewModifyAppAssetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePushAccountTaskRequest() (request *DescribePushAccountTaskRequest) {
	request = &DescribePushAccountTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribePushAccountTask")
	return
}

func NewDescribePushAccountTaskResponse() (response *DescribePushAccountTaskResponse) {
	response = &DescribePushAccountTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账号推送任务
func (c *Client) DescribePushAccountTask(request *DescribePushAccountTaskRequest) (response *DescribePushAccountTaskResponse, err error) {
	if request == nil {
		request = NewDescribePushAccountTaskRequest()
	}
	response = NewDescribePushAccountTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSearchCommandBySidRequest() (request *SearchCommandBySidRequest) {
	request = &SearchCommandBySidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchCommandBySid")
	return
}

func NewSearchCommandBySidResponse() (response *SearchCommandBySidResponse) {
	response = &SearchCommandBySidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据会话Id搜索Command
func (c *Client) SearchCommandBySid(request *SearchCommandBySidRequest) (response *SearchCommandBySidResponse, err error) {
	if request == nil {
		request = NewSearchCommandBySidRequest()
	}
	response = NewSearchCommandBySidResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmSettingRequest() (request *ModifyAlarmSettingRequest) {
	request = &ModifyAlarmSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAlarmSetting")
	return
}

func NewModifyAlarmSettingResponse() (response *ModifyAlarmSettingResponse) {
	response = &ModifyAlarmSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警设置
func (c *Client) ModifyAlarmSetting(request *ModifyAlarmSettingRequest) (response *ModifyAlarmSettingResponse, err error) {
	if request == nil {
		request = NewModifyAlarmSettingRequest()
	}
	response = NewModifyAlarmSettingResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessControlRuleRequest() (request *CreateAccessControlRuleRequest) {
	request = &CreateAccessControlRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAccessControlRule")
	return
}

func NewCreateAccessControlRuleResponse() (response *CreateAccessControlRuleResponse) {
	response = &CreateAccessControlRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建访问控制策略
func (c *Client) CreateAccessControlRule(request *CreateAccessControlRuleRequest) (response *CreateAccessControlRuleResponse, err error) {
	if request == nil {
		request = NewCreateAccessControlRuleRequest()
	}
	response = NewCreateAccessControlRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppAssetsRequest() (request *DescribeAppAssetsRequest) {
	request = &DescribeAppAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAppAssets")
	return
}

func NewDescribeAppAssetsResponse() (response *DescribeAppAssetsResponse) {
	response = &DescribeAppAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用资产
func (c *Client) DescribeAppAssets(request *DescribeAppAssetsRequest) (response *DescribeAppAssetsResponse, err error) {
	if request == nil {
		request = NewDescribeAppAssetsRequest()
	}
	response = NewDescribeAppAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessControlRulesRequest() (request *DeleteAccessControlRulesRequest) {
	request = &DeleteAccessControlRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAccessControlRules")
	return
}

func NewDeleteAccessControlRulesResponse() (response *DeleteAccessControlRulesResponse) {
	response = &DeleteAccessControlRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问控制策略
func (c *Client) DeleteAccessControlRules(request *DeleteAccessControlRulesRequest) (response *DeleteAccessControlRulesResponse, err error) {
	if request == nil {
		request = NewDeleteAccessControlRulesRequest()
	}
	response = NewDeleteAccessControlRulesResponse()
	err = c.Send(request, response)
	return
}

func NewImportDevicesRequest() (request *ImportDevicesRequest) {
	request = &ImportDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ImportDevices")
	return
}

func NewImportDevicesResponse() (response *ImportDevicesResponse) {
	response = &ImportDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入主机信息
func (c *Client) ImportDevices(request *ImportDevicesRequest) (response *ImportDevicesResponse, err error) {
	if request == nil {
		request = NewImportDevicesRequest()
	}
	response = NewImportDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewAddUserGroupMembersRequest() (request *AddUserGroupMembersRequest) {
	request = &AddUserGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "AddUserGroupMembers")
	return
}

func NewAddUserGroupMembersResponse() (response *AddUserGroupMembersResponse) {
	response = &AddUserGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加用户组成员
func (c *Client) AddUserGroupMembers(request *AddUserGroupMembersRequest) (response *AddUserGroupMembersResponse, err error) {
	if request == nil {
		request = NewAddUserGroupMembersRequest()
	}
	response = NewAddUserGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountsWithDeviceCountRequest() (request *DescribeAccountsWithDeviceCountRequest) {
	request = &DescribeAccountsWithDeviceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccountsWithDeviceCount")
	return
}

func NewDescribeAccountsWithDeviceCountResponse() (response *DescribeAccountsWithDeviceCountResponse) {
	response = &DescribeAccountsWithDeviceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询账号,带拥有相同账号的主机数，可以按照主机、主机组查询，不分页
func (c *Client) DescribeAccountsWithDeviceCount(request *DescribeAccountsWithDeviceCountRequest) (response *DescribeAccountsWithDeviceCountResponse, err error) {
	if request == nil {
		request = NewDescribeAccountsWithDeviceCountRequest()
	}
	response = NewDescribeAccountsWithDeviceCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceAccountsRequest() (request *DescribeDeviceAccountsRequest) {
	request = &DescribeDeviceAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceAccounts")
	return
}

func NewDescribeDeviceAccountsResponse() (response *DescribeDeviceAccountsResponse) {
	response = &DescribeDeviceAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询主机账号列表
func (c *Client) DescribeDeviceAccounts(request *DescribeDeviceAccountsRequest) (response *DescribeDeviceAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceAccountsRequest()
	}
	response = NewDescribeDeviceAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceRequest() (request *ModifyResourceRequest) {
	request = &ModifyResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyResource")
	return
}

func NewModifyResourceResponse() (response *ModifyResourceResponse) {
	response = &ModifyResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资源变配
func (c *Client) ModifyResource(request *ModifyResourceRequest) (response *ModifyResourceResponse, err error) {
	if request == nil {
		request = NewModifyResourceRequest()
	}
	response = NewModifyResourceResponse()
	err = c.Send(request, response)
	return
}

func NewRunPushAccountTaskRequest() (request *RunPushAccountTaskRequest) {
	request = &RunPushAccountTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "RunPushAccountTask")
	return
}

func NewRunPushAccountTaskResponse() (response *RunPushAccountTaskResponse) {
	response = &RunPushAccountTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行账号推送任务
func (c *Client) RunPushAccountTask(request *RunPushAccountTaskRequest) (response *RunPushAccountTaskResponse, err error) {
	if request == nil {
		request = NewRunPushAccountTaskRequest()
	}
	response = NewRunPushAccountTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSearchAuditLogRequest() (request *SearchAuditLogRequest) {
	request = &SearchAuditLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchAuditLog")
	return
}

func NewSearchAuditLogResponse() (response *SearchAuditLogResponse) {
	response = &SearchAuditLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索审计日志
func (c *Client) SearchAuditLog(request *SearchAuditLogRequest) (response *SearchAuditLogResponse, err error) {
	if request == nil {
		request = NewSearchAuditLogRequest()
	}
	response = NewSearchAuditLogResponse()
	err = c.Send(request, response)
	return
}

func NewCanCreateTrialResourceRequest() (request *CanCreateTrialResourceRequest) {
	request = &CanCreateTrialResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CanCreateTrialResource")
	return
}

func NewCanCreateTrialResourceResponse() (response *CanCreateTrialResourceResponse) {
	response = &CanCreateTrialResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 判断是否可以创建试用版堡垒机
func (c *Client) CanCreateTrialResource(request *CanCreateTrialResourceRequest) (response *CanCreateTrialResourceResponse, err error) {
	if request == nil {
		request = NewCanCreateTrialResourceRequest()
	}
	response = NewCanCreateTrialResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTasksRequest() (request *DescribeOperationTasksRequest) {
	request = &DescribeOperationTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationTasks")
	return
}

func NewDescribeOperationTasksResponse() (response *DescribeOperationTasksResponse) {
	response = &DescribeOperationTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维任务列表
func (c *Client) DescribeOperationTasks(request *DescribeOperationTasksRequest) (response *DescribeOperationTasksResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTasksRequest()
	}
	response = NewDescribeOperationTasksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDasbResourceRequest() (request *ModifyDasbResourceRequest) {
	request = &ModifyDasbResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDasbResource")
	return
}

func NewModifyDasbResourceResponse() (response *ModifyDasbResourceResponse) {
	response = &ModifyDasbResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 堡垒机变配接口
func (c *Client) ModifyDasbResource(request *ModifyDasbResourceRequest) (response *ModifyDasbResourceResponse, err error) {
	if request == nil {
		request = NewModifyDasbResourceRequest()
	}
	response = NewModifyDasbResourceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExportDeviceTaskRequest() (request *CreateExportDeviceTaskRequest) {
	request = &CreateExportDeviceTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateExportDeviceTask")
	return
}

func NewCreateExportDeviceTaskResponse() (response *CreateExportDeviceTaskResponse) {
	response = &CreateExportDeviceTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建资产导入任务
func (c *Client) CreateExportDeviceTask(request *CreateExportDeviceTaskRequest) (response *CreateExportDeviceTaskResponse, err error) {
	if request == nil {
		request = NewCreateExportDeviceTaskRequest()
	}
	response = NewCreateExportDeviceTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessWhiteListStatusRequest() (request *ModifyAccessWhiteListStatusRequest) {
	request = &ModifyAccessWhiteListStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessWhiteListStatus")
	return
}

func NewModifyAccessWhiteListStatusResponse() (response *ModifyAccessWhiteListStatusResponse) {
	response = &ModifyAccessWhiteListStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问白名单状态：开启或关闭放开全部来源IP。
func (c *Client) ModifyAccessWhiteListStatus(request *ModifyAccessWhiteListStatusRequest) (response *ModifyAccessWhiteListStatusResponse, err error) {
	if request == nil {
		request = NewModifyAccessWhiteListStatusRequest()
	}
	response = NewModifyAccessWhiteListStatusResponse()
	err = c.Send(request, response)
	return
}

func NewApproveTicketRequest() (request *ApproveTicketRequest) {
	request = &ApproveTicketRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ApproveTicket")
	return
}

func NewApproveTicketResponse() (response *ApproveTicketResponse) {
	response = &ApproveTicketResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审批工单
func (c *Client) ApproveTicket(request *ApproveTicketRequest) (response *ApproveTicketResponse, err error) {
	if request == nil {
		request = NewApproveTicketRequest()
	}
	response = NewApproveTicketResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDepartmentRequest() (request *DeleteDepartmentRequest) {
	request = &DeleteDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDepartment")
	return
}

func NewDeleteDepartmentResponse() (response *DeleteDepartmentResponse) {
	response = &DeleteDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除部门
func (c *Client) DeleteDepartment(request *DeleteDepartmentRequest) (response *DeleteDepartmentResponse, err error) {
	if request == nil {
		request = NewDeleteDepartmentRequest()
	}
	response = NewDeleteDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessControlRuleRequest() (request *ModifyAccessControlRuleRequest) {
	request = &ModifyAccessControlRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessControlRule")
	return
}

func NewModifyAccessControlRuleResponse() (response *ModifyAccessControlRuleResponse) {
	response = &ModifyAccessControlRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问控制策略
func (c *Client) ModifyAccessControlRule(request *ModifyAccessControlRuleRequest) (response *ModifyAccessControlRuleResponse, err error) {
	if request == nil {
		request = NewModifyAccessControlRuleRequest()
	}
	response = NewModifyAccessControlRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUnlockUserRequest() (request *UnlockUserRequest) {
	request = &UnlockUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "UnlockUser")
	return
}

func NewUnlockUserResponse() (response *UnlockUserResponse) {
	response = &UnlockUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解锁用户
func (c *Client) UnlockUser(request *UnlockUserRequest) (response *UnlockUserResponse, err error) {
	if request == nil {
		request = NewUnlockUserRequest()
	}
	response = NewUnlockUserResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExportUserTaskRequest() (request *CreateExportUserTaskRequest) {
	request = &CreateExportUserTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateExportUserTask")
	return
}

func NewCreateExportUserTaskResponse() (response *CreateExportUserTaskResponse) {
	response = &CreateExportUserTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用户导出任务
func (c *Client) CreateExportUserTask(request *CreateExportUserTaskRequest) (response *CreateExportUserTaskResponse, err error) {
	if request == nil {
		request = NewCreateExportUserTaskRequest()
	}
	response = NewCreateExportUserTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSearchFileRequest() (request *SearchFileRequest) {
	request = &SearchFileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchFile")
	return
}

func NewSearchFileResponse() (response *SearchFileResponse) {
	response = &SearchFileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 文件传输检索
func (c *Client) SearchFile(request *SearchFileRequest) (response *SearchFileResponse, err error) {
	if request == nil {
		request = NewSearchFileRequest()
	}
	response = NewSearchFileResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDepartmentsRequest() (request *DescribeDepartmentsRequest) {
	request = &DescribeDepartmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDepartments")
	return
}

func NewDescribeDepartmentsResponse() (response *DescribeDepartmentsResponse) {
	response = &DescribeDepartmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部门信息
func (c *Client) DescribeDepartments(request *DescribeDepartmentsRequest) (response *DescribeDepartmentsResponse, err error) {
	if request == nil {
		request = NewDescribeDepartmentsRequest()
	}
	response = NewDescribeDepartmentsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchEventRequest() (request *SearchEventRequest) {
	request = &SearchEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchEvent")
	return
}

func NewSearchEventResponse() (response *SearchEventResponse) {
	response = &SearchEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索高危事件
func (c *Client) SearchEvent(request *SearchEventRequest) (response *SearchEventResponse, err error) {
	if request == nil {
		request = NewSearchEventRequest()
	}
	response = NewSearchEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableInstanceTypesRequest() (request *DescribeAvailableInstanceTypesRequest) {
	request = &DescribeAvailableInstanceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAvailableInstanceTypes")
	return
}

func NewDescribeAvailableInstanceTypesResponse() (response *DescribeAvailableInstanceTypesResponse) {
	response = &DescribeAvailableInstanceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询有效的实例类型列表
func (c *Client) DescribeAvailableInstanceTypes(request *DescribeAvailableInstanceTypesRequest) (response *DescribeAvailableInstanceTypesResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableInstanceTypesRequest()
	}
	response = NewDescribeAvailableInstanceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePushAccountTaskRequest() (request *CreatePushAccountTaskRequest) {
	request = &CreatePushAccountTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreatePushAccountTask")
	return
}

func NewCreatePushAccountTaskResponse() (response *CreatePushAccountTaskResponse) {
	response = &CreatePushAccountTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建账号推送任务
func (c *Client) CreatePushAccountTask(request *CreatePushAccountTaskRequest) (response *CreatePushAccountTaskResponse, err error) {
	if request == nil {
		request = NewCreatePushAccountTaskRequest()
	}
	response = NewCreatePushAccountTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCmdTemplatesRequest() (request *DeleteCmdTemplatesRequest) {
	request = &DeleteCmdTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteCmdTemplates")
	return
}

func NewDeleteCmdTemplatesResponse() (response *DeleteCmdTemplatesResponse) {
	response = &DeleteCmdTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除高危命令模板
func (c *Client) DeleteCmdTemplates(request *DeleteCmdTemplatesRequest) (response *DeleteCmdTemplatesResponse, err error) {
	if request == nil {
		request = NewDeleteCmdTemplatesRequest()
	}
	response = NewDeleteCmdTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportAuditLogTaskRequest() (request *DescribeExportAuditLogTaskRequest) {
	request = &DescribeExportAuditLogTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeExportAuditLogTask")
	return
}

func NewDescribeExportAuditLogTaskResponse() (response *DescribeExportAuditLogTaskResponse) {
	response = &DescribeExportAuditLogTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取审计日志导出任务列表
func (c *Client) DescribeExportAuditLogTask(request *DescribeExportAuditLogTaskRequest) (response *DescribeExportAuditLogTaskResponse, err error) {
	if request == nil {
		request = NewDescribeExportAuditLogTaskRequest()
	}
	response = NewDescribeExportAuditLogTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppAssetsRequest() (request *DeleteAppAssetsRequest) {
	request = &DeleteAppAssetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAppAssets")
	return
}

func NewDeleteAppAssetsResponse() (response *DeleteAppAssetsResponse) {
	response = &DeleteAppAssetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用资产
func (c *Client) DeleteAppAssets(request *DeleteAppAssetsRequest) (response *DeleteAppAssetsResponse, err error) {
	if request == nil {
		request = NewDeleteAppAssetsRequest()
	}
	response = NewDeleteAppAssetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLogDeliveryRequest() (request *CreateLogDeliveryRequest) {
	request = &CreateLogDeliveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateLogDelivery")
	return
}

func NewCreateLogDeliveryResponse() (response *CreateLogDeliveryResponse) {
	response = &CreateLogDeliveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建日志投递服务
func (c *Client) CreateLogDelivery(request *CreateLogDeliveryRequest) (response *CreateLogDeliveryResponse, err error) {
	if request == nil {
		request = NewCreateLogDeliveryRequest()
	}
	response = NewCreateLogDeliveryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogDeliveryRequest() (request *ModifyLogDeliveryRequest) {
	request = &ModifyLogDeliveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyLogDelivery")
	return
}

func NewModifyLogDeliveryResponse() (response *ModifyLogDeliveryResponse) {
	response = &ModifyLogDeliveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志投递配置
func (c *Client) ModifyLogDelivery(request *ModifyLogDeliveryRequest) (response *ModifyLogDeliveryResponse, err error) {
	if request == nil {
		request = NewModifyLogDeliveryRequest()
	}
	response = NewModifyLogDeliveryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUsersDepartmentRequest() (request *ModifyUsersDepartmentRequest) {
	request = &ModifyUsersDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyUsersDepartment")
	return
}

func NewModifyUsersDepartmentResponse() (response *ModifyUsersDepartmentResponse) {
	response = &ModifyUsersDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改用户部门
func (c *Client) ModifyUsersDepartment(request *ModifyUsersDepartmentRequest) (response *ModifyUsersDepartmentResponse, err error) {
	if request == nil {
		request = NewModifyUsersDepartmentRequest()
	}
	response = NewModifyUsersDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewCreateReportTaskRequest() (request *CreateReportTaskRequest) {
	request = &CreateReportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateReportTask")
	return
}

func NewCreateReportTaskResponse() (response *CreateReportTaskResponse) {
	response = &CreateReportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建运维安全报表任务
func (c *Client) CreateReportTask(request *CreateReportTaskRequest) (response *CreateReportTaskResponse, err error) {
	if request == nil {
		request = NewCreateReportTaskRequest()
	}
	response = NewCreateReportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserGroupRequest() (request *CreateUserGroupRequest) {
	request = &CreateUserGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateUserGroup")
	return
}

func NewCreateUserGroupResponse() (response *CreateUserGroupResponse) {
	response = &CreateUserGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建用户组
func (c *Client) CreateUserGroup(request *CreateUserGroupRequest) (response *CreateUserGroupResponse, err error) {
	if request == nil {
		request = NewCreateUserGroupRequest()
	}
	response = NewCreateUserGroupResponse()
	err = c.Send(request, response)
	return
}

func NewLockUserRequest() (request *LockUserRequest) {
	request = &LockUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "LockUser")
	return
}

func NewLockUserResponse() (response *LockUserResponse) {
	response = &LockUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 锁定用户
func (c *Client) LockUser(request *LockUserRequest) (response *LockUserResponse, err error) {
	if request == nil {
		request = NewLockUserRequest()
	}
	response = NewLockUserResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDevicesSSLRequest() (request *ModifyDevicesSSLRequest) {
	request = &ModifyDevicesSSLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDevicesSSL")
	return
}

func NewModifyDevicesSSLResponse() (response *ModifyDevicesSSLResponse) {
	response = &ModifyDevicesSSLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产SSL信息
func (c *Client) ModifyDevicesSSL(request *ModifyDevicesSSLRequest) (response *ModifyDevicesSSLResponse, err error) {
	if request == nil {
		request = NewModifyDevicesSSLRequest()
	}
	response = NewModifyDevicesSSLResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUKeysRequest() (request *DeleteUKeysRequest) {
	request = &DeleteUKeysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteUKeys")
	return
}

func NewDeleteUKeysResponse() (response *DeleteUKeysResponse) {
	response = &DeleteUKeysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Ukey
func (c *Client) DeleteUKeys(request *DeleteUKeysRequest) (response *DeleteUKeysResponse, err error) {
	if request == nil {
		request = NewDeleteUKeysRequest()
	}
	response = NewDeleteUKeysResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceIdsRequest() (request *DescribeInstanceIdsRequest) {
	request = &DescribeInstanceIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeInstanceIds")
	return
}

func NewDescribeInstanceIdsResponse() (response *DescribeInstanceIdsResponse) {
	response = &DescribeInstanceIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定区域所有主机的InstanceId，不分页
func (c *Client) DescribeInstanceIds(request *DescribeInstanceIdsRequest) (response *DescribeInstanceIdsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceIdsRequest()
	}
	response = NewDescribeInstanceIdsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoginEventRequest() (request *DescribeLoginEventRequest) {
	request = &DescribeLoginEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLoginEvent")
	return
}

func NewDescribeLoginEventResponse() (response *DescribeLoginEventResponse) {
	response = &DescribeLoginEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询登录日志
func (c *Client) DescribeLoginEvent(request *DescribeLoginEventRequest) (response *DescribeLoginEventResponse, err error) {
	if request == nil {
		request = NewDescribeLoginEventRequest()
	}
	response = NewDescribeLoginEventResponse()
	err = c.Send(request, response)
	return
}

func NewResetDeviceAccountPrivateKeyRequest() (request *ResetDeviceAccountPrivateKeyRequest) {
	request = &ResetDeviceAccountPrivateKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetDeviceAccountPrivateKey")
	return
}

func NewResetDeviceAccountPrivateKeyResponse() (response *ResetDeviceAccountPrivateKeyResponse) {
	response = &ResetDeviceAccountPrivateKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清除设备账号绑定的密钥
func (c *Client) ResetDeviceAccountPrivateKey(request *ResetDeviceAccountPrivateKeyRequest) (response *ResetDeviceAccountPrivateKeyResponse, err error) {
	if request == nil {
		request = NewResetDeviceAccountPrivateKeyRequest()
	}
	response = NewResetDeviceAccountPrivateKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppAssetGroupMembersRequest() (request *DescribeAppAssetGroupMembersRequest) {
	request = &DescribeAppAssetGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAppAssetGroupMembers")
	return
}

func NewDescribeAppAssetGroupMembersResponse() (response *DescribeAppAssetGroupMembersResponse) {
	response = &DescribeAppAssetGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用资产组成员
func (c *Client) DescribeAppAssetGroupMembers(request *DescribeAppAssetGroupMembersRequest) (response *DescribeAppAssetGroupMembersResponse, err error) {
	if request == nil {
		request = NewDescribeAppAssetGroupMembersRequest()
	}
	response = NewDescribeAppAssetGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewViewReportRequest() (request *ViewReportRequest) {
	request = &ViewReportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ViewReport")
	return
}

func NewViewReportResponse() (response *ViewReportResponse) {
	response = &ViewReportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 审计报表预览
func (c *Client) ViewReport(request *ViewReportRequest) (response *ViewReportResponse, err error) {
	if request == nil {
		request = NewViewReportRequest()
	}
	response = NewViewReportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlRuleRequest() (request *DescribeAccessControlRuleRequest) {
	request = &DescribeAccessControlRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccessControlRule")
	return
}

func NewDescribeAccessControlRuleResponse() (response *DescribeAccessControlRuleResponse) {
	response = &DescribeAccessControlRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取访问控制策略详情
func (c *Client) DescribeAccessControlRule(request *DescribeAccessControlRuleRequest) (response *DescribeAccessControlRuleResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlRuleRequest()
	}
	response = NewDescribeAccessControlRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsersRequest() (request *DescribeUsersRequest) {
	request = &DescribeUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUsers")
	return
}

func NewDescribeUsersResponse() (response *DescribeUsersResponse) {
	response = &DescribeUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户列表
func (c *Client) DescribeUsers(request *DescribeUsersRequest) (response *DescribeUsersResponse, err error) {
	if request == nil {
		request = NewDescribeUsersRequest()
	}
	response = NewDescribeUsersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAclsRequest() (request *DescribeAclsRequest) {
	request = &DescribeAclsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAcls")
	return
}

func NewDescribeAclsResponse() (response *DescribeAclsResponse) {
	response = &DescribeAclsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问权限列表
func (c *Client) DescribeAcls(request *DescribeAclsRequest) (response *DescribeAclsResponse, err error) {
	if request == nil {
		request = NewDescribeAclsRequest()
	}
	response = NewDescribeAclsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessWhiteListRulesRequest() (request *DeleteAccessWhiteListRulesRequest) {
	request = &DeleteAccessWhiteListRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAccessWhiteListRules")
	return
}

func NewDeleteAccessWhiteListRulesResponse() (response *DeleteAccessWhiteListRulesResponse) {
	response = &DeleteAccessWhiteListRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问白名单规则
func (c *Client) DeleteAccessWhiteListRules(request *DeleteAccessWhiteListRulesRequest) (response *DeleteAccessWhiteListRulesResponse, err error) {
	if request == nil {
		request = NewDeleteAccessWhiteListRulesRequest()
	}
	response = NewDeleteAccessWhiteListRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTicketsRequest() (request *DescribeTicketsRequest) {
	request = &DescribeTicketsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeTickets")
	return
}

func NewDescribeTicketsResponse() (response *DescribeTicketsResponse) {
	response = &DescribeTicketsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维工单列表
func (c *Client) DescribeTickets(request *DescribeTicketsRequest) (response *DescribeTicketsResponse, err error) {
	if request == nil {
		request = NewDescribeTicketsRequest()
	}
	response = NewDescribeTicketsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateChangePwdTaskRequest() (request *CreateChangePwdTaskRequest) {
	request = &CreateChangePwdTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateChangePwdTask")
	return
}

func NewCreateChangePwdTaskResponse() (response *CreateChangePwdTaskResponse) {
	response = &CreateChangePwdTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建修改密码任务
func (c *Client) CreateChangePwdTask(request *CreateChangePwdTaskRequest) (response *CreateChangePwdTaskResponse, err error) {
	if request == nil {
		request = NewCreateChangePwdTaskRequest()
	}
	response = NewCreateChangePwdTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserGroupMembersRequest() (request *DeleteUserGroupMembersRequest) {
	request = &DeleteUserGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteUserGroupMembers")
	return
}

func NewDeleteUserGroupMembersResponse() (response *DeleteUserGroupMembersResponse) {
	response = &DeleteUserGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户组成员
func (c *Client) DeleteUserGroupMembers(request *DeleteUserGroupMembersRequest) (response *DeleteUserGroupMembersResponse, err error) {
	if request == nil {
		request = NewDeleteUserGroupMembersRequest()
	}
	response = NewDeleteUserGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExportDeviceTaskRequest() (request *DeleteExportDeviceTaskRequest) {
	request = &DeleteExportDeviceTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteExportDeviceTask")
	return
}

func NewDeleteExportDeviceTaskResponse() (response *DeleteExportDeviceTaskResponse) {
	response = &DeleteExportDeviceTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资产导出任务
func (c *Client) DeleteExportDeviceTask(request *DeleteExportDeviceTaskRequest) (response *DeleteExportDeviceTaskResponse, err error) {
	if request == nil {
		request = NewDeleteExportDeviceTaskRequest()
	}
	response = NewDeleteExportDeviceTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTrialGuideRequest() (request *DescribeTrialGuideRequest) {
	request = &DescribeTrialGuideRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeTrialGuide")
	return
}

func NewDescribeTrialGuideResponse() (response *DescribeTrialGuideResponse) {
	response = &DescribeTrialGuideResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 堡垒机试用引导
func (c *Client) DescribeTrialGuide(request *DescribeTrialGuideRequest) (response *DescribeTrialGuideResponse, err error) {
	if request == nil {
		request = NewDescribeTrialGuideRequest()
	}
	response = NewDescribeTrialGuideResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserRequest() (request *ModifyUserRequest) {
	request = &ModifyUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyUser")
	return
}

func NewModifyUserResponse() (response *ModifyUserResponse) {
	response = &ModifyUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户信息
func (c *Client) ModifyUser(request *ModifyUserRequest) (response *ModifyUserResponse, err error) {
	if request == nil {
		request = NewModifyUserRequest()
	}
	response = NewModifyUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReportTaskHistoryRequest() (request *DescribeReportTaskHistoryRequest) {
	request = &DescribeReportTaskHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeReportTaskHistory")
	return
}

func NewDescribeReportTaskHistoryResponse() (response *DescribeReportTaskHistoryResponse) {
	response = &DescribeReportTaskHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询审计报表任务记录
func (c *Client) DescribeReportTaskHistory(request *DescribeReportTaskHistoryRequest) (response *DescribeReportTaskHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeReportTaskHistoryRequest()
	}
	response = NewDescribeReportTaskHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDevicesRequest() (request *DeleteDevicesRequest) {
	request = &DeleteDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDevices")
	return
}

func NewDeleteDevicesResponse() (response *DeleteDevicesResponse) {
	response = &DeleteDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除主机
func (c *Client) DeleteDevices(request *DeleteDevicesRequest) (response *DeleteDevicesResponse, err error) {
	if request == nil {
		request = NewDeleteDevicesRequest()
	}
	response = NewDeleteDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDomainsRequest() (request *DeleteDomainsRequest) {
	request = &DeleteDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDomains")
	return
}

func NewDeleteDomainsResponse() (response *DeleteDomainsResponse) {
	response = &DeleteDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除网络域
func (c *Client) DeleteDomains(request *DeleteDomainsRequest) (response *DeleteDomainsResponse, err error) {
	if request == nil {
		request = NewDeleteDomainsRequest()
	}
	response = NewDeleteDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewAddDeviceGroupMembersRequest() (request *AddDeviceGroupMembersRequest) {
	request = &AddDeviceGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "AddDeviceGroupMembers")
	return
}

func NewAddDeviceGroupMembersResponse() (response *AddDeviceGroupMembersResponse) {
	response = &AddDeviceGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加资产组成员
func (c *Client) AddDeviceGroupMembers(request *AddDeviceGroupMembersRequest) (response *AddDeviceGroupMembersResponse, err error) {
	if request == nil {
		request = NewAddDeviceGroupMembersRequest()
	}
	response = NewAddDeviceGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAssetSyncJobRequest() (request *CreateAssetSyncJobRequest) {
	request = &CreateAssetSyncJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAssetSyncJob")
	return
}

func NewCreateAssetSyncJobResponse() (response *CreateAssetSyncJobResponse) {
	response = &CreateAssetSyncJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建手工资产同步任务
func (c *Client) CreateAssetSyncJob(request *CreateAssetSyncJobRequest) (response *CreateAssetSyncJobResponse, err error) {
	if request == nil {
		request = NewCreateAssetSyncJobRequest()
	}
	response = NewCreateAssetSyncJobResponse()
	err = c.Send(request, response)
	return
}

func NewLeaveTrackPageRequest() (request *LeaveTrackPageRequest) {
	request = &LeaveTrackPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "LeaveTrackPage")
	return
}

func NewLeaveTrackPageResponse() (response *LeaveTrackPageResponse) {
	response = &LeaveTrackPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 离开埋点采集页面
func (c *Client) LeaveTrackPage(request *LeaveTrackPageRequest) (response *LeaveTrackPageResponse, err error) {
	if request == nil {
		request = NewLeaveTrackPageRequest()
	}
	response = NewLeaveTrackPageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPasswordSettingRequest() (request *ModifyPasswordSettingRequest) {
	request = &ModifyPasswordSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyPasswordSetting")
	return
}

func NewModifyPasswordSettingResponse() (response *ModifyPasswordSettingResponse) {
	response = &ModifyPasswordSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改密码配置信息
func (c *Client) ModifyPasswordSetting(request *ModifyPasswordSettingRequest) (response *ModifyPasswordSettingResponse, err error) {
	if request == nil {
		request = NewModifyPasswordSettingRequest()
	}
	response = NewModifyPasswordSettingResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserBatchRequest() (request *ModifyUserBatchRequest) {
	request = &ModifyUserBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyUserBatch")
	return
}

func NewModifyUserBatchResponse() (response *ModifyUserBatchResponse) {
	response = &ModifyUserBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改用户
func (c *Client) ModifyUserBatch(request *ModifyUserBatchRequest) (response *ModifyUserBatchResponse, err error) {
	if request == nil {
		request = NewModifyUserBatchRequest()
	}
	response = NewModifyUserBatchResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSubtaskResultByIdRequest() (request *SearchSubtaskResultByIdRequest) {
	request = &SearchSubtaskResultByIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchSubtaskResultById")
	return
}

func NewSearchSubtaskResultByIdResponse() (response *SearchSubtaskResultByIdResponse) {
	response = &SearchSubtaskResultByIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维子任务执行结果
func (c *Client) SearchSubtaskResultById(request *SearchSubtaskResultByIdRequest) (response *SearchSubtaskResultByIdResponse, err error) {
	if request == nil {
		request = NewSearchSubtaskResultByIdRequest()
	}
	response = NewSearchSubtaskResultByIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEnvSettingRequest() (request *DescribeEnvSettingRequest) {
	request = &DescribeEnvSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeEnvSetting")
	return
}

func NewDescribeEnvSettingResponse() (response *DescribeEnvSettingResponse) {
	response = &DescribeEnvSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运行环境设置
func (c *Client) DescribeEnvSetting(request *DescribeEnvSettingRequest) (response *DescribeEnvSettingResponse, err error) {
	if request == nil {
		request = NewDescribeEnvSettingRequest()
	}
	response = NewDescribeEnvSettingResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoginSettingRequest() (request *ModifyLoginSettingRequest) {
	request = &ModifyLoginSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyLoginSetting")
	return
}

func NewModifyLoginSettingResponse() (response *ModifyLoginSettingResponse) {
	response = &ModifyLoginSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改登录相关配置信息
func (c *Client) ModifyLoginSetting(request *ModifyLoginSettingRequest) (response *ModifyLoginSettingResponse, err error) {
	if request == nil {
		request = NewModifyLoginSettingRequest()
	}
	response = NewModifyLoginSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeviceAccountsRequest() (request *DeleteDeviceAccountsRequest) {
	request = &DeleteDeviceAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDeviceAccounts")
	return
}

func NewDeleteDeviceAccountsResponse() (response *DeleteDeviceAccountsResponse) {
	response = &DeleteDeviceAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除主机账号
func (c *Client) DeleteDeviceAccounts(request *DeleteDeviceAccountsRequest) (response *DeleteDeviceAccountsResponse, err error) {
	if request == nil {
		request = NewDeleteDeviceAccountsRequest()
	}
	response = NewDeleteDeviceAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAppAssetsDepartmentRequest() (request *ModifyAppAssetsDepartmentRequest) {
	request = &ModifyAppAssetsDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAppAssetsDepartment")
	return
}

func NewModifyAppAssetsDepartmentResponse() (response *ModifyAppAssetsDepartmentResponse) {
	response = &ModifyAppAssetsDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改应用资产部门
func (c *Client) ModifyAppAssetsDepartment(request *ModifyAppAssetsDepartmentRequest) (response *ModifyAppAssetsDepartmentResponse, err error) {
	if request == nil {
		request = NewModifyAppAssetsDepartmentRequest()
	}
	response = NewModifyAppAssetsDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUsersRequest() (request *DeleteUsersRequest) {
	request = &DeleteUsersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteUsers")
	return
}

func NewDeleteUsersResponse() (response *DeleteUsersResponse) {
	response = &DeleteUsersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户
func (c *Client) DeleteUsers(request *DeleteUsersRequest) (response *DeleteUsersResponse, err error) {
	if request == nil {
		request = NewDeleteUsersRequest()
	}
	response = NewDeleteUsersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessControlTemplateRuleRequest() (request *DeleteAccessControlTemplateRuleRequest) {
	request = &DeleteAccessControlTemplateRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAccessControlTemplateRule")
	return
}

func NewDeleteAccessControlTemplateRuleResponse() (response *DeleteAccessControlTemplateRuleResponse) {
	response = &DeleteAccessControlTemplateRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问控制模板和规则关联
func (c *Client) DeleteAccessControlTemplateRule(request *DeleteAccessControlTemplateRuleRequest) (response *DeleteAccessControlTemplateRuleResponse, err error) {
	if request == nil {
		request = NewDeleteAccessControlTemplateRuleRequest()
	}
	response = NewDeleteAccessControlTemplateRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyOAuthSettingRequest() (request *ModifyOAuthSettingRequest) {
	request = &ModifyOAuthSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyOAuthSetting")
	return
}

func NewModifyOAuthSettingResponse() (response *ModifyOAuthSettingResponse) {
	response = &ModifyOAuthSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置OAuth认证参数
func (c *Client) ModifyOAuthSetting(request *ModifyOAuthSettingRequest) (response *ModifyOAuthSettingResponse, err error) {
	if request == nil {
		request = NewModifyOAuthSettingRequest()
	}
	response = NewModifyOAuthSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserGroupsRequest() (request *DeleteUserGroupsRequest) {
	request = &DeleteUserGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteUserGroups")
	return
}

func NewDeleteUserGroupsResponse() (response *DeleteUserGroupsResponse) {
	response = &DeleteUserGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户组
func (c *Client) DeleteUserGroups(request *DeleteUserGroupsRequest) (response *DeleteUserGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteUserGroupsRequest()
	}
	response = NewDeleteUserGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDeviceRequest() (request *ModifyDeviceRequest) {
	request = &ModifyDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDevice")
	return
}

func NewModifyDeviceResponse() (response *ModifyDeviceResponse) {
	response = &ModifyDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产信息
func (c *Client) ModifyDevice(request *ModifyDeviceRequest) (response *ModifyDeviceResponse, err error) {
	if request == nil {
		request = NewModifyDeviceRequest()
	}
	response = NewModifyDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDevicesRequest() (request *DescribeDevicesRequest) {
	request = &DescribeDevicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDevices")
	return
}

func NewDescribeDevicesResponse() (response *DescribeDevicesResponse) {
	response = &DescribeDevicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产列表
func (c *Client) DescribeDevices(request *DescribeDevicesRequest) (response *DescribeDevicesResponse, err error) {
	if request == nil {
		request = NewDescribeDevicesRequest()
	}
	response = NewDescribeDevicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSystemTaskStatisticsRequest() (request *DescribeSystemTaskStatisticsRequest) {
	request = &DescribeSystemTaskStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeSystemTaskStatistics")
	return
}

func NewDescribeSystemTaskStatisticsResponse() (response *DescribeSystemTaskStatisticsResponse) {
	response = &DescribeSystemTaskStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统任务统计信息
func (c *Client) DescribeSystemTaskStatistics(request *DescribeSystemTaskStatisticsRequest) (response *DescribeSystemTaskStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeSystemTaskStatisticsRequest()
	}
	response = NewDescribeSystemTaskStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceCountSummaryRequest() (request *DescribeDeviceCountSummaryRequest) {
	request = &DescribeDeviceCountSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceCountSummary")
	return
}

func NewDescribeDeviceCountSummaryResponse() (response *DescribeDeviceCountSummaryResponse) {
	response = &DescribeDeviceCountSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户导入的主机数
func (c *Client) DescribeDeviceCountSummary(request *DescribeDeviceCountSummaryRequest) (response *DescribeDeviceCountSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceCountSummaryRequest()
	}
	response = NewDescribeDeviceCountSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDevicesPortRequest() (request *ModifyDevicesPortRequest) {
	request = &ModifyDevicesPortRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDevicesPort")
	return
}

func NewModifyDevicesPortResponse() (response *ModifyDevicesPortResponse) {
	response = &ModifyDevicesPortResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改主机端口
func (c *Client) ModifyDevicesPort(request *ModifyDevicesPortRequest) (response *ModifyDevicesPortResponse, err error) {
	if request == nil {
		request = NewModifyDevicesPortRequest()
	}
	response = NewModifyDevicesPortResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserGroupsRequest() (request *DescribeUserGroupsRequest) {
	request = &DescribeUserGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUserGroups")
	return
}

func NewDescribeUserGroupsResponse() (response *DescribeUserGroupsResponse) {
	response = &DescribeUserGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户组列表
func (c *Client) DescribeUserGroups(request *DescribeUserGroupsRequest) (response *DescribeUserGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeUserGroupsRequest()
	}
	response = NewDescribeUserGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportUserTaskRequest() (request *DescribeExportUserTaskRequest) {
	request = &DescribeExportUserTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeExportUserTask")
	return
}

func NewDescribeExportUserTaskResponse() (response *DescribeExportUserTaskResponse) {
	response = &DescribeExportUserTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户导出任务列表
func (c *Client) DescribeExportUserTask(request *DescribeExportUserTaskRequest) (response *DescribeExportUserTaskResponse, err error) {
	if request == nil {
		request = NewDescribeExportUserTaskRequest()
	}
	response = NewDescribeExportUserTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSearchFileBySidRequest() (request *SearchFileBySidRequest) {
	request = &SearchFileBySidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchFileBySid")
	return
}

func NewSearchFileBySidResponse() (response *SearchFileBySidResponse) {
	response = &SearchFileBySidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索文件传输会话下文件操作列表
func (c *Client) SearchFileBySid(request *SearchFileBySidRequest) (response *SearchFileBySidResponse, err error) {
	if request == nil {
		request = NewSearchFileBySidRequest()
	}
	response = NewSearchFileBySidResponse()
	err = c.Send(request, response)
	return
}

func NewBindAppAssetRequest() (request *BindAppAssetRequest) {
	request = &BindAppAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "BindAppAsset")
	return
}

func NewBindAppAssetResponse() (response *BindAppAssetResponse) {
	response = &BindAppAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 托管应用资产
func (c *Client) BindAppAsset(request *BindAppAssetRequest) (response *BindAppAssetResponse, err error) {
	if request == nil {
		request = NewBindAppAssetRequest()
	}
	response = NewBindAppAssetResponse()
	err = c.Send(request, response)
	return
}

func NewAccessTrackPageRequest() (request *AccessTrackPageRequest) {
	request = &AccessTrackPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "AccessTrackPage")
	return
}

func NewAccessTrackPageResponse() (response *AccessTrackPageResponse) {
	response = &AccessTrackPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 进入埋点页面，埋点数据采集
func (c *Client) AccessTrackPage(request *AccessTrackPageRequest) (response *AccessTrackPageResponse, err error) {
	if request == nil {
		request = NewAccessTrackPageRequest()
	}
	response = NewAccessTrackPageResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserBatchRequest() (request *CreateUserBatchRequest) {
	request = &CreateUserBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateUserBatch")
	return
}

func NewCreateUserBatchResponse() (response *CreateUserBatchResponse) {
	response = &CreateUserBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入用户
func (c *Client) CreateUserBatch(request *CreateUserBatchRequest) (response *CreateUserBatchResponse, err error) {
	if request == nil {
		request = NewCreateUserBatchRequest()
	}
	response = NewCreateUserBatchResponse()
	err = c.Send(request, response)
	return
}

func NewSearchCommandRequest() (request *SearchCommandRequest) {
	request = &SearchCommandRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchCommand")
	return
}

func NewSearchCommandResponse() (response *SearchCommandResponse) {
	response = &SearchCommandResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 命令执行检索
func (c *Client) SearchCommand(request *SearchCommandRequest) (response *SearchCommandResponse, err error) {
	if request == nil {
		request = NewSearchCommandRequest()
	}
	response = NewSearchCommandResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTrialGuideStepRequest() (request *UpdateTrialGuideStepRequest) {
	request = &UpdateTrialGuideStepRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "UpdateTrialGuideStep")
	return
}

func NewUpdateTrialGuideStepResponse() (response *UpdateTrialGuideStepResponse) {
	response = &UpdateTrialGuideStepResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新试用引导步骤
func (c *Client) UpdateTrialGuideStep(request *UpdateTrialGuideStepRequest) (response *UpdateTrialGuideStepResponse, err error) {
	if request == nil {
		request = NewUpdateTrialGuideStepRequest()
	}
	response = NewUpdateTrialGuideStepResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUKeyRequest() (request *ModifyUKeyRequest) {
	request = &ModifyUKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyUKey")
	return
}

func NewModifyUKeyResponse() (response *ModifyUKeyResponse) {
	response = &ModifyUKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改UKey已绑定的用户
func (c *Client) ModifyUKey(request *ModifyUKeyRequest) (response *ModifyUKeyResponse, err error) {
	if request == nil {
		request = NewModifyUKeyRequest()
	}
	response = NewModifyUKeyResponse()
	err = c.Send(request, response)
	return
}

func NewShowGraphRequest() (request *ShowGraphRequest) {
	request = &ShowGraphRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ShowGraph")
	return
}

func NewShowGraphResponse() (response *ShowGraphResponse) {
	response = &ShowGraphResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示折线图
func (c *Client) ShowGraph(request *ShowGraphRequest) (response *ShowGraphResponse, err error) {
	if request == nil {
		request = NewShowGraphRequest()
	}
	response = NewShowGraphResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReportTaskRequest() (request *DeleteReportTaskRequest) {
	request = &DeleteReportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteReportTask")
	return
}

func NewDeleteReportTaskResponse() (response *DeleteReportTaskResponse) {
	response = &DeleteReportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除审计报表任务
func (c *Client) DeleteReportTask(request *DeleteReportTaskRequest) (response *DeleteReportTaskResponse, err error) {
	if request == nil {
		request = NewDeleteReportTaskRequest()
	}
	response = NewDeleteReportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDepartmentRequest() (request *ModifyDepartmentRequest) {
	request = &ModifyDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDepartment")
	return
}

func NewModifyDepartmentResponse() (response *ModifyDepartmentResponse) {
	response = &ModifyDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改部门信息
func (c *Client) ModifyDepartment(request *ModifyDepartmentRequest) (response *ModifyDepartmentResponse, err error) {
	if request == nil {
		request = NewModifyDepartmentRequest()
	}
	response = NewModifyDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessTimePolicyRequest() (request *ModifyAccessTimePolicyRequest) {
	request = &ModifyAccessTimePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessTimePolicy")
	return
}

func NewModifyAccessTimePolicyResponse() (response *ModifyAccessTimePolicyResponse) {
	response = &ModifyAccessTimePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改访问时间段限制
func (c *Client) ModifyAccessTimePolicy(request *ModifyAccessTimePolicyRequest) (response *ModifyAccessTimePolicyResponse, err error) {
	if request == nil {
		request = NewModifyAccessTimePolicyRequest()
	}
	response = NewModifyAccessTimePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserGroupRequest() (request *ModifyUserGroupRequest) {
	request = &ModifyUserGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyUserGroup")
	return
}

func NewModifyUserGroupResponse() (response *ModifyUserGroupResponse) {
	response = &ModifyUserGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户组
func (c *Client) ModifyUserGroup(request *ModifyUserGroupRequest) (response *ModifyUserGroupResponse, err error) {
	if request == nil {
		request = NewModifyUserGroupRequest()
	}
	response = NewModifyUserGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExportAuditLogTaskRequest() (request *DeleteExportAuditLogTaskRequest) {
	request = &DeleteExportAuditLogTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteExportAuditLogTask")
	return
}

func NewDeleteExportAuditLogTaskResponse() (response *DeleteExportAuditLogTaskResponse) {
	response = &DeleteExportAuditLogTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除审计日志导出任务
func (c *Client) DeleteExportAuditLogTask(request *DeleteExportAuditLogTaskRequest) (response *DeleteExportAuditLogTaskResponse, err error) {
	if request == nil {
		request = NewDeleteExportAuditLogTaskRequest()
	}
	response = NewDeleteExportAuditLogTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeviceAccountBatchRequest() (request *CreateDeviceAccountBatchRequest) {
	request = &CreateDeviceAccountBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDeviceAccountBatch")
	return
}

func NewCreateDeviceAccountBatchResponse() (response *CreateDeviceAccountBatchResponse) {
	response = &CreateDeviceAccountBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量添加主机账号
func (c *Client) CreateDeviceAccountBatch(request *CreateDeviceAccountBatchRequest) (response *CreateDeviceAccountBatchResponse, err error) {
	if request == nil {
		request = NewCreateDeviceAccountBatchRequest()
	}
	response = NewCreateDeviceAccountBatchResponse()
	err = c.Send(request, response)
	return
}

func NewSearchChangePwdTaskInfoRequest() (request *SearchChangePwdTaskInfoRequest) {
	request = &SearchChangePwdTaskInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchChangePwdTaskInfo")
	return
}

func NewSearchChangePwdTaskInfoResponse() (response *SearchChangePwdTaskInfoResponse) {
	response = &SearchChangePwdTaskInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索改密任务详情
func (c *Client) SearchChangePwdTaskInfo(request *SearchChangePwdTaskInfoRequest) (response *SearchChangePwdTaskInfoResponse, err error) {
	if request == nil {
		request = NewSearchChangePwdTaskInfoRequest()
	}
	response = NewSearchChangePwdTaskInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDomainRequest() (request *CreateDomainRequest) {
	request = &CreateDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDomain")
	return
}

func NewCreateDomainResponse() (response *CreateDomainResponse) {
	response = &CreateDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加网络域
func (c *Client) CreateDomain(request *CreateDomainRequest) (response *CreateDomainResponse, err error) {
	if request == nil {
		request = NewCreateDomainRequest()
	}
	response = NewCreateDomainResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSessionRequest() (request *SearchSessionRequest) {
	request = &SearchSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchSession")
	return
}

func NewSearchSessionResponse() (response *SearchSessionResponse) {
	response = &SearchSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索会话
func (c *Client) SearchSession(request *SearchSessionRequest) (response *SearchSessionResponse, err error) {
	if request == nil {
		request = NewSearchSessionRequest()
	}
	response = NewSearchSessionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessWhiteListRuleRequest() (request *CreateAccessWhiteListRuleRequest) {
	request = &CreateAccessWhiteListRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAccessWhiteListRule")
	return
}

func NewCreateAccessWhiteListRuleResponse() (response *CreateAccessWhiteListRuleResponse) {
	response = &CreateAccessWhiteListRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加访问白名单规则
func (c *Client) CreateAccessWhiteListRule(request *CreateAccessWhiteListRuleRequest) (response *CreateAccessWhiteListRuleResponse, err error) {
	if request == nil {
		request = NewCreateAccessWhiteListRuleRequest()
	}
	response = NewCreateAccessWhiteListRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLDAPSettingRequest() (request *ModifyLDAPSettingRequest) {
	request = &ModifyLDAPSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyLDAPSetting")
	return
}

func NewModifyLDAPSettingResponse() (response *ModifyLDAPSettingResponse) {
	response = &ModifyLDAPSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改LDAP配置信息
func (c *Client) ModifyLDAPSetting(request *ModifyLDAPSettingRequest) (response *ModifyLDAPSettingResponse, err error) {
	if request == nil {
		request = NewModifyLDAPSettingRequest()
	}
	response = NewModifyLDAPSettingResponse()
	err = c.Send(request, response)
	return
}

func NewModifyChangePwdTaskRequest() (request *ModifyChangePwdTaskRequest) {
	request = &ModifyChangePwdTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyChangePwdTask")
	return
}

func NewModifyChangePwdTaskResponse() (response *ModifyChangePwdTaskResponse) {
	response = &ModifyChangePwdTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新修改密码任务
func (c *Client) ModifyChangePwdTask(request *ModifyChangePwdTaskRequest) (response *ModifyChangePwdTaskResponse, err error) {
	if request == nil {
		request = NewModifyChangePwdTaskRequest()
	}
	response = NewModifyChangePwdTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogOutputSettingsRequest() (request *ModifyLogOutputSettingsRequest) {
	request = &ModifyLogOutputSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyLogOutputSettings")
	return
}

func NewModifyLogOutputSettingsResponse() (response *ModifyLogOutputSettingsResponse) {
	response = &ModifyLogOutputSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改日志外发配置
func (c *Client) ModifyLogOutputSettings(request *ModifyLogOutputSettingsRequest) (response *ModifyLogOutputSettingsResponse, err error) {
	if request == nil {
		request = NewModifyLogOutputSettingsRequest()
	}
	response = NewModifyLogOutputSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAuthModeSettingRequest() (request *ModifyAuthModeSettingRequest) {
	request = &ModifyAuthModeSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAuthModeSetting")
	return
}

func NewModifyAuthModeSettingResponse() (response *ModifyAuthModeSettingResponse) {
	response = &ModifyAuthModeSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改认证方式配置信息
func (c *Client) ModifyAuthModeSetting(request *ModifyAuthModeSettingRequest) (response *ModifyAuthModeSettingResponse, err error) {
	if request == nil {
		request = NewModifyAuthModeSettingRequest()
	}
	response = NewModifyAuthModeSettingResponse()
	err = c.Send(request, response)
	return
}

func NewResetLogDeliveryRequest() (request *ResetLogDeliveryRequest) {
	request = &ResetLogDeliveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetLogDelivery")
	return
}

func NewResetLogDeliveryResponse() (response *ResetLogDeliveryResponse) {
	response = &ResetLogDeliveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置日志投递
func (c *Client) ResetLogDelivery(request *ResetLogDeliveryRequest) (response *ResetLogDeliveryResponse, err error) {
	if request == nil {
		request = NewResetLogDeliveryRequest()
	}
	response = NewResetLogDeliveryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogOutputSettingsRequest() (request *DescribeLogOutputSettingsRequest) {
	request = &DescribeLogOutputSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLogOutputSettings")
	return
}

func NewDescribeLogOutputSettingsResponse() (response *DescribeLogOutputSettingsResponse) {
	response = &DescribeLogOutputSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志外发配置
func (c *Client) DescribeLogOutputSettings(request *DescribeLogOutputSettingsRequest) (response *DescribeLogOutputSettingsResponse, err error) {
	if request == nil {
		request = NewDescribeLogOutputSettingsRequest()
	}
	response = NewDescribeLogOutputSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewEnableIntranetAccessRequest() (request *EnableIntranetAccessRequest) {
	request = &EnableIntranetAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "EnableIntranetAccess")
	return
}

func NewEnableIntranetAccessResponse() (response *EnableIntranetAccessResponse) {
	response = &EnableIntranetAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开通内网访问
func (c *Client) EnableIntranetAccess(request *EnableIntranetAccessRequest) (response *EnableIntranetAccessResponse, err error) {
	if request == nil {
		request = NewEnableIntranetAccessRequest()
	}
	response = NewEnableIntranetAccessResponse()
	err = c.Send(request, response)
	return
}

func NewLoginOpserverRequest() (request *LoginOpserverRequest) {
	request = &LoginOpserverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "LoginOpserver")
	return
}

func NewLoginOpserverResponse() (response *LoginOpserverResponse) {
	response = &LoginOpserverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 控制台切换运维端
func (c *Client) LoginOpserver(request *LoginOpserverRequest) (response *LoginOpserverResponse, err error) {
	if request == nil {
		request = NewLoginOpserverRequest()
	}
	response = NewLoginOpserverResponse()
	err = c.Send(request, response)
	return
}

func NewSyncUserFromCamRequest() (request *SyncUserFromCamRequest) {
	request = &SyncUserFromCamRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SyncUserFromCam")
	return
}

func NewSyncUserFromCamResponse() (response *SyncUserFromCamResponse) {
	response = &SyncUserFromCamResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步cam用户
func (c *Client) SyncUserFromCam(request *SyncUserFromCamRequest) (response *SyncUserFromCamResponse, err error) {
	if request == nil {
		request = NewSyncUserFromCamRequest()
	}
	response = NewSyncUserFromCamResponse()
	err = c.Send(request, response)
	return
}

func NewDisconnectDomainRequest() (request *DisconnectDomainRequest) {
	request = &DisconnectDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DisconnectDomain")
	return
}

func NewDisconnectDomainResponse() (response *DisconnectDomainResponse) {
	response = &DisconnectDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 断开网络域连接
func (c *Client) DisconnectDomain(request *DisconnectDomainRequest) (response *DisconnectDomainResponse, err error) {
	if request == nil {
		request = NewDisconnectDomainRequest()
	}
	response = NewDisconnectDomainResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCmdTemplatesRequest() (request *DescribeCmdTemplatesRequest) {
	request = &DescribeCmdTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeCmdTemplates")
	return
}

func NewDescribeCmdTemplatesResponse() (response *DescribeCmdTemplatesResponse) {
	response = &DescribeCmdTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询命令模板列表
func (c *Client) DescribeCmdTemplates(request *DescribeCmdTemplatesRequest) (response *DescribeCmdTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeCmdTemplatesRequest()
	}
	response = NewDescribeCmdTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceGroupMembersRequest() (request *DescribeDeviceGroupMembersRequest) {
	request = &DescribeDeviceGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceGroupMembers")
	return
}

func NewDescribeDeviceGroupMembersResponse() (response *DescribeDeviceGroupMembersResponse) {
	response = &DescribeDeviceGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产组成员列表
func (c *Client) DescribeDeviceGroupMembers(request *DescribeDeviceGroupMembersRequest) (response *DescribeDeviceGroupMembersResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceGroupMembersRequest()
	}
	response = NewDescribeDeviceGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReportTaskRequest() (request *DescribeReportTaskRequest) {
	request = &DescribeReportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeReportTask")
	return
}

func NewDescribeReportTaskResponse() (response *DescribeReportTaskResponse) {
	response = &DescribeReportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询审计报表任务
func (c *Client) DescribeReportTask(request *DescribeReportTaskRequest) (response *DescribeReportTaskResponse, err error) {
	if request == nil {
		request = NewDescribeReportTaskRequest()
	}
	response = NewDescribeReportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
	request = &CreateUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateUser")
	return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
	response = &CreateUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建用户
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
	if request == nil {
		request = NewCreateUserRequest()
	}
	response = NewCreateUserResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAclsRequest() (request *DeleteAclsRequest) {
	request = &DeleteAclsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAcls")
	return
}

func NewDeleteAclsResponse() (response *DeleteAclsResponse) {
	response = &DeleteAclsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问权限
func (c *Client) DeleteAcls(request *DeleteAclsRequest) (response *DeleteAclsResponse, err error) {
	if request == nil {
		request = NewDeleteAclsRequest()
	}
	response = NewDeleteAclsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUKeysRequest() (request *DescribeUKeysRequest) {
	request = &DescribeUKeysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUKeys")
	return
}

func NewDescribeUKeysResponse() (response *DescribeUKeysResponse) {
	response = &DescribeUKeysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询UKey和绑定用户列表
func (c *Client) DescribeUKeys(request *DescribeUKeysRequest) (response *DescribeUKeysResponse, err error) {
	if request == nil {
		request = NewDescribeUKeysRequest()
	}
	response = NewDescribeUKeysResponse()
	err = c.Send(request, response)
	return
}

func NewCheckLDAPConnectionRequest() (request *CheckLDAPConnectionRequest) {
	request = &CheckLDAPConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CheckLDAPConnection")
	return
}

func NewCheckLDAPConnectionResponse() (response *CheckLDAPConnectionResponse) {
	response = &CheckLDAPConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试LDAP连接
func (c *Client) CheckLDAPConnection(request *CheckLDAPConnectionRequest) (response *CheckLDAPConnectionResponse, err error) {
	if request == nil {
		request = NewCheckLDAPConnectionRequest()
	}
	response = NewCheckLDAPConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeviceGroupRequest() (request *CreateDeviceGroupRequest) {
	request = &CreateDeviceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDeviceGroup")
	return
}

func NewCreateDeviceGroupResponse() (response *CreateDeviceGroupResponse) {
	response = &CreateDeviceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建资产组
func (c *Client) CreateDeviceGroup(request *CreateDeviceGroupRequest) (response *CreateDeviceGroupResponse, err error) {
	if request == nil {
		request = NewCreateDeviceGroupRequest()
	}
	response = NewCreateDeviceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmSettingRequest() (request *DescribeAlarmSettingRequest) {
	request = &DescribeAlarmSettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAlarmSetting")
	return
}

func NewDescribeAlarmSettingResponse() (response *DescribeAlarmSettingResponse) {
	response = &DescribeAlarmSettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询时间告警设置
func (c *Client) DescribeAlarmSetting(request *DescribeAlarmSettingRequest) (response *DescribeAlarmSettingResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmSettingRequest()
	}
	response = NewDescribeAlarmSettingResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTypeRequest() (request *DescribeOperationTypeRequest) {
	request = &DescribeOperationTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationType")
	return
}

func NewDescribeOperationTypeResponse() (response *DescribeOperationTypeResponse) {
	response = &DescribeOperationTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作类型
func (c *Client) DescribeOperationType(request *DescribeOperationTypeRequest) (response *DescribeOperationTypeResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTypeRequest()
	}
	response = NewDescribeOperationTypeResponse()
	err = c.Send(request, response)
	return
}

func NewBindDeviceAccountPrivateKeyRequest() (request *BindDeviceAccountPrivateKeyRequest) {
	request = &BindDeviceAccountPrivateKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "BindDeviceAccountPrivateKey")
	return
}

func NewBindDeviceAccountPrivateKeyResponse() (response *BindDeviceAccountPrivateKeyResponse) {
	response = &BindDeviceAccountPrivateKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定主机账号私钥
func (c *Client) BindDeviceAccountPrivateKey(request *BindDeviceAccountPrivateKeyRequest) (response *BindDeviceAccountPrivateKeyResponse, err error) {
	if request == nil {
		request = NewBindDeviceAccountPrivateKeyRequest()
	}
	response = NewBindDeviceAccountPrivateKeyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDepartmentRequest() (request *CreateDepartmentRequest) {
	request = &CreateDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDepartment")
	return
}

func NewCreateDepartmentResponse() (response *CreateDepartmentResponse) {
	response = &CreateDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建部门
func (c *Client) CreateDepartment(request *CreateDepartmentRequest) (response *CreateDepartmentResponse, err error) {
	if request == nil {
		request = NewCreateDepartmentRequest()
	}
	response = NewCreateDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUKeyBatchRequest() (request *CreateUKeyBatchRequest) {
	request = &CreateUKeyBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateUKeyBatch")
	return
}

func NewCreateUKeyBatchResponse() (response *CreateUKeyBatchResponse) {
	response = &CreateUKeyBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入UKey绑定用户
func (c *Client) CreateUKeyBatch(request *CreateUKeyBatchRequest) (response *CreateUKeyBatchResponse, err error) {
	if request == nil {
		request = NewCreateUKeyBatchRequest()
	}
	response = NewCreateUKeyBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSyncFlagRequest() (request *DescribeAssetSyncFlagRequest) {
	request = &DescribeAssetSyncFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAssetSyncFlag")
	return
}

func NewDescribeAssetSyncFlagResponse() (response *DescribeAssetSyncFlagResponse) {
	response = &DescribeAssetSyncFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产自动同步开关
func (c *Client) DescribeAssetSyncFlag(request *DescribeAssetSyncFlagRequest) (response *DescribeAssetSyncFlagResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSyncFlagRequest()
	}
	response = NewDescribeAssetSyncFlagResponse()
	err = c.Send(request, response)
	return
}

func NewResetDeviceAccountPasswordRequest() (request *ResetDeviceAccountPasswordRequest) {
	request = &ResetDeviceAccountPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetDeviceAccountPassword")
	return
}

func NewResetDeviceAccountPasswordResponse() (response *ResetDeviceAccountPasswordResponse) {
	response = &ResetDeviceAccountPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清除设备账号绑定密码
func (c *Client) ResetDeviceAccountPassword(request *ResetDeviceAccountPasswordRequest) (response *ResetDeviceAccountPasswordResponse, err error) {
	if request == nil {
		request = NewResetDeviceAccountPasswordRequest()
	}
	response = NewResetDeviceAccountPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssetSyncStatusRequest() (request *DescribeAssetSyncStatusRequest) {
	request = &DescribeAssetSyncStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAssetSyncStatus")
	return
}

func NewDescribeAssetSyncStatusResponse() (response *DescribeAssetSyncStatusResponse) {
	response = &DescribeAssetSyncStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产同步状态
func (c *Client) DescribeAssetSyncStatus(request *DescribeAssetSyncStatusRequest) (response *DescribeAssetSyncStatusResponse, err error) {
	if request == nil {
		request = NewDescribeAssetSyncStatusRequest()
	}
	response = NewDescribeAssetSyncStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseUsageRequest() (request *DescribeLicenseUsageRequest) {
	request = &DescribeLicenseUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLicenseUsage")
	return
}

func NewDescribeLicenseUsageResponse() (response *DescribeLicenseUsageResponse) {
	response = &DescribeLicenseUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 供license平台用于查询license用量
func (c *Client) DescribeLicenseUsage(request *DescribeLicenseUsageRequest) (response *DescribeLicenseUsageResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseUsageRequest()
	}
	response = NewDescribeLicenseUsageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTicketSubmitFlagRequest() (request *ModifyTicketSubmitFlagRequest) {
	request = &ModifyTicketSubmitFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyTicketSubmitFlag")
	return
}

func NewModifyTicketSubmitFlagResponse() (response *ModifyTicketSubmitFlagResponse) {
	response = &ModifyTicketSubmitFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改权限工单自动提交设置
func (c *Client) ModifyTicketSubmitFlag(request *ModifyTicketSubmitFlagRequest) (response *ModifyTicketSubmitFlagResponse, err error) {
	if request == nil {
		request = NewModifyTicketSubmitFlagRequest()
	}
	response = NewModifyTicketSubmitFlagResponse()
	err = c.Send(request, response)
	return
}

func NewSearchTaskResultRequest() (request *SearchTaskResultRequest) {
	request = &SearchTaskResultRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchTaskResult")
	return
}

func NewSearchTaskResultResponse() (response *SearchTaskResultResponse) {
	response = &SearchTaskResultResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索运维任务执行结果
func (c *Client) SearchTaskResult(request *SearchTaskResultRequest) (response *SearchTaskResultResponse, err error) {
	if request == nil {
		request = NewSearchTaskResultRequest()
	}
	response = NewSearchTaskResultResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCmdTemplateRequest() (request *CreateCmdTemplateRequest) {
	request = &CreateCmdTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateCmdTemplate")
	return
}

func NewCreateCmdTemplateResponse() (response *CreateCmdTemplateResponse) {
	response = &CreateCmdTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建高危命令模板
func (c *Client) CreateCmdTemplate(request *CreateCmdTemplateRequest) (response *CreateCmdTemplateResponse, err error) {
	if request == nil {
		request = NewCreateCmdTemplateRequest()
	}
	response = NewCreateCmdTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeChangePwdTaskDetailRequest() (request *DescribeChangePwdTaskDetailRequest) {
	request = &DescribeChangePwdTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeChangePwdTaskDetail")
	return
}

func NewDescribeChangePwdTaskDetailResponse() (response *DescribeChangePwdTaskDetailResponse) {
	response = &DescribeChangePwdTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询改密任务详情
func (c *Client) DescribeChangePwdTaskDetail(request *DescribeChangePwdTaskDetailRequest) (response *DescribeChangePwdTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeChangePwdTaskDetailRequest()
	}
	response = NewDescribeChangePwdTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewBindDeviceAccountPasswordRequest() (request *BindDeviceAccountPasswordRequest) {
	request = &BindDeviceAccountPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "BindDeviceAccountPassword")
	return
}

func NewBindDeviceAccountPasswordResponse() (response *BindDeviceAccountPasswordResponse) {
	response = &BindDeviceAccountPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定主机账号密码
func (c *Client) BindDeviceAccountPassword(request *BindDeviceAccountPasswordRequest) (response *BindDeviceAccountPasswordResponse, err error) {
	if request == nil {
		request = NewBindDeviceAccountPasswordRequest()
	}
	response = NewBindDeviceAccountPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShareClbIpsRequest() (request *DescribeShareClbIpsRequest) {
	request = &DescribeShareClbIpsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeShareClbIps")
	return
}

func NewDescribeShareClbIpsResponse() (response *DescribeShareClbIpsResponse) {
	response = &DescribeShareClbIpsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询开通堡垒机可以使用的共享clb
func (c *Client) DescribeShareClbIps(request *DescribeShareClbIpsRequest) (response *DescribeShareClbIpsResponse, err error) {
	if request == nil {
		request = NewDescribeShareClbIpsRequest()
	}
	response = NewDescribeShareClbIpsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAppAssetGroupMembersRequest() (request *DeleteAppAssetGroupMembersRequest) {
	request = &DeleteAppAssetGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAppAssetGroupMembers")
	return
}

func NewDeleteAppAssetGroupMembersResponse() (response *DeleteAppAssetGroupMembersResponse) {
	response = &DeleteAppAssetGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用资产组成员
func (c *Client) DeleteAppAssetGroupMembers(request *DeleteAppAssetGroupMembersRequest) (response *DeleteAppAssetGroupMembersResponse, err error) {
	if request == nil {
		request = NewDeleteAppAssetGroupMembersRequest()
	}
	response = NewDeleteAppAssetGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessWhiteListAutoStatusRequest() (request *ModifyAccessWhiteListAutoStatusRequest) {
	request = &ModifyAccessWhiteListAutoStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessWhiteListAutoStatus")
	return
}

func NewModifyAccessWhiteListAutoStatusResponse() (response *ModifyAccessWhiteListAutoStatusResponse) {
	response = &ModifyAccessWhiteListAutoStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问白名单自动添加IP状态：开启或关闭自动添加IP
func (c *Client) ModifyAccessWhiteListAutoStatus(request *ModifyAccessWhiteListAutoStatusRequest) (response *ModifyAccessWhiteListAutoStatusResponse, err error) {
	if request == nil {
		request = NewModifyAccessWhiteListAutoStatusRequest()
	}
	response = NewModifyAccessWhiteListAutoStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTaskStatisticsRequest() (request *DescribeOperationTaskStatisticsRequest) {
	request = &DescribeOperationTaskStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationTaskStatistics")
	return
}

func NewDescribeOperationTaskStatisticsResponse() (response *DescribeOperationTaskStatisticsResponse) {
	response = &DescribeOperationTaskStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维任务统计信息
func (c *Client) DescribeOperationTaskStatistics(request *DescribeOperationTaskStatisticsRequest) (response *DescribeOperationTaskStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTaskStatisticsRequest()
	}
	response = NewDescribeOperationTaskStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteChangePwdTaskRequest() (request *DeleteChangePwdTaskRequest) {
	request = &DeleteChangePwdTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteChangePwdTask")
	return
}

func NewDeleteChangePwdTaskResponse() (response *DeleteChangePwdTaskResponse) {
	response = &DeleteChangePwdTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除修改密码任务
func (c *Client) DeleteChangePwdTask(request *DeleteChangePwdTaskRequest) (response *DeleteChangePwdTaskResponse, err error) {
	if request == nil {
		request = NewDeleteChangePwdTaskRequest()
	}
	response = NewDeleteChangePwdTaskResponse()
	err = c.Send(request, response)
	return
}

func NewRunChangePwdTaskRequest() (request *RunChangePwdTaskRequest) {
	request = &RunChangePwdTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "RunChangePwdTask")
	return
}

func NewRunChangePwdTaskResponse() (response *RunChangePwdTaskResponse) {
	response = &RunChangePwdTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行改密任务
func (c *Client) RunChangePwdTask(request *RunChangePwdTaskRequest) (response *RunChangePwdTaskResponse, err error) {
	if request == nil {
		request = NewRunChangePwdTaskRequest()
	}
	response = NewRunChangePwdTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSearchFileSessionRequest() (request *SearchFileSessionRequest) {
	request = &SearchFileSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchFileSession")
	return
}

func NewSearchFileSessionResponse() (response *SearchFileSessionResponse) {
	response = &SearchFileSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索传输的文件情况和所属的会话
func (c *Client) SearchFileSession(request *SearchFileSessionRequest) (response *SearchFileSessionResponse, err error) {
	if request == nil {
		request = NewSearchFileSessionRequest()
	}
	response = NewSearchFileSessionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainRequest() (request *ModifyDomainRequest) {
	request = &ModifyDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDomain")
	return
}

func NewModifyDomainResponse() (response *ModifyDomainResponse) {
	response = &ModifyDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加网络域
func (c *Client) ModifyDomain(request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
	if request == nil {
		request = NewModifyDomainRequest()
	}
	response = NewModifyDomainResponse()
	err = c.Send(request, response)
	return
}

func NewSearchStatementBySidRequest() (request *SearchStatementBySidRequest) {
	request = &SearchStatementBySidRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchStatementBySid")
	return
}

func NewSearchStatementBySidResponse() (response *SearchStatementBySidResponse) {
	response = &SearchStatementBySidResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据会话Id搜索数据库操作
func (c *Client) SearchStatementBySid(request *SearchStatementBySidRequest) (response *SearchStatementBySidResponse, err error) {
	if request == nil {
		request = NewSearchStatementBySidRequest()
	}
	response = NewSearchStatementBySidResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPushAccountTaskRequest() (request *ModifyPushAccountTaskRequest) {
	request = &ModifyPushAccountTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyPushAccountTask")
	return
}

func NewModifyPushAccountTaskResponse() (response *ModifyPushAccountTaskResponse) {
	response = &ModifyPushAccountTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改账号推送任务信息
func (c *Client) ModifyPushAccountTask(request *ModifyPushAccountTaskRequest) (response *ModifyPushAccountTaskResponse, err error) {
	if request == nil {
		request = NewModifyPushAccountTaskRequest()
	}
	response = NewModifyPushAccountTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDestroyDasbResourceRequest() (request *DestroyDasbResourceRequest) {
	request = &DestroyDasbResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DestroyDasbResource")
	return
}

func NewDestroyDasbResourceResponse() (response *DestroyDasbResourceResponse) {
	response = &DestroyDasbResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁堡垒机资源接口
func (c *Client) DestroyDasbResource(request *DestroyDasbResourceRequest) (response *DestroyDasbResourceResponse, err error) {
	if request == nil {
		request = NewDestroyDasbResourceRequest()
	}
	response = NewDestroyDasbResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainsRequest() (request *DescribeDomainsRequest) {
	request = &DescribeDomainsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDomains")
	return
}

func NewDescribeDomainsResponse() (response *DescribeDomainsResponse) {
	response = &DescribeDomainsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络域
func (c *Client) DescribeDomains(request *DescribeDomainsRequest) (response *DescribeDomainsResponse, err error) {
	if request == nil {
		request = NewDescribeDomainsRequest()
	}
	response = NewDescribeDomainsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessControlTemplateRequest() (request *ModifyAccessControlTemplateRequest) {
	request = &ModifyAccessControlTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessControlTemplate")
	return
}

func NewModifyAccessControlTemplateResponse() (response *ModifyAccessControlTemplateResponse) {
	response = &ModifyAccessControlTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问控制模板
func (c *Client) ModifyAccessControlTemplate(request *ModifyAccessControlTemplateRequest) (response *ModifyAccessControlTemplateResponse, err error) {
	if request == nil {
		request = NewModifyAccessControlTemplateRequest()
	}
	response = NewModifyAccessControlTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeChangePwdTaskRequest() (request *DescribeChangePwdTaskRequest) {
	request = &DescribeChangePwdTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeChangePwdTask")
	return
}

func NewDescribeChangePwdTaskResponse() (response *DescribeChangePwdTaskResponse) {
	response = &DescribeChangePwdTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询改密任务列表
func (c *Client) DescribeChangePwdTask(request *DescribeChangePwdTaskRequest) (response *DescribeChangePwdTaskResponse, err error) {
	if request == nil {
		request = NewDescribeChangePwdTaskRequest()
	}
	response = NewDescribeChangePwdTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceCountRequest() (request *DescribeDeviceCountRequest) {
	request = &DescribeDeviceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceCount")
	return
}

func NewDescribeDeviceCountResponse() (response *DescribeDeviceCountResponse) {
	response = &DescribeDeviceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户导入的主机数
func (c *Client) DescribeDeviceCount(request *DescribeDeviceCountRequest) (response *DescribeDeviceCountResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceCountRequest()
	}
	response = NewDescribeDeviceCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDeviceGroupRequest() (request *ModifyDeviceGroupRequest) {
	request = &ModifyDeviceGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDeviceGroup")
	return
}

func NewModifyDeviceGroupResponse() (response *ModifyDeviceGroupResponse) {
	response = &ModifyDeviceGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产组
func (c *Client) ModifyDeviceGroup(request *ModifyDeviceGroupRequest) (response *ModifyDeviceGroupResponse, err error) {
	if request == nil {
		request = NewModifyDeviceGroupRequest()
	}
	response = NewModifyDeviceGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAclRequest() (request *ModifyAclRequest) {
	request = &ModifyAclRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAcl")
	return
}

func NewModifyAclResponse() (response *ModifyAclResponse) {
	response = &ModifyAclResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问权限
func (c *Client) ModifyAcl(request *ModifyAclRequest) (response *ModifyAclResponse, err error) {
	if request == nil {
		request = NewModifyAclRequest()
	}
	response = NewModifyAclResponse()
	err = c.Send(request, response)
	return
}

func NewDeployResourceRequest() (request *DeployResourceRequest) {
	request = &DeployResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeployResource")
	return
}

func NewDeployResourceResponse() (response *DeployResourceResponse) {
	response = &DeployResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开通服务，初始化资源，只针对新购资源
func (c *Client) DeployResource(request *DeployResourceRequest) (response *DeployResourceResponse, err error) {
	if request == nil {
		request = NewDeployResourceRequest()
	}
	response = NewDeployResourceResponse()
	err = c.Send(request, response)
	return
}

func NewConnectDomainRequest() (request *ConnectDomainRequest) {
	request = &ConnectDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ConnectDomain")
	return
}

func NewConnectDomainResponse() (response *ConnectDomainResponse) {
	response = &ConnectDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 建立网络域连接
func (c *Client) ConnectDomain(request *ConnectDomainRequest) (response *ConnectDomainResponse, err error) {
	if request == nil {
		request = NewConnectDomainRequest()
	}
	response = NewConnectDomainResponse()
	err = c.Send(request, response)
	return
}

func NewSetLDAPSyncFlagRequest() (request *SetLDAPSyncFlagRequest) {
	request = &SetLDAPSyncFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SetLDAPSyncFlag")
	return
}

func NewSetLDAPSyncFlagResponse() (response *SetLDAPSyncFlagResponse) {
	response = &SetLDAPSyncFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置LDAP 立即同步标记
func (c *Client) SetLDAPSyncFlag(request *SetLDAPSyncFlagRequest) (response *SetLDAPSyncFlagResponse, err error) {
	if request == nil {
		request = NewSetLDAPSyncFlagRequest()
	}
	response = NewSetLDAPSyncFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationEventRequest() (request *DescribeOperationEventRequest) {
	request = &DescribeOperationEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationEvent")
	return
}

func NewDescribeOperationEventResponse() (response *DescribeOperationEventResponse) {
	response = &DescribeOperationEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作日志
func (c *Client) DescribeOperationEvent(request *DescribeOperationEventRequest) (response *DescribeOperationEventResponse, err error) {
	if request == nil {
		request = NewDescribeOperationEventRequest()
	}
	response = NewDescribeOperationEventResponse()
	err = c.Send(request, response)
	return
}

func NewSearchStatementRequest() (request *SearchStatementRequest) {
	request = &SearchStatementRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchStatement")
	return
}

func NewSearchStatementResponse() (response *SearchStatementResponse) {
	response = &SearchStatementResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索数据库操作
func (c *Client) SearchStatement(request *SearchStatementRequest) (response *SearchStatementResponse, err error) {
	if request == nil {
		request = NewSearchStatementRequest()
	}
	response = NewSearchStatementResponse()
	err = c.Send(request, response)
	return
}

func NewImportExternalDeviceRequest() (request *ImportExternalDeviceRequest) {
	request = &ImportExternalDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ImportExternalDevice")
	return
}

func NewImportExternalDeviceResponse() (response *ImportExternalDeviceResponse) {
	response = &ImportExternalDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入外部资产信息
func (c *Client) ImportExternalDevice(request *ImportExternalDeviceRequest) (response *ImportExternalDeviceResponse, err error) {
	if request == nil {
		request = NewImportExternalDeviceRequest()
	}
	response = NewImportExternalDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlTemplatesRequest() (request *DescribeAccessControlTemplatesRequest) {
	request = &DescribeAccessControlTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccessControlTemplates")
	return
}

func NewDescribeAccessControlTemplatesResponse() (response *DescribeAccessControlTemplatesResponse) {
	response = &DescribeAccessControlTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问控制模板列表
func (c *Client) DescribeAccessControlTemplates(request *DescribeAccessControlTemplatesRequest) (response *DescribeAccessControlTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlTemplatesRequest()
	}
	response = NewDescribeAccessControlTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAclRequest() (request *CreateAclRequest) {
	request = &CreateAclRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAcl")
	return
}

func NewCreateAclResponse() (response *CreateAclResponse) {
	response = &CreateAclResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建访问权限
func (c *Client) CreateAcl(request *CreateAclRequest) (response *CreateAclResponse, err error) {
	if request == nil {
		request = NewCreateAclRequest()
	}
	response = NewCreateAclResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecuritySettingRequest() (request *DescribeSecuritySettingRequest) {
	request = &DescribeSecuritySettingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeSecuritySetting")
	return
}

func NewDescribeSecuritySettingResponse() (response *DescribeSecuritySettingResponse) {
	response = &DescribeSecuritySettingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询安全配置信息
func (c *Client) DescribeSecuritySetting(request *DescribeSecuritySettingRequest) (response *DescribeSecuritySettingResponse, err error) {
	if request == nil {
		request = NewDescribeSecuritySettingRequest()
	}
	response = NewDescribeSecuritySettingResponse()
	err = c.Send(request, response)
	return
}

func NewInquireCreateDasbResourceRequest() (request *InquireCreateDasbResourceRequest) {
	request = &InquireCreateDasbResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "InquireCreateDasbResource")
	return
}

func NewInquireCreateDasbResourceResponse() (response *InquireCreateDasbResourceResponse) {
	response = &InquireCreateDasbResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费模式下,【新购堡垒机询价】和【变配询价】都是使用该接口。此接口返回变配后总价。
func (c *Client) InquireCreateDasbResource(request *InquireCreateDasbResourceRequest) (response *InquireCreateDasbResourceResponse, err error) {
	if request == nil {
		request = NewInquireCreateDasbResourceRequest()
	}
	response = NewInquireCreateDasbResourceResponse()
	err = c.Send(request, response)
	return
}

func NewSearchCommandSessionRequest() (request *SearchCommandSessionRequest) {
	request = &SearchCommandSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchCommandSession")
	return
}

func NewSearchCommandSessionResponse() (response *SearchCommandSessionResponse) {
	response = &SearchCommandSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检索命令和所属会话
func (c *Client) SearchCommandSession(request *SearchCommandSessionRequest) (response *SearchCommandSessionResponse, err error) {
	if request == nil {
		request = NewSearchCommandSessionRequest()
	}
	response = NewSearchCommandSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskTemplateRequest() (request *DescribeTaskTemplateRequest) {
	request = &DescribeTaskTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeTaskTemplate")
	return
}

func NewDescribeTaskTemplateResponse() (response *DescribeTaskTemplateResponse) {
	response = &DescribeTaskTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询系统任务模板
func (c *Client) DescribeTaskTemplate(request *DescribeTaskTemplateRequest) (response *DescribeTaskTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeTaskTemplateRequest()
	}
	response = NewDescribeTaskTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserCountRequest() (request *DescribeUserCountRequest) {
	request = &DescribeUserCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUserCount")
	return
}

func NewDescribeUserCountResponse() (response *DescribeUserCountResponse) {
	response = &DescribeUserCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户总数
func (c *Client) DescribeUserCount(request *DescribeUserCountRequest) (response *DescribeUserCountResponse, err error) {
	if request == nil {
		request = NewDescribeUserCountRequest()
	}
	response = NewDescribeUserCountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateResourceRequest() (request *CreateResourceRequest) {
	request = &CreateResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateResource")
	return
}

func NewCreateResourceResponse() (response *CreateResourceResponse) {
	response = &CreateResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建堡垒机实例
func (c *Client) CreateResource(request *CreateResourceRequest) (response *CreateResourceResponse, err error) {
	if request == nil {
		request = NewCreateResourceRequest()
	}
	response = NewCreateResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeviceGroupsRequest() (request *DescribeDeviceGroupsRequest) {
	request = &DescribeDeviceGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDeviceGroups")
	return
}

func NewDescribeDeviceGroupsResponse() (response *DescribeDeviceGroupsResponse) {
	response = &DescribeDeviceGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产组列表
func (c *Client) DescribeDeviceGroups(request *DescribeDeviceGroupsRequest) (response *DescribeDeviceGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDeviceGroupsRequest()
	}
	response = NewDescribeDeviceGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessWhiteListRuleRequest() (request *ModifyAccessWhiteListRuleRequest) {
	request = &ModifyAccessWhiteListRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessWhiteListRule")
	return
}

func NewModifyAccessWhiteListRuleResponse() (response *ModifyAccessWhiteListRuleResponse) {
	response = &ModifyAccessWhiteListRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问白名单规则
func (c *Client) ModifyAccessWhiteListRule(request *ModifyAccessWhiteListRuleRequest) (response *ModifyAccessWhiteListRuleResponse, err error) {
	if request == nil {
		request = NewModifyAccessWhiteListRuleRequest()
	}
	response = NewModifyAccessWhiteListRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExportAuditLogTaskRequest() (request *CreateExportAuditLogTaskRequest) {
	request = &CreateExportAuditLogTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateExportAuditLogTask")
	return
}

func NewCreateExportAuditLogTaskResponse() (response *CreateExportAuditLogTaskResponse) {
	response = &CreateExportAuditLogTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建审计日志导出任务
func (c *Client) CreateExportAuditLogTask(request *CreateExportAuditLogTaskRequest) (response *CreateExportAuditLogTaskResponse, err error) {
	if request == nil {
		request = NewCreateExportAuditLogTaskRequest()
	}
	response = NewCreateExportAuditLogTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportDeviceTaskRequest() (request *DescribeExportDeviceTaskRequest) {
	request = &DescribeExportDeviceTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeExportDeviceTask")
	return
}

func NewDescribeExportDeviceTaskResponse() (response *DescribeExportDeviceTaskResponse) {
	response = &DescribeExportDeviceTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资产导出任务列表
func (c *Client) DescribeExportDeviceTask(request *DescribeExportDeviceTaskRequest) (response *DescribeExportDeviceTaskResponse, err error) {
	if request == nil {
		request = NewDescribeExportDeviceTaskRequest()
	}
	response = NewDescribeExportDeviceTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAccessControlTemplateRuleOrderRequest() (request *ModifyAccessControlTemplateRuleOrderRequest) {
	request = &ModifyAccessControlTemplateRuleOrderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAccessControlTemplateRuleOrder")
	return
}

func NewModifyAccessControlTemplateRuleOrderResponse() (response *ModifyAccessControlTemplateRuleOrderResponse) {
	response = &ModifyAccessControlTemplateRuleOrderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改访问控制模板规则顺序
func (c *Client) ModifyAccessControlTemplateRuleOrder(request *ModifyAccessControlTemplateRuleOrderRequest) (response *ModifyAccessControlTemplateRuleOrderResponse, err error) {
	if request == nil {
		request = NewModifyAccessControlTemplateRuleOrderRequest()
	}
	response = NewModifyAccessControlTemplateRuleOrderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceAllRequest() (request *DescribeResourceAllRequest) {
	request = &DescribeResourceAllRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeResourceAll")
	return
}

func NewDescribeResourceAllResponse() (response *DescribeResourceAllResponse) {
	response = &DescribeResourceAllResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询堡垒机所有租户的实例
func (c *Client) DescribeResourceAll(request *DescribeResourceAllRequest) (response *DescribeResourceAllResponse, err error) {
	if request == nil {
		request = NewDescribeResourceAllRequest()
	}
	response = NewDescribeResourceAllResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDevicesDepartmentRequest() (request *ModifyDevicesDepartmentRequest) {
	request = &ModifyDevicesDepartmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyDevicesDepartment")
	return
}

func NewModifyDevicesDepartmentResponse() (response *ModifyDevicesDepartmentResponse) {
	response = &ModifyDevicesDepartmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量修改资产部门
func (c *Client) ModifyDevicesDepartment(request *ModifyDevicesDepartmentRequest) (response *ModifyDevicesDepartmentResponse, err error) {
	if request == nil {
		request = NewModifyDevicesDepartmentRequest()
	}
	response = NewModifyDevicesDepartmentResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeviceAccountRequest() (request *CreateDeviceAccountRequest) {
	request = &CreateDeviceAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateDeviceAccount")
	return
}

func NewCreateDeviceAccountResponse() (response *CreateDeviceAccountResponse) {
	response = &CreateDeviceAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建主机账号
func (c *Client) CreateDeviceAccount(request *CreateDeviceAccountRequest) (response *CreateDeviceAccountResponse, err error) {
	if request == nil {
		request = NewCreateDeviceAccountRequest()
	}
	response = NewCreateDeviceAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlRulesRequest() (request *DescribeAccessControlRulesRequest) {
	request = &DescribeAccessControlRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccessControlRules")
	return
}

func NewDescribeAccessControlRulesResponse() (response *DescribeAccessControlRulesResponse) {
	response = &DescribeAccessControlRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取访问控制策略列表
func (c *Client) DescribeAccessControlRules(request *DescribeAccessControlRulesRequest) (response *DescribeAccessControlRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlRulesRequest()
	}
	response = NewDescribeAccessControlRulesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCmdTemplateRequest() (request *ModifyCmdTemplateRequest) {
	request = &ModifyCmdTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyCmdTemplate")
	return
}

func NewModifyCmdTemplateResponse() (response *ModifyCmdTemplateResponse) {
	response = &ModifyCmdTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改高危命令模板
func (c *Client) ModifyCmdTemplate(request *ModifyCmdTemplateRequest) (response *ModifyCmdTemplateResponse, err error) {
	if request == nil {
		request = NewModifyCmdTemplateRequest()
	}
	response = NewModifyCmdTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewVisitTrackPageRequest() (request *VisitTrackPageRequest) {
	request = &VisitTrackPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "VisitTrackPage")
	return
}

func NewVisitTrackPageResponse() (response *VisitTrackPageResponse) {
	response = &VisitTrackPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 访问埋点数据采集页面
func (c *Client) VisitTrackPage(request *VisitTrackPageRequest) (response *VisitTrackPageResponse, err error) {
	if request == nil {
		request = NewVisitTrackPageRequest()
	}
	response = NewVisitTrackPageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessControlTemplateRulesRequest() (request *DescribeAccessControlTemplateRulesRequest) {
	request = &DescribeAccessControlTemplateRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccessControlTemplateRules")
	return
}

func NewDescribeAccessControlTemplateRulesResponse() (response *DescribeAccessControlTemplateRulesResponse) {
	response = &DescribeAccessControlTemplateRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询访问控制模板关联的规则列表
func (c *Client) DescribeAccessControlTemplateRules(request *DescribeAccessControlTemplateRulesRequest) (response *DescribeAccessControlTemplateRulesResponse, err error) {
	if request == nil {
		request = NewDescribeAccessControlTemplateRulesRequest()
	}
	response = NewDescribeAccessControlTemplateRulesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchTaskResultDetailRequest() (request *SearchTaskResultDetailRequest) {
	request = &SearchTaskResultDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchTaskResultDetail")
	return
}

func NewSearchTaskResultDetailResponse() (response *SearchTaskResultDetailResponse) {
	response = &SearchTaskResultDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索运维任务执行结果详情
func (c *Client) SearchTaskResultDetail(request *SearchTaskResultDetailRequest) (response *SearchTaskResultDetailResponse, err error) {
	if request == nil {
		request = NewSearchTaskResultDetailRequest()
	}
	response = NewSearchTaskResultDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOperationTaskDetailRequest() (request *DescribeOperationTaskDetailRequest) {
	request = &DescribeOperationTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeOperationTaskDetail")
	return
}

func NewDescribeOperationTaskDetailResponse() (response *DescribeOperationTaskDetailResponse) {
	response = &DescribeOperationTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维任务详情
func (c *Client) DescribeOperationTaskDetail(request *DescribeOperationTaskDetailRequest) (response *DescribeOperationTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeOperationTaskDetailRequest()
	}
	response = NewDescribeOperationTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReportTaskRequest() (request *ModifyReportTaskRequest) {
	request = &ModifyReportTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyReportTask")
	return
}

func NewModifyReportTaskResponse() (response *ModifyReportTaskResponse) {
	response = &ModifyReportTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改审计报表任务
func (c *Client) ModifyReportTask(request *ModifyReportTaskRequest) (response *ModifyReportTaskResponse, err error) {
	if request == nil {
		request = NewModifyReportTaskRequest()
	}
	response = NewModifyReportTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCanCreateShareClbResourceRequest() (request *CanCreateShareClbResourceRequest) {
	request = &CanCreateShareClbResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CanCreateShareClbResource")
	return
}

func NewCanCreateShareClbResourceResponse() (response *CanCreateShareClbResourceResponse) {
	response = &CanCreateShareClbResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询能否购买创建共享clb的堡垒机
func (c *Client) CanCreateShareClbResource(request *CanCreateShareClbResourceRequest) (response *CanCreateShareClbResourceResponse, err error) {
	if request == nil {
		request = NewCanCreateShareClbResourceRequest()
	}
	response = NewCanCreateShareClbResourceResponse()
	err = c.Send(request, response)
	return
}

func NewBindDeviceResourceRequest() (request *BindDeviceResourceRequest) {
	request = &BindDeviceResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "BindDeviceResource")
	return
}

func NewBindDeviceResourceResponse() (response *BindDeviceResourceResponse) {
	response = &BindDeviceResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产绑定的堡垒机服务
func (c *Client) BindDeviceResource(request *BindDeviceResourceRequest) (response *BindDeviceResourceResponse, err error) {
	if request == nil {
		request = NewBindDeviceResourceRequest()
	}
	response = NewBindDeviceResourceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAccessControlTemplateRuleRequest() (request *CreateAccessControlTemplateRuleRequest) {
	request = &CreateAccessControlTemplateRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAccessControlTemplateRule")
	return
}

func NewCreateAccessControlTemplateRuleResponse() (response *CreateAccessControlTemplateRuleResponse) {
	response = &CreateAccessControlTemplateRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建访问控制模板规则关联
func (c *Client) CreateAccessControlTemplateRule(request *CreateAccessControlTemplateRuleRequest) (response *CreateAccessControlTemplateRuleResponse, err error) {
	if request == nil {
		request = NewCreateAccessControlTemplateRuleRequest()
	}
	response = NewCreateAccessControlTemplateRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCkafkaInstanceListRequest() (request *DescribeCkafkaInstanceListRequest) {
	request = &DescribeCkafkaInstanceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeCkafkaInstanceList")
	return
}

func NewDescribeCkafkaInstanceListResponse() (response *DescribeCkafkaInstanceListResponse) {
	response = &DescribeCkafkaInstanceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询ckafka实例列表
func (c *Client) DescribeCkafkaInstanceList(request *DescribeCkafkaInstanceListRequest) (response *DescribeCkafkaInstanceListResponse, err error) {
	if request == nil {
		request = NewDescribeCkafkaInstanceListRequest()
	}
	response = NewDescribeCkafkaInstanceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTicketSubmitFlagRequest() (request *DescribeTicketSubmitFlagRequest) {
	request = &DescribeTicketSubmitFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeTicketSubmitFlag")
	return
}

func NewDescribeTicketSubmitFlagResponse() (response *DescribeTicketSubmitFlagResponse) {
	response = &DescribeTicketSubmitFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询权限工单是否自动提交，不自动提交则是草稿状态，需要再运维端手动提交
func (c *Client) DescribeTicketSubmitFlag(request *DescribeTicketSubmitFlagRequest) (response *DescribeTicketSubmitFlagResponse, err error) {
	if request == nil {
		request = NewDescribeTicketSubmitFlagRequest()
	}
	response = NewDescribeTicketSubmitFlagResponse()
	err = c.Send(request, response)
	return
}

func NewImportDeviceAccountRequest() (request *ImportDeviceAccountRequest) {
	request = &ImportDeviceAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ImportDeviceAccount")
	return
}

func NewImportDeviceAccountResponse() (response *ImportDeviceAccountResponse) {
	response = &ImportDeviceAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入账号
func (c *Client) ImportDeviceAccount(request *ImportDeviceAccountRequest) (response *ImportDeviceAccountResponse, err error) {
	if request == nil {
		request = NewImportDeviceAccountRequest()
	}
	response = NewImportDeviceAccountResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUKeyRequest() (request *CreateUKeyRequest) {
	request = &CreateUKeyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateUKey")
	return
}

func NewCreateUKeyResponse() (response *CreateUKeyResponse) {
	response = &CreateUKeyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定UKey和用户
func (c *Client) CreateUKey(request *CreateUKeyRequest) (response *CreateUKeyResponse, err error) {
	if request == nil {
		request = NewCreateUKeyRequest()
	}
	response = NewCreateUKeyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLDAPUnitSetRequest() (request *DescribeLDAPUnitSetRequest) {
	request = &DescribeLDAPUnitSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLDAPUnitSet")
	return
}

func NewDescribeLDAPUnitSetResponse() (response *DescribeLDAPUnitSetResponse) {
	response = &DescribeLDAPUnitSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取LDAP ou 列表
func (c *Client) DescribeLDAPUnitSet(request *DescribeLDAPUnitSetRequest) (response *DescribeLDAPUnitSetResponse, err error) {
	if request == nil {
		request = NewDescribeLDAPUnitSetRequest()
	}
	response = NewDescribeLDAPUnitSetResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePushAccountTasksRequest() (request *DeletePushAccountTasksRequest) {
	request = &DeletePushAccountTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeletePushAccountTasks")
	return
}

func NewDeletePushAccountTasksResponse() (response *DeletePushAccountTasksResponse) {
	response = &DeletePushAccountTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除账号推送任务
func (c *Client) DeletePushAccountTasks(request *DeletePushAccountTasksRequest) (response *DeletePushAccountTasksResponse, err error) {
	if request == nil {
		request = NewDeletePushAccountTasksRequest()
	}
	response = NewDeletePushAccountTasksResponse()
	err = c.Send(request, response)
	return
}

func NewResetUserPasswordRequest() (request *ResetUserPasswordRequest) {
	request = &ResetUserPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetUserPassword")
	return
}

func NewResetUserPasswordResponse() (response *ResetUserPasswordResponse) {
	response = &ResetUserPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置用户密码（已废弃，由ResetUser取代）
func (c *Client) ResetUserPassword(request *ResetUserPasswordRequest) (response *ResetUserPasswordResponse, err error) {
	if request == nil {
		request = NewResetUserPasswordRequest()
	}
	response = NewResetUserPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewResetUserRequest() (request *ResetUserRequest) {
	request = &ResetUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ResetUser")
	return
}

func NewResetUserResponse() (response *ResetUserResponse) {
	response = &ResetUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置用户
func (c *Client) ResetUser(request *ResetUserRequest) (response *ResetUserResponse, err error) {
	if request == nil {
		request = NewResetUserRequest()
	}
	response = NewResetUserResponse()
	err = c.Send(request, response)
	return
}

func NewSearchKeyboardLoggerRequest() (request *SearchKeyboardLoggerRequest) {
	request = &SearchKeyboardLoggerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchKeyboardLogger")
	return
}

func NewSearchKeyboardLoggerResponse() (response *SearchKeyboardLoggerResponse) {
	response = &SearchKeyboardLoggerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 键盘记录事件列表
func (c *Client) SearchKeyboardLogger(request *SearchKeyboardLoggerRequest) (response *SearchKeyboardLoggerResponse, err error) {
	if request == nil {
		request = NewSearchKeyboardLoggerRequest()
	}
	response = NewSearchKeyboardLoggerResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAppAssetRequest() (request *CreateAppAssetRequest) {
	request = &CreateAppAssetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "CreateAppAsset")
	return
}

func NewCreateAppAssetResponse() (response *CreateAppAssetResponse) {
	response = &CreateAppAssetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建或导入应用资产信息
func (c *Client) CreateAppAsset(request *CreateAppAssetRequest) (response *CreateAppAssetResponse, err error) {
	if request == nil {
		request = NewCreateAppAssetRequest()
	}
	response = NewCreateAppAssetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExternalDeviceRequest() (request *ModifyExternalDeviceRequest) {
	request = &ModifyExternalDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyExternalDevice")
	return
}

func NewModifyExternalDeviceResponse() (response *ModifyExternalDeviceResponse) {
	response = &ModifyExternalDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改外部资产信息
func (c *Client) ModifyExternalDevice(request *ModifyExternalDeviceRequest) (response *ModifyExternalDeviceResponse, err error) {
	if request == nil {
		request = NewModifyExternalDeviceRequest()
	}
	response = NewModifyExternalDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewAccessDeviceRequest() (request *AccessDeviceRequest) {
	request = &AccessDeviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "AccessDevice")
	return
}

func NewAccessDeviceResponse() (response *AccessDeviceResponse) {
	response = &AccessDeviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 访问资产
func (c *Client) AccessDevice(request *AccessDeviceRequest) (response *AccessDeviceResponse, err error) {
	if request == nil {
		request = NewAccessDeviceRequest()
	}
	response = NewAccessDeviceResponse()
	err = c.Send(request, response)
	return
}

func NewAddAppAssetGroupMembersRequest() (request *AddAppAssetGroupMembersRequest) {
	request = &AddAppAssetGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "AddAppAssetGroupMembers")
	return
}

func NewAddAppAssetGroupMembersResponse() (response *AddAppAssetGroupMembersResponse) {
	response = &AddAppAssetGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加应用资产组成员
func (c *Client) AddAppAssetGroupMembers(request *AddAppAssetGroupMembersRequest) (response *AddAppAssetGroupMembersResponse, err error) {
	if request == nil {
		request = NewAddAppAssetGroupMembersRequest()
	}
	response = NewAddAppAssetGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogDeliveryRequest() (request *DescribeLogDeliveryRequest) {
	request = &DescribeLogDeliveryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeLogDelivery")
	return
}

func NewDescribeLogDeliveryResponse() (response *DescribeLogDeliveryResponse) {
	response = &DescribeLogDeliveryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志推送信息
func (c *Client) DescribeLogDelivery(request *DescribeLogDeliveryRequest) (response *DescribeLogDeliveryResponse, err error) {
	if request == nil {
		request = NewDescribeLogDeliveryRequest()
	}
	response = NewDescribeLogDeliveryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeviceGroupMembersRequest() (request *DeleteDeviceGroupMembersRequest) {
	request = &DeleteDeviceGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDeviceGroupMembers")
	return
}

func NewDeleteDeviceGroupMembersResponse() (response *DeleteDeviceGroupMembersResponse) {
	response = &DeleteDeviceGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资产组成员
func (c *Client) DeleteDeviceGroupMembers(request *DeleteDeviceGroupMembersRequest) (response *DeleteDeviceGroupMembersResponse, err error) {
	if request == nil {
		request = NewDeleteDeviceGroupMembersRequest()
	}
	response = NewDeleteDeviceGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAccessControlTemplateRequest() (request *DeleteAccessControlTemplateRequest) {
	request = &DeleteAccessControlTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteAccessControlTemplate")
	return
}

func NewDeleteAccessControlTemplateResponse() (response *DeleteAccessControlTemplateResponse) {
	response = &DeleteAccessControlTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除访问控制模板
func (c *Client) DeleteAccessControlTemplate(request *DeleteAccessControlTemplateRequest) (response *DeleteAccessControlTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteAccessControlTemplateRequest()
	}
	response = NewDeleteAccessControlTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAssetSyncFlagRequest() (request *ModifyAssetSyncFlagRequest) {
	request = &ModifyAssetSyncFlagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ModifyAssetSyncFlag")
	return
}

func NewModifyAssetSyncFlagResponse() (response *ModifyAssetSyncFlagResponse) {
	response = &ModifyAssetSyncFlagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改资产自动同步开关
func (c *Client) ModifyAssetSyncFlag(request *ModifyAssetSyncFlagRequest) (response *ModifyAssetSyncFlagResponse, err error) {
	if request == nil {
		request = NewModifyAssetSyncFlagRequest()
	}
	response = NewModifyAssetSyncFlagResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeviceGroupsRequest() (request *DeleteDeviceGroupsRequest) {
	request = &DeleteDeviceGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteDeviceGroups")
	return
}

func NewDeleteDeviceGroupsResponse() (response *DeleteDeviceGroupsResponse) {
	response = &DeleteDeviceGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除资产组
func (c *Client) DeleteDeviceGroups(request *DeleteDeviceGroupsRequest) (response *DeleteDeviceGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteDeviceGroupsRequest()
	}
	response = NewDeleteDeviceGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDomainInstallScriptRequest() (request *DescribeDomainInstallScriptRequest) {
	request = &DescribeDomainInstallScriptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeDomainInstallScript")
	return
}

func NewDescribeDomainInstallScriptResponse() (response *DescribeDomainInstallScriptResponse) {
	response = &DescribeDomainInstallScriptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网络域安装脚本
func (c *Client) DescribeDomainInstallScript(request *DescribeDomainInstallScriptRequest) (response *DescribeDomainInstallScriptResponse, err error) {
	if request == nil {
		request = NewDescribeDomainInstallScriptRequest()
	}
	response = NewDescribeDomainInstallScriptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserGroupMembersRequest() (request *DescribeUserGroupMembersRequest) {
	request = &DescribeUserGroupMembersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeUserGroupMembers")
	return
}

func NewDescribeUserGroupMembersResponse() (response *DescribeUserGroupMembersResponse) {
	response = &DescribeUserGroupMembersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户组成员列表
func (c *Client) DescribeUserGroupMembers(request *DescribeUserGroupMembersRequest) (response *DescribeUserGroupMembersResponse, err error) {
	if request == nil {
		request = NewDescribeUserGroupMembersRequest()
	}
	response = NewDescribeUserGroupMembersResponse()
	err = c.Send(request, response)
	return
}

func NewShowTopRequest() (request *ShowTopRequest) {
	request = &ShowTopRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "ShowTop")
	return
}

func NewShowTopResponse() (response *ShowTopResponse) {
	response = &ShowTopResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示在线用户数、活跃用户、活跃主机、高危用户
func (c *Client) ShowTop(request *ShowTopRequest) (response *ShowTopResponse, err error) {
	if request == nil {
		request = NewShowTopRequest()
	}
	response = NewShowTopResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccessEntryRequest() (request *DescribeAccessEntryRequest) {
	request = &DescribeAccessEntryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DescribeAccessEntry")
	return
}

func NewDescribeAccessEntryResponse() (response *DescribeAccessEntryResponse) {
	response = &DescribeAccessEntryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询运维人员WEB访问入口
func (c *Client) DescribeAccessEntry(request *DescribeAccessEntryRequest) (response *DescribeAccessEntryResponse, err error) {
	if request == nil {
		request = NewDescribeAccessEntryRequest()
	}
	response = NewDescribeAccessEntryResponse()
	err = c.Send(request, response)
	return
}

func NewMonitorSessionRequest() (request *MonitorSessionRequest) {
	request = &MonitorSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "MonitorSession")
	return
}

func NewMonitorSessionResponse() (response *MonitorSessionResponse) {
	response = &MonitorSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 会话实时监控
func (c *Client) MonitorSession(request *MonitorSessionRequest) (response *MonitorSessionResponse, err error) {
	if request == nil {
		request = NewMonitorSessionRequest()
	}
	response = NewMonitorSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReportTaskHistoryRequest() (request *DeleteReportTaskHistoryRequest) {
	request = &DeleteReportTaskHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "DeleteReportTaskHistory")
	return
}

func NewDeleteReportTaskHistoryResponse() (response *DeleteReportTaskHistoryResponse) {
	response = &DeleteReportTaskHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除审计报表任务记录
func (c *Client) DeleteReportTaskHistory(request *DeleteReportTaskHistoryRequest) (response *DeleteReportTaskHistoryResponse, err error) {
	if request == nil {
		request = NewDeleteReportTaskHistoryRequest()
	}
	response = NewDeleteReportTaskHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSessionCommandRequest() (request *SearchSessionCommandRequest) {
	request = &SearchSessionCommandRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("bhsaas", APIVersion, "SearchSessionCommand")
	return
}

func NewSearchSessionCommandResponse() (response *SearchSessionCommandResponse) {
	response = &SearchSessionCommandResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 命令检索
func (c *Client) SearchSessionCommand(request *SearchSessionCommandRequest) (response *SearchSessionCommandResponse, err error) {
	if request == nil {
		request = NewSearchSessionCommandRequest()
	}
	response = NewSearchSessionCommandResponse()
	err = c.Send(request, response)
	return
}

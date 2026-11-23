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

package v20190107

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-01-07"

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

func NewDescribeInstanceTaskStatusRequest() (request *DescribeInstanceTaskStatusRequest) {
	request = &DescribeInstanceTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeInstanceTaskStatus")
	return
}

func NewDescribeInstanceTaskStatusResponse() (response *DescribeInstanceTaskStatusResponse) {
	response = &DescribeInstanceTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询实力相关的异步流程任务的执行状态，可获知任务是否失败或成功
func (c *Client) DescribeInstanceTaskStatus(request *DescribeInstanceTaskStatusRequest) (response *DescribeInstanceTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceTaskStatusRequest()
	}
	response = NewDescribeInstanceTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupListsRequest() (request *DescribeBackupListsRequest) {
	request = &DescribeBackupListsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeBackupLists")
	return
}

func NewDescribeBackupListsResponse() (response *DescribeBackupListsResponse) {
	response = &DescribeBackupListsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeBackupLists）用于查询实例的备份列表。
func (c *Client) DescribeBackupLists(request *DescribeBackupListsRequest) (response *DescribeBackupListsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupListsRequest()
	}
	response = NewDescribeBackupListsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePGErrorLogRequest() (request *DescribePGErrorLogRequest) {
	request = &DescribePGErrorLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribePGErrorLog")
	return
}

func NewDescribePGErrorLogResponse() (response *DescribePGErrorLogResponse) {
	response = &DescribePGErrorLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribePGErrorLog) 用于查询一个PG实例的错误日志。
func (c *Client) DescribePGErrorLog(request *DescribePGErrorLogRequest) (response *DescribePGErrorLogResponse, err error) {
	if request == nil {
		request = NewDescribePGErrorLogRequest()
	}
	response = NewDescribePGErrorLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupRulesRequest() (request *DescribeBackupRulesRequest) {
	request = &DescribeBackupRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeBackupRules")
	return
}

func NewDescribeBackupRulesResponse() (response *DescribeBackupRulesResponse) {
	response = &DescribeBackupRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeBackupRules）用于查询实例备份规则。
func (c *Client) DescribeBackupRules(request *DescribeBackupRulesRequest) (response *DescribeBackupRulesResponse, err error) {
	if request == nil {
		request = NewDescribeBackupRulesRequest()
	}
	response = NewDescribeBackupRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableRecoveryNodeSpecRequest() (request *DescribeAvailableRecoveryNodeSpecRequest) {
	request = &DescribeAvailableRecoveryNodeSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableRecoveryNodeSpec")
	return
}

func NewDescribeAvailableRecoveryNodeSpecResponse() (response *DescribeAvailableRecoveryNodeSpecResponse) {
	response = &DescribeAvailableRecoveryNodeSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例回档规格
func (c *Client) DescribeAvailableRecoveryNodeSpec(request *DescribeAvailableRecoveryNodeSpecRequest) (response *DescribeAvailableRecoveryNodeSpecResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableRecoveryNodeSpecRequest()
	}
	response = NewDescribeAvailableRecoveryNodeSpecResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDbNameListRequest() (request *DescribeDbNameListRequest) {
	request = &DescribeDbNameListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeDbNameList")
	return
}

func NewDescribeDbNameListResponse() (response *DescribeDbNameListResponse) {
	response = &DescribeDbNameListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeDbNameList) 用于查询一个实例的db名称。
func (c *Client) DescribeDbNameList(request *DescribeDbNameListRequest) (response *DescribeDbNameListResponse, err error) {
	if request == nil {
		request = NewDescribeDbNameListRequest()
	}
	response = NewDescribeDbNameListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSlowLogRequest() (request *DescribeSlowLogRequest) {
	request = &DescribeSlowLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeSlowLog")
	return
}

func NewDescribeSlowLogResponse() (response *DescribeSlowLogResponse) {
	response = &DescribeSlowLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeSlowLog) 用于查询一个实例的慢查询日志。
func (c *Client) DescribeSlowLog(request *DescribeSlowLogRequest) (response *DescribeSlowLogResponse, err error) {
	if request == nil {
		request = NewDescribeSlowLogRequest()
	}
	response = NewDescribeSlowLogResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateInstancesRequest() (request *TerminateInstancesRequest) {
	request = &TerminateInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "TerminateInstances")
	return
}

func NewTerminateInstancesResponse() (response *TerminateInstancesResponse) {
	response = &TerminateInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (TerminateInstances) 用于隔离实例，目前仅支持单个实例隔离。
func (c *Client) TerminateInstances(request *TerminateInstancesRequest) (response *TerminateInstancesResponse, err error) {
	if request == nil {
		request = NewTerminateInstancesRequest()
	}
	response = NewTerminateInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewActivateIsolatedInstanceHourRequest() (request *ActivateIsolatedInstanceHourRequest) {
	request = &ActivateIsolatedInstanceHourRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ActivateIsolatedInstanceHour")
	return
}

func NewActivateIsolatedInstanceHourResponse() (response *ActivateIsolatedInstanceHourResponse) {
	response = &ActivateIsolatedInstanceHourResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ActivateIsolatedInstanceHour) 用于已隔离后付费实例重新上线。
func (c *Client) ActivateIsolatedInstanceHour(request *ActivateIsolatedInstanceHourRequest) (response *ActivateIsolatedInstanceHourResponse, err error) {
	if request == nil {
		request = NewActivateIsolatedInstanceHourRequest()
	}
	response = NewActivateIsolatedInstanceHourResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePGReadOnlyVipRequest() (request *CreatePGReadOnlyVipRequest) {
	request = &CreatePGReadOnlyVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "CreatePGReadOnlyVip")
	return
}

func NewCreatePGReadOnlyVipResponse() (response *CreatePGReadOnlyVipResponse) {
	response = &CreatePGReadOnlyVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（CreatePGReadOnlyVip）用于给postgres实例添加只读VIP，同一实例只允许存在一个只读VIP，VIP可由用户自己指定或系统自行指定。
func (c *Client) CreatePGReadOnlyVip(request *CreatePGReadOnlyVipRequest) (response *CreatePGReadOnlyVipResponse, err error) {
	if request == nil {
		request = NewCreatePGReadOnlyVipRequest()
	}
	response = NewCreatePGReadOnlyVipResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableRecoveryTimeRequest() (request *DescribeAvailableRecoveryTimeRequest) {
	request = &DescribeAvailableRecoveryTimeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableRecoveryTime")
	return
}

func NewDescribeAvailableRecoveryTimeResponse() (response *DescribeAvailableRecoveryTimeResponse) {
	response = &DescribeAvailableRecoveryTimeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 数据恢复时，获取回档时间
func (c *Client) DescribeAvailableRecoveryTime(request *DescribeAvailableRecoveryTimeRequest) (response *DescribeAvailableRecoveryTimeResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableRecoveryTimeRequest()
	}
	response = NewDescribeAvailableRecoveryTimeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBConfigHistoryRequest() (request *DescribeDBConfigHistoryRequest) {
	request = &DescribeDBConfigHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeDBConfigHistory")
	return
}

func NewDescribeDBConfigHistoryResponse() (response *DescribeDBConfigHistoryResponse) {
	response = &DescribeDBConfigHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeDBConfigHistory) 用于查询一个实例的参数修改历史记录。
func (c *Client) DescribeDBConfigHistory(request *DescribeDBConfigHistoryRequest) (response *DescribeDBConfigHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeDBConfigHistoryRequest()
	}
	response = NewDescribeDBConfigHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseIsolatedPGInstanceRequest() (request *ReleaseIsolatedPGInstanceRequest) {
	request = &ReleaseIsolatedPGInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ReleaseIsolatedPGInstance")
	return
}

func NewReleaseIsolatedPGInstanceResponse() (response *ReleaseIsolatedPGInstanceResponse) {
	response = &ReleaseIsolatedPGInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁已隔离的Postgres实例
func (c *Client) ReleaseIsolatedPGInstance(request *ReleaseIsolatedPGInstanceRequest) (response *ReleaseIsolatedPGInstanceResponse, err error) {
	if request == nil {
		request = NewReleaseIsolatedPGInstanceRequest()
	}
	response = NewReleaseIsolatedPGInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableEngineVersionsRequest() (request *DescribeAvailableEngineVersionsRequest) {
	request = &DescribeAvailableEngineVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableEngineVersions")
	return
}

func NewDescribeAvailableEngineVersionsResponse() (response *DescribeAvailableEngineVersionsResponse) {
	response = &DescribeAvailableEngineVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAvailableEngineVersions) 用于查询创建实例时可选择的引擎版本。
func (c *Client) DescribeAvailableEngineVersions(request *DescribeAvailableEngineVersionsRequest) (response *DescribeAvailableEngineVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableEngineVersionsRequest()
	}
	response = NewDescribeAvailableEngineVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableZoneConfigRequest() (request *DescribeAvailableZoneConfigRequest) {
	request = &DescribeAvailableZoneConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableZoneConfig")
	return
}

func NewDescribeAvailableZoneConfigResponse() (response *DescribeAvailableZoneConfigResponse) {
	response = &DescribeAvailableZoneConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeAvailableZoneConfig)用于查询云数据库可售卖的地域和可用区信息。
func (c *Client) DescribeAvailableZoneConfig(request *DescribeAvailableZoneConfigRequest) (response *DescribeAvailableZoneConfigResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableZoneConfigRequest()
	}
	response = NewDescribeAvailableZoneConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceHourRequest() (request *ModifyInstanceHourRequest) {
	request = &ModifyInstanceHourRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ModifyInstanceHour")
	return
}

func NewModifyInstanceHourResponse() (response *ModifyInstanceHourResponse) {
	response = &ModifyInstanceHourResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 后付费变配实例
func (c *Client) ModifyInstanceHour(request *ModifyInstanceHourRequest) (response *ModifyInstanceHourResponse, err error) {
	if request == nil {
		request = NewModifyInstanceHourRequest()
	}
	response = NewModifyInstanceHourResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableRecoveryTimePointRequest() (request *DescribeAvailableRecoveryTimePointRequest) {
	request = &DescribeAvailableRecoveryTimePointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableRecoveryTimePoint")
	return
}

func NewDescribeAvailableRecoveryTimePointResponse() (response *DescribeAvailableRecoveryTimePointResponse) {
	response = &DescribeAvailableRecoveryTimePointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAvailableRecoveryTimePoint）用于数据恢复时获取可用时间点
func (c *Client) DescribeAvailableRecoveryTimePoint(request *DescribeAvailableRecoveryTimePointRequest) (response *DescribeAvailableRecoveryTimePointResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableRecoveryTimePointRequest()
	}
	response = NewDescribeAvailableRecoveryTimePointResponse()
	err = c.Send(request, response)
	return
}

func NewModifyVipRequest() (request *ModifyVipRequest) {
	request = &ModifyVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ModifyVip")
	return
}

func NewModifyVipResponse() (response *ModifyVipResponse) {
	response = &ModifyVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyVip) 用于修改实例的内网地址。
func (c *Client) ModifyVip(request *ModifyVipRequest) (response *ModifyVipResponse, err error) {
	if request == nil {
		request = NewModifyVipRequest()
	}
	response = NewModifyVipResponse()
	err = c.Send(request, response)
	return
}

func NewIsolatePGInstanceHourRequest() (request *IsolatePGInstanceHourRequest) {
	request = &IsolatePGInstanceHourRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "IsolatePGInstanceHour")
	return
}

func NewIsolatePGInstanceHourResponse() (response *IsolatePGInstanceHourResponse) {
	response = &IsolatePGInstanceHourResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (IsolatePGInstanceHour) 用于隔离后付费Postgres实例。
func (c *Client) IsolatePGInstanceHour(request *IsolatePGInstanceHourRequest) (response *IsolatePGInstanceHourResponse, err error) {
	if request == nil {
		request = NewIsolatePGInstanceHourRequest()
	}
	response = NewIsolatePGInstanceHourResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountsRequest() (request *DescribeAccountsRequest) {
	request = &DescribeAccountsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAccounts")
	return
}

func NewDescribeAccountsResponse() (response *DescribeAccountsResponse) {
	response = &DescribeAccountsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeAccounts）用于查询用户列表。
func (c *Client) DescribeAccounts(request *DescribeAccountsRequest) (response *DescribeAccountsResponse, err error) {
	if request == nil {
		request = NewDescribeAccountsRequest()
	}
	response = NewDescribeAccountsResponse()
	err = c.Send(request, response)
	return
}

func NewSecurityGroupsRequest() (request *SecurityGroupsRequest) {
	request = &SecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "SecurityGroups")
	return
}

func NewSecurityGroupsResponse() (response *SecurityGroupsResponse) {
	response = &SecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据筛选查看安全组
func (c *Client) SecurityGroups(request *SecurityGroupsRequest) (response *SecurityGroupsResponse, err error) {
	if request == nil {
		request = NewSecurityGroupsRequest()
	}
	response = NewSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDBParametersRequest() (request *ModifyDBParametersRequest) {
	request = &ModifyDBParametersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ModifyDBParameters")
	return
}

func NewModifyDBParametersResponse() (response *ModifyDBParametersResponse) {
	response = &ModifyDBParametersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyDBParameters) 用于修改实例的参数配置。
func (c *Client) ModifyDBParameters(request *ModifyDBParametersRequest) (response *ModifyDBParametersResponse, err error) {
	if request == nil {
		request = NewModifyDBParametersRequest()
	}
	response = NewModifyDBParametersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeErrorLogRequest() (request *DescribeErrorLogRequest) {
	request = &DescribeErrorLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeErrorLog")
	return
}

func NewDescribeErrorLogResponse() (response *DescribeErrorLogResponse) {
	response = &DescribeErrorLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeErrorLog) 用于查询一个实例的错误日志。
func (c *Client) DescribeErrorLog(request *DescribeErrorLogRequest) (response *DescribeErrorLogResponse, err error) {
	if request == nil {
		request = NewDescribeErrorLogRequest()
	}
	response = NewDescribeErrorLogResponse()
	err = c.Send(request, response)
	return
}

func NewSetBackupRulesRequest() (request *SetBackupRulesRequest) {
	request = &SetBackupRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "SetBackupRules")
	return
}

func NewSetBackupRulesResponse() (response *SetBackupRulesResponse) {
	response = &SetBackupRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（SetBackupRules）用于设置实例备份规则。
func (c *Client) SetBackupRules(request *SetBackupRulesRequest) (response *SetBackupRulesResponse, err error) {
	if request == nil {
		request = NewSetBackupRulesRequest()
	}
	response = NewSetBackupRulesResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePGInstanceHourRequest() (request *CreatePGInstanceHourRequest) {
	request = &CreatePGInstanceHourRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "CreatePGInstanceHour")
	return
}

func NewCreatePGInstanceHourResponse() (response *CreatePGInstanceHourResponse) {
	response = &CreatePGInstanceHourResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (CreatePGInstanceHour) 用于创建Postgres实例。
func (c *Client) CreatePGInstanceHour(request *CreatePGInstanceHourRequest) (response *CreatePGInstanceHourResponse, err error) {
	if request == nil {
		request = NewCreatePGInstanceHourRequest()
	}
	response = NewCreatePGInstanceHourResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceOverviewRequest() (request *DescribeResourceOverviewRequest) {
	request = &DescribeResourceOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeResourceOverview")
	return
}

func NewDescribeResourceOverviewResponse() (response *DescribeResourceOverviewResponse) {
	response = &DescribeResourceOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取本地区资源使用情况概览
func (c *Client) DescribeResourceOverview(request *DescribeResourceOverviewRequest) (response *DescribeResourceOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeResourceOverviewRequest()
	}
	response = NewDescribeResourceOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceDetailsRequest() (request *DescribeInstanceDetailsRequest) {
	request = &DescribeInstanceDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeInstanceDetails")
	return
}

func NewDescribeInstanceDetailsResponse() (response *DescribeInstanceDetailsResponse) {
	response = &DescribeInstanceDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstanceDetails) 用于查询单个实例的详情。
func (c *Client) DescribeInstanceDetails(request *DescribeInstanceDetailsRequest) (response *DescribeInstanceDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceDetailsRequest()
	}
	response = NewDescribeInstanceDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBParametersRequest() (request *DescribeDBParametersRequest) {
	request = &DescribeDBParametersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeDBParameters")
	return
}

func NewDescribeDBParametersResponse() (response *DescribeDBParametersResponse) {
	response = &DescribeDBParametersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeDBParameters) 用于查询一个实例的参数信息。
func (c *Client) DescribeDBParameters(request *DescribeDBParametersRequest) (response *DescribeDBParametersResponse, err error) {
	if request == nil {
		request = NewDescribeDBParametersRequest()
	}
	response = NewDescribeDBParametersResponse()
	err = c.Send(request, response)
	return
}

func NewIsolateInstanceHourRequest() (request *IsolateInstanceHourRequest) {
	request = &IsolateInstanceHourRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "IsolateInstanceHour")
	return
}

func NewIsolateInstanceHourResponse() (response *IsolateInstanceHourResponse) {
	response = &IsolateInstanceHourResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (IsolateInstanceHour) 用于隔离实例。
func (c *Client) IsolateInstanceHour(request *IsolateInstanceHourRequest) (response *IsolateInstanceHourResponse, err error) {
	if request == nil {
		request = NewIsolateInstanceHourRequest()
	}
	response = NewIsolateInstanceHourResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceSyncTypeRequest() (request *ModifyInstanceSyncTypeRequest) {
	request = &ModifyInstanceSyncTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ModifyInstanceSyncType")
	return
}

func NewModifyInstanceSyncTypeResponse() (response *ModifyInstanceSyncTypeResponse) {
	response = &ModifyInstanceSyncTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ModifyInstanceSyncType) 修改实例节点的数据复制类型。
//
// * 可以修改实例节点的数据复制类型（当前仅支持修改数据节点）。
func (c *Client) ModifyInstanceSyncType(request *ModifyInstanceSyncTypeRequest) (response *ModifyInstanceSyncTypeResponse, err error) {
	if request == nil {
		request = NewModifyInstanceSyncTypeRequest()
	}
	response = NewModifyInstanceSyncTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePGAvailableZoneConfigRequest() (request *DescribePGAvailableZoneConfigRequest) {
	request = &DescribePGAvailableZoneConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribePGAvailableZoneConfig")
	return
}

func NewDescribePGAvailableZoneConfigResponse() (response *DescribePGAvailableZoneConfigResponse) {
	response = &DescribePGAvailableZoneConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribePGAvailableZoneConfig)用于查询云数据库Postgres的可售卖地域和可用区信息。
func (c *Client) DescribePGAvailableZoneConfig(request *DescribePGAvailableZoneConfigRequest) (response *DescribePGAvailableZoneConfigResponse, err error) {
	if request == nil {
		request = NewDescribePGAvailableZoneConfigRequest()
	}
	response = NewDescribePGAvailableZoneConfigResponse()
	err = c.Send(request, response)
	return
}

func NewRenewInstancesRequest() (request *RenewInstancesRequest) {
	request = &RenewInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "RenewInstances")
	return
}

func NewRenewInstancesResponse() (response *RenewInstancesResponse) {
	response = &RenewInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (RenewInstances) 用于已隔离实例重新上线，目前仅支持单个实例重新上线。
func (c *Client) RenewInstances(request *RenewInstancesRequest) (response *RenewInstancesResponse, err error) {
	if request == nil {
		request = NewRenewInstancesRequest()
	}
	response = NewRenewInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcePoolsRequest() (request *DescribeResourcePoolsRequest) {
	request = &DescribeResourcePoolsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeResourcePools")
	return
}

func NewDescribeResourcePoolsResponse() (response *DescribeResourcePoolsResponse) {
	response = &DescribeResourcePoolsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于在postgres购买页，查询可部署实例的资源池，能看到的资源池个数与该用户的资源池查看权限有关，权限可在运营端修改
func (c *Client) DescribeResourcePools(request *DescribeResourcePoolsRequest) (response *DescribeResourcePoolsResponse, err error) {
	if request == nil {
		request = NewDescribeResourcePoolsRequest()
	}
	response = NewDescribeResourcePoolsResponse()
	err = c.Send(request, response)
	return
}

func NewActiveIsolatedPGInstanceRequest() (request *ActiveIsolatedPGInstanceRequest) {
	request = &ActiveIsolatedPGInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ActiveIsolatedPGInstance")
	return
}

func NewActiveIsolatedPGInstanceResponse() (response *ActiveIsolatedPGInstanceResponse) {
	response = &ActiveIsolatedPGInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ActiveIsolatedPGInstance) 用于恢复Postgres实例。
func (c *Client) ActiveIsolatedPGInstance(request *ActiveIsolatedPGInstanceRequest) (response *ActiveIsolatedPGInstanceResponse, err error) {
	if request == nil {
		request = NewActiveIsolatedPGInstanceRequest()
	}
	response = NewActiveIsolatedPGInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeInstances) 用于查询一个或多个实例的详细信息。
//
// * 可以根据实例ID、实例名称或者实例计费模式等信息来查询实例的详细信息。过滤信息详细请见过滤器Filter。
// * 如果参数为空，返回当前用户一定数量（Limit所指定的数量，默认为20）的实例。
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableRegionRequest() (request *DescribeAvailableRegionRequest) {
	request = &DescribeAvailableRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableRegion")
	return
}

func NewDescribeAvailableRegionResponse() (response *DescribeAvailableRegionResponse) {
	response = &DescribeAvailableRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeAvailableRegion)用于查询售卖地域信息
func (c *Client) DescribeAvailableRegion(request *DescribeAvailableRegionRequest) (response *DescribeAvailableRegionResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableRegionRequest()
	}
	response = NewDescribeAvailableRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePGInstanceDetailsRequest() (request *DescribePGInstanceDetailsRequest) {
	request = &DescribePGInstanceDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribePGInstanceDetails")
	return
}

func NewDescribePGInstanceDetailsResponse() (response *DescribePGInstanceDetailsResponse) {
	response = &DescribePGInstanceDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribePGInstanceDetails) 用于查询单个Postgres实例的详情。
func (c *Client) DescribePGInstanceDetails(request *DescribePGInstanceDetailsRequest) (response *DescribePGInstanceDetailsResponse, err error) {
	if request == nil {
		request = NewDescribePGInstanceDetailsRequest()
	}
	response = NewDescribePGInstanceDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDBInstanceSecurityGroupsRequest() (request *ModifyDBInstanceSecurityGroupsRequest) {
	request = &ModifyDBInstanceSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ModifyDBInstanceSecurityGroups")
	return
}

func NewModifyDBInstanceSecurityGroupsResponse() (response *ModifyDBInstanceSecurityGroupsResponse) {
	response = &ModifyDBInstanceSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyDBInstanceSecurityGroups）用于修改云数据库安全组
func (c *Client) ModifyDBInstanceSecurityGroups(request *ModifyDBInstanceSecurityGroupsRequest) (response *ModifyDBInstanceSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewModifyDBInstanceSecurityGroupsRequest()
	}
	response = NewModifyDBInstanceSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseIsolatedInstanceRequest() (request *ReleaseIsolatedInstanceRequest) {
	request = &ReleaseIsolatedInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ReleaseIsolatedInstance")
	return
}

func NewReleaseIsolatedInstanceResponse() (response *ReleaseIsolatedInstanceResponse) {
	response = &ReleaseIsolatedInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ReleaseIsolatedInstance) 用于销毁单个已隔离实例。
func (c *Client) ReleaseIsolatedInstance(request *ReleaseIsolatedInstanceRequest) (response *ReleaseIsolatedInstanceResponse, err error) {
	if request == nil {
		request = NewReleaseIsolatedInstanceRequest()
	}
	response = NewReleaseIsolatedInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableZonesRequest() (request *DescribeAvailableZonesRequest) {
	request = &DescribeAvailableZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableZones")
	return
}

func NewDescribeAvailableZonesResponse() (response *DescribeAvailableZonesResponse) {
	response = &DescribeAvailableZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取指定资源池下的可用区信息
func (c *Client) DescribeAvailableZones(request *DescribeAvailableZonesRequest) (response *DescribeAvailableZonesResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableZonesRequest()
	}
	response = NewDescribeAvailableZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBackupDetailsRequest() (request *DescribeBackupDetailsRequest) {
	request = &DescribeBackupDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeBackupDetails")
	return
}

func NewDescribeBackupDetailsResponse() (response *DescribeBackupDetailsResponse) {
	response = &DescribeBackupDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeBackupDetails）用于查询备份详情。
func (c *Client) DescribeBackupDetails(request *DescribeBackupDetailsRequest) (response *DescribeBackupDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeBackupDetailsRequest()
	}
	response = NewDescribeBackupDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
	request = &ResetAccountPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ResetAccountPassword")
	return
}

func NewResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
	response = &ResetAccountPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ResetAccountPassword）用于重置云数据库账号的密码。
func (c *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
	if request == nil {
		request = NewResetAccountPasswordRequest()
	}
	response = NewResetAccountPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBackupRequest() (request *CreateBackupRequest) {
	request = &CreateBackupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "CreateBackup")
	return
}

func NewCreateBackupResponse() (response *CreateBackupResponse) {
	response = &CreateBackupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TbaseV2版本的手动备份
func (c *Client) CreateBackup(request *CreateBackupRequest) (response *CreateBackupResponse, err error) {
	if request == nil {
		request = NewCreateBackupRequest()
	}
	response = NewCreateBackupResponse()
	err = c.Send(request, response)
	return
}

func NewRestartInstanceRequest() (request *RestartInstanceRequest) {
	request = &RestartInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "RestartInstance")
	return
}

func NewRestartInstanceResponse() (response *RestartInstanceResponse) {
	response = &RestartInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (RestartInstance) 用于重启实例。
//
// * 可以根据实例ID、节点类型对不同节点或者整个实例进行重启操作。
func (c *Client) RestartInstance(request *RestartInstanceRequest) (response *RestartInstanceResponse, err error) {
	if request == nil {
		request = NewRestartInstanceRequest()
	}
	response = NewRestartInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePGInstancesRequest() (request *DescribePGInstancesRequest) {
	request = &DescribePGInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribePGInstances")
	return
}

func NewDescribePGInstancesResponse() (response *DescribePGInstancesResponse) {
	response = &DescribePGInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribePGInstances) 用于查询一个或多个Postgres实例的详细信息。
func (c *Client) DescribePGInstances(request *DescribePGInstancesRequest) (response *DescribePGInstancesResponse, err error) {
	if request == nil {
		request = NewDescribePGInstancesRequest()
	}
	response = NewDescribePGInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAvailableCharsetRequest() (request *DescribeAvailableCharsetRequest) {
	request = &DescribeAvailableCharsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeAvailableCharset")
	return
}

func NewDescribeAvailableCharsetResponse() (response *DescribeAvailableCharsetResponse) {
	response = &DescribeAvailableCharsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (DescribeAvailableCharset) 用于查询创建实例时可选择的字符集。
func (c *Client) DescribeAvailableCharset(request *DescribeAvailableCharsetRequest) (response *DescribeAvailableCharsetResponse, err error) {
	if request == nil {
		request = NewDescribeAvailableCharsetRequest()
	}
	response = NewDescribeAvailableCharsetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDBSecurityGroupsRequest() (request *DescribeDBSecurityGroupsRequest) {
	request = &DescribeDBSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DescribeDBSecurityGroups")
	return
}

func NewDescribeDBSecurityGroupsResponse() (response *DescribeDBSecurityGroupsResponse) {
	response = &DescribeDBSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(DescribeDBSecurityGroups)用于查询实例的安全组详情。
func (c *Client) DescribeDBSecurityGroups(request *DescribeDBSecurityGroupsRequest) (response *DescribeDBSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDBSecurityGroupsRequest()
	}
	response = NewDescribeDBSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInstanceHourRequest() (request *CreateInstanceHourRequest) {
	request = &CreateInstanceHourRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "CreateInstanceHour")
	return
}

func NewCreateInstanceHourResponse() (response *CreateInstanceHourResponse) {
	response = &CreateInstanceHourResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (CreateInstanceHour) 用于创建实例。
func (c *Client) CreateInstanceHour(request *CreateInstanceHourRequest) (response *CreateInstanceHourResponse, err error) {
	if request == nil {
		request = NewCreateInstanceHourRequest()
	}
	response = NewCreateInstanceHourResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePGReadOnlyVipRequest() (request *DeletePGReadOnlyVipRequest) {
	request = &DeletePGReadOnlyVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "DeletePGReadOnlyVip")
	return
}

func NewDeletePGReadOnlyVipResponse() (response *DeletePGReadOnlyVipResponse) {
	response = &DeletePGReadOnlyVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于删除指定postgres实例的只读VIP。
func (c *Client) DeletePGReadOnlyVip(request *DeletePGReadOnlyVipRequest) (response *DeletePGReadOnlyVipResponse, err error) {
	if request == nil {
		request = NewDeletePGReadOnlyVipRequest()
	}
	response = NewDeletePGReadOnlyVipResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceNameRequest() (request *ModifyInstanceNameRequest) {
	request = &ModifyInstanceNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ModifyInstanceName")
	return
}

func NewModifyInstanceNameResponse() (response *ModifyInstanceNameResponse) {
	response = &ModifyInstanceNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（ModifyInstanceName）用于修改云数据的名称。
func (c *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (response *ModifyInstanceNameResponse, err error) {
	if request == nil {
		request = NewModifyInstanceNameRequest()
	}
	response = NewModifyInstanceNameResponse()
	err = c.Send(request, response)
	return
}

func NewActiveIsolatedInstanceRequest() (request *ActiveIsolatedInstanceRequest) {
	request = &ActiveIsolatedInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tbase", APIVersion, "ActiveIsolatedInstance")
	return
}

func NewActiveIsolatedInstanceResponse() (response *ActiveIsolatedInstanceResponse) {
	response = &ActiveIsolatedInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口 (ActiveIsolatedInstance) 用于恢复实例。
func (c *Client) ActiveIsolatedInstance(request *ActiveIsolatedInstanceRequest) (response *ActiveIsolatedInstanceResponse, err error) {
	if request == nil {
		request = NewActiveIsolatedInstanceRequest()
	}
	response = NewActiveIsolatedInstanceResponse()
	err = c.Send(request, response)
	return
}

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

package v20180326

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-03-26"

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

func NewDescribeMsApiListRequest() (request *DescribeMsApiListRequest) {
	request = &DescribeMsApiListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsApiList")
	return
}

func NewDescribeMsApiListResponse() (response *DescribeMsApiListResponse) {
	response = &DescribeMsApiListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务API列表
func (c *Client) DescribeMsApiList(request *DescribeMsApiListRequest) (response *DescribeMsApiListResponse, err error) {
	if request == nil {
		request = NewDescribeMsApiListRequest()
	}
	response = NewDescribeMsApiListResponse()
	err = c.Send(request, response)
	return
}

func NewGetConfigpropsRequest() (request *GetConfigpropsRequest) {
	request = &GetConfigpropsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetConfigprops")
	return
}

func NewGetConfigpropsResponse() (response *GetConfigpropsResponse) {
	response = &GetConfigpropsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetConfigprops
func (c *Client) GetConfigprops(request *GetConfigpropsRequest) (response *GetConfigpropsResponse, err error) {
	if request == nil {
		request = NewGetConfigpropsRequest()
	}
	response = NewGetConfigpropsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineMonitorRequest() (request *DescribeMachineMonitorRequest) {
	request = &DescribeMachineMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMachineMonitor")
	return
}

func NewDescribeMachineMonitorResponse() (response *DescribeMachineMonitorResponse) {
	response = &DescribeMachineMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//
// 机器列表监控
func (c *Client) DescribeMachineMonitor(request *DescribeMachineMonitorRequest) (response *DescribeMachineMonitorResponse, err error) {
	if request == nil {
		request = NewDescribeMachineMonitorRequest()
	}
	response = NewDescribeMachineMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFileConfigRequest() (request *CreateFileConfigRequest) {
	request = &CreateFileConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateFileConfig")
	return
}

func NewCreateFileConfigResponse() (response *CreateFileConfigResponse) {
	response = &CreateFileConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建文件配置项
func (c *Client) CreateFileConfig(request *CreateFileConfigRequest) (response *CreateFileConfigResponse, err error) {
	if request == nil {
		request = NewCreateFileConfigRequest()
	}
	response = NewCreateFileConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRepositoryRequest() (request *DeleteRepositoryRequest) {
	request = &DeleteRepositoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteRepository")
	return
}

func NewDeleteRepositoryResponse() (response *DeleteRepositoryResponse) {
	response = &DeleteRepositoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除仓库
func (c *Client) DeleteRepository(request *DeleteRepositoryRequest) (response *DeleteRepositoryResponse, err error) {
	if request == nil {
		request = NewDeleteRepositoryRequest()
	}
	response = NewDeleteRepositoryResponse()
	err = c.Send(request, response)
	return
}

func NewValidateLogSchemaRequest() (request *ValidateLogSchemaRequest) {
	request = &ValidateLogSchemaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ValidateLogSchema")
	return
}

func NewValidateLogSchemaResponse() (response *ValidateLogSchemaResponse) {
	response = &ValidateLogSchemaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验日志Pattern格式合法性
func (c *Client) ValidateLogSchema(request *ValidateLogSchemaRequest) (response *ValidateLogSchemaResponse, err error) {
	if request == nil {
		request = NewValidateLogSchemaRequest()
	}
	response = NewValidateLogSchemaResponse()
	err = c.Send(request, response)
	return
}

func NewDisableTaskFlowRequest() (request *DisableTaskFlowRequest) {
	request = &DisableTaskFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableTaskFlow")
	return
}

func NewDisableTaskFlowResponse() (response *DisableTaskFlowResponse) {
	response = &DisableTaskFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停用工作流
func (c *Client) DisableTaskFlow(request *DisableTaskFlowRequest) (response *DisableTaskFlowResponse, err error) {
	if request == nil {
		request = NewDisableTaskFlowRequest()
	}
	response = NewDisableTaskFlowResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
	request = &CreateRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateRoute")
	return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
	response = &CreateRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建路由
func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
	if request == nil {
		request = NewCreateRouteRequest()
	}
	response = NewCreateRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerMicroServiceListRequest() (request *DescribeContainerMicroServiceListRequest) {
	request = &DescribeContainerMicroServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerMicroServiceList")
	return
}

func NewDescribeContainerMicroServiceListResponse() (response *DescribeContainerMicroServiceListResponse) {
	response = &DescribeContainerMicroServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器微服务详情
func (c *Client) DescribeContainerMicroServiceList(request *DescribeContainerMicroServiceListRequest) (response *DescribeContainerMicroServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeContainerMicroServiceListRequest()
	}
	response = NewDescribeContainerMicroServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewIssueLicenseRequest() (request *IssueLicenseRequest) {
	request = &IssueLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "IssueLicense")
	return
}

func NewIssueLicenseResponse() (response *IssueLicenseResponse) {
	response = &IssueLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 授予许可证
func (c *Client) IssueLicense(request *IssueLicenseRequest) (response *IssueLicenseResponse, err error) {
	if request == nil {
		request = NewIssueLicenseRequest()
	}
	response = NewIssueLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNamespaceRequest() (request *CreateNamespaceRequest) {
	request = &CreateNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateNamespace")
	return
}

func NewCreateNamespaceResponse() (response *CreateNamespaceResponse) {
	response = &CreateNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建命名空间
func (c *Client) CreateNamespace(request *CreateNamespaceRequest) (response *CreateNamespaceResponse, err error) {
	if request == nil {
		request = NewCreateNamespaceRequest()
	}
	response = NewCreateNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReleaseRequest() (request *ModifyReleaseRequest) {
	request = &ModifyReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyRelease")
	return
}

func NewModifyReleaseResponse() (response *ModifyReleaseResponse) {
	response = &ModifyReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改发布单
func (c *Client) ModifyRelease(request *ModifyReleaseRequest) (response *ModifyReleaseResponse, err error) {
	if request == nil {
		request = NewModifyReleaseRequest()
	}
	response = NewModifyReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLicenseApplicationRequest() (request *CreateLicenseApplicationRequest) {
	request = &CreateLicenseApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateLicenseApplication")
	return
}

func NewCreateLicenseApplicationResponse() (response *CreateLicenseApplicationResponse) {
	response = &CreateLicenseApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) CreateLicenseApplication(request *CreateLicenseApplicationRequest) (response *CreateLicenseApplicationResponse, err error) {
	if request == nil {
		request = NewCreateLicenseApplicationRequest()
	}
	response = NewCreateLicenseApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasesRequest() (request *DescribeReleasesRequest) {
	request = &DescribeReleasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleases")
	return
}

func NewDescribeReleasesResponse() (response *DescribeReleasesResponse) {
	response = &DescribeReleasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布单列表
func (c *Client) DescribeReleases(request *DescribeReleasesRequest) (response *DescribeReleasesResponse, err error) {
	if request == nil {
		request = NewDescribeReleasesRequest()
	}
	response = NewDescribeReleasesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayApisRequest() (request *DescribeGatewayApisRequest) {
	request = &DescribeGatewayApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayApis")
	return
}

func NewDescribeGatewayApisResponse() (response *DescribeGatewayApisResponse) {
	response = &DescribeGatewayApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API分组下的Api列表信息
func (c *Client) DescribeGatewayApis(request *DescribeGatewayApisRequest) (response *DescribeGatewayApisResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayApisRequest()
	}
	response = NewDescribeGatewayApisResponse()
	err = c.Send(request, response)
	return
}

func NewReassociateBusinessLogConfigRequest() (request *ReassociateBusinessLogConfigRequest) {
	request = &ReassociateBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReassociateBusinessLogConfig")
	return
}

func NewReassociateBusinessLogConfigResponse() (response *ReassociateBusinessLogConfigResponse) {
	response = &ReassociateBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重关联业务日志配置
func (c *Client) ReassociateBusinessLogConfig(request *ReassociateBusinessLogConfigRequest) (response *ReassociateBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewReassociateBusinessLogConfigRequest()
	}
	response = NewReassociateBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLanesRequest() (request *DescribeLanesRequest) {
	request = &DescribeLanesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLanes")
	return
}

func NewDescribeLanesResponse() (response *DescribeLanesResponse) {
	response = &DescribeLanesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询泳道列表
func (c *Client) DescribeLanes(request *DescribeLanesRequest) (response *DescribeLanesResponse, err error) {
	if request == nil {
		request = NewDescribeLanesRequest()
	}
	response = NewDescribeLanesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDeliveryConfigRequest() (request *CreateDeliveryConfigRequest) {
	request = &CreateDeliveryConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateDeliveryConfig")
	return
}

func NewCreateDeliveryConfigResponse() (response *CreateDeliveryConfigResponse) {
	response = &CreateDeliveryConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建一个配置项，用于用户投递数据
func (c *Client) CreateDeliveryConfig(request *CreateDeliveryConfigRequest) (response *CreateDeliveryConfigResponse, err error) {
	if request == nil {
		request = NewCreateDeliveryConfigRequest()
	}
	response = NewCreateDeliveryConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSetReleasePipelineRunTaskRollbackRequest() (request *SetReleasePipelineRunTaskRollbackRequest) {
	request = &SetReleasePipelineRunTaskRollbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SetReleasePipelineRunTaskRollback")
	return
}

func NewSetReleasePipelineRunTaskRollbackResponse() (response *SetReleasePipelineRunTaskRollbackResponse) {
	response = &SetReleasePipelineRunTaskRollbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置发布单执行任务是否允许回滚
func (c *Client) SetReleasePipelineRunTaskRollback(request *SetReleasePipelineRunTaskRollbackRequest) (response *SetReleasePipelineRunTaskRollbackResponse, err error) {
	if request == nil {
		request = NewSetReleasePipelineRunTaskRollbackRequest()
	}
	response = NewSetReleasePipelineRunTaskRollbackResponse()
	err = c.Send(request, response)
	return
}

func NewTriggerReleasePipelineRunActionRequest() (request *TriggerReleasePipelineRunActionRequest) {
	request = &TriggerReleasePipelineRunActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "TriggerReleasePipelineRunAction")
	return
}

func NewTriggerReleasePipelineRunActionResponse() (response *TriggerReleasePipelineRunActionResponse) {
	response = &TriggerReleasePipelineRunActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发发布单流水操作
func (c *Client) TriggerReleasePipelineRunAction(request *TriggerReleasePipelineRunActionRequest) (response *TriggerReleasePipelineRunActionResponse, err error) {
	if request == nil {
		request = NewTriggerReleasePipelineRunActionRequest()
	}
	response = NewTriggerReleasePipelineRunActionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGCLogRequest() (request *DescribeGCLogRequest) {
	request = &DescribeGCLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGCLog")
	return
}

func NewDescribeGCLogResponse() (response *DescribeGCLogResponse) {
	response = &DescribeGCLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询分析指定目录下GC日志
func (c *Client) DescribeGCLog(request *DescribeGCLogRequest) (response *DescribeGCLogResponse, err error) {
	if request == nil {
		request = NewDescribeGCLogRequest()
	}
	response = NewDescribeGCLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApiRateLimitRuleRequest() (request *CreateApiRateLimitRuleRequest) {
	request = &CreateApiRateLimitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateApiRateLimitRule")
	return
}

func NewCreateApiRateLimitRuleResponse() (response *CreateApiRateLimitRuleResponse) {
	response = &CreateApiRateLimitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建API限流规则
func (c *Client) CreateApiRateLimitRule(request *CreateApiRateLimitRuleRequest) (response *CreateApiRateLimitRuleResponse, err error) {
	if request == nil {
		request = NewCreateApiRateLimitRuleRequest()
	}
	response = NewCreateApiRateLimitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskBatchRecordRequest() (request *DescribeTaskBatchRecordRequest) {
	request = &DescribeTaskBatchRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskBatchRecord")
	return
}

func NewDescribeTaskBatchRecordResponse() (response *DescribeTaskBatchRecordResponse) {
	response = &DescribeTaskBatchRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//  查看任务批次的详情
func (c *Client) DescribeTaskBatchRecord(request *DescribeTaskBatchRecordRequest) (response *DescribeTaskBatchRecordResponse, err error) {
	if request == nil {
		request = NewDescribeTaskBatchRecordRequest()
	}
	response = NewDescribeTaskBatchRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayWithPluginRequest() (request *DescribeGatewayWithPluginRequest) {
	request = &DescribeGatewayWithPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayWithPlugin")
	return
}

func NewDescribeGatewayWithPluginResponse() (response *DescribeGatewayWithPluginResponse) {
	response = &DescribeGatewayWithPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询插件绑定网关
func (c *Client) DescribeGatewayWithPlugin(request *DescribeGatewayWithPluginRequest) (response *DescribeGatewayWithPluginResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayWithPluginRequest()
	}
	response = NewDescribeGatewayWithPluginResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopAvgTimeCostServicesRequest() (request *GetTopAvgTimeCostServicesRequest) {
	request = &GetTopAvgTimeCostServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTopAvgTimeCostServices")
	return
}

func NewGetTopAvgTimeCostServicesResponse() (response *GetTopAvgTimeCostServicesResponse) {
	response = &GetTopAvgTimeCostServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TopN请求平均耗时服务列表
func (c *Client) GetTopAvgTimeCostServices(request *GetTopAvgTimeCostServicesRequest) (response *GetTopAvgTimeCostServicesResponse, err error) {
	if request == nil {
		request = NewGetTopAvgTimeCostServicesRequest()
	}
	response = NewGetTopAvgTimeCostServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductHelpRequest() (request *DeleteProductHelpRequest) {
	request = &DeleteProductHelpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteProductHelp")
	return
}

func NewDeleteProductHelpResponse() (response *DeleteProductHelpResponse) {
	response = &DeleteProductHelpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除产品帮助
func (c *Client) DeleteProductHelp(request *DeleteProductHelpRequest) (response *DeleteProductHelpResponse, err error) {
	if request == nil {
		request = NewDeleteProductHelpRequest()
	}
	response = NewDeleteProductHelpResponse()
	err = c.Send(request, response)
	return
}

func NewGetTraceRequest() (request *GetTraceRequest) {
	request = &GetTraceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTrace")
	return
}

func NewGetTraceResponse() (response *GetTraceResponse) {
	response = &GetTraceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetTrace
func (c *Client) GetTrace(request *GetTraceRequest) (response *GetTraceResponse, err error) {
	if request == nil {
		request = NewGetTraceRequest()
	}
	response = NewGetTraceResponse()
	err = c.Send(request, response)
	return
}

func NewBillingOperationRenewRequest() (request *BillingOperationRenewRequest) {
	request = &BillingOperationRenewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "BillingOperationRenew")
	return
}

func NewBillingOperationRenewResponse() (response *BillingOperationRenewResponse) {
	response = &BillingOperationRenewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费运营端续费
func (c *Client) BillingOperationRenew(request *BillingOperationRenewRequest) (response *BillingOperationRenewResponse, err error) {
	if request == nil {
		request = NewBillingOperationRenewRequest()
	}
	response = NewBillingOperationRenewResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayWeChatMiniProgramLoginPluginRequest() (request *CreateGatewayWeChatMiniProgramLoginPluginRequest) {
	request = &CreateGatewayWeChatMiniProgramLoginPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayWeChatMiniProgramLoginPlugin")
	return
}

func NewCreateGatewayWeChatMiniProgramLoginPluginResponse() (response *CreateGatewayWeChatMiniProgramLoginPluginResponse) {
	response = &CreateGatewayWeChatMiniProgramLoginPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增网关微信小程序登录插件
func (c *Client) CreateGatewayWeChatMiniProgramLoginPlugin(request *CreateGatewayWeChatMiniProgramLoginPluginRequest) (response *CreateGatewayWeChatMiniProgramLoginPluginResponse, err error) {
	if request == nil {
		request = NewCreateGatewayWeChatMiniProgramLoginPluginRequest()
	}
	response = NewCreateGatewayWeChatMiniProgramLoginPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroServiceSummaryListRequest() (request *DescribeMicroServiceSummaryListRequest) {
	request = &DescribeMicroServiceSummaryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroServiceSummaryList")
	return
}

func NewDescribeMicroServiceSummaryListResponse() (response *DescribeMicroServiceSummaryListResponse) {
	response = &DescribeMicroServiceSummaryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务治理概要列表
func (c *Client) DescribeMicroServiceSummaryList(request *DescribeMicroServiceSummaryListRequest) (response *DescribeMicroServiceSummaryListResponse, err error) {
	if request == nil {
		request = NewDescribeMicroServiceSummaryListRequest()
	}
	response = NewDescribeMicroServiceSummaryListResponse()
	err = c.Send(request, response)
	return
}

func NewEnableRouteRequest() (request *EnableRouteRequest) {
	request = &EnableRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableRoute")
	return
}

func NewEnableRouteResponse() (response *EnableRouteResponse) {
	response = &EnableRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用路由
func (c *Client) EnableRoute(request *EnableRouteRequest) (response *EnableRouteResponse, err error) {
	if request == nil {
		request = NewEnableRouteRequest()
	}
	response = NewEnableRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceUsageConfigRequest() (request *DescribeResourceUsageConfigRequest) {
	request = &DescribeResourceUsageConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceUsageConfig")
	return
}

func NewDescribeResourceUsageConfigResponse() (response *DescribeResourceUsageConfigResponse) {
	response = &DescribeResourceUsageConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取概览页配置
func (c *Client) DescribeResourceUsageConfig(request *DescribeResourceUsageConfigRequest) (response *DescribeResourceUsageConfigResponse, err error) {
	if request == nil {
		request = NewDescribeResourceUsageConfigRequest()
	}
	response = NewDescribeResourceUsageConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupInstancesRequest() (request *DescribeGroupInstancesRequest) {
	request = &DescribeGroupInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupInstances")
	return
}

func NewDescribeGroupInstancesResponse() (response *DescribeGroupInstancesResponse) {
	response = &DescribeGroupInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询虚拟机部署组云主机列表
func (c *Client) DescribeGroupInstances(request *DescribeGroupInstancesRequest) (response *DescribeGroupInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeGroupInstancesRequest()
	}
	response = NewDescribeGroupInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseConfigRequest() (request *ReleaseConfigRequest) {
	request = &ReleaseConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleaseConfig")
	return
}

func NewReleaseConfigResponse() (response *ReleaseConfigResponse) {
	response = &ReleaseConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布配置
func (c *Client) ReleaseConfig(request *ReleaseConfigRequest) (response *ReleaseConfigResponse, err error) {
	if request == nil {
		request = NewReleaseConfigRequest()
	}
	response = NewReleaseConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSearchOssStaticBusinessLogRequest() (request *SearchOssStaticBusinessLogRequest) {
	request = &SearchOssStaticBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchOssStaticBusinessLog")
	return
}

func NewSearchOssStaticBusinessLogResponse() (response *SearchOssStaticBusinessLogResponse) {
	response = &SearchOssStaticBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 搜索日志
func (c *Client) SearchOssStaticBusinessLog(request *SearchOssStaticBusinessLogRequest) (response *SearchOssStaticBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchOssStaticBusinessLogRequest()
	}
	response = NewSearchOssStaticBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewValidateDeleteConfigRequest() (request *ValidateDeleteConfigRequest) {
	request = &ValidateDeleteConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ValidateDeleteConfig")
	return
}

func NewValidateDeleteConfigResponse() (response *ValidateDeleteConfigResponse) {
	response = &ValidateDeleteConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验配置项可删除
func (c *Client) ValidateDeleteConfig(request *ValidateDeleteConfigRequest) (response *ValidateDeleteConfigResponse, err error) {
	if request == nil {
		request = NewValidateDeleteConfigRequest()
	}
	response = NewValidateDeleteConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthorizationsRequest() (request *DescribeAuthorizationsRequest) {
	request = &DescribeAuthorizationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorizations")
	return
}

func NewDescribeAuthorizationsResponse() (response *DescribeAuthorizationsResponse) {
	response = &DescribeAuthorizationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务权限配置列表
func (c *Client) DescribeAuthorizations(request *DescribeAuthorizationsRequest) (response *DescribeAuthorizationsResponse, err error) {
	if request == nil {
		request = NewDescribeAuthorizationsRequest()
	}
	response = NewDescribeAuthorizationsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupRequest() (request *DeleteGroupRequest) {
	request = &DeleteGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteGroup")
	return
}

func NewDeleteGroupResponse() (response *DeleteGroupResponse) {
	response = &DeleteGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除容器部署组
func (c *Client) DeleteGroup(request *DeleteGroupRequest) (response *DeleteGroupResponse, err error) {
	if request == nil {
		request = NewDeleteGroupRequest()
	}
	response = NewDeleteGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvocationMetricDataDimensionRequest() (request *DescribeInvocationMetricDataDimensionRequest) {
	request = &DescribeInvocationMetricDataDimensionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationMetricDataDimension")
	return
}

func NewDescribeInvocationMetricDataDimensionResponse() (response *DescribeInvocationMetricDataDimensionResponse) {
	response = &DescribeInvocationMetricDataDimensionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询维度
func (c *Client) DescribeInvocationMetricDataDimension(request *DescribeInvocationMetricDataDimensionRequest) (response *DescribeInvocationMetricDataDimensionResponse, err error) {
	if request == nil {
		request = NewDescribeInvocationMetricDataDimensionRequest()
	}
	response = NewDescribeInvocationMetricDataDimensionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGroupLaneRequest() (request *ModifyGroupLaneRequest) {
	request = &ModifyGroupLaneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyGroupLane")
	return
}

func NewModifyGroupLaneResponse() (response *ModifyGroupLaneResponse) {
	response = &ModifyGroupLaneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新部署组泳道信息
func (c *Client) ModifyGroupLane(request *ModifyGroupLaneRequest) (response *ModifyGroupLaneResponse, err error) {
	if request == nil {
		request = NewModifyGroupLaneRequest()
	}
	response = NewModifyGroupLaneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterSchedulabilityRequest() (request *DescribeClusterSchedulabilityRequest) {
	request = &DescribeClusterSchedulabilityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterSchedulability")
	return
}

func NewDescribeClusterSchedulabilityResponse() (response *DescribeClusterSchedulabilityResponse) {
	response = &DescribeClusterSchedulabilityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建容器部署组时，根据用户所需部署组资源确定集群剩余资源是否满足调度需求
func (c *Client) DescribeClusterSchedulability(request *DescribeClusterSchedulabilityRequest) (response *DescribeClusterSchedulabilityResponse, err error) {
	if request == nil {
		request = NewDescribeClusterSchedulabilityRequest()
	}
	response = NewDescribeClusterSchedulabilityResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePathRewritesRequest() (request *CreatePathRewritesRequest) {
	request = &CreatePathRewritesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreatePathRewrites")
	return
}

func NewCreatePathRewritesResponse() (response *CreatePathRewritesResponse) {
	response = &CreatePathRewritesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建路径重写
func (c *Client) CreatePathRewrites(request *CreatePathRewritesRequest) (response *CreatePathRewritesResponse, err error) {
	if request == nil {
		request = NewCreatePathRewritesRequest()
	}
	response = NewCreatePathRewritesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableTaskFlowRequest() (request *EnableTaskFlowRequest) {
	request = &EnableTaskFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableTaskFlow")
	return
}

func NewEnableTaskFlowResponse() (response *EnableTaskFlowResponse) {
	response = &EnableTaskFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用工作流
func (c *Client) EnableTaskFlow(request *EnableTaskFlowRequest) (response *EnableTaskFlowResponse, err error) {
	if request == nil {
		request = NewEnableTaskFlowRequest()
	}
	response = NewEnableTaskFlowResponse()
	err = c.Send(request, response)
	return
}

func NewListAppRequest() (request *ListAppRequest) {
	request = &ListAppRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListApp")
	return
}

func NewListAppResponse() (response *ListAppResponse) {
	response = &ListAppResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的应用信息
func (c *Client) ListApp(request *ListAppRequest) (response *ListAppResponse, err error) {
	if request == nil {
		request = NewListAppRequest()
	}
	response = NewListAppResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceBatchOperationRequest() (request *DescribeResourceBatchOperationRequest) {
	request = &DescribeResourceBatchOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceBatchOperation")
	return
}

func NewDescribeResourceBatchOperationResponse() (response *DescribeResourceBatchOperationResponse) {
	response = &DescribeResourceBatchOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资源批次操作描述接口
func (c *Client) DescribeResourceBatchOperation(request *DescribeResourceBatchOperationRequest) (response *DescribeResourceBatchOperationResponse, err error) {
	if request == nil {
		request = NewDescribeResourceBatchOperationRequest()
	}
	response = NewDescribeResourceBatchOperationResponse()
	err = c.Send(request, response)
	return
}

func NewListPkgRequest() (request *ListPkgRequest) {
	request = &ListPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListPkg")
	return
}

func NewListPkgResponse() (response *ListPkgResponse) {
	response = &ListPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询程序包列表
func (c *Client) ListPkg(request *ListPkgRequest) (response *ListPkgResponse, err error) {
	if request == nil {
		request = NewListPkgRequest()
	}
	response = NewListPkgResponse()
	err = c.Send(request, response)
	return
}

func NewSearchBusinessLogRequest() (request *SearchBusinessLogRequest) {
	request = &SearchBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchBusinessLog")
	return
}

func NewSearchBusinessLogResponse() (response *SearchBusinessLogResponse) {
	response = &SearchBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务日志搜索
func (c *Client) SearchBusinessLog(request *SearchBusinessLogRequest) (response *SearchBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchBusinessLogRequest()
	}
	response = NewSearchBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFileConfigRequest() (request *DeleteFileConfigRequest) {
	request = &DeleteFileConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteFileConfig")
	return
}

func NewDeleteFileConfigResponse() (response *DeleteFileConfigResponse) {
	response = &DeleteFileConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除文件配置项
func (c *Client) DeleteFileConfig(request *DeleteFileConfigRequest) (response *DeleteFileConfigResponse, err error) {
	if request == nil {
		request = NewDeleteFileConfigRequest()
	}
	response = NewDeleteFileConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPathRewriteRequest() (request *ModifyPathRewriteRequest) {
	request = &ModifyPathRewriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyPathRewrite")
	return
}

func NewModifyPathRewriteResponse() (response *ModifyPathRewriteResponse) {
	response = &ModifyPathRewriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改路径重写
func (c *Client) ModifyPathRewrite(request *ModifyPathRewriteRequest) (response *ModifyPathRewriteResponse, err error) {
	if request == nil {
		request = NewModifyPathRewriteRequest()
	}
	response = NewModifyPathRewriteResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCircuitBreakerRuleRequest() (request *CreateCircuitBreakerRuleRequest) {
	request = &CreateCircuitBreakerRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateCircuitBreakerRule")
	return
}

func NewCreateCircuitBreakerRuleResponse() (response *CreateCircuitBreakerRuleResponse) {
	response = &CreateCircuitBreakerRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建熔断规则
func (c *Client) CreateCircuitBreakerRule(request *CreateCircuitBreakerRuleRequest) (response *CreateCircuitBreakerRuleResponse, err error) {
	if request == nil {
		request = NewCreateCircuitBreakerRuleRequest()
	}
	response = NewCreateCircuitBreakerRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventStatisticsRequest() (request *DescribeEventStatisticsRequest) {
	request = &DescribeEventStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeEventStatistics")
	return
}

func NewDescribeEventStatisticsResponse() (response *DescribeEventStatisticsResponse) {
	response = &DescribeEventStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 事件统计
func (c *Client) DescribeEventStatistics(request *DescribeEventStatisticsRequest) (response *DescribeEventStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeEventStatisticsRequest()
	}
	response = NewDescribeEventStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteTaskFlowRequest() (request *ExecuteTaskFlowRequest) {
	request = &ExecuteTaskFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ExecuteTaskFlow")
	return
}

func NewExecuteTaskFlowResponse() (response *ExecuteTaskFlowResponse) {
	response = &ExecuteTaskFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行一次工作流
func (c *Client) ExecuteTaskFlow(request *ExecuteTaskFlowRequest) (response *ExecuteTaskFlowResponse, err error) {
	if request == nil {
		request = NewExecuteTaskFlowRequest()
	}
	response = NewExecuteTaskFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowAvailableTasksRequest() (request *DescribeFlowAvailableTasksRequest) {
	request = &DescribeFlowAvailableTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowAvailableTasks")
	return
}

func NewDescribeFlowAvailableTasksResponse() (response *DescribeFlowAvailableTasksResponse) {
	response = &DescribeFlowAvailableTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 翻页查询可用于工作流的任务
func (c *Client) DescribeFlowAvailableTasks(request *DescribeFlowAvailableTasksRequest) (response *DescribeFlowAvailableTasksResponse, err error) {
	if request == nil {
		request = NewDescribeFlowAvailableTasksRequest()
	}
	response = NewDescribeFlowAvailableTasksResponse()
	err = c.Send(request, response)
	return
}

func NewAdjustProductNewsPriorityRequest() (request *AdjustProductNewsPriorityRequest) {
	request = &AdjustProductNewsPriorityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AdjustProductNewsPriority")
	return
}

func NewAdjustProductNewsPriorityResponse() (response *AdjustProductNewsPriorityResponse) {
	response = &AdjustProductNewsPriorityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调整产品动态优先级
func (c *Client) AdjustProductNewsPriority(request *AdjustProductNewsPriorityRequest) (response *AdjustProductNewsPriorityResponse, err error) {
	if request == nil {
		request = NewAdjustProductNewsPriorityRequest()
	}
	response = NewAdjustProductNewsPriorityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePriceRequest() (request *DescribePriceRequest) {
	request = &DescribePriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePrice")
	return
}

func NewDescribePriceResponse() (response *DescribePriceResponse) {
	response = &DescribePriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询价格
func (c *Client) DescribePrice(request *DescribePriceRequest) (response *DescribePriceResponse, err error) {
	if request == nil {
		request = NewDescribePriceRequest()
	}
	response = NewDescribePriceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGroupSecretRequest() (request *CreateGroupSecretRequest) {
	request = &CreateGroupSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateGroupSecret")
	return
}

func NewCreateGroupSecretResponse() (response *CreateGroupSecretResponse) {
	response = &CreateGroupSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建API分组的秘钥
func (c *Client) CreateGroupSecret(request *CreateGroupSecretRequest) (response *CreateGroupSecretResponse, err error) {
	if request == nil {
		request = NewCreateGroupSecretRequest()
	}
	response = NewCreateGroupSecretResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmOverviewListRequest() (request *DescribeAlarmOverviewListRequest) {
	request = &DescribeAlarmOverviewListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAlarmOverviewList")
	return
}

func NewDescribeAlarmOverviewListResponse() (response *DescribeAlarmOverviewListResponse) {
	response = &DescribeAlarmOverviewListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询概览页告警信息列表
func (c *Client) DescribeAlarmOverviewList(request *DescribeAlarmOverviewListRequest) (response *DescribeAlarmOverviewListResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmOverviewListRequest()
	}
	response = NewDescribeAlarmOverviewListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLaneGroupExistRequest() (request *DescribeLaneGroupExistRequest) {
	request = &DescribeLaneGroupExistRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLaneGroupExist")
	return
}

func NewDescribeLaneGroupExistResponse() (response *DescribeLaneGroupExistResponse) {
	response = &DescribeLaneGroupExistResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询泳道部署组是否存在
func (c *Client) DescribeLaneGroupExist(request *DescribeLaneGroupExistRequest) (response *DescribeLaneGroupExistResponse, err error) {
	if request == nil {
		request = NewDescribeLaneGroupExistRequest()
	}
	response = NewDescribeLaneGroupExistResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSimpleNamespacesRequest() (request *DescribeSimpleNamespacesRequest) {
	request = &DescribeSimpleNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleNamespaces")
	return
}

func NewDescribeSimpleNamespacesResponse() (response *DescribeSimpleNamespacesResponse) {
	response = &DescribeSimpleNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询简单命名空间列表
func (c *Client) DescribeSimpleNamespaces(request *DescribeSimpleNamespacesRequest) (response *DescribeSimpleNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeSimpleNamespacesRequest()
	}
	response = NewDescribeSimpleNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProductNewsRequest() (request *DeleteProductNewsRequest) {
	request = &DeleteProductNewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteProductNews")
	return
}

func NewDeleteProductNewsResponse() (response *DeleteProductNewsResponse) {
	response = &DeleteProductNewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除产品动态
func (c *Client) DeleteProductNews(request *DeleteProductNewsRequest) (response *DeleteProductNewsResponse, err error) {
	if request == nil {
		request = NewDeleteProductNewsRequest()
	}
	response = NewDeleteProductNewsResponse()
	err = c.Send(request, response)
	return
}

func NewAutoRetryTransactionRequest() (request *AutoRetryTransactionRequest) {
	request = &AutoRetryTransactionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AutoRetryTransaction")
	return
}

func NewAutoRetryTransactionResponse() (response *AutoRetryTransactionResponse) {
	response = &AutoRetryTransactionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启事务自动重试
func (c *Client) AutoRetryTransaction(request *AutoRetryTransactionRequest) (response *AutoRetryTransactionResponse, err error) {
	if request == nil {
		request = NewAutoRetryTransactionRequest()
	}
	response = NewAutoRetryTransactionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupAttributeRequest() (request *DescribeGroupAttributeRequest) {
	request = &DescribeGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupAttribute")
	return
}

func NewDescribeGroupAttributeResponse() (response *DescribeGroupAttributeResponse) {
	response = &DescribeGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取部署组其他属性
func (c *Client) DescribeGroupAttribute(request *DescribeGroupAttributeRequest) (response *DescribeGroupAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeGroupAttributeRequest()
	}
	response = NewDescribeGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackFileConfigRequest() (request *RollbackFileConfigRequest) {
	request = &RollbackFileConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RollbackFileConfig")
	return
}

func NewRollbackFileConfigResponse() (response *RollbackFileConfigResponse) {
	response = &RollbackFileConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚文件配置
func (c *Client) RollbackFileConfig(request *RollbackFileConfigRequest) (response *RollbackFileConfigResponse, err error) {
	if request == nil {
		request = NewRollbackFileConfigRequest()
	}
	response = NewRollbackFileConfigResponse()
	err = c.Send(request, response)
	return
}

func NewStopBuildTaskRequest() (request *StopBuildTaskRequest) {
	request = &StopBuildTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StopBuildTask")
	return
}

func NewStopBuildTaskResponse() (response *StopBuildTaskResponse) {
	response = &StopBuildTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止构建任务
func (c *Client) StopBuildTask(request *StopBuildTaskRequest) (response *StopBuildTaskResponse, err error) {
	if request == nil {
		request = NewStopBuildTaskRequest()
	}
	response = NewStopBuildTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiAccessRequest() (request *DescribeApiAccessRequest) {
	request = &DescribeApiAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiAccess")
	return
}

func NewDescribeApiAccessResponse() (response *DescribeApiAccessResponse) {
	response = &DescribeApiAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Serverless应用的外网访问权限状态
func (c *Client) DescribeApiAccess(request *DescribeApiAccessRequest) (response *DescribeApiAccessResponse, err error) {
	if request == nil {
		request = NewDescribeApiAccessRequest()
	}
	response = NewDescribeApiAccessResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRouteRuleRequest() (request *ModifyRouteRuleRequest) {
	request = &ModifyRouteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyRouteRule")
	return
}

func NewModifyRouteRuleResponse() (response *ModifyRouteRuleResponse) {
	response = &ModifyRouteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新服务路由规则
func (c *Client) ModifyRouteRule(request *ModifyRouteRuleRequest) (response *ModifyRouteRuleResponse, err error) {
	if request == nil {
		request = NewModifyRouteRuleRequest()
	}
	response = NewModifyRouteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
	request = &DescribeGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroup")
	return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
	response = &DescribeGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询虚拟机部署组详情
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
	if request == nil {
		request = NewDescribeGroupRequest()
	}
	response = NewDescribeGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayDailyUseStatisticsRequest() (request *DescribeGatewayDailyUseStatisticsRequest) {
	request = &DescribeGatewayDailyUseStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayDailyUseStatistics")
	return
}

func NewDescribeGatewayDailyUseStatisticsResponse() (response *DescribeGatewayDailyUseStatisticsResponse) {
	response = &DescribeGatewayDailyUseStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关分组或者api日使用统计数据
func (c *Client) DescribeGatewayDailyUseStatistics(request *DescribeGatewayDailyUseStatisticsRequest) (response *DescribeGatewayDailyUseStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayDailyUseStatisticsRequest()
	}
	response = NewDescribeGatewayDailyUseStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceRequestRequest() (request *DescribeInstanceRequestRequest) {
	request = &DescribeInstanceRequestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInstanceRequest")
	return
}

func NewDescribeInstanceRequestResponse() (response *DescribeInstanceRequestResponse) {
	response = &DescribeInstanceRequestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例请求详情
func (c *Client) DescribeInstanceRequest(request *DescribeInstanceRequestRequest) (response *DescribeInstanceRequestResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceRequestRequest()
	}
	response = NewDescribeInstanceRequestResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskGroupsRequest() (request *DescribeTaskGroupsRequest) {
	request = &DescribeTaskGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskGroups")
	return
}

func NewDescribeTaskGroupsResponse() (response *DescribeTaskGroupsResponse) {
	response = &DescribeTaskGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前租户下任务所在的全部分组列表。
func (c *Client) DescribeTaskGroups(request *DescribeTaskGroupsRequest) (response *DescribeTaskGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskGroupsRequest()
	}
	response = NewDescribeTaskGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeployGroupRequest() (request *DeployGroupRequest) {
	request = &DeployGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeployGroup")
	return
}

func NewDeployGroupResponse() (response *DeployGroupResponse) {
	response = &DeployGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署虚拟机部署组应用
func (c *Client) DeployGroup(request *DeployGroupRequest) (response *DeployGroupResponse, err error) {
	if request == nil {
		request = NewDeployGroupRequest()
	}
	response = NewDeployGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterRequest() (request *ModifyClusterRequest) {
	request = &ModifyClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyCluster")
	return
}

func NewModifyClusterResponse() (response *ModifyClusterResponse) {
	response = &ModifyClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群信息
func (c *Client) ModifyCluster(request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
	if request == nil {
		request = NewModifyClusterRequest()
	}
	response = NewModifyClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerGroupsRequest() (request *DescribeContainerGroupsRequest) {
	request = &DescribeContainerGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroups")
	return
}

func NewDescribeContainerGroupsResponse() (response *DescribeContainerGroupsResponse) {
	response = &DescribeContainerGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器部署组列表
func (c *Client) DescribeContainerGroups(request *DescribeContainerGroupsRequest) (response *DescribeContainerGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeContainerGroupsRequest()
	}
	response = NewDescribeContainerGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerGroupDeployInfoRequest() (request *DescribeContainerGroupDeployInfoRequest) {
	request = &DescribeContainerGroupDeployInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDeployInfo")
	return
}

func NewDescribeContainerGroupDeployInfoResponse() (response *DescribeContainerGroupDeployInfoResponse) {
	response = &DescribeContainerGroupDeployInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//  获取部署组详情
func (c *Client) DescribeContainerGroupDeployInfo(request *DescribeContainerGroupDeployInfoRequest) (response *DescribeContainerGroupDeployInfoResponse, err error) {
	if request == nil {
		request = NewDescribeContainerGroupDeployInfoRequest()
	}
	response = NewDescribeContainerGroupDeployInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServicesRequest() (request *DescribeServicesRequest) {
	request = &DescribeServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeServices")
	return
}

func NewDescribeServicesResponse() (response *DescribeServicesResponse) {
	response = &DescribeServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品列表
func (c *Client) DescribeServices(request *DescribeServicesRequest) (response *DescribeServicesResponse, err error) {
	if request == nil {
		request = NewDescribeServicesRequest()
	}
	response = NewDescribeServicesResponse()
	err = c.Send(request, response)
	return
}

func NewShirkGroupRequest() (request *ShirkGroupRequest) {
	request = &ShirkGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ShirkGroup")
	return
}

func NewShirkGroupResponse() (response *ShirkGroupResponse) {
	response = &ShirkGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下线部署组所有机器实例
func (c *Client) ShirkGroup(request *ShirkGroupRequest) (response *ShirkGroupResponse, err error) {
	if request == nil {
		request = NewShirkGroupRequest()
	}
	response = NewShirkGroupResponse()
	err = c.Send(request, response)
	return
}

func NewAdjustProductHelpPriorityRequest() (request *AdjustProductHelpPriorityRequest) {
	request = &AdjustProductHelpPriorityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AdjustProductHelpPriority")
	return
}

func NewAdjustProductHelpPriorityResponse() (response *AdjustProductHelpPriorityResponse) {
	response = &AdjustProductHelpPriorityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调整产品帮助优先级
func (c *Client) AdjustProductHelpPriority(request *AdjustProductHelpPriorityRequest) (response *AdjustProductHelpPriorityResponse, err error) {
	if request == nil {
		request = NewAdjustProductHelpPriorityRequest()
	}
	response = NewAdjustProductHelpPriorityResponse()
	err = c.Send(request, response)
	return
}

func NewRevocationConfigRequest() (request *RevocationConfigRequest) {
	request = &RevocationConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RevocationConfig")
	return
}

func NewRevocationConfigResponse() (response *RevocationConfigResponse) {
	response = &RevocationConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 撤回已发布的配置
func (c *Client) RevocationConfig(request *RevocationConfigRequest) (response *RevocationConfigResponse, err error) {
	if request == nil {
		request = NewRevocationConfigRequest()
	}
	response = NewRevocationConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnitRulesRequest() (request *DescribeUnitRulesRequest) {
	request = &DescribeUnitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRules")
	return
}

func NewDescribeUnitRulesResponse() (response *DescribeUnitRulesResponse) {
	response = &DescribeUnitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单元化规则列表
func (c *Client) DescribeUnitRules(request *DescribeUnitRulesRequest) (response *DescribeUnitRulesResponse, err error) {
	if request == nil {
		request = NewDescribeUnitRulesRequest()
	}
	response = NewDescribeUnitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewImageRegisterRequest() (request *ImageRegisterRequest) {
	request = &ImageRegisterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ImageRegister")
	return
}

func NewImageRegisterResponse() (response *ImageRegisterResponse) {
	response = &ImageRegisterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像用户注册
func (c *Client) ImageRegister(request *ImageRegisterRequest) (response *ImageRegisterResponse, err error) {
	if request == nil {
		request = NewImageRegisterRequest()
	}
	response = NewImageRegisterResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateBusinessLogConfigRequest() (request *AssociateBusinessLogConfigRequest) {
	request = &AssociateBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AssociateBusinessLogConfig")
	return
}

func NewAssociateBusinessLogConfigResponse() (response *AssociateBusinessLogConfigResponse) {
	response = &AssociateBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联日志配置项到应用
func (c *Client) AssociateBusinessLogConfig(request *AssociateBusinessLogConfigRequest) (response *AssociateBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewAssociateBusinessLogConfigRequest()
	}
	response = NewAssociateBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewAddInstancesRequest() (request *AddInstancesRequest) {
	request = &AddInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AddInstances")
	return
}

func NewAddInstancesResponse() (response *AddInstancesResponse) {
	response = &AddInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加云主机节点至TSF集群
func (c *Client) AddInstances(request *AddInstancesRequest) (response *AddInstancesResponse, err error) {
	if request == nil {
		request = NewAddInstancesRequest()
	}
	response = NewAddInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespaceAffinityRequest() (request *DescribeNamespaceAffinityRequest) {
	request = &DescribeNamespaceAffinityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaceAffinity")
	return
}

func NewDescribeNamespaceAffinityResponse() (response *DescribeNamespaceAffinityResponse) {
	response = &DescribeNamespaceAffinityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询命名空间就近访问策略
func (c *Client) DescribeNamespaceAffinity(request *DescribeNamespaceAffinityRequest) (response *DescribeNamespaceAffinityResponse, err error) {
	if request == nil {
		request = NewDescribeNamespaceAffinityRequest()
	}
	response = NewDescribeNamespaceAffinityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageFeaturesRequest() (request *DescribeImageFeaturesRequest) {
	request = &DescribeImageFeaturesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageFeatures")
	return
}

func NewDescribeImageFeaturesResponse() (response *DescribeImageFeaturesResponse) {
	response = &DescribeImageFeaturesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像特征列表，比如支持JDK的类型，konaJDK或openJDK
func (c *Client) DescribeImageFeatures(request *DescribeImageFeaturesRequest) (response *DescribeImageFeaturesResponse, err error) {
	if request == nil {
		request = NewDescribeImageFeaturesRequest()
	}
	response = NewDescribeImageFeaturesResponse()
	err = c.Send(request, response)
	return
}

func NewInstallAgentScriptsRequest() (request *InstallAgentScriptsRequest) {
	request = &InstallAgentScriptsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "InstallAgentScripts")
	return
}

func NewInstallAgentScriptsResponse() (response *InstallAgentScriptsResponse) {
	response = &InstallAgentScriptsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 手动安装agent脚本
func (c *Client) InstallAgentScripts(request *InstallAgentScriptsRequest) (response *InstallAgentScriptsResponse, err error) {
	if request == nil {
		request = NewInstallAgentScriptsRequest()
	}
	response = NewInstallAgentScriptsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
	request = &CreateRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateRole")
	return
}

func NewCreateRoleResponse() (response *CreateRoleResponse) {
	response = &CreateRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建角色
func (c *Client) CreateRole(request *CreateRoleRequest) (response *CreateRoleResponse, err error) {
	if request == nil {
		request = NewCreateRoleRequest()
	}
	response = NewCreateRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsableUnitNamespacesRequest() (request *DescribeUsableUnitNamespacesRequest) {
	request = &DescribeUsableUnitNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUsableUnitNamespaces")
	return
}

func NewDescribeUsableUnitNamespacesResponse() (response *DescribeUsableUnitNamespacesResponse) {
	response = &DescribeUsableUnitNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用于被导入的命名空间列表
func (c *Client) DescribeUsableUnitNamespaces(request *DescribeUsableUnitNamespacesRequest) (response *DescribeUsableUnitNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeUsableUnitNamespacesRequest()
	}
	response = NewDescribeUsableUnitNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayApiRequest() (request *UpdateGatewayApiRequest) {
	request = &UpdateGatewayApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayApi")
	return
}

func NewUpdateGatewayApiResponse() (response *UpdateGatewayApiResponse) {
	response = &UpdateGatewayApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新API
func (c *Client) UpdateGatewayApi(request *UpdateGatewayApiRequest) (response *UpdateGatewayApiResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayApiRequest()
	}
	response = NewUpdateGatewayApiResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTransactionRequest() (request *DescribeTransactionRequest) {
	request = &DescribeTransactionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTransaction")
	return
}

func NewDescribeTransactionResponse() (response *DescribeTransactionResponse) {
	response = &DescribeTransactionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询事务详情
func (c *Client) DescribeTransaction(request *DescribeTransactionRequest) (response *DescribeTransactionResponse, err error) {
	if request == nil {
		request = NewDescribeTransactionRequest()
	}
	response = NewDescribeTransactionResponse()
	err = c.Send(request, response)
	return
}

func NewPingTsfByMultiCloudRequest() (request *PingTsfByMultiCloudRequest) {
	request = &PingTsfByMultiCloudRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "PingTsfByMultiCloud")
	return
}

func NewPingTsfByMultiCloudResponse() (response *PingTsfByMultiCloudResponse) {
	response = &PingTsfByMultiCloudResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查各地域TSF连通性(超云)
func (c *Client) PingTsfByMultiCloud(request *PingTsfByMultiCloudRequest) (response *PingTsfByMultiCloudResponse, err error) {
	if request == nil {
		request = NewPingTsfByMultiCloudRequest()
	}
	response = NewPingTsfByMultiCloudResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasedConfigRequest() (request *DescribeReleasedConfigRequest) {
	request = &DescribeReleasedConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasedConfig")
	return
}

func NewDescribeReleasedConfigResponse() (response *DescribeReleasedConfigResponse) {
	response = &DescribeReleasedConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询group发布的配置
func (c *Client) DescribeReleasedConfig(request *DescribeReleasedConfigRequest) (response *DescribeReleasedConfigResponse, err error) {
	if request == nil {
		request = NewDescribeReleasedConfigRequest()
	}
	response = NewDescribeReleasedConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePublicConfigRequest() (request *CreatePublicConfigRequest) {
	request = &CreatePublicConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreatePublicConfig")
	return
}

func NewCreatePublicConfigResponse() (response *CreatePublicConfigResponse) {
	response = &CreatePublicConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建公共配置项
func (c *Client) CreatePublicConfig(request *CreatePublicConfigRequest) (response *CreatePublicConfigResponse, err error) {
	if request == nil {
		request = NewCreatePublicConfigRequest()
	}
	response = NewCreatePublicConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSdkVersionRequest() (request *DescribeSdkVersionRequest) {
	request = &DescribeSdkVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSdkVersion")
	return
}

func NewDescribeSdkVersionResponse() (response *DescribeSdkVersionResponse) {
	response = &DescribeSdkVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询SDK版本
func (c *Client) DescribeSdkVersion(request *DescribeSdkVersionRequest) (response *DescribeSdkVersionResponse, err error) {
	if request == nil {
		request = NewDescribeSdkVersionRequest()
	}
	response = NewDescribeSdkVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUploadInfoRequest() (request *DescribeUploadInfoRequest) {
	request = &DescribeUploadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUploadInfo")
	return
}

func NewDescribeUploadInfoResponse() (response *DescribeUploadInfoResponse) {
	response = &DescribeUploadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF会将软件包上传到腾讯云对象存储（COS）。调用此接口获取上传信息，如目标地域，桶，包Id，存储路径，鉴权信息等，之后请使用COS API（或SDK）进行上传。
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
func (c *Client) DescribeUploadInfo(request *DescribeUploadInfoRequest) (response *DescribeUploadInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUploadInfoRequest()
	}
	response = NewDescribeUploadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewSetBuildTaskFailedRequest() (request *SetBuildTaskFailedRequest) {
	request = &SetBuildTaskFailedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SetBuildTaskFailed")
	return
}

func NewSetBuildTaskFailedResponse() (response *SetBuildTaskFailedResponse) {
	response = &SetBuildTaskFailedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置构建任务失败
func (c *Client) SetBuildTaskFailed(request *SetBuildTaskFailedRequest) (response *SetBuildTaskFailedResponse, err error) {
	if request == nil {
		request = NewSetBuildTaskFailedRequest()
	}
	response = NewSetBuildTaskFailedResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateFilebeatConfigRequest() (request *DisassociateFilebeatConfigRequest) {
	request = &DisassociateFilebeatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisassociateFilebeatConfig")
	return
}

func NewDisassociateFilebeatConfigResponse() (response *DisassociateFilebeatConfigResponse) {
	response = &DisassociateFilebeatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消通用FilebeatConfig和部署组的关联
func (c *Client) DisassociateFilebeatConfig(request *DisassociateFilebeatConfigRequest) (response *DisassociateFilebeatConfigResponse, err error) {
	if request == nil {
		request = NewDisassociateFilebeatConfigRequest()
	}
	response = NewDisassociateFilebeatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewEnableTaskRequest() (request *EnableTaskRequest) {
	request = &EnableTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableTask")
	return
}

func NewEnableTaskResponse() (response *EnableTaskResponse) {
	response = &EnableTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用任务
func (c *Client) EnableTask(request *EnableTaskRequest) (response *EnableTaskResponse, err error) {
	if request == nil {
		request = NewEnableTaskRequest()
	}
	response = NewEnableTaskResponse()
	err = c.Send(request, response)
	return
}

func NewChangeApiUsableStatusRequest() (request *ChangeApiUsableStatusRequest) {
	request = &ChangeApiUsableStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ChangeApiUsableStatus")
	return
}

func NewChangeApiUsableStatusResponse() (response *ChangeApiUsableStatusResponse) {
	response = &ChangeApiUsableStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用或禁用API
func (c *Client) ChangeApiUsableStatus(request *ChangeApiUsableStatusRequest) (response *ChangeApiUsableStatusResponse, err error) {
	if request == nil {
		request = NewChangeApiUsableStatusRequest()
	}
	response = NewChangeApiUsableStatusResponse()
	err = c.Send(request, response)
	return
}

func NewUploadLicenseApplicationRequest() (request *UploadLicenseApplicationRequest) {
	request = &UploadLicenseApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UploadLicenseApplication")
	return
}

func NewUploadLicenseApplicationResponse() (response *UploadLicenseApplicationResponse) {
	response = &UploadLicenseApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 递交许可申请
func (c *Client) UploadLicenseApplication(request *UploadLicenseApplicationRequest) (response *UploadLicenseApplicationResponse, err error) {
	if request == nil {
		request = NewUploadLicenseApplicationRequest()
	}
	response = NewUploadLicenseApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTsfmanagerZonesRequest() (request *DescribeTsfmanagerZonesRequest) {
	request = &DescribeTsfmanagerZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTsfmanagerZones")
	return
}

func NewDescribeTsfmanagerZonesResponse() (response *DescribeTsfmanagerZonesResponse) {
	response = &DescribeTsfmanagerZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TSF可用区列表
func (c *Client) DescribeTsfmanagerZones(request *DescribeTsfmanagerZonesRequest) (response *DescribeTsfmanagerZonesResponse, err error) {
	if request == nil {
		request = NewDescribeTsfmanagerZonesRequest()
	}
	response = NewDescribeTsfmanagerZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDisableUnitRuleRequest() (request *DisableUnitRuleRequest) {
	request = &DisableUnitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableUnitRule")
	return
}

func NewDisableUnitRuleResponse() (response *DisableUnitRuleResponse) {
	response = &DisableUnitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用单元化规则
func (c *Client) DisableUnitRule(request *DisableUnitRuleRequest) (response *DisableUnitRuleResponse, err error) {
	if request == nil {
		request = NewDisableUnitRuleRequest()
	}
	response = NewDisableUnitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRegionRequest() (request *CreateRegionRequest) {
	request = &CreateRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateRegion")
	return
}

func NewCreateRegionResponse() (response *CreateRegionResponse) {
	response = &CreateRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加业务地域
func (c *Client) CreateRegion(request *CreateRegionRequest) (response *CreateRegionResponse, err error) {
	if request == nil {
		request = NewCreateRegionRequest()
	}
	response = NewCreateRegionResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateBusinessLogConfigRequest() (request *UpdateBusinessLogConfigRequest) {
	request = &UpdateBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateBusinessLogConfig")
	return
}

func NewUpdateBusinessLogConfigResponse() (response *UpdateBusinessLogConfigResponse) {
	response = &UpdateBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新业务日志配置
func (c *Client) UpdateBusinessLogConfig(request *UpdateBusinessLogConfigRequest) (response *UpdateBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewUpdateBusinessLogConfigRequest()
	}
	response = NewUpdateBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileConfigsRequest() (request *DescribeFileConfigsRequest) {
	request = &DescribeFileConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigs")
	return
}

func NewDescribeFileConfigsResponse() (response *DescribeFileConfigsResponse) {
	response = &DescribeFileConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件配置项列表
func (c *Client) DescribeFileConfigs(request *DescribeFileConfigsRequest) (response *DescribeFileConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeFileConfigsRequest()
	}
	response = NewDescribeFileConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlarmPolicyRequest() (request *CreateAlarmPolicyRequest) {
	request = &CreateAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateAlarmPolicy")
	return
}

func NewCreateAlarmPolicyResponse() (response *CreateAlarmPolicyResponse) {
	response = &CreateAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建告警策略
func (c *Client) CreateAlarmPolicy(request *CreateAlarmPolicyRequest) (response *CreateAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewCreateAlarmPolicyRequest()
	}
	response = NewCreateAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	response = NewCreateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUnitNamespacesRequest() (request *CreateUnitNamespacesRequest) {
	request = &CreateUnitNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateUnitNamespaces")
	return
}

func NewCreateUnitNamespacesResponse() (response *CreateUnitNamespacesResponse) {
	response = &CreateUnitNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量创建单元化命名空间
func (c *Client) CreateUnitNamespaces(request *CreateUnitNamespacesRequest) (response *CreateUnitNamespacesResponse, err error) {
	if request == nil {
		request = NewCreateUnitNamespacesRequest()
	}
	response = NewCreateUnitNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewGetOssTopologyGraphRequest() (request *GetOssTopologyGraphRequest) {
	request = &GetOssTopologyGraphRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetOssTopologyGraph")
	return
}

func NewGetOssTopologyGraphResponse() (response *GetOssTopologyGraphResponse) {
	response = &GetOssTopologyGraphResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询依赖拓扑图
func (c *Client) GetOssTopologyGraph(request *GetOssTopologyGraphRequest) (response *GetOssTopologyGraphResponse, err error) {
	if request == nil {
		request = NewGetOssTopologyGraphRequest()
	}
	response = NewGetOssTopologyGraphResponse()
	err = c.Send(request, response)
	return
}

func NewCreateContainerShellSessionRequest() (request *CreateContainerShellSessionRequest) {
	request = &CreateContainerShellSessionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateContainerShellSession")
	return
}

func NewCreateContainerShellSessionResponse() (response *CreateContainerShellSessionResponse) {
	response = &CreateContainerShellSessionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建容器登录的 session
func (c *Client) CreateContainerShellSession(request *CreateContainerShellSessionRequest) (response *CreateContainerShellSessionResponse, err error) {
	if request == nil {
		request = NewCreateContainerShellSessionRequest()
	}
	response = NewCreateContainerShellSessionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicatoinServerLogRequest() (request *DescribeApplicatoinServerLogRequest) {
	request = &DescribeApplicatoinServerLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicatoinServerLog")
	return
}

func NewDescribeApplicatoinServerLogResponse() (response *DescribeApplicatoinServerLogResponse) {
	response = &DescribeApplicatoinServerLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用服务器的日志
func (c *Client) DescribeApplicatoinServerLog(request *DescribeApplicatoinServerLogRequest) (response *DescribeApplicatoinServerLogResponse, err error) {
	if request == nil {
		request = NewDescribeApplicatoinServerLogRequest()
	}
	response = NewDescribeApplicatoinServerLogResponse()
	err = c.Send(request, response)
	return
}

func NewEnableNamespaceAffinityRequest() (request *EnableNamespaceAffinityRequest) {
	request = &EnableNamespaceAffinityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableNamespaceAffinity")
	return
}

func NewEnableNamespaceAffinityResponse() (response *EnableNamespaceAffinityResponse) {
	response = &EnableNamespaceAffinityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用命名空间就近访问策略
func (c *Client) EnableNamespaceAffinity(request *EnableNamespaceAffinityRequest) (response *EnableNamespaceAffinityResponse, err error) {
	if request == nil {
		request = NewEnableNamespaceAffinityRequest()
	}
	response = NewEnableNamespaceAffinityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOverviewResourceRequest() (request *DescribeOverviewResourceRequest) {
	request = &DescribeOverviewResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeOverviewResource")
	return
}

func NewDescribeOverviewResourceResponse() (response *DescribeOverviewResourceResponse) {
	response = &DescribeOverviewResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览页资源信息
func (c *Client) DescribeOverviewResource(request *DescribeOverviewResourceRequest) (response *DescribeOverviewResourceResponse, err error) {
	if request == nil {
		request = NewDescribeOverviewResourceRequest()
	}
	response = NewDescribeOverviewResourceResponse()
	err = c.Send(request, response)
	return
}

func NewGetPkgInfoRequest() (request *GetPkgInfoRequest) {
	request = &GetPkgInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetPkgInfo")
	return
}

func NewGetPkgInfoResponse() (response *GetPkgInfoResponse) {
	response = &GetPkgInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询包详情
func (c *Client) GetPkgInfo(request *GetPkgInfoRequest) (response *GetPkgInfoResponse, err error) {
	if request == nil {
		request = NewGetPkgInfoRequest()
	}
	response = NewGetPkgInfoResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayQQMiniProgramLoginPluginRequest() (request *UpdateGatewayQQMiniProgramLoginPluginRequest) {
	request = &UpdateGatewayQQMiniProgramLoginPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayQQMiniProgramLoginPlugin")
	return
}

func NewUpdateGatewayQQMiniProgramLoginPluginResponse() (response *UpdateGatewayQQMiniProgramLoginPluginResponse) {
	response = &UpdateGatewayQQMiniProgramLoginPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改qq小程序登录插件
func (c *Client) UpdateGatewayQQMiniProgramLoginPlugin(request *UpdateGatewayQQMiniProgramLoginPluginRequest) (response *UpdateGatewayQQMiniProgramLoginPluginResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayQQMiniProgramLoginPluginRequest()
	}
	response = NewUpdateGatewayQQMiniProgramLoginPluginResponse()
	err = c.Send(request, response)
	return
}

func NewRunReleaseRequest() (request *RunReleaseRequest) {
	request = &RunReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RunRelease")
	return
}

func NewRunReleaseResponse() (response *RunReleaseResponse) {
	response = &RunReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 执行发布单
func (c *Client) RunRelease(request *RunReleaseRequest) (response *RunReleaseResponse, err error) {
	if request == nil {
		request = NewRunReleaseRequest()
	}
	response = NewRunReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterRequest() (request *DescribeClusterRequest) {
	request = &DescribeClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeCluster")
	return
}

func NewDescribeClusterResponse() (response *DescribeClusterResponse) {
	response = &DescribeClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群信息
func (c *Client) DescribeCluster(request *DescribeClusterRequest) (response *DescribeClusterResponse, err error) {
	if request == nil {
		request = NewDescribeClusterRequest()
	}
	response = NewDescribeClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSidecarFilterNamesRequest() (request *DescribeSidecarFilterNamesRequest) {
	request = &DescribeSidecarFilterNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSidecarFilterNames")
	return
}

func NewDescribeSidecarFilterNamesResponse() (response *DescribeSidecarFilterNamesResponse) {
	response = &DescribeSidecarFilterNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Sidecar过滤器名称列表
func (c *Client) DescribeSidecarFilterNames(request *DescribeSidecarFilterNamesRequest) (response *DescribeSidecarFilterNamesResponse, err error) {
	if request == nil {
		request = NewDescribeSidecarFilterNamesRequest()
	}
	response = NewDescribeSidecarFilterNamesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTsfZoneRequest() (request *ModifyTsfZoneRequest) {
	request = &ModifyTsfZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyTsfZone")
	return
}

func NewModifyTsfZoneResponse() (response *ModifyTsfZoneResponse) {
	response = &ModifyTsfZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改基础区域..
func (c *Client) ModifyTsfZone(request *ModifyTsfZoneRequest) (response *ModifyTsfZoneResponse, err error) {
	if request == nil {
		request = NewModifyTsfZoneRequest()
	}
	response = NewModifyTsfZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroserviceGroupsRequest() (request *DescribeMicroserviceGroupsRequest) {
	request = &DescribeMicroserviceGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroserviceGroups")
	return
}

func NewDescribeMicroserviceGroupsResponse() (response *DescribeMicroserviceGroupsResponse) {
	response = &DescribeMicroserviceGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微服务下的部署组列表
func (c *Client) DescribeMicroserviceGroups(request *DescribeMicroserviceGroupsRequest) (response *DescribeMicroserviceGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeMicroserviceGroupsRequest()
	}
	response = NewDescribeMicroserviceGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAuthorizationInfoRequest() (request *ModifyAuthorizationInfoRequest) {
	request = &ModifyAuthorizationInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyAuthorizationInfo")
	return
}

func NewModifyAuthorizationInfoResponse() (response *ModifyAuthorizationInfoResponse) {
	response = &ModifyAuthorizationInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置服务权限配置
func (c *Client) ModifyAuthorizationInfo(request *ModifyAuthorizationInfoRequest) (response *ModifyAuthorizationInfoResponse, err error) {
	if request == nil {
		request = NewModifyAuthorizationInfoRequest()
	}
	response = NewModifyAuthorizationInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetDockerForUseRequest() (request *GetDockerForUseRequest) {
	request = &GetDockerForUseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetDockerForUse")
	return
}

func NewGetDockerForUseResponse() (response *GetDockerForUseResponse) {
	response = &GetDockerForUseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取docker使用指引
func (c *Client) GetDockerForUse(request *GetDockerForUseRequest) (response *GetDockerForUseResponse, err error) {
	if request == nil {
		request = NewGetDockerForUseRequest()
	}
	response = NewGetDockerForUseResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGatewayPluginRequest() (request *DeleteGatewayPluginRequest) {
	request = &DeleteGatewayPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteGatewayPlugin")
	return
}

func NewDeleteGatewayPluginResponse() (response *DeleteGatewayPluginResponse) {
	response = &DeleteGatewayPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除网关插件
func (c *Client) DeleteGatewayPlugin(request *DeleteGatewayPluginRequest) (response *DeleteGatewayPluginResponse, err error) {
	if request == nil {
		request = NewDeleteGatewayPluginRequest()
	}
	response = NewDeleteGatewayPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterMonitorRequest() (request *DescribeClusterMonitorRequest) {
	request = &DescribeClusterMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterMonitor")
	return
}

func NewDescribeClusterMonitorResponse() (response *DescribeClusterMonitorResponse) {
	response = &DescribeClusterMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群监控数据
func (c *Client) DescribeClusterMonitor(request *DescribeClusterMonitorRequest) (response *DescribeClusterMonitorResponse, err error) {
	if request == nil {
		request = NewDescribeClusterMonitorRequest()
	}
	response = NewDescribeClusterMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewSearchOssSurroundBusinessLogRequest() (request *SearchOssSurroundBusinessLogRequest) {
	request = &SearchOssSurroundBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchOssSurroundBusinessLog")
	return
}

func NewSearchOssSurroundBusinessLogResponse() (response *SearchOssSurroundBusinessLogResponse) {
	response = &SearchOssSurroundBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务日志上下文检索
func (c *Client) SearchOssSurroundBusinessLog(request *SearchOssSurroundBusinessLogRequest) (response *SearchOssSurroundBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchOssSurroundBusinessLogRequest()
	}
	response = NewSearchOssSurroundBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFilebeatConfigEnableRequest() (request *DescribeFilebeatConfigEnableRequest) {
	request = &DescribeFilebeatConfigEnableRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFilebeatConfigEnable")
	return
}

func NewDescribeFilebeatConfigEnableResponse() (response *DescribeFilebeatConfigEnableResponse) {
	response = &DescribeFilebeatConfigEnableResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 判断filebeat自定义配置是否开启
func (c *Client) DescribeFilebeatConfigEnable(request *DescribeFilebeatConfigEnableRequest) (response *DescribeFilebeatConfigEnableResponse, err error) {
	if request == nil {
		request = NewDescribeFilebeatConfigEnableRequest()
	}
	response = NewDescribeFilebeatConfigEnableResponse()
	err = c.Send(request, response)
	return
}

func NewRedoTaskRequest() (request *RedoTaskRequest) {
	request = &RedoTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RedoTask")
	return
}

func NewRedoTaskResponse() (response *RedoTaskResponse) {
	response = &RedoTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新执行任务
func (c *Client) RedoTask(request *RedoTaskRequest) (response *RedoTaskResponse, err error) {
	if request == nil {
		request = NewRedoTaskRequest()
	}
	response = NewRedoTaskResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseProductHelpRequest() (request *ReleaseProductHelpRequest) {
	request = &ReleaseProductHelpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleaseProductHelp")
	return
}

func NewReleaseProductHelpResponse() (response *ReleaseProductHelpResponse) {
	response = &ReleaseProductHelpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 全量覆盖发布表数据
func (c *Client) ReleaseProductHelp(request *ReleaseProductHelpRequest) (response *ReleaseProductHelpResponse, err error) {
	if request == nil {
		request = NewReleaseProductHelpRequest()
	}
	response = NewReleaseProductHelpResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskRecordsRequest() (request *DescribeTaskRecordsRequest) {
	request = &DescribeTaskRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskRecords")
	return
}

func NewDescribeTaskRecordsResponse() (response *DescribeTaskRecordsResponse) {
	response = &DescribeTaskRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 翻页查询任务列表
func (c *Client) DescribeTaskRecords(request *DescribeTaskRecordsRequest) (response *DescribeTaskRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskRecordsRequest()
	}
	response = NewDescribeTaskRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSidecarFilterRequest() (request *DeleteSidecarFilterRequest) {
	request = &DeleteSidecarFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteSidecarFilter")
	return
}

func NewDeleteSidecarFilterResponse() (response *DeleteSidecarFilterResponse) {
	response = &DeleteSidecarFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Sidecar过滤器
func (c *Client) DeleteSidecarFilter(request *DeleteSidecarFilterRequest) (response *DeleteSidecarFilterResponse, err error) {
	if request == nil {
		request = NewDeleteSidecarFilterRequest()
	}
	response = NewDeleteSidecarFilterResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayOAuthPluginRequest() (request *UpdateGatewayOAuthPluginRequest) {
	request = &UpdateGatewayOAuthPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayOAuthPlugin")
	return
}

func NewUpdateGatewayOAuthPluginResponse() (response *UpdateGatewayOAuthPluginResponse) {
	response = &UpdateGatewayOAuthPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关OAuth插件
func (c *Client) UpdateGatewayOAuthPlugin(request *UpdateGatewayOAuthPluginRequest) (response *UpdateGatewayOAuthPluginResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayOAuthPluginRequest()
	}
	response = NewUpdateGatewayOAuthPluginResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTaskFlowRequest() (request *ModifyTaskFlowRequest) {
	request = &ModifyTaskFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyTaskFlow")
	return
}

func NewModifyTaskFlowResponse() (response *ModifyTaskFlowResponse) {
	response = &ModifyTaskFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改工作流
func (c *Client) ModifyTaskFlow(request *ModifyTaskFlowRequest) (response *ModifyTaskFlowResponse, err error) {
	if request == nil {
		request = NewModifyTaskFlowRequest()
	}
	response = NewModifyTaskFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicConfigRequest() (request *DescribePublicConfigRequest) {
	request = &DescribePublicConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfig")
	return
}

func NewDescribePublicConfigResponse() (response *DescribePublicConfigResponse) {
	response = &DescribePublicConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共配置（单条）
func (c *Client) DescribePublicConfig(request *DescribePublicConfigRequest) (response *DescribePublicConfigResponse, err error) {
	if request == nil {
		request = NewDescribePublicConfigRequest()
	}
	response = NewDescribePublicConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDisableRouteRuleRequest() (request *DisableRouteRuleRequest) {
	request = &DisableRouteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableRouteRule")
	return
}

func NewDisableRouteRuleResponse() (response *DisableRouteRuleResponse) {
	response = &DisableRouteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停用路由规则
func (c *Client) DisableRouteRule(request *DisableRouteRuleRequest) (response *DisableRouteRuleResponse, err error) {
	if request == nil {
		request = NewDisableRouteRuleRequest()
	}
	response = NewDisableRouteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewFindServiceMonitorObjectRequest() (request *FindServiceMonitorObjectRequest) {
	request = &FindServiceMonitorObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "FindServiceMonitorObject")
	return
}

func NewFindServiceMonitorObjectResponse() (response *FindServiceMonitorObjectResponse) {
	response = &FindServiceMonitorObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// FindServiceMonitorObject
func (c *Client) FindServiceMonitorObject(request *FindServiceMonitorObjectRequest) (response *FindServiceMonitorObjectResponse, err error) {
	if request == nil {
		request = NewFindServiceMonitorObjectRequest()
	}
	response = NewFindServiceMonitorObjectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayInstanceRequest() (request *DescribeGatewayInstanceRequest) {
	request = &DescribeGatewayInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayInstance")
	return
}

func NewDescribeGatewayInstanceResponse() (response *DescribeGatewayInstanceResponse) {
	response = &DescribeGatewayInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关实例
func (c *Client) DescribeGatewayInstance(request *DescribeGatewayInstanceRequest) (response *DescribeGatewayInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayInstanceRequest()
	}
	response = NewDescribeGatewayInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLaneRuleRequest() (request *DeleteLaneRuleRequest) {
	request = &DeleteLaneRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteLaneRule")
	return
}

func NewDeleteLaneRuleResponse() (response *DeleteLaneRuleResponse) {
	response = &DeleteLaneRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除泳道规则
func (c *Client) DeleteLaneRule(request *DeleteLaneRuleRequest) (response *DeleteLaneRuleResponse, err error) {
	if request == nil {
		request = NewDeleteLaneRuleRequest()
	}
	response = NewDeleteLaneRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplateRequest() (request *DescribeTemplateRequest) {
	request = &DescribeTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTemplate")
	return
}

func NewDescribeTemplateResponse() (response *DescribeTemplateResponse) {
	response = &DescribeTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取模板
func (c *Client) DescribeTemplate(request *DescribeTemplateRequest) (response *DescribeTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeTemplateRequest()
	}
	response = NewDescribeTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayTagPluginRequest() (request *CreateGatewayTagPluginRequest) {
	request = &CreateGatewayTagPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayTagPlugin")
	return
}

func NewCreateGatewayTagPluginResponse() (response *CreateGatewayTagPluginResponse) {
	response = &CreateGatewayTagPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增网关Tag转化插件
func (c *Client) CreateGatewayTagPlugin(request *CreateGatewayTagPluginRequest) (response *CreateGatewayTagPluginResponse, err error) {
	if request == nil {
		request = NewCreateGatewayTagPluginRequest()
	}
	response = NewCreateGatewayTagPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachinesRequest() (request *DeleteMachinesRequest) {
	request = &DeleteMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteMachines")
	return
}

func NewDeleteMachinesResponse() (response *DeleteMachinesResponse) {
	response = &DeleteMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机器
func (c *Client) DeleteMachines(request *DeleteMachinesRequest) (response *DeleteMachinesResponse, err error) {
	if request == nil {
		request = NewDeleteMachinesRequest()
	}
	response = NewDeleteMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopReqAmountInterfacesRequest() (request *GetTopReqAmountInterfacesRequest) {
	request = &GetTopReqAmountInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTopReqAmountInterfaces")
	return
}

func NewGetTopReqAmountInterfacesResponse() (response *GetTopReqAmountInterfacesResponse) {
	response = &GetTopReqAmountInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TopN请求量接口列表
func (c *Client) GetTopReqAmountInterfaces(request *GetTopReqAmountInterfacesRequest) (response *GetTopReqAmountInterfacesResponse, err error) {
	if request == nil {
		request = NewGetTopReqAmountInterfacesRequest()
	}
	response = NewGetTopReqAmountInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewStartInstanceRequest() (request *StartInstanceRequest) {
	request = &StartInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StartInstance")
	return
}

func NewStartInstanceResponse() (response *StartInstanceResponse) {
	response = &StartInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动机器
func (c *Client) StartInstance(request *StartInstanceRequest) (response *StartInstanceResponse, err error) {
	if request == nil {
		request = NewStartInstanceRequest()
	}
	response = NewStartInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayMonitorOverviewRequest() (request *DescribeGatewayMonitorOverviewRequest) {
	request = &DescribeGatewayMonitorOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayMonitorOverview")
	return
}

func NewDescribeGatewayMonitorOverviewResponse() (response *DescribeGatewayMonitorOverviewResponse) {
	response = &DescribeGatewayMonitorOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关监控概览
func (c *Client) DescribeGatewayMonitorOverview(request *DescribeGatewayMonitorOverviewRequest) (response *DescribeGatewayMonitorOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayMonitorOverviewRequest()
	}
	response = NewDescribeGatewayMonitorOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewSearchStdoutLogRequest() (request *SearchStdoutLogRequest) {
	request = &SearchStdoutLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchStdoutLog")
	return
}

func NewSearchStdoutLogResponse() (response *SearchStdoutLogResponse) {
	response = &SearchStdoutLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标准输出日志搜索
func (c *Client) SearchStdoutLog(request *SearchStdoutLogRequest) (response *SearchStdoutLogResponse, err error) {
	if request == nil {
		request = NewSearchStdoutLogRequest()
	}
	response = NewSearchStdoutLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasePipelineRunRequest() (request *DescribeReleasePipelineRunRequest) {
	request = &DescribeReleasePipelineRunRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasePipelineRun")
	return
}

func NewDescribeReleasePipelineRunResponse() (response *DescribeReleasePipelineRunResponse) {
	response = &DescribeReleasePipelineRunResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布单执行详情
func (c *Client) DescribeReleasePipelineRun(request *DescribeReleasePipelineRunRequest) (response *DescribeReleasePipelineRunResponse, err error) {
	if request == nil {
		request = NewDescribeReleasePipelineRunRequest()
	}
	response = NewDescribeReleasePipelineRunResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServerlessGroupRequest() (request *DeleteServerlessGroupRequest) {
	request = &DeleteServerlessGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteServerlessGroup")
	return
}

func NewDeleteServerlessGroupResponse() (response *DeleteServerlessGroupResponse) {
	response = &DeleteServerlessGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Serverless部署组
func (c *Client) DeleteServerlessGroup(request *DeleteServerlessGroupRequest) (response *DeleteServerlessGroupResponse, err error) {
	if request == nil {
		request = NewDeleteServerlessGroupRequest()
	}
	response = NewDeleteServerlessGroupResponse()
	err = c.Send(request, response)
	return
}

func NewGetTraceInterfacesRequest() (request *GetTraceInterfacesRequest) {
	request = &GetTraceInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTraceInterfaces")
	return
}

func NewGetTraceInterfacesResponse() (response *GetTraceInterfacesResponse) {
	response = &GetTraceInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用链接口列表
func (c *Client) GetTraceInterfaces(request *GetTraceInterfacesRequest) (response *GetTraceInterfacesResponse, err error) {
	if request == nil {
		request = NewGetTraceInterfacesRequest()
	}
	response = NewGetTraceInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlarmReceiversRequest() (request *DeleteAlarmReceiversRequest) {
	request = &DeleteAlarmReceiversRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteAlarmReceivers")
	return
}

func NewDeleteAlarmReceiversResponse() (response *DeleteAlarmReceiversResponse) {
	response = &DeleteAlarmReceiversResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除人员
func (c *Client) DeleteAlarmReceivers(request *DeleteAlarmReceiversRequest) (response *DeleteAlarmReceiversResponse, err error) {
	if request == nil {
		request = NewDeleteAlarmReceiversRequest()
	}
	response = NewDeleteAlarmReceiversResponse()
	err = c.Send(request, response)
	return
}

func NewPreviewLicenseRequest() (request *PreviewLicenseRequest) {
	request = &PreviewLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "PreviewLicense")
	return
}

func NewPreviewLicenseResponse() (response *PreviewLicenseResponse) {
	response = &PreviewLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预览证书
func (c *Client) PreviewLicense(request *PreviewLicenseRequest) (response *PreviewLicenseResponse, err error) {
	if request == nil {
		request = NewPreviewLicenseRequest()
	}
	response = NewPreviewLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiVersionsRequest() (request *DescribeApiVersionsRequest) {
	request = &DescribeApiVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiVersions")
	return
}

func NewDescribeApiVersionsResponse() (response *DescribeApiVersionsResponse) {
	response = &DescribeApiVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API 版本
func (c *Client) DescribeApiVersions(request *DescribeApiVersionsRequest) (response *DescribeApiVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeApiVersionsRequest()
	}
	response = NewDescribeApiVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigRequest() (request *DescribeConfigRequest) {
	request = &DescribeConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfig")
	return
}

func NewDescribeConfigResponse() (response *DescribeConfigResponse) {
	response = &DescribeConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配置
func (c *Client) DescribeConfig(request *DescribeConfigRequest) (response *DescribeConfigResponse, err error) {
	if request == nil {
		request = NewDescribeConfigRequest()
	}
	response = NewDescribeConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWildCardGatewayApiRequest() (request *CreateWildCardGatewayApiRequest) {
	request = &CreateWildCardGatewayApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateWildCardGatewayApi")
	return
}

func NewCreateWildCardGatewayApiResponse() (response *CreateWildCardGatewayApiResponse) {
	response = &CreateWildCardGatewayApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建通配符API
func (c *Client) CreateWildCardGatewayApi(request *CreateWildCardGatewayApiRequest) (response *CreateWildCardGatewayApiResponse, err error) {
	if request == nil {
		request = NewCreateWildCardGatewayApiRequest()
	}
	response = NewCreateWildCardGatewayApiResponse()
	err = c.Send(request, response)
	return
}

func NewDisableLaneGroupEntranceRequest() (request *DisableLaneGroupEntranceRequest) {
	request = &DisableLaneGroupEntranceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableLaneGroupEntrance")
	return
}

func NewDisableLaneGroupEntranceResponse() (response *DisableLaneGroupEntranceResponse) {
	response = &DisableLaneGroupEntranceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消泳道部署组为入口应用
func (c *Client) DisableLaneGroupEntrance(request *DisableLaneGroupEntranceRequest) (response *DisableLaneGroupEntranceResponse, err error) {
	if request == nil {
		request = NewDisableLaneGroupEntranceRequest()
	}
	response = NewDisableLaneGroupEntranceResponse()
	err = c.Send(request, response)
	return
}

func NewListManagerJobLogRequest() (request *ListManagerJobLogRequest) {
	request = &ListManagerJobLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListManagerJobLog")
	return
}

func NewListManagerJobLogResponse() (response *ListManagerJobLogResponse) {
	response = &ListManagerJobLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端任务日志列表
func (c *Client) ListManagerJobLog(request *ListManagerJobLogRequest) (response *ListManagerJobLogResponse, err error) {
	if request == nil {
		request = NewListManagerJobLogRequest()
	}
	response = NewListManagerJobLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFlameGraphRequest() (request *CreateFlameGraphRequest) {
	request = &CreateFlameGraphRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateFlameGraph")
	return
}

func NewCreateFlameGraphResponse() (response *CreateFlameGraphResponse) {
	response = &CreateFlameGraphResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 向对应的java实例下发采集火焰图请求
func (c *Client) CreateFlameGraph(request *CreateFlameGraphRequest) (response *CreateFlameGraphResponse, err error) {
	if request == nil {
		request = NewCreateFlameGraphRequest()
	}
	response = NewCreateFlameGraphResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRepositoriesRequest() (request *DescribeRepositoriesRequest) {
	request = &DescribeRepositoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRepositories")
	return
}

func NewDescribeRepositoriesResponse() (response *DescribeRepositoriesResponse) {
	response = &DescribeRepositoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询仓库列表
func (c *Client) DescribeRepositories(request *DescribeRepositoriesRequest) (response *DescribeRepositoriesResponse, err error) {
	if request == nil {
		request = NewDescribeRepositoriesRequest()
	}
	response = NewDescribeRepositoriesResponse()
	err = c.Send(request, response)
	return
}

func NewCancelHideInstanceRequest() (request *CancelHideInstanceRequest) {
	request = &CancelHideInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CancelHideInstance")
	return
}

func NewCancelHideInstanceResponse() (response *CancelHideInstanceResponse) {
	response = &CancelHideInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消屏蔽注册实例
func (c *Client) CancelHideInstance(request *CancelHideInstanceRequest) (response *CancelHideInstanceResponse, err error) {
	if request == nil {
		request = NewCancelHideInstanceRequest()
	}
	response = NewCancelHideInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterLimitResourceRequest() (request *GetClusterLimitResourceRequest) {
	request = &GetClusterLimitResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetClusterLimitResource")
	return
}

func NewGetClusterLimitResourceResponse() (response *GetClusterLimitResourceResponse) {
	response = &GetClusterLimitResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群资源
func (c *Client) GetClusterLimitResource(request *GetClusterLimitResourceRequest) (response *GetClusterLimitResourceResponse, err error) {
	if request == nil {
		request = NewGetClusterLimitResourceRequest()
	}
	response = NewGetClusterLimitResourceResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePolicyDocumentRequest() (request *CreatePolicyDocumentRequest) {
	request = &CreatePolicyDocumentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreatePolicyDocument")
	return
}

func NewCreatePolicyDocumentResponse() (response *CreatePolicyDocumentResponse) {
	response = &CreatePolicyDocumentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建CAM策略文档
func (c *Client) CreatePolicyDocument(request *CreatePolicyDocumentRequest) (response *CreatePolicyDocumentResponse, err error) {
	if request == nil {
		request = NewCreatePolicyDocumentRequest()
	}
	response = NewCreatePolicyDocumentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileConfigSummaryRequest() (request *DescribeFileConfigSummaryRequest) {
	request = &DescribeFileConfigSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigSummary")
	return
}

func NewDescribeFileConfigSummaryResponse() (response *DescribeFileConfigSummaryResponse) {
	response = &DescribeFileConfigSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件配置汇总列表
func (c *Client) DescribeFileConfigSummary(request *DescribeFileConfigSummaryRequest) (response *DescribeFileConfigSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeFileConfigSummaryRequest()
	}
	response = NewDescribeFileConfigSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubTransactionsRequest() (request *DescribeSubTransactionsRequest) {
	request = &DescribeSubTransactionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSubTransactions")
	return
}

func NewDescribeSubTransactionsResponse() (response *DescribeSubTransactionsResponse) {
	response = &DescribeSubTransactionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询子事务列表
func (c *Client) DescribeSubTransactions(request *DescribeSubTransactionsRequest) (response *DescribeSubTransactionsResponse, err error) {
	if request == nil {
		request = NewDescribeSubTransactionsRequest()
	}
	response = NewDescribeSubTransactionsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteCluster")
	return
}

func NewDeleteClusterResponse() (response *DeleteClusterResponse) {
	response = &DeleteClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群
func (c *Client) DeleteCluster(request *DeleteClusterRequest) (response *DeleteClusterResponse, err error) {
	if request == nil {
		request = NewDeleteClusterRequest()
	}
	response = NewDeleteClusterResponse()
	err = c.Send(request, response)
	return
}

func NewUninstallAgentScriptsRequest() (request *UninstallAgentScriptsRequest) {
	request = &UninstallAgentScriptsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UninstallAgentScripts")
	return
}

func NewUninstallAgentScriptsResponse() (response *UninstallAgentScriptsResponse) {
	response = &UninstallAgentScriptsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 虚拟机集群导入云主机，移除机器卸载Agent脚本
func (c *Client) UninstallAgentScripts(request *UninstallAgentScriptsRequest) (response *UninstallAgentScriptsResponse, err error) {
	if request == nil {
		request = NewUninstallAgentScriptsRequest()
	}
	response = NewUninstallAgentScriptsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespaceAlarmRequest() (request *DescribeNamespaceAlarmRequest) {
	request = &DescribeNamespaceAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaceAlarm")
	return
}

func NewDescribeNamespaceAlarmResponse() (response *DescribeNamespaceAlarmResponse) {
	response = &DescribeNamespaceAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询依赖拓扑命名空间监控告警
func (c *Client) DescribeNamespaceAlarm(request *DescribeNamespaceAlarmRequest) (response *DescribeNamespaceAlarmResponse, err error) {
	if request == nil {
		request = NewDescribeNamespaceAlarmRequest()
	}
	response = NewDescribeNamespaceAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiAccessRequest() (request *DeleteApiAccessRequest) {
	request = &DeleteApiAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteApiAccess")
	return
}

func NewDeleteApiAccessResponse() (response *DeleteApiAccessResponse) {
	response = &DeleteApiAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭Serverless应用外网访问
func (c *Client) DeleteApiAccess(request *DeleteApiAccessRequest) (response *DeleteApiAccessResponse, err error) {
	if request == nil {
		request = NewDeleteApiAccessRequest()
	}
	response = NewDeleteApiAccessResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteContainerGroupsRequest() (request *DeleteContainerGroupsRequest) {
	request = &DeleteContainerGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteContainerGroups")
	return
}

func NewDeleteContainerGroupsResponse() (response *DeleteContainerGroupsResponse) {
	response = &DeleteContainerGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除容器部署组
func (c *Client) DeleteContainerGroups(request *DeleteContainerGroupsRequest) (response *DeleteContainerGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteContainerGroupsRequest()
	}
	response = NewDeleteContainerGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerVolumeOptionsRequest() (request *DescribeContainerVolumeOptionsRequest) {
	request = &DescribeContainerVolumeOptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerVolumeOptions")
	return
}

func NewDescribeContainerVolumeOptionsResponse() (response *DescribeContainerVolumeOptionsResponse) {
	response = &DescribeContainerVolumeOptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取k8s configMap、secret、pvc配置信息
func (c *Client) DescribeContainerVolumeOptions(request *DescribeContainerVolumeOptionsRequest) (response *DescribeContainerVolumeOptionsResponse, err error) {
	if request == nil {
		request = NewDescribeContainerVolumeOptionsRequest()
	}
	response = NewDescribeContainerVolumeOptionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroservicesRequest() (request *DescribeMicroservicesRequest) {
	request = &DescribeMicroservicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservices")
	return
}

func NewDescribeMicroservicesResponse() (response *DescribeMicroservicesResponse) {
	response = &DescribeMicroservicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微服务列表
func (c *Client) DescribeMicroservices(request *DescribeMicroservicesRequest) (response *DescribeMicroservicesResponse, err error) {
	if request == nil {
		request = NewDescribeMicroservicesRequest()
	}
	response = NewDescribeMicroservicesResponse()
	err = c.Send(request, response)
	return
}

func NewCleanAndDisableUnitRouteRequest() (request *CleanAndDisableUnitRouteRequest) {
	request = &CleanAndDisableUnitRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CleanAndDisableUnitRoute")
	return
}

func NewCleanAndDisableUnitRouteResponse() (response *CleanAndDisableUnitRouteResponse) {
	response = &CleanAndDisableUnitRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清理单元化数据并关闭
func (c *Client) CleanAndDisableUnitRoute(request *CleanAndDisableUnitRouteRequest) (response *CleanAndDisableUnitRouteResponse, err error) {
	if request == nil {
		request = NewCleanAndDisableUnitRouteRequest()
	}
	response = NewCleanAndDisableUnitRouteResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBuildTaskRequest() (request *CreateBuildTaskRequest) {
	request = &CreateBuildTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateBuildTask")
	return
}

func NewCreateBuildTaskResponse() (response *CreateBuildTaskResponse) {
	response = &CreateBuildTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建构建任务
func (c *Client) CreateBuildTask(request *CreateBuildTaskRequest) (response *CreateBuildTaskResponse, err error) {
	if request == nil {
		request = NewCreateBuildTaskRequest()
	}
	response = NewCreateBuildTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePodRequest() (request *DeletePodRequest) {
	request = &DeletePodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeletePod")
	return
}

func NewDeletePodResponse() (response *DeletePodResponse) {
	response = &DeletePodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除容器服务实例
func (c *Client) DeletePod(request *DeletePodRequest) (response *DeletePodResponse, err error) {
	if request == nil {
		request = NewDeletePodRequest()
	}
	response = NewDeletePodResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationsRequest() (request *DescribeApplicationsRequest) {
	request = &DescribeApplicationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplications")
	return
}

func NewDescribeApplicationsResponse() (response *DescribeApplicationsResponse) {
	response = &DescribeApplicationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用列表
func (c *Client) DescribeApplications(request *DescribeApplicationsRequest) (response *DescribeApplicationsResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationsRequest()
	}
	response = NewDescribeApplicationsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServerlessGroupRequest() (request *DescribeServerlessGroupRequest) {
	request = &DescribeServerlessGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeServerlessGroup")
	return
}

func NewDescribeServerlessGroupResponse() (response *DescribeServerlessGroupResponse) {
	response = &DescribeServerlessGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Serverless部署组明细
func (c *Client) DescribeServerlessGroup(request *DescribeServerlessGroupRequest) (response *DescribeServerlessGroupResponse, err error) {
	if request == nil {
		request = NewDescribeServerlessGroupRequest()
	}
	response = NewDescribeServerlessGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupUseDetailRequest() (request *DescribeGroupUseDetailRequest) {
	request = &DescribeGroupUseDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupUseDetail")
	return
}

func NewDescribeGroupUseDetailResponse() (response *DescribeGroupUseDetailResponse) {
	response = &DescribeGroupUseDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关分组监控明细数据
func (c *Client) DescribeGroupUseDetail(request *DescribeGroupUseDetailRequest) (response *DescribeGroupUseDetailResponse, err error) {
	if request == nil {
		request = NewDescribeGroupUseDetailRequest()
	}
	response = NewDescribeGroupUseDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupSecretRequest() (request *DeleteGroupSecretRequest) {
	request = &DeleteGroupSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteGroupSecret")
	return
}

func NewDeleteGroupSecretResponse() (response *DeleteGroupSecretResponse) {
	response = &DeleteGroupSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除API分组的秘钥
func (c *Client) DeleteGroupSecret(request *DeleteGroupSecretRequest) (response *DeleteGroupSecretResponse, err error) {
	if request == nil {
		request = NewDeleteGroupSecretRequest()
	}
	response = NewDeleteGroupSecretResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGlobalGroupContextUniqueRestrictRequest() (request *UpdateGlobalGroupContextUniqueRestrictRequest) {
	request = &UpdateGlobalGroupContextUniqueRestrictRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGlobalGroupContextUniqueRestrict")
	return
}

func NewUpdateGlobalGroupContextUniqueRestrictResponse() (response *UpdateGlobalGroupContextUniqueRestrictResponse) {
	response = &UpdateGlobalGroupContextUniqueRestrictResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新所有微服务网关下分组context是否唯一的全局限制（更新以后，对所有微服务网关的分组context都生效）
func (c *Client) UpdateGlobalGroupContextUniqueRestrict(request *UpdateGlobalGroupContextUniqueRestrictRequest) (response *UpdateGlobalGroupContextUniqueRestrictResponse, err error) {
	if request == nil {
		request = NewUpdateGlobalGroupContextUniqueRestrictRequest()
	}
	response = NewUpdateGlobalGroupContextUniqueRestrictResponse()
	err = c.Send(request, response)
	return
}

func NewImageDeleteTagRequest() (request *ImageDeleteTagRequest) {
	request = &ImageDeleteTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ImageDeleteTag")
	return
}

func NewImageDeleteTagResponse() (response *ImageDeleteTagResponse) {
	response = &ImageDeleteTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像版本
func (c *Client) ImageDeleteTag(request *ImageDeleteTagRequest) (response *ImageDeleteTagResponse, err error) {
	if request == nil {
		request = NewImageDeleteTagRequest()
	}
	response = NewImageDeleteTagResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLaneRequest() (request *CreateLaneRequest) {
	request = &CreateLaneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateLane")
	return
}

func NewCreateLaneResponse() (response *CreateLaneResponse) {
	response = &CreateLaneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建泳道
func (c *Client) CreateLane(request *CreateLaneRequest) (response *CreateLaneResponse, err error) {
	if request == nil {
		request = NewCreateLaneRequest()
	}
	response = NewCreateLaneResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceUsageConfigRequest() (request *ModifyResourceUsageConfigRequest) {
	request = &ModifyResourceUsageConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyResourceUsageConfig")
	return
}

func NewModifyResourceUsageConfigResponse() (response *ModifyResourceUsageConfigResponse) {
	response = &ModifyResourceUsageConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改概览页配置
func (c *Client) ModifyResourceUsageConfig(request *ModifyResourceUsageConfigRequest) (response *ModifyResourceUsageConfigResponse, err error) {
	if request == nil {
		request = NewModifyResourceUsageConfigRequest()
	}
	response = NewModifyResourceUsageConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSaveDeployModuleParamsRequest() (request *SaveDeployModuleParamsRequest) {
	request = &SaveDeployModuleParamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SaveDeployModuleParams")
	return
}

func NewSaveDeployModuleParamsResponse() (response *SaveDeployModuleParamsResponse) {
	response = &SaveDeployModuleParamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存参数配置
func (c *Client) SaveDeployModuleParams(request *SaveDeployModuleParamsRequest) (response *SaveDeployModuleParamsResponse, err error) {
	if request == nil {
		request = NewSaveDeployModuleParamsRequest()
	}
	response = NewSaveDeployModuleParamsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayJwtPluginRequest() (request *UpdateGatewayJwtPluginRequest) {
	request = &UpdateGatewayJwtPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayJwtPlugin")
	return
}

func NewUpdateGatewayJwtPluginResponse() (response *UpdateGatewayJwtPluginResponse) {
	response = &UpdateGatewayJwtPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关Jwt插件
func (c *Client) UpdateGatewayJwtPlugin(request *UpdateGatewayJwtPluginRequest) (response *UpdateGatewayJwtPluginResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayJwtPluginRequest()
	}
	response = NewUpdateGatewayJwtPluginResponse()
	err = c.Send(request, response)
	return
}

func NewValidateDeletePublicConfigRequest() (request *ValidateDeletePublicConfigRequest) {
	request = &ValidateDeletePublicConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ValidateDeletePublicConfig")
	return
}

func NewValidateDeletePublicConfigResponse() (response *ValidateDeletePublicConfigResponse) {
	response = &ValidateDeletePublicConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验公共配置项可删除
func (c *Client) ValidateDeletePublicConfig(request *ValidateDeletePublicConfigRequest) (response *ValidateDeletePublicConfigResponse, err error) {
	if request == nil {
		request = NewValidateDeletePublicConfigRequest()
	}
	response = NewValidateDeletePublicConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskExecuteRecordsRequest() (request *DescribeTaskExecuteRecordsRequest) {
	request = &DescribeTaskExecuteRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskExecuteRecords")
	return
}

func NewDescribeTaskExecuteRecordsResponse() (response *DescribeTaskExecuteRecordsResponse) {
	response = &DescribeTaskExecuteRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 翻页查询任务执行列表
func (c *Client) DescribeTaskExecuteRecords(request *DescribeTaskExecuteRecordsRequest) (response *DescribeTaskExecuteRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskExecuteRecordsRequest()
	}
	response = NewDescribeTaskExecuteRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewExpandGroupRequest() (request *ExpandGroupRequest) {
	request = &ExpandGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ExpandGroup")
	return
}

func NewExpandGroupResponse() (response *ExpandGroupResponse) {
	response = &ExpandGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 虚拟机部署组添加实例
func (c *Client) ExpandGroup(request *ExpandGroupRequest) (response *ExpandGroupResponse, err error) {
	if request == nil {
		request = NewExpandGroupRequest()
	}
	response = NewExpandGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateLogCapacityRequest() (request *UpdateLogCapacityRequest) {
	request = &UpdateLogCapacityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateLogCapacity")
	return
}

func NewUpdateLogCapacityResponse() (response *UpdateLogCapacityResponse) {
	response = &UpdateLogCapacityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新日志使用量
func (c *Client) UpdateLogCapacity(request *UpdateLogCapacityRequest) (response *UpdateLogCapacityResponse, err error) {
	if request == nil {
		request = NewUpdateLogCapacityRequest()
	}
	response = NewUpdateLogCapacityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSimpleClustersRequest() (request *DescribeSimpleClustersRequest) {
	request = &DescribeSimpleClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleClusters")
	return
}

func NewDescribeSimpleClustersResponse() (response *DescribeSimpleClustersResponse) {
	response = &DescribeSimpleClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询简单集群列表
func (c *Client) DescribeSimpleClusters(request *DescribeSimpleClustersRequest) (response *DescribeSimpleClustersResponse, err error) {
	if request == nil {
		request = NewDescribeSimpleClustersRequest()
	}
	response = NewDescribeSimpleClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleaseTemplateRequest() (request *DescribeReleaseTemplateRequest) {
	request = &DescribeReleaseTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleaseTemplate")
	return
}

func NewDescribeReleaseTemplateResponse() (response *DescribeReleaseTemplateResponse) {
	response = &DescribeReleaseTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布单模板
func (c *Client) DescribeReleaseTemplate(request *DescribeReleaseTemplateRequest) (response *DescribeReleaseTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeReleaseTemplateRequest()
	}
	response = NewDescribeReleaseTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSimpleGroupsRequest() (request *DescribeSimpleGroupsRequest) {
	request = &DescribeSimpleGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleGroups")
	return
}

func NewDescribeSimpleGroupsResponse() (response *DescribeSimpleGroupsResponse) {
	response = &DescribeSimpleGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询简单部署组列表
func (c *Client) DescribeSimpleGroups(request *DescribeSimpleGroupsRequest) (response *DescribeSimpleGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeSimpleGroupsRequest()
	}
	response = NewDescribeSimpleGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMicroserviceRequest() (request *ModifyMicroserviceRequest) {
	request = &ModifyMicroserviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyMicroservice")
	return
}

func NewModifyMicroserviceResponse() (response *ModifyMicroserviceResponse) {
	response = &ModifyMicroserviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改微服务详情
func (c *Client) ModifyMicroservice(request *ModifyMicroserviceRequest) (response *ModifyMicroserviceResponse, err error) {
	if request == nil {
		request = NewModifyMicroserviceRequest()
	}
	response = NewModifyMicroserviceResponse()
	err = c.Send(request, response)
	return
}

func NewAnalyzeLogSchemaRequest() (request *AnalyzeLogSchemaRequest) {
	request = &AnalyzeLogSchemaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AnalyzeLogSchema")
	return
}

func NewAnalyzeLogSchemaResponse() (response *AnalyzeLogSchemaResponse) {
	response = &AnalyzeLogSchemaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 模拟解析日志
func (c *Client) AnalyzeLogSchema(request *AnalyzeLogSchemaRequest) (response *AnalyzeLogSchemaResponse, err error) {
	if request == nil {
		request = NewAnalyzeLogSchemaRequest()
	}
	response = NewAnalyzeLogSchemaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiUseDetailRequest() (request *DescribeApiUseDetailRequest) {
	request = &DescribeApiUseDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiUseDetail")
	return
}

func NewDescribeApiUseDetailResponse() (response *DescribeApiUseDetailResponse) {
	response = &DescribeApiUseDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关API监控明细数据
func (c *Client) DescribeApiUseDetail(request *DescribeApiUseDetailRequest) (response *DescribeApiUseDetailResponse, err error) {
	if request == nil {
		request = NewDescribeApiUseDetailRequest()
	}
	response = NewDescribeApiUseDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCircuitBreakerRuleRequest() (request *DescribeCircuitBreakerRuleRequest) {
	request = &DescribeCircuitBreakerRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeCircuitBreakerRule")
	return
}

func NewDescribeCircuitBreakerRuleResponse() (response *DescribeCircuitBreakerRuleResponse) {
	response = &DescribeCircuitBreakerRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单条熔断规则
func (c *Client) DescribeCircuitBreakerRule(request *DescribeCircuitBreakerRuleRequest) (response *DescribeCircuitBreakerRuleResponse, err error) {
	if request == nil {
		request = NewDescribeCircuitBreakerRuleRequest()
	}
	response = NewDescribeCircuitBreakerRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskLastStatusRequest() (request *DescribeTaskLastStatusRequest) {
	request = &DescribeTaskLastStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskLastStatus")
	return
}

func NewDescribeTaskLastStatusResponse() (response *DescribeTaskLastStatusResponse) {
	response = &DescribeTaskLastStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务最近一次执行状态
func (c *Client) DescribeTaskLastStatus(request *DescribeTaskLastStatusRequest) (response *DescribeTaskLastStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTaskLastStatusRequest()
	}
	response = NewDescribeTaskLastStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImagesRequest() (request *DescribeImagesRequest) {
	request = &DescribeImagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeImages")
	return
}

func NewDescribeImagesResponse() (response *DescribeImagesResponse) {
	response = &DescribeImagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询镜像列表
func (c *Client) DescribeImages(request *DescribeImagesRequest) (response *DescribeImagesResponse, err error) {
	if request == nil {
		request = NewDescribeImagesRequest()
	}
	response = NewDescribeImagesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyScalableRuleRequest() (request *ModifyScalableRuleRequest) {
	request = &ModifyScalableRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyScalableRule")
	return
}

func NewModifyScalableRuleResponse() (response *ModifyScalableRuleResponse) {
	response = &ModifyScalableRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改弹性扩缩容规则
func (c *Client) ModifyScalableRule(request *ModifyScalableRuleRequest) (response *ModifyScalableRuleResponse, err error) {
	if request == nil {
		request = NewModifyScalableRuleRequest()
	}
	response = NewModifyScalableRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayQQMiniProgramAppSecretRequest() (request *UpdateGatewayQQMiniProgramAppSecretRequest) {
	request = &UpdateGatewayQQMiniProgramAppSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayQQMiniProgramAppSecret")
	return
}

func NewUpdateGatewayQQMiniProgramAppSecretResponse() (response *UpdateGatewayQQMiniProgramAppSecretResponse) {
	response = &UpdateGatewayQQMiniProgramAppSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关QQ小程序应用密钥参数
func (c *Client) UpdateGatewayQQMiniProgramAppSecret(request *UpdateGatewayQQMiniProgramAppSecretRequest) (response *UpdateGatewayQQMiniProgramAppSecretResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayQQMiniProgramAppSecretRequest()
	}
	response = NewUpdateGatewayQQMiniProgramAppSecretResponse()
	err = c.Send(request, response)
	return
}

func NewChangeSecretStatusRequest() (request *ChangeSecretStatusRequest) {
	request = &ChangeSecretStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ChangeSecretStatus")
	return
}

func NewChangeSecretStatusResponse() (response *ChangeSecretStatusResponse) {
	response = &ChangeSecretStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用（禁用）秘钥
func (c *Client) ChangeSecretStatus(request *ChangeSecretStatusRequest) (response *ChangeSecretStatusResponse, err error) {
	if request == nil {
		request = NewChangeSecretStatusRequest()
	}
	response = NewChangeSecretStatusResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAuthorizationRequest() (request *CreateAuthorizationRequest) {
	request = &CreateAuthorizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateAuthorization")
	return
}

func NewCreateAuthorizationResponse() (response *CreateAuthorizationResponse) {
	response = &CreateAuthorizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增服务权限配置
func (c *Client) CreateAuthorization(request *CreateAuthorizationRequest) (response *CreateAuthorizationResponse, err error) {
	if request == nil {
		request = NewCreateAuthorizationRequest()
	}
	response = NewCreateAuthorizationResponse()
	err = c.Send(request, response)
	return
}

func NewDeployInstanceRequest() (request *DeployInstanceRequest) {
	request = &DeployInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeployInstance")
	return
}

func NewDeployInstanceResponse() (response *DeployInstanceResponse) {
	response = &DeployInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署机器
func (c *Client) DeployInstance(request *DeployInstanceRequest) (response *DeployInstanceResponse, err error) {
	if request == nil {
		request = NewDeployInstanceRequest()
	}
	response = NewDeployInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewEnableUnitRouteRequest() (request *EnableUnitRouteRequest) {
	request = &EnableUnitRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableUnitRoute")
	return
}

func NewEnableUnitRouteResponse() (response *EnableUnitRouteResponse) {
	response = &EnableUnitRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用单元化路由
func (c *Client) EnableUnitRoute(request *EnableUnitRouteRequest) (response *EnableUnitRouteResponse, err error) {
	if request == nil {
		request = NewEnableUnitRouteRequest()
	}
	response = NewEnableUnitRouteResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateTaskFlowBatchRequest() (request *TerminateTaskFlowBatchRequest) {
	request = &TerminateTaskFlowBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "TerminateTaskFlowBatch")
	return
}

func NewTerminateTaskFlowBatchResponse() (response *TerminateTaskFlowBatchResponse) {
	response = &TerminateTaskFlowBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止一个工作流批次
func (c *Client) TerminateTaskFlowBatch(request *TerminateTaskFlowBatchRequest) (response *TerminateTaskFlowBatchResponse, err error) {
	if request == nil {
		request = NewTerminateTaskFlowBatchRequest()
	}
	response = NewTerminateTaskFlowBatchResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseFileConfigRequest() (request *ReleaseFileConfigRequest) {
	request = &ReleaseFileConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleaseFileConfig")
	return
}

func NewReleaseFileConfigResponse() (response *ReleaseFileConfigResponse) {
	response = &ReleaseFileConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布文件配置
func (c *Client) ReleaseFileConfig(request *ReleaseFileConfigRequest) (response *ReleaseFileConfigResponse, err error) {
	if request == nil {
		request = NewReleaseFileConfigRequest()
	}
	response = NewReleaseFileConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayMonitorDetailsRequest() (request *DescribeGatewayMonitorDetailsRequest) {
	request = &DescribeGatewayMonitorDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayMonitorDetails")
	return
}

func NewDescribeGatewayMonitorDetailsResponse() (response *DescribeGatewayMonitorDetailsResponse) {
	response = &DescribeGatewayMonitorDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关监控信息
func (c *Client) DescribeGatewayMonitorDetails(request *DescribeGatewayMonitorDetailsRequest) (response *DescribeGatewayMonitorDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayMonitorDetailsRequest()
	}
	response = NewDescribeGatewayMonitorDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckFlowGraphValidityRequest() (request *CheckFlowGraphValidityRequest) {
	request = &CheckFlowGraphValidityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CheckFlowGraphValidity")
	return
}

func NewCheckFlowGraphValidityResponse() (response *CheckFlowGraphValidityResponse) {
	response = &CheckFlowGraphValidityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查工作流拓扑图的有效性
func (c *Client) CheckFlowGraphValidity(request *CheckFlowGraphValidityRequest) (response *CheckFlowGraphValidityResponse, err error) {
	if request == nil {
		request = NewCheckFlowGraphValidityRequest()
	}
	response = NewCheckFlowGraphValidityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillingOperationRecordsRequest() (request *DescribeBillingOperationRecordsRequest) {
	request = &DescribeBillingOperationRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeBillingOperationRecords")
	return
}

func NewDescribeBillingOperationRecordsResponse() (response *DescribeBillingOperationRecordsResponse) {
	response = &DescribeBillingOperationRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询计费运营端操作记录
func (c *Client) DescribeBillingOperationRecords(request *DescribeBillingOperationRecordsRequest) (response *DescribeBillingOperationRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeBillingOperationRecordsRequest()
	}
	response = NewDescribeBillingOperationRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespaceServiceOverviewRequest() (request *DescribeNamespaceServiceOverviewRequest) {
	request = &DescribeNamespaceServiceOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaceServiceOverview")
	return
}

func NewDescribeNamespaceServiceOverviewResponse() (response *DescribeNamespaceServiceOverviewResponse) {
	response = &DescribeNamespaceServiceOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取命名空间维度的服务概览列表
func (c *Client) DescribeNamespaceServiceOverview(request *DescribeNamespaceServiceOverviewRequest) (response *DescribeNamespaceServiceOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeNamespaceServiceOverviewRequest()
	}
	response = NewDescribeNamespaceServiceOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRoleRequest() (request *DescribeRoleRequest) {
	request = &DescribeRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRole")
	return
}

func NewDescribeRoleResponse() (response *DescribeRoleResponse) {
	response = &DescribeRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询角色
func (c *Client) DescribeRole(request *DescribeRoleRequest) (response *DescribeRoleResponse, err error) {
	if request == nil {
		request = NewDescribeRoleRequest()
	}
	response = NewDescribeRoleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMicroserviceRequest() (request *CreateMicroserviceRequest) {
	request = &CreateMicroserviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateMicroservice")
	return
}

func NewCreateMicroserviceResponse() (response *CreateMicroserviceResponse) {
	response = &CreateMicroserviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增微服务
func (c *Client) CreateMicroservice(request *CreateMicroserviceRequest) (response *CreateMicroserviceResponse, err error) {
	if request == nil {
		request = NewCreateMicroserviceRequest()
	}
	response = NewCreateMicroserviceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecordCodesRequest() (request *DescribeRecordCodesRequest) {
	request = &DescribeRecordCodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRecordCodes")
	return
}

func NewDescribeRecordCodesResponse() (response *DescribeRecordCodesResponse) {
	response = &DescribeRecordCodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询审计记录术语描述
func (c *Client) DescribeRecordCodes(request *DescribeRecordCodesRequest) (response *DescribeRecordCodesResponse, err error) {
	if request == nil {
		request = NewDescribeRecordCodesRequest()
	}
	response = NewDescribeRecordCodesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMultiClusterDeliveryConfigRequest() (request *CreateMultiClusterDeliveryConfigRequest) {
	request = &CreateMultiClusterDeliveryConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateMultiClusterDeliveryConfig")
	return
}

func NewCreateMultiClusterDeliveryConfigResponse() (response *CreateMultiClusterDeliveryConfigResponse) {
	response = &CreateMultiClusterDeliveryConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建投递kafka集群配置项
func (c *Client) CreateMultiClusterDeliveryConfig(request *CreateMultiClusterDeliveryConfigRequest) (response *CreateMultiClusterDeliveryConfigResponse, err error) {
	if request == nil {
		request = NewCreateMultiClusterDeliveryConfigRequest()
	}
	response = NewCreateMultiClusterDeliveryConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroserviceMetasRequest() (request *DeleteMicroserviceMetasRequest) {
	request = &DeleteMicroserviceMetasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteMicroserviceMetas")
	return
}

func NewDeleteMicroserviceMetasResponse() (response *DeleteMicroserviceMetasResponse) {
	response = &DeleteMicroserviceMetasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除微服务元数据
func (c *Client) DeleteMicroserviceMetas(request *DeleteMicroserviceMetasRequest) (response *DeleteMicroserviceMetasResponse, err error) {
	if request == nil {
		request = NewDeleteMicroserviceMetasRequest()
	}
	response = NewDeleteMicroserviceMetasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageTagsRequest() (request *DescribeImageTagsRequest) {
	request = &DescribeImageTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageTags")
	return
}

func NewDescribeImageTagsResponse() (response *DescribeImageTagsResponse) {
	response = &DescribeImageTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像版本列表
func (c *Client) DescribeImageTags(request *DescribeImageTagsRequest) (response *DescribeImageTagsResponse, err error) {
	if request == nil {
		request = NewDescribeImageTagsRequest()
	}
	response = NewDescribeImageTagsResponse()
	err = c.Send(request, response)
	return
}

func NewReRelateGroupToScalableRuleRequest() (request *ReRelateGroupToScalableRuleRequest) {
	request = &ReRelateGroupToScalableRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReRelateGroupToScalableRule")
	return
}

func NewReRelateGroupToScalableRuleResponse() (response *ReRelateGroupToScalableRuleResponse) {
	response = &ReRelateGroupToScalableRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新关联弹性伸缩规则到应用
func (c *Client) ReRelateGroupToScalableRule(request *ReRelateGroupToScalableRuleRequest) (response *ReRelateGroupToScalableRuleResponse, err error) {
	if request == nil {
		request = NewReRelateGroupToScalableRuleRequest()
	}
	response = NewReRelateGroupToScalableRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFilebeatConfigRequest() (request *CreateFilebeatConfigRequest) {
	request = &CreateFilebeatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateFilebeatConfig")
	return
}

func NewCreateFilebeatConfigResponse() (response *CreateFilebeatConfigResponse) {
	response = &CreateFilebeatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端配置通用配置（继承运营端配置，租户端优先级大于运营端），一个配置可以对应多个部署组
// 日志配置项，一个配置可以对应多个部署组
// 日志投递项，一个配置可以对应多个部署组
func (c *Client) CreateFilebeatConfig(request *CreateFilebeatConfigRequest) (response *CreateFilebeatConfigResponse, err error) {
	if request == nil {
		request = NewCreateFilebeatConfigRequest()
	}
	response = NewCreateFilebeatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowBatchHistoryRecordsRequest() (request *DescribeFlowBatchHistoryRecordsRequest) {
	request = &DescribeFlowBatchHistoryRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowBatchHistoryRecords")
	return
}

func NewDescribeFlowBatchHistoryRecordsResponse() (response *DescribeFlowBatchHistoryRecordsResponse) {
	response = &DescribeFlowBatchHistoryRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 翻页查询工作流批次历史
func (c *Client) DescribeFlowBatchHistoryRecords(request *DescribeFlowBatchHistoryRecordsRequest) (response *DescribeFlowBatchHistoryRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeFlowBatchHistoryRecordsRequest()
	}
	response = NewDescribeFlowBatchHistoryRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePkgsRequest() (request *DescribePkgsRequest) {
	request = &DescribePkgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePkgs")
	return
}

func NewDescribePkgsResponse() (response *DescribePkgsResponse) {
	response = &DescribePkgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) DescribePkgs(request *DescribePkgsRequest) (response *DescribePkgsResponse, err error) {
	if request == nil {
		request = NewDescribePkgsRequest()
	}
	response = NewDescribePkgsResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseSidecarFilterRequest() (request *ReleaseSidecarFilterRequest) {
	request = &ReleaseSidecarFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleaseSidecarFilter")
	return
}

func NewReleaseSidecarFilterResponse() (response *ReleaseSidecarFilterResponse) {
	response = &ReleaseSidecarFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布Sidecar过滤器
func (c *Client) ReleaseSidecarFilter(request *ReleaseSidecarFilterRequest) (response *ReleaseSidecarFilterResponse, err error) {
	if request == nil {
		request = NewReleaseSidecarFilterRequest()
	}
	response = NewReleaseSidecarFilterResponse()
	err = c.Send(request, response)
	return
}

func NewSearchTraceRequest() (request *SearchTraceRequest) {
	request = &SearchTraceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchTrace")
	return
}

func NewSearchTraceResponse() (response *SearchTraceResponse) {
	response = &SearchTraceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用链
func (c *Client) SearchTrace(request *SearchTraceRequest) (response *SearchTraceResponse, err error) {
	if request == nil {
		request = NewSearchTraceRequest()
	}
	response = NewSearchTraceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeZonesRequest() (request *DescribeZonesRequest) {
	request = &DescribeZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeZones")
	return
}

func NewDescribeZonesResponse() (response *DescribeZonesResponse) {
	response = &DescribeZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务可用区
func (c *Client) DescribeZones(request *DescribeZonesRequest) (response *DescribeZonesResponse, err error) {
	if request == nil {
		request = NewDescribeZonesRequest()
	}
	response = NewDescribeZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRolesRequest() (request *DescribeRolesRequest) {
	request = &DescribeRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRoles")
	return
}

func NewDescribeRolesResponse() (response *DescribeRolesResponse) {
	response = &DescribeRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询角色列表
func (c *Client) DescribeRoles(request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
	if request == nil {
		request = NewDescribeRolesRequest()
	}
	response = NewDescribeRolesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePluginInstancesRequest() (request *DescribePluginInstancesRequest) {
	request = &DescribePluginInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePluginInstances")
	return
}

func NewDescribePluginInstancesResponse() (response *DescribePluginInstancesResponse) {
	response = &DescribePluginInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询网关分组/API绑定（或未绑定）的插件列表
func (c *Client) DescribePluginInstances(request *DescribePluginInstancesRequest) (response *DescribePluginInstancesResponse, err error) {
	if request == nil {
		request = NewDescribePluginInstancesRequest()
	}
	response = NewDescribePluginInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewSynchronizeContainerClusterRequest() (request *SynchronizeContainerClusterRequest) {
	request = &SynchronizeContainerClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SynchronizeContainerCluster")
	return
}

func NewSynchronizeContainerClusterResponse() (response *SynchronizeContainerClusterResponse) {
	response = &SynchronizeContainerClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 同步容器平台集群到tsf
func (c *Client) SynchronizeContainerCluster(request *SynchronizeContainerClusterRequest) (response *SynchronizeContainerClusterResponse, err error) {
	if request == nil {
		request = NewSynchronizeContainerClusterRequest()
	}
	response = NewSynchronizeContainerClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTemplateRequest() (request *DeleteTemplateRequest) {
	request = &DeleteTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteTemplate")
	return
}

func NewDeleteTemplateResponse() (response *DeleteTemplateResponse) {
	response = &DeleteTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除模板
func (c *Client) DeleteTemplate(request *DeleteTemplateRequest) (response *DeleteTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteTemplateRequest()
	}
	response = NewDeleteTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewListContainGroupRequest() (request *ListContainGroupRequest) {
	request = &ListContainGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListContainGroup")
	return
}

func NewListContainGroupResponse() (response *ListContainGroupResponse) {
	response = &ListContainGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器部署组列表
func (c *Client) ListContainGroup(request *ListContainGroupRequest) (response *ListContainGroupResponse, err error) {
	if request == nil {
		request = NewListContainGroupRequest()
	}
	response = NewListContainGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJvmMonitorRequest() (request *DescribeJvmMonitorRequest) {
	request = &DescribeJvmMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeJvmMonitor")
	return
}

func NewDescribeJvmMonitorResponse() (response *DescribeJvmMonitorResponse) {
	response = &DescribeJvmMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询java实例jvm监控数据,返回数据可选
func (c *Client) DescribeJvmMonitor(request *DescribeJvmMonitorRequest) (response *DescribeJvmMonitorResponse, err error) {
	if request == nil {
		request = NewDescribeJvmMonitorRequest()
	}
	response = NewDescribeJvmMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayQQMiniProgramLoginPluginRequest() (request *DescribeGatewayQQMiniProgramLoginPluginRequest) {
	request = &DescribeGatewayQQMiniProgramLoginPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayQQMiniProgramLoginPlugin")
	return
}

func NewDescribeGatewayQQMiniProgramLoginPluginResponse() (response *DescribeGatewayQQMiniProgramLoginPluginResponse) {
	response = &DescribeGatewayQQMiniProgramLoginPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询qq小程序登录插件详情信息
func (c *Client) DescribeGatewayQQMiniProgramLoginPlugin(request *DescribeGatewayQQMiniProgramLoginPluginRequest) (response *DescribeGatewayQQMiniProgramLoginPluginResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayQQMiniProgramLoginPluginRequest()
	}
	response = NewDescribeGatewayQQMiniProgramLoginPluginResponse()
	err = c.Send(request, response)
	return
}

func NewCheckExecuteStatusRequest() (request *CheckExecuteStatusRequest) {
	request = &CheckExecuteStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CheckExecuteStatus")
	return
}

func NewCheckExecuteStatusResponse() (response *CheckExecuteStatusResponse) {
	response = &CheckExecuteStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查运行状态
func (c *Client) CheckExecuteStatus(request *CheckExecuteStatusRequest) (response *CheckExecuteStatusResponse, err error) {
	if request == nil {
		request = NewCheckExecuteStatusRequest()
	}
	response = NewCheckExecuteStatusResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGroupRequest() (request *ModifyGroupRequest) {
	request = &ModifyGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyGroup")
	return
}

func NewModifyGroupResponse() (response *ModifyGroupResponse) {
	response = &ModifyGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新部署组信息
func (c *Client) ModifyGroup(request *ModifyGroupRequest) (response *ModifyGroupResponse, err error) {
	if request == nil {
		request = NewModifyGroupRequest()
	}
	response = NewModifyGroupResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSurroundBusinessLogRequest() (request *SearchSurroundBusinessLogRequest) {
	request = &SearchSurroundBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchSurroundBusinessLog")
	return
}

func NewSearchSurroundBusinessLogResponse() (response *SearchSurroundBusinessLogResponse) {
	response = &SearchSurroundBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务日志上下文检索
func (c *Client) SearchSurroundBusinessLog(request *SearchSurroundBusinessLogRequest) (response *SearchSurroundBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchSurroundBusinessLogRequest()
	}
	response = NewSearchSurroundBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowBatchGraphRequest() (request *DescribeFlowBatchGraphRequest) {
	request = &DescribeFlowBatchGraphRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowBatchGraph")
	return
}

func NewDescribeFlowBatchGraphResponse() (response *DescribeFlowBatchGraphResponse) {
	response = &DescribeFlowBatchGraphResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作流执行批次的图信息
func (c *Client) DescribeFlowBatchGraph(request *DescribeFlowBatchGraphRequest) (response *DescribeFlowBatchGraphResponse, err error) {
	if request == nil {
		request = NewDescribeFlowBatchGraphRequest()
	}
	response = NewDescribeFlowBatchGraphResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnitRulesV2Request() (request *DescribeUnitRulesV2Request) {
	request = &DescribeUnitRulesV2Request{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRulesV2")
	return
}

func NewDescribeUnitRulesV2Response() (response *DescribeUnitRulesV2Response) {
	response = &DescribeUnitRulesV2Response{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单元化规则列表
func (c *Client) DescribeUnitRulesV2(request *DescribeUnitRulesV2Request) (response *DescribeUnitRulesV2Response, err error) {
	if request == nil {
		request = NewDescribeUnitRulesV2Request()
	}
	response = NewDescribeUnitRulesV2Response()
	err = c.Send(request, response)
	return
}

func NewModifyLaneRuleRequest() (request *ModifyLaneRuleRequest) {
	request = &ModifyLaneRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyLaneRule")
	return
}

func NewModifyLaneRuleResponse() (response *ModifyLaneRuleResponse) {
	response = &ModifyLaneRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新泳道规则
func (c *Client) ModifyLaneRule(request *ModifyLaneRuleRequest) (response *ModifyLaneRuleResponse, err error) {
	if request == nil {
		request = NewModifyLaneRuleRequest()
	}
	response = NewModifyLaneRuleResponse()
	err = c.Send(request, response)
	return
}

func NewRunJvmDeadLockCheckRequest() (request *RunJvmDeadLockCheckRequest) {
	request = &RunJvmDeadLockCheckRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RunJvmDeadLockCheck")
	return
}

func NewRunJvmDeadLockCheckResponse() (response *RunJvmDeadLockCheckResponse) {
	response = &RunJvmDeadLockCheckResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对用户java进程进行死锁检测
func (c *Client) RunJvmDeadLockCheck(request *RunJvmDeadLockCheckRequest) (response *RunJvmDeadLockCheckResponse, err error) {
	if request == nil {
		request = NewRunJvmDeadLockCheckRequest()
	}
	response = NewRunJvmDeadLockCheckResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthNamespacesRequest() (request *DescribeAuthNamespacesRequest) {
	request = &DescribeAuthNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthNamespaces")
	return
}

func NewDescribeAuthNamespacesResponse() (response *DescribeAuthNamespacesResponse) {
	response = &DescribeAuthNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询租户下可用于授权的命名空间
func (c *Client) DescribeAuthNamespaces(request *DescribeAuthNamespacesRequest) (response *DescribeAuthNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeAuthNamespacesRequest()
	}
	response = NewDescribeAuthNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCreateGatewayApiStatusRequest() (request *DescribeCreateGatewayApiStatusRequest) {
	request = &DescribeCreateGatewayApiStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeCreateGatewayApiStatus")
	return
}

func NewDescribeCreateGatewayApiStatusResponse() (response *DescribeCreateGatewayApiStatusResponse) {
	response = &DescribeCreateGatewayApiStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询一键导入API分组任务的状态
func (c *Client) DescribeCreateGatewayApiStatus(request *DescribeCreateGatewayApiStatusRequest) (response *DescribeCreateGatewayApiStatusResponse, err error) {
	if request == nil {
		request = NewDescribeCreateGatewayApiStatusRequest()
	}
	response = NewDescribeCreateGatewayApiStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupBindedGatewaysRequest() (request *DescribeGroupBindedGatewaysRequest) {
	request = &DescribeGroupBindedGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupBindedGateways")
	return
}

func NewDescribeGroupBindedGatewaysResponse() (response *DescribeGroupBindedGatewaysResponse) {
	response = &DescribeGroupBindedGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个API分组已绑定的网关部署组信息列表
func (c *Client) DescribeGroupBindedGateways(request *DescribeGroupBindedGatewaysRequest) (response *DescribeGroupBindedGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeGroupBindedGatewaysRequest()
	}
	response = NewDescribeGroupBindedGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewCheckUploadInfoRequest() (request *CheckUploadInfoRequest) {
	request = &CheckUploadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CheckUploadInfo")
	return
}

func NewCheckUploadInfoResponse() (response *CheckUploadInfoResponse) {
	response = &CheckUploadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF会将软件包上传到腾讯云对象存储（COS）。调用此接口检查上传程序包信息，检查程序包是否完整，是否包含启动脚本。
func (c *Client) CheckUploadInfo(request *CheckUploadInfoRequest) (response *CheckUploadInfoResponse, err error) {
	if request == nil {
		request = NewCheckUploadInfoRequest()
	}
	response = NewCheckUploadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePathRewritesRequest() (request *DeletePathRewritesRequest) {
	request = &DeletePathRewritesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeletePathRewrites")
	return
}

func NewDeletePathRewritesResponse() (response *DeletePathRewritesResponse) {
	response = &DeletePathRewritesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除路径重写
func (c *Client) DeletePathRewrites(request *DeletePathRewritesRequest) (response *DeletePathRewritesResponse, err error) {
	if request == nil {
		request = NewDeletePathRewritesRequest()
	}
	response = NewDeletePathRewritesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRouteRequest() (request *ModifyRouteRequest) {
	request = &ModifyRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyRoute")
	return
}

func NewModifyRouteResponse() (response *ModifyRouteResponse) {
	response = &ModifyRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新路由详情
func (c *Client) ModifyRoute(request *ModifyRouteRequest) (response *ModifyRouteResponse, err error) {
	if request == nil {
		request = NewModifyRouteRequest()
	}
	response = NewModifyRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateBusinessLogConfigRequest() (request *DisassociateBusinessLogConfigRequest) {
	request = &DisassociateBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisassociateBusinessLogConfig")
	return
}

func NewDisassociateBusinessLogConfigResponse() (response *DisassociateBusinessLogConfigResponse) {
	response = &DisassociateBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消关联业务日志配置项和应用
func (c *Client) DisassociateBusinessLogConfig(request *DisassociateBusinessLogConfigRequest) (response *DisassociateBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewDisassociateBusinessLogConfigRequest()
	}
	response = NewDisassociateBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInterfaceStatisticRequest() (request *DescribeInterfaceStatisticRequest) {
	request = &DescribeInterfaceStatisticRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInterfaceStatistic")
	return
}

func NewDescribeInterfaceStatisticResponse() (response *DescribeInterfaceStatisticResponse) {
	response = &DescribeInterfaceStatisticResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按照服务,统计每个接口的请求量、响应时间、耗时
func (c *Client) DescribeInterfaceStatistic(request *DescribeInterfaceStatisticRequest) (response *DescribeInterfaceStatisticResponse, err error) {
	if request == nil {
		request = NewDescribeInterfaceStatisticRequest()
	}
	response = NewDescribeInterfaceStatisticResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteReleasesRequest() (request *DeleteReleasesRequest) {
	request = &DeleteReleasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteReleases")
	return
}

func NewDeleteReleasesResponse() (response *DeleteReleasesResponse) {
	response = &DeleteReleasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除发布单
func (c *Client) DeleteReleases(request *DeleteReleasesRequest) (response *DeleteReleasesResponse, err error) {
	if request == nil {
		request = NewDeleteReleasesRequest()
	}
	response = NewDeleteReleasesResponse()
	err = c.Send(request, response)
	return
}

func NewDeployServerlessGroupRequest() (request *DeployServerlessGroupRequest) {
	request = &DeployServerlessGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeployServerlessGroup")
	return
}

func NewDeployServerlessGroupResponse() (response *DeployServerlessGroupResponse) {
	response = &DeployServerlessGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署Serverless应用
func (c *Client) DeployServerlessGroup(request *DeployServerlessGroupRequest) (response *DeployServerlessGroupResponse, err error) {
	if request == nil {
		request = NewDeployServerlessGroupRequest()
	}
	response = NewDeployServerlessGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateCircuitBreakerRuleRequest() (request *UpdateCircuitBreakerRuleRequest) {
	request = &UpdateCircuitBreakerRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateCircuitBreakerRule")
	return
}

func NewUpdateCircuitBreakerRuleResponse() (response *UpdateCircuitBreakerRuleResponse) {
	response = &UpdateCircuitBreakerRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新熔断规则
func (c *Client) UpdateCircuitBreakerRule(request *UpdateCircuitBreakerRuleRequest) (response *UpdateCircuitBreakerRuleResponse, err error) {
	if request == nil {
		request = NewUpdateCircuitBreakerRuleRequest()
	}
	response = NewUpdateCircuitBreakerRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUploadPkgRequest() (request *UploadPkgRequest) {
	request = &UploadPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UploadPkg")
	return
}

func NewUploadPkgResponse() (response *UploadPkgResponse) {
	response = &UploadPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 程序包分片上传接口
func (c *Client) UploadPkg(request *UploadPkgRequest) (response *UploadPkgResponse, err error) {
	if request == nil {
		request = NewUploadPkgRequest()
	}
	response = NewUploadPkgResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayQQMiniProgramLoginPluginRequest() (request *CreateGatewayQQMiniProgramLoginPluginRequest) {
	request = &CreateGatewayQQMiniProgramLoginPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayQQMiniProgramLoginPlugin")
	return
}

func NewCreateGatewayQQMiniProgramLoginPluginResponse() (response *CreateGatewayQQMiniProgramLoginPluginResponse) {
	response = &CreateGatewayQQMiniProgramLoginPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增网关qq小程序登录插件
func (c *Client) CreateGatewayQQMiniProgramLoginPlugin(request *CreateGatewayQQMiniProgramLoginPluginRequest) (response *CreateGatewayQQMiniProgramLoginPluginResponse, err error) {
	if request == nil {
		request = NewCreateGatewayQQMiniProgramLoginPluginRequest()
	}
	response = NewCreateGatewayQQMiniProgramLoginPluginResponse()
	err = c.Send(request, response)
	return
}

func NewStartBuildTaskRequest() (request *StartBuildTaskRequest) {
	request = &StartBuildTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StartBuildTask")
	return
}

func NewStartBuildTaskResponse() (response *StartBuildTaskResponse) {
	response = &StartBuildTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动构建任务
func (c *Client) StartBuildTask(request *StartBuildTaskRequest) (response *StartBuildTaskResponse, err error) {
	if request == nil {
		request = NewStartBuildTaskRequest()
	}
	response = NewStartBuildTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiGroupsRequest() (request *DescribeApiGroupsRequest) {
	request = &DescribeApiGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroups")
	return
}

func NewDescribeApiGroupsResponse() (response *DescribeApiGroupsResponse) {
	response = &DescribeApiGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API 分组信息列表
func (c *Client) DescribeApiGroups(request *DescribeApiGroupsRequest) (response *DescribeApiGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeApiGroupsRequest()
	}
	response = NewDescribeApiGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewEnableUnitRuleRequest() (request *EnableUnitRuleRequest) {
	request = &EnableUnitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableUnitRule")
	return
}

func NewEnableUnitRuleResponse() (response *EnableUnitRuleResponse) {
	response = &EnableUnitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用单元化规则
func (c *Client) EnableUnitRule(request *EnableUnitRuleRequest) (response *EnableUnitRuleResponse, err error) {
	if request == nil {
		request = NewEnableUnitRuleRequest()
	}
	response = NewEnableUnitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewInitializeApmRequest() (request *InitializeApmRequest) {
	request = &InitializeApmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "InitializeApm")
	return
}

func NewInitializeApmResponse() (response *InitializeApmResponse) {
	response = &InitializeApmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 初始化应用性能管理功能，包括默认调用链索引创建，ES权限用户创建和授权等
func (c *Client) InitializeApm(request *InitializeApmRequest) (response *InitializeApmResponse, err error) {
	if request == nil {
		request = NewInitializeApmRequest()
	}
	response = NewInitializeApmResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProgramRequest() (request *CreateProgramRequest) {
	request = &CreateProgramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateProgram")
	return
}

func NewCreateProgramResponse() (response *CreateProgramResponse) {
	response = &CreateProgramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建数据集
func (c *Client) CreateProgram(request *CreateProgramRequest) (response *CreateProgramResponse, err error) {
	if request == nil {
		request = NewCreateProgramRequest()
	}
	response = NewCreateProgramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmPolicyRequest() (request *DescribeAlarmPolicyRequest) {
	request = &DescribeAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAlarmPolicy")
	return
}

func NewDescribeAlarmPolicyResponse() (response *DescribeAlarmPolicyResponse) {
	response = &DescribeAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询告警策略
func (c *Client) DescribeAlarmPolicy(request *DescribeAlarmPolicyRequest) (response *DescribeAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmPolicyRequest()
	}
	response = NewDescribeAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerGroupYamlRequest() (request *DescribeContainerGroupYamlRequest) {
	request = &DescribeContainerGroupYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupYaml")
	return
}

func NewDescribeContainerGroupYamlResponse() (response *DescribeContainerGroupYamlResponse) {
	response = &DescribeContainerGroupYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 TSF 容器部署组的 yaml 信息
func (c *Client) DescribeContainerGroupYaml(request *DescribeContainerGroupYamlRequest) (response *DescribeContainerGroupYamlResponse, err error) {
	if request == nil {
		request = NewDescribeContainerGroupYamlRequest()
	}
	response = NewDescribeContainerGroupYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOverviewInvocationRequest() (request *DescribeOverviewInvocationRequest) {
	request = &DescribeOverviewInvocationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeOverviewInvocation")
	return
}

func NewDescribeOverviewInvocationResponse() (response *DescribeOverviewInvocationResponse) {
	response = &DescribeOverviewInvocationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务调用监控统计概览
func (c *Client) DescribeOverviewInvocation(request *DescribeOverviewInvocationRequest) (response *DescribeOverviewInvocationResponse, err error) {
	if request == nil {
		request = NewDescribeOverviewInvocationRequest()
	}
	response = NewDescribeOverviewInvocationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeStatisticsRequest() (request *DescribeStatisticsRequest) {
	request = &DescribeStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeStatistics")
	return
}

func NewDescribeStatisticsResponse() (response *DescribeStatisticsResponse) {
	response = &DescribeStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务统计页面：接口和服务维度
func (c *Client) DescribeStatistics(request *DescribeStatisticsRequest) (response *DescribeStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeStatisticsRequest()
	}
	response = NewDescribeStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterLimitResourceRequest() (request *DescribeClusterLimitResourceRequest) {
	request = &DescribeClusterLimitResourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterLimitResource")
	return
}

func NewDescribeClusterLimitResourceResponse() (response *DescribeClusterLimitResourceResponse) {
	response = &DescribeClusterLimitResourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群剩余资源
func (c *Client) DescribeClusterLimitResource(request *DescribeClusterLimitResourceRequest) (response *DescribeClusterLimitResourceResponse, err error) {
	if request == nil {
		request = NewDescribeClusterLimitResourceRequest()
	}
	response = NewDescribeClusterLimitResourceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
	request = &DescribeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigs")
	return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
	response = &DescribeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配置项列表
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeConfigsRequest()
	}
	response = NewDescribeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProductHelpRequest() (request *ModifyProductHelpRequest) {
	request = &ModifyProductHelpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyProductHelp")
	return
}

func NewModifyProductHelpResponse() (response *ModifyProductHelpResponse) {
	response = &ModifyProductHelpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改产品帮助
func (c *Client) ModifyProductHelp(request *ModifyProductHelpRequest) (response *ModifyProductHelpResponse, err error) {
	if request == nil {
		request = NewModifyProductHelpRequest()
	}
	response = NewModifyProductHelpResponse()
	err = c.Send(request, response)
	return
}

func NewEnableLaneGroupEntranceRequest() (request *EnableLaneGroupEntranceRequest) {
	request = &EnableLaneGroupEntranceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableLaneGroupEntrance")
	return
}

func NewEnableLaneGroupEntranceResponse() (response *EnableLaneGroupEntranceResponse) {
	response = &EnableLaneGroupEntranceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置泳道部署组入口
func (c *Client) EnableLaneGroupEntrance(request *EnableLaneGroupEntranceRequest) (response *EnableLaneGroupEntranceResponse, err error) {
	if request == nil {
		request = NewEnableLaneGroupEntranceRequest()
	}
	response = NewEnableLaneGroupEntranceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayAllGroupApisRequest() (request *DescribeGatewayAllGroupApisRequest) {
	request = &DescribeGatewayAllGroupApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayAllGroupApis")
	return
}

func NewDescribeGatewayAllGroupApisResponse() (response *DescribeGatewayAllGroupApisResponse) {
	response = &DescribeGatewayAllGroupApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关所有分组下Api列表
func (c *Client) DescribeGatewayAllGroupApis(request *DescribeGatewayAllGroupApisRequest) (response *DescribeGatewayAllGroupApisResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayAllGroupApisRequest()
	}
	response = NewDescribeGatewayAllGroupApisResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApiAccessRequest() (request *CreateApiAccessRequest) {
	request = &CreateApiAccessRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateApiAccess")
	return
}

func NewCreateApiAccessResponse() (response *CreateApiAccessResponse) {
	response = &CreateApiAccessResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启Serverless应用的外网访问
func (c *Client) CreateApiAccess(request *CreateApiAccessRequest) (response *CreateApiAccessResponse, err error) {
	if request == nil {
		request = NewCreateApiAccessRequest()
	}
	response = NewCreateApiAccessResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteServiceInstancesRequest() (request *DeleteServiceInstancesRequest) {
	request = &DeleteServiceInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteServiceInstances")
	return
}

func NewDeleteServiceInstancesResponse() (response *DeleteServiceInstancesResponse) {
	response = &DeleteServiceInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除模块实例节点
func (c *Client) DeleteServiceInstances(request *DeleteServiceInstancesRequest) (response *DeleteServiceInstancesResponse, err error) {
	if request == nil {
		request = NewDeleteServiceInstancesRequest()
	}
	response = NewDeleteServiceInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLaneRuleRequest() (request *CreateLaneRuleRequest) {
	request = &CreateLaneRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateLaneRule")
	return
}

func NewCreateLaneRuleResponse() (response *CreateLaneRuleResponse) {
	response = &CreateLaneRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建泳道规则
func (c *Client) CreateLaneRule(request *CreateLaneRuleRequest) (response *CreateLaneRuleResponse, err error) {
	if request == nil {
		request = NewCreateLaneRuleRequest()
	}
	response = NewCreateLaneRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationRequest() (request *DescribeApplicationRequest) {
	request = &DescribeApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplication")
	return
}

func NewDescribeApplicationResponse() (response *DescribeApplicationResponse) {
	response = &DescribeApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用详情
func (c *Client) DescribeApplication(request *DescribeApplicationRequest) (response *DescribeApplicationResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationRequest()
	}
	response = NewDescribeApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServerlessGroupRequest() (request *CreateServerlessGroupRequest) {
	request = &CreateServerlessGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateServerlessGroup")
	return
}

func NewCreateServerlessGroupResponse() (response *CreateServerlessGroupResponse) {
	response = &CreateServerlessGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Serverless部署组
func (c *Client) CreateServerlessGroup(request *CreateServerlessGroupRequest) (response *CreateServerlessGroupResponse, err error) {
	if request == nil {
		request = NewCreateServerlessGroupRequest()
	}
	response = NewCreateServerlessGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceStatisticsRequest() (request *DescribeServiceStatisticsRequest) {
	request = &DescribeServiceStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeServiceStatistics")
	return
}

func NewDescribeServiceStatisticsResponse() (response *DescribeServiceStatisticsResponse) {
	response = &DescribeServiceStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微服务监控统计
func (c *Client) DescribeServiceStatistics(request *DescribeServiceStatisticsRequest) (response *DescribeServiceStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeServiceStatisticsRequest()
	}
	response = NewDescribeServiceStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAuthorizationRequest() (request *DeleteAuthorizationRequest) {
	request = &DeleteAuthorizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteAuthorization")
	return
}

func NewDeleteAuthorizationResponse() (response *DeleteAuthorizationResponse) {
	response = &DeleteAuthorizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除服务权限配置
func (c *Client) DeleteAuthorization(request *DeleteAuthorizationRequest) (response *DeleteAuthorizationResponse, err error) {
	if request == nil {
		request = NewDeleteAuthorizationRequest()
	}
	response = NewDeleteAuthorizationResponse()
	err = c.Send(request, response)
	return
}

func NewListCloudMicroServiceFindPagedListRequest() (request *ListCloudMicroServiceFindPagedListRequest) {
	request = &ListCloudMicroServiceFindPagedListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListCloudMicroServiceFindPagedList")
	return
}

func NewListCloudMicroServiceFindPagedListResponse() (response *ListCloudMicroServiceFindPagedListResponse) {
	response = &ListCloudMicroServiceFindPagedListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CloudMonitorMicroserviceResult
func (c *Client) ListCloudMicroServiceFindPagedList(request *ListCloudMicroServiceFindPagedListRequest) (response *ListCloudMicroServiceFindPagedListResponse, err error) {
	if request == nil {
		request = NewListCloudMicroServiceFindPagedListRequest()
	}
	response = NewListCloudMicroServiceFindPagedListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceReleasedConfigRequest() (request *DescribeInstanceReleasedConfigRequest) {
	request = &DescribeInstanceReleasedConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInstanceReleasedConfig")
	return
}

func NewDescribeInstanceReleasedConfigResponse() (response *DescribeInstanceReleasedConfigResponse) {
	response = &DescribeInstanceReleasedConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看实例生效配置
func (c *Client) DescribeInstanceReleasedConfig(request *DescribeInstanceReleasedConfigRequest) (response *DescribeInstanceReleasedConfigResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceReleasedConfigRequest()
	}
	response = NewDescribeInstanceReleasedConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLaneGroupsRequest() (request *DescribeLaneGroupsRequest) {
	request = &DescribeLaneGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLaneGroups")
	return
}

func NewDescribeLaneGroupsResponse() (response *DescribeLaneGroupsResponse) {
	response = &DescribeLaneGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询泳道部署组列表
func (c *Client) DescribeLaneGroups(request *DescribeLaneGroupsRequest) (response *DescribeLaneGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeLaneGroupsRequest()
	}
	response = NewDescribeLaneGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCommonGroupAttributeRequest() (request *DescribeCommonGroupAttributeRequest) {
	request = &DescribeCommonGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeCommonGroupAttribute")
	return
}

func NewDescribeCommonGroupAttributeResponse() (response *DescribeCommonGroupAttributeResponse) {
	response = &DescribeCommonGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询通用部署组额外属性
func (c *Client) DescribeCommonGroupAttribute(request *DescribeCommonGroupAttributeRequest) (response *DescribeCommonGroupAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeCommonGroupAttributeRequest()
	}
	response = NewDescribeCommonGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupEnvRequest() (request *DescribeGroupEnvRequest) {
	request = &DescribeGroupEnvRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupEnv")
	return
}

func NewDescribeGroupEnvResponse() (response *DescribeGroupEnvResponse) {
	response = &DescribeGroupEnvResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取虚拟机部署组环境变量
func (c *Client) DescribeGroupEnv(request *DescribeGroupEnvRequest) (response *DescribeGroupEnvResponse, err error) {
	if request == nil {
		request = NewDescribeGroupEnvRequest()
	}
	response = NewDescribeGroupEnvResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
	request = &DescribeTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTemplates")
	return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
	response = &DescribeTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取模板列表
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeTemplatesRequest()
	}
	response = NewDescribeTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReleaseTemplateRequest() (request *ModifyReleaseTemplateRequest) {
	request = &ModifyReleaseTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyReleaseTemplate")
	return
}

func NewModifyReleaseTemplateResponse() (response *ModifyReleaseTemplateResponse) {
	response = &ModifyReleaseTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改发布单模板
func (c *Client) ModifyReleaseTemplate(request *ModifyReleaseTemplateRequest) (response *ModifyReleaseTemplateResponse, err error) {
	if request == nil {
		request = NewModifyReleaseTemplateRequest()
	}
	response = NewModifyReleaseTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApiTimeoutsRequest() (request *UpdateApiTimeoutsRequest) {
	request = &UpdateApiTimeoutsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiTimeouts")
	return
}

func NewUpdateApiTimeoutsResponse() (response *UpdateApiTimeoutsResponse) {
	response = &UpdateApiTimeoutsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量更新API超时
func (c *Client) UpdateApiTimeouts(request *UpdateApiTimeoutsRequest) (response *UpdateApiTimeoutsResponse, err error) {
	if request == nil {
		request = NewUpdateApiTimeoutsRequest()
	}
	response = NewUpdateApiTimeoutsResponse()
	err = c.Send(request, response)
	return
}

func NewOperateManagerJobRequest() (request *OperateManagerJobRequest) {
	request = &OperateManagerJobRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "OperateManagerJob")
	return
}

func NewOperateManagerJobResponse() (response *OperateManagerJobResponse) {
	response = &OperateManagerJobResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 操作运营端任务
func (c *Client) OperateManagerJob(request *OperateManagerJobRequest) (response *OperateManagerJobResponse, err error) {
	if request == nil {
		request = NewOperateManagerJobRequest()
	}
	response = NewOperateManagerJobResponse()
	err = c.Send(request, response)
	return
}

func NewAddClusterInstancesRequest() (request *AddClusterInstancesRequest) {
	request = &AddClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AddClusterInstances")
	return
}

func NewAddClusterInstancesResponse() (response *AddClusterInstancesResponse) {
	response = &AddClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加云主机节点至TSF集群
func (c *Client) AddClusterInstances(request *AddClusterInstancesRequest) (response *AddClusterInstancesResponse, err error) {
	if request == nil {
		request = NewAddClusterInstancesRequest()
	}
	response = NewAddClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskFlowRequest() (request *DescribeTaskFlowRequest) {
	request = &DescribeTaskFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskFlow")
	return
}

func NewDescribeTaskFlowResponse() (response *DescribeTaskFlowResponse) {
	response = &DescribeTaskFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作流的基本信息详情
func (c *Client) DescribeTaskFlow(request *DescribeTaskFlowRequest) (response *DescribeTaskFlowResponse, err error) {
	if request == nil {
		request = NewDescribeTaskFlowRequest()
	}
	response = NewDescribeTaskFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnitApiUseDetailRequest() (request *DescribeUnitApiUseDetailRequest) {
	request = &DescribeUnitApiUseDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitApiUseDetail")
	return
}

func NewDescribeUnitApiUseDetailResponse() (response *DescribeUnitApiUseDetailResponse) {
	response = &DescribeUnitApiUseDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关API监控明细数据（仅单元化网关），非单元化网关使用DescribeApiUseDetail
func (c *Client) DescribeUnitApiUseDetail(request *DescribeUnitApiUseDetailRequest) (response *DescribeUnitApiUseDetailResponse, err error) {
	if request == nil {
		request = NewDescribeUnitApiUseDetailRequest()
	}
	response = NewDescribeUnitApiUseDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBuildTasksRequest() (request *DescribeBuildTasksRequest) {
	request = &DescribeBuildTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeBuildTasks")
	return
}

func NewDescribeBuildTasksResponse() (response *DescribeBuildTasksResponse) {
	response = &DescribeBuildTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询构建任务状态
func (c *Client) DescribeBuildTasks(request *DescribeBuildTasksRequest) (response *DescribeBuildTasksResponse, err error) {
	if request == nil {
		request = NewDescribeBuildTasksRequest()
	}
	response = NewDescribeBuildTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicensesRequest() (request *DescribeLicensesRequest) {
	request = &DescribeLicensesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLicenses")
	return
}

func NewDescribeLicensesResponse() (response *DescribeLicensesResponse) {
	response = &DescribeLicensesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询许可列表
func (c *Client) DescribeLicenses(request *DescribeLicensesRequest) (response *DescribeLicensesResponse, err error) {
	if request == nil {
		request = NewDescribeLicensesRequest()
	}
	response = NewDescribeLicensesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBusinessLogConfigRequest() (request *DescribeBusinessLogConfigRequest) {
	request = &DescribeBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeBusinessLogConfig")
	return
}

func NewDescribeBusinessLogConfigResponse() (response *DescribeBusinessLogConfigResponse) {
	response = &DescribeBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务日志配置项信息
func (c *Client) DescribeBusinessLogConfig(request *DescribeBusinessLogConfigRequest) (response *DescribeBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewDescribeBusinessLogConfigRequest()
	}
	response = NewDescribeBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewTriggerReleasePipelineRunTaskActionRequest() (request *TriggerReleasePipelineRunTaskActionRequest) {
	request = &TriggerReleasePipelineRunTaskActionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "TriggerReleasePipelineRunTaskAction")
	return
}

func NewTriggerReleasePipelineRunTaskActionResponse() (response *TriggerReleasePipelineRunTaskActionResponse) {
	response = &TriggerReleasePipelineRunTaskActionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 触发发布单流水运行任务操作
func (c *Client) TriggerReleasePipelineRunTaskAction(request *TriggerReleasePipelineRunTaskActionRequest) (response *TriggerReleasePipelineRunTaskActionResponse, err error) {
	if request == nil {
		request = NewTriggerReleasePipelineRunTaskActionRequest()
	}
	response = NewTriggerReleasePipelineRunTaskActionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiGroupRequest() (request *DescribeApiGroupRequest) {
	request = &DescribeApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiGroup")
	return
}

func NewDescribeApiGroupResponse() (response *DescribeApiGroupResponse) {
	response = &DescribeApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API分组
func (c *Client) DescribeApiGroup(request *DescribeApiGroupRequest) (response *DescribeApiGroupResponse, err error) {
	if request == nil {
		request = NewDescribeApiGroupRequest()
	}
	response = NewDescribeApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndicesPagedListRequest() (request *DescribeIndicesPagedListRequest) {
	request = &DescribeIndicesPagedListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeIndicesPagedList")
	return
}

func NewDescribeIndicesPagedListResponse() (response *DescribeIndicesPagedListResponse) {
	response = &DescribeIndicesPagedListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所以与分片
func (c *Client) DescribeIndicesPagedList(request *DescribeIndicesPagedListRequest) (response *DescribeIndicesPagedListResponse, err error) {
	if request == nil {
		request = NewDescribeIndicesPagedListRequest()
	}
	response = NewDescribeIndicesPagedListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralMetricDataDimensionRequest() (request *DescribeGeneralMetricDataDimensionRequest) {
	request = &DescribeGeneralMetricDataDimensionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGeneralMetricDataDimension")
	return
}

func NewDescribeGeneralMetricDataDimensionResponse() (response *DescribeGeneralMetricDataDimensionResponse) {
	response = &DescribeGeneralMetricDataDimensionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指标通用索引维度，如微服务网关的服务上游命名空间
func (c *Client) DescribeGeneralMetricDataDimension(request *DescribeGeneralMetricDataDimensionRequest) (response *DescribeGeneralMetricDataDimensionResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralMetricDataDimensionRequest()
	}
	response = NewDescribeGeneralMetricDataDimensionResponse()
	err = c.Send(request, response)
	return
}

func NewImageGetTagListRequest() (request *ImageGetTagListRequest) {
	request = &ImageGetTagListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ImageGetTagList")
	return
}

func NewImageGetTagListResponse() (response *ImageGetTagListResponse) {
	response = &ImageGetTagListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标签页镜像详细信息+部署应用选择镜像 时使用该接口
func (c *Client) ImageGetTagList(request *ImageGetTagListRequest) (response *ImageGetTagListResponse, err error) {
	if request == nil {
		request = NewImageGetTagListRequest()
	}
	response = NewImageGetTagListResponse()
	err = c.Send(request, response)
	return
}

func NewAdjustLaneRulePriorityRequest() (request *AdjustLaneRulePriorityRequest) {
	request = &AdjustLaneRulePriorityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AdjustLaneRulePriority")
	return
}

func NewAdjustLaneRulePriorityResponse() (response *AdjustLaneRulePriorityResponse) {
	response = &AdjustLaneRulePriorityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调整泳道规则优先级
func (c *Client) AdjustLaneRulePriority(request *AdjustLaneRulePriorityRequest) (response *AdjustLaneRulePriorityResponse, err error) {
	if request == nil {
		request = NewAdjustLaneRulePriorityRequest()
	}
	response = NewAdjustLaneRulePriorityResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlarmPolicyRequest() (request *DeleteAlarmPolicyRequest) {
	request = &DeleteAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteAlarmPolicy")
	return
}

func NewDeleteAlarmPolicyResponse() (response *DeleteAlarmPolicyResponse) {
	response = &DeleteAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除告警策略
func (c *Client) DeleteAlarmPolicy(request *DeleteAlarmPolicyRequest) (response *DeleteAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteAlarmPolicyRequest()
	}
	response = NewDeleteAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroserviceRequest() (request *DescribeMicroserviceRequest) {
	request = &DescribeMicroserviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservice")
	return
}

func NewDescribeMicroserviceResponse() (response *DescribeMicroserviceResponse) {
	response = &DescribeMicroserviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微服务详情
func (c *Client) DescribeMicroservice(request *DescribeMicroserviceRequest) (response *DescribeMicroserviceResponse, err error) {
	if request == nil {
		request = NewDescribeMicroserviceRequest()
	}
	response = NewDescribeMicroserviceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvocationStatisticsRequest() (request *DescribeInvocationStatisticsRequest) {
	request = &DescribeInvocationStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationStatistics")
	return
}

func NewDescribeInvocationStatisticsResponse() (response *DescribeInvocationStatisticsResponse) {
	response = &DescribeInvocationStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用监控统计
func (c *Client) DescribeInvocationStatistics(request *DescribeInvocationStatisticsRequest) (response *DescribeInvocationStatisticsResponse, err error) {
	if request == nil {
		request = NewDescribeInvocationStatisticsRequest()
	}
	response = NewDescribeInvocationStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBusinessLogConfigRequest() (request *DeleteBusinessLogConfigRequest) {
	request = &DeleteBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteBusinessLogConfig")
	return
}

func NewDeleteBusinessLogConfigResponse() (response *DeleteBusinessLogConfigResponse) {
	response = &DeleteBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除业务日志配置项
func (c *Client) DeleteBusinessLogConfig(request *DeleteBusinessLogConfigRequest) (response *DeleteBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewDeleteBusinessLogConfigRequest()
	}
	response = NewDeleteBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewFindContainerGroupsRequest() (request *FindContainerGroupsRequest) {
	request = &FindContainerGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "FindContainerGroups")
	return
}

func NewFindContainerGroupsResponse() (response *FindContainerGroupsResponse) {
	response = &FindContainerGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器部署组列表-独立菜单
func (c *Client) FindContainerGroups(request *FindContainerGroupsRequest) (response *FindContainerGroupsResponse, err error) {
	if request == nil {
		request = NewFindContainerGroupsRequest()
	}
	response = NewFindContainerGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewReleasePluginRequest() (request *ReleasePluginRequest) {
	request = &ReleasePluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleasePlugin")
	return
}

func NewReleasePluginResponse() (response *ReleasePluginResponse) {
	response = &ReleasePluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布插件
func (c *Client) ReleasePlugin(request *ReleasePluginRequest) (response *ReleasePluginResponse, err error) {
	if request == nil {
		request = NewReleasePluginRequest()
	}
	response = NewReleasePluginResponse()
	err = c.Send(request, response)
	return
}

func NewBillingOperationDestroyRequest() (request *BillingOperationDestroyRequest) {
	request = &BillingOperationDestroyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "BillingOperationDestroy")
	return
}

func NewBillingOperationDestroyResponse() (response *BillingOperationDestroyResponse) {
	response = &BillingOperationDestroyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费运营端销毁
func (c *Client) BillingOperationDestroy(request *BillingOperationDestroyRequest) (response *BillingOperationDestroyResponse, err error) {
	if request == nil {
		request = NewBillingOperationDestroyRequest()
	}
	response = NewBillingOperationDestroyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRatelimitRequest() (request *DescribeRatelimitRequest) {
	request = &DescribeRatelimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRatelimit")
	return
}

func NewDescribeRatelimitResponse() (response *DescribeRatelimitResponse) {
	response = &DescribeRatelimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 读取限流规则列表
func (c *Client) DescribeRatelimit(request *DescribeRatelimitRequest) (response *DescribeRatelimitResponse, err error) {
	if request == nil {
		request = NewDescribeRatelimitRequest()
	}
	response = NewDescribeRatelimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasePipelineRequest() (request *DescribeReleasePipelineRequest) {
	request = &DescribeReleasePipelineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasePipeline")
	return
}

func NewDescribeReleasePipelineResponse() (response *DescribeReleasePipelineResponse) {
	response = &DescribeReleasePipelineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布单流水线
func (c *Client) DescribeReleasePipeline(request *DescribeReleasePipelineRequest) (response *DescribeReleasePipelineResponse, err error) {
	if request == nil {
		request = NewDescribeReleasePipelineRequest()
	}
	response = NewDescribeReleasePipelineResponse()
	err = c.Send(request, response)
	return
}

func NewOperateUnitNsByMultiCloudRequest() (request *OperateUnitNsByMultiCloudRequest) {
	request = &OperateUnitNsByMultiCloudRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "OperateUnitNsByMultiCloud")
	return
}

func NewOperateUnitNsByMultiCloudResponse() (response *OperateUnitNsByMultiCloudResponse) {
	response = &OperateUnitNsByMultiCloudResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 超云操作单元范围(增删改)
func (c *Client) OperateUnitNsByMultiCloud(request *OperateUnitNsByMultiCloudRequest) (response *OperateUnitNsByMultiCloudResponse, err error) {
	if request == nil {
		request = NewOperateUnitNsByMultiCloudRequest()
	}
	response = NewOperateUnitNsByMultiCloudResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMeshSidecarVersionRequest() (request *DescribeMeshSidecarVersionRequest) {
	request = &DescribeMeshSidecarVersionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMeshSidecarVersion")
	return
}

func NewDescribeMeshSidecarVersionResponse() (response *DescribeMeshSidecarVersionResponse) {
	response = &DescribeMeshSidecarVersionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询sidecar版本
func (c *Client) DescribeMeshSidecarVersion(request *DescribeMeshSidecarVersionRequest) (response *DescribeMeshSidecarVersionResponse, err error) {
	if request == nil {
		request = NewDescribeMeshSidecarVersionRequest()
	}
	response = NewDescribeMeshSidecarVersionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowBatchRecordRequest() (request *DescribeFlowBatchRecordRequest) {
	request = &DescribeFlowBatchRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowBatchRecord")
	return
}

func NewDescribeFlowBatchRecordResponse() (response *DescribeFlowBatchRecordResponse) {
	response = &DescribeFlowBatchRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作流批次的基本信息
func (c *Client) DescribeFlowBatchRecord(request *DescribeFlowBatchRecordRequest) (response *DescribeFlowBatchRecordResponse, err error) {
	if request == nil {
		request = NewDescribeFlowBatchRecordRequest()
	}
	response = NewDescribeFlowBatchRecordResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMsRunningApplicationsRequest() (request *DescribeMsRunningApplicationsRequest) {
	request = &DescribeMsRunningApplicationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsRunningApplications")
	return
}

func NewDescribeMsRunningApplicationsResponse() (response *DescribeMsRunningApplicationsResponse) {
	response = &DescribeMsRunningApplicationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微服务运行态应用列表
func (c *Client) DescribeMsRunningApplications(request *DescribeMsRunningApplicationsRequest) (response *DescribeMsRunningApplicationsResponse, err error) {
	if request == nil {
		request = NewDescribeMsRunningApplicationsRequest()
	}
	response = NewDescribeMsRunningApplicationsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceTsfConfigRequest() (request *DescribeResourceTsfConfigRequest) {
	request = &DescribeResourceTsfConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceTsfConfig")
	return
}

func NewDescribeResourceTsfConfigResponse() (response *DescribeResourceTsfConfigResponse) {
	response = &DescribeResourceTsfConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF配置信息概览
func (c *Client) DescribeResourceTsfConfig(request *DescribeResourceTsfConfigRequest) (response *DescribeResourceTsfConfigResponse, err error) {
	if request == nil {
		request = NewDescribeResourceTsfConfigRequest()
	}
	response = NewDescribeResourceTsfConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayOAuthPluginRequest() (request *CreateGatewayOAuthPluginRequest) {
	request = &CreateGatewayOAuthPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayOAuthPlugin")
	return
}

func NewCreateGatewayOAuthPluginResponse() (response *CreateGatewayOAuthPluginResponse) {
	response = &CreateGatewayOAuthPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增网关OAuth认证插件
func (c *Client) CreateGatewayOAuthPlugin(request *CreateGatewayOAuthPluginRequest) (response *CreateGatewayOAuthPluginResponse, err error) {
	if request == nil {
		request = NewCreateGatewayOAuthPluginRequest()
	}
	response = NewCreateGatewayOAuthPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigTemplatesRequest() (request *DescribeConfigTemplatesRequest) {
	request = &DescribeConfigTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigTemplates")
	return
}

func NewDescribeConfigTemplatesResponse() (response *DescribeConfigTemplatesResponse) {
	response = &DescribeConfigTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配置模板
func (c *Client) DescribeConfigTemplates(request *DescribeConfigTemplatesRequest) (response *DescribeConfigTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeConfigTemplatesRequest()
	}
	response = NewDescribeConfigTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSidecarFilterRequest() (request *DescribeSidecarFilterRequest) {
	request = &DescribeSidecarFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSidecarFilter")
	return
}

func NewDescribeSidecarFilterResponse() (response *DescribeSidecarFilterResponse) {
	response = &DescribeSidecarFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单个Sidecar过滤器
func (c *Client) DescribeSidecarFilter(request *DescribeSidecarFilterRequest) (response *DescribeSidecarFilterResponse, err error) {
	if request == nil {
		request = NewDescribeSidecarFilterRequest()
	}
	response = NewDescribeSidecarFilterResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGatewayApiRequest() (request *DeleteGatewayApiRequest) {
	request = &DeleteGatewayApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteGatewayApi")
	return
}

func NewDeleteGatewayApiResponse() (response *DeleteGatewayApiResponse) {
	response = &DeleteGatewayApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除API
func (c *Client) DeleteGatewayApi(request *DeleteGatewayApiRequest) (response *DeleteGatewayApiResponse, err error) {
	if request == nil {
		request = NewDeleteGatewayApiRequest()
	}
	response = NewDeleteGatewayApiResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventCategoryNameRequest() (request *DescribeEventCategoryNameRequest) {
	request = &DescribeEventCategoryNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeEventCategoryName")
	return
}

func NewDescribeEventCategoryNameResponse() (response *DescribeEventCategoryNameResponse) {
	response = &DescribeEventCategoryNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//  事件名，事件摘要
func (c *Client) DescribeEventCategoryName(request *DescribeEventCategoryNameRequest) (response *DescribeEventCategoryNameResponse, err error) {
	if request == nil {
		request = NewDescribeEventCategoryNameRequest()
	}
	response = NewDescribeEventCategoryNameResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationAttributeRequest() (request *DescribeApplicationAttributeRequest) {
	request = &DescribeApplicationAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationAttribute")
	return
}

func NewDescribeApplicationAttributeResponse() (response *DescribeApplicationAttributeResponse) {
	response = &DescribeApplicationAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用列表其它字段，如实例数量信息等
func (c *Client) DescribeApplicationAttribute(request *DescribeApplicationAttributeRequest) (response *DescribeApplicationAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationAttributeRequest()
	}
	response = NewDescribeApplicationAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewEnableAuthorizationRequest() (request *EnableAuthorizationRequest) {
	request = &EnableAuthorizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableAuthorization")
	return
}

func NewEnableAuthorizationResponse() (response *EnableAuthorizationResponse) {
	response = &EnableAuthorizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用权限
func (c *Client) EnableAuthorization(request *EnableAuthorizationRequest) (response *EnableAuthorizationResponse, err error) {
	if request == nil {
		request = NewEnableAuthorizationRequest()
	}
	response = NewEnableAuthorizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterGroupCountRequest() (request *DescribeClusterGroupCountRequest) {
	request = &DescribeClusterGroupCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterGroupCount")
	return
}

func NewDescribeClusterGroupCountResponse() (response *DescribeClusterGroupCountResponse) {
	response = &DescribeClusterGroupCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群部署组统计信息
func (c *Client) DescribeClusterGroupCount(request *DescribeClusterGroupCountRequest) (response *DescribeClusterGroupCountResponse, err error) {
	if request == nil {
		request = NewDescribeClusterGroupCountRequest()
	}
	response = NewDescribeClusterGroupCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScalableRuleAttributeRequest() (request *DescribeScalableRuleAttributeRequest) {
	request = &DescribeScalableRuleAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeScalableRuleAttribute")
	return
}

func NewDescribeScalableRuleAttributeResponse() (response *DescribeScalableRuleAttributeResponse) {
	response = &DescribeScalableRuleAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性扩缩容规则属性
func (c *Client) DescribeScalableRuleAttribute(request *DescribeScalableRuleAttributeRequest) (response *DescribeScalableRuleAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeScalableRuleAttributeRequest()
	}
	response = NewDescribeScalableRuleAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigReleaseLogsRequest() (request *DescribeConfigReleaseLogsRequest) {
	request = &DescribeConfigReleaseLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleaseLogs")
	return
}

func NewDescribeConfigReleaseLogsResponse() (response *DescribeConfigReleaseLogsResponse) {
	response = &DescribeConfigReleaseLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配置发布历史
func (c *Client) DescribeConfigReleaseLogs(request *DescribeConfigReleaseLogsRequest) (response *DescribeConfigReleaseLogsResponse, err error) {
	if request == nil {
		request = NewDescribeConfigReleaseLogsRequest()
	}
	response = NewDescribeConfigReleaseLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroservicesByAssociateApplicationRequest() (request *DescribeMicroservicesByAssociateApplicationRequest) {
	request = &DescribeMicroservicesByAssociateApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservicesByAssociateApplication")
	return
}

func NewDescribeMicroservicesByAssociateApplicationResponse() (response *DescribeMicroservicesByAssociateApplicationResponse) {
	response = &DescribeMicroservicesByAssociateApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 按关联的应用 ID 获取服务列表
func (c *Client) DescribeMicroservicesByAssociateApplication(request *DescribeMicroservicesByAssociateApplicationRequest) (response *DescribeMicroservicesByAssociateApplicationResponse, err error) {
	if request == nil {
		request = NewDescribeMicroservicesByAssociateApplicationRequest()
	}
	response = NewDescribeMicroservicesByAssociateApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSecretRequest() (request *CreateSecretRequest) {
	request = &CreateSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateSecret")
	return
}

func NewCreateSecretResponse() (response *CreateSecretResponse) {
	response = &CreateSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新建访问凭证信息
func (c *Client) CreateSecret(request *CreateSecretRequest) (response *CreateSecretResponse, err error) {
	if request == nil {
		request = NewCreateSecretRequest()
	}
	response = NewCreateSecretResponse()
	err = c.Send(request, response)
	return
}

func NewBindApiGroupRequest() (request *BindApiGroupRequest) {
	request = &BindApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "BindApiGroup")
	return
}

func NewBindApiGroupResponse() (response *BindApiGroupResponse) {
	response = &BindApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 网关与API分组批量绑定
func (c *Client) BindApiGroup(request *BindApiGroupRequest) (response *BindApiGroupResponse, err error) {
	if request == nil {
		request = NewBindApiGroupRequest()
	}
	response = NewBindApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateConfigTemplateRequest() (request *UpdateConfigTemplateRequest) {
	request = &UpdateConfigTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateConfigTemplate")
	return
}

func NewUpdateConfigTemplateResponse() (response *UpdateConfigTemplateResponse) {
	response = &UpdateConfigTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新参数模板
func (c *Client) UpdateConfigTemplate(request *UpdateConfigTemplateRequest) (response *UpdateConfigTemplateResponse, err error) {
	if request == nil {
		request = NewUpdateConfigTemplateRequest()
	}
	response = NewUpdateConfigTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiDetailRequest() (request *DescribeApiDetailRequest) {
	request = &DescribeApiDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiDetail")
	return
}

func NewDescribeApiDetailResponse() (response *DescribeApiDetailResponse) {
	response = &DescribeApiDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API详情
func (c *Client) DescribeApiDetail(request *DescribeApiDetailRequest) (response *DescribeApiDetailResponse, err error) {
	if request == nil {
		request = NewDescribeApiDetailRequest()
	}
	response = NewDescribeApiDetailResponse()
	err = c.Send(request, response)
	return
}

func NewHideInstanceRequest() (request *HideInstanceRequest) {
	request = &HideInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "HideInstance")
	return
}

func NewHideInstanceResponse() (response *HideInstanceResponse) {
	response = &HideInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 屏蔽注册实例
func (c *Client) HideInstance(request *HideInstanceRequest) (response *HideInstanceResponse, err error) {
	if request == nil {
		request = NewHideInstanceRequest()
	}
	response = NewHideInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProgramsRequest() (request *DescribeProgramsRequest) {
	request = &DescribeProgramsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePrograms")
	return
}

func NewDescribeProgramsResponse() (response *DescribeProgramsResponse) {
	response = &DescribeProgramsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询数据集列表
func (c *Client) DescribePrograms(request *DescribeProgramsRequest) (response *DescribeProgramsResponse, err error) {
	if request == nil {
		request = NewDescribeProgramsRequest()
	}
	response = NewDescribeProgramsResponse()
	err = c.Send(request, response)
	return
}

func NewGetDefaultValueRequest() (request *GetDefaultValueRequest) {
	request = &GetDefaultValueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetDefaultValue")
	return
}

func NewGetDefaultValueResponse() (response *GetDefaultValueResponse) {
	response = &GetDefaultValueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据参数获取默认值
func (c *Client) GetDefaultValue(request *GetDefaultValueRequest) (response *GetDefaultValueResponse, err error) {
	if request == nil {
		request = NewGetDefaultValueRequest()
	}
	response = NewGetDefaultValueResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApisWithPluginRequest() (request *DescribeApisWithPluginRequest) {
	request = &DescribeApisWithPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApisWithPlugin")
	return
}

func NewDescribeApisWithPluginResponse() (response *DescribeApisWithPluginResponse) {
	response = &DescribeApisWithPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个插件下可用于绑定的API
func (c *Client) DescribeApisWithPlugin(request *DescribeApisWithPluginRequest) (response *DescribeApisWithPluginResponse, err error) {
	if request == nil {
		request = NewDescribeApisWithPluginRequest()
	}
	response = NewDescribeApisWithPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeThreadDetailListRequest() (request *DescribeThreadDetailListRequest) {
	request = &DescribeThreadDetailListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeThreadDetailList")
	return
}

func NewDescribeThreadDetailListResponse() (response *DescribeThreadDetailListResponse) {
	response = &DescribeThreadDetailListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询线程详情列表
func (c *Client) DescribeThreadDetailList(request *DescribeThreadDetailListRequest) (response *DescribeThreadDetailListResponse, err error) {
	if request == nil {
		request = NewDescribeThreadDetailListRequest()
	}
	response = NewDescribeThreadDetailListResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRatelimitRequest() (request *UpdateRatelimitRequest) {
	request = &UpdateRatelimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateRatelimit")
	return
}

func NewUpdateRatelimitResponse() (response *UpdateRatelimitResponse) {
	response = &UpdateRatelimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改限流规则
func (c *Client) UpdateRatelimit(request *UpdateRatelimitRequest) (response *UpdateRatelimitResponse, err error) {
	if request == nil {
		request = NewUpdateRatelimitRequest()
	}
	response = NewUpdateRatelimitResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateFilebeatConfigRequest() (request *AssociateFilebeatConfigRequest) {
	request = &AssociateFilebeatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AssociateFilebeatConfig")
	return
}

func NewAssociateFilebeatConfigResponse() (response *AssociateFilebeatConfigResponse) {
	response = &AssociateFilebeatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联通用filebeat配置和部署组
func (c *Client) AssociateFilebeatConfig(request *AssociateFilebeatConfigRequest) (response *AssociateFilebeatConfigResponse, err error) {
	if request == nil {
		request = NewAssociateFilebeatConfigRequest()
	}
	response = NewAssociateFilebeatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSpanBusinessLogRequest() (request *SearchSpanBusinessLogRequest) {
	request = &SearchSpanBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchSpanBusinessLog")
	return
}

func NewSearchSpanBusinessLogResponse() (response *SearchSpanBusinessLogResponse) {
	response = &SearchSpanBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务日志搜索（关联调用链Span）
func (c *Client) SearchSpanBusinessLog(request *SearchSpanBusinessLogRequest) (response *SearchSpanBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchSpanBusinessLogRequest()
	}
	response = NewSearchSpanBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespaceMicroserviceCountRequest() (request *DescribeNamespaceMicroserviceCountRequest) {
	request = &DescribeNamespaceMicroserviceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaceMicroserviceCount")
	return
}

func NewDescribeNamespaceMicroserviceCountResponse() (response *DescribeNamespaceMicroserviceCountResponse) {
	response = &DescribeNamespaceMicroserviceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 服务监控查询命名空间下服务数量
func (c *Client) DescribeNamespaceMicroserviceCount(request *DescribeNamespaceMicroserviceCountRequest) (response *DescribeNamespaceMicroserviceCountResponse, err error) {
	if request == nil {
		request = NewDescribeNamespaceMicroserviceCountRequest()
	}
	response = NewDescribeNamespaceMicroserviceCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupGatewaysRequest() (request *DescribeGroupGatewaysRequest) {
	request = &DescribeGroupGatewaysRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupGateways")
	return
}

func NewDescribeGroupGatewaysResponse() (response *DescribeGroupGatewaysResponse) {
	response = &DescribeGroupGatewaysResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个网关绑定的API 分组信息列表
func (c *Client) DescribeGroupGateways(request *DescribeGroupGatewaysRequest) (response *DescribeGroupGatewaysResponse, err error) {
	if request == nil {
		request = NewDescribeGroupGatewaysRequest()
	}
	response = NewDescribeGroupGatewaysResponse()
	err = c.Send(request, response)
	return
}

func NewDisableAuthorizationRequest() (request *DisableAuthorizationRequest) {
	request = &DisableAuthorizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableAuthorization")
	return
}

func NewDisableAuthorizationResponse() (response *DisableAuthorizationResponse) {
	response = &DisableAuthorizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭权限
func (c *Client) DisableAuthorization(request *DisableAuthorizationRequest) (response *DisableAuthorizationResponse, err error) {
	if request == nil {
		request = NewDisableAuthorizationRequest()
	}
	response = NewDisableAuthorizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationInstancesRequest() (request *DescribeApplicationInstancesRequest) {
	request = &DescribeApplicationInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationInstances")
	return
}

func NewDescribeApplicationInstancesResponse() (response *DescribeApplicationInstancesResponse) {
	response = &DescribeApplicationInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用机器列表
func (c *Client) DescribeApplicationInstances(request *DescribeApplicationInstancesRequest) (response *DescribeApplicationInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationInstancesRequest()
	}
	response = NewDescribeApplicationInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthorizationTypeRequest() (request *DescribeAuthorizationTypeRequest) {
	request = &DescribeAuthorizationTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorizationType")
	return
}

func NewDescribeAuthorizationTypeResponse() (response *DescribeAuthorizationTypeResponse) {
	response = &DescribeAuthorizationTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询鉴权类型，返回：D：未启用；B：黑名单模式；W：白名单模式
func (c *Client) DescribeAuthorizationType(request *DescribeAuthorizationTypeRequest) (response *DescribeAuthorizationTypeResponse, err error) {
	if request == nil {
		request = NewDescribeAuthorizationTypeRequest()
	}
	response = NewDescribeAuthorizationTypeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmPolicyRequest() (request *ModifyAlarmPolicyRequest) {
	request = &ModifyAlarmPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyAlarmPolicy")
	return
}

func NewModifyAlarmPolicyResponse() (response *ModifyAlarmPolicyResponse) {
	response = &ModifyAlarmPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警策略
func (c *Client) ModifyAlarmPolicy(request *ModifyAlarmPolicyRequest) (response *ModifyAlarmPolicyResponse, err error) {
	if request == nil {
		request = NewModifyAlarmPolicyRequest()
	}
	response = NewModifyAlarmPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerEventsRequest() (request *DescribeContainerEventsRequest) {
	request = &DescribeContainerEventsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerEvents")
	return
}

func NewDescribeContainerEventsResponse() (response *DescribeContainerEventsResponse) {
	response = &DescribeContainerEventsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取容器事件列表
func (c *Client) DescribeContainerEvents(request *DescribeContainerEventsRequest) (response *DescribeContainerEventsResponse, err error) {
	if request == nil {
		request = NewDescribeContainerEventsRequest()
	}
	response = NewDescribeContainerEventsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewListAlarmReceiversRequest() (request *ListAlarmReceiversRequest) {
	request = &ListAlarmReceiversRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListAlarmReceivers")
	return
}

func NewListAlarmReceiversResponse() (response *ListAlarmReceiversResponse) {
	response = &ListAlarmReceiversResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警接收者列表
func (c *Client) ListAlarmReceivers(request *ListAlarmReceiversRequest) (response *ListAlarmReceiversResponse, err error) {
	if request == nil {
		request = NewListAlarmReceiversRequest()
	}
	response = NewListAlarmReceiversResponse()
	err = c.Send(request, response)
	return
}

func NewDisableUnitRouteRequest() (request *DisableUnitRouteRequest) {
	request = &DisableUnitRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableUnitRoute")
	return
}

func NewDisableUnitRouteResponse() (response *DisableUnitRouteResponse) {
	response = &DisableUnitRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用单元化路由
func (c *Client) DisableUnitRoute(request *DisableUnitRouteRequest) (response *DisableUnitRouteResponse, err error) {
	if request == nil {
		request = NewDisableUnitRouteRequest()
	}
	response = NewDisableUnitRouteResponse()
	err = c.Send(request, response)
	return
}

func NewExpandInstanceScriptsRequest() (request *ExpandInstanceScriptsRequest) {
	request = &ExpandInstanceScriptsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ExpandInstanceScripts")
	return
}

func NewExpandInstanceScriptsResponse() (response *ExpandInstanceScriptsResponse) {
	response = &ExpandInstanceScriptsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ExpandInstanceScripts
func (c *Client) ExpandInstanceScripts(request *ExpandInstanceScriptsRequest) (response *ExpandInstanceScriptsResponse, err error) {
	if request == nil {
		request = NewExpandInstanceScriptsRequest()
	}
	response = NewExpandInstanceScriptsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasePipelineTaskRequest() (request *DescribeReleasePipelineTaskRequest) {
	request = &DescribeReleasePipelineTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasePipelineTask")
	return
}

func NewDescribeReleasePipelineTaskResponse() (response *DescribeReleasePipelineTaskResponse) {
	response = &DescribeReleasePipelineTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布单流水任务
func (c *Client) DescribeReleasePipelineTask(request *DescribeReleasePipelineTaskRequest) (response *DescribeReleasePipelineTaskResponse, err error) {
	if request == nil {
		request = NewDescribeReleasePipelineTaskRequest()
	}
	response = NewDescribeReleasePipelineTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayWeChatMiniProgramLoginPluginRequest() (request *DescribeGatewayWeChatMiniProgramLoginPluginRequest) {
	request = &DescribeGatewayWeChatMiniProgramLoginPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayWeChatMiniProgramLoginPlugin")
	return
}

func NewDescribeGatewayWeChatMiniProgramLoginPluginResponse() (response *DescribeGatewayWeChatMiniProgramLoginPluginResponse) {
	response = &DescribeGatewayWeChatMiniProgramLoginPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微信登录插件详情信息
func (c *Client) DescribeGatewayWeChatMiniProgramLoginPlugin(request *DescribeGatewayWeChatMiniProgramLoginPluginRequest) (response *DescribeGatewayWeChatMiniProgramLoginPluginResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayWeChatMiniProgramLoginPluginRequest()
	}
	response = NewDescribeGatewayWeChatMiniProgramLoginPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnitRuleCreateTypeRequest() (request *DescribeUnitRuleCreateTypeRequest) {
	request = &DescribeUnitRuleCreateTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRuleCreateType")
	return
}

func NewDescribeUnitRuleCreateTypeResponse() (response *DescribeUnitRuleCreateTypeResponse) {
	response = &DescribeUnitRuleCreateTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单元化规则创建类型
func (c *Client) DescribeUnitRuleCreateType(request *DescribeUnitRuleCreateTypeRequest) (response *DescribeUnitRuleCreateTypeResponse, err error) {
	if request == nil {
		request = NewDescribeUnitRuleCreateTypeRequest()
	}
	response = NewDescribeUnitRuleCreateTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvocationMetricDataPointRequest() (request *DescribeInvocationMetricDataPointRequest) {
	request = &DescribeInvocationMetricDataPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationMetricDataPoint")
	return
}

func NewDescribeInvocationMetricDataPointResponse() (response *DescribeInvocationMetricDataPointResponse) {
	response = &DescribeInvocationMetricDataPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单值指标维度
func (c *Client) DescribeInvocationMetricDataPoint(request *DescribeInvocationMetricDataPointRequest) (response *DescribeInvocationMetricDataPointResponse, err error) {
	if request == nil {
		request = NewDescribeInvocationMetricDataPointRequest()
	}
	response = NewDescribeInvocationMetricDataPointResponse()
	err = c.Send(request, response)
	return
}

func NewEnableRouteRuleRequest() (request *EnableRouteRuleRequest) {
	request = &EnableRouteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableRouteRule")
	return
}

func NewEnableRouteRuleResponse() (response *EnableRouteRuleResponse) {
	response = &EnableRouteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用路由规则
func (c *Client) EnableRouteRule(request *EnableRouteRuleRequest) (response *EnableRouteRuleResponse, err error) {
	if request == nil {
		request = NewEnableRouteRuleRequest()
	}
	response = NewEnableRouteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLaneGroupRequest() (request *DeleteLaneGroupRequest) {
	request = &DeleteLaneGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteLaneGroup")
	return
}

func NewDeleteLaneGroupResponse() (response *DeleteLaneGroupResponse) {
	response = &DeleteLaneGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除泳道部署组
func (c *Client) DeleteLaneGroup(request *DeleteLaneGroupRequest) (response *DeleteLaneGroupResponse, err error) {
	if request == nil {
		request = NewDeleteLaneGroupRequest()
	}
	response = NewDeleteLaneGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBillingDealRecordsRequest() (request *DescribeBillingDealRecordsRequest) {
	request = &DescribeBillingDealRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeBillingDealRecords")
	return
}

func NewDescribeBillingDealRecordsResponse() (response *DescribeBillingDealRecordsResponse) {
	response = &DescribeBillingDealRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询计费订单记录
func (c *Client) DescribeBillingDealRecords(request *DescribeBillingDealRecordsRequest) (response *DescribeBillingDealRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeBillingDealRecordsRequest()
	}
	response = NewDescribeBillingDealRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePublicConfigRequest() (request *DeletePublicConfigRequest) {
	request = &DeletePublicConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeletePublicConfig")
	return
}

func NewDeletePublicConfigResponse() (response *DeletePublicConfigResponse) {
	response = &DeletePublicConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除公共配置项
func (c *Client) DeletePublicConfig(request *DeletePublicConfigRequest) (response *DeletePublicConfigResponse, err error) {
	if request == nil {
		request = NewDeletePublicConfigRequest()
	}
	response = NewDeletePublicConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApiRateLimitRuleRequest() (request *UpdateApiRateLimitRuleRequest) {
	request = &UpdateApiRateLimitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRule")
	return
}

func NewUpdateApiRateLimitRuleResponse() (response *UpdateApiRateLimitRuleResponse) {
	response = &UpdateApiRateLimitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新API限流规则
func (c *Client) UpdateApiRateLimitRule(request *UpdateApiRateLimitRuleRequest) (response *UpdateApiRateLimitRuleResponse, err error) {
	if request == nil {
		request = NewUpdateApiRateLimitRuleRequest()
	}
	response = NewUpdateApiRateLimitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateOrUpdateUnitRuleByMultiCloudRequest() (request *CreateOrUpdateUnitRuleByMultiCloudRequest) {
	request = &CreateOrUpdateUnitRuleByMultiCloudRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateOrUpdateUnitRuleByMultiCloud")
	return
}

func NewCreateOrUpdateUnitRuleByMultiCloudResponse() (response *CreateOrUpdateUnitRuleByMultiCloudResponse) {
	response = &CreateOrUpdateUnitRuleByMultiCloudResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建/更新单元化规则并启用
func (c *Client) CreateOrUpdateUnitRuleByMultiCloud(request *CreateOrUpdateUnitRuleByMultiCloudRequest) (response *CreateOrUpdateUnitRuleByMultiCloudResponse, err error) {
	if request == nil {
		request = NewCreateOrUpdateUnitRuleByMultiCloudRequest()
	}
	response = NewCreateOrUpdateUnitRuleByMultiCloudResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSimpleReleasesRequest() (request *DescribeSimpleReleasesRequest) {
	request = &DescribeSimpleReleasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleReleases")
	return
}

func NewDescribeSimpleReleasesResponse() (response *DescribeSimpleReleasesResponse) {
	response = &DescribeSimpleReleasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询简单发布单列表
func (c *Client) DescribeSimpleReleases(request *DescribeSimpleReleasesRequest) (response *DescribeSimpleReleasesResponse, err error) {
	if request == nil {
		request = NewDescribeSimpleReleasesRequest()
	}
	response = NewDescribeSimpleReleasesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFilebeatConfigRequest() (request *DeleteFilebeatConfigRequest) {
	request = &DeleteFilebeatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteFilebeatConfig")
	return
}

func NewDeleteFilebeatConfigResponse() (response *DeleteFilebeatConfigResponse) {
	response = &DeleteFilebeatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除filebeat配置
func (c *Client) DeleteFilebeatConfig(request *DeleteFilebeatConfigRequest) (response *DeleteFilebeatConfigResponse, err error) {
	if request == nil {
		request = NewDeleteFilebeatConfigRequest()
	}
	response = NewDeleteFilebeatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeHostResourceUsageRequest() (request *DescribeHostResourceUsageRequest) {
	request = &DescribeHostResourceUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeHostResourceUsage")
	return
}

func NewDescribeHostResourceUsageResponse() (response *DescribeHostResourceUsageResponse) {
	response = &DescribeHostResourceUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 实例主机信息概览接口
func (c *Client) DescribeHostResourceUsage(request *DescribeHostResourceUsageRequest) (response *DescribeHostResourceUsageResponse, err error) {
	if request == nil {
		request = NewDescribeHostResourceUsageRequest()
	}
	response = NewDescribeHostResourceUsageResponse()
	err = c.Send(request, response)
	return
}

func NewContinueRunFailedTaskBatchRequest() (request *ContinueRunFailedTaskBatchRequest) {
	request = &ContinueRunFailedTaskBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ContinueRunFailedTaskBatch")
	return
}

func NewContinueRunFailedTaskBatchResponse() (response *ContinueRunFailedTaskBatchResponse) {
	response = &ContinueRunFailedTaskBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对执行失败的任务批次执行续跑
func (c *Client) ContinueRunFailedTaskBatch(request *ContinueRunFailedTaskBatchRequest) (response *ContinueRunFailedTaskBatchResponse, err error) {
	if request == nil {
		request = NewContinueRunFailedTaskBatchRequest()
	}
	response = NewContinueRunFailedTaskBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceResourceUsageRequest() (request *DescribeInstanceResourceUsageRequest) {
	request = &DescribeInstanceResourceUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInstanceResourceUsage")
	return
}

func NewDescribeInstanceResourceUsageResponse() (response *DescribeInstanceResourceUsageResponse) {
	response = &DescribeInstanceResourceUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF实例相关信息概览接口
func (c *Client) DescribeInstanceResourceUsage(request *DescribeInstanceResourceUsageRequest) (response *DescribeInstanceResourceUsageResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceResourceUsageRequest()
	}
	response = NewDescribeInstanceResourceUsageResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
	request = &DeleteConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteConfig")
	return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
	response = &DeleteConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除配置项
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
	if request == nil {
		request = NewDeleteConfigRequest()
	}
	response = NewDeleteConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDisableNamespaceAffinityRequest() (request *DisableNamespaceAffinityRequest) {
	request = &DisableNamespaceAffinityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableNamespaceAffinity")
	return
}

func NewDisableNamespaceAffinityResponse() (response *DisableNamespaceAffinityResponse) {
	response = &DisableNamespaceAffinityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停用命名空间就近访问策略
func (c *Client) DisableNamespaceAffinity(request *DisableNamespaceAffinityRequest) (response *DisableNamespaceAffinityResponse, err error) {
	if request == nil {
		request = NewDisableNamespaceAffinityRequest()
	}
	response = NewDisableNamespaceAffinityResponse()
	err = c.Send(request, response)
	return
}

func NewListGroupPodRequest() (request *ListGroupPodRequest) {
	request = &ListGroupPodRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListGroupPod")
	return
}

func NewListGroupPodResponse() (response *ListGroupPodResponse) {
	response = &ListGroupPodResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取容器部署组实例列表
func (c *Client) ListGroupPod(request *ListGroupPodRequest) (response *ListGroupPodResponse, err error) {
	if request == nil {
		request = NewListGroupPodRequest()
	}
	response = NewListGroupPodResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRealtimeBusinessLogRequest() (request *SearchRealtimeBusinessLogRequest) {
	request = &SearchRealtimeBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchRealtimeBusinessLog")
	return
}

func NewSearchRealtimeBusinessLogResponse() (response *SearchRealtimeBusinessLogResponse) {
	response = &SearchRealtimeBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务日志实时检索
func (c *Client) SearchRealtimeBusinessLog(request *SearchRealtimeBusinessLogRequest) (response *SearchRealtimeBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchRealtimeBusinessLogRequest()
	}
	response = NewSearchRealtimeBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRepositoryRequest() (request *DescribeRepositoryRequest) {
	request = &DescribeRepositoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRepository")
	return
}

func NewDescribeRepositoryResponse() (response *DescribeRepositoryResponse) {
	response = &DescribeRepositoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询仓库信息
func (c *Client) DescribeRepository(request *DescribeRepositoryRequest) (response *DescribeRepositoryResponse, err error) {
	if request == nil {
		request = NewDescribeRepositoryRequest()
	}
	response = NewDescribeRepositoryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskBatchRecordsRequest() (request *DescribeTaskBatchRecordsRequest) {
	request = &DescribeTaskBatchRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskBatchRecords")
	return
}

func NewDescribeTaskBatchRecordsResponse() (response *DescribeTaskBatchRecordsResponse) {
	response = &DescribeTaskBatchRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 翻页查询任务批次信息
func (c *Client) DescribeTaskBatchRecords(request *DescribeTaskBatchRecordsRequest) (response *DescribeTaskBatchRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskBatchRecordsRequest()
	}
	response = NewDescribeTaskBatchRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMicroserviceWithDetailRespRequest() (request *CreateMicroserviceWithDetailRespRequest) {
	request = &CreateMicroserviceWithDetailRespRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateMicroserviceWithDetailResp")
	return
}

func NewCreateMicroserviceWithDetailRespResponse() (response *CreateMicroserviceWithDetailRespResponse) {
	response = &CreateMicroserviceWithDetailRespResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增微服务返回id
func (c *Client) CreateMicroserviceWithDetailResp(request *CreateMicroserviceWithDetailRespRequest) (response *CreateMicroserviceWithDetailRespResponse, err error) {
	if request == nil {
		request = NewCreateMicroserviceWithDetailRespRequest()
	}
	response = NewCreateMicroserviceWithDetailRespResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
	request = &DescribeResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeResources")
	return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
	response = &DescribeResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询资源列表
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeResourcesRequest()
	}
	response = NewDescribeResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationsOtherRequest() (request *DescribeApplicationsOtherRequest) {
	request = &DescribeApplicationsOtherRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationsOther")
	return
}

func NewDescribeApplicationsOtherResponse() (response *DescribeApplicationsOtherResponse) {
	response = &DescribeApplicationsOtherResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) DescribeApplicationsOther(request *DescribeApplicationsOtherRequest) (response *DescribeApplicationsOtherResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationsOtherRequest()
	}
	response = NewDescribeApplicationsOtherResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRelatedGroupRequest() (request *DeleteRelatedGroupRequest) {
	request = &DeleteRelatedGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteRelatedGroup")
	return
}

func NewDeleteRelatedGroupResponse() (response *DeleteRelatedGroupResponse) {
	response = &DeleteRelatedGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除关联部署组
func (c *Client) DeleteRelatedGroup(request *DeleteRelatedGroupRequest) (response *DeleteRelatedGroupResponse, err error) {
	if request == nil {
		request = NewDeleteRelatedGroupRequest()
	}
	response = NewDeleteRelatedGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroserviceMetasRequest() (request *DescribeMicroserviceMetasRequest) {
	request = &DescribeMicroserviceMetasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroserviceMetas")
	return
}

func NewDescribeMicroserviceMetasResponse() (response *DescribeMicroserviceMetasResponse) {
	response = &DescribeMicroserviceMetasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微服务元数据
func (c *Client) DescribeMicroserviceMetas(request *DescribeMicroserviceMetasRequest) (response *DescribeMicroserviceMetasResponse, err error) {
	if request == nil {
		request = NewDescribeMicroserviceMetasRequest()
	}
	response = NewDescribeMicroserviceMetasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeliveryConfigRequest() (request *DescribeDeliveryConfigRequest) {
	request = &DescribeDeliveryConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeliveryConfig")
	return
}

func NewDescribeDeliveryConfigResponse() (response *DescribeDeliveryConfigResponse) {
	response = &DescribeDeliveryConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个投递项配置信息
func (c *Client) DescribeDeliveryConfig(request *DescribeDeliveryConfigRequest) (response *DescribeDeliveryConfigResponse, err error) {
	if request == nil {
		request = NewDescribeDeliveryConfigRequest()
	}
	response = NewDescribeDeliveryConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMsRouteFallbackRequest() (request *DescribeMsRouteFallbackRequest) {
	request = &DescribeMsRouteFallbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMsRouteFallback")
	return
}

func NewDescribeMsRouteFallbackResponse() (response *DescribeMsRouteFallbackResponse) {
	response = &DescribeMsRouteFallbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微服务路由保护策略启停状态
func (c *Client) DescribeMsRouteFallback(request *DescribeMsRouteFallbackRequest) (response *DescribeMsRouteFallbackResponse, err error) {
	if request == nil {
		request = NewDescribeMsRouteFallbackRequest()
	}
	response = NewDescribeMsRouteFallbackResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKubeInjectConfigRequest() (request *DescribeKubeInjectConfigRequest) {
	request = &DescribeKubeInjectConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeKubeInjectConfig")
	return
}

func NewDescribeKubeInjectConfigResponse() (response *DescribeKubeInjectConfigResponse) {
	response = &DescribeKubeInjectConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器轻量化配置
func (c *Client) DescribeKubeInjectConfig(request *DescribeKubeInjectConfigRequest) (response *DescribeKubeInjectConfigResponse, err error) {
	if request == nil {
		request = NewDescribeKubeInjectConfigRequest()
	}
	response = NewDescribeKubeInjectConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigSummaryRequest() (request *DescribeConfigSummaryRequest) {
	request = &DescribeConfigSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigSummary")
	return
}

func NewDescribeConfigSummaryResponse() (response *DescribeConfigSummaryResponse) {
	response = &DescribeConfigSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配置汇总列表
func (c *Client) DescribeConfigSummary(request *DescribeConfigSummaryRequest) (response *DescribeConfigSummaryResponse, err error) {
	if request == nil {
		request = NewDescribeConfigSummaryRequest()
	}
	response = NewDescribeConfigSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicConfigSummaryRequest() (request *DescribePublicConfigSummaryRequest) {
	request = &DescribePublicConfigSummaryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigSummary")
	return
}

func NewDescribePublicConfigSummaryResponse() (response *DescribePublicConfigSummaryResponse) {
	response = &DescribePublicConfigSummaryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共配置汇总列表
func (c *Client) DescribePublicConfigSummary(request *DescribePublicConfigSummaryRequest) (response *DescribePublicConfigSummaryResponse, err error) {
	if request == nil {
		request = NewDescribePublicConfigSummaryRequest()
	}
	response = NewDescribePublicConfigSummaryResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRatelimitCommitConfigRequest() (request *DescribeRatelimitCommitConfigRequest) {
	request = &DescribeRatelimitCommitConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRatelimitCommitConfig")
	return
}

func NewDescribeRatelimitCommitConfigResponse() (response *DescribeRatelimitCommitConfigResponse) {
	response = &DescribeRatelimitCommitConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeRatelimitCommitConfig
func (c *Client) DescribeRatelimitCommitConfig(request *DescribeRatelimitCommitConfigRequest) (response *DescribeRatelimitCommitConfigResponse, err error) {
	if request == nil {
		request = NewDescribeRatelimitCommitConfigRequest()
	}
	response = NewDescribeRatelimitCommitConfigResponse()
	err = c.Send(request, response)
	return
}

func NewResetMeshSidecarMonitorRequest() (request *ResetMeshSidecarMonitorRequest) {
	request = &ResetMeshSidecarMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ResetMeshSidecarMonitor")
	return
}

func NewResetMeshSidecarMonitorResponse() (response *ResetMeshSidecarMonitorResponse) {
	response = &ResetMeshSidecarMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置监控信息
func (c *Client) ResetMeshSidecarMonitor(request *ResetMeshSidecarMonitorRequest) (response *ResetMeshSidecarMonitorResponse, err error) {
	if request == nil {
		request = NewResetMeshSidecarMonitorRequest()
	}
	response = NewResetMeshSidecarMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskFlowGraphRequest() (request *DescribeTaskFlowGraphRequest) {
	request = &DescribeTaskFlowGraphRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskFlowGraph")
	return
}

func NewDescribeTaskFlowGraphResponse() (response *DescribeTaskFlowGraphResponse) {
	response = &DescribeTaskFlowGraphResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作流的任务拓扑图
func (c *Client) DescribeTaskFlowGraph(request *DescribeTaskFlowGraphRequest) (response *DescribeTaskFlowGraphResponse, err error) {
	if request == nil {
		request = NewDescribeTaskFlowGraphRequest()
	}
	response = NewDescribeTaskFlowGraphResponse()
	err = c.Send(request, response)
	return
}

func NewDisableCircuitBreakerRuleRequest() (request *DisableCircuitBreakerRuleRequest) {
	request = &DisableCircuitBreakerRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableCircuitBreakerRule")
	return
}

func NewDisableCircuitBreakerRuleResponse() (response *DisableCircuitBreakerRuleResponse) {
	response = &DisableCircuitBreakerRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用熔断规则
func (c *Client) DisableCircuitBreakerRule(request *DisableCircuitBreakerRuleRequest) (response *DisableCircuitBreakerRuleResponse, err error) {
	if request == nil {
		request = NewDisableCircuitBreakerRuleRequest()
	}
	response = NewDisableCircuitBreakerRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeployTsfModulesRequest() (request *DeployTsfModulesRequest) {
	request = &DeployTsfModulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeployTsfModules")
	return
}

func NewDeployTsfModulesResponse() (response *DeployTsfModulesResponse) {
	response = &DeployTsfModulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 模块部署
func (c *Client) DeployTsfModules(request *DeployTsfModulesRequest) (response *DeployTsfModulesResponse, err error) {
	if request == nil {
		request = NewDeployTsfModulesRequest()
	}
	response = NewDeployTsfModulesResponse()
	err = c.Send(request, response)
	return
}

func NewGetServiceStatisticsRequest() (request *GetServiceStatisticsRequest) {
	request = &GetServiceStatisticsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetServiceStatistics")
	return
}

func NewGetServiceStatisticsResponse() (response *GetServiceStatisticsResponse) {
	response = &GetServiceStatisticsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务运行状态
func (c *Client) GetServiceStatistics(request *GetServiceStatisticsRequest) (response *GetServiceStatisticsResponse, err error) {
	if request == nil {
		request = NewGetServiceStatisticsRequest()
	}
	response = NewGetServiceStatisticsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServerlessGroupsRequest() (request *DescribeServerlessGroupsRequest) {
	request = &DescribeServerlessGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeServerlessGroups")
	return
}

func NewDescribeServerlessGroupsResponse() (response *DescribeServerlessGroupsResponse) {
	response = &DescribeServerlessGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Serverless部署组列表
func (c *Client) DescribeServerlessGroups(request *DescribeServerlessGroupsRequest) (response *DescribeServerlessGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeServerlessGroupsRequest()
	}
	response = NewDescribeServerlessGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecordsRequest() (request *DescribeRecordsRequest) {
	request = &DescribeRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRecords")
	return
}

func NewDescribeRecordsResponse() (response *DescribeRecordsResponse) {
	response = &DescribeRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询操作记录列表.
func (c *Client) DescribeRecords(request *DescribeRecordsRequest) (response *DescribeRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeRecordsRequest()
	}
	response = NewDescribeRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMonitorStatisticsPolicyRequest() (request *CreateMonitorStatisticsPolicyRequest) {
	request = &CreateMonitorStatisticsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateMonitorStatisticsPolicy")
	return
}

func NewCreateMonitorStatisticsPolicyResponse() (response *CreateMonitorStatisticsPolicyResponse) {
	response = &CreateMonitorStatisticsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建监控统计策略
func (c *Client) CreateMonitorStatisticsPolicy(request *CreateMonitorStatisticsPolicyRequest) (response *CreateMonitorStatisticsPolicyResponse, err error) {
	if request == nil {
		request = NewCreateMonitorStatisticsPolicyRequest()
	}
	response = NewCreateMonitorStatisticsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroServiceListRequest() (request *DescribeMicroServiceListRequest) {
	request = &DescribeMicroServiceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroServiceList")
	return
}

func NewDescribeMicroServiceListResponse() (response *DescribeMicroServiceListResponse) {
	response = &DescribeMicroServiceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 云监控中获取tsf 微服务对象列表
func (c *Client) DescribeMicroServiceList(request *DescribeMicroServiceListRequest) (response *DescribeMicroServiceListResponse, err error) {
	if request == nil {
		request = NewDescribeMicroServiceListRequest()
	}
	response = NewDescribeMicroServiceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeMonitorRequest() (request *DescribeNodeMonitorRequest) {
	request = &DescribeNodeMonitorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeNodeMonitor")
	return
}

func NewDescribeNodeMonitorResponse() (response *DescribeNodeMonitorResponse) {
	response = &DescribeNodeMonitorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点监控数据
func (c *Client) DescribeNodeMonitor(request *DescribeNodeMonitorRequest) (response *DescribeNodeMonitorResponse, err error) {
	if request == nil {
		request = NewDescribeNodeMonitorRequest()
	}
	response = NewDescribeNodeMonitorResponse()
	err = c.Send(request, response)
	return
}

func NewDisableSidecarFilterRequest() (request *DisableSidecarFilterRequest) {
	request = &DisableSidecarFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableSidecarFilter")
	return
}

func NewDisableSidecarFilterResponse() (response *DisableSidecarFilterResponse) {
	response = &DisableSidecarFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用Sidecar过滤器
func (c *Client) DisableSidecarFilter(request *DisableSidecarFilterRequest) (response *DisableSidecarFilterResponse, err error) {
	if request == nil {
		request = NewDisableSidecarFilterRequest()
	}
	response = NewDisableSidecarFilterResponse()
	err = c.Send(request, response)
	return
}

func NewFindDeployModuleLogRequest() (request *FindDeployModuleLogRequest) {
	request = &FindDeployModuleLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "FindDeployModuleLog")
	return
}

func NewFindDeployModuleLogResponse() (response *FindDeployModuleLogResponse) {
	response = &FindDeployModuleLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询模块部署日志
func (c *Client) FindDeployModuleLog(request *FindDeployModuleLogRequest) (response *FindDeployModuleLogResponse, err error) {
	if request == nil {
		request = NewFindDeployModuleLogRequest()
	}
	response = NewFindDeployModuleLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFunctionStatusRequest() (request *DescribeFunctionStatusRequest) {
	request = &DescribeFunctionStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFunctionStatus")
	return
}

func NewDescribeFunctionStatusResponse() (response *DescribeFunctionStatusResponse) {
	response = &DescribeFunctionStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述功能状态
func (c *Client) DescribeFunctionStatus(request *DescribeFunctionStatusRequest) (response *DescribeFunctionStatusResponse, err error) {
	if request == nil {
		request = NewDescribeFunctionStatusRequest()
	}
	response = NewDescribeFunctionStatusResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopAvgTimeCostInterfacesRequest() (request *GetTopAvgTimeCostInterfacesRequest) {
	request = &GetTopAvgTimeCostInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTopAvgTimeCostInterfaces")
	return
}

func NewGetTopAvgTimeCostInterfacesResponse() (response *GetTopAvgTimeCostInterfacesResponse) {
	response = &GetTopAvgTimeCostInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TopN请求平均耗时接口列表
func (c *Client) GetTopAvgTimeCostInterfaces(request *GetTopAvgTimeCostInterfacesRequest) (response *GetTopAvgTimeCostInterfacesResponse, err error) {
	if request == nil {
		request = NewGetTopAvgTimeCostInterfacesRequest()
	}
	response = NewGetTopAvgTimeCostInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTsfRegionsRequest() (request *DescribeTsfRegionsRequest) {
	request = &DescribeTsfRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTsfRegions")
	return
}

func NewDescribeTsfRegionsResponse() (response *DescribeTsfRegionsResponse) {
	response = &DescribeTsfRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TSF地域列表
func (c *Client) DescribeTsfRegions(request *DescribeTsfRegionsRequest) (response *DescribeTsfRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeTsfRegionsRequest()
	}
	response = NewDescribeTsfRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTaskRequest() (request *CreateTaskRequest) {
	request = &CreateTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateTask")
	return
}

func NewCreateTaskResponse() (response *CreateTaskResponse) {
	response = &CreateTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建任务
func (c *Client) CreateTask(request *CreateTaskRequest) (response *CreateTaskResponse, err error) {
	if request == nil {
		request = NewCreateTaskRequest()
	}
	response = NewCreateTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApiGroupRequest() (request *DeleteApiGroupRequest) {
	request = &DeleteApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteApiGroup")
	return
}

func NewDeleteApiGroupResponse() (response *DeleteApiGroupResponse) {
	response = &DeleteApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Api分组
func (c *Client) DeleteApiGroup(request *DeleteApiGroupRequest) (response *DeleteApiGroupResponse, err error) {
	if request == nil {
		request = NewDeleteApiGroupRequest()
	}
	response = NewDeleteApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCloudMonitorGroupsRequest() (request *DescribeCloudMonitorGroupsRequest) {
	request = &DescribeCloudMonitorGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeCloudMonitorGroups")
	return
}

func NewDescribeCloudMonitorGroupsResponse() (response *DescribeCloudMonitorGroupsResponse) {
	response = &DescribeCloudMonitorGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取部署组列表，提供给云监控
func (c *Client) DescribeCloudMonitorGroups(request *DescribeCloudMonitorGroupsRequest) (response *DescribeCloudMonitorGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeCloudMonitorGroupsRequest()
	}
	response = NewDescribeCloudMonitorGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewStartContainerGroupRequest() (request *StartContainerGroupRequest) {
	request = &StartContainerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StartContainerGroup")
	return
}

func NewStartContainerGroupResponse() (response *StartContainerGroupResponse) {
	response = &StartContainerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动容器部署组
func (c *Client) StartContainerGroup(request *StartContainerGroupRequest) (response *StartContainerGroupResponse, err error) {
	if request == nil {
		request = NewStartContainerGroupRequest()
	}
	response = NewStartContainerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDscribeTasksRequest() (request *DscribeTasksRequest) {
	request = &DscribeTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DscribeTasks")
	return
}

func NewDscribeTasksResponse() (response *DscribeTasksResponse) {
	response = &DscribeTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取操作日志列表
func (c *Client) DscribeTasks(request *DscribeTasksRequest) (response *DscribeTasksResponse, err error) {
	if request == nil {
		request = NewDscribeTasksRequest()
	}
	response = NewDscribeTasksResponse()
	err = c.Send(request, response)
	return
}

func NewRevocationPublicConfigRequest() (request *RevocationPublicConfigRequest) {
	request = &RevocationPublicConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RevocationPublicConfig")
	return
}

func NewRevocationPublicConfigResponse() (response *RevocationPublicConfigResponse) {
	response = &RevocationPublicConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 撤回已发布的公共配置
func (c *Client) RevocationPublicConfig(request *RevocationPublicConfigRequest) (response *RevocationPublicConfigResponse, err error) {
	if request == nil {
		request = NewRevocationPublicConfigRequest()
	}
	response = NewRevocationPublicConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSimpleApplicationsRequest() (request *DescribeSimpleApplicationsRequest) {
	request = &DescribeSimpleApplicationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleApplications")
	return
}

func NewDescribeSimpleApplicationsResponse() (response *DescribeSimpleApplicationsResponse) {
	response = &DescribeSimpleApplicationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询简单应用列表
func (c *Client) DescribeSimpleApplications(request *DescribeSimpleApplicationsRequest) (response *DescribeSimpleApplicationsResponse, err error) {
	if request == nil {
		request = NewDescribeSimpleApplicationsRequest()
	}
	response = NewDescribeSimpleApplicationsResponse()
	err = c.Send(request, response)
	return
}

func NewPushMultiCloudUnitInfoRequest() (request *PushMultiCloudUnitInfoRequest) {
	request = &PushMultiCloudUnitInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "PushMultiCloudUnitInfo")
	return
}

func NewPushMultiCloudUnitInfoResponse() (response *PushMultiCloudUnitInfoResponse) {
	response = &PushMultiCloudUnitInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 超云推送架构配置到consul
func (c *Client) PushMultiCloudUnitInfo(request *PushMultiCloudUnitInfoRequest) (response *PushMultiCloudUnitInfoResponse, err error) {
	if request == nil {
		request = NewPushMultiCloudUnitInfoRequest()
	}
	response = NewPushMultiCloudUnitInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterInstancesRequest() (request *DescribeClusterInstancesRequest) {
	request = &DescribeClusterInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterInstances")
	return
}

func NewDescribeClusterInstancesResponse() (response *DescribeClusterInstancesResponse) {
	response = &DescribeClusterInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群实例
func (c *Client) DescribeClusterInstances(request *DescribeClusterInstancesRequest) (response *DescribeClusterInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterInstancesRequest()
	}
	response = NewDescribeClusterInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEnabledUnitRuleRequest() (request *DescribeEnabledUnitRuleRequest) {
	request = &DescribeEnabledUnitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeEnabledUnitRule")
	return
}

func NewDescribeEnabledUnitRuleResponse() (response *DescribeEnabledUnitRuleResponse) {
	response = &DescribeEnabledUnitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询生效的单元化规则
func (c *Client) DescribeEnabledUnitRule(request *DescribeEnabledUnitRuleRequest) (response *DescribeEnabledUnitRuleResponse, err error) {
	if request == nil {
		request = NewDescribeEnabledUnitRuleRequest()
	}
	response = NewDescribeEnabledUnitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewListScalableRuleRequest() (request *ListScalableRuleRequest) {
	request = &ListScalableRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListScalableRule")
	return
}

func NewListScalableRuleResponse() (response *ListScalableRuleResponse) {
	response = &ListScalableRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// list弹性扩缩容规则
func (c *Client) ListScalableRule(request *ListScalableRuleRequest) (response *ListScalableRuleResponse, err error) {
	if request == nil {
		request = NewListScalableRuleRequest()
	}
	response = NewListScalableRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayWeChatMiniProgramAppSecretRequest() (request *UpdateGatewayWeChatMiniProgramAppSecretRequest) {
	request = &UpdateGatewayWeChatMiniProgramAppSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayWeChatMiniProgramAppSecret")
	return
}

func NewUpdateGatewayWeChatMiniProgramAppSecretResponse() (response *UpdateGatewayWeChatMiniProgramAppSecretResponse) {
	response = &UpdateGatewayWeChatMiniProgramAppSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关微信小程序应用密钥参数
func (c *Client) UpdateGatewayWeChatMiniProgramAppSecret(request *UpdateGatewayWeChatMiniProgramAppSecretRequest) (response *UpdateGatewayWeChatMiniProgramAppSecretResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayWeChatMiniProgramAppSecretRequest()
	}
	response = NewUpdateGatewayWeChatMiniProgramAppSecretResponse()
	err = c.Send(request, response)
	return
}

func NewDisableFallbackRouteRequest() (request *DisableFallbackRouteRequest) {
	request = &DisableFallbackRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableFallbackRoute")
	return
}

func NewDisableFallbackRouteResponse() (response *DisableFallbackRouteResponse) {
	response = &DisableFallbackRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停用微服务路由保护策略
func (c *Client) DisableFallbackRoute(request *DisableFallbackRouteRequest) (response *DisableFallbackRouteResponse, err error) {
	if request == nil {
		request = NewDisableFallbackRouteRequest()
	}
	response = NewDisableFallbackRouteResponse()
	err = c.Send(request, response)
	return
}

func NewEnableRatelimitRequest() (request *EnableRatelimitRequest) {
	request = &EnableRatelimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableRatelimit")
	return
}

func NewEnableRatelimitResponse() (response *EnableRatelimitResponse) {
	response = &EnableRatelimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用限流
func (c *Client) EnableRatelimit(request *EnableRatelimitRequest) (response *EnableRatelimitResponse, err error) {
	if request == nil {
		request = NewEnableRatelimitRequest()
	}
	response = NewEnableRatelimitResponse()
	err = c.Send(request, response)
	return
}

func NewGetTraceServicesRequest() (request *GetTraceServicesRequest) {
	request = &GetTraceServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTraceServices")
	return
}

func NewGetTraceServicesResponse() (response *GetTraceServicesResponse) {
	response = &GetTraceServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用链服务列表
func (c *Client) GetTraceServices(request *GetTraceServicesRequest) (response *GetTraceServicesResponse, err error) {
	if request == nil {
		request = NewGetTraceServicesRequest()
	}
	response = NewGetTraceServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDraftApiGroupRequest() (request *DraftApiGroupRequest) {
	request = &DraftApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DraftApiGroup")
	return
}

func NewDraftApiGroupResponse() (response *DraftApiGroupResponse) {
	response = &DraftApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下线Api分组
func (c *Client) DraftApiGroup(request *DraftApiGroupRequest) (response *DraftApiGroupResponse, err error) {
	if request == nil {
		request = NewDraftApiGroupRequest()
	}
	response = NewDraftApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseApiGroupRequest() (request *ReleaseApiGroupRequest) {
	request = &ReleaseApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleaseApiGroup")
	return
}

func NewReleaseApiGroupResponse() (response *ReleaseApiGroupResponse) {
	response = &ReleaseApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布Api分组
func (c *Client) ReleaseApiGroup(request *ReleaseApiGroupRequest) (response *ReleaseApiGroupResponse, err error) {
	if request == nil {
		request = NewReleaseApiGroupRequest()
	}
	response = NewReleaseApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeployContainerGroupRequest() (request *DeployContainerGroupRequest) {
	request = &DeployContainerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeployContainerGroup")
	return
}

func NewDeployContainerGroupResponse() (response *DeployContainerGroupResponse) {
	response = &DeployContainerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署容器应用-更新
func (c *Client) DeployContainerGroup(request *DeployContainerGroupRequest) (response *DeployContainerGroupResponse, err error) {
	if request == nil {
		request = NewDeployContainerGroupRequest()
	}
	response = NewDeployContainerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewRetryTransactionRequest() (request *RetryTransactionRequest) {
	request = &RetryTransactionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RetryTransaction")
	return
}

func NewRetryTransactionResponse() (response *RetryTransactionResponse) {
	response = &RetryTransactionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 手动触发事务超时重试
func (c *Client) RetryTransaction(request *RetryTransactionRequest) (response *RetryTransactionResponse, err error) {
	if request == nil {
		request = NewRetryTransactionRequest()
	}
	response = NewRetryTransactionResponse()
	err = c.Send(request, response)
	return
}

func NewTestKuberneteNativeConnectionRequest() (request *TestKuberneteNativeConnectionRequest) {
	request = &TestKuberneteNativeConnectionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "TestKuberneteNativeConnection")
	return
}

func NewTestKuberneteNativeConnectionResponse() (response *TestKuberneteNativeConnectionResponse) {
	response = &TestKuberneteNativeConnectionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查原生集群链接
func (c *Client) TestKuberneteNativeConnection(request *TestKuberneteNativeConnectionRequest) (response *TestKuberneteNativeConnectionResponse, err error) {
	if request == nil {
		request = NewTestKuberneteNativeConnectionRequest()
	}
	response = NewTestKuberneteNativeConnectionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateScalableRuleWithDetailRespRequest() (request *CreateScalableRuleWithDetailRespRequest) {
	request = &CreateScalableRuleWithDetailRespRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateScalableRuleWithDetailResp")
	return
}

func NewCreateScalableRuleWithDetailRespResponse() (response *CreateScalableRuleWithDetailRespResponse) {
	response = &CreateScalableRuleWithDetailRespResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建弹性扩缩容规则返回id
func (c *Client) CreateScalableRuleWithDetailResp(request *CreateScalableRuleWithDetailRespRequest) (response *CreateScalableRuleWithDetailRespResponse, err error) {
	if request == nil {
		request = NewCreateScalableRuleWithDetailRespRequest()
	}
	response = NewCreateScalableRuleWithDetailRespResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerGroupAttributeRequest() (request *DescribeContainerGroupAttributeRequest) {
	request = &DescribeContainerGroupAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupAttribute")
	return
}

func NewDescribeContainerGroupAttributeResponse() (response *DescribeContainerGroupAttributeResponse) {
	response = &DescribeContainerGroupAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取部署组其他字段-用于前端并发调用
func (c *Client) DescribeContainerGroupAttribute(request *DescribeContainerGroupAttributeRequest) (response *DescribeContainerGroupAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeContainerGroupAttributeRequest()
	}
	response = NewDescribeContainerGroupAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContainerReplicasRequest() (request *ModifyContainerReplicasRequest) {
	request = &ModifyContainerReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerReplicas")
	return
}

func NewModifyContainerReplicasResponse() (response *ModifyContainerReplicasResponse) {
	response = &ModifyContainerReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改容器部署组实例数
func (c *Client) ModifyContainerReplicas(request *ModifyContainerReplicasRequest) (response *ModifyContainerReplicasResponse, err error) {
	if request == nil {
		request = NewModifyContainerReplicasRequest()
	}
	response = NewModifyContainerReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAbnormalMetricsConfigRequest() (request *CreateAbnormalMetricsConfigRequest) {
	request = &CreateAbnormalMetricsConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateAbnormalMetricsConfig")
	return
}

func NewCreateAbnormalMetricsConfigResponse() (response *CreateAbnormalMetricsConfigResponse) {
	response = &CreateAbnormalMetricsConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建异常指标配置
func (c *Client) CreateAbnormalMetricsConfig(request *CreateAbnormalMetricsConfigRequest) (response *CreateAbnormalMetricsConfigResponse, err error) {
	if request == nil {
		request = NewCreateAbnormalMetricsConfigRequest()
	}
	response = NewCreateAbnormalMetricsConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateReleaseRequest() (request *CreateReleaseRequest) {
	request = &CreateReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateRelease")
	return
}

func NewCreateReleaseResponse() (response *CreateReleaseResponse) {
	response = &CreateReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建发布单
func (c *Client) CreateRelease(request *CreateReleaseRequest) (response *CreateReleaseResponse, err error) {
	if request == nil {
		request = NewCreateReleaseRequest()
	}
	response = NewCreateReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewGetDownloadInfoRequest() (request *GetDownloadInfoRequest) {
	request = &GetDownloadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetDownloadInfo")
	return
}

func NewGetDownloadInfoResponse() (response *GetDownloadInfoResponse) {
	response = &GetDownloadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF上传的程序包存放在腾讯云对象存储（COS）中，通过该API可以获取从COS下载程序包需要的信息，包括包所在的桶、存储路径、鉴权信息等，之后使用COS API（或SDK）进行下载。
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
func (c *Client) GetDownloadInfo(request *GetDownloadInfoRequest) (response *GetDownloadInfoResponse, err error) {
	if request == nil {
		request = NewGetDownloadInfoRequest()
	}
	response = NewGetDownloadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCkafkaRequest() (request *CreateCkafkaRequest) {
	request = &CreateCkafkaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateCkafka")
	return
}

func NewCreateCkafkaResponse() (response *CreateCkafkaResponse) {
	response = &CreateCkafkaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建ckafka
func (c *Client) CreateCkafka(request *CreateCkafkaRequest) (response *CreateCkafkaResponse, err error) {
	if request == nil {
		request = NewCreateCkafkaRequest()
	}
	response = NewCreateCkafkaResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePkgsRequest() (request *DeletePkgsRequest) {
	request = &DeletePkgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeletePkgs")
	return
}

func NewDeletePkgsResponse() (response *DeletePkgsResponse) {
	response = &DeletePkgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从软件仓库批量删除程序包。
// 一次最多支持删除1000个包，数量超过1000，返回UpperDeleteLimit错误。
func (c *Client) DeletePkgs(request *DeletePkgsRequest) (response *DeletePkgsResponse, err error) {
	if request == nil {
		request = NewDeletePkgsRequest()
	}
	response = NewDeletePkgsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApplicationRequest() (request *CreateApplicationRequest) {
	request = &CreateApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateApplication")
	return
}

func NewCreateApplicationResponse() (response *CreateApplicationResponse) {
	response = &CreateApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建应用
func (c *Client) CreateApplication(request *CreateApplicationRequest) (response *CreateApplicationResponse, err error) {
	if request == nil {
		request = NewCreateApplicationRequest()
	}
	response = NewCreateApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMicroserviceApiRequest() (request *CreateMicroserviceApiRequest) {
	request = &CreateMicroserviceApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateMicroserviceApi")
	return
}

func NewCreateMicroserviceApiResponse() (response *CreateMicroserviceApiResponse) {
	response = &CreateMicroserviceApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增微服务API
func (c *Client) CreateMicroserviceApi(request *CreateMicroserviceApiRequest) (response *CreateMicroserviceApiResponse, err error) {
	if request == nil {
		request = NewCreateMicroserviceApiRequest()
	}
	response = NewCreateMicroserviceApiResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMachinesRequest() (request *ModifyMachinesRequest) {
	request = &ModifyMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyMachines")
	return
}

func NewModifyMachinesResponse() (response *ModifyMachinesResponse) {
	response = &ModifyMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机器列表
func (c *Client) ModifyMachines(request *ModifyMachinesRequest) (response *ModifyMachinesResponse, err error) {
	if request == nil {
		request = NewModifyMachinesRequest()
	}
	response = NewModifyMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayPluginTypesRequest() (request *DescribeGatewayPluginTypesRequest) {
	request = &DescribeGatewayPluginTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayPluginTypes")
	return
}

func NewDescribeGatewayPluginTypesResponse() (response *DescribeGatewayPluginTypesResponse) {
	response = &DescribeGatewayPluginTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关插件支持的类型列表
func (c *Client) DescribeGatewayPluginTypes(request *DescribeGatewayPluginTypesRequest) (response *DescribeGatewayPluginTypesResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayPluginTypesRequest()
	}
	response = NewDescribeGatewayPluginTypesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePermissionCategoriesRequest() (request *DescribePermissionCategoriesRequest) {
	request = &DescribePermissionCategoriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePermissionCategories")
	return
}

func NewDescribePermissionCategoriesResponse() (response *DescribePermissionCategoriesResponse) {
	response = &DescribePermissionCategoriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询权限组列表
func (c *Client) DescribePermissionCategories(request *DescribePermissionCategoriesRequest) (response *DescribePermissionCategoriesResponse, err error) {
	if request == nil {
		request = NewDescribePermissionCategoriesRequest()
	}
	response = NewDescribePermissionCategoriesResponse()
	err = c.Send(request, response)
	return
}

func NewListScalableTasksRequest() (request *ListScalableTasksRequest) {
	request = &ListScalableTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListScalableTasks")
	return
}

func NewListScalableTasksResponse() (response *ListScalableTasksResponse) {
	response = &ListScalableTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性扩缩容任务
func (c *Client) ListScalableTasks(request *ListScalableTasksRequest) (response *ListScalableTasksResponse, err error) {
	if request == nil {
		request = NewListScalableTasksRequest()
	}
	response = NewListScalableTasksResponse()
	err = c.Send(request, response)
	return
}

func NewReleasePublicConfigRequest() (request *ReleasePublicConfigRequest) {
	request = &ReleasePublicConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleasePublicConfig")
	return
}

func NewReleasePublicConfigResponse() (response *ReleasePublicConfigResponse) {
	response = &ReleasePublicConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布公共配置
func (c *Client) ReleasePublicConfig(request *ReleasePublicConfigRequest) (response *ReleasePublicConfigResponse, err error) {
	if request == nil {
		request = NewReleasePublicConfigRequest()
	}
	response = NewReleasePublicConfigResponse()
	err = c.Send(request, response)
	return
}

func NewStartServerlessGroupRequest() (request *StartServerlessGroupRequest) {
	request = &StartServerlessGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StartServerlessGroup")
	return
}

func NewStartServerlessGroupResponse() (response *StartServerlessGroupResponse) {
	response = &StartServerlessGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) StartServerlessGroup(request *StartServerlessGroupRequest) (response *StartServerlessGroupResponse, err error) {
	if request == nil {
		request = NewStartServerlessGroupRequest()
	}
	response = NewStartServerlessGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupBusinessLogConfigsRequest() (request *DescribeGroupBusinessLogConfigsRequest) {
	request = &DescribeGroupBusinessLogConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupBusinessLogConfigs")
	return
}

func NewDescribeGroupBusinessLogConfigsResponse() (response *DescribeGroupBusinessLogConfigsResponse) {
	response = &DescribeGroupBusinessLogConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询分组管理日志配置列表
func (c *Client) DescribeGroupBusinessLogConfigs(request *DescribeGroupBusinessLogConfigsRequest) (response *DescribeGroupBusinessLogConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeGroupBusinessLogConfigsRequest()
	}
	response = NewDescribeGroupBusinessLogConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFilebeatConfigPageRequest() (request *DescribeFilebeatConfigPageRequest) {
	request = &DescribeFilebeatConfigPageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFilebeatConfigPage")
	return
}

func NewDescribeFilebeatConfigPageResponse() (response *DescribeFilebeatConfigPageResponse) {
	response = &DescribeFilebeatConfigPageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询filebeat config配置
func (c *Client) DescribeFilebeatConfigPage(request *DescribeFilebeatConfigPageRequest) (response *DescribeFilebeatConfigPageResponse, err error) {
	if request == nil {
		request = NewDescribeFilebeatConfigPageRequest()
	}
	response = NewDescribeFilebeatConfigPageResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApiRateLimitRulesRequest() (request *UpdateApiRateLimitRulesRequest) {
	request = &UpdateApiRateLimitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiRateLimitRules")
	return
}

func NewUpdateApiRateLimitRulesResponse() (response *UpdateApiRateLimitRulesResponse) {
	response = &UpdateApiRateLimitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量更新API限流规则
func (c *Client) UpdateApiRateLimitRules(request *UpdateApiRateLimitRulesRequest) (response *UpdateApiRateLimitRulesResponse, err error) {
	if request == nil {
		request = NewUpdateApiRateLimitRulesRequest()
	}
	response = NewUpdateApiRateLimitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDownloadInfoRequest() (request *DescribeDownloadInfoRequest) {
	request = &DescribeDownloadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeDownloadInfo")
	return
}

func NewDescribeDownloadInfoResponse() (response *DescribeDownloadInfoResponse) {
	response = &DescribeDownloadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF上传的程序包存放在腾讯云对象存储（COS）中，通过该API可以获取从COS下载程序包需要的信息，包括包所在的桶、存储路径、鉴权信息等，之后使用COS API（或SDK）进行下载。
// COS相关文档请查阅：https://cloud.tencent.com/document/product/436
func (c *Client) DescribeDownloadInfo(request *DescribeDownloadInfoRequest) (response *DescribeDownloadInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDownloadInfoRequest()
	}
	response = NewDescribeDownloadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespaceAffinitiesRequest() (request *DescribeNamespaceAffinitiesRequest) {
	request = &DescribeNamespaceAffinitiesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaceAffinities")
	return
}

func NewDescribeNamespaceAffinitiesResponse() (response *DescribeNamespaceAffinitiesResponse) {
	response = &DescribeNamespaceAffinitiesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询命名空间就近访问策略列表
func (c *Client) DescribeNamespaceAffinities(request *DescribeNamespaceAffinitiesRequest) (response *DescribeNamespaceAffinitiesResponse, err error) {
	if request == nil {
		request = NewDescribeNamespaceAffinitiesRequest()
	}
	response = NewDescribeNamespaceAffinitiesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteRequest() (request *DescribeRouteRequest) {
	request = &DescribeRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRoute")
	return
}

func NewDescribeRouteResponse() (response *DescribeRouteResponse) {
	response = &DescribeRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由详情
func (c *Client) DescribeRoute(request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
	if request == nil {
		request = NewDescribeRouteRequest()
	}
	response = NewDescribeRouteResponse()
	err = c.Send(request, response)
	return
}

func NewSearchMeshLogRequest() (request *SearchMeshLogRequest) {
	request = &SearchMeshLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchMeshLog")
	return
}

func NewSearchMeshLogResponse() (response *SearchMeshLogResponse) {
	response = &SearchMeshLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// mesh日志搜索
func (c *Client) SearchMeshLog(request *SearchMeshLogRequest) (response *SearchMeshLogResponse, err error) {
	if request == nil {
		request = NewSearchMeshLogRequest()
	}
	response = NewSearchMeshLogResponse()
	err = c.Send(request, response)
	return
}

func NewShrinkGroupRequest() (request *ShrinkGroupRequest) {
	request = &ShrinkGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ShrinkGroup")
	return
}

func NewShrinkGroupResponse() (response *ShrinkGroupResponse) {
	response = &ShrinkGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下线部署组所有机器实例
func (c *Client) ShrinkGroup(request *ShrinkGroupRequest) (response *ShrinkGroupResponse, err error) {
	if request == nil {
		request = NewShrinkGroupRequest()
	}
	response = NewShrinkGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayWeChatMiniProgramLoginPluginRequest() (request *UpdateGatewayWeChatMiniProgramLoginPluginRequest) {
	request = &UpdateGatewayWeChatMiniProgramLoginPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayWeChatMiniProgramLoginPlugin")
	return
}

func NewUpdateGatewayWeChatMiniProgramLoginPluginResponse() (response *UpdateGatewayWeChatMiniProgramLoginPluginResponse) {
	response = &UpdateGatewayWeChatMiniProgramLoginPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改微信小程序登录插件
func (c *Client) UpdateGatewayWeChatMiniProgramLoginPlugin(request *UpdateGatewayWeChatMiniProgramLoginPluginRequest) (response *UpdateGatewayWeChatMiniProgramLoginPluginResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayWeChatMiniProgramLoginPluginRequest()
	}
	response = NewUpdateGatewayWeChatMiniProgramLoginPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayInstancesRequest() (request *DescribeGatewayInstancesRequest) {
	request = &DescribeGatewayInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayInstances")
	return
}

func NewDescribeGatewayInstancesResponse() (response *DescribeGatewayInstancesResponse) {
	response = &DescribeGatewayInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关实体列表
func (c *Client) DescribeGatewayInstances(request *DescribeGatewayInstancesRequest) (response *DescribeGatewayInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayInstancesRequest()
	}
	response = NewDescribeGatewayInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDisableRouteRequest() (request *DisableRouteRequest) {
	request = &DisableRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableRoute")
	return
}

func NewDisableRouteResponse() (response *DisableRouteResponse) {
	response = &DisableRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停用服务路由
func (c *Client) DisableRoute(request *DisableRouteRequest) (response *DisableRouteResponse, err error) {
	if request == nil {
		request = NewDisableRouteRequest()
	}
	response = NewDisableRouteResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductHelpRequest() (request *CreateProductHelpRequest) {
	request = &CreateProductHelpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateProductHelp")
	return
}

func NewCreateProductHelpResponse() (response *CreateProductHelpResponse) {
	response = &CreateProductHelpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建产品帮助
func (c *Client) CreateProductHelp(request *CreateProductHelpRequest) (response *CreateProductHelpResponse, err error) {
	if request == nil {
		request = NewCreateProductHelpRequest()
	}
	response = NewCreateProductHelpResponse()
	err = c.Send(request, response)
	return
}

func NewImageGetRepositoryListRequest() (request *ImageGetRepositoryListRequest) {
	request = &ImageGetRepositoryListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ImageGetRepositoryList")
	return
}

func NewImageGetRepositoryListResponse() (response *ImageGetRepositoryListResponse) {
	response = &ImageGetRepositoryListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库列表
func (c *Client) ImageGetRepositoryList(request *ImageGetRepositoryListRequest) (response *ImageGetRepositoryListResponse, err error) {
	if request == nil {
		request = NewImageGetRepositoryListRequest()
	}
	response = NewImageGetRepositoryListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceNamespaceRequest() (request *ModifyInstanceNamespaceRequest) {
	request = &ModifyInstanceNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyInstanceNamespace")
	return
}

func NewModifyInstanceNamespaceResponse() (response *ModifyInstanceNamespaceResponse) {
	response = &ModifyInstanceNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机器命名空间
func (c *Client) ModifyInstanceNamespace(request *ModifyInstanceNamespaceRequest) (response *ModifyInstanceNamespaceResponse, err error) {
	if request == nil {
		request = NewModifyInstanceNamespaceRequest()
	}
	response = NewModifyInstanceNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRoutesRequest() (request *DescribeRoutesRequest) {
	request = &DescribeRoutesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRoutes")
	return
}

func NewDescribeRoutesResponse() (response *DescribeRoutesResponse) {
	response = &DescribeRoutesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由列表
func (c *Client) DescribeRoutes(request *DescribeRoutesRequest) (response *DescribeRoutesResponse, err error) {
	if request == nil {
		request = NewDescribeRoutesRequest()
	}
	response = NewDescribeRoutesResponse()
	err = c.Send(request, response)
	return
}

func NewDisableTaskRequest() (request *DisableTaskRequest) {
	request = &DisableTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableTask")
	return
}

func NewDisableTaskResponse() (response *DisableTaskResponse) {
	response = &DisableTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停用任务
func (c *Client) DisableTask(request *DisableTaskRequest) (response *DisableTaskResponse, err error) {
	if request == nil {
		request = NewDisableTaskRequest()
	}
	response = NewDisableTaskResponse()
	err = c.Send(request, response)
	return
}

func NewGetMetricsRequest() (request *GetMetricsRequest) {
	request = &GetMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetMetrics")
	return
}

func NewGetMetricsResponse() (response *GetMetricsResponse) {
	response = &GetMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取监控端点数据
func (c *Client) GetMetrics(request *GetMetricsRequest) (response *GetMetricsResponse, err error) {
	if request == nil {
		request = NewGetMetricsRequest()
	}
	response = NewGetMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOverviewEventRequest() (request *DescribeOverviewEventRequest) {
	request = &DescribeOverviewEventRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeOverviewEvent")
	return
}

func NewDescribeOverviewEventResponse() (response *DescribeOverviewEventResponse) {
	response = &DescribeOverviewEventResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 事件中心概览
func (c *Client) DescribeOverviewEvent(request *DescribeOverviewEventRequest) (response *DescribeOverviewEventResponse, err error) {
	if request == nil {
		request = NewDescribeOverviewEventRequest()
	}
	response = NewDescribeOverviewEventResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeViewsRequest() (request *DescribeViewsRequest) {
	request = &DescribeViewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeViews")
	return
}

func NewDescribeViewsResponse() (response *DescribeViewsResponse) {
	response = &DescribeViewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有视图列表
func (c *Client) DescribeViews(request *DescribeViewsRequest) (response *DescribeViewsResponse, err error) {
	if request == nil {
		request = NewDescribeViewsRequest()
	}
	response = NewDescribeViewsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateSidecarFilterRequest() (request *UpdateSidecarFilterRequest) {
	request = &UpdateSidecarFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateSidecarFilter")
	return
}

func NewUpdateSidecarFilterResponse() (response *UpdateSidecarFilterResponse) {
	response = &UpdateSidecarFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Sidecar过滤器
func (c *Client) UpdateSidecarFilter(request *UpdateSidecarFilterRequest) (response *UpdateSidecarFilterResponse, err error) {
	if request == nil {
		request = NewUpdateSidecarFilterRequest()
	}
	response = NewUpdateSidecarFilterResponse()
	err = c.Send(request, response)
	return
}

func NewGetOssTraceSpansRequest() (request *GetOssTraceSpansRequest) {
	request = &GetOssTraceSpansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetOssTraceSpans")
	return
}

func NewGetOssTraceSpansResponse() (response *GetOssTraceSpansResponse) {
	response = &GetOssTraceSpansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用TraceId查询调用链Span
func (c *Client) GetOssTraceSpans(request *GetOssTraceSpansRequest) (response *GetOssTraceSpansResponse, err error) {
	if request == nil {
		request = NewGetOssTraceSpansRequest()
	}
	response = NewGetOssTraceSpansResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMonitorStatisticsPolicyRequest() (request *ModifyMonitorStatisticsPolicyRequest) {
	request = &ModifyMonitorStatisticsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyMonitorStatisticsPolicy")
	return
}

func NewModifyMonitorStatisticsPolicyResponse() (response *ModifyMonitorStatisticsPolicyResponse) {
	response = &ModifyMonitorStatisticsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改监控统计策略
func (c *Client) ModifyMonitorStatisticsPolicy(request *ModifyMonitorStatisticsPolicyRequest) (response *ModifyMonitorStatisticsPolicyResponse, err error) {
	if request == nil {
		request = NewModifyMonitorStatisticsPolicyRequest()
	}
	response = NewModifyMonitorStatisticsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewSetEnvoyCloseDebugRequest() (request *SetEnvoyCloseDebugRequest) {
	request = &SetEnvoyCloseDebugRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SetEnvoyCloseDebug")
	return
}

func NewSetEnvoyCloseDebugResponse() (response *SetEnvoyCloseDebugResponse) {
	response = &SetEnvoyCloseDebugResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭envoy debug 级别日志
func (c *Client) SetEnvoyCloseDebug(request *SetEnvoyCloseDebugRequest) (response *SetEnvoyCloseDebugResponse, err error) {
	if request == nil {
		request = NewSetEnvoyCloseDebugRequest()
	}
	response = NewSetEnvoyCloseDebugResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCircuitBreakerRuleRequest() (request *DeleteCircuitBreakerRuleRequest) {
	request = &DeleteCircuitBreakerRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteCircuitBreakerRule")
	return
}

func NewDeleteCircuitBreakerRuleResponse() (response *DeleteCircuitBreakerRuleResponse) {
	response = &DeleteCircuitBreakerRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除熔断规则
func (c *Client) DeleteCircuitBreakerRule(request *DeleteCircuitBreakerRuleRequest) (response *DeleteCircuitBreakerRuleResponse, err error) {
	if request == nil {
		request = NewDeleteCircuitBreakerRuleRequest()
	}
	response = NewDeleteCircuitBreakerRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteGroupsRequest() (request *DeleteGroupsRequest) {
	request = &DeleteGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteGroups")
	return
}

func NewDeleteGroupsResponse() (response *DeleteGroupsResponse) {
	response = &DeleteGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除虚拟机部署组
func (c *Client) DeleteGroups(request *DeleteGroupsRequest) (response *DeleteGroupsResponse, err error) {
	if request == nil {
		request = NewDeleteGroupsRequest()
	}
	response = NewDeleteGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayInstanceProgramsRequest() (request *DescribeGatewayInstanceProgramsRequest) {
	request = &DescribeGatewayInstanceProgramsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayInstancePrograms")
	return
}

func NewDescribeGatewayInstanceProgramsResponse() (response *DescribeGatewayInstanceProgramsResponse) {
	response = &DescribeGatewayInstanceProgramsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据网关实体ID查询其对应绑定的数据集
func (c *Client) DescribeGatewayInstancePrograms(request *DescribeGatewayInstanceProgramsRequest) (response *DescribeGatewayInstanceProgramsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayInstanceProgramsRequest()
	}
	response = NewDescribeGatewayInstanceProgramsResponse()
	err = c.Send(request, response)
	return
}

func NewAddInstanceRequest() (request *AddInstanceRequest) {
	request = &AddInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AddInstance")
	return
}

func NewAddInstanceResponse() (response *AddInstanceResponse) {
	response = &AddInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加机器节点至TSF集群
func (c *Client) AddInstance(request *AddInstanceRequest) (response *AddInstanceResponse, err error) {
	if request == nil {
		request = NewAddInstanceRequest()
	}
	response = NewAddInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerEnvRequest() (request *DescribeContainerEnvRequest) {
	request = &DescribeContainerEnvRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerEnv")
	return
}

func NewDescribeContainerEnvResponse() (response *DescribeContainerEnvResponse) {
	response = &DescribeContainerEnvResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取容器部署组环境变量
func (c *Client) DescribeContainerEnv(request *DescribeContainerEnvRequest) (response *DescribeContainerEnvResponse, err error) {
	if request == nil {
		request = NewDescribeContainerEnvRequest()
	}
	response = NewDescribeContainerEnvResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayInstancesByMultiCloudRequest() (request *DescribeGatewayInstancesByMultiCloudRequest) {
	request = &DescribeGatewayInstancesByMultiCloudRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayInstancesByMultiCloud")
	return
}

func NewDescribeGatewayInstancesByMultiCloudResponse() (response *DescribeGatewayInstancesByMultiCloudResponse) {
	response = &DescribeGatewayInstancesByMultiCloudResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关实体列表(超云)
// 增加返回值:
// MultiCloudStatus 是否禁用绑定
// MultiCloudStatusMsg 禁用提示
func (c *Client) DescribeGatewayInstancesByMultiCloud(request *DescribeGatewayInstancesByMultiCloudRequest) (response *DescribeGatewayInstancesByMultiCloudResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayInstancesByMultiCloudRequest()
	}
	response = NewDescribeGatewayInstancesByMultiCloudResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasePipelineRunsRequest() (request *DescribeReleasePipelineRunsRequest) {
	request = &DescribeReleasePipelineRunsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasePipelineRuns")
	return
}

func NewDescribeReleasePipelineRunsResponse() (response *DescribeReleasePipelineRunsResponse) {
	response = &DescribeReleasePipelineRunsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布单执行流水列表
func (c *Client) DescribeReleasePipelineRuns(request *DescribeReleasePipelineRunsRequest) (response *DescribeReleasePipelineRunsResponse, err error) {
	if request == nil {
		request = NewDescribeReleasePipelineRunsRequest()
	}
	response = NewDescribeReleasePipelineRunsResponse()
	err = c.Send(request, response)
	return
}

func NewRunMsApiRequest() (request *RunMsApiRequest) {
	request = &RunMsApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RunMsApi")
	return
}

func NewRunMsApiResponse() (response *RunMsApiResponse) {
	response = &RunMsApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调试API
func (c *Client) RunMsApi(request *RunMsApiRequest) (response *RunMsApiResponse, err error) {
	if request == nil {
		request = NewRunMsApiRequest()
	}
	response = NewRunMsApiResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicConfigReleaseLogsRequest() (request *DescribePublicConfigReleaseLogsRequest) {
	request = &DescribePublicConfigReleaseLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleaseLogs")
	return
}

func NewDescribePublicConfigReleaseLogsResponse() (response *DescribePublicConfigReleaseLogsResponse) {
	response = &DescribePublicConfigReleaseLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共配置发布历史
func (c *Client) DescribePublicConfigReleaseLogs(request *DescribePublicConfigReleaseLogsRequest) (response *DescribePublicConfigReleaseLogsResponse, err error) {
	if request == nil {
		request = NewDescribePublicConfigReleaseLogsRequest()
	}
	response = NewDescribePublicConfigReleaseLogsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLaneGroupRequest() (request *CreateLaneGroupRequest) {
	request = &CreateLaneGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateLaneGroup")
	return
}

func NewCreateLaneGroupResponse() (response *CreateLaneGroupResponse) {
	response = &CreateLaneGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增泳道部署组
func (c *Client) CreateLaneGroup(request *CreateLaneGroupRequest) (response *CreateLaneGroupResponse, err error) {
	if request == nil {
		request = NewCreateLaneGroupRequest()
	}
	response = NewCreateLaneGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateServiceInstanceRequest() (request *CreateServiceInstanceRequest) {
	request = &CreateServiceInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateServiceInstance")
	return
}

func NewCreateServiceInstanceResponse() (response *CreateServiceInstanceResponse) {
	response = &CreateServiceInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加模块实例节点
func (c *Client) CreateServiceInstance(request *CreateServiceInstanceRequest) (response *CreateServiceInstanceResponse, err error) {
	if request == nil {
		request = NewCreateServiceInstanceRequest()
	}
	response = NewCreateServiceInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTaskFlowRequest() (request *CreateTaskFlowRequest) {
	request = &CreateTaskFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateTaskFlow")
	return
}

func NewCreateTaskFlowResponse() (response *CreateTaskFlowResponse) {
	response = &CreateTaskFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建工作流
func (c *Client) CreateTaskFlow(request *CreateTaskFlowRequest) (response *CreateTaskFlowResponse, err error) {
	if request == nil {
		request = NewCreateTaskFlowRequest()
	}
	response = NewCreateTaskFlowResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskDetailRequest() (request *DescribeTaskDetailRequest) {
	request = &DescribeTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskDetail")
	return
}

func NewDescribeTaskDetailResponse() (response *DescribeTaskDetailResponse) {
	response = &DescribeTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务详情
func (c *Client) DescribeTaskDetail(request *DescribeTaskDetailRequest) (response *DescribeTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeTaskDetailRequest()
	}
	response = NewDescribeTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopFailureRateServicesRequest() (request *GetTopFailureRateServicesRequest) {
	request = &GetTopFailureRateServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTopFailureRateServices")
	return
}

func NewGetTopFailureRateServicesResponse() (response *GetTopFailureRateServicesResponse) {
	response = &GetTopFailureRateServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TopN请求失败率服务列表
func (c *Client) GetTopFailureRateServices(request *GetTopFailureRateServicesRequest) (response *GetTopFailureRateServicesResponse, err error) {
	if request == nil {
		request = NewGetTopFailureRateServicesRequest()
	}
	response = NewGetTopFailureRateServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGlobalGroupContextUniqueRestrictRequest() (request *DescribeGlobalGroupContextUniqueRestrictRequest) {
	request = &DescribeGlobalGroupContextUniqueRestrictRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGlobalGroupContextUniqueRestrict")
	return
}

func NewDescribeGlobalGroupContextUniqueRestrictResponse() (response *DescribeGlobalGroupContextUniqueRestrictResponse) {
	response = &DescribeGlobalGroupContextUniqueRestrictResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有微服务网关下分组context是否唯一的全局限制
func (c *Client) DescribeGlobalGroupContextUniqueRestrict(request *DescribeGlobalGroupContextUniqueRestrictRequest) (response *DescribeGlobalGroupContextUniqueRestrictResponse, err error) {
	if request == nil {
		request = NewDescribeGlobalGroupContextUniqueRestrictRequest()
	}
	response = NewDescribeGlobalGroupContextUniqueRestrictResponse()
	err = c.Send(request, response)
	return
}

func NewGetContainGroupDeployInfoRequest() (request *GetContainGroupDeployInfoRequest) {
	request = &GetContainGroupDeployInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetContainGroupDeployInfo")
	return
}

func NewGetContainGroupDeployInfoResponse() (response *GetContainGroupDeployInfoResponse) {
	response = &GetContainGroupDeployInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取部署组详情
func (c *Client) GetContainGroupDeployInfo(request *GetContainGroupDeployInfoRequest) (response *GetContainGroupDeployInfoResponse, err error) {
	if request == nil {
		request = NewGetContainGroupDeployInfoRequest()
	}
	response = NewGetContainGroupDeployInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLaneRuleRequest() (request *DescribeLaneRuleRequest) {
	request = &DescribeLaneRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLaneRule")
	return
}

func NewDescribeLaneRuleResponse() (response *DescribeLaneRuleResponse) {
	response = &DescribeLaneRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单条泳道规则
func (c *Client) DescribeLaneRule(request *DescribeLaneRuleRequest) (response *DescribeLaneRuleResponse, err error) {
	if request == nil {
		request = NewDescribeLaneRuleRequest()
	}
	response = NewDescribeLaneRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeVisitAccountRequest() (request *DescribeVisitAccountRequest) {
	request = &DescribeVisitAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeVisitAccount")
	return
}

func NewDescribeVisitAccountResponse() (response *DescribeVisitAccountResponse) {
	response = &DescribeVisitAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询体验账号
func (c *Client) DescribeVisitAccount(request *DescribeVisitAccountRequest) (response *DescribeVisitAccountResponse, err error) {
	if request == nil {
		request = NewDescribeVisitAccountRequest()
	}
	response = NewDescribeVisitAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMeshSidecarStatusRequest() (request *DescribeMeshSidecarStatusRequest) {
	request = &DescribeMeshSidecarStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMeshSidecarStatus")
	return
}

func NewDescribeMeshSidecarStatusResponse() (response *DescribeMeshSidecarStatusResponse) {
	response = &DescribeMeshSidecarStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询sidecar状态
func (c *Client) DescribeMeshSidecarStatus(request *DescribeMeshSidecarStatusRequest) (response *DescribeMeshSidecarStatusResponse, err error) {
	if request == nil {
		request = NewDescribeMeshSidecarStatusRequest()
	}
	response = NewDescribeMeshSidecarStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationBusinessLogConfigRequest() (request *DescribeApplicationBusinessLogConfigRequest) {
	request = &DescribeApplicationBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationBusinessLogConfig")
	return
}

func NewDescribeApplicationBusinessLogConfigResponse() (response *DescribeApplicationBusinessLogConfigResponse) {
	response = &DescribeApplicationBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用关联日志配置项信息
func (c *Client) DescribeApplicationBusinessLogConfig(request *DescribeApplicationBusinessLogConfigRequest) (response *DescribeApplicationBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationBusinessLogConfigRequest()
	}
	response = NewDescribeApplicationBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLaneRequest() (request *DeleteLaneRequest) {
	request = &DeleteLaneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteLane")
	return
}

func NewDeleteLaneResponse() (response *DeleteLaneResponse) {
	response = &DeleteLaneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除泳道
func (c *Client) DeleteLane(request *DeleteLaneRequest) (response *DeleteLaneResponse, err error) {
	if request == nil {
		request = NewDeleteLaneRequest()
	}
	response = NewDeleteLaneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceNamesRequest() (request *DescribeServiceNamesRequest) {
	request = &DescribeServiceNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeServiceNames")
	return
}

func NewDescribeServiceNamesResponse() (response *DescribeServiceNamesResponse) {
	response = &DescribeServiceNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取appid 的微服务名列表
func (c *Client) DescribeServiceNames(request *DescribeServiceNamesRequest) (response *DescribeServiceNamesResponse, err error) {
	if request == nil {
		request = NewDescribeServiceNamesRequest()
	}
	response = NewDescribeServiceNamesResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseBusinessLogConfigRequest() (request *ReleaseBusinessLogConfigRequest) {
	request = &ReleaseBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleaseBusinessLogConfig")
	return
}

func NewReleaseBusinessLogConfigResponse() (response *ReleaseBusinessLogConfigResponse) {
	response = &ReleaseBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发布业务日志配置
func (c *Client) ReleaseBusinessLogConfig(request *ReleaseBusinessLogConfigRequest) (response *ReleaseBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewReleaseBusinessLogConfigRequest()
	}
	response = NewReleaseBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
	request = &DeleteRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteRoute")
	return
}

func NewDeleteRouteResponse() (response *DeleteRouteResponse) {
	response = &DeleteRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除路由
func (c *Client) DeleteRoute(request *DeleteRouteRequest) (response *DeleteRouteResponse, err error) {
	if request == nil {
		request = NewDeleteRouteRequest()
	}
	response = NewDeleteRouteResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveInstanceRequest() (request *RemoveInstanceRequest) {
	request = &RemoveInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RemoveInstance")
	return
}

func NewRemoveInstanceResponse() (response *RemoveInstanceResponse) {
	response = &RemoveInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从tsf集群中移除机器节点
func (c *Client) RemoveInstance(request *RemoveInstanceRequest) (response *RemoveInstanceResponse, err error) {
	if request == nil {
		request = NewRemoveInstanceRequest()
	}
	response = NewRemoveInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSidecarFilterRequest() (request *CreateSidecarFilterRequest) {
	request = &CreateSidecarFilterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateSidecarFilter")
	return
}

func NewCreateSidecarFilterResponse() (response *CreateSidecarFilterResponse) {
	response = &CreateSidecarFilterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Sidecar过滤器
func (c *Client) CreateSidecarFilter(request *CreateSidecarFilterRequest) (response *CreateSidecarFilterResponse, err error) {
	if request == nil {
		request = NewCreateSidecarFilterRequest()
	}
	response = NewCreateSidecarFilterResponse()
	err = c.Send(request, response)
	return
}

func NewRedoTaskFlowBatchRequest() (request *RedoTaskFlowBatchRequest) {
	request = &RedoTaskFlowBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskFlowBatch")
	return
}

func NewRedoTaskFlowBatchResponse() (response *RedoTaskFlowBatchResponse) {
	response = &RedoTaskFlowBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新执行工作流批次
func (c *Client) RedoTaskFlowBatch(request *RedoTaskFlowBatchRequest) (response *RedoTaskFlowBatchResponse, err error) {
	if request == nil {
		request = NewRedoTaskFlowBatchRequest()
	}
	response = NewRedoTaskFlowBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileConfigReleaseLogsRequest() (request *DescribeFileConfigReleaseLogsRequest) {
	request = &DescribeFileConfigReleaseLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigReleaseLogs")
	return
}

func NewDescribeFileConfigReleaseLogsResponse() (response *DescribeFileConfigReleaseLogsResponse) {
	response = &DescribeFileConfigReleaseLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件配置发布历史
func (c *Client) DescribeFileConfigReleaseLogs(request *DescribeFileConfigReleaseLogsRequest) (response *DescribeFileConfigReleaseLogsResponse, err error) {
	if request == nil {
		request = NewDescribeFileConfigReleaseLogsRequest()
	}
	response = NewDescribeFileConfigReleaseLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmTaskStatusRequest() (request *DescribeApmTaskStatusRequest) {
	request = &DescribeApmTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApmTaskStatus")
	return
}

func NewDescribeApmTaskStatusResponse() (response *DescribeApmTaskStatusResponse) {
	response = &DescribeApmTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询各种APM耗时任务当前执行状态
func (c *Client) DescribeApmTaskStatus(request *DescribeApmTaskStatusRequest) (response *DescribeApmTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApmTaskStatusRequest()
	}
	response = NewDescribeApmTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewStopTaskBatchRequest() (request *StopTaskBatchRequest) {
	request = &StopTaskBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StopTaskBatch")
	return
}

func NewStopTaskBatchResponse() (response *StopTaskBatchResponse) {
	response = &StopTaskBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止执行中的任务批次， 非运行中的任务不可调用。
func (c *Client) StopTaskBatch(request *StopTaskBatchRequest) (response *StopTaskBatchResponse, err error) {
	if request == nil {
		request = NewStopTaskBatchRequest()
	}
	response = NewStopTaskBatchResponse()
	err = c.Send(request, response)
	return
}

func NewRelateGroupToScalableRuleRequest() (request *RelateGroupToScalableRuleRequest) {
	request = &RelateGroupToScalableRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RelateGroupToScalableRule")
	return
}

func NewRelateGroupToScalableRuleResponse() (response *RelateGroupToScalableRuleResponse) {
	response = &RelateGroupToScalableRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联弹性伸缩规则到应用
func (c *Client) RelateGroupToScalableRule(request *RelateGroupToScalableRuleRequest) (response *RelateGroupToScalableRuleResponse, err error) {
	if request == nil {
		request = NewRelateGroupToScalableRuleRequest()
	}
	response = NewRelateGroupToScalableRuleResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateConfigWithGroupRequest() (request *AssociateConfigWithGroupRequest) {
	request = &AssociateConfigWithGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "AssociateConfigWithGroup")
	return
}

func NewAssociateConfigWithGroupResponse() (response *AssociateConfigWithGroupResponse) {
	response = &AssociateConfigWithGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联投递配置到部署组
func (c *Client) AssociateConfigWithGroup(request *AssociateConfigWithGroupRequest) (response *AssociateConfigWithGroupResponse, err error) {
	if request == nil {
		request = NewAssociateConfigWithGroupRequest()
	}
	response = NewAssociateConfigWithGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationsAttributeRequest() (request *DescribeApplicationsAttributeRequest) {
	request = &DescribeApplicationsAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationsAttribute")
	return
}

func NewDescribeApplicationsAttributeResponse() (response *DescribeApplicationsAttributeResponse) {
	response = &DescribeApplicationsAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) DescribeApplicationsAttribute(request *DescribeApplicationsAttributeRequest) (response *DescribeApplicationsAttributeResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationsAttributeRequest()
	}
	response = NewDescribeApplicationsAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
	request = &DescribeClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusters")
	return
}

func NewDescribeClustersResponse() (response *DescribeClustersResponse) {
	response = &DescribeClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群列表
func (c *Client) DescribeClusters(request *DescribeClustersRequest) (response *DescribeClustersResponse, err error) {
	if request == nil {
		request = NewDescribeClustersRequest()
	}
	response = NewDescribeClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlameGraphRequest() (request *DescribeFlameGraphRequest) {
	request = &DescribeFlameGraphRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlameGraph")
	return
}

func NewDescribeFlameGraphResponse() (response *DescribeFlameGraphResponse) {
	response = &DescribeFlameGraphResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询对应的java实例历史采集的火焰图
func (c *Client) DescribeFlameGraph(request *DescribeFlameGraphRequest) (response *DescribeFlameGraphResponse, err error) {
	if request == nil {
		request = NewDescribeFlameGraphRequest()
	}
	response = NewDescribeFlameGraphResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadMsApiRequest() (request *DownloadMsApiRequest) {
	request = &DownloadMsApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DownloadMsApi")
	return
}

func NewDownloadMsApiResponse() (response *DownloadMsApiResponse) {
	response = &DownloadMsApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出API列表
func (c *Client) DownloadMsApi(request *DownloadMsApiRequest) (response *DownloadMsApiResponse, err error) {
	if request == nil {
		request = NewDownloadMsApiRequest()
	}
	response = NewDownloadMsApiResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiStatusRequest() (request *DescribeApiStatusRequest) {
	request = &DescribeApiStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiStatus")
	return
}

func NewDescribeApiStatusResponse() (response *DescribeApiStatusResponse) {
	response = &DescribeApiStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API在线状态
func (c *Client) DescribeApiStatus(request *DescribeApiStatusRequest) (response *DescribeApiStatusResponse, err error) {
	if request == nil {
		request = NewDescribeApiStatusRequest()
	}
	response = NewDescribeApiStatusResponse()
	err = c.Send(request, response)
	return
}

func NewListGroupsByScalableRuleIdRequest() (request *ListGroupsByScalableRuleIdRequest) {
	request = &ListGroupsByScalableRuleIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListGroupsByScalableRuleId")
	return
}

func NewListGroupsByScalableRuleIdResponse() (response *ListGroupsByScalableRuleIdResponse) {
	response = &ListGroupsByScalableRuleIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取弹性扩缩容应用和部署组
func (c *Client) ListGroupsByScalableRuleId(request *ListGroupsByScalableRuleIdRequest) (response *ListGroupsByScalableRuleIdResponse, err error) {
	if request == nil {
		request = NewListGroupsByScalableRuleIdRequest()
	}
	response = NewListGroupsByScalableRuleIdResponse()
	err = c.Send(request, response)
	return
}

func NewModifyReleasePipelineTaskRequest() (request *ModifyReleasePipelineTaskRequest) {
	request = &ModifyReleasePipelineTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyReleasePipelineTask")
	return
}

func NewModifyReleasePipelineTaskResponse() (response *ModifyReleasePipelineTaskResponse) {
	response = &ModifyReleasePipelineTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改发布单任务
func (c *Client) ModifyReleasePipelineTask(request *ModifyReleasePipelineTaskRequest) (response *ModifyReleasePipelineTaskResponse, err error) {
	if request == nil {
		request = NewModifyReleasePipelineTaskRequest()
	}
	response = NewModifyReleasePipelineTaskResponse()
	err = c.Send(request, response)
	return
}

func NewStopTaskExecuteRequest() (request *StopTaskExecuteRequest) {
	request = &StopTaskExecuteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StopTaskExecute")
	return
}

func NewStopTaskExecuteResponse() (response *StopTaskExecuteResponse) {
	response = &StopTaskExecuteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止正在某个节点上执行的任务
func (c *Client) StopTaskExecute(request *StopTaskExecuteRequest) (response *StopTaskExecuteResponse, err error) {
	if request == nil {
		request = NewStopTaskExecuteRequest()
	}
	response = NewStopTaskExecuteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageUserIsExistsRequest() (request *DescribeImageUserIsExistsRequest) {
	request = &DescribeImageUserIsExistsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageUserIsExists")
	return
}

func NewDescribeImageUserIsExistsResponse() (response *DescribeImageUserIsExistsResponse) {
	response = &DescribeImageUserIsExistsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库用户是否已经开通
func (c *Client) DescribeImageUserIsExists(request *DescribeImageUserIsExistsRequest) (response *DescribeImageUserIsExistsResponse, err error) {
	if request == nil {
		request = NewDescribeImageUserIsExistsRequest()
	}
	response = NewDescribeImageUserIsExistsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCommonPkgRequest() (request *DescribeCommonPkgRequest) {
	request = &DescribeCommonPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeCommonPkg")
	return
}

func NewDescribeCommonPkgResponse() (response *DescribeCommonPkgResponse) {
	response = &DescribeCommonPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取公共包信息
func (c *Client) DescribeCommonPkg(request *DescribeCommonPkgRequest) (response *DescribeCommonPkgResponse, err error) {
	if request == nil {
		request = NewDescribeCommonPkgRequest()
	}
	response = NewDescribeCommonPkgResponse()
	err = c.Send(request, response)
	return
}

func NewDisableRatelimitRequest() (request *DisableRatelimitRequest) {
	request = &DisableRatelimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableRatelimit")
	return
}

func NewDisableRatelimitResponse() (response *DisableRatelimitResponse) {
	response = &DisableRatelimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭限流
func (c *Client) DisableRatelimit(request *DisableRatelimitRequest) (response *DisableRatelimitResponse, err error) {
	if request == nil {
		request = NewDisableRatelimitRequest()
	}
	response = NewDisableRatelimitResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateHealthCheckSettingsRequest() (request *UpdateHealthCheckSettingsRequest) {
	request = &UpdateHealthCheckSettingsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateHealthCheckSettings")
	return
}

func NewUpdateHealthCheckSettingsResponse() (response *UpdateHealthCheckSettingsResponse) {
	response = &UpdateHealthCheckSettingsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新健康检查配置
func (c *Client) UpdateHealthCheckSettings(request *UpdateHealthCheckSettingsRequest) (response *UpdateHealthCheckSettingsResponse, err error) {
	if request == nil {
		request = NewUpdateHealthCheckSettingsRequest()
	}
	response = NewUpdateHealthCheckSettingsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLicenseApplicationRequest() (request *DeleteLicenseApplicationRequest) {
	request = &DeleteLicenseApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteLicenseApplication")
	return
}

func NewDeleteLicenseApplicationResponse() (response *DeleteLicenseApplicationResponse) {
	response = &DeleteLicenseApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除申请
func (c *Client) DeleteLicenseApplication(request *DeleteLicenseApplicationRequest) (response *DeleteLicenseApplicationResponse, err error) {
	if request == nil {
		request = NewDeleteLicenseApplicationRequest()
	}
	response = NewDeleteLicenseApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthorizationRequest() (request *DescribeAuthorizationRequest) {
	request = &DescribeAuthorizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorization")
	return
}

func NewDescribeAuthorizationResponse() (response *DescribeAuthorizationResponse) {
	response = &DescribeAuthorizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务权限配置（废弃）
func (c *Client) DescribeAuthorization(request *DescribeAuthorizationRequest) (response *DescribeAuthorizationResponse, err error) {
	if request == nil {
		request = NewDescribeAuthorizationRequest()
	}
	response = NewDescribeAuthorizationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasePipelineRunTaskStatusRequest() (request *DescribeReleasePipelineRunTaskStatusRequest) {
	request = &DescribeReleasePipelineRunTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasePipelineRunTaskStatus")
	return
}

func NewDescribeReleasePipelineRunTaskStatusResponse() (response *DescribeReleasePipelineRunTaskStatusResponse) {
	response = &DescribeReleasePipelineRunTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布单执行任务的状态信息
func (c *Client) DescribeReleasePipelineRunTaskStatus(request *DescribeReleasePipelineRunTaskStatusRequest) (response *DescribeReleasePipelineRunTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeReleasePipelineRunTaskStatusRequest()
	}
	response = NewDescribeReleasePipelineRunTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewClearResourceDataRequest() (request *ClearResourceDataRequest) {
	request = &ClearResourceDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ClearResourceData")
	return
}

func NewClearResourceDataResponse() (response *ClearResourceDataResponse) {
	response = &ClearResourceDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清理资源
func (c *Client) ClearResourceData(request *ClearResourceDataRequest) (response *ClearResourceDataResponse, err error) {
	if request == nil {
		request = NewClearResourceDataRequest()
	}
	response = NewClearResourceDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventDimensionRequest() (request *DescribeEventDimensionRequest) {
	request = &DescribeEventDimensionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeEventDimension")
	return
}

func NewDescribeEventDimensionResponse() (response *DescribeEventDimensionResponse) {
	response = &DescribeEventDimensionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 事件统计维度表
func (c *Client) DescribeEventDimension(request *DescribeEventDimensionRequest) (response *DescribeEventDimensionResponse, err error) {
	if request == nil {
		request = NewDescribeEventDimensionRequest()
	}
	response = NewDescribeEventDimensionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePodInstancesRequest() (request *DescribePodInstancesRequest) {
	request = &DescribePodInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePodInstances")
	return
}

func NewDescribePodInstancesResponse() (response *DescribePodInstancesResponse) {
	response = &DescribePodInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取部署组实例列表
func (c *Client) DescribePodInstances(request *DescribePodInstancesRequest) (response *DescribePodInstancesResponse, err error) {
	if request == nil {
		request = NewDescribePodInstancesRequest()
	}
	response = NewDescribePodInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyModuleConfRequest() (request *ModifyModuleConfRequest) {
	request = &ModifyModuleConfRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyModuleConf")
	return
}

func NewModifyModuleConfResponse() (response *ModifyModuleConfResponse) {
	response = &ModifyModuleConfResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改模块配置
func (c *Client) ModifyModuleConf(request *ModifyModuleConfRequest) (response *ModifyModuleConfResponse, err error) {
	if request == nil {
		request = NewModifyModuleConfRequest()
	}
	response = NewModifyModuleConfResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateRequest() (request *CreateTemplateRequest) {
	request = &CreateTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateTemplate")
	return
}

func NewCreateTemplateResponse() (response *CreateTemplateResponse) {
	response = &CreateTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存模板草稿
func (c *Client) CreateTemplate(request *CreateTemplateRequest) (response *CreateTemplateResponse, err error) {
	if request == nil {
		request = NewCreateTemplateRequest()
	}
	response = NewCreateTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOverviweServiceRequest() (request *DescribeOverviweServiceRequest) {
	request = &DescribeOverviweServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeOverviweService")
	return
}

func NewDescribeOverviweServiceResponse() (response *DescribeOverviweServiceResponse) {
	response = &DescribeOverviweServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 概览页微服务信息接口
func (c *Client) DescribeOverviweService(request *DescribeOverviweServiceRequest) (response *DescribeOverviweServiceResponse, err error) {
	if request == nil {
		request = NewDescribeOverviweServiceRequest()
	}
	response = NewDescribeOverviweServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupAddibleInstancesRequest() (request *DescribeGroupAddibleInstancesRequest) {
	request = &DescribeGroupAddibleInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupAddibleInstances")
	return
}

func NewDescribeGroupAddibleInstancesResponse() (response *DescribeGroupAddibleInstancesResponse) {
	response = &DescribeGroupAddibleInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//  查询分组可添加的机器列表
func (c *Client) DescribeGroupAddibleInstances(request *DescribeGroupAddibleInstancesRequest) (response *DescribeGroupAddibleInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeGroupAddibleInstancesRequest()
	}
	response = NewDescribeGroupAddibleInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCreateMoreTaskRequest() (request *EnableCreateMoreTaskRequest) {
	request = &EnableCreateMoreTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableCreateMoreTask")
	return
}

func NewEnableCreateMoreTaskResponse() (response *EnableCreateMoreTaskResponse) {
	response = &EnableCreateMoreTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建的任务有上限， 此接口用于判断用户是否可以创建任务。
func (c *Client) EnableCreateMoreTask(request *EnableCreateMoreTaskRequest) (response *EnableCreateMoreTaskResponse, err error) {
	if request == nil {
		request = NewEnableCreateMoreTaskRequest()
	}
	response = NewEnableCreateMoreTaskResponse()
	err = c.Send(request, response)
	return
}

func NewSetJvmMonitorMethodProfileRequest() (request *SetJvmMonitorMethodProfileRequest) {
	request = &SetJvmMonitorMethodProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SetJvmMonitorMethodProfile")
	return
}

func NewSetJvmMonitorMethodProfileResponse() (response *SetJvmMonitorMethodProfileResponse) {
	response = &SetJvmMonitorMethodProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置JVM监控的方法执行分析
func (c *Client) SetJvmMonitorMethodProfile(request *SetJvmMonitorMethodProfileRequest) (response *SetJvmMonitorMethodProfileResponse, err error) {
	if request == nil {
		request = NewSetJvmMonitorMethodProfileRequest()
	}
	response = NewSetJvmMonitorMethodProfileResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopFailureRateInterfacesRequest() (request *GetTopFailureRateInterfacesRequest) {
	request = &GetTopFailureRateInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTopFailureRateInterfaces")
	return
}

func NewGetTopFailureRateInterfacesResponse() (response *GetTopFailureRateInterfacesResponse) {
	response = &GetTopFailureRateInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TopN请求失败率接口列表
func (c *Client) GetTopFailureRateInterfaces(request *GetTopFailureRateInterfacesRequest) (response *GetTopFailureRateInterfacesResponse, err error) {
	if request == nil {
		request = NewGetTopFailureRateInterfacesRequest()
	}
	response = NewGetTopFailureRateInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewOperateKubeInjectRequest() (request *OperateKubeInjectRequest) {
	request = &OperateKubeInjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "OperateKubeInject")
	return
}

func NewOperateKubeInjectResponse() (response *OperateKubeInjectResponse) {
	response = &OperateKubeInjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 容器轻量化操作
func (c *Client) OperateKubeInject(request *OperateKubeInjectRequest) (response *OperateKubeInjectResponse, err error) {
	if request == nil {
		request = NewOperateKubeInjectRequest()
	}
	response = NewOperateKubeInjectResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDockerForUseRequest() (request *DescribeDockerForUseRequest) {
	request = &DescribeDockerForUseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeDockerForUse")
	return
}

func NewDescribeDockerForUseResponse() (response *DescribeDockerForUseResponse) {
	response = &DescribeDockerForUseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取docker使用指引
func (c *Client) DescribeDockerForUse(request *DescribeDockerForUseRequest) (response *DescribeDockerForUseResponse, err error) {
	if request == nil {
		request = NewDescribeDockerForUseRequest()
	}
	response = NewDescribeDockerForUseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSingleContainerGroupsRequest() (request *DescribeSingleContainerGroupsRequest) {
	request = &DescribeSingleContainerGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSingleContainerGroups")
	return
}

func NewDescribeSingleContainerGroupsResponse() (response *DescribeSingleContainerGroupsResponse) {
	response = &DescribeSingleContainerGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取分组列表-独立菜单
func (c *Client) DescribeSingleContainerGroups(request *DescribeSingleContainerGroupsRequest) (response *DescribeSingleContainerGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeSingleContainerGroupsRequest()
	}
	response = NewDescribeSingleContainerGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewListAlarmPoliciesRequest() (request *ListAlarmPoliciesRequest) {
	request = &ListAlarmPoliciesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListAlarmPolicies")
	return
}

func NewListAlarmPoliciesResponse() (response *ListAlarmPoliciesResponse) {
	response = &ListAlarmPoliciesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警策略列表
func (c *Client) ListAlarmPolicies(request *ListAlarmPoliciesRequest) (response *ListAlarmPoliciesResponse, err error) {
	if request == nil {
		request = NewListAlarmPoliciesRequest()
	}
	response = NewListAlarmPoliciesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableFallbackRouteRequest() (request *EnableFallbackRouteRequest) {
	request = &EnableFallbackRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableFallbackRoute")
	return
}

func NewEnableFallbackRouteResponse() (response *EnableFallbackRouteResponse) {
	response = &EnableFallbackRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用微服务路由保护策略
func (c *Client) EnableFallbackRoute(request *EnableFallbackRouteRequest) (response *EnableFallbackRouteResponse, err error) {
	if request == nil {
		request = NewEnableFallbackRouteRequest()
	}
	response = NewEnableFallbackRouteResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApplicationRequest() (request *ModifyApplicationRequest) {
	request = &ModifyApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyApplication")
	return
}

func NewModifyApplicationResponse() (response *ModifyApplicationResponse) {
	response = &ModifyApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改应用
func (c *Client) ModifyApplication(request *ModifyApplicationRequest) (response *ModifyApplicationResponse, err error) {
	if request == nil {
		request = NewModifyApplicationRequest()
	}
	response = NewModifyApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewStopServerlessGroupRequest() (request *StopServerlessGroupRequest) {
	request = &StopServerlessGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StopServerlessGroup")
	return
}

func NewStopServerlessGroupResponse() (response *StopServerlessGroupResponse) {
	response = &StopServerlessGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) StopServerlessGroup(request *StopServerlessGroupRequest) (response *StopServerlessGroupResponse, err error) {
	if request == nil {
		request = NewStopServerlessGroupRequest()
	}
	response = NewStopServerlessGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateScalableRuleRequest() (request *CreateScalableRuleRequest) {
	request = &CreateScalableRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateScalableRule")
	return
}

func NewCreateScalableRuleResponse() (response *CreateScalableRuleResponse) {
	response = &CreateScalableRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建弹性扩缩容规则
func (c *Client) CreateScalableRule(request *CreateScalableRuleRequest) (response *CreateScalableRuleResponse, err error) {
	if request == nil {
		request = NewCreateScalableRuleRequest()
	}
	response = NewCreateScalableRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRepositoryRequest() (request *CreateRepositoryRequest) {
	request = &CreateRepositoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateRepository")
	return
}

func NewCreateRepositoryResponse() (response *CreateRepositoryResponse) {
	response = &CreateRepositoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建仓库
func (c *Client) CreateRepository(request *CreateRepositoryRequest) (response *CreateRepositoryResponse, err error) {
	if request == nil {
		request = NewCreateRepositoryRequest()
	}
	response = NewCreateRepositoryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLaneRequest() (request *ModifyLaneRequest) {
	request = &ModifyLaneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyLane")
	return
}

func NewModifyLaneResponse() (response *ModifyLaneResponse) {
	response = &ModifyLaneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新泳道信息
func (c *Client) ModifyLane(request *ModifyLaneRequest) (response *ModifyLaneResponse, err error) {
	if request == nil {
		request = NewModifyLaneRequest()
	}
	response = NewModifyLaneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupInstancesUnifiedRequest() (request *DescribeGroupInstancesUnifiedRequest) {
	request = &DescribeGroupInstancesUnifiedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupInstancesUnified")
	return
}

func NewDescribeGroupInstancesUnifiedResponse() (response *DescribeGroupInstancesUnifiedResponse) {
	response = &DescribeGroupInstancesUnifiedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 当容器时，返回 pod 列表
func (c *Client) DescribeGroupInstancesUnified(request *DescribeGroupInstancesUnifiedRequest) (response *DescribeGroupInstancesUnifiedResponse, err error) {
	if request == nil {
		request = NewDescribeGroupInstancesUnifiedRequest()
	}
	response = NewDescribeGroupInstancesUnifiedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigReleasesRequest() (request *DescribeConfigReleasesRequest) {
	request = &DescribeConfigReleasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigReleases")
	return
}

func NewDescribeConfigReleasesResponse() (response *DescribeConfigReleasesResponse) {
	response = &DescribeConfigReleasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询配置发布信息
func (c *Client) DescribeConfigReleases(request *DescribeConfigReleasesRequest) (response *DescribeConfigReleasesResponse, err error) {
	if request == nil {
		request = NewDescribeConfigReleasesRequest()
	}
	response = NewDescribeConfigReleasesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchOssSpanBusinessLogRequest() (request *SearchOssSpanBusinessLogRequest) {
	request = &SearchOssSpanBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchOssSpanBusinessLog")
	return
}

func NewSearchOssSpanBusinessLogResponse() (response *SearchOssSpanBusinessLogResponse) {
	response = &SearchOssSpanBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务日志搜索（关联调用链Span）
func (c *Client) SearchOssSpanBusinessLog(request *SearchOssSpanBusinessLogRequest) (response *SearchOssSpanBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchOssSpanBusinessLogRequest()
	}
	response = NewSearchOssSpanBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewGetOssTraceServicesRequest() (request *GetOssTraceServicesRequest) {
	request = &GetOssTraceServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetOssTraceServices")
	return
}

func NewGetOssTraceServicesResponse() (response *GetOssTraceServicesResponse) {
	response = &GetOssTraceServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用链服务列表
func (c *Client) GetOssTraceServices(request *GetOssTraceServicesRequest) (response *GetOssTraceServicesResponse, err error) {
	if request == nil {
		request = NewGetOssTraceServicesRequest()
	}
	response = NewGetOssTraceServicesResponse()
	err = c.Send(request, response)
	return
}

func NewListAppPkgRequest() (request *ListAppPkgRequest) {
	request = &ListAppPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListAppPkg")
	return
}

func NewListAppPkgResponse() (response *ListAppPkgResponse) {
	response = &ListAppPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的应用信息
func (c *Client) ListAppPkg(request *ListAppPkgRequest) (response *ListAppPkgResponse, err error) {
	if request == nil {
		request = NewListAppPkgRequest()
	}
	response = NewListAppPkgResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopReqAmountServicesRequest() (request *GetTopReqAmountServicesRequest) {
	request = &GetTopReqAmountServicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTopReqAmountServices")
	return
}

func NewGetTopReqAmountServicesResponse() (response *GetTopReqAmountServicesResponse) {
	response = &GetTopReqAmountServicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询TopN请求量服务列表
func (c *Client) GetTopReqAmountServices(request *GetTopReqAmountServicesRequest) (response *GetTopReqAmountServicesResponse, err error) {
	if request == nil {
		request = NewGetTopReqAmountServicesRequest()
	}
	response = NewGetTopReqAmountServicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApiRateLimitRulesRequest() (request *DescribeApiRateLimitRulesRequest) {
	request = &DescribeApiRateLimitRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApiRateLimitRules")
	return
}

func NewDescribeApiRateLimitRulesResponse() (response *DescribeApiRateLimitRulesResponse) {
	response = &DescribeApiRateLimitRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询API限流规则
func (c *Client) DescribeApiRateLimitRules(request *DescribeApiRateLimitRulesRequest) (response *DescribeApiRateLimitRulesResponse, err error) {
	if request == nil {
		request = NewDescribeApiRateLimitRulesRequest()
	}
	response = NewDescribeApiRateLimitRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSidecarMetricsRequest() (request *DescribeSidecarMetricsRequest) {
	request = &DescribeSidecarMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSidecarMetrics")
	return
}

func NewDescribeSidecarMetricsResponse() (response *DescribeSidecarMetricsResponse) {
	response = &DescribeSidecarMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取sidecar监控指标
func (c *Client) DescribeSidecarMetrics(request *DescribeSidecarMetricsRequest) (response *DescribeSidecarMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeSidecarMetricsRequest()
	}
	response = NewDescribeSidecarMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUploadInfoRequest() (request *ModifyUploadInfoRequest) {
	request = &ModifyUploadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyUploadInfo")
	return
}

func NewModifyUploadInfoResponse() (response *ModifyUploadInfoResponse) {
	response = &ModifyUploadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用该接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
// 调用此接口完成后，才标志上传包流程结束。
func (c *Client) ModifyUploadInfo(request *ModifyUploadInfoRequest) (response *ModifyUploadInfoResponse, err error) {
	if request == nil {
		request = NewModifyUploadInfoRequest()
	}
	response = NewModifyUploadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewListMachinesRequest() (request *ListMachinesRequest) {
	request = &ListMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListMachines")
	return
}

func NewListMachinesResponse() (response *ListMachinesResponse) {
	response = &ListMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询机器列表
func (c *Client) ListMachines(request *ListMachinesRequest) (response *ListMachinesResponse, err error) {
	if request == nil {
		request = NewListMachinesRequest()
	}
	response = NewListMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnitNamespaceListRequest() (request *DescribeUnitNamespaceListRequest) {
	request = &DescribeUnitNamespaceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitNamespaceList")
	return
}

func NewDescribeUnitNamespaceListResponse() (response *DescribeUnitNamespaceListResponse) {
	response = &DescribeUnitNamespaceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 条件查询单元化命名空间列表
func (c *Client) DescribeUnitNamespaceList(request *DescribeUnitNamespaceListRequest) (response *DescribeUnitNamespaceListResponse, err error) {
	if request == nil {
		request = NewDescribeUnitNamespaceListRequest()
	}
	response = NewDescribeUnitNamespaceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupOtherRequest() (request *DescribeGroupOtherRequest) {
	request = &DescribeGroupOtherRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupOther")
	return
}

func NewDescribeGroupOtherResponse() (response *DescribeGroupOtherResponse) {
	response = &DescribeGroupOtherResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取部署组其他字段
func (c *Client) DescribeGroupOther(request *DescribeGroupOtherRequest) (response *DescribeGroupOtherResponse, err error) {
	if request == nil {
		request = NewDescribeGroupOtherRequest()
	}
	response = NewDescribeGroupOtherResponse()
	err = c.Send(request, response)
	return
}

func NewRedoTaskExecuteRequest() (request *RedoTaskExecuteRequest) {
	request = &RedoTaskExecuteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskExecute")
	return
}

func NewRedoTaskExecuteResponse() (response *RedoTaskExecuteResponse) {
	response = &RedoTaskExecuteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新执行在某个节点上执行任务。
func (c *Client) RedoTaskExecute(request *RedoTaskExecuteRequest) (response *RedoTaskExecuteResponse, err error) {
	if request == nil {
		request = NewRedoTaskExecuteRequest()
	}
	response = NewRedoTaskExecuteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApplicationMicroserviceTypesRequest() (request *DescribeApplicationMicroserviceTypesRequest) {
	request = &DescribeApplicationMicroserviceTypesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeApplicationMicroserviceTypes")
	return
}

func NewDescribeApplicationMicroserviceTypesResponse() (response *DescribeApplicationMicroserviceTypesResponse) {
	response = &DescribeApplicationMicroserviceTypesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过微服务id查询应用微服务类型
func (c *Client) DescribeApplicationMicroserviceTypes(request *DescribeApplicationMicroserviceTypesRequest) (response *DescribeApplicationMicroserviceTypesResponse, err error) {
	if request == nil {
		request = NewDescribeApplicationMicroserviceTypesRequest()
	}
	response = NewDescribeApplicationMicroserviceTypesResponse()
	err = c.Send(request, response)
	return
}

func NewStopGroupRequest() (request *StopGroupRequest) {
	request = &StopGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StopGroup")
	return
}

func NewStopGroupResponse() (response *StopGroupResponse) {
	response = &StopGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止虚拟机部署组
func (c *Client) StopGroup(request *StopGroupRequest) (response *StopGroupResponse, err error) {
	if request == nil {
		request = NewStopGroupRequest()
	}
	response = NewStopGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthorizationInfoRequest() (request *DescribeAuthorizationInfoRequest) {
	request = &DescribeAuthorizationInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorizationInfo")
	return
}

func NewDescribeAuthorizationInfoResponse() (response *DescribeAuthorizationInfoResponse) {
	response = &DescribeAuthorizationInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务权限配置
func (c *Client) DescribeAuthorizationInfo(request *DescribeAuthorizationInfoRequest) (response *DescribeAuthorizationInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAuthorizationInfoRequest()
	}
	response = NewDescribeAuthorizationInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSidecarFiltersRequest() (request *DescribeSidecarFiltersRequest) {
	request = &DescribeSidecarFiltersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSidecarFilters")
	return
}

func NewDescribeSidecarFiltersResponse() (response *DescribeSidecarFiltersResponse) {
	response = &DescribeSidecarFiltersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Sidecar过滤器列表
func (c *Client) DescribeSidecarFilters(request *DescribeSidecarFiltersRequest) (response *DescribeSidecarFiltersResponse, err error) {
	if request == nil {
		request = NewDescribeSidecarFiltersRequest()
	}
	response = NewDescribeSidecarFiltersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRouteRuleRequest() (request *CreateRouteRuleRequest) {
	request = &CreateRouteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateRouteRule")
	return
}

func NewCreateRouteRuleResponse() (response *CreateRouteRuleResponse) {
	response = &CreateRouteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建路由规则
func (c *Client) CreateRouteRule(request *CreateRouteRuleRequest) (response *CreateRouteRuleResponse, err error) {
	if request == nil {
		request = NewCreateRouteRuleRequest()
	}
	response = NewCreateRouteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewInstallLicenseRequest() (request *InstallLicenseRequest) {
	request = &InstallLicenseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "InstallLicense")
	return
}

func NewInstallLicenseResponse() (response *InstallLicenseResponse) {
	response = &InstallLicenseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 装入证书
func (c *Client) InstallLicense(request *InstallLicenseRequest) (response *InstallLicenseResponse, err error) {
	if request == nil {
		request = NewInstallLicenseRequest()
	}
	response = NewInstallLicenseResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateDeliveryConfigRequest() (request *UpdateDeliveryConfigRequest) {
	request = &UpdateDeliveryConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateDeliveryConfig")
	return
}

func NewUpdateDeliveryConfigResponse() (response *UpdateDeliveryConfigResponse) {
	response = &UpdateDeliveryConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新传递配置
func (c *Client) UpdateDeliveryConfig(request *UpdateDeliveryConfigRequest) (response *UpdateDeliveryConfigResponse, err error) {
	if request == nil {
		request = NewUpdateDeliveryConfigRequest()
	}
	response = NewUpdateDeliveryConfigResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUnitRuleRequest() (request *UpdateUnitRuleRequest) {
	request = &UpdateUnitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateUnitRule")
	return
}

func NewUpdateUnitRuleResponse() (response *UpdateUnitRuleResponse) {
	response = &UpdateUnitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新单元化规则
func (c *Client) UpdateUnitRule(request *UpdateUnitRuleRequest) (response *UpdateUnitRuleResponse, err error) {
	if request == nil {
		request = NewUpdateUnitRuleRequest()
	}
	response = NewUpdateUnitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteScalableRuleRequest() (request *DeleteScalableRuleRequest) {
	request = &DeleteScalableRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteScalableRule")
	return
}

func NewDeleteScalableRuleResponse() (response *DeleteScalableRuleResponse) {
	response = &DeleteScalableRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除弹性伸缩规则
func (c *Client) DeleteScalableRule(request *DeleteScalableRuleRequest) (response *DeleteScalableRuleResponse, err error) {
	if request == nil {
		request = NewDeleteScalableRuleRequest()
	}
	response = NewDeleteScalableRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFilebeatConfigRequest() (request *DescribeFilebeatConfigRequest) {
	request = &DescribeFilebeatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFilebeatConfig")
	return
}

func NewDescribeFilebeatConfigResponse() (response *DescribeFilebeatConfigResponse) {
	response = &DescribeFilebeatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端配置通用配置（继承运营端配置，租户端优先级大于运营端），一个配置可以对应多个部署组
func (c *Client) DescribeFilebeatConfig(request *DescribeFilebeatConfigRequest) (response *DescribeFilebeatConfigResponse, err error) {
	if request == nil {
		request = NewDescribeFilebeatConfigRequest()
	}
	response = NewDescribeFilebeatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProgramRequest() (request *DescribeProgramRequest) {
	request = &DescribeProgramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeProgram")
	return
}

func NewDescribeProgramResponse() (response *DescribeProgramResponse) {
	response = &DescribeProgramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询数据集
func (c *Client) DescribeProgram(request *DescribeProgramRequest) (response *DescribeProgramResponse, err error) {
	if request == nil {
		request = NewDescribeProgramRequest()
	}
	response = NewDescribeProgramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTsfZonesRequest() (request *DescribeTsfZonesRequest) {
	request = &DescribeTsfZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTsfZones")
	return
}

func NewDescribeTsfZonesResponse() (response *DescribeTsfZonesResponse) {
	response = &DescribeTsfZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询基础可用区
func (c *Client) DescribeTsfZones(request *DescribeTsfZonesRequest) (response *DescribeTsfZonesResponse, err error) {
	if request == nil {
		request = NewDescribeTsfZonesRequest()
	}
	response = NewDescribeTsfZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLaneRequest() (request *DescribeLaneRequest) {
	request = &DescribeLaneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLane")
	return
}

func NewDescribeLaneResponse() (response *DescribeLaneResponse) {
	response = &DescribeLaneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询泳道详情
func (c *Client) DescribeLane(request *DescribeLaneRequest) (response *DescribeLaneResponse, err error) {
	if request == nil {
		request = NewDescribeLaneRequest()
	}
	response = NewDescribeLaneResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayJwtPluginRequest() (request *CreateGatewayJwtPluginRequest) {
	request = &CreateGatewayJwtPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayJwtPlugin")
	return
}

func NewCreateGatewayJwtPluginResponse() (response *CreateGatewayJwtPluginResponse) {
	response = &CreateGatewayJwtPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增网关JWT认证插件
func (c *Client) CreateGatewayJwtPlugin(request *CreateGatewayJwtPluginRequest) (response *CreateGatewayJwtPluginResponse, err error) {
	if request == nil {
		request = NewCreateGatewayJwtPluginRequest()
	}
	response = NewCreateGatewayJwtPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteReleaseHistoryRequest() (request *DescribeRouteReleaseHistoryRequest) {
	request = &DescribeRouteReleaseHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRouteReleaseHistory")
	return
}

func NewDescribeRouteReleaseHistoryResponse() (response *DescribeRouteReleaseHistoryResponse) {
	response = &DescribeRouteReleaseHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由规则启用 停用记录
func (c *Client) DescribeRouteReleaseHistory(request *DescribeRouteReleaseHistoryRequest) (response *DescribeRouteReleaseHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeRouteReleaseHistoryRequest()
	}
	response = NewDescribeRouteReleaseHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAuthorizationTypeRequest() (request *ModifyAuthorizationTypeRequest) {
	request = &ModifyAuthorizationTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyAuthorizationType")
	return
}

func NewModifyAuthorizationTypeResponse() (response *ModifyAuthorizationTypeResponse) {
	response = &ModifyAuthorizationTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改鉴权类型
func (c *Client) ModifyAuthorizationType(request *ModifyAuthorizationTypeRequest) (response *ModifyAuthorizationTypeResponse, err error) {
	if request == nil {
		request = NewModifyAuthorizationTypeRequest()
	}
	response = NewModifyAuthorizationTypeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteApplicationRequest() (request *DeleteApplicationRequest) {
	request = &DeleteApplicationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteApplication")
	return
}

func NewDeleteApplicationResponse() (response *DeleteApplicationResponse) {
	response = &DeleteApplicationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用
func (c *Client) DeleteApplication(request *DeleteApplicationRequest) (response *DeleteApplicationResponse, err error) {
	if request == nil {
		request = NewDeleteApplicationRequest()
	}
	response = NewDeleteApplicationResponse()
	err = c.Send(request, response)
	return
}

func NewBindPluginRequest() (request *BindPluginRequest) {
	request = &BindPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "BindPlugin")
	return
}

func NewBindPluginResponse() (response *BindPluginResponse) {
	response = &BindPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 插件与网关分组/API批量绑定
func (c *Client) BindPlugin(request *BindPluginRequest) (response *BindPluginResponse, err error) {
	if request == nil {
		request = NewBindPluginRequest()
	}
	response = NewBindPluginResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGroupRequest() (request *CreateGroupRequest) {
	request = &CreateGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateGroup")
	return
}

func NewCreateGroupResponse() (response *CreateGroupResponse) {
	response = &CreateGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建虚拟机部署组
func (c *Client) CreateGroup(request *CreateGroupRequest) (response *CreateGroupResponse, err error) {
	if request == nil {
		request = NewCreateGroupRequest()
	}
	response = NewCreateGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMonitorStatisticsPolicyRequest() (request *DescribeMonitorStatisticsPolicyRequest) {
	request = &DescribeMonitorStatisticsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMonitorStatisticsPolicy")
	return
}

func NewDescribeMonitorStatisticsPolicyResponse() (response *DescribeMonitorStatisticsPolicyResponse) {
	response = &DescribeMonitorStatisticsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定id的监控统计策略
func (c *Client) DescribeMonitorStatisticsPolicy(request *DescribeMonitorStatisticsPolicyRequest) (response *DescribeMonitorStatisticsPolicyResponse, err error) {
	if request == nil {
		request = NewDescribeMonitorStatisticsPolicyRequest()
	}
	response = NewDescribeMonitorStatisticsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpanDetailRequest() (request *DescribeSpanDetailRequest) {
	request = &DescribeSpanDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSpanDetail")
	return
}

func NewDescribeSpanDetailResponse() (response *DescribeSpanDetailResponse) {
	response = &DescribeSpanDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用链单跳数据详情 包含调用方和被调方数据
func (c *Client) DescribeSpanDetail(request *DescribeSpanDetailRequest) (response *DescribeSpanDetailResponse, err error) {
	if request == nil {
		request = NewDescribeSpanDetailRequest()
	}
	response = NewDescribeSpanDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigTemplateRequest() (request *DeleteConfigTemplateRequest) {
	request = &DeleteConfigTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteConfigTemplate")
	return
}

func NewDeleteConfigTemplateResponse() (response *DeleteConfigTemplateResponse) {
	response = &DeleteConfigTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除模板
func (c *Client) DeleteConfigTemplate(request *DeleteConfigTemplateRequest) (response *DeleteConfigTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteConfigTemplateRequest()
	}
	response = NewDeleteConfigTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileConfigRequest() (request *DescribeFileConfigRequest) {
	request = &DescribeFileConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfig")
	return
}

func NewDescribeFileConfigResponse() (response *DescribeFileConfigResponse) {
	response = &DescribeFileConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件配置项
func (c *Client) DescribeFileConfig(request *DescribeFileConfigRequest) (response *DescribeFileConfigResponse, err error) {
	if request == nil {
		request = NewDescribeFileConfigRequest()
	}
	response = NewDescribeFileConfigResponse()
	err = c.Send(request, response)
	return
}

func NewSetEnvoyDebugRequest() (request *SetEnvoyDebugRequest) {
	request = &SetEnvoyDebugRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SetEnvoyDebug")
	return
}

func NewSetEnvoyDebugResponse() (response *SetEnvoyDebugResponse) {
	response = &SetEnvoyDebugResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取sidecar监控指标
func (c *Client) SetEnvoyDebug(request *SetEnvoyDebugRequest) (response *SetEnvoyDebugResponse, err error) {
	if request == nil {
		request = NewSetEnvoyDebugRequest()
	}
	response = NewSetEnvoyDebugResponse()
	err = c.Send(request, response)
	return
}

func NewListMonitorStatisticsPolicyRequest() (request *ListMonitorStatisticsPolicyRequest) {
	request = &ListMonitorStatisticsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListMonitorStatisticsPolicy")
	return
}

func NewListMonitorStatisticsPolicyResponse() (response *ListMonitorStatisticsPolicyResponse) {
	response = &ListMonitorStatisticsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询监控统计策略列表.
func (c *Client) ListMonitorStatisticsPolicy(request *ListMonitorStatisticsPolicyRequest) (response *ListMonitorStatisticsPolicyResponse, err error) {
	if request == nil {
		request = NewListMonitorStatisticsPolicyRequest()
	}
	response = NewListMonitorStatisticsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeliveryConfigByGroupIdRequest() (request *DescribeDeliveryConfigByGroupIdRequest) {
	request = &DescribeDeliveryConfigByGroupIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeliveryConfigByGroupId")
	return
}

func NewDescribeDeliveryConfigByGroupIdResponse() (response *DescribeDeliveryConfigByGroupIdResponse) {
	response = &DescribeDeliveryConfigByGroupIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用部署组id获取绑定信息
func (c *Client) DescribeDeliveryConfigByGroupId(request *DescribeDeliveryConfigByGroupIdRequest) (response *DescribeDeliveryConfigByGroupIdResponse, err error) {
	if request == nil {
		request = NewDescribeDeliveryConfigByGroupIdRequest()
	}
	response = NewDescribeDeliveryConfigByGroupIdResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateFilebeatConfigRequest() (request *UpdateFilebeatConfigRequest) {
	request = &UpdateFilebeatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateFilebeatConfig")
	return
}

func NewUpdateFilebeatConfigResponse() (response *UpdateFilebeatConfigResponse) {
	response = &UpdateFilebeatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端配置通用配置（继承运营端配置，租户端优先级大于运营端），一个配置可以对应多个部署组
// 日志配置项，一个配置可以对应多个部署组
// 日志投递项，一个配置可以对应多个部署组
func (c *Client) UpdateFilebeatConfig(request *UpdateFilebeatConfigRequest) (response *UpdateFilebeatConfigResponse, err error) {
	if request == nil {
		request = NewUpdateFilebeatConfigRequest()
	}
	response = NewUpdateFilebeatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTaskRequest() (request *ModifyTaskRequest) {
	request = &ModifyTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyTask")
	return
}

func NewModifyTaskResponse() (response *ModifyTaskResponse) {
	response = &ModifyTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改任务
func (c *Client) ModifyTask(request *ModifyTaskRequest) (response *ModifyTaskResponse, err error) {
	if request == nil {
		request = NewModifyTaskRequest()
	}
	response = NewModifyTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMachinesRequest() (request *CreateMachinesRequest) {
	request = &CreateMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateMachines")
	return
}

func NewCreateMachinesResponse() (response *CreateMachinesResponse) {
	response = &CreateMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加机器列表
func (c *Client) CreateMachines(request *CreateMachinesRequest) (response *CreateMachinesResponse, err error) {
	if request == nil {
		request = NewCreateMachinesRequest()
	}
	response = NewCreateMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewDeployContainGroupRequest() (request *DeployContainGroupRequest) {
	request = &DeployContainGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeployContainGroup")
	return
}

func NewDeployContainGroupResponse() (response *DeployContainGroupResponse) {
	response = &DeployContainGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署容器应用
func (c *Client) DeployContainGroup(request *DeployContainGroupRequest) (response *DeployContainGroupResponse, err error) {
	if request == nil {
		request = NewDeployContainGroupRequest()
	}
	response = NewDeployContainGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGroupSecretRequest() (request *UpdateGroupSecretRequest) {
	request = &UpdateGroupSecretRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGroupSecret")
	return
}

func NewUpdateGroupSecretResponse() (response *UpdateGroupSecretResponse) {
	response = &UpdateGroupSecretResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新API分组的秘钥
func (c *Client) UpdateGroupSecret(request *UpdateGroupSecretRequest) (response *UpdateGroupSecretResponse, err error) {
	if request == nil {
		request = NewUpdateGroupSecretRequest()
	}
	response = NewUpdateGroupSecretResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateGatewayTagPluginRequest() (request *UpdateGatewayTagPluginRequest) {
	request = &UpdateGatewayTagPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateGatewayTagPlugin")
	return
}

func NewUpdateGatewayTagPluginResponse() (response *UpdateGatewayTagPluginResponse) {
	response = &UpdateGatewayTagPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改网关Tag插件
func (c *Client) UpdateGatewayTagPlugin(request *UpdateGatewayTagPluginRequest) (response *UpdateGatewayTagPluginResponse, err error) {
	if request == nil {
		request = NewUpdateGatewayTagPluginRequest()
	}
	response = NewUpdateGatewayTagPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteViewRequest() (request *DeleteViewRequest) {
	request = &DeleteViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteView")
	return
}

func NewDeleteViewResponse() (response *DeleteViewResponse) {
	response = &DeleteViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除视图
func (c *Client) DeleteView(request *DeleteViewRequest) (response *DeleteViewResponse, err error) {
	if request == nil {
		request = NewDeleteViewRequest()
	}
	response = NewDeleteViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScalableRuleAttributeByGroupRequest() (request *DescribeScalableRuleAttributeByGroupRequest) {
	request = &DescribeScalableRuleAttributeByGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeScalableRuleAttributeByGroup")
	return
}

func NewDescribeScalableRuleAttributeByGroupResponse() (response *DescribeScalableRuleAttributeByGroupResponse) {
	response = &DescribeScalableRuleAttributeByGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 显示关联的弹性伸缩规则
func (c *Client) DescribeScalableRuleAttributeByGroup(request *DescribeScalableRuleAttributeByGroupRequest) (response *DescribeScalableRuleAttributeByGroupResponse, err error) {
	if request == nil {
		request = NewDescribeScalableRuleAttributeByGroupRequest()
	}
	response = NewDescribeScalableRuleAttributeByGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadMultipartPkgRequest() (request *DownloadMultipartPkgRequest) {
	request = &DownloadMultipartPkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DownloadMultipartPkg")
	return
}

func NewDownloadMultipartPkgResponse() (response *DownloadMultipartPkgResponse) {
	response = &DownloadMultipartPkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分片下载程序包
func (c *Client) DownloadMultipartPkg(request *DownloadMultipartPkgRequest) (response *DownloadMultipartPkgResponse, err error) {
	if request == nil {
		request = NewDownloadMultipartPkgRequest()
	}
	response = NewDownloadMultipartPkgResponse()
	err = c.Send(request, response)
	return
}

func NewRevokeFileConfigRequest() (request *RevokeFileConfigRequest) {
	request = &RevokeFileConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RevokeFileConfig")
	return
}

func NewRevokeFileConfigResponse() (response *RevokeFileConfigResponse) {
	response = &RevokeFileConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 撤回已发布的文件配置
func (c *Client) RevokeFileConfig(request *RevokeFileConfigRequest) (response *RevokeFileConfigResponse, err error) {
	if request == nil {
		request = NewRevokeFileConfigRequest()
	}
	response = NewRevokeFileConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLaneRulesRequest() (request *DescribeLaneRulesRequest) {
	request = &DescribeLaneRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLaneRules")
	return
}

func NewDescribeLaneRulesResponse() (response *DescribeLaneRulesResponse) {
	response = &DescribeLaneRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询泳道规则列表
func (c *Client) DescribeLaneRules(request *DescribeLaneRulesRequest) (response *DescribeLaneRulesResponse, err error) {
	if request == nil {
		request = NewDescribeLaneRulesRequest()
	}
	response = NewDescribeLaneRulesResponse()
	err = c.Send(request, response)
	return
}

func NewFindMonitorObjectRequest() (request *FindMonitorObjectRequest) {
	request = &FindMonitorObjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "FindMonitorObject")
	return
}

func NewFindMonitorObjectResponse() (response *FindMonitorObjectResponse) {
	response = &FindMonitorObjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// FindMonitorObject
func (c *Client) FindMonitorObject(request *FindMonitorObjectRequest) (response *FindMonitorObjectResponse, err error) {
	if request == nil {
		request = NewFindMonitorObjectRequest()
	}
	response = NewFindMonitorObjectResponse()
	err = c.Send(request, response)
	return
}

func NewGetDumpRequest() (request *GetDumpRequest) {
	request = &GetDumpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetDump")
	return
}

func NewGetDumpResponse() (response *GetDumpResponse) {
	response = &GetDumpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetDump
func (c *Client) GetDump(request *GetDumpRequest) (response *GetDumpResponse, err error) {
	if request == nil {
		request = NewGetDumpRequest()
	}
	response = NewGetDumpResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteZoneRequest() (request *DeleteZoneRequest) {
	request = &DeleteZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteZone")
	return
}

func NewDeleteZoneResponse() (response *DeleteZoneResponse) {
	response = &DeleteZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除业务可用区
func (c *Client) DeleteZone(request *DeleteZoneRequest) (response *DeleteZoneResponse, err error) {
	if request == nil {
		request = NewDeleteZoneRequest()
	}
	response = NewDeleteZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAddibleInstancesRequest() (request *DescribeAddibleInstancesRequest) {
	request = &DescribeAddibleInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAddibleInstances")
	return
}

func NewDescribeAddibleInstancesResponse() (response *DescribeAddibleInstancesResponse) {
	response = &DescribeAddibleInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群可添加的机器列表
func (c *Client) DescribeAddibleInstances(request *DescribeAddibleInstancesRequest) (response *DescribeAddibleInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeAddibleInstancesRequest()
	}
	response = NewDescribeAddibleInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowLastBatchStateRequest() (request *DescribeFlowLastBatchStateRequest) {
	request = &DescribeFlowLastBatchStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowLastBatchState")
	return
}

func NewDescribeFlowLastBatchStateResponse() (response *DescribeFlowLastBatchStateResponse) {
	response = &DescribeFlowLastBatchStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询工作流最新一个批次的状态信息
func (c *Client) DescribeFlowLastBatchState(request *DescribeFlowLastBatchStateRequest) (response *DescribeFlowLastBatchStateResponse, err error) {
	if request == nil {
		request = NewDescribeFlowLastBatchStateRequest()
	}
	response = NewDescribeFlowLastBatchStateResponse()
	err = c.Send(request, response)
	return
}

func NewShrinkInstancesRequest() (request *ShrinkInstancesRequest) {
	request = &ShrinkInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ShrinkInstances")
	return
}

func NewShrinkInstancesResponse() (response *ShrinkInstancesResponse) {
	response = &ShrinkInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 虚拟机部署组下线实例
func (c *Client) ShrinkInstances(request *ShrinkInstancesRequest) (response *ShrinkInstancesResponse, err error) {
	if request == nil {
		request = NewShrinkInstancesRequest()
	}
	response = NewShrinkInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewShirkNamespaceRequest() (request *ShirkNamespaceRequest) {
	request = &ShirkNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ShirkNamespace")
	return
}

func NewShirkNamespaceResponse() (response *ShirkNamespaceResponse) {
	response = &ShirkNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 移除命名空间机器
func (c *Client) ShirkNamespace(request *ShirkNamespaceRequest) (response *ShirkNamespaceResponse, err error) {
	if request == nil {
		request = NewShirkNamespaceRequest()
	}
	response = NewShirkNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewShrinkinstanceScriptsRequest() (request *ShrinkinstanceScriptsRequest) {
	request = &ShrinkinstanceScriptsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ShrinkinstanceScripts")
	return
}

func NewShrinkinstanceScriptsResponse() (response *ShrinkinstanceScriptsResponse) {
	response = &ShrinkinstanceScriptsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除集群主机脚本
func (c *Client) ShrinkinstanceScripts(request *ShrinkinstanceScriptsRequest) (response *ShrinkinstanceScriptsResponse, err error) {
	if request == nil {
		request = NewShrinkinstanceScriptsRequest()
	}
	response = NewShrinkinstanceScriptsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMonitorStatisticsPolicyRequest() (request *DeleteMonitorStatisticsPolicyRequest) {
	request = &DeleteMonitorStatisticsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteMonitorStatisticsPolicy")
	return
}

func NewDeleteMonitorStatisticsPolicyResponse() (response *DeleteMonitorStatisticsPolicyResponse) {
	response = &DeleteMonitorStatisticsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除监控统计策略
func (c *Client) DeleteMonitorStatisticsPolicy(request *DeleteMonitorStatisticsPolicyRequest) (response *DeleteMonitorStatisticsPolicyResponse, err error) {
	if request == nil {
		request = NewDeleteMonitorStatisticsPolicyRequest()
	}
	response = NewDeleteMonitorStatisticsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroserviceProtocolRequest() (request *DescribeMicroserviceProtocolRequest) {
	request = &DescribeMicroserviceProtocolRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroserviceProtocol")
	return
}

func NewDescribeMicroserviceProtocolResponse() (response *DescribeMicroserviceProtocolResponse) {
	response = &DescribeMicroserviceProtocolResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微服务协议
func (c *Client) DescribeMicroserviceProtocol(request *DescribeMicroserviceProtocolRequest) (response *DescribeMicroserviceProtocolResponse, err error) {
	if request == nil {
		request = NewDescribeMicroserviceProtocolRequest()
	}
	response = NewDescribeMicroserviceProtocolResponse()
	err = c.Send(request, response)
	return
}

func NewListColudMonitorStatisticsPolicyRequest() (request *ListColudMonitorStatisticsPolicyRequest) {
	request = &ListColudMonitorStatisticsPolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListColudMonitorStatisticsPolicy")
	return
}

func NewListColudMonitorStatisticsPolicyResponse() (response *ListColudMonitorStatisticsPolicyResponse) {
	response = &ListColudMonitorStatisticsPolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ListColudMonitorStatisticsPolicy
func (c *Client) ListColudMonitorStatisticsPolicy(request *ListColudMonitorStatisticsPolicyRequest) (response *ListColudMonitorStatisticsPolicyResponse, err error) {
	if request == nil {
		request = NewListColudMonitorStatisticsPolicyRequest()
	}
	response = NewListColudMonitorStatisticsPolicyResponse()
	err = c.Send(request, response)
	return
}

func NewBindGatewayInstanceProgramsRequest() (request *BindGatewayInstanceProgramsRequest) {
	request = &BindGatewayInstanceProgramsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "BindGatewayInstancePrograms")
	return
}

func NewBindGatewayInstanceProgramsResponse() (response *BindGatewayInstanceProgramsResponse) {
	response = &BindGatewayInstanceProgramsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定网关实体与数据集
func (c *Client) BindGatewayInstancePrograms(request *BindGatewayInstanceProgramsRequest) (response *BindGatewayInstanceProgramsResponse, err error) {
	if request == nil {
		request = NewBindGatewayInstanceProgramsRequest()
	}
	response = NewBindGatewayInstanceProgramsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAndDownloadTemplateRequest() (request *CreateAndDownloadTemplateRequest) {
	request = &CreateAndDownloadTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateAndDownloadTemplate")
	return
}

func NewCreateAndDownloadTemplateResponse() (response *CreateAndDownloadTemplateResponse) {
	response = &CreateAndDownloadTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 保存并生成模板
func (c *Client) CreateAndDownloadTemplate(request *CreateAndDownloadTemplateRequest) (response *CreateAndDownloadTemplateResponse, err error) {
	if request == nil {
		request = NewCreateAndDownloadTemplateRequest()
	}
	response = NewCreateAndDownloadTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogCapacityRequest() (request *DescribeLogCapacityRequest) {
	request = &DescribeLogCapacityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLogCapacity")
	return
}

func NewDescribeLogCapacityResponse() (response *DescribeLogCapacityResponse) {
	response = &DescribeLogCapacityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户日志使用量
func (c *Client) DescribeLogCapacity(request *DescribeLogCapacityRequest) (response *DescribeLogCapacityResponse, err error) {
	if request == nil {
		request = NewDescribeLogCapacityRequest()
	}
	response = NewDescribeLogCapacityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskExecuteLogInfoRequest() (request *DescribeTaskExecuteLogInfoRequest) {
	request = &DescribeTaskExecuteLogInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskExecuteLogInfo")
	return
}

func NewDescribeTaskExecuteLogInfoResponse() (response *DescribeTaskExecuteLogInfoResponse) {
	response = &DescribeTaskExecuteLogInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务执行日志
func (c *Client) DescribeTaskExecuteLogInfo(request *DescribeTaskExecuteLogInfoRequest) (response *DescribeTaskExecuteLogInfoResponse, err error) {
	if request == nil {
		request = NewDescribeTaskExecuteLogInfoRequest()
	}
	response = NewDescribeTaskExecuteLogInfoResponse()
	err = c.Send(request, response)
	return
}

func NewGetUploadInfoRequest() (request *GetUploadInfoRequest) {
	request = &GetUploadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetUploadInfo")
	return
}

func NewGetUploadInfoResponse() (response *GetUploadInfoResponse) {
	response = &GetUploadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取上传信息
func (c *Client) GetUploadInfo(request *GetUploadInfoRequest) (response *GetUploadInfoResponse, err error) {
	if request == nil {
		request = NewGetUploadInfoRequest()
	}
	response = NewGetUploadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProductNewsRequest() (request *ModifyProductNewsRequest) {
	request = &ModifyProductNewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyProductNews")
	return
}

func NewModifyProductNewsResponse() (response *ModifyProductNewsResponse) {
	response = &ModifyProductNewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改产品动态
func (c *Client) ModifyProductNews(request *ModifyProductNewsRequest) (response *ModifyProductNewsResponse, err error) {
	if request == nil {
		request = NewModifyProductNewsRequest()
	}
	response = NewModifyProductNewsResponse()
	err = c.Send(request, response)
	return
}

func NewGetResourceBatchIndexRequest() (request *GetResourceBatchIndexRequest) {
	request = &GetResourceBatchIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetResourceBatchIndex")
	return
}

func NewGetResourceBatchIndexResponse() (response *GetResourceBatchIndexResponse) {
	response = &GetResourceBatchIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取资源批次索引
func (c *Client) GetResourceBatchIndex(request *GetResourceBatchIndexRequest) (response *GetResourceBatchIndexResponse, err error) {
	if request == nil {
		request = NewGetResourceBatchIndexRequest()
	}
	response = NewGetResourceBatchIndexResponse()
	err = c.Send(request, response)
	return
}

func NewCheckClusterCIDRRequest() (request *CheckClusterCIDRRequest) {
	request = &CheckClusterCIDRRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CheckClusterCIDR")
	return
}

func NewCheckClusterCIDRResponse() (response *CheckClusterCIDRResponse) {
	response = &CheckClusterCIDRResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查集群CIDR是否可用
func (c *Client) CheckClusterCIDR(request *CheckClusterCIDRRequest) (response *CheckClusterCIDRResponse, err error) {
	if request == nil {
		request = NewCheckClusterCIDRRequest()
	}
	response = NewCheckClusterCIDRResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAssociateRelationRequest() (request *DescribeAssociateRelationRequest) {
	request = &DescribeAssociateRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAssociateRelation")
	return
}

func NewDescribeAssociateRelationResponse() (response *DescribeAssociateRelationResponse) {
	response = &DescribeAssociateRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个部署组关联信息
func (c *Client) DescribeAssociateRelation(request *DescribeAssociateRelationRequest) (response *DescribeAssociateRelationResponse, err error) {
	if request == nil {
		request = NewDescribeAssociateRelationRequest()
	}
	response = NewDescribeAssociateRelationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePathRewriteRequest() (request *DescribePathRewriteRequest) {
	request = &DescribePathRewriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePathRewrite")
	return
}

func NewDescribePathRewriteResponse() (response *DescribePathRewriteResponse) {
	response = &DescribePathRewriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路径重写
func (c *Client) DescribePathRewrite(request *DescribePathRewriteRequest) (response *DescribePathRewriteResponse, err error) {
	if request == nil {
		request = NewDescribePathRewriteRequest()
	}
	response = NewDescribePathRewriteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayTagPluginRequest() (request *DescribeGatewayTagPluginRequest) {
	request = &DescribeGatewayTagPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayTagPlugin")
	return
}

func NewDescribeGatewayTagPluginResponse() (response *DescribeGatewayTagPluginResponse) {
	response = &DescribeGatewayTagPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Tag插件详情信息
func (c *Client) DescribeGatewayTagPlugin(request *DescribeGatewayTagPluginRequest) (response *DescribeGatewayTagPluginResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayTagPluginRequest()
	}
	response = NewDescribeGatewayTagPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupLaneRequest() (request *DescribeGroupLaneRequest) {
	request = &DescribeGroupLaneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupLane")
	return
}

func NewDescribeGroupLaneResponse() (response *DescribeGroupLaneResponse) {
	response = &DescribeGroupLaneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部署组的泳道信息
func (c *Client) DescribeGroupLane(request *DescribeGroupLaneRequest) (response *DescribeGroupLaneResponse, err error) {
	if request == nil {
		request = NewDescribeGroupLaneRequest()
	}
	response = NewDescribeGroupLaneResponse()
	err = c.Send(request, response)
	return
}

func NewListTsfModulesDetailRequest() (request *ListTsfModulesDetailRequest) {
	request = &ListTsfModulesDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListTsfModulesDetail")
	return
}

func NewListTsfModulesDetailResponse() (response *ListTsfModulesDetailResponse) {
	response = &ListTsfModulesDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 模块详情列表
func (c *Client) ListTsfModulesDetail(request *ListTsfModulesDetailRequest) (response *ListTsfModulesDetailResponse, err error) {
	if request == nil {
		request = NewListTsfModulesDetailRequest()
	}
	response = NewListTsfModulesDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGWOverviewInvocationRequest() (request *DescribeGWOverviewInvocationRequest) {
	request = &DescribeGWOverviewInvocationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGWOverviewInvocation")
	return
}

func NewDescribeGWOverviewInvocationResponse() (response *DescribeGWOverviewInvocationResponse) {
	response = &DescribeGWOverviewInvocationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询网关分组或者api监控数据
func (c *Client) DescribeGWOverviewInvocation(request *DescribeGWOverviewInvocationRequest) (response *DescribeGWOverviewInvocationResponse, err error) {
	if request == nil {
		request = NewDescribeGWOverviewInvocationRequest()
	}
	response = NewDescribeGWOverviewInvocationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInovcationIndicatorsRequest() (request *DescribeInovcationIndicatorsRequest) {
	request = &DescribeInovcationIndicatorsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInovcationIndicators")
	return
}

func NewDescribeInovcationIndicatorsResponse() (response *DescribeInovcationIndicatorsResponse) {
	response = &DescribeInovcationIndicatorsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用监控指标
func (c *Client) DescribeInovcationIndicators(request *DescribeInovcationIndicatorsRequest) (response *DescribeInovcationIndicatorsResponse, err error) {
	if request == nil {
		request = NewDescribeInovcationIndicatorsRequest()
	}
	response = NewDescribeInovcationIndicatorsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTransactionsRequest() (request *DescribeTransactionsRequest) {
	request = &DescribeTransactionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTransactions")
	return
}

func NewDescribeTransactionsResponse() (response *DescribeTransactionsResponse) {
	response = &DescribeTransactionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取事务列表
func (c *Client) DescribeTransactions(request *DescribeTransactionsRequest) (response *DescribeTransactionsResponse, err error) {
	if request == nil {
		request = NewDescribeTransactionsRequest()
	}
	response = NewDescribeTransactionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBusinessLogConfigRequest() (request *CreateBusinessLogConfigRequest) {
	request = &CreateBusinessLogConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateBusinessLogConfig")
	return
}

func NewCreateBusinessLogConfigResponse() (response *CreateBusinessLogConfigResponse) {
	response = &CreateBusinessLogConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建业务日志配置项
func (c *Client) CreateBusinessLogConfig(request *CreateBusinessLogConfigRequest) (response *CreateBusinessLogConfigResponse, err error) {
	if request == nil {
		request = NewCreateBusinessLogConfigRequest()
	}
	response = NewCreateBusinessLogConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUnitNamespacesRequest() (request *DeleteUnitNamespacesRequest) {
	request = &DeleteUnitNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteUnitNamespaces")
	return
}

func NewDeleteUnitNamespacesResponse() (response *DeleteUnitNamespacesResponse) {
	response = &DeleteUnitNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除单元化命名空间
func (c *Client) DeleteUnitNamespaces(request *DeleteUnitNamespacesRequest) (response *DeleteUnitNamespacesResponse, err error) {
	if request == nil {
		request = NewDeleteUnitNamespacesRequest()
	}
	response = NewDeleteUnitNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigRequest() (request *CreateConfigRequest) {
	request = &CreateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateConfig")
	return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
	response = &CreateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建配置项
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
	if request == nil {
		request = NewCreateConfigRequest()
	}
	response = NewCreateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePathRewritesRequest() (request *DescribePathRewritesRequest) {
	request = &DescribePathRewritesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePathRewrites")
	return
}

func NewDescribePathRewritesResponse() (response *DescribePathRewritesResponse) {
	response = &DescribePathRewritesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路径重写列表
func (c *Client) DescribePathRewrites(request *DescribePathRewritesRequest) (response *DescribePathRewritesResponse, err error) {
	if request == nil {
		request = NewDescribePathRewritesRequest()
	}
	response = NewDescribePathRewritesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateContainerGroupRequest() (request *UpdateContainerGroupRequest) {
	request = &UpdateContainerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateContainerGroup")
	return
}

func NewUpdateContainerGroupResponse() (response *UpdateContainerGroupResponse) {
	response = &UpdateContainerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改容器部署组
func (c *Client) UpdateContainerGroup(request *UpdateContainerGroupRequest) (response *UpdateContainerGroupResponse, err error) {
	if request == nil {
		request = NewUpdateContainerGroupRequest()
	}
	response = NewUpdateContainerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCheckTaskRuleRequest() (request *CheckTaskRuleRequest) {
	request = &CheckTaskRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CheckTaskRule")
	return
}

func NewCheckTaskRuleResponse() (response *CheckTaskRuleResponse) {
	response = &CheckTaskRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验任务触发规则是否合法
func (c *Client) CheckTaskRule(request *CheckTaskRuleRequest) (response *CheckTaskRuleResponse, err error) {
	if request == nil {
		request = NewCheckTaskRuleRequest()
	}
	response = NewCheckTaskRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTaskFlowRequest() (request *DeleteTaskFlowRequest) {
	request = &DeleteTaskFlowRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteTaskFlow")
	return
}

func NewDeleteTaskFlowResponse() (response *DeleteTaskFlowResponse) {
	response = &DeleteTaskFlowResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除工作流
func (c *Client) DeleteTaskFlow(request *DeleteTaskFlowRequest) (response *DeleteTaskFlowResponse, err error) {
	if request == nil {
		request = NewDeleteTaskFlowRequest()
	}
	response = NewDeleteTaskFlowResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAuthorizationRequest() (request *ModifyAuthorizationRequest) {
	request = &ModifyAuthorizationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyAuthorization")
	return
}

func NewModifyAuthorizationResponse() (response *ModifyAuthorizationResponse) {
	response = &ModifyAuthorizationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新微服务权限规则
func (c *Client) ModifyAuthorization(request *ModifyAuthorizationRequest) (response *ModifyAuthorizationResponse, err error) {
	if request == nil {
		request = NewModifyAuthorizationRequest()
	}
	response = NewModifyAuthorizationResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTsfZoneRequest() (request *CreateTsfZoneRequest) {
	request = &CreateTsfZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateTsfZone")
	return
}

func NewCreateTsfZoneResponse() (response *CreateTsfZoneResponse) {
	response = &CreateTsfZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建基础可用区
func (c *Client) CreateTsfZone(request *CreateTsfZoneRequest) (response *CreateTsfZoneResponse, err error) {
	if request == nil {
		request = NewCreateTsfZoneRequest()
	}
	response = NewCreateTsfZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCurrentAppIdRequest() (request *DescribeCurrentAppIdRequest) {
	request = &DescribeCurrentAppIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeCurrentAppId")
	return
}

func NewDescribeCurrentAppIdResponse() (response *DescribeCurrentAppIdResponse) {
	response = &DescribeCurrentAppIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取当前登录用户的AppId
func (c *Client) DescribeCurrentAppId(request *DescribeCurrentAppIdRequest) (response *DescribeCurrentAppIdResponse, err error) {
	if request == nil {
		request = NewDescribeCurrentAppIdRequest()
	}
	response = NewDescribeCurrentAppIdResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroserviceApiRequest() (request *DeleteMicroserviceApiRequest) {
	request = &DeleteMicroserviceApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteMicroserviceApi")
	return
}

func NewDeleteMicroserviceApiResponse() (response *DeleteMicroserviceApiResponse) {
	response = &DeleteMicroserviceApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除微服务API
func (c *Client) DeleteMicroserviceApi(request *DeleteMicroserviceApiRequest) (response *DeleteMicroserviceApiResponse, err error) {
	if request == nil {
		request = NewDeleteMicroserviceApiRequest()
	}
	response = NewDeleteMicroserviceApiResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvocationStatisticsRatioRequest() (request *DescribeInvocationStatisticsRatioRequest) {
	request = &DescribeInvocationStatisticsRatioRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationStatisticsRatio")
	return
}

func NewDescribeInvocationStatisticsRatioResponse() (response *DescribeInvocationStatisticsRatioResponse) {
	response = &DescribeInvocationStatisticsRatioResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用统计数据对比
func (c *Client) DescribeInvocationStatisticsRatio(request *DescribeInvocationStatisticsRatioRequest) (response *DescribeInvocationStatisticsRatioResponse, err error) {
	if request == nil {
		request = NewDescribeInvocationStatisticsRatioRequest()
	}
	response = NewDescribeInvocationStatisticsRatioResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerGroupPlainYamlRequest() (request *DescribeContainerGroupPlainYamlRequest) {
	request = &DescribeContainerGroupPlainYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupPlainYaml")
	return
}

func NewDescribeContainerGroupPlainYamlResponse() (response *DescribeContainerGroupPlainYamlResponse) {
	response = &DescribeContainerGroupPlainYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 这里的简单 Yaml 指剔除 TSF 信息后的 Yaml。和 DeployContainerGroupByPlainYaml 接口配合使用
func (c *Client) DescribeContainerGroupPlainYaml(request *DescribeContainerGroupPlainYamlRequest) (response *DescribeContainerGroupPlainYamlResponse, err error) {
	if request == nil {
		request = NewDescribeContainerGroupPlainYamlRequest()
	}
	response = NewDescribeContainerGroupPlainYamlResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerNamespacesRequest() (request *DescribeContainerNamespacesRequest) {
	request = &DescribeContainerNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerNamespaces")
	return
}

func NewDescribeContainerNamespacesResponse() (response *DescribeContainerNamespacesResponse) {
	response = &DescribeContainerNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器集群的命名空间列表
func (c *Client) DescribeContainerNamespaces(request *DescribeContainerNamespacesRequest) (response *DescribeContainerNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeContainerNamespacesRequest()
	}
	response = NewDescribeContainerNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteRulesRequest() (request *DescribeRouteRulesRequest) {
	request = &DescribeRouteRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRouteRules")
	return
}

func NewDescribeRouteRulesResponse() (response *DescribeRouteRulesResponse) {
	response = &DescribeRouteRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取路由规则列表
func (c *Client) DescribeRouteRules(request *DescribeRouteRulesRequest) (response *DescribeRouteRulesResponse, err error) {
	if request == nil {
		request = NewDescribeRouteRulesRequest()
	}
	response = NewDescribeRouteRulesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFilebeatConfigGroupRequest() (request *DescribeFilebeatConfigGroupRequest) {
	request = &DescribeFilebeatConfigGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFilebeatConfigGroup")
	return
}

func NewDescribeFilebeatConfigGroupResponse() (response *DescribeFilebeatConfigGroupResponse) {
	response = &DescribeFilebeatConfigGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询filbeat配置绑定的部署组
func (c *Client) DescribeFilebeatConfigGroup(request *DescribeFilebeatConfigGroupRequest) (response *DescribeFilebeatConfigGroupResponse, err error) {
	if request == nil {
		request = NewDescribeFilebeatConfigGroupRequest()
	}
	response = NewDescribeFilebeatConfigGroupResponse()
	err = c.Send(request, response)
	return
}

func NewSetUpKubeInjectRequest() (request *SetUpKubeInjectRequest) {
	request = &SetUpKubeInjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SetUpKubeInject")
	return
}

func NewSetUpKubeInjectResponse() (response *SetUpKubeInjectResponse) {
	response = &SetUpKubeInjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开启容器轻量化接入
func (c *Client) SetUpKubeInject(request *SetUpKubeInjectRequest) (response *SetUpKubeInjectResponse, err error) {
	if request == nil {
		request = NewSetUpKubeInjectRequest()
	}
	response = NewSetUpKubeInjectResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTaskRequest() (request *DeleteTaskRequest) {
	request = &DeleteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteTask")
	return
}

func NewDeleteTaskResponse() (response *DeleteTaskResponse) {
	response = &DeleteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除任务
func (c *Client) DeleteTask(request *DeleteTaskRequest) (response *DeleteTaskResponse, err error) {
	if request == nil {
		request = NewDeleteTaskRequest()
	}
	response = NewDeleteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJvmMonitorMethodProfileRequest() (request *DescribeJvmMonitorMethodProfileRequest) {
	request = &DescribeJvmMonitorMethodProfileRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeJvmMonitorMethodProfile")
	return
}

func NewDescribeJvmMonitorMethodProfileResponse() (response *DescribeJvmMonitorMethodProfileResponse) {
	response = &DescribeJvmMonitorMethodProfileResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询JVM监控的方法执行分析
func (c *Client) DescribeJvmMonitorMethodProfile(request *DescribeJvmMonitorMethodProfileRequest) (response *DescribeJvmMonitorMethodProfileResponse, err error) {
	if request == nil {
		request = NewDescribeJvmMonitorMethodProfileRequest()
	}
	response = NewDescribeJvmMonitorMethodProfileResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNamespaceRequest() (request *DeleteNamespaceRequest) {
	request = &DeleteNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteNamespace")
	return
}

func NewDeleteNamespaceResponse() (response *DeleteNamespaceResponse) {
	response = &DeleteNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除命名空间
func (c *Client) DeleteNamespace(request *DeleteNamespaceRequest) (response *DeleteNamespaceResponse, err error) {
	if request == nil {
		request = NewDeleteNamespaceRequest()
	}
	response = NewDeleteNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewBillingOperationPurchaseRequest() (request *BillingOperationPurchaseRequest) {
	request = &BillingOperationPurchaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "BillingOperationPurchase")
	return
}

func NewBillingOperationPurchaseResponse() (response *BillingOperationPurchaseResponse) {
	response = &BillingOperationPurchaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费运营端购买
func (c *Client) BillingOperationPurchase(request *BillingOperationPurchaseRequest) (response *BillingOperationPurchaseResponse, err error) {
	if request == nil {
		request = NewBillingOperationPurchaseRequest()
	}
	response = NewBillingOperationPurchaseResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupsRequest() (request *DescribeGroupsRequest) {
	request = &DescribeGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroups")
	return
}

func NewDescribeGroupsResponse() (response *DescribeGroupsResponse) {
	response = &DescribeGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取虚拟机部署组列表
func (c *Client) DescribeGroups(request *DescribeGroupsRequest) (response *DescribeGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeGroupsRequest()
	}
	response = NewDescribeGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeployContainerGroupByYamlRequest() (request *DeployContainerGroupByYamlRequest) {
	request = &DeployContainerGroupByYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeployContainerGroupByYaml")
	return
}

func NewDeployContainerGroupByYamlResponse() (response *DeployContainerGroupByYamlResponse) {
	response = &DeployContainerGroupByYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过Yaml部署Tsf容器部署组
func (c *Client) DeployContainerGroupByYaml(request *DeployContainerGroupByYamlRequest) (response *DeployContainerGroupByYamlResponse, err error) {
	if request == nil {
		request = NewDeployContainerGroupByYamlRequest()
	}
	response = NewDeployContainerGroupByYamlResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigTemplateRequest() (request *CreateConfigTemplateRequest) {
	request = &CreateConfigTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateConfigTemplate")
	return
}

func NewCreateConfigTemplateResponse() (response *CreateConfigTemplateResponse) {
	response = &CreateConfigTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建参数模板
func (c *Client) CreateConfigTemplate(request *CreateConfigTemplateRequest) (response *CreateConfigTemplateResponse, err error) {
	if request == nil {
		request = NewCreateConfigTemplateRequest()
	}
	response = NewCreateConfigTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewExecuteTaskRequest() (request *ExecuteTaskRequest) {
	request = &ExecuteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ExecuteTask")
	return
}

func NewExecuteTaskResponse() (response *ExecuteTaskResponse) {
	response = &ExecuteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 手动执行一次任务。
func (c *Client) ExecuteTask(request *ExecuteTaskRequest) (response *ExecuteTaskResponse, err error) {
	if request == nil {
		request = NewExecuteTaskRequest()
	}
	response = NewExecuteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewChangeContainerReplicasRequest() (request *ChangeContainerReplicasRequest) {
	request = &ChangeContainerReplicasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ChangeContainerReplicas")
	return
}

func NewChangeContainerReplicasResponse() (response *ChangeContainerReplicasResponse) {
	response = &ChangeContainerReplicasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改容器部署组实例数
func (c *Client) ChangeContainerReplicas(request *ChangeContainerReplicasRequest) (response *ChangeContainerReplicasResponse, err error) {
	if request == nil {
		request = NewChangeContainerReplicasRequest()
	}
	response = NewChangeContainerReplicasResponse()
	err = c.Send(request, response)
	return
}

func NewListManagerJobInfoRequest() (request *ListManagerJobInfoRequest) {
	request = &ListManagerJobInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListManagerJobInfo")
	return
}

func NewListManagerJobInfoResponse() (response *ListManagerJobInfoResponse) {
	response = &ListManagerJobInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端任务管理列表
func (c *Client) ListManagerJobInfo(request *ListManagerJobInfoRequest) (response *ListManagerJobInfoResponse, err error) {
	if request == nil {
		request = NewListManagerJobInfoRequest()
	}
	response = NewListManagerJobInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateGatewayApiRequest() (request *CreateGatewayApiRequest) {
	request = &CreateGatewayApiRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateGatewayApi")
	return
}

func NewCreateGatewayApiResponse() (response *CreateGatewayApiResponse) {
	response = &CreateGatewayApiResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入API至api分组(也支持新建API到分组)
func (c *Client) CreateGatewayApi(request *CreateGatewayApiRequest) (response *CreateGatewayApiResponse, err error) {
	if request == nil {
		request = NewCreateGatewayApiRequest()
	}
	response = NewCreateGatewayApiResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateUploadInfoRequest() (request *UpdateUploadInfoRequest) {
	request = &UpdateUploadInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateUploadInfo")
	return
}

func NewUpdateUploadInfoResponse() (response *UpdateUploadInfoResponse) {
	response = &UpdateUploadInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 调用TSF的GetUploadInfo接口和COS的上传接口后，需要调用此接口更新TSF中保存的程序包状态。
// 调用此接口完成后，才标志上传包流程结束。
func (c *Client) UpdateUploadInfo(request *UpdateUploadInfoRequest) (response *UpdateUploadInfoResponse, err error) {
	if request == nil {
		request = NewUpdateUploadInfoRequest()
	}
	response = NewUpdateUploadInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBasicResourceUsageRequest() (request *DescribeBasicResourceUsageRequest) {
	request = &DescribeBasicResourceUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeBasicResourceUsage")
	return
}

func NewDescribeBasicResourceUsageResponse() (response *DescribeBasicResourceUsageResponse) {
	response = &DescribeBasicResourceUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF基本资源信息概览接口
func (c *Client) DescribeBasicResourceUsage(request *DescribeBasicResourceUsageRequest) (response *DescribeBasicResourceUsageResponse, err error) {
	if request == nil {
		request = NewDescribeBasicResourceUsageRequest()
	}
	response = NewDescribeBasicResourceUsageResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRealtimeStdoutLogRequest() (request *SearchRealtimeStdoutLogRequest) {
	request = &SearchRealtimeStdoutLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchRealtimeStdoutLog")
	return
}

func NewSearchRealtimeStdoutLogResponse() (response *SearchRealtimeStdoutLogResponse) {
	response = &SearchRealtimeStdoutLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标准输出日志实时检索
func (c *Client) SearchRealtimeStdoutLog(request *SearchRealtimeStdoutLogRequest) (response *SearchRealtimeStdoutLogResponse, err error) {
	if request == nil {
		request = NewSearchRealtimeStdoutLogRequest()
	}
	response = NewSearchRealtimeStdoutLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeImageRepositoryRequest() (request *DescribeImageRepositoryRequest) {
	request = &DescribeImageRepositoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeImageRepository")
	return
}

func NewDescribeImageRepositoryResponse() (response *DescribeImageRepositoryResponse) {
	response = &DescribeImageRepositoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像仓库列表
func (c *Client) DescribeImageRepository(request *DescribeImageRepositoryRequest) (response *DescribeImageRepositoryResponse, err error) {
	if request == nil {
		request = NewDescribeImageRepositoryRequest()
	}
	response = NewDescribeImageRepositoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageTagRequest() (request *DeleteImageTagRequest) {
	request = &DeleteImageTagRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteImageTag")
	return
}

func NewDeleteImageTagResponse() (response *DeleteImageTagResponse) {
	response = &DeleteImageTagResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除镜像版本
func (c *Client) DeleteImageTag(request *DeleteImageTagRequest) (response *DeleteImageTagResponse, err error) {
	if request == nil {
		request = NewDeleteImageTagRequest()
	}
	response = NewDeleteImageTagResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRouteRuleRequest() (request *DeleteRouteRuleRequest) {
	request = &DeleteRouteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteRouteRule")
	return
}

func NewDeleteRouteRuleResponse() (response *DeleteRouteRuleResponse) {
	response = &DeleteRouteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除路由规则
func (c *Client) DeleteRouteRule(request *DeleteRouteRuleRequest) (response *DeleteRouteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteRouteRuleRequest()
	}
	response = NewDeleteRouteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyContainerGroupRequest() (request *ModifyContainerGroupRequest) {
	request = &ModifyContainerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyContainerGroup")
	return
}

func NewModifyContainerGroupResponse() (response *ModifyContainerGroupResponse) {
	response = &ModifyContainerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改容器部署组
func (c *Client) ModifyContainerGroup(request *ModifyContainerGroupRequest) (response *ModifyContainerGroupResponse, err error) {
	if request == nil {
		request = NewModifyContainerGroupRequest()
	}
	response = NewModifyContainerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigTemplateRequest() (request *DescribeConfigTemplateRequest) {
	request = &DescribeConfigTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeConfigTemplate")
	return
}

func NewDescribeConfigTemplateResponse() (response *DescribeConfigTemplateResponse) {
	response = &DescribeConfigTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导入配置
func (c *Client) DescribeConfigTemplate(request *DescribeConfigTemplateRequest) (response *DescribeConfigTemplateResponse, err error) {
	if request == nil {
		request = NewDescribeConfigTemplateRequest()
	}
	response = NewDescribeConfigTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllBillingLicensesRequest() (request *DescribeAllBillingLicensesRequest) {
	request = &DescribeAllBillingLicensesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAllBillingLicenses")
	return
}

func NewDescribeAllBillingLicensesResponse() (response *DescribeAllBillingLicensesResponse) {
	response = &DescribeAllBillingLicensesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询所有计费许可
func (c *Client) DescribeAllBillingLicenses(request *DescribeAllBillingLicensesRequest) (response *DescribeAllBillingLicensesResponse, err error) {
	if request == nil {
		request = NewDescribeAllBillingLicensesRequest()
	}
	response = NewDescribeAllBillingLicensesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupReleaseRequest() (request *DescribeGroupReleaseRequest) {
	request = &DescribeGroupReleaseRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupRelease")
	return
}

func NewDescribeGroupReleaseResponse() (response *DescribeGroupReleaseResponse) {
	response = &DescribeGroupReleaseResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部署组相关的发布信息
func (c *Client) DescribeGroupRelease(request *DescribeGroupReleaseRequest) (response *DescribeGroupReleaseResponse, err error) {
	if request == nil {
		request = NewDescribeGroupReleaseRequest()
	}
	response = NewDescribeGroupReleaseResponse()
	err = c.Send(request, response)
	return
}

func NewRegisterImageUserRequest() (request *RegisterImageUserRequest) {
	request = &RegisterImageUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RegisterImageUser")
	return
}

func NewRegisterImageUserResponse() (response *RegisterImageUserResponse) {
	response = &RegisterImageUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 镜像用户注册
func (c *Client) RegisterImageUser(request *RegisterImageUserRequest) (response *RegisterImageUserResponse, err error) {
	if request == nil {
		request = NewRegisterImageUserRequest()
	}
	response = NewRegisterImageUserResponse()
	err = c.Send(request, response)
	return
}

func NewDisableLaneRuleRequest() (request *DisableLaneRuleRequest) {
	request = &DisableLaneRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisableLaneRule")
	return
}

func NewDisableLaneRuleResponse() (response *DisableLaneRuleResponse) {
	response = &DisableLaneRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 禁用泳道规则
func (c *Client) DisableLaneRule(request *DisableLaneRuleRequest) (response *DisableLaneRuleResponse, err error) {
	if request == nil {
		request = NewDisableLaneRuleRequest()
	}
	response = NewDisableLaneRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseApplicationContentRequest() (request *DescribeLicenseApplicationContentRequest) {
	request = &DescribeLicenseApplicationContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLicenseApplicationContent")
	return
}

func NewDescribeLicenseApplicationContentResponse() (response *DescribeLicenseApplicationContentResponse) {
	response = &DescribeLicenseApplicationContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) DescribeLicenseApplicationContent(request *DescribeLicenseApplicationContentRequest) (response *DescribeLicenseApplicationContentResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseApplicationContentRequest()
	}
	response = NewDescribeLicenseApplicationContentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskBatchHistoryRecordsRequest() (request *DescribeTaskBatchHistoryRecordsRequest) {
	request = &DescribeTaskBatchHistoryRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskBatchHistoryRecords")
	return
}

func NewDescribeTaskBatchHistoryRecordsResponse() (response *DescribeTaskBatchHistoryRecordsResponse) {
	response = &DescribeTaskBatchHistoryRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看任务批次执行历史
func (c *Client) DescribeTaskBatchHistoryRecords(request *DescribeTaskBatchHistoryRecordsRequest) (response *DescribeTaskBatchHistoryRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskBatchHistoryRecordsRequest()
	}
	response = NewDescribeTaskBatchHistoryRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteProgramRequest() (request *DeleteProgramRequest) {
	request = &DeleteProgramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteProgram")
	return
}

func NewDeleteProgramResponse() (response *DeleteProgramResponse) {
	response = &DeleteProgramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除数据集
func (c *Client) DeleteProgram(request *DeleteProgramRequest) (response *DeleteProgramResponse, err error) {
	if request == nil {
		request = NewDeleteProgramRequest()
	}
	response = NewDeleteProgramResponse()
	err = c.Send(request, response)
	return
}

func NewGetApmEsAuthInfoRequest() (request *GetApmEsAuthInfoRequest) {
	request = &GetApmEsAuthInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetApmEsAuthInfo")
	return
}

func NewGetApmEsAuthInfoResponse() (response *GetApmEsAuthInfoResponse) {
	response = &GetApmEsAuthInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用性能管理ES权限信息
func (c *Client) GetApmEsAuthInfo(request *GetApmEsAuthInfoRequest) (response *GetApmEsAuthInfoResponse, err error) {
	if request == nil {
		request = NewGetApmEsAuthInfoRequest()
	}
	response = NewGetApmEsAuthInfoResponse()
	err = c.Send(request, response)
	return
}

func NewEnableCircuitBreakerRuleRequest() (request *EnableCircuitBreakerRuleRequest) {
	request = &EnableCircuitBreakerRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableCircuitBreakerRule")
	return
}

func NewEnableCircuitBreakerRuleResponse() (response *EnableCircuitBreakerRuleResponse) {
	response = &EnableCircuitBreakerRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用熔断规则
func (c *Client) EnableCircuitBreakerRule(request *EnableCircuitBreakerRuleRequest) (response *EnableCircuitBreakerRuleResponse, err error) {
	if request == nil {
		request = NewEnableCircuitBreakerRuleRequest()
	}
	response = NewEnableCircuitBreakerRuleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyZoneRequest() (request *ModifyZoneRequest) {
	request = &ModifyZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyZone")
	return
}

func NewModifyZoneResponse() (response *ModifyZoneResponse) {
	response = &ModifyZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改业务可用区
func (c *Client) ModifyZone(request *ModifyZoneRequest) (response *ModifyZoneResponse, err error) {
	if request == nil {
		request = NewModifyZoneRequest()
	}
	response = NewModifyZoneResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveInstancesRequest() (request *RemoveInstancesRequest) {
	request = &RemoveInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RemoveInstances")
	return
}

func NewRemoveInstancesResponse() (response *RemoveInstancesResponse) {
	response = &RemoveInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 从 TSF 集群中批量移除云主机节点
func (c *Client) RemoveInstances(request *RemoveInstancesRequest) (response *RemoveInstancesResponse, err error) {
	if request == nil {
		request = NewRemoveInstancesRequest()
	}
	response = NewRemoveInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAuthorizationMicroservicesRequest() (request *DescribeAuthorizationMicroservicesRequest) {
	request = &DescribeAuthorizationMicroservicesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAuthorizationMicroservices")
	return
}

func NewDescribeAuthorizationMicroservicesResponse() (response *DescribeAuthorizationMicroservicesResponse) {
	response = &DescribeAuthorizationMicroservicesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询服务权限的微服务列表
func (c *Client) DescribeAuthorizationMicroservices(request *DescribeAuthorizationMicroservicesRequest) (response *DescribeAuthorizationMicroservicesResponse, err error) {
	if request == nil {
		request = NewDescribeAuthorizationMicroservicesRequest()
	}
	response = NewDescribeAuthorizationMicroservicesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskFlowsRequest() (request *DescribeTaskFlowsRequest) {
	request = &DescribeTaskFlowsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskFlows")
	return
}

func NewDescribeTaskFlowsResponse() (response *DescribeTaskFlowsResponse) {
	response = &DescribeTaskFlowsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 翻页查询租户的工作流列表数据
func (c *Client) DescribeTaskFlows(request *DescribeTaskFlowsRequest) (response *DescribeTaskFlowsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskFlowsRequest()
	}
	response = NewDescribeTaskFlowsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchSpanRequest() (request *SearchSpanRequest) {
	request = &SearchSpanRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchSpan")
	return
}

func NewSearchSpanResponse() (response *SearchSpanResponse) {
	response = &SearchSpanResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// span查询
func (c *Client) SearchSpan(request *SearchSpanRequest) (response *SearchSpanResponse, err error) {
	if request == nil {
		request = NewSearchSpanRequest()
	}
	response = NewSearchSpanResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupsWithPluginRequest() (request *DescribeGroupsWithPluginRequest) {
	request = &DescribeGroupsWithPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupsWithPlugin")
	return
}

func NewDescribeGroupsWithPluginResponse() (response *DescribeGroupsWithPluginResponse) {
	response = &DescribeGroupsWithPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个插件下绑定或未绑定的API分组
func (c *Client) DescribeGroupsWithPlugin(request *DescribeGroupsWithPluginRequest) (response *DescribeGroupsWithPluginResponse, err error) {
	if request == nil {
		request = NewDescribeGroupsWithPluginRequest()
	}
	response = NewDescribeGroupsWithPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInterfaceListRequest() (request *DescribeInterfaceListRequest) {
	request = &DescribeInterfaceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInterfaceList")
	return
}

func NewDescribeInterfaceListResponse() (response *DescribeInterfaceListResponse) {
	response = &DescribeInterfaceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取微服务下接口列表，配置接口告警使用
func (c *Client) DescribeInterfaceList(request *DescribeInterfaceListRequest) (response *DescribeInterfaceListResponse, err error) {
	if request == nil {
		request = NewDescribeInterfaceListRequest()
	}
	response = NewDescribeInterfaceListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBusinessLogConfigsRequest() (request *DescribeBusinessLogConfigsRequest) {
	request = &DescribeBusinessLogConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeBusinessLogConfigs")
	return
}

func NewDescribeBusinessLogConfigsResponse() (response *DescribeBusinessLogConfigsResponse) {
	response = &DescribeBusinessLogConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志配置项列表
func (c *Client) DescribeBusinessLogConfigs(request *DescribeBusinessLogConfigsRequest) (response *DescribeBusinessLogConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeBusinessLogConfigsRequest()
	}
	response = NewDescribeBusinessLogConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewReleaseProductNewsRequest() (request *ReleaseProductNewsRequest) {
	request = &ReleaseProductNewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ReleaseProductNews")
	return
}

func NewReleaseProductNewsResponse() (response *ReleaseProductNewsResponse) {
	response = &ReleaseProductNewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 全量覆盖发布表数据
func (c *Client) ReleaseProductNews(request *ReleaseProductNewsRequest) (response *ReleaseProductNewsResponse, err error) {
	if request == nil {
		request = NewReleaseProductNewsRequest()
	}
	response = NewReleaseProductNewsResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindApiGroupRequest() (request *UnbindApiGroupRequest) {
	request = &UnbindApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UnbindApiGroup")
	return
}

func NewUnbindApiGroupResponse() (response *UnbindApiGroupResponse) {
	response = &UnbindApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// API分组批量与网关解绑
func (c *Client) UnbindApiGroup(request *UnbindApiGroupRequest) (response *UnbindApiGroupResponse, err error) {
	if request == nil {
		request = NewUnbindApiGroupRequest()
	}
	response = NewUnbindApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewListApplicationServersRequest() (request *ListApplicationServersRequest) {
	request = &ListApplicationServersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListApplicationServers")
	return
}

func NewListApplicationServersResponse() (response *ListApplicationServersResponse) {
	response = &ListApplicationServersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用服务器列表
func (c *Client) ListApplicationServers(request *ListApplicationServersRequest) (response *ListApplicationServersResponse, err error) {
	if request == nil {
		request = NewListApplicationServersRequest()
	}
	response = NewListApplicationServersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNamespaceRequest() (request *ModifyNamespaceRequest) {
	request = &ModifyNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyNamespace")
	return
}

func NewModifyNamespaceResponse() (response *ModifyNamespaceResponse) {
	response = &ModifyNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改命名空间
func (c *Client) ModifyNamespace(request *ModifyNamespaceRequest) (response *ModifyNamespaceResponse, err error) {
	if request == nil {
		request = NewModifyNamespaceRequest()
	}
	response = NewModifyNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroserviceTotalInstanceCountRequest() (request *DescribeMicroserviceTotalInstanceCountRequest) {
	request = &DescribeMicroserviceTotalInstanceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroserviceTotalInstanceCount")
	return
}

func NewDescribeMicroserviceTotalInstanceCountResponse() (response *DescribeMicroserviceTotalInstanceCountResponse) {
	response = &DescribeMicroserviceTotalInstanceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询微服务下所有部署组的总实例数
func (c *Client) DescribeMicroserviceTotalInstanceCount(request *DescribeMicroserviceTotalInstanceCountRequest) (response *DescribeMicroserviceTotalInstanceCountResponse, err error) {
	if request == nil {
		request = NewDescribeMicroserviceTotalInstanceCountRequest()
	}
	response = NewDescribeMicroserviceTotalInstanceCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeliveryConfigsRequest() (request *DescribeDeliveryConfigsRequest) {
	request = &DescribeDeliveryConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeliveryConfigs")
	return
}

func NewDescribeDeliveryConfigsResponse() (response *DescribeDeliveryConfigsResponse) {
	response = &DescribeDeliveryConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取多个投递项配置
func (c *Client) DescribeDeliveryConfigs(request *DescribeDeliveryConfigsRequest) (response *DescribeDeliveryConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeDeliveryConfigsRequest()
	}
	response = NewDescribeDeliveryConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvocationMetricDataCurveRequest() (request *DescribeInvocationMetricDataCurveRequest) {
	request = &DescribeInvocationMetricDataCurveRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationMetricDataCurve")
	return
}

func NewDescribeInvocationMetricDataCurveResponse() (response *DescribeInvocationMetricDataCurveResponse) {
	response = &DescribeInvocationMetricDataCurveResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用指标数据变化曲线
func (c *Client) DescribeInvocationMetricDataCurve(request *DescribeInvocationMetricDataCurveRequest) (response *DescribeInvocationMetricDataCurveResponse, err error) {
	if request == nil {
		request = NewDescribeInvocationMetricDataCurveRequest()
	}
	response = NewDescribeInvocationMetricDataCurveResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterUsageRequest() (request *DescribeClusterUsageRequest) {
	request = &DescribeClusterUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterUsage")
	return
}

func NewDescribeClusterUsageResponse() (response *DescribeClusterUsageResponse) {
	response = &DescribeClusterUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群使用率
func (c *Client) DescribeClusterUsage(request *DescribeClusterUsageRequest) (response *DescribeClusterUsageResponse, err error) {
	if request == nil {
		request = NewDescribeClusterUsageRequest()
	}
	response = NewDescribeClusterUsageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMeshSidecarServiceInfoRequest() (request *DescribeMeshSidecarServiceInfoRequest) {
	request = &DescribeMeshSidecarServiceInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMeshSidecarServiceInfo")
	return
}

func NewDescribeMeshSidecarServiceInfoResponse() (response *DescribeMeshSidecarServiceInfoResponse) {
	response = &DescribeMeshSidecarServiceInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询mesh注册信息
func (c *Client) DescribeMeshSidecarServiceInfo(request *DescribeMeshSidecarServiceInfoRequest) (response *DescribeMeshSidecarServiceInfoResponse, err error) {
	if request == nil {
		request = NewDescribeMeshSidecarServiceInfoRequest()
	}
	response = NewDescribeMeshSidecarServiceInfoResponse()
	err = c.Send(request, response)
	return
}

func NewFindDeployModuleParamsRequest() (request *FindDeployModuleParamsRequest) {
	request = &FindDeployModuleParamsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "FindDeployModuleParams")
	return
}

func NewFindDeployModuleParamsResponse() (response *FindDeployModuleParamsResponse) {
	response = &FindDeployModuleParamsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询参数配置
func (c *Client) FindDeployModuleParams(request *FindDeployModuleParamsRequest) (response *FindDeployModuleParamsResponse, err error) {
	if request == nil {
		request = NewFindDeployModuleParamsRequest()
	}
	response = NewFindDeployModuleParamsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApiGroupRequest() (request *UpdateApiGroupRequest) {
	request = &UpdateApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateApiGroup")
	return
}

func NewUpdateApiGroupResponse() (response *UpdateApiGroupResponse) {
	response = &UpdateApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Api分组
func (c *Client) UpdateApiGroup(request *UpdateApiGroupRequest) (response *UpdateApiGroupResponse, err error) {
	if request == nil {
		request = NewUpdateApiGroupRequest()
	}
	response = NewUpdateApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewSearchRealtimeMeshLogRequest() (request *SearchRealtimeMeshLogRequest) {
	request = &SearchRealtimeMeshLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchRealtimeMeshLog")
	return
}

func NewSearchRealtimeMeshLogResponse() (response *SearchRealtimeMeshLogResponse) {
	response = &SearchRealtimeMeshLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 标准输出日志实时检索
func (c *Client) SearchRealtimeMeshLog(request *SearchRealtimeMeshLogRequest) (response *SearchRealtimeMeshLogResponse, err error) {
	if request == nil {
		request = NewSearchRealtimeMeshLogRequest()
	}
	response = NewSearchRealtimeMeshLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeManagerLicensesRequest() (request *DescribeManagerLicensesRequest) {
	request = &DescribeManagerLicensesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeManagerLicenses")
	return
}

func NewDescribeManagerLicensesResponse() (response *DescribeManagerLicensesResponse) {
	response = &DescribeManagerLicensesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询许可列表
func (c *Client) DescribeManagerLicenses(request *DescribeManagerLicensesRequest) (response *DescribeManagerLicensesResponse, err error) {
	if request == nil {
		request = NewDescribeManagerLicensesRequest()
	}
	response = NewDescribeManagerLicensesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInterfaceRequestRequest() (request *DescribeInterfaceRequestRequest) {
	request = &DescribeInterfaceRequestRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInterfaceRequest")
	return
}

func NewDescribeInterfaceRequestResponse() (response *DescribeInterfaceRequestResponse) {
	response = &DescribeInterfaceRequestResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口请求详情
func (c *Client) DescribeInterfaceRequest(request *DescribeInterfaceRequestRequest) (response *DescribeInterfaceRequestResponse, err error) {
	if request == nil {
		request = NewDescribeInterfaceRequestRequest()
	}
	response = NewDescribeInterfaceRequestResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterVersionsRequest() (request *DescribeClusterVersionsRequest) {
	request = &DescribeClusterVersionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterVersions")
	return
}

func NewDescribeClusterVersionsResponse() (response *DescribeClusterVersionsResponse) {
	response = &DescribeClusterVersionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看支持的容器集群版本列表
func (c *Client) DescribeClusterVersions(request *DescribeClusterVersionsRequest) (response *DescribeClusterVersionsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterVersionsRequest()
	}
	response = NewDescribeClusterVersionsResponse()
	err = c.Send(request, response)
	return
}

func NewListContainerTaskRequest() (request *ListContainerTaskRequest) {
	request = &ListContainerTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListContainerTask")
	return
}

func NewListContainerTaskResponse() (response *ListContainerTaskResponse) {
	response = &ListContainerTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更记录任务列表
func (c *Client) ListContainerTask(request *ListContainerTaskRequest) (response *ListContainerTaskResponse, err error) {
	if request == nil {
		request = NewListContainerTaskRequest()
	}
	response = NewListContainerTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupResourceUsageRequest() (request *DescribeGroupResourceUsageRequest) {
	request = &DescribeGroupResourceUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupResourceUsage")
	return
}

func NewDescribeGroupResourceUsageResponse() (response *DescribeGroupResourceUsageResponse) {
	response = &DescribeGroupResourceUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 部署组统计信息
func (c *Client) DescribeGroupResourceUsage(request *DescribeGroupResourceUsageRequest) (response *DescribeGroupResourceUsageResponse, err error) {
	if request == nil {
		request = NewDescribeGroupResourceUsageRequest()
	}
	response = NewDescribeGroupResourceUsageResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteImageTagsRequest() (request *DeleteImageTagsRequest) {
	request = &DeleteImageTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteImageTags")
	return
}

func NewDeleteImageTagsResponse() (response *DeleteImageTagsResponse) {
	response = &DeleteImageTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除镜像版本
func (c *Client) DeleteImageTags(request *DeleteImageTagsRequest) (response *DeleteImageTagsResponse, err error) {
	if request == nil {
		request = NewDeleteImageTagsRequest()
	}
	response = NewDeleteImageTagsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeThreadDetailRequest() (request *DescribeThreadDetailRequest) {
	request = &DescribeThreadDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeThreadDetail")
	return
}

func NewDescribeThreadDetailResponse() (response *DescribeThreadDetailResponse) {
	response = &DescribeThreadDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询java实例线程状态以及堆栈列表
func (c *Client) DescribeThreadDetail(request *DescribeThreadDetailRequest) (response *DescribeThreadDetailResponse, err error) {
	if request == nil {
		request = NewDescribeThreadDetailRequest()
	}
	response = NewDescribeThreadDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUnitRuleRequest() (request *CreateUnitRuleRequest) {
	request = &CreateUnitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateUnitRule")
	return
}

func NewCreateUnitRuleResponse() (response *CreateUnitRuleResponse) {
	response = &CreateUnitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建单元化规则
func (c *Client) CreateUnitRule(request *CreateUnitRuleRequest) (response *CreateUnitRuleResponse, err error) {
	if request == nil {
		request = NewCreateUnitRuleRequest()
	}
	response = NewCreateUnitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRegionRequest() (request *DeleteRegionRequest) {
	request = &DeleteRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteRegion")
	return
}

func NewDeleteRegionResponse() (response *DeleteRegionResponse) {
	response = &DeleteRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除业务区域
func (c *Client) DeleteRegion(request *DeleteRegionRequest) (response *DeleteRegionResponse, err error) {
	if request == nil {
		request = NewDeleteRegionRequest()
	}
	response = NewDeleteRegionResponse()
	err = c.Send(request, response)
	return
}

func NewStartGroupRequest() (request *StartGroupRequest) {
	request = &StartGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StartGroup")
	return
}

func NewStartGroupResponse() (response *StartGroupResponse) {
	response = &StartGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动分组
func (c *Client) StartGroup(request *StartGroupRequest) (response *StartGroupResponse, err error) {
	if request == nil {
		request = NewStartGroupRequest()
	}
	response = NewStartGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRegionsRequest() (request *DescribeRegionsRequest) {
	request = &DescribeRegionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRegions")
	return
}

func NewDescribeRegionsResponse() (response *DescribeRegionsResponse) {
	response = &DescribeRegionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询业务地域
func (c *Client) DescribeRegions(request *DescribeRegionsRequest) (response *DescribeRegionsResponse, err error) {
	if request == nil {
		request = NewDescribeRegionsRequest()
	}
	response = NewDescribeRegionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterCpuTypeRequest() (request *DescribeClusterCpuTypeRequest) {
	request = &DescribeClusterCpuTypeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterCpuType")
	return
}

func NewDescribeClusterCpuTypeResponse() (response *DescribeClusterCpuTypeResponse) {
	response = &DescribeClusterCpuTypeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群cpu架构
func (c *Client) DescribeClusterCpuType(request *DescribeClusterCpuTypeRequest) (response *DescribeClusterCpuTypeResponse, err error) {
	if request == nil {
		request = NewDescribeClusterCpuTypeRequest()
	}
	response = NewDescribeClusterCpuTypeResponse()
	err = c.Send(request, response)
	return
}

func NewValidateNamespaceClusterVPCRequest() (request *ValidateNamespaceClusterVPCRequest) {
	request = &ValidateNamespaceClusterVPCRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ValidateNamespaceClusterVPC")
	return
}

func NewValidateNamespaceClusterVPCResponse() (response *ValidateNamespaceClusterVPCResponse) {
	response = &ValidateNamespaceClusterVPCResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验命名空间归属集群VPC一致性
func (c *Client) ValidateNamespaceClusterVPC(request *ValidateNamespaceClusterVPCRequest) (response *ValidateNamespaceClusterVPCResponse, err error) {
	if request == nil {
		request = NewValidateNamespaceClusterVPCRequest()
	}
	response = NewValidateNamespaceClusterVPCResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceUsageRateRequest() (request *DescribeResourceUsageRateRequest) {
	request = &DescribeResourceUsageRateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceUsageRate")
	return
}

func NewDescribeResourceUsageRateResponse() (response *DescribeResourceUsageRateResponse) {
	response = &DescribeResourceUsageRateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看资源使用情况
func (c *Client) DescribeResourceUsageRate(request *DescribeResourceUsageRateRequest) (response *DescribeResourceUsageRateResponse, err error) {
	if request == nil {
		request = NewDescribeResourceUsageRateRequest()
	}
	response = NewDescribeResourceUsageRateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJavaInstanceRequest() (request *DescribeJavaInstanceRequest) {
	request = &DescribeJavaInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeJavaInstance")
	return
}

func NewDescribeJavaInstanceResponse() (response *DescribeJavaInstanceResponse) {
	response = &DescribeJavaInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询java实例当前各种指标,以及基本信息
func (c *Client) DescribeJavaInstance(request *DescribeJavaInstanceRequest) (response *DescribeJavaInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeJavaInstanceRequest()
	}
	response = NewDescribeJavaInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceConfigRequest() (request *DescribeResourceConfigRequest) {
	request = &DescribeResourceConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceConfig")
	return
}

func NewDescribeResourceConfigResponse() (response *DescribeResourceConfigResponse) {
	response = &DescribeResourceConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 无
func (c *Client) DescribeResourceConfig(request *DescribeResourceConfigRequest) (response *DescribeResourceConfigResponse, err error) {
	if request == nil {
		request = NewDescribeResourceConfigRequest()
	}
	response = NewDescribeResourceConfigResponse()
	err = c.Send(request, response)
	return
}

func NewContinueResourceBatchOperationRequest() (request *ContinueResourceBatchOperationRequest) {
	request = &ContinueResourceBatchOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ContinueResourceBatchOperation")
	return
}

func NewContinueResourceBatchOperationResponse() (response *ContinueResourceBatchOperationResponse) {
	response = &ContinueResourceBatchOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 继续执行资源批次操作接口
func (c *Client) ContinueResourceBatchOperation(request *ContinueResourceBatchOperationRequest) (response *ContinueResourceBatchOperationResponse, err error) {
	if request == nil {
		request = NewContinueResourceBatchOperationRequest()
	}
	response = NewContinueResourceBatchOperationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAbnormalMetricsConfigRequest() (request *DescribeAbnormalMetricsConfigRequest) {
	request = &DescribeAbnormalMetricsConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeAbnormalMetricsConfig")
	return
}

func NewDescribeAbnormalMetricsConfigResponse() (response *DescribeAbnormalMetricsConfigResponse) {
	response = &DescribeAbnormalMetricsConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询异常指标配置
func (c *Client) DescribeAbnormalMetricsConfig(request *DescribeAbnormalMetricsConfigRequest) (response *DescribeAbnormalMetricsConfigResponse, err error) {
	if request == nil {
		request = NewDescribeAbnormalMetricsConfigRequest()
	}
	response = NewDescribeAbnormalMetricsConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFlowBatchRecordsRequest() (request *DescribeFlowBatchRecordsRequest) {
	request = &DescribeFlowBatchRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFlowBatchRecords")
	return
}

func NewDescribeFlowBatchRecordsResponse() (response *DescribeFlowBatchRecordsResponse) {
	response = &DescribeFlowBatchRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 翻页查询工作流批次列表
func (c *Client) DescribeFlowBatchRecords(request *DescribeFlowBatchRecordsRequest) (response *DescribeFlowBatchRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeFlowBatchRecordsRequest()
	}
	response = NewDescribeFlowBatchRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateProductNewsRequest() (request *CreateProductNewsRequest) {
	request = &CreateProductNewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateProductNews")
	return
}

func NewCreateProductNewsResponse() (response *CreateProductNewsResponse) {
	response = &CreateProductNewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建产品动态
func (c *Client) CreateProductNews(request *CreateProductNewsRequest) (response *CreateProductNewsResponse, err error) {
	if request == nil {
		request = NewCreateProductNewsRequest()
	}
	response = NewCreateProductNewsResponse()
	err = c.Send(request, response)
	return
}

func NewShirkInstanceRequest() (request *ShirkInstanceRequest) {
	request = &ShirkInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ShirkInstance")
	return
}

func NewShirkInstanceResponse() (response *ShirkInstanceResponse) {
	response = &ShirkInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下线部署组实例
func (c *Client) ShirkInstance(request *ShirkInstanceRequest) (response *ShirkInstanceResponse, err error) {
	if request == nil {
		request = NewShirkInstanceRequest()
	}
	response = NewShirkInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeployGroupRequest() (request *DescribeDeployGroupRequest) {
	request = &DescribeDeployGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeployGroup")
	return
}

func NewDescribeDeployGroupResponse() (response *DescribeDeployGroupResponse) {
	response = &DescribeDeployGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部署组详情
func (c *Client) DescribeDeployGroup(request *DescribeDeployGroupRequest) (response *DescribeDeployGroupResponse, err error) {
	if request == nil {
		request = NewDescribeDeployGroupRequest()
	}
	response = NewDescribeDeployGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeJvmLogsRequest() (request *DescribeJvmLogsRequest) {
	request = &DescribeJvmLogsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeJvmLogs")
	return
}

func NewDescribeJvmLogsResponse() (response *DescribeJvmLogsResponse) {
	response = &DescribeJvmLogsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检索Jvm运行时产生的GC日志
func (c *Client) DescribeJvmLogs(request *DescribeJvmLogsRequest) (response *DescribeJvmLogsResponse, err error) {
	if request == nil {
		request = NewDescribeJvmLogsRequest()
	}
	response = NewDescribeJvmLogsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSecretNamesRequest() (request *DescribeSecretNamesRequest) {
	request = &DescribeSecretNamesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSecretNames")
	return
}

func NewDescribeSecretNamesResponse() (response *DescribeSecretNamesResponse) {
	response = &DescribeSecretNamesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取凭证名称列表
func (c *Client) DescribeSecretNames(request *DescribeSecretNamesRequest) (response *DescribeSecretNamesResponse, err error) {
	if request == nil {
		request = NewDescribeSecretNamesRequest()
	}
	response = NewDescribeSecretNamesResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopologyGraphRequest() (request *GetTopologyGraphRequest) {
	request = &GetTopologyGraphRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTopologyGraph")
	return
}

func NewGetTopologyGraphResponse() (response *GetTopologyGraphResponse) {
	response = &GetTopologyGraphResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询依赖拓扑图详情
func (c *Client) GetTopologyGraph(request *GetTopologyGraphRequest) (response *GetTopologyGraphResponse, err error) {
	if request == nil {
		request = NewGetTopologyGraphRequest()
	}
	response = NewGetTopologyGraphResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceMultiCloudTsfConfigRequest() (request *ModifyResourceMultiCloudTsfConfigRequest) {
	request = &ModifyResourceMultiCloudTsfConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyResourceMultiCloudTsfConfig")
	return
}

func NewModifyResourceMultiCloudTsfConfigResponse() (response *ModifyResourceMultiCloudTsfConfigResponse) {
	response = &ModifyResourceMultiCloudTsfConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// TSF多云配置信息更新
func (c *Client) ModifyResourceMultiCloudTsfConfig(request *ModifyResourceMultiCloudTsfConfigRequest) (response *ModifyResourceMultiCloudTsfConfigResponse, err error) {
	if request == nil {
		request = NewModifyResourceMultiCloudTsfConfigRequest()
	}
	response = NewModifyResourceMultiCloudTsfConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventTrendRequest() (request *DescribeEventTrendRequest) {
	request = &DescribeEventTrendRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeEventTrend")
	return
}

func NewDescribeEventTrendResponse() (response *DescribeEventTrendResponse) {
	response = &DescribeEventTrendResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 事件趋势
func (c *Client) DescribeEventTrend(request *DescribeEventTrendRequest) (response *DescribeEventTrendResponse, err error) {
	if request == nil {
		request = NewDescribeEventTrendRequest()
	}
	response = NewDescribeEventTrendResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayJwtPluginRequest() (request *DescribeGatewayJwtPluginRequest) {
	request = &DescribeGatewayJwtPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayJwtPlugin")
	return
}

func NewDescribeGatewayJwtPluginResponse() (response *DescribeGatewayJwtPluginResponse) {
	response = &DescribeGatewayJwtPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询jwt插件详情信息
func (c *Client) DescribeGatewayJwtPlugin(request *DescribeGatewayJwtPluginRequest) (response *DescribeGatewayJwtPluginResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayJwtPluginRequest()
	}
	response = NewDescribeGatewayJwtPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerTasksRequest() (request *DescribeContainerTasksRequest) {
	request = &DescribeContainerTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerTasks")
	return
}

func NewDescribeContainerTasksResponse() (response *DescribeContainerTasksResponse) {
	response = &DescribeContainerTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 变更记录任务列表
func (c *Client) DescribeContainerTasks(request *DescribeContainerTasksRequest) (response *DescribeContainerTasksResponse, err error) {
	if request == nil {
		request = NewDescribeContainerTasksRequest()
	}
	response = NewDescribeContainerTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserPurchaseInfoRequest() (request *DescribeUserPurchaseInfoRequest) {
	request = &DescribeUserPurchaseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUserPurchaseInfo")
	return
}

func NewDescribeUserPurchaseInfoResponse() (response *DescribeUserPurchaseInfoResponse) {
	response = &DescribeUserPurchaseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 描述用户购买信息
func (c *Client) DescribeUserPurchaseInfo(request *DescribeUserPurchaseInfoRequest) (response *DescribeUserPurchaseInfoResponse, err error) {
	if request == nil {
		request = NewDescribeUserPurchaseInfoRequest()
	}
	response = NewDescribeUserPurchaseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCreateViewRequest() (request *CreateViewRequest) {
	request = &CreateViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateView")
	return
}

func NewCreateViewResponse() (response *CreateViewResponse) {
	response = &CreateViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建视图
func (c *Client) CreateView(request *CreateViewRequest) (response *CreateViewResponse, err error) {
	if request == nil {
		request = NewCreateViewRequest()
	}
	response = NewCreateViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeContainerGroupDetailRequest() (request *DescribeContainerGroupDetailRequest) {
	request = &DescribeContainerGroupDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeContainerGroupDetail")
	return
}

func NewDescribeContainerGroupDetailResponse() (response *DescribeContainerGroupDetailResponse) {
	response = &DescribeContainerGroupDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//  容器部署组详情（已废弃，请使用  DescribeContainerGroupDeployInfo）
func (c *Client) DescribeContainerGroupDetail(request *DescribeContainerGroupDetailRequest) (response *DescribeContainerGroupDetailResponse, err error) {
	if request == nil {
		request = NewDescribeContainerGroupDetailRequest()
	}
	response = NewDescribeContainerGroupDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductNewsRequest() (request *DescribeProductNewsRequest) {
	request = &DescribeProductNewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeProductNews")
	return
}

func NewDescribeProductNewsResponse() (response *DescribeProductNewsResponse) {
	response = &DescribeProductNewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品动态列表
func (c *Client) DescribeProductNews(request *DescribeProductNewsRequest) (response *DescribeProductNewsResponse, err error) {
	if request == nil {
		request = NewDescribeProductNewsRequest()
	}
	response = NewDescribeProductNewsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInvocationMetricScatterPlotRequest() (request *DescribeInvocationMetricScatterPlotRequest) {
	request = &DescribeInvocationMetricScatterPlotRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInvocationMetricScatterPlot")
	return
}

func NewDescribeInvocationMetricScatterPlotResponse() (response *DescribeInvocationMetricScatterPlotResponse) {
	response = &DescribeInvocationMetricScatterPlotResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用指标数据散点图
func (c *Client) DescribeInvocationMetricScatterPlot(request *DescribeInvocationMetricScatterPlotRequest) (response *DescribeInvocationMetricScatterPlotResponse, err error) {
	if request == nil {
		request = NewDescribeInvocationMetricScatterPlotRequest()
	}
	response = NewDescribeInvocationMetricScatterPlotResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRatelimitRequest() (request *CreateRatelimitRequest) {
	request = &CreateRatelimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateRatelimit")
	return
}

func NewCreateRatelimitResponse() (response *CreateRatelimitResponse) {
	response = &CreateRatelimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加限流规则
func (c *Client) CreateRatelimit(request *CreateRatelimitRequest) (response *CreateRatelimitResponse, err error) {
	if request == nil {
		request = NewCreateRatelimitRequest()
	}
	response = NewCreateRatelimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicConfigsRequest() (request *DescribePublicConfigsRequest) {
	request = &DescribePublicConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigs")
	return
}

func NewDescribePublicConfigsResponse() (response *DescribePublicConfigsResponse) {
	response = &DescribePublicConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共配置项列表
func (c *Client) DescribePublicConfigs(request *DescribePublicConfigsRequest) (response *DescribePublicConfigsResponse, err error) {
	if request == nil {
		request = NewDescribePublicConfigsRequest()
	}
	response = NewDescribePublicConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchOssRealtimeBusinessLogRequest() (request *SearchOssRealtimeBusinessLogRequest) {
	request = &SearchOssRealtimeBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchOssRealtimeBusinessLog")
	return
}

func NewSearchOssRealtimeBusinessLogResponse() (response *SearchOssRealtimeBusinessLogResponse) {
	response = &SearchOssRealtimeBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// OSS模块群业务日志实时检索
func (c *Client) SearchOssRealtimeBusinessLog(request *SearchOssRealtimeBusinessLogRequest) (response *SearchOssRealtimeBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchOssRealtimeBusinessLogRequest()
	}
	response = NewSearchOssRealtimeBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasePipelineDeployTaskRequest() (request *DescribeReleasePipelineDeployTaskRequest) {
	request = &DescribeReleasePipelineDeployTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasePipelineDeployTask")
	return
}

func NewDescribeReleasePipelineDeployTaskResponse() (response *DescribeReleasePipelineDeployTaskResponse) {
	response = &DescribeReleasePipelineDeployTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部署任务的部署任务id
func (c *Client) DescribeReleasePipelineDeployTask(request *DescribeReleasePipelineDeployTaskRequest) (response *DescribeReleasePipelineDeployTaskResponse, err error) {
	if request == nil {
		request = NewDescribeReleasePipelineDeployTaskRequest()
	}
	response = NewDescribeReleasePipelineDeployTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReleasePipelineTaskRunRecordRequest() (request *DescribeReleasePipelineTaskRunRecordRequest) {
	request = &DescribeReleasePipelineTaskRunRecordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeReleasePipelineTaskRunRecord")
	return
}

func NewDescribeReleasePipelineTaskRunRecordResponse() (response *DescribeReleasePipelineTaskRunRecordResponse) {
	response = &DescribeReleasePipelineTaskRunRecordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询发布单任务执行记录
func (c *Client) DescribeReleasePipelineTaskRunRecord(request *DescribeReleasePipelineTaskRunRecordRequest) (response *DescribeReleasePipelineTaskRunRecordResponse, err error) {
	if request == nil {
		request = NewDescribeReleasePipelineTaskRunRecordRequest()
	}
	response = NewDescribeReleasePipelineTaskRunRecordResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAllGatewayApiAsyncRequest() (request *CreateAllGatewayApiAsyncRequest) {
	request = &CreateAllGatewayApiAsyncRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateAllGatewayApiAsync")
	return
}

func NewCreateAllGatewayApiAsyncResponse() (response *CreateAllGatewayApiAsyncResponse) {
	response = &CreateAllGatewayApiAsyncResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 一键导入API分组
func (c *Client) CreateAllGatewayApiAsync(request *CreateAllGatewayApiAsyncRequest) (response *CreateAllGatewayApiAsyncResponse, err error) {
	if request == nil {
		request = NewCreateAllGatewayApiAsyncRequest()
	}
	response = NewCreateAllGatewayApiAsyncResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRegionRequest() (request *ModifyRegionRequest) {
	request = &ModifyRegionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyRegion")
	return
}

func NewModifyRegionResponse() (response *ModifyRegionResponse) {
	response = &ModifyRegionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改业务地域
func (c *Client) ModifyRegion(request *ModifyRegionRequest) (response *ModifyRegionResponse, err error) {
	if request == nil {
		request = NewModifyRegionRequest()
	}
	response = NewModifyRegionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCircuitBreakerRulesRequest() (request *DescribeCircuitBreakerRulesRequest) {
	request = &DescribeCircuitBreakerRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeCircuitBreakerRules")
	return
}

func NewDescribeCircuitBreakerRulesResponse() (response *DescribeCircuitBreakerRulesResponse) {
	response = &DescribeCircuitBreakerRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询熔断规则列表
func (c *Client) DescribeCircuitBreakerRules(request *DescribeCircuitBreakerRulesRequest) (response *DescribeCircuitBreakerRulesResponse, err error) {
	if request == nil {
		request = NewDescribeCircuitBreakerRulesRequest()
	}
	response = NewDescribeCircuitBreakerRulesResponse()
	err = c.Send(request, response)
	return
}

func NewEnableLaneRuleRequest() (request *EnableLaneRuleRequest) {
	request = &EnableLaneRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "EnableLaneRule")
	return
}

func NewEnableLaneRuleResponse() (response *EnableLaneRuleResponse) {
	response = &EnableLaneRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启用泳道规则
func (c *Client) EnableLaneRule(request *EnableLaneRuleRequest) (response *EnableLaneRuleResponse, err error) {
	if request == nil {
		request = NewEnableLaneRuleRequest()
	}
	response = NewEnableLaneRuleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateContainGroupRequest() (request *CreateContainGroupRequest) {
	request = &CreateContainGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateContainGroup")
	return
}

func NewCreateContainGroupResponse() (response *CreateContainGroupResponse) {
	response = &CreateContainGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建容器部署组
func (c *Client) CreateContainGroup(request *CreateContainGroupRequest) (response *CreateContainGroupResponse, err error) {
	if request == nil {
		request = NewCreateContainGroupRequest()
	}
	response = NewCreateContainGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDownloadTemplateRequest() (request *DownloadTemplateRequest) {
	request = &DownloadTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DownloadTemplate")
	return
}

func NewDownloadTemplateResponse() (response *DownloadTemplateResponse) {
	response = &DownloadTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成模板
func (c *Client) DownloadTemplate(request *DownloadTemplateRequest) (response *DownloadTemplateResponse, err error) {
	if request == nil {
		request = NewDownloadTemplateRequest()
	}
	response = NewDownloadTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewGetEnvRequest() (request *GetEnvRequest) {
	request = &GetEnvRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetEnv")
	return
}

func NewGetEnvResponse() (response *GetEnvResponse) {
	response = &GetEnvResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// GetEnv
func (c *Client) GetEnv(request *GetEnvRequest) (response *GetEnvResponse, err error) {
	if request == nil {
		request = NewGetEnvRequest()
	}
	response = NewGetEnvResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubTasksRequest() (request *DescribeSubTasksRequest) {
	request = &DescribeSubTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSubTasks")
	return
}

func NewDescribeSubTasksResponse() (response *DescribeSubTasksResponse) {
	response = &DescribeSubTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取子任务详情
func (c *Client) DescribeSubTasks(request *DescribeSubTasksRequest) (response *DescribeSubTasksResponse, err error) {
	if request == nil {
		request = NewDescribeSubTasksRequest()
	}
	response = NewDescribeSubTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceMetricsRequest() (request *DescribeInstanceMetricsRequest) {
	request = &DescribeInstanceMetricsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInstanceMetrics")
	return
}

func NewDescribeInstanceMetricsResponse() (response *DescribeInstanceMetricsResponse) {
	response = &DescribeInstanceMetricsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点监控信息列表
func (c *Client) DescribeInstanceMetrics(request *DescribeInstanceMetricsRequest) (response *DescribeInstanceMetricsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceMetricsRequest()
	}
	response = NewDescribeInstanceMetricsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteRuleRequest() (request *DescribeRouteRuleRequest) {
	request = &DescribeRouteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeRouteRule")
	return
}

func NewDescribeRouteRuleResponse() (response *DescribeRouteRuleResponse) {
	response = &DescribeRouteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询路由规则详情
func (c *Client) DescribeRouteRule(request *DescribeRouteRuleRequest) (response *DescribeRouteRuleResponse, err error) {
	if request == nil {
		request = NewDescribeRouteRuleRequest()
	}
	response = NewDescribeRouteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateMultiClusterDeliveryConfigRequest() (request *UpdateMultiClusterDeliveryConfigRequest) {
	request = &UpdateMultiClusterDeliveryConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateMultiClusterDeliveryConfig")
	return
}

func NewUpdateMultiClusterDeliveryConfigResponse() (response *UpdateMultiClusterDeliveryConfigResponse) {
	response = &UpdateMultiClusterDeliveryConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新投递kafka集群配置项
func (c *Client) UpdateMultiClusterDeliveryConfig(request *UpdateMultiClusterDeliveryConfigRequest) (response *UpdateMultiClusterDeliveryConfigResponse, err error) {
	if request == nil {
		request = NewUpdateMultiClusterDeliveryConfigRequest()
	}
	response = NewUpdateMultiClusterDeliveryConfigResponse()
	err = c.Send(request, response)
	return
}

func NewBillingOperationIsolateRequest() (request *BillingOperationIsolateRequest) {
	request = &BillingOperationIsolateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "BillingOperationIsolate")
	return
}

func NewBillingOperationIsolateResponse() (response *BillingOperationIsolateResponse) {
	response = &BillingOperationIsolateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费运营端隔离
func (c *Client) BillingOperationIsolate(request *BillingOperationIsolateRequest) (response *BillingOperationIsolateResponse, err error) {
	if request == nil {
		request = NewBillingOperationIsolateRequest()
	}
	response = NewBillingOperationIsolateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEventDetailRequest() (request *DescribeEventDetailRequest) {
	request = &DescribeEventDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeEventDetail")
	return
}

func NewDescribeEventDetailResponse() (response *DescribeEventDetailResponse) {
	response = &DescribeEventDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 事件详情
func (c *Client) DescribeEventDetail(request *DescribeEventDetailRequest) (response *DescribeEventDetailResponse, err error) {
	if request == nil {
		request = NewDescribeEventDetailRequest()
	}
	response = NewDescribeEventDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMicroserviceRequest() (request *DeleteMicroserviceRequest) {
	request = &DeleteMicroserviceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteMicroservice")
	return
}

func NewDeleteMicroserviceResponse() (response *DeleteMicroserviceResponse) {
	response = &DeleteMicroserviceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除微服务
func (c *Client) DeleteMicroservice(request *DeleteMicroserviceRequest) (response *DeleteMicroserviceResponse, err error) {
	if request == nil {
		request = NewDeleteMicroserviceRequest()
	}
	response = NewDeleteMicroserviceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateZoneRequest() (request *CreateZoneRequest) {
	request = &CreateZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateZone")
	return
}

func NewCreateZoneResponse() (response *CreateZoneResponse) {
	response = &CreateZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建业务可用区
func (c *Client) CreateZone(request *CreateZoneRequest) (response *CreateZoneResponse, err error) {
	if request == nil {
		request = NewCreateZoneRequest()
	}
	response = NewCreateZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayOAuthPluginRequest() (request *DescribeGatewayOAuthPluginRequest) {
	request = &DescribeGatewayOAuthPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayOAuthPlugin")
	return
}

func NewDescribeGatewayOAuthPluginResponse() (response *DescribeGatewayOAuthPluginResponse) {
	response = &DescribeGatewayOAuthPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询OAuth插件详情信息
func (c *Client) DescribeGatewayOAuthPlugin(request *DescribeGatewayOAuthPluginRequest) (response *DescribeGatewayOAuthPluginResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayOAuthPluginRequest()
	}
	response = NewDescribeGatewayOAuthPluginResponse()
	err = c.Send(request, response)
	return
}

func NewModifyProgramRequest() (request *ModifyProgramRequest) {
	request = &ModifyProgramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyProgram")
	return
}

func NewModifyProgramResponse() (response *ModifyProgramResponse) {
	response = &ModifyProgramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新数据集
func (c *Client) ModifyProgram(request *ModifyProgramRequest) (response *ModifyProgramResponse, err error) {
	if request == nil {
		request = NewModifyProgramRequest()
	}
	response = NewModifyProgramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeThreadDumpRequest() (request *DescribeThreadDumpRequest) {
	request = &DescribeThreadDumpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeThreadDump")
	return
}

func NewDescribeThreadDumpResponse() (response *DescribeThreadDumpResponse) {
	response = &DescribeThreadDumpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询线程分析
func (c *Client) DescribeThreadDump(request *DescribeThreadDumpRequest) (response *DescribeThreadDumpResponse, err error) {
	if request == nil {
		request = NewDescribeThreadDumpRequest()
	}
	response = NewDescribeThreadDumpResponse()
	err = c.Send(request, response)
	return
}

func NewGetContainGroupOtherRequest() (request *GetContainGroupOtherRequest) {
	request = &GetContainGroupOtherRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetContainGroupOther")
	return
}

func NewGetContainGroupOtherResponse() (response *GetContainGroupOtherResponse) {
	response = &GetContainGroupOtherResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取部署组其他字段-用于前端并发调用
func (c *Client) GetContainGroupOther(request *GetContainGroupOtherRequest) (response *GetContainGroupOtherResponse, err error) {
	if request == nil {
		request = NewGetContainGroupOtherRequest()
	}
	response = NewGetContainGroupOtherResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupSecretsRequest() (request *DescribeGroupSecretsRequest) {
	request = &DescribeGroupSecretsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGroupSecrets")
	return
}

func NewDescribeGroupSecretsResponse() (response *DescribeGroupSecretsResponse) {
	response = &DescribeGroupSecretsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询秘钥列表
func (c *Client) DescribeGroupSecrets(request *DescribeGroupSecretsRequest) (response *DescribeGroupSecretsResponse, err error) {
	if request == nil {
		request = NewDescribeGroupSecretsRequest()
	}
	response = NewDescribeGroupSecretsResponse()
	err = c.Send(request, response)
	return
}

func NewGetOssTraceInterfacesRequest() (request *GetOssTraceInterfacesRequest) {
	request = &GetOssTraceInterfacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetOssTraceInterfaces")
	return
}

func NewGetOssTraceInterfacesResponse() (response *GetOssTraceInterfacesResponse) {
	response = &GetOssTraceInterfacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用链接口列表
func (c *Client) GetOssTraceInterfaces(request *GetOssTraceInterfacesRequest) (response *GetOssTraceInterfacesResponse, err error) {
	if request == nil {
		request = NewGetOssTraceInterfacesRequest()
	}
	response = NewGetOssTraceInterfacesResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackPublicConfigRequest() (request *RollbackPublicConfigRequest) {
	request = &RollbackPublicConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RollbackPublicConfig")
	return
}

func NewRollbackPublicConfigResponse() (response *RollbackPublicConfigResponse) {
	response = &RollbackPublicConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚公共配置
func (c *Client) RollbackPublicConfig(request *RollbackPublicConfigRequest) (response *RollbackPublicConfigResponse, err error) {
	if request == nil {
		request = NewRollbackPublicConfigRequest()
	}
	response = NewRollbackPublicConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseInfoRequest() (request *DescribeLicenseInfoRequest) {
	request = &DescribeLicenseInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLicenseInfo")
	return
}

func NewDescribeLicenseInfoResponse() (response *DescribeLicenseInfoResponse) {
	response = &DescribeLicenseInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// license相关信息概览接口
func (c *Client) DescribeLicenseInfo(request *DescribeLicenseInfoRequest) (response *DescribeLicenseInfoResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseInfoRequest()
	}
	response = NewDescribeLicenseInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDraftPluginRequest() (request *DraftPluginRequest) {
	request = &DraftPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DraftPlugin")
	return
}

func NewDraftPluginResponse() (response *DraftPluginResponse) {
	response = &DraftPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下线插件
func (c *Client) DraftPlugin(request *DraftPluginRequest) (response *DraftPluginResponse, err error) {
	if request == nil {
		request = NewDraftPluginRequest()
	}
	response = NewDraftPluginResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKubeDeploymentsRequest() (request *DescribeKubeDeploymentsRequest) {
	request = &DescribeKubeDeploymentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeKubeDeployments")
	return
}

func NewDescribeKubeDeploymentsResponse() (response *DescribeKubeDeploymentsResponse) {
	response = &DescribeKubeDeploymentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群deployment列表
func (c *Client) DescribeKubeDeployments(request *DescribeKubeDeploymentsRequest) (response *DescribeKubeDeploymentsResponse, err error) {
	if request == nil {
		request = NewDescribeKubeDeploymentsRequest()
	}
	response = NewDescribeKubeDeploymentsResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindPluginRequest() (request *UnbindPluginRequest) {
	request = &UnbindPluginRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UnbindPlugin")
	return
}

func NewUnbindPluginResponse() (response *UnbindPluginResponse) {
	response = &UnbindPluginResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 插件批量解绑
func (c *Client) UnbindPlugin(request *UnbindPluginRequest) (response *UnbindPluginResponse, err error) {
	if request == nil {
		request = NewUnbindPluginRequest()
	}
	response = NewUnbindPluginResponse()
	err = c.Send(request, response)
	return
}

func NewStopContainerGroupRequest() (request *StopContainerGroupRequest) {
	request = &StopContainerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StopContainerGroup")
	return
}

func NewStopContainerGroupResponse() (response *StopContainerGroupResponse) {
	response = &StopContainerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止容器部署组
func (c *Client) StopContainerGroup(request *StopContainerGroupRequest) (response *StopContainerGroupResponse, err error) {
	if request == nil {
		request = NewStopContainerGroupRequest()
	}
	response = NewStopContainerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewShrinkInstanceRequest() (request *ShrinkInstanceRequest) {
	request = &ShrinkInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ShrinkInstance")
	return
}

func NewShrinkInstanceResponse() (response *ShrinkInstanceResponse) {
	response = &ShrinkInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 废弃，不可用
func (c *Client) ShrinkInstance(request *ShrinkInstanceRequest) (response *ShrinkInstanceResponse, err error) {
	if request == nil {
		request = NewShrinkInstanceRequest()
	}
	response = NewShrinkInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewExpandNamespaceRequest() (request *ExpandNamespaceRequest) {
	request = &ExpandNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ExpandNamespace")
	return
}

func NewExpandNamespaceResponse() (response *ExpandNamespaceResponse) {
	response = &ExpandNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加机器至命名空间中
func (c *Client) ExpandNamespace(request *ExpandNamespaceRequest) (response *ExpandNamespaceResponse, err error) {
	if request == nil {
		request = NewExpandNamespaceRequest()
	}
	response = NewExpandNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewOperateApplicationTcrBindingRequest() (request *OperateApplicationTcrBindingRequest) {
	request = &OperateApplicationTcrBindingRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "OperateApplicationTcrBinding")
	return
}

func NewOperateApplicationTcrBindingResponse() (response *OperateApplicationTcrBindingResponse) {
	response = &OperateApplicationTcrBindingResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定解绑tcr仓库
func (c *Client) OperateApplicationTcrBinding(request *OperateApplicationTcrBindingRequest) (response *OperateApplicationTcrBindingResponse, err error) {
	if request == nil {
		request = NewOperateApplicationTcrBindingRequest()
	}
	response = NewOperateApplicationTcrBindingResponse()
	err = c.Send(request, response)
	return
}

func NewBillingOperationModifyRequest() (request *BillingOperationModifyRequest) {
	request = &BillingOperationModifyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "BillingOperationModify")
	return
}

func NewBillingOperationModifyResponse() (response *BillingOperationModifyResponse) {
	response = &BillingOperationModifyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 计费运营端变配
func (c *Client) BillingOperationModify(request *BillingOperationModifyRequest) (response *BillingOperationModifyResponse, err error) {
	if request == nil {
		request = NewBillingOperationModifyRequest()
	}
	response = NewBillingOperationModifyResponse()
	err = c.Send(request, response)
	return
}

func NewDisassociateKafkaConfigRequest() (request *DisassociateKafkaConfigRequest) {
	request = &DisassociateKafkaConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DisassociateKafkaConfig")
	return
}

func NewDisassociateKafkaConfigResponse() (response *DisassociateKafkaConfigResponse) {
	response = &DisassociateKafkaConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消关联投递信息和部署组
func (c *Client) DisassociateKafkaConfig(request *DisassociateKafkaConfigRequest) (response *DisassociateKafkaConfigResponse, err error) {
	if request == nil {
		request = NewDisassociateKafkaConfigRequest()
	}
	response = NewDisassociateKafkaConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeliverFilebeatConfigRequest() (request *DeliverFilebeatConfigRequest) {
	request = &DeliverFilebeatConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeliverFilebeatConfig")
	return
}

func NewDeliverFilebeatConfigResponse() (response *DeliverFilebeatConfigResponse) {
	response = &DeliverFilebeatConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下发filebeat config 到master api
func (c *Client) DeliverFilebeatConfig(request *DeliverFilebeatConfigRequest) (response *DeliverFilebeatConfigResponse, err error) {
	if request == nil {
		request = NewDeliverFilebeatConfigRequest()
	}
	response = NewDeliverFilebeatConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRoleRequest() (request *DeleteRoleRequest) {
	request = &DeleteRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteRole")
	return
}

func NewDeleteRoleResponse() (response *DeleteRoleResponse) {
	response = &DeleteRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除角色
func (c *Client) DeleteRole(request *DeleteRoleRequest) (response *DeleteRoleResponse, err error) {
	if request == nil {
		request = NewDeleteRoleRequest()
	}
	response = NewDeleteRoleResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNamespaceCodeRequest() (request *ModifyNamespaceCodeRequest) {
	request = &ModifyNamespaceCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyNamespaceCode")
	return
}

func NewModifyNamespaceCodeResponse() (response *ModifyNamespaceCodeResponse) {
	response = &ModifyNamespaceCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新命名空间编码
func (c *Client) ModifyNamespaceCode(request *ModifyNamespaceCodeRequest) (response *ModifyNamespaceCodeResponse, err error) {
	if request == nil {
		request = NewModifyNamespaceCodeRequest()
	}
	response = NewModifyNamespaceCodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDeployGroupsRequest() (request *DescribeDeployGroupsRequest) {
	request = &DescribeDeployGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeDeployGroups")
	return
}

func NewDescribeDeployGroupsResponse() (response *DescribeDeployGroupsResponse) {
	response = &DescribeDeployGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询部署组列表
func (c *Client) DescribeDeployGroups(request *DescribeDeployGroupsRequest) (response *DescribeDeployGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeDeployGroupsRequest()
	}
	response = NewDescribeDeployGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePublicConfigReleasesRequest() (request *DescribePublicConfigReleasesRequest) {
	request = &DescribePublicConfigReleasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribePublicConfigReleases")
	return
}

func NewDescribePublicConfigReleasesResponse() (response *DescribePublicConfigReleasesResponse) {
	response = &DescribePublicConfigReleasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询公共配置发布信息
func (c *Client) DescribePublicConfigReleases(request *DescribePublicConfigReleasesRequest) (response *DescribePublicConfigReleasesResponse, err error) {
	if request == nil {
		request = NewDescribePublicConfigReleasesRequest()
	}
	response = NewDescribePublicConfigReleasesResponse()
	err = c.Send(request, response)
	return
}

func NewInterruptResourceBatchOperationRequest() (request *InterruptResourceBatchOperationRequest) {
	request = &InterruptResourceBatchOperationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "InterruptResourceBatchOperation")
	return
}

func NewInterruptResourceBatchOperationResponse() (response *InterruptResourceBatchOperationResponse) {
	response = &InterruptResourceBatchOperationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 中断资源批次操作接口
func (c *Client) InterruptResourceBatchOperation(request *InterruptResourceBatchOperationRequest) (response *InterruptResourceBatchOperationResponse, err error) {
	if request == nil {
		request = NewInterruptResourceBatchOperationRequest()
	}
	response = NewInterruptResourceBatchOperationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsableApisRequest() (request *DescribeUsableApisRequest) {
	request = &DescribeUsableApisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUsableApis")
	return
}

func NewDescribeUsableApisResponse() (response *DescribeUsableApisResponse) {
	response = &DescribeUsableApisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用于被导入的API
func (c *Client) DescribeUsableApis(request *DescribeUsableApisRequest) (response *DescribeUsableApisResponse, err error) {
	if request == nil {
		request = NewDescribeUsableApisRequest()
	}
	response = NewDescribeUsableApisResponse()
	err = c.Send(request, response)
	return
}

func NewRedoTaskBatchRequest() (request *RedoTaskBatchRequest) {
	request = &RedoTaskBatchRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RedoTaskBatch")
	return
}

func NewRedoTaskBatchResponse() (response *RedoTaskBatchResponse) {
	response = &RedoTaskBatchResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重新执行任务批次
func (c *Client) RedoTaskBatch(request *RedoTaskBatchRequest) (response *RedoTaskBatchResponse, err error) {
	if request == nil {
		request = NewRedoTaskBatchRequest()
	}
	response = NewRedoTaskBatchResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRatelimitRequest() (request *DeleteRatelimitRequest) {
	request = &DeleteRatelimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteRatelimit")
	return
}

func NewDeleteRatelimitResponse() (response *DeleteRatelimitResponse) {
	response = &DeleteRatelimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除限流规则
func (c *Client) DeleteRatelimit(request *DeleteRatelimitRequest) (response *DeleteRatelimitResponse, err error) {
	if request == nil {
		request = NewDeleteRatelimitRequest()
	}
	response = NewDeleteRatelimitResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlarmReceiverRequest() (request *CreateAlarmReceiverRequest) {
	request = &CreateAlarmReceiverRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateAlarmReceiver")
	return
}

func NewCreateAlarmReceiverResponse() (response *CreateAlarmReceiverResponse) {
	response = &CreateAlarmReceiverResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增绑定人员
func (c *Client) CreateAlarmReceiver(request *CreateAlarmReceiverRequest) (response *CreateAlarmReceiverResponse, err error) {
	if request == nil {
		request = NewCreateAlarmReceiverRequest()
	}
	response = NewCreateAlarmReceiverResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnitNamespacesRequest() (request *DescribeUnitNamespacesRequest) {
	request = &DescribeUnitNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitNamespaces")
	return
}

func NewDescribeUnitNamespacesResponse() (response *DescribeUnitNamespacesResponse) {
	response = &DescribeUnitNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单元化命名空间列表
func (c *Client) DescribeUnitNamespaces(request *DescribeUnitNamespacesRequest) (response *DescribeUnitNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeUnitNamespacesRequest()
	}
	response = NewDescribeUnitNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewStopInstanceRequest() (request *StopInstanceRequest) {
	request = &StopInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "StopInstance")
	return
}

func NewStopInstanceResponse() (response *StopInstanceResponse) {
	response = &StopInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 停止机器
func (c *Client) StopInstance(request *StopInstanceRequest) (response *StopInstanceResponse, err error) {
	if request == nil {
		request = NewStopInstanceRequest()
	}
	response = NewStopInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewRollbackConfigRequest() (request *RollbackConfigRequest) {
	request = &RollbackConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "RollbackConfig")
	return
}

func NewRollbackConfigResponse() (response *RollbackConfigResponse) {
	response = &RollbackConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回滚配置
func (c *Client) RollbackConfig(request *RollbackConfigRequest) (response *RollbackConfigResponse, err error) {
	if request == nil {
		request = NewRollbackConfigRequest()
	}
	response = NewRollbackConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespacesRequest() (request *DescribeNamespacesRequest) {
	request = &DescribeNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeNamespaces")
	return
}

func NewDescribeNamespacesResponse() (response *DescribeNamespacesResponse) {
	response = &DescribeNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

//  获取命名空间列表
func (c *Client) DescribeNamespaces(request *DescribeNamespacesRequest) (response *DescribeNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeNamespacesRequest()
	}
	response = NewDescribeNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProductHelpRequest() (request *DescribeProductHelpRequest) {
	request = &DescribeProductHelpRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeProductHelp")
	return
}

func NewDescribeProductHelpResponse() (response *DescribeProductHelpResponse) {
	response = &DescribeProductHelpResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询产品帮助列表
func (c *Client) DescribeProductHelp(request *DescribeProductHelpRequest) (response *DescribeProductHelpResponse, err error) {
	if request == nil {
		request = NewDescribeProductHelpRequest()
	}
	response = NewDescribeProductHelpResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApiGroupRequest() (request *CreateApiGroupRequest) {
	request = &CreateApiGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "CreateApiGroup")
	return
}

func NewCreateApiGroupResponse() (response *CreateApiGroupResponse) {
	response = &CreateApiGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建API分组
func (c *Client) CreateApiGroup(request *CreateApiGroupRequest) (response *CreateApiGroupResponse, err error) {
	if request == nil {
		request = NewCreateApiGroupRequest()
	}
	response = NewCreateApiGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGatewayPluginsRequest() (request *DescribeGatewayPluginsRequest) {
	request = &DescribeGatewayPluginsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeGatewayPlugins")
	return
}

func NewDescribeGatewayPluginsResponse() (response *DescribeGatewayPluginsResponse) {
	response = &DescribeGatewayPluginsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 分页查询网关插件列表
func (c *Client) DescribeGatewayPlugins(request *DescribeGatewayPluginsRequest) (response *DescribeGatewayPluginsResponse, err error) {
	if request == nil {
		request = NewDescribeGatewayPluginsRequest()
	}
	response = NewDescribeGatewayPluginsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMicroservicesByGroupIdsRequest() (request *DescribeMicroservicesByGroupIdsRequest) {
	request = &DescribeMicroservicesByGroupIdsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeMicroservicesByGroupIds")
	return
}

func NewDescribeMicroservicesByGroupIdsResponse() (response *DescribeMicroservicesByGroupIdsResponse) {
	response = &DescribeMicroservicesByGroupIdsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过部署组ID获取微服务
func (c *Client) DescribeMicroservicesByGroupIds(request *DescribeMicroservicesByGroupIdsRequest) (response *DescribeMicroservicesByGroupIdsResponse, err error) {
	if request == nil {
		request = NewDescribeMicroservicesByGroupIdsRequest()
	}
	response = NewDescribeMicroservicesByGroupIdsResponse()
	err = c.Send(request, response)
	return
}

func NewDeletePkgRequest() (request *DeletePkgRequest) {
	request = &DeletePkgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeletePkg")
	return
}

func NewDeletePkgResponse() (response *DeletePkgResponse) {
	response = &DeletePkgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除包
func (c *Client) DeletePkg(request *DeletePkgRequest) (response *DeletePkgResponse, err error) {
	if request == nil {
		request = NewDeletePkgRequest()
	}
	response = NewDeletePkgResponse()
	err = c.Send(request, response)
	return
}

func NewSearchOssTraceRequest() (request *SearchOssTraceRequest) {
	request = &SearchOssTraceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchOssTrace")
	return
}

func NewSearchOssTraceResponse() (response *SearchOssTraceResponse) {
	response = &SearchOssTraceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用链
func (c *Client) SearchOssTrace(request *SearchOssTraceRequest) (response *SearchOssTraceResponse, err error) {
	if request == nil {
		request = NewSearchOssTraceRequest()
	}
	response = NewSearchOssTraceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceMonitorDetailsRequest() (request *DescribeInstanceMonitorDetailsRequest) {
	request = &DescribeInstanceMonitorDetailsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeInstanceMonitorDetails")
	return
}

func NewDescribeInstanceMonitorDetailsResponse() (response *DescribeInstanceMonitorDetailsResponse) {
	response = &DescribeInstanceMonitorDetailsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询节点监控具体信息
func (c *Client) DescribeInstanceMonitorDetails(request *DescribeInstanceMonitorDetailsRequest) (response *DescribeInstanceMonitorDetailsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceMonitorDetailsRequest()
	}
	response = NewDescribeInstanceMonitorDetailsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUsableGatewayGroupsRequest() (request *DescribeUsableGatewayGroupsRequest) {
	request = &DescribeUsableGatewayGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUsableGatewayGroups")
	return
}

func NewDescribeUsableGatewayGroupsResponse() (response *DescribeUsableGatewayGroupsResponse) {
	response = &DescribeUsableGatewayGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某个API分组下可用于被绑定的网关部署组
func (c *Client) DescribeUsableGatewayGroups(request *DescribeUsableGatewayGroupsRequest) (response *DescribeUsableGatewayGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeUsableGatewayGroupsRequest()
	}
	response = NewDescribeUsableGatewayGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTsfZoneRequest() (request *DeleteTsfZoneRequest) {
	request = &DeleteTsfZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteTsfZone")
	return
}

func NewDeleteTsfZoneResponse() (response *DeleteTsfZoneResponse) {
	response = &DeleteTsfZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除基础可用区
func (c *Client) DeleteTsfZone(request *DeleteTsfZoneRequest) (response *DeleteTsfZoneResponse, err error) {
	if request == nil {
		request = NewDeleteTsfZoneRequest()
	}
	response = NewDeleteTsfZoneResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateRepositoryRequest() (request *UpdateRepositoryRequest) {
	request = &UpdateRepositoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateRepository")
	return
}

func NewUpdateRepositoryResponse() (response *UpdateRepositoryResponse) {
	response = &UpdateRepositoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新仓库信息
func (c *Client) UpdateRepository(request *UpdateRepositoryRequest) (response *UpdateRepositoryResponse, err error) {
	if request == nil {
		request = NewUpdateRepositoryRequest()
	}
	response = NewUpdateRepositoryResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDeliveryConfigRequest() (request *DeleteDeliveryConfigRequest) {
	request = &DeleteDeliveryConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteDeliveryConfig")
	return
}

func NewDeleteDeliveryConfigResponse() (response *DeleteDeliveryConfigResponse) {
	response = &DeleteDeliveryConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除投递配置
func (c *Client) DeleteDeliveryConfig(request *DeleteDeliveryConfigRequest) (response *DeleteDeliveryConfigResponse, err error) {
	if request == nil {
		request = NewDeleteDeliveryConfigRequest()
	}
	response = NewDeleteDeliveryConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetTraceSpansRequest() (request *GetTraceSpansRequest) {
	request = &GetTraceSpansRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "GetTraceSpans")
	return
}

func NewGetTraceSpansResponse() (response *GetTraceSpansResponse) {
	response = &GetTraceSpansResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询调用链Span
func (c *Client) GetTraceSpans(request *GetTraceSpansRequest) (response *GetTraceSpansResponse, err error) {
	if request == nil {
		request = NewGetTraceSpansRequest()
	}
	response = NewGetTraceSpansResponse()
	err = c.Send(request, response)
	return
}

func NewSetNamespaceCodeRequest() (request *SetNamespaceCodeRequest) {
	request = &SetNamespaceCodeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SetNamespaceCode")
	return
}

func NewSetNamespaceCodeResponse() (response *SetNamespaceCodeResponse) {
	response = &SetNamespaceCodeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置命名空间编码
func (c *Client) SetNamespaceCode(request *SetNamespaceCodeRequest) (response *SetNamespaceCodeResponse, err error) {
	if request == nil {
		request = NewSetNamespaceCodeRequest()
	}
	response = NewSetNamespaceCodeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSimpleInstancesRequest() (request *DescribeSimpleInstancesRequest) {
	request = &DescribeSimpleInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeSimpleInstances")
	return
}

func NewDescribeSimpleInstancesResponse() (response *DescribeSimpleInstancesResponse) {
	response = &DescribeSimpleInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询简单机器实例列表
func (c *Client) DescribeSimpleInstances(request *DescribeSimpleInstancesRequest) (response *DescribeSimpleInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeSimpleInstancesRequest()
	}
	response = NewDescribeSimpleInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteContainerGroupRequest() (request *DeleteContainerGroupRequest) {
	request = &DeleteContainerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteContainerGroup")
	return
}

func NewDeleteContainerGroupResponse() (response *DeleteContainerGroupResponse) {
	response = &DeleteContainerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除容器部署组
func (c *Client) DeleteContainerGroup(request *DeleteContainerGroupRequest) (response *DeleteContainerGroupResponse, err error) {
	if request == nil {
		request = NewDeleteContainerGroupRequest()
	}
	response = NewDeleteContainerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewListTsfModuleInstancePortsRequest() (request *ListTsfModuleInstancePortsRequest) {
	request = &ListTsfModuleInstancePortsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ListTsfModuleInstancePorts")
	return
}

func NewListTsfModuleInstancePortsResponse() (response *ListTsfModuleInstancePortsResponse) {
	response = &ListTsfModuleInstancePortsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询模块实例节点端口信息
func (c *Client) ListTsfModuleInstancePorts(request *ListTsfModuleInstancePortsRequest) (response *ListTsfModuleInstancePortsResponse, err error) {
	if request == nil {
		request = NewListTsfModuleInstancePortsRequest()
	}
	response = NewListTsfModuleInstancePortsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLicenseApplicationsRequest() (request *DescribeLicenseApplicationsRequest) {
	request = &DescribeLicenseApplicationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeLicenseApplications")
	return
}

func NewDescribeLicenseApplicationsResponse() (response *DescribeLicenseApplicationsResponse) {
	response = &DescribeLicenseApplicationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取许可申请列表
func (c *Client) DescribeLicenseApplications(request *DescribeLicenseApplicationsRequest) (response *DescribeLicenseApplicationsResponse, err error) {
	if request == nil {
		request = NewDescribeLicenseApplicationsRequest()
	}
	response = NewDescribeLicenseApplicationsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateApplicationAgentRequest() (request *UpdateApplicationAgentRequest) {
	request = &UpdateApplicationAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "UpdateApplicationAgent")
	return
}

func NewUpdateApplicationAgentResponse() (response *UpdateApplicationAgentResponse) {
	response = &UpdateApplicationAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Agent
func (c *Client) UpdateApplicationAgent(request *UpdateApplicationAgentRequest) (response *UpdateApplicationAgentResponse, err error) {
	if request == nil {
		request = NewUpdateApplicationAgentRequest()
	}
	response = NewUpdateApplicationAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterInstanceCountRequest() (request *DescribeClusterInstanceCountRequest) {
	request = &DescribeClusterInstanceCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeClusterInstanceCount")
	return
}

func NewDescribeClusterInstanceCountResponse() (response *DescribeClusterInstanceCountResponse) {
	response = &DescribeClusterInstanceCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询集群机器统计信息
func (c *Client) DescribeClusterInstanceCount(request *DescribeClusterInstanceCountRequest) (response *DescribeClusterInstanceCountResponse, err error) {
	if request == nil {
		request = NewDescribeClusterInstanceCountRequest()
	}
	response = NewDescribeClusterInstanceCountResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUnitRuleRequest() (request *DescribeUnitRuleRequest) {
	request = &DescribeUnitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeUnitRule")
	return
}

func NewDescribeUnitRuleResponse() (response *DescribeUnitRuleResponse) {
	response = &DescribeUnitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询单元化规则详情
func (c *Client) DescribeUnitRule(request *DescribeUnitRuleRequest) (response *DescribeUnitRuleResponse, err error) {
	if request == nil {
		request = NewDescribeUnitRuleRequest()
	}
	response = NewDescribeUnitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceTaskStatusRequest() (request *DescribeResourceTaskStatusRequest) {
	request = &DescribeResourceTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeResourceTaskStatus")
	return
}

func NewDescribeResourceTaskStatusResponse() (response *DescribeResourceTaskStatusResponse) {
	response = &DescribeResourceTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 资源任务的执行状态描述接口
func (c *Client) DescribeResourceTaskStatus(request *DescribeResourceTaskStatusRequest) (response *DescribeResourceTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeResourceTaskStatusRequest()
	}
	response = NewDescribeResourceTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewFindContainerGroupRequest() (request *FindContainerGroupRequest) {
	request = &FindContainerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "FindContainerGroup")
	return
}

func NewFindContainerGroupResponse() (response *FindContainerGroupResponse) {
	response = &FindContainerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查找部署组详情
func (c *Client) FindContainerGroup(request *FindContainerGroupRequest) (response *FindContainerGroupResponse, err error) {
	if request == nil {
		request = NewFindContainerGroupRequest()
	}
	response = NewFindContainerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewSearchContainerStdoutLogRequest() (request *SearchContainerStdoutLogRequest) {
	request = &SearchContainerStdoutLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchContainerStdoutLog")
	return
}

func NewSearchContainerStdoutLogResponse() (response *SearchContainerStdoutLogResponse) {
	response = &SearchContainerStdoutLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询容器标准输出日志
func (c *Client) SearchContainerStdoutLog(request *SearchContainerStdoutLogRequest) (response *SearchContainerStdoutLogResponse, err error) {
	if request == nil {
		request = NewSearchContainerStdoutLogRequest()
	}
	response = NewSearchContainerStdoutLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFileConfigReleasesRequest() (request *DescribeFileConfigReleasesRequest) {
	request = &DescribeFileConfigReleasesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeFileConfigReleases")
	return
}

func NewDescribeFileConfigReleasesResponse() (response *DescribeFileConfigReleasesResponse) {
	response = &DescribeFileConfigReleasesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询文件配置项发布信息
func (c *Client) DescribeFileConfigReleases(request *DescribeFileConfigReleasesRequest) (response *DescribeFileConfigReleasesResponse, err error) {
	if request == nil {
		request = NewDescribeFileConfigReleasesRequest()
	}
	response = NewDescribeFileConfigReleasesResponse()
	err = c.Send(request, response)
	return
}

func NewSearchOssBusinessLogRequest() (request *SearchOssBusinessLogRequest) {
	request = &SearchOssBusinessLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SearchOssBusinessLog")
	return
}

func NewSearchOssBusinessLogResponse() (response *SearchOssBusinessLogResponse) {
	response = &SearchOssBusinessLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务日志搜索
func (c *Client) SearchOssBusinessLog(request *SearchOssBusinessLogRequest) (response *SearchOssBusinessLogResponse, err error) {
	if request == nil {
		request = NewSearchOssBusinessLogRequest()
	}
	response = NewSearchOssBusinessLogResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRoleRequest() (request *ModifyRoleRequest) {
	request = &ModifyRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ModifyRole")
	return
}

func NewModifyRoleResponse() (response *ModifyRoleResponse) {
	response = &ModifyRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新角色
func (c *Client) ModifyRole(request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
	if request == nil {
		request = NewModifyRoleRequest()
	}
	response = NewModifyRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskExecuteHistoryRecordsRequest() (request *DescribeTaskExecuteHistoryRecordsRequest) {
	request = &DescribeTaskExecuteHistoryRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskExecuteHistoryRecords")
	return
}

func NewDescribeTaskExecuteHistoryRecordsResponse() (response *DescribeTaskExecuteHistoryRecordsResponse) {
	response = &DescribeTaskExecuteHistoryRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询某一个执行记录的执行历史流水
func (c *Client) DescribeTaskExecuteHistoryRecords(request *DescribeTaskExecuteHistoryRecordsRequest) (response *DescribeTaskExecuteHistoryRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeTaskExecuteHistoryRecordsRequest()
	}
	response = NewDescribeTaskExecuteHistoryRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUnitRuleRequest() (request *DeleteUnitRuleRequest) {
	request = &DeleteUnitRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeleteUnitRule")
	return
}

func NewDeleteUnitRuleResponse() (response *DeleteUnitRuleResponse) {
	response = &DeleteUnitRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除单元化规则
func (c *Client) DeleteUnitRule(request *DeleteUnitRuleRequest) (response *DeleteUnitRuleResponse, err error) {
	if request == nil {
		request = NewDeleteUnitRuleRequest()
	}
	response = NewDeleteUnitRuleResponse()
	err = c.Send(request, response)
	return
}

func NewSetCkafkaToContainerClusterRequest() (request *SetCkafkaToContainerClusterRequest) {
	request = &SetCkafkaToContainerClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "SetCkafkaToContainerCluster")
	return
}

func NewSetCkafkaToContainerClusterResponse() (response *SetCkafkaToContainerClusterResponse) {
	response = &SetCkafkaToContainerClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置ckafka到容器集群
func (c *Client) SetCkafkaToContainerCluster(request *SetCkafkaToContainerClusterRequest) (response *SetCkafkaToContainerClusterResponse, err error) {
	if request == nil {
		request = NewSetCkafkaToContainerClusterRequest()
	}
	response = NewSetCkafkaToContainerClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskExecuteLogRequest() (request *DescribeTaskExecuteLogRequest) {
	request = &DescribeTaskExecuteLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskExecuteLog")
	return
}

func NewDescribeTaskExecuteLogResponse() (response *DescribeTaskExecuteLogResponse) {
	response = &DescribeTaskExecuteLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务执行日志
func (c *Client) DescribeTaskExecuteLog(request *DescribeTaskExecuteLogRequest) (response *DescribeTaskExecuteLogResponse, err error) {
	if request == nil {
		request = NewDescribeTaskExecuteLogRequest()
	}
	response = NewDescribeTaskExecuteLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskExecuteShardArgumentRequest() (request *DescribeTaskExecuteShardArgumentRequest) {
	request = &DescribeTaskExecuteShardArgumentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DescribeTaskExecuteShardArgument")
	return
}

func NewDescribeTaskExecuteShardArgumentResponse() (response *DescribeTaskExecuteShardArgumentResponse) {
	response = &DescribeTaskExecuteShardArgumentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务执行时的分片参数信息
func (c *Client) DescribeTaskExecuteShardArgument(request *DescribeTaskExecuteShardArgumentRequest) (response *DescribeTaskExecuteShardArgumentResponse, err error) {
	if request == nil {
		request = NewDescribeTaskExecuteShardArgumentRequest()
	}
	response = NewDescribeTaskExecuteShardArgumentResponse()
	err = c.Send(request, response)
	return
}

func NewDeployContainerGroupByPlainYamlRequest() (request *DeployContainerGroupByPlainYamlRequest) {
	request = &DeployContainerGroupByPlainYamlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "DeployContainerGroupByPlainYaml")
	return
}

func NewDeployContainerGroupByPlainYamlResponse() (response *DeployContainerGroupByPlainYamlResponse) {
	response = &DeployContainerGroupByPlainYamlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 这里的简单 Yaml 指 k8s 的 YAML 剔除了 TSF 的信息后的 yaml。和 DescribeContainerGroupPlainYaml 返回的内容一致
func (c *Client) DeployContainerGroupByPlainYaml(request *DeployContainerGroupByPlainYamlRequest) (response *DeployContainerGroupByPlainYamlResponse, err error) {
	if request == nil {
		request = NewDeployContainerGroupByPlainYamlRequest()
	}
	response = NewDeployContainerGroupByPlainYamlResponse()
	err = c.Send(request, response)
	return
}

func NewImageUserIsExistsRequest() (request *ImageUserIsExistsRequest) {
	request = &ImageUserIsExistsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tsf", APIVersion, "ImageUserIsExists")
	return
}

func NewImageUserIsExistsResponse() (response *ImageUserIsExistsResponse) {
	response = &ImageUserIsExistsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 首次点击”镜像标签页”时触发
func (c *Client) ImageUserIsExists(request *ImageUserIsExistsRequest) (response *ImageUserIsExistsResponse, err error) {
	if request == nil {
		request = NewImageUserIsExistsRequest()
	}
	response = NewImageUserIsExistsResponse()
	err = c.Send(request, response)
	return
}

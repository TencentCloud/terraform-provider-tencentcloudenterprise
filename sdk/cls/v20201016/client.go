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

package v20201016

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-10-16"

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

func NewDescribeMetricCorrectDimensionRequest() (request *DescribeMetricCorrectDimensionRequest) {
	request = &DescribeMetricCorrectDimensionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMetricCorrectDimension")
	return
}

func NewDescribeMetricCorrectDimensionResponse() (response *DescribeMetricCorrectDimensionResponse) {
	response = &DescribeMetricCorrectDimensionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指标校准维度信息
func (c *Client) DescribeMetricCorrectDimension(request *DescribeMetricCorrectDimensionRequest) (response *DescribeMetricCorrectDimensionResponse, err error) {
	if request == nil {
		request = NewDescribeMetricCorrectDimensionRequest()
	}
	response = NewDescribeMetricCorrectDimensionResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmNoticeRequest() (request *ModifyAlarmNoticeRequest) {
	request = &ModifyAlarmNoticeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarmNotice")
	return
}

func NewModifyAlarmNoticeResponse() (response *ModifyAlarmNoticeResponse) {
	response = &ModifyAlarmNoticeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于修改告警通知模板。
func (c *Client) ModifyAlarmNotice(request *ModifyAlarmNoticeRequest) (response *ModifyAlarmNoticeResponse, err error) {
	if request == nil {
		request = NewModifyAlarmNoticeRequest()
	}
	response = NewModifyAlarmNoticeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExceptionResourcesRequest() (request *DescribeExceptionResourcesRequest) {
	request = &DescribeExceptionResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeExceptionResources")
	return
}

func NewDescribeExceptionResourcesResponse() (response *DescribeExceptionResourcesResponse) {
	response = &DescribeExceptionResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取异常资源列表
func (c *Client) DescribeExceptionResources(request *DescribeExceptionResourcesRequest) (response *DescribeExceptionResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeExceptionResourcesRequest()
	}
	response = NewDescribeExceptionResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExportsRequest() (request *DescribeExportsRequest) {
	request = &DescribeExportsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeExports")
	return
}

func NewDescribeExportsResponse() (response *DescribeExportsResponse) {
	response = &DescribeExportsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志下载任务列表
func (c *Client) DescribeExports(request *DescribeExportsRequest) (response *DescribeExportsResponse, err error) {
	if request == nil {
		request = NewDescribeExportsRequest()
	}
	response = NewDescribeExportsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMetricConfigRequest() (request *ModifyMetricConfigRequest) {
	request = &ModifyMetricConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyMetricConfig")
	return
}

func NewModifyMetricConfigResponse() (response *ModifyMetricConfigResponse) {
	response = &ModifyMetricConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改指标采集配置
func (c *Client) ModifyMetricConfig(request *ModifyMetricConfigRequest) (response *ModifyMetricConfigResponse, err error) {
	if request == nil {
		request = NewModifyMetricConfigRequest()
	}
	response = NewModifyMetricConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmShieldsRequest() (request *DescribeAlarmShieldsRequest) {
	request = &DescribeAlarmShieldsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarmShields")
	return
}

func NewDescribeAlarmShieldsResponse() (response *DescribeAlarmShieldsResponse) {
	response = &DescribeAlarmShieldsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警屏蔽配置规则
func (c *Client) DescribeAlarmShields(request *DescribeAlarmShieldsRequest) (response *DescribeAlarmShieldsResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmShieldsRequest()
	}
	response = NewDescribeAlarmShieldsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineGroupsRequest() (request *DescribeMachineGroupsRequest) {
	request = &DescribeMachineGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMachineGroups")
	return
}

func NewDescribeMachineGroupsResponse() (response *DescribeMachineGroupsResponse) {
	response = &DescribeMachineGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机器组信息列表
func (c *Client) DescribeMachineGroups(request *DescribeMachineGroupsRequest) (response *DescribeMachineGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeMachineGroupsRequest()
	}
	response = NewDescribeMachineGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchDashboardSubscribeRequest() (request *SearchDashboardSubscribeRequest) {
	request = &SearchDashboardSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "SearchDashboardSubscribe")
	return
}

func NewSearchDashboardSubscribeResponse() (response *SearchDashboardSubscribeResponse) {
	response = &SearchDashboardSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于预览仪表盘订阅
func (c *Client) SearchDashboardSubscribe(request *SearchDashboardSubscribeRequest) (response *SearchDashboardSubscribeResponse, err error) {
	if request == nil {
		request = NewSearchDashboardSubscribeRequest()
	}
	response = NewSearchDashboardSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAlarmChannelRequest() (request *CheckAlarmChannelRequest) {
	request = &CheckAlarmChannelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckAlarmChannel")
	return
}

func NewCheckAlarmChannelResponse() (response *CheckAlarmChannelResponse) {
	response = &CheckAlarmChannelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警渠道检测接口
func (c *Client) CheckAlarmChannel(request *CheckAlarmChannelRequest) (response *CheckAlarmChannelResponse, err error) {
	if request == nil {
		request = NewCheckAlarmChannelRequest()
	}
	response = NewCheckAlarmChannelResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDashboardRequest() (request *ModifyDashboardRequest) {
	request = &ModifyDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyDashboard")
	return
}

func NewModifyDashboardResponse() (response *ModifyDashboardResponse) {
	response = &ModifyDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改仪表盘
func (c *Client) ModifyDashboard(request *ModifyDashboardRequest) (response *ModifyDashboardResponse, err error) {
	if request == nil {
		request = NewModifyDashboardRequest()
	}
	response = NewModifyDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateTopic")
	return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
	response = &CreateTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建日志主题。
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	if request == nil {
		request = NewCreateTopicRequest()
	}
	response = NewCreateTopicResponse()
	err = c.Send(request, response)
	return
}

func NewGetMetricLabelValuesRequest() (request *GetMetricLabelValuesRequest) {
	request = &GetMetricLabelValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GetMetricLabelValues")
	return
}

func NewGetMetricLabelValuesResponse() (response *GetMetricLabelValuesResponse) {
	response = &GetMetricLabelValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取时序label values列表
func (c *Client) GetMetricLabelValues(request *GetMetricLabelValuesRequest) (response *GetMetricLabelValuesResponse, err error) {
	if request == nil {
		request = NewGetMetricLabelValuesRequest()
	}
	response = NewGetMetricLabelValuesResponse()
	err = c.Send(request, response)
	return
}

func NewOpenKafkaConsumeRequest() (request *OpenKafkaConsumeRequest) {
	request = &OpenKafkaConsumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "OpenKafkaConsume")
	return
}

func NewOpenKafkaConsumeResponse() (response *OpenKafkaConsumeResponse) {
	response = &OpenKafkaConsumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 打开kafka消费功能
func (c *Client) OpenKafkaConsume(request *OpenKafkaConsumeRequest) (response *OpenKafkaConsumeResponse, err error) {
	if request == nil {
		request = NewOpenKafkaConsumeRequest()
	}
	response = NewOpenKafkaConsumeResponse()
	err = c.Send(request, response)
	return
}

func NewUpgradeAgentNormalRequest() (request *UpgradeAgentNormalRequest) {
	request = &UpgradeAgentNormalRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UpgradeAgentNormal")
	return
}

func NewUpgradeAgentNormalResponse() (response *UpgradeAgentNormalResponse) {
	response = &UpgradeAgentNormalResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 基于Agent粒度的单次升级任务启动
func (c *Client) UpgradeAgentNormal(request *UpgradeAgentNormalRequest) (response *UpgradeAgentNormalResponse, err error) {
	if request == nil {
		request = NewUpgradeAgentNormalRequest()
	}
	response = NewUpgradeAgentNormalResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRecordingRuleTaskRequest() (request *ModifyRecordingRuleTaskRequest) {
	request = &ModifyRecordingRuleTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyRecordingRuleTask")
	return
}

func NewModifyRecordingRuleTaskResponse() (response *ModifyRecordingRuleTaskResponse) {
	response = &ModifyRecordingRuleTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改预聚合任务
func (c *Client) ModifyRecordingRuleTask(request *ModifyRecordingRuleTaskRequest) (response *ModifyRecordingRuleTaskResponse, err error) {
	if request == nil {
		request = NewModifyRecordingRuleTaskRequest()
	}
	response = NewModifyRecordingRuleTaskResponse()
	err = c.Send(request, response)
	return
}

func NewPreviewKafkaRechargeRequest() (request *PreviewKafkaRechargeRequest) {
	request = &PreviewKafkaRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "PreviewKafkaRecharge")
	return
}

func NewPreviewKafkaRechargeResponse() (response *PreviewKafkaRechargeResponse) {
	response = &PreviewKafkaRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于预览Kafka数据订阅任务客户日志信息
func (c *Client) PreviewKafkaRecharge(request *PreviewKafkaRechargeRequest) (response *PreviewKafkaRechargeResponse, err error) {
	if request == nil {
		request = NewPreviewKafkaRechargeRequest()
	}
	response = NewPreviewKafkaRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMetricSubscribeRequest() (request *CreateMetricSubscribeRequest) {
	request = &CreateMetricSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateMetricSubscribe")
	return
}

func NewCreateMetricSubscribeResponse() (response *CreateMetricSubscribeResponse) {
	response = &CreateMetricSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建指标订阅配置
func (c *Client) CreateMetricSubscribe(request *CreateMetricSubscribeRequest) (response *CreateMetricSubscribeResponse, err error) {
	if request == nil {
		request = NewCreateMetricSubscribeRequest()
	}
	response = NewCreateMetricSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConfigRequest() (request *ModifyConfigRequest) {
	request = &ModifyConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyConfig")
	return
}

func NewModifyConfigResponse() (response *ModifyConfigResponse) {
	response = &ModifyConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改采集规则配置
func (c *Client) ModifyConfig(request *ModifyConfigRequest) (response *ModifyConfigResponse, err error) {
	if request == nil {
		request = NewModifyConfigRequest()
	}
	response = NewModifyConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyScheduledSqlRequest() (request *ModifyScheduledSqlRequest) {
	request = &ModifyScheduledSqlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyScheduledSql")
	return
}

func NewModifyScheduledSqlResponse() (response *ModifyScheduledSqlResponse) {
	response = &ModifyScheduledSqlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改定时SQL分析任务
func (c *Client) ModifyScheduledSql(request *ModifyScheduledSqlRequest) (response *ModifyScheduledSqlResponse, err error) {
	if request == nil {
		request = NewModifyScheduledSqlRequest()
	}
	response = NewModifyScheduledSqlResponse()
	err = c.Send(request, response)
	return
}

func NewModifyKafkaRechargeRequest() (request *ModifyKafkaRechargeRequest) {
	request = &ModifyKafkaRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyKafkaRecharge")
	return
}

func NewModifyKafkaRechargeResponse() (response *ModifyKafkaRechargeResponse) {
	response = &ModifyKafkaRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改Kafka数据订阅任务
func (c *Client) ModifyKafkaRecharge(request *ModifyKafkaRechargeRequest) (response *ModifyKafkaRechargeResponse, err error) {
	if request == nil {
		request = NewModifyKafkaRechargeRequest()
	}
	response = NewModifyKafkaRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBinlogSubscribeRequest() (request *CreateBinlogSubscribeRequest) {
	request = &CreateBinlogSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateBinlogSubscribe")
	return
}

func NewCreateBinlogSubscribeResponse() (response *CreateBinlogSubscribeResponse) {
	response = &CreateBinlogSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建binlog采集配置
func (c *Client) CreateBinlogSubscribe(request *CreateBinlogSubscribeRequest) (response *CreateBinlogSubscribeResponse, err error) {
	if request == nil {
		request = NewCreateBinlogSubscribeRequest()
	}
	response = NewCreateBinlogSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCosRechargeRequest() (request *DeleteCosRechargeRequest) {
	request = &DeleteCosRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteCosRecharge")
	return
}

func NewDeleteCosRechargeResponse() (response *DeleteCosRechargeResponse) {
	response = &DeleteCosRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除cos导入任务
func (c *Client) DeleteCosRecharge(request *DeleteCosRechargeRequest) (response *DeleteCosRechargeResponse, err error) {
	if request == nil {
		request = NewDeleteCosRechargeRequest()
	}
	response = NewDeleteCosRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigExtraRequest() (request *DeleteConfigExtraRequest) {
	request = &DeleteConfigExtraRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConfigExtra")
	return
}

func NewDeleteConfigExtraResponse() (response *DeleteConfigExtraResponse) {
	response = &DeleteConfigExtraResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除特殊采集规则配置
func (c *Client) DeleteConfigExtra(request *DeleteConfigExtraRequest) (response *DeleteConfigExtraResponse, err error) {
	if request == nil {
		request = NewDeleteConfigExtraRequest()
	}
	response = NewDeleteConfigExtraResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRemoteWriteTaskRequest() (request *DeleteRemoteWriteTaskRequest) {
	request = &DeleteRemoteWriteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteRemoteWriteTask")
	return
}

func NewDeleteRemoteWriteTaskResponse() (response *DeleteRemoteWriteTaskResponse) {
	response = &DeleteRemoteWriteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RemoteWrite任务
func (c *Client) DeleteRemoteWriteTask(request *DeleteRemoteWriteTaskRequest) (response *DeleteRemoteWriteTaskResponse, err error) {
	if request == nil {
		request = NewDeleteRemoteWriteTaskRequest()
	}
	response = NewDeleteRemoteWriteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBinlogSubscribePreviewRequest() (request *DescribeBinlogSubscribePreviewRequest) {
	request = &DescribeBinlogSubscribePreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeBinlogSubscribePreview")
	return
}

func NewDescribeBinlogSubscribePreviewResponse() (response *DescribeBinlogSubscribePreviewResponse) {
	response = &DescribeBinlogSubscribePreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// binlog采集预览
func (c *Client) DescribeBinlogSubscribePreview(request *DescribeBinlogSubscribePreviewRequest) (response *DescribeBinlogSubscribePreviewResponse, err error) {
	if request == nil {
		request = NewDescribeBinlogSubscribePreviewRequest()
	}
	response = NewDescribeBinlogSubscribePreviewResponse()
	err = c.Send(request, response)
	return
}

func NewGetClsServiceRequest() (request *GetClsServiceRequest) {
	request = &GetClsServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GetClsService")
	return
}

func NewGetClsServiceResponse() (response *GetClsServiceResponse) {
	response = &GetClsServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询日志服务是否开通
func (c *Client) GetClsService(request *GetClsServiceRequest) (response *GetClsServiceResponse, err error) {
	if request == nil {
		request = NewGetClsServiceRequest()
	}
	response = NewGetClsServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigExtrasRequest() (request *DescribeConfigExtrasRequest) {
	request = &DescribeConfigExtrasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigExtras")
	return
}

func NewDescribeConfigExtrasResponse() (response *DescribeConfigExtrasResponse) {
	response = &DescribeConfigExtrasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取特殊采集配置，特殊采集配置应用于自建K8S环境的采集Agent
func (c *Client) DescribeConfigExtras(request *DescribeConfigExtrasRequest) (response *DescribeConfigExtrasResponse, err error) {
	if request == nil {
		request = NewDescribeConfigExtrasRequest()
	}
	response = NewDescribeConfigExtrasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsumerPreviewRequest() (request *DescribeConsumerPreviewRequest) {
	request = &DescribeConsumerPreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumerPreview")
	return
}

func NewDescribeConsumerPreviewResponse() (response *DescribeConsumerPreviewResponse) {
	response = &DescribeConsumerPreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于kafka投递数据预览
func (c *Client) DescribeConsumerPreview(request *DescribeConsumerPreviewRequest) (response *DescribeConsumerPreviewResponse, err error) {
	if request == nil {
		request = NewDescribeConsumerPreviewRequest()
	}
	response = NewDescribeConsumerPreviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyUserConfigRequest() (request *ModifyUserConfigRequest) {
	request = &ModifyUserConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyUserConfig")
	return
}

func NewModifyUserConfigResponse() (response *ModifyUserConfigResponse) {
	response = &ModifyUserConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改用户配置
func (c *Client) ModifyUserConfig(request *ModifyUserConfigRequest) (response *ModifyUserConfigResponse, err error) {
	if request == nil {
		request = NewModifyUserConfigRequest()
	}
	response = NewModifyUserConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateScheduledSqlRequest() (request *CreateScheduledSqlRequest) {
	request = &CreateScheduledSqlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateScheduledSql")
	return
}

func NewCreateScheduledSqlResponse() (response *CreateScheduledSqlResponse) {
	response = &CreateScheduledSqlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建定时SQL分析任务
func (c *Client) CreateScheduledSql(request *CreateScheduledSqlRequest) (response *CreateScheduledSqlResponse, err error) {
	if request == nil {
		request = NewCreateScheduledSqlRequest()
	}
	response = NewCreateScheduledSqlResponse()
	err = c.Send(request, response)
	return
}

func NewOpenClsServiceRequest() (request *OpenClsServiceRequest) {
	request = &OpenClsServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "OpenClsService")
	return
}

func NewOpenClsServiceResponse() (response *OpenClsServiceResponse) {
	response = &OpenClsServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 开通日志服务
func (c *Client) OpenClsService(request *OpenClsServiceRequest) (response *OpenClsServiceResponse, err error) {
	if request == nil {
		request = NewOpenClsServiceRequest()
	}
	response = NewOpenClsServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndexRequest() (request *DescribeIndexRequest) {
	request = &DescribeIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeIndex")
	return
}

func NewDescribeIndexResponse() (response *DescribeIndexResponse) {
	response = &DescribeIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取索引配置信息
func (c *Client) DescribeIndex(request *DescribeIndexRequest) (response *DescribeIndexResponse, err error) {
	if request == nil {
		request = NewDescribeIndexRequest()
	}
	response = NewDescribeIndexResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlarmNoticeRequest() (request *DeleteAlarmNoticeRequest) {
	request = &DeleteAlarmNoticeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarmNotice")
	return
}

func NewDeleteAlarmNoticeResponse() (response *DeleteAlarmNoticeResponse) {
	response = &DeleteAlarmNoticeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于删除告警通知模板
func (c *Client) DeleteAlarmNotice(request *DeleteAlarmNoticeRequest) (response *DeleteAlarmNoticeResponse, err error) {
	if request == nil {
		request = NewDeleteAlarmNoticeRequest()
	}
	response = NewDeleteAlarmNoticeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteWriteTasksRequest() (request *DescribeRemoteWriteTasksRequest) {
	request = &DescribeRemoteWriteTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeRemoteWriteTasks")
	return
}

func NewDescribeRemoteWriteTasksResponse() (response *DescribeRemoteWriteTasksResponse) {
	response = &DescribeRemoteWriteTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取RemoteWrite投递任务列表页
func (c *Client) DescribeRemoteWriteTasks(request *DescribeRemoteWriteTasksRequest) (response *DescribeRemoteWriteTasksResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteWriteTasksRequest()
	}
	response = NewDescribeRemoteWriteTasksResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRebuildIndexTaskRequest() (request *CreateRebuildIndexTaskRequest) {
	request = &CreateRebuildIndexTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateRebuildIndexTask")
	return
}

func NewCreateRebuildIndexTaskResponse() (response *CreateRebuildIndexTaskResponse) {
	response = &CreateRebuildIndexTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建重建索引任务
func (c *Client) CreateRebuildIndexTask(request *CreateRebuildIndexTaskRequest) (response *CreateRebuildIndexTaskResponse, err error) {
	if request == nil {
		request = NewCreateRebuildIndexTaskRequest()
	}
	response = NewCreateRebuildIndexTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterMetricConfigsRequest() (request *DescribeClusterMetricConfigsRequest) {
	request = &DescribeClusterMetricConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeClusterMetricConfigs")
	return
}

func NewDescribeClusterMetricConfigsResponse() (response *DescribeClusterMetricConfigsResponse) {
	response = &DescribeClusterMetricConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群指标采集配置
func (c *Client) DescribeClusterMetricConfigs(request *DescribeClusterMetricConfigsRequest) (response *DescribeClusterMetricConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterMetricConfigsRequest()
	}
	response = NewDescribeClusterMetricConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateFolderRequest() (request *CreateFolderRequest) {
	request = &CreateFolderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateFolder")
	return
}

func NewCreateFolderResponse() (response *CreateFolderResponse) {
	response = &CreateFolderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建文件夹
func (c *Client) CreateFolder(request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
	if request == nil {
		request = NewCreateFolderRequest()
	}
	response = NewCreateFolderResponse()
	err = c.Send(request, response)
	return
}

func NewQueryRangeMetricRequest() (request *QueryRangeMetricRequest) {
	request = &QueryRangeMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "QueryRangeMetric")
	return
}

func NewQueryRangeMetricResponse() (response *QueryRangeMetricResponse) {
	response = &QueryRangeMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 时序范围检索
func (c *Client) QueryRangeMetric(request *QueryRangeMetricRequest) (response *QueryRangeMetricResponse, err error) {
	if request == nil {
		request = NewQueryRangeMetricRequest()
	}
	response = NewQueryRangeMetricResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExternalDataSourceRequest() (request *DeleteExternalDataSourceRequest) {
	request = &DeleteExternalDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteExternalDataSource")
	return
}

func NewDeleteExternalDataSourceResponse() (response *DeleteExternalDataSourceResponse) {
	response = &DeleteExternalDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除外部数据源
func (c *Client) DeleteExternalDataSource(request *DeleteExternalDataSourceRequest) (response *DeleteExternalDataSourceResponse, err error) {
	if request == nil {
		request = NewDeleteExternalDataSourceRequest()
	}
	response = NewDeleteExternalDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewHeartBeatRequest() (request *HeartBeatRequest) {
	request = &HeartBeatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "HeartBeat")
	return
}

func NewHeartBeatResponse() (response *HeartBeatResponse) {
	response = &HeartBeatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上报当前采集机器的心跳
func (c *Client) HeartBeat(request *HeartBeatRequest) (response *HeartBeatResponse, err error) {
	if request == nil {
		request = NewHeartBeatRequest()
	}
	response = NewHeartBeatResponse()
	err = c.Send(request, response)
	return
}

func NewCreateIndexRequest() (request *CreateIndexRequest) {
	request = &CreateIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateIndex")
	return
}

func NewCreateIndexResponse() (response *CreateIndexResponse) {
	response = &CreateIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建索引
func (c *Client) CreateIndex(request *CreateIndexRequest) (response *CreateIndexResponse, err error) {
	if request == nil {
		request = NewCreateIndexRequest()
	}
	response = NewCreateIndexResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecordingRuleYamlTaskRequest() (request *DescribeRecordingRuleYamlTaskRequest) {
	request = &DescribeRecordingRuleYamlTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeRecordingRuleYamlTask")
	return
}

func NewDescribeRecordingRuleYamlTaskResponse() (response *DescribeRecordingRuleYamlTaskResponse) {
	response = &DescribeRecordingRuleYamlTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取YAML预聚合任务列表
func (c *Client) DescribeRecordingRuleYamlTask(request *DescribeRecordingRuleYamlTaskRequest) (response *DescribeRecordingRuleYamlTaskResponse, err error) {
	if request == nil {
		request = NewDescribeRecordingRuleYamlTaskRequest()
	}
	response = NewDescribeRecordingRuleYamlTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigMachineGroupsRequest() (request *DescribeConfigMachineGroupsRequest) {
	request = &DescribeConfigMachineGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigMachineGroups")
	return
}

func NewDescribeConfigMachineGroupsResponse() (response *DescribeConfigMachineGroupsResponse) {
	response = &DescribeConfigMachineGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集规则配置所绑定的机器组
func (c *Client) DescribeConfigMachineGroups(request *DescribeConfigMachineGroupsRequest) (response *DescribeConfigMachineGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeConfigMachineGroupsRequest()
	}
	response = NewDescribeConfigMachineGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalDataSourcesRequest() (request *DescribeExternalDataSourcesRequest) {
	request = &DescribeExternalDataSourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeExternalDataSources")
	return
}

func NewDescribeExternalDataSourcesResponse() (response *DescribeExternalDataSourcesResponse) {
	response = &DescribeExternalDataSourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取外部数据源信息列表
func (c *Client) DescribeExternalDataSources(request *DescribeExternalDataSourcesRequest) (response *DescribeExternalDataSourcesResponse, err error) {
	if request == nil {
		request = NewDescribeExternalDataSourcesRequest()
	}
	response = NewDescribeExternalDataSourcesResponse()
	err = c.Send(request, response)
	return
}

func NewApplyConfigurationTemplateRequest() (request *ApplyConfigurationTemplateRequest) {
	request = &ApplyConfigurationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ApplyConfigurationTemplate")
	return
}

func NewApplyConfigurationTemplateResponse() (response *ApplyConfigurationTemplateResponse) {
	response = &ApplyConfigurationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用配置模板
func (c *Client) ApplyConfigurationTemplate(request *ApplyConfigurationTemplateRequest) (response *ApplyConfigurationTemplateResponse, err error) {
	if request == nil {
		request = NewApplyConfigurationTemplateRequest()
	}
	response = NewApplyConfigurationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicMetricConfigsRequest() (request *DescribeTopicMetricConfigsRequest) {
	request = &DescribeTopicMetricConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeTopicMetricConfigs")
	return
}

func NewDescribeTopicMetricConfigsResponse() (response *DescribeTopicMetricConfigsResponse) {
	response = &DescribeTopicMetricConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指标采集配置
func (c *Client) DescribeTopicMetricConfigs(request *DescribeTopicMetricConfigsRequest) (response *DescribeTopicMetricConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeTopicMetricConfigsRequest()
	}
	response = NewDescribeTopicMetricConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCheckUserIDRequest() (request *CheckUserIDRequest) {
	request = &CheckUserIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckUserID")
	return
}

func NewCheckUserIDResponse() (response *CheckUserIDResponse) {
	response = &CheckUserIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验用户ID是否符合规范
func (c *Client) CheckUserID(request *CheckUserIDRequest) (response *CheckUserIDResponse, err error) {
	if request == nil {
		request = NewCheckUserIDRequest()
	}
	response = NewCheckUserIDResponse()
	err = c.Send(request, response)
	return
}

func NewGetMetricSeriesRequest() (request *GetMetricSeriesRequest) {
	request = &GetMetricSeriesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GetMetricSeries")
	return
}

func NewGetMetricSeriesResponse() (response *GetMetricSeriesResponse) {
	response = &GetMetricSeriesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取时序label values列表
func (c *Client) GetMetricSeries(request *GetMetricSeriesRequest) (response *GetMetricSeriesResponse, err error) {
	if request == nil {
		request = NewGetMetricSeriesRequest()
	}
	response = NewGetMetricSeriesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRemoteWriteTaskRequest() (request *CreateRemoteWriteTaskRequest) {
	request = &CreateRemoteWriteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateRemoteWriteTask")
	return
}

func NewCreateRemoteWriteTaskResponse() (response *CreateRemoteWriteTaskResponse) {
	response = &CreateRemoteWriteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建RemoteWrite任务
func (c *Client) CreateRemoteWriteTask(request *CreateRemoteWriteTaskRequest) (response *CreateRemoteWriteTaskResponse, err error) {
	if request == nil {
		request = NewCreateRemoteWriteTaskRequest()
	}
	response = NewCreateRemoteWriteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserConfigRequest() (request *DescribeUserConfigRequest) {
	request = &DescribeUserConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeUserConfig")
	return
}

func NewDescribeUserConfigResponse() (response *DescribeUserConfigResponse) {
	response = &DescribeUserConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户配置信息
func (c *Client) DescribeUserConfig(request *DescribeUserConfigRequest) (response *DescribeUserConfigResponse, err error) {
	if request == nil {
		request = NewDescribeUserConfigRequest()
	}
	response = NewDescribeUserConfigResponse()
	err = c.Send(request, response)
	return
}

func NewGetConfigurationTemplateApplyLogRequest() (request *GetConfigurationTemplateApplyLogRequest) {
	request = &GetConfigurationTemplateApplyLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GetConfigurationTemplateApplyLog")
	return
}

func NewGetConfigurationTemplateApplyLogResponse() (response *GetConfigurationTemplateApplyLogResponse) {
	response = &GetConfigurationTemplateApplyLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置模板应用日志
func (c *Client) GetConfigurationTemplateApplyLog(request *GetConfigurationTemplateApplyLogRequest) (response *GetConfigurationTemplateApplyLogResponse, err error) {
	if request == nil {
		request = NewGetConfigurationTemplateApplyLogRequest()
	}
	response = NewGetConfigurationTemplateApplyLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCosRechargesRequest() (request *DescribeCosRechargesRequest) {
	request = &DescribeCosRechargesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeCosRecharges")
	return
}

func NewDescribeCosRechargesResponse() (response *DescribeCosRechargesResponse) {
	response = &DescribeCosRechargesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取对象存储导入配置
func (c *Client) DescribeCosRecharges(request *DescribeCosRechargesRequest) (response *DescribeCosRechargesResponse, err error) {
	if request == nil {
		request = NewDescribeCosRechargesRequest()
	}
	response = NewDescribeCosRechargesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentConfigsRequest() (request *DescribeAgentConfigsRequest) {
	request = &DescribeAgentConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAgentConfigs")
	return
}

func NewDescribeAgentConfigsResponse() (response *DescribeAgentConfigsResponse) {
	response = &DescribeAgentConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取agent对应的采集配置
func (c *Client) DescribeAgentConfigs(request *DescribeAgentConfigsRequest) (response *DescribeAgentConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeAgentConfigsRequest()
	}
	response = NewDescribeAgentConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewAddMachineGroupInfoRequest() (request *AddMachineGroupInfoRequest) {
	request = &AddMachineGroupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "AddMachineGroupInfo")
	return
}

func NewAddMachineGroupInfoResponse() (response *AddMachineGroupInfoResponse) {
	response = &AddMachineGroupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于添加机器组信息
func (c *Client) AddMachineGroupInfo(request *AddMachineGroupInfoRequest) (response *AddMachineGroupInfoResponse, err error) {
	if request == nil {
		request = NewAddMachineGroupInfoRequest()
	}
	response = NewAddMachineGroupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecordingRuleProcessInfoRequest() (request *DescribeRecordingRuleProcessInfoRequest) {
	request = &DescribeRecordingRuleProcessInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeRecordingRuleProcessInfo")
	return
}

func NewDescribeRecordingRuleProcessInfoResponse() (response *DescribeRecordingRuleProcessInfoResponse) {
	response = &DescribeRecordingRuleProcessInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取预聚合分析任务进度信息
func (c *Client) DescribeRecordingRuleProcessInfo(request *DescribeRecordingRuleProcessInfoRequest) (response *DescribeRecordingRuleProcessInfoResponse, err error) {
	if request == nil {
		request = NewDescribeRecordingRuleProcessInfoRequest()
	}
	response = NewDescribeRecordingRuleProcessInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMetricSubscribeRequest() (request *ModifyMetricSubscribeRequest) {
	request = &ModifyMetricSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyMetricSubscribe")
	return
}

func NewModifyMetricSubscribeResponse() (response *ModifyMetricSubscribeResponse) {
	response = &ModifyMetricSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改指标订阅配置
func (c *Client) ModifyMetricSubscribe(request *ModifyMetricSubscribeRequest) (response *ModifyMetricSubscribeResponse, err error) {
	if request == nil {
		request = NewModifyMetricSubscribeRequest()
	}
	response = NewModifyMetricSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteNoticeContentRequest() (request *DeleteNoticeContentRequest) {
	request = &DeleteNoticeContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteNoticeContent")
	return
}

func NewDeleteNoticeContentResponse() (response *DeleteNoticeContentResponse) {
	response = &DeleteNoticeContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除通知内容模板
func (c *Client) DeleteNoticeContent(request *DeleteNoticeContentRequest) (response *DeleteNoticeContentResponse, err error) {
	if request == nil {
		request = NewDeleteNoticeContentRequest()
	}
	response = NewDeleteNoticeContentResponse()
	err = c.Send(request, response)
	return
}

func NewGetAlarmLogRequest() (request *GetAlarmLogRequest) {
	request = &GetAlarmLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GetAlarmLog")
	return
}

func NewGetAlarmLogResponse() (response *GetAlarmLogResponse) {
	response = &GetAlarmLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取告警任务历史
func (c *Client) GetAlarmLog(request *GetAlarmLogRequest) (response *GetAlarmLogResponse, err error) {
	if request == nil {
		request = NewGetAlarmLogRequest()
	}
	response = NewGetAlarmLogResponse()
	err = c.Send(request, response)
	return
}

func NewUploadLogRequest() (request *UploadLogRequest) {
	request = &UploadLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UploadLog")
	return
}

func NewUploadLogResponse() (response *UploadLogResponse) {
	response = &UploadLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 为了保障您日志数据的可靠性以及更高效地使用日志服务，建议您使用CLS优化后的接口[上传结构化日志]
func (c *Client) UploadLog(request *UploadLogRequest) (response *UploadLogResponse, err error) {
	if request == nil {
		request = NewUploadLogRequest()
	}
	response = NewUploadLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourcesRequest() (request *DescribeResourcesRequest) {
	request = &DescribeResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeResources")
	return
}

func NewDescribeResourcesResponse() (response *DescribeResourcesResponse) {
	response = &DescribeResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取全局或指定地域指标资源
func (c *Client) DescribeResources(request *DescribeResourcesRequest) (response *DescribeResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeResourcesRequest()
	}
	response = NewDescribeResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlarmRequest() (request *DeleteAlarmRequest) {
	request = &DeleteAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarm")
	return
}

func NewDeleteAlarmResponse() (response *DeleteAlarmResponse) {
	response = &DeleteAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除告警策略。
func (c *Client) DeleteAlarm(request *DeleteAlarmRequest) (response *DeleteAlarmResponse, err error) {
	if request == nil {
		request = NewDeleteAlarmRequest()
	}
	response = NewDeleteAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteKafkaRechargeRequest() (request *DeleteKafkaRechargeRequest) {
	request = &DeleteKafkaRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteKafkaRecharge")
	return
}

func NewDeleteKafkaRechargeResponse() (response *DeleteKafkaRechargeResponse) {
	response = &DeleteKafkaRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除Kafka数据订阅任务
func (c *Client) DeleteKafkaRecharge(request *DeleteKafkaRechargeRequest) (response *DeleteKafkaRechargeResponse, err error) {
	if request == nil {
		request = NewDeleteKafkaRechargeRequest()
	}
	response = NewDeleteKafkaRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConfigExtraRequest() (request *ModifyConfigExtraRequest) {
	request = &ModifyConfigExtraRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyConfigExtra")
	return
}

func NewModifyConfigExtraResponse() (response *ModifyConfigExtraResponse) {
	response = &ModifyConfigExtraResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改特殊采集配置任务
func (c *Client) ModifyConfigExtra(request *ModifyConfigExtraRequest) (response *ModifyConfigExtraResponse, err error) {
	if request == nil {
		request = NewModifyConfigExtraRequest()
	}
	response = NewModifyConfigExtraResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmsRequest() (request *DescribeAlarmsRequest) {
	request = &DescribeAlarmsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarms")
	return
}

func NewDescribeAlarmsResponse() (response *DescribeAlarmsResponse) {
	response = &DescribeAlarmsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取告警策略。
func (c *Client) DescribeAlarms(request *DescribeAlarmsRequest) (response *DescribeAlarmsResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmsRequest()
	}
	response = NewDescribeAlarmsResponse()
	err = c.Send(request, response)
	return
}

func NewSearchCosRechargeInfoRequest() (request *SearchCosRechargeInfoRequest) {
	request = &SearchCosRechargeInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "SearchCosRechargeInfo")
	return
}

func NewSearchCosRechargeInfoResponse() (response *SearchCosRechargeInfoResponse) {
	response = &SearchCosRechargeInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于预览cos导入信息
func (c *Client) SearchCosRechargeInfo(request *SearchCosRechargeInfoRequest) (response *SearchCosRechargeInfoResponse, err error) {
	if request == nil {
		request = NewSearchCosRechargeInfoRequest()
	}
	response = NewSearchCosRechargeInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsumersRequest() (request *DescribeConsumersRequest) {
	request = &DescribeConsumersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumers")
	return
}

func NewDescribeConsumersResponse() (response *DescribeConsumersResponse) {
	response = &DescribeConsumersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取投递到Ckafka的任务配置列表
func (c *Client) DescribeConsumers(request *DescribeConsumersRequest) (response *DescribeConsumersResponse, err error) {
	if request == nil {
		request = NewDescribeConsumersRequest()
	}
	response = NewDescribeConsumersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDashboardSubscribesRequest() (request *DescribeDashboardSubscribesRequest) {
	request = &DescribeDashboardSubscribesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDashboardSubscribes")
	return
}

func NewDescribeDashboardSubscribesResponse() (response *DescribeDashboardSubscribesResponse) {
	response = &DescribeDashboardSubscribesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取仪表盘订阅列表，支持分页
func (c *Client) DescribeDashboardSubscribes(request *DescribeDashboardSubscribesRequest) (response *DescribeDashboardSubscribesResponse, err error) {
	if request == nil {
		request = NewDescribeDashboardSubscribesRequest()
	}
	response = NewDescribeDashboardSubscribesResponse()
	err = c.Send(request, response)
	return
}

func NewGenBeginRegexRequest() (request *GenBeginRegexRequest) {
	request = &GenBeginRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GenBeginRegex")
	return
}

func NewGenBeginRegexResponse() (response *GenBeginRegexResponse) {
	response = &GenBeginRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成首行正则表达式
func (c *Client) GenBeginRegex(request *GenBeginRegexRequest) (response *GenBeginRegexResponse, err error) {
	if request == nil {
		request = NewGenBeginRegexRequest()
	}
	response = NewGenBeginRegexResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNoticeContentsRequest() (request *DescribeNoticeContentsRequest) {
	request = &DescribeNoticeContentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeNoticeContents")
	return
}

func NewDescribeNoticeContentsResponse() (response *DescribeNoticeContentsResponse) {
	response = &DescribeNoticeContentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取通知内容模板
func (c *Client) DescribeNoticeContents(request *DescribeNoticeContentsRequest) (response *DescribeNoticeContentsResponse, err error) {
	if request == nil {
		request = NewDescribeNoticeContentsRequest()
	}
	response = NewDescribeNoticeContentsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDashboardRequest() (request *CreateDashboardRequest) {
	request = &CreateDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateDashboard")
	return
}

func NewCreateDashboardResponse() (response *CreateDashboardResponse) {
	response = &CreateDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建仪表盘
func (c *Client) CreateDashboard(request *CreateDashboardRequest) (response *CreateDashboardResponse, err error) {
	if request == nil {
		request = NewCreateDashboardRequest()
	}
	response = NewCreateDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConsumerGroupRequest() (request *ModifyConsumerGroupRequest) {
	request = &ModifyConsumerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyConsumerGroup")
	return
}

func NewModifyConsumerGroupResponse() (response *ModifyConsumerGroupResponse) {
	response = &ModifyConsumerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新消费组信息
func (c *Client) ModifyConsumerGroup(request *ModifyConsumerGroupRequest) (response *ModifyConsumerGroupResponse, err error) {
	if request == nil {
		request = NewModifyConsumerGroupRequest()
	}
	response = NewModifyConsumerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachineGroupRequest() (request *DeleteMachineGroupRequest) {
	request = &DeleteMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteMachineGroup")
	return
}

func NewDeleteMachineGroupResponse() (response *DeleteMachineGroupResponse) {
	response = &DeleteMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除机器组
func (c *Client) DeleteMachineGroup(request *DeleteMachineGroupRequest) (response *DeleteMachineGroupResponse, err error) {
	if request == nil {
		request = NewDeleteMachineGroupRequest()
	}
	response = NewDeleteMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetricSubscribePreviewRequest() (request *DescribeMetricSubscribePreviewRequest) {
	request = &DescribeMetricSubscribePreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMetricSubscribePreview")
	return
}

func NewDescribeMetricSubscribePreviewResponse() (response *DescribeMetricSubscribePreviewResponse) {
	response = &DescribeMetricSubscribePreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 指标订阅数据预览
func (c *Client) DescribeMetricSubscribePreview(request *DescribeMetricSubscribePreviewRequest) (response *DescribeMetricSubscribePreviewResponse, err error) {
	if request == nil {
		request = NewDescribeMetricSubscribePreviewRequest()
	}
	response = NewDescribeMetricSubscribePreviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIndexRequest() (request *ModifyIndexRequest) {
	request = &ModifyIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyIndex")
	return
}

func NewModifyIndexResponse() (response *ModifyIndexResponse) {
	response = &ModifyIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改索引配置
func (c *Client) ModifyIndex(request *ModifyIndexRequest) (response *ModifyIndexResponse, err error) {
	if request == nil {
		request = NewModifyIndexRequest()
	}
	response = NewModifyIndexResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRemoteWriteTaskRequest() (request *ModifyRemoteWriteTaskRequest) {
	request = &ModifyRemoteWriteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyRemoteWriteTask")
	return
}

func NewModifyRemoteWriteTaskResponse() (response *ModifyRemoteWriteTaskResponse) {
	response = &ModifyRemoteWriteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改RemoteWrite任务
func (c *Client) ModifyRemoteWriteTask(request *ModifyRemoteWriteTaskRequest) (response *ModifyRemoteWriteTaskResponse, err error) {
	if request == nil {
		request = NewModifyRemoteWriteTaskRequest()
	}
	response = NewModifyRemoteWriteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewQueryMetricRequest() (request *QueryMetricRequest) {
	request = &QueryMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "QueryMetric")
	return
}

func NewQueryMetricResponse() (response *QueryMetricResponse) {
	response = &QueryMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 时序日志主题查询
func (c *Client) QueryMetric(request *QueryMetricRequest) (response *QueryMetricResponse, err error) {
	if request == nil {
		request = NewQueryMetricRequest()
	}
	response = NewQueryMetricResponse()
	err = c.Send(request, response)
	return
}

func NewSplitPartitionRequest() (request *SplitPartitionRequest) {
	request = &SplitPartitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "SplitPartition")
	return
}

func NewSplitPartitionResponse() (response *SplitPartitionResponse) {
	response = &SplitPartitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于分裂主题分区
func (c *Client) SplitPartition(request *SplitPartitionRequest) (response *SplitPartitionResponse, err error) {
	if request == nil {
		request = NewSplitPartitionRequest()
	}
	response = NewSplitPartitionResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDataTransformRequest() (request *CreateDataTransformRequest) {
	request = &CreateDataTransformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateDataTransform")
	return
}

func NewCreateDataTransformResponse() (response *CreateDataTransformResponse) {
	response = &CreateDataTransformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建数据加工任务。
func (c *Client) CreateDataTransform(request *CreateDataTransformRequest) (response *CreateDataTransformResponse, err error) {
	if request == nil {
		request = NewCreateDataTransformRequest()
	}
	response = NewCreateDataTransformResponse()
	err = c.Send(request, response)
	return
}

func NewCreateNoticeContentRequest() (request *CreateNoticeContentRequest) {
	request = &CreateNoticeContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateNoticeContent")
	return
}

func NewCreateNoticeContentResponse() (response *CreateNoticeContentResponse) {
	response = &CreateNoticeContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建通知内容模板
func (c *Client) CreateNoticeContent(request *CreateNoticeContentRequest) (response *CreateNoticeContentResponse, err error) {
	if request == nil {
		request = NewCreateNoticeContentRequest()
	}
	response = NewCreateNoticeContentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterBaseMetricConfigsRequest() (request *DescribeClusterBaseMetricConfigsRequest) {
	request = &DescribeClusterBaseMetricConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeClusterBaseMetricConfigs")
	return
}

func NewDescribeClusterBaseMetricConfigsResponse() (response *DescribeClusterBaseMetricConfigsResponse) {
	response = &DescribeClusterBaseMetricConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群指标基础监控采集配置
func (c *Client) DescribeClusterBaseMetricConfigs(request *DescribeClusterBaseMetricConfigsRequest) (response *DescribeClusterBaseMetricConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeClusterBaseMetricConfigsRequest()
	}
	response = NewDescribeClusterBaseMetricConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmShieldRequest() (request *ModifyAlarmShieldRequest) {
	request = &ModifyAlarmShieldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarmShield")
	return
}

func NewModifyAlarmShieldResponse() (response *ModifyAlarmShieldResponse) {
	response = &ModifyAlarmShieldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警屏蔽规则
func (c *Client) ModifyAlarmShield(request *ModifyAlarmShieldRequest) (response *ModifyAlarmShieldResponse, err error) {
	if request == nil {
		request = NewModifyAlarmShieldRequest()
	}
	response = NewModifyAlarmShieldResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBinlogSubscribesRequest() (request *DescribeBinlogSubscribesRequest) {
	request = &DescribeBinlogSubscribesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeBinlogSubscribes")
	return
}

func NewDescribeBinlogSubscribesResponse() (response *DescribeBinlogSubscribesResponse) {
	response = &DescribeBinlogSubscribesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取binlog采集配置
func (c *Client) DescribeBinlogSubscribes(request *DescribeBinlogSubscribesRequest) (response *DescribeBinlogSubscribesResponse, err error) {
	if request == nil {
		request = NewDescribeBinlogSubscribesRequest()
	}
	response = NewDescribeBinlogSubscribesResponse()
	err = c.Send(request, response)
	return
}

func NewGenKVRegexRequest() (request *GenKVRegexRequest) {
	request = &GenKVRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "GenKVRegex")
	return
}

func NewGenKVRegexResponse() (response *GenKVRegexResponse) {
	response = &GenKVRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 生成提取K-V形式的正则表达式
func (c *Client) GenKVRegex(request *GenKVRegexRequest) (response *GenKVRegexResponse, err error) {
	if request == nil {
		request = NewGenKVRegexRequest()
	}
	response = NewGenKVRegexResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShippersRequest() (request *DescribeShippersRequest) {
	request = &DescribeShippersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeShippers")
	return
}

func NewDescribeShippersResponse() (response *DescribeShippersResponse) {
	response = &DescribeShippersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取投递规则信息列表
func (c *Client) DescribeShippers(request *DescribeShippersRequest) (response *DescribeShippersResponse, err error) {
	if request == nil {
		request = NewDescribeShippersRequest()
	}
	response = NewDescribeShippersResponse()
	err = c.Send(request, response)
	return
}

func NewSearchLogRequest() (request *SearchLogRequest) {
	request = &SearchLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "SearchLog")
	return
}

func NewSearchLogResponse() (response *SearchLogResponse) {
	response = &SearchLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于检索分析日志, 该接口除受默认接口请求频率限制外，针对单个日志主题，查询并发数不能超过15。
func (c *Client) SearchLog(request *SearchLogRequest) (response *SearchLogResponse, err error) {
	if request == nil {
		request = NewSearchLogRequest()
	}
	response = NewSearchLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExportRequest() (request *CreateExportRequest) {
	request = &CreateExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateExport")
	return
}

func NewCreateExportResponse() (response *CreateExportResponse) {
	response = &CreateExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口仅创建下载任务，任务返回的下载地址，请用户调用DescribeExports查看任务列表。其中有下载地址CosPath参数。
func (c *Client) CreateExport(request *CreateExportRequest) (response *CreateExportResponse, err error) {
	if request == nil {
		request = NewCreateExportRequest()
	}
	response = NewCreateExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaUserRequest() (request *DescribeKafkaUserRequest) {
	request = &DescribeKafkaUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaUser")
	return
}

func NewDescribeKafkaUserResponse() (response *DescribeKafkaUserResponse) {
	response = &DescribeKafkaUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取kafka用户信息
func (c *Client) DescribeKafkaUser(request *DescribeKafkaUserRequest) (response *DescribeKafkaUserResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaUserRequest()
	}
	response = NewDescribeKafkaUserResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserConfigRequest() (request *DeleteUserConfigRequest) {
	request = &DeleteUserConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteUserConfig")
	return
}

func NewDeleteUserConfigResponse() (response *DeleteUserConfigResponse) {
	response = &DeleteUserConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户配置
func (c *Client) DeleteUserConfig(request *DeleteUserConfigRequest) (response *DeleteUserConfigResponse, err error) {
	if request == nil {
		request = NewDeleteUserConfigRequest()
	}
	response = NewDeleteUserConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigExtraRequest() (request *CreateConfigExtraRequest) {
	request = &CreateConfigExtraRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateConfigExtra")
	return
}

func NewCreateConfigExtraResponse() (response *CreateConfigExtraResponse) {
	response = &CreateConfigExtraResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建特殊采集配置任务
func (c *Client) CreateConfigExtra(request *CreateConfigExtraRequest) (response *CreateConfigExtraResponse, err error) {
	if request == nil {
		request = NewCreateConfigExtraRequest()
	}
	response = NewCreateConfigExtraResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsumerGroupsRequest() (request *DescribeConsumerGroupsRequest) {
	request = &DescribeConsumerGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumerGroups")
	return
}

func NewDescribeConsumerGroupsResponse() (response *DescribeConsumerGroupsResponse) {
	response = &DescribeConsumerGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取消费组列表
func (c *Client) DescribeConsumerGroups(request *DescribeConsumerGroupsRequest) (response *DescribeConsumerGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeConsumerGroupsRequest()
	}
	response = NewDescribeConsumerGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopicExtendConfigRequest() (request *CreateTopicExtendConfigRequest) {
	request = &CreateTopicExtendConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateTopicExtendConfig")
	return
}

func NewCreateTopicExtendConfigResponse() (response *CreateTopicExtendConfigResponse) {
	response = &CreateTopicExtendConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建采集配置, 不包含临时密钥信息(clb专用)。
func (c *Client) CreateTopicExtendConfig(request *CreateTopicExtendConfigRequest) (response *CreateTopicExtendConfigResponse, err error) {
	if request == nil {
		request = NewCreateTopicExtendConfigRequest()
	}
	response = NewCreateTopicExtendConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteIndexRequest() (request *DeleteIndexRequest) {
	request = &DeleteIndexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteIndex")
	return
}

func NewDeleteIndexResponse() (response *DeleteIndexResponse) {
	response = &DeleteIndexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于日志主题的索引配置
func (c *Client) DeleteIndex(request *DeleteIndexRequest) (response *DeleteIndexResponse, err error) {
	if request == nil {
		request = NewDeleteIndexRequest()
	}
	response = NewDeleteIndexResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCosRechargeRequest() (request *CreateCosRechargeRequest) {
	request = &CreateCosRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateCosRecharge")
	return
}

func NewCreateCosRechargeResponse() (response *CreateCosRechargeResponse) {
	response = &CreateCosRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建对象存储导入任务
func (c *Client) CreateCosRecharge(request *CreateCosRechargeRequest) (response *CreateCosRechargeResponse, err error) {
	if request == nil {
		request = NewCreateCosRechargeRequest()
	}
	response = NewCreateCosRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplatesRequest() (request *DescribeTemplatesRequest) {
	request = &DescribeTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeTemplates")
	return
}

func NewDescribeTemplatesResponse() (response *DescribeTemplatesResponse) {
	response = &DescribeTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取预置的仪表盘列表
func (c *Client) DescribeTemplates(request *DescribeTemplatesRequest) (response *DescribeTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeTemplatesRequest()
	}
	response = NewDescribeTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigRequest() (request *DeleteConfigRequest) {
	request = &DeleteConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConfig")
	return
}

func NewDeleteConfigResponse() (response *DeleteConfigResponse) {
	response = &DeleteConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除采集规则配置
func (c *Client) DeleteConfig(request *DeleteConfigRequest) (response *DeleteConfigResponse, err error) {
	if request == nil {
		request = NewDeleteConfigRequest()
	}
	response = NewDeleteConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShipperPreviewRequest() (request *DescribeShipperPreviewRequest) {
	request = &DescribeShipperPreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeShipperPreview")
	return
}

func NewDescribeShipperPreviewResponse() (response *DescribeShipperPreviewResponse) {
	response = &DescribeShipperPreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于返回cos投递预览数据结果
func (c *Client) DescribeShipperPreview(request *DescribeShipperPreviewRequest) (response *DescribeShipperPreviewResponse, err error) {
	if request == nil {
		request = NewDescribeShipperPreviewRequest()
	}
	response = NewDescribeShipperPreviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConfigurationTemplateRequest() (request *ModifyConfigurationTemplateRequest) {
	request = &ModifyConfigurationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyConfigurationTemplate")
	return
}

func NewModifyConfigurationTemplateResponse() (response *ModifyConfigurationTemplateResponse) {
	response = &ModifyConfigurationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改配置模板
func (c *Client) ModifyConfigurationTemplate(request *ModifyConfigurationTemplateRequest) (response *ModifyConfigurationTemplateResponse, err error) {
	if request == nil {
		request = NewModifyConfigurationTemplateRequest()
	}
	response = NewModifyConfigurationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceAndFolderRelationRequest() (request *ModifyResourceAndFolderRelationRequest) {
	request = &ModifyResourceAndFolderRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyResourceAndFolderRelation")
	return
}

func NewModifyResourceAndFolderRelationResponse() (response *ModifyResourceAndFolderRelationResponse) {
	response = &ModifyResourceAndFolderRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于移动仪表盘到指定的文件夹
func (c *Client) ModifyResourceAndFolderRelation(request *ModifyResourceAndFolderRelationRequest) (response *ModifyResourceAndFolderRelationResponse, err error) {
	if request == nil {
		request = NewModifyResourceAndFolderRelationRequest()
	}
	response = NewModifyResourceAndFolderRelationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicsRequest() (request *DescribeTopicsRequest) {
	request = &DescribeTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeTopics")
	return
}

func NewDescribeTopicsResponse() (response *DescribeTopicsResponse) {
	response = &DescribeTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志主题列表，支持分页
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeTopicsRequest()
	}
	response = NewDescribeTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyIdleResourcePolicyRequest() (request *ModifyIdleResourcePolicyRequest) {
	request = &ModifyIdleResourcePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyIdleResourcePolicy")
	return
}

func NewModifyIdleResourcePolicyResponse() (response *ModifyIdleResourcePolicyResponse) {
	response = &ModifyIdleResourcePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改限制资源冻结策略
func (c *Client) ModifyIdleResourcePolicy(request *ModifyIdleResourcePolicyRequest) (response *ModifyIdleResourcePolicyResponse, err error) {
	if request == nil {
		request = NewModifyIdleResourcePolicyRequest()
	}
	response = NewModifyIdleResourcePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewSendConsumerHeartbeatRequest() (request *SendConsumerHeartbeatRequest) {
	request = &SendConsumerHeartbeatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "SendConsumerHeartbeat")
	return
}

func NewSendConsumerHeartbeatResponse() (response *SendConsumerHeartbeatResponse) {
	response = &SendConsumerHeartbeatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消费组心跳
func (c *Client) SendConsumerHeartbeat(request *SendConsumerHeartbeatRequest) (response *SendConsumerHeartbeatResponse, err error) {
	if request == nil {
		request = NewSendConsumerHeartbeatRequest()
	}
	response = NewSendConsumerHeartbeatResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformProcessInfoRequest() (request *DescribeDataTransformProcessInfoRequest) {
	request = &DescribeDataTransformProcessInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformProcessInfo")
	return
}

func NewDescribeDataTransformProcessInfoResponse() (response *DescribeDataTransformProcessInfoResponse) {
	response = &DescribeDataTransformProcessInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取数据加工任务执行进度详情
func (c *Client) DescribeDataTransformProcessInfo(request *DescribeDataTransformProcessInfoRequest) (response *DescribeDataTransformProcessInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformProcessInfoRequest()
	}
	response = NewDescribeDataTransformProcessInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLogsetRequest() (request *ModifyLogsetRequest) {
	request = &ModifyLogsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyLogset")
	return
}

func NewModifyLogsetResponse() (response *ModifyLogsetResponse) {
	response = &ModifyLogsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改日志集信息
func (c *Client) ModifyLogset(request *ModifyLogsetRequest) (response *ModifyLogsetResponse, err error) {
	if request == nil {
		request = NewModifyLogsetRequest()
	}
	response = NewModifyLogsetResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConsumerRequest() (request *DeleteConsumerRequest) {
	request = &DeleteConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConsumer")
	return
}

func NewDeleteConsumerResponse() (response *DeleteConsumerResponse) {
	response = &DeleteConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除投递配置
func (c *Client) DeleteConsumer(request *DeleteConsumerRequest) (response *DeleteConsumerResponse, err error) {
	if request == nil {
		request = NewDeleteConsumerRequest()
	}
	response = NewDeleteConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMachineGroupInfoRequest() (request *DeleteMachineGroupInfoRequest) {
	request = &DeleteMachineGroupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteMachineGroupInfo")
	return
}

func NewDeleteMachineGroupInfoResponse() (response *DeleteMachineGroupInfoResponse) {
	response = &DeleteMachineGroupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于删除机器组信息
func (c *Client) DeleteMachineGroupInfo(request *DeleteMachineGroupInfoRequest) (response *DeleteMachineGroupInfoResponse, err error) {
	if request == nil {
		request = NewDeleteMachineGroupInfoRequest()
	}
	response = NewDeleteMachineGroupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashboardSubscribeRequest() (request *DeleteDashboardSubscribeRequest) {
	request = &DeleteDashboardSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteDashboardSubscribe")
	return
}

func NewDeleteDashboardSubscribeResponse() (response *DeleteDashboardSubscribeResponse) {
	response = &DeleteDashboardSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于删除仪表盘订阅
func (c *Client) DeleteDashboardSubscribe(request *DeleteDashboardSubscribeRequest) (response *DeleteDashboardSubscribeResponse, err error) {
	if request == nil {
		request = NewDeleteDashboardSubscribeRequest()
	}
	response = NewDeleteDashboardSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyExternalDataSourceRequest() (request *ModifyExternalDataSourceRequest) {
	request = &ModifyExternalDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyExternalDataSource")
	return
}

func NewModifyExternalDataSourceResponse() (response *ModifyExternalDataSourceResponse) {
	response = &ModifyExternalDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改外部数据源
func (c *Client) ModifyExternalDataSource(request *ModifyExternalDataSourceRequest) (response *ModifyExternalDataSourceResponse, err error) {
	if request == nil {
		request = NewModifyExternalDataSourceRequest()
	}
	response = NewModifyExternalDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewCloseKafkaConsumeRequest() (request *CloseKafkaConsumeRequest) {
	request = &CloseKafkaConsumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CloseKafkaConsume")
	return
}

func NewCloseKafkaConsumeResponse() (response *CloseKafkaConsumeResponse) {
	response = &CloseKafkaConsumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭Kafka消费
func (c *Client) CloseKafkaConsume(request *CloseKafkaConsumeRequest) (response *CloseKafkaConsumeResponse, err error) {
	if request == nil {
		request = NewCloseKafkaConsumeRequest()
	}
	response = NewCloseKafkaConsumeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRecordingRuleYamlTaskRequest() (request *DeleteRecordingRuleYamlTaskRequest) {
	request = &DeleteRecordingRuleYamlTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteRecordingRuleYamlTask")
	return
}

func NewDeleteRecordingRuleYamlTaskResponse() (response *DeleteRecordingRuleYamlTaskResponse) {
	response = &DeleteRecordingRuleYamlTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除YAML预聚合任务
func (c *Client) DeleteRecordingRuleYamlTask(request *DeleteRecordingRuleYamlTaskRequest) (response *DeleteRecordingRuleYamlTaskResponse, err error) {
	if request == nil {
		request = NewDeleteRecordingRuleYamlTaskRequest()
	}
	response = NewDeleteRecordingRuleYamlTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMetricConfigRequest() (request *CreateMetricConfigRequest) {
	request = &CreateMetricConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateMetricConfig")
	return
}

func NewCreateMetricConfigResponse() (response *CreateMetricConfigResponse) {
	response = &CreateMetricConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建指标采集配置
func (c *Client) CreateMetricConfig(request *CreateMetricConfigRequest) (response *CreateMetricConfigResponse, err error) {
	if request == nil {
		request = NewCreateMetricConfigRequest()
	}
	response = NewCreateMetricConfigResponse()
	err = c.Send(request, response)
	return
}

func NewCancelRebuildIndexTaskRequest() (request *CancelRebuildIndexTaskRequest) {
	request = &CancelRebuildIndexTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CancelRebuildIndexTask")
	return
}

func NewCancelRebuildIndexTaskResponse() (response *CancelRebuildIndexTaskResponse) {
	response = &CancelRebuildIndexTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 取消重建索引任务
func (c *Client) CancelRebuildIndexTask(request *CancelRebuildIndexTaskRequest) (response *CancelRebuildIndexTaskResponse, err error) {
	if request == nil {
		request = NewCancelRebuildIndexTaskRequest()
	}
	response = NewCancelRebuildIndexTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBinlogSubscribeRequest() (request *DeleteBinlogSubscribeRequest) {
	request = &DeleteBinlogSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteBinlogSubscribe")
	return
}

func NewDeleteBinlogSubscribeResponse() (response *DeleteBinlogSubscribeResponse) {
	response = &DeleteBinlogSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除binlog采集配置
func (c *Client) DeleteBinlogSubscribe(request *DeleteBinlogSubscribeRequest) (response *DeleteBinlogSubscribeResponse, err error) {
	if request == nil {
		request = NewDeleteBinlogSubscribeRequest()
	}
	response = NewDeleteBinlogSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetricSubscribesRequest() (request *DescribeMetricSubscribesRequest) {
	request = &DescribeMetricSubscribesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMetricSubscribes")
	return
}

func NewDescribeMetricSubscribesResponse() (response *DescribeMetricSubscribesResponse) {
	response = &DescribeMetricSubscribesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指标订阅配置
func (c *Client) DescribeMetricSubscribes(request *DescribeMetricSubscribesRequest) (response *DescribeMetricSubscribesResponse, err error) {
	if request == nil {
		request = NewDescribeMetricSubscribesRequest()
	}
	response = NewDescribeMetricSubscribesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRemoteWriteTaskRequest() (request *DescribeRemoteWriteTaskRequest) {
	request = &DescribeRemoteWriteTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeRemoteWriteTask")
	return
}

func NewDescribeRemoteWriteTaskResponse() (response *DescribeRemoteWriteTaskResponse) {
	response = &DescribeRemoteWriteTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取RemoteWrite投递任务列表
func (c *Client) DescribeRemoteWriteTask(request *DescribeRemoteWriteTaskRequest) (response *DescribeRemoteWriteTaskResponse, err error) {
	if request == nil {
		request = NewDescribeRemoteWriteTaskRequest()
	}
	response = NewDescribeRemoteWriteTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCheckAlarmRuleRequest() (request *CheckAlarmRuleRequest) {
	request = &CheckAlarmRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckAlarmRule")
	return
}

func NewCheckAlarmRuleResponse() (response *CheckAlarmRuleResponse) {
	response = &CheckAlarmRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警策略检测接口
func (c *Client) CheckAlarmRule(request *CheckAlarmRuleRequest) (response *CheckAlarmRuleResponse, err error) {
	if request == nil {
		request = NewCheckAlarmRuleRequest()
	}
	response = NewCheckAlarmRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformPreviewInfoRequest() (request *DescribeDataTransformPreviewInfoRequest) {
	request = &DescribeDataTransformPreviewInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformPreviewInfo")
	return
}

func NewDescribeDataTransformPreviewInfoResponse() (response *DescribeDataTransformPreviewInfoResponse) {
	response = &DescribeDataTransformPreviewInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据加工预览任务基本信息
func (c *Client) DescribeDataTransformPreviewInfo(request *DescribeDataTransformPreviewInfoRequest) (response *DescribeDataTransformPreviewInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformPreviewInfoRequest()
	}
	response = NewDescribeDataTransformPreviewInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountRequest() (request *DescribeAccountRequest) {
	request = &DescribeAccountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAccount")
	return
}

func NewDescribeAccountResponse() (response *DescribeAccountResponse) {
	response = &DescribeAccountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取账户状态
func (c *Client) DescribeAccount(request *DescribeAccountRequest) (response *DescribeAccountResponse, err error) {
	if request == nil {
		request = NewDescribeAccountRequest()
	}
	response = NewDescribeAccountResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLogsetRequest() (request *DeleteLogsetRequest) {
	request = &DeleteLogsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteLogset")
	return
}

func NewDeleteLogsetResponse() (response *DeleteLogsetResponse) {
	response = &DeleteLogsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除日志集。
func (c *Client) DeleteLogset(request *DeleteLogsetRequest) (response *DeleteLogsetResponse, err error) {
	if request == nil {
		request = NewDeleteLogsetRequest()
	}
	response = NewDeleteLogsetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDashboardsRequest() (request *DescribeDashboardsRequest) {
	request = &DescribeDashboardsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDashboards")
	return
}

func NewDescribeDashboardsResponse() (response *DescribeDashboardsResponse) {
	response = &DescribeDashboardsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取仪表盘
func (c *Client) DescribeDashboards(request *DescribeDashboardsRequest) (response *DescribeDashboardsResponse, err error) {
	if request == nil {
		request = NewDescribeDashboardsRequest()
	}
	response = NewDescribeDashboardsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyConsumerRequest() (request *ModifyConsumerRequest) {
	request = &ModifyConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyConsumer")
	return
}

func NewModifyConsumerResponse() (response *ModifyConsumerResponse) {
	response = &ModifyConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改投递任务
func (c *Client) ModifyConsumer(request *ModifyConsumerRequest) (response *ModifyConsumerResponse, err error) {
	if request == nil {
		request = NewModifyConsumerRequest()
	}
	response = NewModifyConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewModifyShipperRequest() (request *ModifyShipperRequest) {
	request = &ModifyShipperRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyShipper")
	return
}

func NewModifyShipperResponse() (response *ModifyShipperResponse) {
	response = &ModifyShipperResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改现有的投递规则，客户如果使用此接口，需要自行处理CLS对指定bucket的写权限。
func (c *Client) ModifyShipper(request *ModifyShipperRequest) (response *ModifyShipperResponse, err error) {
	if request == nil {
		request = NewModifyShipperRequest()
	}
	response = NewModifyShipperResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLatestJsonLogRequest() (request *DescribeLatestJsonLogRequest) {
	request = &DescribeLatestJsonLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLatestJsonLog")
	return
}

func NewDescribeLatestJsonLogResponse() (response *DescribeLatestJsonLogResponse) {
	response = &DescribeLatestJsonLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取json格式日志
func (c *Client) DescribeLatestJsonLog(request *DescribeLatestJsonLogRequest) (response *DescribeLatestJsonLogResponse, err error) {
	if request == nil {
		request = NewDescribeLatestJsonLogRequest()
	}
	response = NewDescribeLatestJsonLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePartitionsRequest() (request *DescribePartitionsRequest) {
	request = &DescribePartitionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribePartitions")
	return
}

func NewDescribePartitionsResponse() (response *DescribePartitionsResponse) {
	response = &DescribePartitionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取主题分区列表。
func (c *Client) DescribePartitions(request *DescribePartitionsRequest) (response *DescribePartitionsResponse, err error) {
	if request == nil {
		request = NewDescribePartitionsRequest()
	}
	response = NewDescribePartitionsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateDashboardSubscribeRequest() (request *CreateDashboardSubscribeRequest) {
	request = &CreateDashboardSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateDashboardSubscribe")
	return
}

func NewCreateDashboardSubscribeResponse() (response *CreateDashboardSubscribeResponse) {
	response = &CreateDashboardSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于创建仪表盘订阅
func (c *Client) CreateDashboardSubscribe(request *CreateDashboardSubscribeRequest) (response *CreateDashboardSubscribeResponse, err error) {
	if request == nil {
		request = NewCreateDashboardSubscribeRequest()
	}
	response = NewCreateDashboardSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogFastAnalysisRequest() (request *DescribeLogFastAnalysisRequest) {
	request = &DescribeLogFastAnalysisRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLogFastAnalysis")
	return
}

func NewDescribeLogFastAnalysisResponse() (response *DescribeLogFastAnalysisResponse) {
	response = &DescribeLogFastAnalysisResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 针对日志某个字段可以做快速分析其占比情况
func (c *Client) DescribeLogFastAnalysis(request *DescribeLogFastAnalysisRequest) (response *DescribeLogFastAnalysisResponse, err error) {
	if request == nil {
		request = NewDescribeLogFastAnalysisRequest()
	}
	response = NewDescribeLogFastAnalysisResponse()
	err = c.Send(request, response)
	return
}

func NewCheckRechargeKafkaServerRequest() (request *CheckRechargeKafkaServerRequest) {
	request = &CheckRechargeKafkaServerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckRechargeKafkaServer")
	return
}

func NewCheckRechargeKafkaServerResponse() (response *CheckRechargeKafkaServerResponse) {
	response = &CheckRechargeKafkaServerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于校验Kafka服务集群是否可以正常访问
func (c *Client) CheckRechargeKafkaServer(request *CheckRechargeKafkaServerRequest) (response *CheckRechargeKafkaServerResponse, err error) {
	if request == nil {
		request = NewCheckRechargeKafkaServerRequest()
	}
	response = NewCheckRechargeKafkaServerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogsetsRequest() (request *DescribeLogsetsRequest) {
	request = &DescribeLogsetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLogsets")
	return
}

func NewDescribeLogsetsResponse() (response *DescribeLogsetsResponse) {
	response = &DescribeLogsetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取日志集信息列表。
func (c *Client) DescribeLogsets(request *DescribeLogsetsRequest) (response *DescribeLogsetsResponse, err error) {
	if request == nil {
		request = NewDescribeLogsetsRequest()
	}
	response = NewDescribeLogsetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeShipperTasksRequest() (request *DescribeShipperTasksRequest) {
	request = &DescribeShipperTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeShipperTasks")
	return
}

func NewDescribeShipperTasksResponse() (response *DescribeShipperTasksResponse) {
	response = &DescribeShipperTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取投递任务列表
func (c *Client) DescribeShipperTasks(request *DescribeShipperTasksRequest) (response *DescribeShipperTasksResponse, err error) {
	if request == nil {
		request = NewDescribeShipperTasksRequest()
	}
	response = NewDescribeShipperTasksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmNoticesRequest() (request *DescribeAlarmNoticesRequest) {
	request = &DescribeAlarmNoticesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarmNotices")
	return
}

func NewDescribeAlarmNoticesResponse() (response *DescribeAlarmNoticesResponse) {
	response = &DescribeAlarmNoticesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于获取告警通知模板列表
func (c *Client) DescribeAlarmNotices(request *DescribeAlarmNoticesRequest) (response *DescribeAlarmNoticesResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmNoticesRequest()
	}
	response = NewDescribeAlarmNoticesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigurationTemplateRequest() (request *CreateConfigurationTemplateRequest) {
	request = &CreateConfigurationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateConfigurationTemplate")
	return
}

func NewCreateConfigurationTemplateResponse() (response *CreateConfigurationTemplateResponse) {
	response = &CreateConfigurationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建配置模板
func (c *Client) CreateConfigurationTemplate(request *CreateConfigurationTemplateRequest) (response *CreateConfigurationTemplateResponse, err error) {
	if request == nil {
		request = NewCreateConfigurationTemplateRequest()
	}
	response = NewCreateConfigurationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIdleResourcePolicyRequest() (request *DescribeIdleResourcePolicyRequest) {
	request = &DescribeIdleResourcePolicyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeIdleResourcePolicy")
	return
}

func NewDescribeIdleResourcePolicyResponse() (response *DescribeIdleResourcePolicyResponse) {
	response = &DescribeIdleResourcePolicyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取限制资源策略
func (c *Client) DescribeIdleResourcePolicy(request *DescribeIdleResourcePolicyRequest) (response *DescribeIdleResourcePolicyResponse, err error) {
	if request == nil {
		request = NewDescribeIdleResourcePolicyRequest()
	}
	response = NewDescribeIdleResourcePolicyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentMachineGroupMetadataRequest() (request *DescribeAgentMachineGroupMetadataRequest) {
	request = &DescribeAgentMachineGroupMetadataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAgentMachineGroupMetadata")
	return
}

func NewDescribeAgentMachineGroupMetadataResponse() (response *DescribeAgentMachineGroupMetadataResponse) {
	response = &DescribeAgentMachineGroupMetadataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于Agent获取机器组相关元数据信息
func (c *Client) DescribeAgentMachineGroupMetadata(request *DescribeAgentMachineGroupMetadataRequest) (response *DescribeAgentMachineGroupMetadataResponse, err error) {
	if request == nil {
		request = NewDescribeAgentMachineGroupMetadataRequest()
	}
	response = NewDescribeAgentMachineGroupMetadataResponse()
	err = c.Send(request, response)
	return
}

func NewMergePartitionRequest() (request *MergePartitionRequest) {
	request = &MergePartitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "MergePartition")
	return
}

func NewMergePartitionResponse() (response *MergePartitionResponse) {
	response = &MergePartitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于合并一个读写态的主题分区，合并时指定一个主题分区 ID，日志服务会自动合并范围右相邻的分区。
func (c *Client) MergePartition(request *MergePartitionRequest) (response *MergePartitionResponse, err error) {
	if request == nil {
		request = NewMergePartitionRequest()
	}
	response = NewMergePartitionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaConsumerPreviewRequest() (request *DescribeKafkaConsumerPreviewRequest) {
	request = &DescribeKafkaConsumerPreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumerPreview")
	return
}

func NewDescribeKafkaConsumerPreviewResponse() (response *DescribeKafkaConsumerPreviewResponse) {
	response = &DescribeKafkaConsumerPreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// kafka协议消费预览接口
func (c *Client) DescribeKafkaConsumerPreview(request *DescribeKafkaConsumerPreviewRequest) (response *DescribeKafkaConsumerPreviewResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaConsumerPreviewRequest()
	}
	response = NewDescribeKafkaConsumerPreviewResponse()
	err = c.Send(request, response)
	return
}

func NewCommitConsumerOffsetsRequest() (request *CommitConsumerOffsetsRequest) {
	request = &CommitConsumerOffsetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CommitConsumerOffsets")
	return
}

func NewCommitConsumerOffsetsResponse() (response *CommitConsumerOffsetsResponse) {
	response = &CommitConsumerOffsetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提交消费点位
func (c *Client) CommitConsumerOffsets(request *CommitConsumerOffsetsRequest) (response *CommitConsumerOffsetsResponse, err error) {
	if request == nil {
		request = NewCommitConsumerOffsetsRequest()
	}
	response = NewCommitConsumerOffsetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformInfoRequest() (request *DescribeDataTransformInfoRequest) {
	request = &DescribeDataTransformInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformInfo")
	return
}

func NewDescribeDataTransformInfoResponse() (response *DescribeDataTransformInfoResponse) {
	response = &DescribeDataTransformInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取数据加工任务列表基本信息
func (c *Client) DescribeDataTransformInfo(request *DescribeDataTransformInfoRequest) (response *DescribeDataTransformInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformInfoRequest()
	}
	response = NewDescribeDataTransformInfoResponse()
	err = c.Send(request, response)
	return
}

func NewCloseKafkaConsumerRequest() (request *CloseKafkaConsumerRequest) {
	request = &CloseKafkaConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CloseKafkaConsumer")
	return
}

func NewCloseKafkaConsumerResponse() (response *CloseKafkaConsumerResponse) {
	response = &CloseKafkaConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关闭Kafka协议消费
func (c *Client) CloseKafkaConsumer(request *CloseKafkaConsumerRequest) (response *CloseKafkaConsumerResponse, err error) {
	if request == nil {
		request = NewCloseKafkaConsumerRequest()
	}
	response = NewCloseKafkaConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewCreateKafkaRechargeRequest() (request *CreateKafkaRechargeRequest) {
	request = &CreateKafkaRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateKafkaRecharge")
	return
}

func NewCreateKafkaRechargeResponse() (response *CreateKafkaRechargeResponse) {
	response = &CreateKafkaRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建Kafka数据订阅任务
func (c *Client) CreateKafkaRecharge(request *CreateKafkaRechargeRequest) (response *CreateKafkaRechargeResponse, err error) {
	if request == nil {
		request = NewCreateKafkaRechargeRequest()
	}
	response = NewCreateKafkaRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRecordingRuleYamlTaskRequest() (request *ModifyRecordingRuleYamlTaskRequest) {
	request = &ModifyRecordingRuleYamlTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyRecordingRuleYamlTask")
	return
}

func NewModifyRecordingRuleYamlTaskResponse() (response *ModifyRecordingRuleYamlTaskResponse) {
	response = &ModifyRecordingRuleYamlTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改YAML创建预聚合任务
func (c *Client) ModifyRecordingRuleYamlTask(request *ModifyRecordingRuleYamlTaskRequest) (response *ModifyRecordingRuleYamlTaskResponse, err error) {
	if request == nil {
		request = NewModifyRecordingRuleYamlTaskRequest()
	}
	response = NewModifyRecordingRuleYamlTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCheckFunctionRequest() (request *CheckFunctionRequest) {
	request = &CheckFunctionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckFunction")
	return
}

func NewCheckFunctionResponse() (response *CheckFunctionResponse) {
	response = &CheckFunctionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于语法校验接口。
func (c *Client) CheckFunction(request *CheckFunctionRequest) (response *CheckFunctionResponse, err error) {
	if request == nil {
		request = NewCheckFunctionRequest()
	}
	response = NewCheckFunctionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaConsumerTopicsRequest() (request *DescribeKafkaConsumerTopicsRequest) {
	request = &DescribeKafkaConsumerTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumerTopics")
	return
}

func NewDescribeKafkaConsumerTopicsResponse() (response *DescribeKafkaConsumerTopicsResponse) {
	response = &DescribeKafkaConsumerTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取kafka协议消费主题信息列表
func (c *Client) DescribeKafkaConsumerTopics(request *DescribeKafkaConsumerTopicsRequest) (response *DescribeKafkaConsumerTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaConsumerTopicsRequest()
	}
	response = NewDescribeKafkaConsumerTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewEstimateRebuildIndexTaskRequest() (request *EstimateRebuildIndexTaskRequest) {
	request = &EstimateRebuildIndexTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "EstimateRebuildIndexTask")
	return
}

func NewEstimateRebuildIndexTaskResponse() (response *EstimateRebuildIndexTaskResponse) {
	response = &EstimateRebuildIndexTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 预估重建索引任务
func (c *Client) EstimateRebuildIndexTask(request *EstimateRebuildIndexTaskRequest) (response *EstimateRebuildIndexTaskResponse, err error) {
	if request == nil {
		request = NewEstimateRebuildIndexTaskRequest()
	}
	response = NewEstimateRebuildIndexTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateWebCallbackRequest() (request *CreateWebCallbackRequest) {
	request = &CreateWebCallbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateWebCallback")
	return
}

func NewCreateWebCallbackResponse() (response *CreateWebCallbackResponse) {
	response = &CreateWebCallbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建告警渠道回调配置
func (c *Client) CreateWebCallback(request *CreateWebCallbackRequest) (response *CreateWebCallbackResponse, err error) {
	if request == nil {
		request = NewCreateWebCallbackRequest()
	}
	response = NewCreateWebCallbackResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigsRequest() (request *DescribeConfigsRequest) {
	request = &DescribeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigs")
	return
}

func NewDescribeConfigsResponse() (response *DescribeConfigsResponse) {
	response = &DescribeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集规则配置
func (c *Client) DescribeConfigs(request *DescribeConfigsRequest) (response *DescribeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeConfigsRequest()
	}
	response = NewDescribeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRecordingRuleTaskRequest() (request *DescribeRecordingRuleTaskRequest) {
	request = &DescribeRecordingRuleTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeRecordingRuleTask")
	return
}

func NewDescribeRecordingRuleTaskResponse() (response *DescribeRecordingRuleTaskResponse) {
	response = &DescribeRecordingRuleTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取预聚合任务列表
func (c *Client) DescribeRecordingRuleTask(request *DescribeRecordingRuleTaskRequest) (response *DescribeRecordingRuleTaskResponse, err error) {
	if request == nil {
		request = NewDescribeRecordingRuleTaskRequest()
	}
	response = NewDescribeRecordingRuleTaskResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRecordingRuleTaskRequest() (request *CreateRecordingRuleTaskRequest) {
	request = &CreateRecordingRuleTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateRecordingRuleTask")
	return
}

func NewCreateRecordingRuleTaskResponse() (response *CreateRecordingRuleTaskResponse) {
	response = &CreateRecordingRuleTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建预聚合任务
func (c *Client) CreateRecordingRuleTask(request *CreateRecordingRuleTaskRequest) (response *CreateRecordingRuleTaskResponse, err error) {
	if request == nil {
		request = NewCreateRecordingRuleTaskRequest()
	}
	response = NewCreateRecordingRuleTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMetricSubscribeRequest() (request *DeleteMetricSubscribeRequest) {
	request = &DeleteMetricSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteMetricSubscribe")
	return
}

func NewDeleteMetricSubscribeResponse() (response *DeleteMetricSubscribeResponse) {
	response = &DeleteMetricSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指标订阅配置
func (c *Client) DeleteMetricSubscribe(request *DeleteMetricSubscribeRequest) (response *DeleteMetricSubscribeResponse, err error) {
	if request == nil {
		request = NewDeleteMetricSubscribeRequest()
	}
	response = NewDeleteMetricSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlertRecordHistoryRequest() (request *DescribeAlertRecordHistoryRequest) {
	request = &DescribeAlertRecordHistoryRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAlertRecordHistory")
	return
}

func NewDescribeAlertRecordHistoryResponse() (response *DescribeAlertRecordHistoryResponse) {
	response = &DescribeAlertRecordHistoryResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 告警历史记录
func (c *Client) DescribeAlertRecordHistory(request *DescribeAlertRecordHistoryRequest) (response *DescribeAlertRecordHistoryResponse, err error) {
	if request == nil {
		request = NewDescribeAlertRecordHistoryRequest()
	}
	response = NewDescribeAlertRecordHistoryResponse()
	err = c.Send(request, response)
	return
}

func NewOpenKafkaConsumerRequest() (request *OpenKafkaConsumerRequest) {
	request = &OpenKafkaConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "OpenKafkaConsumer")
	return
}

func NewOpenKafkaConsumerResponse() (response *OpenKafkaConsumerResponse) {
	response = &OpenKafkaConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 打开Kafka协议消费功能
func (c *Client) OpenKafkaConsumer(request *OpenKafkaConsumerRequest) (response *OpenKafkaConsumerResponse, err error) {
	if request == nil {
		request = NewOpenKafkaConsumerRequest()
	}
	response = NewOpenKafkaConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExportRequest() (request *DeleteExportRequest) {
	request = &DeleteExportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteExport")
	return
}

func NewDeleteExportResponse() (response *DeleteExportResponse) {
	response = &DeleteExportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除日志下载任务
func (c *Client) DeleteExport(request *DeleteExportRequest) (response *DeleteExportResponse, err error) {
	if request == nil {
		request = NewDeleteExportRequest()
	}
	response = NewDeleteExportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineGroupConfigsRequest() (request *DescribeMachineGroupConfigsRequest) {
	request = &DescribeMachineGroupConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMachineGroupConfigs")
	return
}

func NewDescribeMachineGroupConfigsResponse() (response *DescribeMachineGroupConfigsResponse) {
	response = &DescribeMachineGroupConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取机器组绑定的采集规则配置
func (c *Client) DescribeMachineGroupConfigs(request *DescribeMachineGroupConfigsRequest) (response *DescribeMachineGroupConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeMachineGroupConfigsRequest()
	}
	response = NewDescribeMachineGroupConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeExternalDataSourcePreviewRequest() (request *DescribeExternalDataSourcePreviewRequest) {
	request = &DescribeExternalDataSourcePreviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeExternalDataSourcePreview")
	return
}

func NewDescribeExternalDataSourcePreviewResponse() (response *DescribeExternalDataSourcePreviewResponse) {
	response = &DescribeExternalDataSourcePreviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 外部数据源预览（验证连通性）
func (c *Client) DescribeExternalDataSourcePreview(request *DescribeExternalDataSourcePreviewRequest) (response *DescribeExternalDataSourcePreviewResponse, err error) {
	if request == nil {
		request = NewDescribeExternalDataSourcePreviewRequest()
	}
	response = NewDescribeExternalDataSourcePreviewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFunctionsRequest() (request *DescribeFunctionsRequest) {
	request = &DescribeFunctionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeFunctions")
	return
}

func NewDescribeFunctionsResponse() (response *DescribeFunctionsResponse) {
	response = &DescribeFunctionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取可使用的函数列表。
func (c *Client) DescribeFunctions(request *DescribeFunctionsRequest) (response *DescribeFunctionsResponse, err error) {
	if request == nil {
		request = NewDescribeFunctionsRequest()
	}
	response = NewDescribeFunctionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogHistogramRequest() (request *DescribeLogHistogramRequest) {
	request = &DescribeLogHistogramRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLogHistogram")
	return
}

func NewDescribeLogHistogramResponse() (response *DescribeLogHistogramResponse) {
	response = &DescribeLogHistogramResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于构建直方图
func (c *Client) DescribeLogHistogram(request *DescribeLogHistogramRequest) (response *DescribeLogHistogramResponse, err error) {
	if request == nil {
		request = NewDescribeLogHistogramRequest()
	}
	response = NewDescribeLogHistogramResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformPreviewDataInfoRequest() (request *DescribeDataTransformPreviewDataInfoRequest) {
	request = &DescribeDataTransformPreviewDataInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformPreviewDataInfo")
	return
}

func NewDescribeDataTransformPreviewDataInfoResponse() (response *DescribeDataTransformPreviewDataInfoResponse) {
	response = &DescribeDataTransformPreviewDataInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取数据加工预览详细数据。
func (c *Client) DescribeDataTransformPreviewDataInfo(request *DescribeDataTransformPreviewDataInfoRequest) (response *DescribeDataTransformPreviewDataInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformPreviewDataInfoRequest()
	}
	response = NewDescribeDataTransformPreviewDataInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteWebCallbackRequest() (request *DeleteWebCallbackRequest) {
	request = &DeleteWebCallbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteWebCallback")
	return
}

func NewDeleteWebCallbackResponse() (response *DeleteWebCallbackResponse) {
	response = &DeleteWebCallbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除告警渠道回调配置
func (c *Client) DeleteWebCallback(request *DeleteWebCallbackRequest) (response *DeleteWebCallbackResponse, err error) {
	if request == nil {
		request = NewDeleteWebCallbackRequest()
	}
	response = NewDeleteWebCallbackResponse()
	err = c.Send(request, response)
	return
}

func NewRetryShipperTaskRequest() (request *RetryShipperTaskRequest) {
	request = &RetryShipperTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "RetryShipperTask")
	return
}

func NewRetryShipperTaskResponse() (response *RetryShipperTaskResponse) {
	response = &RetryShipperTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试失败的投递任务
func (c *Client) RetryShipperTask(request *RetryShipperTaskRequest) (response *RetryShipperTaskResponse, err error) {
	if request == nil {
		request = NewRetryShipperTaskRequest()
	}
	response = NewRetryShipperTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAlarmDetailRequest() (request *DescribeAlarmDetailRequest) {
	request = &DescribeAlarmDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAlarmDetail")
	return
}

func NewDescribeAlarmDetailResponse() (response *DescribeAlarmDetailResponse) {
	response = &DescribeAlarmDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警信息快照日志
func (c *Client) DescribeAlarmDetail(request *DescribeAlarmDetailRequest) (response *DescribeAlarmDetailResponse, err error) {
	if request == nil {
		request = NewDescribeAlarmDetailRequest()
	}
	response = NewDescribeAlarmDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBinlogSubscribeConnectivityRequest() (request *DescribeBinlogSubscribeConnectivityRequest) {
	request = &DescribeBinlogSubscribeConnectivityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeBinlogSubscribeConnectivity")
	return
}

func NewDescribeBinlogSubscribeConnectivityResponse() (response *DescribeBinlogSubscribeConnectivityResponse) {
	response = &DescribeBinlogSubscribeConnectivityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// binlog采集配置连通性校验
func (c *Client) DescribeBinlogSubscribeConnectivity(request *DescribeBinlogSubscribeConnectivityRequest) (response *DescribeBinlogSubscribeConnectivityResponse, err error) {
	if request == nil {
		request = NewDescribeBinlogSubscribeConnectivityRequest()
	}
	response = NewDescribeBinlogSubscribeConnectivityResponse()
	err = c.Send(request, response)
	return
}

func NewCreateExternalDataSourceRequest() (request *CreateExternalDataSourceRequest) {
	request = &CreateExternalDataSourceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateExternalDataSource")
	return
}

func NewCreateExternalDataSourceResponse() (response *CreateExternalDataSourceResponse) {
	response = &CreateExternalDataSourceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建外部数据源。单个日志主题最多可关联20个外部数据（不包含其它日志主题分享到该日志主题的外部数据）
func (c *Client) CreateExternalDataSource(request *CreateExternalDataSourceRequest) (response *CreateExternalDataSourceResponse, err error) {
	if request == nil {
		request = NewCreateExternalDataSourceRequest()
	}
	response = NewCreateExternalDataSourceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRecordingRuleYamlTaskRequest() (request *CreateRecordingRuleYamlTaskRequest) {
	request = &CreateRecordingRuleYamlTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateRecordingRuleYamlTask")
	return
}

func NewCreateRecordingRuleYamlTaskResponse() (response *CreateRecordingRuleYamlTaskResponse) {
	response = &CreateRecordingRuleYamlTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过YAML批量创建预聚合任务
func (c *Client) CreateRecordingRuleYamlTask(request *CreateRecordingRuleYamlTaskRequest) (response *CreateRecordingRuleYamlTaskResponse, err error) {
	if request == nil {
		request = NewCreateRecordingRuleYamlTaskRequest()
	}
	response = NewCreateRecordingRuleYamlTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyNoticeContentRequest() (request *ModifyNoticeContentRequest) {
	request = &ModifyNoticeContentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyNoticeContent")
	return
}

func NewModifyNoticeContentResponse() (response *ModifyNoticeContentResponse) {
	response = &ModifyNoticeContentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改通知内容模板
func (c *Client) ModifyNoticeContent(request *ModifyNoticeContentRequest) (response *ModifyNoticeContentResponse, err error) {
	if request == nil {
		request = NewModifyNoticeContentRequest()
	}
	response = NewModifyNoticeContentResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopicExtendConfigRequest() (request *DeleteTopicExtendConfigRequest) {
	request = &DeleteTopicExtendConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteTopicExtendConfig")
	return
}

func NewDeleteTopicExtendConfigResponse() (response *DeleteTopicExtendConfigResponse) {
	response = &DeleteTopicExtendConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除采集配置(clb专用)。
func (c *Client) DeleteTopicExtendConfig(request *DeleteTopicExtendConfigRequest) (response *DeleteTopicExtendConfigResponse, err error) {
	if request == nil {
		request = NewDeleteTopicExtendConfigRequest()
	}
	response = NewDeleteTopicExtendConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicExtendConfigRequest() (request *DescribeTopicExtendConfigRequest) {
	request = &DescribeTopicExtendConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeTopicExtendConfig")
	return
}

func NewDescribeTopicExtendConfigResponse() (response *DescribeTopicExtendConfigResponse) {
	response = &DescribeTopicExtendConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取采集配置(clb专用)。
func (c *Client) DescribeTopicExtendConfig(request *DescribeTopicExtendConfigRequest) (response *DescribeTopicExtendConfigResponse, err error) {
	if request == nil {
		request = NewDescribeTopicExtendConfigRequest()
	}
	response = NewDescribeTopicExtendConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDashboardSubscribeAckRequest() (request *ModifyDashboardSubscribeAckRequest) {
	request = &ModifyDashboardSubscribeAckRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyDashboardSubscribeAck")
	return
}

func NewModifyDashboardSubscribeAckResponse() (response *ModifyDashboardSubscribeAckResponse) {
	response = &ModifyDashboardSubscribeAckResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于确认仪表盘订阅发送成功
func (c *Client) ModifyDashboardSubscribeAck(request *ModifyDashboardSubscribeAckRequest) (response *ModifyDashboardSubscribeAckResponse, err error) {
	if request == nil {
		request = NewModifyDashboardSubscribeAckRequest()
	}
	response = NewModifyDashboardSubscribeAckResponse()
	err = c.Send(request, response)
	return
}

func NewModifyFolderRequest() (request *ModifyFolderRequest) {
	request = &ModifyFolderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyFolder")
	return
}

func NewModifyFolderResponse() (response *ModifyFolderResponse) {
	response = &ModifyFolderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改文件夹
func (c *Client) ModifyFolder(request *ModifyFolderRequest) (response *ModifyFolderResponse, err error) {
	if request == nil {
		request = NewModifyFolderRequest()
	}
	response = NewModifyFolderResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsumerRequest() (request *DescribeConsumerRequest) {
	request = &DescribeConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumer")
	return
}

func NewDescribeConsumerResponse() (response *DescribeConsumerResponse) {
	response = &DescribeConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取投递配置
func (c *Client) DescribeConsumer(request *DescribeConsumerRequest) (response *DescribeConsumerResponse, err error) {
	if request == nil {
		request = NewDescribeConsumerRequest()
	}
	response = NewDescribeConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformFailLogInfoRequest() (request *DescribeDataTransformFailLogInfoRequest) {
	request = &DescribeDataTransformFailLogInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformFailLogInfo")
	return
}

func NewDescribeDataTransformFailLogInfoResponse() (response *DescribeDataTransformFailLogInfoResponse) {
	response = &DescribeDataTransformFailLogInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取数据加工任务失败日志详请
func (c *Client) DescribeDataTransformFailLogInfo(request *DescribeDataTransformFailLogInfoRequest) (response *DescribeDataTransformFailLogInfoResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformFailLogInfoRequest()
	}
	response = NewDescribeDataTransformFailLogInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLogContextRequest() (request *DescribeLogContextRequest) {
	request = &DescribeLogContextRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLogContext")
	return
}

func NewDescribeLogContextResponse() (response *DescribeLogContextResponse) {
	response = &DescribeLogContextResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于搜索日志上下文附近的内容
func (c *Client) DescribeLogContext(request *DescribeLogContextRequest) (response *DescribeLogContextResponse, err error) {
	if request == nil {
		request = NewDescribeLogContextRequest()
	}
	response = NewDescribeLogContextResponse()
	err = c.Send(request, response)
	return
}

func NewCheckConfigRegexRequest() (request *CheckConfigRegexRequest) {
	request = &CheckConfigRegexRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckConfigRegex")
	return
}

func NewCheckConfigRegexResponse() (response *CheckConfigRegexResponse) {
	response = &CheckConfigRegexResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 校验采集配置正则
func (c *Client) CheckConfigRegex(request *CheckConfigRegexRequest) (response *CheckConfigRegexResponse, err error) {
	if request == nil {
		request = NewCheckConfigRegexRequest()
	}
	response = NewCheckConfigRegexResponse()
	err = c.Send(request, response)
	return
}

func NewApplyConfigToMachineGroupRequest() (request *ApplyConfigToMachineGroupRequest) {
	request = &ApplyConfigToMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ApplyConfigToMachineGroup")
	return
}

func NewApplyConfigToMachineGroupResponse() (response *ApplyConfigToMachineGroupResponse) {
	response = &ApplyConfigToMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用采集配置到指定机器组
func (c *Client) ApplyConfigToMachineGroup(request *ApplyConfigToMachineGroupRequest) (response *ApplyConfigToMachineGroupResponse, err error) {
	if request == nil {
		request = NewApplyConfigToMachineGroupRequest()
	}
	response = NewApplyConfigToMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLogsetRequest() (request *CreateLogsetRequest) {
	request = &CreateLogsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateLogset")
	return
}

func NewCreateLogsetResponse() (response *CreateLogsetResponse) {
	response = &CreateLogsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建日志集，返回新创建的日志集的 ID。
func (c *Client) CreateLogset(request *CreateLogsetRequest) (response *CreateLogsetResponse, err error) {
	if request == nil {
		request = NewCreateLogsetRequest()
	}
	response = NewCreateLogsetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIndexsRequest() (request *DescribeIndexsRequest) {
	request = &DescribeIndexsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeIndexs")
	return
}

func NewDescribeIndexsResponse() (response *DescribeIndexsResponse) {
	response = &DescribeIndexsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取索引配置列表
func (c *Client) DescribeIndexs(request *DescribeIndexsRequest) (response *DescribeIndexsResponse, err error) {
	if request == nil {
		request = NewDescribeIndexsRequest()
	}
	response = NewDescribeIndexsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachinesRequest() (request *DescribeMachinesRequest) {
	request = &DescribeMachinesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeMachines")
	return
}

func NewDescribeMachinesResponse() (response *DescribeMachinesResponse) {
	response = &DescribeMachinesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取制定机器组下的机器状态
func (c *Client) DescribeMachines(request *DescribeMachinesRequest) (response *DescribeMachinesResponse, err error) {
	if request == nil {
		request = NewDescribeMachinesRequest()
	}
	response = NewDescribeMachinesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlarmNoticeRequest() (request *CreateAlarmNoticeRequest) {
	request = &CreateAlarmNoticeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateAlarmNotice")
	return
}

func NewCreateAlarmNoticeResponse() (response *CreateAlarmNoticeResponse) {
	response = &CreateAlarmNoticeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用户创建告警通知模板。
func (c *Client) CreateAlarmNotice(request *CreateAlarmNoticeRequest) (response *CreateAlarmNoticeResponse, err error) {
	if request == nil {
		request = NewCreateAlarmNoticeRequest()
	}
	response = NewCreateAlarmNoticeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScheduledSqlProcessInfoRequest() (request *DescribeScheduledSqlProcessInfoRequest) {
	request = &DescribeScheduledSqlProcessInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeScheduledSqlProcessInfo")
	return
}

func NewDescribeScheduledSqlProcessInfoResponse() (response *DescribeScheduledSqlProcessInfoResponse) {
	response = &DescribeScheduledSqlProcessInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取定时SQL分析任务进度信息
func (c *Client) DescribeScheduledSqlProcessInfo(request *DescribeScheduledSqlProcessInfoRequest) (response *DescribeScheduledSqlProcessInfoResponse, err error) {
	if request == nil {
		request = NewDescribeScheduledSqlProcessInfoRequest()
	}
	response = NewDescribeScheduledSqlProcessInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicBaseMetricConfigsRequest() (request *DescribeTopicBaseMetricConfigsRequest) {
	request = &DescribeTopicBaseMetricConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeTopicBaseMetricConfigs")
	return
}

func NewDescribeTopicBaseMetricConfigsResponse() (response *DescribeTopicBaseMetricConfigsResponse) {
	response = &DescribeTopicBaseMetricConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指标基础监控采集配置
func (c *Client) DescribeTopicBaseMetricConfigs(request *DescribeTopicBaseMetricConfigsRequest) (response *DescribeTopicBaseMetricConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeTopicBaseMetricConfigsRequest()
	}
	response = NewDescribeTopicBaseMetricConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyBinlogSubscribeRequest() (request *ModifyBinlogSubscribeRequest) {
	request = &ModifyBinlogSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyBinlogSubscribe")
	return
}

func NewModifyBinlogSubscribeResponse() (response *ModifyBinlogSubscribeResponse) {
	response = &ModifyBinlogSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改binlog采集配置
func (c *Client) ModifyBinlogSubscribe(request *ModifyBinlogSubscribeRequest) (response *ModifyBinlogSubscribeResponse, err error) {
	if request == nil {
		request = NewModifyBinlogSubscribeRequest()
	}
	response = NewModifyBinlogSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigurationTemplatesRequest() (request *DescribeConfigurationTemplatesRequest) {
	request = &DescribeConfigurationTemplatesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConfigurationTemplates")
	return
}

func NewDescribeConfigurationTemplatesResponse() (response *DescribeConfigurationTemplatesResponse) {
	response = &DescribeConfigurationTemplatesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置模板列表
func (c *Client) DescribeConfigurationTemplates(request *DescribeConfigurationTemplatesRequest) (response *DescribeConfigurationTemplatesResponse, err error) {
	if request == nil {
		request = NewDescribeConfigurationTemplatesRequest()
	}
	response = NewDescribeConfigurationTemplatesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteFolderRequest() (request *DeleteFolderRequest) {
	request = &DeleteFolderRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteFolder")
	return
}

func NewDeleteFolderResponse() (response *DeleteFolderResponse) {
	response = &DeleteFolderResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除文件夹
func (c *Client) DeleteFolder(request *DeleteFolderRequest) (response *DeleteFolderResponse, err error) {
	if request == nil {
		request = NewDeleteFolderRequest()
	}
	response = NewDeleteFolderResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigRequest() (request *CreateConfigRequest) {
	request = &CreateConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateConfig")
	return
}

func NewCreateConfigResponse() (response *CreateConfigResponse) {
	response = &CreateConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建采集规则配置
func (c *Client) CreateConfig(request *CreateConfigRequest) (response *CreateConfigResponse, err error) {
	if request == nil {
		request = NewCreateConfigRequest()
	}
	response = NewCreateConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAlarmShieldRequest() (request *DeleteAlarmShieldRequest) {
	request = &DeleteAlarmShieldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteAlarmShield")
	return
}

func NewDeleteAlarmShieldResponse() (response *DeleteAlarmShieldResponse) {
	response = &DeleteAlarmShieldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除告警屏蔽规则
func (c *Client) DeleteAlarmShield(request *DeleteAlarmShieldRequest) (response *DeleteAlarmShieldResponse, err error) {
	if request == nil {
		request = NewDeleteAlarmShieldRequest()
	}
	response = NewDeleteAlarmShieldResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDataTransformRequest() (request *DeleteDataTransformRequest) {
	request = &DeleteDataTransformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteDataTransform")
	return
}

func NewDeleteDataTransformResponse() (response *DeleteDataTransformResponse) {
	response = &DeleteDataTransformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除数据加工任务
func (c *Client) DeleteDataTransform(request *DeleteDataTransformRequest) (response *DeleteDataTransformResponse, err error) {
	if request == nil {
		request = NewDeleteDataTransformRequest()
	}
	response = NewDeleteDataTransformResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaRechargesRequest() (request *DescribeKafkaRechargesRequest) {
	request = &DescribeKafkaRechargesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaRecharges")
	return
}

func NewDescribeKafkaRechargesResponse() (response *DescribeKafkaRechargesResponse) {
	response = &DescribeKafkaRechargesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取Kafka数据订阅任务
func (c *Client) DescribeKafkaRecharges(request *DescribeKafkaRechargesRequest) (response *DescribeKafkaRechargesResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaRechargesRequest()
	}
	response = NewDescribeKafkaRechargesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAccountInfoRequest() (request *DescribeAccountInfoRequest) {
	request = &DescribeAccountInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeAccountInfo")
	return
}

func NewDescribeAccountInfoResponse() (response *DescribeAccountInfoResponse) {
	response = &DescribeAccountInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取用户信息
func (c *Client) DescribeAccountInfo(request *DescribeAccountInfoRequest) (response *DescribeAccountInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAccountInfoRequest()
	}
	response = NewDescribeAccountInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDashboardSubscribeRequest() (request *ModifyDashboardSubscribeRequest) {
	request = &ModifyDashboardSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyDashboardSubscribe")
	return
}

func NewModifyDashboardSubscribeResponse() (response *ModifyDashboardSubscribeResponse) {
	response = &ModifyDashboardSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于修改仪表盘订阅
func (c *Client) ModifyDashboardSubscribe(request *ModifyDashboardSubscribeRequest) (response *ModifyDashboardSubscribeResponse, err error) {
	if request == nil {
		request = NewModifyDashboardSubscribeRequest()
	}
	response = NewModifyDashboardSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCosRechargeRequest() (request *ModifyCosRechargeRequest) {
	request = &ModifyCosRechargeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyCosRecharge")
	return
}

func NewModifyCosRechargeResponse() (response *ModifyCosRechargeResponse) {
	response = &ModifyCosRechargeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改对象存储导入任务
func (c *Client) ModifyCosRecharge(request *ModifyCosRechargeRequest) (response *ModifyCosRechargeResponse, err error) {
	if request == nil {
		request = NewModifyCosRechargeRequest()
	}
	response = NewModifyCosRechargeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaConsumerRequest() (request *DescribeKafkaConsumerRequest) {
	request = &DescribeKafkaConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsumer")
	return
}

func NewDescribeKafkaConsumerResponse() (response *DescribeKafkaConsumerResponse) {
	response = &DescribeKafkaConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Kafka协议消费信息
func (c *Client) DescribeKafkaConsumer(request *DescribeKafkaConsumerRequest) (response *DescribeKafkaConsumerResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaConsumerRequest()
	}
	response = NewDescribeKafkaConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewUploadServiceLogRequest() (request *UploadServiceLogRequest) {
	request = &UploadServiceLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UploadServiceLog")
	return
}

func NewUploadServiceLogResponse() (response *UploadServiceLogResponse) {
	response = &UploadServiceLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于将Agent服务日志写入到用户名下的服务日志Topic。
func (c *Client) UploadServiceLog(request *UploadServiceLogRequest) (response *UploadServiceLogResponse, err error) {
	if request == nil {
		request = NewUploadServiceLogRequest()
	}
	response = NewUploadServiceLogResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConsumerRequest() (request *CreateConsumerRequest) {
	request = &CreateConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateConsumer")
	return
}

func NewCreateConsumerResponse() (response *CreateConsumerResponse) {
	response = &CreateConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建投递Ckafka任务
func (c *Client) CreateConsumer(request *CreateConsumerRequest) (response *CreateConsumerResponse, err error) {
	if request == nil {
		request = NewCreateConsumerRequest()
	}
	response = NewCreateConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlarmRequest() (request *CreateAlarmRequest) {
	request = &CreateAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateAlarm")
	return
}

func NewCreateAlarmResponse() (response *CreateAlarmResponse) {
	response = &CreateAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于创建告警策略。
func (c *Client) CreateAlarm(request *CreateAlarmRequest) (response *CreateAlarmResponse, err error) {
	if request == nil {
		request = NewCreateAlarmRequest()
	}
	response = NewCreateAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewCreateMachineGroupRequest() (request *CreateMachineGroupRequest) {
	request = &CreateMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateMachineGroup")
	return
}

func NewCreateMachineGroupResponse() (response *CreateMachineGroupResponse) {
	response = &CreateMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建机器组
func (c *Client) CreateMachineGroup(request *CreateMachineGroupRequest) (response *CreateMachineGroupResponse, err error) {
	if request == nil {
		request = NewCreateMachineGroupRequest()
	}
	response = NewCreateMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformAutoCreateLogSetsRequest() (request *DescribeDataTransformAutoCreateLogSetsRequest) {
	request = &DescribeDataTransformAutoCreateLogSetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformAutoCreateLogSets")
	return
}

func NewDescribeDataTransformAutoCreateLogSetsResponse() (response *DescribeDataTransformAutoCreateLogSetsResponse) {
	response = &DescribeDataTransformAutoCreateLogSetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据加工任务创建的日志主题列表
func (c *Client) DescribeDataTransformAutoCreateLogSets(request *DescribeDataTransformAutoCreateLogSetsRequest) (response *DescribeDataTransformAutoCreateLogSetsResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformAutoCreateLogSetsRequest()
	}
	response = NewDescribeDataTransformAutoCreateLogSetsResponse()
	err = c.Send(request, response)
	return
}

func NewRemoveMachineRequest() (request *RemoveMachineRequest) {
	request = &RemoveMachineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "RemoveMachine")
	return
}

func NewRemoveMachineResponse() (response *RemoveMachineResponse) {
	response = &RemoveMachineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于剔除标签机器组中机器
func (c *Client) RemoveMachine(request *RemoveMachineRequest) (response *RemoveMachineResponse, err error) {
	if request == nil {
		request = NewRemoveMachineRequest()
	}
	response = NewRemoveMachineResponse()
	err = c.Send(request, response)
	return
}

func NewCreateShipperRequest() (request *CreateShipperRequest) {
	request = &CreateShipperRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateShipper")
	return
}

func NewCreateShipperResponse() (response *CreateShipperResponse) {
	response = &CreateShipperResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建新的投递规则，客户如果使用此接口，需要自行处理CLS对指定bucket的写权限。
func (c *Client) CreateShipper(request *CreateShipperRequest) (response *CreateShipperResponse, err error) {
	if request == nil {
		request = NewCreateShipperRequest()
	}
	response = NewCreateShipperResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteMetricConfigRequest() (request *DeleteMetricConfigRequest) {
	request = &DeleteMetricConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteMetricConfig")
	return
}

func NewDeleteMetricConfigResponse() (response *DeleteMetricConfigResponse) {
	response = &DeleteMetricConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除指标采集配置
func (c *Client) DeleteMetricConfig(request *DeleteMetricConfigRequest) (response *DeleteMetricConfigResponse, err error) {
	if request == nil {
		request = NewDeleteMetricConfigRequest()
	}
	response = NewDeleteMetricConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRebuildIndexTasksRequest() (request *DescribeRebuildIndexTasksRequest) {
	request = &DescribeRebuildIndexTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeRebuildIndexTasks")
	return
}

func NewDescribeRebuildIndexTasksResponse() (response *DescribeRebuildIndexTasksResponse) {
	response = &DescribeRebuildIndexTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取重建索引任务列表
func (c *Client) DescribeRebuildIndexTasks(request *DescribeRebuildIndexTasksRequest) (response *DescribeRebuildIndexTasksResponse, err error) {
	if request == nil {
		request = NewDescribeRebuildIndexTasksRequest()
	}
	response = NewDescribeRebuildIndexTasksResponse()
	err = c.Send(request, response)
	return
}

func NewModifyKafkaConsumerRequest() (request *ModifyKafkaConsumerRequest) {
	request = &ModifyKafkaConsumerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyKafkaConsumer")
	return
}

func NewModifyKafkaConsumerResponse() (response *ModifyKafkaConsumerResponse) {
	response = &ModifyKafkaConsumerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改Kafka协议消费信息
func (c *Client) ModifyKafkaConsumer(request *ModifyKafkaConsumerRequest) (response *ModifyKafkaConsumerResponse, err error) {
	if request == nil {
		request = NewModifyKafkaConsumerRequest()
	}
	response = NewModifyKafkaConsumerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteDashboardRequest() (request *DeleteDashboardRequest) {
	request = &DeleteDashboardRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteDashboard")
	return
}

func NewDeleteDashboardResponse() (response *DeleteDashboardResponse) {
	response = &DeleteDashboardResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除仪表盘
func (c *Client) DeleteDashboard(request *DeleteDashboardRequest) (response *DeleteDashboardResponse, err error) {
	if request == nil {
		request = NewDeleteDashboardRequest()
	}
	response = NewDeleteDashboardResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConsumerGroupRequest() (request *CreateConsumerGroupRequest) {
	request = &CreateConsumerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateConsumerGroup")
	return
}

func NewCreateConsumerGroupResponse() (response *CreateConsumerGroupResponse) {
	response = &CreateConsumerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建消费组
func (c *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (response *CreateConsumerGroupResponse, err error) {
	if request == nil {
		request = NewCreateConsumerGroupRequest()
	}
	response = NewCreateConsumerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeDataTransformAutoCreateTopicsRequest() (request *DescribeDataTransformAutoCreateTopicsRequest) {
	request = &DescribeDataTransformAutoCreateTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeDataTransformAutoCreateTopics")
	return
}

func NewDescribeDataTransformAutoCreateTopicsResponse() (response *DescribeDataTransformAutoCreateTopicsResponse) {
	response = &DescribeDataTransformAutoCreateTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据加工任务创建的日志主题列表
func (c *Client) DescribeDataTransformAutoCreateTopics(request *DescribeDataTransformAutoCreateTopicsRequest) (response *DescribeDataTransformAutoCreateTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeDataTransformAutoCreateTopicsRequest()
	}
	response = NewDescribeDataTransformAutoCreateTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigurationTemplateRequest() (request *DeleteConfigurationTemplateRequest) {
	request = &DeleteConfigurationTemplateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConfigurationTemplate")
	return
}

func NewDeleteConfigurationTemplateResponse() (response *DeleteConfigurationTemplateResponse) {
	response = &DeleteConfigurationTemplateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除配置模板
func (c *Client) DeleteConfigurationTemplate(request *DeleteConfigurationTemplateRequest) (response *DeleteConfigurationTemplateResponse, err error) {
	if request == nil {
		request = NewDeleteConfigurationTemplateRequest()
	}
	response = NewDeleteConfigurationTemplateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsumerOffsetsRequest() (request *DescribeConsumerOffsetsRequest) {
	request = &DescribeConsumerOffsetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeConsumerOffsets")
	return
}

func NewDescribeConsumerOffsetsResponse() (response *DescribeConsumerOffsetsResponse) {
	response = &DescribeConsumerOffsetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取消费组点位信息
func (c *Client) DescribeConsumerOffsets(request *DescribeConsumerOffsetsRequest) (response *DescribeConsumerOffsetsResponse, err error) {
	if request == nil {
		request = NewDescribeConsumerOffsetsRequest()
	}
	response = NewDescribeConsumerOffsetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeKafkaConsumeRequest() (request *DescribeKafkaConsumeRequest) {
	request = &DescribeKafkaConsumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeKafkaConsume")
	return
}

func NewDescribeKafkaConsumeResponse() (response *DescribeKafkaConsumeResponse) {
	response = &DescribeKafkaConsumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Kakfa消费信息
func (c *Client) DescribeKafkaConsume(request *DescribeKafkaConsumeRequest) (response *DescribeKafkaConsumeResponse, err error) {
	if request == nil {
		request = NewDescribeKafkaConsumeRequest()
	}
	response = NewDescribeKafkaConsumeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConfigFromMachineGroupRequest() (request *DeleteConfigFromMachineGroupRequest) {
	request = &DeleteConfigFromMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConfigFromMachineGroup")
	return
}

func NewDeleteConfigFromMachineGroupResponse() (response *DeleteConfigFromMachineGroupResponse) {
	response = &DeleteConfigFromMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用到机器组的采集配置
func (c *Client) DeleteConfigFromMachineGroup(request *DeleteConfigFromMachineGroupRequest) (response *DeleteConfigFromMachineGroupResponse, err error) {
	if request == nil {
		request = NewDeleteConfigFromMachineGroupRequest()
	}
	response = NewDeleteConfigFromMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
	request = &ModifyTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyTopic")
	return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
	response = &ModifyTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改日志主题。
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
	if request == nil {
		request = NewModifyTopicRequest()
	}
	response = NewModifyTopicResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAlarmShieldRequest() (request *CreateAlarmShieldRequest) {
	request = &CreateAlarmShieldRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CreateAlarmShield")
	return
}

func NewCreateAlarmShieldResponse() (response *CreateAlarmShieldResponse) {
	response = &CreateAlarmShieldResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建告警屏蔽规则
func (c *Client) CreateAlarmShield(request *CreateAlarmShieldRequest) (response *CreateAlarmShieldResponse, err error) {
	if request == nil {
		request = NewCreateAlarmShieldRequest()
	}
	response = NewCreateAlarmShieldResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteScheduledSqlRequest() (request *DeleteScheduledSqlRequest) {
	request = &DeleteScheduledSqlRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteScheduledSql")
	return
}

func NewDeleteScheduledSqlResponse() (response *DeleteScheduledSqlResponse) {
	response = &DeleteScheduledSqlResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除定时SQL分析任务
func (c *Client) DeleteScheduledSql(request *DeleteScheduledSqlRequest) (response *DeleteScheduledSqlResponse, err error) {
	if request == nil {
		request = NewDeleteScheduledSqlRequest()
	}
	response = NewDeleteScheduledSqlResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAlarmRequest() (request *ModifyAlarmRequest) {
	request = &ModifyAlarmRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyAlarm")
	return
}

func NewModifyAlarmResponse() (response *ModifyAlarmResponse) {
	response = &ModifyAlarmResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改告警策略。需要至少修改一项有效内容。
func (c *Client) ModifyAlarm(request *ModifyAlarmRequest) (response *ModifyAlarmResponse, err error) {
	if request == nil {
		request = NewModifyAlarmRequest()
	}
	response = NewModifyAlarmResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDataTransformRequest() (request *ModifyDataTransformRequest) {
	request = &ModifyDataTransformRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyDataTransform")
	return
}

func NewModifyDataTransformResponse() (response *ModifyDataTransformResponse) {
	response = &ModifyDataTransformResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改数据加工任务
func (c *Client) ModifyDataTransform(request *ModifyDataTransformRequest) (response *ModifyDataTransformResponse, err error) {
	if request == nil {
		request = NewModifyDataTransformRequest()
	}
	response = NewModifyDataTransformResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteShipperRequest() (request *DeleteShipperRequest) {
	request = &DeleteShipperRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteShipper")
	return
}

func NewDeleteShipperResponse() (response *DeleteShipperResponse) {
	response = &DeleteShipperResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除投递规则
func (c *Client) DeleteShipper(request *DeleteShipperRequest) (response *DeleteShipperResponse, err error) {
	if request == nil {
		request = NewDeleteShipperRequest()
	}
	response = NewDeleteShipperResponse()
	err = c.Send(request, response)
	return
}

func NewModifyMachineGroupRequest() (request *ModifyMachineGroupRequest) {
	request = &ModifyMachineGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyMachineGroup")
	return
}

func NewModifyMachineGroupResponse() (response *ModifyMachineGroupResponse) {
	response = &ModifyMachineGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改机器组
func (c *Client) ModifyMachineGroup(request *ModifyMachineGroupRequest) (response *ModifyMachineGroupResponse, err error) {
	if request == nil {
		request = NewModifyMachineGroupRequest()
	}
	response = NewModifyMachineGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopicExtendConfigRequest() (request *ModifyTopicExtendConfigRequest) {
	request = &ModifyTopicExtendConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyTopicExtendConfig")
	return
}

func NewModifyTopicExtendConfigResponse() (response *ModifyTopicExtendConfigResponse) {
	response = &ModifyTopicExtendConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改采集配置(clb专用)。
func (c *Client) ModifyTopicExtendConfig(request *ModifyTopicExtendConfigRequest) (response *ModifyTopicExtendConfigResponse, err error) {
	if request == nil {
		request = NewModifyTopicExtendConfigRequest()
	}
	response = NewModifyTopicExtendConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteConsumerGroupRequest() (request *DeleteConsumerGroupRequest) {
	request = &DeleteConsumerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteConsumerGroup")
	return
}

func NewDeleteConsumerGroupResponse() (response *DeleteConsumerGroupResponse) {
	response = &DeleteConsumerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除消费组
func (c *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (response *DeleteConsumerGroupResponse, err error) {
	if request == nil {
		request = NewDeleteConsumerGroupRequest()
	}
	response = NewDeleteConsumerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateAgentStatusRequest() (request *UpdateAgentStatusRequest) {
	request = &UpdateAgentStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "UpdateAgentStatus")
	return
}

func NewUpdateAgentStatusResponse() (response *UpdateAgentStatusResponse) {
	response = &UpdateAgentStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 上报采集机器Agent的升级状态
func (c *Client) UpdateAgentStatus(request *UpdateAgentStatusRequest) (response *UpdateAgentStatusResponse, err error) {
	if request == nil {
		request = NewUpdateAgentStatusRequest()
	}
	response = NewUpdateAgentStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeFoldersRequest() (request *DescribeFoldersRequest) {
	request = &DescribeFoldersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeFolders")
	return
}

func NewDescribeFoldersResponse() (response *DescribeFoldersResponse) {
	response = &DescribeFoldersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取文件夹列表
func (c *Client) DescribeFolders(request *DescribeFoldersRequest) (response *DescribeFoldersResponse, err error) {
	if request == nil {
		request = NewDescribeFoldersRequest()
	}
	response = NewDescribeFoldersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeScheduledSqlInfoRequest() (request *DescribeScheduledSqlInfoRequest) {
	request = &DescribeScheduledSqlInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeScheduledSqlInfo")
	return
}

func NewDescribeScheduledSqlInfoResponse() (response *DescribeScheduledSqlInfoResponse) {
	response = &DescribeScheduledSqlInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于获取定时SQL分析任务列表
func (c *Client) DescribeScheduledSqlInfo(request *DescribeScheduledSqlInfoRequest) (response *DescribeScheduledSqlInfoResponse, err error) {
	if request == nil {
		request = NewDescribeScheduledSqlInfoRequest()
	}
	response = NewDescribeScheduledSqlInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRecordingRuleTaskRequest() (request *DeleteRecordingRuleTaskRequest) {
	request = &DeleteRecordingRuleTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteRecordingRuleTask")
	return
}

func NewDeleteRecordingRuleTaskResponse() (response *DeleteRecordingRuleTaskResponse) {
	response = &DeleteRecordingRuleTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除预聚合任务
func (c *Client) DeleteRecordingRuleTask(request *DeleteRecordingRuleTaskRequest) (response *DeleteRecordingRuleTaskResponse, err error) {
	if request == nil {
		request = NewDeleteRecordingRuleTaskRequest()
	}
	response = NewDeleteRecordingRuleTaskResponse()
	err = c.Send(request, response)
	return
}

func NewModifyWebCallbackRequest() (request *ModifyWebCallbackRequest) {
	request = &ModifyWebCallbackRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "ModifyWebCallback")
	return
}

func NewModifyWebCallbackResponse() (response *ModifyWebCallbackResponse) {
	response = &ModifyWebCallbackResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改告警渠道回调配置
func (c *Client) ModifyWebCallback(request *ModifyWebCallbackRequest) (response *ModifyWebCallbackResponse, err error) {
	if request == nil {
		request = NewModifyWebCallbackRequest()
	}
	response = NewModifyWebCallbackResponse()
	err = c.Send(request, response)
	return
}

func NewCheckRemoteWriteTaskConnectRequest() (request *CheckRemoteWriteTaskConnectRequest) {
	request = &CheckRemoteWriteTaskConnectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "CheckRemoteWriteTaskConnect")
	return
}

func NewCheckRemoteWriteTaskConnectResponse() (response *CheckRemoteWriteTaskConnectResponse) {
	response = &CheckRemoteWriteTaskConnectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检查RemoteWrite任务联通性
func (c *Client) CheckRemoteWriteTaskConnect(request *CheckRemoteWriteTaskConnectRequest) (response *CheckRemoteWriteTaskConnectResponse, err error) {
	if request == nil {
		request = NewCheckRemoteWriteTaskConnectRequest()
	}
	response = NewCheckRemoteWriteTaskConnectResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
	request = &DeleteTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DeleteTopic")
	return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
	response = &DeleteTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于删除日志主题。
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
	if request == nil {
		request = NewDeleteTopicRequest()
	}
	response = NewDeleteTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWebCallbacksRequest() (request *DescribeWebCallbacksRequest) {
	request = &DescribeWebCallbacksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeWebCallbacks")
	return
}

func NewDescribeWebCallbacksResponse() (response *DescribeWebCallbacksResponse) {
	response = &DescribeWebCallbacksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取告警渠道回调配置列表
func (c *Client) DescribeWebCallbacks(request *DescribeWebCallbacksRequest) (response *DescribeWebCallbacksResponse, err error) {
	if request == nil {
		request = NewDescribeWebCallbacksRequest()
	}
	response = NewDescribeWebCallbacksResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLatestUserLogRequest() (request *DescribeLatestUserLogRequest) {
	request = &DescribeLatestUserLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "DescribeLatestUserLog")
	return
}

func NewDescribeLatestUserLogResponse() (response *DescribeLatestUserLogResponse) {
	response = &DescribeLatestUserLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户最新一条日志
func (c *Client) DescribeLatestUserLog(request *DescribeLatestUserLogRequest) (response *DescribeLatestUserLogResponse, err error) {
	if request == nil {
		request = NewDescribeLatestUserLogRequest()
	}
	response = NewDescribeLatestUserLogResponse()
	err = c.Send(request, response)
	return
}

func NewRetryScheduledSqlTaskRequest() (request *RetryScheduledSqlTaskRequest) {
	request = &RetryScheduledSqlTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cls", APIVersion, "RetryScheduledSqlTask")
	return
}

func NewRetryScheduledSqlTaskResponse() (response *RetryScheduledSqlTaskResponse) {
	response = &RetryScheduledSqlTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重试失败的定时SQL分析任务
func (c *Client) RetryScheduledSqlTask(request *RetryScheduledSqlTaskRequest) (response *RetryScheduledSqlTaskResponse, err error) {
	if request == nil {
		request = NewRetryScheduledSqlTaskRequest()
	}
	response = NewRetryScheduledSqlTaskResponse()
	err = c.Send(request, response)
	return
}

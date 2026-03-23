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

package v20210622

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2021-06-22"

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

func NewDescribeServiceRequest() (request *DescribeServiceRequest) {
	request = &DescribeServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeService")
	return
}

func NewDescribeServiceResponse() (response *DescribeServiceResponse) {
	response = &DescribeServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用信息
func (c *Client) DescribeService(request *DescribeServiceRequest) (response *DescribeServiceResponse, err error) {
	if request == nil {
		request = NewDescribeServiceRequest()
	}
	response = NewDescribeServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmServiceMetricRequest() (request *DescribeApmServiceMetricRequest) {
	request = &DescribeApmServiceMetricRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmServiceMetric")
	return
}

func NewDescribeApmServiceMetricResponse() (response *DescribeApmServiceMetricResponse) {
	response = &DescribeApmServiceMetricResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 APM 应用指标列表
func (c *Client) DescribeApmServiceMetric(request *DescribeApmServiceMetricRequest) (response *DescribeApmServiceMetricResponse, err error) {
	if request == nil {
		request = NewDescribeApmServiceMetricRequest()
	}
	response = NewDescribeApmServiceMetricResponse()
	err = c.Send(request, response)
	return
}

func NewCreateQueryViewRequest() (request *CreateQueryViewRequest) {
	request = &CreateQueryViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "CreateQueryView")
	return
}

func NewCreateQueryViewResponse() (response *CreateQueryViewResponse) {
	response = &CreateQueryViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建查询视图
func (c *Client) CreateQueryView(request *CreateQueryViewRequest) (response *CreateQueryViewResponse, err error) {
	if request == nil {
		request = NewCreateQueryViewRequest()
	}
	response = NewCreateQueryViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralMetricDataRequest() (request *DescribeGeneralMetricDataRequest) {
	request = &DescribeGeneralMetricDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralMetricData")
	return
}

func NewDescribeGeneralMetricDataResponse() (response *DescribeGeneralMetricDataResponse) {
	response = &DescribeGeneralMetricDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指标数据通用接口。用户根据需要上送请求参数，返回对应的指标数据。
// 接口调用频率限制为：20次/秒，1200次/分钟。单请求的数据点数限制为1440个。
func (c *Client) DescribeGeneralMetricData(request *DescribeGeneralMetricDataRequest) (response *DescribeGeneralMetricDataResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralMetricDataRequest()
	}
	response = NewDescribeGeneralMetricDataResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentListRequest() (request *DescribeAgentListRequest) {
	request = &DescribeAgentListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeAgentList")
	return
}

func NewDescribeAgentListResponse() (response *DescribeAgentListResponse) {
	response = &DescribeAgentListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询探针agent列表
func (c *Client) DescribeAgentList(request *DescribeAgentListRequest) (response *DescribeAgentListResponse, err error) {
	if request == nil {
		request = NewDescribeAgentListRequest()
	}
	response = NewDescribeAgentListResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateApmInstanceRequest() (request *TerminateApmInstanceRequest) {
	request = &TerminateApmInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "TerminateApmInstance")
	return
}

func NewTerminateApmInstanceResponse() (response *TerminateApmInstanceResponse) {
	response = &TerminateApmInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 销毁 APM 业务系统
func (c *Client) TerminateApmInstance(request *TerminateApmInstanceRequest) (response *TerminateApmInstanceResponse, err error) {
	if request == nil {
		request = NewTerminateApmInstanceRequest()
	}
	response = NewTerminateApmInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceOverviewRequest() (request *DescribeServiceOverviewRequest) {
	request = &DescribeServiceOverviewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceOverview")
	return
}

func NewDescribeServiceOverviewResponse() (response *DescribeServiceOverviewResponse) {
	response = &DescribeServiceOverviewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 应用概览数据拉取
func (c *Client) DescribeServiceOverview(request *DescribeServiceOverviewRequest) (response *DescribeServiceOverviewResponse, err error) {
	if request == nil {
		request = NewDescribeServiceOverviewRequest()
	}
	response = NewDescribeServiceOverviewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApmApplicationConfigRequest() (request *ModifyApmApplicationConfigRequest) {
	request = &ModifyApmApplicationConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyApmApplicationConfig")
	return
}

func NewModifyApmApplicationConfigResponse() (response *ModifyApmApplicationConfigResponse) {
	response = &ModifyApmApplicationConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改应用配置接口
func (c *Client) ModifyApmApplicationConfig(request *ModifyApmApplicationConfigRequest) (response *ModifyApmApplicationConfigResponse, err error) {
	if request == nil {
		request = NewModifyApmApplicationConfigRequest()
	}
	response = NewModifyApmApplicationConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopologyNewRequest() (request *DescribeTopologyNewRequest) {
	request = &DescribeTopologyNewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTopologyNew")
	return
}

func NewDescribeTopologyNewResponse() (response *DescribeTopologyNewResponse) {
	response = &DescribeTopologyNewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据应用名查询服务拓扑图
func (c *Client) DescribeTopologyNew(request *DescribeTopologyNewRequest) (response *DescribeTopologyNewResponse, err error) {
	if request == nil {
		request = NewDescribeTopologyNewRequest()
	}
	response = NewDescribeTopologyNewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetricRecordsRequest() (request *DescribeMetricRecordsRequest) {
	request = &DescribeMetricRecordsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricRecords")
	return
}

func NewDescribeMetricRecordsResponse() (response *DescribeMetricRecordsResponse) {
	response = &DescribeMetricRecordsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指标列表接口，查询指标更推荐使用DescribeGeneralMetricData接口
func (c *Client) DescribeMetricRecords(request *DescribeMetricRecordsRequest) (response *DescribeMetricRecordsResponse, err error) {
	if request == nil {
		request = NewDescribeMetricRecordsRequest()
	}
	response = NewDescribeMetricRecordsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceBriefsRequest() (request *DescribeInstanceBriefsRequest) {
	request = &DescribeInstanceBriefsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeInstanceBriefs")
	return
}

func NewDescribeInstanceBriefsResponse() (response *DescribeInstanceBriefsResponse) {
	response = &DescribeInstanceBriefsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 APM 业务系统简介列表
func (c *Client) DescribeInstanceBriefs(request *DescribeInstanceBriefsRequest) (response *DescribeInstanceBriefsResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceBriefsRequest()
	}
	response = NewDescribeInstanceBriefsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGeneralApmApplicationConfigRequest() (request *ModifyGeneralApmApplicationConfigRequest) {
	request = &ModifyGeneralApmApplicationConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyGeneralApmApplicationConfig")
	return
}

func NewModifyGeneralApmApplicationConfigResponse() (response *ModifyGeneralApmApplicationConfigResponse) {
	response = &ModifyGeneralApmApplicationConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 对外开放的openApi，客户可以灵活的指定需要修改的字段，再加入需要修改的服务列表.
func (c *Client) ModifyGeneralApmApplicationConfig(request *ModifyGeneralApmApplicationConfigRequest) (response *ModifyGeneralApmApplicationConfigResponse, err error) {
	if request == nil {
		request = NewModifyGeneralApmApplicationConfigRequest()
	}
	response = NewModifyGeneralApmApplicationConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentInfoRequest() (request *DescribeAgentInfoRequest) {
	request = &DescribeAgentInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeAgentInfo")
	return
}

func NewDescribeAgentInfoResponse() (response *DescribeAgentInfoResponse) {
	response = &DescribeAgentInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于提示用户当前应用是否存在探针需要进行更新
func (c *Client) DescribeAgentInfo(request *DescribeAgentInfoRequest) (response *DescribeAgentInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAgentInfoRequest()
	}
	response = NewDescribeAgentInfoResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApmServiceRequest() (request *ModifyApmServiceRequest) {
	request = &ModifyApmServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyApmService")
	return
}

func NewModifyApmServiceResponse() (response *ModifyApmServiceResponse) {
	response = &ModifyApmServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改apm应用信息
func (c *Client) ModifyApmService(request *ModifyApmServiceRequest) (response *ModifyApmServiceResponse, err error) {
	if request == nil {
		request = NewModifyApmServiceRequest()
	}
	response = NewModifyApmServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmInstancesRequest() (request *DescribeApmInstancesRequest) {
	request = &DescribeApmInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInstances")
	return
}

func NewDescribeApmInstancesResponse() (response *DescribeApmInstancesResponse) {
	response = &DescribeApmInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 APM 业务系统列表
func (c *Client) DescribeApmInstances(request *DescribeApmInstancesRequest) (response *DescribeApmInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeApmInstancesRequest()
	}
	response = NewDescribeApmInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmServiceRequest() (request *DescribeApmServiceRequest) {
	request = &DescribeApmServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmService")
	return
}

func NewDescribeApmServiceResponse() (response *DescribeApmServiceResponse) {
	response = &DescribeApmServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用详细信息
func (c *Client) DescribeApmService(request *DescribeApmServiceRequest) (response *DescribeApmServiceResponse, err error) {
	if request == nil {
		request = NewDescribeApmServiceRequest()
	}
	response = NewDescribeApmServiceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopologyViewRequest() (request *ModifyTopologyViewRequest) {
	request = &ModifyTopologyViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyTopologyView")
	return
}

func NewModifyTopologyViewResponse() (response *ModifyTopologyViewResponse) {
	response = &ModifyTopologyViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建及修改方案视图，Id为空为创建，非空为修改
func (c *Client) ModifyTopologyView(request *ModifyTopologyViewRequest) (response *ModifyTopologyViewResponse, err error) {
	if request == nil {
		request = NewModifyTopologyViewRequest()
	}
	response = NewModifyTopologyViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagCountValuesRequest() (request *DescribeTagCountValuesRequest) {
	request = &DescribeTagCountValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTagCountValues")
	return
}

func NewDescribeTagCountValuesResponse() (response *DescribeTagCountValuesResponse) {
	response = &DescribeTagCountValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据维度名和过滤条件，查询维度分组数据
func (c *Client) DescribeTagCountValues(request *DescribeTagCountValuesRequest) (response *DescribeTagCountValuesResponse, err error) {
	if request == nil {
		request = NewDescribeTagCountValuesRequest()
	}
	response = NewDescribeTagCountValuesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralApmApplicationConfigRequest() (request *DescribeGeneralApmApplicationConfigRequest) {
	request = &DescribeGeneralApmApplicationConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralApmApplicationConfig")
	return
}

func NewDescribeGeneralApmApplicationConfigResponse() (response *DescribeGeneralApmApplicationConfigResponse) {
	response = &DescribeGeneralApmApplicationConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用配置信息
func (c *Client) DescribeGeneralApmApplicationConfig(request *DescribeGeneralApmApplicationConfigRequest) (response *DescribeGeneralApmApplicationConfigResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralApmApplicationConfigRequest()
	}
	response = NewDescribeGeneralApmApplicationConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceLinkRequest() (request *DescribeServiceLinkRequest) {
	request = &DescribeServiceLinkRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceLink")
	return
}

func NewDescribeServiceLinkResponse() (response *DescribeServiceLinkResponse) {
	response = &DescribeServiceLinkResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过接口获取应用的上游或者下游应用，并且获取它们的对应指标
func (c *Client) DescribeServiceLink(request *DescribeServiceLinkRequest) (response *DescribeServiceLinkResponse, err error) {
	if request == nil {
		request = NewDescribeServiceLinkRequest()
	}
	response = NewDescribeServiceLinkResponse()
	err = c.Send(request, response)
	return
}

func NewModifyQueryViewRequest() (request *ModifyQueryViewRequest) {
	request = &ModifyQueryViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyQueryView")
	return
}

func NewModifyQueryViewResponse() (response *ModifyQueryViewResponse) {
	response = &ModifyQueryViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改查询视图
func (c *Client) ModifyQueryView(request *ModifyQueryViewRequest) (response *ModifyQueryViewResponse, err error) {
	if request == nil {
		request = NewModifyQueryViewRequest()
	}
	response = NewModifyQueryViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpanThreadAnatomyListRequest() (request *DescribeSpanThreadAnatomyListRequest) {
	request = &DescribeSpanThreadAnatomyListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeSpanThreadAnatomyList")
	return
}

func NewDescribeSpanThreadAnatomyListResponse() (response *DescribeSpanThreadAnatomyListResponse) {
	response = &DescribeSpanThreadAnatomyListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询线程剖析信息
func (c *Client) DescribeSpanThreadAnatomyList(request *DescribeSpanThreadAnatomyListRequest) (response *DescribeSpanThreadAnatomyListResponse, err error) {
	if request == nil {
		request = NewDescribeSpanThreadAnatomyListRequest()
	}
	response = NewDescribeSpanThreadAnatomyListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMetricLineDataRequest() (request *DescribeMetricLineDataRequest) {
	request = &DescribeMetricLineDataRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeMetricLineData")
	return
}

func NewDescribeMetricLineDataResponse() (response *DescribeMetricLineDataResponse) {
	response = &DescribeMetricLineDataResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询不同维度聚合的指标数据，返回曲线图数据
func (c *Client) DescribeMetricLineData(request *DescribeMetricLineDataRequest) (response *DescribeMetricLineDataResponse, err error) {
	if request == nil {
		request = NewDescribeMetricLineDataRequest()
	}
	response = NewDescribeMetricLineDataResponse()
	err = c.Send(request, response)
	return
}

func NewTerminateApmServiceRequest() (request *TerminateApmServiceRequest) {
	request = &TerminateApmServiceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "TerminateApmService")
	return
}

func NewTerminateApmServiceResponse() (response *TerminateApmServiceResponse) {
	response = &TerminateApmServiceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除应用资源
func (c *Client) TerminateApmService(request *TerminateApmServiceRequest) (response *TerminateApmServiceResponse, err error) {
	if request == nil {
		request = NewTerminateApmServiceRequest()
	}
	response = NewTerminateApmServiceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpanTreeByIDRequest() (request *DescribeSpanTreeByIDRequest) {
	request = &DescribeSpanTreeByIDRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeSpanTreeByID")
	return
}

func NewDescribeSpanTreeByIDResponse() (response *DescribeSpanTreeByIDResponse) {
	response = &DescribeSpanTreeByIDResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询trace详情接口信息
func (c *Client) DescribeSpanTreeByID(request *DescribeSpanTreeByIDRequest) (response *DescribeSpanTreeByIDResponse, err error) {
	if request == nil {
		request = NewDescribeSpanTreeByIDRequest()
	}
	response = NewDescribeSpanTreeByIDResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopologyViewRequest() (request *DeleteTopologyViewRequest) {
	request = &DeleteTopologyViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DeleteTopologyView")
	return
}

func NewDeleteTopologyViewResponse() (response *DeleteTopologyViewResponse) {
	response = &DeleteTopologyViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过ID删除视图方案
func (c *Client) DeleteTopologyView(request *DeleteTopologyViewRequest) (response *DeleteTopologyViewResponse, err error) {
	if request == nil {
		request = NewDeleteTopologyViewRequest()
	}
	response = NewDeleteTopologyViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGeneralSpanListRequest() (request *DescribeGeneralSpanListRequest) {
	request = &DescribeGeneralSpanListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeGeneralSpanList")
	return
}

func NewDescribeGeneralSpanListResponse() (response *DescribeGeneralSpanListResponse) {
	response = &DescribeGeneralSpanListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通用查询调用链列表
func (c *Client) DescribeGeneralSpanList(request *DescribeGeneralSpanListRequest) (response *DescribeGeneralSpanListResponse, err error) {
	if request == nil {
		request = NewDescribeGeneralSpanListRequest()
	}
	response = NewDescribeGeneralSpanListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeServiceNodesRequest() (request *DescribeServiceNodesRequest) {
	request = &DescribeServiceNodesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeServiceNodes")
	return
}

func NewDescribeServiceNodesResponse() (response *DescribeServiceNodesResponse) {
	response = &DescribeServiceNodesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取业务系统下所有应用节点列表
func (c *Client) DescribeServiceNodes(request *DescribeServiceNodesRequest) (response *DescribeServiceNodesResponse, err error) {
	if request == nil {
		request = NewDescribeServiceNodesRequest()
	}
	response = NewDescribeServiceNodesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTagValuesRequest() (request *DescribeTagValuesRequest) {
	request = &DescribeTagValuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTagValues")
	return
}

func NewDescribeTagValuesResponse() (response *DescribeTagValuesResponse) {
	response = &DescribeTagValuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据维度名和过滤条件，查询维度数据.
func (c *Client) DescribeTagValues(request *DescribeTagValuesRequest) (response *DescribeTagValuesResponse, err error) {
	if request == nil {
		request = NewDescribeTagValuesRequest()
	}
	response = NewDescribeTagValuesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmAgentRequest() (request *DescribeApmAgentRequest) {
	request = &DescribeApmAgentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmAgent")
	return
}

func NewDescribeApmAgentResponse() (response *DescribeApmAgentResponse) {
	response = &DescribeApmAgentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取 APM 接入点
func (c *Client) DescribeApmAgent(request *DescribeApmAgentRequest) (response *DescribeApmAgentResponse, err error) {
	if request == nil {
		request = NewDescribeApmAgentRequest()
	}
	response = NewDescribeApmAgentResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmApplicationConfigRequest() (request *DescribeApmApplicationConfigRequest) {
	request = &DescribeApmApplicationConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmApplicationConfig")
	return
}

func NewDescribeApmApplicationConfigResponse() (response *DescribeApmApplicationConfigResponse) {
	response = &DescribeApmApplicationConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用配置接口
func (c *Client) DescribeApmApplicationConfig(request *DescribeApmApplicationConfigRequest) (response *DescribeApmApplicationConfigResponse, err error) {
	if request == nil {
		request = NewDescribeApmApplicationConfigRequest()
	}
	response = NewDescribeApmApplicationConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmServiceBriefRequest() (request *DescribeApmServiceBriefRequest) {
	request = &DescribeApmServiceBriefRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmServiceBrief")
	return
}

func NewDescribeApmServiceBriefResponse() (response *DescribeApmServiceBriefResponse) {
	response = &DescribeApmServiceBriefResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取应用资源简介列表
func (c *Client) DescribeApmServiceBrief(request *DescribeApmServiceBriefRequest) (response *DescribeApmServiceBriefResponse, err error) {
	if request == nil {
		request = NewDescribeApmServiceBriefRequest()
	}
	response = NewDescribeApmServiceBriefResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeQueryViewRequest() (request *DescribeQueryViewRequest) {
	request = &DescribeQueryViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeQueryView")
	return
}

func NewDescribeQueryViewResponse() (response *DescribeQueryViewResponse) {
	response = &DescribeQueryViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据视图 ID 获取对应的视图详情
func (c *Client) DescribeQueryView(request *DescribeQueryViewRequest) (response *DescribeQueryViewResponse, err error) {
	if request == nil {
		request = NewDescribeQueryViewRequest()
	}
	response = NewDescribeQueryViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopologyViewRequest() (request *DescribeTopologyViewRequest) {
	request = &DescribeTopologyViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTopologyView")
	return
}

func NewDescribeTopologyViewResponse() (response *DescribeTopologyViewResponse) {
	response = &DescribeTopologyViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取业务系统下的所有视图方案
func (c *Client) DescribeTopologyView(request *DescribeTopologyViewRequest) (response *DescribeTopologyViewResponse, err error) {
	if request == nil {
		request = NewDescribeTopologyViewRequest()
	}
	response = NewDescribeTopologyViewResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTraceListRequest() (request *DescribeTraceListRequest) {
	request = &DescribeTraceListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeTraceList")
	return
}

func NewDescribeTraceListResponse() (response *DescribeTraceListResponse) {
	response = &DescribeTraceListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询应用示例下的trace分页列表数据
func (c *Client) DescribeTraceList(request *DescribeTraceListRequest) (response *DescribeTraceListResponse, err error) {
	if request == nil {
		request = NewDescribeTraceListRequest()
	}
	response = NewDescribeTraceListResponse()
	err = c.Send(request, response)
	return
}

func NewCreateApmInstanceRequest() (request *CreateApmInstanceRequest) {
	request = &CreateApmInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "CreateApmInstance")
	return
}

func NewCreateApmInstanceResponse() (response *CreateApmInstanceResponse) {
	response = &CreateApmInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 业务购买 APM 业务系统，调用该接口创建
func (c *Client) CreateApmInstance(request *CreateApmInstanceRequest) (response *CreateApmInstanceResponse, err error) {
	if request == nil {
		request = NewCreateApmInstanceRequest()
	}
	response = NewCreateApmInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteQueryViewRequest() (request *DeleteQueryViewRequest) {
	request = &DeleteQueryViewRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DeleteQueryView")
	return
}

func NewDeleteQueryViewResponse() (response *DeleteQueryViewResponse) {
	response = &DeleteQueryViewResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除查询视图
func (c *Client) DeleteQueryView(request *DeleteQueryViewRequest) (response *DeleteQueryViewResponse, err error) {
	if request == nil {
		request = NewDeleteQueryViewRequest()
	}
	response = NewDeleteQueryViewResponse()
	err = c.Send(request, response)
	return
}

func NewModifyApmInstanceRequest() (request *ModifyApmInstanceRequest) {
	request = &ModifyApmInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "ModifyApmInstance")
	return
}

func NewModifyApmInstanceResponse() (response *ModifyApmInstanceResponse) {
	response = &ModifyApmInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改APM业务系统接口
func (c *Client) ModifyApmInstance(request *ModifyApmInstanceRequest) (response *ModifyApmInstanceResponse, err error) {
	if request == nil {
		request = NewModifyApmInstanceRequest()
	}
	response = NewModifyApmInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteExpiredApplicationConfigRequest() (request *DeleteExpiredApplicationConfigRequest) {
	request = &DeleteExpiredApplicationConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DeleteExpiredApplicationConfig")
	return
}

func NewDeleteExpiredApplicationConfigResponse() (response *DeleteExpiredApplicationConfigResponse) {
	response = &DeleteExpiredApplicationConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除过期的应用配置
func (c *Client) DeleteExpiredApplicationConfig(request *DeleteExpiredApplicationConfigRequest) (response *DeleteExpiredApplicationConfigResponse, err error) {
	if request == nil {
		request = NewDeleteExpiredApplicationConfigRequest()
	}
	response = NewDeleteExpiredApplicationConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeApmInfoByAppIdRequest() (request *DescribeApmInfoByAppIdRequest) {
	request = &DescribeApmInfoByAppIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeApmInfoByAppId")
	return
}

func NewDescribeApmInfoByAppIdResponse() (response *DescribeApmInfoByAppIdResponse) {
	response = &DescribeApmInfoByAppIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回appId对应的剩余上报额度信息
func (c *Client) DescribeApmInfoByAppId(request *DescribeApmInfoByAppIdRequest) (response *DescribeApmInfoByAppIdResponse, err error) {
	if request == nil {
		request = NewDescribeApmInfoByAppIdRequest()
	}
	response = NewDescribeApmInfoByAppIdResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllQueryViewsRequest() (request *DescribeAllQueryViewsRequest) {
	request = &DescribeAllQueryViewsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeAllQueryViews")
	return
}

func NewDescribeAllQueryViewsResponse() (response *DescribeAllQueryViewsResponse) {
	response = &DescribeAllQueryViewsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取对应业务系统下所有的查询视图集合
func (c *Client) DescribeAllQueryViews(request *DescribeAllQueryViewsRequest) (response *DescribeAllQueryViewsResponse, err error) {
	if request == nil {
		request = NewDescribeAllQueryViewsRequest()
	}
	response = NewDescribeAllQueryViewsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSpanTagListRequest() (request *DescribeSpanTagListRequest) {
	request = &DescribeSpanTagListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("apm", APIVersion, "DescribeSpanTagList")
	return
}

func NewDescribeSpanTagListResponse() (response *DescribeSpanTagListResponse) {
	response = &DescribeSpanTagListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 链路列表页查询相关Tag字段展示接口
func (c *Client) DescribeSpanTagList(request *DescribeSpanTagListRequest) (response *DescribeSpanTagListResponse, err error) {
	if request == nil {
		request = NewDescribeSpanTagListRequest()
	}
	response = NewDescribeSpanTagListResponse()
	err = c.Send(request, response)
	return
}

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

package v20180317

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2018-03-17"

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

func NewDeleteRuleRequest() (request *DeleteRuleRequest) {
	request = &DeleteRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteRule")
	return
}

func NewDeleteRuleResponse() (response *DeleteRuleResponse) {
	response = &DeleteRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteRule 接口用来删除负载均衡实例七层监听器下的转发规则。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) DeleteRule(request *DeleteRuleRequest) (response *DeleteRuleResponse, err error) {
	if request == nil {
		request = NewDeleteRuleRequest()
	}
	response = NewDeleteRuleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockIPTaskRequest() (request *DescribeBlockIPTaskRequest) {
	request = &DescribeBlockIPTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBlockIPTask")
	return
}

func NewDescribeBlockIPTaskResponse() (response *DescribeBlockIPTaskResponse) {
	response = &DescribeBlockIPTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据 ModifyBlockIPList 接口返回的异步任务的ID，查询封禁IP（黑名单）异步任务的执行状态。（接口灰度中，如需使用请提工单）
func (c *Client) DescribeBlockIPTask(request *DescribeBlockIPTaskRequest) (response *DescribeBlockIPTaskResponse, err error) {
	if request == nil {
		request = NewDescribeBlockIPTaskRequest()
	}
	response = NewDescribeBlockIPTaskResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeListenersRequest() (request *DescribeListenersRequest) {
	request = &DescribeListenersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeListeners")
	return
}

func NewDescribeListenersResponse() (response *DescribeListenersResponse) {
	response = &DescribeListenersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeListeners 接口可根据负载均衡器 ID，监听器的协议或端口作为过滤条件获取监听器列表。如果不指定任何过滤条件，则返回该负载均衡实例下的所有监听器。
func (c *Client) DescribeListeners(request *DescribeListenersRequest) (response *DescribeListenersResponse, err error) {
	if request == nil {
		request = NewDescribeListenersRequest()
	}
	response = NewDescribeListenersResponse()
	err = c.Send(request, response)
	return
}

func NewUpLoadCertRequest() (request *UpLoadCertRequest) {
	request = &UpLoadCertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "UpLoadCert")
	return
}

func NewUpLoadCertResponse() (response *UpLoadCertResponse) {
	response = &UpLoadCertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端上传ssl证书
func (c *Client) UpLoadCert(request *UpLoadCertRequest) (response *UpLoadCertResponse, err error) {
	if request == nil {
		request = NewUpLoadCertRequest()
	}
	response = NewUpLoadCertResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterResourcesRequest() (request *DescribeClusterResourcesRequest) {
	request = &DescribeClusterResourcesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeClusterResources")
	return
}

func NewDescribeClusterResourcesResponse() (response *DescribeClusterResourcesResponse) {
	response = &DescribeClusterResourcesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询独占集群中资源列表，支持按集群ID、vip、负载均衡ID、是否闲置为过滤条件检索
func (c *Client) DescribeClusterResources(request *DescribeClusterResourcesRequest) (response *DescribeClusterResourcesResponse, err error) {
	if request == nil {
		request = NewDescribeClusterResourcesRequest()
	}
	response = NewDescribeClusterResourcesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCertAliasRequest() (request *ModifyCertAliasRequest) {
	request = &ModifyCertAliasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyCertAlias")
	return
}

func NewModifyCertAliasResponse() (response *ModifyCertAliasResponse) {
	response = &ModifyCertAliasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改租户端证书备注
func (c *Client) ModifyCertAlias(request *ModifyCertAliasRequest) (response *ModifyCertAliasResponse, err error) {
	if request == nil {
		request = NewModifyCertAliasRequest()
	}
	response = NewModifyCertAliasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIspInfoRequest() (request *DescribeIspInfoRequest) {
	request = &DescribeIspInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeIspInfo")
	return
}

func NewDescribeIspInfoResponse() (response *DescribeIspInfoResponse) {
	response = &DescribeIspInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前地域所支持的运营商信息
func (c *Client) DescribeIspInfo(request *DescribeIspInfoRequest) (response *DescribeIspInfoResponse, err error) {
	if request == nil {
		request = NewDescribeIspInfoRequest()
	}
	response = NewDescribeIspInfoResponse()
	err = c.Send(request, response)
	return
}

func NewManualRewriteRequest() (request *ManualRewriteRequest) {
	request = &ManualRewriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ManualRewrite")
	return
}

func NewManualRewriteResponse() (response *ManualRewriteResponse) {
	response = &ManualRewriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户手动配置原访问地址和重定向地址，系统自动将原访问地址的请求重定向至对应路径的目的地址。同一域名下可以配置多条路径作为重定向策略，实现http/https之间请求的自动跳转。设置重定向时，需满足如下约束条件：若A已经重定向至B，则A不能再重定向至C（除非先删除老的重定向关系，再建立新的重定向关系），B不能重定向至任何其它地址。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) ManualRewrite(request *ManualRewriteRequest) (response *ManualRewriteResponse, err error) {
	if request == nil {
		request = NewManualRewriteRequest()
	}
	response = NewManualRewriteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMasterZonesRequest() (request *DescribeMasterZonesRequest) {
	request = &DescribeMasterZonesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeMasterZones")
	return
}

func NewDescribeMasterZonesResponse() (response *DescribeMasterZonesResponse) {
	response = &DescribeMasterZonesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询一个地域的可用区列表
func (c *Client) DescribeMasterZones(request *DescribeMasterZonesRequest) (response *DescribeMasterZonesResponse, err error) {
	if request == nil {
		request = NewDescribeMasterZonesRequest()
	}
	response = NewDescribeMasterZonesResponse()
	err = c.Send(request, response)
	return
}

func NewDeregisterTargetsRequest() (request *DeregisterTargetsRequest) {
	request = &DeregisterTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeregisterTargets")
	return
}

func NewDeregisterTargetsResponse() (response *DeregisterTargetsResponse) {
	response = &DeregisterTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeregisterTargets 接口用来将一台或多台后端服务从负载均衡的监听器或转发规则上解绑，对于四层监听器，只需指定监听器ID即可，对于七层监听器，还需通过LocationId或Domain+Url指定转发规则。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) DeregisterTargets(request *DeregisterTargetsRequest) (response *DeregisterTargetsResponse, err error) {
	if request == nil {
		request = NewDeregisterTargetsRequest()
	}
	response = NewDeregisterTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
	request = &DescribeTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeTaskStatus")
	return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
	response = &DescribeTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于查询异步任务的执行状态，对于非查询类的接口（创建/删除负载均衡实例、监听器、规则以及绑定或解绑后端服务等），在接口调用成功后，都需要使用本接口查询任务最终是否执行成功。
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTaskStatusRequest()
	}
	response = NewDescribeTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoadBalancerListByCertIdRequest() (request *DescribeLoadBalancerListByCertIdRequest) {
	request = &DescribeLoadBalancerListByCertIdRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancerListByCertId")
	return
}

func NewDescribeLoadBalancerListByCertIdResponse() (response *DescribeLoadBalancerListByCertIdResponse) {
	response = &DescribeLoadBalancerListByCertIdResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据证书ID查询其在一个地域中所关联到负载均衡实例列表
func (c *Client) DescribeLoadBalancerListByCertId(request *DescribeLoadBalancerListByCertIdRequest) (response *DescribeLoadBalancerListByCertIdResponse, err error) {
	if request == nil {
		request = NewDescribeLoadBalancerListByCertIdRequest()
	}
	response = NewDescribeLoadBalancerListByCertIdResponse()
	err = c.Send(request, response)
	return
}

func NewModifyLoadBalancerAttributesRequest() (request *ModifyLoadBalancerAttributesRequest) {
	request = &ModifyLoadBalancerAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyLoadBalancerAttributes")
	return
}

func NewModifyLoadBalancerAttributesResponse() (response *ModifyLoadBalancerAttributesResponse) {
	response = &ModifyLoadBalancerAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改负载均衡实例的属性。支持修改负载均衡实例的名称、设置负载均衡的跨域属性。
func (c *Client) ModifyLoadBalancerAttributes(request *ModifyLoadBalancerAttributesRequest) (response *ModifyLoadBalancerAttributesResponse, err error) {
	if request == nil {
		request = NewModifyLoadBalancerAttributesRequest()
	}
	response = NewModifyLoadBalancerAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewSetSecurityGroupForLoadbalancersRequest() (request *SetSecurityGroupForLoadbalancersRequest) {
	request = &SetSecurityGroupForLoadbalancersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetSecurityGroupForLoadbalancers")
	return
}

func NewSetSecurityGroupForLoadbalancersResponse() (response *SetSecurityGroupForLoadbalancersResponse) {
	response = &SetSecurityGroupForLoadbalancersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 绑定或解绑一个安全组到多个公网负载均衡实例。注意：内网负载均衡不支持绑定安全组。
func (c *Client) SetSecurityGroupForLoadbalancers(request *SetSecurityGroupForLoadbalancersRequest) (response *SetSecurityGroupForLoadbalancersResponse, err error) {
	if request == nil {
		request = NewSetSecurityGroupForLoadbalancersRequest()
	}
	response = NewSetSecurityGroupForLoadbalancersResponse()
	err = c.Send(request, response)
	return
}

func NewSetSubUinQuotasRequest() (request *SetSubUinQuotasRequest) {
	request = &SetSubUinQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetSubUinQuotas")
	return
}

func NewSetSubUinQuotasResponse() (response *SetSubUinQuotasResponse) {
	response = &SetSubUinQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过主租户账号设置子账号clb实例相关配额
func (c *Client) SetSubUinQuotas(request *SetSubUinQuotasRequest) (response *SetSubUinQuotasResponse, err error) {
	if request == nil {
		request = NewSetSubUinQuotasRequest()
	}
	response = NewSetSubUinQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTargetHealthRequest() (request *DescribeTargetHealthRequest) {
	request = &DescribeTargetHealthRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetHealth")
	return
}

func NewDescribeTargetHealthResponse() (response *DescribeTargetHealthResponse) {
	response = &DescribeTargetHealthResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTargetHealth 接口用来获取负载均衡后端服务的健康检查结果，不支持传统型负载均衡。
func (c *Client) DescribeTargetHealth(request *DescribeTargetHealthRequest) (response *DescribeTargetHealthResponse, err error) {
	if request == nil {
		request = NewDescribeTargetHealthRequest()
	}
	response = NewDescribeTargetHealthResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLBActionLimitRequest() (request *DescribeLBActionLimitRequest) {
	request = &DescribeLBActionLimitRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeLBActionLimit")
	return
}

func NewDescribeLBActionLimitResponse() (response *DescribeLBActionLimitResponse) {
	response = &DescribeLBActionLimitResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于查询LB相关操作配额
func (c *Client) DescribeLBActionLimit(request *DescribeLBActionLimitRequest) (response *DescribeLBActionLimitResponse, err error) {
	if request == nil {
		request = NewDescribeLBActionLimitRequest()
	}
	response = NewDescribeLBActionLimitResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeWhiteListSupportRequest() (request *DescribeWhiteListSupportRequest) {
	request = &DescribeWhiteListSupportRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeWhiteListSupport")
	return
}

func NewDescribeWhiteListSupportResponse() (response *DescribeWhiteListSupportResponse) {
	response = &DescribeWhiteListSupportResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询当前用户的WhiteList的白名单信息
func (c *Client) DescribeWhiteListSupport(request *DescribeWhiteListSupportRequest) (response *DescribeWhiteListSupportResponse, err error) {
	if request == nil {
		request = NewDescribeWhiteListSupportRequest()
	}
	response = NewDescribeWhiteListSupportResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomizedConfigAssociateListRequest() (request *DescribeCustomizedConfigAssociateListRequest) {
	request = &DescribeCustomizedConfigAssociateListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeCustomizedConfigAssociateList")
	return
}

func NewDescribeCustomizedConfigAssociateListResponse() (response *DescribeCustomizedConfigAssociateListResponse) {
	response = &DescribeCustomizedConfigAssociateListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取配置绑定的 server 或 location，如果 domain 存在，结果将根据 domain 过滤。或拉取配置绑定的 loadbalancer。
func (c *Client) DescribeCustomizedConfigAssociateList(request *DescribeCustomizedConfigAssociateListRequest) (response *DescribeCustomizedConfigAssociateListResponse, err error) {
	if request == nil {
		request = NewDescribeCustomizedConfigAssociateListRequest()
	}
	response = NewDescribeCustomizedConfigAssociateListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoadBalancersForVpcRequest() (request *DescribeLoadBalancersForVpcRequest) {
	request = &DescribeLoadBalancersForVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancersForVpc")
	return
}

func NewDescribeLoadBalancersForVpcResponse() (response *DescribeLoadBalancersForVpcResponse) {
	response = &DescribeLoadBalancersForVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 提供为VPC控制台用，查询负载均衡实例（部分信息没有查）。
func (c *Client) DescribeLoadBalancersForVpc(request *DescribeLoadBalancersForVpcRequest) (response *DescribeLoadBalancersForVpcResponse, err error) {
	if request == nil {
		request = NewDescribeLoadBalancersForVpcRequest()
	}
	response = NewDescribeLoadBalancersForVpcResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppIdLabelRequest() (request *DescribeAppIdLabelRequest) {
	request = &DescribeAppIdLabelRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeAppIdLabel")
	return
}

func NewDescribeAppIdLabelResponse() (response *DescribeAppIdLabelResponse) {
	response = &DescribeAppIdLabelResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户和绑定的标签
func (c *Client) DescribeAppIdLabel(request *DescribeAppIdLabelRequest) (response *DescribeAppIdLabelResponse, err error) {
	if request == nil {
		request = NewDescribeAppIdLabelRequest()
	}
	response = NewDescribeAppIdLabelResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTargetGroupsRequest() (request *DescribeTargetGroupsRequest) {
	request = &DescribeTargetGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetGroups")
	return
}

func NewDescribeTargetGroupsResponse() (response *DescribeTargetGroupsResponse) {
	response = &DescribeTargetGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询目标组信息
func (c *Client) DescribeTargetGroups(request *DescribeTargetGroupsRequest) (response *DescribeTargetGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeTargetGroupsRequest()
	}
	response = NewDescribeTargetGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubUinQuotasRequest() (request *DescribeSubUinQuotasRequest) {
	request = &DescribeSubUinQuotasRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeSubUinQuotas")
	return
}

func NewDescribeSubUinQuotasResponse() (response *DescribeSubUinQuotasResponse) {
	response = &DescribeSubUinQuotasResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 主租户查询子账号配额,若无子账号返回主账号配额
func (c *Client) DescribeSubUinQuotas(request *DescribeSubUinQuotasRequest) (response *DescribeSubUinQuotasResponse, err error) {
	if request == nil {
		request = NewDescribeSubUinQuotasRequest()
	}
	response = NewDescribeSubUinQuotasResponse()
	err = c.Send(request, response)
	return
}

func NewCreateListenerRequest() (request *CreateListenerRequest) {
	request = &CreateListenerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateListener")
	return
}

func NewCreateListenerResponse() (response *CreateListenerResponse) {
	response = &CreateListenerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 在一个负载均衡实例下创建监听器。
// 本接口为异步接口，接口返回成功后，需以返回的 RequestId 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
func (c *Client) CreateListener(request *CreateListenerRequest) (response *CreateListenerResponse, err error) {
	if request == nil {
		request = NewCreateListenerRequest()
	}
	response = NewCreateListenerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoadBalancersRequest() (request *DescribeLoadBalancersRequest) {
	request = &DescribeLoadBalancersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancers")
	return
}

func NewDescribeLoadBalancersResponse() (response *DescribeLoadBalancersResponse) {
	response = &DescribeLoadBalancersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询一个地域的负载均衡实例列表
//
func (c *Client) DescribeLoadBalancers(request *DescribeLoadBalancersRequest) (response *DescribeLoadBalancersResponse, err error) {
	if request == nil {
		request = NewDescribeLoadBalancersRequest()
	}
	response = NewDescribeLoadBalancersResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRuleRequest() (request *ModifyRuleRequest) {
	request = &ModifyRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyRule")
	return
}

func NewModifyRuleResponse() (response *ModifyRuleResponse) {
	response = &ModifyRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyRule 接口用来修改负载均衡七层监听器下的转发规则的各项属性，包括转发路径、健康检查属性、转发策略等。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) ModifyRule(request *ModifyRuleRequest) (response *ModifyRuleResponse, err error) {
	if request == nil {
		request = NewModifyRuleRequest()
	}
	response = NewModifyRuleResponse()
	err = c.Send(request, response)
	return
}

func NewSetLoadBalancerLogRequest() (request *SetLoadBalancerLogRequest) {
	request = &SetLoadBalancerLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetLoadBalancerLog")
	return
}

func NewSetLoadBalancerLogResponse() (response *SetLoadBalancerLogResponse) {
	response = &SetLoadBalancerLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加、删除、更新负载均衡日志访问，该功能目前仅支持以下地域：北京、上海、广州、香港、上海金融、深圳金融。
func (c *Client) SetLoadBalancerLog(request *SetLoadBalancerLogRequest) (response *SetLoadBalancerLogResponse, err error) {
	if request == nil {
		request = NewSetLoadBalancerLogRequest()
	}
	response = NewSetLoadBalancerLogResponse()
	err = c.Send(request, response)
	return
}

func NewQueryBlockIPTaskRequest() (request *QueryBlockIPTaskRequest) {
	request = &QueryBlockIPTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "QueryBlockIPTask")
	return
}

func NewQueryBlockIPTaskResponse() (response *QueryBlockIPTaskResponse) {
	response = &QueryBlockIPTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询封禁IP（黑名单）异步任务的执行状态。（接口灰度中，如需使用请提工单）
func (c *Client) QueryBlockIPTask(request *QueryBlockIPTaskRequest) (response *QueryBlockIPTaskResponse, err error) {
	if request == nil {
		request = NewQueryBlockIPTaskRequest()
	}
	response = NewQueryBlockIPTaskResponse()
	err = c.Send(request, response)
	return
}

func NewAssociateCustomizedConfigRequest() (request *AssociateCustomizedConfigRequest) {
	request = &AssociateCustomizedConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "AssociateCustomizedConfig")
	return
}

func NewAssociateCustomizedConfigResponse() (response *AssociateCustomizedConfigResponse) {
	response = &AssociateCustomizedConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 关联配置到server或location，根据配置类型关联到server或location。
func (c *Client) AssociateCustomizedConfig(request *AssociateCustomizedConfigRequest) (response *AssociateCustomizedConfigResponse, err error) {
	if request == nil {
		request = NewAssociateCustomizedConfigRequest()
	}
	response = NewAssociateCustomizedConfigResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceCreateLoadBalancerRequest() (request *InquiryPriceCreateLoadBalancerRequest) {
	request = &InquiryPriceCreateLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "InquiryPriceCreateLoadBalancer")
	return
}

func NewInquiryPriceCreateLoadBalancerResponse() (response *InquiryPriceCreateLoadBalancerResponse) {
	response = &InquiryPriceCreateLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询创建负载均衡的价格
func (c *Client) InquiryPriceCreateLoadBalancer(request *InquiryPriceCreateLoadBalancerRequest) (response *InquiryPriceCreateLoadBalancerResponse, err error) {
	if request == nil {
		request = NewInquiryPriceCreateLoadBalancerRequest()
	}
	response = NewInquiryPriceCreateLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewBatchModifyTargetWeightRequest() (request *BatchModifyTargetWeightRequest) {
	request = &BatchModifyTargetWeightRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "BatchModifyTargetWeight")
	return
}

func NewBatchModifyTargetWeightResponse() (response *BatchModifyTargetWeightResponse) {
	response = &BatchModifyTargetWeightResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(BatchModifyTargetWeight)用于批量修改负载均衡监听器绑定的后端机器的转发权重，支持负载均衡的4层和7层监听器；不支持传统型负载均衡。
// 本接口为异步接口，本接口返回成功后需以返回的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
func (c *Client) BatchModifyTargetWeight(request *BatchModifyTargetWeightRequest) (response *BatchModifyTargetWeightResponse, err error) {
	if request == nil {
		request = NewBatchModifyTargetWeightRequest()
	}
	response = NewBatchModifyTargetWeightResponse()
	err = c.Send(request, response)
	return
}

func NewCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateLoadBalancer")
	return
}

func NewCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口(CreateLoadBalancer)用来创建负载均衡实例（本接口只支持购买按量计费的负载均衡，包年包月的负载均衡请通过控制台购买）。为了使用负载均衡服务，您必须购买一个或多个负载均衡实例。成功调用该接口后，会返回负载均衡实例的唯一 ID。负载均衡实例的类型分为：公网、内网。详情可参考产品说明中的产品类型。
// 注意：(1)指定可用区申请负载均衡、跨zone容灾(仅香港支持)【如果您需要体验该功能，请通过工单申请】；(2)目前只有北京、上海、广州支持IPv6；(3)一个账号在每个地域的默认购买配额为：公网100个，内网100个。
// 本接口为异步接口，接口成功返回后，可使用 DescribeLoadBalancers 接口查询负载均衡实例的状态（如创建中、正常），以确定是否创建成功。
func (c *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	if request == nil {
		request = NewCreateLoadBalancerRequest()
	}
	response = NewCreateLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewRegisterTargetsRequest() (request *RegisterTargetsRequest) {
	request = &RegisterTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "RegisterTargets")
	return
}

func NewRegisterTargetsResponse() (response *RegisterTargetsResponse) {
	response = &RegisterTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// RegisterTargets 接口用来将一台或多台后端服务绑定到负载均衡的监听器（或7层转发规则），在此之前您需要先行创建相关的4层监听器或7层转发规则。对于四层监听器（TCP、UDP），只需指定监听器ID即可，对于七层监听器（HTTP、HTTPS），还需通过LocationId或者Domain+Url指定转发规则。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) RegisterTargets(request *RegisterTargetsRequest) (response *RegisterTargetsResponse, err error) {
	if request == nil {
		request = NewRegisterTargetsRequest()
	}
	response = NewRegisterTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTargetPortRequest() (request *ModifyTargetPortRequest) {
	request = &ModifyTargetPortRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetPort")
	return
}

func NewModifyTargetPortResponse() (response *ModifyTargetPortResponse) {
	response = &ModifyTargetPortResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyTargetPort接口用于修改监听器绑定的后端服务的端口。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) ModifyTargetPort(request *ModifyTargetPortRequest) (response *ModifyTargetPortResponse, err error) {
	if request == nil {
		request = NewModifyTargetPortRequest()
	}
	response = NewModifyTargetPortResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTargetGroupListRequest() (request *DescribeTargetGroupListRequest) {
	request = &DescribeTargetGroupListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetGroupList")
	return
}

func NewDescribeTargetGroupListResponse() (response *DescribeTargetGroupListResponse) {
	response = &DescribeTargetGroupListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取目标组列表
func (c *Client) DescribeTargetGroupList(request *DescribeTargetGroupListRequest) (response *DescribeTargetGroupListResponse, err error) {
	if request == nil {
		request = NewDescribeTargetGroupListRequest()
	}
	response = NewDescribeTargetGroupListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTargetWeightRequest() (request *ModifyTargetWeightRequest) {
	request = &ModifyTargetWeightRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyTargetWeight")
	return
}

func NewModifyTargetWeightResponse() (response *ModifyTargetWeightResponse) {
	response = &ModifyTargetWeightResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyTargetWeight 接口用于修改负载均衡绑定的后端服务的转发权重。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) ModifyTargetWeight(request *ModifyTargetWeightRequest) (response *ModifyTargetWeightResponse, err error) {
	if request == nil {
		request = NewModifyTargetWeightRequest()
	}
	response = NewModifyTargetWeightResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainAttributesRequest() (request *ModifyDomainAttributesRequest) {
	request = &ModifyDomainAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyDomainAttributes")
	return
}

func NewModifyDomainAttributesResponse() (response *ModifyDomainAttributesResponse) {
	response = &ModifyDomainAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyDomainAttributes接口用于修改负载均衡7层监听器转发规则的域名级别属性，如修改域名、修改DefaultServer、开启/关闭Http2、修改证书。
// 本接口为异步接口，本接口返回成功后，需以返回的RequestId为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) ModifyDomainAttributes(request *ModifyDomainAttributesRequest) (response *ModifyDomainAttributesResponse, err error) {
	if request == nil {
		request = NewModifyDomainAttributesRequest()
	}
	response = NewModifyDomainAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewSetLoadBalancerClsLogRequest() (request *SetLoadBalancerClsLogRequest) {
	request = &SetLoadBalancerClsLogRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetLoadBalancerClsLog")
	return
}

func NewSetLoadBalancerClsLogResponse() (response *SetLoadBalancerClsLogResponse) {
	response = &SetLoadBalancerClsLogResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加、删除、更新负载均衡的日志服务(CLS)主题
func (c *Client) SetLoadBalancerClsLog(request *SetLoadBalancerClsLogRequest) (response *SetLoadBalancerClsLogResponse, err error) {
	if request == nil {
		request = NewSetLoadBalancerClsLogRequest()
	}
	response = NewSetLoadBalancerClsLogResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClassicalLBHealthStatusRequest() (request *DescribeClassicalLBHealthStatusRequest) {
	request = &DescribeClassicalLBHealthStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeClassicalLBHealthStatus")
	return
}

func NewDescribeClassicalLBHealthStatusResponse() (response *DescribeClassicalLBHealthStatusResponse) {
	response = &DescribeClassicalLBHealthStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeClassicalLBHealthStatus用于获取传统型负载均衡后端的健康状态
func (c *Client) DescribeClassicalLBHealthStatus(request *DescribeClassicalLBHealthStatusRequest) (response *DescribeClassicalLBHealthStatusResponse, err error) {
	if request == nil {
		request = NewDescribeClassicalLBHealthStatusRequest()
	}
	response = NewDescribeClassicalLBHealthStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCertsRequest() (request *DescribeCertsRequest) {
	request = &DescribeCertsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeCerts")
	return
}

func NewDescribeCertsResponse() (response *DescribeCertsResponse) {
	response = &DescribeCertsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询可用证书列表
func (c *Client) DescribeCerts(request *DescribeCertsRequest) (response *DescribeCertsResponse, err error) {
	if request == nil {
		request = NewDescribeCertsRequest()
	}
	response = NewDescribeCertsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCertRequest() (request *DeleteCertRequest) {
	request = &DeleteCertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteCert")
	return
}

func NewDeleteCertResponse() (response *DeleteCertResponse) {
	response = &DeleteCertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 租户端CLB删除证书
func (c *Client) DeleteCert(request *DeleteCertRequest) (response *DeleteCertResponse, err error) {
	if request == nil {
		request = NewDeleteCertRequest()
	}
	response = NewDeleteCertResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteLoadBalancerRequest() (request *DeleteLoadBalancerRequest) {
	request = &DeleteLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteLoadBalancer")
	return
}

func NewDeleteLoadBalancerResponse() (response *DeleteLoadBalancerResponse) {
	response = &DeleteLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteLoadBalancer 接口用以删除指定的一个或多个负载均衡实例。
// 本接口为异步接口，接口返回成功后，需以返回的 RequestId 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
func (c *Client) DeleteLoadBalancer(request *DeleteLoadBalancerRequest) (response *DeleteLoadBalancerResponse, err error) {
	if request == nil {
		request = NewDeleteLoadBalancerRequest()
	}
	response = NewDeleteLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewSetCustomizedConfigForLoadBalancerRequest() (request *SetCustomizedConfigForLoadBalancerRequest) {
	request = &SetCustomizedConfigForLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetCustomizedConfigForLoadBalancer")
	return
}

func NewSetCustomizedConfigForLoadBalancerResponse() (response *SetCustomizedConfigForLoadBalancerResponse) {
	response = &SetCustomizedConfigForLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 负载均衡维度的个性化配置相关操作：创建、删除、修改、绑定、解绑
func (c *Client) SetCustomizedConfigForLoadBalancer(request *SetCustomizedConfigForLoadBalancerRequest) (response *SetCustomizedConfigForLoadBalancerResponse, err error) {
	if request == nil {
		request = NewSetCustomizedConfigForLoadBalancerRequest()
	}
	response = NewSetCustomizedConfigForLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRuleRequest() (request *CreateRuleRequest) {
	request = &CreateRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateRule")
	return
}

func NewCreateRuleResponse() (response *CreateRuleResponse) {
	response = &CreateRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateRule 接口用于在一个已存在的负载均衡七层监听器下创建转发规则，七层监听器中，后端服务必须绑定到规则上而非监听器上。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) CreateRule(request *CreateRuleRequest) (response *CreateRuleResponse, err error) {
	if request == nil {
		request = NewCreateRuleRequest()
	}
	response = NewCreateRuleResponse()
	err = c.Send(request, response)
	return
}

func NewReplaceCertRequest() (request *ReplaceCertRequest) {
	request = &ReplaceCertRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ReplaceCert")
	return
}

func NewReplaceCertResponse() (response *ReplaceCertResponse) {
	response = &ReplaceCertResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 使用新证书替换原有证书(注意：若原证书在多个region上使用，则同时替换原证书绑定的规则)
func (c *Client) ReplaceCert(request *ReplaceCertRequest) (response *ReplaceCertResponse, err error) {
	if request == nil {
		request = NewReplaceCertRequest()
	}
	response = NewReplaceCertResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRewriteRequest() (request *DeleteRewriteRequest) {
	request = &DeleteRewriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteRewrite")
	return
}

func NewDeleteRewriteResponse() (response *DeleteRewriteResponse) {
	response = &DeleteRewriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DeleteRewrite 接口支持删除指定转发规则之间的重定向关系。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) DeleteRewrite(request *DeleteRewriteRequest) (response *DeleteRewriteResponse, err error) {
	if request == nil {
		request = NewDeleteRewriteRequest()
	}
	response = NewDeleteRewriteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSetInnerNameRequest() (request *DescribeSetInnerNameRequest) {
	request = &DescribeSetInnerNameRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeSetInnerName")
	return
}

func NewDescribeSetInnerNameResponse() (response *DescribeSetInnerNameResponse) {
	response = &DescribeSetInnerNameResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据集群对外英文名、appid查询中文名
func (c *Client) DescribeSetInnerName(request *DescribeSetInnerNameRequest) (response *DescribeSetInnerNameResponse, err error) {
	if request == nil {
		request = NewDescribeSetInnerNameRequest()
	}
	response = NewDescribeSetInnerNameResponse()
	err = c.Send(request, response)
	return
}

func NewInquiryPriceModifyLoadBalancerRequest() (request *InquiryPriceModifyLoadBalancerRequest) {
	request = &InquiryPriceModifyLoadBalancerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "InquiryPriceModifyLoadBalancer")
	return
}

func NewInquiryPriceModifyLoadBalancerResponse() (response *InquiryPriceModifyLoadBalancerResponse) {
	response = &InquiryPriceModifyLoadBalancerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// InquiryPriceModifyLoadBalancer接口修改负载均衡配置询价。
func (c *Client) InquiryPriceModifyLoadBalancer(request *InquiryPriceModifyLoadBalancerRequest) (response *InquiryPriceModifyLoadBalancerResponse, err error) {
	if request == nil {
		request = NewInquiryPriceModifyLoadBalancerRequest()
	}
	response = NewInquiryPriceModifyLoadBalancerResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteListenerRequest() (request *DeleteListenerRequest) {
	request = &DeleteListenerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DeleteListener")
	return
}

func NewDeleteListenerResponse() (response *DeleteListenerResponse) {
	response = &DeleteListenerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用来删除负载均衡实例下的监听器（四层和七层）。
// 本接口为异步接口，接口返回成功后，需以得到的 RequestID 为入参，调用 DescribeTaskStatus 接口查询本次任务是否成功。
func (c *Client) DeleteListener(request *DeleteListenerRequest) (response *DeleteListenerResponse, err error) {
	if request == nil {
		request = NewDeleteListenerRequest()
	}
	response = NewDeleteListenerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLBListenersRequest() (request *DescribeLBListenersRequest) {
	request = &DescribeLBListenersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeLBListeners")
	return
}

func NewDescribeLBListenersResponse() (response *DescribeLBListenersResponse) {
	response = &DescribeLBListenersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询后端云主机或弹性网卡绑定的负载均衡，支持弹性网卡和cvm查询。
func (c *Client) DescribeLBListeners(request *DescribeLBListenersRequest) (response *DescribeLBListenersResponse, err error) {
	if request == nil {
		request = NewDescribeLBListenersRequest()
	}
	response = NewDescribeLBListenersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCustomizedConfigListRequest() (request *DescribeCustomizedConfigListRequest) {
	request = &DescribeCustomizedConfigListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeCustomizedConfigList")
	return
}

func NewDescribeCustomizedConfigListResponse() (response *DescribeCustomizedConfigListResponse) {
	response = &DescribeCustomizedConfigListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 拉取个性化配置列表，返回用户 AppId 下指定类型的配置。
func (c *Client) DescribeCustomizedConfigList(request *DescribeCustomizedConfigListRequest) (response *DescribeCustomizedConfigListResponse, err error) {
	if request == nil {
		request = NewDescribeCustomizedConfigListRequest()
	}
	response = NewDescribeCustomizedConfigListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTargetCountForLoadBalancersRequest() (request *DescribeTargetCountForLoadBalancersRequest) {
	request = &DescribeTargetCountForLoadBalancersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeTargetCountForLoadBalancers")
	return
}

func NewDescribeTargetCountForLoadBalancersResponse() (response *DescribeTargetCountForLoadBalancersResponse) {
	response = &DescribeTargetCountForLoadBalancersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询负载均衡实例绑定的后端服务总数
func (c *Client) DescribeTargetCountForLoadBalancers(request *DescribeTargetCountForLoadBalancersRequest) (response *DescribeTargetCountForLoadBalancersResponse, err error) {
	if request == nil {
		request = NewDescribeTargetCountForLoadBalancersRequest()
	}
	response = NewDescribeTargetCountForLoadBalancersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSingleIspRequest() (request *DescribeSingleIspRequest) {
	request = &DescribeSingleIspRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeSingleIsp")
	return
}

func NewDescribeSingleIspResponse() (response *DescribeSingleIspResponse) {
	response = &DescribeSingleIspResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 返回某个地域所支持的运营商列表。
func (c *Client) DescribeSingleIsp(request *DescribeSingleIspRequest) (response *DescribeSingleIspResponse, err error) {
	if request == nil {
		request = NewDescribeSingleIspRequest()
	}
	response = NewDescribeSingleIspResponse()
	err = c.Send(request, response)
	return
}

func NewSetLoadBalancerSecurityGroupsRequest() (request *SetLoadBalancerSecurityGroupsRequest) {
	request = &SetLoadBalancerSecurityGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "SetLoadBalancerSecurityGroups")
	return
}

func NewSetLoadBalancerSecurityGroupsResponse() (response *SetLoadBalancerSecurityGroupsResponse) {
	response = &SetLoadBalancerSecurityGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// SetLoadBalancerSecurityGroups 接口支持对一个公网负载均衡实例执行设置（绑定、解绑）安全组操作。查询一个负载均衡实例目前已绑定的安全组，可使用 DescribeLoadBalancers 接口。本接口是set语义，
// 绑定操作时，入参需要传入负载均衡实例要绑定的所有安全组（已绑定的+新增绑定的）。
// 解绑操作时，入参需要传入负载均衡实例执行解绑后所绑定的所有安全组；如果要解绑所有安全组，可不传此参数，或传入空数组。注意：内网负载均衡不支持绑定安全组。
func (c *Client) SetLoadBalancerSecurityGroups(request *SetLoadBalancerSecurityGroupsRequest) (response *SetLoadBalancerSecurityGroupsResponse, err error) {
	if request == nil {
		request = NewSetLoadBalancerSecurityGroupsRequest()
	}
	response = NewSetLoadBalancerSecurityGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewBatchDeregisterTargetsRequest() (request *BatchDeregisterTargetsRequest) {
	request = &BatchDeregisterTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "BatchDeregisterTargets")
	return
}

func NewBatchDeregisterTargetsResponse() (response *BatchDeregisterTargetsResponse) {
	response = &BatchDeregisterTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量解绑四七层后端服务。批量解绑的资源数量上限为500。只支持VPC网络负载均衡。
func (c *Client) BatchDeregisterTargets(request *BatchDeregisterTargetsRequest) (response *BatchDeregisterTargetsResponse, err error) {
	if request == nil {
		request = NewBatchDeregisterTargetsRequest()
	}
	response = NewBatchDeregisterTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTargetsRequest() (request *DescribeTargetsRequest) {
	request = &DescribeTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeTargets")
	return
}

func NewDescribeTargetsResponse() (response *DescribeTargetsResponse) {
	response = &DescribeTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeTargets 接口用来查询负载均衡实例的某些监听器绑定的后端服务列表。
func (c *Client) DescribeTargets(request *DescribeTargetsRequest) (response *DescribeTargetsResponse, err error) {
	if request == nil {
		request = NewDescribeTargetsRequest()
	}
	response = NewDescribeTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewTestEzioRequest() (request *TestEzioRequest) {
	request = &TestEzioRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "TestEzio")
	return
}

func NewTestEzioResponse() (response *TestEzioResponse) {
	response = &TestEzioResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 测试
func (c *Client) TestEzio(request *TestEzioRequest) (response *TestEzioResponse, err error) {
	if request == nil {
		request = NewTestEzioRequest()
	}
	response = NewTestEzioResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBlockIPListRequest() (request *DescribeBlockIPListRequest) {
	request = &DescribeBlockIPListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeBlockIPList")
	return
}

func NewDescribeBlockIPListResponse() (response *DescribeBlockIPListResponse) {
	response = &DescribeBlockIPListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询一个负载均衡所封禁的IP列表（黑名单）。（接口灰度中，如需使用请提工单）
func (c *Client) DescribeBlockIPList(request *DescribeBlockIPListRequest) (response *DescribeBlockIPListResponse, err error) {
	if request == nil {
		request = NewDescribeBlockIPListRequest()
	}
	response = NewDescribeBlockIPListResponse()
	err = c.Send(request, response)
	return
}

func NewModifyListenerRequest() (request *ModifyListenerRequest) {
	request = &ModifyListenerRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyListener")
	return
}

func NewModifyListenerResponse() (response *ModifyListenerResponse) {
	response = &ModifyListenerResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyListener接口用来修改负载均衡监听器的属性，包括监听器名称、健康检查参数、证书信息、转发策略等。本接口不支持传统型负载均衡。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) ModifyListener(request *ModifyListenerRequest) (response *ModifyListenerResponse, err error) {
	if request == nil {
		request = NewModifyListenerRequest()
	}
	response = NewModifyListenerResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeLoadBalancerCountRequest() (request *DescribeLoadBalancerCountRequest) {
	request = &DescribeLoadBalancerCountRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeLoadBalancerCount")
	return
}

func NewDescribeLoadBalancerCountResponse() (response *DescribeLoadBalancerCountResponse) {
	response = &DescribeLoadBalancerCountResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询一个账户在一个地域的负载均衡实例总数
func (c *Client) DescribeLoadBalancerCount(request *DescribeLoadBalancerCountRequest) (response *DescribeLoadBalancerCountResponse, err error) {
	if request == nil {
		request = NewDescribeLoadBalancerCountRequest()
	}
	response = NewDescribeLoadBalancerCountResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDnatRequest() (request *ModifyDnatRequest) {
	request = &ModifyDnatRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyDnat")
	return
}

func NewModifyDnatResponse() (response *ModifyDnatResponse) {
	response = &ModifyDnatResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 打开或关闭clb直通功能
func (c *Client) ModifyDnat(request *ModifyDnatRequest) (response *ModifyDnatResponse, err error) {
	if request == nil {
		request = NewModifyDnatRequest()
	}
	response = NewModifyDnatResponse()
	err = c.Send(request, response)
	return
}

func NewModifyDomainRequest() (request *ModifyDomainRequest) {
	request = &ModifyDomainRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "ModifyDomain")
	return
}

func NewModifyDomainResponse() (response *ModifyDomainResponse) {
	response = &ModifyDomainResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// ModifyDomain接口用来修改负载均衡七层监听器下的域名。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) ModifyDomain(request *ModifyDomainRequest) (response *ModifyDomainResponse, err error) {
	if request == nil {
		request = NewModifyDomainRequest()
	}
	response = NewModifyDomainResponse()
	err = c.Send(request, response)
	return
}

func NewBatchRegisterTargetsRequest() (request *BatchRegisterTargetsRequest) {
	request = &BatchRegisterTargetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "BatchRegisterTargets")
	return
}

func NewBatchRegisterTargetsResponse() (response *BatchRegisterTargetsResponse) {
	response = &BatchRegisterTargetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量绑定虚拟主机或弹性网卡，支持跨域绑定，支持四层、七层（TCP、UDP、HTTP、HTTPS）协议绑定。批量绑定的资源数量上限为500。只支持VPC网络负载均衡。
func (c *Client) BatchRegisterTargets(request *BatchRegisterTargetsRequest) (response *BatchRegisterTargetsResponse, err error) {
	if request == nil {
		request = NewBatchRegisterTargetsRequest()
	}
	response = NewBatchRegisterTargetsResponse()
	err = c.Send(request, response)
	return
}

func NewAutoRewriteRequest() (request *AutoRewriteRequest) {
	request = &AutoRewriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "AutoRewrite")
	return
}

func NewAutoRewriteResponse() (response *AutoRewriteResponse) {
	response = &AutoRewriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户需要先创建出一个HTTPS:443监听器，并在其下创建转发规则。通过调用本接口，系统会自动创建出一个HTTP:80监听器（如果之前不存在），并在其下创建转发规则，与HTTPS:443监听器下的Domains（在入参中指定）对应。创建成功后可以通过HTTP:80地址自动跳转为HTTPS:443地址进行访问。
// 本接口为异步接口，本接口返回成功后需以返回的RequestID为入参，调用DescribeTaskStatus接口查询本次任务是否成功。
func (c *Client) AutoRewrite(request *AutoRewriteRequest) (response *AutoRewriteResponse, err error) {
	if request == nil {
		request = NewAutoRewriteRequest()
	}
	response = NewAutoRewriteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRewriteRequest() (request *DescribeRewriteRequest) {
	request = &DescribeRewriteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeRewrite")
	return
}

func NewDescribeRewriteResponse() (response *DescribeRewriteResponse) {
	response = &DescribeRewriteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeRewrite 接口可根据负载均衡实例ID，查询一个负载均衡实例下转发规则的重定向关系。如果不指定监听器ID或转发规则ID，则返回该负载均衡实例下的所有重定向关系。
func (c *Client) DescribeRewrite(request *DescribeRewriteRequest) (response *DescribeRewriteResponse, err error) {
	if request == nil {
		request = NewDescribeRewriteRequest()
	}
	response = NewDescribeRewriteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClsLogSetRequest() (request *DescribeClsLogSetRequest) {
	request = &DescribeClsLogSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "DescribeClsLogSet")
	return
}

func NewDescribeClsLogSetResponse() (response *DescribeClsLogSetResponse) {
	response = &DescribeClsLogSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeClsLogSet 查询CLB专有日志集
func (c *Client) DescribeClsLogSet(request *DescribeClsLogSetRequest) (response *DescribeClsLogSetResponse, err error) {
	if request == nil {
		request = NewDescribeClsLogSetRequest()
	}
	response = NewDescribeClsLogSetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClsLogSetRequest() (request *CreateClsLogSetRequest) {
	request = &CreateClsLogSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateClsLogSet")
	return
}

func NewCreateClsLogSetResponse() (response *CreateClsLogSetResponse) {
	response = &CreateClsLogSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateClsLogSet 创建CLB专有日志集，用于存储CLB的日志
func (c *Client) CreateClsLogSet(request *CreateClsLogSetRequest) (response *CreateClsLogSetResponse, err error) {
	if request == nil {
		request = NewCreateClsLogSetRequest()
	}
	response = NewCreateClsLogSetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("clb", APIVersion, "CreateTopic")
	return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
	response = &CreateTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CreateTopic 创建CLB日志主题
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	if request == nil {
		request = NewCreateTopicRequest()
	}
	response = NewCreateTopicResponse()
	err = c.Send(request, response)
	return
}

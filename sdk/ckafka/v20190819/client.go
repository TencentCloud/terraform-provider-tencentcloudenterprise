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

package v20190819

import (
	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2019-08-19"

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

func NewDeleteConnectorRequest() (request *DeleteConnectorRequest) {
	request = &DeleteConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DeleteConnector")
	return
}

func NewDeleteConnectorResponse() (response *DeleteConnectorResponse) {
	response = &DeleteConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Connector
func (c *Client) DeleteConnector(request *DeleteConnectorRequest) (response *DeleteConnectorResponse, err error) {
	if request == nil {
		request = NewDeleteConnectorRequest()
	}
	response = NewDeleteConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicRequest() (request *DescribeTopicRequest) {
	request = &DescribeTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopic")
	return
}

func NewDescribeTopicResponse() (response *DescribeTopicResponse) {
	response = &DescribeTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口请求域名：https://ckafka.tencentcloudapi.com
// 本接口（DescribeTopic）用于在用户获取消息队列 CKafka 实例的主题列表
func (c *Client) DescribeTopic(request *DescribeTopicRequest) (response *DescribeTopicResponse, err error) {
	if request == nil {
		request = NewDescribeTopicRequest()
	}
	response = NewDescribeTopicResponse()
	err = c.Send(request, response)
	return
}

func NewHourModifyRequest() (request *HourModifyRequest) {
	request = &HourModifyRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "HourModify")
	return
}

func NewHourModifyResponse() (response *HourModifyResponse) {
	response = &HourModifyResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改按量计费实例配置
func (c *Client) HourModify(request *HourModifyRequest) (response *HourModifyResponse, err error) {
	if request == nil {
		request = NewHourModifyRequest()
	}
	response = NewHourModifyResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicAttributesRequest() (request *DescribeTopicAttributesRequest) {
	request = &DescribeTopicAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicAttributes")
	return
}

func NewDescribeTopicAttributesResponse() (response *DescribeTopicAttributesResponse) {
	response = &DescribeTopicAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主题属性
//
func (c *Client) DescribeTopicAttributes(request *DescribeTopicAttributesRequest) (response *DescribeTopicAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeTopicAttributesRequest()
	}
	response = NewDescribeTopicAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConnectorRequest() (request *CreateConnectorRequest) {
	request = &CreateConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreateConnector")
	return
}

func NewCreateConnectorResponse() (response *CreateConnectorResponse) {
	response = &CreateConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建数据同步任务
//
func (c *Client) CreateConnector(request *CreateConnectorRequest) (response *CreateConnectorResponse, err error) {
	if request == nil {
		request = NewCreateConnectorRequest()
	}
	response = NewCreateConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeIfCommunityRequest() (request *DescribeIfCommunityRequest) {
	request = &DescribeIfCommunityRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeIfCommunity")
	return
}

func NewDescribeIfCommunityResponse() (response *DescribeIfCommunityResponse) {
	response = &DescribeIfCommunityResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询实例是否为社区版
func (c *Client) DescribeIfCommunity(request *DescribeIfCommunityRequest) (response *DescribeIfCommunityResponse, err error) {
	if request == nil {
		request = NewDescribeIfCommunityRequest()
	}
	response = NewDescribeIfCommunityResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConnectorRequest() (request *DescribeConnectorRequest) {
	request = &DescribeConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConnector")
	return
}

func NewDescribeConnectorResponse() (response *DescribeConnectorResponse) {
	response = &DescribeConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取数据同步任务列表
//
func (c *Client) DescribeConnector(request *DescribeConnectorRequest) (response *DescribeConnectorResponse, err error) {
	if request == nil {
		request = NewDescribeConnectorRequest()
	}
	response = NewDescribeConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTaskStatusRequest() (request *DescribeTaskStatusRequest) {
	request = &DescribeTaskStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTaskStatus")
	return
}

func NewDescribeTaskStatusResponse() (response *DescribeTaskStatusResponse) {
	response = &DescribeTaskStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询任务状态
func (c *Client) DescribeTaskStatus(request *DescribeTaskStatusRequest) (response *DescribeTaskStatusResponse, err error) {
	if request == nil {
		request = NewDescribeTaskStatusRequest()
	}
	response = NewDescribeTaskStatusResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeUserRequest() (request *DescribeUserRequest) {
	request = &DescribeUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeUser")
	return
}

func NewDescribeUserResponse() (response *DescribeUserResponse) {
	response = &DescribeUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户信息
func (c *Client) DescribeUser(request *DescribeUserRequest) (response *DescribeUserResponse, err error) {
	if request == nil {
		request = NewDescribeUserRequest()
	}
	response = NewDescribeUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstanceAttributesRequest() (request *DescribeInstanceAttributesRequest) {
	request = &DescribeInstanceAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstanceAttributes")
	return
}

func NewDescribeInstanceAttributesResponse() (response *DescribeInstanceAttributesResponse) {
	response = &DescribeInstanceAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例属性
func (c *Client) DescribeInstanceAttributes(request *DescribeInstanceAttributesRequest) (response *DescribeInstanceAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceAttributesRequest()
	}
	response = NewDescribeInstanceAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreateTopic")
	return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
	response = &CreateTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建ckafka主题
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	if request == nil {
		request = NewCreateTopicRequest()
	}
	response = NewCreateTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesRequest() (request *DescribeInstancesRequest) {
	request = &DescribeInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstances")
	return
}

func NewDescribeInstancesResponse() (response *DescribeInstancesResponse) {
	response = &DescribeInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口（DescribeInstance）用于在用户账户下获取消息队列 CKafka 实例列表
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesRequest()
	}
	response = NewDescribeInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewAuthorizeTokenRequest() (request *AuthorizeTokenRequest) {
	request = &AuthorizeTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "AuthorizeToken")
	return
}

func NewAuthorizeTokenResponse() (response *AuthorizeTokenResponse) {
	response = &AuthorizeTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// AuthorizeToken
func (c *Client) AuthorizeToken(request *AuthorizeTokenRequest) (response *AuthorizeTokenResponse, err error) {
	if request == nil {
		request = NewAuthorizeTokenRequest()
	}
	response = NewAuthorizeTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteInstanceRequest() (request *DeleteInstanceRequest) {
	request = &DeleteInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DeleteInstance")
	return
}

func NewDeleteInstanceResponse() (response *DeleteInstanceResponse) {
	response = &DeleteInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除实例
func (c *Client) DeleteInstance(request *DeleteInstanceRequest) (response *DeleteInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteInstanceRequest()
	}
	response = NewDeleteInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
	request = &DeleteRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DeleteRoute")
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

func NewDescribeInstanceDetailRequest() (request *DescribeInstanceDetailRequest) {
	request = &DescribeInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstanceDetail")
	return
}

func NewDescribeInstanceDetailResponse() (response *DescribeInstanceDetailResponse) {
	response = &DescribeInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户账户下获取实例列表详情
func (c *Client) DescribeInstanceDetail(request *DescribeInstanceDetailRequest) (response *DescribeInstanceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeInstanceDetailRequest()
	}
	response = NewDescribeInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyGroupOffsetsRequest() (request *ModifyGroupOffsetsRequest) {
	request = &ModifyGroupOffsetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "ModifyGroupOffsets")
	return
}

func NewModifyGroupOffsetsResponse() (response *ModifyGroupOffsetsResponse) {
	response = &ModifyGroupOffsetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置Groups 消费分组offset
func (c *Client) ModifyGroupOffsets(request *ModifyGroupOffsetsRequest) (response *ModifyGroupOffsetsResponse, err error) {
	if request == nil {
		request = NewModifyGroupOffsetsRequest()
	}
	response = NewModifyGroupOffsetsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateUserRequest() (request *CreateUserRequest) {
	request = &CreateUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreateUser")
	return
}

func NewCreateUserResponse() (response *CreateUserResponse) {
	response = &CreateUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加用户
func (c *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
	if request == nil {
		request = NewCreateUserRequest()
	}
	response = NewCreateUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupRequest() (request *DescribeGroupRequest) {
	request = &DescribeGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroup")
	return
}

func NewDescribeGroupResponse() (response *DescribeGroupResponse) {
	response = &DescribeGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 枚举消费分组(精简版)
func (c *Client) DescribeGroup(request *DescribeGroupRequest) (response *DescribeGroupResponse, err error) {
	if request == nil {
		request = NewDescribeGroupRequest()
	}
	response = NewDescribeGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicDetailRequest() (request *DescribeTopicDetailRequest) {
	request = &DescribeTopicDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeTopicDetail")
	return
}

func NewDescribeTopicDetailResponse() (response *DescribeTopicDetailResponse) {
	response = &DescribeTopicDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取主题列表详情（仅控制台调用）
func (c *Client) DescribeTopicDetail(request *DescribeTopicDetailRequest) (response *DescribeTopicDetailResponse, err error) {
	if request == nil {
		request = NewDescribeTopicDetailRequest()
	}
	response = NewDescribeTopicDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateInstanceTceRequest() (request *CreateInstanceTceRequest) {
	request = &CreateInstanceTceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreateInstanceTce")
	return
}

func NewCreateInstanceTceResponse() (response *CreateInstanceTceResponse) {
	response = &CreateInstanceTceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建tce实例
func (c *Client) CreateInstanceTce(request *CreateInstanceTceRequest) (response *CreateInstanceTceResponse, err error) {
	if request == nil {
		request = NewCreateInstanceTceRequest()
	}
	response = NewCreateInstanceTceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPasswordRequest() (request *ModifyPasswordRequest) {
	request = &ModifyPasswordRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "ModifyPassword")
	return
}

func NewModifyPasswordResponse() (response *ModifyPasswordResponse) {
	response = &ModifyPasswordResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改密码
func (c *Client) ModifyPassword(request *ModifyPasswordRequest) (response *ModifyPasswordResponse, err error) {
	if request == nil {
		request = NewModifyPasswordRequest()
	}
	response = NewModifyPasswordResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInstanceAttributesRequest() (request *ModifyInstanceAttributesRequest) {
	request = &ModifyInstanceAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "ModifyInstanceAttributes")
	return
}

func NewModifyInstanceAttributesResponse() (response *ModifyInstanceAttributesResponse) {
	response = &ModifyInstanceAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 设置实例属性
func (c *Client) ModifyInstanceAttributes(request *ModifyInstanceAttributesRequest) (response *ModifyInstanceAttributesResponse, err error) {
	if request == nil {
		request = NewModifyInstanceAttributesRequest()
	}
	response = NewModifyInstanceAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConnectorConfigsRequest() (request *DescribeConnectorConfigsRequest) {
	request = &DescribeConnectorConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConnectorConfigs")
	return
}

func NewDescribeConnectorConfigsResponse() (response *DescribeConnectorConfigsResponse) {
	response = &DescribeConnectorConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询connector配置
func (c *Client) DescribeConnectorConfigs(request *DescribeConnectorConfigsRequest) (response *DescribeConnectorConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeConnectorConfigsRequest()
	}
	response = NewDescribeConnectorConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAclRequest() (request *DeleteAclRequest) {
	request = &DeleteAclRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DeleteAcl")
	return
}

func NewDeleteAclResponse() (response *DeleteAclResponse) {
	response = &DeleteAclResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除ACL
func (c *Client) DeleteAcl(request *DeleteAclRequest) (response *DeleteAclResponse, err error) {
	if request == nil {
		request = NewDeleteAclRequest()
	}
	response = NewDeleteAclResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBrokersRequest() (request *DescribeBrokersRequest) {
	request = &DescribeBrokersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeBrokers")
	return
}

func NewDescribeBrokersResponse() (response *DescribeBrokersResponse) {
	response = &DescribeBrokersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于给运维提供枚举broker信息的接口
func (c *Client) DescribeBrokers(request *DescribeBrokersRequest) (response *DescribeBrokersResponse, err error) {
	if request == nil {
		request = NewDescribeBrokersRequest()
	}
	response = NewDescribeBrokersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConsumerGroupRequest() (request *DescribeConsumerGroupRequest) {
	request = &DescribeConsumerGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConsumerGroup")
	return
}

func NewDescribeConsumerGroupResponse() (response *DescribeConsumerGroupResponse) {
	response = &DescribeConsumerGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询消费分组信息
func (c *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (response *DescribeConsumerGroupResponse, err error) {
	if request == nil {
		request = NewDescribeConsumerGroupRequest()
	}
	response = NewDescribeConsumerGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupInfoRequest() (request *DescribeGroupInfoRequest) {
	request = &DescribeGroupInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroupInfo")
	return
}

func NewDescribeGroupInfoResponse() (response *DescribeGroupInfoResponse) {
	response = &DescribeGroupInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取消费分组信息
func (c *Client) DescribeGroupInfo(request *DescribeGroupInfoRequest) (response *DescribeGroupInfoResponse, err error) {
	if request == nil {
		request = NewDescribeGroupInfoRequest()
	}
	response = NewDescribeGroupInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteRequest() (request *DescribeRouteRequest) {
	request = &DescribeRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeRoute")
	return
}

func NewDescribeRouteResponse() (response *DescribeRouteResponse) {
	response = &DescribeRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看路由信息
func (c *Client) DescribeRoute(request *DescribeRouteRequest) (response *DescribeRouteResponse, err error) {
	if request == nil {
		request = NewDescribeRouteRequest()
	}
	response = NewDescribeRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopicIpWhiteListRequest() (request *DeleteTopicIpWhiteListRequest) {
	request = &DeleteTopicIpWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DeleteTopicIpWhiteList")
	return
}

func NewDeleteTopicIpWhiteListResponse() (response *DeleteTopicIpWhiteListResponse) {
	response = &DeleteTopicIpWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除主题IP白名单
func (c *Client) DeleteTopicIpWhiteList(request *DeleteTopicIpWhiteListRequest) (response *DeleteTopicIpWhiteListResponse, err error) {
	if request == nil {
		request = NewDeleteTopicIpWhiteListRequest()
	}
	response = NewDeleteTopicIpWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCkafkaTypeConfigsRequest() (request *DescribeCkafkaTypeConfigsRequest) {
	request = &DescribeCkafkaTypeConfigsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeCkafkaTypeConfigs")
	return
}

func NewDescribeCkafkaTypeConfigsResponse() (response *DescribeCkafkaTypeConfigsResponse) {
	response = &DescribeCkafkaTypeConfigsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取实例规格配置
func (c *Client) DescribeCkafkaTypeConfigs(request *DescribeCkafkaTypeConfigsRequest) (response *DescribeCkafkaTypeConfigsResponse, err error) {
	if request == nil {
		request = NewDescribeCkafkaTypeConfigsRequest()
	}
	response = NewDescribeCkafkaTypeConfigsResponse()
	err = c.Send(request, response)
	return
}

func NewCreatePartitionRequest() (request *CreatePartitionRequest) {
	request = &CreatePartitionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreatePartition")
	return
}

func NewCreatePartitionResponse() (response *CreatePartitionResponse) {
	response = &CreatePartitionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于增加主题中的分区
func (c *Client) CreatePartition(request *CreatePartitionRequest) (response *CreatePartitionResponse, err error) {
	if request == nil {
		request = NewCreatePartitionRequest()
	}
	response = NewCreatePartitionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeACLRequest() (request *DescribeACLRequest) {
	request = &DescribeACLRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeACL")
	return
}

func NewDescribeACLResponse() (response *DescribeACLResponse) {
	response = &DescribeACLResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 枚举ACL
func (c *Client) DescribeACL(request *DescribeACLRequest) (response *DescribeACLResponse, err error) {
	if request == nil {
		request = NewDescribeACLRequest()
	}
	response = NewDescribeACLResponse()
	err = c.Send(request, response)
	return
}

func NewFetchMessageListByOffsetRequest() (request *FetchMessageListByOffsetRequest) {
	request = &FetchMessageListByOffsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "FetchMessageListByOffset")
	return
}

func NewFetchMessageListByOffsetResponse() (response *FetchMessageListByOffsetResponse) {
	response = &FetchMessageListByOffsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据位点查询消息列表
func (c *Client) FetchMessageListByOffset(request *FetchMessageListByOffsetRequest) (response *FetchMessageListByOffsetResponse, err error) {
	if request == nil {
		request = NewFetchMessageListByOffsetRequest()
	}
	response = NewFetchMessageListByOffsetResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopicAttributesRequest() (request *ModifyTopicAttributesRequest) {
	request = &ModifyTopicAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "ModifyTopicAttributes")
	return
}

func NewModifyTopicAttributesResponse() (response *ModifyTopicAttributesResponse) {
	response = &ModifyTopicAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本接口用于修改主题属性。
func (c *Client) ModifyTopicAttributes(request *ModifyTopicAttributesRequest) (response *ModifyTopicAttributesResponse, err error) {
	if request == nil {
		request = NewModifyTopicAttributesRequest()
	}
	response = NewModifyTopicAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCkafkaPriceRequest() (request *DescribeCkafkaPriceRequest) {
	request = &DescribeCkafkaPriceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeCkafkaPrice")
	return
}

func NewDescribeCkafkaPriceResponse() (response *DescribeCkafkaPriceResponse) {
	response = &DescribeCkafkaPriceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 询价
func (c *Client) DescribeCkafkaPrice(request *DescribeCkafkaPriceRequest) (response *DescribeCkafkaPriceResponse, err error) {
	if request == nil {
		request = NewDescribeCkafkaPriceRequest()
	}
	response = NewDescribeCkafkaPriceResponse()
	err = c.Send(request, response)
	return
}

func NewHourCreateRequest() (request *HourCreateRequest) {
	request = &HourCreateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "HourCreate")
	return
}

func NewHourCreateResponse() (response *HourCreateResponse) {
	response = &HourCreateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建按量计费实例
func (c *Client) HourCreate(request *HourCreateRequest) (response *HourCreateResponse, err error) {
	if request == nil {
		request = NewHourCreateRequest()
	}
	response = NewHourCreateResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppInfoRequest() (request *DescribeAppInfoRequest) {
	request = &DescribeAppInfoRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeAppInfo")
	return
}

func NewDescribeAppInfoResponse() (response *DescribeAppInfoResponse) {
	response = &DescribeAppInfoResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户列表
func (c *Client) DescribeAppInfo(request *DescribeAppInfoRequest) (response *DescribeAppInfoResponse, err error) {
	if request == nil {
		request = NewDescribeAppInfoRequest()
	}
	response = NewDescribeAppInfoResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCkafkaZoneRequest() (request *DescribeCkafkaZoneRequest) {
	request = &DescribeCkafkaZoneRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeCkafkaZone")
	return
}

func NewDescribeCkafkaZoneResponse() (response *DescribeCkafkaZoneResponse) {
	response = &DescribeCkafkaZoneResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查看ckafka的可用区列表
func (c *Client) DescribeCkafkaZone(request *DescribeCkafkaZoneRequest) (response *DescribeCkafkaZoneResponse, err error) {
	if request == nil {
		request = NewDescribeCkafkaZoneRequest()
	}
	response = NewDescribeCkafkaZoneResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopicRequest() (request *DeleteTopicRequest) {
	request = &DeleteTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DeleteTopic")
	return
}

func NewDeleteTopicResponse() (response *DeleteTopicResponse) {
	response = &DeleteTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除ckafka主题
func (c *Client) DeleteTopic(request *DeleteTopicRequest) (response *DeleteTopicResponse, err error) {
	if request == nil {
		request = NewDeleteTopicRequest()
	}
	response = NewDeleteTopicResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAclRequest() (request *CreateAclRequest) {
	request = &CreateAclRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreateAcl")
	return
}

func NewCreateAclResponse() (response *CreateAclResponse) {
	response = &CreateAclResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加 ACL 策略
func (c *Client) CreateAcl(request *CreateAclRequest) (response *CreateAclResponse, err error) {
	if request == nil {
		request = NewCreateAclRequest()
	}
	response = NewCreateAclResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConnectorStatusRequest() (request *DescribeConnectorStatusRequest) {
	request = &DescribeConnectorStatusRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeConnectorStatus")
	return
}

func NewDescribeConnectorStatusResponse() (response *DescribeConnectorStatusResponse) {
	response = &DescribeConnectorStatusResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询Connector状态
func (c *Client) DescribeConnectorStatus(request *DescribeConnectorStatusRequest) (response *DescribeConnectorStatusResponse, err error) {
	if request == nil {
		request = NewDescribeConnectorStatusRequest()
	}
	response = NewDescribeConnectorStatusResponse()
	err = c.Send(request, response)
	return
}

func NewPauseConnectorRequest() (request *PauseConnectorRequest) {
	request = &PauseConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "PauseConnector")
	return
}

func NewPauseConnectorResponse() (response *PauseConnectorResponse) {
	response = &PauseConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 暂停数据同步任务
func (c *Client) PauseConnector(request *PauseConnectorRequest) (response *PauseConnectorResponse, err error) {
	if request == nil {
		request = NewPauseConnectorRequest()
	}
	response = NewPauseConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeOrderConfigRequest() (request *DescribeOrderConfigRequest) {
	request = &DescribeOrderConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeOrderConfig")
	return
}

func NewDescribeOrderConfigResponse() (response *DescribeOrderConfigResponse) {
	response = &DescribeOrderConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeOrderConfig
func (c *Client) DescribeOrderConfig(request *DescribeOrderConfigRequest) (response *DescribeOrderConfigResponse, err error) {
	if request == nil {
		request = NewDescribeOrderConfigRequest()
	}
	response = NewDescribeOrderConfigResponse()
	err = c.Send(request, response)
	return
}

func NewModifyResourceTceRequest() (request *ModifyResourceTceRequest) {
	request = &ModifyResourceTceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "ModifyResourceTce")
	return
}

func NewModifyResourceTceResponse() (response *ModifyResourceTceResponse) {
	response = &ModifyResourceTceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改实例
func (c *Client) ModifyResourceTce(request *ModifyResourceTceRequest) (response *ModifyResourceTceResponse, err error) {
	if request == nil {
		request = NewModifyResourceTceRequest()
	}
	response = NewModifyResourceTceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
	request = &CreateRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreateRoute")
	return
}

func NewCreateRouteResponse() (response *CreateRouteResponse) {
	response = &CreateRouteResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 添加实例路由
func (c *Client) CreateRoute(request *CreateRouteRequest) (response *CreateRouteResponse, err error) {
	if request == nil {
		request = NewCreateRouteRequest()
	}
	response = NewCreateRouteResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeGroupOffsetsRequest() (request *DescribeGroupOffsetsRequest) {
	request = &DescribeGroupOffsetsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeGroupOffsets")
	return
}

func NewDescribeGroupOffsetsResponse() (response *DescribeGroupOffsetsResponse) {
	response = &DescribeGroupOffsetsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取消费分组offset
func (c *Client) DescribeGroupOffsets(request *DescribeGroupOffsetsRequest) (response *DescribeGroupOffsetsResponse, err error) {
	if request == nil {
		request = NewDescribeGroupOffsetsRequest()
	}
	response = NewDescribeGroupOffsetsResponse()
	err = c.Send(request, response)
	return
}

func NewFetchMessageListByTimestampRequest() (request *FetchMessageListByTimestampRequest) {
	request = &FetchMessageListByTimestampRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "FetchMessageListByTimestamp")
	return
}

func NewFetchMessageListByTimestampResponse() (response *FetchMessageListByTimestampResponse) {
	response = &FetchMessageListByTimestampResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据时间戳查询消息列表
func (c *Client) FetchMessageListByTimestamp(request *FetchMessageListByTimestampRequest) (response *FetchMessageListByTimestampResponse, err error) {
	if request == nil {
		request = NewFetchMessageListByTimestampRequest()
	}
	response = NewFetchMessageListByTimestampResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopicIpWhiteListRequest() (request *CreateTopicIpWhiteListRequest) {
	request = &CreateTopicIpWhiteListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreateTopicIpWhiteList")
	return
}

func NewCreateTopicIpWhiteListResponse() (response *CreateTopicIpWhiteListResponse) {
	response = &CreateTopicIpWhiteListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建主题ip白名单
func (c *Client) CreateTopicIpWhiteList(request *CreateTopicIpWhiteListRequest) (response *CreateTopicIpWhiteListResponse, err error) {
	if request == nil {
		request = NewCreateTopicIpWhiteListRequest()
	}
	response = NewCreateTopicIpWhiteListResponse()
	err = c.Send(request, response)
	return
}

func NewResumeConnectorRequest() (request *ResumeConnectorRequest) {
	request = &ResumeConnectorRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "ResumeConnector")
	return
}

func NewResumeConnectorResponse() (response *ResumeConnectorResponse) {
	response = &ResumeConnectorResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 启动Connector任务
func (c *Client) ResumeConnector(request *ResumeConnectorRequest) (response *ResumeConnectorResponse, err error) {
	if request == nil {
		request = NewResumeConnectorRequest()
	}
	response = NewResumeConnectorResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAppIdIsVipRequest() (request *DescribeAppIdIsVipRequest) {
	request = &DescribeAppIdIsVipRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeAppIdIsVip")
	return
}

func NewDescribeAppIdIsVipResponse() (response *DescribeAppIdIsVipResponse) {
	response = &DescribeAppIdIsVipResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询客户是否为大客户
func (c *Client) DescribeAppIdIsVip(request *DescribeAppIdIsVipRequest) (response *DescribeAppIdIsVipResponse, err error) {
	if request == nil {
		request = NewDescribeAppIdIsVipRequest()
	}
	response = NewDescribeAppIdIsVipResponse()
	err = c.Send(request, response)
	return
}

func NewFetchMessageByOffsetRequest() (request *FetchMessageByOffsetRequest) {
	request = &FetchMessageByOffsetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "FetchMessageByOffset")
	return
}

func NewFetchMessageByOffsetResponse() (response *FetchMessageByOffsetResponse) {
	response = &FetchMessageByOffsetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据指定offset位置的消息
func (c *Client) FetchMessageByOffset(request *FetchMessageByOffsetRequest) (response *FetchMessageByOffsetResponse, err error) {
	if request == nil {
		request = NewFetchMessageByOffsetRequest()
	}
	response = NewFetchMessageByOffsetResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInstancesDetailRequest() (request *DescribeInstancesDetailRequest) {
	request = &DescribeInstancesDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DescribeInstancesDetail")
	return
}

func NewDescribeInstancesDetailResponse() (response *DescribeInstancesDetailResponse) {
	response = &DescribeInstancesDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用户账户下获取实例列表详情
func (c *Client) DescribeInstancesDetail(request *DescribeInstancesDetailRequest) (response *DescribeInstancesDetailResponse, err error) {
	if request == nil {
		request = NewDescribeInstancesDetailRequest()
	}
	response = NewDescribeInstancesDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTokenRequest() (request *CreateTokenRequest) {
	request = &CreateTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CreateToken")
	return
}

func NewCreateTokenResponse() (response *CreateTokenResponse) {
	response = &CreateTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建最高权限的token
func (c *Client) CreateToken(request *CreateTokenRequest) (response *CreateTokenResponse, err error) {
	if request == nil {
		request = NewCreateTokenRequest()
	}
	response = NewCreateTokenResponse()
	err = c.Send(request, response)
	return
}

func NewCancelAuthorizationTokenRequest() (request *CancelAuthorizationTokenRequest) {
	request = &CancelAuthorizationTokenRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "CancelAuthorizationToken")
	return
}

func NewCancelAuthorizationTokenResponse() (response *CancelAuthorizationTokenResponse) {
	response = &CancelAuthorizationTokenResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// CancelAuthorizationToken
func (c *Client) CancelAuthorizationToken(request *CancelAuthorizationTokenRequest) (response *CancelAuthorizationTokenResponse, err error) {
	if request == nil {
		request = NewCancelAuthorizationTokenRequest()
	}
	response = NewCancelAuthorizationTokenResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteUserRequest() (request *DeleteUserRequest) {
	request = &DeleteUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("ckafka", APIVersion, "DeleteUser")
	return
}

func NewDeleteUserResponse() (response *DeleteUserResponse) {
	response = &DeleteUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除用户
func (c *Client) DeleteUser(request *DeleteUserRequest) (response *DeleteUserResponse, err error) {
	if request == nil {
		request = NewDeleteUserRequest()
	}
	response = NewDeleteUserResponse()
	err = c.Send(request, response)
	return
}

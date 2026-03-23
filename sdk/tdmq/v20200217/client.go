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

package v20200217

import (
	"encoding/json"

	"terraform-provider-tencentcloudenterprise/sdk/common"
	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
	"terraform-provider-tencentcloudenterprise/sdk/common/profile"
)

const APIVersion = "2020-02-17"

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

func NewDescribeDeliveryStatusOptRequest() (request *DescribeDeliveryStatusOptRequest) {
	request = &DescribeDeliveryStatusOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeDeliveryStatusOpt")
	return
}

func NewDescribeDeliveryStatusOptResponse() (response *DescribeDeliveryStatusOptResponse) {
	response = &DescribeDeliveryStatusOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置下发状态信息
func (c *Client) DescribeDeliveryStatusOpt(request *DescribeDeliveryStatusOptRequest) (response *DescribeDeliveryStatusOptResponse, err error) {
	if request == nil {
		request = NewDescribeDeliveryStatusOptRequest()
	}
	response = NewDescribeDeliveryStatusOptResponse()
	err = c.Send(request, response)
	return
}

func NewRebootInstancesOptRequest() (request *RebootInstancesOptRequest) {
	request = &RebootInstancesOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "RebootInstancesOpt")
	return
}

func NewRebootInstancesOptResponse() (response *RebootInstancesOptResponse) {
	response = &RebootInstancesOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端批量重启实例
func (c *Client) RebootInstancesOpt(request *RebootInstancesOptRequest) (response *RebootInstancesOptResponse, err error) {
	if request == nil {
		request = NewRebootInstancesOptRequest()
	}
	response = NewRebootInstancesOptResponse()
	err = c.Send(request, response)
	return
}

func NewResetRocketMQConsumerOffSetRequest() (request *ResetRocketMQConsumerOffSetRequest) {
	request = &ResetRocketMQConsumerOffSetRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ResetRocketMQConsumerOffSet")
	return
}

func NewResetRocketMQConsumerOffSetResponse() (response *ResetRocketMQConsumerOffSetResponse) {
	response = &ResetRocketMQConsumerOffSetResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重置指定Group的消费位点到指定时间戳
func (c *Client) ResetRocketMQConsumerOffSet(request *ResetRocketMQConsumerOffSetRequest) (response *ResetRocketMQConsumerOffSetResponse, err error) {
	if request == nil {
		request = NewResetRocketMQConsumerOffSetRequest()
	}
	response = NewResetRocketMQConsumerOffSetResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCmqTopicRequest() (request *CreateCmqTopicRequest) {
	request = &CreateCmqTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqTopic")
	return
}

func NewCreateCmqTopicResponse() (response *CreateCmqTopicResponse) {
	response = &CreateCmqTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建cmq主题
func (c *Client) CreateCmqTopic(request *CreateCmqTopicRequest) (response *CreateCmqTopicResponse, err error) {
	if request == nil {
		request = NewCreateCmqTopicRequest()
	}
	response = NewCreateCmqTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBindVpcsRequest() (request *DescribeBindVpcsRequest) {
	request = &DescribeBindVpcsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBindVpcs")
	return
}

func NewDescribeBindVpcsResponse() (response *DescribeBindVpcsResponse) {
	response = &DescribeBindVpcsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户VPC绑定关系
func (c *Client) DescribeBindVpcs(request *DescribeBindVpcsRequest) (response *DescribeBindVpcsResponse, err error) {
	if request == nil {
		request = NewDescribeBindVpcsRequest()
	}
	response = NewDescribeBindVpcsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRocketMQClusterRequest() (request *DeleteRocketMQClusterRequest) {
	request = &DeleteRocketMQClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQCluster")
	return
}

func NewDeleteRocketMQClusterResponse() (response *DeleteRocketMQClusterResponse) {
	response = &DeleteRocketMQClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RocketMQ集群
func (c *Client) DeleteRocketMQCluster(request *DeleteRocketMQClusterRequest) (response *DeleteRocketMQClusterResponse, err error) {
	if request == nil {
		request = NewDeleteRocketMQClusterRequest()
	}
	response = NewDeleteRocketMQClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRolesRequest() (request *DescribeRolesRequest) {
	request = &DescribeRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRoles")
	return
}

func NewDescribeRolesResponse() (response *DescribeRolesResponse) {
	response = &DescribeRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取角色列表
func (c *Client) DescribeRoles(request *DescribeRolesRequest) (response *DescribeRolesResponse, err error) {
	if request == nil {
		request = NewDescribeRolesRequest()
	}
	response = NewDescribeRolesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAMQPRouteRelationRequest() (request *DeleteAMQPRouteRelationRequest) {
	request = &DeleteAMQPRouteRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteAMQPRouteRelation")
	return
}

func NewDeleteAMQPRouteRelationResponse() (response *DeleteAMQPRouteRelationResponse) {
	response = &DeleteAMQPRouteRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Amqp路由关系
func (c *Client) DeleteAMQPRouteRelation(request *DeleteAMQPRouteRelationRequest) (response *DeleteAMQPRouteRelationResponse, err error) {
	if request == nil {
		request = NewDeleteAMQPRouteRelationRequest()
	}
	response = NewDeleteAMQPRouteRelationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPVHostsRequest() (request *DescribeAMQPVHostsRequest) {
	request = &DescribeAMQPVHostsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPVHosts")
	return
}

func NewDescribeAMQPVHostsResponse() (response *DescribeAMQPVHostsResponse) {
	response = &DescribeAMQPVHostsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Amqp Vhost 列表
func (c *Client) DescribeAMQPVHosts(request *DescribeAMQPVHostsRequest) (response *DescribeAMQPVHostsResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPVHostsRequest()
	}
	response = NewDescribeAMQPVHostsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAMQPQueueRequest() (request *DeleteAMQPQueueRequest) {
	request = &DeleteAMQPQueueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteAMQPQueue")
	return
}

func NewDeleteAMQPQueueResponse() (response *DeleteAMQPQueueResponse) {
	response = &DeleteAMQPQueueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Amqp队列
func (c *Client) DeleteAMQPQueue(request *DeleteAMQPQueueRequest) (response *DeleteAMQPQueueResponse, err error) {
	if request == nil {
		request = NewDeleteAMQPQueueRequest()
	}
	response = NewDeleteAMQPQueueResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterDetailRequest() (request *DescribeClusterDetailRequest) {
	request = &DescribeClusterDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterDetail")
	return
}

func NewDescribeClusterDetailResponse() (response *DescribeClusterDetailResponse) {
	response = &DescribeClusterDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群的详细信息
func (c *Client) DescribeClusterDetail(request *DescribeClusterDetailRequest) (response *DescribeClusterDetailResponse, err error) {
	if request == nil {
		request = NewDescribeClusterDetailRequest()
	}
	response = NewDescribeClusterDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInternalRocketMQTopicsRequest() (request *DescribeInternalRocketMQTopicsRequest) {
	request = &DescribeInternalRocketMQTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeInternalRocketMQTopics")
	return
}

func NewDescribeInternalRocketMQTopicsResponse() (response *DescribeInternalRocketMQTopicsResponse) {
	response = &DescribeInternalRocketMQTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此为内部接口，用于运营端查询主题列表
func (c *Client) DescribeInternalRocketMQTopics(request *DescribeInternalRocketMQTopicsRequest) (response *DescribeInternalRocketMQTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeInternalRocketMQTopicsRequest()
	}
	response = NewDescribeInternalRocketMQTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCmqQueueRequest() (request *CreateCmqQueueRequest) {
	request = &CreateCmqQueueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqQueue")
	return
}

func NewCreateCmqQueueResponse() (response *CreateCmqQueueResponse) {
	response = &CreateCmqQueueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建cmq队列接口
func (c *Client) CreateCmqQueue(request *CreateCmqQueueRequest) (response *CreateCmqQueueResponse, err error) {
	if request == nil {
		request = NewCreateCmqQueueRequest()
	}
	response = NewCreateCmqQueueResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRocketMQNamespaceRequest() (request *DeleteRocketMQNamespaceRequest) {
	request = &DeleteRocketMQNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQNamespace")
	return
}

func NewDeleteRocketMQNamespaceResponse() (response *DeleteRocketMQNamespaceResponse) {
	response = &DeleteRocketMQNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RocketMQ命名空间
func (c *Client) DeleteRocketMQNamespace(request *DeleteRocketMQNamespaceRequest) (response *DeleteRocketMQNamespaceResponse, err error) {
	if request == nil {
		request = NewDeleteRocketMQNamespaceRequest()
	}
	response = NewDeleteRocketMQNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewPublishCmqMsgRequest() (request *PublishCmqMsgRequest) {
	request = &PublishCmqMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "PublishCmqMsg")
	return
}

func NewPublishCmqMsgResponse() (response *PublishCmqMsgResponse) {
	response = &PublishCmqMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送cmq主题消息
func (c *Client) PublishCmqMsg(request *PublishCmqMsgRequest) (response *PublishCmqMsgResponse, err error) {
	if request == nil {
		request = NewPublishCmqMsgRequest()
	}
	response = NewPublishCmqMsgResponse()
	err = c.Send(request, response)
	return
}

func NewSplitNamespaceBundleOptRequest() (request *SplitNamespaceBundleOptRequest) {
	request = &SplitNamespaceBundleOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SplitNamespaceBundleOpt")
	return
}

func NewSplitNamespaceBundleOptResponse() (response *SplitNamespaceBundleOptResponse) {
	response = &SplitNamespaceBundleOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端命名空间分裂bundle
func (c *Client) SplitNamespaceBundleOpt(request *SplitNamespaceBundleOptRequest) (response *SplitNamespaceBundleOptResponse, err error) {
	if request == nil {
		request = NewSplitNamespaceBundleOptRequest()
	}
	response = NewSplitNamespaceBundleOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateItemsOptRequest() (request *CreateTemplateItemsOptRequest) {
	request = &CreateTemplateItemsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateTemplateItemsOpt")
	return
}

func NewCreateTemplateItemsOptResponse() (response *CreateTemplateItemsOptResponse) {
	response = &CreateTemplateItemsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 模板添加配置项
func (c *Client) CreateTemplateItemsOpt(request *CreateTemplateItemsOptRequest) (response *CreateTemplateItemsOptResponse, err error) {
	if request == nil {
		request = NewCreateTemplateItemsOptRequest()
	}
	response = NewCreateTemplateItemsOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEnvironmentsRequest() (request *DescribeEnvironmentsRequest) {
	request = &DescribeEnvironmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironments")
	return
}

func NewDescribeEnvironmentsResponse() (response *DescribeEnvironmentsResponse) {
	response = &DescribeEnvironmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取租户下命名空间列表
func (c *Client) DescribeEnvironments(request *DescribeEnvironmentsRequest) (response *DescribeEnvironmentsResponse, err error) {
	if request == nil {
		request = NewDescribeEnvironmentsRequest()
	}
	response = NewDescribeEnvironmentsResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRocketMQClusterRequest() (request *CreateRocketMQClusterRequest) {
	request = &CreateRocketMQClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQCluster")
	return
}

func NewCreateRocketMQClusterResponse() (response *CreateRocketMQClusterResponse) {
	response = &CreateRocketMQClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口用于创建一个RocketMQ集群
func (c *Client) CreateRocketMQCluster(request *CreateRocketMQClusterRequest) (response *CreateRocketMQClusterResponse, err error) {
	if request == nil {
		request = NewCreateRocketMQClusterRequest()
	}
	response = NewCreateRocketMQClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPClusterRequest() (request *DescribeAMQPClusterRequest) {
	request = &DescribeAMQPClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPCluster")
	return
}

func NewDescribeAMQPClusterResponse() (response *DescribeAMQPClusterResponse) {
	response = &DescribeAMQPClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个Amqp集群信息
func (c *Client) DescribeAMQPCluster(request *DescribeAMQPClusterRequest) (response *DescribeAMQPClusterResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPClusterRequest()
	}
	response = NewDescribeAMQPClusterResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateConfigOptRequest() (request *UpdateConfigOptRequest) {
	request = &UpdateConfigOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdateConfigOpt")
	return
}

func NewUpdateConfigOptResponse() (response *UpdateConfigOptResponse) {
	response = &UpdateConfigOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新集群配置信息
func (c *Client) UpdateConfigOpt(request *UpdateConfigOptRequest) (response *UpdateConfigOptResponse, err error) {
	if request == nil {
		request = NewUpdateConfigOptRequest()
	}
	response = NewUpdateConfigOptResponse()
	err = c.Send(request, response)
	return
}

func NewSendMessagesRequest() (request *SendMessagesRequest) {
	request = &SendMessagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SendMessages")
	return
}

func NewSendMessagesResponse() (response *SendMessagesResponse) {
	response = &SendMessagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送单条消息
func (c *Client) SendMessages(request *SendMessagesRequest) (response *SendMessagesResponse, err error) {
	if request == nil {
		request = NewSendMessagesRequest()
	}
	response = NewSendMessagesResponse()
	err = c.Send(request, response)
	return
}

func NewSendRocketMQMessageRequest() (request *SendRocketMQMessageRequest) {
	request = &SendRocketMQMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SendRocketMQMessage")
	return
}

func NewSendRocketMQMessageResponse() (response *SendRocketMQMessageResponse) {
	response = &SendRocketMQMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送RocketMQ消息
func (c *Client) SendRocketMQMessage(request *SendRocketMQMessageRequest) (response *SendRocketMQMessageResponse, err error) {
	if request == nil {
		request = NewSendRocketMQMessageRequest()
	}
	response = NewSendRocketMQMessageResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicReplicationClustersRequest() (request *DescribeTopicReplicationClustersRequest) {
	request = &DescribeTopicReplicationClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeTopicReplicationClusters")
	return
}

func NewDescribeTopicReplicationClustersResponse() (response *DescribeTopicReplicationClustersResponse) {
	response = &DescribeTopicReplicationClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询topic跨地域复制绑定关系
func (c *Client) DescribeTopicReplicationClusters(request *DescribeTopicReplicationClustersRequest) (response *DescribeTopicReplicationClustersResponse, err error) {
	if request == nil {
		request = NewDescribeTopicReplicationClustersRequest()
	}
	response = NewDescribeTopicReplicationClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRocketMQTopicRequest() (request *DeleteRocketMQTopicRequest) {
	request = &DeleteRocketMQTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQTopic")
	return
}

func NewDeleteRocketMQTopicResponse() (response *DeleteRocketMQTopicResponse) {
	response = &DeleteRocketMQTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RocketMQ主题
func (c *Client) DeleteRocketMQTopic(request *DeleteRocketMQTopicRequest) (response *DeleteRocketMQTopicResponse, err error) {
	if request == nil {
		request = NewDeleteRocketMQTopicRequest()
	}
	response = NewDeleteRocketMQTopicResponse()
	err = c.Send(request, response)
	return
}

func NewUnloadNamespaceBundleOptRequest() (request *UnloadNamespaceBundleOptRequest) {
	request = &UnloadNamespaceBundleOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UnloadNamespaceBundleOpt")
	return
}

func NewUnloadNamespaceBundleOptResponse() (response *UnloadNamespaceBundleOptResponse) {
	response = &UnloadNamespaceBundleOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端卸载命名空间bundle
func (c *Client) UnloadNamespaceBundleOpt(request *UnloadNamespaceBundleOptRequest) (response *UnloadNamespaceBundleOptResponse, err error) {
	if request == nil {
		request = NewUnloadNamespaceBundleOptRequest()
	}
	response = NewUnloadNamespaceBundleOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTopicRequest() (request *CreateTopicRequest) {
	request = &CreateTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateTopic")
	return
}

func NewCreateTopicResponse() (response *CreateTopicResponse) {
	response = &CreateTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增指定分区、类型的消息主题
func (c *Client) CreateTopic(request *CreateTopicRequest) (response *CreateTopicResponse, err error) {
	if request == nil {
		request = NewCreateTopicRequest()
	}
	response = NewCreateTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterConfigTypeOptRequest() (request *DescribeClusterConfigTypeOptRequest) {
	request = &DescribeClusterConfigTypeOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterConfigTypeOpt")
	return
}

func NewDescribeClusterConfigTypeOptResponse() (response *DescribeClusterConfigTypeOptResponse) {
	response = &DescribeClusterConfigTypeOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群配置类型信息
func (c *Client) DescribeClusterConfigTypeOpt(request *DescribeClusterConfigTypeOptRequest) (response *DescribeClusterConfigTypeOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterConfigTypeOptRequest()
	}
	response = NewDescribeClusterConfigTypeOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPRouteRelationsRequest() (request *DescribeAMQPRouteRelationsRequest) {
	request = &DescribeAMQPRouteRelationsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPRouteRelations")
	return
}

func NewDescribeAMQPRouteRelationsResponse() (response *DescribeAMQPRouteRelationsResponse) {
	response = &DescribeAMQPRouteRelationsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Amqp路由关系列表
func (c *Client) DescribeAMQPRouteRelations(request *DescribeAMQPRouteRelationsRequest) (response *DescribeAMQPRouteRelationsResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPRouteRelationsRequest()
	}
	response = NewDescribeAMQPRouteRelationsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyOffloadThresholdOptRequest() (request *UpdatePolicyOffloadThresholdOptRequest) {
	request = &UpdatePolicyOffloadThresholdOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyOffloadThresholdOpt")
	return
}

func NewUpdatePolicyOffloadThresholdOptResponse() (response *UpdatePolicyOffloadThresholdOptResponse) {
	response = &UpdatePolicyOffloadThresholdOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间消息卸载策略
func (c *Client) UpdatePolicyOffloadThresholdOpt(request *UpdatePolicyOffloadThresholdOptRequest) (response *UpdatePolicyOffloadThresholdOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyOffloadThresholdOptRequest()
	}
	response = NewUpdatePolicyOffloadThresholdOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllRocketMQTenantsRequest() (request *DescribeAllRocketMQTenantsRequest) {
	request = &DescribeAllRocketMQTenantsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAllRocketMQTenants")
	return
}

func NewDescribeAllRocketMQTenantsResponse() (response *DescribeAllRocketMQTenantsResponse) {
	response = &DescribeAllRocketMQTenantsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取某个租户的虚拟集群列表
func (c *Client) DescribeAllRocketMQTenants(request *DescribeAllRocketMQTenantsRequest) (response *DescribeAllRocketMQTenantsResponse, err error) {
	if request == nil {
		request = NewDescribeAllRocketMQTenantsRequest()
	}
	response = NewDescribeAllRocketMQTenantsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyCompactionThresholdOptRequest() (request *UpdatePolicyCompactionThresholdOptRequest) {
	request = &UpdatePolicyCompactionThresholdOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyCompactionThresholdOpt")
	return
}

func NewUpdatePolicyCompactionThresholdOptResponse() (response *UpdatePolicyCompactionThresholdOptResponse) {
	response = &UpdatePolicyCompactionThresholdOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间消息压缩策略
func (c *Client) UpdatePolicyCompactionThresholdOpt(request *UpdatePolicyCompactionThresholdOptRequest) (response *UpdatePolicyCompactionThresholdOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyCompactionThresholdOptRequest()
	}
	response = NewUpdatePolicyCompactionThresholdOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEnvironmentAttributesRequest() (request *DescribeEnvironmentAttributesRequest) {
	request = &DescribeEnvironmentAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironmentAttributes")
	return
}

func NewDescribeEnvironmentAttributesResponse() (response *DescribeEnvironmentAttributesResponse) {
	response = &DescribeEnvironmentAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定命名空间的属性
func (c *Client) DescribeEnvironmentAttributes(request *DescribeEnvironmentAttributesRequest) (response *DescribeEnvironmentAttributesResponse, err error) {
	if request == nil {
		request = NewDescribeEnvironmentAttributesRequest()
	}
	response = NewDescribeEnvironmentAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyPublishRateOptRequest() (request *UpdatePolicyPublishRateOptRequest) {
	request = &UpdatePolicyPublishRateOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyPublishRateOpt")
	return
}

func NewUpdatePolicyPublishRateOptResponse() (response *UpdatePolicyPublishRateOptResponse) {
	response = &UpdatePolicyPublishRateOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间生产限流策略
func (c *Client) UpdatePolicyPublishRateOpt(request *UpdatePolicyPublishRateOptRequest) (response *UpdatePolicyPublishRateOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyPublishRateOptRequest()
	}
	response = NewUpdatePolicyPublishRateOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAgentNodesOptRequest() (request *DescribeAgentNodesOptRequest) {
	request = &DescribeAgentNodesOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAgentNodesOpt")
	return
}

func NewDescribeAgentNodesOptResponse() (response *DescribeAgentNodesOptResponse) {
	response = &DescribeAgentNodesOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营平台接口 获取AgentNodes接口
func (c *Client) DescribeAgentNodesOpt(request *DescribeAgentNodesOptRequest) (response *DescribeAgentNodesOptResponse, err error) {
	if request == nil {
		request = NewDescribeAgentNodesOptRequest()
	}
	response = NewDescribeAgentNodesOptResponse()
	err = c.Send(request, response)
	return
}

func NewUnbindCmqDeadLetterRequest() (request *UnbindCmqDeadLetterRequest) {
	request = &UnbindCmqDeadLetterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UnbindCmqDeadLetter")
	return
}

func NewUnbindCmqDeadLetterResponse() (response *UnbindCmqDeadLetterResponse) {
	response = &UnbindCmqDeadLetterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 解绑cmq死信队列
func (c *Client) UnbindCmqDeadLetter(request *UnbindCmqDeadLetterRequest) (response *UnbindCmqDeadLetterResponse, err error) {
	if request == nil {
		request = NewUnbindCmqDeadLetterRequest()
	}
	response = NewUnbindCmqDeadLetterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQClusterRequest() (request *DescribeRocketMQClusterRequest) {
	request = &DescribeRocketMQClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQCluster")
	return
}

func NewDescribeRocketMQClusterResponse() (response *DescribeRocketMQClusterResponse) {
	response = &DescribeRocketMQClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个RocketMQ集群信息
func (c *Client) DescribeRocketMQCluster(request *DescribeRocketMQClusterRequest) (response *DescribeRocketMQClusterResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQClusterRequest()
	}
	response = NewDescribeRocketMQClusterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRoleRequest() (request *ModifyRoleRequest) {
	request = &ModifyRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRole")
	return
}

func NewModifyRoleResponse() (response *ModifyRoleResponse) {
	response = &ModifyRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 角色修改
func (c *Client) ModifyRole(request *ModifyRoleRequest) (response *ModifyRoleResponse, err error) {
	if request == nil {
		request = NewModifyRoleRequest()
	}
	response = NewModifyRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllVersionConfigsOptRequest() (request *DescribeAllVersionConfigsOptRequest) {
	request = &DescribeAllVersionConfigsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAllVersionConfigsOpt")
	return
}

func NewDescribeAllVersionConfigsOptResponse() (response *DescribeAllVersionConfigsOptResponse) {
	response = &DescribeAllVersionConfigsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群所有版本配置基础信息
func (c *Client) DescribeAllVersionConfigsOpt(request *DescribeAllVersionConfigsOptRequest) (response *DescribeAllVersionConfigsOptResponse, err error) {
	if request == nil {
		request = NewDescribeAllVersionConfigsOptRequest()
	}
	response = NewDescribeAllVersionConfigsOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBookieDiskListOptRequest() (request *DescribeBookieDiskListOptRequest) {
	request = &DescribeBookieDiskListOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBookieDiskListOpt")
	return
}

func NewDescribeBookieDiskListOptResponse() (response *DescribeBookieDiskListOptResponse) {
	response = &DescribeBookieDiskListOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取bookie数据盘列表及使用率
func (c *Client) DescribeBookieDiskListOpt(request *DescribeBookieDiskListOptRequest) (response *DescribeBookieDiskListOptResponse, err error) {
	if request == nil {
		request = NewDescribeBookieDiskListOptRequest()
	}
	response = NewDescribeBookieDiskListOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeLatestMetricsOptRequest() (request *DescribeNodeLatestMetricsOptRequest) {
	request = &DescribeNodeLatestMetricsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeNodeLatestMetricsOpt")
	return
}

func NewDescribeNodeLatestMetricsOptResponse() (response *DescribeNodeLatestMetricsOptResponse) {
	response = &DescribeNodeLatestMetricsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取节点监控指标最新值
func (c *Client) DescribeNodeLatestMetricsOpt(request *DescribeNodeLatestMetricsOptRequest) (response *DescribeNodeLatestMetricsOptResponse, err error) {
	if request == nil {
		request = NewDescribeNodeLatestMetricsOptRequest()
	}
	response = NewDescribeNodeLatestMetricsOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicMsgsRequest() (request *DescribeTopicMsgsRequest) {
	request = &DescribeTopicMsgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeTopicMsgs")
	return
}

func NewDescribeTopicMsgsResponse() (response *DescribeTopicMsgsResponse) {
	response = &DescribeTopicMsgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息查询
func (c *Client) DescribeTopicMsgs(request *DescribeTopicMsgsRequest) (response *DescribeTopicMsgsResponse, err error) {
	if request == nil {
		request = NewDescribeTopicMsgsRequest()
	}
	response = NewDescribeTopicMsgsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopicRequest() (request *ModifyTopicRequest) {
	request = &ModifyTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyTopic")
	return
}

func NewModifyTopicResponse() (response *ModifyTopicResponse) {
	response = &ModifyTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改主题备注和分区数
func (c *Client) ModifyTopic(request *ModifyTopicRequest) (response *ModifyTopicResponse, err error) {
	if request == nil {
		request = NewModifyTopicRequest()
	}
	response = NewModifyTopicResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInternalTenantQuotaRequest() (request *ModifyInternalTenantQuotaRequest) {
	request = &ModifyInternalTenantQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyInternalTenantQuota")
	return
}

func NewModifyInternalTenantQuotaResponse() (response *ModifyInternalTenantQuotaResponse) {
	response = &ModifyInternalTenantQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改虚拟集群配额
func (c *Client) ModifyInternalTenantQuota(request *ModifyInternalTenantQuotaRequest) (response *ModifyInternalTenantQuotaResponse, err error) {
	if request == nil {
		request = NewModifyInternalTenantQuotaRequest()
	}
	response = NewModifyInternalTenantQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewExportRocketMQMessageDetailRequest() (request *ExportRocketMQMessageDetailRequest) {
	request = &ExportRocketMQMessageDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ExportRocketMQMessageDetail")
	return
}

func NewExportRocketMQMessageDetailResponse() (response *ExportRocketMQMessageDetailResponse) {
	response = &ExportRocketMQMessageDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 导出RocketMQ消息详情
func (c *Client) ExportRocketMQMessageDetail(request *ExportRocketMQMessageDetailRequest) (response *ExportRocketMQMessageDetailResponse, err error) {
	if request == nil {
		request = NewExportRocketMQMessageDetailRequest()
	}
	response = NewExportRocketMQMessageDetailResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyDispatchRatePerSubscriptionOptRequest() (request *UpdatePolicyDispatchRatePerSubscriptionOptRequest) {
	request = &UpdatePolicyDispatchRatePerSubscriptionOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyDispatchRatePerSubscriptionOpt")
	return
}

func NewUpdatePolicyDispatchRatePerSubscriptionOptResponse() (response *UpdatePolicyDispatchRatePerSubscriptionOptResponse) {
	response = &UpdatePolicyDispatchRatePerSubscriptionOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间订阅级别请求分发策略限制
func (c *Client) UpdatePolicyDispatchRatePerSubscriptionOpt(request *UpdatePolicyDispatchRatePerSubscriptionOptRequest) (response *UpdatePolicyDispatchRatePerSubscriptionOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyDispatchRatePerSubscriptionOptRequest()
	}
	response = NewUpdatePolicyDispatchRatePerSubscriptionOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQMsgRequest() (request *DescribeRocketMQMsgRequest) {
	request = &DescribeRocketMQMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQMsg")
	return
}

func NewDescribeRocketMQMsgResponse() (response *DescribeRocketMQMsgResponse) {
	response = &DescribeRocketMQMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// rocketmq消息详情
func (c *Client) DescribeRocketMQMsg(request *DescribeRocketMQMsgRequest) (response *DescribeRocketMQMsgResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQMsgRequest()
	}
	response = NewDescribeRocketMQMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRolesRequest() (request *DeleteRolesRequest) {
	request = &DeleteRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRoles")
	return
}

func NewDeleteRolesResponse() (response *DeleteRolesResponse) {
	response = &DeleteRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除角色，支持批量。
func (c *Client) DeleteRoles(request *DeleteRolesRequest) (response *DeleteRolesResponse, err error) {
	if request == nil {
		request = NewDeleteRolesRequest()
	}
	response = NewDeleteRolesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQConsumerConnectionsRequest() (request *DescribeRocketMQConsumerConnectionsRequest) {
	request = &DescribeRocketMQConsumerConnectionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQConsumerConnections")
	return
}

func NewDescribeRocketMQConsumerConnectionsResponse() (response *DescribeRocketMQConsumerConnectionsResponse) {
	response = &DescribeRocketMQConsumerConnectionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定消费组下当前客户端的连接情况
func (c *Client) DescribeRocketMQConsumerConnections(request *DescribeRocketMQConsumerConnectionsRequest) (response *DescribeRocketMQConsumerConnectionsResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQConsumerConnectionsRequest()
	}
	response = NewDescribeRocketMQConsumerConnectionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterOptRequest() (request *DescribeClusterOptRequest) {
	request = &DescribeClusterOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterOpt")
	return
}

func NewDescribeClusterOptResponse() (response *DescribeClusterOptResponse) {
	response = &DescribeClusterOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取物理集群列表
func (c *Client) DescribeClusterOpt(request *DescribeClusterOptRequest) (response *DescribeClusterOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterOptRequest()
	}
	response = NewDescribeClusterOptResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTopicReplicationClustersRequest() (request *ModifyTopicReplicationClustersRequest) {
	request = &ModifyTopicReplicationClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyTopicReplicationClusters")
	return
}

func NewModifyTopicReplicationClustersResponse() (response *ModifyTopicReplicationClustersResponse) {
	response = &ModifyTopicReplicationClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 编辑Topic和跨地域绑定集群关系
func (c *Client) ModifyTopicReplicationClusters(request *ModifyTopicReplicationClustersRequest) (response *ModifyTopicReplicationClustersResponse, err error) {
	if request == nil {
		request = NewModifyTopicReplicationClustersRequest()
	}
	response = NewModifyTopicReplicationClustersResponse()
	err = c.Send(request, response)
	return
}

func NewCheckBatchRoleRequest() (request *CheckBatchRoleRequest) {
	request = &CheckBatchRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CheckBatchRole")
	return
}

func NewCheckBatchRoleResponse() (response *CheckBatchRoleResponse) {
	response = &CheckBatchRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 检验批量导入角色
func (c *Client) CheckBatchRole(request *CheckBatchRoleRequest) (response *CheckBatchRoleResponse, err error) {
	if request == nil {
		request = NewCheckBatchRoleRequest()
	}
	response = NewCheckBatchRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePClustersRequest() (request *DescribePClustersRequest) {
	request = &DescribePClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribePClusters")
	return
}

func NewDescribePClustersResponse() (response *DescribePClustersResponse) {
	response = &DescribePClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户可以使用物理集群列表
func (c *Client) DescribePClusters(request *DescribePClustersRequest) (response *DescribePClustersResponse, err error) {
	if request == nil {
		request = NewDescribePClustersRequest()
	}
	response = NewDescribePClustersResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAMQPExchangeRequest() (request *CreateAMQPExchangeRequest) {
	request = &CreateAMQPExchangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateAMQPExchange")
	return
}

func NewCreateAMQPExchangeResponse() (response *CreateAMQPExchangeResponse) {
	response = &CreateAMQPExchangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建AMQP Exchange
func (c *Client) CreateAMQPExchange(request *CreateAMQPExchangeRequest) (response *CreateAMQPExchangeResponse, err error) {
	if request == nil {
		request = NewCreateAMQPExchangeRequest()
	}
	response = NewCreateAMQPExchangeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCmqQueueAttributeRequest() (request *ModifyCmqQueueAttributeRequest) {
	request = &ModifyCmqQueueAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqQueueAttribute")
	return
}

func NewModifyCmqQueueAttributeResponse() (response *ModifyCmqQueueAttributeResponse) {
	response = &ModifyCmqQueueAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改cmq队列属性
func (c *Client) ModifyCmqQueueAttribute(request *ModifyCmqQueueAttributeRequest) (response *ModifyCmqQueueAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCmqQueueAttributeRequest()
	}
	response = NewModifyCmqQueueAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQRolesRequest() (request *DescribeRocketMQRolesRequest) {
	request = &DescribeRocketMQRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQRoles")
	return
}

func NewDescribeRocketMQRolesResponse() (response *DescribeRocketMQRolesResponse) {
	response = &DescribeRocketMQRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取角色列表
func (c *Client) DescribeRocketMQRoles(request *DescribeRocketMQRolesRequest) (response *DescribeRocketMQRolesResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQRolesRequest()
	}
	response = NewDescribeRocketMQRolesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAMQPQueueRequest() (request *CreateAMQPQueueRequest) {
	request = &CreateAMQPQueueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateAMQPQueue")
	return
}

func NewCreateAMQPQueueResponse() (response *CreateAMQPQueueResponse) {
	response = &CreateAMQPQueueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建AMQP队列
func (c *Client) CreateAMQPQueue(request *CreateAMQPQueueRequest) (response *CreateAMQPQueueResponse, err error) {
	if request == nil {
		request = NewCreateAMQPQueueRequest()
	}
	response = NewCreateAMQPQueueResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterExtraRequest() (request *DescribeClusterExtraRequest) {
	request = &DescribeClusterExtraRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterExtra")
	return
}

func NewDescribeClusterExtraResponse() (response *DescribeClusterExtraResponse) {
	response = &DescribeClusterExtraResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群其他信息
func (c *Client) DescribeClusterExtra(request *DescribeClusterExtraRequest) (response *DescribeClusterExtraResponse, err error) {
	if request == nil {
		request = NewDescribeClusterExtraRequest()
	}
	response = NewDescribeClusterExtraResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyMessageTTLOptRequest() (request *UpdatePolicyMessageTTLOptRequest) {
	request = &UpdatePolicyMessageTTLOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyMessageTTLOpt")
	return
}

func NewUpdatePolicyMessageTTLOptResponse() (response *UpdatePolicyMessageTTLOptResponse) {
	response = &UpdatePolicyMessageTTLOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端设置命名空间消息TTL
func (c *Client) UpdatePolicyMessageTTLOpt(request *UpdatePolicyMessageTTLOptRequest) (response *UpdatePolicyMessageTTLOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyMessageTTLOptRequest()
	}
	response = NewUpdatePolicyMessageTTLOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeResourceRulesOptRequest() (request *DescribeResourceRulesOptRequest) {
	request = &DescribeResourceRulesOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeResourceRulesOpt")
	return
}

func NewDescribeResourceRulesOptResponse() (response *DescribeResourceRulesOptResponse) {
	response = &DescribeResourceRulesOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取集群物理资源分配规则
func (c *Client) DescribeResourceRulesOpt(request *DescribeResourceRulesOptRequest) (response *DescribeResourceRulesOptResponse, err error) {
	if request == nil {
		request = NewDescribeResourceRulesOptRequest()
	}
	response = NewDescribeResourceRulesOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPClustersRequest() (request *DescribeAMQPClustersRequest) {
	request = &DescribeAMQPClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPClusters")
	return
}

func NewDescribeAMQPClustersResponse() (response *DescribeAMQPClustersResponse) {
	response = &DescribeAMQPClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取amqp集群列表
func (c *Client) DescribeAMQPClusters(request *DescribeAMQPClustersRequest) (response *DescribeAMQPClustersResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPClustersRequest()
	}
	response = NewDescribeAMQPClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQBillingUsageRequest() (request *DescribeRocketMQBillingUsageRequest) {
	request = &DescribeRocketMQBillingUsageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQBillingUsage")
	return
}

func NewDescribeRocketMQBillingUsageResponse() (response *DescribeRocketMQBillingUsageResponse) {
	response = &DescribeRocketMQBillingUsageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于获取当月RocketMQ使用量，如API调用次数
func (c *Client) DescribeRocketMQBillingUsage(request *DescribeRocketMQBillingUsageRequest) (response *DescribeRocketMQBillingUsageResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQBillingUsageRequest()
	}
	response = NewDescribeRocketMQBillingUsageResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAMQPVHostRequest() (request *ModifyAMQPVHostRequest) {
	request = &ModifyAMQPVHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyAMQPVHost")
	return
}

func NewModifyAMQPVHostResponse() (response *ModifyAMQPVHostResponse) {
	response = &ModifyAMQPVHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Vhost
func (c *Client) ModifyAMQPVHost(request *ModifyAMQPVHostRequest) (response *ModifyAMQPVHostResponse, err error) {
	if request == nil {
		request = NewModifyAMQPVHostRequest()
	}
	response = NewModifyAMQPVHostResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateTenantQuotaOptRequest() (request *UpdateTenantQuotaOptRequest) {
	request = &UpdateTenantQuotaOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdateTenantQuotaOpt")
	return
}

func NewUpdateTenantQuotaOptResponse() (response *UpdateTenantQuotaOptResponse) {
	response = &UpdateTenantQuotaOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端编辑虚拟集群（租户）配额
func (c *Client) UpdateTenantQuotaOpt(request *UpdateTenantQuotaOptRequest) (response *UpdateTenantQuotaOptResponse, err error) {
	if request == nil {
		request = NewUpdateTenantQuotaOptRequest()
	}
	response = NewUpdateTenantQuotaOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQEnvironmentRolesRequest() (request *DescribeRocketMQEnvironmentRolesRequest) {
	request = &DescribeRocketMQEnvironmentRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQEnvironmentRoles")
	return
}

func NewDescribeRocketMQEnvironmentRolesResponse() (response *DescribeRocketMQEnvironmentRolesResponse) {
	response = &DescribeRocketMQEnvironmentRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取命名空间角色列表
func (c *Client) DescribeRocketMQEnvironmentRoles(request *DescribeRocketMQEnvironmentRolesRequest) (response *DescribeRocketMQEnvironmentRolesResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQEnvironmentRolesRequest()
	}
	response = NewDescribeRocketMQEnvironmentRolesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRocketMQEnvironmentRoleRequest() (request *ModifyRocketMQEnvironmentRoleRequest) {
	request = &ModifyRocketMQEnvironmentRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQEnvironmentRole")
	return
}

func NewModifyRocketMQEnvironmentRoleResponse() (response *ModifyRocketMQEnvironmentRoleResponse) {
	response = &ModifyRocketMQEnvironmentRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改环境角色授权。
func (c *Client) ModifyRocketMQEnvironmentRole(request *ModifyRocketMQEnvironmentRoleRequest) (response *ModifyRocketMQEnvironmentRoleResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQEnvironmentRoleRequest()
	}
	response = NewModifyRocketMQEnvironmentRoleResponse()
	err = c.Send(request, response)
	return
}

func NewAddPulsarResourceAllocationRuleRequest() (request *AddPulsarResourceAllocationRuleRequest) {
	request = &AddPulsarResourceAllocationRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "AddPulsarResourceAllocationRule")
	return
}

func NewAddPulsarResourceAllocationRuleResponse() (response *AddPulsarResourceAllocationRuleResponse) {
	response = &AddPulsarResourceAllocationRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 增加集群物理资源分配规则
func (c *Client) AddPulsarResourceAllocationRule(request *AddPulsarResourceAllocationRuleRequest) (response *AddPulsarResourceAllocationRuleResponse, err error) {
	if request == nil {
		request = NewAddPulsarResourceAllocationRuleRequest()
	}
	response = NewAddPulsarResourceAllocationRuleResponse()
	err = c.Send(request, response)
	return
}

func NewSendBatchMessagesRequest() (request *SendBatchMessagesRequest) {
	request = &SendBatchMessagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SendBatchMessages")
	return
}

func NewSendBatchMessagesResponse() (response *SendBatchMessagesResponse) {
	response = &SendBatchMessagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量发送消息
func (c *Client) SendBatchMessages(request *SendBatchMessagesRequest) (response *SendBatchMessagesResponse, err error) {
	if request == nil {
		request = NewSendBatchMessagesRequest()
	}
	response = NewSendBatchMessagesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAMQPRouteRelationRequest() (request *CreateAMQPRouteRelationRequest) {
	request = &CreateAMQPRouteRelationRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateAMQPRouteRelation")
	return
}

func NewCreateAMQPRouteRelationResponse() (response *CreateAMQPRouteRelationResponse) {
	response = &CreateAMQPRouteRelationResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建AMQP路由关系
func (c *Client) CreateAMQPRouteRelation(request *CreateAMQPRouteRelationRequest) (response *CreateAMQPRouteRelationResponse, err error) {
	if request == nil {
		request = NewCreateAMQPRouteRelationRequest()
	}
	response = NewCreateAMQPRouteRelationResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAllTemplateModuleOptRequest() (request *DescribeAllTemplateModuleOptRequest) {
	request = &DescribeAllTemplateModuleOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAllTemplateModuleOpt")
	return
}

func NewDescribeAllTemplateModuleOptResponse() (response *DescribeAllTemplateModuleOptResponse) {
	response = &DescribeAllTemplateModuleOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看一套模板的所有配置模块
func (c *Client) DescribeAllTemplateModuleOpt(request *DescribeAllTemplateModuleOptRequest) (response *DescribeAllTemplateModuleOptResponse, err error) {
	if request == nil {
		request = NewDescribeAllTemplateModuleOptRequest()
	}
	response = NewDescribeAllTemplateModuleOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTemplateItemOptRequest() (request *DescribeTemplateItemOptRequest) {
	request = &DescribeTemplateItemOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeTemplateItemOpt")
	return
}

func NewDescribeTemplateItemOptResponse() (response *DescribeTemplateItemOptResponse) {
	response = &DescribeTemplateItemOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询模板配置项
func (c *Client) DescribeTemplateItemOpt(request *DescribeTemplateItemOptRequest) (response *DescribeTemplateItemOptResponse, err error) {
	if request == nil {
		request = NewDescribeTemplateItemOptRequest()
	}
	response = NewDescribeTemplateItemOptResponse()
	err = c.Send(request, response)
	return
}

func NewSetRocketMQPublicAccessPointRequest() (request *SetRocketMQPublicAccessPointRequest) {
	request = &SetRocketMQPublicAccessPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SetRocketMQPublicAccessPoint")
	return
}

func NewSetRocketMQPublicAccessPointResponse() (response *SetRocketMQPublicAccessPointResponse) {
	response = &SetRocketMQPublicAccessPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 该接口用于开启关闭公网访问、设置安全访问策略
func (c *Client) SetRocketMQPublicAccessPoint(request *SetRocketMQPublicAccessPointRequest) (response *SetRocketMQPublicAccessPointResponse, err error) {
	if request == nil {
		request = NewSetRocketMQPublicAccessPointRequest()
	}
	response = NewSetRocketMQPublicAccessPointResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBatchRoleRequest() (request *CreateBatchRoleRequest) {
	request = &CreateBatchRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateBatchRole")
	return
}

func NewCreateBatchRoleResponse() (response *CreateBatchRoleResponse) {
	response = &CreateBatchRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入角色
func (c *Client) CreateBatchRole(request *CreateBatchRoleRequest) (response *CreateBatchRoleResponse, err error) {
	if request == nil {
		request = NewCreateBatchRoleRequest()
	}
	response = NewCreateBatchRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBindClustersRequest() (request *DescribeBindClustersRequest) {
	request = &DescribeBindClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBindClusters")
	return
}

func NewDescribeBindClustersResponse() (response *DescribeBindClustersResponse) {
	response = &DescribeBindClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户绑定的专享集群列表
func (c *Client) DescribeBindClusters(request *DescribeBindClustersRequest) (response *DescribeBindClustersResponse, err error) {
	if request == nil {
		request = NewDescribeBindClustersRequest()
	}
	response = NewDescribeBindClustersResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicySchemaCompatibilityStrategyOptRequest() (request *UpdatePolicySchemaCompatibilityStrategyOptRequest) {
	request = &UpdatePolicySchemaCompatibilityStrategyOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicySchemaCompatibilityStrategyOpt")
	return
}

func NewUpdatePolicySchemaCompatibilityStrategyOptResponse() (response *UpdatePolicySchemaCompatibilityStrategyOptResponse) {
	response = &UpdatePolicySchemaCompatibilityStrategyOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间Schema自动更新策略
func (c *Client) UpdatePolicySchemaCompatibilityStrategyOpt(request *UpdatePolicySchemaCompatibilityStrategyOptRequest) (response *UpdatePolicySchemaCompatibilityStrategyOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicySchemaCompatibilityStrategyOptRequest()
	}
	response = NewUpdatePolicySchemaCompatibilityStrategyOptResponse()
	err = c.Send(request, response)
	return
}

func NewClearCmqQueueRequest() (request *ClearCmqQueueRequest) {
	request = &ClearCmqQueueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ClearCmqQueue")
	return
}

func NewClearCmqQueueResponse() (response *ClearCmqQueueResponse) {
	response = &ClearCmqQueueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清空cmq消息队列中的消息
func (c *Client) ClearCmqQueue(request *ClearCmqQueueRequest) (response *ClearCmqQueueResponse, err error) {
	if request == nil {
		request = NewClearCmqQueueRequest()
	}
	response = NewClearCmqQueueResponse()
	err = c.Send(request, response)
	return
}

func NewReceiveMessageRequest() (request *ReceiveMessageRequest) {
	request = &ReceiveMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ReceiveMessage")
	return
}

func NewReceiveMessageResponse() (response *ReceiveMessageResponse) {
	response = &ReceiveMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接收发送到指定 topic 中的消息
func (c *Client) ReceiveMessage(request *ReceiveMessageRequest) (response *ReceiveMessageResponse, err error) {
	if request == nil {
		request = NewReceiveMessageRequest()
	}
	response = NewReceiveMessageResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicySubscribeRateOptRequest() (request *UpdatePolicySubscribeRateOptRequest) {
	request = &UpdatePolicySubscribeRateOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicySubscribeRateOpt")
	return
}

func NewUpdatePolicySubscribeRateOptResponse() (response *UpdatePolicySubscribeRateOptResponse) {
	response = &UpdatePolicySubscribeRateOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间消费者级别限制请求分发策略限制
func (c *Client) UpdatePolicySubscribeRateOpt(request *UpdatePolicySubscribeRateOptRequest) (response *UpdatePolicySubscribeRateOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicySubscribeRateOptRequest()
	}
	response = NewUpdatePolicySubscribeRateOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCmqTopicsRequest() (request *DescribeCmqTopicsRequest) {
	request = &DescribeCmqTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqTopics")
	return
}

func NewDescribeCmqTopicsResponse() (response *DescribeCmqTopicsResponse) {
	response = &DescribeCmqTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 枚举cmq全量主题
func (c *Client) DescribeCmqTopics(request *DescribeCmqTopicsRequest) (response *DescribeCmqTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeCmqTopicsRequest()
	}
	response = NewDescribeCmqTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRocketMQGroupRequest() (request *ModifyRocketMQGroupRequest) {
	request = &ModifyRocketMQGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQGroup")
	return
}

func NewModifyRocketMQGroupResponse() (response *ModifyRocketMQGroupResponse) {
	response = &ModifyRocketMQGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新RocketMQ消费组信息
func (c *Client) ModifyRocketMQGroup(request *ModifyRocketMQGroupRequest) (response *ModifyRocketMQGroupResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQGroupRequest()
	}
	response = NewModifyRocketMQGroupResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicySubscriptionAuthModeOptRequest() (request *UpdatePolicySubscriptionAuthModeOptRequest) {
	request = &UpdatePolicySubscriptionAuthModeOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicySubscriptionAuthModeOpt")
	return
}

func NewUpdatePolicySubscriptionAuthModeOptResponse() (response *UpdatePolicySubscriptionAuthModeOptResponse) {
	response = &UpdatePolicySubscriptionAuthModeOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间订阅鉴权策略
func (c *Client) UpdatePolicySubscriptionAuthModeOpt(request *UpdatePolicySubscriptionAuthModeOptRequest) (response *UpdatePolicySubscriptionAuthModeOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicySubscriptionAuthModeOptRequest()
	}
	response = NewUpdatePolicySubscriptionAuthModeOptResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTopicsRequest() (request *DeleteTopicsRequest) {
	request = &DeleteTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteTopics")
	return
}

func NewDeleteTopicsResponse() (response *DeleteTopicsResponse) {
	response = &DeleteTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除topics
func (c *Client) DeleteTopics(request *DeleteTopicsRequest) (response *DeleteTopicsResponse, err error) {
	if request == nil {
		request = NewDeleteTopicsRequest()
	}
	response = NewDeleteTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewSetClusterOfflineRequest() (request *SetClusterOfflineRequest) {
	request = &SetClusterOfflineRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SetClusterOffline")
	return
}

func NewSetClusterOfflineResponse() (response *SetClusterOfflineResponse) {
	response = &SetClusterOfflineResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 集群下线接口
func (c *Client) SetClusterOffline(request *SetClusterOfflineRequest) (response *SetClusterOfflineResponse, err error) {
	if request == nil {
		request = NewSetClusterOfflineRequest()
	}
	response = NewSetClusterOfflineResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPVHostConnectionsRequest() (request *DescribeAMQPVHostConnectionsRequest) {
	request = &DescribeAMQPVHostConnectionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPVHostConnections")
	return
}

func NewDescribeAMQPVHostConnectionsResponse() (response *DescribeAMQPVHostConnectionsResponse) {
	response = &DescribeAMQPVHostConnectionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Amqp Vhost连接列表
func (c *Client) DescribeAMQPVHostConnections(request *DescribeAMQPVHostConnectionsRequest) (response *DescribeAMQPVHostConnectionsResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPVHostConnectionsRequest()
	}
	response = NewDescribeAMQPVHostConnectionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePortalConfigRequest() (request *DescribePortalConfigRequest) {
	request = &DescribePortalConfigRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribePortalConfig")
	return
}

func NewDescribePortalConfigResponse() (response *DescribePortalConfigResponse) {
	response = &DescribePortalConfigResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取管理端页面配置
func (c *Client) DescribePortalConfig(request *DescribePortalConfigRequest) (response *DescribePortalConfigResponse, err error) {
	if request == nil {
		request = NewDescribePortalConfigRequest()
	}
	response = NewDescribePortalConfigResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCmqSubscriptionDetailRequest() (request *DescribeCmqSubscriptionDetailRequest) {
	request = &DescribeCmqSubscriptionDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqSubscriptionDetail")
	return
}

func NewDescribeCmqSubscriptionDetailResponse() (response *DescribeCmqSubscriptionDetailResponse) {
	response = &DescribeCmqSubscriptionDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cmq订阅详情
func (c *Client) DescribeCmqSubscriptionDetail(request *DescribeCmqSubscriptionDetailRequest) (response *DescribeCmqSubscriptionDetailResponse, err error) {
	if request == nil {
		request = NewDescribeCmqSubscriptionDetailRequest()
	}
	response = NewDescribeCmqSubscriptionDetailResponse()
	err = c.Send(request, response)
	return
}

func NewTpoDescribeResourcesAdminProjectsRequest() (request *TpoDescribeResourcesAdminProjectsRequest) {
	request = &TpoDescribeResourcesAdminProjectsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "TpoDescribeResourcesAdminProjects")
	return
}

func NewTpoDescribeResourcesAdminProjectsResponse() (response *TpoDescribeResourcesAdminProjectsResponse) {
	response = &TpoDescribeResourcesAdminProjectsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 通过获取项目权限列表，用于用户集群项目归属的编辑（添加到某个项目，或者从某个项目变更到另一个项目）
func (c *Client) TpoDescribeResourcesAdminProjects(request *TpoDescribeResourcesAdminProjectsRequest) (response *TpoDescribeResourcesAdminProjectsResponse, err error) {
	if request == nil {
		request = NewTpoDescribeResourcesAdminProjectsRequest()
	}
	response = NewTpoDescribeResourcesAdminProjectsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQVipInstancesRequest() (request *DescribeRocketMQVipInstancesRequest) {
	request = &DescribeRocketMQVipInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQVipInstances")
	return
}

func NewDescribeRocketMQVipInstancesResponse() (response *DescribeRocketMQVipInstancesResponse) {
	response = &DescribeRocketMQVipInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户已购的RocketMQ专享实例列表
func (c *Client) DescribeRocketMQVipInstances(request *DescribeRocketMQVipInstancesRequest) (response *DescribeRocketMQVipInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQVipInstancesRequest()
	}
	response = NewDescribeRocketMQVipInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAMQPExchangeRequest() (request *DeleteAMQPExchangeRequest) {
	request = &DeleteAMQPExchangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteAMQPExchange")
	return
}

func NewDeleteAMQPExchangeResponse() (response *DeleteAMQPExchangeResponse) {
	response = &DeleteAMQPExchangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Amqp交换机
func (c *Client) DeleteAMQPExchange(request *DeleteAMQPExchangeRequest) (response *DeleteAMQPExchangeResponse, err error) {
	if request == nil {
		request = NewDeleteAMQPExchangeRequest()
	}
	response = NewDeleteAMQPExchangeResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCmqTopicRequest() (request *DeleteCmqTopicRequest) {
	request = &DeleteCmqTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqTopic")
	return
}

func NewDeleteCmqTopicResponse() (response *DeleteCmqTopicResponse) {
	response = &DeleteCmqTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除cmq主题
func (c *Client) DeleteCmqTopic(request *DeleteCmqTopicRequest) (response *DeleteCmqTopicResponse, err error) {
	if request == nil {
		request = NewDeleteCmqTopicRequest()
	}
	response = NewDeleteCmqTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterNodesOptRequest() (request *DescribeClusterNodesOptRequest) {
	request = &DescribeClusterNodesOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterNodesOpt")
	return
}

func NewDescribeClusterNodesOptResponse() (response *DescribeClusterNodesOptResponse) {
	response = &DescribeClusterNodesOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取集群节点列表
func (c *Client) DescribeClusterNodesOpt(request *DescribeClusterNodesOptRequest) (response *DescribeClusterNodesOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterNodesOptRequest()
	}
	response = NewDescribeClusterNodesOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterOptRequest() (request *CreateClusterOptRequest) {
	request = &CreateClusterOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateClusterOpt")
	return
}

func NewCreateClusterOptResponse() (response *CreateClusterOptResponse) {
	response = &CreateClusterOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端创建物理集群
func (c *Client) CreateClusterOpt(request *CreateClusterOptRequest) (response *CreateClusterOptResponse, err error) {
	if request == nil {
		request = NewCreateClusterOptRequest()
	}
	response = NewCreateClusterOptResponse()
	err = c.Send(request, response)
	return
}

func NewInitializeClusterConfigOptRequest() (request *InitializeClusterConfigOptRequest) {
	request = &InitializeClusterConfigOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "InitializeClusterConfigOpt")
	return
}

func NewInitializeClusterConfigOptResponse() (response *InitializeClusterConfigOptResponse) {
	response = &InitializeClusterConfigOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 初始化集群配置
func (c *Client) InitializeClusterConfigOpt(request *InitializeClusterConfigOptRequest) (response *InitializeClusterConfigOptResponse, err error) {
	if request == nil {
		request = NewInitializeClusterConfigOptRequest()
	}
	response = NewInitializeClusterConfigOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQGroupsRequest() (request *DescribeRocketMQGroupsRequest) {
	request = &DescribeRocketMQGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQGroups")
	return
}

func NewDescribeRocketMQGroupsResponse() (response *DescribeRocketMQGroupsResponse) {
	response = &DescribeRocketMQGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取RocketMQ消费组列表
func (c *Client) DescribeRocketMQGroups(request *DescribeRocketMQGroupsRequest) (response *DescribeRocketMQGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQGroupsRequest()
	}
	response = NewDescribeRocketMQGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQNamespacesRequest() (request *DescribeRocketMQNamespacesRequest) {
	request = &DescribeRocketMQNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQNamespaces")
	return
}

func NewDescribeRocketMQNamespacesResponse() (response *DescribeRocketMQNamespacesResponse) {
	response = &DescribeRocketMQNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取RocketMQ命名空间列表
func (c *Client) DescribeRocketMQNamespaces(request *DescribeRocketMQNamespacesRequest) (response *DescribeRocketMQNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQNamespacesRequest()
	}
	response = NewDescribeRocketMQNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyMaxConsumerPerSubscriptionOptRequest() (request *UpdatePolicyMaxConsumerPerSubscriptionOptRequest) {
	request = &UpdatePolicyMaxConsumerPerSubscriptionOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyMaxConsumerPerSubscriptionOpt")
	return
}

func NewUpdatePolicyMaxConsumerPerSubscriptionOptResponse() (response *UpdatePolicyMaxConsumerPerSubscriptionOptResponse) {
	response = &UpdatePolicyMaxConsumerPerSubscriptionOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间订阅级别消费者数量限制
func (c *Client) UpdatePolicyMaxConsumerPerSubscriptionOpt(request *UpdatePolicyMaxConsumerPerSubscriptionOptRequest) (response *UpdatePolicyMaxConsumerPerSubscriptionOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyMaxConsumerPerSubscriptionOptRequest()
	}
	response = NewUpdatePolicyMaxConsumerPerSubscriptionOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterOverBookieOptRequest() (request *DescribeClusterOverBookieOptRequest) {
	request = &DescribeClusterOverBookieOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterOverBookieOpt")
	return
}

func NewDescribeClusterOverBookieOptResponse() (response *DescribeClusterOverBookieOptResponse) {
	response = &DescribeClusterOverBookieOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取bookie支撑的物理集群消耗情况列表
func (c *Client) DescribeClusterOverBookieOpt(request *DescribeClusterOverBookieOptRequest) (response *DescribeClusterOverBookieOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterOverBookieOptRequest()
	}
	response = NewDescribeClusterOverBookieOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInternalRocketMQInstancesRequest() (request *DescribeInternalRocketMQInstancesRequest) {
	request = &DescribeInternalRocketMQInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeInternalRocketMQInstances")
	return
}

func NewDescribeInternalRocketMQInstancesResponse() (response *DescribeInternalRocketMQInstancesResponse) {
	response = &DescribeInternalRocketMQInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此为内部接口，用于运营端查询用户RocketMQ实例信息
func (c *Client) DescribeInternalRocketMQInstances(request *DescribeInternalRocketMQInstancesRequest) (response *DescribeInternalRocketMQInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeInternalRocketMQInstancesRequest()
	}
	response = NewDescribeInternalRocketMQInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQPublicAccessPointRequest() (request *DescribeRocketMQPublicAccessPointRequest) {
	request = &DescribeRocketMQPublicAccessPointRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQPublicAccessPoint")
	return
}

func NewDescribeRocketMQPublicAccessPointResponse() (response *DescribeRocketMQPublicAccessPointResponse) {
	response = &DescribeRocketMQPublicAccessPointResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 接口用于查询RocketMQ实例的公网接入信息
func (c *Client) DescribeRocketMQPublicAccessPoint(request *DescribeRocketMQPublicAccessPointRequest) (response *DescribeRocketMQPublicAccessPointResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQPublicAccessPointRequest()
	}
	response = NewDescribeRocketMQPublicAccessPointResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQVipInstanceDetailRequest() (request *DescribeRocketMQVipInstanceDetailRequest) {
	request = &DescribeRocketMQVipInstanceDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQVipInstanceDetail")
	return
}

func NewDescribeRocketMQVipInstanceDetailResponse() (response *DescribeRocketMQVipInstanceDetailResponse) {
	response = &DescribeRocketMQVipInstanceDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个RocketMQ专享集群信息
func (c *Client) DescribeRocketMQVipInstanceDetail(request *DescribeRocketMQVipInstanceDetailRequest) (response *DescribeRocketMQVipInstanceDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQVipInstanceDetailRequest()
	}
	response = NewDescribeRocketMQVipInstanceDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCmqQueueRequest() (request *DeleteCmqQueueRequest) {
	request = &DeleteCmqQueueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqQueue")
	return
}

func NewDeleteCmqQueueResponse() (response *DeleteCmqQueueResponse) {
	response = &DeleteCmqQueueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除cmq队列
func (c *Client) DeleteCmqQueue(request *DeleteCmqQueueRequest) (response *DeleteCmqQueueResponse, err error) {
	if request == nil {
		request = NewDeleteCmqQueueRequest()
	}
	response = NewDeleteCmqQueueResponse()
	err = c.Send(request, response)
	return
}

func NewClearCmqSubscriptionFilterTagsRequest() (request *ClearCmqSubscriptionFilterTagsRequest) {
	request = &ClearCmqSubscriptionFilterTagsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ClearCmqSubscriptionFilterTags")
	return
}

func NewClearCmqSubscriptionFilterTagsResponse() (response *ClearCmqSubscriptionFilterTagsResponse) {
	response = &ClearCmqSubscriptionFilterTagsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 清空订阅者消息标签
func (c *Client) ClearCmqSubscriptionFilterTags(request *ClearCmqSubscriptionFilterTagsRequest) (response *ClearCmqSubscriptionFilterTagsResponse, err error) {
	if request == nil {
		request = NewClearCmqSubscriptionFilterTagsRequest()
	}
	response = NewClearCmqSubscriptionFilterTagsResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyDeduplicationEnabledOptRequest() (request *UpdatePolicyDeduplicationEnabledOptRequest) {
	request = &UpdatePolicyDeduplicationEnabledOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyDeduplicationEnabledOpt")
	return
}

func NewUpdatePolicyDeduplicationEnabledOptResponse() (response *UpdatePolicyDeduplicationEnabledOptResponse) {
	response = &UpdatePolicyDeduplicationEnabledOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间重复数据清除策略
func (c *Client) UpdatePolicyDeduplicationEnabledOpt(request *UpdatePolicyDeduplicationEnabledOptRequest) (response *UpdatePolicyDeduplicationEnabledOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyDeduplicationEnabledOptRequest()
	}
	response = NewUpdatePolicyDeduplicationEnabledOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeEnvironmentRolesRequest() (request *DescribeEnvironmentRolesRequest) {
	request = &DescribeEnvironmentRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeEnvironmentRoles")
	return
}

func NewDescribeEnvironmentRolesResponse() (response *DescribeEnvironmentRolesResponse) {
	response = &DescribeEnvironmentRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取命名空间角色列表
func (c *Client) DescribeEnvironmentRoles(request *DescribeEnvironmentRolesRequest) (response *DescribeEnvironmentRolesResponse, err error) {
	if request == nil {
		request = NewDescribeEnvironmentRolesRequest()
	}
	response = NewDescribeEnvironmentRolesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQConsumerClientsRequest() (request *DescribeRocketMQConsumerClientsRequest) {
	request = &DescribeRocketMQConsumerClientsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQConsumerClients")
	return
}

func NewDescribeRocketMQConsumerClientsResponse() (response *DescribeRocketMQConsumerClientsResponse) {
	response = &DescribeRocketMQConsumerClientsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定消费组下当前在线消费者客户端信息
func (c *Client) DescribeRocketMQConsumerClients(request *DescribeRocketMQConsumerClientsRequest) (response *DescribeRocketMQConsumerClientsResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQConsumerClientsRequest()
	}
	response = NewDescribeRocketMQConsumerClientsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeReplicationClustersRequest() (request *DescribeReplicationClustersRequest) {
	request = &DescribeReplicationClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeReplicationClusters")
	return
}

func NewDescribeReplicationClustersResponse() (response *DescribeReplicationClustersResponse) {
	response = &DescribeReplicationClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询跨地域复制目标集群
func (c *Client) DescribeReplicationClusters(request *DescribeReplicationClustersRequest) (response *DescribeReplicationClustersResponse, err error) {
	if request == nil {
		request = NewDescribeReplicationClustersRequest()
	}
	response = NewDescribeReplicationClustersResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicySchemaValidationEnforcedRequest() (request *UpdatePolicySchemaValidationEnforcedRequest) {
	request = &UpdatePolicySchemaValidationEnforcedRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicySchemaValidationEnforced")
	return
}

func NewUpdatePolicySchemaValidationEnforcedResponse() (response *UpdatePolicySchemaValidationEnforcedResponse) {
	response = &UpdatePolicySchemaValidationEnforcedResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间Schema是否开启强校验
func (c *Client) UpdatePolicySchemaValidationEnforced(request *UpdatePolicySchemaValidationEnforcedRequest) (response *UpdatePolicySchemaValidationEnforcedResponse, err error) {
	if request == nil {
		request = NewUpdatePolicySchemaValidationEnforcedRequest()
	}
	response = NewUpdatePolicySchemaValidationEnforcedResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInternalRocketMQNamespacesRequest() (request *DescribeInternalRocketMQNamespacesRequest) {
	request = &DescribeInternalRocketMQNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeInternalRocketMQNamespaces")
	return
}

func NewDescribeInternalRocketMQNamespacesResponse() (response *DescribeInternalRocketMQNamespacesResponse) {
	response = &DescribeInternalRocketMQNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口为内部接口，用于运营端查询命名空间列表
func (c *Client) DescribeInternalRocketMQNamespaces(request *DescribeInternalRocketMQNamespacesRequest) (response *DescribeInternalRocketMQNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeInternalRocketMQNamespacesRequest()
	}
	response = NewDescribeInternalRocketMQNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterComponentMetricsOptRequest() (request *DescribeClusterComponentMetricsOptRequest) {
	request = &DescribeClusterComponentMetricsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterComponentMetricsOpt")
	return
}

func NewDescribeClusterComponentMetricsOptResponse() (response *DescribeClusterComponentMetricsOptResponse) {
	response = &DescribeClusterComponentMetricsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取组件监控指标值
func (c *Client) DescribeClusterComponentMetricsOpt(request *DescribeClusterComponentMetricsOptRequest) (response *DescribeClusterComponentMetricsOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterComponentMetricsOptRequest()
	}
	response = NewDescribeClusterComponentMetricsOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCmqQueuesRequest() (request *DescribeCmqQueuesRequest) {
	request = &DescribeCmqQueuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqQueues")
	return
}

func NewDescribeCmqQueuesResponse() (response *DescribeCmqQueuesResponse) {
	response = &DescribeCmqQueuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cmq全量队列
func (c *Client) DescribeCmqQueues(request *DescribeCmqQueuesRequest) (response *DescribeCmqQueuesResponse, err error) {
	if request == nil {
		request = NewDescribeCmqQueuesRequest()
	}
	response = NewDescribeCmqQueuesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateTemplateOptRequest() (request *CreateTemplateOptRequest) {
	request = &CreateTemplateOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateTemplateOpt")
	return
}

func NewCreateTemplateOptResponse() (response *CreateTemplateOptResponse) {
	response = &CreateTemplateOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建配置模板
func (c *Client) CreateTemplateOpt(request *CreateTemplateOptRequest) (response *CreateTemplateOptResponse, err error) {
	if request == nil {
		request = NewCreateTemplateOptRequest()
	}
	response = NewCreateTemplateOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeConfigDictOptRequest() (request *DescribeConfigDictOptRequest) {
	request = &DescribeConfigDictOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeConfigDictOpt")
	return
}

func NewDescribeConfigDictOptResponse() (response *DescribeConfigDictOptResponse) {
	response = &DescribeConfigDictOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取配置字典表信息
func (c *Client) DescribeConfigDictOpt(request *DescribeConfigDictOptRequest) (response *DescribeConfigDictOptResponse, err error) {
	if request == nil {
		request = NewDescribeConfigDictOptRequest()
	}
	response = NewDescribeConfigDictOptResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRocketMQClusterRequest() (request *ModifyRocketMQClusterRequest) {
	request = &ModifyRocketMQClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQCluster")
	return
}

func NewModifyRocketMQClusterResponse() (response *ModifyRocketMQClusterResponse) {
	response = &ModifyRocketMQClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新RocketMQ集群信息
func (c *Client) ModifyRocketMQCluster(request *ModifyRocketMQClusterRequest) (response *ModifyRocketMQClusterResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQClusterRequest()
	}
	response = NewModifyRocketMQClusterResponse()
	err = c.Send(request, response)
	return
}

func NewSendConfigOptRequest() (request *SendConfigOptRequest) {
	request = &SendConfigOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SendConfigOpt")
	return
}

func NewSendConfigOptResponse() (response *SendConfigOptResponse) {
	response = &SendConfigOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 配置下发
func (c *Client) SendConfigOpt(request *SendConfigOptRequest) (response *SendConfigOptResponse, err error) {
	if request == nil {
		request = NewSendConfigOptRequest()
	}
	response = NewSendConfigOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInternalNamespacesRequest() (request *DescribeInternalNamespacesRequest) {
	request = &DescribeInternalNamespacesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeInternalNamespaces")
	return
}

func NewDescribeInternalNamespacesResponse() (response *DescribeInternalNamespacesResponse) {
	response = &DescribeInternalNamespacesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取命名空间列表
func (c *Client) DescribeInternalNamespaces(request *DescribeInternalNamespacesRequest) (response *DescribeInternalNamespacesResponse, err error) {
	if request == nil {
		request = NewDescribeInternalNamespacesRequest()
	}
	response = NewDescribeInternalNamespacesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterBrokerNetworkOptRequest() (request *DescribeClusterBrokerNetworkOptRequest) {
	request = &DescribeClusterBrokerNetworkOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterBrokerNetworkOpt")
	return
}

func NewDescribeClusterBrokerNetworkOptResponse() (response *DescribeClusterBrokerNetworkOptResponse) {
	response = &DescribeClusterBrokerNetworkOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取集群网络详情broker网络映射
func (c *Client) DescribeClusterBrokerNetworkOpt(request *DescribeClusterBrokerNetworkOptRequest) (response *DescribeClusterBrokerNetworkOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterBrokerNetworkOptRequest()
	}
	response = NewDescribeClusterBrokerNetworkOptResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyMessageRetentionOptRequest() (request *UpdatePolicyMessageRetentionOptRequest) {
	request = &UpdatePolicyMessageRetentionOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyMessageRetentionOpt")
	return
}

func NewUpdatePolicyMessageRetentionOptResponse() (response *UpdatePolicyMessageRetentionOptResponse) {
	response = &UpdatePolicyMessageRetentionOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间消息保留策略
func (c *Client) UpdatePolicyMessageRetentionOpt(request *UpdatePolicyMessageRetentionOptRequest) (response *UpdatePolicyMessageRetentionOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyMessageRetentionOptRequest()
	}
	response = NewUpdatePolicyMessageRetentionOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateCmqSubscribeRequest() (request *CreateCmqSubscribeRequest) {
	request = &CreateCmqSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateCmqSubscribe")
	return
}

func NewCreateCmqSubscribeResponse() (response *CreateCmqSubscribeResponse) {
	response = &CreateCmqSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建cmq订阅接口
func (c *Client) CreateCmqSubscribe(request *CreateCmqSubscribeRequest) (response *CreateCmqSubscribeResponse, err error) {
	if request == nil {
		request = NewCreateCmqSubscribeRequest()
	}
	response = NewCreateCmqSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQSubscriptionsRequest() (request *DescribeRocketMQSubscriptionsRequest) {
	request = &DescribeRocketMQSubscriptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQSubscriptions")
	return
}

func NewDescribeRocketMQSubscriptionsResponse() (response *DescribeRocketMQSubscriptionsResponse) {
	response = &DescribeRocketMQSubscriptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于获取RocketMQ消费组订阅关系数据
func (c *Client) DescribeRocketMQSubscriptions(request *DescribeRocketMQSubscriptionsRequest) (response *DescribeRocketMQSubscriptionsResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQSubscriptionsRequest()
	}
	response = NewDescribeRocketMQSubscriptionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAMQPClusterRequest() (request *ModifyAMQPClusterRequest) {
	request = &ModifyAMQPClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyAMQPCluster")
	return
}

func NewModifyAMQPClusterResponse() (response *ModifyAMQPClusterResponse) {
	response = &ModifyAMQPClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Amqp集群信息
func (c *Client) ModifyAMQPCluster(request *ModifyAMQPClusterRequest) (response *ModifyAMQPClusterResponse, err error) {
	if request == nil {
		request = NewModifyAMQPClusterRequest()
	}
	response = NewModifyAMQPClusterResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRocketMQRoleRequest() (request *ModifyRocketMQRoleRequest) {
	request = &ModifyRocketMQRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQRole")
	return
}

func NewModifyRocketMQRoleResponse() (response *ModifyRocketMQRoleResponse) {
	response = &ModifyRocketMQRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 角色修改
func (c *Client) ModifyRocketMQRole(request *ModifyRocketMQRoleRequest) (response *ModifyRocketMQRoleResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQRoleRequest()
	}
	response = NewModifyRocketMQRoleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdateClusterOptRequest() (request *UpdateClusterOptRequest) {
	request = &UpdateClusterOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdateClusterOpt")
	return
}

func NewUpdateClusterOptResponse() (response *UpdateClusterOptResponse) {
	response = &UpdateClusterOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新物理集群
func (c *Client) UpdateClusterOpt(request *UpdateClusterOptRequest) (response *UpdateClusterOptResponse, err error) {
	if request == nil {
		request = NewUpdateClusterOptRequest()
	}
	response = NewUpdateClusterOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateConfigOptRequest() (request *CreateConfigOptRequest) {
	request = &CreateConfigOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateConfigOpt")
	return
}

func NewCreateConfigOptResponse() (response *CreateConfigOptResponse) {
	response = &CreateConfigOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建配置
func (c *Client) CreateConfigOpt(request *CreateConfigOptRequest) (response *CreateConfigOptResponse, err error) {
	if request == nil {
		request = NewCreateConfigOptRequest()
	}
	response = NewCreateConfigOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeModuleListOptRequest() (request *DescribeModuleListOptRequest) {
	request = &DescribeModuleListOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeModuleListOpt")
	return
}

func NewDescribeModuleListOptResponse() (response *DescribeModuleListOptResponse) {
	response = &DescribeModuleListOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看配置模块列表
func (c *Client) DescribeModuleListOpt(request *DescribeModuleListOptRequest) (response *DescribeModuleListOptResponse, err error) {
	if request == nil {
		request = NewDescribeModuleListOptRequest()
	}
	response = NewDescribeModuleListOptResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEnvironmentRoleRequest() (request *ModifyEnvironmentRoleRequest) {
	request = &ModifyEnvironmentRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyEnvironmentRole")
	return
}

func NewModifyEnvironmentRoleResponse() (response *ModifyEnvironmentRoleResponse) {
	response = &ModifyEnvironmentRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改环境角色授权。
func (c *Client) ModifyEnvironmentRole(request *ModifyEnvironmentRoleRequest) (response *ModifyEnvironmentRoleResponse, err error) {
	if request == nil {
		request = NewModifyEnvironmentRoleRequest()
	}
	response = NewModifyEnvironmentRoleResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyEncryptionRequiredOptRequest() (request *UpdatePolicyEncryptionRequiredOptRequest) {
	request = &UpdatePolicyEncryptionRequiredOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyEncryptionRequiredOpt")
	return
}

func NewUpdatePolicyEncryptionRequiredOptResponse() (response *UpdatePolicyEncryptionRequiredOptResponse) {
	response = &UpdatePolicyEncryptionRequiredOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间加密开关策略
func (c *Client) UpdatePolicyEncryptionRequiredOpt(request *UpdatePolicyEncryptionRequiredOptRequest) (response *UpdatePolicyEncryptionRequiredOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyEncryptionRequiredOptRequest()
	}
	response = NewUpdatePolicyEncryptionRequiredOptResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteClusterRequest() (request *DeleteClusterRequest) {
	request = &DeleteClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCluster")
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

func NewModifyRocketMQTopicRequest() (request *ModifyRocketMQTopicRequest) {
	request = &ModifyRocketMQTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQTopic")
	return
}

func NewModifyRocketMQTopicResponse() (response *ModifyRocketMQTopicResponse) {
	response = &ModifyRocketMQTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新RocketMQ主题信息
func (c *Client) ModifyRocketMQTopic(request *ModifyRocketMQTopicRequest) (response *ModifyRocketMQTopicResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQTopicRequest()
	}
	response = NewModifyRocketMQTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRouteRequest() (request *DescribeRouteRequest) {
	request = &DescribeRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRoute")
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

func NewUpdateResourceRuleOptRequest() (request *UpdateResourceRuleOptRequest) {
	request = &UpdateResourceRuleOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdateResourceRuleOpt")
	return
}

func NewUpdateResourceRuleOptResponse() (response *UpdateResourceRuleOptResponse) {
	response = &UpdateResourceRuleOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端创建集群物理资源分配规则
func (c *Client) UpdateResourceRuleOpt(request *UpdateResourceRuleOptRequest) (response *UpdateResourceRuleOptResponse, err error) {
	if request == nil {
		request = NewUpdateResourceRuleOptRequest()
	}
	response = NewUpdateResourceRuleOptResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRocketMQGroupRequest() (request *DeleteRocketMQGroupRequest) {
	request = &DeleteRocketMQGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQGroup")
	return
}

func NewDeleteRocketMQGroupResponse() (response *DeleteRocketMQGroupResponse) {
	response = &DeleteRocketMQGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RocketMQ消费组
func (c *Client) DeleteRocketMQGroup(request *DeleteRocketMQGroupRequest) (response *DeleteRocketMQGroupResponse, err error) {
	if request == nil {
		request = NewDeleteRocketMQGroupRequest()
	}
	response = NewDeleteRocketMQGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeSubscriptionsRequest() (request *DescribeSubscriptionsRequest) {
	request = &DescribeSubscriptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeSubscriptions")
	return
}

func NewDescribeSubscriptionsResponse() (response *DescribeSubscriptionsResponse) {
	response = &DescribeSubscriptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定环境和主题下的订阅者列表
func (c *Client) DescribeSubscriptions(request *DescribeSubscriptionsRequest) (response *DescribeSubscriptionsResponse, err error) {
	if request == nil {
		request = NewDescribeSubscriptionsRequest()
	}
	response = NewDescribeSubscriptionsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRocketMQVipInstanceStateRequest() (request *ModifyRocketMQVipInstanceStateRequest) {
	request = &ModifyRocketMQVipInstanceStateRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQVipInstanceState")
	return
}

func NewModifyRocketMQVipInstanceStateResponse() (response *ModifyRocketMQVipInstanceStateResponse) {
	response = &ModifyRocketMQVipInstanceStateResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于修改专享实例状态
func (c *Client) ModifyRocketMQVipInstanceState(request *ModifyRocketMQVipInstanceStateRequest) (response *ModifyRocketMQVipInstanceStateResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQVipInstanceStateRequest()
	}
	response = NewModifyRocketMQVipInstanceStateResponse()
	err = c.Send(request, response)
	return
}

func NewModifyPulsarResourceAllocationRuleRequest() (request *ModifyPulsarResourceAllocationRuleRequest) {
	request = &ModifyPulsarResourceAllocationRuleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyPulsarResourceAllocationRule")
	return
}

func NewModifyPulsarResourceAllocationRuleResponse() (response *ModifyPulsarResourceAllocationRuleResponse) {
	response = &ModifyPulsarResourceAllocationRuleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改集群物理资源分配规则
func (c *Client) ModifyPulsarResourceAllocationRule(request *ModifyPulsarResourceAllocationRuleRequest) (response *ModifyPulsarResourceAllocationRuleResponse, err error) {
	if request == nil {
		request = NewModifyPulsarResourceAllocationRuleRequest()
	}
	response = NewModifyPulsarResourceAllocationRuleResponse()
	err = c.Send(request, response)
	return
}

func NewSendCmqMsgRequest() (request *SendCmqMsgRequest) {
	request = &SendCmqMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SendCmqMsg")
	return
}

func NewSendCmqMsgResponse() (response *SendCmqMsgResponse) {
	response = &SendCmqMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送cmq消息
func (c *Client) SendCmqMsg(request *SendCmqMsgRequest) (response *SendCmqMsgResponse, err error) {
	if request == nil {
		request = NewSendCmqMsgRequest()
	}
	response = NewSendCmqMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQTopicsRequest() (request *DescribeRocketMQTopicsRequest) {
	request = &DescribeRocketMQTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopics")
	return
}

func NewDescribeRocketMQTopicsResponse() (response *DescribeRocketMQTopicsResponse) {
	response = &DescribeRocketMQTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取RocketMQ主题列表
func (c *Client) DescribeRocketMQTopics(request *DescribeRocketMQTopicsRequest) (response *DescribeRocketMQTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQTopicsRequest()
	}
	response = NewDescribeRocketMQTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespacesOptRequest() (request *DescribeNamespacesOptRequest) {
	request = &DescribeNamespacesOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeNamespacesOpt")
	return
}

func NewDescribeNamespacesOptResponse() (response *DescribeNamespacesOptResponse) {
	response = &DescribeNamespacesOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取命名空间列表
func (c *Client) DescribeNamespacesOpt(request *DescribeNamespacesOptRequest) (response *DescribeNamespacesOptResponse, err error) {
	if request == nil {
		request = NewDescribeNamespacesOptRequest()
	}
	response = NewDescribeNamespacesOptResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteSubscriptionsRequest() (request *DeleteSubscriptionsRequest) {
	request = &DeleteSubscriptionsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteSubscriptions")
	return
}

func NewDeleteSubscriptionsResponse() (response *DeleteSubscriptionsResponse) {
	response = &DeleteSubscriptionsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除订阅关系
func (c *Client) DeleteSubscriptions(request *DeleteSubscriptionsRequest) (response *DeleteSubscriptionsResponse, err error) {
	if request == nil {
		request = NewDeleteSubscriptionsRequest()
	}
	response = NewDeleteSubscriptionsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCmqQueueDetailRequest() (request *DescribeCmqQueueDetailRequest) {
	request = &DescribeCmqQueueDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqQueueDetail")
	return
}

func NewDescribeCmqQueueDetailResponse() (response *DescribeCmqQueueDetailResponse) {
	response = &DescribeCmqQueueDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cmq队列详情
func (c *Client) DescribeCmqQueueDetail(request *DescribeCmqQueueDetailRequest) (response *DescribeCmqQueueDetailResponse, err error) {
	if request == nil {
		request = NewDescribeCmqQueueDetailRequest()
	}
	response = NewDescribeCmqQueueDetailResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEnvironmentsRequest() (request *DeleteEnvironmentsRequest) {
	request = &DeleteEnvironmentsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteEnvironments")
	return
}

func NewDeleteEnvironmentsResponse() (response *DeleteEnvironmentsResponse) {
	response = &DeleteEnvironmentsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量删除租户下的命名空间
func (c *Client) DeleteEnvironments(request *DeleteEnvironmentsRequest) (response *DeleteEnvironmentsResponse, err error) {
	if request == nil {
		request = NewDeleteEnvironmentsRequest()
	}
	response = NewDeleteEnvironmentsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyEnvironmentAttributesRequest() (request *ModifyEnvironmentAttributesRequest) {
	request = &ModifyEnvironmentAttributesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyEnvironmentAttributes")
	return
}

func NewModifyEnvironmentAttributesResponse() (response *ModifyEnvironmentAttributesResponse) {
	response = &ModifyEnvironmentAttributesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改指定命名空间的属性值
func (c *Client) ModifyEnvironmentAttributes(request *ModifyEnvironmentAttributesRequest) (response *ModifyEnvironmentAttributesResponse, err error) {
	if request == nil {
		request = NewModifyEnvironmentAttributesRequest()
	}
	response = NewModifyEnvironmentAttributesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeCmqTopicDetailRequest() (request *DescribeCmqTopicDetailRequest) {
	request = &DescribeCmqTopicDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqTopicDetail")
	return
}

func NewDescribeCmqTopicDetailResponse() (response *DescribeCmqTopicDetailResponse) {
	response = &DescribeCmqTopicDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询cmq主题详情
func (c *Client) DescribeCmqTopicDetail(request *DescribeCmqTopicDetailRequest) (response *DescribeCmqTopicDetailResponse, err error) {
	if request == nil {
		request = NewDescribeCmqTopicDetailRequest()
	}
	response = NewDescribeCmqTopicDetailResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyRocketMQConsumeRequest() (request *VerifyRocketMQConsumeRequest) {
	request = &VerifyRocketMQConsumeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "VerifyRocketMQConsume")
	return
}

func NewVerifyRocketMQConsumeResponse() (response *VerifyRocketMQConsumeResponse) {
	response = &VerifyRocketMQConsumeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// Rocketmq消费验证
func (c *Client) VerifyRocketMQConsume(request *VerifyRocketMQConsumeRequest) (response *VerifyRocketMQConsumeResponse, err error) {
	if request == nil {
		request = NewVerifyRocketMQConsumeRequest()
	}
	response = NewVerifyRocketMQConsumeResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRocketMQRoleRequest() (request *CreateRocketMQRoleRequest) {
	request = &CreateRocketMQRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQRole")
	return
}

func NewCreateRocketMQRoleResponse() (response *CreateRocketMQRoleResponse) {
	response = &CreateRocketMQRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建角色
func (c *Client) CreateRocketMQRole(request *CreateRocketMQRoleRequest) (response *CreateRocketMQRoleResponse, err error) {
	if request == nil {
		request = NewCreateRocketMQRoleRequest()
	}
	response = NewCreateRocketMQRoleResponse()
	err = c.Send(request, response)
	return
}

func NewResetMsgSubOffsetByTimestampRequest() (request *ResetMsgSubOffsetByTimestampRequest) {
	request = &ResetMsgSubOffsetByTimestampRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ResetMsgSubOffsetByTimestamp")
	return
}

func NewResetMsgSubOffsetByTimestampResponse() (response *ResetMsgSubOffsetByTimestampResponse) {
	response = &ResetMsgSubOffsetByTimestampResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据时间戳进行消息回溯，精确到毫秒
func (c *Client) ResetMsgSubOffsetByTimestamp(request *ResetMsgSubOffsetByTimestampRequest) (response *ResetMsgSubOffsetByTimestampResponse, err error) {
	if request == nil {
		request = NewResetMsgSubOffsetByTimestampRequest()
	}
	response = NewResetMsgSubOffsetByTimestampResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInternalRocketMQInstanceRequest() (request *ModifyInternalRocketMQInstanceRequest) {
	request = &ModifyInternalRocketMQInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyInternalRocketMQInstance")
	return
}

func NewModifyInternalRocketMQInstanceResponse() (response *ModifyInternalRocketMQInstanceResponse) {
	response = &ModifyInternalRocketMQInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此为内部接口，用于运营端修改用户RocketMQ实例的属性
func (c *Client) ModifyInternalRocketMQInstance(request *ModifyInternalRocketMQInstanceRequest) (response *ModifyInternalRocketMQInstanceResponse, err error) {
	if request == nil {
		request = NewModifyInternalRocketMQInstanceRequest()
	}
	response = NewModifyInternalRocketMQInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateClusterRequest() (request *CreateClusterRequest) {
	request = &CreateClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateCluster")
	return
}

func NewCreateClusterResponse() (response *CreateClusterResponse) {
	response = &CreateClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建用户的集群
func (c *Client) CreateCluster(request *CreateClusterRequest) (response *CreateClusterResponse, err error) {
	if request == nil {
		request = NewCreateClusterRequest()
	}
	response = NewCreateClusterResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopicStatsOptRequest() (request *GetTopicStatsOptRequest) {
	request = &GetTopicStatsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "GetTopicStatsOpt")
	return
}

func NewGetTopicStatsOptResponse() (response *GetTopicStatsOptResponse) {
	response = &GetTopicStatsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取Topic状态
func (c *Client) GetTopicStatsOpt(request *GetTopicStatsOptRequest) (response *GetTopicStatsOptResponse, err error) {
	if request == nil {
		request = NewGetTopicStatsOptRequest()
	}
	response = NewGetTopicStatsOptResponse()
	err = c.Send(request, response)
	return
}

func NewAcknowledgeMessageRequest() (request *AcknowledgeMessageRequest) {
	request = &AcknowledgeMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "AcknowledgeMessage")
	return
}

func NewAcknowledgeMessageResponse() (response *AcknowledgeMessageResponse) {
	response = &AcknowledgeMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 根据提供的 MessageID 确认指定 topic 中的消息
func (c *Client) AcknowledgeMessage(request *AcknowledgeMessageRequest) (response *AcknowledgeMessageResponse, err error) {
	if request == nil {
		request = NewAcknowledgeMessageRequest()
	}
	response = NewAcknowledgeMessageResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEnvironmentRoleRequest() (request *CreateEnvironmentRoleRequest) {
	request = &CreateEnvironmentRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateEnvironmentRole")
	return
}

func NewCreateEnvironmentRoleResponse() (response *CreateEnvironmentRoleResponse) {
	response = &CreateEnvironmentRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建环境角色授权
func (c *Client) CreateEnvironmentRole(request *CreateEnvironmentRoleRequest) (response *CreateEnvironmentRoleResponse, err error) {
	if request == nil {
		request = NewCreateEnvironmentRoleRequest()
	}
	response = NewCreateEnvironmentRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterConfigOptRequest() (request *DescribeClusterConfigOptRequest) {
	request = &DescribeClusterConfigOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterConfigOpt")
	return
}

func NewDescribeClusterConfigOptResponse() (response *DescribeClusterConfigOptResponse) {
	response = &DescribeClusterConfigOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新集群配置详细信息
func (c *Client) DescribeClusterConfigOpt(request *DescribeClusterConfigOptRequest) (response *DescribeClusterConfigOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterConfigOptRequest()
	}
	response = NewDescribeClusterConfigOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterEndpointsOptRequest() (request *DescribeClusterEndpointsOptRequest) {
	request = &DescribeClusterEndpointsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterEndpointsOpt")
	return
}

func NewDescribeClusterEndpointsOptResponse() (response *DescribeClusterEndpointsOptResponse) {
	response = &DescribeClusterEndpointsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取物理集群接入点列表
func (c *Client) DescribeClusterEndpointsOpt(request *DescribeClusterEndpointsOptRequest) (response *DescribeClusterEndpointsOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterEndpointsOptRequest()
	}
	response = NewDescribeClusterEndpointsOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterVersionsOptRequest() (request *DescribeClusterVersionsOptRequest) {
	request = &DescribeClusterVersionsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterVersionsOpt")
	return
}

func NewDescribeClusterVersionsOptResponse() (response *DescribeClusterVersionsOptResponse) {
	response = &DescribeClusterVersionsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取物理集群版本
func (c *Client) DescribeClusterVersionsOpt(request *DescribeClusterVersionsOptRequest) (response *DescribeClusterVersionsOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterVersionsOptRequest()
	}
	response = NewDescribeClusterVersionsOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateResourceRuleOptRequest() (request *CreateResourceRuleOptRequest) {
	request = &CreateResourceRuleOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateResourceRuleOpt")
	return
}

func NewCreateResourceRuleOptResponse() (response *CreateResourceRuleOptResponse) {
	response = &CreateResourceRuleOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端创建集群物理资源分配规则
func (c *Client) CreateResourceRuleOpt(request *CreateResourceRuleOptRequest) (response *CreateResourceRuleOptResponse, err error) {
	if request == nil {
		request = NewCreateResourceRuleOptRequest()
	}
	response = NewCreateResourceRuleOptResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRouteRequest() (request *DeleteRouteRequest) {
	request = &DeleteRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRoute")
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

func NewUpdatePolicyAntiAffinityGroupOptRequest() (request *UpdatePolicyAntiAffinityGroupOptRequest) {
	request = &UpdatePolicyAntiAffinityGroupOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyAntiAffinityGroupOpt")
	return
}

func NewUpdatePolicyAntiAffinityGroupOptResponse() (response *UpdatePolicyAntiAffinityGroupOptResponse) {
	response = &UpdatePolicyAntiAffinityGroupOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间反亲和策略
func (c *Client) UpdatePolicyAntiAffinityGroupOpt(request *UpdatePolicyAntiAffinityGroupOptRequest) (response *UpdatePolicyAntiAffinityGroupOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyAntiAffinityGroupOptRequest()
	}
	response = NewUpdatePolicyAntiAffinityGroupOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterComponentMetricsListOptRequest() (request *DescribeClusterComponentMetricsListOptRequest) {
	request = &DescribeClusterComponentMetricsListOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterComponentMetricsListOpt")
	return
}

func NewDescribeClusterComponentMetricsListOptResponse() (response *DescribeClusterComponentMetricsListOptResponse) {
	response = &DescribeClusterComponentMetricsListOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取组件监控指标列表
func (c *Client) DescribeClusterComponentMetricsListOpt(request *DescribeClusterComponentMetricsListOptRequest) (response *DescribeClusterComponentMetricsListOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterComponentMetricsListOptRequest()
	}
	response = NewDescribeClusterComponentMetricsListOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQAppIdStatsRequest() (request *DescribeRocketMQAppIdStatsRequest) {
	request = &DescribeRocketMQAppIdStatsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQAppIdStats")
	return
}

func NewDescribeRocketMQAppIdStatsResponse() (response *DescribeRocketMQAppIdStatsResponse) {
	response = &DescribeRocketMQAppIdStatsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询RocketMQ用户统计信息
func (c *Client) DescribeRocketMQAppIdStats(request *DescribeRocketMQAppIdStatsRequest) (response *DescribeRocketMQAppIdStatsResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQAppIdStatsRequest()
	}
	response = NewDescribeRocketMQAppIdStatsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteBindVpcRequest() (request *DeleteBindVpcRequest) {
	request = &DeleteBindVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteBindVpc")
	return
}

func NewDeleteBindVpcResponse() (response *DeleteBindVpcResponse) {
	response = &DeleteBindVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除租户VPC绑定关系
func (c *Client) DeleteBindVpc(request *DeleteBindVpcRequest) (response *DeleteBindVpcResponse, err error) {
	if request == nil {
		request = NewDeleteBindVpcRequest()
	}
	response = NewDeleteBindVpcResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteEnvironmentRolesRequest() (request *DeleteEnvironmentRolesRequest) {
	request = &DeleteEnvironmentRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteEnvironmentRoles")
	return
}

func NewDeleteEnvironmentRolesResponse() (response *DeleteEnvironmentRolesResponse) {
	response = &DeleteEnvironmentRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除环境角色授权。
func (c *Client) DeleteEnvironmentRoles(request *DeleteEnvironmentRolesRequest) (response *DeleteEnvironmentRolesResponse, err error) {
	if request == nil {
		request = NewDeleteEnvironmentRolesRequest()
	}
	response = NewDeleteEnvironmentRolesResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteAMQPClusterRequest() (request *DeleteAMQPClusterRequest) {
	request = &DeleteAMQPClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteAMQPCluster")
	return
}

func NewDeleteAMQPClusterResponse() (response *DeleteAMQPClusterResponse) {
	response = &DeleteAMQPClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除AMQP集群
func (c *Client) DeleteAMQPCluster(request *DeleteAMQPClusterRequest) (response *DeleteAMQPClusterResponse, err error) {
	if request == nil {
		request = NewDeleteAMQPClusterRequest()
	}
	response = NewDeleteAMQPClusterResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQTopicsByGroupRequest() (request *DescribeRocketMQTopicsByGroupRequest) {
	request = &DescribeRocketMQTopicsByGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopicsByGroup")
	return
}

func NewDescribeRocketMQTopicsByGroupResponse() (response *DescribeRocketMQTopicsByGroupResponse) {
	response = &DescribeRocketMQTopicsByGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定消费组下订阅的主题列表
func (c *Client) DescribeRocketMQTopicsByGroup(request *DescribeRocketMQTopicsByGroupRequest) (response *DescribeRocketMQTopicsByGroupResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQTopicsByGroupRequest()
	}
	response = NewDescribeRocketMQTopicsByGroupResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCmqTopicAttributeRequest() (request *ModifyCmqTopicAttributeRequest) {
	request = &ModifyCmqTopicAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqTopicAttribute")
	return
}

func NewModifyCmqTopicAttributeResponse() (response *ModifyCmqTopicAttributeResponse) {
	response = &ModifyCmqTopicAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改cmq主题属性
func (c *Client) ModifyCmqTopicAttribute(request *ModifyCmqTopicAttributeRequest) (response *ModifyCmqTopicAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCmqTopicAttributeRequest()
	}
	response = NewModifyCmqTopicAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewRewindCmqQueueRequest() (request *RewindCmqQueueRequest) {
	request = &RewindCmqQueueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "RewindCmqQueue")
	return
}

func NewRewindCmqQueueResponse() (response *RewindCmqQueueResponse) {
	response = &RewindCmqQueueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 回溯cmq队列
func (c *Client) RewindCmqQueue(request *RewindCmqQueueRequest) (response *RewindCmqQueueResponse, err error) {
	if request == nil {
		request = NewRewindCmqQueueRequest()
	}
	response = NewRewindCmqQueueResponse()
	err = c.Send(request, response)
	return
}

func NewClearNamespaceBacklogOptRequest() (request *ClearNamespaceBacklogOptRequest) {
	request = &ClearNamespaceBacklogOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ClearNamespaceBacklogOpt")
	return
}

func NewClearNamespaceBacklogOptResponse() (response *ClearNamespaceBacklogOptResponse) {
	response = &ClearNamespaceBacklogOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端清除命名空间堆积消息
func (c *Client) ClearNamespaceBacklogOpt(request *ClearNamespaceBacklogOptRequest) (response *ClearNamespaceBacklogOptResponse, err error) {
	if request == nil {
		request = NewClearNamespaceBacklogOptRequest()
	}
	response = NewClearNamespaceBacklogOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQMigrationTasksRequest() (request *DescribeRocketMQMigrationTasksRequest) {
	request = &DescribeRocketMQMigrationTasksRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQMigrationTasks")
	return
}

func NewDescribeRocketMQMigrationTasksResponse() (response *DescribeRocketMQMigrationTasksResponse) {
	response = &DescribeRocketMQMigrationTasksResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询RocketMQ数据迁移任务列表
func (c *Client) DescribeRocketMQMigrationTasks(request *DescribeRocketMQMigrationTasksRequest) (response *DescribeRocketMQMigrationTasksResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQMigrationTasksRequest()
	}
	response = NewDescribeRocketMQMigrationTasksResponse()
	err = c.Send(request, response)
	return
}

func NewGetTenantTopListOptRequest() (request *GetTenantTopListOptRequest) {
	request = &GetTenantTopListOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "GetTenantTopListOpt")
	return
}

func NewGetTenantTopListOptResponse() (response *GetTenantTopListOptResponse) {
	response = &GetTenantTopListOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取Top租户列表，可以根据指定维度指标获取
func (c *Client) GetTenantTopListOpt(request *GetTenantTopListOptRequest) (response *GetTenantTopListOptResponse, err error) {
	if request == nil {
		request = NewGetTenantTopListOptRequest()
	}
	response = NewGetTenantTopListOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBrokerOwnedNamespacesOptRequest() (request *DescribeBrokerOwnedNamespacesOptRequest) {
	request = &DescribeBrokerOwnedNamespacesOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBrokerOwnedNamespacesOpt")
	return
}

func NewDescribeBrokerOwnedNamespacesOptResponse() (response *DescribeBrokerOwnedNamespacesOptResponse) {
	response = &DescribeBrokerOwnedNamespacesOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取broker上的命名空间信息（不分页）
func (c *Client) DescribeBrokerOwnedNamespacesOpt(request *DescribeBrokerOwnedNamespacesOptRequest) (response *DescribeBrokerOwnedNamespacesOptResponse, err error) {
	if request == nil {
		request = NewDescribeBrokerOwnedNamespacesOptRequest()
	}
	response = NewDescribeBrokerOwnedNamespacesOptResponse()
	err = c.Send(request, response)
	return
}

func NewRetryRocketMQDlqMessageRequest() (request *RetryRocketMQDlqMessageRequest) {
	request = &RetryRocketMQDlqMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "RetryRocketMQDlqMessage")
	return
}

func NewRetryRocketMQDlqMessageResponse() (response *RetryRocketMQDlqMessageResponse) {
	response = &RetryRocketMQDlqMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 重发RocketMQ死信消息
func (c *Client) RetryRocketMQDlqMessage(request *RetryRocketMQDlqMessageRequest) (response *RetryRocketMQDlqMessageResponse, err error) {
	if request == nil {
		request = NewRetryRocketMQDlqMessageRequest()
	}
	response = NewRetryRocketMQDlqMessageResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRocketMQNamespaceRequest() (request *CreateRocketMQNamespaceRequest) {
	request = &CreateRocketMQNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQNamespace")
	return
}

func NewCreateRocketMQNamespaceResponse() (response *CreateRocketMQNamespaceResponse) {
	response = &CreateRocketMQNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建RocketMQ命名空间
func (c *Client) CreateRocketMQNamespace(request *CreateRocketMQNamespaceRequest) (response *CreateRocketMQNamespaceResponse, err error) {
	if request == nil {
		request = NewCreateRocketMQNamespaceRequest()
	}
	response = NewCreateRocketMQNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAMQPExchangeRequest() (request *ModifyAMQPExchangeRequest) {
	request = &ModifyAMQPExchangeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyAMQPExchange")
	return
}

func NewModifyAMQPExchangeResponse() (response *ModifyAMQPExchangeResponse) {
	response = &ModifyAMQPExchangeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Amqp交换机
func (c *Client) ModifyAMQPExchange(request *ModifyAMQPExchangeRequest) (response *ModifyAMQPExchangeResponse, err error) {
	if request == nil {
		request = NewModifyAMQPExchangeRequest()
	}
	response = NewModifyAMQPExchangeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNodeHealthOptRequest() (request *DescribeNodeHealthOptRequest) {
	request = &DescribeNodeHealthOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeNodeHealthOpt")
	return
}

func NewDescribeNodeHealthOptResponse() (response *DescribeNodeHealthOptResponse) {
	response = &DescribeNodeHealthOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获节点健康状态
func (c *Client) DescribeNodeHealthOpt(request *DescribeNodeHealthOptRequest) (response *DescribeNodeHealthOptResponse, err error) {
	if request == nil {
		request = NewDescribeNodeHealthOptRequest()
	}
	response = NewDescribeNodeHealthOptResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterMetaInitScriptOptRequest() (request *GetClusterMetaInitScriptOptRequest) {
	request = &GetClusterMetaInitScriptOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "GetClusterMetaInitScriptOpt")
	return
}

func NewGetClusterMetaInitScriptOptResponse() (response *GetClusterMetaInitScriptOptResponse) {
	response = &GetClusterMetaInitScriptOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取初始化脚本
func (c *Client) GetClusterMetaInitScriptOpt(request *GetClusterMetaInitScriptOptRequest) (response *GetClusterMetaInitScriptOptResponse, err error) {
	if request == nil {
		request = NewGetClusterMetaInitScriptOptRequest()
	}
	response = NewGetClusterMetaInitScriptOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateEnvironmentRequest() (request *CreateEnvironmentRequest) {
	request = &CreateEnvironmentRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateEnvironment")
	return
}

func NewCreateEnvironmentResponse() (response *CreateEnvironmentResponse) {
	response = &CreateEnvironmentResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于在用户账户下创建消息队列 Tdmq 命名空间
func (c *Client) CreateEnvironment(request *CreateEnvironmentRequest) (response *CreateEnvironmentResponse, err error) {
	if request == nil {
		request = NewCreateEnvironmentRequest()
	}
	response = NewCreateEnvironmentResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRouteRequest() (request *CreateRouteRequest) {
	request = &CreateRouteRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRoute")
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

func NewDeleteAMQPVHostRequest() (request *DeleteAMQPVHostRequest) {
	request = &DeleteAMQPVHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteAMQPVHost")
	return
}

func NewDeleteAMQPVHostResponse() (response *DeleteAMQPVHostResponse) {
	response = &DeleteAMQPVHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除Vhost
func (c *Client) DeleteAMQPVHost(request *DeleteAMQPVHostRequest) (response *DeleteAMQPVHostResponse, err error) {
	if request == nil {
		request = NewDeleteAMQPVHostRequest()
	}
	response = NewDeleteAMQPVHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClustersRequest() (request *DescribeClustersRequest) {
	request = &DescribeClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusters")
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

func NewModifyRocketMQInstanceRequest() (request *ModifyRocketMQInstanceRequest) {
	request = &ModifyRocketMQInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQInstance")
	return
}

func NewModifyRocketMQInstanceResponse() (response *ModifyRocketMQInstanceResponse) {
	response = &ModifyRocketMQInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改RocketMQ专享实例
func (c *Client) ModifyRocketMQInstance(request *ModifyRocketMQInstanceRequest) (response *ModifyRocketMQInstanceResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQInstanceRequest()
	}
	response = NewModifyRocketMQInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribePulsarResourceAllocationRulesRequest() (request *DescribePulsarResourceAllocationRulesRequest) {
	request = &DescribePulsarResourceAllocationRulesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribePulsarResourceAllocationRules")
	return
}

func NewDescribePulsarResourceAllocationRulesResponse() (response *DescribePulsarResourceAllocationRulesResponse) {
	response = &DescribePulsarResourceAllocationRulesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取集群资源分配规则列表
func (c *Client) DescribePulsarResourceAllocationRules(request *DescribePulsarResourceAllocationRulesRequest) (response *DescribePulsarResourceAllocationRulesResponse, err error) {
	if request == nil {
		request = NewDescribePulsarResourceAllocationRulesRequest()
	}
	response = NewDescribePulsarResourceAllocationRulesResponse()
	err = c.Send(request, response)
	return
}

func NewPutClusterOfflineOptRequest() (request *PutClusterOfflineOptRequest) {
	request = &PutClusterOfflineOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "PutClusterOfflineOpt")
	return
}

func NewPutClusterOfflineOptResponse() (response *PutClusterOfflineOptResponse) {
	response = &PutClusterOfflineOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端下线物理集群
func (c *Client) PutClusterOfflineOpt(request *PutClusterOfflineOptRequest) (response *PutClusterOfflineOptResponse, err error) {
	if request == nil {
		request = NewPutClusterOfflineOptRequest()
	}
	response = NewPutClusterOfflineOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTenantsOptRequest() (request *DescribeTenantsOptRequest) {
	request = &DescribeTenantsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeTenantsOpt")
	return
}

func NewDescribeTenantsOptResponse() (response *DescribeTenantsOptResponse) {
	response = &DescribeTenantsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取物虚拟集群（租户）列表
func (c *Client) DescribeTenantsOpt(request *DescribeTenantsOptRequest) (response *DescribeTenantsOptResponse, err error) {
	if request == nil {
		request = NewDescribeTenantsOptRequest()
	}
	response = NewDescribeTenantsOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeTopicsRequest() (request *DescribeTopicsRequest) {
	request = &DescribeTopicsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeTopics")
	return
}

func NewDescribeTopicsResponse() (response *DescribeTopicsResponse) {
	response = &DescribeTopicsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取环境下主题列表
func (c *Client) DescribeTopics(request *DescribeTopicsRequest) (response *DescribeTopicsResponse, err error) {
	if request == nil {
		request = NewDescribeTopicsRequest()
	}
	response = NewDescribeTopicsResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRocketMQRolesRequest() (request *DeleteRocketMQRolesRequest) {
	request = &DeleteRocketMQRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQRoles")
	return
}

func NewDeleteRocketMQRolesResponse() (response *DeleteRocketMQRolesResponse) {
	response = &DeleteRocketMQRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除角色，支持批量。
func (c *Client) DeleteRocketMQRoles(request *DeleteRocketMQRolesRequest) (response *DeleteRocketMQRolesResponse, err error) {
	if request == nil {
		request = NewDeleteRocketMQRolesRequest()
	}
	response = NewDeleteRocketMQRolesResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRocketMQInstanceSpecRequest() (request *ModifyRocketMQInstanceSpecRequest) {
	request = &ModifyRocketMQInstanceSpecRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQInstanceSpec")
	return
}

func NewModifyRocketMQInstanceSpecResponse() (response *ModifyRocketMQInstanceSpecResponse) {
	response = &ModifyRocketMQInstanceSpecResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 本API用于修改RocketMQ专享实例配置，可以支持实例规格、节点数和存储的升配和实例规格的降配。本API发起订单并成功支付后进入实例配置变更的流程，可通过DescribeRocketMQVipInstances查询实例是否已变更完成。
func (c *Client) ModifyRocketMQInstanceSpec(request *ModifyRocketMQInstanceSpecRequest) (response *ModifyRocketMQInstanceSpecResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQInstanceSpecRequest()
	}
	response = NewModifyRocketMQInstanceSpecResponse()
	err = c.Send(request, response)
	return
}

func NewGetClusterBaradMetricsOptRequest() (request *GetClusterBaradMetricsOptRequest) {
	request = &GetClusterBaradMetricsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "GetClusterBaradMetricsOpt")
	return
}

func NewGetClusterBaradMetricsOptResponse() (response *GetClusterBaradMetricsOptResponse) {
	response = &GetClusterBaradMetricsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取物理集群barad监控指标（获取指标最新值和获取时间段监控指标通用）
func (c *Client) GetClusterBaradMetricsOpt(request *GetClusterBaradMetricsOptRequest) (response *GetClusterBaradMetricsOptResponse, err error) {
	if request == nil {
		request = NewGetClusterBaradMetricsOptRequest()
	}
	response = NewGetClusterBaradMetricsOptResponse()
	err = c.Send(request, response)
	return
}

func NewSendAllModulesConfigOptRequest() (request *SendAllModulesConfigOptRequest) {
	request = &SendAllModulesConfigOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SendAllModulesConfigOpt")
	return
}

func NewSendAllModulesConfigOptResponse() (response *SendAllModulesConfigOptResponse) {
	response = &SendAllModulesConfigOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 下发所有模块配置
func (c *Client) SendAllModulesConfigOpt(request *SendAllModulesConfigOptRequest) (response *SendAllModulesConfigOptResponse, err error) {
	if request == nil {
		request = NewSendAllModulesConfigOptRequest()
	}
	response = NewSendAllModulesConfigOptResponse()
	err = c.Send(request, response)
	return
}

func NewGetTopicInternalStatsOptRequest() (request *GetTopicInternalStatsOptRequest) {
	request = &GetTopicInternalStatsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "GetTopicInternalStatsOpt")
	return
}

func NewGetTopicInternalStatsOptResponse() (response *GetTopicInternalStatsOptResponse) {
	response = &GetTopicInternalStatsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取topic内部状态
func (c *Client) GetTopicInternalStatsOpt(request *GetTopicInternalStatsOptRequest) (response *GetTopicInternalStatsOptResponse, err error) {
	if request == nil {
		request = NewGetTopicInternalStatsOptRequest()
	}
	response = NewGetTopicInternalStatsOptResponse()
	err = c.Send(request, response)
	return
}

func NewModifyAMQPQueueRequest() (request *ModifyAMQPQueueRequest) {
	request = &ModifyAMQPQueueRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyAMQPQueue")
	return
}

func NewModifyAMQPQueueResponse() (response *ModifyAMQPQueueResponse) {
	response = &ModifyAMQPQueueResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新Amqp队列
func (c *Client) ModifyAMQPQueue(request *ModifyAMQPQueueRequest) (response *ModifyAMQPQueueResponse, err error) {
	if request == nil {
		request = NewModifyAMQPQueueRequest()
	}
	response = NewModifyAMQPQueueResponse()
	err = c.Send(request, response)
	return
}

func NewCreateBindVpcRequest() (request *CreateBindVpcRequest) {
	request = &CreateBindVpcRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateBindVpc")
	return
}

func NewCreateBindVpcResponse() (response *CreateBindVpcResponse) {
	response = &CreateBindVpcResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建租户VPC绑定关系
func (c *Client) CreateBindVpc(request *CreateBindVpcRequest) (response *CreateBindVpcResponse, err error) {
	if request == nil {
		request = NewCreateBindVpcRequest()
	}
	response = NewCreateBindVpcResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyPersistenceOptRequest() (request *UpdatePolicyPersistenceOptRequest) {
	request = &UpdatePolicyPersistenceOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyPersistenceOpt")
	return
}

func NewUpdatePolicyPersistenceOptResponse() (response *UpdatePolicyPersistenceOptResponse) {
	response = &UpdatePolicyPersistenceOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间消息持久化策略
func (c *Client) UpdatePolicyPersistenceOpt(request *UpdatePolicyPersistenceOptRequest) (response *UpdatePolicyPersistenceOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyPersistenceOptRequest()
	}
	response = NewUpdatePolicyPersistenceOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQMigrationTaskDetailRequest() (request *DescribeRocketMQMigrationTaskDetailRequest) {
	request = &DescribeRocketMQMigrationTaskDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQMigrationTaskDetail")
	return
}

func NewDescribeRocketMQMigrationTaskDetailResponse() (response *DescribeRocketMQMigrationTaskDetailResponse) {
	response = &DescribeRocketMQMigrationTaskDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于获取RocketMQ数据迁移任务详情
func (c *Client) DescribeRocketMQMigrationTaskDetail(request *DescribeRocketMQMigrationTaskDetailRequest) (response *DescribeRocketMQMigrationTaskDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQMigrationTaskDetailRequest()
	}
	response = NewDescribeRocketMQMigrationTaskDetailResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRocketMQNamespaceRequest() (request *ModifyRocketMQNamespaceRequest) {
	request = &ModifyRocketMQNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRocketMQNamespace")
	return
}

func NewModifyRocketMQNamespaceResponse() (response *ModifyRocketMQNamespaceResponse) {
	response = &ModifyRocketMQNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新RocketMQ命名空间
func (c *Client) ModifyRocketMQNamespace(request *ModifyRocketMQNamespaceRequest) (response *ModifyRocketMQNamespaceResponse, err error) {
	if request == nil {
		request = NewModifyRocketMQNamespaceRequest()
	}
	response = NewModifyRocketMQNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewTpoModifyResourceProjectRequest() (request *TpoModifyResourceProjectRequest) {
	request = &TpoModifyResourceProjectRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "TpoModifyResourceProject")
	return
}

func NewTpoModifyResourceProjectResponse() (response *TpoModifyResourceProjectResponse) {
	response = &TpoModifyResourceProjectResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 新增或修改资源对应的项目，新增时 oldProjectId 为空
func (c *Client) TpoModifyResourceProject(request *TpoModifyResourceProjectRequest) (response *TpoModifyResourceProjectResponse, err error) {
	if request == nil {
		request = NewTpoModifyResourceProjectRequest()
	}
	response = NewTpoModifyResourceProjectResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAMQPVHostRequest() (request *CreateAMQPVHostRequest) {
	request = &CreateAMQPVHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateAMQPVHost")
	return
}

func NewCreateAMQPVHostResponse() (response *CreateAMQPVHostResponse) {
	response = &CreateAMQPVHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建Amqp Vhost
func (c *Client) CreateAMQPVHost(request *CreateAMQPVHostRequest) (response *CreateAMQPVHostResponse, err error) {
	if request == nil {
		request = NewCreateAMQPVHostRequest()
	}
	response = NewCreateAMQPVHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMsgTraceRequest() (request *DescribeMsgTraceRequest) {
	request = &DescribeMsgTraceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeMsgTrace")
	return
}

func NewDescribeMsgTraceResponse() (response *DescribeMsgTraceResponse) {
	response = &DescribeMsgTraceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询消息轨迹
func (c *Client) DescribeMsgTrace(request *DescribeMsgTraceRequest) (response *DescribeMsgTraceResponse, err error) {
	if request == nil {
		request = NewDescribeMsgTraceRequest()
	}
	response = NewDescribeMsgTraceResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespacePolicyOptRequest() (request *DescribeNamespacePolicyOptRequest) {
	request = &DescribeNamespacePolicyOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeNamespacePolicyOpt")
	return
}

func NewDescribeNamespacePolicyOptResponse() (response *DescribeNamespacePolicyOptResponse) {
	response = &DescribeNamespacePolicyOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取命名空间策略
func (c *Client) DescribeNamespacePolicyOpt(request *DescribeNamespacePolicyOptRequest) (response *DescribeNamespacePolicyOptResponse, err error) {
	if request == nil {
		request = NewDescribeNamespacePolicyOptRequest()
	}
	response = NewDescribeNamespacePolicyOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRocketMQTopicRequest() (request *CreateRocketMQTopicRequest) {
	request = &CreateRocketMQTopicRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQTopic")
	return
}

func NewCreateRocketMQTopicResponse() (response *CreateRocketMQTopicResponse) {
	response = &CreateRocketMQTopicResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建RocketMQ主题
func (c *Client) CreateRocketMQTopic(request *CreateRocketMQTopicRequest) (response *CreateRocketMQTopicResponse, err error) {
	if request == nil {
		request = NewCreateRocketMQTopicRequest()
	}
	response = NewCreateRocketMQTopicResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPExchangesRequest() (request *DescribeAMQPExchangesRequest) {
	request = &DescribeAMQPExchangesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPExchanges")
	return
}

func NewDescribeAMQPExchangesResponse() (response *DescribeAMQPExchangesResponse) {
	response = &DescribeAMQPExchangesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取AMQP Exchange列表
func (c *Client) DescribeAMQPExchanges(request *DescribeAMQPExchangesRequest) (response *DescribeAMQPExchangesResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPExchangesRequest()
	}
	response = NewDescribeAMQPExchangesResponse()
	err = c.Send(request, response)
	return
}

func NewPutClusterOnlineOptRequest() (request *PutClusterOnlineOptRequest) {
	request = &PutClusterOnlineOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "PutClusterOnlineOpt")
	return
}

func NewPutClusterOnlineOptResponse() (response *PutClusterOnlineOptResponse) {
	response = &PutClusterOnlineOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端上线物理集群
func (c *Client) PutClusterOnlineOpt(request *PutClusterOnlineOptRequest) (response *PutClusterOnlineOptResponse, err error) {
	if request == nil {
		request = NewPutClusterOnlineOptRequest()
	}
	response = NewPutClusterOnlineOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQTopUsagesRequest() (request *DescribeRocketMQTopUsagesRequest) {
	request = &DescribeRocketMQTopUsagesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopUsages")
	return
}

func NewDescribeRocketMQTopUsagesResponse() (response *DescribeRocketMQTopUsagesResponse) {
	response = &DescribeRocketMQTopUsagesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于获取RocketMQ指标排序列表，比如集群实例下占用存储空间最多的主题排序。
func (c *Client) DescribeRocketMQTopUsages(request *DescribeRocketMQTopUsagesRequest) (response *DescribeRocketMQTopUsagesResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQTopUsagesRequest()
	}
	response = NewDescribeRocketMQTopUsagesResponse()
	err = c.Send(request, response)
	return
}

func NewVerifyBatchRoleRequest() (request *VerifyBatchRoleRequest) {
	request = &VerifyBatchRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "VerifyBatchRole")
	return
}

func NewVerifyBatchRoleResponse() (response *VerifyBatchRoleResponse) {
	response = &VerifyBatchRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 批量导入角色预校验
func (c *Client) VerifyBatchRole(request *VerifyBatchRoleRequest) (response *VerifyBatchRoleResponse, err error) {
	if request == nil {
		request = NewVerifyBatchRoleRequest()
	}
	response = NewVerifyBatchRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQCreateQuotaRequest() (request *DescribeRocketMQCreateQuotaRequest) {
	request = &DescribeRocketMQCreateQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQCreateQuota")
	return
}

func NewDescribeRocketMQCreateQuotaResponse() (response *DescribeRocketMQCreateQuotaResponse) {
	response = &DescribeRocketMQCreateQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的配额，如Topic容量，Group容量，剩余可创建集群数
func (c *Client) DescribeRocketMQCreateQuota(request *DescribeRocketMQCreateQuotaRequest) (response *DescribeRocketMQCreateQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQCreateQuotaRequest()
	}
	response = NewDescribeRocketMQCreateQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRocketMQGroupRequest() (request *CreateRocketMQGroupRequest) {
	request = &CreateRocketMQGroupRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQGroup")
	return
}

func NewCreateRocketMQGroupResponse() (response *CreateRocketMQGroupResponse) {
	response = &CreateRocketMQGroupResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建RocketMQ消费组
func (c *Client) CreateRocketMQGroup(request *CreateRocketMQGroupRequest) (response *CreateRocketMQGroupResponse, err error) {
	if request == nil {
		request = NewCreateRocketMQGroupRequest()
	}
	response = NewCreateRocketMQGroupResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeBatchRoleRequest() (request *DescribeBatchRoleRequest) {
	request = &DescribeBatchRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeBatchRole")
	return
}

func NewDescribeBatchRoleResponse() (response *DescribeBatchRoleResponse) {
	response = &DescribeBatchRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// DescribeBatchRole
func (c *Client) DescribeBatchRole(request *DescribeBatchRoleRequest) (response *DescribeBatchRoleResponse, err error) {
	if request == nil {
		request = NewDescribeBatchRoleRequest()
	}
	response = NewDescribeBatchRoleResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRocketMQEnvironmentRoleRequest() (request *CreateRocketMQEnvironmentRoleRequest) {
	request = &CreateRocketMQEnvironmentRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQEnvironmentRole")
	return
}

func NewCreateRocketMQEnvironmentRoleResponse() (response *CreateRocketMQEnvironmentRoleResponse) {
	response = &CreateRocketMQEnvironmentRoleResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建环境角色授权
func (c *Client) CreateRocketMQEnvironmentRole(request *CreateRocketMQEnvironmentRoleRequest) (response *CreateRocketMQEnvironmentRoleResponse, err error) {
	if request == nil {
		request = NewCreateRocketMQEnvironmentRoleRequest()
	}
	response = NewCreateRocketMQEnvironmentRoleResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPQueuesRequest() (request *DescribeAMQPQueuesRequest) {
	request = &DescribeAMQPQueuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPQueues")
	return
}

func NewDescribeAMQPQueuesResponse() (response *DescribeAMQPQueuesResponse) {
	response = &DescribeAMQPQueuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取Amqp队列列表
func (c *Client) DescribeAMQPQueues(request *DescribeAMQPQueuesRequest) (response *DescribeAMQPQueuesResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPQueuesRequest()
	}
	response = NewDescribeAMQPQueuesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProducersRequest() (request *DescribeProducersRequest) {
	request = &DescribeProducersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeProducers")
	return
}

func NewDescribeProducersResponse() (response *DescribeProducersResponse) {
	response = &DescribeProducersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取生产者列表，仅显示在线的生产者
func (c *Client) DescribeProducers(request *DescribeProducersRequest) (response *DescribeProducersResponse, err error) {
	if request == nil {
		request = NewDescribeProducersRequest()
	}
	response = NewDescribeProducersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterLatestMetricsOptRequest() (request *DescribeClusterLatestMetricsOptRequest) {
	request = &DescribeClusterLatestMetricsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterLatestMetricsOpt")
	return
}

func NewDescribeClusterLatestMetricsOptResponse() (response *DescribeClusterLatestMetricsOptResponse) {
	response = &DescribeClusterLatestMetricsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取物理集群监控指标最新值
func (c *Client) DescribeClusterLatestMetricsOpt(request *DescribeClusterLatestMetricsOptRequest) (response *DescribeClusterLatestMetricsOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterLatestMetricsOptRequest()
	}
	response = NewDescribeClusterLatestMetricsOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeProtocolsOptRequest() (request *DescribeProtocolsOptRequest) {
	request = &DescribeProtocolsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeProtocolsOpt")
	return
}

func NewDescribeProtocolsOptResponse() (response *DescribeProtocolsOptResponse) {
	response = &DescribeProtocolsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取协议列表（如pulsar、cmq、rocket）
func (c *Client) DescribeProtocolsOpt(request *DescribeProtocolsOptRequest) (response *DescribeProtocolsOptResponse, err error) {
	if request == nil {
		request = NewDescribeProtocolsOptRequest()
	}
	response = NewDescribeProtocolsOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQConsumerConnectionDetailRequest() (request *DescribeRocketMQConsumerConnectionDetailRequest) {
	request = &DescribeRocketMQConsumerConnectionDetailRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQConsumerConnectionDetail")
	return
}

func NewDescribeRocketMQConsumerConnectionDetailResponse() (response *DescribeRocketMQConsumerConnectionDetailResponse) {
	response = &DescribeRocketMQConsumerConnectionDetailResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取在线消费端详情
func (c *Client) DescribeRocketMQConsumerConnectionDetail(request *DescribeRocketMQConsumerConnectionDetailRequest) (response *DescribeRocketMQConsumerConnectionDetailResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQConsumerConnectionDetailRequest()
	}
	response = NewDescribeRocketMQConsumerConnectionDetailResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRoleRequest() (request *CreateRoleRequest) {
	request = &CreateRoleRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRole")
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

func NewDescribeCmqDeadLetterSourceQueuesRequest() (request *DescribeCmqDeadLetterSourceQueuesRequest) {
	request = &DescribeCmqDeadLetterSourceQueuesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeCmqDeadLetterSourceQueues")
	return
}

func NewDescribeCmqDeadLetterSourceQueuesResponse() (response *DescribeCmqDeadLetterSourceQueuesResponse) {
	response = &DescribeCmqDeadLetterSourceQueuesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 枚举cmq死信队列源队列
func (c *Client) DescribeCmqDeadLetterSourceQueues(request *DescribeCmqDeadLetterSourceQueuesRequest) (response *DescribeCmqDeadLetterSourceQueuesResponse, err error) {
	if request == nil {
		request = NewDescribeCmqDeadLetterSourceQueuesRequest()
	}
	response = NewDescribeCmqDeadLetterSourceQueuesResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyDispatchRatePerTopicOptRequest() (request *UpdatePolicyDispatchRatePerTopicOptRequest) {
	request = &UpdatePolicyDispatchRatePerTopicOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyDispatchRatePerTopicOpt")
	return
}

func NewUpdatePolicyDispatchRatePerTopicOptResponse() (response *UpdatePolicyDispatchRatePerTopicOptResponse) {
	response = &UpdatePolicyDispatchRatePerTopicOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间topic级别请求分发策略限制
func (c *Client) UpdatePolicyDispatchRatePerTopicOpt(request *UpdatePolicyDispatchRatePerTopicOptRequest) (response *UpdatePolicyDispatchRatePerTopicOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyDispatchRatePerTopicOptRequest()
	}
	response = NewUpdatePolicyDispatchRatePerTopicOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPQueueConsumersRequest() (request *DescribeAMQPQueueConsumersRequest) {
	request = &DescribeAMQPQueueConsumersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPQueueConsumers")
	return
}

func NewDescribeAMQPQueueConsumersResponse() (response *DescribeAMQPQueueConsumersResponse) {
	response = &DescribeAMQPQueueConsumersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取指定队列下消费者列表
func (c *Client) DescribeAMQPQueueConsumers(request *DescribeAMQPQueueConsumersRequest) (response *DescribeAMQPQueueConsumersResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPQueueConsumersRequest()
	}
	response = NewDescribeAMQPQueueConsumersResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQInstanceDeliveryProgressRequest() (request *DescribeRocketMQInstanceDeliveryProgressRequest) {
	request = &DescribeRocketMQInstanceDeliveryProgressRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQInstanceDeliveryProgress")
	return
}

func NewDescribeRocketMQInstanceDeliveryProgressResponse() (response *DescribeRocketMQInstanceDeliveryProgressResponse) {
	response = &DescribeRocketMQInstanceDeliveryProgressResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 用于查询RocketMQ实例发货进度
func (c *Client) DescribeRocketMQInstanceDeliveryProgress(request *DescribeRocketMQInstanceDeliveryProgressRequest) (response *DescribeRocketMQInstanceDeliveryProgressResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQInstanceDeliveryProgressRequest()
	}
	response = NewDescribeRocketMQInstanceDeliveryProgressResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteTemplateItemsOptRequest() (request *DeleteTemplateItemsOptRequest) {
	request = &DeleteTemplateItemsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteTemplateItemsOpt")
	return
}

func NewDeleteTemplateItemsOptResponse() (response *DeleteTemplateItemsOptResponse) {
	response = &DeleteTemplateItemsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端删除模板配置项
func (c *Client) DeleteTemplateItemsOpt(request *DeleteTemplateItemsOptRequest) (response *DeleteTemplateItemsOptResponse, err error) {
	if request == nil {
		request = NewDeleteTemplateItemsOptRequest()
	}
	response = NewDeleteTemplateItemsOptResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyMaxProducerPerTopicOptRequest() (request *UpdatePolicyMaxProducerPerTopicOptRequest) {
	request = &UpdatePolicyMaxProducerPerTopicOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyMaxProducerPerTopicOpt")
	return
}

func NewUpdatePolicyMaxProducerPerTopicOptResponse() (response *UpdatePolicyMaxProducerPerTopicOptResponse) {
	response = &UpdatePolicyMaxProducerPerTopicOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间topic级别生产者数量限制
func (c *Client) UpdatePolicyMaxProducerPerTopicOpt(request *UpdatePolicyMaxProducerPerTopicOptRequest) (response *UpdatePolicyMaxProducerPerTopicOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyMaxProducerPerTopicOptRequest()
	}
	response = NewUpdatePolicyMaxProducerPerTopicOptResponse()
	err = c.Send(request, response)
	return
}

func NewCreateSubscriptionRequest() (request *CreateSubscriptionRequest) {
	request = &CreateSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateSubscription")
	return
}

func NewCreateSubscriptionResponse() (response *CreateSubscriptionResponse) {
	response = &CreateSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建一个主题的订阅关系
func (c *Client) CreateSubscription(request *CreateSubscriptionRequest) (response *CreateSubscriptionResponse, err error) {
	if request == nil {
		request = NewCreateSubscriptionRequest()
	}
	response = NewCreateSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRocketMQEnvironmentRolesRequest() (request *DeleteRocketMQEnvironmentRolesRequest) {
	request = &DeleteRocketMQEnvironmentRolesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRocketMQEnvironmentRoles")
	return
}

func NewDeleteRocketMQEnvironmentRolesResponse() (response *DeleteRocketMQEnvironmentRolesResponse) {
	response = &DeleteRocketMQEnvironmentRolesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除环境角色授权。
func (c *Client) DeleteRocketMQEnvironmentRoles(request *DeleteRocketMQEnvironmentRolesRequest) (response *DeleteRocketMQEnvironmentRolesResponse, err error) {
	if request == nil {
		request = NewDeleteRocketMQEnvironmentRolesRequest()
	}
	response = NewDeleteRocketMQEnvironmentRolesResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInternalTenantsRequest() (request *DescribeInternalTenantsRequest) {
	request = &DescribeInternalTenantsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeInternalTenants")
	return
}

func NewDescribeInternalTenantsResponse() (response *DescribeInternalTenantsResponse) {
	response = &DescribeInternalTenantsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取虚拟集群列表
func (c *Client) DescribeInternalTenants(request *DescribeInternalTenantsRequest) (response *DescribeInternalTenantsResponse, err error) {
	if request == nil {
		request = NewDescribeInternalTenantsRequest()
	}
	response = NewDescribeInternalTenantsResponse()
	err = c.Send(request, response)
	return
}

func NewModifyClusterRequest() (request *ModifyClusterRequest) {
	request = &ModifyClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCluster")
	return
}

func NewModifyClusterResponse() (response *ModifyClusterResponse) {
	response = &ModifyClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 更新集群信息
func (c *Client) ModifyCluster(request *ModifyClusterRequest) (response *ModifyClusterResponse, err error) {
	if request == nil {
		request = NewModifyClusterRequest()
	}
	response = NewModifyClusterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateAMQPClusterRequest() (request *CreateAMQPClusterRequest) {
	request = &CreateAMQPClusterRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateAMQPCluster")
	return
}

func NewCreateAMQPClusterResponse() (response *CreateAMQPClusterResponse) {
	response = &CreateAMQPClusterResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建AMQP集群
func (c *Client) CreateAMQPCluster(request *CreateAMQPClusterRequest) (response *CreateAMQPClusterResponse, err error) {
	if request == nil {
		request = NewCreateAMQPClusterRequest()
	}
	response = NewCreateAMQPClusterResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRocketMQMigrationTaskRequest() (request *CreateRocketMQMigrationTaskRequest) {
	request = &CreateRocketMQMigrationTaskRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRocketMQMigrationTask")
	return
}

func NewCreateRocketMQMigrationTaskResponse() (response *CreateRocketMQMigrationTaskResponse) {
	response = &CreateRocketMQMigrationTaskResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建RocketMQ元数据迁移任务，用于批量创建主题和消费组数据
func (c *Client) CreateRocketMQMigrationTask(request *CreateRocketMQMigrationTaskRequest) (response *CreateRocketMQMigrationTaskResponse, err error) {
	if request == nil {
		request = NewCreateRocketMQMigrationTaskRequest()
	}
	response = NewCreateRocketMQMigrationTaskResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyMaxConsumerPerTopicOptRequest() (request *UpdatePolicyMaxConsumerPerTopicOptRequest) {
	request = &UpdatePolicyMaxConsumerPerTopicOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyMaxConsumerPerTopicOpt")
	return
}

func NewUpdatePolicyMaxConsumerPerTopicOptResponse() (response *UpdatePolicyMaxConsumerPerTopicOptResponse) {
	response = &UpdatePolicyMaxConsumerPerTopicOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间topic级别消费者数量限制
func (c *Client) UpdatePolicyMaxConsumerPerTopicOpt(request *UpdatePolicyMaxConsumerPerTopicOptRequest) (response *UpdatePolicyMaxConsumerPerTopicOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyMaxConsumerPerTopicOptRequest()
	}
	response = NewUpdatePolicyMaxConsumerPerTopicOptResponse()
	err = c.Send(request, response)
	return
}

func NewUpdatePolicyBacklogQuotaOptRequest() (request *UpdatePolicyBacklogQuotaOptRequest) {
	request = &UpdatePolicyBacklogQuotaOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "UpdatePolicyBacklogQuotaOpt")
	return
}

func NewUpdatePolicyBacklogQuotaOptResponse() (response *UpdatePolicyBacklogQuotaOptResponse) {
	response = &UpdatePolicyBacklogQuotaOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端更新命名空间消息堆积策略
func (c *Client) UpdatePolicyBacklogQuotaOpt(request *UpdatePolicyBacklogQuotaOptRequest) (response *UpdatePolicyBacklogQuotaOptResponse, err error) {
	if request == nil {
		request = NewUpdatePolicyBacklogQuotaOptRequest()
	}
	response = NewUpdatePolicyBacklogQuotaOptResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTemplateItemsOptRequest() (request *ModifyTemplateItemsOptRequest) {
	request = &ModifyTemplateItemsOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyTemplateItemsOpt")
	return
}

func NewModifyTemplateItemsOptResponse() (response *ModifyTemplateItemsOptResponse) {
	response = &ModifyTemplateItemsOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改模板配置项
func (c *Client) ModifyTemplateItemsOpt(request *ModifyTemplateItemsOptRequest) (response *ModifyTemplateItemsOptResponse, err error) {
	if request == nil {
		request = NewModifyTemplateItemsOptRequest()
	}
	response = NewModifyTemplateItemsOptResponse()
	err = c.Send(request, response)
	return
}

func NewModifyInternalRocketMQNamespaceRequest() (request *ModifyInternalRocketMQNamespaceRequest) {
	request = &ModifyInternalRocketMQNamespaceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyInternalRocketMQNamespace")
	return
}

func NewModifyInternalRocketMQNamespaceResponse() (response *ModifyInternalRocketMQNamespaceResponse) {
	response = &ModifyInternalRocketMQNamespaceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此为内部接口，用于运营端修改命名空间属性
func (c *Client) ModifyInternalRocketMQNamespace(request *ModifyInternalRocketMQNamespaceRequest) (response *ModifyInternalRocketMQNamespaceResponse, err error) {
	if request == nil {
		request = NewModifyInternalRocketMQNamespaceRequest()
	}
	response = NewModifyInternalRocketMQNamespaceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyTemplateOptRequest() (request *ModifyTemplateOptRequest) {
	request = &ModifyTemplateOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyTemplateOpt")
	return
}

func NewModifyTemplateOptResponse() (response *ModifyTemplateOptResponse) {
	response = &ModifyTemplateOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改配置模板
func (c *Client) ModifyTemplateOpt(request *ModifyTemplateOptRequest) (response *ModifyTemplateOptResponse, err error) {
	if request == nil {
		request = NewModifyTemplateOptRequest()
	}
	response = NewModifyTemplateOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQClustersRequest() (request *DescribeRocketMQClustersRequest) {
	request = &DescribeRocketMQClustersRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQClusters")
	return
}

func NewDescribeRocketMQClustersResponse() (response *DescribeRocketMQClustersResponse) {
	response = &DescribeRocketMQClustersResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取RocketMQ集群列表
func (c *Client) DescribeRocketMQClusters(request *DescribeRocketMQClustersRequest) (response *DescribeRocketMQClustersResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQClustersRequest()
	}
	response = NewDescribeRocketMQClustersResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteCmqSubscribeRequest() (request *DeleteCmqSubscribeRequest) {
	request = &DeleteCmqSubscribeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteCmqSubscribe")
	return
}

func NewDeleteCmqSubscribeResponse() (response *DeleteCmqSubscribeResponse) {
	response = &DeleteCmqSubscribeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除cmq订阅
func (c *Client) DeleteCmqSubscribe(request *DeleteCmqSubscribeRequest) (response *DeleteCmqSubscribeResponse, err error) {
	if request == nil {
		request = NewDeleteCmqSubscribeRequest()
	}
	response = NewDeleteCmqSubscribeResponse()
	err = c.Send(request, response)
	return
}

func NewModifyCmqSubscriptionAttributeRequest() (request *ModifyCmqSubscriptionAttributeRequest) {
	request = &ModifyCmqSubscriptionAttributeRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyCmqSubscriptionAttribute")
	return
}

func NewModifyCmqSubscriptionAttributeResponse() (response *ModifyCmqSubscriptionAttributeResponse) {
	response = &ModifyCmqSubscriptionAttributeResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改cmq订阅属性
func (c *Client) ModifyCmqSubscriptionAttribute(request *ModifyCmqSubscriptionAttributeRequest) (response *ModifyCmqSubscriptionAttributeResponse, err error) {
	if request == nil {
		request = NewModifyCmqSubscriptionAttributeRequest()
	}
	response = NewModifyCmqSubscriptionAttributeResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMachineTypeListOptRequest() (request *DescribeMachineTypeListOptRequest) {
	request = &DescribeMachineTypeListOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeMachineTypeListOpt")
	return
}

func NewDescribeMachineTypeListOptResponse() (response *DescribeMachineTypeListOptResponse) {
	response = &DescribeMachineTypeListOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查看机型列表
func (c *Client) DescribeMachineTypeListOpt(request *DescribeMachineTypeListOptRequest) (response *DescribeMachineTypeListOptResponse, err error) {
	if request == nil {
		request = NewDescribeMachineTypeListOptRequest()
	}
	response = NewDescribeMachineTypeListOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQTopicMsgsRequest() (request *DescribeRocketMQTopicMsgsRequest) {
	request = &DescribeRocketMQTopicMsgsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQTopicMsgs")
	return
}

func NewDescribeRocketMQTopicMsgsResponse() (response *DescribeRocketMQTopicMsgsResponse) {
	response = &DescribeRocketMQTopicMsgsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// rocketmq 消息查询
func (c *Client) DescribeRocketMQTopicMsgs(request *DescribeRocketMQTopicMsgsRequest) (response *DescribeRocketMQTopicMsgsResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQTopicMsgsRequest()
	}
	response = NewDescribeRocketMQTopicMsgsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRocketMQMsgTraceRequest() (request *DescribeRocketMQMsgTraceRequest) {
	request = &DescribeRocketMQMsgTraceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRocketMQMsgTrace")
	return
}

func NewDescribeRocketMQMsgTraceResponse() (response *DescribeRocketMQMsgTraceResponse) {
	response = &DescribeRocketMQMsgTraceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询消息轨迹
func (c *Client) DescribeRocketMQMsgTrace(request *DescribeRocketMQMsgTraceRequest) (response *DescribeRocketMQMsgTraceResponse, err error) {
	if request == nil {
		request = NewDescribeRocketMQMsgTraceRequest()
	}
	response = NewDescribeRocketMQMsgTraceResponse()
	err = c.Send(request, response)
	return
}

func NewSendMsgRequest() (request *SendMsgRequest) {
	request = &SendMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "SendMsg")
	return
}

func NewSendMsgResponse() (response *SendMsgResponse) {
	response = &SendMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此接口仅用于测试发生消息，不能作为现网正式生产使用
func (c *Client) SendMsg(request *SendMsgRequest) (response *SendMsgResponse, err error) {
	if request == nil {
		request = NewSendMsgRequest()
	}
	response = NewSendMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeInternalRocketMQGroupsRequest() (request *DescribeInternalRocketMQGroupsRequest) {
	request = &DescribeInternalRocketMQGroupsRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeInternalRocketMQGroups")
	return
}

func NewDescribeInternalRocketMQGroupsResponse() (response *DescribeInternalRocketMQGroupsResponse) {
	response = &DescribeInternalRocketMQGroupsResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 此为内部接口，用于运营端查询用户消费组信息
func (c *Client) DescribeInternalRocketMQGroups(request *DescribeInternalRocketMQGroupsRequest) (response *DescribeInternalRocketMQGroupsResponse, err error) {
	if request == nil {
		request = NewDescribeInternalRocketMQGroupsRequest()
	}
	response = NewDescribeInternalRocketMQGroupsResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRealTimeSubscriptionRequest() (request *DescribeRealTimeSubscriptionRequest) {
	request = &DescribeRealTimeSubscriptionRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRealTimeSubscription")
	return
}

func NewDescribeRealTimeSubscriptionResponse() (response *DescribeRealTimeSubscriptionResponse) {
	response = &DescribeRealTimeSubscriptionResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询指定环境和主题下的单个订阅者实时信息
func (c *Client) DescribeRealTimeSubscription(request *DescribeRealTimeSubscriptionRequest) (response *DescribeRealTimeSubscriptionResponse, err error) {
	if request == nil {
		request = NewDescribeRealTimeSubscriptionRequest()
	}
	response = NewDescribeRealTimeSubscriptionResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeAMQPCreateQuotaRequest() (request *DescribeAMQPCreateQuotaRequest) {
	request = &DescribeAMQPCreateQuotaRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeAMQPCreateQuota")
	return
}

func NewDescribeAMQPCreateQuotaResponse() (response *DescribeAMQPCreateQuotaResponse) {
	response = &DescribeAMQPCreateQuotaResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取用户的配额，如Queue容量，Exchange容量，Vhost容量，单Vhost Tps数,剩余可创建集群数
func (c *Client) DescribeAMQPCreateQuota(request *DescribeAMQPCreateQuotaRequest) (response *DescribeAMQPCreateQuotaResponse, err error) {
	if request == nil {
		request = NewDescribeAMQPCreateQuotaRequest()
	}
	response = NewDescribeAMQPCreateQuotaResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeClusterDetailOptRequest() (request *DescribeClusterDetailOptRequest) {
	request = &DescribeClusterDetailOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeClusterDetailOpt")
	return
}

func NewDescribeClusterDetailOptResponse() (response *DescribeClusterDetailOptResponse) {
	response = &DescribeClusterDetailOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取物理集群详情
func (c *Client) DescribeClusterDetailOpt(request *DescribeClusterDetailOptRequest) (response *DescribeClusterDetailOptResponse, err error) {
	if request == nil {
		request = NewDescribeClusterDetailOptRequest()
	}
	response = NewDescribeClusterDetailOptResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeMsgRequest() (request *DescribeMsgRequest) {
	request = &DescribeMsgRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeMsg")
	return
}

func NewDescribeMsgResponse() (response *DescribeMsgResponse) {
	response = &DescribeMsgResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消息详情
func (c *Client) DescribeMsg(request *DescribeMsgRequest) (response *DescribeMsgResponse, err error) {
	if request == nil {
		request = NewDescribeMsgRequest()
	}
	response = NewDescribeMsgResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeNamespaceBundlesOptRequest() (request *DescribeNamespaceBundlesOptRequest) {
	request = &DescribeNamespaceBundlesOptRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeNamespaceBundlesOpt")
	return
}

func NewDescribeNamespaceBundlesOptResponse() (response *DescribeNamespaceBundlesOptResponse) {
	response = &DescribeNamespaceBundlesOptResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 运营端获取命名空间bundle列表
func (c *Client) DescribeNamespaceBundlesOpt(request *DescribeNamespaceBundlesOptRequest) (response *DescribeNamespaceBundlesOptResponse, err error) {
	if request == nil {
		request = NewDescribeNamespaceBundlesOptRequest()
	}
	response = NewDescribeNamespaceBundlesOptResponse()
	err = c.Send(request, response)
	return
}


func NewDescribeRabbitMQNodeListRequest() (request *DescribeRabbitMQNodeListRequest) {
	request = &DescribeRabbitMQNodeListRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQNodeList")
	return
}

func NewDescribeRabbitMQNodeListResponse() (response *DescribeRabbitMQNodeListResponse) {
	response = &DescribeRabbitMQNodeListResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// RabbitMQ专享版查询节点列表
func (c *Client) DescribeRabbitMQNodeList(request *DescribeRabbitMQNodeListRequest) (response *DescribeRabbitMQNodeListResponse, err error) {
	if request == nil {
		request = NewDescribeRabbitMQNodeListRequest()
	}
	response = NewDescribeRabbitMQNodeListResponse()
	err = c.Send(request, response)
	return
}


func NewDescribeRabbitMQVipInstancesRequest() (request *DescribeRabbitMQVipInstancesRequest) {
	request = &DescribeRabbitMQVipInstancesRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVipInstances")
	return
}

func NewDescribeRabbitMQVipInstancesResponse() (response *DescribeRabbitMQVipInstancesResponse) {
	response = &DescribeRabbitMQVipInstancesResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询用户已购的RabbitMQ专享实例列表
func (c *Client) DescribeRabbitMQVipInstances(request *DescribeRabbitMQVipInstancesRequest) (response *DescribeRabbitMQVipInstancesResponse, err error) {
	if request == nil {
		request = NewDescribeRabbitMQVipInstancesRequest()
	}
	response = NewDescribeRabbitMQVipInstancesResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRabbitMQUserRequest() (request *CreateRabbitMQUserRequest) {
	request = &CreateRabbitMQUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQUser")
	return
}

func NewCreateRabbitMQUserResponse() (response *CreateRabbitMQUserResponse) {
	response = &CreateRabbitMQUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建RabbitMQ的用户
func (c *Client) CreateRabbitMQUser(request *CreateRabbitMQUserRequest) (response *CreateRabbitMQUserResponse, err error) {
	if request == nil {
		request = NewCreateRabbitMQUserRequest()
	}
	response = NewCreateRabbitMQUserResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRabbitMQUserRequest() (request *ModifyRabbitMQUserRequest) {
	request = &ModifyRabbitMQUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQUser")
	return
}

func NewModifyRabbitMQUserResponse() (response *ModifyRabbitMQUserResponse) {
	response = &ModifyRabbitMQUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改RabbitMQ的用户
func (c *Client) ModifyRabbitMQUser(request *ModifyRabbitMQUserRequest) (response *ModifyRabbitMQUserResponse, err error) {
	if request == nil {
		request = NewModifyRabbitMQUserRequest()
	}
	response = NewModifyRabbitMQUserResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRabbitMQUserRequest() (request *DescribeRabbitMQUserRequest) {
	request = &DescribeRabbitMQUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQUser")
	return
}

func NewDescribeRabbitMQUserResponse() (response *DescribeRabbitMQUserResponse) {
	response = &DescribeRabbitMQUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询RabbitMQ用户列表
func (c *Client) DescribeRabbitMQUser(request *DescribeRabbitMQUserRequest) (response *DescribeRabbitMQUserResponse, err error) {
	if request == nil {
		request = NewDescribeRabbitMQUserRequest()
	}
	response = NewDescribeRabbitMQUserResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRabbitMQUserRequest() (request *DeleteRabbitMQUserRequest) {
	request = &DeleteRabbitMQUserRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQUser")
	return
}

func NewDeleteRabbitMQUserResponse() (response *DeleteRabbitMQUserResponse) {
	response = &DeleteRabbitMQUserResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RabbitMQ的用户
func (c *Client) DeleteRabbitMQUser(request *DeleteRabbitMQUserRequest) (response *DeleteRabbitMQUserResponse, err error) {
	if request == nil {
		request = NewDeleteRabbitMQUserRequest()
	}
	response = NewDeleteRabbitMQUserResponse()
	err = c.Send(request, response)
	return
}


func NewDescribeRabbitMQVipInstanceRequest() (request *DescribeRabbitMQVipInstanceRequest) {
	request = &DescribeRabbitMQVipInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVipInstance")
	return
}

func NewDescribeRabbitMQVipInstanceResponse() (response *DescribeRabbitMQVipInstanceResponse) {
	response = &DescribeRabbitMQVipInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 获取单个RabbitMQ专享实例信息
func (c *Client) DescribeRabbitMQVipInstance(request *DescribeRabbitMQVipInstanceRequest) (response *DescribeRabbitMQVipInstanceResponse, err error) {
	if request == nil {
		request = NewDescribeRabbitMQVipInstanceRequest()
	}
	response = NewDescribeRabbitMQVipInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRabbitMQVipInstanceRequest() (request *DeleteRabbitMQVipInstanceRequest) {
	request = &DeleteRabbitMQVipInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQVipInstance")
	return
}

func NewDeleteRabbitMQVipInstanceResponse() (response *DeleteRabbitMQVipInstanceResponse) {
	response = &DeleteRabbitMQVipInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RabbitMQ专享版实例
func (c *Client) DeleteRabbitMQVipInstance(request *DeleteRabbitMQVipInstanceRequest) (response *DeleteRabbitMQVipInstanceResponse, err error) {
	if request == nil {
		request = NewDeleteRabbitMQVipInstanceRequest()
	}
	response = NewDeleteRabbitMQVipInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRabbitMQVipInstanceRequest() (request *CreateRabbitMQVipInstanceRequest) {
	request = &CreateRabbitMQVipInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQVipInstance")
	return
}

func NewCreateRabbitMQVipInstanceResponse() (response *CreateRabbitMQVipInstanceResponse) {
	response = &CreateRabbitMQVipInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建RabbitMQ专享版实例
func (c *Client) CreateRabbitMQVipInstance(request *CreateRabbitMQVipInstanceRequest) (response *CreateRabbitMQVipInstanceResponse, err error) {
	if request == nil {
		request = NewCreateRabbitMQVipInstanceRequest()
	}
	response = NewCreateRabbitMQVipInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewModifyRabbitMQVipInstanceRequest() (request *ModifyRabbitMQVipInstanceRequest) {
	request = &ModifyRabbitMQVipInstanceRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQVipInstance")
	return
}

func NewModifyRabbitMQVipInstanceResponse() (response *ModifyRabbitMQVipInstanceResponse) {
	response = &ModifyRabbitMQVipInstanceResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改RabbitMQ专享版实例
func (c *Client) ModifyRabbitMQVipInstance(request *ModifyRabbitMQVipInstanceRequest) (response *ModifyRabbitMQVipInstanceResponse, err error) {
	if request == nil {
		request = NewModifyRabbitMQVipInstanceRequest()
	}
	response = NewModifyRabbitMQVipInstanceResponse()
	err = c.Send(request, response)
	return
}

func NewCreateRabbitMQVirtualHostRequest() (request *CreateRabbitMQVirtualHostRequest) {
	request = &CreateRabbitMQVirtualHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "CreateRabbitMQVirtualHost")
	return
}

func NewCreateRabbitMQVirtualHostResponse() (response *CreateRabbitMQVirtualHostResponse) {
	response = &CreateRabbitMQVirtualHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 创建RabbitMQ的vhost
func (c *Client) CreateRabbitMQVirtualHost(request *CreateRabbitMQVirtualHostRequest) (response *CreateRabbitMQVirtualHostResponse, err error) {
	if request == nil {
		request = NewCreateRabbitMQVirtualHostRequest()
	}
	response = NewCreateRabbitMQVirtualHostResponse()
	err = c.Send(request, response)
	return
}

func NewDescribeRabbitMQVirtualHostRequest() (request *DescribeRabbitMQVirtualHostRequest) {
	request = &DescribeRabbitMQVirtualHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DescribeRabbitMQVirtualHost")
	return
}

func NewDescribeRabbitMQVirtualHostResponse() (response *DescribeRabbitMQVirtualHostResponse) {
	response = &DescribeRabbitMQVirtualHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 查询RabbitMQ vhost列表
func (c *Client) DescribeRabbitMQVirtualHost(request *DescribeRabbitMQVirtualHostRequest) (response *DescribeRabbitMQVirtualHostResponse, err error) {
	if request == nil {
		request = NewDescribeRabbitMQVirtualHostRequest()
	}
	response = NewDescribeRabbitMQVirtualHostResponse()
	err = c.Send(request, response)
	return
}


func NewModifyRabbitMQVirtualHostRequest() (request *ModifyRabbitMQVirtualHostRequest) {
	request = &ModifyRabbitMQVirtualHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "ModifyRabbitMQVirtualHost")
	return
}

func NewModifyRabbitMQVirtualHostResponse() (response *ModifyRabbitMQVirtualHostResponse) {
	response = &ModifyRabbitMQVirtualHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 修改RabbitMQ的vhost
func (c *Client) ModifyRabbitMQVirtualHost(request *ModifyRabbitMQVirtualHostRequest) (response *ModifyRabbitMQVirtualHostResponse, err error) {
	if request == nil {
		request = NewModifyRabbitMQVirtualHostRequest()
	}
	response = NewModifyRabbitMQVirtualHostResponse()
	err = c.Send(request, response)
	return
}

func NewDeleteRabbitMQVirtualHostRequest() (request *DeleteRabbitMQVirtualHostRequest) {
	request = &DeleteRabbitMQVirtualHostRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("tdmq", APIVersion, "DeleteRabbitMQVirtualHost")
	return
}

func NewDeleteRabbitMQVirtualHostResponse() (response *DeleteRabbitMQVirtualHostResponse) {
	response = &DeleteRabbitMQVirtualHostResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除RabbitMQ的vhost
func (c *Client) DeleteRabbitMQVirtualHost(request *DeleteRabbitMQVirtualHostRequest) (response *DeleteRabbitMQVirtualHostResponse, err error) {
	if request == nil {
		request = NewDeleteRabbitMQVirtualHostRequest()
	}
	response = NewDeleteRabbitMQVirtualHostResponse()
	err = c.Send(request, response)
	return
}

type DeleteRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRabbitMQVirtualHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
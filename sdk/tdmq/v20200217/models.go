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

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ModifyInternalTenantQuotaRequest struct {
	*tchttp.BaseRequest

	// 命名空间配额

	MaxNamespaces *int64 `json:"MaxNamespaces,omitempty" name:"MaxNamespaces"`
	// Topic配额

	MaxTopics *int64 `json:"MaxTopics,omitempty" name:"MaxTopics"`
	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// Exchange配额

	MaxExchanges *int64 `json:"MaxExchanges,omitempty" name:"MaxExchanges"`
	// Queue配额

	MaxQueues *int64 `json:"MaxQueues,omitempty" name:"MaxQueues"`
	// Topic分区配额

	MaxPartitions *int64 `json:"MaxPartitions,omitempty" name:"MaxPartitions"`
	// 命名空间最大堆积，byte为单位

	MaxMsgBacklog *uint64 `json:"MaxMsgBacklog,omitempty" name:"MaxMsgBacklog"`
	// 最长保留时间，毫秒为单位

	MaxRetention *uint64 `json:"MaxRetention,omitempty" name:"MaxRetention"`
	// 客户APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
}

func (r *ModifyInternalTenantQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInternalTenantQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPExchangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPExchangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleListOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置模块列表

		ConfigModules []*ConfigModuleInfoOpt `json:"ConfigModules,omitempty" name:"ConfigModules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleListOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleListOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyAntiAffinityGroupOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// bookie主亲和组

	AffinityGroupPrimary *string `json:"AffinityGroupPrimary,omitempty" name:"AffinityGroupPrimary"`
	// bookie次亲和组

	AffinityGroupSecondary *string `json:"AffinityGroupSecondary,omitempty" name:"AffinityGroupSecondary"`
}

func (r *UpdatePolicyAntiAffinityGroupOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyAntiAffinityGroupOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueueDetailRequest struct {
	*tchttp.BaseRequest

	// 精确匹配QueueName

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DescribeCmqQueueDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqQueueDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleListOptRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeModuleListOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleListOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterMetaInitScriptOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *GetClusterMetaInitScriptOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterMetaInitScriptOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitNamespaceBundleOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SplitNamespaceBundleOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SplitNamespaceBundleOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqSubscriptionDetailRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据SubscriptionName进行模糊搜索

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *DescribeCmqSubscriptionDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqSubscriptionDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqQueueAttributeRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。

	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`
	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。

	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`
	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。

	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`
	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。

	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`
	// 消息保留周期。取值范围 60-1296000 秒（1min-15天），默认值 345600 (4 天)。

	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`
	// 消息最长回溯时间，取值范围0-msgRetentionSeconds，消息的最大回溯之间为消息在队列中的保存周期，0表示不开启消息回溯。

	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`
	// 第一次查询时间

	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`
	// 最大查询次数

	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`
	// 死信队列名称

	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`
	// MaxTimeToLivepolicy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds

	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`
	// 最大接收次数

	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`
	// 死信队列策略

	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`
	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。

	Trace *bool `json:"Trace,omitempty" name:"Trace"`
	// 是否开启事务，1开启，0不开启

	Transaction *uint64 `json:"Transaction,omitempty" name:"Transaction"`
}

func (r *ModifyCmqQueueAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmqQueueAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInternalRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 存储保留时间，按小时

	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`
}

func (r *ModifyInternalRocketMQNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInternalRocketMQNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDeliveryStatusOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 模块名称

		ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
		// 配置下发状态信息

		AgentDeliveryInfos []*DeliveryInfoOpt `json:"AgentDeliveryInfos,omitempty" name:"AgentDeliveryInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDeliveryStatusOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeliveryStatusOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublisherStatsSetOpt struct {

	// 消息生产速率

	MsgRateIn *float64 `json:"MsgRateIn,omitempty" name:"MsgRateIn"`
	// 消息生产吞吐量

	MsgThroughputIn *float64 `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`
	// 平均消息大小

	AverageMsgSize *float64 `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`
	// 分块消息发送速率

	ChunkedMessageRate *float64 `json:"ChunkedMessageRate,omitempty" name:"ChunkedMessageRate"`
	// 生产者ID

	ProducerId *int64 `json:"ProducerId,omitempty" name:"ProducerId"`
}

type TopicStatsSetOpt struct {

	// 消息生产速率

	MsgRateIn *float64 `json:"MsgRateIn,omitempty" name:"MsgRateIn"`
	// 消息生产吞吐量

	MsgThroughputIn *float64 `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`
	// 平均消息大小

	AverageMsgSize *float64 `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`
	// 入字节总数

	BytesInCounter *int64 `json:"BytesInCounter,omitempty" name:"BytesInCounter"`
	// 入消息总数

	MsgInCounter *int64 `json:"MsgInCounter,omitempty" name:"MsgInCounter"`
	// 消息消费速率

	MsgRateOut *float64 `json:"MsgRateOut,omitempty" name:"MsgRateOut"`
	// 消息消费吞吐量

	MsgThroughputOut *float64 `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`
	// 出字节总数

	BytesOutCounter *int64 `json:"BytesOutCounter,omitempty" name:"BytesOutCounter"`
	// 出消息总数

	MsgOutCounter *int64 `json:"MsgOutCounter,omitempty" name:"MsgOutCounter"`
	// 已经发送的分块消息数

	MsgChunkPublished *bool `json:"MsgChunkPublished,omitempty" name:"MsgChunkPublished"`
	// 积压消息存储大小（单位byte）

	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`
	// 积压消息数

	BacklogSize *int64 `json:"BacklogSize,omitempty" name:"BacklogSize"`
	// 生产者状态数组

	Publishers []*PublisherStatsSetOpt `json:"Publishers,omitempty" name:"Publishers"`
	// 订阅状态数组

	SubscriptionSets []*SubscriptionStatsSetOpt `json:"SubscriptionSets,omitempty" name:"SubscriptionSets"`
	// 跨集群复制状态数组

	ReplicationSets []*ConsumerStatsSetOpt `json:"ReplicationSets,omitempty" name:"ReplicationSets"`
	// 分区数量，非分区主题取值为-1

	PartitionIndexes *int64 `json:"PartitionIndexes,omitempty" name:"PartitionIndexes"`
}

type CmqClusterMembers struct {

	// TopicId

	Id *string `json:"Id,omitempty" name:"Id"`
	// 用户id

	SId *string `json:"SId,omitempty" name:"SId"`
	// app id

	Appid *string `json:"Appid,omitempty" name:"Appid"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 类型：1:发布订阅；0:

	Type *string `json:"Type,omitempty" name:"Type"`
	// 类型

	MinorType *string `json:"MinorType,omitempty" name:"MinorType"`
	// 最大消息长度

	MaxMsgSize *string `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`
	// 最大保留时间

	MsgRetentionSeconds *string `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`
	// 单一

	Single *uint64 `json:"Single,omitempty" name:"Single"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户类型

	UserType *string `json:"UserType,omitempty" name:"UserType"`
	// 队列名

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type CmqDeadLetterSource struct {

	// 消息队列ID。

	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`
	// 消息队列名字。

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

type ClearCmqSubscriptionFilterTagsRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *ClearCmqSubscriptionFilterTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearCmqSubscriptionFilterTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitNamespaceBundleOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名，与Broker实例ID二选一

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// bundle名

	Bundle *string `json:"Bundle,omitempty" name:"Bundle"`
	// Broker实例ID，与物理集群名二选一

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *SplitNamespaceBundleOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SplitNamespaceBundleOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMaxProducerPerTopicOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyMaxProducerPerTopicOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMaxProducerPerTopicOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功删除的角色名称数组。

		RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQConsumerClient struct {

	// 客户端ID

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 客户端地址

	ClientAddr *string `json:"ClientAddr,omitempty" name:"ClientAddr"`
}

type DescribeReplicationClustersRequest struct {
	*tchttp.BaseRequest

	// 目标集群所在地域Id

	TargetRegionId *uint64 `json:"TargetRegionId,omitempty" name:"TargetRegionId"`
}

func (r *DescribeReplicationClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReplicationClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendCmqMsgRequest struct {
	*tchttp.BaseRequest

	// 队列名

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
	// 消息内容

	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`
	// 延迟时间

	DelaySeconds *int64 `json:"DelaySeconds,omitempty" name:"DelaySeconds"`
}

func (r *SendCmqMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendCmqMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopicMsgsRequest struct {
	*tchttp.BaseRequest

	// 集群 ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称，查询死信时为groupId

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息 ID

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 消息 key

	MsgKey *string `json:"MsgKey,omitempty" name:"MsgKey"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询偏移

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限额

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 标志一次分页事务

	TaskRequestId *string `json:"TaskRequestId,omitempty" name:"TaskRequestId"`
	// 死信查询时该值为true，只对Rocketmq有效

	QueryDlqMsg *bool `json:"QueryDlqMsg,omitempty" name:"QueryDlqMsg"`
	// 查询最近N条消息 最大不超过1024，默认-1为其他查询条件

	NumOfLatestMsg *int64 `json:"NumOfLatestMsg,omitempty" name:"NumOfLatestMsg"`
}

func (r *DescribeRocketMQTopicMsgsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQTopicMsgsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRocketMQMessageDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息id

		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
		// 消息生成时间戳

		BornTimestamp *int64 `json:"BornTimestamp,omitempty" name:"BornTimestamp"`
		// 消息存储时间戳

		StoreTimestamp *int64 `json:"StoreTimestamp,omitempty" name:"StoreTimestamp"`
		// 消息生产客户端地址

		BornHost *string `json:"BornHost,omitempty" name:"BornHost"`
		// 消息Tag

		MsgTag *string `json:"MsgTag,omitempty" name:"MsgTag"`
		// 消息Key

		MsgKey *string `json:"MsgKey,omitempty" name:"MsgKey"`
		// 消息属性

		Properties *string `json:"Properties,omitempty" name:"Properties"`
		// 消息重试次数

		ReConsumeTimes *uint64 `json:"ReConsumeTimes,omitempty" name:"ReConsumeTimes"`
		// Base64编码格式字符串

		MsgBody *string `json:"MsgBody,omitempty" name:"MsgBody"`
		// 消息内容的CRC32 Code

		MsgBodyCRC *int64 `json:"MsgBodyCRC,omitempty" name:"MsgBodyCRC"`
		// 消息体大小（单位K）
		// 当大于2048时不返回消息

		MsgBodySize *uint64 `json:"MsgBodySize,omitempty" name:"MsgBodySize"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExportRocketMQMessageDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRocketMQMessageDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishCmqMsgRequest struct {
	*tchttp.BaseRequest

	// 主题名

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息内容

	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`
	// 消息标签

	MsgTag []*string `json:"MsgTag,omitempty" name:"MsgTag"`
}

func (r *PublishCmqMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishCmqMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendConfigOptRequest struct {
	*tchttp.BaseRequest

	// 节点IP列表

	NodeIpList []*string `json:"NodeIpList,omitempty" name:"NodeIpList"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 配置版本号

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 是否灰度

	IsGrey *bool `json:"IsGrey,omitempty" name:"IsGrey"`
}

func (r *SendConfigOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendConfigOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateClusterOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQInstanceConfig struct {

	// 单命名空间TPS上线

	MaxTpsPerNamespace *uint64 `json:"MaxTpsPerNamespace,omitempty" name:"MaxTpsPerNamespace"`
	// 最大命名空间数量

	MaxNamespaceNum *uint64 `json:"MaxNamespaceNum,omitempty" name:"MaxNamespaceNum"`
	// 已使用命名空间数量

	UsedNamespaceNum *uint64 `json:"UsedNamespaceNum,omitempty" name:"UsedNamespaceNum"`
	// 最大Topic数量

	MaxTopicNum *uint64 `json:"MaxTopicNum,omitempty" name:"MaxTopicNum"`
	// 已使用Topic数量

	UsedTopicNum *uint64 `json:"UsedTopicNum,omitempty" name:"UsedTopicNum"`
	// 最大Group数量

	MaxGroupNum *uint64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`
	// 已使用Group数量

	UsedGroupNum *uint64 `json:"UsedGroupNum,omitempty" name:"UsedGroupNum"`
	// 集群类型

	ConfigDisplay *string `json:"ConfigDisplay,omitempty" name:"ConfigDisplay"`
	// 集群节点数

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 节点分布情况

	NodeDistribution []*InstanceNodeDistribution `json:"NodeDistribution,omitempty" name:"NodeDistribution"`
	// topic分布情况

	TopicDistribution []*RocketMQTopicDistribution `json:"TopicDistribution,omitempty" name:"TopicDistribution"`
	// 每个主题最大队列数

	MaxQueuesPerTopic *uint64 `json:"MaxQueuesPerTopic,omitempty" name:"MaxQueuesPerTopic"`
}

type DescribeBindVpcsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Vpc集合。

		VpcSets []*VpcBindRecord `json:"VpcSets,omitempty" name:"VpcSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindVpcsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBindVpcsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPulsarResourceAllocationRuleRequest struct {
	*tchttp.BaseRequest

	// 集群分配的规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 集群当前状态，0表示下线，1表示可用

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 备注信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// Pulsar软件版本

	PulsarVersion *string `json:"PulsarVersion,omitempty" name:"PulsarVersion"`
	// Pulsar管理端接口地址

	AdminUrl *string `json:"AdminUrl,omitempty" name:"AdminUrl"`
	// Pulsar管理端接口token

	AdminToken *string `json:"AdminToken,omitempty" name:"AdminToken"`
	// 协议端管理接口地址

	ExternalAdminUrl *string `json:"ExternalAdminUrl,omitempty" name:"ExternalAdminUrl"`
	// 协议端管理接口token

	ExternalAdminToken *string `json:"ExternalAdminToken,omitempty" name:"ExternalAdminToken"`
	// 接入点协议

	EndPointProtocol *string `json:"EndPointProtocol,omitempty" name:"EndPointProtocol"`
	// VPC接入点地址通用后缀

	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`
	// 公网接入点地址通用后缀

	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *ModifyPulsarResourceAllocationRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPulsarResourceAllocationRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TenantTopSetOpt struct {

	// 租户ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 租户名

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// 监控指标复杂类型

	Metrics *CommonMetricsOpt `json:"Metrics,omitempty" name:"Metrics"`
	// 协议名

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type ModifyCmqQueueAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqQueueAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmqQueueAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMaxProducerPerTopicOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 生产者数量限制

	MaxProducer *int64 `json:"MaxProducer,omitempty" name:"MaxProducer"`
}

func (r *UpdatePolicyMaxProducerPerTopicOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMaxProducerPerTopicOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQEnvironmentRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 角色名称数组。

		RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQEnvironmentRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmqPublisherInfo struct {

	// 添加时间

	AddTimeStamp *string `json:"AddTimeStamp,omitempty" name:"AddTimeStamp"`
	// 节点id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 节点ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 修改时间

	ModTimeStamp *string `json:"ModTimeStamp,omitempty" name:"ModTimeStamp"`
	// 状态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 节点vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 节点vport

	Vport *string `json:"Vport,omitempty" name:"Vport"`
}

type NamespacePolicyRateLimitOpt struct {

	// 消息限制量

	Rate *int64 `json:"Rate,omitempty" name:"Rate"`
	// 流量限制量

	Throughput *int64 `json:"Throughput,omitempty" name:"Throughput"`
	// 限流时间窗口，单位为秒

	TimePeriod *int64 `json:"TimePeriod,omitempty" name:"TimePeriod"`
}

type DescribeRocketMQRolesRequest struct {
	*tchttp.BaseRequest

	// 角色名称，模糊查询

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 起始下标，不填默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 必填字段，集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmqDeadLetterPolicy struct {

	// 死信队列。

	DeadLetterQueue *string `json:"DeadLetterQueue,omitempty" name:"DeadLetterQueue"`
	// 死信队列策略。

	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`
	// 最大未消费过期时间。Policy为1时必选。范围300-43200，单位秒，需要小于消息最大保留时间MsgRetentionSeconds。

	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`
	// 最大接收次数。

	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`
}

type EnvironmentRole struct {

	// 环境（命名空间）。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 角色名称。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 授权项，最多只能包含produce、consume两项的非空字符串数组。

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
	// 角色描述。

	RoleDescribe *string `json:"RoleDescribe,omitempty" name:"RoleDescribe"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateAMQPVHostRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// vhost名称，3-64个字符，只能包含字母、数字、“-”及“_”

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 未消费消息的保留时间，以毫秒为单位，60秒-15天

	MsgTtl *uint64 `json:"MsgTtl,omitempty" name:"MsgTtl"`
	// 说明，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateAMQPVHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPVHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功删除的环境（命名空间）数组。

		EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEnvironmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterConfigTypeOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *DescribeClusterConfigTypeOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterConfigTypeOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQInstanceRequest struct {
	*tchttp.BaseRequest

	// 专享实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 实例消息保留时间，小时为单位

	MessageRetention *int64 `json:"MessageRetention,omitempty" name:"MessageRetention"`
}

func (r *ModifyRocketMQInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePortalConfigRequest struct {
	*tchttp.BaseRequest

	// 集群版本号，可提交多个版本号进行查询

	Versions []*string `json:"Versions,omitempty" name:"Versions"`
}

func (r *DescribePortalConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePortalConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnloadNamespaceBundleOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// bundle名

	Bundle *string `json:"Bundle,omitempty" name:"Bundle"`
	// Broker实例ID，与物理集群名二选一

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *UnloadNamespaceBundleOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnloadNamespaceBundleOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyRocketMQConsumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyRocketMQConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyRocketMQConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterComponentMetricsSetOpt struct {

	// 组件ip

	Address *string `json:"Address,omitempty" name:"Address"`
	// 组件监控指标值数组（JSONArray字符串）

	MetricsList *string `json:"MetricsList,omitempty" name:"MetricsList"`
}

type DescribeAMQPClustersRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 按照集群ID关键字搜索

	IdKeyword *string `json:"IdKeyword,omitempty" name:"IdKeyword"`
	// 按照集群名称关键字搜索

	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`
}

func (r *DescribeAMQPClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceBundlesOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 是否需要监控指标，若传false，则不需要传Limit和Offset分页参数

	NeedMetrics *bool `json:"NeedMetrics,omitempty" name:"NeedMetrics"`
	// 查询限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeNamespaceBundlesOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespaceBundlesOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateOptRequest struct {
	*tchttp.BaseRequest

	// 配置模块名称（小写）

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 0-S5 1-SA

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
	// 是否shell脚本文件

	IsShell *bool `json:"IsShell,omitempty" name:"IsShell"`
	// 脚本内容

	ShellScript *string `json:"ShellScript,omitempty" name:"ShellScript"`
}

func (r *CreateTemplateOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPQueueRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 队列名称

	Queue *string `json:"Queue,omitempty" name:"Queue"`
}

func (r *DeleteAMQPQueueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPQueueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmqSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQEnvironmentRolesRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 角色名称数组。

	RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`
	// 必填字段，集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteRocketMQEnvironmentRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQEnvironmentRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBatchRoleRequest struct {
	*tchttp.BaseRequest

	// 指定集群id校验

	ClusterScope *string `json:"ClusterScope,omitempty" name:"ClusterScope"`
	// 批量绑定数据

	BatchRoleBindData []*RoleBindData `json:"BatchRoleBindData,omitempty" name:"BatchRoleBindData"`
}

func (r *CreateBatchRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBatchRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRoleRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 角色名称。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 授权项，最多只能包含produce、consume两项的非空字符串数组。

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
	// 必填字段，集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateEnvironmentRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEnvironmentRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyOffloadThresholdOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 卸载阈值，单位为bytes

	OffloadThreshold *int64 `json:"OffloadThreshold,omitempty" name:"OffloadThreshold"`
}

func (r *UpdatePolicyOffloadThresholdOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyOffloadThresholdOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Role struct {

	// 角色名称。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色token值。

	Token *string `json:"Token,omitempty" name:"Token"`
	// 备注说明。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateCmqTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主题id

		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmqTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 3-64个字符，只能包含字母、数字、“-”及“_”

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 说明信息，不超过128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAMQPClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAMQPClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPExchange struct {

	// Exchange名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Exchange的类别，为枚举类型:Direct, Fanout, Topic

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主绑定数

	SourceBindedNum *uint64 `json:"SourceBindedNum,omitempty" name:"SourceBindedNum"`
	// 说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 被绑定数

	DestBindedNum *uint64 `json:"DestBindedNum,omitempty" name:"DestBindedNum"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否为内部Exchange(以amq.前缀开头的)

	Internal *bool `json:"Internal,omitempty" name:"Internal"`
}

type InitializeClusterConfigOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建信息

		Result *string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitializeClusterConfigOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitializeClusterConfigOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQMsgTraceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// [
		//     {
		//         "Stage": "produce",
		//         "Data": {
		//             "ProducerName": "生产者名",
		//             "ProduceTime": "消息生产时间",
		//             "ProducerAddr": "客户端地址",
		//             "Duration": "耗时ms",
		//             "Status": "状态（0：成功，1：失败）"
		//         }
		//     },
		//     {
		//         "Stage": "persist",
		//         "Data": {
		//             "PersistTime": "存储时间",
		//             "Duration": "耗时ms",
		//             "Status": "状态（0：成功，1：失败）"
		//         }
		//     },
		//     {
		//         "Stage": "consume",
		//         "Data": {
		//             "TotalCount": 2,
		//             "RocketMqConsumeLogs": [
		//                 {
		//                     "ConsumerGroup": "消费组",
		//                     "ConsumeModel": "消费模式",
		//                     "ConsumerAddr": "消费者地址",
		//                     "ConsumeTime": "推送时间",
		//                     "Status": "状态（0:已推送未确认, 2:已确认, 3:转入重试, 4:已重试未确认, 5:已转入死信队列）"
		//                 },
		//                 {
		//                     "ConsumerGroup": "消费组",
		//                     "ConsumeModel": "消费模式",
		//                     "ConsumerAddr": "消费者地址",
		//                     "ConsumeTime": "推送时间",
		//                     "Status": "状态（0:已推送未确认, 2:已确认, 3:转入重试, 4:已重试未确认, 5:已转入死信队列）"
		//                 }
		//             ]
		//         }
		//     }
		// ]

		Result []*TraceResult `json:"Result,omitempty" name:"Result"`
		// 消息轨迹页展示的topic名称

		ShowTopicName *string `json:"ShowTopicName,omitempty" name:"ShowTopicName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQMsgTraceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQMsgTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicReplicationClustersRequest struct {
	*tchttp.BaseRequest

	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 环境（命名空间）名称

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 跨地域复制目标集群列表

	TargetClusterList []*ReplicationClusterSets `json:"TargetClusterList,omitempty" name:"TargetClusterList"`
	// 是否全部清空绑定关系

	UnBindAllTargetCluster *bool `json:"UnBindAllTargetCluster,omitempty" name:"UnBindAllTargetCluster"`
}

func (r *ModifyTopicReplicationClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicReplicationClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcConfig struct {

	// vpc的id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type RocketMQMsgLog struct {

	// 消息id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 消息tag

	MsgTag *string `json:"MsgTag,omitempty" name:"MsgTag"`
	// 消息key

	MsgKey *string `json:"MsgKey,omitempty" name:"MsgKey"`
	// 客户端地址

	ProducerAddr *string `json:"ProducerAddr,omitempty" name:"ProducerAddr"`
	// 消息发送时间

	ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`
	// pulsar消息id

	PulsarMsgId *string `json:"PulsarMsgId,omitempty" name:"PulsarMsgId"`
	// 死信重发次数

	DeadLetterResendTimes *int64 `json:"DeadLetterResendTimes,omitempty" name:"DeadLetterResendTimes"`
	// 死信重发成功次数

	ResendSuccessCount *int64 `json:"ResendSuccessCount,omitempty" name:"ResendSuccessCount"`
}

type SendCmqMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true表示发送成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 消息id

		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCmqMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendCmqMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TpoDescribeResourcesAdminProjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 项目列表

		ProjectSet *ResourceAdminProject `json:"ProjectSet,omitempty" name:"ProjectSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TpoDescribeResourcesAdminProjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TpoDescribeResourcesAdminProjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterConfigTypeOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置类型列表

		ConfigType []*ConfigTypeInfoOpt `json:"ConfigType,omitempty" name:"ConfigType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterConfigTypeOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterConfigTypeOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendConfigOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 每个节点配置下发结果

		SendFileResult []*SendConfigInfoOpt `json:"SendFileResult,omitempty" name:"SendFileResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendConfigOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendConfigOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConfigOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求是否成功

		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateConfigOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConfigOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		ClusterInfo *RocketMQClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`
		// 集群配置

		ClusterConfig *RocketMQClusterConfig `json:"ClusterConfig,omitempty" name:"ClusterConfig"`
		// 集群最近使用量

		ClusterStats *RocketMQClusterRecentStats `json:"ClusterStats,omitempty" name:"ClusterStats"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopicsRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 按主题类型过滤查询结果，可选择Normal, GlobalOrder, PartitionedOrder, Transaction

	FilterType []*string `json:"FilterType,omitempty" name:"FilterType"`
	// 按主题名称搜索，支持模糊查询

	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`
}

func (r *DescribeRocketMQTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQClusterConfig struct {

	// 单命名空间TPS上线

	MaxTpsPerNamespace *uint64 `json:"MaxTpsPerNamespace,omitempty" name:"MaxTpsPerNamespace"`
	// 最大命名空间数量

	MaxNamespaceNum *uint64 `json:"MaxNamespaceNum,omitempty" name:"MaxNamespaceNum"`
	// 已使用命名空间数量

	UsedNamespaceNum *uint64 `json:"UsedNamespaceNum,omitempty" name:"UsedNamespaceNum"`
	// 最大Topic数量

	MaxTopicNum *uint64 `json:"MaxTopicNum,omitempty" name:"MaxTopicNum"`
	// 已使用Topic数量

	UsedTopicNum *uint64 `json:"UsedTopicNum,omitempty" name:"UsedTopicNum"`
	// 最大Group数量

	MaxGroupNum *uint64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`
	// 已使用Group数量

	UsedGroupNum *uint64 `json:"UsedGroupNum,omitempty" name:"UsedGroupNum"`
	// 消息最大保留时间，以毫秒为单位

	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitempty" name:"MaxRetentionTime"`
	// 消息最长延时，以毫秒为单位

	MaxLatencyTime *uint64 `json:"MaxLatencyTime,omitempty" name:"MaxLatencyTime"`
	// 单个主题最大队列数

	MaxQueuesPerTopic *uint64 `json:"MaxQueuesPerTopic,omitempty" name:"MaxQueuesPerTopic"`
}

type DescribeAMQPVHostConnectionsRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
}

func (r *DescribeAMQPVHostConnectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPVHostConnectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTenantTopListOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Top租户列表

		TenantTopSet []*TenantTopSetOpt `json:"TenantTopSet,omitempty" name:"TenantTopSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTenantTopListOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTenantTopListOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqSubscriptionAttributeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 向 Endpoint 推送消息出现错误时，CMQ 推送服务器的重试策略。取值如下：
	// （1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息。
	// （2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始1s，后面是2s，4s，8s···由于 Topic 消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是 EXPONENTIAL_DECAY_RETRY。

	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`
	// 推送内容的格式。取值：（1）JSON；（2）SIMPLIFIED，即 raw 格式。如果 Protocol 是 queue，则取值必须为 SIMPLIFIED。如果 Protocol 是 HTTP，两个值均可以，默认值是 JSON。

	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
	// 消息正文。消息标签（用于消息过滤)。标签数量不能超过5个，每个标签不超过16个字符。与(Batch)PublishMessage的MsgTag参数配合使用，规则：1）如果FilterTag没有设置，则无论MsgTag是否有设置，订阅接收所有发布到Topic的消息；2）如果FilterTag数组有值，则只有数组中至少有一个值在MsgTag数组中也存在时（即FilterTag和MsgTag有交集），订阅才接收该发布到Topic的消息；3）如果FilterTag数组有值，但MsgTag没设置，则不接收任何发布到Topic的消息，可以认为是2）的一种特例，此时FilterTag和MsgTag没有交集。规则整体的设计思想是以订阅者的意愿为主。

	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags"`
	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。

	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`
}

func (r *ModifyCmqSubscriptionAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmqSubscriptionAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 分区数，必须大于或者等于原分区数，若想维持原分区数请输入原数目，修改分区数仅对非全局顺序消息起效果，不允许超过128个分区。

	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`
	// 备注，128字符以内。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindClustersRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeBindClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBindClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReplicationClusterSets struct {

	// 目标集群Id

	TargetClusterId *string `json:"TargetClusterId,omitempty" name:"TargetClusterId"`
	// 目标集群名字

	TargetClusterName *string `json:"TargetClusterName,omitempty" name:"TargetClusterName"`
	// 目标集群地域Id

	TargetRegionId *uint64 `json:"TargetRegionId,omitempty" name:"TargetRegionId"`
	// 状态，0:关闭,1:开启

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 集群接入地址

	BrokerServiceUrl *string `json:"BrokerServiceUrl,omitempty" name:"BrokerServiceUrl"`
	// 集群接入地址

	WebServiceUrl *string `json:"WebServiceUrl,omitempty" name:"WebServiceUrl"`
}

type DeleteBindVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户Vpc Id。

		UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
		// Vpc子网Id。

		UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBindVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBindVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyBatchRoleRequest struct {
	*tchttp.BaseRequest

	// 指定集群id校验

	ClusterScope *string `json:"ClusterScope,omitempty" name:"ClusterScope"`
	// 批量绑定数据

	BatchRoleBindData []*RoleBindData `json:"BatchRoleBindData,omitempty" name:"BatchRoleBindData"`
}

func (r *VerifyBatchRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyBatchRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Topic struct {

	// 最后一次间隔内发布消息的平均byte大小。

	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`
	// 消费者数量。

	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`
	// 被记录下来的消息总数。

	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitempty" name:"LastConfirmedEntry"`
	// 最后一个ledger创建的时间。

	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitempty" name:"LastLedgerCreatedTimestamp"`
	// 本地和复制的发布者每秒发布消息的速率。

	MsgRateIn *string `json:"MsgRateIn,omitempty" name:"MsgRateIn"`
	// 本地和复制的消费者每秒分发消息的数量之和。

	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`
	// 本地和复制的发布者每秒发布消息的byte。

	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`
	// 本地和复制的消费者每秒分发消息的byte。

	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`
	// 被记录下来的消息总数。

	NumberOfEntries *string `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`
	// 分区数<=0：topic下无子分区。

	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`
	// 生产者数量。

	ProducerCount *string `json:"ProducerCount,omitempty" name:"ProducerCount"`
	// 以byte计算的所有消息存储总量。

	TotalSize *string `json:"TotalSize,omitempty" name:"TotalSize"`
	// 分区topic里面的子分区。

	SubTopicSets []*PartitionsTopic `json:"SubTopicSets,omitempty" name:"SubTopicSets"`
	// topic类型描述：
	// 0：普通消息；
	// 1：全局顺序消息；
	// 2：局部顺序消息；
	// 3：重试队列；
	// 4：死信队列；
	// 5：事务消息。

	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`
	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 说明，128个字符以内。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近修改时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 生产者上限。

	ProducerLimit *string `json:"ProducerLimit,omitempty" name:"ProducerLimit"`
	// 消费者上限。

	ConsumerLimit *string `json:"ConsumerLimit,omitempty" name:"ConsumerLimit"`
	// 0: 非持久非分区 1: 非持久分区 2: 持久非分区 3: 持久分区

	PulsarTopicType *int64 `json:"PulsarTopicType,omitempty" name:"PulsarTopicType"`
	// topic所在集群的消息同步状态OriginReplicationOpen  源集群开启消息同步OriginReplicationClose  源集群关闭消息同步TargetReplicationOpen  目标集群开启消息同步TargetReplicationClose  目标集群关闭消息同步

	ReplicationStatus *string `json:"ReplicationStatus,omitempty" name:"ReplicationStatus"`
}

type ModifyRocketMQEnvironmentRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 角色名称。

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 授权项。

		Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQEnvironmentRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 备注说明，长度必须大等于0且小等于128。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 必填字段，集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateRocketMQRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPExchangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPExchangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建结果。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterBrokerNetworkOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 网络协议，取值为DescribeClusterEndpointsOpt接口出参的协议名

	NetworkProtocol *string `json:"NetworkProtocol,omitempty" name:"NetworkProtocol"`
}

func (r *DescribeClusterBrokerNetworkOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterBrokerNetworkOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterBaradMetricsOptRequest struct {
	*tchttp.BaseRequest

	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 指标名，取值：
	// topic_count：topic-partition总数
	// msg_backlog：消息堆积数
	// storage_size：消息堆积大小

	MetricsName *string `json:"MetricsName,omitempty" name:"MetricsName"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 指标间隔（单位为秒）

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *GetClusterBaradMetricsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterBaradMetricsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutClusterOnlineOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutClusterOnlineOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutClusterOnlineOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQSubscriptionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 订阅关系列表

		Subscriptions []*RocketMQSubscription `json:"Subscriptions,omitempty" name:"Subscriptions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQSubscriptionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalRocketMQInstancesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件，支持以下：
	// 1. Name，实例名称模糊查询
	// 2. InstanceId，实例ID查询
	// 3. AppId, 用户APPID查询
	// 4. Uin，用户UIN查询
	// 5. ClusterId, 集群查询
	//

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询起始位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询最大数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInternalRocketMQInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalRocketMQInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerLog struct {

	// 消息ID。

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 消费组。

	ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
	// 消费组名称。

	ConsumerName *string `json:"ConsumerName,omitempty" name:"ConsumerName"`
	// 消费时间。

	ConsumeTime *string `json:"ConsumeTime,omitempty" name:"ConsumeTime"`
	// 消费者客户端地址。

	ConsumerAddr *string `json:"ConsumerAddr,omitempty" name:"ConsumerAddr"`
	// 消费耗时（毫秒）。

	ConsumeUseTime *uint64 `json:"ConsumeUseTime,omitempty" name:"ConsumeUseTime"`
	// 消费状态。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type RocketMQTopic struct {

	// 主题名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主题的类别，为枚举类型，Normal，GlobalOrder，PartitionedOrder，Transaction，Retry及DeadLetter

	Type *string `json:"Type,omitempty" name:"Type"`
	// 订阅组数量

	GroupNum *uint64 `json:"GroupNum,omitempty" name:"GroupNum"`
	// 说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 读写分区数

	PartitionNum *uint64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type CreateRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQConsumerConnectionsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 消费组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 对查询结果排序，此为排序字段，目前支持Accumulative（消息堆积量）

	SortedBy *string `json:"SortedBy,omitempty" name:"SortedBy"`
	// 查询结果排序规则，ASC为升序，DESC为降序

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeRocketMQConsumerConnectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQConsumerConnectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQAppIdStatsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 客户类型：
		//
		// 0: 普通用户
		// 1: 腾讯内部非自研上云客户
		// 2: 外部大客户（KA）客户
		// 3: 内部云梯自研上云客户

		CustomerType *int64 `json:"CustomerType,omitempty" name:"CustomerType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQAppIdStatsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQAppIdStatsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigBaseInfoOpt struct {

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 程序路径

	RootPath *string `json:"RootPath,omitempty" name:"RootPath"`
	// 数据路径

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// 配置路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 配置文件名

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
}

type RocketMQClusterRecentStats struct {

	// Topic数量

	TopicNum *uint64 `json:"TopicNum,omitempty" name:"TopicNum"`
	// API请求数

	ApiRequestNum *uint64 `json:"ApiRequestNum,omitempty" name:"ApiRequestNum"`
	// 消息生产数

	ProducedMsgNum *uint64 `json:"ProducedMsgNum,omitempty" name:"ProducedMsgNum"`
	// 消息消费数

	ConsumedMsgNum *uint64 `json:"ConsumedMsgNum,omitempty" name:"ConsumedMsgNum"`
	// 消息堆积数

	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`
}

type DeleteAMQPVHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPVHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPVHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceRulesOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 规则总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则列表

		ResourceRuleSets []*ResourceRuleSetOpt `json:"ResourceRuleSets,omitempty" name:"ResourceRuleSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceRulesOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceRulesOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyBatchRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0,检测通过，1 检测不通过

		CheckCode *uint64 `json:"CheckCode,omitempty" name:"CheckCode"`
		// 枚举  原因

		CheckMessage *string `json:"CheckMessage,omitempty" name:"CheckMessage"`
		// 异常资源列表

		AbnormalResources []*string `json:"AbnormalResources,omitempty" name:"AbnormalResources"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyBatchRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyBatchRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQMsgRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题，查询死信时传groupId

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// pulsar消息id

	PulsarMsgId *string `json:"PulsarMsgId,omitempty" name:"PulsarMsgId"`
	// 查询死信时该值为true，只对Rocketmq有效

	QueryDlqMsg *bool `json:"QueryDlqMsg,omitempty" name:"QueryDlqMsg"`
}

func (r *DescribeRocketMQMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 主题名。

		TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
		// 0：非分区topic，无分区；非0：具体分区topic的分区数。

		Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`
		// 备注，128字符以内。

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 0： 普通消息；
		// 1 ：全局顺序消息；
		// 2 ：局部顺序消息；
		// 3 ：重试队列；
		// 4 ：死信队列；
		// 5 ：事务消息。

		TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendBatchMessagesRequest struct {
	*tchttp.BaseRequest

	// 消息要发送的topic的名字, 这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 需要发送消息的内容

	Payload *string `json:"Payload,omitempty" name:"Payload"`
	// String 类型的 token，可以不填，系统会自动获取

	StringToken *string `json:"StringToken,omitempty" name:"StringToken"`
	// producer 的名字，要求全局是唯一的，如果不设置，系统会自动生成

	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`
	// 单位：s。消息发送的超时时间。默认值为：30s

	SendTimeout *int64 `json:"SendTimeout,omitempty" name:"SendTimeout"`
	// 内存中允许缓存的生产消息的最大数量，默认值：1000条

	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitempty" name:"MaxPendingMessages"`
	// 每一个batch中消息的最大数量，默认值：1000条/batch

	BatchingMaxMessages *int64 `json:"BatchingMaxMessages,omitempty" name:"BatchingMaxMessages"`
	// 每一个batch最大等待的时间，超过这个时间，不管是否达到指定的batch中消息的数量和大小，都会将该batch发送出去，默认：10ms

	BatchingMaxPublishDelay *int64 `json:"BatchingMaxPublishDelay,omitempty" name:"BatchingMaxPublishDelay"`
	// 每一个batch中最大允许的消息的大小，默认：128KB

	BatchingMaxBytes *int64 `json:"BatchingMaxBytes,omitempty" name:"BatchingMaxBytes"`
}

func (r *SendBatchMessagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendBatchMessagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPulsarResourceAllocationRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPulsarResourceAllocationRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPulsarResourceAllocationRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyRocketMQConsumeRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 消费组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 消息id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 客户端ID

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// topic名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *VerifyRocketMQConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyRocketMQConsumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQVipInstanceStateRequest struct {
	*tchttp.BaseRequest

	// 专享集群实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 仅支持在变配失败状态下传1

	State *int64 `json:"State,omitempty" name:"State"`
}

func (r *ModifyRocketMQVipInstanceStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQVipInstanceStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称，不支持中字以及除了短线和下划线外的特殊字符且不超过16个字符。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 未消费消息过期时间，单位：秒，最小60，最大1296000，（15天）。

	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`
	// 说明，128个字符以内。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 消息保留策略

	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

func (r *CreateEnvironmentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEnvironmentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllTemplateModuleOptRequest struct {
	*tchttp.BaseRequest

	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 0-S5 1-SA

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
}

func (r *DescribeAllTemplateModuleOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllTemplateModuleOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总数

		Total *uint64 `json:"Total,omitempty" name:"Total"`
		// 集群列表

		Pclusters []*string `json:"Pclusters,omitempty" name:"Pclusters"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQConsumerConnectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 在线消费者信息

		Connections []*RocketMQConsumerConnection `json:"Connections,omitempty" name:"Connections"`
		// 订阅组信息

		GroupDetail *RocketMQGroup `json:"GroupDetail,omitempty" name:"GroupDetail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQConsumerConnectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQConsumerConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBindVpcRequest struct {
	*tchttp.BaseRequest

	// 租户Vpc Id。

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 租户Vpc子网Id。

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// 说明，128个字符以内。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateBindVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBindVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPCreateQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAMQPCreateQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPCreateQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishCmqMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// true表示发送成功

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 消息id

		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PublishCmqMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishCmqMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQInstanceSpecRequest struct {
	*tchttp.BaseRequest

	// 专享实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例规格，
	// rocket-vip-basic-1 基础型
	// rocket-vip-basic-2 标准型
	// rocket-vip-basic-3 高阶Ⅰ型
	// rocket-vip-basic-4 高阶Ⅱ型

	Specification *string `json:"Specification,omitempty" name:"Specification"`
	// 节点数量

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 存储空间，GB为单位

	StorageSize *uint64 `json:"StorageSize,omitempty" name:"StorageSize"`
}

func (r *ModifyRocketMQInstanceSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQInstanceSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Subscription struct {

	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 消费者开始连接的时间。

	ConnectedSince *string `json:"ConnectedSince,omitempty" name:"ConnectedSince"`
	// 消费者地址。

	ConsumerAddr *string `json:"ConsumerAddr,omitempty" name:"ConsumerAddr"`
	// 消费者数量。

	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`
	// 消费者名称。

	ConsumerName *string `json:"ConsumerName,omitempty" name:"ConsumerName"`
	// 堆积的消息数量。

	MsgBacklog *string `json:"MsgBacklog,omitempty" name:"MsgBacklog"`
	// 于TTL，此订阅下没有被发送而是被丢弃的比例。

	MsgRateExpired *string `json:"MsgRateExpired,omitempty" name:"MsgRateExpired"`
	// 消费者每秒分发消息的数量之和。

	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`
	// 消费者每秒消息的byte。

	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`
	// 订阅名称。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 消费者集合。

	ConsumerSets []*Consumer `json:"ConsumerSets,omitempty" name:"ConsumerSets"`
	// 是否在线。

	IsOnline *bool `json:"IsOnline,omitempty" name:"IsOnline"`
	// 消费进度集合。

	ConsumersScheduleSets []*ConsumersSchedule `json:"ConsumersScheduleSets,omitempty" name:"ConsumersScheduleSets"`
	// 备注。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近修改时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否由于未 ack 数到达上限而被 block 注意：此字段可能返回 null，表示取不到有效值。

	BlockedSubscriptionOnUnackedMsgs *bool `json:"BlockedSubscriptionOnUnackedMsgs,omitempty" name:"BlockedSubscriptionOnUnackedMsgs"`
	// 订阅类型，Exclusive，Shared，Failover， Key_Shared，空或NULL表示未知， 注意：此字段可能返回 null，表示取不到有效值。

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 未 ack 消息数上限 注意：此字段可能返回 null，表示取不到有效值。

	MaxUnackedMsgNum *int64 `json:"MaxUnackedMsgNum,omitempty" name:"MaxUnackedMsgNum"`
}

type ResetRocketMQConsumerOffSetRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 消费组名称

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 主题名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 重置方式，0表示从最新位点开始，1表示从指定时间点开始

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 重置指定的时间戳，仅在 Type 为1是生效，以毫秒为单位

	ResetTimestamp *uint64 `json:"ResetTimestamp,omitempty" name:"ResetTimestamp"`
}

func (r *ResetRocketMQConsumerOffSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetRocketMQConsumerOffSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQInstanceDeliveryProgressRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRocketMQInstanceDeliveryProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQInstanceDeliveryProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqSubscribeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 订阅的协议，目前支持两种协议：http、queue。使用http协议，用户需自己搭建接受消息的web server。使用queue，消息会自动推送到CMQ queue，用户可以并发地拉取消息。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 接收通知的Endpoint，根据协议Protocol区分：对于http，Endpoint必须以“`http://`”开头，host可以是域名或IP；对于Queue，则填QueueName。 请注意，目前推送服务不能推送到私有网络中，因此Endpoint填写为私有网络域名或地址将接收不到推送的消息，目前支持推送到公网和基础网络。

	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`
	// 向Endpoint推送消息出现错误时，CMQ推送服务器的重试策略。取值有：1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息；2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始1s，后面是2s，4s，8s...由于Topic消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是EXPONENTIAL_DECAY_RETRY。

	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`
	// 消息正文。消息标签（用于消息过滤)。标签数量不能超过5个，每个标签不超过16个字符。与(Batch)PublishMessage的MsgTag参数配合使用，规则：1）如果FilterTag没有设置，则无论MsgTag是否有设置，订阅接收所有发布到Topic的消息；2）如果FilterTag数组有值，则只有数组中至少有一个值在MsgTag数组中也存在时（即FilterTag和MsgTag有交集），订阅才接收该发布到Topic的消息；3）如果FilterTag数组有值，但MsgTag没设置，则不接收任何发布到Topic的消息，可以认为是2）的一种特例，此时FilterTag和MsgTag没有交集。规则整体的设计思想是以订阅者的意愿为主。

	FilterTag []*string `json:"FilterTag,omitempty" name:"FilterTag"`
	// BindingKey数量不超过5个， 每个BindingKey长度不超过64字节，该字段表示订阅接收消息的过滤策略，每个BindingKey最多含有15个“.”， 即最多16个词组。

	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`
	// 推送内容的格式。取值：1）JSON；2）SIMPLIFIED，即raw格式。如果Protocol是queue，则取值必须为SIMPLIFIED。如果Protocol是http，两个值均可以，默认值是JSON。

	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

func (r *CreateCmqSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmqSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DeleteRocketMQNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群列表数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群信息列表

		ClusterSet []*Cluster `json:"ClusterSet,omitempty" name:"ClusterSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 角色数组。

		RoleSets []*Role `json:"RoleSets,omitempty" name:"RoleSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutClusterOfflineOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *PutClusterOfflineOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutClusterOfflineOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceRulesOptRequest struct {
	*tchttp.BaseRequest

	// 如填写，则按照协议类型进行查询。数组内容为协议ID

	Types []*int64 `json:"Types,omitempty" name:"Types"`
	// 如填写，则按照用户APPID进行精确查询

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 限制条目数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeResourceRulesOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceRulesOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQConsumerConnectionDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 消费端主题信息列表

		Details []*RocketMQConsumerTopic `json:"Details,omitempty" name:"Details"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQConsumerConnectionDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQConsumerConnectionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置项（BASE64编码后）

		ConfigItems *string `json:"ConfigItems,omitempty" name:"ConfigItems"`
		// 配置基础信息

		ConfigBaseInfo *ConfigBaseInfoOpt `json:"ConfigBaseInfo,omitempty" name:"ConfigBaseInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentNodesOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 代理节点列表

		AgentNodes []*AgentNodeOpt `json:"AgentNodes,omitempty" name:"AgentNodes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentNodesOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentNodesOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQConsumerClientsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 消费组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 客户端id信息，用于模糊查询

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
}

func (r *DescribeRocketMQConsumerClientsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQConsumerClientsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcknowledgeMessageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 如果为“”，则说明没有错误返回

		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcknowledgeMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcknowledgeMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群的详细信息

		ClusterSet *Cluster `json:"ClusterSet,omitempty" name:"ClusterSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyEncryptionRequiredOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyEncryptionRequiredOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyEncryptionRequiredOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicySchemaValidationEnforcedResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicySchemaValidationEnforcedResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicySchemaValidationEnforcedResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqDeadLetterSourceQueuesRequest struct {
	*tchttp.BaseRequest

	// 死信队列名称

	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`
	// 分页时本页获取主题列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页时本页获取主题的个数，如果不传递该参数，则该参数默认为20，最大值为50。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 根据SourceQueueName过滤

	SourceQueueName *string `json:"SourceQueueName,omitempty" name:"SourceQueueName"`
}

func (r *DescribeCmqDeadLetterSourceQueuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqDeadLetterSourceQueuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQVipInstancesRequest struct {
	*tchttp.BaseRequest

	// 查询条件过滤器，支持的查询条件如下：
	// instanceIds - 实例ID
	// instanceName - 实例名称

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询数目上限，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询起始位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeRocketMQVipInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQVipInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterOverBookieConsumeOpt struct {

	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 数据消耗（GB）

	DataStorage *float64 `json:"DataStorage,omitempty" name:"DataStorage"`
}

type RewindCmqQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
	// 设定该时间，则（Batch）receiveMessage接口，会按照生产消息的先后顺序消费该时间戳以后的消息。

	StartConsumeTime *uint64 `json:"StartConsumeTime,omitempty" name:"StartConsumeTime"`
}

func (r *RewindCmqQueueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RewindCmqQueueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVersionsOptRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClusterVersionsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVersionsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalNamespacesRequest struct {
	*tchttp.BaseRequest

	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 客户APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 租户id

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名称，支持模糊匹配

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInternalNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQVipInstance struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例版本

	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`
	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4 - 异常, 5 - 发货失败，6 - 变配中，7 - 变配失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 节点数量

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 实例配置规格名称

	ConfigDisplay *string `json:"ConfigDisplay,omitempty" name:"ConfigDisplay"`
	// 峰值TPS

	MaxTps *uint64 `json:"MaxTps,omitempty" name:"MaxTps"`
	// 峰值带宽，Mbps为单位

	MaxBandWidth *uint64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`
	// 存储容量，GB为单位

	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 实例到期时间，毫秒为单位

	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费)， 1表示自动续费，2表示明确不自动续费(用户设置)

	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 0-后付费，1-预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 实例配置ID

	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`
	// 是否可以垂直变配

	CanVerticalChange *bool `json:"CanVerticalChange,omitempty" name:"CanVerticalChange"`
	// 最大可设置消息保留时间，小时为单位

	MaxRetention *int64 `json:"MaxRetention,omitempty" name:"MaxRetention"`
	// 最小可设置消息保留时间，小时为单位

	MinRetention *int64 `json:"MinRetention,omitempty" name:"MinRetention"`
	// 实例消息保留时间，小时为单位

	Retention *int64 `json:"Retention,omitempty" name:"Retention"`
}

type DescribeDeliveryStatusOptRequest struct {
	*tchttp.BaseRequest

	// 代理IP列表

	AgentIpList []*string `json:"AgentIpList,omitempty" name:"AgentIpList"`
	// 是否检查所有

	CheckAll *bool `json:"CheckAll,omitempty" name:"CheckAll"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 配置版本号

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribeDeliveryStatusOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDeliveryStatusOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 订阅组列表

		Groups []*RocketMQGroup `json:"Groups,omitempty" name:"Groups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyPublishRateOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 生产流量限制，单位为bytes

	PublishThroughput *int64 `json:"PublishThroughput,omitempty" name:"PublishThroughput"`
	// 生产速率限制

	PublishRate *int64 `json:"PublishRate,omitempty" name:"PublishRate"`
}

func (r *UpdatePolicyPublishRateOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyPublishRateOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateItemsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功

		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTemplateItemsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateItemsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPVHostsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 按名称搜索

	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`
}

func (r *DescribeAMQPVHostsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPVHostsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBookieDiskListOptRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBookieDiskListOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBookieDiskListOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeHealthOptRequest struct {
	*tchttp.BaseRequest

	// 节点实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeNodeHealthOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeHealthOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 租户列表

		TenantSet []*TenantSetOpt `json:"TenantSet,omitempty" name:"TenantSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTenantsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPQueueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAMQPQueueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAMQPQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentRoleRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 角色名称。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 授权项，最多只能包含produce、consume两项的非空字符串数组。

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
	// 必填字段，集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyEnvironmentRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEnvironmentRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmqTaskInfo struct {

	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 状态值

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type MachineDictItemOpt struct {

	// 机型

	ItemKey *string `json:"ItemKey,omitempty" name:"ItemKey"`
	// 编号

	ItemValue *int64 `json:"ItemValue,omitempty" name:"ItemValue"`
}

type DescribeClusterDetailRequest struct {
	*tchttp.BaseRequest

	// 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPExchangeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAMQPExchangeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAMQPExchangeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendBatchMessagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息的唯一标识

		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`
		// 错误消息，返回为 ""，代表没有错误

		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendBatchMessagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendBatchMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllVersionConfigsOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DescribeAllVersionConfigsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllVersionConfigsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalTenantsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 虚拟集群列表

		Tenants []*InternalTenant `json:"Tenants,omitempty" name:"Tenants"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternalTenantsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalTenantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPQueueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPQueueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订阅id

		SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmqSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 未消费消息过期时间，单位：秒。

		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`
		// 说明，128个字符以内。

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 命名空间ID

		NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnvironmentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEnvironmentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPQueueConsumersRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// queue名称

	Queue *string `json:"Queue,omitempty" name:"Queue"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAMQPQueueConsumersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPQueueConsumersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceAdminProject struct {

	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 项目名

	Organization *string `json:"Organization,omitempty" name:"Organization"`
}

type DescribeAMQPExchangesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 主题信息列表

		Exchanges []*AMQPExchange `json:"Exchanges,omitempty" name:"Exchanges"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPExchangesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPExchangesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineTypeListOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机型列表

		MachineTypes []*MachineDictItemOpt `json:"MachineTypes,omitempty" name:"MachineTypes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineTypeListOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineTypeListOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQEnvironmentRoleRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 角色名称。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 授权项，最多只能包含produce、consume两项的非空字符串数组。

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
	// 必填字段，集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateRocketMQEnvironmentRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQEnvironmentRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPVHostRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
}

func (r *DeleteAMQPVHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPVHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindCmqDeadLetterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindCmqDeadLetterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindCmqDeadLetterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicRecord struct {

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type UpdatePolicySubscribeRateOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicySubscribeRateOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicySubscribeRateOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQConsumerTopic struct {

	// 主题名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 主题类型，Default表示普通，GlobalOrder表示全局顺序，PartitionedOrder表示局部顺序，Transaction表示事务，Retry表示重试，DeadLetter表示死信

	Type *string `json:"Type,omitempty" name:"Type"`
	// 分区数

	PartitionNum *uint64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
	// 消息堆积数

	Accumulative *int64 `json:"Accumulative,omitempty" name:"Accumulative"`
	// 最后消费时间，以毫秒为单位

	LastConsumptionTime *uint64 `json:"LastConsumptionTime,omitempty" name:"LastConsumptionTime"`
	// 订阅规则

	SubRule *string `json:"SubRule,omitempty" name:"SubRule"`
}

type DescribeRolesRequest struct {
	*tchttp.BaseRequest

	// 角色名称，模糊查询

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 起始下标，不填默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 必填字段，集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQEnvironmentRoleRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 角色名称。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 授权项，最多只能包含produce、consume两项的非空字符串数组。

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
	// 必填字段，集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyRocketMQEnvironmentRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQEnvironmentRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigDictOptRequest struct {
	*tchttp.BaseRequest

	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

func (r *DescribeConfigDictOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigDictOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQCreateQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRocketMQCreateQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQCreateQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateItemOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否Shell脚本文件

		IsShell *bool `json:"IsShell,omitempty" name:"IsShell"`
		// 配置模块名称（小写）

		ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
		// 脚本文件内容

		ShellScript *string `json:"ShellScript,omitempty" name:"ShellScript"`
		// 模板配置项列表

		TemplateItems []*TemplateItemInfoOpt `json:"TemplateItems,omitempty" name:"TemplateItems"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateItemOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateItemOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportRocketMQMessageDetailRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 应用命名空间

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// Topic名称
	// 如果是死信消息 isDlqMsg=true

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 是否死信消息

	DeadLetterMsg *bool `json:"DeadLetterMsg,omitempty" name:"DeadLetterMsg"`
	// 是否包含消息体

	IncludeMsgBody *bool `json:"IncludeMsgBody,omitempty" name:"IncludeMsgBody"`
}

func (r *ExportRocketMQMessageDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ExportRocketMQMessageDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMessagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息的messageID, 是全局唯一的，用来标识消息的元数据信息

		MessageId *string `json:"MessageId,omitempty" name:"MessageId"`
		// 返回的错误消息，如果返回为 “”，说明没有错误

		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMessagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendMessagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqQueueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearCmqQueueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentNodeOpt struct {

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 代理ID

	AgentId *string `json:"AgentId,omitempty" name:"AgentId"`
	// 代理IP

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// 协议ID

	ProcId *int64 `json:"ProcId,omitempty" name:"ProcId"`
	// 代理版本

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// JDK版本

	JavaVersion *string `json:"JavaVersion,omitempty" name:"JavaVersion"`
}

type RocketMQMessageTrack struct {

	// 消费者组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 消费状态

	ConsumeStatus *string `json:"ConsumeStatus,omitempty" name:"ConsumeStatus"`
	// 消息track类型

	TrackType *string `json:"TrackType,omitempty" name:"TrackType"`
	// 异常信息

	ExceptionDesc *string `json:"ExceptionDesc,omitempty" name:"ExceptionDesc"`
}

type DescribeAMQPQueuesRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 按队列名称搜索，支持模糊查询

	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`
	// 查询结果排序规则，ASC为升序，DESC为降序

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
	// 对查询结果排序，此为排序字段，目前支持Accumulative（消息堆积量）、Tps

	SortedBy *string `json:"SortedBy,omitempty" name:"SortedBy"`
	// 队列名称，指定此参数后将只返回该队列信息

	FilterOneQueue *string `json:"FilterOneQueue,omitempty" name:"FilterOneQueue"`
}

func (r *DescribeAMQPQueuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPQueuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopicMsgsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 消息列表

		TopicMsgLogSets []*RocketMQMsgLog `json:"TopicMsgLogSets,omitempty" name:"TopicMsgLogSets"`
		// 标志一次分页事务

		TaskRequestId *string `json:"TaskRequestId,omitempty" name:"TaskRequestId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQTopicMsgsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQTopicMsgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespaceBundlesOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// bundle列表

		BundleSet []*BundleSetOpt `json:"BundleSet,omitempty" name:"BundleSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespaceBundlesOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespaceBundlesOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 备注说明，长度必须大等于0且小等于128。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 必填字段，集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 批量操作环境

	EnvironmentRoleSets []*EnvironmentRoleSet `json:"EnvironmentRoleSets,omitempty" name:"EnvironmentRoleSets"`
	// 是否全部解绑环境

	UnbindAllEnvironment *bool `json:"UnbindAllEnvironment,omitempty" name:"UnbindAllEnvironment"`
}

func (r *ModifyRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProducerLog struct {

	// 消息ID。

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 生产者名称。

	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`
	// 消息生产时间。

	ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`
	// 生产者客户端。

	ProducerAddr *string `json:"ProducerAddr,omitempty" name:"ProducerAddr"`
	// 生产耗时（秒）。

	ProduceUseTime *uint64 `json:"ProduceUseTime,omitempty" name:"ProduceUseTime"`
	// 状态。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyRocketMQInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterOverBookieOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群消耗情况列表

		ClusterSets []*ClusterOverBookieConsumeOpt `json:"ClusterSets,omitempty" name:"ClusterSets"`
		// 记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterOverBookieOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterOverBookieOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceRuleOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateResourceRuleOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceRuleOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInternalRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInternalRocketMQNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInternalRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主题集合数组。

		TopicSets []*Topic `json:"TopicSets,omitempty" name:"TopicSets"`
		// 主题数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyPersistenceOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyPersistenceOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyPersistenceOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllVersionConfigsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群基础信息

		ConfigBaseInfos []*ConfigBaseInfoOpt `json:"ConfigBaseInfos,omitempty" name:"ConfigBaseInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllVersionConfigsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllVersionConfigsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQInstanceDeliveryProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发货各阶段以及完成情况，有以下几个阶段：
		// TASK_N_CHECK_PARAMS  检查参数
		// TASK_N_CONSTRUCT_CLUSTER_INFO  构建集群信息
		// TASK_N_CONSTRUCT_CONFIG    构建集群配置
		// TASK_N_CREATE_CLUSTER     创建集群
		// TASK_N_CHECK_STATUS_AND_STORE_META   检查集群状态
		// TASK_N_CONSTRUCT_VPC_NET     构建网络信息
		// TASK_N_BIND_VPC_DNS           绑定网络关系

		Stages []*DeliveryStage `json:"Stages,omitempty" name:"Stages"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQInstanceDeliveryProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQInstanceDeliveryProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendRocketMQMessageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 发送结果

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 消息ID

		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendRocketMQMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendRocketMQMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息属性。

		Properties *string `json:"Properties,omitempty" name:"Properties"`
		// 消息体。

		Body *string `json:"Body,omitempty" name:"Body"`
		// 批次ID。

		BatchId *string `json:"BatchId,omitempty" name:"BatchId"`
		// 生产时间。

		ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`
		// 消息ID。

		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
		// 生产者名称。

		ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterNodesOptRequest struct {
	*tchttp.BaseRequest

	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 可用区

	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 版本

	VersionNum *string `json:"VersionNum,omitempty" name:"VersionNum"`
	// 节点状态：0-异常；1-正常

	NodeState *int64 `json:"NodeState,omitempty" name:"NodeState"`
	// 节点类型：1-broker；2-bookie；3-zk

	ComponentType *int64 `json:"ComponentType,omitempty" name:"ComponentType"`
	// 节点地址模糊匹配

	Address *string `json:"Address,omitempty" name:"Address"`
	// 分页参数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 节点实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否需要节点监控指标

	NeedMetrics *bool `json:"NeedMetrics,omitempty" name:"NeedMetrics"`
}

func (r *DescribeClusterNodesOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodesOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。

	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`
	// 消息接收长轮询等待时间。取值范围 0-30 秒，默认值 0。

	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`
	// 消息可见性超时。取值范围 1-43200 秒（即12小时内），默认值 30。

	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`
	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。

	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`
	// 消息保留周期。取值范围 60-1296000 秒（1min-15天），默认值 345600 (4 天)。

	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`
	// 队列是否开启回溯消息能力，该参数取值范围0-msgRetentionSeconds,即最大的回溯时间为消息在队列中的保留周期，0表示不开启。

	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`
	// 1 表示事务队列，0 表示普通队列

	Transaction *uint64 `json:"Transaction,omitempty" name:"Transaction"`
	// 第一次回查间隔

	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`
	// 最大回查次数

	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`
	// 死信队列名称

	DeadLetterQueueName *string `json:"DeadLetterQueueName,omitempty" name:"DeadLetterQueueName"`
	// 死信策略。0为消息被多次消费未删除，1为Time-To-Live过期

	Policy *uint64 `json:"Policy,omitempty" name:"Policy"`
	// 最大接收次数 1-1000

	MaxReceiveCount *uint64 `json:"MaxReceiveCount,omitempty" name:"MaxReceiveCount"`
	// policy为1时必选。最大未消费过期时间。范围300-43200，单位秒，需要小于消息最大保留时间msgRetentionSeconds

	MaxTimeToLive *uint64 `json:"MaxTimeToLive,omitempty" name:"MaxTimeToLive"`
	// 是否开启消息轨迹追踪，当不设置字段时，默认为不开启，该字段为true表示开启，为false表示不开启

	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *CreateCmqQueueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmqQueueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQConsumerConnection struct {

	// 消费者实例ID

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 消费者实例的地址和端口

	ClientAddr *string `json:"ClientAddr,omitempty" name:"ClientAddr"`
	// 消费者应用的语言版本

	Language *string `json:"Language,omitempty" name:"Language"`
	// 消息堆积量

	Accumulative *int64 `json:"Accumulative,omitempty" name:"Accumulative"`
	// 消费端版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type ResourceProjectInfo struct {

	// 资源id

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 旧项目id

	OldProjectId *string `json:"OldProjectId,omitempty" name:"OldProjectId"`
	// 新项目id

	NewProjectId *string `json:"NewProjectId,omitempty" name:"NewProjectId"`
}

type CreateRocketMQGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *uint64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 配置模块名称

	ConfigModule *string `json:"ConfigModule,omitempty" name:"ConfigModule"`
	// 数据存放路径

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// 配置版本号

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 根路径

	RootPath *string `json:"RootPath,omitempty" name:"RootPath"`
	// 配置存放路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
}

func (r *CreateConfigOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 接入点信息数组

		AccessEndpointSets []*AccessEndpointSetOpt `json:"AccessEndpointSets,omitempty" name:"AccessEndpointSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterEndpointsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnloadNamespaceBundleOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnloadNamespaceBundleOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnloadNamespaceBundleOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TpoModifyResourceProjectRequest struct {
	*tchttp.BaseRequest

	// ResourceProjectInfo 数组，用于新增或修改资源对应的权限

	ResourceProjectInfos []*ResourceProjectInfo `json:"ResourceProjectInfos,omitempty" name:"ResourceProjectInfos"`
}

func (r *TpoModifyResourceProjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TpoModifyResourceProjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealTimeSubscriptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订阅者集合数组

		SubscriptionSets []*Subscription `json:"SubscriptionSets,omitempty" name:"SubscriptionSets"`
		// 数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRealTimeSubscriptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRealTimeSubscriptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPVHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAMQPVHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAMQPVHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumersSchedule struct {

	// 当前分区id。

	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`
	// 消息数量。

	NumberOfEntries *uint64 `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`
	// 消息积压数量。

	MsgBacklog *uint64 `json:"MsgBacklog,omitempty" name:"MsgBacklog"`
	// 消费者每秒分发消息的数量之和。

	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`
	// 消费者每秒消息的byte。

	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`
	// 超时丢弃比例。

	MsgRateExpired *string `json:"MsgRateExpired,omitempty" name:"MsgRateExpired"`
}

type CreateRocketMQMigrationTaskRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 任务类型：
	// 0，集群迁移
	// 1，导入到指定命名空间

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 待导入的主题列表

	Topics []*RocketMQTopicConfig `json:"Topics,omitempty" name:"Topics"`
	// 待导入的消费组列表

	Groups []*RocketMQGroupConfig `json:"Groups,omitempty" name:"Groups"`
	// 指定导入的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *CreateRocketMQMigrationTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQMigrationTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQPublicAccessPointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 公网接入点状态：
		// 0， 已开启
		// 1， 已关闭
		// 2，开启中
		// 3，关闭中
		// 4，修改中

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 支付状态：
		// 0, 未知
		// 1，正常
		// 2，欠费

		PayStatus *int64 `json:"PayStatus,omitempty" name:"PayStatus"`
		// 接入点地址

		AccessUrl *string `json:"AccessUrl,omitempty" name:"AccessUrl"`
		// 安全访问规则列表

		Rules []*PublicAccessRule `json:"Rules,omitempty" name:"Rules"`
		// 带宽

		Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
		// 付费模式

		PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQPublicAccessPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQPublicAccessPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscriptionsRequest struct {
	*tchttp.BaseRequest

	// 订阅关系集合，每次最多删除20个。

	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitempty" name:"SubscriptionTopicSets"`
	// pulsar集群Id。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 是否强制删除，默认为false

	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteSubscriptionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubscriptionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalRocketMQNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果的总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 分页后的命名空间信息列表

		Namespaces []*RocketMQNamespace `json:"Namespaces,omitempty" name:"Namespaces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternalRocketMQNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalRocketMQNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterComponentMetricsSetListOpt struct {

	// 指标名

	MetricsName *string `json:"MetricsName,omitempty" name:"MetricsName"`
	// 指标解析

	MetricsExplain *string `json:"MetricsExplain,omitempty" name:"MetricsExplain"`
	// 指标单位枚举，用于前端单位转换：
	// base：条数；time：条数/时间；
	// size：存储大小；bandwidth：带宽

	MetricsUnit *string `json:"MetricsUnit,omitempty" name:"MetricsUnit"`
}

type DescribeClusterComponentMetricsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组件监控指标数组

		MetricsSet []*ClusterComponentMetricsSetOpt `json:"MetricsSet,omitempty" name:"MetricsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterComponentMetricsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterComponentMetricsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 备注说明，长度必须大等于0且小等于128。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 必填字段，集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyRocketMQRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名，不支持中字以及除了短线和下划线外的特殊字符且不超过64个字符。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 0：非分区topic，无分区；非0：具体分区topic的分区数，最大不允许超过128。

	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`
	// 备注，128字符以内。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列。

	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CreateTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTenantsOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 客户APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 虚拟集群名称

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// 协议类型数组

	Protocols []*string `json:"Protocols,omitempty" name:"Protocols"`
	// 排序字段名，支持createTime，updateTime

	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
	// 升序排列ASC，降序排列DESC

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
	// 限制条目数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTenantsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTenantsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPCreateQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户总共可使用集群数量

		MaxClusterNum *uint64 `json:"MaxClusterNum,omitempty" name:"MaxClusterNum"`
		// 租户已创建集群数量

		UsedClusterNum *uint64 `json:"UsedClusterNum,omitempty" name:"UsedClusterNum"`
		// Exchange容量

		ExchangeCapacity *uint64 `json:"ExchangeCapacity,omitempty" name:"ExchangeCapacity"`
		// Queue容量

		QueueCapacity *uint64 `json:"QueueCapacity,omitempty" name:"QueueCapacity"`
		// 单Vhost TPS

		MaxTpsPerVHost *uint64 `json:"MaxTpsPerVHost,omitempty" name:"MaxTpsPerVHost"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPCreateQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPCreateQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInternalTenantQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInternalTenantQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInternalTenantQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色名称

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 备注说明

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 批量操作环境

		EnvironmentRoleSets []*EnvironmentRoleSet `json:"EnvironmentRoleSets,omitempty" name:"EnvironmentRoleSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootInstancesOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPQueuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 队列信息列表

		Queues []*AMQPQueueDetail `json:"Queues,omitempty" name:"Queues"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPQueuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVersionsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本名列表

		Versions []*string `json:"Versions,omitempty" name:"Versions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterVersionsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVersionsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PulsarResourceRuleOpt struct {

	// 集群分配规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 指定用户的APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 指定用户的UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 集群协议类型，支持的值为PULSAR，ROCKETMQ，AMQP，CMQ

	Type *string `json:"Type,omitempty" name:"Type"`
	// 集群当前状态，0表示下线，1表示可用，2表示灰度中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 备注信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则创建时间，ms为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 规则更新时间，ms为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AMQPQueueConsumer struct {

	// 消费者实例的地址和端口

	ClientAddr *string `json:"ClientAddr,omitempty" name:"ClientAddr"`
	// 消费者标签

	ClientTag *string `json:"ClientTag,omitempty" name:"ClientTag"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeProtocolsOptRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeProtocolsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtocolsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopicStatsOptRequest struct {
	*tchttp.BaseRequest

	// 完整topic（partition）名

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *GetTopicStatsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopicStatsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetMsgSubOffsetByTimestampRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅者名称。

	Subscription *string `json:"Subscription,omitempty" name:"Subscription"`
	// 时间戳，精确到毫秒。

	ToTimestamp *uint64 `json:"ToTimestamp,omitempty" name:"ToTimestamp"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ResetMsgSubOffsetByTimestampRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetMsgSubOffsetByTimestampRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQClusterInfo struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 地域信息

	Region *string `json:"Region,omitempty" name:"Region"`
	// 创建时间，毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 集群说明信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 公网接入地址

	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`
	// VPC接入地址

	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`
	// 是否支持命名空间接入点

	SupportNamespaceEndpoint *bool `json:"SupportNamespaceEndpoint,omitempty" name:"SupportNamespaceEndpoint"`
	// VPC信息

	Vpcs []*VpcConfig `json:"Vpcs,omitempty" name:"Vpcs"`
	// 是否为专享实例

	IsVip *bool `json:"IsVip,omitempty" name:"IsVip"`
	// Rocketmq集群标识

	RocketMQFlag *bool `json:"RocketMQFlag,omitempty" name:"RocketMQFlag"`
	// 计费状态，1表示正常，2表示已停服，3表示已销毁

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 欠费停服时间，毫秒为单位

	IsolateTime *int64 `json:"IsolateTime,omitempty" name:"IsolateTime"`
	// HTTP协议公网接入地址

	HttpPublicEndpoint *string `json:"HttpPublicEndpoint,omitempty" name:"HttpPublicEndpoint"`
	// HTTP协议VPC接入地址

	HttpVpcEndpoint *string `json:"HttpVpcEndpoint,omitempty" name:"HttpVpcEndpoint"`
	// TCP内部接入地址

	InternalEndpoint *string `json:"InternalEndpoint,omitempty" name:"InternalEndpoint"`
	// HTTP协议内部接入地址

	HttpInternalEndpoint *string `json:"HttpInternalEndpoint,omitempty" name:"HttpInternalEndpoint"`
}

type UpdatePolicyDeduplicationEnabledOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyDeduplicationEnabledOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyDeduplicationEnabledOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendRocketMQMessageRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// topic名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息key信息

	MsgKey *string `json:"MsgKey,omitempty" name:"MsgKey"`
	// 消息tag信息

	MsgTag *string `json:"MsgTag,omitempty" name:"MsgTag"`
	// 信息内容

	MsgBody *string `json:"MsgBody,omitempty" name:"MsgBody"`
}

func (r *SendRocketMQMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendRocketMQMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePulsarResourceAllocationRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 规则列表

		Rules []*PulsarResourceRuleOpt `json:"Rules,omitempty" name:"Rules"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePulsarResourceAllocationRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePulsarResourceAllocationRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQGroupConfig struct {

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 消费组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否开启广播消费

	ConsumeBroadcastEnable *bool `json:"ConsumeBroadcastEnable,omitempty" name:"ConsumeBroadcastEnable"`
	// 是否开启消费

	ConsumeEnable *bool `json:"ConsumeEnable,omitempty" name:"ConsumeEnable"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DescribeClusterNodesOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点信息

		NodeSet []*NodeSetOpt `json:"NodeSet,omitempty" name:"NodeSet"`
		// 实例类型筛选列表

		InstanceTypeList []*string `json:"InstanceTypeList,omitempty" name:"InstanceTypeList"`
		// 可用区筛选列表

		AvailableZoneList []*string `json:"AvailableZoneList,omitempty" name:"AvailableZoneList"`
		// 版本筛选列表

		VersionList []*string `json:"VersionList,omitempty" name:"VersionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterNodesOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterNodesOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQAppIdStatsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRocketMQAppIdStatsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQAppIdStatsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmqTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateResourceRuleOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 备注信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则当前状态，0表示下线，1表示可用

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 协议名

	ProtocolName *string `json:"ProtocolName,omitempty" name:"ProtocolName"`
}

func (r *UpdateResourceRuleOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateResourceRuleOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TenantSetOpt struct {

	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 虚拟集群名

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// AppId

	CustomerAppId *uint64 `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// Uin

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 所属物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 命名空间个数

	UsedNamespaces *int64 `json:"UsedNamespaces,omitempty" name:"UsedNamespaces"`
	// 命名空间配额

	MaxNamespaces *int64 `json:"MaxNamespaces,omitempty" name:"MaxNamespaces"`
	// topic个数

	UsedTopics *int64 `json:"UsedTopics,omitempty" name:"UsedTopics"`
	// topic配额

	MaxTopics *int64 `json:"MaxTopics,omitempty" name:"MaxTopics"`
	// 存储空间配额，单位为byte

	MaxMsgBacklogSize *float64 `json:"MaxMsgBacklogSize,omitempty" name:"MaxMsgBacklogSize"`
	// 命名空间最大生产TPS

	MaxPublishTps *int64 `json:"MaxPublishTps,omitempty" name:"MaxPublishTps"`
	// 消息保留时间，单位为毫秒

	MaxRetention *int64 `json:"MaxRetention,omitempty" name:"MaxRetention"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 协议类型

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// Topic分区数配额

	MaxPartitions *int64 `json:"MaxPartitions,omitempty" name:"MaxPartitions"`
	// 已使用Topic分区数配额

	UsedPartitions *int64 `json:"UsedPartitions,omitempty" name:"UsedPartitions"`
	// 命名空间最大消费TPS

	MaxDispatchTps *int64 `json:"MaxDispatchTps,omitempty" name:"MaxDispatchTps"`
	// 监控指标

	Metrics *CommonMetricsOpt `json:"Metrics,omitempty" name:"Metrics"`
}

type DeleteSubscriptionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功删除的订阅关系数组。

		SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitempty" name:"SubscriptionTopicSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubscriptionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAMQPClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAMQPClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyBacklogQuotaOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 消息堆积大小限制，单位为bytes

	BacklogQuotaLimit *int64 `json:"BacklogQuotaLimit,omitempty" name:"BacklogQuotaLimit"`
	// 消息堆积限制策略，使用下拉框取值：
	// producer_request_hold、producer_exception、consumer_backlog_eviction

	BacklogQuotaPolicy *string `json:"BacklogQuotaPolicy,omitempty" name:"BacklogQuotaPolicy"`
}

func (r *UpdatePolicyBacklogQuotaOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyBacklogQuotaOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyCompactionThresholdOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 启用压缩阈值，单位为bytes

	CompactionThreshold *int64 `json:"CompactionThreshold,omitempty" name:"CompactionThreshold"`
}

func (r *UpdatePolicyCompactionThresholdOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyCompactionThresholdOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQPublicAccessPointRequest struct {
	*tchttp.BaseRequest

	// 集群ID，当前只支持专享集群

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeRocketMQPublicAccessPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQPublicAccessPointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterIdRes struct {

	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type OwnedNamespaceSetOpt struct {

	// 虚拟集群（租户）名

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 命名空间bundle范围

	BundleBoundary *string `json:"BundleBoundary,omitempty" name:"BundleBoundary"`
}

type DescribeInternalTenantsRequest struct {
	*tchttp.BaseRequest

	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 客户APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 虚拟集群名称

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// 协议类型数组

	Types []*string `json:"Types,omitempty" name:"Types"`
	// 排序字段名，支持createTime，updateTime

	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
	// 升序排列ASC，降序排列DESC

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeInternalTenantsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalTenantsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQTopicConfig struct {

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 主题类型：
	// Normal，普通
	// GlobalOrder， 全局顺序
	// PartitionedOrder, 分区顺序
	// Transaction，事务消息
	// DelayScheduled，延迟/定时消息

	Type *string `json:"Type,omitempty" name:"Type"`
	// 分区个数

	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DeleteAMQPRouteRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPRouteRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPRouteRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间列表

		Namespaces []*RocketMQNamespace `json:"Namespaces,omitempty" name:"Namespaces"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQClusterDetail struct {

	// 集群基本信息

	Info *RocketMQClusterInfo `json:"Info,omitempty" name:"Info"`
	// 集群配置信息

	Config *RocketMQClusterConfig `json:"Config,omitempty" name:"Config"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 集群状态，0:创建中，1:正常，2:销毁中，3:已删除，4: 隔离中，5:创建失败，6: 删除失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DeleteEnvironmentRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 角色名称数组。

		RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEnvironmentRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePortalConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 管理端页面配置信息列表

		Configs []*PortalConfig `json:"Configs,omitempty" name:"Configs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePortalConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePortalConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceRuleSetOpt struct {

	// 分配规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 指定用户的APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 指定用户的UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户名称

	CustomerName *string `json:"CustomerName,omitempty" name:"CustomerName"`
	// 是否大客户

	IsVipCustom *bool `json:"IsVipCustom,omitempty" name:"IsVipCustom"`
	// 协议名称

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 规则创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 规则更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 当前状态，0表示下线，1表示可用，2表示灰度中

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeClusterLatestMetricsOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 指标名：
	// topic_count：topic总数；msg_backlog：消息堆积数；
	// qps：qps；cpu_load：cpu负责；disk_load：存储负载

	MetricsName *string `json:"MetricsName,omitempty" name:"MetricsName"`
}

func (r *DescribeClusterLatestMetricsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterLatestMetricsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateItemsOptRequest struct {
	*tchttp.BaseRequest

	// 配置模块名称（小写）

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 0-S5 1-SA

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
	// 配置项

	TemplateItems []*ConfigTemplateItemOpt `json:"TemplateItems,omitempty" name:"TemplateItems"`
}

func (r *ModifyTemplateItemsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTemplateItemsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmqCluster struct {

	// 队列id

	QueueId *uint64 `json:"QueueId,omitempty" name:"QueueId"`
	// 唯一id

	SId *string `json:"SId,omitempty" name:"SId"`
	// app id

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 队列名

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 类型

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type DescribeClusterExtraResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群版本号

		ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterExtraResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主题详情

		TopicDescribe *CmqTopic `json:"TopicDescribe,omitempty" name:"TopicDescribe"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqTopicDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqTopicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名模糊匹配。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// topic类型描述：
	// 0：普通消息；
	// 1：全局顺序消息；
	// 2：局部顺序消息；
	// 3：重试队列；
	// 4：死信队列；
	// 5：事务消息。

	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// * TopicName
	// 按照主题名字查询，精确查询。
	// 类型：String
	// 必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 内部业务标识

	InternalReqSource *string `json:"InternalReqSource,omitempty" name:"InternalReqSource"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQVipInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		ClusterInfo *RocketMQClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`
		// 集群配置

		InstanceConfig *RocketMQInstanceConfig `json:"InstanceConfig,omitempty" name:"InstanceConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQVipInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQVipInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *DescribeClusterDetailOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeReplicationClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 目标集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 目标集群列表

		ClusterSets []*ReplicationClusterSets `json:"ClusterSets,omitempty" name:"ClusterSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReplicationClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeReplicationClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicySubscribeRateOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 消费速率限制

	DispatchRate *int64 `json:"DispatchRate,omitempty" name:"DispatchRate"`
	// 消费限流时间窗口，单位为秒

	DispatchTimePeriod *int64 `json:"DispatchTimePeriod,omitempty" name:"DispatchTimePeriod"`
}

func (r *UpdatePolicySubscribeRateOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicySubscribeRateOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckBatchRoleRequest struct {
	*tchttp.BaseRequest

	// 指定集群id校验

	ClusterScope *string `json:"ClusterScope,omitempty" name:"ClusterScope"`
	// 批量绑定数据

	BatchRoleBindData []*RoleBindData `json:"BatchRoleBindData,omitempty" name:"BatchRoleBindData"`
}

func (r *CheckBatchRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBatchRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅者名称，不支持中字以及除了短线和下划线外的特殊字符且不超过150个字符。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 备注，128个字符以内。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否幂等创建，若否不允许创建同名的订阅关系。

	IsIdempotent *bool `json:"IsIdempotent," name:"IsIdempotent"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 是否自动创建死信和重试主题，True 表示创建，False表示不创建，默认自动创建死信和重试主题。

	AutoCreatePolicyTopic *bool `json:"AutoCreatePolicyTopic,omitempty" name:"AutoCreatePolicyTopic"`
	// 指定死信和重试主题名称规范，LEGACY表示历史命名规则，COMMUNITY表示Pulsar社区命名规范

	PostFixPattern *string `json:"PostFixPattern,omitempty" name:"PostFixPattern"`
}

func (r *CreateSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {

	// 集群Id。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称。

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 说明信息。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 接入点数量

	EndPointNum *int64 `json:"EndPointNum,omitempty" name:"EndPointNum"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 集群是否健康，1表示健康，0表示异常

	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`
	// 集群健康信息

	HealthyInfo *string `json:"HealthyInfo,omitempty" name:"HealthyInfo"`
	// 集群状态，0:创建中，1:正常，2:删除中，3:已删除，5:创建失败，6: 删除失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 最大命名空间数量

	MaxNamespaceNum *int64 `json:"MaxNamespaceNum,omitempty" name:"MaxNamespaceNum"`
	// 最大Topic数量

	MaxTopicNum *int64 `json:"MaxTopicNum,omitempty" name:"MaxTopicNum"`
	// 最大QPS

	MaxQps *int64 `json:"MaxQps,omitempty" name:"MaxQps"`
	// 消息保留时间

	MessageRetentionTime *int64 `json:"MessageRetentionTime,omitempty" name:"MessageRetentionTime"`
	// 最大存储容量

	MaxStorageCapacity *int64 `json:"MaxStorageCapacity,omitempty" name:"MaxStorageCapacity"`
	// 集群版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 公网访问接入点

	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`
	// VPC访问接入点

	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`
	// 命名空间数量

	NamespaceNum *int64 `json:"NamespaceNum,omitempty" name:"NamespaceNum"`
	// 已使用存储限制，MB为单位

	UsedStorageBudget *int64 `json:"UsedStorageBudget,omitempty" name:"UsedStorageBudget"`
	// 资源对应项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 资源对应项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 计费模型0:按量计费，1：包年包月

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 最大生产消息速率，以条数为单位 注意：此字段可能返回 null，表示取不到有效值。

	MaxPublishRateInMessages *int64 `json:"MaxPublishRateInMessages,omitempty" name:"MaxPublishRateInMessages"`
	// 集群名字

	PclusterName *string `json:"PclusterName,omitempty" name:"PclusterName"`
	// 是否开启公网访问，不填时默认开启 注意：此字段可能返回 null，表示取不到有效值。

	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" name:"PublicAccessEnabled"`
	// 命名空间最大消费带宽，byte为单位 注意：此字段可能返回 null，表示取不到有效值。

	MaxDispatchRateInBytes *int64 `json:"MaxDispatchRateInBytes,omitempty" name:"MaxDispatchRateInBytes"`
	// 标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// topic数量

	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`
	// 内部接入点

	InternalEndPoint *string `json:"InternalEndPoint,omitempty" name:"InternalEndPoint"`
	// 最大推送消息速率，以条数为单位 注意：此字段可能返回 null，表示取不到有效值。

	MaxDispatchRateInMessages *int64 `json:"MaxDispatchRateInMessages,omitempty" name:"MaxDispatchRateInMessages"`
	// 内部pulsar-sdk接入地址

	InternalPulsarEndPoint *string `json:"InternalPulsarEndPoint,omitempty" name:"InternalPulsarEndPoint"`
	// 最大生产消息速率，以字节为单位 注意：此字段可能返回 null，表示取不到有效值。

	MaxPublishRateInBytes *int64 `json:"MaxPublishRateInBytes,omitempty" name:"MaxPublishRateInBytes"`
	// 最长消息延时，以秒为单位 注意：此字段可能返回 null，表示取不到有效值。

	MaxMessageDelayInSeconds *int64 `json:"MaxMessageDelayInSeconds,omitempty" name:"MaxMessageDelayInSeconds"`
	// 内部http接入地址

	InternalHttpEndPoint *string `json:"InternalHttpEndPoint,omitempty" name:"InternalHttpEndPoint"`
}

type DescribeBatchRoleRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeBatchRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccessEndpointSetOpt struct {

	// 协议名

	ProtocolName *string `json:"ProtocolName,omitempty" name:"ProtocolName"`
	// 公网接入点url

	PublicNetwork *string `json:"PublicNetwork,omitempty" name:"PublicNetwork"`
	// vpc接入点url

	VpcNetwork *string `json:"VpcNetwork,omitempty" name:"VpcNetwork"`
	// 内网接入点url

	InternalNetwork *string `json:"InternalNetwork,omitempty" name:"InternalNetwork"`
	// 公网接入点ip端口

	PublicNetworkIp *string `json:"PublicNetworkIp,omitempty" name:"PublicNetworkIp"`
	// VPC接入点ip端口

	VpcNetworkIp *string `json:"VpcNetworkIp,omitempty" name:"VpcNetworkIp"`
	// 内网接入点ip端口

	InternalNetworkIp *string `json:"InternalNetworkIp,omitempty" name:"InternalNetworkIp"`
}

type AddPulsarResourceAllocationRuleRequest struct {
	*tchttp.BaseRequest

	// 集群分配的规则ID

	RuleId *string `json:"RuleId,omitempty" name:"RuleId"`
	// 指定用户的APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 指定用户的UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 集群协议类型，支持的值为PULSAR，ROCKETMQ，AMQP，CMQ

	Type *string `json:"Type,omitempty" name:"Type"`
	// 集群当前状态，0表示下线，1表示可用，2表示灰度中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 备注信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// Pulsar软件版本

	PulsarVersion *string `json:"PulsarVersion,omitempty" name:"PulsarVersion"`
	// Pulsar管理端接口地址

	AdminUrl *string `json:"AdminUrl,omitempty" name:"AdminUrl"`
	// Pulsar管理端接口token

	AdminToken *string `json:"AdminToken,omitempty" name:"AdminToken"`
	// 协议端管理接口地址

	ExternalAdminUrl *string `json:"ExternalAdminUrl,omitempty" name:"ExternalAdminUrl"`
	// 协议端管理接口token

	ExternalAdminToken *string `json:"ExternalAdminToken,omitempty" name:"ExternalAdminToken"`
	// 接入点协议

	EndPointProtocol *string `json:"EndPointProtocol,omitempty" name:"EndPointProtocol"`
	// VPC接入点地址通用后缀

	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`
	// 公网接入点地址通用后缀

	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`
}

func (r *AddPulsarResourceAllocationRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPulsarResourceAllocationRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBrokerOwnedNamespacesOptRequest struct {
	*tchttp.BaseRequest

	// 节点实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 虚拟集群（租户）ID，支持模糊匹配

	TenantIdSearch *string `json:"TenantIdSearch,omitempty" name:"TenantIdSearch"`
	// 命名空间名，支持模糊匹配

	NamespaceNameSearch *string `json:"NamespaceNameSearch,omitempty" name:"NamespaceNameSearch"`
}

func (r *DescribeBrokerOwnedNamespacesOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBrokerOwnedNamespacesOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePulsarResourceAllocationRulesRequest struct {
	*tchttp.BaseRequest

	// 如填写，则按照用户APPID进行精确查询

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 如填写，则按照协议类型进行查询

	Types []*string `json:"Types,omitempty" name:"Types"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制条目数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePulsarResourceAllocationRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePulsarResourceAllocationRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyDispatchRatePerSubscriptionOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 消费流量限制，单位为bytes

	DispatchThroughput *int64 `json:"DispatchThroughput,omitempty" name:"DispatchThroughput"`
	// 消费速率限制

	DispatchRate *int64 `json:"DispatchRate,omitempty" name:"DispatchRate"`
	// 消费限流时间窗口

	DispatchTimePeriod *int64 `json:"DispatchTimePeriod,omitempty" name:"DispatchTimePeriod"`
}

func (r *UpdatePolicyDispatchRatePerSubscriptionOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyDispatchRatePerSubscriptionOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvironmentRoleSet struct {

	// 命名空间Id

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 权限枚举值

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
	// 绑定的集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 角色名字

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
}

type CreateRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称，3-64个字符，只能包含字母、数字、“-”及“_”

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群描述，128个字符以内

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateRocketMQClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicReplicationClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定的目标集群总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 绑定的目标集群列表

		ClusterSets []*ReplicationClusterSets `json:"ClusterSets,omitempty" name:"ClusterSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicReplicationClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicReplicationClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BundleSetOpt struct {

	// bundle名称

	BundleName *string `json:"BundleName,omitempty" name:"BundleName"`
	// 短期负载指标

	ShortTermMetrics *CommonMetricsOpt `json:"ShortTermMetrics,omitempty" name:"ShortTermMetrics"`
	// 长期负载指标

	LongTermMetrics *CommonMetricsOpt `json:"LongTermMetrics,omitempty" name:"LongTermMetrics"`
	// 短期负载采样最大值

	ShortTermSamplesMax *int64 `json:"ShortTermSamplesMax,omitempty" name:"ShortTermSamplesMax"`
	// 短期负载采样实际值

	ShortTermSamplesReal *int64 `json:"ShortTermSamplesReal,omitempty" name:"ShortTermSamplesReal"`
	// 长期负载采样最大值

	LongTermSamplesMax *int64 `json:"LongTermSamplesMax,omitempty" name:"LongTermSamplesMax"`
	// 长期负载采样实际值

	LongTermSamplesReal *int64 `json:"LongTermSamplesReal,omitempty" name:"LongTermSamplesReal"`
}

type PutClusterOfflineOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutClusterOfflineOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutClusterOfflineOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQMigrationTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 迁移任务列表

		Tasks []*RocketMQMigrationTask `json:"Tasks,omitempty" name:"Tasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQMigrationTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQMigrationTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterSubscription struct {

	// 是否仅展示包含真实消费者的订阅。

	ConsumerHasCount *bool `json:"ConsumerHasCount,omitempty" name:"ConsumerHasCount"`
	// 是否仅展示消息堆积的订阅。

	ConsumerHasBacklog *bool `json:"ConsumerHasBacklog,omitempty" name:"ConsumerHasBacklog"`
	// 是否仅展示存在消息超期丢弃的订阅。

	ConsumerHasExpired *bool `json:"ConsumerHasExpired,omitempty" name:"ConsumerHasExpired"`
	// 按照订阅名过滤，精确查询。

	SubscriptionNames []*string `json:"SubscriptionNames,omitempty" name:"SubscriptionNames"`
}

type DescribeAMQPClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeAMQPClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPExchangeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost间名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 交换机名称

	Exchange *string `json:"Exchange,omitempty" name:"Exchange"`
	// 说明信息，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAMQPExchangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAMQPExchangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQBillingUsageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当月API调用量

		APIUsage *uint64 `json:"APIUsage,omitempty" name:"APIUsage"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQBillingUsageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQBillingUsageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQEnvironmentRolesRequest struct {
	*tchttp.BaseRequest

	// 必填字段，环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 起始下标，不填默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 必填字段，Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 角色名称

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQEnvironmentRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQEnvironmentRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearCmqQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *ClearCmqQueueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearCmqQueueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTenantQuotaOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateTenantQuotaOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTenantQuotaOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcknowledgeMessageRequest struct {
	*tchttp.BaseRequest

	// 用作标识消息的唯一的ID（可从 receiveMessage 的返回值中获得）

	MessageId *string `json:"MessageId,omitempty" name:"MessageId"`
	// Topic 名字（可从 receiveMessage 的返回值中获得）这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default

	AckTopic *string `json:"AckTopic,omitempty" name:"AckTopic"`
	// 订阅者的名字，可以从receiveMessage的返回值中获取到。这里尽量与receiveMessage中的订阅者保持一致，否则没办法正确ack 接收回来的消息。

	SubName *string `json:"SubName,omitempty" name:"SubName"`
}

func (r *AcknowledgeMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcknowledgeMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签的key的值

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签的Value的值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeRocketMQTopUsagesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 指标名称，支持以下：
	// consumeLag，消费组堆积数量
	// deadLetterCount，死信数量
	// topicRateIn,   Topic生产速率
	// topicRateOut，Topic消费速率
	// topicStorageSize，Topic存储空间
	// topicApiCalls，Topic API调用次数

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 排序数量，最大20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRocketMQTopUsagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQTopUsagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Producer struct {

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 连接数。

	CountConnect *int64 `json:"CountConnect,omitempty" name:"CountConnect"`
	// 连接集合。

	ConnectionSets []*Connection `json:"ConnectionSets,omitempty" name:"ConnectionSets"`
}

type UpdatePolicySchemaCompatibilityStrategyOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// Schema自动更新策略，下拉框取值：
	// UNDEFINED、ALWAYS_INCOMPATIBLE、ALWAYS_COMPATIBLE、BACKWARD、FORWARD、FULL、BACKWARD_TRANSITIVE、FORWARD_TRANSITIVE、FULL_TRANSITIVE

	SchemaCompatibilityStrategy *string `json:"SchemaCompatibilityStrategy,omitempty" name:"SchemaCompatibilityStrategy"`
}

func (r *UpdatePolicySchemaCompatibilityStrategyOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicySchemaCompatibilityStrategyOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQVipInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRocketMQVipInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQVipInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主题列表

		TopicList []*CmqTopic `json:"TopicList,omitempty" name:"TopicList"`
		// 全量主题数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 未消费消息过期时间，单位：秒，最大1296000。

	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`
	// 备注，字符串最长不超过128。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 消息保留策略

	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

func (r *ModifyEnvironmentAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEnvironmentAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyBacklogQuotaOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyBacklogQuotaOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyBacklogQuotaOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQEnvironmentRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命名空间角色集合。

		EnvironmentRoleSets []*EnvironmentRole `json:"EnvironmentRoleSets,omitempty" name:"EnvironmentRoleSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQEnvironmentRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPRouteRelationsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 路由关系列表

		RouteRelations []*AMQPRouteRelation `json:"RouteRelations,omitempty" name:"RouteRelations"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPRouteRelationsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPRouteRelationsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命名空间角色集合。

		EnvironmentRoleSets []*EnvironmentRole `json:"EnvironmentRoleSets,omitempty" name:"EnvironmentRoleSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnvironmentRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyCompactionThresholdOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyCompactionThresholdOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyCompactionThresholdOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeliveryInfoOpt struct {

	// 代理节点IP

	DeliveryIp *string `json:"DeliveryIp,omitempty" name:"DeliveryIp"`
	// 下发时间

	DeliveryTimes *int64 `json:"DeliveryTimes,omitempty" name:"DeliveryTimes"`
	// 下发状态（实际大小1字节） 0 下发中 1 陈工 2失败

	DeliveryStatus *int64 `json:"DeliveryStatus,omitempty" name:"DeliveryStatus"`
	// 版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 新增时间

	AddTime *int64 `json:"AddTime,omitempty" name:"AddTime"`
}

type SendMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 消息ID。

		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateResourceRuleOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateResourceRuleOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceRuleOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClearNamespaceBacklogOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearNamespaceBacklogOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearNamespaceBacklogOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBindVpcRequest struct {
	*tchttp.BaseRequest

	// 租户Vpc Id。

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// Vpc子网Id。

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteBindVpcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBindVpcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscriptionStatsSetOpt struct {

	// 消息消费速率

	MsgRateOut *float64 `json:"MsgRateOut,omitempty" name:"MsgRateOut"`
	// 消息消费吞吐量

	MsgThroughputOut *float64 `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`
	// 读出字节总数

	BytesOutCounter *int64 `json:"BytesOutCounter,omitempty" name:"BytesOutCounter"`
	// 消息消费总数

	MsgOutCounter *int64 `json:"MsgOutCounter,omitempty" name:"MsgOutCounter"`
	// 消息重传速率

	MsgRateRedeliver *float64 `json:"MsgRateRedeliver,omitempty" name:"MsgRateRedeliver"`
	// 分块消息消费速率

	ChunkedMessageRate *float64 `json:"ChunkedMessageRate,omitempty" name:"ChunkedMessageRate"`
	// 积压消息数

	MsgBacklog *int64 `json:"MsgBacklog,omitempty" name:"MsgBacklog"`
	// 非延迟消息积压消息数

	MsgBacklogNoDelayed *int64 `json:"MsgBacklogNoDelayed,omitempty" name:"MsgBacklogNoDelayed"`
	// 待确认：BlockedSubscriptionOnUnackedMsgs

	BlockedSubscriptionOnUnackedMsgs *bool `json:"BlockedSubscriptionOnUnackedMsgs,omitempty" name:"BlockedSubscriptionOnUnackedMsgs"`
	// 最近一次确认消息时间

	LastAckedTimestamp *int64 `json:"LastAckedTimestamp,omitempty" name:"LastAckedTimestamp"`
	// 最近一次消费时间

	LastConsumedTimestamp *int64 `json:"LastConsumedTimestamp,omitempty" name:"LastConsumedTimestamp"`
	// 延迟消息数

	MsgDelayed *int64 `json:"MsgDelayed,omitempty" name:"MsgDelayed"`
	// 未确认消息数

	UnackedMessages *int64 `json:"UnackedMessages,omitempty" name:"UnackedMessages"`
	// 消息过期速率

	MsgRateExpired *float64 `json:"MsgRateExpired,omitempty" name:"MsgRateExpired"`
	// 最近一次消息过期时间

	LastExpireTimestamp *int64 `json:"LastExpireTimestamp,omitempty" name:"LastExpireTimestamp"`
	// 待确认：LastConsumedFlowTimestamp

	LastConsumedFlowTimestamp *int64 `json:"LastConsumedFlowTimestamp,omitempty" name:"LastConsumedFlowTimestamp"`
	// 是否持久化

	IsDurable *bool `json:"IsDurable,omitempty" name:"IsDurable"`
	// 是否跨地域复制

	IsReplicated *bool `json:"IsReplicated,omitempty" name:"IsReplicated"`
	// 消费者状态数组

	Consumers []*ConsumerStatsSetOpt `json:"Consumers,omitempty" name:"Consumers"`
	// 订阅名

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

type ResetMsgSubOffsetByTimestampResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果。

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetMsgSubOffsetByTimestampResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetMsgSubOffsetByTimestampResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQSubscriptionsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 消费组名称

	Group *string `json:"Group,omitempty" name:"Group"`
	// 查询起始位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRocketMQSubscriptionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQSubscriptionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户Vpc Id。

		UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
		// Vpc子网Id。

		UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentNodesOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 是否在线

	IsOnline *bool `json:"IsOnline,omitempty" name:"IsOnline"`
}

func (r *DescribeAgentNodesOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentNodesOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigTemplateItemOpt struct {

	// 配置属性

	ItemKey *string `json:"ItemKey,omitempty" name:"ItemKey"`
	// 属性值

	ItemValue *string `json:"ItemValue,omitempty" name:"ItemValue"`
	// 配置项备注

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 是否动态配置

	IsDynamic *bool `json:"IsDynamic,omitempty" name:"IsDynamic"`
}

type DeleteRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// 待删除的集群Id。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteRocketMQClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色名称

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 备注说明

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 过滤参数的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数值

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type CreateResourceRuleOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 协议名称

	ProtocolName *string `json:"ProtocolName,omitempty" name:"ProtocolName"`
	// 指定用户的APPID；设置公共集群时为空

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 指定用户的UIN；设置公共集群时为空

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 备注信息

	Comment *string `json:"Comment,omitempty" name:"Comment"`
	// 协议端Admin Url

	ExternalAdminUrl *string `json:"ExternalAdminUrl,omitempty" name:"ExternalAdminUrl"`
	// 协议端Admin Token

	ExternalAdminToken *string `json:"ExternalAdminToken,omitempty" name:"ExternalAdminToken"`
}

func (r *CreateResourceRuleOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateResourceRuleOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// Group名称，8~64个字符

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 说明信息，最长128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 命名空间，目前只支持单个命名空间

	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
	// 是否开启消费

	ReadEnable *bool `json:"ReadEnable,omitempty" name:"ReadEnable"`
	// 是否开启广播消费

	BroadcastEnable *bool `json:"BroadcastEnable,omitempty" name:"BroadcastEnable"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Group类型（TCP/HTTP）

	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`
	// Group最大重试次数

	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitempty" name:"RetryMaxTimes"`
}

func (r *CreateRocketMQGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterMetaInitScriptOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 初始化集群脚本

		InitMetaScript *string `json:"InitMetaScript,omitempty" name:"InitMetaScript"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterMetaInitScriptOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterMetaInitScriptOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgTraceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 生产信息。

		ProducerLog *ProducerLog `json:"ProducerLog,omitempty" name:"ProducerLog"`
		// 服务方信息。

		ServerLog *ServerLog `json:"ServerLog,omitempty" name:"ServerLog"`
		// 消费信息。

		ConsumerLogs *ConsumerLogs `json:"ConsumerLogs,omitempty" name:"ConsumerLogs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsgTraceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMsgTraceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPVHost struct {

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 未消费消息的保留时间，以毫秒单位，范围60秒到15天

	MsgTtl *uint64 `json:"MsgTtl,omitempty" name:"MsgTtl"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

type ClearNamespaceBacklogOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// Broker实例ID，与物理集群名二选一

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ClearNamespaceBacklogOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearNamespaceBacklogOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群信息数组

		ClusterSets []*ClusterSetOpt `json:"ClusterSets,omitempty" name:"ClusterSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmqTopic struct {

	// 主题的 ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息在主题中最长存活时间，从发送到该主题开始经过此参数指定的时间后，不论消息是否被成功推送给用户都将被删除，单位为秒。固定为一天（86400秒），该属性不能修改。

	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`
	// 消息最大长度。取值范围1024 - 1048576Byte（即1 - 1024K），默认值为65536。

	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`
	// 每秒钟发布消息的条数。

	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`
	// 描述用户创建订阅时选择的过滤策略：
	// FilterType = 1表示用户使用 FilterTag 标签过滤;
	// FilterType = 2表示用户使用 BindingKey 过滤。

	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`
	// 主题的创建时间。返回 Unix 时间戳，精确到毫秒。

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后一次修改主题属性的时间。返回 Unix 时间戳，精确到毫秒。

	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`
	// 当前该主题中消息数目（消息堆积数）。

	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`
	// 创建者 Uin，CAM 鉴权 resource 由该字段组合而成。

	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`
	// 关联的标签。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 消息轨迹。true表示开启，false表示不开启。

	Trace *bool `json:"Trace,omitempty" name:"Trace"`
	// 租户id

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type CmqSubscription struct {

	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 订阅 ID。订阅 ID 在拉取监控数据时会用到。

	SubscriptionId *string `json:"SubscriptionId,omitempty" name:"SubscriptionId"`
	// 订阅拥有者的 APPID。

	TopicOwner *uint64 `json:"TopicOwner,omitempty" name:"TopicOwner"`
	// 该订阅待投递的消息数。

	MsgCount *uint64 `json:"MsgCount,omitempty" name:"MsgCount"`
	// 最后一次修改订阅属性的时间。返回 Unix 时间戳，精确到毫秒。

	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`
	// 订阅的创建时间。返回 Unix 时间戳，精确到毫秒。

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 表示订阅接收消息的过滤策略。

	BindingKey []*string `json:"BindingKey,omitempty" name:"BindingKey"`
	// 接收通知的 endpoint，根据协议 protocol 区分：对于 HTTP，endpoint 必须以http://开头，host 可以是域名或 IP；对于 queue，则填 queueName。

	Endpoint *string `json:"Endpoint,omitempty" name:"Endpoint"`
	// 描述用户创建订阅时选择的过滤策略：
	// filterType = 1表示用户使用 filterTag 标签过滤
	// filterType = 2表示用户使用 bindingKey 过滤。

	FilterTags []*string `json:"FilterTags,omitempty" name:"FilterTags"`
	// 订阅的协议，目前支持两种协议：HTTP、queue。使用 HTTP 协议，用户需自己搭建接受消息的 Web Server。使用 queue，消息会自动推送到 CMQ queue，用户可以并发地拉取消息。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 向 endpoint 推送消息出现错误时，CMQ 推送服务器的重试策略。取值有：
	// （1）BACKOFF_RETRY，退避重试。每隔一定时间重试一次，重试够一定次数后，就把该消息丢弃，继续推送下一条消息；
	// （2）EXPONENTIAL_DECAY_RETRY，指数衰退重试。每次重试的间隔是指数递增的，例如开始 1s，后面是 2s，4s，8s...由于 Topic 消息的周期是一天，所以最多重试一天就把消息丢弃。默认值是 EXPONENTIAL_DECAY_RETRY。

	NotifyStrategy *string `json:"NotifyStrategy,omitempty" name:"NotifyStrategy"`
	// 推送内容的格式。取值：（1）JSON；（2）SIMPLIFIED，即 raw 格式。如果 protocol 是 queue，则取值必须为 SIMPLIFIED。如果 protocol 是 HTTP，两个值均可以，默认值是 JSON。

	NotifyContentFormat *string `json:"NotifyContentFormat,omitempty" name:"NotifyContentFormat"`
}

type CreateCmqTopicRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线（-）。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息最大长度。取值范围 1024-65536 Byte（即1-64K），默认值 65536。

	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`
	// 用于指定主题的消息匹配策略。1：表示标签匹配策略；2：表示路由匹配策略，默认值为标签匹配策略。

	FilterType *uint64 `json:"FilterType,omitempty" name:"FilterType"`
	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。

	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`
	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。

	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *CreateCmqTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmqTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 角色名称。

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 授权项。

		Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvironmentRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQMigrationTaskDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子任务总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群id

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 任务状态：
		// 0，迁移中
		// 1，迁移成功
		// 2，迁移完成

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 任务类型，
		// 0，集群迁移
		// 1，指定命名空间

		Type *int64 `json:"Type,omitempty" name:"Type"`
		// 创建时间，毫秒为单位

		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
		// 子任务列表

		SubTasks []*RocketMQMigrationSubTask `json:"SubTasks,omitempty" name:"SubTasks"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQMigrationTaskDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQMigrationTaskDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalRocketMQTopicsRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 过滤条件，支持以下条件：
	// 1. name，topic名称模糊搜索
	// 2. type，按topic类型搜索

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询起始位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询结果最大数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInternalRocketMQTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalRocketMQTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePClustersRequest struct {
	*tchttp.BaseRequest

	// 0表示Pulsar，1表示rocketmq，2表示AMQP，3表示CMQ

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

func (r *DescribePClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReceiveMessageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用作标识消息的唯一主键

		MessageID *string `json:"MessageID,omitempty" name:"MessageID"`
		// 接收消息的内容

		MessagePayload *string `json:"MessagePayload,omitempty" name:"MessagePayload"`
		// 提供给 Ack 接口，用来Ack哪一个topic中的消息

		AckTopic *string `json:"AckTopic,omitempty" name:"AckTopic"`
		// 返回的错误信息，如果为空，说明没有错误

		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
		// 返回订阅者的名字，用来创建 ack consumer时使用

		SubName *string `json:"SubName,omitempty" name:"SubName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReceiveMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReceiveMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功删除的角色名称数组。

		RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPClusterInfo struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 地域信息

	Region *string `json:"Region,omitempty" name:"Region"`
	// 创建时间，毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 集群说明信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 公网接入地址

	PublicEndPoint *string `json:"PublicEndPoint,omitempty" name:"PublicEndPoint"`
	// VPC接入地址

	VpcEndPoint *string `json:"VpcEndPoint,omitempty" name:"VpcEndPoint"`
}

type CmqTransactionPolicy struct {

	// 第一次回查时间。

	FirstQueryInterval *uint64 `json:"FirstQueryInterval,omitempty" name:"FirstQueryInterval"`
	// 最大查询次数。

	MaxQueryCount *uint64 `json:"MaxQueryCount,omitempty" name:"MaxQueryCount"`
}

type DeleteAMQPExchangeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 交换机名称

	Exchange *string `json:"Exchange,omitempty" name:"Exchange"`
}

func (r *DeleteAMQPExchangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPExchangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterComponentMetricsListOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标列表

		MetricsSet []*ClusterComponentMetricsSetListOpt `json:"MetricsSet,omitempty" name:"MetricsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterComponentMetricsListOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterComponentMetricsListOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateOptRequest struct {
	*tchttp.BaseRequest

	// 配置模块名称（小写）

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 0-S5 1-SA

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
	// 脚本内容

	ShellScript *string `json:"ShellScript,omitempty" name:"ShellScript"`
}

func (r *ModifyTemplateOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTemplateOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllRocketMQTenantsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 虚拟集群列表

		Tenants []*InternalTenant `json:"Tenants,omitempty" name:"Tenants"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllRocketMQTenantsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllRocketMQTenantsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeLatestMetricsOptRequest struct {
	*tchttp.BaseRequest

	// 指标名：
	// rate_in：broker生产流量；rate_out：broker消费速率；throughput_in：broker入流量；throughput_in：broker出流量；
	// bookie_read：bookie读消息速率；bookie_write写消息速率；
	// entry_write：bookie写entry成功率；entry_read：bookie读entry成功率；

	MetricsName *string `json:"MetricsName,omitempty" name:"MetricsName"`
	// 节点ip

	NodeAddress *string `json:"NodeAddress,omitempty" name:"NodeAddress"`
}

func (r *DescribeNodeLatestMetricsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeLatestMetricsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 未消费消息过期时间，单位：秒。

		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`
		// 备注，字符串最长不超过128。

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 命名空间ID

		NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvironmentAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyEnvironmentAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		ClusterList []*AMQPClusterDetail `json:"ClusterList,omitempty" name:"ClusterList"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyPublishRateOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyPublishRateOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyPublishRateOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindCluster struct {

	// 物理集群的ID

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 物理集群的名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

type DescribeRocketMQNamespacesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 按名称搜索

	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`
}

func (r *DescribeRocketMQNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterConfigOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 配置版本号

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *DescribeClusterConfigOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterConfigOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacePolicyOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// topic级别消费限流配置

		DispatchRateOnTopic *NamespacePolicyRateLimitOpt `json:"DispatchRateOnTopic,omitempty" name:"DispatchRateOnTopic"`
		// 订阅级别消费限流配置

		DispatchRateOnSubscription *NamespacePolicyRateLimitOpt `json:"DispatchRateOnSubscription,omitempty" name:"DispatchRateOnSubscription"`
		// 消费者级别消费限流配置

		DispatchRateOnConsumer *NamespacePolicyRateLimitOpt `json:"DispatchRateOnConsumer,omitempty" name:"DispatchRateOnConsumer"`
		// 生产限流配置

		PublishRate *NamespacePolicyRateLimitOpt `json:"PublishRate,omitempty" name:"PublishRate"`
		// 消息存活时间，单位为秒

		MessageTTL *int64 `json:"MessageTTL,omitempty" name:"MessageTTL"`
		// 消息保留时间，单位为分钟

		RetentionTime *int64 `json:"RetentionTime,omitempty" name:"RetentionTime"`
		// 消息保留大小，单位为MB

		RetentionSize *int64 `json:"RetentionSize,omitempty" name:"RetentionSize"`
		// 消息堆积大小限制，单位为bytes

		BacklogQuotaLimit *int64 `json:"BacklogQuotaLimit,omitempty" name:"BacklogQuotaLimit"`
		// 消息堆积限制策略

		BacklogQuotaPolicy *string `json:"BacklogQuotaPolicy,omitempty" name:"BacklogQuotaPolicy"`
		// 单个 topic 使用的 bookie 数量

		PersistenceEnsemble *int64 `json:"PersistenceEnsemble,omitempty" name:"PersistenceEnsemble"`
		// 每个 entry 要写入的次数

		PersistenceWriteQuorum *int64 `json:"PersistenceWriteQuorum,omitempty" name:"PersistenceWriteQuorum"`
		// 每个 entry 在等待的 acks（有保证的副本）数量

		PersistenceAckQuorum *int64 `json:"PersistenceAckQuorum,omitempty" name:"PersistenceAckQuorum"`
		// 标记-删除操作的限制速率（0表示无限制）

		PersistenceMarkDeleteRate *float64 `json:"PersistenceMarkDeleteRate,omitempty" name:"PersistenceMarkDeleteRate"`
		// 单个topic生产者数量限制

		MaxProducerOnTopic *int64 `json:"MaxProducerOnTopic,omitempty" name:"MaxProducerOnTopic"`
		// 单个topic消费者数量限制

		MaxConsumerOnTopic *int64 `json:"MaxConsumerOnTopic,omitempty" name:"MaxConsumerOnTopic"`
		// 单个订阅消费者数量限制

		MaxConsumerOnSubscription *int64 `json:"MaxConsumerOnSubscription,omitempty" name:"MaxConsumerOnSubscription"`
		// schema自动更新策略

		SchemaCompatibilityStrategy *string `json:"SchemaCompatibilityStrategy,omitempty" name:"SchemaCompatibilityStrategy"`
		// schema开启强校验

		SchemaValidationEnforced *bool `json:"SchemaValidationEnforced,omitempty" name:"SchemaValidationEnforced"`
		// 订阅鉴权模式

		SubscriptionAuthMode *string `json:"SubscriptionAuthMode,omitempty" name:"SubscriptionAuthMode"`
		// 加密开关状态

		EncryptionRequired *bool `json:"EncryptionRequired,omitempty" name:"EncryptionRequired"`
		// 清除重复消息开关状态

		DeduplicationEnabled *bool `json:"DeduplicationEnabled,omitempty" name:"DeduplicationEnabled"`
		// 压缩阈值，单位为bytes

		CompactionThreshold *int64 `json:"CompactionThreshold,omitempty" name:"CompactionThreshold"`
		// 卸载阈值，单位为bytes

		OffloadThreshold *int64 `json:"OffloadThreshold,omitempty" name:"OffloadThreshold"`
		// bookie主亲和组

		AffinityGroupPrimary *string `json:"AffinityGroupPrimary,omitempty" name:"AffinityGroupPrimary"`
		// bookie次亲和组

		AffinityGroupSecondary *string `json:"AffinityGroupSecondary,omitempty" name:"AffinityGroupSecondary"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespacePolicyOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespacePolicyOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPExchangesRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost ID

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 按路由类型过滤查询结果，可选择Direct, Fanout, Topic

	FilterType []*string `json:"FilterType,omitempty" name:"FilterType"`
	// 按exchange名称搜索，支持模糊查询

	FilterName *string `json:"FilterName,omitempty" name:"FilterName"`
	// 过滤查询内部或者外部exchange

	FilterInternal *bool `json:"FilterInternal,omitempty" name:"FilterInternal"`
}

func (r *DescribeAMQPExchangesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPExchangesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRealTimeSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 订阅者名称，模糊匹配。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 数据过滤条件。

	Filters []*FilterSubscription `json:"Filters,omitempty" name:"Filters"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRealTimeSubscriptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRealTimeSubscriptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicMsgsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 消息ID。

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeTopicMsgsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicMsgsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQMigrationTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQMigrationTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQMigrationTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQMigrationTask struct {

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 任务类型：
	// 0，集群迁移
	// 1，导入指定的命名空间

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 主题总数

	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`
	// 消费组总数

	GroupNum *int64 `json:"GroupNum,omitempty" name:"GroupNum"`
	// 创建时间，单位毫秒

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务状态：
	// 0，迁移中
	// 1，迁移成功
	// 2，迁移完成，只有部分数据完成迁移

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 导入的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type DescribeBindClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 专享集群的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 专享集群的列表

		ClusterSet []*BindCluster `json:"ClusterSet,omitempty" name:"ClusterSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBindClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopicsByGroupRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 消费组名称

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRocketMQTopicsByGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQTopicsByGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicySchemaCompatibilityStrategyOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicySchemaCompatibilityStrategyOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicySchemaCompatibilityStrategyOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPClusterDetail struct {

	// 集群基本信息

	Info *AMQPClusterInfo `json:"Info,omitempty" name:"Info"`
	// 集群配置信息

	Config *AMQPClusterConfig `json:"Config,omitempty" name:"Config"`
}

type Connection struct {

	// 生产者地址。

	Address *string `json:"Address,omitempty" name:"Address"`
	// 主题分区。

	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`
	// 生产者版本。

	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`
	// 生产者名称。

	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`
	// 生产者ID。

	ProducerId *string `json:"ProducerId,omitempty" name:"ProducerId"`
	// 消息平均大小(byte)。

	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`
	// 生成速率(byte/秒)。

	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`
}

type PartitionsTopic struct {

	// 最后一次间隔内发布消息的平均byte大小。

	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`
	// 消费者数量。

	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`
	// 被记录下来的消息总数。

	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitempty" name:"LastConfirmedEntry"`
	// 最后一个ledger创建的时间。

	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitempty" name:"LastLedgerCreatedTimestamp"`
	// 本地和复制的发布者每秒发布消息的速率。

	MsgRateIn *string `json:"MsgRateIn,omitempty" name:"MsgRateIn"`
	// 本地和复制的消费者每秒分发消息的数量之和。

	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`
	// 本地和复制的发布者每秒发布消息的byte。

	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`
	// 本地和复制的消费者每秒分发消息的byte。

	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`
	// 被记录下来的消息总数。

	NumberOfEntries *string `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`
	// 子分区id。

	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`
	// 生产者数量。

	ProducerCount *string `json:"ProducerCount,omitempty" name:"ProducerCount"`
	// 以byte计算的所有消息存储总量。

	TotalSize *string `json:"TotalSize,omitempty" name:"TotalSize"`
	// topic类型描述。

	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`
}

type DeleteCmqQueueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCmqQueueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyDispatchRatePerTopicOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyDispatchRatePerTopicOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyDispatchRatePerTopicOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMessageTTLOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 消息存活时间，单位为秒

	MessageTTL *int64 `json:"MessageTTL,omitempty" name:"MessageTTL"`
}

func (r *UpdatePolicyMessageTTLOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMessageTTLOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTenantTopListOptRequest struct {
	*tchttp.BaseRequest

	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 取值为：
	// 生产：rate_in；消费：rate_out；入流量：throughput_in；出流量：throughput_out；存储大小：storage_size

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 取N条Top列表值

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

func (r *GetTenantTopListOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTenantTopListOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQBillingUsageRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRocketMQBillingUsageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQBillingUsageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteRequest struct {
	*tchttp.BaseRequest

	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendConfigInfoOpt struct {

	// 节点ip地址

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
	// 下发状态

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type DescribeProducersRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 生产者名称，模糊匹配。

	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeProducersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProducersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigTypeInfoOpt struct {

	// 模块组件名称

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 模块配置文件名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 配置模块名

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
}

type Consumer struct {

	// 消费者开始连接的时间。

	ConnectedSince *string `json:"ConnectedSince,omitempty" name:"ConnectedSince"`
	// 消费者地址。

	ConsumerAddr *string `json:"ConsumerAddr,omitempty" name:"ConsumerAddr"`
	// 消费者名称。

	ConsumerName *string `json:"ConsumerName,omitempty" name:"ConsumerName"`
	// 消费者版本。

	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`
	// Partition数量

	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// 集群名称，不支持中字以及除了短线和下划线外的特殊字符且不超过16个字符。

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 用户专享物理集群ID，如果不传，则默认在公共集群上创建用户集群资源。

	BindClusterId *uint64 `json:"BindClusterId,omitempty" name:"BindClusterId"`
	// 说明，128个字符以内。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 集群的标签列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 集群对应的项目id

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 绑定集群名

	BindClusterName *string `json:"BindClusterName,omitempty" name:"BindClusterName"`
}

func (r *CreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQNamespaceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQNamespaceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQNamespaceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 订阅者名称，模糊匹配。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 数据过滤条件。

	Filters []*FilterSubscription `json:"Filters,omitempty" name:"Filters"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeSubscriptionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllRocketMQTenantsRequest struct {
	*tchttp.BaseRequest

	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 虚拟集群名称

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// 协议类型数组

	Types []*string `json:"Types,omitempty" name:"Types"`
	// 排序字段名，支持createTime，updateTime

	SortBy *string `json:"SortBy,omitempty" name:"SortBy"`
	// 升序排列ASC，降序排列DESC

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeAllRocketMQTenantsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllRocketMQTenantsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterOptRequest struct {
	*tchttp.BaseRequest

	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群底层资源

	ClusterResource *string `json:"ClusterResource,omitempty" name:"ClusterResource"`
	// 协议名数组

	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`
	// data-reporter上报地址

	DataReporterServer *string `json:"DataReporterServer,omitempty" name:"DataReporterServer"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *UpdateClusterOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPRouteRelation struct {

	// 路由关系ID

	RouteRelationId *string `json:"RouteRelationId,omitempty" name:"RouteRelationId"`
	// 源Exchange

	SourceExchange *string `json:"SourceExchange,omitempty" name:"SourceExchange"`
	// 目标类型:Queue|Exchange

	DestType *string `json:"DestType,omitempty" name:"DestType"`
	// 目标值

	DestValue *string `json:"DestValue,omitempty" name:"DestValue"`
	// 绑定key

	RoutingKey *string `json:"RoutingKey,omitempty" name:"RoutingKey"`
	// 源路由类型:Direct|Topic|Fanout

	SourceExchangeType *string `json:"SourceExchangeType,omitempty" name:"SourceExchangeType"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 说明信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type CreateAMQPRouteRelationRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 交换机所在的vhost

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 源Exchange名称

	SourceExchange *string `json:"SourceExchange,omitempty" name:"SourceExchange"`
	// 交换机说明，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 目标类型:Queue|Exchange

	DestType *string `json:"DestType,omitempty" name:"DestType"`
	// 目标值

	DestValue *string `json:"DestValue,omitempty" name:"DestValue"`
	// 绑定key,缺省值为default

	RoutingKey *string `json:"RoutingKey,omitempty" name:"RoutingKey"`
}

func (r *CreateAMQPRouteRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPRouteRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCmqQueueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建成功的queueId

		QueueId *string `json:"QueueId,omitempty" name:"QueueId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCmqQueueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterBrokerNetworkOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// broker网络映射信息数组

		BrokerNetworkSets []*BrokerNetworkInfoOpt `json:"BrokerNetworkSets,omitempty" name:"BrokerNetworkSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterBrokerNetworkOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterBrokerNetworkOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQConsumerConnectionDetailRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 消费组名称

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 消费端实例ID

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 按主题类型过滤查询结果，可选择Normal, GlobalOrder, PartitionedOrder, Retry, Transaction, DeadLetter

	FilterType []*string `json:"FilterType,omitempty" name:"FilterType"`
}

func (r *DescribeRocketMQConsumerConnectionDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQConsumerConnectionDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 已废弃

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
	// 已废弃

	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`
	// 说明，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateRocketMQNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryRocketMQDlqMessageRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// group名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 死信消息ID

	MessageIds []*string `json:"MessageIds,omitempty" name:"MessageIds"`
}

func (r *RetryRocketMQDlqMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryRocketMQDlqMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerLogs struct {

	// 记录数。

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 消费日志。

	ConsumerLogSets []*ConsumerLog `json:"ConsumerLogSets,omitempty" name:"ConsumerLogSets"`
}

type DeliveryStage struct {

	// 发货阶段的标识

	Stage *string `json:"Stage,omitempty" name:"Stage"`
	// 该阶段状态，
	// ASSIGNED 未开始
	// RUNNING 运行中
	// SUCCESS 成功
	// FAILURE 失败

	Status *string `json:"Status,omitempty" name:"Status"`
}

type DescribeNamespacesOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命名空间列表

		NamespaceSet []*NamespaceSetOpt `json:"NamespaceSet,omitempty" name:"NamespaceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNamespacesOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespacesOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// 主题名称，3-64个字符，只能包含字母、数字、“-”及“_”

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 主题所在的命名空间，目前支持在单个命名空间下创建主题

	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
	// 主题类型，可选值为Normal, PartitionedOrder, Transaction, DelayScheduled。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 主题说明，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 分区数，全局顺序无效

	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *CreateRocketMQTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyEncryptionRequiredOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 加密开关

	EncryptionRequired *bool `json:"EncryptionRequired,omitempty" name:"EncryptionRequired"`
}

func (r *UpdatePolicyEncryptionRequiredOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyEncryptionRequiredOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消息体

		Body *string `json:"Body,omitempty" name:"Body"`
		// 详情参数

		Properties *string `json:"Properties,omitempty" name:"Properties"`
		// 生产时间

		ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`
		// 消息id

		MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
		// 生产者地址

		ProducerAddr *string `json:"ProducerAddr,omitempty" name:"ProducerAddr"`
		// 消费组消费情况

		MessageTracks []*RocketMQMessageTrack `json:"MessageTracks,omitempty" name:"MessageTracks"`
		// 详情页展示的topic名称

		ShowTopicName *string `json:"ShowTopicName,omitempty" name:"ShowTopicName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalRocketMQNamespacesRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询起始位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询最大数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件，支持Name关键字搜索

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInternalRocketMQNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalRocketMQNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQMigrationTasksRequest struct {
	*tchttp.BaseRequest

	// 查询起始位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询数量上限

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询条件过滤器，支持以下条件：
	// type，任务类型
	// taskId，任务id
	// clusterId，集群id

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQMigrationTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQMigrationTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonMetricsOpt struct {

	// 消息生产速率

	RateIn *float64 `json:"RateIn,omitempty" name:"RateIn"`
	// 消息消费速率

	RateOut *float64 `json:"RateOut,omitempty" name:"RateOut"`
	// 入吞吐量，单位已格式化

	ThroughputIn *string `json:"ThroughputIn,omitempty" name:"ThroughputIn"`
	// 出吞吐量，单位已格式化

	ThroughputOut *string `json:"ThroughputOut,omitempty" name:"ThroughputOut"`
	// 存储消息大小入，单位已格式化

	StorageSize *string `json:"StorageSize,omitempty" name:"StorageSize"`
}

type ServerLog struct {

	// 存储时间。

	SaveTime *string `json:"SaveTime,omitempty" name:"SaveTime"`
	// 状态。

	Status *string `json:"Status,omitempty" name:"Status"`
}

type GetTopicStatsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// topic状态

		TopicStatsSet *TopicStatsSetOpt `json:"TopicStatsSet,omitempty" name:"TopicStatsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicStatsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopicStatsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublicAccessRule struct {

	// ip网段信息

	IpRule *string `json:"IpRule,omitempty" name:"IpRule"`
	// 允许或者拒绝

	Allow *bool `json:"Allow,omitempty" name:"Allow"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DeleteCmqSubscribeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅名字，在单个地域同一帐号的同一主题下唯一。订阅名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

func (r *DeleteCmqSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmqSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPQueueConsumersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 在线消费者信息

		Consumers []*AMQPQueueConsumer `json:"Consumers,omitempty" name:"Consumers"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPQueueConsumersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPQueueConsumersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBatchRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// role绑定关系

		BatchRoleBindData []*RoleBindData `json:"BatchRoleBindData,omitempty" name:"BatchRoleBindData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBatchRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBatchRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetClusterOfflineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetClusterOfflineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetClusterOfflineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQVipInstanceStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQVipInstanceStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQVipInstanceStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterLatestMetricsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标最新值

		Metrics *float64 `json:"Metrics,omitempty" name:"Metrics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterLatestMetricsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterLatestMetricsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMaxConsumerPerSubscriptionOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyMaxConsumerPerSubscriptionOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMaxConsumerPerSubscriptionOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPQueueDetail struct {

	// Queue名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 被绑定数

	DestBindedNum *uint64 `json:"DestBindedNum,omitempty" name:"DestBindedNum"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 在线消费者数

	OnlineConsumerNum *uint64 `json:"OnlineConsumerNum,omitempty" name:"OnlineConsumerNum"`
	// 每秒钟的事务数

	Tps *uint64 `json:"Tps,omitempty" name:"Tps"`
	// 消息堆积数

	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`
	// 是否自动删除

	AutoDelete *bool `json:"AutoDelete,omitempty" name:"AutoDelete"`
	// 死信交换机

	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" name:"DeadLetterExchange"`
	// 死信交换机路由键

	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" name:"DeadLetterRoutingKey"`
}

type RocketMQTopicDistribution struct {

	// topic类型

	TopicType *string `json:"TopicType,omitempty" name:"TopicType"`
	// topic数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
}

type CreateRouteRequest struct {
	*tchttp.BaseRequest

	// 1：公网 2：vpc 4：支撑网

	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 租户Vpc Id

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 租户Vpc子网Id

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// 说明，128个字符以内

	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 加密类型，0:不加密，1:ssl加密
	EncryptType *uint64 `json:"EncryptType,omitempty" name:"EncryptType"`
}

func (r *CreateRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// RocketMQ集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 3-64个字符，只能包含字母、数字、“-”及“_”

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 说明信息，不超过128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRocketMQClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendAllModulesConfigOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *SendAllModulesConfigOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendAllModulesConfigOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetClusterOfflineRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *SetClusterOfflineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetClusterOfflineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyDispatchRatePerTopicOptRequest struct {
	*tchttp.BaseRequest

	// 消费流量限制，单位为bytes

	DispatchThroughput *int64 `json:"DispatchThroughput,omitempty" name:"DispatchThroughput"`
	// 消费速率限制

	DispatchRate *int64 `json:"DispatchRate,omitempty" name:"DispatchRate"`
	// 消费限流时间窗口，单位为秒

	DispatchTimePeriod *int64 `json:"DispatchTimePeriod,omitempty" name:"DispatchTimePeriod"`
	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群(租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *UpdatePolicyDispatchRatePerTopicOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyDispatchRatePerTopicOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQMsgTraceRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题，rocketmq查询死信时值为groupId

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息id

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 消费组、订阅

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 查询死信时该值为true

	QueryDLQMsg *bool `json:"QueryDLQMsg,omitempty" name:"QueryDLQMsg"`
}

func (r *DescribeRocketMQMsgTraceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQMsgTraceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DimensionInstance struct {

	// 实例的维度组合

	Dimensions []*DimensionOpt `json:"Dimensions,omitempty" name:"Dimensions"`
}

type CreateAMQPQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名称，3-64个字符，只能包含字母、数字、“-”及“_”

	Queue *string `json:"Queue,omitempty" name:"Queue"`
	// 队列所在的vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 是否自动清除

	AutoDelete *bool `json:"AutoDelete,omitempty" name:"AutoDelete"`
	// 队列说明，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 死信exchange

	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" name:"DeadLetterExchange"`
	// 路由键

	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" name:"DeadLetterRoutingKey"`
}

func (r *CreateAMQPQueueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPQueueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接入点记录

		RouteRecords []*RouteRecord `json:"RouteRecords,omitempty" name:"RouteRecords"`
		// 记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 是否有支撑路由

		HasSupportRoute *bool `json:"HasSupportRoute,omitempty" name:"HasSupportRoute"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAMQPQueueRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 队列名称

	Queue *string `json:"Queue,omitempty" name:"Queue"`
	// 说明信息，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否自动清除

	AutoDelete *bool `json:"AutoDelete,omitempty" name:"AutoDelete"`
	// 死信exchange

	DeadLetterExchange *string `json:"DeadLetterExchange,omitempty" name:"DeadLetterExchange"`
	// 路由键

	DeadLetterRoutingKey *string `json:"DeadLetterRoutingKey,omitempty" name:"DeadLetterRoutingKey"`
}

func (r *ModifyAMQPQueueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAMQPQueueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicReplicationClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicReplicationClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicReplicationClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRocketMQPublicAccessPointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetRocketMQPublicAccessPointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRocketMQPublicAccessPointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqSubscriptionDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Subscription属性集合

		SubscriptionSet []*CmqSubscription `json:"SubscriptionSet,omitempty" name:"SubscriptionSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqSubscriptionDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqSubscriptionDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateItemsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功

		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTemplateItemsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTemplateItemsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPRouteRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPRouteRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPRouteRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqTopicRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DeleteCmqTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmqTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterOverBookieOptRequest struct {
	*tchttp.BaseRequest

	// bookie的instanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeClusterOverBookieOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterOverBookieOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicySubscriptionAuthModeOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 订阅鉴权模式，下拉框取值：
	// None、Prefix

	SubscriptionAuthMode *string `json:"SubscriptionAuthMode,omitempty" name:"SubscriptionAuthMode"`
}

func (r *UpdatePolicySubscriptionAuthModeOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicySubscriptionAuthModeOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TpoModifyResourceProjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ResourceProjectInfo 数组，新增或修改成功的资源

		ResourceProjectInfos []*ResourceProjectInfo `json:"ResourceProjectInfos,omitempty" name:"ResourceProjectInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TpoModifyResourceProjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TpoModifyResourceProjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群的ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopUsagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标值列表

		Values []*int64 `json:"Values,omitempty" name:"Values"`
		// 指标值对应的维度组合，本接口存在以下几个维度：
		// tenant，namespace，group，topic

		Dimensions []*DimensionInstance `json:"Dimensions,omitempty" name:"Dimensions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQTopUsagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQTopUsagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQInstanceSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单号

		OrderId *string `json:"OrderId,omitempty" name:"OrderId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRocketMQInstanceSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQInstanceSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterOptRequest struct {
	*tchttp.BaseRequest

	// share-共享集群；exclusive-独占集群

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 取值为cvm或eks

	ClusterResource *string `json:"ClusterResource,omitempty" name:"ClusterResource"`
	// 0-异常；1-运行中；2-待上架；3-未就绪

	ClusterState *int64 `json:"ClusterState,omitempty" name:"ClusterState"`
	// 集群名模糊匹配搜索

	ClusterNameSearch *string `json:"ClusterNameSearch,omitempty" name:"ClusterNameSearch"`
	// 查询记录条数，默认为10

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移量，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeClusterOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CmqQueue struct {

	// 消息队列ID。

	QueueId *string `json:"QueueId,omitempty" name:"QueueId"`
	// 消息队列名字。

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
	// 每秒钟生产消息条数的限制，消费消息的大小是该值的1.1倍。

	Qps *uint64 `json:"Qps,omitempty" name:"Qps"`
	// 带宽限制。

	Bps *uint64 `json:"Bps,omitempty" name:"Bps"`
	// 飞行消息最大保留时间。

	MaxDelaySeconds *uint64 `json:"MaxDelaySeconds,omitempty" name:"MaxDelaySeconds"`
	// 最大堆积消息数。取值范围在公测期间为 1,000,000 - 10,000,000，正式上线后范围可达到 1000,000-1000,000,000。默认取值在公测期间为 10,000,000，正式上线后为 100,000,000。

	MaxMsgHeapNum *uint64 `json:"MaxMsgHeapNum,omitempty" name:"MaxMsgHeapNum"`
	// 消息接收长轮询等待时间。取值范围0 - 30秒，默认值0。

	PollingWaitSeconds *uint64 `json:"PollingWaitSeconds,omitempty" name:"PollingWaitSeconds"`
	// 消息保留周期。取值范围60-1296000秒（1min-15天），默认值345600秒（4 天）。

	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`
	// 消息可见性超时。取值范围1 - 43200秒（即12小时内），默认值30。

	VisibilityTimeout *uint64 `json:"VisibilityTimeout,omitempty" name:"VisibilityTimeout"`
	// 消息最大长度。取值范围1024 - 1048576 Byte（即1K - 1024K），默认值65536。

	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`
	// 回溯队列的消息回溯时间最大值，取值范围0 - 43200秒，0表示不开启消息回溯。

	RewindSeconds *uint64 `json:"RewindSeconds,omitempty" name:"RewindSeconds"`
	// 队列的创建时间。返回 Unix 时间戳，精确到毫秒。

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后一次修改队列属性的时间。返回 Unix 时间戳，精确到毫秒。

	LastModifyTime *uint64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`
	// 在队列中处于 Active 状态（不处于被消费状态）的消息总数，为近似值。

	ActiveMsgNum *uint64 `json:"ActiveMsgNum,omitempty" name:"ActiveMsgNum"`
	// 在队列中处于 Inactive 状态（正处于被消费状态）的消息总数，为近似值。

	InactiveMsgNum *uint64 `json:"InactiveMsgNum,omitempty" name:"InactiveMsgNum"`
	// 延迟消息数。

	DelayMsgNum *uint64 `json:"DelayMsgNum,omitempty" name:"DelayMsgNum"`
	// 已调用 DelMsg 接口删除，但还在回溯保留时间内的消息数量。

	RewindMsgNum *uint64 `json:"RewindMsgNum,omitempty" name:"RewindMsgNum"`
	// 消息最小未消费时间，单位为秒。

	MinMsgTime *uint64 `json:"MinMsgTime,omitempty" name:"MinMsgTime"`
	// 事务消息队列。true表示是事务消息，false表示不是事务消息。

	Transaction *bool `json:"Transaction,omitempty" name:"Transaction"`
	// 死信队列。

	DeadLetterSource []*CmqDeadLetterSource `json:"DeadLetterSource,omitempty" name:"DeadLetterSource"`
	// 死信队列策略。

	DeadLetterPolicy *CmqDeadLetterPolicy `json:"DeadLetterPolicy,omitempty" name:"DeadLetterPolicy"`
	// 事务消息策略。

	TransactionPolicy *CmqTransactionPolicy `json:"TransactionPolicy,omitempty" name:"TransactionPolicy"`
	// 创建者Uin。

	CreateUin *uint64 `json:"CreateUin,omitempty" name:"CreateUin"`
	// 关联的标签。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 消息轨迹。true表示开启，false表示不开启。

	Trace *bool `json:"Trace,omitempty" name:"Trace"`
	// 租户id

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

type CreateBindVpcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户Vpc Id。

		UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
		// 租户Vpc子网Id。

		UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
		// 路由Id。

		RouterId *string `json:"RouterId,omitempty" name:"RouterId"`
		// Vpc的Ip。

		Ip *string `json:"Ip,omitempty" name:"Ip"`
		// Vpc的Port。

		Port *uint64 `json:"Port,omitempty" name:"Port"`
		// 说明，128个字符以内。

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBindVpcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBindVpcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterOptRequest struct {
	*tchttp.BaseRequest

	// share-共享集群；exclusive-独占集群

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 取值为cvm或eks

	ClusterResource *string `json:"ClusterResource,omitempty" name:"ClusterResource"`
	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// pulsar admin url

	AdminUrl *string `json:"AdminUrl,omitempty" name:"AdminUrl"`
	// Tracing上报地址

	TracingReport *string `json:"TracingReport,omitempty" name:"TracingReport"`
	// CLB实例ID

	ClbInstanceId *string `json:"ClbInstanceId,omitempty" name:"ClbInstanceId"`
	// broker ip列表

	BrokerIps []*string `json:"BrokerIps,omitempty" name:"BrokerIps"`
	// bookie ip列表

	BookieIps []*string `json:"BookieIps,omitempty" name:"BookieIps"`
	// broker zk ip列表

	BrokerZookeeperIps []*string `json:"BrokerZookeeperIps,omitempty" name:"BrokerZookeeperIps"`
	// bookie zk ip列表

	BookieZookeeperIps []*string `json:"BookieZookeeperIps,omitempty" name:"BookieZookeeperIps"`
	// broker规格：
	// 0-入门；1-基础；2-标准；3-高性能

	BrokerSpecification *int64 `json:"BrokerSpecification,omitempty" name:"BrokerSpecification"`
	// 协议名数组

	ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`
	// 物理集群admin secretKey

	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`
	// 物理集群admin token

	AdminToken *string `json:"AdminToken,omitempty" name:"AdminToken"`
	// data-reporter上报地址

	DataReporterServer *string `json:"DataReporterServer,omitempty" name:"DataReporterServer"`
	// Tracing上报token

	TracingToken *string `json:"TracingToken,omitempty" name:"TracingToken"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

func (r *CreateClusterOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterComponentMetricsListOptRequest struct {
	*tchttp.BaseRequest

	// 视图名，取值：
	// aggregation_metrics：broker监控指标
	// bookeeper_metrics：bookie监控指标
	// broker_zk_metrics：broker-zk监控指标
	// bookie_zk_metrics：bookie-zk监控指标

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

func (r *DescribeClusterComponentMetricsListOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterComponentMetricsListOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueueDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 队列详情列表。

		QueueDescribe *CmqQueue `json:"QueueDescribe,omitempty" name:"QueueDescribe"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqQueueDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqQueueDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClusterBaradMetricsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标值数组

		MetricsList []*float64 `json:"MetricsList,omitempty" name:"MetricsList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClusterBaradMetricsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClusterBaradMetricsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineTypeListOptRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMachineTypeListOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineTypeListOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentRolesRequest struct {
	*tchttp.BaseRequest

	// 必填字段，环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 起始下标，不填默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 必填字段，Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 角色名称

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// * RoleName
	// 按照角色名进行过滤，精确查询。
	// 类型：String
	// 必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEnvironmentRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnvironmentRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPExchangeRequest struct {
	*tchttp.BaseRequest

	// 交换机名称，3-64个字符，只能包含字母、数字、“-”及“_”

	Exchange *string `json:"Exchange,omitempty" name:"Exchange"`
	// 交换机所在的vhost，目前支持在单个vhost下创建主题

	VHosts []*string `json:"VHosts,omitempty" name:"VHosts"`
	// 交换机类型，可选值为Direct, Fanout, Topic

	Type *string `json:"Type,omitempty" name:"Type"`
	// 交换机说明，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 备用交换机名称

	AlternateExchange *string `json:"AlternateExchange,omitempty" name:"AlternateExchange"`
}

func (r *CreateAMQPExchangeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPExchangeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqTopicAttributeRequest struct {
	*tchttp.BaseRequest

	// 主题名字，在单个地域同一帐号下唯一。主题名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息最大长度。取值范围1024 - 65536 Byte（即1 - 64K），默认值65536。

	MaxMsgSize *uint64 `json:"MaxMsgSize,omitempty" name:"MaxMsgSize"`
	// 消息保存时间。取值范围60 - 86400 s（即1分钟 - 1天），默认值86400。

	MsgRetentionSeconds *uint64 `json:"MsgRetentionSeconds,omitempty" name:"MsgRetentionSeconds"`
	// 是否开启消息轨迹标识，true表示开启，false表示不开启，不填表示不开启。

	Trace *bool `json:"Trace,omitempty" name:"Trace"`
}

func (r *ModifyCmqTopicAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmqTopicAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMsgRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称，如果是分区topic需要指定具体分区，如果没有指定则默认发到0分区，例如：my_topic-partition-0。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 消息内容，不能为空且大小不得大于5242880个byte。

	MsgContent *string `json:"MsgContent,omitempty" name:"MsgContent"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *SendMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色名称

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 角色token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 备注说明

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 批量操作环境信息

		EnvironmentRoleSets []*EnvironmentRoleSet `json:"EnvironmentRoleSets,omitempty" name:"EnvironmentRoleSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间记录数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命名空间集合数组。

		EnvironmentSet []*Environment `json:"EnvironmentSet,omitempty" name:"EnvironmentSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterRequest struct {
	*tchttp.BaseRequest

	// Pulsar 集群的ID，需要更新的集群Id。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 更新后的集群名称。

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 说明信息。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetentionPolicy struct {

	// 消息保留时长

	TimeInMinutes *int64 `json:"TimeInMinutes,omitempty" name:"TimeInMinutes"`
	// 消息保留大小

	SizeInMB *int64 `json:"SizeInMB,omitempty" name:"SizeInMB"`
}

type DescribeClusterConfigOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置项集合（BASE64编码后）

		ConfigItems *string `json:"ConfigItems,omitempty" name:"ConfigItems"`
		// 是否Shell

		IsShell *bool `json:"IsShell,omitempty" name:"IsShell"`
		// 配置基础信息

		ConfigBaseInfo *ConfigBaseInfoOpt `json:"ConfigBaseInfo,omitempty" name:"ConfigBaseInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterConfigOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterConfigOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqSubscriptionAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqSubscriptionAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmqSubscriptionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMaxConsumerPerTopicOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 消费者数量限制

	MaxConsumer *int64 `json:"MaxConsumer,omitempty" name:"MaxConsumer"`
}

func (r *UpdatePolicyMaxConsumerPerTopicOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMaxConsumerPerTopicOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQSubscription struct {

	// 主题名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 主题类型：
	// Normal 普通,
	// GlobalOrder 全局顺序,
	// PartitionedOrder 局部顺序,
	// Transaction 事务消息,
	// DelayScheduled 延时消息,
	// Retry 重试,
	// DeadLetter 死信

	Type *string `json:"Type,omitempty" name:"Type"`
	// 分区数

	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
	// 过滤模式，TAG，SQL

	ExpressionType *string `json:"ExpressionType,omitempty" name:"ExpressionType"`
	// 过滤表达式

	SubString *string `json:"SubString,omitempty" name:"SubString"`
	// 订阅状态：
	// 0，订阅关系一致
	// 1，订阅关系不一致
	// 2，未知

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAMQPVHostsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Vhost 列表

		VHosts []*AMQPVHost `json:"VHosts,omitempty" name:"VHosts"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPVHostsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPVHostsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInternalRocketMQInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 修改最多可创建命名空间数量

	MaxNamespaces *uint64 `json:"MaxNamespaces,omitempty" name:"MaxNamespaces"`
	// 修改最多可创建主题数量

	MaxTopics *uint64 `json:"MaxTopics,omitempty" name:"MaxTopics"`
	// 修改最多可创建消费组数量

	MaxGroups *uint64 `json:"MaxGroups,omitempty" name:"MaxGroups"`
	// 修改单节点主题最大可调整队列数量

	MaxQueuesPerTopic *uint64 `json:"MaxQueuesPerTopic,omitempty" name:"MaxQueuesPerTopic"`
	// 修改集群限流次数

	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`
	// 修改实例消息保留时间，仅限专享实例，小时为单位

	MessageRetention *int64 `json:"MessageRetention,omitempty" name:"MessageRetention"`
	// 修改最大可调整消息保留时间，小时为单位

	MaxRetention *int64 `json:"MaxRetention,omitempty" name:"MaxRetention"`
	// 修改最小可调整消息保留时间，小时为单位

	MinRetention *int64 `json:"MinRetention,omitempty" name:"MinRetention"`
	// 是否开启acl，默认开启true

	AclEnabled *string `json:"AclEnabled,omitempty" name:"AclEnabled"`
}

func (r *ModifyInternalRocketMQInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInternalRocketMQInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RoleBindData struct {

	// 角色名字

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 角色说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 集群名字

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 名字空间Id

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 角色权限枚举值

	Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
}

type DescribeAMQPClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		ClusterInfo *AMQPClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`
		// 集群配置

		ClusterConfig *AMQPClusterConfig `json:"ClusterConfig,omitempty" name:"ClusterConfig"`
		// 集群最近使用量

		ClusterStats *AMQPClusterRecentStats `json:"ClusterStats,omitempty" name:"ClusterStats"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeEnvironmentAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnvironmentAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQNamespaceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 已废弃

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
	// 已废弃

	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`
	// 说明，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyRocketMQNamespaceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQNamespaceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MsgLog struct {

	// 消息ID。

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 生产者名称。

	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`
	// 生产时间。

	ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`
	// 生产客户端地址。

	ProducerAddr *string `json:"ProducerAddr,omitempty" name:"ProducerAddr"`
}

type PartitionedTopicInternalStatsSetOpt struct {

	// 分区数量，非分区主题取值为-1

	PartitionIndexes *int64 `json:"PartitionIndexes,omitempty" name:"PartitionIndexes"`
	// 内部主题状态数组

	InternalStatsSets []*InternalStatsSetOpt `json:"InternalStatsSets,omitempty" name:"InternalStatsSets"`
}

type PortalConfig struct {

	// 是否允许用户自行配置路由

	AllowRouteConfig *bool `json:"AllowRouteConfig,omitempty" name:"AllowRouteConfig"`
	// 集群版本号

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeProducersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 生产者集合数组。

		ProducerSets []*Producer `json:"ProducerSets,omitempty" name:"ProducerSets"`
		// 记录总数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProducersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProducersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPVHostConnection struct {

	// 客户端实例的地址和端口

	ClientAddr *string `json:"ClientAddr,omitempty" name:"ClientAddr"`
	// 客户端协议

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 是否支持SSL/TLS

	Security *bool `json:"Security,omitempty" name:"Security"`
	// 信道数

	ChannelNum *uint64 `json:"ChannelNum,omitempty" name:"ChannelNum"`
	// 接受速率

	ReceiveRate *uint64 `json:"ReceiveRate,omitempty" name:"ReceiveRate"`
	// 发送速率

	SendRate *uint64 `json:"SendRate,omitempty" name:"SendRate"`
}

type InstanceNodeDistribution struct {

	// 可用区

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 可用区id

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 节点数

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
}

type InternalTenant struct {

	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 虚拟集群名称

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户的APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群协议类型，支持的值为TDMQ，ROCKETMQ，AMQP，CMQ

	Type *string `json:"Type,omitempty" name:"Type"`
	// 命名空间配额

	MaxNamespaces *int64 `json:"MaxNamespaces,omitempty" name:"MaxNamespaces"`
	// 已使用命名空间配额

	UsedNamespaces *int64 `json:"UsedNamespaces,omitempty" name:"UsedNamespaces"`
	// Topic配额

	MaxTopics *int64 `json:"MaxTopics,omitempty" name:"MaxTopics"`
	// 已使用Topic配额

	UsedTopics *int64 `json:"UsedTopics,omitempty" name:"UsedTopics"`
	// Topic分区数配额

	MaxPartitions *int64 `json:"MaxPartitions,omitempty" name:"MaxPartitions"`
	// 已使用Topic分区数配额

	UsedPartitions *int64 `json:"UsedPartitions,omitempty" name:"UsedPartitions"`
	// 存储配额, byte为单位

	MaxMsgBacklogSize *uint64 `json:"MaxMsgBacklogSize,omitempty" name:"MaxMsgBacklogSize"`
	// 命名空间最大生产TPS

	MaxPublishTps *uint64 `json:"MaxPublishTps,omitempty" name:"MaxPublishTps"`
	// 消息最大保留时间，秒为单位

	MaxRetention *uint64 `json:"MaxRetention,omitempty" name:"MaxRetention"`
	// 创建时间，毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间，毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 命名空间最大消费TPS

	MaxDispatchTps *uint64 `json:"MaxDispatchTps,omitempty" name:"MaxDispatchTps"`
	// 命名空间最大消费带宽，byte为单位

	MaxDispatchRateInBytes *uint64 `json:"MaxDispatchRateInBytes,omitempty" name:"MaxDispatchRateInBytes"`
	// 命名空间最大生产带宽，byte为单位

	MaxPublishRateInBytes *uint64 `json:"MaxPublishRateInBytes,omitempty" name:"MaxPublishRateInBytes"`
	// 消息最大保留空间，MB为单位

	MaxRetentionSizeInMB *uint64 `json:"MaxRetentionSizeInMB,omitempty" name:"MaxRetentionSizeInMB"`
	// public Access Enabled

	PublicAccessEnabled *bool `json:"PublicAccessEnabled,omitempty" name:"PublicAccessEnabled"`
}

type CreateEnvironmentRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 角色名称。

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 授权项。

		Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnvironmentRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteRequest struct {
	*tchttp.BaseRequest

	// 路由id

	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 网络类型

	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`
	// vpc地址

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 子网地址

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
}

func (r *DeleteRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointsOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
}

func (r *DescribeClusterEndpointsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterEndpointsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 主题名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 说明信息，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 分区数，全局类型无效，不可小于当前分区数

	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *ModifyRocketMQTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMaxConsumerPerSubscriptionOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 消费者数量限制

	MaxConsumer *int64 `json:"MaxConsumer,omitempty" name:"MaxConsumer"`
}

func (r *UpdatePolicyMaxConsumerPerSubscriptionOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMaxConsumerPerSubscriptionOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetRocketMQPublicAccessPointRequest struct {
	*tchttp.BaseRequest

	// 集群ID，当前只支持专享集群

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开启或关闭访问

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 带宽大小，开启或者调整公网时必须指定，Mbps为单位

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 付费模式，开启公网时必须指定，0为按小时计费，1为包年包月，当前只支持按小时计费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 公网访问安全规则列表，Enabled为true时必须传入

	Rules []*PublicAccessRule `json:"Rules,omitempty" name:"Rules"`
}

func (r *SetRocketMQPublicAccessPointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetRocketMQPublicAccessPointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTemplateOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功

		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTemplateOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTemplateOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPClusterRequest struct {
	*tchttp.BaseRequest

	// 待删除的集群Id。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteAMQPClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQTopicRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 主题名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
}

func (r *DeleteRocketMQTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigModuleInfoOpt struct {

	// 模块名

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 组件名

	ComponentName *string `json:"ComponentName,omitempty" name:"ComponentName"`
	// 配置文件名

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
	// 是否脚本文件

	IsShell *bool `json:"IsShell,omitempty" name:"IsShell"`
}

type TemplateItemInfoOpt struct {

	// 配置属性名

	ItemKey *string `json:"ItemKey,omitempty" name:"ItemKey"`
	// 配置项默认值

	ItemDefaultValue *string `json:"ItemDefaultValue,omitempty" name:"ItemDefaultValue"`
	// 模板配置项值

	ItemValue *string `json:"ItemValue,omitempty" name:"ItemValue"`
	// 描述

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 是否动态配置，0需要重启，1不需要重启

	IsDynamic *bool `json:"IsDynamic,omitempty" name:"IsDynamic"`
}

type DescribeClusterExtraRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterExtraRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterExtraRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueuesRequest struct {
	*tchttp.BaseRequest

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据QueueName进行过滤

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
	// CMQ 队列名称列表过滤

	QueueNameList []*string `json:"QueueNameList,omitempty" name:"QueueNameList"`
	// 标签过滤查找时，需要设置为 true

	IsTagFilter *bool `json:"IsTagFilter,omitempty" name:"IsTagFilter"`
}

func (r *DescribeCmqQueuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqQueuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RewindCmqQueueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RewindCmqQueueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RewindCmqQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）数组，每次最多删除20个。

	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteEnvironmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEnvironmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTemplateItemsOptRequest struct {
	*tchttp.BaseRequest

	// 配置模块名称（小写）

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 0-S5 1-SA

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
	// 配置项

	TemplateItems []*ConfigTemplateItemOpt `json:"TemplateItems,omitempty" name:"TemplateItems"`
}

func (r *DeleteTemplateItemsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTemplateItemsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReceiveMessageRequest struct {
	*tchttp.BaseRequest

	// 接收消息的topic的名字, 这里尽量需要使用topic的全路径，如果不指定，即：tenant/namespace/topic。默认使用的是：public/default

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 订阅者的名字

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 默认值为1000，consumer接收的消息会首先存储到receiverQueueSize这个队列中，用作调优接收消息的速率

	ReceiverQueueSize *int64 `json:"ReceiverQueueSize,omitempty" name:"ReceiverQueueSize"`
	// 默认值为：Latest。用作判定consumer初始接收消息的位置，可选参数为：Earliest, Latest

	SubInitialPosition *string `json:"SubInitialPosition,omitempty" name:"SubInitialPosition"`
}

func (r *ReceiveMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReceiveMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPClusterRecentStats struct {

	// Queue数量

	QueueNum *uint64 `json:"QueueNum,omitempty" name:"QueueNum"`
	// 消息生产数

	ProducedMsgNum *uint64 `json:"ProducedMsgNum,omitempty" name:"ProducedMsgNum"`
	// 消息堆积数

	AccumulativeMsgNum *uint64 `json:"AccumulativeMsgNum,omitempty" name:"AccumulativeMsgNum"`
	// Exchange数量

	ExchangeNum *uint64 `json:"ExchangeNum,omitempty" name:"ExchangeNum"`
}

type ModifyAMQPVHostRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// vhost名称，3-64个字符，只能包含字母、数字、“-”及“_”

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 未消费消息的保留时间，以毫秒为单位，60秒-15天

	MsgTtl *uint64 `json:"MsgTtl,omitempty" name:"MsgTtl"`
	// 说明，最大128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyAMQPVHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAMQPVHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQGroup struct {

	// 消费组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 在线消费者数量

	ConsumerNum *uint64 `json:"ConsumerNum,omitempty" name:"ConsumerNum"`
	// 消费TPS

	TPS *uint64 `json:"TPS,omitempty" name:"TPS"`
	// 总堆积数量

	TotalAccumulative *int64 `json:"TotalAccumulative,omitempty" name:"TotalAccumulative"`
	// 0表示集群消费模式，1表示广播消费模式，-1表示未知

	ConsumptionMode *int64 `json:"ConsumptionMode,omitempty" name:"ConsumptionMode"`
	// 是否允许消费

	ReadEnabled *bool `json:"ReadEnabled,omitempty" name:"ReadEnabled"`
	// 重试队列分区数

	RetryPartitionNum *uint64 `json:"RetryPartitionNum,omitempty" name:"RetryPartitionNum"`
	// 创建时间，以毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间，以毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 客户端协议

	ClientProtocol *string `json:"ClientProtocol,omitempty" name:"ClientProtocol"`
	// 说明信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 消费者类型，枚举值ACTIVELY, PASSIVELY

	ConsumerType *string `json:"ConsumerType,omitempty" name:"ConsumerType"`
	// 是否开启广播消费

	BroadcastEnabled *bool `json:"BroadcastEnabled,omitempty" name:"BroadcastEnabled"`
	// Group类型

	GroupType *string `json:"GroupType,omitempty" name:"GroupType"`
	// 重试次数

	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitempty" name:"RetryMaxTimes"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type ClearCmqSubscriptionFilterTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClearCmqSubscriptionFilterTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClearCmqSubscriptionFilterTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscriptionTopic struct {

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅名称。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

type DimensionOpt struct {

	// 查询的维度名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 查询维度的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 命名空间名称，模糊搜索。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// * EnvironmentId
	// 按照名称空间进行过滤，精确查询。
	// 类型：String
	// 必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQRolesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 角色数组。

		RoleSets []*Role `json:"RoleSets,omitempty" name:"RoleSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQRolesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQRolesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerStatsSetOpt struct {

	// 消息消费速率

	MsgRateOut *float64 `json:"MsgRateOut,omitempty" name:"MsgRateOut"`
	// 消息消费吞吐量

	MsgThroughputOut *float64 `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`
	// 读出字节总数

	BytesOutCounter *int64 `json:"BytesOutCounter,omitempty" name:"BytesOutCounter"`
	// 消息消费总数

	MsgOutCounter *int64 `json:"MsgOutCounter,omitempty" name:"MsgOutCounter"`
	// 消息重传速率

	MsgRateRedeliver *float64 `json:"MsgRateRedeliver,omitempty" name:"MsgRateRedeliver"`
	// 分块消息消费速率

	ChunkedMessageRate *float64 `json:"ChunkedMessageRate,omitempty" name:"ChunkedMessageRate"`
	// 可用令牌数

	AvailablePermits *int64 `json:"AvailablePermits,omitempty" name:"AvailablePermits"`
	// 未确认消息数

	UnackedMessages *int64 `json:"UnackedMessages,omitempty" name:"UnackedMessages"`
	// Entry中平均消息数

	AvgMessagesPerEntry *int64 `json:"AvgMessagesPerEntry,omitempty" name:"AvgMessagesPerEntry"`
	// 当Consumer未确认消息超过限制时，是否阻塞消费者继续消费

	BlockedConsumerOnUnackedMsgs *bool `json:"BlockedConsumerOnUnackedMsgs,omitempty" name:"BlockedConsumerOnUnackedMsgs"`
	// 最近一次确认消息时间

	LastAckedTimestamp *int64 `json:"LastAckedTimestamp,omitempty" name:"LastAckedTimestamp"`
	// 最近一次消费时间

	LastConsumedTimestamp *int64 `json:"LastConsumedTimestamp,omitempty" name:"LastConsumedTimestamp"`
}

type AddPulsarResourceAllocationRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddPulsarResourceAllocationRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddPulsarResourceAllocationRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTemplateOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功

		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAMQPRouteRelationRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 路由关系ID

	RouteRelationId *string `json:"RouteRelationId,omitempty" name:"RouteRelationId"`
}

func (r *DeleteAMQPRouteRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPRouteRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicDetailRequest struct {
	*tchttp.BaseRequest

	// 精确匹配TopicName。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeCmqTopicDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqTopicDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProtocolsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 协议数组

		ProtocolSets []*ProtocolSetOpt `json:"ProtocolSets,omitempty" name:"ProtocolSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProtocolsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProtocolsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicMsgsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 消息日志列表。

		TopicMsgLogSets []*MsgLog `json:"TopicMsgLogSets,omitempty" name:"TopicMsgLogSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicMsgsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicMsgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClustersRequest struct {
	*tchttp.BaseRequest

	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数目

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 按照集群ID关键字搜索

	IdKeyword *string `json:"IdKeyword,omitempty" name:"IdKeyword"`
	// 按照集群名称关键字搜索

	NameKeyword *string `json:"NameKeyword,omitempty" name:"NameKeyword"`
	// 集群ID列表过滤

	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`
	// 标签过滤查找时，需要设置为true

	IsTagFilter *bool `json:"IsTagFilter,omitempty" name:"IsTagFilter"`
	// 过滤器。目前支持标签过滤。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 主题信息列表

		Topics []*RocketMQTopic `json:"Topics,omitempty" name:"Topics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQMigrationTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 子任务查询起始位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 子任务查询最大数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 子任务查询条件，支持：
	// code，根据返回码查询

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRocketMQMigrationTaskDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQMigrationTaskDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAMQPRouteRelationsRequest struct {
	*tchttp.BaseRequest

	// 查询偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 查询限制数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Vhost名称

	VHostId *string `json:"VHostId,omitempty" name:"VHostId"`
	// 按源exchange名称过滤查询结果，支持模糊查询

	FilterSourceExchange *string `json:"FilterSourceExchange,omitempty" name:"FilterSourceExchange"`
	// 按绑定的目标类型过滤查询结果，可选值:Exchange、Queue

	FilterDestType *string `json:"FilterDestType,omitempty" name:"FilterDestType"`
	// 按目标名称过滤查询结果，支持模糊查询

	FilterDestValue *string `json:"FilterDestValue,omitempty" name:"FilterDestValue"`
}

func (r *DescribeAMQPRouteRelationsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPRouteRelationsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMessageRetentionOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 消息保留时间，单位为分钟

	RetentionTime *int64 `json:"RetentionTime,omitempty" name:"RetentionTime"`
	// 消息保留大小，单位为MB

	RetentionSize *int64 `json:"RetentionSize,omitempty" name:"RetentionSize"`
}

func (r *UpdatePolicyMessageRetentionOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMessageRetentionOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Environment struct {

	// 命名空间名称

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 未消费消息过期时间，单位：秒，最大1296000（15天）

	MsgTTL *int64 `json:"MsgTTL,omitempty" name:"MsgTTL"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 命名空间ID

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// Topic数量

	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`
	// 消息保留策略

	RetentionPolicy *RetentionPolicy `json:"RetentionPolicy,omitempty" name:"RetentionPolicy"`
}

type CreateTemplateItemsOptRequest struct {
	*tchttp.BaseRequest

	// 配置模块名称（小写）

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 0-S5 1-SA

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
	// 配置项

	TemplateItems []*ConfigTemplateItemOpt `json:"TemplateItems,omitempty" name:"TemplateItems"`
}

func (r *CreateTemplateItemsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateItemsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqQueuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 队列列表

		QueueList []*CmqQueue `json:"QueueList,omitempty" name:"QueueList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqQueuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendAllModulesConfigOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下发结果

		SendFileResult []*SendConfigInfoOpt `json:"SendFileResult,omitempty" name:"SendFileResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendAllModulesConfigOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendAllModulesConfigOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订阅者集合数组。

		SubscriptionSets []*Subscription `json:"SubscriptionSets,omitempty" name:"SubscriptionSets"`
		// 数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscriptionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NamespaceSetOpt struct {

	// 命名空间ID

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 客户uin

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户appId

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 虚拟集群（租户）id

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 虚拟集群（租户）名称

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// Bundle数量

	BundleCount *int64 `json:"BundleCount,omitempty" name:"BundleCount"`
	// 生产消息速率限制

	RateInLimit *string `json:"RateInLimit,omitempty" name:"RateInLimit"`
	// 消费消息速率限制

	RateOutLimit *string `json:"RateOutLimit,omitempty" name:"RateOutLimit"`
	// 入吞吐量限制

	ThroughputInLimit *string `json:"ThroughputInLimit,omitempty" name:"ThroughputInLimit"`
	// 出吞吐量限制

	ThroughputOutLimit *string `json:"ThroughputOutLimit,omitempty" name:"ThroughputOutLimit"`
	// TTL时间，单位为秒

	MessageTtl *int64 `json:"MessageTtl,omitempty" name:"MessageTtl"`
	// 消息保留时间，单位为分钟

	RetentionTime *int64 `json:"RetentionTime,omitempty" name:"RetentionTime"`
	// 命名空间监控指标

	Metrics *CommonMetricsOpt `json:"Metrics,omitempty" name:"Metrics"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type VpcBindRecord struct {

	// 租户Vpc Id

	UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
	// 租户Vpc子网Id

	UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
	// 路由Id

	RouterId *string `json:"RouterId,omitempty" name:"RouterId"`
	// Vpc的Id

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Vpc的Port

	Port *uint64 `json:"Port,omitempty" name:"Port"`
	// 说明，128个字符以内

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DescribeBindVpcsRequest struct {
	*tchttp.BaseRequest

	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeBindVpcsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBindVpcsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopicInternalStatsOptRequest struct {
	*tchttp.BaseRequest

	// 完整topic名

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *GetTopicInternalStatsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopicInternalStatsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TpoDescribeResourcesAdminProjectsRequest struct {
	*tchttp.BaseRequest
}

func (r *TpoDescribeResourcesAdminProjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TpoDescribeResourcesAdminProjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalRocketMQTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 分页后的主题列表结果

		Topics []*RocketMQTopic `json:"Topics,omitempty" name:"Topics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternalRocketMQTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalRocketMQTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicySubscriptionAuthModeOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicySubscriptionAuthModeOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicySubscriptionAuthModeOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicsRequest struct {
	*tchttp.BaseRequest

	// 主题集合，每次最多删除20个。

	TopicSets []*TopicRecord `json:"TopicSets,omitempty" name:"TopicSets"`
	// pulsar集群Id。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 是否强制删除，默认为false

	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqDeadLetterSourceQueuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 满足本次条件的队列个数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 死信队列源队列

		QueueSet []*CmqDeadLetterSource `json:"QueueSet,omitempty" name:"QueueSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCmqDeadLetterSourceQueuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqDeadLetterSourceQueuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 消息ID。

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 主题名。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TraceResult struct {

	// 阶段

	Stage *string `json:"Stage,omitempty" name:"Stage"`
	// 内容详情

	Data *string `json:"Data,omitempty" name:"Data"`
}

type DeleteAMQPQueueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAMQPQueueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAMQPQueueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryRocketMQDlqMessageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryRocketMQDlqMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryRocketMQDlqMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPVHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAMQPVHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPVHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterDetailOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 物理集群名

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 集群资源类型：cvm或eks

		ClusterResource *string `json:"ClusterResource,omitempty" name:"ClusterResource"`
		// share-共享集群；exclusive-独占集群

		ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
		// Tracing上报地址

		TracingReport *string `json:"TracingReport,omitempty" name:"TracingReport"`
		// 内网CLB实例ID

		ClbInstanceId *string `json:"ClbInstanceId,omitempty" name:"ClbInstanceId"`
		// 协议名列表

		ProtocolList []*string `json:"ProtocolList,omitempty" name:"ProtocolList"`
		// broker ip列表

		BrokerIps []*string `json:"BrokerIps,omitempty" name:"BrokerIps"`
		// bookie ip列表

		BookieIps []*string `json:"BookieIps,omitempty" name:"BookieIps"`
		// broker zk ip列表

		BrokerZookeeperIps []*string `json:"BrokerZookeeperIps,omitempty" name:"BrokerZookeeperIps"`
		// bookie zk ip列表

		BookieZookeeperIps []*string `json:"BookieZookeeperIps,omitempty" name:"BookieZookeeperIps"`
		// 0-异常；1-运行中；2-待上架；3-未就绪

		ClusterState *int64 `json:"ClusterState,omitempty" name:"ClusterState"`
		// 创建时间

		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
		// 修改时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// broker规格：
		// 0-入门；1-基础；2-标准；3-高性能

		BrokerSpecification *int64 `json:"BrokerSpecification,omitempty" name:"BrokerSpecification"`
		// Admin url

		AdminUrl *string `json:"AdminUrl,omitempty" name:"AdminUrl"`
		// data-reporter上报地址

		DataReporterServer *string `json:"DataReporterServer,omitempty" name:"DataReporterServer"`
		// 集群版本

		ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
		// 备注

		Comment *string `json:"Comment,omitempty" name:"Comment"`
		// 公网CLB实例ID

		PublicClb *string `json:"PublicClb,omitempty" name:"PublicClb"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterDetailOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterDetailOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgTraceRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 消息ID。

	MsgId *string `json:"MsgId,omitempty" name:"MsgId"`
	// 消息生产时间。

	ProduceTime *string `json:"ProduceTime,omitempty" name:"ProduceTime"`
	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 消费组名称模糊匹配。

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeMsgTraceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMsgTraceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateItemOptRequest struct {
	*tchttp.BaseRequest

	// 配置模块名称（小写）

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *int64 `json:"TemplateType,omitempty" name:"TemplateType"`
	// 0-S5 1-SA

	MachineType *int64 `json:"MachineType,omitempty" name:"MachineType"`
}

func (r *DescribeTemplateItemOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateItemOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQConsumerClientsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 在线消费者信息

		Clients []*RocketMQConsumerClient `json:"Clients,omitempty" name:"Clients"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQConsumerClientsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQConsumerClientsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeHealthOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0-异常；1-正常

		NodeState *int64 `json:"NodeState,omitempty" name:"NodeState"`
		// 最近一次健康检查的时间

		LatestHealthCheckTime *string `json:"LatestHealthCheckTime,omitempty" name:"LatestHealthCheckTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeHealthOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeHealthOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyPersistenceOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 单个 topic 使用的 bookie 数量，默认值：2

	PersistenceEnsemble *int64 `json:"PersistenceEnsemble,omitempty" name:"PersistenceEnsemble"`
	// 每个 entry 在等待的 acks（有保证的副本）数量，默认值：2

	PersistenceAckQuorum *int64 `json:"PersistenceAckQuorum,omitempty" name:"PersistenceAckQuorum"`
	// 每个 entry 要写入的次数，默认值：2

	PersistenceWriteQuorum *int64 `json:"PersistenceWriteQuorum,omitempty" name:"PersistenceWriteQuorum"`
	// 标记-删除操作的限制速率（0表示无限制），默认值：0.0

	PersistenceMarkDeleteRate *float64 `json:"PersistenceMarkDeleteRate,omitempty" name:"PersistenceMarkDeleteRate"`
}

func (r *UpdatePolicyPersistenceOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyPersistenceOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AMQPClusterConfig struct {

	// 单Vhost TPS上限

	MaxTpsPerVHost *uint64 `json:"MaxTpsPerVHost,omitempty" name:"MaxTpsPerVHost"`
	// 单Vhost客户端连接数上限

	MaxConnNumPerVHost *uint64 `json:"MaxConnNumPerVHost,omitempty" name:"MaxConnNumPerVHost"`
	// 最大Vhost数量

	MaxVHostNum *uint64 `json:"MaxVHostNum,omitempty" name:"MaxVHostNum"`
	// 最大exchange数量

	MaxExchangeNum *uint64 `json:"MaxExchangeNum,omitempty" name:"MaxExchangeNum"`
	// 最大Queue数量

	MaxQueueNum *uint64 `json:"MaxQueueNum,omitempty" name:"MaxQueueNum"`
	// 消息最大保留时间，以毫秒为单位

	MaxRetentionTime *uint64 `json:"MaxRetentionTime,omitempty" name:"MaxRetentionTime"`
	// 已使用Vhost数量

	UsedVHostNum *uint64 `json:"UsedVHostNum,omitempty" name:"UsedVHostNum"`
	// 已使用exchange数量

	UsedExchangeNum *uint64 `json:"UsedExchangeNum,omitempty" name:"UsedExchangeNum"`
	// 已使用queue数量

	UsedQueueNum *uint64 `json:"UsedQueueNum,omitempty" name:"UsedQueueNum"`
}

type BrokerNetworkInfoOpt struct {

	// broker ip地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 支撑网本地映射

	InternalMappingSelf *string `json:"InternalMappingSelf,omitempty" name:"InternalMappingSelf"`
	// 支撑网CLB映射

	InternalMappingRemote *string `json:"InternalMappingRemote,omitempty" name:"InternalMappingRemote"`
	// VPC本地映射

	VpcMappingSelf *string `json:"VpcMappingSelf,omitempty" name:"VpcMappingSelf"`
	// VPC对端映射

	VpcMappingRemote *string `json:"VpcMappingRemote,omitempty" name:"VpcMappingRemote"`
	// 公网本地映射

	PublicMappingSelf *string `json:"PublicMappingSelf,omitempty" name:"PublicMappingSelf"`
	// 公网CLB映射

	PublicMappingRemote *string `json:"PublicMappingRemote,omitempty" name:"PublicMappingRemote"`
	// 内网打通状态：0-未打通；1-打通成功；2-打通失败

	InternalMappingStatus *int64 `json:"InternalMappingStatus,omitempty" name:"InternalMappingStatus"`
	// VPC打通状态：0-未打通；1-打通成功；2-打通失败

	VpcMappingStatus *int64 `json:"VpcMappingStatus,omitempty" name:"VpcMappingStatus"`
	// 公网打通状态：0-未打通；1-打通成功；2-打通失败

	PublicMappingStatus *int64 `json:"PublicMappingStatus,omitempty" name:"PublicMappingStatus"`
}

type CheckBatchRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0,检测通过，1 检测不通过

		CheckCode *uint64 `json:"CheckCode,omitempty" name:"CheckCode"`
		// 枚举  原因

		CheckMessage *string `json:"CheckMessage,omitempty" name:"CheckMessage"`
		// 异常资源列表

		AbnormalResources []*string `json:"AbnormalResources,omitempty" name:"AbnormalResources"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckBatchRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckBatchRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未消费消息过期时间，单位：秒，最大1296000（15天）。

		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`
		// 消费速率限制，单位：byte/秒，0：不限速。

		RateInByte *uint64 `json:"RateInByte,omitempty" name:"RateInByte"`
		// 消费速率限制，单位：个数/秒，0：不限速。

		RateInSize *uint64 `json:"RateInSize,omitempty" name:"RateInSize"`
		// 已消费消息保存策略，单位：小时，0：消费完马上删除。

		RetentionHours *uint64 `json:"RetentionHours,omitempty" name:"RetentionHours"`
		// 已消费消息保存策略，单位：G，0：消费完马上删除。

		RetentionSize *uint64 `json:"RetentionSize,omitempty" name:"RetentionSize"`
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 副本数。

		Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`
		// 备注。

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEnvironmentAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQRolesRequest struct {
	*tchttp.BaseRequest

	// 角色名称数组。

	RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`
	// 必填字段，集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteRocketMQRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBrokerOwnedNamespacesOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 记录条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 该broker上的命名空间列表

		OwnedNamespaceSet []*OwnedNamespaceSetOpt `json:"OwnedNamespaceSet,omitempty" name:"OwnedNamespaceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBrokerOwnedNamespacesOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBrokerOwnedNamespacesOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBatchRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 批量绑定数据

		BatchRoleBindData []*RoleBindData `json:"BatchRoleBindData,omitempty" name:"BatchRoleBindData"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBatchRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBatchRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindCmqDeadLetterRequest struct {
	*tchttp.BaseRequest

	// 死信策略源队列名称，调用本接口会清空该队列的死信队列策略。

	SourceQueueName *string `json:"SourceQueueName,omitempty" name:"SourceQueueName"`
}

func (r *UnbindCmqDeadLetterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindCmqDeadLetterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternalRocketMQInstance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例所属用户APPID

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 实例所属用户UIN

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 实例付费状态，OK正常，ISOLATED隔离中，DESTROYED已销毁

	PayStatus *string `json:"PayStatus,omitempty" name:"PayStatus"`
	// 实例所属物理集群

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 最大可创建命名空间数量

	MaxNamespaces *uint64 `json:"MaxNamespaces,omitempty" name:"MaxNamespaces"`
	// 最大可创建主题数量

	MaxTopics *uint64 `json:"MaxTopics,omitempty" name:"MaxTopics"`
	// 最大可创建消费组数量

	MaxGroups *uint64 `json:"MaxGroups,omitempty" name:"MaxGroups"`
	// 单节点主题最大可调整队列数

	MaxQueuesPerTopic *uint64 `json:"MaxQueuesPerTopic,omitempty" name:"MaxQueuesPerTopic"`
	// 限流值

	RateLimit *uint64 `json:"RateLimit,omitempty" name:"RateLimit"`
	// 是否专享实例

	IsVIP *bool `json:"IsVIP,omitempty" name:"IsVIP"`
}

type DescribeNodeLatestMetricsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标值

		Metrics *float64 `json:"Metrics,omitempty" name:"Metrics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeLatestMetricsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeLatestMetricsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQGroupsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 限制条数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 主题名称，输入此参数可查询该主题下所有的订阅组

	FilterTopic *string `json:"FilterTopic,omitempty" name:"FilterTopic"`
	// 按消费组名称查询消费组，支持模糊查询

	FilterGroup *string `json:"FilterGroup,omitempty" name:"FilterGroup"`
	// 按照指定字段排序，可选值为tps，accumulative

	SortedBy *string `json:"SortedBy,omitempty" name:"SortedBy"`
	// 按升序或降序排列，可选值为asc，desc

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
	// 订阅组名称，指定此参数后将只返回该订阅组信息

	FilterOneGroup *string `json:"FilterOneGroup,omitempty" name:"FilterOneGroup"`
	// group类型

	Types []*string `json:"Types,omitempty" name:"Types"`
}

func (r *DescribeRocketMQGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateTenantQuotaOptRequest struct {
	*tchttp.BaseRequest

	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 命名空间配额

	MaxNamespaces *int64 `json:"MaxNamespaces,omitempty" name:"MaxNamespaces"`
	// Topic配额

	MaxTopics *int64 `json:"MaxTopics,omitempty" name:"MaxTopics"`
	// Exchange配额

	MaxExchanges *int64 `json:"MaxExchanges,omitempty" name:"MaxExchanges"`
	// Queue配额

	MaxQueues *int64 `json:"MaxQueues,omitempty" name:"MaxQueues"`
	// Topic分区配额

	MaxPartitions *int64 `json:"MaxPartitions,omitempty" name:"MaxPartitions"`
	// 命名空间最大堆积，byte为单

	MaxMsgBacklog *int64 `json:"MaxMsgBacklog,omitempty" name:"MaxMsgBacklog"`
	// 最长保留时间，毫秒为单位

	MaxRetention *int64 `json:"MaxRetention,omitempty" name:"MaxRetention"`
	// 客户APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
}

func (r *UpdateTenantQuotaOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateTenantQuotaOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		ClusterList []*RocketMQClusterDetail `json:"ClusterList,omitempty" name:"ClusterList"`
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQVipInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未分页的总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例信息列表

		Instances []*RocketMQVipInstance `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQVipInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQVipInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalRocketMQInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 分页后的实例信息列表

		Instances []*InternalRocketMQInstance `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternalRocketMQInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalRocketMQInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAMQPClusterRequest struct {
	*tchttp.BaseRequest

	// 3-64个字符，只能包含字母、数字、“-”及“_”

	Name *string `json:"Name,omitempty" name:"Name"`
	// 集群描述，128个字符以内

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateAMQPClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAMQPClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNamespacesOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 客户APPID

	CustomerAppId *string `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 租户id

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 查询限制条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 命名空间ID，获取单条记录时使用

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
}

func (r *DescribeNamespacesOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespacesOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCmqQueueRequest struct {
	*tchttp.BaseRequest

	// 队列名字，在单个地域同一帐号下唯一。队列名称是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)。

	QueueName *string `json:"QueueName,omitempty" name:"QueueName"`
}

func (r *DeleteCmqQueueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCmqQueueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigDictItemOpt struct {

	// 配置项Key

	ItemKey *string `json:"ItemKey,omitempty" name:"ItemKey"`
	// 配置项默认值

	ItemDefaultValue *string `json:"ItemDefaultValue,omitempty" name:"ItemDefaultValue"`
	// 配置项描述

	ItemDesc *string `json:"ItemDesc,omitempty" name:"ItemDesc"`
	// 是否动态生成

	IsDynamic *bool `json:"IsDynamic,omitempty" name:"IsDynamic"`
}

type TenantTopListOpt struct {

	// 租户ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 租户名

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// 生产速率

	RateIn *float64 `json:"RateIn,omitempty" name:"RateIn"`
	// 消费速率

	RateOut *float64 `json:"RateOut,omitempty" name:"RateOut"`
	// 入吞吐量

	ThroughputIn *float64 `json:"ThroughputIn,omitempty" name:"ThroughputIn"`
	// 出吞吐量

	ThroughputOut *float64 `json:"ThroughputOut,omitempty" name:"ThroughputOut"`
	// 存储大小，带单位，后端自动转换

	Storage *string `json:"Storage,omitempty" name:"Storage"`
}

type DescribeBookieDiskListOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据盘总数

		DataDiskCount *int64 `json:"DataDiskCount,omitempty" name:"DataDiskCount"`
		// 数据盘列表

		DataDiskList []*NodeDataDiskOpt `json:"DataDiskList,omitempty" name:"DataDiskList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBookieDiskListOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBookieDiskListOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMessagesRequest struct {
	*tchttp.BaseRequest

	// Token 是用来做鉴权使用的，可以不填，系统会自动获取

	StringToken *string `json:"StringToken,omitempty" name:"StringToken"`
	// 消息要发送的topic的名字, 这里尽量需要使用topic的全路径，即：tenant/namespace/topic。如果不指定，默认使用的是：public/default

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 要发送的消息的内容

	Payload *string `json:"Payload,omitempty" name:"Payload"`
	// 设置 producer 的名字，要求全局唯一，用户不配置，系统会随机生成

	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`
	// 设置消息发送的超时时间，默认为30s

	SendTimeout *int64 `json:"SendTimeout,omitempty" name:"SendTimeout"`
	// 内存中缓存的最大的生产消息的数量，默认为1000条

	MaxPendingMessages *int64 `json:"MaxPendingMessages,omitempty" name:"MaxPendingMessages"`
}

func (r *SendMessagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendMessagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternalNamespace struct {

	// 命名空间名称

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 客户UIN

	CustomerUin *string `json:"CustomerUin,omitempty" name:"CustomerUin"`
	// 客户的APPID

	CustomerAppId *uint64 `json:"CustomerAppId,omitempty" name:"CustomerAppId"`
	// 虚拟集群ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 虚拟集群名称

	TenantName *string `json:"TenantName,omitempty" name:"TenantName"`
	// 创建时间，毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 创建时间，毫秒为单位

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DescribeRocketMQClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRocketMQClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyAntiAffinityGroupOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyAntiAffinityGroupOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyAntiAffinityGroupOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentRolesRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 角色名称数组。

	RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`
	// 必填字段，集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteEnvironmentRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteEnvironmentRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMessageTTLOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyMessageTTLOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMessageTTLOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalRocketMQGroupsRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 如指定主题，则只查询订阅该主题的消费组

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 查询过滤器，支持以下过滤条件：
	// Name，按照消费组名称模糊查询

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询起始位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 查询最大数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInternalRocketMQGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalRocketMQGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProtocolSetOpt struct {

	// 协议ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 协议名

	ProtocolName *string `json:"ProtocolName,omitempty" name:"ProtocolName"`
}

type DescribeAMQPVHostConnectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总记录数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// vhost客户端连接列表

		Connections []*AMQPVHostConnection `json:"Connections,omitempty" name:"Connections"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAMQPVHostConnectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAMQPVHostConnectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTopicInternalStatsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Topic内部状态

		TopicInternalStatsSet *PartitionedTopicInternalStatsSetOpt `json:"TopicInternalStatsSet,omitempty" name:"TopicInternalStatsSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTopicInternalStatsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTopicInternalStatsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// 集群Id，传入需要删除的集群Id。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 被删除的主题数组。

		TopicSets []*TopicRecord `json:"TopicSets,omitempty" name:"TopicSets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternalStatsSetOpt struct {

	// 新增的entry总数

	EntriesAddedCounter *int64 `json:"EntriesAddedCounter,omitempty" name:"EntriesAddedCounter"`
	// 当前entry数量

	NumberOfEntries *int64 `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`
	// 总大小

	TotalSize *int64 `json:"TotalSize,omitempty" name:"TotalSize"`
	// 当前ledger中的entry数量

	CurrentLedgerEntries *int64 `json:"CurrentLedgerEntries,omitempty" name:"CurrentLedgerEntries"`
	// 当前ledger大小

	CurrentLedgerSize *int64 `json:"CurrentLedgerSize,omitempty" name:"CurrentLedgerSize"`
	// 最近一次创建leder时间

	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitempty" name:"LastLedgerCreatedTimestamp"`
	// 等待中的游标数量

	WaitingCursorsCount *int64 `json:"WaitingCursorsCount,omitempty" name:"WaitingCursorsCount"`
	// 等待被持久化的Entry数

	PendingAddEntriesCount *int64 `json:"PendingAddEntriesCount,omitempty" name:"PendingAddEntriesCount"`
	// 最后持久化成功的Entry

	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitempty" name:"LastConfirmedEntry"`
	// 状态

	State *string `json:"State,omitempty" name:"State"`
	// ledger状态列表

	Ledgers []*LedgerInfoSetOpt `json:"Ledgers,omitempty" name:"Ledgers"`
	// 游标状态列表

	CursorStatsSets []*CursorStatsSetOpt `json:"CursorStatsSets,omitempty" name:"CursorStatsSets"`
	// 待确认：CompactedLedger

	CompactedLedger *LedgerInfoSetOpt `json:"CompactedLedger,omitempty" name:"CompactedLedger"`
	// 分区名

	Partition *string `json:"Partition,omitempty" name:"Partition"`
}

type DeleteRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间名称

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 消费组名称

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteRocketMQGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRocketMQTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRolesRequest struct {
	*tchttp.BaseRequest

	// 角色名称数组。

	RoleNames []*string `json:"RoleNames,omitempty" name:"RoleNames"`
	// 必填字段，集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteRolesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRolesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 命名空间列表

		Namespaces []*InternalNamespace `json:"Namespaces,omitempty" name:"Namespaces"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternalNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMaxConsumerPerTopicOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyMaxConsumerPerTopicOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMaxConsumerPerTopicOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户Vpc Id。

		UniqueVpcId *string `json:"UniqueVpcId,omitempty" name:"UniqueVpcId"`
		// 租户Vpc子网Id。

		UniqueSubnetId *string `json:"UniqueSubnetId,omitempty" name:"UniqueSubnetId"`
		// 路由Id。

		RouterId *string `json:"RouterId,omitempty" name:"RouterId"`
		// Vpc的Ip。

		Ip *string `json:"Ip,omitempty" name:"Ip"`
		// Vpc的Port。

		Port *uint64 `json:"Port,omitempty" name:"Port"`
		// 说明，128个字符以内。

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCmqTopicsRequest struct {
	*tchttp.BaseRequest

	// 分页时本页获取队列列表的起始位置。如果填写了该值，必须也要填写 limit 。该值缺省时，后台取默认值 0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页时本页获取队列的个数，如果不传递该参数，则该参数默认为20，最大值为50。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 根据TopicName进行模糊搜索

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// CMQ 主题名称列表过滤

	TopicNameList []*string `json:"TopicNameList,omitempty" name:"TopicNameList"`
	// 标签过滤查找时，需要设置为 true

	IsTagFilter *bool `json:"IsTagFilter,omitempty" name:"IsTagFilter"`
}

func (r *DescribeCmqTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCmqTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQTopicsByGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 主题列表

		Topics []*string `json:"Topics,omitempty" name:"Topics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQTopicsByGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQTopicsByGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicReplicationClustersRequest struct {
	*tchttp.BaseRequest

	// Pulsar 集群的ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 环境（命名空间）名称

	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
	// 主题名

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeTopicReplicationClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicReplicationClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQEnvironmentRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 环境（命名空间）名称。

		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
		// 角色名称。

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 授权项。

		Permissions []*string `json:"Permissions,omitempty" name:"Permissions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQEnvironmentRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQEnvironmentRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInternalRocketMQInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInternalRocketMQInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInternalRocketMQInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRocketMQCreateQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 租户总共可使用集群数量

		MaxClusterNum *uint64 `json:"MaxClusterNum,omitempty" name:"MaxClusterNum"`
		// 租户已创建集群数量

		UsedClusterNum *uint64 `json:"UsedClusterNum,omitempty" name:"UsedClusterNum"`
		// Topic容量

		TopicCapacity *uint64 `json:"TopicCapacity,omitempty" name:"TopicCapacity"`
		// Group容量

		GroupCapacity *uint64 `json:"GroupCapacity,omitempty" name:"GroupCapacity"`
		// 单命名空间TPS

		TpsPerNamespace *uint64 `json:"TpsPerNamespace,omitempty" name:"TpsPerNamespace"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRocketMQCreateQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRocketMQCreateQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Pulsar 集群的ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetRocketMQConsumerOffSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetRocketMQConsumerOffSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetRocketMQConsumerOffSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterComponentMetricsOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// barad视图名，取值：
	// aggregation_metrics：broker监控指标
	// bookeeper_metrics：bookie监控指标
	// zookeeper_request_metrics：zookeeper请求监控指标
	// zookeeper_server_metrics：zookeeper服务端监控指标

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 指标名，通过DescribeClusterComponentMetricsListOpt接口获取

	MetricsName *string `json:"MetricsName,omitempty" name:"MetricsName"`
	// 监控指标开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 监控指标结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 监控指标间隔，单位为秒

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeClusterComponentMetricsOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterComponentMetricsOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyDispatchRatePerSubscriptionOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyDispatchRatePerSubscriptionOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyDispatchRatePerSubscriptionOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LedgerInfoSetOpt struct {

	// Ledger ID

	LedgerId *float64 `json:"LedgerId,omitempty" name:"LedgerId"`
	// entry数量

	Entries *float64 `json:"Entries,omitempty" name:"Entries"`
	// Ledger大小

	Size *int64 `json:"Size,omitempty" name:"Size"`
	// 是否卸载

	Offloaded *bool `json:"Offloaded,omitempty" name:"Offloaded"`
}

type RocketMQNamespace struct {

	// 命名空间名称，3-64个字符，只能包含字母、数字、“-”及“_”

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 已废弃，未消费消息的保留时间，以毫秒单位，范围60秒到15天

	Ttl *uint64 `json:"Ttl,omitempty" name:"Ttl"`
	// 消息持久化后保留的时间，以毫秒单位

	RetentionTime *uint64 `json:"RetentionTime,omitempty" name:"RetentionTime"`
	// 说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 公网接入点地址

	PublicEndpoint *string `json:"PublicEndpoint,omitempty" name:"PublicEndpoint"`
	// VPC接入点地址

	VpcEndpoint *string `json:"VpcEndpoint,omitempty" name:"VpcEndpoint"`
	// 内部接入点地址

	InternalEndpoint *string `json:"InternalEndpoint,omitempty" name:"InternalEndpoint"`
}

type CreateTemplateItemsOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 操作是否成功

		IsSuccess *bool `json:"IsSuccess,omitempty" name:"IsSuccess"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTemplateItemsOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTemplateItemsOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCmqTopicAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCmqTopicAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCmqTopicAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分区数

		Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`
		// 备注，128字符以内。

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RocketMQMigrationSubTask struct {

	// 主题或消费组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 执行结果码：
	// 0，成功
	// 1，已存在，跳过
	// 2，额度已满
	// 3，名称不符合规范
	// 4，系统错误
	// 5，不合法的分区数

	Code *int64 `json:"Code,omitempty" name:"Code"`
	// 元数据类型：
	// 0，主题
	// 1，消费组

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type DescribeNamespacePolicyOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
}

func (r *DescribeNamespacePolicyOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNamespacePolicyOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RebootInstancesOptRequest struct {
	*tchttp.BaseRequest

	// 实例ID列表

	InstanceIdList []*string `json:"InstanceIdList,omitempty" name:"InstanceIdList"`
}

func (r *RebootInstancesOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootInstancesOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateConfigOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 模块名称

	ConfigModule *string `json:"ConfigModule,omitempty" name:"ConfigModule"`
	// 根路径

	RootPath *string `json:"RootPath,omitempty" name:"RootPath"`
	// 数据存放路径

	DataPath *string `json:"DataPath,omitempty" name:"DataPath"`
	// 配置存放路径

	ConfigPath *string `json:"ConfigPath,omitempty" name:"ConfigPath"`
	// 配置项（BASE64编码后）

	ConfigItems *string `json:"ConfigItems,omitempty" name:"ConfigItems"`
	// 配置版本

	ConfigVersion *string `json:"ConfigVersion,omitempty" name:"ConfigVersion"`
}

func (r *UpdateConfigOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateConfigOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterSetOpt struct {

	// 集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 底层资源类型

	ClusterResource *string `json:"ClusterResource,omitempty" name:"ClusterResource"`
	// 集群状态

	ClusterState *string `json:"ClusterState,omitempty" name:"ClusterState"`
	// broker数量

	BrokerCount *int64 `json:"BrokerCount,omitempty" name:"BrokerCount"`
	// 可用broker数量

	AvailableBrokerCount *int64 `json:"AvailableBrokerCount,omitempty" name:"AvailableBrokerCount"`
	// bookie数量

	BookieCount *int64 `json:"BookieCount,omitempty" name:"BookieCount"`
	// 可用bookie数量

	AvailableBookieCount *int64 `json:"AvailableBookieCount,omitempty" name:"AvailableBookieCount"`
	// zk数量

	ZkCount *int64 `json:"ZkCount,omitempty" name:"ZkCount"`
	// 可用zk数量

	AvailableZKCount *int64 `json:"AvailableZKCount,omitempty" name:"AvailableZKCount"`
	// 虚拟集群数量

	TenantCount *int64 `json:"TenantCount,omitempty" name:"TenantCount"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 备注

	Comment *string `json:"Comment,omitempty" name:"Comment"`
}

type DeleteRocketMQClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRocketMQClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRocketMQClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyOffloadThresholdOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyOffloadThresholdOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyOffloadThresholdOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRoleRequest struct {
	*tchttp.BaseRequest

	// 角色名称，不支持中字以及除了短线和下划线外的特殊字符且长度必须大于0且小等于32。

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 备注说明，长度必须大等于0且小等于128。

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 必填字段，集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 批量操作环境

	EnvironmentRoleSets []*EnvironmentRoleSet `json:"EnvironmentRoleSets,omitempty" name:"EnvironmentRoleSets"`
}

func (r *CreateRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CursorStatsSetOpt struct {

	// 订阅名

	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
	// 标记删除位置

	MarkDeletePosition *string `json:"MarkDeletePosition,omitempty" name:"MarkDeletePosition"`
	// 读取位置

	ReadPosition *string `json:"ReadPosition,omitempty" name:"ReadPosition"`
	// 等待中的读请求

	WaitingReadOp *bool `json:"WaitingReadOp,omitempty" name:"WaitingReadOp"`
	// 等待中的读请求数

	PendingReadOps *int64 `json:"PendingReadOps,omitempty" name:"PendingReadOps"`
	// 已被消费的消息总数

	MessagesConsumedCounter *int64 `json:"MessagesConsumedCounter,omitempty" name:"MessagesConsumedCounter"`
	// 游标Ledger

	CursorLedger *int64 `json:"CursorLedger,omitempty" name:"CursorLedger"`
	// 游标Ledger中的最后一个Entry

	CursorLedgerLastEntry *int64 `json:"CursorLedgerLastEntry,omitempty" name:"CursorLedgerLastEntry"`
	// 单独确认的消息

	IndividuallyDeletedMessages *string `json:"IndividuallyDeletedMessages,omitempty" name:"IndividuallyDeletedMessages"`
	// Ledger最后切换时间

	LastLedgerSwitchTimestamp *string `json:"LastLedgerSwitchTimestamp,omitempty" name:"LastLedgerSwitchTimestamp"`
	// 游标状态

	State *string `json:"State,omitempty" name:"State"`
	// 未确认消息跨越的Entry数

	NumberOfEntriesSinceFirstNotAckedMessage *int64 `json:"NumberOfEntriesSinceFirstNotAckedMessage,omitempty" name:"NumberOfEntriesSinceFirstNotAckedMessage"`
	// 未删除的消息范围总数

	TotalNonContiguousDeletedMessagesRange *int64 `json:"TotalNonContiguousDeletedMessagesRange,omitempty" name:"TotalNonContiguousDeletedMessagesRange"`
}

type DescribeConfigDictOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模块名称

		ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
		// 配置字典表信息

		DictItems []*ConfigDictItemOpt `json:"DictItems,omitempty" name:"DictItems"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigDictOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigDictOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutClusterOnlineOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// Admin url

	AdminUrl *string `json:"AdminUrl,omitempty" name:"AdminUrl"`
	// 物理集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 是否内部集群

	IsInternal *bool `json:"IsInternal,omitempty" name:"IsInternal"`
}

func (r *PutClusterOnlineOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutClusterOnlineOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRocketMQRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 角色名称

		RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
		// 角色token

		Token *string `json:"Token,omitempty" name:"Token"`
		// 备注说明

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRocketMQRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRocketMQRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitializeClusterConfigOptRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 是否覆盖

	IsOverWrite *bool `json:"IsOverWrite,omitempty" name:"IsOverWrite"`
	// 0-入门 1- 基础 2- 标准 3- 高性能

	TemplateType *uint64 `json:"TemplateType,omitempty" name:"TemplateType"`
}

func (r *InitializeClusterConfigOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitializeClusterConfigOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyMessageRetentionOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdatePolicyMessageRetentionOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyMessageRetentionOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeSetOpt struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 节点地址

	Address *string `json:"Address,omitempty" name:"Address"`
	// 实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 可用区

	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	// 节点状态

	NodeState *uint64 `json:"NodeState,omitempty" name:"NodeState"`
	// cpu使用率

	CpuUsage *float64 `json:"CpuUsage,omitempty" name:"CpuUsage"`
	// cpu核数

	CpuTotal *int64 `json:"CpuTotal,omitempty" name:"CpuTotal"`
	// 内存使用率

	MemUsage *float64 `json:"MemUsage,omitempty" name:"MemUsage"`
	// 内存大小（G）

	MemTotal *int64 `json:"MemTotal,omitempty" name:"MemTotal"`
	// 内网入带宽

	InBand *float64 `json:"InBand,omitempty" name:"InBand"`
	// 内网出带宽

	OutBand *float64 `json:"OutBand,omitempty" name:"OutBand"`
}

type ModifyRocketMQGroupRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	NamespaceId *string `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 消费组名称

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 说明信息，最长128个字符

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否开启消费

	ReadEnable *bool `json:"ReadEnable,omitempty" name:"ReadEnable"`
	// 是否开启广播消费

	BroadcastEnable *bool `json:"BroadcastEnable,omitempty" name:"BroadcastEnable"`
	// 最大重试次数

	RetryMaxTimes *uint64 `json:"RetryMaxTimes,omitempty" name:"RetryMaxTimes"`
}

func (r *ModifyRocketMQGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRocketMQGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicySchemaValidationEnforcedRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 是否开启强校验

	SchemaValidationEnforced *bool `json:"SchemaValidationEnforced,omitempty" name:"SchemaValidationEnforced"`
}

func (r *UpdatePolicySchemaValidationEnforcedRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicySchemaValidationEnforcedRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInternalRocketMQGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 分页后的消费组列表

		Groups []*RocketMQGroup `json:"Groups,omitempty" name:"Groups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInternalRocketMQGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInternalRocketMQGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllTemplateModuleOptResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置模块列表

		ConfigType []*ConfigTypeInfoOpt `json:"ConfigType,omitempty" name:"ConfigType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllTemplateModuleOptResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllTemplateModuleOptResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// 起始下标，不填默认为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// Pulsar 集群的名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群ID列表过滤

	ClusterIdList []*string `json:"ClusterIdList,omitempty" name:"ClusterIdList"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdatePolicyDeduplicationEnabledOptRequest struct {
	*tchttp.BaseRequest

	// 物理集群名

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 虚拟集群（租户）ID

	TenantId *string `json:"TenantId,omitempty" name:"TenantId"`
	// 命名空间名

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// 重复数据清除开关

	DeduplicationEnabled *bool `json:"DeduplicationEnabled,omitempty" name:"DeduplicationEnabled"`
}

func (r *UpdatePolicyDeduplicationEnabledOptRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdatePolicyDeduplicationEnabledOptRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeDataDiskOpt struct {

	// 数据盘ID

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 数据盘大小（GB）

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 数据盘类型：
	// CLOUD_BASIC：表示普通云硬盘
	// CLOUD_PREMIUM：表示高性能云硬盘
	// CLOUD_SSD：表示SSD云硬盘
	// CLOUD_HSSD：表示增强型SSD云硬盘
	// CLOUD_TSSD：表示极速型SSD云硬盘。

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 数据盘可用区

	AvailableZone *string `json:"AvailableZone,omitempty" name:"AvailableZone"`
	// 数据盘使用率（取值范围0~1）

	DiskUsage *float64 `json:"DiskUsage,omitempty" name:"DiskUsage"`
}

type RouteRecord struct {

	// 私有网络

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 路由id

	RouterId *string `json:"RouterId,omitempty" name:"RouterId"`
	// vip地址

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// vport记录

	Vport *uint64 `json:"Vport,omitempty" name:"Vport"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 路由类型 1: 公网 2：vpc 4：支撑网

	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`
	// 路由Id（int 类型）

	Id *uint64 `json:"Id,omitempty" name:"Id"`

	// EncryptType bool    `json:"EncryptType,omitempty" name:"EncryptType"`
	EncryptType *uint64 `json:"EncryptType,omitempty" name:"EncryptType"`
}

type DescribeRabbitMQNodeListRequest struct {
	*tchttp.BaseRequest

	// rabbitmq集群ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 一页限制

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 模糊搜索节点名字

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 过滤参数的名字和数值
	// 现在只有一个nodeStatus
	// running/down
	// 数组类型，兼容后续添加过滤参数
	//

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 按指定元素排序，现在只有2个
	// cpuUsage/diskUsage

	SortElement *string `json:"SortElement,omitempty" name:"SortElement"`
	// 升序/降序
	// ascend/descend

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
}

func (r *DescribeRabbitMQNodeListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQNodeListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RabbitMQPrivateNode struct {

	// 节点状态

	NodeStatus *string `json:"NodeStatus,omitempty" name:"NodeStatus"`
	// CPU使用率

	CPUUsage *string `json:"CPUUsage,omitempty" name:"CPUUsage"`
	// 内存使用情况，单位MB

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 磁盘使用率

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
	// Rabbitmq的Erlang进程数

	ProcessNumber *uint64 `json:"ProcessNumber,omitempty" name:"ProcessNumber"`
	// 节点名字

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type DescribeRabbitMQNodeListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群列表数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群列表

		NodeList []*RabbitMQPrivateNode `json:"NodeList,omitempty" name:"NodeList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRabbitMQNodeListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQNodeListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRabbitMQVipInstancesRequest struct {
	*tchttp.BaseRequest

	// 查询条件过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询数目上限，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询起始位置

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeRabbitMQVipInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQVipInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRabbitMQVipInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例信息列表

		Instances []*RabbitMQVipInstance `json:"Instances,omitempty" name:"Instances"`
		// 未分页的总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRabbitMQVipInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQVipInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RabbitMQVipInstance struct {

	// 峰值带宽，Mbps为单位

	MaxBandWidth *uint64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`
	// 自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费)，&nbsp;1表示自动续费，2表示明确不自动续费(用户设置)

	AutoRenewFlag *uint64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 底层容器资源标记

	Container *bool `json:"Container,omitempty" name:"Container"`
	// 实例类型，0&nbsp;专享版、1&nbsp;Serverless&nbsp;版

	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例版本

	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`
	// 节点数量

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 实例配置规格名称

	ConfigDisplay *string `json:"ConfigDisplay,omitempty" name:"ConfigDisplay"`
	// VPC&nbsp;接入点列表

	Vpcs []*VpcEndpointInfo `json:"Vpcs,omitempty" name:"Vpcs"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4&nbsp;-&nbsp;异常,&nbsp;5&nbsp;-&nbsp;发货失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 存储容量，GB为单位

	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 是否开启trace

	TraceFlag *bool `json:"TraceFlag,omitempty" name:"TraceFlag"`
	// 创建时间，毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例到期时间，毫秒为单位

	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 可用区信息

	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
	// 实例配置ID

	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`
	// 集群异常。

	ExceptionInformation *string `json:"ExceptionInformation,omitempty" name:"ExceptionInformation"`
	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4&nbsp;-&nbsp;异常,&nbsp;5&nbsp;-&nbsp;发货失败
	// 为了和计费区分开，额外开启一个状态位，用于显示。

	ClusterStatus *int64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 公网接入点

	PublicAccessEndpoint *string `json:"PublicAccessEndpoint,omitempty" name:"PublicAccessEndpoint"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 峰值TPS

	MaxTps *uint64 `json:"MaxTps,omitempty" name:"MaxTps"`
	// 0-后付费，1-预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
}

type VpcEndpointInfo struct {

	// vpc的id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// vpc接入点信息

	VpcEndpoint *string `json:"VpcEndpoint,omitempty" name:"VpcEndpoint"`
	// vpc接入点状态
	// OFF/ON/CREATING/DELETING

	VpcDataStreamEndpointStatus *string `json:"VpcDataStreamEndpointStatus,omitempty" name:"VpcDataStreamEndpointStatus"`
}

type CreateRabbitMQUserRequest struct {
	*tchttp.BaseRequest

	// 用户标签，用于决定改用户访问RabbitMQ&nbsp;Management的权限范围
	// management：普通控制台用户，monitoring：管理型控制台用户，其他值：非控制台用户

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 该用户的最大连接数，不填写则不限制

	MaxConnections *int64 `json:"MaxConnections,omitempty" name:"MaxConnections"`
	// 该用户的最大channel数，不填写则不限制

	MaxChannels *int64 `json:"MaxChannels,omitempty" name:"MaxChannels"`
	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名，登录时使用

	User *string `json:"User,omitempty" name:"User"`
	// 密码，登录时使用

	Password *string `json:"Password,omitempty" name:"Password"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateRabbitMQUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRabbitMQUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRabbitMQUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户名，登录时使用

		User *string `json:"User,omitempty" name:"User"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRabbitMQUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRabbitMQUserRequest struct {
	*tchttp.BaseRequest

	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名，登录时使用

	User *string `json:"User,omitempty" name:"User"`
	// 密码，登录时使用

	Password *string `json:"Password,omitempty" name:"Password"`
	// 描述，不传则不修改

	Description *string `json:"Description,omitempty" name:"Description"`
	// 用户标签，用于决定改用户访问RabbitMQ&nbsp;Management的权限范围，不传则不修改

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 该用户的最大连接数，不传则不修改

	MaxConnections *int64 `json:"MaxConnections,omitempty" name:"MaxConnections"`
	// 该用户的最大channel数，不传则不修改

	MaxChannels *int64 `json:"MaxChannels,omitempty" name:"MaxChannels"`
}

func (r *ModifyRabbitMQUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRabbitMQUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRabbitMQUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRabbitMQUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRabbitMQUserRequest struct {
	*tchttp.BaseRequest

	// 用户标签，根据标签过滤列表

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名检索，支持前缀匹配，后缀匹配

	SearchUser *string `json:"SearchUser,omitempty" name:"SearchUser"`
	// 分页Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 用户名，精确查询

	User *string `json:"User,omitempty" name:"User"`
}

func (r *DescribeRabbitMQUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRabbitMQUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 当前已创建的RabbitMQ用户列表

		RabbitMQUserList []*RabbitMQUser `json:"RabbitMQUserList,omitempty" name:"RabbitMQUserList"`
		// 返回的User数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRabbitMQUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RabbitMQUser struct {

	// 用户最后修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 单个用户最大可用通道数

	MaxChannels *int64 `json:"MaxChannels,omitempty" name:"MaxChannels"`
	// 用户名，登录时使用

	User *string `json:"User,omitempty" name:"User"`
	// 密码，登录时使用

	Password *string `json:"Password,omitempty" name:"Password"`
	// 用户标签，用于决定改用户访问RabbitMQ&nbsp;Management的权限范围

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 用户类型，System：系统创建，User：用户创建

	Type *string `json:"Type,omitempty" name:"Type"`
	// 单个用户最大可用连接数

	MaxConnections *int64 `json:"MaxConnections,omitempty" name:"MaxConnections"`
	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 用户创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DeleteRabbitMQUserRequest struct {
	*tchttp.BaseRequest

	// 用户名，登录时使用

	User *string `json:"User,omitempty" name:"User"`
	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteRabbitMQUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRabbitMQUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRabbitMQUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRabbitMQUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRabbitMQUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeRabbitMQVipInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息

		ClusterInfo *RabbitMQClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`
		// 集群规格信息

		ClusterSpecInfo *RabbitMQClusterSpecInfo `json:"ClusterSpecInfo,omitempty" name:"ClusterSpecInfo"`
		// 集群访问

		ClusterNetInfo *RabbitMQClusterAccessInfo `json:"ClusterNetInfo,omitempty" name:"ClusterNetInfo"`
		// 集群白名单

		ClusterWhiteListInfo *RabbitMQClusterWhiteListInfo `json:"ClusterWhiteListInfo,omitempty" name:"ClusterWhiteListInfo"`
		// vhost配额信息

		VirtualHostQuota *VirtualHostQuota `json:"VirtualHostQuota,omitempty" name:"VirtualHostQuota"`
		// exchange配额信息

		ExchangeQuota *ExchangeQuota `json:"ExchangeQuota,omitempty" name:"ExchangeQuota"`
		// queue配额信息

		QueueQuota *QueueQuota `json:"QueueQuota,omitempty" name:"QueueQuota"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRabbitMQVipInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RabbitMQClusterInfo struct {

	// 集群说明信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// VPC及网络信息

	Vpcs []*VpcEndpointInfo `json:"Vpcs,omitempty" name:"Vpcs"`
	// 可用区信息

	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
	// 虚拟主机数量

	VirtualHostNumber *int64 `json:"VirtualHostNumber,omitempty" name:"VirtualHostNumber"`
	// 队列数量

	QueueNumber *int64 `json:"QueueNumber,omitempty" name:"QueueNumber"`
	// 过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// Connection数量

	ConnectionNumber *int64 `json:"ConnectionNumber,omitempty" name:"ConnectionNumber"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群异常信息

	ExceptionInformation *string `json:"ExceptionInformation,omitempty" name:"ExceptionInformation"`
	// 是否开启镜像队列策略。1表示开启，0表示没开启。

	MirrorQueuePolicyFlag *int64 `json:"MirrorQueuePolicyFlag,omitempty" name:"MirrorQueuePolicyFlag"`
	// 实例类型，0&nbsp;专享版、1&nbsp;Serverless&nbsp;版

	InstanceType *uint64 `json:"InstanceType,omitempty" name:"InstanceType"`
	// Consumer数量

	ConsumerNumber *int64 `json:"ConsumerNumber,omitempty" name:"ConsumerNumber"`
	// 是否开启trace

	TraceFlag *bool `json:"TraceFlag,omitempty" name:"TraceFlag"`
	// 计费模式，0-后付费，1-预付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 创建时间，毫秒为单位

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 堆积消息数&nbsp;单位：条

	MessageStackNumber *int64 `json:"MessageStackNumber,omitempty" name:"MessageStackNumber"`
	// Exchang数量

	ExchangeNumber *int64 `json:"ExchangeNumber,omitempty" name:"ExchangeNumber"`
	// 自动续费标记，0表示默认状态(用户未设置，即初始状态即手动续费)，&nbsp;1表示自动续费，2表示明确不自动续费(用户设置)

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 每秒消费消息数&nbsp;单位：条/秒

	MessageConsumeRate *float64 `json:"MessageConsumeRate,omitempty" name:"MessageConsumeRate"`
	// 集群版本信息

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 每秒生产消息数&nbsp;单位：条/秒

	MessagePublishRate *float64 `json:"MessagePublishRate,omitempty" name:"MessagePublishRate"`
	// 地域信息

	Region *string `json:"Region,omitempty" name:"Region"`
	// Channel数量

	ChannelNumber *int64 `json:"ChannelNumber,omitempty" name:"ChannelNumber"`
	// 实例状态，0表示创建中，1表示正常，2表示隔离中，3表示已销毁，4&nbsp;-&nbsp;异常,&nbsp;5&nbsp;-&nbsp;发货失败

	ClusterStatus *int64 `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

type RabbitMQClusterSpecInfo struct {

	// 集群规格名称

	SpecName *string `json:"SpecName,omitempty" name:"SpecName"`
	// 节点数量

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 峰值tps

	MaxTps *uint64 `json:"MaxTps,omitempty" name:"MaxTps"`
	// 峰值带宽。单位：mbps

	MaxBandWidth *uint64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`
	// 存储容量。单位：GB

	MaxStorage *uint64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 公网带宽tps。单位：Mbps

	PublicNetworkTps *uint64 `json:"PublicNetworkTps,omitempty" name:"PublicNetworkTps"`
}

type RabbitMQClusterAccessInfo struct {

	// 公网管控台开关状态，示例值，OFF/ON/CREATING/DELETING

	PublicWebConsoleSwitchStatus *string `json:"PublicWebConsoleSwitchStatus,omitempty" name:"PublicWebConsoleSwitchStatus"`
	// http://amqp-k3eb47gm.dashboard.rabbitmq.cq.public.tdmq.com:15672/
	// 公网域名接入点

	WebConsoleDomainEndpoint *string `json:"WebConsoleDomainEndpoint,omitempty" name:"WebConsoleDomainEndpoint"`
	// Vpc管控台开关状态，示例值，
	// OFF/ON/CREATING/DELETING

	VpcWebConsoleSwitchStatus *string `json:"VpcWebConsoleSwitchStatus,omitempty" name:"VpcWebConsoleSwitchStatus"`
	// 集群控制台登录用户名

	WebConsoleUsername *string `json:"WebConsoleUsername,omitempty" name:"WebConsoleUsername"`
	// 已废弃

	PublicAccessEndpointStatus *bool `json:"PublicAccessEndpointStatus,omitempty" name:"PublicAccessEndpointStatus"`
	// 已废弃

	VpcControlConsoleSwitchStatus *bool `json:"VpcControlConsoleSwitchStatus,omitempty" name:"VpcControlConsoleSwitchStatus"`
	// 集群控制台访问地址

	WebConsoleEndpoint *string `json:"WebConsoleEndpoint,omitempty" name:"WebConsoleEndpoint"`
	// 集群控制台登录密码

	WebConsolePassword *string `json:"WebConsolePassword,omitempty" name:"WebConsolePassword"`
	// Vpc管控台访问地址，示例值，http://1.1.1.1:15672

	VpcWebConsoleEndpoint *string `json:"VpcWebConsoleEndpoint,omitempty" name:"VpcWebConsoleEndpoint"`
	// Prometheus信息

	PrometheusEndpointInfo *PrometheusEndpointInfo `json:"PrometheusEndpointInfo,omitempty" name:"PrometheusEndpointInfo"`
	// 控制面所使用的VPC信息

	ControlPlaneEndpointInfo *VpcEndpointInfo `json:"ControlPlaneEndpointInfo,omitempty" name:"ControlPlaneEndpointInfo"`
	// 集群公网接入地址

	PublicAccessEndpoint *string `json:"PublicAccessEndpoint,omitempty" name:"PublicAccessEndpoint"`
	// 已废弃

	PublicControlConsoleSwitchStatus *bool `json:"PublicControlConsoleSwitchStatus,omitempty" name:"PublicControlConsoleSwitchStatus"`
	// 公网管控台开关状态，示例值，OFF/ON/CREATING/DELETING

	PublicDataStreamStatus *string `json:"PublicDataStreamStatus,omitempty" name:"PublicDataStreamStatus"`
}

type RabbitMQClusterWhiteListInfo struct {

	// 公网数据流白名单

	PublicDataStreamWhiteList *string `json:"PublicDataStreamWhiteList,omitempty" name:"PublicDataStreamWhiteList"`
	// 公网管控台白名单状态

	PublicControlConsoleWhiteListStatus *string `json:"PublicControlConsoleWhiteListStatus,omitempty" name:"PublicControlConsoleWhiteListStatus"`
	// 公网数据流白名单状态

	PublicDataStreamWhiteListStatus *string `json:"PublicDataStreamWhiteListStatus,omitempty" name:"PublicDataStreamWhiteListStatus"`
	// 废弃

	WhiteList *string `json:"WhiteList,omitempty" name:"WhiteList"`
	// 公网管控台白名单

	PublicControlConsoleWhiteList *string `json:"PublicControlConsoleWhiteList,omitempty" name:"PublicControlConsoleWhiteList"`
}

type VirtualHostQuota struct {

	// 允许创建最大vhost数

	MaxVirtualHost *int64 `json:"MaxVirtualHost,omitempty" name:"MaxVirtualHost"`
	// 已创建vhost数

	UsedVirtualHost *int64 `json:"UsedVirtualHost,omitempty" name:"UsedVirtualHost"`
}

type ExchangeQuota struct {

	// 可创建最大exchange数

	MaxExchange *int64 `json:"MaxExchange,omitempty" name:"MaxExchange"`
	// 已创建exchange数

	UsedExchange *int64 `json:"UsedExchange,omitempty" name:"UsedExchange"`
}

type QueueQuota struct {

	// 可创建最大Queue数

	MaxQueue *int64 `json:"MaxQueue,omitempty" name:"MaxQueue"`
	// 已创建Queue数

	UsedQueue *int64 `json:"UsedQueue,omitempty" name:"UsedQueue"`
}

type PrometheusEndpointInfo struct {

	// vpc信息

	VpcEndpointInfo *VpcEndpointInfo `json:"VpcEndpointInfo,omitempty" name:"VpcEndpointInfo"`
	// Prometheus开关的状态。

	PrometheusEndpointStatus *string `json:"PrometheusEndpointStatus,omitempty" name:"PrometheusEndpointStatus"`
	// prometheus信息

	VpcPrometheusEndpoint []*string `json:"VpcPrometheusEndpoint,omitempty" name:"VpcPrometheusEndpoint"`
	// 节点信息

	NodePrometheusAddress []*string `json:"NodePrometheusAddress,omitempty" name:"NodePrometheusAddress"`
}


type DeleteRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否国际站请求，默认&nbsp;false

	IsIntl *bool `json:"IsIntl,omitempty" name:"IsIntl"`
}

func (r *DeleteRabbitMQVipInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}


type DeleteRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 订单号Id

		TranId *string `json:"TranId,omitempty" name:"TranId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRabbitMQVipInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 单节点存储规格,不传默认为200G

	StorageSize *int64 `json:"StorageSize,omitempty" name:"StorageSize"`
	// 资源标签列表

	ResourceTags []*Tag `json:"ResourceTags,omitempty" name:"ResourceTags"`
	// 可用区

	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
	// 私有网络SubnetId

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 节点数量,多可用区最少为3节点。不传默认单可用区为1,多可用区为3

	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
	// 付费方式，0&nbsp;为后付费，即按量计费；1&nbsp;为预付费，即包年包月。默认包年包月

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 是否国际站请求，默认&nbsp;false

	IsIntl *bool `json:"IsIntl,omitempty" name:"IsIntl"`
	// 私有网络VpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 节点规格,基础型rabbit-vip-basic-1,标准型rabbit-vip-basic-2,高阶1型rabbit-vip-basic-3,高阶2型rabbit-vip-basic-4。不传默认为基础型

	NodeSpec *string `json:"NodeSpec,omitempty" name:"NodeSpec"`
	// 预付费使用。自动续费,不传默认为true

	AutoRenewFlag *bool `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 集群版本，不传默认为&nbsp;3.8.30，可选值为&nbsp;3.8.30&nbsp;和&nbsp;3.11.8

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 镜像队列,不传默认为false

	EnableCreateDefaultHaMirrorQueue *bool `json:"EnableCreateDefaultHaMirrorQueue,omitempty" name:"EnableCreateDefaultHaMirrorQueue"`
	// 购买时长,不传默认为1(月)

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 公网带宽大小，单位&nbsp;M

	Bandwidth *uint64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 是否打开公网接入，不传默认为false

	EnablePublicAccess *bool `json:"EnablePublicAccess,omitempty" name:"EnablePublicAccess"`
}

func (r *CreateRabbitMQVipInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 订单号Id

		TranId *string `json:"TranId,omitempty" name:"TranId"`
		// 实例Id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRabbitMQVipInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRabbitMQVipInstanceRequest struct {
	*tchttp.BaseRequest

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyRabbitMQVipInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRabbitMQVipInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRabbitMQVipInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRabbitMQVipInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRabbitMQVipInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest

	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// vhost名

	VirtualHost *string `json:"VirtualHost,omitempty" name:"VirtualHost"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 消息轨迹开关,true打开,false关闭,默认关闭

	TraceFlag *bool `json:"TraceFlag,omitempty" name:"TraceFlag"`
	// 是否创建镜像队列策略，默认值&nbsp;true

	MirrorQueuePolicyFlag *bool `json:"MirrorQueuePolicyFlag,omitempty" name:"MirrorQueuePolicyFlag"`
}

func (r *CreateRabbitMQVirtualHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// vhost名

		VirtualHost *string `json:"VirtualHost,omitempty" name:"VirtualHost"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRabbitMQVirtualHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest

	// 分页Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// search-virtual-host：vhost名称模糊查询，之前前缀和后缀匹配

	Filters *Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序依据的字段：
	// MessageHeapCount&nbsp;-&nbsp;消息堆积数；
	// MessageRateInOut&nbsp;-&nbsp;生产消费速率之和；
	// MessageRateIn&nbsp;-&nbsp;生产速率；
	// MessageRateOut&nbsp;-&nbsp;消费速率；

	SortElement *string `json:"SortElement,omitempty" name:"SortElement"`
	// 排序顺序，ascend&nbsp;或&nbsp;descend

	SortOrder *string `json:"SortOrder,omitempty" name:"SortOrder"`
	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// vhost名,不传则查询全部

	VirtualHost *string `json:"VirtualHost,omitempty" name:"VirtualHost"`
}

func (r *DescribeRabbitMQVirtualHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// vhost详情列表

		VirtualHostList []*RabbitMQVirtualHostInfo `json:"VirtualHostList,omitempty" name:"VirtualHostList"`
		// 返回vhost数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRabbitMQVirtualHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RabbitMQVirtualHostInfo struct {

	// vhost描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 修改时间

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 消息轨迹开关,true打开,false关闭

	TraceFlag *bool `json:"TraceFlag,omitempty" name:"TraceFlag"`
	// vhost状态，与原生控制台对应，有running、partial、stopped、unknown

	Status *string `json:"Status,omitempty" name:"Status"`
	// 消息堆积数

	MessageHeapCount *int64 `json:"MessageHeapCount,omitempty" name:"MessageHeapCount"`
	// 输入消息速率

	MessageRateIn *float64 `json:"MessageRateIn,omitempty" name:"MessageRateIn"`
	// 输出消息速率

	MessageRateOut *float64 `json:"MessageRateOut,omitempty" name:"MessageRateOut"`
	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// vhost名

	VirtualHost *string `json:"VirtualHost,omitempty" name:"VirtualHost"`
	// vhost标签

	Tags []*string `json:"Tags,omitempty" name:"Tags"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// vhost概览统计信息

	VirtualHostStatistics *RabbitMQVirtualHostStatistics `json:"VirtualHostStatistics,omitempty" name:"VirtualHostStatistics"`
	// 是否存在镜像队列策略，true&nbsp;为存在，false&nbsp;为不存

	MirrorQueuePolicyFlag *bool `json:"MirrorQueuePolicyFlag,omitempty" name:"MirrorQueuePolicyFlag"`
}

type RabbitMQVirtualHostStatistics struct {

	// 当前vhost的queue数量

	CurrentQueues *int64 `json:"CurrentQueues,omitempty" name:"CurrentQueues"`
	// 当前vhost的exchange数量

	CurrentExchanges *int64 `json:"CurrentExchanges,omitempty" name:"CurrentExchanges"`
	// 当前vhost的连接数量

	CurrentConnections *int64 `json:"CurrentConnections,omitempty" name:"CurrentConnections"`
	// 当前vhost的channel数量

	CurrentChannels *int64 `json:"CurrentChannels,omitempty" name:"CurrentChannels"`
	// 当前vhost的用户数量

	CurrentUsers *int64 `json:"CurrentUsers,omitempty" name:"CurrentUsers"`
}

type ModifyRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest

	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// vhost名

	VirtualHost *string `json:"VirtualHost,omitempty" name:"VirtualHost"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 消息轨迹开关,true打开,false关闭

	TraceFlag *bool `json:"TraceFlag,omitempty" name:"TraceFlag"`
}

func (r *ModifyRabbitMQVirtualHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRabbitMQVirtualHostResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRabbitMQVirtualHostResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRabbitMQVirtualHostResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRabbitMQVirtualHostRequest struct {
	*tchttp.BaseRequest

	// 集群实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// vhost名

	VirtualHost *string `json:"VirtualHost,omitempty" name:"VirtualHost"`
}

func (r *DeleteRabbitMQVirtualHostRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRabbitMQVirtualHostRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
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
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// ckafka 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// ckafka 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FetchMessageByOffsetRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 分区id

	Partition *uint64 `json:"Partition,omitempty" name:"Partition"`
	// 位点信息

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *FetchMessageByOffsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FetchMessageByOffsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {

	// ClusterId

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// ZoneId

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// Type

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// Flag

	Flag *int64 `json:"Flag,omitempty" name:"Flag"`
	// TypeDescriptor

	TypeDescriptor *string `json:"TypeDescriptor,omitempty" name:"TypeDescriptor"`
	// BindAppId

	BindAppId *int64 `json:"BindAppId,omitempty" name:"BindAppId"`
	// Status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Controller

	Controller *DescribeClustersController `json:"Controller,omitempty" name:"Controller"`
	// ClusterName

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// Ips

	Ips *string `json:"Ips,omitempty" name:"Ips"`
}

type GroupInfoTopics struct {

	// 分配的 topic 名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 分配的 partition 信息

	Partitions []*int64 `json:"Partitions,omitempty" name:"Partitions"`
}

type DeleteAclResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAclRequest struct {
	*tchttp.BaseRequest

	// 实例id信息

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Acl资源类型，(0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID)，当前只有TOPIC，其它字段用于后续兼容开源kafka的acl时使用

	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// Acl操作方式，(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTEN_WRITE)，当前ckafka只支持READ,WRITE，其它用于后续兼容开源kafka的acl时使用

	Operation *int64 `json:"Operation,omitempty" name:"Operation"`
	// 权限类型，(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)，当前ckakfa支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用

	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`
	// 默认为\*，表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持

	Host *string `json:"Host,omitempty" name:"Host"`
	// 用户列表，默认为*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户

	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

func (r *DeleteAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAclRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的消费分组信息

		Result *ConsumerGroupResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponses struct {

	// 集群信息

	Brokers []*DescribeClustersResponse `json:"Brokers,omitempty" name:"Brokers"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TradeCreateResponse struct {

	// 后付费订单列表

	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
	// 发货任务ID号（为兼容单个资源开通情况保留）

	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
	// 发货任务ID号列表

	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
	// 资源列表

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 交易流水号（一般为大订单号）

	BillId *string `json:"BillId,omitempty" name:"BillId"`
}

type InstanceTypeConfigResponse struct {

	// 实例规格配置结果集合

	InstanceTypeConfigSet []*InstanceTypeConfigDO `json:"InstanceTypeConfigSet,omitempty" name:"InstanceTypeConfigSet"`
	// 最大主题分区

	MaximumTopicPartition *int64 `json:"MaximumTopicPartition,omitempty" name:"MaximumTopicPartition"`
	// 最大实例消费组

	MaximumInstanceConsumerGroup *int64 `json:"MaximumInstanceConsumerGroup,omitempty" name:"MaximumInstanceConsumerGroup"`
}

type CreateUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type DeleteUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// ip白名单列表

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
}

func (r *DeleteTopicIpWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicIpWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GoodsInfo struct {

	// 订单类型ID，接入计费时由计费后台分配

	GoodsCategoryId *int64 `json:"GoodsCategoryId,omitempty" name:"GoodsCategoryId"`
	// 区域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 可用区ID，没有可用区概念则不传

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 新购时表示商品数量，续费与配置变更时固定传1

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 项目ID，适用于公有云

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目ID，适用于专有云、私有云等非公有云平台

	PlatformProjectId *string `json:"PlatformProjectId,omitempty" name:"PlatformProjectId"`
	// 付费模式，0后付费/1预付费

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 平台，0开放平台/1云平台 没有则不传

	Platform *int64 `json:"Platform,omitempty" name:"Platform"`
	// 商品详情，timeSpan表示购买时间长度，timeUnit表示购买时间单位，pid表示定价公式ID

	GoodsDetail *string `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DescribeBrokerResponse struct {

	// ClusterId

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// ZoneId

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// ClusterName

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// BrokerIp

	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`
	// BrokerId

	BrokerId *int64 `json:"BrokerId,omitempty" name:"BrokerId"`
	// MaxNetworkFlow

	MaxNetworkFlow *uint64 `json:"MaxNetworkFlow,omitempty" name:"MaxNetworkFlow"`
	// MaxDiskSpace

	MaxDiskSpace *uint64 `json:"MaxDiskSpace,omitempty" name:"MaxDiskSpace"`
	// MinPort

	MinPort *uint64 `json:"MinPort,omitempty" name:"MinPort"`
	// MaxPort

	MaxPort *uint64 `json:"MaxPort,omitempty" name:"MaxPort"`
	// AvailableNetworkFlow

	AvailableNetworkFlow *uint64 `json:"AvailableNetworkFlow,omitempty" name:"AvailableNetworkFlow"`
	// AvailableDiskSpace

	AvailableDiskSpace *uint64 `json:"AvailableDiskSpace,omitempty" name:"AvailableDiskSpace"`
}

type Partition struct {

	// 分区ID

	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

type SubscribedInfo struct {

	// 订阅的主题名

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 订阅的分区

	Partition []*int64 `json:"Partition,omitempty" name:"Partition"`
	// 分区offset信息

	PartitionOffset []*PartitionOffset `json:"PartitionOffset,omitempty" name:"PartitionOffset"`
}

type DescribeBrokersRequest struct {
	*tchttp.BaseRequest

	// ClusterId

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// BrokerIp

	BrokerIp *string `json:"BrokerIp,omitempty" name:"BrokerIp"`
	// ZoneId

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// SearchWord

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBrokersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBrokersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例的状态。0：创建中，1：运行中，2：删除中 ， 5 隔离中，-1 创建失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 是否开源实例。开源：true，不开源：false

	IfCommunity *bool `json:"IfCommunity,omitempty" name:"IfCommunity"`
}

type CreateRouteRequest struct {
	*tchttp.BaseRequest

	// 实例唯一id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 路由网络类型。目前只支持 3

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// vpc网络Id。当VipType=3时必传

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// vpc子网id。当VipType=3时必传

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 访问类型

	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
	// 是否需要权限管理

	AuthFlag *int64 `json:"AuthFlag,omitempty" name:"AuthFlag"`
	// 调用方appId

	CallerAppid *int64 `json:"CallerAppid,omitempty" name:"CallerAppid"`
}

func (r *CreateRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIfCommunityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *CommunityResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIfCommunityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIfCommunityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Acl struct {

	// Acl资源类型，（0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID）当前只有TOPIC，

	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源名称，和resourceType相关如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 用户列表，默认为User:*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户

	Principal *string `json:"Principal,omitempty" name:"Principal"`
	// 默认为*，表示任何host都可以访问，当前ckafka不支持host为*，但是后面开源kafka的产品化会直接支持

	Host *string `json:"Host,omitempty" name:"Host"`
	// Acl操作方式(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS，12:IDEMPOTEN_WRITE)

	Operation *int64 `json:"Operation,omitempty" name:"Operation"`
	// 权限类型(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)

	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`
}

type InstanceDetail struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 访问实例的vip 信息

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 访问实例的端口信息

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// 虚拟IP列表

	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`
	// 实例的状态。0：创建中，1：运行中，2：删除中：5隔离中， -1 创建失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例带宽，单位Mbps

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 实例的存储大小，单位GB

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 可用区域ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// vpcId，如果为空，说明是基础网络

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 实例是否续费，int  枚举值：1表示自动续费，2表示明确不自动续费

	RenewFlag *int64 `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 实例状态 int：0表示健康，1表示告警，2 表示实例状态异常

	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`
	// 实例状态信息

	HealthyMessage *string `json:"HealthyMessage,omitempty" name:"HealthyMessage"`
	// 实例创建时间时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 实例过期时间

	ExpireTime *int64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 是否为内部客户。值为1 表示内部客户

	IsInternal *int64 `json:"IsInternal,omitempty" name:"IsInternal"`
	// Topic个数

	TopicNum *int64 `json:"TopicNum,omitempty" name:"TopicNum"`
	// 标识tag

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// kafka版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
	// 跨可用区

	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
	// ckafka售卖类型

	Cvm *int64 `json:"Cvm,omitempty" name:"Cvm"`
	// 集群的 CPU 架构。1:x86; 2:arm; 按位标记架构类型，例：3 二进制为 0011，表示 x86、arm 混合架构

	CpuArch *int64 `json:"CpuArch,omitempty" name:"CpuArch"`
	// 实例所在集群的Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例所在集群的名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 付费模式。-1:非付费实例(默认)；0:后付费实例；1:预付费实例

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

type DescribeConsumerGroupRequest struct {
	*tchttp.BaseRequest

	// ckafka实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 可选，用户需要查询的group名称。

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 可选，用户需要查询的group中的对应的topic名称，如果指定了该参数，而group又未指定则忽略该参数。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 本次返回个数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeConsumerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果集列表

		Result *GroupResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的实例详情结果对象

		Result *InstanceDetailResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例日志的最长保留时间，单位分钟，最大30天，0代表不开启日志保留时间回收策略

	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`
	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例配置

	Config *ModifyInstanceAttributesConfig `json:"Config,omitempty" name:"Config"`
}

func (r *ModifyInstanceAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceTceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回参数

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceTceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceTceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTceInstanceResponse struct {

	// flowId

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
	// instanceId数组

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

type GroupResponse struct {

	// 计数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// GroupList

	GroupList []*DescribeGroup `json:"GroupList,omitempty" name:"GroupList"`
}

type CreatePartitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果集

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePartitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的路由信息结果集

		Result *RouteResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type ModifyGroupOffsetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupOffsetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeConnectorRequest struct {
	*tchttp.BaseRequest

	// connectorId

	ConnectorId *string `json:"ConnectorId,omitempty" name:"ConnectorId"`
}

func (r *ResumeConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConnectorResponse struct {

	// 符合条件的 connector 数量。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 数据同步任务列表。

	ConnectorList []*Connector `json:"ConnectorList,omitempty" name:"ConnectorList"`
}

type RouteResponse struct {

	// 路由信息列表

	Routers []*Route `json:"Routers,omitempty" name:"Routers"`
}

type DescribeAppInfoRequest struct {
	*tchttp.BaseRequest

	// 偏移位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 本次查询用户数目最大数量限制，最大值为50，默认50

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAppInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaPriceRequest struct {
	*tchttp.BaseRequest

	// 实例规格，不同规格类型代表实例不同的配置组合（吞吐量、磁盘）。具体取值可调用接口GetCkafkaTypeConfigs获取最新规格表。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例时长，单位：月，最小值1，最大值为36

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 购买实例数量。取值范围：[1，20]

	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
	// 询价标记：不填或者 0 默认为新够买询价  1为续费

	InquireFlag *int64 `json:"InquireFlag,omitempty" name:"InquireFlag"`
	// 磁盘动态扩容的磁盘扩展包

	Disk *int64 `json:"Disk,omitempty" name:"Disk"`
	// 独占集群节点扩展包

	Node *int64 `json:"Node,omitempty" name:"Node"`
	// 区域ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

func (r *DescribeCkafkaPriceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaPriceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommunityResponse struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否社区版

	IfCommunity *bool `json:"IfCommunity,omitempty" name:"IfCommunity"`
}

type ZoneResponse struct {

	// zone列表

	ZoneList []*ZoneInfo `json:"ZoneList,omitempty" name:"ZoneList"`
	// 最大购买实例个数

	MaxBuyInstanceNum *int64 `json:"MaxBuyInstanceNum,omitempty" name:"MaxBuyInstanceNum"`
	// 最大购买带宽 单位Mb/s

	MaxBandwidth *int64 `json:"MaxBandwidth,omitempty" name:"MaxBandwidth"`
	// 后付费单位价格

	UnitPrice *Price `json:"UnitPrice,omitempty" name:"UnitPrice"`
	// 后付费消息单价

	MessagePrice *Price `json:"MessagePrice,omitempty" name:"MessagePrice"`
	// 用户独占集群信息

	ClusterInfo []*ClusterInfo `json:"ClusterInfo,omitempty" name:"ClusterInfo"`
}

type ModifyGroupOffsetsRequest struct {
	*tchttp.BaseRequest

	// kafka实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// kafka 消费分组

	Group *string `json:"Group,omitempty" name:"Group"`
	// 重置offset的策略，入参含义 0. 对齐shift-by参数，代表把offset向前或向后移动shift条 1. 对齐参考(by-duration,to-datetime,to-earliest,to-latest),代表把offset移动到指定timestamp的位置 2. 对齐参考(to-offset)，代表把offset移动到指定的offset位置

	Strategy *int64 `json:"Strategy,omitempty" name:"Strategy"`
	// 表示需要重置的topics， 不填表示全部

	Topics []*string `json:"Topics,omitempty" name:"Topics"`
	// 当strategy为0时，必须包含该字段，可以大于零代表会把offset向后移动shift条，小于零则将offset向前回溯shift条数。正确重置后新的offset应该是(old_offset + shift)，需要注意的是如果新的offset小于partition的earliest则会设置为earliest，如果大于partition 的latest则会设置为latest

	Shift *int64 `json:"Shift,omitempty" name:"Shift"`
	// 单位ms。当strategy为1时，必须包含该字段，其中-2表示重置offset到最开始的位置，-1表示重置到最新的位置(相当于清空)，其它值则代表指定的时间，会获取topic中指定时间的offset然后进行重置，需要注意的时，如果指定的时间不存在消息，则获取最末尾的offset。

	ShiftTimestamp *int64 `json:"ShiftTimestamp,omitempty" name:"ShiftTimestamp"`
	// 需要重新设置的offset位置。当strategy为2，必须包含该字段。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 需要重新设置的partition的列表，如果没有指定Topics参数。则重置全部topics的对应的Partition列表里的partition。指定Topics时则重置指定的topic列表的对应的Partitions列表的partition。

	Partitions []*uint64 `json:"Partitions,omitempty" name:"Partitions"`
}

func (r *ModifyGroupOffsetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGroupOffsetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyDetail struct {

	// 总折扣

	Total *float64 `json:"Total,omitempty" name:"Total"`
	// 用户个人折扣

	User *float64 `json:"User,omitempty" name:"User"`
	// 官网基础折扣

	Common *int64 `json:"Common,omitempty" name:"Common"`
	// Activity

	Activity *int64 `json:"Activity,omitempty" name:"Activity"`
	// 折扣类型

	DiscountType *string `json:"DiscountType,omitempty" name:"DiscountType"`
	// 折扣ID

	DiscountId *int64 `json:"DiscountId,omitempty" name:"DiscountId"`
	// 优惠类型

	PreferentialType *int64 `json:"PreferentialType,omitempty" name:"PreferentialType"`
	// 指定折扣pid

	DiscountSpecifiedPid *int64 `json:"DiscountSpecifiedPid,omitempty" name:"DiscountSpecifiedPid"`
	// Combine

	Combine *float64 `json:"Combine,omitempty" name:"Combine"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果集

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePreResponse struct {

	// 返回的code，0为正常，非0为错误

	ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`
	// 成功消息

	ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
	// 操作型返回的Data数据

	Data *CreateInstancePreData `json:"Data,omitempty" name:"Data"`
}

type TopicPartitionDO struct {

	// Partition ID

	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
	// Leader 运行状态

	LeaderStatus *int64 `json:"LeaderStatus,omitempty" name:"LeaderStatus"`
	// ISR 个数

	IsrNum *int64 `json:"IsrNum,omitempty" name:"IsrNum"`
	// 副本个数

	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
}

type CreateAclRequest struct {
	*tchttp.BaseRequest

	// 实例id信息

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Acl资源类型，(0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID)，当前只有TOPIC，其它字段用于后续兼容开源kafka的acl时使用

	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// Acl操作方式，(0:UNKNOWN，1:ANY，2:ALL，3:READ，4:WRITE，5:CREATE，6:DELETE，7:ALTER，8:DESCRIBE，9:CLUSTER_ACTION，10:DESCRIBE_CONFIGS，11:ALTER_CONFIGS)

	Operation *int64 `json:"Operation,omitempty" name:"Operation"`
	// 权限类型，(0:UNKNOWN，1:ANY，2:DENY，3:ALLOW)，当前ckakfa支持ALLOW(相当于白名单)，其它用于后续兼容开源kafka的acl时使用

	PermissionType *int64 `json:"PermissionType,omitempty" name:"PermissionType"`
	// 默认为\*，表示任何host都可以访问，当前ckafka不支持host为\*，但是后面开源kafka的产品化会直接支持

	Host *string `json:"Host,omitempty" name:"Host"`
	// 用户列表，默认为*，表示任何user都可以访问，当前用户只能是用户列表中包含的用户

	Principal *string `json:"Principal,omitempty" name:"Principal"`
}

func (r *CreateAclRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAclRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Error struct {

	// 错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 错误信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type CreateTopicResp struct {

	// 主题Id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type Topic struct {

	// 主题的ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 主题的名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
}

type CreatePartitionRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 主题分区个数

	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
}

func (r *CreatePartitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePartitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancePreData struct {

	// CreateInstancePre返回固定为0，不能作为CheckTaskStatus的查询条件。只是为了保证和后台数据结构对齐。

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
	// 订单号列表

	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
}

type DescribeClustersController struct {

	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// epoch

	Epoch *int64 `json:"Epoch,omitempty" name:"Epoch"`
}

type DeleteConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributesConfig struct {

	// 自动创建 true 表示开启，false 表示不开启

	AutoCreateTopicEnable *bool `json:"AutoCreateTopicEnable,omitempty" name:"AutoCreateTopicEnable"`
	// 可选，如果auto.create.topic.enable设置为true没有设置该值时，默认设置为3

	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitempty" name:"DefaultNumPartitions"`
	// 如歌auto.create.topic.enable设置为true没有指定该值时默认设置为2

	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitempty" name:"DefaultReplicationFactor"`
}

type Price struct {

	// 折扣价

	RealTotalCost *float64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 原价

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
}

type DescribeACLResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的ACL结果集对象

		Result *AclResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeACLResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeACLResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppIdResponse struct {

	// 符合要求的所有AppId数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 符合要求的App Id列表

	AppIdList []*int64 `json:"AppIdList,omitempty" name:"AppIdList"`
}

type DescribeGroup struct {

	// groupId

	Group *string `json:"Group,omitempty" name:"Group"`
	// 该 group 使用的协议。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
}

type JgwOperateResponse struct {

	// 返回的code，0为正常，非0为错误

	ReturnCode *string `json:"ReturnCode,omitempty" name:"ReturnCode"`
	// 成功消息

	ReturnMessage *string `json:"ReturnMessage,omitempty" name:"ReturnMessage"`
	// 操作型返回的Data数据,可能有flowId等

	Data *OperateResponseData `json:"Data,omitempty" name:"Data"`
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建结果

		Result *CreateTopicResp `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type CreateTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除主题IP白名单结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicIpWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicIpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceTceRequest struct {
	*tchttp.BaseRequest

	// 实例名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 可用区

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// vpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 可选。实例日志的最长保留时间，单位分钟，默认为10080（7天），最大30天，不填默认0，代表不开启日志保留时间回收策略

	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`
	// 实例带宽

	BandWidth *int64 `json:"BandWidth,omitempty" name:"BandWidth"`
	// 实例磁盘

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// ClusterId

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreateInstanceTceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceTceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Config struct {

	// 消息保留时间

	Retention *int64 `json:"Retention,omitempty" name:"Retention"`
	// 最小同步复制数

	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`
	// 日志清理模式，默认 delete。
	// delete：日志按保存时间删除；compact：日志按 key 压缩；compact, delete：日志按 key 压缩且会保存时间删除。

	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`
	// Segment 分片滚动的时长

	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`
	// 0表示 false。 1表示 true。

	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`
	// Segment 分片滚动的字节数

	SegmentBytes *int64 `json:"SegmentBytes,omitempty" name:"SegmentBytes"`
	// 最大消息字节数

	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`
}

type Group struct {

	// 组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type FetchMessageByOffsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *ConsumerRecord `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FetchMessageByOffsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FetchMessageByOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUserRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户密码

	Password *string `json:"Password,omitempty" name:"Password"`
}

func (r *CreateUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果

		Result *InstanceResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果

		Result *TopicResult `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIfCommunityRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeIfCommunityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIfCommunityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的实例详情结果对象

		Result *InstanceDetailResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerGroup struct {

	// 用户组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 订阅信息实体

	SubscribedInfo []*SubscribedInfo `json:"SubscribedInfo,omitempty" name:"SubscribedInfo"`
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名称，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// Partition个数，大于0

	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
	// 副本个数，不能多于 broker 数，最大为3

	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
	// ip白名单开关, 1:打开  0:关闭，默认不打开

	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`
	// Ip白名单列表，配额限制，enableWhileList=1时必选

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// 清理日志策略，日志清理模式，默认为"delete"。"delete"：日志按保存时间删除，"compact"：日志按 key 压缩，"compact, delete"：日志按 key 压缩且会按保存时间删除。

	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`
	// 主题备注，是一个不超过 64 个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线(-)

	Note *string `json:"Note,omitempty" name:"Note"`
	// 默认为1

	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`
	// 是否允许未同步的副本选为leader，false:不允许，true:允许，默认不允许

	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`
	// 可消息选。保留时间，单位ms，当前最小值为60000ms

	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`
	// Segment分片滚动的时长，单位ms，当前最小为3600000ms

	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`
	// 主题消息最大值，单位为 Byte，最大值为12582912Byte（即12MB）。

	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`
}

func (r *CreateTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例ID过滤

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// （过滤条件）按照实例名称过滤，支持模糊查询

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部

	Status []*int64 `json:"Status,omitempty" name:"Status"`
	// 偏移量，不填默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认10，最大值20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 匹配标签key值。

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrderConfigRequest struct {
	*tchttp.BaseRequest

	// 实例规格，不同规格类型代表实例不同的配置组合（吞吐量、磁盘）。具体取值可调用接口GetCkafkaTypeConfigs获取最新规格表。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 实例时长，单位：月，最小值1，最大值为36

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 购买实例数量。取值范围：[1，20]

	InstanceNum *int64 `json:"InstanceNum,omitempty" name:"InstanceNum"`
	// 询价标记：不填或者 0 默认为新够买询价 1为续费

	InquireFlag *int64 `json:"InquireFlag,omitempty" name:"InquireFlag"`
	// 区域ID

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 购买带宽规格，单位：MB

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 购买磁盘规格，默认为带宽对应的最小值，单位：GB

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 购买磁盘类型，默认为CLOUD_TCS

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 消息保留时间，默认为72，单位：小时

	MessageRetention *int64 `json:"MessageRetention,omitempty" name:"MessageRetention"`
	// 购买Topic规格，默认为带宽对应的最小值，单位：个

	Topic *int64 `json:"Topic,omitempty" name:"Topic"`
	// 购买Partition规格，默认为带宽对应的最小值，单位：个

	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
	// 购买实例版本

	BrokerVersion *string `json:"BrokerVersion,omitempty" name:"BrokerVersion"`
	// 标记：purchase新购，renew续费，modify升配

	CategoryAction *string `json:"CategoryAction,omitempty" name:"CategoryAction"`
	// 标准版S2计费类型

	BillType *string `json:"BillType,omitempty" name:"BillType"`
	// 公网带宽计费类型（预付费按带宽结算(包年带宽):BANDWIDTH_PREPAID、带宽按小时后付费(按小时带宽):BANDWIDTH_POSTPAID_BY_HOUR

	PublicNetworkChargeType *string `json:"PublicNetworkChargeType,omitempty" name:"PublicNetworkChargeType"`
	// 公网带宽

	PublicNetworkMonthly *int64 `json:"PublicNetworkMonthly,omitempty" name:"PublicNetworkMonthly"`
	// 实例ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

func (r *DescribeOrderConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrderConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicRequest struct {
	*tchttp.BaseRequest

	// 实例 ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 过滤条件，按照 topicName 过滤，支持模糊查询

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// 偏移量，不填默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认为10，最大值为50

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果

		Result []*GroupInfoResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例ID过滤

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// （过滤条件）按照实例名称过滤，支持模糊查询

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部

	Status []*int64 `json:"Status,omitempty" name:"Status"`
	// 偏移量，不填默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认10，最大值20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 匹配标签key值。

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 过滤器，Name 值有 Ip，ZoneId

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstanceDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupOffsetTopic struct {

	// 主题名称

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 该主题分区数组，其中每个元素为一个 json object

	Partitions []*GroupOffsetPartition `json:"Partitions,omitempty" name:"Partitions"`
}

type User struct {

	// 用户id

	UserId *int64 `json:"UserId,omitempty" name:"UserId"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAppInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的符合要求的App Id列表

		Result *AppIdResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeOrderConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOrderConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOrderConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Connector struct {

	// connectorId。

	ConnectorId *string `json:"ConnectorId,omitempty" name:"ConnectorId"`
	// connector 名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// connector 类型。 source：导入数据到 Kafka；sink：将数据从 Kafka 导出。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 执行该任务的 class 名称。

	ConnectorClass *string `json:"ConnectorClass,omitempty" name:"ConnectorClass"`
	// 源所在地域。

	SourceRegion *string `json:"SourceRegion,omitempty" name:"SourceRegion"`
	// 源地址。

	Source *string `json:"Source,omitempty" name:"Source"`
	// 目的所在地域。

	SinkRegion *string `json:"SinkRegion,omitempty" name:"SinkRegion"`
	// 目的地址。

	Sink *string `json:"Sink,omitempty" name:"Sink"`
	// connector 当前状态。 该状态展示有一定的延迟，任务状态可通过 GetConnectorStatus 接口 获取。
	// UNASSIGNED：任务还未分配；RUNNING：connector 正在运行；PAUSED：connector 已经暂停；FAILED：任务失败；DESTROYED：任务销毁。

	Status *string `json:"Status,omitempty" name:"Status"`
	// connector 描述信息。

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 修改时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TopicDetailResponse struct {

	// 返回的主题详情列表

	TopicList []*TopicDetail `json:"TopicList,omitempty" name:"TopicList"`
	// 符合条件的所有主题详情数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type InstanceResponse struct {

	// 符合条件的实例列表

	InstanceList []*Instance `json:"InstanceList,omitempty" name:"InstanceList"`
	// 符合条件的结果总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type Route struct {

	// 实例接入方式
	// 0：PLAINTEXT (明文方式，没有带用户信息老版本及社区版本都支持)
	// 1：SASL_PLAINTEXT（明文方式，不过在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）
	// 2：SSL（SSL加密通信，没有带用户信息，老版本及社区版本都支持）
	// 3：SASL_SSL（SSL加密通信，在数据开始时，会通过SASL方式登录鉴权，仅社区版本支持）

	AccessType *int64 `json:"AccessType,omitempty" name:"AccessType"`
	// 路由ID

	RouteId *int64 `json:"RouteId,omitempty" name:"RouteId"`
	// vip网络类型（1:外网TGW  2:基础网络 3:VPC网络 4:腾讯云支持环境(一般用于内部实例) 5:SSL外网访问方式访问 6:黑石环境vpc）

	VipType *int64 `json:"VipType,omitempty" name:"VipType"`
	// 虚拟IP列表

	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`
	// 域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 域名port

	DomainPort *int64 `json:"DomainPort,omitempty" name:"DomainPort"`

	// 内部路由的标志(0:非内部 1:内部路由)

	InternalFlag *int64 `json:"InternalFlag,omitempty" name:"InternalFlag"`

	// 内部路由的用途

	UsedFor *string `json:"UsedFor,omitempty" name:"UsedFor"`

	// 路由创建进度（0:已完成 1:处理中）

	Processing *int64 `json:"Processing,omitempty" name:"Processing"`

	// 删除时间戳

	DeleteTimestamp *string `json:"DeleteTimestamp,omitempty" name:"DeleteTimestamp"`

	// broker vip映射列表

	BrokerVipList []*VipEntity `json:"BrokerVipList,omitempty" name:"BrokerVipList"`

	// vpcid

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`

	// 子网信息

	Subnet *string `json:"Subnet,omitempty" name:"Subnet"`

	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
}

type VipEntity struct {

	// 虚拟IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 虚拟端口

	Vport *string `json:"Vport,omitempty" name:"Vport"`
}

type CancelAuthorizationTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelAuthorizationTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelAuthorizationTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaZoneResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果复杂对象实体

		Result *ZoneResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCkafkaZoneResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaZoneResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserResponse struct {

	// 符合条件的用户列表

	Users []*User `json:"Users,omitempty" name:"Users"`
	// 符合条件的总用户数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type GroupOffsetResponse struct {

	// 符合调节的总结果数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 该主题分区数组，其中每个元素为一个 json object

	TopicList []*GroupOffsetTopic `json:"TopicList,omitempty" name:"TopicList"`
}

type CreateInstanceTceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回参数

		Result *CreateTceInstanceResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceTceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceTceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaPriceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result []*PriceResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCkafkaPriceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaPriceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *TaskStatusResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConnectorConfigsResponse struct {

	// 执行该任务的 class 名称。

	ConnectorClass *string `json:"ConnectorClass,omitempty" name:"ConnectorClass"`
	// 目的实例名称。

	DestInstance *string `json:"DestInstance,omitempty" name:"DestInstance"`
	// 主题配置。

	TopicConfig *string `json:"TopicConfig,omitempty" name:"TopicConfig"`
	// connector 版本。

	ConnectorVersion *string `json:"ConnectorVersion,omitempty" name:"ConnectorVersion"`
	// 源实例。

	SrcInstance *string `json:"SrcInstance,omitempty" name:"SrcInstance"`
	// 任务同步启动时候的 offset 设置

	OffsetReset *string `json:"OffsetReset,omitempty" name:"OffsetReset"`
	// 源实例访问 broker 地址和端口。

	SourceBroker *string `json:"SourceBroker,omitempty" name:"SourceBroker"`
	// 源和目的 topic 的 partition 数量是否要求保持一致。
	// true：保持一致；false：不一致。

	KeepPartition *bool `json:"KeepPartition,omitempty" name:"KeepPartition"`
	// connector 名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否开启主题自动创建。true 表示开启，false 表示不开启。

	AutoCreate *bool `json:"AutoCreate,omitempty" name:"AutoCreate"`
}

type DescribeGroupOffsetsRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例 ID 过滤

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Kafka 消费分组

	Group *string `json:"Group,omitempty" name:"Group"`
	// group 订阅的主题名称数组，如果没有该数组，则表示指定的 group 下所有 topic 信息

	Topics []*string `json:"Topics,omitempty" name:"Topics"`
	// 模糊匹配 topicName

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// 本次查询的偏移位置，默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 本次返回结果的最大个数，默认为50，最大值为50

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupOffsetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupOffsetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Assignment struct {

	// assingment版本信息

	Version *int64 `json:"Version,omitempty" name:"Version"`
	// topic信息列表

	Topics []*GroupInfoTopics `json:"Topics,omitempty" name:"Topics"`
}

type DescribeConnectorStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *ConnectorStatus `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConnectorStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 搜索关键字

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最大返回数量

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AuthorizeTokenRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// token

	Tokens *string `json:"Tokens,omitempty" name:"Tokens"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
}

func (r *AuthorizeTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthorizeTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppIdIsVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *AppIdIsVipResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppIdIsVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppIdIsVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 用户当前密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 用户新密码

	PasswordNew *string `json:"PasswordNew,omitempty" name:"PasswordNew"`
}

func (r *ModifyPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SoldOut struct {

	// false未售罄

	Cvm *string `json:"Cvm,omitempty" name:"Cvm"`
}

type ZoneInfo struct {

	// zone的id

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 是否内部APP

	IsInternalApp *int64 `json:"IsInternalApp,omitempty" name:"IsInternalApp"`
	// app id

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 标识

	Flag *bool `json:"Flag,omitempty" name:"Flag"`
	// zone名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// zone状态

	ZoneStatus *int64 `json:"ZoneStatus,omitempty" name:"ZoneStatus"`
	// 额外标识

	Exflag *string `json:"Exflag,omitempty" name:"Exflag"`
	// json对象，key为机型，value true为售罄，false为未售罄

	SoldOut *string `json:"SoldOut,omitempty" name:"SoldOut"`
}

type CreateConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppIdIsVipRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAppIdIsVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppIdIsVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupInfoRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例 ID 过滤。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Kafka 消费分组，Consumer-group，这里是数组形式。如果使用拼接数组元素方式传值（一般 GET 请求用），可以使用格式：GroupList.0=xxx&GroupList.1=yyy。

	GroupList []*string `json:"GroupList,omitempty" name:"GroupList"`
}

func (r *DescribeGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteRequest struct {
	*tchttp.BaseRequest

	// 实例唯一id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FetchMessageListByTimestampResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果。注意，列表中不返回具体的消息内容（key、value），如果需要查询具体消息内容，请使用FetchMessageByOffset接口

		Result []*ConsumerRecord `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FetchMessageListByTimestampResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FetchMessageListByTimestampResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAclResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAclResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAclResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceTceRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例带宽

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 实例磁盘

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

func (r *ModifyResourceTceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceTceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OperateResponseData struct {

	// FlowId

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

type DescribeConnectorConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *ConnectorConfigsResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConnectorConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FetchMessageListByOffsetRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 分区id

	Partition *uint64 `json:"Partition,omitempty" name:"Partition"`
	// 位点信息

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 最大查询条数，默认20，最大20

	SinglePartitionRecordNumber *uint64 `json:"SinglePartitionRecordNumber,omitempty" name:"SinglePartitionRecordNumber"`
}

func (r *FetchMessageListByOffsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FetchMessageListByOffsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Region struct {

	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 区域名称

	AreaName *string `json:"AreaName,omitempty" name:"AreaName"`
	// 地域代码

	RegionCode *string `json:"RegionCode,omitempty" name:"RegionCode"`
	// 地域代码（V3版本）

	RegionCodeV3 *string `json:"RegionCodeV3,omitempty" name:"RegionCodeV3"`
	// NONE:默认值不支持任何特殊机型\nCVM:支持CVM类型

	Support *string `json:"Support,omitempty" name:"Support"`
}

type DescribeCkafkaTypeConfigsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCkafkaTypeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaTypeConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AclResponse struct {

	// 符合条件的总数据条数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// ACL列表

	AclList []*Acl `json:"AclList,omitempty" name:"AclList"`
}

type PartDetail struct {

	// 硬盘价格

	Disk *PriceObject `json:"Disk,omitempty" name:"Disk"`
	// ckafka价格

	CloudKafka *PriceObject `json:"CloudKafka,omitempty" name:"CloudKafka"`
}

type Tag struct {

	// 标签的key

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 标签的值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type CreateRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type DescribeConnectorConfigsRequest struct {
	*tchttp.BaseRequest

	// ConnectorId

	ConnectorId *string `json:"ConnectorId,omitempty" name:"ConnectorId"`
}

func (r *DescribeConnectorConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果对象

		Result *TopicAttributesResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FetchMessageListByOffsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果。注意，列表中不返回具体的消息内容（key、value），如果需要查询具体消息内容，请使用FetchMessageByOffset接口

		Result []*ConsumerRecord `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *FetchMessageListByOffsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FetchMessageListByOffsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FetchMessageListByTimestampRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 分区id

	Partition *uint64 `json:"Partition,omitempty" name:"Partition"`
	// 查询开始时间，13位时间戳

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 最大查询条数，默认20，最大20

	SinglePartitionRecordNumber *uint64 `json:"SinglePartitionRecordNumber,omitempty" name:"SinglePartitionRecordNumber"`
}

func (r *FetchMessageListByTimestampRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *FetchMessageListByTimestampRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRouteRequest struct {
	*tchttp.BaseRequest

	// 实例唯一id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 路由id

	RouteId *int64 `json:"RouteId,omitempty" name:"RouteId"`
	// 调用方appId

	CallerAppid *int64 `json:"CallerAppid,omitempty" name:"CallerAppid"`
}

func (r *DeleteRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConnectorRequest struct {
	*tchttp.BaseRequest

	// zoneId

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 可选。指定在指定的kafka connect集群上进行添加connect任务。指定了该参数可以在独占集群上创建其它用户的实例，一般用于测试独占集群时使用

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 任务描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 必选。可以理解为执行该任务的class名称，不同种类的同步connector有不同的名称，拿我们当前做ckafka不同实例之间的同步就是com.tencent.ckafka.replicator.KafkaSourceConnector，这个主要是相关的kafka connector开发订阅，要保证不同connector不冲突即可。当前现在仅支持com.tencent.ckafka.replicator.KafkaSourceConnector，后面会不断的扩充

	ConnectorClass *string `json:"ConnectorClass,omitempty" name:"ConnectorClass"`
	// 配置

	Config *string `json:"Config,omitempty" name:"Config"`
	// connect 名称配置

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据同步任务返回结果

		Result *ConnectorResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PriceResponse struct {

	// 价格

	Price *int64 `json:"Price,omitempty" name:"Price"`
	// 组成详情

	PartDetail *PartDetail `json:"PartDetail,omitempty" name:"PartDetail"`
	// 是否用户价格

	HasUserPrice *bool `json:"HasUserPrice,omitempty" name:"HasUserPrice"`
	// 总花费

	TotalCost *int64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 货币

	Currency *string `json:"Currency,omitempty" name:"Currency"`
	// 商品的时间单位

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 商品的时间大小

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 资源实例个数

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 类型

	FromType *int64 `json:"FromType,omitempty" name:"FromType"`
	// pid

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 真实总花费

	RealTotalCost *int64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 产品代码

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 子产品代码

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 优惠

	Policy *float64 `json:"Policy,omitempty" name:"Policy"`
	// 优惠详情

	PolicyDetail *PolicyDetail `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
	// 单价

	UnitPrice *int64 `json:"UnitPrice,omitempty" name:"UnitPrice"`
}

type ClusterInfo struct {

	// 集群Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群可用区

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// Broker 的可用区列表

	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
	// 集群的 CPU 架构。1:x86; 2:arm; 按位标记架构类型，例：3 二进制为 0011，表示 x86、arm 混合架构

	CpuArch *int64 `json:"CpuArch,omitempty" name:"CpuArch"`
	// 集群最大磁盘 单位GB

	MaxDiskSize *int64 `json:"MaxDiskSize,omitempty" name:"MaxDiskSize"`
	// 集群最大带宽 单位MB/s

	MaxBandWidth *int64 `json:"MaxBandWidth,omitempty" name:"MaxBandWidth"`
	// 集群当前可用磁盘  单位GB

	AvailableDiskSize *int64 `json:"AvailableDiskSize,omitempty" name:"AvailableDiskSize"`
	// 集群当前可用带宽 单位MB/s

	AvailableBandWidth *int64 `json:"AvailableBandWidth,omitempty" name:"AvailableBandWidth"`
}

type InstanceAttributesResponse struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 接入点 VIP 列表信息

	VipList []*VipEntity `json:"VipList,omitempty" name:"VipList"`
	// 虚拟IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 虚拟端口

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// 实例的状态。0：创建中，1：运行中，2：删除中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 实例带宽，单位：Mbps

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 实例的存储大小，单位：GB

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 可用区

	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// VPC 的 ID，为空表示是基础网络

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网 ID， 为空表示基础网络

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 实例健康状态， 1：健康，2：告警，3：异常

	Healthy *int64 `json:"Healthy,omitempty" name:"Healthy"`
	// 实例健康信息，当前会展示磁盘利用率，最大长度为256

	HealthyMessage *string `json:"HealthyMessage,omitempty" name:"HealthyMessage"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 消息保存时间,单位为分钟

	MsgRetentionTime *int64 `json:"MsgRetentionTime,omitempty" name:"MsgRetentionTime"`
	// 自动创建 Topic 配置， 若该字段为空，则表示未开启自动创建

	Config *InstanceConfigDO `json:"Config,omitempty" name:"Config"`
	// 剩余创建分区数

	RemainderPartitions *int64 `json:"RemainderPartitions,omitempty" name:"RemainderPartitions"`
	// 剩余创建主题数

	RemainderTopics *int64 `json:"RemainderTopics,omitempty" name:"RemainderTopics"`
	// 当前创建分区数

	CreatedPartitions *int64 `json:"CreatedPartitions,omitempty" name:"CreatedPartitions"`
	// 当前创建主题数

	CreatedTopics *int64 `json:"CreatedTopics,omitempty" name:"CreatedTopics"`
	// 标签数组

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 过期时间

	ExpireTime *uint64 `json:"ExpireTime,omitempty" name:"ExpireTime"`
	// 跨可用区

	ZoneIds []*int64 `json:"ZoneIds,omitempty" name:"ZoneIds"`
	// kafka版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
	// 最大分组数

	MaxGroupNum *int64 `json:"MaxGroupNum,omitempty" name:"MaxGroupNum"`
	// 售卖类型

	Cvm *int64 `json:"Cvm,omitempty" name:"Cvm"`
	// 集群的 CPU 架构。1:x86; 2:arm; 按位标记架构类型，例：3 二进制为 0011，表示 x86、arm 混合架构

	CpuArch *int64 `json:"CpuArch,omitempty" name:"CpuArch"`
	// 实例所在集群的Id

	ClusterId *int64 `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例所在集群的名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 付费模式。-1:非付费实例(默认)；0:后付费实例；1:预付费实例

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
}

type CreateTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// token结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConnectorStatusRequest struct {
	*tchttp.BaseRequest

	// Connector的Id

	ConnectorId *string `json:"ConnectorId,omitempty" name:"ConnectorId"`
}

func (r *DescribeConnectorStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID数组

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DeleteInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfoResponse struct {

	// 错误码，正常为0

	ErrorCode *string `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// group 状态描述（常见的为 Empty、Stable、Dead 三种状态）：
	// Dead：消费分组不存在
	// Empty：消费分组，当前没有任何消费者订阅
	// PreparingRebalance：消费分组处于 rebalance 状态
	// CompletingRebalance：消费分组处于 rebalance 状态
	// Stable：消费分组中各个消费者已经加入，处于稳定状态

	State *string `json:"State,omitempty" name:"State"`
	// 消费分组选择的协议类型正常的消费者一般为 consumer 但有些系统采用了自己的协议如 kafka-connect 用的就是 connect。只有标准的 consumer 协议，本接口才知道具体的分配方式的格式，才能解析到具体的 partition 的分配情况

	ProtocolType *string `json:"ProtocolType,omitempty" name:"ProtocolType"`
	// 消费者 partition 分配算法常见的有如下几种(Kafka 消费者 SDK 默认的选择项为 range)：range、 roundrobin、 sticky

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 仅当 state 为 Stable 且 protocol_type 为 consumer 时， 该数组才包含信息

	Members []*GroupInfoMember `json:"Members,omitempty" name:"Members"`
	// Kafka 消费分组

	Group *string `json:"Group,omitempty" name:"Group"`
}

type InstanceDetailResponse struct {

	// 符合条件的实例总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 符合条件的实例详情列表

	InstanceList []*InstanceDetail `json:"InstanceList,omitempty" name:"InstanceList"`
}

type DescribeBrokersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// DescribeBrokerResponses 返回结果

		Result *DescribeBrokersResponses `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBrokersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBrokersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的主题详情实体

		Result *TopicDetailResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InstanceDeleteResponse struct {

	// 删除实例返回的任务Id

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

type DeleteTopicIpWhiteListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除主题IP白名单结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicIpWhiteListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicIpWhiteListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeACLRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Acl资源类型，(0:UNKNOWN，1:ANY，2:TOPIC，3:GROUP，4:CLUSTER，5:TRANSACTIONAL_ID)，当前只有TOPIC，其它字段用于后续兼容开源kafka的acl时使用

	ResourceType *int64 `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源名称，和resourceType相关，如当resourceType为TOPIC时，则该字段表示topic名称，当resourceType为GROUP时，该字段表示group名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
	// 偏移位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 个数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 关键字匹配

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
}

func (r *DescribeACLRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeACLRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerGroupTopic struct {

	// 主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type TopicDetail struct {

	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 分区数

	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
	// 副本数

	ReplicaNum *int64 `json:"ReplicaNum,omitempty" name:"ReplicaNum"`
	// 备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 是否开启ip鉴权白名单，true表示开启，false表示不开启

	EnableWhiteList *bool `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`
	// ip白名单中ip个数

	IpWhiteListCount *int64 `json:"IpWhiteListCount,omitempty" name:"IpWhiteListCount"`
	// 数据备份cos bucket: 转存到cos 的bucket地址

	ForwardCosBucket *string `json:"ForwardCosBucket,omitempty" name:"ForwardCosBucket"`
	// 数据备份cos 状态： 1 不开启数据备份，0 开启数据备份

	ForwardStatus *int64 `json:"ForwardStatus,omitempty" name:"ForwardStatus"`
	// 数据备份到cos的周期频率

	ForwardInterval *int64 `json:"ForwardInterval,omitempty" name:"ForwardInterval"`
	// 高级配置

	Config *Config `json:"Config,omitempty" name:"Config"`
}

type ModifyTopicAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 主题备注，是一个不超过64个字符的字符串，必须以字母为首字符，剩余部分可以包含字母、数字和横划线-。

	Note *string `json:"Note,omitempty" name:"Note"`
	// IP 白名单开关，1：打开；0：关闭。

	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`
	// Ip白名单列表，配额限制，enableWhileList=1时必选

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// 默认为1。

	MinInsyncReplicas *int64 `json:"MinInsyncReplicas,omitempty" name:"MinInsyncReplicas"`
	// 默认为 0，0：false；1：true。

	UncleanLeaderElectionEnable *int64 `json:"UncleanLeaderElectionEnable,omitempty" name:"UncleanLeaderElectionEnable"`
	// 消息保留时间，单位：ms，当前最小值为60000ms。

	RetentionMs *int64 `json:"RetentionMs,omitempty" name:"RetentionMs"`
	// Segment 分片滚动的时长，单位：ms，当前最小为86400000ms。

	SegmentMs *int64 `json:"SegmentMs,omitempty" name:"SegmentMs"`
	// 主题消息最大值，单位为 Byte，最大值为12582912Byte（即12MB）。

	MaxMessageBytes *int64 `json:"MaxMessageBytes,omitempty" name:"MaxMessageBytes"`
	// 消息删除策略，可以选择delete 或者compact

	CleanUpPolicy *string `json:"CleanUpPolicy,omitempty" name:"CleanUpPolicy"`
}

func (r *ModifyTopicAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseConnectorRequest struct {
	*tchttp.BaseRequest

	// ConnectorId

	ConnectorId *string `json:"ConnectorId,omitempty" name:"ConnectorId"`
}

func (r *PauseConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConnectorRequest struct {
	*tchttp.BaseRequest

	// ConnectorId

	ConnectorId *string `json:"ConnectorId,omitempty" name:"ConnectorId"`
}

func (r *DeleteConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 任务唯一标记

	FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
}

func (r *DescribeTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicAttributesRequest struct {
	*tchttp.BaseRequest

	// 实例 ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

func (r *DescribeTopicAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicDetailRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// （过滤条件）按照topicName过滤，支持模糊查询

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// 偏移量，不填默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认 10，最大值20，取值要大于0

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HourCreateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建按量计费实例的响应内容

		Result *TradeCreateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HourCreateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HourCreateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppIdIsVipResponse struct {

	// 0表示小客户 1 表示大客户

	VipFlag *int64 `json:"VipFlag,omitempty" name:"VipFlag"`
	// 0 : 表示公有云客户， 1 表示内部客户

	InternalApp *int64 `json:"InternalApp,omitempty" name:"InternalApp"`
}

type GroupOffsetPartition struct {

	// topic 的 partitionId

	Partition *int64 `json:"Partition,omitempty" name:"Partition"`
	// consumer 提交的 offset 位置

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 支持消费者提交消息时，传入 metadata 作为它用，当前一般为空字符串

	Metadata *string `json:"Metadata,omitempty" name:"Metadata"`
	// 错误码

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 当前 partition 最新的 offset

	LogEndOffset *int64 `json:"LogEndOffset,omitempty" name:"LogEndOffset"`
	// 未消费的消息个数

	Lag *int64 `json:"Lag,omitempty" name:"Lag"`
}

type DescribeConnectorRequest struct {
	*tchttp.BaseRequest

	// connectorId，根据该参数指定 connector 过滤。

	ConnectorId *string `json:"ConnectorId,omitempty" name:"ConnectorId"`
	// （过滤条件）按照 connector 名称过滤，支持模糊查询。

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// 偏移量，不填默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认50，最大值50 。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConnectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConnectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果集

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerGroupResponse struct {

	// 符合条件的消费组数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 主题列表

	TopicList []*ConsumerGroupTopic `json:"TopicList,omitempty" name:"TopicList"`
	// 消费分组List

	GroupList []*ConsumerGroup `json:"GroupList,omitempty" name:"GroupList"`
	// 所有分区数量

	TotalPartition *int64 `json:"TotalPartition,omitempty" name:"TotalPartition"`
	// 监控的分区列表

	PartitionListForMonitor []*Partition `json:"PartitionListForMonitor,omitempty" name:"PartitionListForMonitor"`
	// 主题总数

	TotalTopic *int64 `json:"TotalTopic,omitempty" name:"TotalTopic"`
	// 监控的主题列表

	TopicListForMonitor []*ConsumerGroupTopic `json:"TopicListForMonitor,omitempty" name:"TopicListForMonitor"`
	// 监控的组列表

	GroupListForMonitor []*Group `json:"GroupListForMonitor,omitempty" name:"GroupListForMonitor"`
}

type DeleteUserRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 用户名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HourModifyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 修改按量计费实例的响应内容

		Result *TradeCreateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HourModifyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HourModifyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConnectorStatus struct {

	// connector 状态。
	// UNASSIGNED：任务还未分配；RUNNING：connector 正在运行；PAUSED：connector 已经暂停；FAILED：任务失败；DESTROYED：任务销毁。

	State *string `json:"State,omitempty" name:"State"`
	// connector 类型，有 source 和 sink 两种类型。

	Type *string `json:"Type,omitempty" name:"Type"`
}

type DeleteInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果集

		Result *InstanceDeleteResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例属性返回结果对象

		Result *InstanceAttributesResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeConnectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果

		Result *JgwOperateResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeConnectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeConnectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTokenRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// user

	User *string `json:"User,omitempty" name:"User"`
}

func (r *CreateTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicIpWhiteListRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// ip白名单列表

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
}

func (r *CreateTopicIpWhiteListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicIpWhiteListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回结果列表

		Result *UserResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupOffsetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果对象

		Result *GroupOffsetResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupOffsetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PartitionOffset struct {

	// Partition,例如"0"或"1"

	Partition *string `json:"Partition,omitempty" name:"Partition"`
	// Offset,例如100

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DescribeInstancesDetailRequest struct {
	*tchttp.BaseRequest

	// （过滤条件）按照实例ID过滤

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// （过滤条件）按照实例名称过滤，支持模糊查询

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// （过滤条件）实例的状态。0：创建中，1：运行中，2：删除中，不填默认返回全部

	Status []*int64 `json:"Status,omitempty" name:"Status"`
	// 偏移量，不填默认为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，不填则默认10，最大值20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 匹配标签key值。

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeInstancesDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerRecord struct {

	// 主题名

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 分区id

	Partition *uint64 `json:"Partition,omitempty" name:"Partition"`
	// 位点

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 消息key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 消息value

	Value *string `json:"Value,omitempty" name:"Value"`
	// 消息时间戳

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 消息headers

	Headers *string `json:"Headers,omitempty" name:"Headers"`
}

type PriceDetail struct {

	// 单价

	SinglePrice *int64 `json:"SinglePrice,omitempty" name:"SinglePrice"`
	// 使用数量

	UsedAmount *int64 `json:"UsedAmount,omitempty" name:"UsedAmount"`
	// 花费

	Cost *int64 `json:"Cost,omitempty" name:"Cost"`
}

type TopicAttributesResponse struct {

	// 主题 ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 主题备注

	Note *string `json:"Note,omitempty" name:"Note"`
	// 分区个数

	PartitionNum *int64 `json:"PartitionNum,omitempty" name:"PartitionNum"`
	// IP 白名单开关，1：打开； 0：关闭

	EnableWhiteList *int64 `json:"EnableWhiteList,omitempty" name:"EnableWhiteList"`
	// IP 白名单列表

	IpWhiteList []*string `json:"IpWhiteList,omitempty" name:"IpWhiteList"`
	// topic 配置数组

	Config *Config `json:"Config,omitempty" name:"Config"`
	// 分区详情

	Partitions []*TopicPartitionDO `json:"Partitions,omitempty" name:"Partitions"`
}

type TopicResult struct {

	// 返回的主题信息列表

	TopicList []*Topic `json:"TopicList,omitempty" name:"TopicList"`
	// 符合条件的 topic 数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type HourModifyRequest struct {
	*tchttp.BaseRequest

	// 后付费商品信息

	GoodsInfo *GoodsInfo `json:"GoodsInfo,omitempty" name:"GoodsInfo"`
	// 需要修改的实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *HourModifyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HourModifyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupInfoMember struct {

	// coordinator 为消费分组中的消费者生成的唯一 ID

	MemberId *string `json:"MemberId,omitempty" name:"MemberId"`
	// 客户消费者 SDK 自己设置的 client.id 信息

	ClientId *string `json:"ClientId,omitempty" name:"ClientId"`
	// 一般存储客户的 IP 地址

	ClientHost *string `json:"ClientHost,omitempty" name:"ClientHost"`
	// 存储着分配给该消费者的 partition 信息

	Assignment *Assignment `json:"Assignment,omitempty" name:"Assignment"`
}

type InstanceTypeConfigDO struct {

	// 型号

	Type *string `json:"Type,omitempty" name:"Type"`
	// 类型描述

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 带宽流量大小，单位Mbqs

	Bandwidth *int64 `json:"Bandwidth,omitempty" name:"Bandwidth"`
	// 磁盘大小，单位GB

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 类型对应的pid信息

	Pid *int64 `json:"Pid,omitempty" name:"Pid"`
	// 该规格可以创建的分区数量配额

	MaximumInstancePartition *int64 `json:"MaximumInstancePartition,omitempty" name:"MaximumInstancePartition"`
	// 该规格可以创建的主题数量配额

	MaximumInstanceTopic *int64 `json:"MaximumInstanceTopic,omitempty" name:"MaximumInstanceTopic"`
}

type CancelAuthorizationTokenRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// token

	Tokens *string `json:"Tokens,omitempty" name:"Tokens"`
	// 用户

	User *string `json:"User,omitempty" name:"User"`
}

func (r *CancelAuthorizationTokenRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelAuthorizationTokenRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 按照名称过滤

	SearchWord *string `json:"SearchWord,omitempty" name:"SearchWord"`
	// 偏移

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 本次返回个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBrokersResponses struct {

	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// Brokers

	Brokers []*DescribeBrokerResponse `json:"Brokers,omitempty" name:"Brokers"`
}

type InstanceConfigDO struct {

	// 是否自动创建主题

	AutoCreateTopicsEnable *bool `json:"AutoCreateTopicsEnable,omitempty" name:"AutoCreateTopicsEnable"`
	// 分区数

	DefaultNumPartitions *int64 `json:"DefaultNumPartitions,omitempty" name:"DefaultNumPartitions"`
	// 默认的复制Factor

	DefaultReplicationFactor *int64 `json:"DefaultReplicationFactor,omitempty" name:"DefaultReplicationFactor"`
}

type PriceObject struct {

	// pid

	Pid *string `json:"Pid,omitempty" name:"Pid"`
	// 价格

	Price *int64 `json:"Price,omitempty" name:"Price"`
	// 值

	Value *int64 `json:"Value,omitempty" name:"Value"`
	// 价格模式

	PriceModel *string `json:"PriceModel,omitempty" name:"PriceModel"`
	// 价格详情数组

	PriceDetail []*PriceDetail `json:"PriceDetail,omitempty" name:"PriceDetail"`
	// 总消费

	TotalCost *int64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 账单代码

	BillingItemCode *string `json:"BillingItemCode,omitempty" name:"BillingItemCode"`
	// 子账单代码

	SubBillingItemCode *string `json:"SubBillingItemCode,omitempty" name:"SubBillingItemCode"`
	// 实际总花费

	RealTotalCost *float64 `json:"RealTotalCost,omitempty" name:"RealTotalCost"`
	// 方案

	Policy *float64 `json:"Policy,omitempty" name:"Policy"`
	// 方案详情

	PolicyDetail *PolicyDetail `json:"PolicyDetail,omitempty" name:"PolicyDetail"`
}

type AuthorizeTokenResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AuthorizeTokenResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AuthorizeTokenResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaTypeConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的结果对象

		Result *InstanceTypeConfigResponse `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCkafkaTypeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaTypeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCkafkaZoneRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCkafkaZoneRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCkafkaZoneRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HourCreateRequest struct {
	*tchttp.BaseRequest

	// 后付费商品信息

	GoodsInfoList []*GoodsInfo `json:"GoodsInfoList,omitempty" name:"GoodsInfoList"`
}

func (r *HourCreateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HourCreateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TaskStatusResponse struct {

	// 任务状态:
	// 0 成功
	// 1 失败
	// 2 进行中

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 输出信息

	Output *string `json:"Output,omitempty" name:"Output"`
}

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

package v20170312

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type IdcData struct {

	// idc id,在腾讯云上代表zoneId

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// idc id,在腾讯云上代表zone名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type SetDashboardStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDashboardStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDashboardStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeUnitInfoOutput struct {

	// 指标单位ID

	UnitId *uint64 `json:"UnitId,omitempty" name:"UnitId"`
	// 指标单位名称

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type DescribeBaseMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询得到的指标描述列表

		MetricSet []*MetricObject `json:"MetricSet,omitempty" name:"MetricSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaseMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaseMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDashboardStatusRequest struct {
	*tchttp.BaseRequest

	// 是否第一次获取dashboard详情

	FirstDescribeDashboard *string `json:"FirstDescribeDashboard,omitempty" name:"FirstDescribeDashboard"`
	// 是否第一次添加精选

	FirstAddFeaturedDashboard *string `json:"FirstAddFeaturedDashboard,omitempty" name:"FirstAddFeaturedDashboard"`
	// 是否第一次获取精选指标

	FirstDescribeFeaturedDashboard *string `json:"FirstDescribeFeaturedDashboard,omitempty" name:"FirstDescribeFeaturedDashboard"`
	// 是否第一次查看图表详情

	FirstDescribePanel *string `json:"FirstDescribePanel,omitempty" name:"FirstDescribePanel"`
}

func (r *SetDashboardStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDashboardStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Xxx1 struct {

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// namespace的Id

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
	// 聚合方式

	AggType *string `json:"AggType,omitempty" name:"AggType"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 分类

	Category *string `json:"Category,omitempty" name:"Category"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type MetricSet struct {

	// 命名空间，每个云产品会有一个命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标使用的单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 指标使用的单位

	UnitCname *string `json:"UnitCname,omitempty" name:"UnitCname"`
	// 指标支持的统计周期，单位是秒，如60、300

	Period []*int64 `json:"Period,omitempty" name:"Period"`
	// 统计周期内指标方式

	Periods []*PeriodsSt `json:"Periods,omitempty" name:"Periods"`
	// 统计指标含义解释

	Meaning *MetricObjectMeaning `json:"Meaning,omitempty" name:"Meaning"`
	// 维度描述信息

	Dimensions []*DimensionsDesc `json:"Dimensions,omitempty" name:"Dimensions"`
}

type IdcInfoData struct {

	// 返回数据

	Data []*IdcInfo `json:"Data,omitempty" name:"Data"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AttributeIdOutput struct {

	// 属性ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控指标

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 数据点起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 数据点结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 数据统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 监控数据列表

		DataPoints []*PointsObject `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMonitorDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeAggrValueInfoOutputData struct {

	// 返回聚合数据列表

	Data []*AttributeTimestampValueOutput `json:"Data,omitempty" name:"Data"`
	// 返回聚合数据列表个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type IdcInfo struct {

	// 地域ID

	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`
	// 实例数

	ServerCount *int64 `json:"ServerCount,omitempty" name:"ServerCount"`
	// 地域名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type Metrics struct {

	// 指标ID

	MetricId *uint64 `json:"MetricId,omitempty" name:"MetricId"`
	// 指标的值

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type CreatePolicyGroupCondition struct {

	// 指标Id

	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`
	// 比较类型，范围0-6，分别对应[>,<,>=,<=,==,!=,!]。如果指标有配置默认比较类型值可以不填。

	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`
	// 比较的值，如果指标不必须CalcValue可不填

	CalcValue *float64 `json:"CalcValue,omitempty" name:"CalcValue"`
	// Storm检测周期单位秒，若指标有默认值可不填

	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`
	// 持续几个检测周期触发规则会告警

	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`
	// 告警发送收敛类型。0连续告警，1指数告警

	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次

	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
	// 如果通过模版创建，需要传入模版中该指标的对应RuleId

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
}

type History struct {

	// 告警源id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 告警源名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 主账号uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 子账号uin

	SubUin *int64 `json:"SubUin,omitempty" name:"SubUin"`
	// 告警id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 发生时间

	OccurTime *int64 `json:"OccurTime,omitempty" name:"OccurTime"`
	// 接收时间

	ReceiveTime *int64 `json:"ReceiveTime,omitempty" name:"ReceiveTime"`
	// 告警状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 接收组

	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`
	// 发送方式

	NotifyWays []*int64 `json:"NotifyWays,omitempty" name:"NotifyWays"`
}

type DimensionsDesc struct {

	// 维度名数组

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type CancelStarUnifyDashboardRequest struct {
	*tchttp.BaseRequest

	// dashboard uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

func (r *CancelStarUnifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelStarUnifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeInfoOutput struct {

	// 属性ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性级别

	AttributeLevel *int64 `json:"AttributeLevel,omitempty" name:"AttributeLevel"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 属性类型

	AttributeType *string `json:"AttributeType,omitempty" name:"AttributeType"`
	// 属性数据类型

	DataType *int64 `json:"DataType,omitempty" name:"DataType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 属性负责人

	OwnerName *string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 单位ID

	UnitId *int64 `json:"UnitId,omitempty" name:"UnitId"`
	// 单位名称

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
	// 统计周期

	StatisticalPeriod *int64 `json:"StatisticalPeriod,omitempty" name:"StatisticalPeriod"`
	// 属性唯一字符串，字母、数字、横杠或下划线组成

	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`
}

type AttributeUnitInfoOutputData struct {

	// 属性单位数据列表

	Data []*AttributeUnitInfoOutput `json:"Data,omitempty" name:"Data"`
	// 属性单位数据列表个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ReceiverInfo struct {

	// 有效时段结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 是否需要发送     通知

	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`
	// 告警接收渠道

	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`
	// 电话告警对个人间隔（秒）

	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`
	// 消息接收组列表

	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList"`
	// 接受者类型

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 接收人列表。通过平台接口查询到的接收人id列表

	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`
	// 告警恢复通知方式

	RecoverNotify *string `json:"RecoverNotify,omitempty" name:"RecoverNotify"`
	// 电话告警每轮间隔（秒）

	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`
	// 电话告警轮数

	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`
	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)

	SendFor []*string `json:"SendFor,omitempty" name:"SendFor"`
	// 有效时段开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 电话告警接收者uid

	UIDList []*int64 `json:"UIDList,omitempty" name:"UIDList"`
}

type DescribeBaseMetricsRequest struct {
	*tchttp.BaseRequest

	// 业务命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

func (r *DescribeBaseMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaseMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeInfoInput struct {

	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性类型

	AttributeTypeId *uint64 `json:"AttributeTypeId,omitempty" name:"AttributeTypeId"`
	// 属性数据类型

	DataType *int64 `json:"DataType,omitempty" name:"DataType"`
	// 属性级别

	AttributeLevel *int64 `json:"AttributeLevel,omitempty" name:"AttributeLevel"`
	// 属性单位

	Unit *int64 `json:"Unit,omitempty" name:"Unit"`
	// 统计周期

	StatisticalPeriod *int64 `json:"StatisticalPeriod,omitempty" name:"StatisticalPeriod"`
	// 负责人列表ID。默认为登录用户sub_uin

	OwnerId []*uint64 `json:"OwnerId,omitempty" name:"OwnerId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 属性唯一名称。由字母、数字、横杠或下划线组成

	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`
}

type AttributeInfoOutputData struct {

	// 属性数据列表

	Data []*AttributeInfoOutput `json:"Data,omitempty" name:"Data"`
	// 属性数据列表个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AttributeTimestampValueOutput struct {

	// 时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 属性值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type MetricObject struct {

	// 命名空间，每个云产品会有一个命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标使用的单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 指标支持的统计周期，单位是秒，如60、300

	Period []*int64 `json:"Period,omitempty" name:"Period"`
}

type NameType struct {

	// name

	Name *string `json:"Name,omitempty" name:"Name"`
	// id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

type AttributeServerInfoOutputData struct {

	// 返回上报该属性的服务器列表

	Data []*AttributeServerInfoOutput `json:"Data,omitempty" name:"Data"`
	// 返回上报该属性的服务器列表个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type PointsObject struct {

	// 监控实例的维度组合

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 监控数据点数组，每个点的时间跨度为一个Period值

	Points []*float64 `json:"Points,omitempty" name:"Points"`
}

type AttributeServerInfoOutput struct {

	// 服务器ID

	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`
	// 服务器名称

	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
	// 服务器ip

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 地域ID

	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`
	// 地域名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type MetricObjectMeaning struct {

	// 指标英文解释

	En *string `json:"En,omitempty" name:"En"`
	// 指标中文解释

	Zh *string `json:"Zh,omitempty" name:"Zh"`
}

type DescribePolicyConditionListConfigManualCalcValue struct {

	// 默认值

	Default *float64 `json:"Default,omitempty" name:"Default"`
	// 固定值

	Fixed *float64 `json:"Fixed,omitempty" name:"Fixed"`
	// 最大值

	Max *float64 `json:"Max,omitempty" name:"Max"`
	// 最小值

	Min *float64 `json:"Min,omitempty" name:"Min"`
	// 是否必须

	Need *bool `json:"Need,omitempty" name:"Need"`
}

type ModifyConditionsTemplateRequestEventCondition struct {

	// 告警通知周期

	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
	// 告警通知方式

	AlarmNotifyType *string `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 事件ID

	EventID *string `json:"EventID,omitempty" name:"EventID"`
	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 命名空间，每个云产品会有一个命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 监控统计周期。默认为取值为300，单位为s

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 起始时间，如 2018-01-01 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认为当前时间。 endTime不能小于startTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例对象的维度组合

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 统计方式

	Statistics *string `json:"Statistics,omitempty" name:"Statistics"`
}

func (r *GetMonitorDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonitorDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PeriodsSt struct {

	// 周期

	Period *string `json:"Period,omitempty" name:"Period"`
	// 统计方式

	StatType []*string `json:"StatType,omitempty" name:"StatType"`
}

type AttributeValueInfoOutputData struct {

	// 属性上报数据列表

	Data []*AttributeValueInfoOutput `json:"Data,omitempty" name:"Data"`
	// 属性数据列表个数个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AttributeIdDeleteOutput struct {

	// 删除个数

	DeleteCount *int64 `json:"DeleteCount,omitempty" name:"DeleteCount"`
	// 属性ID

	AttributeId []*uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

type AlertField struct {

	// 自定义告警字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 自定义告警值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CatTaskDetail struct {

	// 任务Id

	TaskId *uint64 `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`
	// 任务周期，单位为分钟。目前支持1，5，15，30几种取值

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 拨测类型。http, https, ping, tcp 之一

	CatTypeName *string `json:"CatTypeName,omitempty" name:"CatTypeName"`
	// 任务状态。1表示暂停，2表示运行中，0为初始态

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 拨测任务的Url

	CgiUrl *string `json:"CgiUrl,omitempty" name:"CgiUrl"`
	// 任务创建时间

	AddTime *string `json:"AddTime,omitempty" name:"AddTime"`
	// 拨测分组id

	AgentGroupId *uint64 `json:"AgentGroupId,omitempty" name:"AgentGroupId"`
	// 告警策略组id

	PolicyGroupId *uint64 `json:"PolicyGroupId,omitempty" name:"PolicyGroupId"`
}

type LogFilterRule struct {

	// 关系 1AND 2OR

	Relation *string `json:"Relation,omitempty" name:"Relation"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 条目ID（只作出参，入参不填）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 指标集ID（只作出参，入参不填）

	MetricSetId *string `json:"MetricSetId,omitempty" name:"MetricSetId"`
}

type AttributeValueInfoOutput struct {

	// 属性值列表

	Values []*AttributeTimestampValueOutput `json:"Values,omitempty" name:"Values"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 开始位置

	StartPosition *int64 `json:"StartPosition,omitempty" name:"StartPosition"`
	// 结束位置

	EndPosition *int64 `json:"EndPosition,omitempty" name:"EndPosition"`
}

type CancelStarUnifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelStarUnifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelStarUnifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

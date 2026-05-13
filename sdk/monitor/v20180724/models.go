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

package v20180724

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type DescribePolicyQuotaV3ResponseData struct {

	// 视图

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 是否为项目配额

	IsQuotaByProject *bool `json:"IsQuotaByProject,omitempty" name:"IsQuotaByProject"`
	// 配额列表

	QuotaList []*DescribePolicyQuotaV3ResponseDataQuotaList `json:"QuotaList,omitempty" name:"QuotaList"`
	// 配额

	Quota *DescribePolicyQuotaV3ResponseDataQuota `json:"Quota,omitempty" name:"Quota"`
}

type CreateConditionsTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板策略组ID

		GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConditionsTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConditionsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果数据json

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标集详情

		Data *ClmMetricSet `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警通知模板 id

	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`
}

func (r *DescribeAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategyStatesInfo struct {

	// 告警策略ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 启停状态。默认为1。 0：停止中， 1：启动中

	IsStart *int64 `json:"IsStart,omitempty" name:"IsStart"`
}

type ModifyMenuUinListRequest struct {
	*tchttp.BaseRequest

	// 菜单名

	Menu *string `json:"Menu,omitempty" name:"Menu"`
	// 进行操作的uin列表

	UinList []*string `json:"UinList,omitempty" name:"UinList"`
	// 操作名：up | down

	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *ModifyMenuUinListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMenuUinListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StarUnifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StarUnifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StarUnifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomAlarmList struct {

	// IP

	LocalIP *string `json:"LocalIP,omitempty" name:"LocalIP"`
	// 消息

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 发生时间

	OccurTime *string `json:"OccurTime,omitempty" name:"OccurTime"`
	// 调用方

	Caller *string `json:"Caller,omitempty" name:"Caller"`
}

type MidQueryCondition struct {

	// 维度

	Key *string `json:"Key,omitempty" name:"Key"`
	// 操作符，支持等于(eq)、不等于(ne)，以及in

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 维度值，当Op是eq、ne时，只使用第一个元素

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DescribeAttributeAllServerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据列表

		Data *AttributeServerInfoOutputData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttributeAllServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttributeAllServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnifyDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dashboard列表

		Dashboards []*UnifyDashboardMeta `json:"Dashboards,omitempty" name:"Dashboards"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnifyDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnifyDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductByIdsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 产品Id列表

	ProductIds []*string `json:"ProductIds,omitempty" name:"ProductIds"`
}

func (r *DescribeMonitorProductByIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorProductByIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmNotice struct {

	// 告警通知模板 ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 告警通知模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 上次修改时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 上次修改人

	UpdatedBy *string `json:"UpdatedBy,omitempty" name:"UpdatedBy"`
	// 告警通知类型 ALARM=未恢复通知 OK=已恢复通知 ALL=全部通知

	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`
	// 用户通知列表

	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices"`
	// 回调通知列表

	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`
	// 是否是系统预设通知模板 0=否 1=是

	IsPreset *int64 `json:"IsPreset,omitempty" name:"IsPreset"`
	// 通知语言 zh-CN=中文 en-US=英文

	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`
	// 告警通知模板绑定的告警策略ID列表

	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
	// 后台 amp consumer id

	AMPConsumerId *string `json:"AMPConsumerId,omitempty" name:"AMPConsumerId"`
	// 其他非公开通知渠道

	ExtraChannels []*string `json:"ExtraChannels,omitempty" name:"ExtraChannels"`
	// BigCustomerNotices

	BigCustomerNotices []*BigCustomerNotice `json:"BigCustomerNotices,omitempty" name:"BigCustomerNotices"`
	// CLS通知渠道列表

	CLSNotices []*CLSNotice `json:"CLSNotices,omitempty" name:"CLSNotices"`
}

// CLSNotice CLS通知渠道
type CLSNotice struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 启用/停用 0=停用 1=启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

type AttributeIdOutput struct {

	// 属性ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

type DescribeBindingPolicyObjectListInstanceGroup struct {

	// 实例分组id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 告警策略类型名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 最后编辑uin

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 实例分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 实例数量

	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 实例所在的地域集合

	Regions []*string `json:"Regions,omitempty" name:"Regions"`
}

type DescribeAlarmHistoriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 告警历史列表

		Histories []*AlarmHistory `json:"Histories,omitempty" name:"Histories"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmHistoriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmHistoriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AmpFrontTunnelRequest struct {
	*tchttp.BaseRequest

	// HTTP请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// HTTP请求地址

	Path *string `json:"Path,omitempty" name:"Path"`
	// HTTP请求体

	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
}

func (r *AmpFrontTunnelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AmpFrontTunnelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyInstanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyInstanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyInstanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmCallbackRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 更新动作, "bind"表示绑定回调地址, "unbind"表示解绑回调地址

	UserAction *string `json:"UserAction,omitempty" name:"UserAction"`
	// 用户回调地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 验证码

	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
	// 是否为新增策略组, 1表示新增的策略组, 0表示已有的策略组

	NewPolicy *int64 `json:"NewPolicy,omitempty" name:"NewPolicy"`
}

func (r *ModifyAlarmCallbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmCallbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListConfigManualContinueTime struct {

	// 默认持续时间，单位：秒

	Default *int64 `json:"Default,omitempty" name:"Default"`
	// 可选持续时间，单位：秒

	Keys []*int64 `json:"Keys,omitempty" name:"Keys"`
	// 是否必须

	Need *bool `json:"Need,omitempty" name:"Need"`
}

type ModifyQuestionnaireResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyQuestionnaireResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyQuestionnaireResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicyRule struct {

	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 聚合维度列表

	AggregationDimensions []*string `json:"AggregationDimensions,omitempty" name:"AggregationDimensions"`
	// 秒数 统计周期

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 数据聚合类型 因复合条件可能涉及多个指标 这里用 json 表示 {"report_metric_gauge":"last"}

	AggregateType *string `json:"AggregateType,omitempty" name:"AggregateType"`
	// 英文运算符
	// intelligent=无阈值智能检测
	// eq=等于
	// ge=大于等于
	// gt=大于
	// le=小于等于
	// lt=小于
	// ne=不等于
	// day_increase=天同比增长
	// day_decrease=天同比下降
	// day_wave=天同比波动
	// week_increase=周同比增长
	// week_decrease=周同比下降
	// week_wave=周同比波动
	// cycle_increase=环比增长
	// cycle_decrease=环比下降
	// cycle_wave=环比波动
	// re=正则匹配

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 阈值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 周期数 持续通知周期 1=持续1个周期 2=持续2个周期...

	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`
	// 秒数 告警间隔  0=不重复 300=每5分钟告警一次 600=每10分钟告警一次 900=每15分钟告警一次 1800=每30分钟告警一次 3600=每1小时告警一次 7200=每2小时告警一次 10800=每3小时告警一次 21600=每6小时告警一次 43200=每12小时告警一次 86400=每1天告警一次

	NoticeFrequency *int64 `json:"NoticeFrequency,omitempty" name:"NoticeFrequency"`
	// 告警频率是否指数增长 0=否 1=是

	IsPowerNotice *int64 `json:"IsPowerNotice,omitempty" name:"IsPowerNotice"`
	// 对于单个触发规则的过滤条件

	Filter *AlarmPolicyFilter `json:"Filter,omitempty" name:"Filter"`
	// 指标展示名，用于出参

	Description *string `json:"Description,omitempty" name:"Description"`
	// 单位，用于出参

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 触发条件类型 STATIC=静态阈值 DYNAMIC=动态阈值

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// ValueMax

	ValueMax *int64 `json:"ValueMax,omitempty" name:"ValueMax"`
	// ValueMin

	ValueMin *int64 `json:"ValueMin,omitempty" name:"ValueMin"`
}

type DescribeInstanceGroupListInstanceGroup struct {

	// 实例分组id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 实例分组策略类型名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 最近更新人uin

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 实例分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 分组内的实例数

	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 更新时间，unix时间戳

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间，unix时间戳

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 绑定的告警策略组列表

	PolicyGroups []*DescribeInstanceGroupListPolicyGroup `json:"PolicyGroups,omitempty" name:"PolicyGroups"`
}

type UnBindingInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 实例组Id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 策略组Id

	PolicyGroupId *int64 `json:"PolicyGroupId,omitempty" name:"PolicyGroupId"`
	// 是否删除策略组，0表示否，1表示是

	IsDelRelatedPolicy *int64 `json:"IsDelRelatedPolicy,omitempty" name:"IsDelRelatedPolicy"`
}

func (r *UnBindingInstanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindingInstanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCMInstanceAttributeDataOutput struct {

	// 指标ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 服务器ID

	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`
	// 指标上报数据

	Values []*AttributeTimestampValueOutput `json:"Values,omitempty" name:"Values"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

type StrategyChannelInfo struct {

	// 告警策略ID

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 告警渠道列表

	AlarmChannel []*string `json:"AlarmChannel,omitempty" name:"AlarmChannel"`
}

type DescribeBasicAlarmListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前取值monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 起始时间，默认一天前的时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认当前时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 根据发生时间排序，取值ASC或DESC

	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`
	// 根据项目ID过滤

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
	// 根据策略类型过滤

	ViewNames []*string `json:"ViewNames,omitempty" name:"ViewNames"`
	// 根据告警状态过滤

	AlarmStatus []*int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
	// 根据告警对象过滤

	ObjLike *string `json:"ObjLike,omitempty" name:"ObjLike"`
	// 根据实例组ID过滤

	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitempty" name:"InstanceGroupIds"`
}

func (r *DescribeBasicAlarmListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicAlarmListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicyTriggerTask struct {

	// 触发任务类型 AS=弹性伸缩

	Type *string `json:"Type,omitempty" name:"Type"`
	// 用 json 表示配置信息 {"Key1":"Value1","Key2":"Value2"}

	TaskConfig *string `json:"TaskConfig,omitempty" name:"TaskConfig"`
}

type DescribeAccidentConfigRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAccidentConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccidentConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecoverNotifyBatchNotifyInfo struct {

	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 告警恢复通知方式列表

	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify"`
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

type MetricAggregation struct {

	// 中文名

	NameCn *string `json:"NameCn,omitempty" name:"NameCn"`
	// 英文名

	NameEn *string `json:"NameEn,omitempty" name:"NameEn"`
	// Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeCvmAgentStatusRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// CVM实例列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeCvmAgentStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmAgentStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteGroupRequest struct {
	*tchttp.BaseRequest

	// 分组Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcServerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *AttributeServerInfoOutputData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultAlarmPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultAlarmPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMLogProfileItem struct {

	// 关联指标集CID（只做出参，入参不填）

	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 字段类型，可选：string、long、double

	Type *string `json:"Type,omitempty" name:"Type"`
	// 创建时间（只做出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只做出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DescribeProductEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件列表

		Events []*DescribeProductEventListEvents `json:"Events,omitempty" name:"Events"`
		// 事件统计

		OverView *DescribeProductEventListOverView `json:"OverView,omitempty" name:"OverView"`
		// 事件总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsForConsoleFontEndData struct {

	// 名字空间

	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标中文名

	MetricCNName *string `json:"MetricCNName,omitempty" name:"MetricCNName"`
	// 指标英文名

	MetricENName *string `json:"MetricENName,omitempty" name:"MetricENName"`
	// 维度

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 单位中午名

	UnitCNName *string `json:"UnitCNName,omitempty" name:"UnitCNName"`
	// 周期与统计方式

	PeriodAndStatTypes []*DescribeBaseMetricsDataPeriod `json:"PeriodAndStatTypes,omitempty" name:"PeriodAndStatTypes"`
	// 指标含义

	Meaning []*DescribeBaseMetricsDataMeaning `json:"Meaning,omitempty" name:"Meaning"`
}

type DescribeMonitorProductsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeMonitorProductsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorProductsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalObjectsThresholdObject struct {

	// 告警id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 告警规则id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 项目id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 告警状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 告警内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 首次发生时间

	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`
	// 最后发生时间

	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`
	// 维度组合，json字符串

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 维度组合md5

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
	// 告警状态，“abnormal”表示告警未恢复，“warning”表示告警已恢复

	AbnormalStatus *string `json:"AbnormalStatus,omitempty" name:"AbnormalStatus"`
	// 对象id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribePolicyGroupInfoCondition struct {

	// 指标名称

	MetricShowName *string `json:"MetricShowName,omitempty" name:"MetricShowName"`
	// 数据聚合周期(单位秒)

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 指标id

	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`
	// 阈值规则id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 指标单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 告警发送收敛类型。0连续告警，1指数告警

	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次

	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等，7表示日同比上涨，8表示日同比下降，9表示周同比上涨，10表示周同比下降，11表示周期环比上涨，12表示周期环比下降

	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`
	// 检测阈值

	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`
	// 持续多长时间触发规则会告警(单位秒)

	ContinueTime *int64 `json:"ContinueTime,omitempty" name:"ContinueTime"`
}

type DescribePolicyConditionListConfigManualStatType struct {

	// 数据聚合方式，周期5秒

	P5 *string `json:"P5,omitempty" name:"P5"`
	// 数据聚合方式，周期10秒

	P10 *string `json:"P10,omitempty" name:"P10"`
	// 数据聚合方式，周期1分钟

	P60 *string `json:"P60,omitempty" name:"P60"`
	// 数据聚合方式，周期5分钟

	P300 *string `json:"P300,omitempty" name:"P300"`
	// 数据聚合方式，周期10分钟

	P600 *string `json:"P600,omitempty" name:"P600"`
	// 数据聚合方式，周期30分钟

	P1800 *string `json:"P1800,omitempty" name:"P1800"`
	// 数据聚合方式，周期1小时

	P3600 *string `json:"P3600,omitempty" name:"P3600"`
	// 数据聚合方式，周期1天

	P86400 *string `json:"P86400,omitempty" name:"P86400"`
}

type DescribePolicyInfoByInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// UniqueID

		UniqueID *string `json:"UniqueID,omitempty" name:"UniqueID"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 策略组信息

		GroupList []*DescribePolicyInfoByInstanceGroupList `json:"GroupList,omitempty" name:"GroupList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyInfoByInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyInfoByInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmHistoryMetric struct {

	// 查询数据使用的命名空间

	QceNamespace *string `json:"QceNamespace,omitempty" name:"QceNamespace"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 统计周期

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 触发告警的数值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 指标的解释

	Tooltip *MetricTooltip `json:"Tooltip,omitempty" name:"Tooltip"`
	// 指标的展示名

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeAlarmSmsQuotaQuota struct {

	// 配额类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 配额名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 免费配额剩余量

	FreeLeft *int64 `json:"FreeLeft,omitempty" name:"FreeLeft"`
	// 付费配额剩余量

	PurchaseLeft *int64 `json:"PurchaseLeft,omitempty" name:"PurchaseLeft"`
	// 已使用量

	Used *int64 `json:"Used,omitempty" name:"Used"`
}

type ModifyAlarmPolicyNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyChannelsRequest struct {
	*tchttp.BaseRequest

	// 告警渠道列表

	Strategy []*StrategyChannelInfo `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *ModifyStrategyChannelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrategyChannelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteStrategysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteStrategysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStrategysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云产品的名字空间

		QceNamespaces []*CommonNamespace `json:"QceNamespaces,omitempty" name:"QceNamespaces"`
		// 自定义监控的命名空间

		CustomNamespaces []*CommonNamespace `json:"CustomNamespaces,omitempty" name:"CustomNamespaces"`
		// CommonNamespaces

		CommonNamespaces []*string `json:"CommonNamespaces,omitempty" name:"CommonNamespaces"`
		// CustomNamespacesNew

		CustomNamespacesNew []*string `json:"CustomNamespacesNew,omitempty" name:"CustomNamespacesNew"`
		// QceNamespacesNew

		QceNamespacesNew []*DashboardNamespace `json:"QceNamespacesNew,omitempty" name:"QceNamespacesNew"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标集列表结果数据

		Data *ClmDescribeMetricSetsData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardDataSource struct {

	// DataSource 唯一标示

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 数据源名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据源类型，如MySQL，Prometheus

	Type *string `json:"Type,omitempty" name:"Type"`
	// 数据源图标

	TypeLogoUrl *string `json:"TypeLogoUrl,omitempty" name:"TypeLogoUrl"`
	// 访问 模式

	Access *string `json:"Access,omitempty" name:"Access"`
	// 数据源访问地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 密码，可以为空字符串

	Password *string `json:"Password,omitempty" name:"Password"`
	// 用户名，可以为空字符串

	User *string `json:"User,omitempty" name:"User"`
	// 数据库

	Database *string `json:"Database,omitempty" name:"Database"`
	// 默认为false

	BasicAuth *bool `json:"BasicAuth,omitempty" name:"BasicAuth"`
	// 是否为默认数据源

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 存储数据源的附加信息，如ES的链接信息"{ \"esVersion\": 5, \"logLevelField\": \"\",\"logMessageField\": \"\", \"maxConcurrentShardRequests\": 256, \"timeField\": \"@timestamp\"}"

	JsonData *string `json:"JsonData,omitempty" name:"JsonData"`
	// 数据源数据是否只读

	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

type CustomAlarmInfo struct {

	// 调用方

	Caller *string `json:"Caller,omitempty" name:"Caller"`
	// IP

	LocalIP *string `json:"LocalIP,omitempty" name:"LocalIP"`
	// 消息内容

	Msg *string `json:"Msg,omitempty" name:"Msg"`
	// 发生时间

	OccurTime *string `json:"OccurTime,omitempty" name:"OccurTime"`
}

type DashboardMetricData struct {

	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 聚合方式

	Aggregate *string `json:"Aggregate,omitempty" name:"Aggregate"`
	// 查询条件数组

	GroupBy []*Dimension `json:"GroupBy,omitempty" name:"GroupBy"`
	// 返回Timestamp数组

	Timestamps []*uint64 `json:"Timestamps,omitempty" name:"Timestamps"`
	// 指标上报数据，返回json序列化字符串，数组形式

	Value *string `json:"Value,omitempty" name:"Value"`
	// 返回查询时间段的开始时间，使用ISO格式 如：2018-04-10T10:15:58.858Z

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 返回查询时间段的开始时间，使用ISO格式 如：2018-04-10T10:15:58.858Z

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 统计周期，单位为秒

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 过滤条件

	Conditions []*DashboardMetricCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
	// 错误码

	Code *string `json:"Code,omitempty" name:"Code"`
	// 查询标识

	SeriesId *string `json:"SeriesId,omitempty" name:"SeriesId"`
	// 产品配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 实例的维度组合

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 排序值

	OrderTag *float64 `json:"OrderTag,omitempty" name:"OrderTag"`
	// 类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

type ReceiverInfo struct {

	// Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 告警时间段开始时间。范围[0,86400)，作为unix时间戳转成北京时间后去掉日期，例如7200表示"10:0:0"

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 告警时间段结束时间。含义同StartTime

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 告警通知方式。可选 "SMS","SITE","EMAIL","CALL","WECHAT"

	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`
	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)

	SendFor []*string `json:"SendFor,omitempty" name:"SendFor"`
	// 电话告警接收者uid

	UidList []*int64 `json:"UidList,omitempty" name:"UidList"`
	// 电话告警轮数

	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`
	// 电话告警对个人间隔（秒）

	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`
	// 电话告警每轮间隔（秒）

	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`
	// 恢复通知方式。可选"SMS"

	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify"`
	// 是否需要电话告警触达提示。0不需要，1需要

	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`
	// 接收人类型。“group” 或 “user”

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 接收组列表。通过平台接口查询到的接收组id列表

	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList"`
	// 接收人列表。通过平台接口查询到的接收人id列表

	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`
	// 接收语言，取值为 en-US 和 zh-CN

	ReceiveLanguage *string `json:"ReceiveLanguage,omitempty" name:"ReceiveLanguage"`
}

type DescribeMenuWhitelistResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 白名单

		Whitelist []*string `json:"Whitelist,omitempty" name:"Whitelist"`
		// 当前使用的用户名单

		CurrentUinList []*string `json:"CurrentUinList,omitempty" name:"CurrentUinList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMenuWhitelistResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMenuWhitelistResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 组Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 组类型，1表示策略组，2表示实例组，3表示条件模板组

	GroupType *int64 `json:"GroupType,omitempty" name:"GroupType"`
	// 修改的字段名，可选：groupName、remark，分别表示修改策略名、描述

	Key *string `json:"Key,omitempty" name:"Key"`
	// 新的字段值

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyPolicyGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolicyGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Contact struct {

	// uid

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
	// uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 接收人名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否可登录,0表示不可登录,1表示可登录

	CanLogin *int64 `json:"CanLogin,omitempty" name:"CanLogin"`
	// 手机号是否验证通过

	PhoneFlag *int64 `json:"PhoneFlag,omitempty" name:"PhoneFlag"`
	// 手机号

	PhoneNum *string `json:"PhoneNum,omitempty" name:"PhoneNum"`
	// 邮箱是否验证通过

	EmailFlag *int64 `json:"EmailFlag,omitempty" name:"EmailFlag"`
	// 邮箱号

	Email *string `json:"Email,omitempty" name:"Email"`
	// 是否是接收负责人

	IsReceiverOwner *int64 `json:"IsReceiverOwner,omitempty" name:"IsReceiverOwner"`
	// ownerUid

	OwnerUid *int64 `json:"OwnerUid,omitempty" name:"OwnerUid"`
}

type PCLMMetricAnalysisPoint struct {

	// 时间戳（秒）

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type StrategyIdInfo struct {

	// 告警策略ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

type GetArgusMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 无

	ResourceId *uint64 `json:"ResourceId,omitempty" name:"ResourceId"`
	// 无

	Query *string `json:"Query,omitempty" name:"Query"`
	// 无

	Epoch *string `json:"Epoch,omitempty" name:"Epoch"`
	// 无

	IsEncrypt *uint64 `json:"IsEncrypt,omitempty" name:"IsEncrypt"`
}

func (r *GetArgusMonitorDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetArgusMonitorDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type ModifyNotifyBatchNotifyInfo struct {

	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 告警方式列表

	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`
}

type ModifyMetricSetRequest struct {
	*tchttp.BaseRequest

	// 指标集ID

	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`
	// 指标集详情

	MetricSet *ClmMetricSet `json:"MetricSet,omitempty" name:"MetricSet"`
}

func (r *ModifyMetricSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMetricSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 列表总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 策略列表

		PolicyList []*MessagePolicy `json:"PolicyList,omitempty" name:"PolicyList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsgPolicyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMsgPolicyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控面板ID

		DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategyData struct {

	// 返回数据

	Data []*StrategyEntry `json:"Data,omitempty" name:"Data"`
	// 查询数据列表总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type PCLMDimension struct {

	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 指标集CID（只作出参，入参不填）

	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
}

type DescribeLogSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果数据json

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMAlertFilterRule struct {

	// 关联的告警策略CID

	AlertPolicyCID *string `json:"AlertPolicyCID,omitempty" name:"AlertPolicyCID"`
	// 过滤的关键字

	Key *string `json:"Key,omitempty" name:"Key"`
	// 操作符

	Operating *string `json:"Operating,omitempty" name:"Operating"`
	// 筛选值（多个值逗号分隔）

	Value *string `json:"Value,omitempty" name:"Value"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DescribeDashboardViewList struct {

	// 视图ID

	ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`
	// 描述名

	DescName *string `json:"DescName,omitempty" name:"DescName"`
	// 名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名列表

	MetricName []*string `json:"MetricName,omitempty" name:"MetricName"`
	// 元数据json

	Meta *string `json:"Meta,omitempty" name:"Meta"`
	// 实例列表

	Instances []*string `json:"Instances,omitempty" name:"Instances"`
}

type DescribePolicyConditionListMetric struct {

	// 指标配置

	ConfigManual *DescribePolicyConditionListConfigManual `json:"ConfigManual,omitempty" name:"ConfigManual"`
	// 指标id

	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`
	// 指标名称

	MetricShowName *string `json:"MetricShowName,omitempty" name:"MetricShowName"`
	// 指标单位

	MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`
}

type TemplateGroup struct {

	// 指标告警规则

	Conditions []*Condition `json:"Conditions,omitempty" name:"Conditions"`
	// 事件告警规则

	EventConditions []*EventCondition `json:"EventConditions,omitempty" name:"EventConditions"`
	// 关联告警策略组

	PolicyGroups []*PolicyGroup `json:"PolicyGroups,omitempty" name:"PolicyGroups"`
	// 模板策略组ID

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 模板策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 最后修改人UIN

	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 视图

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 是否为与关系

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type ShieldPolicyAlarmRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 告警屏蔽状态，1表示屏蔽，0表示不屏蔽

	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`
	// 设置指定实例的屏蔽状态

	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *ShieldPolicyAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShieldPolicyAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonNamespace struct {

	// 命名空间标示

	Id *string `json:"Id,omitempty" name:"Id"`
	// 命名空间名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 配置信息

	Config *string `json:"Config,omitempty" name:"Config"`
	// 支持地域列表

	AvailableRegions []*string `json:"AvailableRegions,omitempty" name:"AvailableRegions"`
	// 排序Id

	SortId *int64 `json:"SortId,omitempty" name:"SortId"`
	// Dashboard中的唯一表示

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
}

type DashboardDimensionSource struct {

	// 数据源

	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 维度名

	DimensionName *string `json:"DimensionName,omitempty" name:"DimensionName"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
}

type DescribeBasicAlarmListAlarms struct {

	// 该条告警的ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 项目ID

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名称

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 告警状态ID

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 告警状态

	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
	// 策略组ID

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 发生时间

	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`
	// 持续时间，单位s

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// 结束时间

	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`
	// 告警内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 告警对象

	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`
	// 告警对象ID

	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`
	// 策略类型

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// VPC，只有CVM有

	Vpc *string `json:"Vpc,omitempty" name:"Vpc"`
	// 指标ID

	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 告警类型，0表示指标告警，2表示产品事件告警，3表示平台事件告警

	AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 告警对象维度信息

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 通知方式

	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`
	// 所属实例组信息

	InstanceGroup []*InstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`
}

type DescribeEventPolicyConfigDimension struct {

	// 维度名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度key值

	Key *string `json:"Key,omitempty" name:"Key"`
}

type DescribePolicyConditionListConfigManual struct {

	// 检测方式

	CalcType *DescribePolicyConditionListConfigManualCalcType `json:"CalcType,omitempty" name:"CalcType"`
	// 检测阈值

	CalcValue *DescribePolicyConditionListConfigManualCalcValue `json:"CalcValue,omitempty" name:"CalcValue"`
	// 持续时间

	ContinueTime *DescribePolicyConditionListConfigManualContinueTime `json:"ContinueTime,omitempty" name:"ContinueTime"`
	// 数据周期

	Period *DescribePolicyConditionListConfigManualPeriod `json:"Period,omitempty" name:"Period"`
	// 持续周期个数

	PeriodNum *DescribePolicyConditionListConfigManualPeriodNum `json:"PeriodNum,omitempty" name:"PeriodNum"`
	// 聚合方式

	StatType *DescribePolicyConditionListConfigManualStatType `json:"StatType,omitempty" name:"StatType"`
}

type MetricTooltip struct {

	// 文本内容

	Message *string `json:"Message,omitempty" name:"Message"`
	// 详情链接

	URL *string `json:"URL,omitempty" name:"URL"`
	// 链接名称

	UrlName *string `json:"UrlName,omitempty" name:"UrlName"`
}

type DeleteDashboardViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
	// 主账号 uid 用于创建预设通知

	OwnerUid *int64 `json:"OwnerUid,omitempty" name:"OwnerUid"`
	// 告警通知模板名称 用来模糊搜索

	Name *string `json:"Name,omitempty" name:"Name"`
	// 根据接收人过滤告警通知模板需要选定通知用户类型 USER=用户 GROUP=用户组 传空=不按接收人过滤

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 接收对象列表

	UserIds []*int64 `json:"UserIds,omitempty" name:"UserIds"`
	// 接收组列表

	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// 页码 最小为1

	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 分页大小 1～200

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 按更新时间排序方式 ASC=正序 DESC=倒序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 根据后台 amp consumer id 列表过滤，传空不过滤

	AMPConsumerIds []*string `json:"AMPConsumerIds,omitempty" name:"AMPConsumerIds"`
	// 根据通知模板 id 过滤，空数组/不传则不过滤

	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcServerRequest struct {
	*tchttp.BaseRequest

	// 地域ID列表

	IdcId []*int64 `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *DescribeIdcServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcServerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddDashboardPanelsRequest struct {
	*tchttp.BaseRequest

	// 源列表

	SrcList *DashboardPanel `json:"SrcList,omitempty" name:"SrcList"`
	// 目标

	Dst *DashboardPanel `json:"Dst,omitempty" name:"Dst"`
}

func (r *AddDashboardPanelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDashboardPanelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CgrpGroupNode struct {

	// 组ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 组深度

	High *int64 `json:"High,omitempty" name:"High"`
	// 分组ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 是否是叶子节点

	IsLeaf *bool `json:"IsLeaf,omitempty" name:"IsLeaf"`
	// 当前在树中的层级

	Level *int64 `json:"Level,omitempty" name:"Level"`
	// 树节点

	Name *string `json:"Name,omitempty" name:"Name"`
	// 父节点Id

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// SubUin

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 允许叶子节点存放实例类型

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type DescribeStorageDurationData struct {

	// 周期

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 持续时间

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
}

type DescribeAlarmPoliciesRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 页数，从 1 开始计数，默认 1

	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页的数量，取值1~100，默认20

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 按策略名称模糊搜索

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 根据监控类型过滤 不选默认查所有类型 "MT_QCE"=云产品监控

	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`
	// 根据命名空间过滤

	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
	// 根据策略类型（viewName）过滤 （云产品监控用）

	Measurements []*string `json:"Measurements,omitempty" name:"Measurements"`
	// 告警对象列表

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 根据接收人搜索

	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids"`
	// 根据接收组搜索

	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`
	// 根据默认策略筛选 不传展示全部策略 DEFAULT=展示默认策略 NOT_DEFAULT=展示非默认策略

	PolicyType []*string `json:"PolicyType,omitempty" name:"PolicyType"`
	// 排序字段

	Field *string `json:"Field,omitempty" name:"Field"`
	// 排序顺序：升序：ASC  降序：DESC

	Order *string `json:"Order,omitempty" name:"Order"`
	// 项目id数组

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
	// 字段废弃，启停，1：启用   0：停止

	AlarmStatus []*int64 `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
	// 告警通知id列表

	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
	// 根据触发条件筛选 不传展示全部策略 STATIC=展示静态阈值策略 DYNAMIC=展示动态阈值策略

	RuleTypes []*string `json:"RuleTypes,omitempty" name:"RuleTypes"`
	// 启停，1：启用   0：停止

	Enable []*int64 `json:"Enable,omitempty" name:"Enable"`
	// 是否未配置通知规则，1：未配置，0：配置

	NotBindingNoticeRule *int64 `json:"NotBindingNoticeRule,omitempty" name:"NotBindingNoticeRule"`
	// 实例组id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 条件模板id

	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`
	// 模糊匹配

	Like *string `json:"Like,omitempty" name:"Like"`
	// 限制

	Limit *string `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAlarmPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConditionsTemplateListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 模板策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 模板策略组ID

	GroupID *string `json:"GroupID,omitempty" name:"GroupID"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 按更新时间排序方式，枚举值(asc, desc)

	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`
}

func (r *DescribeConditionsTemplateListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConditionsTemplateListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardPanelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result []*string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardPanelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardPanelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeCallbacksRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmNoticeCallbacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticeCallbacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeUnitsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *AttributeUnitInfoOutputData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttributeUnitsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttributeUnitsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerDatasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*ServerAttributeDataOutput `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerDatasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerDatasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控面板ID

		DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributeRequest struct {
	*tchttp.BaseRequest

	// 属性ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

func (r *DeleteAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAbnormalObjectsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警类型名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 起始时间，格式:"2019-01-01"

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 项目id列表

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
	// 告警状态，"all"表示所有实例，“abnormal”表示异常实例（告警未恢复），“warning”表示提醒实例（告警已恢复）

	AbnormalStatus *string `json:"AbnormalStatus,omitempty" name:"AbnormalStatus"`
	// 是否汇聚，不填表示不汇聚，"instance"表示按实例汇聚

	GroupBy *string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 按维度过滤，json字符串

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *DescribeAbnormalObjectsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalObjectsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMetricSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标集ID

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMetricSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMetricSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警指标列表

		Metrics []*Metric `json:"Metrics,omitempty" name:"Metrics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CLMDescribeMetricSetsRequest struct {
	*tchttp.BaseRequest

	// 指标集CID

	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
	// 指标集名称（模糊搜索）

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志集名称（模糊搜索）

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题名称（模糊搜索）

	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`
	// 偏移量，默认为0

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回数量，默认为20，最大100

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *CLMDescribeMetricSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CLMDescribeMetricSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyListRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 过滤条件

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMsgPolicyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMsgPolicyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMsgPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略id

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 接收用户组id列表

	ReceiverGroupIds []*int64 `json:"ReceiverGroupIds,omitempty" name:"ReceiverGroupIds"`
	// 告警方式列表

	NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays"`
	// 电话告警配置

	VoiceConfig *ModifyMsgPolicyVoiceConfig `json:"VoiceConfig,omitempty" name:"VoiceConfig"`
}

func (r *ModifyMsgPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMsgPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusMetricOut struct {

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

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeNewUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否是新用户

		IsNewUser *bool `json:"IsNewUser,omitempty" name:"IsNewUser"`
		// Uid订阅的告警策略Ids

		SubGroupIds []*int64 `json:"SubGroupIds,omitempty" name:"SubGroupIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNewUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略组id

		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolicyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindingPolicyTagRequest struct {
	*tchttp.BaseRequest

	// 固定取值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 策略标签

	Tag *PolicyTag `json:"Tag,omitempty" name:"Tag"`
	// 产品类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *UnbindingPolicyTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindingPolicyTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllNamespacesRequest struct {
	*tchttp.BaseRequest

	// 根据监控类型过滤 "ST_DASHBOARD"=云产品监控 "MT_CUSTOM"=自定义监控

	SceneType *string `json:"SceneType,omitempty" name:"SceneType"`
	// 根据namespace的Id过滤 不填默认查询所有

	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`
	// 根据使用场景过滤 "ST_DASHBOARD"=Dashboard类型 或 "ST_ALARM"=告警类型

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAllNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 主题ID列表，逗号分割

	TopicIds *string `json:"TopicIds,omitempty" name:"TopicIds"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询

	Query *string `json:"Query,omitempty" name:"Query"`
	// 条数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 加载更多使用，透传上次返回的 context 值

	Context *string `json:"Context,omitempty" name:"Context"`
	// 排序

	Sort *string `json:"Sort,omitempty" name:"Sort"`
}

func (r *SearchLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribeInfoReceiver struct {

	// 用户名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 用户Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 用户Uid

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

type BindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 策略分组Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 实例分组ID

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 需要绑定的对象维度信息

	Dimensions []*BindingPolicyObjectDimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 必填。固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略ID，使用此字段时GroupId可传入任意值

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *BindingPolicyObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindingPolicyObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 实例组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 实例列表

	InstanceList []*BindingPolicyObjectDimension `json:"InstanceList,omitempty" name:"InstanceList"`
}

func (r *CreateInstanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMsgPolicyVoiceConfig struct {

	// uid列表

	UidList []*int64 `json:"UidList,omitempty" name:"UidList"`
	// 告警轮数, 1到5之间

	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`
	// 每轮告警之间的时间间隔, 单位秒

	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`
	// 每个人之间的时间间隔, 单位秒

	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`
	// 是否需要通知, 1表示需要, 0表示不需要

	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`
}

type URLNotice struct {

	// 回调 url（限长256字符）

	URL *string `json:"URL,omitempty" name:"URL"`
	// 是否通过验证 0=否 1=是

	IsValid *int64 `json:"IsValid,omitempty" name:"IsValid"`
	// 验证码

	ValidationCode *string `json:"ValidationCode,omitempty" name:"ValidationCode"`
	// 通知开始时间 00:00:00 开始的秒数（取值范围0-86399）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 通知结束时间 00:00:00 开始的秒数（取值范围0-86399）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 星期描述，0代表周日，1~6分别代表周一至周六

	Weekday []*int64 `json:"Weekday,omitempty" name:"Weekday"`
}

type DescribePolicyQuotaRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribePolicyQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventPolicyConfig struct {

	// 事件ID

	EventID *string `json:"EventID,omitempty" name:"EventID"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品英文名

	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`
	// 产品中文名

	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`
	// 事件名

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 事件英文名

	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`
	// 事件中文名

	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`
	// 维度

	Dimensions []*DescribeEventPolicyConfigDimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

type DescribePolicySituationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 未订阅告警策略组数量

		UnSubscribePolicyGroupCount *int64 `json:"UnSubscribePolicyGroupCount,omitempty" name:"UnSubscribePolicyGroupCount"`
		// 告警状态信息

		AlarmSituation *DescribePolicySituationAlarmSituation `json:"AlarmSituation,omitempty" name:"AlarmSituation"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicySituationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicySituationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyQuotaV3ResponseDataQuotaList struct {

	// 用户

	Used *int64 `json:"Used,omitempty" name:"Used"`
	// 总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 项目ID

	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
}

type DescribePolicyInfoByInstanceConditionsTemp struct {

	// 组ID

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 最后编辑者UIN

	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 是否为与关系

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ClmMetricAnalysisCompareData struct {

	// 对比数据

	Compare []*PCLMMetricAnalysisPoint `json:"Compare,omitempty" name:"Compare"`
}

type DeleteMetricSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除成功失败布尔值

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMetricSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetricSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMChartsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *CCMChartData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMChartsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMChartsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListEventMetric struct {

	// 事件id

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 事件名称

	EventShowName *string `json:"EventShowName,omitempty" name:"EventShowName"`
	// 是否需要恢复

	NeedRecovered *bool `json:"NeedRecovered,omitempty" name:"NeedRecovered"`
	// 事件类型，预留字段，当前固定取值为2

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type MetricObjectMeaning struct {

	// 指标英文解释

	En *string `json:"En,omitempty" name:"En"`
	// 指标中文解释

	Zh *string `json:"Zh,omitempty" name:"Zh"`
}

type UserNotice struct {

	// 接收者类型 USER=用户 GROUP=用户组

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 用户 uid 列表

	UserIds []*int64 `json:"UserIds,omitempty" name:"UserIds"`
	// 用户组 group id 列表

	GroupIds []*int64 `json:"GroupIds,omitempty" name:"GroupIds"`
	// 通知开始时间 00:00:00 开始的秒数（取值范围0-86399）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 通知结束时间 00:00:00 开始的秒数（取值范围0-86399）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 通知渠道列表 EMAIL=邮件 SMS=短信 CALL=电话 WECHAT=微信

	NoticeWay []*string `json:"NoticeWay,omitempty" name:"NoticeWay"`
	// 电话轮询列表

	PhoneOrder []*int64 `json:"PhoneOrder,omitempty" name:"PhoneOrder"`
	// 电话轮询次数 （取值范围1-5）

	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitempty" name:"PhoneCircleTimes"`
	// 单次轮询内拨打间隔 秒数 （取值范围60-900）

	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitempty" name:"PhoneInnerInterval"`
	// 两次轮询间隔 秒数（取值范围60-900）

	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitempty" name:"PhoneCircleInterval"`
	// 是否需要触达通知 0=否 1=是

	NeedPhoneArriveNotice *int64 `json:"NeedPhoneArriveNotice,omitempty" name:"NeedPhoneArriveNotice"`
	// 电话拨打类型 SYNC=同时拨打 CIRCLE=轮询拨打 不指定时默认是轮询

	PhoneCallType *string `json:"PhoneCallType,omitempty" name:"PhoneCallType"`
	// 通知周期 1-7表示周一到周日

	Weekday []*int64 `json:"Weekday,omitempty" name:"Weekday"`
}

type DeleteInstancesRequest struct {
	*tchttp.BaseRequest

	// 分组Id

	Instances []*CgrpInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *DeleteInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDimensionAnalysisDataRequest struct {
	*tchttp.BaseRequest

	// 指标集ID

	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`
	// 开始时间戳

	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`
	// 结束时间戳

	EndTimestamp *int64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`
	// 分析指标字段

	AnalyticMetric *string `json:"AnalyticMetric,omitempty" name:"AnalyticMetric"`
	// 分析的维度字段

	AnalyticDimensionKey *string `json:"AnalyticDimensionKey,omitempty" name:"AnalyticDimensionKey"`
	// 分析的维度值列表

	AnalyticDimensionValues []*string `json:"AnalyticDimensionValues,omitempty" name:"AnalyticDimensionValues"`
	// 维度筛选条件

	Filters []*ClmAnalysisFilter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeDimensionAnalysisDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDimensionAnalysisDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomMetisAbnormOutputAttr struct {

	// 自定义监控属性ID

	AttrId *uint64 `json:"AttrId,omitempty" name:"AttrId"`
	// 异常点

	Abnormals []*uint64 `json:"Abnormals,omitempty" name:"Abnormals"`
}

type DescribeAlarmHistoryByAlarmIdRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警Id

	AlarmId *int64 `json:"AlarmId,omitempty" name:"AlarmId"`
	// 起始时间，比如2020-02-07 15:47:24

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，比如2020-02-26 16:47:18

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAlarmHistoryByAlarmIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmHistoryByAlarmIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type WebStorageInfos struct {

	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 子用户

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 主用户

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 弹窗信息

	WebStorageInfo *string `json:"WebStorageInfo,omitempty" name:"WebStorageInfo"`
}

type ClonePolicyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 复制成功的策略组Id

		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ClonePolicyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClonePolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 视图ID

		ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashboardViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回指标上报数据

		Data []*DashboardMetricData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTypesRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeUserTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUnifyDashboardsRequest struct {
	*tchttp.BaseRequest

	// dashboard uuid 列表

	UUIDList []*string `json:"UUIDList,omitempty" name:"UUIDList"`
}

func (r *DeleteUnifyDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnifyDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsForConsoleFontEndResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标数据

		Data []*DescribeBaseMetricsForConsoleFontEndData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBaseMetricsForConsoleFontEndResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaseMetricsForConsoleFontEndResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyConditionRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略 ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 裸写包含触发规则、过滤条件等内容的告警规则

	PolicyExpression *string `json:"PolicyExpression,omitempty" name:"PolicyExpression"`
	// 触发条件模板 Id，待改造为字符串

	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`
	// 指标触发条件

	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`
	// 事件触发条件

	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`
}

func (r *ModifyAlarmPolicyConditionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyConditionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CLMDescribeMetricSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标集列表结果数据

		Data *PCLMDescribeMetricSetsData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CLMDescribeMetricSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CLMDescribeMetricSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeCallbacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警回调通知

		URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticeCallbacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticeCallbacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyUseListRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 地域id

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域简称

	Area *string `json:"Area,omitempty" name:"Area"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 搜索关键词

	Search *string `json:"Search,omitempty" name:"Search"`
	// 维度组合json字符串

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *DescribePolicyUseListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyUseListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecommendedTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板策略组总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 模板策略组

		GroupTemplateLists []*GroupTemplateList `json:"GroupTemplateLists,omitempty" name:"GroupTemplateLists"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecommendedTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecommendedTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMGroupViewRequest struct {
	*tchttp.BaseRequest

	// 分组ID

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 分组视图是否自动绑定指标。0：否，1：是

	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`
}

func (r *ModifyCCMGroupViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCCMGroupViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusMapOut struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 值列表

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type InstanceGroups struct {

	// 实例组 Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeBaseMetricsForConsoleFontEndRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 名字空间

	NameSpace *string `json:"NameSpace,omitempty" name:"NameSpace"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标名列表

	Batch *string `json:"Batch,omitempty" name:"Batch"`
}

func (r *DescribeBaseMetricsForConsoleFontEndRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBaseMetricsForConsoleFontEndRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分组树

		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
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

type ModifyInstanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategyEntryAlarm struct {

	// 策略的告警规则类型。100：最大值（即>）， 101：最大值（即>=），102：最小值告警（即 <）， 103：最小值告警（即 <=），104：波动值告警

	AlarmType *uint64 `json:"AlarmType,omitempty" name:"AlarmType"`
	// 策略的告警规则属性ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 策略的告警规则阈值

	Threshold *int64 `json:"Threshold,omitempty" name:"Threshold"`
	// 策略的告警规则DetectPeriod分钟内告警次数

	DetectCount *int64 `json:"DetectCount,omitempty" name:"DetectCount"`
	// 策略的告警周期

	DetectPeriod *int64 `json:"DetectPeriod,omitempty" name:"DetectPeriod"`
	// 策略的告警规则告警间隔周期

	TriggerPeriod *int64 `json:"TriggerPeriod,omitempty" name:"TriggerPeriod"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 指标中文描述

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 指标英文描述，唯一名称

	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`
}

type DescribeAlertPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// data

		Data *ClmDescribeAlertPoliciesData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMsgPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 接收用户组id列表

	ReceiverGroupIds []*int64 `json:"ReceiverGroupIds,omitempty" name:"ReceiverGroupIds"`
	// 告警方式列表

	NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays"`
	// 电话告警配置

	VoiceConfig *ModifyMsgPolicyVoiceConfig `json:"VoiceConfig,omitempty" name:"VoiceConfig"`
}

func (r *CreateMsgPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMsgPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略

		Data *ClmAlertPolicy `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewRequest struct {
	*tchttp.BaseRequest

	// 分组ID

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeCCMGroupViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMGroupViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyNoticeRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略 ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 告警通知模板 ID 列表

	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
}

func (r *ModifyAlarmPolicyNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询得到的指标描述列表

		MetricSet []*MetricSet `json:"MetricSet,omitempty" name:"MetricSet"`
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

type InitGroupRequest struct {
	*tchttp.BaseRequest

	// 分组树高度

	High *int64 `json:"High,omitempty" name:"High"`
	// 分组名称,数组大小和高度一致

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 叶子节点存放实例的类型，填SCF或cvm_device

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *InitGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentEventListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 平台事件列表

		Alarms []*DescribeAccidentEventListAlarms `json:"Alarms,omitempty" name:"Alarms"`
		// 平台事件的总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccidentEventListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccidentEventListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *ServerInfoData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentConfigAccidentTypeInfo struct {

	// 平台事件id

	AccidentId *int64 `json:"AccidentId,omitempty" name:"AccidentId"`
	// 平台事件名称

	AccidentName *string `json:"AccidentName,omitempty" name:"AccidentName"`
	// 平台事件英文名称

	AccidentEnName *string `json:"AccidentEnName,omitempty" name:"AccidentEnName"`
}

type GroupTemplateList struct {

	// 模板策略組ID

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 模板策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否为默认策略

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 告警策略启用状态

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 最后修改人UIN

	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 未屏蔽的实例数

	NoShieldedInstanceCount *int64 `json:"NoShieldedInstanceCount,omitempty" name:"NoShieldedInstanceCount"`
	// 总绑定实例数

	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" name:"TotalInstanceCount"`
	// 父策略组ID

	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 视图

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 指标告警触发条件

	Conditions []*Condition `json:"Conditions,omitempty" name:"Conditions"`
	// 事件告警触发条件

	EventConditions []*EventCondition `json:"EventConditions,omitempty" name:"EventConditions"`
	// 是否可设为默认告警策略

	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 修改时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeAlarmPoliciesBasicRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 要查询的策略的【监控类型+1.0 Id】列表

	OriginIds []*AlarmPolicyBasic `json:"OriginIds,omitempty" name:"OriginIds"`
}

func (r *DescribeAlarmPoliciesBasicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPoliciesBasicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeProductEventListOverView struct {

	// 状态变更的事件数量

	StatusChangeAmount *int64 `json:"StatusChangeAmount,omitempty" name:"StatusChangeAmount"`
	// 告警状态未配置的事件数量

	UnConfigAlarmAmount *int64 `json:"UnConfigAlarmAmount,omitempty" name:"UnConfigAlarmAmount"`
	// 异常事件数量

	UnNormalEventAmount *int64 `json:"UnNormalEventAmount,omitempty" name:"UnNormalEventAmount"`
	// 未恢复的事件数量

	UnRecoverAmount *int64 `json:"UnRecoverAmount,omitempty" name:"UnRecoverAmount"`
}

type ProductHealthStatus struct {

	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 健康状态，1表示未恢复

	HealthStatus *int64 `json:"HealthStatus,omitempty" name:"HealthStatus"`
	// 提醒的数量

	WarningCount *int64 `json:"WarningCount,omitempty" name:"WarningCount"`
	// 异常的数量

	AbnormalCount *int64 `json:"AbnormalCount,omitempty" name:"AbnormalCount"`
	// 告警的地域列表

	AlarmRegions []*string `json:"AlarmRegions,omitempty" name:"AlarmRegions"`
}

type TransLogInstanceGroupInfo struct {

	// 实例组Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 实例组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 实例数量

	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 类型

	CategoryType *int64 `json:"CategoryType,omitempty" name:"CategoryType"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近修改时间

	LastModifyTime *int64 `json:"LastModifyTime,omitempty" name:"LastModifyTime"`
	// 最后修改UIN

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
}

type DescribeAccidentEventListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前接口取值monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 起始时间，默认一天前的时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认当前时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 根据UpdateTime排序的规则，取值asc或desc

	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`
	// 根据OccurTime排序的规则，取值asc或desc（优先根据UpdateTimeOrder排序）

	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`
	// 根据事件类型过滤，1表示服务问题，2表示其他订阅

	AccidentType []*int64 `json:"AccidentType,omitempty" name:"AccidentType"`
	// 根据事件过滤，1表示云服务器存储问题，2表示云服务器网络连接问题，3表示云服务器运行异常，202表示运营商网络抖动

	AccidentEvent []*int64 `json:"AccidentEvent,omitempty" name:"AccidentEvent"`
	// 根据事件状态过滤，0表示已恢复，1表示未恢复

	AccidentStatus []*int64 `json:"AccidentStatus,omitempty" name:"AccidentStatus"`
	// 根据事件地域过滤，gz表示广州，sh表示上海等

	AccidentRegion []*string `json:"AccidentRegion,omitempty" name:"AccidentRegion"`
	// 根据影响资源过滤，比如ins-19a06bka

	AffectResource *string `json:"AffectResource,omitempty" name:"AffectResource"`
}

func (r *DescribeAccidentEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccidentEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertPolicyRequest struct {
	*tchttp.BaseRequest

	// 告警策略id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeAlertPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentStatusHistoryData struct {

	// 实例ID

	UnInstanceID *string `json:"UnInstanceID,omitempty" name:"UnInstanceID"`
	// 子机状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 历史数据

	Histories []*DescribeAgentStatusHistory `json:"Histories,omitempty" name:"Histories"`
}

type UnifyDashboardMeta struct {

	// dashboard uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 是否收藏

	IsStarred *bool `json:"IsStarred,omitempty" name:"IsStarred"`
	// dashboard名

	Title *string `json:"Title,omitempty" name:"Title"`
	// 是否默认

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// dashboard类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 最近访问时间

	LastVisitTime *string `json:"LastVisitTime,omitempty" name:"LastVisitTime"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 文件夹uuid

	FolderUUID *string `json:"FolderUUID,omitempty" name:"FolderUUID"`
	// 是否文件夹

	IsFolder *bool `json:"IsFolder,omitempty" name:"IsFolder"`
	// 顺序

	Order *string `json:"Order,omitempty" name:"Order"`
}

type CreatePolicyGroupEventCondition struct {

	// 如果通过模版创建，需要传入模版中该指标的对应RuleId

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 告警事件的Id

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 告警发送收敛类型。0连续告警，1指数告警

	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次

	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
}

type ArgusDimensionIn struct {

	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
}

type DescribeAlarmPolicyQuotaRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控类型过滤 "MT_QCE"=云产品监控 "MT_CUSTOM"=自定义监控

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 触发条件类型 STATIC=静态阈值 DYNAMIC=动态阈值

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
}

func (r *DescribeAlarmPolicyQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPolicyQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListInstance struct {

	// 对象唯一id

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
	// 表示对象实例的维度集合，jsonObj字符串

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 对象是否被屏蔽，0表示未屏蔽，1表示被屏蔽

	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`
	// 对象所在的地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type ServiceInfo struct {

	// 服务ID

	ServiceId *int64 `json:"ServiceId,omitempty" name:"ServiceId"`
	// 是否开通服务

	IsOpen *int64 `json:"IsOpen,omitempty" name:"IsOpen"`
	// 默认视图ID

	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`
}

type DescribeAbnormalObjectsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标告警异常对象列表

		ThresholdObjects []*DescribeAbnormalObjectsThresholdObject `json:"ThresholdObjects,omitempty" name:"ThresholdObjects"`
		// 产品事件告警异常对象列表

		EventObjects []*DescribeAbnormalObjectsEventObject `json:"EventObjects,omitempty" name:"EventObjects"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAbnormalObjectsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAbnormalObjectsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略详情

		Policy *AlarmPolicy `json:"Policy,omitempty" name:"Policy"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeUnitInfoOutput struct {

	// 指标单位ID

	UnitId *uint64 `json:"UnitId,omitempty" name:"UnitId"`
	// 指标单位名称

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricConfig struct {

	// 允许使用的运算符

	Operator []*string `json:"Operator,omitempty" name:"Operator"`
	// 允许配置的数据周期，以秒为单位

	Period []*int64 `json:"Period,omitempty" name:"Period"`
	// 允许配置的持续周期个数

	ContinuePeriod []*int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`
}

type ModifyUnifyDashboardRequest struct {
	*tchttp.BaseRequest

	// dashboard 详细信息

	Data *string `json:"Data,omitempty" name:"Data"`
	// 是否强制覆盖，"true"或者"false"

	Overwrite *string `json:"Overwrite,omitempty" name:"Overwrite"`
	// "1"表示压缩

	Compress *string `json:"Compress,omitempty" name:"Compress"`
	// 文件夹UUID

	FolderUUID *string `json:"FolderUUID,omitempty" name:"FolderUUID"`
}

func (r *ModifyUnifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUnifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShieldPolicyAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ShieldPolicyAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ShieldPolicyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusNamespaceOut struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 数据最大保存时间，单位天

	DataTimeLimit *uint64 `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`
	// 数据最大存储容量，单位GB

	DataDiskLimit *uint64 `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`
	// 管理员列表

	AdminUins []*string `json:"AdminUins,omitempty" name:"AdminUins"`
	// 聚合周期

	AggPeriod *uint64 `json:"AggPeriod,omitempty" name:"AggPeriod"`
	// AppId

	AppId *uint64 `json:"AppId,omitempty" name:"AppId"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

type CreateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		Data []*CgrpInstanceNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCCMChartsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCCMChartsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCCMChartsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMsgPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMsgPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMsgPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmDimension struct {

	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 条目ID（只作出参，入参不填）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 指标集ID（只作出参，入参不填）

	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
}

type DescribeDashboardsRequest struct {
	*tchttp.BaseRequest

	// 产品类型

	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
	// CustomID

	CustomID *string `json:"CustomID,omitempty" name:"CustomID"`
	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribeEventRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 业务id

	BusinessId *int64 `json:"BusinessId,omitempty" name:"BusinessId"`
	// 事件id

	AccidentId *int64 `json:"AccidentId,omitempty" name:"AccidentId"`
	// 通知方式

	NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays"`
	// 接收人(json字符串)

	Receivers *string `json:"Receivers,omitempty" name:"Receivers"`
	// 用户配置(json字符串)

	UserConfig *string `json:"UserConfig,omitempty" name:"UserConfig"`
}

func (r *SubscribeEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubscribeEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmAlertFilterRule struct {

	// 关联的告警策略ID

	AlertPolicyId *int64 `json:"AlertPolicyId,omitempty" name:"AlertPolicyId"`
	// 过滤的关键字

	Key *string `json:"Key,omitempty" name:"Key"`
	// 操作符

	Operating *string `json:"Operating,omitempty" name:"Operating"`
	// 筛选值（多个值逗号分隔）

	Value *string `json:"Value,omitempty" name:"Value"`
	// 过滤条件id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type SetDefaultUnifyDashboardRequest struct {
	*tchttp.BaseRequest

	// dashboard uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

func (r *SetDefaultUnifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultUnifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *ServiceInfo `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *AttributeIdOutput `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMMetricSet struct {

	// CLS日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// CLS日志集名

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// CLS日志主题ID

	LogtopicId *string `json:"LogtopicId,omitempty" name:"LogtopicId"`
	// CLS日志主题名

	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`
	// 指标集名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 统计周期 可选：60、300

	StatisticCycle *int64 `json:"StatisticCycle,omitempty" name:"StatisticCycle"`
	// 时间字段

	TimeField *string `json:"TimeField,omitempty" name:"TimeField"`
	// 日志过滤规则

	LogFilterRules []*PCLMLogFilterRule `json:"LogFilterRules,omitempty" name:"LogFilterRules"`
	// 基础指标规则

	MetricBasicItems []*PCLMMetricBasicItem `json:"MetricBasicItems,omitempty" name:"MetricBasicItems"`
	// 复合指标规则

	MetricCustomItems []*PCLMMetricCustomItem `json:"MetricCustomItems,omitempty" name:"MetricCustomItems"`
	// 维度规则

	Dimensions []*PCLMDimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// CID（只作出参，入参不填）

	CID *string `json:"CID,omitempty" name:"CID"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// APPID（只作出参，入参不填）

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// UIN（只作出参，入参不填）

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 地域（只作出参，入参不填）

	Region *string `json:"Region,omitempty" name:"Region"`
	// 日志数据结构描述

	LogProfileItems []*PCLMLogProfileItem `json:"LogProfileItems,omitempty" name:"LogProfileItems"`
	// SubAccountUin(只作出参，入参不填）

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
}

type CopyUnifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 新dashboard uuid

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyUnifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyUnifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略 ID

		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
		// 用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID

		OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUnifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dashboard uuid

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateUnifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUnifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertPolicyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功为true

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlertPolicyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlertPolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardFoldersRequest struct {
	*tchttp.BaseRequest

	// 文件夹uuid列表

	UUIDList []*string `json:"UUIDList,omitempty" name:"UUIDList"`
	// 是否递归

	Recursive *string `json:"Recursive,omitempty" name:"Recursive"`
}

func (r *DeleteDashboardFoldersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardFoldersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataForMiniProgramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回告警历史数据

		AlarmInfo *string `json:"AlarmInfo,omitempty" name:"AlarmInfo"`
		// 返回告警指标数据

		MetricDatas *MetricData `json:"MetricDatas,omitempty" name:"MetricDatas"`
		// 账号ID

		AppId []*int64 `json:"AppId,omitempty" name:"AppId"`
		// 返回提示信息

		RetMsg *string `json:"RetMsg,omitempty" name:"RetMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataForMiniProgramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataForMiniProgramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSortObjectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 结果对象列表，json字符串

		ObjList *string `json:"ObjList,omitempty" name:"ObjList"`
		// 耗时(秒)

		Cost *float64 `json:"Cost,omitempty" name:"Cost"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSortObjectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSortObjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusOverviewNamespaceOut struct {

	// Namespace的名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// Namespace的描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 数据最大保存时间，单位天

	DataTimeLimit *string `json:"DataTimeLimit,omitempty" name:"DataTimeLimit"`
	// 数据最大存储容量，单位GB

	DataDiskLimit *string `json:"DataDiskLimit,omitempty" name:"DataDiskLimit"`
	// 已使用存储容量，单位GB

	DiskUsage *string `json:"DiskUsage,omitempty" name:"DiskUsage"`
}

type DeleteConditionsTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConditionsTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConditionsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardDimensionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度列表

		DimensionList []*string `json:"DimensionList,omitempty" name:"DimensionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardDimensionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardDimensionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardViewsRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控面板ID

	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
}

func (r *DescribeDashboardViewsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardViewsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例组基础信息

		InstanceGroupInfo *InstanceGroupInfo `json:"InstanceGroupInfo,omitempty" name:"InstanceGroupInfo"`
		// 实例列表

		InstanceList []*TransLogInstance `json:"InstanceList,omitempty" name:"InstanceList"`
		// 实例组关联的策略

		PolicyGroups []*PolicyGroupItem `json:"PolicyGroups,omitempty" name:"PolicyGroups"`
		// 实例所属的地域

		Regions []*string `json:"Regions,omitempty" name:"Regions"`
		// 实例总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardPanel struct {

	// copy|add

	Type *string `json:"Type,omitempty" name:"Type"`
	// dashboard的UUID

	DashboardUUID *string `json:"DashboardUUID,omitempty" name:"DashboardUUID"`
	// panel的ID

	PanelID *int64 `json:"PanelID,omitempty" name:"PanelID"`
	// panel数据

	Panel *string `json:"Panel,omitempty" name:"Panel"`
}

type ModifyConditionsTemplateRequest struct {
	*tchttp.BaseRequest

	// 模板策略ID

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 视图

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 模板名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否为与关系

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
	// 指标告警条件

	Conditions []*ModifyConditionsTemplateRequestCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 事件告警条件

	EventConditions []*ModifyConditionsTemplateRequestEventCondition `json:"EventConditions,omitempty" name:"EventConditions"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 父策略组ID

	ParentGroupID *string `json:"ParentGroupID,omitempty" name:"ParentGroupID"`
	// 是否生效

	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`
	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *ModifyConditionsTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConditionsTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindingAllPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *UnBindingAllPolicyObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindingAllPolicyObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 产品类型过滤，比如"cvm"表示云服务器

	ProductName []*string `json:"ProductName,omitempty" name:"ProductName"`
	// 事件名称过滤，比如"guest_reboot"表示机器重启

	EventName []*string `json:"EventName,omitempty" name:"EventName"`
	// 影响对象，比如ins-19708ino

	InstanceId []*string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 维度过滤，比如外网IP:10.0.0.1

	Dimensions []*DescribeProductEventListDimensions `json:"Dimensions,omitempty" name:"Dimensions"`
	// 地域过滤，比如gz

	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`
	// 事件类型过滤，取值范围["status_change","abnormal"]，分别表示状态变更、异常事件

	Type []*string `json:"Type,omitempty" name:"Type"`
	// 事件状态过滤，取值范围["recover","alarm","-"]，分别表示已恢复、未恢复、无状态

	Status []*string `json:"Status,omitempty" name:"Status"`
	// 项目ID过滤

	Project []*string `json:"Project,omitempty" name:"Project"`
	// 告警状态配置过滤，1表示已配置，0表示未配置

	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`
	// 按更新时间排序，ASC表示升序，DESC表示降序，默认DESC

	TimeOrder *string `json:"TimeOrder,omitempty" name:"TimeOrder"`
	// 起始时间，默认一天前的时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认当前时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 页偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页返回的数量，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeProductEventListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductEventListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerInfo struct {

	// 服务器ID

	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`
	// 服务器名称

	ServerName *string `json:"ServerName,omitempty" name:"ServerName"`
	// 服务器内网IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 地域ID

	IdcId *int64 `json:"IdcId,omitempty" name:"IdcId"`
	// 地域名称

	IdcName *string `json:"IdcName,omitempty" name:"IdcName"`
}

type AttributeUnitInfoOutputData struct {

	// 属性单位数据列表

	Data []*AttributeUnitInfoOutput `json:"Data,omitempty" name:"Data"`
	// 属性单位数据列表个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DescribeSortObjectListDimension struct {

	// 维度名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeGroupIDByTagsRequest struct {
	*tchttp.BaseRequest

	// 产品类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 标签

	Tags []*PolicyTag `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeGroupIDByTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupIDByTagsRequest) FromJsonString(s string) error {
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

type DeleteAlarmNoticesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmNoticesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *StrategyData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMGroupViewStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMGroupViewStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GroupIdTag struct {

	// 1.0策略ID

	GroupIDList []*string `json:"GroupIDList,omitempty" name:"GroupIDList"`
	// 标签Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeDashboardStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否有新版本

		HasNewVersion *bool `json:"HasNewVersion,omitempty" name:"HasNewVersion"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindingInstanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingInstanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindingInstanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CvmAgentStatus struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Agent是否安装

	AgentInstalled *bool `json:"AgentInstalled,omitempty" name:"AgentInstalled"`
}

type DescribeStorageDurationRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeStorageDurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageDurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCMDashboardEntry struct {

	// 自定义监控图表ID

	DashboardId *int64 `json:"DashboardId,omitempty" name:"DashboardId"`
	// 自定义监控图表名称

	DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
	// 自定义监控图表类型。1 明细视图: , 2: 聚合视图

	DashboardType *uint64 `json:"DashboardType,omitempty" name:"DashboardType"`
	// 产品类型。1 基础监控: , 2: 自定义监控

	ProductType *uint64 `json:"ProductType,omitempty" name:"ProductType"`
	// 指标ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 聚合方式列表。列表元素：1 SUM， 2 AVG， 3 MAX， 4 MIN

	Aggregation []*int64 `json:"Aggregation,omitempty" name:"Aggregation"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建者ID

	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新者ID

	UpdaterId *uint64 `json:"UpdaterId,omitempty" name:"UpdaterId"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 实例ID列表。

	InstanceId []*int64 `json:"InstanceId,omitempty" name:"InstanceId"`
}

type DescribeAlarmCallbackHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回调历史列表

		List []*DescribeAlarmCallbackHistory `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmCallbackHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmCallbackHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyUseListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警对象数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 告警对象列表

		List []*DescribePolicyUseList `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyUseListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyUseListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmLogProfileItem struct {

	// id（只做出参，入参不填）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 关联指标集ID（只做出参，入参不填）

	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 字段类型，可选：string、long、double

	Type *string `json:"Type,omitempty" name:"Type"`
	// 创建时间（只做出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只做出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type MoveFailReason struct {

	// dashboard的uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 移动失败原因

	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

type DescribeInstanceGroupListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 按分组名称进行搜索

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 实例分组id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 告警策略类型名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 更新时间的排序, asc 或者 desc

	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`
}

func (r *DescribeInstanceGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusFrontTunnelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// HTTP请求返回结果

		ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ArgusFrontTunnelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ArgusFrontTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardDataSourcesRequest struct {
	*tchttp.BaseRequest

	// 所属模块，默认为 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 应用空间SpaceUUID，默认值为space_default

	SpaceUUID *string `json:"SpaceUUID,omitempty" name:"SpaceUUID"`
}

func (r *DescribeDashboardDataSourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardDataSourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMGroupViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCMGroupViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCCMGroupViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 视图ID

		ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleRequest struct {
	*tchttp.BaseRequest

	// 分组Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteModuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGraphDataRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 查询起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例ID

	UnInstanceID *string `json:"UnInstanceID,omitempty" name:"UnInstanceID"`
	// 实例UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

func (r *DescribeGraphDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGraphDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePreferenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// light|dark

		Theme *string `json:"Theme,omitempty" name:"Theme"`
		// 默认dashboard UUID

		HomeDashboardUUID *string `json:"HomeDashboardUUID,omitempty" name:"HomeDashboardUUID"`
		// 时区

		Timezone *string `json:"Timezone,omitempty" name:"Timezone"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePreferenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePreferenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmReceiversResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmReceiversResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmReceiversResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesInInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 实例组Id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 实例的uniqueId列表

	UniqueIds []*string `json:"UniqueIds,omitempty" name:"UniqueIds"`
}

func (r *DeleteInstancesInInstanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstancesInInstanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultPolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *SetDefaultPolicyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultPolicyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGraphDataPartition struct {

	// 分区名

	Name []*string `json:"Name,omitempty" name:"Name"`
	// 用量

	Usage *float64 `json:"Usage,omitempty" name:"Usage"`
	// 总量

	Total *float64 `json:"Total,omitempty" name:"Total"`
}

type CreateAttributesRequest struct {
	*tchttp.BaseRequest

	// 属性信息

	Attribute []*AttributeInfoInput `json:"Attribute,omitempty" name:"Attribute"`
}

func (r *CreateAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewStrategyRequest struct {
	*tchttp.BaseRequest

	// 实例分组 ID。GroupId或ViewId必选其一，若ViewId存在，则以ViewId为准

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 分组视图ID

	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`
	// 策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 页码

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeCCMGroupViewStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMGroupViewStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductListRequest struct {
	*tchttp.BaseRequest

	// 排序方式：DESC/ASC（区分大小写），默认值DESC

	Order *string `json:"Order,omitempty" name:"Order"`
	// 分页查询的偏移量，默认值0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页查询的每页数据量，默认值20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 固定传值monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeProductListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerInfoData struct {

	// 返回数据

	Data []*ServerInfo `json:"Data,omitempty" name:"Data"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type TransLogItem struct {

	// 变更日志Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 模块Id

	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
	// 日志信息，json字符串

	LogData *string `json:"LogData,omitempty" name:"LogData"`
	// 创建时间

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最后修改UIN

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
}

type EventCondition struct {

	// 告警通知频率

	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
	// 重复通知策略预定义（0 - 只告警一次， 1 - 指数告警，2 - 连接告警）

	AlarmNotifyType *string `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 事件ID

	EventID *string `json:"EventID,omitempty" name:"EventID"`
	// 事件展示名称（对外）

	EventDisplayName *string `json:"EventDisplayName,omitempty" name:"EventDisplayName"`
	// 规则ID

	RuleID *string `json:"RuleID,omitempty" name:"RuleID"`
}

type DescribeCCMChartsRequest struct {
	*tchttp.BaseRequest

	// 页码。默认为1

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量。默认为全量数据总数。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 监控图表ID列表

	ChartId []*int64 `json:"ChartId,omitempty" name:"ChartId"`
	// 监控图表名称

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
	// 监控图表类型。1 明细视图， 2 聚合视图

	ChartType *int64 `json:"ChartType,omitempty" name:"ChartType"`
	// 产品类型。1 基础监控， 2 自定义监控

	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`
	// 指标ID

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 分组ID

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 排序方式。仅支持："asc","ASC","desc","DESC"

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段。仅支持："CreateTime","UpdateTime", "ChartName"

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeCCMChartsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMChartsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户appid

		Appid *int64 `json:"Appid,omitempty" name:"Appid"`
		// 用户所属的大客户类型

		UserTypes []*string `json:"UserTypes,omitempty" name:"UserTypes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警通知规则名称 60字符以内

	Name *string `json:"Name,omitempty" name:"Name"`
	// 通知类型 ALARM=未恢复通知 OK=已恢复通知 ALL=都通知

	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`
	// 通知语言 zh-CN=中文 en-US=英文

	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`
	// 用户通知 最多5个

	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices"`
	// 回调通知 最多3个

	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`
	// 告警通知模板 ID

	NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`
	// 其他非公开通知渠道列表

	ExtraChannels []*string `json:"ExtraChannels,omitempty" name:"ExtraChannels"`
	// CLS通知 最多1个

	CLSNotices []*CLSNotice `json:"CLSNotices,omitempty" name:"CLSNotices"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMMetricBasicItem struct {

	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 指标统计方法：sum count min max

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标描述中文名

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 指标过滤规则

	StatisticFilterRules []*PCLMStatisticFilterRule `json:"StatisticFilterRules,omitempty" name:"StatisticFilterRules"`
	// 指标集CID（只作出参，入参不填）

	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type AttributeIdDeleteOutput struct {

	// 删除个数

	DeleteCount *int64 `json:"DeleteCount,omitempty" name:"DeleteCount"`
	// 属性ID

	AttributeId []*uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

type Metric struct {

	// 告警策略类型

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 仅对云产品监控有效，用于某些接口查询数据

	QceNamespace *string `json:"QceNamespace,omitempty" name:"QceNamespace"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标展示名

	Description *string `json:"Description,omitempty" name:"Description"`
	// 聚合方式

	Aggregations []*MetricAggregation `json:"Aggregations,omitempty" name:"Aggregations"`
	// 最小值

	Min *float64 `json:"Min,omitempty" name:"Min"`
	// 最大值

	Max *float64 `json:"Max,omitempty" name:"Max"`
	// 维度列表

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 富文本提示语

	Tooltip *MetricTooltip `json:"Tooltip,omitempty" name:"Tooltip"`
	// 指标配置

	MetricConfig *MetricConfig `json:"MetricConfig,omitempty" name:"MetricConfig"`
	// Operators

	Operators []*string `json:"Operators,omitempty" name:"Operators"`
	// Periods

	Periods []*int64 `json:"Periods,omitempty" name:"Periods"`
}

type CreateDashboardViewRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控面板视图名

	DescName *string `json:"DescName,omitempty" name:"DescName"`
	// 名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 监控面板ID

	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
	// 面板实例列表

	Instances []*string `json:"Instances,omitempty" name:"Instances"`
	// 元数据，json串

	Meta *string `json:"Meta,omitempty" name:"Meta"`
	// 指标名列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
}

func (r *CreateDashboardViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmCallbackVerifyCodeRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmCallbackVerifyCodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmCallbackVerifyCodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTransLogRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 模块ID，1表示触发模板，2表示实例组，3表示告警策略

	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
	// 实例组ID

	DId *int64 `json:"DId,omitempty" name:"DId"`
	// 每页返回的数量，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 页偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeTransLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTransLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardDimensionValuesRequest struct {
	*tchttp.BaseRequest

	// 应用空间UUID

	SpaceUUID *string `json:"SpaceUUID,omitempty" name:"SpaceUUID"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 数据源

	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 维度名称

	DimensionName *string `json:"DimensionName,omitempty" name:"DimensionName"`
	// 取值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 搜索关键字

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
	// 已选择的维度，用于返回有指标数据的过滤值

	SelectedDimension *MonitorSelectedDimension `json:"SelectedDimension,omitempty" name:"SelectedDimension"`
	// 数据源地域信息

	DataSourceRegion *string `json:"DataSourceRegion,omitempty" name:"DataSourceRegion"`
}

func (r *DescribeDashboardDimensionValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardDimensionValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubscribeEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SubscribeEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SubscribeEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAttributeRequest struct {
	*tchttp.BaseRequest

	// 属性ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 属性名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 属性类型ID

	AttributeTypeId *uint64 `json:"AttributeTypeId,omitempty" name:"AttributeTypeId"`
	// 属性数据类型

	DataType *int64 `json:"DataType,omitempty" name:"DataType"`
	// 属性级别

	AttributeLevel *int64 `json:"AttributeLevel,omitempty" name:"AttributeLevel"`
	// 单位

	Unit *int64 `json:"Unit,omitempty" name:"Unit"`
	// 统计周期

	StatisticalPeriod *int64 `json:"StatisticalPeriod,omitempty" name:"StatisticalPeriod"`
	// 负责人ID

	OwnerId []*uint64 `json:"OwnerId,omitempty" name:"OwnerId"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeValueInfoOutputData struct {

	// 属性上报数据列表

	Data []*AttributeValueInfoOutput `json:"Data,omitempty" name:"Data"`
	// 属性数据列表个数个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DescribePolicyConditionListConfigManualCalcValue struct {

	// 默认值

	Default *string `json:"Default,omitempty" name:"Default"`
	// 固定值

	Fixed *string `json:"Fixed,omitempty" name:"Fixed"`
	// 最大值

	Max *string `json:"Max,omitempty" name:"Max"`
	// 最小值

	Min *string `json:"Min,omitempty" name:"Min"`
	// 是否必须

	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribeAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据列表

		Data *AttributeInfoOutputData `json:"Data,omitempty" name:"Data"`
		// 请求处理时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribePolicyGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMDescribeMetricSetsData struct {

	// 指标集列表数据

	MetricSets []*PCLMMetricSet `json:"MetricSets,omitempty" name:"MetricSets"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DeleteDashboardFoldersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功删除的文件夹列表

		SucList []*string `json:"SucList,omitempty" name:"SucList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardFoldersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPoliciesBasicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略基本信息列表

		PoliciesBasic []*AlarmPolicyBasic `json:"PoliciesBasic,omitempty" name:"PoliciesBasic"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPoliciesBasicResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPoliciesBasicResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMAnalysisFilter struct {

	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 运算符 目前只支持 =
	// （当value数组为多个值时为in关系）

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 值

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type SubscribeInfo struct {

	// 事件类型Id

	BusinessId *int64 `json:"BusinessId,omitempty" name:"BusinessId"`
	// 事件名

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// 问题类型Id

	AccidentId *int64 `json:"AccidentId,omitempty" name:"AccidentId"`
	// 问题名

	AccidentName *string `json:"AccidentName,omitempty" name:"AccidentName"`
	// 通知方式

	NotifyWay []*int64 `json:"NotifyWay,omitempty" name:"NotifyWay"`
	// 接收人列表

	Receivers []*SubscribeInfoReceiver `json:"Receivers,omitempty" name:"Receivers"`
	// 提示

	Tips *string `json:"Tips,omitempty" name:"Tips"`
	// 用户配置，json字符串

	UserConfig *string `json:"UserConfig,omitempty" name:"UserConfig"`
}

type DeletePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组id

	GroupId []*int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeletePolicyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMenuUinListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 信息

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMenuUinListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMenuUinListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusFrontTunnelRequest struct {
	*tchttp.BaseRequest

	// HTTP请求方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// HTTP请求地址

	Path *string `json:"Path,omitempty" name:"Path"`
	// HTTP请求体

	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
}

func (r *ArgusFrontTunnelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ArgusFrontTunnelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmAnalysisFilter struct {

	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 运算符 目前只支持 =
	// （当value数组为多个值时为in关系）

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 值

	Value []*string `json:"Value,omitempty" name:"Value"`
}

type DimensionsDesc struct {

	// 维度名数组

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type PCLMStatisticFilterRule struct {

	// 关联指标条目ID（只做出参，入参不填）

	MetricItemId *int64 `json:"MetricItemId,omitempty" name:"MetricItemId"`
	// 创建时间（只做出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只做出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 关系 1AND 2OR

	Relation *int64 `json:"Relation,omitempty" name:"Relation"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 过滤操作符: 数字型字段支持：gt lt ge le eq ne in 字符型字段支持：eq ne in contains

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 过滤操作值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeMenuWhitelistRequest struct {
	*tchttp.BaseRequest

	// 菜单名

	Menu *string `json:"Menu,omitempty" name:"Menu"`
}

func (r *DescribeMenuWhitelistRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMenuWhitelistRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSetRequest struct {
	*tchttp.BaseRequest

	// 指标集ID

	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`
}

func (r *DescribeMetricSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNotifyBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNotifyBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNotifyBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcServerCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *IdcInfoData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcServerCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcServerCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCustomAlarmListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 策略ID

		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
		// 策略名

		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
		// 消息列表

		AlarmList []*CustomAlarmList `json:"AlarmList,omitempty" name:"AlarmList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCustomAlarmListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomAlarmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmMetricSet struct {

	// CLS日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// CLS日志集名

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// CLS日志主题ID

	LogtopicId *string `json:"LogtopicId,omitempty" name:"LogtopicId"`
	// CLS日志主题名

	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`
	// 指标集名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 统计周期 可选：60、300

	StatisticCycle *int64 `json:"StatisticCycle,omitempty" name:"StatisticCycle"`
	// 时间字段

	TimeField *string `json:"TimeField,omitempty" name:"TimeField"`
	// 日志过滤规则

	LogFilterRules []*ClmLogFilterRule `json:"LogFilterRules,omitempty" name:"LogFilterRules"`
	// 基础指标规则

	MetricBasicItems []*ClmMetricBasicItem `json:"MetricBasicItems,omitempty" name:"MetricBasicItems"`
	// 复合指标规则

	MetricCustomItems []*ClmMetricCustomItem `json:"MetricCustomItems,omitempty" name:"MetricCustomItems"`
	// 维度规则

	Dimensions []*ClmDimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 条目ID（只作出参，入参不填）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// APPID（只作出参，入参不填）

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// UIN（只作出参，入参不填）

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 地域（只作出参，入参不填）

	Region *string `json:"Region,omitempty" name:"Region"`
	// 日志数据结构描述

	LogProfileItems []*ClmLogProfileItem `json:"LogProfileItems,omitempty" name:"LogProfileItems"`
}

type Condition struct {

	// 告警通知频率

	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
	// 重复通知策略预定义（0 - 只告警一次， 1 - 指数告警，2 - 连接告警）

	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 检测方式

	CalcType *string `json:"CalcType,omitempty" name:"CalcType"`
	// 检测值

	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`
	// 持续时间

	ContinueTime *string `json:"ContinueTime,omitempty" name:"ContinueTime"`
	// 指标ID

	MetricID *int64 `json:"MetricID,omitempty" name:"MetricID"`
	// 指标展示名称（对外）

	MetricDisplayName *string `json:"MetricDisplayName,omitempty" name:"MetricDisplayName"`
	// 周期

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 规则ID

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
	// 指标单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type DescribePolicyGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略组列表

		GroupList []*DescribePolicyGroupListGroup `json:"GroupList,omitempty" name:"GroupList"`
		// 策略组总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyQuotaV3ResponseDataQuota struct {

	// 用户

	Used *int64 `json:"Used,omitempty" name:"Used"`
	// 总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
}

type DeleteInstanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmDimensionValuesRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控类型过滤 "MT_QCE"=云产品监控 "MT_CUSTOM"=自定义监控

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 页数，从 1 开始计数，默认 1

	PageNumber *uint64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页的数量，取值1~100，默认20

	PageSize *uint64 `json:"PageSize,omitempty" name:"PageSize"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 维度名

	DimensionName *string `json:"DimensionName,omitempty" name:"DimensionName"`
	// 按维度值模糊搜索

	DimensionValue *string `json:"DimensionValue,omitempty" name:"DimensionValue"`
}

func (r *DescribeAlarmDimensionValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmDimensionValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeAllServerRequest struct {
	*tchttp.BaseRequest

	// 属性ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 查询关键字：服务器ID、服务器名称或服务器ip 搜索

	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
}

func (r *DescribeAttributeAllServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttributeAllServerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyStatesRequest struct {
	*tchttp.BaseRequest

	// 更新启停状态信息

	Strategy []*StrategyStatesInfo `json:"Strategy,omitempty" name:"Strategy"`
}

func (r *ModifyStrategyStatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrategyStatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListConfigManualCalcType struct {

	// CalcType 取值

	Keys []*int64 `json:"Keys,omitempty" name:"Keys"`
	// 是否必须

	Need *bool `json:"Need,omitempty" name:"Need"`
}

type DescribeMonitorDataByAlarmIDDataPoint struct {

	// 是否为空

	IsNull *bool `json:"IsNull,omitempty" name:"IsNull"`
	// 数据值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type MessagePolicy struct {

	// 策略ID

	PolicyID *string `json:"PolicyID,omitempty" name:"PolicyID"`
	// 策略名

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 接收组ID列表

	ReceiverGroupIds []*int64 `json:"ReceiverGroupIds,omitempty" name:"ReceiverGroupIds"`
	// 通知方式

	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`
	// 语音配置

	VoiceConfig *DescribeMsgPolicyListVoiceConfig `json:"VoiceConfig,omitempty" name:"VoiceConfig"`
	// 是否为默认策略

	IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
	// 告警次数

	AlarmCount *int64 `json:"AlarmCount,omitempty" name:"AlarmCount"`
}

type CreateModuleRequest struct {
	*tchttp.BaseRequest

	// 父分组Id

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// 叶子分组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分组允许添加的实例类型

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *CreateModuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateModuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicyBasic struct {

	// 监控类型

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 告警策略1.0 Id

	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
	// 告警策略2.0 Id

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 告警策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
}

type DescribePolicyObjectCountRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribePolicyObjectCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyObjectCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductByIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品列表

		ProductList []*MonitorProductInfo `json:"ProductList,omitempty" name:"ProductList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorProductByIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorProductByIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MiniDashboardAlarmInfo struct {

	// 维度信息

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 无告警策略

	NoPolicy *bool `json:"NoPolicy,omitempty" name:"NoPolicy"`
	// 没有通知规则的默认策略

	DefaultPolicyWithoutNotice *string `json:"DefaultPolicyWithoutNotice,omitempty" name:"DefaultPolicyWithoutNotice"`
}

type CreateDashboardRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控面板名

	DescName *string `json:"DescName,omitempty" name:"DescName"`
	// json描述元数据

	Meta *string `json:"Meta,omitempty" name:"Meta"`
	// 产品类型

	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
	// CustomID

	CustomID *string `json:"CustomID,omitempty" name:"CustomID"`
}

func (r *CreateDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警事件列表

		Events []*AlarmEvent `json:"Events,omitempty" name:"Events"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeProjectsListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectsListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerDatasRequest struct {
	*tchttp.BaseRequest

	// 服务器ID列表

	ServerId []*uint64 `json:"ServerId,omitempty" name:"ServerId"`
	// 指标ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 开始时间。例如："2019-09-11 00:00:00"

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。例如："2019-09-11 00:00:00"

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeServerDatasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerDatasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServerAttributeDataOutput struct {

	// 指标ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 服务器ID

	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`
	// 指标上报值

	Values []*AttributeTimestampValueOutput `json:"Values,omitempty" name:"Values"`
}

type AddDashboardPanelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Result []*string `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddDashboardPanelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddDashboardPanelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindingPolicyObjectDimension struct {

	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域名

	Region *string `json:"Region,omitempty" name:"Region"`
	// 维度信息

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 事件维度信息

	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type DescribeCCMGroupViewAttributeRequest struct {
	*tchttp.BaseRequest

	// 实例分组 ID。GroupId或ViewId必选其一，若ViewId存在，则以ViewId为准

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 分组视图ID

	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`
	// 查询关键字

	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
}

func (r *DescribeCCMGroupViewAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMGroupViewAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InitGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回整棵分组树

		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InitGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InitGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
	// 通知模板名称 60字符以内

	Name *string `json:"Name,omitempty" name:"Name"`
	// 通知类型 ALARM=未恢复通知 OK=已恢复通知 ALL=都通知

	NoticeType *string `json:"NoticeType,omitempty" name:"NoticeType"`
	// 通知语言 zh-CN=中文 en-US=英文

	NoticeLanguage *string `json:"NoticeLanguage,omitempty" name:"NoticeLanguage"`
	// 用户通知 最多5个

	UserNotices []*UserNotice `json:"UserNotices,omitempty" name:"UserNotices"`
	// 回调通知 最多3个

	URLNotices []*URLNotice `json:"URLNotices,omitempty" name:"URLNotices"`
	// 其他非公开通知渠道列表

	ExtraChannels []*string `json:"ExtraChannels,omitempty" name:"ExtraChannels"`
	// 大客户通知 最多1个

	BigCustomerNotices []*BigCustomerNotice `json:"BigCustomerNotices,omitempty" name:"BigCustomerNotices"`
	// CLS通知 最多1个

	CLSNotices []*CLSNotice `json:"CLSNotices,omitempty" name:"CLSNotices"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCCMChartRequest struct {
	*tchttp.BaseRequest

	// 监控图表名称

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
	// 监控图表类型。1 明细视图， 2 聚合视图

	ChartType *int64 `json:"ChartType,omitempty" name:"ChartType"`
	// 产品类型。1 基础监控， 2 自定义监控

	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`
	// 指标ID

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 聚合方式。0 SUM， 1 AVG， 2 MAX， 3 MIN

	Aggregation *int64 `json:"Aggregation,omitempty" name:"Aggregation"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 分组ID

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *CreateCCMChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCCMChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTkeDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回字段列表

		Columns []*string `json:"Columns,omitempty" name:"Columns"`
		// 数据。需要单独对Data部分进行一次JSON解码

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetTkeDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTkeDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendCustomAlarmMsgRequest struct {
	*tchttp.BaseRequest

	// 接口模块名，当前取值monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 消息策略ID，在云监控自定义消息页面配置

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 用户想要发送的自定义消息内容

	Msg *string `json:"Msg,omitempty" name:"Msg"`
}

func (r *SendCustomAlarmMsgRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendCustomAlarmMsgRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmAlertTriggerRule struct {

	// 触发条件ID (出参)

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间 （出参）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（出参）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 关联的告警策略主键id

	AlertPolicyId *int64 `json:"AlertPolicyId,omitempty" name:"AlertPolicyId"`
	// 指标名称（英文）

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标类型 1=普通指标 2=复合指标

	MetricType *int64 `json:"MetricType,omitempty" name:"MetricType"`
	// 操作符

	Operating *string `json:"Operating,omitempty" name:"Operating"`
	// 指标阈值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 持续周期个数

	ContinuousCycleCount *int64 `json:"ContinuousCycleCount,omitempty" name:"ContinuousCycleCount"`
	// 通知频率（通知间隔秒数）

	NoticeFrequencySec *int64 `json:"NoticeFrequencySec,omitempty" name:"NoticeFrequencySec"`
	// 复合指标表达式 （基础指标传空）

	MetricFormula *string `json:"MetricFormula,omitempty" name:"MetricFormula"`
}

type DescribeAlarmSmsQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 总使用量

		Used *int64 `json:"Used,omitempty" name:"Used"`
		// 短信配额信息列表

		QuotaList []*DescribeAlarmSmsQuotaQuota `json:"QuotaList,omitempty" name:"QuotaList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmSmsQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmSmsQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardMetricsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回指标数组

		Data []*DashboardMetricEntry `json:"Data,omitempty" name:"Data"`
		// IsShowCategory

		IsShowCategory *int64 `json:"IsShowCategory,omitempty" name:"IsShowCategory"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardMetricsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardMetricsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTopicIndexRequest struct {
	*tchttp.BaseRequest

	// 无

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeLogTopicIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogTopicIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardDataSourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据源列表

		Data []*DashboardDataSource `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardDataSourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardDataSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmMetricAnalysisCurrentData struct {

	// 当前数据

	Current []*PCLMMetricAnalysisPoint `json:"Current,omitempty" name:"Current"`
}

type PeriodsSt struct {

	// 周期

	Period *string `json:"Period,omitempty" name:"Period"`
	// 统计方式

	StatType []*string `json:"StatType,omitempty" name:"StatType"`
}

type DescribeAlarmMetricsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控类型过滤 "MT_QCE"=云产品监控

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 告警策略类型，由 DescribeAllNamespaces 获得，例如 cvm_device

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

func (r *DescribeAlarmMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSortObjectListRequest struct {
	*tchttp.BaseRequest

	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 排序方式，升序（“asc”）或者降序（“desc”）。默认“desc”

	Order *string `json:"Order,omitempty" name:"Order"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 待排序的指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 时间，精确到分钟，格式：“2019-01-01 05:23:00”

	Time *string `json:"Time,omitempty" name:"Time"`
	// 数据聚合周期，单位秒。不填则默认相应指标的默认周期

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 指定的维度组合

	Dimensions []*DescribeSortObjectListDimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeSortObjectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSortObjectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlertPolicyRequest struct {
	*tchttp.BaseRequest

	// 要删除的策略ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteAlertPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlertPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 实例组Id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 按实例过滤

	InstanceList []*BindingPolicyObjectDimension `json:"InstanceList,omitempty" name:"InstanceList"`
	// 按地域过滤

	Regions []*string `json:"Regions,omitempty" name:"Regions"`
	// 返回的实例列表的分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回的实例列表的分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeInstanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardViewRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图ID

	ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`
	// 名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 监控面板ID

	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
	// 监控面板视图名

	DescName *string `json:"DescName,omitempty" name:"DescName"`
	// 面板实例列表，json字符串数据

	Instances []*string `json:"Instances,omitempty" name:"Instances"`
	// 元数据，json串

	Meta *string `json:"Meta,omitempty" name:"Meta"`
	// 指标名列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
}

func (r *ModifyDashboardViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcTreeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*IdcData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdcTreeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcTreeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStrategyRequest struct {
	*tchttp.BaseRequest

	// 告警策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 告警策略启停状态

	IsStart *int64 `json:"IsStart,omitempty" name:"IsStart"`
	// 告警类型

	StrategyAlarmType *int64 `json:"StrategyAlarmType,omitempty" name:"StrategyAlarmType"`
	// 告警对象类型

	MixType *int64 `json:"MixType,omitempty" name:"MixType"`
	// 当MixType为服务器时，标记服务器子类型

	MixSubType *int64 `json:"MixSubType,omitempty" name:"MixSubType"`
	// 告警对象类型ID

	MixId []*int64 `json:"MixId,omitempty" name:"MixId"`
	// 告警接收类型

	ReceiverType *int64 `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 告警接收用户或用户组ID列表

	ReceiverId []*uint64 `json:"ReceiverId,omitempty" name:"ReceiverId"`
	// 告警有效起始时间

	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" name:"EffectiveStartTime"`
	// 告警有效结束事件

	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" name:"EffectiveEndTime"`
	// 告警渠道列表

	AlarmChannel []*string `json:"AlarmChannel,omitempty" name:"AlarmChannel"`
	// 告警回调url

	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 策略的告警规则列表

	Alarm []*StrategyEntryAlarm `json:"Alarm,omitempty" name:"Alarm"`
}

func (r *CreateStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeContactListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 接收人数量

		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`
		// 接收人列表

		List []*Contact `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContactListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContactListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额描述列表

		Data []*DescribePolicyQuotaV3ResponseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStrategysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *StrategyData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStrategysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStrategysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSetsRequest struct {
	*tchttp.BaseRequest

	// 指标集ID

	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
	// 指标集名称（模糊搜索）

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志集名称（模糊搜索）

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题名称（模糊搜索）

	LogtopicName *string `json:"LogtopicName,omitempty" name:"LogtopicName"`
	// 偏移量，默认为0

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回数量，默认为20，最大100

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeMetricSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConditionsTemplateListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 模板列表

		TemplateGroupList []*TemplateGroup `json:"TemplateGroupList,omitempty" name:"TemplateGroupList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConditionsTemplateListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConditionsTemplateListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMDimensionAnalysisResultData struct {

	// 维度数据

	Current *string `json:"Current,omitempty" name:"Current"`
	// 粒度

	Granularity *int64 `json:"Granularity,omitempty" name:"Granularity"`
}

type DescribeUnifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dashboard数据

		Data *string `json:"Data,omitempty" name:"Data"`
		// 是否默认

		IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
		// 是否收藏

		IsStarred *bool `json:"IsStarred,omitempty" name:"IsStarred"`
		// 文件夹uuid

		FolderUUID *string `json:"FolderUUID,omitempty" name:"FolderUUID"`
		// 是否文件夹

		IsFolder *bool `json:"IsFolder,omitempty" name:"IsFolder"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUnifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceGroupListPolicyGroup struct {

	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type MetricData struct {

	// 无

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 无

	Points *string `json:"Points,omitempty" name:"Points"`
}

type ModifyCCMChartRequest struct {
	*tchttp.BaseRequest

	// 监控图表名称

	ChartId *int64 `json:"ChartId,omitempty" name:"ChartId"`
	// 监控图表类型。1 明细视图， 2 聚合视图

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
	// 产品类型。1 基础监控， 2 自定义监控

	ChartType *int64 `json:"ChartType,omitempty" name:"ChartType"`
	// 指标ID

	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`
	// 聚合方式。0 SUM， 1 AVG， 2 MAX， 3 MIN

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 描述

	Aggregation *int64 `json:"Aggregation,omitempty" name:"Aggregation"`
	// 分组ID

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyCCMChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCCMChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPreferenceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// j结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetPreferenceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPreferenceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentStatusHistory struct {

	// 时间戳

	Timestamp *string `json:"Timestamp,omitempty" name:"Timestamp"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeDashboardDimensionValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度值数组

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardDimensionValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardDimensionValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略条件列表

		Conditions []*DescribePolicyConditionListCondition `json:"Conditions,omitempty" name:"Conditions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyConditionListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyConditionListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceGroupInfo struct {

	// 实例组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 实例组ID

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 实例数量

	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 视图名/策略类型

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 最后更新UIN

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
}

type PCLMAlertPolicy struct {

	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标集CID

	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
	// 状态 1=已开启 2=未开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 告警接收组ID 逗号分隔

	AlertGroupId *string `json:"AlertGroupId,omitempty" name:"AlertGroupId"`
	// 告警接收渠道  "SMS", "EMAIL", "WECHAT", "CALL" 多个逗号分隔

	AlertChannel *string `json:"AlertChannel,omitempty" name:"AlertChannel"`
	// 通知时段开始时间（从00:00:00开始计算的秒数）

	NoticePeriodBegin *int64 `json:"NoticePeriodBegin,omitempty" name:"NoticePeriodBegin"`
	// 通知时段结束时间（从00:00:00开始计算的秒数）

	NoticePeriodEnd *int64 `json:"NoticePeriodEnd,omitempty" name:"NoticePeriodEnd"`
	// 回调url的scheme

	UrlScheme *string `json:"UrlScheme,omitempty" name:"UrlScheme"`
	// 回调url 不包含scheme部分

	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
	// 告警过滤条件列表

	AlertFilterRules []*PCLMAlertFilterRule `json:"AlertFilterRules,omitempty" name:"AlertFilterRules"`
	// 告警触发条件列表

	AlertTriggerRules []*PCLMAlertTriggerRule `json:"AlertTriggerRules,omitempty" name:"AlertTriggerRules"`
	// 告警策略CID （出参）

	CID *string `json:"CID,omitempty" name:"CID"`
	// 创建时间 （出参）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间 （出参）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 告警接收人ID 逗号分隔

	AlertUserId *string `json:"AlertUserId,omitempty" name:"AlertUserId"`
	// 最近告警时间（出参）

	LatestAlertTime *string `json:"LatestAlertTime,omitempty" name:"LatestAlertTime"`
	// 过滤条件之间的关系 1与 2或

	FilterRelation *int64 `json:"FilterRelation,omitempty" name:"FilterRelation"`
	// 触发条件之间的关系 1与 2或

	TriggerRelation *int64 `json:"TriggerRelation,omitempty" name:"TriggerRelation"`
}

type DescribePolicyConditionListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribePolicyConditionListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyConditionListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoReceiverInfo struct {

	// 告警接收组id列表

	ReceiverGroupList []*int64 `json:"ReceiverGroupList,omitempty" name:"ReceiverGroupList"`
	// 告警接收人id列表

	ReceiverUserList []*int64 `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`
	// 告警时间段开始时间。范围[0,86400)，作为unix时间戳转成北京时间后去掉日期，例如7200表示"10:0:0"

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 告警时间段结束时间。含义同StartTime

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 接收类型。“group”(接收组)或“user”(接收人)

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 告警通知方式。可选 "SMS","SITE","EMAIL","CALL","WECHAT"

	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`
	// 电话告警接收者uid

	UidList []*int64 `json:"UidList,omitempty" name:"UidList"`
	// 电话告警轮数

	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`
	// 电话告警每轮间隔（秒）

	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`
	// 电话告警对个人间隔（秒）

	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`
	// 是否需要电话告警触达提示。0不需要，1需要

	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`
	// 电话告警通知时机。可选"OCCUR"(告警时通知),"RECOVER"(恢复时通知)

	SendFor []*string `json:"SendFor,omitempty" name:"SendFor"`
	// 恢复通知方式。可选"SMS"

	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify"`
	// 接收语言，取值为 en-US 和 zh-CN

	ReceiveLanguage *string `json:"ReceiveLanguage,omitempty" name:"ReceiveLanguage"`
}

type ModifyAlertPolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 告警策略状态 1开启 2关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAlertPolicyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlertPolicyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StrategyEntry struct {

	// 告警策略ID

	StrategyId *int64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 告警策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 告警接受类型

	ReceiverType *int64 `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 告警接受用户或用户组ID列表

	ReceiverId []*uint64 `json:"ReceiverId,omitempty" name:"ReceiverId"`
	// 有效开始时间

	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" name:"EffectiveStartTime"`
	// 有效结束时间

	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" name:"EffectiveEndTime"`
	// 告警渠道列表。每个value可以为："SMS","EMAIL","WECHAT"或"CALL"

	AlarmChannel []*string `json:"AlarmChannel,omitempty" name:"AlarmChannel"`
	// 告警回调地址

	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建人ID

	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人ID

	UpdaterId *uint64 `json:"UpdaterId,omitempty" name:"UpdaterId"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 是否启动。 0：停止中， 1：启动中

	IsStart *int64 `json:"IsStart,omitempty" name:"IsStart"`
	// 告警对象类型。 1：服务器，2：视图

	MixType *uint64 `json:"MixType,omitempty" name:"MixType"`
	// 告警对象子类型。MixType为2，则MixId为视图ID列表，MixId为空数组[]， 则后台会自动查询默认视图ID；MixType为1（服务器）时，若MixSubType为1，MixId可不填；若MixSubType为2，则MixId为服务器ID列表，可含有多个服务器ID

	MixSubType *int64 `json:"MixSubType,omitempty" name:"MixSubType"`
	// 告警规则列表

	Alarm []*StrategyEntryAlarm `json:"Alarm,omitempty" name:"Alarm"`
	// 告警对象类型ID列表。

	MixId []*int64 `json:"MixId,omitempty" name:"MixId"`
	// 若mixtype为服务器类型，且子类型为指定服务器id。则存在Ip列表

	Ip []*string `json:"Ip,omitempty" name:"Ip"`
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警通知模板ID

		NoticeId *string `json:"NoticeId,omitempty" name:"NoticeId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警通知模板id列表

	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
}

func (r *DeleteAlarmNoticesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmNoticesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeAggregateDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *AttributeAggrValueInfoOutputData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAttributeAggregateDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttributeAggregateDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 监控数据

		Data []*MetricData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStatisticDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeUnitsRequest struct {
	*tchttp.BaseRequest

	// 指标单位ID列表

	UnitId []*uint64 `json:"UnitId,omitempty" name:"UnitId"`
	// 指标单位名称列表

	UnitName []*string `json:"UnitName,omitempty" name:"UnitName"`
	// 查询关键字：指标单位ID或指标单位名称搜索

	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
}

func (r *DescribeAttributeUnitsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttributeUnitsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardViewsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 视图列表

		ViewList []*DescribeDashboardViewList `json:"ViewList,omitempty" name:"ViewList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardViewsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePolicyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardStatusRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeDashboardStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCMGroupViewEntry struct {

	// 分组视图ID

	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`
	// 分组视图名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 分组视图是否自动绑定指标。0：否，1：是

	IsAuto *int64 `json:"IsAuto,omitempty" name:"IsAuto"`
}

type TransLogInstance struct {

	// 实例Dimension

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 实例UniqueID

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
	// 屏蔽状态，1表示已屏蔽，0表示未屏蔽

	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DeleteMsgPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 自定义消息策略ID

	PolicyID *string `json:"PolicyID,omitempty" name:"PolicyID"`
}

func (r *DeleteMsgPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMsgPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppFlowRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 名字空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 周期

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *DescribeAppFlowRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppFlowRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyQuestionnaireRequest struct {
	*tchttp.BaseRequest

	// 选项值

	Options []*int64 `json:"Options,omitempty" name:"Options"`
	// 建议

	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`
	// 电话号码

	PhoneNumber *string `json:"PhoneNumber,omitempty" name:"PhoneNumber"`
}

func (r *ModifyQuestionnaireRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyQuestionnaireRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductListForDescribeMonitorProducts struct {

	// 支持的地域

	AvaliableRegions []*string `json:"AvaliableRegions,omitempty" name:"AvaliableRegions"`
	// 产品的ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否在Dashboard展示

	IsShowInDashboard *bool `json:"IsShowInDashboard,omitempty" name:"IsShowInDashboard"`
	// JS配置

	Meta *string `json:"Meta,omitempty" name:"Meta"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 策略的视图名

	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
}

type DescribeProductEventListEventsDimensions struct {

	// 维度名（英文）

	Key *string `json:"Key,omitempty" name:"Key"`
	// 维度名（中文）

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyPolicyGroupEventCondition struct {

	// 规则id，不填表示新增，填写了ruleId表示在已存在的规则基础上进行修改

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 事件id

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 告警发送收敛类型。0连续告警，1指数告警

	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次

	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
}

type RegionPolicyObjectCount struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 绑定的实例数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
}

type ArgusAggType struct {

	// 聚合方式

	AggType *string `json:"AggType,omitempty" name:"AggType"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
}

type DescribeDashboardDimensionsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// Dashboard面板ID

	DashboardUUID *string `json:"DashboardUUID,omitempty" name:"DashboardUUID"`
}

func (r *DescribeDashboardDimensionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardDimensionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMidDimensionValueListRequest struct {
	*tchttp.BaseRequest

	// 固定值为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 要查询的维度

	DimensionKey *string `json:"DimensionKey,omitempty" name:"DimensionKey"`
	// 查询过滤条件

	Conditions []*MidQueryCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 起始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 查询的中台地域

	QueryRegion *string `json:"QueryRegion,omitempty" name:"QueryRegion"`
}

func (r *DescribeMidDimensionValueListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMidDimensionValueListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomMetisAbnormOutput struct {

	// 视图结果

	ViewAttrAbnormals []*CustomMetisAbnormOutputView `json:"ViewAttrAbnormals,omitempty" name:"ViewAttrAbnormals"`
}

type DescribeGroupIDByTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果返回

		Result []*GroupIdTag `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGroupIDByTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupIDByTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmMetricAnalysisCustomMetric struct {

	// 复合指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运算表达式

	Formula *string `json:"Formula,omitempty" name:"Formula"`
}

type ModifyMsgPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略id

		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMsgPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMsgPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyGroup struct {

	// 是否可设为默认告警策略

	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`
	// 告警策略组ID

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 告警策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 是否为默认告警策略

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 告警策略启用状态

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 最后修改人UIN

	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 未屏蔽的实例数

	NoShieldedInstanceCount *int64 `json:"NoShieldedInstanceCount,omitempty" name:"NoShieldedInstanceCount"`
	// 父策略组ID

	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`
	// 所属项目ID

	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
	// 告警接收对象信息

	ReceiverInfos []*PolicyGroupReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 修改时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 总绑定实例数

	TotalInstanceCount *int64 `json:"TotalInstanceCount,omitempty" name:"TotalInstanceCount"`
	// 视图

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 是否为与关系规则

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribeAlarmHistoryByAlarmIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警列表

		Alarms []*HistoryAlarmInfo `json:"Alarms,omitempty" name:"Alarms"`
		// 告警总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmHistoryByAlarmIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmHistoryByAlarmIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricAnalysisDataRequest struct {
	*tchttp.BaseRequest

	// 指标集id

	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`
	// 开始时间戳（秒）

	StartTimestamp *int64 `json:"StartTimestamp,omitempty" name:"StartTimestamp"`
	// 结束时间戳（秒）

	EndTimestamp *int64 `json:"EndTimestamp,omitempty" name:"EndTimestamp"`
	// 对比数据开始时间戳（秒）

	CStartTimestamp *int64 `json:"CStartTimestamp,omitempty" name:"CStartTimestamp"`
	// 对比数据结束时间戳（秒）

	CEndTimestamp *int64 `json:"CEndTimestamp,omitempty" name:"CEndTimestamp"`
	// 筛选条件

	Filters []*ClmAnalysisFilter `json:"Filters,omitempty" name:"Filters"`
	// 基础指标列表

	BasicMetrics []*ClmMetricAnalysisBasicMetric `json:"BasicMetrics,omitempty" name:"BasicMetrics"`
	// 复合指标列表

	CustomMetrics []*ClmMetricAnalysisCustomMetric `json:"CustomMetrics,omitempty" name:"CustomMetrics"`
}

func (r *DescribeMetricAnalysisDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricAnalysisDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopDashboardPanelRequest struct {
	*tchttp.BaseRequest

	// 置顶的图表

	Top *DashboardPanel `json:"Top,omitempty" name:"Top"`
	// 取消置顶的图表

	UnTop *DashboardPanel `json:"UnTop,omitempty" name:"UnTop"`
}

func (r *TopDashboardPanelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TopDashboardPanelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConditionsTemplateRequestEventCondition struct {

	// 告警通知周期

	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
	// 告警通知方式

	AlarmNotifyType *string `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 事件ID

	EventID *string `json:"EventID,omitempty" name:"EventID"`
	// 规则ID

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

type BindingPolicyTagRequest struct {
	*tchttp.BaseRequest

	// 固定取值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 策略标签

	Tag *PolicyTag `json:"Tag,omitempty" name:"Tag"`
	// 实例分组ID

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 产品类型

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
}

func (r *BindingPolicyTagRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindingPolicyTagRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteModuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已删除的分组节点

		Data []*CgrpModuleNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteModuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributeAggregateDataRequest struct {
	*tchttp.BaseRequest

	// 指标ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 指标聚合方式

	AttributeAggregation *int64 `json:"AttributeAggregation,omitempty" name:"AttributeAggregation"`
	// 地域列表

	IdcId []*int64 `json:"IdcId,omitempty" name:"IdcId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAttributeAggregateDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttributeAggregateDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略 ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *SetDefaultAlarmPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultAlarmPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicySituationRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 项目id列表

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
}

func (r *DescribePolicySituationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicySituationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimensions struct {

	// 实例维度名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ProductSimple struct {

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 产品中文名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品英文名称

	ProductEnName *string `json:"ProductEnName,omitempty" name:"ProductEnName"`
}

type DescribeAlarmHistoriesRequest struct {
	*tchttp.BaseRequest

	// 页数，从 1 开始计数，默认 1

	PageNumber *int64 `json:"PageNumber,omitempty" name:"PageNumber"`
	// 每页的数量，取值1~100，默认20

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 默认按首次出现时间倒序排列 "ASC"=正序 "DESC"=逆序

	Order *string `json:"Order,omitempty" name:"Order"`
	// 起始时间，默认一天前的时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认当前时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 根据监控类型过滤 不选默认查所有类型 "MT_QCE"=云产品监控 "MT_CUSTOM"=自定义监控 "MT_PROME"=prometheus监控

	MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`
	// 根据告警对象过滤 字符串模糊搜索

	AlarmObject *string `json:"AlarmObject,omitempty" name:"AlarmObject"`
	// 根据告警状态过滤 ALARM=未恢复 OK=已恢复 NO_CONF=已失效 NO_DATA=数据不足，不选默认查所有

	AlarmStatus []*string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
	// 根据项目ID过滤，-1=“-“项目 0=默认项目

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
	// 根据实例组ID过滤

	InstanceGroupIds []*int64 `json:"InstanceGroupIds,omitempty" name:"InstanceGroupIds"`
	// 根据策略类型过滤

	Namespaces *MonitorTypeNamespace `json:"Namespaces,omitempty" name:"Namespaces"`
	// 根据策略名称模糊搜索

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 根据告警内容模糊搜索

	Content *string `json:"Content,omitempty" name:"Content"`
	// 根据接收人搜索

	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids"`
	// 根据接收组搜索

	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 根据告警策略 Id 列表搜索

	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
	// 根据指标名过滤

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 根据告警实例 UniqueId 搜索

	UniqueIds []*string `json:"UniqueIds,omitempty" name:"UniqueIds"`
	// 实例的维度信息

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *DescribeAlarmHistoriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmHistoriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控面板名

	DescName *string `json:"DescName,omitempty" name:"DescName"`
	// json描述元数据

	Meta *string `json:"Meta,omitempty" name:"Meta"`
	// 监控面板ID

	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
}

func (r *ModifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardNamespace struct {

	// 命名空间标示

	Id *string `json:"Id,omitempty" name:"Id"`
	// 命名空间名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 配置信息

	Config *string `json:"Config,omitempty" name:"Config"`
	// 支持地域列表

	AvailableRegions []*string `json:"AvailableRegions,omitempty" name:"AvailableRegions"`
	// 排序Id

	SortId *int64 `json:"SortId,omitempty" name:"SortId"`
	// 命名空间类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// DashboardId

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
}

type DescribeCCMInstanceDatasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*CCMInstanceAttributeDataOutput `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMInstanceDatasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMInstanceDatasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 告警类型

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 指标告警条件的且或关系，1表示且告警，所有指标告警条件都达到才告警，0表示或告警，任意指标告警条件达到都告警

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
	// 指标告警条件规则，不填表示删除已有的所有指标告警条件规则

	Conditions []*ModifyPolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 事件告警条件，不填表示删除已有的事件告警条件

	EventConditions []*ModifyPolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions"`
	// 模板策略组id

	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`
}

func (r *ModifyPolicyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolicyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警Id

		Id *int64 `json:"Id,omitempty" name:"Id"`
		// 策略组Id

		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
		// 策略组名

		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
		// 告警状态，OK、ALARM、NO_CONF、NO_DATA

		AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
		// 告警类型，0（指标告警），2（产品事件告警），3（平台事件告警）

		AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`
		// 阈值

		CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`
		// 当前值

		CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`
		// 告警内容

		Content *string `json:"Content,omitempty" name:"Content"`
		// 首次触发时间

		FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`
		// 最后触发时间

		LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`
		// 持续时长

		Duration *int64 `json:"Duration,omitempty" name:"Duration"`
		// 对象Id

		ObjId *string `json:"ObjId,omitempty" name:"ObjId"`
		// 对象名

		ObjName *string `json:"ObjName,omitempty" name:"ObjName"`
		// 项目Id

		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
		// 项目名

		ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
		// 地域

		Region *string `json:"Region,omitempty" name:"Region"`
		// 取值 0，1，2，3，4，5，其中0表示未恢复，1表示已恢复，4表示已失效，2/3/5表示数据不足。

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 视图名

		ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
		// 通知方式

		NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBasicAlarmListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警列表

		Alarms []*DescribeBasicAlarmListAlarms `json:"Alarms,omitempty" name:"Alarms"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBasicAlarmListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBasicAlarmListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 分页参数，每页返回的数量，取值1~100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 按策略名搜索

	Like *string `json:"Like,omitempty" name:"Like"`
	// 实例分组id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 按更新时间排序, asc 或者 desc

	UpdateTimeOrder *string `json:"UpdateTimeOrder,omitempty" name:"UpdateTimeOrder"`
	// 项目id列表

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
	// 告警策略类型列表

	ViewNames []*string `json:"ViewNames,omitempty" name:"ViewNames"`
	// 是否过滤无接收人策略组, 1表示过滤, 0表示不过滤

	FilterUnuseReceiver *int64 `json:"FilterUnuseReceiver,omitempty" name:"FilterUnuseReceiver"`
	// 过滤条件, 接收组列表

	Receivers []*string `json:"Receivers,omitempty" name:"Receivers"`
	// 过滤条件, 接收人列表

	ReceiverUserList []*string `json:"ReceiverUserList,omitempty" name:"ReceiverUserList"`
	// 维度组合字段(json字符串), 例如[[{"name":"unInstanceId","value":"ins-6e4b2aaa"}]]

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 模板策略组id, 多个id用逗号分隔

	ConditionTempGroupId *string `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`
	// 过滤条件, 接收人或者接收组, user表示接收人, group表示接收组

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
}

func (r *DescribePolicyGroupListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyGroupListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMGroupViewAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*AttributeInfoOutput `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMGroupViewAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMGroupViewAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicy struct {

	// 告警策略 ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 告警策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 监控类型 MT_QCE=云产品监控

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 启停状态 0=停用 1=启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 策略组绑定的实例数

	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`
	// 项目 Id -1=无项目 0=默认项目

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 告警策略类型

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 即 viewName / 视图 云产品监控为 cvm_device 等 自定义监控为 Namespace 下的 Measurement

	Measurement *string `json:"Measurement,omitempty" name:"Measurement"`
	// 裸写包含触发规则、过滤条件等内容的告警规则，适用于 prometheus 监控

	PolicyExpression *string `json:"PolicyExpression,omitempty" name:"PolicyExpression"`
	// 触发条件模板 Id

	ConditionTemplateId *string `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`
	// 指标触发条件

	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`
	// 事件触发条件

	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`
	// 通知规则 id 列表

	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
	// 通知规则 列表

	Notices []*AlarmNotice `json:"Notices,omitempty" name:"Notices"`
	// 触发任务列表

	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`
	// 字段废弃：用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID

	OriginID *string `json:"OriginID,omitempty" name:"OriginID"`
	// 模板策略组
	// 注意：此字段可能返回 null，表示取不到有效值。

	ConditionsTemp *ConditionsTemp `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`
	// 最后编辑的用户uin

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 更新时间
	// 注意：此字段可能返回 null，表示取不到有效值。

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 地域

	Region []*string `json:"Region,omitempty" name:"Region"`
	// namespace显示名字

	NamespaceShowName *string `json:"NamespaceShowName,omitempty" name:"NamespaceShowName"`
	// 是否默认策略，1是，0否

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 能否设置默认策略，1是，0否

	CanSetDefault *int64 `json:"CanSetDefault,omitempty" name:"CanSetDefault"`
	// 实例分组ID

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 实例分组总实例数

	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 实例分组名称

	InstanceGroupName *string `json:"InstanceGroupName,omitempty" name:"InstanceGroupName"`
	// JsConfig 前台配置

	JsConfig *string `json:"JsConfig,omitempty" name:"JsConfig"`
	// 触发条件类型 STATIC=静态阈值 DYNAMIC=动态类型

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 用于实例、实例组绑定和解绑接口（BindingPolicyObject、UnBindingAllPolicyObject、UnBindingPolicyObject）的策略 ID

	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
	// 标签

	TagInstances []*TagInstance `json:"TagInstances,omitempty" name:"TagInstances"`
	// Filter

	Filter *string `json:"Filter,omitempty" name:"Filter"`
	// GroupBy

	GroupBy *string `json:"GroupBy,omitempty" name:"GroupBy"`
}

type DescribeProductHealthStatusListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 按视图名过滤，比如"cvm_device"

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 1表示拉取所有地域，0相反

	AllRegions *int64 `json:"AllRegions,omitempty" name:"AllRegions"`
	// 根据项目id过滤

	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`
}

func (r *DescribeProductHealthStatusListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductHealthStatusListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmCallbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmCallbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopDashboardPanelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TopDashboardPanelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TopDashboardPanelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
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

type DescribeAlarmCallbackVerifyCodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 验证码

		VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmCallbackVerifyCodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmCallbackVerifyCodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorTypesRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeMonitorTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CgrpModuleNode struct {

	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 分组树高度

	High *int64 `json:"High,omitempty" name:"High"`
	// 叶子分组Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 第一层分组Id

	L1Id *int64 `json:"L1Id,omitempty" name:"L1Id"`
	// 第一层分组名称

	L1Name *string `json:"L1Name,omitempty" name:"L1Name"`
	// 第二层分组Id

	L2Id *int64 `json:"L2Id,omitempty" name:"L2Id"`
	// 第二层分组名称

	L2Name *string `json:"L2Name,omitempty" name:"L2Name"`
	// 第三层分组Id

	L3Id *int64 `json:"L3Id,omitempty" name:"L3Id"`
	// 第三层分组名称

	L3Name *string `json:"L3Name,omitempty" name:"L3Name"`
	// 叶子分组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备份负责人

	OwnerBack *string `json:"OwnerBack,omitempty" name:"OwnerBack"`
	// 主要负责人

	OwnerMain *string `json:"OwnerMain,omitempty" name:"OwnerMain"`
	// 父分组Id

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// SubUin

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 叶子分组允许添加的实例类型

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type PCLMMetricAnalysisBasicMetric struct {

	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ClonePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 被复制的策略组Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 是否复制回调方法，0不复制，1复制

	CopyAlarmMethod *int64 `json:"CopyAlarmMethod,omitempty" name:"CopyAlarmMethod"`
}

func (r *ClonePolicyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ClonePolicyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmCallbackHistoryRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmCallbackHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmCallbackHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeUserInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 组Id

		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPolicyGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyPolicyGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetPreferenceRequest struct {
	*tchttp.BaseRequest

	// light|dark

	Theme *string `json:"Theme,omitempty" name:"Theme"`
	// 默认dashboard UUID

	HomeDashboardUUID *string `json:"HomeDashboardUUID,omitempty" name:"HomeDashboardUUID"`
	// 时区

	Timezone *string `json:"Timezone,omitempty" name:"Timezone"`
}

func (r *SetPreferenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetPreferenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CgrpInstanceNode struct {

	// 自增Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例类型

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 叶子分组Id

	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
	// 叶子分组名称

	ModuleName *string `json:"ModuleName,omitempty" name:"ModuleName"`
	// 一级分组Id

	L1Id *int64 `json:"L1Id,omitempty" name:"L1Id"`
	// 二级分组Id

	L2Id *int64 `json:"L2Id,omitempty" name:"L2Id"`
	// 三级分组Id

	L3Id *int64 `json:"L3Id,omitempty" name:"L3Id"`
	// appId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// SubUin

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 分组数组

	Group []*CgrpInstanceGroupNode `json:"Group,omitempty" name:"Group"`
}

type DeleteDashboardPanelsRequest struct {
	*tchttp.BaseRequest

	// dashboard uuid

	DashboardUUID *string `json:"DashboardUUID,omitempty" name:"DashboardUUID"`
	// panel id列表

	PanelIDList []*int64 `json:"PanelIDList,omitempty" name:"PanelIDList"`
}

func (r *DeleteDashboardPanelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardPanelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInstance struct {

	// 标签Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签Value

	Value *string `json:"Value,omitempty" name:"Value"`
	// 实例个数

	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 产品类型，如：cvm

	ServiceType *string `json:"ServiceType,omitempty" name:"ServiceType"`
	// 地域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
	// 绑定状态，2：绑定成功，1：绑定中

	BindingStatus *int64 `json:"BindingStatus,omitempty" name:"BindingStatus"`
	// 标签状态，2：标签存在，1：标签不存在

	TagStatus *int64 `json:"TagStatus,omitempty" name:"TagStatus"`
}

type DescribePolicyObjectCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否为多地域

		IsMultiRegion *bool `json:"IsMultiRegion,omitempty" name:"IsMultiRegion"`
		// 地域统计列表

		RegionList []*RegionPolicyObjectCount `json:"RegionList,omitempty" name:"RegionList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyObjectCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyObjectCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateTagsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 模板变量的标签列表

		TemplateTags []*TemplateTag `json:"TemplateTags,omitempty" name:"TemplateTags"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplateTagsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateTagsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDimensionAnalysisDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度数据结果

		Data *PCLMDimensionAnalysisResultData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDimensionAnalysisDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDimensionAnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupCountRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 实例的维度列表

	Batch []*Instance `json:"Batch,omitempty" name:"Batch"`
}

func (r *DescribePolicyGroupCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyGroupCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件配置列表

		AccidentConfigs []*DescribeAccidentConfigAccident `json:"AccidentConfigs,omitempty" name:"AccidentConfigs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccidentConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccidentConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CgrpInstance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 类型

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 叶子分组节点

	ModuleId *int64 `json:"ModuleId,omitempty" name:"ModuleId"`
}

type DescribePolicyGroupListGroupInstanceGroup struct {

	// 实例分组名称id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 策略类型视图名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 最近编辑的用户uin

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 实例分组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 实例数量

	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type PutMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 上报时自行指定的 IP

	AnnounceIp *string `json:"AnnounceIp,omitempty" name:"AnnounceIp"`
	// 上报时自行指定的时间戳

	AnnounceTimestamp *uint64 `json:"AnnounceTimestamp,omitempty" name:"AnnounceTimestamp"`
	// 一组指标和数据

	Metrics []*MetricDatum `json:"Metrics,omitempty" name:"Metrics"`
	// 上报时自行指定的 IP 或 产品实例ID

	AnnounceInstance *string `json:"AnnounceInstance,omitempty" name:"AnnounceInstance"`
}

func (r *PutMonitorDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutMonitorDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentStatusHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 子机状态信息列表

		Data []*DescribeAgentStatusHistoryData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentStatusHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentStatusHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGraphDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询起始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 查询结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 周期

		Period *int64 `json:"Period,omitempty" name:"Period"`
		// 指标名

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 分区使用列表

		Partitions []*DescribeGraphDataPartition `json:"Partitions,omitempty" name:"Partitions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGraphDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGraphDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerDataRequest struct {
	*tchttp.BaseRequest

	// 服务器ID

	ServerId *uint64 `json:"ServerId,omitempty" name:"ServerId"`
	// 指标ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 开始时间。例如：1970-01-01 00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 截止时间。例如：1970-01-01 00:00:00

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeServerDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmEvent struct {

	// 事件名

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 展示的事件名

	Description *string `json:"Description,omitempty" name:"Description"`
	// 告警策略类型

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 可用于展示的提示信息

	Tooltip *MetricTooltip `json:"Tooltip,omitempty" name:"Tooltip"`
}

type DescribeAlarmDimensionValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度值列表

		DimensionValues []*string `json:"DimensionValues,omitempty" name:"DimensionValues"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmDimensionValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmDimensionValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionsTemp struct {

	// 告警条件id

	ConditionId *int64 `json:"ConditionId,omitempty" name:"ConditionId"`
	// 模版名称

	TemplateName *string `json:"TemplateName,omitempty" name:"TemplateName"`
	// 指标触发条件

	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`
	// 事件触发条件

	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`
}

type TemplateTag struct {

	// 标签名

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 标签类型

	TagType *string `json:"TagType,omitempty" name:"TagType"`
	// 关联的产品ID（基础监控）

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
}

type SearchDimensionValueRequest struct {
	*tchttp.BaseRequest

	// 指标集ID

	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`
	// 维度字段

	DimensionField *string `json:"DimensionField,omitempty" name:"DimensionField"`
	// 搜索内容

	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *SearchDimensionValueRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDimensionValueRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConditionsTemplateRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否为与关系

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 父ID

	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`
	// 是否屏蔽

	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`
	// 指标告警条件

	Conditions []*ModifyConditionsTemplateRequestCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 事件告警条件

	EventConditions []*ModifyConditionsTemplateRequestEventCondition `json:"EventConditions,omitempty" name:"EventConditions"`
}

func (r *CreateConditionsTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConditionsTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertPoliciesRequest struct {
	*tchttp.BaseRequest

	// 分页offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 排序字段和方式 （id desc）

	Order *string `json:"Order,omitempty" name:"Order"`
	// 按ID精确搜索

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 按名字模糊搜索

	Name *string `json:"Name,omitempty" name:"Name"`
	// 按指标集id搜索

	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
}

func (r *DescribeAlertPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略 ID 列表

	PolicyIds []*string `json:"PolicyIds,omitempty" name:"PolicyIds"`
}

func (r *DeleteAlarmPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlertPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功为true

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlertPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGroupRequest struct {
	*tchttp.BaseRequest

	// 分组Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 分组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BindingPolicyTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindingPolicyTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindingPolicyTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分组Id

		Data []*CgrpModuleNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeModuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUnifyDashboardRequest struct {
	*tchttp.BaseRequest

	// dashboard uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

func (r *DescribeUnifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyAlarmCallbackRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *VerifyAlarmCallbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyAlarmCallbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlertPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功为true

		Data *bool `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlertPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeModuleRequest struct {
	*tchttp.BaseRequest

	// 分组Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeModuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeModuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindingAllPolicyObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingAllPolicyObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindingAllPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertPolicyRequest struct {
	*tchttp.BaseRequest

	// 告警策略JSON对象

	AlertPolicy *ClmAlertPolicy `json:"AlertPolicy,omitempty" name:"AlertPolicy"`
}

func (r *CreateAlertPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardViewsInstances struct {

	// 地域ID

	RegionID *string `json:"RegionID,omitempty" name:"RegionID"`
	// 实例ID

	UnInstanceID *string `json:"UnInstanceID,omitempty" name:"UnInstanceID"`
}

type MonitorTypeNamespace struct {

	// 监控类型

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 策略类型值

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
}

type SetDefaultUnifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultUnifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultUnifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmMetricAnalysisBasicMetric struct {

	// 操作符

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ClmMetricBasicItem struct {

	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 指标统计方法：sum count min max

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标描述中文名

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 指标过滤规则

	StatisticFilterRules []*ClmStatisticFilterRule `json:"StatisticFilterRules,omitempty" name:"StatisticFilterRules"`
	// 条目ID（只作出参，入参不填）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 指标集ID（只作出参，入参不填）

	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type MonitorProductInfo struct {

	// 产品Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 产品名

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品视图名

	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`
	// 支持的地域

	AvaliableRegions []*string `json:"AvaliableRegions,omitempty" name:"AvaliableRegions"`
	// Dashboard是否可见

	IsShowInDashboard *bool `json:"IsShowInDashboard,omitempty" name:"IsShowInDashboard"`
	// 产品的详细信息，json字符串

	Meta *string `json:"Meta,omitempty" name:"Meta"`
	// 产品的namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
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

type CgrpInstanceGroupNode struct {

	// 分组Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 分组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeAppFlowResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 周期

		Period *int64 `json:"Period,omitempty" name:"Period"`
		// 指标名

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 数据点

		DataPoints []*DescribeAppFlowConverterResponseDataPoint `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAppFlowResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAppFlowResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 实例组Id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
}

func (r *CopyInstanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyInstanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardFolderRequest struct {
	*tchttp.BaseRequest

	// 文件夹详情

	Data *string `json:"Data,omitempty" name:"Data"`
	// 是否强制覆盖

	Overwrite *string `json:"Overwrite,omitempty" name:"Overwrite"`
}

func (r *ModifyDashboardFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupCondition struct {

	// 指标Id

	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`
	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等。如果指标有配置默认比较类型值可以不填。

	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`
	// 比较的值，如果指标不必须CalcValue可不填

	CalcValue *float64 `json:"CalcValue,omitempty" name:"CalcValue"`
	// 数据聚合周期(单位秒)，若指标有默认值可不填

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

type CustomMetisAbnormlInput struct {

	// 自定义监控视图ID

	ViewId *uint64 `json:"ViewId,omitempty" name:"ViewId"`
	// 自定义监控属性ID

	AttrId []*uint64 `json:"AttrId,omitempty" name:"AttrId"`
}

type DescribeAlarmEventsRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略类型，由 DescribeAllNamespaces 获得，例如 cvm_device

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 监控类型，如 MT_QCE。如果不填默认为 MT_QCE

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
}

func (r *DescribeAlarmEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeSubscribeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindingPolicyObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnBindingPolicyObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindingPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStorageDurationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 存储周期及持续时间列表

		Data []*DescribeStorageDurationData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeStorageDurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStorageDurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyTasksRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略 ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 告警策略触发任务列表，空数据代表解绑

	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`
}

func (r *ModifyAlarmPolicyTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeInfoOutputData struct {

	// 属性数据列表

	Data []*AttributeInfoOutput `json:"Data,omitempty" name:"Data"`
	// 属性数据列表个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type DeleteAttributesRequest struct {
	*tchttp.BaseRequest

	// 指标ID列表

	AttributeId []*uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
}

func (r *DeleteAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmSmsQuotaRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeAlarmSmsQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmSmsQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcTreeRequest struct {
	*tchttp.BaseRequest

	// 0或者为空拉取一级业务树，传其他id表示二三四级业务树

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeIdcTreeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcTreeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *CreateServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyUseList struct {

	// 告警对象唯一id

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
	// 维度字符串

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type MonitorSelectedDimension struct {

	// 维度名

	DimensionName *string `json:"DimensionName,omitempty" name:"DimensionName"`
	// 维度值

	DimensionValue *string `json:"DimensionValue,omitempty" name:"DimensionValue"`
}

type DescribeMonitorAlarmSmsQuotaRequest struct {
	*tchttp.BaseRequest

	// 告警类型

	SmsType []*string `json:"SmsType,omitempty" name:"SmsType"`
	// 模块 /monitor/v1

	Module []*string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeMonitorAlarmSmsQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorAlarmSmsQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计周期

		Period *uint64 `json:"Period,omitempty" name:"Period"`
		// 指标名

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 数据点数组

		DataPoints []*DataPoint `json:"DataPoints,omitempty" name:"DataPoints"`
		// 开始时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 结束时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
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

type DescribeDashboardMetricDataRequest struct {
	*tchttp.BaseRequest

	// 所属模块，默认为 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 应用空间SpaceUUID，默认值为space_default

	SpaceUUID *string `json:"SpaceUUID,omitempty" name:"SpaceUUID"`
	// 指标查询结构体，支持多指标查询

	Query []*DashboardMetricQuery `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeDashboardMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeViewDataRequest struct {
	*tchttp.BaseRequest

	// 视图ID

	ViewId *uint64 `json:"ViewId,omitempty" name:"ViewId"`
	// 指标ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 开始时间。例如："2019-09-11 00:00:00"

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。例如："2019-09-11 00:00:00"

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 地域ID列表

	IdcId []*uint64 `json:"IdcId,omitempty" name:"IdcId"`
	// 聚合方式。不传默认累加值（当前仅支持累加值聚合方式）。 (0: SUM累加值 )， (1: AVG平均值)， (2: MAX最大值)， (3: MIN最小值)

	AttributeAggregation *int64 `json:"AttributeAggregation,omitempty" name:"AttributeAggregation"`
}

func (r *DescribeViewDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeViewDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyGroupItem struct {

	// 策略Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type CCMChartData struct {

	// 返回数据列表

	Data []*CCMChartEntry `json:"Data,omitempty" name:"Data"`
	// 查询数据总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type CCMChartEntry struct {

	// 监控图表ID

	ChartId *int64 `json:"ChartId,omitempty" name:"ChartId"`
	// 分组视图ID

	ViewId *int64 `json:"ViewId,omitempty" name:"ViewId"`
	// 监控图表名称

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
	// 监控图表类型。1 明细视图， 2 聚合视图

	ChartType *int64 `json:"ChartType,omitempty" name:"ChartType"`
	// 产品类型。1 基础监控， 2 自定义监控

	ProductType *int64 `json:"ProductType,omitempty" name:"ProductType"`
	// 指标ID

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 聚合方式。0 SUM， 1 AVG， 2 MAX， 3 MIN

	Aggregation *int64 `json:"Aggregation,omitempty" name:"Aggregation"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 创建者ID

	CreatorId *uint64 `json:"CreatorId,omitempty" name:"CreatorId"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新者ID

	UpdaterId *uint64 `json:"UpdaterId,omitempty" name:"UpdaterId"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 指标名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 指标唯一英文名称

	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`
	// 指标单位名称

	UnitName *string `json:"UnitName,omitempty" name:"UnitName"`
}

type DescribePolicyInfoByInstanceGroupList struct {

	// 策略组ID

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 策略名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否开启

	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 最后编辑者UIN

	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 是否为默认策略

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 是否能设为默认策略

	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`
	// 父策略组ID

	ParentGroupID *int64 `json:"ParentGroupID,omitempty" name:"ParentGroupID"`
	// 备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 是否为与关系策略

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
	// 项目ID

	ProjectID *int64 `json:"ProjectID,omitempty" name:"ProjectID"`
	// 告警规则

	Conditions []*Condition `json:"Conditions,omitempty" name:"Conditions"`
	// 规则模板信息

	ConditionsTemp *DescribePolicyInfoByInstanceConditionsTemp `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`
	// 实例组

	InstanceGroup *DescribePolicyInfoByInstanceInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`
	// 未屏蔽实例数

	NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`
	// 实例总数

	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type DescribeProductEventListEvents struct {

	// 事件ID

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 事件中文名

	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`
	// 事件英文名

	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`
	// 事件简称

	EventName *string `json:"EventName,omitempty" name:"EventName"`
	// 产品中文名

	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`
	// 产品英文名

	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`
	// 产品简称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 是否支持告警

	SupportAlarm *int64 `json:"SupportAlarm,omitempty" name:"SupportAlarm"`
	// 事件类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 实例对象信息

	Dimensions []*DescribeProductEventListEventsDimensions `json:"Dimensions,omitempty" name:"Dimensions"`
	// 实例对象附加信息

	AdditionMsg []*DescribeProductEventListEventsDimensions `json:"AdditionMsg,omitempty" name:"AdditionMsg"`
	// 是否配置告警

	IsAlarmConfig *int64 `json:"IsAlarmConfig,omitempty" name:"IsAlarmConfig"`
	// 策略信息

	GroupInfo []*DescribeProductEventListEventsGroupInfo `json:"GroupInfo,omitempty" name:"GroupInfo"`
}

type ModifyAlarmPolicyConditionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyConditionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyConditionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BigCustomerNotice struct {

	// 是否推送到大客户售后群 0=不推送 1=推送

	NeedPushChatGroup *int64 `json:"NeedPushChatGroup,omitempty" name:"NeedPushChatGroup"`
	// 通知开始时间 00:00:00 开始的秒数（取值范围0-86399）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 通知结束时间 00:00:00 开始的秒数（取值范围0-86399）

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type CreateMetricSetRequest struct {
	*tchttp.BaseRequest

	// 指标集配置

	MetricSet *ClmMetricSet `json:"MetricSet,omitempty" name:"MetricSet"`
}

func (r *CreateMetricSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMetricSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMLogFilterRule struct {

	// 关系 1AND 2OR

	Relation *int64 `json:"Relation,omitempty" name:"Relation"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 过滤操作符: 数字型字段支持：gt lt ge le eq ne in 字符型字段支持：eq ne in contains

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 过滤操作值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 指标集CID（只作出参，入参不填）

	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type DescribeServersRequest struct {
	*tchttp.BaseRequest

	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 服务器内网IP列表

	Ip []*string `json:"Ip,omitempty" name:"Ip"`
	// 服务器id列表

	ServerId []*uint64 `json:"ServerId,omitempty" name:"ServerId"`
	// 查询关键字。服务器ID、服务器名称或服务器IP模糊搜索

	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
	// 地域列表

	IdcId []*int64 `json:"IdcId,omitempty" name:"IdcId"`
	// 排序方式。仅支持："asc","ASC","desc","DESC"

	Order *string `json:"Order,omitempty" name:"Order"`
	// 排序字段。仅支持："ServerId","Ip"

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
}

func (r *DescribeServersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Point struct {

	// 该监控数据点生成的时间点

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 监控数据点的值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DescribeMsgPolicyInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略ID

	PolicyID *string `json:"PolicyID,omitempty" name:"PolicyID"`
}

func (r *DescribeMsgPolicyInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMsgPolicyInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {

	// 实例维度名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例维度值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeAlarmCallbackHistory struct {

	// 回调地址

	Url *string `json:"Url,omitempty" name:"Url"`
	// 回调是否验证

	ValidFlag *int64 `json:"ValidFlag,omitempty" name:"ValidFlag"`
	// 回调地址验证码

	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

type DescribeCustomAlarmListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 自定义策略Id

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 起始时间，默认为一天前

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认为当前系统时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 按发生时间排序方式，asc或desc

	OccurTimeOrder *string `json:"OccurTimeOrder,omitempty" name:"OccurTimeOrder"`
	// 模糊搜索关键词

	Filter *string `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeCustomAlarmListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCustomAlarmListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorAlarmSmsQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorAlarmSmsQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorAlarmSmsQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PCLMDescribeAlertPoliciesData struct {

	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 列表元素

	AlertPolicies []*PCLMAlertPolicy `json:"AlertPolicies,omitempty" name:"AlertPolicies"`
}

type CreateInstanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// InstanceGroupId

		InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventPolicyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 事件策略配置列表项

		Configs []*DescribeEventPolicyConfig `json:"Configs,omitempty" name:"Configs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEventPolicyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventPolicyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServerDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *AttributeValueInfoOutputData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServerDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServerDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebStorageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 弹窗信息

		WebStorageInfos *WebStorageInfos `json:"WebStorageInfos,omitempty" name:"WebStorageInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebStorageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebStorageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicySituationAlarmSituation struct {

	// 告警规则数

	AlarmRuleCount *int64 `json:"AlarmRuleCount,omitempty" name:"AlarmRuleCount"`
	// 告警触达数

	AlarmTouchCount *int64 `json:"AlarmTouchCount,omitempty" name:"AlarmTouchCount"`
}

type CreateCCMChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCCMChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCCMChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回创建的分组节点

		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例列表

	Instances []*CgrpInstance `json:"Instances,omitempty" name:"Instances"`
}

func (r *CreateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeServerInfoOutputData struct {

	// 返回上报该属性的服务器列表

	Data []*AttributeServerInfoOutput `json:"Data,omitempty" name:"Data"`
	// 返回上报该属性的服务器列表个数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AmpFrontTunnelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// HTTP请求返回结果

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AmpFrontTunnelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AmpFrontTunnelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricAnalysisDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// data

		Data *PCLMMetricAnalysisData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricAnalysisDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricAnalysisDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品信息列表

		ProductList []*ProductSimple `json:"ProductList,omitempty" name:"ProductList"`
		// 产品总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控类型，云产品监控为 MT_QCE

		MonitorTypes []*string `json:"MonitorTypes,omitempty" name:"MonitorTypes"`
		// 监控类型详情

		MonitorTypeInfos []*MonitorTypeInfo `json:"MonitorTypeInfos,omitempty" name:"MonitorTypeInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorDataByAlarmIDRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警ID

	AlarmID *string `json:"AlarmID,omitempty" name:"AlarmID"`
	// 起始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeMonitorDataByAlarmIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorDataByAlarmIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID

	StrategyId *uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 告警策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 告警策略启停状态

	IsStart *int64 `json:"IsStart,omitempty" name:"IsStart"`
	// 告警类型

	StrategyAlarmType *int64 `json:"StrategyAlarmType,omitempty" name:"StrategyAlarmType"`
	// 告警对象类型

	MixType *int64 `json:"MixType,omitempty" name:"MixType"`
	// 当MixType为服务器时，标记服务器子类型

	MixSubType *int64 `json:"MixSubType,omitempty" name:"MixSubType"`
	// 告警对象类型ID

	MixId []*int64 `json:"MixId,omitempty" name:"MixId"`
	// 告警接收类型

	ReceiverType *int64 `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 告警接收用户或用户组ID列表

	ReceiverId []*uint64 `json:"ReceiverId,omitempty" name:"ReceiverId"`
	// 告警有效起始时间

	EffectiveStartTime *string `json:"EffectiveStartTime,omitempty" name:"EffectiveStartTime"`
	// 告警有效结束事件

	EffectiveEndTime *string `json:"EffectiveEndTime,omitempty" name:"EffectiveEndTime"`
	// 告警渠道列表

	AlarmChannel []*string `json:"AlarmChannel,omitempty" name:"AlarmChannel"`
	// 告警回调url

	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 策略的告警规则列表

	Alarm []*StrategyEntryAlarm `json:"Alarm,omitempty" name:"Alarm"`
}

func (r *ModifyStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyStatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStrategyStatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrategyStatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMiniDashboardAlarmInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 七天内的告警个数

		AlarmCountInSevenDays *int64 `json:"AlarmCountInSevenDays,omitempty" name:"AlarmCountInSevenDays"`
		// 指定时间段内的告警历史

		AlarmHistory []*AlarmHistory `json:"AlarmHistory,omitempty" name:"AlarmHistory"`
		// 告警历史对应的告警指标

		AlarmMetrics []*AlarmHistoryMetric `json:"AlarmMetrics,omitempty" name:"AlarmMetrics"`
		// 实例是否有策略的相关信息

		PoliciesInfo []*MiniDashboardAlarmInfo `json:"PoliciesInfo,omitempty" name:"PoliciesInfo"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMiniDashboardAlarmInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiniDashboardAlarmInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// dashboard列表

		DashboardList []*DescribeDashboardsList `json:"DashboardList,omitempty" name:"DashboardList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMidDimensionValueListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度值列表

		DimensionValues []*string `json:"DimensionValues,omitempty" name:"DimensionValues"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMidDimensionValueListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMidDimensionValueListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetTkeDataRequest struct {
	*tchttp.BaseRequest

	// 开始时间，格式为Unix时间戳。默认值1小时前

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，格式为Unix时间戳。默认值当前时间。开始时间和结束时间的最大跨度为7天

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 命名空间。当前有tke_cluster、tke_component、tke_namespace、tke_node、tke_workload、tke_pod、tke_container

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 需要拉取的字段列表，包括维度字段和指标字段，不同Namespace的字段也各不相同

	Fields []*string `json:"Fields,omitempty" name:"Fields"`
	// 返回数据条数，默认10，最大值65536

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 查询条件，只能按维度字段进行过滤，多个条件之间是“AND”的关系。需要选择合适的维度条件，避免单次查询返回数据条数超过最大值65536条

	Conditions []*Dimension `json:"Conditions,omitempty" name:"Conditions"`
}

func (r *GetTkeDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetTkeDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecoverNotifyBatchRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 不同策略组的恢复通知方式

	GroupNotifyInfos []*ModifyRecoverNotifyBatchNotifyInfo `json:"GroupNotifyInfos,omitempty" name:"GroupNotifyInfos"`
}

func (r *ModifyRecoverNotifyBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecoverNotifyBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMsgPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建的策略id

		PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMsgPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMsgPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyInfoByInstanceRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 维度

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 偏移量

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribePolicyInfoByInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyInfoByInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMonitorDataRequest struct {
	*tchttp.BaseRequest

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 周期，

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 起始时间，如2018-09-22T19:51:23+08:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，如2018-09-22T20:51:23+08:00，默认为当前时间。 EndTime不能小于StartTime

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例对象的维度组合，格式为key-value键值对形式的集合。

	Instances []*Instance `json:"Instances,omitempty" name:"Instances"`
}

func (r *GetMonitorDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMonitorDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmMetricCustomItem struct {

	// 表达式

	Formula *string `json:"Formula,omitempty" name:"Formula"`
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标描述中文名

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 条目ID（只作出参，入参不填）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 指标集ID（只作出参，入参不填）

	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ViewIdOutput struct {

	// 视图ID

	ViewId *uint64 `json:"ViewId,omitempty" name:"ViewId"`
}

type MonitorTypeInfo struct {

	// 监控类型ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 监控类型

	Name *string `json:"Name,omitempty" name:"Name"`
	// 排列顺序

	SortId *int64 `json:"SortId,omitempty" name:"SortId"`
}

type DescribeDashboardNamespacesRequest struct {
	*tchttp.BaseRequest

	// 所属模块，默认为 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 数据源名称

	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`
	// 应用空间SpaceUUID，默认值为space_default

	SpaceUUID *string `json:"SpaceUUID,omitempty" name:"SpaceUUID"`
	// 应用空间id列表

	Ids []*string `json:"Ids,omitempty" name:"Ids"`
	// 数据源地域信息

	DataSourceRegion *string `json:"DataSourceRegion,omitempty" name:"DataSourceRegion"`
}

func (r *DescribeDashboardNamespacesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardNamespacesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNewUserRequest struct {
	*tchttp.BaseRequest

	// 固定值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警历史开始的时间戳

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 告警历史结束的时间戳

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 用户Uid

	Uid *int64 `json:"Uid,omitempty" name:"Uid"`
}

func (r *DescribeNewUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNewUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConditionsTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略组ID

		GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConditionsTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConditionsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomMetisAbnormOutputView struct {

	// 自定义监控视图ID

	ViewId *uint64 `json:"ViewId,omitempty" name:"ViewId"`
	// 自定义监控智能检测结果View

	AttrAbnormals []*CustomMetisAbnormOutputAttr `json:"AttrAbnormals,omitempty" name:"AttrAbnormals"`
}

type DashboardMetricCondition struct {

	// 数组元素为json序列化字符串

	Dimension []*string `json:"Dimension,omitempty" name:"Dimension"`
	// 域名名称，如：ap-guangzhou

	Region *string `json:"Region,omitempty" name:"Region"`
}

type DescribePolicyConditionListConfigManualPeriodNum struct {

	// 默认周期数

	Default *int64 `json:"Default,omitempty" name:"Default"`
	// 可选周期数

	Keys []*int64 `json:"Keys,omitempty" name:"Keys"`
	// 是否必须

	Need *bool `json:"Need,omitempty" name:"Need"`
}

type ModifyModuleRequest struct {
	*tchttp.BaseRequest

	// 分组Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 分组名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分组节点允许添加的实例类型

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *ModifyModuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyModuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDimensionValueResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度值搜索结果列表。

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchDimensionValueResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDimensionValueResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnsubscribeEventResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnsubscribeEventResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnsubscribeEventResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 策略数组

		Policies []*AlarmPolicy `json:"Policies,omitempty" name:"Policies"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveUnifyDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功移动的dashboard列表

		SucList []*string `json:"SucList,omitempty" name:"SucList"`
		// 移动失败的dashboard列表

		FailList *MoveFailReason `json:"FailList,omitempty" name:"FailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MoveUnifyDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveUnifyDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmAlertPolicy struct {

	// 策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标集ID(无效参数)

	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
	// 状态 1=已开启 2=未开启

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 告警接收组ID 逗号分隔

	AlertGroupId *string `json:"AlertGroupId,omitempty" name:"AlertGroupId"`
	// 告警接收渠道  "SMS", "EMAIL", "WECHAT", "CALL" 多个逗号分隔

	AlertChannel *string `json:"AlertChannel,omitempty" name:"AlertChannel"`
	// 通知时段开始时间（从00:00:00开始计算的秒数）

	NoticePeriodBegin *int64 `json:"NoticePeriodBegin,omitempty" name:"NoticePeriodBegin"`
	// 通知时段结束时间（从00:00:00开始计算的秒数）

	NoticePeriodEnd *int64 `json:"NoticePeriodEnd,omitempty" name:"NoticePeriodEnd"`
	// 回调url的scheme

	UrlScheme *string `json:"UrlScheme,omitempty" name:"UrlScheme"`
	// 回调url 不包含scheme部分

	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
	// 告警过滤条件列表

	AlertFilterRules []*ClmAlertFilterRule `json:"AlertFilterRules,omitempty" name:"AlertFilterRules"`
	// 告警触发条件列表

	AlertTriggerRules []*ClmAlertTriggerRule `json:"AlertTriggerRules,omitempty" name:"AlertTriggerRules"`
	// 告警策略ID （出参）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 创建时间 （出参）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间 （出参）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 告警接收人ID 逗号分隔

	AlertUserId *string `json:"AlertUserId,omitempty" name:"AlertUserId"`
	// 最近告警时间（出参）

	LatestAlertTime *string `json:"LatestAlertTime,omitempty" name:"LatestAlertTime"`
	// 过滤条件之间的关系 1与 2或

	FilterRelation *int64 `json:"FilterRelation,omitempty" name:"FilterRelation"`
	// 触发条件之间的关系 1与 2或

	TriggerRelation *int64 `json:"TriggerRelation,omitempty" name:"TriggerRelation"`
	// 指标集CID

	MetricSetCID *float64 `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
}

type DeleteUnifyDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功删除的dashboard列表

		SucList []*string `json:"SucList,omitempty" name:"SucList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUnifyDashboardsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUnifyDashboardsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警通知模板总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 告警通知模板列表

		Notices []*AlarmNotice `json:"Notices,omitempty" name:"Notices"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 绑定的对象实例列表

		List []*DescribeBindingPolicyObjectListInstance `json:"List,omitempty" name:"List"`
		// 绑定的对象实例总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 未屏蔽的对象实例数

		NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`
		// 绑定的实例分组信息，没有绑定实例分组则为空

		InstanceGroup *DescribeBindingPolicyObjectListInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBindingPolicyObjectListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBindingPolicyObjectListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccidentConfigAccident struct {

	// 业务id

	BusinessId *int64 `json:"BusinessId,omitempty" name:"BusinessId"`
	// 业务名称

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// 业务英文名称

	BusinessEnName *string `json:"BusinessEnName,omitempty" name:"BusinessEnName"`
	// 策略展示名称

	PolicyShowName *string `json:"PolicyShowName,omitempty" name:"PolicyShowName"`
	// 策略展示英文名称

	PolicyShowEnName *string `json:"PolicyShowEnName,omitempty" name:"PolicyShowEnName"`
	// 策略类型名称

	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`
	// 平台事件类型列表

	AccidentTypeInfo []*DescribeAccidentConfigAccidentTypeInfo `json:"AccidentTypeInfo,omitempty" name:"AccidentTypeInfo"`
}

type DeleteStrategysRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID列表

	StrategyId []*uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
}

func (r *DeleteStrategysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteStrategysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmInfoRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警Id

	AlarmId *int64 `json:"AlarmId,omitempty" name:"AlarmId"`
}

func (r *DescribeAlarmInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardDimensionKeysResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度列表

		Data []*DashboardDimensionSource `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardDimensionKeysResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardDimensionKeysResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductEventListDimensions struct {

	// 维度名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 维度值

	Value *string `json:"Value,omitempty" name:"Value"`
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

type AlarmPolicyQuota struct {

	// 账号ID

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 策略类型 ""=无策略类型

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 项目 Id -1=无项目

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 配额总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 已使用量

	Used *int64 `json:"Used,omitempty" name:"Used"`
}

type DescribeBaseMetricsDataMeaning struct {

	// 语言

	Language *string `json:"Language,omitempty" name:"Language"`
	// 含义

	Meaning *string `json:"Meaning,omitempty" name:"Meaning"`
}

type PCLMMetricCustomItem struct {

	// 表达式

	Formula *string `json:"Formula,omitempty" name:"Formula"`
	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标描述中文名

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 指标集CID（只作出参，入参不填）

	MetricSetCID *string `json:"MetricSetCID,omitempty" name:"MetricSetCID"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type ModifyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分组节点

		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMiniDashboardAlarmInfoRequest struct {
	*tchttp.BaseRequest

	// 起始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 监控类型 MT_QCE=云产品监控

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 告警namespace，用于区分云产品

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 云产品的实例信息，json格式

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 七天前告警总数的开始时间

	SevenDaysStartTime *int64 `json:"SevenDaysStartTime,omitempty" name:"SevenDaysStartTime"`
}

func (r *DescribeMiniDashboardAlarmInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMiniDashboardAlarmInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MonitorDimensionSource struct {

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标

	Metric *string `json:"Metric,omitempty" name:"Metric"`
	// 搜索值

	SearchValue *string `json:"SearchValue,omitempty" name:"SearchValue"`
}

type DescribeAgentStatusHistoryRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 查询开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 实例ID列表

	UnInstanceIDs []*string `json:"UnInstanceIDs,omitempty" name:"UnInstanceIDs"`
	// 查询终止时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAgentStatusHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentStatusHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyStrategyChannelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyStrategyChannelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyStrategyChannelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBindingPolicyObjectListDimension struct {

	// 地域id

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域简称

	Region *string `json:"Region,omitempty" name:"Region"`
	// 维度组合json字符串

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 事件维度组合json字符串

	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
}

type DashboardMetricQuery struct {

	// 查询时间段的开始时间，使用ISO格式 如：2018-04-10T10:15:58.858Z

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询时间段的结束时间，使用ISO格式 如：2018-04-10T10:15:58.858Z

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 数据源

	Datasource *string `json:"Datasource,omitempty" name:"Datasource"`
	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 查询指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 统计粒度，默认值为60，单位秒

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 分组，填写指标维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 聚合方式，根据指标类型填写聚合方式

	Aggregate *string `json:"Aggregate,omitempty" name:"Aggregate"`
	// 过滤条件

	Conditions []*DashboardMetricCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 查询标识

	SeriesId *string `json:"SeriesId,omitempty" name:"SeriesId"`
	// 数据源地域名称

	DataSourceRegion *string `json:"DataSourceRegion,omitempty" name:"DataSourceRegion"`
	// 产品配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 查询版本号，老版本不用传，新版本传"2020-10-21"

	QueryVersion *string `json:"QueryVersion,omitempty" name:"QueryVersion"`
	// 排序值的计算方式，avg|sum|min|max

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序顺序，asc|desc

	AscDesc *string `json:"AscDesc,omitempty" name:"AscDesc"`
	// 排序个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type Instance struct {

	// 实例的维度组合

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

type DescribeMonitorDataByAlarmIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标中文名

		MetricCName *string `json:"MetricCName,omitempty" name:"MetricCName"`
		// 指标单位

		MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`
		// 指标名

		MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
		// 周期

		Period *string `json:"Period,omitempty" name:"Period"`
		// 数据点

		DataPoints []*DescribeMonitorDataByAlarmIDDataPoint `json:"DataPoints,omitempty" name:"DataPoints"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorDataByAlarmIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorDataByAlarmIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTransLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 变更日志列表

		List []*TransLogItem `json:"List,omitempty" name:"List"`
		// 总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTransLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTransLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警通知模板详细信息

		Notice *AlarmNotice `json:"Notice,omitempty" name:"Notice"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyInfoRequest struct {
	*tchttp.BaseRequest

	// 模块名，这里填“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略 ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 要修改的字段 NAME=策略名称 REMARK=策略备注

	Key *string `json:"Key,omitempty" name:"Key"`
	// 修改后的值

	Value *string `json:"Value,omitempty" name:"Value"`
}

func (r *ModifyAlarmPolicyInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件夹uuid

		UUID *string `json:"UUID,omitempty" name:"UUID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashboardFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmPolicyQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额粒度 "APPID"=账户下限额 "NAMESPACE"=策略类型下限额 "PROJECT"=项目下限额

		QuotaLevel *string `json:"QuotaLevel,omitempty" name:"QuotaLevel"`
		// 以QuotaLevel为粒度，每个维度的限额

		Quotas []*AlarmPolicyQuota `json:"Quotas,omitempty" name:"Quotas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmPolicyQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPolicyQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePreferenceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePreferenceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePreferenceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendCustomAlarmMsgResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendCustomAlarmMsgResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendCustomAlarmMsgResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *AttributeIdDeleteOutput `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProjectsListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 项目列表

		Projects []*ProjectInfo `json:"Projects,omitempty" name:"Projects"`
		// 项目总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProjectsListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProjectsListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStatisticDataRequest struct {
	*tchttp.BaseRequest

	// 所属模块，固定值，为monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 命名空间，目前只支持QCE/TKE

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 维度条件，操作符支持=、in

	Conditions []*MidQueryCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 统计粒度。默认取值为300，单位为s

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 起始时间，默认为当前时间，如2020-12-08T19:51:23+08:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间，默认为当前时间，如2020-12-08T19:51:23+08:00

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 按指定维度groupBy

	GroupBys []*string `json:"GroupBys,omitempty" name:"GroupBys"`
}

func (r *DescribeStatisticDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStatisticDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductsProductMetaColumn struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 关键字, json字符串, 可能是string数组或者string

	Keys *string `json:"Keys,omitempty" name:"Keys"`
	// 英文名称

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 是否在告警中展示

	IsShowInAlarmList *bool `json:"IsShowInAlarmList,omitempty" name:"IsShowInAlarmList"`
	// 是否在产品列表中展示

	IsShowInMonitorList *bool `json:"IsShowInMonitorList,omitempty" name:"IsShowInMonitorList"`
	// 是否在选择器中展示

	IsShowInSelector *bool `json:"IsShowInSelector,omitempty" name:"IsShowInSelector"`
	// 前端渲染条件

	RenderConditions *string `json:"RenderConditions,omitempty" name:"RenderConditions"`
	// 前端渲染相关字段

	Render *string `json:"Render,omitempty" name:"Render"`
	// 前端相关字段

	NoWrap *bool `json:"NoWrap,omitempty" name:"NoWrap"`
	// 前端相关字段, 是否进行排序

	Order *bool `json:"Order,omitempty" name:"Order"`
	// 前端相关字段, 排序字段

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 是否必须

	Required *bool `json:"Required,omitempty" name:"Required"`
	// 后缀

	Postfixs *string `json:"Postfixs,omitempty" name:"Postfixs"`
}

type CreatePolicyGroupRequest struct {
	*tchttp.BaseRequest

	// 组策略名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 策略组所属视图的名称，若通过模版创建，可不传入

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 策略组所属项目Id，会进行鉴权操作

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 模版策略组Id, 通过模版创建时才需要传

	ConditionTempGroupId *int64 `json:"ConditionTempGroupId,omitempty" name:"ConditionTempGroupId"`
	// 是否屏蔽策略组，0表示不屏蔽，1表示屏蔽。不填默认为0

	IsShielded *int64 `json:"IsShielded,omitempty" name:"IsShielded"`
	// 策略组的备注信息

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 插入时间，戳格式为Unix时间戳，不填则按后台处理时间填充

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 策略组中的阈值告警规则

	Conditions []*CreatePolicyGroupCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 策略组中的时间告警规则

	EventConditions []*CreatePolicyGroupEventCondition `json:"EventConditions,omitempty" name:"EventConditions"`
	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 是否为后端调用。当且仅当值为1时，后台拉取策略模版中的规则填充入Conditions以及EventConditions字段

	BackEndCall *int64 `json:"BackEndCall,omitempty" name:"BackEndCall"`
	// 指标告警规则的且或关系，0表示或规则(满足任意规则就告警)，1表示且规则(满足所有规则才告警)

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

func (r *CreatePolicyGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePolicyGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeStrategysRequest struct {
	*tchttp.BaseRequest

	// 页码。 默认为1

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量。默认为100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 告警策略ID。

	StrategyId []*uint64 `json:"StrategyId,omitempty" name:"StrategyId"`
	// 告警策略名称

	StrategyName *string `json:"StrategyName,omitempty" name:"StrategyName"`
	// 告警对象类型。1：对象，2：视图

	MixType *uint64 `json:"MixType,omitempty" name:"MixType"`
	// 告警对象类型ID列表

	MixId []*int64 `json:"MixId,omitempty" name:"MixId"`
	// 告警接收类型

	ReceiverType *uint64 `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 告警接收人ID列表

	ReceiverUserId []*uint64 `json:"ReceiverUserId,omitempty" name:"ReceiverUserId"`
	// 告警接收用户组列表

	ReceiverGroupId []*uint64 `json:"ReceiverGroupId,omitempty" name:"ReceiverGroupId"`
	// 默认降序desc。排序方式仅支持："asc","ASC","desc","DESC"

	Order *string `json:"Order,omitempty" name:"Order"`
	// 默认创建时间CreateTime。仅支持："CreateTime","UpdateTime"

	OrderField *string `json:"OrderField,omitempty" name:"OrderField"`
	// 告警对象子类型ID。MixType=1，则1:所有服务器, 2: 指定服务器, 3:实例分组； MixType=2，则 1: 默认视图, 2: 非默认视图，指定视图ID；3：分组视图

	MixSubType *uint64 `json:"MixSubType,omitempty" name:"MixSubType"`
}

func (r *DescribeStrategysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeStrategysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NameTypeBak struct {

	// name

	Name *string `json:"Name,omitempty" name:"Name"`
}

type PCLMMetricAnalysisCustomMetric struct {

	// 复合指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 运算表达式

	Formula *string `json:"Formula,omitempty" name:"Formula"`
}

type CreateAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicyFilter struct {

	// 过滤条件类型 DIMENSION=使用 Dimensions 做过滤

	Type *string `json:"Type,omitempty" name:"Type"`
	// 裸写过滤表达式

	Expression *string `json:"Expression,omitempty" name:"Expression"`
	// AlarmPolicyDimension 二维数组序列化后的json字符串，一维数组之间互为或关系，一维数组内的元素互为与关系

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type DescribeQuestionnaireResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否已填写

		IsCompleted *bool `json:"IsCompleted,omitempty" name:"IsCompleted"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuestionnaireResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuestionnaireResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyConditionListCondition struct {

	// 策略视图名称

	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`
	// 事件告警条件

	EventMetrics []*DescribePolicyConditionListEventMetric `json:"EventMetrics,omitempty" name:"EventMetrics"`
	// 是否支持多地域

	IsSupportMultiRegion *bool `json:"IsSupportMultiRegion,omitempty" name:"IsSupportMultiRegion"`
	// 指标告警条件

	Metrics []*DescribePolicyConditionListMetric `json:"Metrics,omitempty" name:"Metrics"`
	// 策略类型名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 排序id

	SortId *int64 `json:"SortId,omitempty" name:"SortId"`
	// 是否支持默认策略

	SupportDefault *bool `json:"SupportDefault,omitempty" name:"SupportDefault"`
	// 支持该策略类型的地域列表

	SupportRegions []*string `json:"SupportRegions,omitempty" name:"SupportRegions"`
	// 过期信息

	DeprecatingInfo *string `json:"DeprecatingInfo,omitempty" name:"DeprecatingInfo"`
}

type DescribePolicyGroupCountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例绑定的告警策略数（json字符串）

		PolicyGroupCount *string `json:"PolicyGroupCount,omitempty" name:"PolicyGroupCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupCountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyGroupCountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmLogFilterRule struct {

	// 关系 1AND 2OR

	Relation *int64 `json:"Relation,omitempty" name:"Relation"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 过滤操作符: 数字型字段支持：gt lt ge le eq ne in 字符型字段支持：eq ne in contains

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 过滤操作值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 条目ID（只作出参，入参不填）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 指标集ID（只作出参，入参不填）

	MetricSetId *int64 `json:"MetricSetId,omitempty" name:"MetricSetId"`
	// 创建时间（只作出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只作出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type PCLMAlertTriggerRule struct {

	// 创建时间 （出参）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（出参）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 关联的告警策略cid

	AlertPolicyCID *string `json:"AlertPolicyCID,omitempty" name:"AlertPolicyCID"`
	// 指标名称（英文）

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标类型 1=普通指标 2=复合指标

	MetricType *int64 `json:"MetricType,omitempty" name:"MetricType"`
	// 操作符

	Operating *string `json:"Operating,omitempty" name:"Operating"`
	// 指标阈值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 持续周期个数

	ContinuousCycleCount *int64 `json:"ContinuousCycleCount,omitempty" name:"ContinuousCycleCount"`
	// 通知频率（通知间隔秒数）

	NoticeFrequencySec *int64 `json:"NoticeFrequencySec,omitempty" name:"NoticeFrequencySec"`
	// 复合指标表达式 （基础指标传空）

	MetricFormula *string `json:"MetricFormula,omitempty" name:"MetricFormula"`
}

type DescribeDashboardNamespacesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 命名空间列表

		Data []*DashboardNamespace `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardNamespacesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardNamespacesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceRequest struct {
	*tchttp.BaseRequest

	// 分组Id，1-4级Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardMetricEntry struct {

	// 命名空间值

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标类型

	MetricType *string `json:"MetricType,omitempty" name:"MetricType"`
	// 指标具有的聚合方式

	Aggregations []*string `json:"Aggregations,omitempty" name:"Aggregations"`
	// 指标展示名称

	Description *string `json:"Description,omitempty" name:"Description"`
	// 指标数据单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 最小值

	Min *float64 `json:"Min,omitempty" name:"Min"`
	// 最大值

	Max *float64 `json:"Max,omitempty" name:"Max"`
	// 维度列表

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type ModifyAlertPolicyRequest struct {
	*tchttp.BaseRequest

	// 修改后的告警策略

	AlertPolicy *ClmAlertPolicy `json:"AlertPolicy,omitempty" name:"AlertPolicy"`
	// 触发条件或筛选条件是否有修改 1有修改 2无修改

	TriggerRuleUpdatedStatus *int64 `json:"TriggerRuleUpdatedStatus,omitempty" name:"TriggerRuleUpdatedStatus"`
}

func (r *ModifyAlertPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlertPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoCallback struct {

	// 用户回调接口地址

	CallbackUrl *string `json:"CallbackUrl,omitempty" name:"CallbackUrl"`
	// 用户回调接口状态，0表示未验证，1表示已验证，2表示存在url但没有通过验证

	ValidFlag *int64 `json:"ValidFlag,omitempty" name:"ValidFlag"`
	// 用户回调接口验证码

	VerifyCode *string `json:"VerifyCode,omitempty" name:"VerifyCode"`
}

type CreateModuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 叶子分组节点

		Data []*CgrpModuleNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateModuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控面板ID

	DashboardID *int64 `json:"DashboardID,omitempty" name:"DashboardID"`
}

func (r *DeleteDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例列表

		Data []*CgrpInstanceNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstancesInInstanceGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstancesInInstanceGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstancesInInstanceGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetArgusMonitorDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetArgusMonitorDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetArgusMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNotifyBatchRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警通知方式列表

	GroupNotifyInfos []*ModifyNotifyBatchNotifyInfo `json:"GroupNotifyInfos,omitempty" name:"GroupNotifyInfos"`
}

func (r *ModifyNotifyBatchRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNotifyBatchRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupListGroup struct {

	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 是否开启

	IsOpen *bool `json:"IsOpen,omitempty" name:"IsOpen"`
	// 策略视图名称

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 最近编辑的用户uin

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 最后修改时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 策略组绑定的实例数

	UseSum *int64 `json:"UseSum,omitempty" name:"UseSum"`
	// 策略组绑定的未屏蔽实例数

	NoShieldedSum *int64 `json:"NoShieldedSum,omitempty" name:"NoShieldedSum"`
	// 是否为默认策略，0表示非默认策略，1表示默认策略

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 是否可以设置成默认策略

	CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`
	// 父策略组id

	ParentGroupId *int64 `json:"ParentGroupId,omitempty" name:"ParentGroupId"`
	// 策略组备注

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 策略组所属项目id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 阈值规则列表

	Conditions []*DescribePolicyGroupInfoCondition `json:"Conditions,omitempty" name:"Conditions"`
	// 产品事件规则列表

	EventConditions []*DescribePolicyGroupInfoEventCondition `json:"EventConditions,omitempty" name:"EventConditions"`
	// 用户接收人列表

	ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`
	// 模板策略组

	ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`
	// 策略组绑定的实例组信息

	InstanceGroup *DescribePolicyGroupListGroupInstanceGroup `json:"InstanceGroup,omitempty" name:"InstanceGroup"`
	// 且或规则标识, 0表示或规则(任意一条规则满足阈值条件就告警), 1表示且规则(所有规则都满足阈值条件才告警)

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DescribeCurrentTimestampRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeCurrentTimestampRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCurrentTimestampRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataPoint struct {

	// 实例对象维度组合

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 时间戳数组，表示那些时间点有数据，缺失的时间戳，没有数据点，可以理解为掉点了

	Timestamps []*float64 `json:"Timestamps,omitempty" name:"Timestamps"`
	// 监控值数组，该数组和Timestamps一一对应

	Values []*float64 `json:"Values,omitempty" name:"Values"`
}

type ModifyCCMGroupStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCMGroupStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCCMGroupStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnbindingPolicyTagResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UnbindingPolicyTagResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnbindingPolicyTagResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 实例分组id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 是否删除关联的告警策略组，1表示删除，其他数值表示不删除

	IsDelRelatedPolicy *int64 `json:"IsDelRelatedPolicy,omitempty" name:"IsDelRelatedPolicy"`
}

func (r *DeleteInstanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecoverNotifyBatchResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecoverNotifyBatchResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecoverNotifyBatchResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmDescribeMetricSetsData struct {

	// 指标集列表数据

	MetricSets []*ClmMetricSet `json:"MetricSets,omitempty" name:"MetricSets"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type ClmModifyAlertPolicyStatusParam struct {

	// 告警策略ID

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 状态：1开启 2关闭

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeDashboardsList struct {

	// DashboardID

	DashboardID *string `json:"DashboardID,omitempty" name:"DashboardID"`
	// 描述名称

	DescName *string `json:"DescName,omitempty" name:"DescName"`
	// 元数据

	Meta *string `json:"Meta,omitempty" name:"Meta"`
}

type CopyConditionsTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CopyConditionsTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyConditionsTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlertPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略ID

		Data *int64 `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlertPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlertPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略组名称

		GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
		// 策略组所属的项目id

		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
		// 是否为默认策略，0表示非默认策略，1表示默认策略

		IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
		// 策略类型

		ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
		// 策略说明

		Remark *string `json:"Remark,omitempty" name:"Remark"`
		// 策略类型名称

		ShowName *string `json:"ShowName,omitempty" name:"ShowName"`
		// 最近编辑的用户uin

		LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
		// 最近编辑时间

		UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
		// 该策略支持的地域

		Region []*string `json:"Region,omitempty" name:"Region"`
		// 策略类型的维度列表

		DimensionGroup []*string `json:"DimensionGroup,omitempty" name:"DimensionGroup"`
		// 阈值规则列表

		ConditionsConfig []*DescribePolicyGroupInfoCondition `json:"ConditionsConfig,omitempty" name:"ConditionsConfig"`
		// 产品事件规则列表

		EventConfig []*DescribePolicyGroupInfoEventCondition `json:"EventConfig,omitempty" name:"EventConfig"`
		// 用户接收人列表

		ReceiverInfos []*DescribePolicyGroupInfoReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`
		// 用户回调信息

		Callback *DescribePolicyGroupInfoCallback `json:"Callback,omitempty" name:"Callback"`
		// 模板策略组

		ConditionsTemp *DescribePolicyGroupInfoConditionTpl `json:"ConditionsTemp,omitempty" name:"ConditionsTemp"`
		// 是否可以设置成默认策略

		CanSetDefault *bool `json:"CanSetDefault,omitempty" name:"CanSetDefault"`
		// 是否且规则

		IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePolicyGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePolicyGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductsProductMetaAlarm struct {

	// 指标告警维度

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 事件告警维度

	EventDimensions *string `json:"EventDimensions,omitempty" name:"EventDimensions"`
	// 告警策略视图名称

	PolicyViewName *string `json:"PolicyViewName,omitempty" name:"PolicyViewName"`
}

type ProjectInfo struct {

	// AppId

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 创建人Uin

	CreatorUin *int64 `json:"CreatorUin,omitempty" name:"CreatorUin"`
	// 项目名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 项目信息

	Info *string `json:"Info,omitempty" name:"Info"`
	// 是否为默认

	IsDefault *int64 `json:"IsDefault,omitempty" name:"IsDefault"`
	// 项目Id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 主账号Uin

	OwnerUin *int64 `json:"OwnerUin,omitempty" name:"OwnerUin"`
	// 源AppId

	SrcAppId *int64 `json:"SrcAppId,omitempty" name:"SrcAppId"`
	// 源平台

	SrcPlat *string `json:"SrcPlat,omitempty" name:"SrcPlat"`
	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type AlarmBindingInstance struct {

	// 实例UUID

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 实例维度组合信息

	Information *string `json:"Information,omitempty" name:"Information"`
}

type PolicyGroupReceiverInfo struct {

	// 有效时段结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 是否需要发送通知

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

	RecoverNotify []*string `json:"RecoverNotify,omitempty" name:"RecoverNotify"`
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

type DescribeAccidentEventListAlarms struct {

	// 事件分类

	BusinessTypeDesc *string `json:"BusinessTypeDesc,omitempty" name:"BusinessTypeDesc"`
	// 事件类型

	AccidentTypeDesc *string `json:"AccidentTypeDesc,omitempty" name:"AccidentTypeDesc"`
	// 事件分类的ID，1表示服务问题，2表示其他订阅

	BusinessID *int64 `json:"BusinessID,omitempty" name:"BusinessID"`
	// 事件状态的ID，0表示已恢复，1表示未恢复

	EventStatus *int64 `json:"EventStatus,omitempty" name:"EventStatus"`
	// 影响的对象

	AffectResource *string `json:"AffectResource,omitempty" name:"AffectResource"`
	// 事件的地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 事件发生的时间

	OccurTime *string `json:"OccurTime,omitempty" name:"OccurTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribePolicyGroupInfoEventCondition struct {

	// 事件id

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 事件告警规则id

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 事件名称

	EventShowName *string `json:"EventShowName,omitempty" name:"EventShowName"`
	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次

	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
	// 告警发送收敛类型。0连续告警，1指数告警

	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
}

type DeleteGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分组节点

		Data []*CgrpGroupNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyModuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 叶子分组节点

		Data []*CgrpModuleNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyModuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyModuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CCMDashboardData struct {

	// 返回数据列表

	Data []*CCMDashboardEntry `json:"Data,omitempty" name:"Data"`
	// 查询数据总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type PCLMMetricAnalysisData struct {

	// 对比数据

	Compare []*PCLMMetricAnalysisPoint `json:"Compare,omitempty" name:"Compare"`
	// 当前数据

	Current []*PCLMMetricAnalysisPoint `json:"Current,omitempty" name:"Current"`
	// 时间粒度 （秒）

	Granularity *int64 `json:"Granularity,omitempty" name:"Granularity"`
}

type ServerAttributeInfo struct {

	// 指标ID

	AttributeId *int64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 指标名称

	AttributeName *string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 指标唯一英文名称

	AttributeUniqueName *string `json:"AttributeUniqueName,omitempty" name:"AttributeUniqueName"`
}

type DescribeAlarmObjectQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配额粒度 "APPID"=账户下限额 "NAMESPACE"=策略类型下限额 "PROJECT"=项目下限额

		QuotaLevel *string `json:"QuotaLevel,omitempty" name:"QuotaLevel"`
		// 以QuotaLevel为粒度，每个维度的限额

		Quotas []*AlarmPolicyQuota `json:"Quotas,omitempty" name:"Quotas"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmObjectQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmObjectQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *ServiceInfo `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeViewDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *AttributeValueInfoOutputData `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeViewDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeViewDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnBindingPolicyObjectRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略id

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 待删除对象实例的唯一id列表

	UniqueId []*string `json:"UniqueId,omitempty" name:"UniqueId"`
	// 实例分组id, 如果按实例分组删除的话UniqueId参数是无效的

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
}

func (r *UnBindingPolicyObjectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnBindingPolicyObjectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardViewRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控面板视图ID

	ViewID *int64 `json:"ViewID,omitempty" name:"ViewID"`
}

func (r *DeleteDashboardViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMetricSetRequest struct {
	*tchttp.BaseRequest

	// 指标集ID

	MetricSetID *int64 `json:"MetricSetID,omitempty" name:"MetricSetID"`
}

func (r *DeleteMetricSetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetricSetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicyCondition struct {

	// 指标触发与或条件，0=或，1=与

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
	// 告警触发条件列表

	Rules []*AlarmPolicyRule `json:"Rules,omitempty" name:"Rules"`
}

type DescribeUserInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 语种

		MsgLang *string `json:"MsgLang,omitempty" name:"MsgLang"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuestionnaireRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeQuestionnaireRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuestionnaireRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusDimensionOut struct {

	// Id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// namespace的Id

	NamespaceId *uint64 `json:"NamespaceId,omitempty" name:"NamespaceId"`
	// 英文名

	EnName *string `json:"EnName,omitempty" name:"EnName"`
	// 中文名

	CnName *string `json:"CnName,omitempty" name:"CnName"`
	// 创建人

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 创建的时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新人

	UpdateUin *string `json:"UpdateUin,omitempty" name:"UpdateUin"`
	// 更新的时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
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

type DescribeAlarmObjectQuotaRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 监控类型过滤 "MT_QCE"=云产品监控 "MT_CUSTOM"=自定义监控

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 触发条件类型 STATIC=静态阈值 DYNAMIC=动态阈值

	RuleType *string `json:"RuleType,omitempty" name:"RuleType"`
	// 告警策略uuid

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
}

func (r *DescribeAlarmObjectQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmObjectQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCCMInstanceDatasRequest struct {
	*tchttp.BaseRequest

	// 实例名称列表

	InstanceName []*string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 指标ID

	AttributeId *uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 开始时间。例如："2019-09-11 00:00:00"

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。例如："2019-09-11 00:00:00"

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 实例类型。输入："cvm_device"或"SCF"

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
}

func (r *DescribeCCMInstanceDatasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMInstanceDatasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogTopicsRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
}

func (r *DescribeLogTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 策略ID

		PolicyID *string `json:"PolicyID,omitempty" name:"PolicyID"`
		// 策略名

		PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
		// 接收组ID

		ReceiverGroupIDs []*int64 `json:"ReceiverGroupIDs,omitempty" name:"ReceiverGroupIDs"`
		// 通知方式

		NotifyWays []*string `json:"NotifyWays,omitempty" name:"NotifyWays"`
		// 是否为默认

		IsDefault *bool `json:"IsDefault,omitempty" name:"IsDefault"`
		// 告警总数

		AlarmCount *int64 `json:"AlarmCount,omitempty" name:"AlarmCount"`
		// 语音配置

		VoiceConfig *DescribeMsgPolicyListVoiceConfig `json:"VoiceConfig,omitempty" name:"VoiceConfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMsgPolicyInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMsgPolicyInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeProductHealthStatusListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各产品的健康状态列表

		List []*ProductHealthStatus `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProductHealthStatusListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeProductHealthStatusListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PutMonitorDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutMonitorDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PutMonitorDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCvmAgentStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例的状态列表

		List []*CvmAgentStatus `json:"List,omitempty" name:"List"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCvmAgentStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCvmAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecommendedTemplateRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

func (r *DescribeRecommendedTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecommendedTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSetsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeLogSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyStatusRequest struct {
	*tchttp.BaseRequest

	// 模块名，固定值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 告警策略 ID

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 启停状态 0=停用 1=启用

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyAlarmPolicyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HistoryAlarmInfo struct {

	// 告警Id

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 策略组Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 告警状态

	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
	// 告警类型，0（指标告警），2（产品事件告警），3（平台事件告警）

	AlarmType *int64 `json:"AlarmType,omitempty" name:"AlarmType"`
	// 告警内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 当前值

	CurrentValue *string `json:"CurrentValue,omitempty" name:"CurrentValue"`
	// 首次触发时间

	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`
	// 最后触发时间

	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`
	// 持续时间

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// 对象Id

	ObjId *string `json:"ObjId,omitempty" name:"ObjId"`
	// 对象名

	ObjName *string `json:"ObjName,omitempty" name:"ObjName"`
	// 项目Id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 状态，其中0表示未恢复，1表示已恢复，4表示已失效，2/3/5表示数据不足

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
}

type DescribeCCMGroupViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *CCMGroupViewEntry `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCCMGroupViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCCMGroupViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAttributesRequest struct {
	*tchttp.BaseRequest

	// 页码

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 每页数量

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 指标id列表

	AttributeId []*uint64 `json:"AttributeId,omitempty" name:"AttributeId"`
	// 指标名称列表

	AttributeName []*string `json:"AttributeName,omitempty" name:"AttributeName"`
	// 指标类型名称列表

	AttributeType []*string `json:"AttributeType,omitempty" name:"AttributeType"`
	// 负责人名称列表

	OwnerName []*string `json:"OwnerName,omitempty" name:"OwnerName"`
	// 查询关键字

	SearchKeyword *string `json:"SearchKeyword,omitempty" name:"SearchKeyword"`
}

func (r *DescribeAttributesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAttributesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UnsubscribeEventRequest struct {
	*tchttp.BaseRequest

	// 固定值"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 业务id

	BusinessId *int64 `json:"BusinessId,omitempty" name:"BusinessId"`
	// 事件id

	AccidentId *int64 `json:"AccidentId,omitempty" name:"AccidentId"`
}

func (r *UnsubscribeEventRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UnsubscribeEventRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConditionsTemplateRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 模板策略组ID

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 是否删除关联策略

	IsDeleteRelatedPolicy *int64 `json:"IsDeleteRelatedPolicy,omitempty" name:"IsDeleteRelatedPolicy"`
	// 所属用户UIN

	OwnerUIN *string `json:"OwnerUIN,omitempty" name:"OwnerUIN"`
}

func (r *DeleteConditionsTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConditionsTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceGroupListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例分组列表

		InstanceGroupList []*DescribeInstanceGroupListInstanceGroup `json:"InstanceGroupList,omitempty" name:"InstanceGroupList"`
		// 实例分组总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceGroupListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceGroupListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMenuAuthorityRequest struct {
	*tchttp.BaseRequest

	// 菜单名

	Menu *string `json:"Menu,omitempty" name:"Menu"`
}

func (r *DescribeMenuAuthorityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMenuAuthorityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMenuAuthorityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否可使用当前页面

		IsAccessible *bool `json:"IsAccessible,omitempty" name:"IsAccessible"`
		// 当前用户版本状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMenuAuthorityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMenuAuthorityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardFolderRequest struct {
	*tchttp.BaseRequest

	// 文件夹详情

	Data *string `json:"Data,omitempty" name:"Data"`
}

func (r *CreateDashboardFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscribeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 事件告警列表

		Events []*SubscribeInfo `json:"Events,omitempty" name:"Events"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscribeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSubscribeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件夹版本

		Version *int64 `json:"Version,omitempty" name:"Version"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMsgPolicyListVoiceConfig struct {

	// 呼叫次数

	RoundNumber *int64 `json:"RoundNumber,omitempty" name:"RoundNumber"`
	// UIN列表

	UIDList []*int64 `json:"UIDList,omitempty" name:"UIDList"`
	// 每轮间隔

	RoundInterval *int64 `json:"RoundInterval,omitempty" name:"RoundInterval"`
	// 没人间隔

	PersonInterval *int64 `json:"PersonInterval,omitempty" name:"PersonInterval"`
	// 是否需要发送提醒

	NeedSendNotice *int64 `json:"NeedSendNotice,omitempty" name:"NeedSendNotice"`
}

type DescribePolicyGroupInfoConditionTpl struct {

	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 策略类型

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 策略组说明

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 最后编辑的用户uin

	LastEditUin *string `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
	// 是否且规则

	IsUnionRule *int64 `json:"IsUnionRule,omitempty" name:"IsUnionRule"`
}

type DeleteAttributesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *AttributeIdDeleteOutput `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAttributesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAttributesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmBindingInstanceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果数量

		Count *int64 `json:"Count,omitempty" name:"Count"`
		// 实例列表

		InstanceList []*AlarmBindingInstance `json:"InstanceList,omitempty" name:"InstanceList"`
		// 本次查询的限额

		Limit *int64 `json:"Limit,omitempty" name:"Limit"`
		// 本次查询的偏移量

		Offset *int64 `json:"Offset,omitempty" name:"Offset"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmBindingInstanceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmBindingInstanceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricDatum struct {

	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标的值

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type DescribeDashboardMetricsRequest struct {
	*tchttp.BaseRequest

	// 所属模块，默认为 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 数据源名称

	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`
	// 命名空间值

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 应用空间SpaceUUID，默认值为space_default

	SpaceUUID *string `json:"SpaceUUID,omitempty" name:"SpaceUUID"`
	// 数据源地域信息

	DataSourceRegion *string `json:"DataSourceRegion,omitempty" name:"DataSourceRegion"`
}

func (r *DescribeDashboardMetricsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardMetricsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataForMiniProgramRequest struct {
	*tchttp.BaseRequest

	// 告警历史AlertId

	AlertId *string `json:"AlertId,omitempty" name:"AlertId"`
}

func (r *DescribeDataForMiniProgramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataForMiniProgramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplateTagsRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *DescribeTemplateTagsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplateTagsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCurrentTimestampResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 服务器当前unix时间戳(秒数)

		Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCurrentTimestampResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCurrentTimestampResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUnifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 版本

		Version *int64 `json:"Version,omitempty" name:"Version"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUnifyDashboardResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUnifyDashboardResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppFlowConverterResponseDataPoint struct {

	// 数据值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DescribeGroupRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果数据json

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ArgusMetricIn struct {

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
}

type DescribeAlarmBindingInstanceListRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 绑定告警策略的AppID

	AppID *int64 `json:"AppID,omitempty" name:"AppID"`
	// 告警策略绑定的视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 查询限制，取值0 - 100，默认100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAlarmBindingInstanceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmBindingInstanceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type StarUnifyDashboardRequest struct {
	*tchttp.BaseRequest

	// dashboard uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
}

func (r *StarUnifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StarUnifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyPolicyGroupCondition struct {

	// 规则id，不填表示新增，填写了ruleId表示在已存在的规则基础上进行修改

	RuleId *int64 `json:"RuleId,omitempty" name:"RuleId"`
	// 指标id

	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`
	// 比较类型，1表示大于，2表示大于等于，3表示小于，4表示小于等于，5表示相等，6表示不相等

	CalcType *int64 `json:"CalcType,omitempty" name:"CalcType"`
	// 检测阈值

	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`
	// 检测指标的数据周期

	CalcPeriod *int64 `json:"CalcPeriod,omitempty" name:"CalcPeriod"`
	// 持续周期个数

	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`
	// 告警发送收敛类型。0连续告警，1指数告警

	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 告警发送周期单位秒。<0 不触发, 0 只触发一次, >0 每隔triggerTime秒触发一次

	AlarmNotifyPeriod *int64 `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
}

type DescribeUnifyDashboardsRequest struct {
	*tchttp.BaseRequest

	// 是否收藏

	IsStarred *string `json:"IsStarred,omitempty" name:"IsStarred"`
	// 关键字搜索

	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`
	// 是否文件夹

	IsFolder *string `json:"IsFolder,omitempty" name:"IsFolder"`
	// 小程序：applet

	Caller *string `json:"Caller,omitempty" name:"Caller"`
}

func (r *DescribeUnifyDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUnifyDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMGroupStrategyRequest struct {
	*tchttp.BaseRequest

	// 实例类型。仅支持："cvm_device","SCF"

	TypeId *string `json:"TypeId,omitempty" name:"TypeId"`
	// 实例名称列表。

	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`
	// 实例分组ID。

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ModifyCCMGroupStrategyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCCMGroupStrategyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetDefaultPolicyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetDefaultPolicyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetDefaultPolicyGroupResponse) FromJsonString(s string) error {
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

type DescribeDashboardDimensionKeysRequest struct {
	*tchttp.BaseRequest

	// 应用空间SpaceUUID，默认值为space_default

	SpaceUUID *string `json:"SpaceUUID,omitempty" name:"SpaceUUID"`
	// 数据源名称

	DataSource *string `json:"DataSource,omitempty" name:"DataSource"`
	// 包含：Namespace（命名空间）、Measurement（指标）、维度名搜索值

	DimensionSource []*MonitorDimensionSource `json:"DimensionSource,omitempty" name:"DimensionSource"`
	// 取值 monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 数据源地域信息

	DataSourceRegion *string `json:"DataSourceRegion,omitempty" name:"DataSourceRegion"`
}

func (r *DescribeDashboardDimensionKeysRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardDimensionKeysRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCCMChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data []*int64 `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCCMChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCCMChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePolicyGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建成功的策略组Id

		GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePolicyGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePolicyGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdcServerCountRequest struct {
	*tchttp.BaseRequest

	// 地域ID数组列表

	IdcId []*int64 `json:"IdcId,omitempty" name:"IdcId"`
}

func (r *DescribeIdcServerCountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdcServerCountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebStorageRequest struct {
	*tchttp.BaseRequest

	// 弹窗信息

	WebStorageInfo []*string `json:"WebStorageInfo,omitempty" name:"WebStorageInfo"`
}

func (r *DescribeWebStorageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebStorageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VerifyAlarmCallbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 验证结果

		VerifyStatus *string `json:"VerifyStatus,omitempty" name:"VerifyStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyAlarmCallbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyAlarmCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AttributeTimestampValueOutput struct {

	// 时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 属性值

	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type DescribeContactListRequest struct {
	*tchttp.BaseRequest

	// 固定值, "monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 偏移量, 从1开始计数

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 限制数，每页返回的数量，取值1~100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeContactListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeContactListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBaseMetricsDataPeriod struct {

	// 周期

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 统计方式

	StatType []*string `json:"StatType,omitempty" name:"StatType"`
}

type IdcInfoData struct {

	// 返回数据

	Data []*IdcInfo `json:"Data,omitempty" name:"Data"`
	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type InstanceGroup struct {

	// 实例组ID

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 实例组名

	InstanceGroupName *string `json:"InstanceGroupName,omitempty" name:"InstanceGroupName"`
}

type DescribeLogTopicIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果数据json

		Data *string `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogTopicIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogTopicIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceGroupRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 实例组Id

	InstanceGroupId *int64 `json:"InstanceGroupId,omitempty" name:"InstanceGroupId"`
	// 实例组名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 实例列表

	InstanceList []*BindingPolicyObjectDimension `json:"InstanceList,omitempty" name:"InstanceList"`
}

func (r *ModifyInstanceGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePolicyInfoByInstanceInstanceGroup struct {

	// 实例组ID

	InstanceGroupID *int64 `json:"InstanceGroupID,omitempty" name:"InstanceGroupID"`
	// 视图

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 最后编辑者UIN

	LastEditUin *int64 `json:"LastEditUin,omitempty" name:"LastEditUin"`
	// 实例组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 实例数

	InstanceSum *int64 `json:"InstanceSum,omitempty" name:"InstanceSum"`
	// 更新时间

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	InsertTime *int64 `json:"InsertTime,omitempty" name:"InsertTime"`
}

type CreateAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略名称，不超过20字符

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 备注，不超过100字符，仅支持中英文、数字、下划线、-

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 监控类型 MT_QCE=云产品监控

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 是否启用 0=停用 1=启用，可不传 默认为1

	Enable *int64 `json:"Enable,omitempty" name:"Enable"`
	// 项目 Id -1=无项目 0=默认项目，可不传 默认为-1

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 告警策略类型，由 DescribeAllNamespaces 获得，例如 cvm_device

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 裸写包含触发规则、过滤条件等内容的告警规则

	PolicyExpression *string `json:"PolicyExpression,omitempty" name:"PolicyExpression"`
	// 触发条件模板 Id ，可不传

	ConditionTemplateId *int64 `json:"ConditionTemplateId,omitempty" name:"ConditionTemplateId"`
	// 指标触发条件

	Condition *AlarmPolicyCondition `json:"Condition,omitempty" name:"Condition"`
	// 事件触发条件

	EventCondition *AlarmPolicyEventCondition `json:"EventCondition,omitempty" name:"EventCondition"`
	// 通知规则 Id 列表，由 DescribeAlarmNotices 获得

	NoticeIds []*string `json:"NoticeIds,omitempty" name:"NoticeIds"`
	// 触发任务列表

	TriggerTasks []*AlarmPolicyTriggerTask `json:"TriggerTasks,omitempty" name:"TriggerTasks"`
}

func (r *CreateAlarmPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeEventPolicyConfigRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 视图名

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 产品列表

	Products []*string `json:"Products,omitempty" name:"Products"`
}

func (r *DescribeEventPolicyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeEventPolicyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCCMChartsRequest struct {
	*tchttp.BaseRequest

	// 监控图表ID列表。

	ChartId []*int64 `json:"ChartId,omitempty" name:"ChartId"`
}

func (r *DeleteCCMChartsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCCMChartsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MoveUnifyDashboardsRequest struct {
	*tchttp.BaseRequest

	// dashboard列表

	DashboardUUIDList []*string `json:"DashboardUUIDList,omitempty" name:"DashboardUUIDList"`
	// 文件夹uuid

	FolderUUID *string `json:"FolderUUID,omitempty" name:"FolderUUID"`
	// 是否覆盖

	Overwrite *string `json:"Overwrite,omitempty" name:"Overwrite"`
}

func (r *MoveUnifyDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MoveUnifyDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmDescribeAlertPoliciesData struct {

	// 总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 列表元素

	AlertPolicies []*ClmAlertPolicy `json:"AlertPolicies,omitempty" name:"AlertPolicies"`
}

type DescribeProductEventListEventsGroupInfo struct {

	// 策略ID

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略名

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
}

type CreateGroupRequest struct {
	*tchttp.BaseRequest

	// 父分组ID

	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`
	// 分组名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClmStatisticFilterRule struct {

	// 条目ID（只做出参，入参不填）

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 关联指标条目ID（只做出参，入参不填）

	MetricItemId *int64 `json:"MetricItemId,omitempty" name:"MetricItemId"`
	// 创建时间（只做出参，入参不填）

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（只做出参，入参不填）

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 关系 1AND 2OR

	Relation *int64 `json:"Relation,omitempty" name:"Relation"`
	// 字段名

	Field *string `json:"Field,omitempty" name:"Field"`
	// 过滤操作符: 数字型字段支持：gt lt ge le eq ne in 字符型字段支持：eq ne in contains

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 过滤操作值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeBindingPolicyObjectListRequest struct {
	*tchttp.BaseRequest

	// 固定值，为"monitor"

	Module *string `json:"Module,omitempty" name:"Module"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略id

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 分页参数，每页返回的数量，取值1~100，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页参数，页偏移量，从0开始计数，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 筛选对象的维度信息

	Dimensions []*DescribeBindingPolicyObjectListDimension `json:"Dimensions,omitempty" name:"Dimensions"`
}

func (r *DescribeBindingPolicyObjectListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBindingPolicyObjectListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CopyUnifyDashboardRequest struct {
	*tchttp.BaseRequest

	// 原dashboard uuid

	UUID *string `json:"UUID,omitempty" name:"UUID"`
	// 新dashboard 标题

	Title *string `json:"Title,omitempty" name:"Title"`
	// 文件夹uuid

	FolderUUID *string `json:"FolderUUID,omitempty" name:"FolderUUID"`
}

func (r *CopyUnifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyUnifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateStrategyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回数据

		Data *StrategyIdInfo `json:"Data,omitempty" name:"Data"`
		// 请求执行时间

		ExecTime *string `json:"ExecTime,omitempty" name:"ExecTime"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateStrategyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateStrategyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IdcData struct {

	// idc id,在腾讯云上代表zoneId

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// idc id,在腾讯云上代表zone名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ModifyMetricSetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标集详情

		Data *ClmMetricSet `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMetricSetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMetricSetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmReceiversRequest struct {
	*tchttp.BaseRequest

	// 需要修改接收人的策略组Id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 新接收人信息, 没有填写则删除所有接收人

	ReceiverInfos []*ReceiverInfo `json:"ReceiverInfos,omitempty" name:"ReceiverInfos"`
	// 必填。固定为“monitor”

	Module *string `json:"Module,omitempty" name:"Module"`
}

func (r *ModifyAlarmReceiversRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmReceiversRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmHistory struct {

	// 告警历史Id

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
	// 监控类型

	MonitorType *string `json:"MonitorType,omitempty" name:"MonitorType"`
	// 策略类型

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 告警对象

	AlarmObject *string `json:"AlarmObject,omitempty" name:"AlarmObject"`
	// 告警内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 时间戳，首次出现时间

	FirstOccurTime *int64 `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`
	// 时间戳，最后出现时间

	LastOccurTime *int64 `json:"LastOccurTime,omitempty" name:"LastOccurTime"`
	// 告警状态，ALARM=未恢复 OK=已恢复 NO_CONF=已失效 NO_DATA=数据不足

	AlarmStatus *string `json:"AlarmStatus,omitempty" name:"AlarmStatus"`
	// 告警策略 Id

	PolicyId *string `json:"PolicyId,omitempty" name:"PolicyId"`
	// 策略名称

	PolicyName *string `json:"PolicyName,omitempty" name:"PolicyName"`
	// 基础产品告警的告警对象所属网络

	VPC *string `json:"VPC,omitempty" name:"VPC"`
	// 项目 Id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名字

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 告警对象所属实例组

	InstanceGroup []*InstanceGroups `json:"InstanceGroup,omitempty" name:"InstanceGroup"`
	// 接收人列表

	ReceiverUids []*int64 `json:"ReceiverUids,omitempty" name:"ReceiverUids"`
	// 接收组列表

	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`
	// 告警渠道列表 SMS=短信 EMAIL=邮件 CALL=电话 WECHAT=微信

	NoticeWays []*string `json:"NoticeWays,omitempty" name:"NoticeWays"`
	// 兼容告警1.0策略组 Id

	OriginId *string `json:"OriginId,omitempty" name:"OriginId"`
	// 告警类型

	AlarmType *string `json:"AlarmType,omitempty" name:"AlarmType"`
	// 事件Id

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 策略是否存在 0=不存在 1=存在

	PolicyExists *int64 `json:"PolicyExists,omitempty" name:"PolicyExists"`
	// 额外信息，json格式的字符串

	ExtraTags *string `json:"ExtraTags,omitempty" name:"ExtraTags"`
	// 指标信息

	MetricsInfo []*AlarmHistoryMetric `json:"MetricsInfo,omitempty" name:"MetricsInfo"`
	// Dimensions

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type ModifyConditionsTemplateRequestCondition struct {

	// 告警通知周期

	AlarmNotifyPeriod *string `json:"AlarmNotifyPeriod,omitempty" name:"AlarmNotifyPeriod"`
	// 告警通知方式

	AlarmNotifyType *int64 `json:"AlarmNotifyType,omitempty" name:"AlarmNotifyType"`
	// 统计周期

	CalcPeriod *string `json:"CalcPeriod,omitempty" name:"CalcPeriod"`
	// 统计方式

	CalcType *string `json:"CalcType,omitempty" name:"CalcType"`
	// 统计值

	CalcValue *string `json:"CalcValue,omitempty" name:"CalcValue"`
	// 持续周期

	ContinuePeriod *string `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`
	// 指标ID

	MetricId *int64 `json:"MetricId,omitempty" name:"MetricId"`
	// 规则ID

	RuleID *int64 `json:"RuleID,omitempty" name:"RuleID"`
}

type CopyConditionsTemplateRequest struct {
	*tchttp.BaseRequest

	// 固定值，monitor

	Module *string `json:"Module,omitempty" name:"Module"`
	// 模板策略组ID

	GroupID *int64 `json:"GroupID,omitempty" name:"GroupID"`
	// 所属UIN

	OwnerUIN *string `json:"OwnerUIN,omitempty" name:"OwnerUIN"`
}

func (r *CopyConditionsTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CopyConditionsTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateUnifyDashboardRequest struct {
	*tchttp.BaseRequest

	// dashboard详细信息

	Data *string `json:"Data,omitempty" name:"Data"`
	// "1"表示压缩

	Compress *string `json:"Compress,omitempty" name:"Compress"`
	// 文件夹UUID

	FolderUUID *string `json:"FolderUUID,omitempty" name:"FolderUUID"`
}

func (r *CreateUnifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateUnifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyTag struct {

	// 标签Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分组列表

		Data []*CgrpInstanceNode `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicyEventCondition struct {

	// 告警触发条件列表

	Rules []*AlarmPolicyRule `json:"Rules,omitempty" name:"Rules"`
}

type DescribeAbnormalObjectsEventObject struct {

	// 告警状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 告警内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 维度组合，json字符串

	Dimensions *string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 告警状态，“abnormal”表示告警未恢复，“warning”表示告警已恢复

	AbnormalStatus *string `json:"AbnormalStatus,omitempty" name:"AbnormalStatus"`
	// 首次发生时间

	FirstOccurTime *string `json:"FirstOccurTime,omitempty" name:"FirstOccurTime"`
	// 最后发生时间

	LastOccurTime *string `json:"LastOccurTime,omitempty" name:"LastOccurTime"`
	// 维度组合md5

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
	// 策略组id

	GroupId *int64 `json:"GroupId,omitempty" name:"GroupId"`
	// 策略组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 项目id

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 产品中文名称

	ProductCName *string `json:"ProductCName,omitempty" name:"ProductCName"`
	// 产品英文名称

	ProductEName *string `json:"ProductEName,omitempty" name:"ProductEName"`
	// 事件英文名称

	EventEName *string `json:"EventEName,omitempty" name:"EventEName"`
	// 事件中文名称

	EventCName *string `json:"EventCName,omitempty" name:"EventCName"`
	// 产品事件id

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 产品维度组合md5

	EventUniqueId *string `json:"EventUniqueId,omitempty" name:"EventUniqueId"`
}

type DescribePolicyConditionListConfigManualPeriod struct {

	// 默认周期，单位：秒

	Default *int64 `json:"Default,omitempty" name:"Default"`
	// 可选周期，单位：秒

	Keys []*int64 `json:"Keys,omitempty" name:"Keys"`
	// 是否必须

	Need *bool `json:"Need,omitempty" name:"Need"`
}

type BindingPolicyObjectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *BindingPolicyObjectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *BindingPolicyObjectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMonitorProductsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云产品列表，json 字符串

		ProductList *string `json:"ProductList,omitempty" name:"ProductList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMonitorProductsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMonitorProductsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

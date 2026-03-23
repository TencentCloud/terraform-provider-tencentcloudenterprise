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
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ApmApplicationConfigView struct {

	// 是否开启日志&nbsp;0&nbsp;关&nbsp;1&nbsp;开

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// 组件列表开关（已废弃）

	InstrumentList []*Instrument `json:"InstrumentList,omitempty" name:"InstrumentList"`
	// 链路压缩开关（已废弃）

	TraceSquash *bool `json:"TraceSquash,omitempty" name:"TraceSquash"`
	// 错误类型过滤

	ExceptionFilter *string `json:"ExceptionFilter,omitempty" name:"ExceptionFilter"`
	// URL&nbsp;收敛规则正则

	UrlConvergence *string `json:"UrlConvergence,omitempty" name:"UrlConvergence"`
	// 业务系统&nbsp;ID

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// HTTP&nbsp;状态码过滤

	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitempty" name:"ErrorCodeFilter"`
	// URL&nbsp;收敛阈值

	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitempty" name:"UrlConvergenceThreshold"`
	// URL&nbsp;排除规则正则

	UrlExclude *string `json:"UrlExclude,omitempty" name:"UrlExclude"`
	// 日志源

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 日志主题

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// 接口过滤

	OperationNameFilter *string `json:"OperationNameFilter,omitempty" name:"OperationNameFilter"`
	// 应用诊断开关（已废弃）

	EventEnable *bool `json:"EventEnable,omitempty" name:"EventEnable"`
	// 日志集&nbsp;

	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`
	// 方法栈快照开关&nbsp;true&nbsp;开启&nbsp;false&nbsp;关闭

	SnapshotEnable *bool `json:"SnapshotEnable,omitempty" name:"SnapshotEnable"`
	// 慢调用监听触发阈值

	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitempty" name:"SnapshotTimeout"`
	// 探针总开关

	AgentEnable *bool `json:"AgentEnable,omitempty" name:"AgentEnable"`
	// URL&nbsp;收敛开关&nbsp;0&nbsp;关&nbsp;1&nbsp;开

	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitempty" name:"UrlConvergenceSwitch"`
}

type Agent struct {

	// 上一次启动时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 123456

	RunningTime *int64 `json:"RunningTime,omitempty" name:"RunningTime"`
	// 版本号

	Version *string `json:"Version,omitempty" name:"Version"`
	// 是否启用该实例的探针
	//

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 应用标签

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 探针IP

	IP *string `json:"IP,omitempty" name:"IP"`
	// 运行状态&nbsp;

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeSpanTagListRequest struct {
	*tchttp.BaseRequest

	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 起始时间（单位：秒）

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 翻页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 页码（从1开始）

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 每页数目

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 是否返回匹配总数

	ExcludeTotal *bool `json:"ExcludeTotal,omitempty" name:"ExcludeTotal"`
	// 维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 结束时间（单位：秒）

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 查询过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeSpanTagListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanTagListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagCountValuesRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 维度名

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeTagCountValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagCountValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmMetricRecord struct {

	// field数组，用于指标的查询结果

	Fields []*ApmField `json:"Fields,omitempty" name:"Fields"`
	// tag数组，用于区分&nbsp;Groupby&nbsp;的对象

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
}

type ServiceDetail struct {

	// 标签

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 应用ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 业务系统ID

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
	// 用户appid

	AppID *int64 `json:"AppID,omitempty" name:"AppID"`
	// 主账号uin

	CreateUIN *string `json:"CreateUIN,omitempty" name:"CreateUIN"`
	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 应用描述

	ServiceDescription *string `json:"ServiceDescription,omitempty" name:"ServiceDescription"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type ThreadAnatomyView struct {

	// 实际耗时（ms）

	MonitorDuration *float64 `json:"MonitorDuration,omitempty" name:"MonitorDuration"`
	// 线程剖析堆栈信息

	AnatomyStacks []*ThreadAnatomyStack `json:"AnatomyStacks,omitempty" name:"AnatomyStacks"`
	// 开始时间

	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
}

type DescribeSpanTreeByIDRequest struct {
	*tchttp.BaseRequest

	// TraceID

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// Span查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// Span查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 代表请求的是否为完整瀑布图

	IsFold *bool `json:"IsFold,omitempty" name:"IsFold"`
	// 开启关键调用视图

	EntrySpanSelected *bool `json:"EntrySpanSelected,omitempty" name:"EntrySpanSelected"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeSpanTreeByIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanTreeByIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricLineDataRequest struct {
	*tchttp.BaseRequest

	// 指标查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 指标查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度（服务级别，接口级别）
	// service
	// interface_name

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 曲线聚合周期（60,3600,86400）

	AggregatePeriod *int64 `json:"AggregatePeriod,omitempty" name:"AggregatePeriod"`
	// Top&nbsp;N&nbsp;曲线参数

	TopLevel *int64 `json:"TopLevel,omitempty" name:"TopLevel"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标名列表

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
}

func (r *DescribeMetricLineDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricLineDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GeneralFilter struct {

	// 过滤值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 过滤维度名

	Key *string `json:"Key,omitempty" name:"Key"`
}

type DescribeApmServiceBriefResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用资源列表

		ServiceList []*ServiceBrief `json:"ServiceList,omitempty" name:"ServiceList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmServiceBriefResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmServiceBriefResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ComponentTopologyView struct {

	// 服务纬度的节点数量

	Service *int64 `json:"Service,omitempty" name:"Service"`
	// 数据库节点数量

	Database *int64 `json:"Database,omitempty" name:"Database"`
	// 消息队列节点数量

	MQ *int64 `json:"MQ,omitempty" name:"MQ"`
}

type ThreadAnatomyStack struct {

	// 方法名称

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 行号

	Line *string `json:"Line,omitempty" name:"Line"`
	// 耗时（ms）

	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
	// 子节点

	ChildrenStack []*ThreadAnatomyStack `json:"ChildrenStack,omitempty" name:"ChildrenStack"`
}

type SpanTreeView struct {

	// 压缩编码后的树结构json串

	CompressedTreeData *string `json:"CompressedTreeData,omitempty" name:"CompressedTreeData"`
	// 整条链路的&nbsp;span&nbsp;数量是否超过限制

	CntExceeded *bool `json:"CntExceeded,omitempty" name:"CntExceeded"`
	// 是否有内部查询超时

	QueryTimeout *bool `json:"QueryTimeout,omitempty" name:"QueryTimeout"`
	// 链路中&nbsp;span&nbsp;总数量

	SpanTotalCnt *int64 `json:"SpanTotalCnt,omitempty" name:"SpanTotalCnt"`
	// 树结构json串

	TreeData *string `json:"TreeData,omitempty" name:"TreeData"`
	// 列表

	ChartData *ChartData `json:"ChartData,omitempty" name:"ChartData"`
	// Trace概览信息

	TraceOverview *TraceOverview `json:"TraceOverview,omitempty" name:"TraceOverview"`
}

type ApmField struct {

	// 指标名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 指标数值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 指标所对应的单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 同比结果数组，推荐使用

	CompareVals []*APMKVItem `json:"CompareVals,omitempty" name:"CompareVals"`
	// 同比上一个周期的具体指标数值

	LastPeriodValue []*APMKV `json:"LastPeriodValue,omitempty" name:"LastPeriodValue"`
	// 同比指标值，已弃用，不建议使用

	CompareVal *string `json:"CompareVal,omitempty" name:"CompareVal"`
}

type DescribeTagValuesRequest struct {
	*tchttp.BaseRequest

	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 维度名

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// 开始时间（单位为秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间（单位为秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// Or&nbsp;过滤条件

	OrFilters []*Filter `json:"OrFilters,omitempty" name:"OrFilters"`
	// 使用类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeTagValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmApplicationConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Apm应用配置

		ApmAppConfig *ApmAppConfig `json:"ApmAppConfig,omitempty" name:"ApmAppConfig"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmApplicationConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralSpanListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Span&nbsp;分页列表

		Spans []*Span `json:"Spans,omitempty" name:"Spans"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralSpanListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralSpanListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopologyEdgeNew struct {

	// 源节点

	Source *string `json:"Source,omitempty" name:"Source"`
	// 目标节点

	Target *string `json:"Target,omitempty" name:"Target"`
	// 错误率

	ErrRate *float64 `json:"ErrRate,omitempty" name:"ErrRate"`
	// 吞吐量

	Qps *float64 `json:"Qps,omitempty" name:"Qps"`
	// 边类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 边颜色

	Color *string `json:"Color,omitempty" name:"Color"`
	// 边上目标节点类型&nbsp;应用/MQ/DB

	TargetComp *string `json:"TargetComp,omitempty" name:"TargetComp"`
	// 组件间调用次数

	ReqCnt *int64 `json:"ReqCnt,omitempty" name:"ReqCnt"`
	// 边ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 边权重

	Weight *float64 `json:"Weight,omitempty" name:"Weight"`
	// 响应时间

	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
	// Sql调用数

	SqlRequestCount *float64 `json:"SqlRequestCount,omitempty" name:"SqlRequestCount"`
	// Sql调用错误数

	SqlErrorRequestCount *float64 `json:"SqlErrorRequestCount,omitempty" name:"SqlErrorRequestCount"`
	// 边上源节点类型&nbsp;应用/MQ/DB

	SourceComp *string `json:"SourceComp,omitempty" name:"SourceComp"`
}

type ModifyQueryViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyQueryViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyQueryViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Line struct {

	// 时间序列

	TimeSerial []*int64 `json:"TimeSerial,omitempty" name:"TimeSerial"`
	// 数据序列

	DataSerial []*float64 `json:"DataSerial,omitempty" name:"DataSerial"`
	// 维度列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 指标数据单位

	MetricUnit *string `json:"MetricUnit,omitempty" name:"MetricUnit"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标中文名

	MetricNameCN *string `json:"MetricNameCN,omitempty" name:"MetricNameCN"`
}

type SelectorView struct {

	// 组件数量

	Component *ComponentTopologyView `json:"Component,omitempty" name:"Component"`
}

type DescribeApmAgentRequest struct {
	*tchttp.BaseRequest

	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 接入方式，现支持&nbsp;skywalking,&nbsp;ot,&nbsp;ebpf&nbsp;方式接入上报，不填默认为&nbsp;ot

	AgentType *string `json:"AgentType,omitempty" name:"AgentType"`
	// 上报环境，现支持&nbsp;pl&nbsp;(内网上报),&nbsp;public&nbsp;(外网),&nbsp;inner&nbsp;(自研&nbsp;VPC&nbsp;)环境上报，不传默认为&nbsp;public

	NetworkMode *string `json:"NetworkMode,omitempty" name:"NetworkMode"`
	// 语言，现支持&nbsp;java,&nbsp;golang,&nbsp;php,&nbsp;python,&nbsp;dotNet,&nbsp;nodejs&nbsp;语言上报，不传默认为&nbsp;golang

	LanguageEnvironment *string `json:"LanguageEnvironment,omitempty" name:"LanguageEnvironment"`
	// 上报方式，已弃用

	ReportMethod *string `json:"ReportMethod,omitempty" name:"ReportMethod"`
}

func (r *DescribeApmAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInfoByAppIdResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// appId相关apm返回具体信息

		Data *ApmAppId `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmInfoByAppIdResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInfoByAppIdResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceBriefsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// APM&nbsp;业务系统列表

		Instances []*ApmInstanceBrief `json:"Instances,omitempty" name:"Instances"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceBriefsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceBriefsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopologyNode struct {

	// 节点是否可以下钻

	CanDrillDown *bool `json:"CanDrillDown,omitempty" name:"CanDrillDown"`
	// 节点颜色

	Color *string `json:"Color,omitempty" name:"Color"`
	// 响应时间

	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
	// 吞吐量

	Qps *float64 `json:"Qps,omitempty" name:"Qps"`
	// 节点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 节点权重

	Weight *float64 `json:"Weight,omitempty" name:"Weight"`
	// 节点位置信息

	Position *Position `json:"Position,omitempty" name:"Position"`
	// 资源层信息

	Resource *Resource `json:"Resource,omitempty" name:"Resource"`
	// MQ&nbsp;消费者视角的响应时间&nbsp;ms

	ConsumerDuration *float64 `json:"ConsumerDuration,omitempty" name:"ConsumerDuration"`
	// MQ&nbsp;消费者视角的错误率&nbsp;%

	ConsumerErrRate *float64 `json:"ConsumerErrRate,omitempty" name:"ConsumerErrRate"`
	// MQ&nbsp;消费者视角的吞吐量

	ConsumerQps *float64 `json:"ConsumerQps,omitempty" name:"ConsumerQps"`
	// 应用&nbsp;ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
	// 错误率

	ErrRate *float64 `json:"ErrRate,omitempty" name:"ErrRate"`
	// 节点类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 节点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 拓扑节点视图名字

	NodeView *string `json:"NodeView,omitempty" name:"NodeView"`
	// 调用次数

	ReqCnt *int64 `json:"ReqCnt,omitempty" name:"ReqCnt"`
	// 节点大小

	Size *string `json:"Size,omitempty" name:"Size"`
	// 消息队列消费者视角的调用次数

	ConsumerReqCnt *int64 `json:"ConsumerReqCnt,omitempty" name:"ConsumerReqCnt"`
	// 节点是否为组件类型

	IsModule *bool `json:"IsModule,omitempty" name:"IsModule"`
	// 节点标签

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
}

type DescribeMetricLineDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标曲线返回数据

		Lines []*Line `json:"Lines,omitempty" name:"Lines"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricLineDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricLineDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 探针列表

		Agents []*Agent `json:"Agents,omitempty" name:"Agents"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Series struct {

	// 字段名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 耗时列表信息

	Data []*int64 `json:"Data,omitempty" name:"Data"`
	// 描述

	Stack *string `json:"Stack,omitempty" name:"Stack"`
}

type ModifyQueryViewRequest struct {
	*tchttp.BaseRequest

	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询视图&nbsp;ID

	ViewId *string `json:"ViewId,omitempty" name:"ViewId"`
	// 查询条件集合

	Filters *string `json:"Filters,omitempty" name:"Filters"`
	// 视图描述/名称

	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifyQueryViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyQueryViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralApmApplicationConfigRequest struct {
	*tchttp.BaseRequest

	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeGeneralApmApplicationConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralApmApplicationConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApmInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 业务系统&nbsp;ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateApmInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmAppId struct {

	// 剩余小时

	RemainderNumberHours *int64 `json:"RemainderNumberHours,omitempty" name:"RemainderNumberHours"`
	// 剩余额度

	RemainderNumberSpans *int64 `json:"RemainderNumberSpans,omitempty" name:"RemainderNumberSpans"`
	// 总额度

	TotalNumbersSpans *int64 `json:"TotalNumbersSpans,omitempty" name:"TotalNumbersSpans"`
	// 当前账号状态。{0:&nbsp;试用版,&nbsp;1:&nbsp;正式,&nbsp;4:&nbsp;未使用产品}

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 试用期过期时间戳

	EndTrialPeriodTime *int64 `json:"EndTrialPeriodTime,omitempty" name:"EndTrialPeriodTime"`
	// 该地域是否存在自动续费资源包

	AutoResourceExist *bool `json:"AutoResourceExist,omitempty" name:"AutoResourceExist"`
	// 当前状态的补充。{0:&nbsp;无,&nbsp;1:&nbsp;账号欠费,&nbsp;2:&nbsp;账号处于试用版到期,&nbsp;3:&nbsp;账号处于试用版限额}

	Reason *int64 `json:"Reason,omitempty" name:"Reason"`
	// 剩余天数

	RemainderNumberDays *int64 `json:"RemainderNumberDays,omitempty" name:"RemainderNumberDays"`
}

type DeleteTopologyViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopologyViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopologyViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagCountKV struct {

	// 维度Key值

	Key *string `json:"Key,omitempty" name:"Key"`
	// 维度对应的count数

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type DescribeTagCountValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// TagCount的键值对

		TagCountKV []*TagCountKV `json:"TagCountKV,omitempty" name:"TagCountKV"`
		// Tag维度分组查询总数

		TotalTagCount *int64 `json:"TotalTagCount,omitempty" name:"TotalTagCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagCountValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagCountValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricRecordsRequest struct {
	*tchttp.BaseRequest

	// 每页大小，默认为1000，合法取值范围为0~1000

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 页码

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页长

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 结束时间（单位为秒）

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// Or&nbsp;过滤条件

	OrFilters []*Filter `json:"OrFilters,omitempty" name:"OrFilters"`
	// 排序
	// 现支持的&nbsp;Key&nbsp;有：
	//
	// -&nbsp;startTime(开始时间)
	// -&nbsp;endTime(结束时间)
	// -&nbsp;duration(响应时间)
	//
	// 现支持的&nbsp;Value&nbsp;有：
	//
	// -&nbsp;desc(降序排序)
	// -&nbsp;asc(升序排序)

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 业务名称，控制台用户请填写taw。

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// 特殊处理查询结果

	Type *string `json:"Type,omitempty" name:"Type"`
	// 分页起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 开始时间（单位为秒）

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeMetricRecordsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricRecordsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpanThreadAnatomyListRequest struct {
	*tchttp.BaseRequest

	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// traceID和spanID
	// 或rid和spanID

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// span查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// span查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeSpanThreadAnatomyListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanThreadAnatomyListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpanProcess struct {

	// 应用服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// Tags&nbsp;标签数组

	Tags []*SpanTag `json:"Tags,omitempty" name:"Tags"`
}

type DescribeServiceNodesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceNodesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceNodesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 111

		TreeListView []*TraceItemView `json:"TreeListView,omitempty" name:"TreeListView"`
		// 列表项总个数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTraceListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTraceListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQueryViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 视图唯一&nbsp;ID

		ViewId *string `json:"ViewId,omitempty" name:"ViewId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQueryViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQueryViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInstancesRequest struct {
	*tchttp.BaseRequest

	// 按业务系统&nbsp;ID&nbsp;过滤

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 是否查询官方&nbsp;Demo&nbsp;业务系统（0=非&nbsp;Demo&nbsp;业务系统，1=Demo&nbsp;业务系统，默认为0）

	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitempty" name:"DemoInstanceFlag"`
	// 是否查询全地域业务系统（0=不查询全地域，1=查询全地域，默认为0）

	AllRegionsFlag *int64 `json:"AllRegionsFlag,omitempty" name:"AllRegionsFlag"`
	// Tag&nbsp;列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 按业务系统名过滤，支持模糊检索

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 按业务系统&nbsp;ID&nbsp;过滤，支持模糊检索

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeApmInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpanTreeByIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 树形图

		SpanTreeView *SpanTreeView `json:"SpanTreeView,omitempty" name:"SpanTreeView"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpanTreeByIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanTreeByIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type YAxis struct {

	// 接口列表

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 接口列表

	Data []*string `json:"Data,omitempty" name:"Data"`
}

type DeleteQueryViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteQueryViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQueryViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopologyViewRequest struct {
	*tchttp.BaseRequest

	// 视图方案ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteTopologyViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopologyViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpanTag struct {

	// 标签类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 标签Key
	// 注意：此字段可能返回&nbsp;null，表示取不到有效值。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值
	// 注意：此字段可能返回&nbsp;null，表示取不到有效值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ApmInstanceBrief struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 业务系统状态。{&nbsp;1:&nbsp;初始化中;&nbsp;2:&nbsp;运行中;&nbsp;4:&nbsp;限流}

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type DescribeAgentListRequest struct {
	*tchttp.BaseRequest

	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 应用实例

	ServiceInstance *string `json:"ServiceInstance,omitempty" name:"ServiceInstance"`
	// 应用标签

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeAgentListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 拓扑视图方案列表

		Views []*TopologyViewResponse `json:"Views,omitempty" name:"Views"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopologyViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceLinkRequest struct {
	*tchttp.BaseRequest

	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 业务名称（默认值：taw）

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// 每页大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 页码

	PageIndex *int64 `json:"PageIndex,omitempty" name:"PageIndex"`
	// 页长

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 分页起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeServiceLinkRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceLinkRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0代表无须提示，1代表存在探针即将过期，2代表存在探针已过期

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteQueryViewRequest struct {
	*tchttp.BaseRequest

	// 视图唯一&nbsp;ID

	ViewId *string `json:"ViewId,omitempty" name:"ViewId"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteQueryViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteQueryViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQueryViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 视图查询条件集合

		Filters *string `json:"Filters,omitempty" name:"Filters"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQueryViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueryViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAllQueryViewsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询视图集合

		Views []*QueryView `json:"Views,omitempty" name:"Views"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllQueryViewsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllQueryViewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmApplicationConfigRequest struct {
	*tchttp.BaseRequest

	// 服务名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeApmApplicationConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmApplicationConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceNodesRequest struct {
	*tchttp.BaseRequest

	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeServiceNodesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceNodesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmAgentInfo struct {

	// Token&nbsp;信息

	Token *string `json:"Token,omitempty" name:"Token"`
	// 外网上报地址

	PublicCollectorURL *string `json:"PublicCollectorURL,omitempty" name:"PublicCollectorURL"`
	// 自研&nbsp;VPC&nbsp;上报地址

	InnerCollectorURL *string `json:"InnerCollectorURL,omitempty" name:"InnerCollectorURL"`
	// 内网上报地址(&nbsp;Private&nbsp;Link&nbsp;上报地址)

	PrivateLinkCollectorURL *string `json:"PrivateLinkCollectorURL,omitempty" name:"PrivateLinkCollectorURL"`
	// Agent&nbsp;下载地址

	AgentDownloadURL *string `json:"AgentDownloadURL,omitempty" name:"AgentDownloadURL"`
	// Collector&nbsp;上报地址

	CollectorURL *string `json:"CollectorURL,omitempty" name:"CollectorURL"`
}

type Position struct {

	// 节点位置横坐标

	X *float64 `json:"X,omitempty" name:"X"`
	// 节点位置纵坐标

	Y *float64 `json:"Y,omitempty" name:"Y"`
}

type DescribeAllQueryViewsRequest struct {
	*tchttp.BaseRequest

	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否是拓扑图视图

	IsTopo *bool `json:"IsTopo,omitempty" name:"IsTopo"`
}

func (r *DescribeAllQueryViewsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAllQueryViewsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmTag struct {

	// 维度Key(列名，标签Key)

	Key *string `json:"Key,omitempty" name:"Key"`
	// 维度值（标签值）

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ModifyGeneralApmApplicationConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回值描述

		Message *string `json:"Message,omitempty" name:"Message"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyGeneralApmApplicationConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGeneralApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceRequest struct {
	*tchttp.BaseRequest

	// 过滤应用的条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpanThreadAnatomyListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 线程剖析总信息

		ThreadAnatomyView *ThreadAnatomyView `json:"ThreadAnatomyView,omitempty" name:"ThreadAnatomyView"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpanThreadAnatomyListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanThreadAnatomyListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateApmInstanceRequest struct {
	*tchttp.BaseRequest

	// 业务系统描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// Trace&nbsp;数据保存时长（单位：天，默认存储时长为3天）

	TraceDuration *int64 `json:"TraceDuration,omitempty" name:"TraceDuration"`
	// 业务系统&nbsp;Tag&nbsp;列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 业务系统上报额度值，默认赋值为0表示不限制上报额度，已废弃

	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitempty" name:"SpanDailyCounters"`
	// 业务系统的计费模式（0=按量付费，1=预付费）

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 是否为免费版业务系统（0=付费版；1=TSF&nbsp;受限免费版；2=免费版）

	Free *int64 `json:"Free,omitempty" name:"Free"`
	// 业务系统名

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateApmInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateApmInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmServiceRequest struct {
	*tchttp.BaseRequest

	// 应用描述

	ServiceDescription *string `json:"ServiceDescription,omitempty" name:"ServiceDescription"`
	// 标签列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 应用ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
}

func (r *ModifyApmServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceBriefsRequest struct {
	*tchttp.BaseRequest

	// Tag&nbsp;列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 是否查询官方&nbsp;Demo&nbsp;业务系统（0:&nbsp;不查询;&nbsp;1:&nbsp;查询）

	DemoInstanceFlag *int64 `json:"DemoInstanceFlag,omitempty" name:"DemoInstanceFlag"`
	// 是否查询所有地域（0:&nbsp;不查询;&nbsp;1:&nbsp;查询）

	AllRegionsFlag *int64 `json:"AllRegionsFlag,omitempty" name:"AllRegionsFlag"`
}

func (r *DescribeInstanceBriefsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceBriefsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmApplicationConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApmApplicationConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTagValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 维度值列表

		Values []*string `json:"Values,omitempty" name:"Values"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTagValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTagValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyViewRequest struct {
	*tchttp.BaseRequest

	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeTopologyViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TkeMeta struct {

	// pod&nbsp;ip

	PodIP *string `json:"PodIP,omitempty" name:"PodIP"`
	// node&nbsp;ip

	NodeIP *string `json:"NodeIP,omitempty" name:"NodeIP"`
	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 集群ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// pod&nbsp;name

	PodName *string `json:"PodName,omitempty" name:"PodName"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载

	Deployment *string `json:"Deployment,omitempty" name:"Deployment"`
}

type TopologyViewResponse struct {

	// 视图方案ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 是否收藏

	Mark *bool `json:"Mark,omitempty" name:"Mark"`
	// 标签

	Labels []*ApmTag `json:"Labels,omitempty" name:"Labels"`
	// 是否默认视图

	Default *bool `json:"Default,omitempty" name:"Default"`
	// 视图方案勾选情况

	Selectors *Selectors `json:"Selectors,omitempty" name:"Selectors"`
	// 视图方案名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type Filter struct {

	// 过滤方式（=,&nbsp;!=,&nbsp;in）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 过滤维度名

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过滤值，in过滤方式用逗号分割多个值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Selectors struct {

	// 组件勾选情况

	Component []*string `json:"Component,omitempty" name:"Component"`
}

type DescribeMetricRecordsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// 查询指标结果集条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricRecordsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricRecordsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type TerminateApmServiceRequest struct {
	*tchttp.BaseRequest

	// 应用资源ID

	ServiceId *string `json:"ServiceId,omitempty" name:"ServiceId"`
}

func (r *TerminateApmServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateApmServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Resource struct {

	// 资源类型

	Type []*string `json:"Type,omitempty" name:"Type"`
	// tke资源层信息

	TKEMeta []*TkeMeta `json:"TKEMeta,omitempty" name:"TKEMeta"`
	// cvm资源信息

	CVMMeta []*CVMMeta `json:"CVMMeta,omitempty" name:"CVMMeta"`
}

type ApmServiceMetric struct {

	// field数组

	Fields []*ApmField `json:"Fields,omitempty" name:"Fields"`
	// tag数组

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 应用信息

	ServiceDetail *ServiceDetail `json:"ServiceDetail,omitempty" name:"ServiceDetail"`
}

type DescribeApmServiceBriefRequest struct {
	*tchttp.BaseRequest

	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 标签

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 是否为demo模式

	Demo *bool `json:"Demo,omitempty" name:"Demo"`
	// 应用ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

func (r *DescribeApmServiceBriefRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmServiceBriefRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTraceListRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 列表项个数

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// trace查询开始时间戳（单位:秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// trace查询结束时间戳（单位:秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// sql通用过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTraceListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTraceListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApmInstanceDetail struct {

	// 是否开启目录遍历检测（0-关闭，1-开启）

	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitempty" name:"IsDirectoryTraversalAnalysis"`
	// 是否开启模板引擎注入检测（0-关闭，1-开启）

	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitempty" name:"IsTemplateEngineInjectionAnalysis"`
	// 采样率（单位：%）

	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
	// 是否&nbsp;TSF&nbsp;默认业务系统（0=否，1=是）

	DefaultTSF *int64 `json:"DefaultTSF,omitempty" name:"DefaultTSF"`
	// 关联的&nbsp;Dashboard&nbsp;ID

	DashboardTopicID *string `json:"DashboardTopicID,omitempty" name:"DashboardTopicID"`
	// 是否开启包含任意文件检测（0-关闭，1-开启）

	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitempty" name:"IsIncludeAnyFileAnalysis"`
	// 是否开启JNDI注入检测（0-关闭，1-开启）

	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitempty" name:"IsJNDIInjectionAnalysis"`
	// URL数字分段收敛阈值

	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitempty" name:"UrlNumberSegmentThreshold"`
	// AppID&nbsp;信息

	AppId *int64 `json:"AppId,omitempty" name:"AppId"`
	// 是否开启错误采样（0=关,&nbsp;1=开）

	ErrorSample *int64 `json:"ErrorSample,omitempty" name:"ErrorSample"`
	// CLS&nbsp;日志所在地域

	LogRegion *string `json:"LogRegion,omitempty" name:"LogRegion"`
	// 是否关联&nbsp;Dashboard（0=关,&nbsp;1=开）

	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitempty" name:"IsRelatedDashboard"`
	// 是否开内存马执行检测（0=关，&nbsp;1=开）

	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitempty" name:"IsMemoryHijackingAnalysis"`
	// URL长分段收敛阈值

	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitempty" name:"UrlLongSegmentThreshold"`
	// 该业务系统服务端应用数量

	ServiceCount *int64 `json:"ServiceCount,omitempty" name:"ServiceCount"`
	// 日均上报&nbsp;Span&nbsp;数

	CountOfReportSpanPerDay *int64 `json:"CountOfReportSpanPerDay,omitempty" name:"CountOfReportSpanPerDay"`
	// Trace&nbsp;数据保存时长（单位：天）

	TraceDuration *int64 `json:"TraceDuration,omitempty" name:"TraceDuration"`
	// 业务系统上报额度

	SpanDailyCounters *int64 `json:"SpanDailyCounters,omitempty" name:"SpanDailyCounters"`
	// 错误警示线（单位：%）

	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitempty" name:"ErrRateThreshold"`
	// 日志主题&nbsp;ID

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// 该业务系统客户端应用数量

	ClientCount *int64 `json:"ClientCount,omitempty" name:"ClientCount"`
	// 业务系统计费模式（1为预付费，0为按量付费）

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 业务系统状态。{
	// 1:&nbsp;初始化中;&nbsp;2:&nbsp;运行中;&nbsp;4:&nbsp;限流}

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 业务系统是否已开通计费（0=未开通，1=已开通）

	BillingInstance *int64 `json:"BillingInstance,omitempty" name:"BillingInstance"`
	// 日志功能开关（0=关，&nbsp;1=开）

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// CLS&nbsp;日志集

	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`
	// 是否开启&nbsp;SQL&nbsp;注入分析（0=关，&nbsp;1=开）

	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitempty" name:"IsSqlInjectionAnalysis"`
	// traceId的索引key:&nbsp;当CLS索引类型为键值索引时生效

	LogTraceIdKey *string `json:"LogTraceIdKey,omitempty" name:"LogTraceIdKey"`
	// 是否开启读取任意文件检测（0-关闭，1-开启）

	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitempty" name:"IsReadAnyFileAnalysis"`
	// 是否开启表达式注入检测（0-关闭，1-开启）

	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitempty" name:"IsExpressionInjectionAnalysis"`
	// 业务系统名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 业务系统&nbsp;Tag&nbsp;列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 存储使用量(单位：MB)

	AmountOfUsedStorage *float64 `json:"AmountOfUsedStorage,omitempty" name:"AmountOfUsedStorage"`
	// 该业务系统最近2天活跃应用数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 限流原因。{
	// 1:&nbsp;正式版限额;
	// 2:&nbsp;试用版限额;
	// 4:&nbsp;试用版到期;
	// 8:&nbsp;账号欠费
	// }

	StopReason *int64 `json:"StopReason,omitempty" name:"StopReason"`
	// 是否开远程命令执行检测（0=关，&nbsp;1=开）

	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitempty" name:"IsRemoteCommandExecutionAnalysis"`
	// CLS索引类型(0=全文索引，1=键值索引)

	LogIndexType *int64 `json:"LogIndexType,omitempty" name:"LogIndexType"`
	// 是否开启删除任意文件检测（0-关闭，1-开启）

	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitempty" name:"IsDeleteAnyFileAnalysis"`
	// 业务系统描述信息

	Description *string `json:"Description,omitempty" name:"Description"`
	// 采样慢调用保存阈值（单位：ms）

	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitempty" name:"SlowRequestSavedThreshold"`
	// 日志源

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// Metric&nbsp;数据保存时长（单位：天）

	MetricDuration *int64 `json:"MetricDuration,omitempty" name:"MetricDuration"`
	// 用户自定义展示标签列表

	CustomShowTags []*string `json:"CustomShowTags,omitempty" name:"CustomShowTags"`
	// 业务系统计费模式是否生效

	PayModeEffective *bool `json:"PayModeEffective,omitempty" name:"PayModeEffective"`
	// 是否免费（0=否，1=限额免费，2=完全免费），默认0

	Free *int64 `json:"Free,omitempty" name:"Free"`
	// 是否开启上传任意文件检测（0-关闭，1-开启）

	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitempty" name:"IsUploadAnyFileAnalysis"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 业务系统所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 创建人&nbsp;Uin

	CreateUin *string `json:"CreateUin,omitempty" name:"CreateUin"`
	// 是否开启脚本引擎注入检测（0-关闭，1-开启）

	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitempty" name:"IsScriptEngineInjectionAnalysis"`
	// 是否开启JNI注入检测（0-关闭，1-开启）

	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitempty" name:"IsJNIInjectionAnalysis"`
	// 是否开启Webshell后门检测（0-关闭，1-开启）

	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitempty" name:"IsWebshellBackdoorAnalysis"`
	// 是否开启反序列化检测（0-关闭，1-开启）

	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitempty" name:"IsDeserializationAnalysis"`
	// 业务系统鉴权&nbsp;token

	Token *string `json:"Token,omitempty" name:"Token"`
	// 响应时间警示线（单位：ms）

	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitempty" name:"ResponseDurationWarningThreshold"`
	// 是否开启组件漏洞检测（0=关，&nbsp;1=开）

	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitempty" name:"IsInstrumentationVulnerabilityScan"`
}

type DescribeApmServiceMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用总数

		ApplicationCount *int64 `json:"ApplicationCount,omitempty" name:"ApplicationCount"`
		// 页码

		Page *int64 `json:"Page,omitempty" name:"Page"`
		// 页大小

		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
		// 应用指标列表
		// 注意：此字段可能返回&nbsp;null，表示取不到有效值。

		ServiceMetricList []*ApmServiceMetric `json:"ServiceMetricList,omitempty" name:"ServiceMetricList"`
		// 符合筛选条件的应用数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 警示异常应用数

		WarningErrorCount *int64 `json:"WarningErrorCount,omitempty" name:"WarningErrorCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmServiceMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmServiceMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopologyViewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 视图方案ID

		Id *string `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopologyViewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopologyViewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CVMMeta struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeApmServiceMetricRequest struct {
	*tchttp.BaseRequest

	// 应用状态筛选

	ServiceStatus *string `json:"ServiceStatus,omitempty" name:"ServiceStatus"`
	// 标签列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 应用ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 是否demo模式

	Demo *bool `json:"Demo,omitempty" name:"Demo"`
}

func (r *DescribeApmServiceMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmServiceMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralSpanListRequest struct {
	*tchttp.BaseRequest

	// 分页

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Span&nbsp;查询开始时间戳（单位：秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// Span&nbsp;查询结束时间戳（单位：秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 通用过滤参数

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序
	// 现支持的&nbsp;Key&nbsp;有：
	//
	// -&nbsp;startTime(开始时间)
	// -&nbsp;endTime(结束时间)
	// -&nbsp;duration(响应时间)
	//
	// 现支持的&nbsp;Value&nbsp;有：
	//
	// -&nbsp;desc(降序排序)
	// -&nbsp;asc(升序排序)

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 业务自身服务名，控制台用户请填写taw

	BusinessName *string `json:"BusinessName,omitempty" name:"BusinessName"`
	// 单页项目个数，默认为10000，合法取值范围为0～10000

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeGeneralSpanListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralSpanListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSpanTagListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总的record数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 列表项

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSpanTagListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSpanTagListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpanReference struct {

	// 关联关系类型

	RefType *string `json:"RefType,omitempty" name:"RefType"`
	// Span&nbsp;ID

	SpanID *string `json:"SpanID,omitempty" name:"SpanID"`
	// Trace&nbsp;ID

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
}

type DescribeTopologyNewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 边集合

		Edges []*TopologyEdgeNew `json:"Edges,omitempty" name:"Edges"`
		// 拓扑图是否有修改

		TopologyModifyFlag *int64 `json:"TopologyModifyFlag,omitempty" name:"TopologyModifyFlag"`
		// 节点数量

		Selectors *SelectorView `json:"Selectors,omitempty" name:"Selectors"`
		// 节点集合

		Nodes []*TopologyNode `json:"Nodes,omitempty" name:"Nodes"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopologyNewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyNewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Agent&nbsp;信息

		ApmAgent *ApmAgentInfo `json:"ApmAgent,omitempty" name:"ApmAgent"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmInstanceRequest struct {
	*tchttp.BaseRequest

	// 是否开启内存马检测

	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitempty" name:"IsMemoryHijackingAnalysis"`
	// 是否开启目录遍历检测（0-关闭，1-开启）

	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitempty" name:"IsDirectoryTraversalAnalysis"`
	// URL数字分段收敛阈值

	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitempty" name:"UrlNumberSegmentThreshold"`
	// 日志集，开启日志功能后才会生效

	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`
	// URL长分段收敛阈值

	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitempty" name:"UrlLongSegmentThreshold"`
	// 是否开启包含任意文件检测（0-关闭，1-开启）

	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitempty" name:"IsIncludeAnyFileAnalysis"`
	// 是否开启脚本引擎注入检测（0-关闭，1-开启）

	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitempty" name:"IsScriptEngineInjectionAnalysis"`
	// traceId的索引key:&nbsp;当CLS索引类型为键值索引时生效

	LogTraceIdKey *string `json:"LogTraceIdKey,omitempty" name:"LogTraceIdKey"`
	// Tag&nbsp;列表

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// CLS&nbsp;日志主题&nbsp;ID，开启日志功能后才会生效

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// 是否开启表达式注入检测（0-关闭，1-开启）

	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitempty" name:"IsExpressionInjectionAnalysis"`
	// 是否开启Webshell后门检测（0-关闭，1-开启）

	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitempty" name:"IsWebshellBackdoorAnalysis"`
	// 业务系统名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 业务系统描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// Trace&nbsp;数据保存时长（单位：天）

	TraceDuration *int64 `json:"TraceDuration,omitempty" name:"TraceDuration"`
	// 是否开启计费

	OpenBilling *bool `json:"OpenBilling,omitempty" name:"OpenBilling"`
	// 业务系统上报额度

	SpanDailyCounters *uint64 `json:"SpanDailyCounters,omitempty" name:"SpanDailyCounters"`
	// 修改计费模式（1为预付费，0为按量付费）

	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`
	// 是否免费（0=付费版；1=TSF&nbsp;受限免费版；2=免费版），默认0

	Free *int64 `json:"Free,omitempty" name:"Free"`
	// 是否开启&nbsp;SQL&nbsp;注入检测（0=关,1=开）

	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitempty" name:"IsSqlInjectionAnalysis"`
	// 是否开启删除任意文件检测（0-关闭，1-开启）

	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitempty" name:"IsDeleteAnyFileAnalysis"`
	// 是否开启读取任意文件检测（0-关闭，1-开启）

	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitempty" name:"IsReadAnyFileAnalysis"`
	// 是否开启JNDI注入检测（0-关闭，1-开启）

	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitempty" name:"IsJNDIInjectionAnalysis"`
	// 是否开启JNI注入检测（0-关闭，1-开启）

	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitempty" name:"IsJNIInjectionAnalysis"`
	// 采样率（单位：%）

	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
	// 是否开启错误采样（0=关,&nbsp;1=开）

	ErrorSample *int64 `json:"ErrorSample,omitempty" name:"ErrorSample"`
	// 是否开启日志功能（0=关,&nbsp;1=开）

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// 用户自定义展示标签列表

	CustomShowTags []*string `json:"CustomShowTags,omitempty" name:"CustomShowTags"`
	// 响应时间警示线

	ResponseDurationWarningThreshold *int64 `json:"ResponseDurationWarningThreshold,omitempty" name:"ResponseDurationWarningThreshold"`
	// 是否开启远程命令攻击检测

	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitempty" name:"IsRemoteCommandExecutionAnalysis"`
	// CLS索引类型(0=全文索引，1=键值索引)

	LogIndexType *int64 `json:"LogIndexType,omitempty" name:"LogIndexType"`
	// 是否开启上传任意文件检测（0-关闭，1-开启）

	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitempty" name:"IsUploadAnyFileAnalysis"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 错误率警示线，当应用的平均错误率超出该阈值时，系统会给出异常提示。

	ErrRateThreshold *int64 `json:"ErrRateThreshold,omitempty" name:"ErrRateThreshold"`
	// 采样慢调用保存阈值（单位：ms）

	SlowRequestSavedThreshold *int64 `json:"SlowRequestSavedThreshold,omitempty" name:"SlowRequestSavedThreshold"`
	// 日志源，开启日志功能后才会生效

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 关联的&nbsp;Dashboard&nbsp;ID，开启关联&nbsp;Dashboard&nbsp;后才会生效

	DashboardTopicID *string `json:"DashboardTopicID,omitempty" name:"DashboardTopicID"`
	// 是否开启组件漏洞检测（0=关,1=开）

	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitempty" name:"IsInstrumentationVulnerabilityScan"`
	// 是否开启模板引擎注入检测（0-关闭，1-开启）

	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitempty" name:"IsTemplateEngineInjectionAnalysis"`
	// 是否开启反序列化检测（0-关闭，1-开启）

	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitempty" name:"IsDeserializationAnalysis"`
	// 日志地域，开启日志功能后才会生效

	LogRegion *string `json:"LogRegion,omitempty" name:"LogRegion"`
	// 是否关联&nbsp;Dashboard（0=关,1=开）

	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitempty" name:"IsRelatedDashboard"`
}

func (r *ModifyApmInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SpanLog struct {

	// 日志时间戳

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 标签

	Fields []*SpanTag `json:"Fields,omitempty" name:"Fields"`
}

type DescribeServiceLinkResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*ApmMetricRecord `json:"Records,omitempty" name:"Records"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeServiceLinkResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceLinkResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeServiceOverviewRequest struct {
	*tchttp.BaseRequest

	// 指标列表

	Metrics []*QueryMetricItem `json:"Metrics,omitempty" name:"Metrics"`
	// 结束时间（单位：秒）

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 聚合维度

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序方式
	// Value&nbsp;填写：
	// -&nbsp;asc：对查询指标进行升序排序
	// -&nbsp;desc：对查询指标进行降序排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 每页大小

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 分页起始点

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 开始时间（单位：秒）

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeServiceOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeServiceOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApmInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMetricItem struct {

	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 同比，现支持&nbsp;CompareByYesterday&nbsp;(与昨天相比)和CompareByLastWeek&nbsp;(与上周相比)&nbsp;

	Compares []*string `json:"Compares,omitempty" name:"Compares"`
	// 同比，已弃用，不建议使用

	Compare *string `json:"Compare,omitempty" name:"Compare"`
}

type TraceItemView struct {

	// 状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// trace唯一标识符

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
	// 入口接口名

	InterfaceName *string `json:"InterfaceName,omitempty" name:"InterfaceName"`
	// 开始采样时间

	SamplingTime *string `json:"SamplingTime,omitempty" name:"SamplingTime"`
	// 入口应用

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 耗时（单位毫秒）

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// span唯一标识符

	SpanID *string `json:"SpanID,omitempty" name:"SpanID"`
}

type ApmAppConfig struct {

	// URL排除正则

	UrlExclude *string `json:"UrlExclude,omitempty" name:"UrlExclude"`
	// 是否开启agent

	AgentEnable *bool `json:"AgentEnable,omitempty" name:"AgentEnable"`
	// 是否开启应用日志配置

	EnableLogConfig *bool `json:"EnableLogConfig,omitempty" name:"EnableLogConfig"`
	// 应用ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 是否开启上传任意文件检测（0-关闭，1-开启）

	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitempty" name:"IsUploadAnyFileAnalysis"`
	// 是否开启反序列化检测（0-关闭，1-开启）

	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitempty" name:"IsDeserializationAnalysis"`
	// 接口名称自动收敛开关（0-关闭，1-开启）

	UrlAutoConvergenceEnable *bool `json:"UrlAutoConvergenceEnable,omitempty" name:"UrlAutoConvergenceEnable"`
	// 错误码过滤

	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitempty" name:"ErrorCodeFilter"`
	// 日志来源

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 日志主题ID

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// 探针每秒上报trace数

	TraceRateLimit *int64 `json:"TraceRateLimit,omitempty" name:"TraceRateLimit"`
	// 是否开启读取任意文件检测（0-关闭，1-开启）

	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitempty" name:"IsReadAnyFileAnalysis"`
	// 是否开启线程剖析

	EnableSnapshot *bool `json:"EnableSnapshot,omitempty" name:"EnableSnapshot"`
	// 组件列表

	InstrumentList []*Instrument `json:"InstrumentList,omitempty" name:"InstrumentList"`
	// 是否开启应用诊断开关

	EventEnable *bool `json:"EventEnable,omitempty" name:"EventEnable"`
	// 是否开启SQL注入分析

	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitempty" name:"IsSqlInjectionAnalysis"`
	// 是否开启删除任意文件检测（0-关闭，1-开启）

	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitempty" name:"IsDeleteAnyFileAnalysis"`
	// 是否开启内存马检测分析

	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitempty" name:"IsMemoryHijackingAnalysis"`
	// CLS索引类型(0=全文索引，1=键值索引)

	LogIndexType *int64 `json:"LogIndexType,omitempty" name:"LogIndexType"`
	// 是否开启包含任意文件检测（0-关闭，1-开启）

	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitempty" name:"IsIncludeAnyFileAnalysis"`
	// 实例ID

	InstanceKey *string `json:"InstanceKey,omitempty" name:"InstanceKey"`
	// URL收敛正则

	UrlConvergence *string `json:"UrlConvergence,omitempty" name:"UrlConvergence"`
	// 服务组件类型

	Components *string `json:"Components,omitempty" name:"Components"`
	// 需过滤的接口名

	IgnoreOperationName *string `json:"IgnoreOperationName,omitempty" name:"IgnoreOperationName"`
	// 是否开启应用级别配置

	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitempty" name:"EnableSecurityConfig"`
	// 是否开启远程命令执行分析

	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitempty" name:"IsRemoteCommandExecutionAnalysis"`
	// traceId的索引key:&nbsp;当CLS索引类型为键值索引时生效

	LogTraceIdKey *string `json:"LogTraceIdKey,omitempty" name:"LogTraceIdKey"`
	// 是否开启目录遍历检测（0-关闭，1-开启）

	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitempty" name:"IsDirectoryTraversalAnalysis"`
	// 是否开启组件漏洞检测

	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitempty" name:"IsInstrumentationVulnerabilityScan"`
	// 是否开启脚本引擎注入检测（0-关闭，1-开启）

	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitempty" name:"IsScriptEngineInjectionAnalysis"`
	// 是否开启表达式注入检测（0-关闭，1-开启）

	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitempty" name:"IsExpressionInjectionAnalysis"`
	// 是否开启Webshell后门检测（0-关闭，1-开启）

	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitempty" name:"IsWebshellBackdoorAnalysis"`
	// URL长分段收敛阈值

	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitempty" name:"UrlLongSegmentThreshold"`
	// URL数字分段收敛阈值

	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitempty" name:"UrlNumberSegmentThreshold"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 是否开启链路压缩

	TraceSquash *bool `json:"TraceSquash,omitempty" name:"TraceSquash"`
	// 应用是否开启dashboard配置：&nbsp;false&nbsp;关（与业务系统保持一致）/true&nbsp;开（应用级配置）

	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitempty" name:"EnableDashboardConfig"`
	// 是否关联dashboard：&nbsp;0&nbsp;关&nbsp;1&nbsp;开

	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitempty" name:"IsRelatedDashboard"`
	// dashboard&nbsp;ID

	DashboardTopicID *string `json:"DashboardTopicID,omitempty" name:"DashboardTopicID"`
	// 是否开启模板引擎注入检测（0-关闭，1-开启）

	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitempty" name:"IsTemplateEngineInjectionAnalysis"`
	// 是否开启JNDI注入检测（0-关闭，1-开启）

	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitempty" name:"IsJNDIInjectionAnalysis"`
	// 是否开启JNI注入检测（0-关闭，1-开启）

	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitempty" name:"IsJNIInjectionAnalysis"`
	// URL收敛开关

	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitempty" name:"UrlConvergenceSwitch"`
	// URL收敛阈值

	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitempty" name:"UrlConvergenceThreshold"`
	// 异常过滤正则

	ExceptionFilter *string `json:"ExceptionFilter,omitempty" name:"ExceptionFilter"`
	// 日志所在地域

	LogRegion *string `json:"LogRegion,omitempty" name:"LogRegion"`
	// 是否开启日志&nbsp;0&nbsp;关&nbsp;1&nbsp;开

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// CLS日志集&nbsp;|&nbsp;ES集群ID

	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`
	// 线程剖析超时阈值

	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitempty" name:"SnapshotTimeout"`
	// 探针接口相关配置

	AgentOperationConfigView *AgentOperationConfigView `json:"AgentOperationConfigView,omitempty" name:"AgentOperationConfigView"`
}

type DescribeQueryViewRequest struct {
	*tchttp.BaseRequest

	// 视图唯一&nbsp;ID

	ViewId *string `json:"ViewId,omitempty" name:"ViewId"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否是拓扑图视图

	IsTopo *bool `json:"IsTopo,omitempty" name:"IsTopo"`
}

func (r *DescribeQueryViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQueryViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopologyViewRequest struct {
	*tchttp.BaseRequest

	// 视图方案名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 收藏标记

	Mark *bool `json:"Mark,omitempty" name:"Mark"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 视图方案ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 标签

	Labels []*ApmTag `json:"Labels,omitempty" name:"Labels"`
	// 组件勾选情况

	Selectors *Selectors `json:"Selectors,omitempty" name:"Selectors"`
	// 节点位置信息

	Snapshot []*Snapshot `json:"Snapshot,omitempty" name:"Snapshot"`
}

func (r *ModifyTopologyViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopologyViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmApplicationConfigRequest struct {
	*tchttp.BaseRequest

	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// URL收敛开关,0&nbsp;关&nbsp;|&nbsp;1&nbsp;开

	UrlConvergenceSwitch *int64 `json:"UrlConvergenceSwitch,omitempty" name:"UrlConvergenceSwitch"`
	// URL排除正则规则，逗号分隔

	UrlExclude *string `json:"UrlExclude,omitempty" name:"UrlExclude"`
	// 探针接口相关配置

	AgentOperationConfigView *AgentOperationConfigView `json:"AgentOperationConfigView,omitempty" name:"AgentOperationConfigView"`
	// 是否关联dashboard：&nbsp;0&nbsp;关&nbsp;1&nbsp;开

	IsRelatedDashboard *int64 `json:"IsRelatedDashboard,omitempty" name:"IsRelatedDashboard"`
	// 是否开启SQL注入分析

	IsSqlInjectionAnalysis *int64 `json:"IsSqlInjectionAnalysis,omitempty" name:"IsSqlInjectionAnalysis"`
	// 是否开启组件漏洞检测

	IsInstrumentationVulnerabilityScan *int64 `json:"IsInstrumentationVulnerabilityScan,omitempty" name:"IsInstrumentationVulnerabilityScan"`
	// 是否开启反序列化检测（0-关闭，1-开启）

	IsDeserializationAnalysis *int64 `json:"IsDeserializationAnalysis,omitempty" name:"IsDeserializationAnalysis"`
	// 错误码过滤，逗号分隔

	ErrorCodeFilter *string `json:"ErrorCodeFilter,omitempty" name:"ErrorCodeFilter"`
	// 日志地域

	LogRegion *string `json:"LogRegion,omitempty" name:"LogRegion"`
	// CLS&nbsp;日志集&nbsp;|&nbsp;ES&nbsp;集群ID

	LogSet *string `json:"LogSet,omitempty" name:"LogSet"`
	// 线程剖析超时阈值

	SnapshotTimeout *int64 `json:"SnapshotTimeout,omitempty" name:"SnapshotTimeout"`
	// 是否开启agent

	AgentEnable *bool `json:"AgentEnable,omitempty" name:"AgentEnable"`
	// 是否开启应用日志配置

	EnableLogConfig *bool `json:"EnableLogConfig,omitempty" name:"EnableLogConfig"`
	// 是否开启应用安全配置

	EnableSecurityConfig *bool `json:"EnableSecurityConfig,omitempty" name:"EnableSecurityConfig"`
	// 是否开启模板引擎注入检测（0-关闭，1-开启）

	IsTemplateEngineInjectionAnalysis *int64 `json:"IsTemplateEngineInjectionAnalysis,omitempty" name:"IsTemplateEngineInjectionAnalysis"`
	// 异常过滤正则规则，逗号分隔

	ExceptionFilter *string `json:"ExceptionFilter,omitempty" name:"ExceptionFilter"`
	// URL收敛正则规则，逗号分隔

	UrlConvergence *string `json:"UrlConvergence,omitempty" name:"UrlConvergence"`
	// 日志开关&nbsp;0&nbsp;关&nbsp;1&nbsp;开

	IsRelatedLog *int64 `json:"IsRelatedLog,omitempty" name:"IsRelatedLog"`
	// 是否开启应用诊断的开关

	EventEnable *bool `json:"EventEnable,omitempty" name:"EventEnable"`
	// traceId的索引key:&nbsp;当CLS索引类型为键值索引时生效

	LogTraceIdKey *string `json:"LogTraceIdKey,omitempty" name:"LogTraceIdKey"`
	// 是否开启包含任意文件检测（0-关闭，1-开启）

	IsIncludeAnyFileAnalysis *int64 `json:"IsIncludeAnyFileAnalysis,omitempty" name:"IsIncludeAnyFileAnalysis"`
	// URL数字分段收敛阈值

	UrlNumberSegmentThreshold *int64 `json:"UrlNumberSegmentThreshold,omitempty" name:"UrlNumberSegmentThreshold"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 日志主题ID

	LogTopicID *string `json:"LogTopicID,omitempty" name:"LogTopicID"`
	// 是否开启链路压缩

	TraceSquash *bool `json:"TraceSquash,omitempty" name:"TraceSquash"`
	// CLS索引类型(0=全文索引，1=键值索引)

	LogIndexType *int64 `json:"LogIndexType,omitempty" name:"LogIndexType"`
	// 是否开启内存马检测

	IsMemoryHijackingAnalysis *int64 `json:"IsMemoryHijackingAnalysis,omitempty" name:"IsMemoryHijackingAnalysis"`
	// 是否开启读取任意文件检测（0-关闭，1-开启）

	IsReadAnyFileAnalysis *int64 `json:"IsReadAnyFileAnalysis,omitempty" name:"IsReadAnyFileAnalysis"`
	// 是否开启脚本引擎注入检测（0-关闭，1-开启）

	IsScriptEngineInjectionAnalysis *int64 `json:"IsScriptEngineInjectionAnalysis,omitempty" name:"IsScriptEngineInjectionAnalysis"`
	// 是否开启表达式注入检测（0-关闭，1-开启）

	IsExpressionInjectionAnalysis *int64 `json:"IsExpressionInjectionAnalysis,omitempty" name:"IsExpressionInjectionAnalysis"`
	// URL收敛阈值

	UrlConvergenceThreshold *int64 `json:"UrlConvergenceThreshold,omitempty" name:"UrlConvergenceThreshold"`
	// 是否开启JNDI注入检测（0-关闭，1-开启）

	IsJNDIInjectionAnalysis *int64 `json:"IsJNDIInjectionAnalysis,omitempty" name:"IsJNDIInjectionAnalysis"`
	// 接口自动收敛开关,0&nbsp;关&nbsp;|&nbsp;1&nbsp;开

	UrlAutoConvergenceEnable *bool `json:"UrlAutoConvergenceEnable,omitempty" name:"UrlAutoConvergenceEnable"`
	// 日志来源&nbsp;CLS&nbsp;|&nbsp;ES

	LogSource *string `json:"LogSource,omitempty" name:"LogSource"`
	// 需过滤的接口

	IgnoreOperationName *string `json:"IgnoreOperationName,omitempty" name:"IgnoreOperationName"`
	// 应用是否开启dashboard配置：&nbsp;false&nbsp;关（与业务系统保持一致）/true&nbsp;开（应用级配置）

	EnableDashboardConfig *bool `json:"EnableDashboardConfig,omitempty" name:"EnableDashboardConfig"`
	// 是否开启远程命令检测

	IsRemoteCommandExecutionAnalysis *int64 `json:"IsRemoteCommandExecutionAnalysis,omitempty" name:"IsRemoteCommandExecutionAnalysis"`
	// 是否开启删除任意文件检测（0-关闭，1-开启）

	IsDeleteAnyFileAnalysis *int64 `json:"IsDeleteAnyFileAnalysis,omitempty" name:"IsDeleteAnyFileAnalysis"`
	// 是否开启Webshell后门检测（0-关闭，1-开启）

	IsWebshellBackdoorAnalysis *int64 `json:"IsWebshellBackdoorAnalysis,omitempty" name:"IsWebshellBackdoorAnalysis"`
	// URL长分段收敛阈值

	UrlLongSegmentThreshold *int64 `json:"UrlLongSegmentThreshold,omitempty" name:"UrlLongSegmentThreshold"`
	// 是否开启线程剖析

	EnableSnapshot *bool `json:"EnableSnapshot,omitempty" name:"EnableSnapshot"`
	// 组件列表

	InstrumentList []*Instrument `json:"InstrumentList,omitempty" name:"InstrumentList"`
	// dashboard&nbsp;ID

	DashboardTopicID *string `json:"DashboardTopicID,omitempty" name:"DashboardTopicID"`
	// 是否开启上传任意文件检测（0-关闭，1-开启）

	IsUploadAnyFileAnalysis *int64 `json:"IsUploadAnyFileAnalysis,omitempty" name:"IsUploadAnyFileAnalysis"`
	// 是否开启目录遍历检测（0-关闭，1-开启）

	IsDirectoryTraversalAnalysis *int64 `json:"IsDirectoryTraversalAnalysis,omitempty" name:"IsDirectoryTraversalAnalysis"`
	// 是否开启JNI注入检测（0-关闭，1-开启）

	IsJNIInjectionAnalysis *int64 `json:"IsJNIInjectionAnalysis,omitempty" name:"IsJNIInjectionAnalysis"`
}

func (r *ModifyApmApplicationConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmApplicationConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateApmInstanceRequest struct {
	*tchttp.BaseRequest

	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *TerminateApmInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateApmInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateApmServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateApmServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateApmServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChartData struct {

	// 接口参数id列表

	Series []*Series `json:"Series,omitempty" name:"Series"`
	// 接口参数列表

	YAxis []*YAxis `json:"YAxis,omitempty" name:"YAxis"`
	// 接口参数列表

	YAxisEx *YAxis `json:"YAxisEx,omitempty" name:"YAxisEx"`
}

type OrderBy struct {

	// 需要排序的字段，现支持&nbsp;startTIme,&nbsp;endTime,&nbsp;duration

	Key *string `json:"Key,omitempty" name:"Key"`
	// asc&nbsp;顺序排序&nbsp;/&nbsp;desc&nbsp;倒序排序

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Instrument struct {

	// 组件名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件开关

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
}

type DescribeAgentInfoRequest struct {
	*tchttp.BaseRequest

	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 应用过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 开始时间

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeAgentInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceBrief struct {

	// 应用ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
}

type APMKVItem struct {

	// Key&nbsp;值定义

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value&nbsp;值定义

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Span struct {

	// 产生时间戳(毫秒)

	Timestamp *int64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 产生时间戳(微秒)

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 持续耗时(微妙)

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// 产生时间戳(毫秒)

	StartTimeMillis *int64 `json:"StartTimeMillis,omitempty" name:"StartTimeMillis"`
	// Parent&nbsp;Span&nbsp;ID

	ParentSpanID *string `json:"ParentSpanID,omitempty" name:"ParentSpanID"`
	// Trace&nbsp;ID

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
	// 标签

	Tags []*SpanTag `json:"Tags,omitempty" name:"Tags"`
	// Span&nbsp;名称

	OperationName *string `json:"OperationName,omitempty" name:"OperationName"`
	// 关联关系

	References []*SpanReference `json:"References,omitempty" name:"References"`
	// Span&nbsp;ID

	SpanID *string `json:"SpanID,omitempty" name:"SpanID"`
	// 日志

	Logs []*SpanLog `json:"Logs,omitempty" name:"Logs"`
	// 上报应用服务信息

	Process *SpanProcess `json:"Process,omitempty" name:"Process"`
}

type DescribeGeneralMetricDataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标结果集

		Records []*Line `json:"Records,omitempty" name:"Records"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralMetricDataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralMetricDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type APMKV struct {

	// Key&nbsp;值定义

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value&nbsp;值定义

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type TraceOverview struct {

	// TraceId查询

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
	// 错误信息

	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
	// 服务名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 耗时

	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
	// 接口名

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// traceID调用状态

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type QueryView struct {

	// 查询视图描述/名称

	Description *string `json:"Description,omitempty" name:"Description"`
	// 查询条件集合

	Filters *string `json:"Filters,omitempty" name:"Filters"`
	// 查询视图&nbsp;ID

	ViewId *string `json:"ViewId,omitempty" name:"ViewId"`
}

type Snapshot struct {

	// 节点id

	ID *string `json:"ID,omitempty" name:"ID"`
	// 节点位置信息

	Position *Position `json:"Position,omitempty" name:"Position"`
}

type DescribeApmServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 页码

		Page *int64 `json:"Page,omitempty" name:"Page"`
		// 页大小

		PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
		// 应用详情列表

		ServiceList []*ServiceDetail `json:"ServiceList,omitempty" name:"ServiceList"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyGeneralApmApplicationConfigRequest struct {
	*tchttp.BaseRequest

	// 需要修改的字段key&nbsp;value分别指定字段名、字段值
	// [具体字段请见](https://cloud.{{conf.main.domain}}/document/product/248/111241)

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 需要修改配置的应用列表名称

	ServiceNames []*string `json:"ServiceNames,omitempty" name:"ServiceNames"`
	// 业务系统Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyGeneralApmApplicationConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyGeneralApmApplicationConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentOperationConfigView struct {

	// 当前接口配置是否开启了接口白名单配置

	RetentionValid *bool `json:"RetentionValid,omitempty" name:"RetentionValid"`
	// RetentionValid为false时生效，接口配置中的黑名单配置，配置中的接口不采集

	IgnoreOperation *string `json:"IgnoreOperation,omitempty" name:"IgnoreOperation"`
	// RetentionValid为true时生效，接口配置中的白名单配置，仅采集配置中的接口

	RetentionOperation *string `json:"RetentionOperation,omitempty" name:"RetentionOperation"`
}

type DescribeGeneralApmApplicationConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 应用配置项

		ApmApplicationConfigView *ApmApplicationConfigView `json:"ApmApplicationConfigView,omitempty" name:"ApmApplicationConfigView"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGeneralApmApplicationConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralApmApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// APM&nbsp;业务系统列表

		Instances []*ApmInstanceDetail `json:"Instances,omitempty" name:"Instances"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeApmInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmInfoByAppIdRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeApmInfoByAppIdRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmInfoByAppIdRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopologyNewRequest struct {
	*tchttp.BaseRequest

	// 不显示的节点类型

	Hidden *Selectors `json:"Hidden,omitempty" name:"Hidden"`
	// 查询top5慢响应节点

	IsSlowTopFive *bool `json:"IsSlowTopFive,omitempty" name:"IsSlowTopFive"`
	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 下游层级

	DownLevel *int64 `json:"DownLevel,omitempty" name:"DownLevel"`
	// 应用实例信息

	ServiceInstance *string `json:"ServiceInstance,omitempty" name:"ServiceInstance"`
	// 表示Topic（MQ拓扑图用）

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// 视图筛选列表

	Selectors *Selectors `json:"Selectors,omitempty" name:"Selectors"`
	// TraceID

	TraceID *string `json:"TraceID,omitempty" name:"TraceID"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询开始时间

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 查询结束时间

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 上游层级

	UpLevel *int64 `json:"UpLevel,omitempty" name:"UpLevel"`
	// 视角

	View *string `json:"View,omitempty" name:"View"`
	// 过滤器

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 是否获取资源层信息

	GetResource *bool `json:"GetResource,omitempty" name:"GetResource"`
	// 视图ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 根据应用标签过滤

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
}

func (r *DescribeTopologyNewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopologyNewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeApmServiceRequest struct {
	*tchttp.BaseRequest

	// 应用名

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 页码

	Page *int64 `json:"Page,omitempty" name:"Page"`
	// 页大小

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 业务系统ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否demo模式

	Demo *bool `json:"Demo,omitempty" name:"Demo"`
	// 标签

	Tags []*ApmTag `json:"Tags,omitempty" name:"Tags"`
	// 应用ID

	ServiceID *string `json:"ServiceID,omitempty" name:"ServiceID"`
}

func (r *DescribeApmServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeApmServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExpiredApplicationConfigRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeleteExpiredApplicationConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExpiredApplicationConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateApmInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateApmInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateApmInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateQueryViewRequest struct {
	*tchttp.BaseRequest

	// 是否为拓扑图视图，默认为&nbsp;false

	IsTopo *bool `json:"IsTopo,omitempty" name:"IsTopo"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 视图描述/名称

	Description *string `json:"Description,omitempty" name:"Description"`
	// 查询条件集合

	Filters *string `json:"Filters,omitempty" name:"Filters"`
}

func (r *CreateQueryViewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateQueryViewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExpiredApplicationConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExpiredApplicationConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExpiredApplicationConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyApmServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyApmServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyApmServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeGeneralMetricDataRequest struct {
	*tchttp.BaseRequest

	// 视图名称，不可自定义输入。[详情请见。](https://cloud.{{conf.main.domain}}/document/product/248/101681)

	ViewName *string `json:"ViewName,omitempty" name:"ViewName"`
	// 要过滤的维度信息，不同视图有对应的指标维度，[详情请见。](https://cloud.{{conf.main.domain}}/document/product/248/101681)

	Filters []*GeneralFilter `json:"Filters,omitempty" name:"Filters"`
	// 结束时间的时间戳，支持查询30天内的指标数据。（单位：秒）

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 需要查询的指标名称，不可自定义输入，[详情请见。](https://cloud.{{conf.main.domain}}/document/product/248/101681)

	Metrics []*string `json:"Metrics,omitempty" name:"Metrics"`
	// 业务系统&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否按固定时间跨度聚合，填入1及大于1的值按1处理，不填按0处理。
	// -&nbsp;填入0，则计算开始时间到截止时间的指标数据。
	// -&nbsp;填入1，则会按照开始时间到截止时间的时间跨度选择聚合粒度：
	// &nbsp;-&nbsp;时间跨度&nbsp;(0,12)&nbsp;小时，则按一分钟粒度聚合。
	// &nbsp;-&nbsp;时间跨度&nbsp;[12,48]&nbsp;小时，则按五分钟粒度聚合。
	// &nbsp;-&nbsp;时间跨度&nbsp;(48,&nbsp;+∞)&nbsp;小时，则按一小时粒度聚合。

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 对查询指标进行排序：
	// Key&nbsp;填写云&nbsp;API&nbsp;指标名称，[详情请见。](https://cloud.{{conf.main.domain}}/document/product/248/101681)
	// Value&nbsp;填写排序方式：&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
	// -&nbsp;asc：对查询指标进行升序排序
	// -&nbsp;desc：对查询指标进行降序排序

	OrderBy *OrderBy `json:"OrderBy,omitempty" name:"OrderBy"`
	// 查询指标的限制条数，目前最多展示50条数据，PageSize取值为1-50，上送PageSize则根据PageSize的值展示限制条数。

	PageSize *int64 `json:"PageSize,omitempty" name:"PageSize"`
	// 聚合维度，不同视图有对应的指标维度，[详情请见。](https://cloud.{{conf.main.domain}}/document/product/248/101681)

	GroupBy []*string `json:"GroupBy,omitempty" name:"GroupBy"`
	// 起始时间的时间戳，支持查询30天内的指标数据。（单位：秒）

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
}

func (r *DescribeGeneralMetricDataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGeneralMetricDataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

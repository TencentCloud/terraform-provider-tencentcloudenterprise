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
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type LogsetETL struct {

	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志集下日志主题的个数

	TopicCount *int64 `json:"TopicCount,omitempty" name:"TopicCount"`
	// 创建时间（秒级时间戳）

	CreatedAt *int64 `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type CreateExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出ID。

		ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigExtraResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigExtraResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKafkaRechargeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteKafkaRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParquetKeyInfo struct {

	// 键值名称

	KeyName *string `json:"KeyName,omitempty" name:"KeyName"`
	// 数据类型，目前支持6种类型：string、boolean、int32、int64、float、double

	KeyType *string `json:"KeyType,omitempty" name:"KeyType"`
	// 解析失败赋值信息

	KeyNonExistingField *string `json:"KeyNonExistingField,omitempty" name:"KeyNonExistingField"`
}

type RecordingRuleYamlTaskInfo struct {

	// yaml配置文件id

	YamlId *string `json:"YamlId,omitempty" name:"YamlId"`
	// 源日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 写入描述的日志主题id

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 任务状态，1:运行&nbsp;2:停止&nbsp;3:异常-找不到源日志主题&nbsp;4:异常-找不到目标主题&nbsp;5:&nbsp;访问权限问题&nbsp;6:内部故障&nbsp;7:其他故障

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务启用状态，1开启,&nbsp;2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 调度开始时间

	ProcessStartTime *uint64 `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 调度周期(分钟)

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 执行周期单位,&nbsp;`0`:&nbsp;Minute,&nbsp;`1`:&nbsp;Second。&nbsp;-&nbsp;ProcessPeriodUnit:1时，ProcessPeriod&nbsp;只支持&nbsp;30（即只支持30秒）。&nbsp;-&nbsp;ProcessPeriodUnit:0时，ProcessPeriod&nbsp;支持范围(0,1440]分钟。

	ProcessPeriodUnit *uint64 `json:"ProcessPeriodUnit,omitempty" name:"ProcessPeriodUnit"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
	// yaml配置文件名称

	YamlConfigName *string `json:"YamlConfigName,omitempty" name:"YamlConfigName"`
	// yaml配置文件内容

	YamlContent *string `json:"YamlContent,omitempty" name:"YamlContent"`
	// yaml文件子任务数量

	SubTaskCount *int64 `json:"SubTaskCount,omitempty" name:"SubTaskCount"`
}

type ApplyConfigToMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyConfigToMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyConfigToMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalDataSourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量大小

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 外部数据源信息

		ExternalDataSources []*ExternalDataSourceInfo `json:"ExternalDataSources,omitempty" name:"ExternalDataSources"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExternalDataSourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExternalDataSourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyShipperResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyShipperResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EventLog struct {

	// 事件通道，支持Application，Security，Setup，System，ALL

	EventChannel *string `json:"EventChannel,omitempty" name:"EventChannel"`
	// 时间类型，1:用户自定义，2:当前时间

	TimeType *uint64 `json:"TimeType,omitempty" name:"TimeType"`
	// 时间，用户选择自定义时间类型时，需要指定时间

	Timestamp *uint64 `json:"Timestamp,omitempty" name:"Timestamp"`
	// 事件ID过滤列表&nbsp;&nbsp;选填，为空表示不做过滤&nbsp;支持正向过滤单个值（例：20）或范围（例：0-20），也支持反向过滤单个值(例：-20)&nbsp;多个过滤项之间可由逗号隔开，例：1-200,-100表示采集1-200范围内除了100以外的事件日志

	EventIDs []*string `json:"EventIDs,omitempty" name:"EventIDs"`
}

type InstanceData struct {

	// 云监控指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// CLS指标名称

	CLSMetricName *string `json:"CLSMetricName,omitempty" name:"CLSMetricName"`
	// 云产品命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 实例信息

	Dimensions []*Dimension `json:"Dimensions,omitempty" name:"Dimensions"`
	// 周期,单位：秒

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 指标统计值

	Value *float64 `json:"Value,omitempty" name:"Value"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type DescribeFunctionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可加工的函数列表

		FunctionInfos []*FunctionInfo `json:"FunctionInfos,omitempty" name:"FunctionInfos"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFunctionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFunctionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetricLabelValuesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时序metric&nbsp;label&nbsp;values

		Values []*string `json:"Values,omitempty" name:"Values"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMetricLabelValuesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricLabelValuesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogConfigInfo struct {

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 采集日志路径列表

	Path *string `json:"Path,omitempty" name:"Path"`
	// 日志类型，可选值：minimalist_log代表单行全文；multiline_log代表多行全文；fullregex_log代表单行完全正则；json_log代表JSON日志；multiline_fullregex_log多行完全正则；user_define_log代表组合解析；delimiter_log代表分隔符；

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 日志格式化格式，可选值：default代表默认格式；nginx_log代表nginx格式日志；

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 黑名单path列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 用户自定义解析字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// 采集配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 使用了元数据的机器组ID列表

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志输入类型，支持linux_file、window_event&nbsp;、windows_file、syslog、k8s_stdout、k8s_file

	InputType *string `json:"InputType,omitempty" name:"InputType"`
}

type CheckAlarmRuleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 监控对象测试结果

		AlarmRuleTestResults []*AlarmRuleTestResult `json:"AlarmRuleTestResults,omitempty" name:"AlarmRuleTestResults"`
		// 触发条件测试结果

		ConditionTestResult *ConditionTestResult `json:"ConditionTestResult,omitempty" name:"ConditionTestResult"`
		// 是否成功检测，0表示成功，-1表示失败

		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
		// 错误原因

		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
		// 多触发条件测试结果。code&nbsp;:&nbsp;0&nbsp;为成功，其他都是失败。

		MultiConditionTestResult []*ConditionTestResult `json:"MultiConditionTestResult,omitempty" name:"MultiConditionTestResult"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAlarmRuleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAlarmRuleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckUserIDRequest struct {
	*tchttp.BaseRequest

	// 用户ID

	UserID *string `json:"UserID,omitempty" name:"UserID"`
	// 用户ID类型，目前只支持topic和logset

	IDType *string `json:"IDType,omitempty" name:"IDType"`
}

func (r *CheckUserIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckUserIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebCallbacksRequest struct {
	*tchttp.BaseRequest

	// name&nbsp;按照【告警渠道回调配置名称】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;webCallbackId&nbsp;按照【告警渠道回调配置ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;type&nbsp;按照【告警渠道回调配置渠道类型】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeWebCallbacksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebCallbacksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type DeleteWebCallbackRequest struct {
	*tchttp.BaseRequest

	// 告警渠道回调配置ID。

	WebCallbackId *string `json:"WebCallbackId,omitempty" name:"WebCallbackId"`
}

func (r *DeleteWebCallbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebCallbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalDataSourcePreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// COS文件(CSV格式)部分行信息

		Data []*CosFieldsData `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExternalDataSourcePreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExternalDataSourcePreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordingRuleYamlTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RecordingRule任务列表信息

		RecordingRuleYamlTaskInfos []*RecordingRuleYamlTaskInfo `json:"RecordingRuleYamlTaskInfos,omitempty" name:"RecordingRuleYamlTaskInfos"`
		// 任务总次数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordingRuleYamlTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordingRuleYamlTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmTemplateConfig struct {

	// 用产品类型。&nbsp;TKE：容器服务&nbsp;CLB：负载均衡

	CloudProduct *string `json:"CloudProduct,omitempty" name:"CloudProduct"`
	// 告警模板&nbsp;id。

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 实例&nbsp;id。当&nbsp;CloudProduct&nbsp;为&nbsp;TKE&nbsp;时，InstanceId&nbsp;为&nbsp;集群&nbsp;id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 告警模板额外信息

	ExtraData *string `json:"ExtraData,omitempty" name:"ExtraData"`
}

type CreateIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaConsumerTopicsRequest struct {
	*tchttp.BaseRequest

	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否需要消费组信息

	NeedGroupInfo *bool `json:"NeedGroupInfo,omitempty" name:"NeedGroupInfo"`
	// 日志主题ID

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *DescribeKafkaConsumerTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumerTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentUpdateInfo struct {

	// 是否需要升级。true表示需要升级，false表示不需要升级。

	NeedUpdate *bool `json:"NeedUpdate,omitempty" name:"NeedUpdate"`
	// 升级类型:0-null,1-manual,2-auto,3-force

	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 升级动作:0-null,1-update,2-revert

	UpdateAction *int64 `json:"UpdateAction,omitempty" name:"UpdateAction"`
	// 重试次数,最大3次

	RetryCount *int64 `json:"RetryCount,omitempty" name:"RetryCount"`
	// 目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
	// 安装包下载链接1

	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`
	// 安装包下载链接2

	DownloadUrlSecond *string `json:"DownloadUrlSecond,omitempty" name:"DownloadUrlSecond"`
	// 安装包文件MD5值

	FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
}

type ConditionTestResult struct {

	// 错误码，可选值：1005代表触发条件表达式语法错误；1006代表触发条件计算错误，请检查参数

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type CreateRemoteWriteTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 目标服务名称

	Target *string `json:"Target,omitempty" name:"Target"`
	// 目标地址

	RemoteWriteURL *string `json:"RemoteWriteURL,omitempty" name:"RemoteWriteURL"`
	// 鉴权类型&nbsp;0:&nbsp;无鉴权&nbsp;1:&nbsp;basic_auth&nbsp;2:&nbsp;token

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 网络类型：&nbsp;1&nbsp;内网&nbsp;2外网

	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 鉴权信息

	AuthInfo *RemoteWriteAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
	// 子网

	SubNet *string `json:"SubNet,omitempty" name:"SubNet"`
	// 后端服务类型&nbsp;0&nbsp;CVM&nbsp;1025&nbsp;CLB

	VirtualGatewayType *int64 `json:"VirtualGatewayType,omitempty" name:"VirtualGatewayType"`
	// 云时序数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否开启投递服务日志。1：关闭，2：开启。&nbsp;默认值：2

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *CreateRemoteWriteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteWriteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递任务是否生效，默认不生效

	Effective *bool `json:"Effective,omitempty" name:"Effective"`
	// 是否投递日志的元数据信息，默认为&nbsp;true。&nbsp;当NeedContent为true时：字段Content有效。&nbsp;当NeedContent为false时：字段Content无效。

	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`
	// 如果需要投递元数据信息，元数据信息的描述

	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`
	// CKafka的描述

	Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`
	// 投递时压缩方式，取值0，2，3。[0：NONE；2：SNAPPY；3：LZ4]

	Compression *int64 `json:"Compression,omitempty" name:"Compression"`
	// 角色访问描述名

	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`
	// 外部ID

	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
	// 高级配置

	AdvancedConfig *AdvancedConsumerConfiguration `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
}

func (r *ModifyConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendConsumerHeartbeatRequest struct {
	*tchttp.BaseRequest

	// 上报心跳的消费组标识

	ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
	// 上报心跳的消费者名称&nbsp;（字母数字下划线，不允许数字、_开头，&nbsp;长度小于256）

	Consumer *string `json:"Consumer,omitempty" name:"Consumer"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// topic&nbsp;分区信息

	TopicPartitionsInfo []*TopicPartitionInfo `json:"TopicPartitionsInfo,omitempty" name:"TopicPartitionsInfo"`
	// 消费者组分区策略

	PartitionStrategy *uint64 `json:"PartitionStrategy,omitempty" name:"PartitionStrategy"`
}

func (r *SendConsumerHeartbeatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendConsumerHeartbeatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformPreviewInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 1:&nbsp;任务完成，&nbsp;2:&nbsp;任务处理中，3:&nbsp;任务处理失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// 错误信息

		FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformPreviewInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformPreviewInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyRegexInfo struct {

	// 需要过滤日志的key

	Key *string `json:"Key,omitempty" name:"Key"`
	// key对应的过滤规则regex

	Regex *string `json:"Regex,omitempty" name:"Regex"`
}

type TopicIdLogFailureInfo struct {

	// 日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志失败信息

	LogFailureInfos []*DataTransformFailureInfo `json:"LogFailureInfos,omitempty" name:"LogFailureInfos"`
}

type DeleteTopicExtendConfigRequest struct {
	*tchttp.BaseRequest

	// clb的业务配置

	LbKeys []*string `json:"LbKeys,omitempty" name:"LbKeys"`
}

func (r *DeleteTopicExtendConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicExtendConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicBaseMetricConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 指标采集配置列表

		Datas []*BaseMetricCollectConfig `json:"Datas,omitempty" name:"Datas"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicBaseMetricConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicBaseMetricConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnalysisDimensional struct {

	// 分析名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分析类型：query，field&nbsp;，original

	Type *string `json:"Type,omitempty" name:"Type"`
	// 分析内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 多维分析配置。&nbsp;&nbsp;当Analysis的Type字段为query（自定义）时，支持&nbsp;{&nbsp;"Key":&nbsp;"SyntaxRule",&nbsp;//&nbsp;语法规则&nbsp;"Value":&nbsp;"1"&nbsp;//0：Lucene语法&nbsp;，1：&nbsp;CQL语法&nbsp;}&nbsp;&nbsp;当Analysis的Type字段为field（top5）时,&nbsp;支持&nbsp;{&nbsp;"Key":&nbsp;"QueryIndex",&nbsp;"Value":&nbsp;"-1"&nbsp;//&nbsp;-1：自定义，&nbsp;1：执行语句1，&nbsp;2：执行语句2&nbsp;},{&nbsp;"Key":&nbsp;"CustomQuery",&nbsp;//检索语句。&nbsp;QueryIndex为-1时有效且必填&nbsp;"Value":&nbsp;"*&nbsp;|&nbsp;select&nbsp;count(*)&nbsp;as&nbsp;count"&nbsp;},{&nbsp;"Key":&nbsp;"SyntaxRule",&nbsp;//&nbsp;查不到这个字段也是老语法（Lucene）&nbsp;"Value":&nbsp;"0"//0:Lucene,&nbsp;1:CQL&nbsp;}&nbsp;&nbsp;当Analysis的Type字段为original（原始日志）时,&nbsp;支持&nbsp;{&nbsp;"Key":&nbsp;"Fields",&nbsp;"Value":&nbsp;"SOURCE,HOSTNAME,TIMESTAMP,PKG_LOGID,TAG.pod_ip"&nbsp;},&nbsp;{&nbsp;"Key":&nbsp;"QueryIndex",&nbsp;"Value":&nbsp;"-1"&nbsp;//&nbsp;-1：自定义，&nbsp;1：执行语句1，&nbsp;2：执行语句2&nbsp;},{&nbsp;"Key":&nbsp;"CustomQuery",&nbsp;//&nbsp;//检索语句。&nbsp;QueryIndex为-1时有效且必填&nbsp;"Value":&nbsp;"*&nbsp;|&nbsp;select&nbsp;count(*)&nbsp;as&nbsp;count"&nbsp;},{&nbsp;"Key":&nbsp;"Format",&nbsp;//显示形式。1：每条日志一行，2：每条日志每个字段一行&nbsp;"Value":&nbsp;"2"&nbsp;},&nbsp;{&nbsp;"Key":&nbsp;"Limit",&nbsp;//最大日志条数&nbsp;"Value":&nbsp;"5"&nbsp;},{&nbsp;"Key":&nbsp;"SyntaxRule",&nbsp;//&nbsp;查不到这个字段也是老语法&nbsp;"Value":&nbsp;"0"//0:Lucene,&nbsp;1:CQL&nbsp;}

	ConfigInfo []*AlarmAnalysisConfig `json:"ConfigInfo,omitempty" name:"ConfigInfo"`
}

type MetricLabel struct {

	// 指标名称

	Key *string `json:"Key,omitempty" name:"Key"`
	// 指标内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CheckRemoteWriteTaskConnectResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 联通行校验&nbsp;0&nbsp;正常&nbsp;1不正常

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckRemoteWriteTaskConnectResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRemoteWriteTaskConnectResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogSubscribesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// binlog采集配置信息

		Datas []*BinlogConfigInfo `json:"Datas,omitempty" name:"Datas"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogSubscribesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBinlogSubscribesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCosRechargeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyCosRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommitConsumerOffsetsRequest struct {
	*tchttp.BaseRequest

	// 消费组标识

	ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
	// 消费机器名称

	Consumer *string `json:"Consumer,omitempty" name:"Consumer"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// topic分区点位信息

	TopicPartitionOffsetsInfo []*TopicPartitionOffsetInfo `json:"TopicPartitionOffsetsInfo,omitempty" name:"TopicPartitionOffsetsInfo"`
}

func (r *CommitConsumerOffsetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CommitConsumerOffsetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigMachineGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集规则配置绑定的机器组列表

		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigMachineGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogHistogramRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 时间间隔:&nbsp;单位ms&nbsp;限制性条件：(To-From)&nbsp;/&nbsp;interval&nbsp;<=&nbsp;200

	Interval *int64 `json:"Interval,omitempty" name:"Interval"`
	// 0（默认值）：不执行语法优化；1：执行语法优化。

	QueryOptimize *uint64 `json:"QueryOptimize,omitempty" name:"QueryOptimize"`
	// 检索语法规则，默认值为0。&nbsp;0：Lucene语法，1：CQL语法。&nbsp;详细说明参见检索条件语法规则

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 要检索分析的日志主题ID列表。TopicId和TopicIds二个参数只能选择其中一个参数进行查询。

	TopicIds []*string `json:"TopicIds,omitempty" name:"TopicIds"`
}

func (r *DescribeLogHistogramRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogHistogramRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIdleResourcePolicyRequest struct {
	*tchttp.BaseRequest

	// 自动冻结策略&nbsp;1&nbsp;开启&nbsp;0关闭

	AutoFreeze *uint64 `json:"AutoFreeze,omitempty" name:"AutoFreeze"`
}

func (r *ModifyIdleResourcePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIdleResourcePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 监控对象列表。

	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`
	// 监控任务运行时间点。

	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`
	// 触发条件&nbsp;注意:&nbsp;&nbsp;Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为2000。

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 告警重复的周期。单位是分钟。取值范围是0~1440。

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 关联的告警通知模板列表。

	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`
	// 是否开启告警策略。默认值为true，true表示开启告警策略，false表示关闭告警策略。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 该参数已废弃，请使用Status参数控制是否开启告警策略。

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 用户自定义告警内容

	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`
	// 用户自定义回调

	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`
	// 多维分析

	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
	// 分组触发状态。&nbsp;默认值false

	GroupTriggerStatus *bool `json:"GroupTriggerStatus,omitempty" name:"GroupTriggerStatus"`
	// 分组触发条件。

	GroupTriggerCondition []*string `json:"GroupTriggerCondition,omitempty" name:"GroupTriggerCondition"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的告警策略。&nbsp;&nbsp;最大支持10个标签键值对，并且不能有重复的键值对。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 监控对象类型。0:执行语句共用监控对象;&nbsp;1:每个执行语句单独选择监控对象。&nbsp;不填则默认为0。&nbsp;当值为1时，AlarmTargets元素个数不能超过10个，AlarmTargets中的Number必须是从1开始的连续正整数，不能重复。

	MonitorObjectType *uint64 `json:"MonitorObjectType,omitempty" name:"MonitorObjectType"`
	// 告警附加分类信息列表。&nbsp;Classifications元素个数不能超过20个。&nbsp;Classifications元素的Key不能为空，不能重复，长度不能超过50个字符，符合正则&nbsp;^[a-z]([a-z0-9_]{0,49})$。&nbsp;Classifications元素的Value长度不能超过200个字符。

	Classifications []*AlarmClassification `json:"Classifications,omitempty" name:"Classifications"`
	// 告警级别&nbsp;0:警告(Warn);&nbsp;1:提醒(Info);&nbsp;2:紧急&nbsp;(Critical)。&nbsp;注意:&nbsp;&nbsp;不填则默认为0。&nbsp;Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。

	AlarmLevel *uint64 `json:"AlarmLevel,omitempty" name:"AlarmLevel"`
	// 交互式触发配置信息。

	ConditionInteractiveConfig *string `json:"ConditionInteractiveConfig,omitempty" name:"ConditionInteractiveConfig"`
	// 多触发条件&nbsp;注意:&nbsp;&nbsp;Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。

	MultiConditions []*MultiCondition `json:"MultiConditions,omitempty" name:"MultiConditions"`
	// 告警模板相关信息

	AlarmTemplateInfo *AlarmTemplateConfig `json:"AlarmTemplateInfo,omitempty" name:"AlarmTemplateInfo"`
}

func (r *CreateAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperTasksRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 查询的开始时间戳，支持最近3天的查询，&nbsp;毫秒。&nbsp;StartTime必须小于EndTime

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 查询的结束时间戳，&nbsp;毫秒。&nbsp;StartTime必须小于EndTime

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeShipperTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShipperTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS对应的topic标识

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *CloseKafkaConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKafkaConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetConfigurationTemplateApplyLogRequest struct {
	*tchttp.BaseRequest

	// 要查询的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 单次查询返回的条数，最大值为1000

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 加载更多详情时使用，透传上次返回的Context值，获取后续的信息

	Context *string `json:"Context,omitempty" name:"Context"`
	// 是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为&nbsp;desc

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 如果Query包含SQL语句，UseNewAnalysis为true时响应参数AnalysisRecords和Columns有效，&nbsp;UseNewAnalysis为false时响应参数AnalysisResults和ColNames有效

	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
}

func (r *GetConfigurationTemplateApplyLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetConfigurationTemplateApplyLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceAndFolderRelationRequest struct {
	*tchttp.BaseRequest

	// 仪表盘Id列表

	DashboardIds []*string `json:"DashboardIds,omitempty" name:"DashboardIds"`
	// 文件夹Id

	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *ModifyResourceAndFolderRelationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceAndFolderRelationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenClsServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *OpenClsServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenClsServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicIdAndRegion struct {

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题id所在的地域id。

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
}

type CreateConfigRequest struct {
	*tchttp.BaseRequest

	// 采集配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 日志采集路径，包含文件名，支持多个路径，多个路径之间英文逗号分隔，文件采集情况下必填

	Path *string `json:"Path,omitempty" name:"Path"`
	// 采集的日志类型，默认为minimalist_log。支持以下类型：&nbsp;&nbsp;json_log代表：JSON-文件日志（详见使用&nbsp;JSON&nbsp;提取模式采集日志）；&nbsp;delimiter_log代表：分隔符-文件日志（详见使用分隔符提取模式采集日志）；&nbsp;minimalist_log代表：单行全文-文件日志（详见使用单行全文提取模式采集日志）；&nbsp;fullregex_log代表：单行完全正则-文件日志（详见使用单行-完全正则提取模式采集日志）；&nbsp;multiline_log代表：多行全文-文件日志（详见使用多行全文提取模式采集日志）；&nbsp;multiline_fullregex_log代表：多行完全正则-文件日志（详见使用多行-完全正则提取模式采集日志）；&nbsp;user_define_log代表：组合解析（适用于多格式嵌套的日志，详见使用组合解析提取模式采集日志）；&nbsp;service_syslog代表：syslog&nbsp;采集（详见采集&nbsp;Syslog）；&nbsp;windows_event_log代表：Windows事件日志（详见采集&nbsp;Windows&nbsp;事件日志）。

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 采集配置所属日志主题ID即TopicId

	Output *string `json:"Output,omitempty" name:"Output"`
	// 用户自定义采集规则，Json格式序列化的字符串。当LogType为user_define_log时，必填。

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// config_extra表主键ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集配置标签

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 高级采集配置。&nbsp;Json字符串，&nbsp;Key/Value定义为如下：&nbsp;&nbsp;ClsAgentFileTimeout(超时属性),&nbsp;取值范围:&nbsp;大于等于0的整数，&nbsp;0为不超时&nbsp;ClsAgentMaxDepth(最大目录深度)，取值范围:&nbsp;大于等于0的整数&nbsp;ClsAgentParseFailMerge(合并解析失败日志)，取值范围:&nbsp;true或false&nbsp;样例：&nbsp;{\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}&nbsp;控制台默认占位值：{\"ClsAgentDefault\":0}

	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
	// 采集配置来源，0：默认来源，1:&nbsp;TKE

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 日志输入类型，支持file、window_event、syslog、k8s_stdout、k8s_file

	InputType *string `json:"InputType,omitempty" name:"InputType"`
}

func (r *CreateConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchCosRechargeInfoRequest struct {
	*tchttp.BaseRequest

	// 日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 投递任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// COS存储桶，详见产品支持的存储桶命名规范。

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// COS存储桶所在地域，详见产品支持的地域列表。

	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
	// COS文件所在文件夹的前缀。默认为空，投递存储桶下所有的文件。

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 压缩模式:&nbsp;"",&nbsp;"gzip",&nbsp;"lzop",&nbsp;"snappy";&nbsp;默认""

	Compress *string `json:"Compress,omitempty" name:"Compress"`
}

func (r *SearchCosRechargeInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCosRechargeInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNoticeContentRequest struct {
	*tchttp.BaseRequest

	// 通知内容模板ID

	NoticeContentId *string `json:"NoticeContentId,omitempty" name:"NoticeContentId"`
}

func (r *DeleteNoticeContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNoticeContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicPartitionOffsetInfo struct {

	// 日志主题id

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 分区点位信息

	PartitionOffsets []*PartitionOffsetInfo `json:"PartitionOffsets,omitempty" name:"PartitionOffsets"`
}

type ModifyConfigurationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigurationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigurationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExcludePathInfo struct {

	// 类型，选填File或Path

	Type *string `json:"Type,omitempty" name:"Type"`
	// Type对应的具体内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type WebCallbackInfo struct {

	// 告警渠道回调配置id。

	WebCallbackId *string `json:"WebCallbackId,omitempty" name:"WebCallbackId"`
	// 告警渠道回调配置名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 渠道类型&nbsp;&nbsp;WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调;

	Type *string `json:"Type,omitempty" name:"Type"`
	// 回调地址

	Webhook *string `json:"Webhook,omitempty" name:"Webhook"`
	// 请求方式。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 秘钥信息。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 主账号。

	Uin *string `json:"Uin,omitempty" name:"Uin"`
	// 子账号。

	SubUin *string `json:"SubUin,omitempty" name:"SubUin"`
	// 创建时间。秒级时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间。秒级时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateWebCallbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回调配置ID。

		WebCallbackId *string `json:"WebCallbackId,omitempty" name:"WebCallbackId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWebCallbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWebCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest

	// topicName&nbsp;按照【日志主题名称】进行过滤，默认为模糊匹配，可使用&nbsp;PreciseSearch&nbsp;参数设置为精确匹配。类型：String。必选：否&nbsp;logsetName&nbsp;按照【日志集名称】进行过滤，默认为模糊匹配，可使用&nbsp;PreciseSearch&nbsp;参数设置为精确匹配。类型：String。必选：否&nbsp;topicId&nbsp;按照【日志主题ID】进行过滤。类型：String。必选：否&nbsp;logsetId&nbsp;按照【日志集ID】进行过滤，可通过调用&nbsp;DescribeLogsets&nbsp;查询已创建的日志集列表或登录控制台进行查看；也可以调用&nbsp;CreateLogset&nbsp;创建新的日志集。类型：String。必选：否&nbsp;tagKey&nbsp;按照【标签键】进行过滤。类型：String。必选：否&nbsp;tag:tagKey&nbsp;按照【标签键值对】进行过滤。tagKey&nbsp;使用具体的标签键进行替换，例如&nbsp;tag:exampleKey。类型：String。必选：否&nbsp;storageType&nbsp;按照【日志主题的存储类型】进行过滤。可选值&nbsp;hot（标准存储），cold（低频存储）类型：String。必选：否&nbsp;注意：每次请求的&nbsp;Filters&nbsp;的上限为10，Filter.Values&nbsp;的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 控制Filters相关字段是否为精确匹配。&nbsp;&nbsp;0:&nbsp;默认值，topicName&nbsp;和&nbsp;logsetName&nbsp;模糊匹配&nbsp;1:&nbsp;topicName&nbsp;精确匹配&nbsp;2:&nbsp;logsetName精确匹配&nbsp;3:&nbsp;topicName&nbsp;和logsetName&nbsp;都精确匹配

	PreciseSearch *uint64 `json:"PreciseSearch,omitempty" name:"PreciseSearch"`
	// 主题类型&nbsp;&nbsp;0:日志主题，默认值&nbsp;1:指标主题

	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformAutoCreateTopicsRequest struct {
	*tchttp.BaseRequest

	// 任务列表id列表

	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`
	// topicId&nbsp;按照【日志主题id】进行过滤。&nbsp;类型：String&nbsp;&nbsp;必选：否&nbsp;&nbsp;&nbsp;&nbsp;topicName&nbsp;按照【日志主题名称】进行过滤。&nbsp;类型：String&nbsp;&nbsp;必选：否&nbsp;&nbsp;&nbsp;&nbsp;logsetId&nbsp;按照【日志集id】进行过滤。&nbsp;类型：String&nbsp;&nbsp;必选：否&nbsp;&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDataTransformAutoCreateTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformAutoCreateTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerGroupInfo struct {

	// 消费组标识

	ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
	// 消费者心跳超时时间（秒）

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// topic列表

	Topics []*string `json:"Topics,omitempty" name:"Topics"`
}

type ShipperInfo struct {

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递的bucket地址

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 投递的前缀目录

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 投递规则的名字

	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`
	// 投递的时间间隔，单位&nbsp;秒

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 投递的文件的最大值，单位&nbsp;MB

	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
	// 是否生效。true表示生效，false表示失效。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 投递日志的过滤规则

	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`
	// 投递日志的分区规则，支持strftime的时间格式表示

	Partition *string `json:"Partition,omitempty" name:"Partition"`
	// 投递日志的压缩配置

	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`
	// 投递日志的内容格式配置

	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
	// 投递日志的创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）

	FilenameMode *uint64 `json:"FilenameMode,omitempty" name:"FilenameMode"`
	// 投递数据范围的开始时间点

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 投递数据范围的结束时间点

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 历史数据投递的进度（仅当用户选择的数据内中历史数据时才有效）

	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
	// 历史数据全部投递完成剩余的时间（仅当用户选择的数据中有历史数据时才有效）

	RemainTime *int64 `json:"RemainTime,omitempty" name:"RemainTime"`
	// 跨账号uin

	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`
	// 历史任务状态：&nbsp;0：实时任务&nbsp;1：任务准备中&nbsp;2：任务运行中&nbsp;3：任务运行异常&nbsp;4：任务运行结束

	HistoryStatus *int64 `json:"HistoryStatus,omitempty" name:"HistoryStatus"`
	// cos桶类型

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 角色访问描述名

	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`
	// 外部ID

	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
	// 任务运行状态。支持0,1,2&nbsp;&nbsp;0:&nbsp;停止&nbsp;1:&nbsp;运行中&nbsp;2:&nbsp;异常

	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
}

type DescribeKafkaConsumerPreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预览数据列表

		PreviewInfos []*string `json:"PreviewInfos,omitempty" name:"PreviewInfos"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaConsumerPreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumerPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuleTagInfo struct {

	// 是否大小写敏感。true表示开启大小写敏感，false表示关闭大小写敏感。

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 元字段索引配置中的字段信息

	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
}

type ModifyTopicExtendConfigRequest struct {
	*tchttp.BaseRequest

	// clb的topic业务配置

	ClbTopicExtendConfigs []*ClbTopicExtendConfig `json:"ClbTopicExtendConfigs,omitempty" name:"ClbTopicExtendConfigs"`
	// 修改模式。&nbsp;mode取值为1：则用lbkey更新ClbTopicExtendConfig；取值为2：则用user_uin更新FederationToken

	Mode *uint64 `json:"Mode,omitempty" name:"Mode"`
}

func (r *ModifyTopicExtendConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicExtendConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionArgument struct {

	// 参数序号，根据参数顺序定义

	ArgIndex *int64 `json:"ArgIndex,omitempty" name:"ArgIndex"`
	// 参数名称

	ArgName *string `json:"ArgName,omitempty" name:"ArgName"`
	// 参数描述

	ArgDesc *string `json:"ArgDesc,omitempty" name:"ArgDesc"`
	// 可接受的参数类型列表，包括字面常量、数组、条件表达式、函数表达式等任意一种或多种

	ArgType *string `json:"ArgType,omitempty" name:"ArgType"`
	// 参数默认值

	ArgValueDefault *string `json:"ArgValueDefault,omitempty" name:"ArgValueDefault"`
	// 范围、枚举类型

	ArgValueType *string `json:"ArgValueType,omitempty" name:"ArgValueType"`
	// 参数值域校验范围，这里仅针对常量进行校验，如果arg_value_type是scope类型，则此数组表示前闭后开区间，否则表示枚举的值类型.&nbsp;&nbsp;如果此值为空，或者空数组，则不进行参数值校验

	ArgValueScope []*string `json:"ArgValueScope,omitempty" name:"ArgValueScope"`
	// 是否必须。true表示必须，false代表不必须。

	IsNecessary *bool `json:"IsNecessary,omitempty" name:"IsNecessary"`
}

type DescribeShippersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递规则列表

		Shippers []*ShipperInfo `json:"Shippers,omitempty" name:"Shippers"`
		// 本次查询获取到的总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShippersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShippersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmTargetInfo struct {

	// 日志集ID。

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称。

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称。

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 查询语句。

	Query *string `json:"Query,omitempty" name:"Query"`
	// 告警对象序号。

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 查询范围起始时间相对于告警执行时间的偏移，单位为分钟，取值为非正，最大值为0，最小值为-1440。

	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`
	// 查询范围终止时间相对于告警执行时间的偏移，单位为分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。

	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
	// 检索语法规则，默认值为0。&nbsp;0：Lucene语法，1：CQL语法。&nbsp;详细说明参见检索条件语法规则

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 查询语句交互模式配置信息。

	QueryInteractiveConfig *string `json:"QueryInteractiveConfig,omitempty" name:"QueryInteractiveConfig"`
	// 主题类型。&nbsp;0:&nbsp;日志主题，1:&nbsp;指标主题

	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`
}

type DashboardTemplateVariable struct {

	// key的值

	Key *string `json:"Key,omitempty" name:"Key"`
	// key对应的values取值values

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeAccountResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户状态，0:未开通，1:正常，2:&nbsp;欠费，&nbsp;3:&nbsp;销毁

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterBaseMetricConfigsRequest struct {
	*tchttp.BaseRequest

	// 机器组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// topicId按照【指标主题id】进行过滤。类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，所有Filter.Values总和上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterBaseMetricConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterBaseMetricConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFunctionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeFunctionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFunctionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FolderInfo struct {

	// 文件夹Id。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 文件夹名称。

	Name *string `json:"Name,omitempty" name:"Name"`
}

type SearchLogTopics struct {

	// 多日志主题检索对应的错误信息

	Errors []*SearchLogErrors `json:"Errors,omitempty" name:"Errors"`
	// 多日志主题检索各日志主题信息

	Infos []*SearchLogInfos `json:"Infos,omitempty" name:"Infos"`
}

type AgentKeyInfo struct {

	// agent&nbsp;IP地址

	IP *string `json:"IP,omitempty" name:"IP"`
	// 实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type HeartBeatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// agent升级任务信息

		UpdateInfo *AgentUpdateInfo `json:"UpdateInfo,omitempty" name:"UpdateInfo"`
		// 服务下发策略

		Policy *PolicyInfo `json:"Policy,omitempty" name:"Policy"`
		// 请求机器关联机器组相关采集配置签名

		MachineConfigSign *string `json:"MachineConfigSign,omitempty" name:"MachineConfigSign"`
		// 用户UIN

		Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *HeartBeatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HeartBeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigRequest struct {
	*tchttp.BaseRequest

	// 采集规则配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 仪表盘详细明细

		DashboardInfos []*DashboardInfo `json:"DashboardInfos,omitempty" name:"DashboardInfos"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type DescribeMetricSubscribesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 指标订阅配置信息

		Datas []*MetricSubscribeInfo `json:"Datas,omitempty" name:"Datas"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricSubscribesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricSubscribesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeliverConfig struct {

	// 地域信息。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 日志主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递数据范围。&nbsp;&nbsp;0:&nbsp;全部日志,&nbsp;包括告警策略日常周期执行的所有日志，也包括告警策略变更产生的日志，默认值&nbsp;&nbsp;1:仅告警触发及恢复日志

	Scope *uint64 `json:"Scope,omitempty" name:"Scope"`
}

type PartitionOffsetInfo struct {

	// 分区id

	PartitionId *uint64 `json:"PartitionId,omitempty" name:"PartitionId"`
	// offset点位

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

type DashboardTopicInfo struct {

	// 主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// topic所在的地域

	Region *string `json:"Region,omitempty" name:"Region"`
}

type ExternalDataSourceSQLInfo struct {

	// 访问方式。&nbsp;1：私有链接(内网访问)；2：外网访问；&nbsp;&nbsp;当DataSource为1时，AccessMode必填

	AccessMode *uint64 `json:"AccessMode,omitempty" name:"AccessMode"`
	// SQL访问IP地址。&nbsp;&nbsp;当DataSource为1且AccessMode为2时，Address必填。

	Address *string `json:"Address,omitempty" name:"Address"`
	// SQL访问端口。&nbsp;&nbsp;当DataSource为1时，Port必填

	Port *string `json:"Port,omitempty" name:"Port"`
	// 私有链接终端节点服务ID。&nbsp;&nbsp;当DataSource为1时，当AccessMode&nbsp;为1时，&nbsp;EndPointServiceId字段必填。

	EndPointServiceId *string `json:"EndPointServiceId,omitempty" name:"EndPointServiceId"`
	// SQL访问用户名

	User *string `json:"User,omitempty" name:"User"`
	// SQL访问密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// SQL访问数据库名

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	// SQL访问表名

	TableName *string `json:"TableName,omitempty" name:"TableName"`
	// InstanceId所属地域。当DataSource为1且AccessMode为3时，Region必填。&nbsp;当DataSource为3、4时，Region必填。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 所属网络VpcId。&nbsp;&nbsp;当DataSource为1且AccessMode为3时，VpcId必填

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 网络服务类型。0：CVM；1025：CLB&nbsp;&nbsp;当DataSource为1且AccessMode为3时，VirtualGatewayType必填

	VirtualGatewayType *uint64 `json:"VirtualGatewayType,omitempty" name:"VirtualGatewayType"`
	// 实例Id。&nbsp;&nbsp;当DataSource为3、4时，InstanceId必填。&nbsp;当DataSource为3时，表示&nbsp;云数据库Mysql&nbsp;实例id，如：cdb-zxcvbnm&nbsp;当DataSource为4时，表示&nbsp;云数据库TDSQL-C&nbsp;Mysql&nbsp;集群id，如：cynosdbmysql-i6bdemo

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

type CreateKafkaRechargeRequest struct {
	*tchttp.BaseRequest

	// 导入CLS目标topic&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// Kafka导入配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 导入Kafka类型，0:&nbsp;CKafka，1:&nbsp;用户自建Kafka

	KafkaType *uint64 `json:"KafkaType,omitempty" name:"KafkaType"`
	// CKafka实例ID，KafkaType为0时必填

	KafkaInstance *string `json:"KafkaInstance,omitempty" name:"KafkaInstance"`
	// 服务地址，KafkaType为1时必填。

	ServerAddr *string `json:"ServerAddr,omitempty" name:"ServerAddr"`
	// ServerAddr是否为加密连接，KafkaType为1时必填

	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitempty" name:"IsEncryptionAddr"`
	// 加密访问协议，IsEncryptionAddr参数为true时必填

	Protocol *KafkaProtocolInfo `json:"Protocol,omitempty" name:"Protocol"`
	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开

	UserKafkaTopics *string `json:"UserKafkaTopics,omitempty" name:"UserKafkaTopics"`
	// 用户Kafka消费组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 导入数据位置，-2:最早（默认），-1：最晚

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 日志导入规则

	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitempty" name:"LogRechargeRule"`
	// 网络信息参数

	NetworkInfo *NetworkInfo `json:"NetworkInfo,omitempty" name:"NetworkInfo"`
}

func (r *CreateKafkaRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKafkaRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigException struct {

	// 采集配置Id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 状态：&nbsp;1：非结构化提取模式

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 采集配置名称

	ConfigName *string `json:"ConfigName,omitempty" name:"ConfigName"`
}

type DeleteMetricSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMetricSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetricSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题ID

		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
		// 是否生效。true表示索引生效，false表示索引失效。

		Status *bool `json:"Status,omitempty" name:"Status"`
		// 索引配置信息

		Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
		// 索引修改时间，初始值为索引创建时间。

		ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
		// 内置保留字段（__FILENAME__，__HOSTNAME__及__SOURCE__）是否包含至全文索引&nbsp;&nbsp;false:不包含&nbsp;true:包含

		IncludeInternalFields *bool `json:"IncludeInternalFields,omitempty" name:"IncludeInternalFields"`
		// 元数据字段（前缀为__TAG__的字段）是否包含至全文索引&nbsp;&nbsp;0:仅包含开启键值索引的元数据字段&nbsp;1:包含所有元数据字段&nbsp;2:不包含任何元数据字段

		MetadataFlag *uint64 `json:"MetadataFlag,omitempty" name:"MetadataFlag"`
		// 自定义日志解析异常存储字段。

		CoverageField *string `json:"CoverageField,omitempty" name:"CoverageField"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmNotice struct {

	// 告警通知模板名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 告警模板的类型。可选值：&nbsp;&nbsp;Trigger&nbsp;-&nbsp;告警触发&nbsp;&nbsp;Recovery&nbsp;-&nbsp;告警恢复&nbsp;&nbsp;All&nbsp;-&nbsp;告警触发和告警恢复

	Type *string `json:"Type,omitempty" name:"Type"`
	// 告警通知模板接收者信息。

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 告警通知模板回调信息。

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 告警通知模板ID。

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 告警通知模板绑定的标签信息。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 调用链接域名。http://&nbsp;或者&nbsp;https://&nbsp;开头，不能/结尾

	JumpDomain *string `json:"JumpDomain,omitempty" name:"JumpDomain"`
	// 通知规则。

	NoticeRules []*NoticeRule `json:"NoticeRules,omitempty" name:"NoticeRules"`
	// 免登录操作告警开关。&nbsp;参数值：&nbsp;1：关闭&nbsp;2：开启（默认开启）

	AlarmShieldStatus *uint64 `json:"AlarmShieldStatus,omitempty" name:"AlarmShieldStatus"`
	// 投递相关信息。

	AlarmNoticeDeliverConfig *AlarmNoticeDeliverConfig `json:"AlarmNoticeDeliverConfig,omitempty" name:"AlarmNoticeDeliverConfig"`
	// 投递日志开关。&nbsp;&nbsp;参数值：&nbsp;&nbsp;1：关闭&nbsp;&nbsp;2：开启

	DeliverStatus *uint64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`
	// 投递日志标识。&nbsp;&nbsp;参数值：&nbsp;&nbsp;1：未启用&nbsp;&nbsp;2：已启用&nbsp;&nbsp;3：投递异常

	DeliverFlag *uint64 `json:"DeliverFlag,omitempty" name:"DeliverFlag"`
	// 通知渠道组配置的告警屏蔽统计状态数量信息。

	AlarmShieldCount *AlarmShieldCount `json:"AlarmShieldCount,omitempty" name:"AlarmShieldCount"`
}

type Demonstration struct {

	// 演示示例资源所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 演示示例类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 演示示例资源

	Resources []*DemonstrationResource `json:"Resources,omitempty" name:"Resources"`
	// 演示示例状态：CREATING,&nbsp;FAILED,&nbsp;SUCCESS,&nbsp;DELETING

	Status *string `json:"Status,omitempty" name:"Status"`
	// 演示示例ID

	DemonstrationId *string `json:"DemonstrationId,omitempty" name:"DemonstrationId"`
	// 是否包含演示示例日志主题

	HasDemonstrationTopic *bool `json:"HasDemonstrationTopic,omitempty" name:"HasDemonstrationTopic"`
	// 演示示例子类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
}

type DescribeUserConfigRequest struct {
	*tchttp.BaseRequest

	// 要查询配置key列表

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
}

func (r *DescribeUserConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmAnalysisConfig struct {

	// 键。支持以下key：&nbsp;SyntaxRule：语法规则，value支持&nbsp;0：Lucene语法；1：&nbsp;CQL语法。&nbsp;QueryIndex：执行语句序号。value支持&nbsp;-1：自定义；&nbsp;1：执行语句1；&nbsp;2：执行语句2。&nbsp;CustomQuery：检索语句。&nbsp;QueryIndex为-1时有效且必填，value示例：&nbsp;"*&nbsp;|&nbsp;select&nbsp;count(*)&nbsp;as&nbsp;count"。&nbsp;Fields：字段。value支持&nbsp;SOURCE；FILENAME；HOSTNAME；TIMESTAMP；INDEX_STATUS；PKG_LOGID；TOPIC。&nbsp;Format：显示形式。value支持&nbsp;1：每条日志一行；2：每条日志每个字段一行。&nbsp;Limit：最大日志条数。&nbsp;value示例：&nbsp;5。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 值。&nbsp;键对应值如下：&nbsp;SyntaxRule：语法规则，value支持&nbsp;0：Lucene语法；1：&nbsp;CQL语法。&nbsp;QueryIndex：执行语句序号。value支持&nbsp;-1：自定义；&nbsp;1：执行语句1；&nbsp;2：执行语句2。&nbsp;CustomQuery：检索语句。&nbsp;QueryIndex为-1时有效且必填，value示例：&nbsp;"*&nbsp;|&nbsp;select&nbsp;count(*)&nbsp;as&nbsp;count"。&nbsp;Fields：字段。value支持&nbsp;SOURCE；FILENAME；HOSTNAME；TIMESTAMP；INDEX_STATUS；PKG_LOGID；TOPIC。&nbsp;Format：显示形式。value支持&nbsp;1：每条日志一行；2：每条日志每个字段一行。&nbsp;Limit：最大日志条数。&nbsp;value示例：&nbsp;5。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type UserConfigInfo struct {

	// 附加配置key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 附加配置内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteExternalDataSourceRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID,&nbsp;此外部数据源配置隶属此日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 外部数据源配置ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteExternalDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExternalDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigurationTemplateInfo struct {

	// 配置模板id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 版本号

	Version *uint64 `json:"Version,omitempty" name:"Version"`
	// 配置模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 配置模板描述

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 投递cos配置模板

	ShipperTemplateInfos []*ShipperTemplateInfo `json:"ShipperTemplateInfos,omitempty" name:"ShipperTemplateInfos"`
	// 创建时间

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ShipperTaskInfo struct {

	// 投递任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 投递信息ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 本批投递的日志的开始时间戳，毫秒

	RangeStart *int64 `json:"RangeStart,omitempty" name:"RangeStart"`
	// 本批投递的日志的结束时间戳，&nbsp;毫秒

	RangeEnd *int64 `json:"RangeEnd,omitempty" name:"RangeEnd"`
	// 本次投递任务的开始时间戳，&nbsp;毫秒

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 本次投递任务的结束时间戳，&nbsp;毫秒

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// 本次投递的结果，"success","running","failed"

	Status *string `json:"Status,omitempty" name:"Status"`
	// 结果的详细信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type DeleteScheduledSqlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScheduledSqlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScheduledSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebCallbackRequest struct {
	*tchttp.BaseRequest

	// 告警渠道回调配置ID。

	WebCallbackId *string `json:"WebCallbackId,omitempty" name:"WebCallbackId"`
	// 告警渠道回调配置名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 渠道类型&nbsp;&nbsp;WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调;

	Type *string `json:"Type,omitempty" name:"Type"`
	// 回调地址。

	Webhook *string `json:"Webhook,omitempty" name:"Webhook"`
	// 请求方式。&nbsp;&nbsp;支持POST、PUT。&nbsp;&nbsp;注意：当Type为Http时，必填。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 秘钥信息。

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *ModifyWebCallbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebCallbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClbTopicExtendConfig struct {

	// LB关键信息，VIP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// LB关键信息，VpcId

	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`
	// clb服务端的公共topic

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// clb用户的topic

	UserTopicId *string `json:"UserTopicId,omitempty" name:"UserTopicId"`
	// clb用户的uin信息

	UserUin *uint64 `json:"UserUin,omitempty" name:"UserUin"`
	// clb用户的appid信息

	UserAppId *uint64 `json:"UserAppId,omitempty" name:"UserAppId"`
	// 临时证书加密密钥ID。最长不超过1024字节。

	UserTmpSecretId *string `json:"UserTmpSecretId,omitempty" name:"UserTmpSecretId"`
	// 临时证书加密密钥Key。最长不超过1024字节。

	UserTmpSecretKey *string `json:"UserTmpSecretKey,omitempty" name:"UserTmpSecretKey"`
	// token,&nbsp;最长不超过4096字节。

	UserToken *string `json:"UserToken,omitempty" name:"UserToken"`
	// 临时证书有效的时间，返回&nbsp;Unix&nbsp;时间戳，精确到秒

	TmpKeyExpired *uint64 `json:"TmpKeyExpired,omitempty" name:"TmpKeyExpired"`
	// 唯一标识clb的一种业务

	LbKey *string `json:"LbKey,omitempty" name:"LbKey"`
	// 公共topic的采样比

	LogSample *string `json:"LogSample,omitempty" name:"LogSample"`
	// 用户topic的采样比

	UserSample *string `json:"UserSample,omitempty" name:"UserSample"`
	// LB健康检查日志&nbsp;Topic&nbsp;ID,&nbsp;和topicId属于另外一种公共的topic

	UserHealthTopicId *string `json:"UserHealthTopicId,omitempty" name:"UserHealthTopicId"`
	// topic的采集配置是否生效,true为生效,false为失效

	UserSampleStatus *bool `json:"UserSampleStatus,omitempty" name:"UserSampleStatus"`
	// 1代表用户topic已删除，0代表用户topic未删除。

	UserTopicStatus *uint64 `json:"UserTopicStatus,omitempty" name:"UserTopicStatus"`
	// lbkey是否要采集到公共topic,&nbsp;true为要采集，&nbsp;false为不采集，默认为false

	Collection *bool `json:"Collection,omitempty" name:"Collection"`
	// LB相关业务ID

	LbID *string `json:"LbID,omitempty" name:"LbID"`
	// 是否展示ServerAddr字段，true&nbsp;展示，false&nbsp;不展示

	ShowServerAddr *bool `json:"ShowServerAddr,omitempty" name:"ShowServerAddr"`
}

type DataTransformTaskInfo struct {

	// 数据加工任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务启用状态，默认为1，正常开启,&nbsp;&nbsp;2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 加工任务类型，1：&nbsp;DSL，&nbsp;2：SQL

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 源日志主题

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 当前加工任务状态（1准备中/2运行中/3停止中/4已停止）

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 加工任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近修改时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 最后启用时间，如果需要重建集群，修改该时间

	LastEnableTime *string `json:"LastEnableTime,omitempty" name:"LastEnableTime"`
	// 日志主题名称

	SrcTopicName *string `json:"SrcTopicName,omitempty" name:"SrcTopicName"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 加工任务目的topic_id以及别名

	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
	// 加工逻辑函数

	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
	// 兜底topic_id

	BackupTopicID *string `json:"BackupTopicID,omitempty" name:"BackupTopicID"`
	// 超限之后是否丢弃日志数据

	BackupGiveUpData *bool `json:"BackupGiveUpData,omitempty" name:"BackupGiveUpData"`
	// 交互式加工逻辑函数配置。

	InteractiveEtlContent *string `json:"InteractiveEtlContent,omitempty" name:"InteractiveEtlContent"`
	// 是否开启投递服务日志。&nbsp;1关闭,2开启

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
	// 任务目标日志主题数量

	TaskDstCount *uint64 `json:"TaskDstCount,omitempty" name:"TaskDstCount"`
	// 数据加工类型。0：标准加工任务；1：前置加工任务。

	DataTransformType *uint64 `json:"DataTransformType,omitempty" name:"DataTransformType"`
	// 保留失败日志状态。&nbsp;1:不保留，2:保留

	KeepFailureLog *uint64 `json:"KeepFailureLog,omitempty" name:"KeepFailureLog"`
	// 失败日志的字段名称

	FailureLogKey *string `json:"FailureLogKey,omitempty" name:"FailureLogKey"`
}

type Instance struct {

	// 实例信息

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type LogItems struct {

	// 分析结果返回的KV数据对

	Data []*LogItem `json:"Data,omitempty" name:"Data"`
}

type QcloudGoodsInfoListInfo struct {

	// 付费模式，0:后付费

	PayMode *uint64 `json:"PayMode,omitempty" name:"PayMode"`
	// 业务产品录入的商品码，业务名称，categoryid中的商品码

	Type *string `json:"Type,omitempty" name:"Type"`
	// 地域ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 区域ID

	ZoneId *uint64 `json:"ZoneId,omitempty" name:"ZoneId"`
	// 商品实例的个数

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 业务参数，用户询价和透传给业务

	GoodsDetail *QcloudGoodsDetailInfo `json:"GoodsDetail,omitempty" name:"GoodsDetail"`
}

type RegexIndexInfo struct {

	// 起始位置

	Start *int64 `json:"Start,omitempty" name:"Start"`
	// 结束位置

	End *int64 `json:"End,omitempty" name:"End"`
}

type CloseKafkaConsumeRequest struct {
	*tchttp.BaseRequest

	// CLS对应的topic标识

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *CloseKafkaConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKafkaConsumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateShipperRequest struct {
	*tchttp.BaseRequest

	// 创建的投递规则所属的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 创建的投递规则投递的bucket

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 创建的投递规则投递目录的前缀

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 投递规则的名字

	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`
	// 投递的时间间隔，单位&nbsp;秒，默认300，范围&nbsp;300-900

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 投递的文件的最大值，单位&nbsp;MB，默认100，范围&nbsp;5-256

	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递

	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`
	// 投递日志的分区规则，支持strftime的时间格式表示

	Partition *string `json:"Partition,omitempty" name:"Partition"`
	// 投递日志的压缩配置

	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`
	// 投递日志的内容格式配置

	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
	// 投递文件命名配置，0：随机数命名，1：投递时间命名，默认0（随机数命名）

	FilenameMode *uint64 `json:"FilenameMode,omitempty" name:"FilenameMode"`
	// 跨账户uin，用于支持跨账户bucket投递

	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`
	// 投递数据范围的开始时间点，不能超出日志主题的生命周期起点。如果用户不填写，默认为用户新建投递任务的时间。

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 投递数据范围的结束时间点，不能填写未来时间。如果用户不填写，默认为持续投递，即无限。

	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
	// cos桶存储类型。支持：STANDARD_IA、ARCHIVE、DEEP_ARCHIVE、STANDARD、MAZ_STANDARD、MAZ_STANDARD_IA、INTELLIGENT_TIERING。&nbsp;&nbsp;STANDARD_IA：低频存储；&nbsp;ARCHIVE：归档存储；&nbsp;DEEP_ARCHIVE：深度归档存储；&nbsp;STANDARD：标准存储；&nbsp;MAZ_STANDARD：标准存储（多&nbsp;AZ）；&nbsp;MAZ_STANDARD_IA：低频存储（多&nbsp;AZ）；&nbsp;INTELLIGENT_TIERING：智能分层存储。

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 角色访问描述名

	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`
	// 外部ID

	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
}

func (r *CreateShipperRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateShipperRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordingRuleTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 源日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 目标主题id

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务启动状态.&nbsp;1开启,&nbsp;2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 调度开始时间,Unix时间戳，单位ms

	ProcessStartTime *int64 `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 调度周期(分钟)，支持范围(0,1440]分钟。

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 执行周期单位,&nbsp;`0`:&nbsp;Minute,&nbsp;`1`:&nbsp;Second。&nbsp;默认：0&nbsp;-&nbsp;ProcessPeriodUnit:1时，ProcessPeriod&nbsp;只支持&nbsp;30（即只支持30秒）。&nbsp;-&nbsp;ProcessPeriodUnit:0时，ProcessPeriod&nbsp;支持范围(0,1440]分钟。

	ProcessPeriodUnit *uint64 `json:"ProcessPeriodUnit,omitempty" name:"ProcessPeriodUnit"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 预聚合检索语句

	RecordingRuleContent *string `json:"RecordingRuleContent,omitempty" name:"RecordingRuleContent"`
	// 自定义指标名称

	CustomMetricLabels []*MetricLabel `json:"CustomMetricLabels,omitempty" name:"CustomMetricLabels"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *ModifyRecordingRuleTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecordingRuleTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigInfo struct {

	// 采集规则配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 采集规则配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 日志采集路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 采集的日志类型。&nbsp;&nbsp;json_log代表：JSON-文件日志（详见使用&nbsp;JSON&nbsp;提取模式采集日志）；&nbsp;delimiter_log代表：分隔符-文件日志（详见使用分隔符提取模式采集日志）；&nbsp;minimalist_log代表：单行全文-文件日志（详见使用单行全文提取模式采集日志）；&nbsp;fullregex_log代表：单行完全正则-文件日志（详见使用单行-完全正则提取模式采集日志）；&nbsp;multiline_log代表：多行全文-文件日志（详见使用多行全文提取模式采集日志）；&nbsp;multiline_fullregex_log代表：多行完全正则-文件日志（详见使用多行-完全正则提取模式采集日志）；&nbsp;user_define_log代表：组合解析（适用于多格式嵌套的日志，详见使用组合解析提取模式采集日志）；&nbsp;service_syslog代表：syslog&nbsp;采集（详见采集&nbsp;Syslog）；&nbsp;windows_event_log代表：Windows事件日志（详见采集&nbsp;Windows&nbsp;事件日志）。

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 采集配置所属日志主题ID即TopicId

	Output *string `json:"Output,omitempty" name:"Output"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 用户自定义解析字符串，详见使用组合解析提取模式采集日志。

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// config_extra主键ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集配置标签

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 高级采集配置。&nbsp;Json字符串，&nbsp;Key/Value定义为如下：&nbsp;&nbsp;ClsAgentFileTimeout(超时属性),&nbsp;取值范围:&nbsp;大于等于0的整数，&nbsp;0为不超时&nbsp;ClsAgentMaxDepth(最大目录深度)，取值范围:&nbsp;大于等于0的整数&nbsp;ClsAgentParseFailMerge(合并解析失败日志)，取值范围:&nbsp;true或false&nbsp;样例：&nbsp;{\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}&nbsp;控制台默认占位值：{\"ClsAgentDefault\":0}

	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
	// 采集配置来源，0:&nbsp;默认来源，1:&nbsp;TKE

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 日志输入类型，支持lfile、window_event、syslog、k8s_stdout、k8s_file

	InputType *string `json:"InputType,omitempty" name:"InputType"`
}

type MachineGroupInfo struct {

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 机器组类型

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 机器组绑定的标签列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启机器组自动更新。true表示开启自动更新，false表示关闭自动更新。

	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 升级开始时间，建议业务低峰期升级LogListener

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 升级结束时间，建议业务低峰期升级LogListener

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 是否开启服务日志，用于记录因Loglistener&nbsp;服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费。true表示开启服务日志，false表示关闭服务日志。

	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
	// TKE标志位，默认值为空字符串。空字符串表示日志不是来自于TKE，label_k8s表示日志来自于TKE。

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// 机器组中机器离线定期清理时间

	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitempty" name:"DelayCleanupTime"`
	// 机器组元数据信息列表

	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`
	// 操作系统类型，0:&nbsp;Linux，1:&nbsp;windows

	OSType *uint64 `json:"OSType,omitempty" name:"OSType"`
	// 专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 集群ID。-&nbsp;当Flag为&nbsp;label_tke，表示TKE集群ID-&nbsp;ClusterId&nbsp;与&nbsp;ClusterRegion&nbsp;必须同时填写，或者为空

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群所在地域-&nbsp;当Flag为&nbsp;label_tke，表示TKE集群地域-&nbsp;ClusterId&nbsp;与&nbsp;ClusterRegion&nbsp;必须同时填写，或者为空

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

type DeleteTopicExtendConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicExtendConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicExtendConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 告警模板名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 通知类型。可选值：&nbsp;&nbsp;Trigger&nbsp;-&nbsp;告警触发&nbsp;Recovery&nbsp;-&nbsp;告警恢复&nbsp;All&nbsp;-&nbsp;告警触发和告警恢复

	Type *string `json:"Type,omitempty" name:"Type"`
	// 告警模板接收者信息。

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 接口回调信息（包括企业微信等）。

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 告警通知模板ID。

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的通知渠道组。最大支持10个标签键值对，并且不能有重复的键值对。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 调用链接域名。http://&nbsp;或者&nbsp;https://&nbsp;开头，不能/结尾

	JumpDomain *string `json:"JumpDomain,omitempty" name:"JumpDomain"`
	// 通知规则。&nbsp;&nbsp;注意:&nbsp;&nbsp;Type、NoticeReceivers和WebCallbacks是一组配置，NoticeRules是另一组配置，2组配置互斥。&nbsp;传其中一组数据，则另一组数据置空。

	NoticeRules []*NoticeRule `json:"NoticeRules,omitempty" name:"NoticeRules"`
	// 投递日志开关。&nbsp;&nbsp;参数值：&nbsp;1：关闭；&nbsp;&nbsp;2：开启

	DeliverStatus *uint64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`
	// 投递日志配置。

	DeliverConfig *DeliverConfig `json:"DeliverConfig,omitempty" name:"DeliverConfig"`
	// 免登录操作告警开关。&nbsp;&nbsp;参数值：&nbsp;1：关闭&nbsp;2：开启（默认开启）

	AlarmShieldStatus *uint64 `json:"AlarmShieldStatus,omitempty" name:"AlarmShieldStatus"`
}

func (r *ModifyAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CallBackInfo struct {

	// 回调时的Body。&nbsp;可将各类告警变量放在请求内容中，详见帮助文档。&nbsp;如下示例：&nbsp;{"data":"data"}

	Body *string `json:"Body,omitempty" name:"Body"`
	// 回调时的HTTP请求头部字段。&nbsp;例如：下面请求头部字段来告知服务器请求主体的内容类型为JSON。&nbsp;&nbsp;"Content-Type:&nbsp;application/json"

	Headers []*string `json:"Headers,omitempty" name:"Headers"`
}

type SearchCosRechargeInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 匹配到的存储桶下的某个文件的前几行数据

		Data []*string `json:"Data,omitempty" name:"Data"`
		// 匹配到的存储桶下的文件个数

		Sum *uint64 `json:"Sum,omitempty" name:"Sum"`
		// 当前预览文件路径

		Path *string `json:"Path,omitempty" name:"Path"`
		// 预览获取数据失败原因

		Msg *string `json:"Msg,omitempty" name:"Msg"`
		// 状态

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchCosRechargeInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchCosRechargeInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyValueInfo struct {

	// 需要配置键值或者元字段索引的字段名称，仅支持字母、数字、下划线和-./@，且不能以下划线开头&nbsp;&nbsp;注意：&nbsp;1，元字段（tag）的Key无需额外添加__TAG__.前缀，与上传日志时对应的字段Key一致即可，云控制台展示时将自动添加__TAG__.前缀&nbsp;2，键值索引（KeyValue）及元字段索引（Tag）中的Key总数不能超过300&nbsp;3，Key的层级不能超过10层，例如a.b.c.d.e.f.g.h.j.k&nbsp;4，不允许同时包含json父子级字段，例如a及a.b

	Key *string `json:"Key,omitempty" name:"Key"`
	// 字段的索引描述信息

	Value *ValueInfo `json:"Value,omitempty" name:"Value"`
}

type UpdateAgentStatusRequest struct {
	*tchttp.BaseRequest

	// Agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// Agent的Instance&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Agent的请求序列号

	AgentSeq *string `json:"AgentSeq,omitempty" name:"AgentSeq"`
	// Agent当前版本

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// Agent升级目标版本

	UpdateVersion *string `json:"UpdateVersion,omitempty" name:"UpdateVersion"`
	// Agent升级状态信息

	AgentStatus *AgentUpdateStatus `json:"AgentStatus,omitempty" name:"AgentStatus"`
}

func (r *UpdateAgentStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAgentStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集规则配置列表

		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGroupConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetAlarmLogRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 查询语句，语句长度最大为1024

	Query *string `json:"Query,omitempty" name:"Query"`
	// 单次查询返回的执行详情条数，最大值为1000

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。&nbsp;注意：&nbsp;&nbsp;透传该参数时，请勿修改除该参数外的其它参数&nbsp;仅当检索分析语句(Query)不包含SQL时有效，SQL获取后续结果参考SQL&nbsp;LIMIT语法

	Context *string `json:"Context,omitempty" name:"Context"`
	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为&nbsp;desc&nbsp;注意：&nbsp;&nbsp;仅当检索分析语句(Query)不包含SQL时有效&nbsp;SQL结果排序方式参考SQL&nbsp;ORDER&nbsp;BY语法

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// true：代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效；&nbsp;false：代表使用老的检索结果返回方式，输出AnalysisResults和ColNames有效；&nbsp;两种返回方式在编码格式上有少量区别，建议使用true。

	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
}

func (r *GetAlarmLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlarmLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigExtraResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConfigExtraResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteKafkaRechargeRequest struct {
	*tchttp.BaseRequest

	// Kafka导入配置ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 导入CLS目标topic&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteKafkaRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteKafkaRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerInfo struct {

	// 投递规则ID

	ConsumerId *string `json:"ConsumerId,omitempty" name:"ConsumerId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递任务是否生效

	Effective *bool `json:"Effective,omitempty" name:"Effective"`
	// CKafka的描述

	Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`
	// 是否投递日志的元数据信息

	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`
	// 如果需要投递元数据信息，元数据信息的描述

	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`
	// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]

	Compression *int64 `json:"Compression,omitempty" name:"Compression"`
	// 投递任务创建毫秒时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 角色访问描述名

	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`
	// 外部ID

	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
	// 任务运行状态。支持`0`,`1`,`2`&nbsp;-&nbsp;`0`:&nbsp;停止&nbsp;-&nbsp;`1`:&nbsp;运行中&nbsp;-&nbsp;`2`:&nbsp;异常

	TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
	// 高级配置

	AdvancedConfig *AdvancedConsumerConfiguration `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
}

type MetricCorrectDimension struct {

	// 指标名称。

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标维度。

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
}

type CreateConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldValueRatioInfos struct {

	// 字段值

	Value *string `json:"Value,omitempty" name:"Value"`
	// 字段值所占的数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 字段值所占的比例

	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
}

type HostFileInfo struct {

	// 日志文件夹

	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
	// 日志文件名

	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`
	// metadata信息

	CustomLabels []*string `json:"CustomLabels,omitempty" name:"CustomLabels"`
	// 日志文件路径信息

	FilePaths []*FilePathInfo `json:"FilePaths,omitempty" name:"FilePaths"`
}

type DeleteConfigFromMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 采集配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteConfigFromMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigFromMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExternalDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyExternalDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExternalDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyWebCallbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWebCallbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyWebCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataTransformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDataTransformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMetricConfigRequest struct {
	*tchttp.BaseRequest

	// 日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 指标采集配置id。

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DeleteMetricConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetricConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformFailLogInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据加工任务失败日志详请

		LogFailureInfos []*DataTransformFailureInfo `json:"LogFailureInfos,omitempty" name:"LogFailureInfos"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformFailLogInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformFailLogInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExternalDataSourcePreviewRequest struct {
	*tchttp.BaseRequest

	// 数据源类型。支持1,2,3,4&nbsp;&nbsp;1：自建MySQL&nbsp;2：COS(csv格式)&nbsp;3：云数据库MySQL&nbsp;4：云数据库TDSQL-C&nbsp;MySQL

	Datasource *uint64 `json:"Datasource,omitempty" name:"Datasource"`
	// SQL相关配置信息，当Datasource值为1时，此字段必填

	SQLInfo *ExternalDataSourceSQLInfo `json:"SQLInfo,omitempty" name:"SQLInfo"`
	// COS相关配置信息，当Datasource值为2时，此字段必填

	COSInfo *ExternalDataSourceCOSInfo `json:"COSInfo,omitempty" name:"COSInfo"`
	// 外部数据源ID

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DescribeExternalDataSourcePreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExternalDataSourcePreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestJsonLogRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeLatestJsonLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestJsonLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BaseMetricCollectConfig struct {

	// 机器组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 基础监控采集配置信息

	Configs []*MetricCollectConfig `json:"Configs,omitempty" name:"Configs"`
}

type ModifyConfigRequest struct {
	*tchttp.BaseRequest

	// 采集规则配置ID，通过获取采集规则配置返回信息获取。

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 采集规则配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 日志采集路径，包含文件名

	Path *string `json:"Path,omitempty" name:"Path"`
	// 采集的日志类型。支持以下类型：&nbsp;&nbsp;json_log代表：JSON-文件日志（详见使用&nbsp;JSON&nbsp;提取模式采集日志）；&nbsp;delimiter_log代表：分隔符-文件日志（详见使用分隔符提取模式采集日志）；&nbsp;minimalist_log代表：单行全文-文件日志（详见使用单行全文提取模式采集日志）；&nbsp;fullregex_log代表：单行完全正则-文件日志（详见使用单行-完全正则提取模式采集日志）；&nbsp;multiline_log代表：多行全文-文件日志（详见使用多行全文提取模式采集日志）；&nbsp;multiline_fullregex_log代表：多行完全正则-文件日志（详见使用多行-完全正则提取模式采集日志）；&nbsp;user_define_log代表：组合解析（适用于多格式嵌套的日志，详见使用组合解析提取模式采集日志）；&nbsp;service_syslog代表：syslog&nbsp;采集（详见采集&nbsp;Syslog）；&nbsp;windows_event_log代表：Windows事件日志（详见采集&nbsp;Windows&nbsp;事件日志）。

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 采集配置关联的日志主题（TopicId）

	Output *string `json:"Output,omitempty" name:"Output"`
	// 用户自定义解析字符串，Json格式序列化的字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// config_extra表主键ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集配置标

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 高级采集配置。&nbsp;Json字符串，&nbsp;Key/Value定义为如下：&nbsp;&nbsp;ClsAgentFileTimeout(超时属性),&nbsp;取值范围:&nbsp;大于等于0的整数，&nbsp;0为不超时&nbsp;ClsAgentMaxDepth(最大目录深度)，取值范围:&nbsp;大于等于0的整数&nbsp;ClsAgentParseFailMerge(合并解析失败日志)，取值范围:&nbsp;true或false&nbsp;样例：&nbsp;{\"ClsAgentFileTimeout\":0,\"ClsAgentMaxDepth\":10,\"ClsAgentParseFailMerge\":true}

	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
	// 采集配置来源，0：默认来源，1:&nbsp;TKE

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 日志输入类型，支持file、window_event、syslog、k8s_stdout、k8s_file

	InputType *string `json:"InputType,omitempty" name:"InputType"`
}

func (r *ModifyConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardSubscribeInfo struct {

	// 仪表盘订阅id。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 仪表盘订阅名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 仪表盘id。

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘订阅时间。

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 仪表盘订阅数据。

	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitempty" name:"SubscribeData"`
	// 仪表盘订阅记录创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 仪表盘订阅记录更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 仪表盘订阅记录最后一次发送成功时间。

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 主账号Id。

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 主账号下的子账号Id。

	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
	// 仪表盘订阅记录最后一次发送的状态。success：全部发送成功，fail：未发送，&nbsp;partialSuccess：部分发送成功。

	LastStatus *string `json:"LastStatus,omitempty" name:"LastStatus"`
}

type NamespaceCorrectDimension struct {

	// 命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标维度信息。

	MetricDimensions []*MetricCorrectDimension `json:"MetricDimensions,omitempty" name:"MetricDimensions"`
}

type PartitionInfo struct {

	// 分区ID

	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
	// 分区的状态（readwrite或者是readonly）

	Status *string `json:"Status,omitempty" name:"Status"`
	// 分区哈希键起始key

	InclusiveBeginKey *string `json:"InclusiveBeginKey,omitempty" name:"InclusiveBeginKey"`
	// 分区哈希键结束key

	ExclusiveEndKey *string `json:"ExclusiveEndKey,omitempty" name:"ExclusiveEndKey"`
	// 分区创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 只读分区数据停止写入时间

	LastWriteTime *string `json:"LastWriteTime,omitempty" name:"LastWriteTime"`
}

type MetricYamlSpec struct {

	// yaml监控类型。&nbsp;支持：&nbsp;-&nbsp;PodMonitor&nbsp;-&nbsp;ServiceMonitor&nbsp;-&nbsp;ScrapeConfig&nbsp;-&nbsp;ScrapeConfig-prometheus&nbsp;`PodMonitor&nbsp;`,`ServiceMonitor&nbsp;`,`ScrapeConfig&nbsp;`&nbsp;属于prometheus-operator&nbsp;`ScrapeConfig-prometheus`&nbsp;属于prometheus

	Type *string `json:"Type,omitempty" name:"Type"`
	// 配置yaml格式。

	Spec *string `json:"Spec,omitempty" name:"Spec"`
}

type MonitorTime struct {

	// 执行周期，&nbsp;可选值：Period、Fixed、Cron。&nbsp;&nbsp;Period：固定频率&nbsp;Fixed：固定时间&nbsp;Cron：Cron表达式

	Type *string `json:"Type,omitempty" name:"Type"`
	// 执行的周期，或者定制执行的时间节点。单位为分钟，取值范围为1~1440。&nbsp;当type为Period,Fixed时，time字段生效。

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 执行的周期cron表达式。示例："*/1&nbsp;*&nbsp;*&nbsp;*&nbsp;*"&nbsp;从左到右每个field的含义&nbsp;Minutes&nbsp;field,&nbsp;Hours&nbsp;field,Day&nbsp;of&nbsp;month&nbsp;field,Month&nbsp;field,Day&nbsp;of&nbsp;week&nbsp;field，&nbsp;不支持秒级别。&nbsp;当type为Cron时，CronExpression字段生效。

	CronExpression *string `json:"CronExpression,omitempty" name:"CronExpression"`
}

type DeleteIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志时间,&nbsp;即SearchLog接口返回信息中Results结构体中的Time，需按照&nbsp;UTC+8&nbsp;时区将该毫秒级Unix时间戳转换为&nbsp;YYYY-mm-dd&nbsp;HH:MM:SS.FFF&nbsp;格式的字符串。

	BTime *string `json:"BTime,omitempty" name:"BTime"`
	// 日志包序号，即SearchLog接口返回信息中Results结构体中的PkgId。

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 日志包内一条日志的序号，即SearchLog接口返回信息中Results结构中的PkgLogId。

	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 前${PrevLogs}条日志，默认值10。

	PrevLogs *int64 `json:"PrevLogs,omitempty" name:"PrevLogs"`
	// 后${NextLogs}条日志，默认值10。

	NextLogs *int64 `json:"NextLogs,omitempty" name:"NextLogs"`
	// 检索语句，对日志上下文进行过滤，最大长度为12KB&nbsp;语句由&nbsp;[检索条件]构成，不支持SQL语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 上下文检索的开始时间，单位：毫秒级时间戳&nbsp;注意：&nbsp;-&nbsp;From为空时，表示上下文检索的开始时间不做限制&nbsp;-&nbsp;From和To非空时，From&nbsp;<&nbsp;To

	From *uint64 `json:"From,omitempty" name:"From"`
	// 上下文检索的结束时间，单位：毫秒级时间戳。&nbsp;注意：&nbsp;-&nbsp;To为空时，表示上下文检索的结束时间不做限制&nbsp;-&nbsp;From和To非空时，From&nbsp;<&nbsp;To

	To *uint64 `json:"To,omitempty" name:"To"`
}

func (r *DescribeLogContextRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogContextRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyUserConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type JsonInfo struct {

	// 启用标志。true表示启用TAG标志，false表示关闭TAG标志。

	EnableTag *bool `json:"EnableTag,omitempty" name:"EnableTag"`
	// 元数据信息列表,&nbsp;可选值为&nbsp;SOURCE、FILENAME、TIMESTAMP、HOSTNAME。

	MetaFields []*string `json:"MetaFields,omitempty" name:"MetaFields"`
	// 投递Json格式，0：字符串方式投递；1:以结构化方式投递

	JsonType *int64 `json:"JsonType,omitempty" name:"JsonType"`
}

type QueryTemplateItem struct {

	// 检索语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 检索语句名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeLogFastAnalysisRequest struct {
	*tchttp.BaseRequest

	// 要查询的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要查询的日志的起始时间，Unix时间戳，单位ms

	From *int64 `json:"From,omitempty" name:"From"`
	// 要查询的日志的结束时间，Unix时间戳，单位ms

	To *int64 `json:"To,omitempty" name:"To"`
	// 字段名

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 查询语句，语句长度最大为4096

	Query *string `json:"Query,omitempty" name:"Query"`
}

func (r *DescribeLogFastAnalysisRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogFastAnalysisRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyResourceAndFolderRelationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyResourceAndFolderRelationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyResourceAndFolderRelationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递规则列表

		Consumers []*ConsumerInfo `json:"Consumers,omitempty" name:"Consumers"`
		// 本次查询获取到的总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExtractRuleInfo struct {

	// 时间字段的key名字，TikeKey和TimeFormat必须成对出现

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// 时间字段的格式，参考c语言的strftime函数对于时间的格式说明输出参数

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 分隔符类型日志的分隔符，只有LogType为delimiter_log时有效

	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`
	// 整条日志匹配规则，只有LogType为fullregex_log时有效

	LogRegex *string `json:"LogRegex,omitempty" name:"LogRegex"`
	// 行首匹配规则，只有LogType为multiline_log或fullregex_log时有效

	BeginRegex *string `json:"BeginRegex,omitempty" name:"BeginRegex"`
	// 取的每个字段的key名字，为空的key代表丢弃这个字段，只有LogType为delimiter_log时有效，json_log的日志使用json本身的key。限制100个。

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 日志过滤规则列表（旧版），需要过滤日志的key，及其对应的regex。&nbsp;注意：2.9.3及以上版本LogListener&nbsp;，建议使用AdvanceFilterRules配置日志过滤规则。

	FilterKeyRegex []*KeyRegexInfo `json:"FilterKeyRegex,omitempty" name:"FilterKeyRegex"`
	// 解析失败日志是否上传，true表示上传，false表示不上传

	UnMatchUpLoadSwitch *bool `json:"UnMatchUpLoadSwitch,omitempty" name:"UnMatchUpLoadSwitch"`
	// 失败日志的key，当UnMatchUpLoadSwitch为true时必填

	UnMatchLogKey *string `json:"UnMatchLogKey,omitempty" name:"UnMatchLogKey"`
	// 增量采集模式下的回溯数据量，默认：-1（全量采集）；其他非负数表示增量采集（从最新的位置，往前采集${Backtracking}字节（Byte）的日志）最大支持1073741824（1G）。&nbsp;注意：&nbsp;&nbsp;COS导入不支持此字段。

	Backtracking *int64 `json:"Backtracking,omitempty" name:"Backtracking"`
	// 是否为Gbk编码。&nbsp;0：否；1：是。&nbsp;注意&nbsp;&nbsp;目前取0值时，表示UTF-8编码&nbsp;COS导入不支持此字段。

	IsGBK *int64 `json:"IsGBK,omitempty" name:"IsGBK"`
	// 是否为标准json。&nbsp;0：否；&nbsp;1：是。&nbsp;&nbsp;标准json指采集器使用业界标准开源解析器进行json解析，非标json指采集器使用CLS自研json解析器进行解析，两种解析器没有本质区别，建议客户使用标准json进行解析。

	JsonStandard *int64 `json:"JsonStandard,omitempty" name:"JsonStandard"`
	// syslog传输协议，取值为tcp或者udp，只有在LogType为service_syslog时生效，其余类型无需填写。&nbsp;注意：&nbsp;&nbsp;该字段适用于：创建采集规则配置、修改采集规则配置。&nbsp;COS导入不支持此字段。

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// syslog系统日志采集指定采集器监听的地址和端口&nbsp;，形式：[ip]:[port]，只有在LogType为service_syslog时生效，其余类型无需填写。&nbsp;注意：&nbsp;&nbsp;该字段适用于：创建采集规则配置、修改采集规则配置。&nbsp;COS导入不支持此字段。

	Address *string `json:"Address,omitempty" name:"Address"`
	// rfc3164：指定系统日志采集使用RFC3164协议解析日志。&nbsp;rfc5424：指定系统日志采集使用RFC5424协议解析日志。&nbsp;auto：自动匹配rfc3164或者rfc5424其中一种协议。&nbsp;只有在LogType为service_syslog时生效，其余类型无需填写。&nbsp;注意：&nbsp;&nbsp;该字段适用于：创建采集规则配置、修改采集规则配置&nbsp;COS导入不支持此字段。

	ParseProtocol *string `json:"ParseProtocol,omitempty" name:"ParseProtocol"`
	// 元数据类型。0:&nbsp;不使用元数据信息；1:使用机器组元数据；2:使用用户自定义元数据；3:使用采集配置路径。&nbsp;注意：&nbsp;&nbsp;COS导入不支持此字段。

	MetadataType *int64 `json:"MetadataType,omitempty" name:"MetadataType"`
	// 采集配置路径正则表达式。&nbsp;&nbsp;请用"()"标识路径中目标字段对应的正则表达式，解析时将"()"视为捕获组，并以__TAG__.{i}:{目标字段}的形式与日志一起上报，其中i为捕获组的序号。&nbsp;若不希望以序号为键名，可以通过命名捕获组"(?<{键名}>{正则})"自定义键名，并以__TAG__.{键名}:{目标字段}的形式与日志一起上报。最多支持5个捕获组&nbsp;注意：&nbsp;&nbsp;MetadataType为3时必填。&nbsp;COS导入不支持此字段。

	PathRegex *string `json:"PathRegex,omitempty" name:"PathRegex"`
	// 用户自定义元数据信息。&nbsp;注意：&nbsp;&nbsp;MetadataType为2时必填。&nbsp;COS导入不支持此字段。

	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`
	// Windows事件日志采集规则，只有在LogType为windows_event_log时生效，其余类型无需填写。

	EventLogRules []*EventLog `json:"EventLogRules,omitempty" name:"EventLogRules"`
	// 日志过滤规则列表（新版）。&nbsp;注意：&nbsp;&nbsp;2.9.3以下版本LogListener不支持，&nbsp;请使用FilterKeyRegex配置日志过滤规则。&nbsp;自建k8s采集配置（CreateConfigExtra、ModifyConfigExtra）不支持此字段。

	AdvanceFilterRules []*AdvanceFilterRuleInfo `json:"AdvanceFilterRules,omitempty" name:"AdvanceFilterRules"`
	// 原始日志的键名称(Key)；所有原始日志，&nbsp;均以您指定的键名称（Key），原始日志内容作为值（Value）进行上传，为空时表示不开启原始日志上传。&nbsp;&nbsp;COS导入不支持此字段。

	RawLogKey *string `json:"RawLogKey,omitempty" name:"RawLogKey"`
}

type CheckRechargeKafkaServerRequest struct {
	*tchttp.BaseRequest

	// 导入Kafka类型，0:&nbsp;CKafka，1:&nbsp;用户自建Kafka

	KafkaType *uint64 `json:"KafkaType,omitempty" name:"KafkaType"`
	// CKafka实例ID，KafkaType为0时必填

	KafkaInstance *string `json:"KafkaInstance,omitempty" name:"KafkaInstance"`
	// 服务地址。&nbsp;KafkaType为1时，ServerAddr必填

	ServerAddr *string `json:"ServerAddr,omitempty" name:"ServerAddr"`
	// ServerAddr是否为加密连接，默认值false。当KafkaType为1用户自建kafka时生效。

	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitempty" name:"IsEncryptionAddr"`
	// 加密访问协议。KafkaType参数为1并且IsEncryptionAddr参数为true时必填。

	Protocol *KafkaProtocolInfo `json:"Protocol,omitempty" name:"Protocol"`
	// 网络信息参数

	NetworkInfo *NetworkInfo `json:"NetworkInfo,omitempty" name:"NetworkInfo"`
}

func (r *CheckRechargeKafkaServerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRechargeKafkaServerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *DeleteLogsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DeleteMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShipperResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteShipperResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNoticeContentsRequest struct {
	*tchttp.BaseRequest

	// name&nbsp;按照【通知内容模板名称】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;noticeContentId&nbsp;按照【通知内容模板ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNoticeContentsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNoticeContentsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBinlogSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyBinlogSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBinlogSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilePathInfo struct {

	// 文件路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 文件名称

	File *string `json:"File,omitempty" name:"File"`
}

type CreateRebuildIndexTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 索引重建任务ID

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRebuildIndexTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRebuildIndexTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyCosRechargeRequest struct {
	*tchttp.BaseRequest

	// COS导入配置Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 日志主题Id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// cos导入任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否启用:&nbsp;&nbsp;&nbsp;0：&nbsp;未启用&nbsp;&nbsp;，&nbsp;1：启用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// COS存储桶，详见产品支持的存储桶命名规范。

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// COS存储桶所在地域，详见产品支持的地域列表。

	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
	// COS文件所在文件夹的前缀。为空串时投递存储桶下所有的文件。

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表单行全文；&nbsp;默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 解析格式。supported:&nbsp;"",&nbsp;"gzip",&nbsp;"lzop",&nbsp;"snappy";&nbsp;默认空

	Compress *string `json:"Compress,omitempty" name:"Compress"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitempty" name:"ExtractRuleInfo"`
	// COS导入任务类型。1：一次性导入任务；2：持续性导入任务。

	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
	// 元数据。支持&nbsp;bucket，object。

	Metadata []*string `json:"Metadata,omitempty" name:"Metadata"`
}

func (r *ModifyCosRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyCosRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 机器组类型。Type：ip，Values中为ip字符串列表机器组；Type：label，Values中为标签字符串列表机器组。

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
	// 标签列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启机器组自动更新。true表示开启自动升级，false表示关闭自动升级。

	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 升级开始时间，建议业务低峰期升级LogListener

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 升级结束时间，建议业务低峰期升级LogListener

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 是否开启服务日志，用于记录因Loglistener&nbsp;服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费

	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
	// TKE标志位，默认值为空字符串。空字符串表示日志不是来自于TKE，label_k8s表示日志来自于TKE。

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// 机器组中机器定期离线清理时间

	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitempty" name:"DelayCleanupTime"`
	// 机器组元数据信息列表

	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`
	// 集群ID。-&nbsp;当Flag为&nbsp;label_tke，表示TKE集群ID-&nbsp;ClusterId&nbsp;与&nbsp;ClusterRegion&nbsp;必须同时填写，或者为空

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群所在地域-&nbsp;当Flag为&nbsp;label_tke，表示TKE集群地域-&nbsp;ClusterId&nbsp;与&nbsp;ClusterRegion&nbsp;必须同时填写，或者为空

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
}

func (r *ModifyMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperPreviewRequest struct {
	*tchttp.BaseRequest

	// 创建的投递规则所属的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 投递日志的内容格式配置

	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
}

func (r *DescribeShipperPreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShipperPreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志源类型。支持&nbsp;container_stdout：容器标准输出；container_file：容器文件路径；host_file：节点文件路径。

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点文件路径类型配置。

	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`
	// 容器文件路径类型配置。

	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`
	// 容器标准输出类型配置。

	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`
	// 采集的日志类型，默认为minimalist_log。支持以下类型：&nbsp;&nbsp;json_log代表：JSON-文件日志（详见使用&nbsp;JSON&nbsp;提取模式采集日志）；&nbsp;delimiter_log代表：分隔符-文件日志（详见使用分隔符提取模式采集日志）；&nbsp;minimalist_log代表：单行全文-文件日志（详见使用单行全文提取模式采集日志）；&nbsp;fullregex_log代表：单行完全正则-文件日志（详见使用单行-完全正则提取模式采集日志）；&nbsp;multiline_log代表：多行全文-文件日志（详见使用多行全文提取模式采集日志）；&nbsp;multiline_fullregex_log代表：多行完全正则-文件日志（详见使用多行-完全正则提取模式采集日志）；&nbsp;user_define_log代表：组合解析（适用于多格式嵌套的日志，详见使用组合解析提取模式采集日志）。

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 日志格式化方式，用于容器采集场景。&nbsp;-&nbsp;已废弃&nbsp;&nbsp;stdout-docker-json：用于docker容器采集场景&nbsp;stdout-containerd：用于containerd容器采集场景

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 组合解析采集规则，用于复杂场景下的日志采集。&nbsp;&nbsp;取值参考：使用组合解析提取模式采集日志

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// 绑定的机器组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 采集配置标记。&nbsp;&nbsp;目前只支持label_k8s，用于标记自建k8s集群使用的采集配置

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集name

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 绑定的机器组id列表

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// 采集相关配置信息。详情见CollectInfo复杂类型配置。

	CollectInfos []*CollectInfo `json:"CollectInfos,omitempty" name:"CollectInfos"`
	// 高级采集配置。&nbsp;Json字符串，&nbsp;Key/Value定义为如下：&nbsp;&nbsp;ClsAgentFileTimeout(超时属性),&nbsp;取值范围:&nbsp;大于等于0的整数，&nbsp;0为不超时&nbsp;ClsAgentMaxDepth(最大目录深度)，取值范围:&nbsp;大于等于0的整数&nbsp;ClsAgentParseFailMerge(合并解析失败日志)，取值范围:&nbsp;true或false&nbsp;ClsAgentDefault(自定义默认值，无特殊含义，用于清空其他选项)，建议取值0

	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
}

func (r *CreateConfigExtraRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigExtraRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateShipperResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递规则ID

		ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateShipperResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateShipperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomMetricSpec struct {

	// 端口。取值范围&nbsp;[1,65535]

	Port *string `json:"Port,omitempty" name:"Port"`
	// Metric地址。校验格式：`^/[a-zA-Z0-9-_./]*$`

	Path *string `json:"Path,omitempty" name:"Path"`
	// 命名空间列表。&nbsp;-&nbsp;最大支持100个&nbsp;-&nbsp;namespace&nbsp;校验格式&nbsp;`[a-z0-9]([-a-z0-9]*[a-z0-9])?`&nbsp;，&nbsp;长度不能超过63&nbsp;-&nbsp;namespace&nbsp;不能重复

	Namespaces []*string `json:"Namespaces,omitempty" name:"Namespaces"`
	// Pod标签。&nbsp;-&nbsp;最大支持100个

	PodLabel []*Label `json:"PodLabel,omitempty" name:"PodLabel"`
}

type DeleteAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmNoticeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmNoticeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集配置列表

		Configs []*ConfigInfo `json:"Configs,omitempty" name:"Configs"`
		// 过滤到的总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CsvInfo struct {

	// csv首行是否打印key。true表示首行打印key，false表示首行不打印key

	PrintKey *bool `json:"PrintKey,omitempty" name:"PrintKey"`
	// 每列key的名字

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 各字段间的分隔符

	Delimiter *string `json:"Delimiter,omitempty" name:"Delimiter"`
	// 若字段内容中包含分隔符，则使用该转义符包裹改字段，只能填写单引号、双引号、空字符串

	EscapeChar *string `json:"EscapeChar,omitempty" name:"EscapeChar"`
	// 对于上面指定的不存在字段使用该内容填充

	NonExistingField *string `json:"NonExistingField,omitempty" name:"NonExistingField"`
}

type ModifyDashboardSubscribeAckRequest struct {
	*tchttp.BaseRequest

	// 仪表盘订阅id。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 仪表盘订阅发送成功时间。

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 仪表盘订阅发送的状态。success：全部发送成功，fail：未发送，&nbsp;partialSuccess：部分发送成功。

	LastStatus *string `json:"LastStatus,omitempty" name:"LastStatus"`
}

func (r *ModifyDashboardSubscribeAckRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardSubscribeAckRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeRule struct {

	// 匹配规则JSON串

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 告警通知接收者信息。

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 告警通知模板回调信息，包括企业微信、钉钉、飞书。

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 告警升级开关。true：开启告警升级、false：关闭告警升级，默认：false

	Escalate *bool `json:"Escalate,omitempty" name:"Escalate"`
	// 告警升级条件。1：无人认领且未恢复、2：未恢复，默认为1&nbsp;&nbsp;无人认领且未恢复：告警没有恢复并且没有人认领则升级&nbsp;未恢复：当前告警持续未恢复则升级

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 告警升级间隔。单位：分钟，范围[1，14400]

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 告警升级后下一个环节的通知渠道配置

	EscalateNotice *EscalateNoticeInfo `json:"EscalateNotice,omitempty" name:"EscalateNotice"`
}

type DeleteDataTransformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDataTransformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteScheduledSqlRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 源日志主题ID

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
}

func (r *DeleteScheduledSqlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteScheduledSqlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestUserLogRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志展示处理方式，默认为1&nbsp;1：自动配置索引数据字段处理&nbsp;2：cos投递csv格式键值索引配置/Parquet字段展示字段处理

	ProcessMode *int64 `json:"ProcessMode,omitempty" name:"ProcessMode"`
}

func (r *DescribeLatestUserLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestUserLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// 采集配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
}

func (r *DescribeConfigMachineGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigMachineGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmNoticeDeliverConfig struct {

	// 通知渠道投递日志配置信息。

	DeliverConfig *DeliverConfig `json:"DeliverConfig,omitempty" name:"DeliverConfig"`
	// 投递失败原因。

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type JsonLogInfo struct {

	// K-V形式日志信息

	Log []*KeyLogInfo `json:"Log,omitempty" name:"Log"`
	// K-V形式标签信息

	Tag []*KeyLogInfo `json:"Tag,omitempty" name:"Tag"`
	// 时间戳

	Time *uint64 `json:"Time,omitempty" name:"Time"`
}

type ApplyConfigToMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 采集配置ID

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *ApplyConfigToMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyConfigToMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataTransformResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDataTransformResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataTransformResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type BinlogDBConfig struct {

	// 访问方式。&nbsp;1:&nbsp;内网&nbsp;2:&nbsp;外网

	AccessMode *uint64 `json:"AccessMode,omitempty" name:"AccessMode"`
	// 用户名。

	User *string `json:"User,omitempty" name:"User"`
	// 密码。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 实例id。&nbsp;数据库类型是云数据库Mysql&nbsp;或者&nbsp;TDSQL-C时必填。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// mysql实例地址。&nbsp;自建MySQL，且访问方式为内网访问时必填。

	Address *string `json:"Address,omitempty" name:"Address"`
	// mysql实例端口。&nbsp;自建MySQL，且访问方式为内网访问时必填。

	Port *string `json:"Port,omitempty" name:"Port"`
	// 私有网络服务ID。&nbsp;自建MySQL，且访问方式为内网访问时必填。

	EndPointServiceId *string `json:"EndPointServiceId,omitempty" name:"EndPointServiceId"`
	// 网络服务类型。&nbsp;1:负载均衡CLB&nbsp;2:云服务器CVM&nbsp;自建MySQL，且访问方式为内网访问时必填。

	NetworkServiceType *uint64 `json:"NetworkServiceType,omitempty" name:"NetworkServiceType"`
}

type TopicWhitelistInfo struct {

	// 当cql返回不为空时，&nbsp;true:可以使用CQL语法，&nbsp;false:不可以使用CQL语法。字段不存在时，不能使用CQL语法。

	CloudQueryLanguage *bool `json:"CloudQueryLanguage,omitempty" name:"CloudQueryLanguage"`
	// 二进制标记位。前端控制台使用。&nbsp;第一位：是否打免费标签

	FlagBit *uint64 `json:"FlagBit,omitempty" name:"FlagBit"`
}

type DescribeAgentMachineGroupMetadataResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器组元数据列表

		GroupMetadata []*MachineGroupMetadataInfo `json:"GroupMetadata,omitempty" name:"GroupMetadata"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentMachineGroupMetadataResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentMachineGroupMetadataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKafkaConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyKafkaConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadServiceLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadServiceLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadServiceLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigurationTemplateRequest struct {
	*tchttp.BaseRequest

	// 模版名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z,&nbsp;_,-,zh字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注，可选参数，长度不超过255个字符

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 投递cos模板配置项

	ShipperTemplateInfos []*ShipperTemplateInfo `json:"ShipperTemplateInfos,omitempty" name:"ShipperTemplateInfos"`
}

func (r *CreateConfigurationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigurationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigExtrasRequest struct {
	*tchttp.BaseRequest

	// 过滤器，支持如下选项：&nbsp;name&nbsp;&nbsp;按照【特殊采集配置名称】进行模糊匹配过滤。&nbsp;类型：String&nbsp;configExtraId&nbsp;&nbsp;按照【特殊采集配置ID】进行过滤。&nbsp;类型：String&nbsp;topicId&nbsp;&nbsp;按照【日志主题】进行过滤。&nbsp;类型：String&nbsp;machineGroupId&nbsp;&nbsp;按照【机器组ID】进行过滤。&nbsp;类型：String&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigExtrasRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigExtrasRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDataTransformRequest struct {
	*tchttp.BaseRequest

	// 任务类型.&nbsp;1:&nbsp;指定主题；2:动态创建。详情请参考创建加工任务文档。

	FuncType *int64 `json:"FuncType,omitempty" name:"FuncType"`
	// 源日志主题

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 加工任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 加工语句。&nbsp;当FuncType为2时，EtlContent必须使用log_auto_output&nbsp;&nbsp;其他参考文档：&nbsp;&nbsp;创建加工任务&nbsp;函数总览

	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
	// 任务启动状态.&nbsp;默认为1:开启,&nbsp;2:关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 加工任务目的topic_id以及别名,当FuncType=1时，该参数必填，当FuncType=2时，无需填写。

	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
	// 加工类型。&nbsp;1：使用源日志主题中的随机数据，进行加工预览；2：使用用户自定义测试数据，进行加工预览；3：创建真实加工任务。

	TaskType *int64 `json:"TaskType,omitempty" name:"TaskType"`
	// 用于预览加工结果的测试数据

	PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitempty" name:"PreviewLogStatistics"`
	// 当FuncType为2时，动态创建的日志集、日志主题的个数超出产品规格限制是否丢弃数据，&nbsp;默认为false。&nbsp;&nbsp;false：创建兜底日志集、日志主题并将日志写入兜底主题；&nbsp;true：丢弃日志数据。

	BackupGiveUpData *bool `json:"BackupGiveUpData,omitempty" name:"BackupGiveUpData"`
	// 交互式加工语句配置内容

	InteractiveEtlContent *string `json:"InteractiveEtlContent,omitempty" name:"InteractiveEtlContent"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
	// 数据加工类型。0：标准加工任务；&nbsp;1：前置加工任务。前置加工任务将采集的日志处理完成后，再写入日志主题。

	DataTransformType *uint64 `json:"DataTransformType,omitempty" name:"DataTransformType"`
	// 保留失败日志状态。&nbsp;1:不保留（默认)，2:保留

	KeepFailureLog *uint64 `json:"KeepFailureLog,omitempty" name:"KeepFailureLog"`
	// 失败日志的字段名称

	FailureLogKey *string `json:"FailureLogKey,omitempty" name:"FailureLogKey"`
}

func (r *CreateDataTransformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDataTransformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘订阅记录Id

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateDashboardSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardRequest struct {
	*tchttp.BaseRequest

	// 仪表盘id

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
}

func (r *DeleteDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryShipperTaskRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// 投递任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *RetryShipperTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryShipperTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateWebCallbackRequest struct {
	*tchttp.BaseRequest

	// 通知内容名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 渠道类型。&nbsp;&nbsp;WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调。

	Type *string `json:"Type,omitempty" name:"Type"`
	// Webhook地址。

	Webhook *string `json:"Webhook,omitempty" name:"Webhook"`
	// 请求方式。&nbsp;支持POST、PUT。&nbsp;&nbsp;当Type为Http时，必填。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 秘钥。

	Key *string `json:"Key,omitempty" name:"Key"`
}

func (r *CreateWebCallbackRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateWebCallbackRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerGroupRequest struct {
	*tchttp.BaseRequest

	// 更新的目标消费者组标识

	ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
	// 消费者心跳超时时间（秒）

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 更新的消费者组包含的日志主题列表

	Topics []*string `json:"Topics,omitempty" name:"Topics"`
	// 日志集Id（日志主题所属的日志集）

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *ModifyConsumerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogInfo struct {

	// 日志时间，单位ms

	Time *int64 `json:"Time,omitempty" name:"Time"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 日志来源IP

	Source *string `json:"Source,omitempty" name:"Source"`
	// 日志文件名称

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 日志上报请求包的ID

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 请求包内日志的ID

	PkgLogId *string `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 日志内容，由多个LogItem&nbsp;(KV结构）组成

	Logs []*LogItem `json:"Logs,omitempty" name:"Logs"`
	// 日志内容的高亮描述信息

	HighLights []*HighLightItem `json:"HighLights,omitempty" name:"HighLights"`
	// 日志内容的Json序列化字符串

	LogJson *string `json:"LogJson,omitempty" name:"LogJson"`
	// 日志来源主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 原始日志(仅在日志创建索引异常时有值)

	RawLog *string `json:"RawLog,omitempty" name:"RawLog"`
	// 日志创建索引异常原因(仅在日志创建索引异常时有值)

	IndexStatus *string `json:"IndexStatus,omitempty" name:"IndexStatus"`
}

type SearchLogInfos struct {

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志存储生命周期

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 冷热边界时间点（ms）

	Border *int64 `json:"Border,omitempty" name:"Border"`
	// 透传本次接口返回的Context值，可获取后续更多日志，过期时间1小时

	Context *string `json:"Context,omitempty" name:"Context"`
}

type ConfigExtraInfo struct {

	// 采集规则扩展配置ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集规则名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 类型：container_stdout、container_file、host_file

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点文件配置信息

	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`
	// 容器文件路径信息

	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`
	// 容器标准输出信息

	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`
	// 日志格式化方式

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表极简日志，multiline_log代表多行日志，fullregex_log代表完整正则，默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 用户自定义解析字符串

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 自建采集配置标

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集name

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题name

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 采集相关配置信息。详情见&nbsp;CollectInfo复杂类型配置。

	CollectInfos []*CollectInfo `json:"CollectInfos,omitempty" name:"CollectInfos"`
	// 高级采集配置。&nbsp;Json字符串，&nbsp;Key/Value定义为如下：&nbsp;&nbsp;ClsAgentFileTimeout(超时属性),&nbsp;取值范围:&nbsp;大于等于0的整数，&nbsp;0为不超时&nbsp;ClsAgentMaxDepth(最大目录深度)，取值范围:&nbsp;大于等于0的整数&nbsp;ClsAgentParseFailMerge(合并解析失败日志)，取值范围:&nbsp;true或false&nbsp;样例：{"ClsAgentFileTimeout":0,"ClsAgentMaxDepth":10,"ClsAgentParseFailMerge":true}

	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
}

type ContainerFileInfo struct {

	// namespace可以多个，用分隔号分割,例如A,B

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 容器名称

	Container *string `json:"Container,omitempty" name:"Container"`
	// 日志文件夹

	LogPath *string `json:"LogPath,omitempty" name:"LogPath"`
	// 日志名称

	FilePattern *string `json:"FilePattern,omitempty" name:"FilePattern"`
	// pod标签信息

	IncludeLabels []*string `json:"IncludeLabels,omitempty" name:"IncludeLabels"`
	// 工作负载信息

	WorkLoad *ContainerWorkLoadInfo `json:"WorkLoad,omitempty" name:"WorkLoad"`
	// 需要排除的namespace可以多个，用分隔号分割,例如A,B

	ExcludeNamespace *string `json:"ExcludeNamespace,omitempty" name:"ExcludeNamespace"`
	// 需要排除的pod标签信息

	ExcludeLabels []*string `json:"ExcludeLabels,omitempty" name:"ExcludeLabels"`
	// 容器名称标记。必填字段，不填默认值为0。&nbsp;&nbsp;默认值0，0:选中Container标记。&nbsp;1:排除Container标记。

	ContainerFlag *uint64 `json:"ContainerFlag,omitempty" name:"ContainerFlag"`
	// 日志文件信息

	FilePaths []*FilePathInfo `json:"FilePaths,omitempty" name:"FilePaths"`
	// metadata信息

	CustomLabels []*string `json:"CustomLabels,omitempty" name:"CustomLabels"`
}

type ExternalDataSourceCOSInfo struct {

	// COS地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// COS存储桶

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// COS文件地址(CSV格式)，仅支持CSV格式文件

	FileAddress *string `json:"FileAddress,omitempty" name:"FileAddress"`
	// 压缩方式&nbsp;0:不压缩;1:gzip;&nbsp;2:lzop;3:snappy

	CompressType *uint64 `json:"CompressType,omitempty" name:"CompressType"`
	// COS文件(CSV格式)内容&nbsp;表头字段与字段类型信息

	FieldsInfo []*FieldInfo `json:"FieldsInfo,omitempty" name:"FieldsInfo"`
}

type ModifyDashboardSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘订阅id。

		Id *uint64 `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerWorkLoadInfo struct {

	// 容器名

	Container *string `json:"Container,omitempty" name:"Container"`
	// 工作负载的类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 工作负载的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 容器名称标记。必填字段，不填默认值为0。&nbsp;&nbsp;0:选中Container标记。默认值&nbsp;1:排除Container标记。

	ContainerFlag *uint64 `json:"ContainerFlag,omitempty" name:"ContainerFlag"`
}

type Relabeling struct {

	// 基于正则表达式匹配执行的动作。&nbsp;-&nbsp;replace:&nbsp;Label替换,&nbsp;必填:&nbsp;SourceLabels,&nbsp;Separator,&nbsp;Regex,&nbsp;TargetLabel,&nbsp;Replacement&nbsp;-&nbsp;labeldrop:&nbsp;丢弃Label,&nbsp;必填:&nbsp;Regex&nbsp;-&nbsp;labelkeep:&nbsp;保留Label,&nbsp;必填:&nbsp;Regex&nbsp;-&nbsp;lowercase:&nbsp;小写化,&nbsp;必填:&nbsp;SourceLabels,&nbsp;Separator,&nbsp;TargetLabel&nbsp;-&nbsp;uppercase:&nbsp;大写化,&nbsp;必填:&nbsp;SourceLabels,&nbsp;Separator,&nbsp;TargetLabel&nbsp;-&nbsp;dropequal:&nbsp;丢弃指标-完全匹配,&nbsp;必填:&nbsp;SourceLabels,&nbsp;Separator,&nbsp;TargetLabel&nbsp;-&nbsp;keepequal:&nbsp;保留指标-完全匹配,&nbsp;必填:&nbsp;SourceLabels,&nbsp;Separator,&nbsp;TargetLabel&nbsp;-&nbsp;drop:&nbsp;丢弃指标-正则匹配,&nbsp;必填:&nbsp;SourceLabels,&nbsp;Separator,&nbsp;Regex&nbsp;-&nbsp;keep:&nbsp;保留指标-正则匹配,&nbsp;必填:&nbsp;SourceLabels,&nbsp;Separator,&nbsp;Regex&nbsp;-&nbsp;hashmod:哈希取模,&nbsp;必填:&nbsp;SourceLabels,&nbsp;Separator,&nbsp;TargetLabel,&nbsp;Modulus&nbsp;-&nbsp;labelmap:Label映射,&nbsp;必填:&nbsp;Regex,&nbsp;Replacement

	Action *string `json:"Action,omitempty" name:"Action"`
	// 原始label

	SourceLabels []*string `json:"SourceLabels,omitempty" name:"SourceLabels"`
	// 原始label连接符。&nbsp;必填时不能为空串，&nbsp;长度不能超过256

	Separator *string `json:"Separator,omitempty" name:"Separator"`
	// 目标label。必填时不能为空串，校验格式：`^[a-zA-Z_][a-zA-Z0-9_]*$`&nbsp;，&nbsp;长度不能超过256

	TargetLabel *string `json:"TargetLabel,omitempty" name:"TargetLabel"`
	// 替换值。如果正则表达式匹配，则对其执行替换操作。&nbsp;-&nbsp;必填时不能为空串，长度不能超过256&nbsp;-&nbsp;当action为LabelMap时，&nbsp;Replacement&nbsp;校验格式：`^(?:(?:[a-zA-Z_]|\$(?:\{\w+\}|\w+))+\w*)+$`

	Replacement *string `json:"Replacement,omitempty" name:"Replacement"`
	// 正则表达式。提取与之匹配值。必填时不能为空串，校验必须是一个合法的&nbsp;RE2

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 获取源标签值的哈希值。必填时不能为空,不能为0

	Modulus *uint64 `json:"Modulus,omitempty" name:"Modulus"`
}

type TopicIdLogStatistic struct {

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 读取的源日志主题的行数

	ReadLines *int64 `json:"ReadLines,omitempty" name:"ReadLines"`
	// 加工后输出到目标日志主题的行数

	WriteLines *int64 `json:"WriteLines,omitempty" name:"WriteLines"`
	// 加工失败的行数

	FailedLines *int64 `json:"FailedLines,omitempty" name:"FailedLines"`
	// 加工过滤的行数

	FilterLines *uint64 `json:"FilterLines,omitempty" name:"FilterLines"`
}

type DashboardInfo struct {

	// 仪表盘id

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘名字

	DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
	// 仪表盘数据

	Data *string `json:"Data,omitempty" name:"Data"`
	// 创建仪表盘的时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// AssumerUin非空则表示创建该日志主题的服务方Uin

	AssumerUin *uint64 `json:"AssumerUin,omitempty" name:"AssumerUin"`
	// RoleName非空则表示创建该日志主题的服务方使用的角色

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// AssumerName非空则表示创建该日志主题的服务方名称

	AssumerName *string `json:"AssumerName,omitempty" name:"AssumerName"`
	// 日志主题绑定的标签信息

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 仪表盘所在地域：&nbsp;为了兼容老的地域。

	DashboardRegion *string `json:"DashboardRegion,omitempty" name:"DashboardRegion"`
	// 修改仪表盘的时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 仪表盘对应的topic相关信息

	DashboardTopicInfos []*DashboardTopicInfo `json:"DashboardTopicInfos,omitempty" name:"DashboardTopicInfos"`
	// 文件夹名称

	FolderName *string `json:"FolderName,omitempty" name:"FolderName"`
	// 文件夹id

	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

type PrivateDomainNames struct {

	// 域名地址

	DomainName *string `json:"DomainName,omitempty" name:"DomainName"`
	// ip地址

	IpAddr *string `json:"IpAddr,omitempty" name:"IpAddr"`
}

type TopicETL struct {

	// 用户id

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 标签

	Tag *string `json:"Tag,omitempty" name:"Tag"`
	// 创建时间（秒级时间戳）

	CreatedAt *int64 `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 更新时间（秒级时间戳）

	UpdatedAt *int64 `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 删除时间（秒级时间戳）

	DeletedAt *int64 `json:"DeletedAt,omitempty" name:"DeletedAt"`
}

type DeleteTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
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

type DescribeFoldersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 文件夹信息列表。

		Data []*FolderInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFoldersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFoldersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardSubscribeAckResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDashboardSubscribeAckResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardSubscribeAckResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
}

func (r *DeleteAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExceptionResourcesRequest struct {
	*tchttp.BaseRequest

	// 异常信息列表

	ExceptionInfo []*ExceptionInfo `json:"ExceptionInfo,omitempty" name:"ExceptionInfo"`
	// 资源信息列表

	ResourceInfo []*ResourceInfo `json:"ResourceInfo,omitempty" name:"ResourceInfo"`
}

func (r *DescribeExceptionResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExceptionResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FullTextInfo struct {

	// 是否大小写敏感。true表示大小写敏感，false代表大小写不敏感。

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 全文索引的分词符，其中的每个字符代表一个分词符；&nbsp;仅支持英文符号、\n\t\r及转义符\；&nbsp;注意：\n\t\r本身已被转义，直接使用双引号包裹即可作为入参，无需再次转义。使用API&nbsp;Explorer进行调试时请使用JSON参数输入方式，以避免\n\t\r被重复转义

	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`
	// 是否包含zh。true代表包含zh，false代表不包含zh。

	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
}

type DeleteAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckRechargeKafkaServerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Kafka集群可访问状态，0：可正常访问，其他值为报错code

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckRechargeKafkaServerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRechargeKafkaServerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigurationTemplateRequest struct {
	*tchttp.BaseRequest

	// 配置模板id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
}

func (r *DeleteConfigurationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigurationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchDashboardSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchDashboardSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDashboardSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerPreviewRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 是否投递日志的元数据信息，默认为&nbsp;true

	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`
	// 如果需要投递元数据信息，元数据信息的描述

	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`
}

func (r *DescribeConsumerPreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerPreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HighLightItem struct {

	// 高亮的日志Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 高亮的语法

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ResourcesInfo struct {

	// 地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 日志集数

	Logsets *uint64 `json:"Logsets,omitempty" name:"Logsets"`
	// 日志主题数

	Topics *uint64 `json:"Topics,omitempty" name:"Topics"`
	// 分区数

	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`
	// 机器数

	Machines *uint64 `json:"Machines,omitempty" name:"Machines"`
	// 心跳正常机器数

	HeartMachines *uint64 `json:"HeartMachines,omitempty" name:"HeartMachines"`
	// 指标主题数

	MetricTopics *uint64 `json:"MetricTopics,omitempty" name:"MetricTopics"`
}

type CreateDashboardSubscribeRequest struct {
	*tchttp.BaseRequest

	// 仪表盘订阅名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 仪表盘id。

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
	// 订阅时间cron表达式，格式为：{秒数}&nbsp;{分钟}&nbsp;{小时}&nbsp;{日期}&nbsp;{月份}&nbsp;{星期}；（有效数据为：{分钟}&nbsp;{小时}&nbsp;{日期}&nbsp;{月份}&nbsp;{星期}）。&nbsp;&nbsp;{秒数}&nbsp;取值范围：&nbsp;0&nbsp;~&nbsp;59&nbsp;{分钟}&nbsp;取值范围：&nbsp;0&nbsp;~&nbsp;59&nbsp;&nbsp;{小时}&nbsp;取值范围：&nbsp;0&nbsp;~&nbsp;23&nbsp;&nbsp;{日期}&nbsp;取值范围：&nbsp;1&nbsp;~&nbsp;31&nbsp;AND&nbsp;(dayOfMonth最后一天：&nbsp;L)&nbsp;{月份}&nbsp;取值范围：&nbsp;1&nbsp;~&nbsp;12&nbsp;{星期}&nbsp;取值范围：&nbsp;0&nbsp;~&nbsp;6&nbsp;【0:星期日，&nbsp;6星期六】

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 仪表盘订阅数据。

	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitempty" name:"SubscribeData"`
}

func (r *CreateDashboardSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigurationTemplateLogsetInfo struct {

	// 日志集id列表

	LogsetIds []*string `json:"LogsetIds,omitempty" name:"LogsetIds"`
	// 动态配置范围。1:新增日志主题,2:已有日志主题。

	MatchRange []*uint64 `json:"MatchRange,omitempty" name:"MatchRange"`
	// 日志主题名称正则匹配表达式。

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 该参数控制&nbsp;Regex参数是否正向匹配。&nbsp;<br>&nbsp;false：Regex为日志主题名称正则匹配表达式。&nbsp;<br>&nbsp;true：Regex为日志主题名称正则不匹配表达式。

	UnMatch *bool `json:"UnMatch,omitempty" name:"UnMatch"`
}

type ConsumerContent struct {

	// 是否投递&nbsp;TAG&nbsp;信息。&nbsp;当EnableTag为true时，表示投递TAG元信息。

	EnableTag *bool `json:"EnableTag,omitempty" name:"EnableTag"`
	// 需要投递的元数据列表，目前仅支持：__SOURCE__，__FILENAME__，__TIMESTAMP__，__HOSTNAME__和__PKGID__

	MetaFields []*string `json:"MetaFields,omitempty" name:"MetaFields"`
	// 当EnableTag为true时，必须填写TagJsonNotTiled字段。&nbsp;TagJsonNotTiled用于标识tag信息是否json平铺。&nbsp;&nbsp;TagJsonNotTiled为true时不平铺，示例：&nbsp;TAG信息：{"__TAG__":{"fieldA":200,"fieldB":"text"}}&nbsp;不平铺：{"__TAG__":{"fieldA":200,"fieldB":"text"}}&nbsp;&nbsp;TagJsonNotTiled为false时平铺，示例：&nbsp;TAG信息：{"__TAG__":{"fieldA":200,"fieldB":"text"}}&nbsp;平铺：{"__TAG__.fieldA":200,"__TAG__.fieldB":"text"}

	TagJsonNotTiled *bool `json:"TagJsonNotTiled,omitempty" name:"TagJsonNotTiled"`
	// 投递时间戳精度，可选项&nbsp;[1：秒；2：毫秒]&nbsp;，默认是1。

	TimestampAccuracy *int64 `json:"TimestampAccuracy,omitempty" name:"TimestampAccuracy"`
	// 投递Json格式。&nbsp;JsonType为0：和原始日志一致，不转义。示例：&nbsp;日志原文：{"a":"aa",&nbsp;"b":{"b1":"b1b1",&nbsp;"c1":"c1c1"}}&nbsp;投递到Ckafka：{"a":"aa",&nbsp;"b":{"b1":"b1b1",&nbsp;"c1":"c1c1"}}&nbsp;&nbsp;JsonType为1：转义。示例：&nbsp;日志原文：{"a":"aa",&nbsp;"b":{"b1":"b1b1",&nbsp;"c1":"c1c1"}}&nbsp;投递到Ckafka：{"a":"aa","b":"{\"b1\":\"b1b1\",&nbsp;\"c1\":\"c1c1\"}"}

	JsonType *int64 `json:"JsonType,omitempty" name:"JsonType"`
}

type CreateTopicExtendConfigRequest struct {
	*tchttp.BaseRequest

	// clb的topic业务配置,数组大小不可以超过100

	ClbTopicExtendConfigs []*ClbTopicExtendConfig `json:"ClbTopicExtendConfigs,omitempty" name:"ClbTopicExtendConfigs"`
}

func (r *CreateTopicExtendConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicExtendConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteWebCallbackResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWebCallbackResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteWebCallbackResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScheduledSqlProcessInfoRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 源日志主题ID

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 实例ID

	ProcessId *string `json:"ProcessId,omitempty" name:"ProcessId"`
	// 执行开始时间戳，单位ms

	ProcessStartTime *int64 `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 执行结束时间戳，单位ms

	ProcessEndTime *int64 `json:"ProcessEndTime,omitempty" name:"ProcessEndTime"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 调度结果，1:运行中&nbsp;2:成功&nbsp;3:失败.&nbsp;默认为0不做过滤

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 排序字段，可选字段为：process_start_time&nbsp;|&nbsp;time_window_start_time。&nbsp;默认为process_start_time

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序顺序，DESC&nbsp;|&nbsp;ASC。&nbsp;默认DESC

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeScheduledSqlProcessInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScheduledSqlProcessInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRangeMetricRequest struct {
	*tchttp.BaseRequest

	// 指标主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 查询语句，使用PromQL语法

	Query *string `json:"Query,omitempty" name:"Query"`
	// 查询起始时间，秒级Unix时间戳

	Start *uint64 `json:"Start,omitempty" name:"Start"`
	// 查询结束时间，秒级Unix时间戳

	End *uint64 `json:"End,omitempty" name:"End"`
	// 查询时间间隔，单位秒

	Step *uint64 `json:"Step,omitempty" name:"Step"`
}

func (r *QueryRangeMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRangeMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PartitionInfoForMonitor struct {

	// 分区id

	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

type DescribeKafkaConsumerPreviewRequest struct {
	*tchttp.BaseRequest

	// 主题名称

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
	// kafka协议消费数据格式

	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitempty" name:"ConsumerContent"`
}

func (r *DescribeKafkaConsumerPreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumerPreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分区列表

		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePartitionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePartitionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceConfig struct {

	// 实例维度

	InstanceDimension []*string `json:"InstanceDimension,omitempty" name:"InstanceDimension"`
	// 实例值

	Instances []*Instance `json:"Instances,omitempty" name:"Instances"`
}

type RebuildIndexTaskInfo struct {

	// 索引重建任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 索引重建任务当前状态，0:索引重建任务已创建，1:创建索引重建资源，2:索引重建资源创建完成，3:重建中，4:暂停，5:重建索引成功，6:重建成功（可检索），7:重建失败，8:撤销，9:删除元数据和索引

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 重建任务开始时间戳

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 重建任务结束时间戳

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 重投预估剩余时间，单位秒

	RemainTime *int64 `json:"RemainTime,omitempty" name:"RemainTime"`
	// 重建任务创建时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 重投完成度，百分比

	Progress *float64 `json:"Progress,omitempty" name:"Progress"`
	// 重建任务更新时间

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 附加状态描述信息（目前仅描述失败时失败原因）

	StatusMessage *string `json:"StatusMessage,omitempty" name:"StatusMessage"`
}

type DescribeDataTransformAutoCreateLogSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题总数

		TotalLogSetCount *uint64 `json:"TotalLogSetCount,omitempty" name:"TotalLogSetCount"`
		// 日志主题信息列表

		LogSets []*LogsetETL `json:"LogSets,omitempty" name:"LogSets"`
		// 重复的日志主题总数

		DuplicateLogSetCount *uint64 `json:"DuplicateLogSetCount,omitempty" name:"DuplicateLogSetCount"`
		// 重复的日志主题信息列表

		DuplicateLogSets []*LogsetETL `json:"DuplicateLogSets,omitempty" name:"DuplicateLogSets"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformAutoCreateLogSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformAutoCreateLogSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KeyLogInfo struct {

	// 日志key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 日志内容

	Value *string `json:"Value,omitempty" name:"Value"`
}

type CloseKafkaConsumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseKafkaConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKafkaConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIndexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIndexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryScheduledSqlTaskRequest struct {
	*tchttp.BaseRequest

	// 源日志主题ID

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 实例ID

	ProcessId *string `json:"ProcessId,omitempty" name:"ProcessId"`
}

func (r *RetryScheduledSqlTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryScheduledSqlTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateAgentStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAgentStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateAgentStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordingRuleProcessInfo struct {

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 实例ID

	ProcessId *string `json:"ProcessId,omitempty" name:"ProcessId"`
	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 目标日志主题id。

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 预聚合语句

	RecordingRuleContent *string `json:"RecordingRuleContent,omitempty" name:"RecordingRuleContent"`
	// 执行时间-开始时间

	ProcessStartTime *string `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 执行时间-结束时间

	ProcessEndTime *string `json:"ProcessEndTime,omitempty" name:"ProcessEndTime"`
	// 执行时间-耗时

	ProcessDuration *int64 `json:"ProcessDuration,omitempty" name:"ProcessDuration"`
	// 处理数据量-输入行数

	ReadLogCount *uint64 `json:"ReadLogCount,omitempty" name:"ReadLogCount"`
	// 处理数据量-输出行数

	WriteLogCount *uint64 `json:"WriteLogCount,omitempty" name:"WriteLogCount"`
	// 调度结果，1:运行中&nbsp;2:成功&nbsp;3:失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 失败原因字段

	StatusFailedMsg *string `json:"StatusFailedMsg,omitempty" name:"StatusFailedMsg"`
	// 任务查询时间

	TaskQueryTime *string `json:"TaskQueryTime,omitempty" name:"TaskQueryTime"`
}

type DeleteRemoteWriteTaskRequest struct {
	*tchttp.BaseRequest

	// RemoteWrite导入任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteRemoteWriteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteWriteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集日志配置

		LogConfigs []*LogConfigInfo `json:"LogConfigs,omitempty" name:"LogConfigs"`
		// 服务日志的配置信息

		ServiceLogConfigs []*ServiceLogConfigInfo `json:"ServiceLogConfigs,omitempty" name:"ServiceLogConfigs"`
		// 弃用

		LastVersion *string `json:"LastVersion,omitempty" name:"LastVersion"`
		// 弃用

		NeedUpdate *bool `json:"NeedUpdate,omitempty" name:"NeedUpdate"`
		// 弃用

		URL *string `json:"URL,omitempty" name:"URL"`
		// 弃用

		FileMd5 *string `json:"FileMd5,omitempty" name:"FileMd5"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAgentConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器状态信息组

		Machines []*MachineInfo `json:"Machines,omitempty" name:"Machines"`
		// 机器组是否开启自动升级功能。&nbsp;0：未开启自动升级；1：开启了自动升级。

		AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
		// 机器组自动升级功能预设开始时间

		UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
		// 机器组自动升级功能预设结束时间

		UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
		// 当前用户可用最新的Loglistener版本

		LatestAgentVersion *string `json:"LatestAgentVersion,omitempty" name:"LatestAgentVersion"`
		// 是否开启服务日志。true表示开启服务日志，false表示不开启服务日志。

		ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
		// TKE标志位，默认值为空字符串。空字符串表示日志不是来自于TKE，label_k8s表示日志来自于TKE。

		Flag *string `json:"Flag,omitempty" name:"Flag"`
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendConsumerHeartbeatResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主题对应的消费组标识

		ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
		// 分区信息

		TopicPartitionsInfo []*TopicPartitionInfo `json:"TopicPartitionsInfo,omitempty" name:"TopicPartitionsInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SendConsumerHeartbeatResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendConsumerHeartbeatResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicIndexInfo struct {

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 索引是否生效。true表示索引生效，false表示索引失效。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 索引配置信息&nbsp;注意：此字段可能返回&nbsp;null，表示取不到有效值。

	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
	// 索引修改时间，初始值为索引创建时间。

	ModifyTime *string `json:"ModifyTime,omitempty" name:"ModifyTime"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 内置保留字段（__FILENAME__，__HOSTNAME__及__SOURCE__）是否包含至全文索引&nbsp;&nbsp;false:不包含&nbsp;true:包含

	IncludeInternalFields *bool `json:"IncludeInternalFields,omitempty" name:"IncludeInternalFields"`
	// 元数据字段（前缀为__TAG__的字段）是否包含至全文索引&nbsp;&nbsp;0:仅包含开启键值索引的元数据字段&nbsp;1:包含所有元数据字段&nbsp;2:不包含任何元数据字段

	MetadataFlag *uint64 `json:"MetadataFlag,omitempty" name:"MetadataFlag"`
	// 自定义日志解析异常存储字段。

	CoverageField *string `json:"CoverageField,omitempty" name:"CoverageField"`
}

type CreateKafkaRechargeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Kafka导入配置ID

		Id *string `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateKafkaRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigurationTemplateRequest struct {
	*tchttp.BaseRequest

	// 配置模板id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 模版名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z,&nbsp;_,-,zh字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注，可选参数，长度不超过255个字符

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 投递cos模板配置项

	ShipperTemplateInfos []*ShipperTemplateInfo `json:"ShipperTemplateInfos,omitempty" name:"ShipperTemplateInfos"`
}

func (r *ModifyConfigurationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigurationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsumerGroupRequest struct {
	*tchttp.BaseRequest

	// 创建的消费者组标识&nbsp;限制：&nbsp;字母数字下划线，不允许数字开头，长度限制256

	ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
	// 消费者心跳超时时间（秒）

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 创建的消费者组包含的日志主题列表

	Topics []*string `json:"Topics,omitempty" name:"Topics"`
	// 日志集Id（日志主题所属的日志集）

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *CreateConsumerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsumerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConditionInfo struct {

	// 条件属性，目前只支持VpcID

	Attributes *string `json:"Attributes,omitempty" name:"Attributes"`
	// 条件规则，1:等于，2:不等于

	Rule *uint64 `json:"Rule,omitempty" name:"Rule"`
	// 对应条件属性的值

	ConditionValue *string `json:"ConditionValue,omitempty" name:"ConditionValue"`
}

type NoticeContentInfo struct {

	// 通知内容模板标题信息。&nbsp;部分通知渠道类型不支持“标题”，请参照控制台页面。

	Title *string `json:"Title,omitempty" name:"Title"`
	// 通知内容模板正文信息。

	Content *string `json:"Content,omitempty" name:"Content"`
	// 请求头（Request&nbsp;Headers）：在HTTP请求中，请求头包含了客户端向服务器发送的附加信息，如用户代理、授权凭证、期望的响应格式等。&nbsp;仅“自定义回调”支持该配置。

	Headers []*string `json:"Headers,omitempty" name:"Headers"`
}

type DescribeMetricSubscribesRequest struct {
	*tchttp.BaseRequest

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// taskId按照【配置id】进行过滤。类型：String&nbsp;必选：否&nbsp;&nbsp;name按照【配置名称】进行过滤。类型：String&nbsp;必选：否&nbsp;&nbsp;status按照【配置状态标记】进行过滤。类型：String&nbsp;必选：否&nbsp;&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMetricSubscribesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricSubscribesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CollectInfo struct {

	// 指定采集类型的采集配置信息。&nbsp;&nbsp;当Type为0时，CollectConfigs不允许为空。&nbsp;当Type为1时，CollectConfigs为空时，表示选择所有Pod&nbsp;Label；否则CollectConfigs为指定Pod&nbsp;Label。

	CollectConfigs []*CollectConfig `json:"CollectConfigs,omitempty" name:"CollectConfigs"`
	// 采集类型，必填字段。&nbsp;&nbsp;0：元数据配置。&nbsp;1：指定Pod&nbsp;Label。

	Type *uint64 `json:"Type,omitempty" name:"Type"`
}

type LogItem struct {

	// 日志Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 日志Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeResourcesRequest struct {
	*tchttp.BaseRequest

	// 获取数据地域类型，all:全地域，ap-region指定地域

	DataRegion *string `json:"DataRegion,omitempty" name:"DataRegion"`
}

func (r *DescribeResourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicExtendConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicExtendConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicExtendConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFoldersRequest struct {
	*tchttp.BaseRequest

	// folderId：按照【文件夹Id】进行过滤。类型：String必选：否&nbsp;folderName按照【按照文件夹名称】进行过滤（支持模糊匹配）。类型：String必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeFoldersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFoldersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmShieldInfo struct {

	// 通知渠道组Id

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
	// 屏蔽规则id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 屏蔽开始时间（秒级时间戳）。

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 屏蔽结束时间（秒级时间戳）。

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 屏蔽类型。1：屏蔽所有通知，2：按照Rule参数屏蔽匹配规则的通知。

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 屏蔽规则，当Type为2时必填。规则填写方式详见产品文档。

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 屏蔽原因。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 规则创建来源。&nbsp;1.&nbsp;控制台，2.api，3.告警通知

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 操作者。

	Operator *string `json:"Operator,omitempty" name:"Operator"`
	// 规则状态。&nbsp;0：暂未生效，1：生效中，2：已失效

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 规则创建时间。

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 规则更新时间。

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteMetricConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMetricConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetricConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePartitionsRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribePartitionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePartitionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RemoveMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AppointLabel struct {

	// 指定标签类型。&nbsp;-&nbsp;0：所有Pod&nbsp;label，Keys字段无效&nbsp;-&nbsp;1：指定Pod&nbsp;label，Keys字段不能为空

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 元数据Pod标签的键。有效标签键有两个部分：可选前缀和名称，以斜杠&nbsp;(/)&nbsp;分隔。名称部分是必需的，并且必须不超过&nbsp;63&nbsp;个字符，以字母数字字符&nbsp;([a-z0-9A-Z])&nbsp;开头和结尾，中间有破折号(-)、下划线(_)、点(.)&nbsp;和字母数字。前缀是可选的。如果指定，前缀必须是&nbsp;DNS&nbsp;子域：一系列以点&nbsp;(.)&nbsp;分隔的&nbsp;DNS&nbsp;标签，总长度不超过&nbsp;253&nbsp;个字符，后跟斜杠&nbsp;(&nbsp;/)。&nbsp;-&nbsp;prefix&nbsp;格式&nbsp;`[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*`&nbsp;-&nbsp;name&nbsp;格式&nbsp;`([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]`&nbsp;-&nbsp;key不能重复

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
}

type MetricSpec struct {

	// 自定义指标采集配置项

	CustomSpecs []*CustomMetricSpec `json:"CustomSpecs,omitempty" name:"CustomSpecs"`
}

type DeleteMachineGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EstimateRebuildIndexTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预估索引重建需要时间，单位秒

		RemainTime *uint64 `json:"RemainTime,omitempty" name:"RemainTime"`
		// 预估写流量大小，单位MB

		WriteTraffic *uint64 `json:"WriteTraffic,omitempty" name:"WriteTraffic"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EstimateRebuildIndexTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EstimateRebuildIndexTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRebuildIndexTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 索引重建任务列表

		RebuildTasks []*RebuildIndexTaskInfo `json:"RebuildTasks,omitempty" name:"RebuildTasks"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRebuildIndexTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRebuildIndexTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordingRuleYamlTaskRequest struct {
	*tchttp.BaseRequest

	// Yaml配置id

	YamlID *string `json:"YamlID,omitempty" name:"YamlID"`
	// 任务数据源日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 任务写入目标主题

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
	// 任务状态；&nbsp;1:开启；2:关闭

	EnableFlag *uint64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 调度开始时间,Unix时间戳，单位ms

	ProcessStartTime *uint64 `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 调度周期(分钟)，支持范围(0,1440]分钟。

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 执行周期单位,&nbsp;`0`:&nbsp;Minute,&nbsp;`1`:&nbsp;Second。&nbsp;默认：0&nbsp;-&nbsp;ProcessPeriodUnit:1时，ProcessPeriod&nbsp;只支持&nbsp;30（即只支持30秒）。&nbsp;-&nbsp;ProcessPeriodUnit:0时，ProcessPeriod&nbsp;支持范围(0,1440]分钟。

	ProcessPeriodUnit *uint64 `json:"ProcessPeriodUnit,omitempty" name:"ProcessPeriodUnit"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// yaml配置名称

	YamlConfigName *string `json:"YamlConfigName,omitempty" name:"YamlConfigName"`
	// yaml配置内容

	YamlContent *string `json:"YamlContent,omitempty" name:"YamlContent"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *ModifyRecordingRuleYamlTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecordingRuleYamlTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterMetricConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 指标采集配置列表

		Datas []*MetricCollectConfig `json:"Datas,omitempty" name:"Datas"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterMetricConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterMetricConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicExtendConfigRequest struct {
	*tchttp.BaseRequest

	// cls的业务标识字段

	LbKeys []*string `json:"LbKeys,omitempty" name:"LbKeys"`
}

func (r *DescribeTopicExtendConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicExtendConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Dimension struct {

	// 实例维度名称,此字段可能返回&nbsp;null，表示取不到有效值。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 实例维度值,此字段可能返回&nbsp;null，表示取不到有效值。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteCosRechargeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteCosRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupsRequest struct {
	*tchttp.BaseRequest

	// 日志集Id（日志主题所属的日志集）

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// topic列表

	Topics []*string `json:"Topics,omitempty" name:"Topics"`
}

func (r *DescribeConsumerGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EscalateNoticeInfo struct {

	// 告警通知模板接收者信息。

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 告警通知模板回调信息。

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 告警升级开关。true：开启告警升级、false：关闭告警升级，默认：false

	Escalate *bool `json:"Escalate,omitempty" name:"Escalate"`
	// 告警升级间隔。单位：分钟，范围[1，14400]

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 升级条件。1：无人认领且未恢复、2：未恢复，默认为1&nbsp;&nbsp;无人认领且未恢复：告警没有恢复并且没有人认领则升级&nbsp;未恢复：当前告警持续未恢复则升级

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 告警升级后下一个环节的通知渠道配置，最多可配置5个环节。

	EscalateNotice *EscalateNoticeInfo `json:"EscalateNotice,omitempty" name:"EscalateNotice"`
}

type KafkaRechargeInfo struct {

	// Kafka数据订阅配置的ID。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// Kafka导入任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 导入Kafka类型，0:&nbsp;CKafka，1:&nbsp;用户自建Kafka

	KafkaType *uint64 `json:"KafkaType,omitempty" name:"KafkaType"`
	// CKafka实例ID，KafkaType为0时必填

	KafkaInstance *string `json:"KafkaInstance,omitempty" name:"KafkaInstance"`
	// 服务地址

	ServerAddr *string `json:"ServerAddr,omitempty" name:"ServerAddr"`
	// ServerAddr是否为加密连接。true表示是加密连接，false表示不是加密连接。

	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitempty" name:"IsEncryptionAddr"`
	// 加密访问协议，IsEncryptionAddr参数为true时必填

	Protocol *KafkaProtocolInfo `json:"Protocol,omitempty" name:"Protocol"`
	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开

	UserKafkaTopics *string `json:"UserKafkaTopics,omitempty" name:"UserKafkaTopics"`
	// 用户Kafka消费组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 状态&nbsp;status&nbsp;1:&nbsp;运行中,&nbsp;2:&nbsp;暂停

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 导入数据位置，-2:最早（默认），-1：最晚

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 日志导入规则

	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitempty" name:"LogRechargeRule"`
	// 私有网络信息

	NetworkInfo *NetworkInfo `json:"NetworkInfo,omitempty" name:"NetworkInfo"`
}

type CreateMetricConfigRequest struct {
	*tchttp.BaseRequest

	// 日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 采集配置来源。支持&nbsp;：`0`、`1`&nbsp;-&nbsp;0:自建k8s&nbsp;-&nbsp;1:TKE

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 机器组id。

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// 监控类型。支持&nbsp;：`0`、`1`，不支持修改&nbsp;-&nbsp;0:基础监控&nbsp;-&nbsp;1:自定义监控,

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 采集配置方式。支持&nbsp;：`0`、`1`，不支持修改&nbsp;-&nbsp;0:普通配置方式，Type字段只能为：``1`&nbsp;-&nbsp;1:YAML导入方式，&nbsp;Type&nbsp;可以是：`0`或者`1`

	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
	// 名称：长度不超过253字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 采集对象,&nbsp;Flag=0时生效

	Spec *MetricSpec `json:"Spec,omitempty" name:"Spec"`
	// 标签处理,&nbsp;Flag=0时生效

	Relabels []*Relabeling `json:"Relabels,omitempty" name:"Relabels"`
	// 标签处理,&nbsp;Flag=0时生效

	MetricRelabels []*Relabeling `json:"MetricRelabels,omitempty" name:"MetricRelabels"`
	// 自定义元数据,&nbsp;Flag=0时生效

	MetricLabel *MetricConfigLabel `json:"MetricLabel,omitempty" name:"MetricLabel"`
	// 通信协议&nbsp;http、https;&nbsp;Flag=0时生效

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 采集频率,&nbsp;Flag=0时生效

	ScrapeInterval *string `json:"ScrapeInterval,omitempty" name:"ScrapeInterval"`
	// 采集超时时间，&nbsp;Flag=0时生效

	ScrapeTimeout *string `json:"ScrapeTimeout,omitempty" name:"ScrapeTimeout"`
	// Prometheus如何处理标签之间的冲突。当Flag=0时生效，支持`true`,`false`&nbsp;-&nbsp;`false`:配置数据中冲突的标签重命名&nbsp;-&nbsp;`true`:忽略冲突的服务器端标签

	HonorLabels *bool `json:"HonorLabels,omitempty" name:"HonorLabels"`
	// 压缩,&nbsp;false:关闭,&nbsp;true:开启,&nbsp;Flag=0时生效

	EnableCompression *bool `json:"EnableCompression,omitempty" name:"EnableCompression"`
	// 采集配置yaml格式字符串,&nbsp;Flag=1时必填

	YamlSpec *MetricYamlSpec `json:"YamlSpec,omitempty" name:"YamlSpec"`
	// 日志主题扩展信息

	TopicExtendInfo *string `json:"TopicExtendInfo,omitempty" name:"TopicExtendInfo"`
}

func (r *CreateMetricConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMetricConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaUserResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 如果返回不为空，代表用户名UserName已经创建成功。

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaUserResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyBinlogSubscribeRequest struct {
	*tchttp.BaseRequest

	// binlog采集任务的日志主题id。必填字段

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// binlog采集任务id。必填字段

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z,&nbsp;_,-,zh字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 数据库类型。&nbsp;1:云mysql&nbsp;2:TDSQL-C&nbsp;3:自建mysql

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 数据库配置信息。

	SqlInfo *BinlogDBConfig `json:"SqlInfo,omitempty" name:"SqlInfo"`
	// binlog采集起始点配置信息。

	InitialPoint *BinlogInitialPoint `json:"InitialPoint,omitempty" name:"InitialPoint"`
	// binlog采集事件配置。

	Events []*string `json:"Events,omitempty" name:"Events"`
	// binlog采集元数据配置。

	Metadatas []*string `json:"Metadatas,omitempty" name:"Metadatas"`
	// 时间戳来源。&nbsp;1:采集时间&nbsp;2:事件时间

	TimestampType *uint64 `json:"TimestampType,omitempty" name:"TimestampType"`
	// 黑名单状态。&nbsp;1:关闭&nbsp;2:开启

	BlockListStatus *uint64 `json:"BlockListStatus,omitempty" name:"BlockListStatus"`
	// 黑名单配置信息，当BlockListStatus&nbsp;=2时，BlockList必填。

	BlockList []*BinlogFilterList `json:"BlockList,omitempty" name:"BlockList"`
	// 白名单状态。&nbsp;1:关闭&nbsp;2:开启

	AllowListStatus *uint64 `json:"AllowListStatus,omitempty" name:"AllowListStatus"`
	// 白名单配置信息，当AllowListStatus=2时，AllowList必填。

	AllowList []*BinlogFilterList `json:"AllowList,omitempty" name:"AllowList"`
	// 任务状态。&nbsp;1：&nbsp;未启用&nbsp;2：&nbsp;启用

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 平铺事件数据标记。&nbsp;默认Flatten为false。&nbsp;Flatten为false时：会将事件数据以数组+JSON格式集中打包到old_data和data两个字段中。如:&nbsp;在数据库操作中执行了两次操作，一次更新了a和b字段，一次d和e，&nbsp;日志值为old_data:[{a:a1,b:b1},&nbsp;{d:d1,e:e1}],&nbsp;data:[{a:a2,b:b2},&nbsp;{d:d2,e:e2}]。&nbsp;Flatten为true时：会将不同的事件拆分为多条日志，并在每条日志中对被操作的字段进行平铺。如：在数据库操作中执行了两次操作，一次更新了a和b字段，一次d和e，&nbsp;则会生成两条日志，&nbsp;分别的值为&nbsp;a:a2,&nbsp;old_a:a1,&nbsp;b:b2,&nbsp;old_b:b1和d:d2,&nbsp;old_d:d1,&nbsp;e:e2,&nbsp;old_e:e1。

	Flatten *bool `json:"Flatten,omitempty" name:"Flatten"`
}

func (r *ModifyBinlogSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyBinlogSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmRuleTestResult struct {

	// 位序

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 错误码。0表示请求正常，-1001表示异常请求

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type DescribeCosRechargesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// cos导入配置信息

		Data []*CosRechargeInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCosRechargesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosRechargesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConsumerGroupInfoForMonitor struct {

	// 消费组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 消费组监控信息

	PartitionListForMonitor []*PartitionInfoForMonitor `json:"PartitionListForMonitor,omitempty" name:"PartitionListForMonitor"`
}

type DescribeMetricSubscribePreviewRequest struct {
	*tchttp.BaseRequest

	// 云产品命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 数据库配置信息。

	Metrics []*MetricConfig `json:"Metrics,omitempty" name:"Metrics"`
	// 实例配置配置。

	InstanceInfo *InstanceConfig `json:"InstanceInfo,omitempty" name:"InstanceInfo"`
}

func (r *DescribeMetricSubscribePreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricSubscribePreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicMetricConfigsRequest struct {
	*tchttp.BaseRequest

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// configId按照【指标采集配置id】进行过滤。类型：String&nbsp;必选：否&nbsp;name按照【配置名称】进行过滤。类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，所有Filter.Values总和上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicMetricConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicMetricConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MergePartitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 合并结果集

		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *MergePartitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MergePartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogsetInfo struct {

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 若AssumerUin非空，则表示创建该日志集的服务方Uin

	AssumerUin *uint64 `json:"AssumerUin,omitempty" name:"AssumerUin"`
	// 若AssumerUin非空，则表示创建该日志集的服务方名称

	AssumerName *string `json:"AssumerName,omitempty" name:"AssumerName"`
	// 若AssumerUin非空，则表示非改服务方的调用者对于日志集的修改权限

	LogsetModifyAcl *int64 `json:"LogsetModifyAcl,omitempty" name:"LogsetModifyAcl"`
	// 日志集绑定的标签

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 日志集下日志主题的数目

	TopicCount *int64 `json:"TopicCount,omitempty" name:"TopicCount"`
	// 若AssumerUin非空，则表示创建该日志集的服务方角色

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 生命周期，单位为天

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 若AssumerUin非空，则表示非改服务方的调用者对于日志集的删除权限

	LogsetDelACL *int64 `json:"LogsetDelACL,omitempty" name:"LogsetDelACL"`
	// 若AssumerUin非空，则表示非改服务方的调用者对于日志集的查询权限

	LogsetShowAcl *uint64 `json:"LogsetShowAcl,omitempty" name:"LogsetShowAcl"`
	// 日志集下指标主题的数目

	MetricTopicCount *int64 `json:"MetricTopicCount,omitempty" name:"MetricTopicCount"`
}

type AddMachineGroupInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddMachineGroupInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMachineGroupInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type DeleteNoticeContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNoticeContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNoticeContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogSubscribePreviewRequest struct {
	*tchttp.BaseRequest

	// 数据库类型。&nbsp;1:云mysql&nbsp;2:TDSQL-C&nbsp;3:自建mysql

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 数据库配置信息。

	SqlInfo *BinlogDBConfig `json:"SqlInfo,omitempty" name:"SqlInfo"`
	// binlog采集起始点配置信息。

	InitialPoint *BinlogInitialPoint `json:"InitialPoint,omitempty" name:"InitialPoint"`
	// binlog采集事件配置。

	Events []*string `json:"Events,omitempty" name:"Events"`
	// binlog采集元数据配置。

	Metadatas []*string `json:"Metadatas,omitempty" name:"Metadatas"`
	// 时间戳来源。&nbsp;1:采集时间&nbsp;2:事件时间

	TimestampType *uint64 `json:"TimestampType,omitempty" name:"TimestampType"`
	// 黑名单状态。&nbsp;1:关闭&nbsp;2:开启

	BlockListStatus *uint64 `json:"BlockListStatus,omitempty" name:"BlockListStatus"`
	// 黑名单配置信息，当BlockListStatus&nbsp;=2时，BlockList必填。

	BlockList []*BinlogFilterList `json:"BlockList,omitempty" name:"BlockList"`
	// 白名单状态。&nbsp;1:关闭&nbsp;2:开启

	AllowListStatus *uint64 `json:"AllowListStatus,omitempty" name:"AllowListStatus"`
	// 白名单配置信息，当AllowListStatus=2时，AllowList必填。

	AllowList []*BinlogFilterList `json:"AllowList,omitempty" name:"AllowList"`
	// 任务id。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 平铺事件数据标记。&nbsp;默认Flatten为false。&nbsp;Flatten为false时：会将事件数据以数组+JSON格式集中打包到old_data和data两个字段中。如:&nbsp;在数据库操作中执行了两次操作，一次更新了a和b字段，一次d和e，&nbsp;日志值为old_data:[{a:a1,b:b1},&nbsp;{d:d1,e:e1}],&nbsp;data:[{a:a2,b:b2},&nbsp;{d:d2,e:e2}]。&nbsp;Flatten为true时：会将不同的事件拆分为多条日志，并在每条日志中对被操作的字段进行平铺。如：在数据库操作中执行了两次操作，一次更新了a和b字段，一次d和e，&nbsp;则会生成两条日志，&nbsp;分别的值为&nbsp;a:a2,&nbsp;old_a:a1,&nbsp;b:b2,&nbsp;old_b:b1和d:d2,&nbsp;old_d:d1,&nbsp;e:e2,&nbsp;old_e:e1。

	Flatten *bool `json:"Flatten,omitempty" name:"Flatten"`
}

func (r *DescribeBinlogSubscribePreviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBinlogSubscribePreviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。

	Key *string `json:"Key,omitempty" name:"Key"`
	// 需要过滤的值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type KafkaConsumerContent struct {

	// 消费数据格式。&nbsp;0：原始内容；1：JSON。

	Format *int64 `json:"Format,omitempty" name:"Format"`
	// 是否投递&nbsp;TAG&nbsp;信息&nbsp;Format为0时，此字段不需要赋值

	EnableTag *bool `json:"EnableTag,omitempty" name:"EnableTag"`
	// tag数据处理方式：1:不平铺（默认值）；2:平铺。&nbsp;&nbsp;不平铺示例：&nbsp;TAG信息：{"__TAG__":{"fieldA":200,"fieldB":"text"}}&nbsp;不平铺：{"__TAG__":{"fieldA":200,"fieldB":"text"}}&nbsp;&nbsp;平铺示例：&nbsp;TAG信息：{"__TAG__":{"fieldA":200,"fieldB":"text"}}&nbsp;平铺：{"__TAG__.fieldA":200,"__TAG__.fieldB":"text"}

	TagTransaction *int64 `json:"TagTransaction,omitempty" name:"TagTransaction"`
	// 元数据信息列表,&nbsp;可选值为：__SOURCE__、__FILENAME__&nbsp;、__TIMESTAMP__、__HOSTNAME__、__PKGID__&nbsp;Format为0时，此字段不需要赋值

	MetaFields []*string `json:"MetaFields,omitempty" name:"MetaFields"`
	// 消费数据Json格式：&nbsp;1：不转义（默认格式）&nbsp;2：转义&nbsp;&nbsp;投递Json格式。&nbsp;JsonType为1：和原始日志一致，不转义。示例：&nbsp;日志原文：{"a":"aa",&nbsp;"b":{"b1":"b1b1",&nbsp;"c1":"c1c1"}}&nbsp;投递到Ckafka：{"a":"aa",&nbsp;"b":{"b1":"b1b1",&nbsp;"c1":"c1c1"}}&nbsp;&nbsp;JsonType为2：转义。示例：&nbsp;日志原文：{"a":"aa",&nbsp;"b":{"b1":"b1b1",&nbsp;"c1":"c1c1"}}&nbsp;投递到Ckafka：{"a":"aa","b":"{\"b1\":\"b1b1\",&nbsp;\"c1\":\"c1c1\"}"}

	JsonType *int64 `json:"JsonType,omitempty" name:"JsonType"`
}

type KafkaProtocolInfo struct {

	// 协议类型，支持的协议类型包括&nbsp;plaintext、sasl_plaintext&nbsp;或&nbsp;sasl_ssl。建议使用&nbsp;sasl_ssl，此协议会进行连接加密同时需要用户认证。&nbsp;入参必填

	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`
	// 加密类型，支持&nbsp;PLAIN、SCRAM-SHA-256&nbsp;或&nbsp;SCRAM-SHA-512。&nbsp;当Protocol为sasl_plaintext或sasl_ssl时必填

	Mechanism *string `json:"Mechanism,omitempty" name:"Mechanism"`
	// 用户名。&nbsp;当Protocol为sasl_plaintext或sasl_ssl时必填

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 用户密码。&nbsp;当Protocol为sasl_plaintext或sasl_ssl时必填

	Password *string `json:"Password,omitempty" name:"Password"`
	// 是否开启客户端证书验证

	EnableClientCertificate *uint64 `json:"EnableClientCertificate,omitempty" name:"EnableClientCertificate"`
	// 是否开启服务端证书验证

	EnableServerCertificate *uint64 `json:"EnableServerCertificate,omitempty" name:"EnableServerCertificate"`
	// 云托管CA证书id

	CACertificateId *string `json:"CACertificateId,omitempty" name:"CACertificateId"`
	// 云托管服务端证书id

	SVRCertificateId *string `json:"SVRCertificateId,omitempty" name:"SVRCertificateId"`
}

type RuleInfo struct {

	// 全文索引配置,&nbsp;为空时代表未开启全文索引

	FullText *FullTextInfo `json:"FullText,omitempty" name:"FullText"`
	// 键值索引配置，为空时代表未开启键值索引

	KeyValue *RuleKeyValueInfo `json:"KeyValue,omitempty" name:"KeyValue"`
	// 元字段索引配置，为空时代表未开启元字段索引

	Tag *RuleTagInfo `json:"Tag,omitempty" name:"Tag"`
	// 键值索引自动配置，为空时代表未开启该功能。&nbsp;启用后自动将日志内的字段添加到键值索引中，包括日志中后续新增的字段。

	DynamicIndex *DynamicIndex `json:"DynamicIndex,omitempty" name:"DynamicIndex"`
}

type CreateIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 是否生效，默认为true。true表示索引生效，false表示索引不生效。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 索引规则

	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
	// 内置保留字段（__FILENAME__，__HOSTNAME__及__SOURCE__）是否包含至全文索引，默认为false，推荐设置为true&nbsp;&nbsp;false:不包含&nbsp;true:包含

	IncludeInternalFields *bool `json:"IncludeInternalFields,omitempty" name:"IncludeInternalFields"`
	// 元数据字段（前缀为__TAG__的字段）是否包含至全文索引，默认为0，推荐设置为1&nbsp;&nbsp;0:仅包含开启键值索引的元数据字段&nbsp;1:包含所有元数据字段&nbsp;2:不包含任何元数据字段

	MetadataFlag *uint64 `json:"MetadataFlag,omitempty" name:"MetadataFlag"`
	// 自定义日志解析异常存储字段。

	CoverageField *string `json:"CoverageField,omitempty" name:"CoverageField"`
}

func (r *CreateIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterMetricConfigsRequest struct {
	*tchttp.BaseRequest

	// 机器组id

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// configId按照【指标采集配置id】进行过滤。类型：String&nbsp;必选：否&nbsp;name按照【配置名称】进行过滤。类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，所有Filter.Values总和上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 是否从缓存获取数据

	FromCache *bool `json:"FromCache,omitempty" name:"FromCache"`
}

func (r *DescribeClusterMetricConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterMetricConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AnonymousInfo struct {

	// 操作列表，支持trackLog(JS/HTTP上传日志&nbsp;)和realtimeProducer(kafka协议上传日志)

	Operations []*string `json:"Operations,omitempty" name:"Operations"`
	// 条件列表

	Conditions []*ConditionInfo `json:"Conditions,omitempty" name:"Conditions"`
}

type RemoteWriteInfo struct {

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// Remote&nbsp;Write任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 网络类型&nbsp;1:&nbsp;内网&nbsp;2:外网

	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 任务运行状态&nbsp;1:&nbsp;运行中&nbsp;2:暂停&nbsp;3:&nbsp;失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 目标服务名称

	Target *string `json:"Target,omitempty" name:"Target"`
	// 目标地址

	RemoteWriteURL *string `json:"RemoteWriteURL,omitempty" name:"RemoteWriteURL"`
	// 鉴权类型&nbsp;0:&nbsp;无鉴权&nbsp;1:&nbsp;basic_auth&nbsp;2:&nbsp;token

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 鉴权信息

	AuthInfo *RemoteWriteAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
	// 日志集

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 任务状态

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 子网

	SubNet *string `json:"SubNet,omitempty" name:"SubNet"`
	// 后端服务类型

	VirtualGatewayType *int64 `json:"VirtualGatewayType,omitempty" name:"VirtualGatewayType"`
	// 云时序数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

type AdvanceFilterRuleInfo struct {

	// 过滤字段

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过滤规则，0:等于，1:字段存在，2:字段不存在

	Rule *uint64 `json:"Rule,omitempty" name:"Rule"`
	// 过滤值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ServiceLogConfigInfo struct {

	// 服务日志的logset信息

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 服务日志的Topic&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 服务日志的Topic&nbsp;Name

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type CreateRecordingRuleTaskRequest struct {
	*tchttp.BaseRequest

	// 任务数据源日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 任务写入目标主题

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
	// 预聚合任务；限制：字母数字下划线，不允许下划线开头，小于256个字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务状态；&nbsp;1:开启；2:关闭

	EnableFlag *uint64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 调度开始时间,Unix时间戳，单位ms

	ProcessStartTime *uint64 `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 调度周期(分钟)，支持范围(0,1440]分钟。

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// 查询语句

	RecordingRuleContent *string `json:"RecordingRuleContent,omitempty" name:"RecordingRuleContent"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 执行周期单位,&nbsp;`0`:&nbsp;Minute,&nbsp;`1`:&nbsp;Second。&nbsp;默认：0&nbsp;-&nbsp;ProcessPeriodUnit:1时，ProcessPeriod&nbsp;只支持&nbsp;30（即只支持30秒）。&nbsp;-&nbsp;ProcessPeriodUnit:0时，ProcessPeriod&nbsp;支持范围(0,1440]分钟。

	ProcessPeriodUnit *uint64 `json:"ProcessPeriodUnit,omitempty" name:"ProcessPeriodUnit"`
	// 指标自定义Label

	CustomMetricLabels []*MetricLabel `json:"CustomMetricLabels,omitempty" name:"CustomMetricLabels"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *CreateRecordingRuleTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRecordingRuleTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardSubscribesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘订阅列表

		DashboardSubscribeInfos []*DashboardSubscribeInfo `json:"DashboardSubscribeInfos,omitempty" name:"DashboardSubscribeInfos"`
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDashboardSubscribesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardSubscribesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FunctionInfo struct {

	// 函数名称

	FuncName *string `json:"FuncName,omitempty" name:"FuncName"`
	// 函数描述

	FuncNameDesc *string `json:"FuncNameDesc,omitempty" name:"FuncNameDesc"`
	// 用来说明函数功能

	FuncUseDesc *string `json:"FuncUseDesc,omitempty" name:"FuncUseDesc"`
	// 语法描述

	FuncSyntaxDesc *string `json:"FuncSyntaxDesc,omitempty" name:"FuncSyntaxDesc"`
	// demo展示

	FuncDemo *string `json:"FuncDemo,omitempty" name:"FuncDemo"`
	// 函数类型

	FuncType *string `json:"FuncType,omitempty" name:"FuncType"`
	// true是可变长度参数的加工函数

	IsVariadic *bool `json:"IsVariadic,omitempty" name:"IsVariadic"`
	// 可变参的参数个数最大限制

	MaxArgumentSize *int64 `json:"MaxArgumentSize,omitempty" name:"MaxArgumentSize"`
	// 函数返回结果类型，用来校验嵌套函数中，返回结果是否和函数参数类型匹配。&nbsp;&nbsp;不同的func_type返回的对象类型不同&nbsp;string/int/bool/condition/func

	ReturnType *string `json:"ReturnType,omitempty" name:"ReturnType"`
	// 函数参数描述

	Arguments []*FunctionArgument `json:"Arguments,omitempty" name:"Arguments"`
}

type MetricConfig struct {

	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 统计周期,单位:秒（s）

	Periods []*uint64 `json:"Periods,omitempty" name:"Periods"`
	// 自定义指标标签

	MetricLabels []*MetricLabel `json:"MetricLabels,omitempty" name:"MetricLabels"`
}

type DescribeScheduledSqlInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ScheduledSQL任务列表信息

		ScheduledSqlTaskInfos []*ScheduledSqlTaskInfo `json:"ScheduledSqlTaskInfos,omitempty" name:"ScheduledSqlTaskInfos"`
		// 任务总次数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScheduledSqlInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScheduledSqlInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigFromMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigFromMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigFromMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmTarget struct {

	// 日志主题ID。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 查询语句。

	Query *string `json:"Query,omitempty" name:"Query"`
	// 告警对象序号；从1开始递增。

	Number *int64 `json:"Number,omitempty" name:"Number"`
	// 查询范围起始时间相对于告警执行时间的偏移，单位为分钟，取值为非正，最大值为0，最小值为-1440。

	StartTimeOffset *int64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`
	// 查询范围终止时间相对于告警执行时间的偏移，单位为分钟，取值为非正，须大于StartTimeOffset，最大值为0，最小值为-1440。

	EndTimeOffset *int64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
	// 日志集ID。

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 检索语法规则，默认值为0。&nbsp;0：Lucene语法，1：CQL语法。&nbsp;详细说明参见检索条件语法规则

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 交互模板查询语句配置信息。

	QueryInteractiveConfig *string `json:"QueryInteractiveConfig,omitempty" name:"QueryInteractiveConfig"`
}

type ContainerStdoutInfo struct {

	// 是否所有容器。true表示是所有容器，false表示不包含所有容器。

	AllContainers *bool `json:"AllContainers,omitempty" name:"AllContainers"`
	// container为空表所有的，不为空采集指定的容器

	Container *string `json:"Container,omitempty" name:"Container"`
	// namespace可以多个，用分隔号分割,例如A,B；为空或者没有这个字段，表示所有namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// pod标签信息

	IncludeLabels []*string `json:"IncludeLabels,omitempty" name:"IncludeLabels"`
	// 工作负载信息

	WorkLoads []*ContainerWorkLoadInfo `json:"WorkLoads,omitempty" name:"WorkLoads"`
	// 需要排除的namespace可以多个，用分隔号分割,例如A,B

	ExcludeNamespace *string `json:"ExcludeNamespace,omitempty" name:"ExcludeNamespace"`
	// 需要排除的pod标签信息

	ExcludeLabels []*string `json:"ExcludeLabels,omitempty" name:"ExcludeLabels"`
	// 容器名称标记。必填字段，不填默认值为0。&nbsp;&nbsp;默认值0，0:选中Container标记。&nbsp;1:排除Container标记。

	ContainerFlag *uint64 `json:"ContainerFlag,omitempty" name:"ContainerFlag"`
	// metadata信息

	CustomLabels *string `json:"CustomLabels,omitempty" name:"CustomLabels"`
}

type DeleteConsumerGroupRequest struct {
	*tchttp.BaseRequest

	// 需要删除的消费者组标识

	ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *DeleteConsumerGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsumerGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesRequest struct {
	*tchttp.BaseRequest

	// name&nbsp;按照【通知渠道组名称】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;alarmNoticeId&nbsp;按照【通知渠道组ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;uid&nbsp;按照【接收用户ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;groupId&nbsp;按照【接收用户组ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;deliverFlag&nbsp;按照【投递状态】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;可选值：&nbsp;"1":未启用,&nbsp;"2":&nbsp;已启用,&nbsp;"3":投递异常&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 是否需要返回通知渠道组配置的告警屏蔽统计状态数量信息。&nbsp;true：需要返回；false：&nbsp;不返回（默认false）。

	HasAlarmShieldCount *bool `json:"HasAlarmShieldCount,omitempty" name:"HasAlarmShieldCount"`
}

func (r *DescribeAlarmNoticesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmNoticesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicMetricConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 指标采集配置列表

		Datas []*MetricCollectConfig `json:"Datas,omitempty" name:"Datas"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicMetricConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicMetricConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoveMachineRequest struct {
	*tchttp.BaseRequest

	// 机器组&nbsp;ID。

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 剔除的Ip数组。

	Ips []*string `json:"Ips,omitempty" name:"Ips"`
	// 剔除的机器实例ID列表

	InstanceIDs []*string `json:"InstanceIDs,omitempty" name:"InstanceIDs"`
}

func (r *RemoveMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RemoveMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QcloudDataInfo struct {

	// 后付费订单号，每个物品对应一个dealName

	DealNames []*string `json:"DealNames,omitempty" name:"DealNames"`
	// 异步发货产品查询发货状态标识

	FlowId *string `json:"FlowId,omitempty" name:"FlowId"`
	// 多物品发货对应的flowId，与dealNames一一对应

	FlowIds []*string `json:"FlowIds,omitempty" name:"FlowIds"`
	// 每个订单号对应的发货资源id列表:{"20200929112744":["ins-kjs4jvkj"],"20200929112745":["disk-8ijw00wy"]}

	ResourceIds *string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 冻结流水，一次开通一个冻结流水

	BillId *string `json:"BillId,omitempty" name:"BillId"`
}

type SendDetail struct {

	// 发送次数

	Sms []*SendDetailItem `json:"Sms,omitempty" name:"Sms"`
	// 发送次数

	Email []*SendDetailItem `json:"Email,omitempty" name:"Email"`
	// 发送次数

	WeChat []*SendDetailItem `json:"WeChat,omitempty" name:"WeChat"`
	// 发送次数

	Phone []*SendDetailItem `json:"Phone,omitempty" name:"Phone"`
	// 发送次数

	Callback []*SendDetailItem `json:"Callback,omitempty" name:"Callback"`
}

type ModifyConsumerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConsumerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryRangeMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标查询结果类型

		ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
		// 指标查询结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryRangeMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryRangeMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardSubscribeRequest struct {
	*tchttp.BaseRequest

	// 仪表盘订阅记录id。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteDashboardSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataTransformFailureInfo struct {

	// 源日志

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
	// 加工失败原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

type WebCallback struct {

	// 回调地址，最大支持1024个字节。&nbsp;也可使用WebCallbackId引用集成配置中的URL，此时该字段请填写为空字符串。

	Url *string `json:"Url,omitempty" name:"Url"`
	// 回调方法。可选值：&nbsp;&nbsp;POST（默认值）&nbsp;PUT&nbsp;注意：&nbsp;&nbsp;参数CallbackType为Http时为必选，其它回调方式无需填写。

	Method *string `json:"Method,omitempty" name:"Method"`
	// 该参数已废弃，请使用NoticeContentId。

	Headers []*string `json:"Headers,omitempty" name:"Headers"`
	// 该参数已废弃，请使用NoticeContentId。

	Body *string `json:"Body,omitempty" name:"Body"`
	// 回调的类型。可选值：Http&nbsp;WeCom&nbsp;DingTalk&nbsp;Lark

	CallbackType *string `json:"CallbackType,omitempty" name:"CallbackType"`
	// 序号。&nbsp;&nbsp;入参无效。&nbsp;出参有效。

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 允许接收信息的开始时间。当做入参时，默认值为00:00:00

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 允许接收信息的结束时间。当做入参时，默认值为23:59:59

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 集成配置ID。

	WebCallbackId *string `json:"WebCallbackId,omitempty" name:"WebCallbackId"`
	// 通知内容模板ID，使用Default-zh引用默认模板（zh），使用Default-en引用DefaultTemplate(English)。

	NoticeContentId *string `json:"NoticeContentId,omitempty" name:"NoticeContentId"`
	// 提醒类型。&nbsp;&nbsp;0：不提醒；1：指定人；2：所有人

	RemindType *uint64 `json:"RemindType,omitempty" name:"RemindType"`
	// 电话列表。

	Mobiles []*string `json:"Mobiles,omitempty" name:"Mobiles"`
	// 用户ID列表。

	UserIds []*string `json:"UserIds,omitempty" name:"UserIds"`
}

type ApplyConfigurationTemplateRequest struct {
	*tchttp.BaseRequest

	// 配置模板id

	TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
	// 引用类型：1:&nbsp;指定日志主题，2：动态匹配

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 日志主题id列表，当Type为1时，此字段不能为空。

	TopicIds []*string `json:"TopicIds,omitempty" name:"TopicIds"`
	// 当Type为2时，此字段不能为空。

	LogsetInfos []*ConfigurationTemplateLogsetInfo `json:"LogsetInfos,omitempty" name:"LogsetInfos"`
	// 应用目标，1:&nbsp;投递COS，...

	ApplyTarget []*uint64 `json:"ApplyTarget,omitempty" name:"ApplyTarget"`
	// COS存储桶。当ApplyTarget有1时，此字段必填。

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
}

func (r *ApplyConfigurationTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyConfigurationTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集名字，不能重名

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 标签描述列表。最大支持10个标签键值对，并且不能有重复的键值对

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 生命周期，单位天；可取值范围1～366

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 日志集ID，格式为：用户自定义部分-用户appid，用户自定义部分仅支持小写字母、数字和-，且不能以-开头和结尾，长度为3至40字符，尾部需要使用-拼接用户appid

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *CreateLogsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExceptionResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异常资源列表

		ExceptionResources []*ExceptionResources `json:"ExceptionResources,omitempty" name:"ExceptionResources"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExceptionResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExceptionResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteDashboardSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteDashboardSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDashboardSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题的索引信息列表

		TopicIndexInfos []*TopicIndexInfo `json:"TopicIndexInfos,omitempty" name:"TopicIndexInfos"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIndexsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMetricSubscribeRequest struct {
	*tchttp.BaseRequest

	// 指标采集任务的日志主题id。必填字段

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 指标采集任务id。必填字段

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z,&nbsp;_,-,zh字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云产品命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标配置信息。

	Metrics []*MetricConfig `json:"Metrics,omitempty" name:"Metrics"`
	// 实例配置信息。

	InstanceInfo *InstanceConfig `json:"InstanceInfo,omitempty" name:"InstanceInfo"`
	// 任务状态。&nbsp;1：&nbsp;未启用&nbsp;2：&nbsp;启用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *ModifyMetricSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMetricSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// KafkaConsumer&nbsp;消费时使用的Topic参数

		TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenKafkaConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExportInfo struct {

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志导出任务ID

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
	// 日志导出查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 日志导出文件名

	FileName *string `json:"FileName,omitempty" name:"FileName"`
	// 日志文件大小

	FileSize *uint64 `json:"FileSize,omitempty" name:"FileSize"`
	// 日志导出时间排序。asc表示正序排序，desc表示倒序排序。

	Order *string `json:"Order,omitempty" name:"Order"`
	// 日志导出格式。json表示日志导出格式为json。

	Format *string `json:"Format,omitempty" name:"Format"`
	// 日志导出数量

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 日志下载状态。Processing:导出正在进行中，Completed:导出完成，Failed:导出失败，Expired:日志导出已过期(三天有效期),&nbsp;Queuing&nbsp;排队中

	Status *string `json:"Status,omitempty" name:"Status"`
	// 日志导出起始时间

	From *int64 `json:"From,omitempty" name:"From"`
	// 日志导出结束时间

	To *int64 `json:"To,omitempty" name:"To"`
	// 日志导出路径

	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`
	// 日志导出创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 语法规则。&nbsp;默认值为0。&nbsp;0：Lucene语法，1：CQL语法。

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 导出字段

	DerivedFields []*string `json:"DerivedFields,omitempty" name:"DerivedFields"`
	// 分隔符。CSV&nbsp;文件中各字段间的分隔符，非必填，默认&nbsp;逗号&nbsp;支持：&nbsp;-&nbsp;空格&nbsp;-&nbsp;`\t`：制表符&nbsp;-&nbsp;`&nbsp;,`：逗号&nbsp;-&nbsp;`;`：竖线

	Separator *string `json:"Separator,omitempty" name:"Separator"`
	// 转义符。CSV&nbsp;文件字段值中出现了分隔符的字符，需用转义符包裹该字符，防止读取数据时被错误识别，非必填，默认&nbsp;双引号&nbsp;支持：&nbsp;-&nbsp;`'&nbsp;`：单引号&nbsp;-&nbsp;`"`：双引号&nbsp;-&nbsp;空串

	EscapeCharacter *string `json:"EscapeCharacter,omitempty" name:"EscapeCharacter"`
	// 填充字段。csv文件字段不存在(无效)时，使用用户指定的值进行填充。&nbsp;默认空&nbsp;支持：&nbsp;&nbsp;空字符串&nbsp;满足正则表达式

	FillField *string `json:"FillField,omitempty" name:"FillField"`
	// 首行Key是否展示。默认&nbsp;true&nbsp;支持：&nbsp;-&nbsp;true：展示&nbsp;-&nbsp;false：不展示

	DisplayHeader *bool `json:"DisplayHeader,omitempty" name:"DisplayHeader"`
}

type MultiTopicSearchInformation struct {

	// 要检索分析的日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时

	Context *string `json:"Context,omitempty" name:"Context"`
}

type CheckFunctionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败错误码

		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
		// 失败错误信息

		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckFunctionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFunctionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExternalDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 外部数据源配置ID

		Id *string `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateExternalDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExternalDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户配置数据

		Data []*UserConfigInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitPartitionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分裂结果集

		Partitions []*PartitionInfo `json:"Partitions,omitempty" name:"Partitions"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SplitPartitionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SplitPartitionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNoticeContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNoticeContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNoticeContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScheduledSqlRequest struct {
	*tchttp.BaseRequest

	// 源日志主题

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 任务启动状态.&nbsp;1开启,&nbsp;2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 定时SQL分析的目标日志主题

	DstResource *ScheduledSqlResouceInfo `json:"DstResource,omitempty" name:"DstResource"`
	// 查询语句

	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitempty" name:"ScheduledSqlContent"`
	// 调度周期(分钟)

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 单次查询的时间窗口.&nbsp;例子中为近15分钟

	ProcessTimeWindow *string `json:"ProcessTimeWindow,omitempty" name:"ProcessTimeWindow"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 源topicId的地域信息

	SrcTopicRegion *string `json:"SrcTopicRegion,omitempty" name:"SrcTopicRegion"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 语法规则。&nbsp;默认值为0。&nbsp;0：Lucene语法，1：CQL语法

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *ModifyScheduledSqlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScheduledSqlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMetricResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标查询结果类型

		ResultType *string `json:"ResultType,omitempty" name:"ResultType"`
		// 指标查询结果

		Result *string `json:"Result,omitempty" name:"Result"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryMetricResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMetricResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsumerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消费组标识

		ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConsumerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordingRuleTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预聚合任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecordingRuleTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecordingRuleTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineInfo struct {

	// 机器的IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 机器状态，0:异常，1:正常

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 机器离线时间，空为正常，异常返回具体时间

	OfflineTime *string `json:"OfflineTime,omitempty" name:"OfflineTime"`
	// 机器是否开启自动升级。0:关闭，1:开启

	AutoUpdate *int64 `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 机器当前版本号。

	Version *string `json:"Version,omitempty" name:"Version"`
	// 机器升级功能状态。&nbsp;0：升级成功；1：升级中；-1：升级失败。

	UpdateStatus *int64 `json:"UpdateStatus,omitempty" name:"UpdateStatus"`
	// 机器升级结果标识。&nbsp;0：成功；1200：升级成功；其他值表示异常。

	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`
	// 机器升级结果信息。&nbsp;“ok”：成功；“update&nbsp;success”：升级成功；其他值为失败原因。

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 机器实例ID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type TopicExtendInfo struct {

	// 日志主题免鉴权配置信息

	AnonymousAccess *AnonymousInfo `json:"AnonymousAccess,omitempty" name:"AnonymousAccess"`
}

type EstimateRebuildIndexTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 预估任务起始时间，毫秒

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 预估任务结束时间，毫秒

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *EstimateRebuildIndexTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EstimateRebuildIndexTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmDetailRequest struct {
	*tchttp.BaseRequest

	// 告警历史快照id

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
}

func (r *DescribeAlarmDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigurationContentInfo struct {

	// 内容格式，只支持json

	Format *string `json:"Format,omitempty" name:"Format"`
	// json格式内容描述

	Json *JsonInfo `json:"Json,omitempty" name:"Json"`
}

type CreateBinlogSubscribeRequest struct {
	*tchttp.BaseRequest

	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z,&nbsp;_,-,zh字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 数据库类型。&nbsp;1:云mysql&nbsp;2:TDSQL-C&nbsp;3:自建mysql

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 数据库配置信息。

	SqlInfo *BinlogDBConfig `json:"SqlInfo,omitempty" name:"SqlInfo"`
	// binlog采集起始点配置信息。

	InitialPoint *BinlogInitialPoint `json:"InitialPoint,omitempty" name:"InitialPoint"`
	// binlog采集事件配置。

	Events []*string `json:"Events,omitempty" name:"Events"`
	// binlog采集元数据配置。

	Metadatas []*string `json:"Metadatas,omitempty" name:"Metadatas"`
	// 时间戳来源。&nbsp;1:采集时间&nbsp;2:事件时间

	TimestampType *uint64 `json:"TimestampType,omitempty" name:"TimestampType"`
	// 黑名单状态。&nbsp;1:关闭&nbsp;2:开启

	BlockListStatus *uint64 `json:"BlockListStatus,omitempty" name:"BlockListStatus"`
	// 黑名单配置信息，当BlockListStatus&nbsp;=2时，BlockList必填。

	BlockList []*BinlogFilterList `json:"BlockList,omitempty" name:"BlockList"`
	// 白名单状态。&nbsp;1:关闭&nbsp;2:开启

	AllowListStatus *uint64 `json:"AllowListStatus,omitempty" name:"AllowListStatus"`
	// 白名单配置信息，当AllowListStatus=2时，AllowList必填。

	AllowList []*BinlogFilterList `json:"AllowList,omitempty" name:"AllowList"`
	// 平铺事件数据标记。&nbsp;默认Flatten为false。&nbsp;Flatten为false时：会将事件数据以数组+JSON格式集中打包到old_data和data两个字段中。如:&nbsp;在数据库操作中执行了两次操作，一次更新了a和b字段，一次d和e，&nbsp;日志值为old_data:[{a:a1,b:b1},&nbsp;{d:d1,e:e1}],&nbsp;data:[{a:a2,b:b2},&nbsp;{d:d2,e:e2}]。&nbsp;Flatten为true时：会将不同的事件拆分为多条日志，并在每条日志中对被操作的字段进行平铺。如：在数据库操作中执行了两次操作，一次更新了a和b字段，一次d和e，&nbsp;则会生成两条日志，&nbsp;分别的值为&nbsp;a:a2,&nbsp;old_a:a1,&nbsp;b:b2,&nbsp;old_b:b1和d:d2,&nbsp;old_d:d1,&nbsp;e:e2,&nbsp;old_e:e1。

	Flatten *bool `json:"Flatten,omitempty" name:"Flatten"`
}

func (r *CreateBinlogSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBinlogSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS控制台创建的TopicId

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
	// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]

	Compression *int64 `json:"Compression,omitempty" name:"Compression"`
	// kafka协议消费数据格式

	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitempty" name:"ConsumerContent"`
	// 是否开启投递服务日志。1：关闭，2：开启。&nbsp;默认值：2

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *OpenKafkaConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKafkaConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardSubscribeData struct {

	// 仪表盘订阅时间，为空标识取仪表盘默认的时间。

	DashboardTime []*string `json:"DashboardTime,omitempty" name:"DashboardTime"`
	// 仪表盘样式布局。0：网格布局，1：单列布局。

	StyleLayout *uint64 `json:"StyleLayout,omitempty" name:"StyleLayout"`
	// 仪表盘订阅模板变量。

	TemplateVariables []*DashboardTemplateVariable `json:"TemplateVariables,omitempty" name:"TemplateVariables"`
	// 仪表盘订阅通知方式。

	NoticeModes []*DashboardNoticeMode `json:"NoticeModes,omitempty" name:"NoticeModes"`
	// 时区。参考：https://en.wikipedia.org/wiki/List_of_tz_database_time_zones

	Timezone *string `json:"Timezone,omitempty" name:"Timezone"`
	// 语言。

	SubscribeLanguage *string `json:"SubscribeLanguage,omitempty" name:"SubscribeLanguage"`
	// 调用链接域名。http://&nbsp;或者&nbsp;https://&nbsp;开头，不能/结尾

	JumpDomain *string `json:"JumpDomain,omitempty" name:"JumpDomain"`
	// 自定义跳转链接。

	JumpUrl *string `json:"JumpUrl,omitempty" name:"JumpUrl"`
}

type DemonstrationResource struct {

	// 资源类型：'LOGSET'&nbsp;|&nbsp;'TOPIC'&nbsp;|&nbsp;'DASHBOARD'&nbsp;|&nbsp;'ALARM'&nbsp;|&nbsp;'ALARM_NOTICE'

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 是否启用。目前只用于表示Topic是否开启日志自动写入。true代表开启日志写入，false代表关闭日志写入

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 资源所在地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 演示示例子类型

	SubType *string `json:"SubType,omitempty" name:"SubType"`
	// 资源类型为TOPIC的上传状态：'INITIAL'&nbsp;|&nbsp;'UPLOADING'&nbsp;|&nbsp;'UPLOAD_FAILED'&nbsp;|&nbsp;'STOPPED'

	UploadStatus *string `json:"UploadStatus,omitempty" name:"UploadStatus"`
	// 模版项ID

	TemplateItemId *string `json:"TemplateItemId,omitempty" name:"TemplateItemId"`
	// 创建者：CLS,&nbsp;USER

	Creator *string `json:"Creator,omitempty" name:"Creator"`
}

type GroupTriggerConditionInfo struct {

	// 分组触发字段名称

	Key *string `json:"Key,omitempty" name:"Key"`
	// 分组触发字段值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type NoticeReceiver struct {

	// 接受者类型。可选值：&nbsp;&nbsp;Uin&nbsp;-&nbsp;用户ID&nbsp;Group&nbsp;-&nbsp;用户组ID&nbsp;暂不支持其余接收者类型

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 接收者。&nbsp;当ReceiverType为Uin时，ReceiverIds的值为用户uid。子用户信息查询&nbsp;当ReceiverType为Group时，ReceiverIds的值为用户组id。CAM用户组

	ReceiverIds []*int64 `json:"ReceiverIds,omitempty" name:"ReceiverIds"`
	// 通知接收渠道。&nbsp;&nbsp;Email&nbsp;-&nbsp;邮件&nbsp;Sms&nbsp;-&nbsp;短信&nbsp;WeChat&nbsp;-&nbsp;微信&nbsp;Phone&nbsp;-&nbsp;电话

	ReceiverChannels []*string `json:"ReceiverChannels,omitempty" name:"ReceiverChannels"`
	// 允许接收信息的开始时间。格式：15:04:05。必填

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 允许接收信息的结束时间。格式：15:04:05。必填

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 位序。&nbsp;&nbsp;入参时无效。&nbsp;出参时有效。

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 通知内容模板ID，使用Default-zh引用默认模板，使用Default-en引用DefaultTemplate(English)。

	NoticeContentId *string `json:"NoticeContentId,omitempty" name:"NoticeContentId"`
}

type QcloudInterfacePara struct {

	// 恒传1

	Multi *uint64 `json:"Multi,omitempty" name:"Multi"`
	// 开通的物品列表，以下参数为每个物品相关的参数

	GoodsInfoList []*QcloudGoodsInfoListInfo `json:"GoodsInfoList,omitempty" name:"GoodsInfoList"`
}

type DescribeBinlogSubscribePreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// binlog日志

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogSubscribePreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBinlogSubscribePreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigurationTemplatesRequest struct {
	*tchttp.BaseRequest

	// templateId按照【配置模板id】进行过滤。类型：String必选：否&nbsp;name按照【配置模板名称】进行过滤。类型：String必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigurationTemplatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigurationTemplatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaConsumerTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// kafka协议消费主题列表

		KafkaConsumers []*KafkaConsumerInfo `json:"KafkaConsumers,omitempty" name:"KafkaConsumers"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaConsumerTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumerTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题ID

		TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
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

type DescribeBinlogSubscribeConnectivityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBinlogSubscribeConnectivityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBinlogSubscribeConnectivityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentConfigsRequest struct {
	*tchttp.BaseRequest

	// agent的版本号

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// 标签列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// agent的instance&nbsp;id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAgentConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigurationTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置模板信息

		Data []*ConfigurationTemplateInfo `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigurationTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigurationTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaUserRequest struct {
	*tchttp.BaseRequest

	// kafka消费用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

func (r *DescribeKafkaUserRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaUserRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogHistogramResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 统计周期：&nbsp;单位ms

		Interval *int64 `json:"Interval,omitempty" name:"Interval"`
		// 命中关键字的日志总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 周期内统计结果详情

		HistogramInfos []*HistogramInfo `json:"HistogramInfos,omitempty" name:"HistogramInfos"`
		// 返回语法优化后的语句（QueryOptimize&nbsp;为&nbsp;1&nbsp;时返回，其他情况返回空字符串）

		Query *string `json:"Query,omitempty" name:"Query"`
		// 日志沉降边界。Unix时间戳（毫秒），在border之前的数据为冷存储，border之后的为标准存储。

		Border *int64 `json:"Border,omitempty" name:"Border"`
		// 多日志主题检索，各个日志主题信息

		Topics *SearchLogTopics `json:"Topics,omitempty" name:"Topics"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogHistogramResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogHistogramResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略ID。

		AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumersRequest struct {
	*tchttp.BaseRequest

	// -&nbsp;consumerId&nbsp;按照【投递规则ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;-&nbsp;topicId&nbsp;按照【日志主题】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;-&nbsp;taskStatus&nbsp;按照【任务运行状态】进行过滤。&nbsp;支持`0`：停止，`1`：运行中，`2`：异常&nbsp;类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为10。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConsumersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRecordingRuleYamlTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRecordingRuleYamlTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRecordingRuleYamlTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyRemoteWriteTaskRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 任务状态&nbsp;0&nbsp;关闭&nbsp;1&nbsp;开启

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// RemoteWrite任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 1&nbsp;内网&nbsp;2外网

	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 目标服务名称

	Target *string `json:"Target,omitempty" name:"Target"`
	// 目标地址

	RemoteWriteURL *string `json:"RemoteWriteURL,omitempty" name:"RemoteWriteURL"`
	// 0:&nbsp;无鉴权&nbsp;1:&nbsp;basic_auth&nbsp;2:&nbsp;token

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 鉴权信息

	AuthInfo *RemoteWriteAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
	// 子网

	SubNet *string `json:"SubNet,omitempty" name:"SubNet"`
	// 后端服务类型&nbsp;-1&nbsp;没有&nbsp;0&nbsp;CVM&nbsp;1025&nbsp;CLB

	VirtualGatewayType *int64 `json:"VirtualGatewayType,omitempty" name:"VirtualGatewayType"`
	// 云时序数据库实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *ModifyRemoteWriteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteWriteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ValueInfo struct {

	// 字段类型，目前支持的类型有：long、text、double

	Type *string `json:"Type,omitempty" name:"Type"`
	// 字段的分词符，其中的每个字符代表一个分词符；&nbsp;仅支持英文符号、\n\t\r及转义符\；&nbsp;long及double类型字段需为空；&nbsp;注意：\n\t\r本身已被转义，直接使用双引号包裹即可作为入参，无需再次转义。使用API&nbsp;Explorer进行调试时请使用JSON参数输入方式，以避免\n\t\r被重复转义

	Tokenizer *string `json:"Tokenizer,omitempty" name:"Tokenizer"`
	// 字段是否开启分析功能。true表示开启分析功能，false表示关闭分析功能。

	SqlFlag *bool `json:"SqlFlag,omitempty" name:"SqlFlag"`
	// 是否包含zh，long及double类型字段需为false。true表示包含zh，false表示不包含zh。

	ContainZH *bool `json:"ContainZH,omitempty" name:"ContainZH"`
	// 索引健值别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 只给子节点开启索引

	OpenIndexForChildOnly *bool `json:"OpenIndexForChildOnly,omitempty" name:"OpenIndexForChildOnly"`
	// 子节点列表

	ChildNode []*KeyValueInfo `json:"ChildNode,omitempty" name:"ChildNode"`
}

type DescribeMachineGroupConfigsRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
}

func (r *DescribeMachineGroupConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGroupConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AdvancedConsumerConfiguration struct {

	// Ckafka分区hash状态。&nbsp;默认&nbsp;false&nbsp;&nbsp;true：开启根据字段&nbsp;Hash&nbsp;值结果相等的信息投递到统一&nbsp;ckafka&nbsp;分区&nbsp;false：关闭根据字段&nbsp;Hash&nbsp;值结果相等的信息投递到统一&nbsp;ckafka&nbsp;分区

	PartitionHashStatus *bool `json:"PartitionHashStatus,omitempty" name:"PartitionHashStatus"`
	// 需要计算&nbsp;hash&nbsp;的字段列表。

	PartitionFields []*string `json:"PartitionFields,omitempty" name:"PartitionFields"`
}

type DescribeCosRechargesRequest struct {
	*tchttp.BaseRequest

	// 日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 状态&nbsp;&nbsp;&nbsp;status&nbsp;0:&nbsp;created,&nbsp;1:&nbsp;running,&nbsp;2:&nbsp;pause,&nbsp;3:&nbsp;finished,&nbsp;4:&nbsp;failed。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 是否启用:&nbsp;&nbsp;&nbsp;0：&nbsp;未启用&nbsp;&nbsp;，&nbsp;1：启用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
}

func (r *DescribeCosRechargesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosRechargesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenClsServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenClsServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenClsServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DashboardNoticeMode struct {

	// 仪表盘通知方式。&nbsp;&nbsp;Uin：云用户&nbsp;Group：云用户组&nbsp;Email：自定义Email&nbsp;WeCom:&nbsp;企业微信回调

	ReceiverType *string `json:"ReceiverType,omitempty" name:"ReceiverType"`
	// 通知方式对应的值。&nbsp;&nbsp;当ReceiverType不是&nbsp;Wecom&nbsp;时，Values必填。

	Values []*string `json:"Values,omitempty" name:"Values"`
	// 仪表盘通知渠道。&nbsp;&nbsp;支持：["Email","Sms","WeChat","Phone"]。&nbsp;当ReceiverType是&nbsp;Email&nbsp;或&nbsp;Wecom&nbsp;时，ReceiverChannels不能赋值。

	ReceiverChannels []*string `json:"ReceiverChannels,omitempty" name:"ReceiverChannels"`
	// 回调Url。&nbsp;&nbsp;当ReceiverType是&nbsp;Wecom&nbsp;时，Url必填。&nbsp;当ReceiverType不是&nbsp;Wecom&nbsp;时，Url不能填写。

	Url *string `json:"Url,omitempty" name:"Url"`
}

type CheckFunctionRequest struct {
	*tchttp.BaseRequest

	// 用户输入的加工语句

	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
	// 加工任务目的topic_id以及别名

	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
	// 数据加工目标主题的类型.&nbsp;1&nbsp;固定主题&nbsp;2动态创建

	FuncType *int64 `json:"FuncType,omitempty" name:"FuncType"`
}

func (r *CheckFunctionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckFunctionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器组ID

		GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdleResourcePolicyRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeIdleResourcePolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdleResourcePolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClsServiceRequest struct {
	*tchttp.BaseRequest
}

func (r *GetClsServiceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClsServiceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmShieldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmShieldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmShieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardSubscribeRequest struct {
	*tchttp.BaseRequest

	// 仪表盘订阅id。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 仪表盘id。

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘订阅名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 订阅时间cron表达式，格式为：{秒数}&nbsp;{分钟}&nbsp;{小时}&nbsp;{日期}&nbsp;{月份}&nbsp;{星期}；（有效数据为：{分钟}&nbsp;{小时}&nbsp;{日期}&nbsp;{月份}&nbsp;{星期}）。

	Cron *string `json:"Cron,omitempty" name:"Cron"`
	// 仪表盘订阅数据。

	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitempty" name:"SubscribeData"`
}

func (r *ModifyDashboardSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyExternalDataSourceRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID,&nbsp;此外部数据源配置隶属此日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 外部数据源ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 日志集ID,&nbsp;表示该日志集下面的所有日志主题都可以使用此外部数据源配置

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 备注，可选参数，长度不超过255个字符

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 数据源类型。支持1,2,3,4&nbsp;&nbsp;1：自建MySQL&nbsp;2：COS(csv格式)&nbsp;3：云数据库MySQL&nbsp;4：云数据库TDSQL-C&nbsp;MySQL

	Datasource *uint64 `json:"Datasource,omitempty" name:"Datasource"`
	// SQL相关配置信息，当Datasource值为1时，此字段生效

	SQLInfo *ExternalDataSourceSQLInfo `json:"SQLInfo,omitempty" name:"SQLInfo"`
	// COS相关配置信息，当Datasource值为2时，此字段生效

	COSInfo *ExternalDataSourceCOSInfo `json:"COSInfo,omitempty" name:"COSInfo"`
}

func (r *ModifyExternalDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyExternalDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformProcessInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据加工任务进度详请

		TaskLogStatistics []*TaskLogStatistic `json:"TaskLogStatistics,omitempty" name:"TaskLogStatistics"`
		// 读取的源日志主题的总行数

		ReadLineSum *uint64 `json:"ReadLineSum,omitempty" name:"ReadLineSum"`
		// 加工后的总行数

		WriteLineSum *uint64 `json:"WriteLineSum,omitempty" name:"WriteLineSum"`
		// 加工失败的总行数

		FailedLineSum *uint64 `json:"FailedLineSum,omitempty" name:"FailedLineSum"`
		// 加工过滤的总行数

		FilterLineSum *uint64 `json:"FilterLineSum,omitempty" name:"FilterLineSum"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformProcessInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformProcessInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTemplatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTemplatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTemplatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNoticeContentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通知内容配置ID

		NoticeContentId *string `json:"NoticeContentId,omitempty" name:"NoticeContentId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNoticeContentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNoticeContentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警策略列表。

		Alarms []*AlarmInfo `json:"Alarms,omitempty" name:"Alarms"`
		// 符合查询条件的告警策略数目。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricSubscribePreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 成功数量

		SuccessCount *uint64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
		// 失败数量

		FailCount *uint64 `json:"FailCount,omitempty" name:"FailCount"`
		// 成功实例数据

		SuccessInstances []*InstanceData `json:"SuccessInstances,omitempty" name:"SuccessInstances"`
		// 失败实例数据

		FailInstances []*InstanceData `json:"FailInstances,omitempty" name:"FailInstances"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricSubscribePreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricSubscribePreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordingRuleProcessInfoRequest struct {
	*tchttp.BaseRequest

	// 源日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 执行开始时间戳，单位ms

	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`
	// 执行结束时间戳，单位ms

	Endtime *int64 `json:"Endtime,omitempty" name:"Endtime"`
	// 实例ID

	ProcessId *string `json:"ProcessId,omitempty" name:"ProcessId"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 调度结果，1:运行中&nbsp;2:成功&nbsp;3:失败.&nbsp;默认为0不做过滤

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 排序字段，可选字段为：process_start_time&nbsp;|&nbsp;time_window_start_time。&nbsp;默认为process_start_time

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序顺序，DESC&nbsp;|&nbsp;ASC。&nbsp;默认DESC

	OrderDirection *string `json:"OrderDirection,omitempty" name:"OrderDirection"`
}

func (r *DescribeRecordingRuleProcessInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordingRuleProcessInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDashboardRequest struct {
	*tchttp.BaseRequest

	// 仪表盘id

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘名称

	DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
	// 仪表盘配置数据

	Data *string `json:"Data,omitempty" name:"Data"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 文件夹id。为空则不修改仪表盘关联的文件夹id。

	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *ModifyDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyScheduledSqlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyScheduledSqlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyScheduledSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExternalDataSourceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExternalDataSourceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExternalDataSourceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetricSeriesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 时序指标series

		Values *string `json:"Values,omitempty" name:"Values"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetMetricSeriesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricSeriesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NoticeContentTemplate struct {

	// 通知内容模板ID。

	NoticeContentId *string `json:"NoticeContentId,omitempty" name:"NoticeContentId"`
	// 通知内容模板名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 语言类型。&nbsp;&nbsp;0：&nbsp;zh&nbsp;1：&nbsp;英文

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 通知内容模板信息。

	NoticeContents []*NoticeContent `json:"NoticeContents,omitempty" name:"NoticeContents"`
	// 通知内容模板标记。&nbsp;&nbsp;0：&nbsp;用户自定义&nbsp;1：&nbsp;系统内置

	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
	// 创建者主账号。

	Uin *uint64 `json:"Uin,omitempty" name:"Uin"`
	// 创建/修改者子账号。

	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
	// 创建时间&nbsp;秒级时间戳。

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间&nbsp;秒级时间戳。

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ScheduledSqlTaskInfo struct {

	// ScheduledSql任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// ScheduledSql任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 源日志主题id

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 源日志主题名称

	SrcTopicName *string `json:"SrcTopicName,omitempty" name:"SrcTopicName"`
	// 定时SQL分析目标主题

	DstResource *ScheduledSqlResouceInfo `json:"DstResource,omitempty" name:"DstResource"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 任务状态，1:运行&nbsp;2:停止&nbsp;3:异常-找不到源日志主题&nbsp;4:异常-找不到目标主题&nbsp;&nbsp;5:&nbsp;访问权限问题&nbsp;6:内部故障&nbsp;7:其他故障

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务启用状态，1开启,&nbsp;2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 查询语句

	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitempty" name:"ScheduledSqlContent"`
	// 调度开始时间

	ProcessStartTime *string `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 调度类型，1:持续运行&nbsp;2:指定时间范围

	ProcessType *int64 `json:"ProcessType,omitempty" name:"ProcessType"`
	// 调度结束时间，当process_type=2时为必传字段

	ProcessEndTime *string `json:"ProcessEndTime,omitempty" name:"ProcessEndTime"`
	// 调度周期(分钟)

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 查询的时间窗口.&nbsp;@m-15m,&nbsp;@m，意为近15分钟

	ProcessTimeWindow *string `json:"ProcessTimeWindow,omitempty" name:"ProcessTimeWindow"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// 源topicId的地域信息

	SrcTopicRegion *string `json:"SrcTopicRegion,omitempty" name:"SrcTopicRegion"`
	// 语法规则，0：Lucene语法，1：CQL语法

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

type DescribeExternalDataSourcesRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID,&nbsp;此外部数据源配置隶属此日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// name按照【名称】进行过滤。类型：String必选：否&nbsp;id按照【外部数据库配置ID】进行过滤。类型：String必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 个数

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExternalDataSourcesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExternalDataSourcesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenKVRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正则表达式

		Regex *string `json:"Regex,omitempty" name:"Regex"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenKVRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenKVRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNoticeContentRequest struct {
	*tchttp.BaseRequest

	// 通知内容模板ID。

	NoticeContentId *string `json:"NoticeContentId,omitempty" name:"NoticeContentId"`
	// 通知内容模板名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 通知内容语言。&nbsp;&nbsp;0：zh&nbsp;1：英文

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 通知内容模板详细信息

	NoticeContents []*NoticeContent `json:"NoticeContents,omitempty" name:"NoticeContents"`
}

func (r *ModifyNoticeContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNoticeContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCosRechargeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// COS导入任务id

		Id *string `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCosRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCosRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// CLS对应topic标识

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *DescribeKafkaConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAgentNormalResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeAgentNormalResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeAgentNormalResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DynamicIndex struct {

	// 键值索引自动配置开关。true表示开启动态索引，false表示关闭动态索引。

	Status *bool `json:"Status,omitempty" name:"Status"`
}

type CreateDashboardRequest struct {
	*tchttp.BaseRequest

	// 仪表盘名称

	DashboardName *string `json:"DashboardName,omitempty" name:"DashboardName"`
	// 仪表盘配置数据

	Data *string `json:"Data,omitempty" name:"Data"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 文件夹id。为空则自动关联到默认文件夹。

	FolderId *string `json:"FolderId,omitempty" name:"FolderId"`
}

func (r *CreateDashboardRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateDashboardRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformProcessInfoRequest struct {
	*tchttp.BaseRequest

	// 数据加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 要查询的起始时间，Unix时间戳，单位ms

	From *uint64 `json:"From,omitempty" name:"From"`
	// 要查询的结束时间，Unix时间戳，单位ms

	To *uint64 `json:"To,omitempty" name:"To"`
	// 是否需要分成多个时间段获取

	NeedMultTimePeriod *bool `json:"NeedMultTimePeriod,omitempty" name:"NeedMultTimePeriod"`
}

func (r *DescribeDataTransformProcessInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformProcessInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 分页的总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志集列表

		Logsets []*LogsetInfo `json:"Logsets,omitempty" name:"Logsets"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogsetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmInfo struct {

	// 告警策略名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 监控对象列表。

	AlarmTargets []*AlarmTargetInfo `json:"AlarmTargets,omitempty" name:"AlarmTargets"`
	// 监控任务运行时间点。

	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`
	// 触发条件。

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为10。

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 告警重复的周期。单位是min。取值范围是0~1440。

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 关联的告警通知模板列表。

	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`
	// 开启状态。true表示开启，false表示关闭。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 告警策略ID。

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
	// 创建时间。

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 最近更新时间。

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 开启状态。true表示开启，false表示关闭。

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 自定义通知模板

	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`
	// 自定义回调模板

	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`
	// 多维分析设置

	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
	// 分组触发状态。1：开启，0：关闭（默认）

	GroupTriggerStatus *bool `json:"GroupTriggerStatus,omitempty" name:"GroupTriggerStatus"`
	// 分组触发条件。

	GroupTriggerCondition []*string `json:"GroupTriggerCondition,omitempty" name:"GroupTriggerCondition"`
	// 告警策略绑定的标签信息。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 监控对象类型。0:执行语句共用监控对象;1:每个执行语句单独选择监控对象。

	MonitorObjectType *uint64 `json:"MonitorObjectType,omitempty" name:"MonitorObjectType"`
	// 告警级别。0:警告(Warn);1:提醒(Info);2:紧急&nbsp;(Critical)。

	AlarmLevel *uint64 `json:"AlarmLevel,omitempty" name:"AlarmLevel"`
	// 告警附加分类信息列表。

	Classifications []*AlarmClassification `json:"Classifications,omitempty" name:"Classifications"`
	// 触发条件交互模式配置信息。

	ConditionInteractiveConfig *string `json:"ConditionInteractiveConfig,omitempty" name:"ConditionInteractiveConfig"`
	// 多触发条件。与&nbsp;Condition互斥。

	MultiConditions []*MultiCondition `json:"MultiConditions,omitempty" name:"MultiConditions"`
	// 告警模板相关信息

	AlarmTemplateInfo *AlarmTemplateConfig `json:"AlarmTemplateInfo,omitempty" name:"AlarmTemplateInfo"`
}

type GetAlarmLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加载后续内容的Context

		Context *string `json:"Context,omitempty" name:"Context"`
		// 指定时间范围内的告警执行详情是否完整返回

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 返回的结果是否为SQL分析结果

		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`
		// 分析结果的列名，如果Query语句有SQL查询，则返回查询字段的列名；&nbsp;否则为空。

		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`
		// 执行详情查询结果。&nbsp;当Query字段无SQL语句时，返回查询结果。&nbsp;当Query字段有SQL语句时，可能返回null。

		Results []*LogInfo `json:"Results,omitempty" name:"Results"`
		// 执行详情统计分析结果。当Query字段有SQL语句时，返回SQL统计结果，否则可能返回null。

		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`
		// 执行详情统计分析结果；UseNewAnalysis为true有效。

		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`
		// 分析结果的列名，&nbsp;UseNewAnalysis为true有效

		Columns []*Column `json:"Columns,omitempty" name:"Columns"`
		// 返回语法优化后的语句（入参如有QueryOptimize&nbsp;且为&nbsp;1&nbsp;时返回，其他情况返回空字符串）

		Query *string `json:"Query,omitempty" name:"Query"`
		// 当前系统使用的采样率，入参如有SamplingRate时生效（主要是当客户输入0时，返回真实后台的AutoSamplingRate值）

		SamplingRate *float64 `json:"SamplingRate,omitempty" name:"SamplingRate"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetAlarmLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetAlarmLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ShipperTemplateInfo struct {

	// 投递规则的名字，允许字符0-9、a-z、A-Z、&nbsp;_、-、zh字符。&nbsp;必须以${TopicName}-${TaskCreateTime}&nbsp;为前缀；${TopicName}&nbsp;字符最多200个，如果超出截断topicname，-${TaskCreateTime}14个，用户可添加41个自定义字符，总长度255。受限于CLS产品规格255。

	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`
	// 创建的投递规则投递目录的前缀,允许字符&nbsp;a-z、A-Z、0-9、_、-、/，并且以非/开头。${TopicName}&nbsp;字符最多200个，如果超出截断topicname，${TaskCreateTime}14个，年月日时分秒25个，如果用户使用了默认参数，那么用户可添加17个自定义字符，总长度255。&nbsp;受限于COS产品规格255，COS允许zh路径。

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 投递日志的分区规则，支持strftime的时间格式表示

	Partition *string `json:"Partition,omitempty" name:"Partition"`
	// 投递日志的压缩配置

	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`
	// 投递的文件的最大值，单位&nbsp;MB，默认256，范围&nbsp;5-256

	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
	// 投递的时间间隔，单位&nbsp;秒，默认300，范围&nbsp;300-900

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 投递日志的内容格式配置。投递格式：仅支持JSON。

	Content *ConfigurationContentInfo `json:"Content,omitempty" name:"Content"`
	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递

	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`
	// 投递数据范围的开始时间点，不能超出日志主题的生命周期起点。如果用户不填写，默认为用户新建投递任务的时间。

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 投递数据范围的结束时间点，不能填写未来时间。如果用户不填写，默认为持续投递，即无限。

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type DeleteAlarmShieldRequest struct {
	*tchttp.BaseRequest

	// 屏蔽规则id。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 通知渠道组id。

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmShieldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmShieldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteShipperRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
}

func (r *DeleteShipperRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteShipperRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExceptionResources struct {

	// 日志主题Id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 索引配置异常列表

	IndexExceptions []*IndexException `json:"IndexExceptions,omitempty" name:"IndexExceptions"`
	// 采集配置异常列表

	ConfigExceptions []*ConfigException `json:"ConfigExceptions,omitempty" name:"ConfigExceptions"`
	// 日志上传异常列表

	UploadLogExceptions []*UploadLogException `json:"UploadLogExceptions,omitempty" name:"UploadLogExceptions"`
}

type UploadLogException struct {

	// 异常事件ID&nbsp;400：参数错误&nbsp;401：鉴权失败/密钥过期&nbsp;413：消息格式过大&nbsp;429：流控-区分超频/超频

	EventId *int64 `json:"EventId,omitempty" name:"EventId"`
	// 异常事件描述

	Message *string `json:"Message,omitempty" name:"Message"`
	// 最近发生时间

	LatestTime *uint64 `json:"LatestTime,omitempty" name:"LatestTime"`
}

type DescribeAlarmDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警历史记录日志

		RecordLog *string `json:"RecordLog,omitempty" name:"RecordLog"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteWriteTasksRequest struct {
	*tchttp.BaseRequest

	// -&nbsp;taskId&nbsp;按照【任务ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;-&nbsp;topicId&nbsp;按照【日志主题】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;-&nbsp;taskStatus&nbsp;按照【任务运行状态】进行过滤。&nbsp;支持`1`：运行中，`2`：停止，`3`：异常&nbsp;类型：String&nbsp;必选：否&nbsp;-&nbsp;name&nbsp;按照【任务名称】进行模糊过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为10。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRemoteWriteTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteWriteTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源列表

		Resources []*ResourcesInfo `json:"Resources,omitempty" name:"Resources"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricConfigLabel struct {

	// 元数据。&nbsp;支持&nbsp;-&nbsp;`namespace`&nbsp;-&nbsp;`pod_name`&nbsp;-&nbsp;`pod_ip`&nbsp;-&nbsp;`pod_uid`&nbsp;-&nbsp;`container_name`&nbsp;-&nbsp;`container_id`&nbsp;-&nbsp;`image_name`&nbsp;-&nbsp;`cluster_id`&nbsp;-&nbsp;`node_id`&nbsp;-&nbsp;`node_ip`

	Metadata []*string `json:"Metadata,omitempty" name:"Metadata"`
	// 元数据Pod&nbsp;Label信息。

	Label *AppointLabel `json:"Label,omitempty" name:"Label"`
	// 自定义label信息。

	CustomLabels []*CustomLabel `json:"CustomLabels,omitempty" name:"CustomLabels"`
}

type GenBeginRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 正则表达式

		Regex *string `json:"Regex,omitempty" name:"Regex"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GenBeginRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenBeginRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyUserConfigRequest struct {
	*tchttp.BaseRequest

	// 用户配置内容

	Data []*UserConfigInfo `json:"Data,omitempty" name:"Data"`
}

func (r *ModifyUserConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyUserConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadServiceLogRequest struct {
	*tchttp.BaseRequest

	// 主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 根据&nbsp;hashkey&nbsp;写入相应范围的主题分区

	HashKey *string `json:"HashKey,omitempty" name:"HashKey"`
	// agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// agent的Instance&nbsp;ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// agent是否开启自动升级功能

	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// agent请求序列号

	AgentSeq *string `json:"AgentSeq,omitempty" name:"AgentSeq"`
	// agent版本

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// 请求Unique&nbsp;ID

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
	// 服务日志类型

	MetricType *string `json:"MetricType,omitempty" name:"MetricType"`
	// agent的label信息

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// 压缩方法

	CompressType *string `json:"CompressType,omitempty" name:"CompressType"`
}

func (r *UploadServiceLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadServiceLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogContextInfo struct {

	// 日志来源设备

	Source *string `json:"Source,omitempty" name:"Source"`
	// 采集路径

	Filename *string `json:"Filename,omitempty" name:"Filename"`
	// 日志内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 日志包序号

	PkgId *string `json:"PkgId,omitempty" name:"PkgId"`
	// 日志包内一条日志的序号

	PkgLogId *int64 `json:"PkgLogId,omitempty" name:"PkgLogId"`
	// 日志时间戳

	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
	// 日志来源主机名称

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 原始日志(仅在日志创建索引异常时有值)

	RawLog *string `json:"RawLog,omitempty" name:"RawLog"`
	// 日志创建索引异常原因(仅在日志创建索引异常时有值)

	IndexStatus *string `json:"IndexStatus,omitempty" name:"IndexStatus"`
	// 日志内容的高亮描述信息

	HighLights []*HighLightItem `json:"HighLights,omitempty" name:"HighLights"`
}

type DescribeDataTransformFailLogInfoRequest struct {
	*tchttp.BaseRequest

	// 数据加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 要查询的起始时间，Unix时间戳，单位ms

	From *uint64 `json:"From,omitempty" name:"From"`
	// 要查询的结束时间，Unix时间戳，单位ms

	To *uint64 `json:"To,omitempty" name:"To"`
	// 目标日志主题id

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
}

func (r *DescribeDataTransformFailLogInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformFailLogInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消费组详情列表

		ConsumerGroupsInfo []*ConsumerGroupInfo `json:"ConsumerGroupsInfo,omitempty" name:"ConsumerGroupsInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaRechargesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// KafkaRechargeInfo&nbsp;信息列表

		Infos []*KafkaRechargeInfo `json:"Infos,omitempty" name:"Infos"`
		// Kafka导入信息总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaRechargesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaRechargesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRebuildIndexTasksRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 索引重建任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 索引重建任务状态，不填返回所有状态任务列表，多种状态之间用逗号分隔，0:索引重建任务已创建，1:已创建索引重建资源，2:重建中，3:重建完成，4:重建成功（可检索），5:任务取消，6:元数据和索引已删除

	Status *string `json:"Status,omitempty" name:"Status"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为10，最大值20。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeRebuildIndexTasksRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRebuildIndexTasksRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PreviewKafkaRechargeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志样例，PreviewType为2时返回

		LogSample *string `json:"LogSample,omitempty" name:"LogSample"`
		// 日志预览结果

		LogData *string `json:"LogData,omitempty" name:"LogData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PreviewKafkaRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreviewKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExportsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordingRuleProcessInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预聚合任务进度信息列表

		RecordingRuleProcessInfos []*RecordingRuleProcessInfo `json:"RecordingRuleProcessInfos,omitempty" name:"RecordingRuleProcessInfos"`
		// 任务总次数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 执行成功数

		TotalSuccess *uint64 `json:"TotalSuccess,omitempty" name:"TotalSuccess"`
		// 执行失败数

		TotalFailed *uint64 `json:"TotalFailed,omitempty" name:"TotalFailed"`
		// 执行运行中

		TotalRunning *uint64 `json:"TotalRunning,omitempty" name:"TotalRunning"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordingRuleProcessInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordingRuleProcessInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 告警通知模板

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

func (r *DeleteAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmShieldsRequest struct {
	*tchttp.BaseRequest

	// 通知渠道组id。

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
	// -&nbsp;taskId:按照【规则id】进行过滤。类型：String&nbsp;必选：否&nbsp;-&nbsp;status:按照【规则状态】进行过滤。类型：String。&nbsp;支持&nbsp;0:暂未生效，1:生效中，2:已失效。&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmShieldsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmShieldsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformAutoCreateLogSetsRequest struct {
	*tchttp.BaseRequest

	// 任务列表id列表

	TaskIds []*string `json:"TaskIds,omitempty" name:"TaskIds"`
	// logsetId&nbsp;按照【日志集id】进行过滤。&nbsp;类型：String&nbsp;&nbsp;必选：否&nbsp;&nbsp;&nbsp;&nbsp;logsetName&nbsp;按照【日志集名称】进行过滤。&nbsp;类型：String&nbsp;&nbsp;必选：否&nbsp;&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDataTransformAutoCreateLogSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformAutoCreateLogSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DescribeIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordingRuleTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRecordingRuleTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecordingRuleTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmShieldsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的规则总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 告警屏蔽未生效数量

		InvalidCount *uint64 `json:"InvalidCount,omitempty" name:"InvalidCount"`
		// 告警屏蔽生效中数量

		ValidCount *uint64 `json:"ValidCount,omitempty" name:"ValidCount"`
		// 告警屏蔽已过期数量

		ExpireCount *uint64 `json:"ExpireCount,omitempty" name:"ExpireCount"`
		// 告警屏蔽规则详情

		AlarmShields []*AlarmShieldInfo `json:"AlarmShields,omitempty" name:"AlarmShields"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlarmShieldsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmShieldsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递任务列表

		Tasks []*ShipperTaskInfo `json:"Tasks,omitempty" name:"Tasks"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShipperTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShipperTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SearchLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 透传本次接口返回的Context值，可获取后续更多日志，过期时间1小时。&nbsp;注意：&nbsp;&nbsp;仅适用于单日志主题检索，检索多个日志主题时，请使用Topics中的Context

		Context *string `json:"Context,omitempty" name:"Context"`
		// 符合检索条件的日志是否已全部返回，如未全部返回可使用Context参数获取后续更多日志&nbsp;注意：仅当检索分析语句(Query)不包含SQL时有效

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 返回的是否为统计分析（即SQL）结果

		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`
		// 日志统计分析结果的列名&nbsp;当UseNewAnalysis为false时生效

		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`
		// 匹配检索条件的原始日志

		Results []*LogInfo `json:"Results,omitempty" name:"Results"`
		// 日志统计分析结果&nbsp;当UseNewAnalysis为false时生效

		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`
		// 日志统计分析结果&nbsp;当UseNewAnalysis为true时生效

		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`
		// 日志统计分析结果的列属性&nbsp;当UseNewAnalysis为true时生效

		Columns []*Column `json:"Columns,omitempty" name:"Columns"`
		// 返回语法优化后的语句（QueryOptimize&nbsp;为&nbsp;1&nbsp;时返回，其他情况返回空字符串）

		Query *string `json:"Query,omitempty" name:"Query"`
		// 本次统计分析使用的采样率

		SamplingRate *float64 `json:"SamplingRate,omitempty" name:"SamplingRate"`
		// 日志沉降边界。Unix时间戳（毫秒），在border之前的数据为冷存储，border之后的为标准存储。

		Border *int64 `json:"Border,omitempty" name:"Border"`
		// 使用多日志主题检索时，各个日志主题的基本信息，例如报错信息。

		Topics *SearchLogTopics `json:"Topics,omitempty" name:"Topics"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type CommitConsumerOffsetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommitConsumerOffsetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CommitConsumerOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachineGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachineGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachineGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ApplyConfigurationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyConfigurationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ApplyConfigurationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckConfigRegexRequest struct {
	*tchttp.BaseRequest

	// 正则表达式

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 要匹配的日志内容

	Content *string `json:"Content,omitempty" name:"Content"`
}

func (r *CheckConfigRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckConfigRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteWriteTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 导入配置ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 任务运行状态&nbsp;1:&nbsp;运行中&nbsp;2:暂停&nbsp;3:&nbsp;失败

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *DescribeRemoteWriteTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteWriteTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMetricSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMetricSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMetricSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetricCollectConfig struct {

	// 采集配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 日志主题id。

	TopicIds []*string `json:"TopicIds,omitempty" name:"TopicIds"`
	// 采集配置来源。支持&nbsp;：`0`、`1`&nbsp;-&nbsp;0:自建k8s&nbsp;-&nbsp;1:TKE

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 机器组id。

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// 监控类型。支持&nbsp;：`0`、`1`，不支持修改&nbsp;-&nbsp;0:基础监控&nbsp;-&nbsp;1:自定义监控,

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 采集配置方式。支持&nbsp;：`0`、`1`，不支持修改&nbsp;-&nbsp;0:普通配置方式，Type字段只能为：`1`&nbsp;-&nbsp;1:YAML导入方式，Type&nbsp;可以是：`0`或者`1`

	Flag *uint64 `json:"Flag,omitempty" name:"Flag"`
	// 名称：长度不超过253字符

	Name *string `json:"Name,omitempty" name:"Name"`
	// 采集对象,&nbsp;Flag=0时生效

	Spec *MetricSpec `json:"Spec,omitempty" name:"Spec"`
	// 标签处理,&nbsp;Flag=0时生效

	Relabels []*Relabeling `json:"Relabels,omitempty" name:"Relabels"`
	// 标签处理,&nbsp;Flag=0时生效

	MetricRelabels []*Relabeling `json:"MetricRelabels,omitempty" name:"MetricRelabels"`
	// 自定义元数据,&nbsp;Flag=0时生效

	MetricLabel *MetricConfigLabel `json:"MetricLabel,omitempty" name:"MetricLabel"`
	// 通信协议&nbsp;`http`、`https`；Flag=0时生效

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 采集频率,&nbsp;Flag=0时生效

	ScrapeInterval *string `json:"ScrapeInterval,omitempty" name:"ScrapeInterval"`
	// 采集超时时间。&nbsp;Flag=0&nbsp;&&&nbsp;Type=1时生效

	ScrapeTimeout *string `json:"ScrapeTimeout,omitempty" name:"ScrapeTimeout"`
	// Prometheus如何处理标签之间的冲突。当Flag=0生效，支持`true`,`false`&nbsp;-&nbsp;`false`:配置数据中冲突的标签重命名&nbsp;-&nbsp;`true`:忽略冲突的服务器端标签

	HonorLabels *bool `json:"HonorLabels,omitempty" name:"HonorLabels"`
	// 压缩,&nbsp;false:关闭,&nbsp;true:开启,&nbsp;Flag=0时生效

	EnableCompression *bool `json:"EnableCompression,omitempty" name:"EnableCompression"`
	// 采集配置yaml格式字符串,&nbsp;Flag=1时必填

	YamlSpec *MetricYamlSpec `json:"YamlSpec,omitempty" name:"YamlSpec"`
	// 操作状态,0:应用,1:暂停

	Operate *uint64 `json:"Operate,omitempty" name:"Operate"`
	// 创建时间戳&nbsp;秒级

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间戳&nbsp;秒级

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 日志主题扩展信息

	TopicExtendInfo *string `json:"TopicExtendInfo,omitempty" name:"TopicExtendInfo"`
}

type CreateMetricSubscribeRequest struct {
	*tchttp.BaseRequest

	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z,&nbsp;_,-,zh字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 云产品命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 数据库配置信息。

	Metrics []*MetricConfig `json:"Metrics,omitempty" name:"Metrics"`
	// 实例配置配置。

	InstanceInfo *InstanceConfig `json:"InstanceInfo,omitempty" name:"InstanceInfo"`
}

func (r *CreateMetricSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMetricSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigExtrasResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集配置列表

		Configs []*ConfigExtraInfo `json:"Configs,omitempty" name:"Configs"`
		// 过滤到的总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConfigExtrasResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigExtrasResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryShipperTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryShipperTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryShipperTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserConfigRequest struct {
	*tchttp.BaseRequest

	// 要删除key列表

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
}

func (r *DeleteUserConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BinlogConfigInfo struct {

	// 配置id。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 名称：长度不超过64字符，以字母开头，接受0-9,a-z,A-Z,&nbsp;_,-,zh字符。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 数据库类型。&nbsp;1:云mysql&nbsp;2:TDSQL-C&nbsp;3:自建mysql

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 数据库配置信息。

	SqlInfo *BinlogDBConfig `json:"SqlInfo,omitempty" name:"SqlInfo"`
	// binlog采集起始点配置信息。

	InitialPoint *BinlogInitialPoint `json:"InitialPoint,omitempty" name:"InitialPoint"`
	// binlog采集事件配置。

	Events []*string `json:"Events,omitempty" name:"Events"`
	// binlog采集元数据配置。

	Metadatas []*string `json:"Metadatas,omitempty" name:"Metadatas"`
	// 时间戳来源。&nbsp;1:采集时间&nbsp;2:事件时间

	TimestampType *uint64 `json:"TimestampType,omitempty" name:"TimestampType"`
	// 黑名单状态。&nbsp;1:关闭&nbsp;2:开启

	BlockListStatus *uint64 `json:"BlockListStatus,omitempty" name:"BlockListStatus"`
	// 黑名单配置。

	BlockList []*BinlogFilterList `json:"BlockList,omitempty" name:"BlockList"`
	// 白名单状态。&nbsp;1:关闭&nbsp;2:开启

	AllowListStatus *uint64 `json:"AllowListStatus,omitempty" name:"AllowListStatus"`
	// 白名单配置。

	AllowList []*BinlogFilterList `json:"AllowList,omitempty" name:"AllowList"`
	// 平铺事件数据标记。&nbsp;默认Flatten为false。&nbsp;Flatten为false时：会将事件数据以数组+JSON格式集中打包到old_data和data两个字段中。如:&nbsp;在数据库操作中执行了两次操作，一次更新了a和b字段，一次d和e，&nbsp;日志值为old_data:[{a:a1,b:b1},&nbsp;{d:d1,e:e1}],&nbsp;data:[{a:a2,b:b2},&nbsp;{d:d2,e:e2}]。&nbsp;Flatten为true时：会将不同的事件拆分为多条日志，并在每条日志中对被操作的字段进行平铺。如：在数据库操作中执行了两次操作，一次更新了a和b字段，一次d和e，&nbsp;则会生成两条日志，&nbsp;分别的值为&nbsp;a:a2,&nbsp;old_a:a1,&nbsp;b:b2,&nbsp;old_b:b1和d:d2,&nbsp;old_d:d1,&nbsp;e:e2,&nbsp;old_e:e1。

	Flatten *bool `json:"Flatten,omitempty" name:"Flatten"`
	// 配置状态。&nbsp;1:&nbsp;未启用&nbsp;2:&nbsp;启用

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 配置状态标记。&nbsp;1:&nbsp;未启用&nbsp;2:&nbsp;启用&nbsp;3:&nbsp;异常

	StatusFlag *uint64 `json:"StatusFlag,omitempty" name:"StatusFlag"`
	// 错误信息。

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 创建时间。秒级时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间。秒级时间戳

	UpdateTime *uint64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DescribeMetricCorrectDimensionRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeMetricCorrectDimensionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricCorrectDimensionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordingRuleYamlTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// yamlConfigName【配置文件名称】进行过滤，模糊匹配。类型：String。必选：否&nbsp;yamlId按照【yamlID】进行过滤，模糊匹配。类型：String。必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRecordingRuleYamlTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordingRuleYamlTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetClsServiceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 账户服务开通状态，0:服务已开通，1:服务未开通

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetClsServiceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetClsServiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestUserLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json格式日志内容

		LogInfo *JsonLogInfo `json:"LogInfo,omitempty" name:"LogInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLatestUserLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestUserLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogFastAnalysisResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 字段取值的个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 字段取值的占比信息

		FieldValueRatioInfos []*FieldValueRatioInfos `json:"FieldValueRatioInfos,omitempty" name:"FieldValueRatioInfos"`
		// 日志沉降边界。Unix时间戳（毫秒），在border之前的数据为冷存储，border之后的为标准存储。

		Border *int64 `json:"Border,omitempty" name:"Border"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogFastAnalysisResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogFastAnalysisResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KafkaInfo struct {

	// 可消费topic名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// ACL模式用户名

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// ACL模式密码

	Password *string `json:"Password,omitempty" name:"Password"`
	// 可消费kafka实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// SASL接入点信息

	BootstrapServers *string `json:"BootstrapServers,omitempty" name:"BootstrapServers"`
}

type ModifyIdleResourcePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyIdleResourcePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIdleResourcePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMetricConfigRequest struct {
	*tchttp.BaseRequest

	// 日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 指标采集配置id

	ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
	// 采集配置来源。支持&nbsp;：`0`、`1`&nbsp;-&nbsp;0:自建k8s&nbsp;-&nbsp;1:TKE

	Source *uint64 `json:"Source,omitempty" name:"Source"`
	// 机器组id。

	GroupIds []*string `json:"GroupIds,omitempty" name:"GroupIds"`
	// 操作状态,0:应用,1:暂停

	Operate *uint64 `json:"Operate,omitempty" name:"Operate"`
	// 采集对象,&nbsp;Flag=0时生效

	Spec *MetricSpec `json:"Spec,omitempty" name:"Spec"`
	// 标签处理,&nbsp;Flag=0时生效

	Relabels []*Relabeling `json:"Relabels,omitempty" name:"Relabels"`
	// 标签处理,&nbsp;Flag=0时生效

	MetricRelabels []*Relabeling `json:"MetricRelabels,omitempty" name:"MetricRelabels"`
	// 自定义元数据,&nbsp;Flag=0时生效

	MetricLabel *MetricConfigLabel `json:"MetricLabel,omitempty" name:"MetricLabel"`
	// 通信协议&nbsp;`http`、`https`；Flag=0时生效

	Scheme *string `json:"Scheme,omitempty" name:"Scheme"`
	// 采集频率,&nbsp;Flag=0时生效

	ScrapeInterval *string `json:"ScrapeInterval,omitempty" name:"ScrapeInterval"`
	// 采集超时时间。

	ScrapeTimeout *string `json:"ScrapeTimeout,omitempty" name:"ScrapeTimeout"`
	// Prometheus如何处理标签之间的冲突。当Flag=0&nbsp;&&&nbsp;Type=1时生效，支持`true`,`false`&nbsp;-&nbsp;`false`:配置数据中冲突的标签重命名&nbsp;-&nbsp;`true`:忽略冲突的服务器端标签

	HonorLabels *bool `json:"HonorLabels,omitempty" name:"HonorLabels"`
	// 压缩,&nbsp;false:关闭,&nbsp;true:开启,&nbsp;Flag=0时生效

	EnableCompression *bool `json:"EnableCompression,omitempty" name:"EnableCompression"`
	// 采集配置yaml格式字符串,&nbsp;Flag=1时必填

	YamlSpec *MetricYamlSpec `json:"YamlSpec,omitempty" name:"YamlSpec"`
}

func (r *ModifyMetricConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMetricConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmShieldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 屏蔽规则ID。

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAlarmShieldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmShieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBinlogSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBinlogSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBinlogSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetricSeriesRequest struct {
	*tchttp.BaseRequest

	// 时序主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 起始时间

	Start *uint64 `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *uint64 `json:"End,omitempty" name:"End"`
	// Label匹配规则

	Match []*string `json:"Match,omitempty" name:"Match"`
}

func (r *GetMetricSeriesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricSeriesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CompressInfo struct {

	// 压缩格式，支持gzip、lzop、snappy和none不压缩

	Format *string `json:"Format,omitempty" name:"Format"`
}

type CreateAlarmShieldRequest struct {
	*tchttp.BaseRequest

	// 通知渠道组id。

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
	// 屏蔽开始时间（秒级时间戳）。

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 屏蔽结束时间（秒级时间戳）。

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 屏蔽类型。1：屏蔽所有通知，2：按照Rule参数屏蔽匹配规则的通知。

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 屏蔽规则，当Type为2时必填。规则填写方式详见产品文档。

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 屏蔽原因。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
}

func (r *CreateAlarmShieldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmShieldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志集ID

		LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIndexsRequest struct {
	*tchttp.BaseRequest

	// topicName按照【日志主题名称】进行过滤。类型：String必选：否&nbsp;topicId按照【日志主题ID】进行过滤。类型：String必选：否&nbsp;logsetId按照【日志集ID】进行过滤，可通过调用DescribeLogsets查询已创建的日志集列表或登录控制台进行查看；也可以调用CreateLogset创建新的日志集。类型：String必选：否&nbsp;tagKey按照【标签键】进行过滤。类型：String必选：否&nbsp;tag:tagKey按照【标签键值对】进行过滤。tag-key使用具体的标签键进行替换。使用请参考示例2。类型：String必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeIndexsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIndexsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FilterRuleInfo struct {

	// 过滤规则Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 过滤规则

	Regex *string `json:"Regex,omitempty" name:"Regex"`
	// 过滤规则Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ResourceInfo struct {

	// 需要检索的topic_id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type AddMachineGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组类型&nbsp;目前type支持&nbsp;ip&nbsp;和&nbsp;label

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
}

func (r *AddMachineGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddMachineGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicExtendConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicExtendConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicExtendConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicInfo struct {

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 主题分区个数

	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// 主题是否开启索引（主题类型需为日志主题）。true表示开启索引，false表示关闭索引。

	Index *bool `json:"Index,omitempty" name:"Index"`
	// AssumerUin非空则表示创建该日志主题的服务方Uin

	AssumerUin *uint64 `json:"AssumerUin,omitempty" name:"AssumerUin"`
	// 云产品标识，主题由其它云产品创建时，该字段会显示云产品名称，例如CDN、TKE

	AssumerName *string `json:"AssumerName,omitempty" name:"AssumerName"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 若AssumerUin非空，则表示除服务方外其余调用者修改日志主题的权限

	TopicModifyAcl *int64 `json:"TopicModifyAcl,omitempty" name:"TopicModifyAcl"`
	// 若AssumerUin非空，则表示除服务方外其余调用者展示日志主题的权限

	TopicShowAcl *int64 `json:"TopicShowAcl,omitempty" name:"TopicShowAcl"`
	// 主题是否开启采集，true：开启采集；false：关闭采集。&nbsp;创建日志主题时默认开启，可通过SDK调用ModifyTopic修改此字段。&nbsp;控制台目前不支持修改此参数。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 日志主题绑定的标签信息

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// RoleName非空则表示创建该日志主题的服务方使用的角色

	RoleName *string `json:"RoleName,omitempty" name:"RoleName"`
	// 该主题是否开启自动分裂

	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`
	// 若开启自动分裂的话，该主题能够允许的最大分区数

	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`
	// 主题的存储类型&nbsp;&nbsp;hot:&nbsp;标准存储&nbsp;cold:&nbsp;低频存储

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 生命周期，单位天，可取值范围1~3600。取值为3640时代表永久保存

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 云产品二级标识，日志主题由其它云产品创建时，该字段会显示云产品名称及其日志类型的二级分类，例如TKE-Audit、TKE-Event。部分云产品仅有云产品标识(AssumerName)，无该字段。

	SubAssumerName *string `json:"SubAssumerName,omitempty" name:"SubAssumerName"`
	// 主题对应的日志集信息

	LogsetInfo *LogsetInfo `json:"LogsetInfo,omitempty" name:"LogsetInfo"`
	// 日志主题描述

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 子用户。

	SubUin *uint64 `json:"SubUin,omitempty" name:"SubUin"`
	// 用户自定义抽样配置

	UserSample *string `json:"UserSample,omitempty" name:"UserSample"`
	// 用户采样率状态。true表示开启用户采样率，false表示关闭用户采样率。

	UserSampleStatus *bool `json:"UserSampleStatus,omitempty" name:"UserSampleStatus"`
	// 日志主题名名单信息。

	WhitelistInfo *TopicWhitelistInfo `json:"WhitelistInfo,omitempty" name:"WhitelistInfo"`
	// 开启日志沉降，标准存储的生命周期，&nbsp;hotPeriod&nbsp;<&nbsp;Period。&nbsp;标准存储为&nbsp;hotPeriod,&nbsp;低频存储则为&nbsp;Period-hotPeriod。（主题类型需为日志主题）&nbsp;HotPeriod=0为没有开启日志沉降。

	HotPeriod *uint64 `json:"HotPeriod,omitempty" name:"HotPeriod"`
	// kms-cls服务秘钥id

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 主题类型。&nbsp;&nbsp;0:&nbsp;日志主题&nbsp;1:&nbsp;指标主题

	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`
	// 免鉴权开关。&nbsp;false：关闭；&nbsp;true：开启。&nbsp;开启后将支持指定操作匿名访问该日志主题。详情请参见日志主题。

	IsWebTracking *bool `json:"IsWebTracking,omitempty" name:"IsWebTracking"`
	// 日志主题扩展信息

	Extends *TopicExtendInfo `json:"Extends,omitempty" name:"Extends"`
	// 专用集群id

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 异步迁移任务ID

	TopicAsyncTaskID *string `json:"TopicAsyncTaskID,omitempty" name:"TopicAsyncTaskID"`
	// 异步迁移状态

	MigrationStatus *uint64 `json:"MigrationStatus,omitempty" name:"MigrationStatus"`
	// 异步迁移完成后，预计生效日期

	EffectiveDate *string `json:"EffectiveDate,omitempty" name:"EffectiveDate"`
	// 异步迁移完成后，新的热存储周期

	NewHotPeriod *uint64 `json:"NewHotPeriod,omitempty" name:"NewHotPeriod"`
	// 异步迁移完成后，新的存储周期

	NewPeriod *uint64 `json:"NewPeriod,omitempty" name:"NewPeriod"`
	// 异步任务迁移完成后，新的存储类型

	NewStorageType *string `json:"NewStorageType,omitempty" name:"NewStorageType"`
}

type DescribeDataTransformPreviewDataInfoRequest struct {
	*tchttp.BaseRequest

	// 任务id，&nbsp;获取加工后数据有效

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 1:&nbsp;获取原始数据。&nbsp;&nbsp;&nbsp;2:&nbsp;获取加工后数据

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// 源日志主题id,&nbsp;&nbsp;获取加工前数据有效

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 数据加工类型。0：标准加工任务；&nbsp;1：前置加工任务。默认为0

	DataTransformType *uint64 `json:"DataTransformType,omitempty" name:"DataTransformType"`
}

func (r *DescribeDataTransformPreviewDataInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformPreviewDataInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosRechargeInfo struct {

	// COS导入配置ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// cos导入任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// cos存储桶

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// cos存储桶地域

	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
	// cos存储桶前缀地址

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表单行全文；&nbsp;默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 状态&nbsp;status&nbsp;0:&nbsp;已创建,&nbsp;1:&nbsp;运行中,&nbsp;2:&nbsp;已停止,&nbsp;3:&nbsp;已完成,&nbsp;4:&nbsp;运行失败。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 是否启用:&nbsp;0：&nbsp;未启用&nbsp;，&nbsp;1：启用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 进度条百分值

	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`
	// 压缩格式，支持&nbsp;"",&nbsp;"gzip",&nbsp;"lzop",&nbsp;"snappy”;&nbsp;默认空

	Compress *string `json:"Compress,omitempty" name:"Compress"`
	// 见：&nbsp;ExtractRuleInfo&nbsp;结构描述

	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitempty" name:"ExtractRuleInfo"`
	// COS导入任务类型。1：一次性导入任务；2：持续性导入任务。

	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
	// 元数据。支持&nbsp;bucket，object。

	Metadata []*string `json:"Metadata,omitempty" name:"Metadata"`
}

type PreviewKafkaRechargeRequest struct {
	*tchttp.BaseRequest

	// 预览类型，1：源数据预览；2：导出结果预览。

	PreviewType *uint64 `json:"PreviewType,omitempty" name:"PreviewType"`
	// 导入Kafka类型，0:&nbsp;CKafka，1:&nbsp;用户自建Kafka

	KafkaType *uint64 `json:"KafkaType,omitempty" name:"KafkaType"`
	// CKafka实例ID，KafkaType为0时必填

	KafkaInstance *string `json:"KafkaInstance,omitempty" name:"KafkaInstance"`
	// 服务地址。&nbsp;KafkaType为1时ServerAddr必填。

	ServerAddr *string `json:"ServerAddr,omitempty" name:"ServerAddr"`
	// ServerAddr是否为加密连接。&nbsp;KafkaType为1时有效。

	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitempty" name:"IsEncryptionAddr"`
	// 加密访问协议。&nbsp;KafkaType为1并且IsEncryptionAddr为true时Protocol必填。

	Protocol *KafkaProtocolInfo `json:"Protocol,omitempty" name:"Protocol"`
	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开。&nbsp;最多支持100个。

	UserKafkaTopics *string `json:"UserKafkaTopics,omitempty" name:"UserKafkaTopics"`
	// 用户Kafka消费组

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 导入数据位置，-2：最早；-1：最晚。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 日志导入规则

	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitempty" name:"LogRechargeRule"`
	// 网络连接参数

	NetworkInfo *NetworkInfo `json:"NetworkInfo,omitempty" name:"NetworkInfo"`
}

func (r *PreviewKafkaRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PreviewKafkaRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PolicyInfo struct {

	// 按位取值标记操作类型。&nbsp;-&nbsp;0b0001&nbsp;标识主动重启命令&nbsp;-&nbsp;0b0010&nbsp;标识强制更新采集配置命令

	OP *uint64 `json:"OP,omitempty" name:"OP"`
}

type CheckRemoteWriteTaskConnectRequest struct {
	*tchttp.BaseRequest

	// 日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 私有网络id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 目标服务名称

	Target *string `json:"Target,omitempty" name:"Target"`
	// 目标地址

	RemoteWriteURL *string `json:"RemoteWriteURL,omitempty" name:"RemoteWriteURL"`
	// 鉴权类型&nbsp;0:&nbsp;无鉴权&nbsp;1:&nbsp;basic_auth&nbsp;2:&nbsp;token

	AuthType *uint64 `json:"AuthType,omitempty" name:"AuthType"`
	// 鉴权信息

	AuthInfo *RemoteWriteAuthInfo `json:"AuthInfo,omitempty" name:"AuthInfo"`
	// 网络类型&nbsp;1&nbsp;内网&nbsp;2外网

	NetType *uint64 `json:"NetType,omitempty" name:"NetType"`
	// 子网

	SubNet *string `json:"SubNet,omitempty" name:"SubNet"`
	// 后端服务类型&nbsp;0&nbsp;CVM&nbsp;1025&nbsp;CLB

	VirtualGatewayType *int64 `json:"VirtualGatewayType,omitempty" name:"VirtualGatewayType"`
}

func (r *CheckRemoteWriteTaskConnectRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckRemoteWriteTaskConnectRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExportsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志导出列表

		Exports []*ExportInfo `json:"Exports,omitempty" name:"Exports"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExportsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExportsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeWebCallbacksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警渠道回调配置列表。

		WebCallbacks []*WebCallbackInfo `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
		// 符合条件的通知内容配置总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWebCallbacksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeWebCallbacksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Column struct {

	// 列的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 列的属性

	Type *string `json:"Type,omitempty" name:"Type"`
}

type CreateRecordingRuleTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecordingRuleTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRecordingRuleTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScheduledSqlResouceInfo struct {

	// 目标主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 主账号Uin

	Uin *int64 `json:"Uin,omitempty" name:"Uin"`
	// 主题的地域信息

	Region *string `json:"Region,omitempty" name:"Region"`
	// 主题类型：0为日志主题，1为指标主题

	BizType *int64 `json:"BizType,omitempty" name:"BizType"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标名称&nbsp;BizType为1时，优先使用MetricNames字段多指标只能填充到MetricNames字段，单指标建议填充到MetricName字段

	MetricNames []*string `json:"MetricNames,omitempty" name:"MetricNames"`
	// 指标维度，不接受时间类型。

	MetricLabels []*string `json:"MetricLabels,omitempty" name:"MetricLabels"`
	// 指标时间戳，默认值为SQL查询时间范围的左侧时间点，您也可以指定其他字段（类型为uinx时间、TimeStamp，精度毫秒）为指标时间戳。

	CustomTime *string `json:"CustomTime,omitempty" name:"CustomTime"`
	// 除了MetricLabels，您还可以使用该参数，为指标补充静态的维度。&nbsp;维度名以字母或下划线开头，后面可以跟字母、数字或下划线，长度小于等于1024&nbsp;字节

	CustomMetricLabels []*MetricLabel `json:"CustomMetricLabels,omitempty" name:"CustomMetricLabels"`
}

type CreateConfigExtraResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集配置扩展信息ID

		ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigExtraResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigExtraResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QcloudGoodsDetailInfo struct {

	// 价格模型

	Pid *uint64 `json:"Pid,omitempty" name:"Pid"`
	// 商品的时间单位：y:年；m:月；d:日；h:小时；M:分钟；s:秒;&nbsp;p:与计费周期无关,一次性购买的产品传p

	TimeUnit *string `json:"TimeUnit,omitempty" name:"TimeUnit"`
	// 商品的时间大小（一次性售卖固定传1）

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 放入goodsDetail内:子产品标签

	SubProductCode *string `json:"SubProductCode,omitempty" name:"SubProductCode"`
	// 新增参数:&nbsp;产品标签

	ProductCode *string `json:"ProductCode,omitempty" name:"ProductCode"`
	// 相关计费项个数

	SvClsPartitionCount *uint64 `json:"SvClsPartitionCount,omitempty" name:"SvClsPartitionCount"`
}

type CreateExportRequest struct {
	*tchttp.BaseRequest

	// 日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志导出检索语句，不支持[SQL语句]

	Query *string `json:"Query,omitempty" name:"Query"`
	// 日志导出数量,&nbsp;&nbsp;最大值1000万

	Count *uint64 `json:"Count,omitempty" name:"Count"`
	// 日志导出时间排序。desc，asc，默认为desc

	Order *string `json:"Order,omitempty" name:"Order"`
	// 日志导出数据格式。json，csv，默认为json

	Format *string `json:"Format,omitempty" name:"Format"`
	// 日志导出起始时间，毫秒时间戳

	From *int64 `json:"From,omitempty" name:"From"`
	// 日志导出结束时间，毫秒时间戳

	To *int64 `json:"To,omitempty" name:"To"`
	// 语法规则,&nbsp;默认值为0。&nbsp;0：Lucene语法，1：CQL语法。

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 导出字段

	DerivedFields []*string `json:"DerivedFields,omitempty" name:"DerivedFields"`
	// 分隔符。CSV&nbsp;文件中各字段间的分隔符，非必填，默认&nbsp;逗号&nbsp;支持：&nbsp;&nbsp;空格&nbsp;\t：制表符&nbsp;,：逗号&nbsp;;：竖线

	Separator *string `json:"Separator,omitempty" name:"Separator"`
	// 转义符。CSV&nbsp;文件字段值中出现了分隔符的字符，需用转义符包裹该字符，防止读取数据时被错误识别，非必填，默认&nbsp;双引号&nbsp;支持：&nbsp;&nbsp;'&nbsp;：单引号&nbsp;"：双引号&nbsp;空串

	EscapeCharacter *string `json:"EscapeCharacter,omitempty" name:"EscapeCharacter"`
	// 填充字段。csv文件字段不存在(无效)时，使用用户指定的值进行填充。&nbsp;默认空&nbsp;支持：&nbsp;&nbsp;空字符串&nbsp;满足正则表达式&nbsp;^[A-Za-z0-9_=+\-!$&*()%#@~^]{0,255}$&nbsp;的字符串

	FillField *string `json:"FillField,omitempty" name:"FillField"`
	// 首行Key是否展示。默认&nbsp;true&nbsp;支持：&nbsp;&nbsp;true：展示&nbsp;false：不展示

	DisplayHeader *bool `json:"DisplayHeader,omitempty" name:"DisplayHeader"`
}

func (r *CreateExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaInfo struct {

	// Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value

	Value *bool `json:"Value,omitempty" name:"Value"`
}

type DescribeRecordingRuleTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RecordingRule任务列表信息

		RecordingRuleTaskInfos []*RecordingRuleTaskInfo `json:"RecordingRuleTaskInfos,omitempty" name:"RecordingRuleTaskInfos"`
		// 任务总次数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRecordingRuleTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordingRuleTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicBaseMetricConfigsRequest struct {
	*tchttp.BaseRequest

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// groupId按照【机器组id】进行过滤。类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，所有Filter.Values总和上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeTopicBaseMetricConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicBaseMetricConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyShipperRequest struct {
	*tchttp.BaseRequest

	// 投递规则ID

	ShipperId *string `json:"ShipperId,omitempty" name:"ShipperId"`
	// COS存储桶，详见产品支持的存储桶命名规范。

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// 投递规则投递的新的目录前缀。&nbsp;&nbsp;仅支持0-9A-Za-z-_/&nbsp;最大支持256个字符

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 投递规则的开关状态。true：开启投递任务；false：关闭投递任务。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 投递规则的名字

	ShipperName *string `json:"ShipperName,omitempty" name:"ShipperName"`
	// 投递的时间间隔，单位&nbsp;秒，默认300，范围&nbsp;300-900

	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`
	// 投递的文件的最大值，单位&nbsp;MB，默认100，范围&nbsp;5-256

	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`
	// 投递日志的过滤规则，匹配的日志进行投递，各rule之间是and关系，最多5个，数组为空则表示不过滤而全部投递

	FilterRules []*FilterRuleInfo `json:"FilterRules,omitempty" name:"FilterRules"`
	// 投递日志的分区规则，支持strftime的时间格式表示

	Partition *string `json:"Partition,omitempty" name:"Partition"`
	// 投递日志的压缩配置

	Compress *CompressInfo `json:"Compress,omitempty" name:"Compress"`
	// 投递日志的内容格式配置

	Content *ContentInfo `json:"Content,omitempty" name:"Content"`
	// 投递文件命名配置，0：随机数命名，1：投递时间命名。

	FilenameMode *uint64 `json:"FilenameMode,omitempty" name:"FilenameMode"`
	// 跨账号uin，用于支持跨账号bucket投递

	CustomUin *uint64 `json:"CustomUin,omitempty" name:"CustomUin"`
	// cos桶存储类型。支持：STANDARD_IA、ARCHIVE、DEEP_ARCHIVE、STANDARD、MAZ_STANDARD、MAZ_STANDARD_IA、INTELLIGENT_TIERING。&nbsp;&nbsp;STANDARD_IA：低频存储；&nbsp;ARCHIVE：归档存储；&nbsp;DEEP_ARCHIVE：深度归档存储；&nbsp;STANDARD：标准存储；&nbsp;MAZ_STANDARD：标准存储（多&nbsp;AZ）；&nbsp;MAZ_STANDARD_IA：低频存储（多&nbsp;AZ）；&nbsp;INTELLIGENT_TIERING：智能分层存储。

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 角色访问描述名

	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`
	// 外部ID

	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
}

func (r *ModifyShipperRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyShipperRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 待消费kafka信息

		Kafka *KafkaInfo `json:"Kafka,omitempty" name:"Kafka"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpenKafkaConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKafkaConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeIdleResourcePolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 自动冻结；0关闭&nbsp;1开启

		AutoFreeze *uint64 `json:"AutoFreeze,omitempty" name:"AutoFreeze"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeIdleResourcePolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeIdleResourcePolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集规则扩展配置ID

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
}

func (r *DeleteConfigExtraRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigExtraRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScheduledSqlProcessInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// ScheduledSQL任务进度信息列表

		ScheduledSqlTaskProcessInfos []*ScheduledSqlTaskProcessInfo `json:"ScheduledSqlTaskProcessInfos,omitempty" name:"ScheduledSqlTaskProcessInfos"`
		// 任务总次数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 执行成功数

		TotalSuccess *uint64 `json:"TotalSuccess,omitempty" name:"TotalSuccess"`
		// 执行失败数

		TotalFailed *uint64 `json:"TotalFailed,omitempty" name:"TotalFailed"`
		// 执行运行中

		TotalRunning *uint64 `json:"TotalRunning,omitempty" name:"TotalRunning"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScheduledSqlProcessInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScheduledSqlProcessInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryMetricRequest struct {
	*tchttp.BaseRequest

	// Prometheus语法规则查询语句

	Query *string `json:"Query,omitempty" name:"Query"`
	// 时间戳

	Time *uint64 `json:"Time,omitempty" name:"Time"`
	// 时序主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *QueryMetricRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryMetricRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendDetailItem struct {

	// Uin或者Gin

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// Uin或者Group或者Http或者Wecom

	Type *string `json:"Type,omitempty" name:"Type"`
	// 成功数

	SuccessCount *int64 `json:"SuccessCount,omitempty" name:"SuccessCount"`
}

type DescribeDataTransformInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 数据加工任务列表信息

		DataTransformTaskInfos []*DataTransformTaskInfo `json:"DataTransformTaskInfos,omitempty" name:"DataTransformTaskInfos"`
		// 任务总次数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadLogRequest struct {
	*tchttp.BaseRequest

	// 该参数已废弃，请勿使用

	HashKey *string `json:"HashKey,omitempty" name:"HashKey"`
	// 主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 压缩方法

	CompressType *string `json:"CompressType,omitempty" name:"CompressType"`
	// Agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// Agent请求的序列号

	AgentSeq *string `json:"AgentSeq,omitempty" name:"AgentSeq"`
	// Agent版本号

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// Agent请求的Unique&nbsp;Id

	UniqueId *string `json:"UniqueId,omitempty" name:"UniqueId"`
}

func (r *UploadLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBinlogSubscribeRequest struct {
	*tchttp.BaseRequest

	// binlog采集任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// binlog采集任务配置的日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteBinlogSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBinlogSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteBinlogSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteBinlogSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteBinlogSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformPreviewInfoRequest struct {
	*tchttp.BaseRequest

	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDataTransformPreviewInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformPreviewInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogContextResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志上下文信息集合

		LogContextInfos []*LogContextInfo `json:"LogContextInfos,omitempty" name:"LogContextInfos"`
		// 上文日志是否已经返回完成（当PrevOver为false，表示有上文日志还未全部返回）。

		PrevOver *bool `json:"PrevOver,omitempty" name:"PrevOver"`
		// 下文日志是否已经返回完成（当NextOver为false，表示有下文日志还未全部返回）。

		NextOver *bool `json:"NextOver,omitempty" name:"NextOver"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogContextResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogContextResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformInfoRequest struct {
	*tchttp.BaseRequest

	// taskName&nbsp;按照【加工任务名称】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;taskId&nbsp;按照【加工任务id】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;topicId&nbsp;按照【源topicId】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;status&nbsp;按照【&nbsp;任务运行状态】进行过滤。&nbsp;1：准备中，2：运行中，3：停止中，4：已停止&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;hasServiceLog&nbsp;按照【是否开启服务日志】进行过滤。&nbsp;1：未开启，2：已开启&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;dstTopicType&nbsp;按照【目标topic类型】进行过滤。&nbsp;1：固定，2：动态&nbsp;类型：String&nbsp;必选：否&nbsp;&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 默认值为2.&nbsp;&nbsp;&nbsp;1:&nbsp;获取单个任务的详细信息&nbsp;2：获取任务列表

	Type *int64 `json:"Type,omitempty" name:"Type"`
	// Type为1，&nbsp;此参数必填

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeDataTransformInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type FieldInfo struct {

	// 字段名

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 字段类型

	FieldType *string `json:"FieldType,omitempty" name:"FieldType"`
}

type DeleteDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type ModifyRemoteWriteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyRemoteWriteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyRemoteWriteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAlarmRuleRequest struct {
	*tchttp.BaseRequest

	// 触发条件

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为2000。

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 报警周期

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 报警时间

	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`
	// 监控对象

	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`
	// 多触发条件。

	MultiConditions []*MultiCondition `json:"MultiConditions,omitempty" name:"MultiConditions"`
}

func (r *CheckAlarmRuleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAlarmRuleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertHistoryRecord struct {

	// 告警历史ID

	RecordId *string `json:"RecordId,omitempty" name:"RecordId"`
	// 告警策略ID

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
	// 告警策略名称

	AlarmName *string `json:"AlarmName,omitempty" name:"AlarmName"`
	// 监控对象ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 监控对象名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 监控对象所属地域

	Region *string `json:"Region,omitempty" name:"Region"`
	// 触发条件

	Trigger *string `json:"Trigger,omitempty" name:"Trigger"`
	// 持续周期，持续满足触发条件TriggerCount个周期后，再进行告警

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 告警通知发送频率，单位为分钟

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 通知渠道组

	Notices []*AlertHistoryNotice `json:"Notices,omitempty" name:"Notices"`
	// 告警持续时间，单位为分钟

	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
	// 告警状态，0代表未恢复，1代表已恢复，2代表已失效

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 告警发生时间，毫秒级Unix时间戳

	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 告警分组触发时对应的分组信息

	GroupTriggerCondition []*GroupTriggerConditionInfo `json:"GroupTriggerCondition,omitempty" name:"GroupTriggerCondition"`
	// 告警级别，0代表警告(Warn)，1代表提醒(Info)，2代表紧急&nbsp;(Critical)

	AlarmLevel *uint64 `json:"AlarmLevel,omitempty" name:"AlarmLevel"`
	// 监控对象类型。&nbsp;0:执行语句共用监控对象;&nbsp;1:每个执行语句单独选择监控对象。

	MonitorObjectType *uint64 `json:"MonitorObjectType,omitempty" name:"MonitorObjectType"`
	// 告警处理人

	ClaimOperator *string `json:"ClaimOperator,omitempty" name:"ClaimOperator"`
}

type DescribeKafkaConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Kafka协议消费打开状态，true表示消费打开，false表示消费关闭

		Status *bool `json:"Status,omitempty" name:"Status"`
		// KafkaConsumer&nbsp;消费时使用的Topic参数

		TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
		// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]

		Compression *int64 `json:"Compression,omitempty" name:"Compression"`
		// kafka协议消费数据格式

		ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitempty" name:"ConsumerContent"`
		// 是否开启投递服务日志。1：关闭，2：开启。

		HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
		// 协议类型，支持&nbsp;0,1,&nbsp;默认&nbsp;0&nbsp;&nbsp;0:&nbsp;ckafka版本&nbsp;1:&nbsp;历史数据版本

		ProtocolType *uint64 `json:"ProtocolType,omitempty" name:"ProtocolType"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckAlarmChannelResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 测试结果

		ChannelTestResults []*ChannelTestResult `json:"ChannelTestResults,omitempty" name:"ChannelTestResults"`
		// 是否成功进行测试

		ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
		// 错误原因

		ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckAlarmChannelResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAlarmChannelResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDataTransformAutoCreateTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题总数

		TotalTopicCount *uint64 `json:"TotalTopicCount,omitempty" name:"TotalTopicCount"`
		// 日志主题信息列表

		Topics []*TopicETL `json:"Topics,omitempty" name:"Topics"`
		// 重复的日志主题总数

		DuplicateTopicCount *uint64 `json:"DuplicateTopicCount,omitempty" name:"DuplicateTopicCount"`
		// 重复的日志主题信息列表

		DuplicateTopics []*TopicETL `json:"DuplicateTopics,omitempty" name:"DuplicateTopics"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformAutoCreateTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformAutoCreateTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogsetsRequest struct {
	*tchttp.BaseRequest

	// logsetName&nbsp;&nbsp;按照【日志集名称】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;logsetId&nbsp;&nbsp;按照【日志集ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;tagKey&nbsp;&nbsp;按照【标签键】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;tag:tagKey&nbsp;&nbsp;按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。&nbsp;类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeLogsetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogsetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScheduledSqlRequest struct {
	*tchttp.BaseRequest

	// 源日志主题

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务启动状态.&nbsp;1开启,&nbsp;2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 定时SQL分析目标日志主题

	DstResource *ScheduledSqlResouceInfo `json:"DstResource,omitempty" name:"DstResource"`
	// 查询语句

	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitempty" name:"ScheduledSqlContent"`
	// 调度开始时间,Unix时间戳，单位ms

	ProcessStartTime *uint64 `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 调度类型，1:持续运行&nbsp;2:指定时间范围

	ProcessType *int64 `json:"ProcessType,omitempty" name:"ProcessType"`
	// 调度结束时间，当ProcessType=2时为必传字段,&nbsp;Unix时间戳，单位ms

	ProcessEndTime *uint64 `json:"ProcessEndTime,omitempty" name:"ProcessEndTime"`
	// 调度周期(分钟)

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 单次查询的时间窗口

	ProcessTimeWindow *string `json:"ProcessTimeWindow,omitempty" name:"ProcessTimeWindow"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// 源topicId的地域信息

	SrcTopicRegion *string `json:"SrcTopicRegion,omitempty" name:"SrcTopicRegion"`
	// 查询语法规则。&nbsp;默认值为0。0：Lucene语法，1：CQL语法

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *CreateScheduledSqlRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScheduledSqlRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateScheduledSqlResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateScheduledSqlResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateScheduledSqlResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKafkaConsumerRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
	// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]

	Compression *int64 `json:"Compression,omitempty" name:"Compression"`
	// kafka协议消费数据格式

	ConsumerContent *KafkaConsumerContent `json:"ConsumerContent,omitempty" name:"ConsumerContent"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *ModifyKafkaConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKafkaConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MultiCondition struct {

	// 触发条件。

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 交互式触发配置信息。

	ConditionInteractiveConfig *string `json:"ConditionInteractiveConfig,omitempty" name:"ConditionInteractiveConfig"`
	// 告警级别。0:警告(Warn);&nbsp;1:提醒(Info);&nbsp;2:紧急&nbsp;(Critical)。&nbsp;&nbsp;不填则默认为0。

	AlarmLevel *uint64 `json:"AlarmLevel,omitempty" name:"AlarmLevel"`
}

type CreateAlarmNoticeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警模板ID

		AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type RuleKeyValueInfo struct {

	// 是否大小写敏感。true表示大小写敏感，false表示大小写不敏感。

	CaseSensitive *bool `json:"CaseSensitive,omitempty" name:"CaseSensitive"`
	// 需要建立索引的键值对信息；最大只能配置100个键值对

	KeyValues []*KeyValueInfo `json:"KeyValues,omitempty" name:"KeyValues"`
	// 索引是否开启动态模板；若开启，则会根据上报的键值对配置索引，但是所有字段类型都是text，大小写敏感，不支持分析，采用默认分词符。已废弃

	TemplateType *string `json:"TemplateType,omitempty" name:"TemplateType"`
}

type DescribeTopicExtendConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// clb的topic业务配置

		ClbTopicExtendConfigs []*ClbTopicExtendConfig `json:"ClbTopicExtendConfigs,omitempty" name:"ClbTopicExtendConfigs"`
		// 总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicExtendConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTopicExtendConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyIndexRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 索引是否生效，默认不生效。true表示索引生效，false表示索引失效。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 索引规则

	Rule *RuleInfo `json:"Rule,omitempty" name:"Rule"`
	// 内置保留字段（__FILENAME__，__HOSTNAME__及__SOURCE__）是否包含至全文索引，默认为false，推荐设置为true&nbsp;&nbsp;false:不包含&nbsp;true:包含

	IncludeInternalFields *bool `json:"IncludeInternalFields,omitempty" name:"IncludeInternalFields"`
	// 元数据字段（前缀为__TAG__的字段）是否包含至全文索引，默认为0，推荐设置为1&nbsp;&nbsp;0:仅包含开启键值索引的元数据字段&nbsp;1:包含所有元数据字段&nbsp;2:不包含任何元数据字段

	MetadataFlag *uint64 `json:"MetadataFlag,omitempty" name:"MetadataFlag"`
	// 自定义日志解析异常存储字段。

	CoverageField *string `json:"CoverageField,omitempty" name:"CoverageField"`
}

func (r *ModifyIndexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyIndexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMetricConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMetricConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMetricConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAccountInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户类型。

		UinType *string `json:"UinType,omitempty" name:"UinType"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShipperPreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预览返回结果

		PreviewInfos []*string `json:"PreviewInfos,omitempty" name:"PreviewInfos"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeShipperPreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShipperPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRebuildIndexTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 重建起始时间戳，毫秒

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 重建结束时间戳，毫秒

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *CreateRebuildIndexTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRebuildIndexTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRecordingRuleTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// yamlId【关联yaml配置ID】进行过滤，模糊匹配。类型：String。必选：否&nbsp;taskName按照【任务名称】进行过滤，模糊匹配。类型：String。必选：否&nbsp;taskId按照【任务ID】进行过滤，模糊匹配。类型：String。必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeRecordingRuleTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRecordingRuleTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRemoteWriteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRemoteWriteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRemoteWriteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpenKafkaConsumeRequest struct {
	*tchttp.BaseRequest

	// CLS相关TopicId

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *OpenKafkaConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpenKafkaConsumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordingRuleYamlTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Yaml配置id，&nbsp;可以关联多子任务

		YamlId *string `json:"YamlId,omitempty" name:"YamlId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRecordingRuleYamlTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRecordingRuleYamlTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardSubscribesRequest struct {
	*tchttp.BaseRequest

	// dashboardId：按照【仪表盘id】进行过滤。类型：String必选：否&nbsp;&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeDashboardSubscribesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardSubscribesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDashboardsRequest struct {
	*tchttp.BaseRequest

	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// dashboardId&nbsp;按照【仪表盘id】进行过滤，类型：String，&nbsp;必选：否。&nbsp;dashboardName&nbsp;按照【仪表盘名字】进行模糊搜索过滤，类型：String，必选：否。&nbsp;dashboardRegion&nbsp;按照【仪表盘地域】进行过滤，为了兼容老的仪表盘，通过云API创建的仪表盘没有地域属性，类型：String，必选：否。&nbsp;tagKey&nbsp;按照【标签键】进行过滤，类型：String，必选：否。&nbsp;tag:tagKey&nbsp;按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换，类型：String，必选：否，使用请参考示例2。&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 按照topicId和regionId过滤。

	TopicIdRegionFilter []*TopicIdAndRegion `json:"TopicIdRegionFilter,omitempty" name:"TopicIdRegionFilter"`
}

func (r *DescribeDashboardsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDashboardsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MergePartitionRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 合并的PartitionId（找到下一个分区InclusiveBeginKey与入参PartitionId对应的ExclusiveEndKey相等，且找到的分区必须是读写分区（Staus:readwrite），入参PartitionId与找到的PartitionId设置为只读分区（Status:readonly）,再新建一个新的读写分区）&nbsp;。获取分区列表&nbsp;&nbsp;入参PartitionId只能是读写分区（Status的值有readonly，readwrite），且能找到入参PartitionId的下一个可读写分区（找到下一个分区InclusiveBeginKey与入参PartitionId对应的ExclusiveEndKey相等）；&nbsp;入参PartitionId不能是最后一个分区（PartitionId的ExclusiveEndKey不能是ffffffffffffffffffffffffffffffff）；&nbsp;topic的分区数量是有限制的（默认50个），合并之后不能超过最大分区，否则不能合并。

	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
}

func (r *MergePartitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *MergePartitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlertHistoryNotice struct {

	// 通知渠道组名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 通知渠道组ID

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
}

type ScheduledSqlTaskProcessInfo struct {

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 实例ID

	ProcessId *string `json:"ProcessId,omitempty" name:"ProcessId"`
	// 加工语句

	ScheduledSqlContent *string `json:"ScheduledSqlContent,omitempty" name:"ScheduledSqlContent"`
	// 执行时间-开始时间

	ProcessStartTime *string `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 执行时间-结束时间

	ProcessEndTime *string `json:"ProcessEndTime,omitempty" name:"ProcessEndTime"`
	// 执行时间-耗时

	ProcessDuration *int64 `json:"ProcessDuration,omitempty" name:"ProcessDuration"`
	// SQL时间窗口-开始时间

	TimeWindowStartTime *string `json:"TimeWindowStartTime,omitempty" name:"TimeWindowStartTime"`
	// SQL时间窗口-结束时间

	TimeWindowEndTime *string `json:"TimeWindowEndTime,omitempty" name:"TimeWindowEndTime"`
	// 处理数据量-输入行数

	ReadLogCount *uint64 `json:"ReadLogCount,omitempty" name:"ReadLogCount"`
	// 处理数据量-输出行数

	WriteLogCount *uint64 `json:"WriteLogCount,omitempty" name:"WriteLogCount"`
	// 调度结果，1:运行中&nbsp;2:成功&nbsp;3:失败

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 源topicid

	SrcTopicId *string `json:"SrcTopicId,omitempty" name:"SrcTopicId"`
	// 失败原因字段

	StatusFailedMsg *string `json:"StatusFailedMsg,omitempty" name:"StatusFailedMsg"`
	// 目标日志主题所在地域信息。

	Region *string `json:"Region,omitempty" name:"Region"`
	// 目标日志主题id。

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
}

type SearchDashboardSubscribeRequest struct {
	*tchttp.BaseRequest

	// 仪表盘订阅Id。

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// 仪表盘id。

	DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
	// 仪表盘订阅名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 仪表盘订阅数据。

	SubscribeData *DashboardSubscribeData `json:"SubscribeData,omitempty" name:"SubscribeData"`
}

func (r *SearchDashboardSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchDashboardSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAgentNormalRequest struct {
	*tchttp.BaseRequest

	// 需要升级的机器IP列表

	Agents []*AgentKeyInfo `json:"Agents,omitempty" name:"Agents"`
	// 升级类型:0-disable,1-manual,2-auto

	UpdateMode *int64 `json:"UpdateMode,omitempty" name:"UpdateMode"`
	// 升级开始时间，如：22:00:00，晚上10点开始

	UpdateStart *string `json:"UpdateStart,omitempty" name:"UpdateStart"`
	// 升级结束时间，如：23:00:00，晚上11点结束

	UpdateStop *string `json:"UpdateStop,omitempty" name:"UpdateStop"`
	// 升级目标版本

	TargetVersion *string `json:"TargetVersion,omitempty" name:"TargetVersion"`
}

func (r *UpgradeAgentNormalRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeAgentNormalRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmShieldCount struct {

	// 符合检索条件的告警屏蔽总数量

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 告警屏蔽未生效数量

	InvalidCount *uint64 `json:"InvalidCount,omitempty" name:"InvalidCount"`
	// 告警屏蔽生效中数量

	ValidCount *uint64 `json:"ValidCount,omitempty" name:"ValidCount"`
	// 告警屏蔽已过期数量

	ExpireCount *uint64 `json:"ExpireCount,omitempty" name:"ExpireCount"`
}

type ExternalDataSourceInfo struct {

	// 外部数据源ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 数据源类型：1表示Mysql数据源，2表示cos数据源

	Datasource *uint64 `json:"Datasource,omitempty" name:"Datasource"`
	// 隶属日志主题ID

	SubjectionTopicId *string `json:"SubjectionTopicId,omitempty" name:"SubjectionTopicId"`
	// 隶属日志主题名称

	SubjectionTopicName *string `json:"SubjectionTopicName,omitempty" name:"SubjectionTopicName"`
	// 当前请求日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// mysql配置信息

	SQLInfo *ExternalDataSourceSQLInfo `json:"SQLInfo,omitempty" name:"SQLInfo"`
	// cos配置信息

	COSInfo *ExternalDataSourceCOSInfo `json:"COSInfo,omitempty" name:"COSInfo"`
	// 创建时间（秒级时间戳）

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间（秒级时间戳）

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type DeleteConfigurationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConfigurationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConfigurationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNoticeContentsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 通知内容模板列表。

		NoticeContents []*NoticeContentTemplate `json:"NoticeContents,omitempty" name:"NoticeContents"`
		// 符合条件的通知内容模板总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNoticeContentsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNoticeContentsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelRebuildIndexTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CancelRebuildIndexTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelRebuildIndexTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckConfigRegexResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 0&nbsp;成功&nbsp;1匹配失败&nbsp;2参数异常

		Status *uint64 `json:"Status,omitempty" name:"Status"`
		// 返回的消息内容

		Message *string `json:"Message,omitempty" name:"Message"`
		// 匹配到的日志内容

		Matches []*string `json:"Matches,omitempty" name:"Matches"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckConfigRegexResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckConfigRegexResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteExportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CloseKafkaConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CloseKafkaConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CloseKafkaConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateDashboardResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘id

		DashboardId *string `json:"DashboardId,omitempty" name:"DashboardId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type DescribeDataTransformPreviewDataInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预览数据详细信息

		PreviewLogStatistics []*PreviewLogStatistic `json:"PreviewLogStatistics,omitempty" name:"PreviewLogStatistics"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDataTransformPreviewDataInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDataTransformPreviewDataInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaConsumeRequest struct {
	*tchttp.BaseRequest

	// CLS对应topic标识

	FromTopicId *string `json:"FromTopicId,omitempty" name:"FromTopicId"`
}

func (r *DescribeKafkaConsumeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMetricCorrectDimensionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标校准信息

		Datas []*NamespaceCorrectDimension `json:"Datas,omitempty" name:"Datas"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMetricCorrectDimensionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMetricCorrectDimensionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteUserConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteUserConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteUserConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScheduledSqlInfoRequest struct {
	*tchttp.BaseRequest

	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// srcTopicName按照【源日志主题名称】进行过滤，模糊匹配。类型：String。必选：否&nbsp;dstTopicName按照【目标日志主题名称】进行过滤，模糊匹配。类型：String。必选：否&nbsp;srcTopicId按照【源日志主题ID】进行过滤。类型：String。必选：否&nbsp;dstTopicId按照【目标日志主题ID】进行过滤。类型：String。必选：否&nbsp;bizType按照【主题类型】进行过滤，0：日志主题；1：指标主题。类型：String。必选：否&nbsp;status按照【任务状态】进行过滤，1：运行；2：停止。类型：String。必选：否&nbsp;taskName按照【任务名称】进行过滤，模糊匹配。类型：String。必选：否&nbsp;taskId按照【任务ID】进行过滤，模糊匹配。类型：String。必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeScheduledSqlInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeScheduledSqlInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetMetricLabelValuesRequest struct {
	*tchttp.BaseRequest

	// 时序主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 时序label名称

	LabelName *string `json:"LabelName,omitempty" name:"LabelName"`
	// 起始时间

	Start *uint64 `json:"Start,omitempty" name:"Start"`
	// 结束时间

	End *uint64 `json:"End,omitempty" name:"End"`
	// Label匹配规则

	Match []*string `json:"Match,omitempty" name:"Match"`
}

func (r *GetMetricLabelValuesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetMetricLabelValuesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NetworkInfo struct {

	// 网络类型。&nbsp;0：公网，1：内网

	NetworkType *uint64 `json:"NetworkType,omitempty" name:"NetworkType"`
	// 私有网络id

	VpcID *string `json:"VpcID,omitempty" name:"VpcID"`
	// 私有网络所属用户app&nbsp;id

	AppID *uint64 `json:"AppID,omitempty" name:"AppID"`
	// 网络服务类型。0：CVM，3：专线网关，11：云联网，1025：CLB

	VirtualGatewayType *uint64 `json:"VirtualGatewayType,omitempty" name:"VirtualGatewayType"`
	// 专线网关id或者云联网id

	VpcGatewayIndex *string `json:"VpcGatewayIndex,omitempty" name:"VpcGatewayIndex"`
	// 私有域名映射地址

	PrivateDomainNames []*PrivateDomainNames `json:"PrivateDomainNames,omitempty" name:"PrivateDomainNames"`
}

type CheckUserIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户ID是否符合规范，1:符合，-1:&nbsp;ID已经存在，-2:&nbsp;输入ID无效

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckUserIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckUserIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateRecordingRuleYamlTaskRequest struct {
	*tchttp.BaseRequest

	// 任务数据源日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 任务写入目标主题

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
	// 任务状态；&nbsp;1:开启；2:关闭

	EnableFlag *uint64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 调度开始时间,Unix时间戳，单位ms

	ProcessStartTime *uint64 `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 调度周期(分钟)，支持范围(0,1440]分钟。

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// yaml配置名称

	YamlConfigName *string `json:"YamlConfigName,omitempty" name:"YamlConfigName"`
	// yaml配置内容

	YamlContent *string `json:"YamlContent,omitempty" name:"YamlContent"`
	// 执行周期单位,&nbsp;`0`:&nbsp;Minute,&nbsp;`1`:&nbsp;Second。&nbsp;默认：0&nbsp;-&nbsp;ProcessPeriodUnit:1时，ProcessPeriod&nbsp;只支持&nbsp;30（即只支持30秒）。&nbsp;-&nbsp;ProcessPeriodUnit:0时，ProcessPeriod&nbsp;支持范围(0,1440]分钟。

	ProcessPeriodUnit *uint64 `json:"ProcessPeriodUnit,omitempty" name:"ProcessPeriodUnit"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
}

func (r *CreateRecordingRuleYamlTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRecordingRuleYamlTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmsRequest struct {
	*tchttp.BaseRequest

	// name&nbsp;&nbsp;按照【告警策略名称】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;alarmId&nbsp;&nbsp;按照【告警策略ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;topicId&nbsp;&nbsp;按照【监控对象的日志主题ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;enable&nbsp;&nbsp;按照【启用状态】进行过滤。&nbsp;类型：String&nbsp;备注：enable参数值范围:&nbsp;1,&nbsp;t,&nbsp;T,&nbsp;TRUE,&nbsp;true,&nbsp;True,&nbsp;0,&nbsp;f,&nbsp;F,&nbsp;FALSE,&nbsp;false,&nbsp;False。&nbsp;其它值将返回参数错误信息.&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlarmsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordingRuleYamlTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteRecordingRuleYamlTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecordingRuleYamlTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineGroupsRequest struct {
	*tchttp.BaseRequest

	// machineGroupName&nbsp;&nbsp;按照【机器组名称】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;machineGroupId&nbsp;&nbsp;按照【机器组ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;osType&nbsp;&nbsp;按照【操作系统类型】进行过滤。&nbsp;类型：Int&nbsp;必选：否&nbsp;tagKey&nbsp;&nbsp;按照【标签键】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;tag:tagKey&nbsp;&nbsp;按照【标签键值对】进行过滤。tagKey使用具体的标签键进行替换。&nbsp;类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMachineGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type KafkaConsumerInfo struct {

	// 消费主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// kafka的主题id

	ToTopicId *string `json:"ToTopicId,omitempty" name:"ToTopicId"`
	// kafka实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 消费组信息

	ConsumerGroupList []*ConsumerGroupInfoForMonitor `json:"ConsumerGroupList,omitempty" name:"ConsumerGroupList"`
}

type SearchLogRequest struct {
	*tchttp.BaseRequest

	// 要检索分析的日志主题ID，仅能指定一个日志主题。&nbsp;如需同时检索多个日志主题，请使用Topics参数。&nbsp;TopicId&nbsp;和&nbsp;Topics&nbsp;不能同时使用，在一次请求中有且只能选择一个。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 要检索分析的日志的起始时间，Unix时间戳（毫秒）

	From *int64 `json:"From,omitempty" name:"From"`
	// 要检索分析的日志的结束时间，Unix时间戳（毫秒）

	To *int64 `json:"To,omitempty" name:"To"`
	// 检索分析语句，最大长度为12KB&nbsp;语句由&nbsp;[检索条件]&nbsp;|&nbsp;[SQL语句]构成，无需对日志进行统计分析时，可省略其中的管道符&nbsp;|&nbsp;及SQL语句&nbsp;使用*或空字符串可查询所有日志

	Query *string `json:"Query,omitempty" name:"Query"`
	// 表示单次查询返回的原始日志条数，默认为100，最大值为1000。&nbsp;注意：&nbsp;&nbsp;仅当检索分析语句(Query)不包含SQL时有效&nbsp;SQL结果条数指定方式参考SQL&nbsp;LIMIT语法&nbsp;可通过两种方式获取后续更多日志：&nbsp;&nbsp;Context:透传上次接口返回的Context值，获取后续更多日志，总计最多可获取1万条原始日志&nbsp;Offset:偏移量，表示从第几行开始返回原始日志，无日志条数限制

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 透传上次接口返回的Context值，可获取后续更多日志，总计最多可获取1万条原始日志，过期时间1小时。&nbsp;注意：&nbsp;&nbsp;透传该参数时，请勿修改除该参数外的其它参数&nbsp;仅适用于单日志主题检索，检索多个日志主题时，请使用Topics中的Context&nbsp;仅当检索分析语句(Query)不包含SQL时有效，SQL获取后续结果参考SQL&nbsp;LIMIT语法

	Context *string `json:"Context,omitempty" name:"Context"`
	// 原始日志是否按时间排序返回；可选值：asc(升序)、desc(降序)，默认为&nbsp;desc&nbsp;注意：&nbsp;&nbsp;仅当检索分析语句(Query)不包含SQL时有效&nbsp;SQL结果排序方式参考SQL&nbsp;ORDER&nbsp;BY语法

	Sort *string `json:"Sort,omitempty" name:"Sort"`
	// 是否返回符合检索条件的关键词，一般用于高亮显示匹配的关键词，仅支持键值检索

	HighLight *bool `json:"HighLight,omitempty" name:"HighLight"`
	// 为true代表使用新的检索结果返回方式，输出参数AnalysisRecords和Columns有效&nbsp;为false时代表使用老的检索结果返回方式,&nbsp;输出AnalysisResults和ColNames有效&nbsp;两种返回方式在编码格式上有少量区别，建议使用true

	UseNewAnalysis *bool `json:"UseNewAnalysis,omitempty" name:"UseNewAnalysis"`
	// 0：不执行语法优化；1：执行语法优化

	QueryOptimize *uint64 `json:"QueryOptimize,omitempty" name:"QueryOptimize"`
	// 执行统计分析（Query中包含SQL）时，是否对原始日志先进行采样，再进行统计分析。&nbsp;0：自动采样;&nbsp;0～1：按指定采样率采样，例如0.02;&nbsp;1：不采样，即精确分析&nbsp;默认值为1

	SamplingRate *float64 `json:"SamplingRate,omitempty" name:"SamplingRate"`
	// 检索语法规则，默认值为0，推荐使用1&nbsp;。&nbsp;&nbsp;0：Lucene语法&nbsp;1：CQL语法（日志服务专用检索语法，控制台默认也使用该语法规则）。&nbsp;详细说明参见检索条件语法规则

	SyntaxRule *uint64 `json:"SyntaxRule,omitempty" name:"SyntaxRule"`
	// 要检索分析的日志主题列表，最大支持50个日志主题。&nbsp;检索单个日志主题时请使用TopicId。&nbsp;TopicId&nbsp;和&nbsp;Topics&nbsp;不能同时使用，在一次请求中有且只能选择一个。

	Topics []*MultiTopicSearchInformation `json:"Topics,omitempty" name:"Topics"`
	// 查询原始日志的偏移量，表示从第几行开始返回原始日志，默认为0。&nbsp;注意：&nbsp;&nbsp;仅当检索分析语句(Query)不包含SQL时有效&nbsp;不能与Context参数同时使用&nbsp;仅适用于单日志主题检索

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *SearchLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SearchLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Label struct {

	// 标签的键。有效标签键有两个部分：可选前缀和名称，以斜杠&nbsp;(/)&nbsp;分隔。名称部分是必需的，并且必须不超过&nbsp;63&nbsp;个字符，以字母数字字符&nbsp;([a-z0-9A-Z])&nbsp;开头和结尾，中间有破折号(-)、下划线(_)、点(.)&nbsp;和字母数字。前缀是可选的。如果指定，前缀必须是&nbsp;DNS&nbsp;子域：一系列以点&nbsp;(.)&nbsp;分隔的&nbsp;DNS&nbsp;标签，总长度不超过&nbsp;253&nbsp;个字符，后跟斜杠&nbsp;(&nbsp;/)。&nbsp;-&nbsp;prefix&nbsp;格式&nbsp;`[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*`&nbsp;-&nbsp;name&nbsp;格式&nbsp;`([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]`&nbsp;-&nbsp;key不能重复

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签键值直接的比较关系。&nbsp;不同业务场景支持的比较符不同，具体支持那些参考接口业务描述。&nbsp;例如：`in`、`notin`

	Operate *string `json:"Operate,omitempty" name:"Operate"`
	// 标签的值.&nbsp;-&nbsp;最大支持63个字符。&nbsp;-&nbsp;格式：`([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]`

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type ParquetInfo struct {

	// ParquetKeyInfo数组

	ParquetKeyInfo []*ParquetKeyInfo `json:"ParquetKeyInfo,omitempty" name:"ParquetKeyInfo"`
}

type DeleteDataTransformRequest struct {
	*tchttp.BaseRequest

	// 数据加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DeleteDataTransformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteDataTransformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlertRecordHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警历史总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 告警历史详情

		Records []*AlertHistoryRecord `json:"Records,omitempty" name:"Records"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAlertRecordHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertRecordHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyFolderRequest struct {
	*tchttp.BaseRequest

	// 文件夹ID。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 文件夹名称。&nbsp;&nbsp;必填项。&nbsp;名称限制不能超过50个字符。&nbsp;重名报错。

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *ModifyFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterBaseMetricConfigsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总数目

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 指标采集配置列表

		Datas []*BaseMetricCollectConfig `json:"Datas,omitempty" name:"Datas"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterBaseMetricConfigsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterBaseMetricConfigsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogsetRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集名称

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志集的绑定的标签键值对。最大支持10个标签键值对，同一个资源只能同时绑定一个标签键。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 生命周期，单位为天；可取值范围1~366

	Period *int64 `json:"Period,omitempty" name:"Period"`
}

func (r *ModifyLogsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AgentUpdateStatus struct {

	// 升级类型:0-null,1-manual,2-auto,3-force

	UpdateType *int64 `json:"UpdateType,omitempty" name:"UpdateType"`
	// 升级动作:0-null,1-update,2-revert

	UpdateAction *int64 `json:"UpdateAction,omitempty" name:"UpdateAction"`
	// 重试次数,最大三次

	RetryCount *int64 `json:"RetryCount,omitempty" name:"RetryCount"`
	// Agent升级状态:0-querying,1-updating,2-reverting,-1-updatefail,-2-revertfail,&nbsp;-10-notsupport

	UpdateStatus *int64 `json:"UpdateStatus,omitempty" name:"UpdateStatus"`
	// 错误码

	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`
	// 错误信息

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type CollectConfig struct {

	// 指定采集类型的采集配置名称信息。&nbsp;&nbsp;当CollectInfo中Type为0：表示元数据配置，name为元数据名称。&nbsp;目前支持"container_id"，"container_name"，"image_name"，"namespace"，"pod_uid"，"pod_name"，"pod_ip"。&nbsp;当CollectInfo中Type为1：指定pod&nbsp;label，name为指定pod&nbsp;label名称。

	Name *string `json:"Name,omitempty" name:"Name"`
}

type LogRechargeRuleInfo struct {

	// 导入类型，支持json_log：json格式日志，minimalist_log:&nbsp;单行全文，fullregex_log:&nbsp;单行完全正则

	RechargeType *string `json:"RechargeType,omitempty" name:"RechargeType"`
	// 整条日志匹配规则，只有RechargeType为fullregex_log时有效

	LogRegex *string `json:"LogRegex,omitempty" name:"LogRegex"`
	// 解析失败日志是否上传，true表示上传，false表示不上传

	UnMatchLogSwitch *bool `json:"UnMatchLogSwitch,omitempty" name:"UnMatchLogSwitch"`
	// 解析失败日志的键名称

	UnMatchLogKey *string `json:"UnMatchLogKey,omitempty" name:"UnMatchLogKey"`
	// 解析失败日志时间来源，0:&nbsp;系统当前时间，1:&nbsp;Kafka消息时间戳

	UnMatchLogTimeSrc *uint64 `json:"UnMatchLogTimeSrc,omitempty" name:"UnMatchLogTimeSrc"`
	// 解析编码格式，0:&nbsp;UTF-8（默认值），1:&nbsp;GBK

	EncodingFormat *uint64 `json:"EncodingFormat,omitempty" name:"EncodingFormat"`
	// 使用默认时间，true：开启（默认值），&nbsp;flase：关闭

	DefaultTimeSwitch *bool `json:"DefaultTimeSwitch,omitempty" name:"DefaultTimeSwitch"`
	// 默认时间来源，0:&nbsp;系统当前时间，1:&nbsp;Kafka消息时间戳

	DefaultTimeSrc *uint64 `json:"DefaultTimeSrc,omitempty" name:"DefaultTimeSrc"`
	// 时间字段

	TimeKey *string `json:"TimeKey,omitempty" name:"TimeKey"`
	// 时间提取正则表达式

	TimeRegex *string `json:"TimeRegex,omitempty" name:"TimeRegex"`
	// 时间字段格式

	TimeFormat *string `json:"TimeFormat,omitempty" name:"TimeFormat"`
	// 时间字段时区

	TimeZone *string `json:"TimeZone,omitempty" name:"TimeZone"`
	// 元数据信息，Kafka导入支持kafka_topic,kafka_partition,kafka_offset,kafka_timestamp

	Metadata []*string `json:"Metadata,omitempty" name:"Metadata"`
	// 日志Key列表，RechargeType为full_regex_log时必填

	Keys []*string `json:"Keys,omitempty" name:"Keys"`
	// 日志过滤规则列表

	AdvanceFilterRules []*AdvanceFilterRuleInfo `json:"AdvanceFilterRules,omitempty" name:"AdvanceFilterRules"`
	// json解析模式，开启首层数据解析

	ParseArray *bool `json:"ParseArray,omitempty" name:"ParseArray"`
}

type DescribeRemoteWriteTasksResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RemoteWrite&nbsp;信息列表

		Infos []*RemoteWriteInfo `json:"Infos,omitempty" name:"Infos"`
		// RemoteWrite信息总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteWriteTasksResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteWriteTasksResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFolderRequest struct {
	*tchttp.BaseRequest

	// 文件夹ID。

	Id *string `json:"Id,omitempty" name:"Id"`
}

func (r *DeleteFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMetricSubscribeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMetricSubscribeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMetricSubscribeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateMetricConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 指标采集配置id列表。

		ConfigIds []*string `json:"ConfigIds,omitempty" name:"ConfigIds"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMetricConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMetricConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyConfigExtraRequest struct {
	*tchttp.BaseRequest

	// 采集配置扩展信息id

	ConfigExtraId *string `json:"ConfigExtraId,omitempty" name:"ConfigExtraId"`
	// 采集配置规程名称，最长63个字符，只能包含小写字符、数字及分隔符（“-”），且必须以小写字符开头，数字或小写字符结尾

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 节点文件配置信息

	HostFile *HostFileInfo `json:"HostFile,omitempty" name:"HostFile"`
	// 容器文件路径信息

	ContainerFile *ContainerFileInfo `json:"ContainerFile,omitempty" name:"ContainerFile"`
	// 容器标准输出信息

	ContainerStdout *ContainerStdoutInfo `json:"ContainerStdout,omitempty" name:"ContainerStdout"`
	// 采集的日志类型，默认为minimalist_log。支持以下类型：&nbsp;&nbsp;json_log代表：JSON-文件日志（详见使用&nbsp;JSON&nbsp;提取模式采集日志）；&nbsp;delimiter_log代表：分隔符-文件日志（详见使用分隔符提取模式采集日志）；&nbsp;minimalist_log代表：单行全文-文件日志（详见使用单行全文提取模式采集日志）；&nbsp;fullregex_log代表：单行完全正则-文件日志（详见使用单行-完全正则提取模式采集日志）；&nbsp;multiline_log代表：多行全文-文件日志（详见使用多行全文提取模式采集日志）；&nbsp;multiline_fullregex_log代表：多行完全正则-文件日志（详见使用多行-完全正则提取模式采集日志）；&nbsp;user_define_log代表：组合解析（适用于多格式嵌套的日志，详见使用组合解析提取模式采集日志）。

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 日志格式化方式，用于容器采集场景。目前已经废弃&nbsp;&nbsp;stdout-docker-json：用于docker容器采集场景&nbsp;stdout-containerd：用于containerd容器采集场景

	LogFormat *string `json:"LogFormat,omitempty" name:"LogFormat"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRule *ExtractRuleInfo `json:"ExtractRule,omitempty" name:"ExtractRule"`
	// 采集黑名单路径列表

	ExcludePaths []*ExcludePathInfo `json:"ExcludePaths,omitempty" name:"ExcludePaths"`
	// 组合解析采集规则，用于复杂场景下的日志采集。&nbsp;&nbsp;取值参考：使用组合解析提取模式采集日志

	UserDefineRule *string `json:"UserDefineRule,omitempty" name:"UserDefineRule"`
	// 类型：container_stdout、container_file、host_file

	Type *string `json:"Type,omitempty" name:"Type"`
	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 自建采集配置标

	ConfigFlag *string `json:"ConfigFlag,omitempty" name:"ConfigFlag"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集name

	LogsetName *string `json:"LogsetName,omitempty" name:"LogsetName"`
	// 日志主题name

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 采集相关配置信息。详情见CollectInfo复杂类型配置。

	CollectInfos []*CollectInfo `json:"CollectInfos,omitempty" name:"CollectInfos"`
	// 高级采集配置。&nbsp;Json字符串，&nbsp;Key/Value定义为如下：&nbsp;&nbsp;ClsAgentFileTimeout(超时属性),&nbsp;取值范围:&nbsp;大于等于0的整数，&nbsp;0为不超时&nbsp;ClsAgentMaxDepth(最大目录深度)，取值范围:&nbsp;大于等于0的整数&nbsp;ClsAgentParseFailMerge(合并解析失败日志)，取值范围:&nbsp;true或false&nbsp;ClsAgentDefault(自定义默认值，无特殊含义，用于清空其他选项)，建议取值0

	AdvancedConfig *string `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
}

func (r *ModifyConfigExtraRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyConfigExtraRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConfigurationTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置模板id

		TemplateId *string `json:"TemplateId,omitempty" name:"TemplateId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigurationTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigurationTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataTransformResouceInfo struct {

	// 目标主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 日志集id

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 主账号Uin

	Uin *string `json:"Uin,omitempty" name:"Uin"`
}

type DescribeMachinesRequest struct {
	*tchttp.BaseRequest

	// 查询的机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// ip&nbsp;&nbsp;按照【ip】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;instance&nbsp;&nbsp;按照【instance】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;version&nbsp;&nbsp;按照【LogListener版本】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;status&nbsp;&nbsp;按照【状态】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;可选值：0：离线，1：正常&nbsp;offlineTime&nbsp;&nbsp;按照【机器离线时间】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;可选值：0：无离线时间，12：12小时内，24：一天内，48：两天内，99：两天前&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目。最大支持100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmShieldResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmShieldResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmShieldResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordingRuleYamlTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	YamlId *string `json:"YamlId,omitempty" name:"YamlId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteRecordingRuleYamlTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecordingRuleYamlTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChannelTestResult struct {

	// 序号

	Index *int64 `json:"Index,omitempty" name:"Index"`
	// 错误码，0是正确，-1001表示无效请求。

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
	// 发送结果

	SendTotal *SendDetail `json:"SendTotal,omitempty" name:"SendTotal"`
}

type DescribeBinlogSubscribesRequest struct {
	*tchttp.BaseRequest

	// 日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// taskId按照【配置id】进行过滤。类型：String&nbsp;必选：否&nbsp;&nbsp;name按照【配置名称】进行过滤。类型：String&nbsp;必选：否&nbsp;&nbsp;statusFlag按照【配置状态标记】进行过滤。类型：String&nbsp;必选：否&nbsp;&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeBinlogSubscribesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBinlogSubscribesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineGroupMetadataInfo struct {

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组元数据列表

	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`
}

type PreviewLogStatistic struct {

	// 目标日志主题

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
	// 日志内容

	LogContent *string `json:"LogContent,omitempty" name:"LogContent"`
	// 失败错误信息，&nbsp;空字符串""表示正常

	FailReason *string `json:"FailReason,omitempty" name:"FailReason"`
	// 行号。从0开始

	LineNum *int64 `json:"LineNum,omitempty" name:"LineNum"`
	// 日志时间，格式：2024-05-07&nbsp;17:13:17.105&nbsp;&nbsp;入参时无效&nbsp;出参时有效，为日志中的时间格式

	Time *string `json:"Time,omitempty" name:"Time"`
	// 目标topic-name

	DstTopicName *string `json:"DstTopicName,omitempty" name:"DstTopicName"`
	// 目标logsetname

	DstLogSetName *string `json:"DstLogSetName,omitempty" name:"DstLogSetName"`
}

type CreateExternalDataSourceRequest struct {
	*tchttp.BaseRequest

	// 名称。仅支持小写字母、数字和_，且不能以_开头和结尾，长度为3至60字符，账号内本地域唯一，不支持修改，在CLS&nbsp;SQL中当作表名。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 备注，可选参数，长度不超过255个字符

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 数据源类型，&nbsp;1:MySQL;&nbsp;2:COS(csv格式)

	Datasource *uint64 `json:"Datasource,omitempty" name:"Datasource"`
	// SQL相关配置信息，当Datasource值为1时，此字段必填

	SQLInfo *ExternalDataSourceSQLInfo `json:"SQLInfo,omitempty" name:"SQLInfo"`
	// COS相关配置信息，当Datasource值为2时，此字段必填

	COSInfo *ExternalDataSourceCOSInfo `json:"COSInfo,omitempty" name:"COSInfo"`
	// 日志主题ID,&nbsp;此外部数据源配置隶属此日志主题

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集ID,&nbsp;不为空时表示该日志集下面的所有日志主题都可以使用此外部数据源配置

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

func (r *CreateExternalDataSourceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateExternalDataSourceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKafkaRechargeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyKafkaRechargeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKafkaRechargeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomLabel struct {

	// 标签的键。&nbsp;-&nbsp;必须以字母或下划线开头，但不可以双下划线（__）开头，后面可以跟任意字母，数字或下划线。&nbsp;-&nbsp;最大支持256个字符。&nbsp;-&nbsp;key不能重复

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签的值。&nbsp;-&nbsp;最大支持256个字符。

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeMachineGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机器组信息列表

		MachineGroups []*MachineGroupInfo `json:"MachineGroups,omitempty" name:"MachineGroups"`
		// 分页的总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteConsumerGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteConsumerGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteConsumerGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SplitPartitionRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 待分裂分区ID

	PartitionId *int64 `json:"PartitionId,omitempty" name:"PartitionId"`
	// 分区切分的哈希key的位置，只在Number=2时有意义

	SplitKey *string `json:"SplitKey,omitempty" name:"SplitKey"`
	// 分区分裂个数(可选)，默认等于2

	Number *int64 `json:"Number,omitempty" name:"Number"`
}

func (r *SplitPartitionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SplitPartitionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerOffsetsRequest struct {
	*tchttp.BaseRequest

	// 日志主题对应的消费组标识

	ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
	// 时间戳(秒级时间戳)

	From *string `json:"From,omitempty" name:"From"`
	// 日志集id(日志主题对应的id)

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志集id(日志主题对应的id)

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 分区id

	PartitionId *string `json:"PartitionId,omitempty" name:"PartitionId"`
}

func (r *DescribeConsumerOffsetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerOffsetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmShieldRequest struct {
	*tchttp.BaseRequest

	// 屏蔽规则ID。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 通知渠道组id。

	AlarmNoticeId *string `json:"AlarmNoticeId,omitempty" name:"AlarmNoticeId"`
	// 屏蔽开始时间（秒级时间戳）。

	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`
	// 屏蔽结束时间（秒级时间戳）。

	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
	// 屏蔽类型。1：屏蔽所有通知，2：按照Rule参数屏蔽匹配规则的通知。

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 屏蔽规则，当Type为2时必填。

	Rule *string `json:"Rule,omitempty" name:"Rule"`
	// 屏蔽原因。

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 规则状态。只有规则状态为生效中（status:1）时，才能将其修改为已失效（status:2）。

	Status *uint64 `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAlarmShieldRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmShieldRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDataTransformRequest struct {
	*tchttp.BaseRequest

	// 加工任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 加工任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 加工语句。&nbsp;当FuncType为2时，EtlContent必须使用log_auto_output&nbsp;&nbsp;其他参考文档：&nbsp;&nbsp;创建加工任务&nbsp;函数总览

	EtlContent *string `json:"EtlContent,omitempty" name:"EtlContent"`
	// 任务启动状态.&nbsp;默认为1，开启,&nbsp;2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 加工任务目的topic_id以及别名

	DstResources []*DataTransformResouceInfo `json:"DstResources,omitempty" name:"DstResources"`
	// 超限之后是否丢弃日志数据

	BackupGiveUpData *bool `json:"BackupGiveUpData,omitempty" name:"BackupGiveUpData"`
	// 交互式加工语句配置内容

	InteractiveEtlContent *string `json:"InteractiveEtlContent,omitempty" name:"InteractiveEtlContent"`
	// 是否开启投递服务日志。1关闭，2开启

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
	// 保留失败日志状态。&nbsp;1:不保留)，2:保留

	KeepFailureLog *uint64 `json:"KeepFailureLog,omitempty" name:"KeepFailureLog"`
	// 失败日志的字段名称

	FailureLogKey *string `json:"FailureLogKey,omitempty" name:"FailureLogKey"`
}

func (r *ModifyDataTransformRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDataTransformRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateConsumerRequest struct {
	*tchttp.BaseRequest

	// 投递任务绑定的日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 是否投递日志的元数据信息，默认为&nbsp;true。&nbsp;当NeedContent为true时：字段Content有效。&nbsp;当NeedContent为false时：字段Content无效。

	NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`
	// 如果需要投递元数据信息，元数据信息的描述

	Content *ConsumerContent `json:"Content,omitempty" name:"Content"`
	// CKafka的描述

	Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`
	// 投递时压缩方式，取值0，2，3。[0：NONE；2：SNAPPY；3：LZ4]

	Compression *int64 `json:"Compression,omitempty" name:"Compression"`
	// 角色访问描述名

	RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`
	// 外部ID

	ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
	// 高级配置项

	AdvancedConfig *AdvancedConsumerConfiguration `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
}

func (r *CreateConsumerRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConsumerRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNoticeContentRequest struct {
	*tchttp.BaseRequest

	// 模板名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 模板内容语言。0：zh1：英文

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 模板详细配置。

	NoticeContents []*NoticeContent `json:"NoticeContents,omitempty" name:"NoticeContents"`
}

func (r *CreateNoticeContentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNoticeContentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLatestJsonLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// json格式日志内容

		LogData *JsonLogInfo `json:"LogData,omitempty" name:"LogData"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLatestJsonLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLatestJsonLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MetaTagInfo struct {

	// 元数据key

	Key *string `json:"Key,omitempty" name:"Key"`
	// 元数据value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeAccountRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HeartBeatRequest struct {
	*tchttp.BaseRequest

	// agent的版本号

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// agent的IP地址

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// agent采集的数据大小

	TotalSize *uint64 `json:"TotalSize,omitempty" name:"TotalSize"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// agent序列号

	AgentSeq *string `json:"AgentSeq,omitempty" name:"AgentSeq"`
	// 标签列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
	// Agent是否开启自动升级功能

	AutoUpdate *string `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 是否需要检查自动升级任务

	CheckUpdate *string `json:"CheckUpdate,omitempty" name:"CheckUpdate"`
}

func (r *HeartBeatRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *HeartBeatRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RemoteWriteAuthInfo struct {

	// basic&nbsp;auth&nbsp;username

	Username *string `json:"Username,omitempty" name:"Username"`
	// basic&nbsp;auth&nbsp;password

	Password *string `json:"Password,omitempty" name:"Password"`
	// basic&nbsp;auth&nbsp;token

	Token *string `json:"Token,omitempty" name:"Token"`
}

type ModifyAlarmRequest struct {
	*tchttp.BaseRequest

	// 告警策略ID。

	AlarmId *string `json:"AlarmId,omitempty" name:"AlarmId"`
	// 告警策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 监控任务运行时间点。

	MonitorTime *MonitorTime `json:"MonitorTime,omitempty" name:"MonitorTime"`
	// 触发条件。&nbsp;&nbsp;注意:&nbsp;&nbsp;Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。

	Condition *string `json:"Condition,omitempty" name:"Condition"`
	// 持续周期。持续满足触发条件TriggerCount个周期后，再进行告警；最小值为1，最大值为2000。

	TriggerCount *int64 `json:"TriggerCount,omitempty" name:"TriggerCount"`
	// 告警重复的周期。单位是分钟。取值范围是0~1440。

	AlarmPeriod *int64 `json:"AlarmPeriod,omitempty" name:"AlarmPeriod"`
	// 关联的告警通知模板列表。

	AlarmNoticeIds []*string `json:"AlarmNoticeIds,omitempty" name:"AlarmNoticeIds"`
	// 监控对象列表。

	AlarmTargets []*AlarmTarget `json:"AlarmTargets,omitempty" name:"AlarmTargets"`
	// 是否开启告警策略。true表示开启告警策略，false表示关闭告警策略。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 该参数已废弃，请使用Status参数控制是否开启告警策略。

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 用户自定义告警内容

	MessageTemplate *string `json:"MessageTemplate,omitempty" name:"MessageTemplate"`
	// 用户自定义回调

	CallBack *CallBackInfo `json:"CallBack,omitempty" name:"CallBack"`
	// 多维分析

	Analysis []*AnalysisDimensional `json:"Analysis,omitempty" name:"Analysis"`
	// 分组触发状态。true：开启，false：关闭（默认）

	GroupTriggerStatus *bool `json:"GroupTriggerStatus,omitempty" name:"GroupTriggerStatus"`
	// 分组触发条件。

	GroupTriggerCondition []*string `json:"GroupTriggerCondition,omitempty" name:"GroupTriggerCondition"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的告警策略。最大支持10个标签键值对，并且不能有重复的键值对。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 监控对象类型。0:执行语句共用监控对象;&nbsp;1:每个执行语句单独选择监控对象。&nbsp;当值为1时，AlarmTargets元素个数不能超过10个，AlarmTargets中的Number必须是从1开始的连续正整数，不能重复。

	MonitorObjectType *uint64 `json:"MonitorObjectType,omitempty" name:"MonitorObjectType"`
	// 告警附加分类信息列表。&nbsp;Classifications元素个数不能超过20个。&nbsp;Classifications元素的Key不能为空，不能重复，长度不能超过50个字符，符合正则&nbsp;^[a-z]([a-z0-9_]{0,49})$。&nbsp;Classifications元素的Value长度不能超过200个字符。

	Classifications []*AlarmClassification `json:"Classifications,omitempty" name:"Classifications"`
	// 告警级别。&nbsp;&nbsp;0:警告(Warn);1:提醒(Info);2:紧急&nbsp;(Critical)&nbsp;&nbsp;注意:&nbsp;&nbsp;Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。

	AlarmLevel *uint64 `json:"AlarmLevel,omitempty" name:"AlarmLevel"`
	// 触发条件交互模式配置信息。

	ConditionInteractiveConfig *string `json:"ConditionInteractiveConfig,omitempty" name:"ConditionInteractiveConfig"`
	// 多触发条件。&nbsp;&nbsp;注意:&nbsp;&nbsp;Condition和AlarmLevel是一组配置，MultiConditions是另一组配置，2组配置互斥。

	MultiConditions []*MultiCondition `json:"MultiConditions,omitempty" name:"MultiConditions"`
	// 告警模板相关信息

	AlarmTemplateInfo *AlarmTemplateConfig `json:"AlarmTemplateInfo,omitempty" name:"AlarmTemplateInfo"`
}

func (r *ModifyAlarmRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BinlogFilterList struct {

	// 数据库名称。

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	// 数据表名称。

	TableNames []*string `json:"TableNames,omitempty" name:"TableNames"`
}

type CheckAlarmChannelRequest struct {
	*tchttp.BaseRequest

	// 告警通知接收者数组

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 回调数组

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 通知类型（Trigger,Recovery,All）

	Type *string `json:"Type,omitempty" name:"Type"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckAlarmChannelRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckAlarmChannelRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBinlogSubscribeConnectivityRequest struct {
	*tchttp.BaseRequest

	// 数据库类型。&nbsp;1:云mysql&nbsp;2:TDSQL-C&nbsp;3:自建mysql

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// 数据库配置信息。

	SqlInfo *BinlogDBConfig `json:"SqlInfo,omitempty" name:"SqlInfo"`
	// 任务id。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeBinlogSubscribeConnectivityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBinlogSubscribeConnectivityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerPreviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 预览数据结果

		PreviewInfos []*string `json:"PreviewInfos,omitempty" name:"PreviewInfos"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerPreviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerPreviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetConfigurationTemplateApplyLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 加载后续详情的Context

		Context *string `json:"Context,omitempty" name:"Context"`
		// 指定时间范围内的配置模板应用记录是否完整返回

		ListOver *bool `json:"ListOver,omitempty" name:"ListOver"`
		// 返回的结果是否为SQL分析结果

		Analysis *bool `json:"Analysis,omitempty" name:"Analysis"`
		// 如果Analysis为True，则返回分析结果的列名，否则为空

		ColNames []*string `json:"ColNames,omitempty" name:"ColNames"`
		// 执行详情查询结果；当Analysis为True时，可能返回为null

		Results []*LogInfo `json:"Results,omitempty" name:"Results"`
		// 执行详情统计分析结果；当Analysis为False时，可能返回为null

		AnalysisResults []*LogItems `json:"AnalysisResults,omitempty" name:"AnalysisResults"`
		// 执行详情统计分析结果;&nbsp;UseNewAnalysis为true有效

		AnalysisRecords []*string `json:"AnalysisRecords,omitempty" name:"AnalysisRecords"`
		// 分析结果的列名，&nbsp;UseNewAnalysis为true有效

		Columns []*Column `json:"Columns,omitempty" name:"Columns"`
		// 返回语法优化后的语句（入参如有QueryOptimize&nbsp;且为&nbsp;1&nbsp;时返回，其他情况返回空字符串）

		Query *string `json:"Query,omitempty" name:"Query"`
		// 当前系统使用的采样率，入参如有SamplingRate时生效（主要是当客户输入0时，返回真实后台的AutoSamplingRate值）

		SamplingRate *float64 `json:"SamplingRate,omitempty" name:"SamplingRate"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetConfigurationTemplateApplyLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetConfigurationTemplateApplyLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContentInfo struct {

	// 内容格式，支持json、csv

	Format *string `json:"Format,omitempty" name:"Format"`
	// csv格式内容描述

	Csv *CsvInfo `json:"Csv,omitempty" name:"Csv"`
	// json格式内容描述

	Json *JsonInfo `json:"Json,omitempty" name:"Json"`
	// parquet格式内容描述

	Parquet *ParquetInfo `json:"Parquet,omitempty" name:"Parquet"`
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，并且不能有重复的键值对。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 主题是否开启采集，true：开启采集；false：关闭采集。&nbsp;控制台目前不支持修改此参数。

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 是否开启自动分裂。true表示开启自动分裂，false表示关闭自动分裂。

	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`
	// 若开启最大分裂，该主题能够能够允许的最大分区数

	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`
	// 生命周期，单位天；可取值范围1~90

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 用户自定义抽样配置

	UserSample *string `json:"UserSample,omitempty" name:"UserSample"`
	// 用户抽样配置开关：默认关闭。true表示开启用户抽样配置，false表示关闭用户抽样配置。

	UserSampleStatus *bool `json:"UserSampleStatus,omitempty" name:"UserSampleStatus"`
	// 存储类型:hot&nbsp;标准存储

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 日志主题描述

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 0：关闭日志沉降。&nbsp;非0：开启日志沉降后标准存储的天数。HotPeriod需要大于等于7，且小于Period。仅在StorageType为&nbsp;hot&nbsp;时生效

	HotPeriod *uint64 `json:"HotPeriod,omitempty" name:"HotPeriod"`
	// 免鉴权开关。&nbsp;false：关闭；&nbsp;true：开启。&nbsp;开启后将支持指定操作匿名访问该日志主题。详情请参见日志主题。

	IsWebTracking *bool `json:"IsWebTracking,omitempty" name:"IsWebTracking"`
	// 隔离topic，true：隔离，false，取消隔离

	Isolated *bool `json:"Isolated,omitempty" name:"Isolated"`
	// 日志主题扩展信息

	Extends *TopicExtendInfo `json:"Extends,omitempty" name:"Extends"`
	// 日志主题分区数量

	PartitionCount *uint64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// 取消切换存储任务的id

	CancelTopicAsyncTaskID *string `json:"CancelTopicAsyncTaskID,omitempty" name:"CancelTopicAsyncTaskID"`
	// 加密相关参数。&nbsp;支持加密地域并且开白用户可以传此参数，其他场景不能传递该参数。&nbsp;只支持传入1：kms-cls&nbsp;云产品秘钥加密

	Encryption *uint64 `json:"Encryption,omitempty" name:"Encryption"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CancelRebuildIndexTaskRequest struct {
	*tchttp.BaseRequest

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 索引重建任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *CancelRebuildIndexTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CancelRebuildIndexTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteExportRequest struct {
	*tchttp.BaseRequest

	// 日志导出ID

	ExportId *string `json:"ExportId,omitempty" name:"ExportId"`
}

func (r *DeleteExportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteExportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RecordingRuleTaskInfo struct {

	// 预聚合任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 源日志主题id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 预聚合任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 任务创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 任务更新时间

	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
	// 任务状态，1:运行&nbsp;2:停止&nbsp;3:异常-找不到源日志主题&nbsp;4:异常-找不到目标主题&nbsp;5:&nbsp;访问权限问题&nbsp;6:内部故障&nbsp;7:其他故障

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 任务启用状态，1开启,&nbsp;2关闭

	EnableFlag *int64 `json:"EnableFlag,omitempty" name:"EnableFlag"`
	// 调度开始时间

	ProcessStartTime *uint64 `json:"ProcessStartTime,omitempty" name:"ProcessStartTime"`
	// 调度周期(分钟)

	ProcessPeriod *int64 `json:"ProcessPeriod,omitempty" name:"ProcessPeriod"`
	// 执行周期单位,&nbsp;`0`:&nbsp;Minute,&nbsp;`1`:&nbsp;Second。&nbsp;-&nbsp;ProcessPeriodUnit:1时，ProcessPeriod&nbsp;只支持&nbsp;30（即只支持30秒）。&nbsp;-&nbsp;ProcessPeriodUnit:0时，ProcessPeriod&nbsp;支持范围(0,1440]分钟。

	ProcessPeriodUnit *uint64 `json:"ProcessPeriodUnit,omitempty" name:"ProcessPeriodUnit"`
	// 执行延迟(秒)

	ProcessDelay *int64 `json:"ProcessDelay,omitempty" name:"ProcessDelay"`
	// 是否开启投递服务日志。1：关闭，2：开启。

	HasServicesLog *uint64 `json:"HasServicesLog,omitempty" name:"HasServicesLog"`
	// 预聚合检索语句

	RecordingRuleContent *string `json:"RecordingRuleContent,omitempty" name:"RecordingRuleContent"`
	// 指标名称

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 自定义指标名称

	CustomMetricLabels []*MetricLabel `json:"CustomMetricLabels,omitempty" name:"CustomMetricLabels"`
	// yaml配置文件id

	YamlId *string `json:"YamlId,omitempty" name:"YamlId"`
	// yaml配置文件名称

	YamlConfigName *string `json:"YamlConfigName,omitempty" name:"YamlConfigName"`
	// 目标日志主题id

	DstTopicId *string `json:"DstTopicId,omitempty" name:"DstTopicId"`
}

type DescribeAccountInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeAccountInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TopicPartitionInfo struct {

	// 日志主题ID

	TopicID *string `json:"TopicID,omitempty" name:"TopicID"`
	// 分区id列表

	Partitions []*uint64 `json:"Partitions,omitempty" name:"Partitions"`
}

type CreateFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件夹ID。

		Id *string `json:"Id,omitempty" name:"Id"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConfigsRequest struct {
	*tchttp.BaseRequest

	// configName&nbsp;&nbsp;按照【采集配置名称】进行模糊匹配过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;configId&nbsp;&nbsp;按照【采集配置ID】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;topicId&nbsp;&nbsp;按照【日志主题】进行过滤。&nbsp;类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为5。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeConfigsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConfigsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题列表

		Topics []*TopicInfo `json:"Topics,omitempty" name:"Topics"`
		// 总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type GenBeginRegexRequest struct {
	*tchttp.BaseRequest

	// 日志样例

	LogSample *string `json:"LogSample,omitempty" name:"LogSample"`
}

func (r *GenBeginRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenBeginRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteFolderResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFolderResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteFolderResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCosRechargeRequest struct {
	*tchttp.BaseRequest

	// 日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 投递任务名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// COS存储桶，详见产品支持的存储桶命名规范。

	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`
	// COS存储桶所在地域，详见产品支持的地域列表。

	BucketRegion *string `json:"BucketRegion,omitempty" name:"BucketRegion"`
	// COS文件所在文件夹的前缀。默认为空，投递存储桶下所有的文件。

	Prefix *string `json:"Prefix,omitempty" name:"Prefix"`
	// 采集的日志类型，json_log代表json格式日志，delimiter_log代表分隔符格式日志，minimalist_log代表单行全文；&nbsp;默认为minimalist_log

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 压缩格式，支持无压缩，gzip，lzop，snappy。默认值为无压缩。

	Compress *string `json:"Compress,omitempty" name:"Compress"`
	// 提取规则，如果设置了ExtractRule，则必须设置LogType

	ExtractRuleInfo *ExtractRuleInfo `json:"ExtractRuleInfo,omitempty" name:"ExtractRuleInfo"`
	// COS导入任务类型。1：一次性导入任务；2：持续性导入任务。默认为1：一次性导入任务

	TaskType *uint64 `json:"TaskType,omitempty" name:"TaskType"`
	// 元数据。

	Metadata []*string `json:"Metadata,omitempty" name:"Metadata"`
}

func (r *CreateCosRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCosRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaRechargesRequest struct {
	*tchttp.BaseRequest

	// 日志主题&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 导入配置ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 状态&nbsp;status&nbsp;1:&nbsp;运行中,&nbsp;2:&nbsp;暂停...

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，默认值为20，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeKafkaRechargesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaRechargesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyLogsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyLogsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyLogsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RetryScheduledSqlTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 调度结果，1:运行中&nbsp;2:成功&nbsp;3:失败

		Status *int64 `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RetryScheduledSqlTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RetryScheduledSqlTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 日志主题分区个数。默认创建1个，最大支持创建10个分区。

	PartitionCount *int64 `json:"PartitionCount,omitempty" name:"PartitionCount"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的日志主题。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启自动分裂，默认值为true。true表示开启自动分裂，false表示关闭自动分裂。

	AutoSplit *bool `json:"AutoSplit,omitempty" name:"AutoSplit"`
	// 开启自动分裂后，每个主题能够允许的最大分区数，默认值为50

	MaxSplitPartitions *int64 `json:"MaxSplitPartitions,omitempty" name:"MaxSplitPartitions"`
	// 日志主题的存储类型，可选值&nbsp;hot（实时存储）；默认为hot。

	StorageType *string `json:"StorageType,omitempty" name:"StorageType"`
	// 生命周期，单位天；可取值范围1~90。默认30天

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 二级产品标识&nbsp;测试

	SubAssumerName *string `json:"SubAssumerName,omitempty" name:"SubAssumerName"`
	// 日志主题描述

	Describes *string `json:"Describes,omitempty" name:"Describes"`
	// 0：关闭日志沉降。&nbsp;非0：开启日志沉降后标准存储的天数。HotPeriod需要大于等于7，且小于Period。仅在StorageType为&nbsp;hot&nbsp;时生效

	HotPeriod *uint64 `json:"HotPeriod,omitempty" name:"HotPeriod"`
	// 加密相关参数。&nbsp;支持加密地域并且开白用户可以传此参数，其他场景不能传递该参数。&nbsp;&nbsp;0或者不传：&nbsp;不加密&nbsp;1：kms-cls&nbsp;云产品秘钥加密

	Encryption *uint64 `json:"Encryption,omitempty" name:"Encryption"`
	// 0:&nbsp;日志主题topic&nbsp;1:&nbsp;时序主题topic

	BizType *uint64 `json:"BizType,omitempty" name:"BizType"`
	// 日志主题ID，格式为：用户自定义部分-用户appid，用户自定义部分仅支持小写字母、数字和-，且不能以-开头和结尾，长度为3至40字符，结尾使用-拼接用户appid

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// webtracking开关；&nbsp;false:&nbsp;关闭&nbsp;true：&nbsp;开启

	IsWebTracking *bool `json:"IsWebTracking,omitempty" name:"IsWebTracking"`
	// 专用集群ID

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
	// 日志主题扩展信息

	Extends *TopicExtendInfo `json:"Extends,omitempty" name:"Extends"`
}

func (r *CreateTopicRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAgentMachineGroupMetadataRequest struct {
	*tchttp.BaseRequest

	// Agent安装机器IP

	AgentIp *string `json:"AgentIp,omitempty" name:"AgentIp"`
	// Agent版本号

	AgentVersion *string `json:"AgentVersion,omitempty" name:"AgentVersion"`
	// Agent实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Agent机器组标签列表

	Labels []*string `json:"Labels,omitempty" name:"Labels"`
}

func (r *DescribeAgentMachineGroupMetadataRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAgentMachineGroupMetadataRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeConsumerOffsetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 消费者组标识

		ConsumerGroup *string `json:"ConsumerGroup,omitempty" name:"ConsumerGroup"`
		// 消费点位信息

		TopicPartitionOffsetsInfo []*TopicPartitionOffsetInfo `json:"TopicPartitionOffsetsInfo,omitempty" name:"TopicPartitionOffsetsInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerOffsetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerOffsetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeShippersRequest struct {
	*tchttp.BaseRequest

	// shipperName：按照【投递规则名称】进行过滤。&nbsp;类型：String。&nbsp;必选：否&nbsp;shipperId：按照【投递规则ID】进行过滤。&nbsp;类型：String。&nbsp;必选：否&nbsp;topicId：按照【日志主题】进行过滤。&nbsp;类型：String。&nbsp;必选：否&nbsp;taskStatus&nbsp;按照【任务运行状态】进行过滤。&nbsp;支持0：停止，1：运行中，2：异常&nbsp;类型：String&nbsp;必选：否&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为10。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 分页的偏移量，默认值为0

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页的限制数目，默认值为20，最大值100

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 控制Filters相关字段是否为精确匹配。&nbsp;0:&nbsp;默认值，shipperName模糊匹配&nbsp;1:&nbsp;shipperName&nbsp;精确匹配

	PreciseSearch *uint64 `json:"PreciseSearch,omitempty" name:"PreciseSearch"`
}

func (r *DescribeShippersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeShippersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IndexException struct {

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 状态：&nbsp;1：未开启索引&nbsp;2：未开启键值索引

	Status *int64 `json:"Status,omitempty" name:"Status"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type TaskLogStatistic struct {

	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 读取的源日志主题的行数

	ReadLines *int64 `json:"ReadLines,omitempty" name:"ReadLines"`
	// 加工后的行数

	WriteLines *int64 `json:"WriteLines,omitempty" name:"WriteLines"`
	// 加工失败的行数

	FailedLines *int64 `json:"FailedLines,omitempty" name:"FailedLines"`
	// 输出到目标日志主题的总体统计数据

	DstTopicLogStatistics []*TopicIdLogStatistic `json:"DstTopicLogStatistics,omitempty" name:"DstTopicLogStatistics"`
	// 加工过滤的行数

	FilterLines *uint64 `json:"FilterLines,omitempty" name:"FilterLines"`
}

type NoticeContent struct {

	// 渠道类型&nbsp;&nbsp;Email:邮件;Sms:短信;WeChat:微信;Phone:电话;WeCom:企业微信;DingTalk:钉钉;Lark:飞书;Http:自定义回调;

	Type *string `json:"Type,omitempty" name:"Type"`
	// 告警触发通知内容模板。

	TriggerContent *NoticeContentInfo `json:"TriggerContent,omitempty" name:"TriggerContent"`
	// 告警恢复通知内容模板。

	RecoveryContent *NoticeContentInfo `json:"RecoveryContent,omitempty" name:"RecoveryContent"`
}

type CreateMachineGroupRequest struct {
	*tchttp.BaseRequest

	// 机器组名字，不能重复

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
	// 创建机器组类型。取值如下：&nbsp;&nbsp;Type：ip，Values中为ip字符串列表创建机器组&nbsp;Type：label，Values中为标签字符串列表创建机器组

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的机器组。最大支持10个标签键值对，同一个资源只能绑定到同一个标签键下。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 是否开启机器组自动更新。true表示开启机器组自动更新，false表示关闭机器组自动更新。是否开启机器组自动更新。默认false

	AutoUpdate *bool `json:"AutoUpdate,omitempty" name:"AutoUpdate"`
	// 升级开始时间，建议业务低峰期升级LogListener

	UpdateStartTime *string `json:"UpdateStartTime,omitempty" name:"UpdateStartTime"`
	// 升级结束时间，建议业务低峰期升级LogListener

	UpdateEndTime *string `json:"UpdateEndTime,omitempty" name:"UpdateEndTime"`
	// 是否开启服务日志，用于记录因Loglistener&nbsp;服务自身产生的log，开启后，会创建内部日志集cls_service_logging和日志主题loglistener_status,loglistener_alarm,loglistener_business，不产生计费。默认false

	ServiceLogging *bool `json:"ServiceLogging,omitempty" name:"ServiceLogging"`
	// TKE标志位，默认值为空字符串。空字符串表示日志不是来自于TKE，label_k8s表示日志来自于TKE。

	Flag *string `json:"Flag,omitempty" name:"Flag"`
	// 机器组中机器离线清理时间。单位：天

	DelayCleanupTime *int64 `json:"DelayCleanupTime,omitempty" name:"DelayCleanupTime"`
	// 机器组元数据信息列表

	MetaTags []*MetaTagInfo `json:"MetaTags,omitempty" name:"MetaTags"`
	// 系统类型，取值如下：&nbsp;&nbsp;0：Linux&nbsp;（默认值）&nbsp;1：Windows

	OSType *uint64 `json:"OSType,omitempty" name:"OSType"`
	// 集群ID。&nbsp;-&nbsp;当Flag为&nbsp;`label_tke`，表示TKE集群ID&nbsp;-&nbsp;ClusterId&nbsp;与&nbsp;ClusterRegion&nbsp;必须同时填写，或者为空

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群所在地域&nbsp;-&nbsp;当Flag为&nbsp;`label_tke`，表示TKE集群地域&nbsp;-&nbsp;ClusterId&nbsp;与&nbsp;ClusterRegion&nbsp;必须同时填写，或者为空

	ClusterRegion *string `json:"ClusterRegion,omitempty" name:"ClusterRegion"`
	// 专用集群ID

	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" name:"DedicatedClusterId"`
}

func (r *CreateMachineGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateMachineGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyKafkaRechargeRequest struct {
	*tchttp.BaseRequest

	// Kafka导入配置ID

	Id *string `json:"Id,omitempty" name:"Id"`
	// 导入CLS目标topic&nbsp;ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// Kafka导入配置名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 导入Kafka类型，0:&nbsp;CKafka，1:&nbsp;用户自建Kafka

	KafkaType *uint64 `json:"KafkaType,omitempty" name:"KafkaType"`
	// CKafka实例ID，KafkaType为0时必填

	KafkaInstance *string `json:"KafkaInstance,omitempty" name:"KafkaInstance"`
	// 服务地址，KafkaType为1时必填。

	ServerAddr *string `json:"ServerAddr,omitempty" name:"ServerAddr"`
	// ServerAddr是否为加密连接

	IsEncryptionAddr *bool `json:"IsEncryptionAddr,omitempty" name:"IsEncryptionAddr"`
	// 加密访问协议，IsEncryptionAddr参数为true时必填

	Protocol *KafkaProtocolInfo `json:"Protocol,omitempty" name:"Protocol"`
	// 用户需要导入的Kafka相关topic列表，多个topic之间使用半角逗号隔开

	UserKafkaTopics *string `json:"UserKafkaTopics,omitempty" name:"UserKafkaTopics"`
	// 用户Kafka消费组名称

	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" name:"ConsumerGroupName"`
	// 日志导入规则

	LogRechargeRule *LogRechargeRuleInfo `json:"LogRechargeRule,omitempty" name:"LogRechargeRule"`
	// 导入控制，1：暂停，2：继续

	StatusControl *uint64 `json:"StatusControl,omitempty" name:"StatusControl"`
	// 私有网络信息参数

	NetworkInfo *NetworkInfo `json:"NetworkInfo,omitempty" name:"NetworkInfo"`
}

func (r *ModifyKafkaRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyKafkaRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAlarmNoticesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 告警通知模板列表。

		AlarmNotices []*AlarmNotice `json:"AlarmNotices,omitempty" name:"AlarmNotices"`
		// 符合条件的告警通知模板总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type MachineGroupTypeInfo struct {

	// 机器组类型。支持&nbsp;ip&nbsp;和&nbsp;label。&nbsp;&nbsp;ip：表示该机器组Values中存的是采集机器的ip地址&nbsp;label：表示该机器组Values中存储的是机器的标签

	Type *string `json:"Type,omitempty" name:"Type"`
	// 机器描述列表。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeConsumerResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 投递任务是否生效

		Effective *bool `json:"Effective,omitempty" name:"Effective"`
		// 是否投递日志的元数据信息

		NeedContent *bool `json:"NeedContent,omitempty" name:"NeedContent"`
		// 如果需要投递元数据信息，元数据信息的描述

		Content *ConsumerContent `json:"Content,omitempty" name:"Content"`
		// CKafka的描述

		Ckafka *Ckafka `json:"Ckafka,omitempty" name:"Ckafka"`
		// 压缩方式[0:NONE；2:SNAPPY；3:LZ4]

		Compression *int64 `json:"Compression,omitempty" name:"Compression"`
		// 任务创建时间

		CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`
		// 角色访问描述名

		RoleArn *string `json:"RoleArn,omitempty" name:"RoleArn"`
		// 外部ID

		ExternalId *string `json:"ExternalId,omitempty" name:"ExternalId"`
		// 任务运行状态。支持0,1,2&nbsp;-&nbsp;0:&nbsp;停止&nbsp;-&nbsp;1:&nbsp;运行中&nbsp;-&nbsp;2:&nbsp;异常

		TaskStatus *uint64 `json:"TaskStatus,omitempty" name:"TaskStatus"`
		// 高级配置

		AdvancedConfig *AdvancedConsumerConfiguration `json:"AdvancedConfig,omitempty" name:"AdvancedConfig"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeConsumerResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeConsumerResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeKafkaConsumeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Kafka消费信息

		Kafka *KafkaInfo `json:"Kafka,omitempty" name:"Kafka"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeKafkaConsumeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeKafkaConsumeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type BinlogInitialPoint struct {

	// 起始位置类型。&nbsp;1:最新位置&nbsp;2:指定位置&nbsp;3.指定GTID

	Type *uint64 `json:"Type,omitempty" name:"Type"`
	// Binlog文件名称。&nbsp;当起始位置类型Type=2时，BinlogFileName&nbsp;必填。

	BinlogFileName *string `json:"BinlogFileName,omitempty" name:"BinlogFileName"`
	// Binlog文件偏移量。&nbsp;当起始位置类型Type=2时，&nbsp;BinlogPosition必填。

	BinlogPosition *uint64 `json:"BinlogPosition,omitempty" name:"BinlogPosition"`
	// GTID位置。&nbsp;当起始位置类型Type=3时，&nbsp;GTIDPosition必填。

	GTIDPosition *string `json:"GTIDPosition,omitempty" name:"GTIDPosition"`
}

type MetricSubscribeInfo struct {

	// 订阅任务id。

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 订阅任务名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 云产品命名空间。

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 指标配置信息。

	Metrics []*MetricConfig `json:"Metrics,omitempty" name:"Metrics"`
	// 实例配置信息。

	InstanceInfo *InstanceConfig `json:"InstanceInfo,omitempty" name:"InstanceInfo"`
	// 订阅任务开关。1:暂停&nbsp;2:启用

	Enable *uint64 `json:"Enable,omitempty" name:"Enable"`
	// 订阅任务运行状态。0:创建中&nbsp;1:暂停&nbsp;2:运行中&nbsp;3:异常

	Status *uint64 `json:"Status,omitempty" name:"Status"`
	// 订阅任务运行异常时的错误信息。

	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
	// 创建时间（秒级时间戳）

	CreateTime *int64 `json:"CreateTime,omitempty" name:"CreateTime"`
	// 更新时间（秒级时间戳）

	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CreateConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 采集配置ID

		ConfigId *string `json:"ConfigId,omitempty" name:"ConfigId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteRecordingRuleTaskRequest struct {
	*tchttp.BaseRequest

	// 任务ID

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteRecordingRuleTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteRecordingRuleTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Ckafka struct {

	// Ckafka&nbsp;的&nbsp;Vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// Ckafka&nbsp;的&nbsp;Vport

	Vport *string `json:"Vport,omitempty" name:"Vport"`
	// Ckafka&nbsp;的&nbsp;InstanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Ckafka&nbsp;的&nbsp;InstanceName

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// Ckafka&nbsp;的&nbsp;TopicId

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// Ckafka&nbsp;的&nbsp;TopicName

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}

type ExceptionInfo struct {

	// 表示该异常问题是否需要展示。true表示需要展示异常问题，false表示不需要展示异常问题

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 异常资源ID：&nbsp;1：索引配置异常&nbsp;2：采集配置异常&nbsp;3：日志上传异常

	ExceptionId *int64 `json:"ExceptionId,omitempty" name:"ExceptionId"`
}

type CreateFolderRequest struct {
	*tchttp.BaseRequest

	// 文件夹名称。&nbsp;&nbsp;必填项。&nbsp;名称限制不能超过50个字符。&nbsp;重名报错。

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateFolderRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateFolderRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GenKVRegexRequest struct {
	*tchttp.BaseRequest

	// 日志样例

	LogSample *string `json:"LogSample,omitempty" name:"LogSample"`
	// 样例索引列表

	Indexes []*RegexIndexInfo `json:"Indexes,omitempty" name:"Indexes"`
	// 是否为多行全文，默认单行

	MultiLine *bool `json:"MultiLine,omitempty" name:"MultiLine"`
}

func (r *GenKVRegexRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GenKVRegexRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CosFieldsData struct {

	// COS文件（csv格式）行字段信息

	RowInfo []*string `json:"RowInfo,omitempty" name:"RowInfo"`
}

type CreateRemoteWriteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// remoteWrite任务id

		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateRemoteWriteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateRemoteWriteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteCosRechargeRequest struct {
	*tchttp.BaseRequest

	// COS导入配置Id

	Id *string `json:"Id,omitempty" name:"Id"`
	// 日志主题Id

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteCosRechargeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteCosRechargeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRemoteWriteTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// RemoteWrite&nbsp;信息列表

		Infos []*RemoteWriteInfo `json:"Infos,omitempty" name:"Infos"`
		// RemoteWrite信息总条数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRemoteWriteTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRemoteWriteTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateAlarmNoticeRequest struct {
	*tchttp.BaseRequest

	// 通知渠道组名称。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）&nbsp;需要发送通知的告警类型。可选值：&nbsp;&nbsp;Trigger&nbsp;-&nbsp;告警触发&nbsp;Recovery&nbsp;-&nbsp;告警恢复&nbsp;All&nbsp;-&nbsp;告警触发和告警恢复

	Type *string `json:"Type,omitempty" name:"Type"`
	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）&nbsp;通知接收对象。

	NoticeReceivers []*NoticeReceiver `json:"NoticeReceivers,omitempty" name:"NoticeReceivers"`
	// 【简易模式】（简易模式/告警模式二选一，分别配置相应参数）&nbsp;接口回调信息（包括企业微信、钉钉、飞书）。

	WebCallbacks []*WebCallback `json:"WebCallbacks,omitempty" name:"WebCallbacks"`
	// 标签描述列表，通过指定该参数可以同时绑定标签到相应的通知渠道组。最大支持50个标签键值对，并且不能有重复的键值对。

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 查询数据链接。http://&nbsp;或者&nbsp;https://&nbsp;开头，不能/结尾

	JumpDomain *string `json:"JumpDomain,omitempty" name:"JumpDomain"`
	// 【高级模式】（简易模式/告警模式二选一，分别配置相应参数）&nbsp;通知规则。

	NoticeRules []*NoticeRule `json:"NoticeRules,omitempty" name:"NoticeRules"`
	// 投递日志开关。可取值如下：&nbsp;1：关闭（默认值）；&nbsp;2：开启&nbsp;投递日志开关开启时，&nbsp;DeliverConfig参数必填。

	DeliverStatus *uint64 `json:"DeliverStatus,omitempty" name:"DeliverStatus"`
	// 投递日志配置参数。当DeliverStatus开启时，必填。

	DeliverConfig *DeliverConfig `json:"DeliverConfig,omitempty" name:"DeliverConfig"`
	// 免登录操作告警开关。可取值如下：&nbsp;-&nbsp;1：关闭&nbsp;-&nbsp;2：开启（默认值）

	AlarmShieldStatus *uint64 `json:"AlarmShieldStatus,omitempty" name:"AlarmShieldStatus"`
}

func (r *CreateAlarmNoticeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateAlarmNoticeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMachineGroupInfoRequest struct {
	*tchttp.BaseRequest

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组类型&nbsp;目前type支持&nbsp;ip&nbsp;和&nbsp;label

	MachineGroupType *MachineGroupTypeInfo `json:"MachineGroupType,omitempty" name:"MachineGroupType"`
}

func (r *DeleteMachineGroupInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMachineGroupInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HistogramInfo struct {

	// 统计周期内的日志条数

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 按&nbsp;period&nbsp;取整后的&nbsp;unix&nbsp;timestamp：&nbsp;单位毫秒

	BTime *int64 `json:"BTime,omitempty" name:"BTime"`
}

type SearchLogErrors struct {

	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 错误码。&nbsp;（已废弃）

	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`
	// 错误信息

	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
	// 错误码

	ErrorCodeStr *string `json:"ErrorCodeStr,omitempty" name:"ErrorCodeStr"`
}

type DeleteMetricSubscribeRequest struct {
	*tchttp.BaseRequest

	// 指标采集任务id

	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
	// 指标采集任务配置的日志主题id。

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

func (r *DeleteMetricSubscribeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMetricSubscribeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmClassification struct {

	// 分类键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 分类值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeAlertRecordHistoryRequest struct {
	*tchttp.BaseRequest

	// alertId：按照告警策略ID进行过滤。类型：String&nbsp;必选：否&nbsp;topicId：按照监控对象ID进行过滤。类型：String&nbsp;必选：否&nbsp;status：按照告警状态进行过滤。类型：String&nbsp;必选：否，0代表未恢复，1代表已恢复，2代表已失效&nbsp;alarmLevel：按照告警等级进行过滤。类型：String&nbsp;必选：否，0代表警告，1代表提醒，2代表紧急&nbsp;每次请求的Filters的上限为10，Filter.Values的上限为100。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 查询时间范围启始时间，毫秒级unix时间戳

	From *uint64 `json:"From,omitempty" name:"From"`
	// 查询时间范围结束时间，毫秒级unix时间戳

	To *uint64 `json:"To,omitempty" name:"To"`
	// 分页的偏移量，默认值为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页单页限制数目，最大值100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeAlertRecordHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlertRecordHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

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

package v20220501

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type CreateHealthCheckPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 健康检测策略名称

		HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateHealthCheckPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHealthCheckPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NativeNodeConvertCostInfo struct {

	// 机器id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 询价价格，单位：分，折扣后

	Cost *float64 `json:"Cost,omitempty" name:"Cost"`
	// 询价价格，单位：分，折扣前

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 原生节点增值服务费询价价格，单位：分，折扣后

	HKServiceFeeCost *float64 `json:"HKServiceFeeCost,omitempty" name:"HKServiceFeeCost"`
	// 原生节点增值服务费询价价格，单位：分，折扣前

	HKServiceFeeTotalCost *float64 `json:"HKServiceFeeTotalCost,omitempty" name:"HKServiceFeeTotalCost"`
}

type NodePool struct {

	// 节点池类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 超级节点池参数，在Type等于Super该字段才有值

	Super *SuperNodePoolInfo `json:"Super,omitempty" name:"Super"`
	// 节点池名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 普通节点池参数，在Type等于Regular该字段才有值

	Regular *RegularNodePoolInfo `json:"Regular,omitempty" name:"Regular"`
	// 第三方节点池参数，在Type等于External该字段才有值

	External *ExternalNodePoolInfo `json:"External,omitempty" name:"External"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池&nbsp;ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 节点污点

	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
	// 节点是否不可调度

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 原生节点池参数

	Native *NativeNodePoolInfo `json:"Native,omitempty" name:"Native"`
	// 是否开启删除保护

	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
	// 原生节点池参数（已弃用）

	Hosted *HostedNodePoolInfo `json:"Hosted,omitempty" name:"Hosted"`
	// 节点&nbsp;&nbsp;Labels

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 节点池状态

	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
	// 节点标签

	Tags []*TagSpecification `json:"Tags,omitempty" name:"Tags"`
	// 节点&nbsp;Annotation&nbsp;列表

	Annotations []*Annotation `json:"Annotations,omitempty" name:"Annotations"`
}

type DescribeOperationEventsRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池&nbsp;ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 最大输出条数，默认20，最大为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// ·&nbsp;&nbsp;type
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【事件类型】进行过滤，事件类型包括Detect、Repair、Upgrade、ScaleUp、ScaleDown、NodeInitialization。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;uuid
	// &nbsp;&nbsp;&nbsp;&nbsp;根据事件的uuid进行查询。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeOperationEventsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationEventsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeMachineConfigurationRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 机型列表

	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 副本数

	Num *int64 `json:"Num,omitempty" name:"Num"`
	// 节点池&nbsp;ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 节点池扩容策略。ZoneEquality：多可用区打散；ZonePriority：首选可用区优先；

	CreatePolicy *string `json:"CreatePolicy,omitempty" name:"CreatePolicy"`
}

func (r *DescribeMachineConfigurationRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineConfigurationRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InternetAccessible struct {

	// 带宽包&nbsp;ID

	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" name:"BandwidthPackageId"`
	// IP&nbsp;地址类型

	AddressType *string `json:"AddressType,omitempty" name:"AddressType"`
	// 公网ID

	AddressId *string `json:"AddressId,omitempty" name:"AddressId"`
	// 带宽

	MaxBandwidthOut *int64 `json:"MaxBandwidthOut,omitempty" name:"MaxBandwidthOut"`
	// 网络计费方式

	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
}

type StartMachinesRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名字列表，一次请求，传入节点数量上限为100个

	MachineNames []*string `json:"MachineNames,omitempty" name:"MachineNames"`
}

func (r *StartMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyMachinesKeyIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMachinesKeyIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachinesKeyIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleNodePoolRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池&nbsp;ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 期望节点数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
}

func (r *ScaleNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleNodePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {

	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 本地专用集群Id

	CdcId *string `json:"CdcId,omitempty" name:"CdcId"`
	// 集群所在vpc的id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群描述

	ClusterDescription *string `json:"ClusterDescription,omitempty" name:"ClusterDescription"`
	// 集群版本（默认值为1.10.5）

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 标签描述列表。

	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
	// 集群状态&nbsp;(Trading&nbsp;集群开通中,Creating&nbsp;创建中,Running&nbsp;运行中,Deleting&nbsp;删除中,Idling&nbsp;闲置中,Recovering&nbsp;唤醒中,Upgrading&nbsp;升级中,NodeUpgrading&nbsp;节点升级中,RuntimeUpgrading&nbsp;节点运行时升级中,MasterScaling&nbsp;Master扩缩容中,ClusterLevelUpgrading&nbsp;调整规格中,ResourceIsolate&nbsp;欠费隔离中,ResourceIsolated&nbsp;欠费已隔离,ResourceReverse&nbsp;冲正恢复中,Abnormal&nbsp;异常)

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 集群等级，针对托管集群生效
	// 注意：此字段可能返回&nbsp;null，表示取不到有效值。

	ClusterLevel *string `json:"ClusterLevel,omitempty" name:"ClusterLevel"`
}

type Disk struct {

	// 云盘大小(G）

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 是否自动化格式盘并挂载

	AutoFormatAndMount *bool `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	// 文件系统

	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`
	// 挂载目录

	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`
	// 云盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
}

type GPUConfig struct {

	// 机型名称

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// GPU相关的参数，包括驱动版本，CUDA版本，cuDNN版本，是否开启MIG以及是否开启Fabric等

	GPUParams *GPUParams `json:"GPUParams,omitempty" name:"GPUParams"`
}

type InquirePriceRenewNativeNodeRequest struct {
	*tchttp.BaseRequest

	// 包年包月节点续费时长，单位：月。

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 资源&nbsp;ID&nbsp;列表

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

func (r *InquirePriceRenewNativeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceRenewNativeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterCreateNativeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下单参数&nbsp;JSON&nbsp;串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterCreateNativeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterCreateNativeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeCountSummary struct {

	// 自动管理的节点

	AutoscalingAdded *AutoscalingAdded `json:"AutoscalingAdded,omitempty" name:"AutoscalingAdded"`
	// 手动管理的节点

	ManuallyAdded *ManuallyAdded `json:"ManuallyAdded,omitempty" name:"ManuallyAdded"`
}

type GPUParams struct {

	// 自定义驱动下载地址

	CustomGPUDriver *string `json:"CustomGPUDriver,omitempty" name:"CustomGPUDriver"`
	// 是否支持qgpu

	QGPU *bool `json:"QGPU,omitempty" name:"QGPU"`
	// GPU驱动版本

	Driver *string `json:"Driver,omitempty" name:"Driver"`
	// CUDA版本

	CUDA *string `json:"CUDA,omitempty" name:"CUDA"`
	// CUDNN版本

	CUDNN *string `json:"CUDNN,omitempty" name:"CUDNN"`
	// 是否启用MIG特性

	MIGEnable *bool `json:"MIGEnable,omitempty" name:"MIGEnable"`
	// 是否启用Fabric特性

	Fabric *bool `json:"Fabric,omitempty" name:"Fabric"`
}

type Annotation struct {

	// map表中的Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// map表中的Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type HealthCheckPolicyBinding struct {

	// 关联节点池数组

	NodePools []*string `json:"NodePools,omitempty" name:"NodePools"`
	// 健康检测策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 规则创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
}

type InquirePriceConvertNativeNodeRequest struct {
	*tchttp.BaseRequest

	// 机器id数组

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *InquirePriceConvertNativeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceConvertNativeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodePoolsRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 查询过滤条件：
	// ·&nbsp;&nbsp;NodePoolsName
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【节点池名】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;NodePoolsId
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【节点池id】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;tags
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【标签键值对】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;tag:tag-key
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【标签键值对】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最大输出条数，默认20，最大为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeNodePoolsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodePoolsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalNodePoolInfo struct {

	// 第三方节点Runtime配置

	RuntimeConfig *RuntimeConfig `json:"RuntimeConfig,omitempty" name:"RuntimeConfig"`
	// Machine系统配置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点池健康检查规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// 节点相关的自定义参数信息

	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
	// base64&nbsp;编码的用户脚本,&nbsp;此脚本会在&nbsp;k8s&nbsp;组件运行后执行,&nbsp;需要用户保证脚本的可重入及重试逻辑,&nbsp;脚本及其生成的日志文件可在节点的&nbsp;/data/ccs_userscript/&nbsp;路径查看,&nbsp;如果要求节点需要在进行初始化完成后才可加入调度,&nbsp;可配合&nbsp;unschedulable&nbsp;参数使用,&nbsp;在&nbsp;userScript&nbsp;最后初始化完成后,&nbsp;添加&nbsp;kubectl&nbsp;uncordon&nbsp;nodename&nbsp;--kubeconfig=/root/.kube/config&nbsp;命令使节点加入调度

	UserScript *string `json:"UserScript,omitempty" name:"UserScript"`
	// 节点数

	NodesNum *uint64 `json:"NodesNum,omitempty" name:"NodesNum"`
	// 节点kubelet版本

	KubeletVersion *string `json:"KubeletVersion,omitempty" name:"KubeletVersion"`
	// 节点网络类型&nbsp;公网：Public&nbsp;专线：Direct

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
}

type MachineSetScaling struct {

	// 节点池最小副本数

	MinReplicas *int64 `json:"MinReplicas,omitempty" name:"MinReplicas"`
	// 节点池最大副本数

	MaxReplicas *int64 `json:"MaxReplicas,omitempty" name:"MaxReplicas"`
	// 节点池扩容策略。ZoneEquality：多可用区打散；ZonePriority：首选可用区优先；

	CreatePolicy *string `json:"CreatePolicy,omitempty" name:"CreatePolicy"`
}

type OperationEvent struct {

	// 事件&nbsp;UUID

	Uuid *string `json:"Uuid,omitempty" name:"Uuid"`
	// 事件类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 事件状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 开始时间

	BeginAt *string `json:"BeginAt,omitempty" name:"BeginAt"`
	// 结束时间

	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 事件相关节点池&nbsp;ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 事件出现原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 事件详细信息

	Message *string `json:"Message,omitempty" name:"Message"`
	// 事件相关对象

	Related *string `json:"Related,omitempty" name:"Related"`
}

type DescribeHealthCheckPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 健康检测策略数组

		HealthCheckPolicies []*HealthCheckPolicy `json:"HealthCheckPolicies,omitempty" name:"HealthCheckPolicies"`
		// 数组总数目

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHealthCheckPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHealthCheckPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceHousekeeperRequest struct {
	*tchttp.BaseRequest

	// 机型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 按量计费：POSTPAID_BY_HOUR，包年包月：PREPAID。默认值：按量计费

	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
	// 询价时长，按量计费：单位秒，包年包月：单位月。默认值按量计费：3600秒，包年包月：一个月。

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 购买数量，默认值1

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 系统盘类型，取值为：CLOUD_SSD、CLOUD_PREMIUM，默认值为高性能云盘：CLOUD_PREMIUM

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘大小，单位：Gi，默认值：50

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
}

func (r *InquirePriceHousekeeperRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceHousekeeperRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LifecycleConfig struct {

	// 节点初始化前自定义脚本

	PreInit *string `json:"PreInit,omitempty" name:"PreInit"`
	// 节点初始化后自定义脚本

	PostInit *string `json:"PostInit,omitempty" name:"PostInit"`
}

type VerifyQGPURequest struct {
	*tchttp.BaseRequest

	// 实例机型名称

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// GPU驱动版本号

	Driver *string `json:"Driver,omitempty" name:"Driver"`
	// QGPU&nbsp;addon&nbsp;版本

	QGPUVersion *string `json:"QGPUVersion,omitempty" name:"QGPUVersion"`
}

func (r *VerifyQGPURequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyQGPURequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeParamUpdateProcesses struct {

	// 更新的参数名称

	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
	// 更新执行的阶段：updating(更新中)、paused（暂停更新）、finished（结束更新）

	UpdateProcess *string `json:"UpdateProcess,omitempty" name:"UpdateProcess"`
	// 存量更新参数成功的节点数目

	UpdatedSuccessNodeNum *int64 `json:"UpdatedSuccessNodeNum,omitempty" name:"UpdatedSuccessNodeNum"`
}

type ModifyNodePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyNodePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegularNodePoolInfo struct {

	// 节点池osName

	NodePoolOs *string `json:"NodePoolOs,omitempty" name:"NodePoolOs"`
	// 节点配置

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)

	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`
	// 期望的节点数量

	DesiredNodesNum *int64 `json:"DesiredNodesNum,omitempty" name:"DesiredNodesNum"`
	// AutoscalingGroupId&nbsp;分组id

	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitempty" name:"AutoscalingGroupId"`
	// NodeCountSummary&nbsp;节点列表

	NodeCountSummary *NodeCountSummary `json:"NodeCountSummary,omitempty" name:"NodeCountSummary"`
	// 状态信息

	AutoscalingGroupStatus *string `json:"AutoscalingGroupStatus,omitempty" name:"AutoscalingGroupStatus"`
	// 最大节点数量

	MaxNodesNum *int64 `json:"MaxNodesNum,omitempty" name:"MaxNodesNum"`
	// 最小节点数量

	MinNodesNum *int64 `json:"MinNodesNum,omitempty" name:"MinNodesNum"`
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// LaunchConfigurationId&nbsp;配置

	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`
}

type TagSpecification struct {

	// 标签对列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
	// 标签绑定的资源类型，当前支持类型：
	// 1.cluster：集群相关接口，TagSpecification&nbsp;的&nbsp;ResourceType&nbsp;传参为&nbsp;cluster
	// 2.machine：节点池相关接口，如：CreateNodePool,&nbsp;DescribeNodePools&nbsp;等，TagSpecification&nbsp;的&nbsp;ResourceType&nbsp;传参为&nbsp;machine

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

type DescribeGPUInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// GPU相关配置

		GPUParams []*GPUParams `json:"GPUParams,omitempty" name:"GPUParams"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGPUInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGPUInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTradeTypesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机型对应的计费参数，是&nbsp;json&nbsp;格式化的数组

		TradeInfo *string `json:"TradeInfo,omitempty" name:"TradeInfo"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTradeTypesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTradeTypesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceAdvancedSettings struct {

	// 该节点属于podCIDR大小自定义模式时，可指定节点上运行的pod数量上限

	DesiredPodNumber *int64 `json:"DesiredPodNumber,omitempty" name:"DesiredPodNumber"`
	// GPU驱动相关参数

	GPUArgs *GPUArgs `json:"GPUArgs,omitempty" name:"GPUArgs"`
	// base64&nbsp;编码的用户脚本，在初始化节点之前执行，目前只对添加已有节点生效

	PreStartUserScript *string `json:"PreStartUserScript,omitempty" name:"PreStartUserScript"`
	// 运行时描述

	RuntimeConfig *RuntimeConfig `json:"RuntimeConfig,omitempty" name:"RuntimeConfig"`
	// base64&nbsp;编码的用户脚本,&nbsp;此脚本会在&nbsp;k8s&nbsp;组件运行后执行,&nbsp;需要用户保证脚本的可重入及重试逻辑,&nbsp;脚本及其生成的日志文件可在节点的&nbsp;/data/ccs_userscript/&nbsp;路径查看,&nbsp;如果要求节点需要在进行初始化完成后才可加入调度,&nbsp;可配合&nbsp;unschedulable&nbsp;参数使用,&nbsp;在&nbsp;userScript&nbsp;最后初始化完成后,&nbsp;添加&nbsp;kubectl&nbsp;uncordon&nbsp;nodename&nbsp;--kubeconfig=/root/.kube/config&nbsp;命令使节点加入调度

	UserScript *string `json:"UserScript,omitempty" name:"UserScript"`
	// 多盘数据盘挂载信息：新建节点时请确保购买CVM的参数传递了购买多个数据盘的信息，如CreateClusterInstances&nbsp;API的RunInstancesPara下的DataDisks也需要设置购买多个数据盘,&nbsp;具体可以参考CreateClusterInstances接口的添加集群节点(多块数据盘)样例；添加已有节点时，请确保填写的分区信息在节点上真实存在

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 节点相关的自定义参数信息

	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
}

type SortBy struct {

	// 排序指标

	FieldName *string `json:"FieldName,omitempty" name:"FieldName"`
	// 排序方式

	OrderType *string `json:"OrderType,omitempty" name:"OrderType"`
}

type HealthCheckTemplateRule struct {

	// 健康检测规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 修复动作

	RepairAction *string `json:"RepairAction,omitempty" name:"RepairAction"`
	// 修复影响

	RepairEffect *string `json:"RepairEffect,omitempty" name:"RepairEffect"`
	// 是否建议开启检测

	ShouldEnable *bool `json:"ShouldEnable,omitempty" name:"ShouldEnable"`
	// 是否建议修复

	ShouldRepair *bool `json:"ShouldRepair,omitempty" name:"ShouldRepair"`
	// 问题严重程度

	Severity *string `json:"Severity,omitempty" name:"Severity"`
	// 健康检测项目名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type ConvertRegularToNativeNodeRequest struct {
	*tchttp.BaseRequest

	// 每个实例的转化超时时间（分钟），默认&nbsp;0，不超时

	Timeout *uint64 `json:"Timeout,omitempty" name:"Timeout"`
	// 转化节点的时候是否重装节点

	ReInstall *bool `json:"ReInstall,omitempty" name:"ReInstall"`
	// 集群的&nbsp;ID

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 普通节点的实例&nbsp;ID&nbsp;列表，如果不传，则转化维度必须是节点池维度

	InstanceIDs []*string `json:"InstanceIDs,omitempty" name:"InstanceIDs"`
	// 转化的维度，Node&nbsp;或&nbsp;NodePool，默认为&nbsp;Node

	ConvertDimension *string `json:"ConvertDimension,omitempty" name:"ConvertDimension"`
	// 当转化维度为节点池维度的时候，这个参数必须传

	OriginalNodePoolID *string `json:"OriginalNodePoolID,omitempty" name:"OriginalNodePoolID"`
	// 转化的目标原生节点池&nbsp;ID

	TargetNodePoolID *string `json:"TargetNodePoolID,omitempty" name:"TargetNodePoolID"`
}

func (r *ConvertRegularToNativeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConvertRegularToNativeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeTradeTypesRequest struct {
	*tchttp.BaseRequest

	// 机型对应的实例族

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
}

func (r *DescribeTradeTypesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeTradeTypesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoUpgradeOptions struct {

	// 自动升级开始时间

	AutoUpgradeStartTime *string `json:"AutoUpgradeStartTime,omitempty" name:"AutoUpgradeStartTime"`
	// 自动升级持续时间

	Duration *string `json:"Duration,omitempty" name:"Duration"`
	// 运维日期

	WeeklyPeriod []*string `json:"WeeklyPeriod,omitempty" name:"WeeklyPeriod"`
}

type RebootMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RebootMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例异常(或者处于初始化中)的原因

	FailedReason *string `json:"FailedReason,omitempty" name:"FailedReason"`
	// 实例的状态
	// -&nbsp;initializing创建中
	// -&nbsp;running&nbsp;运行中
	// -&nbsp;failed&nbsp;异常

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点角色,&nbsp;MASTER,&nbsp;WORKER,&nbsp;ETCD,&nbsp;MASTER_ETCD,ALL,&nbsp;默认为WORKER

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 添加时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 超级节点参数

	Super *SuperNodeInfo `json:"Super,omitempty" name:"Super"`
	// 是否不可调度

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 节点内网IP

	LanIP *string `json:"LanIP,omitempty" name:"LanIP"`
	// 资源池ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 原生节点参数

	Native *NativeNodeInfo `json:"Native,omitempty" name:"Native"`
	// 普通节点参数

	Regular *RegularNodeInfo `json:"Regular,omitempty" name:"Regular"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 第三方节点参数

	External *ExternalNodeInfo `json:"External,omitempty" name:"External"`
	// 实例是否封锁状态

	DrainStatus *string `json:"DrainStatus,omitempty" name:"DrainStatus"`
}

type DescribeClusterMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点池节点列表

		Machines []*Machine `json:"Machines,omitempty" name:"Machines"`
		// 资源总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleNodePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ScaleNodePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ScaleNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstantInspectTaskRequest struct {
	*tchttp.BaseRequest

	// 集群巡检任务名称，在调用CreateInstantInspectJob接口后获得

	JobName *string `json:"JobName,omitempty" name:"JobName"`
}

func (r *DescribeInstantInspectTaskRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstantInspectTaskRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHealthCheckPolicyBindingsRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// ·&nbsp;&nbsp;HealthCheckPolicyName
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【健康检测规则名称】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
	// 最大输出条数，默认20，最大为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeHealthCheckPolicyBindingsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHealthCheckPolicyBindingsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRefundNativeNodeRequest struct {
	*tchttp.BaseRequest

	// 退费资源&nbsp;ID&nbsp;列表

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
}

func (r *InquirePriceRefundNativeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceRefundNativeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateHealthCheckPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 健康检测策略

	HealthCheckPolicy *HealthCheckPolicy `json:"HealthCheckPolicy,omitempty" name:"HealthCheckPolicy"`
}

func (r *CreateHealthCheckPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateHealthCheckPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostedNodePoolInfo struct {

	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 节点计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// machine&nbsp;系统配置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 期望节点数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 伸缩配置

	Scaling *MachineSetScaling `json:"Scaling,omitempty" name:"Scaling"`
	// 安全组列表

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 自动升级配置

	UpgradeSettings *MachineUpgradeSettings `json:"UpgradeSettings,omitempty" name:"UpgradeSettings"`
	// 系统盘配置

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 密钥&nbsp;ID&nbsp;列表

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// JoiningReplicas.

	JoiningReplicas *int64 `json:"JoiningReplicas,omitempty" name:"JoiningReplicas"`
	// JoiningReplicas.

	UpdateExistedNode *bool `json:"UpdateExistedNode,omitempty" name:"UpdateExistedNode"`
	// JoiningReplicas.

	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`
	// kubelet&nbsp;自定义参数

	KubeletArgs []*string `json:"KubeletArgs,omitempty" name:"KubeletArgs"`
	// 运行时根目录

	RuntimeRootDir *string `json:"RuntimeRootDir,omitempty" name:"RuntimeRootDir"`
	// 是否开启弹性伸缩

	EnableAutoscaling *bool `json:"EnableAutoscaling,omitempty" name:"EnableAutoscaling"`
	// 机型列表

	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
	// JoiningReplicas.

	UpgradeableComponents []*UpgradeVersionInfo `json:"UpgradeableComponents,omitempty" name:"UpgradeableComponents"`
	// 是否开启自愈能力

	AutoRepair *bool `json:"AutoRepair,omitempty" name:"AutoRepair"`
	// 包年包月机型计费配置

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 故障自愈规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// 预定义脚本

	Lifecycle *LifecycleConfig `json:"Lifecycle,omitempty" name:"Lifecycle"`
	// JoiningReplicas.

	ReadyReplicas *int64 `json:"ReadyReplicas,omitempty" name:"ReadyReplicas"`
	// JoiningReplicas.

	KubeletVersion *string `json:"KubeletVersion,omitempty" name:"KubeletVersion"`
	// OsName.

	OsName *string `json:"OsName,omitempty" name:"OsName"`
}

type StartMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StartMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceExtraArgs struct {

	// kubelet自定义参数，参数格式为["k1=v1",&nbsp;"k1=v2"]，&nbsp;例如["root-dir=/var/lib/kubelet","feature-gates=PodShareProcessNamespace=true,DynamicKubeletConfig=true"]

	Kubelet []*string `json:"Kubelet,omitempty" name:"Kubelet"`
}

type DescribeHealthCheckPolicyBindingsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 健康检测规则数组

		HealthCheckPolicyBindings []*HealthCheckPolicyBinding `json:"HealthCheckPolicyBindings,omitempty" name:"HealthCheckPolicyBindings"`
		// 健康检测规则数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHealthCheckPolicyBindingsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHealthCheckPolicyBindingsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManuallyAdded struct {

	// 节点总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 过期未回收的节点数量

	Shutdown *int64 `json:"Shutdown,omitempty" name:"Shutdown"`
	// 加入中的节点数量

	Joining *int64 `json:"Joining,omitempty" name:"Joining"`
	// 初始化中的节点数量

	Initializing *int64 `json:"Initializing,omitempty" name:"Initializing"`
	// 正常的节点数量

	Normal *int64 `json:"Normal,omitempty" name:"Normal"`
}

type DescribeClusterResourceLabelsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群资源标签

		Labels []*string `json:"Labels,omitempty" name:"Labels"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterResourceLabelsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResourceLabelsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHousekeeperRegionRequest struct {
	*tchttp.BaseRequest
}

func (r *ListHousekeeperRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHousekeeperRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHealthCheckPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHealthCheckPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHealthCheckPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyHealthCheckPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyHealthCheckPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHealthCheckPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Label struct {

	// map表中的Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// map表中的Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type UpgradeVersionInfo struct {

	// 可升级组件名称

	Component *string `json:"Component,omitempty" name:"Component"`
	// 可升级组件版本

	DesiredVersion *string `json:"DesiredVersion,omitempty" name:"DesiredVersion"`
}

type DescribeNodeParamUpdateProcessResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点池id

		NodePoolID *string `json:"NodePoolID,omitempty" name:"NodePoolID"`
		// 节点池期望节点数

		Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
		// 节点参数更新进度记录

		NodeParamUpdateProcesses []*NodeParamUpdateProcesses `json:"NodeParamUpdateProcesses,omitempty" name:"NodeParamUpdateProcesses"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodeParamUpdateProcessResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeParamUpdateProcessResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartMachineRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
}

func (r *RestartMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterResourceLabelsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// ·&nbsp;&nbsp;Node
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【节点】进行获取标签
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	//
	// ·&nbsp;&nbsp;NodePool
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【节点池】进行获取标签
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeClusterResourceLabelsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterResourceLabelsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportedInstanceFamiliesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否收取原生节点增值服务费用

		EnabledServicefee *bool `json:"EnabledServicefee,omitempty" name:"EnabledServicefee"`
		// 原生节点支持的实例族列表

		InstanceFamilies []*string `json:"InstanceFamilies,omitempty" name:"InstanceFamilies"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSupportedInstanceFamiliesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportedInstanceFamiliesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckPolicy struct {

	// 健康检测策略名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 健康检测策略规则列表

	Rules []*HealthCheckPolicyRule `json:"Rules,omitempty" name:"Rules"`
}

type InquirePricePurchaseNativeNodeRequest struct {
	*tchttp.BaseRequest

	// 机型创建数量配置

	InstanceGoodsInfos []*InstanceGoodsInfo `json:"InstanceGoodsInfos,omitempty" name:"InstanceGoodsInfos"`
	// 机型的可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 高性能集群机型需要传对应的高性能集群id

	HPCClusterId *string `json:"HPCClusterId,omitempty" name:"HPCClusterId"`
	// 按量计费：POSTPAID_BY_HOUR，包年包月：PREPAID。默认值：按量计费

	ChargeType *string `json:"ChargeType,omitempty" name:"ChargeType"`
	// 询价时长，按量计费：单位秒，包年包月：单位月。默认值按量计费：3600秒，包年包月：一个月。

	TimeSpan *int64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 系统盘类型，取值为：CLOUD_SSD、CLOUD_PREMIUM，默认值为高性能云盘：CLOUD_PREMIUM

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 系统盘大小，单位：Gi，默认值：50

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 数据盘列表

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
}

func (r *InquirePricePurchaseNativeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePricePurchaseNativeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetMachineScaleDownProtectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetMachineScaleDownProtectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMachineScaleDownProtectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AutoscalingAdded struct {

	// 正在加入中的节点数量

	Joining *int64 `json:"Joining,omitempty" name:"Joining"`
	// 初始化中的节点数量

	Initializing *int64 `json:"Initializing,omitempty" name:"Initializing"`
	// 正常的节点数量

	Normal *int64 `json:"Normal,omitempty" name:"Normal"`
	// 节点总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 过期未回收的节点数量

	Shutdown *int64 `json:"Shutdown,omitempty" name:"Shutdown"`
}

type DescribeHealthCheckTemplateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 健康检测策略模板

		HealthCheckTemplate *HealthCheckTemplate `json:"HealthCheckTemplate,omitempty" name:"HealthCheckTemplate"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHealthCheckTemplateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHealthCheckTemplateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NativeNodeRenewCostInfo struct {

	// 原生节点增值服务费询价结果，单位：分，折扣后

	HKServiceFeeTotalCost *float64 `json:"HKServiceFeeTotalCost,omitempty" name:"HKServiceFeeTotalCost"`
	// 资源&nbsp;ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 实例询价结果，单位：分，折扣前

	TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
	// 云盘询价结果，单位：分，打折后

	CBSCost *float64 `json:"CBSCost,omitempty" name:"CBSCost"`
	// cvm原生节点cvm机器询价结果，单位：分，折扣前

	CVMTotalCost *float64 `json:"CVMTotalCost,omitempty" name:"CVMTotalCost"`
	// 实例询价结果，单位：分，打折后

	Cost *float64 `json:"Cost,omitempty" name:"Cost"`
	// 云盘询价结果，单位：分，折扣前

	CBSTotalCost *float64 `json:"CBSTotalCost,omitempty" name:"CBSTotalCost"`
	// cvm原生节点cvm机器询价结果，单位：分，折扣后

	CVMCost *float64 `json:"CVMCost,omitempty" name:"CVMCost"`
	// 原生节点增值服务费询价结果，单位：分，折扣前

	HKServiceFeeCost *float64 `json:"HKServiceFeeCost,omitempty" name:"HKServiceFeeCost"`
}

type CUDNN struct {

	// cuDNN的版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// cuDNN的名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// cuDNN的Doc名字

	DocName *string `json:"DocName,omitempty" name:"DocName"`
	// cuDNN的Dev名字

	DevName *string `json:"DevName,omitempty" name:"DevName"`
}

type UpdateExternalNodePoolParam struct {

	// 节点相关自定义参数

	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
	// 运行时参数

	RuntimeConfig *RuntimeConfig `json:"RuntimeConfig,omitempty" name:"RuntimeConfig"`
	// 是否更新存量节点

	UpdateExistedNode *bool `json:"UpdateExistedNode,omitempty" name:"UpdateExistedNode"`
	// base64&nbsp;编码的用户脚本,&nbsp;此脚本会在节点加入集群后执行。

	UserScript *string `json:"UserScript,omitempty" name:"UserScript"`
	// Machine系统配置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 故障自愈规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
}

type SetMachineScaleDownProtectionRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名列表

	MachineNames []*string `json:"MachineNames,omitempty" name:"MachineNames"`
	// 是否开启缩容保护

	ScaleDownDisabled *bool `json:"ScaleDownDisabled,omitempty" name:"ScaleDownDisabled"`
}

func (r *SetMachineScaleDownProtectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMachineScaleDownProtectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IntOrString struct {

	// 整数

	IntVal *int64 `json:"IntVal,omitempty" name:"IntVal"`
	// 字符串

	StrVal *string `json:"StrVal,omitempty" name:"StrVal"`
	// 数值类型，0是int,&nbsp;&nbsp;1是字符串

	Type *int64 `json:"Type,omitempty" name:"Type"`
}

type ClusterInspectionTask struct {

	// 巡检编排

	Orchestration *string `json:"Orchestration,omitempty" name:"Orchestration"`
	// 执行时间

	ExecuteTime *string `json:"ExecuteTime,omitempty" name:"ExecuteTime"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 巡检建议

	Suggest *string `json:"Suggest,omitempty" name:"Suggest"`
	// 任务ID

	TaskID *string `json:"TaskID,omitempty" name:"TaskID"`
	// 任务类型

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 任务名称

	JobName *string `json:"JobName,omitempty" name:"JobName"`
	// waiting,&nbsp;executing,&nbsp;finished

	Status *string `json:"Status,omitempty" name:"Status"`
	// 结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 集群概览

	Cluster *ClusterInspectionTaskOverview `json:"Cluster,omitempty" name:"Cluster"`
}

type DescribeZoneInstanceConfigInfosRequest struct {
	*tchttp.BaseRequest

	// 机型过滤设置

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeZoneInstanceConfigInfosRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneInstanceConfigInfosRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceHousekeeperResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价结果，单位：分，打折后

		Cost *uint64 `json:"Cost,omitempty" name:"Cost"`
		// 询价结果，单位：分，折扣前

		TotalCost *uint64 `json:"TotalCost,omitempty" name:"TotalCost"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceHousekeeperResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceHousekeeperResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// ·&nbsp;&nbsp;ClusterName
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【集群名】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;ClusterType
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【集群类型】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;ClusterStatus
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【集群状态】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;Tags
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【标签键值对】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;vpc-id
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【VPC】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;tag-key
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【标签键】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;tag-value
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【标签值】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;tag:tag-key
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【标签键值对】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 集群类型，例如：MANAGED_CLUSTER

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最大输出条数，默认20，最大为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetMachineLoginResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetMachineLoginResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMachineLoginResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewNativeNodeRequest struct {
	*tchttp.BaseRequest

	// 资源&nbsp;ID&nbsp;列表。格式为&nbsp;cls-adeadbef_np-769jmpw6-lmqcb

	ResourceIds []*string `json:"ResourceIds,omitempty" name:"ResourceIds"`
	// 包年包月节点续费时长，单位：月。

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
}

func (r *SwitchParameterRenewNativeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRenewNativeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePriceRefundNativeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 退费结果列表

		Costs []*NativeNodeRefundCostInfo `json:"Costs,omitempty" name:"Costs"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceRefundNativeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceRefundNativeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeSupportedInstanceFamiliesRequest struct {
	*tchttp.BaseRequest

	// **实例的类型**
	// 支持两种类型：Navite&nbsp;和&nbsp;NativeCVM
	// **Native**:&nbsp;CXM&nbsp;模式，底层的实例是&nbsp;CXM&nbsp;（默认值）
	// **NativeCVM**:&nbsp;CVM&nbsp;模式，底层的实例是&nbsp;CVM

	NativeType *string `json:"NativeType,omitempty" name:"NativeType"`
}

func (r *DescribeSupportedInstanceFamiliesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSupportedInstanceFamiliesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterCreateNativeNodeRequest struct {
	*tchttp.BaseRequest

	// 机型配置列表

	MachineConfigurationSet []*MachineConfiguration `json:"MachineConfigurationSet,omitempty" name:"MachineConfigurationSet"`
	// 节点池&nbsp;ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 原生节点池创建模板

	NodePoolTemplate *NodePoolTemplate `json:"NodePoolTemplate,omitempty" name:"NodePoolTemplate"`
	// 节点池的属性

	NativeType *string `json:"NativeType,omitempty" name:"NativeType"`
	// 按量计费：POSTPAID_BY_HOUR，包年包月：PREPAID。默认值：按量计费

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 系统盘

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 数据盘列表

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 包年包月机型计费配置

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
}

func (r *SwitchParameterCreateNativeNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterCreateNativeNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstantInspectJobRequest struct {
	*tchttp.BaseRequest

	// 集群id,&nbsp;根据集群id初始化任务，如果不为空，则忽略下面的MatchLables字段

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群label

	MatchLabels []*Label `json:"MatchLabels,omitempty" name:"MatchLabels"`
}

func (r *CreateInstantInspectJobRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstantInspectJobRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckTemplate struct {

	// 健康检测项

	Rules []*HealthCheckTemplateRule `json:"Rules,omitempty" name:"Rules"`
}

type CustomDriver struct {

	// 自定义GPU驱动地址链接

	Address *string `json:"Address,omitempty" name:"Address"`
}

type CreateHostedNodePoolParam struct {

	// 系统盘配置

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 自动升级配置

	UpgradeSettings *MachineUpgradeSettings `json:"UpgradeSettings,omitempty" name:"UpgradeSettings"`
	// 是否开启自愈能力

	AutoRepair *bool `json:"AutoRepair,omitempty" name:"AutoRepair"`
	// 节点池&nbsp;Management&nbsp;参数设置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 机型列表

	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
	// 故障自愈规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// 运行时根目录

	RuntimeRootDir *string `json:"RuntimeRootDir,omitempty" name:"RuntimeRootDir"`
	// 是否开启弹性伸缩

	EnableAutoscaling *bool `json:"EnableAutoscaling,omitempty" name:"EnableAutoscaling"`
	// 期望节点数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 包年包月机型计费配置

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 节点池伸缩配置

	Scaling *MachineSetScaling `json:"Scaling,omitempty" name:"Scaling"`
	// 节点计费类型。PREPAID：包年包月；POSTPAID_BY_HOUR：按量计费（默认）；

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 安全组列表

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// kubelet&nbsp;自定义参数

	KubeletArgs []*string `json:"KubeletArgs,omitempty" name:"KubeletArgs"`
	// 预定义脚本

	Lifecycle *LifecycleConfig `json:"Lifecycle,omitempty" name:"Lifecycle"`
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeGPUInfoRequest struct {
	*tchttp.BaseRequest

	// 实例机型名称，默认值""

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 操作系统oskey，默认值""

	OsName *string `json:"OsName,omitempty" name:"OsName"`
}

func (r *DescribeGPUInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeGPUInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineUpgradeSettings struct {

	// 是否开启自动升级

	AutoUpgrade *bool `json:"AutoUpgrade,omitempty" name:"AutoUpgrade"`
	// 运维窗口

	UpgradeOptions *AutoUpgradeOptions `json:"UpgradeOptions,omitempty" name:"UpgradeOptions"`
	// 升级项

	Components []*string `json:"Components,omitempty" name:"Components"`
	// 升级时，最大不可升级的节点数

	MaxUnavailable *IntOrString `json:"MaxUnavailable,omitempty" name:"MaxUnavailable"`
}

type ClusterInspectionTaskOverview struct {

	// 节点数量

	Nodes *uint64 `json:"Nodes,omitempty" name:"Nodes"`
	// Pod数量

	Pods *uint64 `json:"Pods,omitempty" name:"Pods"`
	// CPU核数

	CPUCores *uint64 `json:"CPUCores,omitempty" name:"CPUCores"`
	// 内存GB数

	MemoryGB *uint64 `json:"MemoryGB,omitempty" name:"MemoryGB"`
	// Workload数量

	Workloads *uint64 `json:"Workloads,omitempty" name:"Workloads"`
	// 命名空间数量

	Namespaces *uint64 `json:"Namespaces,omitempty" name:"Namespaces"`
}

type ConvertRegularToNativeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConvertRegularToNativeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ConvertRegularToNativeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHealthCheckPolicyRequest struct {
	*tchttp.BaseRequest

	// 健康检测策略名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteHealthCheckPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHealthCheckPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 属性名称,&nbsp;若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值,&nbsp;若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type UpdateHostedNodePoolParam struct {

	// 包年包月机型计费配置

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 系统盘配置

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 机型列表

	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
	// 是否开启自愈能力

	AutoRepair *bool `json:"AutoRepair,omitempty" name:"AutoRepair"`
	// 故障自愈规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// kubelet&nbsp;自定义参数

	KubeletArgs []*string `json:"KubeletArgs,omitempty" name:"KubeletArgs"`
	// machine&nbsp;系统配置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 安全组列表

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 自动升级配置

	UpgradeSettings *MachineUpgradeSettings `json:"UpgradeSettings,omitempty" name:"UpgradeSettings"`
	// 预定义脚本

	Lifecycle *LifecycleConfig `json:"Lifecycle,omitempty" name:"Lifecycle"`
	// 运行时根目录

	RuntimeRootDir *string `json:"RuntimeRootDir,omitempty" name:"RuntimeRootDir"`
	// 是否开启弹性伸缩

	EnableAutoscaling *bool `json:"EnableAutoscaling,omitempty" name:"EnableAutoscaling"`
	// 是否更新存量节点

	UpdateExistedNode *bool `json:"UpdateExistedNode,omitempty" name:"UpdateExistedNode"`
	// 伸缩配置

	Scaling *MachineSetScaling `json:"Scaling,omitempty" name:"Scaling"`
	// 期望节点数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 节点计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
}

type RegularNodeInfo struct {

	// 节点配置

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 自动伸缩组ID

	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitempty" name:"AutoscalingGroupId"`
}

type StopMachinesRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名字列表，一次请求，传入节点数量上限为100个

	MachineNames []*string `json:"MachineNames,omitempty" name:"MachineNames"`
	// 实例的关闭模式。取值范围：
	// soft_first：表示在正常关闭失败后进行强制关闭
	// hard：直接强制关闭
	// soft：仅软关机

	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

func (r *StopMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GPUArgs struct {

	// 是否启用MIG特性

	MIGEnable *bool `json:"MIGEnable,omitempty" name:"MIGEnable"`
	// GPU驱动版本信息

	Driver *DriverVersion `json:"Driver,omitempty" name:"Driver"`
	// CUDA版本信息

	CUDA *DriverVersion `json:"CUDA,omitempty" name:"CUDA"`
	// cuDNN版本信息

	CUDNN *CUDNN `json:"CUDNN,omitempty" name:"CUDNN"`
	// 自定义GPU驱动信息

	CustomDriver *CustomDriver `json:"CustomDriver,omitempty" name:"CustomDriver"`
}

type InquirePriceRenewNativeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价结果列表

		Costs []*NativeNodeRenewCostInfo `json:"Costs,omitempty" name:"Costs"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceRenewNativeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceRenewNativeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群信息列表

		Clusters []*Cluster `json:"Clusters,omitempty" name:"Clusters"`
		// 错误信息集合

		Errors []*string `json:"Errors,omitempty" name:"Errors"`
		// 集群总个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
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

type ModifyHealthCheckPolicyRequest struct {
	*tchttp.BaseRequest

	// 健康检测策略

	HealthCheckPolicy *HealthCheckPolicy `json:"HealthCheckPolicy,omitempty" name:"HealthCheckPolicy"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ModifyHealthCheckPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyHealthCheckPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HealthCheckPolicyRule struct {

	// 健康检测规则

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否检测此项目

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 是否启用修复

	AutoRepairEnabled *bool `json:"AutoRepairEnabled,omitempty" name:"AutoRepairEnabled"`
}

type SuperNodePoolInfo struct {

	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 安全组列表

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type VerifyQGPUResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否支持QGPU

		Support *bool `json:"Support,omitempty" name:"Support"`
		// 详细信息

		Message *string `json:"Message,omitempty" name:"Message"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *VerifyQGPUResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *VerifyQGPUResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListHousekeeperRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域列表总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域详情

		Regions []*RegionInfo `json:"Regions,omitempty" name:"Regions"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListHousekeeperRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListHousekeeperRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考&nbsp;API&nbsp;[简介](https://{{conf.main_domain}}/document/api/213/15688)中的相关小节。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考&nbsp;API&nbsp;[简介](https://{{conf.main_domain}}/document/api/213/15688)中的相关小节。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件列表:
	// InstanceIds(实例ID),InstanceType(实例类型：Regular，Native，Super，External),VagueIpAddress(模糊匹配IP),Labels(k8s节点label),NodePoolNames(节点池名称),VagueInstanceName(模糊匹配节点名),InstanceStates(节点状态),Unschedulable(是否封锁),NodePoolIds(节点池ID)

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序信息

	SortBy *SortBy `json:"SortBy,omitempty" name:"SortBy"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterMachinesRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名列表

	MachineNames []*string `json:"MachineNames,omitempty" name:"MachineNames"`
	// 删除节点时是否缩容节点池，true为缩容

	EnableScaleDown *bool `json:"EnableScaleDown,omitempty" name:"EnableScaleDown"`
	// 删除原生节点的时候是否同时删除&nbsp;CVM&nbsp;的实例，仅对后付费节点生效

	ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *DeleteClusterMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodeParamUpdateProcessRequest struct {
	*tchttp.BaseRequest

	// 节点池所在集群id

	ClusterID *string `json:"ClusterID,omitempty" name:"ClusterID"`
	// 节点池id

	NodePoolID *string `json:"NodePoolID,omitempty" name:"NodePoolID"`
	// 查询的参数集合

	Params []*string `json:"Params,omitempty" name:"Params"`
}

func (r *DescribeNodeParamUpdateProcessRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodeParamUpdateProcessRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type MachineConfiguration struct {

	// 私有网络信息

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 机型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 机型的&nbsp;CPU&nbsp;配置

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 可用区&nbsp;ID

	ZoneId *string `json:"ZoneId,omitempty" name:"ZoneId"`
	// 机型的内存配置

	Mem *int64 `json:"Mem,omitempty" name:"Mem"`
	// 子网&nbsp;ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 机型需要创建的个数

	Num *int64 `json:"Num,omitempty" name:"Num"`
	// 机型&nbsp;GPU&nbsp;个数

	GPU *int64 `json:"GPU,omitempty" name:"GPU"`
}

type DeleteClusterMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ManagementConfig struct {

	// dns&nbsp;配置

	Nameservers []*string `json:"Nameservers,omitempty" name:"Nameservers"`
	// hosts&nbsp;配置

	Hosts []*string `json:"Hosts,omitempty" name:"Hosts"`
	// 内核参数配置

	KernelArgs []*string `json:"KernelArgs,omitempty" name:"KernelArgs"`
}

type InstanceGoodsInfo struct {

	// 机型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 创建个数

	GoodsNum *uint64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
}

type RestartMachineResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartMachineResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartMachineResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateExternalNodePoolParam struct {

	// base64编码的用户脚本，此脚本会在k8s组件ready后执行

	UserScript *string `json:"UserScript,omitempty" name:"UserScript"`
	// 运行时参数

	RuntimeConfig *RuntimeConfig `json:"RuntimeConfig,omitempty" name:"RuntimeConfig"`
	// 节点池Management参数设置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 节点健康检查规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// 注册节点类型如GPU、CPU

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 公网：Public&nbsp;专线：Direct

	ConnectType *string `json:"ConnectType,omitempty" name:"ConnectType"`
	// 节点相关自定义参数

	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
}

type CreateNodePoolRequest struct {
	*tchttp.BaseRequest

	// 节点&nbsp;Annotation&nbsp;列表

	Annotations []*Annotation `json:"Annotations,omitempty" name:"Annotations"`
	// 注册节点池创建参数

	External *CreateExternalNodePoolParam `json:"External,omitempty" name:"External"`
	// 节点池名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点标签

	Tags []*TagSpecification `json:"Tags,omitempty" name:"Tags"`
	// 节点&nbsp;&nbsp;Labels

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 节点污点

	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
	// 是否开启删除保护

	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
	// 原生节点池创建参数（已弃用）

	Hosted *CreateHostedNodePoolParam `json:"Hosted,omitempty" name:"Hosted"`
	// 节点是否默认不可调度

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 原生节点池创建参数

	Native *CreateNativeNodePoolParam `json:"Native,omitempty" name:"Native"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池类型

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *CreateNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNodePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyNodePoolRequest struct {
	*tchttp.BaseRequest

	// 节点标签

	Tags []*TagSpecification `json:"Tags,omitempty" name:"Tags"`
	// 是否开启删除保护

	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
	// 原生节点池更新参数（已弃用）

	Hosted *UpdateHostedNodePoolParam `json:"Hosted,omitempty" name:"Hosted"`
	// 节点是否不可调度

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 原生节点池更新参数

	Native *UpdateNativeNodePoolParam `json:"Native,omitempty" name:"Native"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点污点

	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
	// 节点&nbsp;Annotation&nbsp;列表

	Annotations []*Annotation `json:"Annotations,omitempty" name:"Annotations"`
	// 节点池&nbsp;ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 节点&nbsp;&nbsp;Labels

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 注册节点更新参数

	External *UpdateExternalNodePoolParam `json:"External,omitempty" name:"External"`
}

func (r *ModifyNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyNodePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NativeNodeRefundCostInfo struct {

	// cvm原生节点cvm机器询价结果，单位：分

	CVMCost *float64 `json:"CVMCost,omitempty" name:"CVMCost"`
	// 原生节点增值服务费询价结果，单位：分

	HKServiceFeeCost *float64 `json:"HKServiceFeeCost,omitempty" name:"HKServiceFeeCost"`
	// 资源&nbsp;ID

	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
	// 实例询价结果，单位：分

	Cost *float64 `json:"Cost,omitempty" name:"Cost"`
	// 云盘询价结果，单位：分

	CBSCost *float64 `json:"CBSCost,omitempty" name:"CBSCost"`
}

type CreateInstantInspectJobResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 巡检任务名称

		JobName *string `json:"JobName,omitempty" name:"JobName"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstantInspectJobResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstantInspectJobResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateNativeNodePoolParam struct {

	// 机型列表

	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
	// 节点安全组存量更新开关，有enable（打开）、disable（关闭）两个状态可选

	UpdateSecurityGroups *string `json:"UpdateSecurityGroups,omitempty" name:"UpdateSecurityGroups"`
	// 原生节点池安装自动化助手开关状态

	AutomationService *bool `json:"AutomationService,omitempty" name:"AutomationService"`
	// 公网信息

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 节点移出策略，有Random（随机）、Newest（优先移出最新实例）、Oldest（优先移出最旧实例）三种可选，默认是Newest

	DeletePolicy *string `json:"DeletePolicy,omitempty" name:"DeletePolicy"`
	// 自动升级配置

	UpgradeSettings *MachineUpgradeSettings `json:"UpgradeSettings,omitempty" name:"UpgradeSettings"`
	// kubelet&nbsp;自定义参数

	KubeletArgs []*string `json:"KubeletArgs,omitempty" name:"KubeletArgs"`
	// 是否开启弹性伸缩

	EnableAutoscaling *bool `json:"EnableAutoscaling,omitempty" name:"EnableAutoscaling"`
	// 数据盘列表

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// qgpu&nbsp;组件版本

	QGPUVersion *string `json:"QGPUVersion,omitempty" name:"QGPUVersion"`
	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// ssh公钥id

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 安全组是否传递到存量节点

	SecurityGroupsSpread *bool `json:"SecurityGroupsSpread,omitempty" name:"SecurityGroupsSpread"`
	// 节点存量更新过程中最大更新不成功节点容忍数，当前值仅支持数字，后续支持百分比

	RollingUpdateMaxConfigUnavailable *string `json:"RollingUpdateMaxConfigUnavailable,omitempty" name:"RollingUpdateMaxConfigUnavailable"`
	// qgpu开关

	QGPUEnable *bool `json:"QGPUEnable,omitempty" name:"QGPUEnable"`
	// ssh公钥id数组

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 节点management参数存量更新开关，有enable（打开）、disable（关闭）两个状态可选

	UpdateMachineManagement *string `json:"UpdateMachineManagement,omitempty" name:"UpdateMachineManagement"`
	// 节点计费类型变更
	// 当前仅支持按量计费转包年包月：
	// -&nbsp;PREPAID
	//

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// Machine&nbsp;系统配置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 故障自愈规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// 原生节点池hostName模式串

	HostNamePattern *string `json:"HostNamePattern,omitempty" name:"HostNamePattern"`
	// 是否更新存量节点

	UpdateExistedNode *bool `json:"UpdateExistedNode,omitempty" name:"UpdateExistedNode"`
	// 节点anno/label/taint参数存量更新开关，有enable（打开）、disable（关闭）两个状态可选

	UpdateMachineMetadata *string `json:"UpdateMachineMetadata,omitempty" name:"UpdateMachineMetadata"`
	// 是否开启自愈能力

	AutoRepair *bool `json:"AutoRepair,omitempty" name:"AutoRepair"`
	// 包年包月机型计费配置

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 系统盘配置

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 预定义脚本

	Lifecycle *LifecycleConfig `json:"Lifecycle,omitempty" name:"Lifecycle"`
	// 运行时根目录

	RuntimeRootDir *string `json:"RuntimeRootDir,omitempty" name:"RuntimeRootDir"`
	// 期望节点数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 伸缩配置

	Scaling *MachineSetScaling `json:"Scaling,omitempty" name:"Scaling"`
	// 安全组列表

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 节点存量更新过程中一次更新最大步长，当前值仅支持数字，后续支持百分比

	RollingUpdateMaxSteps *string `json:"RollingUpdateMaxSteps,omitempty" name:"RollingUpdateMaxSteps"`
	// 节点池&nbsp;GPU&nbsp;配置

	GPUConfigs []*GPUConfig `json:"GPUConfigs,omitempty" name:"GPUConfigs"`
}

type DataDisk struct {

	// 云盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 文件系统(ext3/ext4/xfs)

	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`
	// 挂载设备名或分区名

	DiskPartition *string `json:"DiskPartition,omitempty" name:"DiskPartition"`
	// 购买加密盘时自定义密钥，当传入该参数时,&nbsp;Encrypt入参不为空

	KmsKeyId *string `json:"KmsKeyId,omitempty" name:"KmsKeyId"`
	// 镜像缓存&nbsp;id

	ImageCacheId *string `json:"ImageCacheId,omitempty" name:"ImageCacheId"`
	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照

	SnapshotId *string `json:"SnapshotId,omitempty" name:"SnapshotId"`
	// 云硬盘性能，单位：MB/s。使用此参数可给云硬盘购买额外的性能

	ThroughputPerformance *uint64 `json:"ThroughputPerformance,omitempty" name:"ThroughputPerformance"`
	// 云盘大小(G）

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 是否自动化格式盘并挂载

	AutoFormatAndMount *bool `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	// 磁盘id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// 挂载目录

	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`
	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT

	Encrypt *string `json:"Encrypt,omitempty" name:"Encrypt"`
}

type CreateNativeNodePoolParam struct {

	// 预定义脚本

	Lifecycle *LifecycleConfig `json:"Lifecycle,omitempty" name:"Lifecycle"`
	// 机型和GPU配置相关信息

	GPUConfigs []*GPUConfig `json:"GPUConfigs,omitempty" name:"GPUConfigs"`
	// qgpu开关

	QGPUEnable *bool `json:"QGPUEnable,omitempty" name:"QGPUEnable"`
	// 高性能集群ID

	HPCClusterID *string `json:"HPCClusterID,omitempty" name:"HPCClusterID"`
	// 节点池ssh公钥id数组

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 期望节点数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 原生节点池安装节点自动化助手开关

	AutomationService *bool `json:"AutomationService,omitempty" name:"AutomationService"`
	// 机型列表

	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
	// kubelet&nbsp;自定义参数

	KubeletArgs []*string `json:"KubeletArgs,omitempty" name:"KubeletArgs"`
	// 运行时根目录

	RuntimeRootDir *string `json:"RuntimeRootDir,omitempty" name:"RuntimeRootDir"`
	// qgpu组件版本

	QGPUVersion *string `json:"QGPUVersion,omitempty" name:"QGPUVersion"`
	// 安全组列表

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 包年包月机型计费配置

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// 节点池&nbsp;Management&nbsp;参数设置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 故障自愈规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// 原生节点池数据盘列表

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 自动升级配置

	UpgradeSettings *MachineUpgradeSettings `json:"UpgradeSettings,omitempty" name:"UpgradeSettings"`
	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 节点计费类型。PREPAID：包年包月；POSTPAID_BY_HOUR：按量计费（默认）；

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 是否开启弹性伸缩

	EnableAutoscaling *bool `json:"EnableAutoscaling,omitempty" name:"EnableAutoscaling"`
	// 公网带宽设置

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 节点池ssh公钥id

	KeyId *string `json:"KeyId,omitempty" name:"KeyId"`
	// 节点移出策略，有Random（随机）、Newest（优先移出最新实例）、Oldest（优先移出最旧实例）三种可选，默认是Newest

	DeletePolicy *string `json:"DeletePolicy,omitempty" name:"DeletePolicy"`
	// 节点池类型

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 节点池伸缩配置

	Scaling *MachineSetScaling `json:"Scaling,omitempty" name:"Scaling"`
	// 是否开启自愈能力

	AutoRepair *bool `json:"AutoRepair,omitempty" name:"AutoRepair"`
	// 原生节点池hostName模式串

	HostNamePattern *string `json:"HostNamePattern,omitempty" name:"HostNamePattern"`
	// 系统盘配置

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
}

type NodePoolTemplate struct {

	// 节点是否默认不可调度

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 原生节点池创建参数

	Native *CreateNativeNodePoolParam `json:"Native,omitempty" name:"Native"`
	// 节点标签

	Tags []*TagSpecification `json:"Tags,omitempty" name:"Tags"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 是否开启删除保护

	DeletionProtection *bool `json:"DeletionProtection,omitempty" name:"DeletionProtection"`
	// 节点&nbsp;&nbsp;Labels

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 节点污点

	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
	// 节点池类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点池名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点池标注

	Annotations []*Annotation `json:"Annotations,omitempty" name:"Annotations"`
}

type DescribeHealthCheckTemplateRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeHealthCheckTemplateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHealthCheckTemplateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceChargePrepaid struct {

	// 后付费计费周期，单位（月）：
	// 1，2，3，4，5，，6，7，&nbsp;8，9，10，11，12，24，36，48，60

	Period *uint64 `json:"Period,omitempty" name:"Period"`
	// 预付费续费方式：
	// -&nbsp;NOTIFY_AND_AUTO_RENEW：通知用户过期，且自动续费&nbsp;(默认）
	// -&nbsp;NOTIFY_AND_MANUAL_RENEW：通知用户过期，但不自动续费
	// -&nbsp;DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知用户过期，也不自动续费
	//

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
}

type RebootMachinesRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名字列表，一次请求，传入节点数量上限为100个

	MachineNames []*string `json:"MachineNames,omitempty" name:"MachineNames"`
	// 实例的关闭模式。取值范围：
	// soft_first：表示在正常关闭失败后进行强制关闭
	// hard：直接强制关闭
	// soft：仅软关机默认取值：soft。

	StopType *string `json:"StopType,omitempty" name:"StopType"`
}

func (r *RebootMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RebootMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExternalNodeInfo struct {

	// 第三方节点名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// CPU核数，单位：核

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 节点内存容量，单位：`GB`

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 第三方节点kubelet版本信息

	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`
}

type DescribeMachineConfigurationResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机型配置列表

		MachineConfigurationSet []*MachineConfiguration `json:"MachineConfigurationSet,omitempty" name:"MachineConfigurationSet"`
		// 节点池&nbsp;ID

		NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMachineConfigurationResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeMachineConfigurationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DriverVersion struct {

	// GPU驱动或者CUDA的版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// GPU驱动或者CUDA的名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

type DescribeOperationEventsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 无

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 无

		Events []*OperationEvent `json:"Events,omitempty" name:"Events"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeOperationEventsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeOperationEventsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NativeNodePoolInfo struct {

	// 就绪&nbsp;Machine&nbsp;个数

	ReadyReplicas *int64 `json:"ReadyReplicas,omitempty" name:"ReadyReplicas"`
	// 节点&nbsp;kubelet&nbsp;版本

	KubeletVersion *string `json:"KubeletVersion,omitempty" name:"KubeletVersion"`
	// 修改最大可允许参数更新未成功节点数，当前值仅支持数字，后续支持百分比

	RollingUpdateMaxConfigUnavailable *string `json:"RollingUpdateMaxConfigUnavailable,omitempty" name:"RollingUpdateMaxConfigUnavailable"`
	// 修改存量节点参数安全组开关，包括两种状态，enable(开关打开)、disable(开关关闭)

	UpdateSecurityGroups *string `json:"UpdateSecurityGroups,omitempty" name:"UpdateSecurityGroups"`
	// 原生节点机型&nbsp;Native,&nbsp;NativeCVM

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 伸缩配置

	Scaling *MachineSetScaling `json:"Scaling,omitempty" name:"Scaling"`
	// 是否开启自愈能力

	AutoRepair *bool `json:"AutoRepair,omitempty" name:"AutoRepair"`
	// 可升级的组件列表

	UpgradeableComponents []*UpgradeVersionInfo `json:"UpgradeableComponents,omitempty" name:"UpgradeableComponents"`
	// 修改最大步长，当前值仅支持数字，后续支持百分比

	RollingUpdateMaxSteps *string `json:"RollingUpdateMaxSteps,omitempty" name:"RollingUpdateMaxSteps"`
	// 安全组列表

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 包年包月机型计费配置

	InstanceChargePrepaid *InstanceChargePrepaid `json:"InstanceChargePrepaid,omitempty" name:"InstanceChargePrepaid"`
	// Machine&nbsp;系统配置

	Management *ManagementConfig `json:"Management,omitempty" name:"Management"`
	// 运行时根目录

	RuntimeRootDir *string `json:"RuntimeRootDir,omitempty" name:"RuntimeRootDir"`
	// 是否开启弹性伸缩

	EnableAutoscaling *bool `json:"EnableAutoscaling,omitempty" name:"EnableAutoscaling"`
	// 原生节点池GPU相关配置

	GPUConfigs []*GPUConfig `json:"GPUConfigs,omitempty" name:"GPUConfigs"`
	// qgpu开关状态

	QGPUEnable *bool `json:"QGPUEnable,omitempty" name:"QGPUEnable"`
	// 节点计费类型

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// kubelet&nbsp;自定义参数

	KubeletArgs []*string `json:"KubeletArgs,omitempty" name:"KubeletArgs"`
	// 机型列表

	InstanceTypes []*string `json:"InstanceTypes,omitempty" name:"InstanceTypes"`
	// 正在加入节点池的&nbsp;Machine&nbsp;个数

	JoiningReplicas *int64 `json:"JoiningReplicas,omitempty" name:"JoiningReplicas"`
	// 修改存量节点参数anno/label/taints开关，包括两种状态，enable(开关打开)、disable(开关关闭)

	UpdateMachineMetadata *string `json:"UpdateMachineMetadata,omitempty" name:"UpdateMachineMetadata"`
	// 高性能集群id

	HPCClusterID *string `json:"HPCClusterID,omitempty" name:"HPCClusterID"`
	// 预定义脚本

	Lifecycle *LifecycleConfig `json:"Lifecycle,omitempty" name:"Lifecycle"`
	// 期望节点数

	Replicas *int64 `json:"Replicas,omitempty" name:"Replicas"`
	// 公网带宽设置

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 原生节点池数据盘

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 节点移出策略

	DeletePolicy *string `json:"DeletePolicy,omitempty" name:"DeletePolicy"`
	// 自动升级配置

	UpgradeSettings *MachineUpgradeSettings `json:"UpgradeSettings,omitempty" name:"UpgradeSettings"`
	// 密钥&nbsp;ID&nbsp;列表

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 故障自愈规则名称

	HealthCheckPolicyName *string `json:"HealthCheckPolicyName,omitempty" name:"HealthCheckPolicyName"`
	// 原生节点池hostName模式串

	HostNamePattern *string `json:"HostNamePattern,omitempty" name:"HostNamePattern"`
	// 是否更新存量节点

	UpdateExistedNode *bool `json:"UpdateExistedNode,omitempty" name:"UpdateExistedNode"`
	// 修改存量节点参数management开关，包括两种状态，enable(开关打开)、disable(开关关闭)

	UpdateMachineManagement *string `json:"UpdateMachineManagement,omitempty" name:"UpdateMachineManagement"`
	// 原生节点池安装节点自动化助手开关

	AutomationService *bool `json:"AutomationService,omitempty" name:"AutomationService"`
	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 系统盘配置

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 节点&nbsp;runtime&nbsp;版本

	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`
	// 镜像名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
}

type Taint struct {

	// Taint的Effect

	Effect *string `json:"Effect,omitempty" name:"Effect"`
	// Taint的Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Taint的Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeHealthCheckPoliciesRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// ·&nbsp;&nbsp;HealthCheckPolicyName
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【健康检测策略名称】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 最大输出条数，默认20，最大为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeHealthCheckPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHealthCheckPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Machine struct {

	// 包年包月节点计费过期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// VPC&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 默认登录用户名

	DefaultLoginUser *string `json:"DefaultLoginUser,omitempty" name:"DefaultLoginUser"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// CPU核数，单位：核

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 节点计费模式（已弃用）

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 节点内存容量，单位：`GB`

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 公网带宽相关信息设置

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 机型所属机型族

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 默认登录端口

	DefaultLoginPort *uint64 `json:"DefaultLoginPort,omitempty" name:"DefaultLoginPort"`
	// Machine&nbsp;所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 是否开启缩容保护

	IsProtectedFromScaleIn *bool `json:"IsProtectedFromScaleIn,omitempty" name:"IsProtectedFromScaleIn"`
	// Machine&nbsp;名字

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 自动续费标识

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// GPU核数，单位：核

	GPU *uint64 `json:"GPU,omitempty" name:"GPU"`
	// 节点系统盘配置信息

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 机型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 节点内网&nbsp;IP

	LanIP *string `json:"LanIP,omitempty" name:"LanIP"`
	// 原生节点公网&nbsp;IP

	WanIP *string `json:"WanIP,omitempty" name:"WanIP"`
	// cxm&nbsp;实例&nbsp;id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 节点名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// Machine&nbsp;状态

	MachineState *string `json:"MachineState,omitempty" name:"MachineState"`
	// 节点计费类型。PREPAID：包年包月；POSTPAID_BY_HOUR：按量计费（默认）；

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// Machine&nbsp;登录状态

	LoginStatus *string `json:"LoginStatus,omitempty" name:"LoginStatus"`
}

type ModifyMachinesKeyIdsRequest struct {
	*tchttp.BaseRequest

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 原生节点id

	MachineNames []*string `json:"MachineNames,omitempty" name:"MachineNames"`
	// ssh公钥id

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *ModifyMachinesKeyIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyMachinesKeyIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群中实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群中实例列表

		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 错误信息集合

		Errors []*string `json:"Errors,omitempty" name:"Errors"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NativeNodeInfo struct {

	// Machine&nbsp;状态

	MachineState *string `json:"MachineState,omitempty" name:"MachineState"`
	// VPC&nbsp;唯一&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 节点系统盘配置信息

	SystemDisk *Disk `json:"SystemDisk,omitempty" name:"SystemDisk"`
	// 机型所属机型族

	InstanceFamily *string `json:"InstanceFamily,omitempty" name:"InstanceFamily"`
	// 节点外网&nbsp;IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// **原生节点的&nbsp;Machine&nbsp;类型**
	//
	// -&nbsp;Native&nbsp;表示&nbsp;CXM&nbsp;类型的原生节点
	// -&nbsp;NativeCVM&nbsp;表示&nbsp;CVM&nbsp;类型的原生节点

	MachineType *string `json:"MachineType,omitempty" name:"MachineType"`
	// 自动续费标识

	RenewFlag *string `json:"RenewFlag,omitempty" name:"RenewFlag"`
	// 节点计费模式（已弃用）

	PayMode *string `json:"PayMode,omitempty" name:"PayMode"`
	// 节点内网&nbsp;IP

	LanIp *string `json:"LanIp,omitempty" name:"LanIp"`
	// **原生节点对应的实例&nbsp;ID**
	//
	// -&nbsp;ins-q47ofw6&nbsp;表示这个实例是一个&nbsp;CVM&nbsp;的实例
	// -&nbsp;eks-f8mvyaep&nbsp;表示这个实例是一个&nbsp;CXM&nbsp;的实例

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// Machine&nbsp;登录状态

	LoginStatus *string `json:"LoginStatus,omitempty" name:"LoginStatus"`
	// 机型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 子网唯一&nbsp;ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 公网带宽相关信息设置

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
	// 节点密钥&nbsp;ID&nbsp;列表

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// Machine&nbsp;所在可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 节点计费类型。PREPAID：包年包月；POSTPAID_BY_HOUR：按量计费（默认）；

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// GPU核数，单位：核

	GPU *uint64 `json:"GPU,omitempty" name:"GPU"`
	// 节点GPU相关配置

	GPUParams *GPUParams `json:"GPUParams,omitempty" name:"GPUParams"`
	// 安全组列表

	SecurityGroupIDs []*string `json:"SecurityGroupIDs,omitempty" name:"SecurityGroupIDs"`
	// 节点名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// CPU核数，单位：核

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 包年包月节点计费过期时间

	ExpiredTime *string `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
	// OS的名称

	OsImage *string `json:"OsImage,omitempty" name:"OsImage"`
	// 数据盘列表

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 是否开启缩容保护

	IsProtectedFromScaleIn *bool `json:"IsProtectedFromScaleIn,omitempty" name:"IsProtectedFromScaleIn"`
	// Machine&nbsp;名字

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 节点内存容量，单位：`GB`

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
}

type InquirePriceConvertNativeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价结果列表

		Costs []*NativeNodeConvertCostInfo `json:"Costs,omitempty" name:"Costs"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePriceConvertNativeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePriceConvertNativeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RuntimeConfig struct {

	// 运行时类型

	RuntimeType *string `json:"RuntimeType,omitempty" name:"RuntimeType"`
	// 运行时版本

	RuntimeVersion *string `json:"RuntimeVersion,omitempty" name:"RuntimeVersion"`
	// 运行时根目录

	RuntimeRootDir *string `json:"RuntimeRootDir,omitempty" name:"RuntimeRootDir"`
}

type SetMachineLoginRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名称

	MachineName *string `json:"MachineName,omitempty" name:"MachineName"`
	// 密钥&nbsp;ID&nbsp;列表

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
}

func (r *SetMachineLoginRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetMachineLoginRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNodePoolRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池&nbsp;ID

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
}

func (r *DeleteNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SuperNodeInfo struct {

	// 实例名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 自动续费标识

	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`
	// 节点上&nbsp;Pod&nbsp;的内存总和，单位：Gi。

	UsedMemory *float64 `json:"UsedMemory,omitempty" name:"UsedMemory"`
	// 过期时间

	ExpireAt *string `json:"ExpireAt,omitempty" name:"ExpireAt"`
	// 实例属性

	InstanceAttribute *string `json:"InstanceAttribute,omitempty" name:"InstanceAttribute"`
	// 节点的&nbsp;CPU&nbsp;规格，单位：核。

	CPU *float64 `json:"CPU,omitempty" name:"CPU"`
	// VPC&nbsp;唯一&nbsp;ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网唯一&nbsp;ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 生效时间

	ActiveAt *string `json:"ActiveAt,omitempty" name:"ActiveAt"`
	// 资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 节点上&nbsp;Pod&nbsp;的&nbsp;CPU总和，单位：核。

	UsedCPU *float64 `json:"UsedCPU,omitempty" name:"UsedCPU"`
	// 节点的内存规格，单位：Gi。

	Memory *float64 `json:"Memory,omitempty" name:"Memory"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可调度的单&nbsp;Pod&nbsp;最大&nbsp;CPU&nbsp;规格

	MaxCPUScheduledPod *int64 `json:"MaxCPUScheduledPod,omitempty" name:"MaxCPUScheduledPod"`
}

type StopMachinesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMachinesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *StopMachinesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeNodePoolsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点池列表

		NodePools []*NodePool `json:"NodePools,omitempty" name:"NodePools"`
		// 资源总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 错误信息

		Errors []*string `json:"Errors,omitempty" name:"Errors"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeNodePoolsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeNodePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeZoneInstanceConfigInfosResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 机型配置列表

		InstanceTypeQuotaSet *string `json:"InstanceTypeQuotaSet,omitempty" name:"InstanceTypeQuotaSet"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeZoneInstanceConfigInfosResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeZoneInstanceConfigInfosResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstantInspectTaskResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 巡检任务结果

		Data []*ClusterInspectionTask `json:"Data,omitempty" name:"Data"`
		// 巡检任务状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstantInspectTaskResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstantInspectTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InquirePricePurchaseNativeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 询价结果，单位：分，打折后

		Cost *float64 `json:"Cost,omitempty" name:"Cost"`
		// 询价结果，单位：分，折扣前

		TotalCost *float64 `json:"TotalCost,omitempty" name:"TotalCost"`
		// 云盘询价结果，单位：分，打折后

		CBSCost *float64 `json:"CBSCost,omitempty" name:"CBSCost"`
		// 云盘询价结果，单位：分，折扣前

		CBSTotalCost *float64 `json:"CBSTotalCost,omitempty" name:"CBSTotalCost"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InquirePricePurchaseNativeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InquirePricePurchaseNativeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateNodePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点池&nbsp;ID

		NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateNodePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterMachinesRequest struct {
	*tchttp.BaseRequest

	// 最大输出条数，默认20，最大为100

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点过滤条件，支持以下过滤条件：
	// ·&nbsp;&nbsp;NodePoolsName
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【节点池名】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;NodePoolsId
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【节点池id】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;tags
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【标签键值对】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否
	//
	// ·&nbsp;&nbsp;tag:tag-key
	// &nbsp;&nbsp;&nbsp;&nbsp;按照【标签键值对】进行过滤。
	// &nbsp;&nbsp;&nbsp;&nbsp;类型：String
	// &nbsp;&nbsp;&nbsp;&nbsp;必选：否

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 偏移量，默认0

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterMachinesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterMachinesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterMachineRequest struct {
	*tchttp.BaseRequest

	// 集群&nbsp;ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名列表

	MachineNames []*string `json:"MachineNames,omitempty" name:"MachineNames"`
	// machine的display&nbsp;name

	DisplayName *string `json:"DisplayName,omitempty" name:"DisplayName"`
	// 公网信息

	InternetAccessible *InternetAccessible `json:"InternetAccessible,omitempty" name:"InternetAccessible"`
}

func (r *ModifyClusterMachineRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterMachineRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteNodePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteNodePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchParameterRenewNativeNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下单参数&nbsp;JSON&nbsp;串

		Data *string `json:"Data,omitempty" name:"Data"`
		// Unique request ID, returned on each request. You need to provide this request when locating a problem RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SwitchParameterRenewNativeNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SwitchParameterRenewNativeNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// 地域简称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 地域

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域&nbsp;ID

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
}

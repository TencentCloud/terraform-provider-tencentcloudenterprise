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

package v20180525

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ModifyAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// k8s集群ID

	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`
	// k8s命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// k8s工作负载类型，有deplogyment，job，crontab

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 告警策略ID

	AlarmPolicyId *string `json:"AlarmPolicyId,omitempty" name:"AlarmPolicyId"`
	// 告警策略settings

	AlarmPolicySettings *AlarmPolicySettings `json:"AlarmPolicySettings,omitempty" name:"AlarmPolicySettings"`
	// 告警通知settings

	NotifySettings *NotifySettings `json:"NotifySettings,omitempty" name:"NotifySettings"`
	// 告警屏蔽settings

	ShieldSettings *ShieldSettings `json:"ShieldSettings,omitempty" name:"ShieldSettings"`
}

func (r *ModifyAlarmPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunSecurityServiceEnabled struct {

	// 是否开启[云安全](/tcloud/Security/CWP)服务。取值范围：<br><li>TRUE：表示开启云安全服务<br><li>FALSE：表示不开启云安全服务<br><br>默认取值：TRUE。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type DeleteClusterCIDRFromVbcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 详情

		Detail []*CcnRoute `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterCIDRFromVbcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterCIDRFromVbcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 镜像数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 镜像信息列表

		ImageInstanceSet []*ImageInstance `json:"ImageInstanceSet,omitempty" name:"ImageInstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeImagesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterInspectionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterInspectionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterInspectionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateClusterVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点实例id

		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PauseClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type K8SVersions struct {

	// 组件名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 组件版本

	Version *string `json:"Version,omitempty" name:"Version"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 描述

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type RunClusterInspectionResponseItem struct {

	// 集群实例id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 如果请求未能正常被处理，则Error中将包含错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
}

type DescribeClustersRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表(为空时，
	// 表示获取账号下所有集群)

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最大输出条数，默认20，最大为100

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件,当前只支持按照单个条件ClusterName进行过滤

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHelmChartVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// Results

		Results []*Result `json:"Results,omitempty" name:"Results"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHelmChartVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHelmChartVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterVersionRequest struct {
	*tchttp.BaseRequest

	// 集群 Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 需要升级到的版本

	DstVersion *string `json:"DstVersion,omitempty" name:"DstVersion"`
	// 集群自定义参数

	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
	// 可容忍的最大不可用pod数目

	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitempty" name:"MaxNotReadyPercent"`
	// 是否跳过预检查阶段

	SkipPreCheck *bool `json:"SkipPreCheck,omitempty" name:"SkipPreCheck"`
}

func (r *UpdateClusterVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddClusterCIDRToVbcRequest struct {
	*tchttp.BaseRequest

	// vpc唯一id

	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
	// 集群cidr

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *AddClusterCIDRToVbcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterCIDRToVbcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHelmChartVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否删除成功

		Deleted *bool `json:"Deleted,omitempty" name:"Deleted"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHelmChartVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHelmChartVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterExtraArgsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterExtraArgsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterExtraArgsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterHealthyStatusRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterHealthyStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterHealthyStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVirtualNodeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池id

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// NodeNames数组

	NodeNames []*string `json:"NodeNames,omitempty" name:"NodeNames"`
}

func (r *DescribeClusterVirtualNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVirtualNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVirtualNodePoolsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterVirtualNodePoolsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVirtualNodePoolsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterKubeconfigRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 需要更新证书的用户uin数组

	SubAccounts []*string `json:"SubAccounts,omitempty" name:"SubAccounts"`
}

func (r *UpdateClusterKubeconfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterKubeconfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ServiceAccountVolume struct {

	// data

	Data []*KeyValueData `json:"Data,omitempty" name:"Data"`
}

type ChartVersion struct {

	// Metadata

	Metadata *Metadata `json:"Metadata,omitempty" name:"Metadata"`
	// 创建时间

	Created *string `json:"Created,omitempty" name:"Created"`
	// Digest

	Digest *string `json:"Digest,omitempty" name:"Digest"`
	// URLs

	URLs []*string `json:"URLs,omitempty" name:"URLs"`
	// api

	APIVersion *string `json:"APIVersion,omitempty" name:"APIVersion"`
	// app

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// Description

	Description *string `json:"Description,omitempty" name:"Description"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// Version

	Version *string `json:"Version,omitempty" name:"Version"`
	// home

	Home *string `json:"Home,omitempty" name:"Home"`
	// icon

	Icon *string `json:"Icon,omitempty" name:"Icon"`
	// keywords

	Keywords []*string `json:"Keywords,omitempty" name:"Keywords"`
	// Deprecated

	Deprecated *bool `json:"Deprecated,omitempty" name:"Deprecated"`
	// KubeVersion

	KubeVersion *string `json:"KubeVersion,omitempty" name:"KubeVersion"`
	// Sources

	Sources []*string `json:"Sources,omitempty" name:"Sources"`
}

type NodePoolSet struct {

	// 节点池Id

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 节点池子网Id

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 节点池安全组id

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 节点池名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点池状态

	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
	// 不可调度

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 节点池标签

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 节点池污点

	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
}

type CreateClusterVirtualNodePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟节点池ID

		NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterVirtualNodePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterVirtualNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterEndpointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpUpgradeClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *OpUpgradeClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpUpgradeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCommon struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 集群状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群CIDR

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
	// 集群创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 集群中节点数

	NodeNum *int64 `json:"NodeNum,omitempty" name:"NodeNum"`
	// 集群中master节点数量

	MasterNum *int64 `json:"MasterNum,omitempty" name:"MasterNum"`
	// 操作系统名称

	Os []*string `json:"Os,omitempty" name:"Os"`
	// K8S集群版本

	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
	// 集群访问地址

	ClusterExternalEndpoint *string `json:"ClusterExternalEndpoint,omitempty" name:"ClusterExternalEndpoint"`
	// 集群最大支持的Pod数量

	MaxNodePodNum *int64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`
	// 集群最大支持的服务数量

	MaxClusterServiceNum *int64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`
	// ipvs

	IPVS *int64 `json:"IPVS,omitempty" name:"IPVS"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

type ShieldSettings struct {

	// 是否启动告警屏蔽功能

	EnableShield *bool `json:"EnableShield,omitempty" name:"EnableShield"`
	// 告警屏蔽开始时间，单位s，如8:00:00=> 8 * 60 * 60=28800

	ShieldTimeStart *int64 `json:"ShieldTimeStart,omitempty" name:"ShieldTimeStart"`
	// 告警屏蔽结束时间，单位s，如10:00:00=> 10 * 60 * 60=36000

	ShieldTimeEnd *int64 `json:"ShieldTimeEnd,omitempty" name:"ShieldTimeEnd"`
}

type DescribeAlarmPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群告警策略总个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群告警列表

		AlarmPolicySet []*AlarmPolicy `json:"AlarmPolicySet,omitempty" name:"AlarmPolicySet"`
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

type EKSPodInfo struct {

	// EKSId

	EKSId *string `json:"EKSId,omitempty" name:"EKSId"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// pod的namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 工作负载名称

	KindName *string `json:"KindName,omitempty" name:"KindName"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// vpcId

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// cpu

	CPU *float64 `json:"CPU,omitempty" name:"CPU"`
	// 内存

	Memory *float64 `json:"Memory,omitempty" name:"Memory"`
}

type AddAlarmPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddAlarmPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeImagesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeImagesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeImagesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCredential struct {

	// CA 根证书

	CACert *string `json:"CACert,omitempty" name:"CACert"`
	// 认证用的Token

	Token *string `json:"Token,omitempty" name:"Token"`
}

type DeleteLogCollectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteLogCollectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogCollectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterPodsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的Pod总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Pod信息

		Pods []*PodInfo `json:"Pods,omitempty" name:"Pods"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterPodsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterPodsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionItem struct {

	// 集群实例id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 如果发生错误，则该字段包含错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
	// 如果当前正在巡检，则该字段包含检测巡检

	Progress *ClusterInspectionProgress `json:"Progress,omitempty" name:"Progress"`
	// 最近一次巡检结果概览

	LastResult *ClusterInspectionOverview `json:"LastResult,omitempty" name:"LastResult"`
	// 自动巡检周期，crontab格式

	Cron *string `json:"Cron,omitempty" name:"Cron"`
}

type RouteInfo struct {

	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 目的端CIDR。

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 下一跳地址。

	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
}

type DeleteClusterVirtualNodeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池id

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 是否强制删除，false若有pod运行就不删除

	Force *bool `json:"Force,omitempty" name:"Force"`
	// 虚拟节点名

	NodeNames []*string `json:"NodeNames,omitempty" name:"NodeNames"`
}

func (r *DeleteClusterVirtualNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterVirtualNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerStatus struct {

	// 容器名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 容器ID

	ContainerId *string `json:"ContainerId,omitempty" name:"ContainerId"`
	// 容器状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 容器处于该状态的原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 容器镜像ID

	Image *string `json:"Image,omitempty" name:"Image"`
}

type InstanceUpgradeClusterStatus struct {

	// pod总数

	PodTotal *int64 `json:"PodTotal,omitempty" name:"PodTotal"`
	// NotReady pod总数

	NotReadyPod *int64 `json:"NotReadyPod,omitempty" name:"NotReadyPod"`
}

type SummaryService struct {

	// Service名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Service状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// Service IP

	ServiceIp *string `json:"ServiceIp,omitempty" name:"ServiceIp"`
	// 外网IP

	ExternalIp *string `json:"ExternalIp,omitempty" name:"ExternalIp"`
	// 负载均衡ID

	LbId *string `json:"LbId,omitempty" name:"LbId"`
	// 负载均衡状态

	LbStatus *string `json:"LbStatus,omitempty" name:"LbStatus"`
	// Service访问类型

	AccessType *string `json:"AccessType,omitempty" name:"AccessType"`
	// 期望副本数

	DesiredReplicas *int64 `json:"DesiredReplicas,omitempty" name:"DesiredReplicas"`
	// 当前副本数

	CurrentReplicas *int64 `json:"CurrentReplicas,omitempty" name:"CurrentReplicas"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// ReasonMap

	ReasonMap *string `json:"ReasonMap,omitempty" name:"ReasonMap"`
	// SourceReasonMap

	SourceReasonMap *string `json:"SourceReasonMap,omitempty" name:"SourceReasonMap"`
	// Labels

	Labels *string `json:"Labels,omitempty" name:"Labels"`
	// SysLables

	SysLables *string `json:"SysLables,omitempty" name:"SysLables"`
	// UserLables

	UserLables *string `json:"UserLables,omitempty" name:"UserLables"`
	// ScaleType

	ScaleType *string `json:"ScaleType,omitempty" name:"ScaleType"`
	// HpaPolicy

	HpaPolicy *string `json:"HpaPolicy,omitempty" name:"HpaPolicy"`
}

type DescribeHelmChartRequest struct {
	*tchttp.BaseRequest

	// 搜索内容

	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *DescribeHelmChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHelmChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageRegistryCredential struct {

	// 地址

	Server *string `json:"Server,omitempty" name:"Server"`
	// user

	Username *string `json:"Username,omitempty" name:"Username"`
	// pass

	Password *string `json:"Password,omitempty" name:"Password"`
}

type CreateClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例额外需要设置参数信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// CVM创建透传参数，json化字符串格式，详见[RunInstances](/tcloud/api/Compute/CVM/APIs/云服务器（cvm）/版本（2017-03-12）/实例相关接口/RunInstances)接口。

	RunInstancePara *string `json:"RunInstancePara,omitempty" name:"RunInstancePara"`
}

func (r *CreateClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRouteTablesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群路由表对象。

		RouteTableSet []*RouteTableInfo `json:"RouteTableSet,omitempty" name:"RouteTableSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRouteTablesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRouteTablesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageInstance struct {

	// 镜像别名

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 操作系统名称

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 镜像ID

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)

	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`
	// 镜像的cpu架构，取值:  x86/arm

	Arch *string `json:"Arch,omitempty" name:"Arch"`
	// SportIpv6

	SportIpv6 *string `json:"SportIpv6,omitempty" name:"SportIpv6"`
}

type CheckClusterImageRequest struct {
	*tchttp.BaseRequest

	// 指定有效的镜像ID，格式形如img-xxxx。可通过登录控制台查询，也可调用接口 [DescribeImages](/tcloud/api/Compute/CVM/APIs/云服务器（cvm）/版本（2017-03-12）/镜像相关接口/DescribeImages)，取返回信息中的ImageId字段。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
}

func (r *CheckClusterImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群所属项目

		ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 集群描述

		ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmMetric struct {

	// 告警指标ID

	MetricId *string `json:"MetricId,omitempty" name:"MetricId"`
	// ARGUS系统Measurement

	Measurement *string `json:"Measurement,omitempty" name:"Measurement"`
	// 统计周期

	StatisticsPeriod *int64 `json:"StatisticsPeriod,omitempty" name:"StatisticsPeriod"`
	// 指标名

	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`
	// 指标描述

	ArgusPolicyName *string `json:"ArgusPolicyName,omitempty" name:"ArgusPolicyName"`
	// 告警判断设置

	Evaluator *Evaluator `json:"Evaluator,omitempty" name:"Evaluator"`
	// 持续周期

	ContinuePeriod *int64 `json:"ContinuePeriod,omitempty" name:"ContinuePeriod"`
	// 状态，true表示OK

	Status *bool `json:"Status,omitempty" name:"Status"`
	// 指标单位

	Unit *string `json:"Unit,omitempty" name:"Unit"`
}

type AutoScalingGroupRange struct {

	// 伸缩组最小实例数

	MinSize *int64 `json:"MinSize,omitempty" name:"MinSize"`
	// 伸缩组最大实例数

	MaxSize *int64 `json:"MaxSize,omitempty" name:"MaxSize"`
}

type fileEntry struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// chart值

	Value *string `json:"Value,omitempty" name:"Value"`
	// id

	ID *int64 `json:"ID,omitempty" name:"ID"`
	// fileEntry

	Children []*fileEntry `json:"Children,omitempty" name:"Children"`
}

type InstanceUpgradeProgressItem struct {

	// 节点instanceID

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
	// 任务生命周期
	// process 运行中
	// paused 已停止
	// pauing 正在停止
	// done  已完成
	// timeout 已超时
	// aborted 已取消
	// pending 还未开始

	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
	// 升级开始时间

	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`
	// 升级结束时间

	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`
	// 升级前检查结果

	CheckResult *InstanceUpgradePreCheckResult `json:"CheckResult,omitempty" name:"CheckResult"`
	// 升级步骤详情

	Detail []*TaskStepInfo `json:"Detail,omitempty" name:"Detail"`
}

type DisableClusterAuditResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DisableClusterAuditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableClusterAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeQuotaRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAsGroupAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterAsGroupAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterAsGroupAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterPublicLB struct {

	// 是否开启公网访问LB

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 允许访问的来源CIDR列表

	AllowFromCidrs []*string `json:"AllowFromCidrs,omitempty" name:"AllowFromCidrs"`
}

type CheckLogCollectorNameRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志采集规则名称

	LogCollectorName *string `json:"LogCollectorName,omitempty" name:"LogCollectorName"`
}

func (r *CheckLogCollectorNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLogCollectorNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteTableInfo struct {

	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 路由表CIDR。

	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`
	// VPC实例ID。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

type DescribeAvailableClusterVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		Versions []*string `json:"Versions,omitempty" name:"Versions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableClusterVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableClusterVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求任务的FlowId

		RequestFlowId *int64 `json:"RequestFlowId,omitempty" name:"RequestFlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterEndpointVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterEndpointVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallLogAgentRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// KubeletRoot目录

	KubeletRootDir *string `json:"KubeletRootDir,omitempty" name:"KubeletRootDir"`
}

func (r *InstallLogAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallLogAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HPAInfo struct {

	// 最大副本数

	MaxReplica *int64 `json:"MaxReplica,omitempty" name:"MaxReplica"`
	// 最小副本数

	MinReplica *int64 `json:"MinReplica,omitempty" name:"MinReplica"`
	// 伸缩指标数组

	Metricss []*ScaleMetrics `json:"Metricss,omitempty" name:"Metricss"`
}

type CreateIndependentClusterRequest struct {
	*tchttp.BaseRequest

	// CVM创建透传参数，json化字符串格式, 各种角色的节点配置信息

	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode"`
	// 集群容器网络配置信息

	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`
	// 集群的基本配置信息

	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`
	// 集群高级配置信息

	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`
	// 节点高级配置信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
}

func (r *CreateIndependentClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIndependentClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExistedInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写查询集群列表 接口中返回的 ClusterId 字段（仅通过ClusterId获取需要过滤条件中的VPCID。节点状态比较时会使用该地域下所有集群中的节点进行比较。参数不支持同时指定InstanceIds和ClusterId。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 按照一个或者多个实例ID查询。实例ID形如：ins-xxxxxxxx。（此参数的具体格式可参考API简介的id.N一节）。每次请求的实例的上限为100。参数不支持同时指定InstanceIds和Filters。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 过滤条件,字段和详见[DescribeInstances](/tcloud/api/Compute/CVM/APIs/云服务器（cvm）/版本（2017-03-12）/实例相关接口/DescribeInstances)如果设置了ClusterId，会附加集群的VPCID作为查询字段，在此情况下如果在Filter中指定了"vpc-id"作为过滤字段，指定的VPCID必须与集群的VPCID相同。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 实例IP进行过滤(同时支持内网IP和外网IP)

	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`
	// 实例名称进行过滤

	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`
	// 偏移量，默认为0。用来控制分页的参数；Limit 为单次返回的最多条目数量，Offset 为偏移量。当相应结果是列表形式时，如果数量超过了 Limit 所限定的值，那么只返回 Limit 个值。

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。用来控制分页的参数；Limit 为单次返回的最多条目数量，Offset 为偏移量。当相应结果是列表形式时，如果数量超过了 Limit 所限定的值，那么只返回 Limit 个值。

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeExistedInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExistedInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NotifySettings struct {

	// 告警接收组（用户组）

	ReceiverGroups []*int64 `json:"ReceiverGroups,omitempty" name:"ReceiverGroups"`
	// 告警通知方式。目前有SMS、EMAIL、CALL、WECHAT方式。
	// 分别代表：短信、邮件、电话、微信

	NotifyWay []*string `json:"NotifyWay,omitempty" name:"NotifyWay"`
	// 电话告警顺序。
	// 注：NotifyWay选择CALL，采用该参数。

	PhoneNotifyOrder []*int64 `json:"PhoneNotifyOrder,omitempty" name:"PhoneNotifyOrder"`
	// 电话告警次数。
	// 注：NotifyWay选择CALL，采用该参数。

	PhoneCircleTimes *int64 `json:"PhoneCircleTimes,omitempty" name:"PhoneCircleTimes"`
	// 电话告警轮内间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。

	PhoneInnerInterval *int64 `json:"PhoneInnerInterval,omitempty" name:"PhoneInnerInterval"`
	// 电话告警轮外间隔。单位：秒
	// 注：NotifyWay选择CALL，采用该参数。

	PhoneCircleInterval *int64 `json:"PhoneCircleInterval,omitempty" name:"PhoneCircleInterval"`
	// 电话告警触达通知
	// 注：NotifyWay选择CALL，采用该参数。

	PhoneArriveNotice *int64 `json:"PhoneArriveNotice,omitempty" name:"PhoneArriveNotice"`
}

type DeleteClusterRouteTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckLogCollectorNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志采集规则是否存在（true：存在，false：不存在）

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckLogCollectorNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLogCollectorNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群路由对象。

		RouteSet []*RouteInfo `json:"RouteSet,omitempty" name:"RouteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTableConflictsRequest struct {
	*tchttp.BaseRequest

	// 路由表CIDR

	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`
	// 路由表绑定的VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeRouteTableConflictsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTableConflictsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckIfLogCollectorExistsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志采集规则名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CheckIfLogCollectorExistsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckIfLogCollectorExistsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPoliciesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAlarmPoliciesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmPoliciesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceIdsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点ID列表

		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
		// 数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInstanceIdsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInstanceIdsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterAdvancedSettings struct {

	// 是否启用IPVS

	IPVS *bool `json:"IPVS,omitempty" name:"IPVS"`
	// 是否启用集群节点自动扩缩容(创建集群流程不支持开启此功能)

	AsEnabled *bool `json:"AsEnabled,omitempty" name:"AsEnabled"`
	// 集群使用的runtime类型，包括"docker"和"containerd"两种类型，默认为"docker"

	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	// 集群中节点NodeName类型（包括 hostname,lan-ip两种形式，默认为lan-ip。如果开启了hostname模式，创建节点时需要设置HostName参数，并且InstanceName需要和HostName一致）

	NodeNameType *string `json:"NodeNameType,omitempty" name:"NodeNameType"`
	// 集群自定义参数

	ExtraArgs *ClusterExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
	// 集群网络类型（包括GR(全局路由)和VPC-CNI两种模式，默认为GR。

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
	// 集群VPC-CNI模式是否为非固定IP，默认: FALSE 固定IP。

	IsNonStaticIpMode *bool `json:"IsNonStaticIpMode,omitempty" name:"IsNonStaticIpMode"`
	// 集群是否支持双栈

	IsDualStack *string `json:"IsDualStack,omitempty" name:"IsDualStack"`
}

type File struct {

	// 数据

	Data *string `json:"Data,omitempty" name:"Data"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
}

type RegionInstance struct {

	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域ID

	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
	// 地域状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 地域特性开关(按照JSON的形式返回所有属性)

	FeatureGates *string `json:"FeatureGates,omitempty" name:"FeatureGates"`
	// 地域简称

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 地域白名单

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type GetVbcRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 详情

		Detail []*CcnRoute `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetVbcRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVbcRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogCollectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志收集总个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志收集信息表

		LogCollectorList []*LogCollector `json:"LogCollectorList,omitempty" name:"LogCollectorList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListLogCollectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogCollectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeClusterPublicLB struct {

	// 是否开启公网访问LB

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 允许访问的公网cidr

	AllowFromCidrs []*string `json:"AllowFromCidrs,omitempty" name:"AllowFromCidrs"`
}

type Node struct {

	// 虚拟节点名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 节点池id

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 节点子网id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 节点状态

	Phase *string `json:"Phase,omitempty" name:"Phase"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type DescribeClusterInspectionOverviewsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群巡检报告列表

		OverviewSet []*ClusterInspectionOverview `json:"OverviewSet,omitempty" name:"OverviewSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInspectionOverviewsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInspectionOverviewsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceCreateProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建进度

		Progress []*string `json:"Progress,omitempty" name:"Progress"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceCreateProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceCreateProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOverageRequest struct {
	*tchttp.BaseRequest
}

func (r *QueryOverageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOverageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogCollectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateLogCollectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogCollectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TKEEdgeClusterResources struct {

	// 集群id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群cpu总和

	CPU *float64 `json:"CPU,omitempty" name:"CPU"`
	// 集群memory总和

	Memory *float64 `json:"Memory,omitempty" name:"Memory"`
	// 集群节点数

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
}

type ZoneResourceInfo struct {

	// 可用区，如ap-guangzhou-2

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 资源信息

	ResourceInfo *string `json:"ResourceInfo,omitempty" name:"ResourceInfo"`
}

type DeleteClusterRouteRequest struct {
	*tchttp.BaseRequest

	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 下一跳地址。

	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
	// 目的端CIDR。

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
}

func (r *DeleteClusterRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterCommonNamesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 对应的子用户的CommonName

		CommonNames *CommonNames `json:"CommonNames,omitempty" name:"CommonNames"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterCommonNamesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterCommonNamesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLogCollectorStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetLogCollectorStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLogCollectorStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetVbcRouteRequest struct {
	*tchttp.BaseRequest

	// vpc唯一id

	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
	// 集群cidr

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *GetVbcRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVbcRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExistedInstancesForNode struct {

	// 节点角色，取值:MASTER_ETCD, WORKER。MASTER_ETCD只有在创建 INDEPENDENT_CLUSTER 独立集群时需要指定。MASTER_ETCD节点数量为3～7，建议为奇数。MASTER_ETCD最小配置为4C8G。

	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
	// 已存在实例的重装参数

	ExistedInstancesPara *ExistedInstancesPara `json:"ExistedInstancesPara,omitempty" name:"ExistedInstancesPara"`
	// 节点高级设置，会覆盖集群级别设置的InstanceAdvancedSettings（当前只对节点自定义参数ExtraArgs生效）

	InstanceAdvancedSettingsOverride *InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverride,omitempty" name:"InstanceAdvancedSettingsOverride"`
}

type MachineGroup struct {

	// 机器组ID

	GroupId *string `json:"GroupId,omitempty" name:"GroupId"`
	// 机器组名称

	GroupName *string `json:"GroupName,omitempty" name:"GroupName"`
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

type Step struct {

	// 名称

	Type *string `json:"Type,omitempty" name:"Type"`
	// 最后一次执行时间

	LastProbeTime *string `json:"LastProbeTime,omitempty" name:"LastProbeTime"`
	// 当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 执行信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type GetVbcInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 详情

		Detail []*CcnRoute `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetVbcInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVbcInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAlarmPolicyResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAlarmPolicyResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyAlarmPolicyResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type QueryOverageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 取值true和false,表示license是否超用

		Pass *string `json:"Pass,omitempty" name:"Pass"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *QueryOverageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *QueryOverageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Taint struct {

	// Key

	Key *string `json:"Key,omitempty" name:"Key"`
	// Value

	Value *string `json:"Value,omitempty" name:"Value"`
	// Effect

	Effect *string `json:"Effect,omitempty" name:"Effect"`
}

type CreateClusterRouteTableRequest struct {
	*tchttp.BaseRequest

	// 路由表名称

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 路由表CIDR

	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`
	// 路由表绑定的VPC

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 是否忽略CIDR冲突

	IgnoreClusterCidrConflict *int64 `json:"IgnoreClusterCidrConflict,omitempty" name:"IgnoreClusterCidrConflict"`
}

func (r *CreateClusterRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRouteTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionInfo struct {

	// id

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// gz

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// ap-guangzhou

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 1

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// alluser

	Status *string `json:"Status,omitempty" name:"Status"`
	// time

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// time

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type CreateClusterVirtualNodePoolRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 虚拟节点池名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 子网ID数组

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 安全组数组

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 虚拟节点池的Labels

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 虚拟节点池的Taints

	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
	// 是否可调度

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
}

func (r *CreateClusterVirtualNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterVirtualNodePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CommonNames struct {

	// 对应commonName

	CN *string `json:"CN,omitempty" name:"CN"`
	// 子账户id

	SubaccountUin *string `json:"SubaccountUin,omitempty" name:"SubaccountUin"`
}

type SecretVolume struct {

	// secret

	Data []*KeyValueData `json:"Data,omitempty" name:"Data"`
	// type

	Type *string `json:"Type,omitempty" name:"Type"`
}

type AddClusterCIDRToVbcResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回详情

		Detail []*CcnRoute `json:"Detail,omitempty" name:"Detail"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterCIDRToVbcResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterCIDRToVbcResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogTopicsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志主题的数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志主题集合

		Topics []*LogTopic `json:"Topics,omitempty" name:"Topics"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClsLogTopicsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsLogTopicsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CollectAllCoreRequest struct {
	*tchttp.BaseRequest
}

func (r *CollectAllCoreRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CollectAllCoreRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableVpcPeerClusterRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表名称

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 对等连接实例ID

	VpcPeerId *string `json:"VpcPeerId,omitempty" name:"VpcPeerId"`
}

func (r *EnableVpcPeerClusterRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableVpcPeerClusterRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadHelmChartRequest struct {
	*tchttp.BaseRequest

	// cos文件地址

	CosPath *string `json:"CosPath,omitempty" name:"CosPath"`
	// namespace 默认去UIN

	NamespaceName *string `json:"NamespaceName,omitempty" name:"NamespaceName"`
	// namespace 默认取APPID

	RegistryId *string `json:"RegistryId,omitempty" name:"RegistryId"`
}

func (r *UploadHelmChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadHelmChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AcquireClusterAdminRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AcquireClusterAdminRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcquireClusterAdminRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RouteTableConflict struct {

	// 路由表类型。

	RouteTableType *string `json:"RouteTableType,omitempty" name:"RouteTableType"`
	// 路由表CIDR。

	RouteTableCidrBlock *string `json:"RouteTableCidrBlock,omitempty" name:"RouteTableCidrBlock"`
	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 路由表ID。

	RouteTableId *string `json:"RouteTableId,omitempty" name:"RouteTableId"`
}

type DescribeCcnInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 关联网络实例列表

		InstanceSet []*CcnInstance `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterCreateProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 创建进度

		Progress []*Step `json:"Progress,omitempty" name:"Progress"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterCreateProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterCreateProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesVersionRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表

	ClusterId []*string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeInstancesVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UploadHelmChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UploadHelmChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UploadHelmChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodePool struct {

	// NodePoolId

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// ClusterInstanceId

	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`
	// LifeState

	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
	// LaunchConfigurationId

	LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`
	// AutoscalingGroupId

	AutoscalingGroupId *string `json:"AutoscalingGroupId,omitempty" name:"AutoscalingGroupId"`
	// Labels

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// Taints

	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
}

type DescribeAlarmPoliciesRequest struct {
	*tchttp.BaseRequest

	// k8s集群ID

	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 最大输出条数，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤条件

	Filter *AlarmPolicyFilter `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeAlarmPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAlarmPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointVipStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 端口操作状态 (Creating 创建中  CreateFailed 创建失败 Created 创建完成 Deleting 删除中 DeletedFailed 删除失败 Deleted 已删除 NotFound 未发现操作 )

		Status *string `json:"Status,omitempty" name:"Status"`
		// 操作失败的原因

		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointVipStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterEndpointVipStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterVirtualNodePoolRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 虚拟节点池名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 虚拟节点池的Labels

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 虚拟节点池的Taints

	Taints []*Taint `json:"Taints,omitempty" name:"Taints"`
	// 是否可调度

	Unschedulable *bool `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 节点池id

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
}

func (r *ModifyClusterVirtualNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterVirtualNodePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterCIDRFromCcnRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *DeleteClusterCIDRFromCcnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterCIDRFromCcnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterEndpointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterCommonNamesRequest struct {
	*tchttp.BaseRequest

	// 需要获取commonName的用户uin数组

	SubaccountUins []*string `json:"SubaccountUins,omitempty" name:"SubaccountUins"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterCommonNamesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterCommonNamesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群总个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群信息列表

		Clusters []*Cluster `json:"Clusters,omitempty" name:"Clusters"`
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

type DescribeClusterServicesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 返回的Service总数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Service详细信息

		Services []*SummaryService `json:"Services,omitempty" name:"Services"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterServicesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServicesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceStatus struct {

	// 集群ID

	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`
	// 各资源状态

	Status []*ResourceStatusItem `json:"Status,omitempty" name:"Status"`
}

type AddVpcCniSubnetsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// VPC ID

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 要添加的子网ID

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

func (r *AddVpcCniSubnetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVpcCniSubnetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ImageAttributeSet struct {

	// 镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// 内部镜像id

	InnerImageId *uint64 `json:"InnerImageId,omitempty" name:"InnerImageId"`
}

type ResumeClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResumeClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CertificateInfo struct {

	// kubeconfig信息

	Kubeconfig *string `json:"Kubeconfig,omitempty" name:"Kubeconfig"`
	// 证书状态

	CertificateStatus *string `json:"CertificateStatus,omitempty" name:"CertificateStatus"`
	// 证书certificate

	ClientCertificate *string `json:"ClientCertificate,omitempty" name:"ClientCertificate"`
	// 证书key

	ClientKey *string `json:"ClientKey,omitempty" name:"ClientKey"`
	// 对应用户的昵称

	NickName *string `json:"NickName,omitempty" name:"NickName"`
	// 用户的uin

	SubAccountUin *string `json:"SubAccountUin,omitempty" name:"SubAccountUin"`
	// 证书过期时间

	ExpirationTime *string `json:"ExpirationTime,omitempty" name:"ExpirationTime"`
}

type PodInfo struct {

	// Pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Pod当前状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 处于当前状态原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 在kubernetes中展示的原因

	SourceReason *string `json:"SourceReason,omitempty" name:"SourceReason"`
	// Pod IP

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// Pod重启次数

	RestartCount *int64 `json:"RestartCount,omitempty" name:"RestartCount"`
	// Pod就绪容器数量

	ReadyCount *int64 `json:"ReadyCount,omitempty" name:"ReadyCount"`
	// 所在节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 所在节点IP

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
	// Pod启动时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// Pod所含容器信息

	Containers []*ContainerStatus `json:"Containers,omitempty" name:"Containers"`
}

type CreateTKEClusterRequest struct {
	*tchttp.BaseRequest

	// CVM创建透传参数，json化字符串格式, 各种角色的节点配置信息

	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode"`
	// 集群容器网络配置信息

	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`
	// 集群的基本配置信息

	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`
	// 集群高级配置信息

	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`
	// 节点高级配置信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 是否独立集群

	IndependentCluster *bool `json:"IndependentCluster,omitempty" name:"IndependentCluster"`
}

func (r *CreateTKEClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTKEClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterVirtualNodePoolRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点池id

	NodePoolIds []*string `json:"NodePoolIds,omitempty" name:"NodePoolIds"`
	// 是否强制删除，false若有pod运行就不删除

	Force *bool `json:"Force,omitempty" name:"Force"`
}

func (r *DeleteClusterVirtualNodePoolRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterVirtualNodePoolRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询集群访问端口状态（Created 开启成功，Creating 开启中中，NotFound 未开启）

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterEndpointStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSwitchesRequest struct {
	*tchttp.BaseRequest

	// 集群ID数组

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeLogSwitchesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogSwitchesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInstancesVersionItem struct {

	// 版本

	InstanceVersion *string `json:"InstanceVersion,omitempty" name:"InstanceVersion"`
	// 节点数

	Total *uint64 `json:"Total,omitempty" name:"Total"`
	// 是否可升级

	UpgradeAble *bool `json:"UpgradeAble,omitempty" name:"UpgradeAble"`
}

type EnableLogCollectorRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *EnableLogCollectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableLogCollectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceAdvancedSettings struct {

	// 数据盘挂载点, 默认不挂载数据盘. 已格式化的 ext3，ext4，xfs 文件系统的数据盘将直接挂载，其他文件系统或未格式化的数据盘将自动格式化为ext4 并挂载，请注意备份数据! 无数据盘或有多块数据盘的云主机此设置不生效。

	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`
	// dockerd --graph 指定值, 默认为 /var/lib/docker

	DockerGraphPath *string `json:"DockerGraphPath,omitempty" name:"DockerGraphPath"`
	// base64 编码的用户脚本, 此脚本会在 k8s 组件运行后执行, 需要用户保证脚本的可重入及重试逻辑, 脚本及其生成的日志文件可在节点的 /data/ccs_userscript/ 路径查看, 如果要求节点需要在进行初始化完成后才可加入调度, 可配合 unschedulable 参数使用, 在 userScript 最后初始化完成后, 添加 kubectl uncordon nodename --kubeconfig=/root/.kube/config 命令使节点加入调度

	UserScript *string `json:"UserScript,omitempty" name:"UserScript"`
	// 设置加入的节点是否参与调度，默认值为0，表示参与调度；非0表示不参与调度, 待节点初始化完成之后, 可执行kubectl uncordon nodename使node加入调度.

	Unschedulable *int64 `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// 节点Label数组

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 数据盘相关信息

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// 节点相关的自定义参数信息

	ExtraArgs *InstanceExtraArgs `json:"ExtraArgs,omitempty" name:"ExtraArgs"`
	// base64 编码的用户脚本, 此脚本会在 k8s 组件运行前执行, 需要用户保证脚本的可重入及重试逻辑, 脚本及其生成的日志文件可在节点的 /data/ccs_userscript/ 路径查看.

	PreStartUserScript *string `json:"PreStartUserScript,omitempty" name:"PreStartUserScript"`
}

type Label struct {

	// map表中的Name

	Name *string `json:"Name,omitempty" name:"Name"`
	// map表中的Value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type LogTopic struct {

	// 日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 日志主题名称

	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 日志类型

	LogType *string `json:"LogType,omitempty" name:"LogType"`
	// 是否采集

	Collection *bool `json:"Collection,omitempty" name:"Collection"`
	// 是否有索引

	Index *bool `json:"Index,omitempty" name:"Index"`
}

type CreateIndependentClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateIndependentClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateIndependentClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodePoolOption struct {

	// 是否加入节点池

	AddToNodePool *bool `json:"AddToNodePool,omitempty" name:"AddToNodePool"`
	// 节点池id

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
}

type TaskStepInfo struct {

	// 步骤名称

	Step *string `json:"Step,omitempty" name:"Step"`
	// 生命周期
	// pending : 步骤未开始
	// running: 步骤执行中
	// success: 步骤成功完成
	// failed: 步骤失败

	LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
	// 步骤开始时间

	StartAt *string `json:"StartAt,omitempty" name:"StartAt"`
	// 步骤结束时间

	EndAt *string `json:"EndAt,omitempty" name:"EndAt"`
	// 若步骤生命周期为failed,则此字段显示错误信息

	FailedMsg *string `json:"FailedMsg,omitempty" name:"FailedMsg"`
	// 实例ID(仅节点升级使用)

	InstanceID *string `json:"InstanceID,omitempty" name:"InstanceID"`
}

type DescribeFlowIdStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 开启集群外网访问端口的任务ID

	RequestFlowId *int64 `json:"RequestFlowId,omitempty" name:"RequestFlowId"`
}

func (r *DescribeFlowIdStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowIdStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点实例ID

		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesVersionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 不同版本节点个数统计结果

		Clusters []*ClusterInstancesVersion `json:"Clusters,omitempty" name:"Clusters"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesVersionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesVersionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardIDResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 仪表盘ID

		DashboardID *string `json:"DashboardID,omitempty" name:"DashboardID"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetDashboardIDResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardIDResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type HostNameValue struct {

	// 主机名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 主机数量

	Value *uint64 `json:"Value,omitempty" name:"Value"`
}

type LogSet struct {

	// 日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
	// 日志集名称

	LogSetName []*string `json:"LogSetName,omitempty" name:"LogSetName"`
	// 周期

	Period *int64 `json:"Period,omitempty" name:"Period"`
	// 创建时间

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type DescribeAvailableClusterVersionRequest struct {
	*tchttp.BaseRequest

	// 集群 Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群 Id 列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeAvailableClusterVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableClusterVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstallLogAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *InstallLogAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *InstallLogAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionReportRequest struct {
	*tchttp.BaseRequest

	// 集群实例id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 报告页id

	ReportId *string `json:"ReportId,omitempty" name:"ReportId"`
}

func (r *DescribeClusterInspectionReportRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInspectionReportRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterEndpointRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）

	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *DeleteClusterEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterEndpointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ExistedInstance struct {

	// 实例是否支持加入集群(TRUE 可以加入 FALSE 不能加入)。

	Usable *bool `json:"Usable,omitempty" name:"Usable"`
	// 实例不支持加入的原因。

	UnusableReason *string `json:"UnusableReason,omitempty" name:"UnusableReason"`
	// 实例已经所在的集群ID。

	AlreadyInCluster *string `json:"AlreadyInCluster,omitempty" name:"AlreadyInCluster"`
	// 实例ID形如：ins-xxxxxxxx。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称。

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例主网卡的内网IP列表。

	PrivateIpAddresses []*string `json:"PrivateIpAddresses,omitempty" name:"PrivateIpAddresses"`
	// 实例主网卡的公网IP列表。
	// 注意：此字段可能返回 null，表示取不到有效值。

	PublicIpAddresses []*string `json:"PublicIpAddresses,omitempty" name:"PublicIpAddresses"`
	// 创建时间。按照ISO8601标准表示，并且使用UTC时间。格式为：YYYY-MM-DDThh:mm:ssZ。

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 实例计费模式。取值范围：
	// PREPAID：表示预付费，即包年包月
	// POSTPAID_BY_HOUR：表示后付费，即按量计费
	// CDHPAID：CDH付费，即只对CDH计费，不对CDH上的实例计费。

	InstanceChargeType *string `json:"InstanceChargeType,omitempty" name:"InstanceChargeType"`
	// 实例的CPU核数，单位：核。

	CPU *uint64 `json:"CPU,omitempty" name:"CPU"`
	// 实例内存容量，单位：GB。

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 操作系统名称。

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 实例机型。

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// ipv6地址

	IPv6Addresses []*string `json:"IPv6Addresses,omitempty" name:"IPv6Addresses"`
}

type MachineCore struct {

	// 请求id

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 产品描述

	ProductIdDescribe *string `json:"ProductIdDescribe,omitempty" name:"ProductIdDescribe"`
	// 使用单元

	UsageUnit *string `json:"UsageUnit,omitempty" name:"UsageUnit"`
	// 使用量

	UsageValue *string `json:"UsageValue,omitempty" name:"UsageValue"`
}

type DescribeClusterEndpointVipStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterEndpointVipStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterEndpointVipStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterUpgradingStateRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 操作类型：
	// pause 暂停集群升级
	// resume 继续集群升级
	// abort 取消集群升级

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 如果是取消集群升级，该参数设置为true表示强制取消

	ForceAbort *bool `json:"ForceAbort,omitempty" name:"ForceAbort"`
}

func (r *ModifyClusterUpgradingStateRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterUpgradingStateRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterStatus struct {

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群状态

	ClusterState *string `json:"ClusterState,omitempty" name:"ClusterState"`
	// 集群下机器实例的状态

	ClusterInstanceState *string `json:"ClusterInstanceState,omitempty" name:"ClusterInstanceState"`
	// 集群是否开启监控

	ClusterBMonitor *bool `json:"ClusterBMonitor,omitempty" name:"ClusterBMonitor"`
	// 集群创建中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败

	ClusterInitNodeNum *int64 `json:"ClusterInitNodeNum,omitempty" name:"ClusterInitNodeNum"`
	// 集群运行中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败

	ClusterRunningNodeNum *int64 `json:"ClusterRunningNodeNum,omitempty" name:"ClusterRunningNodeNum"`
	// 集群异常的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败

	ClusterFailedNodeNum *int64 `json:"ClusterFailedNodeNum,omitempty" name:"ClusterFailedNodeNum"`
	// 集群已关机的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败

	ClusterClosedNodeNum *int64 `json:"ClusterClosedNodeNum,omitempty" name:"ClusterClosedNodeNum"`
	// 集群关机中的节点数，-1表示获取节点状态超时，-2表示获取节点状态失败

	ClusterClosingNodeNum *int64 `json:"ClusterClosingNodeNum,omitempty" name:"ClusterClosingNodeNum"`
}

type CheckLogCollectorHostPathResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 主机路径是否合法

		Result *bool `json:"Result,omitempty" name:"Result"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckLogCollectorHostPathResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLogCollectorHostPathResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterKubeconfigRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 返回的config中sever地址类型：MASTER/INNERLB/INTERNETLB

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeClusterKubeconfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterKubeconfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterServicesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 是否使用所有命名空间

	AllNamespace *uint64 `json:"AllNamespace,omitempty" name:"AllNamespace"`
}

func (r *DescribeClusterServicesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterServicesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功启动更新的节点Id列表

		InstanceIdSet []*string `json:"InstanceIdSet,omitempty" name:"InstanceIdSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Chart struct {

	// 文件列表

	Files []*File `json:"Files,omitempty" name:"Files"`
	// Lock

	Lock *string `json:"Lock,omitempty" name:"Lock"`
	// metadata

	Metadata *Metadata `json:"Metadata,omitempty" name:"Metadata"`
	// Schema

	Schema *string `json:"Schema,omitempty" name:"Schema"`
	// Templates

	Templates []*File `json:"Templates,omitempty" name:"Templates"`
	// Values

	Values *string `json:"Values,omitempty" name:"Values"`
}

type ExistedInstancesPara struct {

	// 集群ID

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例额外需要设置参数信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。

	EnhancedService *string `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）

	LoginSettings *string `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 重装系统时，可以指定修改实例的HostName(集群为HostName模式时，此参数必传，规则名称除不支持大写字符外与[RunInstances](/tcloud/api/Compute/CVM/APIs/云服务器（cvm）/版本（2017-03-12）/实例相关接口/RunInstances)接口HostName一致)

	HostName *string `json:"HostName,omitempty" name:"HostName"`
}

type InstanceExtraArgs struct {

	// kubelet自定义参数

	Kubelet []*string `json:"Kubelet,omitempty" name:"Kubelet"`
}

type DescribeClusterSecurityResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群的账号名称

		UserName *string `json:"UserName,omitempty" name:"UserName"`
		// 集群的访问密码

		Password *string `json:"Password,omitempty" name:"Password"`
		// 集群访问CA证书

		CertificationAuthority *string `json:"CertificationAuthority,omitempty" name:"CertificationAuthority"`
		// 集群访问的地址

		ClusterExternalEndpoint *string `json:"ClusterExternalEndpoint,omitempty" name:"ClusterExternalEndpoint"`
		// 集群访问的域名

		Domain *string `json:"Domain,omitempty" name:"Domain"`
		// 集群Endpoint地址

		PgwEndpoint *string `json:"PgwEndpoint,omitempty" name:"PgwEndpoint"`
		// 集群访问策略组

		SecurityPolicy []*string `json:"SecurityPolicy,omitempty" name:"SecurityPolicy"`
		// 用户的kubeconfig文件信息

		Kubeconfig *string `json:"Kubeconfig,omitempty" name:"Kubeconfig"`
		// 集群的jnsgw地址

		JnsGwEndpoint *string `json:"JnsGwEndpoint,omitempty" name:"JnsGwEndpoint"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterSecurityResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterSecurityResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHelmChartRequest struct {
	*tchttp.BaseRequest

	// Chart名称

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
}

func (r *DeleteHelmChartRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHelmChartRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeExistedInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 已经存在的实例信息数组。

		ExistedInstanceSet []*ExistedInstance `json:"ExistedInstanceSet,omitempty" name:"ExistedInstanceSet"`
		// 符合条件的实例数量。

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeExistedInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeExistedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeInstanceProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 升级节点总数

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 已升级节点总数

		Done *int64 `json:"Done,omitempty" name:"Done"`
		// 升级任务生命周期
		// process 运行中
		// paused 已停止
		// pauing 正在停止
		// done  已完成
		// timeout 已超时
		// aborted 已取消

		LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
		// 各节点升级进度详情

		Instances []*InstanceUpgradeProgressItem `json:"Instances,omitempty" name:"Instances"`
		// 集群当前状态

		ClusterStatus *InstanceUpgradeClusterStatus `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeInstanceProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeInstanceProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterKubeconfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功更新证书的用户uin数组

		UpdatedSubAccounts []*string `json:"UpdatedSubAccounts,omitempty" name:"UpdatedSubAccounts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateClusterKubeconfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterKubeconfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogCollectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateLogCollectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogCollectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VersionInstance struct {

	// 版本名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本信息

	Version *string `json:"Version,omitempty" name:"Version"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

type DescribeLogSwitchInfo struct {

	// 是否开启

	Enable *bool `json:"Enable,omitempty" name:"Enable"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
}

type KeyValueData struct {

	// key

	Key *string `json:"Key,omitempty" name:"Key"`
	// value

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DeleteClusterEndpointVipRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DeleteClusterEndpointVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterEndpointVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterImageRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 指定有效的镜像ID，格式形如img-xxxx。可通过登录控制台查询，也可调用接口[DescribeImages](/tcloud/api/Compute/CVM/APIs/云服务器（cvm）/版本（2017-03-12）/镜像相关接口/DescribeImages)，取返回信息中的ImageId字段。

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// "DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)

	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`
}

func (r *ModifyClusterImageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterImageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAsGroupOptionAttributeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群弹性伸缩属性

	ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitempty" name:"ClusterAsGroupOption"`
}

func (r *ModifyClusterAsGroupOptionAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterAsGroupOptionAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SwitchSet struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志运维

	Log *DescribeLogSwitchInfo `json:"Log,omitempty" name:"Log"`
	// 事件运维

	Event *DescribeLogSwitchInfo `json:"Event,omitempty" name:"Event"`
	// 审计运维

	Audit *DescribeLogSwitchInfo `json:"Audit,omitempty" name:"Audit"`
}

type CheckIfLogCollectorExistsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志采集规则是否已经存在

		CheckResult *bool `json:"CheckResult,omitempty" name:"CheckResult"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckIfLogCollectorExistsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckIfLogCollectorExistsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterEndpointVipRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)

	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies"`
}

func (r *CreateClusterEndpointVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterEndpointVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 版本数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 版本列表

		VersionInstanceSet []*VersionInstance `json:"VersionInstanceSet,omitempty" name:"VersionInstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunInstancesForNode struct {

	// 节点角色，取值:MASTER_ETCD, WORKER。MASTER_ETCD只有在创建 INDEPENDENT_CLUSTER 独立集群时需要指定。MASTER_ETCD节点数量为3～7，建议为奇数。MASTER_ETCD节点最小配置为4C8G。

	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
	// CVM创建透传参数，json化字符串格式，详见[RunInstances](/tcloud/api/Compute/CVM/APIs/云服务器（cvm）/版本（2017-03-12）/实例相关接口/RunInstances)接口，传入公共参数外的其他参数即可，其中ImageId会替换为TKE集群OS对应的镜像。

	RunInstancesPara []*string `json:"RunInstancesPara,omitempty" name:"RunInstancesPara"`
	// 节点高级设置，该参数会覆盖集群级别设置的InstanceAdvancedSettings，和上边的RunInstancesPara按照顺序一一对应（当前只对节点自定义参数ExtraArgs生效）。

	InstanceAdvancedSettingsOverrides []*InstanceAdvancedSettings `json:"InstanceAdvancedSettingsOverrides,omitempty" name:"InstanceAdvancedSettingsOverrides"`
}

type AddAlarmPolicyRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 告警策略设置

	AlarmPolicySettings *AlarmPolicySettings `json:"AlarmPolicySettings,omitempty" name:"AlarmPolicySettings"`
	// 告警通知设置

	NotifySettings *NotifySettings `json:"NotifySettings,omitempty" name:"NotifySettings"`
	// 告警屏蔽设置

	ShieldSettings *ShieldSettings `json:"ShieldSettings,omitempty" name:"ShieldSettings"`
	// WebhookUrl

	WebhookUrl *string `json:"WebhookUrl,omitempty" name:"WebhookUrl"`
}

func (r *AddAlarmPolicyRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddAlarmPolicyRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
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

type DescribeClusterInspectionsRequest struct {
	*tchttp.BaseRequest

	// 目标集群实例id数组

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeClusterInspectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInspectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeLogSwitchesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 运维配置

		SwitchSet []*SwitchSet `json:"SwitchSet,omitempty" name:"SwitchSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeLogSwitchesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeLogSwitchesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddExistInstancesParam struct {

	// 目标集群

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例额外需要设置参数信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。

	EnhancedService *string `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）

	LoginSettings *string `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。（目前仅支持设置单个sgId）

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type DeleteClusterAsGroupsRequest struct {
	*tchttp.BaseRequest

	// 集群ID，通过[DescribeClusters](../集群相关接口/DescribeClusters)接口获取。

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群伸缩组ID的列表

	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds"`
	// 是否保留伸缩组中的节点(默认值： false(不保留))

	KeepInstance *bool `json:"KeepInstance,omitempty" name:"KeepInstance"`
}

func (r *DeleteClusterAsGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterAsGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterCIDRFromVbcRequest struct {
	*tchttp.BaseRequest

	// vpc唯一id

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 集群cidr

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *DeleteClusterCIDRFromVbcRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterCIDRFromVbcRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterSecurityRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 是否返回内网地址(默认: FALSE)

	JnsGwEndpointEnable *bool `json:"JnsGwEndpointEnable,omitempty" name:"JnsGwEndpointEnable"`
}

func (r *DescribeClusterSecurityRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterSecurityRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableClusterAuditResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableClusterAuditResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableClusterAuditResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceUpgradePreCheckResultItem struct {

	// 工作负载的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载类型

	WorkLoadKind *string `json:"WorkLoadKind,omitempty" name:"WorkLoadKind"`
	// 工作负载名称

	WorkLoadName *string `json:"WorkLoadName,omitempty" name:"WorkLoadName"`
	// 驱逐节点前工作负载running的pod数目

	Before *uint64 `json:"Before,omitempty" name:"Before"`
	// 驱逐节点后工作负载running的pod数目

	After *uint64 `json:"After,omitempty" name:"After"`
	// 工作负载在本节点上的pod列表

	Pods []*string `json:"Pods,omitempty" name:"Pods"`
}

type RIUtilizationDetail struct {

	// 预留券ID

	ReservedInstanceId *string `json:"ReservedInstanceId,omitempty" name:"ReservedInstanceId"`
	// Pod唯一ID

	EksId *string `json:"EksId,omitempty" name:"EksId"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Pod的名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// Pod的命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 工作负载名称

	KindName *string `json:"KindName,omitempty" name:"KindName"`
	// Pod的uid

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 用量开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 用量结束时间

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 抵扣资源所属产品

	Product *string `json:"Product,omitempty" name:"Product"`
}

type ModifyClusterVirtualNodePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterVirtualNodePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterVirtualNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AvailableExtraArgs struct {

	// kube-apiserver可用的自定义参数

	KubeAPIServer []*Flag `json:"KubeAPIServer,omitempty" name:"KubeAPIServer"`
	// kube-controller-manager可用的自定义参数

	KubeControllerManager []*Flag `json:"KubeControllerManager,omitempty" name:"KubeControllerManager"`
	// kube-scheduler可用的自定义参数

	KubeScheduler []*Flag `json:"KubeScheduler,omitempty" name:"KubeScheduler"`
	// kubelet可用的自定义参数

	Kubelet []*Flag `json:"Kubelet,omitempty" name:"Kubelet"`
}

type DescribeClsLogTopicsRequest struct {
	*tchttp.BaseRequest

	// 日志集ID

	LogSetId *string `json:"LogSetId,omitempty" name:"LogSetId"`
}

func (r *DescribeClsLogTopicsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsLogTopicsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLogCollectorStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否开启了日志采集

		Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
		// 当前POD运行数量

		CurrentNumberScheduled *int64 `json:"CurrentNumberScheduled,omitempty" name:"CurrentNumberScheduled"`
		// 期望POD运行数量

		DesiredNumberScheduled *int64 `json:"DesiredNumberScheduled,omitempty" name:"DesiredNumberScheduled"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLogCollectorStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLogCollectorStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 偏移量，默认为0。用来控制分页的参数；Limit 为单次返回的最多条目数量，Offset 为偏移量。当相应结果是列表形式时，如果数量超过了 Limit 所限定的值，那么只返回 Limit 个值。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。用来控制分页的参数；Limit 为单次返回的最多条目数量，Offset 为偏移量。当相应结果是列表形式时，如果数量超过了 Limit 所限定的值，那么只返回 Limit 个值。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 需要获取的节点实例Id列表。如果为空，表示拉取集群下所有节点实例。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 节点角色, 分为：MASTER_ETCD, WORKER。传空则为获取全部角色

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 过滤条件列表

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *DescribeClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHelmChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		Results []*string `json:"Results,omitempty"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHelmChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHelmChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCLSLogConfigRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志采集配置

	LogConfig *string `json:"LogConfig,omitempty" name:"LogConfig"`
}

func (r *CreateCLSLogConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCLSLogConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListLogCollectorRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 最大输出条数，默认20

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量,默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 过滤条件,当前只支持按照单个条件”日志收集名称“进行过滤

	Search *string `json:"Search,omitempty" name:"Search"`
}

func (r *ListLogCollectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListLogCollectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCIDRSettings struct {

	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突。且网段范围必须在内网网段内，例如:10.1.0.0/14, 192.168.0.1/18,172.16.0.0/16。

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
	// 是否忽略 ClusterCIDR 冲突错误, 默认不忽略

	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitempty" name:"IgnoreClusterCIDRConflict"`
	// 集群中每个Node上最大的Pod数量。取值范围4～256。不为2的幂值时会向上取最接近的2的幂值。

	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`
	// 集群最大的service数量。取值范围32～32768，不为2的幂值时会向上取最接近的2的幂值。

	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`
	// 用于分配集群服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突。且网段范围必须在内网网段内，例如:10.1.0.0/14, 192.168.0.1/18,172.16.0.0/16。

	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`
	// VPC-CNI网络模式下，弹性网卡的子网Id。

	EniSubnetIds []*string `json:"EniSubnetIds,omitempty" name:"EniSubnetIds"`
	// VPC-CNI网络模式下，弹性网卡IP的回收时间，取值范围[300,15768000)

	ClaimExpiredSeconds *int64 `json:"ClaimExpiredSeconds,omitempty" name:"ClaimExpiredSeconds"`
}

type Instance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 节点角色, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, 默认为WORKER

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 实例异常(或者处于初始化中)的原因

	FailedReason *string `json:"FailedReason,omitempty" name:"FailedReason"`
	// 实例的状态（running 运行中，initializing 初始化中，failed 异常）

	InstanceState *string `json:"InstanceState,omitempty" name:"InstanceState"`
	// 实例是否封锁状态

	DrainStatus *string `json:"DrainStatus,omitempty" name:"DrainStatus"`
	// 节点配置

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 添加时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ResourceStatusItem struct {

	// 维度

	Dimension *string `json:"Dimension,omitempty" name:"Dimension"`
	// 总数

	TotalNum *uint64 `json:"TotalNum,omitempty" name:"TotalNum"`
	// 异常数

	AbnormalNum *uint64 `json:"AbnormalNum,omitempty" name:"AbnormalNum"`
	// 异常详情

	AbnormalDetail []*AbnormalDetail `json:"AbnormalDetail,omitempty" name:"AbnormalDetail"`
}

type DescribeClustersResourceStatusRequest struct {
	*tchttp.BaseRequest

	// 集群id，可传列表

	ClusterInstancesIds []*string `json:"ClusterInstancesIds,omitempty" name:"ClusterInstancesIds"`
	// 资源维度，可传列表

	Dimensions []*string `json:"Dimensions,omitempty" name:"Dimensions"`
	// 获取资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeClustersResourceStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersResourceStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterCertificatesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 每页大小

	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量

	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListClusterCertificatesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterCertificatesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterUpgradingStateResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterUpgradingStateResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterUpgradingStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Metadata struct {

	// api版本

	APIVersion *string `json:"APIVersion,omitempty" name:"APIVersion"`
	// app版本

	AppVersion *string `json:"AppVersion,omitempty" name:"AppVersion"`
	// 描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeClusterRoutesRequest struct {
	*tchttp.BaseRequest

	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *DescribeClusterRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TKEEdgeNodeResources struct {

	// 节点cpu配置

	CPU *float64 `json:"CPU,omitempty" name:"CPU"`
	// 节点memory配置

	Memory *float64 `json:"Memory,omitempty" name:"Memory"`
}

type DescribeClsLogSetsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClsLogSetsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsLogSetsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeCluster struct {

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// Vpc Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 集群pod cidr

	PodCIDR *string `json:"PodCIDR,omitempty" name:"PodCIDR"`
	// 集群 service cidr

	ServiceCIDR *string `json:"ServiceCIDR,omitempty" name:"ServiceCIDR"`
	// k8s 版本号

	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`
	// 集群状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群描述信息

	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	// 集群创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type Filter struct {

	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	// 当前只支持按照单个条件ClusterName进行过滤

	Name *string `json:"Name,omitempty" name:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type LogOutputOption struct {

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// topicID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 主机hsot

	Host *string `json:"Host,omitempty" name:"Host"`
	// 端口

	Port *string `json:"Port,omitempty" name:"Port"`
	// topic

	Topic *string `json:"Topic,omitempty" name:"Topic"`
	// LogsetId

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
}

type CreateClusterRouteTableResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterRouteTableResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRouteTableResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeChartDownloadInfoRequest struct {
	*tchttp.BaseRequest

	// Chart名称

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
	// Chart版本

	ChartVersion *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
}

func (r *DescribeChartDownloadInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChartDownloadInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群状态列表

		ClusterStatusSet []*ClusterStatus `json:"ClusterStatusSet,omitempty" name:"ClusterStatusSet"`
		// 集群个数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunClusterInspectionsRequest struct {
	*tchttp.BaseRequest

	// 目标集群列表

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *RunClusterInspectionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunClusterInspectionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceUpgradePreCheckResult struct {

	// 检查是否通过

	CheckPass *bool `json:"CheckPass,omitempty" name:"CheckPass"`
	// 检查项数组

	Items []*InstanceUpgradePreCheckResultItem `json:"Items,omitempty" name:"Items"`
	// 本节点独立pod列表

	SinglePods []*string `json:"SinglePods,omitempty" name:"SinglePods"`
}

type CreateClusterEndpointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicy struct {

	// 告警策略ID

	AlarmPolicyId *string `json:"AlarmPolicyId,omitempty" name:"AlarmPolicyId"`
	// k8s集群ID

	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`
	// k8s命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// k8s工作负载类型

	WorkloadType *string `json:"WorkloadType,omitempty" name:"WorkloadType"`
	// 告警策略settings

	AlarmPolicySettings *AlarmPolicySettings `json:"AlarmPolicySettings,omitempty" name:"AlarmPolicySettings"`
	// 告警通知settings

	NotifySettings *NotifySettings `json:"NotifySettings,omitempty" name:"NotifySettings"`
	// 告警屏蔽settings

	ShieldSettings *ShieldSettings `json:"ShieldSettings,omitempty" name:"ShieldSettings"`
}

type CreateClusterRequest struct {
	*tchttp.BaseRequest

	// CVM创建透传参数，json化字符串格式，详见[RunInstances](/tcloud/api/Compute/CVM/APIs/云服务器（cvm）/版本（2017-03-12）/实例相关接口/RunInstances)接口。总机型(包括地域)数量不超过10个，相同机型(地域)购买多台机器可以通过设置参数中RunInstances中InstanceCount来实现。

	RunInstancesForNode []*RunInstancesForNode `json:"RunInstancesForNode,omitempty" name:"RunInstancesForNode"`
	// 集群容器网络配置信息

	ClusterCIDRSettings *ClusterCIDRSettings `json:"ClusterCIDRSettings,omitempty" name:"ClusterCIDRSettings"`
	// 集群的基本配置信息

	ClusterBasicSettings *ClusterBasicSettings `json:"ClusterBasicSettings,omitempty" name:"ClusterBasicSettings"`
	// 集群高级配置信息

	ClusterAdvancedSettings *ClusterAdvancedSettings `json:"ClusterAdvancedSettings,omitempty" name:"ClusterAdvancedSettings"`
	// 节点高级配置信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 已存在实例的配置信息。所有实例必须在同一个VPC中，最大数量不超过100。

	ExistedInstancesForNode []*ExistedInstancesForNode `json:"ExistedInstancesForNode,omitempty" name:"ExistedInstancesForNode"`
	// CVM类型和其对应的数据盘挂载配置信息

	InstanceDataDiskMountSettings []*InstanceDataDiskMountSetting `json:"InstanceDataDiskMountSettings,omitempty" name:"InstanceDataDiskMountSettings"`
	// 集群的cpu架构，取值:   x86/arm

	ClusterArch *string `json:"ClusterArch,omitempty" name:"ClusterArch"`
}

func (r *CreateClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpgradeClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AlarmPolicySettings struct {

	// 告警策略名称

	AlarmPolicyName *string `json:"AlarmPolicyName,omitempty" name:"AlarmPolicyName"`
	// 告警策略描述

	AlarmPolicyDescription *string `json:"AlarmPolicyDescription,omitempty" name:"AlarmPolicyDescription"`
	// 告警类型，有cluster，node，pod三种类型。

	AlarmPolicyType *string `json:"AlarmPolicyType,omitempty" name:"AlarmPolicyType"`
	// 告警对象绑定类型。
	// all：全部对象绑定。
	// part：只有选择的部分对象绑定。

	AlarmObjectsType *string `json:"AlarmObjectsType,omitempty" name:"AlarmObjectsType"`
	// 告警指标列表。

	AlarmMetrics []*AlarmMetric `json:"AlarmMetrics,omitempty" name:"AlarmMetrics"`
	// 告警对象，多个用逗号分隔。

	AlarmObjects *string `json:"AlarmObjects,omitempty" name:"AlarmObjects"`
}

type InstanceDataDiskMountSetting struct {

	// CVM实例类型

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 数据盘挂载信息

	DataDisks []*DataDisk `json:"DataDisks,omitempty" name:"DataDisks"`
	// CVM实例所属可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
}

type DescribeRegionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeRegionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PauseClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *PauseClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PauseClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterCondition struct {

	// 集群创建过程类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 集群创建过程状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 最后一次探测到该状态的时间

	LastProbeTime *string `json:"LastProbeTime,omitempty" name:"LastProbeTime"`
	// 最后一次转换到该过程的时间

	LastTransitionTime *string `json:"LastTransitionTime,omitempty" name:"LastTransitionTime"`
	// 转换到该过程的简明原因

	Reason *string `json:"Reason,omitempty" name:"Reason"`
	// 转换到该过程的更多信息

	Message *string `json:"Message,omitempty" name:"Message"`
}

type NotReadyPodsItem struct {

	// 命名空间名称

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// pod列表

	Pods []*string `json:"Pods,omitempty" name:"Pods"`
}

type AcquireClusterAdminRoleRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *AcquireClusterAdminRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AcquireClusterAdminRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeHelmChartDetailRequest struct {
	*tchttp.BaseRequest

	// Chart名称

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
	// Chart版本

	ChartVersion *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
}

func (r *DescribeHelmChartDetailRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHelmChartDetailRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeQuotaResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 该账号当前地域支持的最大集群数量

		MaxClustersNum *uint64 `json:"MaxClustersNum,omitempty" name:"MaxClustersNum"`
		// 该账号当前地域单集群支持的最大节点数量

		MaxNodesNum *uint64 `json:"MaxNodesNum,omitempty" name:"MaxNodesNum"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQuotaResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeQuotaResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Tag struct {

	// 标签键

	Key *string `json:"Key,omitempty" name:"Key"`
	// 标签值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type Volume struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// pvc volume信息

	Pvc *PersistentVolumeClaimVolume `json:"Pvc,omitempty" name:"Pvc"`
	// cbs 信息

	TencentCbs *TencentCbsVolume `json:"TencentCbs,omitempty" name:"TencentCbs"`
	// secret

	Secret *SecretVolume `json:"Secret,omitempty" name:"Secret"`
	// Config Map Volume payload

	ConfigMapPayload []*PayLoad `json:"ConfigMapPayload,omitempty" name:"ConfigMapPayload"`
}

type CheckClusterHostNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点主机名校验是否通过(true: 通过 false: 不通过 )

		BEnable *bool `json:"BEnable,omitempty" name:"BEnable"`
		// 节点主机名校验不通过的原因

		ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterHostNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterHostNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 升级类型，目前只支持reset

	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`
	// create 表示开始一次升级任务
	// pause 表示停止任务
	// resume表示继续任务
	// abort表示终止任务

	Operation *string `json:"Operation,omitempty" name:"Operation"`
	// 需要升级的节点列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 当节点重新加入集群时候所使用的参数，参考添加已有节点接口

	ResetParam *UpgradeNodeResetParam `json:"ResetParam,omitempty" name:"ResetParam"`
	// 是否忽略节点升级前检查

	SkipPreCheck *bool `json:"SkipPreCheck,omitempty" name:"SkipPreCheck"`
	// 最大可容忍的不可用Pod比例

	MaxNotReadyPercent *float64 `json:"MaxNotReadyPercent,omitempty" name:"MaxNotReadyPercent"`
}

func (r *UpgradeClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpgradeClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterVirtualNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterVirtualNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterVirtualNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteLogCollectorRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志采集规则名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteLogCollectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteLogCollectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVirtualNodePoolsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟节点池数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 虚拟节点池数组

		NodePoolSet []*NodePoolSet `json:"NodePoolSet,omitempty" name:"NodePoolSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterVirtualNodePoolsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVirtualNodePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityPolicy struct {

	// 安全策略如192.168.0.0/24

	Policy *string `json:"Policy,omitempty" name:"Policy"`
	// 备注

	Description *string `json:"Description,omitempty" name:"Description"`
}

type AddExistedInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 失败的节点ID

		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`
		// 成功的节点ID

		SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds"`
		// 超时未返回出来节点的ID(可能失败，也可能成功)

		TimeoutInstanceIds []*string `json:"TimeoutInstanceIds,omitempty" name:"TimeoutInstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddExistedInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddExistedInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrainClusterNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 节点上pod列表

		Pods []*SimplePodInfo `json:"Pods,omitempty" name:"Pods"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DrainClusterNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DrainClusterNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAsGroupOptionAttributeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterAsGroupOptionAttributeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterAsGroupOptionAttributeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ChargeInfo struct {

	// pod是否计费

	Charge *bool `json:"Charge,omitempty" name:"Charge"`
	// pod计费开始时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// Pod的Uid

	Uid *string `json:"Uid,omitempty" name:"Uid"`
}

type CreateClusterRouteResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterRouteResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRouteResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AddExistedInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 实例额外需要设置参数信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。（目前仅支持设置单个sgId）

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 重装系统时，可以指定修改实例的HostName(集群为HostName模式时，此参数必传，规则名称除不支持大写字符外与[RunInstances](/tcloud/api/Compute/CVM/APIs/云服务器（cvm）/版本（2017-03-12）/实例相关接口/RunInstances)接口HostName一致)

	HostName *string `json:"HostName,omitempty" name:"HostName"`
	// 节点池选项

	NodePool *NodePoolOption `json:"NodePool,omitempty" name:"NodePool"`
}

func (r *AddExistedInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddExistedInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLogCollectorRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志采集规则名字

	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *GetLogCollectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLogCollectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetLogCollectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志采集规则

		LogCollector *LogCollector `json:"LogCollector,omitempty" name:"LogCollector"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetLogCollectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetLogCollectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScaleMetrics struct {

	// 指标名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 指标值

	Value *string `json:"Value,omitempty" name:"Value"`
}

type DescribeClusterAsGroupOptionRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterAsGroupOptionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAsGroupOptionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterAsGroup struct {

	// 伸缩组ID

	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
	// 伸缩组状态(开启 enabled 开启中 enabling 关闭 disabled 关闭中 disabling 更新中 updating 删除中 deleting 开启缩容中 scaleDownEnabling 关闭缩容中 scaleDownDisabling)

	Status *string `json:"Status,omitempty" name:"Status"`
	// 节点是否设置成不可调度

	IsUnschedulable *bool `json:"IsUnschedulable,omitempty" name:"IsUnschedulable"`
	// 伸缩组的label列表

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ClusterNetworkSettings struct {

	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
	// 是否忽略 ClusterCIDR 冲突错误, 默认不忽略

	IgnoreClusterCIDRConflict *bool `json:"IgnoreClusterCIDRConflict,omitempty" name:"IgnoreClusterCIDRConflict"`
	// 集群中每个Node上最大的Pod数量(默认为256)

	MaxNodePodNum *uint64 `json:"MaxNodePodNum,omitempty" name:"MaxNodePodNum"`
	// 集群最大的service数量(默认为256)

	MaxClusterServiceNum *uint64 `json:"MaxClusterServiceNum,omitempty" name:"MaxClusterServiceNum"`
	// 是否启用IPVS(默认不开启)

	Ipvs *bool `json:"Ipvs,omitempty" name:"Ipvs"`
	// 集群的VPCID（如果创建空集群，为必传值，否则自动设置为和集群的节点保持一致）

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 网络插件是否启用CNI(默认开启)

	Cni *bool `json:"Cni,omitempty" name:"Cni"`
	// 集群网络对应子网

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
}

type ClusterResource struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 子网资源列表

	SubnetResources []*SubnetResource `json:"SubnetResources,omitempty" name:"SubnetResources"`
	// 集群内Pod数量配额

	ClusterPodQuota *int64 `json:"ClusterPodQuota,omitempty" name:"ClusterPodQuota"`
	// 集群内已创建的Pod总数量

	ClusterPodNum *int64 `json:"ClusterPodNum,omitempty" name:"ClusterPodNum"`
}

type Result struct {

	// chart

	Chart *string `json:"Chart,omitempty" name:"Chart"`
	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 分数

	Score *int64 `json:"Score,omitempty" name:"Score"`
}

type DeleteClusterAsGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterAsGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterAsGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *DescribeCcnRoutesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRoutesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAsGroupAttributeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群关联的伸缩组属性

	ClusterAsGroupAttribute *ClusterAsGroupAttribute `json:"ClusterAsGroupAttribute,omitempty" name:"ClusterAsGroupAttribute"`
}

func (r *ModifyClusterAsGroupAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterAsGroupAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInspectionProgress struct {

	// 当前步骤名称：
	// init_env: 初始化环境，安装agent。
	// init_k8s_resources：获取kubernetes资源
	// init_components: 获取核心组件参数
	// init_machines：获取节点系统参数
	// diagnostic： 正在检查集群

	Step *string `json:"Step,omitempty" name:"Step"`
	// 检查进度百分比

	Percent *float64 `json:"Percent,omitempty" name:"Percent"`
}

type ReservedInstance struct {

	// 预留实例ID

	ReservedInstanceId *string `json:"ReservedInstanceId,omitempty" name:"ReservedInstanceId"`
	// 预留实例名称

	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" name:"ReservedInstanceName"`
	// 预留券状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 有效期，单位：月

	TimeSpan *uint64 `json:"TimeSpan,omitempty" name:"TimeSpan"`
	// 抵扣资源类型

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 资源核数

	Cpu *float64 `json:"Cpu,omitempty" name:"Cpu"`
	// 资源内存，单位：Gi

	Memory *float64 `json:"Memory,omitempty" name:"Memory"`
	// 预留券的范围，默认值region。

	Scope *string `json:"Scope,omitempty" name:"Scope"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 生效时间

	ActiveAt *string `json:"ActiveAt,omitempty" name:"ActiveAt"`
	// 过期时间

	ExpireAt *string `json:"ExpireAt,omitempty" name:"ExpireAt"`
}

type CheckClusterAdminRoleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterAdminRoleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterAdminRoleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 各集群巡检概览

		ResultSet []*DescribeClusterInspectionItem `json:"ResultSet,omitempty" name:"ResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInspectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInspectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableVpcPeerClusterRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableVpcPeerClusterRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableVpcPeerClusterRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunClusterInspectionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 每个集群的处理结果

		ResultSet []*RunClusterInspectionResponseItem `json:"ResultSet,omitempty" name:"ResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RunClusterInspectionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RunClusterInspectionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterAsGroupOption struct {

	// 是否开启缩容

	IsScaleDownEnabled *bool `json:"IsScaleDownEnabled,omitempty" name:"IsScaleDownEnabled"`
	// 多伸缩组情况下扩容选择算法(random 随机选择，most-pods 最多类型的Pod least-waste 最少的资源浪费，默认为random)

	Expander *string `json:"Expander,omitempty" name:"Expander"`
	// 最大并发缩容数

	MaxEmptyBulkDelete *int64 `json:"MaxEmptyBulkDelete,omitempty" name:"MaxEmptyBulkDelete"`
	// 集群扩容后多少分钟开始判断缩容（默认为10分钟）

	ScaleDownDelay *int64 `json:"ScaleDownDelay,omitempty" name:"ScaleDownDelay"`
	// 节点连续空闲多少分钟后被缩容（默认为 10分钟）

	ScaleDownUnneededTime *int64 `json:"ScaleDownUnneededTime,omitempty" name:"ScaleDownUnneededTime"`
	// 节点资源使用量低于多少(百分比)时认为空闲(默认: 50(百分比))

	ScaleDownUtilizationThreshold *int64 `json:"ScaleDownUtilizationThreshold,omitempty" name:"ScaleDownUtilizationThreshold"`
	// 含有本地存储Pod的节点是否不缩容(默认： FALSE)

	SkipNodesWithLocalStorage *bool `json:"SkipNodesWithLocalStorage,omitempty" name:"SkipNodesWithLocalStorage"`
	// 含有kube-system namespace下非DaemonSet管理的Pod的节点是否不缩容 (默认： FALSE)

	SkipNodesWithSystemPods *bool `json:"SkipNodesWithSystemPods,omitempty" name:"SkipNodesWithSystemPods"`
	// 计算资源使用量时是否默认忽略DaemonSet的实例(默认值: False，不忽略)

	IgnoreDaemonSetsUtilization *bool `json:"IgnoreDaemonSetsUtilization,omitempty" name:"IgnoreDaemonSetsUtilization"`
	// unready节点最大占比

	MaxTotalUnreadyPercentage *int64 `json:"MaxTotalUnreadyPercentage,omitempty" name:"MaxTotalUnreadyPercentage"`
	// unready节点累计数量

	OkTotalUnreadyCount *int64 `json:"OkTotalUnreadyCount,omitempty" name:"OkTotalUnreadyCount"`
	// unready节点缩容时间

	ScaleDownUnreadyTime *int64 `json:"ScaleDownUnreadyTime,omitempty" name:"ScaleDownUnreadyTime"`
	// 未注册成功节点移除耗时

	UnregisteredNodeRemovalTime *int64 `json:"UnregisteredNodeRemovalTime,omitempty" name:"UnregisteredNodeRemovalTime"`
}

type DeleteClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 成功删除的cvm实例id列表

		SuccInstanceIds []*string `json:"SuccInstanceIds,omitempty" name:"SuccInstanceIds"`
		// 未成功删除的cvm实例id列表

		FailedInstanceIds []*string `json:"FailedInstanceIds,omitempty" name:"FailedInstanceIds"`
		// 未找到的cvm实例id列表

		NotFoundInstanceIds []*string `json:"NotFoundInstanceIds,omitempty" name:"NotFoundInstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUpgradeClusterProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群升级整体进度

		Steps []*TaskStepInfo `json:"Steps,omitempty" name:"Steps"`
		// 每个节点的进度

		Instances []*TaskStepInfo `json:"Instances,omitempty" name:"Instances"`
		// 上一次的报错信息

		LastError *string `json:"LastError,omitempty" name:"LastError"`
		// 任务当前的状态

		LifeState *string `json:"LifeState,omitempty" name:"LifeState"`
		// 总数量

		Total *int64 `json:"Total,omitempty" name:"Total"`
		// 已完成节点数量

		Done *int64 `json:"Done,omitempty" name:"Done"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUpgradeClusterProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpgradeClusterProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetDashboardIDRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 仪表盘类型

	DashboardType *string `json:"DashboardType,omitempty" name:"DashboardType"`
}

func (r *GetDashboardIDRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetDashboardIDRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnInstance struct {

	// 云联网实例ID

	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
	// 关联实例类型：
	// VPC：私有网络
	// DIRECTCONNECT：专线网关
	// BMVPC：黑石私有网络

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 关联实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 关联实例所属大区，例如：ap-guangzhou

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 关联实例所属UIN（根账号

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
	// 关联实例CIDR

	Cidrs []*string `json:"Cidrs,omitempty" name:"Cidrs"`
	// 关联实例状态：
	// PENDING：申请中
	// ACTIVE：已连接
	// EXPIRED：已过期
	// REJECTED：已拒绝
	// DELETED：已删除
	// FAILED：失败的（2小时后将异步强制解关联）
	// ATTACHING：关联中
	// DETACHING：解关联中
	// DETACHFAILED：解关联失败（2小时后将异步强制解关联）

	State *string `json:"State,omitempty" name:"State"`
	// 云联网所属UIN（根账号）

	CcnUin *string `json:"CcnUin,omitempty" name:"CcnUin"`
}

type ClusterAlarmSetting struct {

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 是否设置了监控告警

	BMonitor *bool `json:"BMonitor,omitempty" name:"BMonitor"`
}

type CreateClusterAsGroupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 启动配置ID

		LaunchConfigurationId *string `json:"LaunchConfigurationId,omitempty" name:"LaunchConfigurationId"`
		// 伸缩组ID

		AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterAsGroupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterAsGroupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ListClusterCertificatesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 证书总数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 证书对象数组

		CertificateInfos []*CertificateInfo `json:"CertificateInfos,omitempty" name:"CertificateInfos"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListClusterCertificatesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ListClusterCertificatesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 需要升级的集群的 clusterId 字段

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 升级的节点ID列表，请选择集群中需要升级的节点ID列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 描述了实例创建相关的一些高级设置。

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 描述了实例登录相关配置与信息。

	LoginSettings *string `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 描述了实例的增强服务启用情况与其设置，如云安全，云监控等实例 Agent。

	EnhancedService *string `json:"EnhancedService,omitempty" name:"EnhancedService"`
}

func (r *UpdateClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
}

func (r *DescribeCcnInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterEndpointSPResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClusterEndpointSPResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterEndpointSPResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClustersRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserClustersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserClustersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallLogAgentRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *UninstallLogAgentRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallLogAgentRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInfo struct {

	// 集群描述

	ClusterDescription *string `json:"ClusterDescription,omitempty" name:"ClusterDescription"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群类型

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
}

type ClusterInstancesVersion struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 节点版本统计信息

	InstanceVersions []*ClusterInstancesVersionItem `json:"InstanceVersions,omitempty" name:"InstanceVersions"`
	// 出错信息

	Error *string `json:"Error,omitempty" name:"Error"`
	// 是否存在可升级节点

	UpgradeAble *bool `json:"UpgradeAble,omitempty" name:"UpgradeAble"`
}

type Flag struct {

	// 参数名

	Name *string `json:"Name,omitempty" name:"Name"`
	// 参数类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 参数描述

	Usage *string `json:"Usage,omitempty" name:"Usage"`
	// 参数默认值

	Default *string `json:"Default,omitempty" name:"Default"`
	// 参数可选范围（目前包含range和in两种，"[]"代表range，如"[1, 5]"表示参数必须>=1且 <=5, "()"代表in， 如"('aa', 'bb')"表示参数只能为字符串'aa'或者'bb'，该参数为空表示不校验）

	Constraint *string `json:"Constraint,omitempty" name:"Constraint"`
}

type AddClusterCIDRToCcnRequest struct {
	*tchttp.BaseRequest

	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 用于分配集群容器和服务 IP 的 CIDR，不得与 VPC CIDR 冲突，也不得与同 VPC 内其他集群 CIDR 冲突

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
}

func (r *AddClusterCIDRToCcnRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterCIDRToCcnRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckInstancesUpgradeAbleRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 升级类型

	UpgradeType *string `json:"UpgradeType,omitempty" name:"UpgradeType"`
	// 分页Offset

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 分页Limit

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 过滤

	Filter []*Filter `json:"Filter,omitempty" name:"Filter"`
}

func (r *CheckInstancesUpgradeAbleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckInstancesUpgradeAbleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterCreateProgressRequest struct {
	*tchttp.BaseRequest

	// 集群 Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterCreateProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterCreateProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCosInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCosInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterAttributeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群描述

	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	// 集群运行时名称

	ClusterRunTime *string `json:"ClusterRunTime,omitempty" name:"ClusterRunTime"`
}

func (r *ModifyClusterAttributeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterAttributeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInternalLB struct {

	// 是否开启内网访问LB

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 内网访问LB关联的子网Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type ZoneResource struct {

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可创建的规格总数

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type AddVpcCniSubnetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddVpcCniSubnetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddVpcCniSubnetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpdateLogCollectorRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志规则名字

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志规则描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 输入类型

	InputType *string `json:"InputType,omitempty" name:"InputType"`
	// 输出类型

	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
	// 输入选项

	InputOption *LogInputOption `json:"InputOption,omitempty" name:"InputOption"`
	// 输出选项

	OutputOption *LogOutputOption `json:"OutputOption,omitempty" name:"OutputOption"`
}

func (r *UpdateLogCollectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UpdateLogCollectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterVirtualNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟节点数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 虚拟节点数组

		Nodes []*Node `json:"Nodes,omitempty" name:"Nodes"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterVirtualNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterVirtualNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterEndpointSPRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 安全策略放通单个IP或CIDR(例如: "192.168.1.0/24",默认为拒绝所有)

	SecurityPolicies []*string `json:"SecurityPolicies,omitempty" name:"SecurityPolicies"`
}

func (r *ModifyClusterEndpointSPRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterEndpointSPRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TencentCbsVolume struct {

	// cbs id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// fs类型

	FsType *string `json:"FsType,omitempty" name:"FsType"`
	// 是否只读

	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

type DescribeChartDownloadInfoResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 下载地址

		PreSignedDownloadURL *string `json:"PreSignedDownloadURL,omitempty" name:"PreSignedDownloadURL"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeChartDownloadInfoResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeChartDownloadInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CcnRoute struct {

	// 路由策略ID

	RouteId *string `json:"RouteId,omitempty" name:"RouteId"`
	// 目的端

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 下一跳类型（关联实例类型），所有类型：VPC、DIRECTCONNECT

	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`
	// 下一跳（关联实例）

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 下一跳所属地域（关联实例所属地域）

	InstanceRegion *string `json:"InstanceRegion,omitempty" name:"InstanceRegion"`
	// 关联实例所属UIN（根账号）

	InstanceUin *string `json:"InstanceUin,omitempty" name:"InstanceUin"`
}

type DeleteClusterVirtualNodePoolResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterVirtualNodePoolResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterVirtualNodePoolResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UpgradeAbleInstancesItem struct {

	// 节点Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 节点的当前版本

	Version *string `json:"Version,omitempty" name:"Version"`
}

type DescribeClusterInspectionOverviewsRequest struct {
	*tchttp.BaseRequest

	// 集群实例id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DescribeClusterInspectionOverviewsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInspectionOverviewsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Evaluator struct {

	// 告警判断类型，目前支持gt和lt两种

	Type *string `json:"Type,omitempty" name:"Type"`
	// 告警设置的阈值

	Value *float64 `json:"Value,omitempty" name:"Value"`
}

type EnableLogCollectorResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EnableLogCollectorResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableLogCollectorResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCosInfoRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeCosInfoRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCosInfoRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeClusterProgressResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群升级整体进度

		Steps *TaskStepInfo `json:"Steps,omitempty" name:"Steps"`
		// 每个节点的进度

		Instances *TaskStepInfo `json:"Instances,omitempty" name:"Instances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUpgradeClusterProgressResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeClusterProgressResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ContainerProbeResult struct {

	// 探针类型可用的取值为 liveness 或者 readiness

	ProbeType *string `json:"ProbeType,omitempty" name:"ProbeType"`
	// 探针检查结果。在未配置probe规则时为true，在配置规则但是由于通讯问题没有值时为空指针。

	Ready *bool `json:"Ready,omitempty" name:"Ready"`
	// 一个检查结果的有效秒数。一般设置成PeriodSeconds的5倍，由客户端每次set时刷新

	TTL *uint64 `json:"TTL,omitempty" name:"TTL"`
}

type AddClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群 ID，请填写 查询集群列表 接口中返回的 clusterId 字段

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 数据盘挂载点, 默认不挂载数据盘. 已格式化的 ext3，ext4，xfs 文件系统的数据盘将直接挂载，其他文件系统或未格式化的数据盘将自动格式化为ext4 并挂载，请注意备份数据! 无数据盘或有多块数据盘的云主机此设置不生效。

	MountTarget []*string `json:"MountTarget,omitempty" name:"MountTarget"`
	// dockerd --graph 指定值, 默认为 /var/lib/docker

	DockerGraphPath []*string `json:"DockerGraphPath,omitempty" name:"DockerGraphPath"`
	// base64 编码的用户脚本, 此脚本会在 k8s 组件运行后执行, 需要用户保证脚本的可重入及重试逻辑, 脚本及其生成的日志文件可在节点的 /data/ccs_userscript/ 路径查看, 如果要求节点需要在进行初始化完成后才可加入调度, 可配合 unschedulable 参数使用, 在 userScript 最后初始化完成后, 添加 kubectl uncordon nodename --kubeconfig=/root/.kube/config 命令使节点加入调度

	UserScript []*string `json:"UserScript,omitempty" name:"UserScript"`
	// 系统名。centos7.2x86_64 或者 ubuntu16.04.1 LTSx86_64，仅当新建集群为空集群, 第一次向空集群添加节点时需要指定. 当集群系统确定后, 后续添加的节点都是集群系统

	OsName *string `json:"OsName,omitempty" name:"OsName"`
	// 设置加入的节点是否参与调度，默认值为0，表示参与调度；非0表示不参与调度, 待节点初始化完成之后, 可执行kubectl uncordon nodename使node加入调度.

	Unschedulable *int64 `json:"Unschedulable,omitempty" name:"Unschedulable"`
	// cvm创建透传参数，json化字符串格式

	CvmRunInstances *string `json:"CvmRunInstances,omitempty" name:"CvmRunInstances"`
}

func (r *AddClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeUserClustersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总集群数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群数组

		Clusters []*ClusterInfo `json:"Clusters,omitempty" name:"Clusters"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserClustersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUserClustersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePolicyInfo struct {

	// 策略ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// AMP系统主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 策略类型

	AlarmPolicyType *string `json:"AlarmPolicyType,omitempty" name:"AlarmPolicyType"`
	// ARGUS系统策略ID列表

	ArgusPolicyIds []*uint64 `json:"ArgusPolicyIds,omitempty" name:"ArgusPolicyIds"`
}

type EnhancedService struct {

	// 开启云安全服务。若不指定该参数，则默认开启云安全服务。

	SecurityService *RunSecurityServiceEnabled `json:"SecurityService,omitempty" name:"SecurityService"`
	// 开启云安全服务。若不指定该参数，则默认开启云监控服务。

	MonitorService *RunMonitorServiceEnabled `json:"MonitorService,omitempty" name:"MonitorService"`
}

type IPAddress struct {

	// Ip 地址的类型。可为 advertise, public 等

	Type *string `json:"Type,omitempty" name:"Type"`
	// Ip 地址

	Ip *string `json:"Ip,omitempty" name:"Ip"`
	// 网络端口

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

type PayLoad struct {

	// configmap base64 encoded data

	Data *string `json:"Data,omitempty" name:"Data"`
	// configmap file user visible path

	Key *string `json:"Key,omitempty" name:"Key"`
}

type CreateClusterAsGroupRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 伸缩组创建透传参数，json化字符串格式，详见接口[CreateAutoScalingGroup](/tcloud/api/Compute/AS/APIs/弹性伸缩（as）/版本（2018-04-19）/伸缩组相关接口/CreateAutoScalingGroup)。LaunchConfigurationId由LaunchConfigurePara参数创建，不支持填写

	AutoScalingGroupPara *string `json:"AutoScalingGroupPara,omitempty" name:"AutoScalingGroupPara"`
	// 启动配置创建透传参数，json化字符串格式，详见[CreateLaunchConfiguration](/tcloud/api/Compute/AS/APIs/弹性伸缩（as）/版本（2018-04-19）/启动配置相关接口/CreateLaunchConfiguration)接口。另外ImageId参数由于集群维度已经有的ImageId信息，这个字段不需要填写。UserData字段设置通过UserScript设置，这个字段不需要填写。

	LaunchConfigurePara *string `json:"LaunchConfigurePara,omitempty" name:"LaunchConfigurePara"`
	// 节点高级配置信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 节点Label数组

	Labels []*Label `json:"Labels,omitempty" name:"Labels"`
}

func (r *CreateClusterAsGroupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterAsGroupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EdgeClusterInternalLB struct {

	// 是否开启内网访问LB

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
	// 内网访问LB关联的子网Id

	SubnetId []*string `json:"SubnetId,omitempty" name:"SubnetId"`
}

type DescribeHelmChartVersionRequest struct {
	*tchttp.BaseRequest

	// Chart名称

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
}

func (r *DescribeHelmChartVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHelmChartVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetVbcInstanceRequest struct {
	*tchttp.BaseRequest

	// vpc唯一id

	UnVpcId *string `json:"UnVpcId,omitempty" name:"UnVpcId"`
}

func (r *GetVbcInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetVbcInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterHealthyStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群Pod健康状态

		PodsStatus *ClusterHealthyPodsStatus `json:"PodsStatus,omitempty" name:"PodsStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterHealthyStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterHealthyStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OpUpgradeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 节点实例ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 操作类型
	// pause 暂停任务（任务必须处于运行中）
	// resume 恢复任务（任务必须处于已暂停状态）
	// abort 停止任务（任务必须处于已暂停状态）

	Operation *string `json:"Operation,omitempty" name:"Operation"`
}

func (r *OpUpgradeClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *OpUpgradeClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PersistentVolumeClaimVolume struct {

	// pvc名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// cbs id

	DiskId *string `json:"DiskId,omitempty" name:"DiskId"`
	// fs 类型

	FsType *string `json:"FsType,omitempty" name:"FsType"`
	// 是否只读

	ReadOnly *bool `json:"ReadOnly,omitempty" name:"ReadOnly"`
}

type TagSpecification struct {

	// 标签绑定的资源类型，当前支持类型："cluster"

	ResourceType *string `json:"ResourceType,omitempty" name:"ResourceType"`
	// 标签对列表

	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateClusterRouteRequest struct {
	*tchttp.BaseRequest

	// 路由表名称。

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
	// 目的端CIDR。

	DestinationCidrBlock *string `json:"DestinationCidrBlock,omitempty" name:"DestinationCidrBlock"`
	// 下一跳地址。

	GatewayIp *string `json:"GatewayIp,omitempty" name:"GatewayIp"`
}

func (r *CreateClusterRouteRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterRouteRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 主机InstanceId列表

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 集群实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）

	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
	// 是否强制删除(当节点在初始化时，可以指定参数为TRUE)

	ForceDelete *bool `json:"ForceDelete,omitempty" name:"ForceDelete"`
}

func (r *DeleteClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群中实例总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群中实例列表

		InstanceSet []*Instance `json:"InstanceSet,omitempty" name:"InstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
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

type DescribeUpgradeClusterProgressRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// cluster/node

	Type *string `json:"Type,omitempty" name:"Type"`
}

func (r *DescribeUpgradeClusterProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeUpgradeClusterProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PodDetail struct {

	// pod唯一ID

	EksId *string `json:"EksId,omitempty" name:"EksId"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 工作负载类型

	Kind *string `json:"Kind,omitempty" name:"Kind"`
	// 工作负载名称

	KindName *string `json:"KindName,omitempty" name:"KindName"`
	// pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// pod的uid

	Uid *string `json:"Uid,omitempty" name:"Uid"`
	// 记录时间

	RecordTime *string `json:"RecordTime,omitempty" name:"RecordTime"`
}

type SimplePodInfo struct {

	// pod名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// pod所属命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 所属节点IP

	NodeIp *string `json:"NodeIp,omitempty" name:"NodeIp"`
}

type CheckClusterImageResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否支持设置为集群镜像

		Suitable *bool `json:"Suitable,omitempty" name:"Suitable"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterImageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterImageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Banner struct {

	// 序号

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 展示内容

	Content *string `json:"Content,omitempty" name:"Content"`
	// 展示页面

	Tags *string `json:"Tags,omitempty" name:"Tags"`
}

type LogCollector struct {

	// 日志收集名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志收集描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 输入类型

	InputType *string `json:"InputType,omitempty" name:"InputType"`
	// 输出类型

	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
	// 输入选项

	InputOption *LogInputOption `json:"InputOption,omitempty" name:"InputOption"`
	// 输出选项

	OutputOption *LogOutputOption `json:"OutputOption,omitempty" name:"OutputOption"`
}

type CreateClusterVirtualNodeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 虚拟节点池id

	NodePoolId *string `json:"NodePoolId,omitempty" name:"NodePoolId"`
	// 子网ID数组,与SubnetId必写一个

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// 子网id，与SubnetIds必写一个

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

func (r *CreateClusterVirtualNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterVirtualNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterKubeconfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 用户的kubeconfig

		Kubeconfig *string `json:"Kubeconfig,omitempty" name:"Kubeconfig"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterKubeconfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterKubeconfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckClusterCIDRResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否存在CIDR冲突。

		IsConflict *bool `json:"IsConflict,omitempty" name:"IsConflict"`
		// CIDR冲突的类型("CIDR_CONFLICT_WITH_OTHER_CLUSTER" 同VPC其他集群CIDR存在冲突
		// "CIDR_CONFLICT_WITH_VPC_CIDR" 与VPC的CIDR存在冲突
		// "CIDR_CONFLICT_WITH_VPC_GLOBAL_ROUTE" 与同VPC的全局路由存在冲突)。

		ConflictType *string `json:"ConflictType,omitempty" name:"ConflictType"`
		// CIDR冲突描述信息。

		ConflictMsg *string `json:"ConflictMsg,omitempty" name:"ConflictMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckClusterCIDRResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterCIDRResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RunMonitorServiceEnabled struct {

	// 是否开启[云监控](/tcloud/omtool/BARAD)服务。取值范围：<br><li>TRUE：表示开启云监控服务<br><li>FALSE：表示不开启云监控服务<br><br>默认取值：TRUE。

	Enabled *bool `json:"Enabled,omitempty" name:"Enabled"`
}

type CheckLogCollectorHostPathRequest struct {
	*tchttp.BaseRequest

	// 路径名

	Path *string `json:"Path,omitempty" name:"Path"`
}

func (r *CheckLogCollectorHostPathRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckLogCollectorHostPathRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceCreateProgressRequest struct {
	*tchttp.BaseRequest

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceCreateProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceCreateProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRouteTableConflictsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 路由表是否冲突。

		HasConflict *bool `json:"HasConflict,omitempty" name:"HasConflict"`
		// 路由表冲突列表。

		RouteTableConflictSet []*RouteTableConflict `json:"RouteTableConflictSet,omitempty" name:"RouteTableConflictSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRouteTableConflictsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRouteTableConflictsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterHealthyPodsStatus struct {

	// pod总数

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// NotReady的pod总数

	NotReadyTotal *int64 `json:"NotReadyTotal,omitempty" name:"NotReadyTotal"`
	// NotReady的pod详细信息

	NotReadyPods []*NotReadyPodsItem `json:"NotReadyPods,omitempty" name:"NotReadyPods"`
}

type AddClusterCIDRToCcnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网路由数组

		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AddClusterCIDRToCcnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *AddClusterCIDRToCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterExtraArgsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群自定义参数

		ClusterExtraArgs *ClusterExtraArgs `json:"ClusterExtraArgs,omitempty" name:"ClusterExtraArgs"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterExtraArgsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInspectionReportResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 健康检查项目

		ResultSet []*InspectionReportItem `json:"ResultSet,omitempty" name:"ResultSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterInspectionReportResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInspectionReportResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterBasicSettings struct {

	// 集群系统。centos7.2x86_64 或者 ubuntu16.04.1 LTSx86_64，默认取值为ubuntu16.04.1 LTSx86_64

	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`
	// 集群版本,默认值为1.10.5

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群描述

	ClusterDescription *string `json:"ClusterDescription,omitempty" name:"ClusterDescription"`
	// 私有网络ID，形如vpc-xxx。创建托管空集群时必传。

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 集群内新增资源所属项目ID。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到集群实例。

	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
	// 容器的镜像版本，"DOCKER_CUSTOMIZE"(容器定制版),"GENERAL"(普通版本，默认值)

	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`
	// 是否开启节点的默认安全组(默认: 否，Aphla特性)

	NeedWorkSecurityGroup *bool `json:"NeedWorkSecurityGroup,omitempty" name:"NeedWorkSecurityGroup"`
}

type InstanceBaseSettings struct {

	// 节点InstanceId

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 节点名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 节点所在Zone信息(string)

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 节点外网IP

	WanIp *string `json:"WanIp,omitempty" name:"WanIp"`
	// 节点内网IP

	LanIp []*string `json:"LanIp,omitempty" name:"LanIp"`
	// 节点CPU（单位：毫核）

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 节点内存(单位： M)

	Mem *string `json:"Mem,omitempty" name:"Mem"`
	// 节点GPU信息

	Gpu *uint64 `json:"Gpu,omitempty" name:"Gpu"`
	// 节点内核版本

	KernelVersion *string `json:"KernelVersion,omitempty" name:"KernelVersion"`
	// 节点操作系统镜像

	OsImage *string `json:"OsImage,omitempty" name:"OsImage"`
	// 节点是否正常标识(0 异常 1 正常 2 创建中)

	IsNormal *uint64 `json:"IsNormal,omitempty" name:"IsNormal"`
	// 异常原因

	AbnormalReason *string `json:"AbnormalReason,omitempty" name:"AbnormalReason"`
	// K8s版本

	K8sVersion *string `json:"K8sVersion,omitempty" name:"K8sVersion"`
	// 节点角色, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, 默认为WORKER

	NodeRole *string `json:"NodeRole,omitempty" name:"NodeRole"`
	// 节点Label标签列表

	LabelValues []*Label `json:"LabelValues,omitempty" name:"LabelValues"`
	// 是否不可以被调度（0 表示可以被调度， 大于0 表示不可以被调度）

	Unschedulable *uint64 `json:"Unschedulable,omitempty" name:"Unschedulable"`
}

type LogInputOption struct {

	// 路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 标签 类型为map 键值对

	Labels *Label `json:"Labels,omitempty" name:"Labels"`
	// 是否是全部命名空间

	AllNamespaces *bool `json:"AllNamespaces,omitempty" name:"AllNamespaces"`
	// 命名空间

	Namespaces *LogInputOptNamespace `json:"Namespaces,omitempty" name:"Namespaces"`
}

type DeleteClusterEndpointVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterEndpointVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterEndpointVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DisableClusterAuditRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *DisableClusterAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DisableClusterAuditRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrainClusterVirtualNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟节点数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 虚拟节点数组

		Pods []*SimplePodInfo `json:"Pods,omitempty" name:"Pods"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DrainClusterVirtualNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DrainClusterVirtualNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteClusterCIDRFromCcnResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网路由数组

		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClusterCIDRFromCcnResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterCIDRFromCcnResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AbnormalDetail struct {

	// kubernetes namespace

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// kubernetes资源名称

	ResourceName *string `json:"ResourceName,omitempty" name:"ResourceName"`
}

type AlarmPolicyFilter struct {

	// 告警策略名称

	AlarmPolicyName *string `json:"AlarmPolicyName,omitempty" name:"AlarmPolicyName"`
}

type CheckClusterCIDRRequest struct {
	*tchttp.BaseRequest

	// 集群的vpc-id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 集群的CIDR

	ClusterCIDR *string `json:"ClusterCIDR,omitempty" name:"ClusterCIDR"`
	// 集群的网络类型

	NetworkType *string `json:"NetworkType,omitempty" name:"NetworkType"`
}

func (r *CheckClusterCIDRRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterCIDRRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群CA证书

		CertificationAuthority *string `json:"CertificationAuthority,omitempty" name:"CertificationAuthority"`
		// 集群外网访问地址（包含端口）

		ClusterExternalEndpoint *string `json:"ClusterExternalEndpoint,omitempty" name:"ClusterExternalEndpoint"`
		// 集群内网访问地址（包含端口）

		ClusterIntranetEndpoint *string `json:"ClusterIntranetEndpoint,omitempty" name:"ClusterIntranetEndpoint"`
		// 集群默认访问域名

		ClusterDomain *string `json:"ClusterDomain,omitempty" name:"ClusterDomain"`
		// 集群外网访问域名

		ClusterExternalDomain *string `json:"ClusterExternalDomain,omitempty" name:"ClusterExternalDomain"`
		// 集群内外访问域名

		ClusterIntranetDomain *string `json:"ClusterIntranetDomain,omitempty" name:"ClusterIntranetDomain"`
		// 集群外网访问使用的安全组

		SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterEndpointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterEndpointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrainClusterVirtualNodeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点名称

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// true 反馈节点pod信息，false驱逐节点

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *DrainClusterVirtualNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DrainClusterVirtualNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ForwardRequestRequest struct {
	*tchttp.BaseRequest

	// 请求tke-apiserver http请求对应的方法

	Method *string `json:"Method,omitempty" name:"Method"`
	// 请求tke-apiserver  http请求访问路径

	Path *string `json:"Path,omitempty" name:"Path"`
	// 请求tke-apiserver http头中Accept参数

	Accept *string `json:"Accept,omitempty" name:"Accept"`
	// 请求tke-apiserver http头中ContentType参数

	ContentType *string `json:"ContentType,omitempty" name:"ContentType"`
	// 请求tke-apiserver http请求body信息

	RequestBody *string `json:"RequestBody,omitempty" name:"RequestBody"`
	// 请求tke-apiserver http头中X-TKE-ClusterName参数

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 是否编码请求body信息

	EncodedBody *bool `json:"EncodedBody,omitempty" name:"EncodedBody"`
	// 鉴权使用

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ForwardRequestRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ForwardRequestRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InspectionReportItem struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
}

type LoginSettings struct {

	// 实例登录密码。不同操作系统类型密码复杂度限制不一样，具体如下：<br><li>Linux实例密码必须8到16位，至少包括两项[a-z，A-Z]、[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = | { } [ ] : ; ' , . ? / ]中的特殊符号。<br><li>Windows实例密码必须12到16位，至少包括三项[a-z]，[A-Z]，[0-9] 和 [( ) ` ~ ! @ # $ % ^ & * - + = { } [ ] : ; ' , . ? /]中的特殊符号。<br><br>若不指定该参数，则由系统随机生成密码，并通过站内信方式通知到用户。

	Password *string `json:"Password,omitempty" name:"Password"`
	// 密钥ID列表。关联密钥后，就可以通过对应的私钥来访问实例；KeyId可通过接口DescribeKeyPairs获取，密钥与密码不能同时指定，同时Windows操作系统不支持指定密钥。当前仅支持购买的时候指定一个密钥。

	KeyIds []*string `json:"KeyIds,omitempty" name:"KeyIds"`
	// 保持镜像的原始设置。该参数与Password或KeyIds.N不能同时指定。只有使用自定义镜像、共享镜像或外部导入镜像创建实例时才能指定该参数为TRUE。取值范围：<br><li>TRUE：表示保持镜像的登录设置<br><li>FALSE：表示不保持镜像的登录设置<br><br>默认取值：FALSE。

	KeepImageLogin *string `json:"KeepImageLogin,omitempty" name:"KeepImageLogin"`
}

type ContainerEnvs struct {

	// 名称

	Name *string `json:"Name,omitempty" name:"Name"`
	// 是否initContainer

	InitContainer *bool `json:"InitContainer,omitempty" name:"InitContainer"`
	// 环境变量

	Envs []*EnvironmentVariable `json:"Envs,omitempty" name:"Envs"`
}

type ForwardRequestResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 请求tke-apiserver http请求返回的body信息

		ResponseBody *string `json:"ResponseBody,omitempty" name:"ResponseBody"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ForwardRequestResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ForwardRequestResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EksCluster struct {

	// 集群Id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// Vpc Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网列表

	SubnetIds []*string `json:"SubnetIds,omitempty" name:"SubnetIds"`
	// k8s 版本号

	K8SVersion *string `json:"K8SVersion,omitempty" name:"K8SVersion"`
	// 集群状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 集群描述信息

	ClusterDesc *string `json:"ClusterDesc,omitempty" name:"ClusterDesc"`
	// 集群创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// Service 子网Id

	ServiceSubnetId *string `json:"ServiceSubnetId,omitempty" name:"ServiceSubnetId"`
}

type CheckClusterHostNameRequest struct {
	*tchttp.BaseRequest

	// 集群ID(如果往已经存在的集群内加节点，则需要传对应的集群ID)

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点主机名称列表

	HostNames []*HostNameValue `json:"HostNames,omitempty" name:"HostNames"`
}

func (r *CheckClusterHostNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterHostNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeVersionsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClustersResourceStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群资源状态集合

		ResourceStatusSet []*ResourceStatus `json:"ResourceStatusSet,omitempty" name:"ResourceStatusSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClustersResourceStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClustersResourceStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHelmChartVersionRequest struct {
	*tchttp.BaseRequest

	// Chart名称

	ChartName *string `json:"ChartName,omitempty" name:"ChartName"`
	// Chart版本

	ChartVersion *string `json:"ChartVersion,omitempty" name:"ChartVersion"`
}

func (r *DeleteHelmChartVersionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHelmChartVersionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAsGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群关联的伸缩组总数

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 集群关联的伸缩组列表

		ClusterAsGroupSet []*ClusterAsGroup `json:"ClusterAsGroupSet,omitempty" name:"ClusterAsGroupSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAsGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAsGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterEndpointStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）

	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
}

func (r *DescribeClusterEndpointStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterEndpointStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneInfo struct {

	// ID

	Id *uint64 `json:"Id,omitempty" name:"Id"`
	// alias

	Alias *string `json:"Alias,omitempty" name:"Alias"`
	// 形如ap-guangzhou

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 形如ap-guangzhou-2

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// status

	Status *string `json:"Status,omitempty" name:"Status"`
	// Remark

	Remark *string `json:"Remark,omitempty" name:"Remark"`
	// time

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// time

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
}

type CheckClusterAdminRoleRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *CheckClusterAdminRoleRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckClusterAdminRoleRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CheckInstancesUpgradeAbleResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群master当前小版本

		ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
		// 集群master对应的大版本目前最新小版本

		LatestVersion *string `json:"LatestVersion,omitempty" name:"LatestVersion"`
		// 可升级节点列表

		UpgradeAbleInstances []*UpgradeAbleInstancesItem `json:"UpgradeAbleInstances,omitempty" name:"UpgradeAbleInstances"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CheckInstancesUpgradeAbleResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CheckInstancesUpgradeAbleResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterInspectionOverview struct {

	// 本次检查的id，用于与其他检测结果区分

	Id *string `json:"Id,omitempty" name:"Id"`
	// 检查的时间

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 如果检查异常终止，则该字段包含错误信息

	Error *string `json:"Error,omitempty" name:"Error"`
	// 统计信息

	Statistic *InspectionStatistic `json:"Statistic,omitempty" name:"Statistic"`
}

type CreateTKEClusterResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群ID

		ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTKEClusterResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateTKEClusterResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnRoutesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 云联网路由数组

		RouteSet []*CcnRoute `json:"RouteSet,omitempty" name:"RouteSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnRoutesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeCcnRoutesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterPodsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// Deployment名称

	ServiceName *string `json:"ServiceName,omitempty" name:"ServiceName"`
	// 命名空间名称，默认为default

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 过滤参数，从PodIP、NodeIP、PodName中过滤含有该参数的Pod实例

	SearchName *string `json:"SearchName,omitempty" name:"SearchName"`
	// 最大输出条数，默认0

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认0

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeClusterPodsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterPodsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeClusterProgressRequest struct {
	*tchttp.BaseRequest

	// 集群实例ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *GetUpgradeClusterProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeClusterProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateLogCollectorRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志采集规则

	Name *string `json:"Name,omitempty" name:"Name"`
	// 日志收集器描述

	Description *string `json:"Description,omitempty" name:"Description"`
	// 输入类型

	InputType *string `json:"InputType,omitempty" name:"InputType"`
	// 输出类型

	OutputType *string `json:"OutputType,omitempty" name:"OutputType"`
	// 输入选项

	InputOption *LogInputOption `json:"InputOption,omitempty" name:"InputOption"`
	// 输出选项

	OutputOption *LogOutputOption `json:"OutputOption,omitempty" name:"OutputOption"`
}

func (r *CreateLogCollectorRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateLogCollectorRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DrainClusterNodeRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 是否只是拉取列表

	DryRun *bool `json:"DryRun,omitempty" name:"DryRun"`
}

func (r *DrainClusterNodeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DrainClusterNodeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterExtraArgs struct {

	// kube-apiserver自定义参数

	KubeAPIServer []*string `json:"KubeAPIServer,omitempty" name:"KubeAPIServer"`
	// kube-controller-manager自定义参数

	KubeControllerManager []*string `json:"KubeControllerManager,omitempty" name:"KubeControllerManager"`
	// kube-scheduler自定义参数

	KubeScheduler []*string `json:"KubeScheduler,omitempty" name:"KubeScheduler"`
}

type DescribeHelmChartDetailResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 文件

		FileEntry *fileEntry `json:"FileEntry,omitempty" name:"FileEntry"`
		// chart

		Chart *Chart `json:"Chart,omitempty" name:"Chart"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeHelmChartDetailResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeHelmChartDetailResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnableClusterAuditRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 日志集ID

	LogsetId *string `json:"LogsetId,omitempty" name:"LogsetId"`
	// 日志主题ID

	TopicId *string `json:"TopicId,omitempty" name:"TopicId"`
	// 区域ID

	RegionId *string `json:"RegionId,omitempty" name:"RegionId"`
}

func (r *EnableClusterAuditRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *EnableClusterAuditRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type LogInputOptNamespace struct {

	// 命名空间

	Namespace *string `json:"Namespace,omitempty" name:"Namespace"`
	// 是否为命名空间内的所有容器

	AllContainers *bool `json:"AllContainers,omitempty" name:"AllContainers"`
	// Services名称数组

	Services []*string `json:"Services,omitempty" name:"Services"`
}

type UpgradeNodeResetParam struct {

	// 实例额外需要设置参数信息

	InstanceAdvancedSettings *InstanceAdvancedSettings `json:"InstanceAdvancedSettings,omitempty" name:"InstanceAdvancedSettings"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。

	EnhancedService *EnhancedService `json:"EnhancedService,omitempty" name:"EnhancedService"`
	// 节点登录信息（目前仅支持使用Password或者单个KeyIds）

	LoginSettings *LoginSettings `json:"LoginSettings,omitempty" name:"LoginSettings"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。（目前仅支持设置单个sgId）

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
}

type DeleteClusterRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群实例删除时的策略：terminate（销毁实例，仅支持按量计费云主机实例） retain （仅移除，保留实例）

	InstanceDeleteMode *string `json:"InstanceDeleteMode,omitempty" name:"InstanceDeleteMode"`
}

func (r *DeleteClusterRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAvailableExtraArgsRequest struct {
	*tchttp.BaseRequest

	// 集群版本

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群类型(MANAGED_CLUSTER或INDEPENDENT_CLUSTER)

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
}

func (r *DescribeClusterAvailableExtraArgsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAvailableExtraArgsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeRegionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域的数量

		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 地域列表

		RegionInstanceSet []*RegionInstance `json:"RegionInstanceSet,omitempty" name:"RegionInstanceSet"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRegionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeRegionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type GetUpgradeInstanceProgressRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 最多获取多少个节点的进度

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 从第几个节点开始获取进度

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *GetUpgradeInstanceProgressRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *GetUpgradeInstanceProgressRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UninstallLogAgentResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UninstallLogAgentResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *UninstallLogAgentResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DataDisk struct {

	// 云盘类型

	DiskType *string `json:"DiskType,omitempty" name:"DiskType"`
	// 文件系统(ext3/ext4/xfs)

	FileSystem *string `json:"FileSystem,omitempty" name:"FileSystem"`
	// 云盘大小(G）

	DiskSize *int64 `json:"DiskSize,omitempty" name:"DiskSize"`
	// 是否自动化格式盘并挂载

	AutoFormatAndMount *bool `json:"AutoFormatAndMount,omitempty" name:"AutoFormatAndMount"`
	// 挂载目录

	MountTarget *string `json:"MountTarget,omitempty" name:"MountTarget"`
}

type CreateClusterEndpointRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群端口所在的子网ID  (仅在开启非外网访问时需要填，必须为集群所在VPC内的子网)

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 是否为外网访问（TRUE 外网访问 FALSE 内网访问，默认值： FALSE）

	IsExtranet *bool `json:"IsExtranet,omitempty" name:"IsExtranet"`
	// 用户自定义域名

	Domain *string `json:"Domain,omitempty" name:"Domain"`
	// 用户指定安全组

	SecurityGroup *string `json:"SecurityGroup,omitempty" name:"SecurityGroup"`
	// 运营商类型、网络计费模式、宽带上限(仅支持公网)

	ExtensiveParameters *string `json:"ExtensiveParameters,omitempty" name:"ExtensiveParameters"`
}

func (r *CreateClusterEndpointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterEndpointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EnvironmentVariable struct {

	// key

	Name *string `json:"Name,omitempty" name:"Name"`
	// val

	Value *string `json:"Value,omitempty" name:"Value"`
}

type ClusterDefinition struct {

	// 集群地域信息

	RegionId *uint64 `json:"RegionId,omitempty" name:"RegionId"`
	// 项目ID（已经废弃）

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 集群Mater的LB所在的SubnetId

	MasterLBSubnetId *string `json:"MasterLBSubnetId,omitempty" name:"MasterLBSubnetId"`
}

type CollectAllCoreResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 产品名称，固定TKE

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// 资源使用量

		Data []*MachineCore `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CollectAllCoreResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CollectAllCoreResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateCLSLogConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateCLSLogConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateCLSLogConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAvailableExtraArgsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群版本

		ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
		// 可用的自定义参数

		AvailableExtraArgs *AvailableExtraArgs `json:"AvailableExtraArgs,omitempty" name:"AvailableExtraArgs"`
		// 集群类型

		ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAvailableExtraArgsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAvailableExtraArgsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResumeClusterInstancesRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
}

func (r *ResumeClusterInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResumeClusterInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Cluster struct {

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 集群名称

	ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
	// 集群描述

	ClusterDescription *string `json:"ClusterDescription,omitempty" name:"ClusterDescription"`
	// 集群版本（默认值为1.10.5）

	ClusterVersion *string `json:"ClusterVersion,omitempty" name:"ClusterVersion"`
	// 集群系统。centos7.2x86_64 或者 ubuntu16.04.1 LTSx86_64，默认取值为ubuntu16.04.1 LTSx86_64

	ClusterOs *string `json:"ClusterOs,omitempty" name:"ClusterOs"`
	// 集群类型，托管集群：MANAGED_CLUSTER，独立集群：INDEPENDENT_CLUSTER。

	ClusterType *string `json:"ClusterType,omitempty" name:"ClusterType"`
	// 集群网络相关参数

	ClusterNetworkSettings *ClusterNetworkSettings `json:"ClusterNetworkSettings,omitempty" name:"ClusterNetworkSettings"`
	// 集群当前node数量

	ClusterNodeNum *uint64 `json:"ClusterNodeNum,omitempty" name:"ClusterNodeNum"`
	// 集群所属的项目ID

	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 标签描述列表。

	TagSpecification []*TagSpecification `json:"TagSpecification,omitempty" name:"TagSpecification"`
	// 集群状态 (Running 运行中  Creating 创建中 Abnormal 异常  )

	ClusterStatus *string `json:"ClusterStatus,omitempty" name:"ClusterStatus"`
	// 集群属性(包括集群不同属性的MAP，属性字段包括NodeNameType (lan-ip模式和hostname 模式，默认无lan-ip模式))

	Property *string `json:"Property,omitempty" name:"Property"`
	// 集群当前master数量

	ClusterMaterNodeNum *uint64 `json:"ClusterMaterNodeNum,omitempty" name:"ClusterMaterNodeNum"`
	// 集群使用镜像id

	ImageId *string `json:"ImageId,omitempty" name:"ImageId"`
	// OsCustomizeType

	OsCustomizeType *string `json:"OsCustomizeType,omitempty" name:"OsCustomizeType"`
	// 集群运行环境docker或container

	ContainerRuntime *string `json:"ContainerRuntime,omitempty" name:"ContainerRuntime"`
	// 创建时间

	CreatedTime *string `json:"CreatedTime,omitempty" name:"CreatedTime"`
	// 创建的集群架构,取值: X86/ARM

	ClusterArch *string `json:"ClusterArch,omitempty" name:"ClusterArch"`
	// pod直通模式

	DirectAccess *string `json:"DirectAccess,omitempty" name:"DirectAccess"`
	// 集群操作系统名称

	ClusterOsAlias *string `json:"ClusterOsAlias,omitempty" name:"ClusterOsAlias"`
	// 是否为双栈集群

	IsDualStack *string `json:"IsDualStack,omitempty" name:"IsDualStack"`
	// Ipv6 Service Cidr

	Ipv6ServiceCidr *string `json:"Ipv6ServiceCidr,omitempty" name:"Ipv6ServiceCidr"`
}

type InspectionStatistic struct {

	// GoodItem

	GoodItem *uint64 `json:"GoodItem,omitempty" name:"GoodItem"`
	// WarnItem

	WarnItem *uint64 `json:"WarnItem,omitempty" name:"WarnItem"`
	// RiskItem

	RiskItem *uint64 `json:"RiskItem,omitempty" name:"RiskItem"`
	// SeriousItem

	SeriousItem *uint64 `json:"SeriousItem,omitempty" name:"SeriousItem"`
	// FailedItem

	FailedItem *uint64 `json:"FailedItem,omitempty" name:"FailedItem"`
}

type DescribeClusterAsGroupOptionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 集群弹性伸缩属性

		ClusterAsGroupOption *ClusterAsGroupOption `json:"ClusterAsGroupOption,omitempty" name:"ClusterAsGroupOption"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClusterAsGroupOptionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAsGroupOptionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterInstanceIdsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 节点角色, MASTER, WORKER, ETCD, MASTER_ETCD,ALL, 默认为WORKER。默认为WORKER类型。

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 偏移量，默认为0。用来控制分页的参数；Limit 为单次返回的最多条目数量，Offset 为偏移量。当相应结果是列表形式时，如果数量超过了 Limit 所限定的值，那么只返回 Limit 个值。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。用来控制分页的参数；Limit 为单次返回的最多条目数量，Offset 为偏移量。当相应结果是列表形式时，如果数量超过了 Limit 所限定的值，那么只返回 Limit 个值。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 根据节点的IP地址进行搜索，同时搜索内网IP和外网IP

	VagueIpAddress *string `json:"VagueIpAddress,omitempty" name:"VagueIpAddress"`
	// 根据节点的名称进行模糊搜索

	VagueInstanceName *string `json:"VagueInstanceName,omitempty" name:"VagueInstanceName"`
	// 根据节点的状态进行筛选

	InstanceStates []*string `json:"InstanceStates,omitempty" name:"InstanceStates"`
	// 根据节点的标签进行搜索

	Labels *string `json:"Labels,omitempty" name:"Labels"`
}

func (r *DescribeClusterInstanceIdsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterInstanceIdsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteAlarmPoliciesRequest struct {
	*tchttp.BaseRequest

	// k8s集群ID

	ClusterInstanceId *string `json:"ClusterInstanceId,omitempty" name:"ClusterInstanceId"`
	// 告警策略ID列表

	AlarmPolicyIds []*string `json:"AlarmPolicyIds,omitempty" name:"AlarmPolicyIds"`
}

func (r *DeleteAlarmPoliciesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteAlarmPoliciesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateClusterVirtualNodeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 虚拟节点数组

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClusterVirtualNodeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateClusterVirtualNodeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterRouteTablesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeClusterRouteTablesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterRouteTablesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeFlowIdStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 任务错误信息

		ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFlowIdStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeFlowIdStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ClusterAsGroupAttribute struct {

	// 伸缩组ID

	AutoScalingGroupId *string `json:"AutoScalingGroupId,omitempty" name:"AutoScalingGroupId"`
	// 是否开启，伸缩组启用停用的时候需要

	AutoScalingGroupEnabled *bool `json:"AutoScalingGroupEnabled,omitempty" name:"AutoScalingGroupEnabled"`
	// 伸缩组最大最小实例数，调整伸缩组配置的时候需要

	AutoScalingGroupRange *AutoScalingGroupRange `json:"AutoScalingGroupRange,omitempty" name:"AutoScalingGroupRange"`
}

type DeleteClusterRouteTableRequest struct {
	*tchttp.BaseRequest

	// 路由表名称

	RouteTableName *string `json:"RouteTableName,omitempty" name:"RouteTableName"`
}

func (r *DeleteClusterRouteTableRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteClusterRouteTableRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteHelmChartResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 是否删除成功

		Deleted *bool `json:"Deleted,omitempty" name:"Deleted"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteHelmChartResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteHelmChartResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClsLogSetsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 日志集数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 日志集集合

		Logsets []*LogSet `json:"Logsets,omitempty" name:"Logsets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeClsLogSetsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClsLogSetsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterAsGroupsRequest struct {
	*tchttp.BaseRequest

	// 集群ID

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 伸缩组ID列表，如果为空，表示拉取集群关联的所有伸缩组。

	AutoScalingGroupIds []*string `json:"AutoScalingGroupIds,omitempty" name:"AutoScalingGroupIds"`
	// 偏移量，默认为0。用来控制分页的参数；Limit 为单次返回的最多条目数量，Offset 为偏移量。当相应结果是列表形式时，如果数量超过了 Limit 所限定的值，那么只返回 Limit 个值。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 返回数量，默认为20，最大值为100。用来控制分页的参数；Limit 为单次返回的最多条目数量，Offset 为偏移量。当相应结果是列表形式时，如果数量超过了 Limit 所限定的值，那么只返回 Limit 个值。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeClusterAsGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterAsGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeClusterStatusRequest struct {
	*tchttp.BaseRequest

	// 集群ID列表，不传默认拉取所有集群

	ClusterIds []*string `json:"ClusterIds,omitempty" name:"ClusterIds"`
}

func (r *DescribeClusterStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeClusterStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyClusterInspectionRequest struct {
	*tchttp.BaseRequest

	// 集群实例id

	ClusterId *string `json:"ClusterId,omitempty" name:"ClusterId"`
	// 自动巡检周期

	Cron *string `json:"Cron,omitempty" name:"Cron"`
}

func (r *ModifyClusterInspectionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyClusterInspectionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SubnetResource struct {

	// cpu

	CPU *float64 `json:"CPU,omitempty" name:"CPU"`
	// mem

	Memory *float64 `json:"Memory,omitempty" name:"Memory"`
	// pod数量

	PodNum *int64 `json:"PodNum,omitempty" name:"PodNum"`
	// 子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
}

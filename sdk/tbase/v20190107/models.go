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

package v20190107

import (
	"encoding/json"

	tchttp "terraform-provider-tencentcloudenterprise/sdk/common/http"
)

// to suppress unused import error, although ugly
var _ = tchttp.POST
var _ = json.Marshal

type ResourcePool struct {

	// 资源池名

	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`
	// 资源池id

	PoolId *int64 `json:"PoolId,omitempty" name:"PoolId"`
	// 资源池描述信息

	PoolDesc *string `json:"PoolDesc,omitempty" name:"PoolDesc"`
}

type ModifyDBInstanceSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBInstanceSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id。

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyDBParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceSyncTypeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id。

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceSyncTypeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceSyncTypeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步流程Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseIsolatedInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseIsolatedInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConstraintRange struct {

	// 区间的最小值。

	Min *string `json:"Min,omitempty" name:"Min"`
	// 区间的最大值。

	Max *string `json:"Max,omitempty" name:"Max"`
}

type PGZoneSellConf struct {

	// 可用区英文名

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区中文名

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 是否为当前地域的默认可用区

	IsDefaultZone *bool `json:"IsDefaultZone,omitempty" name:"IsDefaultZone"`
	// 可选节点规格

	AvailableSpec []*PGSellSpec `json:"AvailableSpec,omitempty" name:"AvailableSpec"`
}

type DescribeSlowLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的慢查询日志数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 慢查询日志的详细信息。

		SlowLogDetails []*SlowLogDetail `json:"SlowLogDetails,omitempty" name:"SlowLogDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSlowLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSlowLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 过滤条件

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
}

func (r *SecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Instance struct {

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 地域ID

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 私有网络子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 私有网络子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 任务类型

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 任务描述

	TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`
	// 虚拟IP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 虚拟端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// 付费类型，postpaid - 后付费，prepaid - 预付费

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
	// 计算节点列表

	CNNodes []*InstanceNode `json:"CNNodes,omitempty" name:"CNNodes"`
	// 数据节点列表

	DNNodes []*InstanceNode `json:"DNNodes,omitempty" name:"DNNodes"`
	// 管理员账号名

	AdminUserName *string `json:"AdminUserName,omitempty" name:"AdminUserName"`
	// 项目ID，为空表示没有所属项目

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名，为空表示没有所属项目

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
}

type DescribeErrorLogRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询错误日志的结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询错误日志的开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeErrorLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeErrorLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePGInstancesRequest struct {
	*tchttp.BaseRequest

	// 按一个或多个实例Id进行查询，每次最多请求100个。指定该参数后，Filters参数失效。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 过滤条件，目前支持InstanceName、Vip、InstanceId、TagKey。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段，支持 InstanceId, InstanceName, CreatedAt。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序方式，支持ASC和DESC，默认为ASC。

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 偏移量，默认为0。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回数量，默认为20，最大值为100。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 实例状态信息，支持 creating, online, isolate。

	Status *string `json:"Status,omitempty" name:"Status"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribePGInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePGInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 冻结流水Id

		BillId *string `json:"BillId,omitempty" name:"BillId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeletePGReadOnlyVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 删除只读VIP异步流程的任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePGReadOnlyVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePGReadOnlyVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 最早备份开始时间。

		MinBackupStartTime *string `json:"MinBackupStartTime,omitempty" name:"MinBackupStartTime"`
		// 最迟备份开始时间。

		MaxBackupStartTime *string `json:"MaxBackupStartTime,omitempty" name:"MaxBackupStartTime"`
		// 备份保存时间。

		BackupReserveDays *int64 `json:"BackupReserveDays,omitempty" name:"BackupReserveDays"`
		// 是否备份xlog。

		IsBackupXlog *int64 `json:"IsBackupXlog,omitempty" name:"IsBackupXlog"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolatePGInstanceHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步流程Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolatePGInstanceHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolatePGInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBackupRulesRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 最小备份开始时间，形如：02:00:00。取值00:00:00~24:00:00。

	MinBackupStartTime *string `json:"MinBackupStartTime,omitempty" name:"MinBackupStartTime"`
	// 最大备份开始时间，形如：02:00:00。取值00:00:00~24:00:00。

	MaxBackupStartTime *string `json:"MaxBackupStartTime,omitempty" name:"MaxBackupStartTime"`
	// 备份频率，每天备份几次，enum：1。

	BackupFrequency *int64 `json:"BackupFrequency,omitempty" name:"BackupFrequency"`
}

func (r *SetBackupRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBackupRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePGAvailableZoneConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可售卖地域配置详情

		Items []*PGRegionSellConf `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePGAvailableZoneConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePGAvailableZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateInstanceHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步流程Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *IsolateInstanceHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamConstraint struct {

	// 用来描述取值范围类型。section: 数值范围, enum: 枚举, string: 字符串。

	Type *string `json:"Type,omitempty" name:"Type"`
	// type取值为section时的数值范围。

	Range []*ConstraintRange `json:"Range,omitempty" name:"Range"`
	// type取值为enum时的枚举值。

	Enum []*string `json:"Enum,omitempty" name:"Enum"`
	// type取值为string时的数值。

	String *string `json:"String,omitempty" name:"String"`
}

type ConfigHistory struct {

	// 参数修改历史记录唯一id。

	Id *int64 `json:"Id,omitempty" name:"Id"`
	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 参数修改的时间。

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 参数修改的结束时间。

	UpdatedAt *string `json:"UpdatedAt,omitempty" name:"UpdatedAt"`
	// 节点类型。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 参数名。

	ParamName *string `json:"ParamName,omitempty" name:"ParamName"`
	// 修改后的参数值。

	ParamNewValue *string `json:"ParamNewValue,omitempty" name:"ParamNewValue"`
	// 修改前的参数值。

	ParamOldValue *string `json:"ParamOldValue,omitempty" name:"ParamOldValue"`
	// 参数修改状态：success, failed, doing

	Status *string `json:"Status,omitempty" name:"Status"`
}

type ActiveIsolatedInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步流程Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActiveIsolatedInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveIsolatedInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableRecoveryTimePointRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 回档时间

	RecoveryTime *string `json:"RecoveryTime,omitempty" name:"RecoveryTime"`
}

func (r *DescribeAvailableRecoveryTimePointRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableRecoveryTimePointRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetAccountPasswordResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceOverviewDetail struct {

	// 本产品实例数量

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 产品名称

	ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
	// 产品类型

	ProductType *string `json:"ProductType,omitempty" name:"ProductType"`
}

type Node struct {

	// 节点角色

	Role *string `json:"Role,omitempty" name:"Role"`
	// 地域信息

	Region *string `json:"Region,omitempty" name:"Region"`
	// 可用区信息

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 节点所在节点组中的节点类型

	Type *string `json:"Type,omitempty" name:"Type"`
	// 节点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
}

type CreateBackupRequest struct {
	*tchttp.BaseRequest

	// 实例id，当前要备份的实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreateBackupRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tdpg-0wnkx2ex

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *TerminateInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableRecoveryTimePointResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可恢复时间点

		RecoveryTimePoint []*TimePoint `json:"RecoveryTimePoint,omitempty" name:"RecoveryTimePoint"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableRecoveryTimePointResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableRecoveryTimePointResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PGRegionSellConf struct {

	// 地域英文名

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域中文名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 所属大区

	Area *string `json:"Area,omitempty" name:"Area"`
	// 是否为默认地域

	IsDefaultRegion *bool `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`
	// 可用区售卖配置

	AvailableZones []*PGZoneSellConf `json:"AvailableZones,omitempty" name:"AvailableZones"`
}

type DescribeBackupDetailsRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 唯一标识本次备份。

	BackupUid *int64 `json:"BackupUid,omitempty" name:"BackupUid"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeBackupDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBConfigHistoryRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBConfigHistoryRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBConfigHistoryRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RenewInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TimePoint struct {

	// 实例回档时可选择的时间

	Time *string `json:"Time,omitempty" name:"Time"`
}

type DescribePGInstanceDetailsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribePGInstanceDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePGInstanceDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type EngineVersion struct {

	// 引擎类型，例如 TBaseV2,TBaseV3

	Type *string `json:"Type,omitempty" name:"Type"`
	// 引擎版本，例如：1.0.1

	Version *string `json:"Version,omitempty" name:"Version"`
}

type SellSpec struct {

	// 规格码，唯一标识一个规格，创建实例时传该字段用于标识一个规格

	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`
	// 节点类型: cn（计算节点）, dn（存储节点）

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// cpu核心数，单位：核

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小，单位：GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 磁盘最小规格，单位：GB

	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`
	// 磁盘最大规格，单位：GB

	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 磁盘调整步长，单位：GB

	StorageStep *int64 `json:"StorageStep,omitempty" name:"StorageStep"`
	// 磁盘调整步长，单位：GB

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
}

type DescribePGInstanceDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例Id

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例名

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 实例状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 实例状态描述

		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
		// 任务类型

		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
		// 任务描述

		TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`
		// 付费类型, postpaid-后付费, prepaid-预付费

		TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
		// 地域信息

		Region *string `json:"Region,omitempty" name:"Region"`
		// 地域名称

		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
		// 读写vpc信息

		MainVpc *VpcInfo `json:"MainVpc,omitempty" name:"MainVpc"`
		// 只读vpc信息

		RoVpc *VpcInfo `json:"RoVpc,omitempty" name:"RoVpc"`
		// 字符集

		Charset *string `json:"Charset,omitempty" name:"Charset"`
		// 引擎版本

		EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
		// cpu核数

		Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
		// 内存大小

		Memory *int64 `json:"Memory,omitempty" name:"Memory"`
		// 磁盘大小

		Storage *int64 `json:"Storage,omitempty" name:"Storage"`
		// 节点信息

		Nodes []*Node `json:"Nodes,omitempty" name:"Nodes"`
		// 实例创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 项目ID，为空表示没有所属项目

		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
		// 项目名，为空表示没有所属项目

		ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePGInstanceDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePGInstanceDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedPGInstanceRequest struct {
	*tchttp.BaseRequest

	// Postgres实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ReleaseIsolatedPGInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseIsolatedPGInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSaleInfo struct {

	// 可售卖地域总数

	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 可售卖地域详情

	RegionSaleDetail []*RegionSaleDetail `json:"RegionSaleDetail,omitempty" name:"RegionSaleDetail"`
	// 可售卖地域引擎类型TbaseV2/TbaseV3/TbaseV5

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
}

type DescribeAvailableEngineVersionsRequest struct {
	*tchttp.BaseRequest

	// 表示产品类型 (tbase,tdapg,tpostgres)

	Product *string `json:"Product,omitempty" name:"Product"`
}

func (r *DescribeAvailableEngineVersionsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableEngineVersionsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupListsRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeBackupListsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupListsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 按一个或多个实例Id进行查询，每次最多请求100个。指定该参数后，Filters参数失效。

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
	// 过滤条件，目前支持InstanceName、Vip、InstanceId、TagKey。

	Filters []*Filter `json:"Filters,omitempty" name:"Filters"`
	// 排序字段，支持 InstanceId,InstanceName,CreateTime,PeriodEndTime。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序方式，支持ASC和DESC，默认为ASC。

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 偏移量，默认为0。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 返回数量，默认为20，最大值为100。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
	// 实例状态信息，支持 creating, online, isolate

	Status *string `json:"Status,omitempty" name:"Status"`
	// 根据计费类型过滤，支持postpaid后付费，prepaid预付费两种。

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceOverviewRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeResourceOverviewRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceOverviewRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ActiveIsolatedPGInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ActiveIsolatedPGInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveIsolatedPGInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBParametersResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParametersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePGReadOnlyVipRequest struct {
	*tchttp.BaseRequest

	// 唯一标识的vpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 唯一标识的vpc子网Id

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// 选填，可指定vip用于只读

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例唯一标识Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *CreatePGReadOnlyVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePGReadOnlyVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNameRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 新的实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
}

func (r *ModifyInstanceNameRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNameRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Backup struct {

	// 实例备份信息的唯一标识，查询备份详情时的入参。

	BackupUid *int64 `json:"BackupUid,omitempty" name:"BackupUid"`
	// 实例备份开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 实例备份结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 备份策略类型  Auto(自动备份)

	StrategyType *string `json:"StrategyType,omitempty" name:"StrategyType"`
	// 备份类型 Physical(物理备份)

	BackupType *string `json:"BackupType,omitempty" name:"BackupType"`
	// 备份状态。1(备份中)，2（备份成功），-1（备份失败）

	Status *int64 `json:"Status,omitempty" name:"Status"`
}

type ActivateIsolatedInstanceHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActivateIsolatedInstanceHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateIsolatedInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableZoneConfigResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 配置总条数

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可售卖地域配置详情

		Items []*RegionSellConf `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableZoneConfigResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZoneConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RegionSellConf struct {

	// 地域英文名

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域中文名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 所属大区

	Area *string `json:"Area,omitempty" name:"Area"`
	// 是否为默认地域

	IsDefaultRegion *bool `json:"IsDefaultRegion,omitempty" name:"IsDefaultRegion"`
	// 可用区售卖配置

	AvailableZones []*ZoneSellConf `json:"AvailableZones,omitempty" name:"AvailableZones"`
}

type RecoveryNodeSpec struct {

	// 节点类型：0-GTM，1-CN，2-DN，3-FN

	NodeType *uint64 `json:"NodeType,omitempty" name:"NodeType"`
	// 节点组数

	NodeSet *uint64 `json:"NodeSet,omitempty" name:"NodeSet"`
	// 节点个数

	NodeCount *uint64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// CPU，单位：核

	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小，单位：GB

	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`
	// 磁盘大小，单位：GB

	Storage *uint64 `json:"Storage,omitempty" name:"Storage"`
}

type DescribeBackupListsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 查询结果总数。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 备份信息数组。

		BackupList []*Backup `json:"BackupList,omitempty" name:"BackupList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupListsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupListsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Filter struct {

	// 需要过滤的字段。

	Name *string `json:"Name,omitempty" name:"Name"`
	// 字段的过滤值。

	Values []*string `json:"Values,omitempty" name:"Values"`
}

type DescribeAvailableEngineVersionsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 引擎版本数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 可用的引擎版本列表

		EngineVersions []*EngineVersion `json:"EngineVersions,omitempty" name:"EngineVersions"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableEngineVersionsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableEngineVersionsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PGSellSpec struct {

	// 规格码，唯一标识一个规格，创建实例时传该字段用于标识一个规格

	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`
	// 节点类型: pg_dn（postgres实例存储节点）

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// cpu核心数，单位：核

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小，单位：GB

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 磁盘最小规格，单位：GB

	MinStorage *int64 `json:"MinStorage,omitempty" name:"MinStorage"`
	// 磁盘最大规格，单位：GB

	MaxStorage *int64 `json:"MaxStorage,omitempty" name:"MaxStorage"`
	// 磁盘调整步长，单位：GB

	StorageStep *int64 `json:"StorageStep,omitempty" name:"StorageStep"`
	// 磁盘调整步长，单位：GB

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
}

type NodeConfigParams struct {

	// 节点类型，目前支持：cn、dn。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 需要修改的参数配置。

	ConfigParams []*ConfigParam `json:"ConfigParams,omitempty" name:"ConfigParams"`
}

type DescribeAvailableZonesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 资源池下的机器所在可用区

		Zones []*string `json:"Zones,omitempty" name:"Zones"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableZonesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZonesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDbNameListRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// db名称，可进行模糊匹配。

	DbName *string `json:"DbName,omitempty" name:"DbName"`
}

func (r *DescribeDbNameListRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbNameListRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeBackupDetail struct {

	// 节点名。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点类型。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 备份状态。1(备份中)，2(备份成功)，-1(备份失败)

	LatestStatus *int64 `json:"LatestStatus,omitempty" name:"LatestStatus"`
	// 最近备份开始时间。

	LatestStartTime *string `json:"LatestStartTime,omitempty" name:"LatestStartTime"`
	// 最近备份结束时间。

	LatestEndTime *string `json:"LatestEndTime,omitempty" name:"LatestEndTime"`
}

type DescribeSlowLogRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询慢查询日志的开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 查询慢查询日志的结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 数据库名。

	Database *string `json:"Database,omitempty" name:"Database"`
	// 排序字段，支持Calls,CostTime,TotalCallsPercent,TotalCostTimePercent。

	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`
	// 排序方式，支持ASC和DESC。

	OrderByType *string `json:"OrderByType,omitempty" name:"OrderByType"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeSlowLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeSlowLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVipRequest struct {
	*tchttp.BaseRequest

	// 实例id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 要指定的ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 要指定的port

	Port *uint64 `json:"Port,omitempty" name:"Port"`
}

func (r *ModifyVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResetAccountPasswordRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 数据库账号名

	Username *string `json:"Username,omitempty" name:"Username"`
	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为8~32位。

	NewPassword *string `json:"NewPassword,omitempty" name:"NewPassword"`
}

func (r *ResetAccountPasswordRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ResetAccountPasswordRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Outbound struct {

	// 策略，ACCEPT或者DROP。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 地址组id代表的地址集合。

	AddressMod *string `json:"AddressMod,omitempty" name:"AddressMod"`
	// 来源Ip或Ip段，例如192.168.0.0/16。

	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`
	// 描述。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 安全组id代表的地址集合。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 网络协议，支持udp、tcp等。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 端口。

	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`
	// 服务组id代表的协议和端口集合。

	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`
}

type ActivateIsolatedInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ActivateIsolatedInstanceHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActivateIsolatedInstanceHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePGReadOnlyVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步生成只读vip的任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePGReadOnlyVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePGReadOnlyVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableRecoveryTimeRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeAvailableRecoveryTimeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableRecoveryTimeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableRegionResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 可售卖地域详情

		Items []*RegionSaleInfo `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableRegionResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableRegionResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 资源ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 区域ID

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// CN或者DN 节点规格代码

	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`
	// 节点类型，0-GTM,1-CN,2-DN,3-FN

	NodeType *int64 `json:"NodeType,omitempty" name:"NodeType"`
	// 变配后的新配置总节点组数

	SetCount *int64 `json:"SetCount,omitempty" name:"SetCount"`
	// 变配后单个节点磁盘总大小

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 操作类型，1增加节点，2减少节点（暂不支持），3提升节点配置

	ModifyType *int64 `json:"ModifyType,omitempty" name:"ModifyType"`
	// 引擎类型，当前只支持TbaseV2

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 变配后的新配置节点总个数，降配时使用，暂不支持

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 降配时必传ModifyMode=down，降配时使用，暂不支持

	ModifyMode *string `json:"ModifyMode,omitempty" name:"ModifyMode"`
	// 空间回收开始时间

	SpaceRecycleStartTime *string `json:"SpaceRecycleStartTime,omitempty" name:"SpaceRecycleStartTime"`
	// 空间回收结束时间

	SpaceRecycleEndTime *string `json:"SpaceRecycleEndTime,omitempty" name:"SpaceRecycleEndTime"`
}

func (r *ModifyInstanceHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TerminateInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务Id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *TerminateInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *TerminateInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ErrorLogDetail struct {

	// 用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 数据库名。

	Database *string `json:"Database,omitempty" name:"Database"`
	// 错误日志发生时间。

	ErrorTime *string `json:"ErrorTime,omitempty" name:"ErrorTime"`
	// 错误详细信息。

	ErrorMessage *string `json:"ErrorMessage,omitempty" name:"ErrorMessage"`
}

type InstanceOverview struct {

	// 所有类型实例总数量

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
	// 实例资源详情

	Detail []*InstanceOverviewDetail `json:"Detail,omitempty" name:"Detail"`
}

type RecoveryInfo struct {

	// 回档任务ID，会在获取回档时间接口中返回

	RecoveryTimePoint *string `json:"RecoveryTimePoint,omitempty" name:"RecoveryTimePoint"`
	// 回档时间，会在获取回档时间接口中返回，需选择回档时间范围中的某一个时刻点

	RecoveryTime *string `json:"RecoveryTime,omitempty" name:"RecoveryTime"`
	// 要回档的实例ID

	RecoveryInstanceId *string `json:"RecoveryInstanceId,omitempty" name:"RecoveryInstanceId"`
}

type DescribeAvailableZoneConfigRequest struct {
	*tchttp.BaseRequest

	// 引擎类型，支持TbaseV2，TbaseV3

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
}

func (r *DescribeAvailableZoneConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZoneConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePoolsRequest struct {
	*tchttp.BaseRequest

	// 资源池名称，可用于模糊查询

	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`
	// 资源池Id，可用于精准查询

	PoolId *int64 `json:"PoolId,omitempty" name:"PoolId"`
}

func (r *DescribeResourcePoolsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcePoolsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ConfigParam struct {

	// 参数名。

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 参数值。

	ParameterValue *string `json:"ParameterValue,omitempty" name:"ParameterValue"`
}

type InstanceNode struct {

	// 节点类型

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点cpu

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 节点内存

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 节点数据盘大小

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 节点唯一ID

	NodeId *string `json:"NodeId,omitempty" name:"NodeId"`
	// 已用数据存储空间

	UsedStorage *int64 `json:"UsedStorage,omitempty" name:"UsedStorage"`
	// 节点名

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 所在节点组中的节点类型，master-主节点

	Type *string `json:"Type,omitempty" name:"Type"`
	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
}

type RestartInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *RestartInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ZoneSellConf struct {

	// 可用区英文名

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区中文名

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 是否为当前地域的默认可用区

	IsDefaultZone *bool `json:"IsDefaultZone,omitempty" name:"IsDefaultZone"`
	// 可选节点规格

	AvailableSpec []*SellSpec `json:"AvailableSpec,omitempty" name:"AvailableSpec"`
}

type CreatePGInstanceHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 冻结流水Id

		BillId *string `json:"BillId,omitempty" name:"BillId"`
		// 实例Id列表

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 流程Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 订单号

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePGInstanceHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePGInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedPGInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步流程Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ReleaseIsolatedPGInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseIsolatedPGInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RestartInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 如果指定NodeIds，则重启指定节点。如果为空或者包含所有节点id，则重启整个实例。不能和NodeTypes同时使用。

	NodeIds []*string `json:"NodeIds,omitempty" name:"NodeIds"`
	// 如果指定NodeTypes，则重启指定类型的节点。如果为空或者包含实例所有节点类型，则重启整个实例。不能和NodeIds同时使用。

	NodeTypes []*string `json:"NodeTypes,omitempty" name:"NodeTypes"`
}

func (r *RestartInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RestartInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableRecoveryNodeSpecRequest struct {
	*tchttp.BaseRequest

	// 实例ID

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 可恢复时间点，通过该接口DescribeAvailableRecoveryTimePoint获取

	RecoveryTimePoint *string `json:"RecoveryTimePoint,omitempty" name:"RecoveryTimePoint"`
	// 可恢复时间，通过该接口DescribeAvailableRecoveryTime获取

	RecoveryTime *string `json:"RecoveryTime,omitempty" name:"RecoveryTime"`
}

func (r *DescribeAvailableRecoveryNodeSpecRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableRecoveryNodeSpecRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTaskStatusResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务的执行状态，有“success”，“running”，“failure”三种状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTaskStatusResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTaskStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePGErrorLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的错误日志数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 错误日志详细信息。

		ErrorLogDetails []*ErrorLogDetail `json:"ErrorLogDetails,omitempty" name:"ErrorLogDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePGErrorLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePGErrorLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamItem struct {

	// 节点类型。

	NodeType *string `json:"NodeType,omitempty" name:"NodeType"`
	// 节点名。

	NodeName *string `json:"NodeName,omitempty" name:"NodeName"`
	// 节点的各个参数信息。

	Details []*ParamDetail `json:"Details,omitempty" name:"Details"`
	// 参数个数。

	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
}

type CreateInstanceHourResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 冻结流水Id。

		BillId *string `json:"BillId,omitempty" name:"BillId"`
		// 实例Id列表。

		InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
		// 流程Id。

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 订单号

		DealName *string `json:"DealName,omitempty" name:"DealName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateInstanceHourResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceHourResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBConfigHistoryResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的参数操作记录数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 配置变更的操作记录。

		ConfigHistory []*ConfigHistory `json:"ConfigHistory,omitempty" name:"ConfigHistory"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBConfigHistoryResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBConfigHistoryResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTaskStatusRequest struct {
	*tchttp.BaseRequest

	// 流程任务Id

	TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeInstanceTaskStatusRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceTaskStatusRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolateInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolateInstanceHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolateInstanceHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 需修改的节点参数配置。

	NodeConfigParams []*NodeConfigParams `json:"NodeConfigParams,omitempty" name:"NodeConfigParams"`
}

func (r *ModifyDBParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeBackupRulesRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeBackupRulesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupRulesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePGInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// Postgres实例列表。

		Items []*PGInstanceSet `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePGInstancesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePGInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyVipResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务ID

		FlowId *uint64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyVipResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyVipResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyDBInstanceSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：tbase 等。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 要修改的安全组ID列表，一个或者多个安全组Id组成的数组。

	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" name:"SecurityGroupIds"`
	// 实例ID，格式如：tdpg-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ModifyDBInstanceSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyDBInstanceSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SyncTypeInfo struct {

	// 节点的原复制类型，可选项(sync:同步不可退化，async:异步，sync2async:同步可退化)

	SourceSyncType *string `json:"SourceSyncType,omitempty" name:"SourceSyncType"`
	// 备节点的节点id。

	SlaveNodeId *string `json:"SlaveNodeId,omitempty" name:"SlaveNodeId"`
	// 节点的目标复制类型，可选项(sync:同步不可退化，async:异步，sync2async:同步可退化)

	TargetSyncType *string `json:"TargetSyncType,omitempty" name:"TargetSyncType"`
	// 主节点的节点id。

	MasterNodeId *string `json:"MasterNodeId,omitempty" name:"MasterNodeId"`
}

type ActiveIsolatedInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ActiveIsolatedInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveIsolatedInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBSecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 安全组规则

		Groups []*SecurityGroup `json:"Groups,omitempty" name:"Groups"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDBSecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDBParametersRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例节点类型: cn, dn

	NodeTypes []*string `json:"NodeTypes,omitempty" name:"NodeTypes"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeDBParametersRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBParametersRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SecurityGroupsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SecurityGroupsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SecurityGroupsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SlowLogDetail struct {

	// 花费总时间。

	TotalTime *float64 `json:"TotalTime,omitempty" name:"TotalTime"`
	// 调用总次数。

	TotalCalls *int64 `json:"TotalCalls,omitempty" name:"TotalCalls"`
	// 慢查询sql详细信息。

	NormalQuerys []*NormQueryItem `json:"NormalQuerys,omitempty" name:"NormalQuerys"`
}

type ActiveIsolatedPGInstanceResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步流程Id

		FlowId *int64 `json:"FlowId,omitempty" name:"FlowId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ActiveIsolatedPGInstanceResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ActiveIsolatedPGInstanceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的实例数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 实例列表。

		Items []*Instance `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
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

type DescribeErrorLogResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的错误日志数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 错误日志详细信息。

		ErrorLogDetails []*ErrorLogDetail `json:"ErrorLogDetails,omitempty" name:"ErrorLogDetails"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeErrorLogResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeErrorLogResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ResourceUsage struct {

	// 总量

	Total *int64 `json:"Total,omitempty" name:"Total"`
	// 已使用量

	Used *int64 `json:"Used,omitempty" name:"Used"`
}

type DescribeDBSecurityGroupsRequest struct {
	*tchttp.BaseRequest

	// 数据库引擎名称：tbase等。

	Product *string `json:"Product,omitempty" name:"Product"`
	// 实例ID，格式如：tbase-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeDBSecurityGroupsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDBSecurityGroupsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ProductLicenseData struct {

	// 磐石分配的产品 ID，固定为 87870

	ProductId *string `json:"ProductId,omitempty" name:"ProductId"`
	// 产品描述

	ProductIdDescribe *string `json:"ProductIdDescribe,omitempty" name:"ProductIdDescribe"`
	// 数量单位或节点类型

	UsageUnit *string `json:"UsageUnit,omitempty" name:"UsageUnit"`
	// 节点数量

	UsageValue *string `json:"UsageValue,omitempty" name:"UsageValue"`
}

type DescribeAvailableCharsetRequest struct {
	*tchttp.BaseRequest

	// 引擎类型，支持TbaseV2，TbaseV3和TbaseV5,TbaseV5_C

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
}

func (r *DescribeAvailableCharsetRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableCharsetRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type RenewInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如dpag-751k9fi1

	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *RenewInstancesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *RenewInstancesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailsRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DescribeInstanceDetailsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDetailsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type TagInfo struct {

	// TagKey值

	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`
	// TagValue值

	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type DescribeAccountsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的账号数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 账号信息。

		Accounts []*AccountInfo `json:"Accounts,omitempty" name:"Accounts"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAccountsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableRecoveryTimeResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回档允许的最早时间

		StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
		// 回档允许的最晚时间

		EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
		// 本次回档的任务id

		RecoveryTaskId *int64 `json:"RecoveryTaskId,omitempty" name:"RecoveryTaskId"`
		// 当前该实例的回档状态，success成功，init初始化中，running回档中，running状态下不允许再次发起回档请求

		RecoveryStatus *string `json:"RecoveryStatus,omitempty" name:"RecoveryStatus"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableRecoveryTimeResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableRecoveryTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReleaseIsolatedInstanceRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如tbase-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *ReleaseIsolatedInstanceRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReleaseIsolatedInstanceRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PGInstanceSet struct {

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 实例名

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 地域信息

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域名称

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 可用区信息

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 可用区名称

	ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
	// 私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 私有网络子网ID

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 私有网络子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// 创建时间

	CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
	// 实例状态

	Status *string `json:"Status,omitempty" name:"Status"`
	// 实例状态描述

	StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
	// 任务类型

	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
	// 任务描述信息

	TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`
	// 实例ip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 实例端口

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
	// cpu核数

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// 磁盘大小

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 节点数量

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 付费类型，(后付费-postpaid，预付费-prepaid)

	TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
	// 管理员账号名

	AdminUserName *string `json:"AdminUserName,omitempty" name:"AdminUserName"`
	// 只读vip

	RoVip *string `json:"RoVip,omitempty" name:"RoVip"`
	// 只读port

	RoVport *int64 `json:"RoVport,omitempty" name:"RoVport"`
	// 项目ID，为空表示没有所属项目

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 项目名，为空表示没有所属项目

	ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
	// 实例隔离者

	IsolateOperator *string `json:"IsolateOperator,omitempty" name:"IsolateOperator"`
	// 是否欠费，1 - 欠费， 0 - 不欠费

	IsArrears *int64 `json:"IsArrears,omitempty" name:"IsArrears"`
}

type DeletePGReadOnlyVipRequest struct {
	*tchttp.BaseRequest

	// 唯一标识的vpcId

	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`
	// 唯一标识的vpc子网Id

	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`
	// 要删除的VIP

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// 唯一标识的实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *DeletePGReadOnlyVipRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeletePGReadOnlyVipRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableRecoveryNodeSpecResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 回档实例的规格

		NodeSpec *RecoveryNodeSpec `json:"NodeSpec,omitempty" name:"NodeSpec"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableRecoveryNodeSpecResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableRecoveryNodeSpecResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type AccountInfo struct {

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 账号名。

	Username *string `json:"Username,omitempty" name:"Username"`
}

type RegionSaleDetail struct {

	// 地域英文名

	Region *string `json:"Region,omitempty" name:"Region"`
	// 地域中文名

	RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
	// 地域所属大区

	Area *string `json:"Area,omitempty" name:"Area"`
	// 是否售罄

	IsSellOut *bool `json:"IsSellOut,omitempty" name:"IsSellOut"`
}

type DescribeBackupDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的记录数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 节点详情。

		BackupDetailList []*NodeBackupDetail `json:"BackupDetailList,omitempty" name:"BackupDetailList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeBackupDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeBackupDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type IsolatePGInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 实例ID，形如postgres-0wnkx2ex

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *IsolatePGInstanceHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *IsolatePGInstanceHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NodeInfo struct {

	// 要创建实例的某类型节点的总个数（节点组数与每组节点个数的乘积）。

	NodeCount *int64 `json:"NodeCount,omitempty" name:"NodeCount"`
	// 要创建实例的某类型节点的组数。

	SetCount *int64 `json:"SetCount,omitempty" name:"SetCount"`
	// 要创建实例的某类型节点的单节点磁盘大小。

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
	// 要创建实例的某类型节点的规格码，可以通过DescribeAvailableZoneConfig接口获得。

	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`
	// 设置节点的复制类型，目前只有dn设置，可选参数为sync（同步），async（异步），sync2async（异步可退化），默认值为sync2async

	SyncType *string `json:"SyncType,omitempty" name:"SyncType"`
}

type PgNodeInfo struct {

	// 节点规格代码

	SpecCode *string `json:"SpecCode,omitempty" name:"SpecCode"`
	// 磁盘单个节点大小

	Storage *int64 `json:"Storage,omitempty" name:"Storage"`
}

type SecurityGroup struct {

	// 创建时间，时间格式：yyyy-mm-dd hh:mm:ss

	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
	// 入站规则。

	Inbound []*Inbound `json:"Inbound,omitempty" name:"Inbound"`
	// 出站规则。

	Outbound []*Outbound `json:"Outbound,omitempty" name:"Outbound"`
	// 项目ID。

	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`
	// 安全组ID。

	SecurityGroupId *string `json:"SecurityGroupId,omitempty" name:"SecurityGroupId"`
	// 安全组名称。

	SecurityGroupName *string `json:"SecurityGroupName,omitempty" name:"SecurityGroupName"`
	// 安全组备注。

	SecurityGroupRemark *string `json:"SecurityGroupRemark,omitempty" name:"SecurityGroupRemark"`
}

type DescribeAvailableRegionRequest struct {
	*tchttp.BaseRequest

	// 引擎类型，TbaseV2/TbaseV3/TbaseV5

	EngineType []*string `json:"EngineType,omitempty" name:"EngineType"`
}

func (r *DescribeAvailableRegionRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableRegionRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourcePoolsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 总pool数量

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 资源池详细信息

		Items []*ResourcePool `json:"Items,omitempty" name:"Items"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourcePoolsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourcePoolsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceNameResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyInstanceNameResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceNameResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 可用区

	Zone *string `json:"Zone,omitempty" name:"Zone"`
	// 私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 引擎版本

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 管理员账号的密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 字符集，根据DescribeAvailableCharset查询

	Charset *string `json:"Charset,omitempty" name:"Charset"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 网络类型

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 创建实例cn节点相关的规格信息。

	CnNode *NodeInfo `json:"CnNode,omitempty" name:"CnNode"`
	// 创建实例dn节点相关的规格信息。

	DnNode *NodeInfo `json:"DnNode,omitempty" name:"DnNode"`
	// 购买数量，(0,10]，不传默认创建1个

	GoodsNum *int64 `json:"GoodsNum,omitempty" name:"GoodsNum"`
	// 要绑定的安全组

	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`
	// 回档实例必传参数

	Recovery *RecoveryInfo `json:"Recovery,omitempty" name:"Recovery"`
	// 标签

	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`
	// 引擎类型，当前只支持TbaseV2,TbaseV3

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
	// 备可用区列表

	SlaveZones []*string `json:"SlaveZones,omitempty" name:"SlaveZones"`
}

func (r *CreateInstanceHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateInstanceHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceDetailsResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 实例ID

		InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
		// 实例名称

		InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
		// 管理员账号名

		AdminUserName *string `json:"AdminUserName,omitempty" name:"AdminUserName"`
		// 状态

		Status *string `json:"Status,omitempty" name:"Status"`
		// 状态描述

		StatusDesc *string `json:"StatusDesc,omitempty" name:"StatusDesc"`
		// 地域ID

		Region *string `json:"Region,omitempty" name:"Region"`
		// 地域名称

		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
		// 可用区ID

		Zone *string `json:"Zone,omitempty" name:"Zone"`
		// 可用区名称

		ZoneName *string `json:"ZoneName,omitempty" name:"ZoneName"`
		// 私有网络Id

		VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
		// 私有网络名称

		VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
		// 私有网络子网ID

		SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
		// 私有网络子网名称

		SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
		// 创建时间

		CreatedAt *string `json:"CreatedAt,omitempty" name:"CreatedAt"`
		// 任务类型

		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`
		// 任务描述

		TaskDesc *string `json:"TaskDesc,omitempty" name:"TaskDesc"`
		// 虚拟IP

		Vip *string `json:"Vip,omitempty" name:"Vip"`
		// 虚拟端口

		Vport *int64 `json:"Vport,omitempty" name:"Vport"`
		// 付费类型，postpaid - 后付费，prepaid - 预付费

		TradeType *string `json:"TradeType,omitempty" name:"TradeType"`
		// GTM节点列表

		GTMNodes []*InstanceNode `json:"GTMNodes,omitempty" name:"GTMNodes"`
		// 计算节点列表

		CNNodes []*InstanceNode `json:"CNNodes,omitempty" name:"CNNodes"`
		// 数据节点列表

		DNNodes []*InstanceNode `json:"DNNodes,omitempty" name:"DNNodes"`
		// 备份及日志磁盘空间

		BackupStorage *int64 `json:"BackupStorage,omitempty" name:"BackupStorage"`
		// 实例字符集

		Charset *string `json:"Charset,omitempty" name:"Charset"`
		// 引擎类型

		EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
		// 引擎版本

		EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
		// 集群名称

		ClusterName *string `json:"ClusterName,omitempty" name:"ClusterName"`
		// 项目ID，为空表示没有所属项目

		ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
		// 项目名，为空表示没有所属项目

		ProjectName *string `json:"ProjectName,omitempty" name:"ProjectName"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceDetailsResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeInstanceDetailsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribePGErrorLogRequest struct {
	*tchttp.BaseRequest

	// 实例Id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 查询错误日志的结束时间。

	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 查询错误日志的开始时间。

	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribePGErrorLogRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePGErrorLogRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ParamDetail struct {

	// 参数名。

	ParameterName *string `json:"ParameterName,omitempty" name:"ParameterName"`
	// 参数默认值。

	DefaultValue *string `json:"DefaultValue,omitempty" name:"DefaultValue"`
	// 修改此参数是否需要重启。

	NeedRestart *bool `json:"NeedRestart,omitempty" name:"NeedRestart"`
	// 当前运行值。

	RunningValue *string `json:"RunningValue,omitempty" name:"RunningValue"`
	// 参数取值范围。

	ValueRange []*ParamConstraint `json:"ValueRange,omitempty" name:"ValueRange"`
	// 参数单位。

	Unit *string `json:"Unit,omitempty" name:"Unit"`
	// 参数英文简介。

	ShortDesc *string `json:"ShortDesc,omitempty" name:"ShortDesc"`
}

type DescribeAvailableZonesRequest struct {
	*tchttp.BaseRequest

	// 资源池名称

	PoolName *string `json:"PoolName,omitempty" name:"PoolName"`
	// 资源池Id

	PoolId *int64 `json:"PoolId,omitempty" name:"PoolId"`
}

func (r *DescribeAvailableZonesRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableZonesRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type NormQueryItem struct {

	// 数据库名。

	DatabaseName *string `json:"DatabaseName,omitempty" name:"DatabaseName"`
	// 脱敏后的sql语句。

	NormalQuery *string `json:"NormalQuery,omitempty" name:"NormalQuery"`
	// 执行时间最长的语句。

	MaxElapsedQuery *string `json:"MaxElapsedQuery,omitempty" name:"MaxElapsedQuery"`
	// 调用次数。

	Calls *int64 `json:"Calls,omitempty" name:"Calls"`
	// 花费总时间。

	CostTime *float64 `json:"CostTime,omitempty" name:"CostTime"`
	// 客户端ip。

	ClientIp *string `json:"ClientIp,omitempty" name:"ClientIp"`
	// 用户名。

	UserName *string `json:"UserName,omitempty" name:"UserName"`
	// 在总的慢查询中的访问次数占比。

	TotalCallsPercent *float64 `json:"TotalCallsPercent,omitempty" name:"TotalCallsPercent"`
	// 在总的慢查询中的访问时间占比。

	TotalCostTimePercent *float64 `json:"TotalCostTimePercent,omitempty" name:"TotalCostTimePercent"`
	// 花费的最小时间。

	MinCostTime *float64 `json:"MinCostTime,omitempty" name:"MinCostTime"`
	// 花费的最长时间。

	MaxCostTime *float64 `json:"MaxCostTime,omitempty" name:"MaxCostTime"`
	// 最早一条慢查询的sql时间。

	FirstTime *string `json:"FirstTime,omitempty" name:"FirstTime"`
	// 最近一条慢查询的sql时间。

	LastTime *string `json:"LastTime,omitempty" name:"LastTime"`
	// 读共享内存块数。

	SharedReadBlks *int64 `json:"SharedReadBlks,omitempty" name:"SharedReadBlks"`
	// 写共享内存块数。

	SharedWriteBlks *int64 `json:"SharedWriteBlks,omitempty" name:"SharedWriteBlks"`
	// 读io总耗时。

	ReadCostTime *float64 `json:"ReadCostTime,omitempty" name:"ReadCostTime"`
	// 写io总耗时。

	WriteCostTime *float64 `json:"WriteCostTime,omitempty" name:"WriteCostTime"`
}

type DescribePGAvailableZoneConfigRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribePGAvailableZoneConfigRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribePGAvailableZoneConfigRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreatePGInstanceHourRequest struct {
	*tchttp.BaseRequest

	// 私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 子网Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 网络类型

	NetType *string `json:"NetType,omitempty" name:"NetType"`
	// 资源池id

	PoolId *int64 `json:"PoolId,omitempty" name:"PoolId"`
	// 引擎类型，当前只支持PostgreSQL

	EngineType *string `json:"EngineType,omitempty" name:"EngineType"`
	// 引擎版本

	EngineVersion *string `json:"EngineVersion,omitempty" name:"EngineVersion"`
	// 内存大小

	Memory *int64 `json:"Memory,omitempty" name:"Memory"`
	// cpu核数

	Cpu *int64 `json:"Cpu,omitempty" name:"Cpu"`
	// 管理员账号的密码

	AdminPassword *string `json:"AdminPassword,omitempty" name:"AdminPassword"`
	// 字符集，根据DescribeAvailableCharset查询

	Charset *string `json:"Charset,omitempty" name:"Charset"`
	// 创建数量，(0,10]

	Count *int64 `json:"Count,omitempty" name:"Count"`
	// 实例名称

	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`
	// 实例类型，当前只支持master类型

	InstanceRole *string `json:"InstanceRole,omitempty" name:"InstanceRole"`
	// 主节点可用区

	MasterNodeZone *string `json:"MasterNodeZone,omitempty" name:"MasterNodeZone"`
	// 从节点可用区

	FollowNodesZones []*string `json:"FollowNodesZones,omitempty" name:"FollowNodesZones"`
	// 要绑定的安全组

	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" name:"SecurityGroupIdList"`
	// 节点规格信息

	PgNode *PgNodeInfo `json:"PgNode,omitempty" name:"PgNode"`
	// 项目ID

	ProjectId *string `json:"ProjectId,omitempty" name:"ProjectId"`
}

func (r *CreatePGInstanceHourRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreatePGInstanceHourRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAvailableCharsetResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 支持的字符集列表

		AvailableCharsets []*string `json:"AvailableCharsets,omitempty" name:"AvailableCharsets"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAvailableCharsetResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAvailableCharsetResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeResourceOverviewResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 地域英文 ID

		Region *string `json:"Region,omitempty" name:"Region"`
		// region id

		RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`
		// 地域名称

		RegionName *string `json:"RegionName,omitempty" name:"RegionName"`
		// 物理服务器数量

		ServerCount *int64 `json:"ServerCount,omitempty" name:"ServerCount"`
		// 实例统计数据

		Instances *InstanceOverview `json:"Instances,omitempty" name:"Instances"`
		// CPU 使用详情

		Cpu *ResourceUsage `json:"Cpu,omitempty" name:"Cpu"`
		// 内存使用详情

		Memory *ResourceUsage `json:"Memory,omitempty" name:"Memory"`
		// 存储使用详情

		Storage *ResourceUsage `json:"Storage,omitempty" name:"Storage"`
		// 产品名称

		ProductName *string `json:"ProductName,omitempty" name:"ProductName"`
		// 接口版本

		Version *string `json:"Version,omitempty" name:"Version"`
		// 服务器上架时标记为 ALL 或 DN 节点的数量统计

		Data []*ProductLicenseData `json:"Data,omitempty" name:"Data"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeResourceOverviewResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeResourceOverviewResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ModifyInstanceSyncTypeRequest struct {
	*tchttp.BaseRequest

	// 实例id。

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 需要修改复制类型的实例节点信息（可选修改的复制类型包括：同步不可退化、同步可退化、异步）。

	ModifySyncTypeNodes []*SyncTypeInfo `json:"ModifySyncTypeNodes,omitempty" name:"ModifySyncTypeNodes"`
}

func (r *ModifyInstanceSyncTypeRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ModifyInstanceSyncTypeRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CreateBackupResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 任务Id，用于查询手动备份任务是否完成

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateBackupResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *CreateBackupResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SetBackupRulesResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 异步任务id。

		TaskId *int64 `json:"TaskId,omitempty" name:"TaskId"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetBackupRulesResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SetBackupRulesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type VpcInfo struct {

	// 私有网络Id

	VpcId *string `json:"VpcId,omitempty" name:"VpcId"`
	// 私有网络子网Id

	SubnetId *string `json:"SubnetId,omitempty" name:"SubnetId"`
	// 私有网络名称

	VpcName *string `json:"VpcName,omitempty" name:"VpcName"`
	// 私有网络子网名称

	SubnetName *string `json:"SubnetName,omitempty" name:"SubnetName"`
	// vip

	Vip *string `json:"Vip,omitempty" name:"Vip"`
	// vport

	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type DescribeAccountsRequest struct {
	*tchttp.BaseRequest

	// 实例Id

	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
	// 返回数量，默认为20，最大值为100。

	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
	// 偏移量，默认为0。

	Offset *int64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeAccountsRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeAccountsRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeDbNameListResponse struct {
	*tchttp.BaseResponse

	Response *struct {
		// 符合条件的db名称数量。

		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`
		// 符合条件的db名称列表。

		DbNameList []*string `json:"DbNameList,omitempty" name:"DbNameList"`
		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeDbNameListResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DescribeDbNameListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type Inbound struct {

	// 策略，ACCEPT或者DROP。

	Action *string `json:"Action,omitempty" name:"Action"`
	// 地址组id代表的地址集合。

	AddressModule *string `json:"AddressModule,omitempty" name:"AddressModule"`
	// 来源Ip或Ip段，例如192.168.0.0/16。

	CidrIp *string `json:"CidrIp,omitempty" name:"CidrIp"`
	// 描述。

	Desc *string `json:"Desc,omitempty" name:"Desc"`
	// 安全组id代表的地址集合。

	Id *string `json:"Id,omitempty" name:"Id"`
	// 网络协议，支持udp、tcp等。

	IpProtocol *string `json:"IpProtocol,omitempty" name:"IpProtocol"`
	// 端口。

	PortRange *string `json:"PortRange,omitempty" name:"PortRange"`
	// 服务组id代表的协议和端口集合。

	ServiceModule *string `json:"ServiceModule,omitempty" name:"ServiceModule"`
}
